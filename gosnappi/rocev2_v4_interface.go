package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2V4Interface *****
type rocev2V4Interface struct {
	validation
	obj          *otg.Rocev2V4Interface
	marshaller   marshalRocev2V4Interface
	unMarshaller unMarshalRocev2V4Interface
	peersHolder  Rocev2V4InterfaceRocev2V4PeerIter
}

func NewRocev2V4Interface() Rocev2V4Interface {
	obj := rocev2V4Interface{obj: &otg.Rocev2V4Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2V4Interface) msg() *otg.Rocev2V4Interface {
	return obj.obj
}

func (obj *rocev2V4Interface) setMsg(msg *otg.Rocev2V4Interface) Rocev2V4Interface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2V4Interface struct {
	obj *rocev2V4Interface
}

type marshalRocev2V4Interface interface {
	// ToProto marshals Rocev2V4Interface to protobuf object *otg.Rocev2V4Interface
	ToProto() (*otg.Rocev2V4Interface, error)
	// ToPbText marshals Rocev2V4Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2V4Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2V4Interface to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2V4Interface to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2V4Interface struct {
	obj *rocev2V4Interface
}

type unMarshalRocev2V4Interface interface {
	// FromProto unmarshals Rocev2V4Interface from protobuf object *otg.Rocev2V4Interface
	FromProto(msg *otg.Rocev2V4Interface) (Rocev2V4Interface, error)
	// FromPbText unmarshals Rocev2V4Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2V4Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2V4Interface from JSON text
	FromJson(value string) error
}

func (obj *rocev2V4Interface) Marshal() marshalRocev2V4Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2V4Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2V4Interface) Unmarshal() unMarshalRocev2V4Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2V4Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2V4Interface) ToProto() (*otg.Rocev2V4Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2V4Interface) FromProto(msg *otg.Rocev2V4Interface) (Rocev2V4Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2V4Interface) ToPbText() (string, error) {
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

func (m *unMarshalrocev2V4Interface) FromPbText(value string) error {
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

func (m *marshalrocev2V4Interface) ToYaml() (string, error) {
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

func (m *unMarshalrocev2V4Interface) FromYaml(value string) error {
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

func (m *marshalrocev2V4Interface) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalrocev2V4Interface) ToJson() (string, error) {
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

func (m *unMarshalrocev2V4Interface) FromJson(value string) error {
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

func (obj *rocev2V4Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2V4Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2V4Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2V4Interface) Clone() (Rocev2V4Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2V4Interface()
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

func (obj *rocev2V4Interface) setNil() {
	obj.peersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2V4Interface is this contains an array of references to IPv4 interfaces, each having a list of IPv4 peers to various destinations.
type Rocev2V4Interface interface {
	Validation
	// msg marshals Rocev2V4Interface to protobuf object *otg.Rocev2V4Interface
	// and doesn't set defaults
	msg() *otg.Rocev2V4Interface
	// setMsg unmarshals Rocev2V4Interface from protobuf object *otg.Rocev2V4Interface
	// and doesn't set defaults
	setMsg(*otg.Rocev2V4Interface) Rocev2V4Interface
	// provides marshal interface
	Marshal() marshalRocev2V4Interface
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2V4Interface
	// validate validates Rocev2V4Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2V4Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Name returns string, set in Rocev2V4Interface.
	Ipv4Name() string
	// SetIpv4Name assigns string provided by user to Rocev2V4Interface
	SetIpv4Name(value string) Rocev2V4Interface
	// IbMtu returns uint32, set in Rocev2V4Interface.
	IbMtu() uint32
	// SetIbMtu assigns uint32 provided by user to Rocev2V4Interface
	SetIbMtu(value uint32) Rocev2V4Interface
	// HasIbMtu checks if IbMtu has been set in Rocev2V4Interface
	HasIbMtu() bool
	// Peers returns Rocev2V4InterfaceRocev2V4PeerIterIter, set in Rocev2V4Interface
	Peers() Rocev2V4InterfaceRocev2V4PeerIter
	setNil()
}

// The unique name of the IPv4 interface, used as the source IP for this list of RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// Ipv4Name returns a string
func (obj *rocev2V4Interface) Ipv4Name() string {

	return *obj.obj.Ipv4Name

}

// The unique name of the IPv4 interface, used as the source IP for this list of RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// SetIpv4Name sets the string value in the Rocev2V4Interface object
func (obj *rocev2V4Interface) SetIpv4Name(value string) Rocev2V4Interface {

	obj.obj.Ipv4Name = &value
	return obj
}

// The InfiniBand protocol defines several fixed sizes for the Maximum Transmission Unit (IB MTU): 256, 512, 1024, 2048, or 4096 bytes. RDMA messages will have a payload size that corresponds to the configured IB MTU. Additionally, it is possible to configure a custom size.
// IbMtu returns a uint32
func (obj *rocev2V4Interface) IbMtu() uint32 {

	return *obj.obj.IbMtu

}

// The InfiniBand protocol defines several fixed sizes for the Maximum Transmission Unit (IB MTU): 256, 512, 1024, 2048, or 4096 bytes. RDMA messages will have a payload size that corresponds to the configured IB MTU. Additionally, it is possible to configure a custom size.
// IbMtu returns a uint32
func (obj *rocev2V4Interface) HasIbMtu() bool {
	return obj.obj.IbMtu != nil
}

// The InfiniBand protocol defines several fixed sizes for the Maximum Transmission Unit (IB MTU): 256, 512, 1024, 2048, or 4096 bytes. RDMA messages will have a payload size that corresponds to the configured IB MTU. Additionally, it is possible to configure a custom size.
// SetIbMtu sets the uint32 value in the Rocev2V4Interface object
func (obj *rocev2V4Interface) SetIbMtu(value uint32) Rocev2V4Interface {

	obj.obj.IbMtu = &value
	return obj
}

// This contains the list of RoCEv2 peers configured on this interface.
// Peers returns a []Rocev2V4Peer
func (obj *rocev2V4Interface) Peers() Rocev2V4InterfaceRocev2V4PeerIter {
	if len(obj.obj.Peers) == 0 {
		obj.obj.Peers = []*otg.Rocev2V4Peer{}
	}
	if obj.peersHolder == nil {
		obj.peersHolder = newRocev2V4InterfaceRocev2V4PeerIter(&obj.obj.Peers).setMsg(obj)
	}
	return obj.peersHolder
}

type rocev2V4InterfaceRocev2V4PeerIter struct {
	obj               *rocev2V4Interface
	rocev2V4PeerSlice []Rocev2V4Peer
	fieldPtr          *[]*otg.Rocev2V4Peer
}

func newRocev2V4InterfaceRocev2V4PeerIter(ptr *[]*otg.Rocev2V4Peer) Rocev2V4InterfaceRocev2V4PeerIter {
	return &rocev2V4InterfaceRocev2V4PeerIter{fieldPtr: ptr}
}

type Rocev2V4InterfaceRocev2V4PeerIter interface {
	setMsg(*rocev2V4Interface) Rocev2V4InterfaceRocev2V4PeerIter
	Items() []Rocev2V4Peer
	Add() Rocev2V4Peer
	Append(items ...Rocev2V4Peer) Rocev2V4InterfaceRocev2V4PeerIter
	Set(index int, newObj Rocev2V4Peer) Rocev2V4InterfaceRocev2V4PeerIter
	Clear() Rocev2V4InterfaceRocev2V4PeerIter
	clearHolderSlice() Rocev2V4InterfaceRocev2V4PeerIter
	appendHolderSlice(item Rocev2V4Peer) Rocev2V4InterfaceRocev2V4PeerIter
}

func (obj *rocev2V4InterfaceRocev2V4PeerIter) setMsg(msg *rocev2V4Interface) Rocev2V4InterfaceRocev2V4PeerIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rocev2V4Peer{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *rocev2V4InterfaceRocev2V4PeerIter) Items() []Rocev2V4Peer {
	return obj.rocev2V4PeerSlice
}

func (obj *rocev2V4InterfaceRocev2V4PeerIter) Add() Rocev2V4Peer {
	newObj := &otg.Rocev2V4Peer{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rocev2V4Peer{obj: newObj}
	newLibObj.setDefault()
	obj.rocev2V4PeerSlice = append(obj.rocev2V4PeerSlice, newLibObj)
	return newLibObj
}

func (obj *rocev2V4InterfaceRocev2V4PeerIter) Append(items ...Rocev2V4Peer) Rocev2V4InterfaceRocev2V4PeerIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rocev2V4PeerSlice = append(obj.rocev2V4PeerSlice, item)
	}
	return obj
}

func (obj *rocev2V4InterfaceRocev2V4PeerIter) Set(index int, newObj Rocev2V4Peer) Rocev2V4InterfaceRocev2V4PeerIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rocev2V4PeerSlice[index] = newObj
	return obj
}
func (obj *rocev2V4InterfaceRocev2V4PeerIter) Clear() Rocev2V4InterfaceRocev2V4PeerIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Rocev2V4Peer{}
		obj.rocev2V4PeerSlice = []Rocev2V4Peer{}
	}
	return obj
}
func (obj *rocev2V4InterfaceRocev2V4PeerIter) clearHolderSlice() Rocev2V4InterfaceRocev2V4PeerIter {
	if len(obj.rocev2V4PeerSlice) > 0 {
		obj.rocev2V4PeerSlice = []Rocev2V4Peer{}
	}
	return obj
}
func (obj *rocev2V4InterfaceRocev2V4PeerIter) appendHolderSlice(item Rocev2V4Peer) Rocev2V4InterfaceRocev2V4PeerIter {
	obj.rocev2V4PeerSlice = append(obj.rocev2V4PeerSlice, item)
	return obj
}

func (obj *rocev2V4Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Ipv4Name is required
	if obj.obj.Ipv4Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4Name is required field on interface Rocev2V4Interface")
	}

	if obj.obj.IbMtu != nil {

		if *obj.obj.IbMtu > 14000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2V4Interface.IbMtu <= 14000 but Got %d", *obj.obj.IbMtu))
		}

	}

	if len(obj.obj.Peers) != 0 {

		if set_default {
			obj.Peers().clearHolderSlice()
			for _, item := range obj.obj.Peers {
				obj.Peers().appendHolderSlice(&rocev2V4Peer{obj: item})
			}
		}
		for _, item := range obj.Peers().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *rocev2V4Interface) setDefault() {
	if obj.obj.IbMtu == nil {
		obj.SetIbMtu(1024)
	}

}
