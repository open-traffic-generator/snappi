package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpPortInterfaceNameSubType *****
type lldpPortInterfaceNameSubType struct {
	validation
	obj          *otg.LldpPortInterfaceNameSubType
	marshaller   marshalLldpPortInterfaceNameSubType
	unMarshaller unMarshalLldpPortInterfaceNameSubType
}

func NewLldpPortInterfaceNameSubType() LldpPortInterfaceNameSubType {
	obj := lldpPortInterfaceNameSubType{obj: &otg.LldpPortInterfaceNameSubType{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpPortInterfaceNameSubType) msg() *otg.LldpPortInterfaceNameSubType {
	return obj.obj
}

func (obj *lldpPortInterfaceNameSubType) setMsg(msg *otg.LldpPortInterfaceNameSubType) LldpPortInterfaceNameSubType {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpPortInterfaceNameSubType struct {
	obj *lldpPortInterfaceNameSubType
}

type marshalLldpPortInterfaceNameSubType interface {
	// ToProto marshals LldpPortInterfaceNameSubType to protobuf object *otg.LldpPortInterfaceNameSubType
	ToProto() (*otg.LldpPortInterfaceNameSubType, error)
	// ToPbText marshals LldpPortInterfaceNameSubType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpPortInterfaceNameSubType to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpPortInterfaceNameSubType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LldpPortInterfaceNameSubType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallldpPortInterfaceNameSubType struct {
	obj *lldpPortInterfaceNameSubType
}

type unMarshalLldpPortInterfaceNameSubType interface {
	// FromProto unmarshals LldpPortInterfaceNameSubType from protobuf object *otg.LldpPortInterfaceNameSubType
	FromProto(msg *otg.LldpPortInterfaceNameSubType) (LldpPortInterfaceNameSubType, error)
	// FromPbText unmarshals LldpPortInterfaceNameSubType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpPortInterfaceNameSubType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpPortInterfaceNameSubType from JSON text
	FromJson(value string) error
}

func (obj *lldpPortInterfaceNameSubType) Marshal() marshalLldpPortInterfaceNameSubType {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpPortInterfaceNameSubType{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpPortInterfaceNameSubType) Unmarshal() unMarshalLldpPortInterfaceNameSubType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpPortInterfaceNameSubType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpPortInterfaceNameSubType) ToProto() (*otg.LldpPortInterfaceNameSubType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpPortInterfaceNameSubType) FromProto(msg *otg.LldpPortInterfaceNameSubType) (LldpPortInterfaceNameSubType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpPortInterfaceNameSubType) ToPbText() (string, error) {
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

func (m *unMarshallldpPortInterfaceNameSubType) FromPbText(value string) error {
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

func (m *marshallldpPortInterfaceNameSubType) ToYaml() (string, error) {
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

func (m *unMarshallldpPortInterfaceNameSubType) FromYaml(value string) error {
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

func (m *marshallldpPortInterfaceNameSubType) ToJsonRaw() (string, error) {
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

func (m *marshallldpPortInterfaceNameSubType) ToJson() (string, error) {
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

func (m *unMarshallldpPortInterfaceNameSubType) FromJson(value string) error {
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

func (obj *lldpPortInterfaceNameSubType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpPortInterfaceNameSubType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpPortInterfaceNameSubType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpPortInterfaceNameSubType) Clone() (LldpPortInterfaceNameSubType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpPortInterfaceNameSubType()
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

// LldpPortInterfaceNameSubType is the interface name configured in the Port ID TLV.
type LldpPortInterfaceNameSubType interface {
	Validation
	// msg marshals LldpPortInterfaceNameSubType to protobuf object *otg.LldpPortInterfaceNameSubType
	// and doesn't set defaults
	msg() *otg.LldpPortInterfaceNameSubType
	// setMsg unmarshals LldpPortInterfaceNameSubType from protobuf object *otg.LldpPortInterfaceNameSubType
	// and doesn't set defaults
	setMsg(*otg.LldpPortInterfaceNameSubType) LldpPortInterfaceNameSubType
	// provides marshal interface
	Marshal() marshalLldpPortInterfaceNameSubType
	// provides unmarshal interface
	Unmarshal() unMarshalLldpPortInterfaceNameSubType
	// validate validates LldpPortInterfaceNameSubType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpPortInterfaceNameSubType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LldpPortInterfaceNameSubTypeChoiceEnum, set in LldpPortInterfaceNameSubType
	Choice() LldpPortInterfaceNameSubTypeChoiceEnum
	// setChoice assigns LldpPortInterfaceNameSubTypeChoiceEnum provided by user to LldpPortInterfaceNameSubType
	setChoice(value LldpPortInterfaceNameSubTypeChoiceEnum) LldpPortInterfaceNameSubType
	// HasChoice checks if Choice has been set in LldpPortInterfaceNameSubType
	HasChoice() bool
	// Auto returns string, set in LldpPortInterfaceNameSubType.
	Auto() string
	// HasAuto checks if Auto has been set in LldpPortInterfaceNameSubType
	HasAuto() bool
	// Value returns string, set in LldpPortInterfaceNameSubType.
	Value() string
	// SetValue assigns string provided by user to LldpPortInterfaceNameSubType
	SetValue(value string) LldpPortInterfaceNameSubType
	// HasValue checks if Value has been set in LldpPortInterfaceNameSubType
	HasValue() bool
}

type LldpPortInterfaceNameSubTypeChoiceEnum string

// Enum of Choice on LldpPortInterfaceNameSubType
var LldpPortInterfaceNameSubTypeChoice = struct {
	AUTO  LldpPortInterfaceNameSubTypeChoiceEnum
	VALUE LldpPortInterfaceNameSubTypeChoiceEnum
}{
	AUTO:  LldpPortInterfaceNameSubTypeChoiceEnum("auto"),
	VALUE: LldpPortInterfaceNameSubTypeChoiceEnum("value"),
}

func (obj *lldpPortInterfaceNameSubType) Choice() LldpPortInterfaceNameSubTypeChoiceEnum {
	return LldpPortInterfaceNameSubTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// In auto mode the system generated value is set for this property, while if the choice is selected as value,  a user configured value will be used for this property.
// Choice returns a string
func (obj *lldpPortInterfaceNameSubType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *lldpPortInterfaceNameSubType) setChoice(value LldpPortInterfaceNameSubTypeChoiceEnum) LldpPortInterfaceNameSubType {
	intValue, ok := otg.LldpPortInterfaceNameSubType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LldpPortInterfaceNameSubTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LldpPortInterfaceNameSubType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.Auto = nil
	return obj
}

// The OTG implementation must provide a system generated value for this property.
// Auto returns a string
func (obj *lldpPortInterfaceNameSubType) Auto() string {

	if obj.obj.Auto == nil {
		obj.setChoice(LldpPortInterfaceNameSubTypeChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation must provide a system generated value for this property.
// Auto returns a string
func (obj *lldpPortInterfaceNameSubType) HasAuto() bool {
	return obj.obj.Auto != nil
}

// User must specify a value if mode is not auto.
// Value returns a string
func (obj *lldpPortInterfaceNameSubType) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(LldpPortInterfaceNameSubTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// User must specify a value if mode is not auto.
// Value returns a string
func (obj *lldpPortInterfaceNameSubType) HasValue() bool {
	return obj.obj.Value != nil
}

// User must specify a value if mode is not auto.
// SetValue sets the string value in the LldpPortInterfaceNameSubType object
func (obj *lldpPortInterfaceNameSubType) SetValue(value string) LldpPortInterfaceNameSubType {
	obj.setChoice(LldpPortInterfaceNameSubTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *lldpPortInterfaceNameSubType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lldpPortInterfaceNameSubType) setDefault() {
	var choices_set int = 0
	var choice LldpPortInterfaceNameSubTypeChoiceEnum

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = LldpPortInterfaceNameSubTypeChoice.AUTO
	}

	if obj.obj.Value != nil {
		choices_set += 1
		choice = LldpPortInterfaceNameSubTypeChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(LldpPortInterfaceNameSubTypeChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in LldpPortInterfaceNameSubType")
			}
		} else {
			intVal := otg.LldpPortInterfaceNameSubType_Choice_Enum_value[string(choice)]
			enumValue := otg.LldpPortInterfaceNameSubType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
