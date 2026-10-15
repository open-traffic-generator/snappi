package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2OpaqueLsaTlvInformation *****
type ospfv2OpaqueLsaTlvInformation struct {
	validation
	obj                      *otg.Ospfv2OpaqueLsaTlvInformation
	marshaller               marshalOspfv2OpaqueLsaTlvInformation
	unMarshaller             unMarshalOspfv2OpaqueLsaTlvInformation
	trafficEngineeringHolder Ospfv2OpaqueLsaTrafficEngineering
	routerInformationHolder  Ospfv2OpaqueLsaRouterInformation
	extendedPrefixesHolder   Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter
	extendedLinkHolder       Ospfv2OpaqueLsaExtendedLink
	unknownHolder            Ospfv2OpaqueLsaUnknownOpaqueType
}

func NewOspfv2OpaqueLsaTlvInformation() Ospfv2OpaqueLsaTlvInformation {
	obj := ospfv2OpaqueLsaTlvInformation{obj: &otg.Ospfv2OpaqueLsaTlvInformation{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2OpaqueLsaTlvInformation) msg() *otg.Ospfv2OpaqueLsaTlvInformation {
	return obj.obj
}

func (obj *ospfv2OpaqueLsaTlvInformation) setMsg(msg *otg.Ospfv2OpaqueLsaTlvInformation) Ospfv2OpaqueLsaTlvInformation {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2OpaqueLsaTlvInformation struct {
	obj *ospfv2OpaqueLsaTlvInformation
}

type marshalOspfv2OpaqueLsaTlvInformation interface {
	// ToProto marshals Ospfv2OpaqueLsaTlvInformation to protobuf object *otg.Ospfv2OpaqueLsaTlvInformation
	ToProto() (*otg.Ospfv2OpaqueLsaTlvInformation, error)
	// ToPbText marshals Ospfv2OpaqueLsaTlvInformation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2OpaqueLsaTlvInformation to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2OpaqueLsaTlvInformation to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2OpaqueLsaTlvInformation struct {
	obj *ospfv2OpaqueLsaTlvInformation
}

type unMarshalOspfv2OpaqueLsaTlvInformation interface {
	// FromProto unmarshals Ospfv2OpaqueLsaTlvInformation from protobuf object *otg.Ospfv2OpaqueLsaTlvInformation
	FromProto(msg *otg.Ospfv2OpaqueLsaTlvInformation) (Ospfv2OpaqueLsaTlvInformation, error)
	// FromPbText unmarshals Ospfv2OpaqueLsaTlvInformation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2OpaqueLsaTlvInformation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2OpaqueLsaTlvInformation from JSON text
	FromJson(value string) error
}

func (obj *ospfv2OpaqueLsaTlvInformation) Marshal() marshalOspfv2OpaqueLsaTlvInformation {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2OpaqueLsaTlvInformation{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2OpaqueLsaTlvInformation) Unmarshal() unMarshalOspfv2OpaqueLsaTlvInformation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2OpaqueLsaTlvInformation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2OpaqueLsaTlvInformation) ToProto() (*otg.Ospfv2OpaqueLsaTlvInformation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2OpaqueLsaTlvInformation) FromProto(msg *otg.Ospfv2OpaqueLsaTlvInformation) (Ospfv2OpaqueLsaTlvInformation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2OpaqueLsaTlvInformation) ToPbText() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaTlvInformation) FromPbText(value string) error {
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

func (m *marshalospfv2OpaqueLsaTlvInformation) ToYaml() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaTlvInformation) FromYaml(value string) error {
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

func (m *marshalospfv2OpaqueLsaTlvInformation) ToJson() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaTlvInformation) FromJson(value string) error {
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

func (obj *ospfv2OpaqueLsaTlvInformation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaTlvInformation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaTlvInformation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2OpaqueLsaTlvInformation) Clone() (Ospfv2OpaqueLsaTlvInformation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2OpaqueLsaTlvInformation()
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

func (obj *ospfv2OpaqueLsaTlvInformation) setNil() {
	obj.trafficEngineeringHolder = nil
	obj.routerInformationHolder = nil
	obj.extendedPrefixesHolder = nil
	obj.extendedLinkHolder = nil
	obj.unknownHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2OpaqueLsaTlvInformation is the decoded body of an OSPFv2 Opaque LSA, selected by Opaque Type
// (RFC 5250 Section 3).
// choice names the Opaque Type of the LSA, so reading choice is enough to know what
// to read next. There are two cases:
// 1. The Opaque Type is decoded by this model: choice is also the name of the field
// carrying the decoded payload (traffic_engineering, router_information,
// extended_prefixes, extended_link).
// 2. The Opaque Type is not decoded by this model: choice is unknown, the numeric
// Opaque Type is reported in unknown.opaque_type, and the TLVs of the LSA are
// reported raw in the parent LSA's unknown_tlvs. This is the single escape hatch for
// every Opaque Type this model does not decode, whichever the reason - assigned by
// IANA but not decoded here (grace-LSA Opaque Type 3, L1VPN Opaque Type 5,
// Inter-AS-TE-v2 Opaque Type 6 and the rest), not assigned by IANA, reserved for
// private use, or assigned by IANA after this version of the model.
// The choice values are not a copy of the IANA registry and are not kept in step with
// it; the registry is the authoritative list of assigned Opaque Types:
// https://www.iana.org/assignments/ospf-opaque-types/ospf-opaque-types.xhtml
// unknown_tlvs is not limited to case 2: a TLV of a decoded Opaque Type that this
// model does not itself decode is reported there too, which is why it stays on the
// parent LSA rather than inside this object.
type Ospfv2OpaqueLsaTlvInformation interface {
	Validation
	// msg marshals Ospfv2OpaqueLsaTlvInformation to protobuf object *otg.Ospfv2OpaqueLsaTlvInformation
	// and doesn't set defaults
	msg() *otg.Ospfv2OpaqueLsaTlvInformation
	// setMsg unmarshals Ospfv2OpaqueLsaTlvInformation from protobuf object *otg.Ospfv2OpaqueLsaTlvInformation
	// and doesn't set defaults
	setMsg(*otg.Ospfv2OpaqueLsaTlvInformation) Ospfv2OpaqueLsaTlvInformation
	// provides marshal interface
	Marshal() marshalOspfv2OpaqueLsaTlvInformation
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2OpaqueLsaTlvInformation
	// validate validates Ospfv2OpaqueLsaTlvInformation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2OpaqueLsaTlvInformation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv2OpaqueLsaTlvInformationChoiceEnum, set in Ospfv2OpaqueLsaTlvInformation
	Choice() Ospfv2OpaqueLsaTlvInformationChoiceEnum
	// setChoice assigns Ospfv2OpaqueLsaTlvInformationChoiceEnum provided by user to Ospfv2OpaqueLsaTlvInformation
	setChoice(value Ospfv2OpaqueLsaTlvInformationChoiceEnum) Ospfv2OpaqueLsaTlvInformation
	// HasChoice checks if Choice has been set in Ospfv2OpaqueLsaTlvInformation
	HasChoice() bool
	// TrafficEngineering returns Ospfv2OpaqueLsaTrafficEngineering, set in Ospfv2OpaqueLsaTlvInformation.
	// Ospfv2OpaqueLsaTrafficEngineering is the decoded Traffic Engineering Opaque LSA contents, Opaque Type 1
	// (RFC 3630 Section 2.4). A TE LSA contains one top-level TLV (RFC 3630 Section 2.4):
	// either the Router Address TLV, reported as router_address, or the Link TLV,
	// reported as link_id plus link_attributes. Only one Link TLV is carried in each LSA
	// and each of its sub-TLVs occurs at most once (RFC 3630 Section 2.4.2), so both are
	// single-valued here.
	TrafficEngineering() Ospfv2OpaqueLsaTrafficEngineering
	// SetTrafficEngineering assigns Ospfv2OpaqueLsaTrafficEngineering provided by user to Ospfv2OpaqueLsaTlvInformation.
	// Ospfv2OpaqueLsaTrafficEngineering is the decoded Traffic Engineering Opaque LSA contents, Opaque Type 1
	// (RFC 3630 Section 2.4). A TE LSA contains one top-level TLV (RFC 3630 Section 2.4):
	// either the Router Address TLV, reported as router_address, or the Link TLV,
	// reported as link_id plus link_attributes. Only one Link TLV is carried in each LSA
	// and each of its sub-TLVs occurs at most once (RFC 3630 Section 2.4.2), so both are
	// single-valued here.
	SetTrafficEngineering(value Ospfv2OpaqueLsaTrafficEngineering) Ospfv2OpaqueLsaTlvInformation
	// HasTrafficEngineering checks if TrafficEngineering has been set in Ospfv2OpaqueLsaTlvInformation
	HasTrafficEngineering() bool
	// RouterInformation returns Ospfv2OpaqueLsaRouterInformation, set in Ospfv2OpaqueLsaTlvInformation.
	// Ospfv2OpaqueLsaRouterInformation is the decoded contents of one Router Information (RI) Opaque LSA instance, Opaque
	// Type 4 (RFC 7770 Section 2).
	// Everything reported here describes the router named by the parent LSA's
	// header.advertising_router_id, but not necessarily all of it: a router may originate
	// more than one RI LSA instance, for example when its capabilities do not fit in one
	// LSA, and each instance is a separate Opaque LSA carrying its own Opaque ID
	// (RFC 7770 Section 2.1). This object therefore reports what one instance carried,
	// not the complete Router Information of the router.
	RouterInformation() Ospfv2OpaqueLsaRouterInformation
	// SetRouterInformation assigns Ospfv2OpaqueLsaRouterInformation provided by user to Ospfv2OpaqueLsaTlvInformation.
	// Ospfv2OpaqueLsaRouterInformation is the decoded contents of one Router Information (RI) Opaque LSA instance, Opaque
	// Type 4 (RFC 7770 Section 2).
	// Everything reported here describes the router named by the parent LSA's
	// header.advertising_router_id, but not necessarily all of it: a router may originate
	// more than one RI LSA instance, for example when its capabilities do not fit in one
	// LSA, and each instance is a separate Opaque LSA carrying its own Opaque ID
	// (RFC 7770 Section 2.1). This object therefore reports what one instance carried,
	// not the complete Router Information of the router.
	SetRouterInformation(value Ospfv2OpaqueLsaRouterInformation) Ospfv2OpaqueLsaTlvInformation
	// HasRouterInformation checks if RouterInformation has been set in Ospfv2OpaqueLsaTlvInformation
	HasRouterInformation() bool
	// ExtendedPrefixes returns Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIterIter, set in Ospfv2OpaqueLsaTlvInformation
	ExtendedPrefixes() Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter
	// ExtendedLink returns Ospfv2OpaqueLsaExtendedLink, set in Ospfv2OpaqueLsaTlvInformation.
	// Ospfv2OpaqueLsaExtendedLink is a decoded OSPFv2 Extended Link TLV of an Extended Link Opaque LSA, TLV type 1
	// (RFC 7684 Section 3.1).
	// The TLV repeats the Link Type, Link ID and Link Data of the Router-LSA link it extends,
	// so correlate it to that link by matching link_id, and link_data when the same link_id
	// is advertised more than once, against router_lsas[].links[] of the Router-LSA whose
	// header.advertising_router_id equals that of this Opaque LSA.
	ExtendedLink() Ospfv2OpaqueLsaExtendedLink
	// SetExtendedLink assigns Ospfv2OpaqueLsaExtendedLink provided by user to Ospfv2OpaqueLsaTlvInformation.
	// Ospfv2OpaqueLsaExtendedLink is a decoded OSPFv2 Extended Link TLV of an Extended Link Opaque LSA, TLV type 1
	// (RFC 7684 Section 3.1).
	// The TLV repeats the Link Type, Link ID and Link Data of the Router-LSA link it extends,
	// so correlate it to that link by matching link_id, and link_data when the same link_id
	// is advertised more than once, against router_lsas[].links[] of the Router-LSA whose
	// header.advertising_router_id equals that of this Opaque LSA.
	SetExtendedLink(value Ospfv2OpaqueLsaExtendedLink) Ospfv2OpaqueLsaTlvInformation
	// HasExtendedLink checks if ExtendedLink has been set in Ospfv2OpaqueLsaTlvInformation
	HasExtendedLink() bool
	// Unknown returns Ospfv2OpaqueLsaUnknownOpaqueType, set in Ospfv2OpaqueLsaTlvInformation.
	// Ospfv2OpaqueLsaUnknownOpaqueType is the Opaque Type of an OSPFv2 Opaque LSA that this model does not decode
	// (RFC 5250 Section 3).
	// Reported as the numeric wire value, so an Opaque Type this model does not decode
	// is still fully identifiable, whichever the reason - assigned by IANA but not
	// decoded here, not assigned by IANA, reserved for private use, or assigned by IANA
	// after this version of the model.
	Unknown() Ospfv2OpaqueLsaUnknownOpaqueType
	// SetUnknown assigns Ospfv2OpaqueLsaUnknownOpaqueType provided by user to Ospfv2OpaqueLsaTlvInformation.
	// Ospfv2OpaqueLsaUnknownOpaqueType is the Opaque Type of an OSPFv2 Opaque LSA that this model does not decode
	// (RFC 5250 Section 3).
	// Reported as the numeric wire value, so an Opaque Type this model does not decode
	// is still fully identifiable, whichever the reason - assigned by IANA but not
	// decoded here, not assigned by IANA, reserved for private use, or assigned by IANA
	// after this version of the model.
	SetUnknown(value Ospfv2OpaqueLsaUnknownOpaqueType) Ospfv2OpaqueLsaTlvInformation
	// HasUnknown checks if Unknown has been set in Ospfv2OpaqueLsaTlvInformation
	HasUnknown() bool
	setNil()
}

type Ospfv2OpaqueLsaTlvInformationChoiceEnum string

// Enum of Choice on Ospfv2OpaqueLsaTlvInformation
var Ospfv2OpaqueLsaTlvInformationChoice = struct {
	TRAFFIC_ENGINEERING Ospfv2OpaqueLsaTlvInformationChoiceEnum
	ROUTER_INFORMATION  Ospfv2OpaqueLsaTlvInformationChoiceEnum
	EXTENDED_PREFIXES   Ospfv2OpaqueLsaTlvInformationChoiceEnum
	EXTENDED_LINK       Ospfv2OpaqueLsaTlvInformationChoiceEnum
	UNKNOWN             Ospfv2OpaqueLsaTlvInformationChoiceEnum
}{
	TRAFFIC_ENGINEERING: Ospfv2OpaqueLsaTlvInformationChoiceEnum("traffic_engineering"),
	ROUTER_INFORMATION:  Ospfv2OpaqueLsaTlvInformationChoiceEnum("router_information"),
	EXTENDED_PREFIXES:   Ospfv2OpaqueLsaTlvInformationChoiceEnum("extended_prefixes"),
	EXTENDED_LINK:       Ospfv2OpaqueLsaTlvInformationChoiceEnum("extended_link"),
	UNKNOWN:             Ospfv2OpaqueLsaTlvInformationChoiceEnum("unknown"),
}

func (obj *ospfv2OpaqueLsaTlvInformation) Choice() Ospfv2OpaqueLsaTlvInformationChoiceEnum {
	return Ospfv2OpaqueLsaTlvInformationChoiceEnum(obj.obj.Choice.Enum().String())
}

// The Opaque Type carried in the most significant octet of the LSA's Link State ID (RFC 5250 Section 3), when this model decodes that Opaque Type; unknown otherwise.
// Choice returns a string
func (obj *ospfv2OpaqueLsaTlvInformation) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv2OpaqueLsaTlvInformation) setChoice(value Ospfv2OpaqueLsaTlvInformationChoiceEnum) Ospfv2OpaqueLsaTlvInformation {
	intValue, ok := otg.Ospfv2OpaqueLsaTlvInformation_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2OpaqueLsaTlvInformationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2OpaqueLsaTlvInformation_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Unknown = nil
	obj.unknownHolder = nil
	obj.obj.ExtendedLink = nil
	obj.extendedLinkHolder = nil
	obj.obj.ExtendedPrefixes = nil
	obj.extendedPrefixesHolder = nil
	obj.obj.RouterInformation = nil
	obj.routerInformationHolder = nil
	obj.obj.TrafficEngineering = nil
	obj.trafficEngineeringHolder = nil

	if value == Ospfv2OpaqueLsaTlvInformationChoice.TRAFFIC_ENGINEERING {
		obj.obj.TrafficEngineering = NewOspfv2OpaqueLsaTrafficEngineering().msg()
	}

	if value == Ospfv2OpaqueLsaTlvInformationChoice.ROUTER_INFORMATION {
		obj.obj.RouterInformation = NewOspfv2OpaqueLsaRouterInformation().msg()
	}

	if value == Ospfv2OpaqueLsaTlvInformationChoice.EXTENDED_PREFIXES {
		obj.obj.ExtendedPrefixes = []*otg.Ospfv2OpaqueLsaExtendedPrefix{}
	}

	if value == Ospfv2OpaqueLsaTlvInformationChoice.EXTENDED_LINK {
		obj.obj.ExtendedLink = NewOspfv2OpaqueLsaExtendedLink().msg()
	}

	if value == Ospfv2OpaqueLsaTlvInformationChoice.UNKNOWN {
		obj.obj.Unknown = NewOspfv2OpaqueLsaUnknownOpaqueType().msg()
	}

	return obj
}

// The decoded contents of a Traffic Engineering Opaque LSA
// (RFC 3630 Section 2.4).
// TrafficEngineering returns a Ospfv2OpaqueLsaTrafficEngineering
func (obj *ospfv2OpaqueLsaTlvInformation) TrafficEngineering() Ospfv2OpaqueLsaTrafficEngineering {
	if obj.obj.TrafficEngineering == nil {
		obj.setChoice(Ospfv2OpaqueLsaTlvInformationChoice.TRAFFIC_ENGINEERING)
	}
	if obj.trafficEngineeringHolder == nil {
		obj.trafficEngineeringHolder = &ospfv2OpaqueLsaTrafficEngineering{obj: obj.obj.TrafficEngineering}
	}
	return obj.trafficEngineeringHolder
}

// The decoded contents of a Traffic Engineering Opaque LSA
// (RFC 3630 Section 2.4).
// TrafficEngineering returns a Ospfv2OpaqueLsaTrafficEngineering
func (obj *ospfv2OpaqueLsaTlvInformation) HasTrafficEngineering() bool {
	return obj.obj.TrafficEngineering != nil
}

// The decoded contents of a Traffic Engineering Opaque LSA
// (RFC 3630 Section 2.4).
// SetTrafficEngineering sets the Ospfv2OpaqueLsaTrafficEngineering value in the Ospfv2OpaqueLsaTlvInformation object
func (obj *ospfv2OpaqueLsaTlvInformation) SetTrafficEngineering(value Ospfv2OpaqueLsaTrafficEngineering) Ospfv2OpaqueLsaTlvInformation {
	obj.setChoice(Ospfv2OpaqueLsaTlvInformationChoice.TRAFFIC_ENGINEERING)
	obj.trafficEngineeringHolder = nil
	obj.obj.TrafficEngineering = value.msg()

	return obj
}

// The decoded contents of this Router Information (RI) Opaque LSA instance
// (RFC 7770 Section 2). Multiple Router Information LSA instances may be
// originated by the same router, with each instance represented by a separate
// Opaque LSA, distinguished by the Opaque ID reported in the parent LSA's id.
// The contents describe the router identified by the parent LSA's
// header.advertising_router_id. Correlate them to the Router-LSA of that router
// by matching router_lsas[].header.advertising_router_id, and expect to merge
// the contents of every RI LSA instance that router originated in this scope.
// RouterInformation returns a Ospfv2OpaqueLsaRouterInformation
func (obj *ospfv2OpaqueLsaTlvInformation) RouterInformation() Ospfv2OpaqueLsaRouterInformation {
	if obj.obj.RouterInformation == nil {
		obj.setChoice(Ospfv2OpaqueLsaTlvInformationChoice.ROUTER_INFORMATION)
	}
	if obj.routerInformationHolder == nil {
		obj.routerInformationHolder = &ospfv2OpaqueLsaRouterInformation{obj: obj.obj.RouterInformation}
	}
	return obj.routerInformationHolder
}

// The decoded contents of this Router Information (RI) Opaque LSA instance
// (RFC 7770 Section 2). Multiple Router Information LSA instances may be
// originated by the same router, with each instance represented by a separate
// Opaque LSA, distinguished by the Opaque ID reported in the parent LSA's id.
// The contents describe the router identified by the parent LSA's
// header.advertising_router_id. Correlate them to the Router-LSA of that router
// by matching router_lsas[].header.advertising_router_id, and expect to merge
// the contents of every RI LSA instance that router originated in this scope.
// RouterInformation returns a Ospfv2OpaqueLsaRouterInformation
func (obj *ospfv2OpaqueLsaTlvInformation) HasRouterInformation() bool {
	return obj.obj.RouterInformation != nil
}

// The decoded contents of this Router Information (RI) Opaque LSA instance
// (RFC 7770 Section 2). Multiple Router Information LSA instances may be
// originated by the same router, with each instance represented by a separate
// Opaque LSA, distinguished by the Opaque ID reported in the parent LSA's id.
// The contents describe the router identified by the parent LSA's
// header.advertising_router_id. Correlate them to the Router-LSA of that router
// by matching router_lsas[].header.advertising_router_id, and expect to merge
// the contents of every RI LSA instance that router originated in this scope.
// SetRouterInformation sets the Ospfv2OpaqueLsaRouterInformation value in the Ospfv2OpaqueLsaTlvInformation object
func (obj *ospfv2OpaqueLsaTlvInformation) SetRouterInformation(value Ospfv2OpaqueLsaRouterInformation) Ospfv2OpaqueLsaTlvInformation {
	obj.setChoice(Ospfv2OpaqueLsaTlvInformationChoice.ROUTER_INFORMATION)
	obj.routerInformationHolder = nil
	obj.obj.RouterInformation = value.msg()

	return obj
}

// The decoded OSPFv2 Extended Prefix TLVs of an Extended Prefix Opaque LSA.
// One Extended Prefix Opaque LSA can carry more than one Extended Prefix TLV, each
// describing a different prefix (RFC 7684 Section 2.1).
// ExtendedPrefixes returns a []Ospfv2OpaqueLsaExtendedPrefix
func (obj *ospfv2OpaqueLsaTlvInformation) ExtendedPrefixes() Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter {
	if len(obj.obj.ExtendedPrefixes) == 0 {
		obj.setChoice(Ospfv2OpaqueLsaTlvInformationChoice.EXTENDED_PREFIXES)
	}
	if obj.extendedPrefixesHolder == nil {
		obj.extendedPrefixesHolder = newOspfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter(&obj.obj.ExtendedPrefixes).setMsg(obj)
	}
	return obj.extendedPrefixesHolder
}

type ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter struct {
	obj                                *ospfv2OpaqueLsaTlvInformation
	ospfv2OpaqueLsaExtendedPrefixSlice []Ospfv2OpaqueLsaExtendedPrefix
	fieldPtr                           *[]*otg.Ospfv2OpaqueLsaExtendedPrefix
}

func newOspfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter(ptr *[]*otg.Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter {
	return &ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter{fieldPtr: ptr}
}

type Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter interface {
	setMsg(*ospfv2OpaqueLsaTlvInformation) Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter
	Items() []Ospfv2OpaqueLsaExtendedPrefix
	Add() Ospfv2OpaqueLsaExtendedPrefix
	Append(items ...Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter
	Set(index int, newObj Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter
	Clear() Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter
	clearHolderSlice() Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter
	appendHolderSlice(item Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter
}

func (obj *ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter) setMsg(msg *ospfv2OpaqueLsaTlvInformation) Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2OpaqueLsaExtendedPrefix{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter) Items() []Ospfv2OpaqueLsaExtendedPrefix {
	return obj.ospfv2OpaqueLsaExtendedPrefixSlice
}

func (obj *ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter) Add() Ospfv2OpaqueLsaExtendedPrefix {
	newObj := &otg.Ospfv2OpaqueLsaExtendedPrefix{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2OpaqueLsaExtendedPrefix{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2OpaqueLsaExtendedPrefixSlice = append(obj.ospfv2OpaqueLsaExtendedPrefixSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter) Append(items ...Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2OpaqueLsaExtendedPrefixSlice = append(obj.ospfv2OpaqueLsaExtendedPrefixSlice, item)
	}
	return obj
}

func (obj *ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter) Set(index int, newObj Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2OpaqueLsaExtendedPrefixSlice[index] = newObj
	return obj
}
func (obj *ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter) Clear() Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2OpaqueLsaExtendedPrefix{}
		obj.ospfv2OpaqueLsaExtendedPrefixSlice = []Ospfv2OpaqueLsaExtendedPrefix{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter) clearHolderSlice() Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter {
	if len(obj.ospfv2OpaqueLsaExtendedPrefixSlice) > 0 {
		obj.ospfv2OpaqueLsaExtendedPrefixSlice = []Ospfv2OpaqueLsaExtendedPrefix{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter) appendHolderSlice(item Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaTlvInformationOspfv2OpaqueLsaExtendedPrefixIter {
	obj.ospfv2OpaqueLsaExtendedPrefixSlice = append(obj.ospfv2OpaqueLsaExtendedPrefixSlice, item)
	return obj
}

// The decoded OSPFv2 Extended Link TLV of an Extended Link Opaque LSA. Only one
// Extended Link TLV is advertised in each Extended Link Opaque LSA
// (RFC 7684 Section 3.1), so this is a single object rather than a list.
// ExtendedLink returns a Ospfv2OpaqueLsaExtendedLink
func (obj *ospfv2OpaqueLsaTlvInformation) ExtendedLink() Ospfv2OpaqueLsaExtendedLink {
	if obj.obj.ExtendedLink == nil {
		obj.setChoice(Ospfv2OpaqueLsaTlvInformationChoice.EXTENDED_LINK)
	}
	if obj.extendedLinkHolder == nil {
		obj.extendedLinkHolder = &ospfv2OpaqueLsaExtendedLink{obj: obj.obj.ExtendedLink}
	}
	return obj.extendedLinkHolder
}

// The decoded OSPFv2 Extended Link TLV of an Extended Link Opaque LSA. Only one
// Extended Link TLV is advertised in each Extended Link Opaque LSA
// (RFC 7684 Section 3.1), so this is a single object rather than a list.
// ExtendedLink returns a Ospfv2OpaqueLsaExtendedLink
func (obj *ospfv2OpaqueLsaTlvInformation) HasExtendedLink() bool {
	return obj.obj.ExtendedLink != nil
}

// The decoded OSPFv2 Extended Link TLV of an Extended Link Opaque LSA. Only one
// Extended Link TLV is advertised in each Extended Link Opaque LSA
// (RFC 7684 Section 3.1), so this is a single object rather than a list.
// SetExtendedLink sets the Ospfv2OpaqueLsaExtendedLink value in the Ospfv2OpaqueLsaTlvInformation object
func (obj *ospfv2OpaqueLsaTlvInformation) SetExtendedLink(value Ospfv2OpaqueLsaExtendedLink) Ospfv2OpaqueLsaTlvInformation {
	obj.setChoice(Ospfv2OpaqueLsaTlvInformationChoice.EXTENDED_LINK)
	obj.extendedLinkHolder = nil
	obj.obj.ExtendedLink = value.msg()

	return obj
}

// The Opaque Type when this model does not decode it (RFC 5250 Section 3). The
// body of such an LSA is not decoded, so its TLVs are reported raw in the parent
// LSA's unknown_tlvs.
// Unknown returns a Ospfv2OpaqueLsaUnknownOpaqueType
func (obj *ospfv2OpaqueLsaTlvInformation) Unknown() Ospfv2OpaqueLsaUnknownOpaqueType {
	if obj.obj.Unknown == nil {
		obj.setChoice(Ospfv2OpaqueLsaTlvInformationChoice.UNKNOWN)
	}
	if obj.unknownHolder == nil {
		obj.unknownHolder = &ospfv2OpaqueLsaUnknownOpaqueType{obj: obj.obj.Unknown}
	}
	return obj.unknownHolder
}

// The Opaque Type when this model does not decode it (RFC 5250 Section 3). The
// body of such an LSA is not decoded, so its TLVs are reported raw in the parent
// LSA's unknown_tlvs.
// Unknown returns a Ospfv2OpaqueLsaUnknownOpaqueType
func (obj *ospfv2OpaqueLsaTlvInformation) HasUnknown() bool {
	return obj.obj.Unknown != nil
}

// The Opaque Type when this model does not decode it (RFC 5250 Section 3). The
// body of such an LSA is not decoded, so its TLVs are reported raw in the parent
// LSA's unknown_tlvs.
// SetUnknown sets the Ospfv2OpaqueLsaUnknownOpaqueType value in the Ospfv2OpaqueLsaTlvInformation object
func (obj *ospfv2OpaqueLsaTlvInformation) SetUnknown(value Ospfv2OpaqueLsaUnknownOpaqueType) Ospfv2OpaqueLsaTlvInformation {
	obj.setChoice(Ospfv2OpaqueLsaTlvInformationChoice.UNKNOWN)
	obj.unknownHolder = nil
	obj.obj.Unknown = value.msg()

	return obj
}

func (obj *ospfv2OpaqueLsaTlvInformation) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TrafficEngineering != nil {

		obj.TrafficEngineering().validateObj(vObj, set_default)
	}

	if obj.obj.RouterInformation != nil {

		obj.RouterInformation().validateObj(vObj, set_default)
	}

	if len(obj.obj.ExtendedPrefixes) != 0 {

		if set_default {
			obj.ExtendedPrefixes().clearHolderSlice()
			for _, item := range obj.obj.ExtendedPrefixes {
				obj.ExtendedPrefixes().appendHolderSlice(&ospfv2OpaqueLsaExtendedPrefix{obj: item})
			}
		}
		for _, item := range obj.ExtendedPrefixes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.ExtendedLink != nil {

		obj.ExtendedLink().validateObj(vObj, set_default)
	}

	if obj.obj.Unknown != nil {

		obj.Unknown().validateObj(vObj, set_default)
	}

}

func (obj *ospfv2OpaqueLsaTlvInformation) setDefault() {
	var choices_set int = 0
	var choice Ospfv2OpaqueLsaTlvInformationChoiceEnum

	if obj.obj.TrafficEngineering != nil {
		choices_set += 1
		choice = Ospfv2OpaqueLsaTlvInformationChoice.TRAFFIC_ENGINEERING
	}

	if obj.obj.RouterInformation != nil {
		choices_set += 1
		choice = Ospfv2OpaqueLsaTlvInformationChoice.ROUTER_INFORMATION
	}

	if len(obj.obj.ExtendedPrefixes) > 0 {
		choices_set += 1
		choice = Ospfv2OpaqueLsaTlvInformationChoice.EXTENDED_PREFIXES
	}

	if obj.obj.ExtendedLink != nil {
		choices_set += 1
		choice = Ospfv2OpaqueLsaTlvInformationChoice.EXTENDED_LINK
	}

	if obj.obj.Unknown != nil {
		choices_set += 1
		choice = Ospfv2OpaqueLsaTlvInformationChoice.UNKNOWN
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv2OpaqueLsaTlvInformation")
			}
		} else {
			intVal := otg.Ospfv2OpaqueLsaTlvInformation_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv2OpaqueLsaTlvInformation_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
