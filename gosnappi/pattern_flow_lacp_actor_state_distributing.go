package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateDistributing *****
type patternFlowLacpActorStateDistributing struct {
	validation
	obj             *otg.PatternFlowLacpActorStateDistributing
	marshaller      marshalPatternFlowLacpActorStateDistributing
	unMarshaller    unMarshalPatternFlowLacpActorStateDistributing
	incrementHolder PatternFlowLacpActorStateDistributingCounter
	decrementHolder PatternFlowLacpActorStateDistributingCounter
}

func NewPatternFlowLacpActorStateDistributing() PatternFlowLacpActorStateDistributing {
	obj := patternFlowLacpActorStateDistributing{obj: &otg.PatternFlowLacpActorStateDistributing{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateDistributing) msg() *otg.PatternFlowLacpActorStateDistributing {
	return obj.obj
}

func (obj *patternFlowLacpActorStateDistributing) setMsg(msg *otg.PatternFlowLacpActorStateDistributing) PatternFlowLacpActorStateDistributing {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateDistributing struct {
	obj *patternFlowLacpActorStateDistributing
}

type marshalPatternFlowLacpActorStateDistributing interface {
	// ToProto marshals PatternFlowLacpActorStateDistributing to protobuf object *otg.PatternFlowLacpActorStateDistributing
	ToProto() (*otg.PatternFlowLacpActorStateDistributing, error)
	// ToPbText marshals PatternFlowLacpActorStateDistributing to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateDistributing to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateDistributing to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateDistributing struct {
	obj *patternFlowLacpActorStateDistributing
}

type unMarshalPatternFlowLacpActorStateDistributing interface {
	// FromProto unmarshals PatternFlowLacpActorStateDistributing from protobuf object *otg.PatternFlowLacpActorStateDistributing
	FromProto(msg *otg.PatternFlowLacpActorStateDistributing) (PatternFlowLacpActorStateDistributing, error)
	// FromPbText unmarshals PatternFlowLacpActorStateDistributing from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateDistributing from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateDistributing from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateDistributing) Marshal() marshalPatternFlowLacpActorStateDistributing {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateDistributing{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateDistributing) Unmarshal() unMarshalPatternFlowLacpActorStateDistributing {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateDistributing{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateDistributing) ToProto() (*otg.PatternFlowLacpActorStateDistributing, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateDistributing) FromProto(msg *otg.PatternFlowLacpActorStateDistributing) (PatternFlowLacpActorStateDistributing, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateDistributing) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateDistributing) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateDistributing) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateDistributing) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateDistributing) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateDistributing) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateDistributing) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateDistributing) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateDistributing) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateDistributing) Clone() (PatternFlowLacpActorStateDistributing, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateDistributing()
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

func (obj *patternFlowLacpActorStateDistributing) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorStateDistributing is distributing (1=Enabled, 0=Disabled)
type PatternFlowLacpActorStateDistributing interface {
	Validation
	// msg marshals PatternFlowLacpActorStateDistributing to protobuf object *otg.PatternFlowLacpActorStateDistributing
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateDistributing
	// setMsg unmarshals PatternFlowLacpActorStateDistributing from protobuf object *otg.PatternFlowLacpActorStateDistributing
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateDistributing) PatternFlowLacpActorStateDistributing
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateDistributing
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateDistributing
	// validate validates PatternFlowLacpActorStateDistributing
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateDistributing, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorStateDistributingChoiceEnum, set in PatternFlowLacpActorStateDistributing
	Choice() PatternFlowLacpActorStateDistributingChoiceEnum
	// setChoice assigns PatternFlowLacpActorStateDistributingChoiceEnum provided by user to PatternFlowLacpActorStateDistributing
	setChoice(value PatternFlowLacpActorStateDistributingChoiceEnum) PatternFlowLacpActorStateDistributing
	// HasChoice checks if Choice has been set in PatternFlowLacpActorStateDistributing
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorStateDistributing.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorStateDistributing
	SetValue(value uint32) PatternFlowLacpActorStateDistributing
	// HasValue checks if Value has been set in PatternFlowLacpActorStateDistributing
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorStateDistributing.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorStateDistributing
	SetValues(value []uint32) PatternFlowLacpActorStateDistributing
	// Increment returns PatternFlowLacpActorStateDistributingCounter, set in PatternFlowLacpActorStateDistributing.
	// PatternFlowLacpActorStateDistributingCounter is integer counter pattern
	Increment() PatternFlowLacpActorStateDistributingCounter
	// SetIncrement assigns PatternFlowLacpActorStateDistributingCounter provided by user to PatternFlowLacpActorStateDistributing.
	// PatternFlowLacpActorStateDistributingCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorStateDistributingCounter) PatternFlowLacpActorStateDistributing
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorStateDistributing
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorStateDistributingCounter, set in PatternFlowLacpActorStateDistributing.
	// PatternFlowLacpActorStateDistributingCounter is integer counter pattern
	Decrement() PatternFlowLacpActorStateDistributingCounter
	// SetDecrement assigns PatternFlowLacpActorStateDistributingCounter provided by user to PatternFlowLacpActorStateDistributing.
	// PatternFlowLacpActorStateDistributingCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorStateDistributingCounter) PatternFlowLacpActorStateDistributing
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorStateDistributing
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorStateDistributingChoiceEnum string

// Enum of Choice on PatternFlowLacpActorStateDistributing
var PatternFlowLacpActorStateDistributingChoice = struct {
	VALUE     PatternFlowLacpActorStateDistributingChoiceEnum
	VALUES    PatternFlowLacpActorStateDistributingChoiceEnum
	INCREMENT PatternFlowLacpActorStateDistributingChoiceEnum
	DECREMENT PatternFlowLacpActorStateDistributingChoiceEnum
}{
	VALUE:     PatternFlowLacpActorStateDistributingChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorStateDistributingChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorStateDistributingChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorStateDistributingChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorStateDistributing) Choice() PatternFlowLacpActorStateDistributingChoiceEnum {
	return PatternFlowLacpActorStateDistributingChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorStateDistributing) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorStateDistributing) setChoice(value PatternFlowLacpActorStateDistributingChoiceEnum) PatternFlowLacpActorStateDistributing {
	intValue, ok := otg.PatternFlowLacpActorStateDistributing_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorStateDistributingChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorStateDistributing_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorStateDistributingChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorStateDistributingChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorStateDistributingChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorStateDistributingCounter().msg()
	}

	if value == PatternFlowLacpActorStateDistributingChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorStateDistributingCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateDistributing) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorStateDistributingChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateDistributing) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorStateDistributing object
func (obj *patternFlowLacpActorStateDistributing) SetValue(value uint32) PatternFlowLacpActorStateDistributing {
	obj.setChoice(PatternFlowLacpActorStateDistributingChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorStateDistributing) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorStateDistributing object
func (obj *patternFlowLacpActorStateDistributing) SetValues(value []uint32) PatternFlowLacpActorStateDistributing {
	obj.setChoice(PatternFlowLacpActorStateDistributingChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateDistributingCounter
func (obj *patternFlowLacpActorStateDistributing) Increment() PatternFlowLacpActorStateDistributingCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorStateDistributingChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorStateDistributingCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateDistributingCounter
func (obj *patternFlowLacpActorStateDistributing) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorStateDistributingCounter value in the PatternFlowLacpActorStateDistributing object
func (obj *patternFlowLacpActorStateDistributing) SetIncrement(value PatternFlowLacpActorStateDistributingCounter) PatternFlowLacpActorStateDistributing {
	obj.setChoice(PatternFlowLacpActorStateDistributingChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateDistributingCounter
func (obj *patternFlowLacpActorStateDistributing) Decrement() PatternFlowLacpActorStateDistributingCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorStateDistributingChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorStateDistributingCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateDistributingCounter
func (obj *patternFlowLacpActorStateDistributing) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorStateDistributingCounter value in the PatternFlowLacpActorStateDistributing object
func (obj *patternFlowLacpActorStateDistributing) SetDecrement(value PatternFlowLacpActorStateDistributingCounter) PatternFlowLacpActorStateDistributing {
	obj.setChoice(PatternFlowLacpActorStateDistributingChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorStateDistributing) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateDistributing.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorStateDistributing.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpActorStateDistributing) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorStateDistributingChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateDistributingChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorStateDistributingChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateDistributingChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateDistributingChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorStateDistributingChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorStateDistributing")
			}
		} else {
			intVal := otg.PatternFlowLacpActorStateDistributing_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorStateDistributing_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
