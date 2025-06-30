package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceIpv6Loopback *****
type deviceIpv6Loopback struct {
	validation
	obj          *otg.DeviceIpv6Loopback
	marshaller   marshalDeviceIpv6Loopback
	unMarshaller unMarshalDeviceIpv6Loopback
}

func NewDeviceIpv6Loopback() DeviceIpv6Loopback {
	obj := deviceIpv6Loopback{obj: &otg.DeviceIpv6Loopback{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceIpv6Loopback) msg() *otg.DeviceIpv6Loopback {
	return obj.obj
}

func (obj *deviceIpv6Loopback) setMsg(msg *otg.DeviceIpv6Loopback) DeviceIpv6Loopback {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceIpv6Loopback struct {
	obj *deviceIpv6Loopback
}

type marshalDeviceIpv6Loopback interface {
	// ToProto marshals DeviceIpv6Loopback to protobuf object *otg.DeviceIpv6Loopback
	ToProto() (*otg.DeviceIpv6Loopback, error)
	// ToPbText marshals DeviceIpv6Loopback to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceIpv6Loopback to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceIpv6Loopback to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceIpv6Loopback struct {
	obj *deviceIpv6Loopback
}

type unMarshalDeviceIpv6Loopback interface {
	// FromProto unmarshals DeviceIpv6Loopback from protobuf object *otg.DeviceIpv6Loopback
	FromProto(msg *otg.DeviceIpv6Loopback) (DeviceIpv6Loopback, error)
	// FromPbText unmarshals DeviceIpv6Loopback from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceIpv6Loopback from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceIpv6Loopback from JSON text
	FromJson(value string) error
}

func (obj *deviceIpv6Loopback) Marshal() marshalDeviceIpv6Loopback {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceIpv6Loopback{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceIpv6Loopback) Unmarshal() unMarshalDeviceIpv6Loopback {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceIpv6Loopback{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceIpv6Loopback) ToProto() (*otg.DeviceIpv6Loopback, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceIpv6Loopback) FromProto(msg *otg.DeviceIpv6Loopback) (DeviceIpv6Loopback, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceIpv6Loopback) ToPbText() (string, error) {
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

func (m *unMarshaldeviceIpv6Loopback) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshaldeviceIpv6Loopback) ToYaml() (string, error) {
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

func (m *unMarshaldeviceIpv6Loopback) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshaldeviceIpv6Loopback) ToJson() (string, error) {
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

func (m *unMarshaldeviceIpv6Loopback) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *deviceIpv6Loopback) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceIpv6Loopback) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceIpv6Loopback) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceIpv6Loopback) Clone() (DeviceIpv6Loopback, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceIpv6Loopback()
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

// DeviceIpv6Loopback is an IPv6 Loopback interface
type DeviceIpv6Loopback interface {
	Validation
	// msg marshals DeviceIpv6Loopback to protobuf object *otg.DeviceIpv6Loopback
	// and doesn't set defaults
	msg() *otg.DeviceIpv6Loopback
	// setMsg unmarshals DeviceIpv6Loopback from protobuf object *otg.DeviceIpv6Loopback
	// and doesn't set defaults
	setMsg(*otg.DeviceIpv6Loopback) DeviceIpv6Loopback
	// provides marshal interface
	Marshal() marshalDeviceIpv6Loopback
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceIpv6Loopback
	// validate validates DeviceIpv6Loopback
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceIpv6Loopback, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthName returns string, set in DeviceIpv6Loopback.
	EthName() string
	// SetEthName assigns string provided by user to DeviceIpv6Loopback
	SetEthName(value string) DeviceIpv6Loopback
	// Address returns string, set in DeviceIpv6Loopback.
	Address() string
	// SetAddress assigns string provided by user to DeviceIpv6Loopback
	SetAddress(value string) DeviceIpv6Loopback
	// HasAddress checks if Address has been set in DeviceIpv6Loopback
	HasAddress() bool
	// Name returns string, set in DeviceIpv6Loopback.
	Name() string
	// SetName assigns string provided by user to DeviceIpv6Loopback
	SetName(value string) DeviceIpv6Loopback
}

// The unique name of the Ethernet interface behind which this Loopback
// interface will be created.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// EthName returns a string
func (obj *deviceIpv6Loopback) EthName() string {

	return *obj.obj.EthName

}

// The unique name of the Ethernet interface behind which this Loopback
// interface will be created.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// SetEthName sets the string value in the DeviceIpv6Loopback object
func (obj *deviceIpv6Loopback) SetEthName(value string) DeviceIpv6Loopback {

	obj.obj.EthName = &value
	return obj
}

// The IPv6 Loopback address with prefix length of 128.
// Address returns a string
func (obj *deviceIpv6Loopback) Address() string {

	return *obj.obj.Address

}

// The IPv6 Loopback address with prefix length of 128.
// Address returns a string
func (obj *deviceIpv6Loopback) HasAddress() bool {
	return obj.obj.Address != nil
}

// The IPv6 Loopback address with prefix length of 128.
// SetAddress sets the string value in the DeviceIpv6Loopback object
func (obj *deviceIpv6Loopback) SetAddress(value string) DeviceIpv6Loopback {

	obj.obj.Address = &value
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceIpv6Loopback) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceIpv6Loopback object
func (obj *deviceIpv6Loopback) SetName(value string) DeviceIpv6Loopback {

	obj.obj.Name = &value
	return obj
}

func (obj *deviceIpv6Loopback) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EthName is required
	if obj.obj.EthName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EthName is required field on interface DeviceIpv6Loopback")
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv6(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv6Loopback.Address"))
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceIpv6Loopback")
	}
}

func (obj *deviceIpv6Loopback) setDefault() {
	if obj.obj.Address == nil {
		obj.SetAddress("::0")
	}

}
