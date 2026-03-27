package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorLength *****
type patternFlowLacpduActorLength struct {
	validation
	obj             *otg.PatternFlowLacpduActorLength
	marshaller      marshalPatternFlowLacpduActorLength
	unMarshaller    unMarshalPatternFlowLacpduActorLength
	incrementHolder PatternFlowLacpduActorLengthCounter
	decrementHolder PatternFlowLacpduActorLengthCounter
}

func NewPatternFlowLacpduActorLength() PatternFlowLacpduActorLength {
	obj := patternFlowLacpduActorLength{obj: &otg.PatternFlowLacpduActorLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorLength) msg() *otg.PatternFlowLacpduActorLength {
	return obj.obj
}

func (obj *patternFlowLacpduActorLength) setMsg(msg *otg.PatternFlowLacpduActorLength) PatternFlowLacpduActorLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorLength struct {
	obj *patternFlowLacpduActorLength
}

type marshalPatternFlowLacpduActorLength interface {
	// ToProto marshals PatternFlowLacpduActorLength to protobuf object *otg.PatternFlowLacpduActorLength
	ToProto() (*otg.PatternFlowLacpduActorLength, error)
	// ToPbText marshals PatternFlowLacpduActorLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorLength struct {
	obj *patternFlowLacpduActorLength
}

type unMarshalPatternFlowLacpduActorLength interface {
	// FromProto unmarshals PatternFlowLacpduActorLength from protobuf object *otg.PatternFlowLacpduActorLength
	FromProto(msg *otg.PatternFlowLacpduActorLength) (PatternFlowLacpduActorLength, error)
	// FromPbText unmarshals PatternFlowLacpduActorLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorLength) Marshal() marshalPatternFlowLacpduActorLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorLength) Unmarshal() unMarshalPatternFlowLacpduActorLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorLength) ToProto() (*otg.PatternFlowLacpduActorLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorLength) FromProto(msg *otg.PatternFlowLacpduActorLength) (PatternFlowLacpduActorLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorLength) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorLength) Clone() (PatternFlowLacpduActorLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorLength()
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

func (obj *patternFlowLacpduActorLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorLength is the length of the Actor Information TLV payload in bytes.
type PatternFlowLacpduActorLength interface {
	Validation
	// msg marshals PatternFlowLacpduActorLength to protobuf object *otg.PatternFlowLacpduActorLength
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorLength
	// setMsg unmarshals PatternFlowLacpduActorLength from protobuf object *otg.PatternFlowLacpduActorLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorLength) PatternFlowLacpduActorLength
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorLength
	// validate validates PatternFlowLacpduActorLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorLengthChoiceEnum, set in PatternFlowLacpduActorLength
	Choice() PatternFlowLacpduActorLengthChoiceEnum
	// setChoice assigns PatternFlowLacpduActorLengthChoiceEnum provided by user to PatternFlowLacpduActorLength
	setChoice(value PatternFlowLacpduActorLengthChoiceEnum) PatternFlowLacpduActorLength
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorLength
	SetValue(value uint32) PatternFlowLacpduActorLength
	// HasValue checks if Value has been set in PatternFlowLacpduActorLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorLength
	SetValues(value []uint32) PatternFlowLacpduActorLength
	// Increment returns PatternFlowLacpduActorLengthCounter, set in PatternFlowLacpduActorLength.
	// PatternFlowLacpduActorLengthCounter is integer counter pattern
	Increment() PatternFlowLacpduActorLengthCounter
	// SetIncrement assigns PatternFlowLacpduActorLengthCounter provided by user to PatternFlowLacpduActorLength.
	// PatternFlowLacpduActorLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorLengthCounter) PatternFlowLacpduActorLength
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorLength
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorLengthCounter, set in PatternFlowLacpduActorLength.
	// PatternFlowLacpduActorLengthCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorLengthCounter
	// SetDecrement assigns PatternFlowLacpduActorLengthCounter provided by user to PatternFlowLacpduActorLength.
	// PatternFlowLacpduActorLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorLengthCounter) PatternFlowLacpduActorLength
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorLength
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorLengthChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorLength
var PatternFlowLacpduActorLengthChoice = struct {
	VALUE     PatternFlowLacpduActorLengthChoiceEnum
	VALUES    PatternFlowLacpduActorLengthChoiceEnum
	INCREMENT PatternFlowLacpduActorLengthChoiceEnum
	DECREMENT PatternFlowLacpduActorLengthChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorLengthChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorLengthChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorLengthChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorLength) Choice() PatternFlowLacpduActorLengthChoiceEnum {
	return PatternFlowLacpduActorLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorLength) setChoice(value PatternFlowLacpduActorLengthChoiceEnum) PatternFlowLacpduActorLength {
	intValue, ok := otg.PatternFlowLacpduActorLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorLengthChoice.VALUE {
		defaultValue := uint32(20)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorLengthChoice.VALUES {
		defaultValue := []uint32{20}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorLengthCounter().msg()
	}

	if value == PatternFlowLacpduActorLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorLength object
func (obj *patternFlowLacpduActorLength) SetValue(value uint32) PatternFlowLacpduActorLength {
	obj.setChoice(PatternFlowLacpduActorLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{20})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorLength object
func (obj *patternFlowLacpduActorLength) SetValues(value []uint32) PatternFlowLacpduActorLength {
	obj.setChoice(PatternFlowLacpduActorLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorLengthCounter
func (obj *patternFlowLacpduActorLength) Increment() PatternFlowLacpduActorLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorLengthCounter
func (obj *patternFlowLacpduActorLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorLengthCounter value in the PatternFlowLacpduActorLength object
func (obj *patternFlowLacpduActorLength) SetIncrement(value PatternFlowLacpduActorLengthCounter) PatternFlowLacpduActorLength {
	obj.setChoice(PatternFlowLacpduActorLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorLengthCounter
func (obj *patternFlowLacpduActorLength) Decrement() PatternFlowLacpduActorLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorLengthCounter
func (obj *patternFlowLacpduActorLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorLengthCounter value in the PatternFlowLacpduActorLength object
func (obj *patternFlowLacpduActorLength) SetDecrement(value PatternFlowLacpduActorLengthCounter) PatternFlowLacpduActorLength {
	obj.setChoice(PatternFlowLacpduActorLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpduActorLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorLength")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
