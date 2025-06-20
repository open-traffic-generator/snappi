package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpCapability *****
type bgpCapability struct {
	validation
	obj          *otg.BgpCapability
	marshaller   marshalBgpCapability
	unMarshaller unMarshalBgpCapability
}

func NewBgpCapability() BgpCapability {
	obj := bgpCapability{obj: &otg.BgpCapability{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpCapability) msg() *otg.BgpCapability {
	return obj.obj
}

func (obj *bgpCapability) setMsg(msg *otg.BgpCapability) BgpCapability {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpCapability struct {
	obj *bgpCapability
}

type marshalBgpCapability interface {
	// ToProto marshals BgpCapability to protobuf object *otg.BgpCapability
	ToProto() (*otg.BgpCapability, error)
	// ToPbText marshals BgpCapability to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpCapability to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpCapability to JSON text
	ToJson() (string, error)
}

type unMarshalbgpCapability struct {
	obj *bgpCapability
}

type unMarshalBgpCapability interface {
	// FromProto unmarshals BgpCapability from protobuf object *otg.BgpCapability
	FromProto(msg *otg.BgpCapability) (BgpCapability, error)
	// FromPbText unmarshals BgpCapability from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpCapability from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpCapability from JSON text
	FromJson(value string) error
}

func (obj *bgpCapability) Marshal() marshalBgpCapability {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpCapability{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpCapability) Unmarshal() unMarshalBgpCapability {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpCapability{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpCapability) ToProto() (*otg.BgpCapability, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpCapability) FromProto(msg *otg.BgpCapability) (BgpCapability, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpCapability) ToPbText() (string, error) {
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

func (m *unMarshalbgpCapability) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalbgpCapability) ToYaml() (string, error) {
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

func (m *unMarshalbgpCapability) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalbgpCapability) ToJson() (string, error) {
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

func (m *unMarshalbgpCapability) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *bgpCapability) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpCapability) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpCapability) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpCapability) Clone() (BgpCapability, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpCapability()
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

// BgpCapability is configuration for BGP capability settings.
type BgpCapability interface {
	Validation
	// msg marshals BgpCapability to protobuf object *otg.BgpCapability
	// and doesn't set defaults
	msg() *otg.BgpCapability
	// setMsg unmarshals BgpCapability from protobuf object *otg.BgpCapability
	// and doesn't set defaults
	setMsg(*otg.BgpCapability) BgpCapability
	// provides marshal interface
	Marshal() marshalBgpCapability
	// provides unmarshal interface
	Unmarshal() unMarshalBgpCapability
	// validate validates BgpCapability
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpCapability, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Unicast returns bool, set in BgpCapability.
	Ipv4Unicast() bool
	// SetIpv4Unicast assigns bool provided by user to BgpCapability
	SetIpv4Unicast(value bool) BgpCapability
	// HasIpv4Unicast checks if Ipv4Unicast has been set in BgpCapability
	HasIpv4Unicast() bool
	// Ipv4Multicast returns bool, set in BgpCapability.
	Ipv4Multicast() bool
	// SetIpv4Multicast assigns bool provided by user to BgpCapability
	SetIpv4Multicast(value bool) BgpCapability
	// HasIpv4Multicast checks if Ipv4Multicast has been set in BgpCapability
	HasIpv4Multicast() bool
	// Ipv6Unicast returns bool, set in BgpCapability.
	Ipv6Unicast() bool
	// SetIpv6Unicast assigns bool provided by user to BgpCapability
	SetIpv6Unicast(value bool) BgpCapability
	// HasIpv6Unicast checks if Ipv6Unicast has been set in BgpCapability
	HasIpv6Unicast() bool
	// Ipv6Multicast returns bool, set in BgpCapability.
	Ipv6Multicast() bool
	// SetIpv6Multicast assigns bool provided by user to BgpCapability
	SetIpv6Multicast(value bool) BgpCapability
	// HasIpv6Multicast checks if Ipv6Multicast has been set in BgpCapability
	HasIpv6Multicast() bool
	// Vpls returns bool, set in BgpCapability.
	Vpls() bool
	// SetVpls assigns bool provided by user to BgpCapability
	SetVpls(value bool) BgpCapability
	// HasVpls checks if Vpls has been set in BgpCapability
	HasVpls() bool
	// RouteRefresh returns bool, set in BgpCapability.
	RouteRefresh() bool
	// SetRouteRefresh assigns bool provided by user to BgpCapability
	SetRouteRefresh(value bool) BgpCapability
	// HasRouteRefresh checks if RouteRefresh has been set in BgpCapability
	HasRouteRefresh() bool
	// RouteConstraint returns bool, set in BgpCapability.
	RouteConstraint() bool
	// SetRouteConstraint assigns bool provided by user to BgpCapability
	SetRouteConstraint(value bool) BgpCapability
	// HasRouteConstraint checks if RouteConstraint has been set in BgpCapability
	HasRouteConstraint() bool
	// LinkStateNonVpn returns bool, set in BgpCapability.
	LinkStateNonVpn() bool
	// SetLinkStateNonVpn assigns bool provided by user to BgpCapability
	SetLinkStateNonVpn(value bool) BgpCapability
	// HasLinkStateNonVpn checks if LinkStateNonVpn has been set in BgpCapability
	HasLinkStateNonVpn() bool
	// LinkStateVpn returns bool, set in BgpCapability.
	LinkStateVpn() bool
	// SetLinkStateVpn assigns bool provided by user to BgpCapability
	SetLinkStateVpn(value bool) BgpCapability
	// HasLinkStateVpn checks if LinkStateVpn has been set in BgpCapability
	HasLinkStateVpn() bool
	// Evpn returns bool, set in BgpCapability.
	Evpn() bool
	// SetEvpn assigns bool provided by user to BgpCapability
	SetEvpn(value bool) BgpCapability
	// HasEvpn checks if Evpn has been set in BgpCapability
	HasEvpn() bool
	// ExtendedNextHopEncoding returns bool, set in BgpCapability.
	ExtendedNextHopEncoding() bool
	// SetExtendedNextHopEncoding assigns bool provided by user to BgpCapability
	SetExtendedNextHopEncoding(value bool) BgpCapability
	// HasExtendedNextHopEncoding checks if ExtendedNextHopEncoding has been set in BgpCapability
	HasExtendedNextHopEncoding() bool
	// Ipv4MulticastVpn returns bool, set in BgpCapability.
	Ipv4MulticastVpn() bool
	// SetIpv4MulticastVpn assigns bool provided by user to BgpCapability
	SetIpv4MulticastVpn(value bool) BgpCapability
	// HasIpv4MulticastVpn checks if Ipv4MulticastVpn has been set in BgpCapability
	HasIpv4MulticastVpn() bool
	// Ipv4MplsVpn returns bool, set in BgpCapability.
	Ipv4MplsVpn() bool
	// SetIpv4MplsVpn assigns bool provided by user to BgpCapability
	SetIpv4MplsVpn(value bool) BgpCapability
	// HasIpv4MplsVpn checks if Ipv4MplsVpn has been set in BgpCapability
	HasIpv4MplsVpn() bool
	// Ipv4Mdt returns bool, set in BgpCapability.
	Ipv4Mdt() bool
	// SetIpv4Mdt assigns bool provided by user to BgpCapability
	SetIpv4Mdt(value bool) BgpCapability
	// HasIpv4Mdt checks if Ipv4Mdt has been set in BgpCapability
	HasIpv4Mdt() bool
	// Ipv4MulticastMplsVpn returns bool, set in BgpCapability.
	Ipv4MulticastMplsVpn() bool
	// SetIpv4MulticastMplsVpn assigns bool provided by user to BgpCapability
	SetIpv4MulticastMplsVpn(value bool) BgpCapability
	// HasIpv4MulticastMplsVpn checks if Ipv4MulticastMplsVpn has been set in BgpCapability
	HasIpv4MulticastMplsVpn() bool
	// Ipv4UnicastFlowSpec returns bool, set in BgpCapability.
	Ipv4UnicastFlowSpec() bool
	// SetIpv4UnicastFlowSpec assigns bool provided by user to BgpCapability
	SetIpv4UnicastFlowSpec(value bool) BgpCapability
	// HasIpv4UnicastFlowSpec checks if Ipv4UnicastFlowSpec has been set in BgpCapability
	HasIpv4UnicastFlowSpec() bool
	// Ipv4SrTePolicy returns bool, set in BgpCapability.
	Ipv4SrTePolicy() bool
	// SetIpv4SrTePolicy assigns bool provided by user to BgpCapability
	SetIpv4SrTePolicy(value bool) BgpCapability
	// HasIpv4SrTePolicy checks if Ipv4SrTePolicy has been set in BgpCapability
	HasIpv4SrTePolicy() bool
	// Ipv4UnicastAddPath returns bool, set in BgpCapability.
	Ipv4UnicastAddPath() bool
	// SetIpv4UnicastAddPath assigns bool provided by user to BgpCapability
	SetIpv4UnicastAddPath(value bool) BgpCapability
	// HasIpv4UnicastAddPath checks if Ipv4UnicastAddPath has been set in BgpCapability
	HasIpv4UnicastAddPath() bool
	// Ipv6MulticastVpn returns bool, set in BgpCapability.
	Ipv6MulticastVpn() bool
	// SetIpv6MulticastVpn assigns bool provided by user to BgpCapability
	SetIpv6MulticastVpn(value bool) BgpCapability
	// HasIpv6MulticastVpn checks if Ipv6MulticastVpn has been set in BgpCapability
	HasIpv6MulticastVpn() bool
	// Ipv6MplsVpn returns bool, set in BgpCapability.
	Ipv6MplsVpn() bool
	// SetIpv6MplsVpn assigns bool provided by user to BgpCapability
	SetIpv6MplsVpn(value bool) BgpCapability
	// HasIpv6MplsVpn checks if Ipv6MplsVpn has been set in BgpCapability
	HasIpv6MplsVpn() bool
	// Ipv6Mdt returns bool, set in BgpCapability.
	Ipv6Mdt() bool
	// SetIpv6Mdt assigns bool provided by user to BgpCapability
	SetIpv6Mdt(value bool) BgpCapability
	// HasIpv6Mdt checks if Ipv6Mdt has been set in BgpCapability
	HasIpv6Mdt() bool
	// Ipv6MulticastMplsVpn returns bool, set in BgpCapability.
	Ipv6MulticastMplsVpn() bool
	// SetIpv6MulticastMplsVpn assigns bool provided by user to BgpCapability
	SetIpv6MulticastMplsVpn(value bool) BgpCapability
	// HasIpv6MulticastMplsVpn checks if Ipv6MulticastMplsVpn has been set in BgpCapability
	HasIpv6MulticastMplsVpn() bool
	// Ipv6UnicastFlowSpec returns bool, set in BgpCapability.
	Ipv6UnicastFlowSpec() bool
	// SetIpv6UnicastFlowSpec assigns bool provided by user to BgpCapability
	SetIpv6UnicastFlowSpec(value bool) BgpCapability
	// HasIpv6UnicastFlowSpec checks if Ipv6UnicastFlowSpec has been set in BgpCapability
	HasIpv6UnicastFlowSpec() bool
	// Ipv6SrTePolicy returns bool, set in BgpCapability.
	Ipv6SrTePolicy() bool
	// SetIpv6SrTePolicy assigns bool provided by user to BgpCapability
	SetIpv6SrTePolicy(value bool) BgpCapability
	// HasIpv6SrTePolicy checks if Ipv6SrTePolicy has been set in BgpCapability
	HasIpv6SrTePolicy() bool
	// Ipv6UnicastAddPath returns bool, set in BgpCapability.
	Ipv6UnicastAddPath() bool
	// SetIpv6UnicastAddPath assigns bool provided by user to BgpCapability
	SetIpv6UnicastAddPath(value bool) BgpCapability
	// HasIpv6UnicastAddPath checks if Ipv6UnicastAddPath has been set in BgpCapability
	HasIpv6UnicastAddPath() bool
}

// Support for the IPv4 Unicast address family.
// Ipv4Unicast returns a bool
func (obj *bgpCapability) Ipv4Unicast() bool {

	return *obj.obj.Ipv4Unicast

}

// Support for the IPv4 Unicast address family.
// Ipv4Unicast returns a bool
func (obj *bgpCapability) HasIpv4Unicast() bool {
	return obj.obj.Ipv4Unicast != nil
}

// Support for the IPv4 Unicast address family.
// SetIpv4Unicast sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv4Unicast(value bool) BgpCapability {

	obj.obj.Ipv4Unicast = &value
	return obj
}

// Support for the IPv4 Multicast address family.
// Ipv4Multicast returns a bool
func (obj *bgpCapability) Ipv4Multicast() bool {

	return *obj.obj.Ipv4Multicast

}

// Support for the IPv4 Multicast address family.
// Ipv4Multicast returns a bool
func (obj *bgpCapability) HasIpv4Multicast() bool {
	return obj.obj.Ipv4Multicast != nil
}

// Support for the IPv4 Multicast address family.
// SetIpv4Multicast sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv4Multicast(value bool) BgpCapability {

	obj.obj.Ipv4Multicast = &value
	return obj
}

// Support for the IPv4 Unicast address family.
// Ipv6Unicast returns a bool
func (obj *bgpCapability) Ipv6Unicast() bool {

	return *obj.obj.Ipv6Unicast

}

// Support for the IPv4 Unicast address family.
// Ipv6Unicast returns a bool
func (obj *bgpCapability) HasIpv6Unicast() bool {
	return obj.obj.Ipv6Unicast != nil
}

// Support for the IPv4 Unicast address family.
// SetIpv6Unicast sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv6Unicast(value bool) BgpCapability {

	obj.obj.Ipv6Unicast = &value
	return obj
}

// Support for the IPv6 Multicast address family.
// Ipv6Multicast returns a bool
func (obj *bgpCapability) Ipv6Multicast() bool {

	return *obj.obj.Ipv6Multicast

}

// Support for the IPv6 Multicast address family.
// Ipv6Multicast returns a bool
func (obj *bgpCapability) HasIpv6Multicast() bool {
	return obj.obj.Ipv6Multicast != nil
}

// Support for the IPv6 Multicast address family.
// SetIpv6Multicast sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv6Multicast(value bool) BgpCapability {

	obj.obj.Ipv6Multicast = &value
	return obj
}

// Support for VPLS as below.
// RFC4761 - Virtual Private LAN Service (VPLS) using BGP for Auto-Discovery
// and Signaling.
// RFC6624 - Layer 2 Virtual Private Networks using BGP for Auto-Discovery
// and Signaling.
// Vpls returns a bool
func (obj *bgpCapability) Vpls() bool {

	return *obj.obj.Vpls

}

// Support for VPLS as below.
// RFC4761 - Virtual Private LAN Service (VPLS) using BGP for Auto-Discovery
// and Signaling.
// RFC6624 - Layer 2 Virtual Private Networks using BGP for Auto-Discovery
// and Signaling.
// Vpls returns a bool
func (obj *bgpCapability) HasVpls() bool {
	return obj.obj.Vpls != nil
}

// Support for VPLS as below.
// RFC4761 - Virtual Private LAN Service (VPLS) using BGP for Auto-Discovery
// and Signaling.
// RFC6624 - Layer 2 Virtual Private Networks using BGP for Auto-Discovery
// and Signaling.
// SetVpls sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetVpls(value bool) BgpCapability {

	obj.obj.Vpls = &value
	return obj
}

// Support for the route refresh capabilities. Route Refresh allows the dynamic exchange of route refresh requests  and routing information between BGP peers and the subsequent re-advertisement  of the outbound or inbound routing table.
// RouteRefresh returns a bool
func (obj *bgpCapability) RouteRefresh() bool {

	return *obj.obj.RouteRefresh

}

// Support for the route refresh capabilities. Route Refresh allows the dynamic exchange of route refresh requests  and routing information between BGP peers and the subsequent re-advertisement  of the outbound or inbound routing table.
// RouteRefresh returns a bool
func (obj *bgpCapability) HasRouteRefresh() bool {
	return obj.obj.RouteRefresh != nil
}

// Support for the route refresh capabilities. Route Refresh allows the dynamic exchange of route refresh requests  and routing information between BGP peers and the subsequent re-advertisement  of the outbound or inbound routing table.
// SetRouteRefresh sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetRouteRefresh(value bool) BgpCapability {

	obj.obj.RouteRefresh = &value
	return obj
}

// Supports for the route constraint capabilities. Route Constraint allows the advertisement of Route Target Membership  information. The BGP peers exchange Route Target Reachability Information,  which is used to build a route distribution graph.  This limits the propagation of VPN Network Layer Reachability Information (NLRI) between different autonomous systems  or distinct clusters of the same autonomous system. This is supported for Layer 3 Virtual Private Network scenario.
// RouteConstraint returns a bool
func (obj *bgpCapability) RouteConstraint() bool {

	return *obj.obj.RouteConstraint

}

// Supports for the route constraint capabilities. Route Constraint allows the advertisement of Route Target Membership  information. The BGP peers exchange Route Target Reachability Information,  which is used to build a route distribution graph.  This limits the propagation of VPN Network Layer Reachability Information (NLRI) between different autonomous systems  or distinct clusters of the same autonomous system. This is supported for Layer 3 Virtual Private Network scenario.
// RouteConstraint returns a bool
func (obj *bgpCapability) HasRouteConstraint() bool {
	return obj.obj.RouteConstraint != nil
}

// Supports for the route constraint capabilities. Route Constraint allows the advertisement of Route Target Membership  information. The BGP peers exchange Route Target Reachability Information,  which is used to build a route distribution graph.  This limits the propagation of VPN Network Layer Reachability Information (NLRI) between different autonomous systems  or distinct clusters of the same autonomous system. This is supported for Layer 3 Virtual Private Network scenario.
// SetRouteConstraint sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetRouteConstraint(value bool) BgpCapability {

	obj.obj.RouteConstraint = &value
	return obj
}

// Support for BGP Link State for ISIS and OSPF.
// LinkStateNonVpn returns a bool
func (obj *bgpCapability) LinkStateNonVpn() bool {

	return *obj.obj.LinkStateNonVpn

}

// Support for BGP Link State for ISIS and OSPF.
// LinkStateNonVpn returns a bool
func (obj *bgpCapability) HasLinkStateNonVpn() bool {
	return obj.obj.LinkStateNonVpn != nil
}

// Support for BGP Link State for ISIS and OSPF.
// SetLinkStateNonVpn sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetLinkStateNonVpn(value bool) BgpCapability {

	obj.obj.LinkStateNonVpn = &value
	return obj
}

// Capability advertisement of BGP Link State for VPNs.
// LinkStateVpn returns a bool
func (obj *bgpCapability) LinkStateVpn() bool {

	return *obj.obj.LinkStateVpn

}

// Capability advertisement of BGP Link State for VPNs.
// LinkStateVpn returns a bool
func (obj *bgpCapability) HasLinkStateVpn() bool {
	return obj.obj.LinkStateVpn != nil
}

// Capability advertisement of BGP Link State for VPNs.
// SetLinkStateVpn sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetLinkStateVpn(value bool) BgpCapability {

	obj.obj.LinkStateVpn = &value
	return obj
}

// Support for the EVPN address family.
// Evpn returns a bool
func (obj *bgpCapability) Evpn() bool {

	return *obj.obj.Evpn

}

// Support for the EVPN address family.
// Evpn returns a bool
func (obj *bgpCapability) HasEvpn() bool {
	return obj.obj.Evpn != nil
}

// Support for the EVPN address family.
// SetEvpn sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetEvpn(value bool) BgpCapability {

	obj.obj.Evpn = &value
	return obj
}

// Support for extended Next Hop Encoding for Nexthop field in  IPv4 routes advertisement.  This allows IPv4 routes being advertised by IPv6 peers to  include an IPv6 Nexthop.
// ExtendedNextHopEncoding returns a bool
func (obj *bgpCapability) ExtendedNextHopEncoding() bool {

	return *obj.obj.ExtendedNextHopEncoding

}

// Support for extended Next Hop Encoding for Nexthop field in  IPv4 routes advertisement.  This allows IPv4 routes being advertised by IPv6 peers to  include an IPv6 Nexthop.
// ExtendedNextHopEncoding returns a bool
func (obj *bgpCapability) HasExtendedNextHopEncoding() bool {
	return obj.obj.ExtendedNextHopEncoding != nil
}

// Support for extended Next Hop Encoding for Nexthop field in  IPv4 routes advertisement.  This allows IPv4 routes being advertised by IPv6 peers to  include an IPv6 Nexthop.
// SetExtendedNextHopEncoding sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetExtendedNextHopEncoding(value bool) BgpCapability {

	obj.obj.ExtendedNextHopEncoding = &value
	return obj
}

// Support for the IPv4 Multicast VPN address family.
// Ipv4MulticastVpn returns a bool
func (obj *bgpCapability) Ipv4MulticastVpn() bool {

	return *obj.obj.Ipv4MulticastVpn

}

// Support for the IPv4 Multicast VPN address family.
// Ipv4MulticastVpn returns a bool
func (obj *bgpCapability) HasIpv4MulticastVpn() bool {
	return obj.obj.Ipv4MulticastVpn != nil
}

// Support for the IPv4 Multicast VPN address family.
// SetIpv4MulticastVpn sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv4MulticastVpn(value bool) BgpCapability {

	obj.obj.Ipv4MulticastVpn = &value
	return obj
}

// Support for the IPv4 MPLS L3VPN address family.
// Ipv4MplsVpn returns a bool
func (obj *bgpCapability) Ipv4MplsVpn() bool {

	return *obj.obj.Ipv4MplsVpn

}

// Support for the IPv4 MPLS L3VPN address family.
// Ipv4MplsVpn returns a bool
func (obj *bgpCapability) HasIpv4MplsVpn() bool {
	return obj.obj.Ipv4MplsVpn != nil
}

// Support for the IPv4 MPLS L3VPN address family.
// SetIpv4MplsVpn sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv4MplsVpn(value bool) BgpCapability {

	obj.obj.Ipv4MplsVpn = &value
	return obj
}

// Supports for IPv4 MDT address family messages.
// Ipv4Mdt returns a bool
func (obj *bgpCapability) Ipv4Mdt() bool {

	return *obj.obj.Ipv4Mdt

}

// Supports for IPv4 MDT address family messages.
// Ipv4Mdt returns a bool
func (obj *bgpCapability) HasIpv4Mdt() bool {
	return obj.obj.Ipv4Mdt != nil
}

// Supports for IPv4 MDT address family messages.
// SetIpv4Mdt sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv4Mdt(value bool) BgpCapability {

	obj.obj.Ipv4Mdt = &value
	return obj
}

// Support for the IPv4 Multicast VPN address family.
// Ipv4MulticastMplsVpn returns a bool
func (obj *bgpCapability) Ipv4MulticastMplsVpn() bool {

	return *obj.obj.Ipv4MulticastMplsVpn

}

// Support for the IPv4 Multicast VPN address family.
// Ipv4MulticastMplsVpn returns a bool
func (obj *bgpCapability) HasIpv4MulticastMplsVpn() bool {
	return obj.obj.Ipv4MulticastMplsVpn != nil
}

// Support for the IPv4 Multicast VPN address family.
// SetIpv4MulticastMplsVpn sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv4MulticastMplsVpn(value bool) BgpCapability {

	obj.obj.Ipv4MulticastMplsVpn = &value
	return obj
}

// Support for propagation of IPv4 unicast flow specification rules.
// Ipv4UnicastFlowSpec returns a bool
func (obj *bgpCapability) Ipv4UnicastFlowSpec() bool {

	return *obj.obj.Ipv4UnicastFlowSpec

}

// Support for propagation of IPv4 unicast flow specification rules.
// Ipv4UnicastFlowSpec returns a bool
func (obj *bgpCapability) HasIpv4UnicastFlowSpec() bool {
	return obj.obj.Ipv4UnicastFlowSpec != nil
}

// Support for propagation of IPv4 unicast flow specification rules.
// SetIpv4UnicastFlowSpec sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv4UnicastFlowSpec(value bool) BgpCapability {

	obj.obj.Ipv4UnicastFlowSpec = &value
	return obj
}

// Support for IPv4 SRTE policy.
// Ipv4SrTePolicy returns a bool
func (obj *bgpCapability) Ipv4SrTePolicy() bool {

	return *obj.obj.Ipv4SrTePolicy

}

// Support for IPv4 SRTE policy.
// Ipv4SrTePolicy returns a bool
func (obj *bgpCapability) HasIpv4SrTePolicy() bool {
	return obj.obj.Ipv4SrTePolicy != nil
}

// Support for IPv4 SRTE policy.
// SetIpv4SrTePolicy sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv4SrTePolicy(value bool) BgpCapability {

	obj.obj.Ipv4SrTePolicy = &value
	return obj
}

// Support for IPv4 Unicast Add Path Capability.
// Ipv4UnicastAddPath returns a bool
func (obj *bgpCapability) Ipv4UnicastAddPath() bool {

	return *obj.obj.Ipv4UnicastAddPath

}

// Support for IPv4 Unicast Add Path Capability.
// Ipv4UnicastAddPath returns a bool
func (obj *bgpCapability) HasIpv4UnicastAddPath() bool {
	return obj.obj.Ipv4UnicastAddPath != nil
}

// Support for IPv4 Unicast Add Path Capability.
// SetIpv4UnicastAddPath sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv4UnicastAddPath(value bool) BgpCapability {

	obj.obj.Ipv4UnicastAddPath = &value
	return obj
}

// Support for the IPv6 Multicast VPN address family.
// Ipv6MulticastVpn returns a bool
func (obj *bgpCapability) Ipv6MulticastVpn() bool {

	return *obj.obj.Ipv6MulticastVpn

}

// Support for the IPv6 Multicast VPN address family.
// Ipv6MulticastVpn returns a bool
func (obj *bgpCapability) HasIpv6MulticastVpn() bool {
	return obj.obj.Ipv6MulticastVpn != nil
}

// Support for the IPv6 Multicast VPN address family.
// SetIpv6MulticastVpn sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv6MulticastVpn(value bool) BgpCapability {

	obj.obj.Ipv6MulticastVpn = &value
	return obj
}

// Support for the IPv6 MPLS L3VPN address family.
// Ipv6MplsVpn returns a bool
func (obj *bgpCapability) Ipv6MplsVpn() bool {

	return *obj.obj.Ipv6MplsVpn

}

// Support for the IPv6 MPLS L3VPN address family.
// Ipv6MplsVpn returns a bool
func (obj *bgpCapability) HasIpv6MplsVpn() bool {
	return obj.obj.Ipv6MplsVpn != nil
}

// Support for the IPv6 MPLS L3VPN address family.
// SetIpv6MplsVpn sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv6MplsVpn(value bool) BgpCapability {

	obj.obj.Ipv6MplsVpn = &value
	return obj
}

// Support for IPv6 MDT address family messages.
// Ipv6Mdt returns a bool
func (obj *bgpCapability) Ipv6Mdt() bool {

	return *obj.obj.Ipv6Mdt

}

// Support for IPv6 MDT address family messages.
// Ipv6Mdt returns a bool
func (obj *bgpCapability) HasIpv6Mdt() bool {
	return obj.obj.Ipv6Mdt != nil
}

// Support for IPv6 MDT address family messages.
// SetIpv6Mdt sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv6Mdt(value bool) BgpCapability {

	obj.obj.Ipv6Mdt = &value
	return obj
}

// Support for the IPv6 Multicast VPN address family.
// Ipv6MulticastMplsVpn returns a bool
func (obj *bgpCapability) Ipv6MulticastMplsVpn() bool {

	return *obj.obj.Ipv6MulticastMplsVpn

}

// Support for the IPv6 Multicast VPN address family.
// Ipv6MulticastMplsVpn returns a bool
func (obj *bgpCapability) HasIpv6MulticastMplsVpn() bool {
	return obj.obj.Ipv6MulticastMplsVpn != nil
}

// Support for the IPv6 Multicast VPN address family.
// SetIpv6MulticastMplsVpn sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv6MulticastMplsVpn(value bool) BgpCapability {

	obj.obj.Ipv6MulticastMplsVpn = &value
	return obj
}

// Support for propagation of IPv6 unicast flow specification rules.
// Ipv6UnicastFlowSpec returns a bool
func (obj *bgpCapability) Ipv6UnicastFlowSpec() bool {

	return *obj.obj.Ipv6UnicastFlowSpec

}

// Support for propagation of IPv6 unicast flow specification rules.
// Ipv6UnicastFlowSpec returns a bool
func (obj *bgpCapability) HasIpv6UnicastFlowSpec() bool {
	return obj.obj.Ipv6UnicastFlowSpec != nil
}

// Support for propagation of IPv6 unicast flow specification rules.
// SetIpv6UnicastFlowSpec sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv6UnicastFlowSpec(value bool) BgpCapability {

	obj.obj.Ipv6UnicastFlowSpec = &value
	return obj
}

// Support for IPv6 SRTE policy.
// Ipv6SrTePolicy returns a bool
func (obj *bgpCapability) Ipv6SrTePolicy() bool {

	return *obj.obj.Ipv6SrTePolicy

}

// Support for IPv6 SRTE policy.
// Ipv6SrTePolicy returns a bool
func (obj *bgpCapability) HasIpv6SrTePolicy() bool {
	return obj.obj.Ipv6SrTePolicy != nil
}

// Support for IPv6 SRTE policy.
// SetIpv6SrTePolicy sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv6SrTePolicy(value bool) BgpCapability {

	obj.obj.Ipv6SrTePolicy = &value
	return obj
}

// Support for IPv6 Unicast Add Path Capability.
// Ipv6UnicastAddPath returns a bool
func (obj *bgpCapability) Ipv6UnicastAddPath() bool {

	return *obj.obj.Ipv6UnicastAddPath

}

// Support for IPv6 Unicast Add Path Capability.
// Ipv6UnicastAddPath returns a bool
func (obj *bgpCapability) HasIpv6UnicastAddPath() bool {
	return obj.obj.Ipv6UnicastAddPath != nil
}

// Support for IPv6 Unicast Add Path Capability.
// SetIpv6UnicastAddPath sets the bool value in the BgpCapability object
func (obj *bgpCapability) SetIpv6UnicastAddPath(value bool) BgpCapability {

	obj.obj.Ipv6UnicastAddPath = &value
	return obj
}

func (obj *bgpCapability) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpCapability) setDefault() {
	if obj.obj.Ipv4Unicast == nil {
		obj.SetIpv4Unicast(true)
	}
	if obj.obj.Ipv4Multicast == nil {
		obj.SetIpv4Multicast(false)
	}
	if obj.obj.Ipv6Unicast == nil {
		obj.SetIpv6Unicast(true)
	}
	if obj.obj.Ipv6Multicast == nil {
		obj.SetIpv6Multicast(false)
	}
	if obj.obj.Vpls == nil {
		obj.SetVpls(false)
	}
	if obj.obj.RouteRefresh == nil {
		obj.SetRouteRefresh(true)
	}
	if obj.obj.RouteConstraint == nil {
		obj.SetRouteConstraint(false)
	}
	if obj.obj.LinkStateNonVpn == nil {
		obj.SetLinkStateNonVpn(false)
	}
	if obj.obj.LinkStateVpn == nil {
		obj.SetLinkStateVpn(false)
	}
	if obj.obj.Evpn == nil {
		obj.SetEvpn(false)
	}
	if obj.obj.ExtendedNextHopEncoding == nil {
		obj.SetExtendedNextHopEncoding(false)
	}
	if obj.obj.Ipv4MulticastVpn == nil {
		obj.SetIpv4MulticastVpn(false)
	}
	if obj.obj.Ipv4MplsVpn == nil {
		obj.SetIpv4MplsVpn(false)
	}
	if obj.obj.Ipv4Mdt == nil {
		obj.SetIpv4Mdt(false)
	}
	if obj.obj.Ipv4MulticastMplsVpn == nil {
		obj.SetIpv4MulticastMplsVpn(false)
	}
	if obj.obj.Ipv4UnicastFlowSpec == nil {
		obj.SetIpv4UnicastFlowSpec(false)
	}
	if obj.obj.Ipv4SrTePolicy == nil {
		obj.SetIpv4SrTePolicy(false)
	}
	if obj.obj.Ipv4UnicastAddPath == nil {
		obj.SetIpv4UnicastAddPath(false)
	}
	if obj.obj.Ipv6MulticastVpn == nil {
		obj.SetIpv6MulticastVpn(false)
	}
	if obj.obj.Ipv6MplsVpn == nil {
		obj.SetIpv6MplsVpn(false)
	}
	if obj.obj.Ipv6Mdt == nil {
		obj.SetIpv6Mdt(false)
	}
	if obj.obj.Ipv6MulticastMplsVpn == nil {
		obj.SetIpv6MulticastMplsVpn(false)
	}
	if obj.obj.Ipv6UnicastFlowSpec == nil {
		obj.SetIpv6UnicastFlowSpec(false)
	}
	if obj.obj.Ipv6SrTePolicy == nil {
		obj.SetIpv6SrTePolicy(false)
	}
	if obj.obj.Ipv6UnicastAddPath == nil {
		obj.SetIpv6UnicastAddPath(false)
	}

}
