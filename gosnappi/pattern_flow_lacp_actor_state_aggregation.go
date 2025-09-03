package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateAggregation *****
type patternFlowLacpActorStateAggregation struct {
	validation
	obj             *otg.PatternFlowLacpActorStateAggregation
	marshaller      marshalPatternFlowLacpActorStateAggregation
	unMarshaller    unMarshalPatternFlowLacpActorStateAggregation
	incrementHolder PatternFlowLacpActorStateAggregationCounter
	decrementHolder PatternFlowLacpActorStateAggregationCounter
}

func NewPatternFlowLacpActorStateAggregation() PatternFlowLacpActorStateAggregation {
	obj := patternFlowLacpActorStateAggregation{obj: &otg.PatternFlowLacpActorStateAggregation{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateAggregation) msg() *otg.PatternFlowLacpActorStateAggregation {
	return obj.obj
}

func (obj *patternFlowLacpActorStateAggregation) setMsg(msg *otg.PatternFlowLacpActorStateAggregation) PatternFlowLacpActorStateAggregation {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateAggregation struct {
	obj *patternFlowLacpActorStateAggregation
}

type marshalPatternFlowLacpActorStateAggregation interface {
	// ToProto marshals PatternFlowLacpActorStateAggregation to protobuf object *otg.PatternFlowLacpActorStateAggregation
	ToProto() (*otg.PatternFlowLacpActorStateAggregation, error)
	// ToPbText marshals PatternFlowLacpActorStateAggregation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateAggregation to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateAggregation to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateAggregation struct {
	obj *patternFlowLacpActorStateAggregation
}

type unMarshalPatternFlowLacpActorStateAggregation interface {
	// FromProto unmarshals PatternFlowLacpActorStateAggregation from protobuf object *otg.PatternFlowLacpActorStateAggregation
	FromProto(msg *otg.PatternFlowLacpActorStateAggregation) (PatternFlowLacpActorStateAggregation, error)
	// FromPbText unmarshals PatternFlowLacpActorStateAggregation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateAggregation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateAggregation from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateAggregation) Marshal() marshalPatternFlowLacpActorStateAggregation {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateAggregation{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateAggregation) Unmarshal() unMarshalPatternFlowLacpActorStateAggregation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateAggregation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateAggregation) ToProto() (*otg.PatternFlowLacpActorStateAggregation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateAggregation) FromProto(msg *otg.PatternFlowLacpActorStateAggregation) (PatternFlowLacpActorStateAggregation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateAggregation) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateAggregation) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateAggregation) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateAggregation) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateAggregation) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateAggregation) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateAggregation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateAggregation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateAggregation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateAggregation) Clone() (PatternFlowLacpActorStateAggregation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateAggregation()
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

func (obj *patternFlowLacpActorStateAggregation) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorStateAggregation is aggregation (1=Aggregatable, 0=Individual)
type PatternFlowLacpActorStateAggregation interface {
	Validation
	// msg marshals PatternFlowLacpActorStateAggregation to protobuf object *otg.PatternFlowLacpActorStateAggregation
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateAggregation
	// setMsg unmarshals PatternFlowLacpActorStateAggregation from protobuf object *otg.PatternFlowLacpActorStateAggregation
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateAggregation) PatternFlowLacpActorStateAggregation
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateAggregation
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateAggregation
	// validate validates PatternFlowLacpActorStateAggregation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateAggregation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorStateAggregationChoiceEnum, set in PatternFlowLacpActorStateAggregation
	Choice() PatternFlowLacpActorStateAggregationChoiceEnum
	// setChoice assigns PatternFlowLacpActorStateAggregationChoiceEnum provided by user to PatternFlowLacpActorStateAggregation
	setChoice(value PatternFlowLacpActorStateAggregationChoiceEnum) PatternFlowLacpActorStateAggregation
	// HasChoice checks if Choice has been set in PatternFlowLacpActorStateAggregation
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorStateAggregation.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorStateAggregation
	SetValue(value uint32) PatternFlowLacpActorStateAggregation
	// HasValue checks if Value has been set in PatternFlowLacpActorStateAggregation
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorStateAggregation.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorStateAggregation
	SetValues(value []uint32) PatternFlowLacpActorStateAggregation
	// Increment returns PatternFlowLacpActorStateAggregationCounter, set in PatternFlowLacpActorStateAggregation.
	// PatternFlowLacpActorStateAggregationCounter is integer counter pattern
	Increment() PatternFlowLacpActorStateAggregationCounter
	// SetIncrement assigns PatternFlowLacpActorStateAggregationCounter provided by user to PatternFlowLacpActorStateAggregation.
	// PatternFlowLacpActorStateAggregationCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorStateAggregationCounter) PatternFlowLacpActorStateAggregation
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorStateAggregation
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorStateAggregationCounter, set in PatternFlowLacpActorStateAggregation.
	// PatternFlowLacpActorStateAggregationCounter is integer counter pattern
	Decrement() PatternFlowLacpActorStateAggregationCounter
	// SetDecrement assigns PatternFlowLacpActorStateAggregationCounter provided by user to PatternFlowLacpActorStateAggregation.
	// PatternFlowLacpActorStateAggregationCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorStateAggregationCounter) PatternFlowLacpActorStateAggregation
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorStateAggregation
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorStateAggregationChoiceEnum string

// Enum of Choice on PatternFlowLacpActorStateAggregation
var PatternFlowLacpActorStateAggregationChoice = struct {
	VALUE     PatternFlowLacpActorStateAggregationChoiceEnum
	VALUES    PatternFlowLacpActorStateAggregationChoiceEnum
	INCREMENT PatternFlowLacpActorStateAggregationChoiceEnum
	DECREMENT PatternFlowLacpActorStateAggregationChoiceEnum
}{
	VALUE:     PatternFlowLacpActorStateAggregationChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorStateAggregationChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorStateAggregationChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorStateAggregationChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorStateAggregation) Choice() PatternFlowLacpActorStateAggregationChoiceEnum {
	return PatternFlowLacpActorStateAggregationChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorStateAggregation) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorStateAggregation) setChoice(value PatternFlowLacpActorStateAggregationChoiceEnum) PatternFlowLacpActorStateAggregation {
	intValue, ok := otg.PatternFlowLacpActorStateAggregation_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorStateAggregationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorStateAggregation_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorStateAggregationChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorStateAggregationChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorStateAggregationChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorStateAggregationCounter().msg()
	}

	if value == PatternFlowLacpActorStateAggregationChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorStateAggregationCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateAggregation) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorStateAggregationChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateAggregation) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorStateAggregation object
func (obj *patternFlowLacpActorStateAggregation) SetValue(value uint32) PatternFlowLacpActorStateAggregation {
	obj.setChoice(PatternFlowLacpActorStateAggregationChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorStateAggregation) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorStateAggregation object
func (obj *patternFlowLacpActorStateAggregation) SetValues(value []uint32) PatternFlowLacpActorStateAggregation {
	obj.setChoice(PatternFlowLacpActorStateAggregationChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateAggregationCounter
func (obj *patternFlowLacpActorStateAggregation) Increment() PatternFlowLacpActorStateAggregationCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorStateAggregationChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorStateAggregationCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateAggregationCounter
func (obj *patternFlowLacpActorStateAggregation) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorStateAggregationCounter value in the PatternFlowLacpActorStateAggregation object
func (obj *patternFlowLacpActorStateAggregation) SetIncrement(value PatternFlowLacpActorStateAggregationCounter) PatternFlowLacpActorStateAggregation {
	obj.setChoice(PatternFlowLacpActorStateAggregationChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateAggregationCounter
func (obj *patternFlowLacpActorStateAggregation) Decrement() PatternFlowLacpActorStateAggregationCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorStateAggregationChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorStateAggregationCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateAggregationCounter
func (obj *patternFlowLacpActorStateAggregation) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorStateAggregationCounter value in the PatternFlowLacpActorStateAggregation object
func (obj *patternFlowLacpActorStateAggregation) SetDecrement(value PatternFlowLacpActorStateAggregationCounter) PatternFlowLacpActorStateAggregation {
	obj.setChoice(PatternFlowLacpActorStateAggregationChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorStateAggregation) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateAggregation.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorStateAggregation.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpActorStateAggregation) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorStateAggregationChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateAggregationChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorStateAggregationChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateAggregationChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateAggregationChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorStateAggregationChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorStateAggregation")
			}
		} else {
			intVal := otg.PatternFlowLacpActorStateAggregation_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorStateAggregation_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
