package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspExtendedIsReachabilityTlv *****
type isisLspExtendedIsReachabilityTlv struct {
	validation
	obj             *otg.IsisLspExtendedIsReachabilityTlv
	marshaller      marshalIsisLspExtendedIsReachabilityTlv
	unMarshaller    unMarshalIsisLspExtendedIsReachabilityTlv
	neighborsHolder IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter
}

func NewIsisLspExtendedIsReachabilityTlv() IsisLspExtendedIsReachabilityTlv {
	obj := isisLspExtendedIsReachabilityTlv{obj: &otg.IsisLspExtendedIsReachabilityTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspExtendedIsReachabilityTlv) msg() *otg.IsisLspExtendedIsReachabilityTlv {
	return obj.obj
}

func (obj *isisLspExtendedIsReachabilityTlv) setMsg(msg *otg.IsisLspExtendedIsReachabilityTlv) IsisLspExtendedIsReachabilityTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspExtendedIsReachabilityTlv struct {
	obj *isisLspExtendedIsReachabilityTlv
}

type marshalIsisLspExtendedIsReachabilityTlv interface {
	// ToProto marshals IsisLspExtendedIsReachabilityTlv to protobuf object *otg.IsisLspExtendedIsReachabilityTlv
	ToProto() (*otg.IsisLspExtendedIsReachabilityTlv, error)
	// ToPbText marshals IsisLspExtendedIsReachabilityTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspExtendedIsReachabilityTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspExtendedIsReachabilityTlv to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspExtendedIsReachabilityTlv struct {
	obj *isisLspExtendedIsReachabilityTlv
}

type unMarshalIsisLspExtendedIsReachabilityTlv interface {
	// FromProto unmarshals IsisLspExtendedIsReachabilityTlv from protobuf object *otg.IsisLspExtendedIsReachabilityTlv
	FromProto(msg *otg.IsisLspExtendedIsReachabilityTlv) (IsisLspExtendedIsReachabilityTlv, error)
	// FromPbText unmarshals IsisLspExtendedIsReachabilityTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspExtendedIsReachabilityTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspExtendedIsReachabilityTlv from JSON text
	FromJson(value string) error
}

func (obj *isisLspExtendedIsReachabilityTlv) Marshal() marshalIsisLspExtendedIsReachabilityTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspExtendedIsReachabilityTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspExtendedIsReachabilityTlv) Unmarshal() unMarshalIsisLspExtendedIsReachabilityTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspExtendedIsReachabilityTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspExtendedIsReachabilityTlv) ToProto() (*otg.IsisLspExtendedIsReachabilityTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspExtendedIsReachabilityTlv) FromProto(msg *otg.IsisLspExtendedIsReachabilityTlv) (IsisLspExtendedIsReachabilityTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspExtendedIsReachabilityTlv) ToPbText() (string, error) {
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

func (m *unMarshalisisLspExtendedIsReachabilityTlv) FromPbText(value string) error {
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

func (m *marshalisisLspExtendedIsReachabilityTlv) ToYaml() (string, error) {
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

func (m *unMarshalisisLspExtendedIsReachabilityTlv) FromYaml(value string) error {
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

func (m *marshalisisLspExtendedIsReachabilityTlv) ToJson() (string, error) {
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

func (m *unMarshalisisLspExtendedIsReachabilityTlv) FromJson(value string) error {
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

func (obj *isisLspExtendedIsReachabilityTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspExtendedIsReachabilityTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspExtendedIsReachabilityTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspExtendedIsReachabilityTlv) Clone() (IsisLspExtendedIsReachabilityTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspExtendedIsReachabilityTlv()
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

func (obj *isisLspExtendedIsReachabilityTlv) setNil() {
	obj.neighborsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspExtendedIsReachabilityTlv is this is list of ISIS neighbors and attributes in Extended-IS-Reachability TLV (type 22).
type IsisLspExtendedIsReachabilityTlv interface {
	Validation
	// msg marshals IsisLspExtendedIsReachabilityTlv to protobuf object *otg.IsisLspExtendedIsReachabilityTlv
	// and doesn't set defaults
	msg() *otg.IsisLspExtendedIsReachabilityTlv
	// setMsg unmarshals IsisLspExtendedIsReachabilityTlv from protobuf object *otg.IsisLspExtendedIsReachabilityTlv
	// and doesn't set defaults
	setMsg(*otg.IsisLspExtendedIsReachabilityTlv) IsisLspExtendedIsReachabilityTlv
	// provides marshal interface
	Marshal() marshalIsisLspExtendedIsReachabilityTlv
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspExtendedIsReachabilityTlv
	// validate validates IsisLspExtendedIsReachabilityTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspExtendedIsReachabilityTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Neighbors returns IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIterIter, set in IsisLspExtendedIsReachabilityTlv
	Neighbors() IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter
	setNil()
}

// This container describes IS neighbors.
// Neighbors returns a []IsisLspExtendedNeighbor
func (obj *isisLspExtendedIsReachabilityTlv) Neighbors() IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter {
	if len(obj.obj.Neighbors) == 0 {
		obj.obj.Neighbors = []*otg.IsisLspExtendedNeighbor{}
	}
	if obj.neighborsHolder == nil {
		obj.neighborsHolder = newIsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter(&obj.obj.Neighbors).setMsg(obj)
	}
	return obj.neighborsHolder
}

type isisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter struct {
	obj                          *isisLspExtendedIsReachabilityTlv
	isisLspExtendedNeighborSlice []IsisLspExtendedNeighbor
	fieldPtr                     *[]*otg.IsisLspExtendedNeighbor
}

func newIsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter(ptr *[]*otg.IsisLspExtendedNeighbor) IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter {
	return &isisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter{fieldPtr: ptr}
}

type IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter interface {
	setMsg(*isisLspExtendedIsReachabilityTlv) IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter
	Items() []IsisLspExtendedNeighbor
	Add() IsisLspExtendedNeighbor
	Append(items ...IsisLspExtendedNeighbor) IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter
	Set(index int, newObj IsisLspExtendedNeighbor) IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter
	Clear() IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter
	clearHolderSlice() IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter
	appendHolderSlice(item IsisLspExtendedNeighbor) IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter
}

func (obj *isisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter) setMsg(msg *isisLspExtendedIsReachabilityTlv) IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspExtendedNeighbor{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter) Items() []IsisLspExtendedNeighbor {
	return obj.isisLspExtendedNeighborSlice
}

func (obj *isisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter) Add() IsisLspExtendedNeighbor {
	newObj := &otg.IsisLspExtendedNeighbor{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspExtendedNeighbor{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspExtendedNeighborSlice = append(obj.isisLspExtendedNeighborSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter) Append(items ...IsisLspExtendedNeighbor) IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspExtendedNeighborSlice = append(obj.isisLspExtendedNeighborSlice, item)
	}
	return obj
}

func (obj *isisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter) Set(index int, newObj IsisLspExtendedNeighbor) IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspExtendedNeighborSlice[index] = newObj
	return obj
}
func (obj *isisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter) Clear() IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspExtendedNeighbor{}
		obj.isisLspExtendedNeighborSlice = []IsisLspExtendedNeighbor{}
	}
	return obj
}
func (obj *isisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter) clearHolderSlice() IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter {
	if len(obj.isisLspExtendedNeighborSlice) > 0 {
		obj.isisLspExtendedNeighborSlice = []IsisLspExtendedNeighbor{}
	}
	return obj
}
func (obj *isisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter) appendHolderSlice(item IsisLspExtendedNeighbor) IsisLspExtendedIsReachabilityTlvIsisLspExtendedNeighborIter {
	obj.isisLspExtendedNeighborSlice = append(obj.isisLspExtendedNeighborSlice, item)
	return obj
}

func (obj *isisLspExtendedIsReachabilityTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Neighbors) != 0 {

		if set_default {
			obj.Neighbors().clearHolderSlice()
			for _, item := range obj.obj.Neighbors {
				obj.Neighbors().appendHolderSlice(&isisLspExtendedNeighbor{obj: item})
			}
		}
		for _, item := range obj.Neighbors().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisLspExtendedIsReachabilityTlv) setDefault() {

}
