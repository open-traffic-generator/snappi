package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime *****
type patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	marshaller      marshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	unMarshaller    unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	incrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	decrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime() PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) (PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) (PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime()
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime is length of time in seconds that addresses generated from the prefix via stateless address autoconfiguration remain preferred. Must be less than or equal to valid_lifetime.
type PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum, set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	Choice() PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	// Increment returns PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
var PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice = struct {
	VALUE     PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum
	VALUES    PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) Choice() PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum {
	return PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime {
	intValue, ok := otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.VALUE {
		defaultValue := uint32(10)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.VALUES {
		defaultValue := []uint32{10}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{10})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) Increment() PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
