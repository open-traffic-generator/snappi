package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateSynchronization *****
type patternFlowLacpduActorStateSynchronization struct {
	validation
	obj             *otg.PatternFlowLacpduActorStateSynchronization
	marshaller      marshalPatternFlowLacpduActorStateSynchronization
	unMarshaller    unMarshalPatternFlowLacpduActorStateSynchronization
	incrementHolder PatternFlowLacpduActorStateSynchronizationCounter
	decrementHolder PatternFlowLacpduActorStateSynchronizationCounter
}

func NewPatternFlowLacpduActorStateSynchronization() PatternFlowLacpduActorStateSynchronization {
	obj := patternFlowLacpduActorStateSynchronization{obj: &otg.PatternFlowLacpduActorStateSynchronization{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateSynchronization) msg() *otg.PatternFlowLacpduActorStateSynchronization {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateSynchronization) setMsg(msg *otg.PatternFlowLacpduActorStateSynchronization) PatternFlowLacpduActorStateSynchronization {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateSynchronization struct {
	obj *patternFlowLacpduActorStateSynchronization
}

type marshalPatternFlowLacpduActorStateSynchronization interface {
	// ToProto marshals PatternFlowLacpduActorStateSynchronization to protobuf object *otg.PatternFlowLacpduActorStateSynchronization
	ToProto() (*otg.PatternFlowLacpduActorStateSynchronization, error)
	// ToPbText marshals PatternFlowLacpduActorStateSynchronization to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateSynchronization to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateSynchronization to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateSynchronization struct {
	obj *patternFlowLacpduActorStateSynchronization
}

type unMarshalPatternFlowLacpduActorStateSynchronization interface {
	// FromProto unmarshals PatternFlowLacpduActorStateSynchronization from protobuf object *otg.PatternFlowLacpduActorStateSynchronization
	FromProto(msg *otg.PatternFlowLacpduActorStateSynchronization) (PatternFlowLacpduActorStateSynchronization, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateSynchronization from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateSynchronization from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateSynchronization from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateSynchronization) Marshal() marshalPatternFlowLacpduActorStateSynchronization {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateSynchronization{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateSynchronization) Unmarshal() unMarshalPatternFlowLacpduActorStateSynchronization {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateSynchronization{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateSynchronization) ToProto() (*otg.PatternFlowLacpduActorStateSynchronization, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateSynchronization) FromProto(msg *otg.PatternFlowLacpduActorStateSynchronization) (PatternFlowLacpduActorStateSynchronization, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateSynchronization) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateSynchronization) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateSynchronization) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateSynchronization) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateSynchronization) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateSynchronization) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateSynchronization) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateSynchronization) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateSynchronization) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateSynchronization) Clone() (PatternFlowLacpduActorStateSynchronization, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateSynchronization()
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

func (obj *patternFlowLacpduActorStateSynchronization) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorStateSynchronization is synchronization (1=In Sync, 0=Out of Sync)
type PatternFlowLacpduActorStateSynchronization interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateSynchronization to protobuf object *otg.PatternFlowLacpduActorStateSynchronization
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateSynchronization
	// setMsg unmarshals PatternFlowLacpduActorStateSynchronization from protobuf object *otg.PatternFlowLacpduActorStateSynchronization
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateSynchronization) PatternFlowLacpduActorStateSynchronization
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateSynchronization
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateSynchronization
	// validate validates PatternFlowLacpduActorStateSynchronization
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateSynchronization, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorStateSynchronizationChoiceEnum, set in PatternFlowLacpduActorStateSynchronization
	Choice() PatternFlowLacpduActorStateSynchronizationChoiceEnum
	// setChoice assigns PatternFlowLacpduActorStateSynchronizationChoiceEnum provided by user to PatternFlowLacpduActorStateSynchronization
	setChoice(value PatternFlowLacpduActorStateSynchronizationChoiceEnum) PatternFlowLacpduActorStateSynchronization
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorStateSynchronization
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorStateSynchronization.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorStateSynchronization
	SetValue(value uint32) PatternFlowLacpduActorStateSynchronization
	// HasValue checks if Value has been set in PatternFlowLacpduActorStateSynchronization
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorStateSynchronization.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorStateSynchronization
	SetValues(value []uint32) PatternFlowLacpduActorStateSynchronization
	// Increment returns PatternFlowLacpduActorStateSynchronizationCounter, set in PatternFlowLacpduActorStateSynchronization.
	// PatternFlowLacpduActorStateSynchronizationCounter is integer counter pattern
	Increment() PatternFlowLacpduActorStateSynchronizationCounter
	// SetIncrement assigns PatternFlowLacpduActorStateSynchronizationCounter provided by user to PatternFlowLacpduActorStateSynchronization.
	// PatternFlowLacpduActorStateSynchronizationCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorStateSynchronizationCounter) PatternFlowLacpduActorStateSynchronization
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorStateSynchronization
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorStateSynchronizationCounter, set in PatternFlowLacpduActorStateSynchronization.
	// PatternFlowLacpduActorStateSynchronizationCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorStateSynchronizationCounter
	// SetDecrement assigns PatternFlowLacpduActorStateSynchronizationCounter provided by user to PatternFlowLacpduActorStateSynchronization.
	// PatternFlowLacpduActorStateSynchronizationCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorStateSynchronizationCounter) PatternFlowLacpduActorStateSynchronization
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorStateSynchronization
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorStateSynchronizationChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorStateSynchronization
var PatternFlowLacpduActorStateSynchronizationChoice = struct {
	VALUE     PatternFlowLacpduActorStateSynchronizationChoiceEnum
	VALUES    PatternFlowLacpduActorStateSynchronizationChoiceEnum
	INCREMENT PatternFlowLacpduActorStateSynchronizationChoiceEnum
	DECREMENT PatternFlowLacpduActorStateSynchronizationChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorStateSynchronizationChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorStateSynchronizationChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorStateSynchronizationChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorStateSynchronizationChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorStateSynchronization) Choice() PatternFlowLacpduActorStateSynchronizationChoiceEnum {
	return PatternFlowLacpduActorStateSynchronizationChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorStateSynchronization) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorStateSynchronization) setChoice(value PatternFlowLacpduActorStateSynchronizationChoiceEnum) PatternFlowLacpduActorStateSynchronization {
	intValue, ok := otg.PatternFlowLacpduActorStateSynchronization_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorStateSynchronizationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorStateSynchronization_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorStateSynchronizationChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorStateSynchronizationChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorStateSynchronizationChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorStateSynchronizationCounter().msg()
	}

	if value == PatternFlowLacpduActorStateSynchronizationChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorStateSynchronizationCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateSynchronization) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorStateSynchronizationChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateSynchronization) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorStateSynchronization object
func (obj *patternFlowLacpduActorStateSynchronization) SetValue(value uint32) PatternFlowLacpduActorStateSynchronization {
	obj.setChoice(PatternFlowLacpduActorStateSynchronizationChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorStateSynchronization) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorStateSynchronization object
func (obj *patternFlowLacpduActorStateSynchronization) SetValues(value []uint32) PatternFlowLacpduActorStateSynchronization {
	obj.setChoice(PatternFlowLacpduActorStateSynchronizationChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateSynchronizationCounter
func (obj *patternFlowLacpduActorStateSynchronization) Increment() PatternFlowLacpduActorStateSynchronizationCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorStateSynchronizationChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorStateSynchronizationCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateSynchronizationCounter
func (obj *patternFlowLacpduActorStateSynchronization) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorStateSynchronizationCounter value in the PatternFlowLacpduActorStateSynchronization object
func (obj *patternFlowLacpduActorStateSynchronization) SetIncrement(value PatternFlowLacpduActorStateSynchronizationCounter) PatternFlowLacpduActorStateSynchronization {
	obj.setChoice(PatternFlowLacpduActorStateSynchronizationChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateSynchronizationCounter
func (obj *patternFlowLacpduActorStateSynchronization) Decrement() PatternFlowLacpduActorStateSynchronizationCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorStateSynchronizationChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorStateSynchronizationCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateSynchronizationCounter
func (obj *patternFlowLacpduActorStateSynchronization) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorStateSynchronizationCounter value in the PatternFlowLacpduActorStateSynchronization object
func (obj *patternFlowLacpduActorStateSynchronization) SetDecrement(value PatternFlowLacpduActorStateSynchronizationCounter) PatternFlowLacpduActorStateSynchronization {
	obj.setChoice(PatternFlowLacpduActorStateSynchronizationChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorStateSynchronization) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateSynchronization.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorStateSynchronization.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduActorStateSynchronization) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorStateSynchronizationChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateSynchronizationChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorStateSynchronizationChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateSynchronizationChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateSynchronizationChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorStateSynchronizationChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorStateSynchronization")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorStateSynchronization_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorStateSynchronization_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
