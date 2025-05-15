package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspIsReachabilityTlv *****
type isisLspIsReachabilityTlv struct {
	validation
	obj             *otg.IsisLspIsReachabilityTlv
	marshaller      marshalIsisLspIsReachabilityTlv
	unMarshaller    unMarshalIsisLspIsReachabilityTlv
	neighborsHolder IsisLspIsReachabilityTlvIsisLspneighborIter
}

func NewIsisLspIsReachabilityTlv() IsisLspIsReachabilityTlv {
	obj := isisLspIsReachabilityTlv{obj: &otg.IsisLspIsReachabilityTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspIsReachabilityTlv) msg() *otg.IsisLspIsReachabilityTlv {
	return obj.obj
}

func (obj *isisLspIsReachabilityTlv) setMsg(msg *otg.IsisLspIsReachabilityTlv) IsisLspIsReachabilityTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspIsReachabilityTlv struct {
	obj *isisLspIsReachabilityTlv
}

type marshalIsisLspIsReachabilityTlv interface {
	// ToProto marshals IsisLspIsReachabilityTlv to protobuf object *otg.IsisLspIsReachabilityTlv
	ToProto() (*otg.IsisLspIsReachabilityTlv, error)
	// ToPbText marshals IsisLspIsReachabilityTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspIsReachabilityTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspIsReachabilityTlv to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspIsReachabilityTlv struct {
	obj *isisLspIsReachabilityTlv
}

type unMarshalIsisLspIsReachabilityTlv interface {
	// FromProto unmarshals IsisLspIsReachabilityTlv from protobuf object *otg.IsisLspIsReachabilityTlv
	FromProto(msg *otg.IsisLspIsReachabilityTlv) (IsisLspIsReachabilityTlv, error)
	// FromPbText unmarshals IsisLspIsReachabilityTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspIsReachabilityTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspIsReachabilityTlv from JSON text
	FromJson(value string) error
}

func (obj *isisLspIsReachabilityTlv) Marshal() marshalIsisLspIsReachabilityTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspIsReachabilityTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspIsReachabilityTlv) Unmarshal() unMarshalIsisLspIsReachabilityTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspIsReachabilityTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspIsReachabilityTlv) ToProto() (*otg.IsisLspIsReachabilityTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspIsReachabilityTlv) FromProto(msg *otg.IsisLspIsReachabilityTlv) (IsisLspIsReachabilityTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspIsReachabilityTlv) ToPbText() (string, error) {
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

func (m *unMarshalisisLspIsReachabilityTlv) FromPbText(value string) error {
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

func (m *marshalisisLspIsReachabilityTlv) ToYaml() (string, error) {
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

func (m *unMarshalisisLspIsReachabilityTlv) FromYaml(value string) error {
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

func (m *marshalisisLspIsReachabilityTlv) ToJson() (string, error) {
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

func (m *unMarshalisisLspIsReachabilityTlv) FromJson(value string) error {
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

func (obj *isisLspIsReachabilityTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspIsReachabilityTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspIsReachabilityTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspIsReachabilityTlv) Clone() (IsisLspIsReachabilityTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspIsReachabilityTlv()
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

func (obj *isisLspIsReachabilityTlv) setNil() {
	obj.neighborsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspIsReachabilityTlv is this container describes list of ISIS neighbors and attributes in IS-Reachability TLV (type 2).
type IsisLspIsReachabilityTlv interface {
	Validation
	// msg marshals IsisLspIsReachabilityTlv to protobuf object *otg.IsisLspIsReachabilityTlv
	// and doesn't set defaults
	msg() *otg.IsisLspIsReachabilityTlv
	// setMsg unmarshals IsisLspIsReachabilityTlv from protobuf object *otg.IsisLspIsReachabilityTlv
	// and doesn't set defaults
	setMsg(*otg.IsisLspIsReachabilityTlv) IsisLspIsReachabilityTlv
	// provides marshal interface
	Marshal() marshalIsisLspIsReachabilityTlv
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspIsReachabilityTlv
	// validate validates IsisLspIsReachabilityTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspIsReachabilityTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Neighbors returns IsisLspIsReachabilityTlvIsisLspneighborIterIter, set in IsisLspIsReachabilityTlv
	Neighbors() IsisLspIsReachabilityTlvIsisLspneighborIter
	setNil()
}

// This container describes Intermediate System (IS) neighbors.
// Neighbors returns a []IsisLspneighbor
func (obj *isisLspIsReachabilityTlv) Neighbors() IsisLspIsReachabilityTlvIsisLspneighborIter {
	if len(obj.obj.Neighbors) == 0 {
		obj.obj.Neighbors = []*otg.IsisLspneighbor{}
	}
	if obj.neighborsHolder == nil {
		obj.neighborsHolder = newIsisLspIsReachabilityTlvIsisLspneighborIter(&obj.obj.Neighbors).setMsg(obj)
	}
	return obj.neighborsHolder
}

type isisLspIsReachabilityTlvIsisLspneighborIter struct {
	obj                  *isisLspIsReachabilityTlv
	isisLspneighborSlice []IsisLspneighbor
	fieldPtr             *[]*otg.IsisLspneighbor
}

func newIsisLspIsReachabilityTlvIsisLspneighborIter(ptr *[]*otg.IsisLspneighbor) IsisLspIsReachabilityTlvIsisLspneighborIter {
	return &isisLspIsReachabilityTlvIsisLspneighborIter{fieldPtr: ptr}
}

type IsisLspIsReachabilityTlvIsisLspneighborIter interface {
	setMsg(*isisLspIsReachabilityTlv) IsisLspIsReachabilityTlvIsisLspneighborIter
	Items() []IsisLspneighbor
	Add() IsisLspneighbor
	Append(items ...IsisLspneighbor) IsisLspIsReachabilityTlvIsisLspneighborIter
	Set(index int, newObj IsisLspneighbor) IsisLspIsReachabilityTlvIsisLspneighborIter
	Clear() IsisLspIsReachabilityTlvIsisLspneighborIter
	clearHolderSlice() IsisLspIsReachabilityTlvIsisLspneighborIter
	appendHolderSlice(item IsisLspneighbor) IsisLspIsReachabilityTlvIsisLspneighborIter
}

func (obj *isisLspIsReachabilityTlvIsisLspneighborIter) setMsg(msg *isisLspIsReachabilityTlv) IsisLspIsReachabilityTlvIsisLspneighborIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspneighbor{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspIsReachabilityTlvIsisLspneighborIter) Items() []IsisLspneighbor {
	return obj.isisLspneighborSlice
}

func (obj *isisLspIsReachabilityTlvIsisLspneighborIter) Add() IsisLspneighbor {
	newObj := &otg.IsisLspneighbor{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspneighbor{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspneighborSlice = append(obj.isisLspneighborSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspIsReachabilityTlvIsisLspneighborIter) Append(items ...IsisLspneighbor) IsisLspIsReachabilityTlvIsisLspneighborIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspneighborSlice = append(obj.isisLspneighborSlice, item)
	}
	return obj
}

func (obj *isisLspIsReachabilityTlvIsisLspneighborIter) Set(index int, newObj IsisLspneighbor) IsisLspIsReachabilityTlvIsisLspneighborIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspneighborSlice[index] = newObj
	return obj
}
func (obj *isisLspIsReachabilityTlvIsisLspneighborIter) Clear() IsisLspIsReachabilityTlvIsisLspneighborIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspneighbor{}
		obj.isisLspneighborSlice = []IsisLspneighbor{}
	}
	return obj
}
func (obj *isisLspIsReachabilityTlvIsisLspneighborIter) clearHolderSlice() IsisLspIsReachabilityTlvIsisLspneighborIter {
	if len(obj.isisLspneighborSlice) > 0 {
		obj.isisLspneighborSlice = []IsisLspneighbor{}
	}
	return obj
}
func (obj *isisLspIsReachabilityTlvIsisLspneighborIter) appendHolderSlice(item IsisLspneighbor) IsisLspIsReachabilityTlvIsisLspneighborIter {
	obj.isisLspneighborSlice = append(obj.isisLspneighborSlice, item)
	return obj
}

func (obj *isisLspIsReachabilityTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Neighbors) != 0 {

		if set_default {
			obj.Neighbors().clearHolderSlice()
			for _, item := range obj.obj.Neighbors {
				obj.Neighbors().appendHolderSlice(&isisLspneighbor{obj: item})
			}
		}
		for _, item := range obj.Neighbors().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisLspIsReachabilityTlv) setDefault() {

}
