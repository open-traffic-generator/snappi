package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress *****
type patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	marshaller      marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	unMarshaller    unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	incrementHolder PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	decrementHolder PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
}

func NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress() PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress {
	obj := patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress{obj: &otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) msg() *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) setMsg(msg *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress struct {
	obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
}

type marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress interface {
	// ToProto marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress to protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	ToProto() (*otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress, error)
	// ToPbText marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress struct {
	obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
}

type unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress from protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	FromProto(msg *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) (PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) Marshal() marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) Unmarshal() unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) ToProto() (*otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) FromProto(msg *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) (PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) Clone() (PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress()
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

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress is link-layer address of the sender (source option) or target (target option).
type PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress to protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	// setMsg unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress from protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	// validate validates PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	Choice() PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	setChoice(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	HasChoice() bool
	// Value returns string, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	SetValue(value string) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	HasValue() bool
	// Values returns []string, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	SetValues(value []string) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	// Increment returns PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress.
	// PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter is mac counter pattern
	Increment() PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress.
	// PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter is mac counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress.
	// PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter is mac counter pattern
	Decrement() PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress.
	// PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter is mac counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress
var PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice = struct {
	VALUE     PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum
	VALUES    PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) Choice() PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum {
	return PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) setChoice(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress {
	intValue, ok := otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) SetValue(value string) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress {
	obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) SetValues(value []string) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress {
	obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) Increment() PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) SetIncrement(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress {
	obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) Decrement() PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) SetDecrement(value PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressCounter) PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress {
	obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddressChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLinkLayerAddress_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
