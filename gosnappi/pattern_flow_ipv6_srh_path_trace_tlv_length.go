package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPathTraceTlvLength *****
type patternFlowIpv6SRHPathTraceTlvLength struct {
	validation
	obj             *otg.PatternFlowIpv6SRHPathTraceTlvLength
	marshaller      marshalPatternFlowIpv6SRHPathTraceTlvLength
	unMarshaller    unMarshalPatternFlowIpv6SRHPathTraceTlvLength
	incrementHolder PatternFlowIpv6SRHPathTraceTlvLengthCounter
	decrementHolder PatternFlowIpv6SRHPathTraceTlvLengthCounter
}

func NewPatternFlowIpv6SRHPathTraceTlvLength() PatternFlowIpv6SRHPathTraceTlvLength {
	obj := patternFlowIpv6SRHPathTraceTlvLength{obj: &otg.PatternFlowIpv6SRHPathTraceTlvLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvLength) msg() *otg.PatternFlowIpv6SRHPathTraceTlvLength {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvLength) setMsg(msg *otg.PatternFlowIpv6SRHPathTraceTlvLength) PatternFlowIpv6SRHPathTraceTlvLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPathTraceTlvLength struct {
	obj *patternFlowIpv6SRHPathTraceTlvLength
}

type marshalPatternFlowIpv6SRHPathTraceTlvLength interface {
	// ToProto marshals PatternFlowIpv6SRHPathTraceTlvLength to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvLength
	ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvLength, error)
	// ToPbText marshals PatternFlowIpv6SRHPathTraceTlvLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPathTraceTlvLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPathTraceTlvLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPathTraceTlvLength struct {
	obj *patternFlowIpv6SRHPathTraceTlvLength
}

type unMarshalPatternFlowIpv6SRHPathTraceTlvLength interface {
	// FromProto unmarshals PatternFlowIpv6SRHPathTraceTlvLength from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvLength
	FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvLength) (PatternFlowIpv6SRHPathTraceTlvLength, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPathTraceTlvLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPathTraceTlvLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPathTraceTlvLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPathTraceTlvLength) Marshal() marshalPatternFlowIpv6SRHPathTraceTlvLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPathTraceTlvLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPathTraceTlvLength) Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPathTraceTlvLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvLength) ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvLength) FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvLength) (PatternFlowIpv6SRHPathTraceTlvLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvLength) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPathTraceTlvLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPathTraceTlvLength) Clone() (PatternFlowIpv6SRHPathTraceTlvLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPathTraceTlvLength()
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

func (obj *patternFlowIpv6SRHPathTraceTlvLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHPathTraceTlvLength is length of the TLV Value field in bytes (draft-ietf-spring-srv6-path-tracing). When auto is assigned the implementation derives this from the configured value fields.
type PatternFlowIpv6SRHPathTraceTlvLength interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPathTraceTlvLength to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPathTraceTlvLength
	// setMsg unmarshals PatternFlowIpv6SRHPathTraceTlvLength from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPathTraceTlvLength) PatternFlowIpv6SRHPathTraceTlvLength
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPathTraceTlvLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvLength
	// validate validates PatternFlowIpv6SRHPathTraceTlvLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPathTraceTlvLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum, set in PatternFlowIpv6SRHPathTraceTlvLength
	Choice() PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum provided by user to PatternFlowIpv6SRHPathTraceTlvLength
	setChoice(value PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum) PatternFlowIpv6SRHPathTraceTlvLength
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHPathTraceTlvLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHPathTraceTlvLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvLength
	SetValue(value uint32) PatternFlowIpv6SRHPathTraceTlvLength
	// HasValue checks if Value has been set in PatternFlowIpv6SRHPathTraceTlvLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHPathTraceTlvLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvLength
	SetValues(value []uint32) PatternFlowIpv6SRHPathTraceTlvLength
	// Auto returns uint32, set in PatternFlowIpv6SRHPathTraceTlvLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SRHPathTraceTlvLength
	HasAuto() bool
	// Increment returns PatternFlowIpv6SRHPathTraceTlvLengthCounter, set in PatternFlowIpv6SRHPathTraceTlvLength.
	// PatternFlowIpv6SRHPathTraceTlvLengthCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHPathTraceTlvLengthCounter
	// SetIncrement assigns PatternFlowIpv6SRHPathTraceTlvLengthCounter provided by user to PatternFlowIpv6SRHPathTraceTlvLength.
	// PatternFlowIpv6SRHPathTraceTlvLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHPathTraceTlvLengthCounter) PatternFlowIpv6SRHPathTraceTlvLength
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHPathTraceTlvLength
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHPathTraceTlvLengthCounter, set in PatternFlowIpv6SRHPathTraceTlvLength.
	// PatternFlowIpv6SRHPathTraceTlvLengthCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHPathTraceTlvLengthCounter
	// SetDecrement assigns PatternFlowIpv6SRHPathTraceTlvLengthCounter provided by user to PatternFlowIpv6SRHPathTraceTlvLength.
	// PatternFlowIpv6SRHPathTraceTlvLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHPathTraceTlvLengthCounter) PatternFlowIpv6SRHPathTraceTlvLength
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHPathTraceTlvLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHPathTraceTlvLength
var PatternFlowIpv6SRHPathTraceTlvLengthChoice = struct {
	VALUE     PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum
	VALUES    PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum
	AUTO      PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum
	INCREMENT PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum
	DECREMENT PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHPathTraceTlvLength) Choice() PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum {
	return PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHPathTraceTlvLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHPathTraceTlvLength) setChoice(value PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum) PatternFlowIpv6SRHPathTraceTlvLength {
	intValue, ok := otg.PatternFlowIpv6SRHPathTraceTlvLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHPathTraceTlvLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHPathTraceTlvLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvLengthChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvLengthChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHPathTraceTlvLengthCounter().msg()
	}

	if value == PatternFlowIpv6SRHPathTraceTlvLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHPathTraceTlvLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvLength object
func (obj *patternFlowIpv6SRHPathTraceTlvLength) SetValue(value uint32) PatternFlowIpv6SRHPathTraceTlvLength {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHPathTraceTlvLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHPathTraceTlvLength object
func (obj *patternFlowIpv6SRHPathTraceTlvLength) SetValues(value []uint32) PatternFlowIpv6SRHPathTraceTlvLength {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvLengthChoice.VALUES)
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
func (obj *patternFlowIpv6SRHPathTraceTlvLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPathTraceTlvLengthCounter
func (obj *patternFlowIpv6SRHPathTraceTlvLength) Increment() PatternFlowIpv6SRHPathTraceTlvLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHPathTraceTlvLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPathTraceTlvLengthCounter
func (obj *patternFlowIpv6SRHPathTraceTlvLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHPathTraceTlvLengthCounter value in the PatternFlowIpv6SRHPathTraceTlvLength object
func (obj *patternFlowIpv6SRHPathTraceTlvLength) SetIncrement(value PatternFlowIpv6SRHPathTraceTlvLengthCounter) PatternFlowIpv6SRHPathTraceTlvLength {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPathTraceTlvLengthCounter
func (obj *patternFlowIpv6SRHPathTraceTlvLength) Decrement() PatternFlowIpv6SRHPathTraceTlvLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHPathTraceTlvLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPathTraceTlvLengthCounter
func (obj *patternFlowIpv6SRHPathTraceTlvLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHPathTraceTlvLengthCounter value in the PatternFlowIpv6SRHPathTraceTlvLength object
func (obj *patternFlowIpv6SRHPathTraceTlvLength) SetDecrement(value PatternFlowIpv6SRHPathTraceTlvLengthCounter) PatternFlowIpv6SRHPathTraceTlvLength {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SRHPathTraceTlvLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvLength.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHPathTraceTlvLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHPathTraceTlvLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHPathTraceTlvLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHPathTraceTlvLength")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHPathTraceTlvLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHPathTraceTlvLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
