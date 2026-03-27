package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorReserved *****
type patternFlowLacpduActorReserved struct {
	validation
	obj             *otg.PatternFlowLacpduActorReserved
	marshaller      marshalPatternFlowLacpduActorReserved
	unMarshaller    unMarshalPatternFlowLacpduActorReserved
	incrementHolder PatternFlowLacpduActorReservedCounter
	decrementHolder PatternFlowLacpduActorReservedCounter
}

func NewPatternFlowLacpduActorReserved() PatternFlowLacpduActorReserved {
	obj := patternFlowLacpduActorReserved{obj: &otg.PatternFlowLacpduActorReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorReserved) msg() *otg.PatternFlowLacpduActorReserved {
	return obj.obj
}

func (obj *patternFlowLacpduActorReserved) setMsg(msg *otg.PatternFlowLacpduActorReserved) PatternFlowLacpduActorReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorReserved struct {
	obj *patternFlowLacpduActorReserved
}

type marshalPatternFlowLacpduActorReserved interface {
	// ToProto marshals PatternFlowLacpduActorReserved to protobuf object *otg.PatternFlowLacpduActorReserved
	ToProto() (*otg.PatternFlowLacpduActorReserved, error)
	// ToPbText marshals PatternFlowLacpduActorReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorReserved struct {
	obj *patternFlowLacpduActorReserved
}

type unMarshalPatternFlowLacpduActorReserved interface {
	// FromProto unmarshals PatternFlowLacpduActorReserved from protobuf object *otg.PatternFlowLacpduActorReserved
	FromProto(msg *otg.PatternFlowLacpduActorReserved) (PatternFlowLacpduActorReserved, error)
	// FromPbText unmarshals PatternFlowLacpduActorReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorReserved) Marshal() marshalPatternFlowLacpduActorReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorReserved) Unmarshal() unMarshalPatternFlowLacpduActorReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorReserved) ToProto() (*otg.PatternFlowLacpduActorReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorReserved) FromProto(msg *otg.PatternFlowLacpduActorReserved) (PatternFlowLacpduActorReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorReserved) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorReserved) Clone() (PatternFlowLacpduActorReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorReserved()
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

func (obj *patternFlowLacpduActorReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorReserved is reserved field for future use in the Actor TLV. Should be set to all zeros.
type PatternFlowLacpduActorReserved interface {
	Validation
	// msg marshals PatternFlowLacpduActorReserved to protobuf object *otg.PatternFlowLacpduActorReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorReserved
	// setMsg unmarshals PatternFlowLacpduActorReserved from protobuf object *otg.PatternFlowLacpduActorReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorReserved) PatternFlowLacpduActorReserved
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorReserved
	// validate validates PatternFlowLacpduActorReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorReservedChoiceEnum, set in PatternFlowLacpduActorReserved
	Choice() PatternFlowLacpduActorReservedChoiceEnum
	// setChoice assigns PatternFlowLacpduActorReservedChoiceEnum provided by user to PatternFlowLacpduActorReserved
	setChoice(value PatternFlowLacpduActorReservedChoiceEnum) PatternFlowLacpduActorReserved
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorReserved
	SetValue(value uint32) PatternFlowLacpduActorReserved
	// HasValue checks if Value has been set in PatternFlowLacpduActorReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorReserved
	SetValues(value []uint32) PatternFlowLacpduActorReserved
	// Increment returns PatternFlowLacpduActorReservedCounter, set in PatternFlowLacpduActorReserved.
	// PatternFlowLacpduActorReservedCounter is integer counter pattern
	Increment() PatternFlowLacpduActorReservedCounter
	// SetIncrement assigns PatternFlowLacpduActorReservedCounter provided by user to PatternFlowLacpduActorReserved.
	// PatternFlowLacpduActorReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorReservedCounter) PatternFlowLacpduActorReserved
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorReserved
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorReservedCounter, set in PatternFlowLacpduActorReserved.
	// PatternFlowLacpduActorReservedCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorReservedCounter
	// SetDecrement assigns PatternFlowLacpduActorReservedCounter provided by user to PatternFlowLacpduActorReserved.
	// PatternFlowLacpduActorReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorReservedCounter) PatternFlowLacpduActorReserved
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorReservedChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorReserved
var PatternFlowLacpduActorReservedChoice = struct {
	VALUE     PatternFlowLacpduActorReservedChoiceEnum
	VALUES    PatternFlowLacpduActorReservedChoiceEnum
	INCREMENT PatternFlowLacpduActorReservedChoiceEnum
	DECREMENT PatternFlowLacpduActorReservedChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorReservedChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorReservedChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorReservedChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorReserved) Choice() PatternFlowLacpduActorReservedChoiceEnum {
	return PatternFlowLacpduActorReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorReserved) setChoice(value PatternFlowLacpduActorReservedChoiceEnum) PatternFlowLacpduActorReserved {
	intValue, ok := otg.PatternFlowLacpduActorReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorReservedCounter().msg()
	}

	if value == PatternFlowLacpduActorReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorReserved object
func (obj *patternFlowLacpduActorReserved) SetValue(value uint32) PatternFlowLacpduActorReserved {
	obj.setChoice(PatternFlowLacpduActorReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorReserved object
func (obj *patternFlowLacpduActorReserved) SetValues(value []uint32) PatternFlowLacpduActorReserved {
	obj.setChoice(PatternFlowLacpduActorReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorReservedCounter
func (obj *patternFlowLacpduActorReserved) Increment() PatternFlowLacpduActorReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorReservedCounter
func (obj *patternFlowLacpduActorReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorReservedCounter value in the PatternFlowLacpduActorReserved object
func (obj *patternFlowLacpduActorReserved) SetIncrement(value PatternFlowLacpduActorReservedCounter) PatternFlowLacpduActorReserved {
	obj.setChoice(PatternFlowLacpduActorReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorReservedCounter
func (obj *patternFlowLacpduActorReserved) Decrement() PatternFlowLacpduActorReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorReservedCounter
func (obj *patternFlowLacpduActorReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorReservedCounter value in the PatternFlowLacpduActorReserved object
func (obj *patternFlowLacpduActorReserved) SetDecrement(value PatternFlowLacpduActorReservedCounter) PatternFlowLacpduActorReserved {
	obj.setChoice(PatternFlowLacpduActorReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorReserved.Value <= 16777215 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 16777215 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorReserved.Values <= 16777215 but Got %d", item))
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

func (obj *patternFlowLacpduActorReserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorReserved")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorReserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorReserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
