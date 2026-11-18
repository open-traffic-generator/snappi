package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpMtuOptionReserved *****
type patternFlowIcmpv6NdpMtuOptionReserved struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpMtuOptionReserved
	marshaller      marshalPatternFlowIcmpv6NdpMtuOptionReserved
	unMarshaller    unMarshalPatternFlowIcmpv6NdpMtuOptionReserved
	incrementHolder PatternFlowIcmpv6NdpMtuOptionReservedCounter
	decrementHolder PatternFlowIcmpv6NdpMtuOptionReservedCounter
}

func NewPatternFlowIcmpv6NdpMtuOptionReserved() PatternFlowIcmpv6NdpMtuOptionReserved {
	obj := patternFlowIcmpv6NdpMtuOptionReserved{obj: &otg.PatternFlowIcmpv6NdpMtuOptionReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionReserved) msg() *otg.PatternFlowIcmpv6NdpMtuOptionReserved {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionReserved) setMsg(msg *otg.PatternFlowIcmpv6NdpMtuOptionReserved) PatternFlowIcmpv6NdpMtuOptionReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpMtuOptionReserved struct {
	obj *patternFlowIcmpv6NdpMtuOptionReserved
}

type marshalPatternFlowIcmpv6NdpMtuOptionReserved interface {
	// ToProto marshals PatternFlowIcmpv6NdpMtuOptionReserved to protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionReserved
	ToProto() (*otg.PatternFlowIcmpv6NdpMtuOptionReserved, error)
	// ToPbText marshals PatternFlowIcmpv6NdpMtuOptionReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpMtuOptionReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpMtuOptionReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpMtuOptionReserved struct {
	obj *patternFlowIcmpv6NdpMtuOptionReserved
}

type unMarshalPatternFlowIcmpv6NdpMtuOptionReserved interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpMtuOptionReserved from protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionReserved
	FromProto(msg *otg.PatternFlowIcmpv6NdpMtuOptionReserved) (PatternFlowIcmpv6NdpMtuOptionReserved, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpMtuOptionReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpMtuOptionReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpMtuOptionReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpMtuOptionReserved) Marshal() marshalPatternFlowIcmpv6NdpMtuOptionReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpMtuOptionReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpMtuOptionReserved) Unmarshal() unMarshalPatternFlowIcmpv6NdpMtuOptionReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpMtuOptionReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpMtuOptionReserved) ToProto() (*otg.PatternFlowIcmpv6NdpMtuOptionReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionReserved) FromProto(msg *otg.PatternFlowIcmpv6NdpMtuOptionReserved) (PatternFlowIcmpv6NdpMtuOptionReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpMtuOptionReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpMtuOptionReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpMtuOptionReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionReserved) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpMtuOptionReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpMtuOptionReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpMtuOptionReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpMtuOptionReserved) Clone() (PatternFlowIcmpv6NdpMtuOptionReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpMtuOptionReserved()
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

func (obj *patternFlowIcmpv6NdpMtuOptionReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpMtuOptionReserved is reserved. Must be initialized to zero by the sender.
type PatternFlowIcmpv6NdpMtuOptionReserved interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpMtuOptionReserved to protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpMtuOptionReserved
	// setMsg unmarshals PatternFlowIcmpv6NdpMtuOptionReserved from protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpMtuOptionReserved) PatternFlowIcmpv6NdpMtuOptionReserved
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpMtuOptionReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpMtuOptionReserved
	// validate validates PatternFlowIcmpv6NdpMtuOptionReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpMtuOptionReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum, set in PatternFlowIcmpv6NdpMtuOptionReserved
	Choice() PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum provided by user to PatternFlowIcmpv6NdpMtuOptionReserved
	setChoice(value PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum) PatternFlowIcmpv6NdpMtuOptionReserved
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpMtuOptionReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpMtuOptionReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionReserved
	SetValue(value uint32) PatternFlowIcmpv6NdpMtuOptionReserved
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpMtuOptionReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpMtuOptionReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionReserved
	SetValues(value []uint32) PatternFlowIcmpv6NdpMtuOptionReserved
	// Increment returns PatternFlowIcmpv6NdpMtuOptionReservedCounter, set in PatternFlowIcmpv6NdpMtuOptionReserved.
	// PatternFlowIcmpv6NdpMtuOptionReservedCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpMtuOptionReservedCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpMtuOptionReservedCounter provided by user to PatternFlowIcmpv6NdpMtuOptionReserved.
	// PatternFlowIcmpv6NdpMtuOptionReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpMtuOptionReservedCounter) PatternFlowIcmpv6NdpMtuOptionReserved
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpMtuOptionReserved
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpMtuOptionReservedCounter, set in PatternFlowIcmpv6NdpMtuOptionReserved.
	// PatternFlowIcmpv6NdpMtuOptionReservedCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpMtuOptionReservedCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpMtuOptionReservedCounter provided by user to PatternFlowIcmpv6NdpMtuOptionReserved.
	// PatternFlowIcmpv6NdpMtuOptionReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpMtuOptionReservedCounter) PatternFlowIcmpv6NdpMtuOptionReserved
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpMtuOptionReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpMtuOptionReserved
var PatternFlowIcmpv6NdpMtuOptionReservedChoice = struct {
	VALUE     PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum
	VALUES    PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpMtuOptionReserved) Choice() PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum {
	return PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpMtuOptionReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpMtuOptionReserved) setChoice(value PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum) PatternFlowIcmpv6NdpMtuOptionReserved {
	intValue, ok := otg.PatternFlowIcmpv6NdpMtuOptionReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpMtuOptionReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpMtuOptionReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpMtuOptionReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpMtuOptionReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpMtuOptionReservedCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpMtuOptionReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpMtuOptionReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpMtuOptionReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpMtuOptionReserved object
func (obj *patternFlowIcmpv6NdpMtuOptionReserved) SetValue(value uint32) PatternFlowIcmpv6NdpMtuOptionReserved {
	obj.setChoice(PatternFlowIcmpv6NdpMtuOptionReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpMtuOptionReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpMtuOptionReserved object
func (obj *patternFlowIcmpv6NdpMtuOptionReserved) SetValues(value []uint32) PatternFlowIcmpv6NdpMtuOptionReserved {
	obj.setChoice(PatternFlowIcmpv6NdpMtuOptionReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpMtuOptionReservedCounter
func (obj *patternFlowIcmpv6NdpMtuOptionReserved) Increment() PatternFlowIcmpv6NdpMtuOptionReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpMtuOptionReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpMtuOptionReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpMtuOptionReservedCounter
func (obj *patternFlowIcmpv6NdpMtuOptionReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpMtuOptionReservedCounter value in the PatternFlowIcmpv6NdpMtuOptionReserved object
func (obj *patternFlowIcmpv6NdpMtuOptionReserved) SetIncrement(value PatternFlowIcmpv6NdpMtuOptionReservedCounter) PatternFlowIcmpv6NdpMtuOptionReserved {
	obj.setChoice(PatternFlowIcmpv6NdpMtuOptionReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpMtuOptionReservedCounter
func (obj *patternFlowIcmpv6NdpMtuOptionReserved) Decrement() PatternFlowIcmpv6NdpMtuOptionReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpMtuOptionReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpMtuOptionReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpMtuOptionReservedCounter
func (obj *patternFlowIcmpv6NdpMtuOptionReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpMtuOptionReservedCounter value in the PatternFlowIcmpv6NdpMtuOptionReserved object
func (obj *patternFlowIcmpv6NdpMtuOptionReserved) SetDecrement(value PatternFlowIcmpv6NdpMtuOptionReservedCounter) PatternFlowIcmpv6NdpMtuOptionReserved {
	obj.setChoice(PatternFlowIcmpv6NdpMtuOptionReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpMtuOptionReserved.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6NdpMtuOptionReserved.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowIcmpv6NdpMtuOptionReserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpMtuOptionReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpMtuOptionReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpMtuOptionReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpMtuOptionReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpMtuOptionReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpMtuOptionReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpMtuOptionReserved")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpMtuOptionReserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpMtuOptionReserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
