package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorPortNumber *****
type patternFlowLacpduActorPortNumber struct {
	validation
	obj             *otg.PatternFlowLacpduActorPortNumber
	marshaller      marshalPatternFlowLacpduActorPortNumber
	unMarshaller    unMarshalPatternFlowLacpduActorPortNumber
	incrementHolder PatternFlowLacpduActorPortNumberCounter
	decrementHolder PatternFlowLacpduActorPortNumberCounter
}

func NewPatternFlowLacpduActorPortNumber() PatternFlowLacpduActorPortNumber {
	obj := patternFlowLacpduActorPortNumber{obj: &otg.PatternFlowLacpduActorPortNumber{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorPortNumber) msg() *otg.PatternFlowLacpduActorPortNumber {
	return obj.obj
}

func (obj *patternFlowLacpduActorPortNumber) setMsg(msg *otg.PatternFlowLacpduActorPortNumber) PatternFlowLacpduActorPortNumber {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorPortNumber struct {
	obj *patternFlowLacpduActorPortNumber
}

type marshalPatternFlowLacpduActorPortNumber interface {
	// ToProto marshals PatternFlowLacpduActorPortNumber to protobuf object *otg.PatternFlowLacpduActorPortNumber
	ToProto() (*otg.PatternFlowLacpduActorPortNumber, error)
	// ToPbText marshals PatternFlowLacpduActorPortNumber to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorPortNumber to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorPortNumber to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorPortNumber struct {
	obj *patternFlowLacpduActorPortNumber
}

type unMarshalPatternFlowLacpduActorPortNumber interface {
	// FromProto unmarshals PatternFlowLacpduActorPortNumber from protobuf object *otg.PatternFlowLacpduActorPortNumber
	FromProto(msg *otg.PatternFlowLacpduActorPortNumber) (PatternFlowLacpduActorPortNumber, error)
	// FromPbText unmarshals PatternFlowLacpduActorPortNumber from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorPortNumber from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorPortNumber from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorPortNumber) Marshal() marshalPatternFlowLacpduActorPortNumber {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorPortNumber{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorPortNumber) Unmarshal() unMarshalPatternFlowLacpduActorPortNumber {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorPortNumber{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorPortNumber) ToProto() (*otg.PatternFlowLacpduActorPortNumber, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorPortNumber) FromProto(msg *otg.PatternFlowLacpduActorPortNumber) (PatternFlowLacpduActorPortNumber, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorPortNumber) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorPortNumber) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorPortNumber) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorPortNumber) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorPortNumber) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorPortNumber) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorPortNumber) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorPortNumber) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorPortNumber) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorPortNumber) Clone() (PatternFlowLacpduActorPortNumber, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorPortNumber()
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

func (obj *patternFlowLacpduActorPortNumber) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorPortNumber is the port number assigned to the Actor's port. It is a unique identifier for the port within the system.
type PatternFlowLacpduActorPortNumber interface {
	Validation
	// msg marshals PatternFlowLacpduActorPortNumber to protobuf object *otg.PatternFlowLacpduActorPortNumber
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorPortNumber
	// setMsg unmarshals PatternFlowLacpduActorPortNumber from protobuf object *otg.PatternFlowLacpduActorPortNumber
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorPortNumber) PatternFlowLacpduActorPortNumber
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorPortNumber
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorPortNumber
	// validate validates PatternFlowLacpduActorPortNumber
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorPortNumber, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorPortNumberChoiceEnum, set in PatternFlowLacpduActorPortNumber
	Choice() PatternFlowLacpduActorPortNumberChoiceEnum
	// setChoice assigns PatternFlowLacpduActorPortNumberChoiceEnum provided by user to PatternFlowLacpduActorPortNumber
	setChoice(value PatternFlowLacpduActorPortNumberChoiceEnum) PatternFlowLacpduActorPortNumber
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorPortNumber
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorPortNumber.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorPortNumber
	SetValue(value uint32) PatternFlowLacpduActorPortNumber
	// HasValue checks if Value has been set in PatternFlowLacpduActorPortNumber
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorPortNumber.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorPortNumber
	SetValues(value []uint32) PatternFlowLacpduActorPortNumber
	// Increment returns PatternFlowLacpduActorPortNumberCounter, set in PatternFlowLacpduActorPortNumber.
	// PatternFlowLacpduActorPortNumberCounter is integer counter pattern
	Increment() PatternFlowLacpduActorPortNumberCounter
	// SetIncrement assigns PatternFlowLacpduActorPortNumberCounter provided by user to PatternFlowLacpduActorPortNumber.
	// PatternFlowLacpduActorPortNumberCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorPortNumberCounter) PatternFlowLacpduActorPortNumber
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorPortNumber
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorPortNumberCounter, set in PatternFlowLacpduActorPortNumber.
	// PatternFlowLacpduActorPortNumberCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorPortNumberCounter
	// SetDecrement assigns PatternFlowLacpduActorPortNumberCounter provided by user to PatternFlowLacpduActorPortNumber.
	// PatternFlowLacpduActorPortNumberCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorPortNumberCounter) PatternFlowLacpduActorPortNumber
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorPortNumber
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorPortNumberChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorPortNumber
var PatternFlowLacpduActorPortNumberChoice = struct {
	VALUE     PatternFlowLacpduActorPortNumberChoiceEnum
	VALUES    PatternFlowLacpduActorPortNumberChoiceEnum
	INCREMENT PatternFlowLacpduActorPortNumberChoiceEnum
	DECREMENT PatternFlowLacpduActorPortNumberChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorPortNumberChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorPortNumberChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorPortNumberChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorPortNumberChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorPortNumber) Choice() PatternFlowLacpduActorPortNumberChoiceEnum {
	return PatternFlowLacpduActorPortNumberChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorPortNumber) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorPortNumber) setChoice(value PatternFlowLacpduActorPortNumberChoiceEnum) PatternFlowLacpduActorPortNumber {
	intValue, ok := otg.PatternFlowLacpduActorPortNumber_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorPortNumberChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorPortNumber_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorPortNumberChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorPortNumberChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorPortNumberChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorPortNumberCounter().msg()
	}

	if value == PatternFlowLacpduActorPortNumberChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorPortNumberCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorPortNumber) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorPortNumberChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorPortNumber) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorPortNumber object
func (obj *patternFlowLacpduActorPortNumber) SetValue(value uint32) PatternFlowLacpduActorPortNumber {
	obj.setChoice(PatternFlowLacpduActorPortNumberChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorPortNumber) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorPortNumber object
func (obj *patternFlowLacpduActorPortNumber) SetValues(value []uint32) PatternFlowLacpduActorPortNumber {
	obj.setChoice(PatternFlowLacpduActorPortNumberChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorPortNumberCounter
func (obj *patternFlowLacpduActorPortNumber) Increment() PatternFlowLacpduActorPortNumberCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorPortNumberChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorPortNumberCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorPortNumberCounter
func (obj *patternFlowLacpduActorPortNumber) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorPortNumberCounter value in the PatternFlowLacpduActorPortNumber object
func (obj *patternFlowLacpduActorPortNumber) SetIncrement(value PatternFlowLacpduActorPortNumberCounter) PatternFlowLacpduActorPortNumber {
	obj.setChoice(PatternFlowLacpduActorPortNumberChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorPortNumberCounter
func (obj *patternFlowLacpduActorPortNumber) Decrement() PatternFlowLacpduActorPortNumberCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorPortNumberChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorPortNumberCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorPortNumberCounter
func (obj *patternFlowLacpduActorPortNumber) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorPortNumberCounter value in the PatternFlowLacpduActorPortNumber object
func (obj *patternFlowLacpduActorPortNumber) SetDecrement(value PatternFlowLacpduActorPortNumberCounter) PatternFlowLacpduActorPortNumber {
	obj.setChoice(PatternFlowLacpduActorPortNumberChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorPortNumber) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorPortNumber.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorPortNumber.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpduActorPortNumber) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorPortNumberChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorPortNumberChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorPortNumberChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorPortNumberChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorPortNumberChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorPortNumberChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorPortNumber")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorPortNumber_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorPortNumber_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
