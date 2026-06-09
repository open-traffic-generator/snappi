package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength *****
type patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	marshaller      marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	incrementHolder PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	decrementHolder PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
}

func NewPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength {
	obj := patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength{obj: &otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
}

type marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
}

type unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) (PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) (PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength()
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

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength is length of the Locator Block in bits (RFC 9800 Section 3).
// Valid range: 0-112.
// For NEXT-CSID (locator_length > 0): high-order locator_length bits of
// locator form the LB; usids are packed after it in forward order.
// For REPLACE-CSID first container (locator_length > 0): same as NEXT-CSID
// structure; use a single usid = Locator-Node+Function value (LNFL bits).
// For REPLACE-CSID packed containers (locator_length = 0): the locator
// field is ignored; the 128-bit SRH entry is K = floor(128/LNFL) slots
// of LNFL bits each; usids are provided in wire order (MSB first) - see usids
// field description for examples (RFC 9800 Section 4.2).
type PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	// validate validates PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	Choice() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	setChoice(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	// Increment returns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
var PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) Choice() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum {
	return PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) setChoice(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.VALUE {
		defaultValue := uint32(32)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.VALUES {
		defaultValue := []uint32{32}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) SetValue(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{32})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) Increment() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) SetIncrement(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) Decrement() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) SetDecrement(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
