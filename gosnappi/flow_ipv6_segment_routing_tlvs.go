package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingTlvs *****
type flowIpv6SegmentRoutingTlvs struct {
	validation
	obj             *otg.FlowIpv6SegmentRoutingTlvs
	marshaller      marshalFlowIpv6SegmentRoutingTlvs
	unMarshaller    unMarshalFlowIpv6SegmentRoutingTlvs
	lengthHolder    PatternFlowIpv6SegmentRoutingTlvsLength
	ingressHolder   FlowIpv6SegmentRoutingTlvIngressNode
	egressHolder    FlowIpv6SegmentRoutingTlvEgressNode
	opaqueHolder    FlowIpv6SegmentRoutingTlvOpaqueContainer
	paddingHolder   FlowIpv6SegmentRoutingTlvPad
	pathTraceHolder FlowIpv6SegmentRoutingTlvPathTrace
}

func NewFlowIpv6SegmentRoutingTlvs() FlowIpv6SegmentRoutingTlvs {
	obj := flowIpv6SegmentRoutingTlvs{obj: &otg.FlowIpv6SegmentRoutingTlvs{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingTlvs) msg() *otg.FlowIpv6SegmentRoutingTlvs {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingTlvs) setMsg(msg *otg.FlowIpv6SegmentRoutingTlvs) FlowIpv6SegmentRoutingTlvs {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingTlvs struct {
	obj *flowIpv6SegmentRoutingTlvs
}

type marshalFlowIpv6SegmentRoutingTlvs interface {
	// ToProto marshals FlowIpv6SegmentRoutingTlvs to protobuf object *otg.FlowIpv6SegmentRoutingTlvs
	ToProto() (*otg.FlowIpv6SegmentRoutingTlvs, error)
	// ToPbText marshals FlowIpv6SegmentRoutingTlvs to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingTlvs to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingTlvs to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingTlvs struct {
	obj *flowIpv6SegmentRoutingTlvs
}

type unMarshalFlowIpv6SegmentRoutingTlvs interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingTlvs from protobuf object *otg.FlowIpv6SegmentRoutingTlvs
	FromProto(msg *otg.FlowIpv6SegmentRoutingTlvs) (FlowIpv6SegmentRoutingTlvs, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingTlvs from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingTlvs from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingTlvs from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingTlvs) Marshal() marshalFlowIpv6SegmentRoutingTlvs {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingTlvs{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingTlvs) Unmarshal() unMarshalFlowIpv6SegmentRoutingTlvs {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingTlvs{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingTlvs) ToProto() (*otg.FlowIpv6SegmentRoutingTlvs, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingTlvs) FromProto(msg *otg.FlowIpv6SegmentRoutingTlvs) (FlowIpv6SegmentRoutingTlvs, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingTlvs) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvs) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingTlvs) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvs) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingTlvs) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvs) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingTlvs) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingTlvs) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingTlvs) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingTlvs) Clone() (FlowIpv6SegmentRoutingTlvs, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingTlvs()
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

func (obj *flowIpv6SegmentRoutingTlvs) setNil() {
	obj.lengthHolder = nil
	obj.ingressHolder = nil
	obj.egressHolder = nil
	obj.opaqueHolder = nil
	obj.paddingHolder = nil
	obj.pathTraceHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRoutingTlvs is a list of optional Type-Length-Value (TLV) objects that carry metadata for segment processing. TLVs are only present if the Hdr Ext Len is greater than what is required for the fixed header and the segment list.
type FlowIpv6SegmentRoutingTlvs interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingTlvs to protobuf object *otg.FlowIpv6SegmentRoutingTlvs
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingTlvs
	// setMsg unmarshals FlowIpv6SegmentRoutingTlvs from protobuf object *otg.FlowIpv6SegmentRoutingTlvs
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingTlvs) FlowIpv6SegmentRoutingTlvs
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingTlvs
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingTlvs
	// validate validates FlowIpv6SegmentRoutingTlvs
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingTlvs, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIpv6SegmentRoutingTlvsChoiceEnum, set in FlowIpv6SegmentRoutingTlvs
	Choice() FlowIpv6SegmentRoutingTlvsChoiceEnum
	// setChoice assigns FlowIpv6SegmentRoutingTlvsChoiceEnum provided by user to FlowIpv6SegmentRoutingTlvs
	setChoice(value FlowIpv6SegmentRoutingTlvsChoiceEnum) FlowIpv6SegmentRoutingTlvs
	// HasChoice checks if Choice has been set in FlowIpv6SegmentRoutingTlvs
	HasChoice() bool
	// Length returns PatternFlowIpv6SegmentRoutingTlvsLength, set in FlowIpv6SegmentRoutingTlvs.
	// PatternFlowIpv6SegmentRoutingTlvsLength is the length of the TLV value in bytes. If auto is selected the implementation should automatically set the length field to the correct value as per the selected TLV and attributes.
	Length() PatternFlowIpv6SegmentRoutingTlvsLength
	// SetLength assigns PatternFlowIpv6SegmentRoutingTlvsLength provided by user to FlowIpv6SegmentRoutingTlvs.
	// PatternFlowIpv6SegmentRoutingTlvsLength is the length of the TLV value in bytes. If auto is selected the implementation should automatically set the length field to the correct value as per the selected TLV and attributes.
	SetLength(value PatternFlowIpv6SegmentRoutingTlvsLength) FlowIpv6SegmentRoutingTlvs
	// HasLength checks if Length has been set in FlowIpv6SegmentRoutingTlvs
	HasLength() bool
	// Ingress returns FlowIpv6SegmentRoutingTlvIngressNode, set in FlowIpv6SegmentRoutingTlvs.
	// FlowIpv6SegmentRoutingTlvIngressNode is a TLV that may be used to carry the IPv6 address of the ingress node of the SR domain.
	Ingress() FlowIpv6SegmentRoutingTlvIngressNode
	// SetIngress assigns FlowIpv6SegmentRoutingTlvIngressNode provided by user to FlowIpv6SegmentRoutingTlvs.
	// FlowIpv6SegmentRoutingTlvIngressNode is a TLV that may be used to carry the IPv6 address of the ingress node of the SR domain.
	SetIngress(value FlowIpv6SegmentRoutingTlvIngressNode) FlowIpv6SegmentRoutingTlvs
	// HasIngress checks if Ingress has been set in FlowIpv6SegmentRoutingTlvs
	HasIngress() bool
	// Egress returns FlowIpv6SegmentRoutingTlvEgressNode, set in FlowIpv6SegmentRoutingTlvs.
	// FlowIpv6SegmentRoutingTlvEgressNode is a TLV that identifies the egress node of the SR policy.
	Egress() FlowIpv6SegmentRoutingTlvEgressNode
	// SetEgress assigns FlowIpv6SegmentRoutingTlvEgressNode provided by user to FlowIpv6SegmentRoutingTlvs.
	// FlowIpv6SegmentRoutingTlvEgressNode is a TLV that identifies the egress node of the SR policy.
	SetEgress(value FlowIpv6SegmentRoutingTlvEgressNode) FlowIpv6SegmentRoutingTlvs
	// HasEgress checks if Egress has been set in FlowIpv6SegmentRoutingTlvs
	HasEgress() bool
	// Opaque returns FlowIpv6SegmentRoutingTlvOpaqueContainer, set in FlowIpv6SegmentRoutingTlvs.
	// FlowIpv6SegmentRoutingTlvOpaqueContainer is a TLV used to carry opaque data across the SR domain.
	Opaque() FlowIpv6SegmentRoutingTlvOpaqueContainer
	// SetOpaque assigns FlowIpv6SegmentRoutingTlvOpaqueContainer provided by user to FlowIpv6SegmentRoutingTlvs.
	// FlowIpv6SegmentRoutingTlvOpaqueContainer is a TLV used to carry opaque data across the SR domain.
	SetOpaque(value FlowIpv6SegmentRoutingTlvOpaqueContainer) FlowIpv6SegmentRoutingTlvs
	// HasOpaque checks if Opaque has been set in FlowIpv6SegmentRoutingTlvs
	HasOpaque() bool
	// Padding returns FlowIpv6SegmentRoutingTlvPad, set in FlowIpv6SegmentRoutingTlvs.
	// FlowIpv6SegmentRoutingTlvPad is padding TLV used when more than one byte of padding is required to align subsequent TLVs or pad the header to an 8-octet boundary.
	Padding() FlowIpv6SegmentRoutingTlvPad
	// SetPadding assigns FlowIpv6SegmentRoutingTlvPad provided by user to FlowIpv6SegmentRoutingTlvs.
	// FlowIpv6SegmentRoutingTlvPad is padding TLV used when more than one byte of padding is required to align subsequent TLVs or pad the header to an 8-octet boundary.
	SetPadding(value FlowIpv6SegmentRoutingTlvPad) FlowIpv6SegmentRoutingTlvs
	// HasPadding checks if Padding has been set in FlowIpv6SegmentRoutingTlvs
	HasPadding() bool
	// PathTrace returns FlowIpv6SegmentRoutingTlvPathTrace, set in FlowIpv6SegmentRoutingTlvs.
	// FlowIpv6SegmentRoutingTlvPathTrace is a TLV used for path tracing and performance monitoring within the SR domain.
	PathTrace() FlowIpv6SegmentRoutingTlvPathTrace
	// SetPathTrace assigns FlowIpv6SegmentRoutingTlvPathTrace provided by user to FlowIpv6SegmentRoutingTlvs.
	// FlowIpv6SegmentRoutingTlvPathTrace is a TLV used for path tracing and performance monitoring within the SR domain.
	SetPathTrace(value FlowIpv6SegmentRoutingTlvPathTrace) FlowIpv6SegmentRoutingTlvs
	// HasPathTrace checks if PathTrace has been set in FlowIpv6SegmentRoutingTlvs
	HasPathTrace() bool
	setNil()
}

type FlowIpv6SegmentRoutingTlvsChoiceEnum string

// Enum of Choice on FlowIpv6SegmentRoutingTlvs
var FlowIpv6SegmentRoutingTlvsChoice = struct {
	INGRESS    FlowIpv6SegmentRoutingTlvsChoiceEnum
	EGRESS     FlowIpv6SegmentRoutingTlvsChoiceEnum
	OPAQUE     FlowIpv6SegmentRoutingTlvsChoiceEnum
	PADDING    FlowIpv6SegmentRoutingTlvsChoiceEnum
	PATH_TRACE FlowIpv6SegmentRoutingTlvsChoiceEnum
}{
	INGRESS:    FlowIpv6SegmentRoutingTlvsChoiceEnum("ingress"),
	EGRESS:     FlowIpv6SegmentRoutingTlvsChoiceEnum("egress"),
	OPAQUE:     FlowIpv6SegmentRoutingTlvsChoiceEnum("opaque"),
	PADDING:    FlowIpv6SegmentRoutingTlvsChoiceEnum("padding"),
	PATH_TRACE: FlowIpv6SegmentRoutingTlvsChoiceEnum("path_trace"),
}

func (obj *flowIpv6SegmentRoutingTlvs) Choice() FlowIpv6SegmentRoutingTlvsChoiceEnum {
	return FlowIpv6SegmentRoutingTlvsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowIpv6SegmentRoutingTlvs) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowIpv6SegmentRoutingTlvs) setChoice(value FlowIpv6SegmentRoutingTlvsChoiceEnum) FlowIpv6SegmentRoutingTlvs {
	intValue, ok := otg.FlowIpv6SegmentRoutingTlvs_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIpv6SegmentRoutingTlvsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIpv6SegmentRoutingTlvs_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.PathTrace = nil
	obj.pathTraceHolder = nil
	obj.obj.Padding = nil
	obj.paddingHolder = nil
	obj.obj.Opaque = nil
	obj.opaqueHolder = nil
	obj.obj.Egress = nil
	obj.egressHolder = nil
	obj.obj.Ingress = nil
	obj.ingressHolder = nil

	if value == FlowIpv6SegmentRoutingTlvsChoice.INGRESS {
		obj.obj.Ingress = NewFlowIpv6SegmentRoutingTlvIngressNode().msg()
	}

	if value == FlowIpv6SegmentRoutingTlvsChoice.EGRESS {
		obj.obj.Egress = NewFlowIpv6SegmentRoutingTlvEgressNode().msg()
	}

	if value == FlowIpv6SegmentRoutingTlvsChoice.OPAQUE {
		obj.obj.Opaque = NewFlowIpv6SegmentRoutingTlvOpaqueContainer().msg()
	}

	if value == FlowIpv6SegmentRoutingTlvsChoice.PADDING {
		obj.obj.Padding = NewFlowIpv6SegmentRoutingTlvPad().msg()
	}

	if value == FlowIpv6SegmentRoutingTlvsChoice.PATH_TRACE {
		obj.obj.PathTrace = NewFlowIpv6SegmentRoutingTlvPathTrace().msg()
	}

	return obj
}

// description is TBD
// Length returns a PatternFlowIpv6SegmentRoutingTlvsLength
func (obj *flowIpv6SegmentRoutingTlvs) Length() PatternFlowIpv6SegmentRoutingTlvsLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowIpv6SegmentRoutingTlvsLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowIpv6SegmentRoutingTlvsLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowIpv6SegmentRoutingTlvsLength
func (obj *flowIpv6SegmentRoutingTlvs) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowIpv6SegmentRoutingTlvsLength value in the FlowIpv6SegmentRoutingTlvs object
func (obj *flowIpv6SegmentRoutingTlvs) SetLength(value PatternFlowIpv6SegmentRoutingTlvsLength) FlowIpv6SegmentRoutingTlvs {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// Ingress returns a FlowIpv6SegmentRoutingTlvIngressNode
func (obj *flowIpv6SegmentRoutingTlvs) Ingress() FlowIpv6SegmentRoutingTlvIngressNode {
	if obj.obj.Ingress == nil {
		obj.setChoice(FlowIpv6SegmentRoutingTlvsChoice.INGRESS)
	}
	if obj.ingressHolder == nil {
		obj.ingressHolder = &flowIpv6SegmentRoutingTlvIngressNode{obj: obj.obj.Ingress}
	}
	return obj.ingressHolder
}

// description is TBD
// Ingress returns a FlowIpv6SegmentRoutingTlvIngressNode
func (obj *flowIpv6SegmentRoutingTlvs) HasIngress() bool {
	return obj.obj.Ingress != nil
}

// description is TBD
// SetIngress sets the FlowIpv6SegmentRoutingTlvIngressNode value in the FlowIpv6SegmentRoutingTlvs object
func (obj *flowIpv6SegmentRoutingTlvs) SetIngress(value FlowIpv6SegmentRoutingTlvIngressNode) FlowIpv6SegmentRoutingTlvs {
	obj.setChoice(FlowIpv6SegmentRoutingTlvsChoice.INGRESS)
	obj.ingressHolder = nil
	obj.obj.Ingress = value.msg()

	return obj
}

// description is TBD
// Egress returns a FlowIpv6SegmentRoutingTlvEgressNode
func (obj *flowIpv6SegmentRoutingTlvs) Egress() FlowIpv6SegmentRoutingTlvEgressNode {
	if obj.obj.Egress == nil {
		obj.setChoice(FlowIpv6SegmentRoutingTlvsChoice.EGRESS)
	}
	if obj.egressHolder == nil {
		obj.egressHolder = &flowIpv6SegmentRoutingTlvEgressNode{obj: obj.obj.Egress}
	}
	return obj.egressHolder
}

// description is TBD
// Egress returns a FlowIpv6SegmentRoutingTlvEgressNode
func (obj *flowIpv6SegmentRoutingTlvs) HasEgress() bool {
	return obj.obj.Egress != nil
}

// description is TBD
// SetEgress sets the FlowIpv6SegmentRoutingTlvEgressNode value in the FlowIpv6SegmentRoutingTlvs object
func (obj *flowIpv6SegmentRoutingTlvs) SetEgress(value FlowIpv6SegmentRoutingTlvEgressNode) FlowIpv6SegmentRoutingTlvs {
	obj.setChoice(FlowIpv6SegmentRoutingTlvsChoice.EGRESS)
	obj.egressHolder = nil
	obj.obj.Egress = value.msg()

	return obj
}

// description is TBD
// Opaque returns a FlowIpv6SegmentRoutingTlvOpaqueContainer
func (obj *flowIpv6SegmentRoutingTlvs) Opaque() FlowIpv6SegmentRoutingTlvOpaqueContainer {
	if obj.obj.Opaque == nil {
		obj.setChoice(FlowIpv6SegmentRoutingTlvsChoice.OPAQUE)
	}
	if obj.opaqueHolder == nil {
		obj.opaqueHolder = &flowIpv6SegmentRoutingTlvOpaqueContainer{obj: obj.obj.Opaque}
	}
	return obj.opaqueHolder
}

// description is TBD
// Opaque returns a FlowIpv6SegmentRoutingTlvOpaqueContainer
func (obj *flowIpv6SegmentRoutingTlvs) HasOpaque() bool {
	return obj.obj.Opaque != nil
}

// description is TBD
// SetOpaque sets the FlowIpv6SegmentRoutingTlvOpaqueContainer value in the FlowIpv6SegmentRoutingTlvs object
func (obj *flowIpv6SegmentRoutingTlvs) SetOpaque(value FlowIpv6SegmentRoutingTlvOpaqueContainer) FlowIpv6SegmentRoutingTlvs {
	obj.setChoice(FlowIpv6SegmentRoutingTlvsChoice.OPAQUE)
	obj.opaqueHolder = nil
	obj.obj.Opaque = value.msg()

	return obj
}

// description is TBD
// Padding returns a FlowIpv6SegmentRoutingTlvPad
func (obj *flowIpv6SegmentRoutingTlvs) Padding() FlowIpv6SegmentRoutingTlvPad {
	if obj.obj.Padding == nil {
		obj.setChoice(FlowIpv6SegmentRoutingTlvsChoice.PADDING)
	}
	if obj.paddingHolder == nil {
		obj.paddingHolder = &flowIpv6SegmentRoutingTlvPad{obj: obj.obj.Padding}
	}
	return obj.paddingHolder
}

// description is TBD
// Padding returns a FlowIpv6SegmentRoutingTlvPad
func (obj *flowIpv6SegmentRoutingTlvs) HasPadding() bool {
	return obj.obj.Padding != nil
}

// description is TBD
// SetPadding sets the FlowIpv6SegmentRoutingTlvPad value in the FlowIpv6SegmentRoutingTlvs object
func (obj *flowIpv6SegmentRoutingTlvs) SetPadding(value FlowIpv6SegmentRoutingTlvPad) FlowIpv6SegmentRoutingTlvs {
	obj.setChoice(FlowIpv6SegmentRoutingTlvsChoice.PADDING)
	obj.paddingHolder = nil
	obj.obj.Padding = value.msg()

	return obj
}

// description is TBD
// PathTrace returns a FlowIpv6SegmentRoutingTlvPathTrace
func (obj *flowIpv6SegmentRoutingTlvs) PathTrace() FlowIpv6SegmentRoutingTlvPathTrace {
	if obj.obj.PathTrace == nil {
		obj.setChoice(FlowIpv6SegmentRoutingTlvsChoice.PATH_TRACE)
	}
	if obj.pathTraceHolder == nil {
		obj.pathTraceHolder = &flowIpv6SegmentRoutingTlvPathTrace{obj: obj.obj.PathTrace}
	}
	return obj.pathTraceHolder
}

// description is TBD
// PathTrace returns a FlowIpv6SegmentRoutingTlvPathTrace
func (obj *flowIpv6SegmentRoutingTlvs) HasPathTrace() bool {
	return obj.obj.PathTrace != nil
}

// description is TBD
// SetPathTrace sets the FlowIpv6SegmentRoutingTlvPathTrace value in the FlowIpv6SegmentRoutingTlvs object
func (obj *flowIpv6SegmentRoutingTlvs) SetPathTrace(value FlowIpv6SegmentRoutingTlvPathTrace) FlowIpv6SegmentRoutingTlvs {
	obj.setChoice(FlowIpv6SegmentRoutingTlvsChoice.PATH_TRACE)
	obj.pathTraceHolder = nil
	obj.obj.PathTrace = value.msg()

	return obj
}

func (obj *flowIpv6SegmentRoutingTlvs) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Ingress != nil {

		obj.Ingress().validateObj(vObj, set_default)
	}

	if obj.obj.Egress != nil {

		obj.Egress().validateObj(vObj, set_default)
	}

	if obj.obj.Opaque != nil {

		obj.Opaque().validateObj(vObj, set_default)
	}

	if obj.obj.Padding != nil {

		obj.Padding().validateObj(vObj, set_default)
	}

	if obj.obj.PathTrace != nil {

		obj.PathTrace().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6SegmentRoutingTlvs) setDefault() {
	var choices_set int = 0
	var choice FlowIpv6SegmentRoutingTlvsChoiceEnum

	if obj.obj.Ingress != nil {
		choices_set += 1
		choice = FlowIpv6SegmentRoutingTlvsChoice.INGRESS
	}

	if obj.obj.Egress != nil {
		choices_set += 1
		choice = FlowIpv6SegmentRoutingTlvsChoice.EGRESS
	}

	if obj.obj.Opaque != nil {
		choices_set += 1
		choice = FlowIpv6SegmentRoutingTlvsChoice.OPAQUE
	}

	if obj.obj.Padding != nil {
		choices_set += 1
		choice = FlowIpv6SegmentRoutingTlvsChoice.PADDING
	}

	if obj.obj.PathTrace != nil {
		choices_set += 1
		choice = FlowIpv6SegmentRoutingTlvsChoice.PATH_TRACE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowIpv6SegmentRoutingTlvsChoice.INGRESS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowIpv6SegmentRoutingTlvs")
			}
		} else {
			intVal := otg.FlowIpv6SegmentRoutingTlvs_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowIpv6SegmentRoutingTlvs_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
