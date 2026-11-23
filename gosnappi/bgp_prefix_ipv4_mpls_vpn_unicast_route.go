package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv4MplsVpnUnicastRoute *****
type bgpPrefixIpv4MplsVpnUnicastRoute struct {
	validation
	obj                       *otg.BgpPrefixIpv4MplsVpnUnicastRoute
	marshaller                marshalBgpPrefixIpv4MplsVpnUnicastRoute
	unMarshaller              unMarshalBgpPrefixIpv4MplsVpnUnicastRoute
	communitiesHolder         BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter
	extendedCommunitiesHolder BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter
	asPathHolder              ResultBgpAsPath
}

func NewBgpPrefixIpv4MplsVpnUnicastRoute() BgpPrefixIpv4MplsVpnUnicastRoute {
	obj := bgpPrefixIpv4MplsVpnUnicastRoute{obj: &otg.BgpPrefixIpv4MplsVpnUnicastRoute{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) msg() *otg.BgpPrefixIpv4MplsVpnUnicastRoute {
	return obj.obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) setMsg(msg *otg.BgpPrefixIpv4MplsVpnUnicastRoute) BgpPrefixIpv4MplsVpnUnicastRoute {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv4MplsVpnUnicastRoute struct {
	obj *bgpPrefixIpv4MplsVpnUnicastRoute
}

type marshalBgpPrefixIpv4MplsVpnUnicastRoute interface {
	// ToProto marshals BgpPrefixIpv4MplsVpnUnicastRoute to protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastRoute
	ToProto() (*otg.BgpPrefixIpv4MplsVpnUnicastRoute, error)
	// ToPbText marshals BgpPrefixIpv4MplsVpnUnicastRoute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv4MplsVpnUnicastRoute to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv4MplsVpnUnicastRoute to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv4MplsVpnUnicastRoute struct {
	obj *bgpPrefixIpv4MplsVpnUnicastRoute
}

type unMarshalBgpPrefixIpv4MplsVpnUnicastRoute interface {
	// FromProto unmarshals BgpPrefixIpv4MplsVpnUnicastRoute from protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastRoute
	FromProto(msg *otg.BgpPrefixIpv4MplsVpnUnicastRoute) (BgpPrefixIpv4MplsVpnUnicastRoute, error)
	// FromPbText unmarshals BgpPrefixIpv4MplsVpnUnicastRoute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv4MplsVpnUnicastRoute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv4MplsVpnUnicastRoute from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) Marshal() marshalBgpPrefixIpv4MplsVpnUnicastRoute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv4MplsVpnUnicastRoute{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) Unmarshal() unMarshalBgpPrefixIpv4MplsVpnUnicastRoute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv4MplsVpnUnicastRoute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv4MplsVpnUnicastRoute) ToProto() (*otg.BgpPrefixIpv4MplsVpnUnicastRoute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastRoute) FromProto(msg *otg.BgpPrefixIpv4MplsVpnUnicastRoute) (BgpPrefixIpv4MplsVpnUnicastRoute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv4MplsVpnUnicastRoute) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastRoute) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsVpnUnicastRoute) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastRoute) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsVpnUnicastRoute) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastRoute) FromJson(value string) error {
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

func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) Clone() (BgpPrefixIpv4MplsVpnUnicastRoute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv4MplsVpnUnicastRoute()
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

func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) setNil() {
	obj.communitiesHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv4MplsVpnUnicastRoute is a VPN-IPv4 prefix imported into a VRF, with the path attributes received for that NLRI (RFC 4364).
type BgpPrefixIpv4MplsVpnUnicastRoute interface {
	Validation
	// msg marshals BgpPrefixIpv4MplsVpnUnicastRoute to protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastRoute
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv4MplsVpnUnicastRoute
	// setMsg unmarshals BgpPrefixIpv4MplsVpnUnicastRoute from protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastRoute
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv4MplsVpnUnicastRoute) BgpPrefixIpv4MplsVpnUnicastRoute
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv4MplsVpnUnicastRoute
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv4MplsVpnUnicastRoute
	// validate validates BgpPrefixIpv4MplsVpnUnicastRoute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv4MplsVpnUnicastRoute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouteDistinguisher returns string, set in BgpPrefixIpv4MplsVpnUnicastRoute.
	RouteDistinguisher() string
	// SetRouteDistinguisher assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastRoute
	SetRouteDistinguisher(value string) BgpPrefixIpv4MplsVpnUnicastRoute
	// HasRouteDistinguisher checks if RouteDistinguisher has been set in BgpPrefixIpv4MplsVpnUnicastRoute
	HasRouteDistinguisher() bool
	// Ipv4Address returns string, set in BgpPrefixIpv4MplsVpnUnicastRoute.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastRoute
	SetIpv4Address(value string) BgpPrefixIpv4MplsVpnUnicastRoute
	// HasIpv4Address checks if Ipv4Address has been set in BgpPrefixIpv4MplsVpnUnicastRoute
	HasIpv4Address() bool
	// PrefixLength returns uint32, set in BgpPrefixIpv4MplsVpnUnicastRoute.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastRoute
	SetPrefixLength(value uint32) BgpPrefixIpv4MplsVpnUnicastRoute
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv4MplsVpnUnicastRoute
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum, set in BgpPrefixIpv4MplsVpnUnicastRoute
	Origin() BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum
	// SetOrigin assigns BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum provided by user to BgpPrefixIpv4MplsVpnUnicastRoute
	SetOrigin(value BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum) BgpPrefixIpv4MplsVpnUnicastRoute
	// HasOrigin checks if Origin has been set in BgpPrefixIpv4MplsVpnUnicastRoute
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv4MplsVpnUnicastRoute.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastRoute
	SetPathId(value uint32) BgpPrefixIpv4MplsVpnUnicastRoute
	// HasPathId checks if PathId has been set in BgpPrefixIpv4MplsVpnUnicastRoute
	HasPathId() bool
	// Ipv4NextHop returns string, set in BgpPrefixIpv4MplsVpnUnicastRoute.
	Ipv4NextHop() string
	// SetIpv4NextHop assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastRoute
	SetIpv4NextHop(value string) BgpPrefixIpv4MplsVpnUnicastRoute
	// HasIpv4NextHop checks if Ipv4NextHop has been set in BgpPrefixIpv4MplsVpnUnicastRoute
	HasIpv4NextHop() bool
	// Ipv6NextHop returns string, set in BgpPrefixIpv4MplsVpnUnicastRoute.
	Ipv6NextHop() string
	// SetIpv6NextHop assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastRoute
	SetIpv6NextHop(value string) BgpPrefixIpv4MplsVpnUnicastRoute
	// HasIpv6NextHop checks if Ipv6NextHop has been set in BgpPrefixIpv4MplsVpnUnicastRoute
	HasIpv6NextHop() bool
	// Label returns uint32, set in BgpPrefixIpv4MplsVpnUnicastRoute.
	Label() uint32
	// SetLabel assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastRoute
	SetLabel(value uint32) BgpPrefixIpv4MplsVpnUnicastRoute
	// HasLabel checks if Label has been set in BgpPrefixIpv4MplsVpnUnicastRoute
	HasLabel() bool
	// Communities returns BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIterIter, set in BgpPrefixIpv4MplsVpnUnicastRoute
	Communities() BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter
	// ExtendedCommunities returns BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIterIter, set in BgpPrefixIpv4MplsVpnUnicastRoute
	ExtendedCommunities() BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter
	// AsPath returns ResultBgpAsPath, set in BgpPrefixIpv4MplsVpnUnicastRoute.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	AsPath() ResultBgpAsPath
	// SetAsPath assigns ResultBgpAsPath provided by user to BgpPrefixIpv4MplsVpnUnicastRoute.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	SetAsPath(value ResultBgpAsPath) BgpPrefixIpv4MplsVpnUnicastRoute
	// HasAsPath checks if AsPath has been set in BgpPrefixIpv4MplsVpnUnicastRoute
	HasAsPath() bool
	// LocalPreference returns uint32, set in BgpPrefixIpv4MplsVpnUnicastRoute.
	LocalPreference() uint32
	// SetLocalPreference assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastRoute
	SetLocalPreference(value uint32) BgpPrefixIpv4MplsVpnUnicastRoute
	// HasLocalPreference checks if LocalPreference has been set in BgpPrefixIpv4MplsVpnUnicastRoute
	HasLocalPreference() bool
	// MultiExitDiscriminator returns uint32, set in BgpPrefixIpv4MplsVpnUnicastRoute.
	MultiExitDiscriminator() uint32
	// SetMultiExitDiscriminator assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastRoute
	SetMultiExitDiscriminator(value uint32) BgpPrefixIpv4MplsVpnUnicastRoute
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BgpPrefixIpv4MplsVpnUnicastRoute
	HasMultiExitDiscriminator() bool
	setNil()
}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv4 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// RouteDistinguisher returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) RouteDistinguisher() string {

	return *obj.obj.RouteDistinguisher

}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv4 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// RouteDistinguisher returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) HasRouteDistinguisher() bool {
	return obj.obj.RouteDistinguisher != nil
}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv4 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// SetRouteDistinguisher sets the string value in the BgpPrefixIpv4MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) SetRouteDistinguisher(value string) BgpPrefixIpv4MplsVpnUnicastRoute {

	obj.obj.RouteDistinguisher = &value
	return obj
}

// An IPv4 unicast address.
// Ipv4Address returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// An IPv4 unicast address.
// Ipv4Address returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// An IPv4 unicast address.
// SetIpv4Address sets the string value in the BgpPrefixIpv4MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) SetIpv4Address(value string) BgpPrefixIpv4MplsVpnUnicastRoute {

	obj.obj.Ipv4Address = &value
	return obj
}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// description is TBD
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) SetPrefixLength(value uint32) BgpPrefixIpv4MplsVpnUnicastRoute {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum string

// Enum of Origin on BgpPrefixIpv4MplsVpnUnicastRoute
var BgpPrefixIpv4MplsVpnUnicastRouteOrigin = struct {
	IGP        BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum
	EGP        BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum
	INCOMPLETE BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum
}{
	IGP:        BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum("igp"),
	EGP:        BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) Origin() BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum {
	return BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin of the prefix.
// Origin returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) SetOrigin(value BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum) BgpPrefixIpv4MplsVpnUnicastRoute {
	intValue, ok := otg.BgpPrefixIpv4MplsVpnUnicastRoute_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv4MplsVpnUnicastRouteOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv4MplsVpnUnicastRoute_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id.
// SetPathId sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) SetPathId(value uint32) BgpPrefixIpv4MplsVpnUnicastRoute {

	obj.obj.PathId = &value
	return obj
}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) Ipv4NextHop() string {

	return *obj.obj.Ipv4NextHop

}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) HasIpv4NextHop() bool {
	return obj.obj.Ipv4NextHop != nil
}

// The IPv4 address of the egress interface.
// SetIpv4NextHop sets the string value in the BgpPrefixIpv4MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) SetIpv4NextHop(value string) BgpPrefixIpv4MplsVpnUnicastRoute {

	obj.obj.Ipv4NextHop = &value
	return obj
}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) Ipv6NextHop() string {

	return *obj.obj.Ipv6NextHop

}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) HasIpv6NextHop() bool {
	return obj.obj.Ipv6NextHop != nil
}

// The IPv6 address of the egress interface.
// SetIpv6NextHop sets the string value in the BgpPrefixIpv4MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) SetIpv6NextHop(value string) BgpPrefixIpv4MplsVpnUnicastRoute {

	obj.obj.Ipv6NextHop = &value
	return obj
}

// The single VPN dataplane MPLS label (RFC 4364 Section 3) bound to this VPN-IPv4 prefix. RFC 4364 assigns exactly one label per route, matching the configured mpls_labels range on Bgp.L3vpn.V4RouteRange.
// Label returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) Label() uint32 {

	return *obj.obj.Label

}

// The single VPN dataplane MPLS label (RFC 4364 Section 3) bound to this VPN-IPv4 prefix. RFC 4364 assigns exactly one label per route, matching the configured mpls_labels range on Bgp.L3vpn.V4RouteRange.
// Label returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) HasLabel() bool {
	return obj.obj.Label != nil
}

// The single VPN dataplane MPLS label (RFC 4364 Section 3) bound to this VPN-IPv4 prefix. RFC 4364 assigns exactly one label per route, matching the configured mpls_labels range on Bgp.L3vpn.V4RouteRange.
// SetLabel sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) SetLabel(value uint32) BgpPrefixIpv4MplsVpnUnicastRoute {

	obj.obj.Label = &value
	return obj
}

// Optional community attributes.
// Communities returns a []ResultBgpCommunity
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) Communities() BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.ResultBgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter struct {
	obj                     *bgpPrefixIpv4MplsVpnUnicastRoute
	resultBgpCommunitySlice []ResultBgpCommunity
	fieldPtr                *[]*otg.ResultBgpCommunity
}

func newBgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter(ptr *[]*otg.ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter {
	return &bgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter interface {
	setMsg(*bgpPrefixIpv4MplsVpnUnicastRoute) BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter
	Items() []ResultBgpCommunity
	Add() ResultBgpCommunity
	Append(items ...ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter
	Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter
	Clear() BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter
	clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter
	appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter) setMsg(msg *bgpPrefixIpv4MplsVpnUnicastRoute) BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultBgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter) Items() []ResultBgpCommunity {
	return obj.resultBgpCommunitySlice
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter) Add() ResultBgpCommunity {
	newObj := &otg.ResultBgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultBgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter) Append(items ...ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter) Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultBgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter) Clear() BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultBgpCommunity{}
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter) clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter {
	if len(obj.resultBgpCommunitySlice) > 0 {
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter) appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultBgpCommunityIter {
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	return obj
}

// Optional received Extended Community attributes, including the Route Target(s) (RFC 4360) attached to this VPN-IPv4 route. Each received Extended Community attribute is available for retrieval in two forms. Support of the 'raw' format in which all 8 bytes (16 hex characters) is always present and available for use. In addition, if supported by the implementation, the Extended Community attribute may also be retrieved in the 'structured' format which is an optional field.
// ExtendedCommunities returns a []ResultExtendedCommunity
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) ExtendedCommunities() BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.ResultExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter struct {
	obj                          *bgpPrefixIpv4MplsVpnUnicastRoute
	resultExtendedCommunitySlice []ResultExtendedCommunity
	fieldPtr                     *[]*otg.ResultExtendedCommunity
}

func newBgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter(ptr *[]*otg.ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter {
	return &bgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter interface {
	setMsg(*bgpPrefixIpv4MplsVpnUnicastRoute) BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter
	Items() []ResultExtendedCommunity
	Add() ResultExtendedCommunity
	Append(items ...ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter
	Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter
	Clear() BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter
	clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter
	appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter) setMsg(msg *bgpPrefixIpv4MplsVpnUnicastRoute) BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter) Items() []ResultExtendedCommunity {
	return obj.resultExtendedCommunitySlice
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter) Add() ResultExtendedCommunity {
	newObj := &otg.ResultExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter) Append(items ...ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter) Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter) Clear() BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultExtendedCommunity{}
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter) clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter {
	if len(obj.resultExtendedCommunitySlice) > 0 {
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter) appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastRouteResultExtendedCommunityIter {
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) AsPath() ResultBgpAsPath {
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
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the ResultBgpAsPath value in the BgpPrefixIpv4MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) SetAsPath(value ResultBgpAsPath) BgpPrefixIpv4MplsVpnUnicastRoute {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) LocalPreference() uint32 {

	return *obj.obj.LocalPreference

}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// SetLocalPreference sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) SetLocalPreference(value uint32) BgpPrefixIpv4MplsVpnUnicastRoute {

	obj.obj.LocalPreference = &value
	return obj
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) MultiExitDiscriminator() uint32 {

	return *obj.obj.MultiExitDiscriminator

}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// SetMultiExitDiscriminator sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) SetMultiExitDiscriminator(value uint32) BgpPrefixIpv4MplsVpnUnicastRoute {

	obj.obj.MultiExitDiscriminator = &value
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv4MplsVpnUnicastRoute.PrefixLength <= 32 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4MplsVpnUnicastRoute.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4MplsVpnUnicastRoute.Ipv6NextHop"))
		}

	}

	if obj.obj.Label != nil {

		if *obj.obj.Label > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv4MplsVpnUnicastRoute.Label <= 1048575 but Got %d", *obj.obj.Label))
		}

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

func (obj *bgpPrefixIpv4MplsVpnUnicastRoute) setDefault() {

}
