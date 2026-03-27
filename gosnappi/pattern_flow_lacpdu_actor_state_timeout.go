package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateTimeout *****
type patternFlowLacpduActorStateTimeout struct {
	validation
	obj             *otg.PatternFlowLacpduActorStateTimeout
	marshaller      marshalPatternFlowLacpduActorStateTimeout
	unMarshaller    unMarshalPatternFlowLacpduActorStateTimeout
	incrementHolder PatternFlowLacpduActorStateTimeoutCounter
	decrementHolder PatternFlowLacpduActorStateTimeoutCounter
}

func NewPatternFlowLacpduActorStateTimeout() PatternFlowLacpduActorStateTimeout {
	obj := patternFlowLacpduActorStateTimeout{obj: &otg.PatternFlowLacpduActorStateTimeout{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateTimeout) msg() *otg.PatternFlowLacpduActorStateTimeout {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateTimeout) setMsg(msg *otg.PatternFlowLacpduActorStateTimeout) PatternFlowLacpduActorStateTimeout {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateTimeout struct {
	obj *patternFlowLacpduActorStateTimeout
}

type marshalPatternFlowLacpduActorStateTimeout interface {
	// ToProto marshals PatternFlowLacpduActorStateTimeout to protobuf object *otg.PatternFlowLacpduActorStateTimeout
	ToProto() (*otg.PatternFlowLacpduActorStateTimeout, error)
	// ToPbText marshals PatternFlowLacpduActorStateTimeout to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateTimeout to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateTimeout to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateTimeout struct {
	obj *patternFlowLacpduActorStateTimeout
}

type unMarshalPatternFlowLacpduActorStateTimeout interface {
	// FromProto unmarshals PatternFlowLacpduActorStateTimeout from protobuf object *otg.PatternFlowLacpduActorStateTimeout
	FromProto(msg *otg.PatternFlowLacpduActorStateTimeout) (PatternFlowLacpduActorStateTimeout, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateTimeout from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateTimeout from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateTimeout from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateTimeout) Marshal() marshalPatternFlowLacpduActorStateTimeout {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateTimeout{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateTimeout) Unmarshal() unMarshalPatternFlowLacpduActorStateTimeout {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateTimeout{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateTimeout) ToProto() (*otg.PatternFlowLacpduActorStateTimeout, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateTimeout) FromProto(msg *otg.PatternFlowLacpduActorStateTimeout) (PatternFlowLacpduActorStateTimeout, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateTimeout) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateTimeout) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateTimeout) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateTimeout) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateTimeout) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateTimeout) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateTimeout) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateTimeout) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateTimeout) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateTimeout) Clone() (PatternFlowLacpduActorStateTimeout, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateTimeout()
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

func (obj *patternFlowLacpduActorStateTimeout) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorStateTimeout is lACP Timeout (1=Fast timeout, 0=Slow timeout)
type PatternFlowLacpduActorStateTimeout interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateTimeout to protobuf object *otg.PatternFlowLacpduActorStateTimeout
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateTimeout
	// setMsg unmarshals PatternFlowLacpduActorStateTimeout from protobuf object *otg.PatternFlowLacpduActorStateTimeout
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateTimeout) PatternFlowLacpduActorStateTimeout
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateTimeout
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateTimeout
	// validate validates PatternFlowLacpduActorStateTimeout
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateTimeout, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorStateTimeoutChoiceEnum, set in PatternFlowLacpduActorStateTimeout
	Choice() PatternFlowLacpduActorStateTimeoutChoiceEnum
	// setChoice assigns PatternFlowLacpduActorStateTimeoutChoiceEnum provided by user to PatternFlowLacpduActorStateTimeout
	setChoice(value PatternFlowLacpduActorStateTimeoutChoiceEnum) PatternFlowLacpduActorStateTimeout
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorStateTimeout
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorStateTimeout.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorStateTimeout
	SetValue(value uint32) PatternFlowLacpduActorStateTimeout
	// HasValue checks if Value has been set in PatternFlowLacpduActorStateTimeout
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorStateTimeout.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorStateTimeout
	SetValues(value []uint32) PatternFlowLacpduActorStateTimeout
	// Increment returns PatternFlowLacpduActorStateTimeoutCounter, set in PatternFlowLacpduActorStateTimeout.
	// PatternFlowLacpduActorStateTimeoutCounter is integer counter pattern
	Increment() PatternFlowLacpduActorStateTimeoutCounter
	// SetIncrement assigns PatternFlowLacpduActorStateTimeoutCounter provided by user to PatternFlowLacpduActorStateTimeout.
	// PatternFlowLacpduActorStateTimeoutCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorStateTimeoutCounter) PatternFlowLacpduActorStateTimeout
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorStateTimeout
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorStateTimeoutCounter, set in PatternFlowLacpduActorStateTimeout.
	// PatternFlowLacpduActorStateTimeoutCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorStateTimeoutCounter
	// SetDecrement assigns PatternFlowLacpduActorStateTimeoutCounter provided by user to PatternFlowLacpduActorStateTimeout.
	// PatternFlowLacpduActorStateTimeoutCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorStateTimeoutCounter) PatternFlowLacpduActorStateTimeout
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorStateTimeout
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorStateTimeoutChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorStateTimeout
var PatternFlowLacpduActorStateTimeoutChoice = struct {
	VALUE     PatternFlowLacpduActorStateTimeoutChoiceEnum
	VALUES    PatternFlowLacpduActorStateTimeoutChoiceEnum
	INCREMENT PatternFlowLacpduActorStateTimeoutChoiceEnum
	DECREMENT PatternFlowLacpduActorStateTimeoutChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorStateTimeoutChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorStateTimeoutChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorStateTimeoutChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorStateTimeoutChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorStateTimeout) Choice() PatternFlowLacpduActorStateTimeoutChoiceEnum {
	return PatternFlowLacpduActorStateTimeoutChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorStateTimeout) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorStateTimeout) setChoice(value PatternFlowLacpduActorStateTimeoutChoiceEnum) PatternFlowLacpduActorStateTimeout {
	intValue, ok := otg.PatternFlowLacpduActorStateTimeout_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorStateTimeoutChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorStateTimeout_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorStateTimeoutChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorStateTimeoutChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorStateTimeoutChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorStateTimeoutCounter().msg()
	}

	if value == PatternFlowLacpduActorStateTimeoutChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorStateTimeoutCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateTimeout) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorStateTimeoutChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateTimeout) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorStateTimeout object
func (obj *patternFlowLacpduActorStateTimeout) SetValue(value uint32) PatternFlowLacpduActorStateTimeout {
	obj.setChoice(PatternFlowLacpduActorStateTimeoutChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorStateTimeout) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorStateTimeout object
func (obj *patternFlowLacpduActorStateTimeout) SetValues(value []uint32) PatternFlowLacpduActorStateTimeout {
	obj.setChoice(PatternFlowLacpduActorStateTimeoutChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateTimeoutCounter
func (obj *patternFlowLacpduActorStateTimeout) Increment() PatternFlowLacpduActorStateTimeoutCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorStateTimeoutChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorStateTimeoutCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateTimeoutCounter
func (obj *patternFlowLacpduActorStateTimeout) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorStateTimeoutCounter value in the PatternFlowLacpduActorStateTimeout object
func (obj *patternFlowLacpduActorStateTimeout) SetIncrement(value PatternFlowLacpduActorStateTimeoutCounter) PatternFlowLacpduActorStateTimeout {
	obj.setChoice(PatternFlowLacpduActorStateTimeoutChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateTimeoutCounter
func (obj *patternFlowLacpduActorStateTimeout) Decrement() PatternFlowLacpduActorStateTimeoutCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorStateTimeoutChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorStateTimeoutCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateTimeoutCounter
func (obj *patternFlowLacpduActorStateTimeout) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorStateTimeoutCounter value in the PatternFlowLacpduActorStateTimeout object
func (obj *patternFlowLacpduActorStateTimeout) SetDecrement(value PatternFlowLacpduActorStateTimeoutCounter) PatternFlowLacpduActorStateTimeout {
	obj.setChoice(PatternFlowLacpduActorStateTimeoutChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorStateTimeout) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateTimeout.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorStateTimeout.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduActorStateTimeout) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorStateTimeoutChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateTimeoutChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorStateTimeoutChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateTimeoutChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateTimeoutChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorStateTimeoutChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorStateTimeout")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorStateTimeout_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorStateTimeout_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
