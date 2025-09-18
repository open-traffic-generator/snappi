package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorPortPriority *****
type patternFlowLacpduActorPortPriority struct {
	validation
	obj             *otg.PatternFlowLacpduActorPortPriority
	marshaller      marshalPatternFlowLacpduActorPortPriority
	unMarshaller    unMarshalPatternFlowLacpduActorPortPriority
	incrementHolder PatternFlowLacpduActorPortPriorityCounter
	decrementHolder PatternFlowLacpduActorPortPriorityCounter
}

func NewPatternFlowLacpduActorPortPriority() PatternFlowLacpduActorPortPriority {
	obj := patternFlowLacpduActorPortPriority{obj: &otg.PatternFlowLacpduActorPortPriority{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorPortPriority) msg() *otg.PatternFlowLacpduActorPortPriority {
	return obj.obj
}

func (obj *patternFlowLacpduActorPortPriority) setMsg(msg *otg.PatternFlowLacpduActorPortPriority) PatternFlowLacpduActorPortPriority {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorPortPriority struct {
	obj *patternFlowLacpduActorPortPriority
}

type marshalPatternFlowLacpduActorPortPriority interface {
	// ToProto marshals PatternFlowLacpduActorPortPriority to protobuf object *otg.PatternFlowLacpduActorPortPriority
	ToProto() (*otg.PatternFlowLacpduActorPortPriority, error)
	// ToPbText marshals PatternFlowLacpduActorPortPriority to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorPortPriority to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorPortPriority to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorPortPriority struct {
	obj *patternFlowLacpduActorPortPriority
}

type unMarshalPatternFlowLacpduActorPortPriority interface {
	// FromProto unmarshals PatternFlowLacpduActorPortPriority from protobuf object *otg.PatternFlowLacpduActorPortPriority
	FromProto(msg *otg.PatternFlowLacpduActorPortPriority) (PatternFlowLacpduActorPortPriority, error)
	// FromPbText unmarshals PatternFlowLacpduActorPortPriority from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorPortPriority from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorPortPriority from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorPortPriority) Marshal() marshalPatternFlowLacpduActorPortPriority {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorPortPriority{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorPortPriority) Unmarshal() unMarshalPatternFlowLacpduActorPortPriority {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorPortPriority{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorPortPriority) ToProto() (*otg.PatternFlowLacpduActorPortPriority, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorPortPriority) FromProto(msg *otg.PatternFlowLacpduActorPortPriority) (PatternFlowLacpduActorPortPriority, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorPortPriority) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorPortPriority) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorPortPriority) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorPortPriority) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorPortPriority) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorPortPriority) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorPortPriority) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorPortPriority) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorPortPriority) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorPortPriority) Clone() (PatternFlowLacpduActorPortPriority, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorPortPriority()
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

func (obj *patternFlowLacpduActorPortPriority) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorPortPriority is the priority assigned to the Actor's port. A lower numerical value indicates a higher priority. Used to prioritize ports for inclusion in a Link Aggregation Group (LAG) when the group is full.
type PatternFlowLacpduActorPortPriority interface {
	Validation
	// msg marshals PatternFlowLacpduActorPortPriority to protobuf object *otg.PatternFlowLacpduActorPortPriority
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorPortPriority
	// setMsg unmarshals PatternFlowLacpduActorPortPriority from protobuf object *otg.PatternFlowLacpduActorPortPriority
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorPortPriority) PatternFlowLacpduActorPortPriority
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorPortPriority
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorPortPriority
	// validate validates PatternFlowLacpduActorPortPriority
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorPortPriority, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorPortPriorityChoiceEnum, set in PatternFlowLacpduActorPortPriority
	Choice() PatternFlowLacpduActorPortPriorityChoiceEnum
	// setChoice assigns PatternFlowLacpduActorPortPriorityChoiceEnum provided by user to PatternFlowLacpduActorPortPriority
	setChoice(value PatternFlowLacpduActorPortPriorityChoiceEnum) PatternFlowLacpduActorPortPriority
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorPortPriority
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorPortPriority.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorPortPriority
	SetValue(value uint32) PatternFlowLacpduActorPortPriority
	// HasValue checks if Value has been set in PatternFlowLacpduActorPortPriority
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorPortPriority.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorPortPriority
	SetValues(value []uint32) PatternFlowLacpduActorPortPriority
	// Increment returns PatternFlowLacpduActorPortPriorityCounter, set in PatternFlowLacpduActorPortPriority.
	// PatternFlowLacpduActorPortPriorityCounter is integer counter pattern
	Increment() PatternFlowLacpduActorPortPriorityCounter
	// SetIncrement assigns PatternFlowLacpduActorPortPriorityCounter provided by user to PatternFlowLacpduActorPortPriority.
	// PatternFlowLacpduActorPortPriorityCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorPortPriorityCounter) PatternFlowLacpduActorPortPriority
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorPortPriority
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorPortPriorityCounter, set in PatternFlowLacpduActorPortPriority.
	// PatternFlowLacpduActorPortPriorityCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorPortPriorityCounter
	// SetDecrement assigns PatternFlowLacpduActorPortPriorityCounter provided by user to PatternFlowLacpduActorPortPriority.
	// PatternFlowLacpduActorPortPriorityCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorPortPriorityCounter) PatternFlowLacpduActorPortPriority
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorPortPriority
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorPortPriorityChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorPortPriority
var PatternFlowLacpduActorPortPriorityChoice = struct {
	VALUE     PatternFlowLacpduActorPortPriorityChoiceEnum
	VALUES    PatternFlowLacpduActorPortPriorityChoiceEnum
	INCREMENT PatternFlowLacpduActorPortPriorityChoiceEnum
	DECREMENT PatternFlowLacpduActorPortPriorityChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorPortPriorityChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorPortPriorityChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorPortPriorityChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorPortPriorityChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorPortPriority) Choice() PatternFlowLacpduActorPortPriorityChoiceEnum {
	return PatternFlowLacpduActorPortPriorityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorPortPriority) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorPortPriority) setChoice(value PatternFlowLacpduActorPortPriorityChoiceEnum) PatternFlowLacpduActorPortPriority {
	intValue, ok := otg.PatternFlowLacpduActorPortPriority_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorPortPriorityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorPortPriority_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorPortPriorityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorPortPriorityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorPortPriorityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorPortPriorityCounter().msg()
	}

	if value == PatternFlowLacpduActorPortPriorityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorPortPriorityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorPortPriority) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorPortPriorityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorPortPriority) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorPortPriority object
func (obj *patternFlowLacpduActorPortPriority) SetValue(value uint32) PatternFlowLacpduActorPortPriority {
	obj.setChoice(PatternFlowLacpduActorPortPriorityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorPortPriority) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorPortPriority object
func (obj *patternFlowLacpduActorPortPriority) SetValues(value []uint32) PatternFlowLacpduActorPortPriority {
	obj.setChoice(PatternFlowLacpduActorPortPriorityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorPortPriorityCounter
func (obj *patternFlowLacpduActorPortPriority) Increment() PatternFlowLacpduActorPortPriorityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorPortPriorityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorPortPriorityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorPortPriorityCounter
func (obj *patternFlowLacpduActorPortPriority) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorPortPriorityCounter value in the PatternFlowLacpduActorPortPriority object
func (obj *patternFlowLacpduActorPortPriority) SetIncrement(value PatternFlowLacpduActorPortPriorityCounter) PatternFlowLacpduActorPortPriority {
	obj.setChoice(PatternFlowLacpduActorPortPriorityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorPortPriorityCounter
func (obj *patternFlowLacpduActorPortPriority) Decrement() PatternFlowLacpduActorPortPriorityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorPortPriorityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorPortPriorityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorPortPriorityCounter
func (obj *patternFlowLacpduActorPortPriority) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorPortPriorityCounter value in the PatternFlowLacpduActorPortPriority object
func (obj *patternFlowLacpduActorPortPriority) SetDecrement(value PatternFlowLacpduActorPortPriorityCounter) PatternFlowLacpduActorPortPriority {
	obj.setChoice(PatternFlowLacpduActorPortPriorityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorPortPriority) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorPortPriority.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorPortPriority.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpduActorPortPriority) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorPortPriorityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorPortPriorityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorPortPriorityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorPortPriorityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorPortPriorityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorPortPriorityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorPortPriority")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorPortPriority_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorPortPriority_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
