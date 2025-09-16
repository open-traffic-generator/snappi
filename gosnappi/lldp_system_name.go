package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpSystemName *****
type lldpSystemName struct {
	validation
	obj          *otg.LldpSystemName
	marshaller   marshalLldpSystemName
	unMarshaller unMarshalLldpSystemName
}

func NewLldpSystemName() LldpSystemName {
	obj := lldpSystemName{obj: &otg.LldpSystemName{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpSystemName) msg() *otg.LldpSystemName {
	return obj.obj
}

func (obj *lldpSystemName) setMsg(msg *otg.LldpSystemName) LldpSystemName {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpSystemName struct {
	obj *lldpSystemName
}

type marshalLldpSystemName interface {
	// ToProto marshals LldpSystemName to protobuf object *otg.LldpSystemName
	ToProto() (*otg.LldpSystemName, error)
	// ToPbText marshals LldpSystemName to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpSystemName to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpSystemName to JSON text
	ToJson() (string, error)
}

type unMarshallldpSystemName struct {
	obj *lldpSystemName
}

type unMarshalLldpSystemName interface {
	// FromProto unmarshals LldpSystemName from protobuf object *otg.LldpSystemName
	FromProto(msg *otg.LldpSystemName) (LldpSystemName, error)
	// FromPbText unmarshals LldpSystemName from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpSystemName from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpSystemName from JSON text
	FromJson(value string) error
}

func (obj *lldpSystemName) Marshal() marshalLldpSystemName {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpSystemName{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpSystemName) Unmarshal() unMarshalLldpSystemName {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpSystemName{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpSystemName) ToProto() (*otg.LldpSystemName, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpSystemName) FromProto(msg *otg.LldpSystemName) (LldpSystemName, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpSystemName) ToPbText() (string, error) {
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

func (m *unMarshallldpSystemName) FromPbText(value string) error {
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

func (m *marshallldpSystemName) ToYaml() (string, error) {
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

func (m *unMarshallldpSystemName) FromYaml(value string) error {
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

func (m *marshallldpSystemName) ToJson() (string, error) {
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

func (m *unMarshallldpSystemName) FromJson(value string) error {
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

func (obj *lldpSystemName) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpSystemName) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpSystemName) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpSystemName) Clone() (LldpSystemName, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpSystemName()
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

// LldpSystemName is the system Name configured in the System Name TLV.
type LldpSystemName interface {
	Validation
	// msg marshals LldpSystemName to protobuf object *otg.LldpSystemName
	// and doesn't set defaults
	msg() *otg.LldpSystemName
	// setMsg unmarshals LldpSystemName from protobuf object *otg.LldpSystemName
	// and doesn't set defaults
	setMsg(*otg.LldpSystemName) LldpSystemName
	// provides marshal interface
	Marshal() marshalLldpSystemName
	// provides unmarshal interface
	Unmarshal() unMarshalLldpSystemName
	// validate validates LldpSystemName
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpSystemName, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LldpSystemNameChoiceEnum, set in LldpSystemName
	Choice() LldpSystemNameChoiceEnum
	// setChoice assigns LldpSystemNameChoiceEnum provided by user to LldpSystemName
	setChoice(value LldpSystemNameChoiceEnum) LldpSystemName
	// HasChoice checks if Choice has been set in LldpSystemName
	HasChoice() bool
	// Auto returns string, set in LldpSystemName.
	Auto() string
	// HasAuto checks if Auto has been set in LldpSystemName
	HasAuto() bool
	// Value returns string, set in LldpSystemName.
	Value() string
	// SetValue assigns string provided by user to LldpSystemName
	SetValue(value string) LldpSystemName
	// HasValue checks if Value has been set in LldpSystemName
	HasValue() bool
}

type LldpSystemNameChoiceEnum string

// Enum of Choice on LldpSystemName
var LldpSystemNameChoice = struct {
	AUTO  LldpSystemNameChoiceEnum
	VALUE LldpSystemNameChoiceEnum
}{
	AUTO:  LldpSystemNameChoiceEnum("auto"),
	VALUE: LldpSystemNameChoiceEnum("value"),
}

func (obj *lldpSystemName) Choice() LldpSystemNameChoiceEnum {
	return LldpSystemNameChoiceEnum(obj.obj.Choice.Enum().String())
}

// In auto mode the system generated value is set for this property, while if the choice is selected as value,  a user configured value will be used for this property.
// Choice returns a string
func (obj *lldpSystemName) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *lldpSystemName) setChoice(value LldpSystemNameChoiceEnum) LldpSystemName {
	intValue, ok := otg.LldpSystemName_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LldpSystemNameChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LldpSystemName_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.Auto = nil
	return obj
}

// The OTG implementation must provide a system generated value for this property.
// Auto returns a string
func (obj *lldpSystemName) Auto() string {

	if obj.obj.Auto == nil {
		obj.setChoice(LldpSystemNameChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation must provide a system generated value for this property.
// Auto returns a string
func (obj *lldpSystemName) HasAuto() bool {
	return obj.obj.Auto != nil
}

// User must specify a value if mode is not auto.
// Value returns a string
func (obj *lldpSystemName) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(LldpSystemNameChoice.VALUE)
	}

	return *obj.obj.Value

}

// User must specify a value if mode is not auto.
// Value returns a string
func (obj *lldpSystemName) HasValue() bool {
	return obj.obj.Value != nil
}

// User must specify a value if mode is not auto.
// SetValue sets the string value in the LldpSystemName object
func (obj *lldpSystemName) SetValue(value string) LldpSystemName {
	obj.setChoice(LldpSystemNameChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *lldpSystemName) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lldpSystemName) setDefault() {
	var choices_set int = 0
	var choice LldpSystemNameChoiceEnum

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = LldpSystemNameChoice.AUTO
	}

	if obj.obj.Value != nil {
		choices_set += 1
		choice = LldpSystemNameChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(LldpSystemNameChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in LldpSystemName")
			}
		} else {
			intVal := otg.LldpSystemName_Choice_Enum_value[string(choice)]
			enumValue := otg.LldpSystemName_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
