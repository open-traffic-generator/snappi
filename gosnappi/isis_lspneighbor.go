package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspneighbor *****
type isisLspneighbor struct {
	validation
	obj                 *otg.IsisLspneighbor
	marshaller          marshalIsisLspneighbor
	unMarshaller        unMarshalIsisLspneighbor
	adjacencySidsHolder IsisLspneighborIsisLspAdjacencySIDIter
}

func NewIsisLspneighbor() IsisLspneighbor {
	obj := isisLspneighbor{obj: &otg.IsisLspneighbor{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspneighbor) msg() *otg.IsisLspneighbor {
	return obj.obj
}

func (obj *isisLspneighbor) setMsg(msg *otg.IsisLspneighbor) IsisLspneighbor {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspneighbor struct {
	obj *isisLspneighbor
}

type marshalIsisLspneighbor interface {
	// ToProto marshals IsisLspneighbor to protobuf object *otg.IsisLspneighbor
	ToProto() (*otg.IsisLspneighbor, error)
	// ToPbText marshals IsisLspneighbor to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspneighbor to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspneighbor to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspneighbor struct {
	obj *isisLspneighbor
}

type unMarshalIsisLspneighbor interface {
	// FromProto unmarshals IsisLspneighbor from protobuf object *otg.IsisLspneighbor
	FromProto(msg *otg.IsisLspneighbor) (IsisLspneighbor, error)
	// FromPbText unmarshals IsisLspneighbor from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspneighbor from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspneighbor from JSON text
	FromJson(value string) error
}

func (obj *isisLspneighbor) Marshal() marshalIsisLspneighbor {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspneighbor{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspneighbor) Unmarshal() unMarshalIsisLspneighbor {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspneighbor{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspneighbor) ToProto() (*otg.IsisLspneighbor, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspneighbor) FromProto(msg *otg.IsisLspneighbor) (IsisLspneighbor, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspneighbor) ToPbText() (string, error) {
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

func (m *unMarshalisisLspneighbor) FromPbText(value string) error {
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

func (m *marshalisisLspneighbor) ToYaml() (string, error) {
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

func (m *unMarshalisisLspneighbor) FromYaml(value string) error {
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

func (m *marshalisisLspneighbor) ToJson() (string, error) {
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

func (m *unMarshalisisLspneighbor) FromJson(value string) error {
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

func (obj *isisLspneighbor) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspneighbor) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspneighbor) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspneighbor) Clone() (IsisLspneighbor, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspneighbor()
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

func (obj *isisLspneighbor) setNil() {
	obj.adjacencySidsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspneighbor is this contains IS neighbors.
type IsisLspneighbor interface {
	Validation
	// msg marshals IsisLspneighbor to protobuf object *otg.IsisLspneighbor
	// and doesn't set defaults
	msg() *otg.IsisLspneighbor
	// setMsg unmarshals IsisLspneighbor from protobuf object *otg.IsisLspneighbor
	// and doesn't set defaults
	setMsg(*otg.IsisLspneighbor) IsisLspneighbor
	// provides marshal interface
	Marshal() marshalIsisLspneighbor
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspneighbor
	// validate validates IsisLspneighbor
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspneighbor, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SystemId returns string, set in IsisLspneighbor.
	SystemId() string
	// SetSystemId assigns string provided by user to IsisLspneighbor
	SetSystemId(value string) IsisLspneighbor
	// HasSystemId checks if SystemId has been set in IsisLspneighbor
	HasSystemId() bool
	// AdjacencySids returns IsisLspneighborIsisLspAdjacencySIDIterIter, set in IsisLspneighbor
	AdjacencySids() IsisLspneighborIsisLspAdjacencySIDIter
	setNil()
}

// The System ID for this emulated ISIS router, e.g. "640100010000".
// SystemId returns a string
func (obj *isisLspneighbor) SystemId() string {

	return *obj.obj.SystemId

}

// The System ID for this emulated ISIS router, e.g. "640100010000".
// SystemId returns a string
func (obj *isisLspneighbor) HasSystemId() bool {
	return obj.obj.SystemId != nil
}

// The System ID for this emulated ISIS router, e.g. "640100010000".
// SetSystemId sets the string value in the IsisLspneighbor object
func (obj *isisLspneighbor) SetSystemId(value string) IsisLspneighbor {

	obj.obj.SystemId = &value
	return obj
}

// List of segment routing adjacency SIDs.
// AdjacencySids returns a []IsisLspAdjacencySID
func (obj *isisLspneighbor) AdjacencySids() IsisLspneighborIsisLspAdjacencySIDIter {
	if len(obj.obj.AdjacencySids) == 0 {
		obj.obj.AdjacencySids = []*otg.IsisLspAdjacencySID{}
	}
	if obj.adjacencySidsHolder == nil {
		obj.adjacencySidsHolder = newIsisLspneighborIsisLspAdjacencySIDIter(&obj.obj.AdjacencySids).setMsg(obj)
	}
	return obj.adjacencySidsHolder
}

type isisLspneighborIsisLspAdjacencySIDIter struct {
	obj                      *isisLspneighbor
	isisLspAdjacencySIDSlice []IsisLspAdjacencySID
	fieldPtr                 *[]*otg.IsisLspAdjacencySID
}

func newIsisLspneighborIsisLspAdjacencySIDIter(ptr *[]*otg.IsisLspAdjacencySID) IsisLspneighborIsisLspAdjacencySIDIter {
	return &isisLspneighborIsisLspAdjacencySIDIter{fieldPtr: ptr}
}

type IsisLspneighborIsisLspAdjacencySIDIter interface {
	setMsg(*isisLspneighbor) IsisLspneighborIsisLspAdjacencySIDIter
	Items() []IsisLspAdjacencySID
	Add() IsisLspAdjacencySID
	Append(items ...IsisLspAdjacencySID) IsisLspneighborIsisLspAdjacencySIDIter
	Set(index int, newObj IsisLspAdjacencySID) IsisLspneighborIsisLspAdjacencySIDIter
	Clear() IsisLspneighborIsisLspAdjacencySIDIter
	clearHolderSlice() IsisLspneighborIsisLspAdjacencySIDIter
	appendHolderSlice(item IsisLspAdjacencySID) IsisLspneighborIsisLspAdjacencySIDIter
}

func (obj *isisLspneighborIsisLspAdjacencySIDIter) setMsg(msg *isisLspneighbor) IsisLspneighborIsisLspAdjacencySIDIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspAdjacencySID{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspneighborIsisLspAdjacencySIDIter) Items() []IsisLspAdjacencySID {
	return obj.isisLspAdjacencySIDSlice
}

func (obj *isisLspneighborIsisLspAdjacencySIDIter) Add() IsisLspAdjacencySID {
	newObj := &otg.IsisLspAdjacencySID{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspAdjacencySID{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspAdjacencySIDSlice = append(obj.isisLspAdjacencySIDSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspneighborIsisLspAdjacencySIDIter) Append(items ...IsisLspAdjacencySID) IsisLspneighborIsisLspAdjacencySIDIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspAdjacencySIDSlice = append(obj.isisLspAdjacencySIDSlice, item)
	}
	return obj
}

func (obj *isisLspneighborIsisLspAdjacencySIDIter) Set(index int, newObj IsisLspAdjacencySID) IsisLspneighborIsisLspAdjacencySIDIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspAdjacencySIDSlice[index] = newObj
	return obj
}
func (obj *isisLspneighborIsisLspAdjacencySIDIter) Clear() IsisLspneighborIsisLspAdjacencySIDIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspAdjacencySID{}
		obj.isisLspAdjacencySIDSlice = []IsisLspAdjacencySID{}
	}
	return obj
}
func (obj *isisLspneighborIsisLspAdjacencySIDIter) clearHolderSlice() IsisLspneighborIsisLspAdjacencySIDIter {
	if len(obj.isisLspAdjacencySIDSlice) > 0 {
		obj.isisLspAdjacencySIDSlice = []IsisLspAdjacencySID{}
	}
	return obj
}
func (obj *isisLspneighborIsisLspAdjacencySIDIter) appendHolderSlice(item IsisLspAdjacencySID) IsisLspneighborIsisLspAdjacencySIDIter {
	obj.isisLspAdjacencySIDSlice = append(obj.isisLspAdjacencySIDSlice, item)
	return obj
}

func (obj *isisLspneighbor) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SystemId != nil {

		err := obj.validateHex(obj.SystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspneighbor.SystemId"))
		}

	}

	if len(obj.obj.AdjacencySids) != 0 {

		if set_default {
			obj.AdjacencySids().clearHolderSlice()
			for _, item := range obj.obj.AdjacencySids {
				obj.AdjacencySids().appendHolderSlice(&isisLspAdjacencySID{obj: item})
			}
		}
		for _, item := range obj.AdjacencySids().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisLspneighbor) setDefault() {

}
