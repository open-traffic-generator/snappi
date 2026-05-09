package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsU2Flag *****
type patternFlowIpv6SegmentRoutingFlagsU2Flag struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag
	marshaller      marshalPatternFlowIpv6SegmentRoutingFlagsU2Flag
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingFlagsU2Flag
	incrementHolder PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	decrementHolder PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsU2Flag() PatternFlowIpv6SegmentRoutingFlagsU2Flag {
	obj := patternFlowIpv6SegmentRoutingFlagsU2Flag{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag) PatternFlowIpv6SegmentRoutingFlagsU2Flag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsU2Flag struct {
	obj *patternFlowIpv6SegmentRoutingFlagsU2Flag
}

type marshalPatternFlowIpv6SegmentRoutingFlagsU2Flag interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsU2Flag to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsU2Flag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsU2Flag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsU2Flag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsU2Flag struct {
	obj *patternFlowIpv6SegmentRoutingFlagsU2Flag
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsU2Flag interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsU2Flag from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag) (PatternFlowIpv6SegmentRoutingFlagsU2Flag, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsU2Flag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsU2Flag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsU2Flag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsU2Flag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsU2Flag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsU2Flag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsU2Flag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsU2Flag) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsU2Flag) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag) (PatternFlowIpv6SegmentRoutingFlagsU2Flag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsU2Flag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsU2Flag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsU2Flag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsU2Flag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsU2Flag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsU2Flag) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) Clone() (PatternFlowIpv6SegmentRoutingFlagsU2Flag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsU2Flag()
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

func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingFlagsU2Flag is reserved bit U2 (RFC 8754 Section 2.1). MUST be zero on normal transmission. Exposed for negative or conformance testing only.
type PatternFlowIpv6SegmentRoutingFlagsU2Flag interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsU2Flag to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsU2Flag from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag) PatternFlowIpv6SegmentRoutingFlagsU2Flag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsU2Flag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsU2Flag
	// validate validates PatternFlowIpv6SegmentRoutingFlagsU2Flag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsU2Flag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum, set in PatternFlowIpv6SegmentRoutingFlagsU2Flag
	Choice() PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingFlagsU2Flag
	setChoice(value PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsU2Flag
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingFlagsU2Flag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsU2Flag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsU2Flag
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsU2Flag
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingFlagsU2Flag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingFlagsU2Flag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsU2Flag
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsU2Flag
	// Increment returns PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter, set in PatternFlowIpv6SegmentRoutingFlagsU2Flag.
	// PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsU2Flag.
	// PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) PatternFlowIpv6SegmentRoutingFlagsU2Flag
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingFlagsU2Flag
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter, set in PatternFlowIpv6SegmentRoutingFlagsU2Flag.
	// PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsU2Flag.
	// PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) PatternFlowIpv6SegmentRoutingFlagsU2Flag
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingFlagsU2Flag
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingFlagsU2Flag
var PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) Choice() PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum {
	return PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) setChoice(value PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsU2Flag {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingFlagsU2FlagCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingFlagsU2FlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsU2Flag object
func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsU2Flag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingFlagsU2Flag object
func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsU2Flag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) Increment() PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingFlagsU2FlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter value in the PatternFlowIpv6SegmentRoutingFlagsU2Flag object
func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) PatternFlowIpv6SegmentRoutingFlagsU2Flag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) Decrement() PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingFlagsU2FlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter value in the PatternFlowIpv6SegmentRoutingFlagsU2Flag object
func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) PatternFlowIpv6SegmentRoutingFlagsU2Flag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsU2Flag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingFlagsU2Flag.Values <= 1 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingFlagsU2Flag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingFlagsU2FlagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU2FlagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingFlagsU2Flag")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsU2Flag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
