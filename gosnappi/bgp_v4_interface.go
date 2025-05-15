package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV4Interface *****
type bgpV4Interface struct {
	validation
	obj          *otg.BgpV4Interface
	marshaller   marshalBgpV4Interface
	unMarshaller unMarshalBgpV4Interface
	peersHolder  BgpV4InterfaceBgpV4PeerIter
}

func NewBgpV4Interface() BgpV4Interface {
	obj := bgpV4Interface{obj: &otg.BgpV4Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV4Interface) msg() *otg.BgpV4Interface {
	return obj.obj
}

func (obj *bgpV4Interface) setMsg(msg *otg.BgpV4Interface) BgpV4Interface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV4Interface struct {
	obj *bgpV4Interface
}

type marshalBgpV4Interface interface {
	// ToProto marshals BgpV4Interface to protobuf object *otg.BgpV4Interface
	ToProto() (*otg.BgpV4Interface, error)
	// ToPbText marshals BgpV4Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV4Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV4Interface to JSON text
	ToJson() (string, error)
}

type unMarshalbgpV4Interface struct {
	obj *bgpV4Interface
}

type unMarshalBgpV4Interface interface {
	// FromProto unmarshals BgpV4Interface from protobuf object *otg.BgpV4Interface
	FromProto(msg *otg.BgpV4Interface) (BgpV4Interface, error)
	// FromPbText unmarshals BgpV4Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV4Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV4Interface from JSON text
	FromJson(value string) error
}

func (obj *bgpV4Interface) Marshal() marshalBgpV4Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV4Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV4Interface) Unmarshal() unMarshalBgpV4Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV4Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV4Interface) ToProto() (*otg.BgpV4Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV4Interface) FromProto(msg *otg.BgpV4Interface) (BgpV4Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV4Interface) ToPbText() (string, error) {
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

func (m *unMarshalbgpV4Interface) FromPbText(value string) error {
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

func (m *marshalbgpV4Interface) ToYaml() (string, error) {
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

func (m *unMarshalbgpV4Interface) FromYaml(value string) error {
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

func (m *marshalbgpV4Interface) ToJson() (string, error) {
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

func (m *unMarshalbgpV4Interface) FromJson(value string) error {
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

func (obj *bgpV4Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV4Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV4Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV4Interface) Clone() (BgpV4Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV4Interface()
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

func (obj *bgpV4Interface) setNil() {
	obj.peersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpV4Interface is configuration for emulated BGPv4 peers and routes on a single IPv4 interface.
type BgpV4Interface interface {
	Validation
	// msg marshals BgpV4Interface to protobuf object *otg.BgpV4Interface
	// and doesn't set defaults
	msg() *otg.BgpV4Interface
	// setMsg unmarshals BgpV4Interface from protobuf object *otg.BgpV4Interface
	// and doesn't set defaults
	setMsg(*otg.BgpV4Interface) BgpV4Interface
	// provides marshal interface
	Marshal() marshalBgpV4Interface
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV4Interface
	// validate validates BgpV4Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV4Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Name returns string, set in BgpV4Interface.
	Ipv4Name() string
	// SetIpv4Name assigns string provided by user to BgpV4Interface
	SetIpv4Name(value string) BgpV4Interface
	// Peers returns BgpV4InterfaceBgpV4PeerIterIter, set in BgpV4Interface
	Peers() BgpV4InterfaceBgpV4PeerIter
	setNil()
}

// The unique name of the IPv4,  Loopback IPv4 interface or DHCPv4 client used as the source IP for this list of BGP peers.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// Ipv4Name returns a string
func (obj *bgpV4Interface) Ipv4Name() string {

	return *obj.obj.Ipv4Name

}

// The unique name of the IPv4,  Loopback IPv4 interface or DHCPv4 client used as the source IP for this list of BGP peers.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// SetIpv4Name sets the string value in the BgpV4Interface object
func (obj *bgpV4Interface) SetIpv4Name(value string) BgpV4Interface {

	obj.obj.Ipv4Name = &value
	return obj
}

// This contains the list of BGPv4 peers configured on this interface.
// Peers returns a []BgpV4Peer
func (obj *bgpV4Interface) Peers() BgpV4InterfaceBgpV4PeerIter {
	if len(obj.obj.Peers) == 0 {
		obj.obj.Peers = []*otg.BgpV4Peer{}
	}
	if obj.peersHolder == nil {
		obj.peersHolder = newBgpV4InterfaceBgpV4PeerIter(&obj.obj.Peers).setMsg(obj)
	}
	return obj.peersHolder
}

type bgpV4InterfaceBgpV4PeerIter struct {
	obj            *bgpV4Interface
	bgpV4PeerSlice []BgpV4Peer
	fieldPtr       *[]*otg.BgpV4Peer
}

func newBgpV4InterfaceBgpV4PeerIter(ptr *[]*otg.BgpV4Peer) BgpV4InterfaceBgpV4PeerIter {
	return &bgpV4InterfaceBgpV4PeerIter{fieldPtr: ptr}
}

type BgpV4InterfaceBgpV4PeerIter interface {
	setMsg(*bgpV4Interface) BgpV4InterfaceBgpV4PeerIter
	Items() []BgpV4Peer
	Add() BgpV4Peer
	Append(items ...BgpV4Peer) BgpV4InterfaceBgpV4PeerIter
	Set(index int, newObj BgpV4Peer) BgpV4InterfaceBgpV4PeerIter
	Clear() BgpV4InterfaceBgpV4PeerIter
	clearHolderSlice() BgpV4InterfaceBgpV4PeerIter
	appendHolderSlice(item BgpV4Peer) BgpV4InterfaceBgpV4PeerIter
}

func (obj *bgpV4InterfaceBgpV4PeerIter) setMsg(msg *bgpV4Interface) BgpV4InterfaceBgpV4PeerIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV4Peer{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4InterfaceBgpV4PeerIter) Items() []BgpV4Peer {
	return obj.bgpV4PeerSlice
}

func (obj *bgpV4InterfaceBgpV4PeerIter) Add() BgpV4Peer {
	newObj := &otg.BgpV4Peer{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV4Peer{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV4PeerSlice = append(obj.bgpV4PeerSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4InterfaceBgpV4PeerIter) Append(items ...BgpV4Peer) BgpV4InterfaceBgpV4PeerIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV4PeerSlice = append(obj.bgpV4PeerSlice, item)
	}
	return obj
}

func (obj *bgpV4InterfaceBgpV4PeerIter) Set(index int, newObj BgpV4Peer) BgpV4InterfaceBgpV4PeerIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV4PeerSlice[index] = newObj
	return obj
}
func (obj *bgpV4InterfaceBgpV4PeerIter) Clear() BgpV4InterfaceBgpV4PeerIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV4Peer{}
		obj.bgpV4PeerSlice = []BgpV4Peer{}
	}
	return obj
}
func (obj *bgpV4InterfaceBgpV4PeerIter) clearHolderSlice() BgpV4InterfaceBgpV4PeerIter {
	if len(obj.bgpV4PeerSlice) > 0 {
		obj.bgpV4PeerSlice = []BgpV4Peer{}
	}
	return obj
}
func (obj *bgpV4InterfaceBgpV4PeerIter) appendHolderSlice(item BgpV4Peer) BgpV4InterfaceBgpV4PeerIter {
	obj.bgpV4PeerSlice = append(obj.bgpV4PeerSlice, item)
	return obj
}

func (obj *bgpV4Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Ipv4Name is required
	if obj.obj.Ipv4Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4Name is required field on interface BgpV4Interface")
	}

	if len(obj.obj.Peers) != 0 {

		if set_default {
			obj.Peers().clearHolderSlice()
			for _, item := range obj.obj.Peers {
				obj.Peers().appendHolderSlice(&bgpV4Peer{obj: item})
			}
		}
		for _, item := range obj.Peers().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpV4Interface) setDefault() {

}
