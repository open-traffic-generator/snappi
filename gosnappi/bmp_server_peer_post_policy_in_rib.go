package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BmpServerPeerPostPolicyInRib *****
type bmpServerPeerPostPolicyInRib struct {
	validation
	obj                       *otg.BmpServerPeerPostPolicyInRib
	marshaller                marshalBmpServerPeerPostPolicyInRib
	unMarshaller              unMarshalBmpServerPeerPostPolicyInRib
	ipv4UnicastPrefixesHolder BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter
	ipv6UnicastPrefixesHolder BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter
}

func NewBmpServerPeerPostPolicyInRib() BmpServerPeerPostPolicyInRib {
	obj := bmpServerPeerPostPolicyInRib{obj: &otg.BmpServerPeerPostPolicyInRib{}}
	obj.setDefault()
	return &obj
}

func (obj *bmpServerPeerPostPolicyInRib) msg() *otg.BmpServerPeerPostPolicyInRib {
	return obj.obj
}

func (obj *bmpServerPeerPostPolicyInRib) setMsg(msg *otg.BmpServerPeerPostPolicyInRib) BmpServerPeerPostPolicyInRib {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbmpServerPeerPostPolicyInRib struct {
	obj *bmpServerPeerPostPolicyInRib
}

type marshalBmpServerPeerPostPolicyInRib interface {
	// ToProto marshals BmpServerPeerPostPolicyInRib to protobuf object *otg.BmpServerPeerPostPolicyInRib
	ToProto() (*otg.BmpServerPeerPostPolicyInRib, error)
	// ToPbText marshals BmpServerPeerPostPolicyInRib to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BmpServerPeerPostPolicyInRib to YAML text
	ToYaml() (string, error)
	// ToJson marshals BmpServerPeerPostPolicyInRib to JSON text
	ToJson() (string, error)
}

type unMarshalbmpServerPeerPostPolicyInRib struct {
	obj *bmpServerPeerPostPolicyInRib
}

type unMarshalBmpServerPeerPostPolicyInRib interface {
	// FromProto unmarshals BmpServerPeerPostPolicyInRib from protobuf object *otg.BmpServerPeerPostPolicyInRib
	FromProto(msg *otg.BmpServerPeerPostPolicyInRib) (BmpServerPeerPostPolicyInRib, error)
	// FromPbText unmarshals BmpServerPeerPostPolicyInRib from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BmpServerPeerPostPolicyInRib from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BmpServerPeerPostPolicyInRib from JSON text
	FromJson(value string) error
}

func (obj *bmpServerPeerPostPolicyInRib) Marshal() marshalBmpServerPeerPostPolicyInRib {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbmpServerPeerPostPolicyInRib{obj: obj}
	}
	return obj.marshaller
}

func (obj *bmpServerPeerPostPolicyInRib) Unmarshal() unMarshalBmpServerPeerPostPolicyInRib {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbmpServerPeerPostPolicyInRib{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbmpServerPeerPostPolicyInRib) ToProto() (*otg.BmpServerPeerPostPolicyInRib, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbmpServerPeerPostPolicyInRib) FromProto(msg *otg.BmpServerPeerPostPolicyInRib) (BmpServerPeerPostPolicyInRib, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbmpServerPeerPostPolicyInRib) ToPbText() (string, error) {
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

func (m *unMarshalbmpServerPeerPostPolicyInRib) FromPbText(value string) error {
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

func (m *marshalbmpServerPeerPostPolicyInRib) ToYaml() (string, error) {
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

func (m *unMarshalbmpServerPeerPostPolicyInRib) FromYaml(value string) error {
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

func (m *marshalbmpServerPeerPostPolicyInRib) ToJson() (string, error) {
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

func (m *unMarshalbmpServerPeerPostPolicyInRib) FromJson(value string) error {
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

func (obj *bmpServerPeerPostPolicyInRib) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bmpServerPeerPostPolicyInRib) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bmpServerPeerPostPolicyInRib) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bmpServerPeerPostPolicyInRib) Clone() (BmpServerPeerPostPolicyInRib, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBmpServerPeerPostPolicyInRib()
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

func (obj *bmpServerPeerPostPolicyInRib) setNil() {
	obj.ipv4UnicastPrefixesHolder = nil
	obj.ipv6UnicastPrefixesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BmpServerPeerPostPolicyInRib is the set of routes received from this peer as reported using BMP Route Monitoring messages by the BMP client  which are part of the Adj-RIB-In database *after* inbound policies were applied. This is determined by checking that the L flag is set in the Per-Peer Header in the received Route Monitoring message. It should also be ensured that the O flag as defined in RFC8671 is not set in the flags. (indicates Adj-RIB-Out)  Note that routes which have been advertised initially and currently in withdrawn state can be included in the returned set of routes by setting the status of that specific route to 'withdrawn'. Routes which have been received from this peer and part of current Post-Policy In-Rib DB should have the status set to 'advertised'. Note that if prefix_storage.ipv4_unicast.discard or/and prefix_storage.ipv6_unicast.discard is configured or exceptions in the prefix_storage are specified such that all incoming routes are filtered, the corresponding prefix database list will be empty. .
type BmpServerPeerPostPolicyInRib interface {
	Validation
	// msg marshals BmpServerPeerPostPolicyInRib to protobuf object *otg.BmpServerPeerPostPolicyInRib
	// and doesn't set defaults
	msg() *otg.BmpServerPeerPostPolicyInRib
	// setMsg unmarshals BmpServerPeerPostPolicyInRib from protobuf object *otg.BmpServerPeerPostPolicyInRib
	// and doesn't set defaults
	setMsg(*otg.BmpServerPeerPostPolicyInRib) BmpServerPeerPostPolicyInRib
	// provides marshal interface
	Marshal() marshalBmpServerPeerPostPolicyInRib
	// provides unmarshal interface
	Unmarshal() unMarshalBmpServerPeerPostPolicyInRib
	// validate validates BmpServerPeerPostPolicyInRib
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BmpServerPeerPostPolicyInRib, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4UnicastPrefixes returns BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIterIter, set in BmpServerPeerPostPolicyInRib
	Ipv4UnicastPrefixes() BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter
	// Ipv6UnicastPrefixes returns BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIterIter, set in BmpServerPeerPostPolicyInRib
	Ipv6UnicastPrefixes() BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter
	setNil()
}

// description is TBD
// Ipv4UnicastPrefixes returns a []BmpPrefixIpv4UnicastState
func (obj *bmpServerPeerPostPolicyInRib) Ipv4UnicastPrefixes() BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter {
	if len(obj.obj.Ipv4UnicastPrefixes) == 0 {
		obj.obj.Ipv4UnicastPrefixes = []*otg.BmpPrefixIpv4UnicastState{}
	}
	if obj.ipv4UnicastPrefixesHolder == nil {
		obj.ipv4UnicastPrefixesHolder = newBmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter(&obj.obj.Ipv4UnicastPrefixes).setMsg(obj)
	}
	return obj.ipv4UnicastPrefixesHolder
}

type bmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter struct {
	obj                            *bmpServerPeerPostPolicyInRib
	bmpPrefixIpv4UnicastStateSlice []BmpPrefixIpv4UnicastState
	fieldPtr                       *[]*otg.BmpPrefixIpv4UnicastState
}

func newBmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter(ptr *[]*otg.BmpPrefixIpv4UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter {
	return &bmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter{fieldPtr: ptr}
}

type BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter interface {
	setMsg(*bmpServerPeerPostPolicyInRib) BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter
	Items() []BmpPrefixIpv4UnicastState
	Add() BmpPrefixIpv4UnicastState
	Append(items ...BmpPrefixIpv4UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter
	Set(index int, newObj BmpPrefixIpv4UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter
	Clear() BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter
	clearHolderSlice() BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter
	appendHolderSlice(item BmpPrefixIpv4UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter
}

func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter) setMsg(msg *bmpServerPeerPostPolicyInRib) BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bmpPrefixIpv4UnicastState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter) Items() []BmpPrefixIpv4UnicastState {
	return obj.bmpPrefixIpv4UnicastStateSlice
}

func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter) Add() BmpPrefixIpv4UnicastState {
	newObj := &otg.BmpPrefixIpv4UnicastState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bmpPrefixIpv4UnicastState{obj: newObj}
	newLibObj.setDefault()
	obj.bmpPrefixIpv4UnicastStateSlice = append(obj.bmpPrefixIpv4UnicastStateSlice, newLibObj)
	return newLibObj
}

func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter) Append(items ...BmpPrefixIpv4UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bmpPrefixIpv4UnicastStateSlice = append(obj.bmpPrefixIpv4UnicastStateSlice, item)
	}
	return obj
}

func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter) Set(index int, newObj BmpPrefixIpv4UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bmpPrefixIpv4UnicastStateSlice[index] = newObj
	return obj
}
func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter) Clear() BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BmpPrefixIpv4UnicastState{}
		obj.bmpPrefixIpv4UnicastStateSlice = []BmpPrefixIpv4UnicastState{}
	}
	return obj
}
func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter) clearHolderSlice() BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter {
	if len(obj.bmpPrefixIpv4UnicastStateSlice) > 0 {
		obj.bmpPrefixIpv4UnicastStateSlice = []BmpPrefixIpv4UnicastState{}
	}
	return obj
}
func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter) appendHolderSlice(item BmpPrefixIpv4UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv4UnicastStateIter {
	obj.bmpPrefixIpv4UnicastStateSlice = append(obj.bmpPrefixIpv4UnicastStateSlice, item)
	return obj
}

// description is TBD
// Ipv6UnicastPrefixes returns a []BmpPrefixIpv6UnicastState
func (obj *bmpServerPeerPostPolicyInRib) Ipv6UnicastPrefixes() BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter {
	if len(obj.obj.Ipv6UnicastPrefixes) == 0 {
		obj.obj.Ipv6UnicastPrefixes = []*otg.BmpPrefixIpv6UnicastState{}
	}
	if obj.ipv6UnicastPrefixesHolder == nil {
		obj.ipv6UnicastPrefixesHolder = newBmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter(&obj.obj.Ipv6UnicastPrefixes).setMsg(obj)
	}
	return obj.ipv6UnicastPrefixesHolder
}

type bmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter struct {
	obj                            *bmpServerPeerPostPolicyInRib
	bmpPrefixIpv6UnicastStateSlice []BmpPrefixIpv6UnicastState
	fieldPtr                       *[]*otg.BmpPrefixIpv6UnicastState
}

func newBmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter(ptr *[]*otg.BmpPrefixIpv6UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter {
	return &bmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter{fieldPtr: ptr}
}

type BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter interface {
	setMsg(*bmpServerPeerPostPolicyInRib) BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter
	Items() []BmpPrefixIpv6UnicastState
	Add() BmpPrefixIpv6UnicastState
	Append(items ...BmpPrefixIpv6UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter
	Set(index int, newObj BmpPrefixIpv6UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter
	Clear() BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter
	clearHolderSlice() BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter
	appendHolderSlice(item BmpPrefixIpv6UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter
}

func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter) setMsg(msg *bmpServerPeerPostPolicyInRib) BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bmpPrefixIpv6UnicastState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter) Items() []BmpPrefixIpv6UnicastState {
	return obj.bmpPrefixIpv6UnicastStateSlice
}

func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter) Add() BmpPrefixIpv6UnicastState {
	newObj := &otg.BmpPrefixIpv6UnicastState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bmpPrefixIpv6UnicastState{obj: newObj}
	newLibObj.setDefault()
	obj.bmpPrefixIpv6UnicastStateSlice = append(obj.bmpPrefixIpv6UnicastStateSlice, newLibObj)
	return newLibObj
}

func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter) Append(items ...BmpPrefixIpv6UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bmpPrefixIpv6UnicastStateSlice = append(obj.bmpPrefixIpv6UnicastStateSlice, item)
	}
	return obj
}

func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter) Set(index int, newObj BmpPrefixIpv6UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bmpPrefixIpv6UnicastStateSlice[index] = newObj
	return obj
}
func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter) Clear() BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BmpPrefixIpv6UnicastState{}
		obj.bmpPrefixIpv6UnicastStateSlice = []BmpPrefixIpv6UnicastState{}
	}
	return obj
}
func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter) clearHolderSlice() BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter {
	if len(obj.bmpPrefixIpv6UnicastStateSlice) > 0 {
		obj.bmpPrefixIpv6UnicastStateSlice = []BmpPrefixIpv6UnicastState{}
	}
	return obj
}
func (obj *bmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter) appendHolderSlice(item BmpPrefixIpv6UnicastState) BmpServerPeerPostPolicyInRibBmpPrefixIpv6UnicastStateIter {
	obj.bmpPrefixIpv6UnicastStateSlice = append(obj.bmpPrefixIpv6UnicastStateSlice, item)
	return obj
}

func (obj *bmpServerPeerPostPolicyInRib) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4UnicastPrefixes) != 0 {

		if set_default {
			obj.Ipv4UnicastPrefixes().clearHolderSlice()
			for _, item := range obj.obj.Ipv4UnicastPrefixes {
				obj.Ipv4UnicastPrefixes().appendHolderSlice(&bmpPrefixIpv4UnicastState{obj: item})
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
				obj.Ipv6UnicastPrefixes().appendHolderSlice(&bmpPrefixIpv6UnicastState{obj: item})
			}
		}
		for _, item := range obj.Ipv6UnicastPrefixes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bmpServerPeerPostPolicyInRib) setDefault() {

}
