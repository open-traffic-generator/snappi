package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvEgressNodeValue *****
type patternFlowIpv6SegmentRoutingTlvEgressNodeValue struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	marshaller      marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	incrementHolder PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	decrementHolder PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvEgressNodeValue() PatternFlowIpv6SegmentRoutingTlvEgressNodeValue {
	obj := patternFlowIpv6SegmentRoutingTlvEgressNodeValue{obj: &otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) msg() *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue) PatternFlowIpv6SegmentRoutingTlvEgressNodeValue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValue struct {
	obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue
}

type marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValue interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValue to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValue to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValue struct {
	obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValue interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValue from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue) (PatternFlowIpv6SegmentRoutingTlvEgressNodeValue, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValue from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValue) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValue) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue) (PatternFlowIpv6SegmentRoutingTlvEgressNodeValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValue) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValue) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValue) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValue) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValue) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValue) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) Clone() (PatternFlowIpv6SegmentRoutingTlvEgressNodeValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvEgressNodeValue()
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

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingTlvEgressNodeValue is the 128-bit IPv6 address of the egress node.
type PatternFlowIpv6SegmentRoutingTlvEgressNodeValue interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValue to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValue from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue) PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	// validate validates PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvEgressNodeValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	Choice() PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	setChoice(value PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum) PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValue.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	SetValue(value string) PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValue.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	SetValues(value []string) PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	// Increment returns PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValue.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter is ipv6 counter pattern
	Increment() PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeValue.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter is ipv6 counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValue.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter is ipv6 counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeValue.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter is ipv6 counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
var PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) Choice() PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum {
	return PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) setChoice(value PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum) PatternFlowIpv6SegmentRoutingTlvEgressNodeValue {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.VALUE {
		defaultValue := "::"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.VALUES {
		defaultValue := []string{"::"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeValue object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) SetValue(value string) PatternFlowIpv6SegmentRoutingTlvEgressNodeValue {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeValue object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) SetValues(value []string) PatternFlowIpv6SegmentRoutingTlvEgressNodeValue {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) Increment() PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeValue object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) SetIncrement(value PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeValue {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) Decrement() PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeValue object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) SetDecrement(value PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeValue {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingTlvEgressNodeValue.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingTlvEgressNodeValue.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValue) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeValueChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingTlvEgressNodeValue")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValue_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
