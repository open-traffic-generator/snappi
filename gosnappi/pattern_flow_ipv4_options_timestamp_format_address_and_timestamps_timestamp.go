package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp *****
type patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp struct {
	validation
	obj             *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	marshaller      marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	unMarshaller    unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	incrementHolder PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	decrementHolder PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
}

func NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp {
	obj := patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp{obj: &otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
}

type marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
}

type unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp()
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

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp is a 32-bit value representing the time of packet processing, recorded as the number of milliseconds elapsed since midnight Universal Time (UT).
type PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	// validate validates PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	Choice() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum
	// setChoice assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	setChoice(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	// HasChoice checks if Choice has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	SetValue(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	// HasValue checks if Value has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	SetValues(value []uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	// Increment returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter is integer counter pattern
	Increment() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	// SetIncrement assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	// HasIncrement checks if Increment has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter is integer counter pattern
	Decrement() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	// SetDecrement assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum string

// Enum of Choice on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
var PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice = struct {
	VALUE     PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum
	VALUES    PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum
	INCREMENT PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum
	DECREMENT PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum
}{
	VALUE:     PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum("value"),
	VALUES:    PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) Choice() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum {
	return PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) setChoice(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp {
	intValue, ok := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter().msg()
	}

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) SetValue(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) SetValues(value []uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) Increment() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) SetIncrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) Decrement() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) SetDecrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp")
			}
		} else {
			intVal := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
