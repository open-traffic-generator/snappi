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

// ***** DeviceIpv6 *****
type deviceIpv6 struct {
	validation
	obj              *otg.DeviceIpv6
	marshaller       marshalDeviceIpv6
	unMarshaller     unMarshalDeviceIpv6
	gatewayMacHolder DeviceIpv6GatewayMAC
}

func NewDeviceIpv6() DeviceIpv6 {
	obj := deviceIpv6{obj: &otg.DeviceIpv6{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceIpv6) msg() *otg.DeviceIpv6 {
	return obj.obj
}

func (obj *deviceIpv6) setMsg(msg *otg.DeviceIpv6) DeviceIpv6 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceIpv6 struct {
	obj *deviceIpv6
}

type marshalDeviceIpv6 interface {
	// ToProto marshals DeviceIpv6 to protobuf object *otg.DeviceIpv6
	ToProto() (*otg.DeviceIpv6, error)
	// ToPbText marshals DeviceIpv6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceIpv6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceIpv6 to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceIpv6 struct {
	obj *deviceIpv6
}

type unMarshalDeviceIpv6 interface {
	// FromProto unmarshals DeviceIpv6 from protobuf object *otg.DeviceIpv6
	FromProto(msg *otg.DeviceIpv6) (DeviceIpv6, error)
	// FromPbText unmarshals DeviceIpv6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceIpv6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceIpv6 from JSON text
	FromJson(value string) error
}

func (obj *deviceIpv6) Marshal() marshalDeviceIpv6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceIpv6{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceIpv6) Unmarshal() unMarshalDeviceIpv6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceIpv6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceIpv6) ToProto() (*otg.DeviceIpv6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceIpv6) FromProto(msg *otg.DeviceIpv6) (DeviceIpv6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceIpv6) ToPbText() (string, error) {
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

func (m *unMarshaldeviceIpv6) FromPbText(value string) error {
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

func (m *marshaldeviceIpv6) ToYaml() (string, error) {
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

func (m *unMarshaldeviceIpv6) FromYaml(value string) error {
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

func (m *marshaldeviceIpv6) ToJson() (string, error) {
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

func (m *unMarshaldeviceIpv6) FromJson(value string) error {
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

func (obj *deviceIpv6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceIpv6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceIpv6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceIpv6) Clone() (DeviceIpv6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceIpv6()
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

func (obj *deviceIpv6) setNil() {
	obj.gatewayMacHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceIpv6 is an IPv6 interface with gateway.
type DeviceIpv6 interface {
	Validation
	// msg marshals DeviceIpv6 to protobuf object *otg.DeviceIpv6
	// and doesn't set defaults
	msg() *otg.DeviceIpv6
	// setMsg unmarshals DeviceIpv6 from protobuf object *otg.DeviceIpv6
	// and doesn't set defaults
	setMsg(*otg.DeviceIpv6) DeviceIpv6
	// provides marshal interface
	Marshal() marshalDeviceIpv6
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceIpv6
	// validate validates DeviceIpv6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceIpv6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Gateway returns string, set in DeviceIpv6.
	Gateway() string
	// SetGateway assigns string provided by user to DeviceIpv6
	SetGateway(value string) DeviceIpv6
	// GatewayMac returns DeviceIpv6GatewayMAC, set in DeviceIpv6.
	// DeviceIpv6GatewayMAC is by default auto(resolved gateway mac) is set. Setting a value would mean that ND will not be used for learning MAC of connected device. The user-configured MAC address will be used for auto-filling the destination
	// MAC address in the control and data packets sent from this IPv6 endpoint
	// whenever applicable.
	GatewayMac() DeviceIpv6GatewayMAC
	// SetGatewayMac assigns DeviceIpv6GatewayMAC provided by user to DeviceIpv6.
	// DeviceIpv6GatewayMAC is by default auto(resolved gateway mac) is set. Setting a value would mean that ND will not be used for learning MAC of connected device. The user-configured MAC address will be used for auto-filling the destination
	// MAC address in the control and data packets sent from this IPv6 endpoint
	// whenever applicable.
	SetGatewayMac(value DeviceIpv6GatewayMAC) DeviceIpv6
	// HasGatewayMac checks if GatewayMac has been set in DeviceIpv6
	HasGatewayMac() bool
	// Address returns string, set in DeviceIpv6.
	Address() string
	// SetAddress assigns string provided by user to DeviceIpv6
	SetAddress(value string) DeviceIpv6
	// Prefix returns uint32, set in DeviceIpv6.
	Prefix() uint32
	// SetPrefix assigns uint32 provided by user to DeviceIpv6
	SetPrefix(value uint32) DeviceIpv6
	// HasPrefix checks if Prefix has been set in DeviceIpv6
	HasPrefix() bool
	// Name returns string, set in DeviceIpv6.
	Name() string
	// SetName assigns string provided by user to DeviceIpv6
	SetName(value string) DeviceIpv6
	setNil()
}

// The IPv6 gateway address.
// Gateway returns a string
func (obj *deviceIpv6) Gateway() string {

	return *obj.obj.Gateway

}

// The IPv6 gateway address.
// SetGateway sets the string value in the DeviceIpv6 object
func (obj *deviceIpv6) SetGateway(value string) DeviceIpv6 {

	obj.obj.Gateway = &value
	return obj
}

// description is TBD
// GatewayMac returns a DeviceIpv6GatewayMAC
func (obj *deviceIpv6) GatewayMac() DeviceIpv6GatewayMAC {
	if obj.obj.GatewayMac == nil {
		obj.obj.GatewayMac = NewDeviceIpv6GatewayMAC().msg()
	}
	if obj.gatewayMacHolder == nil {
		obj.gatewayMacHolder = &deviceIpv6GatewayMAC{obj: obj.obj.GatewayMac}
	}
	return obj.gatewayMacHolder
}

// description is TBD
// GatewayMac returns a DeviceIpv6GatewayMAC
func (obj *deviceIpv6) HasGatewayMac() bool {
	return obj.obj.GatewayMac != nil
}

// description is TBD
// SetGatewayMac sets the DeviceIpv6GatewayMAC value in the DeviceIpv6 object
func (obj *deviceIpv6) SetGatewayMac(value DeviceIpv6GatewayMAC) DeviceIpv6 {

	obj.gatewayMacHolder = nil
	obj.obj.GatewayMac = value.msg()

	return obj
}

// The IPv6 address.
// Address returns a string
func (obj *deviceIpv6) Address() string {

	return *obj.obj.Address

}

// The IPv6 address.
// SetAddress sets the string value in the DeviceIpv6 object
func (obj *deviceIpv6) SetAddress(value string) DeviceIpv6 {

	obj.obj.Address = &value
	return obj
}

// The network prefix.
// Prefix returns a uint32
func (obj *deviceIpv6) Prefix() uint32 {

	return *obj.obj.Prefix

}

// The network prefix.
// Prefix returns a uint32
func (obj *deviceIpv6) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The network prefix.
// SetPrefix sets the uint32 value in the DeviceIpv6 object
func (obj *deviceIpv6) SetPrefix(value uint32) DeviceIpv6 {

	obj.obj.Prefix = &value
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceIpv6) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceIpv6 object
func (obj *deviceIpv6) SetName(value string) DeviceIpv6 {

	obj.obj.Name = &value
	return obj
}

func (obj *deviceIpv6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Gateway is required
	if obj.obj.Gateway == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Gateway is required field on interface DeviceIpv6")
	}
	if obj.obj.Gateway != nil {

		err := obj.validateIpv6(obj.Gateway())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv6.Gateway"))
		}

	}

	if obj.obj.GatewayMac != nil {

		obj.GatewayMac().validateObj(vObj, set_default)
	}

	// Address is required
	if obj.obj.Address == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Address is required field on interface DeviceIpv6")
	}
	if obj.obj.Address != nil {

		err := obj.validateIpv6(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv6.Address"))
		}

	}

	if obj.obj.Prefix != nil {

		if *obj.obj.Prefix < 1 || *obj.obj.Prefix > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceIpv6.Prefix <= 128 but Got %d", *obj.obj.Prefix))
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceIpv6")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"DeviceIpv6.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

}

func (obj *deviceIpv6) setDefault() {
	if obj.obj.Prefix == nil {
		obj.SetPrefix(64)
	}

}
