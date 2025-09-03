package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateActivity *****
type patternFlowLacpPartnerStateActivity struct {
	validation
	obj             *otg.PatternFlowLacpPartnerStateActivity
	marshaller      marshalPatternFlowLacpPartnerStateActivity
	unMarshaller    unMarshalPatternFlowLacpPartnerStateActivity
	incrementHolder PatternFlowLacpPartnerStateActivityCounter
	decrementHolder PatternFlowLacpPartnerStateActivityCounter
}

func NewPatternFlowLacpPartnerStateActivity() PatternFlowLacpPartnerStateActivity {
	obj := patternFlowLacpPartnerStateActivity{obj: &otg.PatternFlowLacpPartnerStateActivity{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateActivity) msg() *otg.PatternFlowLacpPartnerStateActivity {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateActivity) setMsg(msg *otg.PatternFlowLacpPartnerStateActivity) PatternFlowLacpPartnerStateActivity {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateActivity struct {
	obj *patternFlowLacpPartnerStateActivity
}

type marshalPatternFlowLacpPartnerStateActivity interface {
	// ToProto marshals PatternFlowLacpPartnerStateActivity to protobuf object *otg.PatternFlowLacpPartnerStateActivity
	ToProto() (*otg.PatternFlowLacpPartnerStateActivity, error)
	// ToPbText marshals PatternFlowLacpPartnerStateActivity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateActivity to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateActivity to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateActivity struct {
	obj *patternFlowLacpPartnerStateActivity
}

type unMarshalPatternFlowLacpPartnerStateActivity interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateActivity from protobuf object *otg.PatternFlowLacpPartnerStateActivity
	FromProto(msg *otg.PatternFlowLacpPartnerStateActivity) (PatternFlowLacpPartnerStateActivity, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateActivity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateActivity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateActivity from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateActivity) Marshal() marshalPatternFlowLacpPartnerStateActivity {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateActivity{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateActivity) Unmarshal() unMarshalPatternFlowLacpPartnerStateActivity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateActivity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateActivity) ToProto() (*otg.PatternFlowLacpPartnerStateActivity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateActivity) FromProto(msg *otg.PatternFlowLacpPartnerStateActivity) (PatternFlowLacpPartnerStateActivity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateActivity) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateActivity) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateActivity) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateActivity) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateActivity) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateActivity) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateActivity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateActivity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateActivity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateActivity) Clone() (PatternFlowLacpPartnerStateActivity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateActivity()
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

func (obj *patternFlowLacpPartnerStateActivity) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerStateActivity is lACP Activity (1=Active, 0=Passive)
type PatternFlowLacpPartnerStateActivity interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateActivity to protobuf object *otg.PatternFlowLacpPartnerStateActivity
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateActivity
	// setMsg unmarshals PatternFlowLacpPartnerStateActivity from protobuf object *otg.PatternFlowLacpPartnerStateActivity
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateActivity) PatternFlowLacpPartnerStateActivity
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateActivity
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateActivity
	// validate validates PatternFlowLacpPartnerStateActivity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateActivity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerStateActivityChoiceEnum, set in PatternFlowLacpPartnerStateActivity
	Choice() PatternFlowLacpPartnerStateActivityChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerStateActivityChoiceEnum provided by user to PatternFlowLacpPartnerStateActivity
	setChoice(value PatternFlowLacpPartnerStateActivityChoiceEnum) PatternFlowLacpPartnerStateActivity
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerStateActivity
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerStateActivity.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerStateActivity
	SetValue(value uint32) PatternFlowLacpPartnerStateActivity
	// HasValue checks if Value has been set in PatternFlowLacpPartnerStateActivity
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerStateActivity.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerStateActivity
	SetValues(value []uint32) PatternFlowLacpPartnerStateActivity
	// Increment returns PatternFlowLacpPartnerStateActivityCounter, set in PatternFlowLacpPartnerStateActivity.
	// PatternFlowLacpPartnerStateActivityCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerStateActivityCounter
	// SetIncrement assigns PatternFlowLacpPartnerStateActivityCounter provided by user to PatternFlowLacpPartnerStateActivity.
	// PatternFlowLacpPartnerStateActivityCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerStateActivityCounter) PatternFlowLacpPartnerStateActivity
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerStateActivity
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerStateActivityCounter, set in PatternFlowLacpPartnerStateActivity.
	// PatternFlowLacpPartnerStateActivityCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerStateActivityCounter
	// SetDecrement assigns PatternFlowLacpPartnerStateActivityCounter provided by user to PatternFlowLacpPartnerStateActivity.
	// PatternFlowLacpPartnerStateActivityCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerStateActivityCounter) PatternFlowLacpPartnerStateActivity
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerStateActivity
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerStateActivityChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerStateActivity
var PatternFlowLacpPartnerStateActivityChoice = struct {
	VALUE     PatternFlowLacpPartnerStateActivityChoiceEnum
	VALUES    PatternFlowLacpPartnerStateActivityChoiceEnum
	INCREMENT PatternFlowLacpPartnerStateActivityChoiceEnum
	DECREMENT PatternFlowLacpPartnerStateActivityChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerStateActivityChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerStateActivityChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerStateActivityChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerStateActivityChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerStateActivity) Choice() PatternFlowLacpPartnerStateActivityChoiceEnum {
	return PatternFlowLacpPartnerStateActivityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerStateActivity) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerStateActivity) setChoice(value PatternFlowLacpPartnerStateActivityChoiceEnum) PatternFlowLacpPartnerStateActivity {
	intValue, ok := otg.PatternFlowLacpPartnerStateActivity_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerStateActivityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerStateActivity_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerStateActivityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerStateActivityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerStateActivityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerStateActivityCounter().msg()
	}

	if value == PatternFlowLacpPartnerStateActivityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerStateActivityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateActivity) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerStateActivityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateActivity) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerStateActivity object
func (obj *patternFlowLacpPartnerStateActivity) SetValue(value uint32) PatternFlowLacpPartnerStateActivity {
	obj.setChoice(PatternFlowLacpPartnerStateActivityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerStateActivity) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerStateActivity object
func (obj *patternFlowLacpPartnerStateActivity) SetValues(value []uint32) PatternFlowLacpPartnerStateActivity {
	obj.setChoice(PatternFlowLacpPartnerStateActivityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateActivityCounter
func (obj *patternFlowLacpPartnerStateActivity) Increment() PatternFlowLacpPartnerStateActivityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerStateActivityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerStateActivityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateActivityCounter
func (obj *patternFlowLacpPartnerStateActivity) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerStateActivityCounter value in the PatternFlowLacpPartnerStateActivity object
func (obj *patternFlowLacpPartnerStateActivity) SetIncrement(value PatternFlowLacpPartnerStateActivityCounter) PatternFlowLacpPartnerStateActivity {
	obj.setChoice(PatternFlowLacpPartnerStateActivityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateActivityCounter
func (obj *patternFlowLacpPartnerStateActivity) Decrement() PatternFlowLacpPartnerStateActivityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerStateActivityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerStateActivityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateActivityCounter
func (obj *patternFlowLacpPartnerStateActivity) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerStateActivityCounter value in the PatternFlowLacpPartnerStateActivity object
func (obj *patternFlowLacpPartnerStateActivity) SetDecrement(value PatternFlowLacpPartnerStateActivityCounter) PatternFlowLacpPartnerStateActivity {
	obj.setChoice(PatternFlowLacpPartnerStateActivityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerStateActivity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateActivity.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerStateActivity.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpPartnerStateActivity) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerStateActivityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateActivityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateActivityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateActivityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateActivityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerStateActivityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerStateActivity")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerStateActivity_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerStateActivity_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
