package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPadTlvLength *****
type patternFlowIpv6SRHPadTlvLength struct {
	validation
	obj             *otg.PatternFlowIpv6SRHPadTlvLength
	marshaller      marshalPatternFlowIpv6SRHPadTlvLength
	unMarshaller    unMarshalPatternFlowIpv6SRHPadTlvLength
	incrementHolder PatternFlowIpv6SRHPadTlvLengthCounter
	decrementHolder PatternFlowIpv6SRHPadTlvLengthCounter
}

func NewPatternFlowIpv6SRHPadTlvLength() PatternFlowIpv6SRHPadTlvLength {
	obj := patternFlowIpv6SRHPadTlvLength{obj: &otg.PatternFlowIpv6SRHPadTlvLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPadTlvLength) msg() *otg.PatternFlowIpv6SRHPadTlvLength {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPadTlvLength) setMsg(msg *otg.PatternFlowIpv6SRHPadTlvLength) PatternFlowIpv6SRHPadTlvLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPadTlvLength struct {
	obj *patternFlowIpv6SRHPadTlvLength
}

type marshalPatternFlowIpv6SRHPadTlvLength interface {
	// ToProto marshals PatternFlowIpv6SRHPadTlvLength to protobuf object *otg.PatternFlowIpv6SRHPadTlvLength
	ToProto() (*otg.PatternFlowIpv6SRHPadTlvLength, error)
	// ToPbText marshals PatternFlowIpv6SRHPadTlvLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPadTlvLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPadTlvLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPadTlvLength struct {
	obj *patternFlowIpv6SRHPadTlvLength
}

type unMarshalPatternFlowIpv6SRHPadTlvLength interface {
	// FromProto unmarshals PatternFlowIpv6SRHPadTlvLength from protobuf object *otg.PatternFlowIpv6SRHPadTlvLength
	FromProto(msg *otg.PatternFlowIpv6SRHPadTlvLength) (PatternFlowIpv6SRHPadTlvLength, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPadTlvLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPadTlvLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPadTlvLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPadTlvLength) Marshal() marshalPatternFlowIpv6SRHPadTlvLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPadTlvLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPadTlvLength) Unmarshal() unMarshalPatternFlowIpv6SRHPadTlvLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPadTlvLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPadTlvLength) ToProto() (*otg.PatternFlowIpv6SRHPadTlvLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPadTlvLength) FromProto(msg *otg.PatternFlowIpv6SRHPadTlvLength) (PatternFlowIpv6SRHPadTlvLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPadTlvLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPadTlvLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPadTlvLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPadTlvLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPadTlvLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPadTlvLength) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPadTlvLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPadTlvLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPadTlvLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPadTlvLength) Clone() (PatternFlowIpv6SRHPadTlvLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPadTlvLength()
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

func (obj *patternFlowIpv6SRHPadTlvLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHPadTlvLength is number of padding bytes following the Type and Length fields (RFC 8754 Section 2.1). When auto is assigned the implementation computes the value needed to align the TLV section to an 8-byte boundary.
type PatternFlowIpv6SRHPadTlvLength interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPadTlvLength to protobuf object *otg.PatternFlowIpv6SRHPadTlvLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPadTlvLength
	// setMsg unmarshals PatternFlowIpv6SRHPadTlvLength from protobuf object *otg.PatternFlowIpv6SRHPadTlvLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPadTlvLength) PatternFlowIpv6SRHPadTlvLength
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPadTlvLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPadTlvLength
	// validate validates PatternFlowIpv6SRHPadTlvLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPadTlvLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHPadTlvLengthChoiceEnum, set in PatternFlowIpv6SRHPadTlvLength
	Choice() PatternFlowIpv6SRHPadTlvLengthChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHPadTlvLengthChoiceEnum provided by user to PatternFlowIpv6SRHPadTlvLength
	setChoice(value PatternFlowIpv6SRHPadTlvLengthChoiceEnum) PatternFlowIpv6SRHPadTlvLength
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHPadTlvLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHPadTlvLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHPadTlvLength
	SetValue(value uint32) PatternFlowIpv6SRHPadTlvLength
	// HasValue checks if Value has been set in PatternFlowIpv6SRHPadTlvLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHPadTlvLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHPadTlvLength
	SetValues(value []uint32) PatternFlowIpv6SRHPadTlvLength
	// Auto returns uint32, set in PatternFlowIpv6SRHPadTlvLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SRHPadTlvLength
	HasAuto() bool
	// Increment returns PatternFlowIpv6SRHPadTlvLengthCounter, set in PatternFlowIpv6SRHPadTlvLength.
	// PatternFlowIpv6SRHPadTlvLengthCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHPadTlvLengthCounter
	// SetIncrement assigns PatternFlowIpv6SRHPadTlvLengthCounter provided by user to PatternFlowIpv6SRHPadTlvLength.
	// PatternFlowIpv6SRHPadTlvLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHPadTlvLengthCounter) PatternFlowIpv6SRHPadTlvLength
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHPadTlvLength
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHPadTlvLengthCounter, set in PatternFlowIpv6SRHPadTlvLength.
	// PatternFlowIpv6SRHPadTlvLengthCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHPadTlvLengthCounter
	// SetDecrement assigns PatternFlowIpv6SRHPadTlvLengthCounter provided by user to PatternFlowIpv6SRHPadTlvLength.
	// PatternFlowIpv6SRHPadTlvLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHPadTlvLengthCounter) PatternFlowIpv6SRHPadTlvLength
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHPadTlvLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHPadTlvLengthChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHPadTlvLength
var PatternFlowIpv6SRHPadTlvLengthChoice = struct {
	VALUE     PatternFlowIpv6SRHPadTlvLengthChoiceEnum
	VALUES    PatternFlowIpv6SRHPadTlvLengthChoiceEnum
	AUTO      PatternFlowIpv6SRHPadTlvLengthChoiceEnum
	INCREMENT PatternFlowIpv6SRHPadTlvLengthChoiceEnum
	DECREMENT PatternFlowIpv6SRHPadTlvLengthChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHPadTlvLengthChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHPadTlvLengthChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SRHPadTlvLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SRHPadTlvLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHPadTlvLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHPadTlvLength) Choice() PatternFlowIpv6SRHPadTlvLengthChoiceEnum {
	return PatternFlowIpv6SRHPadTlvLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHPadTlvLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHPadTlvLength) setChoice(value PatternFlowIpv6SRHPadTlvLengthChoiceEnum) PatternFlowIpv6SRHPadTlvLength {
	intValue, ok := otg.PatternFlowIpv6SRHPadTlvLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHPadTlvLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHPadTlvLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHPadTlvLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHPadTlvLengthChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHPadTlvLengthChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SRHPadTlvLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHPadTlvLengthCounter().msg()
	}

	if value == PatternFlowIpv6SRHPadTlvLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHPadTlvLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPadTlvLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHPadTlvLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPadTlvLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHPadTlvLength object
func (obj *patternFlowIpv6SRHPadTlvLength) SetValue(value uint32) PatternFlowIpv6SRHPadTlvLength {
	obj.setChoice(PatternFlowIpv6SRHPadTlvLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHPadTlvLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHPadTlvLength object
func (obj *patternFlowIpv6SRHPadTlvLength) SetValues(value []uint32) PatternFlowIpv6SRHPadTlvLength {
	obj.setChoice(PatternFlowIpv6SRHPadTlvLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SRHPadTlvLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SRHPadTlvLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SRHPadTlvLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPadTlvLengthCounter
func (obj *patternFlowIpv6SRHPadTlvLength) Increment() PatternFlowIpv6SRHPadTlvLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHPadTlvLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHPadTlvLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPadTlvLengthCounter
func (obj *patternFlowIpv6SRHPadTlvLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHPadTlvLengthCounter value in the PatternFlowIpv6SRHPadTlvLength object
func (obj *patternFlowIpv6SRHPadTlvLength) SetIncrement(value PatternFlowIpv6SRHPadTlvLengthCounter) PatternFlowIpv6SRHPadTlvLength {
	obj.setChoice(PatternFlowIpv6SRHPadTlvLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPadTlvLengthCounter
func (obj *patternFlowIpv6SRHPadTlvLength) Decrement() PatternFlowIpv6SRHPadTlvLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHPadTlvLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHPadTlvLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPadTlvLengthCounter
func (obj *patternFlowIpv6SRHPadTlvLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHPadTlvLengthCounter value in the PatternFlowIpv6SRHPadTlvLength object
func (obj *patternFlowIpv6SRHPadTlvLength) SetDecrement(value PatternFlowIpv6SRHPadTlvLengthCounter) PatternFlowIpv6SRHPadTlvLength {
	obj.setChoice(PatternFlowIpv6SRHPadTlvLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHPadTlvLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPadTlvLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SRHPadTlvLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPadTlvLength.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHPadTlvLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHPadTlvLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPadTlvLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHPadTlvLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPadTlvLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPadTlvLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPadTlvLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHPadTlvLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHPadTlvLength")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHPadTlvLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHPadTlvLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
