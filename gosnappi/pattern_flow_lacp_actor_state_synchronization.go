package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateSynchronization *****
type patternFlowLacpActorStateSynchronization struct {
	validation
	obj             *otg.PatternFlowLacpActorStateSynchronization
	marshaller      marshalPatternFlowLacpActorStateSynchronization
	unMarshaller    unMarshalPatternFlowLacpActorStateSynchronization
	incrementHolder PatternFlowLacpActorStateSynchronizationCounter
	decrementHolder PatternFlowLacpActorStateSynchronizationCounter
}

func NewPatternFlowLacpActorStateSynchronization() PatternFlowLacpActorStateSynchronization {
	obj := patternFlowLacpActorStateSynchronization{obj: &otg.PatternFlowLacpActorStateSynchronization{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateSynchronization) msg() *otg.PatternFlowLacpActorStateSynchronization {
	return obj.obj
}

func (obj *patternFlowLacpActorStateSynchronization) setMsg(msg *otg.PatternFlowLacpActorStateSynchronization) PatternFlowLacpActorStateSynchronization {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateSynchronization struct {
	obj *patternFlowLacpActorStateSynchronization
}

type marshalPatternFlowLacpActorStateSynchronization interface {
	// ToProto marshals PatternFlowLacpActorStateSynchronization to protobuf object *otg.PatternFlowLacpActorStateSynchronization
	ToProto() (*otg.PatternFlowLacpActorStateSynchronization, error)
	// ToPbText marshals PatternFlowLacpActorStateSynchronization to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateSynchronization to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateSynchronization to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateSynchronization struct {
	obj *patternFlowLacpActorStateSynchronization
}

type unMarshalPatternFlowLacpActorStateSynchronization interface {
	// FromProto unmarshals PatternFlowLacpActorStateSynchronization from protobuf object *otg.PatternFlowLacpActorStateSynchronization
	FromProto(msg *otg.PatternFlowLacpActorStateSynchronization) (PatternFlowLacpActorStateSynchronization, error)
	// FromPbText unmarshals PatternFlowLacpActorStateSynchronization from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateSynchronization from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateSynchronization from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateSynchronization) Marshal() marshalPatternFlowLacpActorStateSynchronization {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateSynchronization{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateSynchronization) Unmarshal() unMarshalPatternFlowLacpActorStateSynchronization {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateSynchronization{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateSynchronization) ToProto() (*otg.PatternFlowLacpActorStateSynchronization, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateSynchronization) FromProto(msg *otg.PatternFlowLacpActorStateSynchronization) (PatternFlowLacpActorStateSynchronization, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateSynchronization) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateSynchronization) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateSynchronization) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateSynchronization) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateSynchronization) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateSynchronization) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateSynchronization) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateSynchronization) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateSynchronization) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateSynchronization) Clone() (PatternFlowLacpActorStateSynchronization, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateSynchronization()
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

func (obj *patternFlowLacpActorStateSynchronization) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorStateSynchronization is synchronization (1=In Sync, 0=Out of Sync)
type PatternFlowLacpActorStateSynchronization interface {
	Validation
	// msg marshals PatternFlowLacpActorStateSynchronization to protobuf object *otg.PatternFlowLacpActorStateSynchronization
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateSynchronization
	// setMsg unmarshals PatternFlowLacpActorStateSynchronization from protobuf object *otg.PatternFlowLacpActorStateSynchronization
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateSynchronization) PatternFlowLacpActorStateSynchronization
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateSynchronization
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateSynchronization
	// validate validates PatternFlowLacpActorStateSynchronization
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateSynchronization, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorStateSynchronizationChoiceEnum, set in PatternFlowLacpActorStateSynchronization
	Choice() PatternFlowLacpActorStateSynchronizationChoiceEnum
	// setChoice assigns PatternFlowLacpActorStateSynchronizationChoiceEnum provided by user to PatternFlowLacpActorStateSynchronization
	setChoice(value PatternFlowLacpActorStateSynchronizationChoiceEnum) PatternFlowLacpActorStateSynchronization
	// HasChoice checks if Choice has been set in PatternFlowLacpActorStateSynchronization
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorStateSynchronization.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorStateSynchronization
	SetValue(value uint32) PatternFlowLacpActorStateSynchronization
	// HasValue checks if Value has been set in PatternFlowLacpActorStateSynchronization
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorStateSynchronization.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorStateSynchronization
	SetValues(value []uint32) PatternFlowLacpActorStateSynchronization
	// Increment returns PatternFlowLacpActorStateSynchronizationCounter, set in PatternFlowLacpActorStateSynchronization.
	// PatternFlowLacpActorStateSynchronizationCounter is integer counter pattern
	Increment() PatternFlowLacpActorStateSynchronizationCounter
	// SetIncrement assigns PatternFlowLacpActorStateSynchronizationCounter provided by user to PatternFlowLacpActorStateSynchronization.
	// PatternFlowLacpActorStateSynchronizationCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorStateSynchronizationCounter) PatternFlowLacpActorStateSynchronization
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorStateSynchronization
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorStateSynchronizationCounter, set in PatternFlowLacpActorStateSynchronization.
	// PatternFlowLacpActorStateSynchronizationCounter is integer counter pattern
	Decrement() PatternFlowLacpActorStateSynchronizationCounter
	// SetDecrement assigns PatternFlowLacpActorStateSynchronizationCounter provided by user to PatternFlowLacpActorStateSynchronization.
	// PatternFlowLacpActorStateSynchronizationCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorStateSynchronizationCounter) PatternFlowLacpActorStateSynchronization
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorStateSynchronization
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorStateSynchronizationChoiceEnum string

// Enum of Choice on PatternFlowLacpActorStateSynchronization
var PatternFlowLacpActorStateSynchronizationChoice = struct {
	VALUE     PatternFlowLacpActorStateSynchronizationChoiceEnum
	VALUES    PatternFlowLacpActorStateSynchronizationChoiceEnum
	INCREMENT PatternFlowLacpActorStateSynchronizationChoiceEnum
	DECREMENT PatternFlowLacpActorStateSynchronizationChoiceEnum
}{
	VALUE:     PatternFlowLacpActorStateSynchronizationChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorStateSynchronizationChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorStateSynchronizationChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorStateSynchronizationChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorStateSynchronization) Choice() PatternFlowLacpActorStateSynchronizationChoiceEnum {
	return PatternFlowLacpActorStateSynchronizationChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorStateSynchronization) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorStateSynchronization) setChoice(value PatternFlowLacpActorStateSynchronizationChoiceEnum) PatternFlowLacpActorStateSynchronization {
	intValue, ok := otg.PatternFlowLacpActorStateSynchronization_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorStateSynchronizationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorStateSynchronization_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorStateSynchronizationChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorStateSynchronizationChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorStateSynchronizationChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorStateSynchronizationCounter().msg()
	}

	if value == PatternFlowLacpActorStateSynchronizationChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorStateSynchronizationCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateSynchronization) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorStateSynchronizationChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateSynchronization) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorStateSynchronization object
func (obj *patternFlowLacpActorStateSynchronization) SetValue(value uint32) PatternFlowLacpActorStateSynchronization {
	obj.setChoice(PatternFlowLacpActorStateSynchronizationChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorStateSynchronization) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorStateSynchronization object
func (obj *patternFlowLacpActorStateSynchronization) SetValues(value []uint32) PatternFlowLacpActorStateSynchronization {
	obj.setChoice(PatternFlowLacpActorStateSynchronizationChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateSynchronizationCounter
func (obj *patternFlowLacpActorStateSynchronization) Increment() PatternFlowLacpActorStateSynchronizationCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorStateSynchronizationChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorStateSynchronizationCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateSynchronizationCounter
func (obj *patternFlowLacpActorStateSynchronization) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorStateSynchronizationCounter value in the PatternFlowLacpActorStateSynchronization object
func (obj *patternFlowLacpActorStateSynchronization) SetIncrement(value PatternFlowLacpActorStateSynchronizationCounter) PatternFlowLacpActorStateSynchronization {
	obj.setChoice(PatternFlowLacpActorStateSynchronizationChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateSynchronizationCounter
func (obj *patternFlowLacpActorStateSynchronization) Decrement() PatternFlowLacpActorStateSynchronizationCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorStateSynchronizationChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorStateSynchronizationCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateSynchronizationCounter
func (obj *patternFlowLacpActorStateSynchronization) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorStateSynchronizationCounter value in the PatternFlowLacpActorStateSynchronization object
func (obj *patternFlowLacpActorStateSynchronization) SetDecrement(value PatternFlowLacpActorStateSynchronizationCounter) PatternFlowLacpActorStateSynchronization {
	obj.setChoice(PatternFlowLacpActorStateSynchronizationChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorStateSynchronization) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateSynchronization.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorStateSynchronization.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpActorStateSynchronization) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorStateSynchronizationChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateSynchronizationChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorStateSynchronizationChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateSynchronizationChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateSynchronizationChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorStateSynchronizationChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorStateSynchronization")
			}
		} else {
			intVal := otg.PatternFlowLacpActorStateSynchronization_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorStateSynchronization_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
