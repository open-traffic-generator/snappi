package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHOpaqueContainerTlvLength *****
type patternFlowIpv6SRHOpaqueContainerTlvLength struct {
	validation
	obj             *otg.PatternFlowIpv6SRHOpaqueContainerTlvLength
	marshaller      marshalPatternFlowIpv6SRHOpaqueContainerTlvLength
	unMarshaller    unMarshalPatternFlowIpv6SRHOpaqueContainerTlvLength
	incrementHolder PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	decrementHolder PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
}

func NewPatternFlowIpv6SRHOpaqueContainerTlvLength() PatternFlowIpv6SRHOpaqueContainerTlvLength {
	obj := patternFlowIpv6SRHOpaqueContainerTlvLength{obj: &otg.PatternFlowIpv6SRHOpaqueContainerTlvLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) msg() *otg.PatternFlowIpv6SRHOpaqueContainerTlvLength {
	return obj.obj
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) setMsg(msg *otg.PatternFlowIpv6SRHOpaqueContainerTlvLength) PatternFlowIpv6SRHOpaqueContainerTlvLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHOpaqueContainerTlvLength struct {
	obj *patternFlowIpv6SRHOpaqueContainerTlvLength
}

type marshalPatternFlowIpv6SRHOpaqueContainerTlvLength interface {
	// ToProto marshals PatternFlowIpv6SRHOpaqueContainerTlvLength to protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvLength
	ToProto() (*otg.PatternFlowIpv6SRHOpaqueContainerTlvLength, error)
	// ToPbText marshals PatternFlowIpv6SRHOpaqueContainerTlvLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHOpaqueContainerTlvLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHOpaqueContainerTlvLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHOpaqueContainerTlvLength struct {
	obj *patternFlowIpv6SRHOpaqueContainerTlvLength
}

type unMarshalPatternFlowIpv6SRHOpaqueContainerTlvLength interface {
	// FromProto unmarshals PatternFlowIpv6SRHOpaqueContainerTlvLength from protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvLength
	FromProto(msg *otg.PatternFlowIpv6SRHOpaqueContainerTlvLength) (PatternFlowIpv6SRHOpaqueContainerTlvLength, error)
	// FromPbText unmarshals PatternFlowIpv6SRHOpaqueContainerTlvLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHOpaqueContainerTlvLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHOpaqueContainerTlvLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) Marshal() marshalPatternFlowIpv6SRHOpaqueContainerTlvLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHOpaqueContainerTlvLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) Unmarshal() unMarshalPatternFlowIpv6SRHOpaqueContainerTlvLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHOpaqueContainerTlvLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvLength) ToProto() (*otg.PatternFlowIpv6SRHOpaqueContainerTlvLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvLength) FromProto(msg *otg.PatternFlowIpv6SRHOpaqueContainerTlvLength) (PatternFlowIpv6SRHOpaqueContainerTlvLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvLength) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) Clone() (PatternFlowIpv6SRHOpaqueContainerTlvLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHOpaqueContainerTlvLength()
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

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHOpaqueContainerTlvLength is length of the TLV Value field in bytes (RFC 8754 Section 2.1). When auto is assigned the implementation derives this from the length of the value string.
type PatternFlowIpv6SRHOpaqueContainerTlvLength interface {
	Validation
	// msg marshals PatternFlowIpv6SRHOpaqueContainerTlvLength to protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHOpaqueContainerTlvLength
	// setMsg unmarshals PatternFlowIpv6SRHOpaqueContainerTlvLength from protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHOpaqueContainerTlvLength) PatternFlowIpv6SRHOpaqueContainerTlvLength
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHOpaqueContainerTlvLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHOpaqueContainerTlvLength
	// validate validates PatternFlowIpv6SRHOpaqueContainerTlvLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHOpaqueContainerTlvLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum, set in PatternFlowIpv6SRHOpaqueContainerTlvLength
	Choice() PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum provided by user to PatternFlowIpv6SRHOpaqueContainerTlvLength
	setChoice(value PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum) PatternFlowIpv6SRHOpaqueContainerTlvLength
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHOpaqueContainerTlvLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHOpaqueContainerTlvLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHOpaqueContainerTlvLength
	SetValue(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvLength
	// HasValue checks if Value has been set in PatternFlowIpv6SRHOpaqueContainerTlvLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHOpaqueContainerTlvLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHOpaqueContainerTlvLength
	SetValues(value []uint32) PatternFlowIpv6SRHOpaqueContainerTlvLength
	// Auto returns uint32, set in PatternFlowIpv6SRHOpaqueContainerTlvLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SRHOpaqueContainerTlvLength
	HasAuto() bool
	// Increment returns PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter, set in PatternFlowIpv6SRHOpaqueContainerTlvLength.
	// PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	// SetIncrement assigns PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter provided by user to PatternFlowIpv6SRHOpaqueContainerTlvLength.
	// PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) PatternFlowIpv6SRHOpaqueContainerTlvLength
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHOpaqueContainerTlvLength
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter, set in PatternFlowIpv6SRHOpaqueContainerTlvLength.
	// PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	// SetDecrement assigns PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter provided by user to PatternFlowIpv6SRHOpaqueContainerTlvLength.
	// PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) PatternFlowIpv6SRHOpaqueContainerTlvLength
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHOpaqueContainerTlvLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHOpaqueContainerTlvLength
var PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice = struct {
	VALUE     PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum
	VALUES    PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum
	AUTO      PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum
	INCREMENT PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum
	DECREMENT PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) Choice() PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum {
	return PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) setChoice(value PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum) PatternFlowIpv6SRHOpaqueContainerTlvLength {
	intValue, ok := otg.PatternFlowIpv6SRHOpaqueContainerTlvLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHOpaqueContainerTlvLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHOpaqueContainerTlvLengthCounter().msg()
	}

	if value == PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHOpaqueContainerTlvLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHOpaqueContainerTlvLength object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) SetValue(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvLength {
	obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHOpaqueContainerTlvLength object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) SetValues(value []uint32) PatternFlowIpv6SRHOpaqueContainerTlvLength {
	obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.VALUES)
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
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) Increment() PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHOpaqueContainerTlvLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter value in the PatternFlowIpv6SRHOpaqueContainerTlvLength object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) SetIncrement(value PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) PatternFlowIpv6SRHOpaqueContainerTlvLength {
	obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) Decrement() PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHOpaqueContainerTlvLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter value in the PatternFlowIpv6SRHOpaqueContainerTlvLength object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) SetDecrement(value PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) PatternFlowIpv6SRHOpaqueContainerTlvLength {
	obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHOpaqueContainerTlvLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SRHOpaqueContainerTlvLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHOpaqueContainerTlvLength.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHOpaqueContainerTlvLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHOpaqueContainerTlvLength")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHOpaqueContainerTlvLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHOpaqueContainerTlvLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
