package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NeighborSolicitationTargetAddress *****
type patternFlowIcmpv6NeighborSolicitationTargetAddress struct {
	validation
	obj             *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress
	marshaller      marshalPatternFlowIcmpv6NeighborSolicitationTargetAddress
	unMarshaller    unMarshalPatternFlowIcmpv6NeighborSolicitationTargetAddress
	incrementHolder PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	decrementHolder PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
}

func NewPatternFlowIcmpv6NeighborSolicitationTargetAddress() PatternFlowIcmpv6NeighborSolicitationTargetAddress {
	obj := patternFlowIcmpv6NeighborSolicitationTargetAddress{obj: &otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) msg() *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress {
	return obj.obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) setMsg(msg *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress) PatternFlowIcmpv6NeighborSolicitationTargetAddress {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NeighborSolicitationTargetAddress struct {
	obj *patternFlowIcmpv6NeighborSolicitationTargetAddress
}

type marshalPatternFlowIcmpv6NeighborSolicitationTargetAddress interface {
	// ToProto marshals PatternFlowIcmpv6NeighborSolicitationTargetAddress to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress
	ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress, error)
	// ToPbText marshals PatternFlowIcmpv6NeighborSolicitationTargetAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NeighborSolicitationTargetAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NeighborSolicitationTargetAddress to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NeighborSolicitationTargetAddress struct {
	obj *patternFlowIcmpv6NeighborSolicitationTargetAddress
}

type unMarshalPatternFlowIcmpv6NeighborSolicitationTargetAddress interface {
	// FromProto unmarshals PatternFlowIcmpv6NeighborSolicitationTargetAddress from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress
	FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress) (PatternFlowIcmpv6NeighborSolicitationTargetAddress, error)
	// FromPbText unmarshals PatternFlowIcmpv6NeighborSolicitationTargetAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NeighborSolicitationTargetAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NeighborSolicitationTargetAddress from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) Marshal() marshalPatternFlowIcmpv6NeighborSolicitationTargetAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NeighborSolicitationTargetAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationTargetAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NeighborSolicitationTargetAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationTargetAddress) ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationTargetAddress) FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress) (PatternFlowIcmpv6NeighborSolicitationTargetAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationTargetAddress) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationTargetAddress) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationTargetAddress) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationTargetAddress) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationTargetAddress) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationTargetAddress) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) Clone() (PatternFlowIcmpv6NeighborSolicitationTargetAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NeighborSolicitationTargetAddress()
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

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NeighborSolicitationTargetAddress is the IPv6 address of the target of the solicitation. Must not be a multicast address.
type PatternFlowIcmpv6NeighborSolicitationTargetAddress interface {
	Validation
	// msg marshals PatternFlowIcmpv6NeighborSolicitationTargetAddress to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress
	// setMsg unmarshals PatternFlowIcmpv6NeighborSolicitationTargetAddress from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress) PatternFlowIcmpv6NeighborSolicitationTargetAddress
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NeighborSolicitationTargetAddress
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationTargetAddress
	// validate validates PatternFlowIcmpv6NeighborSolicitationTargetAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NeighborSolicitationTargetAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum, set in PatternFlowIcmpv6NeighborSolicitationTargetAddress
	Choice() PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum provided by user to PatternFlowIcmpv6NeighborSolicitationTargetAddress
	setChoice(value PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum) PatternFlowIcmpv6NeighborSolicitationTargetAddress
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NeighborSolicitationTargetAddress
	HasChoice() bool
	// Value returns string, set in PatternFlowIcmpv6NeighborSolicitationTargetAddress.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIcmpv6NeighborSolicitationTargetAddress
	SetValue(value string) PatternFlowIcmpv6NeighborSolicitationTargetAddress
	// HasValue checks if Value has been set in PatternFlowIcmpv6NeighborSolicitationTargetAddress
	HasValue() bool
	// Values returns []string, set in PatternFlowIcmpv6NeighborSolicitationTargetAddress.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIcmpv6NeighborSolicitationTargetAddress
	SetValues(value []string) PatternFlowIcmpv6NeighborSolicitationTargetAddress
	// Increment returns PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter, set in PatternFlowIcmpv6NeighborSolicitationTargetAddress.
	// PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter is ipv6 counter pattern
	Increment() PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	// SetIncrement assigns PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter provided by user to PatternFlowIcmpv6NeighborSolicitationTargetAddress.
	// PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter is ipv6 counter pattern
	SetIncrement(value PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) PatternFlowIcmpv6NeighborSolicitationTargetAddress
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NeighborSolicitationTargetAddress
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter, set in PatternFlowIcmpv6NeighborSolicitationTargetAddress.
	// PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter is ipv6 counter pattern
	Decrement() PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	// SetDecrement assigns PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter provided by user to PatternFlowIcmpv6NeighborSolicitationTargetAddress.
	// PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter is ipv6 counter pattern
	SetDecrement(value PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) PatternFlowIcmpv6NeighborSolicitationTargetAddress
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NeighborSolicitationTargetAddress
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NeighborSolicitationTargetAddress
var PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice = struct {
	VALUE     PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum
	VALUES    PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum
	INCREMENT PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum
	DECREMENT PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) Choice() PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum {
	return PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) setChoice(value PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum) PatternFlowIcmpv6NeighborSolicitationTargetAddress {
	intValue, ok := otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.VALUE {
		defaultValue := "::1"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.VALUES {
		defaultValue := []string{"::1"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NeighborSolicitationTargetAddressCounter().msg()
	}

	if value == PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NeighborSolicitationTargetAddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIcmpv6NeighborSolicitationTargetAddress object
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) SetValue(value string) PatternFlowIcmpv6NeighborSolicitationTargetAddress {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::1"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIcmpv6NeighborSolicitationTargetAddress object
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) SetValues(value []string) PatternFlowIcmpv6NeighborSolicitationTargetAddress {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) Increment() PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NeighborSolicitationTargetAddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter value in the PatternFlowIcmpv6NeighborSolicitationTargetAddress object
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) SetIncrement(value PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) PatternFlowIcmpv6NeighborSolicitationTargetAddress {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) Decrement() PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NeighborSolicitationTargetAddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter value in the PatternFlowIcmpv6NeighborSolicitationTargetAddress object
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) SetDecrement(value PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) PatternFlowIcmpv6NeighborSolicitationTargetAddress {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIcmpv6NeighborSolicitationTargetAddress.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIcmpv6NeighborSolicitationTargetAddress.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddress) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NeighborSolicitationTargetAddressChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTargetAddressChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NeighborSolicitationTargetAddress")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NeighborSolicitationTargetAddress_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
