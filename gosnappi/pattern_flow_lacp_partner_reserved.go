package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerReserved *****
type patternFlowLacpPartnerReserved struct {
	validation
	obj             *otg.PatternFlowLacpPartnerReserved
	marshaller      marshalPatternFlowLacpPartnerReserved
	unMarshaller    unMarshalPatternFlowLacpPartnerReserved
	incrementHolder PatternFlowLacpPartnerReservedCounter
	decrementHolder PatternFlowLacpPartnerReservedCounter
}

func NewPatternFlowLacpPartnerReserved() PatternFlowLacpPartnerReserved {
	obj := patternFlowLacpPartnerReserved{obj: &otg.PatternFlowLacpPartnerReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerReserved) msg() *otg.PatternFlowLacpPartnerReserved {
	return obj.obj
}

func (obj *patternFlowLacpPartnerReserved) setMsg(msg *otg.PatternFlowLacpPartnerReserved) PatternFlowLacpPartnerReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerReserved struct {
	obj *patternFlowLacpPartnerReserved
}

type marshalPatternFlowLacpPartnerReserved interface {
	// ToProto marshals PatternFlowLacpPartnerReserved to protobuf object *otg.PatternFlowLacpPartnerReserved
	ToProto() (*otg.PatternFlowLacpPartnerReserved, error)
	// ToPbText marshals PatternFlowLacpPartnerReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerReserved struct {
	obj *patternFlowLacpPartnerReserved
}

type unMarshalPatternFlowLacpPartnerReserved interface {
	// FromProto unmarshals PatternFlowLacpPartnerReserved from protobuf object *otg.PatternFlowLacpPartnerReserved
	FromProto(msg *otg.PatternFlowLacpPartnerReserved) (PatternFlowLacpPartnerReserved, error)
	// FromPbText unmarshals PatternFlowLacpPartnerReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerReserved) Marshal() marshalPatternFlowLacpPartnerReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerReserved) Unmarshal() unMarshalPatternFlowLacpPartnerReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerReserved) ToProto() (*otg.PatternFlowLacpPartnerReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerReserved) FromProto(msg *otg.PatternFlowLacpPartnerReserved) (PatternFlowLacpPartnerReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerReserved) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerReserved) Clone() (PatternFlowLacpPartnerReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerReserved()
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

func (obj *patternFlowLacpPartnerReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerReserved is reserved field for future use in the Partner TLV. Should be set to all zeros.
type PatternFlowLacpPartnerReserved interface {
	Validation
	// msg marshals PatternFlowLacpPartnerReserved to protobuf object *otg.PatternFlowLacpPartnerReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerReserved
	// setMsg unmarshals PatternFlowLacpPartnerReserved from protobuf object *otg.PatternFlowLacpPartnerReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerReserved) PatternFlowLacpPartnerReserved
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerReserved
	// validate validates PatternFlowLacpPartnerReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerReservedChoiceEnum, set in PatternFlowLacpPartnerReserved
	Choice() PatternFlowLacpPartnerReservedChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerReservedChoiceEnum provided by user to PatternFlowLacpPartnerReserved
	setChoice(value PatternFlowLacpPartnerReservedChoiceEnum) PatternFlowLacpPartnerReserved
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerReserved
	SetValue(value uint32) PatternFlowLacpPartnerReserved
	// HasValue checks if Value has been set in PatternFlowLacpPartnerReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerReserved
	SetValues(value []uint32) PatternFlowLacpPartnerReserved
	// Increment returns PatternFlowLacpPartnerReservedCounter, set in PatternFlowLacpPartnerReserved.
	// PatternFlowLacpPartnerReservedCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerReservedCounter
	// SetIncrement assigns PatternFlowLacpPartnerReservedCounter provided by user to PatternFlowLacpPartnerReserved.
	// PatternFlowLacpPartnerReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerReservedCounter) PatternFlowLacpPartnerReserved
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerReserved
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerReservedCounter, set in PatternFlowLacpPartnerReserved.
	// PatternFlowLacpPartnerReservedCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerReservedCounter
	// SetDecrement assigns PatternFlowLacpPartnerReservedCounter provided by user to PatternFlowLacpPartnerReserved.
	// PatternFlowLacpPartnerReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerReservedCounter) PatternFlowLacpPartnerReserved
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerReservedChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerReserved
var PatternFlowLacpPartnerReservedChoice = struct {
	VALUE     PatternFlowLacpPartnerReservedChoiceEnum
	VALUES    PatternFlowLacpPartnerReservedChoiceEnum
	INCREMENT PatternFlowLacpPartnerReservedChoiceEnum
	DECREMENT PatternFlowLacpPartnerReservedChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerReservedChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerReservedChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerReservedChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerReserved) Choice() PatternFlowLacpPartnerReservedChoiceEnum {
	return PatternFlowLacpPartnerReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerReserved) setChoice(value PatternFlowLacpPartnerReservedChoiceEnum) PatternFlowLacpPartnerReserved {
	intValue, ok := otg.PatternFlowLacpPartnerReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerReservedCounter().msg()
	}

	if value == PatternFlowLacpPartnerReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerReserved object
func (obj *patternFlowLacpPartnerReserved) SetValue(value uint32) PatternFlowLacpPartnerReserved {
	obj.setChoice(PatternFlowLacpPartnerReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerReserved object
func (obj *patternFlowLacpPartnerReserved) SetValues(value []uint32) PatternFlowLacpPartnerReserved {
	obj.setChoice(PatternFlowLacpPartnerReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerReservedCounter
func (obj *patternFlowLacpPartnerReserved) Increment() PatternFlowLacpPartnerReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerReservedCounter
func (obj *patternFlowLacpPartnerReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerReservedCounter value in the PatternFlowLacpPartnerReserved object
func (obj *patternFlowLacpPartnerReserved) SetIncrement(value PatternFlowLacpPartnerReservedCounter) PatternFlowLacpPartnerReserved {
	obj.setChoice(PatternFlowLacpPartnerReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerReservedCounter
func (obj *patternFlowLacpPartnerReserved) Decrement() PatternFlowLacpPartnerReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerReservedCounter
func (obj *patternFlowLacpPartnerReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerReservedCounter value in the PatternFlowLacpPartnerReserved object
func (obj *patternFlowLacpPartnerReserved) SetDecrement(value PatternFlowLacpPartnerReservedCounter) PatternFlowLacpPartnerReserved {
	obj.setChoice(PatternFlowLacpPartnerReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerReserved.Value <= 16777215 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 16777215 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerReserved.Values <= 16777215 but Got %d", item))
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

func (obj *patternFlowLacpPartnerReserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerReserved")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerReserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerReserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
