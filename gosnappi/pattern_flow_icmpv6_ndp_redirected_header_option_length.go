package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpRedirectedHeaderOptionLength *****
type patternFlowIcmpv6NdpRedirectedHeaderOptionLength struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	marshaller      marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	unMarshaller    unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	incrementHolder PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	decrementHolder PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
}

func NewPatternFlowIcmpv6NdpRedirectedHeaderOptionLength() PatternFlowIcmpv6NdpRedirectedHeaderOptionLength {
	obj := patternFlowIcmpv6NdpRedirectedHeaderOptionLength{obj: &otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) msg() *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) setMsg(msg *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength) PatternFlowIcmpv6NdpRedirectedHeaderOptionLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLength struct {
	obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength
}

type marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLength interface {
	// ToProto marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLength to protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	ToProto() (*otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength, error)
	// ToPbText marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLength struct {
	obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength
}

type unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLength interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLength from protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	FromProto(msg *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength) (PatternFlowIcmpv6NdpRedirectedHeaderOptionLength, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) Marshal() marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) Unmarshal() unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLength) ToProto() (*otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLength) FromProto(msg *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength) (PatternFlowIcmpv6NdpRedirectedHeaderOptionLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLength) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) Clone() (PatternFlowIcmpv6NdpRedirectedHeaderOptionLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpRedirectedHeaderOptionLength()
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

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpRedirectedHeaderOptionLength is option length in units of 8 octets, including type/length bytes, reserved bytes and the truncated original packet.
type PatternFlowIcmpv6NdpRedirectedHeaderOptionLength interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLength to protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	// setMsg unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLength from protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength) PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	// validate validates PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpRedirectedHeaderOptionLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	Choice() PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	setChoice(value PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum) PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	SetValue(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	SetValues(value []uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	// Auto returns uint32, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	HasAuto() bool
	// Increment returns PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLength.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionLength.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLength.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionLength.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
var PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice = struct {
	VALUE     PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum
	VALUES    PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum
	AUTO      PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum("values"),
	AUTO:      PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) Choice() PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum {
	return PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) setChoice(value PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum) PatternFlowIcmpv6NdpRedirectedHeaderOptionLength {
	intValue, ok := otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.AUTO {
		defaultValue := uint32(1)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionLength object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) SetValue(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionLength object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) SetValues(value []uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.VALUES)
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
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) Increment() PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionLength object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) SetIncrement(value PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) PatternFlowIcmpv6NdpRedirectedHeaderOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) Decrement() PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionLength object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) SetDecrement(value PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) PatternFlowIcmpv6NdpRedirectedHeaderOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpRedirectedHeaderOptionLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6NdpRedirectedHeaderOptionLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpRedirectedHeaderOptionLength.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpRedirectedHeaderOptionLength")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
