package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StatesResponse *****
type statesResponse struct {
	validation
	obj                    *otg.StatesResponse
	marshaller             marshalStatesResponse
	unMarshaller           unMarshalStatesResponse
	ipv4NeighborsHolder    StatesResponseNeighborsv4StateIter
	ipv6NeighborsHolder    StatesResponseNeighborsv6StateIter
	bgpPrefixesHolder      StatesResponseBgpPrefixesStateIter
	isisLspsHolder         StatesResponseIsisLspsStateIter
	lldpNeighborsHolder    StatesResponseLldpNeighborsStateIter
	rsvpLspsHolder         StatesResponseRsvpLspsStateIter
	dhcpv4InterfacesHolder StatesResponseDhcpv4InterfaceStateIter
	dhcpv4LeasesHolder     StatesResponseDhcpv4LeasesStateIter
	dhcpv6InterfacesHolder StatesResponseDhcpv6InterfaceStateIter
	dhcpv6LeasesHolder     StatesResponseDhcpv6LeasesStateIter
	ospfv2LsasHolder       StatesResponseOspfv2LsaStateIter
}

func NewStatesResponse() StatesResponse {
	obj := statesResponse{obj: &otg.StatesResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *statesResponse) msg() *otg.StatesResponse {
	return obj.obj
}

func (obj *statesResponse) setMsg(msg *otg.StatesResponse) StatesResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstatesResponse struct {
	obj *statesResponse
}

type marshalStatesResponse interface {
	// ToProto marshals StatesResponse to protobuf object *otg.StatesResponse
	ToProto() (*otg.StatesResponse, error)
	// ToPbText marshals StatesResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StatesResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals StatesResponse to JSON text
	ToJson() (string, error)
}

type unMarshalstatesResponse struct {
	obj *statesResponse
}

type unMarshalStatesResponse interface {
	// FromProto unmarshals StatesResponse from protobuf object *otg.StatesResponse
	FromProto(msg *otg.StatesResponse) (StatesResponse, error)
	// FromPbText unmarshals StatesResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StatesResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StatesResponse from JSON text
	FromJson(value string) error
}

func (obj *statesResponse) Marshal() marshalStatesResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstatesResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *statesResponse) Unmarshal() unMarshalStatesResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstatesResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstatesResponse) ToProto() (*otg.StatesResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstatesResponse) FromProto(msg *otg.StatesResponse) (StatesResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstatesResponse) ToPbText() (string, error) {
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

func (m *unMarshalstatesResponse) FromPbText(value string) error {
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

func (m *marshalstatesResponse) ToYaml() (string, error) {
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

func (m *unMarshalstatesResponse) FromYaml(value string) error {
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

func (m *marshalstatesResponse) ToJson() (string, error) {
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

func (m *unMarshalstatesResponse) FromJson(value string) error {
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

func (obj *statesResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *statesResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *statesResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *statesResponse) Clone() (StatesResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStatesResponse()
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

func (obj *statesResponse) setNil() {
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
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StatesResponse is response containing chosen traffic generator states
type StatesResponse interface {
	Validation
	// msg marshals StatesResponse to protobuf object *otg.StatesResponse
	// and doesn't set defaults
	msg() *otg.StatesResponse
	// setMsg unmarshals StatesResponse from protobuf object *otg.StatesResponse
	// and doesn't set defaults
	setMsg(*otg.StatesResponse) StatesResponse
	// provides marshal interface
	Marshal() marshalStatesResponse
	// provides unmarshal interface
	Unmarshal() unMarshalStatesResponse
	// validate validates StatesResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StatesResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StatesResponseChoiceEnum, set in StatesResponse
	Choice() StatesResponseChoiceEnum
	// setChoice assigns StatesResponseChoiceEnum provided by user to StatesResponse
	setChoice(value StatesResponseChoiceEnum) StatesResponse
	// HasChoice checks if Choice has been set in StatesResponse
	HasChoice() bool
	// Ipv4Neighbors returns StatesResponseNeighborsv4StateIterIter, set in StatesResponse
	Ipv4Neighbors() StatesResponseNeighborsv4StateIter
	// Ipv6Neighbors returns StatesResponseNeighborsv6StateIterIter, set in StatesResponse
	Ipv6Neighbors() StatesResponseNeighborsv6StateIter
	// BgpPrefixes returns StatesResponseBgpPrefixesStateIterIter, set in StatesResponse
	BgpPrefixes() StatesResponseBgpPrefixesStateIter
	// IsisLsps returns StatesResponseIsisLspsStateIterIter, set in StatesResponse
	IsisLsps() StatesResponseIsisLspsStateIter
	// LldpNeighbors returns StatesResponseLldpNeighborsStateIterIter, set in StatesResponse
	LldpNeighbors() StatesResponseLldpNeighborsStateIter
	// RsvpLsps returns StatesResponseRsvpLspsStateIterIter, set in StatesResponse
	RsvpLsps() StatesResponseRsvpLspsStateIter
	// Dhcpv4Interfaces returns StatesResponseDhcpv4InterfaceStateIterIter, set in StatesResponse
	Dhcpv4Interfaces() StatesResponseDhcpv4InterfaceStateIter
	// Dhcpv4Leases returns StatesResponseDhcpv4LeasesStateIterIter, set in StatesResponse
	Dhcpv4Leases() StatesResponseDhcpv4LeasesStateIter
	// Dhcpv6Interfaces returns StatesResponseDhcpv6InterfaceStateIterIter, set in StatesResponse
	Dhcpv6Interfaces() StatesResponseDhcpv6InterfaceStateIter
	// Dhcpv6Leases returns StatesResponseDhcpv6LeasesStateIterIter, set in StatesResponse
	Dhcpv6Leases() StatesResponseDhcpv6LeasesStateIter
	// Ospfv2Lsas returns StatesResponseOspfv2LsaStateIterIter, set in StatesResponse
	Ospfv2Lsas() StatesResponseOspfv2LsaStateIter
	setNil()
}

type StatesResponseChoiceEnum string

// Enum of Choice on StatesResponse
var StatesResponseChoice = struct {
	IPV4_NEIGHBORS    StatesResponseChoiceEnum
	IPV6_NEIGHBORS    StatesResponseChoiceEnum
	BGP_PREFIXES      StatesResponseChoiceEnum
	ISIS_LSPS         StatesResponseChoiceEnum
	LLDP_NEIGHBORS    StatesResponseChoiceEnum
	RSVP_LSPS         StatesResponseChoiceEnum
	DHCPV4_INTERFACES StatesResponseChoiceEnum
	DHCPV4_LEASES     StatesResponseChoiceEnum
	DHCPV6_INTERFACES StatesResponseChoiceEnum
	DHCPV6_LEASES     StatesResponseChoiceEnum
	OSPFV2_LSAS       StatesResponseChoiceEnum
}{
	IPV4_NEIGHBORS:    StatesResponseChoiceEnum("ipv4_neighbors"),
	IPV6_NEIGHBORS:    StatesResponseChoiceEnum("ipv6_neighbors"),
	BGP_PREFIXES:      StatesResponseChoiceEnum("bgp_prefixes"),
	ISIS_LSPS:         StatesResponseChoiceEnum("isis_lsps"),
	LLDP_NEIGHBORS:    StatesResponseChoiceEnum("lldp_neighbors"),
	RSVP_LSPS:         StatesResponseChoiceEnum("rsvp_lsps"),
	DHCPV4_INTERFACES: StatesResponseChoiceEnum("dhcpv4_interfaces"),
	DHCPV4_LEASES:     StatesResponseChoiceEnum("dhcpv4_leases"),
	DHCPV6_INTERFACES: StatesResponseChoiceEnum("dhcpv6_interfaces"),
	DHCPV6_LEASES:     StatesResponseChoiceEnum("dhcpv6_leases"),
	OSPFV2_LSAS:       StatesResponseChoiceEnum("ospfv2_lsas"),
}

func (obj *statesResponse) Choice() StatesResponseChoiceEnum {
	return StatesResponseChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *statesResponse) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *statesResponse) setChoice(value StatesResponseChoiceEnum) StatesResponse {
	intValue, ok := otg.StatesResponse_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StatesResponseChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StatesResponse_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
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

	if value == StatesResponseChoice.IPV4_NEIGHBORS {
		obj.obj.Ipv4Neighbors = []*otg.Neighborsv4State{}
	}

	if value == StatesResponseChoice.IPV6_NEIGHBORS {
		obj.obj.Ipv6Neighbors = []*otg.Neighborsv6State{}
	}

	if value == StatesResponseChoice.BGP_PREFIXES {
		obj.obj.BgpPrefixes = []*otg.BgpPrefixesState{}
	}

	if value == StatesResponseChoice.ISIS_LSPS {
		obj.obj.IsisLsps = []*otg.IsisLspsState{}
	}

	if value == StatesResponseChoice.LLDP_NEIGHBORS {
		obj.obj.LldpNeighbors = []*otg.LldpNeighborsState{}
	}

	if value == StatesResponseChoice.RSVP_LSPS {
		obj.obj.RsvpLsps = []*otg.RsvpLspsState{}
	}

	if value == StatesResponseChoice.DHCPV4_INTERFACES {
		obj.obj.Dhcpv4Interfaces = []*otg.Dhcpv4InterfaceState{}
	}

	if value == StatesResponseChoice.DHCPV4_LEASES {
		obj.obj.Dhcpv4Leases = []*otg.Dhcpv4LeasesState{}
	}

	if value == StatesResponseChoice.DHCPV6_INTERFACES {
		obj.obj.Dhcpv6Interfaces = []*otg.Dhcpv6InterfaceState{}
	}

	if value == StatesResponseChoice.DHCPV6_LEASES {
		obj.obj.Dhcpv6Leases = []*otg.Dhcpv6LeasesState{}
	}

	if value == StatesResponseChoice.OSPFV2_LSAS {
		obj.obj.Ospfv2Lsas = []*otg.Ospfv2LsaState{}
	}

	return obj
}

// description is TBD
// Ipv4Neighbors returns a []Neighborsv4State
func (obj *statesResponse) Ipv4Neighbors() StatesResponseNeighborsv4StateIter {
	if len(obj.obj.Ipv4Neighbors) == 0 {
		obj.setChoice(StatesResponseChoice.IPV4_NEIGHBORS)
	}
	if obj.ipv4NeighborsHolder == nil {
		obj.ipv4NeighborsHolder = newStatesResponseNeighborsv4StateIter(&obj.obj.Ipv4Neighbors).setMsg(obj)
	}
	return obj.ipv4NeighborsHolder
}

type statesResponseNeighborsv4StateIter struct {
	obj                   *statesResponse
	neighborsv4StateSlice []Neighborsv4State
	fieldPtr              *[]*otg.Neighborsv4State
}

func newStatesResponseNeighborsv4StateIter(ptr *[]*otg.Neighborsv4State) StatesResponseNeighborsv4StateIter {
	return &statesResponseNeighborsv4StateIter{fieldPtr: ptr}
}

type StatesResponseNeighborsv4StateIter interface {
	setMsg(*statesResponse) StatesResponseNeighborsv4StateIter
	Items() []Neighborsv4State
	Add() Neighborsv4State
	Append(items ...Neighborsv4State) StatesResponseNeighborsv4StateIter
	Set(index int, newObj Neighborsv4State) StatesResponseNeighborsv4StateIter
	Clear() StatesResponseNeighborsv4StateIter
	clearHolderSlice() StatesResponseNeighborsv4StateIter
	appendHolderSlice(item Neighborsv4State) StatesResponseNeighborsv4StateIter
}

func (obj *statesResponseNeighborsv4StateIter) setMsg(msg *statesResponse) StatesResponseNeighborsv4StateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&neighborsv4State{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *statesResponseNeighborsv4StateIter) Items() []Neighborsv4State {
	return obj.neighborsv4StateSlice
}

func (obj *statesResponseNeighborsv4StateIter) Add() Neighborsv4State {
	newObj := &otg.Neighborsv4State{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &neighborsv4State{obj: newObj}
	newLibObj.setDefault()
	obj.neighborsv4StateSlice = append(obj.neighborsv4StateSlice, newLibObj)
	return newLibObj
}

func (obj *statesResponseNeighborsv4StateIter) Append(items ...Neighborsv4State) StatesResponseNeighborsv4StateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.neighborsv4StateSlice = append(obj.neighborsv4StateSlice, item)
	}
	return obj
}

func (obj *statesResponseNeighborsv4StateIter) Set(index int, newObj Neighborsv4State) StatesResponseNeighborsv4StateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.neighborsv4StateSlice[index] = newObj
	return obj
}
func (obj *statesResponseNeighborsv4StateIter) Clear() StatesResponseNeighborsv4StateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Neighborsv4State{}
		obj.neighborsv4StateSlice = []Neighborsv4State{}
	}
	return obj
}
func (obj *statesResponseNeighborsv4StateIter) clearHolderSlice() StatesResponseNeighborsv4StateIter {
	if len(obj.neighborsv4StateSlice) > 0 {
		obj.neighborsv4StateSlice = []Neighborsv4State{}
	}
	return obj
}
func (obj *statesResponseNeighborsv4StateIter) appendHolderSlice(item Neighborsv4State) StatesResponseNeighborsv4StateIter {
	obj.neighborsv4StateSlice = append(obj.neighborsv4StateSlice, item)
	return obj
}

// description is TBD
// Ipv6Neighbors returns a []Neighborsv6State
func (obj *statesResponse) Ipv6Neighbors() StatesResponseNeighborsv6StateIter {
	if len(obj.obj.Ipv6Neighbors) == 0 {
		obj.setChoice(StatesResponseChoice.IPV6_NEIGHBORS)
	}
	if obj.ipv6NeighborsHolder == nil {
		obj.ipv6NeighborsHolder = newStatesResponseNeighborsv6StateIter(&obj.obj.Ipv6Neighbors).setMsg(obj)
	}
	return obj.ipv6NeighborsHolder
}

type statesResponseNeighborsv6StateIter struct {
	obj                   *statesResponse
	neighborsv6StateSlice []Neighborsv6State
	fieldPtr              *[]*otg.Neighborsv6State
}

func newStatesResponseNeighborsv6StateIter(ptr *[]*otg.Neighborsv6State) StatesResponseNeighborsv6StateIter {
	return &statesResponseNeighborsv6StateIter{fieldPtr: ptr}
}

type StatesResponseNeighborsv6StateIter interface {
	setMsg(*statesResponse) StatesResponseNeighborsv6StateIter
	Items() []Neighborsv6State
	Add() Neighborsv6State
	Append(items ...Neighborsv6State) StatesResponseNeighborsv6StateIter
	Set(index int, newObj Neighborsv6State) StatesResponseNeighborsv6StateIter
	Clear() StatesResponseNeighborsv6StateIter
	clearHolderSlice() StatesResponseNeighborsv6StateIter
	appendHolderSlice(item Neighborsv6State) StatesResponseNeighborsv6StateIter
}

func (obj *statesResponseNeighborsv6StateIter) setMsg(msg *statesResponse) StatesResponseNeighborsv6StateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&neighborsv6State{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *statesResponseNeighborsv6StateIter) Items() []Neighborsv6State {
	return obj.neighborsv6StateSlice
}

func (obj *statesResponseNeighborsv6StateIter) Add() Neighborsv6State {
	newObj := &otg.Neighborsv6State{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &neighborsv6State{obj: newObj}
	newLibObj.setDefault()
	obj.neighborsv6StateSlice = append(obj.neighborsv6StateSlice, newLibObj)
	return newLibObj
}

func (obj *statesResponseNeighborsv6StateIter) Append(items ...Neighborsv6State) StatesResponseNeighborsv6StateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.neighborsv6StateSlice = append(obj.neighborsv6StateSlice, item)
	}
	return obj
}

func (obj *statesResponseNeighborsv6StateIter) Set(index int, newObj Neighborsv6State) StatesResponseNeighborsv6StateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.neighborsv6StateSlice[index] = newObj
	return obj
}
func (obj *statesResponseNeighborsv6StateIter) Clear() StatesResponseNeighborsv6StateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Neighborsv6State{}
		obj.neighborsv6StateSlice = []Neighborsv6State{}
	}
	return obj
}
func (obj *statesResponseNeighborsv6StateIter) clearHolderSlice() StatesResponseNeighborsv6StateIter {
	if len(obj.neighborsv6StateSlice) > 0 {
		obj.neighborsv6StateSlice = []Neighborsv6State{}
	}
	return obj
}
func (obj *statesResponseNeighborsv6StateIter) appendHolderSlice(item Neighborsv6State) StatesResponseNeighborsv6StateIter {
	obj.neighborsv6StateSlice = append(obj.neighborsv6StateSlice, item)
	return obj
}

// description is TBD
// BgpPrefixes returns a []BgpPrefixesState
func (obj *statesResponse) BgpPrefixes() StatesResponseBgpPrefixesStateIter {
	if len(obj.obj.BgpPrefixes) == 0 {
		obj.setChoice(StatesResponseChoice.BGP_PREFIXES)
	}
	if obj.bgpPrefixesHolder == nil {
		obj.bgpPrefixesHolder = newStatesResponseBgpPrefixesStateIter(&obj.obj.BgpPrefixes).setMsg(obj)
	}
	return obj.bgpPrefixesHolder
}

type statesResponseBgpPrefixesStateIter struct {
	obj                   *statesResponse
	bgpPrefixesStateSlice []BgpPrefixesState
	fieldPtr              *[]*otg.BgpPrefixesState
}

func newStatesResponseBgpPrefixesStateIter(ptr *[]*otg.BgpPrefixesState) StatesResponseBgpPrefixesStateIter {
	return &statesResponseBgpPrefixesStateIter{fieldPtr: ptr}
}

type StatesResponseBgpPrefixesStateIter interface {
	setMsg(*statesResponse) StatesResponseBgpPrefixesStateIter
	Items() []BgpPrefixesState
	Add() BgpPrefixesState
	Append(items ...BgpPrefixesState) StatesResponseBgpPrefixesStateIter
	Set(index int, newObj BgpPrefixesState) StatesResponseBgpPrefixesStateIter
	Clear() StatesResponseBgpPrefixesStateIter
	clearHolderSlice() StatesResponseBgpPrefixesStateIter
	appendHolderSlice(item BgpPrefixesState) StatesResponseBgpPrefixesStateIter
}

func (obj *statesResponseBgpPrefixesStateIter) setMsg(msg *statesResponse) StatesResponseBgpPrefixesStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpPrefixesState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *statesResponseBgpPrefixesStateIter) Items() []BgpPrefixesState {
	return obj.bgpPrefixesStateSlice
}

func (obj *statesResponseBgpPrefixesStateIter) Add() BgpPrefixesState {
	newObj := &otg.BgpPrefixesState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpPrefixesState{obj: newObj}
	newLibObj.setDefault()
	obj.bgpPrefixesStateSlice = append(obj.bgpPrefixesStateSlice, newLibObj)
	return newLibObj
}

func (obj *statesResponseBgpPrefixesStateIter) Append(items ...BgpPrefixesState) StatesResponseBgpPrefixesStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpPrefixesStateSlice = append(obj.bgpPrefixesStateSlice, item)
	}
	return obj
}

func (obj *statesResponseBgpPrefixesStateIter) Set(index int, newObj BgpPrefixesState) StatesResponseBgpPrefixesStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpPrefixesStateSlice[index] = newObj
	return obj
}
func (obj *statesResponseBgpPrefixesStateIter) Clear() StatesResponseBgpPrefixesStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpPrefixesState{}
		obj.bgpPrefixesStateSlice = []BgpPrefixesState{}
	}
	return obj
}
func (obj *statesResponseBgpPrefixesStateIter) clearHolderSlice() StatesResponseBgpPrefixesStateIter {
	if len(obj.bgpPrefixesStateSlice) > 0 {
		obj.bgpPrefixesStateSlice = []BgpPrefixesState{}
	}
	return obj
}
func (obj *statesResponseBgpPrefixesStateIter) appendHolderSlice(item BgpPrefixesState) StatesResponseBgpPrefixesStateIter {
	obj.bgpPrefixesStateSlice = append(obj.bgpPrefixesStateSlice, item)
	return obj
}

// description is TBD
// IsisLsps returns a []IsisLspsState
func (obj *statesResponse) IsisLsps() StatesResponseIsisLspsStateIter {
	if len(obj.obj.IsisLsps) == 0 {
		obj.setChoice(StatesResponseChoice.ISIS_LSPS)
	}
	if obj.isisLspsHolder == nil {
		obj.isisLspsHolder = newStatesResponseIsisLspsStateIter(&obj.obj.IsisLsps).setMsg(obj)
	}
	return obj.isisLspsHolder
}

type statesResponseIsisLspsStateIter struct {
	obj                *statesResponse
	isisLspsStateSlice []IsisLspsState
	fieldPtr           *[]*otg.IsisLspsState
}

func newStatesResponseIsisLspsStateIter(ptr *[]*otg.IsisLspsState) StatesResponseIsisLspsStateIter {
	return &statesResponseIsisLspsStateIter{fieldPtr: ptr}
}

type StatesResponseIsisLspsStateIter interface {
	setMsg(*statesResponse) StatesResponseIsisLspsStateIter
	Items() []IsisLspsState
	Add() IsisLspsState
	Append(items ...IsisLspsState) StatesResponseIsisLspsStateIter
	Set(index int, newObj IsisLspsState) StatesResponseIsisLspsStateIter
	Clear() StatesResponseIsisLspsStateIter
	clearHolderSlice() StatesResponseIsisLspsStateIter
	appendHolderSlice(item IsisLspsState) StatesResponseIsisLspsStateIter
}

func (obj *statesResponseIsisLspsStateIter) setMsg(msg *statesResponse) StatesResponseIsisLspsStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspsState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *statesResponseIsisLspsStateIter) Items() []IsisLspsState {
	return obj.isisLspsStateSlice
}

func (obj *statesResponseIsisLspsStateIter) Add() IsisLspsState {
	newObj := &otg.IsisLspsState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspsState{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspsStateSlice = append(obj.isisLspsStateSlice, newLibObj)
	return newLibObj
}

func (obj *statesResponseIsisLspsStateIter) Append(items ...IsisLspsState) StatesResponseIsisLspsStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspsStateSlice = append(obj.isisLspsStateSlice, item)
	}
	return obj
}

func (obj *statesResponseIsisLspsStateIter) Set(index int, newObj IsisLspsState) StatesResponseIsisLspsStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspsStateSlice[index] = newObj
	return obj
}
func (obj *statesResponseIsisLspsStateIter) Clear() StatesResponseIsisLspsStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspsState{}
		obj.isisLspsStateSlice = []IsisLspsState{}
	}
	return obj
}
func (obj *statesResponseIsisLspsStateIter) clearHolderSlice() StatesResponseIsisLspsStateIter {
	if len(obj.isisLspsStateSlice) > 0 {
		obj.isisLspsStateSlice = []IsisLspsState{}
	}
	return obj
}
func (obj *statesResponseIsisLspsStateIter) appendHolderSlice(item IsisLspsState) StatesResponseIsisLspsStateIter {
	obj.isisLspsStateSlice = append(obj.isisLspsStateSlice, item)
	return obj
}

// description is TBD
// LldpNeighbors returns a []LldpNeighborsState
func (obj *statesResponse) LldpNeighbors() StatesResponseLldpNeighborsStateIter {
	if len(obj.obj.LldpNeighbors) == 0 {
		obj.setChoice(StatesResponseChoice.LLDP_NEIGHBORS)
	}
	if obj.lldpNeighborsHolder == nil {
		obj.lldpNeighborsHolder = newStatesResponseLldpNeighborsStateIter(&obj.obj.LldpNeighbors).setMsg(obj)
	}
	return obj.lldpNeighborsHolder
}

type statesResponseLldpNeighborsStateIter struct {
	obj                     *statesResponse
	lldpNeighborsStateSlice []LldpNeighborsState
	fieldPtr                *[]*otg.LldpNeighborsState
}

func newStatesResponseLldpNeighborsStateIter(ptr *[]*otg.LldpNeighborsState) StatesResponseLldpNeighborsStateIter {
	return &statesResponseLldpNeighborsStateIter{fieldPtr: ptr}
}

type StatesResponseLldpNeighborsStateIter interface {
	setMsg(*statesResponse) StatesResponseLldpNeighborsStateIter
	Items() []LldpNeighborsState
	Add() LldpNeighborsState
	Append(items ...LldpNeighborsState) StatesResponseLldpNeighborsStateIter
	Set(index int, newObj LldpNeighborsState) StatesResponseLldpNeighborsStateIter
	Clear() StatesResponseLldpNeighborsStateIter
	clearHolderSlice() StatesResponseLldpNeighborsStateIter
	appendHolderSlice(item LldpNeighborsState) StatesResponseLldpNeighborsStateIter
}

func (obj *statesResponseLldpNeighborsStateIter) setMsg(msg *statesResponse) StatesResponseLldpNeighborsStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&lldpNeighborsState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *statesResponseLldpNeighborsStateIter) Items() []LldpNeighborsState {
	return obj.lldpNeighborsStateSlice
}

func (obj *statesResponseLldpNeighborsStateIter) Add() LldpNeighborsState {
	newObj := &otg.LldpNeighborsState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &lldpNeighborsState{obj: newObj}
	newLibObj.setDefault()
	obj.lldpNeighborsStateSlice = append(obj.lldpNeighborsStateSlice, newLibObj)
	return newLibObj
}

func (obj *statesResponseLldpNeighborsStateIter) Append(items ...LldpNeighborsState) StatesResponseLldpNeighborsStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.lldpNeighborsStateSlice = append(obj.lldpNeighborsStateSlice, item)
	}
	return obj
}

func (obj *statesResponseLldpNeighborsStateIter) Set(index int, newObj LldpNeighborsState) StatesResponseLldpNeighborsStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.lldpNeighborsStateSlice[index] = newObj
	return obj
}
func (obj *statesResponseLldpNeighborsStateIter) Clear() StatesResponseLldpNeighborsStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LldpNeighborsState{}
		obj.lldpNeighborsStateSlice = []LldpNeighborsState{}
	}
	return obj
}
func (obj *statesResponseLldpNeighborsStateIter) clearHolderSlice() StatesResponseLldpNeighborsStateIter {
	if len(obj.lldpNeighborsStateSlice) > 0 {
		obj.lldpNeighborsStateSlice = []LldpNeighborsState{}
	}
	return obj
}
func (obj *statesResponseLldpNeighborsStateIter) appendHolderSlice(item LldpNeighborsState) StatesResponseLldpNeighborsStateIter {
	obj.lldpNeighborsStateSlice = append(obj.lldpNeighborsStateSlice, item)
	return obj
}

// description is TBD
// RsvpLsps returns a []RsvpLspsState
func (obj *statesResponse) RsvpLsps() StatesResponseRsvpLspsStateIter {
	if len(obj.obj.RsvpLsps) == 0 {
		obj.setChoice(StatesResponseChoice.RSVP_LSPS)
	}
	if obj.rsvpLspsHolder == nil {
		obj.rsvpLspsHolder = newStatesResponseRsvpLspsStateIter(&obj.obj.RsvpLsps).setMsg(obj)
	}
	return obj.rsvpLspsHolder
}

type statesResponseRsvpLspsStateIter struct {
	obj                *statesResponse
	rsvpLspsStateSlice []RsvpLspsState
	fieldPtr           *[]*otg.RsvpLspsState
}

func newStatesResponseRsvpLspsStateIter(ptr *[]*otg.RsvpLspsState) StatesResponseRsvpLspsStateIter {
	return &statesResponseRsvpLspsStateIter{fieldPtr: ptr}
}

type StatesResponseRsvpLspsStateIter interface {
	setMsg(*statesResponse) StatesResponseRsvpLspsStateIter
	Items() []RsvpLspsState
	Add() RsvpLspsState
	Append(items ...RsvpLspsState) StatesResponseRsvpLspsStateIter
	Set(index int, newObj RsvpLspsState) StatesResponseRsvpLspsStateIter
	Clear() StatesResponseRsvpLspsStateIter
	clearHolderSlice() StatesResponseRsvpLspsStateIter
	appendHolderSlice(item RsvpLspsState) StatesResponseRsvpLspsStateIter
}

func (obj *statesResponseRsvpLspsStateIter) setMsg(msg *statesResponse) StatesResponseRsvpLspsStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rsvpLspsState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *statesResponseRsvpLspsStateIter) Items() []RsvpLspsState {
	return obj.rsvpLspsStateSlice
}

func (obj *statesResponseRsvpLspsStateIter) Add() RsvpLspsState {
	newObj := &otg.RsvpLspsState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rsvpLspsState{obj: newObj}
	newLibObj.setDefault()
	obj.rsvpLspsStateSlice = append(obj.rsvpLspsStateSlice, newLibObj)
	return newLibObj
}

func (obj *statesResponseRsvpLspsStateIter) Append(items ...RsvpLspsState) StatesResponseRsvpLspsStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rsvpLspsStateSlice = append(obj.rsvpLspsStateSlice, item)
	}
	return obj
}

func (obj *statesResponseRsvpLspsStateIter) Set(index int, newObj RsvpLspsState) StatesResponseRsvpLspsStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rsvpLspsStateSlice[index] = newObj
	return obj
}
func (obj *statesResponseRsvpLspsStateIter) Clear() StatesResponseRsvpLspsStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RsvpLspsState{}
		obj.rsvpLspsStateSlice = []RsvpLspsState{}
	}
	return obj
}
func (obj *statesResponseRsvpLspsStateIter) clearHolderSlice() StatesResponseRsvpLspsStateIter {
	if len(obj.rsvpLspsStateSlice) > 0 {
		obj.rsvpLspsStateSlice = []RsvpLspsState{}
	}
	return obj
}
func (obj *statesResponseRsvpLspsStateIter) appendHolderSlice(item RsvpLspsState) StatesResponseRsvpLspsStateIter {
	obj.rsvpLspsStateSlice = append(obj.rsvpLspsStateSlice, item)
	return obj
}

// description is TBD
// Dhcpv4Interfaces returns a []Dhcpv4InterfaceState
func (obj *statesResponse) Dhcpv4Interfaces() StatesResponseDhcpv4InterfaceStateIter {
	if len(obj.obj.Dhcpv4Interfaces) == 0 {
		obj.setChoice(StatesResponseChoice.DHCPV4_INTERFACES)
	}
	if obj.dhcpv4InterfacesHolder == nil {
		obj.dhcpv4InterfacesHolder = newStatesResponseDhcpv4InterfaceStateIter(&obj.obj.Dhcpv4Interfaces).setMsg(obj)
	}
	return obj.dhcpv4InterfacesHolder
}

type statesResponseDhcpv4InterfaceStateIter struct {
	obj                       *statesResponse
	dhcpv4InterfaceStateSlice []Dhcpv4InterfaceState
	fieldPtr                  *[]*otg.Dhcpv4InterfaceState
}

func newStatesResponseDhcpv4InterfaceStateIter(ptr *[]*otg.Dhcpv4InterfaceState) StatesResponseDhcpv4InterfaceStateIter {
	return &statesResponseDhcpv4InterfaceStateIter{fieldPtr: ptr}
}

type StatesResponseDhcpv4InterfaceStateIter interface {
	setMsg(*statesResponse) StatesResponseDhcpv4InterfaceStateIter
	Items() []Dhcpv4InterfaceState
	Add() Dhcpv4InterfaceState
	Append(items ...Dhcpv4InterfaceState) StatesResponseDhcpv4InterfaceStateIter
	Set(index int, newObj Dhcpv4InterfaceState) StatesResponseDhcpv4InterfaceStateIter
	Clear() StatesResponseDhcpv4InterfaceStateIter
	clearHolderSlice() StatesResponseDhcpv4InterfaceStateIter
	appendHolderSlice(item Dhcpv4InterfaceState) StatesResponseDhcpv4InterfaceStateIter
}

func (obj *statesResponseDhcpv4InterfaceStateIter) setMsg(msg *statesResponse) StatesResponseDhcpv4InterfaceStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv4InterfaceState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *statesResponseDhcpv4InterfaceStateIter) Items() []Dhcpv4InterfaceState {
	return obj.dhcpv4InterfaceStateSlice
}

func (obj *statesResponseDhcpv4InterfaceStateIter) Add() Dhcpv4InterfaceState {
	newObj := &otg.Dhcpv4InterfaceState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv4InterfaceState{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv4InterfaceStateSlice = append(obj.dhcpv4InterfaceStateSlice, newLibObj)
	return newLibObj
}

func (obj *statesResponseDhcpv4InterfaceStateIter) Append(items ...Dhcpv4InterfaceState) StatesResponseDhcpv4InterfaceStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv4InterfaceStateSlice = append(obj.dhcpv4InterfaceStateSlice, item)
	}
	return obj
}

func (obj *statesResponseDhcpv4InterfaceStateIter) Set(index int, newObj Dhcpv4InterfaceState) StatesResponseDhcpv4InterfaceStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv4InterfaceStateSlice[index] = newObj
	return obj
}
func (obj *statesResponseDhcpv4InterfaceStateIter) Clear() StatesResponseDhcpv4InterfaceStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv4InterfaceState{}
		obj.dhcpv4InterfaceStateSlice = []Dhcpv4InterfaceState{}
	}
	return obj
}
func (obj *statesResponseDhcpv4InterfaceStateIter) clearHolderSlice() StatesResponseDhcpv4InterfaceStateIter {
	if len(obj.dhcpv4InterfaceStateSlice) > 0 {
		obj.dhcpv4InterfaceStateSlice = []Dhcpv4InterfaceState{}
	}
	return obj
}
func (obj *statesResponseDhcpv4InterfaceStateIter) appendHolderSlice(item Dhcpv4InterfaceState) StatesResponseDhcpv4InterfaceStateIter {
	obj.dhcpv4InterfaceStateSlice = append(obj.dhcpv4InterfaceStateSlice, item)
	return obj
}

// description is TBD
// Dhcpv4Leases returns a []Dhcpv4LeasesState
func (obj *statesResponse) Dhcpv4Leases() StatesResponseDhcpv4LeasesStateIter {
	if len(obj.obj.Dhcpv4Leases) == 0 {
		obj.setChoice(StatesResponseChoice.DHCPV4_LEASES)
	}
	if obj.dhcpv4LeasesHolder == nil {
		obj.dhcpv4LeasesHolder = newStatesResponseDhcpv4LeasesStateIter(&obj.obj.Dhcpv4Leases).setMsg(obj)
	}
	return obj.dhcpv4LeasesHolder
}

type statesResponseDhcpv4LeasesStateIter struct {
	obj                    *statesResponse
	dhcpv4LeasesStateSlice []Dhcpv4LeasesState
	fieldPtr               *[]*otg.Dhcpv4LeasesState
}

func newStatesResponseDhcpv4LeasesStateIter(ptr *[]*otg.Dhcpv4LeasesState) StatesResponseDhcpv4LeasesStateIter {
	return &statesResponseDhcpv4LeasesStateIter{fieldPtr: ptr}
}

type StatesResponseDhcpv4LeasesStateIter interface {
	setMsg(*statesResponse) StatesResponseDhcpv4LeasesStateIter
	Items() []Dhcpv4LeasesState
	Add() Dhcpv4LeasesState
	Append(items ...Dhcpv4LeasesState) StatesResponseDhcpv4LeasesStateIter
	Set(index int, newObj Dhcpv4LeasesState) StatesResponseDhcpv4LeasesStateIter
	Clear() StatesResponseDhcpv4LeasesStateIter
	clearHolderSlice() StatesResponseDhcpv4LeasesStateIter
	appendHolderSlice(item Dhcpv4LeasesState) StatesResponseDhcpv4LeasesStateIter
}

func (obj *statesResponseDhcpv4LeasesStateIter) setMsg(msg *statesResponse) StatesResponseDhcpv4LeasesStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv4LeasesState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *statesResponseDhcpv4LeasesStateIter) Items() []Dhcpv4LeasesState {
	return obj.dhcpv4LeasesStateSlice
}

func (obj *statesResponseDhcpv4LeasesStateIter) Add() Dhcpv4LeasesState {
	newObj := &otg.Dhcpv4LeasesState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv4LeasesState{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv4LeasesStateSlice = append(obj.dhcpv4LeasesStateSlice, newLibObj)
	return newLibObj
}

func (obj *statesResponseDhcpv4LeasesStateIter) Append(items ...Dhcpv4LeasesState) StatesResponseDhcpv4LeasesStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv4LeasesStateSlice = append(obj.dhcpv4LeasesStateSlice, item)
	}
	return obj
}

func (obj *statesResponseDhcpv4LeasesStateIter) Set(index int, newObj Dhcpv4LeasesState) StatesResponseDhcpv4LeasesStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv4LeasesStateSlice[index] = newObj
	return obj
}
func (obj *statesResponseDhcpv4LeasesStateIter) Clear() StatesResponseDhcpv4LeasesStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv4LeasesState{}
		obj.dhcpv4LeasesStateSlice = []Dhcpv4LeasesState{}
	}
	return obj
}
func (obj *statesResponseDhcpv4LeasesStateIter) clearHolderSlice() StatesResponseDhcpv4LeasesStateIter {
	if len(obj.dhcpv4LeasesStateSlice) > 0 {
		obj.dhcpv4LeasesStateSlice = []Dhcpv4LeasesState{}
	}
	return obj
}
func (obj *statesResponseDhcpv4LeasesStateIter) appendHolderSlice(item Dhcpv4LeasesState) StatesResponseDhcpv4LeasesStateIter {
	obj.dhcpv4LeasesStateSlice = append(obj.dhcpv4LeasesStateSlice, item)
	return obj
}

// description is TBD
// Dhcpv6Interfaces returns a []Dhcpv6InterfaceState
func (obj *statesResponse) Dhcpv6Interfaces() StatesResponseDhcpv6InterfaceStateIter {
	if len(obj.obj.Dhcpv6Interfaces) == 0 {
		obj.setChoice(StatesResponseChoice.DHCPV6_INTERFACES)
	}
	if obj.dhcpv6InterfacesHolder == nil {
		obj.dhcpv6InterfacesHolder = newStatesResponseDhcpv6InterfaceStateIter(&obj.obj.Dhcpv6Interfaces).setMsg(obj)
	}
	return obj.dhcpv6InterfacesHolder
}

type statesResponseDhcpv6InterfaceStateIter struct {
	obj                       *statesResponse
	dhcpv6InterfaceStateSlice []Dhcpv6InterfaceState
	fieldPtr                  *[]*otg.Dhcpv6InterfaceState
}

func newStatesResponseDhcpv6InterfaceStateIter(ptr *[]*otg.Dhcpv6InterfaceState) StatesResponseDhcpv6InterfaceStateIter {
	return &statesResponseDhcpv6InterfaceStateIter{fieldPtr: ptr}
}

type StatesResponseDhcpv6InterfaceStateIter interface {
	setMsg(*statesResponse) StatesResponseDhcpv6InterfaceStateIter
	Items() []Dhcpv6InterfaceState
	Add() Dhcpv6InterfaceState
	Append(items ...Dhcpv6InterfaceState) StatesResponseDhcpv6InterfaceStateIter
	Set(index int, newObj Dhcpv6InterfaceState) StatesResponseDhcpv6InterfaceStateIter
	Clear() StatesResponseDhcpv6InterfaceStateIter
	clearHolderSlice() StatesResponseDhcpv6InterfaceStateIter
	appendHolderSlice(item Dhcpv6InterfaceState) StatesResponseDhcpv6InterfaceStateIter
}

func (obj *statesResponseDhcpv6InterfaceStateIter) setMsg(msg *statesResponse) StatesResponseDhcpv6InterfaceStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6InterfaceState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *statesResponseDhcpv6InterfaceStateIter) Items() []Dhcpv6InterfaceState {
	return obj.dhcpv6InterfaceStateSlice
}

func (obj *statesResponseDhcpv6InterfaceStateIter) Add() Dhcpv6InterfaceState {
	newObj := &otg.Dhcpv6InterfaceState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6InterfaceState{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6InterfaceStateSlice = append(obj.dhcpv6InterfaceStateSlice, newLibObj)
	return newLibObj
}

func (obj *statesResponseDhcpv6InterfaceStateIter) Append(items ...Dhcpv6InterfaceState) StatesResponseDhcpv6InterfaceStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6InterfaceStateSlice = append(obj.dhcpv6InterfaceStateSlice, item)
	}
	return obj
}

func (obj *statesResponseDhcpv6InterfaceStateIter) Set(index int, newObj Dhcpv6InterfaceState) StatesResponseDhcpv6InterfaceStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6InterfaceStateSlice[index] = newObj
	return obj
}
func (obj *statesResponseDhcpv6InterfaceStateIter) Clear() StatesResponseDhcpv6InterfaceStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6InterfaceState{}
		obj.dhcpv6InterfaceStateSlice = []Dhcpv6InterfaceState{}
	}
	return obj
}
func (obj *statesResponseDhcpv6InterfaceStateIter) clearHolderSlice() StatesResponseDhcpv6InterfaceStateIter {
	if len(obj.dhcpv6InterfaceStateSlice) > 0 {
		obj.dhcpv6InterfaceStateSlice = []Dhcpv6InterfaceState{}
	}
	return obj
}
func (obj *statesResponseDhcpv6InterfaceStateIter) appendHolderSlice(item Dhcpv6InterfaceState) StatesResponseDhcpv6InterfaceStateIter {
	obj.dhcpv6InterfaceStateSlice = append(obj.dhcpv6InterfaceStateSlice, item)
	return obj
}

// description is TBD
// Dhcpv6Leases returns a []Dhcpv6LeasesState
func (obj *statesResponse) Dhcpv6Leases() StatesResponseDhcpv6LeasesStateIter {
	if len(obj.obj.Dhcpv6Leases) == 0 {
		obj.setChoice(StatesResponseChoice.DHCPV6_LEASES)
	}
	if obj.dhcpv6LeasesHolder == nil {
		obj.dhcpv6LeasesHolder = newStatesResponseDhcpv6LeasesStateIter(&obj.obj.Dhcpv6Leases).setMsg(obj)
	}
	return obj.dhcpv6LeasesHolder
}

type statesResponseDhcpv6LeasesStateIter struct {
	obj                    *statesResponse
	dhcpv6LeasesStateSlice []Dhcpv6LeasesState
	fieldPtr               *[]*otg.Dhcpv6LeasesState
}

func newStatesResponseDhcpv6LeasesStateIter(ptr *[]*otg.Dhcpv6LeasesState) StatesResponseDhcpv6LeasesStateIter {
	return &statesResponseDhcpv6LeasesStateIter{fieldPtr: ptr}
}

type StatesResponseDhcpv6LeasesStateIter interface {
	setMsg(*statesResponse) StatesResponseDhcpv6LeasesStateIter
	Items() []Dhcpv6LeasesState
	Add() Dhcpv6LeasesState
	Append(items ...Dhcpv6LeasesState) StatesResponseDhcpv6LeasesStateIter
	Set(index int, newObj Dhcpv6LeasesState) StatesResponseDhcpv6LeasesStateIter
	Clear() StatesResponseDhcpv6LeasesStateIter
	clearHolderSlice() StatesResponseDhcpv6LeasesStateIter
	appendHolderSlice(item Dhcpv6LeasesState) StatesResponseDhcpv6LeasesStateIter
}

func (obj *statesResponseDhcpv6LeasesStateIter) setMsg(msg *statesResponse) StatesResponseDhcpv6LeasesStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6LeasesState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *statesResponseDhcpv6LeasesStateIter) Items() []Dhcpv6LeasesState {
	return obj.dhcpv6LeasesStateSlice
}

func (obj *statesResponseDhcpv6LeasesStateIter) Add() Dhcpv6LeasesState {
	newObj := &otg.Dhcpv6LeasesState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6LeasesState{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6LeasesStateSlice = append(obj.dhcpv6LeasesStateSlice, newLibObj)
	return newLibObj
}

func (obj *statesResponseDhcpv6LeasesStateIter) Append(items ...Dhcpv6LeasesState) StatesResponseDhcpv6LeasesStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6LeasesStateSlice = append(obj.dhcpv6LeasesStateSlice, item)
	}
	return obj
}

func (obj *statesResponseDhcpv6LeasesStateIter) Set(index int, newObj Dhcpv6LeasesState) StatesResponseDhcpv6LeasesStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6LeasesStateSlice[index] = newObj
	return obj
}
func (obj *statesResponseDhcpv6LeasesStateIter) Clear() StatesResponseDhcpv6LeasesStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6LeasesState{}
		obj.dhcpv6LeasesStateSlice = []Dhcpv6LeasesState{}
	}
	return obj
}
func (obj *statesResponseDhcpv6LeasesStateIter) clearHolderSlice() StatesResponseDhcpv6LeasesStateIter {
	if len(obj.dhcpv6LeasesStateSlice) > 0 {
		obj.dhcpv6LeasesStateSlice = []Dhcpv6LeasesState{}
	}
	return obj
}
func (obj *statesResponseDhcpv6LeasesStateIter) appendHolderSlice(item Dhcpv6LeasesState) StatesResponseDhcpv6LeasesStateIter {
	obj.dhcpv6LeasesStateSlice = append(obj.dhcpv6LeasesStateSlice, item)
	return obj
}

// description is TBD
// Ospfv2Lsas returns a []Ospfv2LsaState
func (obj *statesResponse) Ospfv2Lsas() StatesResponseOspfv2LsaStateIter {
	if len(obj.obj.Ospfv2Lsas) == 0 {
		obj.setChoice(StatesResponseChoice.OSPFV2_LSAS)
	}
	if obj.ospfv2LsasHolder == nil {
		obj.ospfv2LsasHolder = newStatesResponseOspfv2LsaStateIter(&obj.obj.Ospfv2Lsas).setMsg(obj)
	}
	return obj.ospfv2LsasHolder
}

type statesResponseOspfv2LsaStateIter struct {
	obj                 *statesResponse
	ospfv2LsaStateSlice []Ospfv2LsaState
	fieldPtr            *[]*otg.Ospfv2LsaState
}

func newStatesResponseOspfv2LsaStateIter(ptr *[]*otg.Ospfv2LsaState) StatesResponseOspfv2LsaStateIter {
	return &statesResponseOspfv2LsaStateIter{fieldPtr: ptr}
}

type StatesResponseOspfv2LsaStateIter interface {
	setMsg(*statesResponse) StatesResponseOspfv2LsaStateIter
	Items() []Ospfv2LsaState
	Add() Ospfv2LsaState
	Append(items ...Ospfv2LsaState) StatesResponseOspfv2LsaStateIter
	Set(index int, newObj Ospfv2LsaState) StatesResponseOspfv2LsaStateIter
	Clear() StatesResponseOspfv2LsaStateIter
	clearHolderSlice() StatesResponseOspfv2LsaStateIter
	appendHolderSlice(item Ospfv2LsaState) StatesResponseOspfv2LsaStateIter
}

func (obj *statesResponseOspfv2LsaStateIter) setMsg(msg *statesResponse) StatesResponseOspfv2LsaStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2LsaState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *statesResponseOspfv2LsaStateIter) Items() []Ospfv2LsaState {
	return obj.ospfv2LsaStateSlice
}

func (obj *statesResponseOspfv2LsaStateIter) Add() Ospfv2LsaState {
	newObj := &otg.Ospfv2LsaState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2LsaState{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2LsaStateSlice = append(obj.ospfv2LsaStateSlice, newLibObj)
	return newLibObj
}

func (obj *statesResponseOspfv2LsaStateIter) Append(items ...Ospfv2LsaState) StatesResponseOspfv2LsaStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2LsaStateSlice = append(obj.ospfv2LsaStateSlice, item)
	}
	return obj
}

func (obj *statesResponseOspfv2LsaStateIter) Set(index int, newObj Ospfv2LsaState) StatesResponseOspfv2LsaStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2LsaStateSlice[index] = newObj
	return obj
}
func (obj *statesResponseOspfv2LsaStateIter) Clear() StatesResponseOspfv2LsaStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2LsaState{}
		obj.ospfv2LsaStateSlice = []Ospfv2LsaState{}
	}
	return obj
}
func (obj *statesResponseOspfv2LsaStateIter) clearHolderSlice() StatesResponseOspfv2LsaStateIter {
	if len(obj.ospfv2LsaStateSlice) > 0 {
		obj.ospfv2LsaStateSlice = []Ospfv2LsaState{}
	}
	return obj
}
func (obj *statesResponseOspfv2LsaStateIter) appendHolderSlice(item Ospfv2LsaState) StatesResponseOspfv2LsaStateIter {
	obj.ospfv2LsaStateSlice = append(obj.ospfv2LsaStateSlice, item)
	return obj
}

func (obj *statesResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4Neighbors) != 0 {

		if set_default {
			obj.Ipv4Neighbors().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Neighbors {
				obj.Ipv4Neighbors().appendHolderSlice(&neighborsv4State{obj: item})
			}
		}
		for _, item := range obj.Ipv4Neighbors().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv6Neighbors) != 0 {

		if set_default {
			obj.Ipv6Neighbors().clearHolderSlice()
			for _, item := range obj.obj.Ipv6Neighbors {
				obj.Ipv6Neighbors().appendHolderSlice(&neighborsv6State{obj: item})
			}
		}
		for _, item := range obj.Ipv6Neighbors().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.BgpPrefixes) != 0 {

		if set_default {
			obj.BgpPrefixes().clearHolderSlice()
			for _, item := range obj.obj.BgpPrefixes {
				obj.BgpPrefixes().appendHolderSlice(&bgpPrefixesState{obj: item})
			}
		}
		for _, item := range obj.BgpPrefixes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.IsisLsps) != 0 {

		if set_default {
			obj.IsisLsps().clearHolderSlice()
			for _, item := range obj.obj.IsisLsps {
				obj.IsisLsps().appendHolderSlice(&isisLspsState{obj: item})
			}
		}
		for _, item := range obj.IsisLsps().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.LldpNeighbors) != 0 {

		if set_default {
			obj.LldpNeighbors().clearHolderSlice()
			for _, item := range obj.obj.LldpNeighbors {
				obj.LldpNeighbors().appendHolderSlice(&lldpNeighborsState{obj: item})
			}
		}
		for _, item := range obj.LldpNeighbors().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.RsvpLsps) != 0 {

		if set_default {
			obj.RsvpLsps().clearHolderSlice()
			for _, item := range obj.obj.RsvpLsps {
				obj.RsvpLsps().appendHolderSlice(&rsvpLspsState{obj: item})
			}
		}
		for _, item := range obj.RsvpLsps().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Dhcpv4Interfaces) != 0 {

		if set_default {
			obj.Dhcpv4Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Dhcpv4Interfaces {
				obj.Dhcpv4Interfaces().appendHolderSlice(&dhcpv4InterfaceState{obj: item})
			}
		}
		for _, item := range obj.Dhcpv4Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Dhcpv4Leases) != 0 {

		if set_default {
			obj.Dhcpv4Leases().clearHolderSlice()
			for _, item := range obj.obj.Dhcpv4Leases {
				obj.Dhcpv4Leases().appendHolderSlice(&dhcpv4LeasesState{obj: item})
			}
		}
		for _, item := range obj.Dhcpv4Leases().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Dhcpv6Interfaces) != 0 {

		if set_default {
			obj.Dhcpv6Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Dhcpv6Interfaces {
				obj.Dhcpv6Interfaces().appendHolderSlice(&dhcpv6InterfaceState{obj: item})
			}
		}
		for _, item := range obj.Dhcpv6Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Dhcpv6Leases) != 0 {

		if set_default {
			obj.Dhcpv6Leases().clearHolderSlice()
			for _, item := range obj.obj.Dhcpv6Leases {
				obj.Dhcpv6Leases().appendHolderSlice(&dhcpv6LeasesState{obj: item})
			}
		}
		for _, item := range obj.Dhcpv6Leases().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ospfv2Lsas) != 0 {

		if set_default {
			obj.Ospfv2Lsas().clearHolderSlice()
			for _, item := range obj.obj.Ospfv2Lsas {
				obj.Ospfv2Lsas().appendHolderSlice(&ospfv2LsaState{obj: item})
			}
		}
		for _, item := range obj.Ospfv2Lsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *statesResponse) setDefault() {
	var choices_set int = 0
	var choice StatesResponseChoiceEnum

	if len(obj.obj.Ipv4Neighbors) > 0 {
		choices_set += 1
		choice = StatesResponseChoice.IPV4_NEIGHBORS
	}

	if len(obj.obj.Ipv6Neighbors) > 0 {
		choices_set += 1
		choice = StatesResponseChoice.IPV6_NEIGHBORS
	}

	if len(obj.obj.BgpPrefixes) > 0 {
		choices_set += 1
		choice = StatesResponseChoice.BGP_PREFIXES
	}

	if len(obj.obj.IsisLsps) > 0 {
		choices_set += 1
		choice = StatesResponseChoice.ISIS_LSPS
	}

	if len(obj.obj.LldpNeighbors) > 0 {
		choices_set += 1
		choice = StatesResponseChoice.LLDP_NEIGHBORS
	}

	if len(obj.obj.RsvpLsps) > 0 {
		choices_set += 1
		choice = StatesResponseChoice.RSVP_LSPS
	}

	if len(obj.obj.Dhcpv4Interfaces) > 0 {
		choices_set += 1
		choice = StatesResponseChoice.DHCPV4_INTERFACES
	}

	if len(obj.obj.Dhcpv4Leases) > 0 {
		choices_set += 1
		choice = StatesResponseChoice.DHCPV4_LEASES
	}

	if len(obj.obj.Dhcpv6Interfaces) > 0 {
		choices_set += 1
		choice = StatesResponseChoice.DHCPV6_INTERFACES
	}

	if len(obj.obj.Dhcpv6Leases) > 0 {
		choices_set += 1
		choice = StatesResponseChoice.DHCPV6_LEASES
	}

	if len(obj.obj.Ospfv2Lsas) > 0 {
		choices_set += 1
		choice = StatesResponseChoice.OSPFV2_LSAS
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(StatesResponseChoice.IPV4_NEIGHBORS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in StatesResponse")
			}
		} else {
			intVal := otg.StatesResponse_Choice_Enum_value[string(choice)]
			enumValue := otg.StatesResponse_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
