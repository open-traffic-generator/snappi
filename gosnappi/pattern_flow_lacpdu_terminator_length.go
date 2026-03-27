package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduTerminatorLength *****
type patternFlowLacpduTerminatorLength struct {
	validation
	obj             *otg.PatternFlowLacpduTerminatorLength
	marshaller      marshalPatternFlowLacpduTerminatorLength
	unMarshaller    unMarshalPatternFlowLacpduTerminatorLength
	incrementHolder PatternFlowLacpduTerminatorLengthCounter
	decrementHolder PatternFlowLacpduTerminatorLengthCounter
}

func NewPatternFlowLacpduTerminatorLength() PatternFlowLacpduTerminatorLength {
	obj := patternFlowLacpduTerminatorLength{obj: &otg.PatternFlowLacpduTerminatorLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduTerminatorLength) msg() *otg.PatternFlowLacpduTerminatorLength {
	return obj.obj
}

func (obj *patternFlowLacpduTerminatorLength) setMsg(msg *otg.PatternFlowLacpduTerminatorLength) PatternFlowLacpduTerminatorLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduTerminatorLength struct {
	obj *patternFlowLacpduTerminatorLength
}

type marshalPatternFlowLacpduTerminatorLength interface {
	// ToProto marshals PatternFlowLacpduTerminatorLength to protobuf object *otg.PatternFlowLacpduTerminatorLength
	ToProto() (*otg.PatternFlowLacpduTerminatorLength, error)
	// ToPbText marshals PatternFlowLacpduTerminatorLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduTerminatorLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduTerminatorLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduTerminatorLength struct {
	obj *patternFlowLacpduTerminatorLength
}

type unMarshalPatternFlowLacpduTerminatorLength interface {
	// FromProto unmarshals PatternFlowLacpduTerminatorLength from protobuf object *otg.PatternFlowLacpduTerminatorLength
	FromProto(msg *otg.PatternFlowLacpduTerminatorLength) (PatternFlowLacpduTerminatorLength, error)
	// FromPbText unmarshals PatternFlowLacpduTerminatorLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduTerminatorLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduTerminatorLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduTerminatorLength) Marshal() marshalPatternFlowLacpduTerminatorLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduTerminatorLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduTerminatorLength) Unmarshal() unMarshalPatternFlowLacpduTerminatorLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduTerminatorLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduTerminatorLength) ToProto() (*otg.PatternFlowLacpduTerminatorLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduTerminatorLength) FromProto(msg *otg.PatternFlowLacpduTerminatorLength) (PatternFlowLacpduTerminatorLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduTerminatorLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduTerminatorLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduTerminatorLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduTerminatorLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduTerminatorLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduTerminatorLength) FromJson(value string) error {
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

func (obj *patternFlowLacpduTerminatorLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduTerminatorLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduTerminatorLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduTerminatorLength) Clone() (PatternFlowLacpduTerminatorLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduTerminatorLength()
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

func (obj *patternFlowLacpduTerminatorLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduTerminatorLength is the length of the Terminator TLV payload. The value must be 0.
type PatternFlowLacpduTerminatorLength interface {
	Validation
	// msg marshals PatternFlowLacpduTerminatorLength to protobuf object *otg.PatternFlowLacpduTerminatorLength
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduTerminatorLength
	// setMsg unmarshals PatternFlowLacpduTerminatorLength from protobuf object *otg.PatternFlowLacpduTerminatorLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduTerminatorLength) PatternFlowLacpduTerminatorLength
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduTerminatorLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduTerminatorLength
	// validate validates PatternFlowLacpduTerminatorLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduTerminatorLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduTerminatorLengthChoiceEnum, set in PatternFlowLacpduTerminatorLength
	Choice() PatternFlowLacpduTerminatorLengthChoiceEnum
	// setChoice assigns PatternFlowLacpduTerminatorLengthChoiceEnum provided by user to PatternFlowLacpduTerminatorLength
	setChoice(value PatternFlowLacpduTerminatorLengthChoiceEnum) PatternFlowLacpduTerminatorLength
	// HasChoice checks if Choice has been set in PatternFlowLacpduTerminatorLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduTerminatorLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduTerminatorLength
	SetValue(value uint32) PatternFlowLacpduTerminatorLength
	// HasValue checks if Value has been set in PatternFlowLacpduTerminatorLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduTerminatorLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduTerminatorLength
	SetValues(value []uint32) PatternFlowLacpduTerminatorLength
	// Increment returns PatternFlowLacpduTerminatorLengthCounter, set in PatternFlowLacpduTerminatorLength.
	// PatternFlowLacpduTerminatorLengthCounter is integer counter pattern
	Increment() PatternFlowLacpduTerminatorLengthCounter
	// SetIncrement assigns PatternFlowLacpduTerminatorLengthCounter provided by user to PatternFlowLacpduTerminatorLength.
	// PatternFlowLacpduTerminatorLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduTerminatorLengthCounter) PatternFlowLacpduTerminatorLength
	// HasIncrement checks if Increment has been set in PatternFlowLacpduTerminatorLength
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduTerminatorLengthCounter, set in PatternFlowLacpduTerminatorLength.
	// PatternFlowLacpduTerminatorLengthCounter is integer counter pattern
	Decrement() PatternFlowLacpduTerminatorLengthCounter
	// SetDecrement assigns PatternFlowLacpduTerminatorLengthCounter provided by user to PatternFlowLacpduTerminatorLength.
	// PatternFlowLacpduTerminatorLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduTerminatorLengthCounter) PatternFlowLacpduTerminatorLength
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduTerminatorLength
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduTerminatorLengthChoiceEnum string

// Enum of Choice on PatternFlowLacpduTerminatorLength
var PatternFlowLacpduTerminatorLengthChoice = struct {
	VALUE     PatternFlowLacpduTerminatorLengthChoiceEnum
	VALUES    PatternFlowLacpduTerminatorLengthChoiceEnum
	INCREMENT PatternFlowLacpduTerminatorLengthChoiceEnum
	DECREMENT PatternFlowLacpduTerminatorLengthChoiceEnum
}{
	VALUE:     PatternFlowLacpduTerminatorLengthChoiceEnum("value"),
	VALUES:    PatternFlowLacpduTerminatorLengthChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduTerminatorLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduTerminatorLengthChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduTerminatorLength) Choice() PatternFlowLacpduTerminatorLengthChoiceEnum {
	return PatternFlowLacpduTerminatorLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduTerminatorLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduTerminatorLength) setChoice(value PatternFlowLacpduTerminatorLengthChoiceEnum) PatternFlowLacpduTerminatorLength {
	intValue, ok := otg.PatternFlowLacpduTerminatorLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduTerminatorLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduTerminatorLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduTerminatorLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduTerminatorLengthChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduTerminatorLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduTerminatorLengthCounter().msg()
	}

	if value == PatternFlowLacpduTerminatorLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduTerminatorLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduTerminatorLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduTerminatorLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduTerminatorLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduTerminatorLength object
func (obj *patternFlowLacpduTerminatorLength) SetValue(value uint32) PatternFlowLacpduTerminatorLength {
	obj.setChoice(PatternFlowLacpduTerminatorLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduTerminatorLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduTerminatorLength object
func (obj *patternFlowLacpduTerminatorLength) SetValues(value []uint32) PatternFlowLacpduTerminatorLength {
	obj.setChoice(PatternFlowLacpduTerminatorLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduTerminatorLengthCounter
func (obj *patternFlowLacpduTerminatorLength) Increment() PatternFlowLacpduTerminatorLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduTerminatorLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduTerminatorLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduTerminatorLengthCounter
func (obj *patternFlowLacpduTerminatorLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduTerminatorLengthCounter value in the PatternFlowLacpduTerminatorLength object
func (obj *patternFlowLacpduTerminatorLength) SetIncrement(value PatternFlowLacpduTerminatorLengthCounter) PatternFlowLacpduTerminatorLength {
	obj.setChoice(PatternFlowLacpduTerminatorLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduTerminatorLengthCounter
func (obj *patternFlowLacpduTerminatorLength) Decrement() PatternFlowLacpduTerminatorLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduTerminatorLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduTerminatorLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduTerminatorLengthCounter
func (obj *patternFlowLacpduTerminatorLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduTerminatorLengthCounter value in the PatternFlowLacpduTerminatorLength object
func (obj *patternFlowLacpduTerminatorLength) SetDecrement(value PatternFlowLacpduTerminatorLengthCounter) PatternFlowLacpduTerminatorLength {
	obj.setChoice(PatternFlowLacpduTerminatorLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduTerminatorLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduTerminatorLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduTerminatorLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpduTerminatorLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduTerminatorLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduTerminatorLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduTerminatorLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduTerminatorLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduTerminatorLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduTerminatorLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduTerminatorLength")
			}
		} else {
			intVal := otg.PatternFlowLacpduTerminatorLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduTerminatorLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
