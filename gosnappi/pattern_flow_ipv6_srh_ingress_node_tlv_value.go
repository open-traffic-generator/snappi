package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHIngressNodeTlvValue *****
type patternFlowIpv6SRHIngressNodeTlvValue struct {
	validation
	obj             *otg.PatternFlowIpv6SRHIngressNodeTlvValue
	marshaller      marshalPatternFlowIpv6SRHIngressNodeTlvValue
	unMarshaller    unMarshalPatternFlowIpv6SRHIngressNodeTlvValue
	incrementHolder PatternFlowIpv6SRHIngressNodeTlvValueCounter
	decrementHolder PatternFlowIpv6SRHIngressNodeTlvValueCounter
}

func NewPatternFlowIpv6SRHIngressNodeTlvValue() PatternFlowIpv6SRHIngressNodeTlvValue {
	obj := patternFlowIpv6SRHIngressNodeTlvValue{obj: &otg.PatternFlowIpv6SRHIngressNodeTlvValue{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValue) msg() *otg.PatternFlowIpv6SRHIngressNodeTlvValue {
	return obj.obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValue) setMsg(msg *otg.PatternFlowIpv6SRHIngressNodeTlvValue) PatternFlowIpv6SRHIngressNodeTlvValue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHIngressNodeTlvValue struct {
	obj *patternFlowIpv6SRHIngressNodeTlvValue
}

type marshalPatternFlowIpv6SRHIngressNodeTlvValue interface {
	// ToProto marshals PatternFlowIpv6SRHIngressNodeTlvValue to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvValue
	ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvValue, error)
	// ToPbText marshals PatternFlowIpv6SRHIngressNodeTlvValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHIngressNodeTlvValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHIngressNodeTlvValue to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHIngressNodeTlvValue struct {
	obj *patternFlowIpv6SRHIngressNodeTlvValue
}

type unMarshalPatternFlowIpv6SRHIngressNodeTlvValue interface {
	// FromProto unmarshals PatternFlowIpv6SRHIngressNodeTlvValue from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvValue
	FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvValue) (PatternFlowIpv6SRHIngressNodeTlvValue, error)
	// FromPbText unmarshals PatternFlowIpv6SRHIngressNodeTlvValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHIngressNodeTlvValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHIngressNodeTlvValue from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValue) Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHIngressNodeTlvValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValue) Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHIngressNodeTlvValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvValue) ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvValue) FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvValue) (PatternFlowIpv6SRHIngressNodeTlvValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvValue) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvValue) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvValue) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvValue) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvValue) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvValue) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHIngressNodeTlvValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValue) Clone() (PatternFlowIpv6SRHIngressNodeTlvValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHIngressNodeTlvValue()
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

func (obj *patternFlowIpv6SRHIngressNodeTlvValue) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHIngressNodeTlvValue is global IPv6 address of the SRv6 ingress node (RFC 9259 Section 3.1). This is the address of the node that imposed the SRH on the packet.
type PatternFlowIpv6SRHIngressNodeTlvValue interface {
	Validation
	// msg marshals PatternFlowIpv6SRHIngressNodeTlvValue to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvValue
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHIngressNodeTlvValue
	// setMsg unmarshals PatternFlowIpv6SRHIngressNodeTlvValue from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvValue
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHIngressNodeTlvValue) PatternFlowIpv6SRHIngressNodeTlvValue
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvValue
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvValue
	// validate validates PatternFlowIpv6SRHIngressNodeTlvValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHIngressNodeTlvValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum, set in PatternFlowIpv6SRHIngressNodeTlvValue
	Choice() PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum provided by user to PatternFlowIpv6SRHIngressNodeTlvValue
	setChoice(value PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum) PatternFlowIpv6SRHIngressNodeTlvValue
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHIngressNodeTlvValue
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv6SRHIngressNodeTlvValue.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv6SRHIngressNodeTlvValue
	SetValue(value string) PatternFlowIpv6SRHIngressNodeTlvValue
	// HasValue checks if Value has been set in PatternFlowIpv6SRHIngressNodeTlvValue
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv6SRHIngressNodeTlvValue.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv6SRHIngressNodeTlvValue
	SetValues(value []string) PatternFlowIpv6SRHIngressNodeTlvValue
	// Increment returns PatternFlowIpv6SRHIngressNodeTlvValueCounter, set in PatternFlowIpv6SRHIngressNodeTlvValue.
	// PatternFlowIpv6SRHIngressNodeTlvValueCounter is ipv6 counter pattern
	Increment() PatternFlowIpv6SRHIngressNodeTlvValueCounter
	// SetIncrement assigns PatternFlowIpv6SRHIngressNodeTlvValueCounter provided by user to PatternFlowIpv6SRHIngressNodeTlvValue.
	// PatternFlowIpv6SRHIngressNodeTlvValueCounter is ipv6 counter pattern
	SetIncrement(value PatternFlowIpv6SRHIngressNodeTlvValueCounter) PatternFlowIpv6SRHIngressNodeTlvValue
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHIngressNodeTlvValue
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHIngressNodeTlvValueCounter, set in PatternFlowIpv6SRHIngressNodeTlvValue.
	// PatternFlowIpv6SRHIngressNodeTlvValueCounter is ipv6 counter pattern
	Decrement() PatternFlowIpv6SRHIngressNodeTlvValueCounter
	// SetDecrement assigns PatternFlowIpv6SRHIngressNodeTlvValueCounter provided by user to PatternFlowIpv6SRHIngressNodeTlvValue.
	// PatternFlowIpv6SRHIngressNodeTlvValueCounter is ipv6 counter pattern
	SetDecrement(value PatternFlowIpv6SRHIngressNodeTlvValueCounter) PatternFlowIpv6SRHIngressNodeTlvValue
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHIngressNodeTlvValue
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHIngressNodeTlvValue
var PatternFlowIpv6SRHIngressNodeTlvValueChoice = struct {
	VALUE     PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum
	VALUES    PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum
	INCREMENT PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum
	DECREMENT PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValue) Choice() PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum {
	return PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHIngressNodeTlvValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValue) setChoice(value PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum) PatternFlowIpv6SRHIngressNodeTlvValue {
	intValue, ok := otg.PatternFlowIpv6SRHIngressNodeTlvValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHIngressNodeTlvValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHIngressNodeTlvValueChoice.VALUE {
		defaultValue := "::0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvValueChoice.VALUES {
		defaultValue := []string{"::0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvValueChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHIngressNodeTlvValueCounter().msg()
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvValueChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHIngressNodeTlvValueCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SRHIngressNodeTlvValue) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvValueChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SRHIngressNodeTlvValue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv6SRHIngressNodeTlvValue object
func (obj *patternFlowIpv6SRHIngressNodeTlvValue) SetValue(value string) PatternFlowIpv6SRHIngressNodeTlvValue {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvValueChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv6SRHIngressNodeTlvValue) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv6SRHIngressNodeTlvValue object
func (obj *patternFlowIpv6SRHIngressNodeTlvValue) SetValues(value []string) PatternFlowIpv6SRHIngressNodeTlvValue {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvValueChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHIngressNodeTlvValueCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvValue) Increment() PatternFlowIpv6SRHIngressNodeTlvValueCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvValueChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHIngressNodeTlvValueCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHIngressNodeTlvValueCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvValue) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHIngressNodeTlvValueCounter value in the PatternFlowIpv6SRHIngressNodeTlvValue object
func (obj *patternFlowIpv6SRHIngressNodeTlvValue) SetIncrement(value PatternFlowIpv6SRHIngressNodeTlvValueCounter) PatternFlowIpv6SRHIngressNodeTlvValue {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvValueChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHIngressNodeTlvValueCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvValue) Decrement() PatternFlowIpv6SRHIngressNodeTlvValueCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvValueChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHIngressNodeTlvValueCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHIngressNodeTlvValueCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvValue) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHIngressNodeTlvValueCounter value in the PatternFlowIpv6SRHIngressNodeTlvValue object
func (obj *patternFlowIpv6SRHIngressNodeTlvValue) SetDecrement(value PatternFlowIpv6SRHIngressNodeTlvValueCounter) PatternFlowIpv6SRHIngressNodeTlvValue {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvValueChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SRHIngressNodeTlvValue.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SRHIngressNodeTlvValue.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHIngressNodeTlvValue) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHIngressNodeTlvValueChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvValueChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvValueChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvValueChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvValueChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvValueChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHIngressNodeTlvValue")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHIngressNodeTlvValue_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHIngressNodeTlvValue_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
