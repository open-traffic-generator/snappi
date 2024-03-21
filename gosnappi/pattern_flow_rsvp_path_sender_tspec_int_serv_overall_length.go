package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServOverallLength *****
type patternFlowRSVPPathSenderTspecIntServOverallLength struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTspecIntServOverallLength
	marshaller      marshalPatternFlowRSVPPathSenderTspecIntServOverallLength
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTspecIntServOverallLength
	incrementHolder PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	decrementHolder PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServOverallLength() PatternFlowRSVPPathSenderTspecIntServOverallLength {
	obj := patternFlowRSVPPathSenderTspecIntServOverallLength{obj: &otg.PatternFlowRSVPPathSenderTspecIntServOverallLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) msg() *otg.PatternFlowRSVPPathSenderTspecIntServOverallLength {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServOverallLength) PatternFlowRSVPPathSenderTspecIntServOverallLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServOverallLength struct {
	obj *patternFlowRSVPPathSenderTspecIntServOverallLength
}

type marshalPatternFlowRSVPPathSenderTspecIntServOverallLength interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServOverallLength to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServOverallLength
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServOverallLength, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServOverallLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServOverallLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServOverallLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServOverallLength struct {
	obj *patternFlowRSVPPathSenderTspecIntServOverallLength
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServOverallLength interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServOverallLength from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServOverallLength
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServOverallLength) (PatternFlowRSVPPathSenderTspecIntServOverallLength, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServOverallLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServOverallLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServOverallLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServOverallLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServOverallLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServOverallLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServOverallLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServOverallLength) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServOverallLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServOverallLength) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServOverallLength) (PatternFlowRSVPPathSenderTspecIntServOverallLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServOverallLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServOverallLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServOverallLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServOverallLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServOverallLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServOverallLength) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) Clone() (PatternFlowRSVPPathSenderTspecIntServOverallLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServOverallLength()
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

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTspecIntServOverallLength is overall length (7 words not including header).
type PatternFlowRSVPPathSenderTspecIntServOverallLength interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServOverallLength to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServOverallLength
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServOverallLength
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServOverallLength from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServOverallLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServOverallLength) PatternFlowRSVPPathSenderTspecIntServOverallLength
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServOverallLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServOverallLength
	// validate validates PatternFlowRSVPPathSenderTspecIntServOverallLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServOverallLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum, set in PatternFlowRSVPPathSenderTspecIntServOverallLength
	Choice() PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum provided by user to PatternFlowRSVPPathSenderTspecIntServOverallLength
	setChoice(value PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum) PatternFlowRSVPPathSenderTspecIntServOverallLength
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTspecIntServOverallLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTspecIntServOverallLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServOverallLength
	SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServOverallLength
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTspecIntServOverallLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTspecIntServOverallLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServOverallLength
	SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServOverallLength
	// Increment returns PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter, set in PatternFlowRSVPPathSenderTspecIntServOverallLength.
	// PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	// SetIncrement assigns PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter provided by user to PatternFlowRSVPPathSenderTspecIntServOverallLength.
	// PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) PatternFlowRSVPPathSenderTspecIntServOverallLength
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTspecIntServOverallLength
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter, set in PatternFlowRSVPPathSenderTspecIntServOverallLength.
	// PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	// SetDecrement assigns PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter provided by user to PatternFlowRSVPPathSenderTspecIntServOverallLength.
	// PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) PatternFlowRSVPPathSenderTspecIntServOverallLength
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTspecIntServOverallLength
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTspecIntServOverallLength
var PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice = struct {
	VALUE     PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) Choice() PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum {
	return PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) setChoice(value PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum) PatternFlowRSVPPathSenderTspecIntServOverallLength {
	intValue, ok := otg.PatternFlowRSVPPathSenderTspecIntServOverallLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTspecIntServOverallLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.VALUE {
		defaultValue := uint32(7)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.VALUES {
		defaultValue := []uint32{7}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTspecIntServOverallLengthCounter().msg()
	}

	if value == PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTspecIntServOverallLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServOverallLength object
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServOverallLength {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{7})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTspecIntServOverallLength object
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServOverallLength {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) Increment() PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTspecIntServOverallLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter value in the PatternFlowRSVPPathSenderTspecIntServOverallLength object
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) SetIncrement(value PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) PatternFlowRSVPPathSenderTspecIntServOverallLength {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) Decrement() PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTspecIntServOverallLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter value in the PatternFlowRSVPPathSenderTspecIntServOverallLength object
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) SetDecrement(value PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) PatternFlowRSVPPathSenderTspecIntServOverallLength {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServOverallLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSenderTspecIntServOverallLength.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSenderTspecIntServOverallLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSenderTspecIntServOverallLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSenderTspecIntServOverallLength")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSenderTspecIntServOverallLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSenderTspecIntServOverallLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
