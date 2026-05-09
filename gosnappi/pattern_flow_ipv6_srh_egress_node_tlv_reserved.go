package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHEgressNodeTlvReserved *****
type patternFlowIpv6SRHEgressNodeTlvReserved struct {
	validation
	obj             *otg.PatternFlowIpv6SRHEgressNodeTlvReserved
	marshaller      marshalPatternFlowIpv6SRHEgressNodeTlvReserved
	unMarshaller    unMarshalPatternFlowIpv6SRHEgressNodeTlvReserved
	incrementHolder PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	decrementHolder PatternFlowIpv6SRHEgressNodeTlvReservedCounter
}

func NewPatternFlowIpv6SRHEgressNodeTlvReserved() PatternFlowIpv6SRHEgressNodeTlvReserved {
	obj := patternFlowIpv6SRHEgressNodeTlvReserved{obj: &otg.PatternFlowIpv6SRHEgressNodeTlvReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) msg() *otg.PatternFlowIpv6SRHEgressNodeTlvReserved {
	return obj.obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) setMsg(msg *otg.PatternFlowIpv6SRHEgressNodeTlvReserved) PatternFlowIpv6SRHEgressNodeTlvReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHEgressNodeTlvReserved struct {
	obj *patternFlowIpv6SRHEgressNodeTlvReserved
}

type marshalPatternFlowIpv6SRHEgressNodeTlvReserved interface {
	// ToProto marshals PatternFlowIpv6SRHEgressNodeTlvReserved to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvReserved
	ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvReserved, error)
	// ToPbText marshals PatternFlowIpv6SRHEgressNodeTlvReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHEgressNodeTlvReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHEgressNodeTlvReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHEgressNodeTlvReserved struct {
	obj *patternFlowIpv6SRHEgressNodeTlvReserved
}

type unMarshalPatternFlowIpv6SRHEgressNodeTlvReserved interface {
	// FromProto unmarshals PatternFlowIpv6SRHEgressNodeTlvReserved from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvReserved
	FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvReserved) (PatternFlowIpv6SRHEgressNodeTlvReserved, error)
	// FromPbText unmarshals PatternFlowIpv6SRHEgressNodeTlvReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHEgressNodeTlvReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHEgressNodeTlvReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHEgressNodeTlvReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHEgressNodeTlvReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvReserved) ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvReserved) FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvReserved) (PatternFlowIpv6SRHEgressNodeTlvReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvReserved) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) Clone() (PatternFlowIpv6SRHEgressNodeTlvReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHEgressNodeTlvReserved()
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

func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHEgressNodeTlvReserved is reserved byte following the Length field (RFC 9259 Section 3.2). Must be zero on transmission. The adjacent Node ID Length byte is fixed at 16 for IPv6 addresses and is not separately configurable.
type PatternFlowIpv6SRHEgressNodeTlvReserved interface {
	Validation
	// msg marshals PatternFlowIpv6SRHEgressNodeTlvReserved to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHEgressNodeTlvReserved
	// setMsg unmarshals PatternFlowIpv6SRHEgressNodeTlvReserved from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHEgressNodeTlvReserved) PatternFlowIpv6SRHEgressNodeTlvReserved
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvReserved
	// validate validates PatternFlowIpv6SRHEgressNodeTlvReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHEgressNodeTlvReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum, set in PatternFlowIpv6SRHEgressNodeTlvReserved
	Choice() PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum provided by user to PatternFlowIpv6SRHEgressNodeTlvReserved
	setChoice(value PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum) PatternFlowIpv6SRHEgressNodeTlvReserved
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHEgressNodeTlvReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvReserved
	SetValue(value uint32) PatternFlowIpv6SRHEgressNodeTlvReserved
	// HasValue checks if Value has been set in PatternFlowIpv6SRHEgressNodeTlvReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHEgressNodeTlvReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvReserved
	SetValues(value []uint32) PatternFlowIpv6SRHEgressNodeTlvReserved
	// Increment returns PatternFlowIpv6SRHEgressNodeTlvReservedCounter, set in PatternFlowIpv6SRHEgressNodeTlvReserved.
	// PatternFlowIpv6SRHEgressNodeTlvReservedCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	// SetIncrement assigns PatternFlowIpv6SRHEgressNodeTlvReservedCounter provided by user to PatternFlowIpv6SRHEgressNodeTlvReserved.
	// PatternFlowIpv6SRHEgressNodeTlvReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHEgressNodeTlvReservedCounter) PatternFlowIpv6SRHEgressNodeTlvReserved
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHEgressNodeTlvReserved
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHEgressNodeTlvReservedCounter, set in PatternFlowIpv6SRHEgressNodeTlvReserved.
	// PatternFlowIpv6SRHEgressNodeTlvReservedCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	// SetDecrement assigns PatternFlowIpv6SRHEgressNodeTlvReservedCounter provided by user to PatternFlowIpv6SRHEgressNodeTlvReserved.
	// PatternFlowIpv6SRHEgressNodeTlvReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHEgressNodeTlvReservedCounter) PatternFlowIpv6SRHEgressNodeTlvReserved
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHEgressNodeTlvReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHEgressNodeTlvReserved
var PatternFlowIpv6SRHEgressNodeTlvReservedChoice = struct {
	VALUE     PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum
	VALUES    PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum
	INCREMENT PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum
	DECREMENT PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) Choice() PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum {
	return PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) setChoice(value PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum) PatternFlowIpv6SRHEgressNodeTlvReserved {
	intValue, ok := otg.PatternFlowIpv6SRHEgressNodeTlvReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHEgressNodeTlvReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHEgressNodeTlvReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHEgressNodeTlvReservedCounter().msg()
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHEgressNodeTlvReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHEgressNodeTlvReserved object
func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) SetValue(value uint32) PatternFlowIpv6SRHEgressNodeTlvReserved {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHEgressNodeTlvReserved object
func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) SetValues(value []uint32) PatternFlowIpv6SRHEgressNodeTlvReserved {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHEgressNodeTlvReservedCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) Increment() PatternFlowIpv6SRHEgressNodeTlvReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHEgressNodeTlvReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHEgressNodeTlvReservedCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHEgressNodeTlvReservedCounter value in the PatternFlowIpv6SRHEgressNodeTlvReserved object
func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) SetIncrement(value PatternFlowIpv6SRHEgressNodeTlvReservedCounter) PatternFlowIpv6SRHEgressNodeTlvReserved {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHEgressNodeTlvReservedCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) Decrement() PatternFlowIpv6SRHEgressNodeTlvReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHEgressNodeTlvReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHEgressNodeTlvReservedCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHEgressNodeTlvReservedCounter value in the PatternFlowIpv6SRHEgressNodeTlvReserved object
func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) SetDecrement(value PatternFlowIpv6SRHEgressNodeTlvReservedCounter) PatternFlowIpv6SRHEgressNodeTlvReserved {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvReserved.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SRHEgressNodeTlvReserved.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIpv6SRHEgressNodeTlvReserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHEgressNodeTlvReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHEgressNodeTlvReserved")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHEgressNodeTlvReserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHEgressNodeTlvReserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
