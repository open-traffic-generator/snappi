package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmCfmShortMaNameLength *****
type patternFlowCfmCfmShortMaNameLength struct {
	validation
	obj             *otg.PatternFlowCfmCfmShortMaNameLength
	marshaller      marshalPatternFlowCfmCfmShortMaNameLength
	unMarshaller    unMarshalPatternFlowCfmCfmShortMaNameLength
	incrementHolder PatternFlowCfmCfmShortMaNameLengthCounter
	decrementHolder PatternFlowCfmCfmShortMaNameLengthCounter
}

func NewPatternFlowCfmCfmShortMaNameLength() PatternFlowCfmCfmShortMaNameLength {
	obj := patternFlowCfmCfmShortMaNameLength{obj: &otg.PatternFlowCfmCfmShortMaNameLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmCfmShortMaNameLength) msg() *otg.PatternFlowCfmCfmShortMaNameLength {
	return obj.obj
}

func (obj *patternFlowCfmCfmShortMaNameLength) setMsg(msg *otg.PatternFlowCfmCfmShortMaNameLength) PatternFlowCfmCfmShortMaNameLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmCfmShortMaNameLength struct {
	obj *patternFlowCfmCfmShortMaNameLength
}

type marshalPatternFlowCfmCfmShortMaNameLength interface {
	// ToProto marshals PatternFlowCfmCfmShortMaNameLength to protobuf object *otg.PatternFlowCfmCfmShortMaNameLength
	ToProto() (*otg.PatternFlowCfmCfmShortMaNameLength, error)
	// ToPbText marshals PatternFlowCfmCfmShortMaNameLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmCfmShortMaNameLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmCfmShortMaNameLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmCfmShortMaNameLength struct {
	obj *patternFlowCfmCfmShortMaNameLength
}

type unMarshalPatternFlowCfmCfmShortMaNameLength interface {
	// FromProto unmarshals PatternFlowCfmCfmShortMaNameLength from protobuf object *otg.PatternFlowCfmCfmShortMaNameLength
	FromProto(msg *otg.PatternFlowCfmCfmShortMaNameLength) (PatternFlowCfmCfmShortMaNameLength, error)
	// FromPbText unmarshals PatternFlowCfmCfmShortMaNameLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmCfmShortMaNameLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmCfmShortMaNameLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmCfmShortMaNameLength) Marshal() marshalPatternFlowCfmCfmShortMaNameLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmCfmShortMaNameLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmCfmShortMaNameLength) Unmarshal() unMarshalPatternFlowCfmCfmShortMaNameLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmCfmShortMaNameLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmCfmShortMaNameLength) ToProto() (*otg.PatternFlowCfmCfmShortMaNameLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmCfmShortMaNameLength) FromProto(msg *otg.PatternFlowCfmCfmShortMaNameLength) (PatternFlowCfmCfmShortMaNameLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmCfmShortMaNameLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmCfmShortMaNameLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmCfmShortMaNameLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmCfmShortMaNameLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmCfmShortMaNameLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmCfmShortMaNameLength) FromJson(value string) error {
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

func (obj *patternFlowCfmCfmShortMaNameLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmCfmShortMaNameLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmCfmShortMaNameLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmCfmShortMaNameLength) Clone() (PatternFlowCfmCfmShortMaNameLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmCfmShortMaNameLength()
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

func (obj *patternFlowCfmCfmShortMaNameLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmCfmShortMaNameLength is short MA name length
type PatternFlowCfmCfmShortMaNameLength interface {
	Validation
	// msg marshals PatternFlowCfmCfmShortMaNameLength to protobuf object *otg.PatternFlowCfmCfmShortMaNameLength
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmCfmShortMaNameLength
	// setMsg unmarshals PatternFlowCfmCfmShortMaNameLength from protobuf object *otg.PatternFlowCfmCfmShortMaNameLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmCfmShortMaNameLength) PatternFlowCfmCfmShortMaNameLength
	// provides marshal interface
	Marshal() marshalPatternFlowCfmCfmShortMaNameLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmCfmShortMaNameLength
	// validate validates PatternFlowCfmCfmShortMaNameLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmCfmShortMaNameLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmCfmShortMaNameLengthChoiceEnum, set in PatternFlowCfmCfmShortMaNameLength
	Choice() PatternFlowCfmCfmShortMaNameLengthChoiceEnum
	// setChoice assigns PatternFlowCfmCfmShortMaNameLengthChoiceEnum provided by user to PatternFlowCfmCfmShortMaNameLength
	setChoice(value PatternFlowCfmCfmShortMaNameLengthChoiceEnum) PatternFlowCfmCfmShortMaNameLength
	// HasChoice checks if Choice has been set in PatternFlowCfmCfmShortMaNameLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowCfmCfmShortMaNameLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowCfmCfmShortMaNameLength
	SetValue(value uint32) PatternFlowCfmCfmShortMaNameLength
	// HasValue checks if Value has been set in PatternFlowCfmCfmShortMaNameLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowCfmCfmShortMaNameLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowCfmCfmShortMaNameLength
	SetValues(value []uint32) PatternFlowCfmCfmShortMaNameLength
	// Increment returns PatternFlowCfmCfmShortMaNameLengthCounter, set in PatternFlowCfmCfmShortMaNameLength.
	// PatternFlowCfmCfmShortMaNameLengthCounter is integer counter pattern
	Increment() PatternFlowCfmCfmShortMaNameLengthCounter
	// SetIncrement assigns PatternFlowCfmCfmShortMaNameLengthCounter provided by user to PatternFlowCfmCfmShortMaNameLength.
	// PatternFlowCfmCfmShortMaNameLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmCfmShortMaNameLengthCounter) PatternFlowCfmCfmShortMaNameLength
	// HasIncrement checks if Increment has been set in PatternFlowCfmCfmShortMaNameLength
	HasIncrement() bool
	// Decrement returns PatternFlowCfmCfmShortMaNameLengthCounter, set in PatternFlowCfmCfmShortMaNameLength.
	// PatternFlowCfmCfmShortMaNameLengthCounter is integer counter pattern
	Decrement() PatternFlowCfmCfmShortMaNameLengthCounter
	// SetDecrement assigns PatternFlowCfmCfmShortMaNameLengthCounter provided by user to PatternFlowCfmCfmShortMaNameLength.
	// PatternFlowCfmCfmShortMaNameLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmCfmShortMaNameLengthCounter) PatternFlowCfmCfmShortMaNameLength
	// HasDecrement checks if Decrement has been set in PatternFlowCfmCfmShortMaNameLength
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmCfmShortMaNameLengthChoiceEnum string

// Enum of Choice on PatternFlowCfmCfmShortMaNameLength
var PatternFlowCfmCfmShortMaNameLengthChoice = struct {
	VALUE     PatternFlowCfmCfmShortMaNameLengthChoiceEnum
	VALUES    PatternFlowCfmCfmShortMaNameLengthChoiceEnum
	INCREMENT PatternFlowCfmCfmShortMaNameLengthChoiceEnum
	DECREMENT PatternFlowCfmCfmShortMaNameLengthChoiceEnum
}{
	VALUE:     PatternFlowCfmCfmShortMaNameLengthChoiceEnum("value"),
	VALUES:    PatternFlowCfmCfmShortMaNameLengthChoiceEnum("values"),
	INCREMENT: PatternFlowCfmCfmShortMaNameLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmCfmShortMaNameLengthChoiceEnum("decrement"),
}

func (obj *patternFlowCfmCfmShortMaNameLength) Choice() PatternFlowCfmCfmShortMaNameLengthChoiceEnum {
	return PatternFlowCfmCfmShortMaNameLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmCfmShortMaNameLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmCfmShortMaNameLength) setChoice(value PatternFlowCfmCfmShortMaNameLengthChoiceEnum) PatternFlowCfmCfmShortMaNameLength {
	intValue, ok := otg.PatternFlowCfmCfmShortMaNameLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmCfmShortMaNameLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmCfmShortMaNameLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmCfmShortMaNameLengthChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmCfmShortMaNameLengthChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmCfmShortMaNameLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmCfmShortMaNameLengthCounter().msg()
	}

	if value == PatternFlowCfmCfmShortMaNameLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmCfmShortMaNameLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmCfmShortMaNameLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmCfmShortMaNameLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmCfmShortMaNameLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowCfmCfmShortMaNameLength object
func (obj *patternFlowCfmCfmShortMaNameLength) SetValue(value uint32) PatternFlowCfmCfmShortMaNameLength {
	obj.setChoice(PatternFlowCfmCfmShortMaNameLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowCfmCfmShortMaNameLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowCfmCfmShortMaNameLength object
func (obj *patternFlowCfmCfmShortMaNameLength) SetValues(value []uint32) PatternFlowCfmCfmShortMaNameLength {
	obj.setChoice(PatternFlowCfmCfmShortMaNameLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmCfmShortMaNameLengthCounter
func (obj *patternFlowCfmCfmShortMaNameLength) Increment() PatternFlowCfmCfmShortMaNameLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmCfmShortMaNameLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmCfmShortMaNameLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmCfmShortMaNameLengthCounter
func (obj *patternFlowCfmCfmShortMaNameLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmCfmShortMaNameLengthCounter value in the PatternFlowCfmCfmShortMaNameLength object
func (obj *patternFlowCfmCfmShortMaNameLength) SetIncrement(value PatternFlowCfmCfmShortMaNameLengthCounter) PatternFlowCfmCfmShortMaNameLength {
	obj.setChoice(PatternFlowCfmCfmShortMaNameLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmCfmShortMaNameLengthCounter
func (obj *patternFlowCfmCfmShortMaNameLength) Decrement() PatternFlowCfmCfmShortMaNameLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmCfmShortMaNameLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmCfmShortMaNameLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmCfmShortMaNameLengthCounter
func (obj *patternFlowCfmCfmShortMaNameLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmCfmShortMaNameLengthCounter value in the PatternFlowCfmCfmShortMaNameLength object
func (obj *patternFlowCfmCfmShortMaNameLength) SetDecrement(value PatternFlowCfmCfmShortMaNameLengthCounter) PatternFlowCfmCfmShortMaNameLength {
	obj.setChoice(PatternFlowCfmCfmShortMaNameLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmCfmShortMaNameLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmCfmShortMaNameLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowCfmCfmShortMaNameLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowCfmCfmShortMaNameLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmCfmShortMaNameLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmCfmShortMaNameLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmCfmShortMaNameLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmCfmShortMaNameLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmCfmShortMaNameLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmCfmShortMaNameLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmCfmShortMaNameLength")
			}
		} else {
			intVal := otg.PatternFlowCfmCfmShortMaNameLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmCfmShortMaNameLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
