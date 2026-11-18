package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2OpaqueLsaRouterInformation *****
type ospfv2OpaqueLsaRouterInformation struct {
	validation
	obj                  *otg.Ospfv2OpaqueLsaRouterInformation
	marshaller           marshalOspfv2OpaqueLsaRouterInformation
	unMarshaller         unMarshalOspfv2OpaqueLsaRouterInformation
	srCapabilityHolder   Ospfv2LsaSrCapability
	riCapabilitiesHolder Ospfv2LsaRiCapabilities
	nodeMsdHolder        Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter
}

func NewOspfv2OpaqueLsaRouterInformation() Ospfv2OpaqueLsaRouterInformation {
	obj := ospfv2OpaqueLsaRouterInformation{obj: &otg.Ospfv2OpaqueLsaRouterInformation{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2OpaqueLsaRouterInformation) msg() *otg.Ospfv2OpaqueLsaRouterInformation {
	return obj.obj
}

func (obj *ospfv2OpaqueLsaRouterInformation) setMsg(msg *otg.Ospfv2OpaqueLsaRouterInformation) Ospfv2OpaqueLsaRouterInformation {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2OpaqueLsaRouterInformation struct {
	obj *ospfv2OpaqueLsaRouterInformation
}

type marshalOspfv2OpaqueLsaRouterInformation interface {
	// ToProto marshals Ospfv2OpaqueLsaRouterInformation to protobuf object *otg.Ospfv2OpaqueLsaRouterInformation
	ToProto() (*otg.Ospfv2OpaqueLsaRouterInformation, error)
	// ToPbText marshals Ospfv2OpaqueLsaRouterInformation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2OpaqueLsaRouterInformation to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2OpaqueLsaRouterInformation to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2OpaqueLsaRouterInformation struct {
	obj *ospfv2OpaqueLsaRouterInformation
}

type unMarshalOspfv2OpaqueLsaRouterInformation interface {
	// FromProto unmarshals Ospfv2OpaqueLsaRouterInformation from protobuf object *otg.Ospfv2OpaqueLsaRouterInformation
	FromProto(msg *otg.Ospfv2OpaqueLsaRouterInformation) (Ospfv2OpaqueLsaRouterInformation, error)
	// FromPbText unmarshals Ospfv2OpaqueLsaRouterInformation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2OpaqueLsaRouterInformation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2OpaqueLsaRouterInformation from JSON text
	FromJson(value string) error
}

func (obj *ospfv2OpaqueLsaRouterInformation) Marshal() marshalOspfv2OpaqueLsaRouterInformation {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2OpaqueLsaRouterInformation{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2OpaqueLsaRouterInformation) Unmarshal() unMarshalOspfv2OpaqueLsaRouterInformation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2OpaqueLsaRouterInformation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2OpaqueLsaRouterInformation) ToProto() (*otg.Ospfv2OpaqueLsaRouterInformation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2OpaqueLsaRouterInformation) FromProto(msg *otg.Ospfv2OpaqueLsaRouterInformation) (Ospfv2OpaqueLsaRouterInformation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2OpaqueLsaRouterInformation) ToPbText() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaRouterInformation) FromPbText(value string) error {
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

func (m *marshalospfv2OpaqueLsaRouterInformation) ToYaml() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaRouterInformation) FromYaml(value string) error {
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

func (m *marshalospfv2OpaqueLsaRouterInformation) ToJson() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaRouterInformation) FromJson(value string) error {
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

func (obj *ospfv2OpaqueLsaRouterInformation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaRouterInformation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaRouterInformation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2OpaqueLsaRouterInformation) Clone() (Ospfv2OpaqueLsaRouterInformation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2OpaqueLsaRouterInformation()
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

func (obj *ospfv2OpaqueLsaRouterInformation) setNil() {
	obj.srCapabilityHolder = nil
	obj.riCapabilitiesHolder = nil
	obj.nodeMsdHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2OpaqueLsaRouterInformation is the decoded contents of one Router Information (RI) Opaque LSA instance, Opaque
// Type 4 (RFC 7770 Section 2).
// Everything reported here describes the router named by the parent LSA's
// header.advertising_router_id, but not necessarily all of it: a router may originate
// more than one RI LSA instance, for example when its capabilities do not fit in one
// LSA, and each instance is a separate Opaque LSA carrying its own Opaque ID
// (RFC 7770 Section 2.1). This object therefore reports what one instance carried,
// not the complete Router Information of the router.
type Ospfv2OpaqueLsaRouterInformation interface {
	Validation
	// msg marshals Ospfv2OpaqueLsaRouterInformation to protobuf object *otg.Ospfv2OpaqueLsaRouterInformation
	// and doesn't set defaults
	msg() *otg.Ospfv2OpaqueLsaRouterInformation
	// setMsg unmarshals Ospfv2OpaqueLsaRouterInformation from protobuf object *otg.Ospfv2OpaqueLsaRouterInformation
	// and doesn't set defaults
	setMsg(*otg.Ospfv2OpaqueLsaRouterInformation) Ospfv2OpaqueLsaRouterInformation
	// provides marshal interface
	Marshal() marshalOspfv2OpaqueLsaRouterInformation
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2OpaqueLsaRouterInformation
	// validate validates Ospfv2OpaqueLsaRouterInformation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2OpaqueLsaRouterInformation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SrCapability returns Ospfv2LsaSrCapability, set in Ospfv2OpaqueLsaRouterInformation.
	// Ospfv2LsaSrCapability is the Segment Routing capability learned from the Router Information (RI) Opaque LSA:
	// the SR-Algorithm TLV, SID/Label Range (SRGB) TLV and SR Local Block (SRLB) TLV.
	// Reference: https://datatracker.ietf.org/doc/html/rfc8665.
	SrCapability() Ospfv2LsaSrCapability
	// SetSrCapability assigns Ospfv2LsaSrCapability provided by user to Ospfv2OpaqueLsaRouterInformation.
	// Ospfv2LsaSrCapability is the Segment Routing capability learned from the Router Information (RI) Opaque LSA:
	// the SR-Algorithm TLV, SID/Label Range (SRGB) TLV and SR Local Block (SRLB) TLV.
	// Reference: https://datatracker.ietf.org/doc/html/rfc8665.
	SetSrCapability(value Ospfv2LsaSrCapability) Ospfv2OpaqueLsaRouterInformation
	// HasSrCapability checks if SrCapability has been set in Ospfv2OpaqueLsaRouterInformation
	HasSrCapability() bool
	// RiCapabilities returns Ospfv2LsaRiCapabilities, set in Ospfv2OpaqueLsaRouterInformation.
	// Ospfv2LsaRiCapabilities is the Router Informational Capabilities (TLV type 1) and Router Functional Capabilities
	// (TLV type 2) learned from the Router Information (RI) Opaque LSA (RFC 7770 Sections 2.5,
	// 2.6).
	RiCapabilities() Ospfv2LsaRiCapabilities
	// SetRiCapabilities assigns Ospfv2LsaRiCapabilities provided by user to Ospfv2OpaqueLsaRouterInformation.
	// Ospfv2LsaRiCapabilities is the Router Informational Capabilities (TLV type 1) and Router Functional Capabilities
	// (TLV type 2) learned from the Router Information (RI) Opaque LSA (RFC 7770 Sections 2.5,
	// 2.6).
	SetRiCapabilities(value Ospfv2LsaRiCapabilities) Ospfv2OpaqueLsaRouterInformation
	// HasRiCapabilities checks if RiCapabilities has been set in Ospfv2OpaqueLsaRouterInformation
	HasRiCapabilities() bool
	// NodeMsd returns Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIterIter, set in Ospfv2OpaqueLsaRouterInformation
	NodeMsd() Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter
	setNil()
}

// The Segment Routing capability of the advertising router, decoded from the
// SR-Algorithm TLV, SID/Label Range (SRGB) TLV and SR Local Block (SRLB) TLV
// (RFC 8665 Sections 3.1, 3.2, 3.3).
// SrCapability returns a Ospfv2LsaSrCapability
func (obj *ospfv2OpaqueLsaRouterInformation) SrCapability() Ospfv2LsaSrCapability {
	if obj.obj.SrCapability == nil {
		obj.obj.SrCapability = NewOspfv2LsaSrCapability().msg()
	}
	if obj.srCapabilityHolder == nil {
		obj.srCapabilityHolder = &ospfv2LsaSrCapability{obj: obj.obj.SrCapability}
	}
	return obj.srCapabilityHolder
}

// The Segment Routing capability of the advertising router, decoded from the
// SR-Algorithm TLV, SID/Label Range (SRGB) TLV and SR Local Block (SRLB) TLV
// (RFC 8665 Sections 3.1, 3.2, 3.3).
// SrCapability returns a Ospfv2LsaSrCapability
func (obj *ospfv2OpaqueLsaRouterInformation) HasSrCapability() bool {
	return obj.obj.SrCapability != nil
}

// The Segment Routing capability of the advertising router, decoded from the
// SR-Algorithm TLV, SID/Label Range (SRGB) TLV and SR Local Block (SRLB) TLV
// (RFC 8665 Sections 3.1, 3.2, 3.3).
// SetSrCapability sets the Ospfv2LsaSrCapability value in the Ospfv2OpaqueLsaRouterInformation object
func (obj *ospfv2OpaqueLsaRouterInformation) SetSrCapability(value Ospfv2LsaSrCapability) Ospfv2OpaqueLsaRouterInformation {

	obj.srCapabilityHolder = nil
	obj.obj.SrCapability = value.msg()

	return obj
}

// The Router Informational and Functional Capabilities of the advertising router,
// decoded from the Router Informational Capabilities TLV and the Router Functional
// Capabilities TLV (RFC 7770 Sections 2.5, 2.6).
// RiCapabilities returns a Ospfv2LsaRiCapabilities
func (obj *ospfv2OpaqueLsaRouterInformation) RiCapabilities() Ospfv2LsaRiCapabilities {
	if obj.obj.RiCapabilities == nil {
		obj.obj.RiCapabilities = NewOspfv2LsaRiCapabilities().msg()
	}
	if obj.riCapabilitiesHolder == nil {
		obj.riCapabilitiesHolder = &ospfv2LsaRiCapabilities{obj: obj.obj.RiCapabilities}
	}
	return obj.riCapabilitiesHolder
}

// The Router Informational and Functional Capabilities of the advertising router,
// decoded from the Router Informational Capabilities TLV and the Router Functional
// Capabilities TLV (RFC 7770 Sections 2.5, 2.6).
// RiCapabilities returns a Ospfv2LsaRiCapabilities
func (obj *ospfv2OpaqueLsaRouterInformation) HasRiCapabilities() bool {
	return obj.obj.RiCapabilities != nil
}

// The Router Informational and Functional Capabilities of the advertising router,
// decoded from the Router Informational Capabilities TLV and the Router Functional
// Capabilities TLV (RFC 7770 Sections 2.5, 2.6).
// SetRiCapabilities sets the Ospfv2LsaRiCapabilities value in the Ospfv2OpaqueLsaRouterInformation object
func (obj *ospfv2OpaqueLsaRouterInformation) SetRiCapabilities(value Ospfv2LsaRiCapabilities) Ospfv2OpaqueLsaRouterInformation {

	obj.riCapabilitiesHolder = nil
	obj.obj.RiCapabilities = value.msg()

	return obj
}

// One or more Maximum SID Depth (MSD) values of the advertising router, decoded from
// the Node MSD TLV (RFC 8476 Section 2).
// NodeMsd returns a []Ospfv2LsaMsd
func (obj *ospfv2OpaqueLsaRouterInformation) NodeMsd() Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter {
	if len(obj.obj.NodeMsd) == 0 {
		obj.obj.NodeMsd = []*otg.Ospfv2LsaMsd{}
	}
	if obj.nodeMsdHolder == nil {
		obj.nodeMsdHolder = newOspfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter(&obj.obj.NodeMsd).setMsg(obj)
	}
	return obj.nodeMsdHolder
}

type ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter struct {
	obj               *ospfv2OpaqueLsaRouterInformation
	ospfv2LsaMsdSlice []Ospfv2LsaMsd
	fieldPtr          *[]*otg.Ospfv2LsaMsd
}

func newOspfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter(ptr *[]*otg.Ospfv2LsaMsd) Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter {
	return &ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter{fieldPtr: ptr}
}

type Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter interface {
	setMsg(*ospfv2OpaqueLsaRouterInformation) Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter
	Items() []Ospfv2LsaMsd
	Add() Ospfv2LsaMsd
	Append(items ...Ospfv2LsaMsd) Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter
	Set(index int, newObj Ospfv2LsaMsd) Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter
	Clear() Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter
	clearHolderSlice() Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter
	appendHolderSlice(item Ospfv2LsaMsd) Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter
}

func (obj *ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter) setMsg(msg *ospfv2OpaqueLsaRouterInformation) Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2LsaMsd{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter) Items() []Ospfv2LsaMsd {
	return obj.ospfv2LsaMsdSlice
}

func (obj *ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter) Add() Ospfv2LsaMsd {
	newObj := &otg.Ospfv2LsaMsd{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2LsaMsd{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2LsaMsdSlice = append(obj.ospfv2LsaMsdSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter) Append(items ...Ospfv2LsaMsd) Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2LsaMsdSlice = append(obj.ospfv2LsaMsdSlice, item)
	}
	return obj
}

func (obj *ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter) Set(index int, newObj Ospfv2LsaMsd) Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2LsaMsdSlice[index] = newObj
	return obj
}
func (obj *ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter) Clear() Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2LsaMsd{}
		obj.ospfv2LsaMsdSlice = []Ospfv2LsaMsd{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter) clearHolderSlice() Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter {
	if len(obj.ospfv2LsaMsdSlice) > 0 {
		obj.ospfv2LsaMsdSlice = []Ospfv2LsaMsd{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter) appendHolderSlice(item Ospfv2LsaMsd) Ospfv2OpaqueLsaRouterInformationOspfv2LsaMsdIter {
	obj.ospfv2LsaMsdSlice = append(obj.ospfv2LsaMsdSlice, item)
	return obj
}

func (obj *ospfv2OpaqueLsaRouterInformation) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SrCapability != nil {

		obj.SrCapability().validateObj(vObj, set_default)
	}

	if obj.obj.RiCapabilities != nil {

		obj.RiCapabilities().validateObj(vObj, set_default)
	}

	if len(obj.obj.NodeMsd) != 0 {

		if set_default {
			obj.NodeMsd().clearHolderSlice()
			for _, item := range obj.obj.NodeMsd {
				obj.NodeMsd().appendHolderSlice(&ospfv2LsaMsd{obj: item})
			}
		}
		for _, item := range obj.NodeMsd().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv2OpaqueLsaRouterInformation) setDefault() {

}
