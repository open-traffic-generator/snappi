package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowSnmpv2CVariableBindingValue *****
type flowSnmpv2CVariableBindingValue struct {
	validation
	obj                        *otg.FlowSnmpv2CVariableBindingValue
	marshaller                 marshalFlowSnmpv2CVariableBindingValue
	unMarshaller               unMarshalFlowSnmpv2CVariableBindingValue
	integerValueHolder         PatternFlowSnmpv2CVariableBindingValueIntegerValue
	stringValueHolder          FlowSnmpv2CVariableBindingStringValue
	ipAddressValueHolder       PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	counterValueHolder         PatternFlowSnmpv2CVariableBindingValueCounterValue
	timeticksValueHolder       PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	bigCounterValueHolder      PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	unsignedIntegerValueHolder PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
}

func NewFlowSnmpv2CVariableBindingValue() FlowSnmpv2CVariableBindingValue {
	obj := flowSnmpv2CVariableBindingValue{obj: &otg.FlowSnmpv2CVariableBindingValue{}}
	obj.setDefault()
	return &obj
}

func (obj *flowSnmpv2CVariableBindingValue) msg() *otg.FlowSnmpv2CVariableBindingValue {
	return obj.obj
}

func (obj *flowSnmpv2CVariableBindingValue) setMsg(msg *otg.FlowSnmpv2CVariableBindingValue) FlowSnmpv2CVariableBindingValue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowSnmpv2CVariableBindingValue struct {
	obj *flowSnmpv2CVariableBindingValue
}

type marshalFlowSnmpv2CVariableBindingValue interface {
	// ToProto marshals FlowSnmpv2CVariableBindingValue to protobuf object *otg.FlowSnmpv2CVariableBindingValue
	ToProto() (*otg.FlowSnmpv2CVariableBindingValue, error)
	// ToPbText marshals FlowSnmpv2CVariableBindingValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowSnmpv2CVariableBindingValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowSnmpv2CVariableBindingValue to JSON text
	ToJson() (string, error)
}

type unMarshalflowSnmpv2CVariableBindingValue struct {
	obj *flowSnmpv2CVariableBindingValue
}

type unMarshalFlowSnmpv2CVariableBindingValue interface {
	// FromProto unmarshals FlowSnmpv2CVariableBindingValue from protobuf object *otg.FlowSnmpv2CVariableBindingValue
	FromProto(msg *otg.FlowSnmpv2CVariableBindingValue) (FlowSnmpv2CVariableBindingValue, error)
	// FromPbText unmarshals FlowSnmpv2CVariableBindingValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowSnmpv2CVariableBindingValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowSnmpv2CVariableBindingValue from JSON text
	FromJson(value string) error
}

func (obj *flowSnmpv2CVariableBindingValue) Marshal() marshalFlowSnmpv2CVariableBindingValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowSnmpv2CVariableBindingValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowSnmpv2CVariableBindingValue) Unmarshal() unMarshalFlowSnmpv2CVariableBindingValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowSnmpv2CVariableBindingValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowSnmpv2CVariableBindingValue) ToProto() (*otg.FlowSnmpv2CVariableBindingValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowSnmpv2CVariableBindingValue) FromProto(msg *otg.FlowSnmpv2CVariableBindingValue) (FlowSnmpv2CVariableBindingValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowSnmpv2CVariableBindingValue) ToPbText() (string, error) {
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

func (m *unMarshalflowSnmpv2CVariableBindingValue) FromPbText(value string) error {
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

func (m *marshalflowSnmpv2CVariableBindingValue) ToYaml() (string, error) {
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

func (m *unMarshalflowSnmpv2CVariableBindingValue) FromYaml(value string) error {
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

func (m *marshalflowSnmpv2CVariableBindingValue) ToJson() (string, error) {
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

func (m *unMarshalflowSnmpv2CVariableBindingValue) FromJson(value string) error {
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

func (obj *flowSnmpv2CVariableBindingValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowSnmpv2CVariableBindingValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowSnmpv2CVariableBindingValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowSnmpv2CVariableBindingValue) Clone() (FlowSnmpv2CVariableBindingValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowSnmpv2CVariableBindingValue()
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

func (obj *flowSnmpv2CVariableBindingValue) setNil() {
	obj.integerValueHolder = nil
	obj.stringValueHolder = nil
	obj.ipAddressValueHolder = nil
	obj.counterValueHolder = nil
	obj.timeticksValueHolder = nil
	obj.bigCounterValueHolder = nil
	obj.unsignedIntegerValueHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowSnmpv2CVariableBindingValue is the value for the object_identifier as per RFC2578.
type FlowSnmpv2CVariableBindingValue interface {
	Validation
	// msg marshals FlowSnmpv2CVariableBindingValue to protobuf object *otg.FlowSnmpv2CVariableBindingValue
	// and doesn't set defaults
	msg() *otg.FlowSnmpv2CVariableBindingValue
	// setMsg unmarshals FlowSnmpv2CVariableBindingValue from protobuf object *otg.FlowSnmpv2CVariableBindingValue
	// and doesn't set defaults
	setMsg(*otg.FlowSnmpv2CVariableBindingValue) FlowSnmpv2CVariableBindingValue
	// provides marshal interface
	Marshal() marshalFlowSnmpv2CVariableBindingValue
	// provides unmarshal interface
	Unmarshal() unMarshalFlowSnmpv2CVariableBindingValue
	// validate validates FlowSnmpv2CVariableBindingValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowSnmpv2CVariableBindingValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowSnmpv2CVariableBindingValueChoiceEnum, set in FlowSnmpv2CVariableBindingValue
	Choice() FlowSnmpv2CVariableBindingValueChoiceEnum
	// setChoice assigns FlowSnmpv2CVariableBindingValueChoiceEnum provided by user to FlowSnmpv2CVariableBindingValue
	setChoice(value FlowSnmpv2CVariableBindingValueChoiceEnum) FlowSnmpv2CVariableBindingValue
	// HasChoice checks if Choice has been set in FlowSnmpv2CVariableBindingValue
	HasChoice() bool
	// getter for NoValue to set choice.
	NoValue()
	// IntegerValue returns PatternFlowSnmpv2CVariableBindingValueIntegerValue, set in FlowSnmpv2CVariableBindingValue.
	IntegerValue() PatternFlowSnmpv2CVariableBindingValueIntegerValue
	// SetIntegerValue assigns PatternFlowSnmpv2CVariableBindingValueIntegerValue provided by user to FlowSnmpv2CVariableBindingValue.
	SetIntegerValue(value PatternFlowSnmpv2CVariableBindingValueIntegerValue) FlowSnmpv2CVariableBindingValue
	// HasIntegerValue checks if IntegerValue has been set in FlowSnmpv2CVariableBindingValue
	HasIntegerValue() bool
	// StringValue returns FlowSnmpv2CVariableBindingStringValue, set in FlowSnmpv2CVariableBindingValue.
	StringValue() FlowSnmpv2CVariableBindingStringValue
	// SetStringValue assigns FlowSnmpv2CVariableBindingStringValue provided by user to FlowSnmpv2CVariableBindingValue.
	SetStringValue(value FlowSnmpv2CVariableBindingStringValue) FlowSnmpv2CVariableBindingValue
	// HasStringValue checks if StringValue has been set in FlowSnmpv2CVariableBindingValue
	HasStringValue() bool
	// ObjectIdentifierValue returns string, set in FlowSnmpv2CVariableBindingValue.
	ObjectIdentifierValue() string
	// SetObjectIdentifierValue assigns string provided by user to FlowSnmpv2CVariableBindingValue
	SetObjectIdentifierValue(value string) FlowSnmpv2CVariableBindingValue
	// HasObjectIdentifierValue checks if ObjectIdentifierValue has been set in FlowSnmpv2CVariableBindingValue
	HasObjectIdentifierValue() bool
	// IpAddressValue returns PatternFlowSnmpv2CVariableBindingValueIpAddressValue, set in FlowSnmpv2CVariableBindingValue.
	IpAddressValue() PatternFlowSnmpv2CVariableBindingValueIpAddressValue
	// SetIpAddressValue assigns PatternFlowSnmpv2CVariableBindingValueIpAddressValue provided by user to FlowSnmpv2CVariableBindingValue.
	SetIpAddressValue(value PatternFlowSnmpv2CVariableBindingValueIpAddressValue) FlowSnmpv2CVariableBindingValue
	// HasIpAddressValue checks if IpAddressValue has been set in FlowSnmpv2CVariableBindingValue
	HasIpAddressValue() bool
	// CounterValue returns PatternFlowSnmpv2CVariableBindingValueCounterValue, set in FlowSnmpv2CVariableBindingValue.
	CounterValue() PatternFlowSnmpv2CVariableBindingValueCounterValue
	// SetCounterValue assigns PatternFlowSnmpv2CVariableBindingValueCounterValue provided by user to FlowSnmpv2CVariableBindingValue.
	SetCounterValue(value PatternFlowSnmpv2CVariableBindingValueCounterValue) FlowSnmpv2CVariableBindingValue
	// HasCounterValue checks if CounterValue has been set in FlowSnmpv2CVariableBindingValue
	HasCounterValue() bool
	// TimeticksValue returns PatternFlowSnmpv2CVariableBindingValueTimeticksValue, set in FlowSnmpv2CVariableBindingValue.
	TimeticksValue() PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	// SetTimeticksValue assigns PatternFlowSnmpv2CVariableBindingValueTimeticksValue provided by user to FlowSnmpv2CVariableBindingValue.
	SetTimeticksValue(value PatternFlowSnmpv2CVariableBindingValueTimeticksValue) FlowSnmpv2CVariableBindingValue
	// HasTimeticksValue checks if TimeticksValue has been set in FlowSnmpv2CVariableBindingValue
	HasTimeticksValue() bool
	// ArbitraryValue returns string, set in FlowSnmpv2CVariableBindingValue.
	ArbitraryValue() string
	// SetArbitraryValue assigns string provided by user to FlowSnmpv2CVariableBindingValue
	SetArbitraryValue(value string) FlowSnmpv2CVariableBindingValue
	// HasArbitraryValue checks if ArbitraryValue has been set in FlowSnmpv2CVariableBindingValue
	HasArbitraryValue() bool
	// BigCounterValue returns PatternFlowSnmpv2CVariableBindingValueBigCounterValue, set in FlowSnmpv2CVariableBindingValue.
	BigCounterValue() PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	// SetBigCounterValue assigns PatternFlowSnmpv2CVariableBindingValueBigCounterValue provided by user to FlowSnmpv2CVariableBindingValue.
	SetBigCounterValue(value PatternFlowSnmpv2CVariableBindingValueBigCounterValue) FlowSnmpv2CVariableBindingValue
	// HasBigCounterValue checks if BigCounterValue has been set in FlowSnmpv2CVariableBindingValue
	HasBigCounterValue() bool
	// UnsignedIntegerValue returns PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue, set in FlowSnmpv2CVariableBindingValue.
	UnsignedIntegerValue() PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	// SetUnsignedIntegerValue assigns PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue provided by user to FlowSnmpv2CVariableBindingValue.
	SetUnsignedIntegerValue(value PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) FlowSnmpv2CVariableBindingValue
	// HasUnsignedIntegerValue checks if UnsignedIntegerValue has been set in FlowSnmpv2CVariableBindingValue
	HasUnsignedIntegerValue() bool
	setNil()
}

type FlowSnmpv2CVariableBindingValueChoiceEnum string

// Enum of Choice on FlowSnmpv2CVariableBindingValue
var FlowSnmpv2CVariableBindingValueChoice = struct {
	NO_VALUE                FlowSnmpv2CVariableBindingValueChoiceEnum
	INTEGER_VALUE           FlowSnmpv2CVariableBindingValueChoiceEnum
	STRING_VALUE            FlowSnmpv2CVariableBindingValueChoiceEnum
	OBJECT_IDENTIFIER_VALUE FlowSnmpv2CVariableBindingValueChoiceEnum
	IP_ADDRESS_VALUE        FlowSnmpv2CVariableBindingValueChoiceEnum
	COUNTER_VALUE           FlowSnmpv2CVariableBindingValueChoiceEnum
	TIMETICKS_VALUE         FlowSnmpv2CVariableBindingValueChoiceEnum
	ARBITRARY_VALUE         FlowSnmpv2CVariableBindingValueChoiceEnum
	BIG_COUNTER_VALUE       FlowSnmpv2CVariableBindingValueChoiceEnum
	UNSIGNED_INTEGER_VALUE  FlowSnmpv2CVariableBindingValueChoiceEnum
}{
	NO_VALUE:                FlowSnmpv2CVariableBindingValueChoiceEnum("no_value"),
	INTEGER_VALUE:           FlowSnmpv2CVariableBindingValueChoiceEnum("integer_value"),
	STRING_VALUE:            FlowSnmpv2CVariableBindingValueChoiceEnum("string_value"),
	OBJECT_IDENTIFIER_VALUE: FlowSnmpv2CVariableBindingValueChoiceEnum("object_identifier_value"),
	IP_ADDRESS_VALUE:        FlowSnmpv2CVariableBindingValueChoiceEnum("ip_address_value"),
	COUNTER_VALUE:           FlowSnmpv2CVariableBindingValueChoiceEnum("counter_value"),
	TIMETICKS_VALUE:         FlowSnmpv2CVariableBindingValueChoiceEnum("timeticks_value"),
	ARBITRARY_VALUE:         FlowSnmpv2CVariableBindingValueChoiceEnum("arbitrary_value"),
	BIG_COUNTER_VALUE:       FlowSnmpv2CVariableBindingValueChoiceEnum("big_counter_value"),
	UNSIGNED_INTEGER_VALUE:  FlowSnmpv2CVariableBindingValueChoiceEnum("unsigned_integer_value"),
}

func (obj *flowSnmpv2CVariableBindingValue) Choice() FlowSnmpv2CVariableBindingValueChoiceEnum {
	return FlowSnmpv2CVariableBindingValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for NoValue to set choice
func (obj *flowSnmpv2CVariableBindingValue) NoValue() {
	obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.NO_VALUE)
}

// description is TBD
// Choice returns a string
func (obj *flowSnmpv2CVariableBindingValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowSnmpv2CVariableBindingValue) setChoice(value FlowSnmpv2CVariableBindingValueChoiceEnum) FlowSnmpv2CVariableBindingValue {
	intValue, ok := otg.FlowSnmpv2CVariableBindingValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowSnmpv2CVariableBindingValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowSnmpv2CVariableBindingValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.UnsignedIntegerValue = nil
	obj.unsignedIntegerValueHolder = nil
	obj.obj.BigCounterValue = nil
	obj.bigCounterValueHolder = nil
	obj.obj.ArbitraryValue = nil
	obj.obj.TimeticksValue = nil
	obj.timeticksValueHolder = nil
	obj.obj.CounterValue = nil
	obj.counterValueHolder = nil
	obj.obj.IpAddressValue = nil
	obj.ipAddressValueHolder = nil
	obj.obj.ObjectIdentifierValue = nil
	obj.obj.StringValue = nil
	obj.stringValueHolder = nil
	obj.obj.IntegerValue = nil
	obj.integerValueHolder = nil

	if value == FlowSnmpv2CVariableBindingValueChoice.INTEGER_VALUE {
		obj.obj.IntegerValue = NewPatternFlowSnmpv2CVariableBindingValueIntegerValue().msg()
	}

	if value == FlowSnmpv2CVariableBindingValueChoice.STRING_VALUE {
		obj.obj.StringValue = NewFlowSnmpv2CVariableBindingStringValue().msg()
	}

	if value == FlowSnmpv2CVariableBindingValueChoice.OBJECT_IDENTIFIER_VALUE {
		defaultValue := "0.1"
		obj.obj.ObjectIdentifierValue = &defaultValue
	}

	if value == FlowSnmpv2CVariableBindingValueChoice.IP_ADDRESS_VALUE {
		obj.obj.IpAddressValue = NewPatternFlowSnmpv2CVariableBindingValueIpAddressValue().msg()
	}

	if value == FlowSnmpv2CVariableBindingValueChoice.COUNTER_VALUE {
		obj.obj.CounterValue = NewPatternFlowSnmpv2CVariableBindingValueCounterValue().msg()
	}

	if value == FlowSnmpv2CVariableBindingValueChoice.TIMETICKS_VALUE {
		obj.obj.TimeticksValue = NewPatternFlowSnmpv2CVariableBindingValueTimeticksValue().msg()
	}

	if value == FlowSnmpv2CVariableBindingValueChoice.ARBITRARY_VALUE {
		defaultValue := "00"
		obj.obj.ArbitraryValue = &defaultValue
	}

	if value == FlowSnmpv2CVariableBindingValueChoice.BIG_COUNTER_VALUE {
		obj.obj.BigCounterValue = NewPatternFlowSnmpv2CVariableBindingValueBigCounterValue().msg()
	}

	if value == FlowSnmpv2CVariableBindingValueChoice.UNSIGNED_INTEGER_VALUE {
		obj.obj.UnsignedIntegerValue = NewPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue().msg()
	}

	return obj
}

// description is TBD
// IntegerValue returns a PatternFlowSnmpv2CVariableBindingValueIntegerValue
func (obj *flowSnmpv2CVariableBindingValue) IntegerValue() PatternFlowSnmpv2CVariableBindingValueIntegerValue {
	if obj.obj.IntegerValue == nil {
		obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.INTEGER_VALUE)
	}
	if obj.integerValueHolder == nil {
		obj.integerValueHolder = &patternFlowSnmpv2CVariableBindingValueIntegerValue{obj: obj.obj.IntegerValue}
	}
	return obj.integerValueHolder
}

// description is TBD
// IntegerValue returns a PatternFlowSnmpv2CVariableBindingValueIntegerValue
func (obj *flowSnmpv2CVariableBindingValue) HasIntegerValue() bool {
	return obj.obj.IntegerValue != nil
}

// description is TBD
// SetIntegerValue sets the PatternFlowSnmpv2CVariableBindingValueIntegerValue value in the FlowSnmpv2CVariableBindingValue object
func (obj *flowSnmpv2CVariableBindingValue) SetIntegerValue(value PatternFlowSnmpv2CVariableBindingValueIntegerValue) FlowSnmpv2CVariableBindingValue {
	obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.INTEGER_VALUE)
	obj.integerValueHolder = nil
	obj.obj.IntegerValue = value.msg()

	return obj
}

// description is TBD
// StringValue returns a FlowSnmpv2CVariableBindingStringValue
func (obj *flowSnmpv2CVariableBindingValue) StringValue() FlowSnmpv2CVariableBindingStringValue {
	if obj.obj.StringValue == nil {
		obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.STRING_VALUE)
	}
	if obj.stringValueHolder == nil {
		obj.stringValueHolder = &flowSnmpv2CVariableBindingStringValue{obj: obj.obj.StringValue}
	}
	return obj.stringValueHolder
}

// description is TBD
// StringValue returns a FlowSnmpv2CVariableBindingStringValue
func (obj *flowSnmpv2CVariableBindingValue) HasStringValue() bool {
	return obj.obj.StringValue != nil
}

// description is TBD
// SetStringValue sets the FlowSnmpv2CVariableBindingStringValue value in the FlowSnmpv2CVariableBindingValue object
func (obj *flowSnmpv2CVariableBindingValue) SetStringValue(value FlowSnmpv2CVariableBindingStringValue) FlowSnmpv2CVariableBindingValue {
	obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.STRING_VALUE)
	obj.stringValueHolder = nil
	obj.obj.StringValue = value.msg()

	return obj
}

// The Object Identifier points to a particular parameter in the SNMP agent.
// - Encoding of this field follows RFC2578(section 3.5) and ASN.1 X.690(section 8.1.3.6) specification.
// Refer: http://www.itu.int/ITU-T/asn1/
// - According to BER, the first two numbers of any OID (x.y) are encoded as one value using the formula (40*x)+y.
// Example, the first two numbers of an SNMP OID 1.3... are encoded as 43 or 0x2B, because (40*1)+3 = 43.
// - After the first two numbers are encoded, the subsequent numbers in the OID are each encoded as a byte.
// - However, a special rule is required for large numbers because one byte can only represent a number from 0-127.
// - The rule for large numbers states that only the lower 7 bits in the byte are used for holding the value (0-127).
// - The highest order bit(8th) is used as a flag to indicate that this number spans more than one byte. Therefore, any number over 127 must be encoded using more than one byte.
// - Example, the number 2680 in the OID '1.3.6.1.4.1.2680.1.2.7.3.2.0' cannot be encoded using a single byte.
// According to this rule, the number 2680 must be encoded as 0x94 0x78.
// Since the most significant bit is set in the first byte (0x94), it indicates that number spans to the next byte.
// Since the most significant bit is not set in the next byte (0x78), it indicates that the number ends at the second byte.
// The value is derived by appending 7 bits from each of the concatenated bytes i.e (0x14 *128^1) + (0x78 * 128^0) = 2680.
// ObjectIdentifierValue returns a string
func (obj *flowSnmpv2CVariableBindingValue) ObjectIdentifierValue() string {

	if obj.obj.ObjectIdentifierValue == nil {
		obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.OBJECT_IDENTIFIER_VALUE)
	}

	return *obj.obj.ObjectIdentifierValue

}

// The Object Identifier points to a particular parameter in the SNMP agent.
// - Encoding of this field follows RFC2578(section 3.5) and ASN.1 X.690(section 8.1.3.6) specification.
// Refer: http://www.itu.int/ITU-T/asn1/
// - According to BER, the first two numbers of any OID (x.y) are encoded as one value using the formula (40*x)+y.
// Example, the first two numbers of an SNMP OID 1.3... are encoded as 43 or 0x2B, because (40*1)+3 = 43.
// - After the first two numbers are encoded, the subsequent numbers in the OID are each encoded as a byte.
// - However, a special rule is required for large numbers because one byte can only represent a number from 0-127.
// - The rule for large numbers states that only the lower 7 bits in the byte are used for holding the value (0-127).
// - The highest order bit(8th) is used as a flag to indicate that this number spans more than one byte. Therefore, any number over 127 must be encoded using more than one byte.
// - Example, the number 2680 in the OID '1.3.6.1.4.1.2680.1.2.7.3.2.0' cannot be encoded using a single byte.
// According to this rule, the number 2680 must be encoded as 0x94 0x78.
// Since the most significant bit is set in the first byte (0x94), it indicates that number spans to the next byte.
// Since the most significant bit is not set in the next byte (0x78), it indicates that the number ends at the second byte.
// The value is derived by appending 7 bits from each of the concatenated bytes i.e (0x14 *128^1) + (0x78 * 128^0) = 2680.
// ObjectIdentifierValue returns a string
func (obj *flowSnmpv2CVariableBindingValue) HasObjectIdentifierValue() bool {
	return obj.obj.ObjectIdentifierValue != nil
}

// The Object Identifier points to a particular parameter in the SNMP agent.
// - Encoding of this field follows RFC2578(section 3.5) and ASN.1 X.690(section 8.1.3.6) specification.
// Refer: http://www.itu.int/ITU-T/asn1/
// - According to BER, the first two numbers of any OID (x.y) are encoded as one value using the formula (40*x)+y.
// Example, the first two numbers of an SNMP OID 1.3... are encoded as 43 or 0x2B, because (40*1)+3 = 43.
// - After the first two numbers are encoded, the subsequent numbers in the OID are each encoded as a byte.
// - However, a special rule is required for large numbers because one byte can only represent a number from 0-127.
// - The rule for large numbers states that only the lower 7 bits in the byte are used for holding the value (0-127).
// - The highest order bit(8th) is used as a flag to indicate that this number spans more than one byte. Therefore, any number over 127 must be encoded using more than one byte.
// - Example, the number 2680 in the OID '1.3.6.1.4.1.2680.1.2.7.3.2.0' cannot be encoded using a single byte.
// According to this rule, the number 2680 must be encoded as 0x94 0x78.
// Since the most significant bit is set in the first byte (0x94), it indicates that number spans to the next byte.
// Since the most significant bit is not set in the next byte (0x78), it indicates that the number ends at the second byte.
// The value is derived by appending 7 bits from each of the concatenated bytes i.e (0x14 *128^1) + (0x78 * 128^0) = 2680.
// SetObjectIdentifierValue sets the string value in the FlowSnmpv2CVariableBindingValue object
func (obj *flowSnmpv2CVariableBindingValue) SetObjectIdentifierValue(value string) FlowSnmpv2CVariableBindingValue {
	obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.OBJECT_IDENTIFIER_VALUE)
	obj.obj.ObjectIdentifierValue = &value
	return obj
}

// description is TBD
// IpAddressValue returns a PatternFlowSnmpv2CVariableBindingValueIpAddressValue
func (obj *flowSnmpv2CVariableBindingValue) IpAddressValue() PatternFlowSnmpv2CVariableBindingValueIpAddressValue {
	if obj.obj.IpAddressValue == nil {
		obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.IP_ADDRESS_VALUE)
	}
	if obj.ipAddressValueHolder == nil {
		obj.ipAddressValueHolder = &patternFlowSnmpv2CVariableBindingValueIpAddressValue{obj: obj.obj.IpAddressValue}
	}
	return obj.ipAddressValueHolder
}

// description is TBD
// IpAddressValue returns a PatternFlowSnmpv2CVariableBindingValueIpAddressValue
func (obj *flowSnmpv2CVariableBindingValue) HasIpAddressValue() bool {
	return obj.obj.IpAddressValue != nil
}

// description is TBD
// SetIpAddressValue sets the PatternFlowSnmpv2CVariableBindingValueIpAddressValue value in the FlowSnmpv2CVariableBindingValue object
func (obj *flowSnmpv2CVariableBindingValue) SetIpAddressValue(value PatternFlowSnmpv2CVariableBindingValueIpAddressValue) FlowSnmpv2CVariableBindingValue {
	obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.IP_ADDRESS_VALUE)
	obj.ipAddressValueHolder = nil
	obj.obj.IpAddressValue = value.msg()

	return obj
}

// description is TBD
// CounterValue returns a PatternFlowSnmpv2CVariableBindingValueCounterValue
func (obj *flowSnmpv2CVariableBindingValue) CounterValue() PatternFlowSnmpv2CVariableBindingValueCounterValue {
	if obj.obj.CounterValue == nil {
		obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.COUNTER_VALUE)
	}
	if obj.counterValueHolder == nil {
		obj.counterValueHolder = &patternFlowSnmpv2CVariableBindingValueCounterValue{obj: obj.obj.CounterValue}
	}
	return obj.counterValueHolder
}

// description is TBD
// CounterValue returns a PatternFlowSnmpv2CVariableBindingValueCounterValue
func (obj *flowSnmpv2CVariableBindingValue) HasCounterValue() bool {
	return obj.obj.CounterValue != nil
}

// description is TBD
// SetCounterValue sets the PatternFlowSnmpv2CVariableBindingValueCounterValue value in the FlowSnmpv2CVariableBindingValue object
func (obj *flowSnmpv2CVariableBindingValue) SetCounterValue(value PatternFlowSnmpv2CVariableBindingValueCounterValue) FlowSnmpv2CVariableBindingValue {
	obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.COUNTER_VALUE)
	obj.counterValueHolder = nil
	obj.obj.CounterValue = value.msg()

	return obj
}

// description is TBD
// TimeticksValue returns a PatternFlowSnmpv2CVariableBindingValueTimeticksValue
func (obj *flowSnmpv2CVariableBindingValue) TimeticksValue() PatternFlowSnmpv2CVariableBindingValueTimeticksValue {
	if obj.obj.TimeticksValue == nil {
		obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.TIMETICKS_VALUE)
	}
	if obj.timeticksValueHolder == nil {
		obj.timeticksValueHolder = &patternFlowSnmpv2CVariableBindingValueTimeticksValue{obj: obj.obj.TimeticksValue}
	}
	return obj.timeticksValueHolder
}

// description is TBD
// TimeticksValue returns a PatternFlowSnmpv2CVariableBindingValueTimeticksValue
func (obj *flowSnmpv2CVariableBindingValue) HasTimeticksValue() bool {
	return obj.obj.TimeticksValue != nil
}

// description is TBD
// SetTimeticksValue sets the PatternFlowSnmpv2CVariableBindingValueTimeticksValue value in the FlowSnmpv2CVariableBindingValue object
func (obj *flowSnmpv2CVariableBindingValue) SetTimeticksValue(value PatternFlowSnmpv2CVariableBindingValueTimeticksValue) FlowSnmpv2CVariableBindingValue {
	obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.TIMETICKS_VALUE)
	obj.timeticksValueHolder = nil
	obj.obj.TimeticksValue = value.msg()

	return obj
}

// It contains the hex bytes of the value to be sent.  As of now it is restricted to 10000 bytes.
// ArbitraryValue returns a string
func (obj *flowSnmpv2CVariableBindingValue) ArbitraryValue() string {

	if obj.obj.ArbitraryValue == nil {
		obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.ARBITRARY_VALUE)
	}

	return *obj.obj.ArbitraryValue

}

// It contains the hex bytes of the value to be sent.  As of now it is restricted to 10000 bytes.
// ArbitraryValue returns a string
func (obj *flowSnmpv2CVariableBindingValue) HasArbitraryValue() bool {
	return obj.obj.ArbitraryValue != nil
}

// It contains the hex bytes of the value to be sent.  As of now it is restricted to 10000 bytes.
// SetArbitraryValue sets the string value in the FlowSnmpv2CVariableBindingValue object
func (obj *flowSnmpv2CVariableBindingValue) SetArbitraryValue(value string) FlowSnmpv2CVariableBindingValue {
	obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.ARBITRARY_VALUE)
	obj.obj.ArbitraryValue = &value
	return obj
}

// description is TBD
// BigCounterValue returns a PatternFlowSnmpv2CVariableBindingValueBigCounterValue
func (obj *flowSnmpv2CVariableBindingValue) BigCounterValue() PatternFlowSnmpv2CVariableBindingValueBigCounterValue {
	if obj.obj.BigCounterValue == nil {
		obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.BIG_COUNTER_VALUE)
	}
	if obj.bigCounterValueHolder == nil {
		obj.bigCounterValueHolder = &patternFlowSnmpv2CVariableBindingValueBigCounterValue{obj: obj.obj.BigCounterValue}
	}
	return obj.bigCounterValueHolder
}

// description is TBD
// BigCounterValue returns a PatternFlowSnmpv2CVariableBindingValueBigCounterValue
func (obj *flowSnmpv2CVariableBindingValue) HasBigCounterValue() bool {
	return obj.obj.BigCounterValue != nil
}

// description is TBD
// SetBigCounterValue sets the PatternFlowSnmpv2CVariableBindingValueBigCounterValue value in the FlowSnmpv2CVariableBindingValue object
func (obj *flowSnmpv2CVariableBindingValue) SetBigCounterValue(value PatternFlowSnmpv2CVariableBindingValueBigCounterValue) FlowSnmpv2CVariableBindingValue {
	obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.BIG_COUNTER_VALUE)
	obj.bigCounterValueHolder = nil
	obj.obj.BigCounterValue = value.msg()

	return obj
}

// description is TBD
// UnsignedIntegerValue returns a PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
func (obj *flowSnmpv2CVariableBindingValue) UnsignedIntegerValue() PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue {
	if obj.obj.UnsignedIntegerValue == nil {
		obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.UNSIGNED_INTEGER_VALUE)
	}
	if obj.unsignedIntegerValueHolder == nil {
		obj.unsignedIntegerValueHolder = &patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue{obj: obj.obj.UnsignedIntegerValue}
	}
	return obj.unsignedIntegerValueHolder
}

// description is TBD
// UnsignedIntegerValue returns a PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
func (obj *flowSnmpv2CVariableBindingValue) HasUnsignedIntegerValue() bool {
	return obj.obj.UnsignedIntegerValue != nil
}

// description is TBD
// SetUnsignedIntegerValue sets the PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue value in the FlowSnmpv2CVariableBindingValue object
func (obj *flowSnmpv2CVariableBindingValue) SetUnsignedIntegerValue(value PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) FlowSnmpv2CVariableBindingValue {
	obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.UNSIGNED_INTEGER_VALUE)
	obj.unsignedIntegerValueHolder = nil
	obj.obj.UnsignedIntegerValue = value.msg()

	return obj
}

func (obj *flowSnmpv2CVariableBindingValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.IntegerValue != nil {

		obj.IntegerValue().validateObj(vObj, set_default)
	}

	if obj.obj.StringValue != nil {

		obj.StringValue().validateObj(vObj, set_default)
	}

	if obj.obj.ObjectIdentifierValue != nil {

		err := obj.validateOid(obj.ObjectIdentifierValue())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowSnmpv2CVariableBindingValue.ObjectIdentifierValue"))
		}

	}

	if obj.obj.IpAddressValue != nil {

		obj.IpAddressValue().validateObj(vObj, set_default)
	}

	if obj.obj.CounterValue != nil {

		obj.CounterValue().validateObj(vObj, set_default)
	}

	if obj.obj.TimeticksValue != nil {

		obj.TimeticksValue().validateObj(vObj, set_default)
	}

	if obj.obj.ArbitraryValue != nil {

		if len(*obj.obj.ArbitraryValue) > 10000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of FlowSnmpv2CVariableBindingValue.ArbitraryValue <= 10000 but Got %d",
					len(*obj.obj.ArbitraryValue)))
		}

	}

	if obj.obj.BigCounterValue != nil {

		obj.BigCounterValue().validateObj(vObj, set_default)
	}

	if obj.obj.UnsignedIntegerValue != nil {

		obj.UnsignedIntegerValue().validateObj(vObj, set_default)
	}

}

func (obj *flowSnmpv2CVariableBindingValue) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowSnmpv2CVariableBindingValueChoice.NO_VALUE)

	}

}
