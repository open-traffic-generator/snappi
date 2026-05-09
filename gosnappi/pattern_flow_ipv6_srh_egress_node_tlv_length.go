package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHEgressNodeTlvLength *****
type patternFlowIpv6SRHEgressNodeTlvLength struct {
	validation
	obj             *otg.PatternFlowIpv6SRHEgressNodeTlvLength
	marshaller      marshalPatternFlowIpv6SRHEgressNodeTlvLength
	unMarshaller    unMarshalPatternFlowIpv6SRHEgressNodeTlvLength
	incrementHolder PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	decrementHolder PatternFlowIpv6SRHEgressNodeTlvLengthCounter
}

func NewPatternFlowIpv6SRHEgressNodeTlvLength() PatternFlowIpv6SRHEgressNodeTlvLength {
	obj := patternFlowIpv6SRHEgressNodeTlvLength{obj: &otg.PatternFlowIpv6SRHEgressNodeTlvLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLength) msg() *otg.PatternFlowIpv6SRHEgressNodeTlvLength {
	return obj.obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLength) setMsg(msg *otg.PatternFlowIpv6SRHEgressNodeTlvLength) PatternFlowIpv6SRHEgressNodeTlvLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHEgressNodeTlvLength struct {
	obj *patternFlowIpv6SRHEgressNodeTlvLength
}

type marshalPatternFlowIpv6SRHEgressNodeTlvLength interface {
	// ToProto marshals PatternFlowIpv6SRHEgressNodeTlvLength to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvLength
	ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvLength, error)
	// ToPbText marshals PatternFlowIpv6SRHEgressNodeTlvLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHEgressNodeTlvLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHEgressNodeTlvLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHEgressNodeTlvLength struct {
	obj *patternFlowIpv6SRHEgressNodeTlvLength
}

type unMarshalPatternFlowIpv6SRHEgressNodeTlvLength interface {
	// FromProto unmarshals PatternFlowIpv6SRHEgressNodeTlvLength from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvLength
	FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvLength) (PatternFlowIpv6SRHEgressNodeTlvLength, error)
	// FromPbText unmarshals PatternFlowIpv6SRHEgressNodeTlvLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHEgressNodeTlvLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHEgressNodeTlvLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLength) Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHEgressNodeTlvLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLength) Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHEgressNodeTlvLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvLength) ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvLength) FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvLength) (PatternFlowIpv6SRHEgressNodeTlvLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvLength) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHEgressNodeTlvLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLength) Clone() (PatternFlowIpv6SRHEgressNodeTlvLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHEgressNodeTlvLength()
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

func (obj *patternFlowIpv6SRHEgressNodeTlvLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHEgressNodeTlvLength is length of the TLV Value field in bytes (RFC 9259 Section 3.2). When auto is assigned the implementation sets this to 18: 1-byte Reserved + 1-byte Node ID Len (fixed at 16) + 16-byte IPv6 address.
type PatternFlowIpv6SRHEgressNodeTlvLength interface {
	Validation
	// msg marshals PatternFlowIpv6SRHEgressNodeTlvLength to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHEgressNodeTlvLength
	// setMsg unmarshals PatternFlowIpv6SRHEgressNodeTlvLength from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHEgressNodeTlvLength) PatternFlowIpv6SRHEgressNodeTlvLength
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvLength
	// validate validates PatternFlowIpv6SRHEgressNodeTlvLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHEgressNodeTlvLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum, set in PatternFlowIpv6SRHEgressNodeTlvLength
	Choice() PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum provided by user to PatternFlowIpv6SRHEgressNodeTlvLength
	setChoice(value PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum) PatternFlowIpv6SRHEgressNodeTlvLength
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHEgressNodeTlvLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvLength
	SetValue(value uint32) PatternFlowIpv6SRHEgressNodeTlvLength
	// HasValue checks if Value has been set in PatternFlowIpv6SRHEgressNodeTlvLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHEgressNodeTlvLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvLength
	SetValues(value []uint32) PatternFlowIpv6SRHEgressNodeTlvLength
	// Auto returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SRHEgressNodeTlvLength
	HasAuto() bool
	// Increment returns PatternFlowIpv6SRHEgressNodeTlvLengthCounter, set in PatternFlowIpv6SRHEgressNodeTlvLength.
	// PatternFlowIpv6SRHEgressNodeTlvLengthCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	// SetIncrement assigns PatternFlowIpv6SRHEgressNodeTlvLengthCounter provided by user to PatternFlowIpv6SRHEgressNodeTlvLength.
	// PatternFlowIpv6SRHEgressNodeTlvLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHEgressNodeTlvLengthCounter) PatternFlowIpv6SRHEgressNodeTlvLength
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHEgressNodeTlvLength
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHEgressNodeTlvLengthCounter, set in PatternFlowIpv6SRHEgressNodeTlvLength.
	// PatternFlowIpv6SRHEgressNodeTlvLengthCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	// SetDecrement assigns PatternFlowIpv6SRHEgressNodeTlvLengthCounter provided by user to PatternFlowIpv6SRHEgressNodeTlvLength.
	// PatternFlowIpv6SRHEgressNodeTlvLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHEgressNodeTlvLengthCounter) PatternFlowIpv6SRHEgressNodeTlvLength
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHEgressNodeTlvLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHEgressNodeTlvLength
var PatternFlowIpv6SRHEgressNodeTlvLengthChoice = struct {
	VALUE     PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum
	VALUES    PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum
	AUTO      PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum
	INCREMENT PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum
	DECREMENT PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLength) Choice() PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum {
	return PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLength) setChoice(value PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum) PatternFlowIpv6SRHEgressNodeTlvLength {
	intValue, ok := otg.PatternFlowIpv6SRHEgressNodeTlvLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHEgressNodeTlvLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHEgressNodeTlvLengthChoice.VALUE {
		defaultValue := uint32(18)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvLengthChoice.VALUES {
		defaultValue := []uint32{18}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvLengthChoice.AUTO {
		defaultValue := uint32(18)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHEgressNodeTlvLengthCounter().msg()
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHEgressNodeTlvLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHEgressNodeTlvLength object
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) SetValue(value uint32) PatternFlowIpv6SRHEgressNodeTlvLength {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{18})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHEgressNodeTlvLength object
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) SetValues(value []uint32) PatternFlowIpv6SRHEgressNodeTlvLength {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvLengthChoice.VALUES)
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
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHEgressNodeTlvLengthCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) Increment() PatternFlowIpv6SRHEgressNodeTlvLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHEgressNodeTlvLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHEgressNodeTlvLengthCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHEgressNodeTlvLengthCounter value in the PatternFlowIpv6SRHEgressNodeTlvLength object
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) SetIncrement(value PatternFlowIpv6SRHEgressNodeTlvLengthCounter) PatternFlowIpv6SRHEgressNodeTlvLength {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHEgressNodeTlvLengthCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) Decrement() PatternFlowIpv6SRHEgressNodeTlvLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHEgressNodeTlvLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHEgressNodeTlvLengthCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHEgressNodeTlvLengthCounter value in the PatternFlowIpv6SRHEgressNodeTlvLength object
func (obj *patternFlowIpv6SRHEgressNodeTlvLength) SetDecrement(value PatternFlowIpv6SRHEgressNodeTlvLengthCounter) PatternFlowIpv6SRHEgressNodeTlvLength {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SRHEgressNodeTlvLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvLength.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHEgressNodeTlvLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHEgressNodeTlvLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHEgressNodeTlvLength")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHEgressNodeTlvLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHEgressNodeTlvLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
