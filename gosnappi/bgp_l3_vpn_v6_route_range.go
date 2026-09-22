package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpL3VpnV6RouteRange *****
type bgpL3VpnV6RouteRange struct {
	validation
	obj                       *otg.BgpL3VpnV6RouteRange
	marshaller                marshalBgpL3VpnV6RouteRange
	unMarshaller              unMarshalBgpL3VpnV6RouteRange
	addressesHolder           BgpL3VpnV6RouteRangeV6RouteAddressIter
	advancedHolder            BgpRouteAdvanced
	communitiesHolder         BgpL3VpnV6RouteRangeBgpCommunityIter
	asPathHolder              BgpAsPath
	addPathHolder             BgpAddPath
	extendedCommunitiesHolder BgpL3VpnV6RouteRangeBgpExtendedCommunityIter
	serviceBindingHolder      BgpL3VpnV6ServiceBinding
}

func NewBgpL3VpnV6RouteRange() BgpL3VpnV6RouteRange {
	obj := bgpL3VpnV6RouteRange{obj: &otg.BgpL3VpnV6RouteRange{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpL3VpnV6RouteRange) msg() *otg.BgpL3VpnV6RouteRange {
	return obj.obj
}

func (obj *bgpL3VpnV6RouteRange) setMsg(msg *otg.BgpL3VpnV6RouteRange) BgpL3VpnV6RouteRange {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpL3VpnV6RouteRange struct {
	obj *bgpL3VpnV6RouteRange
}

type marshalBgpL3VpnV6RouteRange interface {
	// ToProto marshals BgpL3VpnV6RouteRange to protobuf object *otg.BgpL3VpnV6RouteRange
	ToProto() (*otg.BgpL3VpnV6RouteRange, error)
	// ToPbText marshals BgpL3VpnV6RouteRange to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpL3VpnV6RouteRange to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpL3VpnV6RouteRange to JSON text
	ToJson() (string, error)
}

type unMarshalbgpL3VpnV6RouteRange struct {
	obj *bgpL3VpnV6RouteRange
}

type unMarshalBgpL3VpnV6RouteRange interface {
	// FromProto unmarshals BgpL3VpnV6RouteRange from protobuf object *otg.BgpL3VpnV6RouteRange
	FromProto(msg *otg.BgpL3VpnV6RouteRange) (BgpL3VpnV6RouteRange, error)
	// FromPbText unmarshals BgpL3VpnV6RouteRange from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpL3VpnV6RouteRange from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpL3VpnV6RouteRange from JSON text
	FromJson(value string) error
}

func (obj *bgpL3VpnV6RouteRange) Marshal() marshalBgpL3VpnV6RouteRange {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpL3VpnV6RouteRange{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpL3VpnV6RouteRange) Unmarshal() unMarshalBgpL3VpnV6RouteRange {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpL3VpnV6RouteRange{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpL3VpnV6RouteRange) ToProto() (*otg.BgpL3VpnV6RouteRange, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpL3VpnV6RouteRange) FromProto(msg *otg.BgpL3VpnV6RouteRange) (BgpL3VpnV6RouteRange, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpL3VpnV6RouteRange) ToPbText() (string, error) {
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

func (m *unMarshalbgpL3VpnV6RouteRange) FromPbText(value string) error {
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

func (m *marshalbgpL3VpnV6RouteRange) ToYaml() (string, error) {
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

func (m *unMarshalbgpL3VpnV6RouteRange) FromYaml(value string) error {
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

func (m *marshalbgpL3VpnV6RouteRange) ToJson() (string, error) {
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

func (m *unMarshalbgpL3VpnV6RouteRange) FromJson(value string) error {
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

func (obj *bgpL3VpnV6RouteRange) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpL3VpnV6RouteRange) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpL3VpnV6RouteRange) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpL3VpnV6RouteRange) Clone() (BgpL3VpnV6RouteRange, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpL3VpnV6RouteRange()
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

func (obj *bgpL3VpnV6RouteRange) setNil() {
	obj.addressesHolder = nil
	obj.advancedHolder = nil
	obj.communitiesHolder = nil
	obj.asPathHolder = nil
	obj.addPathHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.serviceBindingHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpL3VpnV6RouteRange is emulated VPN-IPv6 customer route range belonging to a Bgp.L3vpn.Vrf (RFC 4659, 6VPE). Same shape as the plain Bgp.V6RouteRange, except the dataplane binding is selected via service_binding, a choice that currently offers only a VPN MPLS label (RFC 4364 Section 3) but keeps the route range independent of the shared Bgp.V6RouteRange/ Bgp.MplsLabelBindings schemas so that an additional dataplane binding (for example an SRv6 Service SID, RFC 9252) can be added later as a new choice value without a breaking change.
type BgpL3VpnV6RouteRange interface {
	Validation
	// msg marshals BgpL3VpnV6RouteRange to protobuf object *otg.BgpL3VpnV6RouteRange
	// and doesn't set defaults
	msg() *otg.BgpL3VpnV6RouteRange
	// setMsg unmarshals BgpL3VpnV6RouteRange from protobuf object *otg.BgpL3VpnV6RouteRange
	// and doesn't set defaults
	setMsg(*otg.BgpL3VpnV6RouteRange) BgpL3VpnV6RouteRange
	// provides marshal interface
	Marshal() marshalBgpL3VpnV6RouteRange
	// provides unmarshal interface
	Unmarshal() unMarshalBgpL3VpnV6RouteRange
	// validate validates BgpL3VpnV6RouteRange
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpL3VpnV6RouteRange, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Addresses returns BgpL3VpnV6RouteRangeV6RouteAddressIterIter, set in BgpL3VpnV6RouteRange
	Addresses() BgpL3VpnV6RouteRangeV6RouteAddressIter
	// NextHopMode returns BgpL3VpnV6RouteRangeNextHopModeEnum, set in BgpL3VpnV6RouteRange
	NextHopMode() BgpL3VpnV6RouteRangeNextHopModeEnum
	// SetNextHopMode assigns BgpL3VpnV6RouteRangeNextHopModeEnum provided by user to BgpL3VpnV6RouteRange
	SetNextHopMode(value BgpL3VpnV6RouteRangeNextHopModeEnum) BgpL3VpnV6RouteRange
	// HasNextHopMode checks if NextHopMode has been set in BgpL3VpnV6RouteRange
	HasNextHopMode() bool
	// NextHopAddressType returns BgpL3VpnV6RouteRangeNextHopAddressTypeEnum, set in BgpL3VpnV6RouteRange
	NextHopAddressType() BgpL3VpnV6RouteRangeNextHopAddressTypeEnum
	// SetNextHopAddressType assigns BgpL3VpnV6RouteRangeNextHopAddressTypeEnum provided by user to BgpL3VpnV6RouteRange
	SetNextHopAddressType(value BgpL3VpnV6RouteRangeNextHopAddressTypeEnum) BgpL3VpnV6RouteRange
	// HasNextHopAddressType checks if NextHopAddressType has been set in BgpL3VpnV6RouteRange
	HasNextHopAddressType() bool
	// NextHopIpv4Address returns string, set in BgpL3VpnV6RouteRange.
	NextHopIpv4Address() string
	// SetNextHopIpv4Address assigns string provided by user to BgpL3VpnV6RouteRange
	SetNextHopIpv4Address(value string) BgpL3VpnV6RouteRange
	// HasNextHopIpv4Address checks if NextHopIpv4Address has been set in BgpL3VpnV6RouteRange
	HasNextHopIpv4Address() bool
	// NextHopIpv6Address returns string, set in BgpL3VpnV6RouteRange.
	NextHopIpv6Address() string
	// SetNextHopIpv6Address assigns string provided by user to BgpL3VpnV6RouteRange
	SetNextHopIpv6Address(value string) BgpL3VpnV6RouteRange
	// HasNextHopIpv6Address checks if NextHopIpv6Address has been set in BgpL3VpnV6RouteRange
	HasNextHopIpv6Address() bool
	// Advanced returns BgpRouteAdvanced, set in BgpL3VpnV6RouteRange.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	Advanced() BgpRouteAdvanced
	// SetAdvanced assigns BgpRouteAdvanced provided by user to BgpL3VpnV6RouteRange.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	SetAdvanced(value BgpRouteAdvanced) BgpL3VpnV6RouteRange
	// HasAdvanced checks if Advanced has been set in BgpL3VpnV6RouteRange
	HasAdvanced() bool
	// Communities returns BgpL3VpnV6RouteRangeBgpCommunityIterIter, set in BgpL3VpnV6RouteRange
	Communities() BgpL3VpnV6RouteRangeBgpCommunityIter
	// AsPath returns BgpAsPath, set in BgpL3VpnV6RouteRange.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	AsPath() BgpAsPath
	// SetAsPath assigns BgpAsPath provided by user to BgpL3VpnV6RouteRange.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	SetAsPath(value BgpAsPath) BgpL3VpnV6RouteRange
	// HasAsPath checks if AsPath has been set in BgpL3VpnV6RouteRange
	HasAsPath() bool
	// AddPath returns BgpAddPath, set in BgpL3VpnV6RouteRange.
	// BgpAddPath is the BGP Additional Paths feature is a BGP extension that allows the  advertisement of multiple paths for the same prefix without the new  paths implicitly replacing any previous paths.
	AddPath() BgpAddPath
	// SetAddPath assigns BgpAddPath provided by user to BgpL3VpnV6RouteRange.
	// BgpAddPath is the BGP Additional Paths feature is a BGP extension that allows the  advertisement of multiple paths for the same prefix without the new  paths implicitly replacing any previous paths.
	SetAddPath(value BgpAddPath) BgpL3VpnV6RouteRange
	// HasAddPath checks if AddPath has been set in BgpL3VpnV6RouteRange
	HasAddPath() bool
	// Name returns string, set in BgpL3VpnV6RouteRange.
	Name() string
	// SetName assigns string provided by user to BgpL3VpnV6RouteRange
	SetName(value string) BgpL3VpnV6RouteRange
	// ExtendedCommunities returns BgpL3VpnV6RouteRangeBgpExtendedCommunityIterIter, set in BgpL3VpnV6RouteRange
	ExtendedCommunities() BgpL3VpnV6RouteRangeBgpExtendedCommunityIter
	// ServiceBinding returns BgpL3VpnV6ServiceBinding, set in BgpL3VpnV6RouteRange.
	ServiceBinding() BgpL3VpnV6ServiceBinding
	// SetServiceBinding assigns BgpL3VpnV6ServiceBinding provided by user to BgpL3VpnV6RouteRange.
	SetServiceBinding(value BgpL3VpnV6ServiceBinding) BgpL3VpnV6RouteRange
	// HasServiceBinding checks if ServiceBinding has been set in BgpL3VpnV6RouteRange
	HasServiceBinding() bool
	setNil()
}

// A list of group of IPv6 route addresses.
// Addresses returns a []V6RouteAddress
func (obj *bgpL3VpnV6RouteRange) Addresses() BgpL3VpnV6RouteRangeV6RouteAddressIter {
	if len(obj.obj.Addresses) == 0 {
		obj.obj.Addresses = []*otg.V6RouteAddress{}
	}
	if obj.addressesHolder == nil {
		obj.addressesHolder = newBgpL3VpnV6RouteRangeV6RouteAddressIter(&obj.obj.Addresses).setMsg(obj)
	}
	return obj.addressesHolder
}

type bgpL3VpnV6RouteRangeV6RouteAddressIter struct {
	obj                 *bgpL3VpnV6RouteRange
	v6RouteAddressSlice []V6RouteAddress
	fieldPtr            *[]*otg.V6RouteAddress
}

func newBgpL3VpnV6RouteRangeV6RouteAddressIter(ptr *[]*otg.V6RouteAddress) BgpL3VpnV6RouteRangeV6RouteAddressIter {
	return &bgpL3VpnV6RouteRangeV6RouteAddressIter{fieldPtr: ptr}
}

type BgpL3VpnV6RouteRangeV6RouteAddressIter interface {
	setMsg(*bgpL3VpnV6RouteRange) BgpL3VpnV6RouteRangeV6RouteAddressIter
	Items() []V6RouteAddress
	Add() V6RouteAddress
	Append(items ...V6RouteAddress) BgpL3VpnV6RouteRangeV6RouteAddressIter
	Set(index int, newObj V6RouteAddress) BgpL3VpnV6RouteRangeV6RouteAddressIter
	Clear() BgpL3VpnV6RouteRangeV6RouteAddressIter
	clearHolderSlice() BgpL3VpnV6RouteRangeV6RouteAddressIter
	appendHolderSlice(item V6RouteAddress) BgpL3VpnV6RouteRangeV6RouteAddressIter
}

func (obj *bgpL3VpnV6RouteRangeV6RouteAddressIter) setMsg(msg *bgpL3VpnV6RouteRange) BgpL3VpnV6RouteRangeV6RouteAddressIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&v6RouteAddress{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpL3VpnV6RouteRangeV6RouteAddressIter) Items() []V6RouteAddress {
	return obj.v6RouteAddressSlice
}

func (obj *bgpL3VpnV6RouteRangeV6RouteAddressIter) Add() V6RouteAddress {
	newObj := &otg.V6RouteAddress{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &v6RouteAddress{obj: newObj}
	newLibObj.setDefault()
	obj.v6RouteAddressSlice = append(obj.v6RouteAddressSlice, newLibObj)
	return newLibObj
}

func (obj *bgpL3VpnV6RouteRangeV6RouteAddressIter) Append(items ...V6RouteAddress) BgpL3VpnV6RouteRangeV6RouteAddressIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.v6RouteAddressSlice = append(obj.v6RouteAddressSlice, item)
	}
	return obj
}

func (obj *bgpL3VpnV6RouteRangeV6RouteAddressIter) Set(index int, newObj V6RouteAddress) BgpL3VpnV6RouteRangeV6RouteAddressIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.v6RouteAddressSlice[index] = newObj
	return obj
}
func (obj *bgpL3VpnV6RouteRangeV6RouteAddressIter) Clear() BgpL3VpnV6RouteRangeV6RouteAddressIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.V6RouteAddress{}
		obj.v6RouteAddressSlice = []V6RouteAddress{}
	}
	return obj
}
func (obj *bgpL3VpnV6RouteRangeV6RouteAddressIter) clearHolderSlice() BgpL3VpnV6RouteRangeV6RouteAddressIter {
	if len(obj.v6RouteAddressSlice) > 0 {
		obj.v6RouteAddressSlice = []V6RouteAddress{}
	}
	return obj
}
func (obj *bgpL3VpnV6RouteRangeV6RouteAddressIter) appendHolderSlice(item V6RouteAddress) BgpL3VpnV6RouteRangeV6RouteAddressIter {
	obj.v6RouteAddressSlice = append(obj.v6RouteAddressSlice, item)
	return obj
}

type BgpL3VpnV6RouteRangeNextHopModeEnum string

// Enum of NextHopMode on BgpL3VpnV6RouteRange
var BgpL3VpnV6RouteRangeNextHopMode = struct {
	LOCAL_IP BgpL3VpnV6RouteRangeNextHopModeEnum
	MANUAL   BgpL3VpnV6RouteRangeNextHopModeEnum
}{
	LOCAL_IP: BgpL3VpnV6RouteRangeNextHopModeEnum("local_ip"),
	MANUAL:   BgpL3VpnV6RouteRangeNextHopModeEnum("manual"),
}

func (obj *bgpL3VpnV6RouteRange) NextHopMode() BgpL3VpnV6RouteRangeNextHopModeEnum {
	return BgpL3VpnV6RouteRangeNextHopModeEnum(obj.obj.NextHopMode.Enum().String())
}

// Specify the NextHop in MP REACH NLRI. The mode for setting the IP address  of the NextHop in the MP REACH NLRI can be one of the following:
// Local IP: Automatically fills the Nexthop with the Local IP of the BGP
// peer.
// If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// Manual: Override the Nexthop with any arbitrary IPv4/IPv6 address.
// NextHopMode returns a string
func (obj *bgpL3VpnV6RouteRange) HasNextHopMode() bool {
	return obj.obj.NextHopMode != nil
}

func (obj *bgpL3VpnV6RouteRange) SetNextHopMode(value BgpL3VpnV6RouteRangeNextHopModeEnum) BgpL3VpnV6RouteRange {
	intValue, ok := otg.BgpL3VpnV6RouteRange_NextHopMode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpL3VpnV6RouteRangeNextHopModeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpL3VpnV6RouteRange_NextHopMode_Enum(intValue)
	obj.obj.NextHopMode = &enumValue

	return obj
}

type BgpL3VpnV6RouteRangeNextHopAddressTypeEnum string

// Enum of NextHopAddressType on BgpL3VpnV6RouteRange
var BgpL3VpnV6RouteRangeNextHopAddressType = struct {
	IPV4 BgpL3VpnV6RouteRangeNextHopAddressTypeEnum
	IPV6 BgpL3VpnV6RouteRangeNextHopAddressTypeEnum
}{
	IPV4: BgpL3VpnV6RouteRangeNextHopAddressTypeEnum("ipv4"),
	IPV6: BgpL3VpnV6RouteRangeNextHopAddressTypeEnum("ipv6"),
}

func (obj *bgpL3VpnV6RouteRange) NextHopAddressType() BgpL3VpnV6RouteRangeNextHopAddressTypeEnum {
	return BgpL3VpnV6RouteRangeNextHopAddressTypeEnum(obj.obj.NextHopAddressType.Enum().String())
}

// If the Nexthop Mode is Manual, it sets the type of the NextHop IP address.
// NextHopAddressType returns a string
func (obj *bgpL3VpnV6RouteRange) HasNextHopAddressType() bool {
	return obj.obj.NextHopAddressType != nil
}

func (obj *bgpL3VpnV6RouteRange) SetNextHopAddressType(value BgpL3VpnV6RouteRangeNextHopAddressTypeEnum) BgpL3VpnV6RouteRange {
	intValue, ok := otg.BgpL3VpnV6RouteRange_NextHopAddressType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpL3VpnV6RouteRangeNextHopAddressTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpL3VpnV6RouteRange_NextHopAddressType_Enum(intValue)
	obj.obj.NextHopAddressType = &enumValue

	return obj
}

// The IPv4 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// NextHopIpv4Address returns a string
func (obj *bgpL3VpnV6RouteRange) NextHopIpv4Address() string {

	return *obj.obj.NextHopIpv4Address

}

// The IPv4 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// NextHopIpv4Address returns a string
func (obj *bgpL3VpnV6RouteRange) HasNextHopIpv4Address() bool {
	return obj.obj.NextHopIpv4Address != nil
}

// The IPv4 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// SetNextHopIpv4Address sets the string value in the BgpL3VpnV6RouteRange object
func (obj *bgpL3VpnV6RouteRange) SetNextHopIpv4Address(value string) BgpL3VpnV6RouteRange {

	obj.obj.NextHopIpv4Address = &value
	return obj
}

// The IPv6 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv6.
// NextHopIpv6Address returns a string
func (obj *bgpL3VpnV6RouteRange) NextHopIpv6Address() string {

	return *obj.obj.NextHopIpv6Address

}

// The IPv6 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv6.
// NextHopIpv6Address returns a string
func (obj *bgpL3VpnV6RouteRange) HasNextHopIpv6Address() bool {
	return obj.obj.NextHopIpv6Address != nil
}

// The IPv6 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv6.
// SetNextHopIpv6Address sets the string value in the BgpL3VpnV6RouteRange object
func (obj *bgpL3VpnV6RouteRange) SetNextHopIpv6Address(value string) BgpL3VpnV6RouteRange {

	obj.obj.NextHopIpv6Address = &value
	return obj
}

// description is TBD
// Advanced returns a BgpRouteAdvanced
func (obj *bgpL3VpnV6RouteRange) Advanced() BgpRouteAdvanced {
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
func (obj *bgpL3VpnV6RouteRange) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// description is TBD
// SetAdvanced sets the BgpRouteAdvanced value in the BgpL3VpnV6RouteRange object
func (obj *bgpL3VpnV6RouteRange) SetAdvanced(value BgpRouteAdvanced) BgpL3VpnV6RouteRange {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// Optional community settings.
// Communities returns a []BgpCommunity
func (obj *bgpL3VpnV6RouteRange) Communities() BgpL3VpnV6RouteRangeBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.BgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpL3VpnV6RouteRangeBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpL3VpnV6RouteRangeBgpCommunityIter struct {
	obj               *bgpL3VpnV6RouteRange
	bgpCommunitySlice []BgpCommunity
	fieldPtr          *[]*otg.BgpCommunity
}

func newBgpL3VpnV6RouteRangeBgpCommunityIter(ptr *[]*otg.BgpCommunity) BgpL3VpnV6RouteRangeBgpCommunityIter {
	return &bgpL3VpnV6RouteRangeBgpCommunityIter{fieldPtr: ptr}
}

type BgpL3VpnV6RouteRangeBgpCommunityIter interface {
	setMsg(*bgpL3VpnV6RouteRange) BgpL3VpnV6RouteRangeBgpCommunityIter
	Items() []BgpCommunity
	Add() BgpCommunity
	Append(items ...BgpCommunity) BgpL3VpnV6RouteRangeBgpCommunityIter
	Set(index int, newObj BgpCommunity) BgpL3VpnV6RouteRangeBgpCommunityIter
	Clear() BgpL3VpnV6RouteRangeBgpCommunityIter
	clearHolderSlice() BgpL3VpnV6RouteRangeBgpCommunityIter
	appendHolderSlice(item BgpCommunity) BgpL3VpnV6RouteRangeBgpCommunityIter
}

func (obj *bgpL3VpnV6RouteRangeBgpCommunityIter) setMsg(msg *bgpL3VpnV6RouteRange) BgpL3VpnV6RouteRangeBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpL3VpnV6RouteRangeBgpCommunityIter) Items() []BgpCommunity {
	return obj.bgpCommunitySlice
}

func (obj *bgpL3VpnV6RouteRangeBgpCommunityIter) Add() BgpCommunity {
	newObj := &otg.BgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpL3VpnV6RouteRangeBgpCommunityIter) Append(items ...BgpCommunity) BgpL3VpnV6RouteRangeBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpL3VpnV6RouteRangeBgpCommunityIter) Set(index int, newObj BgpCommunity) BgpL3VpnV6RouteRangeBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpL3VpnV6RouteRangeBgpCommunityIter) Clear() BgpL3VpnV6RouteRangeBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpCommunity{}
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpL3VpnV6RouteRangeBgpCommunityIter) clearHolderSlice() BgpL3VpnV6RouteRangeBgpCommunityIter {
	if len(obj.bgpCommunitySlice) > 0 {
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpL3VpnV6RouteRangeBgpCommunityIter) appendHolderSlice(item BgpCommunity) BgpL3VpnV6RouteRangeBgpCommunityIter {
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a BgpAsPath
func (obj *bgpL3VpnV6RouteRange) AsPath() BgpAsPath {
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
func (obj *bgpL3VpnV6RouteRange) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the BgpAsPath value in the BgpL3VpnV6RouteRange object
func (obj *bgpL3VpnV6RouteRange) SetAsPath(value BgpAsPath) BgpL3VpnV6RouteRange {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// description is TBD
// AddPath returns a BgpAddPath
func (obj *bgpL3VpnV6RouteRange) AddPath() BgpAddPath {
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
func (obj *bgpL3VpnV6RouteRange) HasAddPath() bool {
	return obj.obj.AddPath != nil
}

// description is TBD
// SetAddPath sets the BgpAddPath value in the BgpL3VpnV6RouteRange object
func (obj *bgpL3VpnV6RouteRange) SetAddPath(value BgpAddPath) BgpL3VpnV6RouteRange {

	obj.addPathHolder = nil
	obj.obj.AddPath = value.msg()

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpL3VpnV6RouteRange) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpL3VpnV6RouteRange object
func (obj *bgpL3VpnV6RouteRange) SetName(value string) BgpL3VpnV6RouteRange {

	obj.obj.Name = &value
	return obj
}

// Optional Extended Community settings. The Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an eight byte value. It is divided into two main parts. The first two bytes of the community encode a type and sub-type fields and the last six bytes carry a unique set of data in a format defined by the type and sub-type field. Extended communities provide a larger range for grouping or categorizing communities.
// ExtendedCommunities returns a []BgpExtendedCommunity
func (obj *bgpL3VpnV6RouteRange) ExtendedCommunities() BgpL3VpnV6RouteRangeBgpExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.BgpExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpL3VpnV6RouteRangeBgpExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpL3VpnV6RouteRangeBgpExtendedCommunityIter struct {
	obj                       *bgpL3VpnV6RouteRange
	bgpExtendedCommunitySlice []BgpExtendedCommunity
	fieldPtr                  *[]*otg.BgpExtendedCommunity
}

func newBgpL3VpnV6RouteRangeBgpExtendedCommunityIter(ptr *[]*otg.BgpExtendedCommunity) BgpL3VpnV6RouteRangeBgpExtendedCommunityIter {
	return &bgpL3VpnV6RouteRangeBgpExtendedCommunityIter{fieldPtr: ptr}
}

type BgpL3VpnV6RouteRangeBgpExtendedCommunityIter interface {
	setMsg(*bgpL3VpnV6RouteRange) BgpL3VpnV6RouteRangeBgpExtendedCommunityIter
	Items() []BgpExtendedCommunity
	Add() BgpExtendedCommunity
	Append(items ...BgpExtendedCommunity) BgpL3VpnV6RouteRangeBgpExtendedCommunityIter
	Set(index int, newObj BgpExtendedCommunity) BgpL3VpnV6RouteRangeBgpExtendedCommunityIter
	Clear() BgpL3VpnV6RouteRangeBgpExtendedCommunityIter
	clearHolderSlice() BgpL3VpnV6RouteRangeBgpExtendedCommunityIter
	appendHolderSlice(item BgpExtendedCommunity) BgpL3VpnV6RouteRangeBgpExtendedCommunityIter
}

func (obj *bgpL3VpnV6RouteRangeBgpExtendedCommunityIter) setMsg(msg *bgpL3VpnV6RouteRange) BgpL3VpnV6RouteRangeBgpExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpL3VpnV6RouteRangeBgpExtendedCommunityIter) Items() []BgpExtendedCommunity {
	return obj.bgpExtendedCommunitySlice
}

func (obj *bgpL3VpnV6RouteRangeBgpExtendedCommunityIter) Add() BgpExtendedCommunity {
	newObj := &otg.BgpExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpL3VpnV6RouteRangeBgpExtendedCommunityIter) Append(items ...BgpExtendedCommunity) BgpL3VpnV6RouteRangeBgpExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpL3VpnV6RouteRangeBgpExtendedCommunityIter) Set(index int, newObj BgpExtendedCommunity) BgpL3VpnV6RouteRangeBgpExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpL3VpnV6RouteRangeBgpExtendedCommunityIter) Clear() BgpL3VpnV6RouteRangeBgpExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtendedCommunity{}
		obj.bgpExtendedCommunitySlice = []BgpExtendedCommunity{}
	}
	return obj
}
func (obj *bgpL3VpnV6RouteRangeBgpExtendedCommunityIter) clearHolderSlice() BgpL3VpnV6RouteRangeBgpExtendedCommunityIter {
	if len(obj.bgpExtendedCommunitySlice) > 0 {
		obj.bgpExtendedCommunitySlice = []BgpExtendedCommunity{}
	}
	return obj
}
func (obj *bgpL3VpnV6RouteRangeBgpExtendedCommunityIter) appendHolderSlice(item BgpExtendedCommunity) BgpL3VpnV6RouteRangeBgpExtendedCommunityIter {
	obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, item)
	return obj
}

// Selects and configures this route range's VPN dataplane binding. Only the VPN MPLS label is currently supported; the choice structure exists so that an additional dataplane binding can be added later without breaking this field.
// ServiceBinding returns a BgpL3VpnV6ServiceBinding
func (obj *bgpL3VpnV6RouteRange) ServiceBinding() BgpL3VpnV6ServiceBinding {
	if obj.obj.ServiceBinding == nil {
		obj.obj.ServiceBinding = NewBgpL3VpnV6ServiceBinding().msg()
	}
	if obj.serviceBindingHolder == nil {
		obj.serviceBindingHolder = &bgpL3VpnV6ServiceBinding{obj: obj.obj.ServiceBinding}
	}
	return obj.serviceBindingHolder
}

// Selects and configures this route range's VPN dataplane binding. Only the VPN MPLS label is currently supported; the choice structure exists so that an additional dataplane binding can be added later without breaking this field.
// ServiceBinding returns a BgpL3VpnV6ServiceBinding
func (obj *bgpL3VpnV6RouteRange) HasServiceBinding() bool {
	return obj.obj.ServiceBinding != nil
}

// Selects and configures this route range's VPN dataplane binding. Only the VPN MPLS label is currently supported; the choice structure exists so that an additional dataplane binding can be added later without breaking this field.
// SetServiceBinding sets the BgpL3VpnV6ServiceBinding value in the BgpL3VpnV6RouteRange object
func (obj *bgpL3VpnV6RouteRange) SetServiceBinding(value BgpL3VpnV6ServiceBinding) BgpL3VpnV6RouteRange {

	obj.serviceBindingHolder = nil
	obj.obj.ServiceBinding = value.msg()

	return obj
}

func (obj *bgpL3VpnV6RouteRange) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Addresses) != 0 {

		if set_default {
			obj.Addresses().clearHolderSlice()
			for _, item := range obj.obj.Addresses {
				obj.Addresses().appendHolderSlice(&v6RouteAddress{obj: item})
			}
		}
		for _, item := range obj.Addresses().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.NextHopIpv4Address != nil {

		err := obj.validateIpv4(obj.NextHopIpv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpL3VpnV6RouteRange.NextHopIpv4Address"))
		}

	}

	if obj.obj.NextHopIpv6Address != nil {

		err := obj.validateIpv6(obj.NextHopIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpL3VpnV6RouteRange.NextHopIpv6Address"))
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
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpL3VpnV6RouteRange")
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

	if obj.obj.ServiceBinding != nil {

		obj.ServiceBinding().validateObj(vObj, set_default)
	}

}

func (obj *bgpL3VpnV6RouteRange) setDefault() {
	if obj.obj.NextHopMode == nil {
		obj.SetNextHopMode(BgpL3VpnV6RouteRangeNextHopMode.LOCAL_IP)

	}
	if obj.obj.NextHopAddressType == nil {
		obj.SetNextHopAddressType(BgpL3VpnV6RouteRangeNextHopAddressType.IPV6)

	}
	if obj.obj.NextHopIpv4Address == nil {
		obj.SetNextHopIpv4Address("0.0.0.0")
	}
	if obj.obj.NextHopIpv6Address == nil {
		obj.SetNextHopIpv6Address("::0")
	}

}
