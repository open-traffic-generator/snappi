package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerReserved *****
type patternFlowLacpduPartnerReserved struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerReserved
	marshaller      marshalPatternFlowLacpduPartnerReserved
	unMarshaller    unMarshalPatternFlowLacpduPartnerReserved
	incrementHolder PatternFlowLacpduPartnerReservedCounter
	decrementHolder PatternFlowLacpduPartnerReservedCounter
}

func NewPatternFlowLacpduPartnerReserved() PatternFlowLacpduPartnerReserved {
	obj := patternFlowLacpduPartnerReserved{obj: &otg.PatternFlowLacpduPartnerReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerReserved) msg() *otg.PatternFlowLacpduPartnerReserved {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerReserved) setMsg(msg *otg.PatternFlowLacpduPartnerReserved) PatternFlowLacpduPartnerReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerReserved struct {
	obj *patternFlowLacpduPartnerReserved
}

type marshalPatternFlowLacpduPartnerReserved interface {
	// ToProto marshals PatternFlowLacpduPartnerReserved to protobuf object *otg.PatternFlowLacpduPartnerReserved
	ToProto() (*otg.PatternFlowLacpduPartnerReserved, error)
	// ToPbText marshals PatternFlowLacpduPartnerReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerReserved struct {
	obj *patternFlowLacpduPartnerReserved
}

type unMarshalPatternFlowLacpduPartnerReserved interface {
	// FromProto unmarshals PatternFlowLacpduPartnerReserved from protobuf object *otg.PatternFlowLacpduPartnerReserved
	FromProto(msg *otg.PatternFlowLacpduPartnerReserved) (PatternFlowLacpduPartnerReserved, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerReserved) Marshal() marshalPatternFlowLacpduPartnerReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerReserved) Unmarshal() unMarshalPatternFlowLacpduPartnerReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerReserved) ToProto() (*otg.PatternFlowLacpduPartnerReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerReserved) FromProto(msg *otg.PatternFlowLacpduPartnerReserved) (PatternFlowLacpduPartnerReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerReserved) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerReserved) Clone() (PatternFlowLacpduPartnerReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerReserved()
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

func (obj *patternFlowLacpduPartnerReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerReserved is reserved field for future use in the Partner TLV. Should be set to all zeros.
type PatternFlowLacpduPartnerReserved interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerReserved to protobuf object *otg.PatternFlowLacpduPartnerReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerReserved
	// setMsg unmarshals PatternFlowLacpduPartnerReserved from protobuf object *otg.PatternFlowLacpduPartnerReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerReserved) PatternFlowLacpduPartnerReserved
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerReserved
	// validate validates PatternFlowLacpduPartnerReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerReservedChoiceEnum, set in PatternFlowLacpduPartnerReserved
	Choice() PatternFlowLacpduPartnerReservedChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerReservedChoiceEnum provided by user to PatternFlowLacpduPartnerReserved
	setChoice(value PatternFlowLacpduPartnerReservedChoiceEnum) PatternFlowLacpduPartnerReserved
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerReserved
	SetValue(value uint32) PatternFlowLacpduPartnerReserved
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerReserved
	SetValues(value []uint32) PatternFlowLacpduPartnerReserved
	// Increment returns PatternFlowLacpduPartnerReservedCounter, set in PatternFlowLacpduPartnerReserved.
	// PatternFlowLacpduPartnerReservedCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerReservedCounter
	// SetIncrement assigns PatternFlowLacpduPartnerReservedCounter provided by user to PatternFlowLacpduPartnerReserved.
	// PatternFlowLacpduPartnerReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerReservedCounter) PatternFlowLacpduPartnerReserved
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerReserved
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerReservedCounter, set in PatternFlowLacpduPartnerReserved.
	// PatternFlowLacpduPartnerReservedCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerReservedCounter
	// SetDecrement assigns PatternFlowLacpduPartnerReservedCounter provided by user to PatternFlowLacpduPartnerReserved.
	// PatternFlowLacpduPartnerReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerReservedCounter) PatternFlowLacpduPartnerReserved
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerReservedChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerReserved
var PatternFlowLacpduPartnerReservedChoice = struct {
	VALUE     PatternFlowLacpduPartnerReservedChoiceEnum
	VALUES    PatternFlowLacpduPartnerReservedChoiceEnum
	INCREMENT PatternFlowLacpduPartnerReservedChoiceEnum
	DECREMENT PatternFlowLacpduPartnerReservedChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerReservedChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerReservedChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerReservedChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerReserved) Choice() PatternFlowLacpduPartnerReservedChoiceEnum {
	return PatternFlowLacpduPartnerReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerReserved) setChoice(value PatternFlowLacpduPartnerReservedChoiceEnum) PatternFlowLacpduPartnerReserved {
	intValue, ok := otg.PatternFlowLacpduPartnerReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerReservedCounter().msg()
	}

	if value == PatternFlowLacpduPartnerReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerReserved object
func (obj *patternFlowLacpduPartnerReserved) SetValue(value uint32) PatternFlowLacpduPartnerReserved {
	obj.setChoice(PatternFlowLacpduPartnerReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerReserved object
func (obj *patternFlowLacpduPartnerReserved) SetValues(value []uint32) PatternFlowLacpduPartnerReserved {
	obj.setChoice(PatternFlowLacpduPartnerReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerReservedCounter
func (obj *patternFlowLacpduPartnerReserved) Increment() PatternFlowLacpduPartnerReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerReservedCounter
func (obj *patternFlowLacpduPartnerReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerReservedCounter value in the PatternFlowLacpduPartnerReserved object
func (obj *patternFlowLacpduPartnerReserved) SetIncrement(value PatternFlowLacpduPartnerReservedCounter) PatternFlowLacpduPartnerReserved {
	obj.setChoice(PatternFlowLacpduPartnerReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerReservedCounter
func (obj *patternFlowLacpduPartnerReserved) Decrement() PatternFlowLacpduPartnerReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerReservedCounter
func (obj *patternFlowLacpduPartnerReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerReservedCounter value in the PatternFlowLacpduPartnerReserved object
func (obj *patternFlowLacpduPartnerReserved) SetDecrement(value PatternFlowLacpduPartnerReservedCounter) PatternFlowLacpduPartnerReserved {
	obj.setChoice(PatternFlowLacpduPartnerReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerReserved.Value <= 16777215 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 16777215 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerReserved.Values <= 16777215 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerReserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerReserved")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerReserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerReserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
