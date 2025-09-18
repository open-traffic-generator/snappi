package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2V4Peer *****
type rocev2V4Peer struct {
	validation
	obj          *otg.Rocev2V4Peer
	marshaller   marshalRocev2V4Peer
	unMarshaller unMarshalRocev2V4Peer
	qpsHolder    Rocev2V4PeerRocev2QPsIter
}

func NewRocev2V4Peer() Rocev2V4Peer {
	obj := rocev2V4Peer{obj: &otg.Rocev2V4Peer{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2V4Peer) msg() *otg.Rocev2V4Peer {
	return obj.obj
}

func (obj *rocev2V4Peer) setMsg(msg *otg.Rocev2V4Peer) Rocev2V4Peer {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2V4Peer struct {
	obj *rocev2V4Peer
}

type marshalRocev2V4Peer interface {
	// ToProto marshals Rocev2V4Peer to protobuf object *otg.Rocev2V4Peer
	ToProto() (*otg.Rocev2V4Peer, error)
	// ToPbText marshals Rocev2V4Peer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2V4Peer to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2V4Peer to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2V4Peer struct {
	obj *rocev2V4Peer
}

type unMarshalRocev2V4Peer interface {
	// FromProto unmarshals Rocev2V4Peer from protobuf object *otg.Rocev2V4Peer
	FromProto(msg *otg.Rocev2V4Peer) (Rocev2V4Peer, error)
	// FromPbText unmarshals Rocev2V4Peer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2V4Peer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2V4Peer from JSON text
	FromJson(value string) error
}

func (obj *rocev2V4Peer) Marshal() marshalRocev2V4Peer {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2V4Peer{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2V4Peer) Unmarshal() unMarshalRocev2V4Peer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2V4Peer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2V4Peer) ToProto() (*otg.Rocev2V4Peer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2V4Peer) FromProto(msg *otg.Rocev2V4Peer) (Rocev2V4Peer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2V4Peer) ToPbText() (string, error) {
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

func (m *unMarshalrocev2V4Peer) FromPbText(value string) error {
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

func (m *marshalrocev2V4Peer) ToYaml() (string, error) {
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

func (m *unMarshalrocev2V4Peer) FromYaml(value string) error {
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

func (m *marshalrocev2V4Peer) ToJson() (string, error) {
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

func (m *unMarshalrocev2V4Peer) FromJson(value string) error {
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

func (obj *rocev2V4Peer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2V4Peer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2V4Peer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2V4Peer) Clone() (Rocev2V4Peer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2V4Peer()
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

func (obj *rocev2V4Peer) setNil() {
	obj.qpsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2V4Peer is configuration for RoCEv2 IPv4 peers.
type Rocev2V4Peer interface {
	Validation
	// msg marshals Rocev2V4Peer to protobuf object *otg.Rocev2V4Peer
	// and doesn't set defaults
	msg() *otg.Rocev2V4Peer
	// setMsg unmarshals Rocev2V4Peer from protobuf object *otg.Rocev2V4Peer
	// and doesn't set defaults
	setMsg(*otg.Rocev2V4Peer) Rocev2V4Peer
	// provides marshal interface
	Marshal() marshalRocev2V4Peer
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2V4Peer
	// validate validates Rocev2V4Peer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2V4Peer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Rocev2V4Peer.
	Name() string
	// SetName assigns string provided by user to Rocev2V4Peer
	SetName(value string) Rocev2V4Peer
	// DestinationIpAddress returns string, set in Rocev2V4Peer.
	DestinationIpAddress() string
	// SetDestinationIpAddress assigns string provided by user to Rocev2V4Peer
	SetDestinationIpAddress(value string) Rocev2V4Peer
	// Qps returns Rocev2V4PeerRocev2QPsIterIter, set in Rocev2V4Peer
	Qps() Rocev2V4PeerRocev2QPsIter
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *rocev2V4Peer) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Rocev2V4Peer object
func (obj *rocev2V4Peer) SetName(value string) Rocev2V4Peer {

	obj.obj.Name = &value
	return obj
}

// Specify the destination ip address.
// DestinationIpAddress returns a string
func (obj *rocev2V4Peer) DestinationIpAddress() string {

	return *obj.obj.DestinationIpAddress

}

// Specify the destination ip address.
// SetDestinationIpAddress sets the string value in the Rocev2V4Peer object
func (obj *rocev2V4Peer) SetDestinationIpAddress(value string) Rocev2V4Peer {

	obj.obj.DestinationIpAddress = &value
	return obj
}

// This allows the user to set  multiple QPs and its properties between a pair of source and destination RoCEv2 devices.
// Qps returns a []Rocev2QPs
func (obj *rocev2V4Peer) Qps() Rocev2V4PeerRocev2QPsIter {
	if len(obj.obj.Qps) == 0 {
		obj.obj.Qps = []*otg.Rocev2QPs{}
	}
	if obj.qpsHolder == nil {
		obj.qpsHolder = newRocev2V4PeerRocev2QPsIter(&obj.obj.Qps).setMsg(obj)
	}
	return obj.qpsHolder
}

type rocev2V4PeerRocev2QPsIter struct {
	obj            *rocev2V4Peer
	rocev2QPsSlice []Rocev2QPs
	fieldPtr       *[]*otg.Rocev2QPs
}

func newRocev2V4PeerRocev2QPsIter(ptr *[]*otg.Rocev2QPs) Rocev2V4PeerRocev2QPsIter {
	return &rocev2V4PeerRocev2QPsIter{fieldPtr: ptr}
}

type Rocev2V4PeerRocev2QPsIter interface {
	setMsg(*rocev2V4Peer) Rocev2V4PeerRocev2QPsIter
	Items() []Rocev2QPs
	Add() Rocev2QPs
	Append(items ...Rocev2QPs) Rocev2V4PeerRocev2QPsIter
	Set(index int, newObj Rocev2QPs) Rocev2V4PeerRocev2QPsIter
	Clear() Rocev2V4PeerRocev2QPsIter
	clearHolderSlice() Rocev2V4PeerRocev2QPsIter
	appendHolderSlice(item Rocev2QPs) Rocev2V4PeerRocev2QPsIter
}

func (obj *rocev2V4PeerRocev2QPsIter) setMsg(msg *rocev2V4Peer) Rocev2V4PeerRocev2QPsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rocev2QPs{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *rocev2V4PeerRocev2QPsIter) Items() []Rocev2QPs {
	return obj.rocev2QPsSlice
}

func (obj *rocev2V4PeerRocev2QPsIter) Add() Rocev2QPs {
	newObj := &otg.Rocev2QPs{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rocev2QPs{obj: newObj}
	newLibObj.setDefault()
	obj.rocev2QPsSlice = append(obj.rocev2QPsSlice, newLibObj)
	return newLibObj
}

func (obj *rocev2V4PeerRocev2QPsIter) Append(items ...Rocev2QPs) Rocev2V4PeerRocev2QPsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rocev2QPsSlice = append(obj.rocev2QPsSlice, item)
	}
	return obj
}

func (obj *rocev2V4PeerRocev2QPsIter) Set(index int, newObj Rocev2QPs) Rocev2V4PeerRocev2QPsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rocev2QPsSlice[index] = newObj
	return obj
}
func (obj *rocev2V4PeerRocev2QPsIter) Clear() Rocev2V4PeerRocev2QPsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Rocev2QPs{}
		obj.rocev2QPsSlice = []Rocev2QPs{}
	}
	return obj
}
func (obj *rocev2V4PeerRocev2QPsIter) clearHolderSlice() Rocev2V4PeerRocev2QPsIter {
	if len(obj.rocev2QPsSlice) > 0 {
		obj.rocev2QPsSlice = []Rocev2QPs{}
	}
	return obj
}
func (obj *rocev2V4PeerRocev2QPsIter) appendHolderSlice(item Rocev2QPs) Rocev2V4PeerRocev2QPsIter {
	obj.rocev2QPsSlice = append(obj.rocev2QPsSlice, item)
	return obj
}

func (obj *rocev2V4Peer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Rocev2V4Peer")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"Rocev2V4Peer.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	// DestinationIpAddress is required
	if obj.obj.DestinationIpAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "DestinationIpAddress is required field on interface Rocev2V4Peer")
	}
	if obj.obj.DestinationIpAddress != nil {

		err := obj.validateIpv4(obj.DestinationIpAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Rocev2V4Peer.DestinationIpAddress"))
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

func (obj *rocev2V4Peer) setDefault() {

}
