package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduCollectorType *****
type patternFlowLacpduCollectorType struct {
	validation
	obj             *otg.PatternFlowLacpduCollectorType
	marshaller      marshalPatternFlowLacpduCollectorType
	unMarshaller    unMarshalPatternFlowLacpduCollectorType
	incrementHolder PatternFlowLacpduCollectorTypeCounter
	decrementHolder PatternFlowLacpduCollectorTypeCounter
}

func NewPatternFlowLacpduCollectorType() PatternFlowLacpduCollectorType {
	obj := patternFlowLacpduCollectorType{obj: &otg.PatternFlowLacpduCollectorType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduCollectorType) msg() *otg.PatternFlowLacpduCollectorType {
	return obj.obj
}

func (obj *patternFlowLacpduCollectorType) setMsg(msg *otg.PatternFlowLacpduCollectorType) PatternFlowLacpduCollectorType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduCollectorType struct {
	obj *patternFlowLacpduCollectorType
}

type marshalPatternFlowLacpduCollectorType interface {
	// ToProto marshals PatternFlowLacpduCollectorType to protobuf object *otg.PatternFlowLacpduCollectorType
	ToProto() (*otg.PatternFlowLacpduCollectorType, error)
	// ToPbText marshals PatternFlowLacpduCollectorType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduCollectorType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduCollectorType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduCollectorType struct {
	obj *patternFlowLacpduCollectorType
}

type unMarshalPatternFlowLacpduCollectorType interface {
	// FromProto unmarshals PatternFlowLacpduCollectorType from protobuf object *otg.PatternFlowLacpduCollectorType
	FromProto(msg *otg.PatternFlowLacpduCollectorType) (PatternFlowLacpduCollectorType, error)
	// FromPbText unmarshals PatternFlowLacpduCollectorType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduCollectorType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduCollectorType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduCollectorType) Marshal() marshalPatternFlowLacpduCollectorType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduCollectorType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduCollectorType) Unmarshal() unMarshalPatternFlowLacpduCollectorType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduCollectorType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduCollectorType) ToProto() (*otg.PatternFlowLacpduCollectorType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduCollectorType) FromProto(msg *otg.PatternFlowLacpduCollectorType) (PatternFlowLacpduCollectorType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduCollectorType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorType) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduCollectorType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorType) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduCollectorType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorType) FromJson(value string) error {
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

func (obj *patternFlowLacpduCollectorType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduCollectorType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduCollectorType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduCollectorType) Clone() (PatternFlowLacpduCollectorType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduCollectorType()
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

func (obj *patternFlowLacpduCollectorType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduCollectorType is tLV Type for Collector Information. The value 0x03 identifies this TLV.
type PatternFlowLacpduCollectorType interface {
	Validation
	// msg marshals PatternFlowLacpduCollectorType to protobuf object *otg.PatternFlowLacpduCollectorType
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduCollectorType
	// setMsg unmarshals PatternFlowLacpduCollectorType from protobuf object *otg.PatternFlowLacpduCollectorType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduCollectorType) PatternFlowLacpduCollectorType
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduCollectorType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduCollectorType
	// validate validates PatternFlowLacpduCollectorType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduCollectorType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduCollectorTypeChoiceEnum, set in PatternFlowLacpduCollectorType
	Choice() PatternFlowLacpduCollectorTypeChoiceEnum
	// setChoice assigns PatternFlowLacpduCollectorTypeChoiceEnum provided by user to PatternFlowLacpduCollectorType
	setChoice(value PatternFlowLacpduCollectorTypeChoiceEnum) PatternFlowLacpduCollectorType
	// HasChoice checks if Choice has been set in PatternFlowLacpduCollectorType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduCollectorType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduCollectorType
	SetValue(value uint32) PatternFlowLacpduCollectorType
	// HasValue checks if Value has been set in PatternFlowLacpduCollectorType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduCollectorType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduCollectorType
	SetValues(value []uint32) PatternFlowLacpduCollectorType
	// Increment returns PatternFlowLacpduCollectorTypeCounter, set in PatternFlowLacpduCollectorType.
	// PatternFlowLacpduCollectorTypeCounter is integer counter pattern
	Increment() PatternFlowLacpduCollectorTypeCounter
	// SetIncrement assigns PatternFlowLacpduCollectorTypeCounter provided by user to PatternFlowLacpduCollectorType.
	// PatternFlowLacpduCollectorTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduCollectorTypeCounter) PatternFlowLacpduCollectorType
	// HasIncrement checks if Increment has been set in PatternFlowLacpduCollectorType
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduCollectorTypeCounter, set in PatternFlowLacpduCollectorType.
	// PatternFlowLacpduCollectorTypeCounter is integer counter pattern
	Decrement() PatternFlowLacpduCollectorTypeCounter
	// SetDecrement assigns PatternFlowLacpduCollectorTypeCounter provided by user to PatternFlowLacpduCollectorType.
	// PatternFlowLacpduCollectorTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduCollectorTypeCounter) PatternFlowLacpduCollectorType
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduCollectorType
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduCollectorTypeChoiceEnum string

// Enum of Choice on PatternFlowLacpduCollectorType
var PatternFlowLacpduCollectorTypeChoice = struct {
	VALUE     PatternFlowLacpduCollectorTypeChoiceEnum
	VALUES    PatternFlowLacpduCollectorTypeChoiceEnum
	INCREMENT PatternFlowLacpduCollectorTypeChoiceEnum
	DECREMENT PatternFlowLacpduCollectorTypeChoiceEnum
}{
	VALUE:     PatternFlowLacpduCollectorTypeChoiceEnum("value"),
	VALUES:    PatternFlowLacpduCollectorTypeChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduCollectorTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduCollectorTypeChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduCollectorType) Choice() PatternFlowLacpduCollectorTypeChoiceEnum {
	return PatternFlowLacpduCollectorTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduCollectorType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduCollectorType) setChoice(value PatternFlowLacpduCollectorTypeChoiceEnum) PatternFlowLacpduCollectorType {
	intValue, ok := otg.PatternFlowLacpduCollectorType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduCollectorTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduCollectorType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduCollectorTypeChoice.VALUE {
		defaultValue := uint32(3)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduCollectorTypeChoice.VALUES {
		defaultValue := []uint32{3}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduCollectorTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduCollectorTypeCounter().msg()
	}

	if value == PatternFlowLacpduCollectorTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduCollectorTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduCollectorType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduCollectorTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduCollectorType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduCollectorType object
func (obj *patternFlowLacpduCollectorType) SetValue(value uint32) PatternFlowLacpduCollectorType {
	obj.setChoice(PatternFlowLacpduCollectorTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduCollectorType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{3})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduCollectorType object
func (obj *patternFlowLacpduCollectorType) SetValues(value []uint32) PatternFlowLacpduCollectorType {
	obj.setChoice(PatternFlowLacpduCollectorTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduCollectorTypeCounter
func (obj *patternFlowLacpduCollectorType) Increment() PatternFlowLacpduCollectorTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduCollectorTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduCollectorTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduCollectorTypeCounter
func (obj *patternFlowLacpduCollectorType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduCollectorTypeCounter value in the PatternFlowLacpduCollectorType object
func (obj *patternFlowLacpduCollectorType) SetIncrement(value PatternFlowLacpduCollectorTypeCounter) PatternFlowLacpduCollectorType {
	obj.setChoice(PatternFlowLacpduCollectorTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduCollectorTypeCounter
func (obj *patternFlowLacpduCollectorType) Decrement() PatternFlowLacpduCollectorTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduCollectorTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduCollectorTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduCollectorTypeCounter
func (obj *patternFlowLacpduCollectorType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduCollectorTypeCounter value in the PatternFlowLacpduCollectorType object
func (obj *patternFlowLacpduCollectorType) SetDecrement(value PatternFlowLacpduCollectorTypeCounter) PatternFlowLacpduCollectorType {
	obj.setChoice(PatternFlowLacpduCollectorTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduCollectorType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduCollectorType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduCollectorType.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpduCollectorType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduCollectorTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduCollectorTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduCollectorTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduCollectorTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduCollectorTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduCollectorTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduCollectorType")
			}
		} else {
			intVal := otg.PatternFlowLacpduCollectorType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduCollectorType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
