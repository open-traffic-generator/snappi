package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2Interface *****
type ospfv2Interface struct {
	validation
	obj                      *otg.Ospfv2Interface
	marshaller               marshalOspfv2Interface
	unMarshaller             unMarshalOspfv2Interface
	neighborsHolder          Ospfv2InterfaceOspfv2InterfaceNeighborIter
	trafficEngineeringHolder Ospfv2InterfaceLinkStateTEIter
	authenticationHolder     Ospfv2InterfaceAuthentication
	advancedHolder           Ospfv2InterfaceAdvanced
	linkProtectionHolder     Ospfv2InterfaceLinkProtection
}

func NewOspfv2Interface() Ospfv2Interface {
	obj := ospfv2Interface{obj: &otg.Ospfv2Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2Interface) msg() *otg.Ospfv2Interface {
	return obj.obj
}

func (obj *ospfv2Interface) setMsg(msg *otg.Ospfv2Interface) Ospfv2Interface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2Interface struct {
	obj *ospfv2Interface
}

type marshalOspfv2Interface interface {
	// ToProto marshals Ospfv2Interface to protobuf object *otg.Ospfv2Interface
	ToProto() (*otg.Ospfv2Interface, error)
	// ToPbText marshals Ospfv2Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2Interface to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2Interface struct {
	obj *ospfv2Interface
}

type unMarshalOspfv2Interface interface {
	// FromProto unmarshals Ospfv2Interface from protobuf object *otg.Ospfv2Interface
	FromProto(msg *otg.Ospfv2Interface) (Ospfv2Interface, error)
	// FromPbText unmarshals Ospfv2Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2Interface from JSON text
	FromJson(value string) error
}

func (obj *ospfv2Interface) Marshal() marshalOspfv2Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2Interface) Unmarshal() unMarshalOspfv2Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2Interface) ToProto() (*otg.Ospfv2Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2Interface) FromProto(msg *otg.Ospfv2Interface) (Ospfv2Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2Interface) ToPbText() (string, error) {
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

func (m *unMarshalospfv2Interface) FromPbText(value string) error {
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

func (m *marshalospfv2Interface) ToYaml() (string, error) {
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

func (m *unMarshalospfv2Interface) FromYaml(value string) error {
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

func (m *marshalospfv2Interface) ToJson() (string, error) {
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

func (m *unMarshalospfv2Interface) FromJson(value string) error {
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

func (obj *ospfv2Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2Interface) Clone() (Ospfv2Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2Interface()
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

func (obj *ospfv2Interface) setNil() {
	obj.neighborsHolder = nil
	obj.trafficEngineeringHolder = nil
	obj.authenticationHolder = nil
	obj.advancedHolder = nil
	obj.linkProtectionHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2Interface is configuration for single OSPFv2 interface.
type Ospfv2Interface interface {
	Validation
	// msg marshals Ospfv2Interface to protobuf object *otg.Ospfv2Interface
	// and doesn't set defaults
	msg() *otg.Ospfv2Interface
	// setMsg unmarshals Ospfv2Interface from protobuf object *otg.Ospfv2Interface
	// and doesn't set defaults
	setMsg(*otg.Ospfv2Interface) Ospfv2Interface
	// provides marshal interface
	Marshal() marshalOspfv2Interface
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2Interface
	// validate validates Ospfv2Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Ospfv2Interface.
	Name() string
	// SetName assigns string provided by user to Ospfv2Interface
	SetName(value string) Ospfv2Interface
	// Choice returns Ospfv2InterfaceChoiceEnum, set in Ospfv2Interface
	Choice() Ospfv2InterfaceChoiceEnum
	// setChoice assigns Ospfv2InterfaceChoiceEnum provided by user to Ospfv2Interface
	setChoice(value Ospfv2InterfaceChoiceEnum) Ospfv2Interface
	// HasChoice checks if Choice has been set in Ospfv2Interface
	HasChoice() bool
	// Ipv4Name returns string, set in Ospfv2Interface.
	Ipv4Name() string
	// SetIpv4Name assigns string provided by user to Ospfv2Interface
	SetIpv4Name(value string) Ospfv2Interface
	// HasIpv4Name checks if Ipv4Name has been set in Ospfv2Interface
	HasIpv4Name() bool
	// DhcpIntfName returns string, set in Ospfv2Interface.
	DhcpIntfName() string
	// SetDhcpIntfName assigns string provided by user to Ospfv2Interface
	SetDhcpIntfName(value string) Ospfv2Interface
	// HasDhcpIntfName checks if DhcpIntfName has been set in Ospfv2Interface
	HasDhcpIntfName() bool
	// AreaType returns Ospfv2InterfaceAreaTypeEnum, set in Ospfv2Interface
	AreaType() Ospfv2InterfaceAreaTypeEnum
	// SetAreaType assigns Ospfv2InterfaceAreaTypeEnum provided by user to Ospfv2Interface
	SetAreaType(value Ospfv2InterfaceAreaTypeEnum) Ospfv2Interface
	// HasAreaType checks if AreaType has been set in Ospfv2Interface
	HasAreaType() bool
	// AreaId returns uint32, set in Ospfv2Interface.
	AreaId() uint32
	// SetAreaId assigns uint32 provided by user to Ospfv2Interface
	SetAreaId(value uint32) Ospfv2Interface
	// HasAreaId checks if AreaId has been set in Ospfv2Interface
	HasAreaId() bool
	// AreaIdAsIp returns string, set in Ospfv2Interface.
	AreaIdAsIp() string
	// SetAreaIdAsIp assigns string provided by user to Ospfv2Interface
	SetAreaIdAsIp(value string) Ospfv2Interface
	// HasAreaIdAsIp checks if AreaIdAsIp has been set in Ospfv2Interface
	HasAreaIdAsIp() bool
	// NetworkType returns Ospfv2InterfaceNetworkTypeEnum, set in Ospfv2Interface
	NetworkType() Ospfv2InterfaceNetworkTypeEnum
	// SetNetworkType assigns Ospfv2InterfaceNetworkTypeEnum provided by user to Ospfv2Interface
	SetNetworkType(value Ospfv2InterfaceNetworkTypeEnum) Ospfv2Interface
	// HasNetworkType checks if NetworkType has been set in Ospfv2Interface
	HasNetworkType() bool
	// Neighbors returns Ospfv2InterfaceOspfv2InterfaceNeighborIterIter, set in Ospfv2Interface
	Neighbors() Ospfv2InterfaceOspfv2InterfaceNeighborIter
	// TrafficEngineering returns Ospfv2InterfaceLinkStateTEIterIter, set in Ospfv2Interface
	TrafficEngineering() Ospfv2InterfaceLinkStateTEIter
	// Authentication returns Ospfv2InterfaceAuthentication, set in Ospfv2Interface.
	// Ospfv2InterfaceAuthentication is this contains OSPFv2 authentication properties.
	// Reference: https://www.rfc-editor.org/rfc/rfc2328#appendix-D
	Authentication() Ospfv2InterfaceAuthentication
	// SetAuthentication assigns Ospfv2InterfaceAuthentication provided by user to Ospfv2Interface.
	// Ospfv2InterfaceAuthentication is this contains OSPFv2 authentication properties.
	// Reference: https://www.rfc-editor.org/rfc/rfc2328#appendix-D
	SetAuthentication(value Ospfv2InterfaceAuthentication) Ospfv2Interface
	// HasAuthentication checks if Authentication has been set in Ospfv2Interface
	HasAuthentication() bool
	// Advanced returns Ospfv2InterfaceAdvanced, set in Ospfv2Interface.
	// Ospfv2InterfaceAdvanced is contains OSPFv2 advanced properties.
	Advanced() Ospfv2InterfaceAdvanced
	// SetAdvanced assigns Ospfv2InterfaceAdvanced provided by user to Ospfv2Interface.
	// Ospfv2InterfaceAdvanced is contains OSPFv2 advanced properties.
	SetAdvanced(value Ospfv2InterfaceAdvanced) Ospfv2Interface
	// HasAdvanced checks if Advanced has been set in Ospfv2Interface
	HasAdvanced() bool
	// LinkProtection returns Ospfv2InterfaceLinkProtection, set in Ospfv2Interface.
	// Ospfv2InterfaceLinkProtection is optional container for the link protection sub TLV (type 20).
	LinkProtection() Ospfv2InterfaceLinkProtection
	// SetLinkProtection assigns Ospfv2InterfaceLinkProtection provided by user to Ospfv2Interface.
	// Ospfv2InterfaceLinkProtection is optional container for the link protection sub TLV (type 20).
	SetLinkProtection(value Ospfv2InterfaceLinkProtection) Ospfv2Interface
	// HasLinkProtection checks if LinkProtection has been set in Ospfv2Interface
	HasLinkProtection() bool
	// SrlgValues returns []uint32, set in Ospfv2Interface.
	SrlgValues() []uint32
	// SetSrlgValues assigns []uint32 provided by user to Ospfv2Interface
	SetSrlgValues(value []uint32) Ospfv2Interface
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *ospfv2Interface) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetName(value string) Ospfv2Interface {

	obj.obj.Name = &value
	return obj
}

type Ospfv2InterfaceChoiceEnum string

// Enum of Choice on Ospfv2Interface
var Ospfv2InterfaceChoice = struct {
	IPV4_NAME      Ospfv2InterfaceChoiceEnum
	DHCP_INTF_NAME Ospfv2InterfaceChoiceEnum
}{
	IPV4_NAME:      Ospfv2InterfaceChoiceEnum("ipv4_name"),
	DHCP_INTF_NAME: Ospfv2InterfaceChoiceEnum("dhcp_intf_name"),
}

func (obj *ospfv2Interface) Choice() Ospfv2InterfaceChoiceEnum {
	return Ospfv2InterfaceChoiceEnum(obj.obj.Choice.Enum().String())
}

// The name of the interface top of which OSPFv2 is configured.
// - ipv4_name: The globally unique name of the IPv4 interface connected to the DUT.
// This name must match the "name" field of the "ipv4_addresses" on top which this OSPFv2 interface is configured.
// - dhcp_intf_name: The Area ID in IPv4 address format.
// Choice returns a string
func (obj *ospfv2Interface) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv2Interface) setChoice(value Ospfv2InterfaceChoiceEnum) Ospfv2Interface {
	intValue, ok := otg.Ospfv2Interface_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2InterfaceChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2Interface_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.DhcpIntfName = nil
	obj.obj.Ipv4Name = nil
	return obj
}

// The globally unique name of the IPv4 interface connected to the DUT.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// Ipv4Name returns a string
func (obj *ospfv2Interface) Ipv4Name() string {

	if obj.obj.Ipv4Name == nil {
		obj.setChoice(Ospfv2InterfaceChoice.IPV4_NAME)
	}

	return *obj.obj.Ipv4Name

}

// The globally unique name of the IPv4 interface connected to the DUT.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// Ipv4Name returns a string
func (obj *ospfv2Interface) HasIpv4Name() bool {
	return obj.obj.Ipv4Name != nil
}

// The globally unique name of the IPv4 interface connected to the DUT.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// SetIpv4Name sets the string value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetIpv4Name(value string) Ospfv2Interface {
	obj.setChoice(Ospfv2InterfaceChoice.IPV4_NAME)
	obj.obj.Ipv4Name = &value
	return obj
}

// The globally unique name of the DHCPv4 interface connected to the DUT.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// DhcpIntfName returns a string
func (obj *ospfv2Interface) DhcpIntfName() string {

	if obj.obj.DhcpIntfName == nil {
		obj.setChoice(Ospfv2InterfaceChoice.DHCP_INTF_NAME)
	}

	return *obj.obj.DhcpIntfName

}

// The globally unique name of the DHCPv4 interface connected to the DUT.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// DhcpIntfName returns a string
func (obj *ospfv2Interface) HasDhcpIntfName() bool {
	return obj.obj.DhcpIntfName != nil
}

// The globally unique name of the DHCPv4 interface connected to the DUT.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// SetDhcpIntfName sets the string value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetDhcpIntfName(value string) Ospfv2Interface {
	obj.setChoice(Ospfv2InterfaceChoice.DHCP_INTF_NAME)
	obj.obj.DhcpIntfName = &value
	return obj
}

type Ospfv2InterfaceAreaTypeEnum string

// Enum of AreaType on Ospfv2Interface
var Ospfv2InterfaceAreaType = struct {
	AREA_ID       Ospfv2InterfaceAreaTypeEnum
	AREA_ID_AS_IP Ospfv2InterfaceAreaTypeEnum
}{
	AREA_ID:       Ospfv2InterfaceAreaTypeEnum("area_id"),
	AREA_ID_AS_IP: Ospfv2InterfaceAreaTypeEnum("area_id_as_ip"),
}

func (obj *ospfv2Interface) AreaType() Ospfv2InterfaceAreaTypeEnum {
	return Ospfv2InterfaceAreaTypeEnum(obj.obj.AreaType.Enum().String())
}

// The OSPF Area ID identifies the routing area to which the host belongs. Area ID type can be following format.
// - area_id: A 32-bit number identifying the area.
// - area_id_as_ip: The Area ID in IPv4 address format.
// AreaType returns a string
func (obj *ospfv2Interface) HasAreaType() bool {
	return obj.obj.AreaType != nil
}

func (obj *ospfv2Interface) SetAreaType(value Ospfv2InterfaceAreaTypeEnum) Ospfv2Interface {
	intValue, ok := otg.Ospfv2Interface_AreaType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2InterfaceAreaTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2Interface_AreaType_Enum(intValue)
	obj.obj.AreaType = &enumValue

	return obj
}

// The Area ID.
// AreaId returns a uint32
func (obj *ospfv2Interface) AreaId() uint32 {

	return *obj.obj.AreaId

}

// The Area ID.
// AreaId returns a uint32
func (obj *ospfv2Interface) HasAreaId() bool {
	return obj.obj.AreaId != nil
}

// The Area ID.
// SetAreaId sets the uint32 value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetAreaId(value uint32) Ospfv2Interface {

	obj.obj.AreaId = &value
	return obj
}

// The Area ID in IPv4 address format.
// AreaIdAsIp returns a string
func (obj *ospfv2Interface) AreaIdAsIp() string {

	return *obj.obj.AreaIdAsIp

}

// The Area ID in IPv4 address format.
// AreaIdAsIp returns a string
func (obj *ospfv2Interface) HasAreaIdAsIp() bool {
	return obj.obj.AreaIdAsIp != nil
}

// The Area ID in IPv4 address format.
// SetAreaIdAsIp sets the string value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetAreaIdAsIp(value string) Ospfv2Interface {

	obj.obj.AreaIdAsIp = &value
	return obj
}

type Ospfv2InterfaceNetworkTypeEnum string

// Enum of NetworkType on Ospfv2Interface
var Ospfv2InterfaceNetworkType = struct {
	BROADCAST          Ospfv2InterfaceNetworkTypeEnum
	POINT_TO_POINT     Ospfv2InterfaceNetworkTypeEnum
	POINT_TO_MULTICAST Ospfv2InterfaceNetworkTypeEnum
}{
	BROADCAST:          Ospfv2InterfaceNetworkTypeEnum("broadcast"),
	POINT_TO_POINT:     Ospfv2InterfaceNetworkTypeEnum("point_to_point"),
	POINT_TO_MULTICAST: Ospfv2InterfaceNetworkTypeEnum("point_to_multicast"),
}

func (obj *ospfv2Interface) NetworkType() Ospfv2InterfaceNetworkTypeEnum {
	return Ospfv2InterfaceNetworkTypeEnum(obj.obj.NetworkType.Enum().String())
}

// The OSPF network link type options.
// - Point to Point:
// - Broadcast:
// - Point to Multipoint: In this case, at least a neigbor to be configured.
// NetworkType returns a string
func (obj *ospfv2Interface) HasNetworkType() bool {
	return obj.obj.NetworkType != nil
}

func (obj *ospfv2Interface) SetNetworkType(value Ospfv2InterfaceNetworkTypeEnum) Ospfv2Interface {
	intValue, ok := otg.Ospfv2Interface_NetworkType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2InterfaceNetworkTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2Interface_NetworkType_Enum(intValue)
	obj.obj.NetworkType = &enumValue

	return obj
}

// Configuration of the list of neighbors.
// Neighbors returns a []Ospfv2InterfaceNeighbor
func (obj *ospfv2Interface) Neighbors() Ospfv2InterfaceOspfv2InterfaceNeighborIter {
	if len(obj.obj.Neighbors) == 0 {
		obj.obj.Neighbors = []*otg.Ospfv2InterfaceNeighbor{}
	}
	if obj.neighborsHolder == nil {
		obj.neighborsHolder = newOspfv2InterfaceOspfv2InterfaceNeighborIter(&obj.obj.Neighbors).setMsg(obj)
	}
	return obj.neighborsHolder
}

type ospfv2InterfaceOspfv2InterfaceNeighborIter struct {
	obj                          *ospfv2Interface
	ospfv2InterfaceNeighborSlice []Ospfv2InterfaceNeighbor
	fieldPtr                     *[]*otg.Ospfv2InterfaceNeighbor
}

func newOspfv2InterfaceOspfv2InterfaceNeighborIter(ptr *[]*otg.Ospfv2InterfaceNeighbor) Ospfv2InterfaceOspfv2InterfaceNeighborIter {
	return &ospfv2InterfaceOspfv2InterfaceNeighborIter{fieldPtr: ptr}
}

type Ospfv2InterfaceOspfv2InterfaceNeighborIter interface {
	setMsg(*ospfv2Interface) Ospfv2InterfaceOspfv2InterfaceNeighborIter
	Items() []Ospfv2InterfaceNeighbor
	Add() Ospfv2InterfaceNeighbor
	Append(items ...Ospfv2InterfaceNeighbor) Ospfv2InterfaceOspfv2InterfaceNeighborIter
	Set(index int, newObj Ospfv2InterfaceNeighbor) Ospfv2InterfaceOspfv2InterfaceNeighborIter
	Clear() Ospfv2InterfaceOspfv2InterfaceNeighborIter
	clearHolderSlice() Ospfv2InterfaceOspfv2InterfaceNeighborIter
	appendHolderSlice(item Ospfv2InterfaceNeighbor) Ospfv2InterfaceOspfv2InterfaceNeighborIter
}

func (obj *ospfv2InterfaceOspfv2InterfaceNeighborIter) setMsg(msg *ospfv2Interface) Ospfv2InterfaceOspfv2InterfaceNeighborIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2InterfaceNeighbor{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2InterfaceOspfv2InterfaceNeighborIter) Items() []Ospfv2InterfaceNeighbor {
	return obj.ospfv2InterfaceNeighborSlice
}

func (obj *ospfv2InterfaceOspfv2InterfaceNeighborIter) Add() Ospfv2InterfaceNeighbor {
	newObj := &otg.Ospfv2InterfaceNeighbor{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2InterfaceNeighbor{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2InterfaceNeighborSlice = append(obj.ospfv2InterfaceNeighborSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2InterfaceOspfv2InterfaceNeighborIter) Append(items ...Ospfv2InterfaceNeighbor) Ospfv2InterfaceOspfv2InterfaceNeighborIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2InterfaceNeighborSlice = append(obj.ospfv2InterfaceNeighborSlice, item)
	}
	return obj
}

func (obj *ospfv2InterfaceOspfv2InterfaceNeighborIter) Set(index int, newObj Ospfv2InterfaceNeighbor) Ospfv2InterfaceOspfv2InterfaceNeighborIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2InterfaceNeighborSlice[index] = newObj
	return obj
}
func (obj *ospfv2InterfaceOspfv2InterfaceNeighborIter) Clear() Ospfv2InterfaceOspfv2InterfaceNeighborIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2InterfaceNeighbor{}
		obj.ospfv2InterfaceNeighborSlice = []Ospfv2InterfaceNeighbor{}
	}
	return obj
}
func (obj *ospfv2InterfaceOspfv2InterfaceNeighborIter) clearHolderSlice() Ospfv2InterfaceOspfv2InterfaceNeighborIter {
	if len(obj.ospfv2InterfaceNeighborSlice) > 0 {
		obj.ospfv2InterfaceNeighborSlice = []Ospfv2InterfaceNeighbor{}
	}
	return obj
}
func (obj *ospfv2InterfaceOspfv2InterfaceNeighborIter) appendHolderSlice(item Ospfv2InterfaceNeighbor) Ospfv2InterfaceOspfv2InterfaceNeighborIter {
	obj.ospfv2InterfaceNeighborSlice = append(obj.ospfv2InterfaceNeighborSlice, item)
	return obj
}

// Contains a list of Traffic Engineering attributes.
// TrafficEngineering returns a []LinkStateTE
func (obj *ospfv2Interface) TrafficEngineering() Ospfv2InterfaceLinkStateTEIter {
	if len(obj.obj.TrafficEngineering) == 0 {
		obj.obj.TrafficEngineering = []*otg.LinkStateTE{}
	}
	if obj.trafficEngineeringHolder == nil {
		obj.trafficEngineeringHolder = newOspfv2InterfaceLinkStateTEIter(&obj.obj.TrafficEngineering).setMsg(obj)
	}
	return obj.trafficEngineeringHolder
}

type ospfv2InterfaceLinkStateTEIter struct {
	obj              *ospfv2Interface
	linkStateTESlice []LinkStateTE
	fieldPtr         *[]*otg.LinkStateTE
}

func newOspfv2InterfaceLinkStateTEIter(ptr *[]*otg.LinkStateTE) Ospfv2InterfaceLinkStateTEIter {
	return &ospfv2InterfaceLinkStateTEIter{fieldPtr: ptr}
}

type Ospfv2InterfaceLinkStateTEIter interface {
	setMsg(*ospfv2Interface) Ospfv2InterfaceLinkStateTEIter
	Items() []LinkStateTE
	Add() LinkStateTE
	Append(items ...LinkStateTE) Ospfv2InterfaceLinkStateTEIter
	Set(index int, newObj LinkStateTE) Ospfv2InterfaceLinkStateTEIter
	Clear() Ospfv2InterfaceLinkStateTEIter
	clearHolderSlice() Ospfv2InterfaceLinkStateTEIter
	appendHolderSlice(item LinkStateTE) Ospfv2InterfaceLinkStateTEIter
}

func (obj *ospfv2InterfaceLinkStateTEIter) setMsg(msg *ospfv2Interface) Ospfv2InterfaceLinkStateTEIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&linkStateTE{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2InterfaceLinkStateTEIter) Items() []LinkStateTE {
	return obj.linkStateTESlice
}

func (obj *ospfv2InterfaceLinkStateTEIter) Add() LinkStateTE {
	newObj := &otg.LinkStateTE{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &linkStateTE{obj: newObj}
	newLibObj.setDefault()
	obj.linkStateTESlice = append(obj.linkStateTESlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2InterfaceLinkStateTEIter) Append(items ...LinkStateTE) Ospfv2InterfaceLinkStateTEIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.linkStateTESlice = append(obj.linkStateTESlice, item)
	}
	return obj
}

func (obj *ospfv2InterfaceLinkStateTEIter) Set(index int, newObj LinkStateTE) Ospfv2InterfaceLinkStateTEIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.linkStateTESlice[index] = newObj
	return obj
}
func (obj *ospfv2InterfaceLinkStateTEIter) Clear() Ospfv2InterfaceLinkStateTEIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LinkStateTE{}
		obj.linkStateTESlice = []LinkStateTE{}
	}
	return obj
}
func (obj *ospfv2InterfaceLinkStateTEIter) clearHolderSlice() Ospfv2InterfaceLinkStateTEIter {
	if len(obj.linkStateTESlice) > 0 {
		obj.linkStateTESlice = []LinkStateTE{}
	}
	return obj
}
func (obj *ospfv2InterfaceLinkStateTEIter) appendHolderSlice(item LinkStateTE) Ospfv2InterfaceLinkStateTEIter {
	obj.linkStateTESlice = append(obj.linkStateTESlice, item)
	return obj
}

// OSPFv2 authentication properties.
// Authentication returns a Ospfv2InterfaceAuthentication
func (obj *ospfv2Interface) Authentication() Ospfv2InterfaceAuthentication {
	if obj.obj.Authentication == nil {
		obj.obj.Authentication = NewOspfv2InterfaceAuthentication().msg()
	}
	if obj.authenticationHolder == nil {
		obj.authenticationHolder = &ospfv2InterfaceAuthentication{obj: obj.obj.Authentication}
	}
	return obj.authenticationHolder
}

// OSPFv2 authentication properties.
// Authentication returns a Ospfv2InterfaceAuthentication
func (obj *ospfv2Interface) HasAuthentication() bool {
	return obj.obj.Authentication != nil
}

// OSPFv2 authentication properties.
// SetAuthentication sets the Ospfv2InterfaceAuthentication value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetAuthentication(value Ospfv2InterfaceAuthentication) Ospfv2Interface {

	obj.authenticationHolder = nil
	obj.obj.Authentication = value.msg()

	return obj
}

// Optional container for advanced interface properties.
// Advanced returns a Ospfv2InterfaceAdvanced
func (obj *ospfv2Interface) Advanced() Ospfv2InterfaceAdvanced {
	if obj.obj.Advanced == nil {
		obj.obj.Advanced = NewOspfv2InterfaceAdvanced().msg()
	}
	if obj.advancedHolder == nil {
		obj.advancedHolder = &ospfv2InterfaceAdvanced{obj: obj.obj.Advanced}
	}
	return obj.advancedHolder
}

// Optional container for advanced interface properties.
// Advanced returns a Ospfv2InterfaceAdvanced
func (obj *ospfv2Interface) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// Optional container for advanced interface properties.
// SetAdvanced sets the Ospfv2InterfaceAdvanced value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetAdvanced(value Ospfv2InterfaceAdvanced) Ospfv2Interface {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// Link protection on the OSPFv2 link between two interfaces.
// LinkProtection returns a Ospfv2InterfaceLinkProtection
func (obj *ospfv2Interface) LinkProtection() Ospfv2InterfaceLinkProtection {
	if obj.obj.LinkProtection == nil {
		obj.obj.LinkProtection = NewOspfv2InterfaceLinkProtection().msg()
	}
	if obj.linkProtectionHolder == nil {
		obj.linkProtectionHolder = &ospfv2InterfaceLinkProtection{obj: obj.obj.LinkProtection}
	}
	return obj.linkProtectionHolder
}

// Link protection on the OSPFv2 link between two interfaces.
// LinkProtection returns a Ospfv2InterfaceLinkProtection
func (obj *ospfv2Interface) HasLinkProtection() bool {
	return obj.obj.LinkProtection != nil
}

// Link protection on the OSPFv2 link between two interfaces.
// SetLinkProtection sets the Ospfv2InterfaceLinkProtection value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetLinkProtection(value Ospfv2InterfaceLinkProtection) Ospfv2Interface {

	obj.linkProtectionHolder = nil
	obj.obj.LinkProtection = value.msg()

	return obj
}

// This contains list of SRLG values for the link between two interfaces.
// SrlgValues returns a []uint32
func (obj *ospfv2Interface) SrlgValues() []uint32 {
	if obj.obj.SrlgValues == nil {
		obj.obj.SrlgValues = make([]uint32, 0)
	}
	return obj.obj.SrlgValues
}

// This contains list of SRLG values for the link between two interfaces.
// SetSrlgValues sets the []uint32 value in the Ospfv2Interface object
func (obj *ospfv2Interface) SetSrlgValues(value []uint32) Ospfv2Interface {

	if obj.obj.SrlgValues == nil {
		obj.obj.SrlgValues = make([]uint32, 0)
	}
	obj.obj.SrlgValues = value

	return obj
}

func (obj *ospfv2Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Ospfv2Interface")
	}

	if obj.obj.AreaIdAsIp != nil {

		err := obj.validateIpv4(obj.AreaIdAsIp())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2Interface.AreaIdAsIp"))
		}

	}

	if len(obj.obj.Neighbors) != 0 {

		if set_default {
			obj.Neighbors().clearHolderSlice()
			for _, item := range obj.obj.Neighbors {
				obj.Neighbors().appendHolderSlice(&ospfv2InterfaceNeighbor{obj: item})
			}
		}
		for _, item := range obj.Neighbors().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.TrafficEngineering) != 0 {

		if set_default {
			obj.TrafficEngineering().clearHolderSlice()
			for _, item := range obj.obj.TrafficEngineering {
				obj.TrafficEngineering().appendHolderSlice(&linkStateTE{obj: item})
			}
		}
		for _, item := range obj.TrafficEngineering().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Authentication != nil {

		obj.Authentication().validateObj(vObj, set_default)
	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(vObj, set_default)
	}

	if obj.obj.LinkProtection != nil {

		obj.LinkProtection().validateObj(vObj, set_default)
	}

	if obj.obj.SrlgValues != nil {

		for _, item := range obj.obj.SrlgValues {
			if item > 5 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= Ospfv2Interface.SrlgValues <= 5 but Got %d", item))
			}

		}

	}

}

func (obj *ospfv2Interface) setDefault() {
	var choices_set int = 0
	var choice Ospfv2InterfaceChoiceEnum

	if obj.obj.Ipv4Name != nil {
		choices_set += 1
		choice = Ospfv2InterfaceChoice.IPV4_NAME
	}

	if obj.obj.DhcpIntfName != nil {
		choices_set += 1
		choice = Ospfv2InterfaceChoice.DHCP_INTF_NAME
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv2InterfaceChoice.IPV4_NAME)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv2Interface")
			}
		} else {
			intVal := otg.Ospfv2Interface_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv2Interface_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.AreaType == nil {
		obj.SetAreaType(Ospfv2InterfaceAreaType.AREA_ID)

	}
	if obj.obj.AreaId == nil {
		obj.SetAreaId(0)
	}
	if obj.obj.NetworkType == nil {
		obj.SetNetworkType(Ospfv2InterfaceNetworkType.BROADCAST)

	}

}
