package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceDhcpv6Client *****
type deviceDhcpv6Client struct {
	validation
	obj                  *otg.DeviceDhcpv6Client
	marshaller           marshalDeviceDhcpv6Client
	unMarshaller         unMarshalDeviceDhcpv6Client
	iaTypeHolder         DeviceDhcpv6ClientIaType
	duidTypeHolder       DeviceDhcpv6ClientDuidType
	optionsRequestHolder DeviceDhcpv6ClientOptionsRequest
	optionsHolder        DeviceDhcpv6ClientOptions
}

func NewDeviceDhcpv6Client() DeviceDhcpv6Client {
	obj := deviceDhcpv6Client{obj: &otg.DeviceDhcpv6Client{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceDhcpv6Client) msg() *otg.DeviceDhcpv6Client {
	return obj.obj
}

func (obj *deviceDhcpv6Client) setMsg(msg *otg.DeviceDhcpv6Client) DeviceDhcpv6Client {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceDhcpv6Client struct {
	obj *deviceDhcpv6Client
}

type marshalDeviceDhcpv6Client interface {
	// ToProto marshals DeviceDhcpv6Client to protobuf object *otg.DeviceDhcpv6Client
	ToProto() (*otg.DeviceDhcpv6Client, error)
	// ToPbText marshals DeviceDhcpv6Client to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceDhcpv6Client to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceDhcpv6Client to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceDhcpv6Client struct {
	obj *deviceDhcpv6Client
}

type unMarshalDeviceDhcpv6Client interface {
	// FromProto unmarshals DeviceDhcpv6Client from protobuf object *otg.DeviceDhcpv6Client
	FromProto(msg *otg.DeviceDhcpv6Client) (DeviceDhcpv6Client, error)
	// FromPbText unmarshals DeviceDhcpv6Client from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceDhcpv6Client from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceDhcpv6Client from JSON text
	FromJson(value string) error
}

func (obj *deviceDhcpv6Client) Marshal() marshalDeviceDhcpv6Client {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceDhcpv6Client{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceDhcpv6Client) Unmarshal() unMarshalDeviceDhcpv6Client {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceDhcpv6Client{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceDhcpv6Client) ToProto() (*otg.DeviceDhcpv6Client, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceDhcpv6Client) FromProto(msg *otg.DeviceDhcpv6Client) (DeviceDhcpv6Client, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceDhcpv6Client) ToPbText() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (m *unMarshaldeviceDhcpv6Client) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshaldeviceDhcpv6Client) ToYaml() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshaldeviceDhcpv6Client) FromYaml(value string) error {
	if value == "" {
		value = "{}"
	}
	data, err := yaml.YAMLToJSON([]byte(value))
	if err != nil {
		return err
	}
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	uError := opts.Unmarshal([]byte(data), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshaldeviceDhcpv6Client) ToJson() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshaldeviceDhcpv6Client) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *deviceDhcpv6Client) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceDhcpv6Client) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceDhcpv6Client) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceDhcpv6Client) Clone() (DeviceDhcpv6Client, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceDhcpv6Client()
	data, err := proto.Marshal(obj.msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *deviceDhcpv6Client) setNil() {
	obj.iaTypeHolder = nil
	obj.duidTypeHolder = nil
	obj.optionsRequestHolder = nil
	obj.optionsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceDhcpv6Client is configuration for emulated DHCPv6 Client on a single Interface. If the DHCPv6 Client receives one or more  DHCPv6 ADVERTISE messages from one or more servers then the client chooses one server from which to request  configuration parameters, based on the configuration parameters offered by the server in the DHCPv6 ADVERTISE  messages. If all configuration parameters match then the first server will be chosen. https://www.rfc-editor.org/rfc/rfc8415.html
type DeviceDhcpv6Client interface {
	Validation
	// msg marshals DeviceDhcpv6Client to protobuf object *otg.DeviceDhcpv6Client
	// and doesn't set defaults
	msg() *otg.DeviceDhcpv6Client
	// setMsg unmarshals DeviceDhcpv6Client from protobuf object *otg.DeviceDhcpv6Client
	// and doesn't set defaults
	setMsg(*otg.DeviceDhcpv6Client) DeviceDhcpv6Client
	// provides marshal interface
	Marshal() marshalDeviceDhcpv6Client
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceDhcpv6Client
	// validate validates DeviceDhcpv6Client
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceDhcpv6Client, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in DeviceDhcpv6Client.
	Name() string
	// SetName assigns string provided by user to DeviceDhcpv6Client
	SetName(value string) DeviceDhcpv6Client
	// RapidCommit returns bool, set in DeviceDhcpv6Client.
	RapidCommit() bool
	// SetRapidCommit assigns bool provided by user to DeviceDhcpv6Client
	SetRapidCommit(value bool) DeviceDhcpv6Client
	// HasRapidCommit checks if RapidCommit has been set in DeviceDhcpv6Client
	HasRapidCommit() bool
	// IaType returns DeviceDhcpv6ClientIaType, set in DeviceDhcpv6Client.
	IaType() DeviceDhcpv6ClientIaType
	// SetIaType assigns DeviceDhcpv6ClientIaType provided by user to DeviceDhcpv6Client.
	SetIaType(value DeviceDhcpv6ClientIaType) DeviceDhcpv6Client
	// HasIaType checks if IaType has been set in DeviceDhcpv6Client
	HasIaType() bool
	// DuidType returns DeviceDhcpv6ClientDuidType, set in DeviceDhcpv6Client.
	DuidType() DeviceDhcpv6ClientDuidType
	// SetDuidType assigns DeviceDhcpv6ClientDuidType provided by user to DeviceDhcpv6Client.
	SetDuidType(value DeviceDhcpv6ClientDuidType) DeviceDhcpv6Client
	// HasDuidType checks if DuidType has been set in DeviceDhcpv6Client
	HasDuidType() bool
	// OptionsRequest returns DeviceDhcpv6ClientOptionsRequest, set in DeviceDhcpv6Client.
	// DeviceDhcpv6ClientOptionsRequest is dHCP client options, these configured options are sent in Dhcp client messages.
	OptionsRequest() DeviceDhcpv6ClientOptionsRequest
	// SetOptionsRequest assigns DeviceDhcpv6ClientOptionsRequest provided by user to DeviceDhcpv6Client.
	// DeviceDhcpv6ClientOptionsRequest is dHCP client options, these configured options are sent in Dhcp client messages.
	SetOptionsRequest(value DeviceDhcpv6ClientOptionsRequest) DeviceDhcpv6Client
	// HasOptionsRequest checks if OptionsRequest has been set in DeviceDhcpv6Client
	HasOptionsRequest() bool
	// Options returns DeviceDhcpv6ClientOptions, set in DeviceDhcpv6Client.
	// DeviceDhcpv6ClientOptions is dHCP client options, these configured options are sent in Dhcp client messages.
	Options() DeviceDhcpv6ClientOptions
	// SetOptions assigns DeviceDhcpv6ClientOptions provided by user to DeviceDhcpv6Client.
	// DeviceDhcpv6ClientOptions is dHCP client options, these configured options are sent in Dhcp client messages.
	SetOptions(value DeviceDhcpv6ClientOptions) DeviceDhcpv6Client
	// HasOptions checks if Options has been set in DeviceDhcpv6Client
	HasOptions() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceDhcpv6Client) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetName(value string) DeviceDhcpv6Client {

	obj.obj.Name = &value
	return obj
}

// If Rapid Commit is set, client initiates Rapid Commit two-message exchange by including Rapid Commit option  in Solicit message.
// RapidCommit returns a bool
func (obj *deviceDhcpv6Client) RapidCommit() bool {

	return *obj.obj.RapidCommit

}

// If Rapid Commit is set, client initiates Rapid Commit two-message exchange by including Rapid Commit option  in Solicit message.
// RapidCommit returns a bool
func (obj *deviceDhcpv6Client) HasRapidCommit() bool {
	return obj.obj.RapidCommit != nil
}

// If Rapid Commit is set, client initiates Rapid Commit two-message exchange by including Rapid Commit option  in Solicit message.
// SetRapidCommit sets the bool value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetRapidCommit(value bool) DeviceDhcpv6Client {

	obj.obj.RapidCommit = &value
	return obj
}

// Each IA has an associated IAID. Differnet IA options represent different types of IPv6 addresses and parameters
// accepted by DHCPv6 clients each used in different context by an IPv6 node.
// IaType returns a DeviceDhcpv6ClientIaType
func (obj *deviceDhcpv6Client) IaType() DeviceDhcpv6ClientIaType {
	if obj.obj.IaType == nil {
		obj.obj.IaType = NewDeviceDhcpv6ClientIaType().msg()
	}
	if obj.iaTypeHolder == nil {
		obj.iaTypeHolder = &deviceDhcpv6ClientIaType{obj: obj.obj.IaType}
	}
	return obj.iaTypeHolder
}

// Each IA has an associated IAID. Differnet IA options represent different types of IPv6 addresses and parameters
// accepted by DHCPv6 clients each used in different context by an IPv6 node.
// IaType returns a DeviceDhcpv6ClientIaType
func (obj *deviceDhcpv6Client) HasIaType() bool {
	return obj.obj.IaType != nil
}

// Each IA has an associated IAID. Differnet IA options represent different types of IPv6 addresses and parameters
// accepted by DHCPv6 clients each used in different context by an IPv6 node.
// SetIaType sets the DeviceDhcpv6ClientIaType value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetIaType(value DeviceDhcpv6ClientIaType) DeviceDhcpv6Client {

	obj.iaTypeHolder = nil
	obj.obj.IaType = value.msg()

	return obj
}

// Each DHCP client and server has a DUID. DHCP clients and servers use DUIDs to identify each other.
// DuidType returns a DeviceDhcpv6ClientDuidType
func (obj *deviceDhcpv6Client) DuidType() DeviceDhcpv6ClientDuidType {
	if obj.obj.DuidType == nil {
		obj.obj.DuidType = NewDeviceDhcpv6ClientDuidType().msg()
	}
	if obj.duidTypeHolder == nil {
		obj.duidTypeHolder = &deviceDhcpv6ClientDuidType{obj: obj.obj.DuidType}
	}
	return obj.duidTypeHolder
}

// Each DHCP client and server has a DUID. DHCP clients and servers use DUIDs to identify each other.
// DuidType returns a DeviceDhcpv6ClientDuidType
func (obj *deviceDhcpv6Client) HasDuidType() bool {
	return obj.obj.DuidType != nil
}

// Each DHCP client and server has a DUID. DHCP clients and servers use DUIDs to identify each other.
// SetDuidType sets the DeviceDhcpv6ClientDuidType value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetDuidType(value DeviceDhcpv6ClientDuidType) DeviceDhcpv6Client {

	obj.duidTypeHolder = nil
	obj.obj.DuidType = value.msg()

	return obj
}

// The options requested by a client from a server in the options request option.
// OptionsRequest returns a DeviceDhcpv6ClientOptionsRequest
func (obj *deviceDhcpv6Client) OptionsRequest() DeviceDhcpv6ClientOptionsRequest {
	if obj.obj.OptionsRequest == nil {
		obj.obj.OptionsRequest = NewDeviceDhcpv6ClientOptionsRequest().msg()
	}
	if obj.optionsRequestHolder == nil {
		obj.optionsRequestHolder = &deviceDhcpv6ClientOptionsRequest{obj: obj.obj.OptionsRequest}
	}
	return obj.optionsRequestHolder
}

// The options requested by a client from a server in the options request option.
// OptionsRequest returns a DeviceDhcpv6ClientOptionsRequest
func (obj *deviceDhcpv6Client) HasOptionsRequest() bool {
	return obj.obj.OptionsRequest != nil
}

// The options requested by a client from a server in the options request option.
// SetOptionsRequest sets the DeviceDhcpv6ClientOptionsRequest value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetOptionsRequest(value DeviceDhcpv6ClientOptionsRequest) DeviceDhcpv6Client {

	obj.optionsRequestHolder = nil
	obj.obj.OptionsRequest = value.msg()

	return obj
}

// Optional DHCPv4 Client options that are sent in Dhcp client messages.
// Options returns a DeviceDhcpv6ClientOptions
func (obj *deviceDhcpv6Client) Options() DeviceDhcpv6ClientOptions {
	if obj.obj.Options == nil {
		obj.obj.Options = NewDeviceDhcpv6ClientOptions().msg()
	}
	if obj.optionsHolder == nil {
		obj.optionsHolder = &deviceDhcpv6ClientOptions{obj: obj.obj.Options}
	}
	return obj.optionsHolder
}

// Optional DHCPv4 Client options that are sent in Dhcp client messages.
// Options returns a DeviceDhcpv6ClientOptions
func (obj *deviceDhcpv6Client) HasOptions() bool {
	return obj.obj.Options != nil
}

// Optional DHCPv4 Client options that are sent in Dhcp client messages.
// SetOptions sets the DeviceDhcpv6ClientOptions value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetOptions(value DeviceDhcpv6ClientOptions) DeviceDhcpv6Client {

	obj.optionsHolder = nil
	obj.obj.Options = value.msg()

	return obj
}

func (obj *deviceDhcpv6Client) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceDhcpv6Client")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"DeviceDhcpv6Client.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.IaType != nil {

		obj.IaType().validateObj(vObj, set_default)
	}

	if obj.obj.DuidType != nil {

		obj.DuidType().validateObj(vObj, set_default)
	}

	if obj.obj.OptionsRequest != nil {

		obj.OptionsRequest().validateObj(vObj, set_default)
	}

	if obj.obj.Options != nil {

		obj.Options().validateObj(vObj, set_default)
	}

}

func (obj *deviceDhcpv6Client) setDefault() {
	if obj.obj.RapidCommit == nil {
		obj.SetRapidCommit(false)
	}

}
