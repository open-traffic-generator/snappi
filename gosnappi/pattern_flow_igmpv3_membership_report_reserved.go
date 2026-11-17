package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipReportReserved *****
type patternFlowIgmpv3MembershipReportReserved struct {
	validation
	obj             *otg.PatternFlowIgmpv3MembershipReportReserved
	marshaller      marshalPatternFlowIgmpv3MembershipReportReserved
	unMarshaller    unMarshalPatternFlowIgmpv3MembershipReportReserved
	incrementHolder PatternFlowIgmpv3MembershipReportReservedCounter
	decrementHolder PatternFlowIgmpv3MembershipReportReservedCounter
}

func NewPatternFlowIgmpv3MembershipReportReserved() PatternFlowIgmpv3MembershipReportReserved {
	obj := patternFlowIgmpv3MembershipReportReserved{obj: &otg.PatternFlowIgmpv3MembershipReportReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipReportReserved) msg() *otg.PatternFlowIgmpv3MembershipReportReserved {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipReportReserved) setMsg(msg *otg.PatternFlowIgmpv3MembershipReportReserved) PatternFlowIgmpv3MembershipReportReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipReportReserved struct {
	obj *patternFlowIgmpv3MembershipReportReserved
}

type marshalPatternFlowIgmpv3MembershipReportReserved interface {
	// ToProto marshals PatternFlowIgmpv3MembershipReportReserved to protobuf object *otg.PatternFlowIgmpv3MembershipReportReserved
	ToProto() (*otg.PatternFlowIgmpv3MembershipReportReserved, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipReportReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipReportReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipReportReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipReportReserved struct {
	obj *patternFlowIgmpv3MembershipReportReserved
}

type unMarshalPatternFlowIgmpv3MembershipReportReserved interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipReportReserved from protobuf object *otg.PatternFlowIgmpv3MembershipReportReserved
	FromProto(msg *otg.PatternFlowIgmpv3MembershipReportReserved) (PatternFlowIgmpv3MembershipReportReserved, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipReportReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipReportReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipReportReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipReportReserved) Marshal() marshalPatternFlowIgmpv3MembershipReportReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipReportReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipReportReserved) Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipReportReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipReportReserved) ToProto() (*otg.PatternFlowIgmpv3MembershipReportReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipReportReserved) FromProto(msg *otg.PatternFlowIgmpv3MembershipReportReserved) (PatternFlowIgmpv3MembershipReportReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipReportReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportReserved) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipReportReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipReportReserved) Clone() (PatternFlowIgmpv3MembershipReportReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipReportReserved()
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

func (obj *patternFlowIgmpv3MembershipReportReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3MembershipReportReserved is reserved field, set to zero by the sender and ignored by the receiver.
type PatternFlowIgmpv3MembershipReportReserved interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipReportReserved to protobuf object *otg.PatternFlowIgmpv3MembershipReportReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipReportReserved
	// setMsg unmarshals PatternFlowIgmpv3MembershipReportReserved from protobuf object *otg.PatternFlowIgmpv3MembershipReportReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipReportReserved) PatternFlowIgmpv3MembershipReportReserved
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipReportReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportReserved
	// validate validates PatternFlowIgmpv3MembershipReportReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipReportReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipReportReservedChoiceEnum, set in PatternFlowIgmpv3MembershipReportReserved
	Choice() PatternFlowIgmpv3MembershipReportReservedChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipReportReservedChoiceEnum provided by user to PatternFlowIgmpv3MembershipReportReserved
	setChoice(value PatternFlowIgmpv3MembershipReportReservedChoiceEnum) PatternFlowIgmpv3MembershipReportReserved
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipReportReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3MembershipReportReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportReserved
	SetValue(value uint32) PatternFlowIgmpv3MembershipReportReserved
	// HasValue checks if Value has been set in PatternFlowIgmpv3MembershipReportReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3MembershipReportReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3MembershipReportReserved
	SetValues(value []uint32) PatternFlowIgmpv3MembershipReportReserved
	// Increment returns PatternFlowIgmpv3MembershipReportReservedCounter, set in PatternFlowIgmpv3MembershipReportReserved.
	// PatternFlowIgmpv3MembershipReportReservedCounter is integer counter pattern
	Increment() PatternFlowIgmpv3MembershipReportReservedCounter
	// SetIncrement assigns PatternFlowIgmpv3MembershipReportReservedCounter provided by user to PatternFlowIgmpv3MembershipReportReserved.
	// PatternFlowIgmpv3MembershipReportReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3MembershipReportReservedCounter) PatternFlowIgmpv3MembershipReportReserved
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3MembershipReportReserved
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3MembershipReportReservedCounter, set in PatternFlowIgmpv3MembershipReportReserved.
	// PatternFlowIgmpv3MembershipReportReservedCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3MembershipReportReservedCounter
	// SetDecrement assigns PatternFlowIgmpv3MembershipReportReservedCounter provided by user to PatternFlowIgmpv3MembershipReportReserved.
	// PatternFlowIgmpv3MembershipReportReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3MembershipReportReservedCounter) PatternFlowIgmpv3MembershipReportReserved
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3MembershipReportReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3MembershipReportReservedChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipReportReserved
var PatternFlowIgmpv3MembershipReportReservedChoice = struct {
	VALUE     PatternFlowIgmpv3MembershipReportReservedChoiceEnum
	VALUES    PatternFlowIgmpv3MembershipReportReservedChoiceEnum
	INCREMENT PatternFlowIgmpv3MembershipReportReservedChoiceEnum
	DECREMENT PatternFlowIgmpv3MembershipReportReservedChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3MembershipReportReservedChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3MembershipReportReservedChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv3MembershipReportReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3MembershipReportReservedChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3MembershipReportReserved) Choice() PatternFlowIgmpv3MembershipReportReservedChoiceEnum {
	return PatternFlowIgmpv3MembershipReportReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipReportReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipReportReserved) setChoice(value PatternFlowIgmpv3MembershipReportReservedChoiceEnum) PatternFlowIgmpv3MembershipReportReserved {
	intValue, ok := otg.PatternFlowIgmpv3MembershipReportReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipReportReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipReportReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3MembershipReportReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipReportReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3MembershipReportReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3MembershipReportReservedCounter().msg()
	}

	if value == PatternFlowIgmpv3MembershipReportReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3MembershipReportReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipReportReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipReportReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3MembershipReportReserved object
func (obj *patternFlowIgmpv3MembershipReportReserved) SetValue(value uint32) PatternFlowIgmpv3MembershipReportReserved {
	obj.setChoice(PatternFlowIgmpv3MembershipReportReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3MembershipReportReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3MembershipReportReserved object
func (obj *patternFlowIgmpv3MembershipReportReserved) SetValues(value []uint32) PatternFlowIgmpv3MembershipReportReserved {
	obj.setChoice(PatternFlowIgmpv3MembershipReportReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipReportReservedCounter
func (obj *patternFlowIgmpv3MembershipReportReserved) Increment() PatternFlowIgmpv3MembershipReportReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3MembershipReportReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipReportReservedCounter
func (obj *patternFlowIgmpv3MembershipReportReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3MembershipReportReservedCounter value in the PatternFlowIgmpv3MembershipReportReserved object
func (obj *patternFlowIgmpv3MembershipReportReserved) SetIncrement(value PatternFlowIgmpv3MembershipReportReservedCounter) PatternFlowIgmpv3MembershipReportReserved {
	obj.setChoice(PatternFlowIgmpv3MembershipReportReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipReportReservedCounter
func (obj *patternFlowIgmpv3MembershipReportReserved) Decrement() PatternFlowIgmpv3MembershipReportReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3MembershipReportReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipReportReservedCounter
func (obj *patternFlowIgmpv3MembershipReportReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3MembershipReportReservedCounter value in the PatternFlowIgmpv3MembershipReportReserved object
func (obj *patternFlowIgmpv3MembershipReportReserved) SetDecrement(value PatternFlowIgmpv3MembershipReportReservedCounter) PatternFlowIgmpv3MembershipReportReserved {
	obj.setChoice(PatternFlowIgmpv3MembershipReportReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3MembershipReportReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportReserved.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3MembershipReportReserved.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIgmpv3MembershipReportReserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipReportReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipReportReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipReportReserved")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipReportReserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipReportReserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
