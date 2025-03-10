package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteV4Policy *****
type bgpSrteV4Policy struct {
	validation
	obj                  *otg.BgpSrteV4Policy
	marshaller           marshalBgpSrteV4Policy
	unMarshaller         unMarshalBgpSrteV4Policy
	advancedHolder       BgpRouteAdvanced
	addPathHolder        BgpAddPath
	asPathHolder         BgpAsPath
	communitiesHolder    BgpSrteV4PolicyBgpCommunityIter
	extCommunitiesHolder BgpSrteV4PolicyBgpExtCommunityIter
	tunnelTlvsHolder     BgpSrteV4PolicyBgpSrteV4TunnelTlvIter
}

func NewBgpSrteV4Policy() BgpSrteV4Policy {
	obj := bgpSrteV4Policy{obj: &otg.BgpSrteV4Policy{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteV4Policy) msg() *otg.BgpSrteV4Policy {
	return obj.obj
}

func (obj *bgpSrteV4Policy) setMsg(msg *otg.BgpSrteV4Policy) BgpSrteV4Policy {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteV4Policy struct {
	obj *bgpSrteV4Policy
}

type marshalBgpSrteV4Policy interface {
	// ToProto marshals BgpSrteV4Policy to protobuf object *otg.BgpSrteV4Policy
	ToProto() (*otg.BgpSrteV4Policy, error)
	// ToPbText marshals BgpSrteV4Policy to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteV4Policy to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteV4Policy to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpSrteV4Policy to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpSrteV4Policy struct {
	obj *bgpSrteV4Policy
}

type unMarshalBgpSrteV4Policy interface {
	// FromProto unmarshals BgpSrteV4Policy from protobuf object *otg.BgpSrteV4Policy
	FromProto(msg *otg.BgpSrteV4Policy) (BgpSrteV4Policy, error)
	// FromPbText unmarshals BgpSrteV4Policy from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteV4Policy from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteV4Policy from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteV4Policy) Marshal() marshalBgpSrteV4Policy {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteV4Policy{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteV4Policy) Unmarshal() unMarshalBgpSrteV4Policy {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteV4Policy{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteV4Policy) ToProto() (*otg.BgpSrteV4Policy, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteV4Policy) FromProto(msg *otg.BgpSrteV4Policy) (BgpSrteV4Policy, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteV4Policy) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteV4Policy) FromPbText(value string) error {
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

func (m *marshalbgpSrteV4Policy) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteV4Policy) FromYaml(value string) error {
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

func (m *marshalbgpSrteV4Policy) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalbgpSrteV4Policy) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteV4Policy) FromJson(value string) error {
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

func (obj *bgpSrteV4Policy) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteV4Policy) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteV4Policy) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteV4Policy) Clone() (BgpSrteV4Policy, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteV4Policy()
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

func (obj *bgpSrteV4Policy) setNil() {
	obj.advancedHolder = nil
	obj.addPathHolder = nil
	obj.asPathHolder = nil
	obj.communitiesHolder = nil
	obj.extCommunitiesHolder = nil
	obj.tunnelTlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteV4Policy is configuration for BGP Segment Routing Traffic Engineering(SRTE)
// policy.

type BgpSrteV4Policy interface {
	Validation
	// msg marshals BgpSrteV4Policy to protobuf object *otg.BgpSrteV4Policy
	// and doesn't set defaults
	msg() *otg.BgpSrteV4Policy
	// setMsg unmarshals BgpSrteV4Policy from protobuf object *otg.BgpSrteV4Policy
	// and doesn't set defaults
	setMsg(*otg.BgpSrteV4Policy) BgpSrteV4Policy
	// provides marshal interface
	Marshal() marshalBgpSrteV4Policy
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteV4Policy
	// validate validates BgpSrteV4Policy
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteV4Policy, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Distinguisher returns uint32, set in BgpSrteV4Policy.
	Distinguisher() uint32
	// SetDistinguisher assigns uint32 provided by user to BgpSrteV4Policy
	SetDistinguisher(value uint32) BgpSrteV4Policy
	// HasDistinguisher checks if Distinguisher has been set in BgpSrteV4Policy
	HasDistinguisher() bool
	// Color returns uint32, set in BgpSrteV4Policy.
	Color() uint32
	// SetColor assigns uint32 provided by user to BgpSrteV4Policy
	SetColor(value uint32) BgpSrteV4Policy
	// HasColor checks if Color has been set in BgpSrteV4Policy
	HasColor() bool
	// Ipv4Endpoint returns string, set in BgpSrteV4Policy.
	Ipv4Endpoint() string
	// SetIpv4Endpoint assigns string provided by user to BgpSrteV4Policy
	SetIpv4Endpoint(value string) BgpSrteV4Policy
	// NextHopMode returns BgpSrteV4PolicyNextHopModeEnum, set in BgpSrteV4Policy
	NextHopMode() BgpSrteV4PolicyNextHopModeEnum
	// SetNextHopMode assigns BgpSrteV4PolicyNextHopModeEnum provided by user to BgpSrteV4Policy
	SetNextHopMode(value BgpSrteV4PolicyNextHopModeEnum) BgpSrteV4Policy
	// HasNextHopMode checks if NextHopMode has been set in BgpSrteV4Policy
	HasNextHopMode() bool
	// NextHopAddressType returns BgpSrteV4PolicyNextHopAddressTypeEnum, set in BgpSrteV4Policy
	NextHopAddressType() BgpSrteV4PolicyNextHopAddressTypeEnum
	// SetNextHopAddressType assigns BgpSrteV4PolicyNextHopAddressTypeEnum provided by user to BgpSrteV4Policy
	SetNextHopAddressType(value BgpSrteV4PolicyNextHopAddressTypeEnum) BgpSrteV4Policy
	// HasNextHopAddressType checks if NextHopAddressType has been set in BgpSrteV4Policy
	HasNextHopAddressType() bool
	// NextHopIpv4Address returns string, set in BgpSrteV4Policy.
	NextHopIpv4Address() string
	// SetNextHopIpv4Address assigns string provided by user to BgpSrteV4Policy
	SetNextHopIpv4Address(value string) BgpSrteV4Policy
	// HasNextHopIpv4Address checks if NextHopIpv4Address has been set in BgpSrteV4Policy
	HasNextHopIpv4Address() bool
	// NextHopIpv6Address returns string, set in BgpSrteV4Policy.
	NextHopIpv6Address() string
	// SetNextHopIpv6Address assigns string provided by user to BgpSrteV4Policy
	SetNextHopIpv6Address(value string) BgpSrteV4Policy
	// HasNextHopIpv6Address checks if NextHopIpv6Address has been set in BgpSrteV4Policy
	HasNextHopIpv6Address() bool
	// Advanced returns BgpRouteAdvanced, set in BgpSrteV4Policy.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	Advanced() BgpRouteAdvanced
	// SetAdvanced assigns BgpRouteAdvanced provided by user to BgpSrteV4Policy.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	SetAdvanced(value BgpRouteAdvanced) BgpSrteV4Policy
	// HasAdvanced checks if Advanced has been set in BgpSrteV4Policy
	HasAdvanced() bool
	// AddPath returns BgpAddPath, set in BgpSrteV4Policy.
	// BgpAddPath is the BGP Additional Paths feature is a BGP extension that allows the  advertisement of multiple paths for the same prefix without the new  paths implicitly replacing any previous paths.
	AddPath() BgpAddPath
	// SetAddPath assigns BgpAddPath provided by user to BgpSrteV4Policy.
	// BgpAddPath is the BGP Additional Paths feature is a BGP extension that allows the  advertisement of multiple paths for the same prefix without the new  paths implicitly replacing any previous paths.
	SetAddPath(value BgpAddPath) BgpSrteV4Policy
	// HasAddPath checks if AddPath has been set in BgpSrteV4Policy
	HasAddPath() bool
	// AsPath returns BgpAsPath, set in BgpSrteV4Policy.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	AsPath() BgpAsPath
	// SetAsPath assigns BgpAsPath provided by user to BgpSrteV4Policy.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	SetAsPath(value BgpAsPath) BgpSrteV4Policy
	// HasAsPath checks if AsPath has been set in BgpSrteV4Policy
	HasAsPath() bool
	// Communities returns BgpSrteV4PolicyBgpCommunityIterIter, set in BgpSrteV4Policy
	Communities() BgpSrteV4PolicyBgpCommunityIter
	// ExtCommunities returns BgpSrteV4PolicyBgpExtCommunityIterIter, set in BgpSrteV4Policy
	ExtCommunities() BgpSrteV4PolicyBgpExtCommunityIter
	// TunnelTlvs returns BgpSrteV4PolicyBgpSrteV4TunnelTlvIterIter, set in BgpSrteV4Policy
	TunnelTlvs() BgpSrteV4PolicyBgpSrteV4TunnelTlvIter
	// Name returns string, set in BgpSrteV4Policy.
	Name() string
	// SetName assigns string provided by user to BgpSrteV4Policy
	SetName(value string) BgpSrteV4Policy
	// Active returns bool, set in BgpSrteV4Policy.
	Active() bool
	// SetActive assigns bool provided by user to BgpSrteV4Policy
	SetActive(value bool) BgpSrteV4Policy
	// HasActive checks if Active has been set in BgpSrteV4Policy
	HasActive() bool
	setNil()
}

// 4-octet value uniquely identifying the policy in the context of (color, endpoint) tuple. It is used by the SR Policy originator to make unique (from an NLRI perspective)  both for multiple candidate  paths of the same SR Policy as well as candidate paths  of different SR Policies (i.e. with different segment list) with the same Color  and Endpoint but meant for different head-ends.
// Distinguisher returns a uint32
func (obj *bgpSrteV4Policy) Distinguisher() uint32 {

	return *obj.obj.Distinguisher

}

// 4-octet value uniquely identifying the policy in the context of (color, endpoint) tuple. It is used by the SR Policy originator to make unique (from an NLRI perspective)  both for multiple candidate  paths of the same SR Policy as well as candidate paths  of different SR Policies (i.e. with different segment list) with the same Color  and Endpoint but meant for different head-ends.
// Distinguisher returns a uint32
func (obj *bgpSrteV4Policy) HasDistinguisher() bool {
	return obj.obj.Distinguisher != nil
}

// 4-octet value uniquely identifying the policy in the context of (color, endpoint) tuple. It is used by the SR Policy originator to make unique (from an NLRI perspective)  both for multiple candidate  paths of the same SR Policy as well as candidate paths  of different SR Policies (i.e. with different segment list) with the same Color  and Endpoint but meant for different head-ends.
// SetDistinguisher sets the uint32 value in the BgpSrteV4Policy object
func (obj *bgpSrteV4Policy) SetDistinguisher(value uint32) BgpSrteV4Policy {

	obj.obj.Distinguisher = &value
	return obj
}

// Policy color is used to match the color of the destination prefixes to steer traffic into the SR Policy.
// Color returns a uint32
func (obj *bgpSrteV4Policy) Color() uint32 {

	return *obj.obj.Color

}

// Policy color is used to match the color of the destination prefixes to steer traffic into the SR Policy.
// Color returns a uint32
func (obj *bgpSrteV4Policy) HasColor() bool {
	return obj.obj.Color != nil
}

// Policy color is used to match the color of the destination prefixes to steer traffic into the SR Policy.
// SetColor sets the uint32 value in the BgpSrteV4Policy object
func (obj *bgpSrteV4Policy) SetColor(value uint32) BgpSrteV4Policy {

	obj.obj.Color = &value
	return obj
}

// Specifies a single node or a set of nodes (e.g. an anycast address). It is selected on the basis of the SR Policy type (AFI).
// Ipv4Endpoint returns a string
func (obj *bgpSrteV4Policy) Ipv4Endpoint() string {

	return *obj.obj.Ipv4Endpoint

}

// Specifies a single node or a set of nodes (e.g. an anycast address). It is selected on the basis of the SR Policy type (AFI).
// SetIpv4Endpoint sets the string value in the BgpSrteV4Policy object
func (obj *bgpSrteV4Policy) SetIpv4Endpoint(value string) BgpSrteV4Policy {

	obj.obj.Ipv4Endpoint = &value
	return obj
}

type BgpSrteV4PolicyNextHopModeEnum string

// Enum of NextHopMode on BgpSrteV4Policy
var BgpSrteV4PolicyNextHopMode = struct {
	LOCAL_IP BgpSrteV4PolicyNextHopModeEnum
	MANUAL   BgpSrteV4PolicyNextHopModeEnum
}{
	LOCAL_IP: BgpSrteV4PolicyNextHopModeEnum("local_ip"),
	MANUAL:   BgpSrteV4PolicyNextHopModeEnum("manual"),
}

func (obj *bgpSrteV4Policy) NextHopMode() BgpSrteV4PolicyNextHopModeEnum {
	return BgpSrteV4PolicyNextHopModeEnum(obj.obj.NextHopMode.Enum().String())
}

// Mode for choosing the NextHop in MP REACH NLRI. Available modes are : Local IP: Automatically fills the Nexthop with the Local IP of the BGP peer. For IPv6 BGP peer the Nexthop Encoding capability should be enabled. Manual: Override the Nexthop with any arbitrary IPv4/IPv6 address.
// NextHopMode returns a string
func (obj *bgpSrteV4Policy) HasNextHopMode() bool {
	return obj.obj.NextHopMode != nil
}

func (obj *bgpSrteV4Policy) SetNextHopMode(value BgpSrteV4PolicyNextHopModeEnum) BgpSrteV4Policy {
	intValue, ok := otg.BgpSrteV4Policy_NextHopMode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpSrteV4PolicyNextHopModeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpSrteV4Policy_NextHopMode_Enum(intValue)
	obj.obj.NextHopMode = &enumValue

	return obj
}

type BgpSrteV4PolicyNextHopAddressTypeEnum string

// Enum of NextHopAddressType on BgpSrteV4Policy
var BgpSrteV4PolicyNextHopAddressType = struct {
	IPV4 BgpSrteV4PolicyNextHopAddressTypeEnum
	IPV6 BgpSrteV4PolicyNextHopAddressTypeEnum
}{
	IPV4: BgpSrteV4PolicyNextHopAddressTypeEnum("ipv4"),
	IPV6: BgpSrteV4PolicyNextHopAddressTypeEnum("ipv6"),
}

func (obj *bgpSrteV4Policy) NextHopAddressType() BgpSrteV4PolicyNextHopAddressTypeEnum {
	return BgpSrteV4PolicyNextHopAddressTypeEnum(obj.obj.NextHopAddressType.Enum().String())
}

// Type of next hop IP address to be used when 'next_hop_mode' is set to 'manual'.
// NextHopAddressType returns a string
func (obj *bgpSrteV4Policy) HasNextHopAddressType() bool {
	return obj.obj.NextHopAddressType != nil
}

func (obj *bgpSrteV4Policy) SetNextHopAddressType(value BgpSrteV4PolicyNextHopAddressTypeEnum) BgpSrteV4Policy {
	intValue, ok := otg.BgpSrteV4Policy_NextHopAddressType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpSrteV4PolicyNextHopAddressTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpSrteV4Policy_NextHopAddressType_Enum(intValue)
	obj.obj.NextHopAddressType = &enumValue

	return obj
}

// The IPv4 address of the next hop if the Nexthop type 'next_hop_mode' is 'manual' and the Nexthop type 'next_hop_address_type' is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability extended_next_hop_encoding should be enabled.
// NextHopIpv4Address returns a string
func (obj *bgpSrteV4Policy) NextHopIpv4Address() string {

	return *obj.obj.NextHopIpv4Address

}

// The IPv4 address of the next hop if the Nexthop type 'next_hop_mode' is 'manual' and the Nexthop type 'next_hop_address_type' is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability extended_next_hop_encoding should be enabled.
// NextHopIpv4Address returns a string
func (obj *bgpSrteV4Policy) HasNextHopIpv4Address() bool {
	return obj.obj.NextHopIpv4Address != nil
}

// The IPv4 address of the next hop if the Nexthop type 'next_hop_mode' is 'manual' and the Nexthop type 'next_hop_address_type' is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability extended_next_hop_encoding should be enabled.
// SetNextHopIpv4Address sets the string value in the BgpSrteV4Policy object
func (obj *bgpSrteV4Policy) SetNextHopIpv4Address(value string) BgpSrteV4Policy {

	obj.obj.NextHopIpv4Address = &value
	return obj
}

// The IPv6 address of the next hop if the Nexthop Mode 'next_hop_address_type' is 'manual' and the Nexthop type 'next_hop_address_type' is IPv6.
// NextHopIpv6Address returns a string
func (obj *bgpSrteV4Policy) NextHopIpv6Address() string {

	return *obj.obj.NextHopIpv6Address

}

// The IPv6 address of the next hop if the Nexthop Mode 'next_hop_address_type' is 'manual' and the Nexthop type 'next_hop_address_type' is IPv6.
// NextHopIpv6Address returns a string
func (obj *bgpSrteV4Policy) HasNextHopIpv6Address() bool {
	return obj.obj.NextHopIpv6Address != nil
}

// The IPv6 address of the next hop if the Nexthop Mode 'next_hop_address_type' is 'manual' and the Nexthop type 'next_hop_address_type' is IPv6.
// SetNextHopIpv6Address sets the string value in the BgpSrteV4Policy object
func (obj *bgpSrteV4Policy) SetNextHopIpv6Address(value string) BgpSrteV4Policy {

	obj.obj.NextHopIpv6Address = &value
	return obj
}

// description is TBD
// Advanced returns a BgpRouteAdvanced
func (obj *bgpSrteV4Policy) Advanced() BgpRouteAdvanced {
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
func (obj *bgpSrteV4Policy) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// description is TBD
// SetAdvanced sets the BgpRouteAdvanced value in the BgpSrteV4Policy object
func (obj *bgpSrteV4Policy) SetAdvanced(value BgpRouteAdvanced) BgpSrteV4Policy {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// description is TBD
// AddPath returns a BgpAddPath
func (obj *bgpSrteV4Policy) AddPath() BgpAddPath {
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
func (obj *bgpSrteV4Policy) HasAddPath() bool {
	return obj.obj.AddPath != nil
}

// description is TBD
// SetAddPath sets the BgpAddPath value in the BgpSrteV4Policy object
func (obj *bgpSrteV4Policy) SetAddPath(value BgpAddPath) BgpSrteV4Policy {

	obj.addPathHolder = nil
	obj.obj.AddPath = value.msg()

	return obj
}

// description is TBD
// AsPath returns a BgpAsPath
func (obj *bgpSrteV4Policy) AsPath() BgpAsPath {
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
func (obj *bgpSrteV4Policy) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the BgpAsPath value in the BgpSrteV4Policy object
func (obj *bgpSrteV4Policy) SetAsPath(value BgpAsPath) BgpSrteV4Policy {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// Optional Community settings.
// Communities returns a []BgpCommunity
func (obj *bgpSrteV4Policy) Communities() BgpSrteV4PolicyBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.BgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpSrteV4PolicyBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpSrteV4PolicyBgpCommunityIter struct {
	obj               *bgpSrteV4Policy
	bgpCommunitySlice []BgpCommunity
	fieldPtr          *[]*otg.BgpCommunity
}

func newBgpSrteV4PolicyBgpCommunityIter(ptr *[]*otg.BgpCommunity) BgpSrteV4PolicyBgpCommunityIter {
	return &bgpSrteV4PolicyBgpCommunityIter{fieldPtr: ptr}
}

type BgpSrteV4PolicyBgpCommunityIter interface {
	setMsg(*bgpSrteV4Policy) BgpSrteV4PolicyBgpCommunityIter
	Items() []BgpCommunity
	Add() BgpCommunity
	Append(items ...BgpCommunity) BgpSrteV4PolicyBgpCommunityIter
	Set(index int, newObj BgpCommunity) BgpSrteV4PolicyBgpCommunityIter
	Clear() BgpSrteV4PolicyBgpCommunityIter
	clearHolderSlice() BgpSrteV4PolicyBgpCommunityIter
	appendHolderSlice(item BgpCommunity) BgpSrteV4PolicyBgpCommunityIter
}

func (obj *bgpSrteV4PolicyBgpCommunityIter) setMsg(msg *bgpSrteV4Policy) BgpSrteV4PolicyBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpSrteV4PolicyBgpCommunityIter) Items() []BgpCommunity {
	return obj.bgpCommunitySlice
}

func (obj *bgpSrteV4PolicyBgpCommunityIter) Add() BgpCommunity {
	newObj := &otg.BgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpSrteV4PolicyBgpCommunityIter) Append(items ...BgpCommunity) BgpSrteV4PolicyBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpSrteV4PolicyBgpCommunityIter) Set(index int, newObj BgpCommunity) BgpSrteV4PolicyBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpSrteV4PolicyBgpCommunityIter) Clear() BgpSrteV4PolicyBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpCommunity{}
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpSrteV4PolicyBgpCommunityIter) clearHolderSlice() BgpSrteV4PolicyBgpCommunityIter {
	if len(obj.bgpCommunitySlice) > 0 {
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpSrteV4PolicyBgpCommunityIter) appendHolderSlice(item BgpCommunity) BgpSrteV4PolicyBgpCommunityIter {
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	return obj
}

// Optional Extended Community settings. The Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an 8-Bytes value. It is divided into two main parts. The first 2 Bytes of the community encode a type and sub-type fields and the last 6 Bytes carry a unique set of data in a format defined by the type and sub-type field. Extended communities provide a larger  range for grouping or categorizing communities. When type is administrator_as_2octet or administrator_as_4octet, the valid sub types are route target and origin. The valid value for  administrator_as_2octet and administrator_as_4octet type is either two byte AS followed by four byte local administrator id or four byte AS followed by two  byte local administrator id.  When type is administrator_ipv4_address the valid sub types are route target and origin. The valid value for  administrator_ipv4_address is a four byte IPv4 address followed by a two byte local administrator id.  When type is opaque, valid sub types are color and encapsulation. When sub type is color, first two bytes of the value field contain flags and last four bytes  contains the value of the color. When sub type is encapsulation the first four bytes of value field are reserved and last two bytes carries the tunnel type from  IANA's "ETHER TYPES" registry e.g IPv4 (protocol type = 0x0800), IPv6 (protocol type = 0x86dd), and MPLS (protocol type = 0x8847). When type is administrator_as_2octet_link_bandwidth the valid sub type is extended_bandwidth. The first two bytes of the value field contains the AS number and the last four bytes contains the bandwidth in IEEE floating point format.  When type is evpn the valid subtype is mac_address. In the value field the low-order bit of the first byte(Flags) is defined as the "Sticky/static" flag and may be set to 1, indicating the MAC address is static and cannot move. The second byte is reserved and the  last four bytes contain the sequence number which is used to ensure that PEs retain the correct MAC/IP Advertisement route when multiple updates  occur for the same MAC address.
// ExtCommunities returns a []BgpExtCommunity
func (obj *bgpSrteV4Policy) ExtCommunities() BgpSrteV4PolicyBgpExtCommunityIter {
	if len(obj.obj.ExtCommunities) == 0 {
		obj.obj.ExtCommunities = []*otg.BgpExtCommunity{}
	}
	if obj.extCommunitiesHolder == nil {
		obj.extCommunitiesHolder = newBgpSrteV4PolicyBgpExtCommunityIter(&obj.obj.ExtCommunities).setMsg(obj)
	}
	return obj.extCommunitiesHolder
}

type bgpSrteV4PolicyBgpExtCommunityIter struct {
	obj                  *bgpSrteV4Policy
	bgpExtCommunitySlice []BgpExtCommunity
	fieldPtr             *[]*otg.BgpExtCommunity
}

func newBgpSrteV4PolicyBgpExtCommunityIter(ptr *[]*otg.BgpExtCommunity) BgpSrteV4PolicyBgpExtCommunityIter {
	return &bgpSrteV4PolicyBgpExtCommunityIter{fieldPtr: ptr}
}

type BgpSrteV4PolicyBgpExtCommunityIter interface {
	setMsg(*bgpSrteV4Policy) BgpSrteV4PolicyBgpExtCommunityIter
	Items() []BgpExtCommunity
	Add() BgpExtCommunity
	Append(items ...BgpExtCommunity) BgpSrteV4PolicyBgpExtCommunityIter
	Set(index int, newObj BgpExtCommunity) BgpSrteV4PolicyBgpExtCommunityIter
	Clear() BgpSrteV4PolicyBgpExtCommunityIter
	clearHolderSlice() BgpSrteV4PolicyBgpExtCommunityIter
	appendHolderSlice(item BgpExtCommunity) BgpSrteV4PolicyBgpExtCommunityIter
}

func (obj *bgpSrteV4PolicyBgpExtCommunityIter) setMsg(msg *bgpSrteV4Policy) BgpSrteV4PolicyBgpExtCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpSrteV4PolicyBgpExtCommunityIter) Items() []BgpExtCommunity {
	return obj.bgpExtCommunitySlice
}

func (obj *bgpSrteV4PolicyBgpExtCommunityIter) Add() BgpExtCommunity {
	newObj := &otg.BgpExtCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpSrteV4PolicyBgpExtCommunityIter) Append(items ...BgpExtCommunity) BgpSrteV4PolicyBgpExtCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	}
	return obj
}

func (obj *bgpSrteV4PolicyBgpExtCommunityIter) Set(index int, newObj BgpExtCommunity) BgpSrteV4PolicyBgpExtCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpSrteV4PolicyBgpExtCommunityIter) Clear() BgpSrteV4PolicyBgpExtCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtCommunity{}
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpSrteV4PolicyBgpExtCommunityIter) clearHolderSlice() BgpSrteV4PolicyBgpExtCommunityIter {
	if len(obj.bgpExtCommunitySlice) > 0 {
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpSrteV4PolicyBgpExtCommunityIter) appendHolderSlice(item BgpExtCommunity) BgpSrteV4PolicyBgpExtCommunityIter {
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	return obj
}

// List Tunnel Encapsulation Attributes.
// TunnelTlvs returns a []BgpSrteV4TunnelTlv
func (obj *bgpSrteV4Policy) TunnelTlvs() BgpSrteV4PolicyBgpSrteV4TunnelTlvIter {
	if len(obj.obj.TunnelTlvs) == 0 {
		obj.obj.TunnelTlvs = []*otg.BgpSrteV4TunnelTlv{}
	}
	if obj.tunnelTlvsHolder == nil {
		obj.tunnelTlvsHolder = newBgpSrteV4PolicyBgpSrteV4TunnelTlvIter(&obj.obj.TunnelTlvs).setMsg(obj)
	}
	return obj.tunnelTlvsHolder
}

type bgpSrteV4PolicyBgpSrteV4TunnelTlvIter struct {
	obj                     *bgpSrteV4Policy
	bgpSrteV4TunnelTlvSlice []BgpSrteV4TunnelTlv
	fieldPtr                *[]*otg.BgpSrteV4TunnelTlv
}

func newBgpSrteV4PolicyBgpSrteV4TunnelTlvIter(ptr *[]*otg.BgpSrteV4TunnelTlv) BgpSrteV4PolicyBgpSrteV4TunnelTlvIter {
	return &bgpSrteV4PolicyBgpSrteV4TunnelTlvIter{fieldPtr: ptr}
}

type BgpSrteV4PolicyBgpSrteV4TunnelTlvIter interface {
	setMsg(*bgpSrteV4Policy) BgpSrteV4PolicyBgpSrteV4TunnelTlvIter
	Items() []BgpSrteV4TunnelTlv
	Add() BgpSrteV4TunnelTlv
	Append(items ...BgpSrteV4TunnelTlv) BgpSrteV4PolicyBgpSrteV4TunnelTlvIter
	Set(index int, newObj BgpSrteV4TunnelTlv) BgpSrteV4PolicyBgpSrteV4TunnelTlvIter
	Clear() BgpSrteV4PolicyBgpSrteV4TunnelTlvIter
	clearHolderSlice() BgpSrteV4PolicyBgpSrteV4TunnelTlvIter
	appendHolderSlice(item BgpSrteV4TunnelTlv) BgpSrteV4PolicyBgpSrteV4TunnelTlvIter
}

func (obj *bgpSrteV4PolicyBgpSrteV4TunnelTlvIter) setMsg(msg *bgpSrteV4Policy) BgpSrteV4PolicyBgpSrteV4TunnelTlvIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpSrteV4TunnelTlv{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpSrteV4PolicyBgpSrteV4TunnelTlvIter) Items() []BgpSrteV4TunnelTlv {
	return obj.bgpSrteV4TunnelTlvSlice
}

func (obj *bgpSrteV4PolicyBgpSrteV4TunnelTlvIter) Add() BgpSrteV4TunnelTlv {
	newObj := &otg.BgpSrteV4TunnelTlv{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpSrteV4TunnelTlv{obj: newObj}
	newLibObj.setDefault()
	obj.bgpSrteV4TunnelTlvSlice = append(obj.bgpSrteV4TunnelTlvSlice, newLibObj)
	return newLibObj
}

func (obj *bgpSrteV4PolicyBgpSrteV4TunnelTlvIter) Append(items ...BgpSrteV4TunnelTlv) BgpSrteV4PolicyBgpSrteV4TunnelTlvIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpSrteV4TunnelTlvSlice = append(obj.bgpSrteV4TunnelTlvSlice, item)
	}
	return obj
}

func (obj *bgpSrteV4PolicyBgpSrteV4TunnelTlvIter) Set(index int, newObj BgpSrteV4TunnelTlv) BgpSrteV4PolicyBgpSrteV4TunnelTlvIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpSrteV4TunnelTlvSlice[index] = newObj
	return obj
}
func (obj *bgpSrteV4PolicyBgpSrteV4TunnelTlvIter) Clear() BgpSrteV4PolicyBgpSrteV4TunnelTlvIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpSrteV4TunnelTlv{}
		obj.bgpSrteV4TunnelTlvSlice = []BgpSrteV4TunnelTlv{}
	}
	return obj
}
func (obj *bgpSrteV4PolicyBgpSrteV4TunnelTlvIter) clearHolderSlice() BgpSrteV4PolicyBgpSrteV4TunnelTlvIter {
	if len(obj.bgpSrteV4TunnelTlvSlice) > 0 {
		obj.bgpSrteV4TunnelTlvSlice = []BgpSrteV4TunnelTlv{}
	}
	return obj
}
func (obj *bgpSrteV4PolicyBgpSrteV4TunnelTlvIter) appendHolderSlice(item BgpSrteV4TunnelTlv) BgpSrteV4PolicyBgpSrteV4TunnelTlvIter {
	obj.bgpSrteV4TunnelTlvSlice = append(obj.bgpSrteV4TunnelTlvSlice, item)
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpSrteV4Policy) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpSrteV4Policy object
func (obj *bgpSrteV4Policy) SetName(value string) BgpSrteV4Policy {

	obj.obj.Name = &value
	return obj
}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// Active returns a bool
func (obj *bgpSrteV4Policy) Active() bool {

	return *obj.obj.Active

}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// Active returns a bool
func (obj *bgpSrteV4Policy) HasActive() bool {
	return obj.obj.Active != nil
}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// SetActive sets the bool value in the BgpSrteV4Policy object
func (obj *bgpSrteV4Policy) SetActive(value bool) BgpSrteV4Policy {

	obj.obj.Active = &value
	return obj
}

func (obj *bgpSrteV4Policy) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Ipv4Endpoint is required
	if obj.obj.Ipv4Endpoint == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4Endpoint is required field on interface BgpSrteV4Policy")
	}
	if obj.obj.Ipv4Endpoint != nil {

		err := obj.validateIpv4(obj.Ipv4Endpoint())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteV4Policy.Ipv4Endpoint"))
		}

	}

	if obj.obj.NextHopIpv4Address != nil {

		err := obj.validateIpv4(obj.NextHopIpv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteV4Policy.NextHopIpv4Address"))
		}

	}

	if obj.obj.NextHopIpv6Address != nil {

		err := obj.validateIpv6(obj.NextHopIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteV4Policy.NextHopIpv6Address"))
		}

	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(vObj, set_default)
	}

	if obj.obj.AddPath != nil {

		obj.AddPath().validateObj(vObj, set_default)
	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(vObj, set_default)
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

	if len(obj.obj.TunnelTlvs) != 0 {

		if set_default {
			obj.TunnelTlvs().clearHolderSlice()
			for _, item := range obj.obj.TunnelTlvs {
				obj.TunnelTlvs().appendHolderSlice(&bgpSrteV4TunnelTlv{obj: item})
			}
		}
		for _, item := range obj.TunnelTlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpSrteV4Policy")
	}
}

func (obj *bgpSrteV4Policy) setDefault() {
	if obj.obj.Distinguisher == nil {
		obj.SetDistinguisher(1)
	}
	if obj.obj.Color == nil {
		obj.SetColor(100)
	}
	if obj.obj.NextHopMode == nil {
		obj.SetNextHopMode(BgpSrteV4PolicyNextHopMode.LOCAL_IP)

	}
	if obj.obj.NextHopAddressType == nil {
		obj.SetNextHopAddressType(BgpSrteV4PolicyNextHopAddressType.IPV4)

	}
	if obj.obj.Active == nil {
		obj.SetActive(true)
	}

}
