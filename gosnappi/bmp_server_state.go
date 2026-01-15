package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BmpServerState *****
type bmpServerState struct {
	validation
	obj          *otg.BmpServerState
	marshaller   marshalBmpServerState
	unMarshaller unMarshalBmpServerState
	peersHolder  BmpServerStateBmpServerPeerStateIter
}

func NewBmpServerState() BmpServerState {
	obj := bmpServerState{obj: &otg.BmpServerState{}}
	obj.setDefault()
	return &obj
}

func (obj *bmpServerState) msg() *otg.BmpServerState {
	return obj.obj
}

func (obj *bmpServerState) setMsg(msg *otg.BmpServerState) BmpServerState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbmpServerState struct {
	obj *bmpServerState
}

type marshalBmpServerState interface {
	// ToProto marshals BmpServerState to protobuf object *otg.BmpServerState
	ToProto() (*otg.BmpServerState, error)
	// ToPbText marshals BmpServerState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BmpServerState to YAML text
	ToYaml() (string, error)
	// ToJson marshals BmpServerState to JSON text
	ToJson() (string, error)
}

type unMarshalbmpServerState struct {
	obj *bmpServerState
}

type unMarshalBmpServerState interface {
	// FromProto unmarshals BmpServerState from protobuf object *otg.BmpServerState
	FromProto(msg *otg.BmpServerState) (BmpServerState, error)
	// FromPbText unmarshals BmpServerState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BmpServerState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BmpServerState from JSON text
	FromJson(value string) error
}

func (obj *bmpServerState) Marshal() marshalBmpServerState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbmpServerState{obj: obj}
	}
	return obj.marshaller
}

func (obj *bmpServerState) Unmarshal() unMarshalBmpServerState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbmpServerState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbmpServerState) ToProto() (*otg.BmpServerState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbmpServerState) FromProto(msg *otg.BmpServerState) (BmpServerState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbmpServerState) ToPbText() (string, error) {
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

func (m *unMarshalbmpServerState) FromPbText(value string) error {
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

func (m *marshalbmpServerState) ToYaml() (string, error) {
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

func (m *unMarshalbmpServerState) FromYaml(value string) error {
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

func (m *marshalbmpServerState) ToJson() (string, error) {
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

func (m *unMarshalbmpServerState) FromJson(value string) error {
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

func (obj *bmpServerState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bmpServerState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bmpServerState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bmpServerState) Clone() (BmpServerState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBmpServerState()
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

func (obj *bmpServerState) setNil() {
	obj.peersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BmpServerState is the information learned by a specific BMP Server. Note that each BMP Server is connected to a single BMP client, though multiple BMP Servers can be configured on the  same IPv4 or IPv6 interface , each connected to a different BMP client. Each BMP client can in turn report information about multiple BGP sessions on that router.
type BmpServerState interface {
	Validation
	// msg marshals BmpServerState to protobuf object *otg.BmpServerState
	// and doesn't set defaults
	msg() *otg.BmpServerState
	// setMsg unmarshals BmpServerState from protobuf object *otg.BmpServerState
	// and doesn't set defaults
	setMsg(*otg.BmpServerState) BmpServerState
	// provides marshal interface
	Marshal() marshalBmpServerState
	// provides unmarshal interface
	Unmarshal() unMarshalBmpServerState
	// validate validates BmpServerState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BmpServerState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// BmpServerName returns string, set in BmpServerState.
	BmpServerName() string
	// SetBmpServerName assigns string provided by user to BmpServerState
	SetBmpServerName(value string) BmpServerState
	// HasBmpServerName checks if BmpServerName has been set in BmpServerState
	HasBmpServerName() bool
	// Peers returns BmpServerStateBmpServerPeerStateIterIter, set in BmpServerState
	Peers() BmpServerStateBmpServerPeerStateIter
	setNil()
}

// The name of the BMP Server configured for which the information is being returned.
// BmpServerName returns a string
func (obj *bmpServerState) BmpServerName() string {

	return *obj.obj.BmpServerName

}

// The name of the BMP Server configured for which the information is being returned.
// BmpServerName returns a string
func (obj *bmpServerState) HasBmpServerName() bool {
	return obj.obj.BmpServerName != nil
}

// The name of the BMP Server configured for which the information is being returned.
// SetBmpServerName sets the string value in the BmpServerState object
func (obj *bmpServerState) SetBmpServerName(value string) BmpServerState {

	obj.obj.BmpServerName = &value
	return obj
}

// description is TBD
// Peers returns a []BmpServerPeerState
func (obj *bmpServerState) Peers() BmpServerStateBmpServerPeerStateIter {
	if len(obj.obj.Peers) == 0 {
		obj.obj.Peers = []*otg.BmpServerPeerState{}
	}
	if obj.peersHolder == nil {
		obj.peersHolder = newBmpServerStateBmpServerPeerStateIter(&obj.obj.Peers).setMsg(obj)
	}
	return obj.peersHolder
}

type bmpServerStateBmpServerPeerStateIter struct {
	obj                     *bmpServerState
	bmpServerPeerStateSlice []BmpServerPeerState
	fieldPtr                *[]*otg.BmpServerPeerState
}

func newBmpServerStateBmpServerPeerStateIter(ptr *[]*otg.BmpServerPeerState) BmpServerStateBmpServerPeerStateIter {
	return &bmpServerStateBmpServerPeerStateIter{fieldPtr: ptr}
}

type BmpServerStateBmpServerPeerStateIter interface {
	setMsg(*bmpServerState) BmpServerStateBmpServerPeerStateIter
	Items() []BmpServerPeerState
	Add() BmpServerPeerState
	Append(items ...BmpServerPeerState) BmpServerStateBmpServerPeerStateIter
	Set(index int, newObj BmpServerPeerState) BmpServerStateBmpServerPeerStateIter
	Clear() BmpServerStateBmpServerPeerStateIter
	clearHolderSlice() BmpServerStateBmpServerPeerStateIter
	appendHolderSlice(item BmpServerPeerState) BmpServerStateBmpServerPeerStateIter
}

func (obj *bmpServerStateBmpServerPeerStateIter) setMsg(msg *bmpServerState) BmpServerStateBmpServerPeerStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bmpServerPeerState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bmpServerStateBmpServerPeerStateIter) Items() []BmpServerPeerState {
	return obj.bmpServerPeerStateSlice
}

func (obj *bmpServerStateBmpServerPeerStateIter) Add() BmpServerPeerState {
	newObj := &otg.BmpServerPeerState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bmpServerPeerState{obj: newObj}
	newLibObj.setDefault()
	obj.bmpServerPeerStateSlice = append(obj.bmpServerPeerStateSlice, newLibObj)
	return newLibObj
}

func (obj *bmpServerStateBmpServerPeerStateIter) Append(items ...BmpServerPeerState) BmpServerStateBmpServerPeerStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bmpServerPeerStateSlice = append(obj.bmpServerPeerStateSlice, item)
	}
	return obj
}

func (obj *bmpServerStateBmpServerPeerStateIter) Set(index int, newObj BmpServerPeerState) BmpServerStateBmpServerPeerStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bmpServerPeerStateSlice[index] = newObj
	return obj
}
func (obj *bmpServerStateBmpServerPeerStateIter) Clear() BmpServerStateBmpServerPeerStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BmpServerPeerState{}
		obj.bmpServerPeerStateSlice = []BmpServerPeerState{}
	}
	return obj
}
func (obj *bmpServerStateBmpServerPeerStateIter) clearHolderSlice() BmpServerStateBmpServerPeerStateIter {
	if len(obj.bmpServerPeerStateSlice) > 0 {
		obj.bmpServerPeerStateSlice = []BmpServerPeerState{}
	}
	return obj
}
func (obj *bmpServerStateBmpServerPeerStateIter) appendHolderSlice(item BmpServerPeerState) BmpServerStateBmpServerPeerStateIter {
	obj.bmpServerPeerStateSlice = append(obj.bmpServerPeerStateSlice, item)
	return obj
}

func (obj *bmpServerState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Peers) != 0 {

		if set_default {
			obj.Peers().clearHolderSlice()
			for _, item := range obj.obj.Peers {
				obj.Peers().appendHolderSlice(&bmpServerPeerState{obj: item})
			}
		}
		for _, item := range obj.Peers().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bmpServerState) setDefault() {

}
