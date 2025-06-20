package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength *****
type patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength struct {
	validation
	obj             *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	marshaller      marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	unMarshaller    unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	incrementHolder PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	decrementHolder PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
}

func NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength() PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength {
	obj := patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength{obj: &otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) msg() *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength {
	return obj.obj
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) setMsg(msg *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength struct {
	obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
}

type marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength interface {
	// ToProto marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength to protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	ToProto() (*otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength, error)
	// ToPbText marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength struct {
	obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
}

type unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength interface {
	// FromProto unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength from protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	FromProto(msg *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) (PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength, error)
	// FromPbText unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) Marshal() marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) Unmarshal() unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) ToProto() (*otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) FromProto(msg *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) (PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) Clone() (PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength()
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

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength is prefix-length of IPv4 address.
type PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength interface {
	Validation
	// msg marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength to protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	// setMsg unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength from protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	// validate validates PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	Choice() PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum
	// setChoice assigns PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	setChoice(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	SetValue(value uint32) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	// HasValue checks if Value has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	SetValues(value []uint32) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	// Increment returns PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength.
	// PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter is integer counter pattern
	Increment() PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	// SetIncrement assigns PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength.
	// PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength.
	// PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	// SetDecrement assigns PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength.
	// PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
var PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice = struct {
	VALUE     PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum
	VALUES    PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum
	INCREMENT PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum
	DECREMENT PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) Choice() PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum {
	return PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) setChoice(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength {
	intValue, ok := otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.VALUE {
		defaultValue := uint32(32)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.VALUES {
		defaultValue := []uint32{32}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter().msg()
	}

	if value == PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) SetValue(value uint32) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength {
	obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{32})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) SetValues(value []uint32) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength {
	obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) Increment() PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) SetIncrement(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength {
	obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) Decrement() PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) SetDecrement(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength {
	obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
