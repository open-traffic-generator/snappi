package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvIngressNodeValue *****
type patternFlowIpv6SegmentRoutingTlvIngressNodeValue struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	marshaller      marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	incrementHolder PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	decrementHolder PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvIngressNodeValue() PatternFlowIpv6SegmentRoutingTlvIngressNodeValue {
	obj := patternFlowIpv6SegmentRoutingTlvIngressNodeValue{obj: &otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) msg() *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue) PatternFlowIpv6SegmentRoutingTlvIngressNodeValue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValue struct {
	obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue
}

type marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValue interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValue to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValue to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValue struct {
	obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValue interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValue from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue) (PatternFlowIpv6SegmentRoutingTlvIngressNodeValue, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValue from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValue) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValue) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue) (PatternFlowIpv6SegmentRoutingTlvIngressNodeValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValue) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValue) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValue) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValue) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValue) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValue) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) Clone() (PatternFlowIpv6SegmentRoutingTlvIngressNodeValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvIngressNodeValue()
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

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingTlvIngressNodeValue is the 128-bit IPv6 address of the ingress node.
type PatternFlowIpv6SegmentRoutingTlvIngressNodeValue interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValue to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValue from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue) PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	// validate validates PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvIngressNodeValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	Choice() PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	setChoice(value PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum) PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValue.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	SetValue(value string) PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValue.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	SetValues(value []string) PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	// Increment returns PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValue.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter is ipv6 counter pattern
	Increment() PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeValue.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter is ipv6 counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValue.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter is ipv6 counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeValue.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter is ipv6 counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
var PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) Choice() PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum {
	return PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) setChoice(value PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum) PatternFlowIpv6SegmentRoutingTlvIngressNodeValue {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.VALUE {
		defaultValue := "::"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.VALUES {
		defaultValue := []string{"::"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeValue object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) SetValue(value string) PatternFlowIpv6SegmentRoutingTlvIngressNodeValue {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeValue object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) SetValues(value []string) PatternFlowIpv6SegmentRoutingTlvIngressNodeValue {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) Increment() PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeValue object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) SetIncrement(value PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeValue {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) Decrement() PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeValue object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) SetDecrement(value PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeValue {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingTlvIngressNodeValue.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingTlvIngressNodeValue.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValue) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeValueChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingTlvIngressNodeValue")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValue_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
