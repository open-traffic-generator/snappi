package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmMDNameSpecifiedMdNameLength *****
type patternFlowCfmMDNameSpecifiedMdNameLength struct {
	validation
	obj             *otg.PatternFlowCfmMDNameSpecifiedMdNameLength
	marshaller      marshalPatternFlowCfmMDNameSpecifiedMdNameLength
	unMarshaller    unMarshalPatternFlowCfmMDNameSpecifiedMdNameLength
	incrementHolder PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	decrementHolder PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
}

func NewPatternFlowCfmMDNameSpecifiedMdNameLength() PatternFlowCfmMDNameSpecifiedMdNameLength {
	obj := patternFlowCfmMDNameSpecifiedMdNameLength{obj: &otg.PatternFlowCfmMDNameSpecifiedMdNameLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) msg() *otg.PatternFlowCfmMDNameSpecifiedMdNameLength {
	return obj.obj
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) setMsg(msg *otg.PatternFlowCfmMDNameSpecifiedMdNameLength) PatternFlowCfmMDNameSpecifiedMdNameLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmMDNameSpecifiedMdNameLength struct {
	obj *patternFlowCfmMDNameSpecifiedMdNameLength
}

type marshalPatternFlowCfmMDNameSpecifiedMdNameLength interface {
	// ToProto marshals PatternFlowCfmMDNameSpecifiedMdNameLength to protobuf object *otg.PatternFlowCfmMDNameSpecifiedMdNameLength
	ToProto() (*otg.PatternFlowCfmMDNameSpecifiedMdNameLength, error)
	// ToPbText marshals PatternFlowCfmMDNameSpecifiedMdNameLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmMDNameSpecifiedMdNameLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmMDNameSpecifiedMdNameLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmMDNameSpecifiedMdNameLength struct {
	obj *patternFlowCfmMDNameSpecifiedMdNameLength
}

type unMarshalPatternFlowCfmMDNameSpecifiedMdNameLength interface {
	// FromProto unmarshals PatternFlowCfmMDNameSpecifiedMdNameLength from protobuf object *otg.PatternFlowCfmMDNameSpecifiedMdNameLength
	FromProto(msg *otg.PatternFlowCfmMDNameSpecifiedMdNameLength) (PatternFlowCfmMDNameSpecifiedMdNameLength, error)
	// FromPbText unmarshals PatternFlowCfmMDNameSpecifiedMdNameLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmMDNameSpecifiedMdNameLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmMDNameSpecifiedMdNameLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) Marshal() marshalPatternFlowCfmMDNameSpecifiedMdNameLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmMDNameSpecifiedMdNameLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) Unmarshal() unMarshalPatternFlowCfmMDNameSpecifiedMdNameLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmMDNameSpecifiedMdNameLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmMDNameSpecifiedMdNameLength) ToProto() (*otg.PatternFlowCfmMDNameSpecifiedMdNameLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmMDNameSpecifiedMdNameLength) FromProto(msg *otg.PatternFlowCfmMDNameSpecifiedMdNameLength) (PatternFlowCfmMDNameSpecifiedMdNameLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmMDNameSpecifiedMdNameLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmMDNameSpecifiedMdNameLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmMDNameSpecifiedMdNameLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmMDNameSpecifiedMdNameLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmMDNameSpecifiedMdNameLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmMDNameSpecifiedMdNameLength) FromJson(value string) error {
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

func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) Clone() (PatternFlowCfmMDNameSpecifiedMdNameLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmMDNameSpecifiedMdNameLength()
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

func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmMDNameSpecifiedMdNameLength is maintenance domain name length
type PatternFlowCfmMDNameSpecifiedMdNameLength interface {
	Validation
	// msg marshals PatternFlowCfmMDNameSpecifiedMdNameLength to protobuf object *otg.PatternFlowCfmMDNameSpecifiedMdNameLength
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmMDNameSpecifiedMdNameLength
	// setMsg unmarshals PatternFlowCfmMDNameSpecifiedMdNameLength from protobuf object *otg.PatternFlowCfmMDNameSpecifiedMdNameLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmMDNameSpecifiedMdNameLength) PatternFlowCfmMDNameSpecifiedMdNameLength
	// provides marshal interface
	Marshal() marshalPatternFlowCfmMDNameSpecifiedMdNameLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmMDNameSpecifiedMdNameLength
	// validate validates PatternFlowCfmMDNameSpecifiedMdNameLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmMDNameSpecifiedMdNameLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum, set in PatternFlowCfmMDNameSpecifiedMdNameLength
	Choice() PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum
	// setChoice assigns PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum provided by user to PatternFlowCfmMDNameSpecifiedMdNameLength
	setChoice(value PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum) PatternFlowCfmMDNameSpecifiedMdNameLength
	// HasChoice checks if Choice has been set in PatternFlowCfmMDNameSpecifiedMdNameLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowCfmMDNameSpecifiedMdNameLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowCfmMDNameSpecifiedMdNameLength
	SetValue(value uint32) PatternFlowCfmMDNameSpecifiedMdNameLength
	// HasValue checks if Value has been set in PatternFlowCfmMDNameSpecifiedMdNameLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowCfmMDNameSpecifiedMdNameLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowCfmMDNameSpecifiedMdNameLength
	SetValues(value []uint32) PatternFlowCfmMDNameSpecifiedMdNameLength
	// Increment returns PatternFlowCfmMDNameSpecifiedMdNameLengthCounter, set in PatternFlowCfmMDNameSpecifiedMdNameLength.
	// PatternFlowCfmMDNameSpecifiedMdNameLengthCounter is integer counter pattern
	Increment() PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	// SetIncrement assigns PatternFlowCfmMDNameSpecifiedMdNameLengthCounter provided by user to PatternFlowCfmMDNameSpecifiedMdNameLength.
	// PatternFlowCfmMDNameSpecifiedMdNameLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmMDNameSpecifiedMdNameLengthCounter) PatternFlowCfmMDNameSpecifiedMdNameLength
	// HasIncrement checks if Increment has been set in PatternFlowCfmMDNameSpecifiedMdNameLength
	HasIncrement() bool
	// Decrement returns PatternFlowCfmMDNameSpecifiedMdNameLengthCounter, set in PatternFlowCfmMDNameSpecifiedMdNameLength.
	// PatternFlowCfmMDNameSpecifiedMdNameLengthCounter is integer counter pattern
	Decrement() PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
	// SetDecrement assigns PatternFlowCfmMDNameSpecifiedMdNameLengthCounter provided by user to PatternFlowCfmMDNameSpecifiedMdNameLength.
	// PatternFlowCfmMDNameSpecifiedMdNameLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmMDNameSpecifiedMdNameLengthCounter) PatternFlowCfmMDNameSpecifiedMdNameLength
	// HasDecrement checks if Decrement has been set in PatternFlowCfmMDNameSpecifiedMdNameLength
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum string

// Enum of Choice on PatternFlowCfmMDNameSpecifiedMdNameLength
var PatternFlowCfmMDNameSpecifiedMdNameLengthChoice = struct {
	VALUE     PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum
	VALUES    PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum
	INCREMENT PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum
	DECREMENT PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum
}{
	VALUE:     PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum("value"),
	VALUES:    PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum("values"),
	INCREMENT: PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum("decrement"),
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) Choice() PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum {
	return PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) setChoice(value PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum) PatternFlowCfmMDNameSpecifiedMdNameLength {
	intValue, ok := otg.PatternFlowCfmMDNameSpecifiedMdNameLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmMDNameSpecifiedMdNameLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmMDNameSpecifiedMdNameLengthCounter().msg()
	}

	if value == PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmMDNameSpecifiedMdNameLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowCfmMDNameSpecifiedMdNameLength object
func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) SetValue(value uint32) PatternFlowCfmMDNameSpecifiedMdNameLength {
	obj.setChoice(PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowCfmMDNameSpecifiedMdNameLength object
func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) SetValues(value []uint32) PatternFlowCfmMDNameSpecifiedMdNameLength {
	obj.setChoice(PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) Increment() PatternFlowCfmMDNameSpecifiedMdNameLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmMDNameSpecifiedMdNameLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmMDNameSpecifiedMdNameLengthCounter value in the PatternFlowCfmMDNameSpecifiedMdNameLength object
func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) SetIncrement(value PatternFlowCfmMDNameSpecifiedMdNameLengthCounter) PatternFlowCfmMDNameSpecifiedMdNameLength {
	obj.setChoice(PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) Decrement() PatternFlowCfmMDNameSpecifiedMdNameLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmMDNameSpecifiedMdNameLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmMDNameSpecifiedMdNameLengthCounter
func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmMDNameSpecifiedMdNameLengthCounter value in the PatternFlowCfmMDNameSpecifiedMdNameLength object
func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) SetDecrement(value PatternFlowCfmMDNameSpecifiedMdNameLengthCounter) PatternFlowCfmMDNameSpecifiedMdNameLength {
	obj.setChoice(PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmMDNameSpecifiedMdNameLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowCfmMDNameSpecifiedMdNameLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowCfmMDNameSpecifiedMdNameLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmMDNameSpecifiedMdNameLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmMDNameSpecifiedMdNameLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmMDNameSpecifiedMdNameLength")
			}
		} else {
			intVal := otg.PatternFlowCfmMDNameSpecifiedMdNameLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmMDNameSpecifiedMdNameLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
