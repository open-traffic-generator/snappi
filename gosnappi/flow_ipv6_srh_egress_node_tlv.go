package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SRHEgressNodeTlv *****
type flowIpv6SRHEgressNodeTlv struct {
	validation
	obj            *otg.FlowIpv6SRHEgressNodeTlv
	marshaller     marshalFlowIpv6SRHEgressNodeTlv
	unMarshaller   unMarshalFlowIpv6SRHEgressNodeTlv
	typeHolder     PatternFlowIpv6SRHEgressNodeTlvType
	lengthHolder   PatternFlowIpv6SRHEgressNodeTlvLength
	reservedHolder PatternFlowIpv6SRHEgressNodeTlvReserved
	valueHolder    PatternFlowIpv6SRHEgressNodeTlvValue
}

func NewFlowIpv6SRHEgressNodeTlv() FlowIpv6SRHEgressNodeTlv {
	obj := flowIpv6SRHEgressNodeTlv{obj: &otg.FlowIpv6SRHEgressNodeTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SRHEgressNodeTlv) msg() *otg.FlowIpv6SRHEgressNodeTlv {
	return obj.obj
}

func (obj *flowIpv6SRHEgressNodeTlv) setMsg(msg *otg.FlowIpv6SRHEgressNodeTlv) FlowIpv6SRHEgressNodeTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SRHEgressNodeTlv struct {
	obj *flowIpv6SRHEgressNodeTlv
}

type marshalFlowIpv6SRHEgressNodeTlv interface {
	// ToProto marshals FlowIpv6SRHEgressNodeTlv to protobuf object *otg.FlowIpv6SRHEgressNodeTlv
	ToProto() (*otg.FlowIpv6SRHEgressNodeTlv, error)
	// ToPbText marshals FlowIpv6SRHEgressNodeTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SRHEgressNodeTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SRHEgressNodeTlv to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SRHEgressNodeTlv struct {
	obj *flowIpv6SRHEgressNodeTlv
}

type unMarshalFlowIpv6SRHEgressNodeTlv interface {
	// FromProto unmarshals FlowIpv6SRHEgressNodeTlv from protobuf object *otg.FlowIpv6SRHEgressNodeTlv
	FromProto(msg *otg.FlowIpv6SRHEgressNodeTlv) (FlowIpv6SRHEgressNodeTlv, error)
	// FromPbText unmarshals FlowIpv6SRHEgressNodeTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SRHEgressNodeTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SRHEgressNodeTlv from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SRHEgressNodeTlv) Marshal() marshalFlowIpv6SRHEgressNodeTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SRHEgressNodeTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SRHEgressNodeTlv) Unmarshal() unMarshalFlowIpv6SRHEgressNodeTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SRHEgressNodeTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SRHEgressNodeTlv) ToProto() (*otg.FlowIpv6SRHEgressNodeTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SRHEgressNodeTlv) FromProto(msg *otg.FlowIpv6SRHEgressNodeTlv) (FlowIpv6SRHEgressNodeTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SRHEgressNodeTlv) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SRHEgressNodeTlv) FromPbText(value string) error {
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

func (m *marshalflowIpv6SRHEgressNodeTlv) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SRHEgressNodeTlv) FromYaml(value string) error {
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

func (m *marshalflowIpv6SRHEgressNodeTlv) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SRHEgressNodeTlv) FromJson(value string) error {
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

func (obj *flowIpv6SRHEgressNodeTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SRHEgressNodeTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SRHEgressNodeTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SRHEgressNodeTlv) Clone() (FlowIpv6SRHEgressNodeTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SRHEgressNodeTlv()
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

func (obj *flowIpv6SRHEgressNodeTlv) setNil() {
	obj.typeHolder = nil
	obj.lengthHolder = nil
	obj.reservedHolder = nil
	obj.valueHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SRHEgressNodeTlv is sRH Egress Node TLV (type 2, RFC 9259 Section 3.2). Identifies the global IPv6 address of the intended SRv6 egress node - the last segment endpoint of the SR policy. OAM-capable endpoint nodes inspect this TLV to verify that the packet will exit the SR domain at the expected node. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 9259 Section 3.2.
type FlowIpv6SRHEgressNodeTlv interface {
	Validation
	// msg marshals FlowIpv6SRHEgressNodeTlv to protobuf object *otg.FlowIpv6SRHEgressNodeTlv
	// and doesn't set defaults
	msg() *otg.FlowIpv6SRHEgressNodeTlv
	// setMsg unmarshals FlowIpv6SRHEgressNodeTlv from protobuf object *otg.FlowIpv6SRHEgressNodeTlv
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SRHEgressNodeTlv) FlowIpv6SRHEgressNodeTlv
	// provides marshal interface
	Marshal() marshalFlowIpv6SRHEgressNodeTlv
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SRHEgressNodeTlv
	// validate validates FlowIpv6SRHEgressNodeTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SRHEgressNodeTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowIpv6SRHEgressNodeTlvType, set in FlowIpv6SRHEgressNodeTlv.
	// PatternFlowIpv6SRHEgressNodeTlvType is tLV type code (RFC 9259 Section 3.2). The canonical value is 2. Set to a different value only for negative or conformance testing.
	Type() PatternFlowIpv6SRHEgressNodeTlvType
	// SetType assigns PatternFlowIpv6SRHEgressNodeTlvType provided by user to FlowIpv6SRHEgressNodeTlv.
	// PatternFlowIpv6SRHEgressNodeTlvType is tLV type code (RFC 9259 Section 3.2). The canonical value is 2. Set to a different value only for negative or conformance testing.
	SetType(value PatternFlowIpv6SRHEgressNodeTlvType) FlowIpv6SRHEgressNodeTlv
	// HasType checks if Type has been set in FlowIpv6SRHEgressNodeTlv
	HasType() bool
	// Length returns PatternFlowIpv6SRHEgressNodeTlvLength, set in FlowIpv6SRHEgressNodeTlv.
	// PatternFlowIpv6SRHEgressNodeTlvLength is length of the TLV Value field in bytes (RFC 9259 Section 3.2). When auto is assigned the implementation sets this to 18: 1-byte Reserved + 1-byte Node ID Len (fixed at 16) + 16-byte IPv6 address.
	Length() PatternFlowIpv6SRHEgressNodeTlvLength
	// SetLength assigns PatternFlowIpv6SRHEgressNodeTlvLength provided by user to FlowIpv6SRHEgressNodeTlv.
	// PatternFlowIpv6SRHEgressNodeTlvLength is length of the TLV Value field in bytes (RFC 9259 Section 3.2). When auto is assigned the implementation sets this to 18: 1-byte Reserved + 1-byte Node ID Len (fixed at 16) + 16-byte IPv6 address.
	SetLength(value PatternFlowIpv6SRHEgressNodeTlvLength) FlowIpv6SRHEgressNodeTlv
	// HasLength checks if Length has been set in FlowIpv6SRHEgressNodeTlv
	HasLength() bool
	// Reserved returns PatternFlowIpv6SRHEgressNodeTlvReserved, set in FlowIpv6SRHEgressNodeTlv.
	// PatternFlowIpv6SRHEgressNodeTlvReserved is reserved byte following the Length field (RFC 9259 Section 3.2). Must be zero on transmission. The adjacent Node ID Length byte is fixed at 16 for IPv6 addresses and is not separately configurable.
	Reserved() PatternFlowIpv6SRHEgressNodeTlvReserved
	// SetReserved assigns PatternFlowIpv6SRHEgressNodeTlvReserved provided by user to FlowIpv6SRHEgressNodeTlv.
	// PatternFlowIpv6SRHEgressNodeTlvReserved is reserved byte following the Length field (RFC 9259 Section 3.2). Must be zero on transmission. The adjacent Node ID Length byte is fixed at 16 for IPv6 addresses and is not separately configurable.
	SetReserved(value PatternFlowIpv6SRHEgressNodeTlvReserved) FlowIpv6SRHEgressNodeTlv
	// HasReserved checks if Reserved has been set in FlowIpv6SRHEgressNodeTlv
	HasReserved() bool
	// Value returns PatternFlowIpv6SRHEgressNodeTlvValue, set in FlowIpv6SRHEgressNodeTlv.
	// PatternFlowIpv6SRHEgressNodeTlvValue is global IPv6 address of the intended SRv6 egress node (RFC 9259 Section 3.2). This is the address of the final segment endpoint of the SR policy.
	Value() PatternFlowIpv6SRHEgressNodeTlvValue
	// SetValue assigns PatternFlowIpv6SRHEgressNodeTlvValue provided by user to FlowIpv6SRHEgressNodeTlv.
	// PatternFlowIpv6SRHEgressNodeTlvValue is global IPv6 address of the intended SRv6 egress node (RFC 9259 Section 3.2). This is the address of the final segment endpoint of the SR policy.
	SetValue(value PatternFlowIpv6SRHEgressNodeTlvValue) FlowIpv6SRHEgressNodeTlv
	// HasValue checks if Value has been set in FlowIpv6SRHEgressNodeTlv
	HasValue() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowIpv6SRHEgressNodeTlvType
func (obj *flowIpv6SRHEgressNodeTlv) Type() PatternFlowIpv6SRHEgressNodeTlvType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowIpv6SRHEgressNodeTlvType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowIpv6SRHEgressNodeTlvType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowIpv6SRHEgressNodeTlvType
func (obj *flowIpv6SRHEgressNodeTlv) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowIpv6SRHEgressNodeTlvType value in the FlowIpv6SRHEgressNodeTlv object
func (obj *flowIpv6SRHEgressNodeTlv) SetType(value PatternFlowIpv6SRHEgressNodeTlvType) FlowIpv6SRHEgressNodeTlv {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowIpv6SRHEgressNodeTlvLength
func (obj *flowIpv6SRHEgressNodeTlv) Length() PatternFlowIpv6SRHEgressNodeTlvLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowIpv6SRHEgressNodeTlvLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowIpv6SRHEgressNodeTlvLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowIpv6SRHEgressNodeTlvLength
func (obj *flowIpv6SRHEgressNodeTlv) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowIpv6SRHEgressNodeTlvLength value in the FlowIpv6SRHEgressNodeTlv object
func (obj *flowIpv6SRHEgressNodeTlv) SetLength(value PatternFlowIpv6SRHEgressNodeTlvLength) FlowIpv6SRHEgressNodeTlv {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowIpv6SRHEgressNodeTlvReserved
func (obj *flowIpv6SRHEgressNodeTlv) Reserved() PatternFlowIpv6SRHEgressNodeTlvReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowIpv6SRHEgressNodeTlvReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowIpv6SRHEgressNodeTlvReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowIpv6SRHEgressNodeTlvReserved
func (obj *flowIpv6SRHEgressNodeTlv) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowIpv6SRHEgressNodeTlvReserved value in the FlowIpv6SRHEgressNodeTlv object
func (obj *flowIpv6SRHEgressNodeTlv) SetReserved(value PatternFlowIpv6SRHEgressNodeTlvReserved) FlowIpv6SRHEgressNodeTlv {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// Value returns a PatternFlowIpv6SRHEgressNodeTlvValue
func (obj *flowIpv6SRHEgressNodeTlv) Value() PatternFlowIpv6SRHEgressNodeTlvValue {
	if obj.obj.Value == nil {
		obj.obj.Value = NewPatternFlowIpv6SRHEgressNodeTlvValue().msg()
	}
	if obj.valueHolder == nil {
		obj.valueHolder = &patternFlowIpv6SRHEgressNodeTlvValue{obj: obj.obj.Value}
	}
	return obj.valueHolder
}

// description is TBD
// Value returns a PatternFlowIpv6SRHEgressNodeTlvValue
func (obj *flowIpv6SRHEgressNodeTlv) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the PatternFlowIpv6SRHEgressNodeTlvValue value in the FlowIpv6SRHEgressNodeTlv object
func (obj *flowIpv6SRHEgressNodeTlv) SetValue(value PatternFlowIpv6SRHEgressNodeTlvValue) FlowIpv6SRHEgressNodeTlv {

	obj.valueHolder = nil
	obj.obj.Value = value.msg()

	return obj
}

func (obj *flowIpv6SRHEgressNodeTlv) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowIpv6SRHEgressNodeTlv) setDefault() {

}
