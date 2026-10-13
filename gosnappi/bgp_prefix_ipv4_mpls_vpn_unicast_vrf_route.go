package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv4MplsVpnUnicastVrfRoute *****
type bgpPrefixIpv4MplsVpnUnicastVrfRoute struct {
	validation
	obj                       *otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute
	marshaller                marshalBgpPrefixIpv4MplsVpnUnicastVrfRoute
	unMarshaller              unMarshalBgpPrefixIpv4MplsVpnUnicastVrfRoute
	communitiesHolder         BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter
	extendedCommunitiesHolder BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	asPathHolder              ResultBgpAsPath
}

func NewBgpPrefixIpv4MplsVpnUnicastVrfRoute() BgpPrefixIpv4MplsVpnUnicastVrfRoute {
	obj := bgpPrefixIpv4MplsVpnUnicastVrfRoute{obj: &otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) msg() *otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute {
	return obj.obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) setMsg(msg *otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute) BgpPrefixIpv4MplsVpnUnicastVrfRoute {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv4MplsVpnUnicastVrfRoute struct {
	obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute
}

type marshalBgpPrefixIpv4MplsVpnUnicastVrfRoute interface {
	// ToProto marshals BgpPrefixIpv4MplsVpnUnicastVrfRoute to protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute
	ToProto() (*otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute, error)
	// ToPbText marshals BgpPrefixIpv4MplsVpnUnicastVrfRoute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv4MplsVpnUnicastVrfRoute to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv4MplsVpnUnicastVrfRoute to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv4MplsVpnUnicastVrfRoute struct {
	obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute
}

type unMarshalBgpPrefixIpv4MplsVpnUnicastVrfRoute interface {
	// FromProto unmarshals BgpPrefixIpv4MplsVpnUnicastVrfRoute from protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute
	FromProto(msg *otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute) (BgpPrefixIpv4MplsVpnUnicastVrfRoute, error)
	// FromPbText unmarshals BgpPrefixIpv4MplsVpnUnicastVrfRoute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv4MplsVpnUnicastVrfRoute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv4MplsVpnUnicastVrfRoute from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) Marshal() marshalBgpPrefixIpv4MplsVpnUnicastVrfRoute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv4MplsVpnUnicastVrfRoute{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) Unmarshal() unMarshalBgpPrefixIpv4MplsVpnUnicastVrfRoute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv4MplsVpnUnicastVrfRoute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv4MplsVpnUnicastVrfRoute) ToProto() (*otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastVrfRoute) FromProto(msg *otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute) (BgpPrefixIpv4MplsVpnUnicastVrfRoute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv4MplsVpnUnicastVrfRoute) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastVrfRoute) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsVpnUnicastVrfRoute) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastVrfRoute) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsVpnUnicastVrfRoute) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastVrfRoute) FromJson(value string) error {
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

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) Clone() (BgpPrefixIpv4MplsVpnUnicastVrfRoute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv4MplsVpnUnicastVrfRoute()
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

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) setNil() {
	obj.communitiesHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv4MplsVpnUnicastVrfRoute is a VPN-IPv4 prefix as imported into a single local VRF, with the path attributes received for that NLRI (RFC 4364).
type BgpPrefixIpv4MplsVpnUnicastVrfRoute interface {
	Validation
	// msg marshals BgpPrefixIpv4MplsVpnUnicastVrfRoute to protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// setMsg unmarshals BgpPrefixIpv4MplsVpnUnicastVrfRoute from protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute) BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv4MplsVpnUnicastVrfRoute
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv4MplsVpnUnicastVrfRoute
	// validate validates BgpPrefixIpv4MplsVpnUnicastVrfRoute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv4MplsVpnUnicastVrfRoute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute.
	Name() string
	// SetName assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastVrfRoute
	SetName(value string) BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// HasName checks if Name has been set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	HasName() bool
	// RouteDistinguisher returns string, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute.
	RouteDistinguisher() string
	// SetRouteDistinguisher assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastVrfRoute
	SetRouteDistinguisher(value string) BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// HasRouteDistinguisher checks if RouteDistinguisher has been set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	HasRouteDistinguisher() bool
	// Ipv4Address returns string, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastVrfRoute
	SetIpv4Address(value string) BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// HasIpv4Address checks if Ipv4Address has been set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	HasIpv4Address() bool
	// PrefixLength returns uint32, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastVrfRoute
	SetPrefixLength(value uint32) BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	Origin() BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum
	// SetOrigin assigns BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum provided by user to BgpPrefixIpv4MplsVpnUnicastVrfRoute
	SetOrigin(value BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum) BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// HasOrigin checks if Origin has been set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastVrfRoute
	SetPathId(value uint32) BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// HasPathId checks if PathId has been set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	HasPathId() bool
	// Ipv4NextHop returns string, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute.
	Ipv4NextHop() string
	// SetIpv4NextHop assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastVrfRoute
	SetIpv4NextHop(value string) BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// HasIpv4NextHop checks if Ipv4NextHop has been set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	HasIpv4NextHop() bool
	// Ipv6NextHop returns string, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute.
	Ipv6NextHop() string
	// SetIpv6NextHop assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastVrfRoute
	SetIpv6NextHop(value string) BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// HasIpv6NextHop checks if Ipv6NextHop has been set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	HasIpv6NextHop() bool
	// Label returns uint32, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute.
	Label() uint32
	// SetLabel assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastVrfRoute
	SetLabel(value uint32) BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// HasLabel checks if Label has been set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	HasLabel() bool
	// Communities returns BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIterIter, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	Communities() BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter
	// ExtendedCommunities returns BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIterIter, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	ExtendedCommunities() BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	// AsPath returns ResultBgpAsPath, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	AsPath() ResultBgpAsPath
	// SetAsPath assigns ResultBgpAsPath provided by user to BgpPrefixIpv4MplsVpnUnicastVrfRoute.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	SetAsPath(value ResultBgpAsPath) BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// HasAsPath checks if AsPath has been set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	HasAsPath() bool
	// LocalPreference returns uint32, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute.
	LocalPreference() uint32
	// SetLocalPreference assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastVrfRoute
	SetLocalPreference(value uint32) BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// HasLocalPreference checks if LocalPreference has been set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	HasLocalPreference() bool
	// MultiExitDiscriminator returns uint32, set in BgpPrefixIpv4MplsVpnUnicastVrfRoute.
	MultiExitDiscriminator() uint32
	// SetMultiExitDiscriminator assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastVrfRoute
	SetMultiExitDiscriminator(value uint32) BgpPrefixIpv4MplsVpnUnicastVrfRoute
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BgpPrefixIpv4MplsVpnUnicastVrfRoute
	HasMultiExitDiscriminator() bool
	setNil()
}

// The name of the local VRF (Bgp.L3vpn.Vrf) that imported this prefix.
// Name returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) Name() string {

	return *obj.obj.Name

}

// The name of the local VRF (Bgp.L3vpn.Vrf) that imported this prefix.
// Name returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) HasName() bool {
	return obj.obj.Name != nil
}

// The name of the local VRF (Bgp.L3vpn.Vrf) that imported this prefix.
// SetName sets the string value in the BgpPrefixIpv4MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) SetName(value string) BgpPrefixIpv4MplsVpnUnicastVrfRoute {

	obj.obj.Name = &value
	return obj
}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv4 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// RouteDistinguisher returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) RouteDistinguisher() string {

	return *obj.obj.RouteDistinguisher

}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv4 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// RouteDistinguisher returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) HasRouteDistinguisher() bool {
	return obj.obj.RouteDistinguisher != nil
}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv4 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// SetRouteDistinguisher sets the string value in the BgpPrefixIpv4MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) SetRouteDistinguisher(value string) BgpPrefixIpv4MplsVpnUnicastVrfRoute {

	obj.obj.RouteDistinguisher = &value
	return obj
}

// An IPv4 unicast address.
// Ipv4Address returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// An IPv4 unicast address.
// Ipv4Address returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// An IPv4 unicast address.
// SetIpv4Address sets the string value in the BgpPrefixIpv4MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) SetIpv4Address(value string) BgpPrefixIpv4MplsVpnUnicastVrfRoute {

	obj.obj.Ipv4Address = &value
	return obj
}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// description is TBD
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) SetPrefixLength(value uint32) BgpPrefixIpv4MplsVpnUnicastVrfRoute {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum string

// Enum of Origin on BgpPrefixIpv4MplsVpnUnicastVrfRoute
var BgpPrefixIpv4MplsVpnUnicastVrfRouteOrigin = struct {
	IGP        BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum
	EGP        BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum
	INCOMPLETE BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum
}{
	IGP:        BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum("igp"),
	EGP:        BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) Origin() BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum {
	return BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin of the prefix.
// Origin returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) SetOrigin(value BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum) BgpPrefixIpv4MplsVpnUnicastVrfRoute {
	intValue, ok := otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv4MplsVpnUnicastVrfRouteOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv4MplsVpnUnicastVrfRoute_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id.
// SetPathId sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) SetPathId(value uint32) BgpPrefixIpv4MplsVpnUnicastVrfRoute {

	obj.obj.PathId = &value
	return obj
}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) Ipv4NextHop() string {

	return *obj.obj.Ipv4NextHop

}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) HasIpv4NextHop() bool {
	return obj.obj.Ipv4NextHop != nil
}

// The IPv4 address of the egress interface.
// SetIpv4NextHop sets the string value in the BgpPrefixIpv4MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) SetIpv4NextHop(value string) BgpPrefixIpv4MplsVpnUnicastVrfRoute {

	obj.obj.Ipv4NextHop = &value
	return obj
}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) Ipv6NextHop() string {

	return *obj.obj.Ipv6NextHop

}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) HasIpv6NextHop() bool {
	return obj.obj.Ipv6NextHop != nil
}

// The IPv6 address of the egress interface.
// SetIpv6NextHop sets the string value in the BgpPrefixIpv4MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) SetIpv6NextHop(value string) BgpPrefixIpv4MplsVpnUnicastVrfRoute {

	obj.obj.Ipv6NextHop = &value
	return obj
}

// The single VPN dataplane MPLS label (RFC 4364 Section 3) bound to this VPN-IPv4 prefix. RFC 4364 assigns exactly one label per route, matching the configured mpls_labels range on Bgp.L3vpn.V4RouteRange.
// Label returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) Label() uint32 {

	return *obj.obj.Label

}

// The single VPN dataplane MPLS label (RFC 4364 Section 3) bound to this VPN-IPv4 prefix. RFC 4364 assigns exactly one label per route, matching the configured mpls_labels range on Bgp.L3vpn.V4RouteRange.
// Label returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) HasLabel() bool {
	return obj.obj.Label != nil
}

// The single VPN dataplane MPLS label (RFC 4364 Section 3) bound to this VPN-IPv4 prefix. RFC 4364 assigns exactly one label per route, matching the configured mpls_labels range on Bgp.L3vpn.V4RouteRange.
// SetLabel sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) SetLabel(value uint32) BgpPrefixIpv4MplsVpnUnicastVrfRoute {

	obj.obj.Label = &value
	return obj
}

// Optional community attributes.
// Communities returns a []ResultBgpCommunity
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) Communities() BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.ResultBgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter struct {
	obj                     *bgpPrefixIpv4MplsVpnUnicastVrfRoute
	resultBgpCommunitySlice []ResultBgpCommunity
	fieldPtr                *[]*otg.ResultBgpCommunity
}

func newBgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter(ptr *[]*otg.ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	return &bgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter interface {
	setMsg(*bgpPrefixIpv4MplsVpnUnicastVrfRoute) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter
	Items() []ResultBgpCommunity
	Add() ResultBgpCommunity
	Append(items ...ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter
	Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter
	Clear() BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter
	clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter
	appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter) setMsg(msg *bgpPrefixIpv4MplsVpnUnicastVrfRoute) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultBgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter) Items() []ResultBgpCommunity {
	return obj.resultBgpCommunitySlice
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter) Add() ResultBgpCommunity {
	newObj := &otg.ResultBgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultBgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter) Append(items ...ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter) Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultBgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter) Clear() BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultBgpCommunity{}
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter) clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	if len(obj.resultBgpCommunitySlice) > 0 {
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter) appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultBgpCommunityIter {
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	return obj
}

// Optional received Extended Community attributes, including the Route Target(s) (RFC 4360) attached to this VPN-IPv4 route. Each received Extended Community attribute is available for retrieval in two forms. Support of the 'raw' format in which all 8 bytes (16 hex characters) is always present and available for use. In addition, if supported by the implementation, the Extended Community attribute may also be retrieved in the 'structured' format which is an optional field.
// ExtendedCommunities returns a []ResultExtendedCommunity
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) ExtendedCommunities() BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.ResultExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter struct {
	obj                          *bgpPrefixIpv4MplsVpnUnicastVrfRoute
	resultExtendedCommunitySlice []ResultExtendedCommunity
	fieldPtr                     *[]*otg.ResultExtendedCommunity
}

func newBgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter(ptr *[]*otg.ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	return &bgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter interface {
	setMsg(*bgpPrefixIpv4MplsVpnUnicastVrfRoute) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	Items() []ResultExtendedCommunity
	Add() ResultExtendedCommunity
	Append(items ...ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	Clear() BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter
	appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter) setMsg(msg *bgpPrefixIpv4MplsVpnUnicastVrfRoute) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter) Items() []ResultExtendedCommunity {
	return obj.resultExtendedCommunitySlice
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter) Add() ResultExtendedCommunity {
	newObj := &otg.ResultExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter) Append(items ...ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter) Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter) Clear() BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultExtendedCommunity{}
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter) clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	if len(obj.resultExtendedCommunitySlice) > 0 {
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter) appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastVrfRouteResultExtendedCommunityIter {
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) AsPath() ResultBgpAsPath {
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
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the ResultBgpAsPath value in the BgpPrefixIpv4MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) SetAsPath(value ResultBgpAsPath) BgpPrefixIpv4MplsVpnUnicastVrfRoute {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) LocalPreference() uint32 {

	return *obj.obj.LocalPreference

}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// SetLocalPreference sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) SetLocalPreference(value uint32) BgpPrefixIpv4MplsVpnUnicastVrfRoute {

	obj.obj.LocalPreference = &value
	return obj
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) MultiExitDiscriminator() uint32 {

	return *obj.obj.MultiExitDiscriminator

}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// SetMultiExitDiscriminator sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastVrfRoute object
func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) SetMultiExitDiscriminator(value uint32) BgpPrefixIpv4MplsVpnUnicastVrfRoute {

	obj.obj.MultiExitDiscriminator = &value
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv4MplsVpnUnicastVrfRoute.PrefixLength <= 32 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4MplsVpnUnicastVrfRoute.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4MplsVpnUnicastVrfRoute.Ipv6NextHop"))
		}

	}

	if obj.obj.Label != nil {

		if *obj.obj.Label > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv4MplsVpnUnicastVrfRoute.Label <= 1048575 but Got %d", *obj.obj.Label))
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

func (obj *bgpPrefixIpv4MplsVpnUnicastVrfRoute) setDefault() {

}
