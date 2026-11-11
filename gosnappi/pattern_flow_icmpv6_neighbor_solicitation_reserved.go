package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NeighborSolicitationReserved *****
type patternFlowIcmpv6NeighborSolicitationReserved struct {
	validation
	obj             *otg.PatternFlowIcmpv6NeighborSolicitationReserved
	marshaller      marshalPatternFlowIcmpv6NeighborSolicitationReserved
	unMarshaller    unMarshalPatternFlowIcmpv6NeighborSolicitationReserved
	incrementHolder PatternFlowIcmpv6NeighborSolicitationReservedCounter
	decrementHolder PatternFlowIcmpv6NeighborSolicitationReservedCounter
}

func NewPatternFlowIcmpv6NeighborSolicitationReserved() PatternFlowIcmpv6NeighborSolicitationReserved {
	obj := patternFlowIcmpv6NeighborSolicitationReserved{obj: &otg.PatternFlowIcmpv6NeighborSolicitationReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationReserved) msg() *otg.PatternFlowIcmpv6NeighborSolicitationReserved {
	return obj.obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationReserved) setMsg(msg *otg.PatternFlowIcmpv6NeighborSolicitationReserved) PatternFlowIcmpv6NeighborSolicitationReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NeighborSolicitationReserved struct {
	obj *patternFlowIcmpv6NeighborSolicitationReserved
}

type marshalPatternFlowIcmpv6NeighborSolicitationReserved interface {
	// ToProto marshals PatternFlowIcmpv6NeighborSolicitationReserved to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationReserved
	ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationReserved, error)
	// ToPbText marshals PatternFlowIcmpv6NeighborSolicitationReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NeighborSolicitationReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NeighborSolicitationReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NeighborSolicitationReserved struct {
	obj *patternFlowIcmpv6NeighborSolicitationReserved
}

type unMarshalPatternFlowIcmpv6NeighborSolicitationReserved interface {
	// FromProto unmarshals PatternFlowIcmpv6NeighborSolicitationReserved from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationReserved
	FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationReserved) (PatternFlowIcmpv6NeighborSolicitationReserved, error)
	// FromPbText unmarshals PatternFlowIcmpv6NeighborSolicitationReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NeighborSolicitationReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NeighborSolicitationReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NeighborSolicitationReserved) Marshal() marshalPatternFlowIcmpv6NeighborSolicitationReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NeighborSolicitationReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NeighborSolicitationReserved) Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NeighborSolicitationReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationReserved) ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationReserved) FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationReserved) (PatternFlowIcmpv6NeighborSolicitationReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationReserved) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NeighborSolicitationReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NeighborSolicitationReserved) Clone() (PatternFlowIcmpv6NeighborSolicitationReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NeighborSolicitationReserved()
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

func (obj *patternFlowIcmpv6NeighborSolicitationReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NeighborSolicitationReserved is reserved field. Must be initialized to zero by the sender and must be ignored by the receiver.
type PatternFlowIcmpv6NeighborSolicitationReserved interface {
	Validation
	// msg marshals PatternFlowIcmpv6NeighborSolicitationReserved to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NeighborSolicitationReserved
	// setMsg unmarshals PatternFlowIcmpv6NeighborSolicitationReserved from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NeighborSolicitationReserved) PatternFlowIcmpv6NeighborSolicitationReserved
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NeighborSolicitationReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationReserved
	// validate validates PatternFlowIcmpv6NeighborSolicitationReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NeighborSolicitationReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum, set in PatternFlowIcmpv6NeighborSolicitationReserved
	Choice() PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum provided by user to PatternFlowIcmpv6NeighborSolicitationReserved
	setChoice(value PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum) PatternFlowIcmpv6NeighborSolicitationReserved
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NeighborSolicitationReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NeighborSolicitationReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationReserved
	SetValue(value uint32) PatternFlowIcmpv6NeighborSolicitationReserved
	// HasValue checks if Value has been set in PatternFlowIcmpv6NeighborSolicitationReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NeighborSolicitationReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationReserved
	SetValues(value []uint32) PatternFlowIcmpv6NeighborSolicitationReserved
	// Increment returns PatternFlowIcmpv6NeighborSolicitationReservedCounter, set in PatternFlowIcmpv6NeighborSolicitationReserved.
	// PatternFlowIcmpv6NeighborSolicitationReservedCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NeighborSolicitationReservedCounter
	// SetIncrement assigns PatternFlowIcmpv6NeighborSolicitationReservedCounter provided by user to PatternFlowIcmpv6NeighborSolicitationReserved.
	// PatternFlowIcmpv6NeighborSolicitationReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NeighborSolicitationReservedCounter) PatternFlowIcmpv6NeighborSolicitationReserved
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NeighborSolicitationReserved
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NeighborSolicitationReservedCounter, set in PatternFlowIcmpv6NeighborSolicitationReserved.
	// PatternFlowIcmpv6NeighborSolicitationReservedCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NeighborSolicitationReservedCounter
	// SetDecrement assigns PatternFlowIcmpv6NeighborSolicitationReservedCounter provided by user to PatternFlowIcmpv6NeighborSolicitationReserved.
	// PatternFlowIcmpv6NeighborSolicitationReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NeighborSolicitationReservedCounter) PatternFlowIcmpv6NeighborSolicitationReserved
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NeighborSolicitationReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NeighborSolicitationReserved
var PatternFlowIcmpv6NeighborSolicitationReservedChoice = struct {
	VALUE     PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum
	VALUES    PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum
	INCREMENT PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum
	DECREMENT PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NeighborSolicitationReserved) Choice() PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum {
	return PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NeighborSolicitationReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NeighborSolicitationReserved) setChoice(value PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum) PatternFlowIcmpv6NeighborSolicitationReserved {
	intValue, ok := otg.PatternFlowIcmpv6NeighborSolicitationReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NeighborSolicitationReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NeighborSolicitationReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NeighborSolicitationReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NeighborSolicitationReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NeighborSolicitationReservedCounter().msg()
	}

	if value == PatternFlowIcmpv6NeighborSolicitationReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NeighborSolicitationReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NeighborSolicitationReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationReserved object
func (obj *patternFlowIcmpv6NeighborSolicitationReserved) SetValue(value uint32) PatternFlowIcmpv6NeighborSolicitationReserved {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NeighborSolicitationReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NeighborSolicitationReserved object
func (obj *patternFlowIcmpv6NeighborSolicitationReserved) SetValues(value []uint32) PatternFlowIcmpv6NeighborSolicitationReserved {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NeighborSolicitationReservedCounter
func (obj *patternFlowIcmpv6NeighborSolicitationReserved) Increment() PatternFlowIcmpv6NeighborSolicitationReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NeighborSolicitationReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NeighborSolicitationReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NeighborSolicitationReservedCounter
func (obj *patternFlowIcmpv6NeighborSolicitationReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NeighborSolicitationReservedCounter value in the PatternFlowIcmpv6NeighborSolicitationReserved object
func (obj *patternFlowIcmpv6NeighborSolicitationReserved) SetIncrement(value PatternFlowIcmpv6NeighborSolicitationReservedCounter) PatternFlowIcmpv6NeighborSolicitationReserved {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NeighborSolicitationReservedCounter
func (obj *patternFlowIcmpv6NeighborSolicitationReserved) Decrement() PatternFlowIcmpv6NeighborSolicitationReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NeighborSolicitationReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NeighborSolicitationReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NeighborSolicitationReservedCounter
func (obj *patternFlowIcmpv6NeighborSolicitationReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NeighborSolicitationReservedCounter value in the PatternFlowIcmpv6NeighborSolicitationReserved object
func (obj *patternFlowIcmpv6NeighborSolicitationReserved) SetDecrement(value PatternFlowIcmpv6NeighborSolicitationReservedCounter) PatternFlowIcmpv6NeighborSolicitationReserved {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIcmpv6NeighborSolicitationReserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NeighborSolicitationReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NeighborSolicitationReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NeighborSolicitationReserved")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NeighborSolicitationReserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NeighborSolicitationReserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
