package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVariableBindingValueIpAddressValue *****
type patternFlowSnmpv2CVariableBindingValueIpAddressValue struct {
	validation
	obj             *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	marshaller      marshalPatternFlowSnmpv2CVariableBindingValueIpAddressValue
	unMarshaller    unMarshalPatternFlowSnmpv2CVariableBindingValueIpAddressValue
	incrementHolder PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	decrementHolder PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
}

func NewPatternFlowSnmpv2CVariableBindingValueIpAddressValue() PatternFlowSnmpv2CVariableBindingValueIpAddressValue {
	obj := patternFlowSnmpv2CVariableBindingValueIpAddressValue{obj: &otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) msg() *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) setMsg(msg *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue) PatternFlowSnmpv2CVariableBindingValueIpAddressValue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVariableBindingValueIpAddressValue struct {
	obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue
}

type marshalPatternFlowSnmpv2CVariableBindingValueIpAddressValue interface {
	// ToProto marshals PatternFlowSnmpv2CVariableBindingValueIpAddressValue to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue, error)
	// ToPbText marshals PatternFlowSnmpv2CVariableBindingValueIpAddressValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVariableBindingValueIpAddressValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVariableBindingValueIpAddressValue to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowSnmpv2CVariableBindingValueIpAddressValue struct {
	obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue
}

type unMarshalPatternFlowSnmpv2CVariableBindingValueIpAddressValue interface {
	// FromProto unmarshals PatternFlowSnmpv2CVariableBindingValueIpAddressValue from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue) (PatternFlowSnmpv2CVariableBindingValueIpAddressValue, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVariableBindingValueIpAddressValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVariableBindingValueIpAddressValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVariableBindingValueIpAddressValue from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) Marshal() marshalPatternFlowSnmpv2CVariableBindingValueIpAddressValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVariableBindingValueIpAddressValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueIpAddressValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVariableBindingValueIpAddressValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIpAddressValue) ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIpAddressValue) FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue) (PatternFlowSnmpv2CVariableBindingValueIpAddressValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIpAddressValue) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIpAddressValue) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIpAddressValue) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIpAddressValue) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIpAddressValue) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIpAddressValue) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) Clone() (PatternFlowSnmpv2CVariableBindingValueIpAddressValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVariableBindingValueIpAddressValue()
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

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowSnmpv2CVariableBindingValueIpAddressValue is iPv4 address returned for the requested OID.
type PatternFlowSnmpv2CVariableBindingValueIpAddressValue interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVariableBindingValueIpAddressValue to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	// setMsg unmarshals PatternFlowSnmpv2CVariableBindingValueIpAddressValue from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue) PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVariableBindingValueIpAddressValue
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueIpAddressValue
	// validate validates PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVariableBindingValueIpAddressValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum, set in PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	Choice() PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum
	// setChoice assigns PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum provided by user to PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	setChoice(value PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum) PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	// HasChoice checks if Choice has been set in PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	HasChoice() bool
	// Value returns string, set in PatternFlowSnmpv2CVariableBindingValueIpAddressValue.
	Value() string
	// SetValue assigns string provided by user to PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	SetValue(value string) PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	// HasValue checks if Value has been set in PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	HasValue() bool
	// Values returns []string, set in PatternFlowSnmpv2CVariableBindingValueIpAddressValue.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	SetValues(value []string) PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	// Increment returns PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter, set in PatternFlowSnmpv2CVariableBindingValueIpAddressValue.
	Increment() PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	// SetIncrement assigns PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter provided by user to PatternFlowSnmpv2CVariableBindingValueIpAddressValue.
	SetIncrement(value PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	// HasIncrement checks if Increment has been set in PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	HasIncrement() bool
	// Decrement returns PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter, set in PatternFlowSnmpv2CVariableBindingValueIpAddressValue.
	Decrement() PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	// SetDecrement assigns PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter provided by user to PatternFlowSnmpv2CVariableBindingValueIpAddressValue.
	SetDecrement(value PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	// HasDecrement checks if Decrement has been set in PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	HasDecrement() bool
	setNil()
}

type PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum string

// Enum of Choice on PatternFlowSnmpv2CVariableBindingValueIpAddressValue
var PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice = struct {
	VALUE     PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum
	VALUES    PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum
	INCREMENT PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum
	DECREMENT PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum
}{
	VALUE:     PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum("value"),
	VALUES:    PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum("values"),
	INCREMENT: PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum("increment"),
	DECREMENT: PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum("decrement"),
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) Choice() PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum {
	return PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) setChoice(value PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum) PatternFlowSnmpv2CVariableBindingValueIpAddressValue {
	intValue, ok := otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter().msg()
	}

	if value == PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowSnmpv2CVariableBindingValueIpAddressValue object
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) SetValue(value string) PatternFlowSnmpv2CVariableBindingValueIpAddressValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowSnmpv2CVariableBindingValueIpAddressValue object
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) SetValues(value []string) PatternFlowSnmpv2CVariableBindingValueIpAddressValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) Increment() PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter value in the PatternFlowSnmpv2CVariableBindingValueIpAddressValue object
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) SetIncrement(value PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) PatternFlowSnmpv2CVariableBindingValueIpAddressValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) Decrement() PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter value in the PatternFlowSnmpv2CVariableBindingValueIpAddressValue object
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) SetDecrement(value PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) PatternFlowSnmpv2CVariableBindingValueIpAddressValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowSnmpv2CVariableBindingValueIpAddressValue.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowSnmpv2CVariableBindingValueIpAddressValue.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValue) setDefault() {
	var choices_set int = 0
	var choice PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIpAddressValueChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowSnmpv2CVariableBindingValueIpAddressValue")
			}
		} else {
			intVal := otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValue_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
