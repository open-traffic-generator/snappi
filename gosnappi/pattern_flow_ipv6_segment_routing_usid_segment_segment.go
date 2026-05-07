package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingUsidSegmentSegment *****
type patternFlowIpv6SegmentRoutingUsidSegmentSegment struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	marshaller      marshalPatternFlowIpv6SegmentRoutingUsidSegmentSegment
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentSegment
	incrementHolder PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	decrementHolder PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
}

func NewPatternFlowIpv6SegmentRoutingUsidSegmentSegment() PatternFlowIpv6SegmentRoutingUsidSegmentSegment {
	obj := patternFlowIpv6SegmentRoutingUsidSegmentSegment{obj: &otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment) PatternFlowIpv6SegmentRoutingUsidSegmentSegment {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingUsidSegmentSegment struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment
}

type marshalPatternFlowIpv6SegmentRoutingUsidSegmentSegment interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingUsidSegmentSegment to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingUsidSegmentSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingUsidSegmentSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingUsidSegmentSegment to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentSegment struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment
}

type unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentSegment interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentSegment from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment) (PatternFlowIpv6SegmentRoutingUsidSegmentSegment, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentSegment from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingUsidSegmentSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentSegment) ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentSegment) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment) (PatternFlowIpv6SegmentRoutingUsidSegmentSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentSegment) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentSegment) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentSegment) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentSegment) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentSegment) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentSegment) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingUsidSegmentSegment()
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

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingUsidSegmentSegment is a pre-computed 128-bit uSID container value as an IPv6 address (RFC 9800 Section 4). Example: "fc00:0:2:1:3:1::".
type PatternFlowIpv6SegmentRoutingUsidSegmentSegment interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingUsidSegmentSegment to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentSegment from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment) PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentSegment
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentSegment
	// validate validates PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum, set in PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	Choice() PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	setChoice(value PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum) PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv6SegmentRoutingUsidSegmentSegment.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	SetValue(value string) PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv6SegmentRoutingUsidSegmentSegment.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	SetValues(value []string) PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	// Increment returns PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter, set in PatternFlowIpv6SegmentRoutingUsidSegmentSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter is ipv6 counter pattern
	Increment() PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter is ipv6 counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter, set in PatternFlowIpv6SegmentRoutingUsidSegmentSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter is ipv6 counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter is ipv6 counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingUsidSegmentSegment
var PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) Choice() PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum {
	return PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) setChoice(value PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum) PatternFlowIpv6SegmentRoutingUsidSegmentSegment {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.VALUE {
		defaultValue := "::0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.VALUES {
		defaultValue := []string{"::0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv6SegmentRoutingUsidSegmentSegment object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) SetValue(value string) PatternFlowIpv6SegmentRoutingUsidSegmentSegment {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv6SegmentRoutingUsidSegmentSegment object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) SetValues(value []string) PatternFlowIpv6SegmentRoutingUsidSegmentSegment {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) Increment() PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter value in the PatternFlowIpv6SegmentRoutingUsidSegmentSegment object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) SetIncrement(value PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) PatternFlowIpv6SegmentRoutingUsidSegmentSegment {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) Decrement() PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter value in the PatternFlowIpv6SegmentRoutingUsidSegmentSegment object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) SetDecrement(value PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) PatternFlowIpv6SegmentRoutingUsidSegmentSegment {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingUsidSegmentSegment.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingUsidSegmentSegment.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegment) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentSegmentChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingUsidSegmentSegment")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegment_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
