package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerSystemPriority *****
type patternFlowLacpduPartnerSystemPriority struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerSystemPriority
	marshaller      marshalPatternFlowLacpduPartnerSystemPriority
	unMarshaller    unMarshalPatternFlowLacpduPartnerSystemPriority
	incrementHolder PatternFlowLacpduPartnerSystemPriorityCounter
	decrementHolder PatternFlowLacpduPartnerSystemPriorityCounter
}

func NewPatternFlowLacpduPartnerSystemPriority() PatternFlowLacpduPartnerSystemPriority {
	obj := patternFlowLacpduPartnerSystemPriority{obj: &otg.PatternFlowLacpduPartnerSystemPriority{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerSystemPriority) msg() *otg.PatternFlowLacpduPartnerSystemPriority {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerSystemPriority) setMsg(msg *otg.PatternFlowLacpduPartnerSystemPriority) PatternFlowLacpduPartnerSystemPriority {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerSystemPriority struct {
	obj *patternFlowLacpduPartnerSystemPriority
}

type marshalPatternFlowLacpduPartnerSystemPriority interface {
	// ToProto marshals PatternFlowLacpduPartnerSystemPriority to protobuf object *otg.PatternFlowLacpduPartnerSystemPriority
	ToProto() (*otg.PatternFlowLacpduPartnerSystemPriority, error)
	// ToPbText marshals PatternFlowLacpduPartnerSystemPriority to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerSystemPriority to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerSystemPriority to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerSystemPriority struct {
	obj *patternFlowLacpduPartnerSystemPriority
}

type unMarshalPatternFlowLacpduPartnerSystemPriority interface {
	// FromProto unmarshals PatternFlowLacpduPartnerSystemPriority from protobuf object *otg.PatternFlowLacpduPartnerSystemPriority
	FromProto(msg *otg.PatternFlowLacpduPartnerSystemPriority) (PatternFlowLacpduPartnerSystemPriority, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerSystemPriority from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerSystemPriority from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerSystemPriority from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerSystemPriority) Marshal() marshalPatternFlowLacpduPartnerSystemPriority {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerSystemPriority{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerSystemPriority) Unmarshal() unMarshalPatternFlowLacpduPartnerSystemPriority {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerSystemPriority{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerSystemPriority) ToProto() (*otg.PatternFlowLacpduPartnerSystemPriority, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerSystemPriority) FromProto(msg *otg.PatternFlowLacpduPartnerSystemPriority) (PatternFlowLacpduPartnerSystemPriority, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerSystemPriority) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerSystemPriority) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerSystemPriority) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerSystemPriority) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerSystemPriority) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerSystemPriority) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerSystemPriority) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerSystemPriority) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerSystemPriority) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerSystemPriority) Clone() (PatternFlowLacpduPartnerSystemPriority, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerSystemPriority()
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

func (obj *patternFlowLacpduPartnerSystemPriority) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerSystemPriority is the priority of the Partner's system, as received by the Actor.
type PatternFlowLacpduPartnerSystemPriority interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerSystemPriority to protobuf object *otg.PatternFlowLacpduPartnerSystemPriority
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerSystemPriority
	// setMsg unmarshals PatternFlowLacpduPartnerSystemPriority from protobuf object *otg.PatternFlowLacpduPartnerSystemPriority
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerSystemPriority) PatternFlowLacpduPartnerSystemPriority
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerSystemPriority
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerSystemPriority
	// validate validates PatternFlowLacpduPartnerSystemPriority
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerSystemPriority, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerSystemPriorityChoiceEnum, set in PatternFlowLacpduPartnerSystemPriority
	Choice() PatternFlowLacpduPartnerSystemPriorityChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerSystemPriorityChoiceEnum provided by user to PatternFlowLacpduPartnerSystemPriority
	setChoice(value PatternFlowLacpduPartnerSystemPriorityChoiceEnum) PatternFlowLacpduPartnerSystemPriority
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerSystemPriority
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerSystemPriority.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerSystemPriority
	SetValue(value uint32) PatternFlowLacpduPartnerSystemPriority
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerSystemPriority
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerSystemPriority.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerSystemPriority
	SetValues(value []uint32) PatternFlowLacpduPartnerSystemPriority
	// Increment returns PatternFlowLacpduPartnerSystemPriorityCounter, set in PatternFlowLacpduPartnerSystemPriority.
	// PatternFlowLacpduPartnerSystemPriorityCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerSystemPriorityCounter
	// SetIncrement assigns PatternFlowLacpduPartnerSystemPriorityCounter provided by user to PatternFlowLacpduPartnerSystemPriority.
	// PatternFlowLacpduPartnerSystemPriorityCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerSystemPriorityCounter) PatternFlowLacpduPartnerSystemPriority
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerSystemPriority
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerSystemPriorityCounter, set in PatternFlowLacpduPartnerSystemPriority.
	// PatternFlowLacpduPartnerSystemPriorityCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerSystemPriorityCounter
	// SetDecrement assigns PatternFlowLacpduPartnerSystemPriorityCounter provided by user to PatternFlowLacpduPartnerSystemPriority.
	// PatternFlowLacpduPartnerSystemPriorityCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerSystemPriorityCounter) PatternFlowLacpduPartnerSystemPriority
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerSystemPriority
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerSystemPriorityChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerSystemPriority
var PatternFlowLacpduPartnerSystemPriorityChoice = struct {
	VALUE     PatternFlowLacpduPartnerSystemPriorityChoiceEnum
	VALUES    PatternFlowLacpduPartnerSystemPriorityChoiceEnum
	INCREMENT PatternFlowLacpduPartnerSystemPriorityChoiceEnum
	DECREMENT PatternFlowLacpduPartnerSystemPriorityChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerSystemPriorityChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerSystemPriorityChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerSystemPriorityChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerSystemPriorityChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerSystemPriority) Choice() PatternFlowLacpduPartnerSystemPriorityChoiceEnum {
	return PatternFlowLacpduPartnerSystemPriorityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerSystemPriority) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerSystemPriority) setChoice(value PatternFlowLacpduPartnerSystemPriorityChoiceEnum) PatternFlowLacpduPartnerSystemPriority {
	intValue, ok := otg.PatternFlowLacpduPartnerSystemPriority_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerSystemPriorityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerSystemPriority_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerSystemPriorityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerSystemPriorityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerSystemPriorityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerSystemPriorityCounter().msg()
	}

	if value == PatternFlowLacpduPartnerSystemPriorityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerSystemPriorityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerSystemPriority) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerSystemPriorityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerSystemPriority) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerSystemPriority object
func (obj *patternFlowLacpduPartnerSystemPriority) SetValue(value uint32) PatternFlowLacpduPartnerSystemPriority {
	obj.setChoice(PatternFlowLacpduPartnerSystemPriorityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerSystemPriority) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerSystemPriority object
func (obj *patternFlowLacpduPartnerSystemPriority) SetValues(value []uint32) PatternFlowLacpduPartnerSystemPriority {
	obj.setChoice(PatternFlowLacpduPartnerSystemPriorityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerSystemPriorityCounter
func (obj *patternFlowLacpduPartnerSystemPriority) Increment() PatternFlowLacpduPartnerSystemPriorityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerSystemPriorityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerSystemPriorityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerSystemPriorityCounter
func (obj *patternFlowLacpduPartnerSystemPriority) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerSystemPriorityCounter value in the PatternFlowLacpduPartnerSystemPriority object
func (obj *patternFlowLacpduPartnerSystemPriority) SetIncrement(value PatternFlowLacpduPartnerSystemPriorityCounter) PatternFlowLacpduPartnerSystemPriority {
	obj.setChoice(PatternFlowLacpduPartnerSystemPriorityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerSystemPriorityCounter
func (obj *patternFlowLacpduPartnerSystemPriority) Decrement() PatternFlowLacpduPartnerSystemPriorityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerSystemPriorityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerSystemPriorityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerSystemPriorityCounter
func (obj *patternFlowLacpduPartnerSystemPriority) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerSystemPriorityCounter value in the PatternFlowLacpduPartnerSystemPriority object
func (obj *patternFlowLacpduPartnerSystemPriority) SetDecrement(value PatternFlowLacpduPartnerSystemPriorityCounter) PatternFlowLacpduPartnerSystemPriority {
	obj.setChoice(PatternFlowLacpduPartnerSystemPriorityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerSystemPriority) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerSystemPriority.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerSystemPriority.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerSystemPriority) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerSystemPriorityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerSystemPriorityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerSystemPriorityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerSystemPriorityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerSystemPriorityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerSystemPriorityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerSystemPriority")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerSystemPriority_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerSystemPriority_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
