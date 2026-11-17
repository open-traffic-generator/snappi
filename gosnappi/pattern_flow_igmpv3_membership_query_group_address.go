package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryGroupAddress *****
type patternFlowIgmpv3MembershipQueryGroupAddress struct {
	validation
	obj             *otg.PatternFlowIgmpv3MembershipQueryGroupAddress
	marshaller      marshalPatternFlowIgmpv3MembershipQueryGroupAddress
	unMarshaller    unMarshalPatternFlowIgmpv3MembershipQueryGroupAddress
	incrementHolder PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	decrementHolder PatternFlowIgmpv3MembershipQueryGroupAddressCounter
}

func NewPatternFlowIgmpv3MembershipQueryGroupAddress() PatternFlowIgmpv3MembershipQueryGroupAddress {
	obj := patternFlowIgmpv3MembershipQueryGroupAddress{obj: &otg.PatternFlowIgmpv3MembershipQueryGroupAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) msg() *otg.PatternFlowIgmpv3MembershipQueryGroupAddress {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryGroupAddress) PatternFlowIgmpv3MembershipQueryGroupAddress {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryGroupAddress struct {
	obj *patternFlowIgmpv3MembershipQueryGroupAddress
}

type marshalPatternFlowIgmpv3MembershipQueryGroupAddress interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryGroupAddress to protobuf object *otg.PatternFlowIgmpv3MembershipQueryGroupAddress
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryGroupAddress, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryGroupAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryGroupAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryGroupAddress to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryGroupAddress struct {
	obj *patternFlowIgmpv3MembershipQueryGroupAddress
}

type unMarshalPatternFlowIgmpv3MembershipQueryGroupAddress interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryGroupAddress from protobuf object *otg.PatternFlowIgmpv3MembershipQueryGroupAddress
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryGroupAddress) (PatternFlowIgmpv3MembershipQueryGroupAddress, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryGroupAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryGroupAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryGroupAddress from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) Marshal() marshalPatternFlowIgmpv3MembershipQueryGroupAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryGroupAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryGroupAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryGroupAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryGroupAddress) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryGroupAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryGroupAddress) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryGroupAddress) (PatternFlowIgmpv3MembershipQueryGroupAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryGroupAddress) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryGroupAddress) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryGroupAddress) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryGroupAddress) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryGroupAddress) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryGroupAddress) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) Clone() (PatternFlowIgmpv3MembershipQueryGroupAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryGroupAddress()
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

func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3MembershipQueryGroupAddress is group address, zero for a general query and the group being queried for a group-specific or group-and-source-specific query.
type PatternFlowIgmpv3MembershipQueryGroupAddress interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryGroupAddress to protobuf object *otg.PatternFlowIgmpv3MembershipQueryGroupAddress
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryGroupAddress
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryGroupAddress from protobuf object *otg.PatternFlowIgmpv3MembershipQueryGroupAddress
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryGroupAddress) PatternFlowIgmpv3MembershipQueryGroupAddress
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryGroupAddress
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryGroupAddress
	// validate validates PatternFlowIgmpv3MembershipQueryGroupAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryGroupAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum, set in PatternFlowIgmpv3MembershipQueryGroupAddress
	Choice() PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum provided by user to PatternFlowIgmpv3MembershipQueryGroupAddress
	setChoice(value PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum) PatternFlowIgmpv3MembershipQueryGroupAddress
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipQueryGroupAddress
	HasChoice() bool
	// Value returns string, set in PatternFlowIgmpv3MembershipQueryGroupAddress.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIgmpv3MembershipQueryGroupAddress
	SetValue(value string) PatternFlowIgmpv3MembershipQueryGroupAddress
	// HasValue checks if Value has been set in PatternFlowIgmpv3MembershipQueryGroupAddress
	HasValue() bool
	// Values returns []string, set in PatternFlowIgmpv3MembershipQueryGroupAddress.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIgmpv3MembershipQueryGroupAddress
	SetValues(value []string) PatternFlowIgmpv3MembershipQueryGroupAddress
	// Increment returns PatternFlowIgmpv3MembershipQueryGroupAddressCounter, set in PatternFlowIgmpv3MembershipQueryGroupAddress.
	// PatternFlowIgmpv3MembershipQueryGroupAddressCounter is ipv4 counter pattern
	Increment() PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	// SetIncrement assigns PatternFlowIgmpv3MembershipQueryGroupAddressCounter provided by user to PatternFlowIgmpv3MembershipQueryGroupAddress.
	// PatternFlowIgmpv3MembershipQueryGroupAddressCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowIgmpv3MembershipQueryGroupAddressCounter) PatternFlowIgmpv3MembershipQueryGroupAddress
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3MembershipQueryGroupAddress
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3MembershipQueryGroupAddressCounter, set in PatternFlowIgmpv3MembershipQueryGroupAddress.
	// PatternFlowIgmpv3MembershipQueryGroupAddressCounter is ipv4 counter pattern
	Decrement() PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	// SetDecrement assigns PatternFlowIgmpv3MembershipQueryGroupAddressCounter provided by user to PatternFlowIgmpv3MembershipQueryGroupAddress.
	// PatternFlowIgmpv3MembershipQueryGroupAddressCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowIgmpv3MembershipQueryGroupAddressCounter) PatternFlowIgmpv3MembershipQueryGroupAddress
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3MembershipQueryGroupAddress
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipQueryGroupAddress
var PatternFlowIgmpv3MembershipQueryGroupAddressChoice = struct {
	VALUE     PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum
	VALUES    PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum
	INCREMENT PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum
	DECREMENT PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) Choice() PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum {
	return PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) setChoice(value PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum) PatternFlowIgmpv3MembershipQueryGroupAddress {
	intValue, ok := otg.PatternFlowIgmpv3MembershipQueryGroupAddress_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipQueryGroupAddress_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3MembershipQueryGroupAddressChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryGroupAddressChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryGroupAddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3MembershipQueryGroupAddressCounter().msg()
	}

	if value == PatternFlowIgmpv3MembershipQueryGroupAddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3MembershipQueryGroupAddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryGroupAddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIgmpv3MembershipQueryGroupAddress object
func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) SetValue(value string) PatternFlowIgmpv3MembershipQueryGroupAddress {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryGroupAddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIgmpv3MembershipQueryGroupAddress object
func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) SetValues(value []string) PatternFlowIgmpv3MembershipQueryGroupAddress {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryGroupAddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryGroupAddressCounter
func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) Increment() PatternFlowIgmpv3MembershipQueryGroupAddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryGroupAddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3MembershipQueryGroupAddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryGroupAddressCounter
func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3MembershipQueryGroupAddressCounter value in the PatternFlowIgmpv3MembershipQueryGroupAddress object
func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) SetIncrement(value PatternFlowIgmpv3MembershipQueryGroupAddressCounter) PatternFlowIgmpv3MembershipQueryGroupAddress {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryGroupAddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryGroupAddressCounter
func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) Decrement() PatternFlowIgmpv3MembershipQueryGroupAddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryGroupAddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3MembershipQueryGroupAddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryGroupAddressCounter
func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3MembershipQueryGroupAddressCounter value in the PatternFlowIgmpv3MembershipQueryGroupAddress object
func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) SetDecrement(value PatternFlowIgmpv3MembershipQueryGroupAddressCounter) PatternFlowIgmpv3MembershipQueryGroupAddress {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryGroupAddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv3MembershipQueryGroupAddress.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv3MembershipQueryGroupAddress.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddress) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipQueryGroupAddressChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryGroupAddressChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryGroupAddressChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryGroupAddressChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryGroupAddressChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipQueryGroupAddressChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipQueryGroupAddress")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipQueryGroupAddress_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipQueryGroupAddress_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
