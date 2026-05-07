package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingUsid *****
type flowIpv6SegmentRoutingUsid struct {
	validation
	obj                *otg.FlowIpv6SegmentRoutingUsid
	marshaller         marshalFlowIpv6SegmentRoutingUsid
	unMarshaller       unMarshalFlowIpv6SegmentRoutingUsid
	segmentsLeftHolder PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	lastEntryHolder    PatternFlowIpv6SegmentRoutingUsidLastEntry
	flagsHolder        FlowIpv6SegmentRoutingFlags
	tagHolder          PatternFlowIpv6SegmentRoutingUsidTag
	segmentListHolder  FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter
}

func NewFlowIpv6SegmentRoutingUsid() FlowIpv6SegmentRoutingUsid {
	obj := flowIpv6SegmentRoutingUsid{obj: &otg.FlowIpv6SegmentRoutingUsid{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingUsid) msg() *otg.FlowIpv6SegmentRoutingUsid {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingUsid) setMsg(msg *otg.FlowIpv6SegmentRoutingUsid) FlowIpv6SegmentRoutingUsid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingUsid struct {
	obj *flowIpv6SegmentRoutingUsid
}

type marshalFlowIpv6SegmentRoutingUsid interface {
	// ToProto marshals FlowIpv6SegmentRoutingUsid to protobuf object *otg.FlowIpv6SegmentRoutingUsid
	ToProto() (*otg.FlowIpv6SegmentRoutingUsid, error)
	// ToPbText marshals FlowIpv6SegmentRoutingUsid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingUsid to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingUsid to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingUsid struct {
	obj *flowIpv6SegmentRoutingUsid
}

type unMarshalFlowIpv6SegmentRoutingUsid interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingUsid from protobuf object *otg.FlowIpv6SegmentRoutingUsid
	FromProto(msg *otg.FlowIpv6SegmentRoutingUsid) (FlowIpv6SegmentRoutingUsid, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingUsid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingUsid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingUsid from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingUsid) Marshal() marshalFlowIpv6SegmentRoutingUsid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingUsid{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingUsid) Unmarshal() unMarshalFlowIpv6SegmentRoutingUsid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingUsid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingUsid) ToProto() (*otg.FlowIpv6SegmentRoutingUsid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingUsid) FromProto(msg *otg.FlowIpv6SegmentRoutingUsid) (FlowIpv6SegmentRoutingUsid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingUsid) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsid) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingUsid) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsid) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingUsid) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsid) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingUsid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUsid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUsid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingUsid) Clone() (FlowIpv6SegmentRoutingUsid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingUsid()
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

func (obj *flowIpv6SegmentRoutingUsid) setNil() {
	obj.segmentsLeftHolder = nil
	obj.lastEntryHolder = nil
	obj.flagsHolder = nil
	obj.tagHolder = nil
	obj.segmentListHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRoutingUsid is iPv6 Segment Routing Header (SRH) with uSID containers (RFC 9800 Section 4).
// Each segment list entry is a pre-computed 128-bit uSID container value
// supplied as a plain IPv6 address. The container packs multiple Micro-SIDs
// but appears on the wire as a single 128-bit address - no decomposition
// fields (locator block length, node length, function length) are present
// in the packet. The user pre-computes the container value and supplies it
// directly, allowing raw traffic generation independent of any control
// plane configuration on the emulated device.
// The Routing Type field is always 4 (RFC 8754 Section 2.1).
//
// Packet stack example - 1 container (F3216, Node2 + Node3 packed into one container):
//
// Container value assembly (done by user, outside the schema):
// lb=fc00::/32 | node2-fn1 (0002:0001) | node3-fn1 (0003:0001) | zeros
// => fc00:0:2:1:3:1::
//
// ```
// +--------------------------------------------------+
// | Ethernet                                         |
// +--------------------------------------------------+
// | Outer IPv6                                       |
// |   next_header : 43 (Routing Extension Header)   |
// |   dst         : fc00:0:2:1:3:1:: (container)    |
// +--------------------------------------------------+
// | SRH (Routing Type 4)                            |
// |   segments_left : 0  (auto)                     |
// |   last_entry    : 0  (auto)                     |
// |   flags, tag    : 0                             |
// |   Segment[0]    : fc00:0:2:1:3:1::              |
// +--------------------------------------------------+
// | Inner payload (IPv4, IPv6, etc.)                 |
// +--------------------------------------------------+
// ```
//
// At Node2: reads first uSID (node2-fn1), shifts container left,
// outer dst becomes fc00:0:3:1::, forwards.
// At Node3: reads next uSID (node3-fn1), container exhausted,
// strips SRH, delivers payload.
//
// Packet stack example - 2 containers (3-hop path: Node2 + Node3, then Node4):
//
// ```
// +--------------------------------------------------+
// | Ethernet                                         |
// +--------------------------------------------------+
// | Outer IPv6                                       |
// |   dst : fc00:0:2:1:3:1:: (Segment[1], active)   |
// +--------------------------------------------------+
// | SRH (Routing Type 4)                            |
// |   segments_left : 1  (auto)                     |
// |   last_entry    : 1  (auto)                     |
// |   flags, tag    : 0                             |
// |   Segment[0]    : fc00:0:4:1::      (last hop)  |
// |   Segment[1]    : fc00:0:2:1:3:1::  (first)     |
// +--------------------------------------------------+
// | Inner payload (IPv4, IPv6, etc.)                 |
// +--------------------------------------------------+
// ```
//
// The segment list is encoded in reverse path order: Segment[0] is the
// last container (last hop), Segment[n-1] is the first container to visit.
type FlowIpv6SegmentRoutingUsid interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingUsid to protobuf object *otg.FlowIpv6SegmentRoutingUsid
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingUsid
	// setMsg unmarshals FlowIpv6SegmentRoutingUsid from protobuf object *otg.FlowIpv6SegmentRoutingUsid
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingUsid) FlowIpv6SegmentRoutingUsid
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingUsid
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingUsid
	// validate validates FlowIpv6SegmentRoutingUsid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingUsid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SegmentsLeft returns PatternFlowIpv6SegmentRoutingUsidSegmentsLeft, set in FlowIpv6SegmentRoutingUsid.
	// PatternFlowIpv6SegmentRoutingUsidSegmentsLeft is number of uSID containers remaining before the packet reaches its final destination (RFC 8754 Section 2.1). Decremented when the active container is exhausted and the next container becomes active. When auto is assigned the value is set to the number of containers specified.
	SegmentsLeft() PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	// SetSegmentsLeft assigns PatternFlowIpv6SegmentRoutingUsidSegmentsLeft provided by user to FlowIpv6SegmentRoutingUsid.
	// PatternFlowIpv6SegmentRoutingUsidSegmentsLeft is number of uSID containers remaining before the packet reaches its final destination (RFC 8754 Section 2.1). Decremented when the active container is exhausted and the next container becomes active. When auto is assigned the value is set to the number of containers specified.
	SetSegmentsLeft(value PatternFlowIpv6SegmentRoutingUsidSegmentsLeft) FlowIpv6SegmentRoutingUsid
	// HasSegmentsLeft checks if SegmentsLeft has been set in FlowIpv6SegmentRoutingUsid
	HasSegmentsLeft() bool
	// LastEntry returns PatternFlowIpv6SegmentRoutingUsidLastEntry, set in FlowIpv6SegmentRoutingUsid.
	// PatternFlowIpv6SegmentRoutingUsidLastEntry is zero-based index of the last element in the uSID container list (RFC 8754 Section 2.1). When auto is assigned the value is set to one less than the number of containers specified.
	LastEntry() PatternFlowIpv6SegmentRoutingUsidLastEntry
	// SetLastEntry assigns PatternFlowIpv6SegmentRoutingUsidLastEntry provided by user to FlowIpv6SegmentRoutingUsid.
	// PatternFlowIpv6SegmentRoutingUsidLastEntry is zero-based index of the last element in the uSID container list (RFC 8754 Section 2.1). When auto is assigned the value is set to one less than the number of containers specified.
	SetLastEntry(value PatternFlowIpv6SegmentRoutingUsidLastEntry) FlowIpv6SegmentRoutingUsid
	// HasLastEntry checks if LastEntry has been set in FlowIpv6SegmentRoutingUsid
	HasLastEntry() bool
	// Flags returns FlowIpv6SegmentRoutingFlags, set in FlowIpv6SegmentRoutingUsid.
	// FlowIpv6SegmentRoutingFlags is 8-bit flags field in the SRH (RFC 8754 Section 2.1). RFC 8754 reserves all bits as unused; earlier drafts defined OAM, HMAC, and FRR protection flags.
	Flags() FlowIpv6SegmentRoutingFlags
	// SetFlags assigns FlowIpv6SegmentRoutingFlags provided by user to FlowIpv6SegmentRoutingUsid.
	// FlowIpv6SegmentRoutingFlags is 8-bit flags field in the SRH (RFC 8754 Section 2.1). RFC 8754 reserves all bits as unused; earlier drafts defined OAM, HMAC, and FRR protection flags.
	SetFlags(value FlowIpv6SegmentRoutingFlags) FlowIpv6SegmentRoutingUsid
	// HasFlags checks if Flags has been set in FlowIpv6SegmentRoutingUsid
	HasFlags() bool
	// Tag returns PatternFlowIpv6SegmentRoutingUsidTag, set in FlowIpv6SegmentRoutingUsid.
	// PatternFlowIpv6SegmentRoutingUsidTag is 16-bit field used to tag a packet as part of a class or group (RFC 8754 Section 2.1). If not used it MUST be set to zero.
	Tag() PatternFlowIpv6SegmentRoutingUsidTag
	// SetTag assigns PatternFlowIpv6SegmentRoutingUsidTag provided by user to FlowIpv6SegmentRoutingUsid.
	// PatternFlowIpv6SegmentRoutingUsidTag is 16-bit field used to tag a packet as part of a class or group (RFC 8754 Section 2.1). If not used it MUST be set to zero.
	SetTag(value PatternFlowIpv6SegmentRoutingUsidTag) FlowIpv6SegmentRoutingUsid
	// HasTag checks if Tag has been set in FlowIpv6SegmentRoutingUsid
	HasTag() bool
	// SegmentList returns FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIterIter, set in FlowIpv6SegmentRoutingUsid
	SegmentList() FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter
	setNil()
}

// description is TBD
// SegmentsLeft returns a PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
func (obj *flowIpv6SegmentRoutingUsid) SegmentsLeft() PatternFlowIpv6SegmentRoutingUsidSegmentsLeft {
	if obj.obj.SegmentsLeft == nil {
		obj.obj.SegmentsLeft = NewPatternFlowIpv6SegmentRoutingUsidSegmentsLeft().msg()
	}
	if obj.segmentsLeftHolder == nil {
		obj.segmentsLeftHolder = &patternFlowIpv6SegmentRoutingUsidSegmentsLeft{obj: obj.obj.SegmentsLeft}
	}
	return obj.segmentsLeftHolder
}

// description is TBD
// SegmentsLeft returns a PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
func (obj *flowIpv6SegmentRoutingUsid) HasSegmentsLeft() bool {
	return obj.obj.SegmentsLeft != nil
}

// description is TBD
// SetSegmentsLeft sets the PatternFlowIpv6SegmentRoutingUsidSegmentsLeft value in the FlowIpv6SegmentRoutingUsid object
func (obj *flowIpv6SegmentRoutingUsid) SetSegmentsLeft(value PatternFlowIpv6SegmentRoutingUsidSegmentsLeft) FlowIpv6SegmentRoutingUsid {

	obj.segmentsLeftHolder = nil
	obj.obj.SegmentsLeft = value.msg()

	return obj
}

// description is TBD
// LastEntry returns a PatternFlowIpv6SegmentRoutingUsidLastEntry
func (obj *flowIpv6SegmentRoutingUsid) LastEntry() PatternFlowIpv6SegmentRoutingUsidLastEntry {
	if obj.obj.LastEntry == nil {
		obj.obj.LastEntry = NewPatternFlowIpv6SegmentRoutingUsidLastEntry().msg()
	}
	if obj.lastEntryHolder == nil {
		obj.lastEntryHolder = &patternFlowIpv6SegmentRoutingUsidLastEntry{obj: obj.obj.LastEntry}
	}
	return obj.lastEntryHolder
}

// description is TBD
// LastEntry returns a PatternFlowIpv6SegmentRoutingUsidLastEntry
func (obj *flowIpv6SegmentRoutingUsid) HasLastEntry() bool {
	return obj.obj.LastEntry != nil
}

// description is TBD
// SetLastEntry sets the PatternFlowIpv6SegmentRoutingUsidLastEntry value in the FlowIpv6SegmentRoutingUsid object
func (obj *flowIpv6SegmentRoutingUsid) SetLastEntry(value PatternFlowIpv6SegmentRoutingUsidLastEntry) FlowIpv6SegmentRoutingUsid {

	obj.lastEntryHolder = nil
	obj.obj.LastEntry = value.msg()

	return obj
}

// description is TBD
// Flags returns a FlowIpv6SegmentRoutingFlags
func (obj *flowIpv6SegmentRoutingUsid) Flags() FlowIpv6SegmentRoutingFlags {
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
func (obj *flowIpv6SegmentRoutingUsid) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the FlowIpv6SegmentRoutingFlags value in the FlowIpv6SegmentRoutingUsid object
func (obj *flowIpv6SegmentRoutingUsid) SetFlags(value FlowIpv6SegmentRoutingFlags) FlowIpv6SegmentRoutingUsid {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// description is TBD
// Tag returns a PatternFlowIpv6SegmentRoutingUsidTag
func (obj *flowIpv6SegmentRoutingUsid) Tag() PatternFlowIpv6SegmentRoutingUsidTag {
	if obj.obj.Tag == nil {
		obj.obj.Tag = NewPatternFlowIpv6SegmentRoutingUsidTag().msg()
	}
	if obj.tagHolder == nil {
		obj.tagHolder = &patternFlowIpv6SegmentRoutingUsidTag{obj: obj.obj.Tag}
	}
	return obj.tagHolder
}

// description is TBD
// Tag returns a PatternFlowIpv6SegmentRoutingUsidTag
func (obj *flowIpv6SegmentRoutingUsid) HasTag() bool {
	return obj.obj.Tag != nil
}

// description is TBD
// SetTag sets the PatternFlowIpv6SegmentRoutingUsidTag value in the FlowIpv6SegmentRoutingUsid object
func (obj *flowIpv6SegmentRoutingUsid) SetTag(value PatternFlowIpv6SegmentRoutingUsidTag) FlowIpv6SegmentRoutingUsid {

	obj.tagHolder = nil
	obj.obj.Tag = value.msg()

	return obj
}

// List of pre-computed 128-bit uSID container values (RFC 9800 Section 4), encoded in reverse path order: Segment[0] is the last container (last hop), Segment[n-1] is the first container to visit. Each value is a plain IPv6 address supplied by the user.
// Maximum capacity per SRH (RFC 8754 Section 2.1): The hdr_ext_len field is 8 bits; the SRH total size is (hdr_ext_len + 1) x 8 bytes. With hdr_ext_len max = 255 the SRH can hold at most (255+1)*8 - 8 = 2040 bytes of segment list, giving a theoretical maximum of 127 containers (each 16 bytes). For F3216 (3 uSIDs per container) this is 127 x 3 = 381 micro-SIDs. In practice a standard Ethernet MTU of 1500 bytes limits the SRH to approximately 88 containers (~264 micro-SIDs for F3216) after accounting for the outer IPv6 header (40 bytes) and inner payload.
// SegmentList returns a []FlowIpv6SegmentRoutingUsidSegment
func (obj *flowIpv6SegmentRoutingUsid) SegmentList() FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter {
	if len(obj.obj.SegmentList) == 0 {
		obj.obj.SegmentList = []*otg.FlowIpv6SegmentRoutingUsidSegment{}
	}
	if obj.segmentListHolder == nil {
		obj.segmentListHolder = newFlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter(&obj.obj.SegmentList).setMsg(obj)
	}
	return obj.segmentListHolder
}

type flowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter struct {
	obj                                    *flowIpv6SegmentRoutingUsid
	flowIpv6SegmentRoutingUsidSegmentSlice []FlowIpv6SegmentRoutingUsidSegment
	fieldPtr                               *[]*otg.FlowIpv6SegmentRoutingUsidSegment
}

func newFlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter(ptr *[]*otg.FlowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter {
	return &flowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter{fieldPtr: ptr}
}

type FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter interface {
	setMsg(*flowIpv6SegmentRoutingUsid) FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter
	Items() []FlowIpv6SegmentRoutingUsidSegment
	Add() FlowIpv6SegmentRoutingUsidSegment
	Append(items ...FlowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter
	Set(index int, newObj FlowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter
	Clear() FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter
	clearHolderSlice() FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter
	appendHolderSlice(item FlowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter
}

func (obj *flowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter) setMsg(msg *flowIpv6SegmentRoutingUsid) FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowIpv6SegmentRoutingUsidSegment{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter) Items() []FlowIpv6SegmentRoutingUsidSegment {
	return obj.flowIpv6SegmentRoutingUsidSegmentSlice
}

func (obj *flowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter) Add() FlowIpv6SegmentRoutingUsidSegment {
	newObj := &otg.FlowIpv6SegmentRoutingUsidSegment{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowIpv6SegmentRoutingUsidSegment{obj: newObj}
	newLibObj.setDefault()
	obj.flowIpv6SegmentRoutingUsidSegmentSlice = append(obj.flowIpv6SegmentRoutingUsidSegmentSlice, newLibObj)
	return newLibObj
}

func (obj *flowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter) Append(items ...FlowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowIpv6SegmentRoutingUsidSegmentSlice = append(obj.flowIpv6SegmentRoutingUsidSegmentSlice, item)
	}
	return obj
}

func (obj *flowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter) Set(index int, newObj FlowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowIpv6SegmentRoutingUsidSegmentSlice[index] = newObj
	return obj
}
func (obj *flowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter) Clear() FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowIpv6SegmentRoutingUsidSegment{}
		obj.flowIpv6SegmentRoutingUsidSegmentSlice = []FlowIpv6SegmentRoutingUsidSegment{}
	}
	return obj
}
func (obj *flowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter) clearHolderSlice() FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter {
	if len(obj.flowIpv6SegmentRoutingUsidSegmentSlice) > 0 {
		obj.flowIpv6SegmentRoutingUsidSegmentSlice = []FlowIpv6SegmentRoutingUsidSegment{}
	}
	return obj
}
func (obj *flowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter) appendHolderSlice(item FlowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidFlowIpv6SegmentRoutingUsidSegmentIter {
	obj.flowIpv6SegmentRoutingUsidSegmentSlice = append(obj.flowIpv6SegmentRoutingUsidSegmentSlice, item)
	return obj
}

func (obj *flowIpv6SegmentRoutingUsid) validateObj(vObj *validation, set_default bool) {
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
				obj.SegmentList().appendHolderSlice(&flowIpv6SegmentRoutingUsidSegment{obj: item})
			}
		}
		for _, item := range obj.SegmentList().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowIpv6SegmentRoutingUsid) setDefault() {

}
