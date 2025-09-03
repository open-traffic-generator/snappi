package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateSynchronization *****
type patternFlowLacpPartnerStateSynchronization struct {
	validation
	obj             *otg.PatternFlowLacpPartnerStateSynchronization
	marshaller      marshalPatternFlowLacpPartnerStateSynchronization
	unMarshaller    unMarshalPatternFlowLacpPartnerStateSynchronization
	incrementHolder PatternFlowLacpPartnerStateSynchronizationCounter
	decrementHolder PatternFlowLacpPartnerStateSynchronizationCounter
}

func NewPatternFlowLacpPartnerStateSynchronization() PatternFlowLacpPartnerStateSynchronization {
	obj := patternFlowLacpPartnerStateSynchronization{obj: &otg.PatternFlowLacpPartnerStateSynchronization{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateSynchronization) msg() *otg.PatternFlowLacpPartnerStateSynchronization {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateSynchronization) setMsg(msg *otg.PatternFlowLacpPartnerStateSynchronization) PatternFlowLacpPartnerStateSynchronization {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateSynchronization struct {
	obj *patternFlowLacpPartnerStateSynchronization
}

type marshalPatternFlowLacpPartnerStateSynchronization interface {
	// ToProto marshals PatternFlowLacpPartnerStateSynchronization to protobuf object *otg.PatternFlowLacpPartnerStateSynchronization
	ToProto() (*otg.PatternFlowLacpPartnerStateSynchronization, error)
	// ToPbText marshals PatternFlowLacpPartnerStateSynchronization to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateSynchronization to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateSynchronization to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateSynchronization struct {
	obj *patternFlowLacpPartnerStateSynchronization
}

type unMarshalPatternFlowLacpPartnerStateSynchronization interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateSynchronization from protobuf object *otg.PatternFlowLacpPartnerStateSynchronization
	FromProto(msg *otg.PatternFlowLacpPartnerStateSynchronization) (PatternFlowLacpPartnerStateSynchronization, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateSynchronization from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateSynchronization from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateSynchronization from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateSynchronization) Marshal() marshalPatternFlowLacpPartnerStateSynchronization {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateSynchronization{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateSynchronization) Unmarshal() unMarshalPatternFlowLacpPartnerStateSynchronization {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateSynchronization{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateSynchronization) ToProto() (*otg.PatternFlowLacpPartnerStateSynchronization, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateSynchronization) FromProto(msg *otg.PatternFlowLacpPartnerStateSynchronization) (PatternFlowLacpPartnerStateSynchronization, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateSynchronization) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateSynchronization) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateSynchronization) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateSynchronization) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateSynchronization) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateSynchronization) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateSynchronization) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateSynchronization) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateSynchronization) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateSynchronization) Clone() (PatternFlowLacpPartnerStateSynchronization, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateSynchronization()
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

func (obj *patternFlowLacpPartnerStateSynchronization) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerStateSynchronization is synchronization (1=In Sync, 0=Out of Sync)
type PatternFlowLacpPartnerStateSynchronization interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateSynchronization to protobuf object *otg.PatternFlowLacpPartnerStateSynchronization
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateSynchronization
	// setMsg unmarshals PatternFlowLacpPartnerStateSynchronization from protobuf object *otg.PatternFlowLacpPartnerStateSynchronization
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateSynchronization) PatternFlowLacpPartnerStateSynchronization
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateSynchronization
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateSynchronization
	// validate validates PatternFlowLacpPartnerStateSynchronization
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateSynchronization, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerStateSynchronizationChoiceEnum, set in PatternFlowLacpPartnerStateSynchronization
	Choice() PatternFlowLacpPartnerStateSynchronizationChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerStateSynchronizationChoiceEnum provided by user to PatternFlowLacpPartnerStateSynchronization
	setChoice(value PatternFlowLacpPartnerStateSynchronizationChoiceEnum) PatternFlowLacpPartnerStateSynchronization
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerStateSynchronization
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerStateSynchronization.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerStateSynchronization
	SetValue(value uint32) PatternFlowLacpPartnerStateSynchronization
	// HasValue checks if Value has been set in PatternFlowLacpPartnerStateSynchronization
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerStateSynchronization.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerStateSynchronization
	SetValues(value []uint32) PatternFlowLacpPartnerStateSynchronization
	// Increment returns PatternFlowLacpPartnerStateSynchronizationCounter, set in PatternFlowLacpPartnerStateSynchronization.
	// PatternFlowLacpPartnerStateSynchronizationCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerStateSynchronizationCounter
	// SetIncrement assigns PatternFlowLacpPartnerStateSynchronizationCounter provided by user to PatternFlowLacpPartnerStateSynchronization.
	// PatternFlowLacpPartnerStateSynchronizationCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerStateSynchronizationCounter) PatternFlowLacpPartnerStateSynchronization
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerStateSynchronization
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerStateSynchronizationCounter, set in PatternFlowLacpPartnerStateSynchronization.
	// PatternFlowLacpPartnerStateSynchronizationCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerStateSynchronizationCounter
	// SetDecrement assigns PatternFlowLacpPartnerStateSynchronizationCounter provided by user to PatternFlowLacpPartnerStateSynchronization.
	// PatternFlowLacpPartnerStateSynchronizationCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerStateSynchronizationCounter) PatternFlowLacpPartnerStateSynchronization
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerStateSynchronization
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerStateSynchronizationChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerStateSynchronization
var PatternFlowLacpPartnerStateSynchronizationChoice = struct {
	VALUE     PatternFlowLacpPartnerStateSynchronizationChoiceEnum
	VALUES    PatternFlowLacpPartnerStateSynchronizationChoiceEnum
	INCREMENT PatternFlowLacpPartnerStateSynchronizationChoiceEnum
	DECREMENT PatternFlowLacpPartnerStateSynchronizationChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerStateSynchronizationChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerStateSynchronizationChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerStateSynchronizationChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerStateSynchronizationChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerStateSynchronization) Choice() PatternFlowLacpPartnerStateSynchronizationChoiceEnum {
	return PatternFlowLacpPartnerStateSynchronizationChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerStateSynchronization) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerStateSynchronization) setChoice(value PatternFlowLacpPartnerStateSynchronizationChoiceEnum) PatternFlowLacpPartnerStateSynchronization {
	intValue, ok := otg.PatternFlowLacpPartnerStateSynchronization_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerStateSynchronizationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerStateSynchronization_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerStateSynchronizationChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerStateSynchronizationChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerStateSynchronizationChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerStateSynchronizationCounter().msg()
	}

	if value == PatternFlowLacpPartnerStateSynchronizationChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerStateSynchronizationCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateSynchronization) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerStateSynchronizationChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateSynchronization) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerStateSynchronization object
func (obj *patternFlowLacpPartnerStateSynchronization) SetValue(value uint32) PatternFlowLacpPartnerStateSynchronization {
	obj.setChoice(PatternFlowLacpPartnerStateSynchronizationChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerStateSynchronization) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerStateSynchronization object
func (obj *patternFlowLacpPartnerStateSynchronization) SetValues(value []uint32) PatternFlowLacpPartnerStateSynchronization {
	obj.setChoice(PatternFlowLacpPartnerStateSynchronizationChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateSynchronizationCounter
func (obj *patternFlowLacpPartnerStateSynchronization) Increment() PatternFlowLacpPartnerStateSynchronizationCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerStateSynchronizationChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerStateSynchronizationCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateSynchronizationCounter
func (obj *patternFlowLacpPartnerStateSynchronization) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerStateSynchronizationCounter value in the PatternFlowLacpPartnerStateSynchronization object
func (obj *patternFlowLacpPartnerStateSynchronization) SetIncrement(value PatternFlowLacpPartnerStateSynchronizationCounter) PatternFlowLacpPartnerStateSynchronization {
	obj.setChoice(PatternFlowLacpPartnerStateSynchronizationChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateSynchronizationCounter
func (obj *patternFlowLacpPartnerStateSynchronization) Decrement() PatternFlowLacpPartnerStateSynchronizationCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerStateSynchronizationChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerStateSynchronizationCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateSynchronizationCounter
func (obj *patternFlowLacpPartnerStateSynchronization) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerStateSynchronizationCounter value in the PatternFlowLacpPartnerStateSynchronization object
func (obj *patternFlowLacpPartnerStateSynchronization) SetDecrement(value PatternFlowLacpPartnerStateSynchronizationCounter) PatternFlowLacpPartnerStateSynchronization {
	obj.setChoice(PatternFlowLacpPartnerStateSynchronizationChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerStateSynchronization) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateSynchronization.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerStateSynchronization.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpPartnerStateSynchronization) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerStateSynchronizationChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateSynchronizationChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateSynchronizationChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateSynchronizationChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateSynchronizationChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerStateSynchronizationChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerStateSynchronization")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerStateSynchronization_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerStateSynchronization_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
