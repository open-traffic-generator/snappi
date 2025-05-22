package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3RouterLsa *****
type ospfv3RouterLsa struct {
	validation
	obj          *otg.Ospfv3RouterLsa
	marshaller   marshalOspfv3RouterLsa
	unMarshaller unMarshalOspfv3RouterLsa
	headerHolder Ospfv3LsaHeader
	linksHolder  Ospfv3RouterLsaOspfv3LinkIter
}

func NewOspfv3RouterLsa() Ospfv3RouterLsa {
	obj := ospfv3RouterLsa{obj: &otg.Ospfv3RouterLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3RouterLsa) msg() *otg.Ospfv3RouterLsa {
	return obj.obj
}

func (obj *ospfv3RouterLsa) setMsg(msg *otg.Ospfv3RouterLsa) Ospfv3RouterLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3RouterLsa struct {
	obj *ospfv3RouterLsa
}

type marshalOspfv3RouterLsa interface {
	// ToProto marshals Ospfv3RouterLsa to protobuf object *otg.Ospfv3RouterLsa
	ToProto() (*otg.Ospfv3RouterLsa, error)
	// ToPbText marshals Ospfv3RouterLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3RouterLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3RouterLsa to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv3RouterLsa to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv3RouterLsa struct {
	obj *ospfv3RouterLsa
}

type unMarshalOspfv3RouterLsa interface {
	// FromProto unmarshals Ospfv3RouterLsa from protobuf object *otg.Ospfv3RouterLsa
	FromProto(msg *otg.Ospfv3RouterLsa) (Ospfv3RouterLsa, error)
	// FromPbText unmarshals Ospfv3RouterLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3RouterLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3RouterLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv3RouterLsa) Marshal() marshalOspfv3RouterLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3RouterLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3RouterLsa) Unmarshal() unMarshalOspfv3RouterLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3RouterLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3RouterLsa) ToProto() (*otg.Ospfv3RouterLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3RouterLsa) FromProto(msg *otg.Ospfv3RouterLsa) (Ospfv3RouterLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3RouterLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv3RouterLsa) FromPbText(value string) error {
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

func (m *marshalospfv3RouterLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv3RouterLsa) FromYaml(value string) error {
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

func (m *marshalospfv3RouterLsa) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalospfv3RouterLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv3RouterLsa) FromJson(value string) error {
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

func (obj *ospfv3RouterLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3RouterLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3RouterLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3RouterLsa) Clone() (Ospfv3RouterLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3RouterLsa()
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

func (obj *ospfv3RouterLsa) setNil() {
	obj.headerHolder = nil
	obj.linksHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3RouterLsa is contents of the router LSA.
type Ospfv3RouterLsa interface {
	Validation
	// msg marshals Ospfv3RouterLsa to protobuf object *otg.Ospfv3RouterLsa
	// and doesn't set defaults
	msg() *otg.Ospfv3RouterLsa
	// setMsg unmarshals Ospfv3RouterLsa from protobuf object *otg.Ospfv3RouterLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv3RouterLsa) Ospfv3RouterLsa
	// provides marshal interface
	Marshal() marshalOspfv3RouterLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3RouterLsa
	// validate validates Ospfv3RouterLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3RouterLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Header returns Ospfv3LsaHeader, set in Ospfv3RouterLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	Header() Ospfv3LsaHeader
	// SetHeader assigns Ospfv3LsaHeader provided by user to Ospfv3RouterLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	SetHeader(value Ospfv3LsaHeader) Ospfv3RouterLsa
	// HasHeader checks if Header has been set in Ospfv3RouterLsa
	HasHeader() bool
	// NeighborRouterId returns string, set in Ospfv3RouterLsa.
	NeighborRouterId() string
	// SetNeighborRouterId assigns string provided by user to Ospfv3RouterLsa
	SetNeighborRouterId(value string) Ospfv3RouterLsa
	// HasNeighborRouterId checks if NeighborRouterId has been set in Ospfv3RouterLsa
	HasNeighborRouterId() bool
	// Links returns Ospfv3RouterLsaOspfv3LinkIterIter, set in Ospfv3RouterLsa
	Links() Ospfv3RouterLsaOspfv3LinkIter
	setNil()
}

// Contents of the LSA header.
// Header returns a Ospfv3LsaHeader
func (obj *ospfv3RouterLsa) Header() Ospfv3LsaHeader {
	if obj.obj.Header == nil {
		obj.obj.Header = NewOspfv3LsaHeader().msg()
	}
	if obj.headerHolder == nil {
		obj.headerHolder = &ospfv3LsaHeader{obj: obj.obj.Header}
	}
	return obj.headerHolder
}

// Contents of the LSA header.
// Header returns a Ospfv3LsaHeader
func (obj *ospfv3RouterLsa) HasHeader() bool {
	return obj.obj.Header != nil
}

// Contents of the LSA header.
// SetHeader sets the Ospfv3LsaHeader value in the Ospfv3RouterLsa object
func (obj *ospfv3RouterLsa) SetHeader(value Ospfv3LsaHeader) Ospfv3RouterLsa {

	obj.headerHolder = nil
	obj.obj.Header = value.msg()

	return obj
}

// Neighbor router id that is described within the LSA.
// NeighborRouterId returns a string
func (obj *ospfv3RouterLsa) NeighborRouterId() string {

	return *obj.obj.NeighborRouterId

}

// Neighbor router id that is described within the LSA.
// NeighborRouterId returns a string
func (obj *ospfv3RouterLsa) HasNeighborRouterId() bool {
	return obj.obj.NeighborRouterId != nil
}

// Neighbor router id that is described within the LSA.
// SetNeighborRouterId sets the string value in the Ospfv3RouterLsa object
func (obj *ospfv3RouterLsa) SetNeighborRouterId(value string) Ospfv3RouterLsa {

	obj.obj.NeighborRouterId = &value
	return obj
}

// Links that are described within the LSA.
// Links returns a []Ospfv3Link
func (obj *ospfv3RouterLsa) Links() Ospfv3RouterLsaOspfv3LinkIter {
	if len(obj.obj.Links) == 0 {
		obj.obj.Links = []*otg.Ospfv3Link{}
	}
	if obj.linksHolder == nil {
		obj.linksHolder = newOspfv3RouterLsaOspfv3LinkIter(&obj.obj.Links).setMsg(obj)
	}
	return obj.linksHolder
}

type ospfv3RouterLsaOspfv3LinkIter struct {
	obj             *ospfv3RouterLsa
	ospfv3LinkSlice []Ospfv3Link
	fieldPtr        *[]*otg.Ospfv3Link
}

func newOspfv3RouterLsaOspfv3LinkIter(ptr *[]*otg.Ospfv3Link) Ospfv3RouterLsaOspfv3LinkIter {
	return &ospfv3RouterLsaOspfv3LinkIter{fieldPtr: ptr}
}

type Ospfv3RouterLsaOspfv3LinkIter interface {
	setMsg(*ospfv3RouterLsa) Ospfv3RouterLsaOspfv3LinkIter
	Items() []Ospfv3Link
	Add() Ospfv3Link
	Append(items ...Ospfv3Link) Ospfv3RouterLsaOspfv3LinkIter
	Set(index int, newObj Ospfv3Link) Ospfv3RouterLsaOspfv3LinkIter
	Clear() Ospfv3RouterLsaOspfv3LinkIter
	clearHolderSlice() Ospfv3RouterLsaOspfv3LinkIter
	appendHolderSlice(item Ospfv3Link) Ospfv3RouterLsaOspfv3LinkIter
}

func (obj *ospfv3RouterLsaOspfv3LinkIter) setMsg(msg *ospfv3RouterLsa) Ospfv3RouterLsaOspfv3LinkIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3Link{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv3RouterLsaOspfv3LinkIter) Items() []Ospfv3Link {
	return obj.ospfv3LinkSlice
}

func (obj *ospfv3RouterLsaOspfv3LinkIter) Add() Ospfv3Link {
	newObj := &otg.Ospfv3Link{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3Link{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3LinkSlice = append(obj.ospfv3LinkSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv3RouterLsaOspfv3LinkIter) Append(items ...Ospfv3Link) Ospfv3RouterLsaOspfv3LinkIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3LinkSlice = append(obj.ospfv3LinkSlice, item)
	}
	return obj
}

func (obj *ospfv3RouterLsaOspfv3LinkIter) Set(index int, newObj Ospfv3Link) Ospfv3RouterLsaOspfv3LinkIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3LinkSlice[index] = newObj
	return obj
}
func (obj *ospfv3RouterLsaOspfv3LinkIter) Clear() Ospfv3RouterLsaOspfv3LinkIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3Link{}
		obj.ospfv3LinkSlice = []Ospfv3Link{}
	}
	return obj
}
func (obj *ospfv3RouterLsaOspfv3LinkIter) clearHolderSlice() Ospfv3RouterLsaOspfv3LinkIter {
	if len(obj.ospfv3LinkSlice) > 0 {
		obj.ospfv3LinkSlice = []Ospfv3Link{}
	}
	return obj
}
func (obj *ospfv3RouterLsaOspfv3LinkIter) appendHolderSlice(item Ospfv3Link) Ospfv3RouterLsaOspfv3LinkIter {
	obj.ospfv3LinkSlice = append(obj.ospfv3LinkSlice, item)
	return obj
}

func (obj *ospfv3RouterLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Header != nil {

		obj.Header().validateObj(vObj, set_default)
	}

	if obj.obj.NeighborRouterId != nil {

		err := obj.validateIpv4(obj.NeighborRouterId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3RouterLsa.NeighborRouterId"))
		}

	}

	if len(obj.obj.Links) != 0 {

		if set_default {
			obj.Links().clearHolderSlice()
			for _, item := range obj.obj.Links {
				obj.Links().appendHolderSlice(&ospfv3Link{obj: item})
			}
		}
		for _, item := range obj.Links().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv3RouterLsa) setDefault() {

}
