package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateCollecting *****
type patternFlowLacpduActorStateCollecting struct {
	validation
	obj             *otg.PatternFlowLacpduActorStateCollecting
	marshaller      marshalPatternFlowLacpduActorStateCollecting
	unMarshaller    unMarshalPatternFlowLacpduActorStateCollecting
	incrementHolder PatternFlowLacpduActorStateCollectingCounter
	decrementHolder PatternFlowLacpduActorStateCollectingCounter
}

func NewPatternFlowLacpduActorStateCollecting() PatternFlowLacpduActorStateCollecting {
	obj := patternFlowLacpduActorStateCollecting{obj: &otg.PatternFlowLacpduActorStateCollecting{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateCollecting) msg() *otg.PatternFlowLacpduActorStateCollecting {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateCollecting) setMsg(msg *otg.PatternFlowLacpduActorStateCollecting) PatternFlowLacpduActorStateCollecting {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateCollecting struct {
	obj *patternFlowLacpduActorStateCollecting
}

type marshalPatternFlowLacpduActorStateCollecting interface {
	// ToProto marshals PatternFlowLacpduActorStateCollecting to protobuf object *otg.PatternFlowLacpduActorStateCollecting
	ToProto() (*otg.PatternFlowLacpduActorStateCollecting, error)
	// ToPbText marshals PatternFlowLacpduActorStateCollecting to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateCollecting to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateCollecting to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateCollecting struct {
	obj *patternFlowLacpduActorStateCollecting
}

type unMarshalPatternFlowLacpduActorStateCollecting interface {
	// FromProto unmarshals PatternFlowLacpduActorStateCollecting from protobuf object *otg.PatternFlowLacpduActorStateCollecting
	FromProto(msg *otg.PatternFlowLacpduActorStateCollecting) (PatternFlowLacpduActorStateCollecting, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateCollecting from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateCollecting from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateCollecting from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateCollecting) Marshal() marshalPatternFlowLacpduActorStateCollecting {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateCollecting{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateCollecting) Unmarshal() unMarshalPatternFlowLacpduActorStateCollecting {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateCollecting{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateCollecting) ToProto() (*otg.PatternFlowLacpduActorStateCollecting, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateCollecting) FromProto(msg *otg.PatternFlowLacpduActorStateCollecting) (PatternFlowLacpduActorStateCollecting, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateCollecting) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateCollecting) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateCollecting) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateCollecting) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateCollecting) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateCollecting) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateCollecting) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateCollecting) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateCollecting) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateCollecting) Clone() (PatternFlowLacpduActorStateCollecting, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateCollecting()
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

func (obj *patternFlowLacpduActorStateCollecting) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorStateCollecting is collecting (1=Enabled, 0=Disabled)
type PatternFlowLacpduActorStateCollecting interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateCollecting to protobuf object *otg.PatternFlowLacpduActorStateCollecting
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateCollecting
	// setMsg unmarshals PatternFlowLacpduActorStateCollecting from protobuf object *otg.PatternFlowLacpduActorStateCollecting
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateCollecting) PatternFlowLacpduActorStateCollecting
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateCollecting
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateCollecting
	// validate validates PatternFlowLacpduActorStateCollecting
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateCollecting, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorStateCollectingChoiceEnum, set in PatternFlowLacpduActorStateCollecting
	Choice() PatternFlowLacpduActorStateCollectingChoiceEnum
	// setChoice assigns PatternFlowLacpduActorStateCollectingChoiceEnum provided by user to PatternFlowLacpduActorStateCollecting
	setChoice(value PatternFlowLacpduActorStateCollectingChoiceEnum) PatternFlowLacpduActorStateCollecting
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorStateCollecting
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorStateCollecting.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorStateCollecting
	SetValue(value uint32) PatternFlowLacpduActorStateCollecting
	// HasValue checks if Value has been set in PatternFlowLacpduActorStateCollecting
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorStateCollecting.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorStateCollecting
	SetValues(value []uint32) PatternFlowLacpduActorStateCollecting
	// Increment returns PatternFlowLacpduActorStateCollectingCounter, set in PatternFlowLacpduActorStateCollecting.
	// PatternFlowLacpduActorStateCollectingCounter is integer counter pattern
	Increment() PatternFlowLacpduActorStateCollectingCounter
	// SetIncrement assigns PatternFlowLacpduActorStateCollectingCounter provided by user to PatternFlowLacpduActorStateCollecting.
	// PatternFlowLacpduActorStateCollectingCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorStateCollectingCounter) PatternFlowLacpduActorStateCollecting
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorStateCollecting
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorStateCollectingCounter, set in PatternFlowLacpduActorStateCollecting.
	// PatternFlowLacpduActorStateCollectingCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorStateCollectingCounter
	// SetDecrement assigns PatternFlowLacpduActorStateCollectingCounter provided by user to PatternFlowLacpduActorStateCollecting.
	// PatternFlowLacpduActorStateCollectingCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorStateCollectingCounter) PatternFlowLacpduActorStateCollecting
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorStateCollecting
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorStateCollectingChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorStateCollecting
var PatternFlowLacpduActorStateCollectingChoice = struct {
	VALUE     PatternFlowLacpduActorStateCollectingChoiceEnum
	VALUES    PatternFlowLacpduActorStateCollectingChoiceEnum
	INCREMENT PatternFlowLacpduActorStateCollectingChoiceEnum
	DECREMENT PatternFlowLacpduActorStateCollectingChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorStateCollectingChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorStateCollectingChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorStateCollectingChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorStateCollectingChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorStateCollecting) Choice() PatternFlowLacpduActorStateCollectingChoiceEnum {
	return PatternFlowLacpduActorStateCollectingChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorStateCollecting) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorStateCollecting) setChoice(value PatternFlowLacpduActorStateCollectingChoiceEnum) PatternFlowLacpduActorStateCollecting {
	intValue, ok := otg.PatternFlowLacpduActorStateCollecting_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorStateCollectingChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorStateCollecting_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorStateCollectingChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorStateCollectingChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorStateCollectingChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorStateCollectingCounter().msg()
	}

	if value == PatternFlowLacpduActorStateCollectingChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorStateCollectingCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateCollecting) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorStateCollectingChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateCollecting) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorStateCollecting object
func (obj *patternFlowLacpduActorStateCollecting) SetValue(value uint32) PatternFlowLacpduActorStateCollecting {
	obj.setChoice(PatternFlowLacpduActorStateCollectingChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorStateCollecting) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorStateCollecting object
func (obj *patternFlowLacpduActorStateCollecting) SetValues(value []uint32) PatternFlowLacpduActorStateCollecting {
	obj.setChoice(PatternFlowLacpduActorStateCollectingChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateCollectingCounter
func (obj *patternFlowLacpduActorStateCollecting) Increment() PatternFlowLacpduActorStateCollectingCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorStateCollectingChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorStateCollectingCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateCollectingCounter
func (obj *patternFlowLacpduActorStateCollecting) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorStateCollectingCounter value in the PatternFlowLacpduActorStateCollecting object
func (obj *patternFlowLacpduActorStateCollecting) SetIncrement(value PatternFlowLacpduActorStateCollectingCounter) PatternFlowLacpduActorStateCollecting {
	obj.setChoice(PatternFlowLacpduActorStateCollectingChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateCollectingCounter
func (obj *patternFlowLacpduActorStateCollecting) Decrement() PatternFlowLacpduActorStateCollectingCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorStateCollectingChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorStateCollectingCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateCollectingCounter
func (obj *patternFlowLacpduActorStateCollecting) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorStateCollectingCounter value in the PatternFlowLacpduActorStateCollecting object
func (obj *patternFlowLacpduActorStateCollecting) SetDecrement(value PatternFlowLacpduActorStateCollectingCounter) PatternFlowLacpduActorStateCollecting {
	obj.setChoice(PatternFlowLacpduActorStateCollectingChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorStateCollecting) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateCollecting.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorStateCollecting.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduActorStateCollecting) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorStateCollectingChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateCollectingChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorStateCollectingChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateCollectingChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateCollectingChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorStateCollectingChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorStateCollecting")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorStateCollecting_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorStateCollecting_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
