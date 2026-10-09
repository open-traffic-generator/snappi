package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaExtendedLinkAttributes *****
type ospfv2LsaExtendedLinkAttributes struct {
	validation
	obj           *otg.Ospfv2LsaExtendedLinkAttributes
	marshaller    marshalOspfv2LsaExtendedLinkAttributes
	unMarshaller  unMarshalOspfv2LsaExtendedLinkAttributes
	linkMsdHolder Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter
}

func NewOspfv2LsaExtendedLinkAttributes() Ospfv2LsaExtendedLinkAttributes {
	obj := ospfv2LsaExtendedLinkAttributes{obj: &otg.Ospfv2LsaExtendedLinkAttributes{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaExtendedLinkAttributes) msg() *otg.Ospfv2LsaExtendedLinkAttributes {
	return obj.obj
}

func (obj *ospfv2LsaExtendedLinkAttributes) setMsg(msg *otg.Ospfv2LsaExtendedLinkAttributes) Ospfv2LsaExtendedLinkAttributes {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaExtendedLinkAttributes struct {
	obj *ospfv2LsaExtendedLinkAttributes
}

type marshalOspfv2LsaExtendedLinkAttributes interface {
	// ToProto marshals Ospfv2LsaExtendedLinkAttributes to protobuf object *otg.Ospfv2LsaExtendedLinkAttributes
	ToProto() (*otg.Ospfv2LsaExtendedLinkAttributes, error)
	// ToPbText marshals Ospfv2LsaExtendedLinkAttributes to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaExtendedLinkAttributes to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaExtendedLinkAttributes to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaExtendedLinkAttributes struct {
	obj *ospfv2LsaExtendedLinkAttributes
}

type unMarshalOspfv2LsaExtendedLinkAttributes interface {
	// FromProto unmarshals Ospfv2LsaExtendedLinkAttributes from protobuf object *otg.Ospfv2LsaExtendedLinkAttributes
	FromProto(msg *otg.Ospfv2LsaExtendedLinkAttributes) (Ospfv2LsaExtendedLinkAttributes, error)
	// FromPbText unmarshals Ospfv2LsaExtendedLinkAttributes from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaExtendedLinkAttributes from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaExtendedLinkAttributes from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaExtendedLinkAttributes) Marshal() marshalOspfv2LsaExtendedLinkAttributes {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaExtendedLinkAttributes{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaExtendedLinkAttributes) Unmarshal() unMarshalOspfv2LsaExtendedLinkAttributes {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaExtendedLinkAttributes{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaExtendedLinkAttributes) ToProto() (*otg.Ospfv2LsaExtendedLinkAttributes, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaExtendedLinkAttributes) FromProto(msg *otg.Ospfv2LsaExtendedLinkAttributes) (Ospfv2LsaExtendedLinkAttributes, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaExtendedLinkAttributes) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaExtendedLinkAttributes) FromPbText(value string) error {
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

func (m *marshalospfv2LsaExtendedLinkAttributes) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaExtendedLinkAttributes) FromYaml(value string) error {
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

func (m *marshalospfv2LsaExtendedLinkAttributes) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaExtendedLinkAttributes) FromJson(value string) error {
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

func (obj *ospfv2LsaExtendedLinkAttributes) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaExtendedLinkAttributes) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaExtendedLinkAttributes) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaExtendedLinkAttributes) Clone() (Ospfv2LsaExtendedLinkAttributes, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaExtendedLinkAttributes()
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

func (obj *ospfv2LsaExtendedLinkAttributes) setNil() {
	obj.linkMsdHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2LsaExtendedLinkAttributes is traffic engineering attributes of a link, sourced from the sub-TLVs of the OSPFv2
// Extended Link TLV: the application-specific ASLA sub-TLV (RFC 9492 Section 6), the
// direct (non-application-specific) Maximum Bandwidth sub-TLV (RFC 9492 Section 7,
// sub-type 23), and the Link MSD sub-TLV (RFC 8476 Section 3, sub-type 6).
// link_type, local/remote interface addresses, maximum_reservable_bandwidth and
// unreserved_bandwidths are not included: RFC 9492 defines no OSPFv2 encoding,
// application-specific or direct, for any of them on this TLV.
type Ospfv2LsaExtendedLinkAttributes interface {
	Validation
	// msg marshals Ospfv2LsaExtendedLinkAttributes to protobuf object *otg.Ospfv2LsaExtendedLinkAttributes
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaExtendedLinkAttributes
	// setMsg unmarshals Ospfv2LsaExtendedLinkAttributes from protobuf object *otg.Ospfv2LsaExtendedLinkAttributes
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaExtendedLinkAttributes) Ospfv2LsaExtendedLinkAttributes
	// provides marshal interface
	Marshal() marshalOspfv2LsaExtendedLinkAttributes
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaExtendedLinkAttributes
	// validate validates Ospfv2LsaExtendedLinkAttributes
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaExtendedLinkAttributes, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TeMetric returns uint32, set in Ospfv2LsaExtendedLinkAttributes.
	TeMetric() uint32
	// SetTeMetric assigns uint32 provided by user to Ospfv2LsaExtendedLinkAttributes
	SetTeMetric(value uint32) Ospfv2LsaExtendedLinkAttributes
	// HasTeMetric checks if TeMetric has been set in Ospfv2LsaExtendedLinkAttributes
	HasTeMetric() bool
	// MaximumBandwidth returns float32, set in Ospfv2LsaExtendedLinkAttributes.
	MaximumBandwidth() float32
	// SetMaximumBandwidth assigns float32 provided by user to Ospfv2LsaExtendedLinkAttributes
	SetMaximumBandwidth(value float32) Ospfv2LsaExtendedLinkAttributes
	// HasMaximumBandwidth checks if MaximumBandwidth has been set in Ospfv2LsaExtendedLinkAttributes
	HasMaximumBandwidth() bool
	// AdministrativeGroup returns uint32, set in Ospfv2LsaExtendedLinkAttributes.
	AdministrativeGroup() uint32
	// SetAdministrativeGroup assigns uint32 provided by user to Ospfv2LsaExtendedLinkAttributes
	SetAdministrativeGroup(value uint32) Ospfv2LsaExtendedLinkAttributes
	// HasAdministrativeGroup checks if AdministrativeGroup has been set in Ospfv2LsaExtendedLinkAttributes
	HasAdministrativeGroup() bool
	// ExtendedAdministrativeGroup returns []uint32, set in Ospfv2LsaExtendedLinkAttributes.
	ExtendedAdministrativeGroup() []uint32
	// SetExtendedAdministrativeGroup assigns []uint32 provided by user to Ospfv2LsaExtendedLinkAttributes
	SetExtendedAdministrativeGroup(value []uint32) Ospfv2LsaExtendedLinkAttributes
	// Srlg returns []uint32, set in Ospfv2LsaExtendedLinkAttributes.
	Srlg() []uint32
	// SetSrlg assigns []uint32 provided by user to Ospfv2LsaExtendedLinkAttributes
	SetSrlg(value []uint32) Ospfv2LsaExtendedLinkAttributes
	// LinkMsd returns Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIterIter, set in Ospfv2LsaExtendedLinkAttributes
	LinkMsd() Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter
	setNil()
}

// The TE Metric ASLA sub-TLV, type 22 (RFC 9492 Section 6).
// TeMetric returns a uint32
func (obj *ospfv2LsaExtendedLinkAttributes) TeMetric() uint32 {

	return *obj.obj.TeMetric

}

// The TE Metric ASLA sub-TLV, type 22 (RFC 9492 Section 6).
// TeMetric returns a uint32
func (obj *ospfv2LsaExtendedLinkAttributes) HasTeMetric() bool {
	return obj.obj.TeMetric != nil
}

// The TE Metric ASLA sub-TLV, type 22 (RFC 9492 Section 6).
// SetTeMetric sets the uint32 value in the Ospfv2LsaExtendedLinkAttributes object
func (obj *ospfv2LsaExtendedLinkAttributes) SetTeMetric(value uint32) Ospfv2LsaExtendedLinkAttributes {

	obj.obj.TeMetric = &value
	return obj
}

// The Maximum Bandwidth sub-TLV, sub-type 23 (RFC 9492 Section 7), in bytes per
// second. Maximum Bandwidth is an application-independent attribute and MUST NOT
// be advertised inside the ASLA sub-TLV; this direct sub-TLV, using the same format
// as RFC 3630, is how it is advertised on this LSA instead.
// MaximumBandwidth returns a float32
func (obj *ospfv2LsaExtendedLinkAttributes) MaximumBandwidth() float32 {

	return *obj.obj.MaximumBandwidth

}

// The Maximum Bandwidth sub-TLV, sub-type 23 (RFC 9492 Section 7), in bytes per
// second. Maximum Bandwidth is an application-independent attribute and MUST NOT
// be advertised inside the ASLA sub-TLV; this direct sub-TLV, using the same format
// as RFC 3630, is how it is advertised on this LSA instead.
// MaximumBandwidth returns a float32
func (obj *ospfv2LsaExtendedLinkAttributes) HasMaximumBandwidth() bool {
	return obj.obj.MaximumBandwidth != nil
}

// The Maximum Bandwidth sub-TLV, sub-type 23 (RFC 9492 Section 7), in bytes per
// second. Maximum Bandwidth is an application-independent attribute and MUST NOT
// be advertised inside the ASLA sub-TLV; this direct sub-TLV, using the same format
// as RFC 3630, is how it is advertised on this LSA instead.
// SetMaximumBandwidth sets the float32 value in the Ospfv2LsaExtendedLinkAttributes object
func (obj *ospfv2LsaExtendedLinkAttributes) SetMaximumBandwidth(value float32) Ospfv2LsaExtendedLinkAttributes {

	obj.obj.MaximumBandwidth = &value
	return obj
}

// The Administrative Group ASLA sub-TLV, type 19 (RFC 9492 Section 6.2).
// AdministrativeGroup returns a uint32
func (obj *ospfv2LsaExtendedLinkAttributes) AdministrativeGroup() uint32 {

	return *obj.obj.AdministrativeGroup

}

// The Administrative Group ASLA sub-TLV, type 19 (RFC 9492 Section 6.2).
// AdministrativeGroup returns a uint32
func (obj *ospfv2LsaExtendedLinkAttributes) HasAdministrativeGroup() bool {
	return obj.obj.AdministrativeGroup != nil
}

// The Administrative Group ASLA sub-TLV, type 19 (RFC 9492 Section 6.2).
// SetAdministrativeGroup sets the uint32 value in the Ospfv2LsaExtendedLinkAttributes object
func (obj *ospfv2LsaExtendedLinkAttributes) SetAdministrativeGroup(value uint32) Ospfv2LsaExtendedLinkAttributes {

	obj.obj.AdministrativeGroup = &value
	return obj
}

// The Extended Administrative Group ASLA sub-TLV, type 20 (RFC 9492 Section 6.3).
// ExtendedAdministrativeGroup returns a []uint32
func (obj *ospfv2LsaExtendedLinkAttributes) ExtendedAdministrativeGroup() []uint32 {
	if obj.obj.ExtendedAdministrativeGroup == nil {
		obj.obj.ExtendedAdministrativeGroup = make([]uint32, 0)
	}
	return obj.obj.ExtendedAdministrativeGroup
}

// The Extended Administrative Group ASLA sub-TLV, type 20 (RFC 9492 Section 6.3).
// SetExtendedAdministrativeGroup sets the []uint32 value in the Ospfv2LsaExtendedLinkAttributes object
func (obj *ospfv2LsaExtendedLinkAttributes) SetExtendedAdministrativeGroup(value []uint32) Ospfv2LsaExtendedLinkAttributes {

	if obj.obj.ExtendedAdministrativeGroup == nil {
		obj.obj.ExtendedAdministrativeGroup = make([]uint32, 0)
	}
	obj.obj.ExtendedAdministrativeGroup = value

	return obj
}

// The Shared Risk Link Group (SRLG) ASLA sub-TLV, type 11 (RFC 9492 Section 6.1).
// Srlg returns a []uint32
func (obj *ospfv2LsaExtendedLinkAttributes) Srlg() []uint32 {
	if obj.obj.Srlg == nil {
		obj.obj.Srlg = make([]uint32, 0)
	}
	return obj.obj.Srlg
}

// The Shared Risk Link Group (SRLG) ASLA sub-TLV, type 11 (RFC 9492 Section 6.1).
// SetSrlg sets the []uint32 value in the Ospfv2LsaExtendedLinkAttributes object
func (obj *ospfv2LsaExtendedLinkAttributes) SetSrlg(value []uint32) Ospfv2LsaExtendedLinkAttributes {

	if obj.obj.Srlg == nil {
		obj.obj.Srlg = make([]uint32, 0)
	}
	obj.obj.Srlg = value

	return obj
}

// The Link MSD sub-TLV, sub-type 6 (RFC 8476 Section 3).
// LinkMsd returns a []Ospfv2LsaMsd
func (obj *ospfv2LsaExtendedLinkAttributes) LinkMsd() Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter {
	if len(obj.obj.LinkMsd) == 0 {
		obj.obj.LinkMsd = []*otg.Ospfv2LsaMsd{}
	}
	if obj.linkMsdHolder == nil {
		obj.linkMsdHolder = newOspfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter(&obj.obj.LinkMsd).setMsg(obj)
	}
	return obj.linkMsdHolder
}

type ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter struct {
	obj               *ospfv2LsaExtendedLinkAttributes
	ospfv2LsaMsdSlice []Ospfv2LsaMsd
	fieldPtr          *[]*otg.Ospfv2LsaMsd
}

func newOspfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter(ptr *[]*otg.Ospfv2LsaMsd) Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter {
	return &ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter{fieldPtr: ptr}
}

type Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter interface {
	setMsg(*ospfv2LsaExtendedLinkAttributes) Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter
	Items() []Ospfv2LsaMsd
	Add() Ospfv2LsaMsd
	Append(items ...Ospfv2LsaMsd) Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter
	Set(index int, newObj Ospfv2LsaMsd) Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter
	Clear() Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter
	clearHolderSlice() Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter
	appendHolderSlice(item Ospfv2LsaMsd) Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter
}

func (obj *ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter) setMsg(msg *ospfv2LsaExtendedLinkAttributes) Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2LsaMsd{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter) Items() []Ospfv2LsaMsd {
	return obj.ospfv2LsaMsdSlice
}

func (obj *ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter) Add() Ospfv2LsaMsd {
	newObj := &otg.Ospfv2LsaMsd{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2LsaMsd{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2LsaMsdSlice = append(obj.ospfv2LsaMsdSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter) Append(items ...Ospfv2LsaMsd) Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2LsaMsdSlice = append(obj.ospfv2LsaMsdSlice, item)
	}
	return obj
}

func (obj *ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter) Set(index int, newObj Ospfv2LsaMsd) Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2LsaMsdSlice[index] = newObj
	return obj
}
func (obj *ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter) Clear() Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2LsaMsd{}
		obj.ospfv2LsaMsdSlice = []Ospfv2LsaMsd{}
	}
	return obj
}
func (obj *ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter) clearHolderSlice() Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter {
	if len(obj.ospfv2LsaMsdSlice) > 0 {
		obj.ospfv2LsaMsdSlice = []Ospfv2LsaMsd{}
	}
	return obj
}
func (obj *ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter) appendHolderSlice(item Ospfv2LsaMsd) Ospfv2LsaExtendedLinkAttributesOspfv2LsaMsdIter {
	obj.ospfv2LsaMsdSlice = append(obj.ospfv2LsaMsdSlice, item)
	return obj
}

func (obj *ospfv2LsaExtendedLinkAttributes) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.LinkMsd) != 0 {

		if set_default {
			obj.LinkMsd().clearHolderSlice()
			for _, item := range obj.obj.LinkMsd {
				obj.LinkMsd().appendHolderSlice(&ospfv2LsaMsd{obj: item})
			}
		}
		for _, item := range obj.LinkMsd().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv2LsaExtendedLinkAttributes) setDefault() {

}
