package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorPortNumber *****
type patternFlowLacpActorPortNumber struct {
	validation
	obj             *otg.PatternFlowLacpActorPortNumber
	marshaller      marshalPatternFlowLacpActorPortNumber
	unMarshaller    unMarshalPatternFlowLacpActorPortNumber
	incrementHolder PatternFlowLacpActorPortNumberCounter
	decrementHolder PatternFlowLacpActorPortNumberCounter
}

func NewPatternFlowLacpActorPortNumber() PatternFlowLacpActorPortNumber {
	obj := patternFlowLacpActorPortNumber{obj: &otg.PatternFlowLacpActorPortNumber{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorPortNumber) msg() *otg.PatternFlowLacpActorPortNumber {
	return obj.obj
}

func (obj *patternFlowLacpActorPortNumber) setMsg(msg *otg.PatternFlowLacpActorPortNumber) PatternFlowLacpActorPortNumber {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorPortNumber struct {
	obj *patternFlowLacpActorPortNumber
}

type marshalPatternFlowLacpActorPortNumber interface {
	// ToProto marshals PatternFlowLacpActorPortNumber to protobuf object *otg.PatternFlowLacpActorPortNumber
	ToProto() (*otg.PatternFlowLacpActorPortNumber, error)
	// ToPbText marshals PatternFlowLacpActorPortNumber to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorPortNumber to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorPortNumber to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorPortNumber struct {
	obj *patternFlowLacpActorPortNumber
}

type unMarshalPatternFlowLacpActorPortNumber interface {
	// FromProto unmarshals PatternFlowLacpActorPortNumber from protobuf object *otg.PatternFlowLacpActorPortNumber
	FromProto(msg *otg.PatternFlowLacpActorPortNumber) (PatternFlowLacpActorPortNumber, error)
	// FromPbText unmarshals PatternFlowLacpActorPortNumber from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorPortNumber from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorPortNumber from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorPortNumber) Marshal() marshalPatternFlowLacpActorPortNumber {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorPortNumber{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorPortNumber) Unmarshal() unMarshalPatternFlowLacpActorPortNumber {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorPortNumber{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorPortNumber) ToProto() (*otg.PatternFlowLacpActorPortNumber, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorPortNumber) FromProto(msg *otg.PatternFlowLacpActorPortNumber) (PatternFlowLacpActorPortNumber, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorPortNumber) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorPortNumber) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorPortNumber) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorPortNumber) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorPortNumber) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorPortNumber) FromJson(value string) error {
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

func (obj *patternFlowLacpActorPortNumber) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorPortNumber) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorPortNumber) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorPortNumber) Clone() (PatternFlowLacpActorPortNumber, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorPortNumber()
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

func (obj *patternFlowLacpActorPortNumber) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorPortNumber is the port number assigned to the Actor's port. It is a unique identifier for the port within the system.
type PatternFlowLacpActorPortNumber interface {
	Validation
	// msg marshals PatternFlowLacpActorPortNumber to protobuf object *otg.PatternFlowLacpActorPortNumber
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorPortNumber
	// setMsg unmarshals PatternFlowLacpActorPortNumber from protobuf object *otg.PatternFlowLacpActorPortNumber
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorPortNumber) PatternFlowLacpActorPortNumber
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorPortNumber
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorPortNumber
	// validate validates PatternFlowLacpActorPortNumber
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorPortNumber, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorPortNumberChoiceEnum, set in PatternFlowLacpActorPortNumber
	Choice() PatternFlowLacpActorPortNumberChoiceEnum
	// setChoice assigns PatternFlowLacpActorPortNumberChoiceEnum provided by user to PatternFlowLacpActorPortNumber
	setChoice(value PatternFlowLacpActorPortNumberChoiceEnum) PatternFlowLacpActorPortNumber
	// HasChoice checks if Choice has been set in PatternFlowLacpActorPortNumber
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorPortNumber.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorPortNumber
	SetValue(value uint32) PatternFlowLacpActorPortNumber
	// HasValue checks if Value has been set in PatternFlowLacpActorPortNumber
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorPortNumber.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorPortNumber
	SetValues(value []uint32) PatternFlowLacpActorPortNumber
	// Increment returns PatternFlowLacpActorPortNumberCounter, set in PatternFlowLacpActorPortNumber.
	// PatternFlowLacpActorPortNumberCounter is integer counter pattern
	Increment() PatternFlowLacpActorPortNumberCounter
	// SetIncrement assigns PatternFlowLacpActorPortNumberCounter provided by user to PatternFlowLacpActorPortNumber.
	// PatternFlowLacpActorPortNumberCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorPortNumberCounter) PatternFlowLacpActorPortNumber
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorPortNumber
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorPortNumberCounter, set in PatternFlowLacpActorPortNumber.
	// PatternFlowLacpActorPortNumberCounter is integer counter pattern
	Decrement() PatternFlowLacpActorPortNumberCounter
	// SetDecrement assigns PatternFlowLacpActorPortNumberCounter provided by user to PatternFlowLacpActorPortNumber.
	// PatternFlowLacpActorPortNumberCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorPortNumberCounter) PatternFlowLacpActorPortNumber
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorPortNumber
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorPortNumberChoiceEnum string

// Enum of Choice on PatternFlowLacpActorPortNumber
var PatternFlowLacpActorPortNumberChoice = struct {
	VALUE     PatternFlowLacpActorPortNumberChoiceEnum
	VALUES    PatternFlowLacpActorPortNumberChoiceEnum
	INCREMENT PatternFlowLacpActorPortNumberChoiceEnum
	DECREMENT PatternFlowLacpActorPortNumberChoiceEnum
}{
	VALUE:     PatternFlowLacpActorPortNumberChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorPortNumberChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorPortNumberChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorPortNumberChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorPortNumber) Choice() PatternFlowLacpActorPortNumberChoiceEnum {
	return PatternFlowLacpActorPortNumberChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorPortNumber) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorPortNumber) setChoice(value PatternFlowLacpActorPortNumberChoiceEnum) PatternFlowLacpActorPortNumber {
	intValue, ok := otg.PatternFlowLacpActorPortNumber_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorPortNumberChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorPortNumber_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorPortNumberChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorPortNumberChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorPortNumberChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorPortNumberCounter().msg()
	}

	if value == PatternFlowLacpActorPortNumberChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorPortNumberCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorPortNumber) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorPortNumberChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorPortNumber) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorPortNumber object
func (obj *patternFlowLacpActorPortNumber) SetValue(value uint32) PatternFlowLacpActorPortNumber {
	obj.setChoice(PatternFlowLacpActorPortNumberChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorPortNumber) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorPortNumber object
func (obj *patternFlowLacpActorPortNumber) SetValues(value []uint32) PatternFlowLacpActorPortNumber {
	obj.setChoice(PatternFlowLacpActorPortNumberChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorPortNumberCounter
func (obj *patternFlowLacpActorPortNumber) Increment() PatternFlowLacpActorPortNumberCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorPortNumberChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorPortNumberCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorPortNumberCounter
func (obj *patternFlowLacpActorPortNumber) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorPortNumberCounter value in the PatternFlowLacpActorPortNumber object
func (obj *patternFlowLacpActorPortNumber) SetIncrement(value PatternFlowLacpActorPortNumberCounter) PatternFlowLacpActorPortNumber {
	obj.setChoice(PatternFlowLacpActorPortNumberChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorPortNumberCounter
func (obj *patternFlowLacpActorPortNumber) Decrement() PatternFlowLacpActorPortNumberCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorPortNumberChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorPortNumberCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorPortNumberCounter
func (obj *patternFlowLacpActorPortNumber) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorPortNumberCounter value in the PatternFlowLacpActorPortNumber object
func (obj *patternFlowLacpActorPortNumber) SetDecrement(value PatternFlowLacpActorPortNumberCounter) PatternFlowLacpActorPortNumber {
	obj.setChoice(PatternFlowLacpActorPortNumberChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorPortNumber) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorPortNumber.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorPortNumber.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpActorPortNumber) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorPortNumberChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorPortNumberChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorPortNumberChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorPortNumberChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorPortNumberChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorPortNumberChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorPortNumber")
			}
		} else {
			intVal := otg.PatternFlowLacpActorPortNumber_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorPortNumber_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
