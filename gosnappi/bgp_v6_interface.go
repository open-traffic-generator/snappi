package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV6Interface *****
type bgpV6Interface struct {
	validation
	obj          *otg.BgpV6Interface
	marshaller   marshalBgpV6Interface
	unMarshaller unMarshalBgpV6Interface
	peersHolder  BgpV6InterfaceBgpV6PeerIter
}

func NewBgpV6Interface() BgpV6Interface {
	obj := bgpV6Interface{obj: &otg.BgpV6Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV6Interface) msg() *otg.BgpV6Interface {
	return obj.obj
}

func (obj *bgpV6Interface) setMsg(msg *otg.BgpV6Interface) BgpV6Interface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV6Interface struct {
	obj *bgpV6Interface
}

type marshalBgpV6Interface interface {
	// ToProto marshals BgpV6Interface to protobuf object *otg.BgpV6Interface
	ToProto() (*otg.BgpV6Interface, error)
	// ToPbText marshals BgpV6Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV6Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV6Interface to JSON text
	ToJson() (string, error)
}

type unMarshalbgpV6Interface struct {
	obj *bgpV6Interface
}

type unMarshalBgpV6Interface interface {
	// FromProto unmarshals BgpV6Interface from protobuf object *otg.BgpV6Interface
	FromProto(msg *otg.BgpV6Interface) (BgpV6Interface, error)
	// FromPbText unmarshals BgpV6Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV6Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV6Interface from JSON text
	FromJson(value string) error
}

func (obj *bgpV6Interface) Marshal() marshalBgpV6Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV6Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV6Interface) Unmarshal() unMarshalBgpV6Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV6Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV6Interface) ToProto() (*otg.BgpV6Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV6Interface) FromProto(msg *otg.BgpV6Interface) (BgpV6Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV6Interface) ToPbText() (string, error) {
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

func (m *unMarshalbgpV6Interface) FromPbText(value string) error {
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

func (m *marshalbgpV6Interface) ToYaml() (string, error) {
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

func (m *unMarshalbgpV6Interface) FromYaml(value string) error {
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

func (m *marshalbgpV6Interface) ToJson() (string, error) {
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

func (m *unMarshalbgpV6Interface) FromJson(value string) error {
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

func (obj *bgpV6Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV6Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV6Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV6Interface) Clone() (BgpV6Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV6Interface()
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

func (obj *bgpV6Interface) setNil() {
	obj.peersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpV6Interface is configuration for emulated BGPv6 peers and routes on a single IPv6 interface.
type BgpV6Interface interface {
	Validation
	// msg marshals BgpV6Interface to protobuf object *otg.BgpV6Interface
	// and doesn't set defaults
	msg() *otg.BgpV6Interface
	// setMsg unmarshals BgpV6Interface from protobuf object *otg.BgpV6Interface
	// and doesn't set defaults
	setMsg(*otg.BgpV6Interface) BgpV6Interface
	// provides marshal interface
	Marshal() marshalBgpV6Interface
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV6Interface
	// validate validates BgpV6Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV6Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv6Name returns string, set in BgpV6Interface.
	Ipv6Name() string
	// SetIpv6Name assigns string provided by user to BgpV6Interface
	SetIpv6Name(value string) BgpV6Interface
	// Peers returns BgpV6InterfaceBgpV6PeerIterIter, set in BgpV6Interface
	Peers() BgpV6InterfaceBgpV6PeerIter
	setNil()
}

// The unique name of IPv6 Loopback IPv6 interface or DHCPv4 client used as the source IP for this list of BGP peers.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Device.Ipv6Loopback/properties/name
// - /components/schemas/DhcpClient.V6/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Device.Ipv6Loopback/properties/name
// - /components/schemas/DhcpClient.V6/properties/name
//
// Ipv6Name returns a string
func (obj *bgpV6Interface) Ipv6Name() string {

	return *obj.obj.Ipv6Name

}

// The unique name of IPv6 Loopback IPv6 interface or DHCPv4 client used as the source IP for this list of BGP peers.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Device.Ipv6Loopback/properties/name
// - /components/schemas/DhcpClient.V6/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Device.Ipv6Loopback/properties/name
// - /components/schemas/DhcpClient.V6/properties/name
//
// SetIpv6Name sets the string value in the BgpV6Interface object
func (obj *bgpV6Interface) SetIpv6Name(value string) BgpV6Interface {

	obj.obj.Ipv6Name = &value
	return obj
}

// This contains the list of BGPv6 peers configured on this interface.
// Peers returns a []BgpV6Peer
func (obj *bgpV6Interface) Peers() BgpV6InterfaceBgpV6PeerIter {
	if len(obj.obj.Peers) == 0 {
		obj.obj.Peers = []*otg.BgpV6Peer{}
	}
	if obj.peersHolder == nil {
		obj.peersHolder = newBgpV6InterfaceBgpV6PeerIter(&obj.obj.Peers).setMsg(obj)
	}
	return obj.peersHolder
}

type bgpV6InterfaceBgpV6PeerIter struct {
	obj            *bgpV6Interface
	bgpV6PeerSlice []BgpV6Peer
	fieldPtr       *[]*otg.BgpV6Peer
}

func newBgpV6InterfaceBgpV6PeerIter(ptr *[]*otg.BgpV6Peer) BgpV6InterfaceBgpV6PeerIter {
	return &bgpV6InterfaceBgpV6PeerIter{fieldPtr: ptr}
}

type BgpV6InterfaceBgpV6PeerIter interface {
	setMsg(*bgpV6Interface) BgpV6InterfaceBgpV6PeerIter
	Items() []BgpV6Peer
	Add() BgpV6Peer
	Append(items ...BgpV6Peer) BgpV6InterfaceBgpV6PeerIter
	Set(index int, newObj BgpV6Peer) BgpV6InterfaceBgpV6PeerIter
	Clear() BgpV6InterfaceBgpV6PeerIter
	clearHolderSlice() BgpV6InterfaceBgpV6PeerIter
	appendHolderSlice(item BgpV6Peer) BgpV6InterfaceBgpV6PeerIter
}

func (obj *bgpV6InterfaceBgpV6PeerIter) setMsg(msg *bgpV6Interface) BgpV6InterfaceBgpV6PeerIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV6Peer{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6InterfaceBgpV6PeerIter) Items() []BgpV6Peer {
	return obj.bgpV6PeerSlice
}

func (obj *bgpV6InterfaceBgpV6PeerIter) Add() BgpV6Peer {
	newObj := &otg.BgpV6Peer{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV6Peer{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV6PeerSlice = append(obj.bgpV6PeerSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6InterfaceBgpV6PeerIter) Append(items ...BgpV6Peer) BgpV6InterfaceBgpV6PeerIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV6PeerSlice = append(obj.bgpV6PeerSlice, item)
	}
	return obj
}

func (obj *bgpV6InterfaceBgpV6PeerIter) Set(index int, newObj BgpV6Peer) BgpV6InterfaceBgpV6PeerIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV6PeerSlice[index] = newObj
	return obj
}
func (obj *bgpV6InterfaceBgpV6PeerIter) Clear() BgpV6InterfaceBgpV6PeerIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV6Peer{}
		obj.bgpV6PeerSlice = []BgpV6Peer{}
	}
	return obj
}
func (obj *bgpV6InterfaceBgpV6PeerIter) clearHolderSlice() BgpV6InterfaceBgpV6PeerIter {
	if len(obj.bgpV6PeerSlice) > 0 {
		obj.bgpV6PeerSlice = []BgpV6Peer{}
	}
	return obj
}
func (obj *bgpV6InterfaceBgpV6PeerIter) appendHolderSlice(item BgpV6Peer) BgpV6InterfaceBgpV6PeerIter {
	obj.bgpV6PeerSlice = append(obj.bgpV6PeerSlice, item)
	return obj
}

func (obj *bgpV6Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Ipv6Name is required
	if obj.obj.Ipv6Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv6Name is required field on interface BgpV6Interface")
	}

	if len(obj.obj.Peers) != 0 {

		if set_default {
			obj.Peers().clearHolderSlice()
			for _, item := range obj.obj.Peers {
				obj.Peers().appendHolderSlice(&bgpV6Peer{obj: item})
			}
		}
		for _, item := range obj.Peers().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpV6Interface) setDefault() {

}
