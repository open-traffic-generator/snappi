package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateActivity *****
type patternFlowLacpduPartnerStateActivity struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerStateActivity
	marshaller      marshalPatternFlowLacpduPartnerStateActivity
	unMarshaller    unMarshalPatternFlowLacpduPartnerStateActivity
	incrementHolder PatternFlowLacpduPartnerStateActivityCounter
	decrementHolder PatternFlowLacpduPartnerStateActivityCounter
}

func NewPatternFlowLacpduPartnerStateActivity() PatternFlowLacpduPartnerStateActivity {
	obj := patternFlowLacpduPartnerStateActivity{obj: &otg.PatternFlowLacpduPartnerStateActivity{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateActivity) msg() *otg.PatternFlowLacpduPartnerStateActivity {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateActivity) setMsg(msg *otg.PatternFlowLacpduPartnerStateActivity) PatternFlowLacpduPartnerStateActivity {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateActivity struct {
	obj *patternFlowLacpduPartnerStateActivity
}

type marshalPatternFlowLacpduPartnerStateActivity interface {
	// ToProto marshals PatternFlowLacpduPartnerStateActivity to protobuf object *otg.PatternFlowLacpduPartnerStateActivity
	ToProto() (*otg.PatternFlowLacpduPartnerStateActivity, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateActivity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateActivity to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateActivity to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateActivity struct {
	obj *patternFlowLacpduPartnerStateActivity
}

type unMarshalPatternFlowLacpduPartnerStateActivity interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateActivity from protobuf object *otg.PatternFlowLacpduPartnerStateActivity
	FromProto(msg *otg.PatternFlowLacpduPartnerStateActivity) (PatternFlowLacpduPartnerStateActivity, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateActivity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateActivity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateActivity from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateActivity) Marshal() marshalPatternFlowLacpduPartnerStateActivity {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateActivity{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateActivity) Unmarshal() unMarshalPatternFlowLacpduPartnerStateActivity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateActivity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateActivity) ToProto() (*otg.PatternFlowLacpduPartnerStateActivity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateActivity) FromProto(msg *otg.PatternFlowLacpduPartnerStateActivity) (PatternFlowLacpduPartnerStateActivity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateActivity) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateActivity) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateActivity) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateActivity) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateActivity) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateActivity) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateActivity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateActivity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateActivity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateActivity) Clone() (PatternFlowLacpduPartnerStateActivity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateActivity()
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

func (obj *patternFlowLacpduPartnerStateActivity) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerStateActivity is lACP Activity (1=Active, 0=Passive)
type PatternFlowLacpduPartnerStateActivity interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateActivity to protobuf object *otg.PatternFlowLacpduPartnerStateActivity
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateActivity
	// setMsg unmarshals PatternFlowLacpduPartnerStateActivity from protobuf object *otg.PatternFlowLacpduPartnerStateActivity
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateActivity) PatternFlowLacpduPartnerStateActivity
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateActivity
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateActivity
	// validate validates PatternFlowLacpduPartnerStateActivity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateActivity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerStateActivityChoiceEnum, set in PatternFlowLacpduPartnerStateActivity
	Choice() PatternFlowLacpduPartnerStateActivityChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerStateActivityChoiceEnum provided by user to PatternFlowLacpduPartnerStateActivity
	setChoice(value PatternFlowLacpduPartnerStateActivityChoiceEnum) PatternFlowLacpduPartnerStateActivity
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerStateActivity
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerStateActivity.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerStateActivity
	SetValue(value uint32) PatternFlowLacpduPartnerStateActivity
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerStateActivity
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerStateActivity.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerStateActivity
	SetValues(value []uint32) PatternFlowLacpduPartnerStateActivity
	// Increment returns PatternFlowLacpduPartnerStateActivityCounter, set in PatternFlowLacpduPartnerStateActivity.
	// PatternFlowLacpduPartnerStateActivityCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerStateActivityCounter
	// SetIncrement assigns PatternFlowLacpduPartnerStateActivityCounter provided by user to PatternFlowLacpduPartnerStateActivity.
	// PatternFlowLacpduPartnerStateActivityCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerStateActivityCounter) PatternFlowLacpduPartnerStateActivity
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerStateActivity
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerStateActivityCounter, set in PatternFlowLacpduPartnerStateActivity.
	// PatternFlowLacpduPartnerStateActivityCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerStateActivityCounter
	// SetDecrement assigns PatternFlowLacpduPartnerStateActivityCounter provided by user to PatternFlowLacpduPartnerStateActivity.
	// PatternFlowLacpduPartnerStateActivityCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerStateActivityCounter) PatternFlowLacpduPartnerStateActivity
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerStateActivity
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerStateActivityChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerStateActivity
var PatternFlowLacpduPartnerStateActivityChoice = struct {
	VALUE     PatternFlowLacpduPartnerStateActivityChoiceEnum
	VALUES    PatternFlowLacpduPartnerStateActivityChoiceEnum
	INCREMENT PatternFlowLacpduPartnerStateActivityChoiceEnum
	DECREMENT PatternFlowLacpduPartnerStateActivityChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerStateActivityChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerStateActivityChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerStateActivityChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerStateActivityChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerStateActivity) Choice() PatternFlowLacpduPartnerStateActivityChoiceEnum {
	return PatternFlowLacpduPartnerStateActivityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerStateActivity) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerStateActivity) setChoice(value PatternFlowLacpduPartnerStateActivityChoiceEnum) PatternFlowLacpduPartnerStateActivity {
	intValue, ok := otg.PatternFlowLacpduPartnerStateActivity_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerStateActivityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerStateActivity_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerStateActivityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerStateActivityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerStateActivityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerStateActivityCounter().msg()
	}

	if value == PatternFlowLacpduPartnerStateActivityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerStateActivityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateActivity) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateActivityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateActivity) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerStateActivity object
func (obj *patternFlowLacpduPartnerStateActivity) SetValue(value uint32) PatternFlowLacpduPartnerStateActivity {
	obj.setChoice(PatternFlowLacpduPartnerStateActivityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerStateActivity) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerStateActivity object
func (obj *patternFlowLacpduPartnerStateActivity) SetValues(value []uint32) PatternFlowLacpduPartnerStateActivity {
	obj.setChoice(PatternFlowLacpduPartnerStateActivityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateActivityCounter
func (obj *patternFlowLacpduPartnerStateActivity) Increment() PatternFlowLacpduPartnerStateActivityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateActivityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerStateActivityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateActivityCounter
func (obj *patternFlowLacpduPartnerStateActivity) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerStateActivityCounter value in the PatternFlowLacpduPartnerStateActivity object
func (obj *patternFlowLacpduPartnerStateActivity) SetIncrement(value PatternFlowLacpduPartnerStateActivityCounter) PatternFlowLacpduPartnerStateActivity {
	obj.setChoice(PatternFlowLacpduPartnerStateActivityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateActivityCounter
func (obj *patternFlowLacpduPartnerStateActivity) Decrement() PatternFlowLacpduPartnerStateActivityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateActivityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerStateActivityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateActivityCounter
func (obj *patternFlowLacpduPartnerStateActivity) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerStateActivityCounter value in the PatternFlowLacpduPartnerStateActivity object
func (obj *patternFlowLacpduPartnerStateActivity) SetDecrement(value PatternFlowLacpduPartnerStateActivityCounter) PatternFlowLacpduPartnerStateActivity {
	obj.setChoice(PatternFlowLacpduPartnerStateActivityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerStateActivity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateActivity.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerStateActivity.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerStateActivity) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerStateActivityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateActivityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateActivityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateActivityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateActivityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerStateActivityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerStateActivity")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerStateActivity_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerStateActivity_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
