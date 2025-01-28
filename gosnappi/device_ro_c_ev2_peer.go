package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceRoCEv2Peer *****
type deviceRoCEv2Peer struct {
	validation
	obj                  *otg.DeviceRoCEv2Peer
	marshaller           marshalDeviceRoCEv2Peer
	unMarshaller         unMarshalDeviceRoCEv2Peer
	ipv4InterfacesHolder DeviceRoCEv2PeerRoCEv2V4InterfaceIter
	ipv6InterfacesHolder DeviceRoCEv2PeerRoCEv2V6InterfaceIter
}

func NewDeviceRoCEv2Peer() DeviceRoCEv2Peer {
	obj := deviceRoCEv2Peer{obj: &otg.DeviceRoCEv2Peer{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceRoCEv2Peer) msg() *otg.DeviceRoCEv2Peer {
	return obj.obj
}

func (obj *deviceRoCEv2Peer) setMsg(msg *otg.DeviceRoCEv2Peer) DeviceRoCEv2Peer {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceRoCEv2Peer struct {
	obj *deviceRoCEv2Peer
}

type marshalDeviceRoCEv2Peer interface {
	// ToProto marshals DeviceRoCEv2Peer to protobuf object *otg.DeviceRoCEv2Peer
	ToProto() (*otg.DeviceRoCEv2Peer, error)
	// ToPbText marshals DeviceRoCEv2Peer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceRoCEv2Peer to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceRoCEv2Peer to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceRoCEv2Peer struct {
	obj *deviceRoCEv2Peer
}

type unMarshalDeviceRoCEv2Peer interface {
	// FromProto unmarshals DeviceRoCEv2Peer from protobuf object *otg.DeviceRoCEv2Peer
	FromProto(msg *otg.DeviceRoCEv2Peer) (DeviceRoCEv2Peer, error)
	// FromPbText unmarshals DeviceRoCEv2Peer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceRoCEv2Peer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceRoCEv2Peer from JSON text
	FromJson(value string) error
}

func (obj *deviceRoCEv2Peer) Marshal() marshalDeviceRoCEv2Peer {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceRoCEv2Peer{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceRoCEv2Peer) Unmarshal() unMarshalDeviceRoCEv2Peer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceRoCEv2Peer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceRoCEv2Peer) ToProto() (*otg.DeviceRoCEv2Peer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceRoCEv2Peer) FromProto(msg *otg.DeviceRoCEv2Peer) (DeviceRoCEv2Peer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceRoCEv2Peer) ToPbText() (string, error) {
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

func (m *unMarshaldeviceRoCEv2Peer) FromPbText(value string) error {
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

func (m *marshaldeviceRoCEv2Peer) ToYaml() (string, error) {
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

func (m *unMarshaldeviceRoCEv2Peer) FromYaml(value string) error {
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

func (m *marshaldeviceRoCEv2Peer) ToJson() (string, error) {
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

func (m *unMarshaldeviceRoCEv2Peer) FromJson(value string) error {
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

func (obj *deviceRoCEv2Peer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceRoCEv2Peer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceRoCEv2Peer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceRoCEv2Peer) Clone() (DeviceRoCEv2Peer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceRoCEv2Peer()
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

func (obj *deviceRoCEv2Peer) setNil() {
	obj.ipv4InterfacesHolder = nil
	obj.ipv6InterfacesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceRoCEv2Peer is configuration for one or more IPv4 or IPv6 RoCEv2 Peers.
type DeviceRoCEv2Peer interface {
	Validation
	// msg marshals DeviceRoCEv2Peer to protobuf object *otg.DeviceRoCEv2Peer
	// and doesn't set defaults
	msg() *otg.DeviceRoCEv2Peer
	// setMsg unmarshals DeviceRoCEv2Peer from protobuf object *otg.DeviceRoCEv2Peer
	// and doesn't set defaults
	setMsg(*otg.DeviceRoCEv2Peer) DeviceRoCEv2Peer
	// provides marshal interface
	Marshal() marshalDeviceRoCEv2Peer
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceRoCEv2Peer
	// validate validates DeviceRoCEv2Peer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceRoCEv2Peer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Interfaces returns DeviceRoCEv2PeerRoCEv2V4InterfaceIterIter, set in DeviceRoCEv2Peer
	Ipv4Interfaces() DeviceRoCEv2PeerRoCEv2V4InterfaceIter
	// Ipv6Interfaces returns DeviceRoCEv2PeerRoCEv2V6InterfaceIterIter, set in DeviceRoCEv2Peer
	Ipv6Interfaces() DeviceRoCEv2PeerRoCEv2V6InterfaceIter
	setNil()
}

// This contains an array of references to IPv4 interfaces,  each of which will have list of peers to different destinations.
// Ipv4Interfaces returns a []RoCEv2V4Interface
func (obj *deviceRoCEv2Peer) Ipv4Interfaces() DeviceRoCEv2PeerRoCEv2V4InterfaceIter {
	if len(obj.obj.Ipv4Interfaces) == 0 {
		obj.obj.Ipv4Interfaces = []*otg.RoCEv2V4Interface{}
	}
	if obj.ipv4InterfacesHolder == nil {
		obj.ipv4InterfacesHolder = newDeviceRoCEv2PeerRoCEv2V4InterfaceIter(&obj.obj.Ipv4Interfaces).setMsg(obj)
	}
	return obj.ipv4InterfacesHolder
}

type deviceRoCEv2PeerRoCEv2V4InterfaceIter struct {
	obj                    *deviceRoCEv2Peer
	roCEv2V4InterfaceSlice []RoCEv2V4Interface
	fieldPtr               *[]*otg.RoCEv2V4Interface
}

func newDeviceRoCEv2PeerRoCEv2V4InterfaceIter(ptr *[]*otg.RoCEv2V4Interface) DeviceRoCEv2PeerRoCEv2V4InterfaceIter {
	return &deviceRoCEv2PeerRoCEv2V4InterfaceIter{fieldPtr: ptr}
}

type DeviceRoCEv2PeerRoCEv2V4InterfaceIter interface {
	setMsg(*deviceRoCEv2Peer) DeviceRoCEv2PeerRoCEv2V4InterfaceIter
	Items() []RoCEv2V4Interface
	Add() RoCEv2V4Interface
	Append(items ...RoCEv2V4Interface) DeviceRoCEv2PeerRoCEv2V4InterfaceIter
	Set(index int, newObj RoCEv2V4Interface) DeviceRoCEv2PeerRoCEv2V4InterfaceIter
	Clear() DeviceRoCEv2PeerRoCEv2V4InterfaceIter
	clearHolderSlice() DeviceRoCEv2PeerRoCEv2V4InterfaceIter
	appendHolderSlice(item RoCEv2V4Interface) DeviceRoCEv2PeerRoCEv2V4InterfaceIter
}

func (obj *deviceRoCEv2PeerRoCEv2V4InterfaceIter) setMsg(msg *deviceRoCEv2Peer) DeviceRoCEv2PeerRoCEv2V4InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&roCEv2V4Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceRoCEv2PeerRoCEv2V4InterfaceIter) Items() []RoCEv2V4Interface {
	return obj.roCEv2V4InterfaceSlice
}

func (obj *deviceRoCEv2PeerRoCEv2V4InterfaceIter) Add() RoCEv2V4Interface {
	newObj := &otg.RoCEv2V4Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &roCEv2V4Interface{obj: newObj}
	newLibObj.setDefault()
	obj.roCEv2V4InterfaceSlice = append(obj.roCEv2V4InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceRoCEv2PeerRoCEv2V4InterfaceIter) Append(items ...RoCEv2V4Interface) DeviceRoCEv2PeerRoCEv2V4InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.roCEv2V4InterfaceSlice = append(obj.roCEv2V4InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceRoCEv2PeerRoCEv2V4InterfaceIter) Set(index int, newObj RoCEv2V4Interface) DeviceRoCEv2PeerRoCEv2V4InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.roCEv2V4InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceRoCEv2PeerRoCEv2V4InterfaceIter) Clear() DeviceRoCEv2PeerRoCEv2V4InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RoCEv2V4Interface{}
		obj.roCEv2V4InterfaceSlice = []RoCEv2V4Interface{}
	}
	return obj
}
func (obj *deviceRoCEv2PeerRoCEv2V4InterfaceIter) clearHolderSlice() DeviceRoCEv2PeerRoCEv2V4InterfaceIter {
	if len(obj.roCEv2V4InterfaceSlice) > 0 {
		obj.roCEv2V4InterfaceSlice = []RoCEv2V4Interface{}
	}
	return obj
}
func (obj *deviceRoCEv2PeerRoCEv2V4InterfaceIter) appendHolderSlice(item RoCEv2V4Interface) DeviceRoCEv2PeerRoCEv2V4InterfaceIter {
	obj.roCEv2V4InterfaceSlice = append(obj.roCEv2V4InterfaceSlice, item)
	return obj
}

// This contains an array of references to IPv6 interfaces,  each of which will have list of peers to different destinations.
// Ipv6Interfaces returns a []RoCEv2V6Interface
func (obj *deviceRoCEv2Peer) Ipv6Interfaces() DeviceRoCEv2PeerRoCEv2V6InterfaceIter {
	if len(obj.obj.Ipv6Interfaces) == 0 {
		obj.obj.Ipv6Interfaces = []*otg.RoCEv2V6Interface{}
	}
	if obj.ipv6InterfacesHolder == nil {
		obj.ipv6InterfacesHolder = newDeviceRoCEv2PeerRoCEv2V6InterfaceIter(&obj.obj.Ipv6Interfaces).setMsg(obj)
	}
	return obj.ipv6InterfacesHolder
}

type deviceRoCEv2PeerRoCEv2V6InterfaceIter struct {
	obj                    *deviceRoCEv2Peer
	roCEv2V6InterfaceSlice []RoCEv2V6Interface
	fieldPtr               *[]*otg.RoCEv2V6Interface
}

func newDeviceRoCEv2PeerRoCEv2V6InterfaceIter(ptr *[]*otg.RoCEv2V6Interface) DeviceRoCEv2PeerRoCEv2V6InterfaceIter {
	return &deviceRoCEv2PeerRoCEv2V6InterfaceIter{fieldPtr: ptr}
}

type DeviceRoCEv2PeerRoCEv2V6InterfaceIter interface {
	setMsg(*deviceRoCEv2Peer) DeviceRoCEv2PeerRoCEv2V6InterfaceIter
	Items() []RoCEv2V6Interface
	Add() RoCEv2V6Interface
	Append(items ...RoCEv2V6Interface) DeviceRoCEv2PeerRoCEv2V6InterfaceIter
	Set(index int, newObj RoCEv2V6Interface) DeviceRoCEv2PeerRoCEv2V6InterfaceIter
	Clear() DeviceRoCEv2PeerRoCEv2V6InterfaceIter
	clearHolderSlice() DeviceRoCEv2PeerRoCEv2V6InterfaceIter
	appendHolderSlice(item RoCEv2V6Interface) DeviceRoCEv2PeerRoCEv2V6InterfaceIter
}

func (obj *deviceRoCEv2PeerRoCEv2V6InterfaceIter) setMsg(msg *deviceRoCEv2Peer) DeviceRoCEv2PeerRoCEv2V6InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&roCEv2V6Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceRoCEv2PeerRoCEv2V6InterfaceIter) Items() []RoCEv2V6Interface {
	return obj.roCEv2V6InterfaceSlice
}

func (obj *deviceRoCEv2PeerRoCEv2V6InterfaceIter) Add() RoCEv2V6Interface {
	newObj := &otg.RoCEv2V6Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &roCEv2V6Interface{obj: newObj}
	newLibObj.setDefault()
	obj.roCEv2V6InterfaceSlice = append(obj.roCEv2V6InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceRoCEv2PeerRoCEv2V6InterfaceIter) Append(items ...RoCEv2V6Interface) DeviceRoCEv2PeerRoCEv2V6InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.roCEv2V6InterfaceSlice = append(obj.roCEv2V6InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceRoCEv2PeerRoCEv2V6InterfaceIter) Set(index int, newObj RoCEv2V6Interface) DeviceRoCEv2PeerRoCEv2V6InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.roCEv2V6InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceRoCEv2PeerRoCEv2V6InterfaceIter) Clear() DeviceRoCEv2PeerRoCEv2V6InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RoCEv2V6Interface{}
		obj.roCEv2V6InterfaceSlice = []RoCEv2V6Interface{}
	}
	return obj
}
func (obj *deviceRoCEv2PeerRoCEv2V6InterfaceIter) clearHolderSlice() DeviceRoCEv2PeerRoCEv2V6InterfaceIter {
	if len(obj.roCEv2V6InterfaceSlice) > 0 {
		obj.roCEv2V6InterfaceSlice = []RoCEv2V6Interface{}
	}
	return obj
}
func (obj *deviceRoCEv2PeerRoCEv2V6InterfaceIter) appendHolderSlice(item RoCEv2V6Interface) DeviceRoCEv2PeerRoCEv2V6InterfaceIter {
	obj.roCEv2V6InterfaceSlice = append(obj.roCEv2V6InterfaceSlice, item)
	return obj
}

func (obj *deviceRoCEv2Peer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4Interfaces) != 0 {

		if set_default {
			obj.Ipv4Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Interfaces {
				obj.Ipv4Interfaces().appendHolderSlice(&roCEv2V4Interface{obj: item})
			}
		}
		for _, item := range obj.Ipv4Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv6Interfaces) != 0 {

		if set_default {
			obj.Ipv6Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Ipv6Interfaces {
				obj.Ipv6Interfaces().appendHolderSlice(&roCEv2V6Interface{obj: item})
			}
		}
		for _, item := range obj.Ipv6Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceRoCEv2Peer) setDefault() {

}
