package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved *****
type patternFlowIpv6SegmentRoutingTlvIngressNodeReserved struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	marshaller      marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	incrementHolder PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	decrementHolder PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvIngressNodeReserved() PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved {
	obj := patternFlowIpv6SegmentRoutingTlvIngressNodeReserved{obj: &otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) msg() *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReserved struct {
	obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved
}

type marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReserved interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReserved struct {
	obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReserved interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) (PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) (PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) Clone() (PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvIngressNodeReserved()
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

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved is reserved field. Must be 0.
type PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	// validate validates PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	Choice() PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	setChoice(value PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum) PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	// Increment returns PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
var PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) Choice() PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum {
	return PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) setChoice(value PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum) PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) Increment() PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) SetIncrement(value PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) Decrement() PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) SetDecrement(value PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
