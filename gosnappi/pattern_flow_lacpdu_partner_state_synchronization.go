package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateSynchronization *****
type patternFlowLacpduPartnerStateSynchronization struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerStateSynchronization
	marshaller      marshalPatternFlowLacpduPartnerStateSynchronization
	unMarshaller    unMarshalPatternFlowLacpduPartnerStateSynchronization
	incrementHolder PatternFlowLacpduPartnerStateSynchronizationCounter
	decrementHolder PatternFlowLacpduPartnerStateSynchronizationCounter
}

func NewPatternFlowLacpduPartnerStateSynchronization() PatternFlowLacpduPartnerStateSynchronization {
	obj := patternFlowLacpduPartnerStateSynchronization{obj: &otg.PatternFlowLacpduPartnerStateSynchronization{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateSynchronization) msg() *otg.PatternFlowLacpduPartnerStateSynchronization {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateSynchronization) setMsg(msg *otg.PatternFlowLacpduPartnerStateSynchronization) PatternFlowLacpduPartnerStateSynchronization {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateSynchronization struct {
	obj *patternFlowLacpduPartnerStateSynchronization
}

type marshalPatternFlowLacpduPartnerStateSynchronization interface {
	// ToProto marshals PatternFlowLacpduPartnerStateSynchronization to protobuf object *otg.PatternFlowLacpduPartnerStateSynchronization
	ToProto() (*otg.PatternFlowLacpduPartnerStateSynchronization, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateSynchronization to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateSynchronization to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateSynchronization to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateSynchronization struct {
	obj *patternFlowLacpduPartnerStateSynchronization
}

type unMarshalPatternFlowLacpduPartnerStateSynchronization interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateSynchronization from protobuf object *otg.PatternFlowLacpduPartnerStateSynchronization
	FromProto(msg *otg.PatternFlowLacpduPartnerStateSynchronization) (PatternFlowLacpduPartnerStateSynchronization, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateSynchronization from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateSynchronization from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateSynchronization from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateSynchronization) Marshal() marshalPatternFlowLacpduPartnerStateSynchronization {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateSynchronization{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateSynchronization) Unmarshal() unMarshalPatternFlowLacpduPartnerStateSynchronization {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateSynchronization{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateSynchronization) ToProto() (*otg.PatternFlowLacpduPartnerStateSynchronization, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateSynchronization) FromProto(msg *otg.PatternFlowLacpduPartnerStateSynchronization) (PatternFlowLacpduPartnerStateSynchronization, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateSynchronization) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateSynchronization) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateSynchronization) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateSynchronization) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateSynchronization) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateSynchronization) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateSynchronization) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateSynchronization) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateSynchronization) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateSynchronization) Clone() (PatternFlowLacpduPartnerStateSynchronization, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateSynchronization()
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

func (obj *patternFlowLacpduPartnerStateSynchronization) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerStateSynchronization is synchronization (1=In Sync, 0=Out of Sync)
type PatternFlowLacpduPartnerStateSynchronization interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateSynchronization to protobuf object *otg.PatternFlowLacpduPartnerStateSynchronization
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateSynchronization
	// setMsg unmarshals PatternFlowLacpduPartnerStateSynchronization from protobuf object *otg.PatternFlowLacpduPartnerStateSynchronization
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateSynchronization) PatternFlowLacpduPartnerStateSynchronization
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateSynchronization
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateSynchronization
	// validate validates PatternFlowLacpduPartnerStateSynchronization
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateSynchronization, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerStateSynchronizationChoiceEnum, set in PatternFlowLacpduPartnerStateSynchronization
	Choice() PatternFlowLacpduPartnerStateSynchronizationChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerStateSynchronizationChoiceEnum provided by user to PatternFlowLacpduPartnerStateSynchronization
	setChoice(value PatternFlowLacpduPartnerStateSynchronizationChoiceEnum) PatternFlowLacpduPartnerStateSynchronization
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerStateSynchronization
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerStateSynchronization.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerStateSynchronization
	SetValue(value uint32) PatternFlowLacpduPartnerStateSynchronization
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerStateSynchronization
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerStateSynchronization.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerStateSynchronization
	SetValues(value []uint32) PatternFlowLacpduPartnerStateSynchronization
	// Increment returns PatternFlowLacpduPartnerStateSynchronizationCounter, set in PatternFlowLacpduPartnerStateSynchronization.
	// PatternFlowLacpduPartnerStateSynchronizationCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerStateSynchronizationCounter
	// SetIncrement assigns PatternFlowLacpduPartnerStateSynchronizationCounter provided by user to PatternFlowLacpduPartnerStateSynchronization.
	// PatternFlowLacpduPartnerStateSynchronizationCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerStateSynchronizationCounter) PatternFlowLacpduPartnerStateSynchronization
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerStateSynchronization
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerStateSynchronizationCounter, set in PatternFlowLacpduPartnerStateSynchronization.
	// PatternFlowLacpduPartnerStateSynchronizationCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerStateSynchronizationCounter
	// SetDecrement assigns PatternFlowLacpduPartnerStateSynchronizationCounter provided by user to PatternFlowLacpduPartnerStateSynchronization.
	// PatternFlowLacpduPartnerStateSynchronizationCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerStateSynchronizationCounter) PatternFlowLacpduPartnerStateSynchronization
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerStateSynchronization
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerStateSynchronizationChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerStateSynchronization
var PatternFlowLacpduPartnerStateSynchronizationChoice = struct {
	VALUE     PatternFlowLacpduPartnerStateSynchronizationChoiceEnum
	VALUES    PatternFlowLacpduPartnerStateSynchronizationChoiceEnum
	INCREMENT PatternFlowLacpduPartnerStateSynchronizationChoiceEnum
	DECREMENT PatternFlowLacpduPartnerStateSynchronizationChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerStateSynchronizationChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerStateSynchronizationChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerStateSynchronizationChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerStateSynchronizationChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerStateSynchronization) Choice() PatternFlowLacpduPartnerStateSynchronizationChoiceEnum {
	return PatternFlowLacpduPartnerStateSynchronizationChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerStateSynchronization) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerStateSynchronization) setChoice(value PatternFlowLacpduPartnerStateSynchronizationChoiceEnum) PatternFlowLacpduPartnerStateSynchronization {
	intValue, ok := otg.PatternFlowLacpduPartnerStateSynchronization_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerStateSynchronizationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerStateSynchronization_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerStateSynchronizationChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerStateSynchronizationChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerStateSynchronizationChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerStateSynchronizationCounter().msg()
	}

	if value == PatternFlowLacpduPartnerStateSynchronizationChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerStateSynchronizationCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateSynchronization) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateSynchronizationChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateSynchronization) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerStateSynchronization object
func (obj *patternFlowLacpduPartnerStateSynchronization) SetValue(value uint32) PatternFlowLacpduPartnerStateSynchronization {
	obj.setChoice(PatternFlowLacpduPartnerStateSynchronizationChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerStateSynchronization) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerStateSynchronization object
func (obj *patternFlowLacpduPartnerStateSynchronization) SetValues(value []uint32) PatternFlowLacpduPartnerStateSynchronization {
	obj.setChoice(PatternFlowLacpduPartnerStateSynchronizationChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateSynchronizationCounter
func (obj *patternFlowLacpduPartnerStateSynchronization) Increment() PatternFlowLacpduPartnerStateSynchronizationCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateSynchronizationChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerStateSynchronizationCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateSynchronizationCounter
func (obj *patternFlowLacpduPartnerStateSynchronization) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerStateSynchronizationCounter value in the PatternFlowLacpduPartnerStateSynchronization object
func (obj *patternFlowLacpduPartnerStateSynchronization) SetIncrement(value PatternFlowLacpduPartnerStateSynchronizationCounter) PatternFlowLacpduPartnerStateSynchronization {
	obj.setChoice(PatternFlowLacpduPartnerStateSynchronizationChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateSynchronizationCounter
func (obj *patternFlowLacpduPartnerStateSynchronization) Decrement() PatternFlowLacpduPartnerStateSynchronizationCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateSynchronizationChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerStateSynchronizationCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateSynchronizationCounter
func (obj *patternFlowLacpduPartnerStateSynchronization) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerStateSynchronizationCounter value in the PatternFlowLacpduPartnerStateSynchronization object
func (obj *patternFlowLacpduPartnerStateSynchronization) SetDecrement(value PatternFlowLacpduPartnerStateSynchronizationCounter) PatternFlowLacpduPartnerStateSynchronization {
	obj.setChoice(PatternFlowLacpduPartnerStateSynchronizationChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerStateSynchronization) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateSynchronization.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerStateSynchronization.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerStateSynchronization) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerStateSynchronizationChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateSynchronizationChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateSynchronizationChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateSynchronizationChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateSynchronizationChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerStateSynchronizationChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerStateSynchronization")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerStateSynchronization_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerStateSynchronization_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
