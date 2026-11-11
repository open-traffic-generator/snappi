package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionLength *****
type patternFlowIcmpv6NdpPrefixInformationOptionLength struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength
	marshaller      marshalPatternFlowIcmpv6NdpPrefixInformationOptionLength
	unMarshaller    unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLength
	incrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	decrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionLength() PatternFlowIcmpv6NdpPrefixInformationOptionLength {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionLength{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength) PatternFlowIcmpv6NdpPrefixInformationOptionLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionLength struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionLength
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionLength interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionLength to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLength struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionLength
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLength interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLength from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength) (PatternFlowIcmpv6NdpPrefixInformationOptionLength, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLength) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLength) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength) (PatternFlowIcmpv6NdpPrefixInformationOptionLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLength) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionLength()
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpPrefixInformationOptionLength is option length in units of 8 octets. Fixed at 4 for the Prefix Information option.
type PatternFlowIcmpv6NdpPrefixInformationOptionLength interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionLength to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLength from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength) PatternFlowIcmpv6NdpPrefixInformationOptionLength
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLength
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum, set in PatternFlowIcmpv6NdpPrefixInformationOptionLength
	Choice() PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLength
	setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionLength
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLength
	SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLength
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLength
	SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLength
	// Auto returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLength
	HasAuto() bool
	// Increment returns PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionLength.
	// PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLength.
	// PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) PatternFlowIcmpv6NdpPrefixInformationOptionLength
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLength
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionLength.
	// PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLength.
	// PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) PatternFlowIcmpv6NdpPrefixInformationOptionLength
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpPrefixInformationOptionLength
var PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice = struct {
	VALUE     PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum
	VALUES    PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum
	AUTO      PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum("values"),
	AUTO:      PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) Choice() PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum {
	return PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionLength {
	intValue, ok := otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.VALUE {
		defaultValue := uint32(4)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.VALUES {
		defaultValue := []uint32{4}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.AUTO {
		defaultValue := uint32(4)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionLength object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{4})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionLength object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.VALUES)
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
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) Increment() PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionLength object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) PatternFlowIcmpv6NdpPrefixInformationOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionLength object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) PatternFlowIcmpv6NdpPrefixInformationOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6NdpPrefixInformationOptionLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionLength.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpPrefixInformationOptionLength")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
