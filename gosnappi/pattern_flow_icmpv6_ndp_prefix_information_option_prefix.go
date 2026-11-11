package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionPrefix *****
type patternFlowIcmpv6NdpPrefixInformationOptionPrefix struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	marshaller      marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	unMarshaller    unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	incrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	decrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefix() PatternFlowIcmpv6NdpPrefixInformationOptionPrefix {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionPrefix{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix) PatternFlowIcmpv6NdpPrefixInformationOptionPrefix {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefix struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefix interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefix to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefix to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefix to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefix to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefix struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefix interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefix from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix) (PatternFlowIcmpv6NdpPrefixInformationOptionPrefix, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefix from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefix from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefix from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefix {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefix{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefix {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefix{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefix) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefix) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix) (PatternFlowIcmpv6NdpPrefixInformationOptionPrefix, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefix) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefix) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefix) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefix) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefix) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefix) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionPrefix, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefix()
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpPrefixInformationOptionPrefix is iP address or prefix of an IP address. The Prefix Length field contains the number of valid leading bits.
type PatternFlowIcmpv6NdpPrefixInformationOptionPrefix interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefix to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefix from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix) PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionPrefix, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	Choice() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	HasChoice() bool
	// Value returns string, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefix.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	SetValue(value string) PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	HasValue() bool
	// Values returns []string, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefix.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	SetValues(value []string) PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	// Increment returns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefix.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter is ipv6 counter pattern
	Increment() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefix.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter is ipv6 counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefix.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter is ipv6 counter pattern
	Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefix.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter is ipv6 counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
var PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice = struct {
	VALUE     PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum
	VALUES    PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) Choice() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum {
	return PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionPrefix {
	intValue, ok := otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.VALUE {
		defaultValue := "::"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.VALUES {
		defaultValue := []string{"::"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefix object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) SetValue(value string) PatternFlowIcmpv6NdpPrefixInformationOptionPrefix {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefix object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) SetValues(value []string) PatternFlowIcmpv6NdpPrefixInformationOptionPrefix {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) Increment() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefix object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPrefix {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefix object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPrefix {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIcmpv6NdpPrefixInformationOptionPrefix.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIcmpv6NdpPrefixInformationOptionPrefix.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefix) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPrefixChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpPrefixInformationOptionPrefix")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefix_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
