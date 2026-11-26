package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3SourceAddress *****
type patternFlowIgmpv3SourceAddress struct {
	validation
	obj             *otg.PatternFlowIgmpv3SourceAddress
	marshaller      marshalPatternFlowIgmpv3SourceAddress
	unMarshaller    unMarshalPatternFlowIgmpv3SourceAddress
	incrementHolder PatternFlowIgmpv3SourceAddressCounter
	decrementHolder PatternFlowIgmpv3SourceAddressCounter
}

func NewPatternFlowIgmpv3SourceAddress() PatternFlowIgmpv3SourceAddress {
	obj := patternFlowIgmpv3SourceAddress{obj: &otg.PatternFlowIgmpv3SourceAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3SourceAddress) msg() *otg.PatternFlowIgmpv3SourceAddress {
	return obj.obj
}

func (obj *patternFlowIgmpv3SourceAddress) setMsg(msg *otg.PatternFlowIgmpv3SourceAddress) PatternFlowIgmpv3SourceAddress {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3SourceAddress struct {
	obj *patternFlowIgmpv3SourceAddress
}

type marshalPatternFlowIgmpv3SourceAddress interface {
	// ToProto marshals PatternFlowIgmpv3SourceAddress to protobuf object *otg.PatternFlowIgmpv3SourceAddress
	ToProto() (*otg.PatternFlowIgmpv3SourceAddress, error)
	// ToPbText marshals PatternFlowIgmpv3SourceAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3SourceAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3SourceAddress to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3SourceAddress struct {
	obj *patternFlowIgmpv3SourceAddress
}

type unMarshalPatternFlowIgmpv3SourceAddress interface {
	// FromProto unmarshals PatternFlowIgmpv3SourceAddress from protobuf object *otg.PatternFlowIgmpv3SourceAddress
	FromProto(msg *otg.PatternFlowIgmpv3SourceAddress) (PatternFlowIgmpv3SourceAddress, error)
	// FromPbText unmarshals PatternFlowIgmpv3SourceAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3SourceAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3SourceAddress from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3SourceAddress) Marshal() marshalPatternFlowIgmpv3SourceAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3SourceAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3SourceAddress) Unmarshal() unMarshalPatternFlowIgmpv3SourceAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3SourceAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3SourceAddress) ToProto() (*otg.PatternFlowIgmpv3SourceAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3SourceAddress) FromProto(msg *otg.PatternFlowIgmpv3SourceAddress) (PatternFlowIgmpv3SourceAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3SourceAddress) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3SourceAddress) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3SourceAddress) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3SourceAddress) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3SourceAddress) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3SourceAddress) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3SourceAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3SourceAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3SourceAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3SourceAddress) Clone() (PatternFlowIgmpv3SourceAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3SourceAddress()
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

func (obj *patternFlowIgmpv3SourceAddress) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3SourceAddress is unicast IPv4 source address. A counter or a values list changes the address from packet to packet across the flow, it does not add sources to the message.
type PatternFlowIgmpv3SourceAddress interface {
	Validation
	// msg marshals PatternFlowIgmpv3SourceAddress to protobuf object *otg.PatternFlowIgmpv3SourceAddress
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3SourceAddress
	// setMsg unmarshals PatternFlowIgmpv3SourceAddress from protobuf object *otg.PatternFlowIgmpv3SourceAddress
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3SourceAddress) PatternFlowIgmpv3SourceAddress
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3SourceAddress
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3SourceAddress
	// validate validates PatternFlowIgmpv3SourceAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3SourceAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3SourceAddressChoiceEnum, set in PatternFlowIgmpv3SourceAddress
	Choice() PatternFlowIgmpv3SourceAddressChoiceEnum
	// setChoice assigns PatternFlowIgmpv3SourceAddressChoiceEnum provided by user to PatternFlowIgmpv3SourceAddress
	setChoice(value PatternFlowIgmpv3SourceAddressChoiceEnum) PatternFlowIgmpv3SourceAddress
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3SourceAddress
	HasChoice() bool
	// Value returns string, set in PatternFlowIgmpv3SourceAddress.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIgmpv3SourceAddress
	SetValue(value string) PatternFlowIgmpv3SourceAddress
	// HasValue checks if Value has been set in PatternFlowIgmpv3SourceAddress
	HasValue() bool
	// Values returns []string, set in PatternFlowIgmpv3SourceAddress.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIgmpv3SourceAddress
	SetValues(value []string) PatternFlowIgmpv3SourceAddress
	// Increment returns PatternFlowIgmpv3SourceAddressCounter, set in PatternFlowIgmpv3SourceAddress.
	// PatternFlowIgmpv3SourceAddressCounter is ipv4 counter pattern
	Increment() PatternFlowIgmpv3SourceAddressCounter
	// SetIncrement assigns PatternFlowIgmpv3SourceAddressCounter provided by user to PatternFlowIgmpv3SourceAddress.
	// PatternFlowIgmpv3SourceAddressCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowIgmpv3SourceAddressCounter) PatternFlowIgmpv3SourceAddress
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3SourceAddress
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3SourceAddressCounter, set in PatternFlowIgmpv3SourceAddress.
	// PatternFlowIgmpv3SourceAddressCounter is ipv4 counter pattern
	Decrement() PatternFlowIgmpv3SourceAddressCounter
	// SetDecrement assigns PatternFlowIgmpv3SourceAddressCounter provided by user to PatternFlowIgmpv3SourceAddress.
	// PatternFlowIgmpv3SourceAddressCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowIgmpv3SourceAddressCounter) PatternFlowIgmpv3SourceAddress
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3SourceAddress
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3SourceAddressChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3SourceAddress
var PatternFlowIgmpv3SourceAddressChoice = struct {
	VALUE     PatternFlowIgmpv3SourceAddressChoiceEnum
	VALUES    PatternFlowIgmpv3SourceAddressChoiceEnum
	INCREMENT PatternFlowIgmpv3SourceAddressChoiceEnum
	DECREMENT PatternFlowIgmpv3SourceAddressChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3SourceAddressChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3SourceAddressChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv3SourceAddressChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3SourceAddressChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3SourceAddress) Choice() PatternFlowIgmpv3SourceAddressChoiceEnum {
	return PatternFlowIgmpv3SourceAddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3SourceAddress) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3SourceAddress) setChoice(value PatternFlowIgmpv3SourceAddressChoiceEnum) PatternFlowIgmpv3SourceAddress {
	intValue, ok := otg.PatternFlowIgmpv3SourceAddress_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3SourceAddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3SourceAddress_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3SourceAddressChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3SourceAddressChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3SourceAddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3SourceAddressCounter().msg()
	}

	if value == PatternFlowIgmpv3SourceAddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3SourceAddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIgmpv3SourceAddress) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3SourceAddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIgmpv3SourceAddress) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIgmpv3SourceAddress object
func (obj *patternFlowIgmpv3SourceAddress) SetValue(value string) PatternFlowIgmpv3SourceAddress {
	obj.setChoice(PatternFlowIgmpv3SourceAddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIgmpv3SourceAddress) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIgmpv3SourceAddress object
func (obj *patternFlowIgmpv3SourceAddress) SetValues(value []string) PatternFlowIgmpv3SourceAddress {
	obj.setChoice(PatternFlowIgmpv3SourceAddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv3SourceAddressCounter
func (obj *patternFlowIgmpv3SourceAddress) Increment() PatternFlowIgmpv3SourceAddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3SourceAddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3SourceAddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3SourceAddressCounter
func (obj *patternFlowIgmpv3SourceAddress) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3SourceAddressCounter value in the PatternFlowIgmpv3SourceAddress object
func (obj *patternFlowIgmpv3SourceAddress) SetIncrement(value PatternFlowIgmpv3SourceAddressCounter) PatternFlowIgmpv3SourceAddress {
	obj.setChoice(PatternFlowIgmpv3SourceAddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3SourceAddressCounter
func (obj *patternFlowIgmpv3SourceAddress) Decrement() PatternFlowIgmpv3SourceAddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3SourceAddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3SourceAddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3SourceAddressCounter
func (obj *patternFlowIgmpv3SourceAddress) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3SourceAddressCounter value in the PatternFlowIgmpv3SourceAddress object
func (obj *patternFlowIgmpv3SourceAddress) SetDecrement(value PatternFlowIgmpv3SourceAddressCounter) PatternFlowIgmpv3SourceAddress {
	obj.setChoice(PatternFlowIgmpv3SourceAddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3SourceAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv3SourceAddress.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv3SourceAddress.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIgmpv3SourceAddress) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3SourceAddressChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3SourceAddressChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3SourceAddressChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3SourceAddressChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3SourceAddressChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3SourceAddressChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3SourceAddress")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3SourceAddress_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3SourceAddress_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
