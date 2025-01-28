package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2V6Interface *****
type roCEv2V6Interface struct {
	validation
	obj          *otg.RoCEv2V6Interface
	marshaller   marshalRoCEv2V6Interface
	unMarshaller unMarshalRoCEv2V6Interface
	peersHolder  RoCEv2V6InterfaceRoCEv2V6PeerIter
}

func NewRoCEv2V6Interface() RoCEv2V6Interface {
	obj := roCEv2V6Interface{obj: &otg.RoCEv2V6Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2V6Interface) msg() *otg.RoCEv2V6Interface {
	return obj.obj
}

func (obj *roCEv2V6Interface) setMsg(msg *otg.RoCEv2V6Interface) RoCEv2V6Interface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2V6Interface struct {
	obj *roCEv2V6Interface
}

type marshalRoCEv2V6Interface interface {
	// ToProto marshals RoCEv2V6Interface to protobuf object *otg.RoCEv2V6Interface
	ToProto() (*otg.RoCEv2V6Interface, error)
	// ToPbText marshals RoCEv2V6Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2V6Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2V6Interface to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2V6Interface struct {
	obj *roCEv2V6Interface
}

type unMarshalRoCEv2V6Interface interface {
	// FromProto unmarshals RoCEv2V6Interface from protobuf object *otg.RoCEv2V6Interface
	FromProto(msg *otg.RoCEv2V6Interface) (RoCEv2V6Interface, error)
	// FromPbText unmarshals RoCEv2V6Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2V6Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2V6Interface from JSON text
	FromJson(value string) error
}

func (obj *roCEv2V6Interface) Marshal() marshalRoCEv2V6Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2V6Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2V6Interface) Unmarshal() unMarshalRoCEv2V6Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2V6Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2V6Interface) ToProto() (*otg.RoCEv2V6Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2V6Interface) FromProto(msg *otg.RoCEv2V6Interface) (RoCEv2V6Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2V6Interface) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2V6Interface) FromPbText(value string) error {
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

func (m *marshalroCEv2V6Interface) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2V6Interface) FromYaml(value string) error {
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

func (m *marshalroCEv2V6Interface) ToJson() (string, error) {
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

func (m *unMarshalroCEv2V6Interface) FromJson(value string) error {
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

func (obj *roCEv2V6Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2V6Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2V6Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2V6Interface) Clone() (RoCEv2V6Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2V6Interface()
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

func (obj *roCEv2V6Interface) setNil() {
	obj.peersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RoCEv2V6Interface is configuration for RoCEv2 IPv6 peers.
type RoCEv2V6Interface interface {
	Validation
	// msg marshals RoCEv2V6Interface to protobuf object *otg.RoCEv2V6Interface
	// and doesn't set defaults
	msg() *otg.RoCEv2V6Interface
	// setMsg unmarshals RoCEv2V6Interface from protobuf object *otg.RoCEv2V6Interface
	// and doesn't set defaults
	setMsg(*otg.RoCEv2V6Interface) RoCEv2V6Interface
	// provides marshal interface
	Marshal() marshalRoCEv2V6Interface
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2V6Interface
	// validate validates RoCEv2V6Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2V6Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv6Name returns string, set in RoCEv2V6Interface.
	Ipv6Name() string
	// SetIpv6Name assigns string provided by user to RoCEv2V6Interface
	SetIpv6Name(value string) RoCEv2V6Interface
	// Peers returns RoCEv2V6InterfaceRoCEv2V6PeerIterIter, set in RoCEv2V6Interface
	Peers() RoCEv2V6InterfaceRoCEv2V6PeerIter
	setNil()
}

// The unique name of IPv6 used as the source IP for this list of RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// Ipv6Name returns a string
func (obj *roCEv2V6Interface) Ipv6Name() string {

	return *obj.obj.Ipv6Name

}

// The unique name of IPv6 used as the source IP for this list of RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// SetIpv6Name sets the string value in the RoCEv2V6Interface object
func (obj *roCEv2V6Interface) SetIpv6Name(value string) RoCEv2V6Interface {

	obj.obj.Ipv6Name = &value
	return obj
}

// This contains the list of RoCEv2 IPv6 peers configured on this interface.
// Peers returns a []RoCEv2V6Peer
func (obj *roCEv2V6Interface) Peers() RoCEv2V6InterfaceRoCEv2V6PeerIter {
	if len(obj.obj.Peers) == 0 {
		obj.obj.Peers = []*otg.RoCEv2V6Peer{}
	}
	if obj.peersHolder == nil {
		obj.peersHolder = newRoCEv2V6InterfaceRoCEv2V6PeerIter(&obj.obj.Peers).setMsg(obj)
	}
	return obj.peersHolder
}

type roCEv2V6InterfaceRoCEv2V6PeerIter struct {
	obj               *roCEv2V6Interface
	roCEv2V6PeerSlice []RoCEv2V6Peer
	fieldPtr          *[]*otg.RoCEv2V6Peer
}

func newRoCEv2V6InterfaceRoCEv2V6PeerIter(ptr *[]*otg.RoCEv2V6Peer) RoCEv2V6InterfaceRoCEv2V6PeerIter {
	return &roCEv2V6InterfaceRoCEv2V6PeerIter{fieldPtr: ptr}
}

type RoCEv2V6InterfaceRoCEv2V6PeerIter interface {
	setMsg(*roCEv2V6Interface) RoCEv2V6InterfaceRoCEv2V6PeerIter
	Items() []RoCEv2V6Peer
	Add() RoCEv2V6Peer
	Append(items ...RoCEv2V6Peer) RoCEv2V6InterfaceRoCEv2V6PeerIter
	Set(index int, newObj RoCEv2V6Peer) RoCEv2V6InterfaceRoCEv2V6PeerIter
	Clear() RoCEv2V6InterfaceRoCEv2V6PeerIter
	clearHolderSlice() RoCEv2V6InterfaceRoCEv2V6PeerIter
	appendHolderSlice(item RoCEv2V6Peer) RoCEv2V6InterfaceRoCEv2V6PeerIter
}

func (obj *roCEv2V6InterfaceRoCEv2V6PeerIter) setMsg(msg *roCEv2V6Interface) RoCEv2V6InterfaceRoCEv2V6PeerIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&roCEv2V6Peer{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *roCEv2V6InterfaceRoCEv2V6PeerIter) Items() []RoCEv2V6Peer {
	return obj.roCEv2V6PeerSlice
}

func (obj *roCEv2V6InterfaceRoCEv2V6PeerIter) Add() RoCEv2V6Peer {
	newObj := &otg.RoCEv2V6Peer{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &roCEv2V6Peer{obj: newObj}
	newLibObj.setDefault()
	obj.roCEv2V6PeerSlice = append(obj.roCEv2V6PeerSlice, newLibObj)
	return newLibObj
}

func (obj *roCEv2V6InterfaceRoCEv2V6PeerIter) Append(items ...RoCEv2V6Peer) RoCEv2V6InterfaceRoCEv2V6PeerIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.roCEv2V6PeerSlice = append(obj.roCEv2V6PeerSlice, item)
	}
	return obj
}

func (obj *roCEv2V6InterfaceRoCEv2V6PeerIter) Set(index int, newObj RoCEv2V6Peer) RoCEv2V6InterfaceRoCEv2V6PeerIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.roCEv2V6PeerSlice[index] = newObj
	return obj
}
func (obj *roCEv2V6InterfaceRoCEv2V6PeerIter) Clear() RoCEv2V6InterfaceRoCEv2V6PeerIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RoCEv2V6Peer{}
		obj.roCEv2V6PeerSlice = []RoCEv2V6Peer{}
	}
	return obj
}
func (obj *roCEv2V6InterfaceRoCEv2V6PeerIter) clearHolderSlice() RoCEv2V6InterfaceRoCEv2V6PeerIter {
	if len(obj.roCEv2V6PeerSlice) > 0 {
		obj.roCEv2V6PeerSlice = []RoCEv2V6Peer{}
	}
	return obj
}
func (obj *roCEv2V6InterfaceRoCEv2V6PeerIter) appendHolderSlice(item RoCEv2V6Peer) RoCEv2V6InterfaceRoCEv2V6PeerIter {
	obj.roCEv2V6PeerSlice = append(obj.roCEv2V6PeerSlice, item)
	return obj
}

func (obj *roCEv2V6Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Ipv6Name is required
	if obj.obj.Ipv6Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv6Name is required field on interface RoCEv2V6Interface")
	}

	if len(obj.obj.Peers) != 0 {

		if set_default {
			obj.Peers().clearHolderSlice()
			for _, item := range obj.obj.Peers {
				obj.Peers().appendHolderSlice(&roCEv2V6Peer{obj: item})
			}
		}
		for _, item := range obj.Peers().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *roCEv2V6Interface) setDefault() {

}
