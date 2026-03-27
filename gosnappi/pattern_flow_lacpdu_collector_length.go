package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduCollectorLength *****
type patternFlowLacpduCollectorLength struct {
	validation
	obj             *otg.PatternFlowLacpduCollectorLength
	marshaller      marshalPatternFlowLacpduCollectorLength
	unMarshaller    unMarshalPatternFlowLacpduCollectorLength
	incrementHolder PatternFlowLacpduCollectorLengthCounter
	decrementHolder PatternFlowLacpduCollectorLengthCounter
}

func NewPatternFlowLacpduCollectorLength() PatternFlowLacpduCollectorLength {
	obj := patternFlowLacpduCollectorLength{obj: &otg.PatternFlowLacpduCollectorLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduCollectorLength) msg() *otg.PatternFlowLacpduCollectorLength {
	return obj.obj
}

func (obj *patternFlowLacpduCollectorLength) setMsg(msg *otg.PatternFlowLacpduCollectorLength) PatternFlowLacpduCollectorLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduCollectorLength struct {
	obj *patternFlowLacpduCollectorLength
}

type marshalPatternFlowLacpduCollectorLength interface {
	// ToProto marshals PatternFlowLacpduCollectorLength to protobuf object *otg.PatternFlowLacpduCollectorLength
	ToProto() (*otg.PatternFlowLacpduCollectorLength, error)
	// ToPbText marshals PatternFlowLacpduCollectorLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduCollectorLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduCollectorLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduCollectorLength struct {
	obj *patternFlowLacpduCollectorLength
}

type unMarshalPatternFlowLacpduCollectorLength interface {
	// FromProto unmarshals PatternFlowLacpduCollectorLength from protobuf object *otg.PatternFlowLacpduCollectorLength
	FromProto(msg *otg.PatternFlowLacpduCollectorLength) (PatternFlowLacpduCollectorLength, error)
	// FromPbText unmarshals PatternFlowLacpduCollectorLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduCollectorLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduCollectorLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduCollectorLength) Marshal() marshalPatternFlowLacpduCollectorLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduCollectorLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduCollectorLength) Unmarshal() unMarshalPatternFlowLacpduCollectorLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduCollectorLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduCollectorLength) ToProto() (*otg.PatternFlowLacpduCollectorLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduCollectorLength) FromProto(msg *otg.PatternFlowLacpduCollectorLength) (PatternFlowLacpduCollectorLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduCollectorLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduCollectorLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduCollectorLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorLength) FromJson(value string) error {
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

func (obj *patternFlowLacpduCollectorLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduCollectorLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduCollectorLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduCollectorLength) Clone() (PatternFlowLacpduCollectorLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduCollectorLength()
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

func (obj *patternFlowLacpduCollectorLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduCollectorLength is the length of the Collector Information TLV payload in bytes. The value must be 16 (0x10).
type PatternFlowLacpduCollectorLength interface {
	Validation
	// msg marshals PatternFlowLacpduCollectorLength to protobuf object *otg.PatternFlowLacpduCollectorLength
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduCollectorLength
	// setMsg unmarshals PatternFlowLacpduCollectorLength from protobuf object *otg.PatternFlowLacpduCollectorLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduCollectorLength) PatternFlowLacpduCollectorLength
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduCollectorLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduCollectorLength
	// validate validates PatternFlowLacpduCollectorLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduCollectorLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduCollectorLengthChoiceEnum, set in PatternFlowLacpduCollectorLength
	Choice() PatternFlowLacpduCollectorLengthChoiceEnum
	// setChoice assigns PatternFlowLacpduCollectorLengthChoiceEnum provided by user to PatternFlowLacpduCollectorLength
	setChoice(value PatternFlowLacpduCollectorLengthChoiceEnum) PatternFlowLacpduCollectorLength
	// HasChoice checks if Choice has been set in PatternFlowLacpduCollectorLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduCollectorLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduCollectorLength
	SetValue(value uint32) PatternFlowLacpduCollectorLength
	// HasValue checks if Value has been set in PatternFlowLacpduCollectorLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduCollectorLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduCollectorLength
	SetValues(value []uint32) PatternFlowLacpduCollectorLength
	// Increment returns PatternFlowLacpduCollectorLengthCounter, set in PatternFlowLacpduCollectorLength.
	// PatternFlowLacpduCollectorLengthCounter is integer counter pattern
	Increment() PatternFlowLacpduCollectorLengthCounter
	// SetIncrement assigns PatternFlowLacpduCollectorLengthCounter provided by user to PatternFlowLacpduCollectorLength.
	// PatternFlowLacpduCollectorLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduCollectorLengthCounter) PatternFlowLacpduCollectorLength
	// HasIncrement checks if Increment has been set in PatternFlowLacpduCollectorLength
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduCollectorLengthCounter, set in PatternFlowLacpduCollectorLength.
	// PatternFlowLacpduCollectorLengthCounter is integer counter pattern
	Decrement() PatternFlowLacpduCollectorLengthCounter
	// SetDecrement assigns PatternFlowLacpduCollectorLengthCounter provided by user to PatternFlowLacpduCollectorLength.
	// PatternFlowLacpduCollectorLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduCollectorLengthCounter) PatternFlowLacpduCollectorLength
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduCollectorLength
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduCollectorLengthChoiceEnum string

// Enum of Choice on PatternFlowLacpduCollectorLength
var PatternFlowLacpduCollectorLengthChoice = struct {
	VALUE     PatternFlowLacpduCollectorLengthChoiceEnum
	VALUES    PatternFlowLacpduCollectorLengthChoiceEnum
	INCREMENT PatternFlowLacpduCollectorLengthChoiceEnum
	DECREMENT PatternFlowLacpduCollectorLengthChoiceEnum
}{
	VALUE:     PatternFlowLacpduCollectorLengthChoiceEnum("value"),
	VALUES:    PatternFlowLacpduCollectorLengthChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduCollectorLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduCollectorLengthChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduCollectorLength) Choice() PatternFlowLacpduCollectorLengthChoiceEnum {
	return PatternFlowLacpduCollectorLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduCollectorLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduCollectorLength) setChoice(value PatternFlowLacpduCollectorLengthChoiceEnum) PatternFlowLacpduCollectorLength {
	intValue, ok := otg.PatternFlowLacpduCollectorLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduCollectorLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduCollectorLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduCollectorLengthChoice.VALUE {
		defaultValue := uint32(16)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduCollectorLengthChoice.VALUES {
		defaultValue := []uint32{16}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduCollectorLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduCollectorLengthCounter().msg()
	}

	if value == PatternFlowLacpduCollectorLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduCollectorLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduCollectorLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduCollectorLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduCollectorLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduCollectorLength object
func (obj *patternFlowLacpduCollectorLength) SetValue(value uint32) PatternFlowLacpduCollectorLength {
	obj.setChoice(PatternFlowLacpduCollectorLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduCollectorLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{16})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduCollectorLength object
func (obj *patternFlowLacpduCollectorLength) SetValues(value []uint32) PatternFlowLacpduCollectorLength {
	obj.setChoice(PatternFlowLacpduCollectorLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduCollectorLengthCounter
func (obj *patternFlowLacpduCollectorLength) Increment() PatternFlowLacpduCollectorLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduCollectorLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduCollectorLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduCollectorLengthCounter
func (obj *patternFlowLacpduCollectorLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduCollectorLengthCounter value in the PatternFlowLacpduCollectorLength object
func (obj *patternFlowLacpduCollectorLength) SetIncrement(value PatternFlowLacpduCollectorLengthCounter) PatternFlowLacpduCollectorLength {
	obj.setChoice(PatternFlowLacpduCollectorLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduCollectorLengthCounter
func (obj *patternFlowLacpduCollectorLength) Decrement() PatternFlowLacpduCollectorLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduCollectorLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduCollectorLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduCollectorLengthCounter
func (obj *patternFlowLacpduCollectorLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduCollectorLengthCounter value in the PatternFlowLacpduCollectorLength object
func (obj *patternFlowLacpduCollectorLength) SetDecrement(value PatternFlowLacpduCollectorLengthCounter) PatternFlowLacpduCollectorLength {
	obj.setChoice(PatternFlowLacpduCollectorLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduCollectorLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduCollectorLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduCollectorLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpduCollectorLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduCollectorLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduCollectorLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduCollectorLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduCollectorLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduCollectorLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduCollectorLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduCollectorLength")
			}
		} else {
			intVal := otg.PatternFlowLacpduCollectorLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduCollectorLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
