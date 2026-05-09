package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsU1Flag *****
type patternFlowIpv6SegmentRoutingFlagsU1Flag struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag
	marshaller      marshalPatternFlowIpv6SegmentRoutingFlagsU1Flag
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingFlagsU1Flag
	incrementHolder PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter
	decrementHolder PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsU1Flag() PatternFlowIpv6SegmentRoutingFlagsU1Flag {
	obj := patternFlowIpv6SegmentRoutingFlagsU1Flag{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag) PatternFlowIpv6SegmentRoutingFlagsU1Flag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsU1Flag struct {
	obj *patternFlowIpv6SegmentRoutingFlagsU1Flag
}

type marshalPatternFlowIpv6SegmentRoutingFlagsU1Flag interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsU1Flag to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsU1Flag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsU1Flag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsU1Flag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsU1Flag struct {
	obj *patternFlowIpv6SegmentRoutingFlagsU1Flag
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsU1Flag interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsU1Flag from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag) (PatternFlowIpv6SegmentRoutingFlagsU1Flag, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsU1Flag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsU1Flag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsU1Flag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsU1Flag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsU1Flag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsU1Flag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsU1Flag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsU1Flag) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsU1Flag) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag) (PatternFlowIpv6SegmentRoutingFlagsU1Flag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsU1Flag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsU1Flag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsU1Flag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsU1Flag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsU1Flag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsU1Flag) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) Clone() (PatternFlowIpv6SegmentRoutingFlagsU1Flag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsU1Flag()
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

func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingFlagsU1Flag is reserved bit U1 (RFC 8754 Section 2.1). MUST be zero on normal transmission. Exposed for negative or conformance testing only.
type PatternFlowIpv6SegmentRoutingFlagsU1Flag interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsU1Flag to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsU1Flag from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag) PatternFlowIpv6SegmentRoutingFlagsU1Flag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsU1Flag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsU1Flag
	// validate validates PatternFlowIpv6SegmentRoutingFlagsU1Flag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsU1Flag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum, set in PatternFlowIpv6SegmentRoutingFlagsU1Flag
	Choice() PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingFlagsU1Flag
	setChoice(value PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsU1Flag
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingFlagsU1Flag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsU1Flag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsU1Flag
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsU1Flag
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingFlagsU1Flag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingFlagsU1Flag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsU1Flag
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsU1Flag
	// Increment returns PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter, set in PatternFlowIpv6SegmentRoutingFlagsU1Flag.
	// PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsU1Flag.
	// PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter) PatternFlowIpv6SegmentRoutingFlagsU1Flag
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingFlagsU1Flag
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter, set in PatternFlowIpv6SegmentRoutingFlagsU1Flag.
	// PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsU1Flag.
	// PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter) PatternFlowIpv6SegmentRoutingFlagsU1Flag
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingFlagsU1Flag
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingFlagsU1Flag
var PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) Choice() PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum {
	return PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) setChoice(value PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsU1Flag {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingFlagsU1FlagCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingFlagsU1FlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsU1Flag object
func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsU1Flag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingFlagsU1Flag object
func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsU1Flag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) Increment() PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingFlagsU1FlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter value in the PatternFlowIpv6SegmentRoutingFlagsU1Flag object
func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter) PatternFlowIpv6SegmentRoutingFlagsU1Flag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) Decrement() PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingFlagsU1FlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter value in the PatternFlowIpv6SegmentRoutingFlagsU1Flag object
func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsU1FlagCounter) PatternFlowIpv6SegmentRoutingFlagsU1Flag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsU1Flag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingFlagsU1Flag.Values <= 1 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingFlagsU1Flag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingFlagsU1FlagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsU1FlagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingFlagsU1Flag")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsU1Flag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
