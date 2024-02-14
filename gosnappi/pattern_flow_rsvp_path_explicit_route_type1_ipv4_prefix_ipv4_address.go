package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address *****
type patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address struct {
	validation
	obj             *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	marshaller      marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	unMarshaller    unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	incrementHolder PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	decrementHolder PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
}

func NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address {
	obj := patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address{obj: &otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) msg() *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address {
	return obj.obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) setMsg(msg *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address struct {
	obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
}

type marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address interface {
	// ToProto marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address to protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	ToProto() (*otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address, error)
	// ToPbText marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address struct {
	obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
}

type unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address interface {
	// FromProto unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address from protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	FromProto(msg *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address, error)
	// FromPbText unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) Marshal() marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) Unmarshal() unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) ToProto() (*otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) FromProto(msg *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) Clone() (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address()
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

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address is this IPv4 address is treated as a prefix based on the prefix length value below.  Bits beyond the prefix are ignored on receipt and SHOULD be set to zero on transmission.
type PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address interface {
	Validation
	// msg marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address to protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	// setMsg unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address from protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	// validate validates PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	Choice() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum
	// setChoice assigns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	setChoice(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	HasChoice() bool
	// Value returns string, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address.
	Value() string
	// SetValue assigns string provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	SetValue(value string) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	// HasValue checks if Value has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	HasValue() bool
	// Values returns []string, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	SetValues(value []string) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	// Increment returns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address.
	// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter is ipv4 counter pattern
	Increment() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	// SetIncrement assigns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address.
	// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address.
	// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter is ipv4 counter pattern
	Decrement() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
	// SetDecrement assigns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address.
	// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
var PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoice = struct {
	VALUE     PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum
	VALUES    PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum
	INCREMENT PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum
	DECREMENT PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) Choice() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum {
	return PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) setChoice(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address {
	intValue, ok := otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter().msg()
	}

	if value == PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) SetValue(value string) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address {
	obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) SetValues(value []string) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address {
	obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) Increment() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) SetIncrement(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address {
	obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) Decrement() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) SetDecrement(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressCounter) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address {
	obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4AddressChoice.VALUE)

	}

}
