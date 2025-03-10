package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathRsvpHopIpv4Ipv4Address *****
type patternFlowRSVPPathRsvpHopIpv4Ipv4Address struct {
	validation
	obj             *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	marshaller      marshalPatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	unMarshaller    unMarshalPatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	incrementHolder PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	decrementHolder PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
}

func NewPatternFlowRSVPPathRsvpHopIpv4Ipv4Address() PatternFlowRSVPPathRsvpHopIpv4Ipv4Address {
	obj := patternFlowRSVPPathRsvpHopIpv4Ipv4Address{obj: &otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) msg() *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address {
	return obj.obj
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) setMsg(msg *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address) PatternFlowRSVPPathRsvpHopIpv4Ipv4Address {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathRsvpHopIpv4Ipv4Address struct {
	obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address
}

type marshalPatternFlowRSVPPathRsvpHopIpv4Ipv4Address interface {
	// ToProto marshals PatternFlowRSVPPathRsvpHopIpv4Ipv4Address to protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	ToProto() (*otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address, error)
	// ToPbText marshals PatternFlowRSVPPathRsvpHopIpv4Ipv4Address to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathRsvpHopIpv4Ipv4Address to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathRsvpHopIpv4Ipv4Address to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathRsvpHopIpv4Ipv4Address to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathRsvpHopIpv4Ipv4Address struct {
	obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address
}

type unMarshalPatternFlowRSVPPathRsvpHopIpv4Ipv4Address interface {
	// FromProto unmarshals PatternFlowRSVPPathRsvpHopIpv4Ipv4Address from protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	FromProto(msg *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address) (PatternFlowRSVPPathRsvpHopIpv4Ipv4Address, error)
	// FromPbText unmarshals PatternFlowRSVPPathRsvpHopIpv4Ipv4Address from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathRsvpHopIpv4Ipv4Address from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathRsvpHopIpv4Ipv4Address from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) Marshal() marshalPatternFlowRSVPPathRsvpHopIpv4Ipv4Address {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathRsvpHopIpv4Ipv4Address{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) Unmarshal() unMarshalPatternFlowRSVPPathRsvpHopIpv4Ipv4Address {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathRsvpHopIpv4Ipv4Address{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4Ipv4Address) ToProto() (*otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4Ipv4Address) FromProto(msg *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address) (PatternFlowRSVPPathRsvpHopIpv4Ipv4Address, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4Ipv4Address) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4Ipv4Address) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4Ipv4Address) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4Ipv4Address) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4Ipv4Address) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4Ipv4Address) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4Ipv4Address) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) Clone() (PatternFlowRSVPPathRsvpHopIpv4Ipv4Address, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathRsvpHopIpv4Ipv4Address()
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

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathRsvpHopIpv4Ipv4Address is the IPv4 address of the interface through which the last RSVP-knowledgeable hop forwarded this message.
type PatternFlowRSVPPathRsvpHopIpv4Ipv4Address interface {
	Validation
	// msg marshals PatternFlowRSVPPathRsvpHopIpv4Ipv4Address to protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	// setMsg unmarshals PatternFlowRSVPPathRsvpHopIpv4Ipv4Address from protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address) PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	// validate validates PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathRsvpHopIpv4Ipv4Address, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum, set in PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	Choice() PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum
	// setChoice assigns PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum provided by user to PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	setChoice(value PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum) PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	HasChoice() bool
	// Value returns string, set in PatternFlowRSVPPathRsvpHopIpv4Ipv4Address.
	Value() string
	// SetValue assigns string provided by user to PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	SetValue(value string) PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	// HasValue checks if Value has been set in PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	HasValue() bool
	// Values returns []string, set in PatternFlowRSVPPathRsvpHopIpv4Ipv4Address.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	SetValues(value []string) PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	// Increment returns PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter, set in PatternFlowRSVPPathRsvpHopIpv4Ipv4Address.
	// PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter is ipv4 counter pattern
	Increment() PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	// SetIncrement assigns PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter provided by user to PatternFlowRSVPPathRsvpHopIpv4Ipv4Address.
	// PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter, set in PatternFlowRSVPPathRsvpHopIpv4Ipv4Address.
	// PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter is ipv4 counter pattern
	Decrement() PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	// SetDecrement assigns PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter provided by user to PatternFlowRSVPPathRsvpHopIpv4Ipv4Address.
	// PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
var PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice = struct {
	VALUE     PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum
	VALUES    PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum
	INCREMENT PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum
	DECREMENT PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) Choice() PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum {
	return PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) setChoice(value PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum) PatternFlowRSVPPathRsvpHopIpv4Ipv4Address {
	intValue, ok := otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter().msg()
	}

	if value == PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowRSVPPathRsvpHopIpv4Ipv4Address object
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) SetValue(value string) PatternFlowRSVPPathRsvpHopIpv4Ipv4Address {
	obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowRSVPPathRsvpHopIpv4Ipv4Address object
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) SetValues(value []string) PatternFlowRSVPPathRsvpHopIpv4Ipv4Address {
	obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) Increment() PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter value in the PatternFlowRSVPPathRsvpHopIpv4Ipv4Address object
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) SetIncrement(value PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) PatternFlowRSVPPathRsvpHopIpv4Ipv4Address {
	obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) Decrement() PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter value in the PatternFlowRSVPPathRsvpHopIpv4Ipv4Address object
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) SetDecrement(value PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) PatternFlowRSVPPathRsvpHopIpv4Ipv4Address {
	obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathRsvpHopIpv4Ipv4Address.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathRsvpHopIpv4Ipv4Address.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4Address) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathRsvpHopIpv4Ipv4Address")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4Address_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
