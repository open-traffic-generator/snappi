package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress *****
type patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress struct {
	validation
	obj             *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	marshaller      marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	unMarshaller    unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	incrementHolder PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	decrementHolder PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
}

func NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress {
	obj := patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress{obj: &otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
}

type marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
}

type unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress()
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

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress is iPv4 address of the router interface that handled the packet, serving as a network identifier for the corresponding timestamp entry.
type PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	// validate validates PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	Choice() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum
	// setChoice assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	setChoice(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	// HasChoice checks if Choice has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	SetValue(value string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	// HasValue checks if Value has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	SetValues(value []string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	// Increment returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter is ipv4 counter pattern
	Increment() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	// SetIncrement assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	// HasIncrement checks if Increment has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter is ipv4 counter pattern
	Decrement() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	// SetDecrement assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum string

// Enum of Choice on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
var PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice = struct {
	VALUE     PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum
	VALUES    PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum
	INCREMENT PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum
	DECREMENT PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum
}{
	VALUE:     PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum("value"),
	VALUES:    PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) Choice() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum {
	return PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) setChoice(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress {
	intValue, ok := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter().msg()
	}

	if value == PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) SetValue(value string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) SetValues(value []string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) Increment() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) SetIncrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) Decrement() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) SetDecrement(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress")
			}
		} else {
			intVal := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
