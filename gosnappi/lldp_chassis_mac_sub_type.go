package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpChassisMacSubType *****
type lldpChassisMacSubType struct {
	validation
	obj          *otg.LldpChassisMacSubType
	marshaller   marshalLldpChassisMacSubType
	unMarshaller unMarshalLldpChassisMacSubType
}

func NewLldpChassisMacSubType() LldpChassisMacSubType {
	obj := lldpChassisMacSubType{obj: &otg.LldpChassisMacSubType{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpChassisMacSubType) msg() *otg.LldpChassisMacSubType {
	return obj.obj
}

func (obj *lldpChassisMacSubType) setMsg(msg *otg.LldpChassisMacSubType) LldpChassisMacSubType {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpChassisMacSubType struct {
	obj *lldpChassisMacSubType
}

type marshalLldpChassisMacSubType interface {
	// ToProto marshals LldpChassisMacSubType to protobuf object *otg.LldpChassisMacSubType
	ToProto() (*otg.LldpChassisMacSubType, error)
	// ToPbText marshals LldpChassisMacSubType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpChassisMacSubType to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpChassisMacSubType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LldpChassisMacSubType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallldpChassisMacSubType struct {
	obj *lldpChassisMacSubType
}

type unMarshalLldpChassisMacSubType interface {
	// FromProto unmarshals LldpChassisMacSubType from protobuf object *otg.LldpChassisMacSubType
	FromProto(msg *otg.LldpChassisMacSubType) (LldpChassisMacSubType, error)
	// FromPbText unmarshals LldpChassisMacSubType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpChassisMacSubType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpChassisMacSubType from JSON text
	FromJson(value string) error
}

func (obj *lldpChassisMacSubType) Marshal() marshalLldpChassisMacSubType {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpChassisMacSubType{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpChassisMacSubType) Unmarshal() unMarshalLldpChassisMacSubType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpChassisMacSubType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpChassisMacSubType) ToProto() (*otg.LldpChassisMacSubType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpChassisMacSubType) FromProto(msg *otg.LldpChassisMacSubType) (LldpChassisMacSubType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpChassisMacSubType) ToPbText() (string, error) {
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

func (m *unMarshallldpChassisMacSubType) FromPbText(value string) error {
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

func (m *marshallldpChassisMacSubType) ToYaml() (string, error) {
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

func (m *unMarshallldpChassisMacSubType) FromYaml(value string) error {
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

func (m *marshallldpChassisMacSubType) ToJsonRaw() (string, error) {
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

func (m *marshallldpChassisMacSubType) ToJson() (string, error) {
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

func (m *unMarshallldpChassisMacSubType) FromJson(value string) error {
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

func (obj *lldpChassisMacSubType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpChassisMacSubType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpChassisMacSubType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpChassisMacSubType) Clone() (LldpChassisMacSubType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpChassisMacSubType()
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

// LldpChassisMacSubType is the MAC address configured in the Chassis ID TLV.
type LldpChassisMacSubType interface {
	Validation
	// msg marshals LldpChassisMacSubType to protobuf object *otg.LldpChassisMacSubType
	// and doesn't set defaults
	msg() *otg.LldpChassisMacSubType
	// setMsg unmarshals LldpChassisMacSubType from protobuf object *otg.LldpChassisMacSubType
	// and doesn't set defaults
	setMsg(*otg.LldpChassisMacSubType) LldpChassisMacSubType
	// provides marshal interface
	Marshal() marshalLldpChassisMacSubType
	// provides unmarshal interface
	Unmarshal() unMarshalLldpChassisMacSubType
	// validate validates LldpChassisMacSubType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpChassisMacSubType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LldpChassisMacSubTypeChoiceEnum, set in LldpChassisMacSubType
	Choice() LldpChassisMacSubTypeChoiceEnum
	// setChoice assigns LldpChassisMacSubTypeChoiceEnum provided by user to LldpChassisMacSubType
	setChoice(value LldpChassisMacSubTypeChoiceEnum) LldpChassisMacSubType
	// HasChoice checks if Choice has been set in LldpChassisMacSubType
	HasChoice() bool
	// Auto returns string, set in LldpChassisMacSubType.
	Auto() string
	// HasAuto checks if Auto has been set in LldpChassisMacSubType
	HasAuto() bool
	// Value returns string, set in LldpChassisMacSubType.
	Value() string
	// SetValue assigns string provided by user to LldpChassisMacSubType
	SetValue(value string) LldpChassisMacSubType
	// HasValue checks if Value has been set in LldpChassisMacSubType
	HasValue() bool
}

type LldpChassisMacSubTypeChoiceEnum string

// Enum of Choice on LldpChassisMacSubType
var LldpChassisMacSubTypeChoice = struct {
	AUTO  LldpChassisMacSubTypeChoiceEnum
	VALUE LldpChassisMacSubTypeChoiceEnum
}{
	AUTO:  LldpChassisMacSubTypeChoiceEnum("auto"),
	VALUE: LldpChassisMacSubTypeChoiceEnum("value"),
}

func (obj *lldpChassisMacSubType) Choice() LldpChassisMacSubTypeChoiceEnum {
	return LldpChassisMacSubTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// In auto mode the system generated value is set for this property, while if the choice is selected as value,  a user configured value will be used for this property.
// Choice returns a string
func (obj *lldpChassisMacSubType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *lldpChassisMacSubType) setChoice(value LldpChassisMacSubTypeChoiceEnum) LldpChassisMacSubType {
	intValue, ok := otg.LldpChassisMacSubType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LldpChassisMacSubTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LldpChassisMacSubType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.Auto = nil
	return obj
}

// The OTG implementation must provide a system generated value for this property.
// Auto returns a string
func (obj *lldpChassisMacSubType) Auto() string {

	if obj.obj.Auto == nil {
		obj.setChoice(LldpChassisMacSubTypeChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation must provide a system generated value for this property.
// Auto returns a string
func (obj *lldpChassisMacSubType) HasAuto() bool {
	return obj.obj.Auto != nil
}

// User must specify a value if mode is not auto.
// Value returns a string
func (obj *lldpChassisMacSubType) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(LldpChassisMacSubTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// User must specify a value if mode is not auto.
// Value returns a string
func (obj *lldpChassisMacSubType) HasValue() bool {
	return obj.obj.Value != nil
}

// User must specify a value if mode is not auto.
// SetValue sets the string value in the LldpChassisMacSubType object
func (obj *lldpChassisMacSubType) SetValue(value string) LldpChassisMacSubType {
	obj.setChoice(LldpChassisMacSubTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *lldpChassisMacSubType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Auto != nil {

		err := obj.validateMac(obj.Auto())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LldpChassisMacSubType.Auto"))
		}

	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LldpChassisMacSubType.Value"))
		}

	}

}

func (obj *lldpChassisMacSubType) setDefault() {
	var choices_set int = 0
	var choice LldpChassisMacSubTypeChoiceEnum

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = LldpChassisMacSubTypeChoice.AUTO
	}

	if obj.obj.Value != nil {
		choices_set += 1
		choice = LldpChassisMacSubTypeChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(LldpChassisMacSubTypeChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in LldpChassisMacSubType")
			}
		} else {
			intVal := otg.LldpChassisMacSubType_Choice_Enum_value[string(choice)]
			enumValue := otg.LldpChassisMacSubType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
