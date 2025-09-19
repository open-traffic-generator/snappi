package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvsLength *****
type patternFlowIpv6SegmentRoutingTlvsLength struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingTlvsLength
	marshaller      marshalPatternFlowIpv6SegmentRoutingTlvsLength
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingTlvsLength
	incrementHolder PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	decrementHolder PatternFlowIpv6SegmentRoutingTlvsLengthCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvsLength() PatternFlowIpv6SegmentRoutingTlvsLength {
	obj := patternFlowIpv6SegmentRoutingTlvsLength{obj: &otg.PatternFlowIpv6SegmentRoutingTlvsLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLength) msg() *otg.PatternFlowIpv6SegmentRoutingTlvsLength {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLength) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvsLength) PatternFlowIpv6SegmentRoutingTlvsLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvsLength struct {
	obj *patternFlowIpv6SegmentRoutingTlvsLength
}

type marshalPatternFlowIpv6SegmentRoutingTlvsLength interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvsLength to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvsLength
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvsLength, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvsLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvsLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvsLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvsLength struct {
	obj *patternFlowIpv6SegmentRoutingTlvsLength
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvsLength interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvsLength from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvsLength
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvsLength) (PatternFlowIpv6SegmentRoutingTlvsLength, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvsLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvsLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvsLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLength) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvsLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvsLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLength) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvsLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvsLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvsLength) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvsLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvsLength) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvsLength) (PatternFlowIpv6SegmentRoutingTlvsLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvsLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvsLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvsLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvsLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvsLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvsLength) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvsLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLength) Clone() (PatternFlowIpv6SegmentRoutingTlvsLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvsLength()
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

func (obj *patternFlowIpv6SegmentRoutingTlvsLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingTlvsLength is the length of the TLV value in bytes. If auto is selected the implementation should automatically set the length field to the correct value as per the selected TLV and attributes.
type PatternFlowIpv6SegmentRoutingTlvsLength interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvsLength to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvsLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvsLength
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvsLength from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvsLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvsLength) PatternFlowIpv6SegmentRoutingTlvsLength
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvsLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvsLength
	// validate validates PatternFlowIpv6SegmentRoutingTlvsLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvsLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum, set in PatternFlowIpv6SegmentRoutingTlvsLength
	Choice() PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingTlvsLength
	setChoice(value PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum) PatternFlowIpv6SegmentRoutingTlvsLength
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingTlvsLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingTlvsLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvsLength
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvsLength
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingTlvsLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingTlvsLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvsLength
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvsLength
	// Auto returns uint32, set in PatternFlowIpv6SegmentRoutingTlvsLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SegmentRoutingTlvsLength
	HasAuto() bool
	// Increment returns PatternFlowIpv6SegmentRoutingTlvsLengthCounter, set in PatternFlowIpv6SegmentRoutingTlvsLength.
	// PatternFlowIpv6SegmentRoutingTlvsLengthCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingTlvsLengthCounter provided by user to PatternFlowIpv6SegmentRoutingTlvsLength.
	// PatternFlowIpv6SegmentRoutingTlvsLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingTlvsLengthCounter) PatternFlowIpv6SegmentRoutingTlvsLength
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingTlvsLength
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingTlvsLengthCounter, set in PatternFlowIpv6SegmentRoutingTlvsLength.
	// PatternFlowIpv6SegmentRoutingTlvsLengthCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingTlvsLengthCounter provided by user to PatternFlowIpv6SegmentRoutingTlvsLength.
	// PatternFlowIpv6SegmentRoutingTlvsLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingTlvsLengthCounter) PatternFlowIpv6SegmentRoutingTlvsLength
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingTlvsLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingTlvsLength
var PatternFlowIpv6SegmentRoutingTlvsLengthChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum
	AUTO      PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLength) Choice() PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum {
	return PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLength) setChoice(value PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum) PatternFlowIpv6SegmentRoutingTlvsLength {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingTlvsLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingTlvsLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingTlvsLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvsLengthChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvsLengthChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvsLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingTlvsLengthCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingTlvsLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingTlvsLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvsLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvsLength object
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvsLength {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvsLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingTlvsLength object
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvsLength {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvsLengthChoice.VALUES)
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
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvsLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvsLengthCounter
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) Increment() PatternFlowIpv6SegmentRoutingTlvsLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvsLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingTlvsLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvsLengthCounter
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingTlvsLengthCounter value in the PatternFlowIpv6SegmentRoutingTlvsLength object
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) SetIncrement(value PatternFlowIpv6SegmentRoutingTlvsLengthCounter) PatternFlowIpv6SegmentRoutingTlvsLength {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvsLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvsLengthCounter
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) Decrement() PatternFlowIpv6SegmentRoutingTlvsLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvsLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingTlvsLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvsLengthCounter
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingTlvsLengthCounter value in the PatternFlowIpv6SegmentRoutingTlvsLength object
func (obj *patternFlowIpv6SegmentRoutingTlvsLength) SetDecrement(value PatternFlowIpv6SegmentRoutingTlvsLengthCounter) PatternFlowIpv6SegmentRoutingTlvsLength {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvsLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvsLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingTlvsLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvsLength.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvsLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingTlvsLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvsLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvsLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvsLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvsLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvsLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingTlvsLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingTlvsLength")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingTlvsLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingTlvsLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
