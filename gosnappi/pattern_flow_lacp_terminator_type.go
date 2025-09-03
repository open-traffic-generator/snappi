package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpTerminatorType *****
type patternFlowLacpTerminatorType struct {
	validation
	obj             *otg.PatternFlowLacpTerminatorType
	marshaller      marshalPatternFlowLacpTerminatorType
	unMarshaller    unMarshalPatternFlowLacpTerminatorType
	incrementHolder PatternFlowLacpTerminatorTypeCounter
	decrementHolder PatternFlowLacpTerminatorTypeCounter
}

func NewPatternFlowLacpTerminatorType() PatternFlowLacpTerminatorType {
	obj := patternFlowLacpTerminatorType{obj: &otg.PatternFlowLacpTerminatorType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpTerminatorType) msg() *otg.PatternFlowLacpTerminatorType {
	return obj.obj
}

func (obj *patternFlowLacpTerminatorType) setMsg(msg *otg.PatternFlowLacpTerminatorType) PatternFlowLacpTerminatorType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpTerminatorType struct {
	obj *patternFlowLacpTerminatorType
}

type marshalPatternFlowLacpTerminatorType interface {
	// ToProto marshals PatternFlowLacpTerminatorType to protobuf object *otg.PatternFlowLacpTerminatorType
	ToProto() (*otg.PatternFlowLacpTerminatorType, error)
	// ToPbText marshals PatternFlowLacpTerminatorType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpTerminatorType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpTerminatorType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpTerminatorType struct {
	obj *patternFlowLacpTerminatorType
}

type unMarshalPatternFlowLacpTerminatorType interface {
	// FromProto unmarshals PatternFlowLacpTerminatorType from protobuf object *otg.PatternFlowLacpTerminatorType
	FromProto(msg *otg.PatternFlowLacpTerminatorType) (PatternFlowLacpTerminatorType, error)
	// FromPbText unmarshals PatternFlowLacpTerminatorType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpTerminatorType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpTerminatorType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpTerminatorType) Marshal() marshalPatternFlowLacpTerminatorType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpTerminatorType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpTerminatorType) Unmarshal() unMarshalPatternFlowLacpTerminatorType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpTerminatorType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpTerminatorType) ToProto() (*otg.PatternFlowLacpTerminatorType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpTerminatorType) FromProto(msg *otg.PatternFlowLacpTerminatorType) (PatternFlowLacpTerminatorType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpTerminatorType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpTerminatorType) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpTerminatorType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpTerminatorType) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpTerminatorType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpTerminatorType) FromJson(value string) error {
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

func (obj *patternFlowLacpTerminatorType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpTerminatorType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpTerminatorType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpTerminatorType) Clone() (PatternFlowLacpTerminatorType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpTerminatorType()
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

func (obj *patternFlowLacpTerminatorType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpTerminatorType is tLV Type for Terminator Information. The value 0x00 indicates the end of the TLV list.
type PatternFlowLacpTerminatorType interface {
	Validation
	// msg marshals PatternFlowLacpTerminatorType to protobuf object *otg.PatternFlowLacpTerminatorType
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpTerminatorType
	// setMsg unmarshals PatternFlowLacpTerminatorType from protobuf object *otg.PatternFlowLacpTerminatorType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpTerminatorType) PatternFlowLacpTerminatorType
	// provides marshal interface
	Marshal() marshalPatternFlowLacpTerminatorType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpTerminatorType
	// validate validates PatternFlowLacpTerminatorType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpTerminatorType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpTerminatorTypeChoiceEnum, set in PatternFlowLacpTerminatorType
	Choice() PatternFlowLacpTerminatorTypeChoiceEnum
	// setChoice assigns PatternFlowLacpTerminatorTypeChoiceEnum provided by user to PatternFlowLacpTerminatorType
	setChoice(value PatternFlowLacpTerminatorTypeChoiceEnum) PatternFlowLacpTerminatorType
	// HasChoice checks if Choice has been set in PatternFlowLacpTerminatorType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpTerminatorType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpTerminatorType
	SetValue(value uint32) PatternFlowLacpTerminatorType
	// HasValue checks if Value has been set in PatternFlowLacpTerminatorType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpTerminatorType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpTerminatorType
	SetValues(value []uint32) PatternFlowLacpTerminatorType
	// Increment returns PatternFlowLacpTerminatorTypeCounter, set in PatternFlowLacpTerminatorType.
	// PatternFlowLacpTerminatorTypeCounter is integer counter pattern
	Increment() PatternFlowLacpTerminatorTypeCounter
	// SetIncrement assigns PatternFlowLacpTerminatorTypeCounter provided by user to PatternFlowLacpTerminatorType.
	// PatternFlowLacpTerminatorTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpTerminatorTypeCounter) PatternFlowLacpTerminatorType
	// HasIncrement checks if Increment has been set in PatternFlowLacpTerminatorType
	HasIncrement() bool
	// Decrement returns PatternFlowLacpTerminatorTypeCounter, set in PatternFlowLacpTerminatorType.
	// PatternFlowLacpTerminatorTypeCounter is integer counter pattern
	Decrement() PatternFlowLacpTerminatorTypeCounter
	// SetDecrement assigns PatternFlowLacpTerminatorTypeCounter provided by user to PatternFlowLacpTerminatorType.
	// PatternFlowLacpTerminatorTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpTerminatorTypeCounter) PatternFlowLacpTerminatorType
	// HasDecrement checks if Decrement has been set in PatternFlowLacpTerminatorType
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpTerminatorTypeChoiceEnum string

// Enum of Choice on PatternFlowLacpTerminatorType
var PatternFlowLacpTerminatorTypeChoice = struct {
	VALUE     PatternFlowLacpTerminatorTypeChoiceEnum
	VALUES    PatternFlowLacpTerminatorTypeChoiceEnum
	INCREMENT PatternFlowLacpTerminatorTypeChoiceEnum
	DECREMENT PatternFlowLacpTerminatorTypeChoiceEnum
}{
	VALUE:     PatternFlowLacpTerminatorTypeChoiceEnum("value"),
	VALUES:    PatternFlowLacpTerminatorTypeChoiceEnum("values"),
	INCREMENT: PatternFlowLacpTerminatorTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpTerminatorTypeChoiceEnum("decrement"),
}

func (obj *patternFlowLacpTerminatorType) Choice() PatternFlowLacpTerminatorTypeChoiceEnum {
	return PatternFlowLacpTerminatorTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpTerminatorType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpTerminatorType) setChoice(value PatternFlowLacpTerminatorTypeChoiceEnum) PatternFlowLacpTerminatorType {
	intValue, ok := otg.PatternFlowLacpTerminatorType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpTerminatorTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpTerminatorType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpTerminatorTypeChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpTerminatorTypeChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpTerminatorTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpTerminatorTypeCounter().msg()
	}

	if value == PatternFlowLacpTerminatorTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpTerminatorTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpTerminatorType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpTerminatorTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpTerminatorType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpTerminatorType object
func (obj *patternFlowLacpTerminatorType) SetValue(value uint32) PatternFlowLacpTerminatorType {
	obj.setChoice(PatternFlowLacpTerminatorTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpTerminatorType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpTerminatorType object
func (obj *patternFlowLacpTerminatorType) SetValues(value []uint32) PatternFlowLacpTerminatorType {
	obj.setChoice(PatternFlowLacpTerminatorTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpTerminatorTypeCounter
func (obj *patternFlowLacpTerminatorType) Increment() PatternFlowLacpTerminatorTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpTerminatorTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpTerminatorTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpTerminatorTypeCounter
func (obj *patternFlowLacpTerminatorType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpTerminatorTypeCounter value in the PatternFlowLacpTerminatorType object
func (obj *patternFlowLacpTerminatorType) SetIncrement(value PatternFlowLacpTerminatorTypeCounter) PatternFlowLacpTerminatorType {
	obj.setChoice(PatternFlowLacpTerminatorTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpTerminatorTypeCounter
func (obj *patternFlowLacpTerminatorType) Decrement() PatternFlowLacpTerminatorTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpTerminatorTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpTerminatorTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpTerminatorTypeCounter
func (obj *patternFlowLacpTerminatorType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpTerminatorTypeCounter value in the PatternFlowLacpTerminatorType object
func (obj *patternFlowLacpTerminatorType) SetDecrement(value PatternFlowLacpTerminatorTypeCounter) PatternFlowLacpTerminatorType {
	obj.setChoice(PatternFlowLacpTerminatorTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpTerminatorType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpTerminatorType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpTerminatorType.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpTerminatorType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpTerminatorTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpTerminatorTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpTerminatorTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpTerminatorTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpTerminatorTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpTerminatorTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpTerminatorType")
			}
		} else {
			intVal := otg.PatternFlowLacpTerminatorType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpTerminatorType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
