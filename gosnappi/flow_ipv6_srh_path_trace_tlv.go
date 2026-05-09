package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SRHPathTraceTlv *****
type flowIpv6SRHPathTraceTlv struct {
	validation
	obj             *otg.FlowIpv6SRHPathTraceTlv
	marshaller      marshalFlowIpv6SRHPathTraceTlv
	unMarshaller    unMarshalFlowIpv6SRHPathTraceTlv
	typeHolder      PatternFlowIpv6SRHPathTraceTlvType
	lengthHolder    PatternFlowIpv6SRHPathTraceTlvLength
	reservedHolder  PatternFlowIpv6SRHPathTraceTlvReserved
	timestampHolder PatternFlowIpv6SRHPathTraceTlvTimestamp
	pmDataHolder    PatternFlowIpv6SRHPathTraceTlvPmData
	valueHolder     PatternFlowIpv6SRHPathTraceTlvValue
}

func NewFlowIpv6SRHPathTraceTlv() FlowIpv6SRHPathTraceTlv {
	obj := flowIpv6SRHPathTraceTlv{obj: &otg.FlowIpv6SRHPathTraceTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SRHPathTraceTlv) msg() *otg.FlowIpv6SRHPathTraceTlv {
	return obj.obj
}

func (obj *flowIpv6SRHPathTraceTlv) setMsg(msg *otg.FlowIpv6SRHPathTraceTlv) FlowIpv6SRHPathTraceTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SRHPathTraceTlv struct {
	obj *flowIpv6SRHPathTraceTlv
}

type marshalFlowIpv6SRHPathTraceTlv interface {
	// ToProto marshals FlowIpv6SRHPathTraceTlv to protobuf object *otg.FlowIpv6SRHPathTraceTlv
	ToProto() (*otg.FlowIpv6SRHPathTraceTlv, error)
	// ToPbText marshals FlowIpv6SRHPathTraceTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SRHPathTraceTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SRHPathTraceTlv to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SRHPathTraceTlv struct {
	obj *flowIpv6SRHPathTraceTlv
}

type unMarshalFlowIpv6SRHPathTraceTlv interface {
	// FromProto unmarshals FlowIpv6SRHPathTraceTlv from protobuf object *otg.FlowIpv6SRHPathTraceTlv
	FromProto(msg *otg.FlowIpv6SRHPathTraceTlv) (FlowIpv6SRHPathTraceTlv, error)
	// FromPbText unmarshals FlowIpv6SRHPathTraceTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SRHPathTraceTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SRHPathTraceTlv from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SRHPathTraceTlv) Marshal() marshalFlowIpv6SRHPathTraceTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SRHPathTraceTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SRHPathTraceTlv) Unmarshal() unMarshalFlowIpv6SRHPathTraceTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SRHPathTraceTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SRHPathTraceTlv) ToProto() (*otg.FlowIpv6SRHPathTraceTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SRHPathTraceTlv) FromProto(msg *otg.FlowIpv6SRHPathTraceTlv) (FlowIpv6SRHPathTraceTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SRHPathTraceTlv) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SRHPathTraceTlv) FromPbText(value string) error {
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

func (m *marshalflowIpv6SRHPathTraceTlv) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SRHPathTraceTlv) FromYaml(value string) error {
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

func (m *marshalflowIpv6SRHPathTraceTlv) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SRHPathTraceTlv) FromJson(value string) error {
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

func (obj *flowIpv6SRHPathTraceTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SRHPathTraceTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SRHPathTraceTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SRHPathTraceTlv) Clone() (FlowIpv6SRHPathTraceTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SRHPathTraceTlv()
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

func (obj *flowIpv6SRHPathTraceTlv) setNil() {
	obj.typeHolder = nil
	obj.lengthHolder = nil
	obj.reservedHolder = nil
	obj.timestampHolder = nil
	obj.pmDataHolder = nil
	obj.valueHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SRHPathTraceTlv is sRH Path Trace TLV (type 128, draft-ietf-spring-srv6-path-tracing). Records path information as the packet traverses SRv6 segment endpoints. The ingress node initializes the TLV; each subsequent SRv6 endpoint appends an 8-byte block containing its node identity, timestamp, and performance monitoring data. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: draft-ietf-spring-srv6-path-tracing.
type FlowIpv6SRHPathTraceTlv interface {
	Validation
	// msg marshals FlowIpv6SRHPathTraceTlv to protobuf object *otg.FlowIpv6SRHPathTraceTlv
	// and doesn't set defaults
	msg() *otg.FlowIpv6SRHPathTraceTlv
	// setMsg unmarshals FlowIpv6SRHPathTraceTlv from protobuf object *otg.FlowIpv6SRHPathTraceTlv
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SRHPathTraceTlv) FlowIpv6SRHPathTraceTlv
	// provides marshal interface
	Marshal() marshalFlowIpv6SRHPathTraceTlv
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SRHPathTraceTlv
	// validate validates FlowIpv6SRHPathTraceTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SRHPathTraceTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowIpv6SRHPathTraceTlvType, set in FlowIpv6SRHPathTraceTlv.
	// PatternFlowIpv6SRHPathTraceTlvType is tLV type code (draft-ietf-spring-srv6-path-tracing). The canonical value is 128. Set to a different value only for negative or conformance testing.
	Type() PatternFlowIpv6SRHPathTraceTlvType
	// SetType assigns PatternFlowIpv6SRHPathTraceTlvType provided by user to FlowIpv6SRHPathTraceTlv.
	// PatternFlowIpv6SRHPathTraceTlvType is tLV type code (draft-ietf-spring-srv6-path-tracing). The canonical value is 128. Set to a different value only for negative or conformance testing.
	SetType(value PatternFlowIpv6SRHPathTraceTlvType) FlowIpv6SRHPathTraceTlv
	// HasType checks if Type has been set in FlowIpv6SRHPathTraceTlv
	HasType() bool
	// Length returns PatternFlowIpv6SRHPathTraceTlvLength, set in FlowIpv6SRHPathTraceTlv.
	// PatternFlowIpv6SRHPathTraceTlvLength is length of the TLV Value field in bytes (draft-ietf-spring-srv6-path-tracing). When auto is assigned the implementation derives this from the configured value fields.
	Length() PatternFlowIpv6SRHPathTraceTlvLength
	// SetLength assigns PatternFlowIpv6SRHPathTraceTlvLength provided by user to FlowIpv6SRHPathTraceTlv.
	// PatternFlowIpv6SRHPathTraceTlvLength is length of the TLV Value field in bytes (draft-ietf-spring-srv6-path-tracing). When auto is assigned the implementation derives this from the configured value fields.
	SetLength(value PatternFlowIpv6SRHPathTraceTlvLength) FlowIpv6SRHPathTraceTlv
	// HasLength checks if Length has been set in FlowIpv6SRHPathTraceTlv
	HasLength() bool
	// Reserved returns PatternFlowIpv6SRHPathTraceTlvReserved, set in FlowIpv6SRHPathTraceTlv.
	// PatternFlowIpv6SRHPathTraceTlvReserved is reserved field (draft-ietf-spring-srv6-path-tracing). Must be zero on transmission.
	Reserved() PatternFlowIpv6SRHPathTraceTlvReserved
	// SetReserved assigns PatternFlowIpv6SRHPathTraceTlvReserved provided by user to FlowIpv6SRHPathTraceTlv.
	// PatternFlowIpv6SRHPathTraceTlvReserved is reserved field (draft-ietf-spring-srv6-path-tracing). Must be zero on transmission.
	SetReserved(value PatternFlowIpv6SRHPathTraceTlvReserved) FlowIpv6SRHPathTraceTlv
	// HasReserved checks if Reserved has been set in FlowIpv6SRHPathTraceTlv
	HasReserved() bool
	// Timestamp returns PatternFlowIpv6SRHPathTraceTlvTimestamp, set in FlowIpv6SRHPathTraceTlv.
	// PatternFlowIpv6SRHPathTraceTlvTimestamp is timestamp inserted by the ingress node (draft-ietf-spring-srv6-path-tracing). Encoded as a 32-bit truncated PTP timestamp value.
	Timestamp() PatternFlowIpv6SRHPathTraceTlvTimestamp
	// SetTimestamp assigns PatternFlowIpv6SRHPathTraceTlvTimestamp provided by user to FlowIpv6SRHPathTraceTlv.
	// PatternFlowIpv6SRHPathTraceTlvTimestamp is timestamp inserted by the ingress node (draft-ietf-spring-srv6-path-tracing). Encoded as a 32-bit truncated PTP timestamp value.
	SetTimestamp(value PatternFlowIpv6SRHPathTraceTlvTimestamp) FlowIpv6SRHPathTraceTlv
	// HasTimestamp checks if Timestamp has been set in FlowIpv6SRHPathTraceTlv
	HasTimestamp() bool
	// PmData returns PatternFlowIpv6SRHPathTraceTlvPmData, set in FlowIpv6SRHPathTraceTlv.
	// PatternFlowIpv6SRHPathTraceTlvPmData is performance Monitoring (PM) data inserted by the ingress node (draft-ietf-spring-srv6-path-tracing). Carries interface and load information encoded as a 32-bit value.
	PmData() PatternFlowIpv6SRHPathTraceTlvPmData
	// SetPmData assigns PatternFlowIpv6SRHPathTraceTlvPmData provided by user to FlowIpv6SRHPathTraceTlv.
	// PatternFlowIpv6SRHPathTraceTlvPmData is performance Monitoring (PM) data inserted by the ingress node (draft-ietf-spring-srv6-path-tracing). Carries interface and load information encoded as a 32-bit value.
	SetPmData(value PatternFlowIpv6SRHPathTraceTlvPmData) FlowIpv6SRHPathTraceTlv
	// HasPmData checks if PmData has been set in FlowIpv6SRHPathTraceTlv
	HasPmData() bool
	// Value returns PatternFlowIpv6SRHPathTraceTlvValue, set in FlowIpv6SRHPathTraceTlv.
	// PatternFlowIpv6SRHPathTraceTlvValue is initial node identification value inserted by the ingress node (draft-ietf-spring-srv6-path-tracing). Identifies the ingress node within the path trace record.
	Value() PatternFlowIpv6SRHPathTraceTlvValue
	// SetValue assigns PatternFlowIpv6SRHPathTraceTlvValue provided by user to FlowIpv6SRHPathTraceTlv.
	// PatternFlowIpv6SRHPathTraceTlvValue is initial node identification value inserted by the ingress node (draft-ietf-spring-srv6-path-tracing). Identifies the ingress node within the path trace record.
	SetValue(value PatternFlowIpv6SRHPathTraceTlvValue) FlowIpv6SRHPathTraceTlv
	// HasValue checks if Value has been set in FlowIpv6SRHPathTraceTlv
	HasValue() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowIpv6SRHPathTraceTlvType
func (obj *flowIpv6SRHPathTraceTlv) Type() PatternFlowIpv6SRHPathTraceTlvType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowIpv6SRHPathTraceTlvType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowIpv6SRHPathTraceTlvType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowIpv6SRHPathTraceTlvType
func (obj *flowIpv6SRHPathTraceTlv) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowIpv6SRHPathTraceTlvType value in the FlowIpv6SRHPathTraceTlv object
func (obj *flowIpv6SRHPathTraceTlv) SetType(value PatternFlowIpv6SRHPathTraceTlvType) FlowIpv6SRHPathTraceTlv {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowIpv6SRHPathTraceTlvLength
func (obj *flowIpv6SRHPathTraceTlv) Length() PatternFlowIpv6SRHPathTraceTlvLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowIpv6SRHPathTraceTlvLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowIpv6SRHPathTraceTlvLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowIpv6SRHPathTraceTlvLength
func (obj *flowIpv6SRHPathTraceTlv) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowIpv6SRHPathTraceTlvLength value in the FlowIpv6SRHPathTraceTlv object
func (obj *flowIpv6SRHPathTraceTlv) SetLength(value PatternFlowIpv6SRHPathTraceTlvLength) FlowIpv6SRHPathTraceTlv {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowIpv6SRHPathTraceTlvReserved
func (obj *flowIpv6SRHPathTraceTlv) Reserved() PatternFlowIpv6SRHPathTraceTlvReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowIpv6SRHPathTraceTlvReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowIpv6SRHPathTraceTlvReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowIpv6SRHPathTraceTlvReserved
func (obj *flowIpv6SRHPathTraceTlv) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowIpv6SRHPathTraceTlvReserved value in the FlowIpv6SRHPathTraceTlv object
func (obj *flowIpv6SRHPathTraceTlv) SetReserved(value PatternFlowIpv6SRHPathTraceTlvReserved) FlowIpv6SRHPathTraceTlv {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// Timestamp returns a PatternFlowIpv6SRHPathTraceTlvTimestamp
func (obj *flowIpv6SRHPathTraceTlv) Timestamp() PatternFlowIpv6SRHPathTraceTlvTimestamp {
	if obj.obj.Timestamp == nil {
		obj.obj.Timestamp = NewPatternFlowIpv6SRHPathTraceTlvTimestamp().msg()
	}
	if obj.timestampHolder == nil {
		obj.timestampHolder = &patternFlowIpv6SRHPathTraceTlvTimestamp{obj: obj.obj.Timestamp}
	}
	return obj.timestampHolder
}

// description is TBD
// Timestamp returns a PatternFlowIpv6SRHPathTraceTlvTimestamp
func (obj *flowIpv6SRHPathTraceTlv) HasTimestamp() bool {
	return obj.obj.Timestamp != nil
}

// description is TBD
// SetTimestamp sets the PatternFlowIpv6SRHPathTraceTlvTimestamp value in the FlowIpv6SRHPathTraceTlv object
func (obj *flowIpv6SRHPathTraceTlv) SetTimestamp(value PatternFlowIpv6SRHPathTraceTlvTimestamp) FlowIpv6SRHPathTraceTlv {

	obj.timestampHolder = nil
	obj.obj.Timestamp = value.msg()

	return obj
}

// description is TBD
// PmData returns a PatternFlowIpv6SRHPathTraceTlvPmData
func (obj *flowIpv6SRHPathTraceTlv) PmData() PatternFlowIpv6SRHPathTraceTlvPmData {
	if obj.obj.PmData == nil {
		obj.obj.PmData = NewPatternFlowIpv6SRHPathTraceTlvPmData().msg()
	}
	if obj.pmDataHolder == nil {
		obj.pmDataHolder = &patternFlowIpv6SRHPathTraceTlvPmData{obj: obj.obj.PmData}
	}
	return obj.pmDataHolder
}

// description is TBD
// PmData returns a PatternFlowIpv6SRHPathTraceTlvPmData
func (obj *flowIpv6SRHPathTraceTlv) HasPmData() bool {
	return obj.obj.PmData != nil
}

// description is TBD
// SetPmData sets the PatternFlowIpv6SRHPathTraceTlvPmData value in the FlowIpv6SRHPathTraceTlv object
func (obj *flowIpv6SRHPathTraceTlv) SetPmData(value PatternFlowIpv6SRHPathTraceTlvPmData) FlowIpv6SRHPathTraceTlv {

	obj.pmDataHolder = nil
	obj.obj.PmData = value.msg()

	return obj
}

// description is TBD
// Value returns a PatternFlowIpv6SRHPathTraceTlvValue
func (obj *flowIpv6SRHPathTraceTlv) Value() PatternFlowIpv6SRHPathTraceTlvValue {
	if obj.obj.Value == nil {
		obj.obj.Value = NewPatternFlowIpv6SRHPathTraceTlvValue().msg()
	}
	if obj.valueHolder == nil {
		obj.valueHolder = &patternFlowIpv6SRHPathTraceTlvValue{obj: obj.obj.Value}
	}
	return obj.valueHolder
}

// description is TBD
// Value returns a PatternFlowIpv6SRHPathTraceTlvValue
func (obj *flowIpv6SRHPathTraceTlv) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the PatternFlowIpv6SRHPathTraceTlvValue value in the FlowIpv6SRHPathTraceTlv object
func (obj *flowIpv6SRHPathTraceTlv) SetValue(value PatternFlowIpv6SRHPathTraceTlvValue) FlowIpv6SRHPathTraceTlv {

	obj.valueHolder = nil
	obj.obj.Value = value.msg()

	return obj
}

func (obj *flowIpv6SRHPathTraceTlv) validateObj(vObj *validation, set_default bool) {
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

	if obj.obj.Timestamp != nil {

		obj.Timestamp().validateObj(vObj, set_default)
	}

	if obj.obj.PmData != nil {

		obj.PmData().validateObj(vObj, set_default)
	}

	if obj.obj.Value != nil {

		obj.Value().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6SRHPathTraceTlv) setDefault() {

}
