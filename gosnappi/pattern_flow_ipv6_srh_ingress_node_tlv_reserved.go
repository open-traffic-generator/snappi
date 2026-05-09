package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHIngressNodeTlvReserved *****
type patternFlowIpv6SRHIngressNodeTlvReserved struct {
	validation
	obj             *otg.PatternFlowIpv6SRHIngressNodeTlvReserved
	marshaller      marshalPatternFlowIpv6SRHIngressNodeTlvReserved
	unMarshaller    unMarshalPatternFlowIpv6SRHIngressNodeTlvReserved
	incrementHolder PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	decrementHolder PatternFlowIpv6SRHIngressNodeTlvReservedCounter
}

func NewPatternFlowIpv6SRHIngressNodeTlvReserved() PatternFlowIpv6SRHIngressNodeTlvReserved {
	obj := patternFlowIpv6SRHIngressNodeTlvReserved{obj: &otg.PatternFlowIpv6SRHIngressNodeTlvReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) msg() *otg.PatternFlowIpv6SRHIngressNodeTlvReserved {
	return obj.obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) setMsg(msg *otg.PatternFlowIpv6SRHIngressNodeTlvReserved) PatternFlowIpv6SRHIngressNodeTlvReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHIngressNodeTlvReserved struct {
	obj *patternFlowIpv6SRHIngressNodeTlvReserved
}

type marshalPatternFlowIpv6SRHIngressNodeTlvReserved interface {
	// ToProto marshals PatternFlowIpv6SRHIngressNodeTlvReserved to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvReserved
	ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvReserved, error)
	// ToPbText marshals PatternFlowIpv6SRHIngressNodeTlvReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHIngressNodeTlvReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHIngressNodeTlvReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHIngressNodeTlvReserved struct {
	obj *patternFlowIpv6SRHIngressNodeTlvReserved
}

type unMarshalPatternFlowIpv6SRHIngressNodeTlvReserved interface {
	// FromProto unmarshals PatternFlowIpv6SRHIngressNodeTlvReserved from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvReserved
	FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvReserved) (PatternFlowIpv6SRHIngressNodeTlvReserved, error)
	// FromPbText unmarshals PatternFlowIpv6SRHIngressNodeTlvReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHIngressNodeTlvReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHIngressNodeTlvReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHIngressNodeTlvReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHIngressNodeTlvReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvReserved) ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvReserved) FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvReserved) (PatternFlowIpv6SRHIngressNodeTlvReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvReserved) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) Clone() (PatternFlowIpv6SRHIngressNodeTlvReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHIngressNodeTlvReserved()
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

func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHIngressNodeTlvReserved is reserved byte following the Length field (RFC 9259 Section 3.1). Must be zero on transmission. The adjacent Node ID Length byte is fixed at 16 for IPv6 addresses and is not separately configurable.
type PatternFlowIpv6SRHIngressNodeTlvReserved interface {
	Validation
	// msg marshals PatternFlowIpv6SRHIngressNodeTlvReserved to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHIngressNodeTlvReserved
	// setMsg unmarshals PatternFlowIpv6SRHIngressNodeTlvReserved from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHIngressNodeTlvReserved) PatternFlowIpv6SRHIngressNodeTlvReserved
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvReserved
	// validate validates PatternFlowIpv6SRHIngressNodeTlvReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHIngressNodeTlvReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum, set in PatternFlowIpv6SRHIngressNodeTlvReserved
	Choice() PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum provided by user to PatternFlowIpv6SRHIngressNodeTlvReserved
	setChoice(value PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum) PatternFlowIpv6SRHIngressNodeTlvReserved
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHIngressNodeTlvReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvReserved
	SetValue(value uint32) PatternFlowIpv6SRHIngressNodeTlvReserved
	// HasValue checks if Value has been set in PatternFlowIpv6SRHIngressNodeTlvReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHIngressNodeTlvReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvReserved
	SetValues(value []uint32) PatternFlowIpv6SRHIngressNodeTlvReserved
	// Increment returns PatternFlowIpv6SRHIngressNodeTlvReservedCounter, set in PatternFlowIpv6SRHIngressNodeTlvReserved.
	// PatternFlowIpv6SRHIngressNodeTlvReservedCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	// SetIncrement assigns PatternFlowIpv6SRHIngressNodeTlvReservedCounter provided by user to PatternFlowIpv6SRHIngressNodeTlvReserved.
	// PatternFlowIpv6SRHIngressNodeTlvReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHIngressNodeTlvReservedCounter) PatternFlowIpv6SRHIngressNodeTlvReserved
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHIngressNodeTlvReserved
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHIngressNodeTlvReservedCounter, set in PatternFlowIpv6SRHIngressNodeTlvReserved.
	// PatternFlowIpv6SRHIngressNodeTlvReservedCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	// SetDecrement assigns PatternFlowIpv6SRHIngressNodeTlvReservedCounter provided by user to PatternFlowIpv6SRHIngressNodeTlvReserved.
	// PatternFlowIpv6SRHIngressNodeTlvReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHIngressNodeTlvReservedCounter) PatternFlowIpv6SRHIngressNodeTlvReserved
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHIngressNodeTlvReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHIngressNodeTlvReserved
var PatternFlowIpv6SRHIngressNodeTlvReservedChoice = struct {
	VALUE     PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum
	VALUES    PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum
	INCREMENT PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum
	DECREMENT PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) Choice() PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum {
	return PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) setChoice(value PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum) PatternFlowIpv6SRHIngressNodeTlvReserved {
	intValue, ok := otg.PatternFlowIpv6SRHIngressNodeTlvReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHIngressNodeTlvReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHIngressNodeTlvReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHIngressNodeTlvReservedCounter().msg()
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHIngressNodeTlvReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHIngressNodeTlvReserved object
func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) SetValue(value uint32) PatternFlowIpv6SRHIngressNodeTlvReserved {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHIngressNodeTlvReserved object
func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) SetValues(value []uint32) PatternFlowIpv6SRHIngressNodeTlvReserved {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHIngressNodeTlvReservedCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) Increment() PatternFlowIpv6SRHIngressNodeTlvReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHIngressNodeTlvReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHIngressNodeTlvReservedCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHIngressNodeTlvReservedCounter value in the PatternFlowIpv6SRHIngressNodeTlvReserved object
func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) SetIncrement(value PatternFlowIpv6SRHIngressNodeTlvReservedCounter) PatternFlowIpv6SRHIngressNodeTlvReserved {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHIngressNodeTlvReservedCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) Decrement() PatternFlowIpv6SRHIngressNodeTlvReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHIngressNodeTlvReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHIngressNodeTlvReservedCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHIngressNodeTlvReservedCounter value in the PatternFlowIpv6SRHIngressNodeTlvReserved object
func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) SetDecrement(value PatternFlowIpv6SRHIngressNodeTlvReservedCounter) PatternFlowIpv6SRHIngressNodeTlvReserved {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvReserved.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SRHIngressNodeTlvReserved.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIpv6SRHIngressNodeTlvReserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHIngressNodeTlvReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHIngressNodeTlvReserved")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHIngressNodeTlvReserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHIngressNodeTlvReserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
