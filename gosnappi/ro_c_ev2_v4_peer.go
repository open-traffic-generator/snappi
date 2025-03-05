package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2V4Peer *****
type roCEv2V4Peer struct {
	validation
	obj                *otg.RoCEv2V4Peer
	marshaller         marshalRoCEv2V4Peer
	unMarshaller       unMarshalRoCEv2V4Peer
	flowSettingsHolder RoCEv2V4PeerRoCEv2FlowSettingsIter
}

func NewRoCEv2V4Peer() RoCEv2V4Peer {
	obj := roCEv2V4Peer{obj: &otg.RoCEv2V4Peer{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2V4Peer) msg() *otg.RoCEv2V4Peer {
	return obj.obj
}

func (obj *roCEv2V4Peer) setMsg(msg *otg.RoCEv2V4Peer) RoCEv2V4Peer {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2V4Peer struct {
	obj *roCEv2V4Peer
}

type marshalRoCEv2V4Peer interface {
	// ToProto marshals RoCEv2V4Peer to protobuf object *otg.RoCEv2V4Peer
	ToProto() (*otg.RoCEv2V4Peer, error)
	// ToPbText marshals RoCEv2V4Peer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2V4Peer to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2V4Peer to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2V4Peer struct {
	obj *roCEv2V4Peer
}

type unMarshalRoCEv2V4Peer interface {
	// FromProto unmarshals RoCEv2V4Peer from protobuf object *otg.RoCEv2V4Peer
	FromProto(msg *otg.RoCEv2V4Peer) (RoCEv2V4Peer, error)
	// FromPbText unmarshals RoCEv2V4Peer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2V4Peer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2V4Peer from JSON text
	FromJson(value string) error
}

func (obj *roCEv2V4Peer) Marshal() marshalRoCEv2V4Peer {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2V4Peer{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2V4Peer) Unmarshal() unMarshalRoCEv2V4Peer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2V4Peer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2V4Peer) ToProto() (*otg.RoCEv2V4Peer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2V4Peer) FromProto(msg *otg.RoCEv2V4Peer) (RoCEv2V4Peer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2V4Peer) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2V4Peer) FromPbText(value string) error {
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

func (m *marshalroCEv2V4Peer) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2V4Peer) FromYaml(value string) error {
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

func (m *marshalroCEv2V4Peer) ToJson() (string, error) {
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

func (m *unMarshalroCEv2V4Peer) FromJson(value string) error {
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

func (obj *roCEv2V4Peer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2V4Peer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2V4Peer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2V4Peer) Clone() (RoCEv2V4Peer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2V4Peer()
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

func (obj *roCEv2V4Peer) setNil() {
	obj.flowSettingsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RoCEv2V4Peer is configuration for RoCEv2 IPv4 peers.
type RoCEv2V4Peer interface {
	Validation
	// msg marshals RoCEv2V4Peer to protobuf object *otg.RoCEv2V4Peer
	// and doesn't set defaults
	msg() *otg.RoCEv2V4Peer
	// setMsg unmarshals RoCEv2V4Peer from protobuf object *otg.RoCEv2V4Peer
	// and doesn't set defaults
	setMsg(*otg.RoCEv2V4Peer) RoCEv2V4Peer
	// provides marshal interface
	Marshal() marshalRoCEv2V4Peer
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2V4Peer
	// validate validates RoCEv2V4Peer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2V4Peer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// NumOfQps returns uint32, set in RoCEv2V4Peer.
	NumOfQps() uint32
	// SetNumOfQps assigns uint32 provided by user to RoCEv2V4Peer
	SetNumOfQps(value uint32) RoCEv2V4Peer
	// DestinationIpAddress returns string, set in RoCEv2V4Peer.
	DestinationIpAddress() string
	// SetDestinationIpAddress assigns string provided by user to RoCEv2V4Peer
	SetDestinationIpAddress(value string) RoCEv2V4Peer
	// IbMtu returns uint32, set in RoCEv2V4Peer.
	IbMtu() uint32
	// SetIbMtu assigns uint32 provided by user to RoCEv2V4Peer
	SetIbMtu(value uint32) RoCEv2V4Peer
	// Name returns string, set in RoCEv2V4Peer.
	Name() string
	// SetName assigns string provided by user to RoCEv2V4Peer
	SetName(value string) RoCEv2V4Peer
	// HasName checks if Name has been set in RoCEv2V4Peer
	HasName() bool
	// FlowSettings returns RoCEv2V4PeerRoCEv2FlowSettingsIterIter, set in RoCEv2V4Peer
	FlowSettings() RoCEv2V4PeerRoCEv2FlowSettingsIter
	setNil()
}

// Maximum number of QP per Endpoint.
// NumOfQps returns a uint32
func (obj *roCEv2V4Peer) NumOfQps() uint32 {

	return *obj.obj.NumOfQps

}

// Maximum number of QP per Endpoint.
// SetNumOfQps sets the uint32 value in the RoCEv2V4Peer object
func (obj *roCEv2V4Peer) SetNumOfQps(value uint32) RoCEv2V4Peer {

	obj.obj.NumOfQps = &value
	return obj
}

// Specify the IP address of External NIC i.e Destination Point IP Address.
// DestinationIpAddress returns a string
func (obj *roCEv2V4Peer) DestinationIpAddress() string {

	return *obj.obj.DestinationIpAddress

}

// Specify the IP address of External NIC i.e Destination Point IP Address.
// SetDestinationIpAddress sets the string value in the RoCEv2V4Peer object
func (obj *roCEv2V4Peer) SetDestinationIpAddress(value string) RoCEv2V4Peer {

	obj.obj.DestinationIpAddress = &value
	return obj
}

// InfiniBand protocol Maximum Transmission Unit (MTU) defines several fix size MTU: 256, 512, 1024, 2048 or 4096 bytes. RDMA write message will have payload size same as configured IB MTU. You can configure custom size also.
// IbMtu returns a uint32
func (obj *roCEv2V4Peer) IbMtu() uint32 {

	return *obj.obj.IbMtu

}

// InfiniBand protocol Maximum Transmission Unit (MTU) defines several fix size MTU: 256, 512, 1024, 2048 or 4096 bytes. RDMA write message will have payload size same as configured IB MTU. You can configure custom size also.
// SetIbMtu sets the uint32 value in the RoCEv2V4Peer object
func (obj *roCEv2V4Peer) SetIbMtu(value uint32) RoCEv2V4Peer {

	obj.obj.IbMtu = &value
	return obj
}

// The name of RoCEv2 Peer
// Name returns a string
func (obj *roCEv2V4Peer) Name() string {

	return *obj.obj.Name

}

// The name of RoCEv2 Peer
// Name returns a string
func (obj *roCEv2V4Peer) HasName() bool {
	return obj.obj.Name != nil
}

// The name of RoCEv2 Peer
// SetName sets the string value in the RoCEv2V4Peer object
func (obj *roCEv2V4Peer) SetName(value string) RoCEv2V4Peer {

	obj.obj.Name = &value
	return obj
}

// This configuration allows you to configure RDMA flow over the same QP number from same source and destination.
// FlowSettings returns a []RoCEv2FlowSettings
func (obj *roCEv2V4Peer) FlowSettings() RoCEv2V4PeerRoCEv2FlowSettingsIter {
	if len(obj.obj.FlowSettings) == 0 {
		obj.obj.FlowSettings = []*otg.RoCEv2FlowSettings{}
	}
	if obj.flowSettingsHolder == nil {
		obj.flowSettingsHolder = newRoCEv2V4PeerRoCEv2FlowSettingsIter(&obj.obj.FlowSettings).setMsg(obj)
	}
	return obj.flowSettingsHolder
}

type roCEv2V4PeerRoCEv2FlowSettingsIter struct {
	obj                     *roCEv2V4Peer
	roCEv2FlowSettingsSlice []RoCEv2FlowSettings
	fieldPtr                *[]*otg.RoCEv2FlowSettings
}

func newRoCEv2V4PeerRoCEv2FlowSettingsIter(ptr *[]*otg.RoCEv2FlowSettings) RoCEv2V4PeerRoCEv2FlowSettingsIter {
	return &roCEv2V4PeerRoCEv2FlowSettingsIter{fieldPtr: ptr}
}

type RoCEv2V4PeerRoCEv2FlowSettingsIter interface {
	setMsg(*roCEv2V4Peer) RoCEv2V4PeerRoCEv2FlowSettingsIter
	Items() []RoCEv2FlowSettings
	Add() RoCEv2FlowSettings
	Append(items ...RoCEv2FlowSettings) RoCEv2V4PeerRoCEv2FlowSettingsIter
	Set(index int, newObj RoCEv2FlowSettings) RoCEv2V4PeerRoCEv2FlowSettingsIter
	Clear() RoCEv2V4PeerRoCEv2FlowSettingsIter
	clearHolderSlice() RoCEv2V4PeerRoCEv2FlowSettingsIter
	appendHolderSlice(item RoCEv2FlowSettings) RoCEv2V4PeerRoCEv2FlowSettingsIter
}

func (obj *roCEv2V4PeerRoCEv2FlowSettingsIter) setMsg(msg *roCEv2V4Peer) RoCEv2V4PeerRoCEv2FlowSettingsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&roCEv2FlowSettings{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *roCEv2V4PeerRoCEv2FlowSettingsIter) Items() []RoCEv2FlowSettings {
	return obj.roCEv2FlowSettingsSlice
}

func (obj *roCEv2V4PeerRoCEv2FlowSettingsIter) Add() RoCEv2FlowSettings {
	newObj := &otg.RoCEv2FlowSettings{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &roCEv2FlowSettings{obj: newObj}
	newLibObj.setDefault()
	obj.roCEv2FlowSettingsSlice = append(obj.roCEv2FlowSettingsSlice, newLibObj)
	return newLibObj
}

func (obj *roCEv2V4PeerRoCEv2FlowSettingsIter) Append(items ...RoCEv2FlowSettings) RoCEv2V4PeerRoCEv2FlowSettingsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.roCEv2FlowSettingsSlice = append(obj.roCEv2FlowSettingsSlice, item)
	}
	return obj
}

func (obj *roCEv2V4PeerRoCEv2FlowSettingsIter) Set(index int, newObj RoCEv2FlowSettings) RoCEv2V4PeerRoCEv2FlowSettingsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.roCEv2FlowSettingsSlice[index] = newObj
	return obj
}
func (obj *roCEv2V4PeerRoCEv2FlowSettingsIter) Clear() RoCEv2V4PeerRoCEv2FlowSettingsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RoCEv2FlowSettings{}
		obj.roCEv2FlowSettingsSlice = []RoCEv2FlowSettings{}
	}
	return obj
}
func (obj *roCEv2V4PeerRoCEv2FlowSettingsIter) clearHolderSlice() RoCEv2V4PeerRoCEv2FlowSettingsIter {
	if len(obj.roCEv2FlowSettingsSlice) > 0 {
		obj.roCEv2FlowSettingsSlice = []RoCEv2FlowSettings{}
	}
	return obj
}
func (obj *roCEv2V4PeerRoCEv2FlowSettingsIter) appendHolderSlice(item RoCEv2FlowSettings) RoCEv2V4PeerRoCEv2FlowSettingsIter {
	obj.roCEv2FlowSettingsSlice = append(obj.roCEv2FlowSettingsSlice, item)
	return obj
}

func (obj *roCEv2V4Peer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// NumOfQps is required
	if obj.obj.NumOfQps == nil {
		vObj.validationErrors = append(vObj.validationErrors, "NumOfQps is required field on interface RoCEv2V4Peer")
	}

	// DestinationIpAddress is required
	if obj.obj.DestinationIpAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "DestinationIpAddress is required field on interface RoCEv2V4Peer")
	}

	// IbMtu is required
	if obj.obj.IbMtu == nil {
		vObj.validationErrors = append(vObj.validationErrors, "IbMtu is required field on interface RoCEv2V4Peer")
	}
	if obj.obj.IbMtu != nil {

		if *obj.obj.IbMtu > 14000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2V4Peer.IbMtu <= 14000 but Got %d", *obj.obj.IbMtu))
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

func (obj *roCEv2V4Peer) setDefault() {

}
