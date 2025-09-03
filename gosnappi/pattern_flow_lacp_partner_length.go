package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerLength *****
type patternFlowLacpPartnerLength struct {
	validation
	obj             *otg.PatternFlowLacpPartnerLength
	marshaller      marshalPatternFlowLacpPartnerLength
	unMarshaller    unMarshalPatternFlowLacpPartnerLength
	incrementHolder PatternFlowLacpPartnerLengthCounter
	decrementHolder PatternFlowLacpPartnerLengthCounter
}

func NewPatternFlowLacpPartnerLength() PatternFlowLacpPartnerLength {
	obj := patternFlowLacpPartnerLength{obj: &otg.PatternFlowLacpPartnerLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerLength) msg() *otg.PatternFlowLacpPartnerLength {
	return obj.obj
}

func (obj *patternFlowLacpPartnerLength) setMsg(msg *otg.PatternFlowLacpPartnerLength) PatternFlowLacpPartnerLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerLength struct {
	obj *patternFlowLacpPartnerLength
}

type marshalPatternFlowLacpPartnerLength interface {
	// ToProto marshals PatternFlowLacpPartnerLength to protobuf object *otg.PatternFlowLacpPartnerLength
	ToProto() (*otg.PatternFlowLacpPartnerLength, error)
	// ToPbText marshals PatternFlowLacpPartnerLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerLength struct {
	obj *patternFlowLacpPartnerLength
}

type unMarshalPatternFlowLacpPartnerLength interface {
	// FromProto unmarshals PatternFlowLacpPartnerLength from protobuf object *otg.PatternFlowLacpPartnerLength
	FromProto(msg *otg.PatternFlowLacpPartnerLength) (PatternFlowLacpPartnerLength, error)
	// FromPbText unmarshals PatternFlowLacpPartnerLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerLength) Marshal() marshalPatternFlowLacpPartnerLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerLength) Unmarshal() unMarshalPatternFlowLacpPartnerLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerLength) ToProto() (*otg.PatternFlowLacpPartnerLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerLength) FromProto(msg *otg.PatternFlowLacpPartnerLength) (PatternFlowLacpPartnerLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerLength) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerLength) Clone() (PatternFlowLacpPartnerLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerLength()
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

func (obj *patternFlowLacpPartnerLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerLength is the length of the Partner Information TLV payload in bytes.
type PatternFlowLacpPartnerLength interface {
	Validation
	// msg marshals PatternFlowLacpPartnerLength to protobuf object *otg.PatternFlowLacpPartnerLength
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerLength
	// setMsg unmarshals PatternFlowLacpPartnerLength from protobuf object *otg.PatternFlowLacpPartnerLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerLength) PatternFlowLacpPartnerLength
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerLength
	// validate validates PatternFlowLacpPartnerLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerLengthChoiceEnum, set in PatternFlowLacpPartnerLength
	Choice() PatternFlowLacpPartnerLengthChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerLengthChoiceEnum provided by user to PatternFlowLacpPartnerLength
	setChoice(value PatternFlowLacpPartnerLengthChoiceEnum) PatternFlowLacpPartnerLength
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerLength
	SetValue(value uint32) PatternFlowLacpPartnerLength
	// HasValue checks if Value has been set in PatternFlowLacpPartnerLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerLength
	SetValues(value []uint32) PatternFlowLacpPartnerLength
	// Increment returns PatternFlowLacpPartnerLengthCounter, set in PatternFlowLacpPartnerLength.
	// PatternFlowLacpPartnerLengthCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerLengthCounter
	// SetIncrement assigns PatternFlowLacpPartnerLengthCounter provided by user to PatternFlowLacpPartnerLength.
	// PatternFlowLacpPartnerLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerLengthCounter) PatternFlowLacpPartnerLength
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerLength
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerLengthCounter, set in PatternFlowLacpPartnerLength.
	// PatternFlowLacpPartnerLengthCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerLengthCounter
	// SetDecrement assigns PatternFlowLacpPartnerLengthCounter provided by user to PatternFlowLacpPartnerLength.
	// PatternFlowLacpPartnerLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerLengthCounter) PatternFlowLacpPartnerLength
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerLength
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerLengthChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerLength
var PatternFlowLacpPartnerLengthChoice = struct {
	VALUE     PatternFlowLacpPartnerLengthChoiceEnum
	VALUES    PatternFlowLacpPartnerLengthChoiceEnum
	INCREMENT PatternFlowLacpPartnerLengthChoiceEnum
	DECREMENT PatternFlowLacpPartnerLengthChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerLengthChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerLengthChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerLengthChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerLength) Choice() PatternFlowLacpPartnerLengthChoiceEnum {
	return PatternFlowLacpPartnerLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerLength) setChoice(value PatternFlowLacpPartnerLengthChoiceEnum) PatternFlowLacpPartnerLength {
	intValue, ok := otg.PatternFlowLacpPartnerLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerLengthChoice.VALUE {
		defaultValue := uint32(20)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerLengthChoice.VALUES {
		defaultValue := []uint32{20}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerLengthCounter().msg()
	}

	if value == PatternFlowLacpPartnerLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerLength object
func (obj *patternFlowLacpPartnerLength) SetValue(value uint32) PatternFlowLacpPartnerLength {
	obj.setChoice(PatternFlowLacpPartnerLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{20})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerLength object
func (obj *patternFlowLacpPartnerLength) SetValues(value []uint32) PatternFlowLacpPartnerLength {
	obj.setChoice(PatternFlowLacpPartnerLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerLengthCounter
func (obj *patternFlowLacpPartnerLength) Increment() PatternFlowLacpPartnerLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerLengthCounter
func (obj *patternFlowLacpPartnerLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerLengthCounter value in the PatternFlowLacpPartnerLength object
func (obj *patternFlowLacpPartnerLength) SetIncrement(value PatternFlowLacpPartnerLengthCounter) PatternFlowLacpPartnerLength {
	obj.setChoice(PatternFlowLacpPartnerLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerLengthCounter
func (obj *patternFlowLacpPartnerLength) Decrement() PatternFlowLacpPartnerLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerLengthCounter
func (obj *patternFlowLacpPartnerLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerLengthCounter value in the PatternFlowLacpPartnerLength object
func (obj *patternFlowLacpPartnerLength) SetDecrement(value PatternFlowLacpPartnerLengthCounter) PatternFlowLacpPartnerLength {
	obj.setChoice(PatternFlowLacpPartnerLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpPartnerLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerLength")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
