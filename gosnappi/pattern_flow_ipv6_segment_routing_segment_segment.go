package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingSegmentSegment *****
type patternFlowIpv6SegmentRoutingSegmentSegment struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingSegmentSegment
	marshaller      marshalPatternFlowIpv6SegmentRoutingSegmentSegment
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingSegmentSegment
	incrementHolder PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	decrementHolder PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
}

func NewPatternFlowIpv6SegmentRoutingSegmentSegment() PatternFlowIpv6SegmentRoutingSegmentSegment {
	obj := patternFlowIpv6SegmentRoutingSegmentSegment{obj: &otg.PatternFlowIpv6SegmentRoutingSegmentSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) msg() *otg.PatternFlowIpv6SegmentRoutingSegmentSegment {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingSegmentSegment) PatternFlowIpv6SegmentRoutingSegmentSegment {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingSegmentSegment struct {
	obj *patternFlowIpv6SegmentRoutingSegmentSegment
}

type marshalPatternFlowIpv6SegmentRoutingSegmentSegment interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingSegmentSegment to protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentSegment
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingSegmentSegment, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingSegmentSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingSegmentSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingSegmentSegment to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingSegmentSegment struct {
	obj *patternFlowIpv6SegmentRoutingSegmentSegment
}

type unMarshalPatternFlowIpv6SegmentRoutingSegmentSegment interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingSegmentSegment from protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentSegment
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingSegmentSegment) (PatternFlowIpv6SegmentRoutingSegmentSegment, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingSegmentSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingSegmentSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingSegmentSegment from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) Marshal() marshalPatternFlowIpv6SegmentRoutingSegmentSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingSegmentSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingSegmentSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingSegmentSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentSegment) ToProto() (*otg.PatternFlowIpv6SegmentRoutingSegmentSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentSegment) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingSegmentSegment) (PatternFlowIpv6SegmentRoutingSegmentSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentSegment) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentSegment) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentSegment) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentSegment) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentSegment) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentSegment) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) Clone() (PatternFlowIpv6SegmentRoutingSegmentSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingSegmentSegment()
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

func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingSegmentSegment is a 128-bit IPv6 address segment.
type PatternFlowIpv6SegmentRoutingSegmentSegment interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingSegmentSegment to protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentSegment
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingSegmentSegment
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingSegmentSegment from protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentSegment
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingSegmentSegment) PatternFlowIpv6SegmentRoutingSegmentSegment
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingSegmentSegment
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingSegmentSegment
	// validate validates PatternFlowIpv6SegmentRoutingSegmentSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingSegmentSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum, set in PatternFlowIpv6SegmentRoutingSegmentSegment
	Choice() PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingSegmentSegment
	setChoice(value PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum) PatternFlowIpv6SegmentRoutingSegmentSegment
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingSegmentSegment
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv6SegmentRoutingSegmentSegment.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv6SegmentRoutingSegmentSegment
	SetValue(value string) PatternFlowIpv6SegmentRoutingSegmentSegment
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingSegmentSegment
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv6SegmentRoutingSegmentSegment.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv6SegmentRoutingSegmentSegment
	SetValues(value []string) PatternFlowIpv6SegmentRoutingSegmentSegment
	// Increment returns PatternFlowIpv6SegmentRoutingSegmentSegmentCounter, set in PatternFlowIpv6SegmentRoutingSegmentSegment.
	// PatternFlowIpv6SegmentRoutingSegmentSegmentCounter is ipv6 counter pattern
	Increment() PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingSegmentSegmentCounter provided by user to PatternFlowIpv6SegmentRoutingSegmentSegment.
	// PatternFlowIpv6SegmentRoutingSegmentSegmentCounter is ipv6 counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingSegmentSegmentCounter) PatternFlowIpv6SegmentRoutingSegmentSegment
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingSegmentSegment
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingSegmentSegmentCounter, set in PatternFlowIpv6SegmentRoutingSegmentSegment.
	// PatternFlowIpv6SegmentRoutingSegmentSegmentCounter is ipv6 counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingSegmentSegmentCounter provided by user to PatternFlowIpv6SegmentRoutingSegmentSegment.
	// PatternFlowIpv6SegmentRoutingSegmentSegmentCounter is ipv6 counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingSegmentSegmentCounter) PatternFlowIpv6SegmentRoutingSegmentSegment
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingSegmentSegment
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingSegmentSegment
var PatternFlowIpv6SegmentRoutingSegmentSegmentChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) Choice() PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum {
	return PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) setChoice(value PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum) PatternFlowIpv6SegmentRoutingSegmentSegment {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingSegmentSegment_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingSegmentSegment_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.VALUE {
		defaultValue := "::0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.VALUES {
		defaultValue := []string{"::0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingSegmentSegmentCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingSegmentSegmentCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv6SegmentRoutingSegmentSegment object
func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) SetValue(value string) PatternFlowIpv6SegmentRoutingSegmentSegment {
	obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv6SegmentRoutingSegmentSegment object
func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) SetValues(value []string) PatternFlowIpv6SegmentRoutingSegmentSegment {
	obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) Increment() PatternFlowIpv6SegmentRoutingSegmentSegmentCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingSegmentSegmentCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingSegmentSegmentCounter value in the PatternFlowIpv6SegmentRoutingSegmentSegment object
func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) SetIncrement(value PatternFlowIpv6SegmentRoutingSegmentSegmentCounter) PatternFlowIpv6SegmentRoutingSegmentSegment {
	obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) Decrement() PatternFlowIpv6SegmentRoutingSegmentSegmentCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingSegmentSegmentCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingSegmentSegmentCounter value in the PatternFlowIpv6SegmentRoutingSegmentSegment object
func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) SetDecrement(value PatternFlowIpv6SegmentRoutingSegmentSegmentCounter) PatternFlowIpv6SegmentRoutingSegmentSegment {
	obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingSegmentSegment.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingSegmentSegment.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegment) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingSegmentSegmentChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentSegmentChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingSegmentSegment")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingSegmentSegment_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingSegmentSegment_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
