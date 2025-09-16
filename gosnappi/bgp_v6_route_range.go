package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV6RouteRange *****
type bgpV6RouteRange struct {
	validation
	obj                       *otg.BgpV6RouteRange
	marshaller                marshalBgpV6RouteRange
	unMarshaller              unMarshalBgpV6RouteRange
	addressesHolder           BgpV6RouteRangeV6RouteAddressIter
	advancedHolder            BgpRouteAdvanced
	communitiesHolder         BgpV6RouteRangeBgpCommunityIter
	asPathHolder              BgpAsPath
	addPathHolder             BgpAddPath
	extCommunitiesHolder      BgpV6RouteRangeBgpExtCommunityIter
	extendedCommunitiesHolder BgpV6RouteRangeBgpExtendedCommunityIter
}

func NewBgpV6RouteRange() BgpV6RouteRange {
	obj := bgpV6RouteRange{obj: &otg.BgpV6RouteRange{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV6RouteRange) msg() *otg.BgpV6RouteRange {
	return obj.obj
}

func (obj *bgpV6RouteRange) setMsg(msg *otg.BgpV6RouteRange) BgpV6RouteRange {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV6RouteRange struct {
	obj *bgpV6RouteRange
}

type marshalBgpV6RouteRange interface {
	// ToProto marshals BgpV6RouteRange to protobuf object *otg.BgpV6RouteRange
	ToProto() (*otg.BgpV6RouteRange, error)
	// ToPbText marshals BgpV6RouteRange to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV6RouteRange to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV6RouteRange to JSON text
	ToJson() (string, error)
}

type unMarshalbgpV6RouteRange struct {
	obj *bgpV6RouteRange
}

type unMarshalBgpV6RouteRange interface {
	// FromProto unmarshals BgpV6RouteRange from protobuf object *otg.BgpV6RouteRange
	FromProto(msg *otg.BgpV6RouteRange) (BgpV6RouteRange, error)
	// FromPbText unmarshals BgpV6RouteRange from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV6RouteRange from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV6RouteRange from JSON text
	FromJson(value string) error
}

func (obj *bgpV6RouteRange) Marshal() marshalBgpV6RouteRange {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV6RouteRange{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV6RouteRange) Unmarshal() unMarshalBgpV6RouteRange {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV6RouteRange{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV6RouteRange) ToProto() (*otg.BgpV6RouteRange, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV6RouteRange) FromProto(msg *otg.BgpV6RouteRange) (BgpV6RouteRange, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV6RouteRange) ToPbText() (string, error) {
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

func (m *unMarshalbgpV6RouteRange) FromPbText(value string) error {
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

func (m *marshalbgpV6RouteRange) ToYaml() (string, error) {
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

func (m *unMarshalbgpV6RouteRange) FromYaml(value string) error {
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

func (m *marshalbgpV6RouteRange) ToJson() (string, error) {
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

func (m *unMarshalbgpV6RouteRange) FromJson(value string) error {
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

func (obj *bgpV6RouteRange) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV6RouteRange) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV6RouteRange) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV6RouteRange) Clone() (BgpV6RouteRange, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV6RouteRange()
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

func (obj *bgpV6RouteRange) setNil() {
	obj.addressesHolder = nil
	obj.advancedHolder = nil
	obj.communitiesHolder = nil
	obj.asPathHolder = nil
	obj.addPathHolder = nil
	obj.extCommunitiesHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpV6RouteRange is emulated BGPv6 route range.
type BgpV6RouteRange interface {
	Validation
	// msg marshals BgpV6RouteRange to protobuf object *otg.BgpV6RouteRange
	// and doesn't set defaults
	msg() *otg.BgpV6RouteRange
	// setMsg unmarshals BgpV6RouteRange from protobuf object *otg.BgpV6RouteRange
	// and doesn't set defaults
	setMsg(*otg.BgpV6RouteRange) BgpV6RouteRange
	// provides marshal interface
	Marshal() marshalBgpV6RouteRange
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV6RouteRange
	// validate validates BgpV6RouteRange
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV6RouteRange, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Addresses returns BgpV6RouteRangeV6RouteAddressIterIter, set in BgpV6RouteRange
	Addresses() BgpV6RouteRangeV6RouteAddressIter
	// NextHopMode returns BgpV6RouteRangeNextHopModeEnum, set in BgpV6RouteRange
	NextHopMode() BgpV6RouteRangeNextHopModeEnum
	// SetNextHopMode assigns BgpV6RouteRangeNextHopModeEnum provided by user to BgpV6RouteRange
	SetNextHopMode(value BgpV6RouteRangeNextHopModeEnum) BgpV6RouteRange
	// HasNextHopMode checks if NextHopMode has been set in BgpV6RouteRange
	HasNextHopMode() bool
	// NextHopAddressType returns BgpV6RouteRangeNextHopAddressTypeEnum, set in BgpV6RouteRange
	NextHopAddressType() BgpV6RouteRangeNextHopAddressTypeEnum
	// SetNextHopAddressType assigns BgpV6RouteRangeNextHopAddressTypeEnum provided by user to BgpV6RouteRange
	SetNextHopAddressType(value BgpV6RouteRangeNextHopAddressTypeEnum) BgpV6RouteRange
	// HasNextHopAddressType checks if NextHopAddressType has been set in BgpV6RouteRange
	HasNextHopAddressType() bool
	// NextHopIpv4Address returns string, set in BgpV6RouteRange.
	NextHopIpv4Address() string
	// SetNextHopIpv4Address assigns string provided by user to BgpV6RouteRange
	SetNextHopIpv4Address(value string) BgpV6RouteRange
	// HasNextHopIpv4Address checks if NextHopIpv4Address has been set in BgpV6RouteRange
	HasNextHopIpv4Address() bool
	// NextHopIpv6Address returns string, set in BgpV6RouteRange.
	NextHopIpv6Address() string
	// SetNextHopIpv6Address assigns string provided by user to BgpV6RouteRange
	SetNextHopIpv6Address(value string) BgpV6RouteRange
	// HasNextHopIpv6Address checks if NextHopIpv6Address has been set in BgpV6RouteRange
	HasNextHopIpv6Address() bool
	// Advanced returns BgpRouteAdvanced, set in BgpV6RouteRange.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	Advanced() BgpRouteAdvanced
	// SetAdvanced assigns BgpRouteAdvanced provided by user to BgpV6RouteRange.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	SetAdvanced(value BgpRouteAdvanced) BgpV6RouteRange
	// HasAdvanced checks if Advanced has been set in BgpV6RouteRange
	HasAdvanced() bool
	// Communities returns BgpV6RouteRangeBgpCommunityIterIter, set in BgpV6RouteRange
	Communities() BgpV6RouteRangeBgpCommunityIter
	// AsPath returns BgpAsPath, set in BgpV6RouteRange.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	AsPath() BgpAsPath
	// SetAsPath assigns BgpAsPath provided by user to BgpV6RouteRange.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	SetAsPath(value BgpAsPath) BgpV6RouteRange
	// HasAsPath checks if AsPath has been set in BgpV6RouteRange
	HasAsPath() bool
	// AddPath returns BgpAddPath, set in BgpV6RouteRange.
	// BgpAddPath is the BGP Additional Paths feature is a BGP extension that allows the  advertisement of multiple paths for the same prefix without the new  paths implicitly replacing any previous paths.
	AddPath() BgpAddPath
	// SetAddPath assigns BgpAddPath provided by user to BgpV6RouteRange.
	// BgpAddPath is the BGP Additional Paths feature is a BGP extension that allows the  advertisement of multiple paths for the same prefix without the new  paths implicitly replacing any previous paths.
	SetAddPath(value BgpAddPath) BgpV6RouteRange
	// HasAddPath checks if AddPath has been set in BgpV6RouteRange
	HasAddPath() bool
	// Name returns string, set in BgpV6RouteRange.
	Name() string
	// SetName assigns string provided by user to BgpV6RouteRange
	SetName(value string) BgpV6RouteRange
	// ExtCommunities returns BgpV6RouteRangeBgpExtCommunityIterIter, set in BgpV6RouteRange
	ExtCommunities() BgpV6RouteRangeBgpExtCommunityIter
	// ExtendedCommunities returns BgpV6RouteRangeBgpExtendedCommunityIterIter, set in BgpV6RouteRange
	ExtendedCommunities() BgpV6RouteRangeBgpExtendedCommunityIter
	setNil()
}

// A list of group of IPv6 route addresses.
// Addresses returns a []V6RouteAddress
func (obj *bgpV6RouteRange) Addresses() BgpV6RouteRangeV6RouteAddressIter {
	if len(obj.obj.Addresses) == 0 {
		obj.obj.Addresses = []*otg.V6RouteAddress{}
	}
	if obj.addressesHolder == nil {
		obj.addressesHolder = newBgpV6RouteRangeV6RouteAddressIter(&obj.obj.Addresses).setMsg(obj)
	}
	return obj.addressesHolder
}

type bgpV6RouteRangeV6RouteAddressIter struct {
	obj                 *bgpV6RouteRange
	v6RouteAddressSlice []V6RouteAddress
	fieldPtr            *[]*otg.V6RouteAddress
}

func newBgpV6RouteRangeV6RouteAddressIter(ptr *[]*otg.V6RouteAddress) BgpV6RouteRangeV6RouteAddressIter {
	return &bgpV6RouteRangeV6RouteAddressIter{fieldPtr: ptr}
}

type BgpV6RouteRangeV6RouteAddressIter interface {
	setMsg(*bgpV6RouteRange) BgpV6RouteRangeV6RouteAddressIter
	Items() []V6RouteAddress
	Add() V6RouteAddress
	Append(items ...V6RouteAddress) BgpV6RouteRangeV6RouteAddressIter
	Set(index int, newObj V6RouteAddress) BgpV6RouteRangeV6RouteAddressIter
	Clear() BgpV6RouteRangeV6RouteAddressIter
	clearHolderSlice() BgpV6RouteRangeV6RouteAddressIter
	appendHolderSlice(item V6RouteAddress) BgpV6RouteRangeV6RouteAddressIter
}

func (obj *bgpV6RouteRangeV6RouteAddressIter) setMsg(msg *bgpV6RouteRange) BgpV6RouteRangeV6RouteAddressIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&v6RouteAddress{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6RouteRangeV6RouteAddressIter) Items() []V6RouteAddress {
	return obj.v6RouteAddressSlice
}

func (obj *bgpV6RouteRangeV6RouteAddressIter) Add() V6RouteAddress {
	newObj := &otg.V6RouteAddress{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &v6RouteAddress{obj: newObj}
	newLibObj.setDefault()
	obj.v6RouteAddressSlice = append(obj.v6RouteAddressSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6RouteRangeV6RouteAddressIter) Append(items ...V6RouteAddress) BgpV6RouteRangeV6RouteAddressIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.v6RouteAddressSlice = append(obj.v6RouteAddressSlice, item)
	}
	return obj
}

func (obj *bgpV6RouteRangeV6RouteAddressIter) Set(index int, newObj V6RouteAddress) BgpV6RouteRangeV6RouteAddressIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.v6RouteAddressSlice[index] = newObj
	return obj
}
func (obj *bgpV6RouteRangeV6RouteAddressIter) Clear() BgpV6RouteRangeV6RouteAddressIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.V6RouteAddress{}
		obj.v6RouteAddressSlice = []V6RouteAddress{}
	}
	return obj
}
func (obj *bgpV6RouteRangeV6RouteAddressIter) clearHolderSlice() BgpV6RouteRangeV6RouteAddressIter {
	if len(obj.v6RouteAddressSlice) > 0 {
		obj.v6RouteAddressSlice = []V6RouteAddress{}
	}
	return obj
}
func (obj *bgpV6RouteRangeV6RouteAddressIter) appendHolderSlice(item V6RouteAddress) BgpV6RouteRangeV6RouteAddressIter {
	obj.v6RouteAddressSlice = append(obj.v6RouteAddressSlice, item)
	return obj
}

type BgpV6RouteRangeNextHopModeEnum string

// Enum of NextHopMode on BgpV6RouteRange
var BgpV6RouteRangeNextHopMode = struct {
	LOCAL_IP BgpV6RouteRangeNextHopModeEnum
	MANUAL   BgpV6RouteRangeNextHopModeEnum
}{
	LOCAL_IP: BgpV6RouteRangeNextHopModeEnum("local_ip"),
	MANUAL:   BgpV6RouteRangeNextHopModeEnum("manual"),
}

func (obj *bgpV6RouteRange) NextHopMode() BgpV6RouteRangeNextHopModeEnum {
	return BgpV6RouteRangeNextHopModeEnum(obj.obj.NextHopMode.Enum().String())
}

// Specify the NextHop in MP REACH NLRI. The mode for setting the IP address  of the NextHop in the MP REACH NLRI can be one of the following:
// Local IP: Automatically fills the Nexthop with the Local IP of the BGP
// peer.
// If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// Manual: Override the Nexthop with any arbitrary IPv4/IPv6 address.
// NextHopMode returns a string
func (obj *bgpV6RouteRange) HasNextHopMode() bool {
	return obj.obj.NextHopMode != nil
}

func (obj *bgpV6RouteRange) SetNextHopMode(value BgpV6RouteRangeNextHopModeEnum) BgpV6RouteRange {
	intValue, ok := otg.BgpV6RouteRange_NextHopMode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV6RouteRangeNextHopModeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV6RouteRange_NextHopMode_Enum(intValue)
	obj.obj.NextHopMode = &enumValue

	return obj
}

type BgpV6RouteRangeNextHopAddressTypeEnum string

// Enum of NextHopAddressType on BgpV6RouteRange
var BgpV6RouteRangeNextHopAddressType = struct {
	IPV4 BgpV6RouteRangeNextHopAddressTypeEnum
	IPV6 BgpV6RouteRangeNextHopAddressTypeEnum
}{
	IPV4: BgpV6RouteRangeNextHopAddressTypeEnum("ipv4"),
	IPV6: BgpV6RouteRangeNextHopAddressTypeEnum("ipv6"),
}

func (obj *bgpV6RouteRange) NextHopAddressType() BgpV6RouteRangeNextHopAddressTypeEnum {
	return BgpV6RouteRangeNextHopAddressTypeEnum(obj.obj.NextHopAddressType.Enum().String())
}

// If the Nexthop Mode is Manual, it sets the type of the NextHop IP address.
// NextHopAddressType returns a string
func (obj *bgpV6RouteRange) HasNextHopAddressType() bool {
	return obj.obj.NextHopAddressType != nil
}

func (obj *bgpV6RouteRange) SetNextHopAddressType(value BgpV6RouteRangeNextHopAddressTypeEnum) BgpV6RouteRange {
	intValue, ok := otg.BgpV6RouteRange_NextHopAddressType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV6RouteRangeNextHopAddressTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV6RouteRange_NextHopAddressType_Enum(intValue)
	obj.obj.NextHopAddressType = &enumValue

	return obj
}

// The IPv4 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// NextHopIpv4Address returns a string
func (obj *bgpV6RouteRange) NextHopIpv4Address() string {

	return *obj.obj.NextHopIpv4Address

}

// The IPv4 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// NextHopIpv4Address returns a string
func (obj *bgpV6RouteRange) HasNextHopIpv4Address() bool {
	return obj.obj.NextHopIpv4Address != nil
}

// The IPv4 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// SetNextHopIpv4Address sets the string value in the BgpV6RouteRange object
func (obj *bgpV6RouteRange) SetNextHopIpv4Address(value string) BgpV6RouteRange {

	obj.obj.NextHopIpv4Address = &value
	return obj
}

// The IPv6 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv6.
// NextHopIpv6Address returns a string
func (obj *bgpV6RouteRange) NextHopIpv6Address() string {

	return *obj.obj.NextHopIpv6Address

}

// The IPv6 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv6.
// NextHopIpv6Address returns a string
func (obj *bgpV6RouteRange) HasNextHopIpv6Address() bool {
	return obj.obj.NextHopIpv6Address != nil
}

// The IPv6 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv6.
// SetNextHopIpv6Address sets the string value in the BgpV6RouteRange object
func (obj *bgpV6RouteRange) SetNextHopIpv6Address(value string) BgpV6RouteRange {

	obj.obj.NextHopIpv6Address = &value
	return obj
}

// description is TBD
// Advanced returns a BgpRouteAdvanced
func (obj *bgpV6RouteRange) Advanced() BgpRouteAdvanced {
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
func (obj *bgpV6RouteRange) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// description is TBD
// SetAdvanced sets the BgpRouteAdvanced value in the BgpV6RouteRange object
func (obj *bgpV6RouteRange) SetAdvanced(value BgpRouteAdvanced) BgpV6RouteRange {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// Optional community settings.
// Communities returns a []BgpCommunity
func (obj *bgpV6RouteRange) Communities() BgpV6RouteRangeBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.BgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpV6RouteRangeBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpV6RouteRangeBgpCommunityIter struct {
	obj               *bgpV6RouteRange
	bgpCommunitySlice []BgpCommunity
	fieldPtr          *[]*otg.BgpCommunity
}

func newBgpV6RouteRangeBgpCommunityIter(ptr *[]*otg.BgpCommunity) BgpV6RouteRangeBgpCommunityIter {
	return &bgpV6RouteRangeBgpCommunityIter{fieldPtr: ptr}
}

type BgpV6RouteRangeBgpCommunityIter interface {
	setMsg(*bgpV6RouteRange) BgpV6RouteRangeBgpCommunityIter
	Items() []BgpCommunity
	Add() BgpCommunity
	Append(items ...BgpCommunity) BgpV6RouteRangeBgpCommunityIter
	Set(index int, newObj BgpCommunity) BgpV6RouteRangeBgpCommunityIter
	Clear() BgpV6RouteRangeBgpCommunityIter
	clearHolderSlice() BgpV6RouteRangeBgpCommunityIter
	appendHolderSlice(item BgpCommunity) BgpV6RouteRangeBgpCommunityIter
}

func (obj *bgpV6RouteRangeBgpCommunityIter) setMsg(msg *bgpV6RouteRange) BgpV6RouteRangeBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6RouteRangeBgpCommunityIter) Items() []BgpCommunity {
	return obj.bgpCommunitySlice
}

func (obj *bgpV6RouteRangeBgpCommunityIter) Add() BgpCommunity {
	newObj := &otg.BgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6RouteRangeBgpCommunityIter) Append(items ...BgpCommunity) BgpV6RouteRangeBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV6RouteRangeBgpCommunityIter) Set(index int, newObj BgpCommunity) BgpV6RouteRangeBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV6RouteRangeBgpCommunityIter) Clear() BgpV6RouteRangeBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpCommunity{}
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpV6RouteRangeBgpCommunityIter) clearHolderSlice() BgpV6RouteRangeBgpCommunityIter {
	if len(obj.bgpCommunitySlice) > 0 {
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpV6RouteRangeBgpCommunityIter) appendHolderSlice(item BgpCommunity) BgpV6RouteRangeBgpCommunityIter {
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a BgpAsPath
func (obj *bgpV6RouteRange) AsPath() BgpAsPath {
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
func (obj *bgpV6RouteRange) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the BgpAsPath value in the BgpV6RouteRange object
func (obj *bgpV6RouteRange) SetAsPath(value BgpAsPath) BgpV6RouteRange {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// description is TBD
// AddPath returns a BgpAddPath
func (obj *bgpV6RouteRange) AddPath() BgpAddPath {
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
func (obj *bgpV6RouteRange) HasAddPath() bool {
	return obj.obj.AddPath != nil
}

// description is TBD
// SetAddPath sets the BgpAddPath value in the BgpV6RouteRange object
func (obj *bgpV6RouteRange) SetAddPath(value BgpAddPath) BgpV6RouteRange {

	obj.addPathHolder = nil
	obj.obj.AddPath = value.msg()

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpV6RouteRange) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpV6RouteRange object
func (obj *bgpV6RouteRange) SetName(value string) BgpV6RouteRange {

	obj.obj.Name = &value
	return obj
}

// Deprecated: This property is deprecated in favor of property extended_communities
//
// Deprecated: This property is deprecated in favor of property extended_communities
//
// Optional Extended Community settings. The Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an 8-Bytes value. It is divided into two main parts. The first 2 Bytes of the community encode a type and sub-type fields and the last 6 Bytes carry a unique set of data in a format defined by the type and sub-type field. Extended communities provide a larger  range for grouping or categorizing communities. When type is administrator_as_2octet or administrator_as_4octet, the valid sub types are route target and origin. The valid value for  administrator_as_2octet and administrator_as_4octet type is either two byte AS followed by four byte local administrator id or four byte AS followed by two  byte local administrator id.  When type is administrator_ipv4_address the valid sub types are route target and origin. The valid value for  administrator_ipv4_address is a four byte IPv4 address followed by a two byte local administrator id.  When type is opaque, valid sub types are color and encapsulation. When sub type is color, first two bytes of the value field contain flags and last four bytes  contains the value of the color. When sub type is encapsulation the first four bytes of value field are reserved and last two bytes carries the tunnel type from  IANA's "ETHER TYPES" registry e.g IPv4 (protocol type = 0x0800), IPv6 (protocol type = 0x86dd), and MPLS (protocol type = 0x8847). When type is administrator_as_2octet_link_bandwidth the valid sub type is extended_bandwidth. The first two bytes of the value field contains the AS number and the last four bytes contains the bandwidth in IEEE floating point format.  When type is evpn the valid subtype is mac_address. In the value field the low-order bit of the first byte(Flags) is defined as the "Sticky/static" flag and may be set to 1, indicating the MAC address is static and cannot move. The second byte is reserved and the  last four bytes contain the sequence number which is used to ensure that PEs retain the correct MAC/IP Advertisement route when multiple updates  occur for the same MAC address.  Note evpn type is defined mainly for use with evpn route updates and not for IPv4 and IPv6 route updates.
// ExtCommunities returns a []BgpExtCommunity
func (obj *bgpV6RouteRange) ExtCommunities() BgpV6RouteRangeBgpExtCommunityIter {
	if len(obj.obj.ExtCommunities) == 0 {
		obj.obj.ExtCommunities = []*otg.BgpExtCommunity{}
	}
	if obj.extCommunitiesHolder == nil {
		obj.extCommunitiesHolder = newBgpV6RouteRangeBgpExtCommunityIter(&obj.obj.ExtCommunities).setMsg(obj)
	}
	return obj.extCommunitiesHolder
}

type bgpV6RouteRangeBgpExtCommunityIter struct {
	obj                  *bgpV6RouteRange
	bgpExtCommunitySlice []BgpExtCommunity
	fieldPtr             *[]*otg.BgpExtCommunity
}

func newBgpV6RouteRangeBgpExtCommunityIter(ptr *[]*otg.BgpExtCommunity) BgpV6RouteRangeBgpExtCommunityIter {
	return &bgpV6RouteRangeBgpExtCommunityIter{fieldPtr: ptr}
}

type BgpV6RouteRangeBgpExtCommunityIter interface {
	setMsg(*bgpV6RouteRange) BgpV6RouteRangeBgpExtCommunityIter
	Items() []BgpExtCommunity
	Add() BgpExtCommunity
	Append(items ...BgpExtCommunity) BgpV6RouteRangeBgpExtCommunityIter
	Set(index int, newObj BgpExtCommunity) BgpV6RouteRangeBgpExtCommunityIter
	Clear() BgpV6RouteRangeBgpExtCommunityIter
	clearHolderSlice() BgpV6RouteRangeBgpExtCommunityIter
	appendHolderSlice(item BgpExtCommunity) BgpV6RouteRangeBgpExtCommunityIter
}

func (obj *bgpV6RouteRangeBgpExtCommunityIter) setMsg(msg *bgpV6RouteRange) BgpV6RouteRangeBgpExtCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6RouteRangeBgpExtCommunityIter) Items() []BgpExtCommunity {
	return obj.bgpExtCommunitySlice
}

func (obj *bgpV6RouteRangeBgpExtCommunityIter) Add() BgpExtCommunity {
	newObj := &otg.BgpExtCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6RouteRangeBgpExtCommunityIter) Append(items ...BgpExtCommunity) BgpV6RouteRangeBgpExtCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV6RouteRangeBgpExtCommunityIter) Set(index int, newObj BgpExtCommunity) BgpV6RouteRangeBgpExtCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV6RouteRangeBgpExtCommunityIter) Clear() BgpV6RouteRangeBgpExtCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtCommunity{}
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpV6RouteRangeBgpExtCommunityIter) clearHolderSlice() BgpV6RouteRangeBgpExtCommunityIter {
	if len(obj.bgpExtCommunitySlice) > 0 {
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpV6RouteRangeBgpExtCommunityIter) appendHolderSlice(item BgpExtCommunity) BgpV6RouteRangeBgpExtCommunityIter {
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	return obj
}

// Optional Extended Community settings. The Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an eight byte value. It is divided into two main parts. The first two bytes of the community encode a type and sub-type fields and the last six bytes carry a unique set of data in a format defined by the type and sub-type field. Extended communities provide a larger range for grouping or categorizing communities.
// ExtendedCommunities returns a []BgpExtendedCommunity
func (obj *bgpV6RouteRange) ExtendedCommunities() BgpV6RouteRangeBgpExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.BgpExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpV6RouteRangeBgpExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpV6RouteRangeBgpExtendedCommunityIter struct {
	obj                       *bgpV6RouteRange
	bgpExtendedCommunitySlice []BgpExtendedCommunity
	fieldPtr                  *[]*otg.BgpExtendedCommunity
}

func newBgpV6RouteRangeBgpExtendedCommunityIter(ptr *[]*otg.BgpExtendedCommunity) BgpV6RouteRangeBgpExtendedCommunityIter {
	return &bgpV6RouteRangeBgpExtendedCommunityIter{fieldPtr: ptr}
}

type BgpV6RouteRangeBgpExtendedCommunityIter interface {
	setMsg(*bgpV6RouteRange) BgpV6RouteRangeBgpExtendedCommunityIter
	Items() []BgpExtendedCommunity
	Add() BgpExtendedCommunity
	Append(items ...BgpExtendedCommunity) BgpV6RouteRangeBgpExtendedCommunityIter
	Set(index int, newObj BgpExtendedCommunity) BgpV6RouteRangeBgpExtendedCommunityIter
	Clear() BgpV6RouteRangeBgpExtendedCommunityIter
	clearHolderSlice() BgpV6RouteRangeBgpExtendedCommunityIter
	appendHolderSlice(item BgpExtendedCommunity) BgpV6RouteRangeBgpExtendedCommunityIter
}

func (obj *bgpV6RouteRangeBgpExtendedCommunityIter) setMsg(msg *bgpV6RouteRange) BgpV6RouteRangeBgpExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6RouteRangeBgpExtendedCommunityIter) Items() []BgpExtendedCommunity {
	return obj.bgpExtendedCommunitySlice
}

func (obj *bgpV6RouteRangeBgpExtendedCommunityIter) Add() BgpExtendedCommunity {
	newObj := &otg.BgpExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6RouteRangeBgpExtendedCommunityIter) Append(items ...BgpExtendedCommunity) BgpV6RouteRangeBgpExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV6RouteRangeBgpExtendedCommunityIter) Set(index int, newObj BgpExtendedCommunity) BgpV6RouteRangeBgpExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV6RouteRangeBgpExtendedCommunityIter) Clear() BgpV6RouteRangeBgpExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtendedCommunity{}
		obj.bgpExtendedCommunitySlice = []BgpExtendedCommunity{}
	}
	return obj
}
func (obj *bgpV6RouteRangeBgpExtendedCommunityIter) clearHolderSlice() BgpV6RouteRangeBgpExtendedCommunityIter {
	if len(obj.bgpExtendedCommunitySlice) > 0 {
		obj.bgpExtendedCommunitySlice = []BgpExtendedCommunity{}
	}
	return obj
}
func (obj *bgpV6RouteRangeBgpExtendedCommunityIter) appendHolderSlice(item BgpExtendedCommunity) BgpV6RouteRangeBgpExtendedCommunityIter {
	obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, item)
	return obj
}

func (obj *bgpV6RouteRange) validateObj(vObj *validation, set_default bool) {
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
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpV6RouteRange.NextHopIpv4Address"))
		}

	}

	if obj.obj.NextHopIpv6Address != nil {

		err := obj.validateIpv6(obj.NextHopIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpV6RouteRange.NextHopIpv6Address"))
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
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpV6RouteRange")
	}

	if len(obj.obj.ExtCommunities) != 0 {
		obj.addWarnings("ExtCommunities property in schema BgpV6RouteRange is deprecated, This property is deprecated in favor of property extended_communities")

		if set_default {
			obj.ExtCommunities().clearHolderSlice()
			for _, item := range obj.obj.ExtCommunities {
				obj.ExtCommunities().appendHolderSlice(&bgpExtCommunity{obj: item})
			}
		}
		for _, item := range obj.ExtCommunities().Items() {
			item.validateObj(vObj, set_default)
		}

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

}

func (obj *bgpV6RouteRange) setDefault() {
	if obj.obj.NextHopMode == nil {
		obj.SetNextHopMode(BgpV6RouteRangeNextHopMode.LOCAL_IP)

	}
	if obj.obj.NextHopAddressType == nil {
		obj.SetNextHopAddressType(BgpV6RouteRangeNextHopAddressType.IPV6)

	}
	if obj.obj.NextHopIpv4Address == nil {
		obj.SetNextHopIpv4Address("0.0.0.0")
	}
	if obj.obj.NextHopIpv6Address == nil {
		obj.SetNextHopIpv6Address("::0")
	}

}
