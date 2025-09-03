package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorLength *****
type patternFlowLacpActorLength struct {
	validation
	obj             *otg.PatternFlowLacpActorLength
	marshaller      marshalPatternFlowLacpActorLength
	unMarshaller    unMarshalPatternFlowLacpActorLength
	incrementHolder PatternFlowLacpActorLengthCounter
	decrementHolder PatternFlowLacpActorLengthCounter
}

func NewPatternFlowLacpActorLength() PatternFlowLacpActorLength {
	obj := patternFlowLacpActorLength{obj: &otg.PatternFlowLacpActorLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorLength) msg() *otg.PatternFlowLacpActorLength {
	return obj.obj
}

func (obj *patternFlowLacpActorLength) setMsg(msg *otg.PatternFlowLacpActorLength) PatternFlowLacpActorLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorLength struct {
	obj *patternFlowLacpActorLength
}

type marshalPatternFlowLacpActorLength interface {
	// ToProto marshals PatternFlowLacpActorLength to protobuf object *otg.PatternFlowLacpActorLength
	ToProto() (*otg.PatternFlowLacpActorLength, error)
	// ToPbText marshals PatternFlowLacpActorLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorLength struct {
	obj *patternFlowLacpActorLength
}

type unMarshalPatternFlowLacpActorLength interface {
	// FromProto unmarshals PatternFlowLacpActorLength from protobuf object *otg.PatternFlowLacpActorLength
	FromProto(msg *otg.PatternFlowLacpActorLength) (PatternFlowLacpActorLength, error)
	// FromPbText unmarshals PatternFlowLacpActorLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorLength) Marshal() marshalPatternFlowLacpActorLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorLength) Unmarshal() unMarshalPatternFlowLacpActorLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorLength) ToProto() (*otg.PatternFlowLacpActorLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorLength) FromProto(msg *otg.PatternFlowLacpActorLength) (PatternFlowLacpActorLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorLength) FromJson(value string) error {
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

func (obj *patternFlowLacpActorLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorLength) Clone() (PatternFlowLacpActorLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorLength()
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

func (obj *patternFlowLacpActorLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorLength is the length of the Actor Information TLV payload in bytes.
type PatternFlowLacpActorLength interface {
	Validation
	// msg marshals PatternFlowLacpActorLength to protobuf object *otg.PatternFlowLacpActorLength
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorLength
	// setMsg unmarshals PatternFlowLacpActorLength from protobuf object *otg.PatternFlowLacpActorLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorLength) PatternFlowLacpActorLength
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorLength
	// validate validates PatternFlowLacpActorLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorLengthChoiceEnum, set in PatternFlowLacpActorLength
	Choice() PatternFlowLacpActorLengthChoiceEnum
	// setChoice assigns PatternFlowLacpActorLengthChoiceEnum provided by user to PatternFlowLacpActorLength
	setChoice(value PatternFlowLacpActorLengthChoiceEnum) PatternFlowLacpActorLength
	// HasChoice checks if Choice has been set in PatternFlowLacpActorLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorLength
	SetValue(value uint32) PatternFlowLacpActorLength
	// HasValue checks if Value has been set in PatternFlowLacpActorLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorLength
	SetValues(value []uint32) PatternFlowLacpActorLength
	// Increment returns PatternFlowLacpActorLengthCounter, set in PatternFlowLacpActorLength.
	// PatternFlowLacpActorLengthCounter is integer counter pattern
	Increment() PatternFlowLacpActorLengthCounter
	// SetIncrement assigns PatternFlowLacpActorLengthCounter provided by user to PatternFlowLacpActorLength.
	// PatternFlowLacpActorLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorLengthCounter) PatternFlowLacpActorLength
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorLength
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorLengthCounter, set in PatternFlowLacpActorLength.
	// PatternFlowLacpActorLengthCounter is integer counter pattern
	Decrement() PatternFlowLacpActorLengthCounter
	// SetDecrement assigns PatternFlowLacpActorLengthCounter provided by user to PatternFlowLacpActorLength.
	// PatternFlowLacpActorLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorLengthCounter) PatternFlowLacpActorLength
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorLength
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorLengthChoiceEnum string

// Enum of Choice on PatternFlowLacpActorLength
var PatternFlowLacpActorLengthChoice = struct {
	VALUE     PatternFlowLacpActorLengthChoiceEnum
	VALUES    PatternFlowLacpActorLengthChoiceEnum
	INCREMENT PatternFlowLacpActorLengthChoiceEnum
	DECREMENT PatternFlowLacpActorLengthChoiceEnum
}{
	VALUE:     PatternFlowLacpActorLengthChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorLengthChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorLengthChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorLength) Choice() PatternFlowLacpActorLengthChoiceEnum {
	return PatternFlowLacpActorLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorLength) setChoice(value PatternFlowLacpActorLengthChoiceEnum) PatternFlowLacpActorLength {
	intValue, ok := otg.PatternFlowLacpActorLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorLengthChoice.VALUE {
		defaultValue := uint32(20)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorLengthChoice.VALUES {
		defaultValue := []uint32{20}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorLengthCounter().msg()
	}

	if value == PatternFlowLacpActorLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorLength object
func (obj *patternFlowLacpActorLength) SetValue(value uint32) PatternFlowLacpActorLength {
	obj.setChoice(PatternFlowLacpActorLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{20})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorLength object
func (obj *patternFlowLacpActorLength) SetValues(value []uint32) PatternFlowLacpActorLength {
	obj.setChoice(PatternFlowLacpActorLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorLengthCounter
func (obj *patternFlowLacpActorLength) Increment() PatternFlowLacpActorLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorLengthCounter
func (obj *patternFlowLacpActorLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorLengthCounter value in the PatternFlowLacpActorLength object
func (obj *patternFlowLacpActorLength) SetIncrement(value PatternFlowLacpActorLengthCounter) PatternFlowLacpActorLength {
	obj.setChoice(PatternFlowLacpActorLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorLengthCounter
func (obj *patternFlowLacpActorLength) Decrement() PatternFlowLacpActorLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorLengthCounter
func (obj *patternFlowLacpActorLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorLengthCounter value in the PatternFlowLacpActorLength object
func (obj *patternFlowLacpActorLength) SetDecrement(value PatternFlowLacpActorLengthCounter) PatternFlowLacpActorLength {
	obj.setChoice(PatternFlowLacpActorLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpActorLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorLength")
			}
		} else {
			intVal := otg.PatternFlowLacpActorLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
