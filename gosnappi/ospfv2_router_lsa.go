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
	obj                *otg.Ospfv2RouterLsa
	marshaller         marshalOspfv2RouterLsa
	unMarshaller       unMarshalOspfv2RouterLsa
	headerHolder       Ospfv2LsaHeader
	linksHolder        Ospfv2RouterLsaOspfv2LinkIter
	srCapabilityHolder Ospfv2LsaSrCapability
	prefixSidHolder    Ospfv2LsaPrefixSid
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
	obj.headerHolder = nil
	obj.linksHolder = nil
	obj.srCapabilityHolder = nil
	obj.prefixSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2RouterLsa is contents of the router LSA.
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
	// Header returns Ospfv2LsaHeader, set in Ospfv2RouterLsa.
	// Ospfv2LsaHeader is attributes in LSA Header.
	Header() Ospfv2LsaHeader
	// SetHeader assigns Ospfv2LsaHeader provided by user to Ospfv2RouterLsa.
	// Ospfv2LsaHeader is attributes in LSA Header.
	SetHeader(value Ospfv2LsaHeader) Ospfv2RouterLsa
	// HasHeader checks if Header has been set in Ospfv2RouterLsa
	HasHeader() bool
	// Links returns Ospfv2RouterLsaOspfv2LinkIterIter, set in Ospfv2RouterLsa
	Links() Ospfv2RouterLsaOspfv2LinkIter
	// SrCapability returns Ospfv2LsaSrCapability, set in Ospfv2RouterLsa.
	// Ospfv2LsaSrCapability is the Segment Routing capability learned from the Router Information (RI) Opaque LSA:
	// the SR-Algorithm TLV, SID/Label Range (SRGB) TLV and SR Local Block (SRLB) TLV.
	// Reference: https://datatracker.ietf.org/doc/html/rfc8665.
	SrCapability() Ospfv2LsaSrCapability
	// SetSrCapability assigns Ospfv2LsaSrCapability provided by user to Ospfv2RouterLsa.
	// Ospfv2LsaSrCapability is the Segment Routing capability learned from the Router Information (RI) Opaque LSA:
	// the SR-Algorithm TLV, SID/Label Range (SRGB) TLV and SR Local Block (SRLB) TLV.
	// Reference: https://datatracker.ietf.org/doc/html/rfc8665.
	SetSrCapability(value Ospfv2LsaSrCapability) Ospfv2RouterLsa
	// HasSrCapability checks if SrCapability has been set in Ospfv2RouterLsa
	HasSrCapability() bool
	// PrefixSid returns Ospfv2LsaPrefixSid, set in Ospfv2RouterLsa.
	// Ospfv2LsaPrefixSid is the learned OSPFv2 Prefix-SID and its attributes, decoded from the Prefix-SID sub-TLV of
	// the Extended Prefix Opaque LSA (RFC 8665).
	PrefixSid() Ospfv2LsaPrefixSid
	// SetPrefixSid assigns Ospfv2LsaPrefixSid provided by user to Ospfv2RouterLsa.
	// Ospfv2LsaPrefixSid is the learned OSPFv2 Prefix-SID and its attributes, decoded from the Prefix-SID sub-TLV of
	// the Extended Prefix Opaque LSA (RFC 8665).
	SetPrefixSid(value Ospfv2LsaPrefixSid) Ospfv2RouterLsa
	// HasPrefixSid checks if PrefixSid has been set in Ospfv2RouterLsa
	HasPrefixSid() bool
	setNil()
}

// Contents of the LSA header.
// Header returns a Ospfv2LsaHeader
func (obj *ospfv2RouterLsa) Header() Ospfv2LsaHeader {
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
func (obj *ospfv2RouterLsa) HasHeader() bool {
	return obj.obj.Header != nil
}

// Contents of the LSA header.
// SetHeader sets the Ospfv2LsaHeader value in the Ospfv2RouterLsa object
func (obj *ospfv2RouterLsa) SetHeader(value Ospfv2LsaHeader) Ospfv2RouterLsa {

	obj.headerHolder = nil
	obj.obj.Header = value.msg()

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

// The Segment Routing capability learned for this router, decoded from the Router
// Information (RI) Opaque LSA: the SR-Algorithm TLV, SID/Label Range (SRGB) TLV and
// SR Local Block (SRLB) TLV (RFC 8665).
// SrCapability returns a Ospfv2LsaSrCapability
func (obj *ospfv2RouterLsa) SrCapability() Ospfv2LsaSrCapability {
	if obj.obj.SrCapability == nil {
		obj.obj.SrCapability = NewOspfv2LsaSrCapability().msg()
	}
	if obj.srCapabilityHolder == nil {
		obj.srCapabilityHolder = &ospfv2LsaSrCapability{obj: obj.obj.SrCapability}
	}
	return obj.srCapabilityHolder
}

// The Segment Routing capability learned for this router, decoded from the Router
// Information (RI) Opaque LSA: the SR-Algorithm TLV, SID/Label Range (SRGB) TLV and
// SR Local Block (SRLB) TLV (RFC 8665).
// SrCapability returns a Ospfv2LsaSrCapability
func (obj *ospfv2RouterLsa) HasSrCapability() bool {
	return obj.obj.SrCapability != nil
}

// The Segment Routing capability learned for this router, decoded from the Router
// Information (RI) Opaque LSA: the SR-Algorithm TLV, SID/Label Range (SRGB) TLV and
// SR Local Block (SRLB) TLV (RFC 8665).
// SetSrCapability sets the Ospfv2LsaSrCapability value in the Ospfv2RouterLsa object
func (obj *ospfv2RouterLsa) SetSrCapability(value Ospfv2LsaSrCapability) Ospfv2RouterLsa {

	obj.srCapabilityHolder = nil
	obj.obj.SrCapability = value.msg()

	return obj
}

// The Node/Prefix-SID learned for this router, decoded from the Prefix-SID sub-TLV of
// the OSPFv2 Extended Prefix Opaque LSA advertised for the router's own prefix (RFC 8665).
// PrefixSid returns a Ospfv2LsaPrefixSid
func (obj *ospfv2RouterLsa) PrefixSid() Ospfv2LsaPrefixSid {
	if obj.obj.PrefixSid == nil {
		obj.obj.PrefixSid = NewOspfv2LsaPrefixSid().msg()
	}
	if obj.prefixSidHolder == nil {
		obj.prefixSidHolder = &ospfv2LsaPrefixSid{obj: obj.obj.PrefixSid}
	}
	return obj.prefixSidHolder
}

// The Node/Prefix-SID learned for this router, decoded from the Prefix-SID sub-TLV of
// the OSPFv2 Extended Prefix Opaque LSA advertised for the router's own prefix (RFC 8665).
// PrefixSid returns a Ospfv2LsaPrefixSid
func (obj *ospfv2RouterLsa) HasPrefixSid() bool {
	return obj.obj.PrefixSid != nil
}

// The Node/Prefix-SID learned for this router, decoded from the Prefix-SID sub-TLV of
// the OSPFv2 Extended Prefix Opaque LSA advertised for the router's own prefix (RFC 8665).
// SetPrefixSid sets the Ospfv2LsaPrefixSid value in the Ospfv2RouterLsa object
func (obj *ospfv2RouterLsa) SetPrefixSid(value Ospfv2LsaPrefixSid) Ospfv2RouterLsa {

	obj.prefixSidHolder = nil
	obj.obj.PrefixSid = value.msg()

	return obj
}

func (obj *ospfv2RouterLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Header != nil {

		obj.Header().validateObj(vObj, set_default)
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

	if obj.obj.SrCapability != nil {

		obj.SrCapability().validateObj(vObj, set_default)
	}

	if obj.obj.PrefixSid != nil {

		obj.PrefixSid().validateObj(vObj, set_default)
	}

}

func (obj *ospfv2RouterLsa) setDefault() {

}
