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
	obj                 *otg.StatesRequest
	marshaller          marshalStatesRequest
	unMarshaller        unMarshalStatesRequest
	ipv4NeighborsHolder Neighborsv4StatesRequest
	ipv6NeighborsHolder Neighborsv6StatesRequest
	bgpPrefixesHolder   BgpPrefixStateRequest
	isisLspsHolder      IsisLspsStateRequest
	lldpNeighborsHolder LldpNeighborsStateRequest
	rsvpLspsHolder      RsvpLspsStateRequest
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
	setNil()
}

type StatesRequestChoiceEnum string

// Enum of Choice on StatesRequest
var StatesRequestChoice = struct {
	IPV4_NEIGHBORS StatesRequestChoiceEnum
	IPV6_NEIGHBORS StatesRequestChoiceEnum
	BGP_PREFIXES   StatesRequestChoiceEnum
	ISIS_LSPS      StatesRequestChoiceEnum
	LLDP_NEIGHBORS StatesRequestChoiceEnum
	RSVP_LSPS      StatesRequestChoiceEnum
}{
	IPV4_NEIGHBORS: StatesRequestChoiceEnum("ipv4_neighbors"),
	IPV6_NEIGHBORS: StatesRequestChoiceEnum("ipv6_neighbors"),
	BGP_PREFIXES:   StatesRequestChoiceEnum("bgp_prefixes"),
	ISIS_LSPS:      StatesRequestChoiceEnum("isis_lsps"),
	LLDP_NEIGHBORS: StatesRequestChoiceEnum("lldp_neighbors"),
	RSVP_LSPS:      StatesRequestChoiceEnum("rsvp_lsps"),
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

}

func (obj *statesRequest) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(StatesRequestChoice.IPV4_NEIGHBORS)

	}

}