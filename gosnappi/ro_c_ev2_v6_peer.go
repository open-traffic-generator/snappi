package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2V6Peer *****
type roCEv2V6Peer struct {
	validation
	obj                *otg.RoCEv2V6Peer
	marshaller         marshalRoCEv2V6Peer
	unMarshaller       unMarshalRoCEv2V6Peer
	flowSettingsHolder RoCEv2V6PeerRoCEv2FlowSettingsIter
}

func NewRoCEv2V6Peer() RoCEv2V6Peer {
	obj := roCEv2V6Peer{obj: &otg.RoCEv2V6Peer{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2V6Peer) msg() *otg.RoCEv2V6Peer {
	return obj.obj
}

func (obj *roCEv2V6Peer) setMsg(msg *otg.RoCEv2V6Peer) RoCEv2V6Peer {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2V6Peer struct {
	obj *roCEv2V6Peer
}

type marshalRoCEv2V6Peer interface {
	// ToProto marshals RoCEv2V6Peer to protobuf object *otg.RoCEv2V6Peer
	ToProto() (*otg.RoCEv2V6Peer, error)
	// ToPbText marshals RoCEv2V6Peer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2V6Peer to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2V6Peer to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2V6Peer struct {
	obj *roCEv2V6Peer
}

type unMarshalRoCEv2V6Peer interface {
	// FromProto unmarshals RoCEv2V6Peer from protobuf object *otg.RoCEv2V6Peer
	FromProto(msg *otg.RoCEv2V6Peer) (RoCEv2V6Peer, error)
	// FromPbText unmarshals RoCEv2V6Peer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2V6Peer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2V6Peer from JSON text
	FromJson(value string) error
}

func (obj *roCEv2V6Peer) Marshal() marshalRoCEv2V6Peer {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2V6Peer{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2V6Peer) Unmarshal() unMarshalRoCEv2V6Peer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2V6Peer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2V6Peer) ToProto() (*otg.RoCEv2V6Peer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2V6Peer) FromProto(msg *otg.RoCEv2V6Peer) (RoCEv2V6Peer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2V6Peer) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2V6Peer) FromPbText(value string) error {
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

func (m *marshalroCEv2V6Peer) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2V6Peer) FromYaml(value string) error {
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

func (m *marshalroCEv2V6Peer) ToJson() (string, error) {
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

func (m *unMarshalroCEv2V6Peer) FromJson(value string) error {
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

func (obj *roCEv2V6Peer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2V6Peer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2V6Peer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2V6Peer) Clone() (RoCEv2V6Peer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2V6Peer()
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

func (obj *roCEv2V6Peer) setNil() {
	obj.flowSettingsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RoCEv2V6Peer is configuration for RoCEv2 IPv6 peer settings.
type RoCEv2V6Peer interface {
	Validation
	// msg marshals RoCEv2V6Peer to protobuf object *otg.RoCEv2V6Peer
	// and doesn't set defaults
	msg() *otg.RoCEv2V6Peer
	// setMsg unmarshals RoCEv2V6Peer from protobuf object *otg.RoCEv2V6Peer
	// and doesn't set defaults
	setMsg(*otg.RoCEv2V6Peer) RoCEv2V6Peer
	// provides marshal interface
	Marshal() marshalRoCEv2V6Peer
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2V6Peer
	// validate validates RoCEv2V6Peer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2V6Peer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// NumOfQps returns int32, set in RoCEv2V6Peer.
	NumOfQps() int32
	// SetNumOfQps assigns int32 provided by user to RoCEv2V6Peer
	SetNumOfQps(value int32) RoCEv2V6Peer
	// DestinationIpAddress returns string, set in RoCEv2V6Peer.
	DestinationIpAddress() string
	// SetDestinationIpAddress assigns string provided by user to RoCEv2V6Peer
	SetDestinationIpAddress(value string) RoCEv2V6Peer
	// IbMtu returns uint32, set in RoCEv2V6Peer.
	IbMtu() uint32
	// SetIbMtu assigns uint32 provided by user to RoCEv2V6Peer
	SetIbMtu(value uint32) RoCEv2V6Peer
	// ConnectionType returns RoCEv2V6PeerConnectionTypeEnum, set in RoCEv2V6Peer
	ConnectionType() RoCEv2V6PeerConnectionTypeEnum
	// SetConnectionType assigns RoCEv2V6PeerConnectionTypeEnum provided by user to RoCEv2V6Peer
	SetConnectionType(value RoCEv2V6PeerConnectionTypeEnum) RoCEv2V6Peer
	// HasConnectionType checks if ConnectionType has been set in RoCEv2V6Peer
	HasConnectionType() bool
	// TrafficBurstMode returns RoCEv2V6PeerTrafficBurstModeEnum, set in RoCEv2V6Peer
	TrafficBurstMode() RoCEv2V6PeerTrafficBurstModeEnum
	// SetTrafficBurstMode assigns RoCEv2V6PeerTrafficBurstModeEnum provided by user to RoCEv2V6Peer
	SetTrafficBurstMode(value RoCEv2V6PeerTrafficBurstModeEnum) RoCEv2V6Peer
	// HasTrafficBurstMode checks if TrafficBurstMode has been set in RoCEv2V6Peer
	HasTrafficBurstMode() bool
	// TrafficBurstCount returns uint32, set in RoCEv2V6Peer.
	TrafficBurstCount() uint32
	// SetTrafficBurstCount assigns uint32 provided by user to RoCEv2V6Peer
	SetTrafficBurstCount(value uint32) RoCEv2V6Peer
	// HasTrafficBurstCount checks if TrafficBurstCount has been set in RoCEv2V6Peer
	HasTrafficBurstCount() bool
	// Name returns string, set in RoCEv2V6Peer.
	Name() string
	// SetName assigns string provided by user to RoCEv2V6Peer
	SetName(value string) RoCEv2V6Peer
	// HasName checks if Name has been set in RoCEv2V6Peer
	HasName() bool
	// FlowSettings returns RoCEv2V6PeerRoCEv2FlowSettingsIterIter, set in RoCEv2V6Peer
	FlowSettings() RoCEv2V6PeerRoCEv2FlowSettingsIter
	setNil()
}

// Maximum number of QP per Endpoint.
// NumOfQps returns a int32
func (obj *roCEv2V6Peer) NumOfQps() int32 {

	return *obj.obj.NumOfQps

}

// Maximum number of QP per Endpoint.
// SetNumOfQps sets the int32 value in the RoCEv2V6Peer object
func (obj *roCEv2V6Peer) SetNumOfQps(value int32) RoCEv2V6Peer {

	obj.obj.NumOfQps = &value
	return obj
}

// Specify the IP address of External NIC i.e Destination Point IP Address.
// DestinationIpAddress returns a string
func (obj *roCEv2V6Peer) DestinationIpAddress() string {

	return *obj.obj.DestinationIpAddress

}

// Specify the IP address of External NIC i.e Destination Point IP Address.
// SetDestinationIpAddress sets the string value in the RoCEv2V6Peer object
func (obj *roCEv2V6Peer) SetDestinationIpAddress(value string) RoCEv2V6Peer {

	obj.obj.DestinationIpAddress = &value
	return obj
}

// InfiniBand protocol Maximum Transmission Unit (MTU) defines several fix size MTU: 256, 512, 1024, 2048 or 4096 bytes. RDMA write message will have payload size same as configured IB MTU. You can configure custom size also.
// IbMtu returns a uint32
func (obj *roCEv2V6Peer) IbMtu() uint32 {

	return *obj.obj.IbMtu

}

// InfiniBand protocol Maximum Transmission Unit (MTU) defines several fix size MTU: 256, 512, 1024, 2048 or 4096 bytes. RDMA write message will have payload size same as configured IB MTU. You can configure custom size also.
// SetIbMtu sets the uint32 value in the RoCEv2V6Peer object
func (obj *roCEv2V6Peer) SetIbMtu(value uint32) RoCEv2V6Peer {

	obj.obj.IbMtu = &value
	return obj
}

type RoCEv2V6PeerConnectionTypeEnum string

// Enum of ConnectionType on RoCEv2V6Peer
var RoCEv2V6PeerConnectionType = struct {
	RELIABLE_CONNECTION          RoCEv2V6PeerConnectionTypeEnum
	RELIABLE_DATAGRAM            RoCEv2V6PeerConnectionTypeEnum
	EXTENDED_RELIABLE_CONNECTION RoCEv2V6PeerConnectionTypeEnum
	UNRELIABLE_DATAGRAM          RoCEv2V6PeerConnectionTypeEnum
	UNRELIABLE_CONNECTION        RoCEv2V6PeerConnectionTypeEnum
	RAW_IPV6_DATAGRAM            RoCEv2V6PeerConnectionTypeEnum
	RAW_ETHERNET_DATAGRAM        RoCEv2V6PeerConnectionTypeEnum
}{
	RELIABLE_CONNECTION:          RoCEv2V6PeerConnectionTypeEnum("reliable_connection"),
	RELIABLE_DATAGRAM:            RoCEv2V6PeerConnectionTypeEnum("reliable_datagram"),
	EXTENDED_RELIABLE_CONNECTION: RoCEv2V6PeerConnectionTypeEnum("extended_reliable_connection"),
	UNRELIABLE_DATAGRAM:          RoCEv2V6PeerConnectionTypeEnum("unreliable_datagram"),
	UNRELIABLE_CONNECTION:        RoCEv2V6PeerConnectionTypeEnum("unreliable_connection"),
	RAW_IPV6_DATAGRAM:            RoCEv2V6PeerConnectionTypeEnum("raw_ipv6_datagram"),
	RAW_ETHERNET_DATAGRAM:        RoCEv2V6PeerConnectionTypeEnum("raw_ethernet_datagram"),
}

func (obj *roCEv2V6Peer) ConnectionType() RoCEv2V6PeerConnectionTypeEnum {
	return RoCEv2V6PeerConnectionTypeEnum(obj.obj.ConnectionType.Enum().String())
}

// There are multiple connection types. Valid values are :  Reliable Connection (RC), Reliable Datagram (RD), Extended Reliable Connection (XRC), Unreliable Datagram (UD),  Unreliable Connection (UC), Raw IPv6 Datagram, Raw Ethertype Datagram.
// ConnectionType returns a string
func (obj *roCEv2V6Peer) HasConnectionType() bool {
	return obj.obj.ConnectionType != nil
}

func (obj *roCEv2V6Peer) SetConnectionType(value RoCEv2V6PeerConnectionTypeEnum) RoCEv2V6Peer {
	intValue, ok := otg.RoCEv2V6Peer_ConnectionType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2V6PeerConnectionTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2V6Peer_ConnectionType_Enum(intValue)
	obj.obj.ConnectionType = &enumValue

	return obj
}

type RoCEv2V6PeerTrafficBurstModeEnum string

// Enum of TrafficBurstMode on RoCEv2V6Peer
var RoCEv2V6PeerTrafficBurstMode = struct {
	CONTINUOUS RoCEv2V6PeerTrafficBurstModeEnum
	FIXED      RoCEv2V6PeerTrafficBurstModeEnum
}{
	CONTINUOUS: RoCEv2V6PeerTrafficBurstModeEnum("Continuous"),
	FIXED:      RoCEv2V6PeerTrafficBurstModeEnum("Fixed"),
}

func (obj *roCEv2V6Peer) TrafficBurstMode() RoCEv2V6PeerTrafficBurstModeEnum {
	return RoCEv2V6PeerTrafficBurstModeEnum(obj.obj.TrafficBurstMode.Enum().String())
}

// Traffic Burst Mode to applied in RoCEv2 Traffic , Continuous or Fixed.
// TrafficBurstMode returns a string
func (obj *roCEv2V6Peer) HasTrafficBurstMode() bool {
	return obj.obj.TrafficBurstMode != nil
}

func (obj *roCEv2V6Peer) SetTrafficBurstMode(value RoCEv2V6PeerTrafficBurstModeEnum) RoCEv2V6Peer {
	intValue, ok := otg.RoCEv2V6Peer_TrafficBurstMode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2V6PeerTrafficBurstModeEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2V6Peer_TrafficBurstMode_Enum(intValue)
	obj.obj.TrafficBurstMode = &enumValue

	return obj
}

// Burst Count to applied in RoCEv2 Traffic item.
// TrafficBurstCount returns a uint32
func (obj *roCEv2V6Peer) TrafficBurstCount() uint32 {

	return *obj.obj.TrafficBurstCount

}

// Burst Count to applied in RoCEv2 Traffic item.
// TrafficBurstCount returns a uint32
func (obj *roCEv2V6Peer) HasTrafficBurstCount() bool {
	return obj.obj.TrafficBurstCount != nil
}

// Burst Count to applied in RoCEv2 Traffic item.
// SetTrafficBurstCount sets the uint32 value in the RoCEv2V6Peer object
func (obj *roCEv2V6Peer) SetTrafficBurstCount(value uint32) RoCEv2V6Peer {

	obj.obj.TrafficBurstCount = &value
	return obj
}

// The name of RoCEv2 Peer
// Name returns a string
func (obj *roCEv2V6Peer) Name() string {

	return *obj.obj.Name

}

// The name of RoCEv2 Peer
// Name returns a string
func (obj *roCEv2V6Peer) HasName() bool {
	return obj.obj.Name != nil
}

// The name of RoCEv2 Peer
// SetName sets the string value in the RoCEv2V6Peer object
func (obj *roCEv2V6Peer) SetName(value string) RoCEv2V6Peer {

	obj.obj.Name = &value
	return obj
}

// This configuration allows you to configure RDMA flow over the same QP number from same source and destination.
// FlowSettings returns a []RoCEv2FlowSettings
func (obj *roCEv2V6Peer) FlowSettings() RoCEv2V6PeerRoCEv2FlowSettingsIter {
	if len(obj.obj.FlowSettings) == 0 {
		obj.obj.FlowSettings = []*otg.RoCEv2FlowSettings{}
	}
	if obj.flowSettingsHolder == nil {
		obj.flowSettingsHolder = newRoCEv2V6PeerRoCEv2FlowSettingsIter(&obj.obj.FlowSettings).setMsg(obj)
	}
	return obj.flowSettingsHolder
}

type roCEv2V6PeerRoCEv2FlowSettingsIter struct {
	obj                     *roCEv2V6Peer
	roCEv2FlowSettingsSlice []RoCEv2FlowSettings
	fieldPtr                *[]*otg.RoCEv2FlowSettings
}

func newRoCEv2V6PeerRoCEv2FlowSettingsIter(ptr *[]*otg.RoCEv2FlowSettings) RoCEv2V6PeerRoCEv2FlowSettingsIter {
	return &roCEv2V6PeerRoCEv2FlowSettingsIter{fieldPtr: ptr}
}

type RoCEv2V6PeerRoCEv2FlowSettingsIter interface {
	setMsg(*roCEv2V6Peer) RoCEv2V6PeerRoCEv2FlowSettingsIter
	Items() []RoCEv2FlowSettings
	Add() RoCEv2FlowSettings
	Append(items ...RoCEv2FlowSettings) RoCEv2V6PeerRoCEv2FlowSettingsIter
	Set(index int, newObj RoCEv2FlowSettings) RoCEv2V6PeerRoCEv2FlowSettingsIter
	Clear() RoCEv2V6PeerRoCEv2FlowSettingsIter
	clearHolderSlice() RoCEv2V6PeerRoCEv2FlowSettingsIter
	appendHolderSlice(item RoCEv2FlowSettings) RoCEv2V6PeerRoCEv2FlowSettingsIter
}

func (obj *roCEv2V6PeerRoCEv2FlowSettingsIter) setMsg(msg *roCEv2V6Peer) RoCEv2V6PeerRoCEv2FlowSettingsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&roCEv2FlowSettings{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *roCEv2V6PeerRoCEv2FlowSettingsIter) Items() []RoCEv2FlowSettings {
	return obj.roCEv2FlowSettingsSlice
}

func (obj *roCEv2V6PeerRoCEv2FlowSettingsIter) Add() RoCEv2FlowSettings {
	newObj := &otg.RoCEv2FlowSettings{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &roCEv2FlowSettings{obj: newObj}
	newLibObj.setDefault()
	obj.roCEv2FlowSettingsSlice = append(obj.roCEv2FlowSettingsSlice, newLibObj)
	return newLibObj
}

func (obj *roCEv2V6PeerRoCEv2FlowSettingsIter) Append(items ...RoCEv2FlowSettings) RoCEv2V6PeerRoCEv2FlowSettingsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.roCEv2FlowSettingsSlice = append(obj.roCEv2FlowSettingsSlice, item)
	}
	return obj
}

func (obj *roCEv2V6PeerRoCEv2FlowSettingsIter) Set(index int, newObj RoCEv2FlowSettings) RoCEv2V6PeerRoCEv2FlowSettingsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.roCEv2FlowSettingsSlice[index] = newObj
	return obj
}
func (obj *roCEv2V6PeerRoCEv2FlowSettingsIter) Clear() RoCEv2V6PeerRoCEv2FlowSettingsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RoCEv2FlowSettings{}
		obj.roCEv2FlowSettingsSlice = []RoCEv2FlowSettings{}
	}
	return obj
}
func (obj *roCEv2V6PeerRoCEv2FlowSettingsIter) clearHolderSlice() RoCEv2V6PeerRoCEv2FlowSettingsIter {
	if len(obj.roCEv2FlowSettingsSlice) > 0 {
		obj.roCEv2FlowSettingsSlice = []RoCEv2FlowSettings{}
	}
	return obj
}
func (obj *roCEv2V6PeerRoCEv2FlowSettingsIter) appendHolderSlice(item RoCEv2FlowSettings) RoCEv2V6PeerRoCEv2FlowSettingsIter {
	obj.roCEv2FlowSettingsSlice = append(obj.roCEv2FlowSettingsSlice, item)
	return obj
}

func (obj *roCEv2V6Peer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// NumOfQps is required
	if obj.obj.NumOfQps == nil {
		vObj.validationErrors = append(vObj.validationErrors, "NumOfQps is required field on interface RoCEv2V6Peer")
	}

	// DestinationIpAddress is required
	if obj.obj.DestinationIpAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "DestinationIpAddress is required field on interface RoCEv2V6Peer")
	}

	// IbMtu is required
	if obj.obj.IbMtu == nil {
		vObj.validationErrors = append(vObj.validationErrors, "IbMtu is required field on interface RoCEv2V6Peer")
	}
	if obj.obj.IbMtu != nil {

		if *obj.obj.IbMtu > 14000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2V6Peer.IbMtu <= 14000 but Got %d", *obj.obj.IbMtu))
		}

	}

	if obj.obj.TrafficBurstCount != nil {

		if *obj.obj.TrafficBurstCount > 16777216 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2V6Peer.TrafficBurstCount <= 16777216 but Got %d", *obj.obj.TrafficBurstCount))
		}

	}

	if len(obj.obj.FlowSettings) != 0 {

		if set_default {
			obj.FlowSettings().clearHolderSlice()
			for _, item := range obj.obj.FlowSettings {
				obj.FlowSettings().appendHolderSlice(&roCEv2FlowSettings{obj: item})
			}
		}
		for _, item := range obj.FlowSettings().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *roCEv2V6Peer) setDefault() {

}
