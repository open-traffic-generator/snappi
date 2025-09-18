package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerLength *****
type patternFlowLacpduPartnerLength struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerLength
	marshaller      marshalPatternFlowLacpduPartnerLength
	unMarshaller    unMarshalPatternFlowLacpduPartnerLength
	incrementHolder PatternFlowLacpduPartnerLengthCounter
	decrementHolder PatternFlowLacpduPartnerLengthCounter
}

func NewPatternFlowLacpduPartnerLength() PatternFlowLacpduPartnerLength {
	obj := patternFlowLacpduPartnerLength{obj: &otg.PatternFlowLacpduPartnerLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerLength) msg() *otg.PatternFlowLacpduPartnerLength {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerLength) setMsg(msg *otg.PatternFlowLacpduPartnerLength) PatternFlowLacpduPartnerLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerLength struct {
	obj *patternFlowLacpduPartnerLength
}

type marshalPatternFlowLacpduPartnerLength interface {
	// ToProto marshals PatternFlowLacpduPartnerLength to protobuf object *otg.PatternFlowLacpduPartnerLength
	ToProto() (*otg.PatternFlowLacpduPartnerLength, error)
	// ToPbText marshals PatternFlowLacpduPartnerLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerLength struct {
	obj *patternFlowLacpduPartnerLength
}

type unMarshalPatternFlowLacpduPartnerLength interface {
	// FromProto unmarshals PatternFlowLacpduPartnerLength from protobuf object *otg.PatternFlowLacpduPartnerLength
	FromProto(msg *otg.PatternFlowLacpduPartnerLength) (PatternFlowLacpduPartnerLength, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerLength) Marshal() marshalPatternFlowLacpduPartnerLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerLength) Unmarshal() unMarshalPatternFlowLacpduPartnerLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerLength) ToProto() (*otg.PatternFlowLacpduPartnerLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerLength) FromProto(msg *otg.PatternFlowLacpduPartnerLength) (PatternFlowLacpduPartnerLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerLength) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerLength) Clone() (PatternFlowLacpduPartnerLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerLength()
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

func (obj *patternFlowLacpduPartnerLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerLength is the length of the Partner Information TLV payload in bytes.
type PatternFlowLacpduPartnerLength interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerLength to protobuf object *otg.PatternFlowLacpduPartnerLength
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerLength
	// setMsg unmarshals PatternFlowLacpduPartnerLength from protobuf object *otg.PatternFlowLacpduPartnerLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerLength) PatternFlowLacpduPartnerLength
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerLength
	// validate validates PatternFlowLacpduPartnerLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerLengthChoiceEnum, set in PatternFlowLacpduPartnerLength
	Choice() PatternFlowLacpduPartnerLengthChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerLengthChoiceEnum provided by user to PatternFlowLacpduPartnerLength
	setChoice(value PatternFlowLacpduPartnerLengthChoiceEnum) PatternFlowLacpduPartnerLength
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerLength
	SetValue(value uint32) PatternFlowLacpduPartnerLength
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerLength
	SetValues(value []uint32) PatternFlowLacpduPartnerLength
	// Increment returns PatternFlowLacpduPartnerLengthCounter, set in PatternFlowLacpduPartnerLength.
	// PatternFlowLacpduPartnerLengthCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerLengthCounter
	// SetIncrement assigns PatternFlowLacpduPartnerLengthCounter provided by user to PatternFlowLacpduPartnerLength.
	// PatternFlowLacpduPartnerLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerLengthCounter) PatternFlowLacpduPartnerLength
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerLength
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerLengthCounter, set in PatternFlowLacpduPartnerLength.
	// PatternFlowLacpduPartnerLengthCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerLengthCounter
	// SetDecrement assigns PatternFlowLacpduPartnerLengthCounter provided by user to PatternFlowLacpduPartnerLength.
	// PatternFlowLacpduPartnerLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerLengthCounter) PatternFlowLacpduPartnerLength
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerLength
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerLengthChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerLength
var PatternFlowLacpduPartnerLengthChoice = struct {
	VALUE     PatternFlowLacpduPartnerLengthChoiceEnum
	VALUES    PatternFlowLacpduPartnerLengthChoiceEnum
	INCREMENT PatternFlowLacpduPartnerLengthChoiceEnum
	DECREMENT PatternFlowLacpduPartnerLengthChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerLengthChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerLengthChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerLengthChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerLength) Choice() PatternFlowLacpduPartnerLengthChoiceEnum {
	return PatternFlowLacpduPartnerLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerLength) setChoice(value PatternFlowLacpduPartnerLengthChoiceEnum) PatternFlowLacpduPartnerLength {
	intValue, ok := otg.PatternFlowLacpduPartnerLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerLengthChoice.VALUE {
		defaultValue := uint32(20)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerLengthChoice.VALUES {
		defaultValue := []uint32{20}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerLengthCounter().msg()
	}

	if value == PatternFlowLacpduPartnerLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerLength object
func (obj *patternFlowLacpduPartnerLength) SetValue(value uint32) PatternFlowLacpduPartnerLength {
	obj.setChoice(PatternFlowLacpduPartnerLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{20})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerLength object
func (obj *patternFlowLacpduPartnerLength) SetValues(value []uint32) PatternFlowLacpduPartnerLength {
	obj.setChoice(PatternFlowLacpduPartnerLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerLengthCounter
func (obj *patternFlowLacpduPartnerLength) Increment() PatternFlowLacpduPartnerLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerLengthCounter
func (obj *patternFlowLacpduPartnerLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerLengthCounter value in the PatternFlowLacpduPartnerLength object
func (obj *patternFlowLacpduPartnerLength) SetIncrement(value PatternFlowLacpduPartnerLengthCounter) PatternFlowLacpduPartnerLength {
	obj.setChoice(PatternFlowLacpduPartnerLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerLengthCounter
func (obj *patternFlowLacpduPartnerLength) Decrement() PatternFlowLacpduPartnerLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerLengthCounter
func (obj *patternFlowLacpduPartnerLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerLengthCounter value in the PatternFlowLacpduPartnerLength object
func (obj *patternFlowLacpduPartnerLength) SetDecrement(value PatternFlowLacpduPartnerLengthCounter) PatternFlowLacpduPartnerLength {
	obj.setChoice(PatternFlowLacpduPartnerLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerLength")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
