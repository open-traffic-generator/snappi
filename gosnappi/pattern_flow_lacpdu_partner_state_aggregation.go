package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateAggregation *****
type patternFlowLacpduPartnerStateAggregation struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerStateAggregation
	marshaller      marshalPatternFlowLacpduPartnerStateAggregation
	unMarshaller    unMarshalPatternFlowLacpduPartnerStateAggregation
	incrementHolder PatternFlowLacpduPartnerStateAggregationCounter
	decrementHolder PatternFlowLacpduPartnerStateAggregationCounter
}

func NewPatternFlowLacpduPartnerStateAggregation() PatternFlowLacpduPartnerStateAggregation {
	obj := patternFlowLacpduPartnerStateAggregation{obj: &otg.PatternFlowLacpduPartnerStateAggregation{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateAggregation) msg() *otg.PatternFlowLacpduPartnerStateAggregation {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateAggregation) setMsg(msg *otg.PatternFlowLacpduPartnerStateAggregation) PatternFlowLacpduPartnerStateAggregation {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateAggregation struct {
	obj *patternFlowLacpduPartnerStateAggregation
}

type marshalPatternFlowLacpduPartnerStateAggregation interface {
	// ToProto marshals PatternFlowLacpduPartnerStateAggregation to protobuf object *otg.PatternFlowLacpduPartnerStateAggregation
	ToProto() (*otg.PatternFlowLacpduPartnerStateAggregation, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateAggregation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateAggregation to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateAggregation to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateAggregation struct {
	obj *patternFlowLacpduPartnerStateAggregation
}

type unMarshalPatternFlowLacpduPartnerStateAggregation interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateAggregation from protobuf object *otg.PatternFlowLacpduPartnerStateAggregation
	FromProto(msg *otg.PatternFlowLacpduPartnerStateAggregation) (PatternFlowLacpduPartnerStateAggregation, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateAggregation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateAggregation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateAggregation from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateAggregation) Marshal() marshalPatternFlowLacpduPartnerStateAggregation {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateAggregation{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateAggregation) Unmarshal() unMarshalPatternFlowLacpduPartnerStateAggregation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateAggregation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateAggregation) ToProto() (*otg.PatternFlowLacpduPartnerStateAggregation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateAggregation) FromProto(msg *otg.PatternFlowLacpduPartnerStateAggregation) (PatternFlowLacpduPartnerStateAggregation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateAggregation) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateAggregation) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateAggregation) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateAggregation) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateAggregation) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateAggregation) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateAggregation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateAggregation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateAggregation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateAggregation) Clone() (PatternFlowLacpduPartnerStateAggregation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateAggregation()
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

func (obj *patternFlowLacpduPartnerStateAggregation) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerStateAggregation is aggregation (1=Aggregatable, 0=Individual)
type PatternFlowLacpduPartnerStateAggregation interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateAggregation to protobuf object *otg.PatternFlowLacpduPartnerStateAggregation
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateAggregation
	// setMsg unmarshals PatternFlowLacpduPartnerStateAggregation from protobuf object *otg.PatternFlowLacpduPartnerStateAggregation
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateAggregation) PatternFlowLacpduPartnerStateAggregation
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateAggregation
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateAggregation
	// validate validates PatternFlowLacpduPartnerStateAggregation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateAggregation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerStateAggregationChoiceEnum, set in PatternFlowLacpduPartnerStateAggregation
	Choice() PatternFlowLacpduPartnerStateAggregationChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerStateAggregationChoiceEnum provided by user to PatternFlowLacpduPartnerStateAggregation
	setChoice(value PatternFlowLacpduPartnerStateAggregationChoiceEnum) PatternFlowLacpduPartnerStateAggregation
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerStateAggregation
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerStateAggregation.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerStateAggregation
	SetValue(value uint32) PatternFlowLacpduPartnerStateAggregation
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerStateAggregation
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerStateAggregation.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerStateAggregation
	SetValues(value []uint32) PatternFlowLacpduPartnerStateAggregation
	// Increment returns PatternFlowLacpduPartnerStateAggregationCounter, set in PatternFlowLacpduPartnerStateAggregation.
	// PatternFlowLacpduPartnerStateAggregationCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerStateAggregationCounter
	// SetIncrement assigns PatternFlowLacpduPartnerStateAggregationCounter provided by user to PatternFlowLacpduPartnerStateAggregation.
	// PatternFlowLacpduPartnerStateAggregationCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerStateAggregationCounter) PatternFlowLacpduPartnerStateAggregation
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerStateAggregation
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerStateAggregationCounter, set in PatternFlowLacpduPartnerStateAggregation.
	// PatternFlowLacpduPartnerStateAggregationCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerStateAggregationCounter
	// SetDecrement assigns PatternFlowLacpduPartnerStateAggregationCounter provided by user to PatternFlowLacpduPartnerStateAggregation.
	// PatternFlowLacpduPartnerStateAggregationCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerStateAggregationCounter) PatternFlowLacpduPartnerStateAggregation
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerStateAggregation
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerStateAggregationChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerStateAggregation
var PatternFlowLacpduPartnerStateAggregationChoice = struct {
	VALUE     PatternFlowLacpduPartnerStateAggregationChoiceEnum
	VALUES    PatternFlowLacpduPartnerStateAggregationChoiceEnum
	INCREMENT PatternFlowLacpduPartnerStateAggregationChoiceEnum
	DECREMENT PatternFlowLacpduPartnerStateAggregationChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerStateAggregationChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerStateAggregationChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerStateAggregationChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerStateAggregationChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerStateAggregation) Choice() PatternFlowLacpduPartnerStateAggregationChoiceEnum {
	return PatternFlowLacpduPartnerStateAggregationChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerStateAggregation) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerStateAggregation) setChoice(value PatternFlowLacpduPartnerStateAggregationChoiceEnum) PatternFlowLacpduPartnerStateAggregation {
	intValue, ok := otg.PatternFlowLacpduPartnerStateAggregation_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerStateAggregationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerStateAggregation_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerStateAggregationChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerStateAggregationChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerStateAggregationChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerStateAggregationCounter().msg()
	}

	if value == PatternFlowLacpduPartnerStateAggregationChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerStateAggregationCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateAggregation) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateAggregationChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateAggregation) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerStateAggregation object
func (obj *patternFlowLacpduPartnerStateAggregation) SetValue(value uint32) PatternFlowLacpduPartnerStateAggregation {
	obj.setChoice(PatternFlowLacpduPartnerStateAggregationChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerStateAggregation) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerStateAggregation object
func (obj *patternFlowLacpduPartnerStateAggregation) SetValues(value []uint32) PatternFlowLacpduPartnerStateAggregation {
	obj.setChoice(PatternFlowLacpduPartnerStateAggregationChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateAggregationCounter
func (obj *patternFlowLacpduPartnerStateAggregation) Increment() PatternFlowLacpduPartnerStateAggregationCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateAggregationChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerStateAggregationCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateAggregationCounter
func (obj *patternFlowLacpduPartnerStateAggregation) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerStateAggregationCounter value in the PatternFlowLacpduPartnerStateAggregation object
func (obj *patternFlowLacpduPartnerStateAggregation) SetIncrement(value PatternFlowLacpduPartnerStateAggregationCounter) PatternFlowLacpduPartnerStateAggregation {
	obj.setChoice(PatternFlowLacpduPartnerStateAggregationChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateAggregationCounter
func (obj *patternFlowLacpduPartnerStateAggregation) Decrement() PatternFlowLacpduPartnerStateAggregationCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateAggregationChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerStateAggregationCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateAggregationCounter
func (obj *patternFlowLacpduPartnerStateAggregation) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerStateAggregationCounter value in the PatternFlowLacpduPartnerStateAggregation object
func (obj *patternFlowLacpduPartnerStateAggregation) SetDecrement(value PatternFlowLacpduPartnerStateAggregationCounter) PatternFlowLacpduPartnerStateAggregation {
	obj.setChoice(PatternFlowLacpduPartnerStateAggregationChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerStateAggregation) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateAggregation.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerStateAggregation.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerStateAggregation) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerStateAggregationChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateAggregationChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateAggregationChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateAggregationChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateAggregationChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerStateAggregationChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerStateAggregation")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerStateAggregation_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerStateAggregation_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
