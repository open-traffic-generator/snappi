package gosnappi

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type grpcTransport struct {
	clientConnection *grpc.ClientConn
	location         string
	requestTimeout   time.Duration
	dialTimeout      time.Duration
}

type GrpcTransport interface {
	// SetLocation set client connection to the given grpc target
	SetLocation(value string) GrpcTransport
	// Location get grpc target
	Location() string
	// SetRequestTimeout set timeout in grpc request
	SetRequestTimeout(value time.Duration) GrpcTransport
	// RequestTimeout get timeout in grpc request
	RequestTimeout() time.Duration
	// SetDialTimeout set timeout in grpc dial
	SetDialTimeout(value time.Duration) GrpcTransport
	// DialTimeout get timeout in grpc dial
	DialTimeout() time.Duration
	// SetClientConnection set grpc DialContext
	// SetClientConnection and (SetLocation, SetDialTimeout) are mutually exclusive
	SetClientConnection(con *grpc.ClientConn) GrpcTransport
	// ClientConnection get grpc DialContext
	ClientConnection() *grpc.ClientConn
}

// Location
func (obj *grpcTransport) Location() string {
	return obj.location
}

// SetLocation
func (obj *grpcTransport) SetLocation(value string) GrpcTransport {
	obj.location = value
	return obj
}

// RequestTimeout returns the grpc request timeout in seconds
func (obj *grpcTransport) RequestTimeout() time.Duration {
	return obj.requestTimeout
}

// SetRequestTimeout contains the timeout value in seconds for a grpc request
func (obj *grpcTransport) SetRequestTimeout(value time.Duration) GrpcTransport {
	obj.requestTimeout = value
	return obj
}
func (obj *grpcTransport) DialTimeout() time.Duration {
	return obj.dialTimeout
}

func (obj *grpcTransport) SetDialTimeout(value time.Duration) GrpcTransport {
	obj.dialTimeout = value
	return obj
}

func (obj *grpcTransport) ClientConnection() *grpc.ClientConn {
	return obj.clientConnection
}

func (obj *grpcTransport) SetClientConnection(con *grpc.ClientConn) GrpcTransport {
	obj.clientConnection = con
	return obj
}

type httpTransport struct {
	location string
	verify   bool
	conn     net.Conn
}

type HttpTransport interface {
	SetLocation(value string) HttpTransport
	Location() string
	SetVerify(value bool) HttpTransport
	Verify() bool
}

// Location
func (obj *httpTransport) Location() string {
	return obj.location
}

// SetLocation
func (obj *httpTransport) SetLocation(value string) HttpTransport {
	obj.location = value
	return obj
}

// Verify returns whether or not TLS certificates will be verified by the server
func (obj *httpTransport) Verify() bool {
	return obj.verify
}

// SetVerify determines whether or not TLS certificates will be verified by the server
func (obj *httpTransport) SetVerify(value bool) HttpTransport {
	obj.verify = value
	return obj
}

type api struct {
	grpc     *grpcTransport
	http     *httpTransport
	tracer   *telemetry
	warnings string
}

type Api interface {
	Telemetry() Telemetry
	NewGrpcTransport() GrpcTransport
	hasGrpcTransport() bool
	NewHttpTransport() HttpTransport
	hasHttpTransport() bool
	Close() error
	// Warnings Api is only for testing purpose
	// and not intended to use in production
	Warnings() string
	deprecated(message string)
	under_review(message string)
	addWarnings(message string)
	fromHttpError(statusCode int, body []byte) Error
	fromGrpcError(err error) (Error, bool)
	FromError(err error) (Error, bool)
}

// NewGrpcTransport sets the underlying transport of the Api as grpc
func (api *api) NewGrpcTransport() GrpcTransport {
	api.grpc = &grpcTransport{
		location:       "localhost:5050",
		requestTimeout: 10 * time.Second,
		dialTimeout:    10 * time.Second,
	}
	api.http = nil
	return api.grpc
}

// HasGrpcTransport will return True for gRPC transport
func (api *api) hasGrpcTransport() bool {
	return api.grpc != nil
}

// NewHttpTransport sets the underlying transport of the Api as http
func (api *api) NewHttpTransport() HttpTransport {
	api.http = &httpTransport{
		location: "https://localhost:443",
		verify:   false,
	}
	if api.grpc != nil {
		if api.grpc.clientConnection != nil {
			api.grpc.clientConnection.Close()
		}
	}
	api.grpc = nil
	return api.http
}

func (api *api) hasHttpTransport() bool {
	return api.http != nil
}

func (api *api) Warnings() string {
	return api.warnings
}

func (api *api) addWarnings(message string) {
	fmt.Printf("[WARNING]: %s\n", message)
	api.warnings = message
}

func (api *api) deprecated(message string) {
	api.warnings = message
	fmt.Printf("warning: %s\n", message)
}

func (api *api) under_review(message string) {
	api.warnings = message
	fmt.Printf("warning: %s\n", message)
}

func (api *api) FromError(err error) (Error, bool) {
	if rErr, ok := err.(Error); ok {
		return rErr, true
	}

	rErr := NewError()
	if err := rErr.FromJson(err.Error()); err == nil {
		return rErr, true
	}

	return api.fromGrpcError(err)
}

func (api *api) setResponseErr(obj Error, code int32, message string) {
	errors := []string{}
	errors = append(errors, message)
	obj.Msg().Code = code
	obj.Msg().Errors = errors
}

func (api *api) fromGrpcError(err error) (Error, bool) {
	st, ok := status.FromError(err)
	if ok {
		rErr := NewError()
		if err := rErr.FromJson(st.Message()); err == nil {
			rErr.Msg().Code = int32(st.Code())
			return rErr, true
		}

		api.setResponseErr(rErr, int32(st.Code()), st.Message())
		return rErr, true
	}

	return nil, false
}

func (api *api) fromHttpError(statusCode int, body []byte) Error {
	rErr := NewError()
	bStr := string(body)
	if err := rErr.FromJson(bStr); err == nil {
		return rErr
	}

	api.setResponseErr(rErr, int32(statusCode), bStr)

	return rErr
}

// Returns instance of telemetry operations
func (api *api) Telemetry() Telemetry {
	return api.tracer
}

// HttpRequestDoer will return True for HTTP transport
type httpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type httpClient struct {
	client httpRequestDoer
	ctx    context.Context
}

// All methods that perform validation will add errors here
// All api rpcs MUST call Validate
type Constraints interface {
	ValueOf(name string) interface{}
}

type validation struct {
	validationErrors []string
	warnings         []string
	constraints      map[string]map[string]Constraints
}

type Validation interface {
	validationResult() error
	deprecated(message string)
	under_review(message string)
	Warnings() []string
	addWarnings(message string)
}

func (obj *validation) validationResult() error {
	obj.constraints = make(map[string]map[string]Constraints)
	if len(obj.validationErrors) > 0 {
		errors := strings.Join(obj.validationErrors, "\n")
		obj.validationErrors = nil
		return fmt.Errorf(errors)
	}
	return nil
}

func (obj *validation) Warnings() []string {
	if len(obj.warnings) > 0 {
		warns := obj.warnings
		obj.warnings = nil
		return warns
	}
	return obj.warnings
}

func (obj *validation) addWarnings(message string) {
	fmt.Printf("[WARNING]: %s\n", message)
	obj.warnings = append(obj.warnings, message)
}

func (obj *validation) deprecated(message string) {
	fmt.Printf("warning: %s\n", message)
	obj.warnings = append(obj.warnings, message)
}

func (obj *validation) under_review(message string) {
	fmt.Printf("warning: %s\n", message)
	obj.warnings = append(obj.warnings, message)
}

func (obj *validation) validateMac(mac string) error {
	macSlice := strings.Split(mac, ":")
	if len(macSlice) != 6 {
		return fmt.Errorf(fmt.Sprintf("Invalid Mac address %s", mac))
	}
	octInd := []string{"0th", "1st", "2nd", "3rd", "4th", "5th"}
	for ind, val := range macSlice {
		num, err := strconv.ParseUint(val, 16, 32)
		if err != nil || num > 255 {
			return fmt.Errorf(fmt.Sprintf("Invalid Mac address at %s octet in %s mac", octInd[ind], mac))
		}
	}
	return nil
}

func (obj *validation) validateIpv4(ip string) error {
	ipSlice := strings.Split(ip, ".")
	if len(ipSlice) != 4 {
		return fmt.Errorf(fmt.Sprintf("Invalid Ipv4 address %s", ip))
	}
	octInd := []string{"1st", "2nd", "3rd", "4th"}
	for ind, val := range ipSlice {
		num, err := strconv.ParseUint(val, 10, 32)
		if err != nil || num > 255 {
			return fmt.Errorf(fmt.Sprintf("Invalid Ipv4 address at %s octet in %s ipv4", octInd[ind], ip))
		}
	}
	return nil
}

func (obj *validation) validateIpv6(ip string) error {
	ip = strings.Trim(ip, " \t")
	if strings.Count(ip, " ") > 0 || strings.Count(ip, ":") > 7 ||
		strings.Count(ip, "::") > 1 || strings.Count(ip, ":::") > 0 ||
		strings.Count(ip, ":") == 0 {
		return fmt.Errorf(fmt.Sprintf("Invalid ipv6 address %s", ip))
	}
	if (string(ip[0]) == ":" && string(ip[:2]) != "::") || (string(ip[len(ip)-1]) == ":" && string(ip[len(ip)-2:]) != "::") {
		return fmt.Errorf(fmt.Sprintf("Invalid ipv6 address %s", ip))
	}
	if strings.Count(ip, "::") == 0 && strings.Count(ip, ":") != 7 {
		return fmt.Errorf(fmt.Sprintf("Invalid ipv6 address %s", ip))
	}
	if ip == "::" {
		return nil
	}
	if ip[:2] == "::" {
		r := strings.NewReplacer("::", "0:")
		ip = r.Replace(ip)
	} else if ip[len(ip)-2:] == "::" {
		r := strings.NewReplacer("::", ":0")
		ip = r.Replace(ip)
	} else {
		r := strings.NewReplacer("::", ":0:")
		ip = r.Replace(ip)
	}
	octInd := []string{"1st", "2nd", "3rd", "4th", "5th", "6th", "7th", "8th"}

	ipSlice := strings.Split(ip, ":")

	for ind, val := range ipSlice {
		num, err := strconv.ParseUint(val, 16, 64)
		if err != nil || num > 65535 {
			return fmt.Errorf(fmt.Sprintf("Invalid Ipv6 address at %s octet in %s ipv6", octInd[ind], ip))
		}
	}

	return nil
}

func (obj *validation) validateHex(hex string) error {
	matched, err := regexp.MatchString(`^[0-9a-fA-F]+$|^0[x|X][0-9a-fA-F]+$`, hex)
	if err != nil || !matched {
		return fmt.Errorf(fmt.Sprintf("Invalid hex value %s", hex))
	}
	return nil
}

func (obj *validation) validateSlice(valSlice []string, sliceType string) error {
	indices := []string{}
	var err error
	for i, val := range valSlice {
		if sliceType == "mac" {
			err = obj.validateMac(val)
		} else if sliceType == "ipv4" {
			err = obj.validateIpv4(val)

		} else if sliceType == "ipv6" {
			err = obj.validateIpv6(val)
		} else if sliceType == "hex" {
			err = obj.validateHex(val)
		} else {
			return fmt.Errorf(fmt.Sprintf("Invalid slice type received <%s>", sliceType))
		}

		if err != nil {
			indices = append(indices, fmt.Sprintf("%d", i))
		}
	}
	if len(indices) > 0 {
		return fmt.Errorf(
			fmt.Sprintf("Invalid %s addresses at indices %s", sliceType, strings.Join(indices, ",")),
		)
	}
	return nil
}

func (obj *validation) validateMacSlice(mac []string) error {
	return obj.validateSlice(mac, "mac")
}

func (obj *validation) validateIpv4Slice(ip []string) error {
	return obj.validateSlice(ip, "ipv4")
}

func (obj *validation) validateIpv6Slice(ip []string) error {
	return obj.validateSlice(ip, "ipv6")
}

func (obj *validation) validateHexSlice(hex []string) error {
	return obj.validateSlice(hex, "hex")
}

// TODO: restore behavior
// func (obj *validation) createMap(objName string) {
// 	if obj.constraints == nil {
// 		obj.constraints = make(map[string]map[string]Constraints)
// 	}
// 	_, ok := obj.constraints[objName]
// 	if !ok {
// 		obj.constraints[objName] = make(map[string]Constraints)
// 	}
// }

// TODO: restore behavior
// func (obj *validation) isUnique(objectName, value string, object Constraints) bool {
// 	if value == "" {
// 		return true
// 	}

// 	obj.createMap("globals")
// 	_, ok := obj.constraints["globals"][value]
// 	unique := false
// 	if !ok {
// 		obj.constraints["globals"][value] = object
// 		obj.createMap(objectName)
// 		obj.constraints[objectName][value] = object
// 		unique = true
// 	}
// 	return unique
// }

// TODO: restore behavior
// func (obj *validation) validateConstraint(objectName []string, value string) bool {
// 	if value == "" {
// 		return false
// 	}
// 	found := false
// 	for _, object := range objectName {
// 		obj_ := strings.Split(object, ".")
// 		strukt, ok := obj.constraints[obj_[0]]
// 		if !ok {
// 			continue
// 		}
// 		for _, object := range strukt {
// 			intf := object.ValueOf(obj_[1])
// 			if intf == nil {
// 				continue
// 			}
// 			if value == fmt.Sprintf("%v", intf) {
// 				found = true
// 				break
// 			}
// 		}
// 		if found {
// 			break
// 		}
// 	}
// 	return found
// }

func checkClientServerVersionCompatibility(clientVer string, serverVer string, componentName string) error {

	c, err := semver.NewVersion(clientVer)
	if err != nil {
		return fmt.Errorf("client %s version '%s' is not a valid semver", componentName, clientVer)
	}

	s, err := semver.NewConstraint(serverVer)
	if err != nil {
		return fmt.Errorf("server %s version '%s' is not a valid semver constraint", componentName, serverVer)
	}

	err = fmt.Errorf("client %s version '%s' is not semver compatible with server %s version constraint '%s'", componentName, clientVer, componentName, serverVer)
	valid, errs := s.Validate(c)
	if len(errs) != 0 {
		return fmt.Errorf("%v: %v", err, errs)
	}

	if !valid {
		return err
	}

	return nil
}
