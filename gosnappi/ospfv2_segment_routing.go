package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2SegmentRouting *****
type ospfv2SegmentRouting struct {
	validation
	obj                 *otg.Ospfv2SegmentRouting
	marshaller          marshalOspfv2SegmentRouting
	unMarshaller        unMarshalOspfv2SegmentRouting
	srgbRangesHolder    Ospfv2SegmentRoutingOspfv2SRSrgbIter
	srlbRangesHolder    Ospfv2SegmentRoutingOspfv2SRSrlbIter
	nodePrefixSidHolder Ospfv2SRRouterNodeSid
}

func NewOspfv2SegmentRouting() Ospfv2SegmentRouting {
	obj := ospfv2SegmentRouting{obj: &otg.Ospfv2SegmentRouting{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2SegmentRouting) msg() *otg.Ospfv2SegmentRouting {
	return obj.obj
}

func (obj *ospfv2SegmentRouting) setMsg(msg *otg.Ospfv2SegmentRouting) Ospfv2SegmentRouting {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2SegmentRouting struct {
	obj *ospfv2SegmentRouting
}

type marshalOspfv2SegmentRouting interface {
	// ToProto marshals Ospfv2SegmentRouting to protobuf object *otg.Ospfv2SegmentRouting
	ToProto() (*otg.Ospfv2SegmentRouting, error)
	// ToPbText marshals Ospfv2SegmentRouting to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2SegmentRouting to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2SegmentRouting to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2SegmentRouting struct {
	obj *ospfv2SegmentRouting
}

type unMarshalOspfv2SegmentRouting interface {
	// FromProto unmarshals Ospfv2SegmentRouting from protobuf object *otg.Ospfv2SegmentRouting
	FromProto(msg *otg.Ospfv2SegmentRouting) (Ospfv2SegmentRouting, error)
	// FromPbText unmarshals Ospfv2SegmentRouting from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2SegmentRouting from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2SegmentRouting from JSON text
	FromJson(value string) error
}

func (obj *ospfv2SegmentRouting) Marshal() marshalOspfv2SegmentRouting {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2SegmentRouting{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2SegmentRouting) Unmarshal() unMarshalOspfv2SegmentRouting {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2SegmentRouting{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2SegmentRouting) ToProto() (*otg.Ospfv2SegmentRouting, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2SegmentRouting) FromProto(msg *otg.Ospfv2SegmentRouting) (Ospfv2SegmentRouting, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2SegmentRouting) ToPbText() (string, error) {
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

func (m *unMarshalospfv2SegmentRouting) FromPbText(value string) error {
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

func (m *marshalospfv2SegmentRouting) ToYaml() (string, error) {
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

func (m *unMarshalospfv2SegmentRouting) FromYaml(value string) error {
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

func (m *marshalospfv2SegmentRouting) ToJson() (string, error) {
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

func (m *unMarshalospfv2SegmentRouting) FromJson(value string) error {
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

func (obj *ospfv2SegmentRouting) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2SegmentRouting) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2SegmentRouting) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2SegmentRouting) Clone() (Ospfv2SegmentRouting, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2SegmentRouting()
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

func (obj *ospfv2SegmentRouting) setNil() {
	obj.srgbRangesHolder = nil
	obj.srlbRangesHolder = nil
	obj.nodePrefixSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2SegmentRouting is segment Routing (SR) allows for a flexible definition of end-to-end paths within IGP
// topologies by encoding paths as sequences of topological sub-paths, called "segments".
// In OSPFv2 the SR-specific information is advertised using Opaque LSAs. The router level
// Segment Routing capabilities (SR-Algorithm, SID/Label Range (SRGB), SR Local Block (SRLB))
// are carried in the Router Information (RI) Opaque LSA, and the router's own Node Prefix-SID
// is carried in the Extended Prefix Opaque LSA for the router loopback.
// Reference: https://datatracker.ietf.org/doc/html/rfc8665.
// An implementation may advertise the SR capabilities with default values if a user does
// not set any of the properties of Segment Routing.
type Ospfv2SegmentRouting interface {
	Validation
	// msg marshals Ospfv2SegmentRouting to protobuf object *otg.Ospfv2SegmentRouting
	// and doesn't set defaults
	msg() *otg.Ospfv2SegmentRouting
	// setMsg unmarshals Ospfv2SegmentRouting from protobuf object *otg.Ospfv2SegmentRouting
	// and doesn't set defaults
	setMsg(*otg.Ospfv2SegmentRouting) Ospfv2SegmentRouting
	// provides marshal interface
	Marshal() marshalOspfv2SegmentRouting
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2SegmentRouting
	// validate validates Ospfv2SegmentRouting
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2SegmentRouting, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Algorithms returns []uint32, set in Ospfv2SegmentRouting.
	Algorithms() []uint32
	// SetAlgorithms assigns []uint32 provided by user to Ospfv2SegmentRouting
	SetAlgorithms(value []uint32) Ospfv2SegmentRouting
	// SrgbRanges returns Ospfv2SegmentRoutingOspfv2SRSrgbIterIter, set in Ospfv2SegmentRouting
	SrgbRanges() Ospfv2SegmentRoutingOspfv2SRSrgbIter
	// SrlbRanges returns Ospfv2SegmentRoutingOspfv2SRSrlbIterIter, set in Ospfv2SegmentRouting
	SrlbRanges() Ospfv2SegmentRoutingOspfv2SRSrlbIter
	// NodePrefixSid returns Ospfv2SRRouterNodeSid, set in Ospfv2SegmentRouting.
	// Ospfv2SRRouterNodeSid is the Node (loopback) Prefix-SID advertised by this router for its own loopback address.
	// It is advertised as a Prefix-SID sub-TLV inside the Extended Prefix TLV of the Extended
	// Prefix Opaque LSA, together with the one-octet Extended Prefix flags.
	// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-prefix-sid-sub-tlv.
	NodePrefixSid() Ospfv2SRRouterNodeSid
	// SetNodePrefixSid assigns Ospfv2SRRouterNodeSid provided by user to Ospfv2SegmentRouting.
	// Ospfv2SRRouterNodeSid is the Node (loopback) Prefix-SID advertised by this router for its own loopback address.
	// It is advertised as a Prefix-SID sub-TLV inside the Extended Prefix TLV of the Extended
	// Prefix Opaque LSA, together with the one-octet Extended Prefix flags.
	// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-prefix-sid-sub-tlv.
	SetNodePrefixSid(value Ospfv2SRRouterNodeSid) Ospfv2SegmentRouting
	// HasNodePrefixSid checks if NodePrefixSid has been set in Ospfv2SegmentRouting
	HasNodePrefixSid() bool
	setNil()
}

// The SR-Algorithm TLV, if present, carries one or more Segment Routing Algorithms that
// the router supports when calculating reachability to other nodes or to prefixes
// attached to these nodes.
// - 0: SPF algorithm based on link metric.
// - 1: Strict SPF algorithm based on link metric.
// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-sr-algorithm-tlv.
// When the originating router does not advertise the SR-Algorithm TLV, it implies that
// algorithm 0 is the only algorithm supported. When advertised, algorithm 0 MUST be
// present while non-zero algorithms MAY be present.
// Algorithms returns a []uint32
func (obj *ospfv2SegmentRouting) Algorithms() []uint32 {
	if obj.obj.Algorithms == nil {
		obj.obj.Algorithms = make([]uint32, 0)
	}
	return obj.obj.Algorithms
}

// The SR-Algorithm TLV, if present, carries one or more Segment Routing Algorithms that
// the router supports when calculating reachability to other nodes or to prefixes
// attached to these nodes.
// - 0: SPF algorithm based on link metric.
// - 1: Strict SPF algorithm based on link metric.
// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-sr-algorithm-tlv.
// When the originating router does not advertise the SR-Algorithm TLV, it implies that
// algorithm 0 is the only algorithm supported. When advertised, algorithm 0 MUST be
// present while non-zero algorithms MAY be present.
// SetAlgorithms sets the []uint32 value in the Ospfv2SegmentRouting object
func (obj *ospfv2SegmentRouting) SetAlgorithms(value []uint32) Ospfv2SegmentRouting {

	if obj.obj.Algorithms == nil {
		obj.obj.Algorithms = make([]uint32, 0)
	}
	obj.obj.Algorithms = value

	return obj
}

// The list of Segment Routing Global Block (SRGB) ranges advertised in the
// SID/Label Range TLV of the RI Opaque LSA.
// If no SRGB range is configured, an implementation should advertise one SRGB range with
// default values.
// SrgbRanges returns a []Ospfv2SRSrgb
func (obj *ospfv2SegmentRouting) SrgbRanges() Ospfv2SegmentRoutingOspfv2SRSrgbIter {
	if len(obj.obj.SrgbRanges) == 0 {
		obj.obj.SrgbRanges = []*otg.Ospfv2SRSrgb{}
	}
	if obj.srgbRangesHolder == nil {
		obj.srgbRangesHolder = newOspfv2SegmentRoutingOspfv2SRSrgbIter(&obj.obj.SrgbRanges).setMsg(obj)
	}
	return obj.srgbRangesHolder
}

type ospfv2SegmentRoutingOspfv2SRSrgbIter struct {
	obj               *ospfv2SegmentRouting
	ospfv2SRSrgbSlice []Ospfv2SRSrgb
	fieldPtr          *[]*otg.Ospfv2SRSrgb
}

func newOspfv2SegmentRoutingOspfv2SRSrgbIter(ptr *[]*otg.Ospfv2SRSrgb) Ospfv2SegmentRoutingOspfv2SRSrgbIter {
	return &ospfv2SegmentRoutingOspfv2SRSrgbIter{fieldPtr: ptr}
}

type Ospfv2SegmentRoutingOspfv2SRSrgbIter interface {
	setMsg(*ospfv2SegmentRouting) Ospfv2SegmentRoutingOspfv2SRSrgbIter
	Items() []Ospfv2SRSrgb
	Add() Ospfv2SRSrgb
	Append(items ...Ospfv2SRSrgb) Ospfv2SegmentRoutingOspfv2SRSrgbIter
	Set(index int, newObj Ospfv2SRSrgb) Ospfv2SegmentRoutingOspfv2SRSrgbIter
	Clear() Ospfv2SegmentRoutingOspfv2SRSrgbIter
	clearHolderSlice() Ospfv2SegmentRoutingOspfv2SRSrgbIter
	appendHolderSlice(item Ospfv2SRSrgb) Ospfv2SegmentRoutingOspfv2SRSrgbIter
}

func (obj *ospfv2SegmentRoutingOspfv2SRSrgbIter) setMsg(msg *ospfv2SegmentRouting) Ospfv2SegmentRoutingOspfv2SRSrgbIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2SRSrgb{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2SegmentRoutingOspfv2SRSrgbIter) Items() []Ospfv2SRSrgb {
	return obj.ospfv2SRSrgbSlice
}

func (obj *ospfv2SegmentRoutingOspfv2SRSrgbIter) Add() Ospfv2SRSrgb {
	newObj := &otg.Ospfv2SRSrgb{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2SRSrgb{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2SRSrgbSlice = append(obj.ospfv2SRSrgbSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2SegmentRoutingOspfv2SRSrgbIter) Append(items ...Ospfv2SRSrgb) Ospfv2SegmentRoutingOspfv2SRSrgbIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2SRSrgbSlice = append(obj.ospfv2SRSrgbSlice, item)
	}
	return obj
}

func (obj *ospfv2SegmentRoutingOspfv2SRSrgbIter) Set(index int, newObj Ospfv2SRSrgb) Ospfv2SegmentRoutingOspfv2SRSrgbIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2SRSrgbSlice[index] = newObj
	return obj
}
func (obj *ospfv2SegmentRoutingOspfv2SRSrgbIter) Clear() Ospfv2SegmentRoutingOspfv2SRSrgbIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2SRSrgb{}
		obj.ospfv2SRSrgbSlice = []Ospfv2SRSrgb{}
	}
	return obj
}
func (obj *ospfv2SegmentRoutingOspfv2SRSrgbIter) clearHolderSlice() Ospfv2SegmentRoutingOspfv2SRSrgbIter {
	if len(obj.ospfv2SRSrgbSlice) > 0 {
		obj.ospfv2SRSrgbSlice = []Ospfv2SRSrgb{}
	}
	return obj
}
func (obj *ospfv2SegmentRoutingOspfv2SRSrgbIter) appendHolderSlice(item Ospfv2SRSrgb) Ospfv2SegmentRoutingOspfv2SRSrgbIter {
	obj.ospfv2SRSrgbSlice = append(obj.ospfv2SRSrgbSlice, item)
	return obj
}

// The list of SR Local Block (SRLB) ranges advertised in the SR Local Block TLV of the
// RI Opaque LSA. The SRLB contains the range of labels the node has reserved for local
// SIDs, for example Adjacency SIDs.
// SrlbRanges returns a []Ospfv2SRSrlb
func (obj *ospfv2SegmentRouting) SrlbRanges() Ospfv2SegmentRoutingOspfv2SRSrlbIter {
	if len(obj.obj.SrlbRanges) == 0 {
		obj.obj.SrlbRanges = []*otg.Ospfv2SRSrlb{}
	}
	if obj.srlbRangesHolder == nil {
		obj.srlbRangesHolder = newOspfv2SegmentRoutingOspfv2SRSrlbIter(&obj.obj.SrlbRanges).setMsg(obj)
	}
	return obj.srlbRangesHolder
}

type ospfv2SegmentRoutingOspfv2SRSrlbIter struct {
	obj               *ospfv2SegmentRouting
	ospfv2SRSrlbSlice []Ospfv2SRSrlb
	fieldPtr          *[]*otg.Ospfv2SRSrlb
}

func newOspfv2SegmentRoutingOspfv2SRSrlbIter(ptr *[]*otg.Ospfv2SRSrlb) Ospfv2SegmentRoutingOspfv2SRSrlbIter {
	return &ospfv2SegmentRoutingOspfv2SRSrlbIter{fieldPtr: ptr}
}

type Ospfv2SegmentRoutingOspfv2SRSrlbIter interface {
	setMsg(*ospfv2SegmentRouting) Ospfv2SegmentRoutingOspfv2SRSrlbIter
	Items() []Ospfv2SRSrlb
	Add() Ospfv2SRSrlb
	Append(items ...Ospfv2SRSrlb) Ospfv2SegmentRoutingOspfv2SRSrlbIter
	Set(index int, newObj Ospfv2SRSrlb) Ospfv2SegmentRoutingOspfv2SRSrlbIter
	Clear() Ospfv2SegmentRoutingOspfv2SRSrlbIter
	clearHolderSlice() Ospfv2SegmentRoutingOspfv2SRSrlbIter
	appendHolderSlice(item Ospfv2SRSrlb) Ospfv2SegmentRoutingOspfv2SRSrlbIter
}

func (obj *ospfv2SegmentRoutingOspfv2SRSrlbIter) setMsg(msg *ospfv2SegmentRouting) Ospfv2SegmentRoutingOspfv2SRSrlbIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2SRSrlb{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2SegmentRoutingOspfv2SRSrlbIter) Items() []Ospfv2SRSrlb {
	return obj.ospfv2SRSrlbSlice
}

func (obj *ospfv2SegmentRoutingOspfv2SRSrlbIter) Add() Ospfv2SRSrlb {
	newObj := &otg.Ospfv2SRSrlb{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2SRSrlb{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2SRSrlbSlice = append(obj.ospfv2SRSrlbSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2SegmentRoutingOspfv2SRSrlbIter) Append(items ...Ospfv2SRSrlb) Ospfv2SegmentRoutingOspfv2SRSrlbIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2SRSrlbSlice = append(obj.ospfv2SRSrlbSlice, item)
	}
	return obj
}

func (obj *ospfv2SegmentRoutingOspfv2SRSrlbIter) Set(index int, newObj Ospfv2SRSrlb) Ospfv2SegmentRoutingOspfv2SRSrlbIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2SRSrlbSlice[index] = newObj
	return obj
}
func (obj *ospfv2SegmentRoutingOspfv2SRSrlbIter) Clear() Ospfv2SegmentRoutingOspfv2SRSrlbIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2SRSrlb{}
		obj.ospfv2SRSrlbSlice = []Ospfv2SRSrlb{}
	}
	return obj
}
func (obj *ospfv2SegmentRoutingOspfv2SRSrlbIter) clearHolderSlice() Ospfv2SegmentRoutingOspfv2SRSrlbIter {
	if len(obj.ospfv2SRSrlbSlice) > 0 {
		obj.ospfv2SRSrlbSlice = []Ospfv2SRSrlb{}
	}
	return obj
}
func (obj *ospfv2SegmentRoutingOspfv2SRSrlbIter) appendHolderSlice(item Ospfv2SRSrlb) Ospfv2SegmentRoutingOspfv2SRSrlbIter {
	obj.ospfv2SRSrlbSlice = append(obj.ospfv2SRSrlbSlice, item)
	return obj
}

// Optional Node (loopback) Prefix-SID advertised by this router in the Extended Prefix
// Opaque LSA for its own loopback address. The Node Prefix-SID identifies the router in
// the Segment Routing domain.
// NodePrefixSid returns a Ospfv2SRRouterNodeSid
func (obj *ospfv2SegmentRouting) NodePrefixSid() Ospfv2SRRouterNodeSid {
	if obj.obj.NodePrefixSid == nil {
		obj.obj.NodePrefixSid = NewOspfv2SRRouterNodeSid().msg()
	}
	if obj.nodePrefixSidHolder == nil {
		obj.nodePrefixSidHolder = &ospfv2SRRouterNodeSid{obj: obj.obj.NodePrefixSid}
	}
	return obj.nodePrefixSidHolder
}

// Optional Node (loopback) Prefix-SID advertised by this router in the Extended Prefix
// Opaque LSA for its own loopback address. The Node Prefix-SID identifies the router in
// the Segment Routing domain.
// NodePrefixSid returns a Ospfv2SRRouterNodeSid
func (obj *ospfv2SegmentRouting) HasNodePrefixSid() bool {
	return obj.obj.NodePrefixSid != nil
}

// Optional Node (loopback) Prefix-SID advertised by this router in the Extended Prefix
// Opaque LSA for its own loopback address. The Node Prefix-SID identifies the router in
// the Segment Routing domain.
// SetNodePrefixSid sets the Ospfv2SRRouterNodeSid value in the Ospfv2SegmentRouting object
func (obj *ospfv2SegmentRouting) SetNodePrefixSid(value Ospfv2SRRouterNodeSid) Ospfv2SegmentRouting {

	obj.nodePrefixSidHolder = nil
	obj.obj.NodePrefixSid = value.msg()

	return obj
}

func (obj *ospfv2SegmentRouting) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Algorithms != nil {

		for _, item := range obj.obj.Algorithms {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("0 <= Ospfv2SegmentRouting.Algorithms <= 255 but Got %d", item))
			}

		}

	}

	if len(obj.obj.SrgbRanges) != 0 {

		if set_default {
			obj.SrgbRanges().clearHolderSlice()
			for _, item := range obj.obj.SrgbRanges {
				obj.SrgbRanges().appendHolderSlice(&ospfv2SRSrgb{obj: item})
			}
		}
		for _, item := range obj.SrgbRanges().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.SrlbRanges) != 0 {

		if set_default {
			obj.SrlbRanges().clearHolderSlice()
			for _, item := range obj.obj.SrlbRanges {
				obj.SrlbRanges().appendHolderSlice(&ospfv2SRSrlb{obj: item})
			}
		}
		for _, item := range obj.SrlbRanges().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.NodePrefixSid != nil {

		obj.NodePrefixSid().validateObj(vObj, set_default)
	}

}

func (obj *ospfv2SegmentRouting) setDefault() {

}
