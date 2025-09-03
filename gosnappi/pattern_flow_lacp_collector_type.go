package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpCollectorType *****
type patternFlowLacpCollectorType struct {
	validation
	obj             *otg.PatternFlowLacpCollectorType
	marshaller      marshalPatternFlowLacpCollectorType
	unMarshaller    unMarshalPatternFlowLacpCollectorType
	incrementHolder PatternFlowLacpCollectorTypeCounter
	decrementHolder PatternFlowLacpCollectorTypeCounter
}

func NewPatternFlowLacpCollectorType() PatternFlowLacpCollectorType {
	obj := patternFlowLacpCollectorType{obj: &otg.PatternFlowLacpCollectorType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpCollectorType) msg() *otg.PatternFlowLacpCollectorType {
	return obj.obj
}

func (obj *patternFlowLacpCollectorType) setMsg(msg *otg.PatternFlowLacpCollectorType) PatternFlowLacpCollectorType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpCollectorType struct {
	obj *patternFlowLacpCollectorType
}

type marshalPatternFlowLacpCollectorType interface {
	// ToProto marshals PatternFlowLacpCollectorType to protobuf object *otg.PatternFlowLacpCollectorType
	ToProto() (*otg.PatternFlowLacpCollectorType, error)
	// ToPbText marshals PatternFlowLacpCollectorType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpCollectorType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpCollectorType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpCollectorType struct {
	obj *patternFlowLacpCollectorType
}

type unMarshalPatternFlowLacpCollectorType interface {
	// FromProto unmarshals PatternFlowLacpCollectorType from protobuf object *otg.PatternFlowLacpCollectorType
	FromProto(msg *otg.PatternFlowLacpCollectorType) (PatternFlowLacpCollectorType, error)
	// FromPbText unmarshals PatternFlowLacpCollectorType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpCollectorType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpCollectorType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpCollectorType) Marshal() marshalPatternFlowLacpCollectorType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpCollectorType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpCollectorType) Unmarshal() unMarshalPatternFlowLacpCollectorType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpCollectorType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpCollectorType) ToProto() (*otg.PatternFlowLacpCollectorType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpCollectorType) FromProto(msg *otg.PatternFlowLacpCollectorType) (PatternFlowLacpCollectorType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpCollectorType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorType) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpCollectorType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorType) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpCollectorType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorType) FromJson(value string) error {
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

func (obj *patternFlowLacpCollectorType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpCollectorType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpCollectorType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpCollectorType) Clone() (PatternFlowLacpCollectorType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpCollectorType()
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

func (obj *patternFlowLacpCollectorType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpCollectorType is tLV Type for Collector Information. The value 0x03 identifies this TLV.
type PatternFlowLacpCollectorType interface {
	Validation
	// msg marshals PatternFlowLacpCollectorType to protobuf object *otg.PatternFlowLacpCollectorType
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpCollectorType
	// setMsg unmarshals PatternFlowLacpCollectorType from protobuf object *otg.PatternFlowLacpCollectorType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpCollectorType) PatternFlowLacpCollectorType
	// provides marshal interface
	Marshal() marshalPatternFlowLacpCollectorType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpCollectorType
	// validate validates PatternFlowLacpCollectorType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpCollectorType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpCollectorTypeChoiceEnum, set in PatternFlowLacpCollectorType
	Choice() PatternFlowLacpCollectorTypeChoiceEnum
	// setChoice assigns PatternFlowLacpCollectorTypeChoiceEnum provided by user to PatternFlowLacpCollectorType
	setChoice(value PatternFlowLacpCollectorTypeChoiceEnum) PatternFlowLacpCollectorType
	// HasChoice checks if Choice has been set in PatternFlowLacpCollectorType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpCollectorType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpCollectorType
	SetValue(value uint32) PatternFlowLacpCollectorType
	// HasValue checks if Value has been set in PatternFlowLacpCollectorType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpCollectorType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpCollectorType
	SetValues(value []uint32) PatternFlowLacpCollectorType
	// Increment returns PatternFlowLacpCollectorTypeCounter, set in PatternFlowLacpCollectorType.
	// PatternFlowLacpCollectorTypeCounter is integer counter pattern
	Increment() PatternFlowLacpCollectorTypeCounter
	// SetIncrement assigns PatternFlowLacpCollectorTypeCounter provided by user to PatternFlowLacpCollectorType.
	// PatternFlowLacpCollectorTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpCollectorTypeCounter) PatternFlowLacpCollectorType
	// HasIncrement checks if Increment has been set in PatternFlowLacpCollectorType
	HasIncrement() bool
	// Decrement returns PatternFlowLacpCollectorTypeCounter, set in PatternFlowLacpCollectorType.
	// PatternFlowLacpCollectorTypeCounter is integer counter pattern
	Decrement() PatternFlowLacpCollectorTypeCounter
	// SetDecrement assigns PatternFlowLacpCollectorTypeCounter provided by user to PatternFlowLacpCollectorType.
	// PatternFlowLacpCollectorTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpCollectorTypeCounter) PatternFlowLacpCollectorType
	// HasDecrement checks if Decrement has been set in PatternFlowLacpCollectorType
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpCollectorTypeChoiceEnum string

// Enum of Choice on PatternFlowLacpCollectorType
var PatternFlowLacpCollectorTypeChoice = struct {
	VALUE     PatternFlowLacpCollectorTypeChoiceEnum
	VALUES    PatternFlowLacpCollectorTypeChoiceEnum
	INCREMENT PatternFlowLacpCollectorTypeChoiceEnum
	DECREMENT PatternFlowLacpCollectorTypeChoiceEnum
}{
	VALUE:     PatternFlowLacpCollectorTypeChoiceEnum("value"),
	VALUES:    PatternFlowLacpCollectorTypeChoiceEnum("values"),
	INCREMENT: PatternFlowLacpCollectorTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpCollectorTypeChoiceEnum("decrement"),
}

func (obj *patternFlowLacpCollectorType) Choice() PatternFlowLacpCollectorTypeChoiceEnum {
	return PatternFlowLacpCollectorTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpCollectorType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpCollectorType) setChoice(value PatternFlowLacpCollectorTypeChoiceEnum) PatternFlowLacpCollectorType {
	intValue, ok := otg.PatternFlowLacpCollectorType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpCollectorTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpCollectorType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpCollectorTypeChoice.VALUE {
		defaultValue := uint32(3)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpCollectorTypeChoice.VALUES {
		defaultValue := []uint32{3}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpCollectorTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpCollectorTypeCounter().msg()
	}

	if value == PatternFlowLacpCollectorTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpCollectorTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpCollectorType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpCollectorTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpCollectorType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpCollectorType object
func (obj *patternFlowLacpCollectorType) SetValue(value uint32) PatternFlowLacpCollectorType {
	obj.setChoice(PatternFlowLacpCollectorTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpCollectorType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{3})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpCollectorType object
func (obj *patternFlowLacpCollectorType) SetValues(value []uint32) PatternFlowLacpCollectorType {
	obj.setChoice(PatternFlowLacpCollectorTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpCollectorTypeCounter
func (obj *patternFlowLacpCollectorType) Increment() PatternFlowLacpCollectorTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpCollectorTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpCollectorTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpCollectorTypeCounter
func (obj *patternFlowLacpCollectorType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpCollectorTypeCounter value in the PatternFlowLacpCollectorType object
func (obj *patternFlowLacpCollectorType) SetIncrement(value PatternFlowLacpCollectorTypeCounter) PatternFlowLacpCollectorType {
	obj.setChoice(PatternFlowLacpCollectorTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpCollectorTypeCounter
func (obj *patternFlowLacpCollectorType) Decrement() PatternFlowLacpCollectorTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpCollectorTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpCollectorTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpCollectorTypeCounter
func (obj *patternFlowLacpCollectorType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpCollectorTypeCounter value in the PatternFlowLacpCollectorType object
func (obj *patternFlowLacpCollectorType) SetDecrement(value PatternFlowLacpCollectorTypeCounter) PatternFlowLacpCollectorType {
	obj.setChoice(PatternFlowLacpCollectorTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpCollectorType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpCollectorType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpCollectorType.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpCollectorType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpCollectorTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpCollectorTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpCollectorTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpCollectorTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpCollectorTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpCollectorTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpCollectorType")
			}
		} else {
			intVal := otg.PatternFlowLacpCollectorType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpCollectorType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
