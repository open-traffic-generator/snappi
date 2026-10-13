package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv6MplsVpnUnicastVrfRoute *****
type bgpPrefixIpv6MplsVpnUnicastVrfRoute struct {
	validation
	obj                       *otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute
	marshaller                marshalBgpPrefixIpv6MplsVpnUnicastVrfRoute
	unMarshaller              unMarshalBgpPrefixIpv6MplsVpnUnicastVrfRoute
	serviceBindingHolder      BgpPrefixIpv6MplsVpnUnicastServiceBinding
	communitiesHolder         BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter
	extendedCommunitiesHolder BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	asPathHolder              ResultBgpAsPath
}

func NewBgpPrefixIpv6MplsVpnUnicastVrfRoute() BgpPrefixIpv6MplsVpnUnicastVrfRoute {
	obj := bgpPrefixIpv6MplsVpnUnicastVrfRoute{obj: &otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) msg() *otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute {
	return obj.obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) setMsg(msg *otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute) BgpPrefixIpv6MplsVpnUnicastVrfRoute {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv6MplsVpnUnicastVrfRoute struct {
	obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute
}

type marshalBgpPrefixIpv6MplsVpnUnicastVrfRoute interface {
	// ToProto marshals BgpPrefixIpv6MplsVpnUnicastVrfRoute to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute
	ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute, error)
	// ToPbText marshals BgpPrefixIpv6MplsVpnUnicastVrfRoute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv6MplsVpnUnicastVrfRoute to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv6MplsVpnUnicastVrfRoute to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv6MplsVpnUnicastVrfRoute struct {
	obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute
}

type unMarshalBgpPrefixIpv6MplsVpnUnicastVrfRoute interface {
	// FromProto unmarshals BgpPrefixIpv6MplsVpnUnicastVrfRoute from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute
	FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute) (BgpPrefixIpv6MplsVpnUnicastVrfRoute, error)
	// FromPbText unmarshals BgpPrefixIpv6MplsVpnUnicastVrfRoute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv6MplsVpnUnicastVrfRoute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv6MplsVpnUnicastVrfRoute from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) Marshal() marshalBgpPrefixIpv6MplsVpnUnicastVrfRoute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv6MplsVpnUnicastVrfRoute{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastVrfRoute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv6MplsVpnUnicastVrfRoute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastVrfRoute) ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastVrfRoute) FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute) (BgpPrefixIpv6MplsVpnUnicastVrfRoute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastVrfRoute) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastVrfRoute) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastVrfRoute) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastVrfRoute) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastVrfRoute) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastVrfRoute) FromJson(value string) error {
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

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) Clone() (BgpPrefixIpv6MplsVpnUnicastVrfRoute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv6MplsVpnUnicastVrfRoute()
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

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) setNil() {
	obj.serviceBindingHolder = nil
	obj.communitiesHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv6MplsVpnUnicastVrfRoute is a VPN-IPv6 prefix as imported into a single local VRF, with the path attributes received for that NLRI (RFC 4364, RFC 4659).
type BgpPrefixIpv6MplsVpnUnicastVrfRoute interface {
	Validation
	// msg marshals BgpPrefixIpv6MplsVpnUnicastVrfRoute to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// setMsg unmarshals BgpPrefixIpv6MplsVpnUnicastVrfRoute from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute) BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv6MplsVpnUnicastVrfRoute
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastVrfRoute
	// validate validates BgpPrefixIpv6MplsVpnUnicastVrfRoute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv6MplsVpnUnicastVrfRoute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute.
	Name() string
	// SetName assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastVrfRoute
	SetName(value string) BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// HasName checks if Name has been set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	HasName() bool
	// RouteDistinguisher returns string, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute.
	RouteDistinguisher() string
	// SetRouteDistinguisher assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastVrfRoute
	SetRouteDistinguisher(value string) BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// HasRouteDistinguisher checks if RouteDistinguisher has been set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	HasRouteDistinguisher() bool
	// Ipv6Address returns string, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute.
	Ipv6Address() string
	// SetIpv6Address assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastVrfRoute
	SetIpv6Address(value string) BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// HasIpv6Address checks if Ipv6Address has been set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	HasIpv6Address() bool
	// PrefixLength returns uint32, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastVrfRoute
	SetPrefixLength(value uint32) BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	Origin() BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum
	// SetOrigin assigns BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum provided by user to BgpPrefixIpv6MplsVpnUnicastVrfRoute
	SetOrigin(value BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum) BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// HasOrigin checks if Origin has been set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastVrfRoute
	SetPathId(value uint32) BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// HasPathId checks if PathId has been set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	HasPathId() bool
	// Ipv4NextHop returns string, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute.
	Ipv4NextHop() string
	// SetIpv4NextHop assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastVrfRoute
	SetIpv4NextHop(value string) BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// HasIpv4NextHop checks if Ipv4NextHop has been set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	HasIpv4NextHop() bool
	// Ipv6NextHop returns string, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute.
	Ipv6NextHop() string
	// SetIpv6NextHop assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastVrfRoute
	SetIpv6NextHop(value string) BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// HasIpv6NextHop checks if Ipv6NextHop has been set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	HasIpv6NextHop() bool
	// ServiceBinding returns BgpPrefixIpv6MplsVpnUnicastServiceBinding, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute.
	// BgpPrefixIpv6MplsVpnUnicastServiceBinding is the VPN dataplane binding received for a VPN-IPv6 route, mirroring the Bgp.L3vpn.V6ServiceBinding choice on the configured route range. Currently the only defined choice is the VPN MPLS label (RFC 4364 Section 3, RFC 4659); kept as a choice, independent of the shared label field, so a future dataplane binding (for example an SRv6 Service SID, RFC 9252) can be reported as a new sibling choice value without restructuring this schema.
	ServiceBinding() BgpPrefixIpv6MplsVpnUnicastServiceBinding
	// SetServiceBinding assigns BgpPrefixIpv6MplsVpnUnicastServiceBinding provided by user to BgpPrefixIpv6MplsVpnUnicastVrfRoute.
	// BgpPrefixIpv6MplsVpnUnicastServiceBinding is the VPN dataplane binding received for a VPN-IPv6 route, mirroring the Bgp.L3vpn.V6ServiceBinding choice on the configured route range. Currently the only defined choice is the VPN MPLS label (RFC 4364 Section 3, RFC 4659); kept as a choice, independent of the shared label field, so a future dataplane binding (for example an SRv6 Service SID, RFC 9252) can be reported as a new sibling choice value without restructuring this schema.
	SetServiceBinding(value BgpPrefixIpv6MplsVpnUnicastServiceBinding) BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// HasServiceBinding checks if ServiceBinding has been set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	HasServiceBinding() bool
	// Communities returns BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIterIter, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	Communities() BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter
	// ExtendedCommunities returns BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIterIter, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	ExtendedCommunities() BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	// AsPath returns ResultBgpAsPath, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	AsPath() ResultBgpAsPath
	// SetAsPath assigns ResultBgpAsPath provided by user to BgpPrefixIpv6MplsVpnUnicastVrfRoute.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	SetAsPath(value ResultBgpAsPath) BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// HasAsPath checks if AsPath has been set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	HasAsPath() bool
	// LocalPreference returns uint32, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute.
	LocalPreference() uint32
	// SetLocalPreference assigns uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastVrfRoute
	SetLocalPreference(value uint32) BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// HasLocalPreference checks if LocalPreference has been set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	HasLocalPreference() bool
	// MultiExitDiscriminator returns uint32, set in BgpPrefixIpv6MplsVpnUnicastVrfRoute.
	MultiExitDiscriminator() uint32
	// SetMultiExitDiscriminator assigns uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastVrfRoute
	SetMultiExitDiscriminator(value uint32) BgpPrefixIpv6MplsVpnUnicastVrfRoute
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BgpPrefixIpv6MplsVpnUnicastVrfRoute
	HasMultiExitDiscriminator() bool
	setNil()
}

// The name of the local VRF (Bgp.L3vpn.Vrf) that imported this prefix.
// Name returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) Name() string {

	return *obj.obj.Name

}

// The name of the local VRF (Bgp.L3vpn.Vrf) that imported this prefix.
// Name returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) HasName() bool {
	return obj.obj.Name != nil
}

// The name of the local VRF (Bgp.L3vpn.Vrf) that imported this prefix.
// SetName sets the string value in the BgpPrefixIpv6MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) SetName(value string) BgpPrefixIpv6MplsVpnUnicastVrfRoute {

	obj.obj.Name = &value
	return obj
}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv6 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// RouteDistinguisher returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) RouteDistinguisher() string {

	return *obj.obj.RouteDistinguisher

}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv6 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// RouteDistinguisher returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) HasRouteDistinguisher() bool {
	return obj.obj.RouteDistinguisher != nil
}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv6 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// SetRouteDistinguisher sets the string value in the BgpPrefixIpv6MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) SetRouteDistinguisher(value string) BgpPrefixIpv6MplsVpnUnicastVrfRoute {

	obj.obj.RouteDistinguisher = &value
	return obj
}

// An IPv6 unicast address.
// Ipv6Address returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) Ipv6Address() string {

	return *obj.obj.Ipv6Address

}

// An IPv6 unicast address.
// Ipv6Address returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) HasIpv6Address() bool {
	return obj.obj.Ipv6Address != nil
}

// An IPv6 unicast address.
// SetIpv6Address sets the string value in the BgpPrefixIpv6MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) SetIpv6Address(value string) BgpPrefixIpv6MplsVpnUnicastVrfRoute {

	obj.obj.Ipv6Address = &value
	return obj
}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// description is TBD
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv6MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) SetPrefixLength(value uint32) BgpPrefixIpv6MplsVpnUnicastVrfRoute {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum string

// Enum of Origin on BgpPrefixIpv6MplsVpnUnicastVrfRoute
var BgpPrefixIpv6MplsVpnUnicastVrfRouteOrigin = struct {
	IGP        BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum
	EGP        BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum
	INCOMPLETE BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum
}{
	IGP:        BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum("igp"),
	EGP:        BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) Origin() BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum {
	return BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin of the prefix.
// Origin returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) SetOrigin(value BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum) BgpPrefixIpv6MplsVpnUnicastVrfRoute {
	intValue, ok := otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv6MplsVpnUnicastVrfRouteOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv6MplsVpnUnicastVrfRoute_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id.
// SetPathId sets the uint32 value in the BgpPrefixIpv6MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) SetPathId(value uint32) BgpPrefixIpv6MplsVpnUnicastVrfRoute {

	obj.obj.PathId = &value
	return obj
}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) Ipv4NextHop() string {

	return *obj.obj.Ipv4NextHop

}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) HasIpv4NextHop() bool {
	return obj.obj.Ipv4NextHop != nil
}

// The IPv4 address of the egress interface.
// SetIpv4NextHop sets the string value in the BgpPrefixIpv6MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) SetIpv4NextHop(value string) BgpPrefixIpv6MplsVpnUnicastVrfRoute {

	obj.obj.Ipv4NextHop = &value
	return obj
}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) Ipv6NextHop() string {

	return *obj.obj.Ipv6NextHop

}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) HasIpv6NextHop() bool {
	return obj.obj.Ipv6NextHop != nil
}

// The IPv6 address of the egress interface.
// SetIpv6NextHop sets the string value in the BgpPrefixIpv6MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) SetIpv6NextHop(value string) BgpPrefixIpv6MplsVpnUnicastVrfRoute {

	obj.obj.Ipv6NextHop = &value
	return obj
}

// The VPN dataplane binding received for this VPN-IPv6 prefix.
// ServiceBinding returns a BgpPrefixIpv6MplsVpnUnicastServiceBinding
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) ServiceBinding() BgpPrefixIpv6MplsVpnUnicastServiceBinding {
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
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) HasServiceBinding() bool {
	return obj.obj.ServiceBinding != nil
}

// The VPN dataplane binding received for this VPN-IPv6 prefix.
// SetServiceBinding sets the BgpPrefixIpv6MplsVpnUnicastServiceBinding value in the BgpPrefixIpv6MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) SetServiceBinding(value BgpPrefixIpv6MplsVpnUnicastServiceBinding) BgpPrefixIpv6MplsVpnUnicastVrfRoute {

	obj.serviceBindingHolder = nil
	obj.obj.ServiceBinding = value.msg()

	return obj
}

// Optional community attributes.
// Communities returns a []ResultBgpCommunity
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) Communities() BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.ResultBgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter struct {
	obj                     *bgpPrefixIpv6MplsVpnUnicastVrfRoute
	resultBgpCommunitySlice []ResultBgpCommunity
	fieldPtr                *[]*otg.ResultBgpCommunity
}

func newBgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter(ptr *[]*otg.ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	return &bgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter interface {
	setMsg(*bgpPrefixIpv6MplsVpnUnicastVrfRoute) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter
	Items() []ResultBgpCommunity
	Add() ResultBgpCommunity
	Append(items ...ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter
	Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter
	Clear() BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter
	clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter
	appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter) setMsg(msg *bgpPrefixIpv6MplsVpnUnicastVrfRoute) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultBgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter) Items() []ResultBgpCommunity {
	return obj.resultBgpCommunitySlice
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter) Add() ResultBgpCommunity {
	newObj := &otg.ResultBgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultBgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter) Append(items ...ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter) Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultBgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter) Clear() BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultBgpCommunity{}
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter) clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	if len(obj.resultBgpCommunitySlice) > 0 {
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter) appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	return obj
}

// Optional received Extended Community attributes, including the Route Target(s) (RFC 4360) attached to this VPN-IPv6 route. Each received Extended Community attribute is available for retrieval in two forms. Support of the 'raw' format in which all 8 bytes (16 hex characters) is always present and available for use. In addition, if supported by the implementation, the Extended Community attribute may also be retrieved in the 'structured' format which is an optional field.
// ExtendedCommunities returns a []ResultExtendedCommunity
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) ExtendedCommunities() BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.ResultExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter struct {
	obj                          *bgpPrefixIpv6MplsVpnUnicastVrfRoute
	resultExtendedCommunitySlice []ResultExtendedCommunity
	fieldPtr                     *[]*otg.ResultExtendedCommunity
}

func newBgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter(ptr *[]*otg.ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	return &bgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter interface {
	setMsg(*bgpPrefixIpv6MplsVpnUnicastVrfRoute) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	Items() []ResultExtendedCommunity
	Add() ResultExtendedCommunity
	Append(items ...ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	Clear() BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter) setMsg(msg *bgpPrefixIpv6MplsVpnUnicastVrfRoute) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter) Items() []ResultExtendedCommunity {
	return obj.resultExtendedCommunitySlice
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter) Add() ResultExtendedCommunity {
	newObj := &otg.ResultExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter) Append(items ...ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter) Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter) Clear() BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultExtendedCommunity{}
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter) clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	if len(obj.resultExtendedCommunitySlice) > 0 {
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter) appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) AsPath() ResultBgpAsPath {
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
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the ResultBgpAsPath value in the BgpPrefixIpv6MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) SetAsPath(value ResultBgpAsPath) BgpPrefixIpv6MplsVpnUnicastVrfRoute {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) LocalPreference() uint32 {

	return *obj.obj.LocalPreference

}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// SetLocalPreference sets the uint32 value in the BgpPrefixIpv6MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) SetLocalPreference(value uint32) BgpPrefixIpv6MplsVpnUnicastVrfRoute {

	obj.obj.LocalPreference = &value
	return obj
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) MultiExitDiscriminator() uint32 {

	return *obj.obj.MultiExitDiscriminator

}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// SetMultiExitDiscriminator sets the uint32 value in the BgpPrefixIpv6MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) SetMultiExitDiscriminator(value uint32) BgpPrefixIpv6MplsVpnUnicastVrfRoute {

	obj.obj.MultiExitDiscriminator = &value
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv6MplsVpnUnicastVrfRoute.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6MplsVpnUnicastVrfRoute.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6MplsVpnUnicastVrfRoute.Ipv6NextHop"))
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

func (obj *bgpPrefixIpv6MplsVpnUnicastVrfRoute) setDefault() {

}
