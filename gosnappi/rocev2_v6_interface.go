package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2V6Interface *****
type rocev2V6Interface struct {
	validation
	obj          *otg.Rocev2V6Interface
	marshaller   marshalRocev2V6Interface
	unMarshaller unMarshalRocev2V6Interface
	peersHolder  Rocev2V6InterfaceRocev2V6PeerIter
}

func NewRocev2V6Interface() Rocev2V6Interface {
	obj := rocev2V6Interface{obj: &otg.Rocev2V6Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2V6Interface) msg() *otg.Rocev2V6Interface {
	return obj.obj
}

func (obj *rocev2V6Interface) setMsg(msg *otg.Rocev2V6Interface) Rocev2V6Interface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2V6Interface struct {
	obj *rocev2V6Interface
}

type marshalRocev2V6Interface interface {
	// ToProto marshals Rocev2V6Interface to protobuf object *otg.Rocev2V6Interface
	ToProto() (*otg.Rocev2V6Interface, error)
	// ToPbText marshals Rocev2V6Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2V6Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2V6Interface to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2V6Interface to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2V6Interface struct {
	obj *rocev2V6Interface
}

type unMarshalRocev2V6Interface interface {
	// FromProto unmarshals Rocev2V6Interface from protobuf object *otg.Rocev2V6Interface
	FromProto(msg *otg.Rocev2V6Interface) (Rocev2V6Interface, error)
	// FromPbText unmarshals Rocev2V6Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2V6Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2V6Interface from JSON text
	FromJson(value string) error
}

func (obj *rocev2V6Interface) Marshal() marshalRocev2V6Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2V6Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2V6Interface) Unmarshal() unMarshalRocev2V6Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2V6Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2V6Interface) ToProto() (*otg.Rocev2V6Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2V6Interface) FromProto(msg *otg.Rocev2V6Interface) (Rocev2V6Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2V6Interface) ToPbText() (string, error) {
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

func (m *unMarshalrocev2V6Interface) FromPbText(value string) error {
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

func (m *marshalrocev2V6Interface) ToYaml() (string, error) {
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

func (m *unMarshalrocev2V6Interface) FromYaml(value string) error {
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

func (m *marshalrocev2V6Interface) ToJsonRaw() (string, error) {
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

func (m *marshalrocev2V6Interface) ToJson() (string, error) {
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

func (m *unMarshalrocev2V6Interface) FromJson(value string) error {
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

func (obj *rocev2V6Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2V6Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2V6Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2V6Interface) Clone() (Rocev2V6Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2V6Interface()
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

func (obj *rocev2V6Interface) setNil() {
	obj.peersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2V6Interface is this contains an array of references to IPv6 interfaces, each having a list of IPv6 peers to various destinations.
type Rocev2V6Interface interface {
	Validation
	// msg marshals Rocev2V6Interface to protobuf object *otg.Rocev2V6Interface
	// and doesn't set defaults
	msg() *otg.Rocev2V6Interface
	// setMsg unmarshals Rocev2V6Interface from protobuf object *otg.Rocev2V6Interface
	// and doesn't set defaults
	setMsg(*otg.Rocev2V6Interface) Rocev2V6Interface
	// provides marshal interface
	Marshal() marshalRocev2V6Interface
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2V6Interface
	// validate validates Rocev2V6Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2V6Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv6Name returns string, set in Rocev2V6Interface.
	Ipv6Name() string
	// SetIpv6Name assigns string provided by user to Rocev2V6Interface
	SetIpv6Name(value string) Rocev2V6Interface
	// IbMtu returns uint32, set in Rocev2V6Interface.
	IbMtu() uint32
	// SetIbMtu assigns uint32 provided by user to Rocev2V6Interface
	SetIbMtu(value uint32) Rocev2V6Interface
	// HasIbMtu checks if IbMtu has been set in Rocev2V6Interface
	HasIbMtu() bool
	// Peers returns Rocev2V6InterfaceRocev2V6PeerIterIter, set in Rocev2V6Interface
	Peers() Rocev2V6InterfaceRocev2V6PeerIter
	setNil()
}

// The unique name of IPv6 used as the source IP for this list of RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// Ipv6Name returns a string
func (obj *rocev2V6Interface) Ipv6Name() string {

	return *obj.obj.Ipv6Name

}

// The unique name of IPv6 used as the source IP for this list of RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// SetIpv6Name sets the string value in the Rocev2V6Interface object
func (obj *rocev2V6Interface) SetIpv6Name(value string) Rocev2V6Interface {

	obj.obj.Ipv6Name = &value
	return obj
}

// The InfiniBand protocol defines several fixed sizes for the Maximum Transmission Unit (IB MTU): 256, 512, 1024, 2048, or 4096 bytes. RDMA messages will have a payload size that corresponds to the configured IB MTU. Additionally, it is possible to configure a custom size.
// IbMtu returns a uint32
func (obj *rocev2V6Interface) IbMtu() uint32 {

	return *obj.obj.IbMtu

}

// The InfiniBand protocol defines several fixed sizes for the Maximum Transmission Unit (IB MTU): 256, 512, 1024, 2048, or 4096 bytes. RDMA messages will have a payload size that corresponds to the configured IB MTU. Additionally, it is possible to configure a custom size.
// IbMtu returns a uint32
func (obj *rocev2V6Interface) HasIbMtu() bool {
	return obj.obj.IbMtu != nil
}

// The InfiniBand protocol defines several fixed sizes for the Maximum Transmission Unit (IB MTU): 256, 512, 1024, 2048, or 4096 bytes. RDMA messages will have a payload size that corresponds to the configured IB MTU. Additionally, it is possible to configure a custom size.
// SetIbMtu sets the uint32 value in the Rocev2V6Interface object
func (obj *rocev2V6Interface) SetIbMtu(value uint32) Rocev2V6Interface {

	obj.obj.IbMtu = &value
	return obj
}

// This contains the list of RoCEv2 IPv6 peers configured on this interface.
// Peers returns a []Rocev2V6Peer
func (obj *rocev2V6Interface) Peers() Rocev2V6InterfaceRocev2V6PeerIter {
	if len(obj.obj.Peers) == 0 {
		obj.obj.Peers = []*otg.Rocev2V6Peer{}
	}
	if obj.peersHolder == nil {
		obj.peersHolder = newRocev2V6InterfaceRocev2V6PeerIter(&obj.obj.Peers).setMsg(obj)
	}
	return obj.peersHolder
}

type rocev2V6InterfaceRocev2V6PeerIter struct {
	obj               *rocev2V6Interface
	rocev2V6PeerSlice []Rocev2V6Peer
	fieldPtr          *[]*otg.Rocev2V6Peer
}

func newRocev2V6InterfaceRocev2V6PeerIter(ptr *[]*otg.Rocev2V6Peer) Rocev2V6InterfaceRocev2V6PeerIter {
	return &rocev2V6InterfaceRocev2V6PeerIter{fieldPtr: ptr}
}

type Rocev2V6InterfaceRocev2V6PeerIter interface {
	setMsg(*rocev2V6Interface) Rocev2V6InterfaceRocev2V6PeerIter
	Items() []Rocev2V6Peer
	Add() Rocev2V6Peer
	Append(items ...Rocev2V6Peer) Rocev2V6InterfaceRocev2V6PeerIter
	Set(index int, newObj Rocev2V6Peer) Rocev2V6InterfaceRocev2V6PeerIter
	Clear() Rocev2V6InterfaceRocev2V6PeerIter
	clearHolderSlice() Rocev2V6InterfaceRocev2V6PeerIter
	appendHolderSlice(item Rocev2V6Peer) Rocev2V6InterfaceRocev2V6PeerIter
}

func (obj *rocev2V6InterfaceRocev2V6PeerIter) setMsg(msg *rocev2V6Interface) Rocev2V6InterfaceRocev2V6PeerIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rocev2V6Peer{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *rocev2V6InterfaceRocev2V6PeerIter) Items() []Rocev2V6Peer {
	return obj.rocev2V6PeerSlice
}

func (obj *rocev2V6InterfaceRocev2V6PeerIter) Add() Rocev2V6Peer {
	newObj := &otg.Rocev2V6Peer{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rocev2V6Peer{obj: newObj}
	newLibObj.setDefault()
	obj.rocev2V6PeerSlice = append(obj.rocev2V6PeerSlice, newLibObj)
	return newLibObj
}

func (obj *rocev2V6InterfaceRocev2V6PeerIter) Append(items ...Rocev2V6Peer) Rocev2V6InterfaceRocev2V6PeerIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rocev2V6PeerSlice = append(obj.rocev2V6PeerSlice, item)
	}
	return obj
}

func (obj *rocev2V6InterfaceRocev2V6PeerIter) Set(index int, newObj Rocev2V6Peer) Rocev2V6InterfaceRocev2V6PeerIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rocev2V6PeerSlice[index] = newObj
	return obj
}
func (obj *rocev2V6InterfaceRocev2V6PeerIter) Clear() Rocev2V6InterfaceRocev2V6PeerIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Rocev2V6Peer{}
		obj.rocev2V6PeerSlice = []Rocev2V6Peer{}
	}
	return obj
}
func (obj *rocev2V6InterfaceRocev2V6PeerIter) clearHolderSlice() Rocev2V6InterfaceRocev2V6PeerIter {
	if len(obj.rocev2V6PeerSlice) > 0 {
		obj.rocev2V6PeerSlice = []Rocev2V6Peer{}
	}
	return obj
}
func (obj *rocev2V6InterfaceRocev2V6PeerIter) appendHolderSlice(item Rocev2V6Peer) Rocev2V6InterfaceRocev2V6PeerIter {
	obj.rocev2V6PeerSlice = append(obj.rocev2V6PeerSlice, item)
	return obj
}

func (obj *rocev2V6Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Ipv6Name is required
	if obj.obj.Ipv6Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv6Name is required field on interface Rocev2V6Interface")
	}

	if obj.obj.IbMtu != nil {

		if *obj.obj.IbMtu > 14000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2V6Interface.IbMtu <= 14000 but Got %d", *obj.obj.IbMtu))
		}

	}

	if len(obj.obj.Peers) != 0 {

		if set_default {
			obj.Peers().clearHolderSlice()
			for _, item := range obj.obj.Peers {
				obj.Peers().appendHolderSlice(&rocev2V6Peer{obj: item})
			}
		}
		for _, item := range obj.Peers().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *rocev2V6Interface) setDefault() {
	if obj.obj.IbMtu == nil {
		obj.SetIbMtu(1024)
	}

}
