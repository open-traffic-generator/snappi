package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2OpaqueLsaTrafficEngineering *****
type ospfv2OpaqueLsaTrafficEngineering struct {
	validation
	obj                  *otg.Ospfv2OpaqueLsaTrafficEngineering
	marshaller           marshalOspfv2OpaqueLsaTrafficEngineering
	unMarshaller         unMarshalOspfv2OpaqueLsaTrafficEngineering
	linkAttributesHolder Ospfv2LsaLinkTrafficEngineering
}

func NewOspfv2OpaqueLsaTrafficEngineering() Ospfv2OpaqueLsaTrafficEngineering {
	obj := ospfv2OpaqueLsaTrafficEngineering{obj: &otg.Ospfv2OpaqueLsaTrafficEngineering{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2OpaqueLsaTrafficEngineering) msg() *otg.Ospfv2OpaqueLsaTrafficEngineering {
	return obj.obj
}

func (obj *ospfv2OpaqueLsaTrafficEngineering) setMsg(msg *otg.Ospfv2OpaqueLsaTrafficEngineering) Ospfv2OpaqueLsaTrafficEngineering {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2OpaqueLsaTrafficEngineering struct {
	obj *ospfv2OpaqueLsaTrafficEngineering
}

type marshalOspfv2OpaqueLsaTrafficEngineering interface {
	// ToProto marshals Ospfv2OpaqueLsaTrafficEngineering to protobuf object *otg.Ospfv2OpaqueLsaTrafficEngineering
	ToProto() (*otg.Ospfv2OpaqueLsaTrafficEngineering, error)
	// ToPbText marshals Ospfv2OpaqueLsaTrafficEngineering to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2OpaqueLsaTrafficEngineering to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2OpaqueLsaTrafficEngineering to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2OpaqueLsaTrafficEngineering struct {
	obj *ospfv2OpaqueLsaTrafficEngineering
}

type unMarshalOspfv2OpaqueLsaTrafficEngineering interface {
	// FromProto unmarshals Ospfv2OpaqueLsaTrafficEngineering from protobuf object *otg.Ospfv2OpaqueLsaTrafficEngineering
	FromProto(msg *otg.Ospfv2OpaqueLsaTrafficEngineering) (Ospfv2OpaqueLsaTrafficEngineering, error)
	// FromPbText unmarshals Ospfv2OpaqueLsaTrafficEngineering from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2OpaqueLsaTrafficEngineering from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2OpaqueLsaTrafficEngineering from JSON text
	FromJson(value string) error
}

func (obj *ospfv2OpaqueLsaTrafficEngineering) Marshal() marshalOspfv2OpaqueLsaTrafficEngineering {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2OpaqueLsaTrafficEngineering{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2OpaqueLsaTrafficEngineering) Unmarshal() unMarshalOspfv2OpaqueLsaTrafficEngineering {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2OpaqueLsaTrafficEngineering{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2OpaqueLsaTrafficEngineering) ToProto() (*otg.Ospfv2OpaqueLsaTrafficEngineering, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2OpaqueLsaTrafficEngineering) FromProto(msg *otg.Ospfv2OpaqueLsaTrafficEngineering) (Ospfv2OpaqueLsaTrafficEngineering, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2OpaqueLsaTrafficEngineering) ToPbText() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaTrafficEngineering) FromPbText(value string) error {
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

func (m *marshalospfv2OpaqueLsaTrafficEngineering) ToYaml() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaTrafficEngineering) FromYaml(value string) error {
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

func (m *marshalospfv2OpaqueLsaTrafficEngineering) ToJson() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaTrafficEngineering) FromJson(value string) error {
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

func (obj *ospfv2OpaqueLsaTrafficEngineering) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaTrafficEngineering) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaTrafficEngineering) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2OpaqueLsaTrafficEngineering) Clone() (Ospfv2OpaqueLsaTrafficEngineering, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2OpaqueLsaTrafficEngineering()
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

func (obj *ospfv2OpaqueLsaTrafficEngineering) setNil() {
	obj.linkAttributesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2OpaqueLsaTrafficEngineering is the decoded Traffic Engineering Opaque LSA contents, Opaque Type 1
// (RFC 3630 Section 2.4). A TE LSA contains one top-level TLV (RFC 3630 Section 2.4):
// either the Router Address TLV, reported as router_address, or the Link TLV,
// reported as link_id plus link_attributes. Only one Link TLV is carried in each LSA
// and each of its sub-TLVs occurs at most once (RFC 3630 Section 2.4.2), so both are
// single-valued here.
type Ospfv2OpaqueLsaTrafficEngineering interface {
	Validation
	// msg marshals Ospfv2OpaqueLsaTrafficEngineering to protobuf object *otg.Ospfv2OpaqueLsaTrafficEngineering
	// and doesn't set defaults
	msg() *otg.Ospfv2OpaqueLsaTrafficEngineering
	// setMsg unmarshals Ospfv2OpaqueLsaTrafficEngineering from protobuf object *otg.Ospfv2OpaqueLsaTrafficEngineering
	// and doesn't set defaults
	setMsg(*otg.Ospfv2OpaqueLsaTrafficEngineering) Ospfv2OpaqueLsaTrafficEngineering
	// provides marshal interface
	Marshal() marshalOspfv2OpaqueLsaTrafficEngineering
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2OpaqueLsaTrafficEngineering
	// validate validates Ospfv2OpaqueLsaTrafficEngineering
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2OpaqueLsaTrafficEngineering, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterAddress returns string, set in Ospfv2OpaqueLsaTrafficEngineering.
	RouterAddress() string
	// SetRouterAddress assigns string provided by user to Ospfv2OpaqueLsaTrafficEngineering
	SetRouterAddress(value string) Ospfv2OpaqueLsaTrafficEngineering
	// HasRouterAddress checks if RouterAddress has been set in Ospfv2OpaqueLsaTrafficEngineering
	HasRouterAddress() bool
	// LinkId returns string, set in Ospfv2OpaqueLsaTrafficEngineering.
	LinkId() string
	// SetLinkId assigns string provided by user to Ospfv2OpaqueLsaTrafficEngineering
	SetLinkId(value string) Ospfv2OpaqueLsaTrafficEngineering
	// HasLinkId checks if LinkId has been set in Ospfv2OpaqueLsaTrafficEngineering
	HasLinkId() bool
	// LinkAttributes returns Ospfv2LsaLinkTrafficEngineering, set in Ospfv2OpaqueLsaTrafficEngineering.
	// Ospfv2LsaLinkTrafficEngineering is traffic engineering attributes for a link, sourced from the Link TLV sub-TLVs of the
	// Traffic Engineering Opaque LSA (RFC 3630 Section 2.5) and the corresponding sub-TLVs of
	// the Extended Link TLV of the OSPFv2 Extended Link Opaque LSA (RFC 9492).
	// This is the attribute set of the link, not a transcription of the sub-TLVs that
	// carried it: a property here names the attribute, and the description of each names
	// the sub-TLV or sub-TLVs it can be sourced from, in either of the two encodings.
	LinkAttributes() Ospfv2LsaLinkTrafficEngineering
	// SetLinkAttributes assigns Ospfv2LsaLinkTrafficEngineering provided by user to Ospfv2OpaqueLsaTrafficEngineering.
	// Ospfv2LsaLinkTrafficEngineering is traffic engineering attributes for a link, sourced from the Link TLV sub-TLVs of the
	// Traffic Engineering Opaque LSA (RFC 3630 Section 2.5) and the corresponding sub-TLVs of
	// the Extended Link TLV of the OSPFv2 Extended Link Opaque LSA (RFC 9492).
	// This is the attribute set of the link, not a transcription of the sub-TLVs that
	// carried it: a property here names the attribute, and the description of each names
	// the sub-TLV or sub-TLVs it can be sourced from, in either of the two encodings.
	SetLinkAttributes(value Ospfv2LsaLinkTrafficEngineering) Ospfv2OpaqueLsaTrafficEngineering
	// HasLinkAttributes checks if LinkAttributes has been set in Ospfv2OpaqueLsaTrafficEngineering
	HasLinkAttributes() bool
	setNil()
}

// The stable IPv4 address of the advertising router, decoded from the Router Address
// TLV, TLV type 1 (RFC 3630 Section 2.4.1). Present instead of link_id and
// link_attributes when the LSA carries a Router Address TLV.
// RouterAddress returns a string
func (obj *ospfv2OpaqueLsaTrafficEngineering) RouterAddress() string {

	return *obj.obj.RouterAddress

}

// The stable IPv4 address of the advertising router, decoded from the Router Address
// TLV, TLV type 1 (RFC 3630 Section 2.4.1). Present instead of link_id and
// link_attributes when the LSA carries a Router Address TLV.
// RouterAddress returns a string
func (obj *ospfv2OpaqueLsaTrafficEngineering) HasRouterAddress() bool {
	return obj.obj.RouterAddress != nil
}

// The stable IPv4 address of the advertising router, decoded from the Router Address
// TLV, TLV type 1 (RFC 3630 Section 2.4.1). Present instead of link_id and
// link_attributes when the LSA carries a Router Address TLV.
// SetRouterAddress sets the string value in the Ospfv2OpaqueLsaTrafficEngineering object
func (obj *ospfv2OpaqueLsaTrafficEngineering) SetRouterAddress(value string) Ospfv2OpaqueLsaTrafficEngineering {

	obj.obj.RouterAddress = &value
	return obj
}

// The identifier of the link the Link TLV describes, decoded from its Link ID
// sub-TLV, sub-type 2 (RFC 3630 Section 2.5.2): the Router ID of the neighbor for a
// point-to-point link, or the interface address of the Designated Router for a
// multi-access link. Correlate the link attributes to the link they describe by
// matching router_lsas[].links[].id of the Router-LSA whose
// header.advertising_router_id equals that of this Opaque LSA.
// LinkId returns a string
func (obj *ospfv2OpaqueLsaTrafficEngineering) LinkId() string {

	return *obj.obj.LinkId

}

// The identifier of the link the Link TLV describes, decoded from its Link ID
// sub-TLV, sub-type 2 (RFC 3630 Section 2.5.2): the Router ID of the neighbor for a
// point-to-point link, or the interface address of the Designated Router for a
// multi-access link. Correlate the link attributes to the link they describe by
// matching router_lsas[].links[].id of the Router-LSA whose
// header.advertising_router_id equals that of this Opaque LSA.
// LinkId returns a string
func (obj *ospfv2OpaqueLsaTrafficEngineering) HasLinkId() bool {
	return obj.obj.LinkId != nil
}

// The identifier of the link the Link TLV describes, decoded from its Link ID
// sub-TLV, sub-type 2 (RFC 3630 Section 2.5.2): the Router ID of the neighbor for a
// point-to-point link, or the interface address of the Designated Router for a
// multi-access link. Correlate the link attributes to the link they describe by
// matching router_lsas[].links[].id of the Router-LSA whose
// header.advertising_router_id equals that of this Opaque LSA.
// SetLinkId sets the string value in the Ospfv2OpaqueLsaTrafficEngineering object
func (obj *ospfv2OpaqueLsaTrafficEngineering) SetLinkId(value string) Ospfv2OpaqueLsaTrafficEngineering {

	obj.obj.LinkId = &value
	return obj
}

// The traffic engineering attributes of the link, decoded from the sub-TLVs of the
// Link TLV, TLV type 2 (RFC 3630 Section 2.5).
// LinkAttributes returns a Ospfv2LsaLinkTrafficEngineering
func (obj *ospfv2OpaqueLsaTrafficEngineering) LinkAttributes() Ospfv2LsaLinkTrafficEngineering {
	if obj.obj.LinkAttributes == nil {
		obj.obj.LinkAttributes = NewOspfv2LsaLinkTrafficEngineering().msg()
	}
	if obj.linkAttributesHolder == nil {
		obj.linkAttributesHolder = &ospfv2LsaLinkTrafficEngineering{obj: obj.obj.LinkAttributes}
	}
	return obj.linkAttributesHolder
}

// The traffic engineering attributes of the link, decoded from the sub-TLVs of the
// Link TLV, TLV type 2 (RFC 3630 Section 2.5).
// LinkAttributes returns a Ospfv2LsaLinkTrafficEngineering
func (obj *ospfv2OpaqueLsaTrafficEngineering) HasLinkAttributes() bool {
	return obj.obj.LinkAttributes != nil
}

// The traffic engineering attributes of the link, decoded from the sub-TLVs of the
// Link TLV, TLV type 2 (RFC 3630 Section 2.5).
// SetLinkAttributes sets the Ospfv2LsaLinkTrafficEngineering value in the Ospfv2OpaqueLsaTrafficEngineering object
func (obj *ospfv2OpaqueLsaTrafficEngineering) SetLinkAttributes(value Ospfv2LsaLinkTrafficEngineering) Ospfv2OpaqueLsaTrafficEngineering {

	obj.linkAttributesHolder = nil
	obj.obj.LinkAttributes = value.msg()

	return obj
}

func (obj *ospfv2OpaqueLsaTrafficEngineering) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RouterAddress != nil {

		err := obj.validateIpv4(obj.RouterAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2OpaqueLsaTrafficEngineering.RouterAddress"))
		}

	}

	if obj.obj.LinkId != nil {

		err := obj.validateIpv4(obj.LinkId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2OpaqueLsaTrafficEngineering.LinkId"))
		}

	}

	if obj.obj.LinkAttributes != nil {

		obj.LinkAttributes().validateObj(vObj, set_default)
	}

}

func (obj *ospfv2OpaqueLsaTrafficEngineering) setDefault() {

}
