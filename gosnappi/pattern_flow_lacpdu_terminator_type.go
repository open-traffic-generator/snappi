package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduTerminatorType *****
type patternFlowLacpduTerminatorType struct {
	validation
	obj             *otg.PatternFlowLacpduTerminatorType
	marshaller      marshalPatternFlowLacpduTerminatorType
	unMarshaller    unMarshalPatternFlowLacpduTerminatorType
	incrementHolder PatternFlowLacpduTerminatorTypeCounter
	decrementHolder PatternFlowLacpduTerminatorTypeCounter
}

func NewPatternFlowLacpduTerminatorType() PatternFlowLacpduTerminatorType {
	obj := patternFlowLacpduTerminatorType{obj: &otg.PatternFlowLacpduTerminatorType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduTerminatorType) msg() *otg.PatternFlowLacpduTerminatorType {
	return obj.obj
}

func (obj *patternFlowLacpduTerminatorType) setMsg(msg *otg.PatternFlowLacpduTerminatorType) PatternFlowLacpduTerminatorType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduTerminatorType struct {
	obj *patternFlowLacpduTerminatorType
}

type marshalPatternFlowLacpduTerminatorType interface {
	// ToProto marshals PatternFlowLacpduTerminatorType to protobuf object *otg.PatternFlowLacpduTerminatorType
	ToProto() (*otg.PatternFlowLacpduTerminatorType, error)
	// ToPbText marshals PatternFlowLacpduTerminatorType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduTerminatorType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduTerminatorType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduTerminatorType struct {
	obj *patternFlowLacpduTerminatorType
}

type unMarshalPatternFlowLacpduTerminatorType interface {
	// FromProto unmarshals PatternFlowLacpduTerminatorType from protobuf object *otg.PatternFlowLacpduTerminatorType
	FromProto(msg *otg.PatternFlowLacpduTerminatorType) (PatternFlowLacpduTerminatorType, error)
	// FromPbText unmarshals PatternFlowLacpduTerminatorType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduTerminatorType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduTerminatorType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduTerminatorType) Marshal() marshalPatternFlowLacpduTerminatorType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduTerminatorType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduTerminatorType) Unmarshal() unMarshalPatternFlowLacpduTerminatorType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduTerminatorType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduTerminatorType) ToProto() (*otg.PatternFlowLacpduTerminatorType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduTerminatorType) FromProto(msg *otg.PatternFlowLacpduTerminatorType) (PatternFlowLacpduTerminatorType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduTerminatorType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduTerminatorType) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduTerminatorType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduTerminatorType) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduTerminatorType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduTerminatorType) FromJson(value string) error {
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

func (obj *patternFlowLacpduTerminatorType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduTerminatorType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduTerminatorType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduTerminatorType) Clone() (PatternFlowLacpduTerminatorType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduTerminatorType()
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

func (obj *patternFlowLacpduTerminatorType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduTerminatorType is tLV Type for Terminator Information. The value 0x00 indicates the end of the TLV list.
type PatternFlowLacpduTerminatorType interface {
	Validation
	// msg marshals PatternFlowLacpduTerminatorType to protobuf object *otg.PatternFlowLacpduTerminatorType
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduTerminatorType
	// setMsg unmarshals PatternFlowLacpduTerminatorType from protobuf object *otg.PatternFlowLacpduTerminatorType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduTerminatorType) PatternFlowLacpduTerminatorType
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduTerminatorType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduTerminatorType
	// validate validates PatternFlowLacpduTerminatorType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduTerminatorType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduTerminatorTypeChoiceEnum, set in PatternFlowLacpduTerminatorType
	Choice() PatternFlowLacpduTerminatorTypeChoiceEnum
	// setChoice assigns PatternFlowLacpduTerminatorTypeChoiceEnum provided by user to PatternFlowLacpduTerminatorType
	setChoice(value PatternFlowLacpduTerminatorTypeChoiceEnum) PatternFlowLacpduTerminatorType
	// HasChoice checks if Choice has been set in PatternFlowLacpduTerminatorType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduTerminatorType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduTerminatorType
	SetValue(value uint32) PatternFlowLacpduTerminatorType
	// HasValue checks if Value has been set in PatternFlowLacpduTerminatorType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduTerminatorType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduTerminatorType
	SetValues(value []uint32) PatternFlowLacpduTerminatorType
	// Increment returns PatternFlowLacpduTerminatorTypeCounter, set in PatternFlowLacpduTerminatorType.
	// PatternFlowLacpduTerminatorTypeCounter is integer counter pattern
	Increment() PatternFlowLacpduTerminatorTypeCounter
	// SetIncrement assigns PatternFlowLacpduTerminatorTypeCounter provided by user to PatternFlowLacpduTerminatorType.
	// PatternFlowLacpduTerminatorTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduTerminatorTypeCounter) PatternFlowLacpduTerminatorType
	// HasIncrement checks if Increment has been set in PatternFlowLacpduTerminatorType
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduTerminatorTypeCounter, set in PatternFlowLacpduTerminatorType.
	// PatternFlowLacpduTerminatorTypeCounter is integer counter pattern
	Decrement() PatternFlowLacpduTerminatorTypeCounter
	// SetDecrement assigns PatternFlowLacpduTerminatorTypeCounter provided by user to PatternFlowLacpduTerminatorType.
	// PatternFlowLacpduTerminatorTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduTerminatorTypeCounter) PatternFlowLacpduTerminatorType
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduTerminatorType
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduTerminatorTypeChoiceEnum string

// Enum of Choice on PatternFlowLacpduTerminatorType
var PatternFlowLacpduTerminatorTypeChoice = struct {
	VALUE     PatternFlowLacpduTerminatorTypeChoiceEnum
	VALUES    PatternFlowLacpduTerminatorTypeChoiceEnum
	INCREMENT PatternFlowLacpduTerminatorTypeChoiceEnum
	DECREMENT PatternFlowLacpduTerminatorTypeChoiceEnum
}{
	VALUE:     PatternFlowLacpduTerminatorTypeChoiceEnum("value"),
	VALUES:    PatternFlowLacpduTerminatorTypeChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduTerminatorTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduTerminatorTypeChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduTerminatorType) Choice() PatternFlowLacpduTerminatorTypeChoiceEnum {
	return PatternFlowLacpduTerminatorTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduTerminatorType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduTerminatorType) setChoice(value PatternFlowLacpduTerminatorTypeChoiceEnum) PatternFlowLacpduTerminatorType {
	intValue, ok := otg.PatternFlowLacpduTerminatorType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduTerminatorTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduTerminatorType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduTerminatorTypeChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduTerminatorTypeChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduTerminatorTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduTerminatorTypeCounter().msg()
	}

	if value == PatternFlowLacpduTerminatorTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduTerminatorTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduTerminatorType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduTerminatorTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduTerminatorType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduTerminatorType object
func (obj *patternFlowLacpduTerminatorType) SetValue(value uint32) PatternFlowLacpduTerminatorType {
	obj.setChoice(PatternFlowLacpduTerminatorTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduTerminatorType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduTerminatorType object
func (obj *patternFlowLacpduTerminatorType) SetValues(value []uint32) PatternFlowLacpduTerminatorType {
	obj.setChoice(PatternFlowLacpduTerminatorTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduTerminatorTypeCounter
func (obj *patternFlowLacpduTerminatorType) Increment() PatternFlowLacpduTerminatorTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduTerminatorTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduTerminatorTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduTerminatorTypeCounter
func (obj *patternFlowLacpduTerminatorType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduTerminatorTypeCounter value in the PatternFlowLacpduTerminatorType object
func (obj *patternFlowLacpduTerminatorType) SetIncrement(value PatternFlowLacpduTerminatorTypeCounter) PatternFlowLacpduTerminatorType {
	obj.setChoice(PatternFlowLacpduTerminatorTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduTerminatorTypeCounter
func (obj *patternFlowLacpduTerminatorType) Decrement() PatternFlowLacpduTerminatorTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduTerminatorTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduTerminatorTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduTerminatorTypeCounter
func (obj *patternFlowLacpduTerminatorType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduTerminatorTypeCounter value in the PatternFlowLacpduTerminatorType object
func (obj *patternFlowLacpduTerminatorType) SetDecrement(value PatternFlowLacpduTerminatorTypeCounter) PatternFlowLacpduTerminatorType {
	obj.setChoice(PatternFlowLacpduTerminatorTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduTerminatorType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduTerminatorType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduTerminatorType.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpduTerminatorType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduTerminatorTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduTerminatorTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduTerminatorTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduTerminatorTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduTerminatorTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduTerminatorTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduTerminatorType")
			}
		} else {
			intVal := otg.PatternFlowLacpduTerminatorType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduTerminatorType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
