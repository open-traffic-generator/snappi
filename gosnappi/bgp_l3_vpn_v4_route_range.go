package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpL3VpnV4RouteRange *****
type bgpL3VpnV4RouteRange struct {
	validation
	obj                       *otg.BgpL3VpnV4RouteRange
	marshaller                marshalBgpL3VpnV4RouteRange
	unMarshaller              unMarshalBgpL3VpnV4RouteRange
	addressesHolder           BgpL3VpnV4RouteRangeV4RouteAddressIter
	advancedHolder            BgpRouteAdvanced
	communitiesHolder         BgpL3VpnV4RouteRangeBgpCommunityIter
	asPathHolder              BgpAsPath
	addPathHolder             BgpAddPath
	extendedCommunitiesHolder BgpL3VpnV4RouteRangeBgpExtendedCommunityIter
	mplsLabelsHolder          RouteMplsLabelValue
	routeDistinguisherHolder  BgpRouteDistinguisher
}

func NewBgpL3VpnV4RouteRange() BgpL3VpnV4RouteRange {
	obj := bgpL3VpnV4RouteRange{obj: &otg.BgpL3VpnV4RouteRange{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpL3VpnV4RouteRange) msg() *otg.BgpL3VpnV4RouteRange {
	return obj.obj
}

func (obj *bgpL3VpnV4RouteRange) setMsg(msg *otg.BgpL3VpnV4RouteRange) BgpL3VpnV4RouteRange {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpL3VpnV4RouteRange struct {
	obj *bgpL3VpnV4RouteRange
}

type marshalBgpL3VpnV4RouteRange interface {
	// ToProto marshals BgpL3VpnV4RouteRange to protobuf object *otg.BgpL3VpnV4RouteRange
	ToProto() (*otg.BgpL3VpnV4RouteRange, error)
	// ToPbText marshals BgpL3VpnV4RouteRange to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpL3VpnV4RouteRange to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpL3VpnV4RouteRange to JSON text
	ToJson() (string, error)
}

type unMarshalbgpL3VpnV4RouteRange struct {
	obj *bgpL3VpnV4RouteRange
}

type unMarshalBgpL3VpnV4RouteRange interface {
	// FromProto unmarshals BgpL3VpnV4RouteRange from protobuf object *otg.BgpL3VpnV4RouteRange
	FromProto(msg *otg.BgpL3VpnV4RouteRange) (BgpL3VpnV4RouteRange, error)
	// FromPbText unmarshals BgpL3VpnV4RouteRange from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpL3VpnV4RouteRange from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpL3VpnV4RouteRange from JSON text
	FromJson(value string) error
}

func (obj *bgpL3VpnV4RouteRange) Marshal() marshalBgpL3VpnV4RouteRange {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpL3VpnV4RouteRange{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpL3VpnV4RouteRange) Unmarshal() unMarshalBgpL3VpnV4RouteRange {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpL3VpnV4RouteRange{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpL3VpnV4RouteRange) ToProto() (*otg.BgpL3VpnV4RouteRange, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpL3VpnV4RouteRange) FromProto(msg *otg.BgpL3VpnV4RouteRange) (BgpL3VpnV4RouteRange, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpL3VpnV4RouteRange) ToPbText() (string, error) {
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

func (m *unMarshalbgpL3VpnV4RouteRange) FromPbText(value string) error {
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

func (m *marshalbgpL3VpnV4RouteRange) ToYaml() (string, error) {
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

func (m *unMarshalbgpL3VpnV4RouteRange) FromYaml(value string) error {
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

func (m *marshalbgpL3VpnV4RouteRange) ToJson() (string, error) {
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

func (m *unMarshalbgpL3VpnV4RouteRange) FromJson(value string) error {
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

func (obj *bgpL3VpnV4RouteRange) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpL3VpnV4RouteRange) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpL3VpnV4RouteRange) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpL3VpnV4RouteRange) Clone() (BgpL3VpnV4RouteRange, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpL3VpnV4RouteRange()
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

func (obj *bgpL3VpnV4RouteRange) setNil() {
	obj.addressesHolder = nil
	obj.advancedHolder = nil
	obj.communitiesHolder = nil
	obj.asPathHolder = nil
	obj.addPathHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.mplsLabelsHolder = nil
	obj.routeDistinguisherHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpL3VpnV4RouteRange is emulated VPN-IPv4 customer route range belonging to a Bgp.L3vpn.Vrf (RFC 4364). Same shape as the plain Bgp.V4RouteRange, plus a per route range route_distinguisher. The Route Distinguisher (RFC 4364 Section 4.1) is carried here, on the route range, rather than once on the parent VRF - matching RFC 4364's per-NLRI RD semantics. This lets different route ranges in the same VRF advertise different RDs (RFC 4364 Section 4.1 / 4.3.5, e.g. a multihomed CE whose prefix must stay distinct per PE). The VPN dataplane label is the existing mpls_labels field (RFC 4364 Section 3). The route range is kept independent of the shared Bgp.V4RouteRange so the VPN-only route_distinguisher does not leak into plain (non-VPN) BGP route ranges.
type BgpL3VpnV4RouteRange interface {
	Validation
	// msg marshals BgpL3VpnV4RouteRange to protobuf object *otg.BgpL3VpnV4RouteRange
	// and doesn't set defaults
	msg() *otg.BgpL3VpnV4RouteRange
	// setMsg unmarshals BgpL3VpnV4RouteRange from protobuf object *otg.BgpL3VpnV4RouteRange
	// and doesn't set defaults
	setMsg(*otg.BgpL3VpnV4RouteRange) BgpL3VpnV4RouteRange
	// provides marshal interface
	Marshal() marshalBgpL3VpnV4RouteRange
	// provides unmarshal interface
	Unmarshal() unMarshalBgpL3VpnV4RouteRange
	// validate validates BgpL3VpnV4RouteRange
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpL3VpnV4RouteRange, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Addresses returns BgpL3VpnV4RouteRangeV4RouteAddressIterIter, set in BgpL3VpnV4RouteRange
	Addresses() BgpL3VpnV4RouteRangeV4RouteAddressIter
	// NextHopMode returns BgpL3VpnV4RouteRangeNextHopModeEnum, set in BgpL3VpnV4RouteRange
	NextHopMode() BgpL3VpnV4RouteRangeNextHopModeEnum
	// SetNextHopMode assigns BgpL3VpnV4RouteRangeNextHopModeEnum provided by user to BgpL3VpnV4RouteRange
	SetNextHopMode(value BgpL3VpnV4RouteRangeNextHopModeEnum) BgpL3VpnV4RouteRange
	// HasNextHopMode checks if NextHopMode has been set in BgpL3VpnV4RouteRange
	HasNextHopMode() bool
	// NextHopAddressType returns BgpL3VpnV4RouteRangeNextHopAddressTypeEnum, set in BgpL3VpnV4RouteRange
	NextHopAddressType() BgpL3VpnV4RouteRangeNextHopAddressTypeEnum
	// SetNextHopAddressType assigns BgpL3VpnV4RouteRangeNextHopAddressTypeEnum provided by user to BgpL3VpnV4RouteRange
	SetNextHopAddressType(value BgpL3VpnV4RouteRangeNextHopAddressTypeEnum) BgpL3VpnV4RouteRange
	// HasNextHopAddressType checks if NextHopAddressType has been set in BgpL3VpnV4RouteRange
	HasNextHopAddressType() bool
	// NextHopIpv4Address returns string, set in BgpL3VpnV4RouteRange.
	NextHopIpv4Address() string
	// SetNextHopIpv4Address assigns string provided by user to BgpL3VpnV4RouteRange
	SetNextHopIpv4Address(value string) BgpL3VpnV4RouteRange
	// HasNextHopIpv4Address checks if NextHopIpv4Address has been set in BgpL3VpnV4RouteRange
	HasNextHopIpv4Address() bool
	// NextHopIpv6Address returns string, set in BgpL3VpnV4RouteRange.
	NextHopIpv6Address() string
	// SetNextHopIpv6Address assigns string provided by user to BgpL3VpnV4RouteRange
	SetNextHopIpv6Address(value string) BgpL3VpnV4RouteRange
	// HasNextHopIpv6Address checks if NextHopIpv6Address has been set in BgpL3VpnV4RouteRange
	HasNextHopIpv6Address() bool
	// Advanced returns BgpRouteAdvanced, set in BgpL3VpnV4RouteRange.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	Advanced() BgpRouteAdvanced
	// SetAdvanced assigns BgpRouteAdvanced provided by user to BgpL3VpnV4RouteRange.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	SetAdvanced(value BgpRouteAdvanced) BgpL3VpnV4RouteRange
	// HasAdvanced checks if Advanced has been set in BgpL3VpnV4RouteRange
	HasAdvanced() bool
	// Communities returns BgpL3VpnV4RouteRangeBgpCommunityIterIter, set in BgpL3VpnV4RouteRange
	Communities() BgpL3VpnV4RouteRangeBgpCommunityIter
	// AsPath returns BgpAsPath, set in BgpL3VpnV4RouteRange.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	AsPath() BgpAsPath
	// SetAsPath assigns BgpAsPath provided by user to BgpL3VpnV4RouteRange.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	SetAsPath(value BgpAsPath) BgpL3VpnV4RouteRange
	// HasAsPath checks if AsPath has been set in BgpL3VpnV4RouteRange
	HasAsPath() bool
	// AddPath returns BgpAddPath, set in BgpL3VpnV4RouteRange.
	// BgpAddPath is the BGP Additional Paths feature is a BGP extension that allows the  advertisement of multiple paths for the same prefix without the new  paths implicitly replacing any previous paths.
	AddPath() BgpAddPath
	// SetAddPath assigns BgpAddPath provided by user to BgpL3VpnV4RouteRange.
	// BgpAddPath is the BGP Additional Paths feature is a BGP extension that allows the  advertisement of multiple paths for the same prefix without the new  paths implicitly replacing any previous paths.
	SetAddPath(value BgpAddPath) BgpL3VpnV4RouteRange
	// HasAddPath checks if AddPath has been set in BgpL3VpnV4RouteRange
	HasAddPath() bool
	// Name returns string, set in BgpL3VpnV4RouteRange.
	Name() string
	// SetName assigns string provided by user to BgpL3VpnV4RouteRange
	SetName(value string) BgpL3VpnV4RouteRange
	// ExtendedCommunities returns BgpL3VpnV4RouteRangeBgpExtendedCommunityIterIter, set in BgpL3VpnV4RouteRange
	ExtendedCommunities() BgpL3VpnV4RouteRangeBgpExtendedCommunityIter
	// MplsLabels returns RouteMplsLabelValue, set in BgpL3VpnV4RouteRange.
	// RouteMplsLabelValue is a container of MPLS Prefix Label Value/Index in the address range.
	MplsLabels() RouteMplsLabelValue
	// SetMplsLabels assigns RouteMplsLabelValue provided by user to BgpL3VpnV4RouteRange.
	// RouteMplsLabelValue is a container of MPLS Prefix Label Value/Index in the address range.
	SetMplsLabels(value RouteMplsLabelValue) BgpL3VpnV4RouteRange
	// RouteDistinguisher returns BgpRouteDistinguisher, set in BgpL3VpnV4RouteRange.
	// BgpRouteDistinguisher is bGP Route Distinguisher.
	RouteDistinguisher() BgpRouteDistinguisher
	// SetRouteDistinguisher assigns BgpRouteDistinguisher provided by user to BgpL3VpnV4RouteRange.
	// BgpRouteDistinguisher is bGP Route Distinguisher.
	SetRouteDistinguisher(value BgpRouteDistinguisher) BgpL3VpnV4RouteRange
	setNil()
}

// A list of group of IPv4 route addresses.
// Addresses returns a []V4RouteAddress
func (obj *bgpL3VpnV4RouteRange) Addresses() BgpL3VpnV4RouteRangeV4RouteAddressIter {
	if len(obj.obj.Addresses) == 0 {
		obj.obj.Addresses = []*otg.V4RouteAddress{}
	}
	if obj.addressesHolder == nil {
		obj.addressesHolder = newBgpL3VpnV4RouteRangeV4RouteAddressIter(&obj.obj.Addresses).setMsg(obj)
	}
	return obj.addressesHolder
}

type bgpL3VpnV4RouteRangeV4RouteAddressIter struct {
	obj                 *bgpL3VpnV4RouteRange
	v4RouteAddressSlice []V4RouteAddress
	fieldPtr            *[]*otg.V4RouteAddress
}

func newBgpL3VpnV4RouteRangeV4RouteAddressIter(ptr *[]*otg.V4RouteAddress) BgpL3VpnV4RouteRangeV4RouteAddressIter {
	return &bgpL3VpnV4RouteRangeV4RouteAddressIter{fieldPtr: ptr}
}

type BgpL3VpnV4RouteRangeV4RouteAddressIter interface {
	setMsg(*bgpL3VpnV4RouteRange) BgpL3VpnV4RouteRangeV4RouteAddressIter
	Items() []V4RouteAddress
	Add() V4RouteAddress
	Append(items ...V4RouteAddress) BgpL3VpnV4RouteRangeV4RouteAddressIter
	Set(index int, newObj V4RouteAddress) BgpL3VpnV4RouteRangeV4RouteAddressIter
	Clear() BgpL3VpnV4RouteRangeV4RouteAddressIter
	clearHolderSlice() BgpL3VpnV4RouteRangeV4RouteAddressIter
	appendHolderSlice(item V4RouteAddress) BgpL3VpnV4RouteRangeV4RouteAddressIter
}

func (obj *bgpL3VpnV4RouteRangeV4RouteAddressIter) setMsg(msg *bgpL3VpnV4RouteRange) BgpL3VpnV4RouteRangeV4RouteAddressIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&v4RouteAddress{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpL3VpnV4RouteRangeV4RouteAddressIter) Items() []V4RouteAddress {
	return obj.v4RouteAddressSlice
}

func (obj *bgpL3VpnV4RouteRangeV4RouteAddressIter) Add() V4RouteAddress {
	newObj := &otg.V4RouteAddress{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &v4RouteAddress{obj: newObj}
	newLibObj.setDefault()
	obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, newLibObj)
	return newLibObj
}

func (obj *bgpL3VpnV4RouteRangeV4RouteAddressIter) Append(items ...V4RouteAddress) BgpL3VpnV4RouteRangeV4RouteAddressIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, item)
	}
	return obj
}

func (obj *bgpL3VpnV4RouteRangeV4RouteAddressIter) Set(index int, newObj V4RouteAddress) BgpL3VpnV4RouteRangeV4RouteAddressIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.v4RouteAddressSlice[index] = newObj
	return obj
}
func (obj *bgpL3VpnV4RouteRangeV4RouteAddressIter) Clear() BgpL3VpnV4RouteRangeV4RouteAddressIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.V4RouteAddress{}
		obj.v4RouteAddressSlice = []V4RouteAddress{}
	}
	return obj
}
func (obj *bgpL3VpnV4RouteRangeV4RouteAddressIter) clearHolderSlice() BgpL3VpnV4RouteRangeV4RouteAddressIter {
	if len(obj.v4RouteAddressSlice) > 0 {
		obj.v4RouteAddressSlice = []V4RouteAddress{}
	}
	return obj
}
func (obj *bgpL3VpnV4RouteRangeV4RouteAddressIter) appendHolderSlice(item V4RouteAddress) BgpL3VpnV4RouteRangeV4RouteAddressIter {
	obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, item)
	return obj
}

type BgpL3VpnV4RouteRangeNextHopModeEnum string

// Enum of NextHopMode on BgpL3VpnV4RouteRange
var BgpL3VpnV4RouteRangeNextHopMode = struct {
	LOCAL_IP BgpL3VpnV4RouteRangeNextHopModeEnum
	MANUAL   BgpL3VpnV4RouteRangeNextHopModeEnum
}{
	LOCAL_IP: BgpL3VpnV4RouteRangeNextHopModeEnum("local_ip"),
	MANUAL:   BgpL3VpnV4RouteRangeNextHopModeEnum("manual"),
}

func (obj *bgpL3VpnV4RouteRange) NextHopMode() BgpL3VpnV4RouteRangeNextHopModeEnum {
	return BgpL3VpnV4RouteRangeNextHopModeEnum(obj.obj.NextHopMode.Enum().String())
}

// Specify the NextHop in MP REACH NLRI. The mode for setting the IP address  of the NextHop in the MP REACH NLRI can be one of the following:
// Local IP: Automatically fills the Nexthop with the Local IP of the BGP
// peer.
// If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// Manual: Override the Nexthop with any arbitrary IPv4/IPv6 address.
// NextHopMode returns a string
func (obj *bgpL3VpnV4RouteRange) HasNextHopMode() bool {
	return obj.obj.NextHopMode != nil
}

func (obj *bgpL3VpnV4RouteRange) SetNextHopMode(value BgpL3VpnV4RouteRangeNextHopModeEnum) BgpL3VpnV4RouteRange {
	intValue, ok := otg.BgpL3VpnV4RouteRange_NextHopMode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpL3VpnV4RouteRangeNextHopModeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpL3VpnV4RouteRange_NextHopMode_Enum(intValue)
	obj.obj.NextHopMode = &enumValue

	return obj
}

type BgpL3VpnV4RouteRangeNextHopAddressTypeEnum string

// Enum of NextHopAddressType on BgpL3VpnV4RouteRange
var BgpL3VpnV4RouteRangeNextHopAddressType = struct {
	IPV4 BgpL3VpnV4RouteRangeNextHopAddressTypeEnum
	IPV6 BgpL3VpnV4RouteRangeNextHopAddressTypeEnum
}{
	IPV4: BgpL3VpnV4RouteRangeNextHopAddressTypeEnum("ipv4"),
	IPV6: BgpL3VpnV4RouteRangeNextHopAddressTypeEnum("ipv6"),
}

func (obj *bgpL3VpnV4RouteRange) NextHopAddressType() BgpL3VpnV4RouteRangeNextHopAddressTypeEnum {
	return BgpL3VpnV4RouteRangeNextHopAddressTypeEnum(obj.obj.NextHopAddressType.Enum().String())
}

// If the Nexthop Mode is Manual, it sets the type of the NextHop IP address.
// NextHopAddressType returns a string
func (obj *bgpL3VpnV4RouteRange) HasNextHopAddressType() bool {
	return obj.obj.NextHopAddressType != nil
}

func (obj *bgpL3VpnV4RouteRange) SetNextHopAddressType(value BgpL3VpnV4RouteRangeNextHopAddressTypeEnum) BgpL3VpnV4RouteRange {
	intValue, ok := otg.BgpL3VpnV4RouteRange_NextHopAddressType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpL3VpnV4RouteRangeNextHopAddressTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpL3VpnV4RouteRange_NextHopAddressType_Enum(intValue)
	obj.obj.NextHopAddressType = &enumValue

	return obj
}

// The IPv4 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// NextHopIpv4Address returns a string
func (obj *bgpL3VpnV4RouteRange) NextHopIpv4Address() string {

	return *obj.obj.NextHopIpv4Address

}

// The IPv4 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// NextHopIpv4Address returns a string
func (obj *bgpL3VpnV4RouteRange) HasNextHopIpv4Address() bool {
	return obj.obj.NextHopIpv4Address != nil
}

// The IPv4 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// SetNextHopIpv4Address sets the string value in the BgpL3VpnV4RouteRange object
func (obj *bgpL3VpnV4RouteRange) SetNextHopIpv4Address(value string) BgpL3VpnV4RouteRange {

	obj.obj.NextHopIpv4Address = &value
	return obj
}

// The IPv6 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv6.
// NextHopIpv6Address returns a string
func (obj *bgpL3VpnV4RouteRange) NextHopIpv6Address() string {

	return *obj.obj.NextHopIpv6Address

}

// The IPv6 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv6.
// NextHopIpv6Address returns a string
func (obj *bgpL3VpnV4RouteRange) HasNextHopIpv6Address() bool {
	return obj.obj.NextHopIpv6Address != nil
}

// The IPv6 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv6.
// SetNextHopIpv6Address sets the string value in the BgpL3VpnV4RouteRange object
func (obj *bgpL3VpnV4RouteRange) SetNextHopIpv6Address(value string) BgpL3VpnV4RouteRange {

	obj.obj.NextHopIpv6Address = &value
	return obj
}

// description is TBD
// Advanced returns a BgpRouteAdvanced
func (obj *bgpL3VpnV4RouteRange) Advanced() BgpRouteAdvanced {
	if obj.obj.Advanced == nil {
		obj.obj.Advanced = NewBgpRouteAdvanced().msg()
	}
	if obj.advancedHolder == nil {
		obj.advancedHolder = &bgpRouteAdvanced{obj: obj.obj.Advanced}
	}
	return obj.advancedHolder
}

// description is TBD
// Advanced returns a BgpRouteAdvanced
func (obj *bgpL3VpnV4RouteRange) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// description is TBD
// SetAdvanced sets the BgpRouteAdvanced value in the BgpL3VpnV4RouteRange object
func (obj *bgpL3VpnV4RouteRange) SetAdvanced(value BgpRouteAdvanced) BgpL3VpnV4RouteRange {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// Optional community settings.
// Communities returns a []BgpCommunity
func (obj *bgpL3VpnV4RouteRange) Communities() BgpL3VpnV4RouteRangeBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.BgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpL3VpnV4RouteRangeBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpL3VpnV4RouteRangeBgpCommunityIter struct {
	obj               *bgpL3VpnV4RouteRange
	bgpCommunitySlice []BgpCommunity
	fieldPtr          *[]*otg.BgpCommunity
}

func newBgpL3VpnV4RouteRangeBgpCommunityIter(ptr *[]*otg.BgpCommunity) BgpL3VpnV4RouteRangeBgpCommunityIter {
	return &bgpL3VpnV4RouteRangeBgpCommunityIter{fieldPtr: ptr}
}

type BgpL3VpnV4RouteRangeBgpCommunityIter interface {
	setMsg(*bgpL3VpnV4RouteRange) BgpL3VpnV4RouteRangeBgpCommunityIter
	Items() []BgpCommunity
	Add() BgpCommunity
	Append(items ...BgpCommunity) BgpL3VpnV4RouteRangeBgpCommunityIter
	Set(index int, newObj BgpCommunity) BgpL3VpnV4RouteRangeBgpCommunityIter
	Clear() BgpL3VpnV4RouteRangeBgpCommunityIter
	clearHolderSlice() BgpL3VpnV4RouteRangeBgpCommunityIter
	appendHolderSlice(item BgpCommunity) BgpL3VpnV4RouteRangeBgpCommunityIter
}

func (obj *bgpL3VpnV4RouteRangeBgpCommunityIter) setMsg(msg *bgpL3VpnV4RouteRange) BgpL3VpnV4RouteRangeBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpL3VpnV4RouteRangeBgpCommunityIter) Items() []BgpCommunity {
	return obj.bgpCommunitySlice
}

func (obj *bgpL3VpnV4RouteRangeBgpCommunityIter) Add() BgpCommunity {
	newObj := &otg.BgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpL3VpnV4RouteRangeBgpCommunityIter) Append(items ...BgpCommunity) BgpL3VpnV4RouteRangeBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpL3VpnV4RouteRangeBgpCommunityIter) Set(index int, newObj BgpCommunity) BgpL3VpnV4RouteRangeBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpL3VpnV4RouteRangeBgpCommunityIter) Clear() BgpL3VpnV4RouteRangeBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpCommunity{}
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpL3VpnV4RouteRangeBgpCommunityIter) clearHolderSlice() BgpL3VpnV4RouteRangeBgpCommunityIter {
	if len(obj.bgpCommunitySlice) > 0 {
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpL3VpnV4RouteRangeBgpCommunityIter) appendHolderSlice(item BgpCommunity) BgpL3VpnV4RouteRangeBgpCommunityIter {
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a BgpAsPath
func (obj *bgpL3VpnV4RouteRange) AsPath() BgpAsPath {
	if obj.obj.AsPath == nil {
		obj.obj.AsPath = NewBgpAsPath().msg()
	}
	if obj.asPathHolder == nil {
		obj.asPathHolder = &bgpAsPath{obj: obj.obj.AsPath}
	}
	return obj.asPathHolder
}

// description is TBD
// AsPath returns a BgpAsPath
func (obj *bgpL3VpnV4RouteRange) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the BgpAsPath value in the BgpL3VpnV4RouteRange object
func (obj *bgpL3VpnV4RouteRange) SetAsPath(value BgpAsPath) BgpL3VpnV4RouteRange {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// description is TBD
// AddPath returns a BgpAddPath
func (obj *bgpL3VpnV4RouteRange) AddPath() BgpAddPath {
	if obj.obj.AddPath == nil {
		obj.obj.AddPath = NewBgpAddPath().msg()
	}
	if obj.addPathHolder == nil {
		obj.addPathHolder = &bgpAddPath{obj: obj.obj.AddPath}
	}
	return obj.addPathHolder
}

// description is TBD
// AddPath returns a BgpAddPath
func (obj *bgpL3VpnV4RouteRange) HasAddPath() bool {
	return obj.obj.AddPath != nil
}

// description is TBD
// SetAddPath sets the BgpAddPath value in the BgpL3VpnV4RouteRange object
func (obj *bgpL3VpnV4RouteRange) SetAddPath(value BgpAddPath) BgpL3VpnV4RouteRange {

	obj.addPathHolder = nil
	obj.obj.AddPath = value.msg()

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpL3VpnV4RouteRange) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpL3VpnV4RouteRange object
func (obj *bgpL3VpnV4RouteRange) SetName(value string) BgpL3VpnV4RouteRange {

	obj.obj.Name = &value
	return obj
}

// Optional Extended Community settings. The Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an eight byte value. It is divided into two main parts. The first two bytes of the community encode a type and sub-type fields and the last six bytes carry a unique set of data in a format defined by the type and sub-type field. Extended communities provide a larger range for grouping or categorizing communities.
// ExtendedCommunities returns a []BgpExtendedCommunity
func (obj *bgpL3VpnV4RouteRange) ExtendedCommunities() BgpL3VpnV4RouteRangeBgpExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.BgpExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpL3VpnV4RouteRangeBgpExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpL3VpnV4RouteRangeBgpExtendedCommunityIter struct {
	obj                       *bgpL3VpnV4RouteRange
	bgpExtendedCommunitySlice []BgpExtendedCommunity
	fieldPtr                  *[]*otg.BgpExtendedCommunity
}

func newBgpL3VpnV4RouteRangeBgpExtendedCommunityIter(ptr *[]*otg.BgpExtendedCommunity) BgpL3VpnV4RouteRangeBgpExtendedCommunityIter {
	return &bgpL3VpnV4RouteRangeBgpExtendedCommunityIter{fieldPtr: ptr}
}

type BgpL3VpnV4RouteRangeBgpExtendedCommunityIter interface {
	setMsg(*bgpL3VpnV4RouteRange) BgpL3VpnV4RouteRangeBgpExtendedCommunityIter
	Items() []BgpExtendedCommunity
	Add() BgpExtendedCommunity
	Append(items ...BgpExtendedCommunity) BgpL3VpnV4RouteRangeBgpExtendedCommunityIter
	Set(index int, newObj BgpExtendedCommunity) BgpL3VpnV4RouteRangeBgpExtendedCommunityIter
	Clear() BgpL3VpnV4RouteRangeBgpExtendedCommunityIter
	clearHolderSlice() BgpL3VpnV4RouteRangeBgpExtendedCommunityIter
	appendHolderSlice(item BgpExtendedCommunity) BgpL3VpnV4RouteRangeBgpExtendedCommunityIter
}

func (obj *bgpL3VpnV4RouteRangeBgpExtendedCommunityIter) setMsg(msg *bgpL3VpnV4RouteRange) BgpL3VpnV4RouteRangeBgpExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpL3VpnV4RouteRangeBgpExtendedCommunityIter) Items() []BgpExtendedCommunity {
	return obj.bgpExtendedCommunitySlice
}

func (obj *bgpL3VpnV4RouteRangeBgpExtendedCommunityIter) Add() BgpExtendedCommunity {
	newObj := &otg.BgpExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpL3VpnV4RouteRangeBgpExtendedCommunityIter) Append(items ...BgpExtendedCommunity) BgpL3VpnV4RouteRangeBgpExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpL3VpnV4RouteRangeBgpExtendedCommunityIter) Set(index int, newObj BgpExtendedCommunity) BgpL3VpnV4RouteRangeBgpExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpL3VpnV4RouteRangeBgpExtendedCommunityIter) Clear() BgpL3VpnV4RouteRangeBgpExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtendedCommunity{}
		obj.bgpExtendedCommunitySlice = []BgpExtendedCommunity{}
	}
	return obj
}
func (obj *bgpL3VpnV4RouteRangeBgpExtendedCommunityIter) clearHolderSlice() BgpL3VpnV4RouteRangeBgpExtendedCommunityIter {
	if len(obj.bgpExtendedCommunitySlice) > 0 {
		obj.bgpExtendedCommunitySlice = []BgpExtendedCommunity{}
	}
	return obj
}
func (obj *bgpL3VpnV4RouteRangeBgpExtendedCommunityIter) appendHolderSlice(item BgpExtendedCommunity) BgpL3VpnV4RouteRangeBgpExtendedCommunityIter {
	obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, item)
	return obj
}

// The single VPN dataplane MPLS label (RFC 4364 Section 3) bound to every IPv4 prefix in this route range. RFC 4364 assigns exactly one label per route - Section 4.3.2 only discusses whether that value is shared across routes in a VRF or attachment circuit, not stacking multiple labels per route.
// MplsLabels returns a RouteMplsLabelValue
func (obj *bgpL3VpnV4RouteRange) MplsLabels() RouteMplsLabelValue {
	if obj.obj.MplsLabels == nil {
		obj.obj.MplsLabels = NewRouteMplsLabelValue().msg()
	}
	if obj.mplsLabelsHolder == nil {
		obj.mplsLabelsHolder = &routeMplsLabelValue{obj: obj.obj.MplsLabels}
	}
	return obj.mplsLabelsHolder
}

// The single VPN dataplane MPLS label (RFC 4364 Section 3) bound to every IPv4 prefix in this route range. RFC 4364 assigns exactly one label per route - Section 4.3.2 only discusses whether that value is shared across routes in a VRF or attachment circuit, not stacking multiple labels per route.
// SetMplsLabels sets the RouteMplsLabelValue value in the BgpL3VpnV4RouteRange object
func (obj *bgpL3VpnV4RouteRange) SetMplsLabels(value RouteMplsLabelValue) BgpL3VpnV4RouteRange {

	obj.mplsLabelsHolder = nil
	obj.obj.MplsLabels = value.msg()

	return obj
}

// The Route Distinguisher (RFC 4364 Section 4.1) prepended to every IPv4 prefix in this route range, forming the VPN-IPv4 NLRI (AFI 1, SAFI 128). Carried per route range to match RFC 4364's per-NLRI RD semantics; route ranges in the same VRF may use different RDs.
// RouteDistinguisher returns a BgpRouteDistinguisher
func (obj *bgpL3VpnV4RouteRange) RouteDistinguisher() BgpRouteDistinguisher {
	if obj.obj.RouteDistinguisher == nil {
		obj.obj.RouteDistinguisher = NewBgpRouteDistinguisher().msg()
	}
	if obj.routeDistinguisherHolder == nil {
		obj.routeDistinguisherHolder = &bgpRouteDistinguisher{obj: obj.obj.RouteDistinguisher}
	}
	return obj.routeDistinguisherHolder
}

// The Route Distinguisher (RFC 4364 Section 4.1) prepended to every IPv4 prefix in this route range, forming the VPN-IPv4 NLRI (AFI 1, SAFI 128). Carried per route range to match RFC 4364's per-NLRI RD semantics; route ranges in the same VRF may use different RDs.
// SetRouteDistinguisher sets the BgpRouteDistinguisher value in the BgpL3VpnV4RouteRange object
func (obj *bgpL3VpnV4RouteRange) SetRouteDistinguisher(value BgpRouteDistinguisher) BgpL3VpnV4RouteRange {

	obj.routeDistinguisherHolder = nil
	obj.obj.RouteDistinguisher = value.msg()

	return obj
}

func (obj *bgpL3VpnV4RouteRange) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Addresses) != 0 {

		if set_default {
			obj.Addresses().clearHolderSlice()
			for _, item := range obj.obj.Addresses {
				obj.Addresses().appendHolderSlice(&v4RouteAddress{obj: item})
			}
		}
		for _, item := range obj.Addresses().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.NextHopIpv4Address != nil {

		err := obj.validateIpv4(obj.NextHopIpv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpL3VpnV4RouteRange.NextHopIpv4Address"))
		}

	}

	if obj.obj.NextHopIpv6Address != nil {

		err := obj.validateIpv6(obj.NextHopIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpL3VpnV4RouteRange.NextHopIpv6Address"))
		}

	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(vObj, set_default)
	}

	if len(obj.obj.Communities) != 0 {

		if set_default {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				obj.Communities().appendHolderSlice(&bgpCommunity{obj: item})
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(vObj, set_default)
	}

	if obj.obj.AddPath != nil {

		obj.AddPath().validateObj(vObj, set_default)
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpL3VpnV4RouteRange")
	}

	if len(obj.obj.ExtendedCommunities) != 0 {

		if set_default {
			obj.ExtendedCommunities().clearHolderSlice()
			for _, item := range obj.obj.ExtendedCommunities {
				obj.ExtendedCommunities().appendHolderSlice(&bgpExtendedCommunity{obj: item})
			}
		}
		for _, item := range obj.ExtendedCommunities().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// MplsLabels is required
	if obj.obj.MplsLabels == nil {
		vObj.validationErrors = append(vObj.validationErrors, "MplsLabels is required field on interface BgpL3VpnV4RouteRange")
	}

	if obj.obj.MplsLabels != nil {

		obj.MplsLabels().validateObj(vObj, set_default)
	}

	// RouteDistinguisher is required
	if obj.obj.RouteDistinguisher == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RouteDistinguisher is required field on interface BgpL3VpnV4RouteRange")
	}

	if obj.obj.RouteDistinguisher != nil {

		obj.RouteDistinguisher().validateObj(vObj, set_default)
	}

}

func (obj *bgpL3VpnV4RouteRange) setDefault() {
	if obj.obj.NextHopMode == nil {
		obj.SetNextHopMode(BgpL3VpnV4RouteRangeNextHopMode.LOCAL_IP)

	}
	if obj.obj.NextHopAddressType == nil {
		obj.SetNextHopAddressType(BgpL3VpnV4RouteRangeNextHopAddressType.IPV4)

	}
	if obj.obj.NextHopIpv4Address == nil {
		obj.SetNextHopIpv4Address("0.0.0.0")
	}
	if obj.obj.NextHopIpv6Address == nil {
		obj.SetNextHopIpv6Address("::0")
	}

}
