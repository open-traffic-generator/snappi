package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorType *****
type patternFlowLacpActorType struct {
	validation
	obj             *otg.PatternFlowLacpActorType
	marshaller      marshalPatternFlowLacpActorType
	unMarshaller    unMarshalPatternFlowLacpActorType
	incrementHolder PatternFlowLacpActorTypeCounter
	decrementHolder PatternFlowLacpActorTypeCounter
}

func NewPatternFlowLacpActorType() PatternFlowLacpActorType {
	obj := patternFlowLacpActorType{obj: &otg.PatternFlowLacpActorType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorType) msg() *otg.PatternFlowLacpActorType {
	return obj.obj
}

func (obj *patternFlowLacpActorType) setMsg(msg *otg.PatternFlowLacpActorType) PatternFlowLacpActorType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorType struct {
	obj *patternFlowLacpActorType
}

type marshalPatternFlowLacpActorType interface {
	// ToProto marshals PatternFlowLacpActorType to protobuf object *otg.PatternFlowLacpActorType
	ToProto() (*otg.PatternFlowLacpActorType, error)
	// ToPbText marshals PatternFlowLacpActorType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorType struct {
	obj *patternFlowLacpActorType
}

type unMarshalPatternFlowLacpActorType interface {
	// FromProto unmarshals PatternFlowLacpActorType from protobuf object *otg.PatternFlowLacpActorType
	FromProto(msg *otg.PatternFlowLacpActorType) (PatternFlowLacpActorType, error)
	// FromPbText unmarshals PatternFlowLacpActorType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorType) Marshal() marshalPatternFlowLacpActorType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorType) Unmarshal() unMarshalPatternFlowLacpActorType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorType) ToProto() (*otg.PatternFlowLacpActorType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorType) FromProto(msg *otg.PatternFlowLacpActorType) (PatternFlowLacpActorType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorType) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorType) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorType) FromJson(value string) error {
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

func (obj *patternFlowLacpActorType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorType) Clone() (PatternFlowLacpActorType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorType()
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

func (obj *patternFlowLacpActorType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorType is tLV Type for Actor Information. The value 0x01 identifies this TLV as containing information about the sending device (the Actor).
type PatternFlowLacpActorType interface {
	Validation
	// msg marshals PatternFlowLacpActorType to protobuf object *otg.PatternFlowLacpActorType
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorType
	// setMsg unmarshals PatternFlowLacpActorType from protobuf object *otg.PatternFlowLacpActorType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorType) PatternFlowLacpActorType
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorType
	// validate validates PatternFlowLacpActorType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorTypeChoiceEnum, set in PatternFlowLacpActorType
	Choice() PatternFlowLacpActorTypeChoiceEnum
	// setChoice assigns PatternFlowLacpActorTypeChoiceEnum provided by user to PatternFlowLacpActorType
	setChoice(value PatternFlowLacpActorTypeChoiceEnum) PatternFlowLacpActorType
	// HasChoice checks if Choice has been set in PatternFlowLacpActorType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorType
	SetValue(value uint32) PatternFlowLacpActorType
	// HasValue checks if Value has been set in PatternFlowLacpActorType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorType
	SetValues(value []uint32) PatternFlowLacpActorType
	// Increment returns PatternFlowLacpActorTypeCounter, set in PatternFlowLacpActorType.
	// PatternFlowLacpActorTypeCounter is integer counter pattern
	Increment() PatternFlowLacpActorTypeCounter
	// SetIncrement assigns PatternFlowLacpActorTypeCounter provided by user to PatternFlowLacpActorType.
	// PatternFlowLacpActorTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorTypeCounter) PatternFlowLacpActorType
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorType
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorTypeCounter, set in PatternFlowLacpActorType.
	// PatternFlowLacpActorTypeCounter is integer counter pattern
	Decrement() PatternFlowLacpActorTypeCounter
	// SetDecrement assigns PatternFlowLacpActorTypeCounter provided by user to PatternFlowLacpActorType.
	// PatternFlowLacpActorTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorTypeCounter) PatternFlowLacpActorType
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorType
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorTypeChoiceEnum string

// Enum of Choice on PatternFlowLacpActorType
var PatternFlowLacpActorTypeChoice = struct {
	VALUE     PatternFlowLacpActorTypeChoiceEnum
	VALUES    PatternFlowLacpActorTypeChoiceEnum
	INCREMENT PatternFlowLacpActorTypeChoiceEnum
	DECREMENT PatternFlowLacpActorTypeChoiceEnum
}{
	VALUE:     PatternFlowLacpActorTypeChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorTypeChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorTypeChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorType) Choice() PatternFlowLacpActorTypeChoiceEnum {
	return PatternFlowLacpActorTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorType) setChoice(value PatternFlowLacpActorTypeChoiceEnum) PatternFlowLacpActorType {
	intValue, ok := otg.PatternFlowLacpActorType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorTypeChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorTypeChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorTypeCounter().msg()
	}

	if value == PatternFlowLacpActorTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorType object
func (obj *patternFlowLacpActorType) SetValue(value uint32) PatternFlowLacpActorType {
	obj.setChoice(PatternFlowLacpActorTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorType object
func (obj *patternFlowLacpActorType) SetValues(value []uint32) PatternFlowLacpActorType {
	obj.setChoice(PatternFlowLacpActorTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorTypeCounter
func (obj *patternFlowLacpActorType) Increment() PatternFlowLacpActorTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorTypeCounter
func (obj *patternFlowLacpActorType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorTypeCounter value in the PatternFlowLacpActorType object
func (obj *patternFlowLacpActorType) SetIncrement(value PatternFlowLacpActorTypeCounter) PatternFlowLacpActorType {
	obj.setChoice(PatternFlowLacpActorTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorTypeCounter
func (obj *patternFlowLacpActorType) Decrement() PatternFlowLacpActorTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorTypeCounter
func (obj *patternFlowLacpActorType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorTypeCounter value in the PatternFlowLacpActorType object
func (obj *patternFlowLacpActorType) SetDecrement(value PatternFlowLacpActorTypeCounter) PatternFlowLacpActorType {
	obj.setChoice(PatternFlowLacpActorTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorType.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpActorType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorType")
			}
		} else {
			intVal := otg.PatternFlowLacpActorType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
