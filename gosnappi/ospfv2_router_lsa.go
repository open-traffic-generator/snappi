package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2RouterLsa *****
type ospfv2RouterLsa struct {
	validation
	obj               *otg.Ospfv2RouterLsa
	marshaller        marshalOspfv2RouterLsa
	unMarshaller      unMarshalOspfv2RouterLsa
	commonAttrsHolder Ospfv2CommonAttrs
	linksHolder       Ospfv2RouterLsaOspfv2LinkIter
}

func NewOspfv2RouterLsa() Ospfv2RouterLsa {
	obj := ospfv2RouterLsa{obj: &otg.Ospfv2RouterLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2RouterLsa) msg() *otg.Ospfv2RouterLsa {
	return obj.obj
}

func (obj *ospfv2RouterLsa) setMsg(msg *otg.Ospfv2RouterLsa) Ospfv2RouterLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2RouterLsa struct {
	obj *ospfv2RouterLsa
}

type marshalOspfv2RouterLsa interface {
	// ToProto marshals Ospfv2RouterLsa to protobuf object *otg.Ospfv2RouterLsa
	ToProto() (*otg.Ospfv2RouterLsa, error)
	// ToPbText marshals Ospfv2RouterLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2RouterLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2RouterLsa to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2RouterLsa struct {
	obj *ospfv2RouterLsa
}

type unMarshalOspfv2RouterLsa interface {
	// FromProto unmarshals Ospfv2RouterLsa from protobuf object *otg.Ospfv2RouterLsa
	FromProto(msg *otg.Ospfv2RouterLsa) (Ospfv2RouterLsa, error)
	// FromPbText unmarshals Ospfv2RouterLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2RouterLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2RouterLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv2RouterLsa) Marshal() marshalOspfv2RouterLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2RouterLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2RouterLsa) Unmarshal() unMarshalOspfv2RouterLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2RouterLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2RouterLsa) ToProto() (*otg.Ospfv2RouterLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2RouterLsa) FromProto(msg *otg.Ospfv2RouterLsa) (Ospfv2RouterLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2RouterLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv2RouterLsa) FromPbText(value string) error {
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

func (m *marshalospfv2RouterLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv2RouterLsa) FromYaml(value string) error {
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

func (m *marshalospfv2RouterLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv2RouterLsa) FromJson(value string) error {
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

func (obj *ospfv2RouterLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2RouterLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2RouterLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2RouterLsa) Clone() (Ospfv2RouterLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2RouterLsa()
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

func (obj *ospfv2RouterLsa) setNil() {
	obj.commonAttrsHolder = nil
	obj.linksHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2RouterLsa is contents of the router LSA
type Ospfv2RouterLsa interface {
	Validation
	// msg marshals Ospfv2RouterLsa to protobuf object *otg.Ospfv2RouterLsa
	// and doesn't set defaults
	msg() *otg.Ospfv2RouterLsa
	// setMsg unmarshals Ospfv2RouterLsa from protobuf object *otg.Ospfv2RouterLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv2RouterLsa) Ospfv2RouterLsa
	// provides marshal interface
	Marshal() marshalOspfv2RouterLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2RouterLsa
	// validate validates Ospfv2RouterLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2RouterLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CommonAttrs returns Ospfv2CommonAttrs, set in Ospfv2RouterLsa.
	// Ospfv2CommonAttrs is attributes in LSA Header.
	CommonAttrs() Ospfv2CommonAttrs
	// SetCommonAttrs assigns Ospfv2CommonAttrs provided by user to Ospfv2RouterLsa.
	// Ospfv2CommonAttrs is attributes in LSA Header.
	SetCommonAttrs(value Ospfv2CommonAttrs) Ospfv2RouterLsa
	// HasCommonAttrs checks if CommonAttrs has been set in Ospfv2RouterLsa
	HasCommonAttrs() bool
	// Links returns Ospfv2RouterLsaOspfv2LinkIterIter, set in Ospfv2RouterLsa
	Links() Ospfv2RouterLsaOspfv2LinkIter
	setNil()
}

// Contents of the router LSA.
// CommonAttrs returns a Ospfv2CommonAttrs
func (obj *ospfv2RouterLsa) CommonAttrs() Ospfv2CommonAttrs {
	if obj.obj.CommonAttrs == nil {
		obj.obj.CommonAttrs = NewOspfv2CommonAttrs().msg()
	}
	if obj.commonAttrsHolder == nil {
		obj.commonAttrsHolder = &ospfv2CommonAttrs{obj: obj.obj.CommonAttrs}
	}
	return obj.commonAttrsHolder
}

// Contents of the router LSA.
// CommonAttrs returns a Ospfv2CommonAttrs
func (obj *ospfv2RouterLsa) HasCommonAttrs() bool {
	return obj.obj.CommonAttrs != nil
}

// Contents of the router LSA.
// SetCommonAttrs sets the Ospfv2CommonAttrs value in the Ospfv2RouterLsa object
func (obj *ospfv2RouterLsa) SetCommonAttrs(value Ospfv2CommonAttrs) Ospfv2RouterLsa {

	obj.commonAttrsHolder = nil
	obj.obj.CommonAttrs = value.msg()

	return obj
}

// Links that are described within the LSA.
// Links returns a []Ospfv2Link
func (obj *ospfv2RouterLsa) Links() Ospfv2RouterLsaOspfv2LinkIter {
	if len(obj.obj.Links) == 0 {
		obj.obj.Links = []*otg.Ospfv2Link{}
	}
	if obj.linksHolder == nil {
		obj.linksHolder = newOspfv2RouterLsaOspfv2LinkIter(&obj.obj.Links).setMsg(obj)
	}
	return obj.linksHolder
}

type ospfv2RouterLsaOspfv2LinkIter struct {
	obj             *ospfv2RouterLsa
	ospfv2LinkSlice []Ospfv2Link
	fieldPtr        *[]*otg.Ospfv2Link
}

func newOspfv2RouterLsaOspfv2LinkIter(ptr *[]*otg.Ospfv2Link) Ospfv2RouterLsaOspfv2LinkIter {
	return &ospfv2RouterLsaOspfv2LinkIter{fieldPtr: ptr}
}

type Ospfv2RouterLsaOspfv2LinkIter interface {
	setMsg(*ospfv2RouterLsa) Ospfv2RouterLsaOspfv2LinkIter
	Items() []Ospfv2Link
	Add() Ospfv2Link
	Append(items ...Ospfv2Link) Ospfv2RouterLsaOspfv2LinkIter
	Set(index int, newObj Ospfv2Link) Ospfv2RouterLsaOspfv2LinkIter
	Clear() Ospfv2RouterLsaOspfv2LinkIter
	clearHolderSlice() Ospfv2RouterLsaOspfv2LinkIter
	appendHolderSlice(item Ospfv2Link) Ospfv2RouterLsaOspfv2LinkIter
}

func (obj *ospfv2RouterLsaOspfv2LinkIter) setMsg(msg *ospfv2RouterLsa) Ospfv2RouterLsaOspfv2LinkIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2Link{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2RouterLsaOspfv2LinkIter) Items() []Ospfv2Link {
	return obj.ospfv2LinkSlice
}

func (obj *ospfv2RouterLsaOspfv2LinkIter) Add() Ospfv2Link {
	newObj := &otg.Ospfv2Link{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2Link{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2LinkSlice = append(obj.ospfv2LinkSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2RouterLsaOspfv2LinkIter) Append(items ...Ospfv2Link) Ospfv2RouterLsaOspfv2LinkIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2LinkSlice = append(obj.ospfv2LinkSlice, item)
	}
	return obj
}

func (obj *ospfv2RouterLsaOspfv2LinkIter) Set(index int, newObj Ospfv2Link) Ospfv2RouterLsaOspfv2LinkIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2LinkSlice[index] = newObj
	return obj
}
func (obj *ospfv2RouterLsaOspfv2LinkIter) Clear() Ospfv2RouterLsaOspfv2LinkIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2Link{}
		obj.ospfv2LinkSlice = []Ospfv2Link{}
	}
	return obj
}
func (obj *ospfv2RouterLsaOspfv2LinkIter) clearHolderSlice() Ospfv2RouterLsaOspfv2LinkIter {
	if len(obj.ospfv2LinkSlice) > 0 {
		obj.ospfv2LinkSlice = []Ospfv2Link{}
	}
	return obj
}
func (obj *ospfv2RouterLsaOspfv2LinkIter) appendHolderSlice(item Ospfv2Link) Ospfv2RouterLsaOspfv2LinkIter {
	obj.ospfv2LinkSlice = append(obj.ospfv2LinkSlice, item)
	return obj
}

func (obj *ospfv2RouterLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CommonAttrs != nil {

		obj.CommonAttrs().validateObj(vObj, set_default)
	}

	if len(obj.obj.Links) != 0 {

		if set_default {
			obj.Links().clearHolderSlice()
			for _, item := range obj.obj.Links {
				obj.Links().appendHolderSlice(&ospfv2Link{obj: item})
			}
		}
		for _, item := range obj.Links().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv2RouterLsa) setDefault() {

}
