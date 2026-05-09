package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHEgressNodeTlvValue *****
type patternFlowIpv6SRHEgressNodeTlvValue struct {
	validation
	obj             *otg.PatternFlowIpv6SRHEgressNodeTlvValue
	marshaller      marshalPatternFlowIpv6SRHEgressNodeTlvValue
	unMarshaller    unMarshalPatternFlowIpv6SRHEgressNodeTlvValue
	incrementHolder PatternFlowIpv6SRHEgressNodeTlvValueCounter
	decrementHolder PatternFlowIpv6SRHEgressNodeTlvValueCounter
}

func NewPatternFlowIpv6SRHEgressNodeTlvValue() PatternFlowIpv6SRHEgressNodeTlvValue {
	obj := patternFlowIpv6SRHEgressNodeTlvValue{obj: &otg.PatternFlowIpv6SRHEgressNodeTlvValue{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValue) msg() *otg.PatternFlowIpv6SRHEgressNodeTlvValue {
	return obj.obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValue) setMsg(msg *otg.PatternFlowIpv6SRHEgressNodeTlvValue) PatternFlowIpv6SRHEgressNodeTlvValue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHEgressNodeTlvValue struct {
	obj *patternFlowIpv6SRHEgressNodeTlvValue
}

type marshalPatternFlowIpv6SRHEgressNodeTlvValue interface {
	// ToProto marshals PatternFlowIpv6SRHEgressNodeTlvValue to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvValue
	ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvValue, error)
	// ToPbText marshals PatternFlowIpv6SRHEgressNodeTlvValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHEgressNodeTlvValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHEgressNodeTlvValue to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHEgressNodeTlvValue struct {
	obj *patternFlowIpv6SRHEgressNodeTlvValue
}

type unMarshalPatternFlowIpv6SRHEgressNodeTlvValue interface {
	// FromProto unmarshals PatternFlowIpv6SRHEgressNodeTlvValue from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvValue
	FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvValue) (PatternFlowIpv6SRHEgressNodeTlvValue, error)
	// FromPbText unmarshals PatternFlowIpv6SRHEgressNodeTlvValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHEgressNodeTlvValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHEgressNodeTlvValue from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValue) Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHEgressNodeTlvValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValue) Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHEgressNodeTlvValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvValue) ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvValue) FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvValue) (PatternFlowIpv6SRHEgressNodeTlvValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvValue) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvValue) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvValue) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvValue) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvValue) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvValue) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHEgressNodeTlvValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValue) Clone() (PatternFlowIpv6SRHEgressNodeTlvValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHEgressNodeTlvValue()
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

func (obj *patternFlowIpv6SRHEgressNodeTlvValue) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHEgressNodeTlvValue is global IPv6 address of the intended SRv6 egress node (RFC 9259 Section 3.2). This is the address of the final segment endpoint of the SR policy.
type PatternFlowIpv6SRHEgressNodeTlvValue interface {
	Validation
	// msg marshals PatternFlowIpv6SRHEgressNodeTlvValue to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvValue
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHEgressNodeTlvValue
	// setMsg unmarshals PatternFlowIpv6SRHEgressNodeTlvValue from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvValue
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHEgressNodeTlvValue) PatternFlowIpv6SRHEgressNodeTlvValue
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvValue
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvValue
	// validate validates PatternFlowIpv6SRHEgressNodeTlvValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHEgressNodeTlvValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum, set in PatternFlowIpv6SRHEgressNodeTlvValue
	Choice() PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum provided by user to PatternFlowIpv6SRHEgressNodeTlvValue
	setChoice(value PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum) PatternFlowIpv6SRHEgressNodeTlvValue
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHEgressNodeTlvValue
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv6SRHEgressNodeTlvValue.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv6SRHEgressNodeTlvValue
	SetValue(value string) PatternFlowIpv6SRHEgressNodeTlvValue
	// HasValue checks if Value has been set in PatternFlowIpv6SRHEgressNodeTlvValue
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv6SRHEgressNodeTlvValue.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv6SRHEgressNodeTlvValue
	SetValues(value []string) PatternFlowIpv6SRHEgressNodeTlvValue
	// Increment returns PatternFlowIpv6SRHEgressNodeTlvValueCounter, set in PatternFlowIpv6SRHEgressNodeTlvValue.
	// PatternFlowIpv6SRHEgressNodeTlvValueCounter is ipv6 counter pattern
	Increment() PatternFlowIpv6SRHEgressNodeTlvValueCounter
	// SetIncrement assigns PatternFlowIpv6SRHEgressNodeTlvValueCounter provided by user to PatternFlowIpv6SRHEgressNodeTlvValue.
	// PatternFlowIpv6SRHEgressNodeTlvValueCounter is ipv6 counter pattern
	SetIncrement(value PatternFlowIpv6SRHEgressNodeTlvValueCounter) PatternFlowIpv6SRHEgressNodeTlvValue
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHEgressNodeTlvValue
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHEgressNodeTlvValueCounter, set in PatternFlowIpv6SRHEgressNodeTlvValue.
	// PatternFlowIpv6SRHEgressNodeTlvValueCounter is ipv6 counter pattern
	Decrement() PatternFlowIpv6SRHEgressNodeTlvValueCounter
	// SetDecrement assigns PatternFlowIpv6SRHEgressNodeTlvValueCounter provided by user to PatternFlowIpv6SRHEgressNodeTlvValue.
	// PatternFlowIpv6SRHEgressNodeTlvValueCounter is ipv6 counter pattern
	SetDecrement(value PatternFlowIpv6SRHEgressNodeTlvValueCounter) PatternFlowIpv6SRHEgressNodeTlvValue
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHEgressNodeTlvValue
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHEgressNodeTlvValue
var PatternFlowIpv6SRHEgressNodeTlvValueChoice = struct {
	VALUE     PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum
	VALUES    PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum
	INCREMENT PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum
	DECREMENT PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValue) Choice() PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum {
	return PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHEgressNodeTlvValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValue) setChoice(value PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum) PatternFlowIpv6SRHEgressNodeTlvValue {
	intValue, ok := otg.PatternFlowIpv6SRHEgressNodeTlvValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHEgressNodeTlvValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHEgressNodeTlvValueChoice.VALUE {
		defaultValue := "::0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvValueChoice.VALUES {
		defaultValue := []string{"::0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvValueChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHEgressNodeTlvValueCounter().msg()
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvValueChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHEgressNodeTlvValueCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SRHEgressNodeTlvValue) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvValueChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SRHEgressNodeTlvValue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv6SRHEgressNodeTlvValue object
func (obj *patternFlowIpv6SRHEgressNodeTlvValue) SetValue(value string) PatternFlowIpv6SRHEgressNodeTlvValue {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvValueChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv6SRHEgressNodeTlvValue) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv6SRHEgressNodeTlvValue object
func (obj *patternFlowIpv6SRHEgressNodeTlvValue) SetValues(value []string) PatternFlowIpv6SRHEgressNodeTlvValue {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvValueChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHEgressNodeTlvValueCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvValue) Increment() PatternFlowIpv6SRHEgressNodeTlvValueCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvValueChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHEgressNodeTlvValueCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHEgressNodeTlvValueCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvValue) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHEgressNodeTlvValueCounter value in the PatternFlowIpv6SRHEgressNodeTlvValue object
func (obj *patternFlowIpv6SRHEgressNodeTlvValue) SetIncrement(value PatternFlowIpv6SRHEgressNodeTlvValueCounter) PatternFlowIpv6SRHEgressNodeTlvValue {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvValueChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHEgressNodeTlvValueCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvValue) Decrement() PatternFlowIpv6SRHEgressNodeTlvValueCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvValueChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHEgressNodeTlvValueCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHEgressNodeTlvValueCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvValue) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHEgressNodeTlvValueCounter value in the PatternFlowIpv6SRHEgressNodeTlvValue object
func (obj *patternFlowIpv6SRHEgressNodeTlvValue) SetDecrement(value PatternFlowIpv6SRHEgressNodeTlvValueCounter) PatternFlowIpv6SRHEgressNodeTlvValue {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvValueChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SRHEgressNodeTlvValue.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SRHEgressNodeTlvValue.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHEgressNodeTlvValue) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHEgressNodeTlvValueChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvValueChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvValueChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvValueChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvValueChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvValueChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHEgressNodeTlvValue")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHEgressNodeTlvValue_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHEgressNodeTlvValue_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
