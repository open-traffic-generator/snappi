package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2NetworkLsa *****
type ospfv2NetworkLsa struct {
	validation
	obj             *otg.Ospfv2NetworkLsa
	marshaller      marshalOspfv2NetworkLsa
	unMarshaller    unMarshalOspfv2NetworkLsa
	headerHolder    Ospfv2LsaHeader
	neighborsHolder Ospfv2NetworkLsaOspfv2LsaNeighborIter
}

func NewOspfv2NetworkLsa() Ospfv2NetworkLsa {
	obj := ospfv2NetworkLsa{obj: &otg.Ospfv2NetworkLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2NetworkLsa) msg() *otg.Ospfv2NetworkLsa {
	return obj.obj
}

func (obj *ospfv2NetworkLsa) setMsg(msg *otg.Ospfv2NetworkLsa) Ospfv2NetworkLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2NetworkLsa struct {
	obj *ospfv2NetworkLsa
}

type marshalOspfv2NetworkLsa interface {
	// ToProto marshals Ospfv2NetworkLsa to protobuf object *otg.Ospfv2NetworkLsa
	ToProto() (*otg.Ospfv2NetworkLsa, error)
	// ToPbText marshals Ospfv2NetworkLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2NetworkLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2NetworkLsa to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2NetworkLsa struct {
	obj *ospfv2NetworkLsa
}

type unMarshalOspfv2NetworkLsa interface {
	// FromProto unmarshals Ospfv2NetworkLsa from protobuf object *otg.Ospfv2NetworkLsa
	FromProto(msg *otg.Ospfv2NetworkLsa) (Ospfv2NetworkLsa, error)
	// FromPbText unmarshals Ospfv2NetworkLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2NetworkLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2NetworkLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv2NetworkLsa) Marshal() marshalOspfv2NetworkLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2NetworkLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2NetworkLsa) Unmarshal() unMarshalOspfv2NetworkLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2NetworkLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2NetworkLsa) ToProto() (*otg.Ospfv2NetworkLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2NetworkLsa) FromProto(msg *otg.Ospfv2NetworkLsa) (Ospfv2NetworkLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2NetworkLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv2NetworkLsa) FromPbText(value string) error {
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

func (m *marshalospfv2NetworkLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv2NetworkLsa) FromYaml(value string) error {
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

func (m *marshalospfv2NetworkLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv2NetworkLsa) FromJson(value string) error {
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

func (obj *ospfv2NetworkLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2NetworkLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2NetworkLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2NetworkLsa) Clone() (Ospfv2NetworkLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2NetworkLsa()
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

func (obj *ospfv2NetworkLsa) setNil() {
	obj.headerHolder = nil
	obj.neighborsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2NetworkLsa is contents of the Network LSA.
type Ospfv2NetworkLsa interface {
	Validation
	// msg marshals Ospfv2NetworkLsa to protobuf object *otg.Ospfv2NetworkLsa
	// and doesn't set defaults
	msg() *otg.Ospfv2NetworkLsa
	// setMsg unmarshals Ospfv2NetworkLsa from protobuf object *otg.Ospfv2NetworkLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv2NetworkLsa) Ospfv2NetworkLsa
	// provides marshal interface
	Marshal() marshalOspfv2NetworkLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2NetworkLsa
	// validate validates Ospfv2NetworkLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2NetworkLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Header returns Ospfv2LsaHeader, set in Ospfv2NetworkLsa.
	// Ospfv2LsaHeader is attributes in LSA Header.
	Header() Ospfv2LsaHeader
	// SetHeader assigns Ospfv2LsaHeader provided by user to Ospfv2NetworkLsa.
	// Ospfv2LsaHeader is attributes in LSA Header.
	SetHeader(value Ospfv2LsaHeader) Ospfv2NetworkLsa
	// HasHeader checks if Header has been set in Ospfv2NetworkLsa
	HasHeader() bool
	// Neighbors returns Ospfv2NetworkLsaOspfv2LsaNeighborIterIter, set in Ospfv2NetworkLsa
	Neighbors() Ospfv2NetworkLsaOspfv2LsaNeighborIter
	setNil()
}

// Contents of the LSA header.
// Header returns a Ospfv2LsaHeader
func (obj *ospfv2NetworkLsa) Header() Ospfv2LsaHeader {
	if obj.obj.Header == nil {
		obj.obj.Header = NewOspfv2LsaHeader().msg()
	}
	if obj.headerHolder == nil {
		obj.headerHolder = &ospfv2LsaHeader{obj: obj.obj.Header}
	}
	return obj.headerHolder
}

// Contents of the LSA header.
// Header returns a Ospfv2LsaHeader
func (obj *ospfv2NetworkLsa) HasHeader() bool {
	return obj.obj.Header != nil
}

// Contents of the LSA header.
// SetHeader sets the Ospfv2LsaHeader value in the Ospfv2NetworkLsa object
func (obj *ospfv2NetworkLsa) SetHeader(value Ospfv2LsaHeader) Ospfv2NetworkLsa {

	obj.headerHolder = nil
	obj.obj.Header = value.msg()

	return obj
}

// Neighbors that are described within the LSA.
// Neighbors returns a []Ospfv2LsaNeighbor
func (obj *ospfv2NetworkLsa) Neighbors() Ospfv2NetworkLsaOspfv2LsaNeighborIter {
	if len(obj.obj.Neighbors) == 0 {
		obj.obj.Neighbors = []*otg.Ospfv2LsaNeighbor{}
	}
	if obj.neighborsHolder == nil {
		obj.neighborsHolder = newOspfv2NetworkLsaOspfv2LsaNeighborIter(&obj.obj.Neighbors).setMsg(obj)
	}
	return obj.neighborsHolder
}

type ospfv2NetworkLsaOspfv2LsaNeighborIter struct {
	obj                    *ospfv2NetworkLsa
	ospfv2LsaNeighborSlice []Ospfv2LsaNeighbor
	fieldPtr               *[]*otg.Ospfv2LsaNeighbor
}

func newOspfv2NetworkLsaOspfv2LsaNeighborIter(ptr *[]*otg.Ospfv2LsaNeighbor) Ospfv2NetworkLsaOspfv2LsaNeighborIter {
	return &ospfv2NetworkLsaOspfv2LsaNeighborIter{fieldPtr: ptr}
}

type Ospfv2NetworkLsaOspfv2LsaNeighborIter interface {
	setMsg(*ospfv2NetworkLsa) Ospfv2NetworkLsaOspfv2LsaNeighborIter
	Items() []Ospfv2LsaNeighbor
	Add() Ospfv2LsaNeighbor
	Append(items ...Ospfv2LsaNeighbor) Ospfv2NetworkLsaOspfv2LsaNeighborIter
	Set(index int, newObj Ospfv2LsaNeighbor) Ospfv2NetworkLsaOspfv2LsaNeighborIter
	Clear() Ospfv2NetworkLsaOspfv2LsaNeighborIter
	clearHolderSlice() Ospfv2NetworkLsaOspfv2LsaNeighborIter
	appendHolderSlice(item Ospfv2LsaNeighbor) Ospfv2NetworkLsaOspfv2LsaNeighborIter
}

func (obj *ospfv2NetworkLsaOspfv2LsaNeighborIter) setMsg(msg *ospfv2NetworkLsa) Ospfv2NetworkLsaOspfv2LsaNeighborIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2LsaNeighbor{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2NetworkLsaOspfv2LsaNeighborIter) Items() []Ospfv2LsaNeighbor {
	return obj.ospfv2LsaNeighborSlice
}

func (obj *ospfv2NetworkLsaOspfv2LsaNeighborIter) Add() Ospfv2LsaNeighbor {
	newObj := &otg.Ospfv2LsaNeighbor{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2LsaNeighbor{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2LsaNeighborSlice = append(obj.ospfv2LsaNeighborSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2NetworkLsaOspfv2LsaNeighborIter) Append(items ...Ospfv2LsaNeighbor) Ospfv2NetworkLsaOspfv2LsaNeighborIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2LsaNeighborSlice = append(obj.ospfv2LsaNeighborSlice, item)
	}
	return obj
}

func (obj *ospfv2NetworkLsaOspfv2LsaNeighborIter) Set(index int, newObj Ospfv2LsaNeighbor) Ospfv2NetworkLsaOspfv2LsaNeighborIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2LsaNeighborSlice[index] = newObj
	return obj
}
func (obj *ospfv2NetworkLsaOspfv2LsaNeighborIter) Clear() Ospfv2NetworkLsaOspfv2LsaNeighborIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2LsaNeighbor{}
		obj.ospfv2LsaNeighborSlice = []Ospfv2LsaNeighbor{}
	}
	return obj
}
func (obj *ospfv2NetworkLsaOspfv2LsaNeighborIter) clearHolderSlice() Ospfv2NetworkLsaOspfv2LsaNeighborIter {
	if len(obj.ospfv2LsaNeighborSlice) > 0 {
		obj.ospfv2LsaNeighborSlice = []Ospfv2LsaNeighbor{}
	}
	return obj
}
func (obj *ospfv2NetworkLsaOspfv2LsaNeighborIter) appendHolderSlice(item Ospfv2LsaNeighbor) Ospfv2NetworkLsaOspfv2LsaNeighborIter {
	obj.ospfv2LsaNeighborSlice = append(obj.ospfv2LsaNeighborSlice, item)
	return obj
}

func (obj *ospfv2NetworkLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Header != nil {

		obj.Header().validateObj(vObj, set_default)
	}

	if len(obj.obj.Neighbors) != 0 {

		if set_default {
			obj.Neighbors().clearHolderSlice()
			for _, item := range obj.obj.Neighbors {
				obj.Neighbors().appendHolderSlice(&ospfv2LsaNeighbor{obj: item})
			}
		}
		for _, item := range obj.Neighbors().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv2NetworkLsa) setDefault() {

}
