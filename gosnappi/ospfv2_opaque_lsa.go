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
	obj          *otg.Ospfv2OpaqueLsa
	marshaller   marshalOspfv2OpaqueLsa
	unMarshaller unMarshalOspfv2OpaqueLsa
	headerHolder Ospfv2LsaHeader
	tlvsHolder   Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter
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
	obj.tlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2OpaqueLsa is contents of OSPFv2 Opaque LSA - Type 9/10/11 (RFC 5250).
// The Link State ID of an Opaque LSA is not a plain IPv4 address; it is split into
// an Opaque Type (most significant octet) and an Opaque ID (remaining three octets),
// decoded here as tlv_information and id (RFC 5250 Section 3). header.lsa_id carries
// the raw, undecoded Link State ID value.
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
	// Tlvs returns Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIterIter, set in Ospfv2OpaqueLsa
	Tlvs() Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter
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

// The raw, undecoded TLVs carried in the body of the Opaque LSA, in the generic
// type/length/value TLV format used by all OSPFv2 Opaque LSAs (RFC 7770 Section 2,
// RFC 8665, RFC 9492).
// Tlvs returns a []Ospfv2OpaqueLsaTlv
func (obj *ospfv2OpaqueLsa) Tlvs() Ospfv2OpaqueLsaOspfv2OpaqueLsaTlvIter {
	if len(obj.obj.Tlvs) == 0 {
		obj.obj.Tlvs = []*otg.Ospfv2OpaqueLsaTlv{}
	}
	if obj.tlvsHolder == nil {
		obj.tlvsHolder = newOspfv2OpaqueLsaOspfv2OpaqueLsaTlvIter(&obj.obj.Tlvs).setMsg(obj)
	}
	return obj.tlvsHolder
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

	if len(obj.obj.Tlvs) != 0 {

		if set_default {
			obj.Tlvs().clearHolderSlice()
			for _, item := range obj.obj.Tlvs {
				obj.Tlvs().appendHolderSlice(&ospfv2OpaqueLsaTlv{obj: item})
			}
		}
		for _, item := range obj.Tlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv2OpaqueLsa) setDefault() {

}
