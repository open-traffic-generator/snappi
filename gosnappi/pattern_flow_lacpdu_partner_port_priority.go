package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerPortPriority *****
type patternFlowLacpduPartnerPortPriority struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerPortPriority
	marshaller      marshalPatternFlowLacpduPartnerPortPriority
	unMarshaller    unMarshalPatternFlowLacpduPartnerPortPriority
	incrementHolder PatternFlowLacpduPartnerPortPriorityCounter
	decrementHolder PatternFlowLacpduPartnerPortPriorityCounter
}

func NewPatternFlowLacpduPartnerPortPriority() PatternFlowLacpduPartnerPortPriority {
	obj := patternFlowLacpduPartnerPortPriority{obj: &otg.PatternFlowLacpduPartnerPortPriority{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerPortPriority) msg() *otg.PatternFlowLacpduPartnerPortPriority {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerPortPriority) setMsg(msg *otg.PatternFlowLacpduPartnerPortPriority) PatternFlowLacpduPartnerPortPriority {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerPortPriority struct {
	obj *patternFlowLacpduPartnerPortPriority
}

type marshalPatternFlowLacpduPartnerPortPriority interface {
	// ToProto marshals PatternFlowLacpduPartnerPortPriority to protobuf object *otg.PatternFlowLacpduPartnerPortPriority
	ToProto() (*otg.PatternFlowLacpduPartnerPortPriority, error)
	// ToPbText marshals PatternFlowLacpduPartnerPortPriority to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerPortPriority to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerPortPriority to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerPortPriority struct {
	obj *patternFlowLacpduPartnerPortPriority
}

type unMarshalPatternFlowLacpduPartnerPortPriority interface {
	// FromProto unmarshals PatternFlowLacpduPartnerPortPriority from protobuf object *otg.PatternFlowLacpduPartnerPortPriority
	FromProto(msg *otg.PatternFlowLacpduPartnerPortPriority) (PatternFlowLacpduPartnerPortPriority, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerPortPriority from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerPortPriority from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerPortPriority from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerPortPriority) Marshal() marshalPatternFlowLacpduPartnerPortPriority {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerPortPriority{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerPortPriority) Unmarshal() unMarshalPatternFlowLacpduPartnerPortPriority {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerPortPriority{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerPortPriority) ToProto() (*otg.PatternFlowLacpduPartnerPortPriority, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerPortPriority) FromProto(msg *otg.PatternFlowLacpduPartnerPortPriority) (PatternFlowLacpduPartnerPortPriority, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerPortPriority) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerPortPriority) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerPortPriority) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerPortPriority) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerPortPriority) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerPortPriority) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerPortPriority) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerPortPriority) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerPortPriority) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerPortPriority) Clone() (PatternFlowLacpduPartnerPortPriority, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerPortPriority()
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

func (obj *patternFlowLacpduPartnerPortPriority) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerPortPriority is the priority of the Partner's port, as received by the Actor.
type PatternFlowLacpduPartnerPortPriority interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerPortPriority to protobuf object *otg.PatternFlowLacpduPartnerPortPriority
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerPortPriority
	// setMsg unmarshals PatternFlowLacpduPartnerPortPriority from protobuf object *otg.PatternFlowLacpduPartnerPortPriority
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerPortPriority) PatternFlowLacpduPartnerPortPriority
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerPortPriority
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerPortPriority
	// validate validates PatternFlowLacpduPartnerPortPriority
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerPortPriority, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerPortPriorityChoiceEnum, set in PatternFlowLacpduPartnerPortPriority
	Choice() PatternFlowLacpduPartnerPortPriorityChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerPortPriorityChoiceEnum provided by user to PatternFlowLacpduPartnerPortPriority
	setChoice(value PatternFlowLacpduPartnerPortPriorityChoiceEnum) PatternFlowLacpduPartnerPortPriority
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerPortPriority
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerPortPriority.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerPortPriority
	SetValue(value uint32) PatternFlowLacpduPartnerPortPriority
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerPortPriority
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerPortPriority.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerPortPriority
	SetValues(value []uint32) PatternFlowLacpduPartnerPortPriority
	// Increment returns PatternFlowLacpduPartnerPortPriorityCounter, set in PatternFlowLacpduPartnerPortPriority.
	// PatternFlowLacpduPartnerPortPriorityCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerPortPriorityCounter
	// SetIncrement assigns PatternFlowLacpduPartnerPortPriorityCounter provided by user to PatternFlowLacpduPartnerPortPriority.
	// PatternFlowLacpduPartnerPortPriorityCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerPortPriorityCounter) PatternFlowLacpduPartnerPortPriority
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerPortPriority
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerPortPriorityCounter, set in PatternFlowLacpduPartnerPortPriority.
	// PatternFlowLacpduPartnerPortPriorityCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerPortPriorityCounter
	// SetDecrement assigns PatternFlowLacpduPartnerPortPriorityCounter provided by user to PatternFlowLacpduPartnerPortPriority.
	// PatternFlowLacpduPartnerPortPriorityCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerPortPriorityCounter) PatternFlowLacpduPartnerPortPriority
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerPortPriority
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerPortPriorityChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerPortPriority
var PatternFlowLacpduPartnerPortPriorityChoice = struct {
	VALUE     PatternFlowLacpduPartnerPortPriorityChoiceEnum
	VALUES    PatternFlowLacpduPartnerPortPriorityChoiceEnum
	INCREMENT PatternFlowLacpduPartnerPortPriorityChoiceEnum
	DECREMENT PatternFlowLacpduPartnerPortPriorityChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerPortPriorityChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerPortPriorityChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerPortPriorityChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerPortPriorityChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerPortPriority) Choice() PatternFlowLacpduPartnerPortPriorityChoiceEnum {
	return PatternFlowLacpduPartnerPortPriorityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerPortPriority) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerPortPriority) setChoice(value PatternFlowLacpduPartnerPortPriorityChoiceEnum) PatternFlowLacpduPartnerPortPriority {
	intValue, ok := otg.PatternFlowLacpduPartnerPortPriority_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerPortPriorityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerPortPriority_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerPortPriorityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerPortPriorityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerPortPriorityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerPortPriorityCounter().msg()
	}

	if value == PatternFlowLacpduPartnerPortPriorityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerPortPriorityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerPortPriority) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerPortPriorityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerPortPriority) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerPortPriority object
func (obj *patternFlowLacpduPartnerPortPriority) SetValue(value uint32) PatternFlowLacpduPartnerPortPriority {
	obj.setChoice(PatternFlowLacpduPartnerPortPriorityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerPortPriority) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerPortPriority object
func (obj *patternFlowLacpduPartnerPortPriority) SetValues(value []uint32) PatternFlowLacpduPartnerPortPriority {
	obj.setChoice(PatternFlowLacpduPartnerPortPriorityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerPortPriorityCounter
func (obj *patternFlowLacpduPartnerPortPriority) Increment() PatternFlowLacpduPartnerPortPriorityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerPortPriorityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerPortPriorityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerPortPriorityCounter
func (obj *patternFlowLacpduPartnerPortPriority) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerPortPriorityCounter value in the PatternFlowLacpduPartnerPortPriority object
func (obj *patternFlowLacpduPartnerPortPriority) SetIncrement(value PatternFlowLacpduPartnerPortPriorityCounter) PatternFlowLacpduPartnerPortPriority {
	obj.setChoice(PatternFlowLacpduPartnerPortPriorityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerPortPriorityCounter
func (obj *patternFlowLacpduPartnerPortPriority) Decrement() PatternFlowLacpduPartnerPortPriorityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerPortPriorityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerPortPriorityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerPortPriorityCounter
func (obj *patternFlowLacpduPartnerPortPriority) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerPortPriorityCounter value in the PatternFlowLacpduPartnerPortPriority object
func (obj *patternFlowLacpduPartnerPortPriority) SetDecrement(value PatternFlowLacpduPartnerPortPriorityCounter) PatternFlowLacpduPartnerPortPriority {
	obj.setChoice(PatternFlowLacpduPartnerPortPriorityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerPortPriority) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerPortPriority.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerPortPriority.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerPortPriority) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerPortPriorityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerPortPriorityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerPortPriorityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerPortPriorityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerPortPriorityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerPortPriorityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerPortPriority")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerPortPriority_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerPortPriority_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
