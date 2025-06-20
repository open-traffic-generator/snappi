package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceIpv4 *****
type deviceIpv4 struct {
	validation
	obj              *otg.DeviceIpv4
	marshaller       marshalDeviceIpv4
	unMarshaller     unMarshalDeviceIpv4
	gatewayMacHolder DeviceIpv4GatewayMAC
}

func NewDeviceIpv4() DeviceIpv4 {
	obj := deviceIpv4{obj: &otg.DeviceIpv4{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceIpv4) msg() *otg.DeviceIpv4 {
	return obj.obj
}

func (obj *deviceIpv4) setMsg(msg *otg.DeviceIpv4) DeviceIpv4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceIpv4 struct {
	obj *deviceIpv4
}

type marshalDeviceIpv4 interface {
	// ToProto marshals DeviceIpv4 to protobuf object *otg.DeviceIpv4
	ToProto() (*otg.DeviceIpv4, error)
	// ToPbText marshals DeviceIpv4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceIpv4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceIpv4 to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceIpv4 struct {
	obj *deviceIpv4
}

type unMarshalDeviceIpv4 interface {
	// FromProto unmarshals DeviceIpv4 from protobuf object *otg.DeviceIpv4
	FromProto(msg *otg.DeviceIpv4) (DeviceIpv4, error)
	// FromPbText unmarshals DeviceIpv4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceIpv4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceIpv4 from JSON text
	FromJson(value string) error
}

func (obj *deviceIpv4) Marshal() marshalDeviceIpv4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceIpv4{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceIpv4) Unmarshal() unMarshalDeviceIpv4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceIpv4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceIpv4) ToProto() (*otg.DeviceIpv4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceIpv4) FromProto(msg *otg.DeviceIpv4) (DeviceIpv4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceIpv4) ToPbText() (string, error) {
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

func (m *unMarshaldeviceIpv4) FromPbText(value string) error {
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

func (m *marshaldeviceIpv4) ToYaml() (string, error) {
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

func (m *unMarshaldeviceIpv4) FromYaml(value string) error {
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

func (m *marshaldeviceIpv4) ToJson() (string, error) {
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

func (m *unMarshaldeviceIpv4) FromJson(value string) error {
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

func (obj *deviceIpv4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceIpv4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceIpv4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceIpv4) Clone() (DeviceIpv4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceIpv4()
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

func (obj *deviceIpv4) setNil() {
	obj.gatewayMacHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceIpv4 is an IPv4 interface with gateway
type DeviceIpv4 interface {
	Validation
	// msg marshals DeviceIpv4 to protobuf object *otg.DeviceIpv4
	// and doesn't set defaults
	msg() *otg.DeviceIpv4
	// setMsg unmarshals DeviceIpv4 from protobuf object *otg.DeviceIpv4
	// and doesn't set defaults
	setMsg(*otg.DeviceIpv4) DeviceIpv4
	// provides marshal interface
	Marshal() marshalDeviceIpv4
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceIpv4
	// validate validates DeviceIpv4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceIpv4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Gateway returns string, set in DeviceIpv4.
	Gateway() string
	// SetGateway assigns string provided by user to DeviceIpv4
	SetGateway(value string) DeviceIpv4
	// GatewayMac returns DeviceIpv4GatewayMAC, set in DeviceIpv4.
	// DeviceIpv4GatewayMAC is by default auto(resolved gateway mac) is set.  Setting a value would mean that ARP will not be used for learning MAC of connected device. The user-configured MAC address will be used for auto-filling the destination
	// MAC address in the control and data packets sent from this IPv4 endpoint
	// whenever applicable.
	GatewayMac() DeviceIpv4GatewayMAC
	// SetGatewayMac assigns DeviceIpv4GatewayMAC provided by user to DeviceIpv4.
	// DeviceIpv4GatewayMAC is by default auto(resolved gateway mac) is set.  Setting a value would mean that ARP will not be used for learning MAC of connected device. The user-configured MAC address will be used for auto-filling the destination
	// MAC address in the control and data packets sent from this IPv4 endpoint
	// whenever applicable.
	SetGatewayMac(value DeviceIpv4GatewayMAC) DeviceIpv4
	// HasGatewayMac checks if GatewayMac has been set in DeviceIpv4
	HasGatewayMac() bool
	// Address returns string, set in DeviceIpv4.
	Address() string
	// SetAddress assigns string provided by user to DeviceIpv4
	SetAddress(value string) DeviceIpv4
	// Prefix returns uint32, set in DeviceIpv4.
	Prefix() uint32
	// SetPrefix assigns uint32 provided by user to DeviceIpv4
	SetPrefix(value uint32) DeviceIpv4
	// HasPrefix checks if Prefix has been set in DeviceIpv4
	HasPrefix() bool
	// Name returns string, set in DeviceIpv4.
	Name() string
	// SetName assigns string provided by user to DeviceIpv4
	SetName(value string) DeviceIpv4
	setNil()
}

// The IPv4 address of the gateway
// Gateway returns a string
func (obj *deviceIpv4) Gateway() string {

	return *obj.obj.Gateway

}

// The IPv4 address of the gateway
// SetGateway sets the string value in the DeviceIpv4 object
func (obj *deviceIpv4) SetGateway(value string) DeviceIpv4 {

	obj.obj.Gateway = &value
	return obj
}

// description is TBD
// GatewayMac returns a DeviceIpv4GatewayMAC
func (obj *deviceIpv4) GatewayMac() DeviceIpv4GatewayMAC {
	if obj.obj.GatewayMac == nil {
		obj.obj.GatewayMac = NewDeviceIpv4GatewayMAC().msg()
	}
	if obj.gatewayMacHolder == nil {
		obj.gatewayMacHolder = &deviceIpv4GatewayMAC{obj: obj.obj.GatewayMac}
	}
	return obj.gatewayMacHolder
}

// description is TBD
// GatewayMac returns a DeviceIpv4GatewayMAC
func (obj *deviceIpv4) HasGatewayMac() bool {
	return obj.obj.GatewayMac != nil
}

// description is TBD
// SetGatewayMac sets the DeviceIpv4GatewayMAC value in the DeviceIpv4 object
func (obj *deviceIpv4) SetGatewayMac(value DeviceIpv4GatewayMAC) DeviceIpv4 {

	obj.gatewayMacHolder = nil
	obj.obj.GatewayMac = value.msg()

	return obj
}

// The IPv4 address
// Address returns a string
func (obj *deviceIpv4) Address() string {

	return *obj.obj.Address

}

// The IPv4 address
// SetAddress sets the string value in the DeviceIpv4 object
func (obj *deviceIpv4) SetAddress(value string) DeviceIpv4 {

	obj.obj.Address = &value
	return obj
}

// The prefix of the IPv4 address.
// Prefix returns a uint32
func (obj *deviceIpv4) Prefix() uint32 {

	return *obj.obj.Prefix

}

// The prefix of the IPv4 address.
// Prefix returns a uint32
func (obj *deviceIpv4) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The prefix of the IPv4 address.
// SetPrefix sets the uint32 value in the DeviceIpv4 object
func (obj *deviceIpv4) SetPrefix(value uint32) DeviceIpv4 {

	obj.obj.Prefix = &value
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceIpv4) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceIpv4 object
func (obj *deviceIpv4) SetName(value string) DeviceIpv4 {

	obj.obj.Name = &value
	return obj
}

func (obj *deviceIpv4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Gateway is required
	if obj.obj.Gateway == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Gateway is required field on interface DeviceIpv4")
	}
	if obj.obj.Gateway != nil {

		err := obj.validateIpv4(obj.Gateway())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv4.Gateway"))
		}

	}

	if obj.obj.GatewayMac != nil {

		obj.GatewayMac().validateObj(vObj, set_default)
	}

	// Address is required
	if obj.obj.Address == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Address is required field on interface DeviceIpv4")
	}
	if obj.obj.Address != nil {

		err := obj.validateIpv4(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv4.Address"))
		}

	}

	if obj.obj.Prefix != nil {

		if *obj.obj.Prefix < 1 || *obj.obj.Prefix > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceIpv4.Prefix <= 32 but Got %d", *obj.obj.Prefix))
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceIpv4")
	}
}

func (obj *deviceIpv4) setDefault() {
	if obj.obj.Prefix == nil {
		obj.SetPrefix(24)
	}

}
