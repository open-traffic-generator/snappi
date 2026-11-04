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
	obj                *otg.FlowIpv6SegmentRouting
	marshaller         marshalFlowIpv6SegmentRouting
	unMarshaller       unMarshalFlowIpv6SegmentRouting
	segmentsLeftHolder PatternFlowIpv6SegmentRoutingSegmentsLeft
	lastEntryHolder    PatternFlowIpv6SegmentRoutingLastEntry
	flagsHolder        FlowIpv6SegmentRoutingFlags
	tagHolder          PatternFlowIpv6SegmentRoutingTag
	segmentListHolder  FlowIpv6SegmentRoutingFlowIpv6SegmentRoutingSegmentIter
	padTlvHolder       FlowIpv6SRHPadTlv
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
	obj.padTlvHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRouting is iPv6 Segment Routing Header (SRH, Routing Type 4, RFC 8754 Section 2) carrying full 128-bit SRv6 SIDs. Each segment list entry is a complete SID (locator + function + argument) as a 128-bit IPv6 address. Segment list encoded in reverse path order: Segment[0] is the last hop, Segment[n-1] is the first hop.
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
	// FlowIpv6SegmentRoutingFlags is an 8-bit field containing flags. While RFC 8754 reserves all bits as unused, earlier drafts defined specific flags for behavior such as OAM, HMAC, and FRR protection.
	Flags() FlowIpv6SegmentRoutingFlags
	// SetFlags assigns FlowIpv6SegmentRoutingFlags provided by user to FlowIpv6SegmentRouting.
	// FlowIpv6SegmentRoutingFlags is an 8-bit field containing flags. While RFC 8754 reserves all bits as unused, earlier drafts defined specific flags for behavior such as OAM, HMAC, and FRR protection.
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
	// PadTlv returns FlowIpv6SRHPadTlv, set in FlowIpv6SegmentRouting.
	// FlowIpv6SRHPadTlv is sRH Padding TLV (type 4, RFC 8754 Section 2.1). Used to align the SRH TLV block to an 8-byte boundary. The padding bytes are set to zero and are skipped by transit nodes. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 8754 Section 2.1.
	PadTlv() FlowIpv6SRHPadTlv
	// SetPadTlv assigns FlowIpv6SRHPadTlv provided by user to FlowIpv6SegmentRouting.
	// FlowIpv6SRHPadTlv is sRH Padding TLV (type 4, RFC 8754 Section 2.1). Used to align the SRH TLV block to an 8-byte boundary. The padding bytes are set to zero and are skipped by transit nodes. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 8754 Section 2.1.
	SetPadTlv(value FlowIpv6SRHPadTlv) FlowIpv6SegmentRouting
	// HasPadTlv checks if PadTlv has been set in FlowIpv6SegmentRouting
	HasPadTlv() bool
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

	if obj.obj.PadTlv != nil {

		obj.PadTlv().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6SegmentRouting) setDefault() {

}
