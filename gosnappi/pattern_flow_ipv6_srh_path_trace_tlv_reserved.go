package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPathTraceTlvReserved *****
type patternFlowIpv6SRHPathTraceTlvReserved struct {
	validation
	obj             *otg.PatternFlowIpv6SRHPathTraceTlvReserved
	marshaller      marshalPatternFlowIpv6SRHPathTraceTlvReserved
	unMarshaller    unMarshalPatternFlowIpv6SRHPathTraceTlvReserved
	incrementHolder PatternFlowIpv6SRHPathTraceTlvReservedCounter
	decrementHolder PatternFlowIpv6SRHPathTraceTlvReservedCounter
}

func NewPatternFlowIpv6SRHPathTraceTlvReserved() PatternFlowIpv6SRHPathTraceTlvReserved {
	obj := patternFlowIpv6SRHPathTraceTlvReserved{obj: &otg.PatternFlowIpv6SRHPathTraceTlvReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvReserved) msg() *otg.PatternFlowIpv6SRHPathTraceTlvReserved {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvReserved) setMsg(msg *otg.PatternFlowIpv6SRHPathTraceTlvReserved) PatternFlowIpv6SRHPathTraceTlvReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPathTraceTlvReserved struct {
	obj *patternFlowIpv6SRHPathTraceTlvReserved
}

type marshalPatternFlowIpv6SRHPathTraceTlvReserved interface {
	// ToProto marshals PatternFlowIpv6SRHPathTraceTlvReserved to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvReserved
	ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvReserved, error)
	// ToPbText marshals PatternFlowIpv6SRHPathTraceTlvReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPathTraceTlvReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPathTraceTlvReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPathTraceTlvReserved struct {
	obj *patternFlowIpv6SRHPathTraceTlvReserved
}

type unMarshalPatternFlowIpv6SRHPathTraceTlvReserved interface {
	// FromProto unmarshals PatternFlowIpv6SRHPathTraceTlvReserved from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvReserved
	FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvReserved) (PatternFlowIpv6SRHPathTraceTlvReserved, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPathTraceTlvReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPathTraceTlvReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPathTraceTlvReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPathTraceTlvReserved) Marshal() marshalPatternFlowIpv6SRHPathTraceTlvReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPathTraceTlvReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPathTraceTlvReserved) Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPathTraceTlvReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvReserved) ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvReserved) FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvReserved) (PatternFlowIpv6SRHPathTraceTlvReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvReserved) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPathTraceTlvReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPathTraceTlvReserved) Clone() (PatternFlowIpv6SRHPathTraceTlvReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPathTraceTlvReserved()
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

func (obj *patternFlowIpv6SRHPathTraceTlvReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHPathTraceTlvReserved is reserved field (draft-ietf-spring-srv6-path-tracing). Must be zero on transmission.
type PatternFlowIpv6SRHPathTraceTlvReserved interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPathTraceTlvReserved to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPathTraceTlvReserved
	// setMsg unmarshals PatternFlowIpv6SRHPathTraceTlvReserved from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPathTraceTlvReserved) PatternFlowIpv6SRHPathTraceTlvReserved
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPathTraceTlvReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvReserved
	// validate validates PatternFlowIpv6SRHPathTraceTlvReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPathTraceTlvReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum, set in PatternFlowIpv6SRHPathTraceTlvReserved
	Choice() PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum provided by user to PatternFlowIpv6SRHPathTraceTlvReserved
	setChoice(value PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum) PatternFlowIpv6SRHPathTraceTlvReserved
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHPathTraceTlvReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHPathTraceTlvReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvReserved
	SetValue(value uint32) PatternFlowIpv6SRHPathTraceTlvReserved
	// HasValue checks if Value has been set in PatternFlowIpv6SRHPathTraceTlvReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHPathTraceTlvReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvReserved
	SetValues(value []uint32) PatternFlowIpv6SRHPathTraceTlvReserved
	// Increment returns PatternFlowIpv6SRHPathTraceTlvReservedCounter, set in PatternFlowIpv6SRHPathTraceTlvReserved.
	// PatternFlowIpv6SRHPathTraceTlvReservedCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHPathTraceTlvReservedCounter
	// SetIncrement assigns PatternFlowIpv6SRHPathTraceTlvReservedCounter provided by user to PatternFlowIpv6SRHPathTraceTlvReserved.
	// PatternFlowIpv6SRHPathTraceTlvReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHPathTraceTlvReservedCounter) PatternFlowIpv6SRHPathTraceTlvReserved
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHPathTraceTlvReserved
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHPathTraceTlvReservedCounter, set in PatternFlowIpv6SRHPathTraceTlvReserved.
	// PatternFlowIpv6SRHPathTraceTlvReservedCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHPathTraceTlvReservedCounter
	// SetDecrement assigns PatternFlowIpv6SRHPathTraceTlvReservedCounter provided by user to PatternFlowIpv6SRHPathTraceTlvReserved.
	// PatternFlowIpv6SRHPathTraceTlvReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHPathTraceTlvReservedCounter) PatternFlowIpv6SRHPathTraceTlvReserved
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHPathTraceTlvReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHPathTraceTlvReserved
var PatternFlowIpv6SRHPathTraceTlvReservedChoice = struct {
	VALUE     PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum
	VALUES    PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum
	INCREMENT PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum
	DECREMENT PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHPathTraceTlvReserved) Choice() PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum {
	return PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHPathTraceTlvReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHPathTraceTlvReserved) setChoice(value PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum) PatternFlowIpv6SRHPathTraceTlvReserved {
	intValue, ok := otg.PatternFlowIpv6SRHPathTraceTlvReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHPathTraceTlvReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHPathTraceTlvReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHPathTraceTlvReservedCounter().msg()
	}

	if value == PatternFlowIpv6SRHPathTraceTlvReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHPathTraceTlvReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvReserved object
func (obj *patternFlowIpv6SRHPathTraceTlvReserved) SetValue(value uint32) PatternFlowIpv6SRHPathTraceTlvReserved {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHPathTraceTlvReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHPathTraceTlvReserved object
func (obj *patternFlowIpv6SRHPathTraceTlvReserved) SetValues(value []uint32) PatternFlowIpv6SRHPathTraceTlvReserved {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPathTraceTlvReservedCounter
func (obj *patternFlowIpv6SRHPathTraceTlvReserved) Increment() PatternFlowIpv6SRHPathTraceTlvReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHPathTraceTlvReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPathTraceTlvReservedCounter
func (obj *patternFlowIpv6SRHPathTraceTlvReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHPathTraceTlvReservedCounter value in the PatternFlowIpv6SRHPathTraceTlvReserved object
func (obj *patternFlowIpv6SRHPathTraceTlvReserved) SetIncrement(value PatternFlowIpv6SRHPathTraceTlvReservedCounter) PatternFlowIpv6SRHPathTraceTlvReserved {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPathTraceTlvReservedCounter
func (obj *patternFlowIpv6SRHPathTraceTlvReserved) Decrement() PatternFlowIpv6SRHPathTraceTlvReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHPathTraceTlvReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPathTraceTlvReservedCounter
func (obj *patternFlowIpv6SRHPathTraceTlvReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHPathTraceTlvReservedCounter value in the PatternFlowIpv6SRHPathTraceTlvReserved object
func (obj *patternFlowIpv6SRHPathTraceTlvReserved) SetDecrement(value PatternFlowIpv6SRHPathTraceTlvReservedCounter) PatternFlowIpv6SRHPathTraceTlvReserved {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvReserved.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SRHPathTraceTlvReserved.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIpv6SRHPathTraceTlvReserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHPathTraceTlvReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHPathTraceTlvReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHPathTraceTlvReserved")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHPathTraceTlvReserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHPathTraceTlvReserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
