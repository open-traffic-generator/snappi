package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRsvpReserved *****
type patternFlowRsvpReserved struct {
	validation
	obj             *otg.PatternFlowRsvpReserved
	marshaller      marshalPatternFlowRsvpReserved
	unMarshaller    unMarshalPatternFlowRsvpReserved
	incrementHolder PatternFlowRsvpReservedCounter
	decrementHolder PatternFlowRsvpReservedCounter
}

func NewPatternFlowRsvpReserved() PatternFlowRsvpReserved {
	obj := patternFlowRsvpReserved{obj: &otg.PatternFlowRsvpReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRsvpReserved) msg() *otg.PatternFlowRsvpReserved {
	return obj.obj
}

func (obj *patternFlowRsvpReserved) setMsg(msg *otg.PatternFlowRsvpReserved) PatternFlowRsvpReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRsvpReserved struct {
	obj *patternFlowRsvpReserved
}

type marshalPatternFlowRsvpReserved interface {
	// ToProto marshals PatternFlowRsvpReserved to protobuf object *otg.PatternFlowRsvpReserved
	ToProto() (*otg.PatternFlowRsvpReserved, error)
	// ToPbText marshals PatternFlowRsvpReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRsvpReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRsvpReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRsvpReserved struct {
	obj *patternFlowRsvpReserved
}

type unMarshalPatternFlowRsvpReserved interface {
	// FromProto unmarshals PatternFlowRsvpReserved from protobuf object *otg.PatternFlowRsvpReserved
	FromProto(msg *otg.PatternFlowRsvpReserved) (PatternFlowRsvpReserved, error)
	// FromPbText unmarshals PatternFlowRsvpReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRsvpReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRsvpReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRsvpReserved) Marshal() marshalPatternFlowRsvpReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRsvpReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRsvpReserved) Unmarshal() unMarshalPatternFlowRsvpReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRsvpReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRsvpReserved) ToProto() (*otg.PatternFlowRsvpReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRsvpReserved) FromProto(msg *otg.PatternFlowRsvpReserved) (PatternFlowRsvpReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRsvpReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRsvpReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowRsvpReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRsvpReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowRsvpReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRsvpReserved) FromJson(value string) error {
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

func (obj *patternFlowRsvpReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRsvpReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRsvpReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRsvpReserved) Clone() (PatternFlowRsvpReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRsvpReserved()
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

func (obj *patternFlowRsvpReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRsvpReserved is reserved
type PatternFlowRsvpReserved interface {
	Validation
	// msg marshals PatternFlowRsvpReserved to protobuf object *otg.PatternFlowRsvpReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowRsvpReserved
	// setMsg unmarshals PatternFlowRsvpReserved from protobuf object *otg.PatternFlowRsvpReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRsvpReserved) PatternFlowRsvpReserved
	// provides marshal interface
	Marshal() marshalPatternFlowRsvpReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRsvpReserved
	// validate validates PatternFlowRsvpReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRsvpReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRsvpReservedChoiceEnum, set in PatternFlowRsvpReserved
	Choice() PatternFlowRsvpReservedChoiceEnum
	// setChoice assigns PatternFlowRsvpReservedChoiceEnum provided by user to PatternFlowRsvpReserved
	setChoice(value PatternFlowRsvpReservedChoiceEnum) PatternFlowRsvpReserved
	// HasChoice checks if Choice has been set in PatternFlowRsvpReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRsvpReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRsvpReserved
	SetValue(value uint32) PatternFlowRsvpReserved
	// HasValue checks if Value has been set in PatternFlowRsvpReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRsvpReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRsvpReserved
	SetValues(value []uint32) PatternFlowRsvpReserved
	// Increment returns PatternFlowRsvpReservedCounter, set in PatternFlowRsvpReserved.
	// PatternFlowRsvpReservedCounter is integer counter pattern
	Increment() PatternFlowRsvpReservedCounter
	// SetIncrement assigns PatternFlowRsvpReservedCounter provided by user to PatternFlowRsvpReserved.
	// PatternFlowRsvpReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowRsvpReservedCounter) PatternFlowRsvpReserved
	// HasIncrement checks if Increment has been set in PatternFlowRsvpReserved
	HasIncrement() bool
	// Decrement returns PatternFlowRsvpReservedCounter, set in PatternFlowRsvpReserved.
	// PatternFlowRsvpReservedCounter is integer counter pattern
	Decrement() PatternFlowRsvpReservedCounter
	// SetDecrement assigns PatternFlowRsvpReservedCounter provided by user to PatternFlowRsvpReserved.
	// PatternFlowRsvpReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowRsvpReservedCounter) PatternFlowRsvpReserved
	// HasDecrement checks if Decrement has been set in PatternFlowRsvpReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowRsvpReservedChoiceEnum string

// Enum of Choice on PatternFlowRsvpReserved
var PatternFlowRsvpReservedChoice = struct {
	VALUE     PatternFlowRsvpReservedChoiceEnum
	VALUES    PatternFlowRsvpReservedChoiceEnum
	INCREMENT PatternFlowRsvpReservedChoiceEnum
	DECREMENT PatternFlowRsvpReservedChoiceEnum
}{
	VALUE:     PatternFlowRsvpReservedChoiceEnum("value"),
	VALUES:    PatternFlowRsvpReservedChoiceEnum("values"),
	INCREMENT: PatternFlowRsvpReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowRsvpReservedChoiceEnum("decrement"),
}

func (obj *patternFlowRsvpReserved) Choice() PatternFlowRsvpReservedChoiceEnum {
	return PatternFlowRsvpReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRsvpReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRsvpReserved) setChoice(value PatternFlowRsvpReservedChoiceEnum) PatternFlowRsvpReserved {
	intValue, ok := otg.PatternFlowRsvpReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRsvpReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRsvpReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRsvpReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRsvpReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRsvpReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRsvpReservedCounter().msg()
	}

	if value == PatternFlowRsvpReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRsvpReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRsvpReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRsvpReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRsvpReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRsvpReserved object
func (obj *patternFlowRsvpReserved) SetValue(value uint32) PatternFlowRsvpReserved {
	obj.setChoice(PatternFlowRsvpReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRsvpReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRsvpReserved object
func (obj *patternFlowRsvpReserved) SetValues(value []uint32) PatternFlowRsvpReserved {
	obj.setChoice(PatternFlowRsvpReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRsvpReservedCounter
func (obj *patternFlowRsvpReserved) Increment() PatternFlowRsvpReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRsvpReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRsvpReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRsvpReservedCounter
func (obj *patternFlowRsvpReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRsvpReservedCounter value in the PatternFlowRsvpReserved object
func (obj *patternFlowRsvpReserved) SetIncrement(value PatternFlowRsvpReservedCounter) PatternFlowRsvpReserved {
	obj.setChoice(PatternFlowRsvpReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRsvpReservedCounter
func (obj *patternFlowRsvpReserved) Decrement() PatternFlowRsvpReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRsvpReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRsvpReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRsvpReservedCounter
func (obj *patternFlowRsvpReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRsvpReservedCounter value in the PatternFlowRsvpReserved object
func (obj *patternFlowRsvpReserved) SetDecrement(value PatternFlowRsvpReservedCounter) PatternFlowRsvpReserved {
	obj.setChoice(PatternFlowRsvpReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRsvpReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRsvpReserved.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRsvpReserved.Values <= 255 but Got %d", item))
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

func (obj *patternFlowRsvpReserved) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRsvpReservedChoice.VALUE)

	}

}
