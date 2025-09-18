package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateActivity *****
type patternFlowLacpduActorStateActivity struct {
	validation
	obj             *otg.PatternFlowLacpduActorStateActivity
	marshaller      marshalPatternFlowLacpduActorStateActivity
	unMarshaller    unMarshalPatternFlowLacpduActorStateActivity
	incrementHolder PatternFlowLacpduActorStateActivityCounter
	decrementHolder PatternFlowLacpduActorStateActivityCounter
}

func NewPatternFlowLacpduActorStateActivity() PatternFlowLacpduActorStateActivity {
	obj := patternFlowLacpduActorStateActivity{obj: &otg.PatternFlowLacpduActorStateActivity{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateActivity) msg() *otg.PatternFlowLacpduActorStateActivity {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateActivity) setMsg(msg *otg.PatternFlowLacpduActorStateActivity) PatternFlowLacpduActorStateActivity {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateActivity struct {
	obj *patternFlowLacpduActorStateActivity
}

type marshalPatternFlowLacpduActorStateActivity interface {
	// ToProto marshals PatternFlowLacpduActorStateActivity to protobuf object *otg.PatternFlowLacpduActorStateActivity
	ToProto() (*otg.PatternFlowLacpduActorStateActivity, error)
	// ToPbText marshals PatternFlowLacpduActorStateActivity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateActivity to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateActivity to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateActivity struct {
	obj *patternFlowLacpduActorStateActivity
}

type unMarshalPatternFlowLacpduActorStateActivity interface {
	// FromProto unmarshals PatternFlowLacpduActorStateActivity from protobuf object *otg.PatternFlowLacpduActorStateActivity
	FromProto(msg *otg.PatternFlowLacpduActorStateActivity) (PatternFlowLacpduActorStateActivity, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateActivity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateActivity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateActivity from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateActivity) Marshal() marshalPatternFlowLacpduActorStateActivity {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateActivity{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateActivity) Unmarshal() unMarshalPatternFlowLacpduActorStateActivity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateActivity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateActivity) ToProto() (*otg.PatternFlowLacpduActorStateActivity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateActivity) FromProto(msg *otg.PatternFlowLacpduActorStateActivity) (PatternFlowLacpduActorStateActivity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateActivity) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateActivity) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateActivity) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateActivity) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateActivity) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateActivity) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateActivity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateActivity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateActivity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateActivity) Clone() (PatternFlowLacpduActorStateActivity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateActivity()
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

func (obj *patternFlowLacpduActorStateActivity) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorStateActivity is lACP Activity (1=Active, 0=Passive)
type PatternFlowLacpduActorStateActivity interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateActivity to protobuf object *otg.PatternFlowLacpduActorStateActivity
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateActivity
	// setMsg unmarshals PatternFlowLacpduActorStateActivity from protobuf object *otg.PatternFlowLacpduActorStateActivity
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateActivity) PatternFlowLacpduActorStateActivity
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateActivity
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateActivity
	// validate validates PatternFlowLacpduActorStateActivity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateActivity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorStateActivityChoiceEnum, set in PatternFlowLacpduActorStateActivity
	Choice() PatternFlowLacpduActorStateActivityChoiceEnum
	// setChoice assigns PatternFlowLacpduActorStateActivityChoiceEnum provided by user to PatternFlowLacpduActorStateActivity
	setChoice(value PatternFlowLacpduActorStateActivityChoiceEnum) PatternFlowLacpduActorStateActivity
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorStateActivity
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorStateActivity.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorStateActivity
	SetValue(value uint32) PatternFlowLacpduActorStateActivity
	// HasValue checks if Value has been set in PatternFlowLacpduActorStateActivity
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorStateActivity.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorStateActivity
	SetValues(value []uint32) PatternFlowLacpduActorStateActivity
	// Increment returns PatternFlowLacpduActorStateActivityCounter, set in PatternFlowLacpduActorStateActivity.
	// PatternFlowLacpduActorStateActivityCounter is integer counter pattern
	Increment() PatternFlowLacpduActorStateActivityCounter
	// SetIncrement assigns PatternFlowLacpduActorStateActivityCounter provided by user to PatternFlowLacpduActorStateActivity.
	// PatternFlowLacpduActorStateActivityCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorStateActivityCounter) PatternFlowLacpduActorStateActivity
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorStateActivity
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorStateActivityCounter, set in PatternFlowLacpduActorStateActivity.
	// PatternFlowLacpduActorStateActivityCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorStateActivityCounter
	// SetDecrement assigns PatternFlowLacpduActorStateActivityCounter provided by user to PatternFlowLacpduActorStateActivity.
	// PatternFlowLacpduActorStateActivityCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorStateActivityCounter) PatternFlowLacpduActorStateActivity
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorStateActivity
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorStateActivityChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorStateActivity
var PatternFlowLacpduActorStateActivityChoice = struct {
	VALUE     PatternFlowLacpduActorStateActivityChoiceEnum
	VALUES    PatternFlowLacpduActorStateActivityChoiceEnum
	INCREMENT PatternFlowLacpduActorStateActivityChoiceEnum
	DECREMENT PatternFlowLacpduActorStateActivityChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorStateActivityChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorStateActivityChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorStateActivityChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorStateActivityChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorStateActivity) Choice() PatternFlowLacpduActorStateActivityChoiceEnum {
	return PatternFlowLacpduActorStateActivityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorStateActivity) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorStateActivity) setChoice(value PatternFlowLacpduActorStateActivityChoiceEnum) PatternFlowLacpduActorStateActivity {
	intValue, ok := otg.PatternFlowLacpduActorStateActivity_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorStateActivityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorStateActivity_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorStateActivityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorStateActivityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorStateActivityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorStateActivityCounter().msg()
	}

	if value == PatternFlowLacpduActorStateActivityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorStateActivityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateActivity) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorStateActivityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateActivity) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorStateActivity object
func (obj *patternFlowLacpduActorStateActivity) SetValue(value uint32) PatternFlowLacpduActorStateActivity {
	obj.setChoice(PatternFlowLacpduActorStateActivityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorStateActivity) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorStateActivity object
func (obj *patternFlowLacpduActorStateActivity) SetValues(value []uint32) PatternFlowLacpduActorStateActivity {
	obj.setChoice(PatternFlowLacpduActorStateActivityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateActivityCounter
func (obj *patternFlowLacpduActorStateActivity) Increment() PatternFlowLacpduActorStateActivityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorStateActivityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorStateActivityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateActivityCounter
func (obj *patternFlowLacpduActorStateActivity) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorStateActivityCounter value in the PatternFlowLacpduActorStateActivity object
func (obj *patternFlowLacpduActorStateActivity) SetIncrement(value PatternFlowLacpduActorStateActivityCounter) PatternFlowLacpduActorStateActivity {
	obj.setChoice(PatternFlowLacpduActorStateActivityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateActivityCounter
func (obj *patternFlowLacpduActorStateActivity) Decrement() PatternFlowLacpduActorStateActivityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorStateActivityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorStateActivityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateActivityCounter
func (obj *patternFlowLacpduActorStateActivity) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorStateActivityCounter value in the PatternFlowLacpduActorStateActivity object
func (obj *patternFlowLacpduActorStateActivity) SetDecrement(value PatternFlowLacpduActorStateActivityCounter) PatternFlowLacpduActorStateActivity {
	obj.setChoice(PatternFlowLacpduActorStateActivityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorStateActivity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateActivity.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorStateActivity.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduActorStateActivity) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorStateActivityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateActivityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorStateActivityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateActivityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateActivityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorStateActivityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorStateActivity")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorStateActivity_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorStateActivity_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
