package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHIngressNodeTlvLength *****
type patternFlowIpv6SRHIngressNodeTlvLength struct {
	validation
	obj             *otg.PatternFlowIpv6SRHIngressNodeTlvLength
	marshaller      marshalPatternFlowIpv6SRHIngressNodeTlvLength
	unMarshaller    unMarshalPatternFlowIpv6SRHIngressNodeTlvLength
	incrementHolder PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	decrementHolder PatternFlowIpv6SRHIngressNodeTlvLengthCounter
}

func NewPatternFlowIpv6SRHIngressNodeTlvLength() PatternFlowIpv6SRHIngressNodeTlvLength {
	obj := patternFlowIpv6SRHIngressNodeTlvLength{obj: &otg.PatternFlowIpv6SRHIngressNodeTlvLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLength) msg() *otg.PatternFlowIpv6SRHIngressNodeTlvLength {
	return obj.obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLength) setMsg(msg *otg.PatternFlowIpv6SRHIngressNodeTlvLength) PatternFlowIpv6SRHIngressNodeTlvLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHIngressNodeTlvLength struct {
	obj *patternFlowIpv6SRHIngressNodeTlvLength
}

type marshalPatternFlowIpv6SRHIngressNodeTlvLength interface {
	// ToProto marshals PatternFlowIpv6SRHIngressNodeTlvLength to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvLength
	ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvLength, error)
	// ToPbText marshals PatternFlowIpv6SRHIngressNodeTlvLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHIngressNodeTlvLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHIngressNodeTlvLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHIngressNodeTlvLength struct {
	obj *patternFlowIpv6SRHIngressNodeTlvLength
}

type unMarshalPatternFlowIpv6SRHIngressNodeTlvLength interface {
	// FromProto unmarshals PatternFlowIpv6SRHIngressNodeTlvLength from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvLength
	FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvLength) (PatternFlowIpv6SRHIngressNodeTlvLength, error)
	// FromPbText unmarshals PatternFlowIpv6SRHIngressNodeTlvLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHIngressNodeTlvLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHIngressNodeTlvLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLength) Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHIngressNodeTlvLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLength) Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHIngressNodeTlvLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvLength) ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvLength) FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvLength) (PatternFlowIpv6SRHIngressNodeTlvLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvLength) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHIngressNodeTlvLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLength) Clone() (PatternFlowIpv6SRHIngressNodeTlvLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHIngressNodeTlvLength()
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

func (obj *patternFlowIpv6SRHIngressNodeTlvLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHIngressNodeTlvLength is length of the TLV Value field in bytes (RFC 9259 Section 3.1). When auto is assigned the implementation sets this to 18: 1-byte Reserved + 1-byte Node ID Len (fixed at 16) + 16-byte IPv6 address.
type PatternFlowIpv6SRHIngressNodeTlvLength interface {
	Validation
	// msg marshals PatternFlowIpv6SRHIngressNodeTlvLength to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHIngressNodeTlvLength
	// setMsg unmarshals PatternFlowIpv6SRHIngressNodeTlvLength from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHIngressNodeTlvLength) PatternFlowIpv6SRHIngressNodeTlvLength
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvLength
	// validate validates PatternFlowIpv6SRHIngressNodeTlvLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHIngressNodeTlvLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum, set in PatternFlowIpv6SRHIngressNodeTlvLength
	Choice() PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum provided by user to PatternFlowIpv6SRHIngressNodeTlvLength
	setChoice(value PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum) PatternFlowIpv6SRHIngressNodeTlvLength
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHIngressNodeTlvLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvLength
	SetValue(value uint32) PatternFlowIpv6SRHIngressNodeTlvLength
	// HasValue checks if Value has been set in PatternFlowIpv6SRHIngressNodeTlvLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHIngressNodeTlvLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvLength
	SetValues(value []uint32) PatternFlowIpv6SRHIngressNodeTlvLength
	// Auto returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SRHIngressNodeTlvLength
	HasAuto() bool
	// Increment returns PatternFlowIpv6SRHIngressNodeTlvLengthCounter, set in PatternFlowIpv6SRHIngressNodeTlvLength.
	// PatternFlowIpv6SRHIngressNodeTlvLengthCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	// SetIncrement assigns PatternFlowIpv6SRHIngressNodeTlvLengthCounter provided by user to PatternFlowIpv6SRHIngressNodeTlvLength.
	// PatternFlowIpv6SRHIngressNodeTlvLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHIngressNodeTlvLengthCounter) PatternFlowIpv6SRHIngressNodeTlvLength
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHIngressNodeTlvLength
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHIngressNodeTlvLengthCounter, set in PatternFlowIpv6SRHIngressNodeTlvLength.
	// PatternFlowIpv6SRHIngressNodeTlvLengthCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	// SetDecrement assigns PatternFlowIpv6SRHIngressNodeTlvLengthCounter provided by user to PatternFlowIpv6SRHIngressNodeTlvLength.
	// PatternFlowIpv6SRHIngressNodeTlvLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHIngressNodeTlvLengthCounter) PatternFlowIpv6SRHIngressNodeTlvLength
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHIngressNodeTlvLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHIngressNodeTlvLength
var PatternFlowIpv6SRHIngressNodeTlvLengthChoice = struct {
	VALUE     PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum
	VALUES    PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum
	AUTO      PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum
	INCREMENT PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum
	DECREMENT PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLength) Choice() PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum {
	return PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLength) setChoice(value PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum) PatternFlowIpv6SRHIngressNodeTlvLength {
	intValue, ok := otg.PatternFlowIpv6SRHIngressNodeTlvLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHIngressNodeTlvLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHIngressNodeTlvLengthChoice.VALUE {
		defaultValue := uint32(18)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvLengthChoice.VALUES {
		defaultValue := []uint32{18}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvLengthChoice.AUTO {
		defaultValue := uint32(18)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHIngressNodeTlvLengthCounter().msg()
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHIngressNodeTlvLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHIngressNodeTlvLength object
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) SetValue(value uint32) PatternFlowIpv6SRHIngressNodeTlvLength {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{18})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHIngressNodeTlvLength object
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) SetValues(value []uint32) PatternFlowIpv6SRHIngressNodeTlvLength {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvLengthChoice.VALUES)
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
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHIngressNodeTlvLengthCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) Increment() PatternFlowIpv6SRHIngressNodeTlvLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHIngressNodeTlvLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHIngressNodeTlvLengthCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHIngressNodeTlvLengthCounter value in the PatternFlowIpv6SRHIngressNodeTlvLength object
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) SetIncrement(value PatternFlowIpv6SRHIngressNodeTlvLengthCounter) PatternFlowIpv6SRHIngressNodeTlvLength {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHIngressNodeTlvLengthCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) Decrement() PatternFlowIpv6SRHIngressNodeTlvLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHIngressNodeTlvLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHIngressNodeTlvLengthCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHIngressNodeTlvLengthCounter value in the PatternFlowIpv6SRHIngressNodeTlvLength object
func (obj *patternFlowIpv6SRHIngressNodeTlvLength) SetDecrement(value PatternFlowIpv6SRHIngressNodeTlvLengthCounter) PatternFlowIpv6SRHIngressNodeTlvLength {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SRHIngressNodeTlvLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvLength.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHIngressNodeTlvLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHIngressNodeTlvLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHIngressNodeTlvLength")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHIngressNodeTlvLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHIngressNodeTlvLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
