package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerPortNumber *****
type patternFlowLacpduPartnerPortNumber struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerPortNumber
	marshaller      marshalPatternFlowLacpduPartnerPortNumber
	unMarshaller    unMarshalPatternFlowLacpduPartnerPortNumber
	incrementHolder PatternFlowLacpduPartnerPortNumberCounter
	decrementHolder PatternFlowLacpduPartnerPortNumberCounter
}

func NewPatternFlowLacpduPartnerPortNumber() PatternFlowLacpduPartnerPortNumber {
	obj := patternFlowLacpduPartnerPortNumber{obj: &otg.PatternFlowLacpduPartnerPortNumber{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerPortNumber) msg() *otg.PatternFlowLacpduPartnerPortNumber {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerPortNumber) setMsg(msg *otg.PatternFlowLacpduPartnerPortNumber) PatternFlowLacpduPartnerPortNumber {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerPortNumber struct {
	obj *patternFlowLacpduPartnerPortNumber
}

type marshalPatternFlowLacpduPartnerPortNumber interface {
	// ToProto marshals PatternFlowLacpduPartnerPortNumber to protobuf object *otg.PatternFlowLacpduPartnerPortNumber
	ToProto() (*otg.PatternFlowLacpduPartnerPortNumber, error)
	// ToPbText marshals PatternFlowLacpduPartnerPortNumber to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerPortNumber to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerPortNumber to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerPortNumber struct {
	obj *patternFlowLacpduPartnerPortNumber
}

type unMarshalPatternFlowLacpduPartnerPortNumber interface {
	// FromProto unmarshals PatternFlowLacpduPartnerPortNumber from protobuf object *otg.PatternFlowLacpduPartnerPortNumber
	FromProto(msg *otg.PatternFlowLacpduPartnerPortNumber) (PatternFlowLacpduPartnerPortNumber, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerPortNumber from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerPortNumber from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerPortNumber from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerPortNumber) Marshal() marshalPatternFlowLacpduPartnerPortNumber {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerPortNumber{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerPortNumber) Unmarshal() unMarshalPatternFlowLacpduPartnerPortNumber {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerPortNumber{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerPortNumber) ToProto() (*otg.PatternFlowLacpduPartnerPortNumber, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerPortNumber) FromProto(msg *otg.PatternFlowLacpduPartnerPortNumber) (PatternFlowLacpduPartnerPortNumber, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerPortNumber) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerPortNumber) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerPortNumber) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerPortNumber) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerPortNumber) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerPortNumber) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerPortNumber) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerPortNumber) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerPortNumber) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerPortNumber) Clone() (PatternFlowLacpduPartnerPortNumber, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerPortNumber()
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

func (obj *patternFlowLacpduPartnerPortNumber) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerPortNumber is the port number of the Partner's port, as received by the Actor.
type PatternFlowLacpduPartnerPortNumber interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerPortNumber to protobuf object *otg.PatternFlowLacpduPartnerPortNumber
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerPortNumber
	// setMsg unmarshals PatternFlowLacpduPartnerPortNumber from protobuf object *otg.PatternFlowLacpduPartnerPortNumber
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerPortNumber) PatternFlowLacpduPartnerPortNumber
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerPortNumber
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerPortNumber
	// validate validates PatternFlowLacpduPartnerPortNumber
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerPortNumber, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerPortNumberChoiceEnum, set in PatternFlowLacpduPartnerPortNumber
	Choice() PatternFlowLacpduPartnerPortNumberChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerPortNumberChoiceEnum provided by user to PatternFlowLacpduPartnerPortNumber
	setChoice(value PatternFlowLacpduPartnerPortNumberChoiceEnum) PatternFlowLacpduPartnerPortNumber
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerPortNumber
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerPortNumber.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerPortNumber
	SetValue(value uint32) PatternFlowLacpduPartnerPortNumber
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerPortNumber
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerPortNumber.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerPortNumber
	SetValues(value []uint32) PatternFlowLacpduPartnerPortNumber
	// Increment returns PatternFlowLacpduPartnerPortNumberCounter, set in PatternFlowLacpduPartnerPortNumber.
	// PatternFlowLacpduPartnerPortNumberCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerPortNumberCounter
	// SetIncrement assigns PatternFlowLacpduPartnerPortNumberCounter provided by user to PatternFlowLacpduPartnerPortNumber.
	// PatternFlowLacpduPartnerPortNumberCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerPortNumberCounter) PatternFlowLacpduPartnerPortNumber
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerPortNumber
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerPortNumberCounter, set in PatternFlowLacpduPartnerPortNumber.
	// PatternFlowLacpduPartnerPortNumberCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerPortNumberCounter
	// SetDecrement assigns PatternFlowLacpduPartnerPortNumberCounter provided by user to PatternFlowLacpduPartnerPortNumber.
	// PatternFlowLacpduPartnerPortNumberCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerPortNumberCounter) PatternFlowLacpduPartnerPortNumber
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerPortNumber
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerPortNumberChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerPortNumber
var PatternFlowLacpduPartnerPortNumberChoice = struct {
	VALUE     PatternFlowLacpduPartnerPortNumberChoiceEnum
	VALUES    PatternFlowLacpduPartnerPortNumberChoiceEnum
	INCREMENT PatternFlowLacpduPartnerPortNumberChoiceEnum
	DECREMENT PatternFlowLacpduPartnerPortNumberChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerPortNumberChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerPortNumberChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerPortNumberChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerPortNumberChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerPortNumber) Choice() PatternFlowLacpduPartnerPortNumberChoiceEnum {
	return PatternFlowLacpduPartnerPortNumberChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerPortNumber) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerPortNumber) setChoice(value PatternFlowLacpduPartnerPortNumberChoiceEnum) PatternFlowLacpduPartnerPortNumber {
	intValue, ok := otg.PatternFlowLacpduPartnerPortNumber_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerPortNumberChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerPortNumber_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerPortNumberChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerPortNumberChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerPortNumberChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerPortNumberCounter().msg()
	}

	if value == PatternFlowLacpduPartnerPortNumberChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerPortNumberCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerPortNumber) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerPortNumberChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerPortNumber) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerPortNumber object
func (obj *patternFlowLacpduPartnerPortNumber) SetValue(value uint32) PatternFlowLacpduPartnerPortNumber {
	obj.setChoice(PatternFlowLacpduPartnerPortNumberChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerPortNumber) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerPortNumber object
func (obj *patternFlowLacpduPartnerPortNumber) SetValues(value []uint32) PatternFlowLacpduPartnerPortNumber {
	obj.setChoice(PatternFlowLacpduPartnerPortNumberChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerPortNumberCounter
func (obj *patternFlowLacpduPartnerPortNumber) Increment() PatternFlowLacpduPartnerPortNumberCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerPortNumberChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerPortNumberCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerPortNumberCounter
func (obj *patternFlowLacpduPartnerPortNumber) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerPortNumberCounter value in the PatternFlowLacpduPartnerPortNumber object
func (obj *patternFlowLacpduPartnerPortNumber) SetIncrement(value PatternFlowLacpduPartnerPortNumberCounter) PatternFlowLacpduPartnerPortNumber {
	obj.setChoice(PatternFlowLacpduPartnerPortNumberChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerPortNumberCounter
func (obj *patternFlowLacpduPartnerPortNumber) Decrement() PatternFlowLacpduPartnerPortNumberCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerPortNumberChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerPortNumberCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerPortNumberCounter
func (obj *patternFlowLacpduPartnerPortNumber) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerPortNumberCounter value in the PatternFlowLacpduPartnerPortNumber object
func (obj *patternFlowLacpduPartnerPortNumber) SetDecrement(value PatternFlowLacpduPartnerPortNumberCounter) PatternFlowLacpduPartnerPortNumber {
	obj.setChoice(PatternFlowLacpduPartnerPortNumberChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerPortNumber) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerPortNumber.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerPortNumber.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerPortNumber) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerPortNumberChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerPortNumberChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerPortNumberChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerPortNumberChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerPortNumberChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerPortNumberChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerPortNumber")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerPortNumber_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerPortNumber_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
