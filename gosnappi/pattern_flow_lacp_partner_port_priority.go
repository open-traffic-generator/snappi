package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerPortPriority *****
type patternFlowLacpPartnerPortPriority struct {
	validation
	obj             *otg.PatternFlowLacpPartnerPortPriority
	marshaller      marshalPatternFlowLacpPartnerPortPriority
	unMarshaller    unMarshalPatternFlowLacpPartnerPortPriority
	incrementHolder PatternFlowLacpPartnerPortPriorityCounter
	decrementHolder PatternFlowLacpPartnerPortPriorityCounter
}

func NewPatternFlowLacpPartnerPortPriority() PatternFlowLacpPartnerPortPriority {
	obj := patternFlowLacpPartnerPortPriority{obj: &otg.PatternFlowLacpPartnerPortPriority{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerPortPriority) msg() *otg.PatternFlowLacpPartnerPortPriority {
	return obj.obj
}

func (obj *patternFlowLacpPartnerPortPriority) setMsg(msg *otg.PatternFlowLacpPartnerPortPriority) PatternFlowLacpPartnerPortPriority {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerPortPriority struct {
	obj *patternFlowLacpPartnerPortPriority
}

type marshalPatternFlowLacpPartnerPortPriority interface {
	// ToProto marshals PatternFlowLacpPartnerPortPriority to protobuf object *otg.PatternFlowLacpPartnerPortPriority
	ToProto() (*otg.PatternFlowLacpPartnerPortPriority, error)
	// ToPbText marshals PatternFlowLacpPartnerPortPriority to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerPortPriority to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerPortPriority to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerPortPriority struct {
	obj *patternFlowLacpPartnerPortPriority
}

type unMarshalPatternFlowLacpPartnerPortPriority interface {
	// FromProto unmarshals PatternFlowLacpPartnerPortPriority from protobuf object *otg.PatternFlowLacpPartnerPortPriority
	FromProto(msg *otg.PatternFlowLacpPartnerPortPriority) (PatternFlowLacpPartnerPortPriority, error)
	// FromPbText unmarshals PatternFlowLacpPartnerPortPriority from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerPortPriority from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerPortPriority from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerPortPriority) Marshal() marshalPatternFlowLacpPartnerPortPriority {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerPortPriority{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerPortPriority) Unmarshal() unMarshalPatternFlowLacpPartnerPortPriority {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerPortPriority{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerPortPriority) ToProto() (*otg.PatternFlowLacpPartnerPortPriority, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerPortPriority) FromProto(msg *otg.PatternFlowLacpPartnerPortPriority) (PatternFlowLacpPartnerPortPriority, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerPortPriority) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerPortPriority) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerPortPriority) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerPortPriority) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerPortPriority) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerPortPriority) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerPortPriority) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerPortPriority) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerPortPriority) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerPortPriority) Clone() (PatternFlowLacpPartnerPortPriority, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerPortPriority()
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

func (obj *patternFlowLacpPartnerPortPriority) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerPortPriority is the priority of the Partner's port, as received by the Actor.
type PatternFlowLacpPartnerPortPriority interface {
	Validation
	// msg marshals PatternFlowLacpPartnerPortPriority to protobuf object *otg.PatternFlowLacpPartnerPortPriority
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerPortPriority
	// setMsg unmarshals PatternFlowLacpPartnerPortPriority from protobuf object *otg.PatternFlowLacpPartnerPortPriority
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerPortPriority) PatternFlowLacpPartnerPortPriority
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerPortPriority
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerPortPriority
	// validate validates PatternFlowLacpPartnerPortPriority
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerPortPriority, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerPortPriorityChoiceEnum, set in PatternFlowLacpPartnerPortPriority
	Choice() PatternFlowLacpPartnerPortPriorityChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerPortPriorityChoiceEnum provided by user to PatternFlowLacpPartnerPortPriority
	setChoice(value PatternFlowLacpPartnerPortPriorityChoiceEnum) PatternFlowLacpPartnerPortPriority
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerPortPriority
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerPortPriority.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerPortPriority
	SetValue(value uint32) PatternFlowLacpPartnerPortPriority
	// HasValue checks if Value has been set in PatternFlowLacpPartnerPortPriority
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerPortPriority.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerPortPriority
	SetValues(value []uint32) PatternFlowLacpPartnerPortPriority
	// Increment returns PatternFlowLacpPartnerPortPriorityCounter, set in PatternFlowLacpPartnerPortPriority.
	// PatternFlowLacpPartnerPortPriorityCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerPortPriorityCounter
	// SetIncrement assigns PatternFlowLacpPartnerPortPriorityCounter provided by user to PatternFlowLacpPartnerPortPriority.
	// PatternFlowLacpPartnerPortPriorityCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerPortPriorityCounter) PatternFlowLacpPartnerPortPriority
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerPortPriority
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerPortPriorityCounter, set in PatternFlowLacpPartnerPortPriority.
	// PatternFlowLacpPartnerPortPriorityCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerPortPriorityCounter
	// SetDecrement assigns PatternFlowLacpPartnerPortPriorityCounter provided by user to PatternFlowLacpPartnerPortPriority.
	// PatternFlowLacpPartnerPortPriorityCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerPortPriorityCounter) PatternFlowLacpPartnerPortPriority
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerPortPriority
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerPortPriorityChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerPortPriority
var PatternFlowLacpPartnerPortPriorityChoice = struct {
	VALUE     PatternFlowLacpPartnerPortPriorityChoiceEnum
	VALUES    PatternFlowLacpPartnerPortPriorityChoiceEnum
	INCREMENT PatternFlowLacpPartnerPortPriorityChoiceEnum
	DECREMENT PatternFlowLacpPartnerPortPriorityChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerPortPriorityChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerPortPriorityChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerPortPriorityChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerPortPriorityChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerPortPriority) Choice() PatternFlowLacpPartnerPortPriorityChoiceEnum {
	return PatternFlowLacpPartnerPortPriorityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerPortPriority) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerPortPriority) setChoice(value PatternFlowLacpPartnerPortPriorityChoiceEnum) PatternFlowLacpPartnerPortPriority {
	intValue, ok := otg.PatternFlowLacpPartnerPortPriority_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerPortPriorityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerPortPriority_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerPortPriorityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerPortPriorityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerPortPriorityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerPortPriorityCounter().msg()
	}

	if value == PatternFlowLacpPartnerPortPriorityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerPortPriorityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerPortPriority) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerPortPriorityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerPortPriority) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerPortPriority object
func (obj *patternFlowLacpPartnerPortPriority) SetValue(value uint32) PatternFlowLacpPartnerPortPriority {
	obj.setChoice(PatternFlowLacpPartnerPortPriorityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerPortPriority) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerPortPriority object
func (obj *patternFlowLacpPartnerPortPriority) SetValues(value []uint32) PatternFlowLacpPartnerPortPriority {
	obj.setChoice(PatternFlowLacpPartnerPortPriorityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerPortPriorityCounter
func (obj *patternFlowLacpPartnerPortPriority) Increment() PatternFlowLacpPartnerPortPriorityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerPortPriorityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerPortPriorityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerPortPriorityCounter
func (obj *patternFlowLacpPartnerPortPriority) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerPortPriorityCounter value in the PatternFlowLacpPartnerPortPriority object
func (obj *patternFlowLacpPartnerPortPriority) SetIncrement(value PatternFlowLacpPartnerPortPriorityCounter) PatternFlowLacpPartnerPortPriority {
	obj.setChoice(PatternFlowLacpPartnerPortPriorityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerPortPriorityCounter
func (obj *patternFlowLacpPartnerPortPriority) Decrement() PatternFlowLacpPartnerPortPriorityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerPortPriorityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerPortPriorityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerPortPriorityCounter
func (obj *patternFlowLacpPartnerPortPriority) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerPortPriorityCounter value in the PatternFlowLacpPartnerPortPriority object
func (obj *patternFlowLacpPartnerPortPriority) SetDecrement(value PatternFlowLacpPartnerPortPriorityCounter) PatternFlowLacpPartnerPortPriority {
	obj.setChoice(PatternFlowLacpPartnerPortPriorityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerPortPriority) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerPortPriority.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerPortPriority.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpPartnerPortPriority) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerPortPriorityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerPortPriorityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerPortPriorityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerPortPriorityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerPortPriorityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerPortPriorityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerPortPriority")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerPortPriority_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerPortPriority_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
