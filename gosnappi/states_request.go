package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StatesRequest *****
type statesRequest struct {
	validation
	obj                    *otg.StatesRequest
	marshaller             marshalStatesRequest
	unMarshaller           unMarshalStatesRequest
	ipv4NeighborsHolder    Neighborsv4StatesRequest
	ipv6NeighborsHolder    Neighborsv6StatesRequest
	bgpPrefixesHolder      BgpPrefixStateRequest
	isisLspsHolder         IsisLspsStateRequest
	lldpNeighborsHolder    LldpNeighborsStateRequest
	rsvpLspsHolder         RsvpLspsStateRequest
	dhcpv4InterfacesHolder Dhcpv4InterfaceStateRequest
	dhcpv4LeasesHolder     Dhcpv4LeaseStateRequest
	dhcpv6InterfacesHolder Dhcpv6InterfaceStateRequest
	dhcpv6LeasesHolder     Dhcpv6LeaseStateRequest
	ospfv2LsasHolder       Ospfv2LsasStateRequest
	ospfv3LsasHolder       Ospfv3LsasStateRequest
}

func NewStatesRequest() StatesRequest {
	obj := statesRequest{obj: &otg.StatesRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *statesRequest) msg() *otg.StatesRequest {
	return obj.obj
}

func (obj *statesRequest) setMsg(msg *otg.StatesRequest) StatesRequest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstatesRequest struct {
	obj *statesRequest
}

type marshalStatesRequest interface {
	// ToProto marshals StatesRequest to protobuf object *otg.StatesRequest
	ToProto() (*otg.StatesRequest, error)
	// ToPbText marshals StatesRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StatesRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals StatesRequest to JSON text
	ToJson() (string, error)
}

type unMarshalstatesRequest struct {
	obj *statesRequest
}

type unMarshalStatesRequest interface {
	// FromProto unmarshals StatesRequest from protobuf object *otg.StatesRequest
	FromProto(msg *otg.StatesRequest) (StatesRequest, error)
	// FromPbText unmarshals StatesRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StatesRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StatesRequest from JSON text
	FromJson(value string) error
}

func (obj *statesRequest) Marshal() marshalStatesRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstatesRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *statesRequest) Unmarshal() unMarshalStatesRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstatesRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstatesRequest) ToProto() (*otg.StatesRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstatesRequest) FromProto(msg *otg.StatesRequest) (StatesRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstatesRequest) ToPbText() (string, error) {
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

func (m *unMarshalstatesRequest) FromPbText(value string) error {
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

func (m *marshalstatesRequest) ToYaml() (string, error) {
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

func (m *unMarshalstatesRequest) FromYaml(value string) error {
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

func (m *marshalstatesRequest) ToJson() (string, error) {
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

func (m *unMarshalstatesRequest) FromJson(value string) error {
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

func (obj *statesRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *statesRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *statesRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *statesRequest) Clone() (StatesRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStatesRequest()
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

func (obj *statesRequest) setNil() {
	obj.ipv4NeighborsHolder = nil
	obj.ipv6NeighborsHolder = nil
	obj.bgpPrefixesHolder = nil
	obj.isisLspsHolder = nil
	obj.lldpNeighborsHolder = nil
	obj.rsvpLspsHolder = nil
	obj.dhcpv4InterfacesHolder = nil
	obj.dhcpv4LeasesHolder = nil
	obj.dhcpv6InterfacesHolder = nil
	obj.dhcpv6LeasesHolder = nil
	obj.ospfv2LsasHolder = nil
	obj.ospfv3LsasHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StatesRequest is request to traffic generator for states of choice
type StatesRequest interface {
	Validation
	// msg marshals StatesRequest to protobuf object *otg.StatesRequest
	// and doesn't set defaults
	msg() *otg.StatesRequest
	// setMsg unmarshals StatesRequest from protobuf object *otg.StatesRequest
	// and doesn't set defaults
	setMsg(*otg.StatesRequest) StatesRequest
	// provides marshal interface
	Marshal() marshalStatesRequest
	// provides unmarshal interface
	Unmarshal() unMarshalStatesRequest
	// validate validates StatesRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StatesRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StatesRequestChoiceEnum, set in StatesRequest
	Choice() StatesRequestChoiceEnum
	// setChoice assigns StatesRequestChoiceEnum provided by user to StatesRequest
	setChoice(value StatesRequestChoiceEnum) StatesRequest
	// HasChoice checks if Choice has been set in StatesRequest
	HasChoice() bool
	// Ipv4Neighbors returns Neighborsv4StatesRequest, set in StatesRequest.
	// Neighborsv4StatesRequest is the request to retrieve IPv4 Neighbor state (ARP cache entries) of a network interface(s).
	Ipv4Neighbors() Neighborsv4StatesRequest
	// SetIpv4Neighbors assigns Neighborsv4StatesRequest provided by user to StatesRequest.
	// Neighborsv4StatesRequest is the request to retrieve IPv4 Neighbor state (ARP cache entries) of a network interface(s).
	SetIpv4Neighbors(value Neighborsv4StatesRequest) StatesRequest
	// HasIpv4Neighbors checks if Ipv4Neighbors has been set in StatesRequest
	HasIpv4Neighbors() bool
	// Ipv6Neighbors returns Neighborsv6StatesRequest, set in StatesRequest.
	// Neighborsv6StatesRequest is the request to retrieve IPv6 Neighbor state (NDISC cache entries) of a network interface(s).
	Ipv6Neighbors() Neighborsv6StatesRequest
	// SetIpv6Neighbors assigns Neighborsv6StatesRequest provided by user to StatesRequest.
	// Neighborsv6StatesRequest is the request to retrieve IPv6 Neighbor state (NDISC cache entries) of a network interface(s).
	SetIpv6Neighbors(value Neighborsv6StatesRequest) StatesRequest
	// HasIpv6Neighbors checks if Ipv6Neighbors has been set in StatesRequest
	HasIpv6Neighbors() bool
	// BgpPrefixes returns BgpPrefixStateRequest, set in StatesRequest.
	// BgpPrefixStateRequest is the request to retrieve BGP peer prefix information.
	BgpPrefixes() BgpPrefixStateRequest
	// SetBgpPrefixes assigns BgpPrefixStateRequest provided by user to StatesRequest.
	// BgpPrefixStateRequest is the request to retrieve BGP peer prefix information.
	SetBgpPrefixes(value BgpPrefixStateRequest) StatesRequest
	// HasBgpPrefixes checks if BgpPrefixes has been set in StatesRequest
	HasBgpPrefixes() bool
	// IsisLsps returns IsisLspsStateRequest, set in StatesRequest.
	// IsisLspsStateRequest is the request to retrieve ISIS Link State PDU (LSP) information learned by the router.
	IsisLsps() IsisLspsStateRequest
	// SetIsisLsps assigns IsisLspsStateRequest provided by user to StatesRequest.
	// IsisLspsStateRequest is the request to retrieve ISIS Link State PDU (LSP) information learned by the router.
	SetIsisLsps(value IsisLspsStateRequest) StatesRequest
	// HasIsisLsps checks if IsisLsps has been set in StatesRequest
	HasIsisLsps() bool
	// LldpNeighbors returns LldpNeighborsStateRequest, set in StatesRequest.
	// LldpNeighborsStateRequest is the request to retrieve LLDP neighbor information for a given instance.
	LldpNeighbors() LldpNeighborsStateRequest
	// SetLldpNeighbors assigns LldpNeighborsStateRequest provided by user to StatesRequest.
	// LldpNeighborsStateRequest is the request to retrieve LLDP neighbor information for a given instance.
	SetLldpNeighbors(value LldpNeighborsStateRequest) StatesRequest
	// HasLldpNeighbors checks if LldpNeighbors has been set in StatesRequest
	HasLldpNeighbors() bool
	// RsvpLsps returns RsvpLspsStateRequest, set in StatesRequest.
	// RsvpLspsStateRequest is the request to retrieve RSVP Label Switched Path (LSP) information learned by the router.
	RsvpLsps() RsvpLspsStateRequest
	// SetRsvpLsps assigns RsvpLspsStateRequest provided by user to StatesRequest.
	// RsvpLspsStateRequest is the request to retrieve RSVP Label Switched Path (LSP) information learned by the router.
	SetRsvpLsps(value RsvpLspsStateRequest) StatesRequest
	// HasRsvpLsps checks if RsvpLsps has been set in StatesRequest
	HasRsvpLsps() bool
	// Dhcpv4Interfaces returns Dhcpv4InterfaceStateRequest, set in StatesRequest.
	// Dhcpv4InterfaceStateRequest is the request for assigned IPv4 address information associated with DHCP Client sessions.
	Dhcpv4Interfaces() Dhcpv4InterfaceStateRequest
	// SetDhcpv4Interfaces assigns Dhcpv4InterfaceStateRequest provided by user to StatesRequest.
	// Dhcpv4InterfaceStateRequest is the request for assigned IPv4 address information associated with DHCP Client sessions.
	SetDhcpv4Interfaces(value Dhcpv4InterfaceStateRequest) StatesRequest
	// HasDhcpv4Interfaces checks if Dhcpv4Interfaces has been set in StatesRequest
	HasDhcpv4Interfaces() bool
	// Dhcpv4Leases returns Dhcpv4LeaseStateRequest, set in StatesRequest.
	// Dhcpv4LeaseStateRequest is the request to retrieve DHCP Server host allocated status.
	Dhcpv4Leases() Dhcpv4LeaseStateRequest
	// SetDhcpv4Leases assigns Dhcpv4LeaseStateRequest provided by user to StatesRequest.
	// Dhcpv4LeaseStateRequest is the request to retrieve DHCP Server host allocated status.
	SetDhcpv4Leases(value Dhcpv4LeaseStateRequest) StatesRequest
	// HasDhcpv4Leases checks if Dhcpv4Leases has been set in StatesRequest
	HasDhcpv4Leases() bool
	// Dhcpv6Interfaces returns Dhcpv6InterfaceStateRequest, set in StatesRequest.
	// Dhcpv6InterfaceStateRequest is the request for assigned IPv6 address information associated with DHCP Client sessions.
	Dhcpv6Interfaces() Dhcpv6InterfaceStateRequest
	// SetDhcpv6Interfaces assigns Dhcpv6InterfaceStateRequest provided by user to StatesRequest.
	// Dhcpv6InterfaceStateRequest is the request for assigned IPv6 address information associated with DHCP Client sessions.
	SetDhcpv6Interfaces(value Dhcpv6InterfaceStateRequest) StatesRequest
	// HasDhcpv6Interfaces checks if Dhcpv6Interfaces has been set in StatesRequest
	HasDhcpv6Interfaces() bool
	// Dhcpv6Leases returns Dhcpv6LeaseStateRequest, set in StatesRequest.
	// Dhcpv6LeaseStateRequest is the request to retrieve DHCP Server host allocated status.
	Dhcpv6Leases() Dhcpv6LeaseStateRequest
	// SetDhcpv6Leases assigns Dhcpv6LeaseStateRequest provided by user to StatesRequest.
	// Dhcpv6LeaseStateRequest is the request to retrieve DHCP Server host allocated status.
	SetDhcpv6Leases(value Dhcpv6LeaseStateRequest) StatesRequest
	// HasDhcpv6Leases checks if Dhcpv6Leases has been set in StatesRequest
	HasDhcpv6Leases() bool
	// Ospfv2Lsas returns Ospfv2LsasStateRequest, set in StatesRequest.
	// Ospfv2LsasStateRequest is the request to retrieve OSPFv2 Link State Advertisements (LSA) information learned by the routers.
	Ospfv2Lsas() Ospfv2LsasStateRequest
	// SetOspfv2Lsas assigns Ospfv2LsasStateRequest provided by user to StatesRequest.
	// Ospfv2LsasStateRequest is the request to retrieve OSPFv2 Link State Advertisements (LSA) information learned by the routers.
	SetOspfv2Lsas(value Ospfv2LsasStateRequest) StatesRequest
	// HasOspfv2Lsas checks if Ospfv2Lsas has been set in StatesRequest
	HasOspfv2Lsas() bool
	// Ospfv3Lsas returns Ospfv3LsasStateRequest, set in StatesRequest.
	// Ospfv3LsasStateRequest is the request to retrieve OSPFv3 Link State Advertisements (LSA) information learned by the routers.
	Ospfv3Lsas() Ospfv3LsasStateRequest
	// SetOspfv3Lsas assigns Ospfv3LsasStateRequest provided by user to StatesRequest.
	// Ospfv3LsasStateRequest is the request to retrieve OSPFv3 Link State Advertisements (LSA) information learned by the routers.
	SetOspfv3Lsas(value Ospfv3LsasStateRequest) StatesRequest
	// HasOspfv3Lsas checks if Ospfv3Lsas has been set in StatesRequest
	HasOspfv3Lsas() bool
	setNil()
}

type StatesRequestChoiceEnum string

// Enum of Choice on StatesRequest
var StatesRequestChoice = struct {
	IPV4_NEIGHBORS    StatesRequestChoiceEnum
	IPV6_NEIGHBORS    StatesRequestChoiceEnum
	BGP_PREFIXES      StatesRequestChoiceEnum
	ISIS_LSPS         StatesRequestChoiceEnum
	LLDP_NEIGHBORS    StatesRequestChoiceEnum
	RSVP_LSPS         StatesRequestChoiceEnum
	DHCPV4_INTERFACES StatesRequestChoiceEnum
	DHCPV4_LEASES     StatesRequestChoiceEnum
	DHCPV6_INTERFACES StatesRequestChoiceEnum
	DHCPV6_LEASES     StatesRequestChoiceEnum
	OSPFV2_LSAS       StatesRequestChoiceEnum
	OSPFV3_LSAS       StatesRequestChoiceEnum
}{
	IPV4_NEIGHBORS:    StatesRequestChoiceEnum("ipv4_neighbors"),
	IPV6_NEIGHBORS:    StatesRequestChoiceEnum("ipv6_neighbors"),
	BGP_PREFIXES:      StatesRequestChoiceEnum("bgp_prefixes"),
	ISIS_LSPS:         StatesRequestChoiceEnum("isis_lsps"),
	LLDP_NEIGHBORS:    StatesRequestChoiceEnum("lldp_neighbors"),
	RSVP_LSPS:         StatesRequestChoiceEnum("rsvp_lsps"),
	DHCPV4_INTERFACES: StatesRequestChoiceEnum("dhcpv4_interfaces"),
	DHCPV4_LEASES:     StatesRequestChoiceEnum("dhcpv4_leases"),
	DHCPV6_INTERFACES: StatesRequestChoiceEnum("dhcpv6_interfaces"),
	DHCPV6_LEASES:     StatesRequestChoiceEnum("dhcpv6_leases"),
	OSPFV2_LSAS:       StatesRequestChoiceEnum("ospfv2_lsas"),
	OSPFV3_LSAS:       StatesRequestChoiceEnum("ospfv3_lsas"),
}

func (obj *statesRequest) Choice() StatesRequestChoiceEnum {
	return StatesRequestChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *statesRequest) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *statesRequest) setChoice(value StatesRequestChoiceEnum) StatesRequest {
	intValue, ok := otg.StatesRequest_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StatesRequestChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StatesRequest_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ospfv3Lsas = nil
	obj.ospfv3LsasHolder = nil
	obj.obj.Ospfv2Lsas = nil
	obj.ospfv2LsasHolder = nil
	obj.obj.Dhcpv6Leases = nil
	obj.dhcpv6LeasesHolder = nil
	obj.obj.Dhcpv6Interfaces = nil
	obj.dhcpv6InterfacesHolder = nil
	obj.obj.Dhcpv4Leases = nil
	obj.dhcpv4LeasesHolder = nil
	obj.obj.Dhcpv4Interfaces = nil
	obj.dhcpv4InterfacesHolder = nil
	obj.obj.RsvpLsps = nil
	obj.rsvpLspsHolder = nil
	obj.obj.LldpNeighbors = nil
	obj.lldpNeighborsHolder = nil
	obj.obj.IsisLsps = nil
	obj.isisLspsHolder = nil
	obj.obj.BgpPrefixes = nil
	obj.bgpPrefixesHolder = nil
	obj.obj.Ipv6Neighbors = nil
	obj.ipv6NeighborsHolder = nil
	obj.obj.Ipv4Neighbors = nil
	obj.ipv4NeighborsHolder = nil

	if value == StatesRequestChoice.IPV4_NEIGHBORS {
		obj.obj.Ipv4Neighbors = NewNeighborsv4StatesRequest().msg()
	}

	if value == StatesRequestChoice.IPV6_NEIGHBORS {
		obj.obj.Ipv6Neighbors = NewNeighborsv6StatesRequest().msg()
	}

	if value == StatesRequestChoice.BGP_PREFIXES {
		obj.obj.BgpPrefixes = NewBgpPrefixStateRequest().msg()
	}

	if value == StatesRequestChoice.ISIS_LSPS {
		obj.obj.IsisLsps = NewIsisLspsStateRequest().msg()
	}

	if value == StatesRequestChoice.LLDP_NEIGHBORS {
		obj.obj.LldpNeighbors = NewLldpNeighborsStateRequest().msg()
	}

	if value == StatesRequestChoice.RSVP_LSPS {
		obj.obj.RsvpLsps = NewRsvpLspsStateRequest().msg()
	}

	if value == StatesRequestChoice.DHCPV4_INTERFACES {
		obj.obj.Dhcpv4Interfaces = NewDhcpv4InterfaceStateRequest().msg()
	}

	if value == StatesRequestChoice.DHCPV4_LEASES {
		obj.obj.Dhcpv4Leases = NewDhcpv4LeaseStateRequest().msg()
	}

	if value == StatesRequestChoice.DHCPV6_INTERFACES {
		obj.obj.Dhcpv6Interfaces = NewDhcpv6InterfaceStateRequest().msg()
	}

	if value == StatesRequestChoice.DHCPV6_LEASES {
		obj.obj.Dhcpv6Leases = NewDhcpv6LeaseStateRequest().msg()
	}

	if value == StatesRequestChoice.OSPFV2_LSAS {
		obj.obj.Ospfv2Lsas = NewOspfv2LsasStateRequest().msg()
	}

	if value == StatesRequestChoice.OSPFV3_LSAS {
		obj.obj.Ospfv3Lsas = NewOspfv3LsasStateRequest().msg()
	}

	return obj
}

// description is TBD
// Ipv4Neighbors returns a Neighborsv4StatesRequest
func (obj *statesRequest) Ipv4Neighbors() Neighborsv4StatesRequest {
	if obj.obj.Ipv4Neighbors == nil {
		obj.setChoice(StatesRequestChoice.IPV4_NEIGHBORS)
	}
	if obj.ipv4NeighborsHolder == nil {
		obj.ipv4NeighborsHolder = &neighborsv4StatesRequest{obj: obj.obj.Ipv4Neighbors}
	}
	return obj.ipv4NeighborsHolder
}

// description is TBD
// Ipv4Neighbors returns a Neighborsv4StatesRequest
func (obj *statesRequest) HasIpv4Neighbors() bool {
	return obj.obj.Ipv4Neighbors != nil
}

// description is TBD
// SetIpv4Neighbors sets the Neighborsv4StatesRequest value in the StatesRequest object
func (obj *statesRequest) SetIpv4Neighbors(value Neighborsv4StatesRequest) StatesRequest {
	obj.setChoice(StatesRequestChoice.IPV4_NEIGHBORS)
	obj.ipv4NeighborsHolder = nil
	obj.obj.Ipv4Neighbors = value.msg()

	return obj
}

// description is TBD
// Ipv6Neighbors returns a Neighborsv6StatesRequest
func (obj *statesRequest) Ipv6Neighbors() Neighborsv6StatesRequest {
	if obj.obj.Ipv6Neighbors == nil {
		obj.setChoice(StatesRequestChoice.IPV6_NEIGHBORS)
	}
	if obj.ipv6NeighborsHolder == nil {
		obj.ipv6NeighborsHolder = &neighborsv6StatesRequest{obj: obj.obj.Ipv6Neighbors}
	}
	return obj.ipv6NeighborsHolder
}

// description is TBD
// Ipv6Neighbors returns a Neighborsv6StatesRequest
func (obj *statesRequest) HasIpv6Neighbors() bool {
	return obj.obj.Ipv6Neighbors != nil
}

// description is TBD
// SetIpv6Neighbors sets the Neighborsv6StatesRequest value in the StatesRequest object
func (obj *statesRequest) SetIpv6Neighbors(value Neighborsv6StatesRequest) StatesRequest {
	obj.setChoice(StatesRequestChoice.IPV6_NEIGHBORS)
	obj.ipv6NeighborsHolder = nil
	obj.obj.Ipv6Neighbors = value.msg()

	return obj
}

// description is TBD
// BgpPrefixes returns a BgpPrefixStateRequest
func (obj *statesRequest) BgpPrefixes() BgpPrefixStateRequest {
	if obj.obj.BgpPrefixes == nil {
		obj.setChoice(StatesRequestChoice.BGP_PREFIXES)
	}
	if obj.bgpPrefixesHolder == nil {
		obj.bgpPrefixesHolder = &bgpPrefixStateRequest{obj: obj.obj.BgpPrefixes}
	}
	return obj.bgpPrefixesHolder
}

// description is TBD
// BgpPrefixes returns a BgpPrefixStateRequest
func (obj *statesRequest) HasBgpPrefixes() bool {
	return obj.obj.BgpPrefixes != nil
}

// description is TBD
// SetBgpPrefixes sets the BgpPrefixStateRequest value in the StatesRequest object
func (obj *statesRequest) SetBgpPrefixes(value BgpPrefixStateRequest) StatesRequest {
	obj.setChoice(StatesRequestChoice.BGP_PREFIXES)
	obj.bgpPrefixesHolder = nil
	obj.obj.BgpPrefixes = value.msg()

	return obj
}

// description is TBD
// IsisLsps returns a IsisLspsStateRequest
func (obj *statesRequest) IsisLsps() IsisLspsStateRequest {
	if obj.obj.IsisLsps == nil {
		obj.setChoice(StatesRequestChoice.ISIS_LSPS)
	}
	if obj.isisLspsHolder == nil {
		obj.isisLspsHolder = &isisLspsStateRequest{obj: obj.obj.IsisLsps}
	}
	return obj.isisLspsHolder
}

// description is TBD
// IsisLsps returns a IsisLspsStateRequest
func (obj *statesRequest) HasIsisLsps() bool {
	return obj.obj.IsisLsps != nil
}

// description is TBD
// SetIsisLsps sets the IsisLspsStateRequest value in the StatesRequest object
func (obj *statesRequest) SetIsisLsps(value IsisLspsStateRequest) StatesRequest {
	obj.setChoice(StatesRequestChoice.ISIS_LSPS)
	obj.isisLspsHolder = nil
	obj.obj.IsisLsps = value.msg()

	return obj
}

// description is TBD
// LldpNeighbors returns a LldpNeighborsStateRequest
func (obj *statesRequest) LldpNeighbors() LldpNeighborsStateRequest {
	if obj.obj.LldpNeighbors == nil {
		obj.setChoice(StatesRequestChoice.LLDP_NEIGHBORS)
	}
	if obj.lldpNeighborsHolder == nil {
		obj.lldpNeighborsHolder = &lldpNeighborsStateRequest{obj: obj.obj.LldpNeighbors}
	}
	return obj.lldpNeighborsHolder
}

// description is TBD
// LldpNeighbors returns a LldpNeighborsStateRequest
func (obj *statesRequest) HasLldpNeighbors() bool {
	return obj.obj.LldpNeighbors != nil
}

// description is TBD
// SetLldpNeighbors sets the LldpNeighborsStateRequest value in the StatesRequest object
func (obj *statesRequest) SetLldpNeighbors(value LldpNeighborsStateRequest) StatesRequest {
	obj.setChoice(StatesRequestChoice.LLDP_NEIGHBORS)
	obj.lldpNeighborsHolder = nil
	obj.obj.LldpNeighbors = value.msg()

	return obj
}

// description is TBD
// RsvpLsps returns a RsvpLspsStateRequest
func (obj *statesRequest) RsvpLsps() RsvpLspsStateRequest {
	if obj.obj.RsvpLsps == nil {
		obj.setChoice(StatesRequestChoice.RSVP_LSPS)
	}
	if obj.rsvpLspsHolder == nil {
		obj.rsvpLspsHolder = &rsvpLspsStateRequest{obj: obj.obj.RsvpLsps}
	}
	return obj.rsvpLspsHolder
}

// description is TBD
// RsvpLsps returns a RsvpLspsStateRequest
func (obj *statesRequest) HasRsvpLsps() bool {
	return obj.obj.RsvpLsps != nil
}

// description is TBD
// SetRsvpLsps sets the RsvpLspsStateRequest value in the StatesRequest object
func (obj *statesRequest) SetRsvpLsps(value RsvpLspsStateRequest) StatesRequest {
	obj.setChoice(StatesRequestChoice.RSVP_LSPS)
	obj.rsvpLspsHolder = nil
	obj.obj.RsvpLsps = value.msg()

	return obj
}

// description is TBD
// Dhcpv4Interfaces returns a Dhcpv4InterfaceStateRequest
func (obj *statesRequest) Dhcpv4Interfaces() Dhcpv4InterfaceStateRequest {
	if obj.obj.Dhcpv4Interfaces == nil {
		obj.setChoice(StatesRequestChoice.DHCPV4_INTERFACES)
	}
	if obj.dhcpv4InterfacesHolder == nil {
		obj.dhcpv4InterfacesHolder = &dhcpv4InterfaceStateRequest{obj: obj.obj.Dhcpv4Interfaces}
	}
	return obj.dhcpv4InterfacesHolder
}

// description is TBD
// Dhcpv4Interfaces returns a Dhcpv4InterfaceStateRequest
func (obj *statesRequest) HasDhcpv4Interfaces() bool {
	return obj.obj.Dhcpv4Interfaces != nil
}

// description is TBD
// SetDhcpv4Interfaces sets the Dhcpv4InterfaceStateRequest value in the StatesRequest object
func (obj *statesRequest) SetDhcpv4Interfaces(value Dhcpv4InterfaceStateRequest) StatesRequest {
	obj.setChoice(StatesRequestChoice.DHCPV4_INTERFACES)
	obj.dhcpv4InterfacesHolder = nil
	obj.obj.Dhcpv4Interfaces = value.msg()

	return obj
}

// description is TBD
// Dhcpv4Leases returns a Dhcpv4LeaseStateRequest
func (obj *statesRequest) Dhcpv4Leases() Dhcpv4LeaseStateRequest {
	if obj.obj.Dhcpv4Leases == nil {
		obj.setChoice(StatesRequestChoice.DHCPV4_LEASES)
	}
	if obj.dhcpv4LeasesHolder == nil {
		obj.dhcpv4LeasesHolder = &dhcpv4LeaseStateRequest{obj: obj.obj.Dhcpv4Leases}
	}
	return obj.dhcpv4LeasesHolder
}

// description is TBD
// Dhcpv4Leases returns a Dhcpv4LeaseStateRequest
func (obj *statesRequest) HasDhcpv4Leases() bool {
	return obj.obj.Dhcpv4Leases != nil
}

// description is TBD
// SetDhcpv4Leases sets the Dhcpv4LeaseStateRequest value in the StatesRequest object
func (obj *statesRequest) SetDhcpv4Leases(value Dhcpv4LeaseStateRequest) StatesRequest {
	obj.setChoice(StatesRequestChoice.DHCPV4_LEASES)
	obj.dhcpv4LeasesHolder = nil
	obj.obj.Dhcpv4Leases = value.msg()

	return obj
}

// description is TBD
// Dhcpv6Interfaces returns a Dhcpv6InterfaceStateRequest
func (obj *statesRequest) Dhcpv6Interfaces() Dhcpv6InterfaceStateRequest {
	if obj.obj.Dhcpv6Interfaces == nil {
		obj.setChoice(StatesRequestChoice.DHCPV6_INTERFACES)
	}
	if obj.dhcpv6InterfacesHolder == nil {
		obj.dhcpv6InterfacesHolder = &dhcpv6InterfaceStateRequest{obj: obj.obj.Dhcpv6Interfaces}
	}
	return obj.dhcpv6InterfacesHolder
}

// description is TBD
// Dhcpv6Interfaces returns a Dhcpv6InterfaceStateRequest
func (obj *statesRequest) HasDhcpv6Interfaces() bool {
	return obj.obj.Dhcpv6Interfaces != nil
}

// description is TBD
// SetDhcpv6Interfaces sets the Dhcpv6InterfaceStateRequest value in the StatesRequest object
func (obj *statesRequest) SetDhcpv6Interfaces(value Dhcpv6InterfaceStateRequest) StatesRequest {
	obj.setChoice(StatesRequestChoice.DHCPV6_INTERFACES)
	obj.dhcpv6InterfacesHolder = nil
	obj.obj.Dhcpv6Interfaces = value.msg()

	return obj
}

// description is TBD
// Dhcpv6Leases returns a Dhcpv6LeaseStateRequest
func (obj *statesRequest) Dhcpv6Leases() Dhcpv6LeaseStateRequest {
	if obj.obj.Dhcpv6Leases == nil {
		obj.setChoice(StatesRequestChoice.DHCPV6_LEASES)
	}
	if obj.dhcpv6LeasesHolder == nil {
		obj.dhcpv6LeasesHolder = &dhcpv6LeaseStateRequest{obj: obj.obj.Dhcpv6Leases}
	}
	return obj.dhcpv6LeasesHolder
}

// description is TBD
// Dhcpv6Leases returns a Dhcpv6LeaseStateRequest
func (obj *statesRequest) HasDhcpv6Leases() bool {
	return obj.obj.Dhcpv6Leases != nil
}

// description is TBD
// SetDhcpv6Leases sets the Dhcpv6LeaseStateRequest value in the StatesRequest object
func (obj *statesRequest) SetDhcpv6Leases(value Dhcpv6LeaseStateRequest) StatesRequest {
	obj.setChoice(StatesRequestChoice.DHCPV6_LEASES)
	obj.dhcpv6LeasesHolder = nil
	obj.obj.Dhcpv6Leases = value.msg()

	return obj
}

// description is TBD
// Ospfv2Lsas returns a Ospfv2LsasStateRequest
func (obj *statesRequest) Ospfv2Lsas() Ospfv2LsasStateRequest {
	if obj.obj.Ospfv2Lsas == nil {
		obj.setChoice(StatesRequestChoice.OSPFV2_LSAS)
	}
	if obj.ospfv2LsasHolder == nil {
		obj.ospfv2LsasHolder = &ospfv2LsasStateRequest{obj: obj.obj.Ospfv2Lsas}
	}
	return obj.ospfv2LsasHolder
}

// description is TBD
// Ospfv2Lsas returns a Ospfv2LsasStateRequest
func (obj *statesRequest) HasOspfv2Lsas() bool {
	return obj.obj.Ospfv2Lsas != nil
}

// description is TBD
// SetOspfv2Lsas sets the Ospfv2LsasStateRequest value in the StatesRequest object
func (obj *statesRequest) SetOspfv2Lsas(value Ospfv2LsasStateRequest) StatesRequest {
	obj.setChoice(StatesRequestChoice.OSPFV2_LSAS)
	obj.ospfv2LsasHolder = nil
	obj.obj.Ospfv2Lsas = value.msg()

	return obj
}

// description is TBD
// Ospfv3Lsas returns a Ospfv3LsasStateRequest
func (obj *statesRequest) Ospfv3Lsas() Ospfv3LsasStateRequest {
	if obj.obj.Ospfv3Lsas == nil {
		obj.setChoice(StatesRequestChoice.OSPFV3_LSAS)
	}
	if obj.ospfv3LsasHolder == nil {
		obj.ospfv3LsasHolder = &ospfv3LsasStateRequest{obj: obj.obj.Ospfv3Lsas}
	}
	return obj.ospfv3LsasHolder
}

// description is TBD
// Ospfv3Lsas returns a Ospfv3LsasStateRequest
func (obj *statesRequest) HasOspfv3Lsas() bool {
	return obj.obj.Ospfv3Lsas != nil
}

// description is TBD
// SetOspfv3Lsas sets the Ospfv3LsasStateRequest value in the StatesRequest object
func (obj *statesRequest) SetOspfv3Lsas(value Ospfv3LsasStateRequest) StatesRequest {
	obj.setChoice(StatesRequestChoice.OSPFV3_LSAS)
	obj.ospfv3LsasHolder = nil
	obj.obj.Ospfv3Lsas = value.msg()

	return obj
}

func (obj *statesRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4Neighbors != nil {

		obj.Ipv4Neighbors().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv6Neighbors != nil {

		obj.Ipv6Neighbors().validateObj(vObj, set_default)
	}

	if obj.obj.BgpPrefixes != nil {

		obj.BgpPrefixes().validateObj(vObj, set_default)
	}

	if obj.obj.IsisLsps != nil {

		obj.IsisLsps().validateObj(vObj, set_default)
	}

	if obj.obj.LldpNeighbors != nil {

		obj.LldpNeighbors().validateObj(vObj, set_default)
	}

	if obj.obj.RsvpLsps != nil {

		obj.RsvpLsps().validateObj(vObj, set_default)
	}

	if obj.obj.Dhcpv4Interfaces != nil {

		obj.Dhcpv4Interfaces().validateObj(vObj, set_default)
	}

	if obj.obj.Dhcpv4Leases != nil {

		obj.Dhcpv4Leases().validateObj(vObj, set_default)
	}

	if obj.obj.Dhcpv6Interfaces != nil {

		obj.Dhcpv6Interfaces().validateObj(vObj, set_default)
	}

	if obj.obj.Dhcpv6Leases != nil {

		obj.Dhcpv6Leases().validateObj(vObj, set_default)
	}

	if obj.obj.Ospfv2Lsas != nil {

		obj.Ospfv2Lsas().validateObj(vObj, set_default)
	}

	if obj.obj.Ospfv3Lsas != nil {

		obj.Ospfv3Lsas().validateObj(vObj, set_default)
	}

}

func (obj *statesRequest) setDefault() {
	var choices_set int = 0
	var choice StatesRequestChoiceEnum

	if obj.obj.Ipv4Neighbors != nil {
		choices_set += 1
		choice = StatesRequestChoice.IPV4_NEIGHBORS
	}

	if obj.obj.Ipv6Neighbors != nil {
		choices_set += 1
		choice = StatesRequestChoice.IPV6_NEIGHBORS
	}

	if obj.obj.BgpPrefixes != nil {
		choices_set += 1
		choice = StatesRequestChoice.BGP_PREFIXES
	}

	if obj.obj.IsisLsps != nil {
		choices_set += 1
		choice = StatesRequestChoice.ISIS_LSPS
	}

	if obj.obj.LldpNeighbors != nil {
		choices_set += 1
		choice = StatesRequestChoice.LLDP_NEIGHBORS
	}

	if obj.obj.RsvpLsps != nil {
		choices_set += 1
		choice = StatesRequestChoice.RSVP_LSPS
	}

	if obj.obj.Dhcpv4Interfaces != nil {
		choices_set += 1
		choice = StatesRequestChoice.DHCPV4_INTERFACES
	}

	if obj.obj.Dhcpv4Leases != nil {
		choices_set += 1
		choice = StatesRequestChoice.DHCPV4_LEASES
	}

	if obj.obj.Dhcpv6Interfaces != nil {
		choices_set += 1
		choice = StatesRequestChoice.DHCPV6_INTERFACES
	}

	if obj.obj.Dhcpv6Leases != nil {
		choices_set += 1
		choice = StatesRequestChoice.DHCPV6_LEASES
	}

	if obj.obj.Ospfv2Lsas != nil {
		choices_set += 1
		choice = StatesRequestChoice.OSPFV2_LSAS
	}

	if obj.obj.Ospfv3Lsas != nil {
		choices_set += 1
		choice = StatesRequestChoice.OSPFV3_LSAS
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(StatesRequestChoice.IPV4_NEIGHBORS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in StatesRequest")
			}
		} else {
			intVal := otg.StatesRequest_Choice_Enum_value[string(choice)]
			enumValue := otg.StatesRequest_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
