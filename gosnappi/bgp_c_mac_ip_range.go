package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpCMacIpRange *****
type bgpCMacIpRange struct {
	validation
	obj                  *otg.BgpCMacIpRange
	marshaller           marshalBgpCMacIpRange
	unMarshaller         unMarshalBgpCMacIpRange
	macAddressesHolder   MACRouteAddress
	ipv4AddressesHolder  V4RouteAddress
	ipv6AddressesHolder  V6RouteAddress
	advancedHolder       BgpRouteAdvanced
	communitiesHolder    BgpCMacIpRangeBgpCommunityIter
	extCommunitiesHolder BgpCMacIpRangeBgpExtCommunityIter
	asPathHolder         BgpAsPath
}

func NewBgpCMacIpRange() BgpCMacIpRange {
	obj := bgpCMacIpRange{obj: &otg.BgpCMacIpRange{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpCMacIpRange) msg() *otg.BgpCMacIpRange {
	return obj.obj
}

func (obj *bgpCMacIpRange) setMsg(msg *otg.BgpCMacIpRange) BgpCMacIpRange {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpCMacIpRange struct {
	obj *bgpCMacIpRange
}

type marshalBgpCMacIpRange interface {
	// ToProto marshals BgpCMacIpRange to protobuf object *otg.BgpCMacIpRange
	ToProto() (*otg.BgpCMacIpRange, error)
	// ToPbText marshals BgpCMacIpRange to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpCMacIpRange to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpCMacIpRange to JSON text
	ToJson() (string, error)
}

type unMarshalbgpCMacIpRange struct {
	obj *bgpCMacIpRange
}

type unMarshalBgpCMacIpRange interface {
	// FromProto unmarshals BgpCMacIpRange from protobuf object *otg.BgpCMacIpRange
	FromProto(msg *otg.BgpCMacIpRange) (BgpCMacIpRange, error)
	// FromPbText unmarshals BgpCMacIpRange from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpCMacIpRange from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpCMacIpRange from JSON text
	FromJson(value string) error
}

func (obj *bgpCMacIpRange) Marshal() marshalBgpCMacIpRange {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpCMacIpRange{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpCMacIpRange) Unmarshal() unMarshalBgpCMacIpRange {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpCMacIpRange{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpCMacIpRange) ToProto() (*otg.BgpCMacIpRange, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpCMacIpRange) FromProto(msg *otg.BgpCMacIpRange) (BgpCMacIpRange, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpCMacIpRange) ToPbText() (string, error) {
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

func (m *unMarshalbgpCMacIpRange) FromPbText(value string) error {
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

func (m *marshalbgpCMacIpRange) ToYaml() (string, error) {
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

func (m *unMarshalbgpCMacIpRange) FromYaml(value string) error {
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

func (m *marshalbgpCMacIpRange) ToJson() (string, error) {
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

func (m *unMarshalbgpCMacIpRange) FromJson(value string) error {
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

func (obj *bgpCMacIpRange) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpCMacIpRange) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpCMacIpRange) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpCMacIpRange) Clone() (BgpCMacIpRange, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpCMacIpRange()
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

func (obj *bgpCMacIpRange) setNil() {
	obj.macAddressesHolder = nil
	obj.ipv4AddressesHolder = nil
	obj.ipv6AddressesHolder = nil
	obj.advancedHolder = nil
	obj.communitiesHolder = nil
	obj.extCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpCMacIpRange is configuration for MAC/IP Ranges per Broadcast Domain.
//
// Advertises following route -
//
// Type 2 - MAC/IP Advertisement Route.
type BgpCMacIpRange interface {
	Validation
	// msg marshals BgpCMacIpRange to protobuf object *otg.BgpCMacIpRange
	// and doesn't set defaults
	msg() *otg.BgpCMacIpRange
	// setMsg unmarshals BgpCMacIpRange from protobuf object *otg.BgpCMacIpRange
	// and doesn't set defaults
	setMsg(*otg.BgpCMacIpRange) BgpCMacIpRange
	// provides marshal interface
	Marshal() marshalBgpCMacIpRange
	// provides unmarshal interface
	Unmarshal() unMarshalBgpCMacIpRange
	// validate validates BgpCMacIpRange
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpCMacIpRange, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MacAddresses returns MACRouteAddress, set in BgpCMacIpRange.
	// MACRouteAddress is a container for MAC route addresses.
	MacAddresses() MACRouteAddress
	// SetMacAddresses assigns MACRouteAddress provided by user to BgpCMacIpRange.
	// MACRouteAddress is a container for MAC route addresses.
	SetMacAddresses(value MACRouteAddress) BgpCMacIpRange
	// HasMacAddresses checks if MacAddresses has been set in BgpCMacIpRange
	HasMacAddresses() bool
	// L2Vni returns uint32, set in BgpCMacIpRange.
	L2Vni() uint32
	// SetL2Vni assigns uint32 provided by user to BgpCMacIpRange
	SetL2Vni(value uint32) BgpCMacIpRange
	// HasL2Vni checks if L2Vni has been set in BgpCMacIpRange
	HasL2Vni() bool
	// Ipv4Addresses returns V4RouteAddress, set in BgpCMacIpRange.
	// V4RouteAddress is a container for IPv4 route addresses.
	Ipv4Addresses() V4RouteAddress
	// SetIpv4Addresses assigns V4RouteAddress provided by user to BgpCMacIpRange.
	// V4RouteAddress is a container for IPv4 route addresses.
	SetIpv4Addresses(value V4RouteAddress) BgpCMacIpRange
	// HasIpv4Addresses checks if Ipv4Addresses has been set in BgpCMacIpRange
	HasIpv4Addresses() bool
	// Ipv6Addresses returns V6RouteAddress, set in BgpCMacIpRange.
	// V6RouteAddress is a container for IPv6 route addresses.
	Ipv6Addresses() V6RouteAddress
	// SetIpv6Addresses assigns V6RouteAddress provided by user to BgpCMacIpRange.
	// V6RouteAddress is a container for IPv6 route addresses.
	SetIpv6Addresses(value V6RouteAddress) BgpCMacIpRange
	// HasIpv6Addresses checks if Ipv6Addresses has been set in BgpCMacIpRange
	HasIpv6Addresses() bool
	// L3Vni returns uint32, set in BgpCMacIpRange.
	L3Vni() uint32
	// SetL3Vni assigns uint32 provided by user to BgpCMacIpRange
	SetL3Vni(value uint32) BgpCMacIpRange
	// HasL3Vni checks if L3Vni has been set in BgpCMacIpRange
	HasL3Vni() bool
	// IncludeDefaultGateway returns bool, set in BgpCMacIpRange.
	IncludeDefaultGateway() bool
	// SetIncludeDefaultGateway assigns bool provided by user to BgpCMacIpRange
	SetIncludeDefaultGateway(value bool) BgpCMacIpRange
	// HasIncludeDefaultGateway checks if IncludeDefaultGateway has been set in BgpCMacIpRange
	HasIncludeDefaultGateway() bool
	// Advanced returns BgpRouteAdvanced, set in BgpCMacIpRange.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	Advanced() BgpRouteAdvanced
	// SetAdvanced assigns BgpRouteAdvanced provided by user to BgpCMacIpRange.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	SetAdvanced(value BgpRouteAdvanced) BgpCMacIpRange
	// HasAdvanced checks if Advanced has been set in BgpCMacIpRange
	HasAdvanced() bool
	// Communities returns BgpCMacIpRangeBgpCommunityIterIter, set in BgpCMacIpRange
	Communities() BgpCMacIpRangeBgpCommunityIter
	// ExtCommunities returns BgpCMacIpRangeBgpExtCommunityIterIter, set in BgpCMacIpRange
	ExtCommunities() BgpCMacIpRangeBgpExtCommunityIter
	// AsPath returns BgpAsPath, set in BgpCMacIpRange.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	AsPath() BgpAsPath
	// SetAsPath assigns BgpAsPath provided by user to BgpCMacIpRange.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	SetAsPath(value BgpAsPath) BgpCMacIpRange
	// HasAsPath checks if AsPath has been set in BgpCMacIpRange
	HasAsPath() bool
	// Name returns string, set in BgpCMacIpRange.
	Name() string
	// SetName assigns string provided by user to BgpCMacIpRange
	SetName(value string) BgpCMacIpRange
	setNil()
}

// Host MAC address range per Broadcast Domain.
// MacAddresses returns a MACRouteAddress
func (obj *bgpCMacIpRange) MacAddresses() MACRouteAddress {
	if obj.obj.MacAddresses == nil {
		obj.obj.MacAddresses = NewMACRouteAddress().msg()
	}
	if obj.macAddressesHolder == nil {
		obj.macAddressesHolder = &mACRouteAddress{obj: obj.obj.MacAddresses}
	}
	return obj.macAddressesHolder
}

// Host MAC address range per Broadcast Domain.
// MacAddresses returns a MACRouteAddress
func (obj *bgpCMacIpRange) HasMacAddresses() bool {
	return obj.obj.MacAddresses != nil
}

// Host MAC address range per Broadcast Domain.
// SetMacAddresses sets the MACRouteAddress value in the BgpCMacIpRange object
func (obj *bgpCMacIpRange) SetMacAddresses(value MACRouteAddress) BgpCMacIpRange {

	obj.macAddressesHolder = nil
	obj.obj.MacAddresses = value.msg()

	return obj
}

// Layer 2 Virtual Network Identifier (L2VNI) to be advertised with MAC/IP Advertisement Route (Type 2)
// L2Vni returns a uint32
func (obj *bgpCMacIpRange) L2Vni() uint32 {

	return *obj.obj.L2Vni

}

// Layer 2 Virtual Network Identifier (L2VNI) to be advertised with MAC/IP Advertisement Route (Type 2)
// L2Vni returns a uint32
func (obj *bgpCMacIpRange) HasL2Vni() bool {
	return obj.obj.L2Vni != nil
}

// Layer 2 Virtual Network Identifier (L2VNI) to be advertised with MAC/IP Advertisement Route (Type 2)
// SetL2Vni sets the uint32 value in the BgpCMacIpRange object
func (obj *bgpCMacIpRange) SetL2Vni(value uint32) BgpCMacIpRange {

	obj.obj.L2Vni = &value
	return obj
}

// Host IPv4 address range per Broadcast Domain.
// Ipv4Addresses returns a V4RouteAddress
func (obj *bgpCMacIpRange) Ipv4Addresses() V4RouteAddress {
	if obj.obj.Ipv4Addresses == nil {
		obj.obj.Ipv4Addresses = NewV4RouteAddress().msg()
	}
	if obj.ipv4AddressesHolder == nil {
		obj.ipv4AddressesHolder = &v4RouteAddress{obj: obj.obj.Ipv4Addresses}
	}
	return obj.ipv4AddressesHolder
}

// Host IPv4 address range per Broadcast Domain.
// Ipv4Addresses returns a V4RouteAddress
func (obj *bgpCMacIpRange) HasIpv4Addresses() bool {
	return obj.obj.Ipv4Addresses != nil
}

// Host IPv4 address range per Broadcast Domain.
// SetIpv4Addresses sets the V4RouteAddress value in the BgpCMacIpRange object
func (obj *bgpCMacIpRange) SetIpv4Addresses(value V4RouteAddress) BgpCMacIpRange {

	obj.ipv4AddressesHolder = nil
	obj.obj.Ipv4Addresses = value.msg()

	return obj
}

// Host IPv6 address range per Broadcast Domain.
// Ipv6Addresses returns a V6RouteAddress
func (obj *bgpCMacIpRange) Ipv6Addresses() V6RouteAddress {
	if obj.obj.Ipv6Addresses == nil {
		obj.obj.Ipv6Addresses = NewV6RouteAddress().msg()
	}
	if obj.ipv6AddressesHolder == nil {
		obj.ipv6AddressesHolder = &v6RouteAddress{obj: obj.obj.Ipv6Addresses}
	}
	return obj.ipv6AddressesHolder
}

// Host IPv6 address range per Broadcast Domain.
// Ipv6Addresses returns a V6RouteAddress
func (obj *bgpCMacIpRange) HasIpv6Addresses() bool {
	return obj.obj.Ipv6Addresses != nil
}

// Host IPv6 address range per Broadcast Domain.
// SetIpv6Addresses sets the V6RouteAddress value in the BgpCMacIpRange object
func (obj *bgpCMacIpRange) SetIpv6Addresses(value V6RouteAddress) BgpCMacIpRange {

	obj.ipv6AddressesHolder = nil
	obj.obj.Ipv6Addresses = value.msg()

	return obj
}

// Layer 3 Virtual Network Identifier (L3VNI) to be advertised with MAC/IP Advertisement Route (Type 2).
// L3Vni returns a uint32
func (obj *bgpCMacIpRange) L3Vni() uint32 {

	return *obj.obj.L3Vni

}

// Layer 3 Virtual Network Identifier (L3VNI) to be advertised with MAC/IP Advertisement Route (Type 2).
// L3Vni returns a uint32
func (obj *bgpCMacIpRange) HasL3Vni() bool {
	return obj.obj.L3Vni != nil
}

// Layer 3 Virtual Network Identifier (L3VNI) to be advertised with MAC/IP Advertisement Route (Type 2).
// SetL3Vni sets the uint32 value in the BgpCMacIpRange object
func (obj *bgpCMacIpRange) SetL3Vni(value uint32) BgpCMacIpRange {

	obj.obj.L3Vni = &value
	return obj
}

// Include default Gateway Extended Community in MAC/IP Advertisement Route (Type 2).
// IncludeDefaultGateway returns a bool
func (obj *bgpCMacIpRange) IncludeDefaultGateway() bool {

	return *obj.obj.IncludeDefaultGateway

}

// Include default Gateway Extended Community in MAC/IP Advertisement Route (Type 2).
// IncludeDefaultGateway returns a bool
func (obj *bgpCMacIpRange) HasIncludeDefaultGateway() bool {
	return obj.obj.IncludeDefaultGateway != nil
}

// Include default Gateway Extended Community in MAC/IP Advertisement Route (Type 2).
// SetIncludeDefaultGateway sets the bool value in the BgpCMacIpRange object
func (obj *bgpCMacIpRange) SetIncludeDefaultGateway(value bool) BgpCMacIpRange {

	obj.obj.IncludeDefaultGateway = &value
	return obj
}

// description is TBD
// Advanced returns a BgpRouteAdvanced
func (obj *bgpCMacIpRange) Advanced() BgpRouteAdvanced {
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
func (obj *bgpCMacIpRange) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// description is TBD
// SetAdvanced sets the BgpRouteAdvanced value in the BgpCMacIpRange object
func (obj *bgpCMacIpRange) SetAdvanced(value BgpRouteAdvanced) BgpCMacIpRange {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// Optional community settings.
// Communities returns a []BgpCommunity
func (obj *bgpCMacIpRange) Communities() BgpCMacIpRangeBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.BgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpCMacIpRangeBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpCMacIpRangeBgpCommunityIter struct {
	obj               *bgpCMacIpRange
	bgpCommunitySlice []BgpCommunity
	fieldPtr          *[]*otg.BgpCommunity
}

func newBgpCMacIpRangeBgpCommunityIter(ptr *[]*otg.BgpCommunity) BgpCMacIpRangeBgpCommunityIter {
	return &bgpCMacIpRangeBgpCommunityIter{fieldPtr: ptr}
}

type BgpCMacIpRangeBgpCommunityIter interface {
	setMsg(*bgpCMacIpRange) BgpCMacIpRangeBgpCommunityIter
	Items() []BgpCommunity
	Add() BgpCommunity
	Append(items ...BgpCommunity) BgpCMacIpRangeBgpCommunityIter
	Set(index int, newObj BgpCommunity) BgpCMacIpRangeBgpCommunityIter
	Clear() BgpCMacIpRangeBgpCommunityIter
	clearHolderSlice() BgpCMacIpRangeBgpCommunityIter
	appendHolderSlice(item BgpCommunity) BgpCMacIpRangeBgpCommunityIter
}

func (obj *bgpCMacIpRangeBgpCommunityIter) setMsg(msg *bgpCMacIpRange) BgpCMacIpRangeBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpCMacIpRangeBgpCommunityIter) Items() []BgpCommunity {
	return obj.bgpCommunitySlice
}

func (obj *bgpCMacIpRangeBgpCommunityIter) Add() BgpCommunity {
	newObj := &otg.BgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpCMacIpRangeBgpCommunityIter) Append(items ...BgpCommunity) BgpCMacIpRangeBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpCMacIpRangeBgpCommunityIter) Set(index int, newObj BgpCommunity) BgpCMacIpRangeBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpCMacIpRangeBgpCommunityIter) Clear() BgpCMacIpRangeBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpCommunity{}
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpCMacIpRangeBgpCommunityIter) clearHolderSlice() BgpCMacIpRangeBgpCommunityIter {
	if len(obj.bgpCommunitySlice) > 0 {
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpCMacIpRangeBgpCommunityIter) appendHolderSlice(item BgpCommunity) BgpCMacIpRangeBgpCommunityIter {
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	return obj
}

// Optional Extended Community settings. The Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an 8-Bytes value. It is divided into two main parts. The first 2 Bytes of the community encode a type and sub-type fields and the last 6 Bytes carry a unique set of data in a format defined by the type and sub-type field. Extended communities provide a larger  range for grouping or categorizing communities. When type is administrator_as_2octet or administrator_as_4octet, the valid sub types are route target and origin. The valid value for  administrator_as_2octet and administrator_as_4octet type is either two byte AS followed by four byte local administrator id or four byte AS followed by two  byte local administrator id.  When type is administrator_ipv4_address the valid sub types are route target and origin. The valid value for  administrator_ipv4_address is a four byte IPv4 address followed by a two byte local administrator id.  When type is opaque, valid sub types are color and encapsulation. When sub type is color, first two bytes of the value field contain flags and last four bytes  contains the value of the color. When sub type is encapsulation the first four bytes of value field are reserved and last two bytes carries the tunnel type from  IANA's "ETHER TYPES" registry e.g IPv4 (protocol type = 0x0800), IPv6 (protocol type = 0x86dd), and MPLS (protocol type = 0x8847). When type is administrator_as_2octet_link_bandwidth the valid sub type is extended_bandwidth. The first two bytes of the value field contains the AS number and the last four bytes contains the bandwidth in IEEE floating point format.  When type is evpn the valid subtype is mac_address. In the value field the low-order bit of the first byte(Flags) is defined as the "Sticky/static" flag and may be set to 1, indicating the MAC address is static and cannot move. The second byte is reserved and the  last four bytes contain the sequence number which is used to ensure that PEs retain the correct MAC/IP Advertisement route when multiple updates  occur for the same MAC address.
// ExtCommunities returns a []BgpExtCommunity
func (obj *bgpCMacIpRange) ExtCommunities() BgpCMacIpRangeBgpExtCommunityIter {
	if len(obj.obj.ExtCommunities) == 0 {
		obj.obj.ExtCommunities = []*otg.BgpExtCommunity{}
	}
	if obj.extCommunitiesHolder == nil {
		obj.extCommunitiesHolder = newBgpCMacIpRangeBgpExtCommunityIter(&obj.obj.ExtCommunities).setMsg(obj)
	}
	return obj.extCommunitiesHolder
}

type bgpCMacIpRangeBgpExtCommunityIter struct {
	obj                  *bgpCMacIpRange
	bgpExtCommunitySlice []BgpExtCommunity
	fieldPtr             *[]*otg.BgpExtCommunity
}

func newBgpCMacIpRangeBgpExtCommunityIter(ptr *[]*otg.BgpExtCommunity) BgpCMacIpRangeBgpExtCommunityIter {
	return &bgpCMacIpRangeBgpExtCommunityIter{fieldPtr: ptr}
}

type BgpCMacIpRangeBgpExtCommunityIter interface {
	setMsg(*bgpCMacIpRange) BgpCMacIpRangeBgpExtCommunityIter
	Items() []BgpExtCommunity
	Add() BgpExtCommunity
	Append(items ...BgpExtCommunity) BgpCMacIpRangeBgpExtCommunityIter
	Set(index int, newObj BgpExtCommunity) BgpCMacIpRangeBgpExtCommunityIter
	Clear() BgpCMacIpRangeBgpExtCommunityIter
	clearHolderSlice() BgpCMacIpRangeBgpExtCommunityIter
	appendHolderSlice(item BgpExtCommunity) BgpCMacIpRangeBgpExtCommunityIter
}

func (obj *bgpCMacIpRangeBgpExtCommunityIter) setMsg(msg *bgpCMacIpRange) BgpCMacIpRangeBgpExtCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpCMacIpRangeBgpExtCommunityIter) Items() []BgpExtCommunity {
	return obj.bgpExtCommunitySlice
}

func (obj *bgpCMacIpRangeBgpExtCommunityIter) Add() BgpExtCommunity {
	newObj := &otg.BgpExtCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpCMacIpRangeBgpExtCommunityIter) Append(items ...BgpExtCommunity) BgpCMacIpRangeBgpExtCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	}
	return obj
}

func (obj *bgpCMacIpRangeBgpExtCommunityIter) Set(index int, newObj BgpExtCommunity) BgpCMacIpRangeBgpExtCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpCMacIpRangeBgpExtCommunityIter) Clear() BgpCMacIpRangeBgpExtCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtCommunity{}
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpCMacIpRangeBgpExtCommunityIter) clearHolderSlice() BgpCMacIpRangeBgpExtCommunityIter {
	if len(obj.bgpExtCommunitySlice) > 0 {
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpCMacIpRangeBgpExtCommunityIter) appendHolderSlice(item BgpExtCommunity) BgpCMacIpRangeBgpExtCommunityIter {
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	return obj
}

// Optional AS PATH settings.
// AsPath returns a BgpAsPath
func (obj *bgpCMacIpRange) AsPath() BgpAsPath {
	if obj.obj.AsPath == nil {
		obj.obj.AsPath = NewBgpAsPath().msg()
	}
	if obj.asPathHolder == nil {
		obj.asPathHolder = &bgpAsPath{obj: obj.obj.AsPath}
	}
	return obj.asPathHolder
}

// Optional AS PATH settings.
// AsPath returns a BgpAsPath
func (obj *bgpCMacIpRange) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// Optional AS PATH settings.
// SetAsPath sets the BgpAsPath value in the BgpCMacIpRange object
func (obj *bgpCMacIpRange) SetAsPath(value BgpAsPath) BgpCMacIpRange {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpCMacIpRange) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpCMacIpRange object
func (obj *bgpCMacIpRange) SetName(value string) BgpCMacIpRange {

	obj.obj.Name = &value
	return obj
}

func (obj *bgpCMacIpRange) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MacAddresses != nil {

		obj.MacAddresses().validateObj(vObj, set_default)
	}

	if obj.obj.L2Vni != nil {

		if *obj.obj.L2Vni > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpCMacIpRange.L2Vni <= 16777215 but Got %d", *obj.obj.L2Vni))
		}

	}

	if obj.obj.Ipv4Addresses != nil {

		obj.Ipv4Addresses().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv6Addresses != nil {

		obj.Ipv6Addresses().validateObj(vObj, set_default)
	}

	if obj.obj.L3Vni != nil {

		if *obj.obj.L3Vni > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpCMacIpRange.L3Vni <= 16777215 but Got %d", *obj.obj.L3Vni))
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

	if len(obj.obj.ExtCommunities) != 0 {

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

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(vObj, set_default)
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpCMacIpRange")
	}
}

func (obj *bgpCMacIpRange) setDefault() {
	if obj.obj.L2Vni == nil {
		obj.SetL2Vni(0)
	}
	if obj.obj.L3Vni == nil {
		obj.SetL3Vni(0)
	}
	if obj.obj.IncludeDefaultGateway == nil {
		obj.SetIncludeDefaultGateway(false)
	}

}
