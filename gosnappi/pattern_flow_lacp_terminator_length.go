package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpTerminatorLength *****
type patternFlowLacpTerminatorLength struct {
	validation
	obj             *otg.PatternFlowLacpTerminatorLength
	marshaller      marshalPatternFlowLacpTerminatorLength
	unMarshaller    unMarshalPatternFlowLacpTerminatorLength
	incrementHolder PatternFlowLacpTerminatorLengthCounter
	decrementHolder PatternFlowLacpTerminatorLengthCounter
}

func NewPatternFlowLacpTerminatorLength() PatternFlowLacpTerminatorLength {
	obj := patternFlowLacpTerminatorLength{obj: &otg.PatternFlowLacpTerminatorLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpTerminatorLength) msg() *otg.PatternFlowLacpTerminatorLength {
	return obj.obj
}

func (obj *patternFlowLacpTerminatorLength) setMsg(msg *otg.PatternFlowLacpTerminatorLength) PatternFlowLacpTerminatorLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpTerminatorLength struct {
	obj *patternFlowLacpTerminatorLength
}

type marshalPatternFlowLacpTerminatorLength interface {
	// ToProto marshals PatternFlowLacpTerminatorLength to protobuf object *otg.PatternFlowLacpTerminatorLength
	ToProto() (*otg.PatternFlowLacpTerminatorLength, error)
	// ToPbText marshals PatternFlowLacpTerminatorLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpTerminatorLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpTerminatorLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpTerminatorLength struct {
	obj *patternFlowLacpTerminatorLength
}

type unMarshalPatternFlowLacpTerminatorLength interface {
	// FromProto unmarshals PatternFlowLacpTerminatorLength from protobuf object *otg.PatternFlowLacpTerminatorLength
	FromProto(msg *otg.PatternFlowLacpTerminatorLength) (PatternFlowLacpTerminatorLength, error)
	// FromPbText unmarshals PatternFlowLacpTerminatorLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpTerminatorLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpTerminatorLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpTerminatorLength) Marshal() marshalPatternFlowLacpTerminatorLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpTerminatorLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpTerminatorLength) Unmarshal() unMarshalPatternFlowLacpTerminatorLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpTerminatorLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpTerminatorLength) ToProto() (*otg.PatternFlowLacpTerminatorLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpTerminatorLength) FromProto(msg *otg.PatternFlowLacpTerminatorLength) (PatternFlowLacpTerminatorLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpTerminatorLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpTerminatorLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpTerminatorLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpTerminatorLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpTerminatorLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpTerminatorLength) FromJson(value string) error {
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

func (obj *patternFlowLacpTerminatorLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpTerminatorLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpTerminatorLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpTerminatorLength) Clone() (PatternFlowLacpTerminatorLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpTerminatorLength()
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

func (obj *patternFlowLacpTerminatorLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpTerminatorLength is the length of the Terminator TLV payload. The value must be 0.
type PatternFlowLacpTerminatorLength interface {
	Validation
	// msg marshals PatternFlowLacpTerminatorLength to protobuf object *otg.PatternFlowLacpTerminatorLength
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpTerminatorLength
	// setMsg unmarshals PatternFlowLacpTerminatorLength from protobuf object *otg.PatternFlowLacpTerminatorLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpTerminatorLength) PatternFlowLacpTerminatorLength
	// provides marshal interface
	Marshal() marshalPatternFlowLacpTerminatorLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpTerminatorLength
	// validate validates PatternFlowLacpTerminatorLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpTerminatorLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpTerminatorLengthChoiceEnum, set in PatternFlowLacpTerminatorLength
	Choice() PatternFlowLacpTerminatorLengthChoiceEnum
	// setChoice assigns PatternFlowLacpTerminatorLengthChoiceEnum provided by user to PatternFlowLacpTerminatorLength
	setChoice(value PatternFlowLacpTerminatorLengthChoiceEnum) PatternFlowLacpTerminatorLength
	// HasChoice checks if Choice has been set in PatternFlowLacpTerminatorLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpTerminatorLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpTerminatorLength
	SetValue(value uint32) PatternFlowLacpTerminatorLength
	// HasValue checks if Value has been set in PatternFlowLacpTerminatorLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpTerminatorLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpTerminatorLength
	SetValues(value []uint32) PatternFlowLacpTerminatorLength
	// Increment returns PatternFlowLacpTerminatorLengthCounter, set in PatternFlowLacpTerminatorLength.
	// PatternFlowLacpTerminatorLengthCounter is integer counter pattern
	Increment() PatternFlowLacpTerminatorLengthCounter
	// SetIncrement assigns PatternFlowLacpTerminatorLengthCounter provided by user to PatternFlowLacpTerminatorLength.
	// PatternFlowLacpTerminatorLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpTerminatorLengthCounter) PatternFlowLacpTerminatorLength
	// HasIncrement checks if Increment has been set in PatternFlowLacpTerminatorLength
	HasIncrement() bool
	// Decrement returns PatternFlowLacpTerminatorLengthCounter, set in PatternFlowLacpTerminatorLength.
	// PatternFlowLacpTerminatorLengthCounter is integer counter pattern
	Decrement() PatternFlowLacpTerminatorLengthCounter
	// SetDecrement assigns PatternFlowLacpTerminatorLengthCounter provided by user to PatternFlowLacpTerminatorLength.
	// PatternFlowLacpTerminatorLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpTerminatorLengthCounter) PatternFlowLacpTerminatorLength
	// HasDecrement checks if Decrement has been set in PatternFlowLacpTerminatorLength
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpTerminatorLengthChoiceEnum string

// Enum of Choice on PatternFlowLacpTerminatorLength
var PatternFlowLacpTerminatorLengthChoice = struct {
	VALUE     PatternFlowLacpTerminatorLengthChoiceEnum
	VALUES    PatternFlowLacpTerminatorLengthChoiceEnum
	INCREMENT PatternFlowLacpTerminatorLengthChoiceEnum
	DECREMENT PatternFlowLacpTerminatorLengthChoiceEnum
}{
	VALUE:     PatternFlowLacpTerminatorLengthChoiceEnum("value"),
	VALUES:    PatternFlowLacpTerminatorLengthChoiceEnum("values"),
	INCREMENT: PatternFlowLacpTerminatorLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpTerminatorLengthChoiceEnum("decrement"),
}

func (obj *patternFlowLacpTerminatorLength) Choice() PatternFlowLacpTerminatorLengthChoiceEnum {
	return PatternFlowLacpTerminatorLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpTerminatorLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpTerminatorLength) setChoice(value PatternFlowLacpTerminatorLengthChoiceEnum) PatternFlowLacpTerminatorLength {
	intValue, ok := otg.PatternFlowLacpTerminatorLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpTerminatorLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpTerminatorLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpTerminatorLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpTerminatorLengthChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpTerminatorLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpTerminatorLengthCounter().msg()
	}

	if value == PatternFlowLacpTerminatorLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpTerminatorLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpTerminatorLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpTerminatorLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpTerminatorLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpTerminatorLength object
func (obj *patternFlowLacpTerminatorLength) SetValue(value uint32) PatternFlowLacpTerminatorLength {
	obj.setChoice(PatternFlowLacpTerminatorLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpTerminatorLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpTerminatorLength object
func (obj *patternFlowLacpTerminatorLength) SetValues(value []uint32) PatternFlowLacpTerminatorLength {
	obj.setChoice(PatternFlowLacpTerminatorLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpTerminatorLengthCounter
func (obj *patternFlowLacpTerminatorLength) Increment() PatternFlowLacpTerminatorLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpTerminatorLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpTerminatorLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpTerminatorLengthCounter
func (obj *patternFlowLacpTerminatorLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpTerminatorLengthCounter value in the PatternFlowLacpTerminatorLength object
func (obj *patternFlowLacpTerminatorLength) SetIncrement(value PatternFlowLacpTerminatorLengthCounter) PatternFlowLacpTerminatorLength {
	obj.setChoice(PatternFlowLacpTerminatorLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpTerminatorLengthCounter
func (obj *patternFlowLacpTerminatorLength) Decrement() PatternFlowLacpTerminatorLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpTerminatorLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpTerminatorLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpTerminatorLengthCounter
func (obj *patternFlowLacpTerminatorLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpTerminatorLengthCounter value in the PatternFlowLacpTerminatorLength object
func (obj *patternFlowLacpTerminatorLength) SetDecrement(value PatternFlowLacpTerminatorLengthCounter) PatternFlowLacpTerminatorLength {
	obj.setChoice(PatternFlowLacpTerminatorLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpTerminatorLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpTerminatorLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpTerminatorLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpTerminatorLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpTerminatorLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpTerminatorLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpTerminatorLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpTerminatorLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpTerminatorLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpTerminatorLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpTerminatorLength")
			}
		} else {
			intVal := otg.PatternFlowLacpTerminatorLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpTerminatorLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
