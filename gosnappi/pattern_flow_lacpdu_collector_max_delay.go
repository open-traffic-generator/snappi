package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduCollectorMaxDelay *****
type patternFlowLacpduCollectorMaxDelay struct {
	validation
	obj             *otg.PatternFlowLacpduCollectorMaxDelay
	marshaller      marshalPatternFlowLacpduCollectorMaxDelay
	unMarshaller    unMarshalPatternFlowLacpduCollectorMaxDelay
	incrementHolder PatternFlowLacpduCollectorMaxDelayCounter
	decrementHolder PatternFlowLacpduCollectorMaxDelayCounter
}

func NewPatternFlowLacpduCollectorMaxDelay() PatternFlowLacpduCollectorMaxDelay {
	obj := patternFlowLacpduCollectorMaxDelay{obj: &otg.PatternFlowLacpduCollectorMaxDelay{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduCollectorMaxDelay) msg() *otg.PatternFlowLacpduCollectorMaxDelay {
	return obj.obj
}

func (obj *patternFlowLacpduCollectorMaxDelay) setMsg(msg *otg.PatternFlowLacpduCollectorMaxDelay) PatternFlowLacpduCollectorMaxDelay {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduCollectorMaxDelay struct {
	obj *patternFlowLacpduCollectorMaxDelay
}

type marshalPatternFlowLacpduCollectorMaxDelay interface {
	// ToProto marshals PatternFlowLacpduCollectorMaxDelay to protobuf object *otg.PatternFlowLacpduCollectorMaxDelay
	ToProto() (*otg.PatternFlowLacpduCollectorMaxDelay, error)
	// ToPbText marshals PatternFlowLacpduCollectorMaxDelay to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduCollectorMaxDelay to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduCollectorMaxDelay to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduCollectorMaxDelay struct {
	obj *patternFlowLacpduCollectorMaxDelay
}

type unMarshalPatternFlowLacpduCollectorMaxDelay interface {
	// FromProto unmarshals PatternFlowLacpduCollectorMaxDelay from protobuf object *otg.PatternFlowLacpduCollectorMaxDelay
	FromProto(msg *otg.PatternFlowLacpduCollectorMaxDelay) (PatternFlowLacpduCollectorMaxDelay, error)
	// FromPbText unmarshals PatternFlowLacpduCollectorMaxDelay from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduCollectorMaxDelay from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduCollectorMaxDelay from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduCollectorMaxDelay) Marshal() marshalPatternFlowLacpduCollectorMaxDelay {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduCollectorMaxDelay{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduCollectorMaxDelay) Unmarshal() unMarshalPatternFlowLacpduCollectorMaxDelay {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduCollectorMaxDelay{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduCollectorMaxDelay) ToProto() (*otg.PatternFlowLacpduCollectorMaxDelay, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduCollectorMaxDelay) FromProto(msg *otg.PatternFlowLacpduCollectorMaxDelay) (PatternFlowLacpduCollectorMaxDelay, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduCollectorMaxDelay) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorMaxDelay) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduCollectorMaxDelay) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorMaxDelay) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduCollectorMaxDelay) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorMaxDelay) FromJson(value string) error {
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

func (obj *patternFlowLacpduCollectorMaxDelay) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduCollectorMaxDelay) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduCollectorMaxDelay) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduCollectorMaxDelay) Clone() (PatternFlowLacpduCollectorMaxDelay, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduCollectorMaxDelay()
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

func (obj *patternFlowLacpduCollectorMaxDelay) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduCollectorMaxDelay is indicates the maximum delay, in units of 10 microseconds, that the transmitting system's aggregator will take to buffer frames from its collector before emitting a frame.
type PatternFlowLacpduCollectorMaxDelay interface {
	Validation
	// msg marshals PatternFlowLacpduCollectorMaxDelay to protobuf object *otg.PatternFlowLacpduCollectorMaxDelay
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduCollectorMaxDelay
	// setMsg unmarshals PatternFlowLacpduCollectorMaxDelay from protobuf object *otg.PatternFlowLacpduCollectorMaxDelay
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduCollectorMaxDelay) PatternFlowLacpduCollectorMaxDelay
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduCollectorMaxDelay
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduCollectorMaxDelay
	// validate validates PatternFlowLacpduCollectorMaxDelay
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduCollectorMaxDelay, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduCollectorMaxDelayChoiceEnum, set in PatternFlowLacpduCollectorMaxDelay
	Choice() PatternFlowLacpduCollectorMaxDelayChoiceEnum
	// setChoice assigns PatternFlowLacpduCollectorMaxDelayChoiceEnum provided by user to PatternFlowLacpduCollectorMaxDelay
	setChoice(value PatternFlowLacpduCollectorMaxDelayChoiceEnum) PatternFlowLacpduCollectorMaxDelay
	// HasChoice checks if Choice has been set in PatternFlowLacpduCollectorMaxDelay
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduCollectorMaxDelay.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduCollectorMaxDelay
	SetValue(value uint32) PatternFlowLacpduCollectorMaxDelay
	// HasValue checks if Value has been set in PatternFlowLacpduCollectorMaxDelay
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduCollectorMaxDelay.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduCollectorMaxDelay
	SetValues(value []uint32) PatternFlowLacpduCollectorMaxDelay
	// Increment returns PatternFlowLacpduCollectorMaxDelayCounter, set in PatternFlowLacpduCollectorMaxDelay.
	// PatternFlowLacpduCollectorMaxDelayCounter is integer counter pattern
	Increment() PatternFlowLacpduCollectorMaxDelayCounter
	// SetIncrement assigns PatternFlowLacpduCollectorMaxDelayCounter provided by user to PatternFlowLacpduCollectorMaxDelay.
	// PatternFlowLacpduCollectorMaxDelayCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduCollectorMaxDelayCounter) PatternFlowLacpduCollectorMaxDelay
	// HasIncrement checks if Increment has been set in PatternFlowLacpduCollectorMaxDelay
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduCollectorMaxDelayCounter, set in PatternFlowLacpduCollectorMaxDelay.
	// PatternFlowLacpduCollectorMaxDelayCounter is integer counter pattern
	Decrement() PatternFlowLacpduCollectorMaxDelayCounter
	// SetDecrement assigns PatternFlowLacpduCollectorMaxDelayCounter provided by user to PatternFlowLacpduCollectorMaxDelay.
	// PatternFlowLacpduCollectorMaxDelayCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduCollectorMaxDelayCounter) PatternFlowLacpduCollectorMaxDelay
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduCollectorMaxDelay
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduCollectorMaxDelayChoiceEnum string

// Enum of Choice on PatternFlowLacpduCollectorMaxDelay
var PatternFlowLacpduCollectorMaxDelayChoice = struct {
	VALUE     PatternFlowLacpduCollectorMaxDelayChoiceEnum
	VALUES    PatternFlowLacpduCollectorMaxDelayChoiceEnum
	INCREMENT PatternFlowLacpduCollectorMaxDelayChoiceEnum
	DECREMENT PatternFlowLacpduCollectorMaxDelayChoiceEnum
}{
	VALUE:     PatternFlowLacpduCollectorMaxDelayChoiceEnum("value"),
	VALUES:    PatternFlowLacpduCollectorMaxDelayChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduCollectorMaxDelayChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduCollectorMaxDelayChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduCollectorMaxDelay) Choice() PatternFlowLacpduCollectorMaxDelayChoiceEnum {
	return PatternFlowLacpduCollectorMaxDelayChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduCollectorMaxDelay) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduCollectorMaxDelay) setChoice(value PatternFlowLacpduCollectorMaxDelayChoiceEnum) PatternFlowLacpduCollectorMaxDelay {
	intValue, ok := otg.PatternFlowLacpduCollectorMaxDelay_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduCollectorMaxDelayChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduCollectorMaxDelay_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduCollectorMaxDelayChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduCollectorMaxDelayChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduCollectorMaxDelayChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduCollectorMaxDelayCounter().msg()
	}

	if value == PatternFlowLacpduCollectorMaxDelayChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduCollectorMaxDelayCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduCollectorMaxDelay) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduCollectorMaxDelayChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduCollectorMaxDelay) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduCollectorMaxDelay object
func (obj *patternFlowLacpduCollectorMaxDelay) SetValue(value uint32) PatternFlowLacpduCollectorMaxDelay {
	obj.setChoice(PatternFlowLacpduCollectorMaxDelayChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduCollectorMaxDelay) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduCollectorMaxDelay object
func (obj *patternFlowLacpduCollectorMaxDelay) SetValues(value []uint32) PatternFlowLacpduCollectorMaxDelay {
	obj.setChoice(PatternFlowLacpduCollectorMaxDelayChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduCollectorMaxDelayCounter
func (obj *patternFlowLacpduCollectorMaxDelay) Increment() PatternFlowLacpduCollectorMaxDelayCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduCollectorMaxDelayChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduCollectorMaxDelayCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduCollectorMaxDelayCounter
func (obj *patternFlowLacpduCollectorMaxDelay) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduCollectorMaxDelayCounter value in the PatternFlowLacpduCollectorMaxDelay object
func (obj *patternFlowLacpduCollectorMaxDelay) SetIncrement(value PatternFlowLacpduCollectorMaxDelayCounter) PatternFlowLacpduCollectorMaxDelay {
	obj.setChoice(PatternFlowLacpduCollectorMaxDelayChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduCollectorMaxDelayCounter
func (obj *patternFlowLacpduCollectorMaxDelay) Decrement() PatternFlowLacpduCollectorMaxDelayCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduCollectorMaxDelayChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduCollectorMaxDelayCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduCollectorMaxDelayCounter
func (obj *patternFlowLacpduCollectorMaxDelay) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduCollectorMaxDelayCounter value in the PatternFlowLacpduCollectorMaxDelay object
func (obj *patternFlowLacpduCollectorMaxDelay) SetDecrement(value PatternFlowLacpduCollectorMaxDelayCounter) PatternFlowLacpduCollectorMaxDelay {
	obj.setChoice(PatternFlowLacpduCollectorMaxDelayChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduCollectorMaxDelay) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduCollectorMaxDelay.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduCollectorMaxDelay.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpduCollectorMaxDelay) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduCollectorMaxDelayChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduCollectorMaxDelayChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduCollectorMaxDelayChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduCollectorMaxDelayChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduCollectorMaxDelayChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduCollectorMaxDelayChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduCollectorMaxDelay")
			}
		} else {
			intVal := otg.PatternFlowLacpduCollectorMaxDelay_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduCollectorMaxDelay_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
