package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpMtuOptionLength *****
type patternFlowIcmpv6NdpMtuOptionLength struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpMtuOptionLength
	marshaller      marshalPatternFlowIcmpv6NdpMtuOptionLength
	unMarshaller    unMarshalPatternFlowIcmpv6NdpMtuOptionLength
	incrementHolder PatternFlowIcmpv6NdpMtuOptionLengthCounter
	decrementHolder PatternFlowIcmpv6NdpMtuOptionLengthCounter
}

func NewPatternFlowIcmpv6NdpMtuOptionLength() PatternFlowIcmpv6NdpMtuOptionLength {
	obj := patternFlowIcmpv6NdpMtuOptionLength{obj: &otg.PatternFlowIcmpv6NdpMtuOptionLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionLength) msg() *otg.PatternFlowIcmpv6NdpMtuOptionLength {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionLength) setMsg(msg *otg.PatternFlowIcmpv6NdpMtuOptionLength) PatternFlowIcmpv6NdpMtuOptionLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpMtuOptionLength struct {
	obj *patternFlowIcmpv6NdpMtuOptionLength
}

type marshalPatternFlowIcmpv6NdpMtuOptionLength interface {
	// ToProto marshals PatternFlowIcmpv6NdpMtuOptionLength to protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionLength
	ToProto() (*otg.PatternFlowIcmpv6NdpMtuOptionLength, error)
	// ToPbText marshals PatternFlowIcmpv6NdpMtuOptionLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpMtuOptionLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpMtuOptionLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpMtuOptionLength struct {
	obj *patternFlowIcmpv6NdpMtuOptionLength
}

type unMarshalPatternFlowIcmpv6NdpMtuOptionLength interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpMtuOptionLength from protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionLength
	FromProto(msg *otg.PatternFlowIcmpv6NdpMtuOptionLength) (PatternFlowIcmpv6NdpMtuOptionLength, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpMtuOptionLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpMtuOptionLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpMtuOptionLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpMtuOptionLength) Marshal() marshalPatternFlowIcmpv6NdpMtuOptionLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpMtuOptionLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpMtuOptionLength) Unmarshal() unMarshalPatternFlowIcmpv6NdpMtuOptionLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpMtuOptionLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpMtuOptionLength) ToProto() (*otg.PatternFlowIcmpv6NdpMtuOptionLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionLength) FromProto(msg *otg.PatternFlowIcmpv6NdpMtuOptionLength) (PatternFlowIcmpv6NdpMtuOptionLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpMtuOptionLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpMtuOptionLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpMtuOptionLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionLength) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpMtuOptionLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpMtuOptionLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpMtuOptionLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpMtuOptionLength) Clone() (PatternFlowIcmpv6NdpMtuOptionLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpMtuOptionLength()
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

func (obj *patternFlowIcmpv6NdpMtuOptionLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpMtuOptionLength is option length in units of 8 octets. Fixed at 1 for the MTU option.
type PatternFlowIcmpv6NdpMtuOptionLength interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpMtuOptionLength to protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpMtuOptionLength
	// setMsg unmarshals PatternFlowIcmpv6NdpMtuOptionLength from protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpMtuOptionLength) PatternFlowIcmpv6NdpMtuOptionLength
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpMtuOptionLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpMtuOptionLength
	// validate validates PatternFlowIcmpv6NdpMtuOptionLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpMtuOptionLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum, set in PatternFlowIcmpv6NdpMtuOptionLength
	Choice() PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum provided by user to PatternFlowIcmpv6NdpMtuOptionLength
	setChoice(value PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum) PatternFlowIcmpv6NdpMtuOptionLength
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpMtuOptionLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpMtuOptionLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionLength
	SetValue(value uint32) PatternFlowIcmpv6NdpMtuOptionLength
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpMtuOptionLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpMtuOptionLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionLength
	SetValues(value []uint32) PatternFlowIcmpv6NdpMtuOptionLength
	// Auto returns uint32, set in PatternFlowIcmpv6NdpMtuOptionLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIcmpv6NdpMtuOptionLength
	HasAuto() bool
	// Increment returns PatternFlowIcmpv6NdpMtuOptionLengthCounter, set in PatternFlowIcmpv6NdpMtuOptionLength.
	// PatternFlowIcmpv6NdpMtuOptionLengthCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpMtuOptionLengthCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpMtuOptionLengthCounter provided by user to PatternFlowIcmpv6NdpMtuOptionLength.
	// PatternFlowIcmpv6NdpMtuOptionLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpMtuOptionLengthCounter) PatternFlowIcmpv6NdpMtuOptionLength
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpMtuOptionLength
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpMtuOptionLengthCounter, set in PatternFlowIcmpv6NdpMtuOptionLength.
	// PatternFlowIcmpv6NdpMtuOptionLengthCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpMtuOptionLengthCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpMtuOptionLengthCounter provided by user to PatternFlowIcmpv6NdpMtuOptionLength.
	// PatternFlowIcmpv6NdpMtuOptionLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpMtuOptionLengthCounter) PatternFlowIcmpv6NdpMtuOptionLength
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpMtuOptionLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpMtuOptionLength
var PatternFlowIcmpv6NdpMtuOptionLengthChoice = struct {
	VALUE     PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum
	VALUES    PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum
	AUTO      PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum("values"),
	AUTO:      PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpMtuOptionLength) Choice() PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum {
	return PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpMtuOptionLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpMtuOptionLength) setChoice(value PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum) PatternFlowIcmpv6NdpMtuOptionLength {
	intValue, ok := otg.PatternFlowIcmpv6NdpMtuOptionLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpMtuOptionLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpMtuOptionLengthChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpMtuOptionLengthChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpMtuOptionLengthChoice.AUTO {
		defaultValue := uint32(1)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpMtuOptionLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpMtuOptionLengthCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpMtuOptionLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpMtuOptionLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpMtuOptionLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpMtuOptionLength object
func (obj *patternFlowIcmpv6NdpMtuOptionLength) SetValue(value uint32) PatternFlowIcmpv6NdpMtuOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpMtuOptionLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpMtuOptionLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpMtuOptionLength object
func (obj *patternFlowIcmpv6NdpMtuOptionLength) SetValues(value []uint32) PatternFlowIcmpv6NdpMtuOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpMtuOptionLengthChoice.VALUES)
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
func (obj *patternFlowIcmpv6NdpMtuOptionLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIcmpv6NdpMtuOptionLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpMtuOptionLengthCounter
func (obj *patternFlowIcmpv6NdpMtuOptionLength) Increment() PatternFlowIcmpv6NdpMtuOptionLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpMtuOptionLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpMtuOptionLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpMtuOptionLengthCounter
func (obj *patternFlowIcmpv6NdpMtuOptionLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpMtuOptionLengthCounter value in the PatternFlowIcmpv6NdpMtuOptionLength object
func (obj *patternFlowIcmpv6NdpMtuOptionLength) SetIncrement(value PatternFlowIcmpv6NdpMtuOptionLengthCounter) PatternFlowIcmpv6NdpMtuOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpMtuOptionLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpMtuOptionLengthCounter
func (obj *patternFlowIcmpv6NdpMtuOptionLength) Decrement() PatternFlowIcmpv6NdpMtuOptionLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpMtuOptionLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpMtuOptionLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpMtuOptionLengthCounter
func (obj *patternFlowIcmpv6NdpMtuOptionLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpMtuOptionLengthCounter value in the PatternFlowIcmpv6NdpMtuOptionLength object
func (obj *patternFlowIcmpv6NdpMtuOptionLength) SetDecrement(value PatternFlowIcmpv6NdpMtuOptionLengthCounter) PatternFlowIcmpv6NdpMtuOptionLength {
	obj.setChoice(PatternFlowIcmpv6NdpMtuOptionLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpMtuOptionLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6NdpMtuOptionLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpMtuOptionLength.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIcmpv6NdpMtuOptionLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpMtuOptionLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpMtuOptionLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpMtuOptionLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpMtuOptionLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpMtuOptionLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpMtuOptionLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpMtuOptionLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpMtuOptionLength")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpMtuOptionLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpMtuOptionLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
