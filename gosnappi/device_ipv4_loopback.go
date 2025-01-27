package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceIpv4Loopback *****
type deviceIpv4Loopback struct {
	validation
	obj          *otg.DeviceIpv4Loopback
	marshaller   marshalDeviceIpv4Loopback
	unMarshaller unMarshalDeviceIpv4Loopback
}

func NewDeviceIpv4Loopback() DeviceIpv4Loopback {
	obj := deviceIpv4Loopback{obj: &otg.DeviceIpv4Loopback{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceIpv4Loopback) msg() *otg.DeviceIpv4Loopback {
	return obj.obj
}

func (obj *deviceIpv4Loopback) setMsg(msg *otg.DeviceIpv4Loopback) DeviceIpv4Loopback {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceIpv4Loopback struct {
	obj *deviceIpv4Loopback
}

type marshalDeviceIpv4Loopback interface {
	// ToProto marshals DeviceIpv4Loopback to protobuf object *otg.DeviceIpv4Loopback
	ToProto() (*otg.DeviceIpv4Loopback, error)
	// ToPbText marshals DeviceIpv4Loopback to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceIpv4Loopback to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceIpv4Loopback to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals DeviceIpv4Loopback to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldeviceIpv4Loopback struct {
	obj *deviceIpv4Loopback
}

type unMarshalDeviceIpv4Loopback interface {
	// FromProto unmarshals DeviceIpv4Loopback from protobuf object *otg.DeviceIpv4Loopback
	FromProto(msg *otg.DeviceIpv4Loopback) (DeviceIpv4Loopback, error)
	// FromPbText unmarshals DeviceIpv4Loopback from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceIpv4Loopback from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceIpv4Loopback from JSON text
	FromJson(value string) error
}

func (obj *deviceIpv4Loopback) Marshal() marshalDeviceIpv4Loopback {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceIpv4Loopback{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceIpv4Loopback) Unmarshal() unMarshalDeviceIpv4Loopback {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceIpv4Loopback{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceIpv4Loopback) ToProto() (*otg.DeviceIpv4Loopback, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceIpv4Loopback) FromProto(msg *otg.DeviceIpv4Loopback) (DeviceIpv4Loopback, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceIpv4Loopback) ToPbText() (string, error) {
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

func (m *unMarshaldeviceIpv4Loopback) FromPbText(value string) error {
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

func (m *marshaldeviceIpv4Loopback) ToYaml() (string, error) {
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

func (m *unMarshaldeviceIpv4Loopback) FromYaml(value string) error {
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

func (m *marshaldeviceIpv4Loopback) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshaldeviceIpv4Loopback) ToJson() (string, error) {
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

func (m *unMarshaldeviceIpv4Loopback) FromJson(value string) error {
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

func (obj *deviceIpv4Loopback) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceIpv4Loopback) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceIpv4Loopback) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceIpv4Loopback) Clone() (DeviceIpv4Loopback, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceIpv4Loopback()
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

// DeviceIpv4Loopback is an IPv4 Loopback interface.
type DeviceIpv4Loopback interface {
	Validation
	// msg marshals DeviceIpv4Loopback to protobuf object *otg.DeviceIpv4Loopback
	// and doesn't set defaults
	msg() *otg.DeviceIpv4Loopback
	// setMsg unmarshals DeviceIpv4Loopback from protobuf object *otg.DeviceIpv4Loopback
	// and doesn't set defaults
	setMsg(*otg.DeviceIpv4Loopback) DeviceIpv4Loopback
	// provides marshal interface
	Marshal() marshalDeviceIpv4Loopback
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceIpv4Loopback
	// validate validates DeviceIpv4Loopback
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceIpv4Loopback, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthName returns string, set in DeviceIpv4Loopback.
	EthName() string
	// SetEthName assigns string provided by user to DeviceIpv4Loopback
	SetEthName(value string) DeviceIpv4Loopback
	// Address returns string, set in DeviceIpv4Loopback.
	Address() string
	// SetAddress assigns string provided by user to DeviceIpv4Loopback
	SetAddress(value string) DeviceIpv4Loopback
	// HasAddress checks if Address has been set in DeviceIpv4Loopback
	HasAddress() bool
	// Name returns string, set in DeviceIpv4Loopback.
	Name() string
	// SetName assigns string provided by user to DeviceIpv4Loopback
	SetName(value string) DeviceIpv4Loopback
}

// The unique name of the Ethernet interface behind which this Loopback  interface will be created.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// EthName returns a string
func (obj *deviceIpv4Loopback) EthName() string {

	return *obj.obj.EthName

}

// The unique name of the Ethernet interface behind which this Loopback  interface will be created.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// SetEthName sets the string value in the DeviceIpv4Loopback object
func (obj *deviceIpv4Loopback) SetEthName(value string) DeviceIpv4Loopback {

	obj.obj.EthName = &value
	return obj
}

// The IPv4 Loopback address with prefix length of 32.
// Address returns a string
func (obj *deviceIpv4Loopback) Address() string {

	return *obj.obj.Address

}

// The IPv4 Loopback address with prefix length of 32.
// Address returns a string
func (obj *deviceIpv4Loopback) HasAddress() bool {
	return obj.obj.Address != nil
}

// The IPv4 Loopback address with prefix length of 32.
// SetAddress sets the string value in the DeviceIpv4Loopback object
func (obj *deviceIpv4Loopback) SetAddress(value string) DeviceIpv4Loopback {

	obj.obj.Address = &value
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceIpv4Loopback) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceIpv4Loopback object
func (obj *deviceIpv4Loopback) SetName(value string) DeviceIpv4Loopback {

	obj.obj.Name = &value
	return obj
}

func (obj *deviceIpv4Loopback) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EthName is required
	if obj.obj.EthName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EthName is required field on interface DeviceIpv4Loopback")
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv4(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceIpv4Loopback.Address"))
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceIpv4Loopback")
	}
}

func (obj *deviceIpv4Loopback) setDefault() {
	if obj.obj.Address == nil {
		obj.SetAddress("0.0.0.0")
	}

}
