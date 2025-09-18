package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags *****
type patternFlowIpv6SegmentRoutingTlvIngressNodeFlags struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	marshaller      marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	incrementHolder PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	decrementHolder PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvIngressNodeFlags() PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags {
	obj := patternFlowIpv6SegmentRoutingTlvIngressNodeFlags{obj: &otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) msg() *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlags struct {
	obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags
}

type marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlags interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlags struct {
	obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlags interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) (PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) (PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) Clone() (PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvIngressNodeFlags()
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

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags is 8-bit flags field for this TLV.
type PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	// validate validates PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	Choice() PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	setChoice(value PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	// Increment returns PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
var PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) Choice() PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum {
	return PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) setChoice(value PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) Increment() PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) SetIncrement(value PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) Decrement() PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) SetDecrement(value PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlags) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
