package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateTimeout *****
type patternFlowLacpActorStateTimeout struct {
	validation
	obj             *otg.PatternFlowLacpActorStateTimeout
	marshaller      marshalPatternFlowLacpActorStateTimeout
	unMarshaller    unMarshalPatternFlowLacpActorStateTimeout
	incrementHolder PatternFlowLacpActorStateTimeoutCounter
	decrementHolder PatternFlowLacpActorStateTimeoutCounter
}

func NewPatternFlowLacpActorStateTimeout() PatternFlowLacpActorStateTimeout {
	obj := patternFlowLacpActorStateTimeout{obj: &otg.PatternFlowLacpActorStateTimeout{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateTimeout) msg() *otg.PatternFlowLacpActorStateTimeout {
	return obj.obj
}

func (obj *patternFlowLacpActorStateTimeout) setMsg(msg *otg.PatternFlowLacpActorStateTimeout) PatternFlowLacpActorStateTimeout {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateTimeout struct {
	obj *patternFlowLacpActorStateTimeout
}

type marshalPatternFlowLacpActorStateTimeout interface {
	// ToProto marshals PatternFlowLacpActorStateTimeout to protobuf object *otg.PatternFlowLacpActorStateTimeout
	ToProto() (*otg.PatternFlowLacpActorStateTimeout, error)
	// ToPbText marshals PatternFlowLacpActorStateTimeout to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateTimeout to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateTimeout to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateTimeout struct {
	obj *patternFlowLacpActorStateTimeout
}

type unMarshalPatternFlowLacpActorStateTimeout interface {
	// FromProto unmarshals PatternFlowLacpActorStateTimeout from protobuf object *otg.PatternFlowLacpActorStateTimeout
	FromProto(msg *otg.PatternFlowLacpActorStateTimeout) (PatternFlowLacpActorStateTimeout, error)
	// FromPbText unmarshals PatternFlowLacpActorStateTimeout from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateTimeout from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateTimeout from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateTimeout) Marshal() marshalPatternFlowLacpActorStateTimeout {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateTimeout{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateTimeout) Unmarshal() unMarshalPatternFlowLacpActorStateTimeout {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateTimeout{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateTimeout) ToProto() (*otg.PatternFlowLacpActorStateTimeout, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateTimeout) FromProto(msg *otg.PatternFlowLacpActorStateTimeout) (PatternFlowLacpActorStateTimeout, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateTimeout) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateTimeout) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateTimeout) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateTimeout) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateTimeout) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateTimeout) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateTimeout) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateTimeout) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateTimeout) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateTimeout) Clone() (PatternFlowLacpActorStateTimeout, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateTimeout()
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

func (obj *patternFlowLacpActorStateTimeout) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorStateTimeout is lACP Timeout (1=Fast timeout, 0=Slow timeout)
type PatternFlowLacpActorStateTimeout interface {
	Validation
	// msg marshals PatternFlowLacpActorStateTimeout to protobuf object *otg.PatternFlowLacpActorStateTimeout
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateTimeout
	// setMsg unmarshals PatternFlowLacpActorStateTimeout from protobuf object *otg.PatternFlowLacpActorStateTimeout
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateTimeout) PatternFlowLacpActorStateTimeout
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateTimeout
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateTimeout
	// validate validates PatternFlowLacpActorStateTimeout
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateTimeout, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorStateTimeoutChoiceEnum, set in PatternFlowLacpActorStateTimeout
	Choice() PatternFlowLacpActorStateTimeoutChoiceEnum
	// setChoice assigns PatternFlowLacpActorStateTimeoutChoiceEnum provided by user to PatternFlowLacpActorStateTimeout
	setChoice(value PatternFlowLacpActorStateTimeoutChoiceEnum) PatternFlowLacpActorStateTimeout
	// HasChoice checks if Choice has been set in PatternFlowLacpActorStateTimeout
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorStateTimeout.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorStateTimeout
	SetValue(value uint32) PatternFlowLacpActorStateTimeout
	// HasValue checks if Value has been set in PatternFlowLacpActorStateTimeout
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorStateTimeout.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorStateTimeout
	SetValues(value []uint32) PatternFlowLacpActorStateTimeout
	// Increment returns PatternFlowLacpActorStateTimeoutCounter, set in PatternFlowLacpActorStateTimeout.
	// PatternFlowLacpActorStateTimeoutCounter is integer counter pattern
	Increment() PatternFlowLacpActorStateTimeoutCounter
	// SetIncrement assigns PatternFlowLacpActorStateTimeoutCounter provided by user to PatternFlowLacpActorStateTimeout.
	// PatternFlowLacpActorStateTimeoutCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorStateTimeoutCounter) PatternFlowLacpActorStateTimeout
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorStateTimeout
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorStateTimeoutCounter, set in PatternFlowLacpActorStateTimeout.
	// PatternFlowLacpActorStateTimeoutCounter is integer counter pattern
	Decrement() PatternFlowLacpActorStateTimeoutCounter
	// SetDecrement assigns PatternFlowLacpActorStateTimeoutCounter provided by user to PatternFlowLacpActorStateTimeout.
	// PatternFlowLacpActorStateTimeoutCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorStateTimeoutCounter) PatternFlowLacpActorStateTimeout
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorStateTimeout
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorStateTimeoutChoiceEnum string

// Enum of Choice on PatternFlowLacpActorStateTimeout
var PatternFlowLacpActorStateTimeoutChoice = struct {
	VALUE     PatternFlowLacpActorStateTimeoutChoiceEnum
	VALUES    PatternFlowLacpActorStateTimeoutChoiceEnum
	INCREMENT PatternFlowLacpActorStateTimeoutChoiceEnum
	DECREMENT PatternFlowLacpActorStateTimeoutChoiceEnum
}{
	VALUE:     PatternFlowLacpActorStateTimeoutChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorStateTimeoutChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorStateTimeoutChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorStateTimeoutChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorStateTimeout) Choice() PatternFlowLacpActorStateTimeoutChoiceEnum {
	return PatternFlowLacpActorStateTimeoutChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorStateTimeout) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorStateTimeout) setChoice(value PatternFlowLacpActorStateTimeoutChoiceEnum) PatternFlowLacpActorStateTimeout {
	intValue, ok := otg.PatternFlowLacpActorStateTimeout_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorStateTimeoutChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorStateTimeout_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorStateTimeoutChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorStateTimeoutChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorStateTimeoutChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorStateTimeoutCounter().msg()
	}

	if value == PatternFlowLacpActorStateTimeoutChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorStateTimeoutCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateTimeout) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorStateTimeoutChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateTimeout) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorStateTimeout object
func (obj *patternFlowLacpActorStateTimeout) SetValue(value uint32) PatternFlowLacpActorStateTimeout {
	obj.setChoice(PatternFlowLacpActorStateTimeoutChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorStateTimeout) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorStateTimeout object
func (obj *patternFlowLacpActorStateTimeout) SetValues(value []uint32) PatternFlowLacpActorStateTimeout {
	obj.setChoice(PatternFlowLacpActorStateTimeoutChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateTimeoutCounter
func (obj *patternFlowLacpActorStateTimeout) Increment() PatternFlowLacpActorStateTimeoutCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorStateTimeoutChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorStateTimeoutCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateTimeoutCounter
func (obj *patternFlowLacpActorStateTimeout) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorStateTimeoutCounter value in the PatternFlowLacpActorStateTimeout object
func (obj *patternFlowLacpActorStateTimeout) SetIncrement(value PatternFlowLacpActorStateTimeoutCounter) PatternFlowLacpActorStateTimeout {
	obj.setChoice(PatternFlowLacpActorStateTimeoutChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateTimeoutCounter
func (obj *patternFlowLacpActorStateTimeout) Decrement() PatternFlowLacpActorStateTimeoutCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorStateTimeoutChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorStateTimeoutCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateTimeoutCounter
func (obj *patternFlowLacpActorStateTimeout) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorStateTimeoutCounter value in the PatternFlowLacpActorStateTimeout object
func (obj *patternFlowLacpActorStateTimeout) SetDecrement(value PatternFlowLacpActorStateTimeoutCounter) PatternFlowLacpActorStateTimeout {
	obj.setChoice(PatternFlowLacpActorStateTimeoutChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorStateTimeout) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateTimeout.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorStateTimeout.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpActorStateTimeout) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorStateTimeoutChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateTimeoutChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorStateTimeoutChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateTimeoutChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateTimeoutChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorStateTimeoutChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorStateTimeout")
			}
		} else {
			intVal := otg.PatternFlowLacpActorStateTimeout_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorStateTimeout_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
