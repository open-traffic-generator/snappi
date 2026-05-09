package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SRHIngressNodeTlv *****
type flowIpv6SRHIngressNodeTlv struct {
	validation
	obj            *otg.FlowIpv6SRHIngressNodeTlv
	marshaller     marshalFlowIpv6SRHIngressNodeTlv
	unMarshaller   unMarshalFlowIpv6SRHIngressNodeTlv
	typeHolder     PatternFlowIpv6SRHIngressNodeTlvType
	lengthHolder   PatternFlowIpv6SRHIngressNodeTlvLength
	reservedHolder PatternFlowIpv6SRHIngressNodeTlvReserved
	valueHolder    PatternFlowIpv6SRHIngressNodeTlvValue
}

func NewFlowIpv6SRHIngressNodeTlv() FlowIpv6SRHIngressNodeTlv {
	obj := flowIpv6SRHIngressNodeTlv{obj: &otg.FlowIpv6SRHIngressNodeTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SRHIngressNodeTlv) msg() *otg.FlowIpv6SRHIngressNodeTlv {
	return obj.obj
}

func (obj *flowIpv6SRHIngressNodeTlv) setMsg(msg *otg.FlowIpv6SRHIngressNodeTlv) FlowIpv6SRHIngressNodeTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SRHIngressNodeTlv struct {
	obj *flowIpv6SRHIngressNodeTlv
}

type marshalFlowIpv6SRHIngressNodeTlv interface {
	// ToProto marshals FlowIpv6SRHIngressNodeTlv to protobuf object *otg.FlowIpv6SRHIngressNodeTlv
	ToProto() (*otg.FlowIpv6SRHIngressNodeTlv, error)
	// ToPbText marshals FlowIpv6SRHIngressNodeTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SRHIngressNodeTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SRHIngressNodeTlv to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SRHIngressNodeTlv struct {
	obj *flowIpv6SRHIngressNodeTlv
}

type unMarshalFlowIpv6SRHIngressNodeTlv interface {
	// FromProto unmarshals FlowIpv6SRHIngressNodeTlv from protobuf object *otg.FlowIpv6SRHIngressNodeTlv
	FromProto(msg *otg.FlowIpv6SRHIngressNodeTlv) (FlowIpv6SRHIngressNodeTlv, error)
	// FromPbText unmarshals FlowIpv6SRHIngressNodeTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SRHIngressNodeTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SRHIngressNodeTlv from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SRHIngressNodeTlv) Marshal() marshalFlowIpv6SRHIngressNodeTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SRHIngressNodeTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SRHIngressNodeTlv) Unmarshal() unMarshalFlowIpv6SRHIngressNodeTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SRHIngressNodeTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SRHIngressNodeTlv) ToProto() (*otg.FlowIpv6SRHIngressNodeTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SRHIngressNodeTlv) FromProto(msg *otg.FlowIpv6SRHIngressNodeTlv) (FlowIpv6SRHIngressNodeTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SRHIngressNodeTlv) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SRHIngressNodeTlv) FromPbText(value string) error {
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

func (m *marshalflowIpv6SRHIngressNodeTlv) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SRHIngressNodeTlv) FromYaml(value string) error {
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

func (m *marshalflowIpv6SRHIngressNodeTlv) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SRHIngressNodeTlv) FromJson(value string) error {
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

func (obj *flowIpv6SRHIngressNodeTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SRHIngressNodeTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SRHIngressNodeTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SRHIngressNodeTlv) Clone() (FlowIpv6SRHIngressNodeTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SRHIngressNodeTlv()
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

func (obj *flowIpv6SRHIngressNodeTlv) setNil() {
	obj.typeHolder = nil
	obj.lengthHolder = nil
	obj.reservedHolder = nil
	obj.valueHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SRHIngressNodeTlv is sRH Ingress Node TLV (type 1, RFC 9259 Section 3.1). Identifies the global IPv6 address of the SRv6 ingress node that imposed the SRH on the packet. OAM-capable endpoint nodes (those that process packets with the O-flag set) inspect this TLV to verify that the packet entered the SR domain at the expected node. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 9259 Section 3.1.
type FlowIpv6SRHIngressNodeTlv interface {
	Validation
	// msg marshals FlowIpv6SRHIngressNodeTlv to protobuf object *otg.FlowIpv6SRHIngressNodeTlv
	// and doesn't set defaults
	msg() *otg.FlowIpv6SRHIngressNodeTlv
	// setMsg unmarshals FlowIpv6SRHIngressNodeTlv from protobuf object *otg.FlowIpv6SRHIngressNodeTlv
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SRHIngressNodeTlv) FlowIpv6SRHIngressNodeTlv
	// provides marshal interface
	Marshal() marshalFlowIpv6SRHIngressNodeTlv
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SRHIngressNodeTlv
	// validate validates FlowIpv6SRHIngressNodeTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SRHIngressNodeTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowIpv6SRHIngressNodeTlvType, set in FlowIpv6SRHIngressNodeTlv.
	// PatternFlowIpv6SRHIngressNodeTlvType is tLV type code (RFC 9259 Section 3.1). The canonical value is 1. Set to a different value only for negative or conformance testing.
	Type() PatternFlowIpv6SRHIngressNodeTlvType
	// SetType assigns PatternFlowIpv6SRHIngressNodeTlvType provided by user to FlowIpv6SRHIngressNodeTlv.
	// PatternFlowIpv6SRHIngressNodeTlvType is tLV type code (RFC 9259 Section 3.1). The canonical value is 1. Set to a different value only for negative or conformance testing.
	SetType(value PatternFlowIpv6SRHIngressNodeTlvType) FlowIpv6SRHIngressNodeTlv
	// HasType checks if Type has been set in FlowIpv6SRHIngressNodeTlv
	HasType() bool
	// Length returns PatternFlowIpv6SRHIngressNodeTlvLength, set in FlowIpv6SRHIngressNodeTlv.
	// PatternFlowIpv6SRHIngressNodeTlvLength is length of the TLV Value field in bytes (RFC 9259 Section 3.1). When auto is assigned the implementation sets this to 18: 1-byte Reserved + 1-byte Node ID Len (fixed at 16) + 16-byte IPv6 address.
	Length() PatternFlowIpv6SRHIngressNodeTlvLength
	// SetLength assigns PatternFlowIpv6SRHIngressNodeTlvLength provided by user to FlowIpv6SRHIngressNodeTlv.
	// PatternFlowIpv6SRHIngressNodeTlvLength is length of the TLV Value field in bytes (RFC 9259 Section 3.1). When auto is assigned the implementation sets this to 18: 1-byte Reserved + 1-byte Node ID Len (fixed at 16) + 16-byte IPv6 address.
	SetLength(value PatternFlowIpv6SRHIngressNodeTlvLength) FlowIpv6SRHIngressNodeTlv
	// HasLength checks if Length has been set in FlowIpv6SRHIngressNodeTlv
	HasLength() bool
	// Reserved returns PatternFlowIpv6SRHIngressNodeTlvReserved, set in FlowIpv6SRHIngressNodeTlv.
	// PatternFlowIpv6SRHIngressNodeTlvReserved is reserved byte following the Length field (RFC 9259 Section 3.1). Must be zero on transmission. The adjacent Node ID Length byte is fixed at 16 for IPv6 addresses and is not separately configurable.
	Reserved() PatternFlowIpv6SRHIngressNodeTlvReserved
	// SetReserved assigns PatternFlowIpv6SRHIngressNodeTlvReserved provided by user to FlowIpv6SRHIngressNodeTlv.
	// PatternFlowIpv6SRHIngressNodeTlvReserved is reserved byte following the Length field (RFC 9259 Section 3.1). Must be zero on transmission. The adjacent Node ID Length byte is fixed at 16 for IPv6 addresses and is not separately configurable.
	SetReserved(value PatternFlowIpv6SRHIngressNodeTlvReserved) FlowIpv6SRHIngressNodeTlv
	// HasReserved checks if Reserved has been set in FlowIpv6SRHIngressNodeTlv
	HasReserved() bool
	// Value returns PatternFlowIpv6SRHIngressNodeTlvValue, set in FlowIpv6SRHIngressNodeTlv.
	// PatternFlowIpv6SRHIngressNodeTlvValue is global IPv6 address of the SRv6 ingress node (RFC 9259 Section 3.1). This is the address of the node that imposed the SRH on the packet.
	Value() PatternFlowIpv6SRHIngressNodeTlvValue
	// SetValue assigns PatternFlowIpv6SRHIngressNodeTlvValue provided by user to FlowIpv6SRHIngressNodeTlv.
	// PatternFlowIpv6SRHIngressNodeTlvValue is global IPv6 address of the SRv6 ingress node (RFC 9259 Section 3.1). This is the address of the node that imposed the SRH on the packet.
	SetValue(value PatternFlowIpv6SRHIngressNodeTlvValue) FlowIpv6SRHIngressNodeTlv
	// HasValue checks if Value has been set in FlowIpv6SRHIngressNodeTlv
	HasValue() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowIpv6SRHIngressNodeTlvType
func (obj *flowIpv6SRHIngressNodeTlv) Type() PatternFlowIpv6SRHIngressNodeTlvType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowIpv6SRHIngressNodeTlvType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowIpv6SRHIngressNodeTlvType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowIpv6SRHIngressNodeTlvType
func (obj *flowIpv6SRHIngressNodeTlv) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowIpv6SRHIngressNodeTlvType value in the FlowIpv6SRHIngressNodeTlv object
func (obj *flowIpv6SRHIngressNodeTlv) SetType(value PatternFlowIpv6SRHIngressNodeTlvType) FlowIpv6SRHIngressNodeTlv {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowIpv6SRHIngressNodeTlvLength
func (obj *flowIpv6SRHIngressNodeTlv) Length() PatternFlowIpv6SRHIngressNodeTlvLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowIpv6SRHIngressNodeTlvLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowIpv6SRHIngressNodeTlvLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowIpv6SRHIngressNodeTlvLength
func (obj *flowIpv6SRHIngressNodeTlv) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowIpv6SRHIngressNodeTlvLength value in the FlowIpv6SRHIngressNodeTlv object
func (obj *flowIpv6SRHIngressNodeTlv) SetLength(value PatternFlowIpv6SRHIngressNodeTlvLength) FlowIpv6SRHIngressNodeTlv {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowIpv6SRHIngressNodeTlvReserved
func (obj *flowIpv6SRHIngressNodeTlv) Reserved() PatternFlowIpv6SRHIngressNodeTlvReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowIpv6SRHIngressNodeTlvReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowIpv6SRHIngressNodeTlvReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowIpv6SRHIngressNodeTlvReserved
func (obj *flowIpv6SRHIngressNodeTlv) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowIpv6SRHIngressNodeTlvReserved value in the FlowIpv6SRHIngressNodeTlv object
func (obj *flowIpv6SRHIngressNodeTlv) SetReserved(value PatternFlowIpv6SRHIngressNodeTlvReserved) FlowIpv6SRHIngressNodeTlv {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// Value returns a PatternFlowIpv6SRHIngressNodeTlvValue
func (obj *flowIpv6SRHIngressNodeTlv) Value() PatternFlowIpv6SRHIngressNodeTlvValue {
	if obj.obj.Value == nil {
		obj.obj.Value = NewPatternFlowIpv6SRHIngressNodeTlvValue().msg()
	}
	if obj.valueHolder == nil {
		obj.valueHolder = &patternFlowIpv6SRHIngressNodeTlvValue{obj: obj.obj.Value}
	}
	return obj.valueHolder
}

// description is TBD
// Value returns a PatternFlowIpv6SRHIngressNodeTlvValue
func (obj *flowIpv6SRHIngressNodeTlv) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the PatternFlowIpv6SRHIngressNodeTlvValue value in the FlowIpv6SRHIngressNodeTlv object
func (obj *flowIpv6SRHIngressNodeTlv) SetValue(value PatternFlowIpv6SRHIngressNodeTlvValue) FlowIpv6SRHIngressNodeTlv {

	obj.valueHolder = nil
	obj.obj.Value = value.msg()

	return obj
}

func (obj *flowIpv6SRHIngressNodeTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(vObj, set_default)
	}

	if obj.obj.Value != nil {

		obj.Value().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6SRHIngressNodeTlv) setDefault() {

}
