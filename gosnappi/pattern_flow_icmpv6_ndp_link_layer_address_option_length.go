package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpLinkLayerAddressOptionLength *****
type patternFlowIcmpv6NdpLinkLayerAddressOptionLength struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	marshaller      marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	unMarshaller    unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	incrementHolder PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	decrementHolder PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
}

func NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLength() PatternFlowIcmpv6NdpLinkLayerAddressOptionLength {
	obj := patternFlowIcmpv6NdpLinkLayerAddressOptionLength{obj: &otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) msg() *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) setMsg(msg *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength) PatternFlowIcmpv6NdpLinkLayerAddressOptionLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLength struct {
	obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength
}

type marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLength interface {
	// ToProto marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLength to protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	ToProto() (*otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength, error)
	// ToPbText marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLength struct {
	obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength
}

type unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLength interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLength from protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	FromProto(msg *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength) (PatternFlowIcmpv6NdpLinkLayerAddressOptionLength, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) Marshal() marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) Unmarshal() unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLength) ToProto() (*otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLength) FromProto(msg *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength) (PatternFlowIcmpv6NdpLinkLayerAddressOptionLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLength) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) Clone() (PatternFlowIcmpv6NdpLinkLayerAddressOptionLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLength()
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

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpLinkLayerAddressOptionLength is option length in units of 8 octets, including the type and length bytes. For a 6-octet Ethernet address this is 1.
type PatternFlowIcmpv6NdpLinkLayerAddressOptionLength interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLength to protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	// setMsg unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLength from protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength) PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	// validate validates PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpLinkLayerAddressOptionLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	Choice() PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	setChoice(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum) PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	SetValue(value uint32) PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	SetValues(value []uint32) PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	// Auto returns uint32, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	HasAuto() bool
	// Increment returns PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLength.
	// PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLength.
	// PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLength.
	// PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLength.
	// PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpLinkLayerAddressOptionLength
var PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice = struct {
	VALUE     PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum
	VALUES    PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum
	AUTO      PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum("values"),
	AUTO:      PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) Choice() PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum {
	return PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) setChoice(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum) PatternFlowIcmpv6NdpLinkLayerAddressOptionLength {
	intValue, ok := otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.AUTO {
		defaultValue := uint32(1)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLength object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) SetValue(value uint32) PatternFlowIcmpv6NdpLinkLayerAddressOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLength object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) SetValues(value []uint32) PatternFlowIcmpv6NdpLinkLayerAddressOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.VALUES)
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
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) Increment() PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLength object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) SetIncrement(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) PatternFlowIcmpv6NdpLinkLayerAddressOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) Decrement() PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLength object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) SetDecrement(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) PatternFlowIcmpv6NdpLinkLayerAddressOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpLinkLayerAddressOptionLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6NdpLinkLayerAddressOptionLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpLinkLayerAddressOptionLength.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpLinkLayerAddressOptionLength")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
