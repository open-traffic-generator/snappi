package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3GroupRecordMulticastAddress *****
type patternFlowIgmpv3GroupRecordMulticastAddress struct {
	validation
	obj             *otg.PatternFlowIgmpv3GroupRecordMulticastAddress
	marshaller      marshalPatternFlowIgmpv3GroupRecordMulticastAddress
	unMarshaller    unMarshalPatternFlowIgmpv3GroupRecordMulticastAddress
	incrementHolder PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	decrementHolder PatternFlowIgmpv3GroupRecordMulticastAddressCounter
}

func NewPatternFlowIgmpv3GroupRecordMulticastAddress() PatternFlowIgmpv3GroupRecordMulticastAddress {
	obj := patternFlowIgmpv3GroupRecordMulticastAddress{obj: &otg.PatternFlowIgmpv3GroupRecordMulticastAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) msg() *otg.PatternFlowIgmpv3GroupRecordMulticastAddress {
	return obj.obj
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) setMsg(msg *otg.PatternFlowIgmpv3GroupRecordMulticastAddress) PatternFlowIgmpv3GroupRecordMulticastAddress {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3GroupRecordMulticastAddress struct {
	obj *patternFlowIgmpv3GroupRecordMulticastAddress
}

type marshalPatternFlowIgmpv3GroupRecordMulticastAddress interface {
	// ToProto marshals PatternFlowIgmpv3GroupRecordMulticastAddress to protobuf object *otg.PatternFlowIgmpv3GroupRecordMulticastAddress
	ToProto() (*otg.PatternFlowIgmpv3GroupRecordMulticastAddress, error)
	// ToPbText marshals PatternFlowIgmpv3GroupRecordMulticastAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3GroupRecordMulticastAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3GroupRecordMulticastAddress to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3GroupRecordMulticastAddress struct {
	obj *patternFlowIgmpv3GroupRecordMulticastAddress
}

type unMarshalPatternFlowIgmpv3GroupRecordMulticastAddress interface {
	// FromProto unmarshals PatternFlowIgmpv3GroupRecordMulticastAddress from protobuf object *otg.PatternFlowIgmpv3GroupRecordMulticastAddress
	FromProto(msg *otg.PatternFlowIgmpv3GroupRecordMulticastAddress) (PatternFlowIgmpv3GroupRecordMulticastAddress, error)
	// FromPbText unmarshals PatternFlowIgmpv3GroupRecordMulticastAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3GroupRecordMulticastAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3GroupRecordMulticastAddress from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) Marshal() marshalPatternFlowIgmpv3GroupRecordMulticastAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3GroupRecordMulticastAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordMulticastAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3GroupRecordMulticastAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3GroupRecordMulticastAddress) ToProto() (*otg.PatternFlowIgmpv3GroupRecordMulticastAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3GroupRecordMulticastAddress) FromProto(msg *otg.PatternFlowIgmpv3GroupRecordMulticastAddress) (PatternFlowIgmpv3GroupRecordMulticastAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3GroupRecordMulticastAddress) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordMulticastAddress) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordMulticastAddress) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordMulticastAddress) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordMulticastAddress) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordMulticastAddress) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) Clone() (PatternFlowIgmpv3GroupRecordMulticastAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3GroupRecordMulticastAddress()
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

func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3GroupRecordMulticastAddress is multicast group address that this record reports on.
type PatternFlowIgmpv3GroupRecordMulticastAddress interface {
	Validation
	// msg marshals PatternFlowIgmpv3GroupRecordMulticastAddress to protobuf object *otg.PatternFlowIgmpv3GroupRecordMulticastAddress
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3GroupRecordMulticastAddress
	// setMsg unmarshals PatternFlowIgmpv3GroupRecordMulticastAddress from protobuf object *otg.PatternFlowIgmpv3GroupRecordMulticastAddress
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3GroupRecordMulticastAddress) PatternFlowIgmpv3GroupRecordMulticastAddress
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3GroupRecordMulticastAddress
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordMulticastAddress
	// validate validates PatternFlowIgmpv3GroupRecordMulticastAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3GroupRecordMulticastAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum, set in PatternFlowIgmpv3GroupRecordMulticastAddress
	Choice() PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum
	// setChoice assigns PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum provided by user to PatternFlowIgmpv3GroupRecordMulticastAddress
	setChoice(value PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum) PatternFlowIgmpv3GroupRecordMulticastAddress
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3GroupRecordMulticastAddress
	HasChoice() bool
	// Value returns string, set in PatternFlowIgmpv3GroupRecordMulticastAddress.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIgmpv3GroupRecordMulticastAddress
	SetValue(value string) PatternFlowIgmpv3GroupRecordMulticastAddress
	// HasValue checks if Value has been set in PatternFlowIgmpv3GroupRecordMulticastAddress
	HasValue() bool
	// Values returns []string, set in PatternFlowIgmpv3GroupRecordMulticastAddress.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIgmpv3GroupRecordMulticastAddress
	SetValues(value []string) PatternFlowIgmpv3GroupRecordMulticastAddress
	// Increment returns PatternFlowIgmpv3GroupRecordMulticastAddressCounter, set in PatternFlowIgmpv3GroupRecordMulticastAddress.
	// PatternFlowIgmpv3GroupRecordMulticastAddressCounter is ipv4 counter pattern
	Increment() PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	// SetIncrement assigns PatternFlowIgmpv3GroupRecordMulticastAddressCounter provided by user to PatternFlowIgmpv3GroupRecordMulticastAddress.
	// PatternFlowIgmpv3GroupRecordMulticastAddressCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowIgmpv3GroupRecordMulticastAddressCounter) PatternFlowIgmpv3GroupRecordMulticastAddress
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3GroupRecordMulticastAddress
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3GroupRecordMulticastAddressCounter, set in PatternFlowIgmpv3GroupRecordMulticastAddress.
	// PatternFlowIgmpv3GroupRecordMulticastAddressCounter is ipv4 counter pattern
	Decrement() PatternFlowIgmpv3GroupRecordMulticastAddressCounter
	// SetDecrement assigns PatternFlowIgmpv3GroupRecordMulticastAddressCounter provided by user to PatternFlowIgmpv3GroupRecordMulticastAddress.
	// PatternFlowIgmpv3GroupRecordMulticastAddressCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowIgmpv3GroupRecordMulticastAddressCounter) PatternFlowIgmpv3GroupRecordMulticastAddress
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3GroupRecordMulticastAddress
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3GroupRecordMulticastAddress
var PatternFlowIgmpv3GroupRecordMulticastAddressChoice = struct {
	VALUE     PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum
	VALUES    PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum
	INCREMENT PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum
	DECREMENT PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) Choice() PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum {
	return PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) setChoice(value PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum) PatternFlowIgmpv3GroupRecordMulticastAddress {
	intValue, ok := otg.PatternFlowIgmpv3GroupRecordMulticastAddress_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3GroupRecordMulticastAddress_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3GroupRecordMulticastAddressChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3GroupRecordMulticastAddressChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3GroupRecordMulticastAddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3GroupRecordMulticastAddressCounter().msg()
	}

	if value == PatternFlowIgmpv3GroupRecordMulticastAddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3GroupRecordMulticastAddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordMulticastAddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIgmpv3GroupRecordMulticastAddress object
func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) SetValue(value string) PatternFlowIgmpv3GroupRecordMulticastAddress {
	obj.setChoice(PatternFlowIgmpv3GroupRecordMulticastAddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIgmpv3GroupRecordMulticastAddress object
func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) SetValues(value []string) PatternFlowIgmpv3GroupRecordMulticastAddress {
	obj.setChoice(PatternFlowIgmpv3GroupRecordMulticastAddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv3GroupRecordMulticastAddressCounter
func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) Increment() PatternFlowIgmpv3GroupRecordMulticastAddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordMulticastAddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3GroupRecordMulticastAddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3GroupRecordMulticastAddressCounter
func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3GroupRecordMulticastAddressCounter value in the PatternFlowIgmpv3GroupRecordMulticastAddress object
func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) SetIncrement(value PatternFlowIgmpv3GroupRecordMulticastAddressCounter) PatternFlowIgmpv3GroupRecordMulticastAddress {
	obj.setChoice(PatternFlowIgmpv3GroupRecordMulticastAddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3GroupRecordMulticastAddressCounter
func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) Decrement() PatternFlowIgmpv3GroupRecordMulticastAddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordMulticastAddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3GroupRecordMulticastAddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3GroupRecordMulticastAddressCounter
func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3GroupRecordMulticastAddressCounter value in the PatternFlowIgmpv3GroupRecordMulticastAddress object
func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) SetDecrement(value PatternFlowIgmpv3GroupRecordMulticastAddressCounter) PatternFlowIgmpv3GroupRecordMulticastAddress {
	obj.setChoice(PatternFlowIgmpv3GroupRecordMulticastAddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv3GroupRecordMulticastAddress.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv3GroupRecordMulticastAddress.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIgmpv3GroupRecordMulticastAddress) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3GroupRecordMulticastAddressChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordMulticastAddressChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordMulticastAddressChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordMulticastAddressChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordMulticastAddressChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3GroupRecordMulticastAddressChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3GroupRecordMulticastAddress")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3GroupRecordMulticastAddress_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3GroupRecordMulticastAddress_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
