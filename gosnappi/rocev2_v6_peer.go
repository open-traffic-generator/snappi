package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2V6Peer *****
type rocev2V6Peer struct {
	validation
	obj          *otg.Rocev2V6Peer
	marshaller   marshalRocev2V6Peer
	unMarshaller unMarshalRocev2V6Peer
	qpsHolder    Rocev2V6PeerRocev2QPsIter
}

func NewRocev2V6Peer() Rocev2V6Peer {
	obj := rocev2V6Peer{obj: &otg.Rocev2V6Peer{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2V6Peer) msg() *otg.Rocev2V6Peer {
	return obj.obj
}

func (obj *rocev2V6Peer) setMsg(msg *otg.Rocev2V6Peer) Rocev2V6Peer {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2V6Peer struct {
	obj *rocev2V6Peer
}

type marshalRocev2V6Peer interface {
	// ToProto marshals Rocev2V6Peer to protobuf object *otg.Rocev2V6Peer
	ToProto() (*otg.Rocev2V6Peer, error)
	// ToPbText marshals Rocev2V6Peer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2V6Peer to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2V6Peer to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2V6Peer to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2V6Peer struct {
	obj *rocev2V6Peer
}

type unMarshalRocev2V6Peer interface {
	// FromProto unmarshals Rocev2V6Peer from protobuf object *otg.Rocev2V6Peer
	FromProto(msg *otg.Rocev2V6Peer) (Rocev2V6Peer, error)
	// FromPbText unmarshals Rocev2V6Peer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2V6Peer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2V6Peer from JSON text
	FromJson(value string) error
}

func (obj *rocev2V6Peer) Marshal() marshalRocev2V6Peer {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2V6Peer{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2V6Peer) Unmarshal() unMarshalRocev2V6Peer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2V6Peer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2V6Peer) ToProto() (*otg.Rocev2V6Peer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2V6Peer) FromProto(msg *otg.Rocev2V6Peer) (Rocev2V6Peer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2V6Peer) ToPbText() (string, error) {
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

func (m *unMarshalrocev2V6Peer) FromPbText(value string) error {
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

func (m *marshalrocev2V6Peer) ToYaml() (string, error) {
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

func (m *unMarshalrocev2V6Peer) FromYaml(value string) error {
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

func (m *marshalrocev2V6Peer) ToJsonRaw() (string, error) {
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

func (m *marshalrocev2V6Peer) ToJson() (string, error) {
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

func (m *unMarshalrocev2V6Peer) FromJson(value string) error {
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

func (obj *rocev2V6Peer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2V6Peer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2V6Peer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2V6Peer) Clone() (Rocev2V6Peer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2V6Peer()
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

func (obj *rocev2V6Peer) setNil() {
	obj.qpsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2V6Peer is configuration for RoCEv2 IPv6 peer settings.
type Rocev2V6Peer interface {
	Validation
	// msg marshals Rocev2V6Peer to protobuf object *otg.Rocev2V6Peer
	// and doesn't set defaults
	msg() *otg.Rocev2V6Peer
	// setMsg unmarshals Rocev2V6Peer from protobuf object *otg.Rocev2V6Peer
	// and doesn't set defaults
	setMsg(*otg.Rocev2V6Peer) Rocev2V6Peer
	// provides marshal interface
	Marshal() marshalRocev2V6Peer
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2V6Peer
	// validate validates Rocev2V6Peer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2V6Peer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Rocev2V6Peer.
	Name() string
	// SetName assigns string provided by user to Rocev2V6Peer
	SetName(value string) Rocev2V6Peer
	// DestinationIpAddress returns string, set in Rocev2V6Peer.
	DestinationIpAddress() string
	// SetDestinationIpAddress assigns string provided by user to Rocev2V6Peer
	SetDestinationIpAddress(value string) Rocev2V6Peer
	// Qps returns Rocev2V6PeerRocev2QPsIterIter, set in Rocev2V6Peer
	Qps() Rocev2V6PeerRocev2QPsIter
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *rocev2V6Peer) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Rocev2V6Peer object
func (obj *rocev2V6Peer) SetName(value string) Rocev2V6Peer {

	obj.obj.Name = &value
	return obj
}

// Specify the destination ip address.
// DestinationIpAddress returns a string
func (obj *rocev2V6Peer) DestinationIpAddress() string {

	return *obj.obj.DestinationIpAddress

}

// Specify the destination ip address.
// SetDestinationIpAddress sets the string value in the Rocev2V6Peer object
func (obj *rocev2V6Peer) SetDestinationIpAddress(value string) Rocev2V6Peer {

	obj.obj.DestinationIpAddress = &value
	return obj
}

// This allows the user to set  multiple QPs and its properties between a pair of source and destination RoCEv2 devices.
// Qps returns a []Rocev2QPs
func (obj *rocev2V6Peer) Qps() Rocev2V6PeerRocev2QPsIter {
	if len(obj.obj.Qps) == 0 {
		obj.obj.Qps = []*otg.Rocev2QPs{}
	}
	if obj.qpsHolder == nil {
		obj.qpsHolder = newRocev2V6PeerRocev2QPsIter(&obj.obj.Qps).setMsg(obj)
	}
	return obj.qpsHolder
}

type rocev2V6PeerRocev2QPsIter struct {
	obj            *rocev2V6Peer
	rocev2QPsSlice []Rocev2QPs
	fieldPtr       *[]*otg.Rocev2QPs
}

func newRocev2V6PeerRocev2QPsIter(ptr *[]*otg.Rocev2QPs) Rocev2V6PeerRocev2QPsIter {
	return &rocev2V6PeerRocev2QPsIter{fieldPtr: ptr}
}

type Rocev2V6PeerRocev2QPsIter interface {
	setMsg(*rocev2V6Peer) Rocev2V6PeerRocev2QPsIter
	Items() []Rocev2QPs
	Add() Rocev2QPs
	Append(items ...Rocev2QPs) Rocev2V6PeerRocev2QPsIter
	Set(index int, newObj Rocev2QPs) Rocev2V6PeerRocev2QPsIter
	Clear() Rocev2V6PeerRocev2QPsIter
	clearHolderSlice() Rocev2V6PeerRocev2QPsIter
	appendHolderSlice(item Rocev2QPs) Rocev2V6PeerRocev2QPsIter
}

func (obj *rocev2V6PeerRocev2QPsIter) setMsg(msg *rocev2V6Peer) Rocev2V6PeerRocev2QPsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rocev2QPs{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *rocev2V6PeerRocev2QPsIter) Items() []Rocev2QPs {
	return obj.rocev2QPsSlice
}

func (obj *rocev2V6PeerRocev2QPsIter) Add() Rocev2QPs {
	newObj := &otg.Rocev2QPs{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rocev2QPs{obj: newObj}
	newLibObj.setDefault()
	obj.rocev2QPsSlice = append(obj.rocev2QPsSlice, newLibObj)
	return newLibObj
}

func (obj *rocev2V6PeerRocev2QPsIter) Append(items ...Rocev2QPs) Rocev2V6PeerRocev2QPsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rocev2QPsSlice = append(obj.rocev2QPsSlice, item)
	}
	return obj
}

func (obj *rocev2V6PeerRocev2QPsIter) Set(index int, newObj Rocev2QPs) Rocev2V6PeerRocev2QPsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rocev2QPsSlice[index] = newObj
	return obj
}
func (obj *rocev2V6PeerRocev2QPsIter) Clear() Rocev2V6PeerRocev2QPsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Rocev2QPs{}
		obj.rocev2QPsSlice = []Rocev2QPs{}
	}
	return obj
}
func (obj *rocev2V6PeerRocev2QPsIter) clearHolderSlice() Rocev2V6PeerRocev2QPsIter {
	if len(obj.rocev2QPsSlice) > 0 {
		obj.rocev2QPsSlice = []Rocev2QPs{}
	}
	return obj
}
func (obj *rocev2V6PeerRocev2QPsIter) appendHolderSlice(item Rocev2QPs) Rocev2V6PeerRocev2QPsIter {
	obj.rocev2QPsSlice = append(obj.rocev2QPsSlice, item)
	return obj
}

func (obj *rocev2V6Peer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Rocev2V6Peer")
	}

	// DestinationIpAddress is required
	if obj.obj.DestinationIpAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "DestinationIpAddress is required field on interface Rocev2V6Peer")
	}
	if obj.obj.DestinationIpAddress != nil {

		err := obj.validateIpv6(obj.DestinationIpAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Rocev2V6Peer.DestinationIpAddress"))
		}

	}

	if len(obj.obj.Qps) != 0 {

		if set_default {
			obj.Qps().clearHolderSlice()
			for _, item := range obj.obj.Qps {
				obj.Qps().appendHolderSlice(&rocev2QPs{obj: item})
			}
		}
		for _, item := range obj.Qps().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *rocev2V6Peer) setDefault() {

}
