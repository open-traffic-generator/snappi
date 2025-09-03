package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerSystemPriority *****
type patternFlowLacpPartnerSystemPriority struct {
	validation
	obj             *otg.PatternFlowLacpPartnerSystemPriority
	marshaller      marshalPatternFlowLacpPartnerSystemPriority
	unMarshaller    unMarshalPatternFlowLacpPartnerSystemPriority
	incrementHolder PatternFlowLacpPartnerSystemPriorityCounter
	decrementHolder PatternFlowLacpPartnerSystemPriorityCounter
}

func NewPatternFlowLacpPartnerSystemPriority() PatternFlowLacpPartnerSystemPriority {
	obj := patternFlowLacpPartnerSystemPriority{obj: &otg.PatternFlowLacpPartnerSystemPriority{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerSystemPriority) msg() *otg.PatternFlowLacpPartnerSystemPriority {
	return obj.obj
}

func (obj *patternFlowLacpPartnerSystemPriority) setMsg(msg *otg.PatternFlowLacpPartnerSystemPriority) PatternFlowLacpPartnerSystemPriority {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerSystemPriority struct {
	obj *patternFlowLacpPartnerSystemPriority
}

type marshalPatternFlowLacpPartnerSystemPriority interface {
	// ToProto marshals PatternFlowLacpPartnerSystemPriority to protobuf object *otg.PatternFlowLacpPartnerSystemPriority
	ToProto() (*otg.PatternFlowLacpPartnerSystemPriority, error)
	// ToPbText marshals PatternFlowLacpPartnerSystemPriority to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerSystemPriority to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerSystemPriority to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerSystemPriority struct {
	obj *patternFlowLacpPartnerSystemPriority
}

type unMarshalPatternFlowLacpPartnerSystemPriority interface {
	// FromProto unmarshals PatternFlowLacpPartnerSystemPriority from protobuf object *otg.PatternFlowLacpPartnerSystemPriority
	FromProto(msg *otg.PatternFlowLacpPartnerSystemPriority) (PatternFlowLacpPartnerSystemPriority, error)
	// FromPbText unmarshals PatternFlowLacpPartnerSystemPriority from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerSystemPriority from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerSystemPriority from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerSystemPriority) Marshal() marshalPatternFlowLacpPartnerSystemPriority {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerSystemPriority{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerSystemPriority) Unmarshal() unMarshalPatternFlowLacpPartnerSystemPriority {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerSystemPriority{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerSystemPriority) ToProto() (*otg.PatternFlowLacpPartnerSystemPriority, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerSystemPriority) FromProto(msg *otg.PatternFlowLacpPartnerSystemPriority) (PatternFlowLacpPartnerSystemPriority, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerSystemPriority) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerSystemPriority) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerSystemPriority) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerSystemPriority) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerSystemPriority) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerSystemPriority) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerSystemPriority) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerSystemPriority) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerSystemPriority) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerSystemPriority) Clone() (PatternFlowLacpPartnerSystemPriority, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerSystemPriority()
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

func (obj *patternFlowLacpPartnerSystemPriority) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerSystemPriority is the priority of the Partner's system, as received by the Actor.
type PatternFlowLacpPartnerSystemPriority interface {
	Validation
	// msg marshals PatternFlowLacpPartnerSystemPriority to protobuf object *otg.PatternFlowLacpPartnerSystemPriority
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerSystemPriority
	// setMsg unmarshals PatternFlowLacpPartnerSystemPriority from protobuf object *otg.PatternFlowLacpPartnerSystemPriority
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerSystemPriority) PatternFlowLacpPartnerSystemPriority
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerSystemPriority
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerSystemPriority
	// validate validates PatternFlowLacpPartnerSystemPriority
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerSystemPriority, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerSystemPriorityChoiceEnum, set in PatternFlowLacpPartnerSystemPriority
	Choice() PatternFlowLacpPartnerSystemPriorityChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerSystemPriorityChoiceEnum provided by user to PatternFlowLacpPartnerSystemPriority
	setChoice(value PatternFlowLacpPartnerSystemPriorityChoiceEnum) PatternFlowLacpPartnerSystemPriority
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerSystemPriority
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerSystemPriority.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerSystemPriority
	SetValue(value uint32) PatternFlowLacpPartnerSystemPriority
	// HasValue checks if Value has been set in PatternFlowLacpPartnerSystemPriority
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerSystemPriority.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerSystemPriority
	SetValues(value []uint32) PatternFlowLacpPartnerSystemPriority
	// Increment returns PatternFlowLacpPartnerSystemPriorityCounter, set in PatternFlowLacpPartnerSystemPriority.
	// PatternFlowLacpPartnerSystemPriorityCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerSystemPriorityCounter
	// SetIncrement assigns PatternFlowLacpPartnerSystemPriorityCounter provided by user to PatternFlowLacpPartnerSystemPriority.
	// PatternFlowLacpPartnerSystemPriorityCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerSystemPriorityCounter) PatternFlowLacpPartnerSystemPriority
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerSystemPriority
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerSystemPriorityCounter, set in PatternFlowLacpPartnerSystemPriority.
	// PatternFlowLacpPartnerSystemPriorityCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerSystemPriorityCounter
	// SetDecrement assigns PatternFlowLacpPartnerSystemPriorityCounter provided by user to PatternFlowLacpPartnerSystemPriority.
	// PatternFlowLacpPartnerSystemPriorityCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerSystemPriorityCounter) PatternFlowLacpPartnerSystemPriority
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerSystemPriority
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerSystemPriorityChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerSystemPriority
var PatternFlowLacpPartnerSystemPriorityChoice = struct {
	VALUE     PatternFlowLacpPartnerSystemPriorityChoiceEnum
	VALUES    PatternFlowLacpPartnerSystemPriorityChoiceEnum
	INCREMENT PatternFlowLacpPartnerSystemPriorityChoiceEnum
	DECREMENT PatternFlowLacpPartnerSystemPriorityChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerSystemPriorityChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerSystemPriorityChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerSystemPriorityChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerSystemPriorityChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerSystemPriority) Choice() PatternFlowLacpPartnerSystemPriorityChoiceEnum {
	return PatternFlowLacpPartnerSystemPriorityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerSystemPriority) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerSystemPriority) setChoice(value PatternFlowLacpPartnerSystemPriorityChoiceEnum) PatternFlowLacpPartnerSystemPriority {
	intValue, ok := otg.PatternFlowLacpPartnerSystemPriority_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerSystemPriorityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerSystemPriority_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerSystemPriorityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerSystemPriorityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerSystemPriorityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerSystemPriorityCounter().msg()
	}

	if value == PatternFlowLacpPartnerSystemPriorityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerSystemPriorityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerSystemPriority) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerSystemPriorityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerSystemPriority) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerSystemPriority object
func (obj *patternFlowLacpPartnerSystemPriority) SetValue(value uint32) PatternFlowLacpPartnerSystemPriority {
	obj.setChoice(PatternFlowLacpPartnerSystemPriorityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerSystemPriority) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerSystemPriority object
func (obj *patternFlowLacpPartnerSystemPriority) SetValues(value []uint32) PatternFlowLacpPartnerSystemPriority {
	obj.setChoice(PatternFlowLacpPartnerSystemPriorityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerSystemPriorityCounter
func (obj *patternFlowLacpPartnerSystemPriority) Increment() PatternFlowLacpPartnerSystemPriorityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerSystemPriorityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerSystemPriorityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerSystemPriorityCounter
func (obj *patternFlowLacpPartnerSystemPriority) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerSystemPriorityCounter value in the PatternFlowLacpPartnerSystemPriority object
func (obj *patternFlowLacpPartnerSystemPriority) SetIncrement(value PatternFlowLacpPartnerSystemPriorityCounter) PatternFlowLacpPartnerSystemPriority {
	obj.setChoice(PatternFlowLacpPartnerSystemPriorityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerSystemPriorityCounter
func (obj *patternFlowLacpPartnerSystemPriority) Decrement() PatternFlowLacpPartnerSystemPriorityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerSystemPriorityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerSystemPriorityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerSystemPriorityCounter
func (obj *patternFlowLacpPartnerSystemPriority) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerSystemPriorityCounter value in the PatternFlowLacpPartnerSystemPriority object
func (obj *patternFlowLacpPartnerSystemPriority) SetDecrement(value PatternFlowLacpPartnerSystemPriorityCounter) PatternFlowLacpPartnerSystemPriority {
	obj.setChoice(PatternFlowLacpPartnerSystemPriorityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerSystemPriority) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerSystemPriority.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerSystemPriority.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpPartnerSystemPriority) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerSystemPriorityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerSystemPriorityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerSystemPriorityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerSystemPriorityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerSystemPriorityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerSystemPriorityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerSystemPriority")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerSystemPriority_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerSystemPriority_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
