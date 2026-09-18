package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2OpaqueLsa *****
type ospfv2OpaqueLsa struct {
	validation
	obj                      *otg.Ospfv2OpaqueLsa
	marshaller               marshalOspfv2OpaqueLsa
	unMarshaller             unMarshalOspfv2OpaqueLsa
	headerHolder             Ospfv2LsaHeader
	routerInformationHolder  Ospfv2OpaqueLsaRouterInformation
	trafficEngineeringHolder Ospfv2OpaqueLsaTrafficEngineering
	extendedPrefixesHolder   Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter
	extendedLinksHolder      Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter
	unknownTlvsHolder        Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter
}

func NewOspfv2OpaqueLsa() Ospfv2OpaqueLsa {
	obj := ospfv2OpaqueLsa{obj: &otg.Ospfv2OpaqueLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2OpaqueLsa) msg() *otg.Ospfv2OpaqueLsa {
	return obj.obj
}

func (obj *ospfv2OpaqueLsa) setMsg(msg *otg.Ospfv2OpaqueLsa) Ospfv2OpaqueLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2OpaqueLsa struct {
	obj *ospfv2OpaqueLsa
}

type marshalOspfv2OpaqueLsa interface {
	// ToProto marshals Ospfv2OpaqueLsa to protobuf object *otg.Ospfv2OpaqueLsa
	ToProto() (*otg.Ospfv2OpaqueLsa, error)
	// ToPbText marshals Ospfv2OpaqueLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2OpaqueLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2OpaqueLsa to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2OpaqueLsa struct {
	obj *ospfv2OpaqueLsa
}

type unMarshalOspfv2OpaqueLsa interface {
	// FromProto unmarshals Ospfv2OpaqueLsa from protobuf object *otg.Ospfv2OpaqueLsa
	FromProto(msg *otg.Ospfv2OpaqueLsa) (Ospfv2OpaqueLsa, error)
	// FromPbText unmarshals Ospfv2OpaqueLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2OpaqueLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2OpaqueLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv2OpaqueLsa) Marshal() marshalOspfv2OpaqueLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2OpaqueLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2OpaqueLsa) Unmarshal() unMarshalOspfv2OpaqueLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2OpaqueLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2OpaqueLsa) ToProto() (*otg.Ospfv2OpaqueLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2OpaqueLsa) FromProto(msg *otg.Ospfv2OpaqueLsa) (Ospfv2OpaqueLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2OpaqueLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsa) FromPbText(value string) error {
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

func (m *marshalospfv2OpaqueLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsa) FromYaml(value string) error {
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

func (m *marshalospfv2OpaqueLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsa) FromJson(value string) error {
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

func (obj *ospfv2OpaqueLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2OpaqueLsa) Clone() (Ospfv2OpaqueLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2OpaqueLsa()
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

func (obj *ospfv2OpaqueLsa) setNil() {
	obj.headerHolder = nil
	obj.routerInformationHolder = nil
	obj.trafficEngineeringHolder = nil
	obj.extendedPrefixesHolder = nil
	obj.extendedLinksHolder = nil
	obj.unknownTlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2OpaqueLsa is contents of OSPFv2 Opaque LSA - Type 9/10/11 (RFC 5250).
// The Link State ID of an Opaque LSA is not a plain IPv4 address; it is split into
// an Opaque Type (most significant octet) and an Opaque ID (remaining three octets),
// decoded here as tlv_information and id (RFC 5250 Section 3). header.lsa_id carries
// the raw, undecoded Link State ID value.
// The Segment Routing and Traffic Engineering information carried by an Opaque LSA is
// reported on this object, in the LSA it was actually advertised in, rather than on the
// Router-LSA, Summary-LSA, AS-External-LSA or NSSA-LSA that describes the router, link
// or prefix it applies to. Exactly one of router_information, traffic_engineering,
// extended_prefixes or extended_links is populated, selected by tlv_information; every
// other top-level TLV of the LSA is reported raw in unknown_tlvs. Each of those objects
// documents how to correlate it back to the LSA that advertises the router, link or
// prefix it describes.
type Ospfv2OpaqueLsa interface {
	Validation
	// msg marshals Ospfv2OpaqueLsa to protobuf object *otg.Ospfv2OpaqueLsa
	// and doesn't set defaults
	msg() *otg.Ospfv2OpaqueLsa
	// setMsg unmarshals Ospfv2OpaqueLsa from protobuf object *otg.Ospfv2OpaqueLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv2OpaqueLsa) Ospfv2OpaqueLsa
	// provides marshal interface
	Marshal() marshalOspfv2OpaqueLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2OpaqueLsa
	// validate validates Ospfv2OpaqueLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2OpaqueLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Header returns Ospfv2LsaHeader, set in Ospfv2OpaqueLsa.
	// Ospfv2LsaHeader is attributes in LSA Header.
	Header() Ospfv2LsaHeader
	// SetHeader assigns Ospfv2LsaHeader provided by user to Ospfv2OpaqueLsa.
	// Ospfv2LsaHeader is attributes in LSA Header.
	SetHeader(value Ospfv2LsaHeader) Ospfv2OpaqueLsa
	// HasHeader checks if Header has been set in Ospfv2OpaqueLsa
	HasHeader() bool
	// Type returns Ospfv2OpaqueLsaTypeEnum, set in Ospfv2OpaqueLsa
	Type() Ospfv2OpaqueLsaTypeEnum
	// SetType assigns Ospfv2OpaqueLsaTypeEnum provided by user to Ospfv2OpaqueLsa
	SetType(value Ospfv2OpaqueLsaTypeEnum) Ospfv2OpaqueLsa
	// HasType checks if Type has been set in Ospfv2OpaqueLsa
	HasType() bool
	// TlvInformation returns Ospfv2OpaqueLsaTlvInformationEnum, set in Ospfv2OpaqueLsa
	TlvInformation() Ospfv2OpaqueLsaTlvInformationEnum
	// SetTlvInformation assigns Ospfv2OpaqueLsaTlvInformationEnum provided by user to Ospfv2OpaqueLsa
	SetTlvInformation(value Ospfv2OpaqueLsaTlvInformationEnum) Ospfv2OpaqueLsa
	// HasTlvInformation checks if TlvInformation has been set in Ospfv2OpaqueLsa
	HasTlvInformation() bool
	// Id returns uint32, set in Ospfv2OpaqueLsa.
	Id() uint32
	// SetId assigns uint32 provided by user to Ospfv2OpaqueLsa
	SetId(value uint32) Ospfv2OpaqueLsa
	// HasId checks if Id has been set in Ospfv2OpaqueLsa
	HasId() bool
	// RouterInformation returns Ospfv2OpaqueLsaRouterInformation, set in Ospfv2OpaqueLsa.
	// Ospfv2OpaqueLsaRouterInformation is the decoded Router Information (RI) Opaque LSA contents, Opaque Type 4
	// (RFC 7770 Section 2). The RI Opaque LSA is originated once per router, so everything
	// reported here describes the router named by the parent LSA's
	// header.advertising_router_id.
	RouterInformation() Ospfv2OpaqueLsaRouterInformation
	// SetRouterInformation assigns Ospfv2OpaqueLsaRouterInformation provided by user to Ospfv2OpaqueLsa.
	// Ospfv2OpaqueLsaRouterInformation is the decoded Router Information (RI) Opaque LSA contents, Opaque Type 4
	// (RFC 7770 Section 2). The RI Opaque LSA is originated once per router, so everything
	// reported here describes the router named by the parent LSA's
	// header.advertising_router_id.
	SetRouterInformation(value Ospfv2OpaqueLsaRouterInformation) Ospfv2OpaqueLsa
	// HasRouterInformation checks if RouterInformation has been set in Ospfv2OpaqueLsa
	HasRouterInformation() bool
	// TrafficEngineering returns Ospfv2OpaqueLsaTrafficEngineering, set in Ospfv2OpaqueLsa.
	// Ospfv2OpaqueLsaTrafficEngineering is the decoded Traffic Engineering Opaque LSA contents, Opaque Type 1
	// (RFC 3630 Section 2.4). A TE LSA carries a single top-level TLV: either the Router
	// Address TLV, reported as router_address, or the Link TLV, reported as link_id plus
	// link_attributes.
	TrafficEngineering() Ospfv2OpaqueLsaTrafficEngineering
	// SetTrafficEngineering assigns Ospfv2OpaqueLsaTrafficEngineering provided by user to Ospfv2OpaqueLsa.
	// Ospfv2OpaqueLsaTrafficEngineering is the decoded Traffic Engineering Opaque LSA contents, Opaque Type 1
	// (RFC 3630 Section 2.4). A TE LSA carries a single top-level TLV: either the Router
	// Address TLV, reported as router_address, or the Link TLV, reported as link_id plus
	// link_attributes.
	SetTrafficEngineering(value Ospfv2OpaqueLsaTrafficEngineering) Ospfv2OpaqueLsa
	// HasTrafficEngineering checks if TrafficEngineering has been set in Ospfv2OpaqueLsa
	HasTrafficEngineering() bool
	// ExtendedPrefixes returns Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIterIter, set in Ospfv2OpaqueLsa
	ExtendedPrefixes() Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter
	// ExtendedLinks returns Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIterIter, set in Ospfv2OpaqueLsa
	ExtendedLinks() Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter
	// UnknownTlvs returns Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIterIter, set in Ospfv2OpaqueLsa
	UnknownTlvs() Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter
	setNil()
}

// Contents of the LSA header.
// Header returns a Ospfv2LsaHeader
func (obj *ospfv2OpaqueLsa) Header() Ospfv2LsaHeader {
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
func (obj *ospfv2OpaqueLsa) HasHeader() bool {
	return obj.obj.Header != nil
}

// Contents of the LSA header.
// SetHeader sets the Ospfv2LsaHeader value in the Ospfv2OpaqueLsa object
func (obj *ospfv2OpaqueLsa) SetHeader(value Ospfv2LsaHeader) Ospfv2OpaqueLsa {

	obj.headerHolder = nil
	obj.obj.Header = value.msg()

	return obj
}

type Ospfv2OpaqueLsaTypeEnum string

// Enum of Type on Ospfv2OpaqueLsa
var Ospfv2OpaqueLsaType = struct {
	LOCAL  Ospfv2OpaqueLsaTypeEnum
	AREA   Ospfv2OpaqueLsaTypeEnum
	DOMAIN Ospfv2OpaqueLsaTypeEnum
}{
	LOCAL:  Ospfv2OpaqueLsaTypeEnum("local"),
	AREA:   Ospfv2OpaqueLsaTypeEnum("area"),
	DOMAIN: Ospfv2OpaqueLsaTypeEnum("domain"),
}

func (obj *ospfv2OpaqueLsa) Type() Ospfv2OpaqueLsaTypeEnum {
	return Ospfv2OpaqueLsaTypeEnum(obj.obj.Type.Enum().String())
}

// The flooding scope of the Opaque LSA, determined by the LSA's LS Type
// (RFC 5250 Section 4): local (Type 9, not flooded beyond the local link),
// area (Type 10, flooded throughout the area) or domain (Type 11, flooded
// throughout the Autonomous System, excluding stub areas).
// Type returns a string
func (obj *ospfv2OpaqueLsa) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *ospfv2OpaqueLsa) SetType(value Ospfv2OpaqueLsaTypeEnum) Ospfv2OpaqueLsa {
	intValue, ok := otg.Ospfv2OpaqueLsa_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2OpaqueLsaTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2OpaqueLsa_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

type Ospfv2OpaqueLsaTlvInformationEnum string

// Enum of TlvInformation on Ospfv2OpaqueLsa
var Ospfv2OpaqueLsaTlvInformation = struct {
	TRAFFIC_ENGINEERING       Ospfv2OpaqueLsaTlvInformationEnum
	SYCAMORE_OPTICAL_TOPOLOGY Ospfv2OpaqueLsaTlvInformationEnum
	GRACE                     Ospfv2OpaqueLsaTlvInformationEnum
	ROUTER_INFORMATION        Ospfv2OpaqueLsaTlvInformationEnum
	L1VPN                     Ospfv2OpaqueLsaTlvInformationEnum
	INTER_AS_TE_V2            Ospfv2OpaqueLsaTlvInformationEnum
	EXTENDED_PREFIX           Ospfv2OpaqueLsaTlvInformationEnum
	EXTENDED_LINK             Ospfv2OpaqueLsaTlvInformationEnum
	TTZ                       Ospfv2OpaqueLsaTlvInformationEnum
	DYNAMIC_FLOODING          Ospfv2OpaqueLsaTlvInformationEnum
	EXTENDED_INTER_AREA_ASBR  Ospfv2OpaqueLsaTlvInformationEnum
}{
	TRAFFIC_ENGINEERING:       Ospfv2OpaqueLsaTlvInformationEnum("traffic_engineering"),
	SYCAMORE_OPTICAL_TOPOLOGY: Ospfv2OpaqueLsaTlvInformationEnum("sycamore_optical_topology"),
	GRACE:                     Ospfv2OpaqueLsaTlvInformationEnum("grace"),
	ROUTER_INFORMATION:        Ospfv2OpaqueLsaTlvInformationEnum("router_information"),
	L1VPN:                     Ospfv2OpaqueLsaTlvInformationEnum("l1vpn"),
	INTER_AS_TE_V2:            Ospfv2OpaqueLsaTlvInformationEnum("inter_as_te_v2"),
	EXTENDED_PREFIX:           Ospfv2OpaqueLsaTlvInformationEnum("extended_prefix"),
	EXTENDED_LINK:             Ospfv2OpaqueLsaTlvInformationEnum("extended_link"),
	TTZ:                       Ospfv2OpaqueLsaTlvInformationEnum("ttz"),
	DYNAMIC_FLOODING:          Ospfv2OpaqueLsaTlvInformationEnum("dynamic_flooding"),
	EXTENDED_INTER_AREA_ASBR:  Ospfv2OpaqueLsaTlvInformationEnum("extended_inter_area_asbr"),
}

func (obj *ospfv2OpaqueLsa) TlvInformation() Ospfv2OpaqueLsaTlvInformationEnum {
	return Ospfv2OpaqueLsaTlvInformationEnum(obj.obj.TlvInformation.Enum().String())
}

// The Opaque Type, decoded from the most significant octet of the LSA's Link
// State ID (RFC 5250 Section 3). Identifies the type of information carried in
// the tlvs (IANA Opaque LSA Option Types registry).
// TlvInformation returns a string
func (obj *ospfv2OpaqueLsa) HasTlvInformation() bool {
	return obj.obj.TlvInformation != nil
}

func (obj *ospfv2OpaqueLsa) SetTlvInformation(value Ospfv2OpaqueLsaTlvInformationEnum) Ospfv2OpaqueLsa {
	intValue, ok := otg.Ospfv2OpaqueLsa_TlvInformation_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2OpaqueLsaTlvInformationEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2OpaqueLsa_TlvInformation_Enum(intValue)
	obj.obj.TlvInformation = &enumValue

	return obj
}

// The Opaque ID, decoded from the least significant three octets of the LSA's Link State ID (RFC 5250 Section 3). Used to further distinguish LSAs of the same Opaque Type originated by the same router.
// Id returns a uint32
func (obj *ospfv2OpaqueLsa) Id() uint32 {

	return *obj.obj.Id

}

// The Opaque ID, decoded from the least significant three octets of the LSA's Link State ID (RFC 5250 Section 3). Used to further distinguish LSAs of the same Opaque Type originated by the same router.
// Id returns a uint32
func (obj *ospfv2OpaqueLsa) HasId() bool {
	return obj.obj.Id != nil
}

// The Opaque ID, decoded from the least significant three octets of the LSA's Link State ID (RFC 5250 Section 3). Used to further distinguish LSAs of the same Opaque Type originated by the same router.
// SetId sets the uint32 value in the Ospfv2OpaqueLsa object
func (obj *ospfv2OpaqueLsa) SetId(value uint32) Ospfv2OpaqueLsa {

	obj.obj.Id = &value
	return obj
}

// The decoded contents of a Router Information (RI) Opaque LSA, present when
// tlv_information is router_information (RFC 7770 Section 2).
// An RI Opaque LSA is originated per router, so its contents describe the router
// identified by header.advertising_router_id. Correlate it to the Router-LSA of that
// router by matching router_lsas[].header.advertising_router_id.
// RouterInformation returns a Ospfv2OpaqueLsaRouterInformation
func (obj *ospfv2OpaqueLsa) RouterInformation() Ospfv2OpaqueLsaRouterInformation {
	if obj.obj.RouterInformation == nil {
		obj.obj.RouterInformation = NewOspfv2OpaqueLsaRouterInformation().msg()
	}
	if obj.routerInformationHolder == nil {
		obj.routerInformationHolder = &ospfv2OpaqueLsaRouterInformation{obj: obj.obj.RouterInformation}
	}
	return obj.routerInformationHolder
}

// The decoded contents of a Router Information (RI) Opaque LSA, present when
// tlv_information is router_information (RFC 7770 Section 2).
// An RI Opaque LSA is originated per router, so its contents describe the router
// identified by header.advertising_router_id. Correlate it to the Router-LSA of that
// router by matching router_lsas[].header.advertising_router_id.
// RouterInformation returns a Ospfv2OpaqueLsaRouterInformation
func (obj *ospfv2OpaqueLsa) HasRouterInformation() bool {
	return obj.obj.RouterInformation != nil
}

// The decoded contents of a Router Information (RI) Opaque LSA, present when
// tlv_information is router_information (RFC 7770 Section 2).
// An RI Opaque LSA is originated per router, so its contents describe the router
// identified by header.advertising_router_id. Correlate it to the Router-LSA of that
// router by matching router_lsas[].header.advertising_router_id.
// SetRouterInformation sets the Ospfv2OpaqueLsaRouterInformation value in the Ospfv2OpaqueLsa object
func (obj *ospfv2OpaqueLsa) SetRouterInformation(value Ospfv2OpaqueLsaRouterInformation) Ospfv2OpaqueLsa {

	obj.routerInformationHolder = nil
	obj.obj.RouterInformation = value.msg()

	return obj
}

// The decoded contents of a Traffic Engineering Opaque LSA, present when
// tlv_information is traffic_engineering (RFC 3630 Section 2.4).
// TrafficEngineering returns a Ospfv2OpaqueLsaTrafficEngineering
func (obj *ospfv2OpaqueLsa) TrafficEngineering() Ospfv2OpaqueLsaTrafficEngineering {
	if obj.obj.TrafficEngineering == nil {
		obj.obj.TrafficEngineering = NewOspfv2OpaqueLsaTrafficEngineering().msg()
	}
	if obj.trafficEngineeringHolder == nil {
		obj.trafficEngineeringHolder = &ospfv2OpaqueLsaTrafficEngineering{obj: obj.obj.TrafficEngineering}
	}
	return obj.trafficEngineeringHolder
}

// The decoded contents of a Traffic Engineering Opaque LSA, present when
// tlv_information is traffic_engineering (RFC 3630 Section 2.4).
// TrafficEngineering returns a Ospfv2OpaqueLsaTrafficEngineering
func (obj *ospfv2OpaqueLsa) HasTrafficEngineering() bool {
	return obj.obj.TrafficEngineering != nil
}

// The decoded contents of a Traffic Engineering Opaque LSA, present when
// tlv_information is traffic_engineering (RFC 3630 Section 2.4).
// SetTrafficEngineering sets the Ospfv2OpaqueLsaTrafficEngineering value in the Ospfv2OpaqueLsa object
func (obj *ospfv2OpaqueLsa) SetTrafficEngineering(value Ospfv2OpaqueLsaTrafficEngineering) Ospfv2OpaqueLsa {

	obj.trafficEngineeringHolder = nil
	obj.obj.TrafficEngineering = value.msg()

	return obj
}

// The decoded OSPFv2 Extended Prefix TLVs of an Extended Prefix Opaque LSA, present
// when tlv_information is extended_prefix. One Extended Prefix Opaque LSA can carry
// more than one Extended Prefix TLV, each describing a different prefix
// (RFC 7684 Section 2.1).
// ExtendedPrefixes returns a []Ospfv2OpaqueLsaExtendedPrefix
func (obj *ospfv2OpaqueLsa) ExtendedPrefixes() Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter {
	if len(obj.obj.ExtendedPrefixes) == 0 {
		obj.obj.ExtendedPrefixes = []*otg.Ospfv2OpaqueLsaExtendedPrefix{}
	}
	if obj.extendedPrefixesHolder == nil {
		obj.extendedPrefixesHolder = newOspfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter(&obj.obj.ExtendedPrefixes).setMsg(obj)
	}
	return obj.extendedPrefixesHolder
}

type ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter struct {
	obj                                *ospfv2OpaqueLsa
	ospfv2OpaqueLsaExtendedPrefixSlice []Ospfv2OpaqueLsaExtendedPrefix
	fieldPtr                           *[]*otg.Ospfv2OpaqueLsaExtendedPrefix
}

func newOspfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter(ptr *[]*otg.Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter {
	return &ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter{fieldPtr: ptr}
}

type Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter interface {
	setMsg(*ospfv2OpaqueLsa) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter
	Items() []Ospfv2OpaqueLsaExtendedPrefix
	Add() Ospfv2OpaqueLsaExtendedPrefix
	Append(items ...Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter
	Set(index int, newObj Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter
	Clear() Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter
	clearHolderSlice() Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter
	appendHolderSlice(item Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter) setMsg(msg *ospfv2OpaqueLsa) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2OpaqueLsaExtendedPrefix{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter) Items() []Ospfv2OpaqueLsaExtendedPrefix {
	return obj.ospfv2OpaqueLsaExtendedPrefixSlice
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter) Add() Ospfv2OpaqueLsaExtendedPrefix {
	newObj := &otg.Ospfv2OpaqueLsaExtendedPrefix{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2OpaqueLsaExtendedPrefix{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2OpaqueLsaExtendedPrefixSlice = append(obj.ospfv2OpaqueLsaExtendedPrefixSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter) Append(items ...Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2OpaqueLsaExtendedPrefixSlice = append(obj.ospfv2OpaqueLsaExtendedPrefixSlice, item)
	}
	return obj
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter) Set(index int, newObj Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2OpaqueLsaExtendedPrefixSlice[index] = newObj
	return obj
}
func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter) Clear() Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2OpaqueLsaExtendedPrefix{}
		obj.ospfv2OpaqueLsaExtendedPrefixSlice = []Ospfv2OpaqueLsaExtendedPrefix{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter) clearHolderSlice() Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter {
	if len(obj.ospfv2OpaqueLsaExtendedPrefixSlice) > 0 {
		obj.ospfv2OpaqueLsaExtendedPrefixSlice = []Ospfv2OpaqueLsaExtendedPrefix{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter) appendHolderSlice(item Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedPrefixIter {
	obj.ospfv2OpaqueLsaExtendedPrefixSlice = append(obj.ospfv2OpaqueLsaExtendedPrefixSlice, item)
	return obj
}

// The decoded OSPFv2 Extended Link TLVs of an Extended Link Opaque LSA, present when
// tlv_information is extended_link (RFC 7684 Section 3.1).
// ExtendedLinks returns a []Ospfv2OpaqueLsaExtendedLink
func (obj *ospfv2OpaqueLsa) ExtendedLinks() Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter {
	if len(obj.obj.ExtendedLinks) == 0 {
		obj.obj.ExtendedLinks = []*otg.Ospfv2OpaqueLsaExtendedLink{}
	}
	if obj.extendedLinksHolder == nil {
		obj.extendedLinksHolder = newOspfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter(&obj.obj.ExtendedLinks).setMsg(obj)
	}
	return obj.extendedLinksHolder
}

type ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter struct {
	obj                              *ospfv2OpaqueLsa
	ospfv2OpaqueLsaExtendedLinkSlice []Ospfv2OpaqueLsaExtendedLink
	fieldPtr                         *[]*otg.Ospfv2OpaqueLsaExtendedLink
}

func newOspfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter(ptr *[]*otg.Ospfv2OpaqueLsaExtendedLink) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter {
	return &ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter{fieldPtr: ptr}
}

type Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter interface {
	setMsg(*ospfv2OpaqueLsa) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter
	Items() []Ospfv2OpaqueLsaExtendedLink
	Add() Ospfv2OpaqueLsaExtendedLink
	Append(items ...Ospfv2OpaqueLsaExtendedLink) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter
	Set(index int, newObj Ospfv2OpaqueLsaExtendedLink) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter
	Clear() Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter
	clearHolderSlice() Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter
	appendHolderSlice(item Ospfv2OpaqueLsaExtendedLink) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter) setMsg(msg *ospfv2OpaqueLsa) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2OpaqueLsaExtendedLink{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter) Items() []Ospfv2OpaqueLsaExtendedLink {
	return obj.ospfv2OpaqueLsaExtendedLinkSlice
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter) Add() Ospfv2OpaqueLsaExtendedLink {
	newObj := &otg.Ospfv2OpaqueLsaExtendedLink{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2OpaqueLsaExtendedLink{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2OpaqueLsaExtendedLinkSlice = append(obj.ospfv2OpaqueLsaExtendedLinkSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter) Append(items ...Ospfv2OpaqueLsaExtendedLink) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2OpaqueLsaExtendedLinkSlice = append(obj.ospfv2OpaqueLsaExtendedLinkSlice, item)
	}
	return obj
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter) Set(index int, newObj Ospfv2OpaqueLsaExtendedLink) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2OpaqueLsaExtendedLinkSlice[index] = newObj
	return obj
}
func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter) Clear() Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2OpaqueLsaExtendedLink{}
		obj.ospfv2OpaqueLsaExtendedLinkSlice = []Ospfv2OpaqueLsaExtendedLink{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter) clearHolderSlice() Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter {
	if len(obj.ospfv2OpaqueLsaExtendedLinkSlice) > 0 {
		obj.ospfv2OpaqueLsaExtendedLinkSlice = []Ospfv2OpaqueLsaExtendedLink{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter) appendHolderSlice(item Ospfv2OpaqueLsaExtendedLink) Ospfv2OpaqueLsaOspfv2OpaqueLsaExtendedLinkIter {
	obj.ospfv2OpaqueLsaExtendedLinkSlice = append(obj.ospfv2OpaqueLsaExtendedLinkSlice, item)
	return obj
}

// TLVs carried in the body of the Opaque LSA that are not decoded into
// router_information, traffic_engineering, extended_prefixes or extended_links,
// returned raw in the generic type/length/value TLV format used by all OSPFv2 Opaque
// LSAs (RFC 7770 Section 2, RFC 8665, RFC 9492).
// UnknownTlvs returns a []Ospfv2OpaqueLsaTlv
func (obj *ospfv2OpaqueLsa) UnknownTlvs() Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter {
	if len(obj.obj.UnknownTlvs) == 0 {
		obj.obj.UnknownTlvs = []*otg.Ospfv2OpaqueLsaTlv{}
	}
	if obj.unknownTlvsHolder == nil {
		obj.unknownTlvsHolder = newOspfv2OpaqueLsaOspfv2OpaqueLsaTlvIter(&obj.obj.UnknownTlvs).setMsg(obj)
	}
	return obj.unknownTlvsHolder
}

type ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter struct {
	obj                     *ospfv2OpaqueLsa
	ospfv2OpaqueLsaTlvSlice []Ospfv2OpaqueLsaTlv
	fieldPtr                *[]*otg.Ospfv2OpaqueLsaTlv
}

func newOspfv2OpaqueLsaOspfv2OpaqueLsaTlvIter(ptr *[]*otg.Ospfv2OpaqueLsaTlv) Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter {
	return &ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter{fieldPtr: ptr}
}

type Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter interface {
	setMsg(*ospfv2OpaqueLsa) Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter
	Items() []Ospfv2OpaqueLsaTlv
	Add() Ospfv2OpaqueLsaTlv
	Append(items ...Ospfv2OpaqueLsaTlv) Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter
	Set(index int, newObj Ospfv2OpaqueLsaTlv) Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter
	Clear() Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter
	clearHolderSlice() Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter
	appendHolderSlice(item Ospfv2OpaqueLsaTlv) Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter) setMsg(msg *ospfv2OpaqueLsa) Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2OpaqueLsaTlv{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter) Items() []Ospfv2OpaqueLsaTlv {
	return obj.ospfv2OpaqueLsaTlvSlice
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter) Add() Ospfv2OpaqueLsaTlv {
	newObj := &otg.Ospfv2OpaqueLsaTlv{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2OpaqueLsaTlv{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2OpaqueLsaTlvSlice = append(obj.ospfv2OpaqueLsaTlvSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter) Append(items ...Ospfv2OpaqueLsaTlv) Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2OpaqueLsaTlvSlice = append(obj.ospfv2OpaqueLsaTlvSlice, item)
	}
	return obj
}

func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter) Set(index int, newObj Ospfv2OpaqueLsaTlv) Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2OpaqueLsaTlvSlice[index] = newObj
	return obj
}
func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter) Clear() Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2OpaqueLsaTlv{}
		obj.ospfv2OpaqueLsaTlvSlice = []Ospfv2OpaqueLsaTlv{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter) clearHolderSlice() Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter {
	if len(obj.ospfv2OpaqueLsaTlvSlice) > 0 {
		obj.ospfv2OpaqueLsaTlvSlice = []Ospfv2OpaqueLsaTlv{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter) appendHolderSlice(item Ospfv2OpaqueLsaTlv) Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter {
	obj.ospfv2OpaqueLsaTlvSlice = append(obj.ospfv2OpaqueLsaTlvSlice, item)
	return obj
}

func (obj *ospfv2OpaqueLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Header != nil {

		obj.Header().validateObj(vObj, set_default)
	}

	if obj.obj.Id != nil {

		if *obj.obj.Id > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv2OpaqueLsa.Id <= 16777215 but Got %d", *obj.obj.Id))
		}

	}

	if obj.obj.RouterInformation != nil {

		obj.RouterInformation().validateObj(vObj, set_default)
	}

	if obj.obj.TrafficEngineering != nil {

		obj.TrafficEngineering().validateObj(vObj, set_default)
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

	if len(obj.obj.ExtendedLinks) != 0 {

		if set_default {
			obj.ExtendedLinks().clearHolderSlice()
			for _, item := range obj.obj.ExtendedLinks {
				obj.ExtendedLinks().appendHolderSlice(&ospfv2OpaqueLsaExtendedLink{obj: item})
			}
		}
		for _, item := range obj.ExtendedLinks().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.UnknownTlvs) != 0 {

		if set_default {
			obj.UnknownTlvs().clearHolderSlice()
			for _, item := range obj.obj.UnknownTlvs {
				obj.UnknownTlvs().appendHolderSlice(&ospfv2OpaqueLsaTlv{obj: item})
			}
		}
		for _, item := range obj.UnknownTlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv2OpaqueLsa) setDefault() {

}
