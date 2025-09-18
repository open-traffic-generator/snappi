package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags *****
type patternFlowIpv6SegmentRoutingTlvEgressNodeFlags struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	marshaller      marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	incrementHolder PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	decrementHolder PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvEgressNodeFlags() PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags {
	obj := patternFlowIpv6SegmentRoutingTlvEgressNodeFlags{obj: &otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) msg() *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlags struct {
	obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags
}

type marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlags interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlags struct {
	obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlags interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) (PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) (PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) Clone() (PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvEgressNodeFlags()
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

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags is 8-bit flags field for this TLV.
type PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	// validate validates PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	Choice() PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	setChoice(value PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	// Increment returns PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
var PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) Choice() PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum {
	return PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) setChoice(value PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) Increment() PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) SetIncrement(value PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) Decrement() PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) SetDecrement(value PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlags) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
