package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV4RouteRange *****
type bgpV4RouteRange struct {
	validation
	obj                       *otg.BgpV4RouteRange
	marshaller                marshalBgpV4RouteRange
	unMarshaller              unMarshalBgpV4RouteRange
	addressesHolder           BgpV4RouteRangeV4RouteAddressIter
	advancedHolder            BgpRouteAdvanced
	communitiesHolder         BgpV4RouteRangeBgpCommunityIter
	asPathHolder              BgpAsPath
	addPathHolder             BgpAddPath
	extCommunitiesHolder      BgpV4RouteRangeBgpExtCommunityIter
	extendedCommunitiesHolder BgpV4RouteRangeBgpExtendedCommunityIter
}

func NewBgpV4RouteRange() BgpV4RouteRange {
	obj := bgpV4RouteRange{obj: &otg.BgpV4RouteRange{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV4RouteRange) msg() *otg.BgpV4RouteRange {
	return obj.obj
}

func (obj *bgpV4RouteRange) setMsg(msg *otg.BgpV4RouteRange) BgpV4RouteRange {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV4RouteRange struct {
	obj *bgpV4RouteRange
}

type marshalBgpV4RouteRange interface {
	// ToProto marshals BgpV4RouteRange to protobuf object *otg.BgpV4RouteRange
	ToProto() (*otg.BgpV4RouteRange, error)
	// ToPbText marshals BgpV4RouteRange to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV4RouteRange to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV4RouteRange to JSON text
	ToJson() (string, error)
}

type unMarshalbgpV4RouteRange struct {
	obj *bgpV4RouteRange
}

type unMarshalBgpV4RouteRange interface {
	// FromProto unmarshals BgpV4RouteRange from protobuf object *otg.BgpV4RouteRange
	FromProto(msg *otg.BgpV4RouteRange) (BgpV4RouteRange, error)
	// FromPbText unmarshals BgpV4RouteRange from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV4RouteRange from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV4RouteRange from JSON text
	FromJson(value string) error
}

func (obj *bgpV4RouteRange) Marshal() marshalBgpV4RouteRange {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV4RouteRange{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV4RouteRange) Unmarshal() unMarshalBgpV4RouteRange {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV4RouteRange{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV4RouteRange) ToProto() (*otg.BgpV4RouteRange, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV4RouteRange) FromProto(msg *otg.BgpV4RouteRange) (BgpV4RouteRange, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV4RouteRange) ToPbText() (string, error) {
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

func (m *unMarshalbgpV4RouteRange) FromPbText(value string) error {
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

func (m *marshalbgpV4RouteRange) ToYaml() (string, error) {
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

func (m *unMarshalbgpV4RouteRange) FromYaml(value string) error {
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

func (m *marshalbgpV4RouteRange) ToJson() (string, error) {
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

func (m *unMarshalbgpV4RouteRange) FromJson(value string) error {
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

func (obj *bgpV4RouteRange) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV4RouteRange) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV4RouteRange) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV4RouteRange) Clone() (BgpV4RouteRange, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV4RouteRange()
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

func (obj *bgpV4RouteRange) setNil() {
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

// BgpV4RouteRange is emulated BGPv4 route range.
type BgpV4RouteRange interface {
	Validation
	// msg marshals BgpV4RouteRange to protobuf object *otg.BgpV4RouteRange
	// and doesn't set defaults
	msg() *otg.BgpV4RouteRange
	// setMsg unmarshals BgpV4RouteRange from protobuf object *otg.BgpV4RouteRange
	// and doesn't set defaults
	setMsg(*otg.BgpV4RouteRange) BgpV4RouteRange
	// provides marshal interface
	Marshal() marshalBgpV4RouteRange
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV4RouteRange
	// validate validates BgpV4RouteRange
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV4RouteRange, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Addresses returns BgpV4RouteRangeV4RouteAddressIterIter, set in BgpV4RouteRange
	Addresses() BgpV4RouteRangeV4RouteAddressIter
	// NextHopMode returns BgpV4RouteRangeNextHopModeEnum, set in BgpV4RouteRange
	NextHopMode() BgpV4RouteRangeNextHopModeEnum
	// SetNextHopMode assigns BgpV4RouteRangeNextHopModeEnum provided by user to BgpV4RouteRange
	SetNextHopMode(value BgpV4RouteRangeNextHopModeEnum) BgpV4RouteRange
	// HasNextHopMode checks if NextHopMode has been set in BgpV4RouteRange
	HasNextHopMode() bool
	// NextHopAddressType returns BgpV4RouteRangeNextHopAddressTypeEnum, set in BgpV4RouteRange
	NextHopAddressType() BgpV4RouteRangeNextHopAddressTypeEnum
	// SetNextHopAddressType assigns BgpV4RouteRangeNextHopAddressTypeEnum provided by user to BgpV4RouteRange
	SetNextHopAddressType(value BgpV4RouteRangeNextHopAddressTypeEnum) BgpV4RouteRange
	// HasNextHopAddressType checks if NextHopAddressType has been set in BgpV4RouteRange
	HasNextHopAddressType() bool
	// NextHopIpv4Address returns string, set in BgpV4RouteRange.
	NextHopIpv4Address() string
	// SetNextHopIpv4Address assigns string provided by user to BgpV4RouteRange
	SetNextHopIpv4Address(value string) BgpV4RouteRange
	// HasNextHopIpv4Address checks if NextHopIpv4Address has been set in BgpV4RouteRange
	HasNextHopIpv4Address() bool
	// NextHopIpv6Address returns string, set in BgpV4RouteRange.
	NextHopIpv6Address() string
	// SetNextHopIpv6Address assigns string provided by user to BgpV4RouteRange
	SetNextHopIpv6Address(value string) BgpV4RouteRange
	// HasNextHopIpv6Address checks if NextHopIpv6Address has been set in BgpV4RouteRange
	HasNextHopIpv6Address() bool
	// Advanced returns BgpRouteAdvanced, set in BgpV4RouteRange.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	Advanced() BgpRouteAdvanced
	// SetAdvanced assigns BgpRouteAdvanced provided by user to BgpV4RouteRange.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	SetAdvanced(value BgpRouteAdvanced) BgpV4RouteRange
	// HasAdvanced checks if Advanced has been set in BgpV4RouteRange
	HasAdvanced() bool
	// Communities returns BgpV4RouteRangeBgpCommunityIterIter, set in BgpV4RouteRange
	Communities() BgpV4RouteRangeBgpCommunityIter
	// AsPath returns BgpAsPath, set in BgpV4RouteRange.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	AsPath() BgpAsPath
	// SetAsPath assigns BgpAsPath provided by user to BgpV4RouteRange.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	SetAsPath(value BgpAsPath) BgpV4RouteRange
	// HasAsPath checks if AsPath has been set in BgpV4RouteRange
	HasAsPath() bool
	// AddPath returns BgpAddPath, set in BgpV4RouteRange.
	// BgpAddPath is the BGP Additional Paths feature is a BGP extension that allows the  advertisement of multiple paths for the same prefix without the new  paths implicitly replacing any previous paths.
	AddPath() BgpAddPath
	// SetAddPath assigns BgpAddPath provided by user to BgpV4RouteRange.
	// BgpAddPath is the BGP Additional Paths feature is a BGP extension that allows the  advertisement of multiple paths for the same prefix without the new  paths implicitly replacing any previous paths.
	SetAddPath(value BgpAddPath) BgpV4RouteRange
	// HasAddPath checks if AddPath has been set in BgpV4RouteRange
	HasAddPath() bool
	// Name returns string, set in BgpV4RouteRange.
	Name() string
	// SetName assigns string provided by user to BgpV4RouteRange
	SetName(value string) BgpV4RouteRange
	// ExtCommunities returns BgpV4RouteRangeBgpExtCommunityIterIter, set in BgpV4RouteRange
	ExtCommunities() BgpV4RouteRangeBgpExtCommunityIter
	// ExtendedCommunities returns BgpV4RouteRangeBgpExtendedCommunityIterIter, set in BgpV4RouteRange
	ExtendedCommunities() BgpV4RouteRangeBgpExtendedCommunityIter
	setNil()
}

// A list of group of IPv4 route addresses.
// Addresses returns a []V4RouteAddress
func (obj *bgpV4RouteRange) Addresses() BgpV4RouteRangeV4RouteAddressIter {
	if len(obj.obj.Addresses) == 0 {
		obj.obj.Addresses = []*otg.V4RouteAddress{}
	}
	if obj.addressesHolder == nil {
		obj.addressesHolder = newBgpV4RouteRangeV4RouteAddressIter(&obj.obj.Addresses).setMsg(obj)
	}
	return obj.addressesHolder
}

type bgpV4RouteRangeV4RouteAddressIter struct {
	obj                 *bgpV4RouteRange
	v4RouteAddressSlice []V4RouteAddress
	fieldPtr            *[]*otg.V4RouteAddress
}

func newBgpV4RouteRangeV4RouteAddressIter(ptr *[]*otg.V4RouteAddress) BgpV4RouteRangeV4RouteAddressIter {
	return &bgpV4RouteRangeV4RouteAddressIter{fieldPtr: ptr}
}

type BgpV4RouteRangeV4RouteAddressIter interface {
	setMsg(*bgpV4RouteRange) BgpV4RouteRangeV4RouteAddressIter
	Items() []V4RouteAddress
	Add() V4RouteAddress
	Append(items ...V4RouteAddress) BgpV4RouteRangeV4RouteAddressIter
	Set(index int, newObj V4RouteAddress) BgpV4RouteRangeV4RouteAddressIter
	Clear() BgpV4RouteRangeV4RouteAddressIter
	clearHolderSlice() BgpV4RouteRangeV4RouteAddressIter
	appendHolderSlice(item V4RouteAddress) BgpV4RouteRangeV4RouteAddressIter
}

func (obj *bgpV4RouteRangeV4RouteAddressIter) setMsg(msg *bgpV4RouteRange) BgpV4RouteRangeV4RouteAddressIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&v4RouteAddress{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4RouteRangeV4RouteAddressIter) Items() []V4RouteAddress {
	return obj.v4RouteAddressSlice
}

func (obj *bgpV4RouteRangeV4RouteAddressIter) Add() V4RouteAddress {
	newObj := &otg.V4RouteAddress{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &v4RouteAddress{obj: newObj}
	newLibObj.setDefault()
	obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4RouteRangeV4RouteAddressIter) Append(items ...V4RouteAddress) BgpV4RouteRangeV4RouteAddressIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, item)
	}
	return obj
}

func (obj *bgpV4RouteRangeV4RouteAddressIter) Set(index int, newObj V4RouteAddress) BgpV4RouteRangeV4RouteAddressIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.v4RouteAddressSlice[index] = newObj
	return obj
}
func (obj *bgpV4RouteRangeV4RouteAddressIter) Clear() BgpV4RouteRangeV4RouteAddressIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.V4RouteAddress{}
		obj.v4RouteAddressSlice = []V4RouteAddress{}
	}
	return obj
}
func (obj *bgpV4RouteRangeV4RouteAddressIter) clearHolderSlice() BgpV4RouteRangeV4RouteAddressIter {
	if len(obj.v4RouteAddressSlice) > 0 {
		obj.v4RouteAddressSlice = []V4RouteAddress{}
	}
	return obj
}
func (obj *bgpV4RouteRangeV4RouteAddressIter) appendHolderSlice(item V4RouteAddress) BgpV4RouteRangeV4RouteAddressIter {
	obj.v4RouteAddressSlice = append(obj.v4RouteAddressSlice, item)
	return obj
}

type BgpV4RouteRangeNextHopModeEnum string

// Enum of NextHopMode on BgpV4RouteRange
var BgpV4RouteRangeNextHopMode = struct {
	LOCAL_IP BgpV4RouteRangeNextHopModeEnum
	MANUAL   BgpV4RouteRangeNextHopModeEnum
}{
	LOCAL_IP: BgpV4RouteRangeNextHopModeEnum("local_ip"),
	MANUAL:   BgpV4RouteRangeNextHopModeEnum("manual"),
}

func (obj *bgpV4RouteRange) NextHopMode() BgpV4RouteRangeNextHopModeEnum {
	return BgpV4RouteRangeNextHopModeEnum(obj.obj.NextHopMode.Enum().String())
}

// Specify the NextHop in MP REACH NLRI. The mode for setting the IP address  of the NextHop in the MP REACH NLRI can be one of the following:
// Local IP: Automatically fills the Nexthop with the Local IP of the BGP
// peer.
// If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// Manual: Override the Nexthop with any arbitrary IPv4/IPv6 address.
// NextHopMode returns a string
func (obj *bgpV4RouteRange) HasNextHopMode() bool {
	return obj.obj.NextHopMode != nil
}

func (obj *bgpV4RouteRange) SetNextHopMode(value BgpV4RouteRangeNextHopModeEnum) BgpV4RouteRange {
	intValue, ok := otg.BgpV4RouteRange_NextHopMode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV4RouteRangeNextHopModeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV4RouteRange_NextHopMode_Enum(intValue)
	obj.obj.NextHopMode = &enumValue

	return obj
}

type BgpV4RouteRangeNextHopAddressTypeEnum string

// Enum of NextHopAddressType on BgpV4RouteRange
var BgpV4RouteRangeNextHopAddressType = struct {
	IPV4 BgpV4RouteRangeNextHopAddressTypeEnum
	IPV6 BgpV4RouteRangeNextHopAddressTypeEnum
}{
	IPV4: BgpV4RouteRangeNextHopAddressTypeEnum("ipv4"),
	IPV6: BgpV4RouteRangeNextHopAddressTypeEnum("ipv6"),
}

func (obj *bgpV4RouteRange) NextHopAddressType() BgpV4RouteRangeNextHopAddressTypeEnum {
	return BgpV4RouteRangeNextHopAddressTypeEnum(obj.obj.NextHopAddressType.Enum().String())
}

// If the Nexthop Mode is Manual, it sets the type of the NextHop IP address.
// NextHopAddressType returns a string
func (obj *bgpV4RouteRange) HasNextHopAddressType() bool {
	return obj.obj.NextHopAddressType != nil
}

func (obj *bgpV4RouteRange) SetNextHopAddressType(value BgpV4RouteRangeNextHopAddressTypeEnum) BgpV4RouteRange {
	intValue, ok := otg.BgpV4RouteRange_NextHopAddressType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV4RouteRangeNextHopAddressTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV4RouteRange_NextHopAddressType_Enum(intValue)
	obj.obj.NextHopAddressType = &enumValue

	return obj
}

// The IPv4 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// NextHopIpv4Address returns a string
func (obj *bgpV4RouteRange) NextHopIpv4Address() string {

	return *obj.obj.NextHopIpv4Address

}

// The IPv4 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// NextHopIpv4Address returns a string
func (obj *bgpV4RouteRange) HasNextHopIpv4Address() bool {
	return obj.obj.NextHopIpv4Address != nil
}

// The IPv4 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability should be enabled.
// SetNextHopIpv4Address sets the string value in the BgpV4RouteRange object
func (obj *bgpV4RouteRange) SetNextHopIpv4Address(value string) BgpV4RouteRange {

	obj.obj.NextHopIpv4Address = &value
	return obj
}

// The IPv6 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv6.
// NextHopIpv6Address returns a string
func (obj *bgpV4RouteRange) NextHopIpv6Address() string {

	return *obj.obj.NextHopIpv6Address

}

// The IPv6 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv6.
// NextHopIpv6Address returns a string
func (obj *bgpV4RouteRange) HasNextHopIpv6Address() bool {
	return obj.obj.NextHopIpv6Address != nil
}

// The IPv6 address of the next hop if the Nexthop Mode is manual and the Nexthop type is IPv6.
// SetNextHopIpv6Address sets the string value in the BgpV4RouteRange object
func (obj *bgpV4RouteRange) SetNextHopIpv6Address(value string) BgpV4RouteRange {

	obj.obj.NextHopIpv6Address = &value
	return obj
}

// description is TBD
// Advanced returns a BgpRouteAdvanced
func (obj *bgpV4RouteRange) Advanced() BgpRouteAdvanced {
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
func (obj *bgpV4RouteRange) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// description is TBD
// SetAdvanced sets the BgpRouteAdvanced value in the BgpV4RouteRange object
func (obj *bgpV4RouteRange) SetAdvanced(value BgpRouteAdvanced) BgpV4RouteRange {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// Optional community settings.
// Communities returns a []BgpCommunity
func (obj *bgpV4RouteRange) Communities() BgpV4RouteRangeBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.BgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpV4RouteRangeBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpV4RouteRangeBgpCommunityIter struct {
	obj               *bgpV4RouteRange
	bgpCommunitySlice []BgpCommunity
	fieldPtr          *[]*otg.BgpCommunity
}

func newBgpV4RouteRangeBgpCommunityIter(ptr *[]*otg.BgpCommunity) BgpV4RouteRangeBgpCommunityIter {
	return &bgpV4RouteRangeBgpCommunityIter{fieldPtr: ptr}
}

type BgpV4RouteRangeBgpCommunityIter interface {
	setMsg(*bgpV4RouteRange) BgpV4RouteRangeBgpCommunityIter
	Items() []BgpCommunity
	Add() BgpCommunity
	Append(items ...BgpCommunity) BgpV4RouteRangeBgpCommunityIter
	Set(index int, newObj BgpCommunity) BgpV4RouteRangeBgpCommunityIter
	Clear() BgpV4RouteRangeBgpCommunityIter
	clearHolderSlice() BgpV4RouteRangeBgpCommunityIter
	appendHolderSlice(item BgpCommunity) BgpV4RouteRangeBgpCommunityIter
}

func (obj *bgpV4RouteRangeBgpCommunityIter) setMsg(msg *bgpV4RouteRange) BgpV4RouteRangeBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4RouteRangeBgpCommunityIter) Items() []BgpCommunity {
	return obj.bgpCommunitySlice
}

func (obj *bgpV4RouteRangeBgpCommunityIter) Add() BgpCommunity {
	newObj := &otg.BgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4RouteRangeBgpCommunityIter) Append(items ...BgpCommunity) BgpV4RouteRangeBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV4RouteRangeBgpCommunityIter) Set(index int, newObj BgpCommunity) BgpV4RouteRangeBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV4RouteRangeBgpCommunityIter) Clear() BgpV4RouteRangeBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpCommunity{}
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpV4RouteRangeBgpCommunityIter) clearHolderSlice() BgpV4RouteRangeBgpCommunityIter {
	if len(obj.bgpCommunitySlice) > 0 {
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpV4RouteRangeBgpCommunityIter) appendHolderSlice(item BgpCommunity) BgpV4RouteRangeBgpCommunityIter {
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a BgpAsPath
func (obj *bgpV4RouteRange) AsPath() BgpAsPath {
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
func (obj *bgpV4RouteRange) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the BgpAsPath value in the BgpV4RouteRange object
func (obj *bgpV4RouteRange) SetAsPath(value BgpAsPath) BgpV4RouteRange {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// description is TBD
// AddPath returns a BgpAddPath
func (obj *bgpV4RouteRange) AddPath() BgpAddPath {
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
func (obj *bgpV4RouteRange) HasAddPath() bool {
	return obj.obj.AddPath != nil
}

// description is TBD
// SetAddPath sets the BgpAddPath value in the BgpV4RouteRange object
func (obj *bgpV4RouteRange) SetAddPath(value BgpAddPath) BgpV4RouteRange {

	obj.addPathHolder = nil
	obj.obj.AddPath = value.msg()

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpV4RouteRange) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpV4RouteRange object
func (obj *bgpV4RouteRange) SetName(value string) BgpV4RouteRange {

	obj.obj.Name = &value
	return obj
}

// Deprecated: This property is deprecated in favor of property extended_communities
//
// Deprecated: This property is deprecated in favor of property extended_communities
//
// Optional Extended Community settings. The Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an 8-Bytes value. It is divided into two main parts. The first 2 Bytes of the community encode a type and sub-type fields and the last 6 Bytes carry a unique set of data in a format defined by the type and sub-type field. Extended communities provide a larger  range for grouping or categorizing communities. When type is administrator_as_2octet or administrator_as_4octet, the valid sub types are route target and origin. The valid value for  administrator_as_2octet and administrator_as_4octet type is either two byte AS followed by four byte local administrator id or four byte AS followed by two  byte local administrator id.  When type is administrator_ipv4_address the valid sub types are route target and origin. The valid value for  administrator_ipv4_address is a four byte IPv4 address followed by a two byte local administrator id.  When type is opaque, valid sub types are color and encapsulation. When sub type is color, first two bytes of the value field contain flags and last four bytes  contains the value of the color. When sub type is encapsulation the first four bytes of value field are reserved and last two bytes carries the tunnel type from  IANA's "ETHER TYPES" registry e.g IPv4 (protocol type = 0x0800), IPv6 (protocol type = 0x86dd), and MPLS (protocol type = 0x8847). When type is administrator_as_2octet_link_bandwidth the valid sub type is extended_bandwidth. The first two bytes of the value field contains the AS number and the last four bytes contains the bandwidth in IEEE floating point format.  When type is evpn the valid subtype is mac_address. In the value field the low-order bit of the first byte(Flags) is defined as the "Sticky/static" flag and may be set to 1, indicating the MAC address is static and cannot move. The second byte is reserved and the  last four bytes contain the sequence number which is used to ensure that PEs retain the correct MAC/IP Advertisement route when multiple updates  occur for the same MAC address.  Note evpn type is defined mainly for use with evpn route updates and not for IPv4 and IPv6 route updates.
// ExtCommunities returns a []BgpExtCommunity
func (obj *bgpV4RouteRange) ExtCommunities() BgpV4RouteRangeBgpExtCommunityIter {
	if len(obj.obj.ExtCommunities) == 0 {
		obj.obj.ExtCommunities = []*otg.BgpExtCommunity{}
	}
	if obj.extCommunitiesHolder == nil {
		obj.extCommunitiesHolder = newBgpV4RouteRangeBgpExtCommunityIter(&obj.obj.ExtCommunities).setMsg(obj)
	}
	return obj.extCommunitiesHolder
}

type bgpV4RouteRangeBgpExtCommunityIter struct {
	obj                  *bgpV4RouteRange
	bgpExtCommunitySlice []BgpExtCommunity
	fieldPtr             *[]*otg.BgpExtCommunity
}

func newBgpV4RouteRangeBgpExtCommunityIter(ptr *[]*otg.BgpExtCommunity) BgpV4RouteRangeBgpExtCommunityIter {
	return &bgpV4RouteRangeBgpExtCommunityIter{fieldPtr: ptr}
}

type BgpV4RouteRangeBgpExtCommunityIter interface {
	setMsg(*bgpV4RouteRange) BgpV4RouteRangeBgpExtCommunityIter
	Items() []BgpExtCommunity
	Add() BgpExtCommunity
	Append(items ...BgpExtCommunity) BgpV4RouteRangeBgpExtCommunityIter
	Set(index int, newObj BgpExtCommunity) BgpV4RouteRangeBgpExtCommunityIter
	Clear() BgpV4RouteRangeBgpExtCommunityIter
	clearHolderSlice() BgpV4RouteRangeBgpExtCommunityIter
	appendHolderSlice(item BgpExtCommunity) BgpV4RouteRangeBgpExtCommunityIter
}

func (obj *bgpV4RouteRangeBgpExtCommunityIter) setMsg(msg *bgpV4RouteRange) BgpV4RouteRangeBgpExtCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4RouteRangeBgpExtCommunityIter) Items() []BgpExtCommunity {
	return obj.bgpExtCommunitySlice
}

func (obj *bgpV4RouteRangeBgpExtCommunityIter) Add() BgpExtCommunity {
	newObj := &otg.BgpExtCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4RouteRangeBgpExtCommunityIter) Append(items ...BgpExtCommunity) BgpV4RouteRangeBgpExtCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV4RouteRangeBgpExtCommunityIter) Set(index int, newObj BgpExtCommunity) BgpV4RouteRangeBgpExtCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV4RouteRangeBgpExtCommunityIter) Clear() BgpV4RouteRangeBgpExtCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtCommunity{}
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpV4RouteRangeBgpExtCommunityIter) clearHolderSlice() BgpV4RouteRangeBgpExtCommunityIter {
	if len(obj.bgpExtCommunitySlice) > 0 {
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpV4RouteRangeBgpExtCommunityIter) appendHolderSlice(item BgpExtCommunity) BgpV4RouteRangeBgpExtCommunityIter {
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	return obj
}

// Optional Extended Community settings. The Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an eight byte value. It is divided into two main parts. The first two bytes of the community encode a type and sub-type fields and the last six bytes carry a unique set of data in a format defined by the type and sub-type field. Extended communities provide a larger range for grouping or categorizing communities.
// ExtendedCommunities returns a []BgpExtendedCommunity
func (obj *bgpV4RouteRange) ExtendedCommunities() BgpV4RouteRangeBgpExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.BgpExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpV4RouteRangeBgpExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpV4RouteRangeBgpExtendedCommunityIter struct {
	obj                       *bgpV4RouteRange
	bgpExtendedCommunitySlice []BgpExtendedCommunity
	fieldPtr                  *[]*otg.BgpExtendedCommunity
}

func newBgpV4RouteRangeBgpExtendedCommunityIter(ptr *[]*otg.BgpExtendedCommunity) BgpV4RouteRangeBgpExtendedCommunityIter {
	return &bgpV4RouteRangeBgpExtendedCommunityIter{fieldPtr: ptr}
}

type BgpV4RouteRangeBgpExtendedCommunityIter interface {
	setMsg(*bgpV4RouteRange) BgpV4RouteRangeBgpExtendedCommunityIter
	Items() []BgpExtendedCommunity
	Add() BgpExtendedCommunity
	Append(items ...BgpExtendedCommunity) BgpV4RouteRangeBgpExtendedCommunityIter
	Set(index int, newObj BgpExtendedCommunity) BgpV4RouteRangeBgpExtendedCommunityIter
	Clear() BgpV4RouteRangeBgpExtendedCommunityIter
	clearHolderSlice() BgpV4RouteRangeBgpExtendedCommunityIter
	appendHolderSlice(item BgpExtendedCommunity) BgpV4RouteRangeBgpExtendedCommunityIter
}

func (obj *bgpV4RouteRangeBgpExtendedCommunityIter) setMsg(msg *bgpV4RouteRange) BgpV4RouteRangeBgpExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4RouteRangeBgpExtendedCommunityIter) Items() []BgpExtendedCommunity {
	return obj.bgpExtendedCommunitySlice
}

func (obj *bgpV4RouteRangeBgpExtendedCommunityIter) Add() BgpExtendedCommunity {
	newObj := &otg.BgpExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4RouteRangeBgpExtendedCommunityIter) Append(items ...BgpExtendedCommunity) BgpV4RouteRangeBgpExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV4RouteRangeBgpExtendedCommunityIter) Set(index int, newObj BgpExtendedCommunity) BgpV4RouteRangeBgpExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV4RouteRangeBgpExtendedCommunityIter) Clear() BgpV4RouteRangeBgpExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtendedCommunity{}
		obj.bgpExtendedCommunitySlice = []BgpExtendedCommunity{}
	}
	return obj
}
func (obj *bgpV4RouteRangeBgpExtendedCommunityIter) clearHolderSlice() BgpV4RouteRangeBgpExtendedCommunityIter {
	if len(obj.bgpExtendedCommunitySlice) > 0 {
		obj.bgpExtendedCommunitySlice = []BgpExtendedCommunity{}
	}
	return obj
}
func (obj *bgpV4RouteRangeBgpExtendedCommunityIter) appendHolderSlice(item BgpExtendedCommunity) BgpV4RouteRangeBgpExtendedCommunityIter {
	obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, item)
	return obj
}

func (obj *bgpV4RouteRange) validateObj(vObj *validation, set_default bool) {
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
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpV4RouteRange.NextHopIpv4Address"))
		}

	}

	if obj.obj.NextHopIpv6Address != nil {

		err := obj.validateIpv6(obj.NextHopIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpV4RouteRange.NextHopIpv6Address"))
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
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpV4RouteRange")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"BgpV4RouteRange.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if len(obj.obj.ExtCommunities) != 0 {
		obj.addWarnings("ExtCommunities property in schema BgpV4RouteRange is deprecated, This property is deprecated in favor of property extended_communities")

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

func (obj *bgpV4RouteRange) setDefault() {
	if obj.obj.NextHopMode == nil {
		obj.SetNextHopMode(BgpV4RouteRangeNextHopMode.LOCAL_IP)

	}
	if obj.obj.NextHopAddressType == nil {
		obj.SetNextHopAddressType(BgpV4RouteRangeNextHopAddressType.IPV4)

	}
	if obj.obj.NextHopIpv4Address == nil {
		obj.SetNextHopIpv4Address("0.0.0.0")
	}
	if obj.obj.NextHopIpv6Address == nil {
		obj.SetNextHopIpv6Address("::0")
	}

}
