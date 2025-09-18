package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorSystemPriority *****
type patternFlowLacpduActorSystemPriority struct {
	validation
	obj             *otg.PatternFlowLacpduActorSystemPriority
	marshaller      marshalPatternFlowLacpduActorSystemPriority
	unMarshaller    unMarshalPatternFlowLacpduActorSystemPriority
	incrementHolder PatternFlowLacpduActorSystemPriorityCounter
	decrementHolder PatternFlowLacpduActorSystemPriorityCounter
}

func NewPatternFlowLacpduActorSystemPriority() PatternFlowLacpduActorSystemPriority {
	obj := patternFlowLacpduActorSystemPriority{obj: &otg.PatternFlowLacpduActorSystemPriority{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorSystemPriority) msg() *otg.PatternFlowLacpduActorSystemPriority {
	return obj.obj
}

func (obj *patternFlowLacpduActorSystemPriority) setMsg(msg *otg.PatternFlowLacpduActorSystemPriority) PatternFlowLacpduActorSystemPriority {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorSystemPriority struct {
	obj *patternFlowLacpduActorSystemPriority
}

type marshalPatternFlowLacpduActorSystemPriority interface {
	// ToProto marshals PatternFlowLacpduActorSystemPriority to protobuf object *otg.PatternFlowLacpduActorSystemPriority
	ToProto() (*otg.PatternFlowLacpduActorSystemPriority, error)
	// ToPbText marshals PatternFlowLacpduActorSystemPriority to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorSystemPriority to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorSystemPriority to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorSystemPriority struct {
	obj *patternFlowLacpduActorSystemPriority
}

type unMarshalPatternFlowLacpduActorSystemPriority interface {
	// FromProto unmarshals PatternFlowLacpduActorSystemPriority from protobuf object *otg.PatternFlowLacpduActorSystemPriority
	FromProto(msg *otg.PatternFlowLacpduActorSystemPriority) (PatternFlowLacpduActorSystemPriority, error)
	// FromPbText unmarshals PatternFlowLacpduActorSystemPriority from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorSystemPriority from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorSystemPriority from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorSystemPriority) Marshal() marshalPatternFlowLacpduActorSystemPriority {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorSystemPriority{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorSystemPriority) Unmarshal() unMarshalPatternFlowLacpduActorSystemPriority {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorSystemPriority{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorSystemPriority) ToProto() (*otg.PatternFlowLacpduActorSystemPriority, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorSystemPriority) FromProto(msg *otg.PatternFlowLacpduActorSystemPriority) (PatternFlowLacpduActorSystemPriority, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorSystemPriority) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorSystemPriority) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorSystemPriority) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorSystemPriority) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorSystemPriority) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorSystemPriority) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorSystemPriority) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorSystemPriority) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorSystemPriority) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorSystemPriority) Clone() (PatternFlowLacpduActorSystemPriority, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorSystemPriority()
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

func (obj *patternFlowLacpduActorSystemPriority) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorSystemPriority is the priority assigned to the Actor's system for use in aggregation. A lower numerical value indicates a higher priority. Used to select the active System ID when forming an aggregator.
type PatternFlowLacpduActorSystemPriority interface {
	Validation
	// msg marshals PatternFlowLacpduActorSystemPriority to protobuf object *otg.PatternFlowLacpduActorSystemPriority
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorSystemPriority
	// setMsg unmarshals PatternFlowLacpduActorSystemPriority from protobuf object *otg.PatternFlowLacpduActorSystemPriority
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorSystemPriority) PatternFlowLacpduActorSystemPriority
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorSystemPriority
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorSystemPriority
	// validate validates PatternFlowLacpduActorSystemPriority
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorSystemPriority, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorSystemPriorityChoiceEnum, set in PatternFlowLacpduActorSystemPriority
	Choice() PatternFlowLacpduActorSystemPriorityChoiceEnum
	// setChoice assigns PatternFlowLacpduActorSystemPriorityChoiceEnum provided by user to PatternFlowLacpduActorSystemPriority
	setChoice(value PatternFlowLacpduActorSystemPriorityChoiceEnum) PatternFlowLacpduActorSystemPriority
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorSystemPriority
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorSystemPriority.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorSystemPriority
	SetValue(value uint32) PatternFlowLacpduActorSystemPriority
	// HasValue checks if Value has been set in PatternFlowLacpduActorSystemPriority
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorSystemPriority.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorSystemPriority
	SetValues(value []uint32) PatternFlowLacpduActorSystemPriority
	// Increment returns PatternFlowLacpduActorSystemPriorityCounter, set in PatternFlowLacpduActorSystemPriority.
	// PatternFlowLacpduActorSystemPriorityCounter is integer counter pattern
	Increment() PatternFlowLacpduActorSystemPriorityCounter
	// SetIncrement assigns PatternFlowLacpduActorSystemPriorityCounter provided by user to PatternFlowLacpduActorSystemPriority.
	// PatternFlowLacpduActorSystemPriorityCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorSystemPriorityCounter) PatternFlowLacpduActorSystemPriority
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorSystemPriority
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorSystemPriorityCounter, set in PatternFlowLacpduActorSystemPriority.
	// PatternFlowLacpduActorSystemPriorityCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorSystemPriorityCounter
	// SetDecrement assigns PatternFlowLacpduActorSystemPriorityCounter provided by user to PatternFlowLacpduActorSystemPriority.
	// PatternFlowLacpduActorSystemPriorityCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorSystemPriorityCounter) PatternFlowLacpduActorSystemPriority
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorSystemPriority
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorSystemPriorityChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorSystemPriority
var PatternFlowLacpduActorSystemPriorityChoice = struct {
	VALUE     PatternFlowLacpduActorSystemPriorityChoiceEnum
	VALUES    PatternFlowLacpduActorSystemPriorityChoiceEnum
	INCREMENT PatternFlowLacpduActorSystemPriorityChoiceEnum
	DECREMENT PatternFlowLacpduActorSystemPriorityChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorSystemPriorityChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorSystemPriorityChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorSystemPriorityChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorSystemPriorityChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorSystemPriority) Choice() PatternFlowLacpduActorSystemPriorityChoiceEnum {
	return PatternFlowLacpduActorSystemPriorityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorSystemPriority) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorSystemPriority) setChoice(value PatternFlowLacpduActorSystemPriorityChoiceEnum) PatternFlowLacpduActorSystemPriority {
	intValue, ok := otg.PatternFlowLacpduActorSystemPriority_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorSystemPriorityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorSystemPriority_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorSystemPriorityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorSystemPriorityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorSystemPriorityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorSystemPriorityCounter().msg()
	}

	if value == PatternFlowLacpduActorSystemPriorityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorSystemPriorityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorSystemPriority) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorSystemPriorityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorSystemPriority) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorSystemPriority object
func (obj *patternFlowLacpduActorSystemPriority) SetValue(value uint32) PatternFlowLacpduActorSystemPriority {
	obj.setChoice(PatternFlowLacpduActorSystemPriorityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorSystemPriority) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorSystemPriority object
func (obj *patternFlowLacpduActorSystemPriority) SetValues(value []uint32) PatternFlowLacpduActorSystemPriority {
	obj.setChoice(PatternFlowLacpduActorSystemPriorityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorSystemPriorityCounter
func (obj *patternFlowLacpduActorSystemPriority) Increment() PatternFlowLacpduActorSystemPriorityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorSystemPriorityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorSystemPriorityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorSystemPriorityCounter
func (obj *patternFlowLacpduActorSystemPriority) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorSystemPriorityCounter value in the PatternFlowLacpduActorSystemPriority object
func (obj *patternFlowLacpduActorSystemPriority) SetIncrement(value PatternFlowLacpduActorSystemPriorityCounter) PatternFlowLacpduActorSystemPriority {
	obj.setChoice(PatternFlowLacpduActorSystemPriorityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorSystemPriorityCounter
func (obj *patternFlowLacpduActorSystemPriority) Decrement() PatternFlowLacpduActorSystemPriorityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorSystemPriorityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorSystemPriorityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorSystemPriorityCounter
func (obj *patternFlowLacpduActorSystemPriority) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorSystemPriorityCounter value in the PatternFlowLacpduActorSystemPriority object
func (obj *patternFlowLacpduActorSystemPriority) SetDecrement(value PatternFlowLacpduActorSystemPriorityCounter) PatternFlowLacpduActorSystemPriority {
	obj.setChoice(PatternFlowLacpduActorSystemPriorityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorSystemPriority) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorSystemPriority.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorSystemPriority.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpduActorSystemPriority) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorSystemPriorityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorSystemPriorityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorSystemPriorityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorSystemPriorityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorSystemPriorityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorSystemPriorityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorSystemPriority")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorSystemPriority_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorSystemPriority_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
