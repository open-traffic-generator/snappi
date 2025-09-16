package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV6Peer *****
type bgpV6Peer struct {
	validation
	obj                            *otg.BgpV6Peer
	marshaller                     marshalBgpV6Peer
	unMarshaller                   unMarshalBgpV6Peer
	segmentRoutingHolder           BgpV6SegmentRouting
	evpnEthernetSegmentsHolder     BgpV6PeerBgpV6EthernetSegmentIter
	advancedHolder                 BgpAdvanced
	capabilityHolder               BgpCapability
	learnedInformationFilterHolder BgpLearnedInformationFilter
	v4RoutesHolder                 BgpV6PeerBgpV4RouteRangeIter
	v6RoutesHolder                 BgpV6PeerBgpV6RouteRangeIter
	v4SrtePoliciesHolder           BgpV6PeerBgpSrteV4PolicyIter
	v6SrtePoliciesHolder           BgpV6PeerBgpSrteV6PolicyIter
	gracefulRestartHolder          BgpGracefulRestart
	replayUpdatesHolder            BgpUpdateReplay
}

func NewBgpV6Peer() BgpV6Peer {
	obj := bgpV6Peer{obj: &otg.BgpV6Peer{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV6Peer) msg() *otg.BgpV6Peer {
	return obj.obj
}

func (obj *bgpV6Peer) setMsg(msg *otg.BgpV6Peer) BgpV6Peer {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV6Peer struct {
	obj *bgpV6Peer
}

type marshalBgpV6Peer interface {
	// ToProto marshals BgpV6Peer to protobuf object *otg.BgpV6Peer
	ToProto() (*otg.BgpV6Peer, error)
	// ToPbText marshals BgpV6Peer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV6Peer to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV6Peer to JSON text
	ToJson() (string, error)
}

type unMarshalbgpV6Peer struct {
	obj *bgpV6Peer
}

type unMarshalBgpV6Peer interface {
	// FromProto unmarshals BgpV6Peer from protobuf object *otg.BgpV6Peer
	FromProto(msg *otg.BgpV6Peer) (BgpV6Peer, error)
	// FromPbText unmarshals BgpV6Peer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV6Peer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV6Peer from JSON text
	FromJson(value string) error
}

func (obj *bgpV6Peer) Marshal() marshalBgpV6Peer {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV6Peer{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV6Peer) Unmarshal() unMarshalBgpV6Peer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV6Peer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV6Peer) ToProto() (*otg.BgpV6Peer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV6Peer) FromProto(msg *otg.BgpV6Peer) (BgpV6Peer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV6Peer) ToPbText() (string, error) {
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

func (m *unMarshalbgpV6Peer) FromPbText(value string) error {
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

func (m *marshalbgpV6Peer) ToYaml() (string, error) {
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

func (m *unMarshalbgpV6Peer) FromYaml(value string) error {
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

func (m *marshalbgpV6Peer) ToJson() (string, error) {
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

func (m *unMarshalbgpV6Peer) FromJson(value string) error {
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

func (obj *bgpV6Peer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV6Peer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV6Peer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV6Peer) Clone() (BgpV6Peer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV6Peer()
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

func (obj *bgpV6Peer) setNil() {
	obj.segmentRoutingHolder = nil
	obj.evpnEthernetSegmentsHolder = nil
	obj.advancedHolder = nil
	obj.capabilityHolder = nil
	obj.learnedInformationFilterHolder = nil
	obj.v4RoutesHolder = nil
	obj.v6RoutesHolder = nil
	obj.v4SrtePoliciesHolder = nil
	obj.v6SrtePoliciesHolder = nil
	obj.gracefulRestartHolder = nil
	obj.replayUpdatesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpV6Peer is configuration for BGPv6 peer settings and routes.
type BgpV6Peer interface {
	Validation
	// msg marshals BgpV6Peer to protobuf object *otg.BgpV6Peer
	// and doesn't set defaults
	msg() *otg.BgpV6Peer
	// setMsg unmarshals BgpV6Peer from protobuf object *otg.BgpV6Peer
	// and doesn't set defaults
	setMsg(*otg.BgpV6Peer) BgpV6Peer
	// provides marshal interface
	Marshal() marshalBgpV6Peer
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV6Peer
	// validate validates BgpV6Peer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV6Peer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PeerAddress returns string, set in BgpV6Peer.
	PeerAddress() string
	// SetPeerAddress assigns string provided by user to BgpV6Peer
	SetPeerAddress(value string) BgpV6Peer
	// SegmentRouting returns BgpV6SegmentRouting, set in BgpV6Peer.
	// BgpV6SegmentRouting is configuration for BGPv6 segment routing settings.
	SegmentRouting() BgpV6SegmentRouting
	// SetSegmentRouting assigns BgpV6SegmentRouting provided by user to BgpV6Peer.
	// BgpV6SegmentRouting is configuration for BGPv6 segment routing settings.
	SetSegmentRouting(value BgpV6SegmentRouting) BgpV6Peer
	// HasSegmentRouting checks if SegmentRouting has been set in BgpV6Peer
	HasSegmentRouting() bool
	// EvpnEthernetSegments returns BgpV6PeerBgpV6EthernetSegmentIterIter, set in BgpV6Peer
	EvpnEthernetSegments() BgpV6PeerBgpV6EthernetSegmentIter
	// AsType returns BgpV6PeerAsTypeEnum, set in BgpV6Peer
	AsType() BgpV6PeerAsTypeEnum
	// SetAsType assigns BgpV6PeerAsTypeEnum provided by user to BgpV6Peer
	SetAsType(value BgpV6PeerAsTypeEnum) BgpV6Peer
	// AsNumber returns uint32, set in BgpV6Peer.
	AsNumber() uint32
	// SetAsNumber assigns uint32 provided by user to BgpV6Peer
	SetAsNumber(value uint32) BgpV6Peer
	// AsNumberWidth returns BgpV6PeerAsNumberWidthEnum, set in BgpV6Peer
	AsNumberWidth() BgpV6PeerAsNumberWidthEnum
	// SetAsNumberWidth assigns BgpV6PeerAsNumberWidthEnum provided by user to BgpV6Peer
	SetAsNumberWidth(value BgpV6PeerAsNumberWidthEnum) BgpV6Peer
	// HasAsNumberWidth checks if AsNumberWidth has been set in BgpV6Peer
	HasAsNumberWidth() bool
	// Advanced returns BgpAdvanced, set in BgpV6Peer.
	// BgpAdvanced is configuration for BGP advanced settings.
	Advanced() BgpAdvanced
	// SetAdvanced assigns BgpAdvanced provided by user to BgpV6Peer.
	// BgpAdvanced is configuration for BGP advanced settings.
	SetAdvanced(value BgpAdvanced) BgpV6Peer
	// HasAdvanced checks if Advanced has been set in BgpV6Peer
	HasAdvanced() bool
	// Capability returns BgpCapability, set in BgpV6Peer.
	// BgpCapability is configuration for BGP capability settings.
	Capability() BgpCapability
	// SetCapability assigns BgpCapability provided by user to BgpV6Peer.
	// BgpCapability is configuration for BGP capability settings.
	SetCapability(value BgpCapability) BgpV6Peer
	// HasCapability checks if Capability has been set in BgpV6Peer
	HasCapability() bool
	// LearnedInformationFilter returns BgpLearnedInformationFilter, set in BgpV6Peer.
	// BgpLearnedInformationFilter is configuration for controlling storage of BGP learned information recieved from the peer.
	LearnedInformationFilter() BgpLearnedInformationFilter
	// SetLearnedInformationFilter assigns BgpLearnedInformationFilter provided by user to BgpV6Peer.
	// BgpLearnedInformationFilter is configuration for controlling storage of BGP learned information recieved from the peer.
	SetLearnedInformationFilter(value BgpLearnedInformationFilter) BgpV6Peer
	// HasLearnedInformationFilter checks if LearnedInformationFilter has been set in BgpV6Peer
	HasLearnedInformationFilter() bool
	// V4Routes returns BgpV6PeerBgpV4RouteRangeIterIter, set in BgpV6Peer
	V4Routes() BgpV6PeerBgpV4RouteRangeIter
	// V6Routes returns BgpV6PeerBgpV6RouteRangeIterIter, set in BgpV6Peer
	V6Routes() BgpV6PeerBgpV6RouteRangeIter
	// V4SrtePolicies returns BgpV6PeerBgpSrteV4PolicyIterIter, set in BgpV6Peer
	V4SrtePolicies() BgpV6PeerBgpSrteV4PolicyIter
	// V6SrtePolicies returns BgpV6PeerBgpSrteV6PolicyIterIter, set in BgpV6Peer
	V6SrtePolicies() BgpV6PeerBgpSrteV6PolicyIter
	// Name returns string, set in BgpV6Peer.
	Name() string
	// SetName assigns string provided by user to BgpV6Peer
	SetName(value string) BgpV6Peer
	// GracefulRestart returns BgpGracefulRestart, set in BgpV6Peer.
	// BgpGracefulRestart is the Graceful Restart Capability (RFC 4724) is a BGP capability that can be used by a BGP speaker to indicate its ability to preserve its forwarding state during BGP restart. The Graceful Restart (GR) capability is advertised in OPEN messages sent between BGP peers. After a BGP session has been established, and the initial routing update has been completed,  an End-of-RIB (Routing Information Base) marker is sent in an UPDATE message to convey information  about routing convergence.
	GracefulRestart() BgpGracefulRestart
	// SetGracefulRestart assigns BgpGracefulRestart provided by user to BgpV6Peer.
	// BgpGracefulRestart is the Graceful Restart Capability (RFC 4724) is a BGP capability that can be used by a BGP speaker to indicate its ability to preserve its forwarding state during BGP restart. The Graceful Restart (GR) capability is advertised in OPEN messages sent between BGP peers. After a BGP session has been established, and the initial routing update has been completed,  an End-of-RIB (Routing Information Base) marker is sent in an UPDATE message to convey information  about routing convergence.
	SetGracefulRestart(value BgpGracefulRestart) BgpV6Peer
	// HasGracefulRestart checks if GracefulRestart has been set in BgpV6Peer
	HasGracefulRestart() bool
	// ReplayUpdates returns BgpUpdateReplay, set in BgpV6Peer.
	// BgpUpdateReplay is ordered BGP Updates ( including both Advertise and Withdraws ) to be sent in the order given in the input to the peer after the BGP session is established.
	ReplayUpdates() BgpUpdateReplay
	// SetReplayUpdates assigns BgpUpdateReplay provided by user to BgpV6Peer.
	// BgpUpdateReplay is ordered BGP Updates ( including both Advertise and Withdraws ) to be sent in the order given in the input to the peer after the BGP session is established.
	SetReplayUpdates(value BgpUpdateReplay) BgpV6Peer
	// HasReplayUpdates checks if ReplayUpdates has been set in BgpV6Peer
	HasReplayUpdates() bool
	setNil()
}

// IPv6 address of the BGP peer for the session
// PeerAddress returns a string
func (obj *bgpV6Peer) PeerAddress() string {

	return *obj.obj.PeerAddress

}

// IPv6 address of the BGP peer for the session
// SetPeerAddress sets the string value in the BgpV6Peer object
func (obj *bgpV6Peer) SetPeerAddress(value string) BgpV6Peer {

	obj.obj.PeerAddress = &value
	return obj
}

// description is TBD
// SegmentRouting returns a BgpV6SegmentRouting
func (obj *bgpV6Peer) SegmentRouting() BgpV6SegmentRouting {
	if obj.obj.SegmentRouting == nil {
		obj.obj.SegmentRouting = NewBgpV6SegmentRouting().msg()
	}
	if obj.segmentRoutingHolder == nil {
		obj.segmentRoutingHolder = &bgpV6SegmentRouting{obj: obj.obj.SegmentRouting}
	}
	return obj.segmentRoutingHolder
}

// description is TBD
// SegmentRouting returns a BgpV6SegmentRouting
func (obj *bgpV6Peer) HasSegmentRouting() bool {
	return obj.obj.SegmentRouting != nil
}

// description is TBD
// SetSegmentRouting sets the BgpV6SegmentRouting value in the BgpV6Peer object
func (obj *bgpV6Peer) SetSegmentRouting(value BgpV6SegmentRouting) BgpV6Peer {

	obj.segmentRoutingHolder = nil
	obj.obj.SegmentRouting = value.msg()

	return obj
}

// This contains the list of Ethernet Virtual Private Network (EVPN) Ethernet Segments (ES) Per BGP Peer for IPv6 Address Family Identifier (AFI).
//
// Each Ethernet Segment contains a list of EVPN Instances (EVIs) .
// Each EVI contains a list of Broadcast Domains.
// Each Broadcast Domain contains a list of MAC/IP Ranges.
//
// <Ethernet Segment, EVI, Broadcast Domain> is responsible for advertising Ethernet Auto-discovery Route Per EVI (Type 1).
//
// <Ethernet Segment, EVI> is responsible for advertising Ethernet Auto-discovery Route Per Ethernet Segment (Type 1).
//
// <Ethernet Segment, EVI, Broadcast Domain, MAC/IP> is responsible for advertising MAC/IP Advertisement Route (Type 2).
//
// <Ethernet Segment, EVI, Broadcast Domain> is responsible for advertising Inclusive Multicast Ethernet Tag Route (Type 3).
//
// Ethernet Segment is responsible for advertising Ethernet Segment Route (Type 4).
// EvpnEthernetSegments returns a []BgpV6EthernetSegment
func (obj *bgpV6Peer) EvpnEthernetSegments() BgpV6PeerBgpV6EthernetSegmentIter {
	if len(obj.obj.EvpnEthernetSegments) == 0 {
		obj.obj.EvpnEthernetSegments = []*otg.BgpV6EthernetSegment{}
	}
	if obj.evpnEthernetSegmentsHolder == nil {
		obj.evpnEthernetSegmentsHolder = newBgpV6PeerBgpV6EthernetSegmentIter(&obj.obj.EvpnEthernetSegments).setMsg(obj)
	}
	return obj.evpnEthernetSegmentsHolder
}

type bgpV6PeerBgpV6EthernetSegmentIter struct {
	obj                       *bgpV6Peer
	bgpV6EthernetSegmentSlice []BgpV6EthernetSegment
	fieldPtr                  *[]*otg.BgpV6EthernetSegment
}

func newBgpV6PeerBgpV6EthernetSegmentIter(ptr *[]*otg.BgpV6EthernetSegment) BgpV6PeerBgpV6EthernetSegmentIter {
	return &bgpV6PeerBgpV6EthernetSegmentIter{fieldPtr: ptr}
}

type BgpV6PeerBgpV6EthernetSegmentIter interface {
	setMsg(*bgpV6Peer) BgpV6PeerBgpV6EthernetSegmentIter
	Items() []BgpV6EthernetSegment
	Add() BgpV6EthernetSegment
	Append(items ...BgpV6EthernetSegment) BgpV6PeerBgpV6EthernetSegmentIter
	Set(index int, newObj BgpV6EthernetSegment) BgpV6PeerBgpV6EthernetSegmentIter
	Clear() BgpV6PeerBgpV6EthernetSegmentIter
	clearHolderSlice() BgpV6PeerBgpV6EthernetSegmentIter
	appendHolderSlice(item BgpV6EthernetSegment) BgpV6PeerBgpV6EthernetSegmentIter
}

func (obj *bgpV6PeerBgpV6EthernetSegmentIter) setMsg(msg *bgpV6Peer) BgpV6PeerBgpV6EthernetSegmentIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV6EthernetSegment{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6PeerBgpV6EthernetSegmentIter) Items() []BgpV6EthernetSegment {
	return obj.bgpV6EthernetSegmentSlice
}

func (obj *bgpV6PeerBgpV6EthernetSegmentIter) Add() BgpV6EthernetSegment {
	newObj := &otg.BgpV6EthernetSegment{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV6EthernetSegment{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV6EthernetSegmentSlice = append(obj.bgpV6EthernetSegmentSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6PeerBgpV6EthernetSegmentIter) Append(items ...BgpV6EthernetSegment) BgpV6PeerBgpV6EthernetSegmentIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV6EthernetSegmentSlice = append(obj.bgpV6EthernetSegmentSlice, item)
	}
	return obj
}

func (obj *bgpV6PeerBgpV6EthernetSegmentIter) Set(index int, newObj BgpV6EthernetSegment) BgpV6PeerBgpV6EthernetSegmentIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV6EthernetSegmentSlice[index] = newObj
	return obj
}
func (obj *bgpV6PeerBgpV6EthernetSegmentIter) Clear() BgpV6PeerBgpV6EthernetSegmentIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV6EthernetSegment{}
		obj.bgpV6EthernetSegmentSlice = []BgpV6EthernetSegment{}
	}
	return obj
}
func (obj *bgpV6PeerBgpV6EthernetSegmentIter) clearHolderSlice() BgpV6PeerBgpV6EthernetSegmentIter {
	if len(obj.bgpV6EthernetSegmentSlice) > 0 {
		obj.bgpV6EthernetSegmentSlice = []BgpV6EthernetSegment{}
	}
	return obj
}
func (obj *bgpV6PeerBgpV6EthernetSegmentIter) appendHolderSlice(item BgpV6EthernetSegment) BgpV6PeerBgpV6EthernetSegmentIter {
	obj.bgpV6EthernetSegmentSlice = append(obj.bgpV6EthernetSegmentSlice, item)
	return obj
}

type BgpV6PeerAsTypeEnum string

// Enum of AsType on BgpV6Peer
var BgpV6PeerAsType = struct {
	IBGP BgpV6PeerAsTypeEnum
	EBGP BgpV6PeerAsTypeEnum
}{
	IBGP: BgpV6PeerAsTypeEnum("ibgp"),
	EBGP: BgpV6PeerAsTypeEnum("ebgp"),
}

func (obj *bgpV6Peer) AsType() BgpV6PeerAsTypeEnum {
	return BgpV6PeerAsTypeEnum(obj.obj.AsType.Enum().String())
}

func (obj *bgpV6Peer) SetAsType(value BgpV6PeerAsTypeEnum) BgpV6Peer {
	intValue, ok := otg.BgpV6Peer_AsType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV6PeerAsTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV6Peer_AsType_Enum(intValue)
	obj.obj.AsType = &enumValue

	return obj
}

// Autonomous System Number (AS number or ASN)
// AsNumber returns a uint32
func (obj *bgpV6Peer) AsNumber() uint32 {

	return *obj.obj.AsNumber

}

// Autonomous System Number (AS number or ASN)
// SetAsNumber sets the uint32 value in the BgpV6Peer object
func (obj *bgpV6Peer) SetAsNumber(value uint32) BgpV6Peer {

	obj.obj.AsNumber = &value
	return obj
}

type BgpV6PeerAsNumberWidthEnum string

// Enum of AsNumberWidth on BgpV6Peer
var BgpV6PeerAsNumberWidth = struct {
	TWO  BgpV6PeerAsNumberWidthEnum
	FOUR BgpV6PeerAsNumberWidthEnum
}{
	TWO:  BgpV6PeerAsNumberWidthEnum("two"),
	FOUR: BgpV6PeerAsNumberWidthEnum("four"),
}

func (obj *bgpV6Peer) AsNumberWidth() BgpV6PeerAsNumberWidthEnum {
	return BgpV6PeerAsNumberWidthEnum(obj.obj.AsNumberWidth.Enum().String())
}

// The width in bytes of the as_number values. Any as_number values that exceeds the width MUST result in an error.
// AsNumberWidth returns a string
func (obj *bgpV6Peer) HasAsNumberWidth() bool {
	return obj.obj.AsNumberWidth != nil
}

func (obj *bgpV6Peer) SetAsNumberWidth(value BgpV6PeerAsNumberWidthEnum) BgpV6Peer {
	intValue, ok := otg.BgpV6Peer_AsNumberWidth_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV6PeerAsNumberWidthEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV6Peer_AsNumberWidth_Enum(intValue)
	obj.obj.AsNumberWidth = &enumValue

	return obj
}

// description is TBD
// Advanced returns a BgpAdvanced
func (obj *bgpV6Peer) Advanced() BgpAdvanced {
	if obj.obj.Advanced == nil {
		obj.obj.Advanced = NewBgpAdvanced().msg()
	}
	if obj.advancedHolder == nil {
		obj.advancedHolder = &bgpAdvanced{obj: obj.obj.Advanced}
	}
	return obj.advancedHolder
}

// description is TBD
// Advanced returns a BgpAdvanced
func (obj *bgpV6Peer) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// description is TBD
// SetAdvanced sets the BgpAdvanced value in the BgpV6Peer object
func (obj *bgpV6Peer) SetAdvanced(value BgpAdvanced) BgpV6Peer {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// description is TBD
// Capability returns a BgpCapability
func (obj *bgpV6Peer) Capability() BgpCapability {
	if obj.obj.Capability == nil {
		obj.obj.Capability = NewBgpCapability().msg()
	}
	if obj.capabilityHolder == nil {
		obj.capabilityHolder = &bgpCapability{obj: obj.obj.Capability}
	}
	return obj.capabilityHolder
}

// description is TBD
// Capability returns a BgpCapability
func (obj *bgpV6Peer) HasCapability() bool {
	return obj.obj.Capability != nil
}

// description is TBD
// SetCapability sets the BgpCapability value in the BgpV6Peer object
func (obj *bgpV6Peer) SetCapability(value BgpCapability) BgpV6Peer {

	obj.capabilityHolder = nil
	obj.obj.Capability = value.msg()

	return obj
}

// description is TBD
// LearnedInformationFilter returns a BgpLearnedInformationFilter
func (obj *bgpV6Peer) LearnedInformationFilter() BgpLearnedInformationFilter {
	if obj.obj.LearnedInformationFilter == nil {
		obj.obj.LearnedInformationFilter = NewBgpLearnedInformationFilter().msg()
	}
	if obj.learnedInformationFilterHolder == nil {
		obj.learnedInformationFilterHolder = &bgpLearnedInformationFilter{obj: obj.obj.LearnedInformationFilter}
	}
	return obj.learnedInformationFilterHolder
}

// description is TBD
// LearnedInformationFilter returns a BgpLearnedInformationFilter
func (obj *bgpV6Peer) HasLearnedInformationFilter() bool {
	return obj.obj.LearnedInformationFilter != nil
}

// description is TBD
// SetLearnedInformationFilter sets the BgpLearnedInformationFilter value in the BgpV6Peer object
func (obj *bgpV6Peer) SetLearnedInformationFilter(value BgpLearnedInformationFilter) BgpV6Peer {

	obj.learnedInformationFilterHolder = nil
	obj.obj.LearnedInformationFilter = value.msg()

	return obj
}

// Emulated BGPv4 route ranges.
// V4Routes returns a []BgpV4RouteRange
func (obj *bgpV6Peer) V4Routes() BgpV6PeerBgpV4RouteRangeIter {
	if len(obj.obj.V4Routes) == 0 {
		obj.obj.V4Routes = []*otg.BgpV4RouteRange{}
	}
	if obj.v4RoutesHolder == nil {
		obj.v4RoutesHolder = newBgpV6PeerBgpV4RouteRangeIter(&obj.obj.V4Routes).setMsg(obj)
	}
	return obj.v4RoutesHolder
}

type bgpV6PeerBgpV4RouteRangeIter struct {
	obj                  *bgpV6Peer
	bgpV4RouteRangeSlice []BgpV4RouteRange
	fieldPtr             *[]*otg.BgpV4RouteRange
}

func newBgpV6PeerBgpV4RouteRangeIter(ptr *[]*otg.BgpV4RouteRange) BgpV6PeerBgpV4RouteRangeIter {
	return &bgpV6PeerBgpV4RouteRangeIter{fieldPtr: ptr}
}

type BgpV6PeerBgpV4RouteRangeIter interface {
	setMsg(*bgpV6Peer) BgpV6PeerBgpV4RouteRangeIter
	Items() []BgpV4RouteRange
	Add() BgpV4RouteRange
	Append(items ...BgpV4RouteRange) BgpV6PeerBgpV4RouteRangeIter
	Set(index int, newObj BgpV4RouteRange) BgpV6PeerBgpV4RouteRangeIter
	Clear() BgpV6PeerBgpV4RouteRangeIter
	clearHolderSlice() BgpV6PeerBgpV4RouteRangeIter
	appendHolderSlice(item BgpV4RouteRange) BgpV6PeerBgpV4RouteRangeIter
}

func (obj *bgpV6PeerBgpV4RouteRangeIter) setMsg(msg *bgpV6Peer) BgpV6PeerBgpV4RouteRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV4RouteRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6PeerBgpV4RouteRangeIter) Items() []BgpV4RouteRange {
	return obj.bgpV4RouteRangeSlice
}

func (obj *bgpV6PeerBgpV4RouteRangeIter) Add() BgpV4RouteRange {
	newObj := &otg.BgpV4RouteRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV4RouteRange{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV4RouteRangeSlice = append(obj.bgpV4RouteRangeSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6PeerBgpV4RouteRangeIter) Append(items ...BgpV4RouteRange) BgpV6PeerBgpV4RouteRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV4RouteRangeSlice = append(obj.bgpV4RouteRangeSlice, item)
	}
	return obj
}

func (obj *bgpV6PeerBgpV4RouteRangeIter) Set(index int, newObj BgpV4RouteRange) BgpV6PeerBgpV4RouteRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV4RouteRangeSlice[index] = newObj
	return obj
}
func (obj *bgpV6PeerBgpV4RouteRangeIter) Clear() BgpV6PeerBgpV4RouteRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV4RouteRange{}
		obj.bgpV4RouteRangeSlice = []BgpV4RouteRange{}
	}
	return obj
}
func (obj *bgpV6PeerBgpV4RouteRangeIter) clearHolderSlice() BgpV6PeerBgpV4RouteRangeIter {
	if len(obj.bgpV4RouteRangeSlice) > 0 {
		obj.bgpV4RouteRangeSlice = []BgpV4RouteRange{}
	}
	return obj
}
func (obj *bgpV6PeerBgpV4RouteRangeIter) appendHolderSlice(item BgpV4RouteRange) BgpV6PeerBgpV4RouteRangeIter {
	obj.bgpV4RouteRangeSlice = append(obj.bgpV4RouteRangeSlice, item)
	return obj
}

// Emulated BGPv6 route ranges.
// V6Routes returns a []BgpV6RouteRange
func (obj *bgpV6Peer) V6Routes() BgpV6PeerBgpV6RouteRangeIter {
	if len(obj.obj.V6Routes) == 0 {
		obj.obj.V6Routes = []*otg.BgpV6RouteRange{}
	}
	if obj.v6RoutesHolder == nil {
		obj.v6RoutesHolder = newBgpV6PeerBgpV6RouteRangeIter(&obj.obj.V6Routes).setMsg(obj)
	}
	return obj.v6RoutesHolder
}

type bgpV6PeerBgpV6RouteRangeIter struct {
	obj                  *bgpV6Peer
	bgpV6RouteRangeSlice []BgpV6RouteRange
	fieldPtr             *[]*otg.BgpV6RouteRange
}

func newBgpV6PeerBgpV6RouteRangeIter(ptr *[]*otg.BgpV6RouteRange) BgpV6PeerBgpV6RouteRangeIter {
	return &bgpV6PeerBgpV6RouteRangeIter{fieldPtr: ptr}
}

type BgpV6PeerBgpV6RouteRangeIter interface {
	setMsg(*bgpV6Peer) BgpV6PeerBgpV6RouteRangeIter
	Items() []BgpV6RouteRange
	Add() BgpV6RouteRange
	Append(items ...BgpV6RouteRange) BgpV6PeerBgpV6RouteRangeIter
	Set(index int, newObj BgpV6RouteRange) BgpV6PeerBgpV6RouteRangeIter
	Clear() BgpV6PeerBgpV6RouteRangeIter
	clearHolderSlice() BgpV6PeerBgpV6RouteRangeIter
	appendHolderSlice(item BgpV6RouteRange) BgpV6PeerBgpV6RouteRangeIter
}

func (obj *bgpV6PeerBgpV6RouteRangeIter) setMsg(msg *bgpV6Peer) BgpV6PeerBgpV6RouteRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV6RouteRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6PeerBgpV6RouteRangeIter) Items() []BgpV6RouteRange {
	return obj.bgpV6RouteRangeSlice
}

func (obj *bgpV6PeerBgpV6RouteRangeIter) Add() BgpV6RouteRange {
	newObj := &otg.BgpV6RouteRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV6RouteRange{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV6RouteRangeSlice = append(obj.bgpV6RouteRangeSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6PeerBgpV6RouteRangeIter) Append(items ...BgpV6RouteRange) BgpV6PeerBgpV6RouteRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV6RouteRangeSlice = append(obj.bgpV6RouteRangeSlice, item)
	}
	return obj
}

func (obj *bgpV6PeerBgpV6RouteRangeIter) Set(index int, newObj BgpV6RouteRange) BgpV6PeerBgpV6RouteRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV6RouteRangeSlice[index] = newObj
	return obj
}
func (obj *bgpV6PeerBgpV6RouteRangeIter) Clear() BgpV6PeerBgpV6RouteRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV6RouteRange{}
		obj.bgpV6RouteRangeSlice = []BgpV6RouteRange{}
	}
	return obj
}
func (obj *bgpV6PeerBgpV6RouteRangeIter) clearHolderSlice() BgpV6PeerBgpV6RouteRangeIter {
	if len(obj.bgpV6RouteRangeSlice) > 0 {
		obj.bgpV6RouteRangeSlice = []BgpV6RouteRange{}
	}
	return obj
}
func (obj *bgpV6PeerBgpV6RouteRangeIter) appendHolderSlice(item BgpV6RouteRange) BgpV6PeerBgpV6RouteRangeIter {
	obj.bgpV6RouteRangeSlice = append(obj.bgpV6RouteRangeSlice, item)
	return obj
}

// Segment Routing Traffic Engineering (SR TE) Policies for IPv4 Address Family Identifier (AFI).
// V4SrtePolicies returns a []BgpSrteV4Policy
func (obj *bgpV6Peer) V4SrtePolicies() BgpV6PeerBgpSrteV4PolicyIter {
	if len(obj.obj.V4SrtePolicies) == 0 {
		obj.obj.V4SrtePolicies = []*otg.BgpSrteV4Policy{}
	}
	if obj.v4SrtePoliciesHolder == nil {
		obj.v4SrtePoliciesHolder = newBgpV6PeerBgpSrteV4PolicyIter(&obj.obj.V4SrtePolicies).setMsg(obj)
	}
	return obj.v4SrtePoliciesHolder
}

type bgpV6PeerBgpSrteV4PolicyIter struct {
	obj                  *bgpV6Peer
	bgpSrteV4PolicySlice []BgpSrteV4Policy
	fieldPtr             *[]*otg.BgpSrteV4Policy
}

func newBgpV6PeerBgpSrteV4PolicyIter(ptr *[]*otg.BgpSrteV4Policy) BgpV6PeerBgpSrteV4PolicyIter {
	return &bgpV6PeerBgpSrteV4PolicyIter{fieldPtr: ptr}
}

type BgpV6PeerBgpSrteV4PolicyIter interface {
	setMsg(*bgpV6Peer) BgpV6PeerBgpSrteV4PolicyIter
	Items() []BgpSrteV4Policy
	Add() BgpSrteV4Policy
	Append(items ...BgpSrteV4Policy) BgpV6PeerBgpSrteV4PolicyIter
	Set(index int, newObj BgpSrteV4Policy) BgpV6PeerBgpSrteV4PolicyIter
	Clear() BgpV6PeerBgpSrteV4PolicyIter
	clearHolderSlice() BgpV6PeerBgpSrteV4PolicyIter
	appendHolderSlice(item BgpSrteV4Policy) BgpV6PeerBgpSrteV4PolicyIter
}

func (obj *bgpV6PeerBgpSrteV4PolicyIter) setMsg(msg *bgpV6Peer) BgpV6PeerBgpSrteV4PolicyIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpSrteV4Policy{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6PeerBgpSrteV4PolicyIter) Items() []BgpSrteV4Policy {
	return obj.bgpSrteV4PolicySlice
}

func (obj *bgpV6PeerBgpSrteV4PolicyIter) Add() BgpSrteV4Policy {
	newObj := &otg.BgpSrteV4Policy{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpSrteV4Policy{obj: newObj}
	newLibObj.setDefault()
	obj.bgpSrteV4PolicySlice = append(obj.bgpSrteV4PolicySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6PeerBgpSrteV4PolicyIter) Append(items ...BgpSrteV4Policy) BgpV6PeerBgpSrteV4PolicyIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpSrteV4PolicySlice = append(obj.bgpSrteV4PolicySlice, item)
	}
	return obj
}

func (obj *bgpV6PeerBgpSrteV4PolicyIter) Set(index int, newObj BgpSrteV4Policy) BgpV6PeerBgpSrteV4PolicyIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpSrteV4PolicySlice[index] = newObj
	return obj
}
func (obj *bgpV6PeerBgpSrteV4PolicyIter) Clear() BgpV6PeerBgpSrteV4PolicyIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpSrteV4Policy{}
		obj.bgpSrteV4PolicySlice = []BgpSrteV4Policy{}
	}
	return obj
}
func (obj *bgpV6PeerBgpSrteV4PolicyIter) clearHolderSlice() BgpV6PeerBgpSrteV4PolicyIter {
	if len(obj.bgpSrteV4PolicySlice) > 0 {
		obj.bgpSrteV4PolicySlice = []BgpSrteV4Policy{}
	}
	return obj
}
func (obj *bgpV6PeerBgpSrteV4PolicyIter) appendHolderSlice(item BgpSrteV4Policy) BgpV6PeerBgpSrteV4PolicyIter {
	obj.bgpSrteV4PolicySlice = append(obj.bgpSrteV4PolicySlice, item)
	return obj
}

// Segment Routing Traffic Engineering (SR TE) Policies for IPv6 Address Family Identifier (AFI).
// V6SrtePolicies returns a []BgpSrteV6Policy
func (obj *bgpV6Peer) V6SrtePolicies() BgpV6PeerBgpSrteV6PolicyIter {
	if len(obj.obj.V6SrtePolicies) == 0 {
		obj.obj.V6SrtePolicies = []*otg.BgpSrteV6Policy{}
	}
	if obj.v6SrtePoliciesHolder == nil {
		obj.v6SrtePoliciesHolder = newBgpV6PeerBgpSrteV6PolicyIter(&obj.obj.V6SrtePolicies).setMsg(obj)
	}
	return obj.v6SrtePoliciesHolder
}

type bgpV6PeerBgpSrteV6PolicyIter struct {
	obj                  *bgpV6Peer
	bgpSrteV6PolicySlice []BgpSrteV6Policy
	fieldPtr             *[]*otg.BgpSrteV6Policy
}

func newBgpV6PeerBgpSrteV6PolicyIter(ptr *[]*otg.BgpSrteV6Policy) BgpV6PeerBgpSrteV6PolicyIter {
	return &bgpV6PeerBgpSrteV6PolicyIter{fieldPtr: ptr}
}

type BgpV6PeerBgpSrteV6PolicyIter interface {
	setMsg(*bgpV6Peer) BgpV6PeerBgpSrteV6PolicyIter
	Items() []BgpSrteV6Policy
	Add() BgpSrteV6Policy
	Append(items ...BgpSrteV6Policy) BgpV6PeerBgpSrteV6PolicyIter
	Set(index int, newObj BgpSrteV6Policy) BgpV6PeerBgpSrteV6PolicyIter
	Clear() BgpV6PeerBgpSrteV6PolicyIter
	clearHolderSlice() BgpV6PeerBgpSrteV6PolicyIter
	appendHolderSlice(item BgpSrteV6Policy) BgpV6PeerBgpSrteV6PolicyIter
}

func (obj *bgpV6PeerBgpSrteV6PolicyIter) setMsg(msg *bgpV6Peer) BgpV6PeerBgpSrteV6PolicyIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpSrteV6Policy{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6PeerBgpSrteV6PolicyIter) Items() []BgpSrteV6Policy {
	return obj.bgpSrteV6PolicySlice
}

func (obj *bgpV6PeerBgpSrteV6PolicyIter) Add() BgpSrteV6Policy {
	newObj := &otg.BgpSrteV6Policy{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpSrteV6Policy{obj: newObj}
	newLibObj.setDefault()
	obj.bgpSrteV6PolicySlice = append(obj.bgpSrteV6PolicySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6PeerBgpSrteV6PolicyIter) Append(items ...BgpSrteV6Policy) BgpV6PeerBgpSrteV6PolicyIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpSrteV6PolicySlice = append(obj.bgpSrteV6PolicySlice, item)
	}
	return obj
}

func (obj *bgpV6PeerBgpSrteV6PolicyIter) Set(index int, newObj BgpSrteV6Policy) BgpV6PeerBgpSrteV6PolicyIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpSrteV6PolicySlice[index] = newObj
	return obj
}
func (obj *bgpV6PeerBgpSrteV6PolicyIter) Clear() BgpV6PeerBgpSrteV6PolicyIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpSrteV6Policy{}
		obj.bgpSrteV6PolicySlice = []BgpSrteV6Policy{}
	}
	return obj
}
func (obj *bgpV6PeerBgpSrteV6PolicyIter) clearHolderSlice() BgpV6PeerBgpSrteV6PolicyIter {
	if len(obj.bgpSrteV6PolicySlice) > 0 {
		obj.bgpSrteV6PolicySlice = []BgpSrteV6Policy{}
	}
	return obj
}
func (obj *bgpV6PeerBgpSrteV6PolicyIter) appendHolderSlice(item BgpSrteV6Policy) BgpV6PeerBgpSrteV6PolicyIter {
	obj.bgpSrteV6PolicySlice = append(obj.bgpSrteV6PolicySlice, item)
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpV6Peer) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpV6Peer object
func (obj *bgpV6Peer) SetName(value string) BgpV6Peer {

	obj.obj.Name = &value
	return obj
}

// description is TBD
// GracefulRestart returns a BgpGracefulRestart
func (obj *bgpV6Peer) GracefulRestart() BgpGracefulRestart {
	if obj.obj.GracefulRestart == nil {
		obj.obj.GracefulRestart = NewBgpGracefulRestart().msg()
	}
	if obj.gracefulRestartHolder == nil {
		obj.gracefulRestartHolder = &bgpGracefulRestart{obj: obj.obj.GracefulRestart}
	}
	return obj.gracefulRestartHolder
}

// description is TBD
// GracefulRestart returns a BgpGracefulRestart
func (obj *bgpV6Peer) HasGracefulRestart() bool {
	return obj.obj.GracefulRestart != nil
}

// description is TBD
// SetGracefulRestart sets the BgpGracefulRestart value in the BgpV6Peer object
func (obj *bgpV6Peer) SetGracefulRestart(value BgpGracefulRestart) BgpV6Peer {

	obj.gracefulRestartHolder = nil
	obj.obj.GracefulRestart = value.msg()

	return obj
}

// BGP Updates to be sent to the peer as specified after the session is established.
// ReplayUpdates returns a BgpUpdateReplay
func (obj *bgpV6Peer) ReplayUpdates() BgpUpdateReplay {
	if obj.obj.ReplayUpdates == nil {
		obj.obj.ReplayUpdates = NewBgpUpdateReplay().msg()
	}
	if obj.replayUpdatesHolder == nil {
		obj.replayUpdatesHolder = &bgpUpdateReplay{obj: obj.obj.ReplayUpdates}
	}
	return obj.replayUpdatesHolder
}

// BGP Updates to be sent to the peer as specified after the session is established.
// ReplayUpdates returns a BgpUpdateReplay
func (obj *bgpV6Peer) HasReplayUpdates() bool {
	return obj.obj.ReplayUpdates != nil
}

// BGP Updates to be sent to the peer as specified after the session is established.
// SetReplayUpdates sets the BgpUpdateReplay value in the BgpV6Peer object
func (obj *bgpV6Peer) SetReplayUpdates(value BgpUpdateReplay) BgpV6Peer {

	obj.replayUpdatesHolder = nil
	obj.obj.ReplayUpdates = value.msg()

	return obj
}

func (obj *bgpV6Peer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// PeerAddress is required
	if obj.obj.PeerAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "PeerAddress is required field on interface BgpV6Peer")
	}
	if obj.obj.PeerAddress != nil {

		err := obj.validateIpv6(obj.PeerAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpV6Peer.PeerAddress"))
		}

	}

	if obj.obj.SegmentRouting != nil {

		obj.SegmentRouting().validateObj(vObj, set_default)
	}

	if len(obj.obj.EvpnEthernetSegments) != 0 {

		if set_default {
			obj.EvpnEthernetSegments().clearHolderSlice()
			for _, item := range obj.obj.EvpnEthernetSegments {
				obj.EvpnEthernetSegments().appendHolderSlice(&bgpV6EthernetSegment{obj: item})
			}
		}
		for _, item := range obj.EvpnEthernetSegments().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// AsType is required
	if obj.obj.AsType == nil {
		vObj.validationErrors = append(vObj.validationErrors, "AsType is required field on interface BgpV6Peer")
	}

	// AsNumber is required
	if obj.obj.AsNumber == nil {
		vObj.validationErrors = append(vObj.validationErrors, "AsNumber is required field on interface BgpV6Peer")
	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(vObj, set_default)
	}

	if obj.obj.Capability != nil {

		obj.Capability().validateObj(vObj, set_default)
	}

	if obj.obj.LearnedInformationFilter != nil {

		obj.LearnedInformationFilter().validateObj(vObj, set_default)
	}

	if len(obj.obj.V4Routes) != 0 {

		if set_default {
			obj.V4Routes().clearHolderSlice()
			for _, item := range obj.obj.V4Routes {
				obj.V4Routes().appendHolderSlice(&bgpV4RouteRange{obj: item})
			}
		}
		for _, item := range obj.V4Routes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.V6Routes) != 0 {

		if set_default {
			obj.V6Routes().clearHolderSlice()
			for _, item := range obj.obj.V6Routes {
				obj.V6Routes().appendHolderSlice(&bgpV6RouteRange{obj: item})
			}
		}
		for _, item := range obj.V6Routes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.V4SrtePolicies) != 0 {

		if set_default {
			obj.V4SrtePolicies().clearHolderSlice()
			for _, item := range obj.obj.V4SrtePolicies {
				obj.V4SrtePolicies().appendHolderSlice(&bgpSrteV4Policy{obj: item})
			}
		}
		for _, item := range obj.V4SrtePolicies().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.V6SrtePolicies) != 0 {

		if set_default {
			obj.V6SrtePolicies().clearHolderSlice()
			for _, item := range obj.obj.V6SrtePolicies {
				obj.V6SrtePolicies().appendHolderSlice(&bgpSrteV6Policy{obj: item})
			}
		}
		for _, item := range obj.V6SrtePolicies().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpV6Peer")
	}

	if obj.obj.GracefulRestart != nil {

		obj.GracefulRestart().validateObj(vObj, set_default)
	}

	if obj.obj.ReplayUpdates != nil {

		obj.ReplayUpdates().validateObj(vObj, set_default)
	}

}

func (obj *bgpV6Peer) setDefault() {
	if obj.obj.AsNumberWidth == nil {
		obj.SetAsNumberWidth(BgpV6PeerAsNumberWidth.FOUR)

	}

}
