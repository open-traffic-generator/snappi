package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingUsidSegmentLocator *****
type patternFlowIpv6SegmentRoutingUsidSegmentLocator struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	marshaller      marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocator
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocator
	incrementHolder PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	decrementHolder PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
}

func NewPatternFlowIpv6SegmentRoutingUsidSegmentLocator() PatternFlowIpv6SegmentRoutingUsidSegmentLocator {
	obj := patternFlowIpv6SegmentRoutingUsidSegmentLocator{obj: &otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator) PatternFlowIpv6SegmentRoutingUsidSegmentLocator {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocator struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator
}

type marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocator interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocator to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocator to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocator to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocator to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocator struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator
}

type unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocator interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocator from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator) (PatternFlowIpv6SegmentRoutingUsidSegmentLocator, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocator from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocator from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocator from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocator {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocator{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocator {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocator{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocator) ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocator) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator) (PatternFlowIpv6SegmentRoutingUsidSegmentLocator, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocator) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocator) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocator) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocator) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocator) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocator) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentLocator, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingUsidSegmentLocator()
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

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingUsidSegmentLocator is the Locator Block (LB) IPv6 prefix shared by all uSIDs in this container (RFC 9800 Section 3). Defines the common high-order bits of every uSID assembled from this entry. For F3216, this is a /32 prefix (e.g., fc00::). The prefix length is given by locator_length.
type PatternFlowIpv6SegmentRoutingUsidSegmentLocator interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocator to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocator from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator) PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocator
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocator
	// validate validates PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentLocator, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	Choice() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	setChoice(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum) PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocator.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	SetValue(value string) PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocator.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	SetValues(value []string) PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	// Increment returns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocator.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter is ipv6 counter pattern
	Increment() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocator.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter is ipv6 counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocator.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter is ipv6 counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocator.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter is ipv6 counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingUsidSegmentLocator
var PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) Choice() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum {
	return PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) setChoice(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum) PatternFlowIpv6SegmentRoutingUsidSegmentLocator {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.VALUE {
		defaultValue := "::0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.VALUES {
		defaultValue := []string{"::0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocator object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) SetValue(value string) PatternFlowIpv6SegmentRoutingUsidSegmentLocator {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocator object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) SetValues(value []string) PatternFlowIpv6SegmentRoutingUsidSegmentLocator {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) Increment() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocator object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) SetIncrement(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) PatternFlowIpv6SegmentRoutingUsidSegmentLocator {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) Decrement() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocator object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) SetDecrement(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) PatternFlowIpv6SegmentRoutingUsidSegmentLocator {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingUsidSegmentLocator.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingUsidSegmentLocator.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocator) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentLocatorChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingUsidSegmentLocator")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocator_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
