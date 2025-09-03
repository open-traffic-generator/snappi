package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorReserved *****
type patternFlowLacpActorReserved struct {
	validation
	obj             *otg.PatternFlowLacpActorReserved
	marshaller      marshalPatternFlowLacpActorReserved
	unMarshaller    unMarshalPatternFlowLacpActorReserved
	incrementHolder PatternFlowLacpActorReservedCounter
	decrementHolder PatternFlowLacpActorReservedCounter
}

func NewPatternFlowLacpActorReserved() PatternFlowLacpActorReserved {
	obj := patternFlowLacpActorReserved{obj: &otg.PatternFlowLacpActorReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorReserved) msg() *otg.PatternFlowLacpActorReserved {
	return obj.obj
}

func (obj *patternFlowLacpActorReserved) setMsg(msg *otg.PatternFlowLacpActorReserved) PatternFlowLacpActorReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorReserved struct {
	obj *patternFlowLacpActorReserved
}

type marshalPatternFlowLacpActorReserved interface {
	// ToProto marshals PatternFlowLacpActorReserved to protobuf object *otg.PatternFlowLacpActorReserved
	ToProto() (*otg.PatternFlowLacpActorReserved, error)
	// ToPbText marshals PatternFlowLacpActorReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorReserved struct {
	obj *patternFlowLacpActorReserved
}

type unMarshalPatternFlowLacpActorReserved interface {
	// FromProto unmarshals PatternFlowLacpActorReserved from protobuf object *otg.PatternFlowLacpActorReserved
	FromProto(msg *otg.PatternFlowLacpActorReserved) (PatternFlowLacpActorReserved, error)
	// FromPbText unmarshals PatternFlowLacpActorReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorReserved) Marshal() marshalPatternFlowLacpActorReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorReserved) Unmarshal() unMarshalPatternFlowLacpActorReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorReserved) ToProto() (*otg.PatternFlowLacpActorReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorReserved) FromProto(msg *otg.PatternFlowLacpActorReserved) (PatternFlowLacpActorReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorReserved) FromJson(value string) error {
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

func (obj *patternFlowLacpActorReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorReserved) Clone() (PatternFlowLacpActorReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorReserved()
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

func (obj *patternFlowLacpActorReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorReserved is reserved field for future use in the Actor TLV. Should be set to all zeros.
type PatternFlowLacpActorReserved interface {
	Validation
	// msg marshals PatternFlowLacpActorReserved to protobuf object *otg.PatternFlowLacpActorReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorReserved
	// setMsg unmarshals PatternFlowLacpActorReserved from protobuf object *otg.PatternFlowLacpActorReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorReserved) PatternFlowLacpActorReserved
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorReserved
	// validate validates PatternFlowLacpActorReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorReservedChoiceEnum, set in PatternFlowLacpActorReserved
	Choice() PatternFlowLacpActorReservedChoiceEnum
	// setChoice assigns PatternFlowLacpActorReservedChoiceEnum provided by user to PatternFlowLacpActorReserved
	setChoice(value PatternFlowLacpActorReservedChoiceEnum) PatternFlowLacpActorReserved
	// HasChoice checks if Choice has been set in PatternFlowLacpActorReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorReserved
	SetValue(value uint32) PatternFlowLacpActorReserved
	// HasValue checks if Value has been set in PatternFlowLacpActorReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorReserved
	SetValues(value []uint32) PatternFlowLacpActorReserved
	// Increment returns PatternFlowLacpActorReservedCounter, set in PatternFlowLacpActorReserved.
	// PatternFlowLacpActorReservedCounter is integer counter pattern
	Increment() PatternFlowLacpActorReservedCounter
	// SetIncrement assigns PatternFlowLacpActorReservedCounter provided by user to PatternFlowLacpActorReserved.
	// PatternFlowLacpActorReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorReservedCounter) PatternFlowLacpActorReserved
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorReserved
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorReservedCounter, set in PatternFlowLacpActorReserved.
	// PatternFlowLacpActorReservedCounter is integer counter pattern
	Decrement() PatternFlowLacpActorReservedCounter
	// SetDecrement assigns PatternFlowLacpActorReservedCounter provided by user to PatternFlowLacpActorReserved.
	// PatternFlowLacpActorReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorReservedCounter) PatternFlowLacpActorReserved
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorReservedChoiceEnum string

// Enum of Choice on PatternFlowLacpActorReserved
var PatternFlowLacpActorReservedChoice = struct {
	VALUE     PatternFlowLacpActorReservedChoiceEnum
	VALUES    PatternFlowLacpActorReservedChoiceEnum
	INCREMENT PatternFlowLacpActorReservedChoiceEnum
	DECREMENT PatternFlowLacpActorReservedChoiceEnum
}{
	VALUE:     PatternFlowLacpActorReservedChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorReservedChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorReservedChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorReserved) Choice() PatternFlowLacpActorReservedChoiceEnum {
	return PatternFlowLacpActorReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorReserved) setChoice(value PatternFlowLacpActorReservedChoiceEnum) PatternFlowLacpActorReserved {
	intValue, ok := otg.PatternFlowLacpActorReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorReservedCounter().msg()
	}

	if value == PatternFlowLacpActorReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorReserved object
func (obj *patternFlowLacpActorReserved) SetValue(value uint32) PatternFlowLacpActorReserved {
	obj.setChoice(PatternFlowLacpActorReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorReserved object
func (obj *patternFlowLacpActorReserved) SetValues(value []uint32) PatternFlowLacpActorReserved {
	obj.setChoice(PatternFlowLacpActorReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorReservedCounter
func (obj *patternFlowLacpActorReserved) Increment() PatternFlowLacpActorReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorReservedCounter
func (obj *patternFlowLacpActorReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorReservedCounter value in the PatternFlowLacpActorReserved object
func (obj *patternFlowLacpActorReserved) SetIncrement(value PatternFlowLacpActorReservedCounter) PatternFlowLacpActorReserved {
	obj.setChoice(PatternFlowLacpActorReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorReservedCounter
func (obj *patternFlowLacpActorReserved) Decrement() PatternFlowLacpActorReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorReservedCounter
func (obj *patternFlowLacpActorReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorReservedCounter value in the PatternFlowLacpActorReserved object
func (obj *patternFlowLacpActorReserved) SetDecrement(value PatternFlowLacpActorReservedCounter) PatternFlowLacpActorReserved {
	obj.setChoice(PatternFlowLacpActorReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorReserved.Value <= 16777215 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 16777215 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorReserved.Values <= 16777215 but Got %d", item))
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

func (obj *patternFlowLacpActorReserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorReserved")
			}
		} else {
			intVal := otg.PatternFlowLacpActorReserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorReserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
