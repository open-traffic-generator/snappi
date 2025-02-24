package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceDhcpv6ClientDuidType *****
type deviceDhcpv6ClientDuidType struct {
	validation
	obj          *otg.DeviceDhcpv6ClientDuidType
	marshaller   marshalDeviceDhcpv6ClientDuidType
	unMarshaller unMarshalDeviceDhcpv6ClientDuidType
	lltHolder    DeviceDhcpv6ClientNoDuid
	enHolder     DeviceDhcpv6ClientDuidValue
	llHolder     DeviceDhcpv6ClientNoDuid
}

func NewDeviceDhcpv6ClientDuidType() DeviceDhcpv6ClientDuidType {
	obj := deviceDhcpv6ClientDuidType{obj: &otg.DeviceDhcpv6ClientDuidType{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceDhcpv6ClientDuidType) msg() *otg.DeviceDhcpv6ClientDuidType {
	return obj.obj
}

func (obj *deviceDhcpv6ClientDuidType) setMsg(msg *otg.DeviceDhcpv6ClientDuidType) DeviceDhcpv6ClientDuidType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceDhcpv6ClientDuidType struct {
	obj *deviceDhcpv6ClientDuidType
}

type marshalDeviceDhcpv6ClientDuidType interface {
	// ToProto marshals DeviceDhcpv6ClientDuidType to protobuf object *otg.DeviceDhcpv6ClientDuidType
	ToProto() (*otg.DeviceDhcpv6ClientDuidType, error)
	// ToPbText marshals DeviceDhcpv6ClientDuidType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceDhcpv6ClientDuidType to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceDhcpv6ClientDuidType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals DeviceDhcpv6ClientDuidType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldeviceDhcpv6ClientDuidType struct {
	obj *deviceDhcpv6ClientDuidType
}

type unMarshalDeviceDhcpv6ClientDuidType interface {
	// FromProto unmarshals DeviceDhcpv6ClientDuidType from protobuf object *otg.DeviceDhcpv6ClientDuidType
	FromProto(msg *otg.DeviceDhcpv6ClientDuidType) (DeviceDhcpv6ClientDuidType, error)
	// FromPbText unmarshals DeviceDhcpv6ClientDuidType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceDhcpv6ClientDuidType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceDhcpv6ClientDuidType from JSON text
	FromJson(value string) error
}

func (obj *deviceDhcpv6ClientDuidType) Marshal() marshalDeviceDhcpv6ClientDuidType {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceDhcpv6ClientDuidType{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceDhcpv6ClientDuidType) Unmarshal() unMarshalDeviceDhcpv6ClientDuidType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceDhcpv6ClientDuidType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceDhcpv6ClientDuidType) ToProto() (*otg.DeviceDhcpv6ClientDuidType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceDhcpv6ClientDuidType) FromProto(msg *otg.DeviceDhcpv6ClientDuidType) (DeviceDhcpv6ClientDuidType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceDhcpv6ClientDuidType) ToPbText() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientDuidType) FromPbText(value string) error {
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

func (m *marshaldeviceDhcpv6ClientDuidType) ToYaml() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientDuidType) FromYaml(value string) error {
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

func (m *marshaldeviceDhcpv6ClientDuidType) ToJsonRaw() (string, error) {
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

func (m *marshaldeviceDhcpv6ClientDuidType) ToJson() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientDuidType) FromJson(value string) error {
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

func (obj *deviceDhcpv6ClientDuidType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientDuidType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientDuidType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceDhcpv6ClientDuidType) Clone() (DeviceDhcpv6ClientDuidType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceDhcpv6ClientDuidType()
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

func (obj *deviceDhcpv6ClientDuidType) setNil() {
	obj.lltHolder = nil
	obj.enHolder = nil
	obj.llHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceDhcpv6ClientDuidType is description is TBD
type DeviceDhcpv6ClientDuidType interface {
	Validation
	// msg marshals DeviceDhcpv6ClientDuidType to protobuf object *otg.DeviceDhcpv6ClientDuidType
	// and doesn't set defaults
	msg() *otg.DeviceDhcpv6ClientDuidType
	// setMsg unmarshals DeviceDhcpv6ClientDuidType from protobuf object *otg.DeviceDhcpv6ClientDuidType
	// and doesn't set defaults
	setMsg(*otg.DeviceDhcpv6ClientDuidType) DeviceDhcpv6ClientDuidType
	// provides marshal interface
	Marshal() marshalDeviceDhcpv6ClientDuidType
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceDhcpv6ClientDuidType
	// validate validates DeviceDhcpv6ClientDuidType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceDhcpv6ClientDuidType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns DeviceDhcpv6ClientDuidTypeChoiceEnum, set in DeviceDhcpv6ClientDuidType
	Choice() DeviceDhcpv6ClientDuidTypeChoiceEnum
	// setChoice assigns DeviceDhcpv6ClientDuidTypeChoiceEnum provided by user to DeviceDhcpv6ClientDuidType
	setChoice(value DeviceDhcpv6ClientDuidTypeChoiceEnum) DeviceDhcpv6ClientDuidType
	// HasChoice checks if Choice has been set in DeviceDhcpv6ClientDuidType
	HasChoice() bool
	// Llt returns DeviceDhcpv6ClientNoDuid, set in DeviceDhcpv6ClientDuidType.
	Llt() DeviceDhcpv6ClientNoDuid
	// SetLlt assigns DeviceDhcpv6ClientNoDuid provided by user to DeviceDhcpv6ClientDuidType.
	SetLlt(value DeviceDhcpv6ClientNoDuid) DeviceDhcpv6ClientDuidType
	// HasLlt checks if Llt has been set in DeviceDhcpv6ClientDuidType
	HasLlt() bool
	// En returns DeviceDhcpv6ClientDuidValue, set in DeviceDhcpv6ClientDuidType.
	En() DeviceDhcpv6ClientDuidValue
	// SetEn assigns DeviceDhcpv6ClientDuidValue provided by user to DeviceDhcpv6ClientDuidType.
	SetEn(value DeviceDhcpv6ClientDuidValue) DeviceDhcpv6ClientDuidType
	// HasEn checks if En has been set in DeviceDhcpv6ClientDuidType
	HasEn() bool
	// Ll returns DeviceDhcpv6ClientNoDuid, set in DeviceDhcpv6ClientDuidType.
	Ll() DeviceDhcpv6ClientNoDuid
	// SetLl assigns DeviceDhcpv6ClientNoDuid provided by user to DeviceDhcpv6ClientDuidType.
	SetLl(value DeviceDhcpv6ClientNoDuid) DeviceDhcpv6ClientDuidType
	// HasLl checks if Ll has been set in DeviceDhcpv6ClientDuidType
	HasLl() bool
	setNil()
}

type DeviceDhcpv6ClientDuidTypeChoiceEnum string

// Enum of Choice on DeviceDhcpv6ClientDuidType
var DeviceDhcpv6ClientDuidTypeChoice = struct {
	LLT DeviceDhcpv6ClientDuidTypeChoiceEnum
	EN  DeviceDhcpv6ClientDuidTypeChoiceEnum
	LL  DeviceDhcpv6ClientDuidTypeChoiceEnum
}{
	LLT: DeviceDhcpv6ClientDuidTypeChoiceEnum("llt"),
	EN:  DeviceDhcpv6ClientDuidTypeChoiceEnum("en"),
	LL:  DeviceDhcpv6ClientDuidTypeChoiceEnum("ll"),
}

func (obj *deviceDhcpv6ClientDuidType) Choice() DeviceDhcpv6ClientDuidTypeChoiceEnum {
	return DeviceDhcpv6ClientDuidTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// Each DHCP client and server has a DUID. DHCP clients use DUIDs to identify a server in messages where a server needs  to be identified.
// Choice returns a string
func (obj *deviceDhcpv6ClientDuidType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *deviceDhcpv6ClientDuidType) setChoice(value DeviceDhcpv6ClientDuidTypeChoiceEnum) DeviceDhcpv6ClientDuidType {
	intValue, ok := otg.DeviceDhcpv6ClientDuidType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceDhcpv6ClientDuidTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceDhcpv6ClientDuidType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ll = nil
	obj.llHolder = nil
	obj.obj.En = nil
	obj.enHolder = nil
	obj.obj.Llt = nil
	obj.lltHolder = nil

	if value == DeviceDhcpv6ClientDuidTypeChoice.LLT {
		obj.obj.Llt = NewDeviceDhcpv6ClientNoDuid().msg()
	}

	if value == DeviceDhcpv6ClientDuidTypeChoice.EN {
		obj.obj.En = NewDeviceDhcpv6ClientDuidValue().msg()
	}

	if value == DeviceDhcpv6ClientDuidTypeChoice.LL {
		obj.obj.Ll = NewDeviceDhcpv6ClientNoDuid().msg()
	}

	return obj
}

// description is TBD
// Llt returns a DeviceDhcpv6ClientNoDuid
func (obj *deviceDhcpv6ClientDuidType) Llt() DeviceDhcpv6ClientNoDuid {
	if obj.obj.Llt == nil {
		obj.setChoice(DeviceDhcpv6ClientDuidTypeChoice.LLT)
	}
	if obj.lltHolder == nil {
		obj.lltHolder = &deviceDhcpv6ClientNoDuid{obj: obj.obj.Llt}
	}
	return obj.lltHolder
}

// description is TBD
// Llt returns a DeviceDhcpv6ClientNoDuid
func (obj *deviceDhcpv6ClientDuidType) HasLlt() bool {
	return obj.obj.Llt != nil
}

// description is TBD
// SetLlt sets the DeviceDhcpv6ClientNoDuid value in the DeviceDhcpv6ClientDuidType object
func (obj *deviceDhcpv6ClientDuidType) SetLlt(value DeviceDhcpv6ClientNoDuid) DeviceDhcpv6ClientDuidType {
	obj.setChoice(DeviceDhcpv6ClientDuidTypeChoice.LLT)
	obj.lltHolder = nil
	obj.obj.Llt = value.msg()

	return obj
}

// description is TBD
// En returns a DeviceDhcpv6ClientDuidValue
func (obj *deviceDhcpv6ClientDuidType) En() DeviceDhcpv6ClientDuidValue {
	if obj.obj.En == nil {
		obj.setChoice(DeviceDhcpv6ClientDuidTypeChoice.EN)
	}
	if obj.enHolder == nil {
		obj.enHolder = &deviceDhcpv6ClientDuidValue{obj: obj.obj.En}
	}
	return obj.enHolder
}

// description is TBD
// En returns a DeviceDhcpv6ClientDuidValue
func (obj *deviceDhcpv6ClientDuidType) HasEn() bool {
	return obj.obj.En != nil
}

// description is TBD
// SetEn sets the DeviceDhcpv6ClientDuidValue value in the DeviceDhcpv6ClientDuidType object
func (obj *deviceDhcpv6ClientDuidType) SetEn(value DeviceDhcpv6ClientDuidValue) DeviceDhcpv6ClientDuidType {
	obj.setChoice(DeviceDhcpv6ClientDuidTypeChoice.EN)
	obj.enHolder = nil
	obj.obj.En = value.msg()

	return obj
}

// description is TBD
// Ll returns a DeviceDhcpv6ClientNoDuid
func (obj *deviceDhcpv6ClientDuidType) Ll() DeviceDhcpv6ClientNoDuid {
	if obj.obj.Ll == nil {
		obj.setChoice(DeviceDhcpv6ClientDuidTypeChoice.LL)
	}
	if obj.llHolder == nil {
		obj.llHolder = &deviceDhcpv6ClientNoDuid{obj: obj.obj.Ll}
	}
	return obj.llHolder
}

// description is TBD
// Ll returns a DeviceDhcpv6ClientNoDuid
func (obj *deviceDhcpv6ClientDuidType) HasLl() bool {
	return obj.obj.Ll != nil
}

// description is TBD
// SetLl sets the DeviceDhcpv6ClientNoDuid value in the DeviceDhcpv6ClientDuidType object
func (obj *deviceDhcpv6ClientDuidType) SetLl(value DeviceDhcpv6ClientNoDuid) DeviceDhcpv6ClientDuidType {
	obj.setChoice(DeviceDhcpv6ClientDuidTypeChoice.LL)
	obj.llHolder = nil
	obj.obj.Ll = value.msg()

	return obj
}

func (obj *deviceDhcpv6ClientDuidType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Llt != nil {

		obj.Llt().validateObj(vObj, set_default)
	}

	if obj.obj.En != nil {

		obj.En().validateObj(vObj, set_default)
	}

	if obj.obj.Ll != nil {

		obj.Ll().validateObj(vObj, set_default)
	}

}

func (obj *deviceDhcpv6ClientDuidType) setDefault() {
	var choices_set int = 0
	var choice DeviceDhcpv6ClientDuidTypeChoiceEnum

	if obj.obj.Llt != nil {
		choices_set += 1
		choice = DeviceDhcpv6ClientDuidTypeChoice.LLT
	}

	if obj.obj.En != nil {
		choices_set += 1
		choice = DeviceDhcpv6ClientDuidTypeChoice.EN
	}

	if obj.obj.Ll != nil {
		choices_set += 1
		choice = DeviceDhcpv6ClientDuidTypeChoice.LL
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(DeviceDhcpv6ClientDuidTypeChoice.LLT)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in DeviceDhcpv6ClientDuidType")
			}
		} else {
			intVal := otg.DeviceDhcpv6ClientDuidType_Choice_Enum_value[string(choice)]
			enumValue := otg.DeviceDhcpv6ClientDuidType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
