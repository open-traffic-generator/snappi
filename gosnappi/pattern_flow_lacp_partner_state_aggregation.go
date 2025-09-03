package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateAggregation *****
type patternFlowLacpPartnerStateAggregation struct {
	validation
	obj             *otg.PatternFlowLacpPartnerStateAggregation
	marshaller      marshalPatternFlowLacpPartnerStateAggregation
	unMarshaller    unMarshalPatternFlowLacpPartnerStateAggregation
	incrementHolder PatternFlowLacpPartnerStateAggregationCounter
	decrementHolder PatternFlowLacpPartnerStateAggregationCounter
}

func NewPatternFlowLacpPartnerStateAggregation() PatternFlowLacpPartnerStateAggregation {
	obj := patternFlowLacpPartnerStateAggregation{obj: &otg.PatternFlowLacpPartnerStateAggregation{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateAggregation) msg() *otg.PatternFlowLacpPartnerStateAggregation {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateAggregation) setMsg(msg *otg.PatternFlowLacpPartnerStateAggregation) PatternFlowLacpPartnerStateAggregation {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateAggregation struct {
	obj *patternFlowLacpPartnerStateAggregation
}

type marshalPatternFlowLacpPartnerStateAggregation interface {
	// ToProto marshals PatternFlowLacpPartnerStateAggregation to protobuf object *otg.PatternFlowLacpPartnerStateAggregation
	ToProto() (*otg.PatternFlowLacpPartnerStateAggregation, error)
	// ToPbText marshals PatternFlowLacpPartnerStateAggregation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateAggregation to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateAggregation to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateAggregation struct {
	obj *patternFlowLacpPartnerStateAggregation
}

type unMarshalPatternFlowLacpPartnerStateAggregation interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateAggregation from protobuf object *otg.PatternFlowLacpPartnerStateAggregation
	FromProto(msg *otg.PatternFlowLacpPartnerStateAggregation) (PatternFlowLacpPartnerStateAggregation, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateAggregation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateAggregation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateAggregation from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateAggregation) Marshal() marshalPatternFlowLacpPartnerStateAggregation {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateAggregation{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateAggregation) Unmarshal() unMarshalPatternFlowLacpPartnerStateAggregation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateAggregation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateAggregation) ToProto() (*otg.PatternFlowLacpPartnerStateAggregation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateAggregation) FromProto(msg *otg.PatternFlowLacpPartnerStateAggregation) (PatternFlowLacpPartnerStateAggregation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateAggregation) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateAggregation) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateAggregation) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateAggregation) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateAggregation) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateAggregation) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateAggregation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateAggregation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateAggregation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateAggregation) Clone() (PatternFlowLacpPartnerStateAggregation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateAggregation()
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

func (obj *patternFlowLacpPartnerStateAggregation) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerStateAggregation is aggregation (1=Aggregatable, 0=Individual)
type PatternFlowLacpPartnerStateAggregation interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateAggregation to protobuf object *otg.PatternFlowLacpPartnerStateAggregation
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateAggregation
	// setMsg unmarshals PatternFlowLacpPartnerStateAggregation from protobuf object *otg.PatternFlowLacpPartnerStateAggregation
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateAggregation) PatternFlowLacpPartnerStateAggregation
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateAggregation
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateAggregation
	// validate validates PatternFlowLacpPartnerStateAggregation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateAggregation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerStateAggregationChoiceEnum, set in PatternFlowLacpPartnerStateAggregation
	Choice() PatternFlowLacpPartnerStateAggregationChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerStateAggregationChoiceEnum provided by user to PatternFlowLacpPartnerStateAggregation
	setChoice(value PatternFlowLacpPartnerStateAggregationChoiceEnum) PatternFlowLacpPartnerStateAggregation
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerStateAggregation
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerStateAggregation.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerStateAggregation
	SetValue(value uint32) PatternFlowLacpPartnerStateAggregation
	// HasValue checks if Value has been set in PatternFlowLacpPartnerStateAggregation
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerStateAggregation.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerStateAggregation
	SetValues(value []uint32) PatternFlowLacpPartnerStateAggregation
	// Increment returns PatternFlowLacpPartnerStateAggregationCounter, set in PatternFlowLacpPartnerStateAggregation.
	// PatternFlowLacpPartnerStateAggregationCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerStateAggregationCounter
	// SetIncrement assigns PatternFlowLacpPartnerStateAggregationCounter provided by user to PatternFlowLacpPartnerStateAggregation.
	// PatternFlowLacpPartnerStateAggregationCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerStateAggregationCounter) PatternFlowLacpPartnerStateAggregation
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerStateAggregation
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerStateAggregationCounter, set in PatternFlowLacpPartnerStateAggregation.
	// PatternFlowLacpPartnerStateAggregationCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerStateAggregationCounter
	// SetDecrement assigns PatternFlowLacpPartnerStateAggregationCounter provided by user to PatternFlowLacpPartnerStateAggregation.
	// PatternFlowLacpPartnerStateAggregationCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerStateAggregationCounter) PatternFlowLacpPartnerStateAggregation
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerStateAggregation
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerStateAggregationChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerStateAggregation
var PatternFlowLacpPartnerStateAggregationChoice = struct {
	VALUE     PatternFlowLacpPartnerStateAggregationChoiceEnum
	VALUES    PatternFlowLacpPartnerStateAggregationChoiceEnum
	INCREMENT PatternFlowLacpPartnerStateAggregationChoiceEnum
	DECREMENT PatternFlowLacpPartnerStateAggregationChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerStateAggregationChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerStateAggregationChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerStateAggregationChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerStateAggregationChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerStateAggregation) Choice() PatternFlowLacpPartnerStateAggregationChoiceEnum {
	return PatternFlowLacpPartnerStateAggregationChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerStateAggregation) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerStateAggregation) setChoice(value PatternFlowLacpPartnerStateAggregationChoiceEnum) PatternFlowLacpPartnerStateAggregation {
	intValue, ok := otg.PatternFlowLacpPartnerStateAggregation_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerStateAggregationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerStateAggregation_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerStateAggregationChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerStateAggregationChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerStateAggregationChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerStateAggregationCounter().msg()
	}

	if value == PatternFlowLacpPartnerStateAggregationChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerStateAggregationCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateAggregation) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerStateAggregationChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateAggregation) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerStateAggregation object
func (obj *patternFlowLacpPartnerStateAggregation) SetValue(value uint32) PatternFlowLacpPartnerStateAggregation {
	obj.setChoice(PatternFlowLacpPartnerStateAggregationChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerStateAggregation) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerStateAggregation object
func (obj *patternFlowLacpPartnerStateAggregation) SetValues(value []uint32) PatternFlowLacpPartnerStateAggregation {
	obj.setChoice(PatternFlowLacpPartnerStateAggregationChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateAggregationCounter
func (obj *patternFlowLacpPartnerStateAggregation) Increment() PatternFlowLacpPartnerStateAggregationCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerStateAggregationChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerStateAggregationCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateAggregationCounter
func (obj *patternFlowLacpPartnerStateAggregation) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerStateAggregationCounter value in the PatternFlowLacpPartnerStateAggregation object
func (obj *patternFlowLacpPartnerStateAggregation) SetIncrement(value PatternFlowLacpPartnerStateAggregationCounter) PatternFlowLacpPartnerStateAggregation {
	obj.setChoice(PatternFlowLacpPartnerStateAggregationChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateAggregationCounter
func (obj *patternFlowLacpPartnerStateAggregation) Decrement() PatternFlowLacpPartnerStateAggregationCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerStateAggregationChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerStateAggregationCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateAggregationCounter
func (obj *patternFlowLacpPartnerStateAggregation) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerStateAggregationCounter value in the PatternFlowLacpPartnerStateAggregation object
func (obj *patternFlowLacpPartnerStateAggregation) SetDecrement(value PatternFlowLacpPartnerStateAggregationCounter) PatternFlowLacpPartnerStateAggregation {
	obj.setChoice(PatternFlowLacpPartnerStateAggregationChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerStateAggregation) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateAggregation.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerStateAggregation.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpPartnerStateAggregation) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerStateAggregationChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateAggregationChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateAggregationChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateAggregationChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateAggregationChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerStateAggregationChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerStateAggregation")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerStateAggregation_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerStateAggregation_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
