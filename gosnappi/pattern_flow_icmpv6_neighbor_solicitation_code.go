package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NeighborSolicitationCode *****
type patternFlowIcmpv6NeighborSolicitationCode struct {
	validation
	obj             *otg.PatternFlowIcmpv6NeighborSolicitationCode
	marshaller      marshalPatternFlowIcmpv6NeighborSolicitationCode
	unMarshaller    unMarshalPatternFlowIcmpv6NeighborSolicitationCode
	incrementHolder PatternFlowIcmpv6NeighborSolicitationCodeCounter
	decrementHolder PatternFlowIcmpv6NeighborSolicitationCodeCounter
}

func NewPatternFlowIcmpv6NeighborSolicitationCode() PatternFlowIcmpv6NeighborSolicitationCode {
	obj := patternFlowIcmpv6NeighborSolicitationCode{obj: &otg.PatternFlowIcmpv6NeighborSolicitationCode{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationCode) msg() *otg.PatternFlowIcmpv6NeighborSolicitationCode {
	return obj.obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationCode) setMsg(msg *otg.PatternFlowIcmpv6NeighborSolicitationCode) PatternFlowIcmpv6NeighborSolicitationCode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NeighborSolicitationCode struct {
	obj *patternFlowIcmpv6NeighborSolicitationCode
}

type marshalPatternFlowIcmpv6NeighborSolicitationCode interface {
	// ToProto marshals PatternFlowIcmpv6NeighborSolicitationCode to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationCode
	ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationCode, error)
	// ToPbText marshals PatternFlowIcmpv6NeighborSolicitationCode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NeighborSolicitationCode to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NeighborSolicitationCode to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NeighborSolicitationCode struct {
	obj *patternFlowIcmpv6NeighborSolicitationCode
}

type unMarshalPatternFlowIcmpv6NeighborSolicitationCode interface {
	// FromProto unmarshals PatternFlowIcmpv6NeighborSolicitationCode from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationCode
	FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationCode) (PatternFlowIcmpv6NeighborSolicitationCode, error)
	// FromPbText unmarshals PatternFlowIcmpv6NeighborSolicitationCode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NeighborSolicitationCode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NeighborSolicitationCode from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NeighborSolicitationCode) Marshal() marshalPatternFlowIcmpv6NeighborSolicitationCode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NeighborSolicitationCode{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NeighborSolicitationCode) Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationCode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NeighborSolicitationCode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationCode) ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationCode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationCode) FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationCode) (PatternFlowIcmpv6NeighborSolicitationCode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationCode) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationCode) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationCode) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationCode) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationCode) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationCode) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NeighborSolicitationCode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationCode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationCode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NeighborSolicitationCode) Clone() (PatternFlowIcmpv6NeighborSolicitationCode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NeighborSolicitationCode()
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

func (obj *patternFlowIcmpv6NeighborSolicitationCode) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NeighborSolicitationCode is iCMPv6 message code. Must be 0 for Neighbor Solicitation.
type PatternFlowIcmpv6NeighborSolicitationCode interface {
	Validation
	// msg marshals PatternFlowIcmpv6NeighborSolicitationCode to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationCode
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NeighborSolicitationCode
	// setMsg unmarshals PatternFlowIcmpv6NeighborSolicitationCode from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationCode
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NeighborSolicitationCode) PatternFlowIcmpv6NeighborSolicitationCode
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NeighborSolicitationCode
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationCode
	// validate validates PatternFlowIcmpv6NeighborSolicitationCode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NeighborSolicitationCode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum, set in PatternFlowIcmpv6NeighborSolicitationCode
	Choice() PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum provided by user to PatternFlowIcmpv6NeighborSolicitationCode
	setChoice(value PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum) PatternFlowIcmpv6NeighborSolicitationCode
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NeighborSolicitationCode
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NeighborSolicitationCode.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationCode
	SetValue(value uint32) PatternFlowIcmpv6NeighborSolicitationCode
	// HasValue checks if Value has been set in PatternFlowIcmpv6NeighborSolicitationCode
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NeighborSolicitationCode.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationCode
	SetValues(value []uint32) PatternFlowIcmpv6NeighborSolicitationCode
	// Increment returns PatternFlowIcmpv6NeighborSolicitationCodeCounter, set in PatternFlowIcmpv6NeighborSolicitationCode.
	// PatternFlowIcmpv6NeighborSolicitationCodeCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NeighborSolicitationCodeCounter
	// SetIncrement assigns PatternFlowIcmpv6NeighborSolicitationCodeCounter provided by user to PatternFlowIcmpv6NeighborSolicitationCode.
	// PatternFlowIcmpv6NeighborSolicitationCodeCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NeighborSolicitationCodeCounter) PatternFlowIcmpv6NeighborSolicitationCode
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NeighborSolicitationCode
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NeighborSolicitationCodeCounter, set in PatternFlowIcmpv6NeighborSolicitationCode.
	// PatternFlowIcmpv6NeighborSolicitationCodeCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NeighborSolicitationCodeCounter
	// SetDecrement assigns PatternFlowIcmpv6NeighborSolicitationCodeCounter provided by user to PatternFlowIcmpv6NeighborSolicitationCode.
	// PatternFlowIcmpv6NeighborSolicitationCodeCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NeighborSolicitationCodeCounter) PatternFlowIcmpv6NeighborSolicitationCode
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NeighborSolicitationCode
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NeighborSolicitationCode
var PatternFlowIcmpv6NeighborSolicitationCodeChoice = struct {
	VALUE     PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum
	VALUES    PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum
	INCREMENT PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum
	DECREMENT PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NeighborSolicitationCode) Choice() PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum {
	return PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NeighborSolicitationCode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NeighborSolicitationCode) setChoice(value PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum) PatternFlowIcmpv6NeighborSolicitationCode {
	intValue, ok := otg.PatternFlowIcmpv6NeighborSolicitationCode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NeighborSolicitationCode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NeighborSolicitationCodeChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NeighborSolicitationCodeChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NeighborSolicitationCodeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NeighborSolicitationCodeCounter().msg()
	}

	if value == PatternFlowIcmpv6NeighborSolicitationCodeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NeighborSolicitationCodeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationCode) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NeighborSolicitationCodeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationCode) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationCode object
func (obj *patternFlowIcmpv6NeighborSolicitationCode) SetValue(value uint32) PatternFlowIcmpv6NeighborSolicitationCode {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationCodeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NeighborSolicitationCode) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NeighborSolicitationCode object
func (obj *patternFlowIcmpv6NeighborSolicitationCode) SetValues(value []uint32) PatternFlowIcmpv6NeighborSolicitationCode {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationCodeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NeighborSolicitationCodeCounter
func (obj *patternFlowIcmpv6NeighborSolicitationCode) Increment() PatternFlowIcmpv6NeighborSolicitationCodeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NeighborSolicitationCodeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NeighborSolicitationCodeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NeighborSolicitationCodeCounter
func (obj *patternFlowIcmpv6NeighborSolicitationCode) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NeighborSolicitationCodeCounter value in the PatternFlowIcmpv6NeighborSolicitationCode object
func (obj *patternFlowIcmpv6NeighborSolicitationCode) SetIncrement(value PatternFlowIcmpv6NeighborSolicitationCodeCounter) PatternFlowIcmpv6NeighborSolicitationCode {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationCodeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NeighborSolicitationCodeCounter
func (obj *patternFlowIcmpv6NeighborSolicitationCode) Decrement() PatternFlowIcmpv6NeighborSolicitationCodeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NeighborSolicitationCodeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NeighborSolicitationCodeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NeighborSolicitationCodeCounter
func (obj *patternFlowIcmpv6NeighborSolicitationCode) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NeighborSolicitationCodeCounter value in the PatternFlowIcmpv6NeighborSolicitationCode object
func (obj *patternFlowIcmpv6NeighborSolicitationCode) SetDecrement(value PatternFlowIcmpv6NeighborSolicitationCodeCounter) PatternFlowIcmpv6NeighborSolicitationCode {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationCodeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationCode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NeighborSolicitationCode.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6NeighborSolicitationCode.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIcmpv6NeighborSolicitationCode) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NeighborSolicitationCodeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationCodeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationCodeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationCodeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationCodeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NeighborSolicitationCodeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NeighborSolicitationCode")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NeighborSolicitationCode_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NeighborSolicitationCode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
