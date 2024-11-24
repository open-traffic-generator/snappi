package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVariableBindingValueIntegerValue *****
type patternFlowSnmpv2CVariableBindingValueIntegerValue struct {
	validation
	obj             *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue
	marshaller      marshalPatternFlowSnmpv2CVariableBindingValueIntegerValue
	unMarshaller    unMarshalPatternFlowSnmpv2CVariableBindingValueIntegerValue
	incrementHolder PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	decrementHolder PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
}

func NewPatternFlowSnmpv2CVariableBindingValueIntegerValue() PatternFlowSnmpv2CVariableBindingValueIntegerValue {
	obj := patternFlowSnmpv2CVariableBindingValueIntegerValue{obj: &otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) msg() *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) setMsg(msg *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue) PatternFlowSnmpv2CVariableBindingValueIntegerValue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVariableBindingValueIntegerValue struct {
	obj *patternFlowSnmpv2CVariableBindingValueIntegerValue
}

type marshalPatternFlowSnmpv2CVariableBindingValueIntegerValue interface {
	// ToProto marshals PatternFlowSnmpv2CVariableBindingValueIntegerValue to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue
	ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue, error)
	// ToPbText marshals PatternFlowSnmpv2CVariableBindingValueIntegerValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVariableBindingValueIntegerValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVariableBindingValueIntegerValue to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowSnmpv2CVariableBindingValueIntegerValue struct {
	obj *patternFlowSnmpv2CVariableBindingValueIntegerValue
}

type unMarshalPatternFlowSnmpv2CVariableBindingValueIntegerValue interface {
	// FromProto unmarshals PatternFlowSnmpv2CVariableBindingValueIntegerValue from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue
	FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue) (PatternFlowSnmpv2CVariableBindingValueIntegerValue, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVariableBindingValueIntegerValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVariableBindingValueIntegerValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVariableBindingValueIntegerValue from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) Marshal() marshalPatternFlowSnmpv2CVariableBindingValueIntegerValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVariableBindingValueIntegerValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueIntegerValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVariableBindingValueIntegerValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIntegerValue) ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIntegerValue) FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue) (PatternFlowSnmpv2CVariableBindingValueIntegerValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIntegerValue) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIntegerValue) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIntegerValue) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIntegerValue) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIntegerValue) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIntegerValue) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) Clone() (PatternFlowSnmpv2CVariableBindingValueIntegerValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVariableBindingValueIntegerValue()
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

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowSnmpv2CVariableBindingValueIntegerValue is integer value returned for the requested OID.
type PatternFlowSnmpv2CVariableBindingValueIntegerValue interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVariableBindingValueIntegerValue to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue
	// setMsg unmarshals PatternFlowSnmpv2CVariableBindingValueIntegerValue from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue) PatternFlowSnmpv2CVariableBindingValueIntegerValue
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVariableBindingValueIntegerValue
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueIntegerValue
	// validate validates PatternFlowSnmpv2CVariableBindingValueIntegerValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVariableBindingValueIntegerValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum, set in PatternFlowSnmpv2CVariableBindingValueIntegerValue
	Choice() PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum
	// setChoice assigns PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum provided by user to PatternFlowSnmpv2CVariableBindingValueIntegerValue
	setChoice(value PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum) PatternFlowSnmpv2CVariableBindingValueIntegerValue
	// HasChoice checks if Choice has been set in PatternFlowSnmpv2CVariableBindingValueIntegerValue
	HasChoice() bool
	// Value returns int32, set in PatternFlowSnmpv2CVariableBindingValueIntegerValue.
	Value() int32
	// SetValue assigns int32 provided by user to PatternFlowSnmpv2CVariableBindingValueIntegerValue
	SetValue(value int32) PatternFlowSnmpv2CVariableBindingValueIntegerValue
	// HasValue checks if Value has been set in PatternFlowSnmpv2CVariableBindingValueIntegerValue
	HasValue() bool
	// Values returns []int32, set in PatternFlowSnmpv2CVariableBindingValueIntegerValue.
	Values() []int32
	// SetValues assigns []int32 provided by user to PatternFlowSnmpv2CVariableBindingValueIntegerValue
	SetValues(value []int32) PatternFlowSnmpv2CVariableBindingValueIntegerValue
	// Increment returns PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter, set in PatternFlowSnmpv2CVariableBindingValueIntegerValue.
	Increment() PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	// SetIncrement assigns PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter provided by user to PatternFlowSnmpv2CVariableBindingValueIntegerValue.
	SetIncrement(value PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) PatternFlowSnmpv2CVariableBindingValueIntegerValue
	// HasIncrement checks if Increment has been set in PatternFlowSnmpv2CVariableBindingValueIntegerValue
	HasIncrement() bool
	// Decrement returns PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter, set in PatternFlowSnmpv2CVariableBindingValueIntegerValue.
	Decrement() PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	// SetDecrement assigns PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter provided by user to PatternFlowSnmpv2CVariableBindingValueIntegerValue.
	SetDecrement(value PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) PatternFlowSnmpv2CVariableBindingValueIntegerValue
	// HasDecrement checks if Decrement has been set in PatternFlowSnmpv2CVariableBindingValueIntegerValue
	HasDecrement() bool
	setNil()
}

type PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum string

// Enum of Choice on PatternFlowSnmpv2CVariableBindingValueIntegerValue
var PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice = struct {
	VALUE     PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum
	VALUES    PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum
	INCREMENT PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum
	DECREMENT PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum
}{
	VALUE:     PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum("value"),
	VALUES:    PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum("values"),
	INCREMENT: PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum("increment"),
	DECREMENT: PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum("decrement"),
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) Choice() PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum {
	return PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) setChoice(value PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum) PatternFlowSnmpv2CVariableBindingValueIntegerValue {
	intValue, ok := otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.VALUE {
		defaultValue := int32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.VALUES {
		defaultValue := []int32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowSnmpv2CVariableBindingValueIntegerValueCounter().msg()
	}

	if value == PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowSnmpv2CVariableBindingValueIntegerValueCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a int32
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) Value() int32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a int32
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the int32 value in the PatternFlowSnmpv2CVariableBindingValueIntegerValue object
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) SetValue(value int32) PatternFlowSnmpv2CVariableBindingValueIntegerValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []int32
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) Values() []int32 {
	if obj.obj.Values == nil {
		obj.SetValues([]int32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []int32 value in the PatternFlowSnmpv2CVariableBindingValueIntegerValue object
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) SetValues(value []int32) PatternFlowSnmpv2CVariableBindingValueIntegerValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]int32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) Increment() PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowSnmpv2CVariableBindingValueIntegerValueCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter value in the PatternFlowSnmpv2CVariableBindingValueIntegerValue object
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) SetIncrement(value PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) PatternFlowSnmpv2CVariableBindingValueIntegerValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) Decrement() PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowSnmpv2CVariableBindingValueIntegerValueCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter value in the PatternFlowSnmpv2CVariableBindingValueIntegerValue object
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) SetDecrement(value PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) PatternFlowSnmpv2CVariableBindingValueIntegerValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValue) setDefault() {
	var choices_set int = 0
	var choice PatternFlowSnmpv2CVariableBindingValueIntegerValueChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowSnmpv2CVariableBindingValueIntegerValueChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowSnmpv2CVariableBindingValueIntegerValue")
			}
		} else {
			intVal := otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowSnmpv2CVariableBindingValueIntegerValue_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
