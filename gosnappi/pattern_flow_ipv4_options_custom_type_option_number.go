package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsCustomTypeOptionNumber *****
type patternFlowIpv4OptionsCustomTypeOptionNumber struct {
	validation
	obj             *otg.PatternFlowIpv4OptionsCustomTypeOptionNumber
	marshaller      marshalPatternFlowIpv4OptionsCustomTypeOptionNumber
	unMarshaller    unMarshalPatternFlowIpv4OptionsCustomTypeOptionNumber
	incrementHolder PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	decrementHolder PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
}

func NewPatternFlowIpv4OptionsCustomTypeOptionNumber() PatternFlowIpv4OptionsCustomTypeOptionNumber {
	obj := patternFlowIpv4OptionsCustomTypeOptionNumber{obj: &otg.PatternFlowIpv4OptionsCustomTypeOptionNumber{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) msg() *otg.PatternFlowIpv4OptionsCustomTypeOptionNumber {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) setMsg(msg *otg.PatternFlowIpv4OptionsCustomTypeOptionNumber) PatternFlowIpv4OptionsCustomTypeOptionNumber {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsCustomTypeOptionNumber struct {
	obj *patternFlowIpv4OptionsCustomTypeOptionNumber
}

type marshalPatternFlowIpv4OptionsCustomTypeOptionNumber interface {
	// ToProto marshals PatternFlowIpv4OptionsCustomTypeOptionNumber to protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionNumber
	ToProto() (*otg.PatternFlowIpv4OptionsCustomTypeOptionNumber, error)
	// ToPbText marshals PatternFlowIpv4OptionsCustomTypeOptionNumber to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsCustomTypeOptionNumber to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsCustomTypeOptionNumber to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsCustomTypeOptionNumber struct {
	obj *patternFlowIpv4OptionsCustomTypeOptionNumber
}

type unMarshalPatternFlowIpv4OptionsCustomTypeOptionNumber interface {
	// FromProto unmarshals PatternFlowIpv4OptionsCustomTypeOptionNumber from protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionNumber
	FromProto(msg *otg.PatternFlowIpv4OptionsCustomTypeOptionNumber) (PatternFlowIpv4OptionsCustomTypeOptionNumber, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsCustomTypeOptionNumber from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsCustomTypeOptionNumber from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsCustomTypeOptionNumber from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) Marshal() marshalPatternFlowIpv4OptionsCustomTypeOptionNumber {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsCustomTypeOptionNumber{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) Unmarshal() unMarshalPatternFlowIpv4OptionsCustomTypeOptionNumber {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsCustomTypeOptionNumber{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionNumber) ToProto() (*otg.PatternFlowIpv4OptionsCustomTypeOptionNumber, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionNumber) FromProto(msg *otg.PatternFlowIpv4OptionsCustomTypeOptionNumber) (PatternFlowIpv4OptionsCustomTypeOptionNumber, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionNumber) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionNumber) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionNumber) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionNumber) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionNumber) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionNumber) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) Clone() (PatternFlowIpv4OptionsCustomTypeOptionNumber, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsCustomTypeOptionNumber()
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

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4OptionsCustomTypeOptionNumber is option Number [Ref:https://www.iana.org/assignments/ip-parameters/ip-parameters.xhtml#ip-parameters-1].
type PatternFlowIpv4OptionsCustomTypeOptionNumber interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsCustomTypeOptionNumber to protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionNumber
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsCustomTypeOptionNumber
	// setMsg unmarshals PatternFlowIpv4OptionsCustomTypeOptionNumber from protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionNumber
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsCustomTypeOptionNumber) PatternFlowIpv4OptionsCustomTypeOptionNumber
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsCustomTypeOptionNumber
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsCustomTypeOptionNumber
	// validate validates PatternFlowIpv4OptionsCustomTypeOptionNumber
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsCustomTypeOptionNumber, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum, set in PatternFlowIpv4OptionsCustomTypeOptionNumber
	Choice() PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum
	// setChoice assigns PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum provided by user to PatternFlowIpv4OptionsCustomTypeOptionNumber
	setChoice(value PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum) PatternFlowIpv4OptionsCustomTypeOptionNumber
	// HasChoice checks if Choice has been set in PatternFlowIpv4OptionsCustomTypeOptionNumber
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4OptionsCustomTypeOptionNumber.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4OptionsCustomTypeOptionNumber
	SetValue(value uint32) PatternFlowIpv4OptionsCustomTypeOptionNumber
	// HasValue checks if Value has been set in PatternFlowIpv4OptionsCustomTypeOptionNumber
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4OptionsCustomTypeOptionNumber.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4OptionsCustomTypeOptionNumber
	SetValues(value []uint32) PatternFlowIpv4OptionsCustomTypeOptionNumber
	// Increment returns PatternFlowIpv4OptionsCustomTypeOptionNumberCounter, set in PatternFlowIpv4OptionsCustomTypeOptionNumber.
	// PatternFlowIpv4OptionsCustomTypeOptionNumberCounter is integer counter pattern
	Increment() PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	// SetIncrement assigns PatternFlowIpv4OptionsCustomTypeOptionNumberCounter provided by user to PatternFlowIpv4OptionsCustomTypeOptionNumber.
	// PatternFlowIpv4OptionsCustomTypeOptionNumberCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4OptionsCustomTypeOptionNumberCounter) PatternFlowIpv4OptionsCustomTypeOptionNumber
	// HasIncrement checks if Increment has been set in PatternFlowIpv4OptionsCustomTypeOptionNumber
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4OptionsCustomTypeOptionNumberCounter, set in PatternFlowIpv4OptionsCustomTypeOptionNumber.
	// PatternFlowIpv4OptionsCustomTypeOptionNumberCounter is integer counter pattern
	Decrement() PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	// SetDecrement assigns PatternFlowIpv4OptionsCustomTypeOptionNumberCounter provided by user to PatternFlowIpv4OptionsCustomTypeOptionNumber.
	// PatternFlowIpv4OptionsCustomTypeOptionNumberCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4OptionsCustomTypeOptionNumberCounter) PatternFlowIpv4OptionsCustomTypeOptionNumber
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4OptionsCustomTypeOptionNumber
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum string

// Enum of Choice on PatternFlowIpv4OptionsCustomTypeOptionNumber
var PatternFlowIpv4OptionsCustomTypeOptionNumberChoice = struct {
	VALUE     PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum
	VALUES    PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum
	INCREMENT PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum
	DECREMENT PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum
}{
	VALUE:     PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum("value"),
	VALUES:    PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) Choice() PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum {
	return PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) setChoice(value PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum) PatternFlowIpv4OptionsCustomTypeOptionNumber {
	intValue, ok := otg.PatternFlowIpv4OptionsCustomTypeOptionNumber_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4OptionsCustomTypeOptionNumber_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4OptionsCustomTypeOptionNumberCounter().msg()
	}

	if value == PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4OptionsCustomTypeOptionNumberCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4OptionsCustomTypeOptionNumber object
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) SetValue(value uint32) PatternFlowIpv4OptionsCustomTypeOptionNumber {
	obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4OptionsCustomTypeOptionNumber object
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) SetValues(value []uint32) PatternFlowIpv4OptionsCustomTypeOptionNumber {
	obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) Increment() PatternFlowIpv4OptionsCustomTypeOptionNumberCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4OptionsCustomTypeOptionNumberCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4OptionsCustomTypeOptionNumberCounter value in the PatternFlowIpv4OptionsCustomTypeOptionNumber object
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) SetIncrement(value PatternFlowIpv4OptionsCustomTypeOptionNumberCounter) PatternFlowIpv4OptionsCustomTypeOptionNumber {
	obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) Decrement() PatternFlowIpv4OptionsCustomTypeOptionNumberCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4OptionsCustomTypeOptionNumberCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4OptionsCustomTypeOptionNumberCounter value in the PatternFlowIpv4OptionsCustomTypeOptionNumber object
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) SetDecrement(value PatternFlowIpv4OptionsCustomTypeOptionNumberCounter) PatternFlowIpv4OptionsCustomTypeOptionNumber {
	obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsCustomTypeOptionNumber.Value <= 31 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 31 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4OptionsCustomTypeOptionNumber.Values <= 31 but Got %d", item))
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

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumber) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4OptionsCustomTypeOptionNumberChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionNumberChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4OptionsCustomTypeOptionNumber")
			}
		} else {
			intVal := otg.PatternFlowIpv4OptionsCustomTypeOptionNumber_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4OptionsCustomTypeOptionNumber_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
