package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixesState *****
type bgpPrefixesState struct {
	validation
	obj                           *otg.BgpPrefixesState
	marshaller                    marshalBgpPrefixesState
	unMarshaller                  unMarshalBgpPrefixesState
	ipv4UnicastPrefixesHolder     BgpPrefixesStateBgpPrefixIpv4UnicastStateIter
	ipv6UnicastPrefixesHolder     BgpPrefixesStateBgpPrefixIpv6UnicastStateIter
	ipv4MplsUnicastPrefixesHolder BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter
	ipv6MplsUnicastPrefixesHolder BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter
}

func NewBgpPrefixesState() BgpPrefixesState {
	obj := bgpPrefixesState{obj: &otg.BgpPrefixesState{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixesState) msg() *otg.BgpPrefixesState {
	return obj.obj
}

func (obj *bgpPrefixesState) setMsg(msg *otg.BgpPrefixesState) BgpPrefixesState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixesState struct {
	obj *bgpPrefixesState
}

type marshalBgpPrefixesState interface {
	// ToProto marshals BgpPrefixesState to protobuf object *otg.BgpPrefixesState
	ToProto() (*otg.BgpPrefixesState, error)
	// ToPbText marshals BgpPrefixesState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixesState to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixesState to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixesState struct {
	obj *bgpPrefixesState
}

type unMarshalBgpPrefixesState interface {
	// FromProto unmarshals BgpPrefixesState from protobuf object *otg.BgpPrefixesState
	FromProto(msg *otg.BgpPrefixesState) (BgpPrefixesState, error)
	// FromPbText unmarshals BgpPrefixesState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixesState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixesState from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixesState) Marshal() marshalBgpPrefixesState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixesState{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixesState) Unmarshal() unMarshalBgpPrefixesState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixesState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixesState) ToProto() (*otg.BgpPrefixesState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixesState) FromProto(msg *otg.BgpPrefixesState) (BgpPrefixesState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixesState) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixesState) FromPbText(value string) error {
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

func (m *marshalbgpPrefixesState) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixesState) FromYaml(value string) error {
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

func (m *marshalbgpPrefixesState) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixesState) FromJson(value string) error {
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

func (obj *bgpPrefixesState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixesState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixesState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixesState) Clone() (BgpPrefixesState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixesState()
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

func (obj *bgpPrefixesState) setNil() {
	obj.ipv4UnicastPrefixesHolder = nil
	obj.ipv6UnicastPrefixesHolder = nil
	obj.ipv4MplsUnicastPrefixesHolder = nil
	obj.ipv6MplsUnicastPrefixesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixesState is bGP peer prefixes.
type BgpPrefixesState interface {
	Validation
	// msg marshals BgpPrefixesState to protobuf object *otg.BgpPrefixesState
	// and doesn't set defaults
	msg() *otg.BgpPrefixesState
	// setMsg unmarshals BgpPrefixesState from protobuf object *otg.BgpPrefixesState
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixesState) BgpPrefixesState
	// provides marshal interface
	Marshal() marshalBgpPrefixesState
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixesState
	// validate validates BgpPrefixesState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixesState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// BgpPeerName returns string, set in BgpPrefixesState.
	BgpPeerName() string
	// SetBgpPeerName assigns string provided by user to BgpPrefixesState
	SetBgpPeerName(value string) BgpPrefixesState
	// HasBgpPeerName checks if BgpPeerName has been set in BgpPrefixesState
	HasBgpPeerName() bool
	// Ipv4UnicastPrefixes returns BgpPrefixesStateBgpPrefixIpv4UnicastStateIterIter, set in BgpPrefixesState
	Ipv4UnicastPrefixes() BgpPrefixesStateBgpPrefixIpv4UnicastStateIter
	// Ipv6UnicastPrefixes returns BgpPrefixesStateBgpPrefixIpv6UnicastStateIterIter, set in BgpPrefixesState
	Ipv6UnicastPrefixes() BgpPrefixesStateBgpPrefixIpv6UnicastStateIter
	// Ipv4MplsUnicastPrefixes returns BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIterIter, set in BgpPrefixesState
	Ipv4MplsUnicastPrefixes() BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter
	// Ipv6MplsUnicastPrefixes returns BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIterIter, set in BgpPrefixesState
	Ipv6MplsUnicastPrefixes() BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter
	setNil()
}

// The name of a BGP peer.
// BgpPeerName returns a string
func (obj *bgpPrefixesState) BgpPeerName() string {

	return *obj.obj.BgpPeerName

}

// The name of a BGP peer.
// BgpPeerName returns a string
func (obj *bgpPrefixesState) HasBgpPeerName() bool {
	return obj.obj.BgpPeerName != nil
}

// The name of a BGP peer.
// SetBgpPeerName sets the string value in the BgpPrefixesState object
func (obj *bgpPrefixesState) SetBgpPeerName(value string) BgpPrefixesState {

	obj.obj.BgpPeerName = &value
	return obj
}

// description is TBD
// Ipv4UnicastPrefixes returns a []BgpPrefixIpv4UnicastState
func (obj *bgpPrefixesState) Ipv4UnicastPrefixes() BgpPrefixesStateBgpPrefixIpv4UnicastStateIter {
	if len(obj.obj.Ipv4UnicastPrefixes) == 0 {
		obj.obj.Ipv4UnicastPrefixes = []*otg.BgpPrefixIpv4UnicastState{}
	}
	if obj.ipv4UnicastPrefixesHolder == nil {
		obj.ipv4UnicastPrefixesHolder = newBgpPrefixesStateBgpPrefixIpv4UnicastStateIter(&obj.obj.Ipv4UnicastPrefixes).setMsg(obj)
	}
	return obj.ipv4UnicastPrefixesHolder
}

type bgpPrefixesStateBgpPrefixIpv4UnicastStateIter struct {
	obj                            *bgpPrefixesState
	bgpPrefixIpv4UnicastStateSlice []BgpPrefixIpv4UnicastState
	fieldPtr                       *[]*otg.BgpPrefixIpv4UnicastState
}

func newBgpPrefixesStateBgpPrefixIpv4UnicastStateIter(ptr *[]*otg.BgpPrefixIpv4UnicastState) BgpPrefixesStateBgpPrefixIpv4UnicastStateIter {
	return &bgpPrefixesStateBgpPrefixIpv4UnicastStateIter{fieldPtr: ptr}
}

type BgpPrefixesStateBgpPrefixIpv4UnicastStateIter interface {
	setMsg(*bgpPrefixesState) BgpPrefixesStateBgpPrefixIpv4UnicastStateIter
	Items() []BgpPrefixIpv4UnicastState
	Add() BgpPrefixIpv4UnicastState
	Append(items ...BgpPrefixIpv4UnicastState) BgpPrefixesStateBgpPrefixIpv4UnicastStateIter
	Set(index int, newObj BgpPrefixIpv4UnicastState) BgpPrefixesStateBgpPrefixIpv4UnicastStateIter
	Clear() BgpPrefixesStateBgpPrefixIpv4UnicastStateIter
	clearHolderSlice() BgpPrefixesStateBgpPrefixIpv4UnicastStateIter
	appendHolderSlice(item BgpPrefixIpv4UnicastState) BgpPrefixesStateBgpPrefixIpv4UnicastStateIter
}

func (obj *bgpPrefixesStateBgpPrefixIpv4UnicastStateIter) setMsg(msg *bgpPrefixesState) BgpPrefixesStateBgpPrefixIpv4UnicastStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpPrefixIpv4UnicastState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixesStateBgpPrefixIpv4UnicastStateIter) Items() []BgpPrefixIpv4UnicastState {
	return obj.bgpPrefixIpv4UnicastStateSlice
}

func (obj *bgpPrefixesStateBgpPrefixIpv4UnicastStateIter) Add() BgpPrefixIpv4UnicastState {
	newObj := &otg.BgpPrefixIpv4UnicastState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpPrefixIpv4UnicastState{obj: newObj}
	newLibObj.setDefault()
	obj.bgpPrefixIpv4UnicastStateSlice = append(obj.bgpPrefixIpv4UnicastStateSlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixesStateBgpPrefixIpv4UnicastStateIter) Append(items ...BgpPrefixIpv4UnicastState) BgpPrefixesStateBgpPrefixIpv4UnicastStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpPrefixIpv4UnicastStateSlice = append(obj.bgpPrefixIpv4UnicastStateSlice, item)
	}
	return obj
}

func (obj *bgpPrefixesStateBgpPrefixIpv4UnicastStateIter) Set(index int, newObj BgpPrefixIpv4UnicastState) BgpPrefixesStateBgpPrefixIpv4UnicastStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpPrefixIpv4UnicastStateSlice[index] = newObj
	return obj
}
func (obj *bgpPrefixesStateBgpPrefixIpv4UnicastStateIter) Clear() BgpPrefixesStateBgpPrefixIpv4UnicastStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpPrefixIpv4UnicastState{}
		obj.bgpPrefixIpv4UnicastStateSlice = []BgpPrefixIpv4UnicastState{}
	}
	return obj
}
func (obj *bgpPrefixesStateBgpPrefixIpv4UnicastStateIter) clearHolderSlice() BgpPrefixesStateBgpPrefixIpv4UnicastStateIter {
	if len(obj.bgpPrefixIpv4UnicastStateSlice) > 0 {
		obj.bgpPrefixIpv4UnicastStateSlice = []BgpPrefixIpv4UnicastState{}
	}
	return obj
}
func (obj *bgpPrefixesStateBgpPrefixIpv4UnicastStateIter) appendHolderSlice(item BgpPrefixIpv4UnicastState) BgpPrefixesStateBgpPrefixIpv4UnicastStateIter {
	obj.bgpPrefixIpv4UnicastStateSlice = append(obj.bgpPrefixIpv4UnicastStateSlice, item)
	return obj
}

// description is TBD
// Ipv6UnicastPrefixes returns a []BgpPrefixIpv6UnicastState
func (obj *bgpPrefixesState) Ipv6UnicastPrefixes() BgpPrefixesStateBgpPrefixIpv6UnicastStateIter {
	if len(obj.obj.Ipv6UnicastPrefixes) == 0 {
		obj.obj.Ipv6UnicastPrefixes = []*otg.BgpPrefixIpv6UnicastState{}
	}
	if obj.ipv6UnicastPrefixesHolder == nil {
		obj.ipv6UnicastPrefixesHolder = newBgpPrefixesStateBgpPrefixIpv6UnicastStateIter(&obj.obj.Ipv6UnicastPrefixes).setMsg(obj)
	}
	return obj.ipv6UnicastPrefixesHolder
}

type bgpPrefixesStateBgpPrefixIpv6UnicastStateIter struct {
	obj                            *bgpPrefixesState
	bgpPrefixIpv6UnicastStateSlice []BgpPrefixIpv6UnicastState
	fieldPtr                       *[]*otg.BgpPrefixIpv6UnicastState
}

func newBgpPrefixesStateBgpPrefixIpv6UnicastStateIter(ptr *[]*otg.BgpPrefixIpv6UnicastState) BgpPrefixesStateBgpPrefixIpv6UnicastStateIter {
	return &bgpPrefixesStateBgpPrefixIpv6UnicastStateIter{fieldPtr: ptr}
}

type BgpPrefixesStateBgpPrefixIpv6UnicastStateIter interface {
	setMsg(*bgpPrefixesState) BgpPrefixesStateBgpPrefixIpv6UnicastStateIter
	Items() []BgpPrefixIpv6UnicastState
	Add() BgpPrefixIpv6UnicastState
	Append(items ...BgpPrefixIpv6UnicastState) BgpPrefixesStateBgpPrefixIpv6UnicastStateIter
	Set(index int, newObj BgpPrefixIpv6UnicastState) BgpPrefixesStateBgpPrefixIpv6UnicastStateIter
	Clear() BgpPrefixesStateBgpPrefixIpv6UnicastStateIter
	clearHolderSlice() BgpPrefixesStateBgpPrefixIpv6UnicastStateIter
	appendHolderSlice(item BgpPrefixIpv6UnicastState) BgpPrefixesStateBgpPrefixIpv6UnicastStateIter
}

func (obj *bgpPrefixesStateBgpPrefixIpv6UnicastStateIter) setMsg(msg *bgpPrefixesState) BgpPrefixesStateBgpPrefixIpv6UnicastStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpPrefixIpv6UnicastState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixesStateBgpPrefixIpv6UnicastStateIter) Items() []BgpPrefixIpv6UnicastState {
	return obj.bgpPrefixIpv6UnicastStateSlice
}

func (obj *bgpPrefixesStateBgpPrefixIpv6UnicastStateIter) Add() BgpPrefixIpv6UnicastState {
	newObj := &otg.BgpPrefixIpv6UnicastState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpPrefixIpv6UnicastState{obj: newObj}
	newLibObj.setDefault()
	obj.bgpPrefixIpv6UnicastStateSlice = append(obj.bgpPrefixIpv6UnicastStateSlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixesStateBgpPrefixIpv6UnicastStateIter) Append(items ...BgpPrefixIpv6UnicastState) BgpPrefixesStateBgpPrefixIpv6UnicastStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpPrefixIpv6UnicastStateSlice = append(obj.bgpPrefixIpv6UnicastStateSlice, item)
	}
	return obj
}

func (obj *bgpPrefixesStateBgpPrefixIpv6UnicastStateIter) Set(index int, newObj BgpPrefixIpv6UnicastState) BgpPrefixesStateBgpPrefixIpv6UnicastStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpPrefixIpv6UnicastStateSlice[index] = newObj
	return obj
}
func (obj *bgpPrefixesStateBgpPrefixIpv6UnicastStateIter) Clear() BgpPrefixesStateBgpPrefixIpv6UnicastStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpPrefixIpv6UnicastState{}
		obj.bgpPrefixIpv6UnicastStateSlice = []BgpPrefixIpv6UnicastState{}
	}
	return obj
}
func (obj *bgpPrefixesStateBgpPrefixIpv6UnicastStateIter) clearHolderSlice() BgpPrefixesStateBgpPrefixIpv6UnicastStateIter {
	if len(obj.bgpPrefixIpv6UnicastStateSlice) > 0 {
		obj.bgpPrefixIpv6UnicastStateSlice = []BgpPrefixIpv6UnicastState{}
	}
	return obj
}
func (obj *bgpPrefixesStateBgpPrefixIpv6UnicastStateIter) appendHolderSlice(item BgpPrefixIpv6UnicastState) BgpPrefixesStateBgpPrefixIpv6UnicastStateIter {
	obj.bgpPrefixIpv6UnicastStateSlice = append(obj.bgpPrefixIpv6UnicastStateSlice, item)
	return obj
}

// description is TBD
// Ipv4MplsUnicastPrefixes returns a []BgpPrefixIpv4MplsUnicastState
func (obj *bgpPrefixesState) Ipv4MplsUnicastPrefixes() BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter {
	if len(obj.obj.Ipv4MplsUnicastPrefixes) == 0 {
		obj.obj.Ipv4MplsUnicastPrefixes = []*otg.BgpPrefixIpv4MplsUnicastState{}
	}
	if obj.ipv4MplsUnicastPrefixesHolder == nil {
		obj.ipv4MplsUnicastPrefixesHolder = newBgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter(&obj.obj.Ipv4MplsUnicastPrefixes).setMsg(obj)
	}
	return obj.ipv4MplsUnicastPrefixesHolder
}

type bgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter struct {
	obj                                *bgpPrefixesState
	bgpPrefixIpv4MplsUnicastStateSlice []BgpPrefixIpv4MplsUnicastState
	fieldPtr                           *[]*otg.BgpPrefixIpv4MplsUnicastState
}

func newBgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter(ptr *[]*otg.BgpPrefixIpv4MplsUnicastState) BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter {
	return &bgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter{fieldPtr: ptr}
}

type BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter interface {
	setMsg(*bgpPrefixesState) BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter
	Items() []BgpPrefixIpv4MplsUnicastState
	Add() BgpPrefixIpv4MplsUnicastState
	Append(items ...BgpPrefixIpv4MplsUnicastState) BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter
	Set(index int, newObj BgpPrefixIpv4MplsUnicastState) BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter
	Clear() BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter
	clearHolderSlice() BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter
	appendHolderSlice(item BgpPrefixIpv4MplsUnicastState) BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter
}

func (obj *bgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter) setMsg(msg *bgpPrefixesState) BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpPrefixIpv4MplsUnicastState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter) Items() []BgpPrefixIpv4MplsUnicastState {
	return obj.bgpPrefixIpv4MplsUnicastStateSlice
}

func (obj *bgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter) Add() BgpPrefixIpv4MplsUnicastState {
	newObj := &otg.BgpPrefixIpv4MplsUnicastState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpPrefixIpv4MplsUnicastState{obj: newObj}
	newLibObj.setDefault()
	obj.bgpPrefixIpv4MplsUnicastStateSlice = append(obj.bgpPrefixIpv4MplsUnicastStateSlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter) Append(items ...BgpPrefixIpv4MplsUnicastState) BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpPrefixIpv4MplsUnicastStateSlice = append(obj.bgpPrefixIpv4MplsUnicastStateSlice, item)
	}
	return obj
}

func (obj *bgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter) Set(index int, newObj BgpPrefixIpv4MplsUnicastState) BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpPrefixIpv4MplsUnicastStateSlice[index] = newObj
	return obj
}
func (obj *bgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter) Clear() BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpPrefixIpv4MplsUnicastState{}
		obj.bgpPrefixIpv4MplsUnicastStateSlice = []BgpPrefixIpv4MplsUnicastState{}
	}
	return obj
}
func (obj *bgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter) clearHolderSlice() BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter {
	if len(obj.bgpPrefixIpv4MplsUnicastStateSlice) > 0 {
		obj.bgpPrefixIpv4MplsUnicastStateSlice = []BgpPrefixIpv4MplsUnicastState{}
	}
	return obj
}
func (obj *bgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter) appendHolderSlice(item BgpPrefixIpv4MplsUnicastState) BgpPrefixesStateBgpPrefixIpv4MplsUnicastStateIter {
	obj.bgpPrefixIpv4MplsUnicastStateSlice = append(obj.bgpPrefixIpv4MplsUnicastStateSlice, item)
	return obj
}

// description is TBD
// Ipv6MplsUnicastPrefixes returns a []BgpPrefixIpv6MplsUnicastState
func (obj *bgpPrefixesState) Ipv6MplsUnicastPrefixes() BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter {
	if len(obj.obj.Ipv6MplsUnicastPrefixes) == 0 {
		obj.obj.Ipv6MplsUnicastPrefixes = []*otg.BgpPrefixIpv6MplsUnicastState{}
	}
	if obj.ipv6MplsUnicastPrefixesHolder == nil {
		obj.ipv6MplsUnicastPrefixesHolder = newBgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter(&obj.obj.Ipv6MplsUnicastPrefixes).setMsg(obj)
	}
	return obj.ipv6MplsUnicastPrefixesHolder
}

type bgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter struct {
	obj                                *bgpPrefixesState
	bgpPrefixIpv6MplsUnicastStateSlice []BgpPrefixIpv6MplsUnicastState
	fieldPtr                           *[]*otg.BgpPrefixIpv6MplsUnicastState
}

func newBgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter(ptr *[]*otg.BgpPrefixIpv6MplsUnicastState) BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter {
	return &bgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter{fieldPtr: ptr}
}

type BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter interface {
	setMsg(*bgpPrefixesState) BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter
	Items() []BgpPrefixIpv6MplsUnicastState
	Add() BgpPrefixIpv6MplsUnicastState
	Append(items ...BgpPrefixIpv6MplsUnicastState) BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter
	Set(index int, newObj BgpPrefixIpv6MplsUnicastState) BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter
	Clear() BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter
	clearHolderSlice() BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter
	appendHolderSlice(item BgpPrefixIpv6MplsUnicastState) BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter
}

func (obj *bgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter) setMsg(msg *bgpPrefixesState) BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpPrefixIpv6MplsUnicastState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter) Items() []BgpPrefixIpv6MplsUnicastState {
	return obj.bgpPrefixIpv6MplsUnicastStateSlice
}

func (obj *bgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter) Add() BgpPrefixIpv6MplsUnicastState {
	newObj := &otg.BgpPrefixIpv6MplsUnicastState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpPrefixIpv6MplsUnicastState{obj: newObj}
	newLibObj.setDefault()
	obj.bgpPrefixIpv6MplsUnicastStateSlice = append(obj.bgpPrefixIpv6MplsUnicastStateSlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter) Append(items ...BgpPrefixIpv6MplsUnicastState) BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpPrefixIpv6MplsUnicastStateSlice = append(obj.bgpPrefixIpv6MplsUnicastStateSlice, item)
	}
	return obj
}

func (obj *bgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter) Set(index int, newObj BgpPrefixIpv6MplsUnicastState) BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpPrefixIpv6MplsUnicastStateSlice[index] = newObj
	return obj
}
func (obj *bgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter) Clear() BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpPrefixIpv6MplsUnicastState{}
		obj.bgpPrefixIpv6MplsUnicastStateSlice = []BgpPrefixIpv6MplsUnicastState{}
	}
	return obj
}
func (obj *bgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter) clearHolderSlice() BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter {
	if len(obj.bgpPrefixIpv6MplsUnicastStateSlice) > 0 {
		obj.bgpPrefixIpv6MplsUnicastStateSlice = []BgpPrefixIpv6MplsUnicastState{}
	}
	return obj
}
func (obj *bgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter) appendHolderSlice(item BgpPrefixIpv6MplsUnicastState) BgpPrefixesStateBgpPrefixIpv6MplsUnicastStateIter {
	obj.bgpPrefixIpv6MplsUnicastStateSlice = append(obj.bgpPrefixIpv6MplsUnicastStateSlice, item)
	return obj
}

func (obj *bgpPrefixesState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4UnicastPrefixes) != 0 {

		if set_default {
			obj.Ipv4UnicastPrefixes().clearHolderSlice()
			for _, item := range obj.obj.Ipv4UnicastPrefixes {
				obj.Ipv4UnicastPrefixes().appendHolderSlice(&bgpPrefixIpv4UnicastState{obj: item})
			}
		}
		for _, item := range obj.Ipv4UnicastPrefixes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv6UnicastPrefixes) != 0 {

		if set_default {
			obj.Ipv6UnicastPrefixes().clearHolderSlice()
			for _, item := range obj.obj.Ipv6UnicastPrefixes {
				obj.Ipv6UnicastPrefixes().appendHolderSlice(&bgpPrefixIpv6UnicastState{obj: item})
			}
		}
		for _, item := range obj.Ipv6UnicastPrefixes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv4MplsUnicastPrefixes) != 0 {

		if set_default {
			obj.Ipv4MplsUnicastPrefixes().clearHolderSlice()
			for _, item := range obj.obj.Ipv4MplsUnicastPrefixes {
				obj.Ipv4MplsUnicastPrefixes().appendHolderSlice(&bgpPrefixIpv4MplsUnicastState{obj: item})
			}
		}
		for _, item := range obj.Ipv4MplsUnicastPrefixes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv6MplsUnicastPrefixes) != 0 {

		if set_default {
			obj.Ipv6MplsUnicastPrefixes().clearHolderSlice()
			for _, item := range obj.obj.Ipv6MplsUnicastPrefixes {
				obj.Ipv6MplsUnicastPrefixes().appendHolderSlice(&bgpPrefixIpv6MplsUnicastState{obj: item})
			}
		}
		for _, item := range obj.Ipv6MplsUnicastPrefixes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpPrefixesState) setDefault() {

}
