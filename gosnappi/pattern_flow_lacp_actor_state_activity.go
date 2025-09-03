package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateActivity *****
type patternFlowLacpActorStateActivity struct {
	validation
	obj             *otg.PatternFlowLacpActorStateActivity
	marshaller      marshalPatternFlowLacpActorStateActivity
	unMarshaller    unMarshalPatternFlowLacpActorStateActivity
	incrementHolder PatternFlowLacpActorStateActivityCounter
	decrementHolder PatternFlowLacpActorStateActivityCounter
}

func NewPatternFlowLacpActorStateActivity() PatternFlowLacpActorStateActivity {
	obj := patternFlowLacpActorStateActivity{obj: &otg.PatternFlowLacpActorStateActivity{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateActivity) msg() *otg.PatternFlowLacpActorStateActivity {
	return obj.obj
}

func (obj *patternFlowLacpActorStateActivity) setMsg(msg *otg.PatternFlowLacpActorStateActivity) PatternFlowLacpActorStateActivity {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateActivity struct {
	obj *patternFlowLacpActorStateActivity
}

type marshalPatternFlowLacpActorStateActivity interface {
	// ToProto marshals PatternFlowLacpActorStateActivity to protobuf object *otg.PatternFlowLacpActorStateActivity
	ToProto() (*otg.PatternFlowLacpActorStateActivity, error)
	// ToPbText marshals PatternFlowLacpActorStateActivity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateActivity to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateActivity to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateActivity struct {
	obj *patternFlowLacpActorStateActivity
}

type unMarshalPatternFlowLacpActorStateActivity interface {
	// FromProto unmarshals PatternFlowLacpActorStateActivity from protobuf object *otg.PatternFlowLacpActorStateActivity
	FromProto(msg *otg.PatternFlowLacpActorStateActivity) (PatternFlowLacpActorStateActivity, error)
	// FromPbText unmarshals PatternFlowLacpActorStateActivity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateActivity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateActivity from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateActivity) Marshal() marshalPatternFlowLacpActorStateActivity {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateActivity{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateActivity) Unmarshal() unMarshalPatternFlowLacpActorStateActivity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateActivity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateActivity) ToProto() (*otg.PatternFlowLacpActorStateActivity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateActivity) FromProto(msg *otg.PatternFlowLacpActorStateActivity) (PatternFlowLacpActorStateActivity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateActivity) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateActivity) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateActivity) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateActivity) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateActivity) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateActivity) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateActivity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateActivity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateActivity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateActivity) Clone() (PatternFlowLacpActorStateActivity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateActivity()
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

func (obj *patternFlowLacpActorStateActivity) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorStateActivity is lACP Activity (1=Active, 0=Passive)
type PatternFlowLacpActorStateActivity interface {
	Validation
	// msg marshals PatternFlowLacpActorStateActivity to protobuf object *otg.PatternFlowLacpActorStateActivity
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateActivity
	// setMsg unmarshals PatternFlowLacpActorStateActivity from protobuf object *otg.PatternFlowLacpActorStateActivity
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateActivity) PatternFlowLacpActorStateActivity
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateActivity
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateActivity
	// validate validates PatternFlowLacpActorStateActivity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateActivity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorStateActivityChoiceEnum, set in PatternFlowLacpActorStateActivity
	Choice() PatternFlowLacpActorStateActivityChoiceEnum
	// setChoice assigns PatternFlowLacpActorStateActivityChoiceEnum provided by user to PatternFlowLacpActorStateActivity
	setChoice(value PatternFlowLacpActorStateActivityChoiceEnum) PatternFlowLacpActorStateActivity
	// HasChoice checks if Choice has been set in PatternFlowLacpActorStateActivity
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorStateActivity.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorStateActivity
	SetValue(value uint32) PatternFlowLacpActorStateActivity
	// HasValue checks if Value has been set in PatternFlowLacpActorStateActivity
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorStateActivity.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorStateActivity
	SetValues(value []uint32) PatternFlowLacpActorStateActivity
	// Increment returns PatternFlowLacpActorStateActivityCounter, set in PatternFlowLacpActorStateActivity.
	// PatternFlowLacpActorStateActivityCounter is integer counter pattern
	Increment() PatternFlowLacpActorStateActivityCounter
	// SetIncrement assigns PatternFlowLacpActorStateActivityCounter provided by user to PatternFlowLacpActorStateActivity.
	// PatternFlowLacpActorStateActivityCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorStateActivityCounter) PatternFlowLacpActorStateActivity
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorStateActivity
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorStateActivityCounter, set in PatternFlowLacpActorStateActivity.
	// PatternFlowLacpActorStateActivityCounter is integer counter pattern
	Decrement() PatternFlowLacpActorStateActivityCounter
	// SetDecrement assigns PatternFlowLacpActorStateActivityCounter provided by user to PatternFlowLacpActorStateActivity.
	// PatternFlowLacpActorStateActivityCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorStateActivityCounter) PatternFlowLacpActorStateActivity
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorStateActivity
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorStateActivityChoiceEnum string

// Enum of Choice on PatternFlowLacpActorStateActivity
var PatternFlowLacpActorStateActivityChoice = struct {
	VALUE     PatternFlowLacpActorStateActivityChoiceEnum
	VALUES    PatternFlowLacpActorStateActivityChoiceEnum
	INCREMENT PatternFlowLacpActorStateActivityChoiceEnum
	DECREMENT PatternFlowLacpActorStateActivityChoiceEnum
}{
	VALUE:     PatternFlowLacpActorStateActivityChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorStateActivityChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorStateActivityChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorStateActivityChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorStateActivity) Choice() PatternFlowLacpActorStateActivityChoiceEnum {
	return PatternFlowLacpActorStateActivityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorStateActivity) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorStateActivity) setChoice(value PatternFlowLacpActorStateActivityChoiceEnum) PatternFlowLacpActorStateActivity {
	intValue, ok := otg.PatternFlowLacpActorStateActivity_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorStateActivityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorStateActivity_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorStateActivityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorStateActivityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorStateActivityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorStateActivityCounter().msg()
	}

	if value == PatternFlowLacpActorStateActivityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorStateActivityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateActivity) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorStateActivityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateActivity) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorStateActivity object
func (obj *patternFlowLacpActorStateActivity) SetValue(value uint32) PatternFlowLacpActorStateActivity {
	obj.setChoice(PatternFlowLacpActorStateActivityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorStateActivity) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorStateActivity object
func (obj *patternFlowLacpActorStateActivity) SetValues(value []uint32) PatternFlowLacpActorStateActivity {
	obj.setChoice(PatternFlowLacpActorStateActivityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateActivityCounter
func (obj *patternFlowLacpActorStateActivity) Increment() PatternFlowLacpActorStateActivityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorStateActivityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorStateActivityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateActivityCounter
func (obj *patternFlowLacpActorStateActivity) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorStateActivityCounter value in the PatternFlowLacpActorStateActivity object
func (obj *patternFlowLacpActorStateActivity) SetIncrement(value PatternFlowLacpActorStateActivityCounter) PatternFlowLacpActorStateActivity {
	obj.setChoice(PatternFlowLacpActorStateActivityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateActivityCounter
func (obj *patternFlowLacpActorStateActivity) Decrement() PatternFlowLacpActorStateActivityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorStateActivityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorStateActivityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateActivityCounter
func (obj *patternFlowLacpActorStateActivity) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorStateActivityCounter value in the PatternFlowLacpActorStateActivity object
func (obj *patternFlowLacpActorStateActivity) SetDecrement(value PatternFlowLacpActorStateActivityCounter) PatternFlowLacpActorStateActivity {
	obj.setChoice(PatternFlowLacpActorStateActivityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorStateActivity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateActivity.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorStateActivity.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpActorStateActivity) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorStateActivityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateActivityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorStateActivityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateActivityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateActivityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorStateActivityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorStateActivity")
			}
		} else {
			intVal := otg.PatternFlowLacpActorStateActivity_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorStateActivity_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
