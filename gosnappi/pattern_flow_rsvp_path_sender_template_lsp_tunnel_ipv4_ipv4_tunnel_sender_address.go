package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress *****
type patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	marshaller      marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	incrementHolder PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	decrementHolder PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
}

func NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress {
	obj := patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress{obj: &otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) msg() *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) setMsg(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress struct {
	obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
}

type marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress interface {
	// ToProto marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress to protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	ToProto() (*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress struct {
	obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
}

type unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress from protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	FromProto(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) Marshal() marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) Unmarshal() unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) ToProto() (*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) FromProto(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) Clone() (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress()
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

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress is iPv4 address for a sender node.
type PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress to protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	// setMsg unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress from protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	// validate validates PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	Choice() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	setChoice(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	HasChoice() bool
	// Value returns string, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress.
	Value() string
	// SetValue assigns string provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	SetValue(value string) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	HasValue() bool
	// Values returns []string, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	SetValues(value []string) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	// Increment returns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter is ipv4 counter pattern
	Increment() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	// SetIncrement assigns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter is ipv4 counter pattern
	Decrement() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
	// SetDecrement assigns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress
var PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice = struct {
	VALUE     PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) Choice() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum {
	return PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) setChoice(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress {
	intValue, ok := otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter().msg()
	}

	if value == PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) SetValue(value string) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress {
	obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) SetValues(value []string) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress {
	obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) Increment() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) SetIncrement(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress {
	obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) Decrement() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) SetDecrement(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress {
	obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddressChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Ipv4TunnelSenderAddress_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
