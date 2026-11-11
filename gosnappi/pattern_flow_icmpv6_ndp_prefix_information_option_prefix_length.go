package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength *****
type patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	marshaller      marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	unMarshaller    unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	incrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	decrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) (PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) (PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength()
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength is number of leading bits of the Prefix field that are valid. Range 0..128.
type PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	Choice() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	// Increment returns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
var PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice = struct {
	VALUE     PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum
	VALUES    PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) Choice() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum {
	return PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength {
	intValue, ok := otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.VALUE {
		defaultValue := uint32(64)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.VALUES {
		defaultValue := []uint32{64}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{64})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) Increment() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
