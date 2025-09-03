package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerPortNumber *****
type patternFlowLacpPartnerPortNumber struct {
	validation
	obj             *otg.PatternFlowLacpPartnerPortNumber
	marshaller      marshalPatternFlowLacpPartnerPortNumber
	unMarshaller    unMarshalPatternFlowLacpPartnerPortNumber
	incrementHolder PatternFlowLacpPartnerPortNumberCounter
	decrementHolder PatternFlowLacpPartnerPortNumberCounter
}

func NewPatternFlowLacpPartnerPortNumber() PatternFlowLacpPartnerPortNumber {
	obj := patternFlowLacpPartnerPortNumber{obj: &otg.PatternFlowLacpPartnerPortNumber{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerPortNumber) msg() *otg.PatternFlowLacpPartnerPortNumber {
	return obj.obj
}

func (obj *patternFlowLacpPartnerPortNumber) setMsg(msg *otg.PatternFlowLacpPartnerPortNumber) PatternFlowLacpPartnerPortNumber {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerPortNumber struct {
	obj *patternFlowLacpPartnerPortNumber
}

type marshalPatternFlowLacpPartnerPortNumber interface {
	// ToProto marshals PatternFlowLacpPartnerPortNumber to protobuf object *otg.PatternFlowLacpPartnerPortNumber
	ToProto() (*otg.PatternFlowLacpPartnerPortNumber, error)
	// ToPbText marshals PatternFlowLacpPartnerPortNumber to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerPortNumber to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerPortNumber to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerPortNumber struct {
	obj *patternFlowLacpPartnerPortNumber
}

type unMarshalPatternFlowLacpPartnerPortNumber interface {
	// FromProto unmarshals PatternFlowLacpPartnerPortNumber from protobuf object *otg.PatternFlowLacpPartnerPortNumber
	FromProto(msg *otg.PatternFlowLacpPartnerPortNumber) (PatternFlowLacpPartnerPortNumber, error)
	// FromPbText unmarshals PatternFlowLacpPartnerPortNumber from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerPortNumber from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerPortNumber from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerPortNumber) Marshal() marshalPatternFlowLacpPartnerPortNumber {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerPortNumber{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerPortNumber) Unmarshal() unMarshalPatternFlowLacpPartnerPortNumber {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerPortNumber{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerPortNumber) ToProto() (*otg.PatternFlowLacpPartnerPortNumber, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerPortNumber) FromProto(msg *otg.PatternFlowLacpPartnerPortNumber) (PatternFlowLacpPartnerPortNumber, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerPortNumber) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerPortNumber) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerPortNumber) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerPortNumber) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerPortNumber) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerPortNumber) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerPortNumber) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerPortNumber) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerPortNumber) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerPortNumber) Clone() (PatternFlowLacpPartnerPortNumber, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerPortNumber()
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

func (obj *patternFlowLacpPartnerPortNumber) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerPortNumber is the port number of the Partner's port, as received by the Actor.
type PatternFlowLacpPartnerPortNumber interface {
	Validation
	// msg marshals PatternFlowLacpPartnerPortNumber to protobuf object *otg.PatternFlowLacpPartnerPortNumber
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerPortNumber
	// setMsg unmarshals PatternFlowLacpPartnerPortNumber from protobuf object *otg.PatternFlowLacpPartnerPortNumber
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerPortNumber) PatternFlowLacpPartnerPortNumber
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerPortNumber
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerPortNumber
	// validate validates PatternFlowLacpPartnerPortNumber
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerPortNumber, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerPortNumberChoiceEnum, set in PatternFlowLacpPartnerPortNumber
	Choice() PatternFlowLacpPartnerPortNumberChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerPortNumberChoiceEnum provided by user to PatternFlowLacpPartnerPortNumber
	setChoice(value PatternFlowLacpPartnerPortNumberChoiceEnum) PatternFlowLacpPartnerPortNumber
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerPortNumber
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerPortNumber.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerPortNumber
	SetValue(value uint32) PatternFlowLacpPartnerPortNumber
	// HasValue checks if Value has been set in PatternFlowLacpPartnerPortNumber
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerPortNumber.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerPortNumber
	SetValues(value []uint32) PatternFlowLacpPartnerPortNumber
	// Increment returns PatternFlowLacpPartnerPortNumberCounter, set in PatternFlowLacpPartnerPortNumber.
	// PatternFlowLacpPartnerPortNumberCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerPortNumberCounter
	// SetIncrement assigns PatternFlowLacpPartnerPortNumberCounter provided by user to PatternFlowLacpPartnerPortNumber.
	// PatternFlowLacpPartnerPortNumberCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerPortNumberCounter) PatternFlowLacpPartnerPortNumber
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerPortNumber
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerPortNumberCounter, set in PatternFlowLacpPartnerPortNumber.
	// PatternFlowLacpPartnerPortNumberCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerPortNumberCounter
	// SetDecrement assigns PatternFlowLacpPartnerPortNumberCounter provided by user to PatternFlowLacpPartnerPortNumber.
	// PatternFlowLacpPartnerPortNumberCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerPortNumberCounter) PatternFlowLacpPartnerPortNumber
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerPortNumber
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerPortNumberChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerPortNumber
var PatternFlowLacpPartnerPortNumberChoice = struct {
	VALUE     PatternFlowLacpPartnerPortNumberChoiceEnum
	VALUES    PatternFlowLacpPartnerPortNumberChoiceEnum
	INCREMENT PatternFlowLacpPartnerPortNumberChoiceEnum
	DECREMENT PatternFlowLacpPartnerPortNumberChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerPortNumberChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerPortNumberChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerPortNumberChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerPortNumberChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerPortNumber) Choice() PatternFlowLacpPartnerPortNumberChoiceEnum {
	return PatternFlowLacpPartnerPortNumberChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerPortNumber) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerPortNumber) setChoice(value PatternFlowLacpPartnerPortNumberChoiceEnum) PatternFlowLacpPartnerPortNumber {
	intValue, ok := otg.PatternFlowLacpPartnerPortNumber_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerPortNumberChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerPortNumber_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerPortNumberChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerPortNumberChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerPortNumberChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerPortNumberCounter().msg()
	}

	if value == PatternFlowLacpPartnerPortNumberChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerPortNumberCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerPortNumber) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerPortNumberChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerPortNumber) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerPortNumber object
func (obj *patternFlowLacpPartnerPortNumber) SetValue(value uint32) PatternFlowLacpPartnerPortNumber {
	obj.setChoice(PatternFlowLacpPartnerPortNumberChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerPortNumber) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerPortNumber object
func (obj *patternFlowLacpPartnerPortNumber) SetValues(value []uint32) PatternFlowLacpPartnerPortNumber {
	obj.setChoice(PatternFlowLacpPartnerPortNumberChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerPortNumberCounter
func (obj *patternFlowLacpPartnerPortNumber) Increment() PatternFlowLacpPartnerPortNumberCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerPortNumberChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerPortNumberCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerPortNumberCounter
func (obj *patternFlowLacpPartnerPortNumber) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerPortNumberCounter value in the PatternFlowLacpPartnerPortNumber object
func (obj *patternFlowLacpPartnerPortNumber) SetIncrement(value PatternFlowLacpPartnerPortNumberCounter) PatternFlowLacpPartnerPortNumber {
	obj.setChoice(PatternFlowLacpPartnerPortNumberChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerPortNumberCounter
func (obj *patternFlowLacpPartnerPortNumber) Decrement() PatternFlowLacpPartnerPortNumberCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerPortNumberChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerPortNumberCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerPortNumberCounter
func (obj *patternFlowLacpPartnerPortNumber) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerPortNumberCounter value in the PatternFlowLacpPartnerPortNumber object
func (obj *patternFlowLacpPartnerPortNumber) SetDecrement(value PatternFlowLacpPartnerPortNumberCounter) PatternFlowLacpPartnerPortNumber {
	obj.setChoice(PatternFlowLacpPartnerPortNumberChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerPortNumber) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerPortNumber.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerPortNumber.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpPartnerPortNumber) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerPortNumberChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerPortNumberChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerPortNumberChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerPortNumberChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerPortNumberChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerPortNumberChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerPortNumber")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerPortNumber_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerPortNumber_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
