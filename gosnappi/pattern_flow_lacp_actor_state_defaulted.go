package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateDefaulted *****
type patternFlowLacpActorStateDefaulted struct {
	validation
	obj             *otg.PatternFlowLacpActorStateDefaulted
	marshaller      marshalPatternFlowLacpActorStateDefaulted
	unMarshaller    unMarshalPatternFlowLacpActorStateDefaulted
	incrementHolder PatternFlowLacpActorStateDefaultedCounter
	decrementHolder PatternFlowLacpActorStateDefaultedCounter
}

func NewPatternFlowLacpActorStateDefaulted() PatternFlowLacpActorStateDefaulted {
	obj := patternFlowLacpActorStateDefaulted{obj: &otg.PatternFlowLacpActorStateDefaulted{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateDefaulted) msg() *otg.PatternFlowLacpActorStateDefaulted {
	return obj.obj
}

func (obj *patternFlowLacpActorStateDefaulted) setMsg(msg *otg.PatternFlowLacpActorStateDefaulted) PatternFlowLacpActorStateDefaulted {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateDefaulted struct {
	obj *patternFlowLacpActorStateDefaulted
}

type marshalPatternFlowLacpActorStateDefaulted interface {
	// ToProto marshals PatternFlowLacpActorStateDefaulted to protobuf object *otg.PatternFlowLacpActorStateDefaulted
	ToProto() (*otg.PatternFlowLacpActorStateDefaulted, error)
	// ToPbText marshals PatternFlowLacpActorStateDefaulted to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateDefaulted to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateDefaulted to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateDefaulted struct {
	obj *patternFlowLacpActorStateDefaulted
}

type unMarshalPatternFlowLacpActorStateDefaulted interface {
	// FromProto unmarshals PatternFlowLacpActorStateDefaulted from protobuf object *otg.PatternFlowLacpActorStateDefaulted
	FromProto(msg *otg.PatternFlowLacpActorStateDefaulted) (PatternFlowLacpActorStateDefaulted, error)
	// FromPbText unmarshals PatternFlowLacpActorStateDefaulted from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateDefaulted from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateDefaulted from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateDefaulted) Marshal() marshalPatternFlowLacpActorStateDefaulted {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateDefaulted{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateDefaulted) Unmarshal() unMarshalPatternFlowLacpActorStateDefaulted {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateDefaulted{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateDefaulted) ToProto() (*otg.PatternFlowLacpActorStateDefaulted, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateDefaulted) FromProto(msg *otg.PatternFlowLacpActorStateDefaulted) (PatternFlowLacpActorStateDefaulted, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateDefaulted) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateDefaulted) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateDefaulted) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateDefaulted) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateDefaulted) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateDefaulted) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateDefaulted) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateDefaulted) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateDefaulted) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateDefaulted) Clone() (PatternFlowLacpActorStateDefaulted, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateDefaulted()
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

func (obj *patternFlowLacpActorStateDefaulted) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorStateDefaulted is defaulted (1=Using defaulted partner info, 0=Using received partner info)
type PatternFlowLacpActorStateDefaulted interface {
	Validation
	// msg marshals PatternFlowLacpActorStateDefaulted to protobuf object *otg.PatternFlowLacpActorStateDefaulted
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateDefaulted
	// setMsg unmarshals PatternFlowLacpActorStateDefaulted from protobuf object *otg.PatternFlowLacpActorStateDefaulted
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateDefaulted) PatternFlowLacpActorStateDefaulted
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateDefaulted
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateDefaulted
	// validate validates PatternFlowLacpActorStateDefaulted
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateDefaulted, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorStateDefaultedChoiceEnum, set in PatternFlowLacpActorStateDefaulted
	Choice() PatternFlowLacpActorStateDefaultedChoiceEnum
	// setChoice assigns PatternFlowLacpActorStateDefaultedChoiceEnum provided by user to PatternFlowLacpActorStateDefaulted
	setChoice(value PatternFlowLacpActorStateDefaultedChoiceEnum) PatternFlowLacpActorStateDefaulted
	// HasChoice checks if Choice has been set in PatternFlowLacpActorStateDefaulted
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorStateDefaulted.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorStateDefaulted
	SetValue(value uint32) PatternFlowLacpActorStateDefaulted
	// HasValue checks if Value has been set in PatternFlowLacpActorStateDefaulted
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorStateDefaulted.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorStateDefaulted
	SetValues(value []uint32) PatternFlowLacpActorStateDefaulted
	// Increment returns PatternFlowLacpActorStateDefaultedCounter, set in PatternFlowLacpActorStateDefaulted.
	// PatternFlowLacpActorStateDefaultedCounter is integer counter pattern
	Increment() PatternFlowLacpActorStateDefaultedCounter
	// SetIncrement assigns PatternFlowLacpActorStateDefaultedCounter provided by user to PatternFlowLacpActorStateDefaulted.
	// PatternFlowLacpActorStateDefaultedCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorStateDefaultedCounter) PatternFlowLacpActorStateDefaulted
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorStateDefaulted
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorStateDefaultedCounter, set in PatternFlowLacpActorStateDefaulted.
	// PatternFlowLacpActorStateDefaultedCounter is integer counter pattern
	Decrement() PatternFlowLacpActorStateDefaultedCounter
	// SetDecrement assigns PatternFlowLacpActorStateDefaultedCounter provided by user to PatternFlowLacpActorStateDefaulted.
	// PatternFlowLacpActorStateDefaultedCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorStateDefaultedCounter) PatternFlowLacpActorStateDefaulted
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorStateDefaulted
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorStateDefaultedChoiceEnum string

// Enum of Choice on PatternFlowLacpActorStateDefaulted
var PatternFlowLacpActorStateDefaultedChoice = struct {
	VALUE     PatternFlowLacpActorStateDefaultedChoiceEnum
	VALUES    PatternFlowLacpActorStateDefaultedChoiceEnum
	INCREMENT PatternFlowLacpActorStateDefaultedChoiceEnum
	DECREMENT PatternFlowLacpActorStateDefaultedChoiceEnum
}{
	VALUE:     PatternFlowLacpActorStateDefaultedChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorStateDefaultedChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorStateDefaultedChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorStateDefaultedChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorStateDefaulted) Choice() PatternFlowLacpActorStateDefaultedChoiceEnum {
	return PatternFlowLacpActorStateDefaultedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorStateDefaulted) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorStateDefaulted) setChoice(value PatternFlowLacpActorStateDefaultedChoiceEnum) PatternFlowLacpActorStateDefaulted {
	intValue, ok := otg.PatternFlowLacpActorStateDefaulted_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorStateDefaultedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorStateDefaulted_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorStateDefaultedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorStateDefaultedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorStateDefaultedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorStateDefaultedCounter().msg()
	}

	if value == PatternFlowLacpActorStateDefaultedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorStateDefaultedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateDefaulted) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorStateDefaultedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateDefaulted) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorStateDefaulted object
func (obj *patternFlowLacpActorStateDefaulted) SetValue(value uint32) PatternFlowLacpActorStateDefaulted {
	obj.setChoice(PatternFlowLacpActorStateDefaultedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorStateDefaulted) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorStateDefaulted object
func (obj *patternFlowLacpActorStateDefaulted) SetValues(value []uint32) PatternFlowLacpActorStateDefaulted {
	obj.setChoice(PatternFlowLacpActorStateDefaultedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateDefaultedCounter
func (obj *patternFlowLacpActorStateDefaulted) Increment() PatternFlowLacpActorStateDefaultedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorStateDefaultedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorStateDefaultedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateDefaultedCounter
func (obj *patternFlowLacpActorStateDefaulted) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorStateDefaultedCounter value in the PatternFlowLacpActorStateDefaulted object
func (obj *patternFlowLacpActorStateDefaulted) SetIncrement(value PatternFlowLacpActorStateDefaultedCounter) PatternFlowLacpActorStateDefaulted {
	obj.setChoice(PatternFlowLacpActorStateDefaultedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateDefaultedCounter
func (obj *patternFlowLacpActorStateDefaulted) Decrement() PatternFlowLacpActorStateDefaultedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorStateDefaultedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorStateDefaultedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateDefaultedCounter
func (obj *patternFlowLacpActorStateDefaulted) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorStateDefaultedCounter value in the PatternFlowLacpActorStateDefaulted object
func (obj *patternFlowLacpActorStateDefaulted) SetDecrement(value PatternFlowLacpActorStateDefaultedCounter) PatternFlowLacpActorStateDefaulted {
	obj.setChoice(PatternFlowLacpActorStateDefaultedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorStateDefaulted) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateDefaulted.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorStateDefaulted.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpActorStateDefaulted) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorStateDefaultedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateDefaultedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorStateDefaultedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateDefaultedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateDefaultedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorStateDefaultedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorStateDefaulted")
			}
		} else {
			intVal := otg.PatternFlowLacpActorStateDefaulted_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorStateDefaulted_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
