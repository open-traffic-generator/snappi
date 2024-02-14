package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRsvpTimeToLive *****
type patternFlowRsvpTimeToLive struct {
	validation
	obj             *otg.PatternFlowRsvpTimeToLive
	marshaller      marshalPatternFlowRsvpTimeToLive
	unMarshaller    unMarshalPatternFlowRsvpTimeToLive
	incrementHolder PatternFlowRsvpTimeToLiveCounter
	decrementHolder PatternFlowRsvpTimeToLiveCounter
}

func NewPatternFlowRsvpTimeToLive() PatternFlowRsvpTimeToLive {
	obj := patternFlowRsvpTimeToLive{obj: &otg.PatternFlowRsvpTimeToLive{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRsvpTimeToLive) msg() *otg.PatternFlowRsvpTimeToLive {
	return obj.obj
}

func (obj *patternFlowRsvpTimeToLive) setMsg(msg *otg.PatternFlowRsvpTimeToLive) PatternFlowRsvpTimeToLive {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRsvpTimeToLive struct {
	obj *patternFlowRsvpTimeToLive
}

type marshalPatternFlowRsvpTimeToLive interface {
	// ToProto marshals PatternFlowRsvpTimeToLive to protobuf object *otg.PatternFlowRsvpTimeToLive
	ToProto() (*otg.PatternFlowRsvpTimeToLive, error)
	// ToPbText marshals PatternFlowRsvpTimeToLive to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRsvpTimeToLive to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRsvpTimeToLive to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRsvpTimeToLive struct {
	obj *patternFlowRsvpTimeToLive
}

type unMarshalPatternFlowRsvpTimeToLive interface {
	// FromProto unmarshals PatternFlowRsvpTimeToLive from protobuf object *otg.PatternFlowRsvpTimeToLive
	FromProto(msg *otg.PatternFlowRsvpTimeToLive) (PatternFlowRsvpTimeToLive, error)
	// FromPbText unmarshals PatternFlowRsvpTimeToLive from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRsvpTimeToLive from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRsvpTimeToLive from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRsvpTimeToLive) Marshal() marshalPatternFlowRsvpTimeToLive {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRsvpTimeToLive{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRsvpTimeToLive) Unmarshal() unMarshalPatternFlowRsvpTimeToLive {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRsvpTimeToLive{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRsvpTimeToLive) ToProto() (*otg.PatternFlowRsvpTimeToLive, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRsvpTimeToLive) FromProto(msg *otg.PatternFlowRsvpTimeToLive) (PatternFlowRsvpTimeToLive, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRsvpTimeToLive) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRsvpTimeToLive) FromPbText(value string) error {
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

func (m *marshalpatternFlowRsvpTimeToLive) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRsvpTimeToLive) FromYaml(value string) error {
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

func (m *marshalpatternFlowRsvpTimeToLive) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRsvpTimeToLive) FromJson(value string) error {
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

func (obj *patternFlowRsvpTimeToLive) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRsvpTimeToLive) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRsvpTimeToLive) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRsvpTimeToLive) Clone() (PatternFlowRsvpTimeToLive, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRsvpTimeToLive()
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

func (obj *patternFlowRsvpTimeToLive) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRsvpTimeToLive is the IP time-to-live(TTL) value with which the message was sent.
type PatternFlowRsvpTimeToLive interface {
	Validation
	// msg marshals PatternFlowRsvpTimeToLive to protobuf object *otg.PatternFlowRsvpTimeToLive
	// and doesn't set defaults
	msg() *otg.PatternFlowRsvpTimeToLive
	// setMsg unmarshals PatternFlowRsvpTimeToLive from protobuf object *otg.PatternFlowRsvpTimeToLive
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRsvpTimeToLive) PatternFlowRsvpTimeToLive
	// provides marshal interface
	Marshal() marshalPatternFlowRsvpTimeToLive
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRsvpTimeToLive
	// validate validates PatternFlowRsvpTimeToLive
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRsvpTimeToLive, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRsvpTimeToLiveChoiceEnum, set in PatternFlowRsvpTimeToLive
	Choice() PatternFlowRsvpTimeToLiveChoiceEnum
	// setChoice assigns PatternFlowRsvpTimeToLiveChoiceEnum provided by user to PatternFlowRsvpTimeToLive
	setChoice(value PatternFlowRsvpTimeToLiveChoiceEnum) PatternFlowRsvpTimeToLive
	// HasChoice checks if Choice has been set in PatternFlowRsvpTimeToLive
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRsvpTimeToLive.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRsvpTimeToLive
	SetValue(value uint32) PatternFlowRsvpTimeToLive
	// HasValue checks if Value has been set in PatternFlowRsvpTimeToLive
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRsvpTimeToLive.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRsvpTimeToLive
	SetValues(value []uint32) PatternFlowRsvpTimeToLive
	// Increment returns PatternFlowRsvpTimeToLiveCounter, set in PatternFlowRsvpTimeToLive.
	// PatternFlowRsvpTimeToLiveCounter is integer counter pattern
	Increment() PatternFlowRsvpTimeToLiveCounter
	// SetIncrement assigns PatternFlowRsvpTimeToLiveCounter provided by user to PatternFlowRsvpTimeToLive.
	// PatternFlowRsvpTimeToLiveCounter is integer counter pattern
	SetIncrement(value PatternFlowRsvpTimeToLiveCounter) PatternFlowRsvpTimeToLive
	// HasIncrement checks if Increment has been set in PatternFlowRsvpTimeToLive
	HasIncrement() bool
	// Decrement returns PatternFlowRsvpTimeToLiveCounter, set in PatternFlowRsvpTimeToLive.
	// PatternFlowRsvpTimeToLiveCounter is integer counter pattern
	Decrement() PatternFlowRsvpTimeToLiveCounter
	// SetDecrement assigns PatternFlowRsvpTimeToLiveCounter provided by user to PatternFlowRsvpTimeToLive.
	// PatternFlowRsvpTimeToLiveCounter is integer counter pattern
	SetDecrement(value PatternFlowRsvpTimeToLiveCounter) PatternFlowRsvpTimeToLive
	// HasDecrement checks if Decrement has been set in PatternFlowRsvpTimeToLive
	HasDecrement() bool
	setNil()
}

type PatternFlowRsvpTimeToLiveChoiceEnum string

// Enum of Choice on PatternFlowRsvpTimeToLive
var PatternFlowRsvpTimeToLiveChoice = struct {
	VALUE     PatternFlowRsvpTimeToLiveChoiceEnum
	VALUES    PatternFlowRsvpTimeToLiveChoiceEnum
	INCREMENT PatternFlowRsvpTimeToLiveChoiceEnum
	DECREMENT PatternFlowRsvpTimeToLiveChoiceEnum
}{
	VALUE:     PatternFlowRsvpTimeToLiveChoiceEnum("value"),
	VALUES:    PatternFlowRsvpTimeToLiveChoiceEnum("values"),
	INCREMENT: PatternFlowRsvpTimeToLiveChoiceEnum("increment"),
	DECREMENT: PatternFlowRsvpTimeToLiveChoiceEnum("decrement"),
}

func (obj *patternFlowRsvpTimeToLive) Choice() PatternFlowRsvpTimeToLiveChoiceEnum {
	return PatternFlowRsvpTimeToLiveChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRsvpTimeToLive) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRsvpTimeToLive) setChoice(value PatternFlowRsvpTimeToLiveChoiceEnum) PatternFlowRsvpTimeToLive {
	intValue, ok := otg.PatternFlowRsvpTimeToLive_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRsvpTimeToLiveChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRsvpTimeToLive_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRsvpTimeToLiveChoice.VALUE {
		defaultValue := uint32(64)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRsvpTimeToLiveChoice.VALUES {
		defaultValue := []uint32{64}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRsvpTimeToLiveChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRsvpTimeToLiveCounter().msg()
	}

	if value == PatternFlowRsvpTimeToLiveChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRsvpTimeToLiveCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRsvpTimeToLive) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRsvpTimeToLiveChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRsvpTimeToLive) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRsvpTimeToLive object
func (obj *patternFlowRsvpTimeToLive) SetValue(value uint32) PatternFlowRsvpTimeToLive {
	obj.setChoice(PatternFlowRsvpTimeToLiveChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRsvpTimeToLive) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{64})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRsvpTimeToLive object
func (obj *patternFlowRsvpTimeToLive) SetValues(value []uint32) PatternFlowRsvpTimeToLive {
	obj.setChoice(PatternFlowRsvpTimeToLiveChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRsvpTimeToLiveCounter
func (obj *patternFlowRsvpTimeToLive) Increment() PatternFlowRsvpTimeToLiveCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRsvpTimeToLiveChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRsvpTimeToLiveCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRsvpTimeToLiveCounter
func (obj *patternFlowRsvpTimeToLive) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRsvpTimeToLiveCounter value in the PatternFlowRsvpTimeToLive object
func (obj *patternFlowRsvpTimeToLive) SetIncrement(value PatternFlowRsvpTimeToLiveCounter) PatternFlowRsvpTimeToLive {
	obj.setChoice(PatternFlowRsvpTimeToLiveChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRsvpTimeToLiveCounter
func (obj *patternFlowRsvpTimeToLive) Decrement() PatternFlowRsvpTimeToLiveCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRsvpTimeToLiveChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRsvpTimeToLiveCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRsvpTimeToLiveCounter
func (obj *patternFlowRsvpTimeToLive) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRsvpTimeToLiveCounter value in the PatternFlowRsvpTimeToLive object
func (obj *patternFlowRsvpTimeToLive) SetDecrement(value PatternFlowRsvpTimeToLiveCounter) PatternFlowRsvpTimeToLive {
	obj.setChoice(PatternFlowRsvpTimeToLiveChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRsvpTimeToLive) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRsvpTimeToLive.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRsvpTimeToLive.Values <= 255 but Got %d", item))
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

func (obj *patternFlowRsvpTimeToLive) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRsvpTimeToLiveChoice.VALUE)

	}

}
