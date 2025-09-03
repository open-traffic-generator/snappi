package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorSystemPriority *****
type patternFlowLacpActorSystemPriority struct {
	validation
	obj             *otg.PatternFlowLacpActorSystemPriority
	marshaller      marshalPatternFlowLacpActorSystemPriority
	unMarshaller    unMarshalPatternFlowLacpActorSystemPriority
	incrementHolder PatternFlowLacpActorSystemPriorityCounter
	decrementHolder PatternFlowLacpActorSystemPriorityCounter
}

func NewPatternFlowLacpActorSystemPriority() PatternFlowLacpActorSystemPriority {
	obj := patternFlowLacpActorSystemPriority{obj: &otg.PatternFlowLacpActorSystemPriority{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorSystemPriority) msg() *otg.PatternFlowLacpActorSystemPriority {
	return obj.obj
}

func (obj *patternFlowLacpActorSystemPriority) setMsg(msg *otg.PatternFlowLacpActorSystemPriority) PatternFlowLacpActorSystemPriority {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorSystemPriority struct {
	obj *patternFlowLacpActorSystemPriority
}

type marshalPatternFlowLacpActorSystemPriority interface {
	// ToProto marshals PatternFlowLacpActorSystemPriority to protobuf object *otg.PatternFlowLacpActorSystemPriority
	ToProto() (*otg.PatternFlowLacpActorSystemPriority, error)
	// ToPbText marshals PatternFlowLacpActorSystemPriority to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorSystemPriority to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorSystemPriority to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorSystemPriority struct {
	obj *patternFlowLacpActorSystemPriority
}

type unMarshalPatternFlowLacpActorSystemPriority interface {
	// FromProto unmarshals PatternFlowLacpActorSystemPriority from protobuf object *otg.PatternFlowLacpActorSystemPriority
	FromProto(msg *otg.PatternFlowLacpActorSystemPriority) (PatternFlowLacpActorSystemPriority, error)
	// FromPbText unmarshals PatternFlowLacpActorSystemPriority from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorSystemPriority from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorSystemPriority from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorSystemPriority) Marshal() marshalPatternFlowLacpActorSystemPriority {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorSystemPriority{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorSystemPriority) Unmarshal() unMarshalPatternFlowLacpActorSystemPriority {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorSystemPriority{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorSystemPriority) ToProto() (*otg.PatternFlowLacpActorSystemPriority, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorSystemPriority) FromProto(msg *otg.PatternFlowLacpActorSystemPriority) (PatternFlowLacpActorSystemPriority, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorSystemPriority) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorSystemPriority) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorSystemPriority) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorSystemPriority) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorSystemPriority) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorSystemPriority) FromJson(value string) error {
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

func (obj *patternFlowLacpActorSystemPriority) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorSystemPriority) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorSystemPriority) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorSystemPriority) Clone() (PatternFlowLacpActorSystemPriority, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorSystemPriority()
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

func (obj *patternFlowLacpActorSystemPriority) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorSystemPriority is the priority assigned to the Actor's system for use in aggregation. A lower numerical value indicates a higher priority. Used to select the active System ID when forming an aggregator.
type PatternFlowLacpActorSystemPriority interface {
	Validation
	// msg marshals PatternFlowLacpActorSystemPriority to protobuf object *otg.PatternFlowLacpActorSystemPriority
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorSystemPriority
	// setMsg unmarshals PatternFlowLacpActorSystemPriority from protobuf object *otg.PatternFlowLacpActorSystemPriority
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorSystemPriority) PatternFlowLacpActorSystemPriority
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorSystemPriority
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorSystemPriority
	// validate validates PatternFlowLacpActorSystemPriority
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorSystemPriority, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorSystemPriorityChoiceEnum, set in PatternFlowLacpActorSystemPriority
	Choice() PatternFlowLacpActorSystemPriorityChoiceEnum
	// setChoice assigns PatternFlowLacpActorSystemPriorityChoiceEnum provided by user to PatternFlowLacpActorSystemPriority
	setChoice(value PatternFlowLacpActorSystemPriorityChoiceEnum) PatternFlowLacpActorSystemPriority
	// HasChoice checks if Choice has been set in PatternFlowLacpActorSystemPriority
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorSystemPriority.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorSystemPriority
	SetValue(value uint32) PatternFlowLacpActorSystemPriority
	// HasValue checks if Value has been set in PatternFlowLacpActorSystemPriority
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorSystemPriority.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorSystemPriority
	SetValues(value []uint32) PatternFlowLacpActorSystemPriority
	// Increment returns PatternFlowLacpActorSystemPriorityCounter, set in PatternFlowLacpActorSystemPriority.
	// PatternFlowLacpActorSystemPriorityCounter is integer counter pattern
	Increment() PatternFlowLacpActorSystemPriorityCounter
	// SetIncrement assigns PatternFlowLacpActorSystemPriorityCounter provided by user to PatternFlowLacpActorSystemPriority.
	// PatternFlowLacpActorSystemPriorityCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorSystemPriorityCounter) PatternFlowLacpActorSystemPriority
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorSystemPriority
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorSystemPriorityCounter, set in PatternFlowLacpActorSystemPriority.
	// PatternFlowLacpActorSystemPriorityCounter is integer counter pattern
	Decrement() PatternFlowLacpActorSystemPriorityCounter
	// SetDecrement assigns PatternFlowLacpActorSystemPriorityCounter provided by user to PatternFlowLacpActorSystemPriority.
	// PatternFlowLacpActorSystemPriorityCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorSystemPriorityCounter) PatternFlowLacpActorSystemPriority
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorSystemPriority
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorSystemPriorityChoiceEnum string

// Enum of Choice on PatternFlowLacpActorSystemPriority
var PatternFlowLacpActorSystemPriorityChoice = struct {
	VALUE     PatternFlowLacpActorSystemPriorityChoiceEnum
	VALUES    PatternFlowLacpActorSystemPriorityChoiceEnum
	INCREMENT PatternFlowLacpActorSystemPriorityChoiceEnum
	DECREMENT PatternFlowLacpActorSystemPriorityChoiceEnum
}{
	VALUE:     PatternFlowLacpActorSystemPriorityChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorSystemPriorityChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorSystemPriorityChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorSystemPriorityChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorSystemPriority) Choice() PatternFlowLacpActorSystemPriorityChoiceEnum {
	return PatternFlowLacpActorSystemPriorityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorSystemPriority) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorSystemPriority) setChoice(value PatternFlowLacpActorSystemPriorityChoiceEnum) PatternFlowLacpActorSystemPriority {
	intValue, ok := otg.PatternFlowLacpActorSystemPriority_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorSystemPriorityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorSystemPriority_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorSystemPriorityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorSystemPriorityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorSystemPriorityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorSystemPriorityCounter().msg()
	}

	if value == PatternFlowLacpActorSystemPriorityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorSystemPriorityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorSystemPriority) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorSystemPriorityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorSystemPriority) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorSystemPriority object
func (obj *patternFlowLacpActorSystemPriority) SetValue(value uint32) PatternFlowLacpActorSystemPriority {
	obj.setChoice(PatternFlowLacpActorSystemPriorityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorSystemPriority) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorSystemPriority object
func (obj *patternFlowLacpActorSystemPriority) SetValues(value []uint32) PatternFlowLacpActorSystemPriority {
	obj.setChoice(PatternFlowLacpActorSystemPriorityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorSystemPriorityCounter
func (obj *patternFlowLacpActorSystemPriority) Increment() PatternFlowLacpActorSystemPriorityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorSystemPriorityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorSystemPriorityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorSystemPriorityCounter
func (obj *patternFlowLacpActorSystemPriority) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorSystemPriorityCounter value in the PatternFlowLacpActorSystemPriority object
func (obj *patternFlowLacpActorSystemPriority) SetIncrement(value PatternFlowLacpActorSystemPriorityCounter) PatternFlowLacpActorSystemPriority {
	obj.setChoice(PatternFlowLacpActorSystemPriorityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorSystemPriorityCounter
func (obj *patternFlowLacpActorSystemPriority) Decrement() PatternFlowLacpActorSystemPriorityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorSystemPriorityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorSystemPriorityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorSystemPriorityCounter
func (obj *patternFlowLacpActorSystemPriority) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorSystemPriorityCounter value in the PatternFlowLacpActorSystemPriority object
func (obj *patternFlowLacpActorSystemPriority) SetDecrement(value PatternFlowLacpActorSystemPriorityCounter) PatternFlowLacpActorSystemPriority {
	obj.setChoice(PatternFlowLacpActorSystemPriorityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorSystemPriority) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorSystemPriority.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorSystemPriority.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpActorSystemPriority) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorSystemPriorityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorSystemPriorityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorSystemPriorityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorSystemPriorityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorSystemPriorityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorSystemPriorityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorSystemPriority")
			}
		} else {
			intVal := otg.PatternFlowLacpActorSystemPriority_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorSystemPriority_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
