package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRouting *****
type flowIpv6SegmentRouting struct {
	validation
	obj                  *otg.FlowIpv6SegmentRouting
	marshaller           marshalFlowIpv6SegmentRouting
	unMarshaller         unMarshalFlowIpv6SegmentRouting
	segmentsLeftHolder   PatternFlowIpv6SegmentRoutingSegmentsLeft
	lastEntryHolder      PatternFlowIpv6SegmentRoutingLastEntry
	flagsHolder          FlowIpv6SegmentRoutingFlags
	tagHolder            PatternFlowIpv6SegmentRoutingTag
	segmentListHolder    FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter
	ingressNodeTlvHolder FlowIpv6SRHIngressNodeTlv
	egressNodeTlvHolder  FlowIpv6SRHEgressNodeTlv
	opaqueTlvHolder      FlowIpv6SRHOpaqueContainerTlv
	padTlvHolder         FlowIpv6SRHPadTlv
	pathTraceTlvHolder   FlowIpv6SRHPathTraceTlv
}

func NewFlowIpv6SegmentRouting() FlowIpv6SegmentRouting {
	obj := flowIpv6SegmentRouting{obj: &otg.FlowIpv6SegmentRouting{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRouting) msg() *otg.FlowIpv6SegmentRouting {
	return obj.obj
}

func (obj *flowIpv6SegmentRouting) setMsg(msg *otg.FlowIpv6SegmentRouting) FlowIpv6SegmentRouting {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRouting struct {
	obj *flowIpv6SegmentRouting
}

type marshalFlowIpv6SegmentRouting interface {
	// ToProto marshals FlowIpv6SegmentRouting to protobuf object *otg.FlowIpv6SegmentRouting
	ToProto() (*otg.FlowIpv6SegmentRouting, error)
	// ToPbText marshals FlowIpv6SegmentRouting to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRouting to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRouting to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRouting struct {
	obj *flowIpv6SegmentRouting
}

type unMarshalFlowIpv6SegmentRouting interface {
	// FromProto unmarshals FlowIpv6SegmentRouting from protobuf object *otg.FlowIpv6SegmentRouting
	FromProto(msg *otg.FlowIpv6SegmentRouting) (FlowIpv6SegmentRouting, error)
	// FromPbText unmarshals FlowIpv6SegmentRouting from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRouting from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRouting from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRouting) Marshal() marshalFlowIpv6SegmentRouting {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRouting{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRouting) Unmarshal() unMarshalFlowIpv6SegmentRouting {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRouting{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRouting) ToProto() (*otg.FlowIpv6SegmentRouting, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRouting) FromProto(msg *otg.FlowIpv6SegmentRouting) (FlowIpv6SegmentRouting, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRouting) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRouting) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRouting) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRouting) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRouting) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRouting) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRouting) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRouting) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRouting) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRouting) Clone() (FlowIpv6SegmentRouting, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRouting()
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

func (obj *flowIpv6SegmentRouting) setNil() {
	obj.segmentsLeftHolder = nil
	obj.lastEntryHolder = nil
	obj.flagsHolder = nil
	obj.tagHolder = nil
	obj.segmentListHolder = nil
	obj.ingressNodeTlvHolder = nil
	obj.egressNodeTlvHolder = nil
	obj.opaqueTlvHolder = nil
	obj.padTlvHolder = nil
	obj.pathTraceTlvHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRouting is iPv6 Segment Routing Header (SRH, Routing Type 4, RFC 8754 Section 2) carrying full 128-bit SRv6 SIDs. Each segment list entry is a complete SID (locator + function + argument) as a 128-bit IPv6 address. Segment list encoded in reverse path order: Segment[0] is the last hop, Segment[n-1] is the first hop (active, placed in outer IPv6 dst).
type FlowIpv6SegmentRouting interface {
	Validation
	// msg marshals FlowIpv6SegmentRouting to protobuf object *otg.FlowIpv6SegmentRouting
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRouting
	// setMsg unmarshals FlowIpv6SegmentRouting from protobuf object *otg.FlowIpv6SegmentRouting
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRouting) FlowIpv6SegmentRouting
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRouting
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRouting
	// validate validates FlowIpv6SegmentRouting
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRouting, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SegmentsLeft returns PatternFlowIpv6SegmentRoutingSegmentsLeft, set in FlowIpv6SegmentRouting.
	// PatternFlowIpv6SegmentRoutingSegmentsLeft is number of route segments remaining before the packet reaches its final destination (RFC 8754 Section 2.1). Decremented at each segment endpoint. When auto is assigned the value is set to the number of segments specified.
	SegmentsLeft() PatternFlowIpv6SegmentRoutingSegmentsLeft
	// SetSegmentsLeft assigns PatternFlowIpv6SegmentRoutingSegmentsLeft provided by user to FlowIpv6SegmentRouting.
	// PatternFlowIpv6SegmentRoutingSegmentsLeft is number of route segments remaining before the packet reaches its final destination (RFC 8754 Section 2.1). Decremented at each segment endpoint. When auto is assigned the value is set to the number of segments specified.
	SetSegmentsLeft(value PatternFlowIpv6SegmentRoutingSegmentsLeft) FlowIpv6SegmentRouting
	// HasSegmentsLeft checks if SegmentsLeft has been set in FlowIpv6SegmentRouting
	HasSegmentsLeft() bool
	// LastEntry returns PatternFlowIpv6SegmentRoutingLastEntry, set in FlowIpv6SegmentRouting.
	// PatternFlowIpv6SegmentRoutingLastEntry is zero-based index of the last element in the Segment List (RFC 8754 Section 2.1). For example, a list of 3 SIDs has Last Entry = 2. When auto is assigned the value is set to one less than the number of segments specified.
	LastEntry() PatternFlowIpv6SegmentRoutingLastEntry
	// SetLastEntry assigns PatternFlowIpv6SegmentRoutingLastEntry provided by user to FlowIpv6SegmentRouting.
	// PatternFlowIpv6SegmentRoutingLastEntry is zero-based index of the last element in the Segment List (RFC 8754 Section 2.1). For example, a list of 3 SIDs has Last Entry = 2. When auto is assigned the value is set to one less than the number of segments specified.
	SetLastEntry(value PatternFlowIpv6SegmentRoutingLastEntry) FlowIpv6SegmentRouting
	// HasLastEntry checks if LastEntry has been set in FlowIpv6SegmentRouting
	HasLastEntry() bool
	// Flags returns FlowIpv6SegmentRoutingFlags, set in FlowIpv6SegmentRouting.
	// FlowIpv6SegmentRoutingFlags is sRH Flags field (RFC 8754 Section 2.1). An 8-bit field; RFC 8754 marks all bits as reserved. IxNetwork exposes the following bits for OAM, HMAC, FRR, and protocol testing: Protected (FRR), Alert, O (OAM, RFC 9259), H (HMAC), and two reserved bits U1 and U2.
	Flags() FlowIpv6SegmentRoutingFlags
	// SetFlags assigns FlowIpv6SegmentRoutingFlags provided by user to FlowIpv6SegmentRouting.
	// FlowIpv6SegmentRoutingFlags is sRH Flags field (RFC 8754 Section 2.1). An 8-bit field; RFC 8754 marks all bits as reserved. IxNetwork exposes the following bits for OAM, HMAC, FRR, and protocol testing: Protected (FRR), Alert, O (OAM, RFC 9259), H (HMAC), and two reserved bits U1 and U2.
	SetFlags(value FlowIpv6SegmentRoutingFlags) FlowIpv6SegmentRouting
	// HasFlags checks if Flags has been set in FlowIpv6SegmentRouting
	HasFlags() bool
	// Tag returns PatternFlowIpv6SegmentRoutingTag, set in FlowIpv6SegmentRouting.
	// PatternFlowIpv6SegmentRoutingTag is 16-bit field used to tag a packet as part of a class or group (RFC 8754 Section 2.1). If not used it MUST be set to zero.
	Tag() PatternFlowIpv6SegmentRoutingTag
	// SetTag assigns PatternFlowIpv6SegmentRoutingTag provided by user to FlowIpv6SegmentRouting.
	// PatternFlowIpv6SegmentRoutingTag is 16-bit field used to tag a packet as part of a class or group (RFC 8754 Section 2.1). If not used it MUST be set to zero.
	SetTag(value PatternFlowIpv6SegmentRoutingTag) FlowIpv6SegmentRouting
	// HasTag checks if Tag has been set in FlowIpv6SegmentRouting
	HasTag() bool
	// SegmentList returns FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIterIter, set in FlowIpv6SegmentRouting
	SegmentList() FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter
	// IngressNodeTlv returns FlowIpv6SRHIngressNodeTlv, set in FlowIpv6SegmentRouting.
	// FlowIpv6SRHIngressNodeTlv is sRH Ingress Node TLV (type 1, RFC 9259 Section 3.1). Identifies the global IPv6 address of the SRv6 ingress node that imposed the SRH on the packet. OAM-capable endpoint nodes (those that process packets with the O-flag set) inspect this TLV to verify that the packet entered the SR domain at the expected node. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 9259 Section 3.1.
	IngressNodeTlv() FlowIpv6SRHIngressNodeTlv
	// SetIngressNodeTlv assigns FlowIpv6SRHIngressNodeTlv provided by user to FlowIpv6SegmentRouting.
	// FlowIpv6SRHIngressNodeTlv is sRH Ingress Node TLV (type 1, RFC 9259 Section 3.1). Identifies the global IPv6 address of the SRv6 ingress node that imposed the SRH on the packet. OAM-capable endpoint nodes (those that process packets with the O-flag set) inspect this TLV to verify that the packet entered the SR domain at the expected node. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 9259 Section 3.1.
	SetIngressNodeTlv(value FlowIpv6SRHIngressNodeTlv) FlowIpv6SegmentRouting
	// HasIngressNodeTlv checks if IngressNodeTlv has been set in FlowIpv6SegmentRouting
	HasIngressNodeTlv() bool
	// EgressNodeTlv returns FlowIpv6SRHEgressNodeTlv, set in FlowIpv6SegmentRouting.
	// FlowIpv6SRHEgressNodeTlv is sRH Egress Node TLV (type 2, RFC 9259 Section 3.2). Identifies the global IPv6 address of the intended SRv6 egress node - the last segment endpoint of the SR policy. OAM-capable endpoint nodes inspect this TLV to verify that the packet will exit the SR domain at the expected node. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 9259 Section 3.2.
	EgressNodeTlv() FlowIpv6SRHEgressNodeTlv
	// SetEgressNodeTlv assigns FlowIpv6SRHEgressNodeTlv provided by user to FlowIpv6SegmentRouting.
	// FlowIpv6SRHEgressNodeTlv is sRH Egress Node TLV (type 2, RFC 9259 Section 3.2). Identifies the global IPv6 address of the intended SRv6 egress node - the last segment endpoint of the SR policy. OAM-capable endpoint nodes inspect this TLV to verify that the packet will exit the SR domain at the expected node. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 9259 Section 3.2.
	SetEgressNodeTlv(value FlowIpv6SRHEgressNodeTlv) FlowIpv6SegmentRouting
	// HasEgressNodeTlv checks if EgressNodeTlv has been set in FlowIpv6SegmentRouting
	HasEgressNodeTlv() bool
	// OpaqueTlv returns FlowIpv6SRHOpaqueContainerTlv, set in FlowIpv6SegmentRouting.
	// FlowIpv6SRHOpaqueContainerTlv is sRH Opaque Container TLV (type 3, RFC 8754 Section 2.1). Carries implementation-specific or application-defined opaque data in the SRH TLV section. Transit nodes do not interpret the contents. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 8754 Section 2.1.
	OpaqueTlv() FlowIpv6SRHOpaqueContainerTlv
	// SetOpaqueTlv assigns FlowIpv6SRHOpaqueContainerTlv provided by user to FlowIpv6SegmentRouting.
	// FlowIpv6SRHOpaqueContainerTlv is sRH Opaque Container TLV (type 3, RFC 8754 Section 2.1). Carries implementation-specific or application-defined opaque data in the SRH TLV section. Transit nodes do not interpret the contents. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 8754 Section 2.1.
	SetOpaqueTlv(value FlowIpv6SRHOpaqueContainerTlv) FlowIpv6SegmentRouting
	// HasOpaqueTlv checks if OpaqueTlv has been set in FlowIpv6SegmentRouting
	HasOpaqueTlv() bool
	// PadTlv returns FlowIpv6SRHPadTlv, set in FlowIpv6SegmentRouting.
	// FlowIpv6SRHPadTlv is sRH Padding TLV (type 4, RFC 8754 Section 2.1). Used to align the SRH TLV block to an 8-byte boundary. The padding bytes are set to zero and are skipped by transit nodes. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 8754 Section 2.1.
	PadTlv() FlowIpv6SRHPadTlv
	// SetPadTlv assigns FlowIpv6SRHPadTlv provided by user to FlowIpv6SegmentRouting.
	// FlowIpv6SRHPadTlv is sRH Padding TLV (type 4, RFC 8754 Section 2.1). Used to align the SRH TLV block to an 8-byte boundary. The padding bytes are set to zero and are skipped by transit nodes. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 8754 Section 2.1.
	SetPadTlv(value FlowIpv6SRHPadTlv) FlowIpv6SegmentRouting
	// HasPadTlv checks if PadTlv has been set in FlowIpv6SegmentRouting
	HasPadTlv() bool
	// PathTraceTlv returns FlowIpv6SRHPathTraceTlv, set in FlowIpv6SegmentRouting.
	// FlowIpv6SRHPathTraceTlv is sRH Path Trace TLV (type 128, draft-ietf-spring-srv6-path-tracing). Records path information as the packet traverses SRv6 segment endpoints. The ingress node initializes the TLV; each subsequent SRv6 endpoint appends an 8-byte block containing its node identity, timestamp, and performance monitoring data. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: draft-ietf-spring-srv6-path-tracing.
	PathTraceTlv() FlowIpv6SRHPathTraceTlv
	// SetPathTraceTlv assigns FlowIpv6SRHPathTraceTlv provided by user to FlowIpv6SegmentRouting.
	// FlowIpv6SRHPathTraceTlv is sRH Path Trace TLV (type 128, draft-ietf-spring-srv6-path-tracing). Records path information as the packet traverses SRv6 segment endpoints. The ingress node initializes the TLV; each subsequent SRv6 endpoint appends an 8-byte block containing its node identity, timestamp, and performance monitoring data. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: draft-ietf-spring-srv6-path-tracing.
	SetPathTraceTlv(value FlowIpv6SRHPathTraceTlv) FlowIpv6SegmentRouting
	// HasPathTraceTlv checks if PathTraceTlv has been set in FlowIpv6SegmentRouting
	HasPathTraceTlv() bool
	setNil()
}

// description is TBD
// SegmentsLeft returns a PatternFlowIpv6SegmentRoutingSegmentsLeft
func (obj *flowIpv6SegmentRouting) SegmentsLeft() PatternFlowIpv6SegmentRoutingSegmentsLeft {
	if obj.obj.SegmentsLeft == nil {
		obj.obj.SegmentsLeft = NewPatternFlowIpv6SegmentRoutingSegmentsLeft().msg()
	}
	if obj.segmentsLeftHolder == nil {
		obj.segmentsLeftHolder = &patternFlowIpv6SegmentRoutingSegmentsLeft{obj: obj.obj.SegmentsLeft}
	}
	return obj.segmentsLeftHolder
}

// description is TBD
// SegmentsLeft returns a PatternFlowIpv6SegmentRoutingSegmentsLeft
func (obj *flowIpv6SegmentRouting) HasSegmentsLeft() bool {
	return obj.obj.SegmentsLeft != nil
}

// description is TBD
// SetSegmentsLeft sets the PatternFlowIpv6SegmentRoutingSegmentsLeft value in the FlowIpv6SegmentRouting object
func (obj *flowIpv6SegmentRouting) SetSegmentsLeft(value PatternFlowIpv6SegmentRoutingSegmentsLeft) FlowIpv6SegmentRouting {

	obj.segmentsLeftHolder = nil
	obj.obj.SegmentsLeft = value.msg()

	return obj
}

// description is TBD
// LastEntry returns a PatternFlowIpv6SegmentRoutingLastEntry
func (obj *flowIpv6SegmentRouting) LastEntry() PatternFlowIpv6SegmentRoutingLastEntry {
	if obj.obj.LastEntry == nil {
		obj.obj.LastEntry = NewPatternFlowIpv6SegmentRoutingLastEntry().msg()
	}
	if obj.lastEntryHolder == nil {
		obj.lastEntryHolder = &patternFlowIpv6SegmentRoutingLastEntry{obj: obj.obj.LastEntry}
	}
	return obj.lastEntryHolder
}

// description is TBD
// LastEntry returns a PatternFlowIpv6SegmentRoutingLastEntry
func (obj *flowIpv6SegmentRouting) HasLastEntry() bool {
	return obj.obj.LastEntry != nil
}

// description is TBD
// SetLastEntry sets the PatternFlowIpv6SegmentRoutingLastEntry value in the FlowIpv6SegmentRouting object
func (obj *flowIpv6SegmentRouting) SetLastEntry(value PatternFlowIpv6SegmentRoutingLastEntry) FlowIpv6SegmentRouting {

	obj.lastEntryHolder = nil
	obj.obj.LastEntry = value.msg()

	return obj
}

// description is TBD
// Flags returns a FlowIpv6SegmentRoutingFlags
func (obj *flowIpv6SegmentRouting) Flags() FlowIpv6SegmentRoutingFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewFlowIpv6SegmentRoutingFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &flowIpv6SegmentRoutingFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// description is TBD
// Flags returns a FlowIpv6SegmentRoutingFlags
func (obj *flowIpv6SegmentRouting) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the FlowIpv6SegmentRoutingFlags value in the FlowIpv6SegmentRouting object
func (obj *flowIpv6SegmentRouting) SetFlags(value FlowIpv6SegmentRoutingFlags) FlowIpv6SegmentRouting {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// description is TBD
// Tag returns a PatternFlowIpv6SegmentRoutingTag
func (obj *flowIpv6SegmentRouting) Tag() PatternFlowIpv6SegmentRoutingTag {
	if obj.obj.Tag == nil {
		obj.obj.Tag = NewPatternFlowIpv6SegmentRoutingTag().msg()
	}
	if obj.tagHolder == nil {
		obj.tagHolder = &patternFlowIpv6SegmentRoutingTag{obj: obj.obj.Tag}
	}
	return obj.tagHolder
}

// description is TBD
// Tag returns a PatternFlowIpv6SegmentRoutingTag
func (obj *flowIpv6SegmentRouting) HasTag() bool {
	return obj.obj.Tag != nil
}

// description is TBD
// SetTag sets the PatternFlowIpv6SegmentRoutingTag value in the FlowIpv6SegmentRouting object
func (obj *flowIpv6SegmentRouting) SetTag(value PatternFlowIpv6SegmentRoutingTag) FlowIpv6SegmentRouting {

	obj.tagHolder = nil
	obj.obj.Tag = value.msg()

	return obj
}

// List of 128-bit SRv6 SID addresses (RFC 8754 Section 2.1), encoded in reverse path order: Segment[0] is the last segment of the SR Policy, Segment[n-1] is the first segment to visit.
// SegmentList returns a []FlowIpv6SegmentRoutingSegment
func (obj *flowIpv6SegmentRouting) SegmentList() FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter {
	if len(obj.obj.SegmentList) == 0 {
		obj.obj.SegmentList = []*otg.FlowIpv6SegmentRoutingSegment{}
	}
	if obj.segmentListHolder == nil {
		obj.segmentListHolder = newFlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter(&obj.obj.SegmentList).setMsg(obj)
	}
	return obj.segmentListHolder
}

type flowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter struct {
	obj                                *flowIpv6SegmentRouting
	flowIpv6SegmentRoutingSegmentSlice []FlowIpv6SegmentRoutingSegment
	fieldPtr                           *[]*otg.FlowIpv6SegmentRoutingSegment
}

func newFlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter(ptr *[]*otg.FlowIpv6SegmentRoutingSegment) FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter {
	return &flowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter{fieldPtr: ptr}
}

type FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter interface {
	setMsg(*flowIpv6SegmentRouting) FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter
	Items() []FlowIpv6SegmentRoutingSegment
	Add() FlowIpv6SegmentRoutingSegment
	Append(items ...FlowIpv6SegmentRoutingSegment) FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter
	Set(index int, newObj FlowIpv6SegmentRoutingSegment) FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter
	Clear() FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter
	clearHolderSlice() FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter
	appendHolderSlice(item FlowIpv6SegmentRoutingSegment) FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter
}

func (obj *flowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter) setMsg(msg *flowIpv6SegmentRouting) FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowIpv6SegmentRoutingSegment{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter) Items() []FlowIpv6SegmentRoutingSegment {
	return obj.flowIpv6SegmentRoutingSegmentSlice
}

func (obj *flowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter) Add() FlowIpv6SegmentRoutingSegment {
	newObj := &otg.FlowIpv6SegmentRoutingSegment{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowIpv6SegmentRoutingSegment{obj: newObj}
	newLibObj.setDefault()
	obj.flowIpv6SegmentRoutingSegmentSlice = append(obj.flowIpv6SegmentRoutingSegmentSlice, newLibObj)
	return newLibObj
}

func (obj *flowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter) Append(items ...FlowIpv6SegmentRoutingSegment) FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowIpv6SegmentRoutingSegmentSlice = append(obj.flowIpv6SegmentRoutingSegmentSlice, item)
	}
	return obj
}

func (obj *flowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter) Set(index int, newObj FlowIpv6SegmentRoutingSegment) FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowIpv6SegmentRoutingSegmentSlice[index] = newObj
	return obj
}
func (obj *flowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter) Clear() FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowIpv6SegmentRoutingSegment{}
		obj.flowIpv6SegmentRoutingSegmentSlice = []FlowIpv6SegmentRoutingSegment{}
	}
	return obj
}
func (obj *flowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter) clearHolderSlice() FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter {
	if len(obj.flowIpv6SegmentRoutingSegmentSlice) > 0 {
		obj.flowIpv6SegmentRoutingSegmentSlice = []FlowIpv6SegmentRoutingSegment{}
	}
	return obj
}
func (obj *flowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter) appendHolderSlice(item FlowIpv6SegmentRoutingSegment) FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter {
	obj.flowIpv6SegmentRoutingSegmentSlice = append(obj.flowIpv6SegmentRoutingSegmentSlice, item)
	return obj
}

// When present, includes an Ingress Node TLV (type 1, RFC 9259 Section 3.1) in the SRH TLV section. Omit to suppress.
// IngressNodeTlv returns a FlowIpv6SRHIngressNodeTlv
func (obj *flowIpv6SegmentRouting) IngressNodeTlv() FlowIpv6SRHIngressNodeTlv {
	if obj.obj.IngressNodeTlv == nil {
		obj.obj.IngressNodeTlv = NewFlowIpv6SRHIngressNodeTlv().msg()
	}
	if obj.ingressNodeTlvHolder == nil {
		obj.ingressNodeTlvHolder = &flowIpv6SRHIngressNodeTlv{obj: obj.obj.IngressNodeTlv}
	}
	return obj.ingressNodeTlvHolder
}

// When present, includes an Ingress Node TLV (type 1, RFC 9259 Section 3.1) in the SRH TLV section. Omit to suppress.
// IngressNodeTlv returns a FlowIpv6SRHIngressNodeTlv
func (obj *flowIpv6SegmentRouting) HasIngressNodeTlv() bool {
	return obj.obj.IngressNodeTlv != nil
}

// When present, includes an Ingress Node TLV (type 1, RFC 9259 Section 3.1) in the SRH TLV section. Omit to suppress.
// SetIngressNodeTlv sets the FlowIpv6SRHIngressNodeTlv value in the FlowIpv6SegmentRouting object
func (obj *flowIpv6SegmentRouting) SetIngressNodeTlv(value FlowIpv6SRHIngressNodeTlv) FlowIpv6SegmentRouting {

	obj.ingressNodeTlvHolder = nil
	obj.obj.IngressNodeTlv = value.msg()

	return obj
}

// When present, includes an Egress Node TLV (type 2, RFC 9259 Section 3.2) in the SRH TLV section. Omit to suppress.
// EgressNodeTlv returns a FlowIpv6SRHEgressNodeTlv
func (obj *flowIpv6SegmentRouting) EgressNodeTlv() FlowIpv6SRHEgressNodeTlv {
	if obj.obj.EgressNodeTlv == nil {
		obj.obj.EgressNodeTlv = NewFlowIpv6SRHEgressNodeTlv().msg()
	}
	if obj.egressNodeTlvHolder == nil {
		obj.egressNodeTlvHolder = &flowIpv6SRHEgressNodeTlv{obj: obj.obj.EgressNodeTlv}
	}
	return obj.egressNodeTlvHolder
}

// When present, includes an Egress Node TLV (type 2, RFC 9259 Section 3.2) in the SRH TLV section. Omit to suppress.
// EgressNodeTlv returns a FlowIpv6SRHEgressNodeTlv
func (obj *flowIpv6SegmentRouting) HasEgressNodeTlv() bool {
	return obj.obj.EgressNodeTlv != nil
}

// When present, includes an Egress Node TLV (type 2, RFC 9259 Section 3.2) in the SRH TLV section. Omit to suppress.
// SetEgressNodeTlv sets the FlowIpv6SRHEgressNodeTlv value in the FlowIpv6SegmentRouting object
func (obj *flowIpv6SegmentRouting) SetEgressNodeTlv(value FlowIpv6SRHEgressNodeTlv) FlowIpv6SegmentRouting {

	obj.egressNodeTlvHolder = nil
	obj.obj.EgressNodeTlv = value.msg()

	return obj
}

// When present, includes an Opaque Container TLV (type 3, RFC 8754 Section 2.1) in the SRH TLV section. Omit to suppress.
// OpaqueTlv returns a FlowIpv6SRHOpaqueContainerTlv
func (obj *flowIpv6SegmentRouting) OpaqueTlv() FlowIpv6SRHOpaqueContainerTlv {
	if obj.obj.OpaqueTlv == nil {
		obj.obj.OpaqueTlv = NewFlowIpv6SRHOpaqueContainerTlv().msg()
	}
	if obj.opaqueTlvHolder == nil {
		obj.opaqueTlvHolder = &flowIpv6SRHOpaqueContainerTlv{obj: obj.obj.OpaqueTlv}
	}
	return obj.opaqueTlvHolder
}

// When present, includes an Opaque Container TLV (type 3, RFC 8754 Section 2.1) in the SRH TLV section. Omit to suppress.
// OpaqueTlv returns a FlowIpv6SRHOpaqueContainerTlv
func (obj *flowIpv6SegmentRouting) HasOpaqueTlv() bool {
	return obj.obj.OpaqueTlv != nil
}

// When present, includes an Opaque Container TLV (type 3, RFC 8754 Section 2.1) in the SRH TLV section. Omit to suppress.
// SetOpaqueTlv sets the FlowIpv6SRHOpaqueContainerTlv value in the FlowIpv6SegmentRouting object
func (obj *flowIpv6SegmentRouting) SetOpaqueTlv(value FlowIpv6SRHOpaqueContainerTlv) FlowIpv6SegmentRouting {

	obj.opaqueTlvHolder = nil
	obj.obj.OpaqueTlv = value.msg()

	return obj
}

// When present, includes a Padding TLV (type 4, RFC 8754 Section 2.1) in the SRH TLV section to align the TLV block to an 8-byte boundary. Omit to suppress.
// PadTlv returns a FlowIpv6SRHPadTlv
func (obj *flowIpv6SegmentRouting) PadTlv() FlowIpv6SRHPadTlv {
	if obj.obj.PadTlv == nil {
		obj.obj.PadTlv = NewFlowIpv6SRHPadTlv().msg()
	}
	if obj.padTlvHolder == nil {
		obj.padTlvHolder = &flowIpv6SRHPadTlv{obj: obj.obj.PadTlv}
	}
	return obj.padTlvHolder
}

// When present, includes a Padding TLV (type 4, RFC 8754 Section 2.1) in the SRH TLV section to align the TLV block to an 8-byte boundary. Omit to suppress.
// PadTlv returns a FlowIpv6SRHPadTlv
func (obj *flowIpv6SegmentRouting) HasPadTlv() bool {
	return obj.obj.PadTlv != nil
}

// When present, includes a Padding TLV (type 4, RFC 8754 Section 2.1) in the SRH TLV section to align the TLV block to an 8-byte boundary. Omit to suppress.
// SetPadTlv sets the FlowIpv6SRHPadTlv value in the FlowIpv6SegmentRouting object
func (obj *flowIpv6SegmentRouting) SetPadTlv(value FlowIpv6SRHPadTlv) FlowIpv6SegmentRouting {

	obj.padTlvHolder = nil
	obj.obj.PadTlv = value.msg()

	return obj
}

// When present, includes a Path Trace TLV (type 128, draft-ietf-spring-srv6-path-tracing) in the SRH TLV section. Each SRv6 endpoint on the path appends an 8-byte block containing node identity, timestamp, and performance monitoring data. Omit to suppress.
// PathTraceTlv returns a FlowIpv6SRHPathTraceTlv
func (obj *flowIpv6SegmentRouting) PathTraceTlv() FlowIpv6SRHPathTraceTlv {
	if obj.obj.PathTraceTlv == nil {
		obj.obj.PathTraceTlv = NewFlowIpv6SRHPathTraceTlv().msg()
	}
	if obj.pathTraceTlvHolder == nil {
		obj.pathTraceTlvHolder = &flowIpv6SRHPathTraceTlv{obj: obj.obj.PathTraceTlv}
	}
	return obj.pathTraceTlvHolder
}

// When present, includes a Path Trace TLV (type 128, draft-ietf-spring-srv6-path-tracing) in the SRH TLV section. Each SRv6 endpoint on the path appends an 8-byte block containing node identity, timestamp, and performance monitoring data. Omit to suppress.
// PathTraceTlv returns a FlowIpv6SRHPathTraceTlv
func (obj *flowIpv6SegmentRouting) HasPathTraceTlv() bool {
	return obj.obj.PathTraceTlv != nil
}

// When present, includes a Path Trace TLV (type 128, draft-ietf-spring-srv6-path-tracing) in the SRH TLV section. Each SRv6 endpoint on the path appends an 8-byte block containing node identity, timestamp, and performance monitoring data. Omit to suppress.
// SetPathTraceTlv sets the FlowIpv6SRHPathTraceTlv value in the FlowIpv6SegmentRouting object
func (obj *flowIpv6SegmentRouting) SetPathTraceTlv(value FlowIpv6SRHPathTraceTlv) FlowIpv6SegmentRouting {

	obj.pathTraceTlvHolder = nil
	obj.obj.PathTraceTlv = value.msg()

	return obj
}

func (obj *flowIpv6SegmentRouting) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SegmentsLeft != nil {

		obj.SegmentsLeft().validateObj(vObj, set_default)
	}

	if obj.obj.LastEntry != nil {

		obj.LastEntry().validateObj(vObj, set_default)
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.Tag != nil {

		obj.Tag().validateObj(vObj, set_default)
	}

	if len(obj.obj.SegmentList) != 0 {

		if set_default {
			obj.SegmentList().clearHolderSlice()
			for _, item := range obj.obj.SegmentList {
				obj.SegmentList().appendHolderSlice(&flowIpv6SegmentRoutingSegment{obj: item})
			}
		}
		for _, item := range obj.SegmentList().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.IngressNodeTlv != nil {

		obj.IngressNodeTlv().validateObj(vObj, set_default)
	}

	if obj.obj.EgressNodeTlv != nil {

		obj.EgressNodeTlv().validateObj(vObj, set_default)
	}

	if obj.obj.OpaqueTlv != nil {

		obj.OpaqueTlv().validateObj(vObj, set_default)
	}

	if obj.obj.PadTlv != nil {

		obj.PadTlv().validateObj(vObj, set_default)
	}

	if obj.obj.PathTraceTlv != nil {

		obj.PathTraceTlv().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6SegmentRouting) setDefault() {

}
