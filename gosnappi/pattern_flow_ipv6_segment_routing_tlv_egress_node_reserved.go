package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved *****
type patternFlowIpv6SegmentRoutingTlvEgressNodeReserved struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	marshaller      marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	incrementHolder PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	decrementHolder PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvEgressNodeReserved() PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved {
	obj := patternFlowIpv6SegmentRoutingTlvEgressNodeReserved{obj: &otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) msg() *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReserved struct {
	obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved
}

type marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReserved interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReserved struct {
	obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReserved interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) (PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) (PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) Clone() (PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvEgressNodeReserved()
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

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved is reserved field. Must be 0.
type PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	// validate validates PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	Choice() PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	setChoice(value PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum) PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	// Increment returns PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
var PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) Choice() PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum {
	return PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) setChoice(value PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum) PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) Increment() PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) SetIncrement(value PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) Decrement() PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) SetDecrement(value PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
