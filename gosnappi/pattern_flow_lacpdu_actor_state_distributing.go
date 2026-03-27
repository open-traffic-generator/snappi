package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateDistributing *****
type patternFlowLacpduActorStateDistributing struct {
	validation
	obj             *otg.PatternFlowLacpduActorStateDistributing
	marshaller      marshalPatternFlowLacpduActorStateDistributing
	unMarshaller    unMarshalPatternFlowLacpduActorStateDistributing
	incrementHolder PatternFlowLacpduActorStateDistributingCounter
	decrementHolder PatternFlowLacpduActorStateDistributingCounter
}

func NewPatternFlowLacpduActorStateDistributing() PatternFlowLacpduActorStateDistributing {
	obj := patternFlowLacpduActorStateDistributing{obj: &otg.PatternFlowLacpduActorStateDistributing{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateDistributing) msg() *otg.PatternFlowLacpduActorStateDistributing {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateDistributing) setMsg(msg *otg.PatternFlowLacpduActorStateDistributing) PatternFlowLacpduActorStateDistributing {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateDistributing struct {
	obj *patternFlowLacpduActorStateDistributing
}

type marshalPatternFlowLacpduActorStateDistributing interface {
	// ToProto marshals PatternFlowLacpduActorStateDistributing to protobuf object *otg.PatternFlowLacpduActorStateDistributing
	ToProto() (*otg.PatternFlowLacpduActorStateDistributing, error)
	// ToPbText marshals PatternFlowLacpduActorStateDistributing to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateDistributing to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateDistributing to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateDistributing struct {
	obj *patternFlowLacpduActorStateDistributing
}

type unMarshalPatternFlowLacpduActorStateDistributing interface {
	// FromProto unmarshals PatternFlowLacpduActorStateDistributing from protobuf object *otg.PatternFlowLacpduActorStateDistributing
	FromProto(msg *otg.PatternFlowLacpduActorStateDistributing) (PatternFlowLacpduActorStateDistributing, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateDistributing from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateDistributing from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateDistributing from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateDistributing) Marshal() marshalPatternFlowLacpduActorStateDistributing {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateDistributing{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateDistributing) Unmarshal() unMarshalPatternFlowLacpduActorStateDistributing {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateDistributing{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateDistributing) ToProto() (*otg.PatternFlowLacpduActorStateDistributing, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateDistributing) FromProto(msg *otg.PatternFlowLacpduActorStateDistributing) (PatternFlowLacpduActorStateDistributing, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateDistributing) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateDistributing) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateDistributing) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateDistributing) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateDistributing) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateDistributing) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateDistributing) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateDistributing) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateDistributing) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateDistributing) Clone() (PatternFlowLacpduActorStateDistributing, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateDistributing()
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

func (obj *patternFlowLacpduActorStateDistributing) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorStateDistributing is distributing (1=Enabled, 0=Disabled)
type PatternFlowLacpduActorStateDistributing interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateDistributing to protobuf object *otg.PatternFlowLacpduActorStateDistributing
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateDistributing
	// setMsg unmarshals PatternFlowLacpduActorStateDistributing from protobuf object *otg.PatternFlowLacpduActorStateDistributing
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateDistributing) PatternFlowLacpduActorStateDistributing
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateDistributing
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateDistributing
	// validate validates PatternFlowLacpduActorStateDistributing
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateDistributing, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorStateDistributingChoiceEnum, set in PatternFlowLacpduActorStateDistributing
	Choice() PatternFlowLacpduActorStateDistributingChoiceEnum
	// setChoice assigns PatternFlowLacpduActorStateDistributingChoiceEnum provided by user to PatternFlowLacpduActorStateDistributing
	setChoice(value PatternFlowLacpduActorStateDistributingChoiceEnum) PatternFlowLacpduActorStateDistributing
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorStateDistributing
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorStateDistributing.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorStateDistributing
	SetValue(value uint32) PatternFlowLacpduActorStateDistributing
	// HasValue checks if Value has been set in PatternFlowLacpduActorStateDistributing
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorStateDistributing.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorStateDistributing
	SetValues(value []uint32) PatternFlowLacpduActorStateDistributing
	// Increment returns PatternFlowLacpduActorStateDistributingCounter, set in PatternFlowLacpduActorStateDistributing.
	// PatternFlowLacpduActorStateDistributingCounter is integer counter pattern
	Increment() PatternFlowLacpduActorStateDistributingCounter
	// SetIncrement assigns PatternFlowLacpduActorStateDistributingCounter provided by user to PatternFlowLacpduActorStateDistributing.
	// PatternFlowLacpduActorStateDistributingCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorStateDistributingCounter) PatternFlowLacpduActorStateDistributing
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorStateDistributing
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorStateDistributingCounter, set in PatternFlowLacpduActorStateDistributing.
	// PatternFlowLacpduActorStateDistributingCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorStateDistributingCounter
	// SetDecrement assigns PatternFlowLacpduActorStateDistributingCounter provided by user to PatternFlowLacpduActorStateDistributing.
	// PatternFlowLacpduActorStateDistributingCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorStateDistributingCounter) PatternFlowLacpduActorStateDistributing
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorStateDistributing
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorStateDistributingChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorStateDistributing
var PatternFlowLacpduActorStateDistributingChoice = struct {
	VALUE     PatternFlowLacpduActorStateDistributingChoiceEnum
	VALUES    PatternFlowLacpduActorStateDistributingChoiceEnum
	INCREMENT PatternFlowLacpduActorStateDistributingChoiceEnum
	DECREMENT PatternFlowLacpduActorStateDistributingChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorStateDistributingChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorStateDistributingChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorStateDistributingChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorStateDistributingChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorStateDistributing) Choice() PatternFlowLacpduActorStateDistributingChoiceEnum {
	return PatternFlowLacpduActorStateDistributingChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorStateDistributing) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorStateDistributing) setChoice(value PatternFlowLacpduActorStateDistributingChoiceEnum) PatternFlowLacpduActorStateDistributing {
	intValue, ok := otg.PatternFlowLacpduActorStateDistributing_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorStateDistributingChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorStateDistributing_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorStateDistributingChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorStateDistributingChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorStateDistributingChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorStateDistributingCounter().msg()
	}

	if value == PatternFlowLacpduActorStateDistributingChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorStateDistributingCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateDistributing) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorStateDistributingChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateDistributing) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorStateDistributing object
func (obj *patternFlowLacpduActorStateDistributing) SetValue(value uint32) PatternFlowLacpduActorStateDistributing {
	obj.setChoice(PatternFlowLacpduActorStateDistributingChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorStateDistributing) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorStateDistributing object
func (obj *patternFlowLacpduActorStateDistributing) SetValues(value []uint32) PatternFlowLacpduActorStateDistributing {
	obj.setChoice(PatternFlowLacpduActorStateDistributingChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateDistributingCounter
func (obj *patternFlowLacpduActorStateDistributing) Increment() PatternFlowLacpduActorStateDistributingCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorStateDistributingChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorStateDistributingCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateDistributingCounter
func (obj *patternFlowLacpduActorStateDistributing) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorStateDistributingCounter value in the PatternFlowLacpduActorStateDistributing object
func (obj *patternFlowLacpduActorStateDistributing) SetIncrement(value PatternFlowLacpduActorStateDistributingCounter) PatternFlowLacpduActorStateDistributing {
	obj.setChoice(PatternFlowLacpduActorStateDistributingChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateDistributingCounter
func (obj *patternFlowLacpduActorStateDistributing) Decrement() PatternFlowLacpduActorStateDistributingCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorStateDistributingChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorStateDistributingCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateDistributingCounter
func (obj *patternFlowLacpduActorStateDistributing) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorStateDistributingCounter value in the PatternFlowLacpduActorStateDistributing object
func (obj *patternFlowLacpduActorStateDistributing) SetDecrement(value PatternFlowLacpduActorStateDistributingCounter) PatternFlowLacpduActorStateDistributing {
	obj.setChoice(PatternFlowLacpduActorStateDistributingChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorStateDistributing) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateDistributing.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorStateDistributing.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduActorStateDistributing) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorStateDistributingChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateDistributingChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorStateDistributingChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateDistributingChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateDistributingChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorStateDistributingChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorStateDistributing")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorStateDistributing_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorStateDistributing_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
