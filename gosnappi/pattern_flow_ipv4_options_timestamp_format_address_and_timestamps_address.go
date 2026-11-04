package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress *****
type patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress struct {
	validation
	obj             *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	marshaller      marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	unMarshaller    unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	incrementHolder PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	decrementHolder PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
}

func NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress {
	obj := patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress{obj: &otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
}

type marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
}

type unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress()
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

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress is iPv4 address of the router interface that handled the packet, serving as a network identifier for the corresponding timestamp entry.
type PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	// validate validates PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	Choice() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum
	// setChoice assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	setChoice(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	// HasChoice checks if Choice has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	SetValue(value string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	// HasValue checks if Value has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	SetValues(value []string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	// Increment returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter is ipv4 counter pattern
	Increment() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	// SetIncrement assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	// HasIncrement checks if Increment has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter is ipv4 counter pattern
	Decrement() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	// SetDecrement assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum string

// Enum of Choice on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
var PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice = struct {
	VALUE     PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum
	VALUES    PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum
	INCREMENT PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum
	DECREMENT PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum
}{
	VALUE:     PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum("value"),
	VALUES:    PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) Choice() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum {
	return PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) setChoice(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress {
	intValue, ok := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter().msg()
	}

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) SetValue(value string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) SetValues(value []string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) Increment() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) SetIncrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) Decrement() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) SetDecrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress")
			}
		} else {
			intVal := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
