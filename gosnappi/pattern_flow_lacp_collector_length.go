package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpCollectorLength *****
type patternFlowLacpCollectorLength struct {
	validation
	obj             *otg.PatternFlowLacpCollectorLength
	marshaller      marshalPatternFlowLacpCollectorLength
	unMarshaller    unMarshalPatternFlowLacpCollectorLength
	incrementHolder PatternFlowLacpCollectorLengthCounter
	decrementHolder PatternFlowLacpCollectorLengthCounter
}

func NewPatternFlowLacpCollectorLength() PatternFlowLacpCollectorLength {
	obj := patternFlowLacpCollectorLength{obj: &otg.PatternFlowLacpCollectorLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpCollectorLength) msg() *otg.PatternFlowLacpCollectorLength {
	return obj.obj
}

func (obj *patternFlowLacpCollectorLength) setMsg(msg *otg.PatternFlowLacpCollectorLength) PatternFlowLacpCollectorLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpCollectorLength struct {
	obj *patternFlowLacpCollectorLength
}

type marshalPatternFlowLacpCollectorLength interface {
	// ToProto marshals PatternFlowLacpCollectorLength to protobuf object *otg.PatternFlowLacpCollectorLength
	ToProto() (*otg.PatternFlowLacpCollectorLength, error)
	// ToPbText marshals PatternFlowLacpCollectorLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpCollectorLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpCollectorLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpCollectorLength struct {
	obj *patternFlowLacpCollectorLength
}

type unMarshalPatternFlowLacpCollectorLength interface {
	// FromProto unmarshals PatternFlowLacpCollectorLength from protobuf object *otg.PatternFlowLacpCollectorLength
	FromProto(msg *otg.PatternFlowLacpCollectorLength) (PatternFlowLacpCollectorLength, error)
	// FromPbText unmarshals PatternFlowLacpCollectorLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpCollectorLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpCollectorLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpCollectorLength) Marshal() marshalPatternFlowLacpCollectorLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpCollectorLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpCollectorLength) Unmarshal() unMarshalPatternFlowLacpCollectorLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpCollectorLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpCollectorLength) ToProto() (*otg.PatternFlowLacpCollectorLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpCollectorLength) FromProto(msg *otg.PatternFlowLacpCollectorLength) (PatternFlowLacpCollectorLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpCollectorLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpCollectorLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpCollectorLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorLength) FromJson(value string) error {
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

func (obj *patternFlowLacpCollectorLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpCollectorLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpCollectorLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpCollectorLength) Clone() (PatternFlowLacpCollectorLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpCollectorLength()
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

func (obj *patternFlowLacpCollectorLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpCollectorLength is the length of the Collector Information TLV payload in bytes. The value must be 16 (0x10).
type PatternFlowLacpCollectorLength interface {
	Validation
	// msg marshals PatternFlowLacpCollectorLength to protobuf object *otg.PatternFlowLacpCollectorLength
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpCollectorLength
	// setMsg unmarshals PatternFlowLacpCollectorLength from protobuf object *otg.PatternFlowLacpCollectorLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpCollectorLength) PatternFlowLacpCollectorLength
	// provides marshal interface
	Marshal() marshalPatternFlowLacpCollectorLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpCollectorLength
	// validate validates PatternFlowLacpCollectorLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpCollectorLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpCollectorLengthChoiceEnum, set in PatternFlowLacpCollectorLength
	Choice() PatternFlowLacpCollectorLengthChoiceEnum
	// setChoice assigns PatternFlowLacpCollectorLengthChoiceEnum provided by user to PatternFlowLacpCollectorLength
	setChoice(value PatternFlowLacpCollectorLengthChoiceEnum) PatternFlowLacpCollectorLength
	// HasChoice checks if Choice has been set in PatternFlowLacpCollectorLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpCollectorLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpCollectorLength
	SetValue(value uint32) PatternFlowLacpCollectorLength
	// HasValue checks if Value has been set in PatternFlowLacpCollectorLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpCollectorLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpCollectorLength
	SetValues(value []uint32) PatternFlowLacpCollectorLength
	// Increment returns PatternFlowLacpCollectorLengthCounter, set in PatternFlowLacpCollectorLength.
	// PatternFlowLacpCollectorLengthCounter is integer counter pattern
	Increment() PatternFlowLacpCollectorLengthCounter
	// SetIncrement assigns PatternFlowLacpCollectorLengthCounter provided by user to PatternFlowLacpCollectorLength.
	// PatternFlowLacpCollectorLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpCollectorLengthCounter) PatternFlowLacpCollectorLength
	// HasIncrement checks if Increment has been set in PatternFlowLacpCollectorLength
	HasIncrement() bool
	// Decrement returns PatternFlowLacpCollectorLengthCounter, set in PatternFlowLacpCollectorLength.
	// PatternFlowLacpCollectorLengthCounter is integer counter pattern
	Decrement() PatternFlowLacpCollectorLengthCounter
	// SetDecrement assigns PatternFlowLacpCollectorLengthCounter provided by user to PatternFlowLacpCollectorLength.
	// PatternFlowLacpCollectorLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpCollectorLengthCounter) PatternFlowLacpCollectorLength
	// HasDecrement checks if Decrement has been set in PatternFlowLacpCollectorLength
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpCollectorLengthChoiceEnum string

// Enum of Choice on PatternFlowLacpCollectorLength
var PatternFlowLacpCollectorLengthChoice = struct {
	VALUE     PatternFlowLacpCollectorLengthChoiceEnum
	VALUES    PatternFlowLacpCollectorLengthChoiceEnum
	INCREMENT PatternFlowLacpCollectorLengthChoiceEnum
	DECREMENT PatternFlowLacpCollectorLengthChoiceEnum
}{
	VALUE:     PatternFlowLacpCollectorLengthChoiceEnum("value"),
	VALUES:    PatternFlowLacpCollectorLengthChoiceEnum("values"),
	INCREMENT: PatternFlowLacpCollectorLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpCollectorLengthChoiceEnum("decrement"),
}

func (obj *patternFlowLacpCollectorLength) Choice() PatternFlowLacpCollectorLengthChoiceEnum {
	return PatternFlowLacpCollectorLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpCollectorLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpCollectorLength) setChoice(value PatternFlowLacpCollectorLengthChoiceEnum) PatternFlowLacpCollectorLength {
	intValue, ok := otg.PatternFlowLacpCollectorLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpCollectorLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpCollectorLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpCollectorLengthChoice.VALUE {
		defaultValue := uint32(16)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpCollectorLengthChoice.VALUES {
		defaultValue := []uint32{16}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpCollectorLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpCollectorLengthCounter().msg()
	}

	if value == PatternFlowLacpCollectorLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpCollectorLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpCollectorLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpCollectorLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpCollectorLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpCollectorLength object
func (obj *patternFlowLacpCollectorLength) SetValue(value uint32) PatternFlowLacpCollectorLength {
	obj.setChoice(PatternFlowLacpCollectorLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpCollectorLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{16})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpCollectorLength object
func (obj *patternFlowLacpCollectorLength) SetValues(value []uint32) PatternFlowLacpCollectorLength {
	obj.setChoice(PatternFlowLacpCollectorLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpCollectorLengthCounter
func (obj *patternFlowLacpCollectorLength) Increment() PatternFlowLacpCollectorLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpCollectorLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpCollectorLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpCollectorLengthCounter
func (obj *patternFlowLacpCollectorLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpCollectorLengthCounter value in the PatternFlowLacpCollectorLength object
func (obj *patternFlowLacpCollectorLength) SetIncrement(value PatternFlowLacpCollectorLengthCounter) PatternFlowLacpCollectorLength {
	obj.setChoice(PatternFlowLacpCollectorLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpCollectorLengthCounter
func (obj *patternFlowLacpCollectorLength) Decrement() PatternFlowLacpCollectorLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpCollectorLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpCollectorLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpCollectorLengthCounter
func (obj *patternFlowLacpCollectorLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpCollectorLengthCounter value in the PatternFlowLacpCollectorLength object
func (obj *patternFlowLacpCollectorLength) SetDecrement(value PatternFlowLacpCollectorLengthCounter) PatternFlowLacpCollectorLength {
	obj.setChoice(PatternFlowLacpCollectorLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpCollectorLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpCollectorLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpCollectorLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpCollectorLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpCollectorLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpCollectorLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpCollectorLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpCollectorLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpCollectorLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpCollectorLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpCollectorLength")
			}
		} else {
			intVal := otg.PatternFlowLacpCollectorLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpCollectorLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
