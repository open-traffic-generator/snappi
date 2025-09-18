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

// ***** BgpSrteV6Policy *****
type bgpSrteV6Policy struct {
	validation
	obj                  *otg.BgpSrteV6Policy
	marshaller           marshalBgpSrteV6Policy
	unMarshaller         unMarshalBgpSrteV6Policy
	advancedHolder       BgpRouteAdvanced
	addPathHolder        BgpAddPath
	asPathHolder         BgpAsPath
	communitiesHolder    BgpSrteV6PolicyBgpCommunityIter
	extcommunitiesHolder BgpSrteV6PolicyBgpExtCommunityIter
	tunnelTlvsHolder     BgpSrteV6PolicyBgpSrteV6TunnelTlvIter
}

func NewBgpSrteV6Policy() BgpSrteV6Policy {
	obj := bgpSrteV6Policy{obj: &otg.BgpSrteV6Policy{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteV6Policy) msg() *otg.BgpSrteV6Policy {
	return obj.obj
}

func (obj *bgpSrteV6Policy) setMsg(msg *otg.BgpSrteV6Policy) BgpSrteV6Policy {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteV6Policy struct {
	obj *bgpSrteV6Policy
}

type marshalBgpSrteV6Policy interface {
	// ToProto marshals BgpSrteV6Policy to protobuf object *otg.BgpSrteV6Policy
	ToProto() (*otg.BgpSrteV6Policy, error)
	// ToPbText marshals BgpSrteV6Policy to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteV6Policy to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteV6Policy to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrteV6Policy struct {
	obj *bgpSrteV6Policy
}

type unMarshalBgpSrteV6Policy interface {
	// FromProto unmarshals BgpSrteV6Policy from protobuf object *otg.BgpSrteV6Policy
	FromProto(msg *otg.BgpSrteV6Policy) (BgpSrteV6Policy, error)
	// FromPbText unmarshals BgpSrteV6Policy from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteV6Policy from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteV6Policy from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteV6Policy) Marshal() marshalBgpSrteV6Policy {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteV6Policy{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteV6Policy) Unmarshal() unMarshalBgpSrteV6Policy {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteV6Policy{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteV6Policy) ToProto() (*otg.BgpSrteV6Policy, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteV6Policy) FromProto(msg *otg.BgpSrteV6Policy) (BgpSrteV6Policy, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteV6Policy) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteV6Policy) FromPbText(value string) error {
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

func (m *marshalbgpSrteV6Policy) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteV6Policy) FromYaml(value string) error {
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

func (m *marshalbgpSrteV6Policy) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteV6Policy) FromJson(value string) error {
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

func (obj *bgpSrteV6Policy) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteV6Policy) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteV6Policy) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteV6Policy) Clone() (BgpSrteV6Policy, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteV6Policy()
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

func (obj *bgpSrteV6Policy) setNil() {
	obj.advancedHolder = nil
	obj.addPathHolder = nil
	obj.asPathHolder = nil
	obj.communitiesHolder = nil
	obj.extcommunitiesHolder = nil
	obj.tunnelTlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteV6Policy is configuration for BGP Segment Routing Traffic Engineering policy.

type BgpSrteV6Policy interface {
	Validation
	// msg marshals BgpSrteV6Policy to protobuf object *otg.BgpSrteV6Policy
	// and doesn't set defaults
	msg() *otg.BgpSrteV6Policy
	// setMsg unmarshals BgpSrteV6Policy from protobuf object *otg.BgpSrteV6Policy
	// and doesn't set defaults
	setMsg(*otg.BgpSrteV6Policy) BgpSrteV6Policy
	// provides marshal interface
	Marshal() marshalBgpSrteV6Policy
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteV6Policy
	// validate validates BgpSrteV6Policy
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteV6Policy, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Distinguisher returns uint32, set in BgpSrteV6Policy.
	Distinguisher() uint32
	// SetDistinguisher assigns uint32 provided by user to BgpSrteV6Policy
	SetDistinguisher(value uint32) BgpSrteV6Policy
	// HasDistinguisher checks if Distinguisher has been set in BgpSrteV6Policy
	HasDistinguisher() bool
	// Color returns uint32, set in BgpSrteV6Policy.
	Color() uint32
	// SetColor assigns uint32 provided by user to BgpSrteV6Policy
	SetColor(value uint32) BgpSrteV6Policy
	// HasColor checks if Color has been set in BgpSrteV6Policy
	HasColor() bool
	// Ipv6Endpoint returns string, set in BgpSrteV6Policy.
	Ipv6Endpoint() string
	// SetIpv6Endpoint assigns string provided by user to BgpSrteV6Policy
	SetIpv6Endpoint(value string) BgpSrteV6Policy
	// NextHopMode returns BgpSrteV6PolicyNextHopModeEnum, set in BgpSrteV6Policy
	NextHopMode() BgpSrteV6PolicyNextHopModeEnum
	// SetNextHopMode assigns BgpSrteV6PolicyNextHopModeEnum provided by user to BgpSrteV6Policy
	SetNextHopMode(value BgpSrteV6PolicyNextHopModeEnum) BgpSrteV6Policy
	// HasNextHopMode checks if NextHopMode has been set in BgpSrteV6Policy
	HasNextHopMode() bool
	// NextHopAddressType returns BgpSrteV6PolicyNextHopAddressTypeEnum, set in BgpSrteV6Policy
	NextHopAddressType() BgpSrteV6PolicyNextHopAddressTypeEnum
	// SetNextHopAddressType assigns BgpSrteV6PolicyNextHopAddressTypeEnum provided by user to BgpSrteV6Policy
	SetNextHopAddressType(value BgpSrteV6PolicyNextHopAddressTypeEnum) BgpSrteV6Policy
	// HasNextHopAddressType checks if NextHopAddressType has been set in BgpSrteV6Policy
	HasNextHopAddressType() bool
	// NextHopIpv4Address returns string, set in BgpSrteV6Policy.
	NextHopIpv4Address() string
	// SetNextHopIpv4Address assigns string provided by user to BgpSrteV6Policy
	SetNextHopIpv4Address(value string) BgpSrteV6Policy
	// HasNextHopIpv4Address checks if NextHopIpv4Address has been set in BgpSrteV6Policy
	HasNextHopIpv4Address() bool
	// NextHopIpv6Address returns string, set in BgpSrteV6Policy.
	NextHopIpv6Address() string
	// SetNextHopIpv6Address assigns string provided by user to BgpSrteV6Policy
	SetNextHopIpv6Address(value string) BgpSrteV6Policy
	// HasNextHopIpv6Address checks if NextHopIpv6Address has been set in BgpSrteV6Policy
	HasNextHopIpv6Address() bool
	// Advanced returns BgpRouteAdvanced, set in BgpSrteV6Policy.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	Advanced() BgpRouteAdvanced
	// SetAdvanced assigns BgpRouteAdvanced provided by user to BgpSrteV6Policy.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	SetAdvanced(value BgpRouteAdvanced) BgpSrteV6Policy
	// HasAdvanced checks if Advanced has been set in BgpSrteV6Policy
	HasAdvanced() bool
	// AddPath returns BgpAddPath, set in BgpSrteV6Policy.
	// BgpAddPath is the BGP Additional Paths feature is a BGP extension that allows the  advertisement of multiple paths for the same prefix without the new  paths implicitly replacing any previous paths.
	AddPath() BgpAddPath
	// SetAddPath assigns BgpAddPath provided by user to BgpSrteV6Policy.
	// BgpAddPath is the BGP Additional Paths feature is a BGP extension that allows the  advertisement of multiple paths for the same prefix without the new  paths implicitly replacing any previous paths.
	SetAddPath(value BgpAddPath) BgpSrteV6Policy
	// HasAddPath checks if AddPath has been set in BgpSrteV6Policy
	HasAddPath() bool
	// AsPath returns BgpAsPath, set in BgpSrteV6Policy.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	AsPath() BgpAsPath
	// SetAsPath assigns BgpAsPath provided by user to BgpSrteV6Policy.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	SetAsPath(value BgpAsPath) BgpSrteV6Policy
	// HasAsPath checks if AsPath has been set in BgpSrteV6Policy
	HasAsPath() bool
	// Communities returns BgpSrteV6PolicyBgpCommunityIterIter, set in BgpSrteV6Policy
	Communities() BgpSrteV6PolicyBgpCommunityIter
	// Extcommunities returns BgpSrteV6PolicyBgpExtCommunityIterIter, set in BgpSrteV6Policy
	Extcommunities() BgpSrteV6PolicyBgpExtCommunityIter
	// TunnelTlvs returns BgpSrteV6PolicyBgpSrteV6TunnelTlvIterIter, set in BgpSrteV6Policy
	TunnelTlvs() BgpSrteV6PolicyBgpSrteV6TunnelTlvIter
	// Name returns string, set in BgpSrteV6Policy.
	Name() string
	// SetName assigns string provided by user to BgpSrteV6Policy
	SetName(value string) BgpSrteV6Policy
	// Active returns bool, set in BgpSrteV6Policy.
	Active() bool
	// SetActive assigns bool provided by user to BgpSrteV6Policy
	SetActive(value bool) BgpSrteV6Policy
	// HasActive checks if Active has been set in BgpSrteV6Policy
	HasActive() bool
	setNil()
}

// Identifies the policy in the context of (color and endpoint) tuple.  It is used by the SR Policy originator to make unique multiple  occurrences of the same SR Policy.
// Distinguisher returns a uint32
func (obj *bgpSrteV6Policy) Distinguisher() uint32 {

	return *obj.obj.Distinguisher

}

// Identifies the policy in the context of (color and endpoint) tuple.  It is used by the SR Policy originator to make unique multiple  occurrences of the same SR Policy.
// Distinguisher returns a uint32
func (obj *bgpSrteV6Policy) HasDistinguisher() bool {
	return obj.obj.Distinguisher != nil
}

// Identifies the policy in the context of (color and endpoint) tuple.  It is used by the SR Policy originator to make unique multiple  occurrences of the same SR Policy.
// SetDistinguisher sets the uint32 value in the BgpSrteV6Policy object
func (obj *bgpSrteV6Policy) SetDistinguisher(value uint32) BgpSrteV6Policy {

	obj.obj.Distinguisher = &value
	return obj
}

// Identifies the policy. It is used to match the color of the  destination prefixes to steer traffic into the SR Policy.
// Color returns a uint32
func (obj *bgpSrteV6Policy) Color() uint32 {

	return *obj.obj.Color

}

// Identifies the policy. It is used to match the color of the  destination prefixes to steer traffic into the SR Policy.
// Color returns a uint32
func (obj *bgpSrteV6Policy) HasColor() bool {
	return obj.obj.Color != nil
}

// Identifies the policy. It is used to match the color of the  destination prefixes to steer traffic into the SR Policy.
// SetColor sets the uint32 value in the BgpSrteV6Policy object
func (obj *bgpSrteV6Policy) SetColor(value uint32) BgpSrteV6Policy {

	obj.obj.Color = &value
	return obj
}

// Specifies a single node or a set of nodes (e.g., an anycast address). It is selected on the basis of the SR Policy type (AFI).
// Ipv6Endpoint returns a string
func (obj *bgpSrteV6Policy) Ipv6Endpoint() string {

	return *obj.obj.Ipv6Endpoint

}

// Specifies a single node or a set of nodes (e.g., an anycast address). It is selected on the basis of the SR Policy type (AFI).
// SetIpv6Endpoint sets the string value in the BgpSrteV6Policy object
func (obj *bgpSrteV6Policy) SetIpv6Endpoint(value string) BgpSrteV6Policy {

	obj.obj.Ipv6Endpoint = &value
	return obj
}

type BgpSrteV6PolicyNextHopModeEnum string

// Enum of NextHopMode on BgpSrteV6Policy
var BgpSrteV6PolicyNextHopMode = struct {
	LOCAL_IP BgpSrteV6PolicyNextHopModeEnum
	MANUAL   BgpSrteV6PolicyNextHopModeEnum
}{
	LOCAL_IP: BgpSrteV6PolicyNextHopModeEnum("local_ip"),
	MANUAL:   BgpSrteV6PolicyNextHopModeEnum("manual"),
}

func (obj *bgpSrteV6Policy) NextHopMode() BgpSrteV6PolicyNextHopModeEnum {
	return BgpSrteV6PolicyNextHopModeEnum(obj.obj.NextHopMode.Enum().String())
}

// Mode for choosing the NextHop in MP REACH NLRI. Available modes are : Local IP: Automatically fills the Nexthop with the Local IP of the BGP peer. For IPv6 BGP peer the Nexthop Encoding capability should be enabled. Manual: Override the Nexthop with any arbitrary IPv4/IPv6 address.
// NextHopMode returns a string
func (obj *bgpSrteV6Policy) HasNextHopMode() bool {
	return obj.obj.NextHopMode != nil
}

func (obj *bgpSrteV6Policy) SetNextHopMode(value BgpSrteV6PolicyNextHopModeEnum) BgpSrteV6Policy {
	intValue, ok := otg.BgpSrteV6Policy_NextHopMode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpSrteV6PolicyNextHopModeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpSrteV6Policy_NextHopMode_Enum(intValue)
	obj.obj.NextHopMode = &enumValue

	return obj
}

type BgpSrteV6PolicyNextHopAddressTypeEnum string

// Enum of NextHopAddressType on BgpSrteV6Policy
var BgpSrteV6PolicyNextHopAddressType = struct {
	IPV4 BgpSrteV6PolicyNextHopAddressTypeEnum
	IPV6 BgpSrteV6PolicyNextHopAddressTypeEnum
}{
	IPV4: BgpSrteV6PolicyNextHopAddressTypeEnum("ipv4"),
	IPV6: BgpSrteV6PolicyNextHopAddressTypeEnum("ipv6"),
}

func (obj *bgpSrteV6Policy) NextHopAddressType() BgpSrteV6PolicyNextHopAddressTypeEnum {
	return BgpSrteV6PolicyNextHopAddressTypeEnum(obj.obj.NextHopAddressType.Enum().String())
}

// Type of next hop IP address to be used when 'next_hop_mode' is set to 'manual'.
// NextHopAddressType returns a string
func (obj *bgpSrteV6Policy) HasNextHopAddressType() bool {
	return obj.obj.NextHopAddressType != nil
}

func (obj *bgpSrteV6Policy) SetNextHopAddressType(value BgpSrteV6PolicyNextHopAddressTypeEnum) BgpSrteV6Policy {
	intValue, ok := otg.BgpSrteV6Policy_NextHopAddressType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpSrteV6PolicyNextHopAddressTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpSrteV6Policy_NextHopAddressType_Enum(intValue)
	obj.obj.NextHopAddressType = &enumValue

	return obj
}

// The IPv4 address of the Nexthop if the 'next_hop_mode' is 'manual' and the Nexthop type 'next_hop_address_type' is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability extended_next_hop_encoding should be enabled.
// NextHopIpv4Address returns a string
func (obj *bgpSrteV6Policy) NextHopIpv4Address() string {

	return *obj.obj.NextHopIpv4Address

}

// The IPv4 address of the Nexthop if the 'next_hop_mode' is 'manual' and the Nexthop type 'next_hop_address_type' is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability extended_next_hop_encoding should be enabled.
// NextHopIpv4Address returns a string
func (obj *bgpSrteV6Policy) HasNextHopIpv4Address() bool {
	return obj.obj.NextHopIpv4Address != nil
}

// The IPv4 address of the Nexthop if the 'next_hop_mode' is 'manual' and the Nexthop type 'next_hop_address_type' is IPv4. If BGP peer is of type IPv6, Nexthop Encoding capability extended_next_hop_encoding should be enabled.
// SetNextHopIpv4Address sets the string value in the BgpSrteV6Policy object
func (obj *bgpSrteV6Policy) SetNextHopIpv4Address(value string) BgpSrteV6Policy {

	obj.obj.NextHopIpv4Address = &value
	return obj
}

// The IPv6 address of the next hop if the Nexthop Mode 'next_hop_address_type' is 'manual' and the Nexthop type 'next_hop_address_type' is IPv6.
// NextHopIpv6Address returns a string
func (obj *bgpSrteV6Policy) NextHopIpv6Address() string {

	return *obj.obj.NextHopIpv6Address

}

// The IPv6 address of the next hop if the Nexthop Mode 'next_hop_address_type' is 'manual' and the Nexthop type 'next_hop_address_type' is IPv6.
// NextHopIpv6Address returns a string
func (obj *bgpSrteV6Policy) HasNextHopIpv6Address() bool {
	return obj.obj.NextHopIpv6Address != nil
}

// The IPv6 address of the next hop if the Nexthop Mode 'next_hop_address_type' is 'manual' and the Nexthop type 'next_hop_address_type' is IPv6.
// SetNextHopIpv6Address sets the string value in the BgpSrteV6Policy object
func (obj *bgpSrteV6Policy) SetNextHopIpv6Address(value string) BgpSrteV6Policy {

	obj.obj.NextHopIpv6Address = &value
	return obj
}

// description is TBD
// Advanced returns a BgpRouteAdvanced
func (obj *bgpSrteV6Policy) Advanced() BgpRouteAdvanced {
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
func (obj *bgpSrteV6Policy) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// description is TBD
// SetAdvanced sets the BgpRouteAdvanced value in the BgpSrteV6Policy object
func (obj *bgpSrteV6Policy) SetAdvanced(value BgpRouteAdvanced) BgpSrteV6Policy {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// description is TBD
// AddPath returns a BgpAddPath
func (obj *bgpSrteV6Policy) AddPath() BgpAddPath {
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
func (obj *bgpSrteV6Policy) HasAddPath() bool {
	return obj.obj.AddPath != nil
}

// description is TBD
// SetAddPath sets the BgpAddPath value in the BgpSrteV6Policy object
func (obj *bgpSrteV6Policy) SetAddPath(value BgpAddPath) BgpSrteV6Policy {

	obj.addPathHolder = nil
	obj.obj.AddPath = value.msg()

	return obj
}

// description is TBD
// AsPath returns a BgpAsPath
func (obj *bgpSrteV6Policy) AsPath() BgpAsPath {
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
func (obj *bgpSrteV6Policy) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the BgpAsPath value in the BgpSrteV6Policy object
func (obj *bgpSrteV6Policy) SetAsPath(value BgpAsPath) BgpSrteV6Policy {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// Optional community settings.
// Communities returns a []BgpCommunity
func (obj *bgpSrteV6Policy) Communities() BgpSrteV6PolicyBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.BgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpSrteV6PolicyBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpSrteV6PolicyBgpCommunityIter struct {
	obj               *bgpSrteV6Policy
	bgpCommunitySlice []BgpCommunity
	fieldPtr          *[]*otg.BgpCommunity
}

func newBgpSrteV6PolicyBgpCommunityIter(ptr *[]*otg.BgpCommunity) BgpSrteV6PolicyBgpCommunityIter {
	return &bgpSrteV6PolicyBgpCommunityIter{fieldPtr: ptr}
}

type BgpSrteV6PolicyBgpCommunityIter interface {
	setMsg(*bgpSrteV6Policy) BgpSrteV6PolicyBgpCommunityIter
	Items() []BgpCommunity
	Add() BgpCommunity
	Append(items ...BgpCommunity) BgpSrteV6PolicyBgpCommunityIter
	Set(index int, newObj BgpCommunity) BgpSrteV6PolicyBgpCommunityIter
	Clear() BgpSrteV6PolicyBgpCommunityIter
	clearHolderSlice() BgpSrteV6PolicyBgpCommunityIter
	appendHolderSlice(item BgpCommunity) BgpSrteV6PolicyBgpCommunityIter
}

func (obj *bgpSrteV6PolicyBgpCommunityIter) setMsg(msg *bgpSrteV6Policy) BgpSrteV6PolicyBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpSrteV6PolicyBgpCommunityIter) Items() []BgpCommunity {
	return obj.bgpCommunitySlice
}

func (obj *bgpSrteV6PolicyBgpCommunityIter) Add() BgpCommunity {
	newObj := &otg.BgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpSrteV6PolicyBgpCommunityIter) Append(items ...BgpCommunity) BgpSrteV6PolicyBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpSrteV6PolicyBgpCommunityIter) Set(index int, newObj BgpCommunity) BgpSrteV6PolicyBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpSrteV6PolicyBgpCommunityIter) Clear() BgpSrteV6PolicyBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpCommunity{}
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpSrteV6PolicyBgpCommunityIter) clearHolderSlice() BgpSrteV6PolicyBgpCommunityIter {
	if len(obj.bgpCommunitySlice) > 0 {
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpSrteV6PolicyBgpCommunityIter) appendHolderSlice(item BgpCommunity) BgpSrteV6PolicyBgpCommunityIter {
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	return obj
}

// Optional Extended Community settings. The Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an 8-Bytes value. It is divided into two main parts. The first 2 Bytes of the community encode a type and sub-type fields and the last 6 Bytes carry a unique set of data in a format defined by the type and sub-type field. Extended communities provide a larger  range for grouping or categorizing communities. When type is administrator_as_2octet or administrator_as_4octet, the valid sub types are route target and origin. The valid value for  administrator_as_2octet and administrator_as_4octet type is either two byte AS followed by four byte local administrator id or four byte AS followed by two  byte local administrator id.  When type is administrator_ipv4_address the valid sub types are route target and origin. The valid value for  administrator_ipv4_address is a four byte IPv4 address followed by a two byte local administrator id.  When type is opaque, valid sub types are color and encapsulation. When sub type is color, first two bytes of the value field contain flags and last four bytes  contains the value of the color. When sub type is encapsulation the first four bytes of value field are reserved and last two bytes carries the tunnel type from  IANA's "ETHER TYPES" registry e.g IPv4 (protocol type = 0x0800), IPv6 (protocol type = 0x86dd), and MPLS (protocol type = 0x8847). When type is administrator_as_2octet_link_bandwidth the valid sub type is extended_bandwidth. The first two bytes of the value field contains the AS number and the last four bytes contains the bandwidth in IEEE floating point format.  When type is evpn the valid subtype is mac_address. In the value field the low-order bit of the first byte(Flags) is defined as the "Sticky/static" flag and may be set to 1, indicating the MAC address is static and cannot move. The second byte is reserved and the  last four bytes contain the sequence number which is used to ensure that PEs retain the correct MAC/IP Advertisement route when multiple updates  occur for the same MAC address.
// Extcommunities returns a []BgpExtCommunity
func (obj *bgpSrteV6Policy) Extcommunities() BgpSrteV6PolicyBgpExtCommunityIter {
	if len(obj.obj.Extcommunities) == 0 {
		obj.obj.Extcommunities = []*otg.BgpExtCommunity{}
	}
	if obj.extcommunitiesHolder == nil {
		obj.extcommunitiesHolder = newBgpSrteV6PolicyBgpExtCommunityIter(&obj.obj.Extcommunities).setMsg(obj)
	}
	return obj.extcommunitiesHolder
}

type bgpSrteV6PolicyBgpExtCommunityIter struct {
	obj                  *bgpSrteV6Policy
	bgpExtCommunitySlice []BgpExtCommunity
	fieldPtr             *[]*otg.BgpExtCommunity
}

func newBgpSrteV6PolicyBgpExtCommunityIter(ptr *[]*otg.BgpExtCommunity) BgpSrteV6PolicyBgpExtCommunityIter {
	return &bgpSrteV6PolicyBgpExtCommunityIter{fieldPtr: ptr}
}

type BgpSrteV6PolicyBgpExtCommunityIter interface {
	setMsg(*bgpSrteV6Policy) BgpSrteV6PolicyBgpExtCommunityIter
	Items() []BgpExtCommunity
	Add() BgpExtCommunity
	Append(items ...BgpExtCommunity) BgpSrteV6PolicyBgpExtCommunityIter
	Set(index int, newObj BgpExtCommunity) BgpSrteV6PolicyBgpExtCommunityIter
	Clear() BgpSrteV6PolicyBgpExtCommunityIter
	clearHolderSlice() BgpSrteV6PolicyBgpExtCommunityIter
	appendHolderSlice(item BgpExtCommunity) BgpSrteV6PolicyBgpExtCommunityIter
}

func (obj *bgpSrteV6PolicyBgpExtCommunityIter) setMsg(msg *bgpSrteV6Policy) BgpSrteV6PolicyBgpExtCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpSrteV6PolicyBgpExtCommunityIter) Items() []BgpExtCommunity {
	return obj.bgpExtCommunitySlice
}

func (obj *bgpSrteV6PolicyBgpExtCommunityIter) Add() BgpExtCommunity {
	newObj := &otg.BgpExtCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpSrteV6PolicyBgpExtCommunityIter) Append(items ...BgpExtCommunity) BgpSrteV6PolicyBgpExtCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	}
	return obj
}

func (obj *bgpSrteV6PolicyBgpExtCommunityIter) Set(index int, newObj BgpExtCommunity) BgpSrteV6PolicyBgpExtCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpSrteV6PolicyBgpExtCommunityIter) Clear() BgpSrteV6PolicyBgpExtCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtCommunity{}
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpSrteV6PolicyBgpExtCommunityIter) clearHolderSlice() BgpSrteV6PolicyBgpExtCommunityIter {
	if len(obj.bgpExtCommunitySlice) > 0 {
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpSrteV6PolicyBgpExtCommunityIter) appendHolderSlice(item BgpExtCommunity) BgpSrteV6PolicyBgpExtCommunityIter {
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	return obj
}

// List of optional tunnel TLV settings.
// TunnelTlvs returns a []BgpSrteV6TunnelTlv
func (obj *bgpSrteV6Policy) TunnelTlvs() BgpSrteV6PolicyBgpSrteV6TunnelTlvIter {
	if len(obj.obj.TunnelTlvs) == 0 {
		obj.obj.TunnelTlvs = []*otg.BgpSrteV6TunnelTlv{}
	}
	if obj.tunnelTlvsHolder == nil {
		obj.tunnelTlvsHolder = newBgpSrteV6PolicyBgpSrteV6TunnelTlvIter(&obj.obj.TunnelTlvs).setMsg(obj)
	}
	return obj.tunnelTlvsHolder
}

type bgpSrteV6PolicyBgpSrteV6TunnelTlvIter struct {
	obj                     *bgpSrteV6Policy
	bgpSrteV6TunnelTlvSlice []BgpSrteV6TunnelTlv
	fieldPtr                *[]*otg.BgpSrteV6TunnelTlv
}

func newBgpSrteV6PolicyBgpSrteV6TunnelTlvIter(ptr *[]*otg.BgpSrteV6TunnelTlv) BgpSrteV6PolicyBgpSrteV6TunnelTlvIter {
	return &bgpSrteV6PolicyBgpSrteV6TunnelTlvIter{fieldPtr: ptr}
}

type BgpSrteV6PolicyBgpSrteV6TunnelTlvIter interface {
	setMsg(*bgpSrteV6Policy) BgpSrteV6PolicyBgpSrteV6TunnelTlvIter
	Items() []BgpSrteV6TunnelTlv
	Add() BgpSrteV6TunnelTlv
	Append(items ...BgpSrteV6TunnelTlv) BgpSrteV6PolicyBgpSrteV6TunnelTlvIter
	Set(index int, newObj BgpSrteV6TunnelTlv) BgpSrteV6PolicyBgpSrteV6TunnelTlvIter
	Clear() BgpSrteV6PolicyBgpSrteV6TunnelTlvIter
	clearHolderSlice() BgpSrteV6PolicyBgpSrteV6TunnelTlvIter
	appendHolderSlice(item BgpSrteV6TunnelTlv) BgpSrteV6PolicyBgpSrteV6TunnelTlvIter
}

func (obj *bgpSrteV6PolicyBgpSrteV6TunnelTlvIter) setMsg(msg *bgpSrteV6Policy) BgpSrteV6PolicyBgpSrteV6TunnelTlvIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpSrteV6TunnelTlv{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpSrteV6PolicyBgpSrteV6TunnelTlvIter) Items() []BgpSrteV6TunnelTlv {
	return obj.bgpSrteV6TunnelTlvSlice
}

func (obj *bgpSrteV6PolicyBgpSrteV6TunnelTlvIter) Add() BgpSrteV6TunnelTlv {
	newObj := &otg.BgpSrteV6TunnelTlv{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpSrteV6TunnelTlv{obj: newObj}
	newLibObj.setDefault()
	obj.bgpSrteV6TunnelTlvSlice = append(obj.bgpSrteV6TunnelTlvSlice, newLibObj)
	return newLibObj
}

func (obj *bgpSrteV6PolicyBgpSrteV6TunnelTlvIter) Append(items ...BgpSrteV6TunnelTlv) BgpSrteV6PolicyBgpSrteV6TunnelTlvIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpSrteV6TunnelTlvSlice = append(obj.bgpSrteV6TunnelTlvSlice, item)
	}
	return obj
}

func (obj *bgpSrteV6PolicyBgpSrteV6TunnelTlvIter) Set(index int, newObj BgpSrteV6TunnelTlv) BgpSrteV6PolicyBgpSrteV6TunnelTlvIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpSrteV6TunnelTlvSlice[index] = newObj
	return obj
}
func (obj *bgpSrteV6PolicyBgpSrteV6TunnelTlvIter) Clear() BgpSrteV6PolicyBgpSrteV6TunnelTlvIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpSrteV6TunnelTlv{}
		obj.bgpSrteV6TunnelTlvSlice = []BgpSrteV6TunnelTlv{}
	}
	return obj
}
func (obj *bgpSrteV6PolicyBgpSrteV6TunnelTlvIter) clearHolderSlice() BgpSrteV6PolicyBgpSrteV6TunnelTlvIter {
	if len(obj.bgpSrteV6TunnelTlvSlice) > 0 {
		obj.bgpSrteV6TunnelTlvSlice = []BgpSrteV6TunnelTlv{}
	}
	return obj
}
func (obj *bgpSrteV6PolicyBgpSrteV6TunnelTlvIter) appendHolderSlice(item BgpSrteV6TunnelTlv) BgpSrteV6PolicyBgpSrteV6TunnelTlvIter {
	obj.bgpSrteV6TunnelTlvSlice = append(obj.bgpSrteV6TunnelTlvSlice, item)
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpSrteV6Policy) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpSrteV6Policy object
func (obj *bgpSrteV6Policy) SetName(value string) BgpSrteV6Policy {

	obj.obj.Name = &value
	return obj
}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// Active returns a bool
func (obj *bgpSrteV6Policy) Active() bool {

	return *obj.obj.Active

}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// Active returns a bool
func (obj *bgpSrteV6Policy) HasActive() bool {
	return obj.obj.Active != nil
}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// SetActive sets the bool value in the BgpSrteV6Policy object
func (obj *bgpSrteV6Policy) SetActive(value bool) BgpSrteV6Policy {

	obj.obj.Active = &value
	return obj
}

func (obj *bgpSrteV6Policy) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Ipv6Endpoint is required
	if obj.obj.Ipv6Endpoint == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv6Endpoint is required field on interface BgpSrteV6Policy")
	}
	if obj.obj.Ipv6Endpoint != nil {

		err := obj.validateIpv6(obj.Ipv6Endpoint())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteV6Policy.Ipv6Endpoint"))
		}

	}

	if obj.obj.NextHopIpv4Address != nil {

		err := obj.validateIpv4(obj.NextHopIpv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteV6Policy.NextHopIpv4Address"))
		}

	}

	if obj.obj.NextHopIpv6Address != nil {

		err := obj.validateIpv6(obj.NextHopIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteV6Policy.NextHopIpv6Address"))
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

	if len(obj.obj.Extcommunities) != 0 {

		if set_default {
			obj.Extcommunities().clearHolderSlice()
			for _, item := range obj.obj.Extcommunities {
				obj.Extcommunities().appendHolderSlice(&bgpExtCommunity{obj: item})
			}
		}
		for _, item := range obj.Extcommunities().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.TunnelTlvs) != 0 {

		if set_default {
			obj.TunnelTlvs().clearHolderSlice()
			for _, item := range obj.obj.TunnelTlvs {
				obj.TunnelTlvs().appendHolderSlice(&bgpSrteV6TunnelTlv{obj: item})
			}
		}
		for _, item := range obj.TunnelTlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpSrteV6Policy")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"BgpSrteV6Policy.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

}

func (obj *bgpSrteV6Policy) setDefault() {
	if obj.obj.Distinguisher == nil {
		obj.SetDistinguisher(1)
	}
	if obj.obj.Color == nil {
		obj.SetColor(100)
	}
	if obj.obj.NextHopMode == nil {
		obj.SetNextHopMode(BgpSrteV6PolicyNextHopMode.LOCAL_IP)

	}
	if obj.obj.NextHopAddressType == nil {
		obj.SetNextHopAddressType(BgpSrteV6PolicyNextHopAddressType.IPV6)

	}
	if obj.obj.NextHopIpv4Address == nil {
		obj.SetNextHopIpv4Address("0.0.0.0")
	}
	if obj.obj.NextHopIpv6Address == nil {
		obj.SetNextHopIpv6Address("::0")
	}
	if obj.obj.Active == nil {
		obj.SetActive(true)
	}

}
