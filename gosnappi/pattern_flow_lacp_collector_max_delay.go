package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpCollectorMaxDelay *****
type patternFlowLacpCollectorMaxDelay struct {
	validation
	obj             *otg.PatternFlowLacpCollectorMaxDelay
	marshaller      marshalPatternFlowLacpCollectorMaxDelay
	unMarshaller    unMarshalPatternFlowLacpCollectorMaxDelay
	incrementHolder PatternFlowLacpCollectorMaxDelayCounter
	decrementHolder PatternFlowLacpCollectorMaxDelayCounter
}

func NewPatternFlowLacpCollectorMaxDelay() PatternFlowLacpCollectorMaxDelay {
	obj := patternFlowLacpCollectorMaxDelay{obj: &otg.PatternFlowLacpCollectorMaxDelay{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpCollectorMaxDelay) msg() *otg.PatternFlowLacpCollectorMaxDelay {
	return obj.obj
}

func (obj *patternFlowLacpCollectorMaxDelay) setMsg(msg *otg.PatternFlowLacpCollectorMaxDelay) PatternFlowLacpCollectorMaxDelay {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpCollectorMaxDelay struct {
	obj *patternFlowLacpCollectorMaxDelay
}

type marshalPatternFlowLacpCollectorMaxDelay interface {
	// ToProto marshals PatternFlowLacpCollectorMaxDelay to protobuf object *otg.PatternFlowLacpCollectorMaxDelay
	ToProto() (*otg.PatternFlowLacpCollectorMaxDelay, error)
	// ToPbText marshals PatternFlowLacpCollectorMaxDelay to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpCollectorMaxDelay to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpCollectorMaxDelay to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpCollectorMaxDelay struct {
	obj *patternFlowLacpCollectorMaxDelay
}

type unMarshalPatternFlowLacpCollectorMaxDelay interface {
	// FromProto unmarshals PatternFlowLacpCollectorMaxDelay from protobuf object *otg.PatternFlowLacpCollectorMaxDelay
	FromProto(msg *otg.PatternFlowLacpCollectorMaxDelay) (PatternFlowLacpCollectorMaxDelay, error)
	// FromPbText unmarshals PatternFlowLacpCollectorMaxDelay from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpCollectorMaxDelay from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpCollectorMaxDelay from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpCollectorMaxDelay) Marshal() marshalPatternFlowLacpCollectorMaxDelay {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpCollectorMaxDelay{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpCollectorMaxDelay) Unmarshal() unMarshalPatternFlowLacpCollectorMaxDelay {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpCollectorMaxDelay{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpCollectorMaxDelay) ToProto() (*otg.PatternFlowLacpCollectorMaxDelay, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpCollectorMaxDelay) FromProto(msg *otg.PatternFlowLacpCollectorMaxDelay) (PatternFlowLacpCollectorMaxDelay, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpCollectorMaxDelay) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorMaxDelay) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpCollectorMaxDelay) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorMaxDelay) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpCollectorMaxDelay) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorMaxDelay) FromJson(value string) error {
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

func (obj *patternFlowLacpCollectorMaxDelay) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpCollectorMaxDelay) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpCollectorMaxDelay) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpCollectorMaxDelay) Clone() (PatternFlowLacpCollectorMaxDelay, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpCollectorMaxDelay()
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

func (obj *patternFlowLacpCollectorMaxDelay) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpCollectorMaxDelay is indicates the maximum delay, in units of 10 microseconds, that the transmitting system's aggregator will take to buffer frames from its collector before emitting a frame.
type PatternFlowLacpCollectorMaxDelay interface {
	Validation
	// msg marshals PatternFlowLacpCollectorMaxDelay to protobuf object *otg.PatternFlowLacpCollectorMaxDelay
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpCollectorMaxDelay
	// setMsg unmarshals PatternFlowLacpCollectorMaxDelay from protobuf object *otg.PatternFlowLacpCollectorMaxDelay
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpCollectorMaxDelay) PatternFlowLacpCollectorMaxDelay
	// provides marshal interface
	Marshal() marshalPatternFlowLacpCollectorMaxDelay
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpCollectorMaxDelay
	// validate validates PatternFlowLacpCollectorMaxDelay
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpCollectorMaxDelay, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpCollectorMaxDelayChoiceEnum, set in PatternFlowLacpCollectorMaxDelay
	Choice() PatternFlowLacpCollectorMaxDelayChoiceEnum
	// setChoice assigns PatternFlowLacpCollectorMaxDelayChoiceEnum provided by user to PatternFlowLacpCollectorMaxDelay
	setChoice(value PatternFlowLacpCollectorMaxDelayChoiceEnum) PatternFlowLacpCollectorMaxDelay
	// HasChoice checks if Choice has been set in PatternFlowLacpCollectorMaxDelay
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpCollectorMaxDelay.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpCollectorMaxDelay
	SetValue(value uint32) PatternFlowLacpCollectorMaxDelay
	// HasValue checks if Value has been set in PatternFlowLacpCollectorMaxDelay
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpCollectorMaxDelay.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpCollectorMaxDelay
	SetValues(value []uint32) PatternFlowLacpCollectorMaxDelay
	// Increment returns PatternFlowLacpCollectorMaxDelayCounter, set in PatternFlowLacpCollectorMaxDelay.
	// PatternFlowLacpCollectorMaxDelayCounter is integer counter pattern
	Increment() PatternFlowLacpCollectorMaxDelayCounter
	// SetIncrement assigns PatternFlowLacpCollectorMaxDelayCounter provided by user to PatternFlowLacpCollectorMaxDelay.
	// PatternFlowLacpCollectorMaxDelayCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpCollectorMaxDelayCounter) PatternFlowLacpCollectorMaxDelay
	// HasIncrement checks if Increment has been set in PatternFlowLacpCollectorMaxDelay
	HasIncrement() bool
	// Decrement returns PatternFlowLacpCollectorMaxDelayCounter, set in PatternFlowLacpCollectorMaxDelay.
	// PatternFlowLacpCollectorMaxDelayCounter is integer counter pattern
	Decrement() PatternFlowLacpCollectorMaxDelayCounter
	// SetDecrement assigns PatternFlowLacpCollectorMaxDelayCounter provided by user to PatternFlowLacpCollectorMaxDelay.
	// PatternFlowLacpCollectorMaxDelayCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpCollectorMaxDelayCounter) PatternFlowLacpCollectorMaxDelay
	// HasDecrement checks if Decrement has been set in PatternFlowLacpCollectorMaxDelay
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpCollectorMaxDelayChoiceEnum string

// Enum of Choice on PatternFlowLacpCollectorMaxDelay
var PatternFlowLacpCollectorMaxDelayChoice = struct {
	VALUE     PatternFlowLacpCollectorMaxDelayChoiceEnum
	VALUES    PatternFlowLacpCollectorMaxDelayChoiceEnum
	INCREMENT PatternFlowLacpCollectorMaxDelayChoiceEnum
	DECREMENT PatternFlowLacpCollectorMaxDelayChoiceEnum
}{
	VALUE:     PatternFlowLacpCollectorMaxDelayChoiceEnum("value"),
	VALUES:    PatternFlowLacpCollectorMaxDelayChoiceEnum("values"),
	INCREMENT: PatternFlowLacpCollectorMaxDelayChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpCollectorMaxDelayChoiceEnum("decrement"),
}

func (obj *patternFlowLacpCollectorMaxDelay) Choice() PatternFlowLacpCollectorMaxDelayChoiceEnum {
	return PatternFlowLacpCollectorMaxDelayChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpCollectorMaxDelay) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpCollectorMaxDelay) setChoice(value PatternFlowLacpCollectorMaxDelayChoiceEnum) PatternFlowLacpCollectorMaxDelay {
	intValue, ok := otg.PatternFlowLacpCollectorMaxDelay_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpCollectorMaxDelayChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpCollectorMaxDelay_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpCollectorMaxDelayChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpCollectorMaxDelayChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpCollectorMaxDelayChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpCollectorMaxDelayCounter().msg()
	}

	if value == PatternFlowLacpCollectorMaxDelayChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpCollectorMaxDelayCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpCollectorMaxDelay) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpCollectorMaxDelayChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpCollectorMaxDelay) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpCollectorMaxDelay object
func (obj *patternFlowLacpCollectorMaxDelay) SetValue(value uint32) PatternFlowLacpCollectorMaxDelay {
	obj.setChoice(PatternFlowLacpCollectorMaxDelayChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpCollectorMaxDelay) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpCollectorMaxDelay object
func (obj *patternFlowLacpCollectorMaxDelay) SetValues(value []uint32) PatternFlowLacpCollectorMaxDelay {
	obj.setChoice(PatternFlowLacpCollectorMaxDelayChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpCollectorMaxDelayCounter
func (obj *patternFlowLacpCollectorMaxDelay) Increment() PatternFlowLacpCollectorMaxDelayCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpCollectorMaxDelayChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpCollectorMaxDelayCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpCollectorMaxDelayCounter
func (obj *patternFlowLacpCollectorMaxDelay) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpCollectorMaxDelayCounter value in the PatternFlowLacpCollectorMaxDelay object
func (obj *patternFlowLacpCollectorMaxDelay) SetIncrement(value PatternFlowLacpCollectorMaxDelayCounter) PatternFlowLacpCollectorMaxDelay {
	obj.setChoice(PatternFlowLacpCollectorMaxDelayChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpCollectorMaxDelayCounter
func (obj *patternFlowLacpCollectorMaxDelay) Decrement() PatternFlowLacpCollectorMaxDelayCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpCollectorMaxDelayChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpCollectorMaxDelayCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpCollectorMaxDelayCounter
func (obj *patternFlowLacpCollectorMaxDelay) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpCollectorMaxDelayCounter value in the PatternFlowLacpCollectorMaxDelay object
func (obj *patternFlowLacpCollectorMaxDelay) SetDecrement(value PatternFlowLacpCollectorMaxDelayCounter) PatternFlowLacpCollectorMaxDelay {
	obj.setChoice(PatternFlowLacpCollectorMaxDelayChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpCollectorMaxDelay) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpCollectorMaxDelay.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpCollectorMaxDelay.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpCollectorMaxDelay) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpCollectorMaxDelayChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpCollectorMaxDelayChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpCollectorMaxDelayChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpCollectorMaxDelayChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpCollectorMaxDelayChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpCollectorMaxDelayChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpCollectorMaxDelay")
			}
		} else {
			intVal := otg.PatternFlowLacpCollectorMaxDelay_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpCollectorMaxDelay_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
