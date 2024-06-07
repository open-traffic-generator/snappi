package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address *****
type patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address struct {
	validation
	obj             *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	marshaller      marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	unMarshaller    unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	incrementHolder PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	decrementHolder PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
}

func NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address() PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address {
	obj := patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address{obj: &otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) msg() *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address {
	return obj.obj
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) setMsg(msg *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address struct {
	obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
}

type marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address interface {
	// ToProto marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address to protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	ToProto() (*otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address, error)
	// ToPbText marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address struct {
	obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
}

type unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address interface {
	// FromProto unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address from protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	FromProto(msg *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) (PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address, error)
	// FromPbText unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) Marshal() marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) Unmarshal() unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) ToProto() (*otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) FromProto(msg *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) (PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) Clone() (PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address()
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

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address is a 32-bit unicast, host address.  Any network-reachable interface address is allowed here. Illegal addresses, such as certain loopback addresses, SHOULD NOT be used.
type PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address interface {
	Validation
	// msg marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address to protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	// setMsg unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address from protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	// validate validates PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	Choice() PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum
	// setChoice assigns PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	setChoice(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	HasChoice() bool
	// Value returns string, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address.
	Value() string
	// SetValue assigns string provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	SetValue(value string) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	// HasValue checks if Value has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	HasValue() bool
	// Values returns []string, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	SetValues(value []string) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	// Increment returns PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address.
	// PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter is ipv4 counter pattern
	Increment() PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	// SetIncrement assigns PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address.
	// PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address.
	// PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter is ipv4 counter pattern
	Decrement() PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	// SetDecrement assigns PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address.
	// PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
var PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice = struct {
	VALUE     PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum
	VALUES    PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum
	INCREMENT PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum
	DECREMENT PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) Choice() PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum {
	return PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) setChoice(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address {
	intValue, ok := otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter().msg()
	}

	if value == PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) SetValue(value string) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address {
	obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) SetValues(value []string) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address {
	obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) Increment() PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) SetIncrement(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address {
	obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) Decrement() PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) SetDecrement(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address {
	obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
