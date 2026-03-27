package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateDefaulted *****
type patternFlowLacpduActorStateDefaulted struct {
	validation
	obj             *otg.PatternFlowLacpduActorStateDefaulted
	marshaller      marshalPatternFlowLacpduActorStateDefaulted
	unMarshaller    unMarshalPatternFlowLacpduActorStateDefaulted
	incrementHolder PatternFlowLacpduActorStateDefaultedCounter
	decrementHolder PatternFlowLacpduActorStateDefaultedCounter
}

func NewPatternFlowLacpduActorStateDefaulted() PatternFlowLacpduActorStateDefaulted {
	obj := patternFlowLacpduActorStateDefaulted{obj: &otg.PatternFlowLacpduActorStateDefaulted{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateDefaulted) msg() *otg.PatternFlowLacpduActorStateDefaulted {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateDefaulted) setMsg(msg *otg.PatternFlowLacpduActorStateDefaulted) PatternFlowLacpduActorStateDefaulted {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateDefaulted struct {
	obj *patternFlowLacpduActorStateDefaulted
}

type marshalPatternFlowLacpduActorStateDefaulted interface {
	// ToProto marshals PatternFlowLacpduActorStateDefaulted to protobuf object *otg.PatternFlowLacpduActorStateDefaulted
	ToProto() (*otg.PatternFlowLacpduActorStateDefaulted, error)
	// ToPbText marshals PatternFlowLacpduActorStateDefaulted to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateDefaulted to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateDefaulted to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateDefaulted struct {
	obj *patternFlowLacpduActorStateDefaulted
}

type unMarshalPatternFlowLacpduActorStateDefaulted interface {
	// FromProto unmarshals PatternFlowLacpduActorStateDefaulted from protobuf object *otg.PatternFlowLacpduActorStateDefaulted
	FromProto(msg *otg.PatternFlowLacpduActorStateDefaulted) (PatternFlowLacpduActorStateDefaulted, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateDefaulted from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateDefaulted from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateDefaulted from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateDefaulted) Marshal() marshalPatternFlowLacpduActorStateDefaulted {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateDefaulted{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateDefaulted) Unmarshal() unMarshalPatternFlowLacpduActorStateDefaulted {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateDefaulted{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateDefaulted) ToProto() (*otg.PatternFlowLacpduActorStateDefaulted, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateDefaulted) FromProto(msg *otg.PatternFlowLacpduActorStateDefaulted) (PatternFlowLacpduActorStateDefaulted, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateDefaulted) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateDefaulted) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateDefaulted) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateDefaulted) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateDefaulted) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateDefaulted) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateDefaulted) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateDefaulted) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateDefaulted) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateDefaulted) Clone() (PatternFlowLacpduActorStateDefaulted, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateDefaulted()
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

func (obj *patternFlowLacpduActorStateDefaulted) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorStateDefaulted is defaulted (1=Using defaulted partner info, 0=Using received partner info)
type PatternFlowLacpduActorStateDefaulted interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateDefaulted to protobuf object *otg.PatternFlowLacpduActorStateDefaulted
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateDefaulted
	// setMsg unmarshals PatternFlowLacpduActorStateDefaulted from protobuf object *otg.PatternFlowLacpduActorStateDefaulted
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateDefaulted) PatternFlowLacpduActorStateDefaulted
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateDefaulted
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateDefaulted
	// validate validates PatternFlowLacpduActorStateDefaulted
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateDefaulted, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorStateDefaultedChoiceEnum, set in PatternFlowLacpduActorStateDefaulted
	Choice() PatternFlowLacpduActorStateDefaultedChoiceEnum
	// setChoice assigns PatternFlowLacpduActorStateDefaultedChoiceEnum provided by user to PatternFlowLacpduActorStateDefaulted
	setChoice(value PatternFlowLacpduActorStateDefaultedChoiceEnum) PatternFlowLacpduActorStateDefaulted
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorStateDefaulted
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorStateDefaulted.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorStateDefaulted
	SetValue(value uint32) PatternFlowLacpduActorStateDefaulted
	// HasValue checks if Value has been set in PatternFlowLacpduActorStateDefaulted
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorStateDefaulted.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorStateDefaulted
	SetValues(value []uint32) PatternFlowLacpduActorStateDefaulted
	// Increment returns PatternFlowLacpduActorStateDefaultedCounter, set in PatternFlowLacpduActorStateDefaulted.
	// PatternFlowLacpduActorStateDefaultedCounter is integer counter pattern
	Increment() PatternFlowLacpduActorStateDefaultedCounter
	// SetIncrement assigns PatternFlowLacpduActorStateDefaultedCounter provided by user to PatternFlowLacpduActorStateDefaulted.
	// PatternFlowLacpduActorStateDefaultedCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorStateDefaultedCounter) PatternFlowLacpduActorStateDefaulted
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorStateDefaulted
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorStateDefaultedCounter, set in PatternFlowLacpduActorStateDefaulted.
	// PatternFlowLacpduActorStateDefaultedCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorStateDefaultedCounter
	// SetDecrement assigns PatternFlowLacpduActorStateDefaultedCounter provided by user to PatternFlowLacpduActorStateDefaulted.
	// PatternFlowLacpduActorStateDefaultedCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorStateDefaultedCounter) PatternFlowLacpduActorStateDefaulted
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorStateDefaulted
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorStateDefaultedChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorStateDefaulted
var PatternFlowLacpduActorStateDefaultedChoice = struct {
	VALUE     PatternFlowLacpduActorStateDefaultedChoiceEnum
	VALUES    PatternFlowLacpduActorStateDefaultedChoiceEnum
	INCREMENT PatternFlowLacpduActorStateDefaultedChoiceEnum
	DECREMENT PatternFlowLacpduActorStateDefaultedChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorStateDefaultedChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorStateDefaultedChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorStateDefaultedChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorStateDefaultedChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorStateDefaulted) Choice() PatternFlowLacpduActorStateDefaultedChoiceEnum {
	return PatternFlowLacpduActorStateDefaultedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorStateDefaulted) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorStateDefaulted) setChoice(value PatternFlowLacpduActorStateDefaultedChoiceEnum) PatternFlowLacpduActorStateDefaulted {
	intValue, ok := otg.PatternFlowLacpduActorStateDefaulted_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorStateDefaultedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorStateDefaulted_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorStateDefaultedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorStateDefaultedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorStateDefaultedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorStateDefaultedCounter().msg()
	}

	if value == PatternFlowLacpduActorStateDefaultedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorStateDefaultedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateDefaulted) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorStateDefaultedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateDefaulted) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorStateDefaulted object
func (obj *patternFlowLacpduActorStateDefaulted) SetValue(value uint32) PatternFlowLacpduActorStateDefaulted {
	obj.setChoice(PatternFlowLacpduActorStateDefaultedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorStateDefaulted) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorStateDefaulted object
func (obj *patternFlowLacpduActorStateDefaulted) SetValues(value []uint32) PatternFlowLacpduActorStateDefaulted {
	obj.setChoice(PatternFlowLacpduActorStateDefaultedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateDefaultedCounter
func (obj *patternFlowLacpduActorStateDefaulted) Increment() PatternFlowLacpduActorStateDefaultedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorStateDefaultedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorStateDefaultedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateDefaultedCounter
func (obj *patternFlowLacpduActorStateDefaulted) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorStateDefaultedCounter value in the PatternFlowLacpduActorStateDefaulted object
func (obj *patternFlowLacpduActorStateDefaulted) SetIncrement(value PatternFlowLacpduActorStateDefaultedCounter) PatternFlowLacpduActorStateDefaulted {
	obj.setChoice(PatternFlowLacpduActorStateDefaultedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateDefaultedCounter
func (obj *patternFlowLacpduActorStateDefaulted) Decrement() PatternFlowLacpduActorStateDefaultedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorStateDefaultedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorStateDefaultedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateDefaultedCounter
func (obj *patternFlowLacpduActorStateDefaulted) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorStateDefaultedCounter value in the PatternFlowLacpduActorStateDefaulted object
func (obj *patternFlowLacpduActorStateDefaulted) SetDecrement(value PatternFlowLacpduActorStateDefaultedCounter) PatternFlowLacpduActorStateDefaulted {
	obj.setChoice(PatternFlowLacpduActorStateDefaultedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorStateDefaulted) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateDefaulted.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorStateDefaulted.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduActorStateDefaulted) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorStateDefaultedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateDefaultedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorStateDefaultedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateDefaultedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateDefaultedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorStateDefaultedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorStateDefaulted")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorStateDefaulted_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorStateDefaulted_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
