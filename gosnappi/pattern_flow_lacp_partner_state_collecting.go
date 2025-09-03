package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateCollecting *****
type patternFlowLacpPartnerStateCollecting struct {
	validation
	obj             *otg.PatternFlowLacpPartnerStateCollecting
	marshaller      marshalPatternFlowLacpPartnerStateCollecting
	unMarshaller    unMarshalPatternFlowLacpPartnerStateCollecting
	incrementHolder PatternFlowLacpPartnerStateCollectingCounter
	decrementHolder PatternFlowLacpPartnerStateCollectingCounter
}

func NewPatternFlowLacpPartnerStateCollecting() PatternFlowLacpPartnerStateCollecting {
	obj := patternFlowLacpPartnerStateCollecting{obj: &otg.PatternFlowLacpPartnerStateCollecting{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateCollecting) msg() *otg.PatternFlowLacpPartnerStateCollecting {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateCollecting) setMsg(msg *otg.PatternFlowLacpPartnerStateCollecting) PatternFlowLacpPartnerStateCollecting {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateCollecting struct {
	obj *patternFlowLacpPartnerStateCollecting
}

type marshalPatternFlowLacpPartnerStateCollecting interface {
	// ToProto marshals PatternFlowLacpPartnerStateCollecting to protobuf object *otg.PatternFlowLacpPartnerStateCollecting
	ToProto() (*otg.PatternFlowLacpPartnerStateCollecting, error)
	// ToPbText marshals PatternFlowLacpPartnerStateCollecting to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateCollecting to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateCollecting to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateCollecting struct {
	obj *patternFlowLacpPartnerStateCollecting
}

type unMarshalPatternFlowLacpPartnerStateCollecting interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateCollecting from protobuf object *otg.PatternFlowLacpPartnerStateCollecting
	FromProto(msg *otg.PatternFlowLacpPartnerStateCollecting) (PatternFlowLacpPartnerStateCollecting, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateCollecting from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateCollecting from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateCollecting from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateCollecting) Marshal() marshalPatternFlowLacpPartnerStateCollecting {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateCollecting{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateCollecting) Unmarshal() unMarshalPatternFlowLacpPartnerStateCollecting {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateCollecting{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateCollecting) ToProto() (*otg.PatternFlowLacpPartnerStateCollecting, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateCollecting) FromProto(msg *otg.PatternFlowLacpPartnerStateCollecting) (PatternFlowLacpPartnerStateCollecting, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateCollecting) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateCollecting) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateCollecting) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateCollecting) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateCollecting) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateCollecting) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateCollecting) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateCollecting) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateCollecting) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateCollecting) Clone() (PatternFlowLacpPartnerStateCollecting, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateCollecting()
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

func (obj *patternFlowLacpPartnerStateCollecting) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerStateCollecting is collecting (1=Enabled, 0=Disabled)
type PatternFlowLacpPartnerStateCollecting interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateCollecting to protobuf object *otg.PatternFlowLacpPartnerStateCollecting
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateCollecting
	// setMsg unmarshals PatternFlowLacpPartnerStateCollecting from protobuf object *otg.PatternFlowLacpPartnerStateCollecting
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateCollecting) PatternFlowLacpPartnerStateCollecting
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateCollecting
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateCollecting
	// validate validates PatternFlowLacpPartnerStateCollecting
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateCollecting, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerStateCollectingChoiceEnum, set in PatternFlowLacpPartnerStateCollecting
	Choice() PatternFlowLacpPartnerStateCollectingChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerStateCollectingChoiceEnum provided by user to PatternFlowLacpPartnerStateCollecting
	setChoice(value PatternFlowLacpPartnerStateCollectingChoiceEnum) PatternFlowLacpPartnerStateCollecting
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerStateCollecting
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerStateCollecting.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerStateCollecting
	SetValue(value uint32) PatternFlowLacpPartnerStateCollecting
	// HasValue checks if Value has been set in PatternFlowLacpPartnerStateCollecting
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerStateCollecting.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerStateCollecting
	SetValues(value []uint32) PatternFlowLacpPartnerStateCollecting
	// Increment returns PatternFlowLacpPartnerStateCollectingCounter, set in PatternFlowLacpPartnerStateCollecting.
	// PatternFlowLacpPartnerStateCollectingCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerStateCollectingCounter
	// SetIncrement assigns PatternFlowLacpPartnerStateCollectingCounter provided by user to PatternFlowLacpPartnerStateCollecting.
	// PatternFlowLacpPartnerStateCollectingCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerStateCollectingCounter) PatternFlowLacpPartnerStateCollecting
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerStateCollecting
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerStateCollectingCounter, set in PatternFlowLacpPartnerStateCollecting.
	// PatternFlowLacpPartnerStateCollectingCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerStateCollectingCounter
	// SetDecrement assigns PatternFlowLacpPartnerStateCollectingCounter provided by user to PatternFlowLacpPartnerStateCollecting.
	// PatternFlowLacpPartnerStateCollectingCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerStateCollectingCounter) PatternFlowLacpPartnerStateCollecting
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerStateCollecting
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerStateCollectingChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerStateCollecting
var PatternFlowLacpPartnerStateCollectingChoice = struct {
	VALUE     PatternFlowLacpPartnerStateCollectingChoiceEnum
	VALUES    PatternFlowLacpPartnerStateCollectingChoiceEnum
	INCREMENT PatternFlowLacpPartnerStateCollectingChoiceEnum
	DECREMENT PatternFlowLacpPartnerStateCollectingChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerStateCollectingChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerStateCollectingChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerStateCollectingChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerStateCollectingChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerStateCollecting) Choice() PatternFlowLacpPartnerStateCollectingChoiceEnum {
	return PatternFlowLacpPartnerStateCollectingChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerStateCollecting) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerStateCollecting) setChoice(value PatternFlowLacpPartnerStateCollectingChoiceEnum) PatternFlowLacpPartnerStateCollecting {
	intValue, ok := otg.PatternFlowLacpPartnerStateCollecting_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerStateCollectingChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerStateCollecting_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerStateCollectingChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerStateCollectingChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerStateCollectingChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerStateCollectingCounter().msg()
	}

	if value == PatternFlowLacpPartnerStateCollectingChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerStateCollectingCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateCollecting) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerStateCollectingChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateCollecting) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerStateCollecting object
func (obj *patternFlowLacpPartnerStateCollecting) SetValue(value uint32) PatternFlowLacpPartnerStateCollecting {
	obj.setChoice(PatternFlowLacpPartnerStateCollectingChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerStateCollecting) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerStateCollecting object
func (obj *patternFlowLacpPartnerStateCollecting) SetValues(value []uint32) PatternFlowLacpPartnerStateCollecting {
	obj.setChoice(PatternFlowLacpPartnerStateCollectingChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateCollectingCounter
func (obj *patternFlowLacpPartnerStateCollecting) Increment() PatternFlowLacpPartnerStateCollectingCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerStateCollectingChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerStateCollectingCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateCollectingCounter
func (obj *patternFlowLacpPartnerStateCollecting) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerStateCollectingCounter value in the PatternFlowLacpPartnerStateCollecting object
func (obj *patternFlowLacpPartnerStateCollecting) SetIncrement(value PatternFlowLacpPartnerStateCollectingCounter) PatternFlowLacpPartnerStateCollecting {
	obj.setChoice(PatternFlowLacpPartnerStateCollectingChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateCollectingCounter
func (obj *patternFlowLacpPartnerStateCollecting) Decrement() PatternFlowLacpPartnerStateCollectingCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerStateCollectingChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerStateCollectingCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateCollectingCounter
func (obj *patternFlowLacpPartnerStateCollecting) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerStateCollectingCounter value in the PatternFlowLacpPartnerStateCollecting object
func (obj *patternFlowLacpPartnerStateCollecting) SetDecrement(value PatternFlowLacpPartnerStateCollectingCounter) PatternFlowLacpPartnerStateCollecting {
	obj.setChoice(PatternFlowLacpPartnerStateCollectingChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerStateCollecting) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateCollecting.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerStateCollecting.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpPartnerStateCollecting) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerStateCollectingChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateCollectingChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateCollectingChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateCollectingChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateCollectingChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerStateCollectingChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerStateCollecting")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerStateCollecting_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerStateCollecting_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
