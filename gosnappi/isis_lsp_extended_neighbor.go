package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspExtendedNeighbor *****
type isisLspExtendedNeighbor struct {
	validation
	obj                 *otg.IsisLspExtendedNeighbor
	marshaller          marshalIsisLspExtendedNeighbor
	unMarshaller        unMarshalIsisLspExtendedNeighbor
	adjacencySidsHolder IsisLspExtendedNeighborIsisLspAdjacencySidIter
	srv6AdjSidsHolder   IsisLspExtendedNeighborIsisLspSRv6AdjSidIter
}

func NewIsisLspExtendedNeighbor() IsisLspExtendedNeighbor {
	obj := isisLspExtendedNeighbor{obj: &otg.IsisLspExtendedNeighbor{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspExtendedNeighbor) msg() *otg.IsisLspExtendedNeighbor {
	return obj.obj
}

func (obj *isisLspExtendedNeighbor) setMsg(msg *otg.IsisLspExtendedNeighbor) IsisLspExtendedNeighbor {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspExtendedNeighbor struct {
	obj *isisLspExtendedNeighbor
}

type marshalIsisLspExtendedNeighbor interface {
	// ToProto marshals IsisLspExtendedNeighbor to protobuf object *otg.IsisLspExtendedNeighbor
	ToProto() (*otg.IsisLspExtendedNeighbor, error)
	// ToPbText marshals IsisLspExtendedNeighbor to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspExtendedNeighbor to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspExtendedNeighbor to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspExtendedNeighbor struct {
	obj *isisLspExtendedNeighbor
}

type unMarshalIsisLspExtendedNeighbor interface {
	// FromProto unmarshals IsisLspExtendedNeighbor from protobuf object *otg.IsisLspExtendedNeighbor
	FromProto(msg *otg.IsisLspExtendedNeighbor) (IsisLspExtendedNeighbor, error)
	// FromPbText unmarshals IsisLspExtendedNeighbor from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspExtendedNeighbor from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspExtendedNeighbor from JSON text
	FromJson(value string) error
}

func (obj *isisLspExtendedNeighbor) Marshal() marshalIsisLspExtendedNeighbor {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspExtendedNeighbor{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspExtendedNeighbor) Unmarshal() unMarshalIsisLspExtendedNeighbor {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspExtendedNeighbor{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspExtendedNeighbor) ToProto() (*otg.IsisLspExtendedNeighbor, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspExtendedNeighbor) FromProto(msg *otg.IsisLspExtendedNeighbor) (IsisLspExtendedNeighbor, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspExtendedNeighbor) ToPbText() (string, error) {
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

func (m *unMarshalisisLspExtendedNeighbor) FromPbText(value string) error {
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

func (m *marshalisisLspExtendedNeighbor) ToYaml() (string, error) {
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

func (m *unMarshalisisLspExtendedNeighbor) FromYaml(value string) error {
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

func (m *marshalisisLspExtendedNeighbor) ToJson() (string, error) {
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

func (m *unMarshalisisLspExtendedNeighbor) FromJson(value string) error {
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

func (obj *isisLspExtendedNeighbor) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspExtendedNeighbor) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspExtendedNeighbor) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspExtendedNeighbor) Clone() (IsisLspExtendedNeighbor, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspExtendedNeighbor()
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

func (obj *isisLspExtendedNeighbor) setNil() {
	obj.adjacencySidsHolder = nil
	obj.srv6AdjSidsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspExtendedNeighbor is this contains IS neighbors.
type IsisLspExtendedNeighbor interface {
	Validation
	// msg marshals IsisLspExtendedNeighbor to protobuf object *otg.IsisLspExtendedNeighbor
	// and doesn't set defaults
	msg() *otg.IsisLspExtendedNeighbor
	// setMsg unmarshals IsisLspExtendedNeighbor from protobuf object *otg.IsisLspExtendedNeighbor
	// and doesn't set defaults
	setMsg(*otg.IsisLspExtendedNeighbor) IsisLspExtendedNeighbor
	// provides marshal interface
	Marshal() marshalIsisLspExtendedNeighbor
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspExtendedNeighbor
	// validate validates IsisLspExtendedNeighbor
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspExtendedNeighbor, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SystemId returns string, set in IsisLspExtendedNeighbor.
	SystemId() string
	// SetSystemId assigns string provided by user to IsisLspExtendedNeighbor
	SetSystemId(value string) IsisLspExtendedNeighbor
	// HasSystemId checks if SystemId has been set in IsisLspExtendedNeighbor
	HasSystemId() bool
	// AdjacencySids returns IsisLspExtendedNeighborIsisLspAdjacencySidIterIter, set in IsisLspExtendedNeighbor
	AdjacencySids() IsisLspExtendedNeighborIsisLspAdjacencySidIter
	// Srv6AdjSids returns IsisLspExtendedNeighborIsisLspSRv6AdjSidIterIter, set in IsisLspExtendedNeighbor
	Srv6AdjSids() IsisLspExtendedNeighborIsisLspSRv6AdjSidIter
	setNil()
}

// The System ID for this emulated ISIS router, e.g. "640100010000".
// SystemId returns a string
func (obj *isisLspExtendedNeighbor) SystemId() string {

	return *obj.obj.SystemId

}

// The System ID for this emulated ISIS router, e.g. "640100010000".
// SystemId returns a string
func (obj *isisLspExtendedNeighbor) HasSystemId() bool {
	return obj.obj.SystemId != nil
}

// The System ID for this emulated ISIS router, e.g. "640100010000".
// SetSystemId sets the string value in the IsisLspExtendedNeighbor object
func (obj *isisLspExtendedNeighbor) SetSystemId(value string) IsisLspExtendedNeighbor {

	obj.obj.SystemId = &value
	return obj
}

// List of SR-MPLS Adjacency SID sub-TLVs (sub-TLV types 31/32,
// RFC 8667) learned from this neighbor.
// AdjacencySids returns a []IsisLspAdjacencySid
func (obj *isisLspExtendedNeighbor) AdjacencySids() IsisLspExtendedNeighborIsisLspAdjacencySidIter {
	if len(obj.obj.AdjacencySids) == 0 {
		obj.obj.AdjacencySids = []*otg.IsisLspAdjacencySid{}
	}
	if obj.adjacencySidsHolder == nil {
		obj.adjacencySidsHolder = newIsisLspExtendedNeighborIsisLspAdjacencySidIter(&obj.obj.AdjacencySids).setMsg(obj)
	}
	return obj.adjacencySidsHolder
}

type isisLspExtendedNeighborIsisLspAdjacencySidIter struct {
	obj                      *isisLspExtendedNeighbor
	isisLspAdjacencySidSlice []IsisLspAdjacencySid
	fieldPtr                 *[]*otg.IsisLspAdjacencySid
}

func newIsisLspExtendedNeighborIsisLspAdjacencySidIter(ptr *[]*otg.IsisLspAdjacencySid) IsisLspExtendedNeighborIsisLspAdjacencySidIter {
	return &isisLspExtendedNeighborIsisLspAdjacencySidIter{fieldPtr: ptr}
}

type IsisLspExtendedNeighborIsisLspAdjacencySidIter interface {
	setMsg(*isisLspExtendedNeighbor) IsisLspExtendedNeighborIsisLspAdjacencySidIter
	Items() []IsisLspAdjacencySid
	Add() IsisLspAdjacencySid
	Append(items ...IsisLspAdjacencySid) IsisLspExtendedNeighborIsisLspAdjacencySidIter
	Set(index int, newObj IsisLspAdjacencySid) IsisLspExtendedNeighborIsisLspAdjacencySidIter
	Clear() IsisLspExtendedNeighborIsisLspAdjacencySidIter
	clearHolderSlice() IsisLspExtendedNeighborIsisLspAdjacencySidIter
	appendHolderSlice(item IsisLspAdjacencySid) IsisLspExtendedNeighborIsisLspAdjacencySidIter
}

func (obj *isisLspExtendedNeighborIsisLspAdjacencySidIter) setMsg(msg *isisLspExtendedNeighbor) IsisLspExtendedNeighborIsisLspAdjacencySidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspAdjacencySid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspExtendedNeighborIsisLspAdjacencySidIter) Items() []IsisLspAdjacencySid {
	return obj.isisLspAdjacencySidSlice
}

func (obj *isisLspExtendedNeighborIsisLspAdjacencySidIter) Add() IsisLspAdjacencySid {
	newObj := &otg.IsisLspAdjacencySid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspAdjacencySid{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspAdjacencySidSlice = append(obj.isisLspAdjacencySidSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspExtendedNeighborIsisLspAdjacencySidIter) Append(items ...IsisLspAdjacencySid) IsisLspExtendedNeighborIsisLspAdjacencySidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspAdjacencySidSlice = append(obj.isisLspAdjacencySidSlice, item)
	}
	return obj
}

func (obj *isisLspExtendedNeighborIsisLspAdjacencySidIter) Set(index int, newObj IsisLspAdjacencySid) IsisLspExtendedNeighborIsisLspAdjacencySidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspAdjacencySidSlice[index] = newObj
	return obj
}
func (obj *isisLspExtendedNeighborIsisLspAdjacencySidIter) Clear() IsisLspExtendedNeighborIsisLspAdjacencySidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspAdjacencySid{}
		obj.isisLspAdjacencySidSlice = []IsisLspAdjacencySid{}
	}
	return obj
}
func (obj *isisLspExtendedNeighborIsisLspAdjacencySidIter) clearHolderSlice() IsisLspExtendedNeighborIsisLspAdjacencySidIter {
	if len(obj.isisLspAdjacencySidSlice) > 0 {
		obj.isisLspAdjacencySidSlice = []IsisLspAdjacencySid{}
	}
	return obj
}
func (obj *isisLspExtendedNeighborIsisLspAdjacencySidIter) appendHolderSlice(item IsisLspAdjacencySid) IsisLspExtendedNeighborIsisLspAdjacencySidIter {
	obj.isisLspAdjacencySidSlice = append(obj.isisLspAdjacencySidSlice, item)
	return obj
}

// List of SRv6 Adjacency SID Sub-TLVs (End.X SID sub-TLV type 43
// for P2P or LAN End.X SID sub-TLV type 44 for broadcast) learned
// from this neighbor. Each entry carries a 128-bit SRv6 SID bound
// to a specific outgoing adjacency.
// Reference: RFC 9352 Sections 8.1-8.2.
// Srv6AdjSids returns a []IsisLspSRv6AdjSid
func (obj *isisLspExtendedNeighbor) Srv6AdjSids() IsisLspExtendedNeighborIsisLspSRv6AdjSidIter {
	if len(obj.obj.Srv6AdjSids) == 0 {
		obj.obj.Srv6AdjSids = []*otg.IsisLspSRv6AdjSid{}
	}
	if obj.srv6AdjSidsHolder == nil {
		obj.srv6AdjSidsHolder = newIsisLspExtendedNeighborIsisLspSRv6AdjSidIter(&obj.obj.Srv6AdjSids).setMsg(obj)
	}
	return obj.srv6AdjSidsHolder
}

type isisLspExtendedNeighborIsisLspSRv6AdjSidIter struct {
	obj                    *isisLspExtendedNeighbor
	isisLspSRv6AdjSidSlice []IsisLspSRv6AdjSid
	fieldPtr               *[]*otg.IsisLspSRv6AdjSid
}

func newIsisLspExtendedNeighborIsisLspSRv6AdjSidIter(ptr *[]*otg.IsisLspSRv6AdjSid) IsisLspExtendedNeighborIsisLspSRv6AdjSidIter {
	return &isisLspExtendedNeighborIsisLspSRv6AdjSidIter{fieldPtr: ptr}
}

type IsisLspExtendedNeighborIsisLspSRv6AdjSidIter interface {
	setMsg(*isisLspExtendedNeighbor) IsisLspExtendedNeighborIsisLspSRv6AdjSidIter
	Items() []IsisLspSRv6AdjSid
	Add() IsisLspSRv6AdjSid
	Append(items ...IsisLspSRv6AdjSid) IsisLspExtendedNeighborIsisLspSRv6AdjSidIter
	Set(index int, newObj IsisLspSRv6AdjSid) IsisLspExtendedNeighborIsisLspSRv6AdjSidIter
	Clear() IsisLspExtendedNeighborIsisLspSRv6AdjSidIter
	clearHolderSlice() IsisLspExtendedNeighborIsisLspSRv6AdjSidIter
	appendHolderSlice(item IsisLspSRv6AdjSid) IsisLspExtendedNeighborIsisLspSRv6AdjSidIter
}

func (obj *isisLspExtendedNeighborIsisLspSRv6AdjSidIter) setMsg(msg *isisLspExtendedNeighbor) IsisLspExtendedNeighborIsisLspSRv6AdjSidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspSRv6AdjSid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspExtendedNeighborIsisLspSRv6AdjSidIter) Items() []IsisLspSRv6AdjSid {
	return obj.isisLspSRv6AdjSidSlice
}

func (obj *isisLspExtendedNeighborIsisLspSRv6AdjSidIter) Add() IsisLspSRv6AdjSid {
	newObj := &otg.IsisLspSRv6AdjSid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspSRv6AdjSid{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspSRv6AdjSidSlice = append(obj.isisLspSRv6AdjSidSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspExtendedNeighborIsisLspSRv6AdjSidIter) Append(items ...IsisLspSRv6AdjSid) IsisLspExtendedNeighborIsisLspSRv6AdjSidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspSRv6AdjSidSlice = append(obj.isisLspSRv6AdjSidSlice, item)
	}
	return obj
}

func (obj *isisLspExtendedNeighborIsisLspSRv6AdjSidIter) Set(index int, newObj IsisLspSRv6AdjSid) IsisLspExtendedNeighborIsisLspSRv6AdjSidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspSRv6AdjSidSlice[index] = newObj
	return obj
}
func (obj *isisLspExtendedNeighborIsisLspSRv6AdjSidIter) Clear() IsisLspExtendedNeighborIsisLspSRv6AdjSidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspSRv6AdjSid{}
		obj.isisLspSRv6AdjSidSlice = []IsisLspSRv6AdjSid{}
	}
	return obj
}
func (obj *isisLspExtendedNeighborIsisLspSRv6AdjSidIter) clearHolderSlice() IsisLspExtendedNeighborIsisLspSRv6AdjSidIter {
	if len(obj.isisLspSRv6AdjSidSlice) > 0 {
		obj.isisLspSRv6AdjSidSlice = []IsisLspSRv6AdjSid{}
	}
	return obj
}
func (obj *isisLspExtendedNeighborIsisLspSRv6AdjSidIter) appendHolderSlice(item IsisLspSRv6AdjSid) IsisLspExtendedNeighborIsisLspSRv6AdjSidIter {
	obj.isisLspSRv6AdjSidSlice = append(obj.isisLspSRv6AdjSidSlice, item)
	return obj
}

func (obj *isisLspExtendedNeighbor) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SystemId != nil {

		err := obj.validateHex(obj.SystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspExtendedNeighbor.SystemId"))
		}

	}

	if len(obj.obj.AdjacencySids) != 0 {

		if set_default {
			obj.AdjacencySids().clearHolderSlice()
			for _, item := range obj.obj.AdjacencySids {
				obj.AdjacencySids().appendHolderSlice(&isisLspAdjacencySid{obj: item})
			}
		}
		for _, item := range obj.AdjacencySids().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Srv6AdjSids) != 0 {

		if set_default {
			obj.Srv6AdjSids().clearHolderSlice()
			for _, item := range obj.obj.Srv6AdjSids {
				obj.Srv6AdjSids().appendHolderSlice(&isisLspSRv6AdjSid{obj: item})
			}
		}
		for _, item := range obj.Srv6AdjSids().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisLspExtendedNeighbor) setDefault() {

}
