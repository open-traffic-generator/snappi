package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6UsidDstLocatorLength *****
type patternFlowIpv6UsidDstLocatorLength struct {
	validation
	obj             *otg.PatternFlowIpv6UsidDstLocatorLength
	marshaller      marshalPatternFlowIpv6UsidDstLocatorLength
	unMarshaller    unMarshalPatternFlowIpv6UsidDstLocatorLength
	incrementHolder PatternFlowIpv6UsidDstLocatorLengthCounter
	decrementHolder PatternFlowIpv6UsidDstLocatorLengthCounter
}

func NewPatternFlowIpv6UsidDstLocatorLength() PatternFlowIpv6UsidDstLocatorLength {
	obj := patternFlowIpv6UsidDstLocatorLength{obj: &otg.PatternFlowIpv6UsidDstLocatorLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6UsidDstLocatorLength) msg() *otg.PatternFlowIpv6UsidDstLocatorLength {
	return obj.obj
}

func (obj *patternFlowIpv6UsidDstLocatorLength) setMsg(msg *otg.PatternFlowIpv6UsidDstLocatorLength) PatternFlowIpv6UsidDstLocatorLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6UsidDstLocatorLength struct {
	obj *patternFlowIpv6UsidDstLocatorLength
}

type marshalPatternFlowIpv6UsidDstLocatorLength interface {
	// ToProto marshals PatternFlowIpv6UsidDstLocatorLength to protobuf object *otg.PatternFlowIpv6UsidDstLocatorLength
	ToProto() (*otg.PatternFlowIpv6UsidDstLocatorLength, error)
	// ToPbText marshals PatternFlowIpv6UsidDstLocatorLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6UsidDstLocatorLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6UsidDstLocatorLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6UsidDstLocatorLength struct {
	obj *patternFlowIpv6UsidDstLocatorLength
}

type unMarshalPatternFlowIpv6UsidDstLocatorLength interface {
	// FromProto unmarshals PatternFlowIpv6UsidDstLocatorLength from protobuf object *otg.PatternFlowIpv6UsidDstLocatorLength
	FromProto(msg *otg.PatternFlowIpv6UsidDstLocatorLength) (PatternFlowIpv6UsidDstLocatorLength, error)
	// FromPbText unmarshals PatternFlowIpv6UsidDstLocatorLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6UsidDstLocatorLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6UsidDstLocatorLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6UsidDstLocatorLength) Marshal() marshalPatternFlowIpv6UsidDstLocatorLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6UsidDstLocatorLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6UsidDstLocatorLength) Unmarshal() unMarshalPatternFlowIpv6UsidDstLocatorLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6UsidDstLocatorLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6UsidDstLocatorLength) ToProto() (*otg.PatternFlowIpv6UsidDstLocatorLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6UsidDstLocatorLength) FromProto(msg *otg.PatternFlowIpv6UsidDstLocatorLength) (PatternFlowIpv6UsidDstLocatorLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6UsidDstLocatorLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6UsidDstLocatorLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6UsidDstLocatorLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6UsidDstLocatorLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6UsidDstLocatorLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6UsidDstLocatorLength) FromJson(value string) error {
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

func (obj *patternFlowIpv6UsidDstLocatorLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6UsidDstLocatorLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6UsidDstLocatorLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6UsidDstLocatorLength) Clone() (PatternFlowIpv6UsidDstLocatorLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6UsidDstLocatorLength()
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

func (obj *patternFlowIpv6UsidDstLocatorLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6UsidDstLocatorLength is length of the Locator Block in bits (RFC 9800 Section 3). Determines how many high-order bits of locator are used as the LB and how many bits remain for uSID packing. For F3216: 32. For F3208: 32. Valid range: 1-112.
type PatternFlowIpv6UsidDstLocatorLength interface {
	Validation
	// msg marshals PatternFlowIpv6UsidDstLocatorLength to protobuf object *otg.PatternFlowIpv6UsidDstLocatorLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6UsidDstLocatorLength
	// setMsg unmarshals PatternFlowIpv6UsidDstLocatorLength from protobuf object *otg.PatternFlowIpv6UsidDstLocatorLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6UsidDstLocatorLength) PatternFlowIpv6UsidDstLocatorLength
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6UsidDstLocatorLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6UsidDstLocatorLength
	// validate validates PatternFlowIpv6UsidDstLocatorLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6UsidDstLocatorLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6UsidDstLocatorLengthChoiceEnum, set in PatternFlowIpv6UsidDstLocatorLength
	Choice() PatternFlowIpv6UsidDstLocatorLengthChoiceEnum
	// setChoice assigns PatternFlowIpv6UsidDstLocatorLengthChoiceEnum provided by user to PatternFlowIpv6UsidDstLocatorLength
	setChoice(value PatternFlowIpv6UsidDstLocatorLengthChoiceEnum) PatternFlowIpv6UsidDstLocatorLength
	// HasChoice checks if Choice has been set in PatternFlowIpv6UsidDstLocatorLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6UsidDstLocatorLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6UsidDstLocatorLength
	SetValue(value uint32) PatternFlowIpv6UsidDstLocatorLength
	// HasValue checks if Value has been set in PatternFlowIpv6UsidDstLocatorLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6UsidDstLocatorLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6UsidDstLocatorLength
	SetValues(value []uint32) PatternFlowIpv6UsidDstLocatorLength
	// Increment returns PatternFlowIpv6UsidDstLocatorLengthCounter, set in PatternFlowIpv6UsidDstLocatorLength.
	// PatternFlowIpv6UsidDstLocatorLengthCounter is integer counter pattern
	Increment() PatternFlowIpv6UsidDstLocatorLengthCounter
	// SetIncrement assigns PatternFlowIpv6UsidDstLocatorLengthCounter provided by user to PatternFlowIpv6UsidDstLocatorLength.
	// PatternFlowIpv6UsidDstLocatorLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6UsidDstLocatorLengthCounter) PatternFlowIpv6UsidDstLocatorLength
	// HasIncrement checks if Increment has been set in PatternFlowIpv6UsidDstLocatorLength
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6UsidDstLocatorLengthCounter, set in PatternFlowIpv6UsidDstLocatorLength.
	// PatternFlowIpv6UsidDstLocatorLengthCounter is integer counter pattern
	Decrement() PatternFlowIpv6UsidDstLocatorLengthCounter
	// SetDecrement assigns PatternFlowIpv6UsidDstLocatorLengthCounter provided by user to PatternFlowIpv6UsidDstLocatorLength.
	// PatternFlowIpv6UsidDstLocatorLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6UsidDstLocatorLengthCounter) PatternFlowIpv6UsidDstLocatorLength
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6UsidDstLocatorLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6UsidDstLocatorLengthChoiceEnum string

// Enum of Choice on PatternFlowIpv6UsidDstLocatorLength
var PatternFlowIpv6UsidDstLocatorLengthChoice = struct {
	VALUE     PatternFlowIpv6UsidDstLocatorLengthChoiceEnum
	VALUES    PatternFlowIpv6UsidDstLocatorLengthChoiceEnum
	INCREMENT PatternFlowIpv6UsidDstLocatorLengthChoiceEnum
	DECREMENT PatternFlowIpv6UsidDstLocatorLengthChoiceEnum
}{
	VALUE:     PatternFlowIpv6UsidDstLocatorLengthChoiceEnum("value"),
	VALUES:    PatternFlowIpv6UsidDstLocatorLengthChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6UsidDstLocatorLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6UsidDstLocatorLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6UsidDstLocatorLength) Choice() PatternFlowIpv6UsidDstLocatorLengthChoiceEnum {
	return PatternFlowIpv6UsidDstLocatorLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6UsidDstLocatorLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6UsidDstLocatorLength) setChoice(value PatternFlowIpv6UsidDstLocatorLengthChoiceEnum) PatternFlowIpv6UsidDstLocatorLength {
	intValue, ok := otg.PatternFlowIpv6UsidDstLocatorLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6UsidDstLocatorLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6UsidDstLocatorLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6UsidDstLocatorLengthChoice.VALUE {
		defaultValue := uint32(32)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6UsidDstLocatorLengthChoice.VALUES {
		defaultValue := []uint32{32}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6UsidDstLocatorLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6UsidDstLocatorLengthCounter().msg()
	}

	if value == PatternFlowIpv6UsidDstLocatorLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6UsidDstLocatorLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6UsidDstLocatorLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6UsidDstLocatorLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6UsidDstLocatorLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6UsidDstLocatorLength object
func (obj *patternFlowIpv6UsidDstLocatorLength) SetValue(value uint32) PatternFlowIpv6UsidDstLocatorLength {
	obj.setChoice(PatternFlowIpv6UsidDstLocatorLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6UsidDstLocatorLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{32})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6UsidDstLocatorLength object
func (obj *patternFlowIpv6UsidDstLocatorLength) SetValues(value []uint32) PatternFlowIpv6UsidDstLocatorLength {
	obj.setChoice(PatternFlowIpv6UsidDstLocatorLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6UsidDstLocatorLengthCounter
func (obj *patternFlowIpv6UsidDstLocatorLength) Increment() PatternFlowIpv6UsidDstLocatorLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6UsidDstLocatorLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6UsidDstLocatorLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6UsidDstLocatorLengthCounter
func (obj *patternFlowIpv6UsidDstLocatorLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6UsidDstLocatorLengthCounter value in the PatternFlowIpv6UsidDstLocatorLength object
func (obj *patternFlowIpv6UsidDstLocatorLength) SetIncrement(value PatternFlowIpv6UsidDstLocatorLengthCounter) PatternFlowIpv6UsidDstLocatorLength {
	obj.setChoice(PatternFlowIpv6UsidDstLocatorLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6UsidDstLocatorLengthCounter
func (obj *patternFlowIpv6UsidDstLocatorLength) Decrement() PatternFlowIpv6UsidDstLocatorLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6UsidDstLocatorLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6UsidDstLocatorLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6UsidDstLocatorLengthCounter
func (obj *patternFlowIpv6UsidDstLocatorLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6UsidDstLocatorLengthCounter value in the PatternFlowIpv6UsidDstLocatorLength object
func (obj *patternFlowIpv6UsidDstLocatorLength) SetDecrement(value PatternFlowIpv6UsidDstLocatorLengthCounter) PatternFlowIpv6UsidDstLocatorLength {
	obj.setChoice(PatternFlowIpv6UsidDstLocatorLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6UsidDstLocatorLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6UsidDstLocatorLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6UsidDstLocatorLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIpv6UsidDstLocatorLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6UsidDstLocatorLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6UsidDstLocatorLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6UsidDstLocatorLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6UsidDstLocatorLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6UsidDstLocatorLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6UsidDstLocatorLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6UsidDstLocatorLength")
			}
		} else {
			intVal := otg.PatternFlowIpv6UsidDstLocatorLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6UsidDstLocatorLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
