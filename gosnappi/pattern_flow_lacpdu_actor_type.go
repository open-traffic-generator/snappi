package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorType *****
type patternFlowLacpduActorType struct {
	validation
	obj             *otg.PatternFlowLacpduActorType
	marshaller      marshalPatternFlowLacpduActorType
	unMarshaller    unMarshalPatternFlowLacpduActorType
	incrementHolder PatternFlowLacpduActorTypeCounter
	decrementHolder PatternFlowLacpduActorTypeCounter
}

func NewPatternFlowLacpduActorType() PatternFlowLacpduActorType {
	obj := patternFlowLacpduActorType{obj: &otg.PatternFlowLacpduActorType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorType) msg() *otg.PatternFlowLacpduActorType {
	return obj.obj
}

func (obj *patternFlowLacpduActorType) setMsg(msg *otg.PatternFlowLacpduActorType) PatternFlowLacpduActorType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorType struct {
	obj *patternFlowLacpduActorType
}

type marshalPatternFlowLacpduActorType interface {
	// ToProto marshals PatternFlowLacpduActorType to protobuf object *otg.PatternFlowLacpduActorType
	ToProto() (*otg.PatternFlowLacpduActorType, error)
	// ToPbText marshals PatternFlowLacpduActorType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorType struct {
	obj *patternFlowLacpduActorType
}

type unMarshalPatternFlowLacpduActorType interface {
	// FromProto unmarshals PatternFlowLacpduActorType from protobuf object *otg.PatternFlowLacpduActorType
	FromProto(msg *otg.PatternFlowLacpduActorType) (PatternFlowLacpduActorType, error)
	// FromPbText unmarshals PatternFlowLacpduActorType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorType) Marshal() marshalPatternFlowLacpduActorType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorType) Unmarshal() unMarshalPatternFlowLacpduActorType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorType) ToProto() (*otg.PatternFlowLacpduActorType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorType) FromProto(msg *otg.PatternFlowLacpduActorType) (PatternFlowLacpduActorType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorType) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorType) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorType) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorType) Clone() (PatternFlowLacpduActorType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorType()
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

func (obj *patternFlowLacpduActorType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorType is tLV Type for Actor Information. The value 0x01 identifies this TLV as containing information about the sending device (the Actor).
type PatternFlowLacpduActorType interface {
	Validation
	// msg marshals PatternFlowLacpduActorType to protobuf object *otg.PatternFlowLacpduActorType
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorType
	// setMsg unmarshals PatternFlowLacpduActorType from protobuf object *otg.PatternFlowLacpduActorType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorType) PatternFlowLacpduActorType
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorType
	// validate validates PatternFlowLacpduActorType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorTypeChoiceEnum, set in PatternFlowLacpduActorType
	Choice() PatternFlowLacpduActorTypeChoiceEnum
	// setChoice assigns PatternFlowLacpduActorTypeChoiceEnum provided by user to PatternFlowLacpduActorType
	setChoice(value PatternFlowLacpduActorTypeChoiceEnum) PatternFlowLacpduActorType
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorType
	SetValue(value uint32) PatternFlowLacpduActorType
	// HasValue checks if Value has been set in PatternFlowLacpduActorType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorType
	SetValues(value []uint32) PatternFlowLacpduActorType
	// Increment returns PatternFlowLacpduActorTypeCounter, set in PatternFlowLacpduActorType.
	// PatternFlowLacpduActorTypeCounter is integer counter pattern
	Increment() PatternFlowLacpduActorTypeCounter
	// SetIncrement assigns PatternFlowLacpduActorTypeCounter provided by user to PatternFlowLacpduActorType.
	// PatternFlowLacpduActorTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorTypeCounter) PatternFlowLacpduActorType
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorType
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorTypeCounter, set in PatternFlowLacpduActorType.
	// PatternFlowLacpduActorTypeCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorTypeCounter
	// SetDecrement assigns PatternFlowLacpduActorTypeCounter provided by user to PatternFlowLacpduActorType.
	// PatternFlowLacpduActorTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorTypeCounter) PatternFlowLacpduActorType
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorType
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorTypeChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorType
var PatternFlowLacpduActorTypeChoice = struct {
	VALUE     PatternFlowLacpduActorTypeChoiceEnum
	VALUES    PatternFlowLacpduActorTypeChoiceEnum
	INCREMENT PatternFlowLacpduActorTypeChoiceEnum
	DECREMENT PatternFlowLacpduActorTypeChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorTypeChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorTypeChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorTypeChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorType) Choice() PatternFlowLacpduActorTypeChoiceEnum {
	return PatternFlowLacpduActorTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorType) setChoice(value PatternFlowLacpduActorTypeChoiceEnum) PatternFlowLacpduActorType {
	intValue, ok := otg.PatternFlowLacpduActorType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorTypeChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorTypeChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorTypeCounter().msg()
	}

	if value == PatternFlowLacpduActorTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorType object
func (obj *patternFlowLacpduActorType) SetValue(value uint32) PatternFlowLacpduActorType {
	obj.setChoice(PatternFlowLacpduActorTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorType object
func (obj *patternFlowLacpduActorType) SetValues(value []uint32) PatternFlowLacpduActorType {
	obj.setChoice(PatternFlowLacpduActorTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorTypeCounter
func (obj *patternFlowLacpduActorType) Increment() PatternFlowLacpduActorTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorTypeCounter
func (obj *patternFlowLacpduActorType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorTypeCounter value in the PatternFlowLacpduActorType object
func (obj *patternFlowLacpduActorType) SetIncrement(value PatternFlowLacpduActorTypeCounter) PatternFlowLacpduActorType {
	obj.setChoice(PatternFlowLacpduActorTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorTypeCounter
func (obj *patternFlowLacpduActorType) Decrement() PatternFlowLacpduActorTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorTypeCounter
func (obj *patternFlowLacpduActorType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorTypeCounter value in the PatternFlowLacpduActorType object
func (obj *patternFlowLacpduActorType) SetDecrement(value PatternFlowLacpduActorTypeCounter) PatternFlowLacpduActorType {
	obj.setChoice(PatternFlowLacpduActorTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorType.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpduActorType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorType")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
