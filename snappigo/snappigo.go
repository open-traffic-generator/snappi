// Open Traffic Generator API 0.4.13
// License: MIT

package snappigo

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	snappipb "github.com/open-traffic-generator/snappi/snappigo/snappipb"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

type grpcTransport struct {
	location       string
	requestTimeout time.Duration
}

type GrpcTransport interface {
	SetLocation(value string) GrpcTransport
	Location() string
	SetRequestTimeout(value int) GrpcTransport
	RequestTimeout() int
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
func (obj *grpcTransport) RequestTimeout() int {
	return int(obj.requestTimeout / time.Second)
}

// SetRequestTimeout contains the timeout value in seconds for a grpc request
func (obj *grpcTransport) SetRequestTimeout(value int) GrpcTransport {
	obj.requestTimeout = time.Duration(value) * time.Second
	return obj
}

type httpTransport struct {
	location string
	verify   bool
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
	grpc *grpcTransport
	http *httpTransport
}

type Api interface {
	NewGrpcTransport() GrpcTransport
	NewHttpTransport() HttpTransport
}

// NewGrpcTransport sets the underlying transport of the Api as grpc
func (api *api) NewGrpcTransport() GrpcTransport {
	api.grpc = &grpcTransport{
		location:       "127.0.0.1:5050",
		requestTimeout: time.Duration(10),
	}
	api.http = nil
	return api.grpc
}

// NewHttpTransport sets the underlying transport of the Api as http
func (api *api) NewHttpTransport() HttpTransport {
	api.http = &httpTransport{
		location: "https://127.0.0.1:443",
		verify:   false,
	}
	api.grpc = nil
	return api.http
}

// All methods that perform validation will add errors here
// All api rpcs MUST call Validate
var validation []string

func Validate() {
	if len(validation) > 0 {
		for _, item := range validation {
			fmt.Println(item)
		}
		panic("validation errors")
	}
}

type snappigoApi struct {
	api
	grpcClient snappipb.OpenapiClient
}

// grpcConnect builds up a grpc connection
func (api *snappigoApi) grpcConnect() error {
	if api.grpcClient == nil {
		conn, err := grpc.Dial(api.grpc.location, grpc.WithInsecure())
		if err != nil {
			return err
		}
		api.grpcClient = snappipb.NewOpenapiClient(conn)
	}
	return nil
}

// NewApi returns a new instance of the top level interface hierarchy
func NewApi() *snappigoApi {
	api := snappigoApi{}
	return &api
}

type SnappigoApi interface {
	Api
	NewConfig() Config
	NewTransmitState() TransmitState
	NewLinkState() LinkState
	NewCaptureState() CaptureState
	NewRouteState() RouteState
	NewMetricsRequest() MetricsRequest
	NewCaptureRequest() CaptureRequest
	SetConfig(config Config) error
	UpdateConfig(config Config) error
	SetTransmitState(transmitState TransmitState) error
	SetLinkState(linkState LinkState) error
	SetCaptureState(captureState CaptureState) error
	SetRouteState(routeState RouteState) error
	GetMetrics(metricsRequest MetricsRequest) error
	GetCapture(captureRequest CaptureRequest) error
}

func (api *snappigoApi) NewConfig() Config {
	return &config{obj: &snappipb.Config{}}
}

func (api *snappigoApi) NewTransmitState() TransmitState {
	return &transmitState{obj: &snappipb.TransmitState{}}
}

func (api *snappigoApi) NewLinkState() LinkState {
	return &linkState{obj: &snappipb.LinkState{}}
}

func (api *snappigoApi) NewCaptureState() CaptureState {
	return &captureState{obj: &snappipb.CaptureState{}}
}

func (api *snappigoApi) NewRouteState() RouteState {
	return &routeState{obj: &snappipb.RouteState{}}
}

func (api *snappigoApi) NewMetricsRequest() MetricsRequest {
	return &metricsRequest{obj: &snappipb.MetricsRequest{}}
}

func (api *snappigoApi) NewCaptureRequest() CaptureRequest {
	return &captureRequest{obj: &snappipb.CaptureRequest{}}
}

func (api *snappigoApi) SetConfig(config Config) error {
	if err := api.grpcConnect(); err != nil {
		return err
	}
	request := snappipb.SetConfigRequest{Config: config.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	client, err := api.grpcClient.SetConfig(ctx, &request)
	if err != nil {
		return err
	}
	resp, _ := client.Recv()
	if resp.GetStatusCode_200() == nil {
		return fmt.Errorf("fail")
	}
	return nil
}

func (api *snappigoApi) UpdateConfig(config Config) error {
	if err := api.grpcConnect(); err != nil {
		return err
	}
	request := snappipb.UpdateConfigRequest{Config: config.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	client, err := api.grpcClient.UpdateConfig(ctx, &request)
	if err != nil {
		return err
	}
	resp, _ := client.Recv()
	if resp.GetStatusCode_200() == nil {
		return fmt.Errorf("fail")
	}
	return nil
}

func (api *snappigoApi) SetTransmitState(transmitState TransmitState) error {
	if err := api.grpcConnect(); err != nil {
		return err
	}
	request := snappipb.SetTransmitStateRequest{TransmitState: transmitState.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	client, err := api.grpcClient.SetTransmitState(ctx, &request)
	if err != nil {
		return err
	}
	resp, _ := client.Recv()
	if resp.GetStatusCode_200() == nil {
		return fmt.Errorf("fail")
	}
	return nil
}

func (api *snappigoApi) SetLinkState(linkState LinkState) error {
	if err := api.grpcConnect(); err != nil {
		return err
	}
	request := snappipb.SetLinkStateRequest{LinkState: linkState.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	client, err := api.grpcClient.SetLinkState(ctx, &request)
	if err != nil {
		return err
	}
	resp, _ := client.Recv()
	if resp.GetStatusCode_200() == nil {
		return fmt.Errorf("fail")
	}
	return nil
}

func (api *snappigoApi) SetCaptureState(captureState CaptureState) error {
	if err := api.grpcConnect(); err != nil {
		return err
	}
	request := snappipb.SetCaptureStateRequest{CaptureState: captureState.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	client, err := api.grpcClient.SetCaptureState(ctx, &request)
	if err != nil {
		return err
	}
	resp, _ := client.Recv()
	if resp.GetStatusCode_200() == nil {
		return fmt.Errorf("fail")
	}
	return nil
}

func (api *snappigoApi) SetRouteState(routeState RouteState) error {
	if err := api.grpcConnect(); err != nil {
		return err
	}
	request := snappipb.SetRouteStateRequest{RouteState: routeState.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	client, err := api.grpcClient.SetRouteState(ctx, &request)
	if err != nil {
		return err
	}
	resp, _ := client.Recv()
	if resp.GetStatusCode_200() == nil {
		return fmt.Errorf("fail")
	}
	return nil
}

func (api *snappigoApi) GetMetrics(metricsRequest MetricsRequest) error {
	if err := api.grpcConnect(); err != nil {
		return err
	}
	request := snappipb.GetMetricsRequest{MetricsRequest: metricsRequest.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	client, err := api.grpcClient.GetMetrics(ctx, &request)
	if err != nil {
		return err
	}
	resp, _ := client.Recv()
	if resp.GetStatusCode_200() == nil {
		return fmt.Errorf("fail")
	}
	return nil
}

func (api *snappigoApi) GetCapture(captureRequest CaptureRequest) error {
	if err := api.grpcConnect(); err != nil {
		return err
	}
	request := snappipb.GetCaptureRequest{CaptureRequest: captureRequest.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	client, err := api.grpcClient.GetCapture(ctx, &request)
	if err != nil {
		return err
	}
	resp, _ := client.Recv()
	if resp.GetStatusCode_200() == nil {
		return fmt.Errorf("fail")
	}
	return nil
}

type config struct {
	obj *snappipb.Config
}

func (obj *config) msg() *snappipb.Config {
	return obj.obj
}

func (obj *config) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *config) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Config interface {
	msg() *snappipb.Config
	Yaml() string
	Json() string
	Ports() []Port
	NewPorts() Port
	Lags() []Lag
	NewLags() Lag
	Layer1() []Layer1
	NewLayer1() Layer1
	Captures() []Capture
	NewCaptures() Capture
	Devices() []Device
	NewDevices() Device
	Flows() []Flow
	NewFlows() Flow
	Events() Event
	Options() ConfigOptions
}

func (obj *config) Ports() []Port {
	if obj.obj.Ports == nil {
		obj.obj.Ports = make([]*snappipb.Port, 0)
	}
	values := make([]Port, 0)
	for _, item := range obj.obj.Ports {
		values = append(values, &port{obj: item})
	}
	return values

}

func (obj *config) NewPorts() Port {
	if obj.obj.Ports == nil {
		obj.obj.Ports = make([]*snappipb.Port, 0)
	}
	slice := append(obj.obj.Ports, &snappipb.Port{})
	obj.obj.Ports = slice
	return &port{obj: slice[len(slice)-1]}
}

func (obj *config) Lags() []Lag {
	if obj.obj.Lags == nil {
		obj.obj.Lags = make([]*snappipb.Lag, 0)
	}
	values := make([]Lag, 0)
	for _, item := range obj.obj.Lags {
		values = append(values, &lag{obj: item})
	}
	return values

}

func (obj *config) NewLags() Lag {
	if obj.obj.Lags == nil {
		obj.obj.Lags = make([]*snappipb.Lag, 0)
	}
	slice := append(obj.obj.Lags, &snappipb.Lag{})
	obj.obj.Lags = slice
	return &lag{obj: slice[len(slice)-1]}
}

func (obj *config) Layer1() []Layer1 {
	if obj.obj.Layer1 == nil {
		obj.obj.Layer1 = make([]*snappipb.Layer1, 0)
	}
	values := make([]Layer1, 0)
	for _, item := range obj.obj.Layer1 {
		values = append(values, &layer1{obj: item})
	}
	return values

}

func (obj *config) NewLayer1() Layer1 {
	if obj.obj.Layer1 == nil {
		obj.obj.Layer1 = make([]*snappipb.Layer1, 0)
	}
	slice := append(obj.obj.Layer1, &snappipb.Layer1{})
	obj.obj.Layer1 = slice
	return &layer1{obj: slice[len(slice)-1]}
}

func (obj *config) Captures() []Capture {
	if obj.obj.Captures == nil {
		obj.obj.Captures = make([]*snappipb.Capture, 0)
	}
	values := make([]Capture, 0)
	for _, item := range obj.obj.Captures {
		values = append(values, &capture{obj: item})
	}
	return values

}

func (obj *config) NewCaptures() Capture {
	if obj.obj.Captures == nil {
		obj.obj.Captures = make([]*snappipb.Capture, 0)
	}
	slice := append(obj.obj.Captures, &snappipb.Capture{})
	obj.obj.Captures = slice
	return &capture{obj: slice[len(slice)-1]}
}

func (obj *config) Devices() []Device {
	if obj.obj.Devices == nil {
		obj.obj.Devices = make([]*snappipb.Device, 0)
	}
	values := make([]Device, 0)
	for _, item := range obj.obj.Devices {
		values = append(values, &device{obj: item})
	}
	return values

}

func (obj *config) NewDevices() Device {
	if obj.obj.Devices == nil {
		obj.obj.Devices = make([]*snappipb.Device, 0)
	}
	slice := append(obj.obj.Devices, &snappipb.Device{})
	obj.obj.Devices = slice
	return &device{obj: slice[len(slice)-1]}
}

func (obj *config) Flows() []Flow {
	if obj.obj.Flows == nil {
		obj.obj.Flows = make([]*snappipb.Flow, 0)
	}
	values := make([]Flow, 0)
	for _, item := range obj.obj.Flows {
		values = append(values, &flow{obj: item})
	}
	return values

}

func (obj *config) NewFlows() Flow {
	if obj.obj.Flows == nil {
		obj.obj.Flows = make([]*snappipb.Flow, 0)
	}
	slice := append(obj.obj.Flows, &snappipb.Flow{})
	obj.obj.Flows = slice
	return &flow{obj: slice[len(slice)-1]}
}

func (obj *config) Events() Event {
	if obj.obj.Events == nil {
		obj.obj.Events = &snappipb.Event{}
	}
	return &event{obj: obj.obj.Events}

}

func (obj *config) Options() ConfigOptions {
	if obj.obj.Options == nil {
		obj.obj.Options = &snappipb.ConfigOptions{}
	}
	return &configOptions{obj: obj.obj.Options}

}

type transmitState struct {
	obj *snappipb.TransmitState
}

func (obj *transmitState) msg() *snappipb.TransmitState {
	return obj.obj
}

func (obj *transmitState) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *transmitState) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type TransmitState interface {
	msg() *snappipb.TransmitState
	Yaml() string
	Json() string
	FlowNames() []string
	SetFlowNames(value []string) TransmitState
}

func (obj *transmitState) FlowNames() []string {
	return obj.obj.FlowNames
}

func (obj *transmitState) SetFlowNames(value []string) TransmitState {
	obj.obj.FlowNames = value
	return obj
}

type linkState struct {
	obj *snappipb.LinkState
}

func (obj *linkState) msg() *snappipb.LinkState {
	return obj.obj
}

func (obj *linkState) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *linkState) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type LinkState interface {
	msg() *snappipb.LinkState
	Yaml() string
	Json() string
	PortNames() []string
	SetPortNames(value []string) LinkState
}

func (obj *linkState) PortNames() []string {
	return obj.obj.PortNames
}

func (obj *linkState) SetPortNames(value []string) LinkState {
	obj.obj.PortNames = value
	return obj
}

type captureState struct {
	obj *snappipb.CaptureState
}

func (obj *captureState) msg() *snappipb.CaptureState {
	return obj.obj
}

func (obj *captureState) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *captureState) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type CaptureState interface {
	msg() *snappipb.CaptureState
	Yaml() string
	Json() string
	PortNames() []string
	SetPortNames(value []string) CaptureState
}

func (obj *captureState) PortNames() []string {
	return obj.obj.PortNames
}

func (obj *captureState) SetPortNames(value []string) CaptureState {
	obj.obj.PortNames = value
	return obj
}

type routeState struct {
	obj *snappipb.RouteState
}

func (obj *routeState) msg() *snappipb.RouteState {
	return obj.obj
}

func (obj *routeState) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *routeState) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type RouteState interface {
	msg() *snappipb.RouteState
	Yaml() string
	Json() string
	Names() []string
	SetNames(value []string) RouteState
}

func (obj *routeState) Names() []string {
	return obj.obj.Names
}

func (obj *routeState) SetNames(value []string) RouteState {
	obj.obj.Names = value
	return obj
}

type metricsRequest struct {
	obj *snappipb.MetricsRequest
}

func (obj *metricsRequest) msg() *snappipb.MetricsRequest {
	return obj.obj
}

func (obj *metricsRequest) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *metricsRequest) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type MetricsRequest interface {
	msg() *snappipb.MetricsRequest
	Yaml() string
	Json() string
	Port() PortMetricsRequest
	Flow() FlowMetricsRequest
	Bgpv4() Bgpv4MetricsRequest
	Bgpv6() Bgpv6MetricsRequest
}

func (obj *metricsRequest) Port() PortMetricsRequest {
	if obj.obj.Port == nil {
		obj.obj.Port = &snappipb.PortMetricsRequest{}
	}
	return &portMetricsRequest{obj: obj.obj.Port}

}

func (obj *metricsRequest) Flow() FlowMetricsRequest {
	if obj.obj.Flow == nil {
		obj.obj.Flow = &snappipb.FlowMetricsRequest{}
	}
	return &flowMetricsRequest{obj: obj.obj.Flow}

}

func (obj *metricsRequest) Bgpv4() Bgpv4MetricsRequest {
	if obj.obj.Bgpv4 == nil {
		obj.obj.Bgpv4 = &snappipb.Bgpv4MetricsRequest{}
	}
	return &bgpv4MetricsRequest{obj: obj.obj.Bgpv4}

}

func (obj *metricsRequest) Bgpv6() Bgpv6MetricsRequest {
	if obj.obj.Bgpv6 == nil {
		obj.obj.Bgpv6 = &snappipb.Bgpv6MetricsRequest{}
	}
	return &bgpv6MetricsRequest{obj: obj.obj.Bgpv6}

}

type captureRequest struct {
	obj *snappipb.CaptureRequest
}

func (obj *captureRequest) msg() *snappipb.CaptureRequest {
	return obj.obj
}

func (obj *captureRequest) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *captureRequest) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type CaptureRequest interface {
	msg() *snappipb.CaptureRequest
	Yaml() string
	Json() string
	PortName() string
	SetPortName(value string) CaptureRequest
}

func (obj *captureRequest) PortName() string {
	return obj.obj.PortName
}

func (obj *captureRequest) SetPortName(value string) CaptureRequest {
	obj.obj.PortName = value
	return obj
}

type port struct {
	obj *snappipb.Port
}

func (obj *port) msg() *snappipb.Port {
	return obj.obj
}

func (obj *port) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *port) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Port interface {
	msg() *snappipb.Port
	Yaml() string
	Json() string
	Location() string
	SetLocation(value string) Port
	Name() string
	SetName(value string) Port
}

func (obj *port) Location() string {
	return *obj.obj.Location
}

func (obj *port) SetLocation(value string) Port {
	obj.obj.Location = &value
	return obj
}

func (obj *port) Name() string {
	return obj.obj.Name
}

func (obj *port) SetName(value string) Port {
	obj.obj.Name = value
	return obj
}

type lag struct {
	obj *snappipb.Lag
}

func (obj *lag) msg() *snappipb.Lag {
	return obj.obj
}

func (obj *lag) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *lag) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Lag interface {
	msg() *snappipb.Lag
	Yaml() string
	Json() string
	Ports() []LagPort
	NewPorts() LagPort
	Name() string
	SetName(value string) Lag
}

func (obj *lag) Ports() []LagPort {
	if obj.obj.Ports == nil {
		obj.obj.Ports = make([]*snappipb.LagPort, 0)
	}
	values := make([]LagPort, 0)
	for _, item := range obj.obj.Ports {
		values = append(values, &lagPort{obj: item})
	}
	return values

}

func (obj *lag) NewPorts() LagPort {
	if obj.obj.Ports == nil {
		obj.obj.Ports = make([]*snappipb.LagPort, 0)
	}
	slice := append(obj.obj.Ports, &snappipb.LagPort{})
	obj.obj.Ports = slice
	return &lagPort{obj: slice[len(slice)-1]}
}

func (obj *lag) Name() string {
	return obj.obj.Name
}

func (obj *lag) SetName(value string) Lag {
	obj.obj.Name = value
	return obj
}

type layer1 struct {
	obj *snappipb.Layer1
}

func (obj *layer1) msg() *snappipb.Layer1 {
	return obj.obj
}

func (obj *layer1) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *layer1) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Layer1 interface {
	msg() *snappipb.Layer1
	Yaml() string
	Json() string
	PortNames() []string
	SetPortNames(value []string) Layer1
	Promiscuous() bool
	SetPromiscuous(value bool) Layer1
	Mtu() int32
	SetMtu(value int32) Layer1
	IeeeMediaDefaults() bool
	SetIeeeMediaDefaults(value bool) Layer1
	AutoNegotiate() bool
	SetAutoNegotiate(value bool) Layer1
	AutoNegotiation() Layer1AutoNegotiation
	FlowControl() Layer1FlowControl
	Name() string
	SetName(value string) Layer1
}

func (obj *layer1) PortNames() []string {
	return obj.obj.PortNames
}

func (obj *layer1) SetPortNames(value []string) Layer1 {
	obj.obj.PortNames = value
	return obj
}

func (obj *layer1) Promiscuous() bool {
	return *obj.obj.Promiscuous
}

func (obj *layer1) SetPromiscuous(value bool) Layer1 {
	obj.obj.Promiscuous = &value
	return obj
}

func (obj *layer1) Mtu() int32 {
	return *obj.obj.Mtu
}

func (obj *layer1) SetMtu(value int32) Layer1 {
	obj.obj.Mtu = &value
	return obj
}

func (obj *layer1) IeeeMediaDefaults() bool {
	return *obj.obj.IeeeMediaDefaults
}

func (obj *layer1) SetIeeeMediaDefaults(value bool) Layer1 {
	obj.obj.IeeeMediaDefaults = &value
	return obj
}

func (obj *layer1) AutoNegotiate() bool {
	return *obj.obj.AutoNegotiate
}

func (obj *layer1) SetAutoNegotiate(value bool) Layer1 {
	obj.obj.AutoNegotiate = &value
	return obj
}

func (obj *layer1) AutoNegotiation() Layer1AutoNegotiation {
	if obj.obj.AutoNegotiation == nil {
		obj.obj.AutoNegotiation = &snappipb.Layer1AutoNegotiation{}
	}
	return &layer1AutoNegotiation{obj: obj.obj.AutoNegotiation}

}

func (obj *layer1) FlowControl() Layer1FlowControl {
	if obj.obj.FlowControl == nil {
		obj.obj.FlowControl = &snappipb.Layer1FlowControl{}
	}
	return &layer1FlowControl{obj: obj.obj.FlowControl}

}

func (obj *layer1) Name() string {
	return obj.obj.Name
}

func (obj *layer1) SetName(value string) Layer1 {
	obj.obj.Name = value
	return obj
}

type capture struct {
	obj *snappipb.Capture
}

func (obj *capture) msg() *snappipb.Capture {
	return obj.obj
}

func (obj *capture) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *capture) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Capture interface {
	msg() *snappipb.Capture
	Yaml() string
	Json() string
	PortNames() []string
	SetPortNames(value []string) Capture
	Filters() []CaptureFilter
	NewFilters() CaptureFilter
	Overwrite() bool
	SetOverwrite(value bool) Capture
	PacketSize() int32
	SetPacketSize(value int32) Capture
	Name() string
	SetName(value string) Capture
}

func (obj *capture) PortNames() []string {
	return obj.obj.PortNames
}

func (obj *capture) SetPortNames(value []string) Capture {
	obj.obj.PortNames = value
	return obj
}

func (obj *capture) Filters() []CaptureFilter {
	if obj.obj.Filters == nil {
		obj.obj.Filters = make([]*snappipb.CaptureFilter, 0)
	}
	values := make([]CaptureFilter, 0)
	for _, item := range obj.obj.Filters {
		values = append(values, &captureFilter{obj: item})
	}
	return values

}

func (obj *capture) NewFilters() CaptureFilter {
	if obj.obj.Filters == nil {
		obj.obj.Filters = make([]*snappipb.CaptureFilter, 0)
	}
	slice := append(obj.obj.Filters, &snappipb.CaptureFilter{})
	obj.obj.Filters = slice
	return &captureFilter{obj: slice[len(slice)-1]}
}

func (obj *capture) Overwrite() bool {
	return *obj.obj.Overwrite
}

func (obj *capture) SetOverwrite(value bool) Capture {
	obj.obj.Overwrite = &value
	return obj
}

func (obj *capture) PacketSize() int32 {
	return *obj.obj.PacketSize
}

func (obj *capture) SetPacketSize(value int32) Capture {
	obj.obj.PacketSize = &value
	return obj
}

func (obj *capture) Name() string {
	return obj.obj.Name
}

func (obj *capture) SetName(value string) Capture {
	obj.obj.Name = value
	return obj
}

type device struct {
	obj *snappipb.Device
}

func (obj *device) msg() *snappipb.Device {
	return obj.obj
}

func (obj *device) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *device) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Device interface {
	msg() *snappipb.Device
	Yaml() string
	Json() string
	ContainerName() string
	SetContainerName(value string) Device
	Ethernet() DeviceEthernet
	Name() string
	SetName(value string) Device
}

func (obj *device) ContainerName() string {
	return obj.obj.ContainerName
}

func (obj *device) SetContainerName(value string) Device {
	obj.obj.ContainerName = value
	return obj
}

func (obj *device) Ethernet() DeviceEthernet {
	return &deviceEthernet{obj: obj.obj.Ethernet}
}

func (obj *device) Name() string {
	return obj.obj.Name
}

func (obj *device) SetName(value string) Device {
	obj.obj.Name = value
	return obj
}

type flow struct {
	obj *snappipb.Flow
}

func (obj *flow) msg() *snappipb.Flow {
	return obj.obj
}

func (obj *flow) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flow) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Flow interface {
	msg() *snappipb.Flow
	Yaml() string
	Json() string
	TxRx() FlowTxRx
	Packet() []FlowHeader
	NewPacket() FlowHeader
	Size() FlowSize
	Rate() FlowRate
	Duration() FlowDuration
	Metrics() FlowMetrics
	Name() string
	SetName(value string) Flow
}

func (obj *flow) TxRx() FlowTxRx {
	return &flowTxRx{obj: obj.obj.TxRx}
}

func (obj *flow) Packet() []FlowHeader {
	if obj.obj.Packet == nil {
		obj.obj.Packet = make([]*snappipb.FlowHeader, 0)
	}
	values := make([]FlowHeader, 0)
	for _, item := range obj.obj.Packet {
		values = append(values, &flowHeader{obj: item})
	}
	return values

}

func (obj *flow) NewPacket() FlowHeader {
	if obj.obj.Packet == nil {
		obj.obj.Packet = make([]*snappipb.FlowHeader, 0)
	}
	slice := append(obj.obj.Packet, &snappipb.FlowHeader{})
	obj.obj.Packet = slice
	return &flowHeader{obj: slice[len(slice)-1]}
}

func (obj *flow) Size() FlowSize {
	if obj.obj.Size == nil {
		obj.obj.Size = &snappipb.FlowSize{}
	}
	return &flowSize{obj: obj.obj.Size}

}

func (obj *flow) Rate() FlowRate {
	if obj.obj.Rate == nil {
		obj.obj.Rate = &snappipb.FlowRate{}
	}
	return &flowRate{obj: obj.obj.Rate}

}

func (obj *flow) Duration() FlowDuration {
	if obj.obj.Duration == nil {
		obj.obj.Duration = &snappipb.FlowDuration{}
	}
	return &flowDuration{obj: obj.obj.Duration}

}

func (obj *flow) Metrics() FlowMetrics {
	if obj.obj.Metrics == nil {
		obj.obj.Metrics = &snappipb.FlowMetrics{}
	}
	return &flowMetrics{obj: obj.obj.Metrics}

}

func (obj *flow) Name() string {
	return obj.obj.Name
}

func (obj *flow) SetName(value string) Flow {
	obj.obj.Name = value
	return obj
}

type event struct {
	obj *snappipb.Event
}

func (obj *event) msg() *snappipb.Event {
	return obj.obj
}

func (obj *event) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *event) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Event interface {
	msg() *snappipb.Event
	Yaml() string
	Json() string
	Enable() bool
	SetEnable(value bool) Event
	Link() EventLink
	RxRateThreshold() EventRxRateThreshold
	RouteAdvertiseWithdraw() EventRouteAdvertiseWithdraw
}

func (obj *event) Enable() bool {
	return *obj.obj.Enable
}

func (obj *event) SetEnable(value bool) Event {
	obj.obj.Enable = &value
	return obj
}

func (obj *event) Link() EventLink {
	if obj.obj.Link == nil {
		obj.obj.Link = &snappipb.EventLink{}
	}
	return &eventLink{obj: obj.obj.Link}

}

func (obj *event) RxRateThreshold() EventRxRateThreshold {
	if obj.obj.RxRateThreshold == nil {
		obj.obj.RxRateThreshold = &snappipb.EventRxRateThreshold{}
	}
	return &eventRxRateThreshold{obj: obj.obj.RxRateThreshold}

}

func (obj *event) RouteAdvertiseWithdraw() EventRouteAdvertiseWithdraw {
	if obj.obj.RouteAdvertiseWithdraw == nil {
		obj.obj.RouteAdvertiseWithdraw = &snappipb.EventRouteAdvertiseWithdraw{}
	}
	return &eventRouteAdvertiseWithdraw{obj: obj.obj.RouteAdvertiseWithdraw}

}

type configOptions struct {
	obj *snappipb.ConfigOptions
}

func (obj *configOptions) msg() *snappipb.ConfigOptions {
	return obj.obj
}

func (obj *configOptions) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *configOptions) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type ConfigOptions interface {
	msg() *snappipb.ConfigOptions
	Yaml() string
	Json() string
	PortOptions() PortOptions
}

func (obj *configOptions) PortOptions() PortOptions {
	if obj.obj.PortOptions == nil {
		obj.obj.PortOptions = &snappipb.PortOptions{}
	}
	return &portOptions{obj: obj.obj.PortOptions}

}

type portMetricsRequest struct {
	obj *snappipb.PortMetricsRequest
}

func (obj *portMetricsRequest) msg() *snappipb.PortMetricsRequest {
	return obj.obj
}

func (obj *portMetricsRequest) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *portMetricsRequest) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PortMetricsRequest interface {
	msg() *snappipb.PortMetricsRequest
	Yaml() string
	Json() string
	PortNames() []string
	SetPortNames(value []string) PortMetricsRequest
}

func (obj *portMetricsRequest) PortNames() []string {
	return obj.obj.PortNames
}

func (obj *portMetricsRequest) SetPortNames(value []string) PortMetricsRequest {
	obj.obj.PortNames = value
	return obj
}

type flowMetricsRequest struct {
	obj *snappipb.FlowMetricsRequest
}

func (obj *flowMetricsRequest) msg() *snappipb.FlowMetricsRequest {
	return obj.obj
}

func (obj *flowMetricsRequest) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowMetricsRequest) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowMetricsRequest interface {
	msg() *snappipb.FlowMetricsRequest
	Yaml() string
	Json() string
	FlowNames() []string
	SetFlowNames(value []string) FlowMetricsRequest
	MetricGroups() FlowMetricGroupRequest
}

func (obj *flowMetricsRequest) FlowNames() []string {
	return obj.obj.FlowNames
}

func (obj *flowMetricsRequest) SetFlowNames(value []string) FlowMetricsRequest {
	obj.obj.FlowNames = value
	return obj
}

func (obj *flowMetricsRequest) MetricGroups() FlowMetricGroupRequest {
	if obj.obj.MetricGroups == nil {
		obj.obj.MetricGroups = &snappipb.FlowMetricGroupRequest{}
	}
	return &flowMetricGroupRequest{obj: obj.obj.MetricGroups}

}

type bgpv4MetricsRequest struct {
	obj *snappipb.Bgpv4MetricsRequest
}

func (obj *bgpv4MetricsRequest) msg() *snappipb.Bgpv4MetricsRequest {
	return obj.obj
}

func (obj *bgpv4MetricsRequest) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *bgpv4MetricsRequest) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Bgpv4MetricsRequest interface {
	msg() *snappipb.Bgpv4MetricsRequest
	Yaml() string
	Json() string
	DeviceNames() []string
	SetDeviceNames(value []string) Bgpv4MetricsRequest
}

func (obj *bgpv4MetricsRequest) DeviceNames() []string {
	return obj.obj.DeviceNames
}

func (obj *bgpv4MetricsRequest) SetDeviceNames(value []string) Bgpv4MetricsRequest {
	obj.obj.DeviceNames = value
	return obj
}

type bgpv6MetricsRequest struct {
	obj *snappipb.Bgpv6MetricsRequest
}

func (obj *bgpv6MetricsRequest) msg() *snappipb.Bgpv6MetricsRequest {
	return obj.obj
}

func (obj *bgpv6MetricsRequest) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *bgpv6MetricsRequest) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Bgpv6MetricsRequest interface {
	msg() *snappipb.Bgpv6MetricsRequest
	Yaml() string
	Json() string
	DeviceNames() []string
	SetDeviceNames(value []string) Bgpv6MetricsRequest
}

func (obj *bgpv6MetricsRequest) DeviceNames() []string {
	return obj.obj.DeviceNames
}

func (obj *bgpv6MetricsRequest) SetDeviceNames(value []string) Bgpv6MetricsRequest {
	obj.obj.DeviceNames = value
	return obj
}

type lagPort struct {
	obj *snappipb.LagPort
}

func (obj *lagPort) msg() *snappipb.LagPort {
	return obj.obj
}

func (obj *lagPort) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *lagPort) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type LagPort interface {
	msg() *snappipb.LagPort
	Yaml() string
	Json() string
	PortName() string
	SetPortName(value string) LagPort
	Protocol() LagProtocol
	Ethernet() DeviceEthernetBase
}

func (obj *lagPort) PortName() string {
	return obj.obj.PortName
}

func (obj *lagPort) SetPortName(value string) LagPort {
	obj.obj.PortName = value
	return obj
}

func (obj *lagPort) Protocol() LagProtocol {
	return &lagProtocol{obj: obj.obj.Protocol}
}

func (obj *lagPort) Ethernet() DeviceEthernetBase {
	return &deviceEthernetBase{obj: obj.obj.Ethernet}
}

type layer1AutoNegotiation struct {
	obj *snappipb.Layer1AutoNegotiation
}

func (obj *layer1AutoNegotiation) msg() *snappipb.Layer1AutoNegotiation {
	return obj.obj
}

func (obj *layer1AutoNegotiation) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *layer1AutoNegotiation) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Layer1AutoNegotiation interface {
	msg() *snappipb.Layer1AutoNegotiation
	Yaml() string
	Json() string
	Advertise1000Mbps() bool
	SetAdvertise1000Mbps(value bool) Layer1AutoNegotiation
	Advertise100FdMbps() bool
	SetAdvertise100FdMbps(value bool) Layer1AutoNegotiation
	Advertise100HdMbps() bool
	SetAdvertise100HdMbps(value bool) Layer1AutoNegotiation
	Advertise10FdMbps() bool
	SetAdvertise10FdMbps(value bool) Layer1AutoNegotiation
	Advertise10HdMbps() bool
	SetAdvertise10HdMbps(value bool) Layer1AutoNegotiation
	LinkTraining() bool
	SetLinkTraining(value bool) Layer1AutoNegotiation
	RsFec() bool
	SetRsFec(value bool) Layer1AutoNegotiation
}

func (obj *layer1AutoNegotiation) Advertise1000Mbps() bool {
	return *obj.obj.Advertise1000Mbps
}

func (obj *layer1AutoNegotiation) SetAdvertise1000Mbps(value bool) Layer1AutoNegotiation {
	obj.obj.Advertise1000Mbps = &value
	return obj
}

func (obj *layer1AutoNegotiation) Advertise100FdMbps() bool {
	return *obj.obj.Advertise100FdMbps
}

func (obj *layer1AutoNegotiation) SetAdvertise100FdMbps(value bool) Layer1AutoNegotiation {
	obj.obj.Advertise100FdMbps = &value
	return obj
}

func (obj *layer1AutoNegotiation) Advertise100HdMbps() bool {
	return *obj.obj.Advertise100HdMbps
}

func (obj *layer1AutoNegotiation) SetAdvertise100HdMbps(value bool) Layer1AutoNegotiation {
	obj.obj.Advertise100HdMbps = &value
	return obj
}

func (obj *layer1AutoNegotiation) Advertise10FdMbps() bool {
	return *obj.obj.Advertise10FdMbps
}

func (obj *layer1AutoNegotiation) SetAdvertise10FdMbps(value bool) Layer1AutoNegotiation {
	obj.obj.Advertise10FdMbps = &value
	return obj
}

func (obj *layer1AutoNegotiation) Advertise10HdMbps() bool {
	return *obj.obj.Advertise10HdMbps
}

func (obj *layer1AutoNegotiation) SetAdvertise10HdMbps(value bool) Layer1AutoNegotiation {
	obj.obj.Advertise10HdMbps = &value
	return obj
}

func (obj *layer1AutoNegotiation) LinkTraining() bool {
	return *obj.obj.LinkTraining
}

func (obj *layer1AutoNegotiation) SetLinkTraining(value bool) Layer1AutoNegotiation {
	obj.obj.LinkTraining = &value
	return obj
}

func (obj *layer1AutoNegotiation) RsFec() bool {
	return *obj.obj.RsFec
}

func (obj *layer1AutoNegotiation) SetRsFec(value bool) Layer1AutoNegotiation {
	obj.obj.RsFec = &value
	return obj
}

type layer1FlowControl struct {
	obj *snappipb.Layer1FlowControl
}

func (obj *layer1FlowControl) msg() *snappipb.Layer1FlowControl {
	return obj.obj
}

func (obj *layer1FlowControl) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *layer1FlowControl) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Layer1FlowControl interface {
	msg() *snappipb.Layer1FlowControl
	Yaml() string
	Json() string
	DirectedAddress() string
	SetDirectedAddress(value string) Layer1FlowControl
	Ieee8021qbb() Layer1Ieee8021qbb
	Ieee8023x() Layer1Ieee8023x
}

func (obj *layer1FlowControl) DirectedAddress() string {
	return *obj.obj.DirectedAddress
}

func (obj *layer1FlowControl) SetDirectedAddress(value string) Layer1FlowControl {
	obj.obj.DirectedAddress = &value
	return obj
}

func (obj *layer1FlowControl) Ieee8021qbb() Layer1Ieee8021qbb {
	if obj.obj.Ieee8021qbb == nil {
		obj.obj.Ieee8021qbb = &snappipb.Layer1Ieee8021qbb{}
	}
	return &layer1Ieee8021qbb{obj: obj.obj.Ieee8021qbb}

}

func (obj *layer1FlowControl) Ieee8023x() Layer1Ieee8023x {
	if obj.obj.Ieee8023x == nil {
		obj.obj.Ieee8023x = &snappipb.Layer1Ieee8023x{}
	}
	return &layer1Ieee8023x{obj: obj.obj.Ieee8023x}

}

type captureFilter struct {
	obj *snappipb.CaptureFilter
}

func (obj *captureFilter) msg() *snappipb.CaptureFilter {
	return obj.obj
}

func (obj *captureFilter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *captureFilter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type CaptureFilter interface {
	msg() *snappipb.CaptureFilter
	Yaml() string
	Json() string
	Custom() CaptureCustom
	Ethernet() CaptureEthernet
	Vlan() CaptureVlan
	Ipv4() CaptureIpv4
	Ipv6() CaptureIpv6
}

func (obj *captureFilter) Custom() CaptureCustom {
	if obj.obj.Custom == nil {
		obj.obj.Custom = &snappipb.CaptureCustom{}
	}
	return &captureCustom{obj: obj.obj.Custom}

}

func (obj *captureFilter) Ethernet() CaptureEthernet {
	if obj.obj.Ethernet == nil {
		obj.obj.Ethernet = &snappipb.CaptureEthernet{}
	}
	return &captureEthernet{obj: obj.obj.Ethernet}

}

func (obj *captureFilter) Vlan() CaptureVlan {
	if obj.obj.Vlan == nil {
		obj.obj.Vlan = &snappipb.CaptureVlan{}
	}
	return &captureVlan{obj: obj.obj.Vlan}

}

func (obj *captureFilter) Ipv4() CaptureIpv4 {
	if obj.obj.Ipv4 == nil {
		obj.obj.Ipv4 = &snappipb.CaptureIpv4{}
	}
	return &captureIpv4{obj: obj.obj.Ipv4}

}

func (obj *captureFilter) Ipv6() CaptureIpv6 {
	if obj.obj.Ipv6 == nil {
		obj.obj.Ipv6 = &snappipb.CaptureIpv6{}
	}
	return &captureIpv6{obj: obj.obj.Ipv6}

}

type deviceEthernet struct {
	obj *snappipb.DeviceEthernet
}

func (obj *deviceEthernet) msg() *snappipb.DeviceEthernet {
	return obj.obj
}

func (obj *deviceEthernet) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceEthernet) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceEthernet interface {
	msg() *snappipb.DeviceEthernet
	Yaml() string
	Json() string
	Ipv4() DeviceIpv4
	Ipv6() DeviceIpv6
	Mac() string
	SetMac(value string) DeviceEthernet
	Mtu() int32
	SetMtu(value int32) DeviceEthernet
	Vlans() []DeviceVlan
	NewVlans() DeviceVlan
	Name() string
	SetName(value string) DeviceEthernet
}

func (obj *deviceEthernet) Ipv4() DeviceIpv4 {
	if obj.obj.Ipv4 == nil {
		obj.obj.Ipv4 = &snappipb.DeviceIpv4{}
	}
	return &deviceIpv4{obj: obj.obj.Ipv4}

}

func (obj *deviceEthernet) Ipv6() DeviceIpv6 {
	if obj.obj.Ipv6 == nil {
		obj.obj.Ipv6 = &snappipb.DeviceIpv6{}
	}
	return &deviceIpv6{obj: obj.obj.Ipv6}

}

func (obj *deviceEthernet) Mac() string {
	return obj.obj.Mac
}

func (obj *deviceEthernet) SetMac(value string) DeviceEthernet {
	obj.obj.Mac = value
	return obj
}

func (obj *deviceEthernet) Mtu() int32 {
	return *obj.obj.Mtu
}

func (obj *deviceEthernet) SetMtu(value int32) DeviceEthernet {
	obj.obj.Mtu = &value
	return obj
}

func (obj *deviceEthernet) Vlans() []DeviceVlan {
	if obj.obj.Vlans == nil {
		obj.obj.Vlans = make([]*snappipb.DeviceVlan, 0)
	}
	values := make([]DeviceVlan, 0)
	for _, item := range obj.obj.Vlans {
		values = append(values, &deviceVlan{obj: item})
	}
	return values

}

func (obj *deviceEthernet) NewVlans() DeviceVlan {
	if obj.obj.Vlans == nil {
		obj.obj.Vlans = make([]*snappipb.DeviceVlan, 0)
	}
	slice := append(obj.obj.Vlans, &snappipb.DeviceVlan{})
	obj.obj.Vlans = slice
	return &deviceVlan{obj: slice[len(slice)-1]}
}

func (obj *deviceEthernet) Name() string {
	return obj.obj.Name
}

func (obj *deviceEthernet) SetName(value string) DeviceEthernet {
	obj.obj.Name = value
	return obj
}

type flowTxRx struct {
	obj *snappipb.FlowTxRx
}

func (obj *flowTxRx) msg() *snappipb.FlowTxRx {
	return obj.obj
}

func (obj *flowTxRx) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowTxRx) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowTxRx interface {
	msg() *snappipb.FlowTxRx
	Yaml() string
	Json() string
	Port() FlowPort
	Device() FlowDevice
}

func (obj *flowTxRx) Port() FlowPort {
	if obj.obj.Port == nil {
		obj.obj.Port = &snappipb.FlowPort{}
	}
	return &flowPort{obj: obj.obj.Port}

}

func (obj *flowTxRx) Device() FlowDevice {
	if obj.obj.Device == nil {
		obj.obj.Device = &snappipb.FlowDevice{}
	}
	return &flowDevice{obj: obj.obj.Device}

}

type flowHeader struct {
	obj *snappipb.FlowHeader
}

func (obj *flowHeader) msg() *snappipb.FlowHeader {
	return obj.obj
}

func (obj *flowHeader) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowHeader) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowHeader interface {
	msg() *snappipb.FlowHeader
	Yaml() string
	Json() string
	Custom() FlowCustom
	Ethernet() FlowEthernet
	Vlan() FlowVlan
	Vxlan() FlowVxlan
	Ipv4() FlowIpv4
	Ipv6() FlowIpv6
	Pfcpause() FlowPfcPause
	Ethernetpause() FlowEthernetPause
	Tcp() FlowTcp
	Udp() FlowUdp
	Gre() FlowGre
	Gtpv1() FlowGtpv1
	Gtpv2() FlowGtpv2
	Arp() FlowArp
	Icmp() FlowIcmp
	Icmpv6() FlowIcmpv6
	Ppp() FlowPpp
	Igmpv1() FlowIgmpv1
}

func (obj *flowHeader) Custom() FlowCustom {
	if obj.obj.Custom == nil {
		obj.obj.Custom = &snappipb.FlowCustom{}
	}
	return &flowCustom{obj: obj.obj.Custom}

}

func (obj *flowHeader) Ethernet() FlowEthernet {
	if obj.obj.Ethernet == nil {
		obj.obj.Ethernet = &snappipb.FlowEthernet{}
	}
	return &flowEthernet{obj: obj.obj.Ethernet}

}

func (obj *flowHeader) Vlan() FlowVlan {
	if obj.obj.Vlan == nil {
		obj.obj.Vlan = &snappipb.FlowVlan{}
	}
	return &flowVlan{obj: obj.obj.Vlan}

}

func (obj *flowHeader) Vxlan() FlowVxlan {
	if obj.obj.Vxlan == nil {
		obj.obj.Vxlan = &snappipb.FlowVxlan{}
	}
	return &flowVxlan{obj: obj.obj.Vxlan}

}

func (obj *flowHeader) Ipv4() FlowIpv4 {
	if obj.obj.Ipv4 == nil {
		obj.obj.Ipv4 = &snappipb.FlowIpv4{}
	}
	return &flowIpv4{obj: obj.obj.Ipv4}

}

func (obj *flowHeader) Ipv6() FlowIpv6 {
	if obj.obj.Ipv6 == nil {
		obj.obj.Ipv6 = &snappipb.FlowIpv6{}
	}
	return &flowIpv6{obj: obj.obj.Ipv6}

}

func (obj *flowHeader) Pfcpause() FlowPfcPause {
	if obj.obj.Pfcpause == nil {
		obj.obj.Pfcpause = &snappipb.FlowPfcPause{}
	}
	return &flowPfcPause{obj: obj.obj.Pfcpause}

}

func (obj *flowHeader) Ethernetpause() FlowEthernetPause {
	if obj.obj.Ethernetpause == nil {
		obj.obj.Ethernetpause = &snappipb.FlowEthernetPause{}
	}
	return &flowEthernetPause{obj: obj.obj.Ethernetpause}

}

func (obj *flowHeader) Tcp() FlowTcp {
	if obj.obj.Tcp == nil {
		obj.obj.Tcp = &snappipb.FlowTcp{}
	}
	return &flowTcp{obj: obj.obj.Tcp}

}

func (obj *flowHeader) Udp() FlowUdp {
	if obj.obj.Udp == nil {
		obj.obj.Udp = &snappipb.FlowUdp{}
	}
	return &flowUdp{obj: obj.obj.Udp}

}

func (obj *flowHeader) Gre() FlowGre {
	if obj.obj.Gre == nil {
		obj.obj.Gre = &snappipb.FlowGre{}
	}
	return &flowGre{obj: obj.obj.Gre}

}

func (obj *flowHeader) Gtpv1() FlowGtpv1 {
	if obj.obj.Gtpv1 == nil {
		obj.obj.Gtpv1 = &snappipb.FlowGtpv1{}
	}
	return &flowGtpv1{obj: obj.obj.Gtpv1}

}

func (obj *flowHeader) Gtpv2() FlowGtpv2 {
	if obj.obj.Gtpv2 == nil {
		obj.obj.Gtpv2 = &snappipb.FlowGtpv2{}
	}
	return &flowGtpv2{obj: obj.obj.Gtpv2}

}

func (obj *flowHeader) Arp() FlowArp {
	if obj.obj.Arp == nil {
		obj.obj.Arp = &snappipb.FlowArp{}
	}
	return &flowArp{obj: obj.obj.Arp}

}

func (obj *flowHeader) Icmp() FlowIcmp {
	if obj.obj.Icmp == nil {
		obj.obj.Icmp = &snappipb.FlowIcmp{}
	}
	return &flowIcmp{obj: obj.obj.Icmp}

}

func (obj *flowHeader) Icmpv6() FlowIcmpv6 {
	if obj.obj.Icmpv6 == nil {
		obj.obj.Icmpv6 = &snappipb.FlowIcmpv6{}
	}
	return &flowIcmpv6{obj: obj.obj.Icmpv6}

}

func (obj *flowHeader) Ppp() FlowPpp {
	if obj.obj.Ppp == nil {
		obj.obj.Ppp = &snappipb.FlowPpp{}
	}
	return &flowPpp{obj: obj.obj.Ppp}

}

func (obj *flowHeader) Igmpv1() FlowIgmpv1 {
	if obj.obj.Igmpv1 == nil {
		obj.obj.Igmpv1 = &snappipb.FlowIgmpv1{}
	}
	return &flowIgmpv1{obj: obj.obj.Igmpv1}

}

type flowSize struct {
	obj *snappipb.FlowSize
}

func (obj *flowSize) msg() *snappipb.FlowSize {
	return obj.obj
}

func (obj *flowSize) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowSize) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowSize interface {
	msg() *snappipb.FlowSize
	Yaml() string
	Json() string
	Fixed() int32
	SetFixed(value int32) FlowSize
	Increment() FlowSizeIncrement
	Random() FlowSizeRandom
}

func (obj *flowSize) Fixed() int32 {
	return *obj.obj.Fixed
}

func (obj *flowSize) SetFixed(value int32) FlowSize {
	obj.obj.Fixed = &value
	return obj
}

func (obj *flowSize) Increment() FlowSizeIncrement {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.FlowSizeIncrement{}
	}
	return &flowSizeIncrement{obj: obj.obj.Increment}

}

func (obj *flowSize) Random() FlowSizeRandom {
	if obj.obj.Random == nil {
		obj.obj.Random = &snappipb.FlowSizeRandom{}
	}
	return &flowSizeRandom{obj: obj.obj.Random}

}

type flowRate struct {
	obj *snappipb.FlowRate
}

func (obj *flowRate) msg() *snappipb.FlowRate {
	return obj.obj
}

func (obj *flowRate) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowRate) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowRate interface {
	msg() *snappipb.FlowRate
	Yaml() string
	Json() string
	Pps() int32
	SetPps(value int32) FlowRate
	Bps() int32
	SetBps(value int32) FlowRate
	Kbps() int32
	SetKbps(value int32) FlowRate
	Mbps() int32
	SetMbps(value int32) FlowRate
	Gbps() int32
	SetGbps(value int32) FlowRate
	Percentage() float32
	SetPercentage(value float32) FlowRate
}

func (obj *flowRate) Pps() int32 {
	return *obj.obj.Pps
}

func (obj *flowRate) SetPps(value int32) FlowRate {
	obj.obj.Pps = &value
	return obj
}

func (obj *flowRate) Bps() int32 {
	return *obj.obj.Bps
}

func (obj *flowRate) SetBps(value int32) FlowRate {
	obj.obj.Bps = &value
	return obj
}

func (obj *flowRate) Kbps() int32 {
	return *obj.obj.Kbps
}

func (obj *flowRate) SetKbps(value int32) FlowRate {
	obj.obj.Kbps = &value
	return obj
}

func (obj *flowRate) Mbps() int32 {
	return *obj.obj.Mbps
}

func (obj *flowRate) SetMbps(value int32) FlowRate {
	obj.obj.Mbps = &value
	return obj
}

func (obj *flowRate) Gbps() int32 {
	return *obj.obj.Gbps
}

func (obj *flowRate) SetGbps(value int32) FlowRate {
	obj.obj.Gbps = &value
	return obj
}

func (obj *flowRate) Percentage() float32 {
	return *obj.obj.Percentage
}

func (obj *flowRate) SetPercentage(value float32) FlowRate {
	obj.obj.Percentage = &value
	return obj
}

type flowDuration struct {
	obj *snappipb.FlowDuration
}

func (obj *flowDuration) msg() *snappipb.FlowDuration {
	return obj.obj
}

func (obj *flowDuration) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowDuration) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowDuration interface {
	msg() *snappipb.FlowDuration
	Yaml() string
	Json() string
	FixedPackets() FlowFixedPackets
	FixedSeconds() FlowFixedSeconds
	Burst() FlowBurst
	Continuous() FlowContinuous
}

func (obj *flowDuration) FixedPackets() FlowFixedPackets {
	if obj.obj.FixedPackets == nil {
		obj.obj.FixedPackets = &snappipb.FlowFixedPackets{}
	}
	return &flowFixedPackets{obj: obj.obj.FixedPackets}

}

func (obj *flowDuration) FixedSeconds() FlowFixedSeconds {
	if obj.obj.FixedSeconds == nil {
		obj.obj.FixedSeconds = &snappipb.FlowFixedSeconds{}
	}
	return &flowFixedSeconds{obj: obj.obj.FixedSeconds}

}

func (obj *flowDuration) Burst() FlowBurst {
	if obj.obj.Burst == nil {
		obj.obj.Burst = &snappipb.FlowBurst{}
	}
	return &flowBurst{obj: obj.obj.Burst}

}

func (obj *flowDuration) Continuous() FlowContinuous {
	if obj.obj.Continuous == nil {
		obj.obj.Continuous = &snappipb.FlowContinuous{}
	}
	return &flowContinuous{obj: obj.obj.Continuous}

}

type flowMetrics struct {
	obj *snappipb.FlowMetrics
}

func (obj *flowMetrics) msg() *snappipb.FlowMetrics {
	return obj.obj
}

func (obj *flowMetrics) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowMetrics) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowMetrics interface {
	msg() *snappipb.FlowMetrics
	Yaml() string
	Json() string
	Enable() bool
	SetEnable(value bool) FlowMetrics
	Loss() bool
	SetLoss(value bool) FlowMetrics
	Timestamps() bool
	SetTimestamps(value bool) FlowMetrics
	Latency() FlowLatencyMetrics
}

func (obj *flowMetrics) Enable() bool {
	return *obj.obj.Enable
}

func (obj *flowMetrics) SetEnable(value bool) FlowMetrics {
	obj.obj.Enable = &value
	return obj
}

func (obj *flowMetrics) Loss() bool {
	return *obj.obj.Loss
}

func (obj *flowMetrics) SetLoss(value bool) FlowMetrics {
	obj.obj.Loss = &value
	return obj
}

func (obj *flowMetrics) Timestamps() bool {
	return *obj.obj.Timestamps
}

func (obj *flowMetrics) SetTimestamps(value bool) FlowMetrics {
	obj.obj.Timestamps = &value
	return obj
}

func (obj *flowMetrics) Latency() FlowLatencyMetrics {
	if obj.obj.Latency == nil {
		obj.obj.Latency = &snappipb.FlowLatencyMetrics{}
	}
	return &flowLatencyMetrics{obj: obj.obj.Latency}

}

type eventLink struct {
	obj *snappipb.EventLink
}

func (obj *eventLink) msg() *snappipb.EventLink {
	return obj.obj
}

func (obj *eventLink) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *eventLink) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type EventLink interface {
	msg() *snappipb.EventLink
	Yaml() string
	Json() string
	Enable() bool
	SetEnable(value bool) EventLink
}

func (obj *eventLink) Enable() bool {
	return *obj.obj.Enable
}

func (obj *eventLink) SetEnable(value bool) EventLink {
	obj.obj.Enable = &value
	return obj
}

type eventRxRateThreshold struct {
	obj *snappipb.EventRxRateThreshold
}

func (obj *eventRxRateThreshold) msg() *snappipb.EventRxRateThreshold {
	return obj.obj
}

func (obj *eventRxRateThreshold) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *eventRxRateThreshold) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type EventRxRateThreshold interface {
	msg() *snappipb.EventRxRateThreshold
	Yaml() string
	Json() string
	Enable() bool
	SetEnable(value bool) EventRxRateThreshold
	Threshold() float32
	SetThreshold(value float32) EventRxRateThreshold
}

func (obj *eventRxRateThreshold) Enable() bool {
	return *obj.obj.Enable
}

func (obj *eventRxRateThreshold) SetEnable(value bool) EventRxRateThreshold {
	obj.obj.Enable = &value
	return obj
}

func (obj *eventRxRateThreshold) Threshold() float32 {
	return *obj.obj.Threshold
}

func (obj *eventRxRateThreshold) SetThreshold(value float32) EventRxRateThreshold {
	obj.obj.Threshold = &value
	return obj
}

type eventRouteAdvertiseWithdraw struct {
	obj *snappipb.EventRouteAdvertiseWithdraw
}

func (obj *eventRouteAdvertiseWithdraw) msg() *snappipb.EventRouteAdvertiseWithdraw {
	return obj.obj
}

func (obj *eventRouteAdvertiseWithdraw) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *eventRouteAdvertiseWithdraw) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type EventRouteAdvertiseWithdraw interface {
	msg() *snappipb.EventRouteAdvertiseWithdraw
	Yaml() string
	Json() string
	Enable() bool
	SetEnable(value bool) EventRouteAdvertiseWithdraw
}

func (obj *eventRouteAdvertiseWithdraw) Enable() bool {
	return *obj.obj.Enable
}

func (obj *eventRouteAdvertiseWithdraw) SetEnable(value bool) EventRouteAdvertiseWithdraw {
	obj.obj.Enable = &value
	return obj
}

type portOptions struct {
	obj *snappipb.PortOptions
}

func (obj *portOptions) msg() *snappipb.PortOptions {
	return obj.obj
}

func (obj *portOptions) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *portOptions) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PortOptions interface {
	msg() *snappipb.PortOptions
	Yaml() string
	Json() string
	LocationPreemption() bool
	SetLocationPreemption(value bool) PortOptions
}

func (obj *portOptions) LocationPreemption() bool {
	return *obj.obj.LocationPreemption
}

func (obj *portOptions) SetLocationPreemption(value bool) PortOptions {
	obj.obj.LocationPreemption = &value
	return obj
}

type flowMetricGroupRequest struct {
	obj *snappipb.FlowMetricGroupRequest
}

func (obj *flowMetricGroupRequest) msg() *snappipb.FlowMetricGroupRequest {
	return obj.obj
}

func (obj *flowMetricGroupRequest) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowMetricGroupRequest) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowMetricGroupRequest interface {
	msg() *snappipb.FlowMetricGroupRequest
	Yaml() string
	Json() string
	Ingress() []string
	SetIngress(value []string) FlowMetricGroupRequest
	Egress() []string
	SetEgress(value []string) FlowMetricGroupRequest
}

func (obj *flowMetricGroupRequest) Ingress() []string {
	return obj.obj.Ingress
}

func (obj *flowMetricGroupRequest) SetIngress(value []string) FlowMetricGroupRequest {
	obj.obj.Ingress = value
	return obj
}

func (obj *flowMetricGroupRequest) Egress() []string {
	return obj.obj.Egress
}

func (obj *flowMetricGroupRequest) SetEgress(value []string) FlowMetricGroupRequest {
	obj.obj.Egress = value
	return obj
}

type lagProtocol struct {
	obj *snappipb.LagProtocol
}

func (obj *lagProtocol) msg() *snappipb.LagProtocol {
	return obj.obj
}

func (obj *lagProtocol) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *lagProtocol) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type LagProtocol interface {
	msg() *snappipb.LagProtocol
	Yaml() string
	Json() string
	Lacp() LagLacp
	Static() LagStatic
}

func (obj *lagProtocol) Lacp() LagLacp {
	if obj.obj.Lacp == nil {
		obj.obj.Lacp = &snappipb.LagLacp{}
	}
	return &lagLacp{obj: obj.obj.Lacp}

}

func (obj *lagProtocol) Static() LagStatic {
	if obj.obj.Static == nil {
		obj.obj.Static = &snappipb.LagStatic{}
	}
	return &lagStatic{obj: obj.obj.Static}

}

type deviceEthernetBase struct {
	obj *snappipb.DeviceEthernetBase
}

func (obj *deviceEthernetBase) msg() *snappipb.DeviceEthernetBase {
	return obj.obj
}

func (obj *deviceEthernetBase) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceEthernetBase) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceEthernetBase interface {
	msg() *snappipb.DeviceEthernetBase
	Yaml() string
	Json() string
	Mac() string
	SetMac(value string) DeviceEthernetBase
	Mtu() int32
	SetMtu(value int32) DeviceEthernetBase
	Vlans() []DeviceVlan
	NewVlans() DeviceVlan
	Name() string
	SetName(value string) DeviceEthernetBase
}

func (obj *deviceEthernetBase) Mac() string {
	return obj.obj.Mac
}

func (obj *deviceEthernetBase) SetMac(value string) DeviceEthernetBase {
	obj.obj.Mac = value
	return obj
}

func (obj *deviceEthernetBase) Mtu() int32 {
	return *obj.obj.Mtu
}

func (obj *deviceEthernetBase) SetMtu(value int32) DeviceEthernetBase {
	obj.obj.Mtu = &value
	return obj
}

func (obj *deviceEthernetBase) Vlans() []DeviceVlan {
	if obj.obj.Vlans == nil {
		obj.obj.Vlans = make([]*snappipb.DeviceVlan, 0)
	}
	values := make([]DeviceVlan, 0)
	for _, item := range obj.obj.Vlans {
		values = append(values, &deviceVlan{obj: item})
	}
	return values

}

func (obj *deviceEthernetBase) NewVlans() DeviceVlan {
	if obj.obj.Vlans == nil {
		obj.obj.Vlans = make([]*snappipb.DeviceVlan, 0)
	}
	slice := append(obj.obj.Vlans, &snappipb.DeviceVlan{})
	obj.obj.Vlans = slice
	return &deviceVlan{obj: slice[len(slice)-1]}
}

func (obj *deviceEthernetBase) Name() string {
	return obj.obj.Name
}

func (obj *deviceEthernetBase) SetName(value string) DeviceEthernetBase {
	obj.obj.Name = value
	return obj
}

type layer1Ieee8021qbb struct {
	obj *snappipb.Layer1Ieee8021qbb
}

func (obj *layer1Ieee8021qbb) msg() *snappipb.Layer1Ieee8021qbb {
	return obj.obj
}

func (obj *layer1Ieee8021qbb) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *layer1Ieee8021qbb) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Layer1Ieee8021qbb interface {
	msg() *snappipb.Layer1Ieee8021qbb
	Yaml() string
	Json() string
	PfcDelay() int32
	SetPfcDelay(value int32) Layer1Ieee8021qbb
	PfcClass0() int32
	SetPfcClass0(value int32) Layer1Ieee8021qbb
	PfcClass1() int32
	SetPfcClass1(value int32) Layer1Ieee8021qbb
	PfcClass2() int32
	SetPfcClass2(value int32) Layer1Ieee8021qbb
	PfcClass3() int32
	SetPfcClass3(value int32) Layer1Ieee8021qbb
	PfcClass4() int32
	SetPfcClass4(value int32) Layer1Ieee8021qbb
	PfcClass5() int32
	SetPfcClass5(value int32) Layer1Ieee8021qbb
	PfcClass6() int32
	SetPfcClass6(value int32) Layer1Ieee8021qbb
	PfcClass7() int32
	SetPfcClass7(value int32) Layer1Ieee8021qbb
}

func (obj *layer1Ieee8021qbb) PfcDelay() int32 {
	return *obj.obj.PfcDelay
}

func (obj *layer1Ieee8021qbb) SetPfcDelay(value int32) Layer1Ieee8021qbb {
	obj.obj.PfcDelay = &value
	return obj
}

func (obj *layer1Ieee8021qbb) PfcClass0() int32 {
	return *obj.obj.PfcClass0
}

func (obj *layer1Ieee8021qbb) SetPfcClass0(value int32) Layer1Ieee8021qbb {
	obj.obj.PfcClass0 = &value
	return obj
}

func (obj *layer1Ieee8021qbb) PfcClass1() int32 {
	return *obj.obj.PfcClass1
}

func (obj *layer1Ieee8021qbb) SetPfcClass1(value int32) Layer1Ieee8021qbb {
	obj.obj.PfcClass1 = &value
	return obj
}

func (obj *layer1Ieee8021qbb) PfcClass2() int32 {
	return *obj.obj.PfcClass2
}

func (obj *layer1Ieee8021qbb) SetPfcClass2(value int32) Layer1Ieee8021qbb {
	obj.obj.PfcClass2 = &value
	return obj
}

func (obj *layer1Ieee8021qbb) PfcClass3() int32 {
	return *obj.obj.PfcClass3
}

func (obj *layer1Ieee8021qbb) SetPfcClass3(value int32) Layer1Ieee8021qbb {
	obj.obj.PfcClass3 = &value
	return obj
}

func (obj *layer1Ieee8021qbb) PfcClass4() int32 {
	return *obj.obj.PfcClass4
}

func (obj *layer1Ieee8021qbb) SetPfcClass4(value int32) Layer1Ieee8021qbb {
	obj.obj.PfcClass4 = &value
	return obj
}

func (obj *layer1Ieee8021qbb) PfcClass5() int32 {
	return *obj.obj.PfcClass5
}

func (obj *layer1Ieee8021qbb) SetPfcClass5(value int32) Layer1Ieee8021qbb {
	obj.obj.PfcClass5 = &value
	return obj
}

func (obj *layer1Ieee8021qbb) PfcClass6() int32 {
	return *obj.obj.PfcClass6
}

func (obj *layer1Ieee8021qbb) SetPfcClass6(value int32) Layer1Ieee8021qbb {
	obj.obj.PfcClass6 = &value
	return obj
}

func (obj *layer1Ieee8021qbb) PfcClass7() int32 {
	return *obj.obj.PfcClass7
}

func (obj *layer1Ieee8021qbb) SetPfcClass7(value int32) Layer1Ieee8021qbb {
	obj.obj.PfcClass7 = &value
	return obj
}

type layer1Ieee8023x struct {
	obj *snappipb.Layer1Ieee8023x
}

func (obj *layer1Ieee8023x) msg() *snappipb.Layer1Ieee8023x {
	return obj.obj
}

func (obj *layer1Ieee8023x) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *layer1Ieee8023x) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Layer1Ieee8023x interface {
	msg() *snappipb.Layer1Ieee8023x
	Yaml() string
	Json() string
}

type captureCustom struct {
	obj *snappipb.CaptureCustom
}

func (obj *captureCustom) msg() *snappipb.CaptureCustom {
	return obj.obj
}

func (obj *captureCustom) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *captureCustom) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type CaptureCustom interface {
	msg() *snappipb.CaptureCustom
	Yaml() string
	Json() string
	Offset() int32
	SetOffset(value int32) CaptureCustom
	BitLength() int32
	SetBitLength(value int32) CaptureCustom
	Value() string
	SetValue(value string) CaptureCustom
	Mask() string
	SetMask(value string) CaptureCustom
	Negate() bool
	SetNegate(value bool) CaptureCustom
}

func (obj *captureCustom) Offset() int32 {
	return *obj.obj.Offset
}

func (obj *captureCustom) SetOffset(value int32) CaptureCustom {
	obj.obj.Offset = &value
	return obj
}

func (obj *captureCustom) BitLength() int32 {
	return *obj.obj.BitLength
}

func (obj *captureCustom) SetBitLength(value int32) CaptureCustom {
	obj.obj.BitLength = &value
	return obj
}

func (obj *captureCustom) Value() string {
	return *obj.obj.Value
}

func (obj *captureCustom) SetValue(value string) CaptureCustom {
	obj.obj.Value = &value
	return obj
}

func (obj *captureCustom) Mask() string {
	return *obj.obj.Mask
}

func (obj *captureCustom) SetMask(value string) CaptureCustom {
	obj.obj.Mask = &value
	return obj
}

func (obj *captureCustom) Negate() bool {
	return *obj.obj.Negate
}

func (obj *captureCustom) SetNegate(value bool) CaptureCustom {
	obj.obj.Negate = &value
	return obj
}

type captureEthernet struct {
	obj *snappipb.CaptureEthernet
}

func (obj *captureEthernet) msg() *snappipb.CaptureEthernet {
	return obj.obj
}

func (obj *captureEthernet) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *captureEthernet) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type CaptureEthernet interface {
	msg() *snappipb.CaptureEthernet
	Yaml() string
	Json() string
	Src() CaptureField
	Dst() CaptureField
	EtherType() CaptureField
	PfcQueue() CaptureField
}

func (obj *captureEthernet) Src() CaptureField {
	if obj.obj.Src == nil {
		obj.obj.Src = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Src}

}

func (obj *captureEthernet) Dst() CaptureField {
	if obj.obj.Dst == nil {
		obj.obj.Dst = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Dst}

}

func (obj *captureEthernet) EtherType() CaptureField {
	if obj.obj.EtherType == nil {
		obj.obj.EtherType = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.EtherType}

}

func (obj *captureEthernet) PfcQueue() CaptureField {
	if obj.obj.PfcQueue == nil {
		obj.obj.PfcQueue = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.PfcQueue}

}

type captureVlan struct {
	obj *snappipb.CaptureVlan
}

func (obj *captureVlan) msg() *snappipb.CaptureVlan {
	return obj.obj
}

func (obj *captureVlan) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *captureVlan) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type CaptureVlan interface {
	msg() *snappipb.CaptureVlan
	Yaml() string
	Json() string
	Priority() CaptureField
	Cfi() CaptureField
	Id() CaptureField
	Protocol() CaptureField
}

func (obj *captureVlan) Priority() CaptureField {
	if obj.obj.Priority == nil {
		obj.obj.Priority = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Priority}

}

func (obj *captureVlan) Cfi() CaptureField {
	if obj.obj.Cfi == nil {
		obj.obj.Cfi = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Cfi}

}

func (obj *captureVlan) Id() CaptureField {
	if obj.obj.Id == nil {
		obj.obj.Id = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Id}

}

func (obj *captureVlan) Protocol() CaptureField {
	if obj.obj.Protocol == nil {
		obj.obj.Protocol = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Protocol}

}

type captureIpv4 struct {
	obj *snappipb.CaptureIpv4
}

func (obj *captureIpv4) msg() *snappipb.CaptureIpv4 {
	return obj.obj
}

func (obj *captureIpv4) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *captureIpv4) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type CaptureIpv4 interface {
	msg() *snappipb.CaptureIpv4
	Yaml() string
	Json() string
	Version() CaptureField
	HeaderLength() CaptureField
	Priority() CaptureField
	TotalLength() CaptureField
	Identification() CaptureField
	Reserved() CaptureField
	DontFragment() CaptureField
	MoreFragments() CaptureField
	FragmentOffset() CaptureField
	TimeToLive() CaptureField
	Protocol() CaptureField
	HeaderChecksum() CaptureField
	Src() CaptureField
	Dst() CaptureField
}

func (obj *captureIpv4) Version() CaptureField {
	if obj.obj.Version == nil {
		obj.obj.Version = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Version}

}

func (obj *captureIpv4) HeaderLength() CaptureField {
	if obj.obj.HeaderLength == nil {
		obj.obj.HeaderLength = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.HeaderLength}

}

func (obj *captureIpv4) Priority() CaptureField {
	if obj.obj.Priority == nil {
		obj.obj.Priority = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Priority}

}

func (obj *captureIpv4) TotalLength() CaptureField {
	if obj.obj.TotalLength == nil {
		obj.obj.TotalLength = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.TotalLength}

}

func (obj *captureIpv4) Identification() CaptureField {
	if obj.obj.Identification == nil {
		obj.obj.Identification = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Identification}

}

func (obj *captureIpv4) Reserved() CaptureField {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Reserved}

}

func (obj *captureIpv4) DontFragment() CaptureField {
	if obj.obj.DontFragment == nil {
		obj.obj.DontFragment = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.DontFragment}

}

func (obj *captureIpv4) MoreFragments() CaptureField {
	if obj.obj.MoreFragments == nil {
		obj.obj.MoreFragments = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.MoreFragments}

}

func (obj *captureIpv4) FragmentOffset() CaptureField {
	if obj.obj.FragmentOffset == nil {
		obj.obj.FragmentOffset = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.FragmentOffset}

}

func (obj *captureIpv4) TimeToLive() CaptureField {
	if obj.obj.TimeToLive == nil {
		obj.obj.TimeToLive = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.TimeToLive}

}

func (obj *captureIpv4) Protocol() CaptureField {
	if obj.obj.Protocol == nil {
		obj.obj.Protocol = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Protocol}

}

func (obj *captureIpv4) HeaderChecksum() CaptureField {
	if obj.obj.HeaderChecksum == nil {
		obj.obj.HeaderChecksum = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.HeaderChecksum}

}

func (obj *captureIpv4) Src() CaptureField {
	if obj.obj.Src == nil {
		obj.obj.Src = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Src}

}

func (obj *captureIpv4) Dst() CaptureField {
	if obj.obj.Dst == nil {
		obj.obj.Dst = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Dst}

}

type captureIpv6 struct {
	obj *snappipb.CaptureIpv6
}

func (obj *captureIpv6) msg() *snappipb.CaptureIpv6 {
	return obj.obj
}

func (obj *captureIpv6) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *captureIpv6) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type CaptureIpv6 interface {
	msg() *snappipb.CaptureIpv6
	Yaml() string
	Json() string
	Version() CaptureField
	TrafficClass() CaptureField
	FlowLabel() CaptureField
	PayloadLength() CaptureField
	NextHeader() CaptureField
	HopLimit() CaptureField
	Src() CaptureField
	Dst() CaptureField
}

func (obj *captureIpv6) Version() CaptureField {
	if obj.obj.Version == nil {
		obj.obj.Version = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Version}

}

func (obj *captureIpv6) TrafficClass() CaptureField {
	if obj.obj.TrafficClass == nil {
		obj.obj.TrafficClass = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.TrafficClass}

}

func (obj *captureIpv6) FlowLabel() CaptureField {
	if obj.obj.FlowLabel == nil {
		obj.obj.FlowLabel = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.FlowLabel}

}

func (obj *captureIpv6) PayloadLength() CaptureField {
	if obj.obj.PayloadLength == nil {
		obj.obj.PayloadLength = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.PayloadLength}

}

func (obj *captureIpv6) NextHeader() CaptureField {
	if obj.obj.NextHeader == nil {
		obj.obj.NextHeader = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.NextHeader}

}

func (obj *captureIpv6) HopLimit() CaptureField {
	if obj.obj.HopLimit == nil {
		obj.obj.HopLimit = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.HopLimit}

}

func (obj *captureIpv6) Src() CaptureField {
	if obj.obj.Src == nil {
		obj.obj.Src = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Src}

}

func (obj *captureIpv6) Dst() CaptureField {
	if obj.obj.Dst == nil {
		obj.obj.Dst = &snappipb.CaptureField{}
	}
	return &captureField{obj: obj.obj.Dst}

}

type deviceIpv4 struct {
	obj *snappipb.DeviceIpv4
}

func (obj *deviceIpv4) msg() *snappipb.DeviceIpv4 {
	return obj.obj
}

func (obj *deviceIpv4) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceIpv4) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceIpv4 interface {
	msg() *snappipb.DeviceIpv4
	Yaml() string
	Json() string
	Gateway() string
	SetGateway(value string) DeviceIpv4
	Address() string
	SetAddress(value string) DeviceIpv4
	Prefix() int32
	SetPrefix(value int32) DeviceIpv4
	Bgpv4() DeviceBgpv4
	Name() string
	SetName(value string) DeviceIpv4
}

func (obj *deviceIpv4) Gateway() string {
	return obj.obj.Gateway
}

func (obj *deviceIpv4) SetGateway(value string) DeviceIpv4 {
	obj.obj.Gateway = value
	return obj
}

func (obj *deviceIpv4) Address() string {
	return obj.obj.Address
}

func (obj *deviceIpv4) SetAddress(value string) DeviceIpv4 {
	obj.obj.Address = value
	return obj
}

func (obj *deviceIpv4) Prefix() int32 {
	return *obj.obj.Prefix
}

func (obj *deviceIpv4) SetPrefix(value int32) DeviceIpv4 {
	obj.obj.Prefix = &value
	return obj
}

func (obj *deviceIpv4) Bgpv4() DeviceBgpv4 {
	if obj.obj.Bgpv4 == nil {
		obj.obj.Bgpv4 = &snappipb.DeviceBgpv4{}
	}
	return &deviceBgpv4{obj: obj.obj.Bgpv4}

}

func (obj *deviceIpv4) Name() string {
	return obj.obj.Name
}

func (obj *deviceIpv4) SetName(value string) DeviceIpv4 {
	obj.obj.Name = value
	return obj
}

type deviceIpv6 struct {
	obj *snappipb.DeviceIpv6
}

func (obj *deviceIpv6) msg() *snappipb.DeviceIpv6 {
	return obj.obj
}

func (obj *deviceIpv6) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceIpv6) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceIpv6 interface {
	msg() *snappipb.DeviceIpv6
	Yaml() string
	Json() string
	Gateway() string
	SetGateway(value string) DeviceIpv6
	Address() string
	SetAddress(value string) DeviceIpv6
	Prefix() int32
	SetPrefix(value int32) DeviceIpv6
	Bgpv6() DeviceBgpv6
	Name() string
	SetName(value string) DeviceIpv6
}

func (obj *deviceIpv6) Gateway() string {
	return obj.obj.Gateway
}

func (obj *deviceIpv6) SetGateway(value string) DeviceIpv6 {
	obj.obj.Gateway = value
	return obj
}

func (obj *deviceIpv6) Address() string {
	return obj.obj.Address
}

func (obj *deviceIpv6) SetAddress(value string) DeviceIpv6 {
	obj.obj.Address = value
	return obj
}

func (obj *deviceIpv6) Prefix() int32 {
	return *obj.obj.Prefix
}

func (obj *deviceIpv6) SetPrefix(value int32) DeviceIpv6 {
	obj.obj.Prefix = &value
	return obj
}

func (obj *deviceIpv6) Bgpv6() DeviceBgpv6 {
	if obj.obj.Bgpv6 == nil {
		obj.obj.Bgpv6 = &snappipb.DeviceBgpv6{}
	}
	return &deviceBgpv6{obj: obj.obj.Bgpv6}

}

func (obj *deviceIpv6) Name() string {
	return obj.obj.Name
}

func (obj *deviceIpv6) SetName(value string) DeviceIpv6 {
	obj.obj.Name = value
	return obj
}

type deviceVlan struct {
	obj *snappipb.DeviceVlan
}

func (obj *deviceVlan) msg() *snappipb.DeviceVlan {
	return obj.obj
}

func (obj *deviceVlan) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceVlan) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceVlan interface {
	msg() *snappipb.DeviceVlan
	Yaml() string
	Json() string
	Priority() int32
	SetPriority(value int32) DeviceVlan
	Id() int32
	SetId(value int32) DeviceVlan
	Name() string
	SetName(value string) DeviceVlan
}

func (obj *deviceVlan) Priority() int32 {
	return *obj.obj.Priority
}

func (obj *deviceVlan) SetPriority(value int32) DeviceVlan {
	obj.obj.Priority = &value
	return obj
}

func (obj *deviceVlan) Id() int32 {
	return *obj.obj.Id
}

func (obj *deviceVlan) SetId(value int32) DeviceVlan {
	obj.obj.Id = &value
	return obj
}

func (obj *deviceVlan) Name() string {
	return obj.obj.Name
}

func (obj *deviceVlan) SetName(value string) DeviceVlan {
	obj.obj.Name = value
	return obj
}

type flowPort struct {
	obj *snappipb.FlowPort
}

func (obj *flowPort) msg() *snappipb.FlowPort {
	return obj.obj
}

func (obj *flowPort) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowPort) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowPort interface {
	msg() *snappipb.FlowPort
	Yaml() string
	Json() string
	TxName() string
	SetTxName(value string) FlowPort
	RxName() string
	SetRxName(value string) FlowPort
}

func (obj *flowPort) TxName() string {
	return obj.obj.TxName
}

func (obj *flowPort) SetTxName(value string) FlowPort {
	obj.obj.TxName = value
	return obj
}

func (obj *flowPort) RxName() string {
	return *obj.obj.RxName
}

func (obj *flowPort) SetRxName(value string) FlowPort {
	obj.obj.RxName = &value
	return obj
}

type flowDevice struct {
	obj *snappipb.FlowDevice
}

func (obj *flowDevice) msg() *snappipb.FlowDevice {
	return obj.obj
}

func (obj *flowDevice) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowDevice) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowDevice interface {
	msg() *snappipb.FlowDevice
	Yaml() string
	Json() string
	TxNames() []string
	SetTxNames(value []string) FlowDevice
	RxNames() []string
	SetRxNames(value []string) FlowDevice
}

func (obj *flowDevice) TxNames() []string {
	return obj.obj.TxNames
}

func (obj *flowDevice) SetTxNames(value []string) FlowDevice {
	obj.obj.TxNames = value
	return obj
}

func (obj *flowDevice) RxNames() []string {
	return obj.obj.RxNames
}

func (obj *flowDevice) SetRxNames(value []string) FlowDevice {
	obj.obj.RxNames = value
	return obj
}

type flowCustom struct {
	obj *snappipb.FlowCustom
}

func (obj *flowCustom) msg() *snappipb.FlowCustom {
	return obj.obj
}

func (obj *flowCustom) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowCustom) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowCustom interface {
	msg() *snappipb.FlowCustom
	Yaml() string
	Json() string
	Bytes() string
	SetBytes(value string) FlowCustom
}

func (obj *flowCustom) Bytes() string {
	return obj.obj.Bytes
}

func (obj *flowCustom) SetBytes(value string) FlowCustom {
	obj.obj.Bytes = value
	return obj
}

type flowEthernet struct {
	obj *snappipb.FlowEthernet
}

func (obj *flowEthernet) msg() *snappipb.FlowEthernet {
	return obj.obj
}

func (obj *flowEthernet) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowEthernet) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowEthernet interface {
	msg() *snappipb.FlowEthernet
	Yaml() string
	Json() string
	Dst() PatternFlowEthernetDst
	Src() PatternFlowEthernetSrc
	EtherType() PatternFlowEthernetEtherType
	PfcQueue() PatternFlowEthernetPfcQueue
}

func (obj *flowEthernet) Dst() PatternFlowEthernetDst {
	if obj.obj.Dst == nil {
		obj.obj.Dst = &snappipb.PatternFlowEthernetDst{}
	}
	return &patternFlowEthernetDst{obj: obj.obj.Dst}

}

func (obj *flowEthernet) Src() PatternFlowEthernetSrc {
	if obj.obj.Src == nil {
		obj.obj.Src = &snappipb.PatternFlowEthernetSrc{}
	}
	return &patternFlowEthernetSrc{obj: obj.obj.Src}

}

func (obj *flowEthernet) EtherType() PatternFlowEthernetEtherType {
	if obj.obj.EtherType == nil {
		obj.obj.EtherType = &snappipb.PatternFlowEthernetEtherType{}
	}
	return &patternFlowEthernetEtherType{obj: obj.obj.EtherType}

}

func (obj *flowEthernet) PfcQueue() PatternFlowEthernetPfcQueue {
	if obj.obj.PfcQueue == nil {
		obj.obj.PfcQueue = &snappipb.PatternFlowEthernetPfcQueue{}
	}
	return &patternFlowEthernetPfcQueue{obj: obj.obj.PfcQueue}

}

type flowVlan struct {
	obj *snappipb.FlowVlan
}

func (obj *flowVlan) msg() *snappipb.FlowVlan {
	return obj.obj
}

func (obj *flowVlan) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowVlan) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowVlan interface {
	msg() *snappipb.FlowVlan
	Yaml() string
	Json() string
	Priority() PatternFlowVlanPriority
	Cfi() PatternFlowVlanCfi
	Id() PatternFlowVlanId
	Tpid() PatternFlowVlanTpid
}

func (obj *flowVlan) Priority() PatternFlowVlanPriority {
	if obj.obj.Priority == nil {
		obj.obj.Priority = &snappipb.PatternFlowVlanPriority{}
	}
	return &patternFlowVlanPriority{obj: obj.obj.Priority}

}

func (obj *flowVlan) Cfi() PatternFlowVlanCfi {
	if obj.obj.Cfi == nil {
		obj.obj.Cfi = &snappipb.PatternFlowVlanCfi{}
	}
	return &patternFlowVlanCfi{obj: obj.obj.Cfi}

}

func (obj *flowVlan) Id() PatternFlowVlanId {
	if obj.obj.Id == nil {
		obj.obj.Id = &snappipb.PatternFlowVlanId{}
	}
	return &patternFlowVlanId{obj: obj.obj.Id}

}

func (obj *flowVlan) Tpid() PatternFlowVlanTpid {
	if obj.obj.Tpid == nil {
		obj.obj.Tpid = &snappipb.PatternFlowVlanTpid{}
	}
	return &patternFlowVlanTpid{obj: obj.obj.Tpid}

}

type flowVxlan struct {
	obj *snappipb.FlowVxlan
}

func (obj *flowVxlan) msg() *snappipb.FlowVxlan {
	return obj.obj
}

func (obj *flowVxlan) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowVxlan) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowVxlan interface {
	msg() *snappipb.FlowVxlan
	Yaml() string
	Json() string
	Flags() PatternFlowVxlanFlags
	Reserved0() PatternFlowVxlanReserved0
	Vni() PatternFlowVxlanVni
	Reserved1() PatternFlowVxlanReserved1
}

func (obj *flowVxlan) Flags() PatternFlowVxlanFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = &snappipb.PatternFlowVxlanFlags{}
	}
	return &patternFlowVxlanFlags{obj: obj.obj.Flags}

}

func (obj *flowVxlan) Reserved0() PatternFlowVxlanReserved0 {
	if obj.obj.Reserved0 == nil {
		obj.obj.Reserved0 = &snappipb.PatternFlowVxlanReserved0{}
	}
	return &patternFlowVxlanReserved0{obj: obj.obj.Reserved0}

}

func (obj *flowVxlan) Vni() PatternFlowVxlanVni {
	if obj.obj.Vni == nil {
		obj.obj.Vni = &snappipb.PatternFlowVxlanVni{}
	}
	return &patternFlowVxlanVni{obj: obj.obj.Vni}

}

func (obj *flowVxlan) Reserved1() PatternFlowVxlanReserved1 {
	if obj.obj.Reserved1 == nil {
		obj.obj.Reserved1 = &snappipb.PatternFlowVxlanReserved1{}
	}
	return &patternFlowVxlanReserved1{obj: obj.obj.Reserved1}

}

type flowIpv4 struct {
	obj *snappipb.FlowIpv4
}

func (obj *flowIpv4) msg() *snappipb.FlowIpv4 {
	return obj.obj
}

func (obj *flowIpv4) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowIpv4) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowIpv4 interface {
	msg() *snappipb.FlowIpv4
	Yaml() string
	Json() string
	Version() PatternFlowIpv4Version
	HeaderLength() PatternFlowIpv4HeaderLength
	Priority() FlowIpv4Priority
	TotalLength() PatternFlowIpv4TotalLength
	Identification() PatternFlowIpv4Identification
	Reserved() PatternFlowIpv4Reserved
	DontFragment() PatternFlowIpv4DontFragment
	MoreFragments() PatternFlowIpv4MoreFragments
	FragmentOffset() PatternFlowIpv4FragmentOffset
	TimeToLive() PatternFlowIpv4TimeToLive
	Protocol() PatternFlowIpv4Protocol
	HeaderChecksum() PatternFlowIpv4HeaderChecksum
	Src() PatternFlowIpv4Src
	Dst() PatternFlowIpv4Dst
}

func (obj *flowIpv4) Version() PatternFlowIpv4Version {
	if obj.obj.Version == nil {
		obj.obj.Version = &snappipb.PatternFlowIpv4Version{}
	}
	return &patternFlowIpv4Version{obj: obj.obj.Version}

}

func (obj *flowIpv4) HeaderLength() PatternFlowIpv4HeaderLength {
	if obj.obj.HeaderLength == nil {
		obj.obj.HeaderLength = &snappipb.PatternFlowIpv4HeaderLength{}
	}
	return &patternFlowIpv4HeaderLength{obj: obj.obj.HeaderLength}

}

func (obj *flowIpv4) Priority() FlowIpv4Priority {
	if obj.obj.Priority == nil {
		obj.obj.Priority = &snappipb.FlowIpv4Priority{}
	}
	return &flowIpv4Priority{obj: obj.obj.Priority}

}

func (obj *flowIpv4) TotalLength() PatternFlowIpv4TotalLength {
	if obj.obj.TotalLength == nil {
		obj.obj.TotalLength = &snappipb.PatternFlowIpv4TotalLength{}
	}
	return &patternFlowIpv4TotalLength{obj: obj.obj.TotalLength}

}

func (obj *flowIpv4) Identification() PatternFlowIpv4Identification {
	if obj.obj.Identification == nil {
		obj.obj.Identification = &snappipb.PatternFlowIpv4Identification{}
	}
	return &patternFlowIpv4Identification{obj: obj.obj.Identification}

}

func (obj *flowIpv4) Reserved() PatternFlowIpv4Reserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = &snappipb.PatternFlowIpv4Reserved{}
	}
	return &patternFlowIpv4Reserved{obj: obj.obj.Reserved}

}

func (obj *flowIpv4) DontFragment() PatternFlowIpv4DontFragment {
	if obj.obj.DontFragment == nil {
		obj.obj.DontFragment = &snappipb.PatternFlowIpv4DontFragment{}
	}
	return &patternFlowIpv4DontFragment{obj: obj.obj.DontFragment}

}

func (obj *flowIpv4) MoreFragments() PatternFlowIpv4MoreFragments {
	if obj.obj.MoreFragments == nil {
		obj.obj.MoreFragments = &snappipb.PatternFlowIpv4MoreFragments{}
	}
	return &patternFlowIpv4MoreFragments{obj: obj.obj.MoreFragments}

}

func (obj *flowIpv4) FragmentOffset() PatternFlowIpv4FragmentOffset {
	if obj.obj.FragmentOffset == nil {
		obj.obj.FragmentOffset = &snappipb.PatternFlowIpv4FragmentOffset{}
	}
	return &patternFlowIpv4FragmentOffset{obj: obj.obj.FragmentOffset}

}

func (obj *flowIpv4) TimeToLive() PatternFlowIpv4TimeToLive {
	if obj.obj.TimeToLive == nil {
		obj.obj.TimeToLive = &snappipb.PatternFlowIpv4TimeToLive{}
	}
	return &patternFlowIpv4TimeToLive{obj: obj.obj.TimeToLive}

}

func (obj *flowIpv4) Protocol() PatternFlowIpv4Protocol {
	if obj.obj.Protocol == nil {
		obj.obj.Protocol = &snappipb.PatternFlowIpv4Protocol{}
	}
	return &patternFlowIpv4Protocol{obj: obj.obj.Protocol}

}

func (obj *flowIpv4) HeaderChecksum() PatternFlowIpv4HeaderChecksum {
	if obj.obj.HeaderChecksum == nil {
		obj.obj.HeaderChecksum = &snappipb.PatternFlowIpv4HeaderChecksum{}
	}
	return &patternFlowIpv4HeaderChecksum{obj: obj.obj.HeaderChecksum}

}

func (obj *flowIpv4) Src() PatternFlowIpv4Src {
	if obj.obj.Src == nil {
		obj.obj.Src = &snappipb.PatternFlowIpv4Src{}
	}
	return &patternFlowIpv4Src{obj: obj.obj.Src}

}

func (obj *flowIpv4) Dst() PatternFlowIpv4Dst {
	if obj.obj.Dst == nil {
		obj.obj.Dst = &snappipb.PatternFlowIpv4Dst{}
	}
	return &patternFlowIpv4Dst{obj: obj.obj.Dst}

}

type flowIpv6 struct {
	obj *snappipb.FlowIpv6
}

func (obj *flowIpv6) msg() *snappipb.FlowIpv6 {
	return obj.obj
}

func (obj *flowIpv6) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowIpv6) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowIpv6 interface {
	msg() *snappipb.FlowIpv6
	Yaml() string
	Json() string
	Version() PatternFlowIpv6Version
	TrafficClass() PatternFlowIpv6TrafficClass
	FlowLabel() PatternFlowIpv6FlowLabel
	PayloadLength() PatternFlowIpv6PayloadLength
	NextHeader() PatternFlowIpv6NextHeader
	HopLimit() PatternFlowIpv6HopLimit
	Src() PatternFlowIpv6Src
	Dst() PatternFlowIpv6Dst
}

func (obj *flowIpv6) Version() PatternFlowIpv6Version {
	if obj.obj.Version == nil {
		obj.obj.Version = &snappipb.PatternFlowIpv6Version{}
	}
	return &patternFlowIpv6Version{obj: obj.obj.Version}

}

func (obj *flowIpv6) TrafficClass() PatternFlowIpv6TrafficClass {
	if obj.obj.TrafficClass == nil {
		obj.obj.TrafficClass = &snappipb.PatternFlowIpv6TrafficClass{}
	}
	return &patternFlowIpv6TrafficClass{obj: obj.obj.TrafficClass}

}

func (obj *flowIpv6) FlowLabel() PatternFlowIpv6FlowLabel {
	if obj.obj.FlowLabel == nil {
		obj.obj.FlowLabel = &snappipb.PatternFlowIpv6FlowLabel{}
	}
	return &patternFlowIpv6FlowLabel{obj: obj.obj.FlowLabel}

}

func (obj *flowIpv6) PayloadLength() PatternFlowIpv6PayloadLength {
	if obj.obj.PayloadLength == nil {
		obj.obj.PayloadLength = &snappipb.PatternFlowIpv6PayloadLength{}
	}
	return &patternFlowIpv6PayloadLength{obj: obj.obj.PayloadLength}

}

func (obj *flowIpv6) NextHeader() PatternFlowIpv6NextHeader {
	if obj.obj.NextHeader == nil {
		obj.obj.NextHeader = &snappipb.PatternFlowIpv6NextHeader{}
	}
	return &patternFlowIpv6NextHeader{obj: obj.obj.NextHeader}

}

func (obj *flowIpv6) HopLimit() PatternFlowIpv6HopLimit {
	if obj.obj.HopLimit == nil {
		obj.obj.HopLimit = &snappipb.PatternFlowIpv6HopLimit{}
	}
	return &patternFlowIpv6HopLimit{obj: obj.obj.HopLimit}

}

func (obj *flowIpv6) Src() PatternFlowIpv6Src {
	if obj.obj.Src == nil {
		obj.obj.Src = &snappipb.PatternFlowIpv6Src{}
	}
	return &patternFlowIpv6Src{obj: obj.obj.Src}

}

func (obj *flowIpv6) Dst() PatternFlowIpv6Dst {
	if obj.obj.Dst == nil {
		obj.obj.Dst = &snappipb.PatternFlowIpv6Dst{}
	}
	return &patternFlowIpv6Dst{obj: obj.obj.Dst}

}

type flowPfcPause struct {
	obj *snappipb.FlowPfcPause
}

func (obj *flowPfcPause) msg() *snappipb.FlowPfcPause {
	return obj.obj
}

func (obj *flowPfcPause) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowPfcPause) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowPfcPause interface {
	msg() *snappipb.FlowPfcPause
	Yaml() string
	Json() string
	Dst() PatternFlowPfcPauseDst
	Src() PatternFlowPfcPauseSrc
	EtherType() PatternFlowPfcPauseEtherType
	ControlOpCode() PatternFlowPfcPauseControlOpCode
	ClassEnableVector() PatternFlowPfcPauseClassEnableVector
	PauseClass0() PatternFlowPfcPausePauseClass0
	PauseClass1() PatternFlowPfcPausePauseClass1
	PauseClass2() PatternFlowPfcPausePauseClass2
	PauseClass3() PatternFlowPfcPausePauseClass3
	PauseClass4() PatternFlowPfcPausePauseClass4
	PauseClass5() PatternFlowPfcPausePauseClass5
	PauseClass6() PatternFlowPfcPausePauseClass6
	PauseClass7() PatternFlowPfcPausePauseClass7
}

func (obj *flowPfcPause) Dst() PatternFlowPfcPauseDst {
	if obj.obj.Dst == nil {
		obj.obj.Dst = &snappipb.PatternFlowPfcPauseDst{}
	}
	return &patternFlowPfcPauseDst{obj: obj.obj.Dst}

}

func (obj *flowPfcPause) Src() PatternFlowPfcPauseSrc {
	if obj.obj.Src == nil {
		obj.obj.Src = &snappipb.PatternFlowPfcPauseSrc{}
	}
	return &patternFlowPfcPauseSrc{obj: obj.obj.Src}

}

func (obj *flowPfcPause) EtherType() PatternFlowPfcPauseEtherType {
	if obj.obj.EtherType == nil {
		obj.obj.EtherType = &snappipb.PatternFlowPfcPauseEtherType{}
	}
	return &patternFlowPfcPauseEtherType{obj: obj.obj.EtherType}

}

func (obj *flowPfcPause) ControlOpCode() PatternFlowPfcPauseControlOpCode {
	if obj.obj.ControlOpCode == nil {
		obj.obj.ControlOpCode = &snappipb.PatternFlowPfcPauseControlOpCode{}
	}
	return &patternFlowPfcPauseControlOpCode{obj: obj.obj.ControlOpCode}

}

func (obj *flowPfcPause) ClassEnableVector() PatternFlowPfcPauseClassEnableVector {
	if obj.obj.ClassEnableVector == nil {
		obj.obj.ClassEnableVector = &snappipb.PatternFlowPfcPauseClassEnableVector{}
	}
	return &patternFlowPfcPauseClassEnableVector{obj: obj.obj.ClassEnableVector}

}

func (obj *flowPfcPause) PauseClass0() PatternFlowPfcPausePauseClass0 {
	if obj.obj.PauseClass0 == nil {
		obj.obj.PauseClass0 = &snappipb.PatternFlowPfcPausePauseClass0{}
	}
	return &patternFlowPfcPausePauseClass0{obj: obj.obj.PauseClass0}

}

func (obj *flowPfcPause) PauseClass1() PatternFlowPfcPausePauseClass1 {
	if obj.obj.PauseClass1 == nil {
		obj.obj.PauseClass1 = &snappipb.PatternFlowPfcPausePauseClass1{}
	}
	return &patternFlowPfcPausePauseClass1{obj: obj.obj.PauseClass1}

}

func (obj *flowPfcPause) PauseClass2() PatternFlowPfcPausePauseClass2 {
	if obj.obj.PauseClass2 == nil {
		obj.obj.PauseClass2 = &snappipb.PatternFlowPfcPausePauseClass2{}
	}
	return &patternFlowPfcPausePauseClass2{obj: obj.obj.PauseClass2}

}

func (obj *flowPfcPause) PauseClass3() PatternFlowPfcPausePauseClass3 {
	if obj.obj.PauseClass3 == nil {
		obj.obj.PauseClass3 = &snappipb.PatternFlowPfcPausePauseClass3{}
	}
	return &patternFlowPfcPausePauseClass3{obj: obj.obj.PauseClass3}

}

func (obj *flowPfcPause) PauseClass4() PatternFlowPfcPausePauseClass4 {
	if obj.obj.PauseClass4 == nil {
		obj.obj.PauseClass4 = &snappipb.PatternFlowPfcPausePauseClass4{}
	}
	return &patternFlowPfcPausePauseClass4{obj: obj.obj.PauseClass4}

}

func (obj *flowPfcPause) PauseClass5() PatternFlowPfcPausePauseClass5 {
	if obj.obj.PauseClass5 == nil {
		obj.obj.PauseClass5 = &snappipb.PatternFlowPfcPausePauseClass5{}
	}
	return &patternFlowPfcPausePauseClass5{obj: obj.obj.PauseClass5}

}

func (obj *flowPfcPause) PauseClass6() PatternFlowPfcPausePauseClass6 {
	if obj.obj.PauseClass6 == nil {
		obj.obj.PauseClass6 = &snappipb.PatternFlowPfcPausePauseClass6{}
	}
	return &patternFlowPfcPausePauseClass6{obj: obj.obj.PauseClass6}

}

func (obj *flowPfcPause) PauseClass7() PatternFlowPfcPausePauseClass7 {
	if obj.obj.PauseClass7 == nil {
		obj.obj.PauseClass7 = &snappipb.PatternFlowPfcPausePauseClass7{}
	}
	return &patternFlowPfcPausePauseClass7{obj: obj.obj.PauseClass7}

}

type flowEthernetPause struct {
	obj *snappipb.FlowEthernetPause
}

func (obj *flowEthernetPause) msg() *snappipb.FlowEthernetPause {
	return obj.obj
}

func (obj *flowEthernetPause) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowEthernetPause) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowEthernetPause interface {
	msg() *snappipb.FlowEthernetPause
	Yaml() string
	Json() string
	Dst() PatternFlowEthernetPauseDst
	Src() PatternFlowEthernetPauseSrc
	EtherType() PatternFlowEthernetPauseEtherType
	ControlOpCode() PatternFlowEthernetPauseControlOpCode
	Time() PatternFlowEthernetPauseTime
}

func (obj *flowEthernetPause) Dst() PatternFlowEthernetPauseDst {
	if obj.obj.Dst == nil {
		obj.obj.Dst = &snappipb.PatternFlowEthernetPauseDst{}
	}
	return &patternFlowEthernetPauseDst{obj: obj.obj.Dst}

}

func (obj *flowEthernetPause) Src() PatternFlowEthernetPauseSrc {
	if obj.obj.Src == nil {
		obj.obj.Src = &snappipb.PatternFlowEthernetPauseSrc{}
	}
	return &patternFlowEthernetPauseSrc{obj: obj.obj.Src}

}

func (obj *flowEthernetPause) EtherType() PatternFlowEthernetPauseEtherType {
	if obj.obj.EtherType == nil {
		obj.obj.EtherType = &snappipb.PatternFlowEthernetPauseEtherType{}
	}
	return &patternFlowEthernetPauseEtherType{obj: obj.obj.EtherType}

}

func (obj *flowEthernetPause) ControlOpCode() PatternFlowEthernetPauseControlOpCode {
	if obj.obj.ControlOpCode == nil {
		obj.obj.ControlOpCode = &snappipb.PatternFlowEthernetPauseControlOpCode{}
	}
	return &patternFlowEthernetPauseControlOpCode{obj: obj.obj.ControlOpCode}

}

func (obj *flowEthernetPause) Time() PatternFlowEthernetPauseTime {
	if obj.obj.Time == nil {
		obj.obj.Time = &snappipb.PatternFlowEthernetPauseTime{}
	}
	return &patternFlowEthernetPauseTime{obj: obj.obj.Time}

}

type flowTcp struct {
	obj *snappipb.FlowTcp
}

func (obj *flowTcp) msg() *snappipb.FlowTcp {
	return obj.obj
}

func (obj *flowTcp) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowTcp) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowTcp interface {
	msg() *snappipb.FlowTcp
	Yaml() string
	Json() string
	SrcPort() PatternFlowTcpSrcPort
	DstPort() PatternFlowTcpDstPort
	SeqNum() PatternFlowTcpSeqNum
	AckNum() PatternFlowTcpAckNum
	DataOffset() PatternFlowTcpDataOffset
	EcnNs() PatternFlowTcpEcnNs
	EcnCwr() PatternFlowTcpEcnCwr
	EcnEcho() PatternFlowTcpEcnEcho
	CtlUrg() PatternFlowTcpCtlUrg
	CtlAck() PatternFlowTcpCtlAck
	CtlPsh() PatternFlowTcpCtlPsh
	CtlRst() PatternFlowTcpCtlRst
	CtlSyn() PatternFlowTcpCtlSyn
	CtlFin() PatternFlowTcpCtlFin
	Window() PatternFlowTcpWindow
}

func (obj *flowTcp) SrcPort() PatternFlowTcpSrcPort {
	if obj.obj.SrcPort == nil {
		obj.obj.SrcPort = &snappipb.PatternFlowTcpSrcPort{}
	}
	return &patternFlowTcpSrcPort{obj: obj.obj.SrcPort}

}

func (obj *flowTcp) DstPort() PatternFlowTcpDstPort {
	if obj.obj.DstPort == nil {
		obj.obj.DstPort = &snappipb.PatternFlowTcpDstPort{}
	}
	return &patternFlowTcpDstPort{obj: obj.obj.DstPort}

}

func (obj *flowTcp) SeqNum() PatternFlowTcpSeqNum {
	if obj.obj.SeqNum == nil {
		obj.obj.SeqNum = &snappipb.PatternFlowTcpSeqNum{}
	}
	return &patternFlowTcpSeqNum{obj: obj.obj.SeqNum}

}

func (obj *flowTcp) AckNum() PatternFlowTcpAckNum {
	if obj.obj.AckNum == nil {
		obj.obj.AckNum = &snappipb.PatternFlowTcpAckNum{}
	}
	return &patternFlowTcpAckNum{obj: obj.obj.AckNum}

}

func (obj *flowTcp) DataOffset() PatternFlowTcpDataOffset {
	if obj.obj.DataOffset == nil {
		obj.obj.DataOffset = &snappipb.PatternFlowTcpDataOffset{}
	}
	return &patternFlowTcpDataOffset{obj: obj.obj.DataOffset}

}

func (obj *flowTcp) EcnNs() PatternFlowTcpEcnNs {
	if obj.obj.EcnNs == nil {
		obj.obj.EcnNs = &snappipb.PatternFlowTcpEcnNs{}
	}
	return &patternFlowTcpEcnNs{obj: obj.obj.EcnNs}

}

func (obj *flowTcp) EcnCwr() PatternFlowTcpEcnCwr {
	if obj.obj.EcnCwr == nil {
		obj.obj.EcnCwr = &snappipb.PatternFlowTcpEcnCwr{}
	}
	return &patternFlowTcpEcnCwr{obj: obj.obj.EcnCwr}

}

func (obj *flowTcp) EcnEcho() PatternFlowTcpEcnEcho {
	if obj.obj.EcnEcho == nil {
		obj.obj.EcnEcho = &snappipb.PatternFlowTcpEcnEcho{}
	}
	return &patternFlowTcpEcnEcho{obj: obj.obj.EcnEcho}

}

func (obj *flowTcp) CtlUrg() PatternFlowTcpCtlUrg {
	if obj.obj.CtlUrg == nil {
		obj.obj.CtlUrg = &snappipb.PatternFlowTcpCtlUrg{}
	}
	return &patternFlowTcpCtlUrg{obj: obj.obj.CtlUrg}

}

func (obj *flowTcp) CtlAck() PatternFlowTcpCtlAck {
	if obj.obj.CtlAck == nil {
		obj.obj.CtlAck = &snappipb.PatternFlowTcpCtlAck{}
	}
	return &patternFlowTcpCtlAck{obj: obj.obj.CtlAck}

}

func (obj *flowTcp) CtlPsh() PatternFlowTcpCtlPsh {
	if obj.obj.CtlPsh == nil {
		obj.obj.CtlPsh = &snappipb.PatternFlowTcpCtlPsh{}
	}
	return &patternFlowTcpCtlPsh{obj: obj.obj.CtlPsh}

}

func (obj *flowTcp) CtlRst() PatternFlowTcpCtlRst {
	if obj.obj.CtlRst == nil {
		obj.obj.CtlRst = &snappipb.PatternFlowTcpCtlRst{}
	}
	return &patternFlowTcpCtlRst{obj: obj.obj.CtlRst}

}

func (obj *flowTcp) CtlSyn() PatternFlowTcpCtlSyn {
	if obj.obj.CtlSyn == nil {
		obj.obj.CtlSyn = &snappipb.PatternFlowTcpCtlSyn{}
	}
	return &patternFlowTcpCtlSyn{obj: obj.obj.CtlSyn}

}

func (obj *flowTcp) CtlFin() PatternFlowTcpCtlFin {
	if obj.obj.CtlFin == nil {
		obj.obj.CtlFin = &snappipb.PatternFlowTcpCtlFin{}
	}
	return &patternFlowTcpCtlFin{obj: obj.obj.CtlFin}

}

func (obj *flowTcp) Window() PatternFlowTcpWindow {
	if obj.obj.Window == nil {
		obj.obj.Window = &snappipb.PatternFlowTcpWindow{}
	}
	return &patternFlowTcpWindow{obj: obj.obj.Window}

}

type flowUdp struct {
	obj *snappipb.FlowUdp
}

func (obj *flowUdp) msg() *snappipb.FlowUdp {
	return obj.obj
}

func (obj *flowUdp) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowUdp) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowUdp interface {
	msg() *snappipb.FlowUdp
	Yaml() string
	Json() string
	SrcPort() PatternFlowUdpSrcPort
	DstPort() PatternFlowUdpDstPort
	Length() PatternFlowUdpLength
	Checksum() PatternFlowUdpChecksum
}

func (obj *flowUdp) SrcPort() PatternFlowUdpSrcPort {
	if obj.obj.SrcPort == nil {
		obj.obj.SrcPort = &snappipb.PatternFlowUdpSrcPort{}
	}
	return &patternFlowUdpSrcPort{obj: obj.obj.SrcPort}

}

func (obj *flowUdp) DstPort() PatternFlowUdpDstPort {
	if obj.obj.DstPort == nil {
		obj.obj.DstPort = &snappipb.PatternFlowUdpDstPort{}
	}
	return &patternFlowUdpDstPort{obj: obj.obj.DstPort}

}

func (obj *flowUdp) Length() PatternFlowUdpLength {
	if obj.obj.Length == nil {
		obj.obj.Length = &snappipb.PatternFlowUdpLength{}
	}
	return &patternFlowUdpLength{obj: obj.obj.Length}

}

func (obj *flowUdp) Checksum() PatternFlowUdpChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = &snappipb.PatternFlowUdpChecksum{}
	}
	return &patternFlowUdpChecksum{obj: obj.obj.Checksum}

}

type flowGre struct {
	obj *snappipb.FlowGre
}

func (obj *flowGre) msg() *snappipb.FlowGre {
	return obj.obj
}

func (obj *flowGre) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowGre) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowGre interface {
	msg() *snappipb.FlowGre
	Yaml() string
	Json() string
	ChecksumPresent() PatternFlowGreChecksumPresent
	Reserved0() PatternFlowGreReserved0
	Version() PatternFlowGreVersion
	Protocol() PatternFlowGreProtocol
	Checksum() PatternFlowGreChecksum
	Reserved1() PatternFlowGreReserved1
}

func (obj *flowGre) ChecksumPresent() PatternFlowGreChecksumPresent {
	if obj.obj.ChecksumPresent == nil {
		obj.obj.ChecksumPresent = &snappipb.PatternFlowGreChecksumPresent{}
	}
	return &patternFlowGreChecksumPresent{obj: obj.obj.ChecksumPresent}

}

func (obj *flowGre) Reserved0() PatternFlowGreReserved0 {
	if obj.obj.Reserved0 == nil {
		obj.obj.Reserved0 = &snappipb.PatternFlowGreReserved0{}
	}
	return &patternFlowGreReserved0{obj: obj.obj.Reserved0}

}

func (obj *flowGre) Version() PatternFlowGreVersion {
	if obj.obj.Version == nil {
		obj.obj.Version = &snappipb.PatternFlowGreVersion{}
	}
	return &patternFlowGreVersion{obj: obj.obj.Version}

}

func (obj *flowGre) Protocol() PatternFlowGreProtocol {
	if obj.obj.Protocol == nil {
		obj.obj.Protocol = &snappipb.PatternFlowGreProtocol{}
	}
	return &patternFlowGreProtocol{obj: obj.obj.Protocol}

}

func (obj *flowGre) Checksum() PatternFlowGreChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = &snappipb.PatternFlowGreChecksum{}
	}
	return &patternFlowGreChecksum{obj: obj.obj.Checksum}

}

func (obj *flowGre) Reserved1() PatternFlowGreReserved1 {
	if obj.obj.Reserved1 == nil {
		obj.obj.Reserved1 = &snappipb.PatternFlowGreReserved1{}
	}
	return &patternFlowGreReserved1{obj: obj.obj.Reserved1}

}

type flowGtpv1 struct {
	obj *snappipb.FlowGtpv1
}

func (obj *flowGtpv1) msg() *snappipb.FlowGtpv1 {
	return obj.obj
}

func (obj *flowGtpv1) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowGtpv1) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowGtpv1 interface {
	msg() *snappipb.FlowGtpv1
	Yaml() string
	Json() string
	Version() PatternFlowGtpv1Version
	ProtocolType() PatternFlowGtpv1ProtocolType
	Reserved() PatternFlowGtpv1Reserved
	EFlag() PatternFlowGtpv1EFlag
	SFlag() PatternFlowGtpv1SFlag
	PnFlag() PatternFlowGtpv1PnFlag
	MessageType() PatternFlowGtpv1MessageType
	MessageLength() PatternFlowGtpv1MessageLength
	Teid() PatternFlowGtpv1Teid
	SquenceNumber() PatternFlowGtpv1SquenceNumber
	NPduNumber() PatternFlowGtpv1NPduNumber
	NextExtensionHeaderType() PatternFlowGtpv1NextExtensionHeaderType
	ExtensionHeaders() []FlowGtpExtension
	NewExtensionHeaders() FlowGtpExtension
}

func (obj *flowGtpv1) Version() PatternFlowGtpv1Version {
	if obj.obj.Version == nil {
		obj.obj.Version = &snappipb.PatternFlowGtpv1Version{}
	}
	return &patternFlowGtpv1Version{obj: obj.obj.Version}

}

func (obj *flowGtpv1) ProtocolType() PatternFlowGtpv1ProtocolType {
	if obj.obj.ProtocolType == nil {
		obj.obj.ProtocolType = &snappipb.PatternFlowGtpv1ProtocolType{}
	}
	return &patternFlowGtpv1ProtocolType{obj: obj.obj.ProtocolType}

}

func (obj *flowGtpv1) Reserved() PatternFlowGtpv1Reserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = &snappipb.PatternFlowGtpv1Reserved{}
	}
	return &patternFlowGtpv1Reserved{obj: obj.obj.Reserved}

}

func (obj *flowGtpv1) EFlag() PatternFlowGtpv1EFlag {
	if obj.obj.EFlag == nil {
		obj.obj.EFlag = &snappipb.PatternFlowGtpv1EFlag{}
	}
	return &patternFlowGtpv1EFlag{obj: obj.obj.EFlag}

}

func (obj *flowGtpv1) SFlag() PatternFlowGtpv1SFlag {
	if obj.obj.SFlag == nil {
		obj.obj.SFlag = &snappipb.PatternFlowGtpv1SFlag{}
	}
	return &patternFlowGtpv1SFlag{obj: obj.obj.SFlag}

}

func (obj *flowGtpv1) PnFlag() PatternFlowGtpv1PnFlag {
	if obj.obj.PnFlag == nil {
		obj.obj.PnFlag = &snappipb.PatternFlowGtpv1PnFlag{}
	}
	return &patternFlowGtpv1PnFlag{obj: obj.obj.PnFlag}

}

func (obj *flowGtpv1) MessageType() PatternFlowGtpv1MessageType {
	if obj.obj.MessageType == nil {
		obj.obj.MessageType = &snappipb.PatternFlowGtpv1MessageType{}
	}
	return &patternFlowGtpv1MessageType{obj: obj.obj.MessageType}

}

func (obj *flowGtpv1) MessageLength() PatternFlowGtpv1MessageLength {
	if obj.obj.MessageLength == nil {
		obj.obj.MessageLength = &snappipb.PatternFlowGtpv1MessageLength{}
	}
	return &patternFlowGtpv1MessageLength{obj: obj.obj.MessageLength}

}

func (obj *flowGtpv1) Teid() PatternFlowGtpv1Teid {
	if obj.obj.Teid == nil {
		obj.obj.Teid = &snappipb.PatternFlowGtpv1Teid{}
	}
	return &patternFlowGtpv1Teid{obj: obj.obj.Teid}

}

func (obj *flowGtpv1) SquenceNumber() PatternFlowGtpv1SquenceNumber {
	if obj.obj.SquenceNumber == nil {
		obj.obj.SquenceNumber = &snappipb.PatternFlowGtpv1SquenceNumber{}
	}
	return &patternFlowGtpv1SquenceNumber{obj: obj.obj.SquenceNumber}

}

func (obj *flowGtpv1) NPduNumber() PatternFlowGtpv1NPduNumber {
	if obj.obj.NPduNumber == nil {
		obj.obj.NPduNumber = &snappipb.PatternFlowGtpv1NPduNumber{}
	}
	return &patternFlowGtpv1NPduNumber{obj: obj.obj.NPduNumber}

}

func (obj *flowGtpv1) NextExtensionHeaderType() PatternFlowGtpv1NextExtensionHeaderType {
	if obj.obj.NextExtensionHeaderType == nil {
		obj.obj.NextExtensionHeaderType = &snappipb.PatternFlowGtpv1NextExtensionHeaderType{}
	}
	return &patternFlowGtpv1NextExtensionHeaderType{obj: obj.obj.NextExtensionHeaderType}

}

func (obj *flowGtpv1) ExtensionHeaders() []FlowGtpExtension {
	if obj.obj.ExtensionHeaders == nil {
		obj.obj.ExtensionHeaders = make([]*snappipb.FlowGtpExtension, 0)
	}
	values := make([]FlowGtpExtension, 0)
	for _, item := range obj.obj.ExtensionHeaders {
		values = append(values, &flowGtpExtension{obj: item})
	}
	return values

}

func (obj *flowGtpv1) NewExtensionHeaders() FlowGtpExtension {
	if obj.obj.ExtensionHeaders == nil {
		obj.obj.ExtensionHeaders = make([]*snappipb.FlowGtpExtension, 0)
	}
	slice := append(obj.obj.ExtensionHeaders, &snappipb.FlowGtpExtension{})
	obj.obj.ExtensionHeaders = slice
	return &flowGtpExtension{obj: slice[len(slice)-1]}
}

type flowGtpv2 struct {
	obj *snappipb.FlowGtpv2
}

func (obj *flowGtpv2) msg() *snappipb.FlowGtpv2 {
	return obj.obj
}

func (obj *flowGtpv2) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowGtpv2) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowGtpv2 interface {
	msg() *snappipb.FlowGtpv2
	Yaml() string
	Json() string
	Version() PatternFlowGtpv2Version
	PiggybackingFlag() PatternFlowGtpv2PiggybackingFlag
	TeidFlag() PatternFlowGtpv2TeidFlag
	Spare1() PatternFlowGtpv2Spare1
	MessageType() PatternFlowGtpv2MessageType
	MessageLength() PatternFlowGtpv2MessageLength
	Teid() PatternFlowGtpv2Teid
	SequenceNumber() PatternFlowGtpv2SequenceNumber
	Spare2() PatternFlowGtpv2Spare2
}

func (obj *flowGtpv2) Version() PatternFlowGtpv2Version {
	if obj.obj.Version == nil {
		obj.obj.Version = &snappipb.PatternFlowGtpv2Version{}
	}
	return &patternFlowGtpv2Version{obj: obj.obj.Version}

}

func (obj *flowGtpv2) PiggybackingFlag() PatternFlowGtpv2PiggybackingFlag {
	if obj.obj.PiggybackingFlag == nil {
		obj.obj.PiggybackingFlag = &snappipb.PatternFlowGtpv2PiggybackingFlag{}
	}
	return &patternFlowGtpv2PiggybackingFlag{obj: obj.obj.PiggybackingFlag}

}

func (obj *flowGtpv2) TeidFlag() PatternFlowGtpv2TeidFlag {
	if obj.obj.TeidFlag == nil {
		obj.obj.TeidFlag = &snappipb.PatternFlowGtpv2TeidFlag{}
	}
	return &patternFlowGtpv2TeidFlag{obj: obj.obj.TeidFlag}

}

func (obj *flowGtpv2) Spare1() PatternFlowGtpv2Spare1 {
	if obj.obj.Spare1 == nil {
		obj.obj.Spare1 = &snappipb.PatternFlowGtpv2Spare1{}
	}
	return &patternFlowGtpv2Spare1{obj: obj.obj.Spare1}

}

func (obj *flowGtpv2) MessageType() PatternFlowGtpv2MessageType {
	if obj.obj.MessageType == nil {
		obj.obj.MessageType = &snappipb.PatternFlowGtpv2MessageType{}
	}
	return &patternFlowGtpv2MessageType{obj: obj.obj.MessageType}

}

func (obj *flowGtpv2) MessageLength() PatternFlowGtpv2MessageLength {
	if obj.obj.MessageLength == nil {
		obj.obj.MessageLength = &snappipb.PatternFlowGtpv2MessageLength{}
	}
	return &patternFlowGtpv2MessageLength{obj: obj.obj.MessageLength}

}

func (obj *flowGtpv2) Teid() PatternFlowGtpv2Teid {
	if obj.obj.Teid == nil {
		obj.obj.Teid = &snappipb.PatternFlowGtpv2Teid{}
	}
	return &patternFlowGtpv2Teid{obj: obj.obj.Teid}

}

func (obj *flowGtpv2) SequenceNumber() PatternFlowGtpv2SequenceNumber {
	if obj.obj.SequenceNumber == nil {
		obj.obj.SequenceNumber = &snappipb.PatternFlowGtpv2SequenceNumber{}
	}
	return &patternFlowGtpv2SequenceNumber{obj: obj.obj.SequenceNumber}

}

func (obj *flowGtpv2) Spare2() PatternFlowGtpv2Spare2 {
	if obj.obj.Spare2 == nil {
		obj.obj.Spare2 = &snappipb.PatternFlowGtpv2Spare2{}
	}
	return &patternFlowGtpv2Spare2{obj: obj.obj.Spare2}

}

type flowArp struct {
	obj *snappipb.FlowArp
}

func (obj *flowArp) msg() *snappipb.FlowArp {
	return obj.obj
}

func (obj *flowArp) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowArp) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowArp interface {
	msg() *snappipb.FlowArp
	Yaml() string
	Json() string
	HardwareType() PatternFlowArpHardwareType
	ProtocolType() PatternFlowArpProtocolType
	HardwareLength() PatternFlowArpHardwareLength
	ProtocolLength() PatternFlowArpProtocolLength
	Operation() PatternFlowArpOperation
	SenderHardwareAddr() PatternFlowArpSenderHardwareAddr
	SenderProtocolAddr() PatternFlowArpSenderProtocolAddr
	TargetHardwareAddr() PatternFlowArpTargetHardwareAddr
	TargetProtocolAddr() PatternFlowArpTargetProtocolAddr
}

func (obj *flowArp) HardwareType() PatternFlowArpHardwareType {
	if obj.obj.HardwareType == nil {
		obj.obj.HardwareType = &snappipb.PatternFlowArpHardwareType{}
	}
	return &patternFlowArpHardwareType{obj: obj.obj.HardwareType}

}

func (obj *flowArp) ProtocolType() PatternFlowArpProtocolType {
	if obj.obj.ProtocolType == nil {
		obj.obj.ProtocolType = &snappipb.PatternFlowArpProtocolType{}
	}
	return &patternFlowArpProtocolType{obj: obj.obj.ProtocolType}

}

func (obj *flowArp) HardwareLength() PatternFlowArpHardwareLength {
	if obj.obj.HardwareLength == nil {
		obj.obj.HardwareLength = &snappipb.PatternFlowArpHardwareLength{}
	}
	return &patternFlowArpHardwareLength{obj: obj.obj.HardwareLength}

}

func (obj *flowArp) ProtocolLength() PatternFlowArpProtocolLength {
	if obj.obj.ProtocolLength == nil {
		obj.obj.ProtocolLength = &snappipb.PatternFlowArpProtocolLength{}
	}
	return &patternFlowArpProtocolLength{obj: obj.obj.ProtocolLength}

}

func (obj *flowArp) Operation() PatternFlowArpOperation {
	if obj.obj.Operation == nil {
		obj.obj.Operation = &snappipb.PatternFlowArpOperation{}
	}
	return &patternFlowArpOperation{obj: obj.obj.Operation}

}

func (obj *flowArp) SenderHardwareAddr() PatternFlowArpSenderHardwareAddr {
	if obj.obj.SenderHardwareAddr == nil {
		obj.obj.SenderHardwareAddr = &snappipb.PatternFlowArpSenderHardwareAddr{}
	}
	return &patternFlowArpSenderHardwareAddr{obj: obj.obj.SenderHardwareAddr}

}

func (obj *flowArp) SenderProtocolAddr() PatternFlowArpSenderProtocolAddr {
	if obj.obj.SenderProtocolAddr == nil {
		obj.obj.SenderProtocolAddr = &snappipb.PatternFlowArpSenderProtocolAddr{}
	}
	return &patternFlowArpSenderProtocolAddr{obj: obj.obj.SenderProtocolAddr}

}

func (obj *flowArp) TargetHardwareAddr() PatternFlowArpTargetHardwareAddr {
	if obj.obj.TargetHardwareAddr == nil {
		obj.obj.TargetHardwareAddr = &snappipb.PatternFlowArpTargetHardwareAddr{}
	}
	return &patternFlowArpTargetHardwareAddr{obj: obj.obj.TargetHardwareAddr}

}

func (obj *flowArp) TargetProtocolAddr() PatternFlowArpTargetProtocolAddr {
	if obj.obj.TargetProtocolAddr == nil {
		obj.obj.TargetProtocolAddr = &snappipb.PatternFlowArpTargetProtocolAddr{}
	}
	return &patternFlowArpTargetProtocolAddr{obj: obj.obj.TargetProtocolAddr}

}

type flowIcmp struct {
	obj *snappipb.FlowIcmp
}

func (obj *flowIcmp) msg() *snappipb.FlowIcmp {
	return obj.obj
}

func (obj *flowIcmp) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowIcmp) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowIcmp interface {
	msg() *snappipb.FlowIcmp
	Yaml() string
	Json() string
	Echo() FlowIcmpEcho
}

func (obj *flowIcmp) Echo() FlowIcmpEcho {
	if obj.obj.Echo == nil {
		obj.obj.Echo = &snappipb.FlowIcmpEcho{}
	}
	return &flowIcmpEcho{obj: obj.obj.Echo}

}

type flowIcmpv6 struct {
	obj *snappipb.FlowIcmpv6
}

func (obj *flowIcmpv6) msg() *snappipb.FlowIcmpv6 {
	return obj.obj
}

func (obj *flowIcmpv6) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowIcmpv6) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowIcmpv6 interface {
	msg() *snappipb.FlowIcmpv6
	Yaml() string
	Json() string
	Echo() FlowIcmpv6Echo
}

func (obj *flowIcmpv6) Echo() FlowIcmpv6Echo {
	if obj.obj.Echo == nil {
		obj.obj.Echo = &snappipb.FlowIcmpv6Echo{}
	}
	return &flowIcmpv6Echo{obj: obj.obj.Echo}

}

type flowPpp struct {
	obj *snappipb.FlowPpp
}

func (obj *flowPpp) msg() *snappipb.FlowPpp {
	return obj.obj
}

func (obj *flowPpp) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowPpp) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowPpp interface {
	msg() *snappipb.FlowPpp
	Yaml() string
	Json() string
	Address() PatternFlowPppAddress
	Control() PatternFlowPppControl
	ProtocolType() PatternFlowPppProtocolType
}

func (obj *flowPpp) Address() PatternFlowPppAddress {
	if obj.obj.Address == nil {
		obj.obj.Address = &snappipb.PatternFlowPppAddress{}
	}
	return &patternFlowPppAddress{obj: obj.obj.Address}

}

func (obj *flowPpp) Control() PatternFlowPppControl {
	if obj.obj.Control == nil {
		obj.obj.Control = &snappipb.PatternFlowPppControl{}
	}
	return &patternFlowPppControl{obj: obj.obj.Control}

}

func (obj *flowPpp) ProtocolType() PatternFlowPppProtocolType {
	if obj.obj.ProtocolType == nil {
		obj.obj.ProtocolType = &snappipb.PatternFlowPppProtocolType{}
	}
	return &patternFlowPppProtocolType{obj: obj.obj.ProtocolType}

}

type flowIgmpv1 struct {
	obj *snappipb.FlowIgmpv1
}

func (obj *flowIgmpv1) msg() *snappipb.FlowIgmpv1 {
	return obj.obj
}

func (obj *flowIgmpv1) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowIgmpv1) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowIgmpv1 interface {
	msg() *snappipb.FlowIgmpv1
	Yaml() string
	Json() string
	Version() PatternFlowIgmpv1Version
	Type() PatternFlowIgmpv1Type
	Unused() PatternFlowIgmpv1Unused
	Checksum() PatternFlowIgmpv1Checksum
	GroupAddress() PatternFlowIgmpv1GroupAddress
}

func (obj *flowIgmpv1) Version() PatternFlowIgmpv1Version {
	if obj.obj.Version == nil {
		obj.obj.Version = &snappipb.PatternFlowIgmpv1Version{}
	}
	return &patternFlowIgmpv1Version{obj: obj.obj.Version}

}

func (obj *flowIgmpv1) Type() PatternFlowIgmpv1Type {
	if obj.obj.Type == nil {
		obj.obj.Type = &snappipb.PatternFlowIgmpv1Type{}
	}
	return &patternFlowIgmpv1Type{obj: obj.obj.Type}

}

func (obj *flowIgmpv1) Unused() PatternFlowIgmpv1Unused {
	if obj.obj.Unused == nil {
		obj.obj.Unused = &snappipb.PatternFlowIgmpv1Unused{}
	}
	return &patternFlowIgmpv1Unused{obj: obj.obj.Unused}

}

func (obj *flowIgmpv1) Checksum() PatternFlowIgmpv1Checksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = &snappipb.PatternFlowIgmpv1Checksum{}
	}
	return &patternFlowIgmpv1Checksum{obj: obj.obj.Checksum}

}

func (obj *flowIgmpv1) GroupAddress() PatternFlowIgmpv1GroupAddress {
	if obj.obj.GroupAddress == nil {
		obj.obj.GroupAddress = &snappipb.PatternFlowIgmpv1GroupAddress{}
	}
	return &patternFlowIgmpv1GroupAddress{obj: obj.obj.GroupAddress}

}

type flowSizeIncrement struct {
	obj *snappipb.FlowSizeIncrement
}

func (obj *flowSizeIncrement) msg() *snappipb.FlowSizeIncrement {
	return obj.obj
}

func (obj *flowSizeIncrement) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowSizeIncrement) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowSizeIncrement interface {
	msg() *snappipb.FlowSizeIncrement
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) FlowSizeIncrement
	End() int32
	SetEnd(value int32) FlowSizeIncrement
	Step() int32
	SetStep(value int32) FlowSizeIncrement
}

func (obj *flowSizeIncrement) Start() int32 {
	return *obj.obj.Start
}

func (obj *flowSizeIncrement) SetStart(value int32) FlowSizeIncrement {
	obj.obj.Start = &value
	return obj
}

func (obj *flowSizeIncrement) End() int32 {
	return *obj.obj.End
}

func (obj *flowSizeIncrement) SetEnd(value int32) FlowSizeIncrement {
	obj.obj.End = &value
	return obj
}

func (obj *flowSizeIncrement) Step() int32 {
	return *obj.obj.Step
}

func (obj *flowSizeIncrement) SetStep(value int32) FlowSizeIncrement {
	obj.obj.Step = &value
	return obj
}

type flowSizeRandom struct {
	obj *snappipb.FlowSizeRandom
}

func (obj *flowSizeRandom) msg() *snappipb.FlowSizeRandom {
	return obj.obj
}

func (obj *flowSizeRandom) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowSizeRandom) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowSizeRandom interface {
	msg() *snappipb.FlowSizeRandom
	Yaml() string
	Json() string
	Min() int32
	SetMin(value int32) FlowSizeRandom
	Max() int32
	SetMax(value int32) FlowSizeRandom
}

func (obj *flowSizeRandom) Min() int32 {
	return *obj.obj.Min
}

func (obj *flowSizeRandom) SetMin(value int32) FlowSizeRandom {
	obj.obj.Min = &value
	return obj
}

func (obj *flowSizeRandom) Max() int32 {
	return *obj.obj.Max
}

func (obj *flowSizeRandom) SetMax(value int32) FlowSizeRandom {
	obj.obj.Max = &value
	return obj
}

type flowFixedPackets struct {
	obj *snappipb.FlowFixedPackets
}

func (obj *flowFixedPackets) msg() *snappipb.FlowFixedPackets {
	return obj.obj
}

func (obj *flowFixedPackets) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowFixedPackets) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowFixedPackets interface {
	msg() *snappipb.FlowFixedPackets
	Yaml() string
	Json() string
	Packets() int32
	SetPackets(value int32) FlowFixedPackets
	Gap() int32
	SetGap(value int32) FlowFixedPackets
	Delay() FlowDelay
}

func (obj *flowFixedPackets) Packets() int32 {
	return *obj.obj.Packets
}

func (obj *flowFixedPackets) SetPackets(value int32) FlowFixedPackets {
	obj.obj.Packets = &value
	return obj
}

func (obj *flowFixedPackets) Gap() int32 {
	return *obj.obj.Gap
}

func (obj *flowFixedPackets) SetGap(value int32) FlowFixedPackets {
	obj.obj.Gap = &value
	return obj
}

func (obj *flowFixedPackets) Delay() FlowDelay {
	if obj.obj.Delay == nil {
		obj.obj.Delay = &snappipb.FlowDelay{}
	}
	return &flowDelay{obj: obj.obj.Delay}

}

type flowFixedSeconds struct {
	obj *snappipb.FlowFixedSeconds
}

func (obj *flowFixedSeconds) msg() *snappipb.FlowFixedSeconds {
	return obj.obj
}

func (obj *flowFixedSeconds) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowFixedSeconds) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowFixedSeconds interface {
	msg() *snappipb.FlowFixedSeconds
	Yaml() string
	Json() string
	Seconds() float32
	SetSeconds(value float32) FlowFixedSeconds
	Gap() int32
	SetGap(value int32) FlowFixedSeconds
	Delay() FlowDelay
}

func (obj *flowFixedSeconds) Seconds() float32 {
	return *obj.obj.Seconds
}

func (obj *flowFixedSeconds) SetSeconds(value float32) FlowFixedSeconds {
	obj.obj.Seconds = &value
	return obj
}

func (obj *flowFixedSeconds) Gap() int32 {
	return *obj.obj.Gap
}

func (obj *flowFixedSeconds) SetGap(value int32) FlowFixedSeconds {
	obj.obj.Gap = &value
	return obj
}

func (obj *flowFixedSeconds) Delay() FlowDelay {
	if obj.obj.Delay == nil {
		obj.obj.Delay = &snappipb.FlowDelay{}
	}
	return &flowDelay{obj: obj.obj.Delay}

}

type flowBurst struct {
	obj *snappipb.FlowBurst
}

func (obj *flowBurst) msg() *snappipb.FlowBurst {
	return obj.obj
}

func (obj *flowBurst) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowBurst) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowBurst interface {
	msg() *snappipb.FlowBurst
	Yaml() string
	Json() string
	Bursts() int32
	SetBursts(value int32) FlowBurst
	Packets() int32
	SetPackets(value int32) FlowBurst
	Gap() int32
	SetGap(value int32) FlowBurst
	InterBurstGap() FlowDurationInterBurstGap
}

func (obj *flowBurst) Bursts() int32 {
	return *obj.obj.Bursts
}

func (obj *flowBurst) SetBursts(value int32) FlowBurst {
	obj.obj.Bursts = &value
	return obj
}

func (obj *flowBurst) Packets() int32 {
	return *obj.obj.Packets
}

func (obj *flowBurst) SetPackets(value int32) FlowBurst {
	obj.obj.Packets = &value
	return obj
}

func (obj *flowBurst) Gap() int32 {
	return *obj.obj.Gap
}

func (obj *flowBurst) SetGap(value int32) FlowBurst {
	obj.obj.Gap = &value
	return obj
}

func (obj *flowBurst) InterBurstGap() FlowDurationInterBurstGap {
	if obj.obj.InterBurstGap == nil {
		obj.obj.InterBurstGap = &snappipb.FlowDurationInterBurstGap{}
	}
	return &flowDurationInterBurstGap{obj: obj.obj.InterBurstGap}

}

type flowContinuous struct {
	obj *snappipb.FlowContinuous
}

func (obj *flowContinuous) msg() *snappipb.FlowContinuous {
	return obj.obj
}

func (obj *flowContinuous) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowContinuous) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowContinuous interface {
	msg() *snappipb.FlowContinuous
	Yaml() string
	Json() string
	Gap() int32
	SetGap(value int32) FlowContinuous
	Delay() FlowDelay
}

func (obj *flowContinuous) Gap() int32 {
	return *obj.obj.Gap
}

func (obj *flowContinuous) SetGap(value int32) FlowContinuous {
	obj.obj.Gap = &value
	return obj
}

func (obj *flowContinuous) Delay() FlowDelay {
	if obj.obj.Delay == nil {
		obj.obj.Delay = &snappipb.FlowDelay{}
	}
	return &flowDelay{obj: obj.obj.Delay}

}

type flowLatencyMetrics struct {
	obj *snappipb.FlowLatencyMetrics
}

func (obj *flowLatencyMetrics) msg() *snappipb.FlowLatencyMetrics {
	return obj.obj
}

func (obj *flowLatencyMetrics) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowLatencyMetrics) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowLatencyMetrics interface {
	msg() *snappipb.FlowLatencyMetrics
	Yaml() string
	Json() string
	Enable() bool
	SetEnable(value bool) FlowLatencyMetrics
}

func (obj *flowLatencyMetrics) Enable() bool {
	return *obj.obj.Enable
}

func (obj *flowLatencyMetrics) SetEnable(value bool) FlowLatencyMetrics {
	obj.obj.Enable = &value
	return obj
}

type lagLacp struct {
	obj *snappipb.LagLacp
}

func (obj *lagLacp) msg() *snappipb.LagLacp {
	return obj.obj
}

func (obj *lagLacp) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *lagLacp) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type LagLacp interface {
	msg() *snappipb.LagLacp
	Yaml() string
	Json() string
	ActorKey() int32
	SetActorKey(value int32) LagLacp
	ActorPortNumber() int32
	SetActorPortNumber(value int32) LagLacp
	ActorPortPriority() int32
	SetActorPortPriority(value int32) LagLacp
	ActorSystemId() string
	SetActorSystemId(value string) LagLacp
	ActorSystemPriority() int32
	SetActorSystemPriority(value int32) LagLacp
	LacpduPeriodicTimeInterval() int32
	SetLacpduPeriodicTimeInterval(value int32) LagLacp
	LacpduTimeout() int32
	SetLacpduTimeout(value int32) LagLacp
}

func (obj *lagLacp) ActorKey() int32 {
	return *obj.obj.ActorKey
}

func (obj *lagLacp) SetActorKey(value int32) LagLacp {
	obj.obj.ActorKey = &value
	return obj
}

func (obj *lagLacp) ActorPortNumber() int32 {
	return *obj.obj.ActorPortNumber
}

func (obj *lagLacp) SetActorPortNumber(value int32) LagLacp {
	obj.obj.ActorPortNumber = &value
	return obj
}

func (obj *lagLacp) ActorPortPriority() int32 {
	return *obj.obj.ActorPortPriority
}

func (obj *lagLacp) SetActorPortPriority(value int32) LagLacp {
	obj.obj.ActorPortPriority = &value
	return obj
}

func (obj *lagLacp) ActorSystemId() string {
	return *obj.obj.ActorSystemId
}

func (obj *lagLacp) SetActorSystemId(value string) LagLacp {
	obj.obj.ActorSystemId = &value
	return obj
}

func (obj *lagLacp) ActorSystemPriority() int32 {
	return *obj.obj.ActorSystemPriority
}

func (obj *lagLacp) SetActorSystemPriority(value int32) LagLacp {
	obj.obj.ActorSystemPriority = &value
	return obj
}

func (obj *lagLacp) LacpduPeriodicTimeInterval() int32 {
	return *obj.obj.LacpduPeriodicTimeInterval
}

func (obj *lagLacp) SetLacpduPeriodicTimeInterval(value int32) LagLacp {
	obj.obj.LacpduPeriodicTimeInterval = &value
	return obj
}

func (obj *lagLacp) LacpduTimeout() int32 {
	return *obj.obj.LacpduTimeout
}

func (obj *lagLacp) SetLacpduTimeout(value int32) LagLacp {
	obj.obj.LacpduTimeout = &value
	return obj
}

type lagStatic struct {
	obj *snappipb.LagStatic
}

func (obj *lagStatic) msg() *snappipb.LagStatic {
	return obj.obj
}

func (obj *lagStatic) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *lagStatic) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type LagStatic interface {
	msg() *snappipb.LagStatic
	Yaml() string
	Json() string
	LagId() int32
	SetLagId(value int32) LagStatic
}

func (obj *lagStatic) LagId() int32 {
	return *obj.obj.LagId
}

func (obj *lagStatic) SetLagId(value int32) LagStatic {
	obj.obj.LagId = &value
	return obj
}

type captureField struct {
	obj *snappipb.CaptureField
}

func (obj *captureField) msg() *snappipb.CaptureField {
	return obj.obj
}

func (obj *captureField) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *captureField) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type CaptureField interface {
	msg() *snappipb.CaptureField
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) CaptureField
	Mask() string
	SetMask(value string) CaptureField
	Negate() bool
	SetNegate(value bool) CaptureField
}

func (obj *captureField) Value() string {
	return *obj.obj.Value
}

func (obj *captureField) SetValue(value string) CaptureField {
	obj.obj.Value = &value
	return obj
}

func (obj *captureField) Mask() string {
	return *obj.obj.Mask
}

func (obj *captureField) SetMask(value string) CaptureField {
	obj.obj.Mask = &value
	return obj
}

func (obj *captureField) Negate() bool {
	return *obj.obj.Negate
}

func (obj *captureField) SetNegate(value bool) CaptureField {
	obj.obj.Negate = &value
	return obj
}

type deviceBgpv4 struct {
	obj *snappipb.DeviceBgpv4
}

func (obj *deviceBgpv4) msg() *snappipb.DeviceBgpv4 {
	return obj.obj
}

func (obj *deviceBgpv4) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpv4) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpv4 interface {
	msg() *snappipb.DeviceBgpv4
	Yaml() string
	Json() string
	LocalAddress() string
	SetLocalAddress(value string) DeviceBgpv4
	DutAddress() string
	SetDutAddress(value string) DeviceBgpv4
	RouterId() string
	SetRouterId(value string) DeviceBgpv4
	AsNumber() int32
	SetAsNumber(value int32) DeviceBgpv4
	Advanced() DeviceBgpAdvanced
	Capability() DeviceBgpCapability
	SrTePolicies() []DeviceBgpSrTePolicy
	NewSrTePolicies() DeviceBgpSrTePolicy
	Bgpv4Routes() []DeviceBgpv4Route
	NewBgpv4Routes() DeviceBgpv4Route
	Bgpv6Routes() []DeviceBgpv6Route
	NewBgpv6Routes() DeviceBgpv6Route
	Name() string
	SetName(value string) DeviceBgpv4
	Active() bool
	SetActive(value bool) DeviceBgpv4
}

func (obj *deviceBgpv4) LocalAddress() string {
	return obj.obj.LocalAddress
}

func (obj *deviceBgpv4) SetLocalAddress(value string) DeviceBgpv4 {
	obj.obj.LocalAddress = value
	return obj
}

func (obj *deviceBgpv4) DutAddress() string {
	return obj.obj.DutAddress
}

func (obj *deviceBgpv4) SetDutAddress(value string) DeviceBgpv4 {
	obj.obj.DutAddress = value
	return obj
}

func (obj *deviceBgpv4) RouterId() string {
	return *obj.obj.RouterId
}

func (obj *deviceBgpv4) SetRouterId(value string) DeviceBgpv4 {
	obj.obj.RouterId = &value
	return obj
}

func (obj *deviceBgpv4) AsNumber() int32 {
	return obj.obj.AsNumber
}

func (obj *deviceBgpv4) SetAsNumber(value int32) DeviceBgpv4 {
	obj.obj.AsNumber = value
	return obj
}

func (obj *deviceBgpv4) Advanced() DeviceBgpAdvanced {
	if obj.obj.Advanced == nil {
		obj.obj.Advanced = &snappipb.DeviceBgpAdvanced{}
	}
	return &deviceBgpAdvanced{obj: obj.obj.Advanced}

}

func (obj *deviceBgpv4) Capability() DeviceBgpCapability {
	if obj.obj.Capability == nil {
		obj.obj.Capability = &snappipb.DeviceBgpCapability{}
	}
	return &deviceBgpCapability{obj: obj.obj.Capability}

}

func (obj *deviceBgpv4) SrTePolicies() []DeviceBgpSrTePolicy {
	if obj.obj.SrTePolicies == nil {
		obj.obj.SrTePolicies = make([]*snappipb.DeviceBgpSrTePolicy, 0)
	}
	values := make([]DeviceBgpSrTePolicy, 0)
	for _, item := range obj.obj.SrTePolicies {
		values = append(values, &deviceBgpSrTePolicy{obj: item})
	}
	return values

}

func (obj *deviceBgpv4) NewSrTePolicies() DeviceBgpSrTePolicy {
	if obj.obj.SrTePolicies == nil {
		obj.obj.SrTePolicies = make([]*snappipb.DeviceBgpSrTePolicy, 0)
	}
	slice := append(obj.obj.SrTePolicies, &snappipb.DeviceBgpSrTePolicy{})
	obj.obj.SrTePolicies = slice
	return &deviceBgpSrTePolicy{obj: slice[len(slice)-1]}
}

func (obj *deviceBgpv4) Bgpv4Routes() []DeviceBgpv4Route {
	if obj.obj.Bgpv4Routes == nil {
		obj.obj.Bgpv4Routes = make([]*snappipb.DeviceBgpv4Route, 0)
	}
	values := make([]DeviceBgpv4Route, 0)
	for _, item := range obj.obj.Bgpv4Routes {
		values = append(values, &deviceBgpv4Route{obj: item})
	}
	return values

}

func (obj *deviceBgpv4) NewBgpv4Routes() DeviceBgpv4Route {
	if obj.obj.Bgpv4Routes == nil {
		obj.obj.Bgpv4Routes = make([]*snappipb.DeviceBgpv4Route, 0)
	}
	slice := append(obj.obj.Bgpv4Routes, &snappipb.DeviceBgpv4Route{})
	obj.obj.Bgpv4Routes = slice
	return &deviceBgpv4Route{obj: slice[len(slice)-1]}
}

func (obj *deviceBgpv4) Bgpv6Routes() []DeviceBgpv6Route {
	if obj.obj.Bgpv6Routes == nil {
		obj.obj.Bgpv6Routes = make([]*snappipb.DeviceBgpv6Route, 0)
	}
	values := make([]DeviceBgpv6Route, 0)
	for _, item := range obj.obj.Bgpv6Routes {
		values = append(values, &deviceBgpv6Route{obj: item})
	}
	return values

}

func (obj *deviceBgpv4) NewBgpv6Routes() DeviceBgpv6Route {
	if obj.obj.Bgpv6Routes == nil {
		obj.obj.Bgpv6Routes = make([]*snappipb.DeviceBgpv6Route, 0)
	}
	slice := append(obj.obj.Bgpv6Routes, &snappipb.DeviceBgpv6Route{})
	obj.obj.Bgpv6Routes = slice
	return &deviceBgpv6Route{obj: slice[len(slice)-1]}
}

func (obj *deviceBgpv4) Name() string {
	return obj.obj.Name
}

func (obj *deviceBgpv4) SetName(value string) DeviceBgpv4 {
	obj.obj.Name = value
	return obj
}

func (obj *deviceBgpv4) Active() bool {
	return *obj.obj.Active
}

func (obj *deviceBgpv4) SetActive(value bool) DeviceBgpv4 {
	obj.obj.Active = &value
	return obj
}

type deviceBgpv6 struct {
	obj *snappipb.DeviceBgpv6
}

func (obj *deviceBgpv6) msg() *snappipb.DeviceBgpv6 {
	return obj.obj
}

func (obj *deviceBgpv6) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpv6) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpv6 interface {
	msg() *snappipb.DeviceBgpv6
	Yaml() string
	Json() string
	LocalAddress() string
	SetLocalAddress(value string) DeviceBgpv6
	DutAddress() string
	SetDutAddress(value string) DeviceBgpv6
	SegmentRouting() DeviceBgpv6SegmentRouting
	RouterId() string
	SetRouterId(value string) DeviceBgpv6
	AsNumber() int32
	SetAsNumber(value int32) DeviceBgpv6
	Advanced() DeviceBgpAdvanced
	Capability() DeviceBgpCapability
	SrTePolicies() []DeviceBgpSrTePolicy
	NewSrTePolicies() DeviceBgpSrTePolicy
	Bgpv4Routes() []DeviceBgpv4Route
	NewBgpv4Routes() DeviceBgpv4Route
	Bgpv6Routes() []DeviceBgpv6Route
	NewBgpv6Routes() DeviceBgpv6Route
	Name() string
	SetName(value string) DeviceBgpv6
	Active() bool
	SetActive(value bool) DeviceBgpv6
}

func (obj *deviceBgpv6) LocalAddress() string {
	return obj.obj.LocalAddress
}

func (obj *deviceBgpv6) SetLocalAddress(value string) DeviceBgpv6 {
	obj.obj.LocalAddress = value
	return obj
}

func (obj *deviceBgpv6) DutAddress() string {
	return obj.obj.DutAddress
}

func (obj *deviceBgpv6) SetDutAddress(value string) DeviceBgpv6 {
	obj.obj.DutAddress = value
	return obj
}

func (obj *deviceBgpv6) SegmentRouting() DeviceBgpv6SegmentRouting {
	if obj.obj.SegmentRouting == nil {
		obj.obj.SegmentRouting = &snappipb.DeviceBgpv6SegmentRouting{}
	}
	return &deviceBgpv6SegmentRouting{obj: obj.obj.SegmentRouting}

}

func (obj *deviceBgpv6) RouterId() string {
	return *obj.obj.RouterId
}

func (obj *deviceBgpv6) SetRouterId(value string) DeviceBgpv6 {
	obj.obj.RouterId = &value
	return obj
}

func (obj *deviceBgpv6) AsNumber() int32 {
	return obj.obj.AsNumber
}

func (obj *deviceBgpv6) SetAsNumber(value int32) DeviceBgpv6 {
	obj.obj.AsNumber = value
	return obj
}

func (obj *deviceBgpv6) Advanced() DeviceBgpAdvanced {
	if obj.obj.Advanced == nil {
		obj.obj.Advanced = &snappipb.DeviceBgpAdvanced{}
	}
	return &deviceBgpAdvanced{obj: obj.obj.Advanced}

}

func (obj *deviceBgpv6) Capability() DeviceBgpCapability {
	if obj.obj.Capability == nil {
		obj.obj.Capability = &snappipb.DeviceBgpCapability{}
	}
	return &deviceBgpCapability{obj: obj.obj.Capability}

}

func (obj *deviceBgpv6) SrTePolicies() []DeviceBgpSrTePolicy {
	if obj.obj.SrTePolicies == nil {
		obj.obj.SrTePolicies = make([]*snappipb.DeviceBgpSrTePolicy, 0)
	}
	values := make([]DeviceBgpSrTePolicy, 0)
	for _, item := range obj.obj.SrTePolicies {
		values = append(values, &deviceBgpSrTePolicy{obj: item})
	}
	return values

}

func (obj *deviceBgpv6) NewSrTePolicies() DeviceBgpSrTePolicy {
	if obj.obj.SrTePolicies == nil {
		obj.obj.SrTePolicies = make([]*snappipb.DeviceBgpSrTePolicy, 0)
	}
	slice := append(obj.obj.SrTePolicies, &snappipb.DeviceBgpSrTePolicy{})
	obj.obj.SrTePolicies = slice
	return &deviceBgpSrTePolicy{obj: slice[len(slice)-1]}
}

func (obj *deviceBgpv6) Bgpv4Routes() []DeviceBgpv4Route {
	if obj.obj.Bgpv4Routes == nil {
		obj.obj.Bgpv4Routes = make([]*snappipb.DeviceBgpv4Route, 0)
	}
	values := make([]DeviceBgpv4Route, 0)
	for _, item := range obj.obj.Bgpv4Routes {
		values = append(values, &deviceBgpv4Route{obj: item})
	}
	return values

}

func (obj *deviceBgpv6) NewBgpv4Routes() DeviceBgpv4Route {
	if obj.obj.Bgpv4Routes == nil {
		obj.obj.Bgpv4Routes = make([]*snappipb.DeviceBgpv4Route, 0)
	}
	slice := append(obj.obj.Bgpv4Routes, &snappipb.DeviceBgpv4Route{})
	obj.obj.Bgpv4Routes = slice
	return &deviceBgpv4Route{obj: slice[len(slice)-1]}
}

func (obj *deviceBgpv6) Bgpv6Routes() []DeviceBgpv6Route {
	if obj.obj.Bgpv6Routes == nil {
		obj.obj.Bgpv6Routes = make([]*snappipb.DeviceBgpv6Route, 0)
	}
	values := make([]DeviceBgpv6Route, 0)
	for _, item := range obj.obj.Bgpv6Routes {
		values = append(values, &deviceBgpv6Route{obj: item})
	}
	return values

}

func (obj *deviceBgpv6) NewBgpv6Routes() DeviceBgpv6Route {
	if obj.obj.Bgpv6Routes == nil {
		obj.obj.Bgpv6Routes = make([]*snappipb.DeviceBgpv6Route, 0)
	}
	slice := append(obj.obj.Bgpv6Routes, &snappipb.DeviceBgpv6Route{})
	obj.obj.Bgpv6Routes = slice
	return &deviceBgpv6Route{obj: slice[len(slice)-1]}
}

func (obj *deviceBgpv6) Name() string {
	return obj.obj.Name
}

func (obj *deviceBgpv6) SetName(value string) DeviceBgpv6 {
	obj.obj.Name = value
	return obj
}

func (obj *deviceBgpv6) Active() bool {
	return *obj.obj.Active
}

func (obj *deviceBgpv6) SetActive(value bool) DeviceBgpv6 {
	obj.obj.Active = &value
	return obj
}

type patternFlowEthernetDst struct {
	obj *snappipb.PatternFlowEthernetDst
}

func (obj *patternFlowEthernetDst) msg() *snappipb.PatternFlowEthernetDst {
	return obj.obj
}

func (obj *patternFlowEthernetDst) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetDst) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetDst interface {
	msg() *snappipb.PatternFlowEthernetDst
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowEthernetDst
	Values() []string
	SetValues(value []string) PatternFlowEthernetDst
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowEthernetDst
	Increment() PatternFlowEthernetDstCounter
	Decrement() PatternFlowEthernetDstCounter
}

func (obj *patternFlowEthernetDst) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowEthernetDst) SetValue(value string) PatternFlowEthernetDst {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowEthernetDst) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowEthernetDst) SetValues(value []string) PatternFlowEthernetDst {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowEthernetDst) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowEthernetDst) SetMetricGroup(value string) PatternFlowEthernetDst {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowEthernetDst) Increment() PatternFlowEthernetDstCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowEthernetDstCounter{}
	}
	return &patternFlowEthernetDstCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowEthernetDst) Decrement() PatternFlowEthernetDstCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowEthernetDstCounter{}
	}
	return &patternFlowEthernetDstCounter{obj: obj.obj.Decrement}

}

type patternFlowEthernetSrc struct {
	obj *snappipb.PatternFlowEthernetSrc
}

func (obj *patternFlowEthernetSrc) msg() *snappipb.PatternFlowEthernetSrc {
	return obj.obj
}

func (obj *patternFlowEthernetSrc) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetSrc) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetSrc interface {
	msg() *snappipb.PatternFlowEthernetSrc
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowEthernetSrc
	Values() []string
	SetValues(value []string) PatternFlowEthernetSrc
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowEthernetSrc
	Increment() PatternFlowEthernetSrcCounter
	Decrement() PatternFlowEthernetSrcCounter
}

func (obj *patternFlowEthernetSrc) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowEthernetSrc) SetValue(value string) PatternFlowEthernetSrc {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowEthernetSrc) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowEthernetSrc) SetValues(value []string) PatternFlowEthernetSrc {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowEthernetSrc) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowEthernetSrc) SetMetricGroup(value string) PatternFlowEthernetSrc {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowEthernetSrc) Increment() PatternFlowEthernetSrcCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowEthernetSrcCounter{}
	}
	return &patternFlowEthernetSrcCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowEthernetSrc) Decrement() PatternFlowEthernetSrcCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowEthernetSrcCounter{}
	}
	return &patternFlowEthernetSrcCounter{obj: obj.obj.Decrement}

}

type patternFlowEthernetEtherType struct {
	obj *snappipb.PatternFlowEthernetEtherType
}

func (obj *patternFlowEthernetEtherType) msg() *snappipb.PatternFlowEthernetEtherType {
	return obj.obj
}

func (obj *patternFlowEthernetEtherType) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetEtherType) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetEtherType interface {
	msg() *snappipb.PatternFlowEthernetEtherType
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowEthernetEtherType
	Values() []int32
	SetValues(value []int32) PatternFlowEthernetEtherType
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowEthernetEtherType
	Increment() PatternFlowEthernetEtherTypeCounter
	Decrement() PatternFlowEthernetEtherTypeCounter
}

func (obj *patternFlowEthernetEtherType) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowEthernetEtherType) SetValue(value int32) PatternFlowEthernetEtherType {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowEthernetEtherType) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowEthernetEtherType) SetValues(value []int32) PatternFlowEthernetEtherType {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowEthernetEtherType) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowEthernetEtherType) SetMetricGroup(value string) PatternFlowEthernetEtherType {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowEthernetEtherType) Increment() PatternFlowEthernetEtherTypeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowEthernetEtherTypeCounter{}
	}
	return &patternFlowEthernetEtherTypeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowEthernetEtherType) Decrement() PatternFlowEthernetEtherTypeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowEthernetEtherTypeCounter{}
	}
	return &patternFlowEthernetEtherTypeCounter{obj: obj.obj.Decrement}

}

type patternFlowEthernetPfcQueue struct {
	obj *snappipb.PatternFlowEthernetPfcQueue
}

func (obj *patternFlowEthernetPfcQueue) msg() *snappipb.PatternFlowEthernetPfcQueue {
	return obj.obj
}

func (obj *patternFlowEthernetPfcQueue) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetPfcQueue) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetPfcQueue interface {
	msg() *snappipb.PatternFlowEthernetPfcQueue
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowEthernetPfcQueue
	Values() []int32
	SetValues(value []int32) PatternFlowEthernetPfcQueue
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowEthernetPfcQueue
	Increment() PatternFlowEthernetPfcQueueCounter
	Decrement() PatternFlowEthernetPfcQueueCounter
}

func (obj *patternFlowEthernetPfcQueue) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowEthernetPfcQueue) SetValue(value int32) PatternFlowEthernetPfcQueue {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowEthernetPfcQueue) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowEthernetPfcQueue) SetValues(value []int32) PatternFlowEthernetPfcQueue {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowEthernetPfcQueue) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowEthernetPfcQueue) SetMetricGroup(value string) PatternFlowEthernetPfcQueue {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowEthernetPfcQueue) Increment() PatternFlowEthernetPfcQueueCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowEthernetPfcQueueCounter{}
	}
	return &patternFlowEthernetPfcQueueCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowEthernetPfcQueue) Decrement() PatternFlowEthernetPfcQueueCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowEthernetPfcQueueCounter{}
	}
	return &patternFlowEthernetPfcQueueCounter{obj: obj.obj.Decrement}

}

type patternFlowVlanPriority struct {
	obj *snappipb.PatternFlowVlanPriority
}

func (obj *patternFlowVlanPriority) msg() *snappipb.PatternFlowVlanPriority {
	return obj.obj
}

func (obj *patternFlowVlanPriority) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVlanPriority) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVlanPriority interface {
	msg() *snappipb.PatternFlowVlanPriority
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowVlanPriority
	Values() []int32
	SetValues(value []int32) PatternFlowVlanPriority
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowVlanPriority
	Increment() PatternFlowVlanPriorityCounter
	Decrement() PatternFlowVlanPriorityCounter
}

func (obj *patternFlowVlanPriority) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowVlanPriority) SetValue(value int32) PatternFlowVlanPriority {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowVlanPriority) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowVlanPriority) SetValues(value []int32) PatternFlowVlanPriority {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowVlanPriority) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowVlanPriority) SetMetricGroup(value string) PatternFlowVlanPriority {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowVlanPriority) Increment() PatternFlowVlanPriorityCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowVlanPriorityCounter{}
	}
	return &patternFlowVlanPriorityCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowVlanPriority) Decrement() PatternFlowVlanPriorityCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowVlanPriorityCounter{}
	}
	return &patternFlowVlanPriorityCounter{obj: obj.obj.Decrement}

}

type patternFlowVlanCfi struct {
	obj *snappipb.PatternFlowVlanCfi
}

func (obj *patternFlowVlanCfi) msg() *snappipb.PatternFlowVlanCfi {
	return obj.obj
}

func (obj *patternFlowVlanCfi) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVlanCfi) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVlanCfi interface {
	msg() *snappipb.PatternFlowVlanCfi
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowVlanCfi
	Values() []int32
	SetValues(value []int32) PatternFlowVlanCfi
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowVlanCfi
	Increment() PatternFlowVlanCfiCounter
	Decrement() PatternFlowVlanCfiCounter
}

func (obj *patternFlowVlanCfi) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowVlanCfi) SetValue(value int32) PatternFlowVlanCfi {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowVlanCfi) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowVlanCfi) SetValues(value []int32) PatternFlowVlanCfi {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowVlanCfi) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowVlanCfi) SetMetricGroup(value string) PatternFlowVlanCfi {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowVlanCfi) Increment() PatternFlowVlanCfiCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowVlanCfiCounter{}
	}
	return &patternFlowVlanCfiCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowVlanCfi) Decrement() PatternFlowVlanCfiCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowVlanCfiCounter{}
	}
	return &patternFlowVlanCfiCounter{obj: obj.obj.Decrement}

}

type patternFlowVlanId struct {
	obj *snappipb.PatternFlowVlanId
}

func (obj *patternFlowVlanId) msg() *snappipb.PatternFlowVlanId {
	return obj.obj
}

func (obj *patternFlowVlanId) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVlanId) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVlanId interface {
	msg() *snappipb.PatternFlowVlanId
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowVlanId
	Values() []int32
	SetValues(value []int32) PatternFlowVlanId
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowVlanId
	Increment() PatternFlowVlanIdCounter
	Decrement() PatternFlowVlanIdCounter
}

func (obj *patternFlowVlanId) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowVlanId) SetValue(value int32) PatternFlowVlanId {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowVlanId) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowVlanId) SetValues(value []int32) PatternFlowVlanId {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowVlanId) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowVlanId) SetMetricGroup(value string) PatternFlowVlanId {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowVlanId) Increment() PatternFlowVlanIdCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowVlanIdCounter{}
	}
	return &patternFlowVlanIdCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowVlanId) Decrement() PatternFlowVlanIdCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowVlanIdCounter{}
	}
	return &patternFlowVlanIdCounter{obj: obj.obj.Decrement}

}

type patternFlowVlanTpid struct {
	obj *snappipb.PatternFlowVlanTpid
}

func (obj *patternFlowVlanTpid) msg() *snappipb.PatternFlowVlanTpid {
	return obj.obj
}

func (obj *patternFlowVlanTpid) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVlanTpid) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVlanTpid interface {
	msg() *snappipb.PatternFlowVlanTpid
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowVlanTpid
	Values() []int32
	SetValues(value []int32) PatternFlowVlanTpid
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowVlanTpid
	Increment() PatternFlowVlanTpidCounter
	Decrement() PatternFlowVlanTpidCounter
}

func (obj *patternFlowVlanTpid) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowVlanTpid) SetValue(value int32) PatternFlowVlanTpid {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowVlanTpid) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowVlanTpid) SetValues(value []int32) PatternFlowVlanTpid {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowVlanTpid) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowVlanTpid) SetMetricGroup(value string) PatternFlowVlanTpid {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowVlanTpid) Increment() PatternFlowVlanTpidCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowVlanTpidCounter{}
	}
	return &patternFlowVlanTpidCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowVlanTpid) Decrement() PatternFlowVlanTpidCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowVlanTpidCounter{}
	}
	return &patternFlowVlanTpidCounter{obj: obj.obj.Decrement}

}

type patternFlowVxlanFlags struct {
	obj *snappipb.PatternFlowVxlanFlags
}

func (obj *patternFlowVxlanFlags) msg() *snappipb.PatternFlowVxlanFlags {
	return obj.obj
}

func (obj *patternFlowVxlanFlags) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVxlanFlags) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVxlanFlags interface {
	msg() *snappipb.PatternFlowVxlanFlags
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowVxlanFlags
	Values() []int32
	SetValues(value []int32) PatternFlowVxlanFlags
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowVxlanFlags
	Increment() PatternFlowVxlanFlagsCounter
	Decrement() PatternFlowVxlanFlagsCounter
}

func (obj *patternFlowVxlanFlags) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowVxlanFlags) SetValue(value int32) PatternFlowVxlanFlags {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowVxlanFlags) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowVxlanFlags) SetValues(value []int32) PatternFlowVxlanFlags {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowVxlanFlags) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowVxlanFlags) SetMetricGroup(value string) PatternFlowVxlanFlags {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowVxlanFlags) Increment() PatternFlowVxlanFlagsCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowVxlanFlagsCounter{}
	}
	return &patternFlowVxlanFlagsCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowVxlanFlags) Decrement() PatternFlowVxlanFlagsCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowVxlanFlagsCounter{}
	}
	return &patternFlowVxlanFlagsCounter{obj: obj.obj.Decrement}

}

type patternFlowVxlanReserved0 struct {
	obj *snappipb.PatternFlowVxlanReserved0
}

func (obj *patternFlowVxlanReserved0) msg() *snappipb.PatternFlowVxlanReserved0 {
	return obj.obj
}

func (obj *patternFlowVxlanReserved0) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVxlanReserved0) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVxlanReserved0 interface {
	msg() *snappipb.PatternFlowVxlanReserved0
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowVxlanReserved0
	Values() []int32
	SetValues(value []int32) PatternFlowVxlanReserved0
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowVxlanReserved0
	Increment() PatternFlowVxlanReserved0Counter
	Decrement() PatternFlowVxlanReserved0Counter
}

func (obj *patternFlowVxlanReserved0) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowVxlanReserved0) SetValue(value int32) PatternFlowVxlanReserved0 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowVxlanReserved0) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowVxlanReserved0) SetValues(value []int32) PatternFlowVxlanReserved0 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowVxlanReserved0) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowVxlanReserved0) SetMetricGroup(value string) PatternFlowVxlanReserved0 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowVxlanReserved0) Increment() PatternFlowVxlanReserved0Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowVxlanReserved0Counter{}
	}
	return &patternFlowVxlanReserved0Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowVxlanReserved0) Decrement() PatternFlowVxlanReserved0Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowVxlanReserved0Counter{}
	}
	return &patternFlowVxlanReserved0Counter{obj: obj.obj.Decrement}

}

type patternFlowVxlanVni struct {
	obj *snappipb.PatternFlowVxlanVni
}

func (obj *patternFlowVxlanVni) msg() *snappipb.PatternFlowVxlanVni {
	return obj.obj
}

func (obj *patternFlowVxlanVni) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVxlanVni) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVxlanVni interface {
	msg() *snappipb.PatternFlowVxlanVni
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowVxlanVni
	Values() []int32
	SetValues(value []int32) PatternFlowVxlanVni
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowVxlanVni
	Increment() PatternFlowVxlanVniCounter
	Decrement() PatternFlowVxlanVniCounter
}

func (obj *patternFlowVxlanVni) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowVxlanVni) SetValue(value int32) PatternFlowVxlanVni {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowVxlanVni) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowVxlanVni) SetValues(value []int32) PatternFlowVxlanVni {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowVxlanVni) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowVxlanVni) SetMetricGroup(value string) PatternFlowVxlanVni {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowVxlanVni) Increment() PatternFlowVxlanVniCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowVxlanVniCounter{}
	}
	return &patternFlowVxlanVniCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowVxlanVni) Decrement() PatternFlowVxlanVniCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowVxlanVniCounter{}
	}
	return &patternFlowVxlanVniCounter{obj: obj.obj.Decrement}

}

type patternFlowVxlanReserved1 struct {
	obj *snappipb.PatternFlowVxlanReserved1
}

func (obj *patternFlowVxlanReserved1) msg() *snappipb.PatternFlowVxlanReserved1 {
	return obj.obj
}

func (obj *patternFlowVxlanReserved1) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVxlanReserved1) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVxlanReserved1 interface {
	msg() *snappipb.PatternFlowVxlanReserved1
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowVxlanReserved1
	Values() []int32
	SetValues(value []int32) PatternFlowVxlanReserved1
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowVxlanReserved1
	Increment() PatternFlowVxlanReserved1Counter
	Decrement() PatternFlowVxlanReserved1Counter
}

func (obj *patternFlowVxlanReserved1) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowVxlanReserved1) SetValue(value int32) PatternFlowVxlanReserved1 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowVxlanReserved1) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowVxlanReserved1) SetValues(value []int32) PatternFlowVxlanReserved1 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowVxlanReserved1) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowVxlanReserved1) SetMetricGroup(value string) PatternFlowVxlanReserved1 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowVxlanReserved1) Increment() PatternFlowVxlanReserved1Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowVxlanReserved1Counter{}
	}
	return &patternFlowVxlanReserved1Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowVxlanReserved1) Decrement() PatternFlowVxlanReserved1Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowVxlanReserved1Counter{}
	}
	return &patternFlowVxlanReserved1Counter{obj: obj.obj.Decrement}

}

type patternFlowIpv4Version struct {
	obj *snappipb.PatternFlowIpv4Version
}

func (obj *patternFlowIpv4Version) msg() *snappipb.PatternFlowIpv4Version {
	return obj.obj
}

func (obj *patternFlowIpv4Version) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4Version) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4Version interface {
	msg() *snappipb.PatternFlowIpv4Version
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4Version
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4Version
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4Version
	Increment() PatternFlowIpv4VersionCounter
	Decrement() PatternFlowIpv4VersionCounter
}

func (obj *patternFlowIpv4Version) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4Version) SetValue(value int32) PatternFlowIpv4Version {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4Version) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4Version) SetValues(value []int32) PatternFlowIpv4Version {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4Version) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4Version) SetMetricGroup(value string) PatternFlowIpv4Version {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4Version) Increment() PatternFlowIpv4VersionCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4VersionCounter{}
	}
	return &patternFlowIpv4VersionCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4Version) Decrement() PatternFlowIpv4VersionCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4VersionCounter{}
	}
	return &patternFlowIpv4VersionCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4HeaderLength struct {
	obj *snappipb.PatternFlowIpv4HeaderLength
}

func (obj *patternFlowIpv4HeaderLength) msg() *snappipb.PatternFlowIpv4HeaderLength {
	return obj.obj
}

func (obj *patternFlowIpv4HeaderLength) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4HeaderLength) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4HeaderLength interface {
	msg() *snappipb.PatternFlowIpv4HeaderLength
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4HeaderLength
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4HeaderLength
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4HeaderLength
	Increment() PatternFlowIpv4HeaderLengthCounter
	Decrement() PatternFlowIpv4HeaderLengthCounter
}

func (obj *patternFlowIpv4HeaderLength) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4HeaderLength) SetValue(value int32) PatternFlowIpv4HeaderLength {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4HeaderLength) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4HeaderLength) SetValues(value []int32) PatternFlowIpv4HeaderLength {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4HeaderLength) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4HeaderLength) SetMetricGroup(value string) PatternFlowIpv4HeaderLength {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4HeaderLength) Increment() PatternFlowIpv4HeaderLengthCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4HeaderLengthCounter{}
	}
	return &patternFlowIpv4HeaderLengthCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4HeaderLength) Decrement() PatternFlowIpv4HeaderLengthCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4HeaderLengthCounter{}
	}
	return &patternFlowIpv4HeaderLengthCounter{obj: obj.obj.Decrement}

}

type flowIpv4Priority struct {
	obj *snappipb.FlowIpv4Priority
}

func (obj *flowIpv4Priority) msg() *snappipb.FlowIpv4Priority {
	return obj.obj
}

func (obj *flowIpv4Priority) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowIpv4Priority) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowIpv4Priority interface {
	msg() *snappipb.FlowIpv4Priority
	Yaml() string
	Json() string
	Raw() PatternFlowIpv4PriorityRaw
	Tos() FlowIpv4Tos
	Dscp() FlowIpv4Dscp
}

func (obj *flowIpv4Priority) Raw() PatternFlowIpv4PriorityRaw {
	if obj.obj.Raw == nil {
		obj.obj.Raw = &snappipb.PatternFlowIpv4PriorityRaw{}
	}
	return &patternFlowIpv4PriorityRaw{obj: obj.obj.Raw}

}

func (obj *flowIpv4Priority) Tos() FlowIpv4Tos {
	if obj.obj.Tos == nil {
		obj.obj.Tos = &snappipb.FlowIpv4Tos{}
	}
	return &flowIpv4Tos{obj: obj.obj.Tos}

}

func (obj *flowIpv4Priority) Dscp() FlowIpv4Dscp {
	if obj.obj.Dscp == nil {
		obj.obj.Dscp = &snappipb.FlowIpv4Dscp{}
	}
	return &flowIpv4Dscp{obj: obj.obj.Dscp}

}

type patternFlowIpv4TotalLength struct {
	obj *snappipb.PatternFlowIpv4TotalLength
}

func (obj *patternFlowIpv4TotalLength) msg() *snappipb.PatternFlowIpv4TotalLength {
	return obj.obj
}

func (obj *patternFlowIpv4TotalLength) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TotalLength) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TotalLength interface {
	msg() *snappipb.PatternFlowIpv4TotalLength
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4TotalLength
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4TotalLength
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4TotalLength
	Increment() PatternFlowIpv4TotalLengthCounter
	Decrement() PatternFlowIpv4TotalLengthCounter
}

func (obj *patternFlowIpv4TotalLength) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4TotalLength) SetValue(value int32) PatternFlowIpv4TotalLength {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4TotalLength) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4TotalLength) SetValues(value []int32) PatternFlowIpv4TotalLength {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4TotalLength) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4TotalLength) SetMetricGroup(value string) PatternFlowIpv4TotalLength {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4TotalLength) Increment() PatternFlowIpv4TotalLengthCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4TotalLengthCounter{}
	}
	return &patternFlowIpv4TotalLengthCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4TotalLength) Decrement() PatternFlowIpv4TotalLengthCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4TotalLengthCounter{}
	}
	return &patternFlowIpv4TotalLengthCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4Identification struct {
	obj *snappipb.PatternFlowIpv4Identification
}

func (obj *patternFlowIpv4Identification) msg() *snappipb.PatternFlowIpv4Identification {
	return obj.obj
}

func (obj *patternFlowIpv4Identification) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4Identification) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4Identification interface {
	msg() *snappipb.PatternFlowIpv4Identification
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4Identification
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4Identification
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4Identification
	Increment() PatternFlowIpv4IdentificationCounter
	Decrement() PatternFlowIpv4IdentificationCounter
}

func (obj *patternFlowIpv4Identification) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4Identification) SetValue(value int32) PatternFlowIpv4Identification {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4Identification) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4Identification) SetValues(value []int32) PatternFlowIpv4Identification {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4Identification) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4Identification) SetMetricGroup(value string) PatternFlowIpv4Identification {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4Identification) Increment() PatternFlowIpv4IdentificationCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4IdentificationCounter{}
	}
	return &patternFlowIpv4IdentificationCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4Identification) Decrement() PatternFlowIpv4IdentificationCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4IdentificationCounter{}
	}
	return &patternFlowIpv4IdentificationCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4Reserved struct {
	obj *snappipb.PatternFlowIpv4Reserved
}

func (obj *patternFlowIpv4Reserved) msg() *snappipb.PatternFlowIpv4Reserved {
	return obj.obj
}

func (obj *patternFlowIpv4Reserved) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4Reserved) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4Reserved interface {
	msg() *snappipb.PatternFlowIpv4Reserved
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4Reserved
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4Reserved
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4Reserved
	Increment() PatternFlowIpv4ReservedCounter
	Decrement() PatternFlowIpv4ReservedCounter
}

func (obj *patternFlowIpv4Reserved) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4Reserved) SetValue(value int32) PatternFlowIpv4Reserved {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4Reserved) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4Reserved) SetValues(value []int32) PatternFlowIpv4Reserved {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4Reserved) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4Reserved) SetMetricGroup(value string) PatternFlowIpv4Reserved {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4Reserved) Increment() PatternFlowIpv4ReservedCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4ReservedCounter{}
	}
	return &patternFlowIpv4ReservedCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4Reserved) Decrement() PatternFlowIpv4ReservedCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4ReservedCounter{}
	}
	return &patternFlowIpv4ReservedCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4DontFragment struct {
	obj *snappipb.PatternFlowIpv4DontFragment
}

func (obj *patternFlowIpv4DontFragment) msg() *snappipb.PatternFlowIpv4DontFragment {
	return obj.obj
}

func (obj *patternFlowIpv4DontFragment) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4DontFragment) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4DontFragment interface {
	msg() *snappipb.PatternFlowIpv4DontFragment
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4DontFragment
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4DontFragment
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4DontFragment
	Increment() PatternFlowIpv4DontFragmentCounter
	Decrement() PatternFlowIpv4DontFragmentCounter
}

func (obj *patternFlowIpv4DontFragment) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4DontFragment) SetValue(value int32) PatternFlowIpv4DontFragment {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4DontFragment) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4DontFragment) SetValues(value []int32) PatternFlowIpv4DontFragment {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4DontFragment) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4DontFragment) SetMetricGroup(value string) PatternFlowIpv4DontFragment {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4DontFragment) Increment() PatternFlowIpv4DontFragmentCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4DontFragmentCounter{}
	}
	return &patternFlowIpv4DontFragmentCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4DontFragment) Decrement() PatternFlowIpv4DontFragmentCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4DontFragmentCounter{}
	}
	return &patternFlowIpv4DontFragmentCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4MoreFragments struct {
	obj *snappipb.PatternFlowIpv4MoreFragments
}

func (obj *patternFlowIpv4MoreFragments) msg() *snappipb.PatternFlowIpv4MoreFragments {
	return obj.obj
}

func (obj *patternFlowIpv4MoreFragments) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4MoreFragments) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4MoreFragments interface {
	msg() *snappipb.PatternFlowIpv4MoreFragments
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4MoreFragments
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4MoreFragments
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4MoreFragments
	Increment() PatternFlowIpv4MoreFragmentsCounter
	Decrement() PatternFlowIpv4MoreFragmentsCounter
}

func (obj *patternFlowIpv4MoreFragments) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4MoreFragments) SetValue(value int32) PatternFlowIpv4MoreFragments {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4MoreFragments) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4MoreFragments) SetValues(value []int32) PatternFlowIpv4MoreFragments {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4MoreFragments) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4MoreFragments) SetMetricGroup(value string) PatternFlowIpv4MoreFragments {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4MoreFragments) Increment() PatternFlowIpv4MoreFragmentsCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4MoreFragmentsCounter{}
	}
	return &patternFlowIpv4MoreFragmentsCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4MoreFragments) Decrement() PatternFlowIpv4MoreFragmentsCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4MoreFragmentsCounter{}
	}
	return &patternFlowIpv4MoreFragmentsCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4FragmentOffset struct {
	obj *snappipb.PatternFlowIpv4FragmentOffset
}

func (obj *patternFlowIpv4FragmentOffset) msg() *snappipb.PatternFlowIpv4FragmentOffset {
	return obj.obj
}

func (obj *patternFlowIpv4FragmentOffset) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4FragmentOffset) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4FragmentOffset interface {
	msg() *snappipb.PatternFlowIpv4FragmentOffset
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4FragmentOffset
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4FragmentOffset
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4FragmentOffset
	Increment() PatternFlowIpv4FragmentOffsetCounter
	Decrement() PatternFlowIpv4FragmentOffsetCounter
}

func (obj *patternFlowIpv4FragmentOffset) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4FragmentOffset) SetValue(value int32) PatternFlowIpv4FragmentOffset {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4FragmentOffset) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4FragmentOffset) SetValues(value []int32) PatternFlowIpv4FragmentOffset {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4FragmentOffset) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4FragmentOffset) SetMetricGroup(value string) PatternFlowIpv4FragmentOffset {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4FragmentOffset) Increment() PatternFlowIpv4FragmentOffsetCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4FragmentOffsetCounter{}
	}
	return &patternFlowIpv4FragmentOffsetCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4FragmentOffset) Decrement() PatternFlowIpv4FragmentOffsetCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4FragmentOffsetCounter{}
	}
	return &patternFlowIpv4FragmentOffsetCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4TimeToLive struct {
	obj *snappipb.PatternFlowIpv4TimeToLive
}

func (obj *patternFlowIpv4TimeToLive) msg() *snappipb.PatternFlowIpv4TimeToLive {
	return obj.obj
}

func (obj *patternFlowIpv4TimeToLive) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TimeToLive) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TimeToLive interface {
	msg() *snappipb.PatternFlowIpv4TimeToLive
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4TimeToLive
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4TimeToLive
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4TimeToLive
	Increment() PatternFlowIpv4TimeToLiveCounter
	Decrement() PatternFlowIpv4TimeToLiveCounter
}

func (obj *patternFlowIpv4TimeToLive) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4TimeToLive) SetValue(value int32) PatternFlowIpv4TimeToLive {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4TimeToLive) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4TimeToLive) SetValues(value []int32) PatternFlowIpv4TimeToLive {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4TimeToLive) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4TimeToLive) SetMetricGroup(value string) PatternFlowIpv4TimeToLive {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4TimeToLive) Increment() PatternFlowIpv4TimeToLiveCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4TimeToLiveCounter{}
	}
	return &patternFlowIpv4TimeToLiveCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4TimeToLive) Decrement() PatternFlowIpv4TimeToLiveCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4TimeToLiveCounter{}
	}
	return &patternFlowIpv4TimeToLiveCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4Protocol struct {
	obj *snappipb.PatternFlowIpv4Protocol
}

func (obj *patternFlowIpv4Protocol) msg() *snappipb.PatternFlowIpv4Protocol {
	return obj.obj
}

func (obj *patternFlowIpv4Protocol) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4Protocol) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4Protocol interface {
	msg() *snappipb.PatternFlowIpv4Protocol
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4Protocol
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4Protocol
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4Protocol
	Increment() PatternFlowIpv4ProtocolCounter
	Decrement() PatternFlowIpv4ProtocolCounter
}

func (obj *patternFlowIpv4Protocol) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4Protocol) SetValue(value int32) PatternFlowIpv4Protocol {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4Protocol) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4Protocol) SetValues(value []int32) PatternFlowIpv4Protocol {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4Protocol) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4Protocol) SetMetricGroup(value string) PatternFlowIpv4Protocol {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4Protocol) Increment() PatternFlowIpv4ProtocolCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4ProtocolCounter{}
	}
	return &patternFlowIpv4ProtocolCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4Protocol) Decrement() PatternFlowIpv4ProtocolCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4ProtocolCounter{}
	}
	return &patternFlowIpv4ProtocolCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4HeaderChecksum struct {
	obj *snappipb.PatternFlowIpv4HeaderChecksum
}

func (obj *patternFlowIpv4HeaderChecksum) msg() *snappipb.PatternFlowIpv4HeaderChecksum {
	return obj.obj
}

func (obj *patternFlowIpv4HeaderChecksum) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4HeaderChecksum) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4HeaderChecksum interface {
	msg() *snappipb.PatternFlowIpv4HeaderChecksum
	Yaml() string
	Json() string
	Custom() int32
	SetCustom(value int32) PatternFlowIpv4HeaderChecksum
}

func (obj *patternFlowIpv4HeaderChecksum) Custom() int32 {
	return *obj.obj.Custom
}

func (obj *patternFlowIpv4HeaderChecksum) SetCustom(value int32) PatternFlowIpv4HeaderChecksum {
	obj.obj.Custom = &value
	return obj
}

type patternFlowIpv4Src struct {
	obj *snappipb.PatternFlowIpv4Src
}

func (obj *patternFlowIpv4Src) msg() *snappipb.PatternFlowIpv4Src {
	return obj.obj
}

func (obj *patternFlowIpv4Src) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4Src) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4Src interface {
	msg() *snappipb.PatternFlowIpv4Src
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowIpv4Src
	Values() []string
	SetValues(value []string) PatternFlowIpv4Src
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4Src
	Increment() PatternFlowIpv4SrcCounter
	Decrement() PatternFlowIpv4SrcCounter
}

func (obj *patternFlowIpv4Src) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4Src) SetValue(value string) PatternFlowIpv4Src {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4Src) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowIpv4Src) SetValues(value []string) PatternFlowIpv4Src {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4Src) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4Src) SetMetricGroup(value string) PatternFlowIpv4Src {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4Src) Increment() PatternFlowIpv4SrcCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4SrcCounter{}
	}
	return &patternFlowIpv4SrcCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4Src) Decrement() PatternFlowIpv4SrcCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4SrcCounter{}
	}
	return &patternFlowIpv4SrcCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4Dst struct {
	obj *snappipb.PatternFlowIpv4Dst
}

func (obj *patternFlowIpv4Dst) msg() *snappipb.PatternFlowIpv4Dst {
	return obj.obj
}

func (obj *patternFlowIpv4Dst) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4Dst) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4Dst interface {
	msg() *snappipb.PatternFlowIpv4Dst
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowIpv4Dst
	Values() []string
	SetValues(value []string) PatternFlowIpv4Dst
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4Dst
	Increment() PatternFlowIpv4DstCounter
	Decrement() PatternFlowIpv4DstCounter
}

func (obj *patternFlowIpv4Dst) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4Dst) SetValue(value string) PatternFlowIpv4Dst {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4Dst) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowIpv4Dst) SetValues(value []string) PatternFlowIpv4Dst {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4Dst) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4Dst) SetMetricGroup(value string) PatternFlowIpv4Dst {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4Dst) Increment() PatternFlowIpv4DstCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4DstCounter{}
	}
	return &patternFlowIpv4DstCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4Dst) Decrement() PatternFlowIpv4DstCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4DstCounter{}
	}
	return &patternFlowIpv4DstCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv6Version struct {
	obj *snappipb.PatternFlowIpv6Version
}

func (obj *patternFlowIpv6Version) msg() *snappipb.PatternFlowIpv6Version {
	return obj.obj
}

func (obj *patternFlowIpv6Version) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6Version) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6Version interface {
	msg() *snappipb.PatternFlowIpv6Version
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv6Version
	Values() []int32
	SetValues(value []int32) PatternFlowIpv6Version
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv6Version
	Increment() PatternFlowIpv6VersionCounter
	Decrement() PatternFlowIpv6VersionCounter
}

func (obj *patternFlowIpv6Version) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv6Version) SetValue(value int32) PatternFlowIpv6Version {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv6Version) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv6Version) SetValues(value []int32) PatternFlowIpv6Version {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv6Version) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv6Version) SetMetricGroup(value string) PatternFlowIpv6Version {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv6Version) Increment() PatternFlowIpv6VersionCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv6VersionCounter{}
	}
	return &patternFlowIpv6VersionCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv6Version) Decrement() PatternFlowIpv6VersionCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv6VersionCounter{}
	}
	return &patternFlowIpv6VersionCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv6TrafficClass struct {
	obj *snappipb.PatternFlowIpv6TrafficClass
}

func (obj *patternFlowIpv6TrafficClass) msg() *snappipb.PatternFlowIpv6TrafficClass {
	return obj.obj
}

func (obj *patternFlowIpv6TrafficClass) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6TrafficClass) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6TrafficClass interface {
	msg() *snappipb.PatternFlowIpv6TrafficClass
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv6TrafficClass
	Values() []int32
	SetValues(value []int32) PatternFlowIpv6TrafficClass
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv6TrafficClass
	Increment() PatternFlowIpv6TrafficClassCounter
	Decrement() PatternFlowIpv6TrafficClassCounter
}

func (obj *patternFlowIpv6TrafficClass) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv6TrafficClass) SetValue(value int32) PatternFlowIpv6TrafficClass {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv6TrafficClass) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv6TrafficClass) SetValues(value []int32) PatternFlowIpv6TrafficClass {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv6TrafficClass) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv6TrafficClass) SetMetricGroup(value string) PatternFlowIpv6TrafficClass {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv6TrafficClass) Increment() PatternFlowIpv6TrafficClassCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv6TrafficClassCounter{}
	}
	return &patternFlowIpv6TrafficClassCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv6TrafficClass) Decrement() PatternFlowIpv6TrafficClassCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv6TrafficClassCounter{}
	}
	return &patternFlowIpv6TrafficClassCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv6FlowLabel struct {
	obj *snappipb.PatternFlowIpv6FlowLabel
}

func (obj *patternFlowIpv6FlowLabel) msg() *snappipb.PatternFlowIpv6FlowLabel {
	return obj.obj
}

func (obj *patternFlowIpv6FlowLabel) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6FlowLabel) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6FlowLabel interface {
	msg() *snappipb.PatternFlowIpv6FlowLabel
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv6FlowLabel
	Values() []int32
	SetValues(value []int32) PatternFlowIpv6FlowLabel
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv6FlowLabel
	Increment() PatternFlowIpv6FlowLabelCounter
	Decrement() PatternFlowIpv6FlowLabelCounter
}

func (obj *patternFlowIpv6FlowLabel) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv6FlowLabel) SetValue(value int32) PatternFlowIpv6FlowLabel {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv6FlowLabel) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv6FlowLabel) SetValues(value []int32) PatternFlowIpv6FlowLabel {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv6FlowLabel) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv6FlowLabel) SetMetricGroup(value string) PatternFlowIpv6FlowLabel {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv6FlowLabel) Increment() PatternFlowIpv6FlowLabelCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv6FlowLabelCounter{}
	}
	return &patternFlowIpv6FlowLabelCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv6FlowLabel) Decrement() PatternFlowIpv6FlowLabelCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv6FlowLabelCounter{}
	}
	return &patternFlowIpv6FlowLabelCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv6PayloadLength struct {
	obj *snappipb.PatternFlowIpv6PayloadLength
}

func (obj *patternFlowIpv6PayloadLength) msg() *snappipb.PatternFlowIpv6PayloadLength {
	return obj.obj
}

func (obj *patternFlowIpv6PayloadLength) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6PayloadLength) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6PayloadLength interface {
	msg() *snappipb.PatternFlowIpv6PayloadLength
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv6PayloadLength
	Values() []int32
	SetValues(value []int32) PatternFlowIpv6PayloadLength
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv6PayloadLength
	Increment() PatternFlowIpv6PayloadLengthCounter
	Decrement() PatternFlowIpv6PayloadLengthCounter
}

func (obj *patternFlowIpv6PayloadLength) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv6PayloadLength) SetValue(value int32) PatternFlowIpv6PayloadLength {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv6PayloadLength) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv6PayloadLength) SetValues(value []int32) PatternFlowIpv6PayloadLength {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv6PayloadLength) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv6PayloadLength) SetMetricGroup(value string) PatternFlowIpv6PayloadLength {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv6PayloadLength) Increment() PatternFlowIpv6PayloadLengthCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv6PayloadLengthCounter{}
	}
	return &patternFlowIpv6PayloadLengthCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv6PayloadLength) Decrement() PatternFlowIpv6PayloadLengthCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv6PayloadLengthCounter{}
	}
	return &patternFlowIpv6PayloadLengthCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv6NextHeader struct {
	obj *snappipb.PatternFlowIpv6NextHeader
}

func (obj *patternFlowIpv6NextHeader) msg() *snappipb.PatternFlowIpv6NextHeader {
	return obj.obj
}

func (obj *patternFlowIpv6NextHeader) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6NextHeader) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6NextHeader interface {
	msg() *snappipb.PatternFlowIpv6NextHeader
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv6NextHeader
	Values() []int32
	SetValues(value []int32) PatternFlowIpv6NextHeader
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv6NextHeader
	Increment() PatternFlowIpv6NextHeaderCounter
	Decrement() PatternFlowIpv6NextHeaderCounter
}

func (obj *patternFlowIpv6NextHeader) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv6NextHeader) SetValue(value int32) PatternFlowIpv6NextHeader {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv6NextHeader) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv6NextHeader) SetValues(value []int32) PatternFlowIpv6NextHeader {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv6NextHeader) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv6NextHeader) SetMetricGroup(value string) PatternFlowIpv6NextHeader {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv6NextHeader) Increment() PatternFlowIpv6NextHeaderCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv6NextHeaderCounter{}
	}
	return &patternFlowIpv6NextHeaderCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv6NextHeader) Decrement() PatternFlowIpv6NextHeaderCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv6NextHeaderCounter{}
	}
	return &patternFlowIpv6NextHeaderCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv6HopLimit struct {
	obj *snappipb.PatternFlowIpv6HopLimit
}

func (obj *patternFlowIpv6HopLimit) msg() *snappipb.PatternFlowIpv6HopLimit {
	return obj.obj
}

func (obj *patternFlowIpv6HopLimit) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6HopLimit) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6HopLimit interface {
	msg() *snappipb.PatternFlowIpv6HopLimit
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv6HopLimit
	Values() []int32
	SetValues(value []int32) PatternFlowIpv6HopLimit
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv6HopLimit
	Increment() PatternFlowIpv6HopLimitCounter
	Decrement() PatternFlowIpv6HopLimitCounter
}

func (obj *patternFlowIpv6HopLimit) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv6HopLimit) SetValue(value int32) PatternFlowIpv6HopLimit {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv6HopLimit) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv6HopLimit) SetValues(value []int32) PatternFlowIpv6HopLimit {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv6HopLimit) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv6HopLimit) SetMetricGroup(value string) PatternFlowIpv6HopLimit {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv6HopLimit) Increment() PatternFlowIpv6HopLimitCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv6HopLimitCounter{}
	}
	return &patternFlowIpv6HopLimitCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv6HopLimit) Decrement() PatternFlowIpv6HopLimitCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv6HopLimitCounter{}
	}
	return &patternFlowIpv6HopLimitCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv6Src struct {
	obj *snappipb.PatternFlowIpv6Src
}

func (obj *patternFlowIpv6Src) msg() *snappipb.PatternFlowIpv6Src {
	return obj.obj
}

func (obj *patternFlowIpv6Src) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6Src) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6Src interface {
	msg() *snappipb.PatternFlowIpv6Src
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowIpv6Src
	Values() []string
	SetValues(value []string) PatternFlowIpv6Src
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv6Src
	Increment() PatternFlowIpv6SrcCounter
	Decrement() PatternFlowIpv6SrcCounter
}

func (obj *patternFlowIpv6Src) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowIpv6Src) SetValue(value string) PatternFlowIpv6Src {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv6Src) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowIpv6Src) SetValues(value []string) PatternFlowIpv6Src {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv6Src) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv6Src) SetMetricGroup(value string) PatternFlowIpv6Src {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv6Src) Increment() PatternFlowIpv6SrcCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv6SrcCounter{}
	}
	return &patternFlowIpv6SrcCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv6Src) Decrement() PatternFlowIpv6SrcCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv6SrcCounter{}
	}
	return &patternFlowIpv6SrcCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv6Dst struct {
	obj *snappipb.PatternFlowIpv6Dst
}

func (obj *patternFlowIpv6Dst) msg() *snappipb.PatternFlowIpv6Dst {
	return obj.obj
}

func (obj *patternFlowIpv6Dst) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6Dst) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6Dst interface {
	msg() *snappipb.PatternFlowIpv6Dst
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowIpv6Dst
	Values() []string
	SetValues(value []string) PatternFlowIpv6Dst
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv6Dst
	Increment() PatternFlowIpv6DstCounter
	Decrement() PatternFlowIpv6DstCounter
}

func (obj *patternFlowIpv6Dst) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowIpv6Dst) SetValue(value string) PatternFlowIpv6Dst {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv6Dst) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowIpv6Dst) SetValues(value []string) PatternFlowIpv6Dst {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv6Dst) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv6Dst) SetMetricGroup(value string) PatternFlowIpv6Dst {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv6Dst) Increment() PatternFlowIpv6DstCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv6DstCounter{}
	}
	return &patternFlowIpv6DstCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv6Dst) Decrement() PatternFlowIpv6DstCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv6DstCounter{}
	}
	return &patternFlowIpv6DstCounter{obj: obj.obj.Decrement}

}

type patternFlowPfcPauseDst struct {
	obj *snappipb.PatternFlowPfcPauseDst
}

func (obj *patternFlowPfcPauseDst) msg() *snappipb.PatternFlowPfcPauseDst {
	return obj.obj
}

func (obj *patternFlowPfcPauseDst) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPauseDst) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPauseDst interface {
	msg() *snappipb.PatternFlowPfcPauseDst
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowPfcPauseDst
	Values() []string
	SetValues(value []string) PatternFlowPfcPauseDst
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPfcPauseDst
	Increment() PatternFlowPfcPauseDstCounter
	Decrement() PatternFlowPfcPauseDstCounter
}

func (obj *patternFlowPfcPauseDst) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowPfcPauseDst) SetValue(value string) PatternFlowPfcPauseDst {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPfcPauseDst) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowPfcPauseDst) SetValues(value []string) PatternFlowPfcPauseDst {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPfcPauseDst) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPfcPauseDst) SetMetricGroup(value string) PatternFlowPfcPauseDst {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPfcPauseDst) Increment() PatternFlowPfcPauseDstCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPfcPauseDstCounter{}
	}
	return &patternFlowPfcPauseDstCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowPfcPauseDst) Decrement() PatternFlowPfcPauseDstCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPfcPauseDstCounter{}
	}
	return &patternFlowPfcPauseDstCounter{obj: obj.obj.Decrement}

}

type patternFlowPfcPauseSrc struct {
	obj *snappipb.PatternFlowPfcPauseSrc
}

func (obj *patternFlowPfcPauseSrc) msg() *snappipb.PatternFlowPfcPauseSrc {
	return obj.obj
}

func (obj *patternFlowPfcPauseSrc) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPauseSrc) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPauseSrc interface {
	msg() *snappipb.PatternFlowPfcPauseSrc
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowPfcPauseSrc
	Values() []string
	SetValues(value []string) PatternFlowPfcPauseSrc
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPfcPauseSrc
	Increment() PatternFlowPfcPauseSrcCounter
	Decrement() PatternFlowPfcPauseSrcCounter
}

func (obj *patternFlowPfcPauseSrc) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowPfcPauseSrc) SetValue(value string) PatternFlowPfcPauseSrc {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPfcPauseSrc) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowPfcPauseSrc) SetValues(value []string) PatternFlowPfcPauseSrc {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPfcPauseSrc) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPfcPauseSrc) SetMetricGroup(value string) PatternFlowPfcPauseSrc {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPfcPauseSrc) Increment() PatternFlowPfcPauseSrcCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPfcPauseSrcCounter{}
	}
	return &patternFlowPfcPauseSrcCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowPfcPauseSrc) Decrement() PatternFlowPfcPauseSrcCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPfcPauseSrcCounter{}
	}
	return &patternFlowPfcPauseSrcCounter{obj: obj.obj.Decrement}

}

type patternFlowPfcPauseEtherType struct {
	obj *snappipb.PatternFlowPfcPauseEtherType
}

func (obj *patternFlowPfcPauseEtherType) msg() *snappipb.PatternFlowPfcPauseEtherType {
	return obj.obj
}

func (obj *patternFlowPfcPauseEtherType) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPauseEtherType) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPauseEtherType interface {
	msg() *snappipb.PatternFlowPfcPauseEtherType
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPfcPauseEtherType
	Values() []int32
	SetValues(value []int32) PatternFlowPfcPauseEtherType
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPfcPauseEtherType
	Increment() PatternFlowPfcPauseEtherTypeCounter
	Decrement() PatternFlowPfcPauseEtherTypeCounter
}

func (obj *patternFlowPfcPauseEtherType) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPfcPauseEtherType) SetValue(value int32) PatternFlowPfcPauseEtherType {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPfcPauseEtherType) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPfcPauseEtherType) SetValues(value []int32) PatternFlowPfcPauseEtherType {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPfcPauseEtherType) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPfcPauseEtherType) SetMetricGroup(value string) PatternFlowPfcPauseEtherType {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPfcPauseEtherType) Increment() PatternFlowPfcPauseEtherTypeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPfcPauseEtherTypeCounter{}
	}
	return &patternFlowPfcPauseEtherTypeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowPfcPauseEtherType) Decrement() PatternFlowPfcPauseEtherTypeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPfcPauseEtherTypeCounter{}
	}
	return &patternFlowPfcPauseEtherTypeCounter{obj: obj.obj.Decrement}

}

type patternFlowPfcPauseControlOpCode struct {
	obj *snappipb.PatternFlowPfcPauseControlOpCode
}

func (obj *patternFlowPfcPauseControlOpCode) msg() *snappipb.PatternFlowPfcPauseControlOpCode {
	return obj.obj
}

func (obj *patternFlowPfcPauseControlOpCode) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPauseControlOpCode) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPauseControlOpCode interface {
	msg() *snappipb.PatternFlowPfcPauseControlOpCode
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPfcPauseControlOpCode
	Values() []int32
	SetValues(value []int32) PatternFlowPfcPauseControlOpCode
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPfcPauseControlOpCode
	Increment() PatternFlowPfcPauseControlOpCodeCounter
	Decrement() PatternFlowPfcPauseControlOpCodeCounter
}

func (obj *patternFlowPfcPauseControlOpCode) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPfcPauseControlOpCode) SetValue(value int32) PatternFlowPfcPauseControlOpCode {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPfcPauseControlOpCode) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPfcPauseControlOpCode) SetValues(value []int32) PatternFlowPfcPauseControlOpCode {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPfcPauseControlOpCode) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPfcPauseControlOpCode) SetMetricGroup(value string) PatternFlowPfcPauseControlOpCode {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPfcPauseControlOpCode) Increment() PatternFlowPfcPauseControlOpCodeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPfcPauseControlOpCodeCounter{}
	}
	return &patternFlowPfcPauseControlOpCodeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowPfcPauseControlOpCode) Decrement() PatternFlowPfcPauseControlOpCodeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPfcPauseControlOpCodeCounter{}
	}
	return &patternFlowPfcPauseControlOpCodeCounter{obj: obj.obj.Decrement}

}

type patternFlowPfcPauseClassEnableVector struct {
	obj *snappipb.PatternFlowPfcPauseClassEnableVector
}

func (obj *patternFlowPfcPauseClassEnableVector) msg() *snappipb.PatternFlowPfcPauseClassEnableVector {
	return obj.obj
}

func (obj *patternFlowPfcPauseClassEnableVector) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPauseClassEnableVector) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPauseClassEnableVector interface {
	msg() *snappipb.PatternFlowPfcPauseClassEnableVector
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPfcPauseClassEnableVector
	Values() []int32
	SetValues(value []int32) PatternFlowPfcPauseClassEnableVector
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPfcPauseClassEnableVector
	Increment() PatternFlowPfcPauseClassEnableVectorCounter
	Decrement() PatternFlowPfcPauseClassEnableVectorCounter
}

func (obj *patternFlowPfcPauseClassEnableVector) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPfcPauseClassEnableVector) SetValue(value int32) PatternFlowPfcPauseClassEnableVector {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPfcPauseClassEnableVector) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPfcPauseClassEnableVector) SetValues(value []int32) PatternFlowPfcPauseClassEnableVector {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPfcPauseClassEnableVector) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPfcPauseClassEnableVector) SetMetricGroup(value string) PatternFlowPfcPauseClassEnableVector {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPfcPauseClassEnableVector) Increment() PatternFlowPfcPauseClassEnableVectorCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPfcPauseClassEnableVectorCounter{}
	}
	return &patternFlowPfcPauseClassEnableVectorCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowPfcPauseClassEnableVector) Decrement() PatternFlowPfcPauseClassEnableVectorCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPfcPauseClassEnableVectorCounter{}
	}
	return &patternFlowPfcPauseClassEnableVectorCounter{obj: obj.obj.Decrement}

}

type patternFlowPfcPausePauseClass0 struct {
	obj *snappipb.PatternFlowPfcPausePauseClass0
}

func (obj *patternFlowPfcPausePauseClass0) msg() *snappipb.PatternFlowPfcPausePauseClass0 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass0) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass0) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass0 interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass0
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPfcPausePauseClass0
	Values() []int32
	SetValues(value []int32) PatternFlowPfcPausePauseClass0
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPfcPausePauseClass0
	Increment() PatternFlowPfcPausePauseClass0Counter
	Decrement() PatternFlowPfcPausePauseClass0Counter
}

func (obj *patternFlowPfcPausePauseClass0) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPfcPausePauseClass0) SetValue(value int32) PatternFlowPfcPausePauseClass0 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass0) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPfcPausePauseClass0) SetValues(value []int32) PatternFlowPfcPausePauseClass0 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPfcPausePauseClass0) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPfcPausePauseClass0) SetMetricGroup(value string) PatternFlowPfcPausePauseClass0 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass0) Increment() PatternFlowPfcPausePauseClass0Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPfcPausePauseClass0Counter{}
	}
	return &patternFlowPfcPausePauseClass0Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowPfcPausePauseClass0) Decrement() PatternFlowPfcPausePauseClass0Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPfcPausePauseClass0Counter{}
	}
	return &patternFlowPfcPausePauseClass0Counter{obj: obj.obj.Decrement}

}

type patternFlowPfcPausePauseClass1 struct {
	obj *snappipb.PatternFlowPfcPausePauseClass1
}

func (obj *patternFlowPfcPausePauseClass1) msg() *snappipb.PatternFlowPfcPausePauseClass1 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass1) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass1) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass1 interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass1
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPfcPausePauseClass1
	Values() []int32
	SetValues(value []int32) PatternFlowPfcPausePauseClass1
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPfcPausePauseClass1
	Increment() PatternFlowPfcPausePauseClass1Counter
	Decrement() PatternFlowPfcPausePauseClass1Counter
}

func (obj *patternFlowPfcPausePauseClass1) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPfcPausePauseClass1) SetValue(value int32) PatternFlowPfcPausePauseClass1 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass1) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPfcPausePauseClass1) SetValues(value []int32) PatternFlowPfcPausePauseClass1 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPfcPausePauseClass1) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPfcPausePauseClass1) SetMetricGroup(value string) PatternFlowPfcPausePauseClass1 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass1) Increment() PatternFlowPfcPausePauseClass1Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPfcPausePauseClass1Counter{}
	}
	return &patternFlowPfcPausePauseClass1Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowPfcPausePauseClass1) Decrement() PatternFlowPfcPausePauseClass1Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPfcPausePauseClass1Counter{}
	}
	return &patternFlowPfcPausePauseClass1Counter{obj: obj.obj.Decrement}

}

type patternFlowPfcPausePauseClass2 struct {
	obj *snappipb.PatternFlowPfcPausePauseClass2
}

func (obj *patternFlowPfcPausePauseClass2) msg() *snappipb.PatternFlowPfcPausePauseClass2 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass2) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass2) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass2 interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass2
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPfcPausePauseClass2
	Values() []int32
	SetValues(value []int32) PatternFlowPfcPausePauseClass2
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPfcPausePauseClass2
	Increment() PatternFlowPfcPausePauseClass2Counter
	Decrement() PatternFlowPfcPausePauseClass2Counter
}

func (obj *patternFlowPfcPausePauseClass2) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPfcPausePauseClass2) SetValue(value int32) PatternFlowPfcPausePauseClass2 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass2) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPfcPausePauseClass2) SetValues(value []int32) PatternFlowPfcPausePauseClass2 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPfcPausePauseClass2) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPfcPausePauseClass2) SetMetricGroup(value string) PatternFlowPfcPausePauseClass2 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass2) Increment() PatternFlowPfcPausePauseClass2Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPfcPausePauseClass2Counter{}
	}
	return &patternFlowPfcPausePauseClass2Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowPfcPausePauseClass2) Decrement() PatternFlowPfcPausePauseClass2Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPfcPausePauseClass2Counter{}
	}
	return &patternFlowPfcPausePauseClass2Counter{obj: obj.obj.Decrement}

}

type patternFlowPfcPausePauseClass3 struct {
	obj *snappipb.PatternFlowPfcPausePauseClass3
}

func (obj *patternFlowPfcPausePauseClass3) msg() *snappipb.PatternFlowPfcPausePauseClass3 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass3) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass3) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass3 interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass3
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPfcPausePauseClass3
	Values() []int32
	SetValues(value []int32) PatternFlowPfcPausePauseClass3
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPfcPausePauseClass3
	Increment() PatternFlowPfcPausePauseClass3Counter
	Decrement() PatternFlowPfcPausePauseClass3Counter
}

func (obj *patternFlowPfcPausePauseClass3) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPfcPausePauseClass3) SetValue(value int32) PatternFlowPfcPausePauseClass3 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass3) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPfcPausePauseClass3) SetValues(value []int32) PatternFlowPfcPausePauseClass3 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPfcPausePauseClass3) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPfcPausePauseClass3) SetMetricGroup(value string) PatternFlowPfcPausePauseClass3 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass3) Increment() PatternFlowPfcPausePauseClass3Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPfcPausePauseClass3Counter{}
	}
	return &patternFlowPfcPausePauseClass3Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowPfcPausePauseClass3) Decrement() PatternFlowPfcPausePauseClass3Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPfcPausePauseClass3Counter{}
	}
	return &patternFlowPfcPausePauseClass3Counter{obj: obj.obj.Decrement}

}

type patternFlowPfcPausePauseClass4 struct {
	obj *snappipb.PatternFlowPfcPausePauseClass4
}

func (obj *patternFlowPfcPausePauseClass4) msg() *snappipb.PatternFlowPfcPausePauseClass4 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass4) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass4) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass4 interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass4
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPfcPausePauseClass4
	Values() []int32
	SetValues(value []int32) PatternFlowPfcPausePauseClass4
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPfcPausePauseClass4
	Increment() PatternFlowPfcPausePauseClass4Counter
	Decrement() PatternFlowPfcPausePauseClass4Counter
}

func (obj *patternFlowPfcPausePauseClass4) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPfcPausePauseClass4) SetValue(value int32) PatternFlowPfcPausePauseClass4 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass4) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPfcPausePauseClass4) SetValues(value []int32) PatternFlowPfcPausePauseClass4 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPfcPausePauseClass4) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPfcPausePauseClass4) SetMetricGroup(value string) PatternFlowPfcPausePauseClass4 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass4) Increment() PatternFlowPfcPausePauseClass4Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPfcPausePauseClass4Counter{}
	}
	return &patternFlowPfcPausePauseClass4Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowPfcPausePauseClass4) Decrement() PatternFlowPfcPausePauseClass4Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPfcPausePauseClass4Counter{}
	}
	return &patternFlowPfcPausePauseClass4Counter{obj: obj.obj.Decrement}

}

type patternFlowPfcPausePauseClass5 struct {
	obj *snappipb.PatternFlowPfcPausePauseClass5
}

func (obj *patternFlowPfcPausePauseClass5) msg() *snappipb.PatternFlowPfcPausePauseClass5 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass5) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass5) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass5 interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass5
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPfcPausePauseClass5
	Values() []int32
	SetValues(value []int32) PatternFlowPfcPausePauseClass5
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPfcPausePauseClass5
	Increment() PatternFlowPfcPausePauseClass5Counter
	Decrement() PatternFlowPfcPausePauseClass5Counter
}

func (obj *patternFlowPfcPausePauseClass5) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPfcPausePauseClass5) SetValue(value int32) PatternFlowPfcPausePauseClass5 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass5) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPfcPausePauseClass5) SetValues(value []int32) PatternFlowPfcPausePauseClass5 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPfcPausePauseClass5) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPfcPausePauseClass5) SetMetricGroup(value string) PatternFlowPfcPausePauseClass5 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass5) Increment() PatternFlowPfcPausePauseClass5Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPfcPausePauseClass5Counter{}
	}
	return &patternFlowPfcPausePauseClass5Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowPfcPausePauseClass5) Decrement() PatternFlowPfcPausePauseClass5Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPfcPausePauseClass5Counter{}
	}
	return &patternFlowPfcPausePauseClass5Counter{obj: obj.obj.Decrement}

}

type patternFlowPfcPausePauseClass6 struct {
	obj *snappipb.PatternFlowPfcPausePauseClass6
}

func (obj *patternFlowPfcPausePauseClass6) msg() *snappipb.PatternFlowPfcPausePauseClass6 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass6) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass6) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass6 interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass6
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPfcPausePauseClass6
	Values() []int32
	SetValues(value []int32) PatternFlowPfcPausePauseClass6
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPfcPausePauseClass6
	Increment() PatternFlowPfcPausePauseClass6Counter
	Decrement() PatternFlowPfcPausePauseClass6Counter
}

func (obj *patternFlowPfcPausePauseClass6) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPfcPausePauseClass6) SetValue(value int32) PatternFlowPfcPausePauseClass6 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass6) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPfcPausePauseClass6) SetValues(value []int32) PatternFlowPfcPausePauseClass6 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPfcPausePauseClass6) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPfcPausePauseClass6) SetMetricGroup(value string) PatternFlowPfcPausePauseClass6 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass6) Increment() PatternFlowPfcPausePauseClass6Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPfcPausePauseClass6Counter{}
	}
	return &patternFlowPfcPausePauseClass6Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowPfcPausePauseClass6) Decrement() PatternFlowPfcPausePauseClass6Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPfcPausePauseClass6Counter{}
	}
	return &patternFlowPfcPausePauseClass6Counter{obj: obj.obj.Decrement}

}

type patternFlowPfcPausePauseClass7 struct {
	obj *snappipb.PatternFlowPfcPausePauseClass7
}

func (obj *patternFlowPfcPausePauseClass7) msg() *snappipb.PatternFlowPfcPausePauseClass7 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass7) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass7) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass7 interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass7
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPfcPausePauseClass7
	Values() []int32
	SetValues(value []int32) PatternFlowPfcPausePauseClass7
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPfcPausePauseClass7
	Increment() PatternFlowPfcPausePauseClass7Counter
	Decrement() PatternFlowPfcPausePauseClass7Counter
}

func (obj *patternFlowPfcPausePauseClass7) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPfcPausePauseClass7) SetValue(value int32) PatternFlowPfcPausePauseClass7 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass7) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPfcPausePauseClass7) SetValues(value []int32) PatternFlowPfcPausePauseClass7 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPfcPausePauseClass7) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPfcPausePauseClass7) SetMetricGroup(value string) PatternFlowPfcPausePauseClass7 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass7) Increment() PatternFlowPfcPausePauseClass7Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPfcPausePauseClass7Counter{}
	}
	return &patternFlowPfcPausePauseClass7Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowPfcPausePauseClass7) Decrement() PatternFlowPfcPausePauseClass7Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPfcPausePauseClass7Counter{}
	}
	return &patternFlowPfcPausePauseClass7Counter{obj: obj.obj.Decrement}

}

type patternFlowEthernetPauseDst struct {
	obj *snappipb.PatternFlowEthernetPauseDst
}

func (obj *patternFlowEthernetPauseDst) msg() *snappipb.PatternFlowEthernetPauseDst {
	return obj.obj
}

func (obj *patternFlowEthernetPauseDst) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetPauseDst) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetPauseDst interface {
	msg() *snappipb.PatternFlowEthernetPauseDst
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowEthernetPauseDst
	Values() []string
	SetValues(value []string) PatternFlowEthernetPauseDst
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowEthernetPauseDst
	Increment() PatternFlowEthernetPauseDstCounter
	Decrement() PatternFlowEthernetPauseDstCounter
}

func (obj *patternFlowEthernetPauseDst) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowEthernetPauseDst) SetValue(value string) PatternFlowEthernetPauseDst {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowEthernetPauseDst) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowEthernetPauseDst) SetValues(value []string) PatternFlowEthernetPauseDst {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowEthernetPauseDst) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowEthernetPauseDst) SetMetricGroup(value string) PatternFlowEthernetPauseDst {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowEthernetPauseDst) Increment() PatternFlowEthernetPauseDstCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowEthernetPauseDstCounter{}
	}
	return &patternFlowEthernetPauseDstCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowEthernetPauseDst) Decrement() PatternFlowEthernetPauseDstCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowEthernetPauseDstCounter{}
	}
	return &patternFlowEthernetPauseDstCounter{obj: obj.obj.Decrement}

}

type patternFlowEthernetPauseSrc struct {
	obj *snappipb.PatternFlowEthernetPauseSrc
}

func (obj *patternFlowEthernetPauseSrc) msg() *snappipb.PatternFlowEthernetPauseSrc {
	return obj.obj
}

func (obj *patternFlowEthernetPauseSrc) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetPauseSrc) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetPauseSrc interface {
	msg() *snappipb.PatternFlowEthernetPauseSrc
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowEthernetPauseSrc
	Values() []string
	SetValues(value []string) PatternFlowEthernetPauseSrc
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowEthernetPauseSrc
	Increment() PatternFlowEthernetPauseSrcCounter
	Decrement() PatternFlowEthernetPauseSrcCounter
}

func (obj *patternFlowEthernetPauseSrc) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowEthernetPauseSrc) SetValue(value string) PatternFlowEthernetPauseSrc {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowEthernetPauseSrc) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowEthernetPauseSrc) SetValues(value []string) PatternFlowEthernetPauseSrc {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowEthernetPauseSrc) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowEthernetPauseSrc) SetMetricGroup(value string) PatternFlowEthernetPauseSrc {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowEthernetPauseSrc) Increment() PatternFlowEthernetPauseSrcCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowEthernetPauseSrcCounter{}
	}
	return &patternFlowEthernetPauseSrcCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowEthernetPauseSrc) Decrement() PatternFlowEthernetPauseSrcCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowEthernetPauseSrcCounter{}
	}
	return &patternFlowEthernetPauseSrcCounter{obj: obj.obj.Decrement}

}

type patternFlowEthernetPauseEtherType struct {
	obj *snappipb.PatternFlowEthernetPauseEtherType
}

func (obj *patternFlowEthernetPauseEtherType) msg() *snappipb.PatternFlowEthernetPauseEtherType {
	return obj.obj
}

func (obj *patternFlowEthernetPauseEtherType) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetPauseEtherType) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetPauseEtherType interface {
	msg() *snappipb.PatternFlowEthernetPauseEtherType
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowEthernetPauseEtherType
	Values() []int32
	SetValues(value []int32) PatternFlowEthernetPauseEtherType
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowEthernetPauseEtherType
	Increment() PatternFlowEthernetPauseEtherTypeCounter
	Decrement() PatternFlowEthernetPauseEtherTypeCounter
}

func (obj *patternFlowEthernetPauseEtherType) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowEthernetPauseEtherType) SetValue(value int32) PatternFlowEthernetPauseEtherType {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowEthernetPauseEtherType) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowEthernetPauseEtherType) SetValues(value []int32) PatternFlowEthernetPauseEtherType {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowEthernetPauseEtherType) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowEthernetPauseEtherType) SetMetricGroup(value string) PatternFlowEthernetPauseEtherType {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowEthernetPauseEtherType) Increment() PatternFlowEthernetPauseEtherTypeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowEthernetPauseEtherTypeCounter{}
	}
	return &patternFlowEthernetPauseEtherTypeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowEthernetPauseEtherType) Decrement() PatternFlowEthernetPauseEtherTypeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowEthernetPauseEtherTypeCounter{}
	}
	return &patternFlowEthernetPauseEtherTypeCounter{obj: obj.obj.Decrement}

}

type patternFlowEthernetPauseControlOpCode struct {
	obj *snappipb.PatternFlowEthernetPauseControlOpCode
}

func (obj *patternFlowEthernetPauseControlOpCode) msg() *snappipb.PatternFlowEthernetPauseControlOpCode {
	return obj.obj
}

func (obj *patternFlowEthernetPauseControlOpCode) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetPauseControlOpCode) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetPauseControlOpCode interface {
	msg() *snappipb.PatternFlowEthernetPauseControlOpCode
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowEthernetPauseControlOpCode
	Values() []int32
	SetValues(value []int32) PatternFlowEthernetPauseControlOpCode
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowEthernetPauseControlOpCode
	Increment() PatternFlowEthernetPauseControlOpCodeCounter
	Decrement() PatternFlowEthernetPauseControlOpCodeCounter
}

func (obj *patternFlowEthernetPauseControlOpCode) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowEthernetPauseControlOpCode) SetValue(value int32) PatternFlowEthernetPauseControlOpCode {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowEthernetPauseControlOpCode) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowEthernetPauseControlOpCode) SetValues(value []int32) PatternFlowEthernetPauseControlOpCode {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowEthernetPauseControlOpCode) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowEthernetPauseControlOpCode) SetMetricGroup(value string) PatternFlowEthernetPauseControlOpCode {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowEthernetPauseControlOpCode) Increment() PatternFlowEthernetPauseControlOpCodeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowEthernetPauseControlOpCodeCounter{}
	}
	return &patternFlowEthernetPauseControlOpCodeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowEthernetPauseControlOpCode) Decrement() PatternFlowEthernetPauseControlOpCodeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowEthernetPauseControlOpCodeCounter{}
	}
	return &patternFlowEthernetPauseControlOpCodeCounter{obj: obj.obj.Decrement}

}

type patternFlowEthernetPauseTime struct {
	obj *snappipb.PatternFlowEthernetPauseTime
}

func (obj *patternFlowEthernetPauseTime) msg() *snappipb.PatternFlowEthernetPauseTime {
	return obj.obj
}

func (obj *patternFlowEthernetPauseTime) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetPauseTime) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetPauseTime interface {
	msg() *snappipb.PatternFlowEthernetPauseTime
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowEthernetPauseTime
	Values() []int32
	SetValues(value []int32) PatternFlowEthernetPauseTime
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowEthernetPauseTime
	Increment() PatternFlowEthernetPauseTimeCounter
	Decrement() PatternFlowEthernetPauseTimeCounter
}

func (obj *patternFlowEthernetPauseTime) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowEthernetPauseTime) SetValue(value int32) PatternFlowEthernetPauseTime {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowEthernetPauseTime) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowEthernetPauseTime) SetValues(value []int32) PatternFlowEthernetPauseTime {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowEthernetPauseTime) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowEthernetPauseTime) SetMetricGroup(value string) PatternFlowEthernetPauseTime {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowEthernetPauseTime) Increment() PatternFlowEthernetPauseTimeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowEthernetPauseTimeCounter{}
	}
	return &patternFlowEthernetPauseTimeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowEthernetPauseTime) Decrement() PatternFlowEthernetPauseTimeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowEthernetPauseTimeCounter{}
	}
	return &patternFlowEthernetPauseTimeCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpSrcPort struct {
	obj *snappipb.PatternFlowTcpSrcPort
}

func (obj *patternFlowTcpSrcPort) msg() *snappipb.PatternFlowTcpSrcPort {
	return obj.obj
}

func (obj *patternFlowTcpSrcPort) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpSrcPort) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpSrcPort interface {
	msg() *snappipb.PatternFlowTcpSrcPort
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpSrcPort
	Values() []int32
	SetValues(value []int32) PatternFlowTcpSrcPort
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpSrcPort
	Increment() PatternFlowTcpSrcPortCounter
	Decrement() PatternFlowTcpSrcPortCounter
}

func (obj *patternFlowTcpSrcPort) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpSrcPort) SetValue(value int32) PatternFlowTcpSrcPort {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpSrcPort) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpSrcPort) SetValues(value []int32) PatternFlowTcpSrcPort {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpSrcPort) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpSrcPort) SetMetricGroup(value string) PatternFlowTcpSrcPort {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpSrcPort) Increment() PatternFlowTcpSrcPortCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpSrcPortCounter{}
	}
	return &patternFlowTcpSrcPortCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpSrcPort) Decrement() PatternFlowTcpSrcPortCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpSrcPortCounter{}
	}
	return &patternFlowTcpSrcPortCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpDstPort struct {
	obj *snappipb.PatternFlowTcpDstPort
}

func (obj *patternFlowTcpDstPort) msg() *snappipb.PatternFlowTcpDstPort {
	return obj.obj
}

func (obj *patternFlowTcpDstPort) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpDstPort) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpDstPort interface {
	msg() *snappipb.PatternFlowTcpDstPort
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpDstPort
	Values() []int32
	SetValues(value []int32) PatternFlowTcpDstPort
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpDstPort
	Increment() PatternFlowTcpDstPortCounter
	Decrement() PatternFlowTcpDstPortCounter
}

func (obj *patternFlowTcpDstPort) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpDstPort) SetValue(value int32) PatternFlowTcpDstPort {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpDstPort) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpDstPort) SetValues(value []int32) PatternFlowTcpDstPort {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpDstPort) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpDstPort) SetMetricGroup(value string) PatternFlowTcpDstPort {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpDstPort) Increment() PatternFlowTcpDstPortCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpDstPortCounter{}
	}
	return &patternFlowTcpDstPortCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpDstPort) Decrement() PatternFlowTcpDstPortCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpDstPortCounter{}
	}
	return &patternFlowTcpDstPortCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpSeqNum struct {
	obj *snappipb.PatternFlowTcpSeqNum
}

func (obj *patternFlowTcpSeqNum) msg() *snappipb.PatternFlowTcpSeqNum {
	return obj.obj
}

func (obj *patternFlowTcpSeqNum) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpSeqNum) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpSeqNum interface {
	msg() *snappipb.PatternFlowTcpSeqNum
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpSeqNum
	Values() []int32
	SetValues(value []int32) PatternFlowTcpSeqNum
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpSeqNum
	Increment() PatternFlowTcpSeqNumCounter
	Decrement() PatternFlowTcpSeqNumCounter
}

func (obj *patternFlowTcpSeqNum) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpSeqNum) SetValue(value int32) PatternFlowTcpSeqNum {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpSeqNum) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpSeqNum) SetValues(value []int32) PatternFlowTcpSeqNum {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpSeqNum) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpSeqNum) SetMetricGroup(value string) PatternFlowTcpSeqNum {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpSeqNum) Increment() PatternFlowTcpSeqNumCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpSeqNumCounter{}
	}
	return &patternFlowTcpSeqNumCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpSeqNum) Decrement() PatternFlowTcpSeqNumCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpSeqNumCounter{}
	}
	return &patternFlowTcpSeqNumCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpAckNum struct {
	obj *snappipb.PatternFlowTcpAckNum
}

func (obj *patternFlowTcpAckNum) msg() *snappipb.PatternFlowTcpAckNum {
	return obj.obj
}

func (obj *patternFlowTcpAckNum) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpAckNum) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpAckNum interface {
	msg() *snappipb.PatternFlowTcpAckNum
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpAckNum
	Values() []int32
	SetValues(value []int32) PatternFlowTcpAckNum
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpAckNum
	Increment() PatternFlowTcpAckNumCounter
	Decrement() PatternFlowTcpAckNumCounter
}

func (obj *patternFlowTcpAckNum) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpAckNum) SetValue(value int32) PatternFlowTcpAckNum {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpAckNum) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpAckNum) SetValues(value []int32) PatternFlowTcpAckNum {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpAckNum) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpAckNum) SetMetricGroup(value string) PatternFlowTcpAckNum {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpAckNum) Increment() PatternFlowTcpAckNumCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpAckNumCounter{}
	}
	return &patternFlowTcpAckNumCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpAckNum) Decrement() PatternFlowTcpAckNumCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpAckNumCounter{}
	}
	return &patternFlowTcpAckNumCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpDataOffset struct {
	obj *snappipb.PatternFlowTcpDataOffset
}

func (obj *patternFlowTcpDataOffset) msg() *snappipb.PatternFlowTcpDataOffset {
	return obj.obj
}

func (obj *patternFlowTcpDataOffset) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpDataOffset) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpDataOffset interface {
	msg() *snappipb.PatternFlowTcpDataOffset
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpDataOffset
	Values() []int32
	SetValues(value []int32) PatternFlowTcpDataOffset
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpDataOffset
	Increment() PatternFlowTcpDataOffsetCounter
	Decrement() PatternFlowTcpDataOffsetCounter
}

func (obj *patternFlowTcpDataOffset) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpDataOffset) SetValue(value int32) PatternFlowTcpDataOffset {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpDataOffset) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpDataOffset) SetValues(value []int32) PatternFlowTcpDataOffset {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpDataOffset) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpDataOffset) SetMetricGroup(value string) PatternFlowTcpDataOffset {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpDataOffset) Increment() PatternFlowTcpDataOffsetCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpDataOffsetCounter{}
	}
	return &patternFlowTcpDataOffsetCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpDataOffset) Decrement() PatternFlowTcpDataOffsetCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpDataOffsetCounter{}
	}
	return &patternFlowTcpDataOffsetCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpEcnNs struct {
	obj *snappipb.PatternFlowTcpEcnNs
}

func (obj *patternFlowTcpEcnNs) msg() *snappipb.PatternFlowTcpEcnNs {
	return obj.obj
}

func (obj *patternFlowTcpEcnNs) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpEcnNs) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpEcnNs interface {
	msg() *snappipb.PatternFlowTcpEcnNs
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpEcnNs
	Values() []int32
	SetValues(value []int32) PatternFlowTcpEcnNs
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpEcnNs
	Increment() PatternFlowTcpEcnNsCounter
	Decrement() PatternFlowTcpEcnNsCounter
}

func (obj *patternFlowTcpEcnNs) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpEcnNs) SetValue(value int32) PatternFlowTcpEcnNs {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpEcnNs) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpEcnNs) SetValues(value []int32) PatternFlowTcpEcnNs {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpEcnNs) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpEcnNs) SetMetricGroup(value string) PatternFlowTcpEcnNs {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpEcnNs) Increment() PatternFlowTcpEcnNsCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpEcnNsCounter{}
	}
	return &patternFlowTcpEcnNsCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpEcnNs) Decrement() PatternFlowTcpEcnNsCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpEcnNsCounter{}
	}
	return &patternFlowTcpEcnNsCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpEcnCwr struct {
	obj *snappipb.PatternFlowTcpEcnCwr
}

func (obj *patternFlowTcpEcnCwr) msg() *snappipb.PatternFlowTcpEcnCwr {
	return obj.obj
}

func (obj *patternFlowTcpEcnCwr) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpEcnCwr) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpEcnCwr interface {
	msg() *snappipb.PatternFlowTcpEcnCwr
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpEcnCwr
	Values() []int32
	SetValues(value []int32) PatternFlowTcpEcnCwr
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpEcnCwr
	Increment() PatternFlowTcpEcnCwrCounter
	Decrement() PatternFlowTcpEcnCwrCounter
}

func (obj *patternFlowTcpEcnCwr) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpEcnCwr) SetValue(value int32) PatternFlowTcpEcnCwr {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpEcnCwr) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpEcnCwr) SetValues(value []int32) PatternFlowTcpEcnCwr {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpEcnCwr) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpEcnCwr) SetMetricGroup(value string) PatternFlowTcpEcnCwr {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpEcnCwr) Increment() PatternFlowTcpEcnCwrCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpEcnCwrCounter{}
	}
	return &patternFlowTcpEcnCwrCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpEcnCwr) Decrement() PatternFlowTcpEcnCwrCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpEcnCwrCounter{}
	}
	return &patternFlowTcpEcnCwrCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpEcnEcho struct {
	obj *snappipb.PatternFlowTcpEcnEcho
}

func (obj *patternFlowTcpEcnEcho) msg() *snappipb.PatternFlowTcpEcnEcho {
	return obj.obj
}

func (obj *patternFlowTcpEcnEcho) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpEcnEcho) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpEcnEcho interface {
	msg() *snappipb.PatternFlowTcpEcnEcho
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpEcnEcho
	Values() []int32
	SetValues(value []int32) PatternFlowTcpEcnEcho
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpEcnEcho
	Increment() PatternFlowTcpEcnEchoCounter
	Decrement() PatternFlowTcpEcnEchoCounter
}

func (obj *patternFlowTcpEcnEcho) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpEcnEcho) SetValue(value int32) PatternFlowTcpEcnEcho {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpEcnEcho) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpEcnEcho) SetValues(value []int32) PatternFlowTcpEcnEcho {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpEcnEcho) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpEcnEcho) SetMetricGroup(value string) PatternFlowTcpEcnEcho {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpEcnEcho) Increment() PatternFlowTcpEcnEchoCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpEcnEchoCounter{}
	}
	return &patternFlowTcpEcnEchoCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpEcnEcho) Decrement() PatternFlowTcpEcnEchoCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpEcnEchoCounter{}
	}
	return &patternFlowTcpEcnEchoCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpCtlUrg struct {
	obj *snappipb.PatternFlowTcpCtlUrg
}

func (obj *patternFlowTcpCtlUrg) msg() *snappipb.PatternFlowTcpCtlUrg {
	return obj.obj
}

func (obj *patternFlowTcpCtlUrg) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpCtlUrg) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpCtlUrg interface {
	msg() *snappipb.PatternFlowTcpCtlUrg
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpCtlUrg
	Values() []int32
	SetValues(value []int32) PatternFlowTcpCtlUrg
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpCtlUrg
	Increment() PatternFlowTcpCtlUrgCounter
	Decrement() PatternFlowTcpCtlUrgCounter
}

func (obj *patternFlowTcpCtlUrg) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpCtlUrg) SetValue(value int32) PatternFlowTcpCtlUrg {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpCtlUrg) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpCtlUrg) SetValues(value []int32) PatternFlowTcpCtlUrg {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpCtlUrg) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpCtlUrg) SetMetricGroup(value string) PatternFlowTcpCtlUrg {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpCtlUrg) Increment() PatternFlowTcpCtlUrgCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpCtlUrgCounter{}
	}
	return &patternFlowTcpCtlUrgCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpCtlUrg) Decrement() PatternFlowTcpCtlUrgCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpCtlUrgCounter{}
	}
	return &patternFlowTcpCtlUrgCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpCtlAck struct {
	obj *snappipb.PatternFlowTcpCtlAck
}

func (obj *patternFlowTcpCtlAck) msg() *snappipb.PatternFlowTcpCtlAck {
	return obj.obj
}

func (obj *patternFlowTcpCtlAck) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpCtlAck) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpCtlAck interface {
	msg() *snappipb.PatternFlowTcpCtlAck
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpCtlAck
	Values() []int32
	SetValues(value []int32) PatternFlowTcpCtlAck
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpCtlAck
	Increment() PatternFlowTcpCtlAckCounter
	Decrement() PatternFlowTcpCtlAckCounter
}

func (obj *patternFlowTcpCtlAck) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpCtlAck) SetValue(value int32) PatternFlowTcpCtlAck {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpCtlAck) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpCtlAck) SetValues(value []int32) PatternFlowTcpCtlAck {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpCtlAck) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpCtlAck) SetMetricGroup(value string) PatternFlowTcpCtlAck {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpCtlAck) Increment() PatternFlowTcpCtlAckCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpCtlAckCounter{}
	}
	return &patternFlowTcpCtlAckCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpCtlAck) Decrement() PatternFlowTcpCtlAckCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpCtlAckCounter{}
	}
	return &patternFlowTcpCtlAckCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpCtlPsh struct {
	obj *snappipb.PatternFlowTcpCtlPsh
}

func (obj *patternFlowTcpCtlPsh) msg() *snappipb.PatternFlowTcpCtlPsh {
	return obj.obj
}

func (obj *patternFlowTcpCtlPsh) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpCtlPsh) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpCtlPsh interface {
	msg() *snappipb.PatternFlowTcpCtlPsh
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpCtlPsh
	Values() []int32
	SetValues(value []int32) PatternFlowTcpCtlPsh
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpCtlPsh
	Increment() PatternFlowTcpCtlPshCounter
	Decrement() PatternFlowTcpCtlPshCounter
}

func (obj *patternFlowTcpCtlPsh) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpCtlPsh) SetValue(value int32) PatternFlowTcpCtlPsh {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpCtlPsh) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpCtlPsh) SetValues(value []int32) PatternFlowTcpCtlPsh {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpCtlPsh) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpCtlPsh) SetMetricGroup(value string) PatternFlowTcpCtlPsh {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpCtlPsh) Increment() PatternFlowTcpCtlPshCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpCtlPshCounter{}
	}
	return &patternFlowTcpCtlPshCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpCtlPsh) Decrement() PatternFlowTcpCtlPshCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpCtlPshCounter{}
	}
	return &patternFlowTcpCtlPshCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpCtlRst struct {
	obj *snappipb.PatternFlowTcpCtlRst
}

func (obj *patternFlowTcpCtlRst) msg() *snappipb.PatternFlowTcpCtlRst {
	return obj.obj
}

func (obj *patternFlowTcpCtlRst) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpCtlRst) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpCtlRst interface {
	msg() *snappipb.PatternFlowTcpCtlRst
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpCtlRst
	Values() []int32
	SetValues(value []int32) PatternFlowTcpCtlRst
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpCtlRst
	Increment() PatternFlowTcpCtlRstCounter
	Decrement() PatternFlowTcpCtlRstCounter
}

func (obj *patternFlowTcpCtlRst) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpCtlRst) SetValue(value int32) PatternFlowTcpCtlRst {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpCtlRst) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpCtlRst) SetValues(value []int32) PatternFlowTcpCtlRst {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpCtlRst) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpCtlRst) SetMetricGroup(value string) PatternFlowTcpCtlRst {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpCtlRst) Increment() PatternFlowTcpCtlRstCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpCtlRstCounter{}
	}
	return &patternFlowTcpCtlRstCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpCtlRst) Decrement() PatternFlowTcpCtlRstCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpCtlRstCounter{}
	}
	return &patternFlowTcpCtlRstCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpCtlSyn struct {
	obj *snappipb.PatternFlowTcpCtlSyn
}

func (obj *patternFlowTcpCtlSyn) msg() *snappipb.PatternFlowTcpCtlSyn {
	return obj.obj
}

func (obj *patternFlowTcpCtlSyn) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpCtlSyn) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpCtlSyn interface {
	msg() *snappipb.PatternFlowTcpCtlSyn
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpCtlSyn
	Values() []int32
	SetValues(value []int32) PatternFlowTcpCtlSyn
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpCtlSyn
	Increment() PatternFlowTcpCtlSynCounter
	Decrement() PatternFlowTcpCtlSynCounter
}

func (obj *patternFlowTcpCtlSyn) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpCtlSyn) SetValue(value int32) PatternFlowTcpCtlSyn {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpCtlSyn) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpCtlSyn) SetValues(value []int32) PatternFlowTcpCtlSyn {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpCtlSyn) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpCtlSyn) SetMetricGroup(value string) PatternFlowTcpCtlSyn {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpCtlSyn) Increment() PatternFlowTcpCtlSynCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpCtlSynCounter{}
	}
	return &patternFlowTcpCtlSynCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpCtlSyn) Decrement() PatternFlowTcpCtlSynCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpCtlSynCounter{}
	}
	return &patternFlowTcpCtlSynCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpCtlFin struct {
	obj *snappipb.PatternFlowTcpCtlFin
}

func (obj *patternFlowTcpCtlFin) msg() *snappipb.PatternFlowTcpCtlFin {
	return obj.obj
}

func (obj *patternFlowTcpCtlFin) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpCtlFin) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpCtlFin interface {
	msg() *snappipb.PatternFlowTcpCtlFin
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpCtlFin
	Values() []int32
	SetValues(value []int32) PatternFlowTcpCtlFin
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpCtlFin
	Increment() PatternFlowTcpCtlFinCounter
	Decrement() PatternFlowTcpCtlFinCounter
}

func (obj *patternFlowTcpCtlFin) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpCtlFin) SetValue(value int32) PatternFlowTcpCtlFin {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpCtlFin) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpCtlFin) SetValues(value []int32) PatternFlowTcpCtlFin {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpCtlFin) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpCtlFin) SetMetricGroup(value string) PatternFlowTcpCtlFin {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpCtlFin) Increment() PatternFlowTcpCtlFinCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpCtlFinCounter{}
	}
	return &patternFlowTcpCtlFinCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpCtlFin) Decrement() PatternFlowTcpCtlFinCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpCtlFinCounter{}
	}
	return &patternFlowTcpCtlFinCounter{obj: obj.obj.Decrement}

}

type patternFlowTcpWindow struct {
	obj *snappipb.PatternFlowTcpWindow
}

func (obj *patternFlowTcpWindow) msg() *snappipb.PatternFlowTcpWindow {
	return obj.obj
}

func (obj *patternFlowTcpWindow) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpWindow) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpWindow interface {
	msg() *snappipb.PatternFlowTcpWindow
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowTcpWindow
	Values() []int32
	SetValues(value []int32) PatternFlowTcpWindow
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowTcpWindow
	Increment() PatternFlowTcpWindowCounter
	Decrement() PatternFlowTcpWindowCounter
}

func (obj *patternFlowTcpWindow) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowTcpWindow) SetValue(value int32) PatternFlowTcpWindow {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowTcpWindow) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowTcpWindow) SetValues(value []int32) PatternFlowTcpWindow {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowTcpWindow) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowTcpWindow) SetMetricGroup(value string) PatternFlowTcpWindow {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowTcpWindow) Increment() PatternFlowTcpWindowCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowTcpWindowCounter{}
	}
	return &patternFlowTcpWindowCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowTcpWindow) Decrement() PatternFlowTcpWindowCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowTcpWindowCounter{}
	}
	return &patternFlowTcpWindowCounter{obj: obj.obj.Decrement}

}

type patternFlowUdpSrcPort struct {
	obj *snappipb.PatternFlowUdpSrcPort
}

func (obj *patternFlowUdpSrcPort) msg() *snappipb.PatternFlowUdpSrcPort {
	return obj.obj
}

func (obj *patternFlowUdpSrcPort) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowUdpSrcPort) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowUdpSrcPort interface {
	msg() *snappipb.PatternFlowUdpSrcPort
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowUdpSrcPort
	Values() []int32
	SetValues(value []int32) PatternFlowUdpSrcPort
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowUdpSrcPort
	Increment() PatternFlowUdpSrcPortCounter
	Decrement() PatternFlowUdpSrcPortCounter
}

func (obj *patternFlowUdpSrcPort) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowUdpSrcPort) SetValue(value int32) PatternFlowUdpSrcPort {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowUdpSrcPort) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowUdpSrcPort) SetValues(value []int32) PatternFlowUdpSrcPort {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowUdpSrcPort) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowUdpSrcPort) SetMetricGroup(value string) PatternFlowUdpSrcPort {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowUdpSrcPort) Increment() PatternFlowUdpSrcPortCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowUdpSrcPortCounter{}
	}
	return &patternFlowUdpSrcPortCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowUdpSrcPort) Decrement() PatternFlowUdpSrcPortCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowUdpSrcPortCounter{}
	}
	return &patternFlowUdpSrcPortCounter{obj: obj.obj.Decrement}

}

type patternFlowUdpDstPort struct {
	obj *snappipb.PatternFlowUdpDstPort
}

func (obj *patternFlowUdpDstPort) msg() *snappipb.PatternFlowUdpDstPort {
	return obj.obj
}

func (obj *patternFlowUdpDstPort) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowUdpDstPort) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowUdpDstPort interface {
	msg() *snappipb.PatternFlowUdpDstPort
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowUdpDstPort
	Values() []int32
	SetValues(value []int32) PatternFlowUdpDstPort
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowUdpDstPort
	Increment() PatternFlowUdpDstPortCounter
	Decrement() PatternFlowUdpDstPortCounter
}

func (obj *patternFlowUdpDstPort) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowUdpDstPort) SetValue(value int32) PatternFlowUdpDstPort {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowUdpDstPort) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowUdpDstPort) SetValues(value []int32) PatternFlowUdpDstPort {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowUdpDstPort) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowUdpDstPort) SetMetricGroup(value string) PatternFlowUdpDstPort {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowUdpDstPort) Increment() PatternFlowUdpDstPortCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowUdpDstPortCounter{}
	}
	return &patternFlowUdpDstPortCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowUdpDstPort) Decrement() PatternFlowUdpDstPortCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowUdpDstPortCounter{}
	}
	return &patternFlowUdpDstPortCounter{obj: obj.obj.Decrement}

}

type patternFlowUdpLength struct {
	obj *snappipb.PatternFlowUdpLength
}

func (obj *patternFlowUdpLength) msg() *snappipb.PatternFlowUdpLength {
	return obj.obj
}

func (obj *patternFlowUdpLength) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowUdpLength) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowUdpLength interface {
	msg() *snappipb.PatternFlowUdpLength
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowUdpLength
	Values() []int32
	SetValues(value []int32) PatternFlowUdpLength
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowUdpLength
	Increment() PatternFlowUdpLengthCounter
	Decrement() PatternFlowUdpLengthCounter
}

func (obj *patternFlowUdpLength) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowUdpLength) SetValue(value int32) PatternFlowUdpLength {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowUdpLength) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowUdpLength) SetValues(value []int32) PatternFlowUdpLength {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowUdpLength) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowUdpLength) SetMetricGroup(value string) PatternFlowUdpLength {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowUdpLength) Increment() PatternFlowUdpLengthCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowUdpLengthCounter{}
	}
	return &patternFlowUdpLengthCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowUdpLength) Decrement() PatternFlowUdpLengthCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowUdpLengthCounter{}
	}
	return &patternFlowUdpLengthCounter{obj: obj.obj.Decrement}

}

type patternFlowUdpChecksum struct {
	obj *snappipb.PatternFlowUdpChecksum
}

func (obj *patternFlowUdpChecksum) msg() *snappipb.PatternFlowUdpChecksum {
	return obj.obj
}

func (obj *patternFlowUdpChecksum) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowUdpChecksum) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowUdpChecksum interface {
	msg() *snappipb.PatternFlowUdpChecksum
	Yaml() string
	Json() string
	Custom() int32
	SetCustom(value int32) PatternFlowUdpChecksum
}

func (obj *patternFlowUdpChecksum) Custom() int32 {
	return *obj.obj.Custom
}

func (obj *patternFlowUdpChecksum) SetCustom(value int32) PatternFlowUdpChecksum {
	obj.obj.Custom = &value
	return obj
}

type patternFlowGreChecksumPresent struct {
	obj *snappipb.PatternFlowGreChecksumPresent
}

func (obj *patternFlowGreChecksumPresent) msg() *snappipb.PatternFlowGreChecksumPresent {
	return obj.obj
}

func (obj *patternFlowGreChecksumPresent) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGreChecksumPresent) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGreChecksumPresent interface {
	msg() *snappipb.PatternFlowGreChecksumPresent
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGreChecksumPresent
	Values() []int32
	SetValues(value []int32) PatternFlowGreChecksumPresent
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGreChecksumPresent
	Increment() PatternFlowGreChecksumPresentCounter
	Decrement() PatternFlowGreChecksumPresentCounter
}

func (obj *patternFlowGreChecksumPresent) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGreChecksumPresent) SetValue(value int32) PatternFlowGreChecksumPresent {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGreChecksumPresent) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGreChecksumPresent) SetValues(value []int32) PatternFlowGreChecksumPresent {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGreChecksumPresent) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGreChecksumPresent) SetMetricGroup(value string) PatternFlowGreChecksumPresent {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGreChecksumPresent) Increment() PatternFlowGreChecksumPresentCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGreChecksumPresentCounter{}
	}
	return &patternFlowGreChecksumPresentCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGreChecksumPresent) Decrement() PatternFlowGreChecksumPresentCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGreChecksumPresentCounter{}
	}
	return &patternFlowGreChecksumPresentCounter{obj: obj.obj.Decrement}

}

type patternFlowGreReserved0 struct {
	obj *snappipb.PatternFlowGreReserved0
}

func (obj *patternFlowGreReserved0) msg() *snappipb.PatternFlowGreReserved0 {
	return obj.obj
}

func (obj *patternFlowGreReserved0) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGreReserved0) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGreReserved0 interface {
	msg() *snappipb.PatternFlowGreReserved0
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGreReserved0
	Values() []int32
	SetValues(value []int32) PatternFlowGreReserved0
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGreReserved0
	Increment() PatternFlowGreReserved0Counter
	Decrement() PatternFlowGreReserved0Counter
}

func (obj *patternFlowGreReserved0) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGreReserved0) SetValue(value int32) PatternFlowGreReserved0 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGreReserved0) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGreReserved0) SetValues(value []int32) PatternFlowGreReserved0 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGreReserved0) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGreReserved0) SetMetricGroup(value string) PatternFlowGreReserved0 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGreReserved0) Increment() PatternFlowGreReserved0Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGreReserved0Counter{}
	}
	return &patternFlowGreReserved0Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowGreReserved0) Decrement() PatternFlowGreReserved0Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGreReserved0Counter{}
	}
	return &patternFlowGreReserved0Counter{obj: obj.obj.Decrement}

}

type patternFlowGreVersion struct {
	obj *snappipb.PatternFlowGreVersion
}

func (obj *patternFlowGreVersion) msg() *snappipb.PatternFlowGreVersion {
	return obj.obj
}

func (obj *patternFlowGreVersion) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGreVersion) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGreVersion interface {
	msg() *snappipb.PatternFlowGreVersion
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGreVersion
	Values() []int32
	SetValues(value []int32) PatternFlowGreVersion
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGreVersion
	Increment() PatternFlowGreVersionCounter
	Decrement() PatternFlowGreVersionCounter
}

func (obj *patternFlowGreVersion) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGreVersion) SetValue(value int32) PatternFlowGreVersion {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGreVersion) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGreVersion) SetValues(value []int32) PatternFlowGreVersion {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGreVersion) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGreVersion) SetMetricGroup(value string) PatternFlowGreVersion {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGreVersion) Increment() PatternFlowGreVersionCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGreVersionCounter{}
	}
	return &patternFlowGreVersionCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGreVersion) Decrement() PatternFlowGreVersionCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGreVersionCounter{}
	}
	return &patternFlowGreVersionCounter{obj: obj.obj.Decrement}

}

type patternFlowGreProtocol struct {
	obj *snappipb.PatternFlowGreProtocol
}

func (obj *patternFlowGreProtocol) msg() *snappipb.PatternFlowGreProtocol {
	return obj.obj
}

func (obj *patternFlowGreProtocol) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGreProtocol) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGreProtocol interface {
	msg() *snappipb.PatternFlowGreProtocol
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGreProtocol
	Values() []int32
	SetValues(value []int32) PatternFlowGreProtocol
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGreProtocol
	Increment() PatternFlowGreProtocolCounter
	Decrement() PatternFlowGreProtocolCounter
}

func (obj *patternFlowGreProtocol) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGreProtocol) SetValue(value int32) PatternFlowGreProtocol {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGreProtocol) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGreProtocol) SetValues(value []int32) PatternFlowGreProtocol {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGreProtocol) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGreProtocol) SetMetricGroup(value string) PatternFlowGreProtocol {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGreProtocol) Increment() PatternFlowGreProtocolCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGreProtocolCounter{}
	}
	return &patternFlowGreProtocolCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGreProtocol) Decrement() PatternFlowGreProtocolCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGreProtocolCounter{}
	}
	return &patternFlowGreProtocolCounter{obj: obj.obj.Decrement}

}

type patternFlowGreChecksum struct {
	obj *snappipb.PatternFlowGreChecksum
}

func (obj *patternFlowGreChecksum) msg() *snappipb.PatternFlowGreChecksum {
	return obj.obj
}

func (obj *patternFlowGreChecksum) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGreChecksum) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGreChecksum interface {
	msg() *snappipb.PatternFlowGreChecksum
	Yaml() string
	Json() string
	Custom() int32
	SetCustom(value int32) PatternFlowGreChecksum
}

func (obj *patternFlowGreChecksum) Custom() int32 {
	return *obj.obj.Custom
}

func (obj *patternFlowGreChecksum) SetCustom(value int32) PatternFlowGreChecksum {
	obj.obj.Custom = &value
	return obj
}

type patternFlowGreReserved1 struct {
	obj *snappipb.PatternFlowGreReserved1
}

func (obj *patternFlowGreReserved1) msg() *snappipb.PatternFlowGreReserved1 {
	return obj.obj
}

func (obj *patternFlowGreReserved1) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGreReserved1) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGreReserved1 interface {
	msg() *snappipb.PatternFlowGreReserved1
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGreReserved1
	Values() []int32
	SetValues(value []int32) PatternFlowGreReserved1
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGreReserved1
	Increment() PatternFlowGreReserved1Counter
	Decrement() PatternFlowGreReserved1Counter
}

func (obj *patternFlowGreReserved1) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGreReserved1) SetValue(value int32) PatternFlowGreReserved1 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGreReserved1) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGreReserved1) SetValues(value []int32) PatternFlowGreReserved1 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGreReserved1) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGreReserved1) SetMetricGroup(value string) PatternFlowGreReserved1 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGreReserved1) Increment() PatternFlowGreReserved1Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGreReserved1Counter{}
	}
	return &patternFlowGreReserved1Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowGreReserved1) Decrement() PatternFlowGreReserved1Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGreReserved1Counter{}
	}
	return &patternFlowGreReserved1Counter{obj: obj.obj.Decrement}

}

type patternFlowGtpv1Version struct {
	obj *snappipb.PatternFlowGtpv1Version
}

func (obj *patternFlowGtpv1Version) msg() *snappipb.PatternFlowGtpv1Version {
	return obj.obj
}

func (obj *patternFlowGtpv1Version) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1Version) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1Version interface {
	msg() *snappipb.PatternFlowGtpv1Version
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv1Version
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv1Version
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv1Version
	Increment() PatternFlowGtpv1VersionCounter
	Decrement() PatternFlowGtpv1VersionCounter
}

func (obj *patternFlowGtpv1Version) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv1Version) SetValue(value int32) PatternFlowGtpv1Version {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv1Version) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv1Version) SetValues(value []int32) PatternFlowGtpv1Version {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv1Version) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv1Version) SetMetricGroup(value string) PatternFlowGtpv1Version {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv1Version) Increment() PatternFlowGtpv1VersionCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv1VersionCounter{}
	}
	return &patternFlowGtpv1VersionCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv1Version) Decrement() PatternFlowGtpv1VersionCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv1VersionCounter{}
	}
	return &patternFlowGtpv1VersionCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv1ProtocolType struct {
	obj *snappipb.PatternFlowGtpv1ProtocolType
}

func (obj *patternFlowGtpv1ProtocolType) msg() *snappipb.PatternFlowGtpv1ProtocolType {
	return obj.obj
}

func (obj *patternFlowGtpv1ProtocolType) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1ProtocolType) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1ProtocolType interface {
	msg() *snappipb.PatternFlowGtpv1ProtocolType
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv1ProtocolType
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv1ProtocolType
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv1ProtocolType
	Increment() PatternFlowGtpv1ProtocolTypeCounter
	Decrement() PatternFlowGtpv1ProtocolTypeCounter
}

func (obj *patternFlowGtpv1ProtocolType) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv1ProtocolType) SetValue(value int32) PatternFlowGtpv1ProtocolType {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv1ProtocolType) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv1ProtocolType) SetValues(value []int32) PatternFlowGtpv1ProtocolType {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv1ProtocolType) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv1ProtocolType) SetMetricGroup(value string) PatternFlowGtpv1ProtocolType {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv1ProtocolType) Increment() PatternFlowGtpv1ProtocolTypeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv1ProtocolTypeCounter{}
	}
	return &patternFlowGtpv1ProtocolTypeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv1ProtocolType) Decrement() PatternFlowGtpv1ProtocolTypeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv1ProtocolTypeCounter{}
	}
	return &patternFlowGtpv1ProtocolTypeCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv1Reserved struct {
	obj *snappipb.PatternFlowGtpv1Reserved
}

func (obj *patternFlowGtpv1Reserved) msg() *snappipb.PatternFlowGtpv1Reserved {
	return obj.obj
}

func (obj *patternFlowGtpv1Reserved) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1Reserved) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1Reserved interface {
	msg() *snappipb.PatternFlowGtpv1Reserved
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv1Reserved
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv1Reserved
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv1Reserved
	Increment() PatternFlowGtpv1ReservedCounter
	Decrement() PatternFlowGtpv1ReservedCounter
}

func (obj *patternFlowGtpv1Reserved) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv1Reserved) SetValue(value int32) PatternFlowGtpv1Reserved {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv1Reserved) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv1Reserved) SetValues(value []int32) PatternFlowGtpv1Reserved {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv1Reserved) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv1Reserved) SetMetricGroup(value string) PatternFlowGtpv1Reserved {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv1Reserved) Increment() PatternFlowGtpv1ReservedCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv1ReservedCounter{}
	}
	return &patternFlowGtpv1ReservedCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv1Reserved) Decrement() PatternFlowGtpv1ReservedCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv1ReservedCounter{}
	}
	return &patternFlowGtpv1ReservedCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv1EFlag struct {
	obj *snappipb.PatternFlowGtpv1EFlag
}

func (obj *patternFlowGtpv1EFlag) msg() *snappipb.PatternFlowGtpv1EFlag {
	return obj.obj
}

func (obj *patternFlowGtpv1EFlag) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1EFlag) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1EFlag interface {
	msg() *snappipb.PatternFlowGtpv1EFlag
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv1EFlag
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv1EFlag
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv1EFlag
	Increment() PatternFlowGtpv1EFlagCounter
	Decrement() PatternFlowGtpv1EFlagCounter
}

func (obj *patternFlowGtpv1EFlag) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv1EFlag) SetValue(value int32) PatternFlowGtpv1EFlag {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv1EFlag) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv1EFlag) SetValues(value []int32) PatternFlowGtpv1EFlag {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv1EFlag) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv1EFlag) SetMetricGroup(value string) PatternFlowGtpv1EFlag {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv1EFlag) Increment() PatternFlowGtpv1EFlagCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv1EFlagCounter{}
	}
	return &patternFlowGtpv1EFlagCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv1EFlag) Decrement() PatternFlowGtpv1EFlagCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv1EFlagCounter{}
	}
	return &patternFlowGtpv1EFlagCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv1SFlag struct {
	obj *snappipb.PatternFlowGtpv1SFlag
}

func (obj *patternFlowGtpv1SFlag) msg() *snappipb.PatternFlowGtpv1SFlag {
	return obj.obj
}

func (obj *patternFlowGtpv1SFlag) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1SFlag) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1SFlag interface {
	msg() *snappipb.PatternFlowGtpv1SFlag
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv1SFlag
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv1SFlag
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv1SFlag
	Increment() PatternFlowGtpv1SFlagCounter
	Decrement() PatternFlowGtpv1SFlagCounter
}

func (obj *patternFlowGtpv1SFlag) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv1SFlag) SetValue(value int32) PatternFlowGtpv1SFlag {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv1SFlag) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv1SFlag) SetValues(value []int32) PatternFlowGtpv1SFlag {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv1SFlag) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv1SFlag) SetMetricGroup(value string) PatternFlowGtpv1SFlag {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv1SFlag) Increment() PatternFlowGtpv1SFlagCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv1SFlagCounter{}
	}
	return &patternFlowGtpv1SFlagCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv1SFlag) Decrement() PatternFlowGtpv1SFlagCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv1SFlagCounter{}
	}
	return &patternFlowGtpv1SFlagCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv1PnFlag struct {
	obj *snappipb.PatternFlowGtpv1PnFlag
}

func (obj *patternFlowGtpv1PnFlag) msg() *snappipb.PatternFlowGtpv1PnFlag {
	return obj.obj
}

func (obj *patternFlowGtpv1PnFlag) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1PnFlag) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1PnFlag interface {
	msg() *snappipb.PatternFlowGtpv1PnFlag
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv1PnFlag
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv1PnFlag
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv1PnFlag
	Increment() PatternFlowGtpv1PnFlagCounter
	Decrement() PatternFlowGtpv1PnFlagCounter
}

func (obj *patternFlowGtpv1PnFlag) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv1PnFlag) SetValue(value int32) PatternFlowGtpv1PnFlag {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv1PnFlag) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv1PnFlag) SetValues(value []int32) PatternFlowGtpv1PnFlag {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv1PnFlag) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv1PnFlag) SetMetricGroup(value string) PatternFlowGtpv1PnFlag {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv1PnFlag) Increment() PatternFlowGtpv1PnFlagCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv1PnFlagCounter{}
	}
	return &patternFlowGtpv1PnFlagCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv1PnFlag) Decrement() PatternFlowGtpv1PnFlagCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv1PnFlagCounter{}
	}
	return &patternFlowGtpv1PnFlagCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv1MessageType struct {
	obj *snappipb.PatternFlowGtpv1MessageType
}

func (obj *patternFlowGtpv1MessageType) msg() *snappipb.PatternFlowGtpv1MessageType {
	return obj.obj
}

func (obj *patternFlowGtpv1MessageType) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1MessageType) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1MessageType interface {
	msg() *snappipb.PatternFlowGtpv1MessageType
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv1MessageType
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv1MessageType
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv1MessageType
	Increment() PatternFlowGtpv1MessageTypeCounter
	Decrement() PatternFlowGtpv1MessageTypeCounter
}

func (obj *patternFlowGtpv1MessageType) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv1MessageType) SetValue(value int32) PatternFlowGtpv1MessageType {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv1MessageType) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv1MessageType) SetValues(value []int32) PatternFlowGtpv1MessageType {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv1MessageType) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv1MessageType) SetMetricGroup(value string) PatternFlowGtpv1MessageType {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv1MessageType) Increment() PatternFlowGtpv1MessageTypeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv1MessageTypeCounter{}
	}
	return &patternFlowGtpv1MessageTypeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv1MessageType) Decrement() PatternFlowGtpv1MessageTypeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv1MessageTypeCounter{}
	}
	return &patternFlowGtpv1MessageTypeCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv1MessageLength struct {
	obj *snappipb.PatternFlowGtpv1MessageLength
}

func (obj *patternFlowGtpv1MessageLength) msg() *snappipb.PatternFlowGtpv1MessageLength {
	return obj.obj
}

func (obj *patternFlowGtpv1MessageLength) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1MessageLength) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1MessageLength interface {
	msg() *snappipb.PatternFlowGtpv1MessageLength
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv1MessageLength
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv1MessageLength
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv1MessageLength
	Increment() PatternFlowGtpv1MessageLengthCounter
	Decrement() PatternFlowGtpv1MessageLengthCounter
}

func (obj *patternFlowGtpv1MessageLength) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv1MessageLength) SetValue(value int32) PatternFlowGtpv1MessageLength {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv1MessageLength) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv1MessageLength) SetValues(value []int32) PatternFlowGtpv1MessageLength {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv1MessageLength) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv1MessageLength) SetMetricGroup(value string) PatternFlowGtpv1MessageLength {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv1MessageLength) Increment() PatternFlowGtpv1MessageLengthCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv1MessageLengthCounter{}
	}
	return &patternFlowGtpv1MessageLengthCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv1MessageLength) Decrement() PatternFlowGtpv1MessageLengthCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv1MessageLengthCounter{}
	}
	return &patternFlowGtpv1MessageLengthCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv1Teid struct {
	obj *snappipb.PatternFlowGtpv1Teid
}

func (obj *patternFlowGtpv1Teid) msg() *snappipb.PatternFlowGtpv1Teid {
	return obj.obj
}

func (obj *patternFlowGtpv1Teid) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1Teid) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1Teid interface {
	msg() *snappipb.PatternFlowGtpv1Teid
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv1Teid
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv1Teid
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv1Teid
	Increment() PatternFlowGtpv1TeidCounter
	Decrement() PatternFlowGtpv1TeidCounter
}

func (obj *patternFlowGtpv1Teid) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv1Teid) SetValue(value int32) PatternFlowGtpv1Teid {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv1Teid) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv1Teid) SetValues(value []int32) PatternFlowGtpv1Teid {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv1Teid) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv1Teid) SetMetricGroup(value string) PatternFlowGtpv1Teid {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv1Teid) Increment() PatternFlowGtpv1TeidCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv1TeidCounter{}
	}
	return &patternFlowGtpv1TeidCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv1Teid) Decrement() PatternFlowGtpv1TeidCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv1TeidCounter{}
	}
	return &patternFlowGtpv1TeidCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv1SquenceNumber struct {
	obj *snappipb.PatternFlowGtpv1SquenceNumber
}

func (obj *patternFlowGtpv1SquenceNumber) msg() *snappipb.PatternFlowGtpv1SquenceNumber {
	return obj.obj
}

func (obj *patternFlowGtpv1SquenceNumber) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1SquenceNumber) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1SquenceNumber interface {
	msg() *snappipb.PatternFlowGtpv1SquenceNumber
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv1SquenceNumber
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv1SquenceNumber
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv1SquenceNumber
	Increment() PatternFlowGtpv1SquenceNumberCounter
	Decrement() PatternFlowGtpv1SquenceNumberCounter
}

func (obj *patternFlowGtpv1SquenceNumber) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv1SquenceNumber) SetValue(value int32) PatternFlowGtpv1SquenceNumber {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv1SquenceNumber) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv1SquenceNumber) SetValues(value []int32) PatternFlowGtpv1SquenceNumber {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv1SquenceNumber) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv1SquenceNumber) SetMetricGroup(value string) PatternFlowGtpv1SquenceNumber {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv1SquenceNumber) Increment() PatternFlowGtpv1SquenceNumberCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv1SquenceNumberCounter{}
	}
	return &patternFlowGtpv1SquenceNumberCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv1SquenceNumber) Decrement() PatternFlowGtpv1SquenceNumberCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv1SquenceNumberCounter{}
	}
	return &patternFlowGtpv1SquenceNumberCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv1NPduNumber struct {
	obj *snappipb.PatternFlowGtpv1NPduNumber
}

func (obj *patternFlowGtpv1NPduNumber) msg() *snappipb.PatternFlowGtpv1NPduNumber {
	return obj.obj
}

func (obj *patternFlowGtpv1NPduNumber) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1NPduNumber) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1NPduNumber interface {
	msg() *snappipb.PatternFlowGtpv1NPduNumber
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv1NPduNumber
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv1NPduNumber
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv1NPduNumber
	Increment() PatternFlowGtpv1NPduNumberCounter
	Decrement() PatternFlowGtpv1NPduNumberCounter
}

func (obj *patternFlowGtpv1NPduNumber) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv1NPduNumber) SetValue(value int32) PatternFlowGtpv1NPduNumber {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv1NPduNumber) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv1NPduNumber) SetValues(value []int32) PatternFlowGtpv1NPduNumber {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv1NPduNumber) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv1NPduNumber) SetMetricGroup(value string) PatternFlowGtpv1NPduNumber {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv1NPduNumber) Increment() PatternFlowGtpv1NPduNumberCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv1NPduNumberCounter{}
	}
	return &patternFlowGtpv1NPduNumberCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv1NPduNumber) Decrement() PatternFlowGtpv1NPduNumberCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv1NPduNumberCounter{}
	}
	return &patternFlowGtpv1NPduNumberCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv1NextExtensionHeaderType struct {
	obj *snappipb.PatternFlowGtpv1NextExtensionHeaderType
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) msg() *snappipb.PatternFlowGtpv1NextExtensionHeaderType {
	return obj.obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1NextExtensionHeaderType interface {
	msg() *snappipb.PatternFlowGtpv1NextExtensionHeaderType
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv1NextExtensionHeaderType
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv1NextExtensionHeaderType
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv1NextExtensionHeaderType
	Increment() PatternFlowGtpv1NextExtensionHeaderTypeCounter
	Decrement() PatternFlowGtpv1NextExtensionHeaderTypeCounter
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) SetValue(value int32) PatternFlowGtpv1NextExtensionHeaderType {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) SetValues(value []int32) PatternFlowGtpv1NextExtensionHeaderType {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) SetMetricGroup(value string) PatternFlowGtpv1NextExtensionHeaderType {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) Increment() PatternFlowGtpv1NextExtensionHeaderTypeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv1NextExtensionHeaderTypeCounter{}
	}
	return &patternFlowGtpv1NextExtensionHeaderTypeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv1NextExtensionHeaderType) Decrement() PatternFlowGtpv1NextExtensionHeaderTypeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv1NextExtensionHeaderTypeCounter{}
	}
	return &patternFlowGtpv1NextExtensionHeaderTypeCounter{obj: obj.obj.Decrement}

}

type flowGtpExtension struct {
	obj *snappipb.FlowGtpExtension
}

func (obj *flowGtpExtension) msg() *snappipb.FlowGtpExtension {
	return obj.obj
}

func (obj *flowGtpExtension) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowGtpExtension) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowGtpExtension interface {
	msg() *snappipb.FlowGtpExtension
	Yaml() string
	Json() string
	ExtensionLength() PatternFlowGtpExtensionExtensionLength
	Contents() PatternFlowGtpExtensionContents
	NextExtensionHeader() PatternFlowGtpExtensionNextExtensionHeader
}

func (obj *flowGtpExtension) ExtensionLength() PatternFlowGtpExtensionExtensionLength {
	if obj.obj.ExtensionLength == nil {
		obj.obj.ExtensionLength = &snappipb.PatternFlowGtpExtensionExtensionLength{}
	}
	return &patternFlowGtpExtensionExtensionLength{obj: obj.obj.ExtensionLength}

}

func (obj *flowGtpExtension) Contents() PatternFlowGtpExtensionContents {
	if obj.obj.Contents == nil {
		obj.obj.Contents = &snappipb.PatternFlowGtpExtensionContents{}
	}
	return &patternFlowGtpExtensionContents{obj: obj.obj.Contents}

}

func (obj *flowGtpExtension) NextExtensionHeader() PatternFlowGtpExtensionNextExtensionHeader {
	if obj.obj.NextExtensionHeader == nil {
		obj.obj.NextExtensionHeader = &snappipb.PatternFlowGtpExtensionNextExtensionHeader{}
	}
	return &patternFlowGtpExtensionNextExtensionHeader{obj: obj.obj.NextExtensionHeader}

}

type patternFlowGtpv2Version struct {
	obj *snappipb.PatternFlowGtpv2Version
}

func (obj *patternFlowGtpv2Version) msg() *snappipb.PatternFlowGtpv2Version {
	return obj.obj
}

func (obj *patternFlowGtpv2Version) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2Version) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2Version interface {
	msg() *snappipb.PatternFlowGtpv2Version
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv2Version
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv2Version
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv2Version
	Increment() PatternFlowGtpv2VersionCounter
	Decrement() PatternFlowGtpv2VersionCounter
}

func (obj *patternFlowGtpv2Version) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv2Version) SetValue(value int32) PatternFlowGtpv2Version {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv2Version) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv2Version) SetValues(value []int32) PatternFlowGtpv2Version {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv2Version) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv2Version) SetMetricGroup(value string) PatternFlowGtpv2Version {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv2Version) Increment() PatternFlowGtpv2VersionCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv2VersionCounter{}
	}
	return &patternFlowGtpv2VersionCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv2Version) Decrement() PatternFlowGtpv2VersionCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv2VersionCounter{}
	}
	return &patternFlowGtpv2VersionCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv2PiggybackingFlag struct {
	obj *snappipb.PatternFlowGtpv2PiggybackingFlag
}

func (obj *patternFlowGtpv2PiggybackingFlag) msg() *snappipb.PatternFlowGtpv2PiggybackingFlag {
	return obj.obj
}

func (obj *patternFlowGtpv2PiggybackingFlag) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2PiggybackingFlag) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2PiggybackingFlag interface {
	msg() *snappipb.PatternFlowGtpv2PiggybackingFlag
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv2PiggybackingFlag
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv2PiggybackingFlag
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv2PiggybackingFlag
	Increment() PatternFlowGtpv2PiggybackingFlagCounter
	Decrement() PatternFlowGtpv2PiggybackingFlagCounter
}

func (obj *patternFlowGtpv2PiggybackingFlag) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv2PiggybackingFlag) SetValue(value int32) PatternFlowGtpv2PiggybackingFlag {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv2PiggybackingFlag) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv2PiggybackingFlag) SetValues(value []int32) PatternFlowGtpv2PiggybackingFlag {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv2PiggybackingFlag) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv2PiggybackingFlag) SetMetricGroup(value string) PatternFlowGtpv2PiggybackingFlag {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv2PiggybackingFlag) Increment() PatternFlowGtpv2PiggybackingFlagCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv2PiggybackingFlagCounter{}
	}
	return &patternFlowGtpv2PiggybackingFlagCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv2PiggybackingFlag) Decrement() PatternFlowGtpv2PiggybackingFlagCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv2PiggybackingFlagCounter{}
	}
	return &patternFlowGtpv2PiggybackingFlagCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv2TeidFlag struct {
	obj *snappipb.PatternFlowGtpv2TeidFlag
}

func (obj *patternFlowGtpv2TeidFlag) msg() *snappipb.PatternFlowGtpv2TeidFlag {
	return obj.obj
}

func (obj *patternFlowGtpv2TeidFlag) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2TeidFlag) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2TeidFlag interface {
	msg() *snappipb.PatternFlowGtpv2TeidFlag
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv2TeidFlag
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv2TeidFlag
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv2TeidFlag
	Increment() PatternFlowGtpv2TeidFlagCounter
	Decrement() PatternFlowGtpv2TeidFlagCounter
}

func (obj *patternFlowGtpv2TeidFlag) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv2TeidFlag) SetValue(value int32) PatternFlowGtpv2TeidFlag {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv2TeidFlag) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv2TeidFlag) SetValues(value []int32) PatternFlowGtpv2TeidFlag {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv2TeidFlag) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv2TeidFlag) SetMetricGroup(value string) PatternFlowGtpv2TeidFlag {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv2TeidFlag) Increment() PatternFlowGtpv2TeidFlagCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv2TeidFlagCounter{}
	}
	return &patternFlowGtpv2TeidFlagCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv2TeidFlag) Decrement() PatternFlowGtpv2TeidFlagCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv2TeidFlagCounter{}
	}
	return &patternFlowGtpv2TeidFlagCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv2Spare1 struct {
	obj *snappipb.PatternFlowGtpv2Spare1
}

func (obj *patternFlowGtpv2Spare1) msg() *snappipb.PatternFlowGtpv2Spare1 {
	return obj.obj
}

func (obj *patternFlowGtpv2Spare1) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2Spare1) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2Spare1 interface {
	msg() *snappipb.PatternFlowGtpv2Spare1
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv2Spare1
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv2Spare1
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv2Spare1
	Increment() PatternFlowGtpv2Spare1Counter
	Decrement() PatternFlowGtpv2Spare1Counter
}

func (obj *patternFlowGtpv2Spare1) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv2Spare1) SetValue(value int32) PatternFlowGtpv2Spare1 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv2Spare1) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv2Spare1) SetValues(value []int32) PatternFlowGtpv2Spare1 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv2Spare1) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv2Spare1) SetMetricGroup(value string) PatternFlowGtpv2Spare1 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv2Spare1) Increment() PatternFlowGtpv2Spare1Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv2Spare1Counter{}
	}
	return &patternFlowGtpv2Spare1Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv2Spare1) Decrement() PatternFlowGtpv2Spare1Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv2Spare1Counter{}
	}
	return &patternFlowGtpv2Spare1Counter{obj: obj.obj.Decrement}

}

type patternFlowGtpv2MessageType struct {
	obj *snappipb.PatternFlowGtpv2MessageType
}

func (obj *patternFlowGtpv2MessageType) msg() *snappipb.PatternFlowGtpv2MessageType {
	return obj.obj
}

func (obj *patternFlowGtpv2MessageType) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2MessageType) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2MessageType interface {
	msg() *snappipb.PatternFlowGtpv2MessageType
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv2MessageType
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv2MessageType
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv2MessageType
	Increment() PatternFlowGtpv2MessageTypeCounter
	Decrement() PatternFlowGtpv2MessageTypeCounter
}

func (obj *patternFlowGtpv2MessageType) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv2MessageType) SetValue(value int32) PatternFlowGtpv2MessageType {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv2MessageType) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv2MessageType) SetValues(value []int32) PatternFlowGtpv2MessageType {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv2MessageType) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv2MessageType) SetMetricGroup(value string) PatternFlowGtpv2MessageType {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv2MessageType) Increment() PatternFlowGtpv2MessageTypeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv2MessageTypeCounter{}
	}
	return &patternFlowGtpv2MessageTypeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv2MessageType) Decrement() PatternFlowGtpv2MessageTypeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv2MessageTypeCounter{}
	}
	return &patternFlowGtpv2MessageTypeCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv2MessageLength struct {
	obj *snappipb.PatternFlowGtpv2MessageLength
}

func (obj *patternFlowGtpv2MessageLength) msg() *snappipb.PatternFlowGtpv2MessageLength {
	return obj.obj
}

func (obj *patternFlowGtpv2MessageLength) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2MessageLength) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2MessageLength interface {
	msg() *snappipb.PatternFlowGtpv2MessageLength
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv2MessageLength
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv2MessageLength
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv2MessageLength
	Increment() PatternFlowGtpv2MessageLengthCounter
	Decrement() PatternFlowGtpv2MessageLengthCounter
}

func (obj *patternFlowGtpv2MessageLength) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv2MessageLength) SetValue(value int32) PatternFlowGtpv2MessageLength {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv2MessageLength) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv2MessageLength) SetValues(value []int32) PatternFlowGtpv2MessageLength {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv2MessageLength) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv2MessageLength) SetMetricGroup(value string) PatternFlowGtpv2MessageLength {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv2MessageLength) Increment() PatternFlowGtpv2MessageLengthCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv2MessageLengthCounter{}
	}
	return &patternFlowGtpv2MessageLengthCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv2MessageLength) Decrement() PatternFlowGtpv2MessageLengthCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv2MessageLengthCounter{}
	}
	return &patternFlowGtpv2MessageLengthCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv2Teid struct {
	obj *snappipb.PatternFlowGtpv2Teid
}

func (obj *patternFlowGtpv2Teid) msg() *snappipb.PatternFlowGtpv2Teid {
	return obj.obj
}

func (obj *patternFlowGtpv2Teid) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2Teid) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2Teid interface {
	msg() *snappipb.PatternFlowGtpv2Teid
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv2Teid
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv2Teid
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv2Teid
	Increment() PatternFlowGtpv2TeidCounter
	Decrement() PatternFlowGtpv2TeidCounter
}

func (obj *patternFlowGtpv2Teid) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv2Teid) SetValue(value int32) PatternFlowGtpv2Teid {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv2Teid) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv2Teid) SetValues(value []int32) PatternFlowGtpv2Teid {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv2Teid) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv2Teid) SetMetricGroup(value string) PatternFlowGtpv2Teid {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv2Teid) Increment() PatternFlowGtpv2TeidCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv2TeidCounter{}
	}
	return &patternFlowGtpv2TeidCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv2Teid) Decrement() PatternFlowGtpv2TeidCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv2TeidCounter{}
	}
	return &patternFlowGtpv2TeidCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv2SequenceNumber struct {
	obj *snappipb.PatternFlowGtpv2SequenceNumber
}

func (obj *patternFlowGtpv2SequenceNumber) msg() *snappipb.PatternFlowGtpv2SequenceNumber {
	return obj.obj
}

func (obj *patternFlowGtpv2SequenceNumber) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2SequenceNumber) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2SequenceNumber interface {
	msg() *snappipb.PatternFlowGtpv2SequenceNumber
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv2SequenceNumber
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv2SequenceNumber
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv2SequenceNumber
	Increment() PatternFlowGtpv2SequenceNumberCounter
	Decrement() PatternFlowGtpv2SequenceNumberCounter
}

func (obj *patternFlowGtpv2SequenceNumber) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv2SequenceNumber) SetValue(value int32) PatternFlowGtpv2SequenceNumber {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv2SequenceNumber) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv2SequenceNumber) SetValues(value []int32) PatternFlowGtpv2SequenceNumber {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv2SequenceNumber) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv2SequenceNumber) SetMetricGroup(value string) PatternFlowGtpv2SequenceNumber {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv2SequenceNumber) Increment() PatternFlowGtpv2SequenceNumberCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv2SequenceNumberCounter{}
	}
	return &patternFlowGtpv2SequenceNumberCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv2SequenceNumber) Decrement() PatternFlowGtpv2SequenceNumberCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv2SequenceNumberCounter{}
	}
	return &patternFlowGtpv2SequenceNumberCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv2Spare2 struct {
	obj *snappipb.PatternFlowGtpv2Spare2
}

func (obj *patternFlowGtpv2Spare2) msg() *snappipb.PatternFlowGtpv2Spare2 {
	return obj.obj
}

func (obj *patternFlowGtpv2Spare2) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2Spare2) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2Spare2 interface {
	msg() *snappipb.PatternFlowGtpv2Spare2
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpv2Spare2
	Values() []int32
	SetValues(value []int32) PatternFlowGtpv2Spare2
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpv2Spare2
	Increment() PatternFlowGtpv2Spare2Counter
	Decrement() PatternFlowGtpv2Spare2Counter
}

func (obj *patternFlowGtpv2Spare2) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpv2Spare2) SetValue(value int32) PatternFlowGtpv2Spare2 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpv2Spare2) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpv2Spare2) SetValues(value []int32) PatternFlowGtpv2Spare2 {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpv2Spare2) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpv2Spare2) SetMetricGroup(value string) PatternFlowGtpv2Spare2 {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpv2Spare2) Increment() PatternFlowGtpv2Spare2Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpv2Spare2Counter{}
	}
	return &patternFlowGtpv2Spare2Counter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpv2Spare2) Decrement() PatternFlowGtpv2Spare2Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpv2Spare2Counter{}
	}
	return &patternFlowGtpv2Spare2Counter{obj: obj.obj.Decrement}

}

type patternFlowArpHardwareType struct {
	obj *snappipb.PatternFlowArpHardwareType
}

func (obj *patternFlowArpHardwareType) msg() *snappipb.PatternFlowArpHardwareType {
	return obj.obj
}

func (obj *patternFlowArpHardwareType) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpHardwareType) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpHardwareType interface {
	msg() *snappipb.PatternFlowArpHardwareType
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowArpHardwareType
	Values() []int32
	SetValues(value []int32) PatternFlowArpHardwareType
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowArpHardwareType
	Increment() PatternFlowArpHardwareTypeCounter
	Decrement() PatternFlowArpHardwareTypeCounter
}

func (obj *patternFlowArpHardwareType) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowArpHardwareType) SetValue(value int32) PatternFlowArpHardwareType {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowArpHardwareType) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowArpHardwareType) SetValues(value []int32) PatternFlowArpHardwareType {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowArpHardwareType) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowArpHardwareType) SetMetricGroup(value string) PatternFlowArpHardwareType {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowArpHardwareType) Increment() PatternFlowArpHardwareTypeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowArpHardwareTypeCounter{}
	}
	return &patternFlowArpHardwareTypeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowArpHardwareType) Decrement() PatternFlowArpHardwareTypeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowArpHardwareTypeCounter{}
	}
	return &patternFlowArpHardwareTypeCounter{obj: obj.obj.Decrement}

}

type patternFlowArpProtocolType struct {
	obj *snappipb.PatternFlowArpProtocolType
}

func (obj *patternFlowArpProtocolType) msg() *snappipb.PatternFlowArpProtocolType {
	return obj.obj
}

func (obj *patternFlowArpProtocolType) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpProtocolType) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpProtocolType interface {
	msg() *snappipb.PatternFlowArpProtocolType
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowArpProtocolType
	Values() []int32
	SetValues(value []int32) PatternFlowArpProtocolType
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowArpProtocolType
	Increment() PatternFlowArpProtocolTypeCounter
	Decrement() PatternFlowArpProtocolTypeCounter
}

func (obj *patternFlowArpProtocolType) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowArpProtocolType) SetValue(value int32) PatternFlowArpProtocolType {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowArpProtocolType) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowArpProtocolType) SetValues(value []int32) PatternFlowArpProtocolType {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowArpProtocolType) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowArpProtocolType) SetMetricGroup(value string) PatternFlowArpProtocolType {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowArpProtocolType) Increment() PatternFlowArpProtocolTypeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowArpProtocolTypeCounter{}
	}
	return &patternFlowArpProtocolTypeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowArpProtocolType) Decrement() PatternFlowArpProtocolTypeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowArpProtocolTypeCounter{}
	}
	return &patternFlowArpProtocolTypeCounter{obj: obj.obj.Decrement}

}

type patternFlowArpHardwareLength struct {
	obj *snappipb.PatternFlowArpHardwareLength
}

func (obj *patternFlowArpHardwareLength) msg() *snappipb.PatternFlowArpHardwareLength {
	return obj.obj
}

func (obj *patternFlowArpHardwareLength) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpHardwareLength) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpHardwareLength interface {
	msg() *snappipb.PatternFlowArpHardwareLength
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowArpHardwareLength
	Values() []int32
	SetValues(value []int32) PatternFlowArpHardwareLength
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowArpHardwareLength
	Increment() PatternFlowArpHardwareLengthCounter
	Decrement() PatternFlowArpHardwareLengthCounter
}

func (obj *patternFlowArpHardwareLength) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowArpHardwareLength) SetValue(value int32) PatternFlowArpHardwareLength {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowArpHardwareLength) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowArpHardwareLength) SetValues(value []int32) PatternFlowArpHardwareLength {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowArpHardwareLength) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowArpHardwareLength) SetMetricGroup(value string) PatternFlowArpHardwareLength {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowArpHardwareLength) Increment() PatternFlowArpHardwareLengthCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowArpHardwareLengthCounter{}
	}
	return &patternFlowArpHardwareLengthCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowArpHardwareLength) Decrement() PatternFlowArpHardwareLengthCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowArpHardwareLengthCounter{}
	}
	return &patternFlowArpHardwareLengthCounter{obj: obj.obj.Decrement}

}

type patternFlowArpProtocolLength struct {
	obj *snappipb.PatternFlowArpProtocolLength
}

func (obj *patternFlowArpProtocolLength) msg() *snappipb.PatternFlowArpProtocolLength {
	return obj.obj
}

func (obj *patternFlowArpProtocolLength) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpProtocolLength) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpProtocolLength interface {
	msg() *snappipb.PatternFlowArpProtocolLength
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowArpProtocolLength
	Values() []int32
	SetValues(value []int32) PatternFlowArpProtocolLength
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowArpProtocolLength
	Increment() PatternFlowArpProtocolLengthCounter
	Decrement() PatternFlowArpProtocolLengthCounter
}

func (obj *patternFlowArpProtocolLength) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowArpProtocolLength) SetValue(value int32) PatternFlowArpProtocolLength {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowArpProtocolLength) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowArpProtocolLength) SetValues(value []int32) PatternFlowArpProtocolLength {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowArpProtocolLength) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowArpProtocolLength) SetMetricGroup(value string) PatternFlowArpProtocolLength {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowArpProtocolLength) Increment() PatternFlowArpProtocolLengthCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowArpProtocolLengthCounter{}
	}
	return &patternFlowArpProtocolLengthCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowArpProtocolLength) Decrement() PatternFlowArpProtocolLengthCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowArpProtocolLengthCounter{}
	}
	return &patternFlowArpProtocolLengthCounter{obj: obj.obj.Decrement}

}

type patternFlowArpOperation struct {
	obj *snappipb.PatternFlowArpOperation
}

func (obj *patternFlowArpOperation) msg() *snappipb.PatternFlowArpOperation {
	return obj.obj
}

func (obj *patternFlowArpOperation) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpOperation) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpOperation interface {
	msg() *snappipb.PatternFlowArpOperation
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowArpOperation
	Values() []int32
	SetValues(value []int32) PatternFlowArpOperation
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowArpOperation
	Increment() PatternFlowArpOperationCounter
	Decrement() PatternFlowArpOperationCounter
}

func (obj *patternFlowArpOperation) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowArpOperation) SetValue(value int32) PatternFlowArpOperation {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowArpOperation) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowArpOperation) SetValues(value []int32) PatternFlowArpOperation {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowArpOperation) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowArpOperation) SetMetricGroup(value string) PatternFlowArpOperation {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowArpOperation) Increment() PatternFlowArpOperationCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowArpOperationCounter{}
	}
	return &patternFlowArpOperationCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowArpOperation) Decrement() PatternFlowArpOperationCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowArpOperationCounter{}
	}
	return &patternFlowArpOperationCounter{obj: obj.obj.Decrement}

}

type patternFlowArpSenderHardwareAddr struct {
	obj *snappipb.PatternFlowArpSenderHardwareAddr
}

func (obj *patternFlowArpSenderHardwareAddr) msg() *snappipb.PatternFlowArpSenderHardwareAddr {
	return obj.obj
}

func (obj *patternFlowArpSenderHardwareAddr) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpSenderHardwareAddr) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpSenderHardwareAddr interface {
	msg() *snappipb.PatternFlowArpSenderHardwareAddr
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowArpSenderHardwareAddr
	Values() []string
	SetValues(value []string) PatternFlowArpSenderHardwareAddr
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowArpSenderHardwareAddr
	Increment() PatternFlowArpSenderHardwareAddrCounter
	Decrement() PatternFlowArpSenderHardwareAddrCounter
}

func (obj *patternFlowArpSenderHardwareAddr) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowArpSenderHardwareAddr) SetValue(value string) PatternFlowArpSenderHardwareAddr {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowArpSenderHardwareAddr) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowArpSenderHardwareAddr) SetValues(value []string) PatternFlowArpSenderHardwareAddr {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowArpSenderHardwareAddr) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowArpSenderHardwareAddr) SetMetricGroup(value string) PatternFlowArpSenderHardwareAddr {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowArpSenderHardwareAddr) Increment() PatternFlowArpSenderHardwareAddrCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowArpSenderHardwareAddrCounter{}
	}
	return &patternFlowArpSenderHardwareAddrCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowArpSenderHardwareAddr) Decrement() PatternFlowArpSenderHardwareAddrCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowArpSenderHardwareAddrCounter{}
	}
	return &patternFlowArpSenderHardwareAddrCounter{obj: obj.obj.Decrement}

}

type patternFlowArpSenderProtocolAddr struct {
	obj *snappipb.PatternFlowArpSenderProtocolAddr
}

func (obj *patternFlowArpSenderProtocolAddr) msg() *snappipb.PatternFlowArpSenderProtocolAddr {
	return obj.obj
}

func (obj *patternFlowArpSenderProtocolAddr) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpSenderProtocolAddr) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpSenderProtocolAddr interface {
	msg() *snappipb.PatternFlowArpSenderProtocolAddr
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowArpSenderProtocolAddr
	Values() []string
	SetValues(value []string) PatternFlowArpSenderProtocolAddr
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowArpSenderProtocolAddr
	Increment() PatternFlowArpSenderProtocolAddrCounter
	Decrement() PatternFlowArpSenderProtocolAddrCounter
}

func (obj *patternFlowArpSenderProtocolAddr) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowArpSenderProtocolAddr) SetValue(value string) PatternFlowArpSenderProtocolAddr {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowArpSenderProtocolAddr) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowArpSenderProtocolAddr) SetValues(value []string) PatternFlowArpSenderProtocolAddr {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowArpSenderProtocolAddr) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowArpSenderProtocolAddr) SetMetricGroup(value string) PatternFlowArpSenderProtocolAddr {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowArpSenderProtocolAddr) Increment() PatternFlowArpSenderProtocolAddrCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowArpSenderProtocolAddrCounter{}
	}
	return &patternFlowArpSenderProtocolAddrCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowArpSenderProtocolAddr) Decrement() PatternFlowArpSenderProtocolAddrCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowArpSenderProtocolAddrCounter{}
	}
	return &patternFlowArpSenderProtocolAddrCounter{obj: obj.obj.Decrement}

}

type patternFlowArpTargetHardwareAddr struct {
	obj *snappipb.PatternFlowArpTargetHardwareAddr
}

func (obj *patternFlowArpTargetHardwareAddr) msg() *snappipb.PatternFlowArpTargetHardwareAddr {
	return obj.obj
}

func (obj *patternFlowArpTargetHardwareAddr) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpTargetHardwareAddr) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpTargetHardwareAddr interface {
	msg() *snappipb.PatternFlowArpTargetHardwareAddr
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowArpTargetHardwareAddr
	Values() []string
	SetValues(value []string) PatternFlowArpTargetHardwareAddr
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowArpTargetHardwareAddr
	Increment() PatternFlowArpTargetHardwareAddrCounter
	Decrement() PatternFlowArpTargetHardwareAddrCounter
}

func (obj *patternFlowArpTargetHardwareAddr) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowArpTargetHardwareAddr) SetValue(value string) PatternFlowArpTargetHardwareAddr {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowArpTargetHardwareAddr) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowArpTargetHardwareAddr) SetValues(value []string) PatternFlowArpTargetHardwareAddr {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowArpTargetHardwareAddr) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowArpTargetHardwareAddr) SetMetricGroup(value string) PatternFlowArpTargetHardwareAddr {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowArpTargetHardwareAddr) Increment() PatternFlowArpTargetHardwareAddrCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowArpTargetHardwareAddrCounter{}
	}
	return &patternFlowArpTargetHardwareAddrCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowArpTargetHardwareAddr) Decrement() PatternFlowArpTargetHardwareAddrCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowArpTargetHardwareAddrCounter{}
	}
	return &patternFlowArpTargetHardwareAddrCounter{obj: obj.obj.Decrement}

}

type patternFlowArpTargetProtocolAddr struct {
	obj *snappipb.PatternFlowArpTargetProtocolAddr
}

func (obj *patternFlowArpTargetProtocolAddr) msg() *snappipb.PatternFlowArpTargetProtocolAddr {
	return obj.obj
}

func (obj *patternFlowArpTargetProtocolAddr) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpTargetProtocolAddr) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpTargetProtocolAddr interface {
	msg() *snappipb.PatternFlowArpTargetProtocolAddr
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowArpTargetProtocolAddr
	Values() []string
	SetValues(value []string) PatternFlowArpTargetProtocolAddr
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowArpTargetProtocolAddr
	Increment() PatternFlowArpTargetProtocolAddrCounter
	Decrement() PatternFlowArpTargetProtocolAddrCounter
}

func (obj *patternFlowArpTargetProtocolAddr) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowArpTargetProtocolAddr) SetValue(value string) PatternFlowArpTargetProtocolAddr {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowArpTargetProtocolAddr) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowArpTargetProtocolAddr) SetValues(value []string) PatternFlowArpTargetProtocolAddr {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowArpTargetProtocolAddr) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowArpTargetProtocolAddr) SetMetricGroup(value string) PatternFlowArpTargetProtocolAddr {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowArpTargetProtocolAddr) Increment() PatternFlowArpTargetProtocolAddrCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowArpTargetProtocolAddrCounter{}
	}
	return &patternFlowArpTargetProtocolAddrCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowArpTargetProtocolAddr) Decrement() PatternFlowArpTargetProtocolAddrCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowArpTargetProtocolAddrCounter{}
	}
	return &patternFlowArpTargetProtocolAddrCounter{obj: obj.obj.Decrement}

}

type flowIcmpEcho struct {
	obj *snappipb.FlowIcmpEcho
}

func (obj *flowIcmpEcho) msg() *snappipb.FlowIcmpEcho {
	return obj.obj
}

func (obj *flowIcmpEcho) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowIcmpEcho) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowIcmpEcho interface {
	msg() *snappipb.FlowIcmpEcho
	Yaml() string
	Json() string
	Type() PatternFlowIcmpEchoType
	Code() PatternFlowIcmpEchoCode
	Checksum() PatternFlowIcmpEchoChecksum
	Identifier() PatternFlowIcmpEchoIdentifier
	SequenceNumber() PatternFlowIcmpEchoSequenceNumber
}

func (obj *flowIcmpEcho) Type() PatternFlowIcmpEchoType {
	if obj.obj.Type == nil {
		obj.obj.Type = &snappipb.PatternFlowIcmpEchoType{}
	}
	return &patternFlowIcmpEchoType{obj: obj.obj.Type}

}

func (obj *flowIcmpEcho) Code() PatternFlowIcmpEchoCode {
	if obj.obj.Code == nil {
		obj.obj.Code = &snappipb.PatternFlowIcmpEchoCode{}
	}
	return &patternFlowIcmpEchoCode{obj: obj.obj.Code}

}

func (obj *flowIcmpEcho) Checksum() PatternFlowIcmpEchoChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = &snappipb.PatternFlowIcmpEchoChecksum{}
	}
	return &patternFlowIcmpEchoChecksum{obj: obj.obj.Checksum}

}

func (obj *flowIcmpEcho) Identifier() PatternFlowIcmpEchoIdentifier {
	if obj.obj.Identifier == nil {
		obj.obj.Identifier = &snappipb.PatternFlowIcmpEchoIdentifier{}
	}
	return &patternFlowIcmpEchoIdentifier{obj: obj.obj.Identifier}

}

func (obj *flowIcmpEcho) SequenceNumber() PatternFlowIcmpEchoSequenceNumber {
	if obj.obj.SequenceNumber == nil {
		obj.obj.SequenceNumber = &snappipb.PatternFlowIcmpEchoSequenceNumber{}
	}
	return &patternFlowIcmpEchoSequenceNumber{obj: obj.obj.SequenceNumber}

}

type flowIcmpv6Echo struct {
	obj *snappipb.FlowIcmpv6Echo
}

func (obj *flowIcmpv6Echo) msg() *snappipb.FlowIcmpv6Echo {
	return obj.obj
}

func (obj *flowIcmpv6Echo) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowIcmpv6Echo) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowIcmpv6Echo interface {
	msg() *snappipb.FlowIcmpv6Echo
	Yaml() string
	Json() string
	Type() PatternFlowIcmpv6EchoType
	Code() PatternFlowIcmpv6EchoCode
	Identifier() PatternFlowIcmpv6EchoIdentifier
	SequenceNumber() PatternFlowIcmpv6EchoSequenceNumber
	Checksum() PatternFlowIcmpv6EchoChecksum
}

func (obj *flowIcmpv6Echo) Type() PatternFlowIcmpv6EchoType {
	if obj.obj.Type == nil {
		obj.obj.Type = &snappipb.PatternFlowIcmpv6EchoType{}
	}
	return &patternFlowIcmpv6EchoType{obj: obj.obj.Type}

}

func (obj *flowIcmpv6Echo) Code() PatternFlowIcmpv6EchoCode {
	if obj.obj.Code == nil {
		obj.obj.Code = &snappipb.PatternFlowIcmpv6EchoCode{}
	}
	return &patternFlowIcmpv6EchoCode{obj: obj.obj.Code}

}

func (obj *flowIcmpv6Echo) Identifier() PatternFlowIcmpv6EchoIdentifier {
	if obj.obj.Identifier == nil {
		obj.obj.Identifier = &snappipb.PatternFlowIcmpv6EchoIdentifier{}
	}
	return &patternFlowIcmpv6EchoIdentifier{obj: obj.obj.Identifier}

}

func (obj *flowIcmpv6Echo) SequenceNumber() PatternFlowIcmpv6EchoSequenceNumber {
	if obj.obj.SequenceNumber == nil {
		obj.obj.SequenceNumber = &snappipb.PatternFlowIcmpv6EchoSequenceNumber{}
	}
	return &patternFlowIcmpv6EchoSequenceNumber{obj: obj.obj.SequenceNumber}

}

func (obj *flowIcmpv6Echo) Checksum() PatternFlowIcmpv6EchoChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = &snappipb.PatternFlowIcmpv6EchoChecksum{}
	}
	return &patternFlowIcmpv6EchoChecksum{obj: obj.obj.Checksum}

}

type patternFlowPppAddress struct {
	obj *snappipb.PatternFlowPppAddress
}

func (obj *patternFlowPppAddress) msg() *snappipb.PatternFlowPppAddress {
	return obj.obj
}

func (obj *patternFlowPppAddress) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPppAddress) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPppAddress interface {
	msg() *snappipb.PatternFlowPppAddress
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPppAddress
	Values() []int32
	SetValues(value []int32) PatternFlowPppAddress
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPppAddress
	Increment() PatternFlowPppAddressCounter
	Decrement() PatternFlowPppAddressCounter
}

func (obj *patternFlowPppAddress) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPppAddress) SetValue(value int32) PatternFlowPppAddress {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPppAddress) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPppAddress) SetValues(value []int32) PatternFlowPppAddress {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPppAddress) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPppAddress) SetMetricGroup(value string) PatternFlowPppAddress {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPppAddress) Increment() PatternFlowPppAddressCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPppAddressCounter{}
	}
	return &patternFlowPppAddressCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowPppAddress) Decrement() PatternFlowPppAddressCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPppAddressCounter{}
	}
	return &patternFlowPppAddressCounter{obj: obj.obj.Decrement}

}

type patternFlowPppControl struct {
	obj *snappipb.PatternFlowPppControl
}

func (obj *patternFlowPppControl) msg() *snappipb.PatternFlowPppControl {
	return obj.obj
}

func (obj *patternFlowPppControl) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPppControl) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPppControl interface {
	msg() *snappipb.PatternFlowPppControl
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPppControl
	Values() []int32
	SetValues(value []int32) PatternFlowPppControl
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPppControl
	Increment() PatternFlowPppControlCounter
	Decrement() PatternFlowPppControlCounter
}

func (obj *patternFlowPppControl) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPppControl) SetValue(value int32) PatternFlowPppControl {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPppControl) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPppControl) SetValues(value []int32) PatternFlowPppControl {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPppControl) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPppControl) SetMetricGroup(value string) PatternFlowPppControl {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPppControl) Increment() PatternFlowPppControlCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPppControlCounter{}
	}
	return &patternFlowPppControlCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowPppControl) Decrement() PatternFlowPppControlCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPppControlCounter{}
	}
	return &patternFlowPppControlCounter{obj: obj.obj.Decrement}

}

type patternFlowPppProtocolType struct {
	obj *snappipb.PatternFlowPppProtocolType
}

func (obj *patternFlowPppProtocolType) msg() *snappipb.PatternFlowPppProtocolType {
	return obj.obj
}

func (obj *patternFlowPppProtocolType) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPppProtocolType) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPppProtocolType interface {
	msg() *snappipb.PatternFlowPppProtocolType
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowPppProtocolType
	Values() []int32
	SetValues(value []int32) PatternFlowPppProtocolType
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowPppProtocolType
	Increment() PatternFlowPppProtocolTypeCounter
	Decrement() PatternFlowPppProtocolTypeCounter
}

func (obj *patternFlowPppProtocolType) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowPppProtocolType) SetValue(value int32) PatternFlowPppProtocolType {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowPppProtocolType) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowPppProtocolType) SetValues(value []int32) PatternFlowPppProtocolType {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowPppProtocolType) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowPppProtocolType) SetMetricGroup(value string) PatternFlowPppProtocolType {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowPppProtocolType) Increment() PatternFlowPppProtocolTypeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowPppProtocolTypeCounter{}
	}
	return &patternFlowPppProtocolTypeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowPppProtocolType) Decrement() PatternFlowPppProtocolTypeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowPppProtocolTypeCounter{}
	}
	return &patternFlowPppProtocolTypeCounter{obj: obj.obj.Decrement}

}

type patternFlowIgmpv1Version struct {
	obj *snappipb.PatternFlowIgmpv1Version
}

func (obj *patternFlowIgmpv1Version) msg() *snappipb.PatternFlowIgmpv1Version {
	return obj.obj
}

func (obj *patternFlowIgmpv1Version) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIgmpv1Version) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIgmpv1Version interface {
	msg() *snappipb.PatternFlowIgmpv1Version
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIgmpv1Version
	Values() []int32
	SetValues(value []int32) PatternFlowIgmpv1Version
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIgmpv1Version
	Increment() PatternFlowIgmpv1VersionCounter
	Decrement() PatternFlowIgmpv1VersionCounter
}

func (obj *patternFlowIgmpv1Version) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIgmpv1Version) SetValue(value int32) PatternFlowIgmpv1Version {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIgmpv1Version) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIgmpv1Version) SetValues(value []int32) PatternFlowIgmpv1Version {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIgmpv1Version) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIgmpv1Version) SetMetricGroup(value string) PatternFlowIgmpv1Version {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIgmpv1Version) Increment() PatternFlowIgmpv1VersionCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIgmpv1VersionCounter{}
	}
	return &patternFlowIgmpv1VersionCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIgmpv1Version) Decrement() PatternFlowIgmpv1VersionCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIgmpv1VersionCounter{}
	}
	return &patternFlowIgmpv1VersionCounter{obj: obj.obj.Decrement}

}

type patternFlowIgmpv1Type struct {
	obj *snappipb.PatternFlowIgmpv1Type
}

func (obj *patternFlowIgmpv1Type) msg() *snappipb.PatternFlowIgmpv1Type {
	return obj.obj
}

func (obj *patternFlowIgmpv1Type) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIgmpv1Type) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIgmpv1Type interface {
	msg() *snappipb.PatternFlowIgmpv1Type
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIgmpv1Type
	Values() []int32
	SetValues(value []int32) PatternFlowIgmpv1Type
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIgmpv1Type
	Increment() PatternFlowIgmpv1TypeCounter
	Decrement() PatternFlowIgmpv1TypeCounter
}

func (obj *patternFlowIgmpv1Type) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIgmpv1Type) SetValue(value int32) PatternFlowIgmpv1Type {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIgmpv1Type) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIgmpv1Type) SetValues(value []int32) PatternFlowIgmpv1Type {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIgmpv1Type) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIgmpv1Type) SetMetricGroup(value string) PatternFlowIgmpv1Type {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIgmpv1Type) Increment() PatternFlowIgmpv1TypeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIgmpv1TypeCounter{}
	}
	return &patternFlowIgmpv1TypeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIgmpv1Type) Decrement() PatternFlowIgmpv1TypeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIgmpv1TypeCounter{}
	}
	return &patternFlowIgmpv1TypeCounter{obj: obj.obj.Decrement}

}

type patternFlowIgmpv1Unused struct {
	obj *snappipb.PatternFlowIgmpv1Unused
}

func (obj *patternFlowIgmpv1Unused) msg() *snappipb.PatternFlowIgmpv1Unused {
	return obj.obj
}

func (obj *patternFlowIgmpv1Unused) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIgmpv1Unused) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIgmpv1Unused interface {
	msg() *snappipb.PatternFlowIgmpv1Unused
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIgmpv1Unused
	Values() []int32
	SetValues(value []int32) PatternFlowIgmpv1Unused
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIgmpv1Unused
	Increment() PatternFlowIgmpv1UnusedCounter
	Decrement() PatternFlowIgmpv1UnusedCounter
}

func (obj *patternFlowIgmpv1Unused) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIgmpv1Unused) SetValue(value int32) PatternFlowIgmpv1Unused {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIgmpv1Unused) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIgmpv1Unused) SetValues(value []int32) PatternFlowIgmpv1Unused {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIgmpv1Unused) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIgmpv1Unused) SetMetricGroup(value string) PatternFlowIgmpv1Unused {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIgmpv1Unused) Increment() PatternFlowIgmpv1UnusedCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIgmpv1UnusedCounter{}
	}
	return &patternFlowIgmpv1UnusedCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIgmpv1Unused) Decrement() PatternFlowIgmpv1UnusedCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIgmpv1UnusedCounter{}
	}
	return &patternFlowIgmpv1UnusedCounter{obj: obj.obj.Decrement}

}

type patternFlowIgmpv1Checksum struct {
	obj *snappipb.PatternFlowIgmpv1Checksum
}

func (obj *patternFlowIgmpv1Checksum) msg() *snappipb.PatternFlowIgmpv1Checksum {
	return obj.obj
}

func (obj *patternFlowIgmpv1Checksum) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIgmpv1Checksum) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIgmpv1Checksum interface {
	msg() *snappipb.PatternFlowIgmpv1Checksum
	Yaml() string
	Json() string
	Custom() int32
	SetCustom(value int32) PatternFlowIgmpv1Checksum
}

func (obj *patternFlowIgmpv1Checksum) Custom() int32 {
	return *obj.obj.Custom
}

func (obj *patternFlowIgmpv1Checksum) SetCustom(value int32) PatternFlowIgmpv1Checksum {
	obj.obj.Custom = &value
	return obj
}

type patternFlowIgmpv1GroupAddress struct {
	obj *snappipb.PatternFlowIgmpv1GroupAddress
}

func (obj *patternFlowIgmpv1GroupAddress) msg() *snappipb.PatternFlowIgmpv1GroupAddress {
	return obj.obj
}

func (obj *patternFlowIgmpv1GroupAddress) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIgmpv1GroupAddress) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIgmpv1GroupAddress interface {
	msg() *snappipb.PatternFlowIgmpv1GroupAddress
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternFlowIgmpv1GroupAddress
	Values() []string
	SetValues(value []string) PatternFlowIgmpv1GroupAddress
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIgmpv1GroupAddress
	Increment() PatternFlowIgmpv1GroupAddressCounter
	Decrement() PatternFlowIgmpv1GroupAddressCounter
}

func (obj *patternFlowIgmpv1GroupAddress) Value() string {
	return *obj.obj.Value
}

func (obj *patternFlowIgmpv1GroupAddress) SetValue(value string) PatternFlowIgmpv1GroupAddress {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIgmpv1GroupAddress) Values() []string {
	return obj.obj.Values
}

func (obj *patternFlowIgmpv1GroupAddress) SetValues(value []string) PatternFlowIgmpv1GroupAddress {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIgmpv1GroupAddress) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIgmpv1GroupAddress) SetMetricGroup(value string) PatternFlowIgmpv1GroupAddress {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIgmpv1GroupAddress) Increment() PatternFlowIgmpv1GroupAddressCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIgmpv1GroupAddressCounter{}
	}
	return &patternFlowIgmpv1GroupAddressCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIgmpv1GroupAddress) Decrement() PatternFlowIgmpv1GroupAddressCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIgmpv1GroupAddressCounter{}
	}
	return &patternFlowIgmpv1GroupAddressCounter{obj: obj.obj.Decrement}

}

type flowDelay struct {
	obj *snappipb.FlowDelay
}

func (obj *flowDelay) msg() *snappipb.FlowDelay {
	return obj.obj
}

func (obj *flowDelay) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowDelay) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowDelay interface {
	msg() *snappipb.FlowDelay
	Yaml() string
	Json() string
	Bytes() float32
	SetBytes(value float32) FlowDelay
	Nanoseconds() float32
	SetNanoseconds(value float32) FlowDelay
	Microseconds() float32
	SetMicroseconds(value float32) FlowDelay
}

func (obj *flowDelay) Bytes() float32 {
	return *obj.obj.Bytes
}

func (obj *flowDelay) SetBytes(value float32) FlowDelay {
	obj.obj.Bytes = &value
	return obj
}

func (obj *flowDelay) Nanoseconds() float32 {
	return *obj.obj.Nanoseconds
}

func (obj *flowDelay) SetNanoseconds(value float32) FlowDelay {
	obj.obj.Nanoseconds = &value
	return obj
}

func (obj *flowDelay) Microseconds() float32 {
	return *obj.obj.Microseconds
}

func (obj *flowDelay) SetMicroseconds(value float32) FlowDelay {
	obj.obj.Microseconds = &value
	return obj
}

type flowDurationInterBurstGap struct {
	obj *snappipb.FlowDurationInterBurstGap
}

func (obj *flowDurationInterBurstGap) msg() *snappipb.FlowDurationInterBurstGap {
	return obj.obj
}

func (obj *flowDurationInterBurstGap) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowDurationInterBurstGap) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowDurationInterBurstGap interface {
	msg() *snappipb.FlowDurationInterBurstGap
	Yaml() string
	Json() string
	Bytes() float64
	SetBytes(value float64) FlowDurationInterBurstGap
	Nanoseconds() float64
	SetNanoseconds(value float64) FlowDurationInterBurstGap
	Microseconds() float64
	SetMicroseconds(value float64) FlowDurationInterBurstGap
}

func (obj *flowDurationInterBurstGap) Bytes() float64 {
	return *obj.obj.Bytes
}

func (obj *flowDurationInterBurstGap) SetBytes(value float64) FlowDurationInterBurstGap {
	obj.obj.Bytes = &value
	return obj
}

func (obj *flowDurationInterBurstGap) Nanoseconds() float64 {
	return *obj.obj.Nanoseconds
}

func (obj *flowDurationInterBurstGap) SetNanoseconds(value float64) FlowDurationInterBurstGap {
	obj.obj.Nanoseconds = &value
	return obj
}

func (obj *flowDurationInterBurstGap) Microseconds() float64 {
	return *obj.obj.Microseconds
}

func (obj *flowDurationInterBurstGap) SetMicroseconds(value float64) FlowDurationInterBurstGap {
	obj.obj.Microseconds = &value
	return obj
}

type deviceBgpAdvanced struct {
	obj *snappipb.DeviceBgpAdvanced
}

func (obj *deviceBgpAdvanced) msg() *snappipb.DeviceBgpAdvanced {
	return obj.obj
}

func (obj *deviceBgpAdvanced) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpAdvanced) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpAdvanced interface {
	msg() *snappipb.DeviceBgpAdvanced
	Yaml() string
	Json() string
	HoldTimeInterval() int32
	SetHoldTimeInterval(value int32) DeviceBgpAdvanced
	KeepAliveInterval() int32
	SetKeepAliveInterval(value int32) DeviceBgpAdvanced
	UpdateInterval() int32
	SetUpdateInterval(value int32) DeviceBgpAdvanced
	TimeToLive() int32
	SetTimeToLive(value int32) DeviceBgpAdvanced
	Md5Key() string
	SetMd5Key(value string) DeviceBgpAdvanced
}

func (obj *deviceBgpAdvanced) HoldTimeInterval() int32 {
	return *obj.obj.HoldTimeInterval
}

func (obj *deviceBgpAdvanced) SetHoldTimeInterval(value int32) DeviceBgpAdvanced {
	obj.obj.HoldTimeInterval = &value
	return obj
}

func (obj *deviceBgpAdvanced) KeepAliveInterval() int32 {
	return *obj.obj.KeepAliveInterval
}

func (obj *deviceBgpAdvanced) SetKeepAliveInterval(value int32) DeviceBgpAdvanced {
	obj.obj.KeepAliveInterval = &value
	return obj
}

func (obj *deviceBgpAdvanced) UpdateInterval() int32 {
	return *obj.obj.UpdateInterval
}

func (obj *deviceBgpAdvanced) SetUpdateInterval(value int32) DeviceBgpAdvanced {
	obj.obj.UpdateInterval = &value
	return obj
}

func (obj *deviceBgpAdvanced) TimeToLive() int32 {
	return *obj.obj.TimeToLive
}

func (obj *deviceBgpAdvanced) SetTimeToLive(value int32) DeviceBgpAdvanced {
	obj.obj.TimeToLive = &value
	return obj
}

func (obj *deviceBgpAdvanced) Md5Key() string {
	return *obj.obj.Md5Key
}

func (obj *deviceBgpAdvanced) SetMd5Key(value string) DeviceBgpAdvanced {
	obj.obj.Md5Key = &value
	return obj
}

type deviceBgpCapability struct {
	obj *snappipb.DeviceBgpCapability
}

func (obj *deviceBgpCapability) msg() *snappipb.DeviceBgpCapability {
	return obj.obj
}

func (obj *deviceBgpCapability) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpCapability) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpCapability interface {
	msg() *snappipb.DeviceBgpCapability
	Yaml() string
	Json() string
	Vpls() bool
	SetVpls(value bool) DeviceBgpCapability
	RouteRefresh() bool
	SetRouteRefresh(value bool) DeviceBgpCapability
	RouteConstraint() bool
	SetRouteConstraint(value bool) DeviceBgpCapability
	LinkStateNonVpn() bool
	SetLinkStateNonVpn(value bool) DeviceBgpCapability
	LinkStateVpn() bool
	SetLinkStateVpn(value bool) DeviceBgpCapability
	Evpn() bool
	SetEvpn(value bool) DeviceBgpCapability
	ExtendedNextHopEncoding() bool
	SetExtendedNextHopEncoding(value bool) DeviceBgpCapability
	Ipv4Unicast() bool
	SetIpv4Unicast(value bool) DeviceBgpCapability
	Ipv4Multicast() bool
	SetIpv4Multicast(value bool) DeviceBgpCapability
	Ipv4MulticastVpn() bool
	SetIpv4MulticastVpn(value bool) DeviceBgpCapability
	Ipv4MplsVpn() bool
	SetIpv4MplsVpn(value bool) DeviceBgpCapability
	Ipv4Mdt() bool
	SetIpv4Mdt(value bool) DeviceBgpCapability
	Ipv4MulticastMplsVpn() bool
	SetIpv4MulticastMplsVpn(value bool) DeviceBgpCapability
	Ipv4UnicastFlowSpec() bool
	SetIpv4UnicastFlowSpec(value bool) DeviceBgpCapability
	Ipv4SrTePolicy() bool
	SetIpv4SrTePolicy(value bool) DeviceBgpCapability
	Ipv4UnicastAddPath() bool
	SetIpv4UnicastAddPath(value bool) DeviceBgpCapability
	Ipv6Unicast() bool
	SetIpv6Unicast(value bool) DeviceBgpCapability
	Ipv6Multicast() bool
	SetIpv6Multicast(value bool) DeviceBgpCapability
	Ipv6MulticastVpn() bool
	SetIpv6MulticastVpn(value bool) DeviceBgpCapability
	Ipv6MplsVpn() bool
	SetIpv6MplsVpn(value bool) DeviceBgpCapability
	Ipv6Mdt() bool
	SetIpv6Mdt(value bool) DeviceBgpCapability
	Ipv6MulticastMplsVpn() bool
	SetIpv6MulticastMplsVpn(value bool) DeviceBgpCapability
	Ipv6UnicastFlowSpec() bool
	SetIpv6UnicastFlowSpec(value bool) DeviceBgpCapability
	Ipv6SrTePolicy() bool
	SetIpv6SrTePolicy(value bool) DeviceBgpCapability
	Ipv6UnicastAddPath() bool
	SetIpv6UnicastAddPath(value bool) DeviceBgpCapability
}

func (obj *deviceBgpCapability) Vpls() bool {
	return *obj.obj.Vpls
}

func (obj *deviceBgpCapability) SetVpls(value bool) DeviceBgpCapability {
	obj.obj.Vpls = &value
	return obj
}

func (obj *deviceBgpCapability) RouteRefresh() bool {
	return *obj.obj.RouteRefresh
}

func (obj *deviceBgpCapability) SetRouteRefresh(value bool) DeviceBgpCapability {
	obj.obj.RouteRefresh = &value
	return obj
}

func (obj *deviceBgpCapability) RouteConstraint() bool {
	return *obj.obj.RouteConstraint
}

func (obj *deviceBgpCapability) SetRouteConstraint(value bool) DeviceBgpCapability {
	obj.obj.RouteConstraint = &value
	return obj
}

func (obj *deviceBgpCapability) LinkStateNonVpn() bool {
	return *obj.obj.LinkStateNonVpn
}

func (obj *deviceBgpCapability) SetLinkStateNonVpn(value bool) DeviceBgpCapability {
	obj.obj.LinkStateNonVpn = &value
	return obj
}

func (obj *deviceBgpCapability) LinkStateVpn() bool {
	return *obj.obj.LinkStateVpn
}

func (obj *deviceBgpCapability) SetLinkStateVpn(value bool) DeviceBgpCapability {
	obj.obj.LinkStateVpn = &value
	return obj
}

func (obj *deviceBgpCapability) Evpn() bool {
	return *obj.obj.Evpn
}

func (obj *deviceBgpCapability) SetEvpn(value bool) DeviceBgpCapability {
	obj.obj.Evpn = &value
	return obj
}

func (obj *deviceBgpCapability) ExtendedNextHopEncoding() bool {
	return *obj.obj.ExtendedNextHopEncoding
}

func (obj *deviceBgpCapability) SetExtendedNextHopEncoding(value bool) DeviceBgpCapability {
	obj.obj.ExtendedNextHopEncoding = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv4Unicast() bool {
	return *obj.obj.Ipv4Unicast
}

func (obj *deviceBgpCapability) SetIpv4Unicast(value bool) DeviceBgpCapability {
	obj.obj.Ipv4Unicast = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv4Multicast() bool {
	return *obj.obj.Ipv4Multicast
}

func (obj *deviceBgpCapability) SetIpv4Multicast(value bool) DeviceBgpCapability {
	obj.obj.Ipv4Multicast = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv4MulticastVpn() bool {
	return *obj.obj.Ipv4MulticastVpn
}

func (obj *deviceBgpCapability) SetIpv4MulticastVpn(value bool) DeviceBgpCapability {
	obj.obj.Ipv4MulticastVpn = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv4MplsVpn() bool {
	return *obj.obj.Ipv4MplsVpn
}

func (obj *deviceBgpCapability) SetIpv4MplsVpn(value bool) DeviceBgpCapability {
	obj.obj.Ipv4MplsVpn = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv4Mdt() bool {
	return *obj.obj.Ipv4Mdt
}

func (obj *deviceBgpCapability) SetIpv4Mdt(value bool) DeviceBgpCapability {
	obj.obj.Ipv4Mdt = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv4MulticastMplsVpn() bool {
	return *obj.obj.Ipv4MulticastMplsVpn
}

func (obj *deviceBgpCapability) SetIpv4MulticastMplsVpn(value bool) DeviceBgpCapability {
	obj.obj.Ipv4MulticastMplsVpn = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv4UnicastFlowSpec() bool {
	return *obj.obj.Ipv4UnicastFlowSpec
}

func (obj *deviceBgpCapability) SetIpv4UnicastFlowSpec(value bool) DeviceBgpCapability {
	obj.obj.Ipv4UnicastFlowSpec = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv4SrTePolicy() bool {
	return *obj.obj.Ipv4SrTePolicy
}

func (obj *deviceBgpCapability) SetIpv4SrTePolicy(value bool) DeviceBgpCapability {
	obj.obj.Ipv4SrTePolicy = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv4UnicastAddPath() bool {
	return *obj.obj.Ipv4UnicastAddPath
}

func (obj *deviceBgpCapability) SetIpv4UnicastAddPath(value bool) DeviceBgpCapability {
	obj.obj.Ipv4UnicastAddPath = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv6Unicast() bool {
	return *obj.obj.Ipv6Unicast
}

func (obj *deviceBgpCapability) SetIpv6Unicast(value bool) DeviceBgpCapability {
	obj.obj.Ipv6Unicast = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv6Multicast() bool {
	return *obj.obj.Ipv6Multicast
}

func (obj *deviceBgpCapability) SetIpv6Multicast(value bool) DeviceBgpCapability {
	obj.obj.Ipv6Multicast = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv6MulticastVpn() bool {
	return *obj.obj.Ipv6MulticastVpn
}

func (obj *deviceBgpCapability) SetIpv6MulticastVpn(value bool) DeviceBgpCapability {
	obj.obj.Ipv6MulticastVpn = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv6MplsVpn() bool {
	return *obj.obj.Ipv6MplsVpn
}

func (obj *deviceBgpCapability) SetIpv6MplsVpn(value bool) DeviceBgpCapability {
	obj.obj.Ipv6MplsVpn = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv6Mdt() bool {
	return *obj.obj.Ipv6Mdt
}

func (obj *deviceBgpCapability) SetIpv6Mdt(value bool) DeviceBgpCapability {
	obj.obj.Ipv6Mdt = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv6MulticastMplsVpn() bool {
	return *obj.obj.Ipv6MulticastMplsVpn
}

func (obj *deviceBgpCapability) SetIpv6MulticastMplsVpn(value bool) DeviceBgpCapability {
	obj.obj.Ipv6MulticastMplsVpn = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv6UnicastFlowSpec() bool {
	return *obj.obj.Ipv6UnicastFlowSpec
}

func (obj *deviceBgpCapability) SetIpv6UnicastFlowSpec(value bool) DeviceBgpCapability {
	obj.obj.Ipv6UnicastFlowSpec = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv6SrTePolicy() bool {
	return *obj.obj.Ipv6SrTePolicy
}

func (obj *deviceBgpCapability) SetIpv6SrTePolicy(value bool) DeviceBgpCapability {
	obj.obj.Ipv6SrTePolicy = &value
	return obj
}

func (obj *deviceBgpCapability) Ipv6UnicastAddPath() bool {
	return *obj.obj.Ipv6UnicastAddPath
}

func (obj *deviceBgpCapability) SetIpv6UnicastAddPath(value bool) DeviceBgpCapability {
	obj.obj.Ipv6UnicastAddPath = &value
	return obj
}

type deviceBgpSrTePolicy struct {
	obj *snappipb.DeviceBgpSrTePolicy
}

func (obj *deviceBgpSrTePolicy) msg() *snappipb.DeviceBgpSrTePolicy {
	return obj.obj
}

func (obj *deviceBgpSrTePolicy) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpSrTePolicy) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpSrTePolicy interface {
	msg() *snappipb.DeviceBgpSrTePolicy
	Yaml() string
	Json() string
	Distinguisher() int32
	SetDistinguisher(value int32) DeviceBgpSrTePolicy
	Color() int32
	SetColor(value int32) DeviceBgpSrTePolicy
	Ipv4Endpoint() string
	SetIpv4Endpoint(value string) DeviceBgpSrTePolicy
	Ipv6Endpoint() string
	SetIpv6Endpoint(value string) DeviceBgpSrTePolicy
	NextHop() DeviceBgpSrTePolicyNextHop
	AddPath() DeviceBgpAddPath
	AsPath() DeviceBgpAsPath
	TunnelTlvs() []DeviceBgpTunnelTlv
	NewTunnelTlvs() DeviceBgpTunnelTlv
	Communities() []DeviceBgpCommunity
	NewCommunities() DeviceBgpCommunity
}

func (obj *deviceBgpSrTePolicy) Distinguisher() int32 {
	return *obj.obj.Distinguisher
}

func (obj *deviceBgpSrTePolicy) SetDistinguisher(value int32) DeviceBgpSrTePolicy {
	obj.obj.Distinguisher = &value
	return obj
}

func (obj *deviceBgpSrTePolicy) Color() int32 {
	return *obj.obj.Color
}

func (obj *deviceBgpSrTePolicy) SetColor(value int32) DeviceBgpSrTePolicy {
	obj.obj.Color = &value
	return obj
}

func (obj *deviceBgpSrTePolicy) Ipv4Endpoint() string {
	return *obj.obj.Ipv4Endpoint
}

func (obj *deviceBgpSrTePolicy) SetIpv4Endpoint(value string) DeviceBgpSrTePolicy {
	obj.obj.Ipv4Endpoint = &value
	return obj
}

func (obj *deviceBgpSrTePolicy) Ipv6Endpoint() string {
	return *obj.obj.Ipv6Endpoint
}

func (obj *deviceBgpSrTePolicy) SetIpv6Endpoint(value string) DeviceBgpSrTePolicy {
	obj.obj.Ipv6Endpoint = &value
	return obj
}

func (obj *deviceBgpSrTePolicy) NextHop() DeviceBgpSrTePolicyNextHop {
	if obj.obj.NextHop == nil {
		obj.obj.NextHop = &snappipb.DeviceBgpSrTePolicyNextHop{}
	}
	return &deviceBgpSrTePolicyNextHop{obj: obj.obj.NextHop}

}

func (obj *deviceBgpSrTePolicy) AddPath() DeviceBgpAddPath {
	if obj.obj.AddPath == nil {
		obj.obj.AddPath = &snappipb.DeviceBgpAddPath{}
	}
	return &deviceBgpAddPath{obj: obj.obj.AddPath}

}

func (obj *deviceBgpSrTePolicy) AsPath() DeviceBgpAsPath {
	if obj.obj.AsPath == nil {
		obj.obj.AsPath = &snappipb.DeviceBgpAsPath{}
	}
	return &deviceBgpAsPath{obj: obj.obj.AsPath}

}

func (obj *deviceBgpSrTePolicy) TunnelTlvs() []DeviceBgpTunnelTlv {
	if obj.obj.TunnelTlvs == nil {
		obj.obj.TunnelTlvs = make([]*snappipb.DeviceBgpTunnelTlv, 0)
	}
	values := make([]DeviceBgpTunnelTlv, 0)
	for _, item := range obj.obj.TunnelTlvs {
		values = append(values, &deviceBgpTunnelTlv{obj: item})
	}
	return values

}

func (obj *deviceBgpSrTePolicy) NewTunnelTlvs() DeviceBgpTunnelTlv {
	if obj.obj.TunnelTlvs == nil {
		obj.obj.TunnelTlvs = make([]*snappipb.DeviceBgpTunnelTlv, 0)
	}
	slice := append(obj.obj.TunnelTlvs, &snappipb.DeviceBgpTunnelTlv{})
	obj.obj.TunnelTlvs = slice
	return &deviceBgpTunnelTlv{obj: slice[len(slice)-1]}
}

func (obj *deviceBgpSrTePolicy) Communities() []DeviceBgpCommunity {
	if obj.obj.Communities == nil {
		obj.obj.Communities = make([]*snappipb.DeviceBgpCommunity, 0)
	}
	values := make([]DeviceBgpCommunity, 0)
	for _, item := range obj.obj.Communities {
		values = append(values, &deviceBgpCommunity{obj: item})
	}
	return values

}

func (obj *deviceBgpSrTePolicy) NewCommunities() DeviceBgpCommunity {
	if obj.obj.Communities == nil {
		obj.obj.Communities = make([]*snappipb.DeviceBgpCommunity, 0)
	}
	slice := append(obj.obj.Communities, &snappipb.DeviceBgpCommunity{})
	obj.obj.Communities = slice
	return &deviceBgpCommunity{obj: slice[len(slice)-1]}
}

type deviceBgpv4Route struct {
	obj *snappipb.DeviceBgpv4Route
}

func (obj *deviceBgpv4Route) msg() *snappipb.DeviceBgpv4Route {
	return obj.obj
}

func (obj *deviceBgpv4Route) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpv4Route) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpv4Route interface {
	msg() *snappipb.DeviceBgpv4Route
	Yaml() string
	Json() string
	Addresses() []DeviceBgpv4RouteAddress
	NewAddresses() DeviceBgpv4RouteAddress
	NextHopAddress() string
	SetNextHopAddress(value string) DeviceBgpv4Route
	Advanced() DeviceBgpRouteAdvanced
	Communities() []DeviceBgpCommunity
	NewCommunities() DeviceBgpCommunity
	AsPath() DeviceBgpAsPath
	AddPath() DeviceBgpAddPath
	Name() string
	SetName(value string) DeviceBgpv4Route
}

func (obj *deviceBgpv4Route) Addresses() []DeviceBgpv4RouteAddress {
	if obj.obj.Addresses == nil {
		obj.obj.Addresses = make([]*snappipb.DeviceBgpv4RouteAddress, 0)
	}
	values := make([]DeviceBgpv4RouteAddress, 0)
	for _, item := range obj.obj.Addresses {
		values = append(values, &deviceBgpv4RouteAddress{obj: item})
	}
	return values

}

func (obj *deviceBgpv4Route) NewAddresses() DeviceBgpv4RouteAddress {
	if obj.obj.Addresses == nil {
		obj.obj.Addresses = make([]*snappipb.DeviceBgpv4RouteAddress, 0)
	}
	slice := append(obj.obj.Addresses, &snappipb.DeviceBgpv4RouteAddress{})
	obj.obj.Addresses = slice
	return &deviceBgpv4RouteAddress{obj: slice[len(slice)-1]}
}

func (obj *deviceBgpv4Route) NextHopAddress() string {
	return *obj.obj.NextHopAddress
}

func (obj *deviceBgpv4Route) SetNextHopAddress(value string) DeviceBgpv4Route {
	obj.obj.NextHopAddress = &value
	return obj
}

func (obj *deviceBgpv4Route) Advanced() DeviceBgpRouteAdvanced {
	if obj.obj.Advanced == nil {
		obj.obj.Advanced = &snappipb.DeviceBgpRouteAdvanced{}
	}
	return &deviceBgpRouteAdvanced{obj: obj.obj.Advanced}

}

func (obj *deviceBgpv4Route) Communities() []DeviceBgpCommunity {
	if obj.obj.Communities == nil {
		obj.obj.Communities = make([]*snappipb.DeviceBgpCommunity, 0)
	}
	values := make([]DeviceBgpCommunity, 0)
	for _, item := range obj.obj.Communities {
		values = append(values, &deviceBgpCommunity{obj: item})
	}
	return values

}

func (obj *deviceBgpv4Route) NewCommunities() DeviceBgpCommunity {
	if obj.obj.Communities == nil {
		obj.obj.Communities = make([]*snappipb.DeviceBgpCommunity, 0)
	}
	slice := append(obj.obj.Communities, &snappipb.DeviceBgpCommunity{})
	obj.obj.Communities = slice
	return &deviceBgpCommunity{obj: slice[len(slice)-1]}
}

func (obj *deviceBgpv4Route) AsPath() DeviceBgpAsPath {
	if obj.obj.AsPath == nil {
		obj.obj.AsPath = &snappipb.DeviceBgpAsPath{}
	}
	return &deviceBgpAsPath{obj: obj.obj.AsPath}

}

func (obj *deviceBgpv4Route) AddPath() DeviceBgpAddPath {
	if obj.obj.AddPath == nil {
		obj.obj.AddPath = &snappipb.DeviceBgpAddPath{}
	}
	return &deviceBgpAddPath{obj: obj.obj.AddPath}

}

func (obj *deviceBgpv4Route) Name() string {
	return obj.obj.Name
}

func (obj *deviceBgpv4Route) SetName(value string) DeviceBgpv4Route {
	obj.obj.Name = value
	return obj
}

type deviceBgpv6Route struct {
	obj *snappipb.DeviceBgpv6Route
}

func (obj *deviceBgpv6Route) msg() *snappipb.DeviceBgpv6Route {
	return obj.obj
}

func (obj *deviceBgpv6Route) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpv6Route) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpv6Route interface {
	msg() *snappipb.DeviceBgpv6Route
	Yaml() string
	Json() string
	Addresses() []DeviceBgpv6RouteAddress
	NewAddresses() DeviceBgpv6RouteAddress
	NextHopAddress() string
	SetNextHopAddress(value string) DeviceBgpv6Route
	Advanced() DeviceBgpRouteAdvanced
	Communities() []DeviceBgpCommunity
	NewCommunities() DeviceBgpCommunity
	AsPath() DeviceBgpAsPath
	AddPath() DeviceBgpAddPath
	Name() string
	SetName(value string) DeviceBgpv6Route
}

func (obj *deviceBgpv6Route) Addresses() []DeviceBgpv6RouteAddress {
	if obj.obj.Addresses == nil {
		obj.obj.Addresses = make([]*snappipb.DeviceBgpv6RouteAddress, 0)
	}
	values := make([]DeviceBgpv6RouteAddress, 0)
	for _, item := range obj.obj.Addresses {
		values = append(values, &deviceBgpv6RouteAddress{obj: item})
	}
	return values

}

func (obj *deviceBgpv6Route) NewAddresses() DeviceBgpv6RouteAddress {
	if obj.obj.Addresses == nil {
		obj.obj.Addresses = make([]*snappipb.DeviceBgpv6RouteAddress, 0)
	}
	slice := append(obj.obj.Addresses, &snappipb.DeviceBgpv6RouteAddress{})
	obj.obj.Addresses = slice
	return &deviceBgpv6RouteAddress{obj: slice[len(slice)-1]}
}

func (obj *deviceBgpv6Route) NextHopAddress() string {
	return *obj.obj.NextHopAddress
}

func (obj *deviceBgpv6Route) SetNextHopAddress(value string) DeviceBgpv6Route {
	obj.obj.NextHopAddress = &value
	return obj
}

func (obj *deviceBgpv6Route) Advanced() DeviceBgpRouteAdvanced {
	if obj.obj.Advanced == nil {
		obj.obj.Advanced = &snappipb.DeviceBgpRouteAdvanced{}
	}
	return &deviceBgpRouteAdvanced{obj: obj.obj.Advanced}

}

func (obj *deviceBgpv6Route) Communities() []DeviceBgpCommunity {
	if obj.obj.Communities == nil {
		obj.obj.Communities = make([]*snappipb.DeviceBgpCommunity, 0)
	}
	values := make([]DeviceBgpCommunity, 0)
	for _, item := range obj.obj.Communities {
		values = append(values, &deviceBgpCommunity{obj: item})
	}
	return values

}

func (obj *deviceBgpv6Route) NewCommunities() DeviceBgpCommunity {
	if obj.obj.Communities == nil {
		obj.obj.Communities = make([]*snappipb.DeviceBgpCommunity, 0)
	}
	slice := append(obj.obj.Communities, &snappipb.DeviceBgpCommunity{})
	obj.obj.Communities = slice
	return &deviceBgpCommunity{obj: slice[len(slice)-1]}
}

func (obj *deviceBgpv6Route) AsPath() DeviceBgpAsPath {
	if obj.obj.AsPath == nil {
		obj.obj.AsPath = &snappipb.DeviceBgpAsPath{}
	}
	return &deviceBgpAsPath{obj: obj.obj.AsPath}

}

func (obj *deviceBgpv6Route) AddPath() DeviceBgpAddPath {
	if obj.obj.AddPath == nil {
		obj.obj.AddPath = &snappipb.DeviceBgpAddPath{}
	}
	return &deviceBgpAddPath{obj: obj.obj.AddPath}

}

func (obj *deviceBgpv6Route) Name() string {
	return obj.obj.Name
}

func (obj *deviceBgpv6Route) SetName(value string) DeviceBgpv6Route {
	obj.obj.Name = value
	return obj
}

type deviceBgpv6SegmentRouting struct {
	obj *snappipb.DeviceBgpv6SegmentRouting
}

func (obj *deviceBgpv6SegmentRouting) msg() *snappipb.DeviceBgpv6SegmentRouting {
	return obj.obj
}

func (obj *deviceBgpv6SegmentRouting) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpv6SegmentRouting) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpv6SegmentRouting interface {
	msg() *snappipb.DeviceBgpv6SegmentRouting
	Yaml() string
	Json() string
	IngressSupportsVpn() bool
	SetIngressSupportsVpn(value bool) DeviceBgpv6SegmentRouting
	ReducedEncapsulation() bool
	SetReducedEncapsulation(value bool) DeviceBgpv6SegmentRouting
	CopyTimeToLive() bool
	SetCopyTimeToLive(value bool) DeviceBgpv6SegmentRouting
	TimeToLive() int32
	SetTimeToLive(value int32) DeviceBgpv6SegmentRouting
	MaxSidsPerSrh() int32
	SetMaxSidsPerSrh(value int32) DeviceBgpv6SegmentRouting
	AutoGenerateSegmentLeftValue() bool
	SetAutoGenerateSegmentLeftValue(value bool) DeviceBgpv6SegmentRouting
	SegmentLeftValue() int32
	SetSegmentLeftValue(value int32) DeviceBgpv6SegmentRouting
	AdvertiseSrTePolicy() bool
	SetAdvertiseSrTePolicy(value bool) DeviceBgpv6SegmentRouting
}

func (obj *deviceBgpv6SegmentRouting) IngressSupportsVpn() bool {
	return *obj.obj.IngressSupportsVpn
}

func (obj *deviceBgpv6SegmentRouting) SetIngressSupportsVpn(value bool) DeviceBgpv6SegmentRouting {
	obj.obj.IngressSupportsVpn = &value
	return obj
}

func (obj *deviceBgpv6SegmentRouting) ReducedEncapsulation() bool {
	return *obj.obj.ReducedEncapsulation
}

func (obj *deviceBgpv6SegmentRouting) SetReducedEncapsulation(value bool) DeviceBgpv6SegmentRouting {
	obj.obj.ReducedEncapsulation = &value
	return obj
}

func (obj *deviceBgpv6SegmentRouting) CopyTimeToLive() bool {
	return *obj.obj.CopyTimeToLive
}

func (obj *deviceBgpv6SegmentRouting) SetCopyTimeToLive(value bool) DeviceBgpv6SegmentRouting {
	obj.obj.CopyTimeToLive = &value
	return obj
}

func (obj *deviceBgpv6SegmentRouting) TimeToLive() int32 {
	return *obj.obj.TimeToLive
}

func (obj *deviceBgpv6SegmentRouting) SetTimeToLive(value int32) DeviceBgpv6SegmentRouting {
	obj.obj.TimeToLive = &value
	return obj
}

func (obj *deviceBgpv6SegmentRouting) MaxSidsPerSrh() int32 {
	return *obj.obj.MaxSidsPerSrh
}

func (obj *deviceBgpv6SegmentRouting) SetMaxSidsPerSrh(value int32) DeviceBgpv6SegmentRouting {
	obj.obj.MaxSidsPerSrh = &value
	return obj
}

func (obj *deviceBgpv6SegmentRouting) AutoGenerateSegmentLeftValue() bool {
	return *obj.obj.AutoGenerateSegmentLeftValue
}

func (obj *deviceBgpv6SegmentRouting) SetAutoGenerateSegmentLeftValue(value bool) DeviceBgpv6SegmentRouting {
	obj.obj.AutoGenerateSegmentLeftValue = &value
	return obj
}

func (obj *deviceBgpv6SegmentRouting) SegmentLeftValue() int32 {
	return *obj.obj.SegmentLeftValue
}

func (obj *deviceBgpv6SegmentRouting) SetSegmentLeftValue(value int32) DeviceBgpv6SegmentRouting {
	obj.obj.SegmentLeftValue = &value
	return obj
}

func (obj *deviceBgpv6SegmentRouting) AdvertiseSrTePolicy() bool {
	return *obj.obj.AdvertiseSrTePolicy
}

func (obj *deviceBgpv6SegmentRouting) SetAdvertiseSrTePolicy(value bool) DeviceBgpv6SegmentRouting {
	obj.obj.AdvertiseSrTePolicy = &value
	return obj
}

type patternFlowEthernetDstCounter struct {
	obj *snappipb.PatternFlowEthernetDstCounter
}

func (obj *patternFlowEthernetDstCounter) msg() *snappipb.PatternFlowEthernetDstCounter {
	return obj.obj
}

func (obj *patternFlowEthernetDstCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetDstCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetDstCounter interface {
	msg() *snappipb.PatternFlowEthernetDstCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowEthernetDstCounter
	Step() string
	SetStep(value string) PatternFlowEthernetDstCounter
	Count() int32
	SetCount(value int32) PatternFlowEthernetDstCounter
}

func (obj *patternFlowEthernetDstCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowEthernetDstCounter) SetStart(value string) PatternFlowEthernetDstCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowEthernetDstCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowEthernetDstCounter) SetStep(value string) PatternFlowEthernetDstCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowEthernetDstCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowEthernetDstCounter) SetCount(value int32) PatternFlowEthernetDstCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowEthernetSrcCounter struct {
	obj *snappipb.PatternFlowEthernetSrcCounter
}

func (obj *patternFlowEthernetSrcCounter) msg() *snappipb.PatternFlowEthernetSrcCounter {
	return obj.obj
}

func (obj *patternFlowEthernetSrcCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetSrcCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetSrcCounter interface {
	msg() *snappipb.PatternFlowEthernetSrcCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowEthernetSrcCounter
	Step() string
	SetStep(value string) PatternFlowEthernetSrcCounter
	Count() int32
	SetCount(value int32) PatternFlowEthernetSrcCounter
}

func (obj *patternFlowEthernetSrcCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowEthernetSrcCounter) SetStart(value string) PatternFlowEthernetSrcCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowEthernetSrcCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowEthernetSrcCounter) SetStep(value string) PatternFlowEthernetSrcCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowEthernetSrcCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowEthernetSrcCounter) SetCount(value int32) PatternFlowEthernetSrcCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowEthernetEtherTypeCounter struct {
	obj *snappipb.PatternFlowEthernetEtherTypeCounter
}

func (obj *patternFlowEthernetEtherTypeCounter) msg() *snappipb.PatternFlowEthernetEtherTypeCounter {
	return obj.obj
}

func (obj *patternFlowEthernetEtherTypeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetEtherTypeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetEtherTypeCounter interface {
	msg() *snappipb.PatternFlowEthernetEtherTypeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowEthernetEtherTypeCounter
	Step() int32
	SetStep(value int32) PatternFlowEthernetEtherTypeCounter
	Count() int32
	SetCount(value int32) PatternFlowEthernetEtherTypeCounter
}

func (obj *patternFlowEthernetEtherTypeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowEthernetEtherTypeCounter) SetStart(value int32) PatternFlowEthernetEtherTypeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowEthernetEtherTypeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowEthernetEtherTypeCounter) SetStep(value int32) PatternFlowEthernetEtherTypeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowEthernetEtherTypeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowEthernetEtherTypeCounter) SetCount(value int32) PatternFlowEthernetEtherTypeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowEthernetPfcQueueCounter struct {
	obj *snappipb.PatternFlowEthernetPfcQueueCounter
}

func (obj *patternFlowEthernetPfcQueueCounter) msg() *snappipb.PatternFlowEthernetPfcQueueCounter {
	return obj.obj
}

func (obj *patternFlowEthernetPfcQueueCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetPfcQueueCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetPfcQueueCounter interface {
	msg() *snappipb.PatternFlowEthernetPfcQueueCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowEthernetPfcQueueCounter
	Step() int32
	SetStep(value int32) PatternFlowEthernetPfcQueueCounter
	Count() int32
	SetCount(value int32) PatternFlowEthernetPfcQueueCounter
}

func (obj *patternFlowEthernetPfcQueueCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowEthernetPfcQueueCounter) SetStart(value int32) PatternFlowEthernetPfcQueueCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowEthernetPfcQueueCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowEthernetPfcQueueCounter) SetStep(value int32) PatternFlowEthernetPfcQueueCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowEthernetPfcQueueCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowEthernetPfcQueueCounter) SetCount(value int32) PatternFlowEthernetPfcQueueCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowVlanPriorityCounter struct {
	obj *snappipb.PatternFlowVlanPriorityCounter
}

func (obj *patternFlowVlanPriorityCounter) msg() *snappipb.PatternFlowVlanPriorityCounter {
	return obj.obj
}

func (obj *patternFlowVlanPriorityCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVlanPriorityCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVlanPriorityCounter interface {
	msg() *snappipb.PatternFlowVlanPriorityCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowVlanPriorityCounter
	Step() int32
	SetStep(value int32) PatternFlowVlanPriorityCounter
	Count() int32
	SetCount(value int32) PatternFlowVlanPriorityCounter
}

func (obj *patternFlowVlanPriorityCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowVlanPriorityCounter) SetStart(value int32) PatternFlowVlanPriorityCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowVlanPriorityCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowVlanPriorityCounter) SetStep(value int32) PatternFlowVlanPriorityCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowVlanPriorityCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowVlanPriorityCounter) SetCount(value int32) PatternFlowVlanPriorityCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowVlanCfiCounter struct {
	obj *snappipb.PatternFlowVlanCfiCounter
}

func (obj *patternFlowVlanCfiCounter) msg() *snappipb.PatternFlowVlanCfiCounter {
	return obj.obj
}

func (obj *patternFlowVlanCfiCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVlanCfiCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVlanCfiCounter interface {
	msg() *snappipb.PatternFlowVlanCfiCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowVlanCfiCounter
	Step() int32
	SetStep(value int32) PatternFlowVlanCfiCounter
	Count() int32
	SetCount(value int32) PatternFlowVlanCfiCounter
}

func (obj *patternFlowVlanCfiCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowVlanCfiCounter) SetStart(value int32) PatternFlowVlanCfiCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowVlanCfiCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowVlanCfiCounter) SetStep(value int32) PatternFlowVlanCfiCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowVlanCfiCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowVlanCfiCounter) SetCount(value int32) PatternFlowVlanCfiCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowVlanIdCounter struct {
	obj *snappipb.PatternFlowVlanIdCounter
}

func (obj *patternFlowVlanIdCounter) msg() *snappipb.PatternFlowVlanIdCounter {
	return obj.obj
}

func (obj *patternFlowVlanIdCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVlanIdCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVlanIdCounter interface {
	msg() *snappipb.PatternFlowVlanIdCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowVlanIdCounter
	Step() int32
	SetStep(value int32) PatternFlowVlanIdCounter
	Count() int32
	SetCount(value int32) PatternFlowVlanIdCounter
}

func (obj *patternFlowVlanIdCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowVlanIdCounter) SetStart(value int32) PatternFlowVlanIdCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowVlanIdCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowVlanIdCounter) SetStep(value int32) PatternFlowVlanIdCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowVlanIdCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowVlanIdCounter) SetCount(value int32) PatternFlowVlanIdCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowVlanTpidCounter struct {
	obj *snappipb.PatternFlowVlanTpidCounter
}

func (obj *patternFlowVlanTpidCounter) msg() *snappipb.PatternFlowVlanTpidCounter {
	return obj.obj
}

func (obj *patternFlowVlanTpidCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVlanTpidCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVlanTpidCounter interface {
	msg() *snappipb.PatternFlowVlanTpidCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowVlanTpidCounter
	Step() int32
	SetStep(value int32) PatternFlowVlanTpidCounter
	Count() int32
	SetCount(value int32) PatternFlowVlanTpidCounter
}

func (obj *patternFlowVlanTpidCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowVlanTpidCounter) SetStart(value int32) PatternFlowVlanTpidCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowVlanTpidCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowVlanTpidCounter) SetStep(value int32) PatternFlowVlanTpidCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowVlanTpidCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowVlanTpidCounter) SetCount(value int32) PatternFlowVlanTpidCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowVxlanFlagsCounter struct {
	obj *snappipb.PatternFlowVxlanFlagsCounter
}

func (obj *patternFlowVxlanFlagsCounter) msg() *snappipb.PatternFlowVxlanFlagsCounter {
	return obj.obj
}

func (obj *patternFlowVxlanFlagsCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVxlanFlagsCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVxlanFlagsCounter interface {
	msg() *snappipb.PatternFlowVxlanFlagsCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowVxlanFlagsCounter
	Step() int32
	SetStep(value int32) PatternFlowVxlanFlagsCounter
	Count() int32
	SetCount(value int32) PatternFlowVxlanFlagsCounter
}

func (obj *patternFlowVxlanFlagsCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowVxlanFlagsCounter) SetStart(value int32) PatternFlowVxlanFlagsCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowVxlanFlagsCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowVxlanFlagsCounter) SetStep(value int32) PatternFlowVxlanFlagsCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowVxlanFlagsCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowVxlanFlagsCounter) SetCount(value int32) PatternFlowVxlanFlagsCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowVxlanReserved0Counter struct {
	obj *snappipb.PatternFlowVxlanReserved0Counter
}

func (obj *patternFlowVxlanReserved0Counter) msg() *snappipb.PatternFlowVxlanReserved0Counter {
	return obj.obj
}

func (obj *patternFlowVxlanReserved0Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVxlanReserved0Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVxlanReserved0Counter interface {
	msg() *snappipb.PatternFlowVxlanReserved0Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowVxlanReserved0Counter
	Step() int32
	SetStep(value int32) PatternFlowVxlanReserved0Counter
	Count() int32
	SetCount(value int32) PatternFlowVxlanReserved0Counter
}

func (obj *patternFlowVxlanReserved0Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowVxlanReserved0Counter) SetStart(value int32) PatternFlowVxlanReserved0Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowVxlanReserved0Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowVxlanReserved0Counter) SetStep(value int32) PatternFlowVxlanReserved0Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowVxlanReserved0Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowVxlanReserved0Counter) SetCount(value int32) PatternFlowVxlanReserved0Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowVxlanVniCounter struct {
	obj *snappipb.PatternFlowVxlanVniCounter
}

func (obj *patternFlowVxlanVniCounter) msg() *snappipb.PatternFlowVxlanVniCounter {
	return obj.obj
}

func (obj *patternFlowVxlanVniCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVxlanVniCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVxlanVniCounter interface {
	msg() *snappipb.PatternFlowVxlanVniCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowVxlanVniCounter
	Step() int32
	SetStep(value int32) PatternFlowVxlanVniCounter
	Count() int32
	SetCount(value int32) PatternFlowVxlanVniCounter
}

func (obj *patternFlowVxlanVniCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowVxlanVniCounter) SetStart(value int32) PatternFlowVxlanVniCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowVxlanVniCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowVxlanVniCounter) SetStep(value int32) PatternFlowVxlanVniCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowVxlanVniCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowVxlanVniCounter) SetCount(value int32) PatternFlowVxlanVniCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowVxlanReserved1Counter struct {
	obj *snappipb.PatternFlowVxlanReserved1Counter
}

func (obj *patternFlowVxlanReserved1Counter) msg() *snappipb.PatternFlowVxlanReserved1Counter {
	return obj.obj
}

func (obj *patternFlowVxlanReserved1Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowVxlanReserved1Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowVxlanReserved1Counter interface {
	msg() *snappipb.PatternFlowVxlanReserved1Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowVxlanReserved1Counter
	Step() int32
	SetStep(value int32) PatternFlowVxlanReserved1Counter
	Count() int32
	SetCount(value int32) PatternFlowVxlanReserved1Counter
}

func (obj *patternFlowVxlanReserved1Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowVxlanReserved1Counter) SetStart(value int32) PatternFlowVxlanReserved1Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowVxlanReserved1Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowVxlanReserved1Counter) SetStep(value int32) PatternFlowVxlanReserved1Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowVxlanReserved1Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowVxlanReserved1Counter) SetCount(value int32) PatternFlowVxlanReserved1Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4VersionCounter struct {
	obj *snappipb.PatternFlowIpv4VersionCounter
}

func (obj *patternFlowIpv4VersionCounter) msg() *snappipb.PatternFlowIpv4VersionCounter {
	return obj.obj
}

func (obj *patternFlowIpv4VersionCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4VersionCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4VersionCounter interface {
	msg() *snappipb.PatternFlowIpv4VersionCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4VersionCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4VersionCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4VersionCounter
}

func (obj *patternFlowIpv4VersionCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4VersionCounter) SetStart(value int32) PatternFlowIpv4VersionCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4VersionCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4VersionCounter) SetStep(value int32) PatternFlowIpv4VersionCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4VersionCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4VersionCounter) SetCount(value int32) PatternFlowIpv4VersionCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4HeaderLengthCounter struct {
	obj *snappipb.PatternFlowIpv4HeaderLengthCounter
}

func (obj *patternFlowIpv4HeaderLengthCounter) msg() *snappipb.PatternFlowIpv4HeaderLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv4HeaderLengthCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4HeaderLengthCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4HeaderLengthCounter interface {
	msg() *snappipb.PatternFlowIpv4HeaderLengthCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4HeaderLengthCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4HeaderLengthCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4HeaderLengthCounter
}

func (obj *patternFlowIpv4HeaderLengthCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4HeaderLengthCounter) SetStart(value int32) PatternFlowIpv4HeaderLengthCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4HeaderLengthCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4HeaderLengthCounter) SetStep(value int32) PatternFlowIpv4HeaderLengthCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4HeaderLengthCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4HeaderLengthCounter) SetCount(value int32) PatternFlowIpv4HeaderLengthCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4PriorityRaw struct {
	obj *snappipb.PatternFlowIpv4PriorityRaw
}

func (obj *patternFlowIpv4PriorityRaw) msg() *snappipb.PatternFlowIpv4PriorityRaw {
	return obj.obj
}

func (obj *patternFlowIpv4PriorityRaw) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4PriorityRaw) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4PriorityRaw interface {
	msg() *snappipb.PatternFlowIpv4PriorityRaw
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4PriorityRaw
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4PriorityRaw
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4PriorityRaw
	Increment() PatternFlowIpv4PriorityRawCounter
	Decrement() PatternFlowIpv4PriorityRawCounter
}

func (obj *patternFlowIpv4PriorityRaw) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4PriorityRaw) SetValue(value int32) PatternFlowIpv4PriorityRaw {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4PriorityRaw) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4PriorityRaw) SetValues(value []int32) PatternFlowIpv4PriorityRaw {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4PriorityRaw) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4PriorityRaw) SetMetricGroup(value string) PatternFlowIpv4PriorityRaw {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4PriorityRaw) Increment() PatternFlowIpv4PriorityRawCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4PriorityRawCounter{}
	}
	return &patternFlowIpv4PriorityRawCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4PriorityRaw) Decrement() PatternFlowIpv4PriorityRawCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4PriorityRawCounter{}
	}
	return &patternFlowIpv4PriorityRawCounter{obj: obj.obj.Decrement}

}

type flowIpv4Tos struct {
	obj *snappipb.FlowIpv4Tos
}

func (obj *flowIpv4Tos) msg() *snappipb.FlowIpv4Tos {
	return obj.obj
}

func (obj *flowIpv4Tos) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowIpv4Tos) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowIpv4Tos interface {
	msg() *snappipb.FlowIpv4Tos
	Yaml() string
	Json() string
	Precedence() PatternFlowIpv4TosPrecedence
	Delay() PatternFlowIpv4TosDelay
	Throughput() PatternFlowIpv4TosThroughput
	Reliability() PatternFlowIpv4TosReliability
	Monetary() PatternFlowIpv4TosMonetary
	Unused() PatternFlowIpv4TosUnused
}

func (obj *flowIpv4Tos) Precedence() PatternFlowIpv4TosPrecedence {
	if obj.obj.Precedence == nil {
		obj.obj.Precedence = &snappipb.PatternFlowIpv4TosPrecedence{}
	}
	return &patternFlowIpv4TosPrecedence{obj: obj.obj.Precedence}

}

func (obj *flowIpv4Tos) Delay() PatternFlowIpv4TosDelay {
	if obj.obj.Delay == nil {
		obj.obj.Delay = &snappipb.PatternFlowIpv4TosDelay{}
	}
	return &patternFlowIpv4TosDelay{obj: obj.obj.Delay}

}

func (obj *flowIpv4Tos) Throughput() PatternFlowIpv4TosThroughput {
	if obj.obj.Throughput == nil {
		obj.obj.Throughput = &snappipb.PatternFlowIpv4TosThroughput{}
	}
	return &patternFlowIpv4TosThroughput{obj: obj.obj.Throughput}

}

func (obj *flowIpv4Tos) Reliability() PatternFlowIpv4TosReliability {
	if obj.obj.Reliability == nil {
		obj.obj.Reliability = &snappipb.PatternFlowIpv4TosReliability{}
	}
	return &patternFlowIpv4TosReliability{obj: obj.obj.Reliability}

}

func (obj *flowIpv4Tos) Monetary() PatternFlowIpv4TosMonetary {
	if obj.obj.Monetary == nil {
		obj.obj.Monetary = &snappipb.PatternFlowIpv4TosMonetary{}
	}
	return &patternFlowIpv4TosMonetary{obj: obj.obj.Monetary}

}

func (obj *flowIpv4Tos) Unused() PatternFlowIpv4TosUnused {
	if obj.obj.Unused == nil {
		obj.obj.Unused = &snappipb.PatternFlowIpv4TosUnused{}
	}
	return &patternFlowIpv4TosUnused{obj: obj.obj.Unused}

}

type flowIpv4Dscp struct {
	obj *snappipb.FlowIpv4Dscp
}

func (obj *flowIpv4Dscp) msg() *snappipb.FlowIpv4Dscp {
	return obj.obj
}

func (obj *flowIpv4Dscp) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *flowIpv4Dscp) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FlowIpv4Dscp interface {
	msg() *snappipb.FlowIpv4Dscp
	Yaml() string
	Json() string
	Phb() PatternFlowIpv4DscpPhb
	Ecn() PatternFlowIpv4DscpEcn
}

func (obj *flowIpv4Dscp) Phb() PatternFlowIpv4DscpPhb {
	if obj.obj.Phb == nil {
		obj.obj.Phb = &snappipb.PatternFlowIpv4DscpPhb{}
	}
	return &patternFlowIpv4DscpPhb{obj: obj.obj.Phb}

}

func (obj *flowIpv4Dscp) Ecn() PatternFlowIpv4DscpEcn {
	if obj.obj.Ecn == nil {
		obj.obj.Ecn = &snappipb.PatternFlowIpv4DscpEcn{}
	}
	return &patternFlowIpv4DscpEcn{obj: obj.obj.Ecn}

}

type patternFlowIpv4TotalLengthCounter struct {
	obj *snappipb.PatternFlowIpv4TotalLengthCounter
}

func (obj *patternFlowIpv4TotalLengthCounter) msg() *snappipb.PatternFlowIpv4TotalLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TotalLengthCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TotalLengthCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TotalLengthCounter interface {
	msg() *snappipb.PatternFlowIpv4TotalLengthCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4TotalLengthCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4TotalLengthCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4TotalLengthCounter
}

func (obj *patternFlowIpv4TotalLengthCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4TotalLengthCounter) SetStart(value int32) PatternFlowIpv4TotalLengthCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4TotalLengthCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4TotalLengthCounter) SetStep(value int32) PatternFlowIpv4TotalLengthCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4TotalLengthCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4TotalLengthCounter) SetCount(value int32) PatternFlowIpv4TotalLengthCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4IdentificationCounter struct {
	obj *snappipb.PatternFlowIpv4IdentificationCounter
}

func (obj *patternFlowIpv4IdentificationCounter) msg() *snappipb.PatternFlowIpv4IdentificationCounter {
	return obj.obj
}

func (obj *patternFlowIpv4IdentificationCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4IdentificationCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4IdentificationCounter interface {
	msg() *snappipb.PatternFlowIpv4IdentificationCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4IdentificationCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4IdentificationCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4IdentificationCounter
}

func (obj *patternFlowIpv4IdentificationCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4IdentificationCounter) SetStart(value int32) PatternFlowIpv4IdentificationCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4IdentificationCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4IdentificationCounter) SetStep(value int32) PatternFlowIpv4IdentificationCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4IdentificationCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4IdentificationCounter) SetCount(value int32) PatternFlowIpv4IdentificationCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4ReservedCounter struct {
	obj *snappipb.PatternFlowIpv4ReservedCounter
}

func (obj *patternFlowIpv4ReservedCounter) msg() *snappipb.PatternFlowIpv4ReservedCounter {
	return obj.obj
}

func (obj *patternFlowIpv4ReservedCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4ReservedCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4ReservedCounter interface {
	msg() *snappipb.PatternFlowIpv4ReservedCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4ReservedCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4ReservedCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4ReservedCounter
}

func (obj *patternFlowIpv4ReservedCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4ReservedCounter) SetStart(value int32) PatternFlowIpv4ReservedCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4ReservedCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4ReservedCounter) SetStep(value int32) PatternFlowIpv4ReservedCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4ReservedCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4ReservedCounter) SetCount(value int32) PatternFlowIpv4ReservedCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4DontFragmentCounter struct {
	obj *snappipb.PatternFlowIpv4DontFragmentCounter
}

func (obj *patternFlowIpv4DontFragmentCounter) msg() *snappipb.PatternFlowIpv4DontFragmentCounter {
	return obj.obj
}

func (obj *patternFlowIpv4DontFragmentCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4DontFragmentCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4DontFragmentCounter interface {
	msg() *snappipb.PatternFlowIpv4DontFragmentCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4DontFragmentCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4DontFragmentCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4DontFragmentCounter
}

func (obj *patternFlowIpv4DontFragmentCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4DontFragmentCounter) SetStart(value int32) PatternFlowIpv4DontFragmentCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4DontFragmentCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4DontFragmentCounter) SetStep(value int32) PatternFlowIpv4DontFragmentCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4DontFragmentCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4DontFragmentCounter) SetCount(value int32) PatternFlowIpv4DontFragmentCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4MoreFragmentsCounter struct {
	obj *snappipb.PatternFlowIpv4MoreFragmentsCounter
}

func (obj *patternFlowIpv4MoreFragmentsCounter) msg() *snappipb.PatternFlowIpv4MoreFragmentsCounter {
	return obj.obj
}

func (obj *patternFlowIpv4MoreFragmentsCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4MoreFragmentsCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4MoreFragmentsCounter interface {
	msg() *snappipb.PatternFlowIpv4MoreFragmentsCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4MoreFragmentsCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4MoreFragmentsCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4MoreFragmentsCounter
}

func (obj *patternFlowIpv4MoreFragmentsCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4MoreFragmentsCounter) SetStart(value int32) PatternFlowIpv4MoreFragmentsCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4MoreFragmentsCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4MoreFragmentsCounter) SetStep(value int32) PatternFlowIpv4MoreFragmentsCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4MoreFragmentsCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4MoreFragmentsCounter) SetCount(value int32) PatternFlowIpv4MoreFragmentsCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4FragmentOffsetCounter struct {
	obj *snappipb.PatternFlowIpv4FragmentOffsetCounter
}

func (obj *patternFlowIpv4FragmentOffsetCounter) msg() *snappipb.PatternFlowIpv4FragmentOffsetCounter {
	return obj.obj
}

func (obj *patternFlowIpv4FragmentOffsetCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4FragmentOffsetCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4FragmentOffsetCounter interface {
	msg() *snappipb.PatternFlowIpv4FragmentOffsetCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4FragmentOffsetCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4FragmentOffsetCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4FragmentOffsetCounter
}

func (obj *patternFlowIpv4FragmentOffsetCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4FragmentOffsetCounter) SetStart(value int32) PatternFlowIpv4FragmentOffsetCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4FragmentOffsetCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4FragmentOffsetCounter) SetStep(value int32) PatternFlowIpv4FragmentOffsetCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4FragmentOffsetCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4FragmentOffsetCounter) SetCount(value int32) PatternFlowIpv4FragmentOffsetCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4TimeToLiveCounter struct {
	obj *snappipb.PatternFlowIpv4TimeToLiveCounter
}

func (obj *patternFlowIpv4TimeToLiveCounter) msg() *snappipb.PatternFlowIpv4TimeToLiveCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TimeToLiveCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TimeToLiveCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TimeToLiveCounter interface {
	msg() *snappipb.PatternFlowIpv4TimeToLiveCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4TimeToLiveCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4TimeToLiveCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4TimeToLiveCounter
}

func (obj *patternFlowIpv4TimeToLiveCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4TimeToLiveCounter) SetStart(value int32) PatternFlowIpv4TimeToLiveCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4TimeToLiveCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4TimeToLiveCounter) SetStep(value int32) PatternFlowIpv4TimeToLiveCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4TimeToLiveCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4TimeToLiveCounter) SetCount(value int32) PatternFlowIpv4TimeToLiveCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4ProtocolCounter struct {
	obj *snappipb.PatternFlowIpv4ProtocolCounter
}

func (obj *patternFlowIpv4ProtocolCounter) msg() *snappipb.PatternFlowIpv4ProtocolCounter {
	return obj.obj
}

func (obj *patternFlowIpv4ProtocolCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4ProtocolCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4ProtocolCounter interface {
	msg() *snappipb.PatternFlowIpv4ProtocolCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4ProtocolCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4ProtocolCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4ProtocolCounter
}

func (obj *patternFlowIpv4ProtocolCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4ProtocolCounter) SetStart(value int32) PatternFlowIpv4ProtocolCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4ProtocolCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4ProtocolCounter) SetStep(value int32) PatternFlowIpv4ProtocolCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4ProtocolCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4ProtocolCounter) SetCount(value int32) PatternFlowIpv4ProtocolCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4SrcCounter struct {
	obj *snappipb.PatternFlowIpv4SrcCounter
}

func (obj *patternFlowIpv4SrcCounter) msg() *snappipb.PatternFlowIpv4SrcCounter {
	return obj.obj
}

func (obj *patternFlowIpv4SrcCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4SrcCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4SrcCounter interface {
	msg() *snappipb.PatternFlowIpv4SrcCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowIpv4SrcCounter
	Step() string
	SetStep(value string) PatternFlowIpv4SrcCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4SrcCounter
}

func (obj *patternFlowIpv4SrcCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4SrcCounter) SetStart(value string) PatternFlowIpv4SrcCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4SrcCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4SrcCounter) SetStep(value string) PatternFlowIpv4SrcCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4SrcCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4SrcCounter) SetCount(value int32) PatternFlowIpv4SrcCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4DstCounter struct {
	obj *snappipb.PatternFlowIpv4DstCounter
}

func (obj *patternFlowIpv4DstCounter) msg() *snappipb.PatternFlowIpv4DstCounter {
	return obj.obj
}

func (obj *patternFlowIpv4DstCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4DstCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4DstCounter interface {
	msg() *snappipb.PatternFlowIpv4DstCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowIpv4DstCounter
	Step() string
	SetStep(value string) PatternFlowIpv4DstCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4DstCounter
}

func (obj *patternFlowIpv4DstCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4DstCounter) SetStart(value string) PatternFlowIpv4DstCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4DstCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4DstCounter) SetStep(value string) PatternFlowIpv4DstCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4DstCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4DstCounter) SetCount(value int32) PatternFlowIpv4DstCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv6VersionCounter struct {
	obj *snappipb.PatternFlowIpv6VersionCounter
}

func (obj *patternFlowIpv6VersionCounter) msg() *snappipb.PatternFlowIpv6VersionCounter {
	return obj.obj
}

func (obj *patternFlowIpv6VersionCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6VersionCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6VersionCounter interface {
	msg() *snappipb.PatternFlowIpv6VersionCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv6VersionCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv6VersionCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv6VersionCounter
}

func (obj *patternFlowIpv6VersionCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv6VersionCounter) SetStart(value int32) PatternFlowIpv6VersionCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv6VersionCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv6VersionCounter) SetStep(value int32) PatternFlowIpv6VersionCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv6VersionCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv6VersionCounter) SetCount(value int32) PatternFlowIpv6VersionCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv6TrafficClassCounter struct {
	obj *snappipb.PatternFlowIpv6TrafficClassCounter
}

func (obj *patternFlowIpv6TrafficClassCounter) msg() *snappipb.PatternFlowIpv6TrafficClassCounter {
	return obj.obj
}

func (obj *patternFlowIpv6TrafficClassCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6TrafficClassCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6TrafficClassCounter interface {
	msg() *snappipb.PatternFlowIpv6TrafficClassCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv6TrafficClassCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv6TrafficClassCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv6TrafficClassCounter
}

func (obj *patternFlowIpv6TrafficClassCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv6TrafficClassCounter) SetStart(value int32) PatternFlowIpv6TrafficClassCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv6TrafficClassCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv6TrafficClassCounter) SetStep(value int32) PatternFlowIpv6TrafficClassCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv6TrafficClassCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv6TrafficClassCounter) SetCount(value int32) PatternFlowIpv6TrafficClassCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv6FlowLabelCounter struct {
	obj *snappipb.PatternFlowIpv6FlowLabelCounter
}

func (obj *patternFlowIpv6FlowLabelCounter) msg() *snappipb.PatternFlowIpv6FlowLabelCounter {
	return obj.obj
}

func (obj *patternFlowIpv6FlowLabelCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6FlowLabelCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6FlowLabelCounter interface {
	msg() *snappipb.PatternFlowIpv6FlowLabelCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv6FlowLabelCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv6FlowLabelCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv6FlowLabelCounter
}

func (obj *patternFlowIpv6FlowLabelCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv6FlowLabelCounter) SetStart(value int32) PatternFlowIpv6FlowLabelCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv6FlowLabelCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv6FlowLabelCounter) SetStep(value int32) PatternFlowIpv6FlowLabelCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv6FlowLabelCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv6FlowLabelCounter) SetCount(value int32) PatternFlowIpv6FlowLabelCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv6PayloadLengthCounter struct {
	obj *snappipb.PatternFlowIpv6PayloadLengthCounter
}

func (obj *patternFlowIpv6PayloadLengthCounter) msg() *snappipb.PatternFlowIpv6PayloadLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv6PayloadLengthCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6PayloadLengthCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6PayloadLengthCounter interface {
	msg() *snappipb.PatternFlowIpv6PayloadLengthCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv6PayloadLengthCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv6PayloadLengthCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv6PayloadLengthCounter
}

func (obj *patternFlowIpv6PayloadLengthCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv6PayloadLengthCounter) SetStart(value int32) PatternFlowIpv6PayloadLengthCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv6PayloadLengthCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv6PayloadLengthCounter) SetStep(value int32) PatternFlowIpv6PayloadLengthCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv6PayloadLengthCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv6PayloadLengthCounter) SetCount(value int32) PatternFlowIpv6PayloadLengthCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv6NextHeaderCounter struct {
	obj *snappipb.PatternFlowIpv6NextHeaderCounter
}

func (obj *patternFlowIpv6NextHeaderCounter) msg() *snappipb.PatternFlowIpv6NextHeaderCounter {
	return obj.obj
}

func (obj *patternFlowIpv6NextHeaderCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6NextHeaderCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6NextHeaderCounter interface {
	msg() *snappipb.PatternFlowIpv6NextHeaderCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv6NextHeaderCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv6NextHeaderCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv6NextHeaderCounter
}

func (obj *patternFlowIpv6NextHeaderCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv6NextHeaderCounter) SetStart(value int32) PatternFlowIpv6NextHeaderCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv6NextHeaderCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv6NextHeaderCounter) SetStep(value int32) PatternFlowIpv6NextHeaderCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv6NextHeaderCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv6NextHeaderCounter) SetCount(value int32) PatternFlowIpv6NextHeaderCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv6HopLimitCounter struct {
	obj *snappipb.PatternFlowIpv6HopLimitCounter
}

func (obj *patternFlowIpv6HopLimitCounter) msg() *snappipb.PatternFlowIpv6HopLimitCounter {
	return obj.obj
}

func (obj *patternFlowIpv6HopLimitCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6HopLimitCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6HopLimitCounter interface {
	msg() *snappipb.PatternFlowIpv6HopLimitCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv6HopLimitCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv6HopLimitCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv6HopLimitCounter
}

func (obj *patternFlowIpv6HopLimitCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv6HopLimitCounter) SetStart(value int32) PatternFlowIpv6HopLimitCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv6HopLimitCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv6HopLimitCounter) SetStep(value int32) PatternFlowIpv6HopLimitCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv6HopLimitCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv6HopLimitCounter) SetCount(value int32) PatternFlowIpv6HopLimitCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv6SrcCounter struct {
	obj *snappipb.PatternFlowIpv6SrcCounter
}

func (obj *patternFlowIpv6SrcCounter) msg() *snappipb.PatternFlowIpv6SrcCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SrcCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6SrcCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6SrcCounter interface {
	msg() *snappipb.PatternFlowIpv6SrcCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowIpv6SrcCounter
	Step() string
	SetStep(value string) PatternFlowIpv6SrcCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv6SrcCounter
}

func (obj *patternFlowIpv6SrcCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowIpv6SrcCounter) SetStart(value string) PatternFlowIpv6SrcCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv6SrcCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowIpv6SrcCounter) SetStep(value string) PatternFlowIpv6SrcCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv6SrcCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv6SrcCounter) SetCount(value int32) PatternFlowIpv6SrcCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv6DstCounter struct {
	obj *snappipb.PatternFlowIpv6DstCounter
}

func (obj *patternFlowIpv6DstCounter) msg() *snappipb.PatternFlowIpv6DstCounter {
	return obj.obj
}

func (obj *patternFlowIpv6DstCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv6DstCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv6DstCounter interface {
	msg() *snappipb.PatternFlowIpv6DstCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowIpv6DstCounter
	Step() string
	SetStep(value string) PatternFlowIpv6DstCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv6DstCounter
}

func (obj *patternFlowIpv6DstCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowIpv6DstCounter) SetStart(value string) PatternFlowIpv6DstCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv6DstCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowIpv6DstCounter) SetStep(value string) PatternFlowIpv6DstCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv6DstCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv6DstCounter) SetCount(value int32) PatternFlowIpv6DstCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPfcPauseDstCounter struct {
	obj *snappipb.PatternFlowPfcPauseDstCounter
}

func (obj *patternFlowPfcPauseDstCounter) msg() *snappipb.PatternFlowPfcPauseDstCounter {
	return obj.obj
}

func (obj *patternFlowPfcPauseDstCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPauseDstCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPauseDstCounter interface {
	msg() *snappipb.PatternFlowPfcPauseDstCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowPfcPauseDstCounter
	Step() string
	SetStep(value string) PatternFlowPfcPauseDstCounter
	Count() int32
	SetCount(value int32) PatternFlowPfcPauseDstCounter
}

func (obj *patternFlowPfcPauseDstCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowPfcPauseDstCounter) SetStart(value string) PatternFlowPfcPauseDstCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPfcPauseDstCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowPfcPauseDstCounter) SetStep(value string) PatternFlowPfcPauseDstCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPfcPauseDstCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPfcPauseDstCounter) SetCount(value int32) PatternFlowPfcPauseDstCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPfcPauseSrcCounter struct {
	obj *snappipb.PatternFlowPfcPauseSrcCounter
}

func (obj *patternFlowPfcPauseSrcCounter) msg() *snappipb.PatternFlowPfcPauseSrcCounter {
	return obj.obj
}

func (obj *patternFlowPfcPauseSrcCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPauseSrcCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPauseSrcCounter interface {
	msg() *snappipb.PatternFlowPfcPauseSrcCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowPfcPauseSrcCounter
	Step() string
	SetStep(value string) PatternFlowPfcPauseSrcCounter
	Count() int32
	SetCount(value int32) PatternFlowPfcPauseSrcCounter
}

func (obj *patternFlowPfcPauseSrcCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowPfcPauseSrcCounter) SetStart(value string) PatternFlowPfcPauseSrcCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPfcPauseSrcCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowPfcPauseSrcCounter) SetStep(value string) PatternFlowPfcPauseSrcCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPfcPauseSrcCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPfcPauseSrcCounter) SetCount(value int32) PatternFlowPfcPauseSrcCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPfcPauseEtherTypeCounter struct {
	obj *snappipb.PatternFlowPfcPauseEtherTypeCounter
}

func (obj *patternFlowPfcPauseEtherTypeCounter) msg() *snappipb.PatternFlowPfcPauseEtherTypeCounter {
	return obj.obj
}

func (obj *patternFlowPfcPauseEtherTypeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPauseEtherTypeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPauseEtherTypeCounter interface {
	msg() *snappipb.PatternFlowPfcPauseEtherTypeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPfcPauseEtherTypeCounter
	Step() int32
	SetStep(value int32) PatternFlowPfcPauseEtherTypeCounter
	Count() int32
	SetCount(value int32) PatternFlowPfcPauseEtherTypeCounter
}

func (obj *patternFlowPfcPauseEtherTypeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPfcPauseEtherTypeCounter) SetStart(value int32) PatternFlowPfcPauseEtherTypeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPfcPauseEtherTypeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPfcPauseEtherTypeCounter) SetStep(value int32) PatternFlowPfcPauseEtherTypeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPfcPauseEtherTypeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPfcPauseEtherTypeCounter) SetCount(value int32) PatternFlowPfcPauseEtherTypeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPfcPauseControlOpCodeCounter struct {
	obj *snappipb.PatternFlowPfcPauseControlOpCodeCounter
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) msg() *snappipb.PatternFlowPfcPauseControlOpCodeCounter {
	return obj.obj
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPauseControlOpCodeCounter interface {
	msg() *snappipb.PatternFlowPfcPauseControlOpCodeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPfcPauseControlOpCodeCounter
	Step() int32
	SetStep(value int32) PatternFlowPfcPauseControlOpCodeCounter
	Count() int32
	SetCount(value int32) PatternFlowPfcPauseControlOpCodeCounter
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) SetStart(value int32) PatternFlowPfcPauseControlOpCodeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) SetStep(value int32) PatternFlowPfcPauseControlOpCodeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPfcPauseControlOpCodeCounter) SetCount(value int32) PatternFlowPfcPauseControlOpCodeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPfcPauseClassEnableVectorCounter struct {
	obj *snappipb.PatternFlowPfcPauseClassEnableVectorCounter
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) msg() *snappipb.PatternFlowPfcPauseClassEnableVectorCounter {
	return obj.obj
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPauseClassEnableVectorCounter interface {
	msg() *snappipb.PatternFlowPfcPauseClassEnableVectorCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPfcPauseClassEnableVectorCounter
	Step() int32
	SetStep(value int32) PatternFlowPfcPauseClassEnableVectorCounter
	Count() int32
	SetCount(value int32) PatternFlowPfcPauseClassEnableVectorCounter
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) SetStart(value int32) PatternFlowPfcPauseClassEnableVectorCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) SetStep(value int32) PatternFlowPfcPauseClassEnableVectorCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPfcPauseClassEnableVectorCounter) SetCount(value int32) PatternFlowPfcPauseClassEnableVectorCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPfcPausePauseClass0Counter struct {
	obj *snappipb.PatternFlowPfcPausePauseClass0Counter
}

func (obj *patternFlowPfcPausePauseClass0Counter) msg() *snappipb.PatternFlowPfcPausePauseClass0Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass0Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass0Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass0Counter interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass0Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPfcPausePauseClass0Counter
	Step() int32
	SetStep(value int32) PatternFlowPfcPausePauseClass0Counter
	Count() int32
	SetCount(value int32) PatternFlowPfcPausePauseClass0Counter
}

func (obj *patternFlowPfcPausePauseClass0Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPfcPausePauseClass0Counter) SetStart(value int32) PatternFlowPfcPausePauseClass0Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass0Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPfcPausePauseClass0Counter) SetStep(value int32) PatternFlowPfcPausePauseClass0Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass0Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPfcPausePauseClass0Counter) SetCount(value int32) PatternFlowPfcPausePauseClass0Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPfcPausePauseClass1Counter struct {
	obj *snappipb.PatternFlowPfcPausePauseClass1Counter
}

func (obj *patternFlowPfcPausePauseClass1Counter) msg() *snappipb.PatternFlowPfcPausePauseClass1Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass1Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass1Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass1Counter interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass1Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPfcPausePauseClass1Counter
	Step() int32
	SetStep(value int32) PatternFlowPfcPausePauseClass1Counter
	Count() int32
	SetCount(value int32) PatternFlowPfcPausePauseClass1Counter
}

func (obj *patternFlowPfcPausePauseClass1Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPfcPausePauseClass1Counter) SetStart(value int32) PatternFlowPfcPausePauseClass1Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass1Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPfcPausePauseClass1Counter) SetStep(value int32) PatternFlowPfcPausePauseClass1Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass1Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPfcPausePauseClass1Counter) SetCount(value int32) PatternFlowPfcPausePauseClass1Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPfcPausePauseClass2Counter struct {
	obj *snappipb.PatternFlowPfcPausePauseClass2Counter
}

func (obj *patternFlowPfcPausePauseClass2Counter) msg() *snappipb.PatternFlowPfcPausePauseClass2Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass2Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass2Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass2Counter interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass2Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPfcPausePauseClass2Counter
	Step() int32
	SetStep(value int32) PatternFlowPfcPausePauseClass2Counter
	Count() int32
	SetCount(value int32) PatternFlowPfcPausePauseClass2Counter
}

func (obj *patternFlowPfcPausePauseClass2Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPfcPausePauseClass2Counter) SetStart(value int32) PatternFlowPfcPausePauseClass2Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass2Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPfcPausePauseClass2Counter) SetStep(value int32) PatternFlowPfcPausePauseClass2Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass2Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPfcPausePauseClass2Counter) SetCount(value int32) PatternFlowPfcPausePauseClass2Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPfcPausePauseClass3Counter struct {
	obj *snappipb.PatternFlowPfcPausePauseClass3Counter
}

func (obj *patternFlowPfcPausePauseClass3Counter) msg() *snappipb.PatternFlowPfcPausePauseClass3Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass3Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass3Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass3Counter interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass3Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPfcPausePauseClass3Counter
	Step() int32
	SetStep(value int32) PatternFlowPfcPausePauseClass3Counter
	Count() int32
	SetCount(value int32) PatternFlowPfcPausePauseClass3Counter
}

func (obj *patternFlowPfcPausePauseClass3Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPfcPausePauseClass3Counter) SetStart(value int32) PatternFlowPfcPausePauseClass3Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass3Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPfcPausePauseClass3Counter) SetStep(value int32) PatternFlowPfcPausePauseClass3Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass3Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPfcPausePauseClass3Counter) SetCount(value int32) PatternFlowPfcPausePauseClass3Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPfcPausePauseClass4Counter struct {
	obj *snappipb.PatternFlowPfcPausePauseClass4Counter
}

func (obj *patternFlowPfcPausePauseClass4Counter) msg() *snappipb.PatternFlowPfcPausePauseClass4Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass4Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass4Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass4Counter interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass4Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPfcPausePauseClass4Counter
	Step() int32
	SetStep(value int32) PatternFlowPfcPausePauseClass4Counter
	Count() int32
	SetCount(value int32) PatternFlowPfcPausePauseClass4Counter
}

func (obj *patternFlowPfcPausePauseClass4Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPfcPausePauseClass4Counter) SetStart(value int32) PatternFlowPfcPausePauseClass4Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass4Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPfcPausePauseClass4Counter) SetStep(value int32) PatternFlowPfcPausePauseClass4Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass4Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPfcPausePauseClass4Counter) SetCount(value int32) PatternFlowPfcPausePauseClass4Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPfcPausePauseClass5Counter struct {
	obj *snappipb.PatternFlowPfcPausePauseClass5Counter
}

func (obj *patternFlowPfcPausePauseClass5Counter) msg() *snappipb.PatternFlowPfcPausePauseClass5Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass5Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass5Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass5Counter interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass5Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPfcPausePauseClass5Counter
	Step() int32
	SetStep(value int32) PatternFlowPfcPausePauseClass5Counter
	Count() int32
	SetCount(value int32) PatternFlowPfcPausePauseClass5Counter
}

func (obj *patternFlowPfcPausePauseClass5Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPfcPausePauseClass5Counter) SetStart(value int32) PatternFlowPfcPausePauseClass5Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass5Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPfcPausePauseClass5Counter) SetStep(value int32) PatternFlowPfcPausePauseClass5Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass5Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPfcPausePauseClass5Counter) SetCount(value int32) PatternFlowPfcPausePauseClass5Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPfcPausePauseClass6Counter struct {
	obj *snappipb.PatternFlowPfcPausePauseClass6Counter
}

func (obj *patternFlowPfcPausePauseClass6Counter) msg() *snappipb.PatternFlowPfcPausePauseClass6Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass6Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass6Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass6Counter interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass6Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPfcPausePauseClass6Counter
	Step() int32
	SetStep(value int32) PatternFlowPfcPausePauseClass6Counter
	Count() int32
	SetCount(value int32) PatternFlowPfcPausePauseClass6Counter
}

func (obj *patternFlowPfcPausePauseClass6Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPfcPausePauseClass6Counter) SetStart(value int32) PatternFlowPfcPausePauseClass6Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass6Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPfcPausePauseClass6Counter) SetStep(value int32) PatternFlowPfcPausePauseClass6Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass6Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPfcPausePauseClass6Counter) SetCount(value int32) PatternFlowPfcPausePauseClass6Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPfcPausePauseClass7Counter struct {
	obj *snappipb.PatternFlowPfcPausePauseClass7Counter
}

func (obj *patternFlowPfcPausePauseClass7Counter) msg() *snappipb.PatternFlowPfcPausePauseClass7Counter {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass7Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPfcPausePauseClass7Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPfcPausePauseClass7Counter interface {
	msg() *snappipb.PatternFlowPfcPausePauseClass7Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPfcPausePauseClass7Counter
	Step() int32
	SetStep(value int32) PatternFlowPfcPausePauseClass7Counter
	Count() int32
	SetCount(value int32) PatternFlowPfcPausePauseClass7Counter
}

func (obj *patternFlowPfcPausePauseClass7Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPfcPausePauseClass7Counter) SetStart(value int32) PatternFlowPfcPausePauseClass7Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass7Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPfcPausePauseClass7Counter) SetStep(value int32) PatternFlowPfcPausePauseClass7Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass7Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPfcPausePauseClass7Counter) SetCount(value int32) PatternFlowPfcPausePauseClass7Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowEthernetPauseDstCounter struct {
	obj *snappipb.PatternFlowEthernetPauseDstCounter
}

func (obj *patternFlowEthernetPauseDstCounter) msg() *snappipb.PatternFlowEthernetPauseDstCounter {
	return obj.obj
}

func (obj *patternFlowEthernetPauseDstCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetPauseDstCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetPauseDstCounter interface {
	msg() *snappipb.PatternFlowEthernetPauseDstCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowEthernetPauseDstCounter
	Step() string
	SetStep(value string) PatternFlowEthernetPauseDstCounter
	Count() int32
	SetCount(value int32) PatternFlowEthernetPauseDstCounter
}

func (obj *patternFlowEthernetPauseDstCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowEthernetPauseDstCounter) SetStart(value string) PatternFlowEthernetPauseDstCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowEthernetPauseDstCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowEthernetPauseDstCounter) SetStep(value string) PatternFlowEthernetPauseDstCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowEthernetPauseDstCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowEthernetPauseDstCounter) SetCount(value int32) PatternFlowEthernetPauseDstCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowEthernetPauseSrcCounter struct {
	obj *snappipb.PatternFlowEthernetPauseSrcCounter
}

func (obj *patternFlowEthernetPauseSrcCounter) msg() *snappipb.PatternFlowEthernetPauseSrcCounter {
	return obj.obj
}

func (obj *patternFlowEthernetPauseSrcCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetPauseSrcCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetPauseSrcCounter interface {
	msg() *snappipb.PatternFlowEthernetPauseSrcCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowEthernetPauseSrcCounter
	Step() string
	SetStep(value string) PatternFlowEthernetPauseSrcCounter
	Count() int32
	SetCount(value int32) PatternFlowEthernetPauseSrcCounter
}

func (obj *patternFlowEthernetPauseSrcCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowEthernetPauseSrcCounter) SetStart(value string) PatternFlowEthernetPauseSrcCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowEthernetPauseSrcCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowEthernetPauseSrcCounter) SetStep(value string) PatternFlowEthernetPauseSrcCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowEthernetPauseSrcCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowEthernetPauseSrcCounter) SetCount(value int32) PatternFlowEthernetPauseSrcCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowEthernetPauseEtherTypeCounter struct {
	obj *snappipb.PatternFlowEthernetPauseEtherTypeCounter
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) msg() *snappipb.PatternFlowEthernetPauseEtherTypeCounter {
	return obj.obj
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetPauseEtherTypeCounter interface {
	msg() *snappipb.PatternFlowEthernetPauseEtherTypeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowEthernetPauseEtherTypeCounter
	Step() int32
	SetStep(value int32) PatternFlowEthernetPauseEtherTypeCounter
	Count() int32
	SetCount(value int32) PatternFlowEthernetPauseEtherTypeCounter
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) SetStart(value int32) PatternFlowEthernetPauseEtherTypeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) SetStep(value int32) PatternFlowEthernetPauseEtherTypeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowEthernetPauseEtherTypeCounter) SetCount(value int32) PatternFlowEthernetPauseEtherTypeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowEthernetPauseControlOpCodeCounter struct {
	obj *snappipb.PatternFlowEthernetPauseControlOpCodeCounter
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) msg() *snappipb.PatternFlowEthernetPauseControlOpCodeCounter {
	return obj.obj
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetPauseControlOpCodeCounter interface {
	msg() *snappipb.PatternFlowEthernetPauseControlOpCodeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowEthernetPauseControlOpCodeCounter
	Step() int32
	SetStep(value int32) PatternFlowEthernetPauseControlOpCodeCounter
	Count() int32
	SetCount(value int32) PatternFlowEthernetPauseControlOpCodeCounter
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) SetStart(value int32) PatternFlowEthernetPauseControlOpCodeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) SetStep(value int32) PatternFlowEthernetPauseControlOpCodeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowEthernetPauseControlOpCodeCounter) SetCount(value int32) PatternFlowEthernetPauseControlOpCodeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowEthernetPauseTimeCounter struct {
	obj *snappipb.PatternFlowEthernetPauseTimeCounter
}

func (obj *patternFlowEthernetPauseTimeCounter) msg() *snappipb.PatternFlowEthernetPauseTimeCounter {
	return obj.obj
}

func (obj *patternFlowEthernetPauseTimeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowEthernetPauseTimeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowEthernetPauseTimeCounter interface {
	msg() *snappipb.PatternFlowEthernetPauseTimeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowEthernetPauseTimeCounter
	Step() int32
	SetStep(value int32) PatternFlowEthernetPauseTimeCounter
	Count() int32
	SetCount(value int32) PatternFlowEthernetPauseTimeCounter
}

func (obj *patternFlowEthernetPauseTimeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowEthernetPauseTimeCounter) SetStart(value int32) PatternFlowEthernetPauseTimeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowEthernetPauseTimeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowEthernetPauseTimeCounter) SetStep(value int32) PatternFlowEthernetPauseTimeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowEthernetPauseTimeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowEthernetPauseTimeCounter) SetCount(value int32) PatternFlowEthernetPauseTimeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpSrcPortCounter struct {
	obj *snappipb.PatternFlowTcpSrcPortCounter
}

func (obj *patternFlowTcpSrcPortCounter) msg() *snappipb.PatternFlowTcpSrcPortCounter {
	return obj.obj
}

func (obj *patternFlowTcpSrcPortCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpSrcPortCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpSrcPortCounter interface {
	msg() *snappipb.PatternFlowTcpSrcPortCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpSrcPortCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpSrcPortCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpSrcPortCounter
}

func (obj *patternFlowTcpSrcPortCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpSrcPortCounter) SetStart(value int32) PatternFlowTcpSrcPortCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpSrcPortCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpSrcPortCounter) SetStep(value int32) PatternFlowTcpSrcPortCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpSrcPortCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpSrcPortCounter) SetCount(value int32) PatternFlowTcpSrcPortCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpDstPortCounter struct {
	obj *snappipb.PatternFlowTcpDstPortCounter
}

func (obj *patternFlowTcpDstPortCounter) msg() *snappipb.PatternFlowTcpDstPortCounter {
	return obj.obj
}

func (obj *patternFlowTcpDstPortCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpDstPortCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpDstPortCounter interface {
	msg() *snappipb.PatternFlowTcpDstPortCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpDstPortCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpDstPortCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpDstPortCounter
}

func (obj *patternFlowTcpDstPortCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpDstPortCounter) SetStart(value int32) PatternFlowTcpDstPortCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpDstPortCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpDstPortCounter) SetStep(value int32) PatternFlowTcpDstPortCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpDstPortCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpDstPortCounter) SetCount(value int32) PatternFlowTcpDstPortCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpSeqNumCounter struct {
	obj *snappipb.PatternFlowTcpSeqNumCounter
}

func (obj *patternFlowTcpSeqNumCounter) msg() *snappipb.PatternFlowTcpSeqNumCounter {
	return obj.obj
}

func (obj *patternFlowTcpSeqNumCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpSeqNumCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpSeqNumCounter interface {
	msg() *snappipb.PatternFlowTcpSeqNumCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpSeqNumCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpSeqNumCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpSeqNumCounter
}

func (obj *patternFlowTcpSeqNumCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpSeqNumCounter) SetStart(value int32) PatternFlowTcpSeqNumCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpSeqNumCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpSeqNumCounter) SetStep(value int32) PatternFlowTcpSeqNumCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpSeqNumCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpSeqNumCounter) SetCount(value int32) PatternFlowTcpSeqNumCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpAckNumCounter struct {
	obj *snappipb.PatternFlowTcpAckNumCounter
}

func (obj *patternFlowTcpAckNumCounter) msg() *snappipb.PatternFlowTcpAckNumCounter {
	return obj.obj
}

func (obj *patternFlowTcpAckNumCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpAckNumCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpAckNumCounter interface {
	msg() *snappipb.PatternFlowTcpAckNumCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpAckNumCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpAckNumCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpAckNumCounter
}

func (obj *patternFlowTcpAckNumCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpAckNumCounter) SetStart(value int32) PatternFlowTcpAckNumCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpAckNumCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpAckNumCounter) SetStep(value int32) PatternFlowTcpAckNumCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpAckNumCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpAckNumCounter) SetCount(value int32) PatternFlowTcpAckNumCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpDataOffsetCounter struct {
	obj *snappipb.PatternFlowTcpDataOffsetCounter
}

func (obj *patternFlowTcpDataOffsetCounter) msg() *snappipb.PatternFlowTcpDataOffsetCounter {
	return obj.obj
}

func (obj *patternFlowTcpDataOffsetCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpDataOffsetCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpDataOffsetCounter interface {
	msg() *snappipb.PatternFlowTcpDataOffsetCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpDataOffsetCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpDataOffsetCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpDataOffsetCounter
}

func (obj *patternFlowTcpDataOffsetCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpDataOffsetCounter) SetStart(value int32) PatternFlowTcpDataOffsetCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpDataOffsetCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpDataOffsetCounter) SetStep(value int32) PatternFlowTcpDataOffsetCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpDataOffsetCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpDataOffsetCounter) SetCount(value int32) PatternFlowTcpDataOffsetCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpEcnNsCounter struct {
	obj *snappipb.PatternFlowTcpEcnNsCounter
}

func (obj *patternFlowTcpEcnNsCounter) msg() *snappipb.PatternFlowTcpEcnNsCounter {
	return obj.obj
}

func (obj *patternFlowTcpEcnNsCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpEcnNsCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpEcnNsCounter interface {
	msg() *snappipb.PatternFlowTcpEcnNsCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpEcnNsCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpEcnNsCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpEcnNsCounter
}

func (obj *patternFlowTcpEcnNsCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpEcnNsCounter) SetStart(value int32) PatternFlowTcpEcnNsCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpEcnNsCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpEcnNsCounter) SetStep(value int32) PatternFlowTcpEcnNsCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpEcnNsCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpEcnNsCounter) SetCount(value int32) PatternFlowTcpEcnNsCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpEcnCwrCounter struct {
	obj *snappipb.PatternFlowTcpEcnCwrCounter
}

func (obj *patternFlowTcpEcnCwrCounter) msg() *snappipb.PatternFlowTcpEcnCwrCounter {
	return obj.obj
}

func (obj *patternFlowTcpEcnCwrCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpEcnCwrCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpEcnCwrCounter interface {
	msg() *snappipb.PatternFlowTcpEcnCwrCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpEcnCwrCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpEcnCwrCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpEcnCwrCounter
}

func (obj *patternFlowTcpEcnCwrCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpEcnCwrCounter) SetStart(value int32) PatternFlowTcpEcnCwrCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpEcnCwrCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpEcnCwrCounter) SetStep(value int32) PatternFlowTcpEcnCwrCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpEcnCwrCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpEcnCwrCounter) SetCount(value int32) PatternFlowTcpEcnCwrCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpEcnEchoCounter struct {
	obj *snappipb.PatternFlowTcpEcnEchoCounter
}

func (obj *patternFlowTcpEcnEchoCounter) msg() *snappipb.PatternFlowTcpEcnEchoCounter {
	return obj.obj
}

func (obj *patternFlowTcpEcnEchoCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpEcnEchoCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpEcnEchoCounter interface {
	msg() *snappipb.PatternFlowTcpEcnEchoCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpEcnEchoCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpEcnEchoCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpEcnEchoCounter
}

func (obj *patternFlowTcpEcnEchoCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpEcnEchoCounter) SetStart(value int32) PatternFlowTcpEcnEchoCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpEcnEchoCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpEcnEchoCounter) SetStep(value int32) PatternFlowTcpEcnEchoCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpEcnEchoCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpEcnEchoCounter) SetCount(value int32) PatternFlowTcpEcnEchoCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpCtlUrgCounter struct {
	obj *snappipb.PatternFlowTcpCtlUrgCounter
}

func (obj *patternFlowTcpCtlUrgCounter) msg() *snappipb.PatternFlowTcpCtlUrgCounter {
	return obj.obj
}

func (obj *patternFlowTcpCtlUrgCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpCtlUrgCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpCtlUrgCounter interface {
	msg() *snappipb.PatternFlowTcpCtlUrgCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpCtlUrgCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpCtlUrgCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpCtlUrgCounter
}

func (obj *patternFlowTcpCtlUrgCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpCtlUrgCounter) SetStart(value int32) PatternFlowTcpCtlUrgCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpCtlUrgCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpCtlUrgCounter) SetStep(value int32) PatternFlowTcpCtlUrgCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpCtlUrgCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpCtlUrgCounter) SetCount(value int32) PatternFlowTcpCtlUrgCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpCtlAckCounter struct {
	obj *snappipb.PatternFlowTcpCtlAckCounter
}

func (obj *patternFlowTcpCtlAckCounter) msg() *snappipb.PatternFlowTcpCtlAckCounter {
	return obj.obj
}

func (obj *patternFlowTcpCtlAckCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpCtlAckCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpCtlAckCounter interface {
	msg() *snappipb.PatternFlowTcpCtlAckCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpCtlAckCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpCtlAckCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpCtlAckCounter
}

func (obj *patternFlowTcpCtlAckCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpCtlAckCounter) SetStart(value int32) PatternFlowTcpCtlAckCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpCtlAckCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpCtlAckCounter) SetStep(value int32) PatternFlowTcpCtlAckCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpCtlAckCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpCtlAckCounter) SetCount(value int32) PatternFlowTcpCtlAckCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpCtlPshCounter struct {
	obj *snappipb.PatternFlowTcpCtlPshCounter
}

func (obj *patternFlowTcpCtlPshCounter) msg() *snappipb.PatternFlowTcpCtlPshCounter {
	return obj.obj
}

func (obj *patternFlowTcpCtlPshCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpCtlPshCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpCtlPshCounter interface {
	msg() *snappipb.PatternFlowTcpCtlPshCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpCtlPshCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpCtlPshCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpCtlPshCounter
}

func (obj *patternFlowTcpCtlPshCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpCtlPshCounter) SetStart(value int32) PatternFlowTcpCtlPshCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpCtlPshCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpCtlPshCounter) SetStep(value int32) PatternFlowTcpCtlPshCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpCtlPshCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpCtlPshCounter) SetCount(value int32) PatternFlowTcpCtlPshCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpCtlRstCounter struct {
	obj *snappipb.PatternFlowTcpCtlRstCounter
}

func (obj *patternFlowTcpCtlRstCounter) msg() *snappipb.PatternFlowTcpCtlRstCounter {
	return obj.obj
}

func (obj *patternFlowTcpCtlRstCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpCtlRstCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpCtlRstCounter interface {
	msg() *snappipb.PatternFlowTcpCtlRstCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpCtlRstCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpCtlRstCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpCtlRstCounter
}

func (obj *patternFlowTcpCtlRstCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpCtlRstCounter) SetStart(value int32) PatternFlowTcpCtlRstCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpCtlRstCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpCtlRstCounter) SetStep(value int32) PatternFlowTcpCtlRstCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpCtlRstCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpCtlRstCounter) SetCount(value int32) PatternFlowTcpCtlRstCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpCtlSynCounter struct {
	obj *snappipb.PatternFlowTcpCtlSynCounter
}

func (obj *patternFlowTcpCtlSynCounter) msg() *snappipb.PatternFlowTcpCtlSynCounter {
	return obj.obj
}

func (obj *patternFlowTcpCtlSynCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpCtlSynCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpCtlSynCounter interface {
	msg() *snappipb.PatternFlowTcpCtlSynCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpCtlSynCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpCtlSynCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpCtlSynCounter
}

func (obj *patternFlowTcpCtlSynCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpCtlSynCounter) SetStart(value int32) PatternFlowTcpCtlSynCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpCtlSynCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpCtlSynCounter) SetStep(value int32) PatternFlowTcpCtlSynCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpCtlSynCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpCtlSynCounter) SetCount(value int32) PatternFlowTcpCtlSynCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpCtlFinCounter struct {
	obj *snappipb.PatternFlowTcpCtlFinCounter
}

func (obj *patternFlowTcpCtlFinCounter) msg() *snappipb.PatternFlowTcpCtlFinCounter {
	return obj.obj
}

func (obj *patternFlowTcpCtlFinCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpCtlFinCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpCtlFinCounter interface {
	msg() *snappipb.PatternFlowTcpCtlFinCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpCtlFinCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpCtlFinCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpCtlFinCounter
}

func (obj *patternFlowTcpCtlFinCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpCtlFinCounter) SetStart(value int32) PatternFlowTcpCtlFinCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpCtlFinCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpCtlFinCounter) SetStep(value int32) PatternFlowTcpCtlFinCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpCtlFinCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpCtlFinCounter) SetCount(value int32) PatternFlowTcpCtlFinCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowTcpWindowCounter struct {
	obj *snappipb.PatternFlowTcpWindowCounter
}

func (obj *patternFlowTcpWindowCounter) msg() *snappipb.PatternFlowTcpWindowCounter {
	return obj.obj
}

func (obj *patternFlowTcpWindowCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowTcpWindowCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowTcpWindowCounter interface {
	msg() *snappipb.PatternFlowTcpWindowCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowTcpWindowCounter
	Step() int32
	SetStep(value int32) PatternFlowTcpWindowCounter
	Count() int32
	SetCount(value int32) PatternFlowTcpWindowCounter
}

func (obj *patternFlowTcpWindowCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowTcpWindowCounter) SetStart(value int32) PatternFlowTcpWindowCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowTcpWindowCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowTcpWindowCounter) SetStep(value int32) PatternFlowTcpWindowCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowTcpWindowCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowTcpWindowCounter) SetCount(value int32) PatternFlowTcpWindowCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowUdpSrcPortCounter struct {
	obj *snappipb.PatternFlowUdpSrcPortCounter
}

func (obj *patternFlowUdpSrcPortCounter) msg() *snappipb.PatternFlowUdpSrcPortCounter {
	return obj.obj
}

func (obj *patternFlowUdpSrcPortCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowUdpSrcPortCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowUdpSrcPortCounter interface {
	msg() *snappipb.PatternFlowUdpSrcPortCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowUdpSrcPortCounter
	Step() int32
	SetStep(value int32) PatternFlowUdpSrcPortCounter
	Count() int32
	SetCount(value int32) PatternFlowUdpSrcPortCounter
}

func (obj *patternFlowUdpSrcPortCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowUdpSrcPortCounter) SetStart(value int32) PatternFlowUdpSrcPortCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowUdpSrcPortCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowUdpSrcPortCounter) SetStep(value int32) PatternFlowUdpSrcPortCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowUdpSrcPortCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowUdpSrcPortCounter) SetCount(value int32) PatternFlowUdpSrcPortCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowUdpDstPortCounter struct {
	obj *snappipb.PatternFlowUdpDstPortCounter
}

func (obj *patternFlowUdpDstPortCounter) msg() *snappipb.PatternFlowUdpDstPortCounter {
	return obj.obj
}

func (obj *patternFlowUdpDstPortCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowUdpDstPortCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowUdpDstPortCounter interface {
	msg() *snappipb.PatternFlowUdpDstPortCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowUdpDstPortCounter
	Step() int32
	SetStep(value int32) PatternFlowUdpDstPortCounter
	Count() int32
	SetCount(value int32) PatternFlowUdpDstPortCounter
}

func (obj *patternFlowUdpDstPortCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowUdpDstPortCounter) SetStart(value int32) PatternFlowUdpDstPortCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowUdpDstPortCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowUdpDstPortCounter) SetStep(value int32) PatternFlowUdpDstPortCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowUdpDstPortCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowUdpDstPortCounter) SetCount(value int32) PatternFlowUdpDstPortCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowUdpLengthCounter struct {
	obj *snappipb.PatternFlowUdpLengthCounter
}

func (obj *patternFlowUdpLengthCounter) msg() *snappipb.PatternFlowUdpLengthCounter {
	return obj.obj
}

func (obj *patternFlowUdpLengthCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowUdpLengthCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowUdpLengthCounter interface {
	msg() *snappipb.PatternFlowUdpLengthCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowUdpLengthCounter
	Step() int32
	SetStep(value int32) PatternFlowUdpLengthCounter
	Count() int32
	SetCount(value int32) PatternFlowUdpLengthCounter
}

func (obj *patternFlowUdpLengthCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowUdpLengthCounter) SetStart(value int32) PatternFlowUdpLengthCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowUdpLengthCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowUdpLengthCounter) SetStep(value int32) PatternFlowUdpLengthCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowUdpLengthCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowUdpLengthCounter) SetCount(value int32) PatternFlowUdpLengthCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGreChecksumPresentCounter struct {
	obj *snappipb.PatternFlowGreChecksumPresentCounter
}

func (obj *patternFlowGreChecksumPresentCounter) msg() *snappipb.PatternFlowGreChecksumPresentCounter {
	return obj.obj
}

func (obj *patternFlowGreChecksumPresentCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGreChecksumPresentCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGreChecksumPresentCounter interface {
	msg() *snappipb.PatternFlowGreChecksumPresentCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGreChecksumPresentCounter
	Step() int32
	SetStep(value int32) PatternFlowGreChecksumPresentCounter
	Count() int32
	SetCount(value int32) PatternFlowGreChecksumPresentCounter
}

func (obj *patternFlowGreChecksumPresentCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGreChecksumPresentCounter) SetStart(value int32) PatternFlowGreChecksumPresentCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGreChecksumPresentCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGreChecksumPresentCounter) SetStep(value int32) PatternFlowGreChecksumPresentCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGreChecksumPresentCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGreChecksumPresentCounter) SetCount(value int32) PatternFlowGreChecksumPresentCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGreReserved0Counter struct {
	obj *snappipb.PatternFlowGreReserved0Counter
}

func (obj *patternFlowGreReserved0Counter) msg() *snappipb.PatternFlowGreReserved0Counter {
	return obj.obj
}

func (obj *patternFlowGreReserved0Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGreReserved0Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGreReserved0Counter interface {
	msg() *snappipb.PatternFlowGreReserved0Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGreReserved0Counter
	Step() int32
	SetStep(value int32) PatternFlowGreReserved0Counter
	Count() int32
	SetCount(value int32) PatternFlowGreReserved0Counter
}

func (obj *patternFlowGreReserved0Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGreReserved0Counter) SetStart(value int32) PatternFlowGreReserved0Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGreReserved0Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGreReserved0Counter) SetStep(value int32) PatternFlowGreReserved0Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGreReserved0Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGreReserved0Counter) SetCount(value int32) PatternFlowGreReserved0Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGreVersionCounter struct {
	obj *snappipb.PatternFlowGreVersionCounter
}

func (obj *patternFlowGreVersionCounter) msg() *snappipb.PatternFlowGreVersionCounter {
	return obj.obj
}

func (obj *patternFlowGreVersionCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGreVersionCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGreVersionCounter interface {
	msg() *snappipb.PatternFlowGreVersionCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGreVersionCounter
	Step() int32
	SetStep(value int32) PatternFlowGreVersionCounter
	Count() int32
	SetCount(value int32) PatternFlowGreVersionCounter
}

func (obj *patternFlowGreVersionCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGreVersionCounter) SetStart(value int32) PatternFlowGreVersionCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGreVersionCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGreVersionCounter) SetStep(value int32) PatternFlowGreVersionCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGreVersionCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGreVersionCounter) SetCount(value int32) PatternFlowGreVersionCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGreProtocolCounter struct {
	obj *snappipb.PatternFlowGreProtocolCounter
}

func (obj *patternFlowGreProtocolCounter) msg() *snappipb.PatternFlowGreProtocolCounter {
	return obj.obj
}

func (obj *patternFlowGreProtocolCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGreProtocolCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGreProtocolCounter interface {
	msg() *snappipb.PatternFlowGreProtocolCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGreProtocolCounter
	Step() int32
	SetStep(value int32) PatternFlowGreProtocolCounter
	Count() int32
	SetCount(value int32) PatternFlowGreProtocolCounter
}

func (obj *patternFlowGreProtocolCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGreProtocolCounter) SetStart(value int32) PatternFlowGreProtocolCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGreProtocolCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGreProtocolCounter) SetStep(value int32) PatternFlowGreProtocolCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGreProtocolCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGreProtocolCounter) SetCount(value int32) PatternFlowGreProtocolCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGreReserved1Counter struct {
	obj *snappipb.PatternFlowGreReserved1Counter
}

func (obj *patternFlowGreReserved1Counter) msg() *snappipb.PatternFlowGreReserved1Counter {
	return obj.obj
}

func (obj *patternFlowGreReserved1Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGreReserved1Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGreReserved1Counter interface {
	msg() *snappipb.PatternFlowGreReserved1Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGreReserved1Counter
	Step() int32
	SetStep(value int32) PatternFlowGreReserved1Counter
	Count() int32
	SetCount(value int32) PatternFlowGreReserved1Counter
}

func (obj *patternFlowGreReserved1Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGreReserved1Counter) SetStart(value int32) PatternFlowGreReserved1Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGreReserved1Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGreReserved1Counter) SetStep(value int32) PatternFlowGreReserved1Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGreReserved1Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGreReserved1Counter) SetCount(value int32) PatternFlowGreReserved1Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv1VersionCounter struct {
	obj *snappipb.PatternFlowGtpv1VersionCounter
}

func (obj *patternFlowGtpv1VersionCounter) msg() *snappipb.PatternFlowGtpv1VersionCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1VersionCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1VersionCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1VersionCounter interface {
	msg() *snappipb.PatternFlowGtpv1VersionCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv1VersionCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv1VersionCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv1VersionCounter
}

func (obj *patternFlowGtpv1VersionCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv1VersionCounter) SetStart(value int32) PatternFlowGtpv1VersionCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv1VersionCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv1VersionCounter) SetStep(value int32) PatternFlowGtpv1VersionCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv1VersionCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv1VersionCounter) SetCount(value int32) PatternFlowGtpv1VersionCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv1ProtocolTypeCounter struct {
	obj *snappipb.PatternFlowGtpv1ProtocolTypeCounter
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) msg() *snappipb.PatternFlowGtpv1ProtocolTypeCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1ProtocolTypeCounter interface {
	msg() *snappipb.PatternFlowGtpv1ProtocolTypeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv1ProtocolTypeCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv1ProtocolTypeCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv1ProtocolTypeCounter
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) SetStart(value int32) PatternFlowGtpv1ProtocolTypeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) SetStep(value int32) PatternFlowGtpv1ProtocolTypeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv1ProtocolTypeCounter) SetCount(value int32) PatternFlowGtpv1ProtocolTypeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv1ReservedCounter struct {
	obj *snappipb.PatternFlowGtpv1ReservedCounter
}

func (obj *patternFlowGtpv1ReservedCounter) msg() *snappipb.PatternFlowGtpv1ReservedCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1ReservedCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1ReservedCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1ReservedCounter interface {
	msg() *snappipb.PatternFlowGtpv1ReservedCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv1ReservedCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv1ReservedCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv1ReservedCounter
}

func (obj *patternFlowGtpv1ReservedCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv1ReservedCounter) SetStart(value int32) PatternFlowGtpv1ReservedCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv1ReservedCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv1ReservedCounter) SetStep(value int32) PatternFlowGtpv1ReservedCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv1ReservedCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv1ReservedCounter) SetCount(value int32) PatternFlowGtpv1ReservedCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv1EFlagCounter struct {
	obj *snappipb.PatternFlowGtpv1EFlagCounter
}

func (obj *patternFlowGtpv1EFlagCounter) msg() *snappipb.PatternFlowGtpv1EFlagCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1EFlagCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1EFlagCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1EFlagCounter interface {
	msg() *snappipb.PatternFlowGtpv1EFlagCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv1EFlagCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv1EFlagCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv1EFlagCounter
}

func (obj *patternFlowGtpv1EFlagCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv1EFlagCounter) SetStart(value int32) PatternFlowGtpv1EFlagCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv1EFlagCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv1EFlagCounter) SetStep(value int32) PatternFlowGtpv1EFlagCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv1EFlagCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv1EFlagCounter) SetCount(value int32) PatternFlowGtpv1EFlagCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv1SFlagCounter struct {
	obj *snappipb.PatternFlowGtpv1SFlagCounter
}

func (obj *patternFlowGtpv1SFlagCounter) msg() *snappipb.PatternFlowGtpv1SFlagCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1SFlagCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1SFlagCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1SFlagCounter interface {
	msg() *snappipb.PatternFlowGtpv1SFlagCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv1SFlagCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv1SFlagCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv1SFlagCounter
}

func (obj *patternFlowGtpv1SFlagCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv1SFlagCounter) SetStart(value int32) PatternFlowGtpv1SFlagCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv1SFlagCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv1SFlagCounter) SetStep(value int32) PatternFlowGtpv1SFlagCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv1SFlagCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv1SFlagCounter) SetCount(value int32) PatternFlowGtpv1SFlagCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv1PnFlagCounter struct {
	obj *snappipb.PatternFlowGtpv1PnFlagCounter
}

func (obj *patternFlowGtpv1PnFlagCounter) msg() *snappipb.PatternFlowGtpv1PnFlagCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1PnFlagCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1PnFlagCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1PnFlagCounter interface {
	msg() *snappipb.PatternFlowGtpv1PnFlagCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv1PnFlagCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv1PnFlagCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv1PnFlagCounter
}

func (obj *patternFlowGtpv1PnFlagCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv1PnFlagCounter) SetStart(value int32) PatternFlowGtpv1PnFlagCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv1PnFlagCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv1PnFlagCounter) SetStep(value int32) PatternFlowGtpv1PnFlagCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv1PnFlagCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv1PnFlagCounter) SetCount(value int32) PatternFlowGtpv1PnFlagCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv1MessageTypeCounter struct {
	obj *snappipb.PatternFlowGtpv1MessageTypeCounter
}

func (obj *patternFlowGtpv1MessageTypeCounter) msg() *snappipb.PatternFlowGtpv1MessageTypeCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1MessageTypeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1MessageTypeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1MessageTypeCounter interface {
	msg() *snappipb.PatternFlowGtpv1MessageTypeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv1MessageTypeCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv1MessageTypeCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv1MessageTypeCounter
}

func (obj *patternFlowGtpv1MessageTypeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv1MessageTypeCounter) SetStart(value int32) PatternFlowGtpv1MessageTypeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv1MessageTypeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv1MessageTypeCounter) SetStep(value int32) PatternFlowGtpv1MessageTypeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv1MessageTypeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv1MessageTypeCounter) SetCount(value int32) PatternFlowGtpv1MessageTypeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv1MessageLengthCounter struct {
	obj *snappipb.PatternFlowGtpv1MessageLengthCounter
}

func (obj *patternFlowGtpv1MessageLengthCounter) msg() *snappipb.PatternFlowGtpv1MessageLengthCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1MessageLengthCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1MessageLengthCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1MessageLengthCounter interface {
	msg() *snappipb.PatternFlowGtpv1MessageLengthCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv1MessageLengthCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv1MessageLengthCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv1MessageLengthCounter
}

func (obj *patternFlowGtpv1MessageLengthCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv1MessageLengthCounter) SetStart(value int32) PatternFlowGtpv1MessageLengthCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv1MessageLengthCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv1MessageLengthCounter) SetStep(value int32) PatternFlowGtpv1MessageLengthCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv1MessageLengthCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv1MessageLengthCounter) SetCount(value int32) PatternFlowGtpv1MessageLengthCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv1TeidCounter struct {
	obj *snappipb.PatternFlowGtpv1TeidCounter
}

func (obj *patternFlowGtpv1TeidCounter) msg() *snappipb.PatternFlowGtpv1TeidCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1TeidCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1TeidCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1TeidCounter interface {
	msg() *snappipb.PatternFlowGtpv1TeidCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv1TeidCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv1TeidCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv1TeidCounter
}

func (obj *patternFlowGtpv1TeidCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv1TeidCounter) SetStart(value int32) PatternFlowGtpv1TeidCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv1TeidCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv1TeidCounter) SetStep(value int32) PatternFlowGtpv1TeidCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv1TeidCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv1TeidCounter) SetCount(value int32) PatternFlowGtpv1TeidCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv1SquenceNumberCounter struct {
	obj *snappipb.PatternFlowGtpv1SquenceNumberCounter
}

func (obj *patternFlowGtpv1SquenceNumberCounter) msg() *snappipb.PatternFlowGtpv1SquenceNumberCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1SquenceNumberCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1SquenceNumberCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1SquenceNumberCounter interface {
	msg() *snappipb.PatternFlowGtpv1SquenceNumberCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv1SquenceNumberCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv1SquenceNumberCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv1SquenceNumberCounter
}

func (obj *patternFlowGtpv1SquenceNumberCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv1SquenceNumberCounter) SetStart(value int32) PatternFlowGtpv1SquenceNumberCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv1SquenceNumberCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv1SquenceNumberCounter) SetStep(value int32) PatternFlowGtpv1SquenceNumberCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv1SquenceNumberCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv1SquenceNumberCounter) SetCount(value int32) PatternFlowGtpv1SquenceNumberCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv1NPduNumberCounter struct {
	obj *snappipb.PatternFlowGtpv1NPduNumberCounter
}

func (obj *patternFlowGtpv1NPduNumberCounter) msg() *snappipb.PatternFlowGtpv1NPduNumberCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1NPduNumberCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1NPduNumberCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1NPduNumberCounter interface {
	msg() *snappipb.PatternFlowGtpv1NPduNumberCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv1NPduNumberCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv1NPduNumberCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv1NPduNumberCounter
}

func (obj *patternFlowGtpv1NPduNumberCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv1NPduNumberCounter) SetStart(value int32) PatternFlowGtpv1NPduNumberCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv1NPduNumberCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv1NPduNumberCounter) SetStep(value int32) PatternFlowGtpv1NPduNumberCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv1NPduNumberCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv1NPduNumberCounter) SetCount(value int32) PatternFlowGtpv1NPduNumberCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv1NextExtensionHeaderTypeCounter struct {
	obj *snappipb.PatternFlowGtpv1NextExtensionHeaderTypeCounter
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) msg() *snappipb.PatternFlowGtpv1NextExtensionHeaderTypeCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv1NextExtensionHeaderTypeCounter interface {
	msg() *snappipb.PatternFlowGtpv1NextExtensionHeaderTypeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv1NextExtensionHeaderTypeCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv1NextExtensionHeaderTypeCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv1NextExtensionHeaderTypeCounter
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) SetStart(value int32) PatternFlowGtpv1NextExtensionHeaderTypeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) SetStep(value int32) PatternFlowGtpv1NextExtensionHeaderTypeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeCounter) SetCount(value int32) PatternFlowGtpv1NextExtensionHeaderTypeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpExtensionExtensionLength struct {
	obj *snappipb.PatternFlowGtpExtensionExtensionLength
}

func (obj *patternFlowGtpExtensionExtensionLength) msg() *snappipb.PatternFlowGtpExtensionExtensionLength {
	return obj.obj
}

func (obj *patternFlowGtpExtensionExtensionLength) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpExtensionExtensionLength) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpExtensionExtensionLength interface {
	msg() *snappipb.PatternFlowGtpExtensionExtensionLength
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpExtensionExtensionLength
	Values() []int32
	SetValues(value []int32) PatternFlowGtpExtensionExtensionLength
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpExtensionExtensionLength
	Increment() PatternFlowGtpExtensionExtensionLengthCounter
	Decrement() PatternFlowGtpExtensionExtensionLengthCounter
}

func (obj *patternFlowGtpExtensionExtensionLength) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpExtensionExtensionLength) SetValue(value int32) PatternFlowGtpExtensionExtensionLength {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpExtensionExtensionLength) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpExtensionExtensionLength) SetValues(value []int32) PatternFlowGtpExtensionExtensionLength {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpExtensionExtensionLength) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpExtensionExtensionLength) SetMetricGroup(value string) PatternFlowGtpExtensionExtensionLength {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpExtensionExtensionLength) Increment() PatternFlowGtpExtensionExtensionLengthCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpExtensionExtensionLengthCounter{}
	}
	return &patternFlowGtpExtensionExtensionLengthCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpExtensionExtensionLength) Decrement() PatternFlowGtpExtensionExtensionLengthCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpExtensionExtensionLengthCounter{}
	}
	return &patternFlowGtpExtensionExtensionLengthCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpExtensionContents struct {
	obj *snappipb.PatternFlowGtpExtensionContents
}

func (obj *patternFlowGtpExtensionContents) msg() *snappipb.PatternFlowGtpExtensionContents {
	return obj.obj
}

func (obj *patternFlowGtpExtensionContents) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpExtensionContents) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpExtensionContents interface {
	msg() *snappipb.PatternFlowGtpExtensionContents
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpExtensionContents
	Values() []int32
	SetValues(value []int32) PatternFlowGtpExtensionContents
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpExtensionContents
	Increment() PatternFlowGtpExtensionContentsCounter
	Decrement() PatternFlowGtpExtensionContentsCounter
}

func (obj *patternFlowGtpExtensionContents) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpExtensionContents) SetValue(value int32) PatternFlowGtpExtensionContents {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpExtensionContents) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpExtensionContents) SetValues(value []int32) PatternFlowGtpExtensionContents {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpExtensionContents) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpExtensionContents) SetMetricGroup(value string) PatternFlowGtpExtensionContents {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpExtensionContents) Increment() PatternFlowGtpExtensionContentsCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpExtensionContentsCounter{}
	}
	return &patternFlowGtpExtensionContentsCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpExtensionContents) Decrement() PatternFlowGtpExtensionContentsCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpExtensionContentsCounter{}
	}
	return &patternFlowGtpExtensionContentsCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpExtensionNextExtensionHeader struct {
	obj *snappipb.PatternFlowGtpExtensionNextExtensionHeader
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) msg() *snappipb.PatternFlowGtpExtensionNextExtensionHeader {
	return obj.obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpExtensionNextExtensionHeader interface {
	msg() *snappipb.PatternFlowGtpExtensionNextExtensionHeader
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowGtpExtensionNextExtensionHeader
	Values() []int32
	SetValues(value []int32) PatternFlowGtpExtensionNextExtensionHeader
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowGtpExtensionNextExtensionHeader
	Increment() PatternFlowGtpExtensionNextExtensionHeaderCounter
	Decrement() PatternFlowGtpExtensionNextExtensionHeaderCounter
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) SetValue(value int32) PatternFlowGtpExtensionNextExtensionHeader {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) SetValues(value []int32) PatternFlowGtpExtensionNextExtensionHeader {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) SetMetricGroup(value string) PatternFlowGtpExtensionNextExtensionHeader {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) Increment() PatternFlowGtpExtensionNextExtensionHeaderCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowGtpExtensionNextExtensionHeaderCounter{}
	}
	return &patternFlowGtpExtensionNextExtensionHeaderCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowGtpExtensionNextExtensionHeader) Decrement() PatternFlowGtpExtensionNextExtensionHeaderCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowGtpExtensionNextExtensionHeaderCounter{}
	}
	return &patternFlowGtpExtensionNextExtensionHeaderCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpv2VersionCounter struct {
	obj *snappipb.PatternFlowGtpv2VersionCounter
}

func (obj *patternFlowGtpv2VersionCounter) msg() *snappipb.PatternFlowGtpv2VersionCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2VersionCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2VersionCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2VersionCounter interface {
	msg() *snappipb.PatternFlowGtpv2VersionCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv2VersionCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv2VersionCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv2VersionCounter
}

func (obj *patternFlowGtpv2VersionCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv2VersionCounter) SetStart(value int32) PatternFlowGtpv2VersionCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv2VersionCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv2VersionCounter) SetStep(value int32) PatternFlowGtpv2VersionCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv2VersionCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv2VersionCounter) SetCount(value int32) PatternFlowGtpv2VersionCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv2PiggybackingFlagCounter struct {
	obj *snappipb.PatternFlowGtpv2PiggybackingFlagCounter
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) msg() *snappipb.PatternFlowGtpv2PiggybackingFlagCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2PiggybackingFlagCounter interface {
	msg() *snappipb.PatternFlowGtpv2PiggybackingFlagCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv2PiggybackingFlagCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv2PiggybackingFlagCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv2PiggybackingFlagCounter
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) SetStart(value int32) PatternFlowGtpv2PiggybackingFlagCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) SetStep(value int32) PatternFlowGtpv2PiggybackingFlagCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) SetCount(value int32) PatternFlowGtpv2PiggybackingFlagCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv2TeidFlagCounter struct {
	obj *snappipb.PatternFlowGtpv2TeidFlagCounter
}

func (obj *patternFlowGtpv2TeidFlagCounter) msg() *snappipb.PatternFlowGtpv2TeidFlagCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2TeidFlagCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2TeidFlagCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2TeidFlagCounter interface {
	msg() *snappipb.PatternFlowGtpv2TeidFlagCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv2TeidFlagCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv2TeidFlagCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv2TeidFlagCounter
}

func (obj *patternFlowGtpv2TeidFlagCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv2TeidFlagCounter) SetStart(value int32) PatternFlowGtpv2TeidFlagCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv2TeidFlagCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv2TeidFlagCounter) SetStep(value int32) PatternFlowGtpv2TeidFlagCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv2TeidFlagCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv2TeidFlagCounter) SetCount(value int32) PatternFlowGtpv2TeidFlagCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv2Spare1Counter struct {
	obj *snappipb.PatternFlowGtpv2Spare1Counter
}

func (obj *patternFlowGtpv2Spare1Counter) msg() *snappipb.PatternFlowGtpv2Spare1Counter {
	return obj.obj
}

func (obj *patternFlowGtpv2Spare1Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2Spare1Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2Spare1Counter interface {
	msg() *snappipb.PatternFlowGtpv2Spare1Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv2Spare1Counter
	Step() int32
	SetStep(value int32) PatternFlowGtpv2Spare1Counter
	Count() int32
	SetCount(value int32) PatternFlowGtpv2Spare1Counter
}

func (obj *patternFlowGtpv2Spare1Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv2Spare1Counter) SetStart(value int32) PatternFlowGtpv2Spare1Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv2Spare1Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv2Spare1Counter) SetStep(value int32) PatternFlowGtpv2Spare1Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv2Spare1Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv2Spare1Counter) SetCount(value int32) PatternFlowGtpv2Spare1Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv2MessageTypeCounter struct {
	obj *snappipb.PatternFlowGtpv2MessageTypeCounter
}

func (obj *patternFlowGtpv2MessageTypeCounter) msg() *snappipb.PatternFlowGtpv2MessageTypeCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2MessageTypeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2MessageTypeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2MessageTypeCounter interface {
	msg() *snappipb.PatternFlowGtpv2MessageTypeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv2MessageTypeCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv2MessageTypeCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv2MessageTypeCounter
}

func (obj *patternFlowGtpv2MessageTypeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv2MessageTypeCounter) SetStart(value int32) PatternFlowGtpv2MessageTypeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv2MessageTypeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv2MessageTypeCounter) SetStep(value int32) PatternFlowGtpv2MessageTypeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv2MessageTypeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv2MessageTypeCounter) SetCount(value int32) PatternFlowGtpv2MessageTypeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv2MessageLengthCounter struct {
	obj *snappipb.PatternFlowGtpv2MessageLengthCounter
}

func (obj *patternFlowGtpv2MessageLengthCounter) msg() *snappipb.PatternFlowGtpv2MessageLengthCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2MessageLengthCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2MessageLengthCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2MessageLengthCounter interface {
	msg() *snappipb.PatternFlowGtpv2MessageLengthCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv2MessageLengthCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv2MessageLengthCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv2MessageLengthCounter
}

func (obj *patternFlowGtpv2MessageLengthCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv2MessageLengthCounter) SetStart(value int32) PatternFlowGtpv2MessageLengthCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv2MessageLengthCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv2MessageLengthCounter) SetStep(value int32) PatternFlowGtpv2MessageLengthCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv2MessageLengthCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv2MessageLengthCounter) SetCount(value int32) PatternFlowGtpv2MessageLengthCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv2TeidCounter struct {
	obj *snappipb.PatternFlowGtpv2TeidCounter
}

func (obj *patternFlowGtpv2TeidCounter) msg() *snappipb.PatternFlowGtpv2TeidCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2TeidCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2TeidCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2TeidCounter interface {
	msg() *snappipb.PatternFlowGtpv2TeidCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv2TeidCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv2TeidCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv2TeidCounter
}

func (obj *patternFlowGtpv2TeidCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv2TeidCounter) SetStart(value int32) PatternFlowGtpv2TeidCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv2TeidCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv2TeidCounter) SetStep(value int32) PatternFlowGtpv2TeidCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv2TeidCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv2TeidCounter) SetCount(value int32) PatternFlowGtpv2TeidCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv2SequenceNumberCounter struct {
	obj *snappipb.PatternFlowGtpv2SequenceNumberCounter
}

func (obj *patternFlowGtpv2SequenceNumberCounter) msg() *snappipb.PatternFlowGtpv2SequenceNumberCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2SequenceNumberCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2SequenceNumberCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2SequenceNumberCounter interface {
	msg() *snappipb.PatternFlowGtpv2SequenceNumberCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv2SequenceNumberCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpv2SequenceNumberCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpv2SequenceNumberCounter
}

func (obj *patternFlowGtpv2SequenceNumberCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv2SequenceNumberCounter) SetStart(value int32) PatternFlowGtpv2SequenceNumberCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv2SequenceNumberCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv2SequenceNumberCounter) SetStep(value int32) PatternFlowGtpv2SequenceNumberCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv2SequenceNumberCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv2SequenceNumberCounter) SetCount(value int32) PatternFlowGtpv2SequenceNumberCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpv2Spare2Counter struct {
	obj *snappipb.PatternFlowGtpv2Spare2Counter
}

func (obj *patternFlowGtpv2Spare2Counter) msg() *snappipb.PatternFlowGtpv2Spare2Counter {
	return obj.obj
}

func (obj *patternFlowGtpv2Spare2Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpv2Spare2Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpv2Spare2Counter interface {
	msg() *snappipb.PatternFlowGtpv2Spare2Counter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpv2Spare2Counter
	Step() int32
	SetStep(value int32) PatternFlowGtpv2Spare2Counter
	Count() int32
	SetCount(value int32) PatternFlowGtpv2Spare2Counter
}

func (obj *patternFlowGtpv2Spare2Counter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpv2Spare2Counter) SetStart(value int32) PatternFlowGtpv2Spare2Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpv2Spare2Counter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpv2Spare2Counter) SetStep(value int32) PatternFlowGtpv2Spare2Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpv2Spare2Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpv2Spare2Counter) SetCount(value int32) PatternFlowGtpv2Spare2Counter {
	obj.obj.Count = &value
	return obj
}

type patternFlowArpHardwareTypeCounter struct {
	obj *snappipb.PatternFlowArpHardwareTypeCounter
}

func (obj *patternFlowArpHardwareTypeCounter) msg() *snappipb.PatternFlowArpHardwareTypeCounter {
	return obj.obj
}

func (obj *patternFlowArpHardwareTypeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpHardwareTypeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpHardwareTypeCounter interface {
	msg() *snappipb.PatternFlowArpHardwareTypeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowArpHardwareTypeCounter
	Step() int32
	SetStep(value int32) PatternFlowArpHardwareTypeCounter
	Count() int32
	SetCount(value int32) PatternFlowArpHardwareTypeCounter
}

func (obj *patternFlowArpHardwareTypeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowArpHardwareTypeCounter) SetStart(value int32) PatternFlowArpHardwareTypeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowArpHardwareTypeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowArpHardwareTypeCounter) SetStep(value int32) PatternFlowArpHardwareTypeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowArpHardwareTypeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowArpHardwareTypeCounter) SetCount(value int32) PatternFlowArpHardwareTypeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowArpProtocolTypeCounter struct {
	obj *snappipb.PatternFlowArpProtocolTypeCounter
}

func (obj *patternFlowArpProtocolTypeCounter) msg() *snappipb.PatternFlowArpProtocolTypeCounter {
	return obj.obj
}

func (obj *patternFlowArpProtocolTypeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpProtocolTypeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpProtocolTypeCounter interface {
	msg() *snappipb.PatternFlowArpProtocolTypeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowArpProtocolTypeCounter
	Step() int32
	SetStep(value int32) PatternFlowArpProtocolTypeCounter
	Count() int32
	SetCount(value int32) PatternFlowArpProtocolTypeCounter
}

func (obj *patternFlowArpProtocolTypeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowArpProtocolTypeCounter) SetStart(value int32) PatternFlowArpProtocolTypeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowArpProtocolTypeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowArpProtocolTypeCounter) SetStep(value int32) PatternFlowArpProtocolTypeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowArpProtocolTypeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowArpProtocolTypeCounter) SetCount(value int32) PatternFlowArpProtocolTypeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowArpHardwareLengthCounter struct {
	obj *snappipb.PatternFlowArpHardwareLengthCounter
}

func (obj *patternFlowArpHardwareLengthCounter) msg() *snappipb.PatternFlowArpHardwareLengthCounter {
	return obj.obj
}

func (obj *patternFlowArpHardwareLengthCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpHardwareLengthCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpHardwareLengthCounter interface {
	msg() *snappipb.PatternFlowArpHardwareLengthCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowArpHardwareLengthCounter
	Step() int32
	SetStep(value int32) PatternFlowArpHardwareLengthCounter
	Count() int32
	SetCount(value int32) PatternFlowArpHardwareLengthCounter
}

func (obj *patternFlowArpHardwareLengthCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowArpHardwareLengthCounter) SetStart(value int32) PatternFlowArpHardwareLengthCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowArpHardwareLengthCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowArpHardwareLengthCounter) SetStep(value int32) PatternFlowArpHardwareLengthCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowArpHardwareLengthCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowArpHardwareLengthCounter) SetCount(value int32) PatternFlowArpHardwareLengthCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowArpProtocolLengthCounter struct {
	obj *snappipb.PatternFlowArpProtocolLengthCounter
}

func (obj *patternFlowArpProtocolLengthCounter) msg() *snappipb.PatternFlowArpProtocolLengthCounter {
	return obj.obj
}

func (obj *patternFlowArpProtocolLengthCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpProtocolLengthCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpProtocolLengthCounter interface {
	msg() *snappipb.PatternFlowArpProtocolLengthCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowArpProtocolLengthCounter
	Step() int32
	SetStep(value int32) PatternFlowArpProtocolLengthCounter
	Count() int32
	SetCount(value int32) PatternFlowArpProtocolLengthCounter
}

func (obj *patternFlowArpProtocolLengthCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowArpProtocolLengthCounter) SetStart(value int32) PatternFlowArpProtocolLengthCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowArpProtocolLengthCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowArpProtocolLengthCounter) SetStep(value int32) PatternFlowArpProtocolLengthCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowArpProtocolLengthCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowArpProtocolLengthCounter) SetCount(value int32) PatternFlowArpProtocolLengthCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowArpOperationCounter struct {
	obj *snappipb.PatternFlowArpOperationCounter
}

func (obj *patternFlowArpOperationCounter) msg() *snappipb.PatternFlowArpOperationCounter {
	return obj.obj
}

func (obj *patternFlowArpOperationCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpOperationCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpOperationCounter interface {
	msg() *snappipb.PatternFlowArpOperationCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowArpOperationCounter
	Step() int32
	SetStep(value int32) PatternFlowArpOperationCounter
	Count() int32
	SetCount(value int32) PatternFlowArpOperationCounter
}

func (obj *patternFlowArpOperationCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowArpOperationCounter) SetStart(value int32) PatternFlowArpOperationCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowArpOperationCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowArpOperationCounter) SetStep(value int32) PatternFlowArpOperationCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowArpOperationCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowArpOperationCounter) SetCount(value int32) PatternFlowArpOperationCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowArpSenderHardwareAddrCounter struct {
	obj *snappipb.PatternFlowArpSenderHardwareAddrCounter
}

func (obj *patternFlowArpSenderHardwareAddrCounter) msg() *snappipb.PatternFlowArpSenderHardwareAddrCounter {
	return obj.obj
}

func (obj *patternFlowArpSenderHardwareAddrCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpSenderHardwareAddrCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpSenderHardwareAddrCounter interface {
	msg() *snappipb.PatternFlowArpSenderHardwareAddrCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowArpSenderHardwareAddrCounter
	Step() string
	SetStep(value string) PatternFlowArpSenderHardwareAddrCounter
	Count() int32
	SetCount(value int32) PatternFlowArpSenderHardwareAddrCounter
}

func (obj *patternFlowArpSenderHardwareAddrCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowArpSenderHardwareAddrCounter) SetStart(value string) PatternFlowArpSenderHardwareAddrCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowArpSenderHardwareAddrCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowArpSenderHardwareAddrCounter) SetStep(value string) PatternFlowArpSenderHardwareAddrCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowArpSenderHardwareAddrCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowArpSenderHardwareAddrCounter) SetCount(value int32) PatternFlowArpSenderHardwareAddrCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowArpSenderProtocolAddrCounter struct {
	obj *snappipb.PatternFlowArpSenderProtocolAddrCounter
}

func (obj *patternFlowArpSenderProtocolAddrCounter) msg() *snappipb.PatternFlowArpSenderProtocolAddrCounter {
	return obj.obj
}

func (obj *patternFlowArpSenderProtocolAddrCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpSenderProtocolAddrCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpSenderProtocolAddrCounter interface {
	msg() *snappipb.PatternFlowArpSenderProtocolAddrCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowArpSenderProtocolAddrCounter
	Step() string
	SetStep(value string) PatternFlowArpSenderProtocolAddrCounter
	Count() int32
	SetCount(value int32) PatternFlowArpSenderProtocolAddrCounter
}

func (obj *patternFlowArpSenderProtocolAddrCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowArpSenderProtocolAddrCounter) SetStart(value string) PatternFlowArpSenderProtocolAddrCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowArpSenderProtocolAddrCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowArpSenderProtocolAddrCounter) SetStep(value string) PatternFlowArpSenderProtocolAddrCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowArpSenderProtocolAddrCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowArpSenderProtocolAddrCounter) SetCount(value int32) PatternFlowArpSenderProtocolAddrCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowArpTargetHardwareAddrCounter struct {
	obj *snappipb.PatternFlowArpTargetHardwareAddrCounter
}

func (obj *patternFlowArpTargetHardwareAddrCounter) msg() *snappipb.PatternFlowArpTargetHardwareAddrCounter {
	return obj.obj
}

func (obj *patternFlowArpTargetHardwareAddrCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpTargetHardwareAddrCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpTargetHardwareAddrCounter interface {
	msg() *snappipb.PatternFlowArpTargetHardwareAddrCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowArpTargetHardwareAddrCounter
	Step() string
	SetStep(value string) PatternFlowArpTargetHardwareAddrCounter
	Count() int32
	SetCount(value int32) PatternFlowArpTargetHardwareAddrCounter
}

func (obj *patternFlowArpTargetHardwareAddrCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowArpTargetHardwareAddrCounter) SetStart(value string) PatternFlowArpTargetHardwareAddrCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowArpTargetHardwareAddrCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowArpTargetHardwareAddrCounter) SetStep(value string) PatternFlowArpTargetHardwareAddrCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowArpTargetHardwareAddrCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowArpTargetHardwareAddrCounter) SetCount(value int32) PatternFlowArpTargetHardwareAddrCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowArpTargetProtocolAddrCounter struct {
	obj *snappipb.PatternFlowArpTargetProtocolAddrCounter
}

func (obj *patternFlowArpTargetProtocolAddrCounter) msg() *snappipb.PatternFlowArpTargetProtocolAddrCounter {
	return obj.obj
}

func (obj *patternFlowArpTargetProtocolAddrCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowArpTargetProtocolAddrCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowArpTargetProtocolAddrCounter interface {
	msg() *snappipb.PatternFlowArpTargetProtocolAddrCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowArpTargetProtocolAddrCounter
	Step() string
	SetStep(value string) PatternFlowArpTargetProtocolAddrCounter
	Count() int32
	SetCount(value int32) PatternFlowArpTargetProtocolAddrCounter
}

func (obj *patternFlowArpTargetProtocolAddrCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowArpTargetProtocolAddrCounter) SetStart(value string) PatternFlowArpTargetProtocolAddrCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowArpTargetProtocolAddrCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowArpTargetProtocolAddrCounter) SetStep(value string) PatternFlowArpTargetProtocolAddrCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowArpTargetProtocolAddrCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowArpTargetProtocolAddrCounter) SetCount(value int32) PatternFlowArpTargetProtocolAddrCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIcmpEchoType struct {
	obj *snappipb.PatternFlowIcmpEchoType
}

func (obj *patternFlowIcmpEchoType) msg() *snappipb.PatternFlowIcmpEchoType {
	return obj.obj
}

func (obj *patternFlowIcmpEchoType) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpEchoType) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpEchoType interface {
	msg() *snappipb.PatternFlowIcmpEchoType
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIcmpEchoType
	Values() []int32
	SetValues(value []int32) PatternFlowIcmpEchoType
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIcmpEchoType
	Increment() PatternFlowIcmpEchoTypeCounter
	Decrement() PatternFlowIcmpEchoTypeCounter
}

func (obj *patternFlowIcmpEchoType) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIcmpEchoType) SetValue(value int32) PatternFlowIcmpEchoType {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIcmpEchoType) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIcmpEchoType) SetValues(value []int32) PatternFlowIcmpEchoType {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIcmpEchoType) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIcmpEchoType) SetMetricGroup(value string) PatternFlowIcmpEchoType {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIcmpEchoType) Increment() PatternFlowIcmpEchoTypeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIcmpEchoTypeCounter{}
	}
	return &patternFlowIcmpEchoTypeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIcmpEchoType) Decrement() PatternFlowIcmpEchoTypeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIcmpEchoTypeCounter{}
	}
	return &patternFlowIcmpEchoTypeCounter{obj: obj.obj.Decrement}

}

type patternFlowIcmpEchoCode struct {
	obj *snappipb.PatternFlowIcmpEchoCode
}

func (obj *patternFlowIcmpEchoCode) msg() *snappipb.PatternFlowIcmpEchoCode {
	return obj.obj
}

func (obj *patternFlowIcmpEchoCode) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpEchoCode) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpEchoCode interface {
	msg() *snappipb.PatternFlowIcmpEchoCode
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIcmpEchoCode
	Values() []int32
	SetValues(value []int32) PatternFlowIcmpEchoCode
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIcmpEchoCode
	Increment() PatternFlowIcmpEchoCodeCounter
	Decrement() PatternFlowIcmpEchoCodeCounter
}

func (obj *patternFlowIcmpEchoCode) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIcmpEchoCode) SetValue(value int32) PatternFlowIcmpEchoCode {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIcmpEchoCode) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIcmpEchoCode) SetValues(value []int32) PatternFlowIcmpEchoCode {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIcmpEchoCode) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIcmpEchoCode) SetMetricGroup(value string) PatternFlowIcmpEchoCode {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIcmpEchoCode) Increment() PatternFlowIcmpEchoCodeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIcmpEchoCodeCounter{}
	}
	return &patternFlowIcmpEchoCodeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIcmpEchoCode) Decrement() PatternFlowIcmpEchoCodeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIcmpEchoCodeCounter{}
	}
	return &patternFlowIcmpEchoCodeCounter{obj: obj.obj.Decrement}

}

type patternFlowIcmpEchoChecksum struct {
	obj *snappipb.PatternFlowIcmpEchoChecksum
}

func (obj *patternFlowIcmpEchoChecksum) msg() *snappipb.PatternFlowIcmpEchoChecksum {
	return obj.obj
}

func (obj *patternFlowIcmpEchoChecksum) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpEchoChecksum) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpEchoChecksum interface {
	msg() *snappipb.PatternFlowIcmpEchoChecksum
	Yaml() string
	Json() string
	Custom() int32
	SetCustom(value int32) PatternFlowIcmpEchoChecksum
}

func (obj *patternFlowIcmpEchoChecksum) Custom() int32 {
	return *obj.obj.Custom
}

func (obj *patternFlowIcmpEchoChecksum) SetCustom(value int32) PatternFlowIcmpEchoChecksum {
	obj.obj.Custom = &value
	return obj
}

type patternFlowIcmpEchoIdentifier struct {
	obj *snappipb.PatternFlowIcmpEchoIdentifier
}

func (obj *patternFlowIcmpEchoIdentifier) msg() *snappipb.PatternFlowIcmpEchoIdentifier {
	return obj.obj
}

func (obj *patternFlowIcmpEchoIdentifier) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpEchoIdentifier) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpEchoIdentifier interface {
	msg() *snappipb.PatternFlowIcmpEchoIdentifier
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIcmpEchoIdentifier
	Values() []int32
	SetValues(value []int32) PatternFlowIcmpEchoIdentifier
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIcmpEchoIdentifier
	Increment() PatternFlowIcmpEchoIdentifierCounter
	Decrement() PatternFlowIcmpEchoIdentifierCounter
}

func (obj *patternFlowIcmpEchoIdentifier) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIcmpEchoIdentifier) SetValue(value int32) PatternFlowIcmpEchoIdentifier {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIcmpEchoIdentifier) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIcmpEchoIdentifier) SetValues(value []int32) PatternFlowIcmpEchoIdentifier {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIcmpEchoIdentifier) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIcmpEchoIdentifier) SetMetricGroup(value string) PatternFlowIcmpEchoIdentifier {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIcmpEchoIdentifier) Increment() PatternFlowIcmpEchoIdentifierCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIcmpEchoIdentifierCounter{}
	}
	return &patternFlowIcmpEchoIdentifierCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIcmpEchoIdentifier) Decrement() PatternFlowIcmpEchoIdentifierCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIcmpEchoIdentifierCounter{}
	}
	return &patternFlowIcmpEchoIdentifierCounter{obj: obj.obj.Decrement}

}

type patternFlowIcmpEchoSequenceNumber struct {
	obj *snappipb.PatternFlowIcmpEchoSequenceNumber
}

func (obj *patternFlowIcmpEchoSequenceNumber) msg() *snappipb.PatternFlowIcmpEchoSequenceNumber {
	return obj.obj
}

func (obj *patternFlowIcmpEchoSequenceNumber) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpEchoSequenceNumber) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpEchoSequenceNumber interface {
	msg() *snappipb.PatternFlowIcmpEchoSequenceNumber
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIcmpEchoSequenceNumber
	Values() []int32
	SetValues(value []int32) PatternFlowIcmpEchoSequenceNumber
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIcmpEchoSequenceNumber
	Increment() PatternFlowIcmpEchoSequenceNumberCounter
	Decrement() PatternFlowIcmpEchoSequenceNumberCounter
}

func (obj *patternFlowIcmpEchoSequenceNumber) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIcmpEchoSequenceNumber) SetValue(value int32) PatternFlowIcmpEchoSequenceNumber {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIcmpEchoSequenceNumber) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIcmpEchoSequenceNumber) SetValues(value []int32) PatternFlowIcmpEchoSequenceNumber {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIcmpEchoSequenceNumber) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIcmpEchoSequenceNumber) SetMetricGroup(value string) PatternFlowIcmpEchoSequenceNumber {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIcmpEchoSequenceNumber) Increment() PatternFlowIcmpEchoSequenceNumberCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIcmpEchoSequenceNumberCounter{}
	}
	return &patternFlowIcmpEchoSequenceNumberCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIcmpEchoSequenceNumber) Decrement() PatternFlowIcmpEchoSequenceNumberCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIcmpEchoSequenceNumberCounter{}
	}
	return &patternFlowIcmpEchoSequenceNumberCounter{obj: obj.obj.Decrement}

}

type patternFlowIcmpv6EchoType struct {
	obj *snappipb.PatternFlowIcmpv6EchoType
}

func (obj *patternFlowIcmpv6EchoType) msg() *snappipb.PatternFlowIcmpv6EchoType {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoType) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpv6EchoType) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpv6EchoType interface {
	msg() *snappipb.PatternFlowIcmpv6EchoType
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIcmpv6EchoType
	Values() []int32
	SetValues(value []int32) PatternFlowIcmpv6EchoType
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIcmpv6EchoType
	Increment() PatternFlowIcmpv6EchoTypeCounter
	Decrement() PatternFlowIcmpv6EchoTypeCounter
}

func (obj *patternFlowIcmpv6EchoType) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIcmpv6EchoType) SetValue(value int32) PatternFlowIcmpv6EchoType {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoType) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIcmpv6EchoType) SetValues(value []int32) PatternFlowIcmpv6EchoType {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIcmpv6EchoType) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIcmpv6EchoType) SetMetricGroup(value string) PatternFlowIcmpv6EchoType {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoType) Increment() PatternFlowIcmpv6EchoTypeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIcmpv6EchoTypeCounter{}
	}
	return &patternFlowIcmpv6EchoTypeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIcmpv6EchoType) Decrement() PatternFlowIcmpv6EchoTypeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIcmpv6EchoTypeCounter{}
	}
	return &patternFlowIcmpv6EchoTypeCounter{obj: obj.obj.Decrement}

}

type patternFlowIcmpv6EchoCode struct {
	obj *snappipb.PatternFlowIcmpv6EchoCode
}

func (obj *patternFlowIcmpv6EchoCode) msg() *snappipb.PatternFlowIcmpv6EchoCode {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoCode) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpv6EchoCode) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpv6EchoCode interface {
	msg() *snappipb.PatternFlowIcmpv6EchoCode
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIcmpv6EchoCode
	Values() []int32
	SetValues(value []int32) PatternFlowIcmpv6EchoCode
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIcmpv6EchoCode
	Increment() PatternFlowIcmpv6EchoCodeCounter
	Decrement() PatternFlowIcmpv6EchoCodeCounter
}

func (obj *patternFlowIcmpv6EchoCode) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIcmpv6EchoCode) SetValue(value int32) PatternFlowIcmpv6EchoCode {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoCode) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIcmpv6EchoCode) SetValues(value []int32) PatternFlowIcmpv6EchoCode {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIcmpv6EchoCode) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIcmpv6EchoCode) SetMetricGroup(value string) PatternFlowIcmpv6EchoCode {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoCode) Increment() PatternFlowIcmpv6EchoCodeCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIcmpv6EchoCodeCounter{}
	}
	return &patternFlowIcmpv6EchoCodeCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIcmpv6EchoCode) Decrement() PatternFlowIcmpv6EchoCodeCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIcmpv6EchoCodeCounter{}
	}
	return &patternFlowIcmpv6EchoCodeCounter{obj: obj.obj.Decrement}

}

type patternFlowIcmpv6EchoIdentifier struct {
	obj *snappipb.PatternFlowIcmpv6EchoIdentifier
}

func (obj *patternFlowIcmpv6EchoIdentifier) msg() *snappipb.PatternFlowIcmpv6EchoIdentifier {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoIdentifier) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpv6EchoIdentifier) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpv6EchoIdentifier interface {
	msg() *snappipb.PatternFlowIcmpv6EchoIdentifier
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIcmpv6EchoIdentifier
	Values() []int32
	SetValues(value []int32) PatternFlowIcmpv6EchoIdentifier
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIcmpv6EchoIdentifier
	Increment() PatternFlowIcmpv6EchoIdentifierCounter
	Decrement() PatternFlowIcmpv6EchoIdentifierCounter
}

func (obj *patternFlowIcmpv6EchoIdentifier) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIcmpv6EchoIdentifier) SetValue(value int32) PatternFlowIcmpv6EchoIdentifier {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoIdentifier) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIcmpv6EchoIdentifier) SetValues(value []int32) PatternFlowIcmpv6EchoIdentifier {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIcmpv6EchoIdentifier) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIcmpv6EchoIdentifier) SetMetricGroup(value string) PatternFlowIcmpv6EchoIdentifier {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoIdentifier) Increment() PatternFlowIcmpv6EchoIdentifierCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIcmpv6EchoIdentifierCounter{}
	}
	return &patternFlowIcmpv6EchoIdentifierCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIcmpv6EchoIdentifier) Decrement() PatternFlowIcmpv6EchoIdentifierCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIcmpv6EchoIdentifierCounter{}
	}
	return &patternFlowIcmpv6EchoIdentifierCounter{obj: obj.obj.Decrement}

}

type patternFlowIcmpv6EchoSequenceNumber struct {
	obj *snappipb.PatternFlowIcmpv6EchoSequenceNumber
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) msg() *snappipb.PatternFlowIcmpv6EchoSequenceNumber {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpv6EchoSequenceNumber interface {
	msg() *snappipb.PatternFlowIcmpv6EchoSequenceNumber
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIcmpv6EchoSequenceNumber
	Values() []int32
	SetValues(value []int32) PatternFlowIcmpv6EchoSequenceNumber
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIcmpv6EchoSequenceNumber
	Increment() PatternFlowIcmpv6EchoSequenceNumberCounter
	Decrement() PatternFlowIcmpv6EchoSequenceNumberCounter
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) SetValue(value int32) PatternFlowIcmpv6EchoSequenceNumber {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) SetValues(value []int32) PatternFlowIcmpv6EchoSequenceNumber {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) SetMetricGroup(value string) PatternFlowIcmpv6EchoSequenceNumber {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) Increment() PatternFlowIcmpv6EchoSequenceNumberCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIcmpv6EchoSequenceNumberCounter{}
	}
	return &patternFlowIcmpv6EchoSequenceNumberCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIcmpv6EchoSequenceNumber) Decrement() PatternFlowIcmpv6EchoSequenceNumberCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIcmpv6EchoSequenceNumberCounter{}
	}
	return &patternFlowIcmpv6EchoSequenceNumberCounter{obj: obj.obj.Decrement}

}

type patternFlowIcmpv6EchoChecksum struct {
	obj *snappipb.PatternFlowIcmpv6EchoChecksum
}

func (obj *patternFlowIcmpv6EchoChecksum) msg() *snappipb.PatternFlowIcmpv6EchoChecksum {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoChecksum) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpv6EchoChecksum) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpv6EchoChecksum interface {
	msg() *snappipb.PatternFlowIcmpv6EchoChecksum
	Yaml() string
	Json() string
	Custom() int32
	SetCustom(value int32) PatternFlowIcmpv6EchoChecksum
}

func (obj *patternFlowIcmpv6EchoChecksum) Custom() int32 {
	return *obj.obj.Custom
}

func (obj *patternFlowIcmpv6EchoChecksum) SetCustom(value int32) PatternFlowIcmpv6EchoChecksum {
	obj.obj.Custom = &value
	return obj
}

type patternFlowPppAddressCounter struct {
	obj *snappipb.PatternFlowPppAddressCounter
}

func (obj *patternFlowPppAddressCounter) msg() *snappipb.PatternFlowPppAddressCounter {
	return obj.obj
}

func (obj *patternFlowPppAddressCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPppAddressCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPppAddressCounter interface {
	msg() *snappipb.PatternFlowPppAddressCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPppAddressCounter
	Step() int32
	SetStep(value int32) PatternFlowPppAddressCounter
	Count() int32
	SetCount(value int32) PatternFlowPppAddressCounter
}

func (obj *patternFlowPppAddressCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPppAddressCounter) SetStart(value int32) PatternFlowPppAddressCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPppAddressCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPppAddressCounter) SetStep(value int32) PatternFlowPppAddressCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPppAddressCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPppAddressCounter) SetCount(value int32) PatternFlowPppAddressCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPppControlCounter struct {
	obj *snappipb.PatternFlowPppControlCounter
}

func (obj *patternFlowPppControlCounter) msg() *snappipb.PatternFlowPppControlCounter {
	return obj.obj
}

func (obj *patternFlowPppControlCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPppControlCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPppControlCounter interface {
	msg() *snappipb.PatternFlowPppControlCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPppControlCounter
	Step() int32
	SetStep(value int32) PatternFlowPppControlCounter
	Count() int32
	SetCount(value int32) PatternFlowPppControlCounter
}

func (obj *patternFlowPppControlCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPppControlCounter) SetStart(value int32) PatternFlowPppControlCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPppControlCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPppControlCounter) SetStep(value int32) PatternFlowPppControlCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPppControlCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPppControlCounter) SetCount(value int32) PatternFlowPppControlCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowPppProtocolTypeCounter struct {
	obj *snappipb.PatternFlowPppProtocolTypeCounter
}

func (obj *patternFlowPppProtocolTypeCounter) msg() *snappipb.PatternFlowPppProtocolTypeCounter {
	return obj.obj
}

func (obj *patternFlowPppProtocolTypeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowPppProtocolTypeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowPppProtocolTypeCounter interface {
	msg() *snappipb.PatternFlowPppProtocolTypeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowPppProtocolTypeCounter
	Step() int32
	SetStep(value int32) PatternFlowPppProtocolTypeCounter
	Count() int32
	SetCount(value int32) PatternFlowPppProtocolTypeCounter
}

func (obj *patternFlowPppProtocolTypeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowPppProtocolTypeCounter) SetStart(value int32) PatternFlowPppProtocolTypeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowPppProtocolTypeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowPppProtocolTypeCounter) SetStep(value int32) PatternFlowPppProtocolTypeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowPppProtocolTypeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowPppProtocolTypeCounter) SetCount(value int32) PatternFlowPppProtocolTypeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIgmpv1VersionCounter struct {
	obj *snappipb.PatternFlowIgmpv1VersionCounter
}

func (obj *patternFlowIgmpv1VersionCounter) msg() *snappipb.PatternFlowIgmpv1VersionCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv1VersionCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIgmpv1VersionCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIgmpv1VersionCounter interface {
	msg() *snappipb.PatternFlowIgmpv1VersionCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIgmpv1VersionCounter
	Step() int32
	SetStep(value int32) PatternFlowIgmpv1VersionCounter
	Count() int32
	SetCount(value int32) PatternFlowIgmpv1VersionCounter
}

func (obj *patternFlowIgmpv1VersionCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIgmpv1VersionCounter) SetStart(value int32) PatternFlowIgmpv1VersionCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIgmpv1VersionCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIgmpv1VersionCounter) SetStep(value int32) PatternFlowIgmpv1VersionCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIgmpv1VersionCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIgmpv1VersionCounter) SetCount(value int32) PatternFlowIgmpv1VersionCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIgmpv1TypeCounter struct {
	obj *snappipb.PatternFlowIgmpv1TypeCounter
}

func (obj *patternFlowIgmpv1TypeCounter) msg() *snappipb.PatternFlowIgmpv1TypeCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv1TypeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIgmpv1TypeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIgmpv1TypeCounter interface {
	msg() *snappipb.PatternFlowIgmpv1TypeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIgmpv1TypeCounter
	Step() int32
	SetStep(value int32) PatternFlowIgmpv1TypeCounter
	Count() int32
	SetCount(value int32) PatternFlowIgmpv1TypeCounter
}

func (obj *patternFlowIgmpv1TypeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIgmpv1TypeCounter) SetStart(value int32) PatternFlowIgmpv1TypeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIgmpv1TypeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIgmpv1TypeCounter) SetStep(value int32) PatternFlowIgmpv1TypeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIgmpv1TypeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIgmpv1TypeCounter) SetCount(value int32) PatternFlowIgmpv1TypeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIgmpv1UnusedCounter struct {
	obj *snappipb.PatternFlowIgmpv1UnusedCounter
}

func (obj *patternFlowIgmpv1UnusedCounter) msg() *snappipb.PatternFlowIgmpv1UnusedCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv1UnusedCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIgmpv1UnusedCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIgmpv1UnusedCounter interface {
	msg() *snappipb.PatternFlowIgmpv1UnusedCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIgmpv1UnusedCounter
	Step() int32
	SetStep(value int32) PatternFlowIgmpv1UnusedCounter
	Count() int32
	SetCount(value int32) PatternFlowIgmpv1UnusedCounter
}

func (obj *patternFlowIgmpv1UnusedCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIgmpv1UnusedCounter) SetStart(value int32) PatternFlowIgmpv1UnusedCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIgmpv1UnusedCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIgmpv1UnusedCounter) SetStep(value int32) PatternFlowIgmpv1UnusedCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIgmpv1UnusedCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIgmpv1UnusedCounter) SetCount(value int32) PatternFlowIgmpv1UnusedCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIgmpv1GroupAddressCounter struct {
	obj *snappipb.PatternFlowIgmpv1GroupAddressCounter
}

func (obj *patternFlowIgmpv1GroupAddressCounter) msg() *snappipb.PatternFlowIgmpv1GroupAddressCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv1GroupAddressCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIgmpv1GroupAddressCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIgmpv1GroupAddressCounter interface {
	msg() *snappipb.PatternFlowIgmpv1GroupAddressCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternFlowIgmpv1GroupAddressCounter
	Step() string
	SetStep(value string) PatternFlowIgmpv1GroupAddressCounter
	Count() int32
	SetCount(value int32) PatternFlowIgmpv1GroupAddressCounter
}

func (obj *patternFlowIgmpv1GroupAddressCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternFlowIgmpv1GroupAddressCounter) SetStart(value string) PatternFlowIgmpv1GroupAddressCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIgmpv1GroupAddressCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternFlowIgmpv1GroupAddressCounter) SetStep(value string) PatternFlowIgmpv1GroupAddressCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIgmpv1GroupAddressCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIgmpv1GroupAddressCounter) SetCount(value int32) PatternFlowIgmpv1GroupAddressCounter {
	obj.obj.Count = &value
	return obj
}

type deviceBgpSrTePolicyNextHop struct {
	obj *snappipb.DeviceBgpSrTePolicyNextHop
}

func (obj *deviceBgpSrTePolicyNextHop) msg() *snappipb.DeviceBgpSrTePolicyNextHop {
	return obj.obj
}

func (obj *deviceBgpSrTePolicyNextHop) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpSrTePolicyNextHop) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpSrTePolicyNextHop interface {
	msg() *snappipb.DeviceBgpSrTePolicyNextHop
	Yaml() string
	Json() string
	Ipv4Address() string
	SetIpv4Address(value string) DeviceBgpSrTePolicyNextHop
	Ipv6Address() string
	SetIpv6Address(value string) DeviceBgpSrTePolicyNextHop
}

func (obj *deviceBgpSrTePolicyNextHop) Ipv4Address() string {
	return *obj.obj.Ipv4Address
}

func (obj *deviceBgpSrTePolicyNextHop) SetIpv4Address(value string) DeviceBgpSrTePolicyNextHop {
	obj.obj.Ipv4Address = &value
	return obj
}

func (obj *deviceBgpSrTePolicyNextHop) Ipv6Address() string {
	return *obj.obj.Ipv6Address
}

func (obj *deviceBgpSrTePolicyNextHop) SetIpv6Address(value string) DeviceBgpSrTePolicyNextHop {
	obj.obj.Ipv6Address = &value
	return obj
}

type deviceBgpAddPath struct {
	obj *snappipb.DeviceBgpAddPath
}

func (obj *deviceBgpAddPath) msg() *snappipb.DeviceBgpAddPath {
	return obj.obj
}

func (obj *deviceBgpAddPath) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpAddPath) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpAddPath interface {
	msg() *snappipb.DeviceBgpAddPath
	Yaml() string
	Json() string
	PathId() int32
	SetPathId(value int32) DeviceBgpAddPath
}

func (obj *deviceBgpAddPath) PathId() int32 {
	return *obj.obj.PathId
}

func (obj *deviceBgpAddPath) SetPathId(value int32) DeviceBgpAddPath {
	obj.obj.PathId = &value
	return obj
}

type deviceBgpAsPath struct {
	obj *snappipb.DeviceBgpAsPath
}

func (obj *deviceBgpAsPath) msg() *snappipb.DeviceBgpAsPath {
	return obj.obj
}

func (obj *deviceBgpAsPath) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpAsPath) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpAsPath interface {
	msg() *snappipb.DeviceBgpAsPath
	Yaml() string
	Json() string
	OverridePeerAsSetMode() bool
	SetOverridePeerAsSetMode(value bool) DeviceBgpAsPath
	AsPathSegments() []DeviceBgpAsPathSegment
	NewAsPathSegments() DeviceBgpAsPathSegment
}

func (obj *deviceBgpAsPath) OverridePeerAsSetMode() bool {
	return *obj.obj.OverridePeerAsSetMode
}

func (obj *deviceBgpAsPath) SetOverridePeerAsSetMode(value bool) DeviceBgpAsPath {
	obj.obj.OverridePeerAsSetMode = &value
	return obj
}

func (obj *deviceBgpAsPath) AsPathSegments() []DeviceBgpAsPathSegment {
	if obj.obj.AsPathSegments == nil {
		obj.obj.AsPathSegments = make([]*snappipb.DeviceBgpAsPathSegment, 0)
	}
	values := make([]DeviceBgpAsPathSegment, 0)
	for _, item := range obj.obj.AsPathSegments {
		values = append(values, &deviceBgpAsPathSegment{obj: item})
	}
	return values

}

func (obj *deviceBgpAsPath) NewAsPathSegments() DeviceBgpAsPathSegment {
	if obj.obj.AsPathSegments == nil {
		obj.obj.AsPathSegments = make([]*snappipb.DeviceBgpAsPathSegment, 0)
	}
	slice := append(obj.obj.AsPathSegments, &snappipb.DeviceBgpAsPathSegment{})
	obj.obj.AsPathSegments = slice
	return &deviceBgpAsPathSegment{obj: slice[len(slice)-1]}
}

type deviceBgpTunnelTlv struct {
	obj *snappipb.DeviceBgpTunnelTlv
}

func (obj *deviceBgpTunnelTlv) msg() *snappipb.DeviceBgpTunnelTlv {
	return obj.obj
}

func (obj *deviceBgpTunnelTlv) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpTunnelTlv) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpTunnelTlv interface {
	msg() *snappipb.DeviceBgpTunnelTlv
	Yaml() string
	Json() string
	SegmentLists() []DeviceBgpSegmentList
	NewSegmentLists() DeviceBgpSegmentList
	RemoteEndpointSubTlv() DeviceBgpRemoteEndpointSubTlv
	PreferenceSubTlv() DeviceBgpPreferenceSubTlv
	BindingSubTlv() DeviceBgpBindingSubTlv
	ExplicitNullLabelPolicySubTlv() DeviceBgpExplicitNullLabelPolicySubTlv
	Active() bool
	SetActive(value bool) DeviceBgpTunnelTlv
}

func (obj *deviceBgpTunnelTlv) SegmentLists() []DeviceBgpSegmentList {
	if obj.obj.SegmentLists == nil {
		obj.obj.SegmentLists = make([]*snappipb.DeviceBgpSegmentList, 0)
	}
	values := make([]DeviceBgpSegmentList, 0)
	for _, item := range obj.obj.SegmentLists {
		values = append(values, &deviceBgpSegmentList{obj: item})
	}
	return values

}

func (obj *deviceBgpTunnelTlv) NewSegmentLists() DeviceBgpSegmentList {
	if obj.obj.SegmentLists == nil {
		obj.obj.SegmentLists = make([]*snappipb.DeviceBgpSegmentList, 0)
	}
	slice := append(obj.obj.SegmentLists, &snappipb.DeviceBgpSegmentList{})
	obj.obj.SegmentLists = slice
	return &deviceBgpSegmentList{obj: slice[len(slice)-1]}
}

func (obj *deviceBgpTunnelTlv) RemoteEndpointSubTlv() DeviceBgpRemoteEndpointSubTlv {
	if obj.obj.RemoteEndpointSubTlv == nil {
		obj.obj.RemoteEndpointSubTlv = &snappipb.DeviceBgpRemoteEndpointSubTlv{}
	}
	return &deviceBgpRemoteEndpointSubTlv{obj: obj.obj.RemoteEndpointSubTlv}

}

func (obj *deviceBgpTunnelTlv) PreferenceSubTlv() DeviceBgpPreferenceSubTlv {
	if obj.obj.PreferenceSubTlv == nil {
		obj.obj.PreferenceSubTlv = &snappipb.DeviceBgpPreferenceSubTlv{}
	}
	return &deviceBgpPreferenceSubTlv{obj: obj.obj.PreferenceSubTlv}

}

func (obj *deviceBgpTunnelTlv) BindingSubTlv() DeviceBgpBindingSubTlv {
	if obj.obj.BindingSubTlv == nil {
		obj.obj.BindingSubTlv = &snappipb.DeviceBgpBindingSubTlv{}
	}
	return &deviceBgpBindingSubTlv{obj: obj.obj.BindingSubTlv}

}

func (obj *deviceBgpTunnelTlv) ExplicitNullLabelPolicySubTlv() DeviceBgpExplicitNullLabelPolicySubTlv {
	if obj.obj.ExplicitNullLabelPolicySubTlv == nil {
		obj.obj.ExplicitNullLabelPolicySubTlv = &snappipb.DeviceBgpExplicitNullLabelPolicySubTlv{}
	}
	return &deviceBgpExplicitNullLabelPolicySubTlv{obj: obj.obj.ExplicitNullLabelPolicySubTlv}

}

func (obj *deviceBgpTunnelTlv) Active() bool {
	return *obj.obj.Active
}

func (obj *deviceBgpTunnelTlv) SetActive(value bool) DeviceBgpTunnelTlv {
	obj.obj.Active = &value
	return obj
}

type deviceBgpCommunity struct {
	obj *snappipb.DeviceBgpCommunity
}

func (obj *deviceBgpCommunity) msg() *snappipb.DeviceBgpCommunity {
	return obj.obj
}

func (obj *deviceBgpCommunity) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpCommunity) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpCommunity interface {
	msg() *snappipb.DeviceBgpCommunity
	Yaml() string
	Json() string
	AsNumber() int32
	SetAsNumber(value int32) DeviceBgpCommunity
	AsCustom() int32
	SetAsCustom(value int32) DeviceBgpCommunity
}

func (obj *deviceBgpCommunity) AsNumber() int32 {
	return *obj.obj.AsNumber
}

func (obj *deviceBgpCommunity) SetAsNumber(value int32) DeviceBgpCommunity {
	obj.obj.AsNumber = &value
	return obj
}

func (obj *deviceBgpCommunity) AsCustom() int32 {
	return *obj.obj.AsCustom
}

func (obj *deviceBgpCommunity) SetAsCustom(value int32) DeviceBgpCommunity {
	obj.obj.AsCustom = &value
	return obj
}

type deviceBgpv4RouteAddress struct {
	obj *snappipb.DeviceBgpv4RouteAddress
}

func (obj *deviceBgpv4RouteAddress) msg() *snappipb.DeviceBgpv4RouteAddress {
	return obj.obj
}

func (obj *deviceBgpv4RouteAddress) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpv4RouteAddress) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpv4RouteAddress interface {
	msg() *snappipb.DeviceBgpv4RouteAddress
	Yaml() string
	Json() string
	Address() string
	SetAddress(value string) DeviceBgpv4RouteAddress
	Prefix() int32
	SetPrefix(value int32) DeviceBgpv4RouteAddress
	Count() int32
	SetCount(value int32) DeviceBgpv4RouteAddress
	Step() int32
	SetStep(value int32) DeviceBgpv4RouteAddress
}

func (obj *deviceBgpv4RouteAddress) Address() string {
	return obj.obj.Address
}

func (obj *deviceBgpv4RouteAddress) SetAddress(value string) DeviceBgpv4RouteAddress {
	obj.obj.Address = value
	return obj
}

func (obj *deviceBgpv4RouteAddress) Prefix() int32 {
	return *obj.obj.Prefix
}

func (obj *deviceBgpv4RouteAddress) SetPrefix(value int32) DeviceBgpv4RouteAddress {
	obj.obj.Prefix = &value
	return obj
}

func (obj *deviceBgpv4RouteAddress) Count() int32 {
	return *obj.obj.Count
}

func (obj *deviceBgpv4RouteAddress) SetCount(value int32) DeviceBgpv4RouteAddress {
	obj.obj.Count = &value
	return obj
}

func (obj *deviceBgpv4RouteAddress) Step() int32 {
	return *obj.obj.Step
}

func (obj *deviceBgpv4RouteAddress) SetStep(value int32) DeviceBgpv4RouteAddress {
	obj.obj.Step = &value
	return obj
}

type deviceBgpRouteAdvanced struct {
	obj *snappipb.DeviceBgpRouteAdvanced
}

func (obj *deviceBgpRouteAdvanced) msg() *snappipb.DeviceBgpRouteAdvanced {
	return obj.obj
}

func (obj *deviceBgpRouteAdvanced) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpRouteAdvanced) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpRouteAdvanced interface {
	msg() *snappipb.DeviceBgpRouteAdvanced
	Yaml() string
	Json() string
	MultiExitDiscriminator() int32
	SetMultiExitDiscriminator(value int32) DeviceBgpRouteAdvanced
}

func (obj *deviceBgpRouteAdvanced) MultiExitDiscriminator() int32 {
	return *obj.obj.MultiExitDiscriminator
}

func (obj *deviceBgpRouteAdvanced) SetMultiExitDiscriminator(value int32) DeviceBgpRouteAdvanced {
	obj.obj.MultiExitDiscriminator = &value
	return obj
}

type deviceBgpv6RouteAddress struct {
	obj *snappipb.DeviceBgpv6RouteAddress
}

func (obj *deviceBgpv6RouteAddress) msg() *snappipb.DeviceBgpv6RouteAddress {
	return obj.obj
}

func (obj *deviceBgpv6RouteAddress) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpv6RouteAddress) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpv6RouteAddress interface {
	msg() *snappipb.DeviceBgpv6RouteAddress
	Yaml() string
	Json() string
	Address() string
	SetAddress(value string) DeviceBgpv6RouteAddress
	Prefix() int32
	SetPrefix(value int32) DeviceBgpv6RouteAddress
	Count() int32
	SetCount(value int32) DeviceBgpv6RouteAddress
	Step() int32
	SetStep(value int32) DeviceBgpv6RouteAddress
}

func (obj *deviceBgpv6RouteAddress) Address() string {
	return obj.obj.Address
}

func (obj *deviceBgpv6RouteAddress) SetAddress(value string) DeviceBgpv6RouteAddress {
	obj.obj.Address = value
	return obj
}

func (obj *deviceBgpv6RouteAddress) Prefix() int32 {
	return *obj.obj.Prefix
}

func (obj *deviceBgpv6RouteAddress) SetPrefix(value int32) DeviceBgpv6RouteAddress {
	obj.obj.Prefix = &value
	return obj
}

func (obj *deviceBgpv6RouteAddress) Count() int32 {
	return *obj.obj.Count
}

func (obj *deviceBgpv6RouteAddress) SetCount(value int32) DeviceBgpv6RouteAddress {
	obj.obj.Count = &value
	return obj
}

func (obj *deviceBgpv6RouteAddress) Step() int32 {
	return *obj.obj.Step
}

func (obj *deviceBgpv6RouteAddress) SetStep(value int32) DeviceBgpv6RouteAddress {
	obj.obj.Step = &value
	return obj
}

type patternFlowIpv4PriorityRawCounter struct {
	obj *snappipb.PatternFlowIpv4PriorityRawCounter
}

func (obj *patternFlowIpv4PriorityRawCounter) msg() *snappipb.PatternFlowIpv4PriorityRawCounter {
	return obj.obj
}

func (obj *patternFlowIpv4PriorityRawCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4PriorityRawCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4PriorityRawCounter interface {
	msg() *snappipb.PatternFlowIpv4PriorityRawCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4PriorityRawCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4PriorityRawCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4PriorityRawCounter
}

func (obj *patternFlowIpv4PriorityRawCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4PriorityRawCounter) SetStart(value int32) PatternFlowIpv4PriorityRawCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4PriorityRawCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4PriorityRawCounter) SetStep(value int32) PatternFlowIpv4PriorityRawCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4PriorityRawCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4PriorityRawCounter) SetCount(value int32) PatternFlowIpv4PriorityRawCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4TosPrecedence struct {
	obj *snappipb.PatternFlowIpv4TosPrecedence
}

func (obj *patternFlowIpv4TosPrecedence) msg() *snappipb.PatternFlowIpv4TosPrecedence {
	return obj.obj
}

func (obj *patternFlowIpv4TosPrecedence) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TosPrecedence) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TosPrecedence interface {
	msg() *snappipb.PatternFlowIpv4TosPrecedence
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4TosPrecedence
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4TosPrecedence
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4TosPrecedence
	Increment() PatternFlowIpv4TosPrecedenceCounter
	Decrement() PatternFlowIpv4TosPrecedenceCounter
}

func (obj *patternFlowIpv4TosPrecedence) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4TosPrecedence) SetValue(value int32) PatternFlowIpv4TosPrecedence {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4TosPrecedence) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4TosPrecedence) SetValues(value []int32) PatternFlowIpv4TosPrecedence {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4TosPrecedence) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4TosPrecedence) SetMetricGroup(value string) PatternFlowIpv4TosPrecedence {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4TosPrecedence) Increment() PatternFlowIpv4TosPrecedenceCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4TosPrecedenceCounter{}
	}
	return &patternFlowIpv4TosPrecedenceCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4TosPrecedence) Decrement() PatternFlowIpv4TosPrecedenceCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4TosPrecedenceCounter{}
	}
	return &patternFlowIpv4TosPrecedenceCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4TosDelay struct {
	obj *snappipb.PatternFlowIpv4TosDelay
}

func (obj *patternFlowIpv4TosDelay) msg() *snappipb.PatternFlowIpv4TosDelay {
	return obj.obj
}

func (obj *patternFlowIpv4TosDelay) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TosDelay) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TosDelay interface {
	msg() *snappipb.PatternFlowIpv4TosDelay
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4TosDelay
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4TosDelay
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4TosDelay
	Increment() PatternFlowIpv4TosDelayCounter
	Decrement() PatternFlowIpv4TosDelayCounter
}

func (obj *patternFlowIpv4TosDelay) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4TosDelay) SetValue(value int32) PatternFlowIpv4TosDelay {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4TosDelay) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4TosDelay) SetValues(value []int32) PatternFlowIpv4TosDelay {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4TosDelay) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4TosDelay) SetMetricGroup(value string) PatternFlowIpv4TosDelay {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4TosDelay) Increment() PatternFlowIpv4TosDelayCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4TosDelayCounter{}
	}
	return &patternFlowIpv4TosDelayCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4TosDelay) Decrement() PatternFlowIpv4TosDelayCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4TosDelayCounter{}
	}
	return &patternFlowIpv4TosDelayCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4TosThroughput struct {
	obj *snappipb.PatternFlowIpv4TosThroughput
}

func (obj *patternFlowIpv4TosThroughput) msg() *snappipb.PatternFlowIpv4TosThroughput {
	return obj.obj
}

func (obj *patternFlowIpv4TosThroughput) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TosThroughput) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TosThroughput interface {
	msg() *snappipb.PatternFlowIpv4TosThroughput
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4TosThroughput
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4TosThroughput
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4TosThroughput
	Increment() PatternFlowIpv4TosThroughputCounter
	Decrement() PatternFlowIpv4TosThroughputCounter
}

func (obj *patternFlowIpv4TosThroughput) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4TosThroughput) SetValue(value int32) PatternFlowIpv4TosThroughput {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4TosThroughput) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4TosThroughput) SetValues(value []int32) PatternFlowIpv4TosThroughput {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4TosThroughput) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4TosThroughput) SetMetricGroup(value string) PatternFlowIpv4TosThroughput {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4TosThroughput) Increment() PatternFlowIpv4TosThroughputCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4TosThroughputCounter{}
	}
	return &patternFlowIpv4TosThroughputCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4TosThroughput) Decrement() PatternFlowIpv4TosThroughputCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4TosThroughputCounter{}
	}
	return &patternFlowIpv4TosThroughputCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4TosReliability struct {
	obj *snappipb.PatternFlowIpv4TosReliability
}

func (obj *patternFlowIpv4TosReliability) msg() *snappipb.PatternFlowIpv4TosReliability {
	return obj.obj
}

func (obj *patternFlowIpv4TosReliability) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TosReliability) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TosReliability interface {
	msg() *snappipb.PatternFlowIpv4TosReliability
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4TosReliability
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4TosReliability
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4TosReliability
	Increment() PatternFlowIpv4TosReliabilityCounter
	Decrement() PatternFlowIpv4TosReliabilityCounter
}

func (obj *patternFlowIpv4TosReliability) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4TosReliability) SetValue(value int32) PatternFlowIpv4TosReliability {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4TosReliability) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4TosReliability) SetValues(value []int32) PatternFlowIpv4TosReliability {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4TosReliability) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4TosReliability) SetMetricGroup(value string) PatternFlowIpv4TosReliability {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4TosReliability) Increment() PatternFlowIpv4TosReliabilityCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4TosReliabilityCounter{}
	}
	return &patternFlowIpv4TosReliabilityCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4TosReliability) Decrement() PatternFlowIpv4TosReliabilityCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4TosReliabilityCounter{}
	}
	return &patternFlowIpv4TosReliabilityCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4TosMonetary struct {
	obj *snappipb.PatternFlowIpv4TosMonetary
}

func (obj *patternFlowIpv4TosMonetary) msg() *snappipb.PatternFlowIpv4TosMonetary {
	return obj.obj
}

func (obj *patternFlowIpv4TosMonetary) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TosMonetary) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TosMonetary interface {
	msg() *snappipb.PatternFlowIpv4TosMonetary
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4TosMonetary
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4TosMonetary
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4TosMonetary
	Increment() PatternFlowIpv4TosMonetaryCounter
	Decrement() PatternFlowIpv4TosMonetaryCounter
}

func (obj *patternFlowIpv4TosMonetary) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4TosMonetary) SetValue(value int32) PatternFlowIpv4TosMonetary {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4TosMonetary) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4TosMonetary) SetValues(value []int32) PatternFlowIpv4TosMonetary {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4TosMonetary) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4TosMonetary) SetMetricGroup(value string) PatternFlowIpv4TosMonetary {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4TosMonetary) Increment() PatternFlowIpv4TosMonetaryCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4TosMonetaryCounter{}
	}
	return &patternFlowIpv4TosMonetaryCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4TosMonetary) Decrement() PatternFlowIpv4TosMonetaryCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4TosMonetaryCounter{}
	}
	return &patternFlowIpv4TosMonetaryCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4TosUnused struct {
	obj *snappipb.PatternFlowIpv4TosUnused
}

func (obj *patternFlowIpv4TosUnused) msg() *snappipb.PatternFlowIpv4TosUnused {
	return obj.obj
}

func (obj *patternFlowIpv4TosUnused) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TosUnused) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TosUnused interface {
	msg() *snappipb.PatternFlowIpv4TosUnused
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4TosUnused
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4TosUnused
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4TosUnused
	Increment() PatternFlowIpv4TosUnusedCounter
	Decrement() PatternFlowIpv4TosUnusedCounter
}

func (obj *patternFlowIpv4TosUnused) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4TosUnused) SetValue(value int32) PatternFlowIpv4TosUnused {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4TosUnused) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4TosUnused) SetValues(value []int32) PatternFlowIpv4TosUnused {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4TosUnused) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4TosUnused) SetMetricGroup(value string) PatternFlowIpv4TosUnused {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4TosUnused) Increment() PatternFlowIpv4TosUnusedCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4TosUnusedCounter{}
	}
	return &patternFlowIpv4TosUnusedCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4TosUnused) Decrement() PatternFlowIpv4TosUnusedCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4TosUnusedCounter{}
	}
	return &patternFlowIpv4TosUnusedCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4DscpPhb struct {
	obj *snappipb.PatternFlowIpv4DscpPhb
}

func (obj *patternFlowIpv4DscpPhb) msg() *snappipb.PatternFlowIpv4DscpPhb {
	return obj.obj
}

func (obj *patternFlowIpv4DscpPhb) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4DscpPhb) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4DscpPhb interface {
	msg() *snappipb.PatternFlowIpv4DscpPhb
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4DscpPhb
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4DscpPhb
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4DscpPhb
	Increment() PatternFlowIpv4DscpPhbCounter
	Decrement() PatternFlowIpv4DscpPhbCounter
}

func (obj *patternFlowIpv4DscpPhb) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4DscpPhb) SetValue(value int32) PatternFlowIpv4DscpPhb {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4DscpPhb) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4DscpPhb) SetValues(value []int32) PatternFlowIpv4DscpPhb {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4DscpPhb) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4DscpPhb) SetMetricGroup(value string) PatternFlowIpv4DscpPhb {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4DscpPhb) Increment() PatternFlowIpv4DscpPhbCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4DscpPhbCounter{}
	}
	return &patternFlowIpv4DscpPhbCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4DscpPhb) Decrement() PatternFlowIpv4DscpPhbCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4DscpPhbCounter{}
	}
	return &patternFlowIpv4DscpPhbCounter{obj: obj.obj.Decrement}

}

type patternFlowIpv4DscpEcn struct {
	obj *snappipb.PatternFlowIpv4DscpEcn
}

func (obj *patternFlowIpv4DscpEcn) msg() *snappipb.PatternFlowIpv4DscpEcn {
	return obj.obj
}

func (obj *patternFlowIpv4DscpEcn) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4DscpEcn) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4DscpEcn interface {
	msg() *snappipb.PatternFlowIpv4DscpEcn
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternFlowIpv4DscpEcn
	Values() []int32
	SetValues(value []int32) PatternFlowIpv4DscpEcn
	MetricGroup() string
	SetMetricGroup(value string) PatternFlowIpv4DscpEcn
	Increment() PatternFlowIpv4DscpEcnCounter
	Decrement() PatternFlowIpv4DscpEcnCounter
}

func (obj *patternFlowIpv4DscpEcn) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternFlowIpv4DscpEcn) SetValue(value int32) PatternFlowIpv4DscpEcn {
	obj.obj.Value = &value
	return obj
}

func (obj *patternFlowIpv4DscpEcn) Values() []int32 {
	return obj.obj.Values
}

func (obj *patternFlowIpv4DscpEcn) SetValues(value []int32) PatternFlowIpv4DscpEcn {
	obj.obj.Values = value
	return obj
}

func (obj *patternFlowIpv4DscpEcn) MetricGroup() string {
	return *obj.obj.MetricGroup
}

func (obj *patternFlowIpv4DscpEcn) SetMetricGroup(value string) PatternFlowIpv4DscpEcn {
	obj.obj.MetricGroup = &value
	return obj
}

func (obj *patternFlowIpv4DscpEcn) Increment() PatternFlowIpv4DscpEcnCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &snappipb.PatternFlowIpv4DscpEcnCounter{}
	}
	return &patternFlowIpv4DscpEcnCounter{obj: obj.obj.Increment}

}

func (obj *patternFlowIpv4DscpEcn) Decrement() PatternFlowIpv4DscpEcnCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &snappipb.PatternFlowIpv4DscpEcnCounter{}
	}
	return &patternFlowIpv4DscpEcnCounter{obj: obj.obj.Decrement}

}

type patternFlowGtpExtensionExtensionLengthCounter struct {
	obj *snappipb.PatternFlowGtpExtensionExtensionLengthCounter
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) msg() *snappipb.PatternFlowGtpExtensionExtensionLengthCounter {
	return obj.obj
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpExtensionExtensionLengthCounter interface {
	msg() *snappipb.PatternFlowGtpExtensionExtensionLengthCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpExtensionExtensionLengthCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpExtensionExtensionLengthCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpExtensionExtensionLengthCounter
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) SetStart(value int32) PatternFlowGtpExtensionExtensionLengthCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) SetStep(value int32) PatternFlowGtpExtensionExtensionLengthCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) SetCount(value int32) PatternFlowGtpExtensionExtensionLengthCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpExtensionContentsCounter struct {
	obj *snappipb.PatternFlowGtpExtensionContentsCounter
}

func (obj *patternFlowGtpExtensionContentsCounter) msg() *snappipb.PatternFlowGtpExtensionContentsCounter {
	return obj.obj
}

func (obj *patternFlowGtpExtensionContentsCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpExtensionContentsCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpExtensionContentsCounter interface {
	msg() *snappipb.PatternFlowGtpExtensionContentsCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpExtensionContentsCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpExtensionContentsCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpExtensionContentsCounter
}

func (obj *patternFlowGtpExtensionContentsCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpExtensionContentsCounter) SetStart(value int32) PatternFlowGtpExtensionContentsCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpExtensionContentsCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpExtensionContentsCounter) SetStep(value int32) PatternFlowGtpExtensionContentsCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpExtensionContentsCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpExtensionContentsCounter) SetCount(value int32) PatternFlowGtpExtensionContentsCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowGtpExtensionNextExtensionHeaderCounter struct {
	obj *snappipb.PatternFlowGtpExtensionNextExtensionHeaderCounter
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) msg() *snappipb.PatternFlowGtpExtensionNextExtensionHeaderCounter {
	return obj.obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowGtpExtensionNextExtensionHeaderCounter interface {
	msg() *snappipb.PatternFlowGtpExtensionNextExtensionHeaderCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowGtpExtensionNextExtensionHeaderCounter
	Step() int32
	SetStep(value int32) PatternFlowGtpExtensionNextExtensionHeaderCounter
	Count() int32
	SetCount(value int32) PatternFlowGtpExtensionNextExtensionHeaderCounter
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) SetStart(value int32) PatternFlowGtpExtensionNextExtensionHeaderCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) SetStep(value int32) PatternFlowGtpExtensionNextExtensionHeaderCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderCounter) SetCount(value int32) PatternFlowGtpExtensionNextExtensionHeaderCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIcmpEchoTypeCounter struct {
	obj *snappipb.PatternFlowIcmpEchoTypeCounter
}

func (obj *patternFlowIcmpEchoTypeCounter) msg() *snappipb.PatternFlowIcmpEchoTypeCounter {
	return obj.obj
}

func (obj *patternFlowIcmpEchoTypeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpEchoTypeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpEchoTypeCounter interface {
	msg() *snappipb.PatternFlowIcmpEchoTypeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIcmpEchoTypeCounter
	Step() int32
	SetStep(value int32) PatternFlowIcmpEchoTypeCounter
	Count() int32
	SetCount(value int32) PatternFlowIcmpEchoTypeCounter
}

func (obj *patternFlowIcmpEchoTypeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIcmpEchoTypeCounter) SetStart(value int32) PatternFlowIcmpEchoTypeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIcmpEchoTypeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIcmpEchoTypeCounter) SetStep(value int32) PatternFlowIcmpEchoTypeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIcmpEchoTypeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIcmpEchoTypeCounter) SetCount(value int32) PatternFlowIcmpEchoTypeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIcmpEchoCodeCounter struct {
	obj *snappipb.PatternFlowIcmpEchoCodeCounter
}

func (obj *patternFlowIcmpEchoCodeCounter) msg() *snappipb.PatternFlowIcmpEchoCodeCounter {
	return obj.obj
}

func (obj *patternFlowIcmpEchoCodeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpEchoCodeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpEchoCodeCounter interface {
	msg() *snappipb.PatternFlowIcmpEchoCodeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIcmpEchoCodeCounter
	Step() int32
	SetStep(value int32) PatternFlowIcmpEchoCodeCounter
	Count() int32
	SetCount(value int32) PatternFlowIcmpEchoCodeCounter
}

func (obj *patternFlowIcmpEchoCodeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIcmpEchoCodeCounter) SetStart(value int32) PatternFlowIcmpEchoCodeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIcmpEchoCodeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIcmpEchoCodeCounter) SetStep(value int32) PatternFlowIcmpEchoCodeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIcmpEchoCodeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIcmpEchoCodeCounter) SetCount(value int32) PatternFlowIcmpEchoCodeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIcmpEchoIdentifierCounter struct {
	obj *snappipb.PatternFlowIcmpEchoIdentifierCounter
}

func (obj *patternFlowIcmpEchoIdentifierCounter) msg() *snappipb.PatternFlowIcmpEchoIdentifierCounter {
	return obj.obj
}

func (obj *patternFlowIcmpEchoIdentifierCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpEchoIdentifierCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpEchoIdentifierCounter interface {
	msg() *snappipb.PatternFlowIcmpEchoIdentifierCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIcmpEchoIdentifierCounter
	Step() int32
	SetStep(value int32) PatternFlowIcmpEchoIdentifierCounter
	Count() int32
	SetCount(value int32) PatternFlowIcmpEchoIdentifierCounter
}

func (obj *patternFlowIcmpEchoIdentifierCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIcmpEchoIdentifierCounter) SetStart(value int32) PatternFlowIcmpEchoIdentifierCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIcmpEchoIdentifierCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIcmpEchoIdentifierCounter) SetStep(value int32) PatternFlowIcmpEchoIdentifierCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIcmpEchoIdentifierCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIcmpEchoIdentifierCounter) SetCount(value int32) PatternFlowIcmpEchoIdentifierCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIcmpEchoSequenceNumberCounter struct {
	obj *snappipb.PatternFlowIcmpEchoSequenceNumberCounter
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) msg() *snappipb.PatternFlowIcmpEchoSequenceNumberCounter {
	return obj.obj
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpEchoSequenceNumberCounter interface {
	msg() *snappipb.PatternFlowIcmpEchoSequenceNumberCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIcmpEchoSequenceNumberCounter
	Step() int32
	SetStep(value int32) PatternFlowIcmpEchoSequenceNumberCounter
	Count() int32
	SetCount(value int32) PatternFlowIcmpEchoSequenceNumberCounter
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) SetStart(value int32) PatternFlowIcmpEchoSequenceNumberCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) SetStep(value int32) PatternFlowIcmpEchoSequenceNumberCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) SetCount(value int32) PatternFlowIcmpEchoSequenceNumberCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIcmpv6EchoTypeCounter struct {
	obj *snappipb.PatternFlowIcmpv6EchoTypeCounter
}

func (obj *patternFlowIcmpv6EchoTypeCounter) msg() *snappipb.PatternFlowIcmpv6EchoTypeCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoTypeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpv6EchoTypeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpv6EchoTypeCounter interface {
	msg() *snappipb.PatternFlowIcmpv6EchoTypeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIcmpv6EchoTypeCounter
	Step() int32
	SetStep(value int32) PatternFlowIcmpv6EchoTypeCounter
	Count() int32
	SetCount(value int32) PatternFlowIcmpv6EchoTypeCounter
}

func (obj *patternFlowIcmpv6EchoTypeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIcmpv6EchoTypeCounter) SetStart(value int32) PatternFlowIcmpv6EchoTypeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoTypeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIcmpv6EchoTypeCounter) SetStep(value int32) PatternFlowIcmpv6EchoTypeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoTypeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIcmpv6EchoTypeCounter) SetCount(value int32) PatternFlowIcmpv6EchoTypeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIcmpv6EchoCodeCounter struct {
	obj *snappipb.PatternFlowIcmpv6EchoCodeCounter
}

func (obj *patternFlowIcmpv6EchoCodeCounter) msg() *snappipb.PatternFlowIcmpv6EchoCodeCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoCodeCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpv6EchoCodeCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpv6EchoCodeCounter interface {
	msg() *snappipb.PatternFlowIcmpv6EchoCodeCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIcmpv6EchoCodeCounter
	Step() int32
	SetStep(value int32) PatternFlowIcmpv6EchoCodeCounter
	Count() int32
	SetCount(value int32) PatternFlowIcmpv6EchoCodeCounter
}

func (obj *patternFlowIcmpv6EchoCodeCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIcmpv6EchoCodeCounter) SetStart(value int32) PatternFlowIcmpv6EchoCodeCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoCodeCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIcmpv6EchoCodeCounter) SetStep(value int32) PatternFlowIcmpv6EchoCodeCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoCodeCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIcmpv6EchoCodeCounter) SetCount(value int32) PatternFlowIcmpv6EchoCodeCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIcmpv6EchoIdentifierCounter struct {
	obj *snappipb.PatternFlowIcmpv6EchoIdentifierCounter
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) msg() *snappipb.PatternFlowIcmpv6EchoIdentifierCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpv6EchoIdentifierCounter interface {
	msg() *snappipb.PatternFlowIcmpv6EchoIdentifierCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIcmpv6EchoIdentifierCounter
	Step() int32
	SetStep(value int32) PatternFlowIcmpv6EchoIdentifierCounter
	Count() int32
	SetCount(value int32) PatternFlowIcmpv6EchoIdentifierCounter
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) SetStart(value int32) PatternFlowIcmpv6EchoIdentifierCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) SetStep(value int32) PatternFlowIcmpv6EchoIdentifierCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) SetCount(value int32) PatternFlowIcmpv6EchoIdentifierCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIcmpv6EchoSequenceNumberCounter struct {
	obj *snappipb.PatternFlowIcmpv6EchoSequenceNumberCounter
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) msg() *snappipb.PatternFlowIcmpv6EchoSequenceNumberCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIcmpv6EchoSequenceNumberCounter interface {
	msg() *snappipb.PatternFlowIcmpv6EchoSequenceNumberCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIcmpv6EchoSequenceNumberCounter
	Step() int32
	SetStep(value int32) PatternFlowIcmpv6EchoSequenceNumberCounter
	Count() int32
	SetCount(value int32) PatternFlowIcmpv6EchoSequenceNumberCounter
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) SetStart(value int32) PatternFlowIcmpv6EchoSequenceNumberCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) SetStep(value int32) PatternFlowIcmpv6EchoSequenceNumberCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) SetCount(value int32) PatternFlowIcmpv6EchoSequenceNumberCounter {
	obj.obj.Count = &value
	return obj
}

type deviceBgpAsPathSegment struct {
	obj *snappipb.DeviceBgpAsPathSegment
}

func (obj *deviceBgpAsPathSegment) msg() *snappipb.DeviceBgpAsPathSegment {
	return obj.obj
}

func (obj *deviceBgpAsPathSegment) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpAsPathSegment) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpAsPathSegment interface {
	msg() *snappipb.DeviceBgpAsPathSegment
	Yaml() string
	Json() string
	AsNumbers() []int32
	SetAsNumbers(value []int32) DeviceBgpAsPathSegment
}

func (obj *deviceBgpAsPathSegment) AsNumbers() []int32 {
	return obj.obj.AsNumbers
}

func (obj *deviceBgpAsPathSegment) SetAsNumbers(value []int32) DeviceBgpAsPathSegment {
	obj.obj.AsNumbers = value
	return obj
}

type deviceBgpSegmentList struct {
	obj *snappipb.DeviceBgpSegmentList
}

func (obj *deviceBgpSegmentList) msg() *snappipb.DeviceBgpSegmentList {
	return obj.obj
}

func (obj *deviceBgpSegmentList) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpSegmentList) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpSegmentList interface {
	msg() *snappipb.DeviceBgpSegmentList
	Yaml() string
	Json() string
	SegmentWeight() int32
	SetSegmentWeight(value int32) DeviceBgpSegmentList
	Segments() []DeviceBgpSegment
	NewSegments() DeviceBgpSegment
	Active() bool
	SetActive(value bool) DeviceBgpSegmentList
}

func (obj *deviceBgpSegmentList) SegmentWeight() int32 {
	return *obj.obj.SegmentWeight
}

func (obj *deviceBgpSegmentList) SetSegmentWeight(value int32) DeviceBgpSegmentList {
	obj.obj.SegmentWeight = &value
	return obj
}

func (obj *deviceBgpSegmentList) Segments() []DeviceBgpSegment {
	if obj.obj.Segments == nil {
		obj.obj.Segments = make([]*snappipb.DeviceBgpSegment, 0)
	}
	values := make([]DeviceBgpSegment, 0)
	for _, item := range obj.obj.Segments {
		values = append(values, &deviceBgpSegment{obj: item})
	}
	return values

}

func (obj *deviceBgpSegmentList) NewSegments() DeviceBgpSegment {
	if obj.obj.Segments == nil {
		obj.obj.Segments = make([]*snappipb.DeviceBgpSegment, 0)
	}
	slice := append(obj.obj.Segments, &snappipb.DeviceBgpSegment{})
	obj.obj.Segments = slice
	return &deviceBgpSegment{obj: slice[len(slice)-1]}
}

func (obj *deviceBgpSegmentList) Active() bool {
	return *obj.obj.Active
}

func (obj *deviceBgpSegmentList) SetActive(value bool) DeviceBgpSegmentList {
	obj.obj.Active = &value
	return obj
}

type deviceBgpRemoteEndpointSubTlv struct {
	obj *snappipb.DeviceBgpRemoteEndpointSubTlv
}

func (obj *deviceBgpRemoteEndpointSubTlv) msg() *snappipb.DeviceBgpRemoteEndpointSubTlv {
	return obj.obj
}

func (obj *deviceBgpRemoteEndpointSubTlv) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpRemoteEndpointSubTlv) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpRemoteEndpointSubTlv interface {
	msg() *snappipb.DeviceBgpRemoteEndpointSubTlv
	Yaml() string
	Json() string
	AsNumber() int32
	SetAsNumber(value int32) DeviceBgpRemoteEndpointSubTlv
	Ipv4Address() string
	SetIpv4Address(value string) DeviceBgpRemoteEndpointSubTlv
	Ipv6Address() string
	SetIpv6Address(value string) DeviceBgpRemoteEndpointSubTlv
}

func (obj *deviceBgpRemoteEndpointSubTlv) AsNumber() int32 {
	return *obj.obj.AsNumber
}

func (obj *deviceBgpRemoteEndpointSubTlv) SetAsNumber(value int32) DeviceBgpRemoteEndpointSubTlv {
	obj.obj.AsNumber = &value
	return obj
}

func (obj *deviceBgpRemoteEndpointSubTlv) Ipv4Address() string {
	return *obj.obj.Ipv4Address
}

func (obj *deviceBgpRemoteEndpointSubTlv) SetIpv4Address(value string) DeviceBgpRemoteEndpointSubTlv {
	obj.obj.Ipv4Address = &value
	return obj
}

func (obj *deviceBgpRemoteEndpointSubTlv) Ipv6Address() string {
	return *obj.obj.Ipv6Address
}

func (obj *deviceBgpRemoteEndpointSubTlv) SetIpv6Address(value string) DeviceBgpRemoteEndpointSubTlv {
	obj.obj.Ipv6Address = &value
	return obj
}

type deviceBgpPreferenceSubTlv struct {
	obj *snappipb.DeviceBgpPreferenceSubTlv
}

func (obj *deviceBgpPreferenceSubTlv) msg() *snappipb.DeviceBgpPreferenceSubTlv {
	return obj.obj
}

func (obj *deviceBgpPreferenceSubTlv) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpPreferenceSubTlv) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpPreferenceSubTlv interface {
	msg() *snappipb.DeviceBgpPreferenceSubTlv
	Yaml() string
	Json() string
	Preference() int32
	SetPreference(value int32) DeviceBgpPreferenceSubTlv
}

func (obj *deviceBgpPreferenceSubTlv) Preference() int32 {
	return *obj.obj.Preference
}

func (obj *deviceBgpPreferenceSubTlv) SetPreference(value int32) DeviceBgpPreferenceSubTlv {
	obj.obj.Preference = &value
	return obj
}

type deviceBgpBindingSubTlv struct {
	obj *snappipb.DeviceBgpBindingSubTlv
}

func (obj *deviceBgpBindingSubTlv) msg() *snappipb.DeviceBgpBindingSubTlv {
	return obj.obj
}

func (obj *deviceBgpBindingSubTlv) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpBindingSubTlv) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpBindingSubTlv interface {
	msg() *snappipb.DeviceBgpBindingSubTlv
	Yaml() string
	Json() string
	FourOctetSid() int32
	SetFourOctetSid(value int32) DeviceBgpBindingSubTlv
	BsidAsMplsLabel() bool
	SetBsidAsMplsLabel(value bool) DeviceBgpBindingSubTlv
	Ipv6Sid() string
	SetIpv6Sid(value string) DeviceBgpBindingSubTlv
	SFlag() bool
	SetSFlag(value bool) DeviceBgpBindingSubTlv
	IFlag() bool
	SetIFlag(value bool) DeviceBgpBindingSubTlv
	RemainingFlagBits() int32
	SetRemainingFlagBits(value int32) DeviceBgpBindingSubTlv
}

func (obj *deviceBgpBindingSubTlv) FourOctetSid() int32 {
	return *obj.obj.FourOctetSid
}

func (obj *deviceBgpBindingSubTlv) SetFourOctetSid(value int32) DeviceBgpBindingSubTlv {
	obj.obj.FourOctetSid = &value
	return obj
}

func (obj *deviceBgpBindingSubTlv) BsidAsMplsLabel() bool {
	return *obj.obj.BsidAsMplsLabel
}

func (obj *deviceBgpBindingSubTlv) SetBsidAsMplsLabel(value bool) DeviceBgpBindingSubTlv {
	obj.obj.BsidAsMplsLabel = &value
	return obj
}

func (obj *deviceBgpBindingSubTlv) Ipv6Sid() string {
	return *obj.obj.Ipv6Sid
}

func (obj *deviceBgpBindingSubTlv) SetIpv6Sid(value string) DeviceBgpBindingSubTlv {
	obj.obj.Ipv6Sid = &value
	return obj
}

func (obj *deviceBgpBindingSubTlv) SFlag() bool {
	return *obj.obj.SFlag
}

func (obj *deviceBgpBindingSubTlv) SetSFlag(value bool) DeviceBgpBindingSubTlv {
	obj.obj.SFlag = &value
	return obj
}

func (obj *deviceBgpBindingSubTlv) IFlag() bool {
	return *obj.obj.IFlag
}

func (obj *deviceBgpBindingSubTlv) SetIFlag(value bool) DeviceBgpBindingSubTlv {
	obj.obj.IFlag = &value
	return obj
}

func (obj *deviceBgpBindingSubTlv) RemainingFlagBits() int32 {
	return *obj.obj.RemainingFlagBits
}

func (obj *deviceBgpBindingSubTlv) SetRemainingFlagBits(value int32) DeviceBgpBindingSubTlv {
	obj.obj.RemainingFlagBits = &value
	return obj
}

type deviceBgpExplicitNullLabelPolicySubTlv struct {
	obj *snappipb.DeviceBgpExplicitNullLabelPolicySubTlv
}

func (obj *deviceBgpExplicitNullLabelPolicySubTlv) msg() *snappipb.DeviceBgpExplicitNullLabelPolicySubTlv {
	return obj.obj
}

func (obj *deviceBgpExplicitNullLabelPolicySubTlv) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpExplicitNullLabelPolicySubTlv) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpExplicitNullLabelPolicySubTlv interface {
	msg() *snappipb.DeviceBgpExplicitNullLabelPolicySubTlv
	Yaml() string
	Json() string
}

type patternFlowIpv4TosPrecedenceCounter struct {
	obj *snappipb.PatternFlowIpv4TosPrecedenceCounter
}

func (obj *patternFlowIpv4TosPrecedenceCounter) msg() *snappipb.PatternFlowIpv4TosPrecedenceCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TosPrecedenceCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TosPrecedenceCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TosPrecedenceCounter interface {
	msg() *snappipb.PatternFlowIpv4TosPrecedenceCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4TosPrecedenceCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4TosPrecedenceCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4TosPrecedenceCounter
}

func (obj *patternFlowIpv4TosPrecedenceCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4TosPrecedenceCounter) SetStart(value int32) PatternFlowIpv4TosPrecedenceCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4TosPrecedenceCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4TosPrecedenceCounter) SetStep(value int32) PatternFlowIpv4TosPrecedenceCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4TosPrecedenceCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4TosPrecedenceCounter) SetCount(value int32) PatternFlowIpv4TosPrecedenceCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4TosDelayCounter struct {
	obj *snappipb.PatternFlowIpv4TosDelayCounter
}

func (obj *patternFlowIpv4TosDelayCounter) msg() *snappipb.PatternFlowIpv4TosDelayCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TosDelayCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TosDelayCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TosDelayCounter interface {
	msg() *snappipb.PatternFlowIpv4TosDelayCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4TosDelayCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4TosDelayCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4TosDelayCounter
}

func (obj *patternFlowIpv4TosDelayCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4TosDelayCounter) SetStart(value int32) PatternFlowIpv4TosDelayCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4TosDelayCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4TosDelayCounter) SetStep(value int32) PatternFlowIpv4TosDelayCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4TosDelayCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4TosDelayCounter) SetCount(value int32) PatternFlowIpv4TosDelayCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4TosThroughputCounter struct {
	obj *snappipb.PatternFlowIpv4TosThroughputCounter
}

func (obj *patternFlowIpv4TosThroughputCounter) msg() *snappipb.PatternFlowIpv4TosThroughputCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TosThroughputCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TosThroughputCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TosThroughputCounter interface {
	msg() *snappipb.PatternFlowIpv4TosThroughputCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4TosThroughputCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4TosThroughputCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4TosThroughputCounter
}

func (obj *patternFlowIpv4TosThroughputCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4TosThroughputCounter) SetStart(value int32) PatternFlowIpv4TosThroughputCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4TosThroughputCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4TosThroughputCounter) SetStep(value int32) PatternFlowIpv4TosThroughputCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4TosThroughputCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4TosThroughputCounter) SetCount(value int32) PatternFlowIpv4TosThroughputCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4TosReliabilityCounter struct {
	obj *snappipb.PatternFlowIpv4TosReliabilityCounter
}

func (obj *patternFlowIpv4TosReliabilityCounter) msg() *snappipb.PatternFlowIpv4TosReliabilityCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TosReliabilityCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TosReliabilityCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TosReliabilityCounter interface {
	msg() *snappipb.PatternFlowIpv4TosReliabilityCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4TosReliabilityCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4TosReliabilityCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4TosReliabilityCounter
}

func (obj *patternFlowIpv4TosReliabilityCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4TosReliabilityCounter) SetStart(value int32) PatternFlowIpv4TosReliabilityCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4TosReliabilityCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4TosReliabilityCounter) SetStep(value int32) PatternFlowIpv4TosReliabilityCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4TosReliabilityCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4TosReliabilityCounter) SetCount(value int32) PatternFlowIpv4TosReliabilityCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4TosMonetaryCounter struct {
	obj *snappipb.PatternFlowIpv4TosMonetaryCounter
}

func (obj *patternFlowIpv4TosMonetaryCounter) msg() *snappipb.PatternFlowIpv4TosMonetaryCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TosMonetaryCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TosMonetaryCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TosMonetaryCounter interface {
	msg() *snappipb.PatternFlowIpv4TosMonetaryCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4TosMonetaryCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4TosMonetaryCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4TosMonetaryCounter
}

func (obj *patternFlowIpv4TosMonetaryCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4TosMonetaryCounter) SetStart(value int32) PatternFlowIpv4TosMonetaryCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4TosMonetaryCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4TosMonetaryCounter) SetStep(value int32) PatternFlowIpv4TosMonetaryCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4TosMonetaryCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4TosMonetaryCounter) SetCount(value int32) PatternFlowIpv4TosMonetaryCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4TosUnusedCounter struct {
	obj *snappipb.PatternFlowIpv4TosUnusedCounter
}

func (obj *patternFlowIpv4TosUnusedCounter) msg() *snappipb.PatternFlowIpv4TosUnusedCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TosUnusedCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4TosUnusedCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4TosUnusedCounter interface {
	msg() *snappipb.PatternFlowIpv4TosUnusedCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4TosUnusedCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4TosUnusedCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4TosUnusedCounter
}

func (obj *patternFlowIpv4TosUnusedCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4TosUnusedCounter) SetStart(value int32) PatternFlowIpv4TosUnusedCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4TosUnusedCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4TosUnusedCounter) SetStep(value int32) PatternFlowIpv4TosUnusedCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4TosUnusedCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4TosUnusedCounter) SetCount(value int32) PatternFlowIpv4TosUnusedCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4DscpPhbCounter struct {
	obj *snappipb.PatternFlowIpv4DscpPhbCounter
}

func (obj *patternFlowIpv4DscpPhbCounter) msg() *snappipb.PatternFlowIpv4DscpPhbCounter {
	return obj.obj
}

func (obj *patternFlowIpv4DscpPhbCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4DscpPhbCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4DscpPhbCounter interface {
	msg() *snappipb.PatternFlowIpv4DscpPhbCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4DscpPhbCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4DscpPhbCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4DscpPhbCounter
}

func (obj *patternFlowIpv4DscpPhbCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4DscpPhbCounter) SetStart(value int32) PatternFlowIpv4DscpPhbCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4DscpPhbCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4DscpPhbCounter) SetStep(value int32) PatternFlowIpv4DscpPhbCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4DscpPhbCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4DscpPhbCounter) SetCount(value int32) PatternFlowIpv4DscpPhbCounter {
	obj.obj.Count = &value
	return obj
}

type patternFlowIpv4DscpEcnCounter struct {
	obj *snappipb.PatternFlowIpv4DscpEcnCounter
}

func (obj *patternFlowIpv4DscpEcnCounter) msg() *snappipb.PatternFlowIpv4DscpEcnCounter {
	return obj.obj
}

func (obj *patternFlowIpv4DscpEcnCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternFlowIpv4DscpEcnCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternFlowIpv4DscpEcnCounter interface {
	msg() *snappipb.PatternFlowIpv4DscpEcnCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternFlowIpv4DscpEcnCounter
	Step() int32
	SetStep(value int32) PatternFlowIpv4DscpEcnCounter
	Count() int32
	SetCount(value int32) PatternFlowIpv4DscpEcnCounter
}

func (obj *patternFlowIpv4DscpEcnCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternFlowIpv4DscpEcnCounter) SetStart(value int32) PatternFlowIpv4DscpEcnCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternFlowIpv4DscpEcnCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternFlowIpv4DscpEcnCounter) SetStep(value int32) PatternFlowIpv4DscpEcnCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternFlowIpv4DscpEcnCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternFlowIpv4DscpEcnCounter) SetCount(value int32) PatternFlowIpv4DscpEcnCounter {
	obj.obj.Count = &value
	return obj
}

type deviceBgpSegment struct {
	obj *snappipb.DeviceBgpSegment
}

func (obj *deviceBgpSegment) msg() *snappipb.DeviceBgpSegment {
	return obj.obj
}

func (obj *deviceBgpSegment) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *deviceBgpSegment) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type DeviceBgpSegment interface {
	msg() *snappipb.DeviceBgpSegment
	Yaml() string
	Json() string
	MplsLabel() int32
	SetMplsLabel(value int32) DeviceBgpSegment
	MplsTc() int32
	SetMplsTc(value int32) DeviceBgpSegment
	MplsTtl() int32
	SetMplsTtl(value int32) DeviceBgpSegment
	VFlag() bool
	SetVFlag(value bool) DeviceBgpSegment
	Ipv6Sid() string
	SetIpv6Sid(value string) DeviceBgpSegment
	RemainingFlagBits() int32
	SetRemainingFlagBits(value int32) DeviceBgpSegment
	Active() bool
	SetActive(value bool) DeviceBgpSegment
}

func (obj *deviceBgpSegment) MplsLabel() int32 {
	return *obj.obj.MplsLabel
}

func (obj *deviceBgpSegment) SetMplsLabel(value int32) DeviceBgpSegment {
	obj.obj.MplsLabel = &value
	return obj
}

func (obj *deviceBgpSegment) MplsTc() int32 {
	return *obj.obj.MplsTc
}

func (obj *deviceBgpSegment) SetMplsTc(value int32) DeviceBgpSegment {
	obj.obj.MplsTc = &value
	return obj
}

func (obj *deviceBgpSegment) MplsTtl() int32 {
	return *obj.obj.MplsTtl
}

func (obj *deviceBgpSegment) SetMplsTtl(value int32) DeviceBgpSegment {
	obj.obj.MplsTtl = &value
	return obj
}

func (obj *deviceBgpSegment) VFlag() bool {
	return *obj.obj.VFlag
}

func (obj *deviceBgpSegment) SetVFlag(value bool) DeviceBgpSegment {
	obj.obj.VFlag = &value
	return obj
}

func (obj *deviceBgpSegment) Ipv6Sid() string {
	return *obj.obj.Ipv6Sid
}

func (obj *deviceBgpSegment) SetIpv6Sid(value string) DeviceBgpSegment {
	obj.obj.Ipv6Sid = &value
	return obj
}

func (obj *deviceBgpSegment) RemainingFlagBits() int32 {
	return *obj.obj.RemainingFlagBits
}

func (obj *deviceBgpSegment) SetRemainingFlagBits(value int32) DeviceBgpSegment {
	obj.obj.RemainingFlagBits = &value
	return obj
}

func (obj *deviceBgpSegment) Active() bool {
	return *obj.obj.Active
}

func (obj *deviceBgpSegment) SetActive(value bool) DeviceBgpSegment {
	obj.obj.Active = &value
	return obj
}
