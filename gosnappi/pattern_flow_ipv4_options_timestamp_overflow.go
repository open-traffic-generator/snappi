package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampOverflow *****
type patternFlowIpv4OptionsTimestampOverflow struct {
	validation
	obj             *otg.PatternFlowIpv4OptionsTimestampOverflow
	marshaller      marshalPatternFlowIpv4OptionsTimestampOverflow
	unMarshaller    unMarshalPatternFlowIpv4OptionsTimestampOverflow
	incrementHolder PatternFlowIpv4OptionsTimestampOverflowCounter
	decrementHolder PatternFlowIpv4OptionsTimestampOverflowCounter
}

func NewPatternFlowIpv4OptionsTimestampOverflow() PatternFlowIpv4OptionsTimestampOverflow {
	obj := patternFlowIpv4OptionsTimestampOverflow{obj: &otg.PatternFlowIpv4OptionsTimestampOverflow{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampOverflow) msg() *otg.PatternFlowIpv4OptionsTimestampOverflow {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampOverflow) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampOverflow) PatternFlowIpv4OptionsTimestampOverflow {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampOverflow struct {
	obj *patternFlowIpv4OptionsTimestampOverflow
}

type marshalPatternFlowIpv4OptionsTimestampOverflow interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampOverflow to protobuf object *otg.PatternFlowIpv4OptionsTimestampOverflow
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampOverflow, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampOverflow to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampOverflow to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampOverflow to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampOverflow struct {
	obj *patternFlowIpv4OptionsTimestampOverflow
}

type unMarshalPatternFlowIpv4OptionsTimestampOverflow interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampOverflow from protobuf object *otg.PatternFlowIpv4OptionsTimestampOverflow
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampOverflow) (PatternFlowIpv4OptionsTimestampOverflow, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampOverflow from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampOverflow from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampOverflow from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampOverflow) Marshal() marshalPatternFlowIpv4OptionsTimestampOverflow {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampOverflow{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampOverflow) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampOverflow {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampOverflow{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampOverflow) ToProto() (*otg.PatternFlowIpv4OptionsTimestampOverflow, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampOverflow) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampOverflow) (PatternFlowIpv4OptionsTimestampOverflow, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampOverflow) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampOverflow) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampOverflow) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampOverflow) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampOverflow) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampOverflow) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampOverflow) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampOverflow) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampOverflow) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampOverflow) Clone() (PatternFlowIpv4OptionsTimestampOverflow, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampOverflow()
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

func (obj *patternFlowIpv4OptionsTimestampOverflow) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4OptionsTimestampOverflow is a counter that indicates the number of intermediate nodes that were unable to record a timestamp because the options data area was full.
type PatternFlowIpv4OptionsTimestampOverflow interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampOverflow to protobuf object *otg.PatternFlowIpv4OptionsTimestampOverflow
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampOverflow
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampOverflow from protobuf object *otg.PatternFlowIpv4OptionsTimestampOverflow
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampOverflow) PatternFlowIpv4OptionsTimestampOverflow
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampOverflow
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampOverflow
	// validate validates PatternFlowIpv4OptionsTimestampOverflow
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampOverflow, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4OptionsTimestampOverflowChoiceEnum, set in PatternFlowIpv4OptionsTimestampOverflow
	Choice() PatternFlowIpv4OptionsTimestampOverflowChoiceEnum
	// setChoice assigns PatternFlowIpv4OptionsTimestampOverflowChoiceEnum provided by user to PatternFlowIpv4OptionsTimestampOverflow
	setChoice(value PatternFlowIpv4OptionsTimestampOverflowChoiceEnum) PatternFlowIpv4OptionsTimestampOverflow
	// HasChoice checks if Choice has been set in PatternFlowIpv4OptionsTimestampOverflow
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4OptionsTimestampOverflow.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampOverflow
	SetValue(value uint32) PatternFlowIpv4OptionsTimestampOverflow
	// HasValue checks if Value has been set in PatternFlowIpv4OptionsTimestampOverflow
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4OptionsTimestampOverflow.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4OptionsTimestampOverflow
	SetValues(value []uint32) PatternFlowIpv4OptionsTimestampOverflow
	// Increment returns PatternFlowIpv4OptionsTimestampOverflowCounter, set in PatternFlowIpv4OptionsTimestampOverflow.
	// PatternFlowIpv4OptionsTimestampOverflowCounter is integer counter pattern
	Increment() PatternFlowIpv4OptionsTimestampOverflowCounter
	// SetIncrement assigns PatternFlowIpv4OptionsTimestampOverflowCounter provided by user to PatternFlowIpv4OptionsTimestampOverflow.
	// PatternFlowIpv4OptionsTimestampOverflowCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4OptionsTimestampOverflowCounter) PatternFlowIpv4OptionsTimestampOverflow
	// HasIncrement checks if Increment has been set in PatternFlowIpv4OptionsTimestampOverflow
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4OptionsTimestampOverflowCounter, set in PatternFlowIpv4OptionsTimestampOverflow.
	// PatternFlowIpv4OptionsTimestampOverflowCounter is integer counter pattern
	Decrement() PatternFlowIpv4OptionsTimestampOverflowCounter
	// SetDecrement assigns PatternFlowIpv4OptionsTimestampOverflowCounter provided by user to PatternFlowIpv4OptionsTimestampOverflow.
	// PatternFlowIpv4OptionsTimestampOverflowCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4OptionsTimestampOverflowCounter) PatternFlowIpv4OptionsTimestampOverflow
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4OptionsTimestampOverflow
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv4OptionsTimestampOverflowChoiceEnum string

// Enum of Choice on PatternFlowIpv4OptionsTimestampOverflow
var PatternFlowIpv4OptionsTimestampOverflowChoice = struct {
	VALUE     PatternFlowIpv4OptionsTimestampOverflowChoiceEnum
	VALUES    PatternFlowIpv4OptionsTimestampOverflowChoiceEnum
	INCREMENT PatternFlowIpv4OptionsTimestampOverflowChoiceEnum
	DECREMENT PatternFlowIpv4OptionsTimestampOverflowChoiceEnum
}{
	VALUE:     PatternFlowIpv4OptionsTimestampOverflowChoiceEnum("value"),
	VALUES:    PatternFlowIpv4OptionsTimestampOverflowChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4OptionsTimestampOverflowChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4OptionsTimestampOverflowChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4OptionsTimestampOverflow) Choice() PatternFlowIpv4OptionsTimestampOverflowChoiceEnum {
	return PatternFlowIpv4OptionsTimestampOverflowChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4OptionsTimestampOverflow) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4OptionsTimestampOverflow) setChoice(value PatternFlowIpv4OptionsTimestampOverflowChoiceEnum) PatternFlowIpv4OptionsTimestampOverflow {
	intValue, ok := otg.PatternFlowIpv4OptionsTimestampOverflow_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4OptionsTimestampOverflowChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4OptionsTimestampOverflow_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4OptionsTimestampOverflowChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampOverflowChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampOverflowChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4OptionsTimestampOverflowCounter().msg()
	}

	if value == PatternFlowIpv4OptionsTimestampOverflowChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4OptionsTimestampOverflowCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsTimestampOverflow) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampOverflowChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsTimestampOverflow) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4OptionsTimestampOverflow object
func (obj *patternFlowIpv4OptionsTimestampOverflow) SetValue(value uint32) PatternFlowIpv4OptionsTimestampOverflow {
	obj.setChoice(PatternFlowIpv4OptionsTimestampOverflowChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4OptionsTimestampOverflow) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4OptionsTimestampOverflow object
func (obj *patternFlowIpv4OptionsTimestampOverflow) SetValues(value []uint32) PatternFlowIpv4OptionsTimestampOverflow {
	obj.setChoice(PatternFlowIpv4OptionsTimestampOverflowChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampOverflowCounter
func (obj *patternFlowIpv4OptionsTimestampOverflow) Increment() PatternFlowIpv4OptionsTimestampOverflowCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampOverflowChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4OptionsTimestampOverflowCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampOverflowCounter
func (obj *patternFlowIpv4OptionsTimestampOverflow) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4OptionsTimestampOverflowCounter value in the PatternFlowIpv4OptionsTimestampOverflow object
func (obj *patternFlowIpv4OptionsTimestampOverflow) SetIncrement(value PatternFlowIpv4OptionsTimestampOverflowCounter) PatternFlowIpv4OptionsTimestampOverflow {
	obj.setChoice(PatternFlowIpv4OptionsTimestampOverflowChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampOverflowCounter
func (obj *patternFlowIpv4OptionsTimestampOverflow) Decrement() PatternFlowIpv4OptionsTimestampOverflowCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampOverflowChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4OptionsTimestampOverflowCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampOverflowCounter
func (obj *patternFlowIpv4OptionsTimestampOverflow) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4OptionsTimestampOverflowCounter value in the PatternFlowIpv4OptionsTimestampOverflow object
func (obj *patternFlowIpv4OptionsTimestampOverflow) SetDecrement(value PatternFlowIpv4OptionsTimestampOverflowCounter) PatternFlowIpv4OptionsTimestampOverflow {
	obj.setChoice(PatternFlowIpv4OptionsTimestampOverflowChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv4OptionsTimestampOverflow) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsTimestampOverflow.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 15 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4OptionsTimestampOverflow.Values <= 15 but Got %d", item))
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

func (obj *patternFlowIpv4OptionsTimestampOverflow) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4OptionsTimestampOverflowChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampOverflowChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampOverflowChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampOverflowChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampOverflowChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4OptionsTimestampOverflowChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4OptionsTimestampOverflow")
			}
		} else {
			intVal := otg.PatternFlowIpv4OptionsTimestampOverflow_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4OptionsTimestampOverflow_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
