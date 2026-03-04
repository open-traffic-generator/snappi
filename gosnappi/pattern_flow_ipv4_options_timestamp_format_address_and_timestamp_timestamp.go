package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp *****
type patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp struct {
	validation
	obj             *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	marshaller      marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	unMarshaller    unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	incrementHolder PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	decrementHolder PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
}

func NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp {
	obj := patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp{obj: &otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
}

type marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
}

type unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp()
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

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp is a 32-bit value representing the time of packet processing, recorded as the number of milliseconds elapsed since midnight Universal Time (UT).
type PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	// validate validates PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	Choice() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum
	// setChoice assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	setChoice(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	// HasChoice checks if Choice has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	SetValue(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	// HasValue checks if Value has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	SetValues(value []uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	// Increment returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter is integer counter pattern
	Increment() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	// SetIncrement assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	// HasIncrement checks if Increment has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter is integer counter pattern
	Decrement() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	// SetDecrement assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum string

// Enum of Choice on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
var PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice = struct {
	VALUE     PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum
	VALUES    PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum
	INCREMENT PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum
	DECREMENT PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum
}{
	VALUE:     PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum("value"),
	VALUES:    PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) Choice() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum {
	return PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) setChoice(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp {
	intValue, ok := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter().msg()
	}

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) SetValue(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) SetValues(value []uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) Increment() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) SetIncrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) Decrement() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) SetDecrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp")
			}
		} else {
			intVal := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
