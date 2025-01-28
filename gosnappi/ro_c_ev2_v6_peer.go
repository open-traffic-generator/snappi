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
	// MaxNumOfQp returns int32, set in RoCEv2V6Peer.
	MaxNumOfQp() int32
	// SetMaxNumOfQp assigns int32 provided by user to RoCEv2V6Peer
	SetMaxNumOfQp(value int32) RoCEv2V6Peer
	// RemoteEndPointIpAddress returns string, set in RoCEv2V6Peer.
	RemoteEndPointIpAddress() string
	// SetRemoteEndPointIpAddress assigns string provided by user to RoCEv2V6Peer
	SetRemoteEndPointIpAddress(value string) RoCEv2V6Peer
	// IbMtu returns uint32, set in RoCEv2V6Peer.
	IbMtu() uint32
	// SetIbMtu assigns uint32 provided by user to RoCEv2V6Peer
	SetIbMtu(value uint32) RoCEv2V6Peer
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

// Maximum number of QP per External Endpoint.
// MaxNumOfQp returns a int32
func (obj *roCEv2V6Peer) MaxNumOfQp() int32 {

	return *obj.obj.MaxNumOfQp

}

// Maximum number of QP per External Endpoint.
// SetMaxNumOfQp sets the int32 value in the RoCEv2V6Peer object
func (obj *roCEv2V6Peer) SetMaxNumOfQp(value int32) RoCEv2V6Peer {

	obj.obj.MaxNumOfQp = &value
	return obj
}

// Specify the IPv4 address of External NIC i.e Remote End Point IP Address.
// RemoteEndPointIpAddress returns a string
func (obj *roCEv2V6Peer) RemoteEndPointIpAddress() string {

	return *obj.obj.RemoteEndPointIpAddress

}

// Specify the IPv4 address of External NIC i.e Remote End Point IP Address.
// SetRemoteEndPointIpAddress sets the string value in the RoCEv2V6Peer object
func (obj *roCEv2V6Peer) SetRemoteEndPointIpAddress(value string) RoCEv2V6Peer {

	obj.obj.RemoteEndPointIpAddress = &value
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

// This section has two views, Local End and Remote End. Both views have same configurations. However, the remote and local peer IP addresses are interchanged. This configuration allows you to configure RDMA flow over the same QP number from same source and destination. Default value for commands at Remote peer is set to None. So that by default, this remote peer does not initiate any traffic flow.
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

	// MaxNumOfQp is required
	if obj.obj.MaxNumOfQp == nil {
		vObj.validationErrors = append(vObj.validationErrors, "MaxNumOfQp is required field on interface RoCEv2V6Peer")
	}

	// RemoteEndPointIpAddress is required
	if obj.obj.RemoteEndPointIpAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RemoteEndPointIpAddress is required field on interface RoCEv2V6Peer")
	}
	if obj.obj.RemoteEndPointIpAddress != nil {

		err := obj.validateIpv4(obj.RemoteEndPointIpAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RoCEv2V6Peer.RemoteEndPointIpAddress"))
		}

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
