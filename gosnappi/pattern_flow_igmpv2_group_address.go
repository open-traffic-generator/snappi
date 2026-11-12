package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv2GroupAddress *****
type patternFlowIgmpv2GroupAddress struct {
	validation
	obj             *otg.PatternFlowIgmpv2GroupAddress
	marshaller      marshalPatternFlowIgmpv2GroupAddress
	unMarshaller    unMarshalPatternFlowIgmpv2GroupAddress
	incrementHolder PatternFlowIgmpv2GroupAddressCounter
	decrementHolder PatternFlowIgmpv2GroupAddressCounter
}

func NewPatternFlowIgmpv2GroupAddress() PatternFlowIgmpv2GroupAddress {
	obj := patternFlowIgmpv2GroupAddress{obj: &otg.PatternFlowIgmpv2GroupAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv2GroupAddress) msg() *otg.PatternFlowIgmpv2GroupAddress {
	return obj.obj
}

func (obj *patternFlowIgmpv2GroupAddress) setMsg(msg *otg.PatternFlowIgmpv2GroupAddress) PatternFlowIgmpv2GroupAddress {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv2GroupAddress struct {
	obj *patternFlowIgmpv2GroupAddress
}

type marshalPatternFlowIgmpv2GroupAddress interface {
	// ToProto marshals PatternFlowIgmpv2GroupAddress to protobuf object *otg.PatternFlowIgmpv2GroupAddress
	ToProto() (*otg.PatternFlowIgmpv2GroupAddress, error)
	// ToPbText marshals PatternFlowIgmpv2GroupAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv2GroupAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv2GroupAddress to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv2GroupAddress struct {
	obj *patternFlowIgmpv2GroupAddress
}

type unMarshalPatternFlowIgmpv2GroupAddress interface {
	// FromProto unmarshals PatternFlowIgmpv2GroupAddress from protobuf object *otg.PatternFlowIgmpv2GroupAddress
	FromProto(msg *otg.PatternFlowIgmpv2GroupAddress) (PatternFlowIgmpv2GroupAddress, error)
	// FromPbText unmarshals PatternFlowIgmpv2GroupAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv2GroupAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv2GroupAddress from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv2GroupAddress) Marshal() marshalPatternFlowIgmpv2GroupAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv2GroupAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv2GroupAddress) Unmarshal() unMarshalPatternFlowIgmpv2GroupAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv2GroupAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv2GroupAddress) ToProto() (*otg.PatternFlowIgmpv2GroupAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv2GroupAddress) FromProto(msg *otg.PatternFlowIgmpv2GroupAddress) (PatternFlowIgmpv2GroupAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv2GroupAddress) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2GroupAddress) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv2GroupAddress) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2GroupAddress) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv2GroupAddress) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2GroupAddress) FromJson(value string) error {
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

func (obj *patternFlowIgmpv2GroupAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2GroupAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2GroupAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv2GroupAddress) Clone() (PatternFlowIgmpv2GroupAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv2GroupAddress()
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

func (obj *patternFlowIgmpv2GroupAddress) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv2GroupAddress is group address is zero in a general Membership Query. It is set to the group being reported/left in Reports and Leave Group messages (RFC 2236 Section 2.4).
type PatternFlowIgmpv2GroupAddress interface {
	Validation
	// msg marshals PatternFlowIgmpv2GroupAddress to protobuf object *otg.PatternFlowIgmpv2GroupAddress
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv2GroupAddress
	// setMsg unmarshals PatternFlowIgmpv2GroupAddress from protobuf object *otg.PatternFlowIgmpv2GroupAddress
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv2GroupAddress) PatternFlowIgmpv2GroupAddress
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv2GroupAddress
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv2GroupAddress
	// validate validates PatternFlowIgmpv2GroupAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv2GroupAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv2GroupAddressChoiceEnum, set in PatternFlowIgmpv2GroupAddress
	Choice() PatternFlowIgmpv2GroupAddressChoiceEnum
	// setChoice assigns PatternFlowIgmpv2GroupAddressChoiceEnum provided by user to PatternFlowIgmpv2GroupAddress
	setChoice(value PatternFlowIgmpv2GroupAddressChoiceEnum) PatternFlowIgmpv2GroupAddress
	// HasChoice checks if Choice has been set in PatternFlowIgmpv2GroupAddress
	HasChoice() bool
	// Value returns string, set in PatternFlowIgmpv2GroupAddress.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIgmpv2GroupAddress
	SetValue(value string) PatternFlowIgmpv2GroupAddress
	// HasValue checks if Value has been set in PatternFlowIgmpv2GroupAddress
	HasValue() bool
	// Values returns []string, set in PatternFlowIgmpv2GroupAddress.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIgmpv2GroupAddress
	SetValues(value []string) PatternFlowIgmpv2GroupAddress
	// Increment returns PatternFlowIgmpv2GroupAddressCounter, set in PatternFlowIgmpv2GroupAddress.
	// PatternFlowIgmpv2GroupAddressCounter is ipv4 counter pattern
	Increment() PatternFlowIgmpv2GroupAddressCounter
	// SetIncrement assigns PatternFlowIgmpv2GroupAddressCounter provided by user to PatternFlowIgmpv2GroupAddress.
	// PatternFlowIgmpv2GroupAddressCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowIgmpv2GroupAddressCounter) PatternFlowIgmpv2GroupAddress
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv2GroupAddress
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv2GroupAddressCounter, set in PatternFlowIgmpv2GroupAddress.
	// PatternFlowIgmpv2GroupAddressCounter is ipv4 counter pattern
	Decrement() PatternFlowIgmpv2GroupAddressCounter
	// SetDecrement assigns PatternFlowIgmpv2GroupAddressCounter provided by user to PatternFlowIgmpv2GroupAddress.
	// PatternFlowIgmpv2GroupAddressCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowIgmpv2GroupAddressCounter) PatternFlowIgmpv2GroupAddress
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv2GroupAddress
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv2GroupAddressChoiceEnum string

// Enum of Choice on PatternFlowIgmpv2GroupAddress
var PatternFlowIgmpv2GroupAddressChoice = struct {
	VALUE     PatternFlowIgmpv2GroupAddressChoiceEnum
	VALUES    PatternFlowIgmpv2GroupAddressChoiceEnum
	INCREMENT PatternFlowIgmpv2GroupAddressChoiceEnum
	DECREMENT PatternFlowIgmpv2GroupAddressChoiceEnum
}{
	VALUE:     PatternFlowIgmpv2GroupAddressChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv2GroupAddressChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv2GroupAddressChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv2GroupAddressChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv2GroupAddress) Choice() PatternFlowIgmpv2GroupAddressChoiceEnum {
	return PatternFlowIgmpv2GroupAddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv2GroupAddress) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv2GroupAddress) setChoice(value PatternFlowIgmpv2GroupAddressChoiceEnum) PatternFlowIgmpv2GroupAddress {
	intValue, ok := otg.PatternFlowIgmpv2GroupAddress_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv2GroupAddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv2GroupAddress_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv2GroupAddressChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv2GroupAddressChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv2GroupAddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv2GroupAddressCounter().msg()
	}

	if value == PatternFlowIgmpv2GroupAddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv2GroupAddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIgmpv2GroupAddress) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv2GroupAddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIgmpv2GroupAddress) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIgmpv2GroupAddress object
func (obj *patternFlowIgmpv2GroupAddress) SetValue(value string) PatternFlowIgmpv2GroupAddress {
	obj.setChoice(PatternFlowIgmpv2GroupAddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIgmpv2GroupAddress) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIgmpv2GroupAddress object
func (obj *patternFlowIgmpv2GroupAddress) SetValues(value []string) PatternFlowIgmpv2GroupAddress {
	obj.setChoice(PatternFlowIgmpv2GroupAddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv2GroupAddressCounter
func (obj *patternFlowIgmpv2GroupAddress) Increment() PatternFlowIgmpv2GroupAddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv2GroupAddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv2GroupAddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv2GroupAddressCounter
func (obj *patternFlowIgmpv2GroupAddress) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv2GroupAddressCounter value in the PatternFlowIgmpv2GroupAddress object
func (obj *patternFlowIgmpv2GroupAddress) SetIncrement(value PatternFlowIgmpv2GroupAddressCounter) PatternFlowIgmpv2GroupAddress {
	obj.setChoice(PatternFlowIgmpv2GroupAddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv2GroupAddressCounter
func (obj *patternFlowIgmpv2GroupAddress) Decrement() PatternFlowIgmpv2GroupAddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv2GroupAddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv2GroupAddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv2GroupAddressCounter
func (obj *patternFlowIgmpv2GroupAddress) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv2GroupAddressCounter value in the PatternFlowIgmpv2GroupAddress object
func (obj *patternFlowIgmpv2GroupAddress) SetDecrement(value PatternFlowIgmpv2GroupAddressCounter) PatternFlowIgmpv2GroupAddress {
	obj.setChoice(PatternFlowIgmpv2GroupAddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv2GroupAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv2GroupAddress.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv2GroupAddress.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIgmpv2GroupAddress) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv2GroupAddressChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv2GroupAddressChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv2GroupAddressChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv2GroupAddressChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv2GroupAddressChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv2GroupAddressChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv2GroupAddress")
			}
		} else {
			intVal := otg.PatternFlowIgmpv2GroupAddress_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv2GroupAddress_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
