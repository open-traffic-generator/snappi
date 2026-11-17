package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3GroupRecordAuxDataLength *****
type patternFlowIgmpv3GroupRecordAuxDataLength struct {
	validation
	obj             *otg.PatternFlowIgmpv3GroupRecordAuxDataLength
	marshaller      marshalPatternFlowIgmpv3GroupRecordAuxDataLength
	unMarshaller    unMarshalPatternFlowIgmpv3GroupRecordAuxDataLength
	incrementHolder PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	decrementHolder PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
}

func NewPatternFlowIgmpv3GroupRecordAuxDataLength() PatternFlowIgmpv3GroupRecordAuxDataLength {
	obj := patternFlowIgmpv3GroupRecordAuxDataLength{obj: &otg.PatternFlowIgmpv3GroupRecordAuxDataLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) msg() *otg.PatternFlowIgmpv3GroupRecordAuxDataLength {
	return obj.obj
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) setMsg(msg *otg.PatternFlowIgmpv3GroupRecordAuxDataLength) PatternFlowIgmpv3GroupRecordAuxDataLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3GroupRecordAuxDataLength struct {
	obj *patternFlowIgmpv3GroupRecordAuxDataLength
}

type marshalPatternFlowIgmpv3GroupRecordAuxDataLength interface {
	// ToProto marshals PatternFlowIgmpv3GroupRecordAuxDataLength to protobuf object *otg.PatternFlowIgmpv3GroupRecordAuxDataLength
	ToProto() (*otg.PatternFlowIgmpv3GroupRecordAuxDataLength, error)
	// ToPbText marshals PatternFlowIgmpv3GroupRecordAuxDataLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3GroupRecordAuxDataLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3GroupRecordAuxDataLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3GroupRecordAuxDataLength struct {
	obj *patternFlowIgmpv3GroupRecordAuxDataLength
}

type unMarshalPatternFlowIgmpv3GroupRecordAuxDataLength interface {
	// FromProto unmarshals PatternFlowIgmpv3GroupRecordAuxDataLength from protobuf object *otg.PatternFlowIgmpv3GroupRecordAuxDataLength
	FromProto(msg *otg.PatternFlowIgmpv3GroupRecordAuxDataLength) (PatternFlowIgmpv3GroupRecordAuxDataLength, error)
	// FromPbText unmarshals PatternFlowIgmpv3GroupRecordAuxDataLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3GroupRecordAuxDataLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3GroupRecordAuxDataLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) Marshal() marshalPatternFlowIgmpv3GroupRecordAuxDataLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3GroupRecordAuxDataLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordAuxDataLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3GroupRecordAuxDataLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3GroupRecordAuxDataLength) ToProto() (*otg.PatternFlowIgmpv3GroupRecordAuxDataLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3GroupRecordAuxDataLength) FromProto(msg *otg.PatternFlowIgmpv3GroupRecordAuxDataLength) (PatternFlowIgmpv3GroupRecordAuxDataLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3GroupRecordAuxDataLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordAuxDataLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordAuxDataLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordAuxDataLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordAuxDataLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordAuxDataLength) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) Clone() (PatternFlowIgmpv3GroupRecordAuxDataLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3GroupRecordAuxDataLength()
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

func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3GroupRecordAuxDataLength is length of aux_data in 32-bit words (RFC 9776 Section 4.2.7). When auto, it is set to the length of aux_data in 32-bit words, so zero when aux_data is omitted. A conforming sender uses zero.
type PatternFlowIgmpv3GroupRecordAuxDataLength interface {
	Validation
	// msg marshals PatternFlowIgmpv3GroupRecordAuxDataLength to protobuf object *otg.PatternFlowIgmpv3GroupRecordAuxDataLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3GroupRecordAuxDataLength
	// setMsg unmarshals PatternFlowIgmpv3GroupRecordAuxDataLength from protobuf object *otg.PatternFlowIgmpv3GroupRecordAuxDataLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3GroupRecordAuxDataLength) PatternFlowIgmpv3GroupRecordAuxDataLength
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3GroupRecordAuxDataLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordAuxDataLength
	// validate validates PatternFlowIgmpv3GroupRecordAuxDataLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3GroupRecordAuxDataLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum, set in PatternFlowIgmpv3GroupRecordAuxDataLength
	Choice() PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum
	// setChoice assigns PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum provided by user to PatternFlowIgmpv3GroupRecordAuxDataLength
	setChoice(value PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum) PatternFlowIgmpv3GroupRecordAuxDataLength
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3GroupRecordAuxDataLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3GroupRecordAuxDataLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3GroupRecordAuxDataLength
	SetValue(value uint32) PatternFlowIgmpv3GroupRecordAuxDataLength
	// HasValue checks if Value has been set in PatternFlowIgmpv3GroupRecordAuxDataLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3GroupRecordAuxDataLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3GroupRecordAuxDataLength
	SetValues(value []uint32) PatternFlowIgmpv3GroupRecordAuxDataLength
	// Auto returns uint32, set in PatternFlowIgmpv3GroupRecordAuxDataLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIgmpv3GroupRecordAuxDataLength
	HasAuto() bool
	// Increment returns PatternFlowIgmpv3GroupRecordAuxDataLengthCounter, set in PatternFlowIgmpv3GroupRecordAuxDataLength.
	// PatternFlowIgmpv3GroupRecordAuxDataLengthCounter is integer counter pattern
	Increment() PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	// SetIncrement assigns PatternFlowIgmpv3GroupRecordAuxDataLengthCounter provided by user to PatternFlowIgmpv3GroupRecordAuxDataLength.
	// PatternFlowIgmpv3GroupRecordAuxDataLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3GroupRecordAuxDataLengthCounter) PatternFlowIgmpv3GroupRecordAuxDataLength
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3GroupRecordAuxDataLength
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3GroupRecordAuxDataLengthCounter, set in PatternFlowIgmpv3GroupRecordAuxDataLength.
	// PatternFlowIgmpv3GroupRecordAuxDataLengthCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
	// SetDecrement assigns PatternFlowIgmpv3GroupRecordAuxDataLengthCounter provided by user to PatternFlowIgmpv3GroupRecordAuxDataLength.
	// PatternFlowIgmpv3GroupRecordAuxDataLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3GroupRecordAuxDataLengthCounter) PatternFlowIgmpv3GroupRecordAuxDataLength
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3GroupRecordAuxDataLength
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3GroupRecordAuxDataLength
var PatternFlowIgmpv3GroupRecordAuxDataLengthChoice = struct {
	VALUE     PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum
	VALUES    PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum
	AUTO      PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum
	INCREMENT PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum
	DECREMENT PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum("values"),
	AUTO:      PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) Choice() PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum {
	return PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) setChoice(value PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum) PatternFlowIgmpv3GroupRecordAuxDataLength {
	intValue, ok := otg.PatternFlowIgmpv3GroupRecordAuxDataLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3GroupRecordAuxDataLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3GroupRecordAuxDataLengthCounter().msg()
	}

	if value == PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3GroupRecordAuxDataLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3GroupRecordAuxDataLength object
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) SetValue(value uint32) PatternFlowIgmpv3GroupRecordAuxDataLength {
	obj.setChoice(PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3GroupRecordAuxDataLength object
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) SetValues(value []uint32) PatternFlowIgmpv3GroupRecordAuxDataLength {
	obj.setChoice(PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.VALUES)
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
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) Increment() PatternFlowIgmpv3GroupRecordAuxDataLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3GroupRecordAuxDataLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3GroupRecordAuxDataLengthCounter value in the PatternFlowIgmpv3GroupRecordAuxDataLength object
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) SetIncrement(value PatternFlowIgmpv3GroupRecordAuxDataLengthCounter) PatternFlowIgmpv3GroupRecordAuxDataLength {
	obj.setChoice(PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) Decrement() PatternFlowIgmpv3GroupRecordAuxDataLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3GroupRecordAuxDataLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3GroupRecordAuxDataLengthCounter
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3GroupRecordAuxDataLengthCounter value in the PatternFlowIgmpv3GroupRecordAuxDataLength object
func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) SetDecrement(value PatternFlowIgmpv3GroupRecordAuxDataLengthCounter) PatternFlowIgmpv3GroupRecordAuxDataLength {
	obj.setChoice(PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordAuxDataLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3GroupRecordAuxDataLength.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordAuxDataLength.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIgmpv3GroupRecordAuxDataLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3GroupRecordAuxDataLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3GroupRecordAuxDataLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3GroupRecordAuxDataLength")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3GroupRecordAuxDataLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3GroupRecordAuxDataLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
