package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2V4Interface *****
type roCEv2V4Interface struct {
	validation
	obj          *otg.RoCEv2V4Interface
	marshaller   marshalRoCEv2V4Interface
	unMarshaller unMarshalRoCEv2V4Interface
	peersHolder  RoCEv2V4InterfaceRoCEv2V4PeerIter
}

func NewRoCEv2V4Interface() RoCEv2V4Interface {
	obj := roCEv2V4Interface{obj: &otg.RoCEv2V4Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2V4Interface) msg() *otg.RoCEv2V4Interface {
	return obj.obj
}

func (obj *roCEv2V4Interface) setMsg(msg *otg.RoCEv2V4Interface) RoCEv2V4Interface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2V4Interface struct {
	obj *roCEv2V4Interface
}

type marshalRoCEv2V4Interface interface {
	// ToProto marshals RoCEv2V4Interface to protobuf object *otg.RoCEv2V4Interface
	ToProto() (*otg.RoCEv2V4Interface, error)
	// ToPbText marshals RoCEv2V4Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2V4Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2V4Interface to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2V4Interface struct {
	obj *roCEv2V4Interface
}

type unMarshalRoCEv2V4Interface interface {
	// FromProto unmarshals RoCEv2V4Interface from protobuf object *otg.RoCEv2V4Interface
	FromProto(msg *otg.RoCEv2V4Interface) (RoCEv2V4Interface, error)
	// FromPbText unmarshals RoCEv2V4Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2V4Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2V4Interface from JSON text
	FromJson(value string) error
}

func (obj *roCEv2V4Interface) Marshal() marshalRoCEv2V4Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2V4Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2V4Interface) Unmarshal() unMarshalRoCEv2V4Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2V4Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2V4Interface) ToProto() (*otg.RoCEv2V4Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2V4Interface) FromProto(msg *otg.RoCEv2V4Interface) (RoCEv2V4Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2V4Interface) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2V4Interface) FromPbText(value string) error {
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

func (m *marshalroCEv2V4Interface) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2V4Interface) FromYaml(value string) error {
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

func (m *marshalroCEv2V4Interface) ToJson() (string, error) {
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

func (m *unMarshalroCEv2V4Interface) FromJson(value string) error {
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

func (obj *roCEv2V4Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2V4Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2V4Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2V4Interface) Clone() (RoCEv2V4Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2V4Interface()
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

func (obj *roCEv2V4Interface) setNil() {
	obj.peersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RoCEv2V4Interface is configuration for RoCEv2 IPv4 peers.
type RoCEv2V4Interface interface {
	Validation
	// msg marshals RoCEv2V4Interface to protobuf object *otg.RoCEv2V4Interface
	// and doesn't set defaults
	msg() *otg.RoCEv2V4Interface
	// setMsg unmarshals RoCEv2V4Interface from protobuf object *otg.RoCEv2V4Interface
	// and doesn't set defaults
	setMsg(*otg.RoCEv2V4Interface) RoCEv2V4Interface
	// provides marshal interface
	Marshal() marshalRoCEv2V4Interface
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2V4Interface
	// validate validates RoCEv2V4Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2V4Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Name returns string, set in RoCEv2V4Interface.
	Ipv4Name() string
	// SetIpv4Name assigns string provided by user to RoCEv2V4Interface
	SetIpv4Name(value string) RoCEv2V4Interface
	// Peers returns RoCEv2V4InterfaceRoCEv2V4PeerIterIter, set in RoCEv2V4Interface
	Peers() RoCEv2V4InterfaceRoCEv2V4PeerIter
	setNil()
}

// The unique name of the IPv4, used as the source IP for this list of RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// Ipv4Name returns a string
func (obj *roCEv2V4Interface) Ipv4Name() string {

	return *obj.obj.Ipv4Name

}

// The unique name of the IPv4, used as the source IP for this list of RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// SetIpv4Name sets the string value in the RoCEv2V4Interface object
func (obj *roCEv2V4Interface) SetIpv4Name(value string) RoCEv2V4Interface {

	obj.obj.Ipv4Name = &value
	return obj
}

// This contains the list of RoCEv2 peers configured on this interface.
// Peers returns a []RoCEv2V4Peer
func (obj *roCEv2V4Interface) Peers() RoCEv2V4InterfaceRoCEv2V4PeerIter {
	if len(obj.obj.Peers) == 0 {
		obj.obj.Peers = []*otg.RoCEv2V4Peer{}
	}
	if obj.peersHolder == nil {
		obj.peersHolder = newRoCEv2V4InterfaceRoCEv2V4PeerIter(&obj.obj.Peers).setMsg(obj)
	}
	return obj.peersHolder
}

type roCEv2V4InterfaceRoCEv2V4PeerIter struct {
	obj               *roCEv2V4Interface
	roCEv2V4PeerSlice []RoCEv2V4Peer
	fieldPtr          *[]*otg.RoCEv2V4Peer
}

func newRoCEv2V4InterfaceRoCEv2V4PeerIter(ptr *[]*otg.RoCEv2V4Peer) RoCEv2V4InterfaceRoCEv2V4PeerIter {
	return &roCEv2V4InterfaceRoCEv2V4PeerIter{fieldPtr: ptr}
}

type RoCEv2V4InterfaceRoCEv2V4PeerIter interface {
	setMsg(*roCEv2V4Interface) RoCEv2V4InterfaceRoCEv2V4PeerIter
	Items() []RoCEv2V4Peer
	Add() RoCEv2V4Peer
	Append(items ...RoCEv2V4Peer) RoCEv2V4InterfaceRoCEv2V4PeerIter
	Set(index int, newObj RoCEv2V4Peer) RoCEv2V4InterfaceRoCEv2V4PeerIter
	Clear() RoCEv2V4InterfaceRoCEv2V4PeerIter
	clearHolderSlice() RoCEv2V4InterfaceRoCEv2V4PeerIter
	appendHolderSlice(item RoCEv2V4Peer) RoCEv2V4InterfaceRoCEv2V4PeerIter
}

func (obj *roCEv2V4InterfaceRoCEv2V4PeerIter) setMsg(msg *roCEv2V4Interface) RoCEv2V4InterfaceRoCEv2V4PeerIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&roCEv2V4Peer{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *roCEv2V4InterfaceRoCEv2V4PeerIter) Items() []RoCEv2V4Peer {
	return obj.roCEv2V4PeerSlice
}

func (obj *roCEv2V4InterfaceRoCEv2V4PeerIter) Add() RoCEv2V4Peer {
	newObj := &otg.RoCEv2V4Peer{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &roCEv2V4Peer{obj: newObj}
	newLibObj.setDefault()
	obj.roCEv2V4PeerSlice = append(obj.roCEv2V4PeerSlice, newLibObj)
	return newLibObj
}

func (obj *roCEv2V4InterfaceRoCEv2V4PeerIter) Append(items ...RoCEv2V4Peer) RoCEv2V4InterfaceRoCEv2V4PeerIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.roCEv2V4PeerSlice = append(obj.roCEv2V4PeerSlice, item)
	}
	return obj
}

func (obj *roCEv2V4InterfaceRoCEv2V4PeerIter) Set(index int, newObj RoCEv2V4Peer) RoCEv2V4InterfaceRoCEv2V4PeerIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.roCEv2V4PeerSlice[index] = newObj
	return obj
}
func (obj *roCEv2V4InterfaceRoCEv2V4PeerIter) Clear() RoCEv2V4InterfaceRoCEv2V4PeerIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RoCEv2V4Peer{}
		obj.roCEv2V4PeerSlice = []RoCEv2V4Peer{}
	}
	return obj
}
func (obj *roCEv2V4InterfaceRoCEv2V4PeerIter) clearHolderSlice() RoCEv2V4InterfaceRoCEv2V4PeerIter {
	if len(obj.roCEv2V4PeerSlice) > 0 {
		obj.roCEv2V4PeerSlice = []RoCEv2V4Peer{}
	}
	return obj
}
func (obj *roCEv2V4InterfaceRoCEv2V4PeerIter) appendHolderSlice(item RoCEv2V4Peer) RoCEv2V4InterfaceRoCEv2V4PeerIter {
	obj.roCEv2V4PeerSlice = append(obj.roCEv2V4PeerSlice, item)
	return obj
}

func (obj *roCEv2V4Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Ipv4Name is required
	if obj.obj.Ipv4Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4Name is required field on interface RoCEv2V4Interface")
	}

	if len(obj.obj.Peers) != 0 {

		if set_default {
			obj.Peers().clearHolderSlice()
			for _, item := range obj.obj.Peers {
				obj.Peers().appendHolderSlice(&roCEv2V4Peer{obj: item})
			}
		}
		for _, item := range obj.Peers().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *roCEv2V4Interface) setDefault() {

}
