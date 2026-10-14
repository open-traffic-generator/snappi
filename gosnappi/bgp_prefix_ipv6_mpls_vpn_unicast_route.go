package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv6MplsVpnUnicastRoute *****
type bgpPrefixIpv6MplsVpnUnicastRoute struct {
	validation
	obj                       *otg.BgpPrefixIpv6MplsVpnUnicastRoute
	marshaller                marshalBgpPrefixIpv6MplsVpnUnicastRoute
	unMarshaller              unMarshalBgpPrefixIpv6MplsVpnUnicastRoute
	serviceBindingHolder      BgpPrefixIpv6MplsVpnUnicastServiceBinding
	communitiesHolder         BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter
	extendedCommunitiesHolder BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter
	asPathHolder              ResultBgpAsPath
}

func NewBgpPrefixIpv6MplsVpnUnicastRoute() BgpPrefixIpv6MplsVpnUnicastRoute {
	obj := bgpPrefixIpv6MplsVpnUnicastRoute{obj: &otg.BgpPrefixIpv6MplsVpnUnicastRoute{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) msg() *otg.BgpPrefixIpv6MplsVpnUnicastRoute {
	return obj.obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) setMsg(msg *otg.BgpPrefixIpv6MplsVpnUnicastRoute) BgpPrefixIpv6MplsVpnUnicastRoute {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv6MplsVpnUnicastRoute struct {
	obj *bgpPrefixIpv6MplsVpnUnicastRoute
}

type marshalBgpPrefixIpv6MplsVpnUnicastRoute interface {
	// ToProto marshals BgpPrefixIpv6MplsVpnUnicastRoute to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastRoute
	ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastRoute, error)
	// ToPbText marshals BgpPrefixIpv6MplsVpnUnicastRoute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv6MplsVpnUnicastRoute to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv6MplsVpnUnicastRoute to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv6MplsVpnUnicastRoute struct {
	obj *bgpPrefixIpv6MplsVpnUnicastRoute
}

type unMarshalBgpPrefixIpv6MplsVpnUnicastRoute interface {
	// FromProto unmarshals BgpPrefixIpv6MplsVpnUnicastRoute from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastRoute
	FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastRoute) (BgpPrefixIpv6MplsVpnUnicastRoute, error)
	// FromPbText unmarshals BgpPrefixIpv6MplsVpnUnicastRoute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv6MplsVpnUnicastRoute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv6MplsVpnUnicastRoute from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) Marshal() marshalBgpPrefixIpv6MplsVpnUnicastRoute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv6MplsVpnUnicastRoute{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastRoute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv6MplsVpnUnicastRoute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastRoute) ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastRoute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastRoute) FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastRoute) (BgpPrefixIpv6MplsVpnUnicastRoute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastRoute) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastRoute) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastRoute) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastRoute) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastRoute) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastRoute) FromJson(value string) error {
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

func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) Clone() (BgpPrefixIpv6MplsVpnUnicastRoute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv6MplsVpnUnicastRoute()
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

func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) setNil() {
	obj.serviceBindingHolder = nil
	obj.communitiesHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv6MplsVpnUnicastRoute is a VPN-IPv6 prefix imported into a VRF, with the path attributes received for that NLRI (RFC 4364, RFC 4659).
type BgpPrefixIpv6MplsVpnUnicastRoute interface {
	Validation
	// msg marshals BgpPrefixIpv6MplsVpnUnicastRoute to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastRoute
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv6MplsVpnUnicastRoute
	// setMsg unmarshals BgpPrefixIpv6MplsVpnUnicastRoute from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastRoute
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv6MplsVpnUnicastRoute) BgpPrefixIpv6MplsVpnUnicastRoute
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv6MplsVpnUnicastRoute
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastRoute
	// validate validates BgpPrefixIpv6MplsVpnUnicastRoute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv6MplsVpnUnicastRoute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouteDistinguisher returns string, set in BgpPrefixIpv6MplsVpnUnicastRoute.
	RouteDistinguisher() string
	// SetRouteDistinguisher assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastRoute
	SetRouteDistinguisher(value string) BgpPrefixIpv6MplsVpnUnicastRoute
	// HasRouteDistinguisher checks if RouteDistinguisher has been set in BgpPrefixIpv6MplsVpnUnicastRoute
	HasRouteDistinguisher() bool
	// Ipv6Address returns string, set in BgpPrefixIpv6MplsVpnUnicastRoute.
	Ipv6Address() string
	// SetIpv6Address assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastRoute
	SetIpv6Address(value string) BgpPrefixIpv6MplsVpnUnicastRoute
	// HasIpv6Address checks if Ipv6Address has been set in BgpPrefixIpv6MplsVpnUnicastRoute
	HasIpv6Address() bool
	// PrefixLength returns uint32, set in BgpPrefixIpv6MplsVpnUnicastRoute.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastRoute
	SetPrefixLength(value uint32) BgpPrefixIpv6MplsVpnUnicastRoute
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv6MplsVpnUnicastRoute
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum, set in BgpPrefixIpv6MplsVpnUnicastRoute
	Origin() BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum
	// SetOrigin assigns BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum provided by user to BgpPrefixIpv6MplsVpnUnicastRoute
	SetOrigin(value BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum) BgpPrefixIpv6MplsVpnUnicastRoute
	// HasOrigin checks if Origin has been set in BgpPrefixIpv6MplsVpnUnicastRoute
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv6MplsVpnUnicastRoute.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastRoute
	SetPathId(value uint32) BgpPrefixIpv6MplsVpnUnicastRoute
	// HasPathId checks if PathId has been set in BgpPrefixIpv6MplsVpnUnicastRoute
	HasPathId() bool
	// Ipv4NextHop returns string, set in BgpPrefixIpv6MplsVpnUnicastRoute.
	Ipv4NextHop() string
	// SetIpv4NextHop assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastRoute
	SetIpv4NextHop(value string) BgpPrefixIpv6MplsVpnUnicastRoute
	// HasIpv4NextHop checks if Ipv4NextHop has been set in BgpPrefixIpv6MplsVpnUnicastRoute
	HasIpv4NextHop() bool
	// Ipv6NextHop returns string, set in BgpPrefixIpv6MplsVpnUnicastRoute.
	Ipv6NextHop() string
	// SetIpv6NextHop assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastRoute
	SetIpv6NextHop(value string) BgpPrefixIpv6MplsVpnUnicastRoute
	// HasIpv6NextHop checks if Ipv6NextHop has been set in BgpPrefixIpv6MplsVpnUnicastRoute
	HasIpv6NextHop() bool
	// ServiceBinding returns BgpPrefixIpv6MplsVpnUnicastServiceBinding, set in BgpPrefixIpv6MplsVpnUnicastRoute.
	// BgpPrefixIpv6MplsVpnUnicastServiceBinding is the VPN dataplane binding received for a VPN-IPv6 route, mirroring the Bgp.L3vpn.V6ServiceBinding choice on the configured route range. Currently the only defined choice is the VPN MPLS label (RFC 4364 Section 3, RFC 4659); kept as a choice, independent of the shared label field, so a future dataplane binding (for example an SRv6 Service SID, RFC 9252) can be reported as a new sibling choice value without restructuring this schema.
	ServiceBinding() BgpPrefixIpv6MplsVpnUnicastServiceBinding
	// SetServiceBinding assigns BgpPrefixIpv6MplsVpnUnicastServiceBinding provided by user to BgpPrefixIpv6MplsVpnUnicastRoute.
	// BgpPrefixIpv6MplsVpnUnicastServiceBinding is the VPN dataplane binding received for a VPN-IPv6 route, mirroring the Bgp.L3vpn.V6ServiceBinding choice on the configured route range. Currently the only defined choice is the VPN MPLS label (RFC 4364 Section 3, RFC 4659); kept as a choice, independent of the shared label field, so a future dataplane binding (for example an SRv6 Service SID, RFC 9252) can be reported as a new sibling choice value without restructuring this schema.
	SetServiceBinding(value BgpPrefixIpv6MplsVpnUnicastServiceBinding) BgpPrefixIpv6MplsVpnUnicastRoute
	// HasServiceBinding checks if ServiceBinding has been set in BgpPrefixIpv6MplsVpnUnicastRoute
	HasServiceBinding() bool
	// Communities returns BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIterIter, set in BgpPrefixIpv6MplsVpnUnicastRoute
	Communities() BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter
	// ExtendedCommunities returns BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIterIter, set in BgpPrefixIpv6MplsVpnUnicastRoute
	ExtendedCommunities() BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter
	// AsPath returns ResultBgpAsPath, set in BgpPrefixIpv6MplsVpnUnicastRoute.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	AsPath() ResultBgpAsPath
	// SetAsPath assigns ResultBgpAsPath provided by user to BgpPrefixIpv6MplsVpnUnicastRoute.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	SetAsPath(value ResultBgpAsPath) BgpPrefixIpv6MplsVpnUnicastRoute
	// HasAsPath checks if AsPath has been set in BgpPrefixIpv6MplsVpnUnicastRoute
	HasAsPath() bool
	// LocalPreference returns uint32, set in BgpPrefixIpv6MplsVpnUnicastRoute.
	LocalPreference() uint32
	// SetLocalPreference assigns uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastRoute
	SetLocalPreference(value uint32) BgpPrefixIpv6MplsVpnUnicastRoute
	// HasLocalPreference checks if LocalPreference has been set in BgpPrefixIpv6MplsVpnUnicastRoute
	HasLocalPreference() bool
	// MultiExitDiscriminator returns uint32, set in BgpPrefixIpv6MplsVpnUnicastRoute.
	MultiExitDiscriminator() uint32
	// SetMultiExitDiscriminator assigns uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastRoute
	SetMultiExitDiscriminator(value uint32) BgpPrefixIpv6MplsVpnUnicastRoute
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BgpPrefixIpv6MplsVpnUnicastRoute
	HasMultiExitDiscriminator() bool
	setNil()
}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv6 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// RouteDistinguisher returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) RouteDistinguisher() string {

	return *obj.obj.RouteDistinguisher

}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv6 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// RouteDistinguisher returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) HasRouteDistinguisher() bool {
	return obj.obj.RouteDistinguisher != nil
}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv6 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// SetRouteDistinguisher sets the string value in the BgpPrefixIpv6MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) SetRouteDistinguisher(value string) BgpPrefixIpv6MplsVpnUnicastRoute {

	obj.obj.RouteDistinguisher = &value
	return obj
}

// An IPv6 unicast address.
// Ipv6Address returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) Ipv6Address() string {

	return *obj.obj.Ipv6Address

}

// An IPv6 unicast address.
// Ipv6Address returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) HasIpv6Address() bool {
	return obj.obj.Ipv6Address != nil
}

// An IPv6 unicast address.
// SetIpv6Address sets the string value in the BgpPrefixIpv6MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) SetIpv6Address(value string) BgpPrefixIpv6MplsVpnUnicastRoute {

	obj.obj.Ipv6Address = &value
	return obj
}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// description is TBD
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv6MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) SetPrefixLength(value uint32) BgpPrefixIpv6MplsVpnUnicastRoute {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum string

// Enum of Origin on BgpPrefixIpv6MplsVpnUnicastRoute
var BgpPrefixIpv6MplsVpnUnicastRouteOrigin = struct {
	IGP        BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum
	EGP        BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum
	INCOMPLETE BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum
}{
	IGP:        BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum("igp"),
	EGP:        BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) Origin() BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum {
	return BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin of the prefix.
// Origin returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) SetOrigin(value BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum) BgpPrefixIpv6MplsVpnUnicastRoute {
	intValue, ok := otg.BgpPrefixIpv6MplsVpnUnicastRoute_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv6MplsVpnUnicastRouteOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv6MplsVpnUnicastRoute_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id.
// SetPathId sets the uint32 value in the BgpPrefixIpv6MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) SetPathId(value uint32) BgpPrefixIpv6MplsVpnUnicastRoute {

	obj.obj.PathId = &value
	return obj
}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) Ipv4NextHop() string {

	return *obj.obj.Ipv4NextHop

}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) HasIpv4NextHop() bool {
	return obj.obj.Ipv4NextHop != nil
}

// The IPv4 address of the egress interface.
// SetIpv4NextHop sets the string value in the BgpPrefixIpv6MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) SetIpv4NextHop(value string) BgpPrefixIpv6MplsVpnUnicastRoute {

	obj.obj.Ipv4NextHop = &value
	return obj
}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) Ipv6NextHop() string {

	return *obj.obj.Ipv6NextHop

}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) HasIpv6NextHop() bool {
	return obj.obj.Ipv6NextHop != nil
}

// The IPv6 address of the egress interface.
// SetIpv6NextHop sets the string value in the BgpPrefixIpv6MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) SetIpv6NextHop(value string) BgpPrefixIpv6MplsVpnUnicastRoute {

	obj.obj.Ipv6NextHop = &value
	return obj
}

// The VPN dataplane binding received for this VPN-IPv6 prefix.
// ServiceBinding returns a BgpPrefixIpv6MplsVpnUnicastServiceBinding
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) ServiceBinding() BgpPrefixIpv6MplsVpnUnicastServiceBinding {
	if obj.obj.ServiceBinding == nil {
		obj.obj.ServiceBinding = NewBgpPrefixIpv6MplsVpnUnicastServiceBinding().msg()
	}
	if obj.serviceBindingHolder == nil {
		obj.serviceBindingHolder = &bgpPrefixIpv6MplsVpnUnicastServiceBinding{obj: obj.obj.ServiceBinding}
	}
	return obj.serviceBindingHolder
}

// The VPN dataplane binding received for this VPN-IPv6 prefix.
// ServiceBinding returns a BgpPrefixIpv6MplsVpnUnicastServiceBinding
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) HasServiceBinding() bool {
	return obj.obj.ServiceBinding != nil
}

// The VPN dataplane binding received for this VPN-IPv6 prefix.
// SetServiceBinding sets the BgpPrefixIpv6MplsVpnUnicastServiceBinding value in the BgpPrefixIpv6MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) SetServiceBinding(value BgpPrefixIpv6MplsVpnUnicastServiceBinding) BgpPrefixIpv6MplsVpnUnicastRoute {

	obj.serviceBindingHolder = nil
	obj.obj.ServiceBinding = value.msg()

	return obj
}

// Optional community attributes.
// Communities returns a []ResultBgpCommunity
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) Communities() BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.ResultBgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter struct {
	obj                     *bgpPrefixIpv6MplsVpnUnicastRoute
	resultBgpCommunitySlice []ResultBgpCommunity
	fieldPtr                *[]*otg.ResultBgpCommunity
}

func newBgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter(ptr *[]*otg.ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter {
	return &bgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter interface {
	setMsg(*bgpPrefixIpv6MplsVpnUnicastRoute) BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter
	Items() []ResultBgpCommunity
	Add() ResultBgpCommunity
	Append(items ...ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter
	Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter
	Clear() BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter
	clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter
	appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter) setMsg(msg *bgpPrefixIpv6MplsVpnUnicastRoute) BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultBgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter) Items() []ResultBgpCommunity {
	return obj.resultBgpCommunitySlice
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter) Add() ResultBgpCommunity {
	newObj := &otg.ResultBgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultBgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter) Append(items ...ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter) Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultBgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter) Clear() BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultBgpCommunity{}
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter) clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter {
	if len(obj.resultBgpCommunitySlice) > 0 {
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter) appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultBgpCommunityIter {
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	return obj
}

// Optional received Extended Community attributes, including the Route Target(s) (RFC 4360) attached to this VPN-IPv6 route. Each received Extended Community attribute is available for retrieval in two forms. Support of the 'raw' format in which all 8 bytes (16 hex characters) is always present and available for use. In addition, if supported by the implementation, the Extended Community attribute may also be retrieved in the 'structured' format which is an optional field.
// ExtendedCommunities returns a []ResultExtendedCommunity
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) ExtendedCommunities() BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.ResultExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter struct {
	obj                          *bgpPrefixIpv6MplsVpnUnicastRoute
	resultExtendedCommunitySlice []ResultExtendedCommunity
	fieldPtr                     *[]*otg.ResultExtendedCommunity
}

func newBgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter(ptr *[]*otg.ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter {
	return &bgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter interface {
	setMsg(*bgpPrefixIpv6MplsVpnUnicastRoute) BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter
	Items() []ResultExtendedCommunity
	Add() ResultExtendedCommunity
	Append(items ...ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter
	Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter
	Clear() BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter
	clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter
	appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter) setMsg(msg *bgpPrefixIpv6MplsVpnUnicastRoute) BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter) Items() []ResultExtendedCommunity {
	return obj.resultExtendedCommunitySlice
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter) Add() ResultExtendedCommunity {
	newObj := &otg.ResultExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter) Append(items ...ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter) Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter) Clear() BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultExtendedCommunity{}
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter) clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter {
	if len(obj.resultExtendedCommunitySlice) > 0 {
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter) appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastRouteResultExtendedCommunityIter {
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) AsPath() ResultBgpAsPath {
	if obj.obj.AsPath == nil {
		obj.obj.AsPath = NewResultBgpAsPath().msg()
	}
	if obj.asPathHolder == nil {
		obj.asPathHolder = &resultBgpAsPath{obj: obj.obj.AsPath}
	}
	return obj.asPathHolder
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the ResultBgpAsPath value in the BgpPrefixIpv6MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) SetAsPath(value ResultBgpAsPath) BgpPrefixIpv6MplsVpnUnicastRoute {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) LocalPreference() uint32 {

	return *obj.obj.LocalPreference

}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// SetLocalPreference sets the uint32 value in the BgpPrefixIpv6MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) SetLocalPreference(value uint32) BgpPrefixIpv6MplsVpnUnicastRoute {

	obj.obj.LocalPreference = &value
	return obj
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) MultiExitDiscriminator() uint32 {

	return *obj.obj.MultiExitDiscriminator

}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// SetMultiExitDiscriminator sets the uint32 value in the BgpPrefixIpv6MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) SetMultiExitDiscriminator(value uint32) BgpPrefixIpv6MplsVpnUnicastRoute {

	obj.obj.MultiExitDiscriminator = &value
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv6MplsVpnUnicastRoute.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6MplsVpnUnicastRoute.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6MplsVpnUnicastRoute.Ipv6NextHop"))
		}

	}

	if obj.obj.ServiceBinding != nil {

		obj.ServiceBinding().validateObj(vObj, set_default)
	}

	if len(obj.obj.Communities) != 0 {

		if set_default {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				obj.Communities().appendHolderSlice(&resultBgpCommunity{obj: item})
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.ExtendedCommunities) != 0 {

		if set_default {
			obj.ExtendedCommunities().clearHolderSlice()
			for _, item := range obj.obj.ExtendedCommunities {
				obj.ExtendedCommunities().appendHolderSlice(&resultExtendedCommunity{obj: item})
			}
		}
		for _, item := range obj.ExtendedCommunities().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(vObj, set_default)
	}

}

func (obj *bgpPrefixIpv6MplsVpnUnicastRoute) setDefault() {

}
