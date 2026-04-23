package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmTlvsLength *****
type patternFlowCfmTlvsLength struct {
	validation
	obj             *otg.PatternFlowCfmTlvsLength
	marshaller      marshalPatternFlowCfmTlvsLength
	unMarshaller    unMarshalPatternFlowCfmTlvsLength
	incrementHolder PatternFlowCfmTlvsLengthCounter
	decrementHolder PatternFlowCfmTlvsLengthCounter
}

func NewPatternFlowCfmTlvsLength() PatternFlowCfmTlvsLength {
	obj := patternFlowCfmTlvsLength{obj: &otg.PatternFlowCfmTlvsLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmTlvsLength) msg() *otg.PatternFlowCfmTlvsLength {
	return obj.obj
}

func (obj *patternFlowCfmTlvsLength) setMsg(msg *otg.PatternFlowCfmTlvsLength) PatternFlowCfmTlvsLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmTlvsLength struct {
	obj *patternFlowCfmTlvsLength
}

type marshalPatternFlowCfmTlvsLength interface {
	// ToProto marshals PatternFlowCfmTlvsLength to protobuf object *otg.PatternFlowCfmTlvsLength
	ToProto() (*otg.PatternFlowCfmTlvsLength, error)
	// ToPbText marshals PatternFlowCfmTlvsLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmTlvsLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmTlvsLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmTlvsLength struct {
	obj *patternFlowCfmTlvsLength
}

type unMarshalPatternFlowCfmTlvsLength interface {
	// FromProto unmarshals PatternFlowCfmTlvsLength from protobuf object *otg.PatternFlowCfmTlvsLength
	FromProto(msg *otg.PatternFlowCfmTlvsLength) (PatternFlowCfmTlvsLength, error)
	// FromPbText unmarshals PatternFlowCfmTlvsLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmTlvsLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmTlvsLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmTlvsLength) Marshal() marshalPatternFlowCfmTlvsLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmTlvsLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmTlvsLength) Unmarshal() unMarshalPatternFlowCfmTlvsLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmTlvsLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmTlvsLength) ToProto() (*otg.PatternFlowCfmTlvsLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmTlvsLength) FromProto(msg *otg.PatternFlowCfmTlvsLength) (PatternFlowCfmTlvsLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmTlvsLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmTlvsLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmTlvsLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmTlvsLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmTlvsLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmTlvsLength) FromJson(value string) error {
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

func (obj *patternFlowCfmTlvsLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmTlvsLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmTlvsLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmTlvsLength) Clone() (PatternFlowCfmTlvsLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmTlvsLength()
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

func (obj *patternFlowCfmTlvsLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmTlvsLength is specifies the length of the tlv data in octets.
type PatternFlowCfmTlvsLength interface {
	Validation
	// msg marshals PatternFlowCfmTlvsLength to protobuf object *otg.PatternFlowCfmTlvsLength
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmTlvsLength
	// setMsg unmarshals PatternFlowCfmTlvsLength from protobuf object *otg.PatternFlowCfmTlvsLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmTlvsLength) PatternFlowCfmTlvsLength
	// provides marshal interface
	Marshal() marshalPatternFlowCfmTlvsLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmTlvsLength
	// validate validates PatternFlowCfmTlvsLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmTlvsLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmTlvsLengthChoiceEnum, set in PatternFlowCfmTlvsLength
	Choice() PatternFlowCfmTlvsLengthChoiceEnum
	// setChoice assigns PatternFlowCfmTlvsLengthChoiceEnum provided by user to PatternFlowCfmTlvsLength
	setChoice(value PatternFlowCfmTlvsLengthChoiceEnum) PatternFlowCfmTlvsLength
	// HasChoice checks if Choice has been set in PatternFlowCfmTlvsLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowCfmTlvsLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowCfmTlvsLength
	SetValue(value uint32) PatternFlowCfmTlvsLength
	// HasValue checks if Value has been set in PatternFlowCfmTlvsLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowCfmTlvsLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowCfmTlvsLength
	SetValues(value []uint32) PatternFlowCfmTlvsLength
	// Auto returns uint32, set in PatternFlowCfmTlvsLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowCfmTlvsLength
	HasAuto() bool
	// Increment returns PatternFlowCfmTlvsLengthCounter, set in PatternFlowCfmTlvsLength.
	// PatternFlowCfmTlvsLengthCounter is integer counter pattern
	Increment() PatternFlowCfmTlvsLengthCounter
	// SetIncrement assigns PatternFlowCfmTlvsLengthCounter provided by user to PatternFlowCfmTlvsLength.
	// PatternFlowCfmTlvsLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmTlvsLengthCounter) PatternFlowCfmTlvsLength
	// HasIncrement checks if Increment has been set in PatternFlowCfmTlvsLength
	HasIncrement() bool
	// Decrement returns PatternFlowCfmTlvsLengthCounter, set in PatternFlowCfmTlvsLength.
	// PatternFlowCfmTlvsLengthCounter is integer counter pattern
	Decrement() PatternFlowCfmTlvsLengthCounter
	// SetDecrement assigns PatternFlowCfmTlvsLengthCounter provided by user to PatternFlowCfmTlvsLength.
	// PatternFlowCfmTlvsLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmTlvsLengthCounter) PatternFlowCfmTlvsLength
	// HasDecrement checks if Decrement has been set in PatternFlowCfmTlvsLength
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmTlvsLengthChoiceEnum string

// Enum of Choice on PatternFlowCfmTlvsLength
var PatternFlowCfmTlvsLengthChoice = struct {
	VALUE     PatternFlowCfmTlvsLengthChoiceEnum
	VALUES    PatternFlowCfmTlvsLengthChoiceEnum
	AUTO      PatternFlowCfmTlvsLengthChoiceEnum
	INCREMENT PatternFlowCfmTlvsLengthChoiceEnum
	DECREMENT PatternFlowCfmTlvsLengthChoiceEnum
}{
	VALUE:     PatternFlowCfmTlvsLengthChoiceEnum("value"),
	VALUES:    PatternFlowCfmTlvsLengthChoiceEnum("values"),
	AUTO:      PatternFlowCfmTlvsLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowCfmTlvsLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmTlvsLengthChoiceEnum("decrement"),
}

func (obj *patternFlowCfmTlvsLength) Choice() PatternFlowCfmTlvsLengthChoiceEnum {
	return PatternFlowCfmTlvsLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmTlvsLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmTlvsLength) setChoice(value PatternFlowCfmTlvsLengthChoiceEnum) PatternFlowCfmTlvsLength {
	intValue, ok := otg.PatternFlowCfmTlvsLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmTlvsLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmTlvsLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmTlvsLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmTlvsLengthChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmTlvsLengthChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowCfmTlvsLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmTlvsLengthCounter().msg()
	}

	if value == PatternFlowCfmTlvsLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmTlvsLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmTlvsLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmTlvsLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmTlvsLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowCfmTlvsLength object
func (obj *patternFlowCfmTlvsLength) SetValue(value uint32) PatternFlowCfmTlvsLength {
	obj.setChoice(PatternFlowCfmTlvsLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowCfmTlvsLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowCfmTlvsLength object
func (obj *patternFlowCfmTlvsLength) SetValues(value []uint32) PatternFlowCfmTlvsLength {
	obj.setChoice(PatternFlowCfmTlvsLengthChoice.VALUES)
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
func (obj *patternFlowCfmTlvsLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowCfmTlvsLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowCfmTlvsLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowCfmTlvsLengthCounter
func (obj *patternFlowCfmTlvsLength) Increment() PatternFlowCfmTlvsLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmTlvsLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmTlvsLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmTlvsLengthCounter
func (obj *patternFlowCfmTlvsLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmTlvsLengthCounter value in the PatternFlowCfmTlvsLength object
func (obj *patternFlowCfmTlvsLength) SetIncrement(value PatternFlowCfmTlvsLengthCounter) PatternFlowCfmTlvsLength {
	obj.setChoice(PatternFlowCfmTlvsLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmTlvsLengthCounter
func (obj *patternFlowCfmTlvsLength) Decrement() PatternFlowCfmTlvsLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmTlvsLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmTlvsLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmTlvsLengthCounter
func (obj *patternFlowCfmTlvsLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmTlvsLengthCounter value in the PatternFlowCfmTlvsLength object
func (obj *patternFlowCfmTlvsLength) SetDecrement(value PatternFlowCfmTlvsLengthCounter) PatternFlowCfmTlvsLength {
	obj.setChoice(PatternFlowCfmTlvsLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmTlvsLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmTlvsLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowCfmTlvsLength.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmTlvsLength.Auto <= 65535 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowCfmTlvsLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmTlvsLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmTlvsLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmTlvsLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowCfmTlvsLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmTlvsLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmTlvsLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmTlvsLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmTlvsLength")
			}
		} else {
			intVal := otg.PatternFlowCfmTlvsLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmTlvsLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
