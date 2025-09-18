package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateAggregation *****
type patternFlowLacpduActorStateAggregation struct {
	validation
	obj             *otg.PatternFlowLacpduActorStateAggregation
	marshaller      marshalPatternFlowLacpduActorStateAggregation
	unMarshaller    unMarshalPatternFlowLacpduActorStateAggregation
	incrementHolder PatternFlowLacpduActorStateAggregationCounter
	decrementHolder PatternFlowLacpduActorStateAggregationCounter
}

func NewPatternFlowLacpduActorStateAggregation() PatternFlowLacpduActorStateAggregation {
	obj := patternFlowLacpduActorStateAggregation{obj: &otg.PatternFlowLacpduActorStateAggregation{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateAggregation) msg() *otg.PatternFlowLacpduActorStateAggregation {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateAggregation) setMsg(msg *otg.PatternFlowLacpduActorStateAggregation) PatternFlowLacpduActorStateAggregation {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateAggregation struct {
	obj *patternFlowLacpduActorStateAggregation
}

type marshalPatternFlowLacpduActorStateAggregation interface {
	// ToProto marshals PatternFlowLacpduActorStateAggregation to protobuf object *otg.PatternFlowLacpduActorStateAggregation
	ToProto() (*otg.PatternFlowLacpduActorStateAggregation, error)
	// ToPbText marshals PatternFlowLacpduActorStateAggregation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateAggregation to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateAggregation to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateAggregation struct {
	obj *patternFlowLacpduActorStateAggregation
}

type unMarshalPatternFlowLacpduActorStateAggregation interface {
	// FromProto unmarshals PatternFlowLacpduActorStateAggregation from protobuf object *otg.PatternFlowLacpduActorStateAggregation
	FromProto(msg *otg.PatternFlowLacpduActorStateAggregation) (PatternFlowLacpduActorStateAggregation, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateAggregation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateAggregation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateAggregation from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateAggregation) Marshal() marshalPatternFlowLacpduActorStateAggregation {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateAggregation{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateAggregation) Unmarshal() unMarshalPatternFlowLacpduActorStateAggregation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateAggregation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateAggregation) ToProto() (*otg.PatternFlowLacpduActorStateAggregation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateAggregation) FromProto(msg *otg.PatternFlowLacpduActorStateAggregation) (PatternFlowLacpduActorStateAggregation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateAggregation) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateAggregation) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateAggregation) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateAggregation) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateAggregation) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateAggregation) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateAggregation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateAggregation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateAggregation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateAggregation) Clone() (PatternFlowLacpduActorStateAggregation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateAggregation()
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

func (obj *patternFlowLacpduActorStateAggregation) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorStateAggregation is aggregation (1=Aggregatable, 0=Individual)
type PatternFlowLacpduActorStateAggregation interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateAggregation to protobuf object *otg.PatternFlowLacpduActorStateAggregation
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateAggregation
	// setMsg unmarshals PatternFlowLacpduActorStateAggregation from protobuf object *otg.PatternFlowLacpduActorStateAggregation
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateAggregation) PatternFlowLacpduActorStateAggregation
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateAggregation
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateAggregation
	// validate validates PatternFlowLacpduActorStateAggregation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateAggregation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorStateAggregationChoiceEnum, set in PatternFlowLacpduActorStateAggregation
	Choice() PatternFlowLacpduActorStateAggregationChoiceEnum
	// setChoice assigns PatternFlowLacpduActorStateAggregationChoiceEnum provided by user to PatternFlowLacpduActorStateAggregation
	setChoice(value PatternFlowLacpduActorStateAggregationChoiceEnum) PatternFlowLacpduActorStateAggregation
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorStateAggregation
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorStateAggregation.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorStateAggregation
	SetValue(value uint32) PatternFlowLacpduActorStateAggregation
	// HasValue checks if Value has been set in PatternFlowLacpduActorStateAggregation
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorStateAggregation.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorStateAggregation
	SetValues(value []uint32) PatternFlowLacpduActorStateAggregation
	// Increment returns PatternFlowLacpduActorStateAggregationCounter, set in PatternFlowLacpduActorStateAggregation.
	// PatternFlowLacpduActorStateAggregationCounter is integer counter pattern
	Increment() PatternFlowLacpduActorStateAggregationCounter
	// SetIncrement assigns PatternFlowLacpduActorStateAggregationCounter provided by user to PatternFlowLacpduActorStateAggregation.
	// PatternFlowLacpduActorStateAggregationCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorStateAggregationCounter) PatternFlowLacpduActorStateAggregation
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorStateAggregation
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorStateAggregationCounter, set in PatternFlowLacpduActorStateAggregation.
	// PatternFlowLacpduActorStateAggregationCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorStateAggregationCounter
	// SetDecrement assigns PatternFlowLacpduActorStateAggregationCounter provided by user to PatternFlowLacpduActorStateAggregation.
	// PatternFlowLacpduActorStateAggregationCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorStateAggregationCounter) PatternFlowLacpduActorStateAggregation
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorStateAggregation
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorStateAggregationChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorStateAggregation
var PatternFlowLacpduActorStateAggregationChoice = struct {
	VALUE     PatternFlowLacpduActorStateAggregationChoiceEnum
	VALUES    PatternFlowLacpduActorStateAggregationChoiceEnum
	INCREMENT PatternFlowLacpduActorStateAggregationChoiceEnum
	DECREMENT PatternFlowLacpduActorStateAggregationChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorStateAggregationChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorStateAggregationChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorStateAggregationChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorStateAggregationChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorStateAggregation) Choice() PatternFlowLacpduActorStateAggregationChoiceEnum {
	return PatternFlowLacpduActorStateAggregationChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorStateAggregation) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorStateAggregation) setChoice(value PatternFlowLacpduActorStateAggregationChoiceEnum) PatternFlowLacpduActorStateAggregation {
	intValue, ok := otg.PatternFlowLacpduActorStateAggregation_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorStateAggregationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorStateAggregation_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorStateAggregationChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorStateAggregationChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorStateAggregationChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorStateAggregationCounter().msg()
	}

	if value == PatternFlowLacpduActorStateAggregationChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorStateAggregationCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateAggregation) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorStateAggregationChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateAggregation) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorStateAggregation object
func (obj *patternFlowLacpduActorStateAggregation) SetValue(value uint32) PatternFlowLacpduActorStateAggregation {
	obj.setChoice(PatternFlowLacpduActorStateAggregationChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorStateAggregation) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorStateAggregation object
func (obj *patternFlowLacpduActorStateAggregation) SetValues(value []uint32) PatternFlowLacpduActorStateAggregation {
	obj.setChoice(PatternFlowLacpduActorStateAggregationChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateAggregationCounter
func (obj *patternFlowLacpduActorStateAggregation) Increment() PatternFlowLacpduActorStateAggregationCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorStateAggregationChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorStateAggregationCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateAggregationCounter
func (obj *patternFlowLacpduActorStateAggregation) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorStateAggregationCounter value in the PatternFlowLacpduActorStateAggregation object
func (obj *patternFlowLacpduActorStateAggregation) SetIncrement(value PatternFlowLacpduActorStateAggregationCounter) PatternFlowLacpduActorStateAggregation {
	obj.setChoice(PatternFlowLacpduActorStateAggregationChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateAggregationCounter
func (obj *patternFlowLacpduActorStateAggregation) Decrement() PatternFlowLacpduActorStateAggregationCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorStateAggregationChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorStateAggregationCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateAggregationCounter
func (obj *patternFlowLacpduActorStateAggregation) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorStateAggregationCounter value in the PatternFlowLacpduActorStateAggregation object
func (obj *patternFlowLacpduActorStateAggregation) SetDecrement(value PatternFlowLacpduActorStateAggregationCounter) PatternFlowLacpduActorStateAggregation {
	obj.setChoice(PatternFlowLacpduActorStateAggregationChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorStateAggregation) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateAggregation.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorStateAggregation.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowLacpduActorStateAggregation) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorStateAggregationChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateAggregationChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorStateAggregationChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateAggregationChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateAggregationChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorStateAggregationChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorStateAggregation")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorStateAggregation_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorStateAggregation_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
