package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MetricsRequest *****
type metricsRequest struct {
	validation
	obj                      *otg.MetricsRequest
	marshaller               marshalMetricsRequest
	unMarshaller             unMarshalMetricsRequest
	portHolder               PortMetricsRequest
	flowHolder               FlowMetricsRequest
	bgpv4Holder              Bgpv4MetricsRequest
	bgpv6Holder              Bgpv6MetricsRequest
	isisHolder               IsisMetricsRequest
	lagHolder                LagMetricsRequest
	lacpHolder               LacpMetricsRequest
	lldpHolder               LldpMetricsRequest
	rsvpHolder               RsvpMetricsRequest
	dhcpv4ClientHolder       Dhcpv4ClientMetricsRequest
	dhcpv4ServerHolder       Dhcpv4ServerMetricsRequest
	dhcpv6ClientHolder       Dhcpv6ClientMetricsRequest
	dhcpv6ServerHolder       Dhcpv6ServerMetricsRequest
	ospfv2Holder             Ospfv2MetricsRequest
	convergenceHolder        ConvergenceRequest
	macsecHolder             MacsecMetricsRequest
	mkaHolder                MkaMetricsRequest
	ospfv3Holder             Ospfv3MetricsRequest
	rocev2Ipv4Holder         Rocev2IPv4MetricsRequest
	rocev2Ipv6Holder         Rocev2IPv6MetricsRequest
	rocev2FlowHolder         Rocev2FlowMetricsRequest
	egressOnlyTrackingHolder EgressOnlyTrackingMetricsRequest
}

func NewMetricsRequest() MetricsRequest {
	obj := metricsRequest{obj: &otg.MetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *metricsRequest) msg() *otg.MetricsRequest {
	return obj.obj
}

func (obj *metricsRequest) setMsg(msg *otg.MetricsRequest) MetricsRequest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmetricsRequest struct {
	obj *metricsRequest
}

type marshalMetricsRequest interface {
	// ToProto marshals MetricsRequest to protobuf object *otg.MetricsRequest
	ToProto() (*otg.MetricsRequest, error)
	// ToPbText marshals MetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals MetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalmetricsRequest struct {
	obj *metricsRequest
}

type unMarshalMetricsRequest interface {
	// FromProto unmarshals MetricsRequest from protobuf object *otg.MetricsRequest
	FromProto(msg *otg.MetricsRequest) (MetricsRequest, error)
	// FromPbText unmarshals MetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *metricsRequest) Marshal() marshalMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *metricsRequest) Unmarshal() unMarshalMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmetricsRequest) ToProto() (*otg.MetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmetricsRequest) FromProto(msg *otg.MetricsRequest) (MetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalmetricsRequest) FromPbText(value string) error {
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

func (m *marshalmetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalmetricsRequest) FromYaml(value string) error {
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

func (m *marshalmetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalmetricsRequest) FromJson(value string) error {
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

func (obj *metricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *metricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *metricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *metricsRequest) Clone() (MetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMetricsRequest()
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

func (obj *metricsRequest) setNil() {
	obj.portHolder = nil
	obj.flowHolder = nil
	obj.bgpv4Holder = nil
	obj.bgpv6Holder = nil
	obj.isisHolder = nil
	obj.lagHolder = nil
	obj.lacpHolder = nil
	obj.lldpHolder = nil
	obj.rsvpHolder = nil
	obj.dhcpv4ClientHolder = nil
	obj.dhcpv4ServerHolder = nil
	obj.dhcpv6ClientHolder = nil
	obj.dhcpv6ServerHolder = nil
	obj.ospfv2Holder = nil
	obj.convergenceHolder = nil
	obj.macsecHolder = nil
	obj.mkaHolder = nil
	obj.ospfv3Holder = nil
	obj.rocev2Ipv4Holder = nil
	obj.rocev2Ipv6Holder = nil
	obj.rocev2FlowHolder = nil
	obj.egressOnlyTrackingHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MetricsRequest is request to traffic generator for metrics of choice.
type MetricsRequest interface {
	Validation
	// msg marshals MetricsRequest to protobuf object *otg.MetricsRequest
	// and doesn't set defaults
	msg() *otg.MetricsRequest
	// setMsg unmarshals MetricsRequest from protobuf object *otg.MetricsRequest
	// and doesn't set defaults
	setMsg(*otg.MetricsRequest) MetricsRequest
	// provides marshal interface
	Marshal() marshalMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalMetricsRequest
	// validate validates MetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MetricsRequestChoiceEnum, set in MetricsRequest
	Choice() MetricsRequestChoiceEnum
	// setChoice assigns MetricsRequestChoiceEnum provided by user to MetricsRequest
	setChoice(value MetricsRequestChoiceEnum) MetricsRequest
	// HasChoice checks if Choice has been set in MetricsRequest
	HasChoice() bool
	// Port returns PortMetricsRequest, set in MetricsRequest.
	// PortMetricsRequest is the port result request to the traffic generator
	Port() PortMetricsRequest
	// SetPort assigns PortMetricsRequest provided by user to MetricsRequest.
	// PortMetricsRequest is the port result request to the traffic generator
	SetPort(value PortMetricsRequest) MetricsRequest
	// HasPort checks if Port has been set in MetricsRequest
	HasPort() bool
	// Flow returns FlowMetricsRequest, set in MetricsRequest.
	// FlowMetricsRequest is the container for a flow metric request.
	Flow() FlowMetricsRequest
	// SetFlow assigns FlowMetricsRequest provided by user to MetricsRequest.
	// FlowMetricsRequest is the container for a flow metric request.
	SetFlow(value FlowMetricsRequest) MetricsRequest
	// HasFlow checks if Flow has been set in MetricsRequest
	HasFlow() bool
	// Bgpv4 returns Bgpv4MetricsRequest, set in MetricsRequest.
	// Bgpv4MetricsRequest is the request to retrieve BGPv4 per peer metrics/statistics.
	Bgpv4() Bgpv4MetricsRequest
	// SetBgpv4 assigns Bgpv4MetricsRequest provided by user to MetricsRequest.
	// Bgpv4MetricsRequest is the request to retrieve BGPv4 per peer metrics/statistics.
	SetBgpv4(value Bgpv4MetricsRequest) MetricsRequest
	// HasBgpv4 checks if Bgpv4 has been set in MetricsRequest
	HasBgpv4() bool
	// Bgpv6 returns Bgpv6MetricsRequest, set in MetricsRequest.
	// Bgpv6MetricsRequest is the request to retrieve BGPv6 per peer metrics/statistics.
	Bgpv6() Bgpv6MetricsRequest
	// SetBgpv6 assigns Bgpv6MetricsRequest provided by user to MetricsRequest.
	// Bgpv6MetricsRequest is the request to retrieve BGPv6 per peer metrics/statistics.
	SetBgpv6(value Bgpv6MetricsRequest) MetricsRequest
	// HasBgpv6 checks if Bgpv6 has been set in MetricsRequest
	HasBgpv6() bool
	// Isis returns IsisMetricsRequest, set in MetricsRequest.
	// IsisMetricsRequest is the request to retrieve ISIS per Router metrics/statistics.
	Isis() IsisMetricsRequest
	// SetIsis assigns IsisMetricsRequest provided by user to MetricsRequest.
	// IsisMetricsRequest is the request to retrieve ISIS per Router metrics/statistics.
	SetIsis(value IsisMetricsRequest) MetricsRequest
	// HasIsis checks if Isis has been set in MetricsRequest
	HasIsis() bool
	// Lag returns LagMetricsRequest, set in MetricsRequest.
	// LagMetricsRequest is the request to retrieve per LAG metrics/statistics.
	Lag() LagMetricsRequest
	// SetLag assigns LagMetricsRequest provided by user to MetricsRequest.
	// LagMetricsRequest is the request to retrieve per LAG metrics/statistics.
	SetLag(value LagMetricsRequest) MetricsRequest
	// HasLag checks if Lag has been set in MetricsRequest
	HasLag() bool
	// Lacp returns LacpMetricsRequest, set in MetricsRequest.
	// LacpMetricsRequest is the request to retrieve LACP per LAG member metrics/statistics.
	Lacp() LacpMetricsRequest
	// SetLacp assigns LacpMetricsRequest provided by user to MetricsRequest.
	// LacpMetricsRequest is the request to retrieve LACP per LAG member metrics/statistics.
	SetLacp(value LacpMetricsRequest) MetricsRequest
	// HasLacp checks if Lacp has been set in MetricsRequest
	HasLacp() bool
	// Lldp returns LldpMetricsRequest, set in MetricsRequest.
	// LldpMetricsRequest is the request to retrieve LLDP per instance metrics/statistics.
	Lldp() LldpMetricsRequest
	// SetLldp assigns LldpMetricsRequest provided by user to MetricsRequest.
	// LldpMetricsRequest is the request to retrieve LLDP per instance metrics/statistics.
	SetLldp(value LldpMetricsRequest) MetricsRequest
	// HasLldp checks if Lldp has been set in MetricsRequest
	HasLldp() bool
	// Rsvp returns RsvpMetricsRequest, set in MetricsRequest.
	// RsvpMetricsRequest is the request to retrieve RSVP-TE per Router metrics/statistics.
	Rsvp() RsvpMetricsRequest
	// SetRsvp assigns RsvpMetricsRequest provided by user to MetricsRequest.
	// RsvpMetricsRequest is the request to retrieve RSVP-TE per Router metrics/statistics.
	SetRsvp(value RsvpMetricsRequest) MetricsRequest
	// HasRsvp checks if Rsvp has been set in MetricsRequest
	HasRsvp() bool
	// Dhcpv4Client returns Dhcpv4ClientMetricsRequest, set in MetricsRequest.
	// Dhcpv4ClientMetricsRequest is the request to retrieve DHCPv4 per client metrics/statistics.
	Dhcpv4Client() Dhcpv4ClientMetricsRequest
	// SetDhcpv4Client assigns Dhcpv4ClientMetricsRequest provided by user to MetricsRequest.
	// Dhcpv4ClientMetricsRequest is the request to retrieve DHCPv4 per client metrics/statistics.
	SetDhcpv4Client(value Dhcpv4ClientMetricsRequest) MetricsRequest
	// HasDhcpv4Client checks if Dhcpv4Client has been set in MetricsRequest
	HasDhcpv4Client() bool
	// Dhcpv4Server returns Dhcpv4ServerMetricsRequest, set in MetricsRequest.
	// Dhcpv4ServerMetricsRequest is the request to retrieve DHCPv4 per Server metrics/statistics.
	Dhcpv4Server() Dhcpv4ServerMetricsRequest
	// SetDhcpv4Server assigns Dhcpv4ServerMetricsRequest provided by user to MetricsRequest.
	// Dhcpv4ServerMetricsRequest is the request to retrieve DHCPv4 per Server metrics/statistics.
	SetDhcpv4Server(value Dhcpv4ServerMetricsRequest) MetricsRequest
	// HasDhcpv4Server checks if Dhcpv4Server has been set in MetricsRequest
	HasDhcpv4Server() bool
	// Dhcpv6Client returns Dhcpv6ClientMetricsRequest, set in MetricsRequest.
	// Dhcpv6ClientMetricsRequest is the request to retrieve DHCPv6 per client metrics/statistics.
	Dhcpv6Client() Dhcpv6ClientMetricsRequest
	// SetDhcpv6Client assigns Dhcpv6ClientMetricsRequest provided by user to MetricsRequest.
	// Dhcpv6ClientMetricsRequest is the request to retrieve DHCPv6 per client metrics/statistics.
	SetDhcpv6Client(value Dhcpv6ClientMetricsRequest) MetricsRequest
	// HasDhcpv6Client checks if Dhcpv6Client has been set in MetricsRequest
	HasDhcpv6Client() bool
	// Dhcpv6Server returns Dhcpv6ServerMetricsRequest, set in MetricsRequest.
	// Dhcpv6ServerMetricsRequest is the request to retrieve DHCPv6 per Server metrics/statistics.
	Dhcpv6Server() Dhcpv6ServerMetricsRequest
	// SetDhcpv6Server assigns Dhcpv6ServerMetricsRequest provided by user to MetricsRequest.
	// Dhcpv6ServerMetricsRequest is the request to retrieve DHCPv6 per Server metrics/statistics.
	SetDhcpv6Server(value Dhcpv6ServerMetricsRequest) MetricsRequest
	// HasDhcpv6Server checks if Dhcpv6Server has been set in MetricsRequest
	HasDhcpv6Server() bool
	// Ospfv2 returns Ospfv2MetricsRequest, set in MetricsRequest.
	// Ospfv2MetricsRequest is the request to retrieve OSPFv2 per Router metrics/statistics.
	Ospfv2() Ospfv2MetricsRequest
	// SetOspfv2 assigns Ospfv2MetricsRequest provided by user to MetricsRequest.
	// Ospfv2MetricsRequest is the request to retrieve OSPFv2 per Router metrics/statistics.
	SetOspfv2(value Ospfv2MetricsRequest) MetricsRequest
	// HasOspfv2 checks if Ospfv2 has been set in MetricsRequest
	HasOspfv2() bool
	// Convergence returns ConvergenceRequest, set in MetricsRequest.
	// ConvergenceRequest is under Review: Convergence metrics is currently under review for pending exploration on use cases.
	//
	// Under Review: Convergence metrics is currently under review for pending exploration on use cases.
	//
	// Container for requesting control-plane and data-plane convergence time metrics for flows.
	Convergence() ConvergenceRequest
	// SetConvergence assigns ConvergenceRequest provided by user to MetricsRequest.
	// ConvergenceRequest is under Review: Convergence metrics is currently under review for pending exploration on use cases.
	//
	// Under Review: Convergence metrics is currently under review for pending exploration on use cases.
	//
	// Container for requesting control-plane and data-plane convergence time metrics for flows.
	SetConvergence(value ConvergenceRequest) MetricsRequest
	// HasConvergence checks if Convergence has been set in MetricsRequest
	HasConvergence() bool
	// Macsec returns MacsecMetricsRequest, set in MetricsRequest.
	// MacsecMetricsRequest is the request to retrieve MACsec per secure entity(secY) metrics/statistics.
	Macsec() MacsecMetricsRequest
	// SetMacsec assigns MacsecMetricsRequest provided by user to MetricsRequest.
	// MacsecMetricsRequest is the request to retrieve MACsec per secure entity(secY) metrics/statistics.
	SetMacsec(value MacsecMetricsRequest) MetricsRequest
	// HasMacsec checks if Macsec has been set in MetricsRequest
	HasMacsec() bool
	// Mka returns MkaMetricsRequest, set in MetricsRequest.
	// MkaMetricsRequest is the request to retrieve MKA per peer metrics/statistics.
	Mka() MkaMetricsRequest
	// SetMka assigns MkaMetricsRequest provided by user to MetricsRequest.
	// MkaMetricsRequest is the request to retrieve MKA per peer metrics/statistics.
	SetMka(value MkaMetricsRequest) MetricsRequest
	// HasMka checks if Mka has been set in MetricsRequest
	HasMka() bool
	// Ospfv3 returns Ospfv3MetricsRequest, set in MetricsRequest.
	// Ospfv3MetricsRequest is the request to retrieve OSPFv3 per router metrics/statistics.
	Ospfv3() Ospfv3MetricsRequest
	// SetOspfv3 assigns Ospfv3MetricsRequest provided by user to MetricsRequest.
	// Ospfv3MetricsRequest is the request to retrieve OSPFv3 per router metrics/statistics.
	SetOspfv3(value Ospfv3MetricsRequest) MetricsRequest
	// HasOspfv3 checks if Ospfv3 has been set in MetricsRequest
	HasOspfv3() bool
	// Rocev2Ipv4 returns Rocev2IPv4MetricsRequest, set in MetricsRequest.
	// Rocev2IPv4MetricsRequest is request to retrieve RoCEv2 over IPv4 per peer metrics/statistics.
	Rocev2Ipv4() Rocev2IPv4MetricsRequest
	// SetRocev2Ipv4 assigns Rocev2IPv4MetricsRequest provided by user to MetricsRequest.
	// Rocev2IPv4MetricsRequest is request to retrieve RoCEv2 over IPv4 per peer metrics/statistics.
	SetRocev2Ipv4(value Rocev2IPv4MetricsRequest) MetricsRequest
	// HasRocev2Ipv4 checks if Rocev2Ipv4 has been set in MetricsRequest
	HasRocev2Ipv4() bool
	// Rocev2Ipv6 returns Rocev2IPv6MetricsRequest, set in MetricsRequest.
	// Rocev2IPv6MetricsRequest is request to retrieve RoCEv2 over IPv6 per peer metrics/statistics.
	Rocev2Ipv6() Rocev2IPv6MetricsRequest
	// SetRocev2Ipv6 assigns Rocev2IPv6MetricsRequest provided by user to MetricsRequest.
	// Rocev2IPv6MetricsRequest is request to retrieve RoCEv2 over IPv6 per peer metrics/statistics.
	SetRocev2Ipv6(value Rocev2IPv6MetricsRequest) MetricsRequest
	// HasRocev2Ipv6 checks if Rocev2Ipv6 has been set in MetricsRequest
	HasRocev2Ipv6() bool
	// Rocev2Flow returns Rocev2FlowMetricsRequest, set in MetricsRequest.
	// Rocev2FlowMetricsRequest is request to retrieve RoCEv2 FLow statistics.
	Rocev2Flow() Rocev2FlowMetricsRequest
	// SetRocev2Flow assigns Rocev2FlowMetricsRequest provided by user to MetricsRequest.
	// Rocev2FlowMetricsRequest is request to retrieve RoCEv2 FLow statistics.
	SetRocev2Flow(value Rocev2FlowMetricsRequest) MetricsRequest
	// HasRocev2Flow checks if Rocev2Flow has been set in MetricsRequest
	HasRocev2Flow() bool
	// EgressOnlyTracking returns EgressOnlyTrackingMetricsRequest, set in MetricsRequest.
	// EgressOnlyTrackingMetricsRequest is the container for a egress only tracking metric request.
	EgressOnlyTracking() EgressOnlyTrackingMetricsRequest
	// SetEgressOnlyTracking assigns EgressOnlyTrackingMetricsRequest provided by user to MetricsRequest.
	// EgressOnlyTrackingMetricsRequest is the container for a egress only tracking metric request.
	SetEgressOnlyTracking(value EgressOnlyTrackingMetricsRequest) MetricsRequest
	// HasEgressOnlyTracking checks if EgressOnlyTracking has been set in MetricsRequest
	HasEgressOnlyTracking() bool
	setNil()
}

type MetricsRequestChoiceEnum string

// Enum of Choice on MetricsRequest
var MetricsRequestChoice = struct {
	PORT                 MetricsRequestChoiceEnum
	FLOW                 MetricsRequestChoiceEnum
	BGPV4                MetricsRequestChoiceEnum
	BGPV6                MetricsRequestChoiceEnum
	ISIS                 MetricsRequestChoiceEnum
	LAG                  MetricsRequestChoiceEnum
	LACP                 MetricsRequestChoiceEnum
	LLDP                 MetricsRequestChoiceEnum
	RSVP                 MetricsRequestChoiceEnum
	DHCPV4_CLIENT        MetricsRequestChoiceEnum
	DHCPV4_SERVER        MetricsRequestChoiceEnum
	DHCPV6_CLIENT        MetricsRequestChoiceEnum
	DHCPV6_SERVER        MetricsRequestChoiceEnum
	OSPFV2               MetricsRequestChoiceEnum
	CONVERGENCE          MetricsRequestChoiceEnum
	MACSEC               MetricsRequestChoiceEnum
	MKA                  MetricsRequestChoiceEnum
	OSPFV3               MetricsRequestChoiceEnum
	ROCEV2_IPV4          MetricsRequestChoiceEnum
	ROCEV2_IPV6          MetricsRequestChoiceEnum
	ROCEV2_FLOW          MetricsRequestChoiceEnum
	EGRESS_ONLY_TRACKING MetricsRequestChoiceEnum
}{
	PORT:                 MetricsRequestChoiceEnum("port"),
	FLOW:                 MetricsRequestChoiceEnum("flow"),
	BGPV4:                MetricsRequestChoiceEnum("bgpv4"),
	BGPV6:                MetricsRequestChoiceEnum("bgpv6"),
	ISIS:                 MetricsRequestChoiceEnum("isis"),
	LAG:                  MetricsRequestChoiceEnum("lag"),
	LACP:                 MetricsRequestChoiceEnum("lacp"),
	LLDP:                 MetricsRequestChoiceEnum("lldp"),
	RSVP:                 MetricsRequestChoiceEnum("rsvp"),
	DHCPV4_CLIENT:        MetricsRequestChoiceEnum("dhcpv4_client"),
	DHCPV4_SERVER:        MetricsRequestChoiceEnum("dhcpv4_server"),
	DHCPV6_CLIENT:        MetricsRequestChoiceEnum("dhcpv6_client"),
	DHCPV6_SERVER:        MetricsRequestChoiceEnum("dhcpv6_server"),
	OSPFV2:               MetricsRequestChoiceEnum("ospfv2"),
	CONVERGENCE:          MetricsRequestChoiceEnum("convergence"),
	MACSEC:               MetricsRequestChoiceEnum("macsec"),
	MKA:                  MetricsRequestChoiceEnum("mka"),
	OSPFV3:               MetricsRequestChoiceEnum("ospfv3"),
	ROCEV2_IPV4:          MetricsRequestChoiceEnum("rocev2_ipv4"),
	ROCEV2_IPV6:          MetricsRequestChoiceEnum("rocev2_ipv6"),
	ROCEV2_FLOW:          MetricsRequestChoiceEnum("rocev2_flow"),
	EGRESS_ONLY_TRACKING: MetricsRequestChoiceEnum("egress_only_tracking"),
}

func (obj *metricsRequest) Choice() MetricsRequestChoiceEnum {
	return MetricsRequestChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *metricsRequest) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *metricsRequest) setChoice(value MetricsRequestChoiceEnum) MetricsRequest {
	intValue, ok := otg.MetricsRequest_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MetricsRequestChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MetricsRequest_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.EgressOnlyTracking = nil
	obj.egressOnlyTrackingHolder = nil
	obj.obj.Rocev2Flow = nil
	obj.rocev2FlowHolder = nil
	obj.obj.Rocev2Ipv6 = nil
	obj.rocev2Ipv6Holder = nil
	obj.obj.Rocev2Ipv4 = nil
	obj.rocev2Ipv4Holder = nil
	obj.obj.Ospfv3 = nil
	obj.ospfv3Holder = nil
	obj.obj.Mka = nil
	obj.mkaHolder = nil
	obj.obj.Macsec = nil
	obj.macsecHolder = nil
	obj.obj.Convergence = nil
	obj.convergenceHolder = nil
	obj.obj.Ospfv2 = nil
	obj.ospfv2Holder = nil
	obj.obj.Dhcpv6Server = nil
	obj.dhcpv6ServerHolder = nil
	obj.obj.Dhcpv6Client = nil
	obj.dhcpv6ClientHolder = nil
	obj.obj.Dhcpv4Server = nil
	obj.dhcpv4ServerHolder = nil
	obj.obj.Dhcpv4Client = nil
	obj.dhcpv4ClientHolder = nil
	obj.obj.Rsvp = nil
	obj.rsvpHolder = nil
	obj.obj.Lldp = nil
	obj.lldpHolder = nil
	obj.obj.Lacp = nil
	obj.lacpHolder = nil
	obj.obj.Lag = nil
	obj.lagHolder = nil
	obj.obj.Isis = nil
	obj.isisHolder = nil
	obj.obj.Bgpv6 = nil
	obj.bgpv6Holder = nil
	obj.obj.Bgpv4 = nil
	obj.bgpv4Holder = nil
	obj.obj.Flow = nil
	obj.flowHolder = nil
	obj.obj.Port = nil
	obj.portHolder = nil

	if value == MetricsRequestChoice.PORT {
		obj.obj.Port = NewPortMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.FLOW {
		obj.obj.Flow = NewFlowMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.BGPV4 {
		obj.obj.Bgpv4 = NewBgpv4MetricsRequest().msg()
	}

	if value == MetricsRequestChoice.BGPV6 {
		obj.obj.Bgpv6 = NewBgpv6MetricsRequest().msg()
	}

	if value == MetricsRequestChoice.ISIS {
		obj.obj.Isis = NewIsisMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.LAG {
		obj.obj.Lag = NewLagMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.LACP {
		obj.obj.Lacp = NewLacpMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.LLDP {
		obj.obj.Lldp = NewLldpMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.RSVP {
		obj.obj.Rsvp = NewRsvpMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.DHCPV4_CLIENT {
		obj.obj.Dhcpv4Client = NewDhcpv4ClientMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.DHCPV4_SERVER {
		obj.obj.Dhcpv4Server = NewDhcpv4ServerMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.DHCPV6_CLIENT {
		obj.obj.Dhcpv6Client = NewDhcpv6ClientMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.DHCPV6_SERVER {
		obj.obj.Dhcpv6Server = NewDhcpv6ServerMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.OSPFV2 {
		obj.obj.Ospfv2 = NewOspfv2MetricsRequest().msg()
	}

	if value == MetricsRequestChoice.CONVERGENCE {
		obj.obj.Convergence = NewConvergenceRequest().msg()
	}

	if value == MetricsRequestChoice.MACSEC {
		obj.obj.Macsec = NewMacsecMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.MKA {
		obj.obj.Mka = NewMkaMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.OSPFV3 {
		obj.obj.Ospfv3 = NewOspfv3MetricsRequest().msg()
	}

	if value == MetricsRequestChoice.ROCEV2_IPV4 {
		obj.obj.Rocev2Ipv4 = NewRocev2IPv4MetricsRequest().msg()
	}

	if value == MetricsRequestChoice.ROCEV2_IPV6 {
		obj.obj.Rocev2Ipv6 = NewRocev2IPv6MetricsRequest().msg()
	}

	if value == MetricsRequestChoice.ROCEV2_FLOW {
		obj.obj.Rocev2Flow = NewRocev2FlowMetricsRequest().msg()
	}

	if value == MetricsRequestChoice.EGRESS_ONLY_TRACKING {
		obj.obj.EgressOnlyTracking = NewEgressOnlyTrackingMetricsRequest().msg()
	}

	return obj
}

// description is TBD
// Port returns a PortMetricsRequest
func (obj *metricsRequest) Port() PortMetricsRequest {
	if obj.obj.Port == nil {
		obj.setChoice(MetricsRequestChoice.PORT)
	}
	if obj.portHolder == nil {
		obj.portHolder = &portMetricsRequest{obj: obj.obj.Port}
	}
	return obj.portHolder
}

// description is TBD
// Port returns a PortMetricsRequest
func (obj *metricsRequest) HasPort() bool {
	return obj.obj.Port != nil
}

// description is TBD
// SetPort sets the PortMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetPort(value PortMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.PORT)
	obj.portHolder = nil
	obj.obj.Port = value.msg()

	return obj
}

// description is TBD
// Flow returns a FlowMetricsRequest
func (obj *metricsRequest) Flow() FlowMetricsRequest {
	if obj.obj.Flow == nil {
		obj.setChoice(MetricsRequestChoice.FLOW)
	}
	if obj.flowHolder == nil {
		obj.flowHolder = &flowMetricsRequest{obj: obj.obj.Flow}
	}
	return obj.flowHolder
}

// description is TBD
// Flow returns a FlowMetricsRequest
func (obj *metricsRequest) HasFlow() bool {
	return obj.obj.Flow != nil
}

// description is TBD
// SetFlow sets the FlowMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetFlow(value FlowMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.FLOW)
	obj.flowHolder = nil
	obj.obj.Flow = value.msg()

	return obj
}

// description is TBD
// Bgpv4 returns a Bgpv4MetricsRequest
func (obj *metricsRequest) Bgpv4() Bgpv4MetricsRequest {
	if obj.obj.Bgpv4 == nil {
		obj.setChoice(MetricsRequestChoice.BGPV4)
	}
	if obj.bgpv4Holder == nil {
		obj.bgpv4Holder = &bgpv4MetricsRequest{obj: obj.obj.Bgpv4}
	}
	return obj.bgpv4Holder
}

// description is TBD
// Bgpv4 returns a Bgpv4MetricsRequest
func (obj *metricsRequest) HasBgpv4() bool {
	return obj.obj.Bgpv4 != nil
}

// description is TBD
// SetBgpv4 sets the Bgpv4MetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetBgpv4(value Bgpv4MetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.BGPV4)
	obj.bgpv4Holder = nil
	obj.obj.Bgpv4 = value.msg()

	return obj
}

// description is TBD
// Bgpv6 returns a Bgpv6MetricsRequest
func (obj *metricsRequest) Bgpv6() Bgpv6MetricsRequest {
	if obj.obj.Bgpv6 == nil {
		obj.setChoice(MetricsRequestChoice.BGPV6)
	}
	if obj.bgpv6Holder == nil {
		obj.bgpv6Holder = &bgpv6MetricsRequest{obj: obj.obj.Bgpv6}
	}
	return obj.bgpv6Holder
}

// description is TBD
// Bgpv6 returns a Bgpv6MetricsRequest
func (obj *metricsRequest) HasBgpv6() bool {
	return obj.obj.Bgpv6 != nil
}

// description is TBD
// SetBgpv6 sets the Bgpv6MetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetBgpv6(value Bgpv6MetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.BGPV6)
	obj.bgpv6Holder = nil
	obj.obj.Bgpv6 = value.msg()

	return obj
}

// description is TBD
// Isis returns a IsisMetricsRequest
func (obj *metricsRequest) Isis() IsisMetricsRequest {
	if obj.obj.Isis == nil {
		obj.setChoice(MetricsRequestChoice.ISIS)
	}
	if obj.isisHolder == nil {
		obj.isisHolder = &isisMetricsRequest{obj: obj.obj.Isis}
	}
	return obj.isisHolder
}

// description is TBD
// Isis returns a IsisMetricsRequest
func (obj *metricsRequest) HasIsis() bool {
	return obj.obj.Isis != nil
}

// description is TBD
// SetIsis sets the IsisMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetIsis(value IsisMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.ISIS)
	obj.isisHolder = nil
	obj.obj.Isis = value.msg()

	return obj
}

// description is TBD
// Lag returns a LagMetricsRequest
func (obj *metricsRequest) Lag() LagMetricsRequest {
	if obj.obj.Lag == nil {
		obj.setChoice(MetricsRequestChoice.LAG)
	}
	if obj.lagHolder == nil {
		obj.lagHolder = &lagMetricsRequest{obj: obj.obj.Lag}
	}
	return obj.lagHolder
}

// description is TBD
// Lag returns a LagMetricsRequest
func (obj *metricsRequest) HasLag() bool {
	return obj.obj.Lag != nil
}

// description is TBD
// SetLag sets the LagMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetLag(value LagMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.LAG)
	obj.lagHolder = nil
	obj.obj.Lag = value.msg()

	return obj
}

// description is TBD
// Lacp returns a LacpMetricsRequest
func (obj *metricsRequest) Lacp() LacpMetricsRequest {
	if obj.obj.Lacp == nil {
		obj.setChoice(MetricsRequestChoice.LACP)
	}
	if obj.lacpHolder == nil {
		obj.lacpHolder = &lacpMetricsRequest{obj: obj.obj.Lacp}
	}
	return obj.lacpHolder
}

// description is TBD
// Lacp returns a LacpMetricsRequest
func (obj *metricsRequest) HasLacp() bool {
	return obj.obj.Lacp != nil
}

// description is TBD
// SetLacp sets the LacpMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetLacp(value LacpMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.LACP)
	obj.lacpHolder = nil
	obj.obj.Lacp = value.msg()

	return obj
}

// description is TBD
// Lldp returns a LldpMetricsRequest
func (obj *metricsRequest) Lldp() LldpMetricsRequest {
	if obj.obj.Lldp == nil {
		obj.setChoice(MetricsRequestChoice.LLDP)
	}
	if obj.lldpHolder == nil {
		obj.lldpHolder = &lldpMetricsRequest{obj: obj.obj.Lldp}
	}
	return obj.lldpHolder
}

// description is TBD
// Lldp returns a LldpMetricsRequest
func (obj *metricsRequest) HasLldp() bool {
	return obj.obj.Lldp != nil
}

// description is TBD
// SetLldp sets the LldpMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetLldp(value LldpMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.LLDP)
	obj.lldpHolder = nil
	obj.obj.Lldp = value.msg()

	return obj
}

// description is TBD
// Rsvp returns a RsvpMetricsRequest
func (obj *metricsRequest) Rsvp() RsvpMetricsRequest {
	if obj.obj.Rsvp == nil {
		obj.setChoice(MetricsRequestChoice.RSVP)
	}
	if obj.rsvpHolder == nil {
		obj.rsvpHolder = &rsvpMetricsRequest{obj: obj.obj.Rsvp}
	}
	return obj.rsvpHolder
}

// description is TBD
// Rsvp returns a RsvpMetricsRequest
func (obj *metricsRequest) HasRsvp() bool {
	return obj.obj.Rsvp != nil
}

// description is TBD
// SetRsvp sets the RsvpMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetRsvp(value RsvpMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.RSVP)
	obj.rsvpHolder = nil
	obj.obj.Rsvp = value.msg()

	return obj
}

// description is TBD
// Dhcpv4Client returns a Dhcpv4ClientMetricsRequest
func (obj *metricsRequest) Dhcpv4Client() Dhcpv4ClientMetricsRequest {
	if obj.obj.Dhcpv4Client == nil {
		obj.setChoice(MetricsRequestChoice.DHCPV4_CLIENT)
	}
	if obj.dhcpv4ClientHolder == nil {
		obj.dhcpv4ClientHolder = &dhcpv4ClientMetricsRequest{obj: obj.obj.Dhcpv4Client}
	}
	return obj.dhcpv4ClientHolder
}

// description is TBD
// Dhcpv4Client returns a Dhcpv4ClientMetricsRequest
func (obj *metricsRequest) HasDhcpv4Client() bool {
	return obj.obj.Dhcpv4Client != nil
}

// description is TBD
// SetDhcpv4Client sets the Dhcpv4ClientMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetDhcpv4Client(value Dhcpv4ClientMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.DHCPV4_CLIENT)
	obj.dhcpv4ClientHolder = nil
	obj.obj.Dhcpv4Client = value.msg()

	return obj
}

// description is TBD
// Dhcpv4Server returns a Dhcpv4ServerMetricsRequest
func (obj *metricsRequest) Dhcpv4Server() Dhcpv4ServerMetricsRequest {
	if obj.obj.Dhcpv4Server == nil {
		obj.setChoice(MetricsRequestChoice.DHCPV4_SERVER)
	}
	if obj.dhcpv4ServerHolder == nil {
		obj.dhcpv4ServerHolder = &dhcpv4ServerMetricsRequest{obj: obj.obj.Dhcpv4Server}
	}
	return obj.dhcpv4ServerHolder
}

// description is TBD
// Dhcpv4Server returns a Dhcpv4ServerMetricsRequest
func (obj *metricsRequest) HasDhcpv4Server() bool {
	return obj.obj.Dhcpv4Server != nil
}

// description is TBD
// SetDhcpv4Server sets the Dhcpv4ServerMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetDhcpv4Server(value Dhcpv4ServerMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.DHCPV4_SERVER)
	obj.dhcpv4ServerHolder = nil
	obj.obj.Dhcpv4Server = value.msg()

	return obj
}

// description is TBD
// Dhcpv6Client returns a Dhcpv6ClientMetricsRequest
func (obj *metricsRequest) Dhcpv6Client() Dhcpv6ClientMetricsRequest {
	if obj.obj.Dhcpv6Client == nil {
		obj.setChoice(MetricsRequestChoice.DHCPV6_CLIENT)
	}
	if obj.dhcpv6ClientHolder == nil {
		obj.dhcpv6ClientHolder = &dhcpv6ClientMetricsRequest{obj: obj.obj.Dhcpv6Client}
	}
	return obj.dhcpv6ClientHolder
}

// description is TBD
// Dhcpv6Client returns a Dhcpv6ClientMetricsRequest
func (obj *metricsRequest) HasDhcpv6Client() bool {
	return obj.obj.Dhcpv6Client != nil
}

// description is TBD
// SetDhcpv6Client sets the Dhcpv6ClientMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetDhcpv6Client(value Dhcpv6ClientMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.DHCPV6_CLIENT)
	obj.dhcpv6ClientHolder = nil
	obj.obj.Dhcpv6Client = value.msg()

	return obj
}

// description is TBD
// Dhcpv6Server returns a Dhcpv6ServerMetricsRequest
func (obj *metricsRequest) Dhcpv6Server() Dhcpv6ServerMetricsRequest {
	if obj.obj.Dhcpv6Server == nil {
		obj.setChoice(MetricsRequestChoice.DHCPV6_SERVER)
	}
	if obj.dhcpv6ServerHolder == nil {
		obj.dhcpv6ServerHolder = &dhcpv6ServerMetricsRequest{obj: obj.obj.Dhcpv6Server}
	}
	return obj.dhcpv6ServerHolder
}

// description is TBD
// Dhcpv6Server returns a Dhcpv6ServerMetricsRequest
func (obj *metricsRequest) HasDhcpv6Server() bool {
	return obj.obj.Dhcpv6Server != nil
}

// description is TBD
// SetDhcpv6Server sets the Dhcpv6ServerMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetDhcpv6Server(value Dhcpv6ServerMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.DHCPV6_SERVER)
	obj.dhcpv6ServerHolder = nil
	obj.obj.Dhcpv6Server = value.msg()

	return obj
}

// description is TBD
// Ospfv2 returns a Ospfv2MetricsRequest
func (obj *metricsRequest) Ospfv2() Ospfv2MetricsRequest {
	if obj.obj.Ospfv2 == nil {
		obj.setChoice(MetricsRequestChoice.OSPFV2)
	}
	if obj.ospfv2Holder == nil {
		obj.ospfv2Holder = &ospfv2MetricsRequest{obj: obj.obj.Ospfv2}
	}
	return obj.ospfv2Holder
}

// description is TBD
// Ospfv2 returns a Ospfv2MetricsRequest
func (obj *metricsRequest) HasOspfv2() bool {
	return obj.obj.Ospfv2 != nil
}

// description is TBD
// SetOspfv2 sets the Ospfv2MetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetOspfv2(value Ospfv2MetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.OSPFV2)
	obj.ospfv2Holder = nil
	obj.obj.Ospfv2 = value.msg()

	return obj
}

// description is TBD
// Convergence returns a ConvergenceRequest
func (obj *metricsRequest) Convergence() ConvergenceRequest {
	if obj.obj.Convergence == nil {
		obj.setChoice(MetricsRequestChoice.CONVERGENCE)
	}
	if obj.convergenceHolder == nil {
		obj.convergenceHolder = &convergenceRequest{obj: obj.obj.Convergence}
	}
	return obj.convergenceHolder
}

// description is TBD
// Convergence returns a ConvergenceRequest
func (obj *metricsRequest) HasConvergence() bool {
	return obj.obj.Convergence != nil
}

// description is TBD
// SetConvergence sets the ConvergenceRequest value in the MetricsRequest object
func (obj *metricsRequest) SetConvergence(value ConvergenceRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.CONVERGENCE)
	obj.convergenceHolder = nil
	obj.obj.Convergence = value.msg()

	return obj
}

// description is TBD
// Macsec returns a MacsecMetricsRequest
func (obj *metricsRequest) Macsec() MacsecMetricsRequest {
	if obj.obj.Macsec == nil {
		obj.setChoice(MetricsRequestChoice.MACSEC)
	}
	if obj.macsecHolder == nil {
		obj.macsecHolder = &macsecMetricsRequest{obj: obj.obj.Macsec}
	}
	return obj.macsecHolder
}

// description is TBD
// Macsec returns a MacsecMetricsRequest
func (obj *metricsRequest) HasMacsec() bool {
	return obj.obj.Macsec != nil
}

// description is TBD
// SetMacsec sets the MacsecMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetMacsec(value MacsecMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.MACSEC)
	obj.macsecHolder = nil
	obj.obj.Macsec = value.msg()

	return obj
}

// description is TBD
// Mka returns a MkaMetricsRequest
func (obj *metricsRequest) Mka() MkaMetricsRequest {
	if obj.obj.Mka == nil {
		obj.setChoice(MetricsRequestChoice.MKA)
	}
	if obj.mkaHolder == nil {
		obj.mkaHolder = &mkaMetricsRequest{obj: obj.obj.Mka}
	}
	return obj.mkaHolder
}

// description is TBD
// Mka returns a MkaMetricsRequest
func (obj *metricsRequest) HasMka() bool {
	return obj.obj.Mka != nil
}

// description is TBD
// SetMka sets the MkaMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetMka(value MkaMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.MKA)
	obj.mkaHolder = nil
	obj.obj.Mka = value.msg()

	return obj
}

// description is TBD
// Ospfv3 returns a Ospfv3MetricsRequest
func (obj *metricsRequest) Ospfv3() Ospfv3MetricsRequest {
	if obj.obj.Ospfv3 == nil {
		obj.setChoice(MetricsRequestChoice.OSPFV3)
	}
	if obj.ospfv3Holder == nil {
		obj.ospfv3Holder = &ospfv3MetricsRequest{obj: obj.obj.Ospfv3}
	}
	return obj.ospfv3Holder
}

// description is TBD
// Ospfv3 returns a Ospfv3MetricsRequest
func (obj *metricsRequest) HasOspfv3() bool {
	return obj.obj.Ospfv3 != nil
}

// description is TBD
// SetOspfv3 sets the Ospfv3MetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetOspfv3(value Ospfv3MetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.OSPFV3)
	obj.ospfv3Holder = nil
	obj.obj.Ospfv3 = value.msg()

	return obj
}

// description is TBD
// Rocev2Ipv4 returns a Rocev2IPv4MetricsRequest
func (obj *metricsRequest) Rocev2Ipv4() Rocev2IPv4MetricsRequest {
	if obj.obj.Rocev2Ipv4 == nil {
		obj.setChoice(MetricsRequestChoice.ROCEV2_IPV4)
	}
	if obj.rocev2Ipv4Holder == nil {
		obj.rocev2Ipv4Holder = &rocev2IPv4MetricsRequest{obj: obj.obj.Rocev2Ipv4}
	}
	return obj.rocev2Ipv4Holder
}

// description is TBD
// Rocev2Ipv4 returns a Rocev2IPv4MetricsRequest
func (obj *metricsRequest) HasRocev2Ipv4() bool {
	return obj.obj.Rocev2Ipv4 != nil
}

// description is TBD
// SetRocev2Ipv4 sets the Rocev2IPv4MetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetRocev2Ipv4(value Rocev2IPv4MetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.ROCEV2_IPV4)
	obj.rocev2Ipv4Holder = nil
	obj.obj.Rocev2Ipv4 = value.msg()

	return obj
}

// description is TBD
// Rocev2Ipv6 returns a Rocev2IPv6MetricsRequest
func (obj *metricsRequest) Rocev2Ipv6() Rocev2IPv6MetricsRequest {
	if obj.obj.Rocev2Ipv6 == nil {
		obj.setChoice(MetricsRequestChoice.ROCEV2_IPV6)
	}
	if obj.rocev2Ipv6Holder == nil {
		obj.rocev2Ipv6Holder = &rocev2IPv6MetricsRequest{obj: obj.obj.Rocev2Ipv6}
	}
	return obj.rocev2Ipv6Holder
}

// description is TBD
// Rocev2Ipv6 returns a Rocev2IPv6MetricsRequest
func (obj *metricsRequest) HasRocev2Ipv6() bool {
	return obj.obj.Rocev2Ipv6 != nil
}

// description is TBD
// SetRocev2Ipv6 sets the Rocev2IPv6MetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetRocev2Ipv6(value Rocev2IPv6MetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.ROCEV2_IPV6)
	obj.rocev2Ipv6Holder = nil
	obj.obj.Rocev2Ipv6 = value.msg()

	return obj
}

// description is TBD
// Rocev2Flow returns a Rocev2FlowMetricsRequest
func (obj *metricsRequest) Rocev2Flow() Rocev2FlowMetricsRequest {
	if obj.obj.Rocev2Flow == nil {
		obj.setChoice(MetricsRequestChoice.ROCEV2_FLOW)
	}
	if obj.rocev2FlowHolder == nil {
		obj.rocev2FlowHolder = &rocev2FlowMetricsRequest{obj: obj.obj.Rocev2Flow}
	}
	return obj.rocev2FlowHolder
}

// description is TBD
// Rocev2Flow returns a Rocev2FlowMetricsRequest
func (obj *metricsRequest) HasRocev2Flow() bool {
	return obj.obj.Rocev2Flow != nil
}

// description is TBD
// SetRocev2Flow sets the Rocev2FlowMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetRocev2Flow(value Rocev2FlowMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.ROCEV2_FLOW)
	obj.rocev2FlowHolder = nil
	obj.obj.Rocev2Flow = value.msg()

	return obj
}

// description is TBD
// EgressOnlyTracking returns a EgressOnlyTrackingMetricsRequest
func (obj *metricsRequest) EgressOnlyTracking() EgressOnlyTrackingMetricsRequest {
	if obj.obj.EgressOnlyTracking == nil {
		obj.setChoice(MetricsRequestChoice.EGRESS_ONLY_TRACKING)
	}
	if obj.egressOnlyTrackingHolder == nil {
		obj.egressOnlyTrackingHolder = &egressOnlyTrackingMetricsRequest{obj: obj.obj.EgressOnlyTracking}
	}
	return obj.egressOnlyTrackingHolder
}

// description is TBD
// EgressOnlyTracking returns a EgressOnlyTrackingMetricsRequest
func (obj *metricsRequest) HasEgressOnlyTracking() bool {
	return obj.obj.EgressOnlyTracking != nil
}

// description is TBD
// SetEgressOnlyTracking sets the EgressOnlyTrackingMetricsRequest value in the MetricsRequest object
func (obj *metricsRequest) SetEgressOnlyTracking(value EgressOnlyTrackingMetricsRequest) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.EGRESS_ONLY_TRACKING)
	obj.egressOnlyTrackingHolder = nil
	obj.obj.EgressOnlyTracking = value.msg()

	return obj
}

func (obj *metricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Port != nil {

		obj.Port().validateObj(vObj, set_default)
	}

	if obj.obj.Flow != nil {

		obj.Flow().validateObj(vObj, set_default)
	}

	if obj.obj.Bgpv4 != nil {

		obj.Bgpv4().validateObj(vObj, set_default)
	}

	if obj.obj.Bgpv6 != nil {

		obj.Bgpv6().validateObj(vObj, set_default)
	}

	if obj.obj.Isis != nil {

		obj.Isis().validateObj(vObj, set_default)
	}

	if obj.obj.Lag != nil {

		obj.Lag().validateObj(vObj, set_default)
	}

	if obj.obj.Lacp != nil {

		obj.Lacp().validateObj(vObj, set_default)
	}

	if obj.obj.Lldp != nil {

		obj.Lldp().validateObj(vObj, set_default)
	}

	if obj.obj.Rsvp != nil {

		obj.Rsvp().validateObj(vObj, set_default)
	}

	if obj.obj.Dhcpv4Client != nil {

		obj.Dhcpv4Client().validateObj(vObj, set_default)
	}

	if obj.obj.Dhcpv4Server != nil {

		obj.Dhcpv4Server().validateObj(vObj, set_default)
	}

	if obj.obj.Dhcpv6Client != nil {

		obj.Dhcpv6Client().validateObj(vObj, set_default)
	}

	if obj.obj.Dhcpv6Server != nil {

		obj.Dhcpv6Server().validateObj(vObj, set_default)
	}

	if obj.obj.Ospfv2 != nil {

		obj.Ospfv2().validateObj(vObj, set_default)
	}

	if obj.obj.Convergence != nil {

		obj.Convergence().validateObj(vObj, set_default)
	}

	if obj.obj.Macsec != nil {

		obj.Macsec().validateObj(vObj, set_default)
	}

	if obj.obj.Mka != nil {

		obj.Mka().validateObj(vObj, set_default)
	}

	if obj.obj.Ospfv3 != nil {

		obj.Ospfv3().validateObj(vObj, set_default)
	}

	if obj.obj.Rocev2Ipv4 != nil {

		obj.Rocev2Ipv4().validateObj(vObj, set_default)
	}

	if obj.obj.Rocev2Ipv6 != nil {

		obj.Rocev2Ipv6().validateObj(vObj, set_default)
	}

	if obj.obj.Rocev2Flow != nil {

		obj.Rocev2Flow().validateObj(vObj, set_default)
	}

	if obj.obj.EgressOnlyTracking != nil {

		obj.EgressOnlyTracking().validateObj(vObj, set_default)
	}

}

func (obj *metricsRequest) setDefault() {
	var choices_set int = 0
	var choice MetricsRequestChoiceEnum

	if obj.obj.Port != nil {
		choices_set += 1
		choice = MetricsRequestChoice.PORT
	}

	if obj.obj.Flow != nil {
		choices_set += 1
		choice = MetricsRequestChoice.FLOW
	}

	if obj.obj.Bgpv4 != nil {
		choices_set += 1
		choice = MetricsRequestChoice.BGPV4
	}

	if obj.obj.Bgpv6 != nil {
		choices_set += 1
		choice = MetricsRequestChoice.BGPV6
	}

	if obj.obj.Isis != nil {
		choices_set += 1
		choice = MetricsRequestChoice.ISIS
	}

	if obj.obj.Lag != nil {
		choices_set += 1
		choice = MetricsRequestChoice.LAG
	}

	if obj.obj.Lacp != nil {
		choices_set += 1
		choice = MetricsRequestChoice.LACP
	}

	if obj.obj.Lldp != nil {
		choices_set += 1
		choice = MetricsRequestChoice.LLDP
	}

	if obj.obj.Rsvp != nil {
		choices_set += 1
		choice = MetricsRequestChoice.RSVP
	}

	if obj.obj.Dhcpv4Client != nil {
		choices_set += 1
		choice = MetricsRequestChoice.DHCPV4_CLIENT
	}

	if obj.obj.Dhcpv4Server != nil {
		choices_set += 1
		choice = MetricsRequestChoice.DHCPV4_SERVER
	}

	if obj.obj.Dhcpv6Client != nil {
		choices_set += 1
		choice = MetricsRequestChoice.DHCPV6_CLIENT
	}

	if obj.obj.Dhcpv6Server != nil {
		choices_set += 1
		choice = MetricsRequestChoice.DHCPV6_SERVER
	}

	if obj.obj.Ospfv2 != nil {
		choices_set += 1
		choice = MetricsRequestChoice.OSPFV2
	}

	if obj.obj.Convergence != nil {
		choices_set += 1
		choice = MetricsRequestChoice.CONVERGENCE
	}

	if obj.obj.Macsec != nil {
		choices_set += 1
		choice = MetricsRequestChoice.MACSEC
	}

	if obj.obj.Mka != nil {
		choices_set += 1
		choice = MetricsRequestChoice.MKA
	}

	if obj.obj.Ospfv3 != nil {
		choices_set += 1
		choice = MetricsRequestChoice.OSPFV3
	}

	if obj.obj.Rocev2Ipv4 != nil {
		choices_set += 1
		choice = MetricsRequestChoice.ROCEV2_IPV4
	}

	if obj.obj.Rocev2Ipv6 != nil {
		choices_set += 1
		choice = MetricsRequestChoice.ROCEV2_IPV6
	}

	if obj.obj.Rocev2Flow != nil {
		choices_set += 1
		choice = MetricsRequestChoice.ROCEV2_FLOW
	}

	if obj.obj.EgressOnlyTracking != nil {
		choices_set += 1
		choice = MetricsRequestChoice.EGRESS_ONLY_TRACKING
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MetricsRequestChoice.PORT)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MetricsRequest")
			}
		} else {
			intVal := otg.MetricsRequest_Choice_Enum_value[string(choice)]
			enumValue := otg.MetricsRequest_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
