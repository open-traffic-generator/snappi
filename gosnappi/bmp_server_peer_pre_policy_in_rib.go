package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BmpServerPeerPrePolicyInRib *****
type bmpServerPeerPrePolicyInRib struct {
	validation
	obj                       *otg.BmpServerPeerPrePolicyInRib
	marshaller                marshalBmpServerPeerPrePolicyInRib
	unMarshaller              unMarshalBmpServerPeerPrePolicyInRib
	ipv4UnicastPrefixesHolder BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter
	ipv6UnicastPrefixesHolder BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter
}

func NewBmpServerPeerPrePolicyInRib() BmpServerPeerPrePolicyInRib {
	obj := bmpServerPeerPrePolicyInRib{obj: &otg.BmpServerPeerPrePolicyInRib{}}
	obj.setDefault()
	return &obj
}

func (obj *bmpServerPeerPrePolicyInRib) msg() *otg.BmpServerPeerPrePolicyInRib {
	return obj.obj
}

func (obj *bmpServerPeerPrePolicyInRib) setMsg(msg *otg.BmpServerPeerPrePolicyInRib) BmpServerPeerPrePolicyInRib {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbmpServerPeerPrePolicyInRib struct {
	obj *bmpServerPeerPrePolicyInRib
}

type marshalBmpServerPeerPrePolicyInRib interface {
	// ToProto marshals BmpServerPeerPrePolicyInRib to protobuf object *otg.BmpServerPeerPrePolicyInRib
	ToProto() (*otg.BmpServerPeerPrePolicyInRib, error)
	// ToPbText marshals BmpServerPeerPrePolicyInRib to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BmpServerPeerPrePolicyInRib to YAML text
	ToYaml() (string, error)
	// ToJson marshals BmpServerPeerPrePolicyInRib to JSON text
	ToJson() (string, error)
}

type unMarshalbmpServerPeerPrePolicyInRib struct {
	obj *bmpServerPeerPrePolicyInRib
}

type unMarshalBmpServerPeerPrePolicyInRib interface {
	// FromProto unmarshals BmpServerPeerPrePolicyInRib from protobuf object *otg.BmpServerPeerPrePolicyInRib
	FromProto(msg *otg.BmpServerPeerPrePolicyInRib) (BmpServerPeerPrePolicyInRib, error)
	// FromPbText unmarshals BmpServerPeerPrePolicyInRib from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BmpServerPeerPrePolicyInRib from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BmpServerPeerPrePolicyInRib from JSON text
	FromJson(value string) error
}

func (obj *bmpServerPeerPrePolicyInRib) Marshal() marshalBmpServerPeerPrePolicyInRib {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbmpServerPeerPrePolicyInRib{obj: obj}
	}
	return obj.marshaller
}

func (obj *bmpServerPeerPrePolicyInRib) Unmarshal() unMarshalBmpServerPeerPrePolicyInRib {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbmpServerPeerPrePolicyInRib{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbmpServerPeerPrePolicyInRib) ToProto() (*otg.BmpServerPeerPrePolicyInRib, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbmpServerPeerPrePolicyInRib) FromProto(msg *otg.BmpServerPeerPrePolicyInRib) (BmpServerPeerPrePolicyInRib, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbmpServerPeerPrePolicyInRib) ToPbText() (string, error) {
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

func (m *unMarshalbmpServerPeerPrePolicyInRib) FromPbText(value string) error {
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

func (m *marshalbmpServerPeerPrePolicyInRib) ToYaml() (string, error) {
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

func (m *unMarshalbmpServerPeerPrePolicyInRib) FromYaml(value string) error {
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

func (m *marshalbmpServerPeerPrePolicyInRib) ToJson() (string, error) {
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

func (m *unMarshalbmpServerPeerPrePolicyInRib) FromJson(value string) error {
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

func (obj *bmpServerPeerPrePolicyInRib) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bmpServerPeerPrePolicyInRib) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bmpServerPeerPrePolicyInRib) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bmpServerPeerPrePolicyInRib) Clone() (BmpServerPeerPrePolicyInRib, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBmpServerPeerPrePolicyInRib()
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

func (obj *bmpServerPeerPrePolicyInRib) setNil() {
	obj.ipv4UnicastPrefixesHolder = nil
	obj.ipv6UnicastPrefixesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BmpServerPeerPrePolicyInRib is the set of routes received from this peer as reported using BMP Route Monitoring messages by the BMP client  which are part of the Adj-RIB-In database *before* inbound policies were applied. This is determined by checking that the L flag is not set in the Per-Peer Header in the received Route Monitoring message. It should also be ensured that the O flag as defined in RFC8671 is not set in the flags. (indicates Adj-RIB-Out)  Note that routes which have been advertised initially and currently in withdrawn state can be included in the returned set of routes by setting the status of that specific route to 'withdrawn'. Routes which have been received from this peer and part of current Pre-Policy In-Rib DB should have the status set to 'advertised'. Note that if prefix_storage.ipv4_unicast.discard or/and prefix_storage.ipv6_unicast.discard is configured or exceptions in the prefix_storage are specified such that all incoming routes are filtered, the corresponding prefix database list will be empty.
type BmpServerPeerPrePolicyInRib interface {
	Validation
	// msg marshals BmpServerPeerPrePolicyInRib to protobuf object *otg.BmpServerPeerPrePolicyInRib
	// and doesn't set defaults
	msg() *otg.BmpServerPeerPrePolicyInRib
	// setMsg unmarshals BmpServerPeerPrePolicyInRib from protobuf object *otg.BmpServerPeerPrePolicyInRib
	// and doesn't set defaults
	setMsg(*otg.BmpServerPeerPrePolicyInRib) BmpServerPeerPrePolicyInRib
	// provides marshal interface
	Marshal() marshalBmpServerPeerPrePolicyInRib
	// provides unmarshal interface
	Unmarshal() unMarshalBmpServerPeerPrePolicyInRib
	// validate validates BmpServerPeerPrePolicyInRib
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BmpServerPeerPrePolicyInRib, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4UnicastPrefixes returns BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIterIter, set in BmpServerPeerPrePolicyInRib
	Ipv4UnicastPrefixes() BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter
	// Ipv6UnicastPrefixes returns BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIterIter, set in BmpServerPeerPrePolicyInRib
	Ipv6UnicastPrefixes() BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter
	setNil()
}

// description is TBD
// Ipv4UnicastPrefixes returns a []BmpPrefixIpv4UnicastState
func (obj *bmpServerPeerPrePolicyInRib) Ipv4UnicastPrefixes() BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter {
	if len(obj.obj.Ipv4UnicastPrefixes) == 0 {
		obj.obj.Ipv4UnicastPrefixes = []*otg.BmpPrefixIpv4UnicastState{}
	}
	if obj.ipv4UnicastPrefixesHolder == nil {
		obj.ipv4UnicastPrefixesHolder = newBmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter(&obj.obj.Ipv4UnicastPrefixes).setMsg(obj)
	}
	return obj.ipv4UnicastPrefixesHolder
}

type bmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter struct {
	obj                            *bmpServerPeerPrePolicyInRib
	bmpPrefixIpv4UnicastStateSlice []BmpPrefixIpv4UnicastState
	fieldPtr                       *[]*otg.BmpPrefixIpv4UnicastState
}

func newBmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter(ptr *[]*otg.BmpPrefixIpv4UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter {
	return &bmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter{fieldPtr: ptr}
}

type BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter interface {
	setMsg(*bmpServerPeerPrePolicyInRib) BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter
	Items() []BmpPrefixIpv4UnicastState
	Add() BmpPrefixIpv4UnicastState
	Append(items ...BmpPrefixIpv4UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter
	Set(index int, newObj BmpPrefixIpv4UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter
	Clear() BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter
	clearHolderSlice() BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter
	appendHolderSlice(item BmpPrefixIpv4UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter
}

func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter) setMsg(msg *bmpServerPeerPrePolicyInRib) BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bmpPrefixIpv4UnicastState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter) Items() []BmpPrefixIpv4UnicastState {
	return obj.bmpPrefixIpv4UnicastStateSlice
}

func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter) Add() BmpPrefixIpv4UnicastState {
	newObj := &otg.BmpPrefixIpv4UnicastState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bmpPrefixIpv4UnicastState{obj: newObj}
	newLibObj.setDefault()
	obj.bmpPrefixIpv4UnicastStateSlice = append(obj.bmpPrefixIpv4UnicastStateSlice, newLibObj)
	return newLibObj
}

func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter) Append(items ...BmpPrefixIpv4UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bmpPrefixIpv4UnicastStateSlice = append(obj.bmpPrefixIpv4UnicastStateSlice, item)
	}
	return obj
}

func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter) Set(index int, newObj BmpPrefixIpv4UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bmpPrefixIpv4UnicastStateSlice[index] = newObj
	return obj
}
func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter) Clear() BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BmpPrefixIpv4UnicastState{}
		obj.bmpPrefixIpv4UnicastStateSlice = []BmpPrefixIpv4UnicastState{}
	}
	return obj
}
func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter) clearHolderSlice() BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter {
	if len(obj.bmpPrefixIpv4UnicastStateSlice) > 0 {
		obj.bmpPrefixIpv4UnicastStateSlice = []BmpPrefixIpv4UnicastState{}
	}
	return obj
}
func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter) appendHolderSlice(item BmpPrefixIpv4UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv4UnicastStateIter {
	obj.bmpPrefixIpv4UnicastStateSlice = append(obj.bmpPrefixIpv4UnicastStateSlice, item)
	return obj
}

// description is TBD
// Ipv6UnicastPrefixes returns a []BmpPrefixIpv6UnicastState
func (obj *bmpServerPeerPrePolicyInRib) Ipv6UnicastPrefixes() BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter {
	if len(obj.obj.Ipv6UnicastPrefixes) == 0 {
		obj.obj.Ipv6UnicastPrefixes = []*otg.BmpPrefixIpv6UnicastState{}
	}
	if obj.ipv6UnicastPrefixesHolder == nil {
		obj.ipv6UnicastPrefixesHolder = newBmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter(&obj.obj.Ipv6UnicastPrefixes).setMsg(obj)
	}
	return obj.ipv6UnicastPrefixesHolder
}

type bmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter struct {
	obj                            *bmpServerPeerPrePolicyInRib
	bmpPrefixIpv6UnicastStateSlice []BmpPrefixIpv6UnicastState
	fieldPtr                       *[]*otg.BmpPrefixIpv6UnicastState
}

func newBmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter(ptr *[]*otg.BmpPrefixIpv6UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter {
	return &bmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter{fieldPtr: ptr}
}

type BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter interface {
	setMsg(*bmpServerPeerPrePolicyInRib) BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter
	Items() []BmpPrefixIpv6UnicastState
	Add() BmpPrefixIpv6UnicastState
	Append(items ...BmpPrefixIpv6UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter
	Set(index int, newObj BmpPrefixIpv6UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter
	Clear() BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter
	clearHolderSlice() BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter
	appendHolderSlice(item BmpPrefixIpv6UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter
}

func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter) setMsg(msg *bmpServerPeerPrePolicyInRib) BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bmpPrefixIpv6UnicastState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter) Items() []BmpPrefixIpv6UnicastState {
	return obj.bmpPrefixIpv6UnicastStateSlice
}

func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter) Add() BmpPrefixIpv6UnicastState {
	newObj := &otg.BmpPrefixIpv6UnicastState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bmpPrefixIpv6UnicastState{obj: newObj}
	newLibObj.setDefault()
	obj.bmpPrefixIpv6UnicastStateSlice = append(obj.bmpPrefixIpv6UnicastStateSlice, newLibObj)
	return newLibObj
}

func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter) Append(items ...BmpPrefixIpv6UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bmpPrefixIpv6UnicastStateSlice = append(obj.bmpPrefixIpv6UnicastStateSlice, item)
	}
	return obj
}

func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter) Set(index int, newObj BmpPrefixIpv6UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bmpPrefixIpv6UnicastStateSlice[index] = newObj
	return obj
}
func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter) Clear() BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BmpPrefixIpv6UnicastState{}
		obj.bmpPrefixIpv6UnicastStateSlice = []BmpPrefixIpv6UnicastState{}
	}
	return obj
}
func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter) clearHolderSlice() BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter {
	if len(obj.bmpPrefixIpv6UnicastStateSlice) > 0 {
		obj.bmpPrefixIpv6UnicastStateSlice = []BmpPrefixIpv6UnicastState{}
	}
	return obj
}
func (obj *bmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter) appendHolderSlice(item BmpPrefixIpv6UnicastState) BmpServerPeerPrePolicyInRibBmpPrefixIpv6UnicastStateIter {
	obj.bmpPrefixIpv6UnicastStateSlice = append(obj.bmpPrefixIpv6UnicastStateSlice, item)
	return obj
}

func (obj *bmpServerPeerPrePolicyInRib) validateObj(vObj *validation, set_default bool) {
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

func (obj *bmpServerPeerPrePolicyInRib) setDefault() {

}
