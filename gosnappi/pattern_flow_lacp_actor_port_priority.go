package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorPortPriority *****
type patternFlowLacpActorPortPriority struct {
	validation
	obj             *otg.PatternFlowLacpActorPortPriority
	marshaller      marshalPatternFlowLacpActorPortPriority
	unMarshaller    unMarshalPatternFlowLacpActorPortPriority
	incrementHolder PatternFlowLacpActorPortPriorityCounter
	decrementHolder PatternFlowLacpActorPortPriorityCounter
}

func NewPatternFlowLacpActorPortPriority() PatternFlowLacpActorPortPriority {
	obj := patternFlowLacpActorPortPriority{obj: &otg.PatternFlowLacpActorPortPriority{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorPortPriority) msg() *otg.PatternFlowLacpActorPortPriority {
	return obj.obj
}

func (obj *patternFlowLacpActorPortPriority) setMsg(msg *otg.PatternFlowLacpActorPortPriority) PatternFlowLacpActorPortPriority {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorPortPriority struct {
	obj *patternFlowLacpActorPortPriority
}

type marshalPatternFlowLacpActorPortPriority interface {
	// ToProto marshals PatternFlowLacpActorPortPriority to protobuf object *otg.PatternFlowLacpActorPortPriority
	ToProto() (*otg.PatternFlowLacpActorPortPriority, error)
	// ToPbText marshals PatternFlowLacpActorPortPriority to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorPortPriority to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorPortPriority to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorPortPriority struct {
	obj *patternFlowLacpActorPortPriority
}

type unMarshalPatternFlowLacpActorPortPriority interface {
	// FromProto unmarshals PatternFlowLacpActorPortPriority from protobuf object *otg.PatternFlowLacpActorPortPriority
	FromProto(msg *otg.PatternFlowLacpActorPortPriority) (PatternFlowLacpActorPortPriority, error)
	// FromPbText unmarshals PatternFlowLacpActorPortPriority from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorPortPriority from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorPortPriority from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorPortPriority) Marshal() marshalPatternFlowLacpActorPortPriority {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorPortPriority{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorPortPriority) Unmarshal() unMarshalPatternFlowLacpActorPortPriority {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorPortPriority{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorPortPriority) ToProto() (*otg.PatternFlowLacpActorPortPriority, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorPortPriority) FromProto(msg *otg.PatternFlowLacpActorPortPriority) (PatternFlowLacpActorPortPriority, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorPortPriority) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorPortPriority) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorPortPriority) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorPortPriority) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorPortPriority) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorPortPriority) FromJson(value string) error {
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

func (obj *patternFlowLacpActorPortPriority) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorPortPriority) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorPortPriority) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorPortPriority) Clone() (PatternFlowLacpActorPortPriority, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorPortPriority()
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

func (obj *patternFlowLacpActorPortPriority) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorPortPriority is the priority assigned to the Actor's port. A lower numerical value indicates a higher priority. Used to prioritize ports for inclusion in a Link Aggregation Group (LAG) when the group is full.
type PatternFlowLacpActorPortPriority interface {
	Validation
	// msg marshals PatternFlowLacpActorPortPriority to protobuf object *otg.PatternFlowLacpActorPortPriority
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorPortPriority
	// setMsg unmarshals PatternFlowLacpActorPortPriority from protobuf object *otg.PatternFlowLacpActorPortPriority
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorPortPriority) PatternFlowLacpActorPortPriority
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorPortPriority
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorPortPriority
	// validate validates PatternFlowLacpActorPortPriority
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorPortPriority, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorPortPriorityChoiceEnum, set in PatternFlowLacpActorPortPriority
	Choice() PatternFlowLacpActorPortPriorityChoiceEnum
	// setChoice assigns PatternFlowLacpActorPortPriorityChoiceEnum provided by user to PatternFlowLacpActorPortPriority
	setChoice(value PatternFlowLacpActorPortPriorityChoiceEnum) PatternFlowLacpActorPortPriority
	// HasChoice checks if Choice has been set in PatternFlowLacpActorPortPriority
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorPortPriority.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorPortPriority
	SetValue(value uint32) PatternFlowLacpActorPortPriority
	// HasValue checks if Value has been set in PatternFlowLacpActorPortPriority
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorPortPriority.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorPortPriority
	SetValues(value []uint32) PatternFlowLacpActorPortPriority
	// Increment returns PatternFlowLacpActorPortPriorityCounter, set in PatternFlowLacpActorPortPriority.
	// PatternFlowLacpActorPortPriorityCounter is integer counter pattern
	Increment() PatternFlowLacpActorPortPriorityCounter
	// SetIncrement assigns PatternFlowLacpActorPortPriorityCounter provided by user to PatternFlowLacpActorPortPriority.
	// PatternFlowLacpActorPortPriorityCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorPortPriorityCounter) PatternFlowLacpActorPortPriority
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorPortPriority
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorPortPriorityCounter, set in PatternFlowLacpActorPortPriority.
	// PatternFlowLacpActorPortPriorityCounter is integer counter pattern
	Decrement() PatternFlowLacpActorPortPriorityCounter
	// SetDecrement assigns PatternFlowLacpActorPortPriorityCounter provided by user to PatternFlowLacpActorPortPriority.
	// PatternFlowLacpActorPortPriorityCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorPortPriorityCounter) PatternFlowLacpActorPortPriority
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorPortPriority
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorPortPriorityChoiceEnum string

// Enum of Choice on PatternFlowLacpActorPortPriority
var PatternFlowLacpActorPortPriorityChoice = struct {
	VALUE     PatternFlowLacpActorPortPriorityChoiceEnum
	VALUES    PatternFlowLacpActorPortPriorityChoiceEnum
	INCREMENT PatternFlowLacpActorPortPriorityChoiceEnum
	DECREMENT PatternFlowLacpActorPortPriorityChoiceEnum
}{
	VALUE:     PatternFlowLacpActorPortPriorityChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorPortPriorityChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorPortPriorityChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorPortPriorityChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorPortPriority) Choice() PatternFlowLacpActorPortPriorityChoiceEnum {
	return PatternFlowLacpActorPortPriorityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorPortPriority) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorPortPriority) setChoice(value PatternFlowLacpActorPortPriorityChoiceEnum) PatternFlowLacpActorPortPriority {
	intValue, ok := otg.PatternFlowLacpActorPortPriority_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorPortPriorityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorPortPriority_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorPortPriorityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorPortPriorityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorPortPriorityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorPortPriorityCounter().msg()
	}

	if value == PatternFlowLacpActorPortPriorityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorPortPriorityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorPortPriority) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorPortPriorityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorPortPriority) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorPortPriority object
func (obj *patternFlowLacpActorPortPriority) SetValue(value uint32) PatternFlowLacpActorPortPriority {
	obj.setChoice(PatternFlowLacpActorPortPriorityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorPortPriority) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorPortPriority object
func (obj *patternFlowLacpActorPortPriority) SetValues(value []uint32) PatternFlowLacpActorPortPriority {
	obj.setChoice(PatternFlowLacpActorPortPriorityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorPortPriorityCounter
func (obj *patternFlowLacpActorPortPriority) Increment() PatternFlowLacpActorPortPriorityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorPortPriorityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorPortPriorityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorPortPriorityCounter
func (obj *patternFlowLacpActorPortPriority) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorPortPriorityCounter value in the PatternFlowLacpActorPortPriority object
func (obj *patternFlowLacpActorPortPriority) SetIncrement(value PatternFlowLacpActorPortPriorityCounter) PatternFlowLacpActorPortPriority {
	obj.setChoice(PatternFlowLacpActorPortPriorityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorPortPriorityCounter
func (obj *patternFlowLacpActorPortPriority) Decrement() PatternFlowLacpActorPortPriorityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorPortPriorityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorPortPriorityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorPortPriorityCounter
func (obj *patternFlowLacpActorPortPriority) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorPortPriorityCounter value in the PatternFlowLacpActorPortPriority object
func (obj *patternFlowLacpActorPortPriority) SetDecrement(value PatternFlowLacpActorPortPriorityCounter) PatternFlowLacpActorPortPriority {
	obj.setChoice(PatternFlowLacpActorPortPriorityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorPortPriority) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorPortPriority.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorPortPriority.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpActorPortPriority) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorPortPriorityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorPortPriorityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorPortPriorityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorPortPriorityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorPortPriorityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorPortPriorityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorPortPriority")
			}
		} else {
			intVal := otg.PatternFlowLacpActorPortPriority_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorPortPriority_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
