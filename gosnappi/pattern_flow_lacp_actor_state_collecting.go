package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateCollecting *****
type patternFlowLacpActorStateCollecting struct {
	validation
	obj             *otg.PatternFlowLacpActorStateCollecting
	marshaller      marshalPatternFlowLacpActorStateCollecting
	unMarshaller    unMarshalPatternFlowLacpActorStateCollecting
	incrementHolder PatternFlowLacpActorStateCollectingCounter
	decrementHolder PatternFlowLacpActorStateCollectingCounter
}

func NewPatternFlowLacpActorStateCollecting() PatternFlowLacpActorStateCollecting {
	obj := patternFlowLacpActorStateCollecting{obj: &otg.PatternFlowLacpActorStateCollecting{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateCollecting) msg() *otg.PatternFlowLacpActorStateCollecting {
	return obj.obj
}

func (obj *patternFlowLacpActorStateCollecting) setMsg(msg *otg.PatternFlowLacpActorStateCollecting) PatternFlowLacpActorStateCollecting {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateCollecting struct {
	obj *patternFlowLacpActorStateCollecting
}

type marshalPatternFlowLacpActorStateCollecting interface {
	// ToProto marshals PatternFlowLacpActorStateCollecting to protobuf object *otg.PatternFlowLacpActorStateCollecting
	ToProto() (*otg.PatternFlowLacpActorStateCollecting, error)
	// ToPbText marshals PatternFlowLacpActorStateCollecting to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateCollecting to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateCollecting to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateCollecting struct {
	obj *patternFlowLacpActorStateCollecting
}

type unMarshalPatternFlowLacpActorStateCollecting interface {
	// FromProto unmarshals PatternFlowLacpActorStateCollecting from protobuf object *otg.PatternFlowLacpActorStateCollecting
	FromProto(msg *otg.PatternFlowLacpActorStateCollecting) (PatternFlowLacpActorStateCollecting, error)
	// FromPbText unmarshals PatternFlowLacpActorStateCollecting from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateCollecting from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateCollecting from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateCollecting) Marshal() marshalPatternFlowLacpActorStateCollecting {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateCollecting{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateCollecting) Unmarshal() unMarshalPatternFlowLacpActorStateCollecting {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateCollecting{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateCollecting) ToProto() (*otg.PatternFlowLacpActorStateCollecting, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateCollecting) FromProto(msg *otg.PatternFlowLacpActorStateCollecting) (PatternFlowLacpActorStateCollecting, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateCollecting) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateCollecting) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateCollecting) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateCollecting) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateCollecting) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateCollecting) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateCollecting) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateCollecting) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateCollecting) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateCollecting) Clone() (PatternFlowLacpActorStateCollecting, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateCollecting()
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

func (obj *patternFlowLacpActorStateCollecting) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorStateCollecting is collecting (1=Enabled, 0=Disabled)
type PatternFlowLacpActorStateCollecting interface {
	Validation
	// msg marshals PatternFlowLacpActorStateCollecting to protobuf object *otg.PatternFlowLacpActorStateCollecting
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateCollecting
	// setMsg unmarshals PatternFlowLacpActorStateCollecting from protobuf object *otg.PatternFlowLacpActorStateCollecting
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateCollecting) PatternFlowLacpActorStateCollecting
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateCollecting
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateCollecting
	// validate validates PatternFlowLacpActorStateCollecting
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateCollecting, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorStateCollectingChoiceEnum, set in PatternFlowLacpActorStateCollecting
	Choice() PatternFlowLacpActorStateCollectingChoiceEnum
	// setChoice assigns PatternFlowLacpActorStateCollectingChoiceEnum provided by user to PatternFlowLacpActorStateCollecting
	setChoice(value PatternFlowLacpActorStateCollectingChoiceEnum) PatternFlowLacpActorStateCollecting
	// HasChoice checks if Choice has been set in PatternFlowLacpActorStateCollecting
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorStateCollecting.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorStateCollecting
	SetValue(value uint32) PatternFlowLacpActorStateCollecting
	// HasValue checks if Value has been set in PatternFlowLacpActorStateCollecting
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorStateCollecting.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorStateCollecting
	SetValues(value []uint32) PatternFlowLacpActorStateCollecting
	// Increment returns PatternFlowLacpActorStateCollectingCounter, set in PatternFlowLacpActorStateCollecting.
	// PatternFlowLacpActorStateCollectingCounter is integer counter pattern
	Increment() PatternFlowLacpActorStateCollectingCounter
	// SetIncrement assigns PatternFlowLacpActorStateCollectingCounter provided by user to PatternFlowLacpActorStateCollecting.
	// PatternFlowLacpActorStateCollectingCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorStateCollectingCounter) PatternFlowLacpActorStateCollecting
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorStateCollecting
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorStateCollectingCounter, set in PatternFlowLacpActorStateCollecting.
	// PatternFlowLacpActorStateCollectingCounter is integer counter pattern
	Decrement() PatternFlowLacpActorStateCollectingCounter
	// SetDecrement assigns PatternFlowLacpActorStateCollectingCounter provided by user to PatternFlowLacpActorStateCollecting.
	// PatternFlowLacpActorStateCollectingCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorStateCollectingCounter) PatternFlowLacpActorStateCollecting
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorStateCollecting
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorStateCollectingChoiceEnum string

// Enum of Choice on PatternFlowLacpActorStateCollecting
var PatternFlowLacpActorStateCollectingChoice = struct {
	VALUE     PatternFlowLacpActorStateCollectingChoiceEnum
	VALUES    PatternFlowLacpActorStateCollectingChoiceEnum
	INCREMENT PatternFlowLacpActorStateCollectingChoiceEnum
	DECREMENT PatternFlowLacpActorStateCollectingChoiceEnum
}{
	VALUE:     PatternFlowLacpActorStateCollectingChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorStateCollectingChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorStateCollectingChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorStateCollectingChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorStateCollecting) Choice() PatternFlowLacpActorStateCollectingChoiceEnum {
	return PatternFlowLacpActorStateCollectingChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorStateCollecting) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorStateCollecting) setChoice(value PatternFlowLacpActorStateCollectingChoiceEnum) PatternFlowLacpActorStateCollecting {
	intValue, ok := otg.PatternFlowLacpActorStateCollecting_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorStateCollectingChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorStateCollecting_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorStateCollectingChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorStateCollectingChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorStateCollectingChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorStateCollectingCounter().msg()
	}

	if value == PatternFlowLacpActorStateCollectingChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorStateCollectingCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateCollecting) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorStateCollectingChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateCollecting) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorStateCollecting object
func (obj *patternFlowLacpActorStateCollecting) SetValue(value uint32) PatternFlowLacpActorStateCollecting {
	obj.setChoice(PatternFlowLacpActorStateCollectingChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorStateCollecting) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorStateCollecting object
func (obj *patternFlowLacpActorStateCollecting) SetValues(value []uint32) PatternFlowLacpActorStateCollecting {
	obj.setChoice(PatternFlowLacpActorStateCollectingChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateCollectingCounter
func (obj *patternFlowLacpActorStateCollecting) Increment() PatternFlowLacpActorStateCollectingCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorStateCollectingChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorStateCollectingCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateCollectingCounter
func (obj *patternFlowLacpActorStateCollecting) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorStateCollectingCounter value in the PatternFlowLacpActorStateCollecting object
func (obj *patternFlowLacpActorStateCollecting) SetIncrement(value PatternFlowLacpActorStateCollectingCounter) PatternFlowLacpActorStateCollecting {
	obj.setChoice(PatternFlowLacpActorStateCollectingChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateCollectingCounter
func (obj *patternFlowLacpActorStateCollecting) Decrement() PatternFlowLacpActorStateCollectingCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorStateCollectingChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorStateCollectingCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateCollectingCounter
func (obj *patternFlowLacpActorStateCollecting) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorStateCollectingCounter value in the PatternFlowLacpActorStateCollecting object
func (obj *patternFlowLacpActorStateCollecting) SetDecrement(value PatternFlowLacpActorStateCollectingCounter) PatternFlowLacpActorStateCollecting {
	obj.setChoice(PatternFlowLacpActorStateCollectingChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorStateCollecting) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateCollecting.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorStateCollecting.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpActorStateCollecting) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorStateCollectingChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateCollectingChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorStateCollectingChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateCollectingChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateCollectingChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorStateCollectingChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorStateCollecting")
			}
		} else {
			intVal := otg.PatternFlowLacpActorStateCollecting_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorStateCollecting_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
