package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmSenderIdTlvChassisIdLength *****
type patternFlowCfmSenderIdTlvChassisIdLength struct {
	validation
	obj             *otg.PatternFlowCfmSenderIdTlvChassisIdLength
	marshaller      marshalPatternFlowCfmSenderIdTlvChassisIdLength
	unMarshaller    unMarshalPatternFlowCfmSenderIdTlvChassisIdLength
	incrementHolder PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	decrementHolder PatternFlowCfmSenderIdTlvChassisIdLengthCounter
}

func NewPatternFlowCfmSenderIdTlvChassisIdLength() PatternFlowCfmSenderIdTlvChassisIdLength {
	obj := patternFlowCfmSenderIdTlvChassisIdLength{obj: &otg.PatternFlowCfmSenderIdTlvChassisIdLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLength) msg() *otg.PatternFlowCfmSenderIdTlvChassisIdLength {
	return obj.obj
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLength) setMsg(msg *otg.PatternFlowCfmSenderIdTlvChassisIdLength) PatternFlowCfmSenderIdTlvChassisIdLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmSenderIdTlvChassisIdLength struct {
	obj *patternFlowCfmSenderIdTlvChassisIdLength
}

type marshalPatternFlowCfmSenderIdTlvChassisIdLength interface {
	// ToProto marshals PatternFlowCfmSenderIdTlvChassisIdLength to protobuf object *otg.PatternFlowCfmSenderIdTlvChassisIdLength
	ToProto() (*otg.PatternFlowCfmSenderIdTlvChassisIdLength, error)
	// ToPbText marshals PatternFlowCfmSenderIdTlvChassisIdLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmSenderIdTlvChassisIdLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmSenderIdTlvChassisIdLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmSenderIdTlvChassisIdLength struct {
	obj *patternFlowCfmSenderIdTlvChassisIdLength
}

type unMarshalPatternFlowCfmSenderIdTlvChassisIdLength interface {
	// FromProto unmarshals PatternFlowCfmSenderIdTlvChassisIdLength from protobuf object *otg.PatternFlowCfmSenderIdTlvChassisIdLength
	FromProto(msg *otg.PatternFlowCfmSenderIdTlvChassisIdLength) (PatternFlowCfmSenderIdTlvChassisIdLength, error)
	// FromPbText unmarshals PatternFlowCfmSenderIdTlvChassisIdLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmSenderIdTlvChassisIdLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmSenderIdTlvChassisIdLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLength) Marshal() marshalPatternFlowCfmSenderIdTlvChassisIdLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmSenderIdTlvChassisIdLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLength) Unmarshal() unMarshalPatternFlowCfmSenderIdTlvChassisIdLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmSenderIdTlvChassisIdLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmSenderIdTlvChassisIdLength) ToProto() (*otg.PatternFlowCfmSenderIdTlvChassisIdLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmSenderIdTlvChassisIdLength) FromProto(msg *otg.PatternFlowCfmSenderIdTlvChassisIdLength) (PatternFlowCfmSenderIdTlvChassisIdLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmSenderIdTlvChassisIdLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmSenderIdTlvChassisIdLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmSenderIdTlvChassisIdLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmSenderIdTlvChassisIdLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmSenderIdTlvChassisIdLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmSenderIdTlvChassisIdLength) FromJson(value string) error {
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

func (obj *patternFlowCfmSenderIdTlvChassisIdLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLength) Clone() (PatternFlowCfmSenderIdTlvChassisIdLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmSenderIdTlvChassisIdLength()
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

func (obj *patternFlowCfmSenderIdTlvChassisIdLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmSenderIdTlvChassisIdLength is specifies the length of the chassis ID field.
type PatternFlowCfmSenderIdTlvChassisIdLength interface {
	Validation
	// msg marshals PatternFlowCfmSenderIdTlvChassisIdLength to protobuf object *otg.PatternFlowCfmSenderIdTlvChassisIdLength
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmSenderIdTlvChassisIdLength
	// setMsg unmarshals PatternFlowCfmSenderIdTlvChassisIdLength from protobuf object *otg.PatternFlowCfmSenderIdTlvChassisIdLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmSenderIdTlvChassisIdLength) PatternFlowCfmSenderIdTlvChassisIdLength
	// provides marshal interface
	Marshal() marshalPatternFlowCfmSenderIdTlvChassisIdLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmSenderIdTlvChassisIdLength
	// validate validates PatternFlowCfmSenderIdTlvChassisIdLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmSenderIdTlvChassisIdLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum, set in PatternFlowCfmSenderIdTlvChassisIdLength
	Choice() PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum
	// setChoice assigns PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum provided by user to PatternFlowCfmSenderIdTlvChassisIdLength
	setChoice(value PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum) PatternFlowCfmSenderIdTlvChassisIdLength
	// HasChoice checks if Choice has been set in PatternFlowCfmSenderIdTlvChassisIdLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowCfmSenderIdTlvChassisIdLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowCfmSenderIdTlvChassisIdLength
	SetValue(value uint32) PatternFlowCfmSenderIdTlvChassisIdLength
	// HasValue checks if Value has been set in PatternFlowCfmSenderIdTlvChassisIdLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowCfmSenderIdTlvChassisIdLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowCfmSenderIdTlvChassisIdLength
	SetValues(value []uint32) PatternFlowCfmSenderIdTlvChassisIdLength
	// Increment returns PatternFlowCfmSenderIdTlvChassisIdLengthCounter, set in PatternFlowCfmSenderIdTlvChassisIdLength.
	// PatternFlowCfmSenderIdTlvChassisIdLengthCounter is integer counter pattern
	Increment() PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	// SetIncrement assigns PatternFlowCfmSenderIdTlvChassisIdLengthCounter provided by user to PatternFlowCfmSenderIdTlvChassisIdLength.
	// PatternFlowCfmSenderIdTlvChassisIdLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmSenderIdTlvChassisIdLengthCounter) PatternFlowCfmSenderIdTlvChassisIdLength
	// HasIncrement checks if Increment has been set in PatternFlowCfmSenderIdTlvChassisIdLength
	HasIncrement() bool
	// Decrement returns PatternFlowCfmSenderIdTlvChassisIdLengthCounter, set in PatternFlowCfmSenderIdTlvChassisIdLength.
	// PatternFlowCfmSenderIdTlvChassisIdLengthCounter is integer counter pattern
	Decrement() PatternFlowCfmSenderIdTlvChassisIdLengthCounter
	// SetDecrement assigns PatternFlowCfmSenderIdTlvChassisIdLengthCounter provided by user to PatternFlowCfmSenderIdTlvChassisIdLength.
	// PatternFlowCfmSenderIdTlvChassisIdLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmSenderIdTlvChassisIdLengthCounter) PatternFlowCfmSenderIdTlvChassisIdLength
	// HasDecrement checks if Decrement has been set in PatternFlowCfmSenderIdTlvChassisIdLength
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum string

// Enum of Choice on PatternFlowCfmSenderIdTlvChassisIdLength
var PatternFlowCfmSenderIdTlvChassisIdLengthChoice = struct {
	VALUE     PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum
	VALUES    PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum
	INCREMENT PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum
	DECREMENT PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum
}{
	VALUE:     PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum("value"),
	VALUES:    PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum("values"),
	INCREMENT: PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum("decrement"),
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLength) Choice() PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum {
	return PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmSenderIdTlvChassisIdLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLength) setChoice(value PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum) PatternFlowCfmSenderIdTlvChassisIdLength {
	intValue, ok := otg.PatternFlowCfmSenderIdTlvChassisIdLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmSenderIdTlvChassisIdLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmSenderIdTlvChassisIdLengthChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmSenderIdTlvChassisIdLengthChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmSenderIdTlvChassisIdLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmSenderIdTlvChassisIdLengthCounter().msg()
	}

	if value == PatternFlowCfmSenderIdTlvChassisIdLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmSenderIdTlvChassisIdLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmSenderIdTlvChassisIdLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmSenderIdTlvChassisIdLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmSenderIdTlvChassisIdLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowCfmSenderIdTlvChassisIdLength object
func (obj *patternFlowCfmSenderIdTlvChassisIdLength) SetValue(value uint32) PatternFlowCfmSenderIdTlvChassisIdLength {
	obj.setChoice(PatternFlowCfmSenderIdTlvChassisIdLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowCfmSenderIdTlvChassisIdLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowCfmSenderIdTlvChassisIdLength object
func (obj *patternFlowCfmSenderIdTlvChassisIdLength) SetValues(value []uint32) PatternFlowCfmSenderIdTlvChassisIdLength {
	obj.setChoice(PatternFlowCfmSenderIdTlvChassisIdLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmSenderIdTlvChassisIdLengthCounter
func (obj *patternFlowCfmSenderIdTlvChassisIdLength) Increment() PatternFlowCfmSenderIdTlvChassisIdLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmSenderIdTlvChassisIdLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmSenderIdTlvChassisIdLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmSenderIdTlvChassisIdLengthCounter
func (obj *patternFlowCfmSenderIdTlvChassisIdLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmSenderIdTlvChassisIdLengthCounter value in the PatternFlowCfmSenderIdTlvChassisIdLength object
func (obj *patternFlowCfmSenderIdTlvChassisIdLength) SetIncrement(value PatternFlowCfmSenderIdTlvChassisIdLengthCounter) PatternFlowCfmSenderIdTlvChassisIdLength {
	obj.setChoice(PatternFlowCfmSenderIdTlvChassisIdLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmSenderIdTlvChassisIdLengthCounter
func (obj *patternFlowCfmSenderIdTlvChassisIdLength) Decrement() PatternFlowCfmSenderIdTlvChassisIdLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmSenderIdTlvChassisIdLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmSenderIdTlvChassisIdLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmSenderIdTlvChassisIdLengthCounter
func (obj *patternFlowCfmSenderIdTlvChassisIdLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmSenderIdTlvChassisIdLengthCounter value in the PatternFlowCfmSenderIdTlvChassisIdLength object
func (obj *patternFlowCfmSenderIdTlvChassisIdLength) SetDecrement(value PatternFlowCfmSenderIdTlvChassisIdLengthCounter) PatternFlowCfmSenderIdTlvChassisIdLength {
	obj.setChoice(PatternFlowCfmSenderIdTlvChassisIdLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmSenderIdTlvChassisIdLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmSenderIdTlvChassisIdLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowCfmSenderIdTlvChassisIdLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowCfmSenderIdTlvChassisIdLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmSenderIdTlvChassisIdLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmSenderIdTlvChassisIdLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmSenderIdTlvChassisIdLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmSenderIdTlvChassisIdLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmSenderIdTlvChassisIdLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmSenderIdTlvChassisIdLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmSenderIdTlvChassisIdLength")
			}
		} else {
			intVal := otg.PatternFlowCfmSenderIdTlvChassisIdLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmSenderIdTlvChassisIdLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
