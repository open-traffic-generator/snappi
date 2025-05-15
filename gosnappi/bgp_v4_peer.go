package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV4Peer *****
type bgpV4Peer struct {
	validation
	obj                            *otg.BgpV4Peer
	marshaller                     marshalBgpV4Peer
	unMarshaller                   unMarshalBgpV4Peer
	evpnEthernetSegmentsHolder     BgpV4PeerBgpV4EthernetSegmentIter
	advancedHolder                 BgpAdvanced
	capabilityHolder               BgpCapability
	learnedInformationFilterHolder BgpLearnedInformationFilter
	v4RoutesHolder                 BgpV4PeerBgpV4RouteRangeIter
	v6RoutesHolder                 BgpV4PeerBgpV6RouteRangeIter
	v4SrtePoliciesHolder           BgpV4PeerBgpSrteV4PolicyIter
	v6SrtePoliciesHolder           BgpV4PeerBgpSrteV6PolicyIter
	gracefulRestartHolder          BgpGracefulRestart
	replayUpdatesHolder            BgpUpdateReplay
}

func NewBgpV4Peer() BgpV4Peer {
	obj := bgpV4Peer{obj: &otg.BgpV4Peer{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV4Peer) msg() *otg.BgpV4Peer {
	return obj.obj
}

func (obj *bgpV4Peer) setMsg(msg *otg.BgpV4Peer) BgpV4Peer {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV4Peer struct {
	obj *bgpV4Peer
}

type marshalBgpV4Peer interface {
	// ToProto marshals BgpV4Peer to protobuf object *otg.BgpV4Peer
	ToProto() (*otg.BgpV4Peer, error)
	// ToPbText marshals BgpV4Peer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV4Peer to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV4Peer to JSON text
	ToJson() (string, error)
}

type unMarshalbgpV4Peer struct {
	obj *bgpV4Peer
}

type unMarshalBgpV4Peer interface {
	// FromProto unmarshals BgpV4Peer from protobuf object *otg.BgpV4Peer
	FromProto(msg *otg.BgpV4Peer) (BgpV4Peer, error)
	// FromPbText unmarshals BgpV4Peer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV4Peer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV4Peer from JSON text
	FromJson(value string) error
}

func (obj *bgpV4Peer) Marshal() marshalBgpV4Peer {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV4Peer{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV4Peer) Unmarshal() unMarshalBgpV4Peer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV4Peer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV4Peer) ToProto() (*otg.BgpV4Peer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV4Peer) FromProto(msg *otg.BgpV4Peer) (BgpV4Peer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV4Peer) ToPbText() (string, error) {
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

func (m *unMarshalbgpV4Peer) FromPbText(value string) error {
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

func (m *marshalbgpV4Peer) ToYaml() (string, error) {
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

func (m *unMarshalbgpV4Peer) FromYaml(value string) error {
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

func (m *marshalbgpV4Peer) ToJson() (string, error) {
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

func (m *unMarshalbgpV4Peer) FromJson(value string) error {
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

func (obj *bgpV4Peer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV4Peer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV4Peer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV4Peer) Clone() (BgpV4Peer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV4Peer()
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

func (obj *bgpV4Peer) setNil() {
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

// BgpV4Peer is configuration for emulated BGPv4 peers and routes.
type BgpV4Peer interface {
	Validation
	// msg marshals BgpV4Peer to protobuf object *otg.BgpV4Peer
	// and doesn't set defaults
	msg() *otg.BgpV4Peer
	// setMsg unmarshals BgpV4Peer from protobuf object *otg.BgpV4Peer
	// and doesn't set defaults
	setMsg(*otg.BgpV4Peer) BgpV4Peer
	// provides marshal interface
	Marshal() marshalBgpV4Peer
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV4Peer
	// validate validates BgpV4Peer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV4Peer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PeerAddress returns string, set in BgpV4Peer.
	PeerAddress() string
	// SetPeerAddress assigns string provided by user to BgpV4Peer
	SetPeerAddress(value string) BgpV4Peer
	// EvpnEthernetSegments returns BgpV4PeerBgpV4EthernetSegmentIterIter, set in BgpV4Peer
	EvpnEthernetSegments() BgpV4PeerBgpV4EthernetSegmentIter
	// AsType returns BgpV4PeerAsTypeEnum, set in BgpV4Peer
	AsType() BgpV4PeerAsTypeEnum
	// SetAsType assigns BgpV4PeerAsTypeEnum provided by user to BgpV4Peer
	SetAsType(value BgpV4PeerAsTypeEnum) BgpV4Peer
	// AsNumber returns uint32, set in BgpV4Peer.
	AsNumber() uint32
	// SetAsNumber assigns uint32 provided by user to BgpV4Peer
	SetAsNumber(value uint32) BgpV4Peer
	// AsNumberWidth returns BgpV4PeerAsNumberWidthEnum, set in BgpV4Peer
	AsNumberWidth() BgpV4PeerAsNumberWidthEnum
	// SetAsNumberWidth assigns BgpV4PeerAsNumberWidthEnum provided by user to BgpV4Peer
	SetAsNumberWidth(value BgpV4PeerAsNumberWidthEnum) BgpV4Peer
	// HasAsNumberWidth checks if AsNumberWidth has been set in BgpV4Peer
	HasAsNumberWidth() bool
	// Advanced returns BgpAdvanced, set in BgpV4Peer.
	// BgpAdvanced is configuration for BGP advanced settings.
	Advanced() BgpAdvanced
	// SetAdvanced assigns BgpAdvanced provided by user to BgpV4Peer.
	// BgpAdvanced is configuration for BGP advanced settings.
	SetAdvanced(value BgpAdvanced) BgpV4Peer
	// HasAdvanced checks if Advanced has been set in BgpV4Peer
	HasAdvanced() bool
	// Capability returns BgpCapability, set in BgpV4Peer.
	// BgpCapability is configuration for BGP capability settings.
	Capability() BgpCapability
	// SetCapability assigns BgpCapability provided by user to BgpV4Peer.
	// BgpCapability is configuration for BGP capability settings.
	SetCapability(value BgpCapability) BgpV4Peer
	// HasCapability checks if Capability has been set in BgpV4Peer
	HasCapability() bool
	// LearnedInformationFilter returns BgpLearnedInformationFilter, set in BgpV4Peer.
	// BgpLearnedInformationFilter is configuration for controlling storage of BGP learned information recieved from the peer.
	LearnedInformationFilter() BgpLearnedInformationFilter
	// SetLearnedInformationFilter assigns BgpLearnedInformationFilter provided by user to BgpV4Peer.
	// BgpLearnedInformationFilter is configuration for controlling storage of BGP learned information recieved from the peer.
	SetLearnedInformationFilter(value BgpLearnedInformationFilter) BgpV4Peer
	// HasLearnedInformationFilter checks if LearnedInformationFilter has been set in BgpV4Peer
	HasLearnedInformationFilter() bool
	// V4Routes returns BgpV4PeerBgpV4RouteRangeIterIter, set in BgpV4Peer
	V4Routes() BgpV4PeerBgpV4RouteRangeIter
	// V6Routes returns BgpV4PeerBgpV6RouteRangeIterIter, set in BgpV4Peer
	V6Routes() BgpV4PeerBgpV6RouteRangeIter
	// V4SrtePolicies returns BgpV4PeerBgpSrteV4PolicyIterIter, set in BgpV4Peer
	V4SrtePolicies() BgpV4PeerBgpSrteV4PolicyIter
	// V6SrtePolicies returns BgpV4PeerBgpSrteV6PolicyIterIter, set in BgpV4Peer
	V6SrtePolicies() BgpV4PeerBgpSrteV6PolicyIter
	// Name returns string, set in BgpV4Peer.
	Name() string
	// SetName assigns string provided by user to BgpV4Peer
	SetName(value string) BgpV4Peer
	// GracefulRestart returns BgpGracefulRestart, set in BgpV4Peer.
	// BgpGracefulRestart is the Graceful Restart Capability (RFC 4724) is a BGP capability that can be used by a BGP speaker to indicate its ability to preserve its forwarding state during BGP restart. The Graceful Restart (GR) capability is advertised in OPEN messages sent between BGP peers. After a BGP session has been established, and the initial routing update has been completed,  an End-of-RIB (Routing Information Base) marker is sent in an UPDATE message to convey information  about routing convergence.
	GracefulRestart() BgpGracefulRestart
	// SetGracefulRestart assigns BgpGracefulRestart provided by user to BgpV4Peer.
	// BgpGracefulRestart is the Graceful Restart Capability (RFC 4724) is a BGP capability that can be used by a BGP speaker to indicate its ability to preserve its forwarding state during BGP restart. The Graceful Restart (GR) capability is advertised in OPEN messages sent between BGP peers. After a BGP session has been established, and the initial routing update has been completed,  an End-of-RIB (Routing Information Base) marker is sent in an UPDATE message to convey information  about routing convergence.
	SetGracefulRestart(value BgpGracefulRestart) BgpV4Peer
	// HasGracefulRestart checks if GracefulRestart has been set in BgpV4Peer
	HasGracefulRestart() bool
	// ReplayUpdates returns BgpUpdateReplay, set in BgpV4Peer.
	// BgpUpdateReplay is ordered BGP Updates ( including both Advertise and Withdraws ) to be sent in the order given in the input to the peer after the BGP session is established.
	ReplayUpdates() BgpUpdateReplay
	// SetReplayUpdates assigns BgpUpdateReplay provided by user to BgpV4Peer.
	// BgpUpdateReplay is ordered BGP Updates ( including both Advertise and Withdraws ) to be sent in the order given in the input to the peer after the BGP session is established.
	SetReplayUpdates(value BgpUpdateReplay) BgpV4Peer
	// HasReplayUpdates checks if ReplayUpdates has been set in BgpV4Peer
	HasReplayUpdates() bool
	setNil()
}

// IPv4 address of the BGP peer for the session.
// PeerAddress returns a string
func (obj *bgpV4Peer) PeerAddress() string {

	return *obj.obj.PeerAddress

}

// IPv4 address of the BGP peer for the session.
// SetPeerAddress sets the string value in the BgpV4Peer object
func (obj *bgpV4Peer) SetPeerAddress(value string) BgpV4Peer {

	obj.obj.PeerAddress = &value
	return obj
}

// This contains the list of Ethernet Virtual Private Network (EVPN) Ethernet Segments (ES) Per BGP Peer for IPv4 Address Family Identifier (AFI).
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
// EvpnEthernetSegments returns a []BgpV4EthernetSegment
func (obj *bgpV4Peer) EvpnEthernetSegments() BgpV4PeerBgpV4EthernetSegmentIter {
	if len(obj.obj.EvpnEthernetSegments) == 0 {
		obj.obj.EvpnEthernetSegments = []*otg.BgpV4EthernetSegment{}
	}
	if obj.evpnEthernetSegmentsHolder == nil {
		obj.evpnEthernetSegmentsHolder = newBgpV4PeerBgpV4EthernetSegmentIter(&obj.obj.EvpnEthernetSegments).setMsg(obj)
	}
	return obj.evpnEthernetSegmentsHolder
}

type bgpV4PeerBgpV4EthernetSegmentIter struct {
	obj                       *bgpV4Peer
	bgpV4EthernetSegmentSlice []BgpV4EthernetSegment
	fieldPtr                  *[]*otg.BgpV4EthernetSegment
}

func newBgpV4PeerBgpV4EthernetSegmentIter(ptr *[]*otg.BgpV4EthernetSegment) BgpV4PeerBgpV4EthernetSegmentIter {
	return &bgpV4PeerBgpV4EthernetSegmentIter{fieldPtr: ptr}
}

type BgpV4PeerBgpV4EthernetSegmentIter interface {
	setMsg(*bgpV4Peer) BgpV4PeerBgpV4EthernetSegmentIter
	Items() []BgpV4EthernetSegment
	Add() BgpV4EthernetSegment
	Append(items ...BgpV4EthernetSegment) BgpV4PeerBgpV4EthernetSegmentIter
	Set(index int, newObj BgpV4EthernetSegment) BgpV4PeerBgpV4EthernetSegmentIter
	Clear() BgpV4PeerBgpV4EthernetSegmentIter
	clearHolderSlice() BgpV4PeerBgpV4EthernetSegmentIter
	appendHolderSlice(item BgpV4EthernetSegment) BgpV4PeerBgpV4EthernetSegmentIter
}

func (obj *bgpV4PeerBgpV4EthernetSegmentIter) setMsg(msg *bgpV4Peer) BgpV4PeerBgpV4EthernetSegmentIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV4EthernetSegment{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4PeerBgpV4EthernetSegmentIter) Items() []BgpV4EthernetSegment {
	return obj.bgpV4EthernetSegmentSlice
}

func (obj *bgpV4PeerBgpV4EthernetSegmentIter) Add() BgpV4EthernetSegment {
	newObj := &otg.BgpV4EthernetSegment{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV4EthernetSegment{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV4EthernetSegmentSlice = append(obj.bgpV4EthernetSegmentSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4PeerBgpV4EthernetSegmentIter) Append(items ...BgpV4EthernetSegment) BgpV4PeerBgpV4EthernetSegmentIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV4EthernetSegmentSlice = append(obj.bgpV4EthernetSegmentSlice, item)
	}
	return obj
}

func (obj *bgpV4PeerBgpV4EthernetSegmentIter) Set(index int, newObj BgpV4EthernetSegment) BgpV4PeerBgpV4EthernetSegmentIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV4EthernetSegmentSlice[index] = newObj
	return obj
}
func (obj *bgpV4PeerBgpV4EthernetSegmentIter) Clear() BgpV4PeerBgpV4EthernetSegmentIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV4EthernetSegment{}
		obj.bgpV4EthernetSegmentSlice = []BgpV4EthernetSegment{}
	}
	return obj
}
func (obj *bgpV4PeerBgpV4EthernetSegmentIter) clearHolderSlice() BgpV4PeerBgpV4EthernetSegmentIter {
	if len(obj.bgpV4EthernetSegmentSlice) > 0 {
		obj.bgpV4EthernetSegmentSlice = []BgpV4EthernetSegment{}
	}
	return obj
}
func (obj *bgpV4PeerBgpV4EthernetSegmentIter) appendHolderSlice(item BgpV4EthernetSegment) BgpV4PeerBgpV4EthernetSegmentIter {
	obj.bgpV4EthernetSegmentSlice = append(obj.bgpV4EthernetSegmentSlice, item)
	return obj
}

type BgpV4PeerAsTypeEnum string

// Enum of AsType on BgpV4Peer
var BgpV4PeerAsType = struct {
	IBGP BgpV4PeerAsTypeEnum
	EBGP BgpV4PeerAsTypeEnum
}{
	IBGP: BgpV4PeerAsTypeEnum("ibgp"),
	EBGP: BgpV4PeerAsTypeEnum("ebgp"),
}

func (obj *bgpV4Peer) AsType() BgpV4PeerAsTypeEnum {
	return BgpV4PeerAsTypeEnum(obj.obj.AsType.Enum().String())
}

func (obj *bgpV4Peer) SetAsType(value BgpV4PeerAsTypeEnum) BgpV4Peer {
	intValue, ok := otg.BgpV4Peer_AsType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV4PeerAsTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV4Peer_AsType_Enum(intValue)
	obj.obj.AsType = &enumValue

	return obj
}

// Autonomous System Number (AS number or ASN)
// AsNumber returns a uint32
func (obj *bgpV4Peer) AsNumber() uint32 {

	return *obj.obj.AsNumber

}

// Autonomous System Number (AS number or ASN)
// SetAsNumber sets the uint32 value in the BgpV4Peer object
func (obj *bgpV4Peer) SetAsNumber(value uint32) BgpV4Peer {

	obj.obj.AsNumber = &value
	return obj
}

type BgpV4PeerAsNumberWidthEnum string

// Enum of AsNumberWidth on BgpV4Peer
var BgpV4PeerAsNumberWidth = struct {
	TWO  BgpV4PeerAsNumberWidthEnum
	FOUR BgpV4PeerAsNumberWidthEnum
}{
	TWO:  BgpV4PeerAsNumberWidthEnum("two"),
	FOUR: BgpV4PeerAsNumberWidthEnum("four"),
}

func (obj *bgpV4Peer) AsNumberWidth() BgpV4PeerAsNumberWidthEnum {
	return BgpV4PeerAsNumberWidthEnum(obj.obj.AsNumberWidth.Enum().String())
}

// The width in bytes of the as_number values. Any as_number values that exceeds the width MUST result in an error.
// AsNumberWidth returns a string
func (obj *bgpV4Peer) HasAsNumberWidth() bool {
	return obj.obj.AsNumberWidth != nil
}

func (obj *bgpV4Peer) SetAsNumberWidth(value BgpV4PeerAsNumberWidthEnum) BgpV4Peer {
	intValue, ok := otg.BgpV4Peer_AsNumberWidth_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV4PeerAsNumberWidthEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV4Peer_AsNumberWidth_Enum(intValue)
	obj.obj.AsNumberWidth = &enumValue

	return obj
}

// description is TBD
// Advanced returns a BgpAdvanced
func (obj *bgpV4Peer) Advanced() BgpAdvanced {
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
func (obj *bgpV4Peer) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// description is TBD
// SetAdvanced sets the BgpAdvanced value in the BgpV4Peer object
func (obj *bgpV4Peer) SetAdvanced(value BgpAdvanced) BgpV4Peer {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// description is TBD
// Capability returns a BgpCapability
func (obj *bgpV4Peer) Capability() BgpCapability {
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
func (obj *bgpV4Peer) HasCapability() bool {
	return obj.obj.Capability != nil
}

// description is TBD
// SetCapability sets the BgpCapability value in the BgpV4Peer object
func (obj *bgpV4Peer) SetCapability(value BgpCapability) BgpV4Peer {

	obj.capabilityHolder = nil
	obj.obj.Capability = value.msg()

	return obj
}

// description is TBD
// LearnedInformationFilter returns a BgpLearnedInformationFilter
func (obj *bgpV4Peer) LearnedInformationFilter() BgpLearnedInformationFilter {
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
func (obj *bgpV4Peer) HasLearnedInformationFilter() bool {
	return obj.obj.LearnedInformationFilter != nil
}

// description is TBD
// SetLearnedInformationFilter sets the BgpLearnedInformationFilter value in the BgpV4Peer object
func (obj *bgpV4Peer) SetLearnedInformationFilter(value BgpLearnedInformationFilter) BgpV4Peer {

	obj.learnedInformationFilterHolder = nil
	obj.obj.LearnedInformationFilter = value.msg()

	return obj
}

// Emulated BGPv4 route ranges.
// V4Routes returns a []BgpV4RouteRange
func (obj *bgpV4Peer) V4Routes() BgpV4PeerBgpV4RouteRangeIter {
	if len(obj.obj.V4Routes) == 0 {
		obj.obj.V4Routes = []*otg.BgpV4RouteRange{}
	}
	if obj.v4RoutesHolder == nil {
		obj.v4RoutesHolder = newBgpV4PeerBgpV4RouteRangeIter(&obj.obj.V4Routes).setMsg(obj)
	}
	return obj.v4RoutesHolder
}

type bgpV4PeerBgpV4RouteRangeIter struct {
	obj                  *bgpV4Peer
	bgpV4RouteRangeSlice []BgpV4RouteRange
	fieldPtr             *[]*otg.BgpV4RouteRange
}

func newBgpV4PeerBgpV4RouteRangeIter(ptr *[]*otg.BgpV4RouteRange) BgpV4PeerBgpV4RouteRangeIter {
	return &bgpV4PeerBgpV4RouteRangeIter{fieldPtr: ptr}
}

type BgpV4PeerBgpV4RouteRangeIter interface {
	setMsg(*bgpV4Peer) BgpV4PeerBgpV4RouteRangeIter
	Items() []BgpV4RouteRange
	Add() BgpV4RouteRange
	Append(items ...BgpV4RouteRange) BgpV4PeerBgpV4RouteRangeIter
	Set(index int, newObj BgpV4RouteRange) BgpV4PeerBgpV4RouteRangeIter
	Clear() BgpV4PeerBgpV4RouteRangeIter
	clearHolderSlice() BgpV4PeerBgpV4RouteRangeIter
	appendHolderSlice(item BgpV4RouteRange) BgpV4PeerBgpV4RouteRangeIter
}

func (obj *bgpV4PeerBgpV4RouteRangeIter) setMsg(msg *bgpV4Peer) BgpV4PeerBgpV4RouteRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV4RouteRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4PeerBgpV4RouteRangeIter) Items() []BgpV4RouteRange {
	return obj.bgpV4RouteRangeSlice
}

func (obj *bgpV4PeerBgpV4RouteRangeIter) Add() BgpV4RouteRange {
	newObj := &otg.BgpV4RouteRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV4RouteRange{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV4RouteRangeSlice = append(obj.bgpV4RouteRangeSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4PeerBgpV4RouteRangeIter) Append(items ...BgpV4RouteRange) BgpV4PeerBgpV4RouteRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV4RouteRangeSlice = append(obj.bgpV4RouteRangeSlice, item)
	}
	return obj
}

func (obj *bgpV4PeerBgpV4RouteRangeIter) Set(index int, newObj BgpV4RouteRange) BgpV4PeerBgpV4RouteRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV4RouteRangeSlice[index] = newObj
	return obj
}
func (obj *bgpV4PeerBgpV4RouteRangeIter) Clear() BgpV4PeerBgpV4RouteRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV4RouteRange{}
		obj.bgpV4RouteRangeSlice = []BgpV4RouteRange{}
	}
	return obj
}
func (obj *bgpV4PeerBgpV4RouteRangeIter) clearHolderSlice() BgpV4PeerBgpV4RouteRangeIter {
	if len(obj.bgpV4RouteRangeSlice) > 0 {
		obj.bgpV4RouteRangeSlice = []BgpV4RouteRange{}
	}
	return obj
}
func (obj *bgpV4PeerBgpV4RouteRangeIter) appendHolderSlice(item BgpV4RouteRange) BgpV4PeerBgpV4RouteRangeIter {
	obj.bgpV4RouteRangeSlice = append(obj.bgpV4RouteRangeSlice, item)
	return obj
}

// Emulated BGPv6 route ranges.
// V6Routes returns a []BgpV6RouteRange
func (obj *bgpV4Peer) V6Routes() BgpV4PeerBgpV6RouteRangeIter {
	if len(obj.obj.V6Routes) == 0 {
		obj.obj.V6Routes = []*otg.BgpV6RouteRange{}
	}
	if obj.v6RoutesHolder == nil {
		obj.v6RoutesHolder = newBgpV4PeerBgpV6RouteRangeIter(&obj.obj.V6Routes).setMsg(obj)
	}
	return obj.v6RoutesHolder
}

type bgpV4PeerBgpV6RouteRangeIter struct {
	obj                  *bgpV4Peer
	bgpV6RouteRangeSlice []BgpV6RouteRange
	fieldPtr             *[]*otg.BgpV6RouteRange
}

func newBgpV4PeerBgpV6RouteRangeIter(ptr *[]*otg.BgpV6RouteRange) BgpV4PeerBgpV6RouteRangeIter {
	return &bgpV4PeerBgpV6RouteRangeIter{fieldPtr: ptr}
}

type BgpV4PeerBgpV6RouteRangeIter interface {
	setMsg(*bgpV4Peer) BgpV4PeerBgpV6RouteRangeIter
	Items() []BgpV6RouteRange
	Add() BgpV6RouteRange
	Append(items ...BgpV6RouteRange) BgpV4PeerBgpV6RouteRangeIter
	Set(index int, newObj BgpV6RouteRange) BgpV4PeerBgpV6RouteRangeIter
	Clear() BgpV4PeerBgpV6RouteRangeIter
	clearHolderSlice() BgpV4PeerBgpV6RouteRangeIter
	appendHolderSlice(item BgpV6RouteRange) BgpV4PeerBgpV6RouteRangeIter
}

func (obj *bgpV4PeerBgpV6RouteRangeIter) setMsg(msg *bgpV4Peer) BgpV4PeerBgpV6RouteRangeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV6RouteRange{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4PeerBgpV6RouteRangeIter) Items() []BgpV6RouteRange {
	return obj.bgpV6RouteRangeSlice
}

func (obj *bgpV4PeerBgpV6RouteRangeIter) Add() BgpV6RouteRange {
	newObj := &otg.BgpV6RouteRange{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV6RouteRange{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV6RouteRangeSlice = append(obj.bgpV6RouteRangeSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4PeerBgpV6RouteRangeIter) Append(items ...BgpV6RouteRange) BgpV4PeerBgpV6RouteRangeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV6RouteRangeSlice = append(obj.bgpV6RouteRangeSlice, item)
	}
	return obj
}

func (obj *bgpV4PeerBgpV6RouteRangeIter) Set(index int, newObj BgpV6RouteRange) BgpV4PeerBgpV6RouteRangeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV6RouteRangeSlice[index] = newObj
	return obj
}
func (obj *bgpV4PeerBgpV6RouteRangeIter) Clear() BgpV4PeerBgpV6RouteRangeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV6RouteRange{}
		obj.bgpV6RouteRangeSlice = []BgpV6RouteRange{}
	}
	return obj
}
func (obj *bgpV4PeerBgpV6RouteRangeIter) clearHolderSlice() BgpV4PeerBgpV6RouteRangeIter {
	if len(obj.bgpV6RouteRangeSlice) > 0 {
		obj.bgpV6RouteRangeSlice = []BgpV6RouteRange{}
	}
	return obj
}
func (obj *bgpV4PeerBgpV6RouteRangeIter) appendHolderSlice(item BgpV6RouteRange) BgpV4PeerBgpV6RouteRangeIter {
	obj.bgpV6RouteRangeSlice = append(obj.bgpV6RouteRangeSlice, item)
	return obj
}

// Segment Routing Traffic Engineering (SR TE) Policies for IPv4 Address Family Identifier (AFI).
// V4SrtePolicies returns a []BgpSrteV4Policy
func (obj *bgpV4Peer) V4SrtePolicies() BgpV4PeerBgpSrteV4PolicyIter {
	if len(obj.obj.V4SrtePolicies) == 0 {
		obj.obj.V4SrtePolicies = []*otg.BgpSrteV4Policy{}
	}
	if obj.v4SrtePoliciesHolder == nil {
		obj.v4SrtePoliciesHolder = newBgpV4PeerBgpSrteV4PolicyIter(&obj.obj.V4SrtePolicies).setMsg(obj)
	}
	return obj.v4SrtePoliciesHolder
}

type bgpV4PeerBgpSrteV4PolicyIter struct {
	obj                  *bgpV4Peer
	bgpSrteV4PolicySlice []BgpSrteV4Policy
	fieldPtr             *[]*otg.BgpSrteV4Policy
}

func newBgpV4PeerBgpSrteV4PolicyIter(ptr *[]*otg.BgpSrteV4Policy) BgpV4PeerBgpSrteV4PolicyIter {
	return &bgpV4PeerBgpSrteV4PolicyIter{fieldPtr: ptr}
}

type BgpV4PeerBgpSrteV4PolicyIter interface {
	setMsg(*bgpV4Peer) BgpV4PeerBgpSrteV4PolicyIter
	Items() []BgpSrteV4Policy
	Add() BgpSrteV4Policy
	Append(items ...BgpSrteV4Policy) BgpV4PeerBgpSrteV4PolicyIter
	Set(index int, newObj BgpSrteV4Policy) BgpV4PeerBgpSrteV4PolicyIter
	Clear() BgpV4PeerBgpSrteV4PolicyIter
	clearHolderSlice() BgpV4PeerBgpSrteV4PolicyIter
	appendHolderSlice(item BgpSrteV4Policy) BgpV4PeerBgpSrteV4PolicyIter
}

func (obj *bgpV4PeerBgpSrteV4PolicyIter) setMsg(msg *bgpV4Peer) BgpV4PeerBgpSrteV4PolicyIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpSrteV4Policy{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4PeerBgpSrteV4PolicyIter) Items() []BgpSrteV4Policy {
	return obj.bgpSrteV4PolicySlice
}

func (obj *bgpV4PeerBgpSrteV4PolicyIter) Add() BgpSrteV4Policy {
	newObj := &otg.BgpSrteV4Policy{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpSrteV4Policy{obj: newObj}
	newLibObj.setDefault()
	obj.bgpSrteV4PolicySlice = append(obj.bgpSrteV4PolicySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4PeerBgpSrteV4PolicyIter) Append(items ...BgpSrteV4Policy) BgpV4PeerBgpSrteV4PolicyIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpSrteV4PolicySlice = append(obj.bgpSrteV4PolicySlice, item)
	}
	return obj
}

func (obj *bgpV4PeerBgpSrteV4PolicyIter) Set(index int, newObj BgpSrteV4Policy) BgpV4PeerBgpSrteV4PolicyIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpSrteV4PolicySlice[index] = newObj
	return obj
}
func (obj *bgpV4PeerBgpSrteV4PolicyIter) Clear() BgpV4PeerBgpSrteV4PolicyIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpSrteV4Policy{}
		obj.bgpSrteV4PolicySlice = []BgpSrteV4Policy{}
	}
	return obj
}
func (obj *bgpV4PeerBgpSrteV4PolicyIter) clearHolderSlice() BgpV4PeerBgpSrteV4PolicyIter {
	if len(obj.bgpSrteV4PolicySlice) > 0 {
		obj.bgpSrteV4PolicySlice = []BgpSrteV4Policy{}
	}
	return obj
}
func (obj *bgpV4PeerBgpSrteV4PolicyIter) appendHolderSlice(item BgpSrteV4Policy) BgpV4PeerBgpSrteV4PolicyIter {
	obj.bgpSrteV4PolicySlice = append(obj.bgpSrteV4PolicySlice, item)
	return obj
}

// Segment Routing Traffic Engineering (SR TE) Policies for IPv6 Address Family Identifier (AFI).
// V6SrtePolicies returns a []BgpSrteV6Policy
func (obj *bgpV4Peer) V6SrtePolicies() BgpV4PeerBgpSrteV6PolicyIter {
	if len(obj.obj.V6SrtePolicies) == 0 {
		obj.obj.V6SrtePolicies = []*otg.BgpSrteV6Policy{}
	}
	if obj.v6SrtePoliciesHolder == nil {
		obj.v6SrtePoliciesHolder = newBgpV4PeerBgpSrteV6PolicyIter(&obj.obj.V6SrtePolicies).setMsg(obj)
	}
	return obj.v6SrtePoliciesHolder
}

type bgpV4PeerBgpSrteV6PolicyIter struct {
	obj                  *bgpV4Peer
	bgpSrteV6PolicySlice []BgpSrteV6Policy
	fieldPtr             *[]*otg.BgpSrteV6Policy
}

func newBgpV4PeerBgpSrteV6PolicyIter(ptr *[]*otg.BgpSrteV6Policy) BgpV4PeerBgpSrteV6PolicyIter {
	return &bgpV4PeerBgpSrteV6PolicyIter{fieldPtr: ptr}
}

type BgpV4PeerBgpSrteV6PolicyIter interface {
	setMsg(*bgpV4Peer) BgpV4PeerBgpSrteV6PolicyIter
	Items() []BgpSrteV6Policy
	Add() BgpSrteV6Policy
	Append(items ...BgpSrteV6Policy) BgpV4PeerBgpSrteV6PolicyIter
	Set(index int, newObj BgpSrteV6Policy) BgpV4PeerBgpSrteV6PolicyIter
	Clear() BgpV4PeerBgpSrteV6PolicyIter
	clearHolderSlice() BgpV4PeerBgpSrteV6PolicyIter
	appendHolderSlice(item BgpSrteV6Policy) BgpV4PeerBgpSrteV6PolicyIter
}

func (obj *bgpV4PeerBgpSrteV6PolicyIter) setMsg(msg *bgpV4Peer) BgpV4PeerBgpSrteV6PolicyIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpSrteV6Policy{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4PeerBgpSrteV6PolicyIter) Items() []BgpSrteV6Policy {
	return obj.bgpSrteV6PolicySlice
}

func (obj *bgpV4PeerBgpSrteV6PolicyIter) Add() BgpSrteV6Policy {
	newObj := &otg.BgpSrteV6Policy{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpSrteV6Policy{obj: newObj}
	newLibObj.setDefault()
	obj.bgpSrteV6PolicySlice = append(obj.bgpSrteV6PolicySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4PeerBgpSrteV6PolicyIter) Append(items ...BgpSrteV6Policy) BgpV4PeerBgpSrteV6PolicyIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpSrteV6PolicySlice = append(obj.bgpSrteV6PolicySlice, item)
	}
	return obj
}

func (obj *bgpV4PeerBgpSrteV6PolicyIter) Set(index int, newObj BgpSrteV6Policy) BgpV4PeerBgpSrteV6PolicyIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpSrteV6PolicySlice[index] = newObj
	return obj
}
func (obj *bgpV4PeerBgpSrteV6PolicyIter) Clear() BgpV4PeerBgpSrteV6PolicyIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpSrteV6Policy{}
		obj.bgpSrteV6PolicySlice = []BgpSrteV6Policy{}
	}
	return obj
}
func (obj *bgpV4PeerBgpSrteV6PolicyIter) clearHolderSlice() BgpV4PeerBgpSrteV6PolicyIter {
	if len(obj.bgpSrteV6PolicySlice) > 0 {
		obj.bgpSrteV6PolicySlice = []BgpSrteV6Policy{}
	}
	return obj
}
func (obj *bgpV4PeerBgpSrteV6PolicyIter) appendHolderSlice(item BgpSrteV6Policy) BgpV4PeerBgpSrteV6PolicyIter {
	obj.bgpSrteV6PolicySlice = append(obj.bgpSrteV6PolicySlice, item)
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpV4Peer) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpV4Peer object
func (obj *bgpV4Peer) SetName(value string) BgpV4Peer {

	obj.obj.Name = &value
	return obj
}

// description is TBD
// GracefulRestart returns a BgpGracefulRestart
func (obj *bgpV4Peer) GracefulRestart() BgpGracefulRestart {
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
func (obj *bgpV4Peer) HasGracefulRestart() bool {
	return obj.obj.GracefulRestart != nil
}

// description is TBD
// SetGracefulRestart sets the BgpGracefulRestart value in the BgpV4Peer object
func (obj *bgpV4Peer) SetGracefulRestart(value BgpGracefulRestart) BgpV4Peer {

	obj.gracefulRestartHolder = nil
	obj.obj.GracefulRestart = value.msg()

	return obj
}

// BGP Updates to be sent to the peer as specified after the session is established.
// ReplayUpdates returns a BgpUpdateReplay
func (obj *bgpV4Peer) ReplayUpdates() BgpUpdateReplay {
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
func (obj *bgpV4Peer) HasReplayUpdates() bool {
	return obj.obj.ReplayUpdates != nil
}

// BGP Updates to be sent to the peer as specified after the session is established.
// SetReplayUpdates sets the BgpUpdateReplay value in the BgpV4Peer object
func (obj *bgpV4Peer) SetReplayUpdates(value BgpUpdateReplay) BgpV4Peer {

	obj.replayUpdatesHolder = nil
	obj.obj.ReplayUpdates = value.msg()

	return obj
}

func (obj *bgpV4Peer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// PeerAddress is required
	if obj.obj.PeerAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "PeerAddress is required field on interface BgpV4Peer")
	}
	if obj.obj.PeerAddress != nil {

		err := obj.validateIpv4(obj.PeerAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpV4Peer.PeerAddress"))
		}

	}

	if len(obj.obj.EvpnEthernetSegments) != 0 {

		if set_default {
			obj.EvpnEthernetSegments().clearHolderSlice()
			for _, item := range obj.obj.EvpnEthernetSegments {
				obj.EvpnEthernetSegments().appendHolderSlice(&bgpV4EthernetSegment{obj: item})
			}
		}
		for _, item := range obj.EvpnEthernetSegments().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// AsType is required
	if obj.obj.AsType == nil {
		vObj.validationErrors = append(vObj.validationErrors, "AsType is required field on interface BgpV4Peer")
	}

	// AsNumber is required
	if obj.obj.AsNumber == nil {
		vObj.validationErrors = append(vObj.validationErrors, "AsNumber is required field on interface BgpV4Peer")
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
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpV4Peer")
	}

	if obj.obj.GracefulRestart != nil {

		obj.GracefulRestart().validateObj(vObj, set_default)
	}

	if obj.obj.ReplayUpdates != nil {

		obj.ReplayUpdates().validateObj(vObj, set_default)
	}

}

func (obj *bgpV4Peer) setDefault() {
	if obj.obj.AsNumberWidth == nil {
		obj.SetAsNumberWidth(BgpV4PeerAsNumberWidth.FOUR)

	}

}
