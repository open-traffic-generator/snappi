package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceRocev2Peer *****
type deviceRocev2Peer struct {
	validation
	obj                  *otg.DeviceRocev2Peer
	marshaller           marshalDeviceRocev2Peer
	unMarshaller         unMarshalDeviceRocev2Peer
	ipv4InterfacesHolder DeviceRocev2PeerRocev2V4InterfaceIter
	ipv6InterfacesHolder DeviceRocev2PeerRocev2V6InterfaceIter
}

func NewDeviceRocev2Peer() DeviceRocev2Peer {
	obj := deviceRocev2Peer{obj: &otg.DeviceRocev2Peer{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceRocev2Peer) msg() *otg.DeviceRocev2Peer {
	return obj.obj
}

func (obj *deviceRocev2Peer) setMsg(msg *otg.DeviceRocev2Peer) DeviceRocev2Peer {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceRocev2Peer struct {
	obj *deviceRocev2Peer
}

type marshalDeviceRocev2Peer interface {
	// ToProto marshals DeviceRocev2Peer to protobuf object *otg.DeviceRocev2Peer
	ToProto() (*otg.DeviceRocev2Peer, error)
	// ToPbText marshals DeviceRocev2Peer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceRocev2Peer to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceRocev2Peer to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceRocev2Peer struct {
	obj *deviceRocev2Peer
}

type unMarshalDeviceRocev2Peer interface {
	// FromProto unmarshals DeviceRocev2Peer from protobuf object *otg.DeviceRocev2Peer
	FromProto(msg *otg.DeviceRocev2Peer) (DeviceRocev2Peer, error)
	// FromPbText unmarshals DeviceRocev2Peer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceRocev2Peer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceRocev2Peer from JSON text
	FromJson(value string) error
}

func (obj *deviceRocev2Peer) Marshal() marshalDeviceRocev2Peer {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceRocev2Peer{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceRocev2Peer) Unmarshal() unMarshalDeviceRocev2Peer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceRocev2Peer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceRocev2Peer) ToProto() (*otg.DeviceRocev2Peer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceRocev2Peer) FromProto(msg *otg.DeviceRocev2Peer) (DeviceRocev2Peer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceRocev2Peer) ToPbText() (string, error) {
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

func (m *unMarshaldeviceRocev2Peer) FromPbText(value string) error {
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

func (m *marshaldeviceRocev2Peer) ToYaml() (string, error) {
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

func (m *unMarshaldeviceRocev2Peer) FromYaml(value string) error {
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

func (m *marshaldeviceRocev2Peer) ToJson() (string, error) {
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

func (m *unMarshaldeviceRocev2Peer) FromJson(value string) error {
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

func (obj *deviceRocev2Peer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceRocev2Peer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceRocev2Peer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceRocev2Peer) Clone() (DeviceRocev2Peer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceRocev2Peer()
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

func (obj *deviceRocev2Peer) setNil() {
	obj.ipv4InterfacesHolder = nil
	obj.ipv6InterfacesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceRocev2Peer is configuration for one or more IPv4 or IPv6 RoCEv2 Peers.
type DeviceRocev2Peer interface {
	Validation
	// msg marshals DeviceRocev2Peer to protobuf object *otg.DeviceRocev2Peer
	// and doesn't set defaults
	msg() *otg.DeviceRocev2Peer
	// setMsg unmarshals DeviceRocev2Peer from protobuf object *otg.DeviceRocev2Peer
	// and doesn't set defaults
	setMsg(*otg.DeviceRocev2Peer) DeviceRocev2Peer
	// provides marshal interface
	Marshal() marshalDeviceRocev2Peer
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceRocev2Peer
	// validate validates DeviceRocev2Peer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceRocev2Peer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Interfaces returns DeviceRocev2PeerRocev2V4InterfaceIterIter, set in DeviceRocev2Peer
	Ipv4Interfaces() DeviceRocev2PeerRocev2V4InterfaceIter
	// Ipv6Interfaces returns DeviceRocev2PeerRocev2V6InterfaceIterIter, set in DeviceRocev2Peer
	Ipv6Interfaces() DeviceRocev2PeerRocev2V6InterfaceIter
	setNil()
}

// This contains an array of references to IPv4 interfaces, each having a list of IPv4 peers to various destinations.
// Ipv4Interfaces returns a []Rocev2V4Interface
func (obj *deviceRocev2Peer) Ipv4Interfaces() DeviceRocev2PeerRocev2V4InterfaceIter {
	if len(obj.obj.Ipv4Interfaces) == 0 {
		obj.obj.Ipv4Interfaces = []*otg.Rocev2V4Interface{}
	}
	if obj.ipv4InterfacesHolder == nil {
		obj.ipv4InterfacesHolder = newDeviceRocev2PeerRocev2V4InterfaceIter(&obj.obj.Ipv4Interfaces).setMsg(obj)
	}
	return obj.ipv4InterfacesHolder
}

type deviceRocev2PeerRocev2V4InterfaceIter struct {
	obj                    *deviceRocev2Peer
	rocev2V4InterfaceSlice []Rocev2V4Interface
	fieldPtr               *[]*otg.Rocev2V4Interface
}

func newDeviceRocev2PeerRocev2V4InterfaceIter(ptr *[]*otg.Rocev2V4Interface) DeviceRocev2PeerRocev2V4InterfaceIter {
	return &deviceRocev2PeerRocev2V4InterfaceIter{fieldPtr: ptr}
}

type DeviceRocev2PeerRocev2V4InterfaceIter interface {
	setMsg(*deviceRocev2Peer) DeviceRocev2PeerRocev2V4InterfaceIter
	Items() []Rocev2V4Interface
	Add() Rocev2V4Interface
	Append(items ...Rocev2V4Interface) DeviceRocev2PeerRocev2V4InterfaceIter
	Set(index int, newObj Rocev2V4Interface) DeviceRocev2PeerRocev2V4InterfaceIter
	Clear() DeviceRocev2PeerRocev2V4InterfaceIter
	clearHolderSlice() DeviceRocev2PeerRocev2V4InterfaceIter
	appendHolderSlice(item Rocev2V4Interface) DeviceRocev2PeerRocev2V4InterfaceIter
}

func (obj *deviceRocev2PeerRocev2V4InterfaceIter) setMsg(msg *deviceRocev2Peer) DeviceRocev2PeerRocev2V4InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rocev2V4Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceRocev2PeerRocev2V4InterfaceIter) Items() []Rocev2V4Interface {
	return obj.rocev2V4InterfaceSlice
}

func (obj *deviceRocev2PeerRocev2V4InterfaceIter) Add() Rocev2V4Interface {
	newObj := &otg.Rocev2V4Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rocev2V4Interface{obj: newObj}
	newLibObj.setDefault()
	obj.rocev2V4InterfaceSlice = append(obj.rocev2V4InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceRocev2PeerRocev2V4InterfaceIter) Append(items ...Rocev2V4Interface) DeviceRocev2PeerRocev2V4InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rocev2V4InterfaceSlice = append(obj.rocev2V4InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceRocev2PeerRocev2V4InterfaceIter) Set(index int, newObj Rocev2V4Interface) DeviceRocev2PeerRocev2V4InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rocev2V4InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceRocev2PeerRocev2V4InterfaceIter) Clear() DeviceRocev2PeerRocev2V4InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Rocev2V4Interface{}
		obj.rocev2V4InterfaceSlice = []Rocev2V4Interface{}
	}
	return obj
}
func (obj *deviceRocev2PeerRocev2V4InterfaceIter) clearHolderSlice() DeviceRocev2PeerRocev2V4InterfaceIter {
	if len(obj.rocev2V4InterfaceSlice) > 0 {
		obj.rocev2V4InterfaceSlice = []Rocev2V4Interface{}
	}
	return obj
}
func (obj *deviceRocev2PeerRocev2V4InterfaceIter) appendHolderSlice(item Rocev2V4Interface) DeviceRocev2PeerRocev2V4InterfaceIter {
	obj.rocev2V4InterfaceSlice = append(obj.rocev2V4InterfaceSlice, item)
	return obj
}

// This contains an array references to IPv6 interfaces, each with a list of IPv6 peers for various destinations.
// Ipv6Interfaces returns a []Rocev2V6Interface
func (obj *deviceRocev2Peer) Ipv6Interfaces() DeviceRocev2PeerRocev2V6InterfaceIter {
	if len(obj.obj.Ipv6Interfaces) == 0 {
		obj.obj.Ipv6Interfaces = []*otg.Rocev2V6Interface{}
	}
	if obj.ipv6InterfacesHolder == nil {
		obj.ipv6InterfacesHolder = newDeviceRocev2PeerRocev2V6InterfaceIter(&obj.obj.Ipv6Interfaces).setMsg(obj)
	}
	return obj.ipv6InterfacesHolder
}

type deviceRocev2PeerRocev2V6InterfaceIter struct {
	obj                    *deviceRocev2Peer
	rocev2V6InterfaceSlice []Rocev2V6Interface
	fieldPtr               *[]*otg.Rocev2V6Interface
}

func newDeviceRocev2PeerRocev2V6InterfaceIter(ptr *[]*otg.Rocev2V6Interface) DeviceRocev2PeerRocev2V6InterfaceIter {
	return &deviceRocev2PeerRocev2V6InterfaceIter{fieldPtr: ptr}
}

type DeviceRocev2PeerRocev2V6InterfaceIter interface {
	setMsg(*deviceRocev2Peer) DeviceRocev2PeerRocev2V6InterfaceIter
	Items() []Rocev2V6Interface
	Add() Rocev2V6Interface
	Append(items ...Rocev2V6Interface) DeviceRocev2PeerRocev2V6InterfaceIter
	Set(index int, newObj Rocev2V6Interface) DeviceRocev2PeerRocev2V6InterfaceIter
	Clear() DeviceRocev2PeerRocev2V6InterfaceIter
	clearHolderSlice() DeviceRocev2PeerRocev2V6InterfaceIter
	appendHolderSlice(item Rocev2V6Interface) DeviceRocev2PeerRocev2V6InterfaceIter
}

func (obj *deviceRocev2PeerRocev2V6InterfaceIter) setMsg(msg *deviceRocev2Peer) DeviceRocev2PeerRocev2V6InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rocev2V6Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceRocev2PeerRocev2V6InterfaceIter) Items() []Rocev2V6Interface {
	return obj.rocev2V6InterfaceSlice
}

func (obj *deviceRocev2PeerRocev2V6InterfaceIter) Add() Rocev2V6Interface {
	newObj := &otg.Rocev2V6Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rocev2V6Interface{obj: newObj}
	newLibObj.setDefault()
	obj.rocev2V6InterfaceSlice = append(obj.rocev2V6InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceRocev2PeerRocev2V6InterfaceIter) Append(items ...Rocev2V6Interface) DeviceRocev2PeerRocev2V6InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rocev2V6InterfaceSlice = append(obj.rocev2V6InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceRocev2PeerRocev2V6InterfaceIter) Set(index int, newObj Rocev2V6Interface) DeviceRocev2PeerRocev2V6InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rocev2V6InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceRocev2PeerRocev2V6InterfaceIter) Clear() DeviceRocev2PeerRocev2V6InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Rocev2V6Interface{}
		obj.rocev2V6InterfaceSlice = []Rocev2V6Interface{}
	}
	return obj
}
func (obj *deviceRocev2PeerRocev2V6InterfaceIter) clearHolderSlice() DeviceRocev2PeerRocev2V6InterfaceIter {
	if len(obj.rocev2V6InterfaceSlice) > 0 {
		obj.rocev2V6InterfaceSlice = []Rocev2V6Interface{}
	}
	return obj
}
func (obj *deviceRocev2PeerRocev2V6InterfaceIter) appendHolderSlice(item Rocev2V6Interface) DeviceRocev2PeerRocev2V6InterfaceIter {
	obj.rocev2V6InterfaceSlice = append(obj.rocev2V6InterfaceSlice, item)
	return obj
}

func (obj *deviceRocev2Peer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4Interfaces) != 0 {

		if set_default {
			obj.Ipv4Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Interfaces {
				obj.Ipv4Interfaces().appendHolderSlice(&rocev2V4Interface{obj: item})
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
				obj.Ipv6Interfaces().appendHolderSlice(&rocev2V6Interface{obj: item})
			}
		}
		for _, item := range obj.Ipv6Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceRocev2Peer) setDefault() {

}
