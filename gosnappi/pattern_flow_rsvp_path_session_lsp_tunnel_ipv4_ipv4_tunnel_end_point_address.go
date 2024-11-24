package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress *****
type patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress struct {
	validation
	obj             *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	marshaller      marshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	unMarshaller    unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	incrementHolder PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	decrementHolder PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
}

func NewPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress() PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress {
	obj := patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress{obj: &otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) msg() *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress {
	return obj.obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) setMsg(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress struct {
	obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
}

type marshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress interface {
	// ToProto marshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress to protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	ToProto() (*otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress, error)
	// ToPbText marshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress struct {
	obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
}

type unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress interface {
	// FromProto unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress from protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	FromProto(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) (PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress, error)
	// FromPbText unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) Marshal() marshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) Unmarshal() unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) ToProto() (*otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) FromProto(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) (PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) Clone() (PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress()
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

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress is iPv4 address of the egress node for the tunnel.
type PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress interface {
	Validation
	// msg marshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress to protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	// setMsg unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress from protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	// validate validates PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum, set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	Choice() PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	setChoice(value PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	HasChoice() bool
	// Value returns string, set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress.
	Value() string
	// SetValue assigns string provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	SetValue(value string) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	// HasValue checks if Value has been set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	HasValue() bool
	// Values returns []string, set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	SetValues(value []string) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	// Increment returns PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter, set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress.
	// PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter is ipv4 counter pattern
	Increment() PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	// SetIncrement assigns PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress.
	// PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter, set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress.
	// PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter is ipv4 counter pattern
	Decrement() PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
	// SetDecrement assigns PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress.
	// PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress
var PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice = struct {
	VALUE     PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum
	VALUES    PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum
	INCREMENT PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum
	DECREMENT PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) Choice() PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum {
	return PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) setChoice(value PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress {
	intValue, ok := otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter().msg()
	}

	if value == PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) SetValue(value string) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress {
	obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) SetValues(value []string) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress {
	obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) Increment() PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter value in the PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) SetIncrement(value PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress {
	obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) Decrement() PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter value in the PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) SetDecrement(value PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressCounter) PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress {
	obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddressChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSessionLspTunnelIpv4Ipv4TunnelEndPointAddress_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
