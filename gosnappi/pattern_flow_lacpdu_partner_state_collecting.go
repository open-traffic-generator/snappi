package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateCollecting *****
type patternFlowLacpduPartnerStateCollecting struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerStateCollecting
	marshaller      marshalPatternFlowLacpduPartnerStateCollecting
	unMarshaller    unMarshalPatternFlowLacpduPartnerStateCollecting
	incrementHolder PatternFlowLacpduPartnerStateCollectingCounter
	decrementHolder PatternFlowLacpduPartnerStateCollectingCounter
}

func NewPatternFlowLacpduPartnerStateCollecting() PatternFlowLacpduPartnerStateCollecting {
	obj := patternFlowLacpduPartnerStateCollecting{obj: &otg.PatternFlowLacpduPartnerStateCollecting{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateCollecting) msg() *otg.PatternFlowLacpduPartnerStateCollecting {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateCollecting) setMsg(msg *otg.PatternFlowLacpduPartnerStateCollecting) PatternFlowLacpduPartnerStateCollecting {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateCollecting struct {
	obj *patternFlowLacpduPartnerStateCollecting
}

type marshalPatternFlowLacpduPartnerStateCollecting interface {
	// ToProto marshals PatternFlowLacpduPartnerStateCollecting to protobuf object *otg.PatternFlowLacpduPartnerStateCollecting
	ToProto() (*otg.PatternFlowLacpduPartnerStateCollecting, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateCollecting to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateCollecting to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateCollecting to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateCollecting struct {
	obj *patternFlowLacpduPartnerStateCollecting
}

type unMarshalPatternFlowLacpduPartnerStateCollecting interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateCollecting from protobuf object *otg.PatternFlowLacpduPartnerStateCollecting
	FromProto(msg *otg.PatternFlowLacpduPartnerStateCollecting) (PatternFlowLacpduPartnerStateCollecting, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateCollecting from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateCollecting from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateCollecting from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateCollecting) Marshal() marshalPatternFlowLacpduPartnerStateCollecting {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateCollecting{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateCollecting) Unmarshal() unMarshalPatternFlowLacpduPartnerStateCollecting {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateCollecting{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateCollecting) ToProto() (*otg.PatternFlowLacpduPartnerStateCollecting, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateCollecting) FromProto(msg *otg.PatternFlowLacpduPartnerStateCollecting) (PatternFlowLacpduPartnerStateCollecting, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateCollecting) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateCollecting) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateCollecting) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateCollecting) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateCollecting) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateCollecting) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateCollecting) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateCollecting) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateCollecting) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateCollecting) Clone() (PatternFlowLacpduPartnerStateCollecting, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateCollecting()
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

func (obj *patternFlowLacpduPartnerStateCollecting) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerStateCollecting is collecting (1=Enabled, 0=Disabled)
type PatternFlowLacpduPartnerStateCollecting interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateCollecting to protobuf object *otg.PatternFlowLacpduPartnerStateCollecting
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateCollecting
	// setMsg unmarshals PatternFlowLacpduPartnerStateCollecting from protobuf object *otg.PatternFlowLacpduPartnerStateCollecting
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateCollecting) PatternFlowLacpduPartnerStateCollecting
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateCollecting
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateCollecting
	// validate validates PatternFlowLacpduPartnerStateCollecting
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateCollecting, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerStateCollectingChoiceEnum, set in PatternFlowLacpduPartnerStateCollecting
	Choice() PatternFlowLacpduPartnerStateCollectingChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerStateCollectingChoiceEnum provided by user to PatternFlowLacpduPartnerStateCollecting
	setChoice(value PatternFlowLacpduPartnerStateCollectingChoiceEnum) PatternFlowLacpduPartnerStateCollecting
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerStateCollecting
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerStateCollecting.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerStateCollecting
	SetValue(value uint32) PatternFlowLacpduPartnerStateCollecting
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerStateCollecting
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerStateCollecting.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerStateCollecting
	SetValues(value []uint32) PatternFlowLacpduPartnerStateCollecting
	// Increment returns PatternFlowLacpduPartnerStateCollectingCounter, set in PatternFlowLacpduPartnerStateCollecting.
	// PatternFlowLacpduPartnerStateCollectingCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerStateCollectingCounter
	// SetIncrement assigns PatternFlowLacpduPartnerStateCollectingCounter provided by user to PatternFlowLacpduPartnerStateCollecting.
	// PatternFlowLacpduPartnerStateCollectingCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerStateCollectingCounter) PatternFlowLacpduPartnerStateCollecting
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerStateCollecting
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerStateCollectingCounter, set in PatternFlowLacpduPartnerStateCollecting.
	// PatternFlowLacpduPartnerStateCollectingCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerStateCollectingCounter
	// SetDecrement assigns PatternFlowLacpduPartnerStateCollectingCounter provided by user to PatternFlowLacpduPartnerStateCollecting.
	// PatternFlowLacpduPartnerStateCollectingCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerStateCollectingCounter) PatternFlowLacpduPartnerStateCollecting
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerStateCollecting
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerStateCollectingChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerStateCollecting
var PatternFlowLacpduPartnerStateCollectingChoice = struct {
	VALUE     PatternFlowLacpduPartnerStateCollectingChoiceEnum
	VALUES    PatternFlowLacpduPartnerStateCollectingChoiceEnum
	INCREMENT PatternFlowLacpduPartnerStateCollectingChoiceEnum
	DECREMENT PatternFlowLacpduPartnerStateCollectingChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerStateCollectingChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerStateCollectingChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerStateCollectingChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerStateCollectingChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerStateCollecting) Choice() PatternFlowLacpduPartnerStateCollectingChoiceEnum {
	return PatternFlowLacpduPartnerStateCollectingChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerStateCollecting) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerStateCollecting) setChoice(value PatternFlowLacpduPartnerStateCollectingChoiceEnum) PatternFlowLacpduPartnerStateCollecting {
	intValue, ok := otg.PatternFlowLacpduPartnerStateCollecting_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerStateCollectingChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerStateCollecting_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerStateCollectingChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerStateCollectingChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerStateCollectingChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerStateCollectingCounter().msg()
	}

	if value == PatternFlowLacpduPartnerStateCollectingChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerStateCollectingCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateCollecting) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateCollectingChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateCollecting) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerStateCollecting object
func (obj *patternFlowLacpduPartnerStateCollecting) SetValue(value uint32) PatternFlowLacpduPartnerStateCollecting {
	obj.setChoice(PatternFlowLacpduPartnerStateCollectingChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerStateCollecting) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerStateCollecting object
func (obj *patternFlowLacpduPartnerStateCollecting) SetValues(value []uint32) PatternFlowLacpduPartnerStateCollecting {
	obj.setChoice(PatternFlowLacpduPartnerStateCollectingChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateCollectingCounter
func (obj *patternFlowLacpduPartnerStateCollecting) Increment() PatternFlowLacpduPartnerStateCollectingCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateCollectingChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerStateCollectingCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateCollectingCounter
func (obj *patternFlowLacpduPartnerStateCollecting) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerStateCollectingCounter value in the PatternFlowLacpduPartnerStateCollecting object
func (obj *patternFlowLacpduPartnerStateCollecting) SetIncrement(value PatternFlowLacpduPartnerStateCollectingCounter) PatternFlowLacpduPartnerStateCollecting {
	obj.setChoice(PatternFlowLacpduPartnerStateCollectingChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateCollectingCounter
func (obj *patternFlowLacpduPartnerStateCollecting) Decrement() PatternFlowLacpduPartnerStateCollectingCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateCollectingChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerStateCollectingCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateCollectingCounter
func (obj *patternFlowLacpduPartnerStateCollecting) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerStateCollectingCounter value in the PatternFlowLacpduPartnerStateCollecting object
func (obj *patternFlowLacpduPartnerStateCollecting) SetDecrement(value PatternFlowLacpduPartnerStateCollectingCounter) PatternFlowLacpduPartnerStateCollecting {
	obj.setChoice(PatternFlowLacpduPartnerStateCollectingChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerStateCollecting) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateCollecting.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerStateCollecting.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerStateCollecting) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerStateCollectingChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateCollectingChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateCollectingChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateCollectingChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateCollectingChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerStateCollectingChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerStateCollecting")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerStateCollecting_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerStateCollecting_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
