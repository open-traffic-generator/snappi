package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv2Type *****
type patternFlowIgmpv2Type struct {
	validation
	obj             *otg.PatternFlowIgmpv2Type
	marshaller      marshalPatternFlowIgmpv2Type
	unMarshaller    unMarshalPatternFlowIgmpv2Type
	incrementHolder PatternFlowIgmpv2TypeCounter
	decrementHolder PatternFlowIgmpv2TypeCounter
}

func NewPatternFlowIgmpv2Type() PatternFlowIgmpv2Type {
	obj := patternFlowIgmpv2Type{obj: &otg.PatternFlowIgmpv2Type{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv2Type) msg() *otg.PatternFlowIgmpv2Type {
	return obj.obj
}

func (obj *patternFlowIgmpv2Type) setMsg(msg *otg.PatternFlowIgmpv2Type) PatternFlowIgmpv2Type {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv2Type struct {
	obj *patternFlowIgmpv2Type
}

type marshalPatternFlowIgmpv2Type interface {
	// ToProto marshals PatternFlowIgmpv2Type to protobuf object *otg.PatternFlowIgmpv2Type
	ToProto() (*otg.PatternFlowIgmpv2Type, error)
	// ToPbText marshals PatternFlowIgmpv2Type to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv2Type to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv2Type to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv2Type struct {
	obj *patternFlowIgmpv2Type
}

type unMarshalPatternFlowIgmpv2Type interface {
	// FromProto unmarshals PatternFlowIgmpv2Type from protobuf object *otg.PatternFlowIgmpv2Type
	FromProto(msg *otg.PatternFlowIgmpv2Type) (PatternFlowIgmpv2Type, error)
	// FromPbText unmarshals PatternFlowIgmpv2Type from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv2Type from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv2Type from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv2Type) Marshal() marshalPatternFlowIgmpv2Type {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv2Type{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv2Type) Unmarshal() unMarshalPatternFlowIgmpv2Type {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv2Type{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv2Type) ToProto() (*otg.PatternFlowIgmpv2Type, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv2Type) FromProto(msg *otg.PatternFlowIgmpv2Type) (PatternFlowIgmpv2Type, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv2Type) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2Type) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv2Type) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2Type) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv2Type) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2Type) FromJson(value string) error {
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

func (obj *patternFlowIgmpv2Type) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2Type) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2Type) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv2Type) Clone() (PatternFlowIgmpv2Type, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv2Type()
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

func (obj *patternFlowIgmpv2Type) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv2Type is iGMPv2 message type
type PatternFlowIgmpv2Type interface {
	Validation
	// msg marshals PatternFlowIgmpv2Type to protobuf object *otg.PatternFlowIgmpv2Type
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv2Type
	// setMsg unmarshals PatternFlowIgmpv2Type from protobuf object *otg.PatternFlowIgmpv2Type
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv2Type) PatternFlowIgmpv2Type
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv2Type
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv2Type
	// validate validates PatternFlowIgmpv2Type
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv2Type, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv2TypeChoiceEnum, set in PatternFlowIgmpv2Type
	Choice() PatternFlowIgmpv2TypeChoiceEnum
	// setChoice assigns PatternFlowIgmpv2TypeChoiceEnum provided by user to PatternFlowIgmpv2Type
	setChoice(value PatternFlowIgmpv2TypeChoiceEnum) PatternFlowIgmpv2Type
	// HasChoice checks if Choice has been set in PatternFlowIgmpv2Type
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv2Type.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv2Type
	SetValue(value uint32) PatternFlowIgmpv2Type
	// HasValue checks if Value has been set in PatternFlowIgmpv2Type
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv2Type.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv2Type
	SetValues(value []uint32) PatternFlowIgmpv2Type
	// Increment returns PatternFlowIgmpv2TypeCounter, set in PatternFlowIgmpv2Type.
	// PatternFlowIgmpv2TypeCounter is integer counter pattern
	Increment() PatternFlowIgmpv2TypeCounter
	// SetIncrement assigns PatternFlowIgmpv2TypeCounter provided by user to PatternFlowIgmpv2Type.
	// PatternFlowIgmpv2TypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv2TypeCounter) PatternFlowIgmpv2Type
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv2Type
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv2TypeCounter, set in PatternFlowIgmpv2Type.
	// PatternFlowIgmpv2TypeCounter is integer counter pattern
	Decrement() PatternFlowIgmpv2TypeCounter
	// SetDecrement assigns PatternFlowIgmpv2TypeCounter provided by user to PatternFlowIgmpv2Type.
	// PatternFlowIgmpv2TypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv2TypeCounter) PatternFlowIgmpv2Type
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv2Type
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv2TypeChoiceEnum string

// Enum of Choice on PatternFlowIgmpv2Type
var PatternFlowIgmpv2TypeChoice = struct {
	VALUE     PatternFlowIgmpv2TypeChoiceEnum
	VALUES    PatternFlowIgmpv2TypeChoiceEnum
	INCREMENT PatternFlowIgmpv2TypeChoiceEnum
	DECREMENT PatternFlowIgmpv2TypeChoiceEnum
}{
	VALUE:     PatternFlowIgmpv2TypeChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv2TypeChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv2TypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv2TypeChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv2Type) Choice() PatternFlowIgmpv2TypeChoiceEnum {
	return PatternFlowIgmpv2TypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv2Type) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv2Type) setChoice(value PatternFlowIgmpv2TypeChoiceEnum) PatternFlowIgmpv2Type {
	intValue, ok := otg.PatternFlowIgmpv2Type_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv2TypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv2Type_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv2TypeChoice.VALUE {
		defaultValue := uint32(17)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv2TypeChoice.VALUES {
		defaultValue := []uint32{17}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv2TypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv2TypeCounter().msg()
	}

	if value == PatternFlowIgmpv2TypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv2TypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv2Type) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv2TypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv2Type) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv2Type object
func (obj *patternFlowIgmpv2Type) SetValue(value uint32) PatternFlowIgmpv2Type {
	obj.setChoice(PatternFlowIgmpv2TypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv2Type) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{17})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv2Type object
func (obj *patternFlowIgmpv2Type) SetValues(value []uint32) PatternFlowIgmpv2Type {
	obj.setChoice(PatternFlowIgmpv2TypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv2TypeCounter
func (obj *patternFlowIgmpv2Type) Increment() PatternFlowIgmpv2TypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv2TypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv2TypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv2TypeCounter
func (obj *patternFlowIgmpv2Type) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv2TypeCounter value in the PatternFlowIgmpv2Type object
func (obj *patternFlowIgmpv2Type) SetIncrement(value PatternFlowIgmpv2TypeCounter) PatternFlowIgmpv2Type {
	obj.setChoice(PatternFlowIgmpv2TypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv2TypeCounter
func (obj *patternFlowIgmpv2Type) Decrement() PatternFlowIgmpv2TypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv2TypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv2TypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv2TypeCounter
func (obj *patternFlowIgmpv2Type) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv2TypeCounter value in the PatternFlowIgmpv2Type object
func (obj *patternFlowIgmpv2Type) SetDecrement(value PatternFlowIgmpv2TypeCounter) PatternFlowIgmpv2Type {
	obj.setChoice(PatternFlowIgmpv2TypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv2Type) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv2Type.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv2Type.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIgmpv2Type) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv2TypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv2TypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv2TypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv2TypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv2TypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv2TypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv2Type")
			}
		} else {
			intVal := otg.PatternFlowIgmpv2Type_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv2Type_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
