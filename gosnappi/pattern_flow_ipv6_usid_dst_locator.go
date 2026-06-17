package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6UsidDstLocator *****
type patternFlowIpv6UsidDstLocator struct {
	validation
	obj             *otg.PatternFlowIpv6UsidDstLocator
	marshaller      marshalPatternFlowIpv6UsidDstLocator
	unMarshaller    unMarshalPatternFlowIpv6UsidDstLocator
	incrementHolder PatternFlowIpv6UsidDstLocatorCounter
	decrementHolder PatternFlowIpv6UsidDstLocatorCounter
}

func NewPatternFlowIpv6UsidDstLocator() PatternFlowIpv6UsidDstLocator {
	obj := patternFlowIpv6UsidDstLocator{obj: &otg.PatternFlowIpv6UsidDstLocator{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6UsidDstLocator) msg() *otg.PatternFlowIpv6UsidDstLocator {
	return obj.obj
}

func (obj *patternFlowIpv6UsidDstLocator) setMsg(msg *otg.PatternFlowIpv6UsidDstLocator) PatternFlowIpv6UsidDstLocator {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6UsidDstLocator struct {
	obj *patternFlowIpv6UsidDstLocator
}

type marshalPatternFlowIpv6UsidDstLocator interface {
	// ToProto marshals PatternFlowIpv6UsidDstLocator to protobuf object *otg.PatternFlowIpv6UsidDstLocator
	ToProto() (*otg.PatternFlowIpv6UsidDstLocator, error)
	// ToPbText marshals PatternFlowIpv6UsidDstLocator to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6UsidDstLocator to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6UsidDstLocator to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6UsidDstLocator struct {
	obj *patternFlowIpv6UsidDstLocator
}

type unMarshalPatternFlowIpv6UsidDstLocator interface {
	// FromProto unmarshals PatternFlowIpv6UsidDstLocator from protobuf object *otg.PatternFlowIpv6UsidDstLocator
	FromProto(msg *otg.PatternFlowIpv6UsidDstLocator) (PatternFlowIpv6UsidDstLocator, error)
	// FromPbText unmarshals PatternFlowIpv6UsidDstLocator from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6UsidDstLocator from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6UsidDstLocator from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6UsidDstLocator) Marshal() marshalPatternFlowIpv6UsidDstLocator {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6UsidDstLocator{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6UsidDstLocator) Unmarshal() unMarshalPatternFlowIpv6UsidDstLocator {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6UsidDstLocator{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6UsidDstLocator) ToProto() (*otg.PatternFlowIpv6UsidDstLocator, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6UsidDstLocator) FromProto(msg *otg.PatternFlowIpv6UsidDstLocator) (PatternFlowIpv6UsidDstLocator, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6UsidDstLocator) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6UsidDstLocator) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6UsidDstLocator) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6UsidDstLocator) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6UsidDstLocator) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6UsidDstLocator) FromJson(value string) error {
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

func (obj *patternFlowIpv6UsidDstLocator) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6UsidDstLocator) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6UsidDstLocator) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6UsidDstLocator) Clone() (PatternFlowIpv6UsidDstLocator, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6UsidDstLocator()
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

func (obj *patternFlowIpv6UsidDstLocator) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6UsidDstLocator is the Locator Block (LB) IPv6 prefix shared by all uSIDs in this container (RFC 9800 Section 3). Defines the common high-order bits packed into the 128-bit dst. For F3216 this is a /32 prefix (e.g., fc00::). The number of bits used is given by locator_length.
type PatternFlowIpv6UsidDstLocator interface {
	Validation
	// msg marshals PatternFlowIpv6UsidDstLocator to protobuf object *otg.PatternFlowIpv6UsidDstLocator
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6UsidDstLocator
	// setMsg unmarshals PatternFlowIpv6UsidDstLocator from protobuf object *otg.PatternFlowIpv6UsidDstLocator
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6UsidDstLocator) PatternFlowIpv6UsidDstLocator
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6UsidDstLocator
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6UsidDstLocator
	// validate validates PatternFlowIpv6UsidDstLocator
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6UsidDstLocator, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6UsidDstLocatorChoiceEnum, set in PatternFlowIpv6UsidDstLocator
	Choice() PatternFlowIpv6UsidDstLocatorChoiceEnum
	// setChoice assigns PatternFlowIpv6UsidDstLocatorChoiceEnum provided by user to PatternFlowIpv6UsidDstLocator
	setChoice(value PatternFlowIpv6UsidDstLocatorChoiceEnum) PatternFlowIpv6UsidDstLocator
	// HasChoice checks if Choice has been set in PatternFlowIpv6UsidDstLocator
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv6UsidDstLocator.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv6UsidDstLocator
	SetValue(value string) PatternFlowIpv6UsidDstLocator
	// HasValue checks if Value has been set in PatternFlowIpv6UsidDstLocator
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv6UsidDstLocator.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv6UsidDstLocator
	SetValues(value []string) PatternFlowIpv6UsidDstLocator
	// Increment returns PatternFlowIpv6UsidDstLocatorCounter, set in PatternFlowIpv6UsidDstLocator.
	// PatternFlowIpv6UsidDstLocatorCounter is ipv6 counter pattern
	Increment() PatternFlowIpv6UsidDstLocatorCounter
	// SetIncrement assigns PatternFlowIpv6UsidDstLocatorCounter provided by user to PatternFlowIpv6UsidDstLocator.
	// PatternFlowIpv6UsidDstLocatorCounter is ipv6 counter pattern
	SetIncrement(value PatternFlowIpv6UsidDstLocatorCounter) PatternFlowIpv6UsidDstLocator
	// HasIncrement checks if Increment has been set in PatternFlowIpv6UsidDstLocator
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6UsidDstLocatorCounter, set in PatternFlowIpv6UsidDstLocator.
	// PatternFlowIpv6UsidDstLocatorCounter is ipv6 counter pattern
	Decrement() PatternFlowIpv6UsidDstLocatorCounter
	// SetDecrement assigns PatternFlowIpv6UsidDstLocatorCounter provided by user to PatternFlowIpv6UsidDstLocator.
	// PatternFlowIpv6UsidDstLocatorCounter is ipv6 counter pattern
	SetDecrement(value PatternFlowIpv6UsidDstLocatorCounter) PatternFlowIpv6UsidDstLocator
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6UsidDstLocator
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6UsidDstLocatorChoiceEnum string

// Enum of Choice on PatternFlowIpv6UsidDstLocator
var PatternFlowIpv6UsidDstLocatorChoice = struct {
	VALUE     PatternFlowIpv6UsidDstLocatorChoiceEnum
	VALUES    PatternFlowIpv6UsidDstLocatorChoiceEnum
	INCREMENT PatternFlowIpv6UsidDstLocatorChoiceEnum
	DECREMENT PatternFlowIpv6UsidDstLocatorChoiceEnum
}{
	VALUE:     PatternFlowIpv6UsidDstLocatorChoiceEnum("value"),
	VALUES:    PatternFlowIpv6UsidDstLocatorChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6UsidDstLocatorChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6UsidDstLocatorChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6UsidDstLocator) Choice() PatternFlowIpv6UsidDstLocatorChoiceEnum {
	return PatternFlowIpv6UsidDstLocatorChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6UsidDstLocator) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6UsidDstLocator) setChoice(value PatternFlowIpv6UsidDstLocatorChoiceEnum) PatternFlowIpv6UsidDstLocator {
	intValue, ok := otg.PatternFlowIpv6UsidDstLocator_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6UsidDstLocatorChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6UsidDstLocator_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6UsidDstLocatorChoice.VALUE {
		defaultValue := "::0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6UsidDstLocatorChoice.VALUES {
		defaultValue := []string{"::0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6UsidDstLocatorChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6UsidDstLocatorCounter().msg()
	}

	if value == PatternFlowIpv6UsidDstLocatorChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6UsidDstLocatorCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6UsidDstLocator) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6UsidDstLocatorChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6UsidDstLocator) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv6UsidDstLocator object
func (obj *patternFlowIpv6UsidDstLocator) SetValue(value string) PatternFlowIpv6UsidDstLocator {
	obj.setChoice(PatternFlowIpv6UsidDstLocatorChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv6UsidDstLocator) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv6UsidDstLocator object
func (obj *patternFlowIpv6UsidDstLocator) SetValues(value []string) PatternFlowIpv6UsidDstLocator {
	obj.setChoice(PatternFlowIpv6UsidDstLocatorChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6UsidDstLocatorCounter
func (obj *patternFlowIpv6UsidDstLocator) Increment() PatternFlowIpv6UsidDstLocatorCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6UsidDstLocatorChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6UsidDstLocatorCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6UsidDstLocatorCounter
func (obj *patternFlowIpv6UsidDstLocator) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6UsidDstLocatorCounter value in the PatternFlowIpv6UsidDstLocator object
func (obj *patternFlowIpv6UsidDstLocator) SetIncrement(value PatternFlowIpv6UsidDstLocatorCounter) PatternFlowIpv6UsidDstLocator {
	obj.setChoice(PatternFlowIpv6UsidDstLocatorChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6UsidDstLocatorCounter
func (obj *patternFlowIpv6UsidDstLocator) Decrement() PatternFlowIpv6UsidDstLocatorCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6UsidDstLocatorChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6UsidDstLocatorCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6UsidDstLocatorCounter
func (obj *patternFlowIpv6UsidDstLocator) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6UsidDstLocatorCounter value in the PatternFlowIpv6UsidDstLocator object
func (obj *patternFlowIpv6UsidDstLocator) SetDecrement(value PatternFlowIpv6UsidDstLocatorCounter) PatternFlowIpv6UsidDstLocator {
	obj.setChoice(PatternFlowIpv6UsidDstLocatorChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6UsidDstLocator) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6UsidDstLocator.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6UsidDstLocator.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6UsidDstLocator) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6UsidDstLocatorChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6UsidDstLocatorChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6UsidDstLocatorChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6UsidDstLocatorChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6UsidDstLocatorChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6UsidDstLocatorChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6UsidDstLocator")
			}
		} else {
			intVal := otg.PatternFlowIpv6UsidDstLocator_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6UsidDstLocator_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
