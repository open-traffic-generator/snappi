package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowSnmpv2CVariableBinding *****
type flowSnmpv2CVariableBinding struct {
	validation
	obj          *otg.FlowSnmpv2CVariableBinding
	marshaller   marshalFlowSnmpv2CVariableBinding
	unMarshaller unMarshalFlowSnmpv2CVariableBinding
	valueHolder  FlowSnmpv2CVariableBindingValue
}

func NewFlowSnmpv2CVariableBinding() FlowSnmpv2CVariableBinding {
	obj := flowSnmpv2CVariableBinding{obj: &otg.FlowSnmpv2CVariableBinding{}}
	obj.setDefault()
	return &obj
}

func (obj *flowSnmpv2CVariableBinding) msg() *otg.FlowSnmpv2CVariableBinding {
	return obj.obj
}

func (obj *flowSnmpv2CVariableBinding) setMsg(msg *otg.FlowSnmpv2CVariableBinding) FlowSnmpv2CVariableBinding {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowSnmpv2CVariableBinding struct {
	obj *flowSnmpv2CVariableBinding
}

type marshalFlowSnmpv2CVariableBinding interface {
	// ToProto marshals FlowSnmpv2CVariableBinding to protobuf object *otg.FlowSnmpv2CVariableBinding
	ToProto() (*otg.FlowSnmpv2CVariableBinding, error)
	// ToPbText marshals FlowSnmpv2CVariableBinding to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowSnmpv2CVariableBinding to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowSnmpv2CVariableBinding to JSON text
	ToJson() (string, error)
}

type unMarshalflowSnmpv2CVariableBinding struct {
	obj *flowSnmpv2CVariableBinding
}

type unMarshalFlowSnmpv2CVariableBinding interface {
	// FromProto unmarshals FlowSnmpv2CVariableBinding from protobuf object *otg.FlowSnmpv2CVariableBinding
	FromProto(msg *otg.FlowSnmpv2CVariableBinding) (FlowSnmpv2CVariableBinding, error)
	// FromPbText unmarshals FlowSnmpv2CVariableBinding from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowSnmpv2CVariableBinding from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowSnmpv2CVariableBinding from JSON text
	FromJson(value string) error
}

func (obj *flowSnmpv2CVariableBinding) Marshal() marshalFlowSnmpv2CVariableBinding {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowSnmpv2CVariableBinding{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowSnmpv2CVariableBinding) Unmarshal() unMarshalFlowSnmpv2CVariableBinding {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowSnmpv2CVariableBinding{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowSnmpv2CVariableBinding) ToProto() (*otg.FlowSnmpv2CVariableBinding, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowSnmpv2CVariableBinding) FromProto(msg *otg.FlowSnmpv2CVariableBinding) (FlowSnmpv2CVariableBinding, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowSnmpv2CVariableBinding) ToPbText() (string, error) {
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

func (m *unMarshalflowSnmpv2CVariableBinding) FromPbText(value string) error {
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

func (m *marshalflowSnmpv2CVariableBinding) ToYaml() (string, error) {
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

func (m *unMarshalflowSnmpv2CVariableBinding) FromYaml(value string) error {
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

func (m *marshalflowSnmpv2CVariableBinding) ToJson() (string, error) {
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

func (m *unMarshalflowSnmpv2CVariableBinding) FromJson(value string) error {
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

func (obj *flowSnmpv2CVariableBinding) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowSnmpv2CVariableBinding) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowSnmpv2CVariableBinding) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowSnmpv2CVariableBinding) Clone() (FlowSnmpv2CVariableBinding, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowSnmpv2CVariableBinding()
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

func (obj *flowSnmpv2CVariableBinding) setNil() {
	obj.valueHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowSnmpv2CVariableBinding is a Sequence of two fields, an object_identifier and the value for/from that object_identifier.
type FlowSnmpv2CVariableBinding interface {
	Validation
	// msg marshals FlowSnmpv2CVariableBinding to protobuf object *otg.FlowSnmpv2CVariableBinding
	// and doesn't set defaults
	msg() *otg.FlowSnmpv2CVariableBinding
	// setMsg unmarshals FlowSnmpv2CVariableBinding from protobuf object *otg.FlowSnmpv2CVariableBinding
	// and doesn't set defaults
	setMsg(*otg.FlowSnmpv2CVariableBinding) FlowSnmpv2CVariableBinding
	// provides marshal interface
	Marshal() marshalFlowSnmpv2CVariableBinding
	// provides unmarshal interface
	Unmarshal() unMarshalFlowSnmpv2CVariableBinding
	// validate validates FlowSnmpv2CVariableBinding
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowSnmpv2CVariableBinding, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ObjectIdentifier returns string, set in FlowSnmpv2CVariableBinding.
	ObjectIdentifier() string
	// SetObjectIdentifier assigns string provided by user to FlowSnmpv2CVariableBinding
	SetObjectIdentifier(value string) FlowSnmpv2CVariableBinding
	// HasObjectIdentifier checks if ObjectIdentifier has been set in FlowSnmpv2CVariableBinding
	HasObjectIdentifier() bool
	// Value returns FlowSnmpv2CVariableBindingValue, set in FlowSnmpv2CVariableBinding.
	Value() FlowSnmpv2CVariableBindingValue
	// SetValue assigns FlowSnmpv2CVariableBindingValue provided by user to FlowSnmpv2CVariableBinding.
	SetValue(value FlowSnmpv2CVariableBindingValue) FlowSnmpv2CVariableBinding
	// HasValue checks if Value has been set in FlowSnmpv2CVariableBinding
	HasValue() bool
	setNil()
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
// ObjectIdentifier returns a string
func (obj *flowSnmpv2CVariableBinding) ObjectIdentifier() string {

	return *obj.obj.ObjectIdentifier

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
// ObjectIdentifier returns a string
func (obj *flowSnmpv2CVariableBinding) HasObjectIdentifier() bool {
	return obj.obj.ObjectIdentifier != nil
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
// SetObjectIdentifier sets the string value in the FlowSnmpv2CVariableBinding object
func (obj *flowSnmpv2CVariableBinding) SetObjectIdentifier(value string) FlowSnmpv2CVariableBinding {

	obj.obj.ObjectIdentifier = &value
	return obj
}

// description is TBD
// Value returns a FlowSnmpv2CVariableBindingValue
func (obj *flowSnmpv2CVariableBinding) Value() FlowSnmpv2CVariableBindingValue {
	if obj.obj.Value == nil {
		obj.obj.Value = NewFlowSnmpv2CVariableBindingValue().msg()
	}
	if obj.valueHolder == nil {
		obj.valueHolder = &flowSnmpv2CVariableBindingValue{obj: obj.obj.Value}
	}
	return obj.valueHolder
}

// description is TBD
// Value returns a FlowSnmpv2CVariableBindingValue
func (obj *flowSnmpv2CVariableBinding) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the FlowSnmpv2CVariableBindingValue value in the FlowSnmpv2CVariableBinding object
func (obj *flowSnmpv2CVariableBinding) SetValue(value FlowSnmpv2CVariableBindingValue) FlowSnmpv2CVariableBinding {

	obj.valueHolder = nil
	obj.obj.Value = value.msg()

	return obj
}

func (obj *flowSnmpv2CVariableBinding) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ObjectIdentifier != nil {

		err := obj.validateOid(obj.ObjectIdentifier())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowSnmpv2CVariableBinding.ObjectIdentifier"))
		}

	}

	if obj.obj.Value != nil {

		obj.Value().validateObj(vObj, set_default)
	}

}

func (obj *flowSnmpv2CVariableBinding) setDefault() {
	if obj.obj.ObjectIdentifier == nil {
		obj.SetObjectIdentifier("0.1")
	}

}
