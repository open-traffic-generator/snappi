package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmSenderIdTlvManagementAddressLength *****
type patternFlowCfmSenderIdTlvManagementAddressLength struct {
	validation
	obj             *otg.PatternFlowCfmSenderIdTlvManagementAddressLength
	marshaller      marshalPatternFlowCfmSenderIdTlvManagementAddressLength
	unMarshaller    unMarshalPatternFlowCfmSenderIdTlvManagementAddressLength
	incrementHolder PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	decrementHolder PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
}

func NewPatternFlowCfmSenderIdTlvManagementAddressLength() PatternFlowCfmSenderIdTlvManagementAddressLength {
	obj := patternFlowCfmSenderIdTlvManagementAddressLength{obj: &otg.PatternFlowCfmSenderIdTlvManagementAddressLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) msg() *otg.PatternFlowCfmSenderIdTlvManagementAddressLength {
	return obj.obj
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) setMsg(msg *otg.PatternFlowCfmSenderIdTlvManagementAddressLength) PatternFlowCfmSenderIdTlvManagementAddressLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmSenderIdTlvManagementAddressLength struct {
	obj *patternFlowCfmSenderIdTlvManagementAddressLength
}

type marshalPatternFlowCfmSenderIdTlvManagementAddressLength interface {
	// ToProto marshals PatternFlowCfmSenderIdTlvManagementAddressLength to protobuf object *otg.PatternFlowCfmSenderIdTlvManagementAddressLength
	ToProto() (*otg.PatternFlowCfmSenderIdTlvManagementAddressLength, error)
	// ToPbText marshals PatternFlowCfmSenderIdTlvManagementAddressLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmSenderIdTlvManagementAddressLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmSenderIdTlvManagementAddressLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmSenderIdTlvManagementAddressLength struct {
	obj *patternFlowCfmSenderIdTlvManagementAddressLength
}

type unMarshalPatternFlowCfmSenderIdTlvManagementAddressLength interface {
	// FromProto unmarshals PatternFlowCfmSenderIdTlvManagementAddressLength from protobuf object *otg.PatternFlowCfmSenderIdTlvManagementAddressLength
	FromProto(msg *otg.PatternFlowCfmSenderIdTlvManagementAddressLength) (PatternFlowCfmSenderIdTlvManagementAddressLength, error)
	// FromPbText unmarshals PatternFlowCfmSenderIdTlvManagementAddressLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmSenderIdTlvManagementAddressLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmSenderIdTlvManagementAddressLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) Marshal() marshalPatternFlowCfmSenderIdTlvManagementAddressLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmSenderIdTlvManagementAddressLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) Unmarshal() unMarshalPatternFlowCfmSenderIdTlvManagementAddressLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmSenderIdTlvManagementAddressLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmSenderIdTlvManagementAddressLength) ToProto() (*otg.PatternFlowCfmSenderIdTlvManagementAddressLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmSenderIdTlvManagementAddressLength) FromProto(msg *otg.PatternFlowCfmSenderIdTlvManagementAddressLength) (PatternFlowCfmSenderIdTlvManagementAddressLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmSenderIdTlvManagementAddressLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmSenderIdTlvManagementAddressLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmSenderIdTlvManagementAddressLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmSenderIdTlvManagementAddressLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmSenderIdTlvManagementAddressLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmSenderIdTlvManagementAddressLength) FromJson(value string) error {
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

func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) Clone() (PatternFlowCfmSenderIdTlvManagementAddressLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmSenderIdTlvManagementAddressLength()
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

func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmSenderIdTlvManagementAddressLength is specifies the length of the management address.
type PatternFlowCfmSenderIdTlvManagementAddressLength interface {
	Validation
	// msg marshals PatternFlowCfmSenderIdTlvManagementAddressLength to protobuf object *otg.PatternFlowCfmSenderIdTlvManagementAddressLength
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmSenderIdTlvManagementAddressLength
	// setMsg unmarshals PatternFlowCfmSenderIdTlvManagementAddressLength from protobuf object *otg.PatternFlowCfmSenderIdTlvManagementAddressLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmSenderIdTlvManagementAddressLength) PatternFlowCfmSenderIdTlvManagementAddressLength
	// provides marshal interface
	Marshal() marshalPatternFlowCfmSenderIdTlvManagementAddressLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmSenderIdTlvManagementAddressLength
	// validate validates PatternFlowCfmSenderIdTlvManagementAddressLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmSenderIdTlvManagementAddressLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum, set in PatternFlowCfmSenderIdTlvManagementAddressLength
	Choice() PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum
	// setChoice assigns PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum provided by user to PatternFlowCfmSenderIdTlvManagementAddressLength
	setChoice(value PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum) PatternFlowCfmSenderIdTlvManagementAddressLength
	// HasChoice checks if Choice has been set in PatternFlowCfmSenderIdTlvManagementAddressLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowCfmSenderIdTlvManagementAddressLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowCfmSenderIdTlvManagementAddressLength
	SetValue(value uint32) PatternFlowCfmSenderIdTlvManagementAddressLength
	// HasValue checks if Value has been set in PatternFlowCfmSenderIdTlvManagementAddressLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowCfmSenderIdTlvManagementAddressLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowCfmSenderIdTlvManagementAddressLength
	SetValues(value []uint32) PatternFlowCfmSenderIdTlvManagementAddressLength
	// Increment returns PatternFlowCfmSenderIdTlvManagementAddressLengthCounter, set in PatternFlowCfmSenderIdTlvManagementAddressLength.
	// PatternFlowCfmSenderIdTlvManagementAddressLengthCounter is integer counter pattern
	Increment() PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	// SetIncrement assigns PatternFlowCfmSenderIdTlvManagementAddressLengthCounter provided by user to PatternFlowCfmSenderIdTlvManagementAddressLength.
	// PatternFlowCfmSenderIdTlvManagementAddressLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmSenderIdTlvManagementAddressLengthCounter) PatternFlowCfmSenderIdTlvManagementAddressLength
	// HasIncrement checks if Increment has been set in PatternFlowCfmSenderIdTlvManagementAddressLength
	HasIncrement() bool
	// Decrement returns PatternFlowCfmSenderIdTlvManagementAddressLengthCounter, set in PatternFlowCfmSenderIdTlvManagementAddressLength.
	// PatternFlowCfmSenderIdTlvManagementAddressLengthCounter is integer counter pattern
	Decrement() PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
	// SetDecrement assigns PatternFlowCfmSenderIdTlvManagementAddressLengthCounter provided by user to PatternFlowCfmSenderIdTlvManagementAddressLength.
	// PatternFlowCfmSenderIdTlvManagementAddressLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmSenderIdTlvManagementAddressLengthCounter) PatternFlowCfmSenderIdTlvManagementAddressLength
	// HasDecrement checks if Decrement has been set in PatternFlowCfmSenderIdTlvManagementAddressLength
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum string

// Enum of Choice on PatternFlowCfmSenderIdTlvManagementAddressLength
var PatternFlowCfmSenderIdTlvManagementAddressLengthChoice = struct {
	VALUE     PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum
	VALUES    PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum
	INCREMENT PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum
	DECREMENT PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum
}{
	VALUE:     PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum("value"),
	VALUES:    PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum("values"),
	INCREMENT: PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum("decrement"),
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) Choice() PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum {
	return PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) setChoice(value PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum) PatternFlowCfmSenderIdTlvManagementAddressLength {
	intValue, ok := otg.PatternFlowCfmSenderIdTlvManagementAddressLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmSenderIdTlvManagementAddressLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmSenderIdTlvManagementAddressLengthCounter().msg()
	}

	if value == PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmSenderIdTlvManagementAddressLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowCfmSenderIdTlvManagementAddressLength object
func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) SetValue(value uint32) PatternFlowCfmSenderIdTlvManagementAddressLength {
	obj.setChoice(PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowCfmSenderIdTlvManagementAddressLength object
func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) SetValues(value []uint32) PatternFlowCfmSenderIdTlvManagementAddressLength {
	obj.setChoice(PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) Increment() PatternFlowCfmSenderIdTlvManagementAddressLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmSenderIdTlvManagementAddressLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmSenderIdTlvManagementAddressLengthCounter value in the PatternFlowCfmSenderIdTlvManagementAddressLength object
func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) SetIncrement(value PatternFlowCfmSenderIdTlvManagementAddressLengthCounter) PatternFlowCfmSenderIdTlvManagementAddressLength {
	obj.setChoice(PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) Decrement() PatternFlowCfmSenderIdTlvManagementAddressLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmSenderIdTlvManagementAddressLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmSenderIdTlvManagementAddressLengthCounter
func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmSenderIdTlvManagementAddressLengthCounter value in the PatternFlowCfmSenderIdTlvManagementAddressLength object
func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) SetDecrement(value PatternFlowCfmSenderIdTlvManagementAddressLengthCounter) PatternFlowCfmSenderIdTlvManagementAddressLength {
	obj.setChoice(PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmSenderIdTlvManagementAddressLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowCfmSenderIdTlvManagementAddressLength.Values <= 255 but Got %d", item))
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

func (obj *patternFlowCfmSenderIdTlvManagementAddressLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmSenderIdTlvManagementAddressLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmSenderIdTlvManagementAddressLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmSenderIdTlvManagementAddressLength")
			}
		} else {
			intVal := otg.PatternFlowCfmSenderIdTlvManagementAddressLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmSenderIdTlvManagementAddressLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
