package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceVlan *****
type deviceVlan struct {
	validation
	obj          *otg.DeviceVlan
	marshaller   marshalDeviceVlan
	unMarshaller unMarshalDeviceVlan
}

func NewDeviceVlan() DeviceVlan {
	obj := deviceVlan{obj: &otg.DeviceVlan{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceVlan) msg() *otg.DeviceVlan {
	return obj.obj
}

func (obj *deviceVlan) setMsg(msg *otg.DeviceVlan) DeviceVlan {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceVlan struct {
	obj *deviceVlan
}

type marshalDeviceVlan interface {
	// ToProto marshals DeviceVlan to protobuf object *otg.DeviceVlan
	ToProto() (*otg.DeviceVlan, error)
	// ToPbText marshals DeviceVlan to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceVlan to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceVlan to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals DeviceVlan to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldeviceVlan struct {
	obj *deviceVlan
}

type unMarshalDeviceVlan interface {
	// FromProto unmarshals DeviceVlan from protobuf object *otg.DeviceVlan
	FromProto(msg *otg.DeviceVlan) (DeviceVlan, error)
	// FromPbText unmarshals DeviceVlan from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceVlan from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceVlan from JSON text
	FromJson(value string) error
}

func (obj *deviceVlan) Marshal() marshalDeviceVlan {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceVlan{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceVlan) Unmarshal() unMarshalDeviceVlan {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceVlan{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceVlan) ToProto() (*otg.DeviceVlan, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceVlan) FromProto(msg *otg.DeviceVlan) (DeviceVlan, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceVlan) ToPbText() (string, error) {
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

func (m *unMarshaldeviceVlan) FromPbText(value string) error {
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

func (m *marshaldeviceVlan) ToYaml() (string, error) {
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

func (m *unMarshaldeviceVlan) FromYaml(value string) error {
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

func (m *marshaldeviceVlan) ToJsonRaw() (string, error) {
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

func (m *marshaldeviceVlan) ToJson() (string, error) {
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

func (m *unMarshaldeviceVlan) FromJson(value string) error {
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

func (obj *deviceVlan) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceVlan) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceVlan) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceVlan) Clone() (DeviceVlan, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceVlan()
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

// DeviceVlan is emulated VLAN protocol.
type DeviceVlan interface {
	Validation
	// msg marshals DeviceVlan to protobuf object *otg.DeviceVlan
	// and doesn't set defaults
	msg() *otg.DeviceVlan
	// setMsg unmarshals DeviceVlan from protobuf object *otg.DeviceVlan
	// and doesn't set defaults
	setMsg(*otg.DeviceVlan) DeviceVlan
	// provides marshal interface
	Marshal() marshalDeviceVlan
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceVlan
	// validate validates DeviceVlan
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceVlan, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Tpid returns DeviceVlanTpidEnum, set in DeviceVlan
	Tpid() DeviceVlanTpidEnum
	// SetTpid assigns DeviceVlanTpidEnum provided by user to DeviceVlan
	SetTpid(value DeviceVlanTpidEnum) DeviceVlan
	// HasTpid checks if Tpid has been set in DeviceVlan
	HasTpid() bool
	// Priority returns uint32, set in DeviceVlan.
	Priority() uint32
	// SetPriority assigns uint32 provided by user to DeviceVlan
	SetPriority(value uint32) DeviceVlan
	// HasPriority checks if Priority has been set in DeviceVlan
	HasPriority() bool
	// Id returns uint32, set in DeviceVlan.
	Id() uint32
	// SetId assigns uint32 provided by user to DeviceVlan
	SetId(value uint32) DeviceVlan
	// HasId checks if Id has been set in DeviceVlan
	HasId() bool
	// Name returns string, set in DeviceVlan.
	Name() string
	// SetName assigns string provided by user to DeviceVlan
	SetName(value string) DeviceVlan
}

type DeviceVlanTpidEnum string

// Enum of Tpid on DeviceVlan
var DeviceVlanTpid = struct {
	X8100 DeviceVlanTpidEnum
	X88A8 DeviceVlanTpidEnum
	X9100 DeviceVlanTpidEnum
	X9200 DeviceVlanTpidEnum
	X9300 DeviceVlanTpidEnum
}{
	X8100: DeviceVlanTpidEnum("x8100"),
	X88A8: DeviceVlanTpidEnum("x88A8"),
	X9100: DeviceVlanTpidEnum("x9100"),
	X9200: DeviceVlanTpidEnum("x9200"),
	X9300: DeviceVlanTpidEnum("x9300"),
}

func (obj *deviceVlan) Tpid() DeviceVlanTpidEnum {
	return DeviceVlanTpidEnum(obj.obj.Tpid.Enum().String())
}

// Tag protocol identifier
// Tpid returns a string
func (obj *deviceVlan) HasTpid() bool {
	return obj.obj.Tpid != nil
}

func (obj *deviceVlan) SetTpid(value DeviceVlanTpidEnum) DeviceVlan {
	intValue, ok := otg.DeviceVlan_Tpid_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceVlanTpidEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceVlan_Tpid_Enum(intValue)
	obj.obj.Tpid = &enumValue

	return obj
}

// Priority code point
// Priority returns a uint32
func (obj *deviceVlan) Priority() uint32 {

	return *obj.obj.Priority

}

// Priority code point
// Priority returns a uint32
func (obj *deviceVlan) HasPriority() bool {
	return obj.obj.Priority != nil
}

// Priority code point
// SetPriority sets the uint32 value in the DeviceVlan object
func (obj *deviceVlan) SetPriority(value uint32) DeviceVlan {

	obj.obj.Priority = &value
	return obj
}

// VLAN identifier
// Id returns a uint32
func (obj *deviceVlan) Id() uint32 {

	return *obj.obj.Id

}

// VLAN identifier
// Id returns a uint32
func (obj *deviceVlan) HasId() bool {
	return obj.obj.Id != nil
}

// VLAN identifier
// SetId sets the uint32 value in the DeviceVlan object
func (obj *deviceVlan) SetId(value uint32) DeviceVlan {

	obj.obj.Id = &value
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceVlan) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceVlan object
func (obj *deviceVlan) SetName(value string) DeviceVlan {

	obj.obj.Name = &value
	return obj
}

func (obj *deviceVlan) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Priority != nil {

		if *obj.obj.Priority > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= DeviceVlan.Priority <= 7 but Got %d", *obj.obj.Priority))
		}

	}

	if obj.obj.Id != nil {

		if *obj.obj.Id > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= DeviceVlan.Id <= 4095 but Got %d", *obj.obj.Id))
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceVlan")
	}
}

func (obj *deviceVlan) setDefault() {
	if obj.obj.Tpid == nil {
		obj.SetTpid(DeviceVlanTpid.X8100)

	}
	if obj.obj.Priority == nil {
		obj.SetPriority(0)
	}
	if obj.obj.Id == nil {
		obj.SetId(1)
	}

}
