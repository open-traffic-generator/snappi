package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspextdNeighbor *****
type isisLspextdNeighbor struct {
	validation
	obj                 *otg.IsisLspextdNeighbor
	marshaller          marshalIsisLspextdNeighbor
	unMarshaller        unMarshalIsisLspextdNeighbor
	adjacencySidsHolder IsisLspextdNeighborIsisLspAdjacencySidIter
}

func NewIsisLspextdNeighbor() IsisLspextdNeighbor {
	obj := isisLspextdNeighbor{obj: &otg.IsisLspextdNeighbor{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspextdNeighbor) msg() *otg.IsisLspextdNeighbor {
	return obj.obj
}

func (obj *isisLspextdNeighbor) setMsg(msg *otg.IsisLspextdNeighbor) IsisLspextdNeighbor {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspextdNeighbor struct {
	obj *isisLspextdNeighbor
}

type marshalIsisLspextdNeighbor interface {
	// ToProto marshals IsisLspextdNeighbor to protobuf object *otg.IsisLspextdNeighbor
	ToProto() (*otg.IsisLspextdNeighbor, error)
	// ToPbText marshals IsisLspextdNeighbor to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspextdNeighbor to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspextdNeighbor to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspextdNeighbor struct {
	obj *isisLspextdNeighbor
}

type unMarshalIsisLspextdNeighbor interface {
	// FromProto unmarshals IsisLspextdNeighbor from protobuf object *otg.IsisLspextdNeighbor
	FromProto(msg *otg.IsisLspextdNeighbor) (IsisLspextdNeighbor, error)
	// FromPbText unmarshals IsisLspextdNeighbor from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspextdNeighbor from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspextdNeighbor from JSON text
	FromJson(value string) error
}

func (obj *isisLspextdNeighbor) Marshal() marshalIsisLspextdNeighbor {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspextdNeighbor{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspextdNeighbor) Unmarshal() unMarshalIsisLspextdNeighbor {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspextdNeighbor{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspextdNeighbor) ToProto() (*otg.IsisLspextdNeighbor, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspextdNeighbor) FromProto(msg *otg.IsisLspextdNeighbor) (IsisLspextdNeighbor, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspextdNeighbor) ToPbText() (string, error) {
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

func (m *unMarshalisisLspextdNeighbor) FromPbText(value string) error {
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

func (m *marshalisisLspextdNeighbor) ToYaml() (string, error) {
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

func (m *unMarshalisisLspextdNeighbor) FromYaml(value string) error {
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

func (m *marshalisisLspextdNeighbor) ToJson() (string, error) {
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

func (m *unMarshalisisLspextdNeighbor) FromJson(value string) error {
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

func (obj *isisLspextdNeighbor) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspextdNeighbor) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspextdNeighbor) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspextdNeighbor) Clone() (IsisLspextdNeighbor, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspextdNeighbor()
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

func (obj *isisLspextdNeighbor) setNil() {
	obj.adjacencySidsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspextdNeighbor is this contains IS neighbors.
type IsisLspextdNeighbor interface {
	Validation
	// msg marshals IsisLspextdNeighbor to protobuf object *otg.IsisLspextdNeighbor
	// and doesn't set defaults
	msg() *otg.IsisLspextdNeighbor
	// setMsg unmarshals IsisLspextdNeighbor from protobuf object *otg.IsisLspextdNeighbor
	// and doesn't set defaults
	setMsg(*otg.IsisLspextdNeighbor) IsisLspextdNeighbor
	// provides marshal interface
	Marshal() marshalIsisLspextdNeighbor
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspextdNeighbor
	// validate validates IsisLspextdNeighbor
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspextdNeighbor, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SystemId returns string, set in IsisLspextdNeighbor.
	SystemId() string
	// SetSystemId assigns string provided by user to IsisLspextdNeighbor
	SetSystemId(value string) IsisLspextdNeighbor
	// HasSystemId checks if SystemId has been set in IsisLspextdNeighbor
	HasSystemId() bool
	// AdjacencySids returns IsisLspextdNeighborIsisLspAdjacencySidIterIter, set in IsisLspextdNeighbor
	AdjacencySids() IsisLspextdNeighborIsisLspAdjacencySidIter
	setNil()
}

// The System ID for this emulated ISIS router, e.g. "640100010000".
// SystemId returns a string
func (obj *isisLspextdNeighbor) SystemId() string {

	return *obj.obj.SystemId

}

// The System ID for this emulated ISIS router, e.g. "640100010000".
// SystemId returns a string
func (obj *isisLspextdNeighbor) HasSystemId() bool {
	return obj.obj.SystemId != nil
}

// The System ID for this emulated ISIS router, e.g. "640100010000".
// SetSystemId sets the string value in the IsisLspextdNeighbor object
func (obj *isisLspextdNeighbor) SetSystemId(value string) IsisLspextdNeighbor {

	obj.obj.SystemId = &value
	return obj
}

// List of segment routing adjacency SIDs.
// AdjacencySids returns a []IsisLspAdjacencySid
func (obj *isisLspextdNeighbor) AdjacencySids() IsisLspextdNeighborIsisLspAdjacencySidIter {
	if len(obj.obj.AdjacencySids) == 0 {
		obj.obj.AdjacencySids = []*otg.IsisLspAdjacencySid{}
	}
	if obj.adjacencySidsHolder == nil {
		obj.adjacencySidsHolder = newIsisLspextdNeighborIsisLspAdjacencySidIter(&obj.obj.AdjacencySids).setMsg(obj)
	}
	return obj.adjacencySidsHolder
}

type isisLspextdNeighborIsisLspAdjacencySidIter struct {
	obj                      *isisLspextdNeighbor
	isisLspAdjacencySidSlice []IsisLspAdjacencySid
	fieldPtr                 *[]*otg.IsisLspAdjacencySid
}

func newIsisLspextdNeighborIsisLspAdjacencySidIter(ptr *[]*otg.IsisLspAdjacencySid) IsisLspextdNeighborIsisLspAdjacencySidIter {
	return &isisLspextdNeighborIsisLspAdjacencySidIter{fieldPtr: ptr}
}

type IsisLspextdNeighborIsisLspAdjacencySidIter interface {
	setMsg(*isisLspextdNeighbor) IsisLspextdNeighborIsisLspAdjacencySidIter
	Items() []IsisLspAdjacencySid
	Add() IsisLspAdjacencySid
	Append(items ...IsisLspAdjacencySid) IsisLspextdNeighborIsisLspAdjacencySidIter
	Set(index int, newObj IsisLspAdjacencySid) IsisLspextdNeighborIsisLspAdjacencySidIter
	Clear() IsisLspextdNeighborIsisLspAdjacencySidIter
	clearHolderSlice() IsisLspextdNeighborIsisLspAdjacencySidIter
	appendHolderSlice(item IsisLspAdjacencySid) IsisLspextdNeighborIsisLspAdjacencySidIter
}

func (obj *isisLspextdNeighborIsisLspAdjacencySidIter) setMsg(msg *isisLspextdNeighbor) IsisLspextdNeighborIsisLspAdjacencySidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspAdjacencySid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspextdNeighborIsisLspAdjacencySidIter) Items() []IsisLspAdjacencySid {
	return obj.isisLspAdjacencySidSlice
}

func (obj *isisLspextdNeighborIsisLspAdjacencySidIter) Add() IsisLspAdjacencySid {
	newObj := &otg.IsisLspAdjacencySid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspAdjacencySid{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspAdjacencySidSlice = append(obj.isisLspAdjacencySidSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspextdNeighborIsisLspAdjacencySidIter) Append(items ...IsisLspAdjacencySid) IsisLspextdNeighborIsisLspAdjacencySidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspAdjacencySidSlice = append(obj.isisLspAdjacencySidSlice, item)
	}
	return obj
}

func (obj *isisLspextdNeighborIsisLspAdjacencySidIter) Set(index int, newObj IsisLspAdjacencySid) IsisLspextdNeighborIsisLspAdjacencySidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspAdjacencySidSlice[index] = newObj
	return obj
}
func (obj *isisLspextdNeighborIsisLspAdjacencySidIter) Clear() IsisLspextdNeighborIsisLspAdjacencySidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspAdjacencySid{}
		obj.isisLspAdjacencySidSlice = []IsisLspAdjacencySid{}
	}
	return obj
}
func (obj *isisLspextdNeighborIsisLspAdjacencySidIter) clearHolderSlice() IsisLspextdNeighborIsisLspAdjacencySidIter {
	if len(obj.isisLspAdjacencySidSlice) > 0 {
		obj.isisLspAdjacencySidSlice = []IsisLspAdjacencySid{}
	}
	return obj
}
func (obj *isisLspextdNeighborIsisLspAdjacencySidIter) appendHolderSlice(item IsisLspAdjacencySid) IsisLspextdNeighborIsisLspAdjacencySidIter {
	obj.isisLspAdjacencySidSlice = append(obj.isisLspAdjacencySidSlice, item)
	return obj
}

func (obj *isisLspextdNeighbor) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SystemId != nil {

		err := obj.validateHex(obj.SystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspextdNeighbor.SystemId"))
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

}

func (obj *isisLspextdNeighbor) setDefault() {

}
