package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVariableBindingValueCounterValue *****
type patternFlowSnmpv2CVariableBindingValueCounterValue struct {
	validation
	obj             *otg.PatternFlowSnmpv2CVariableBindingValueCounterValue
	marshaller      marshalPatternFlowSnmpv2CVariableBindingValueCounterValue
	unMarshaller    unMarshalPatternFlowSnmpv2CVariableBindingValueCounterValue
	incrementHolder PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	decrementHolder PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
}

func NewPatternFlowSnmpv2CVariableBindingValueCounterValue() PatternFlowSnmpv2CVariableBindingValueCounterValue {
	obj := patternFlowSnmpv2CVariableBindingValueCounterValue{obj: &otg.PatternFlowSnmpv2CVariableBindingValueCounterValue{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) msg() *otg.PatternFlowSnmpv2CVariableBindingValueCounterValue {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) setMsg(msg *otg.PatternFlowSnmpv2CVariableBindingValueCounterValue) PatternFlowSnmpv2CVariableBindingValueCounterValue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVariableBindingValueCounterValue struct {
	obj *patternFlowSnmpv2CVariableBindingValueCounterValue
}

type marshalPatternFlowSnmpv2CVariableBindingValueCounterValue interface {
	// ToProto marshals PatternFlowSnmpv2CVariableBindingValueCounterValue to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueCounterValue
	ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueCounterValue, error)
	// ToPbText marshals PatternFlowSnmpv2CVariableBindingValueCounterValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVariableBindingValueCounterValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVariableBindingValueCounterValue to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowSnmpv2CVariableBindingValueCounterValue struct {
	obj *patternFlowSnmpv2CVariableBindingValueCounterValue
}

type unMarshalPatternFlowSnmpv2CVariableBindingValueCounterValue interface {
	// FromProto unmarshals PatternFlowSnmpv2CVariableBindingValueCounterValue from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueCounterValue
	FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueCounterValue) (PatternFlowSnmpv2CVariableBindingValueCounterValue, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVariableBindingValueCounterValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVariableBindingValueCounterValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVariableBindingValueCounterValue from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) Marshal() marshalPatternFlowSnmpv2CVariableBindingValueCounterValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVariableBindingValueCounterValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueCounterValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVariableBindingValueCounterValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueCounterValue) ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueCounterValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueCounterValue) FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueCounterValue) (PatternFlowSnmpv2CVariableBindingValueCounterValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueCounterValue) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueCounterValue) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueCounterValue) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueCounterValue) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueCounterValue) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueCounterValue) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) Clone() (PatternFlowSnmpv2CVariableBindingValueCounterValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVariableBindingValueCounterValue()
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

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowSnmpv2CVariableBindingValueCounterValue is counter returned for the requested OID.
type PatternFlowSnmpv2CVariableBindingValueCounterValue interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVariableBindingValueCounterValue to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueCounterValue
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVariableBindingValueCounterValue
	// setMsg unmarshals PatternFlowSnmpv2CVariableBindingValueCounterValue from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueCounterValue
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVariableBindingValueCounterValue) PatternFlowSnmpv2CVariableBindingValueCounterValue
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVariableBindingValueCounterValue
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueCounterValue
	// validate validates PatternFlowSnmpv2CVariableBindingValueCounterValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVariableBindingValueCounterValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum, set in PatternFlowSnmpv2CVariableBindingValueCounterValue
	Choice() PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum
	// setChoice assigns PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum provided by user to PatternFlowSnmpv2CVariableBindingValueCounterValue
	setChoice(value PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum) PatternFlowSnmpv2CVariableBindingValueCounterValue
	// HasChoice checks if Choice has been set in PatternFlowSnmpv2CVariableBindingValueCounterValue
	HasChoice() bool
	// Value returns uint32, set in PatternFlowSnmpv2CVariableBindingValueCounterValue.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueCounterValue
	SetValue(value uint32) PatternFlowSnmpv2CVariableBindingValueCounterValue
	// HasValue checks if Value has been set in PatternFlowSnmpv2CVariableBindingValueCounterValue
	HasValue() bool
	// Values returns []uint32, set in PatternFlowSnmpv2CVariableBindingValueCounterValue.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueCounterValue
	SetValues(value []uint32) PatternFlowSnmpv2CVariableBindingValueCounterValue
	// Increment returns PatternFlowSnmpv2CVariableBindingValueCounterValueCounter, set in PatternFlowSnmpv2CVariableBindingValueCounterValue.
	Increment() PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	// SetIncrement assigns PatternFlowSnmpv2CVariableBindingValueCounterValueCounter provided by user to PatternFlowSnmpv2CVariableBindingValueCounterValue.
	SetIncrement(value PatternFlowSnmpv2CVariableBindingValueCounterValueCounter) PatternFlowSnmpv2CVariableBindingValueCounterValue
	// HasIncrement checks if Increment has been set in PatternFlowSnmpv2CVariableBindingValueCounterValue
	HasIncrement() bool
	// Decrement returns PatternFlowSnmpv2CVariableBindingValueCounterValueCounter, set in PatternFlowSnmpv2CVariableBindingValueCounterValue.
	Decrement() PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	// SetDecrement assigns PatternFlowSnmpv2CVariableBindingValueCounterValueCounter provided by user to PatternFlowSnmpv2CVariableBindingValueCounterValue.
	SetDecrement(value PatternFlowSnmpv2CVariableBindingValueCounterValueCounter) PatternFlowSnmpv2CVariableBindingValueCounterValue
	// HasDecrement checks if Decrement has been set in PatternFlowSnmpv2CVariableBindingValueCounterValue
	HasDecrement() bool
	setNil()
}

type PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum string

// Enum of Choice on PatternFlowSnmpv2CVariableBindingValueCounterValue
var PatternFlowSnmpv2CVariableBindingValueCounterValueChoice = struct {
	VALUE     PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum
	VALUES    PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum
	INCREMENT PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum
	DECREMENT PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum
}{
	VALUE:     PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum("value"),
	VALUES:    PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum("values"),
	INCREMENT: PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum("increment"),
	DECREMENT: PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum("decrement"),
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) Choice() PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum {
	return PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) setChoice(value PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum) PatternFlowSnmpv2CVariableBindingValueCounterValue {
	intValue, ok := otg.PatternFlowSnmpv2CVariableBindingValueCounterValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowSnmpv2CVariableBindingValueCounterValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowSnmpv2CVariableBindingValueCounterValueCounter().msg()
	}

	if value == PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowSnmpv2CVariableBindingValueCounterValueCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowSnmpv2CVariableBindingValueCounterValue object
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) SetValue(value uint32) PatternFlowSnmpv2CVariableBindingValueCounterValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowSnmpv2CVariableBindingValueCounterValue object
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) SetValues(value []uint32) PatternFlowSnmpv2CVariableBindingValueCounterValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) Increment() PatternFlowSnmpv2CVariableBindingValueCounterValueCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowSnmpv2CVariableBindingValueCounterValueCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowSnmpv2CVariableBindingValueCounterValueCounter value in the PatternFlowSnmpv2CVariableBindingValueCounterValue object
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) SetIncrement(value PatternFlowSnmpv2CVariableBindingValueCounterValueCounter) PatternFlowSnmpv2CVariableBindingValueCounterValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) Decrement() PatternFlowSnmpv2CVariableBindingValueCounterValueCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowSnmpv2CVariableBindingValueCounterValueCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowSnmpv2CVariableBindingValueCounterValueCounter value in the PatternFlowSnmpv2CVariableBindingValueCounterValue object
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) SetDecrement(value PatternFlowSnmpv2CVariableBindingValueCounterValueCounter) PatternFlowSnmpv2CVariableBindingValueCounterValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValue) setDefault() {
	var choices_set int = 0
	var choice PatternFlowSnmpv2CVariableBindingValueCounterValueChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowSnmpv2CVariableBindingValueCounterValueChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowSnmpv2CVariableBindingValueCounterValue")
			}
		} else {
			intVal := otg.PatternFlowSnmpv2CVariableBindingValueCounterValue_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowSnmpv2CVariableBindingValueCounterValue_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
