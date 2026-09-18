package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2OpaqueLsaExtendedLink *****
type ospfv2OpaqueLsaExtendedLink struct {
	validation
	obj                  *otg.Ospfv2OpaqueLsaExtendedLink
	marshaller           marshalOspfv2OpaqueLsaExtendedLink
	unMarshaller         unMarshalOspfv2OpaqueLsaExtendedLink
	adjacencySidHolder   Ospfv2LsaAdjacencySid
	linkAttributesHolder Ospfv2LsaLinkTrafficEngineering
}

func NewOspfv2OpaqueLsaExtendedLink() Ospfv2OpaqueLsaExtendedLink {
	obj := ospfv2OpaqueLsaExtendedLink{obj: &otg.Ospfv2OpaqueLsaExtendedLink{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2OpaqueLsaExtendedLink) msg() *otg.Ospfv2OpaqueLsaExtendedLink {
	return obj.obj
}

func (obj *ospfv2OpaqueLsaExtendedLink) setMsg(msg *otg.Ospfv2OpaqueLsaExtendedLink) Ospfv2OpaqueLsaExtendedLink {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2OpaqueLsaExtendedLink struct {
	obj *ospfv2OpaqueLsaExtendedLink
}

type marshalOspfv2OpaqueLsaExtendedLink interface {
	// ToProto marshals Ospfv2OpaqueLsaExtendedLink to protobuf object *otg.Ospfv2OpaqueLsaExtendedLink
	ToProto() (*otg.Ospfv2OpaqueLsaExtendedLink, error)
	// ToPbText marshals Ospfv2OpaqueLsaExtendedLink to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2OpaqueLsaExtendedLink to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2OpaqueLsaExtendedLink to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2OpaqueLsaExtendedLink struct {
	obj *ospfv2OpaqueLsaExtendedLink
}

type unMarshalOspfv2OpaqueLsaExtendedLink interface {
	// FromProto unmarshals Ospfv2OpaqueLsaExtendedLink from protobuf object *otg.Ospfv2OpaqueLsaExtendedLink
	FromProto(msg *otg.Ospfv2OpaqueLsaExtendedLink) (Ospfv2OpaqueLsaExtendedLink, error)
	// FromPbText unmarshals Ospfv2OpaqueLsaExtendedLink from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2OpaqueLsaExtendedLink from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2OpaqueLsaExtendedLink from JSON text
	FromJson(value string) error
}

func (obj *ospfv2OpaqueLsaExtendedLink) Marshal() marshalOspfv2OpaqueLsaExtendedLink {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2OpaqueLsaExtendedLink{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2OpaqueLsaExtendedLink) Unmarshal() unMarshalOspfv2OpaqueLsaExtendedLink {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2OpaqueLsaExtendedLink{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2OpaqueLsaExtendedLink) ToProto() (*otg.Ospfv2OpaqueLsaExtendedLink, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2OpaqueLsaExtendedLink) FromProto(msg *otg.Ospfv2OpaqueLsaExtendedLink) (Ospfv2OpaqueLsaExtendedLink, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2OpaqueLsaExtendedLink) ToPbText() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaExtendedLink) FromPbText(value string) error {
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

func (m *marshalospfv2OpaqueLsaExtendedLink) ToYaml() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaExtendedLink) FromYaml(value string) error {
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

func (m *marshalospfv2OpaqueLsaExtendedLink) ToJson() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaExtendedLink) FromJson(value string) error {
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

func (obj *ospfv2OpaqueLsaExtendedLink) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaExtendedLink) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaExtendedLink) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2OpaqueLsaExtendedLink) Clone() (Ospfv2OpaqueLsaExtendedLink, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2OpaqueLsaExtendedLink()
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

func (obj *ospfv2OpaqueLsaExtendedLink) setNil() {
	obj.adjacencySidHolder = nil
	obj.linkAttributesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2OpaqueLsaExtendedLink is a decoded OSPFv2 Extended Link TLV of an Extended Link Opaque LSA, TLV type 1
// (RFC 7684 Section 3.1).
// The TLV repeats the Link Type, Link ID and Link Data of the Router-LSA link it extends,
// so correlate it to that link by matching link_id, and link_data when the same link_id
// is advertised more than once, against router_lsas[].links[] of the Router-LSA whose
// header.advertising_router_id equals that of this Opaque LSA.
type Ospfv2OpaqueLsaExtendedLink interface {
	Validation
	// msg marshals Ospfv2OpaqueLsaExtendedLink to protobuf object *otg.Ospfv2OpaqueLsaExtendedLink
	// and doesn't set defaults
	msg() *otg.Ospfv2OpaqueLsaExtendedLink
	// setMsg unmarshals Ospfv2OpaqueLsaExtendedLink from protobuf object *otg.Ospfv2OpaqueLsaExtendedLink
	// and doesn't set defaults
	setMsg(*otg.Ospfv2OpaqueLsaExtendedLink) Ospfv2OpaqueLsaExtendedLink
	// provides marshal interface
	Marshal() marshalOspfv2OpaqueLsaExtendedLink
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2OpaqueLsaExtendedLink
	// validate validates Ospfv2OpaqueLsaExtendedLink
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2OpaqueLsaExtendedLink, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LinkType returns Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum, set in Ospfv2OpaqueLsaExtendedLink
	LinkType() Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum
	// SetLinkType assigns Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum provided by user to Ospfv2OpaqueLsaExtendedLink
	SetLinkType(value Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum) Ospfv2OpaqueLsaExtendedLink
	// HasLinkType checks if LinkType has been set in Ospfv2OpaqueLsaExtendedLink
	HasLinkType() bool
	// LinkId returns string, set in Ospfv2OpaqueLsaExtendedLink.
	LinkId() string
	// SetLinkId assigns string provided by user to Ospfv2OpaqueLsaExtendedLink
	SetLinkId(value string) Ospfv2OpaqueLsaExtendedLink
	// HasLinkId checks if LinkId has been set in Ospfv2OpaqueLsaExtendedLink
	HasLinkId() bool
	// LinkData returns string, set in Ospfv2OpaqueLsaExtendedLink.
	LinkData() string
	// SetLinkData assigns string provided by user to Ospfv2OpaqueLsaExtendedLink
	SetLinkData(value string) Ospfv2OpaqueLsaExtendedLink
	// HasLinkData checks if LinkData has been set in Ospfv2OpaqueLsaExtendedLink
	HasLinkData() bool
	// AdjacencySid returns Ospfv2LsaAdjacencySid, set in Ospfv2OpaqueLsaExtendedLink.
	// Ospfv2LsaAdjacencySid is the learned OSPFv2 Adjacency-SID and its attributes, decoded from the Adj-SID / LAN Adj-SID
	// sub-TLV of the Extended Link Opaque LSA (RFC 8665).
	AdjacencySid() Ospfv2LsaAdjacencySid
	// SetAdjacencySid assigns Ospfv2LsaAdjacencySid provided by user to Ospfv2OpaqueLsaExtendedLink.
	// Ospfv2LsaAdjacencySid is the learned OSPFv2 Adjacency-SID and its attributes, decoded from the Adj-SID / LAN Adj-SID
	// sub-TLV of the Extended Link Opaque LSA (RFC 8665).
	SetAdjacencySid(value Ospfv2LsaAdjacencySid) Ospfv2OpaqueLsaExtendedLink
	// HasAdjacencySid checks if AdjacencySid has been set in Ospfv2OpaqueLsaExtendedLink
	HasAdjacencySid() bool
	// LinkAttributes returns Ospfv2LsaLinkTrafficEngineering, set in Ospfv2OpaqueLsaExtendedLink.
	// Ospfv2LsaLinkTrafficEngineering is traffic engineering attributes for a link, decoded from the Link TLV sub-TLVs of the
	// Traffic Engineering Opaque LSA (RFC 3630 Section 2.5) and the corresponding sub-TLVs of
	// the Extended Link TLV of the OSPFv2 Extended Link Opaque LSA (RFC 9492).
	LinkAttributes() Ospfv2LsaLinkTrafficEngineering
	// SetLinkAttributes assigns Ospfv2LsaLinkTrafficEngineering provided by user to Ospfv2OpaqueLsaExtendedLink.
	// Ospfv2LsaLinkTrafficEngineering is traffic engineering attributes for a link, decoded from the Link TLV sub-TLVs of the
	// Traffic Engineering Opaque LSA (RFC 3630 Section 2.5) and the corresponding sub-TLVs of
	// the Extended Link TLV of the OSPFv2 Extended Link Opaque LSA (RFC 9492).
	SetLinkAttributes(value Ospfv2LsaLinkTrafficEngineering) Ospfv2OpaqueLsaExtendedLink
	// HasLinkAttributes checks if LinkAttributes has been set in Ospfv2OpaqueLsaExtendedLink
	HasLinkAttributes() bool
	setNil()
}

type Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum string

// Enum of LinkType on Ospfv2OpaqueLsaExtendedLink
var Ospfv2OpaqueLsaExtendedLinkLinkType = struct {
	POINT_TO_POINT Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum
	TRANSIT        Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum
	STUB           Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum
	VIRTUAL        Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum
}{
	POINT_TO_POINT: Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum("point_to_point"),
	TRANSIT:        Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum("transit"),
	STUB:           Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum("stub"),
	VIRTUAL:        Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum("virtual"),
}

func (obj *ospfv2OpaqueLsaExtendedLink) LinkType() Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum {
	return Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum(obj.obj.LinkType.Enum().String())
}

// The type of the link the TLV extends, decoded from the Link Type field
// (RFC 7684 Section 3.1). Carries the same value as the Type field of the
// corresponding Router-LSA link (RFC 2328 Section A.4.2).
// LinkType returns a string
func (obj *ospfv2OpaqueLsaExtendedLink) HasLinkType() bool {
	return obj.obj.LinkType != nil
}

func (obj *ospfv2OpaqueLsaExtendedLink) SetLinkType(value Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum) Ospfv2OpaqueLsaExtendedLink {
	intValue, ok := otg.Ospfv2OpaqueLsaExtendedLink_LinkType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2OpaqueLsaExtendedLinkLinkTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2OpaqueLsaExtendedLink_LinkType_Enum(intValue)
	obj.obj.LinkType = &enumValue

	return obj
}

// The identifier of the link the TLV extends, decoded from the Link ID field
// (RFC 7684 Section 3.1).
// LinkId returns a string
func (obj *ospfv2OpaqueLsaExtendedLink) LinkId() string {

	return *obj.obj.LinkId

}

// The identifier of the link the TLV extends, decoded from the Link ID field
// (RFC 7684 Section 3.1).
// LinkId returns a string
func (obj *ospfv2OpaqueLsaExtendedLink) HasLinkId() bool {
	return obj.obj.LinkId != nil
}

// The identifier of the link the TLV extends, decoded from the Link ID field
// (RFC 7684 Section 3.1).
// SetLinkId sets the string value in the Ospfv2OpaqueLsaExtendedLink object
func (obj *ospfv2OpaqueLsaExtendedLink) SetLinkId(value string) Ospfv2OpaqueLsaExtendedLink {

	obj.obj.LinkId = &value
	return obj
}

// The data of the link the TLV extends, decoded from the Link Data field
// (RFC 7684 Section 3.1). Distinguishes two links that share a link_id, such as
// parallel point-to-point links to the same neighbor.
// LinkData returns a string
func (obj *ospfv2OpaqueLsaExtendedLink) LinkData() string {

	return *obj.obj.LinkData

}

// The data of the link the TLV extends, decoded from the Link Data field
// (RFC 7684 Section 3.1). Distinguishes two links that share a link_id, such as
// parallel point-to-point links to the same neighbor.
// LinkData returns a string
func (obj *ospfv2OpaqueLsaExtendedLink) HasLinkData() bool {
	return obj.obj.LinkData != nil
}

// The data of the link the TLV extends, decoded from the Link Data field
// (RFC 7684 Section 3.1). Distinguishes two links that share a link_id, such as
// parallel point-to-point links to the same neighbor.
// SetLinkData sets the string value in the Ospfv2OpaqueLsaExtendedLink object
func (obj *ospfv2OpaqueLsaExtendedLink) SetLinkData(value string) Ospfv2OpaqueLsaExtendedLink {

	obj.obj.LinkData = &value
	return obj
}

// The Adjacency-SID advertised for this link, decoded from the Adj-SID sub-TLV,
// sub-type 2, or the LAN Adj-SID sub-TLV, sub-type 3 (RFC 8665 Sections 6.1, 6.2).
// AdjacencySid returns a Ospfv2LsaAdjacencySid
func (obj *ospfv2OpaqueLsaExtendedLink) AdjacencySid() Ospfv2LsaAdjacencySid {
	if obj.obj.AdjacencySid == nil {
		obj.obj.AdjacencySid = NewOspfv2LsaAdjacencySid().msg()
	}
	if obj.adjacencySidHolder == nil {
		obj.adjacencySidHolder = &ospfv2LsaAdjacencySid{obj: obj.obj.AdjacencySid}
	}
	return obj.adjacencySidHolder
}

// The Adjacency-SID advertised for this link, decoded from the Adj-SID sub-TLV,
// sub-type 2, or the LAN Adj-SID sub-TLV, sub-type 3 (RFC 8665 Sections 6.1, 6.2).
// AdjacencySid returns a Ospfv2LsaAdjacencySid
func (obj *ospfv2OpaqueLsaExtendedLink) HasAdjacencySid() bool {
	return obj.obj.AdjacencySid != nil
}

// The Adjacency-SID advertised for this link, decoded from the Adj-SID sub-TLV,
// sub-type 2, or the LAN Adj-SID sub-TLV, sub-type 3 (RFC 8665 Sections 6.1, 6.2).
// SetAdjacencySid sets the Ospfv2LsaAdjacencySid value in the Ospfv2OpaqueLsaExtendedLink object
func (obj *ospfv2OpaqueLsaExtendedLink) SetAdjacencySid(value Ospfv2LsaAdjacencySid) Ospfv2OpaqueLsaExtendedLink {

	obj.adjacencySidHolder = nil
	obj.obj.AdjacencySid = value.msg()

	return obj
}

// The link attributes advertised for this link, decoded from the application-specific
// and legacy link attribute sub-TLVs of the Extended Link TLV (RFC 9492 Section 6).
// LinkAttributes returns a Ospfv2LsaLinkTrafficEngineering
func (obj *ospfv2OpaqueLsaExtendedLink) LinkAttributes() Ospfv2LsaLinkTrafficEngineering {
	if obj.obj.LinkAttributes == nil {
		obj.obj.LinkAttributes = NewOspfv2LsaLinkTrafficEngineering().msg()
	}
	if obj.linkAttributesHolder == nil {
		obj.linkAttributesHolder = &ospfv2LsaLinkTrafficEngineering{obj: obj.obj.LinkAttributes}
	}
	return obj.linkAttributesHolder
}

// The link attributes advertised for this link, decoded from the application-specific
// and legacy link attribute sub-TLVs of the Extended Link TLV (RFC 9492 Section 6).
// LinkAttributes returns a Ospfv2LsaLinkTrafficEngineering
func (obj *ospfv2OpaqueLsaExtendedLink) HasLinkAttributes() bool {
	return obj.obj.LinkAttributes != nil
}

// The link attributes advertised for this link, decoded from the application-specific
// and legacy link attribute sub-TLVs of the Extended Link TLV (RFC 9492 Section 6).
// SetLinkAttributes sets the Ospfv2LsaLinkTrafficEngineering value in the Ospfv2OpaqueLsaExtendedLink object
func (obj *ospfv2OpaqueLsaExtendedLink) SetLinkAttributes(value Ospfv2LsaLinkTrafficEngineering) Ospfv2OpaqueLsaExtendedLink {

	obj.linkAttributesHolder = nil
	obj.obj.LinkAttributes = value.msg()

	return obj
}

func (obj *ospfv2OpaqueLsaExtendedLink) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LinkId != nil {

		err := obj.validateIpv4(obj.LinkId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2OpaqueLsaExtendedLink.LinkId"))
		}

	}

	if obj.obj.LinkData != nil {

		err := obj.validateIpv4(obj.LinkData())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2OpaqueLsaExtendedLink.LinkData"))
		}

	}

	if obj.obj.AdjacencySid != nil {

		obj.AdjacencySid().validateObj(vObj, set_default)
	}

	if obj.obj.LinkAttributes != nil {

		obj.LinkAttributes().validateObj(vObj, set_default)
	}

}

func (obj *ospfv2OpaqueLsaExtendedLink) setDefault() {

}
