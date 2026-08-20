package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2OpaqueLsaTlv *****
type ospfv2OpaqueLsaTlv struct {
	validation
	obj           *otg.Ospfv2OpaqueLsaTlv
	marshaller    marshalOspfv2OpaqueLsaTlv
	unMarshaller  unMarshalOspfv2OpaqueLsaTlv
	subTlvsHolder Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter
}

func NewOspfv2OpaqueLsaTlv() Ospfv2OpaqueLsaTlv {
	obj := ospfv2OpaqueLsaTlv{obj: &otg.Ospfv2OpaqueLsaTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2OpaqueLsaTlv) msg() *otg.Ospfv2OpaqueLsaTlv {
	return obj.obj
}

func (obj *ospfv2OpaqueLsaTlv) setMsg(msg *otg.Ospfv2OpaqueLsaTlv) Ospfv2OpaqueLsaTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2OpaqueLsaTlv struct {
	obj *ospfv2OpaqueLsaTlv
}

type marshalOspfv2OpaqueLsaTlv interface {
	// ToProto marshals Ospfv2OpaqueLsaTlv to protobuf object *otg.Ospfv2OpaqueLsaTlv
	ToProto() (*otg.Ospfv2OpaqueLsaTlv, error)
	// ToPbText marshals Ospfv2OpaqueLsaTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2OpaqueLsaTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2OpaqueLsaTlv to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2OpaqueLsaTlv struct {
	obj *ospfv2OpaqueLsaTlv
}

type unMarshalOspfv2OpaqueLsaTlv interface {
	// FromProto unmarshals Ospfv2OpaqueLsaTlv from protobuf object *otg.Ospfv2OpaqueLsaTlv
	FromProto(msg *otg.Ospfv2OpaqueLsaTlv) (Ospfv2OpaqueLsaTlv, error)
	// FromPbText unmarshals Ospfv2OpaqueLsaTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2OpaqueLsaTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2OpaqueLsaTlv from JSON text
	FromJson(value string) error
}

func (obj *ospfv2OpaqueLsaTlv) Marshal() marshalOspfv2OpaqueLsaTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2OpaqueLsaTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2OpaqueLsaTlv) Unmarshal() unMarshalOspfv2OpaqueLsaTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2OpaqueLsaTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2OpaqueLsaTlv) ToProto() (*otg.Ospfv2OpaqueLsaTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2OpaqueLsaTlv) FromProto(msg *otg.Ospfv2OpaqueLsaTlv) (Ospfv2OpaqueLsaTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2OpaqueLsaTlv) ToPbText() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaTlv) FromPbText(value string) error {
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

func (m *marshalospfv2OpaqueLsaTlv) ToYaml() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaTlv) FromYaml(value string) error {
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

func (m *marshalospfv2OpaqueLsaTlv) ToJson() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaTlv) FromJson(value string) error {
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

func (obj *ospfv2OpaqueLsaTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2OpaqueLsaTlv) Clone() (Ospfv2OpaqueLsaTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2OpaqueLsaTlv()
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

func (obj *ospfv2OpaqueLsaTlv) setNil() {
	obj.subTlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2OpaqueLsaTlv is a top-level TLV carried in the body of an OSPFv2 Opaque LSA (RFC 7770 Section 2).
type Ospfv2OpaqueLsaTlv interface {
	Validation
	// msg marshals Ospfv2OpaqueLsaTlv to protobuf object *otg.Ospfv2OpaqueLsaTlv
	// and doesn't set defaults
	msg() *otg.Ospfv2OpaqueLsaTlv
	// setMsg unmarshals Ospfv2OpaqueLsaTlv from protobuf object *otg.Ospfv2OpaqueLsaTlv
	// and doesn't set defaults
	setMsg(*otg.Ospfv2OpaqueLsaTlv) Ospfv2OpaqueLsaTlv
	// provides marshal interface
	Marshal() marshalOspfv2OpaqueLsaTlv
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2OpaqueLsaTlv
	// validate validates Ospfv2OpaqueLsaTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2OpaqueLsaTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns Ospfv2OpaqueLsaTlvTypeEnum, set in Ospfv2OpaqueLsaTlv
	Type() Ospfv2OpaqueLsaTlvTypeEnum
	// SetType assigns Ospfv2OpaqueLsaTlvTypeEnum provided by user to Ospfv2OpaqueLsaTlv
	SetType(value Ospfv2OpaqueLsaTlvTypeEnum) Ospfv2OpaqueLsaTlv
	// HasType checks if Type has been set in Ospfv2OpaqueLsaTlv
	HasType() bool
	// Length returns uint32, set in Ospfv2OpaqueLsaTlv.
	Length() uint32
	// SetLength assigns uint32 provided by user to Ospfv2OpaqueLsaTlv
	SetLength(value uint32) Ospfv2OpaqueLsaTlv
	// HasLength checks if Length has been set in Ospfv2OpaqueLsaTlv
	HasLength() bool
	// Value returns string, set in Ospfv2OpaqueLsaTlv.
	Value() string
	// SetValue assigns string provided by user to Ospfv2OpaqueLsaTlv
	SetValue(value string) Ospfv2OpaqueLsaTlv
	// HasValue checks if Value has been set in Ospfv2OpaqueLsaTlv
	HasValue() bool
	// SubTlvs returns Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIterIter, set in Ospfv2OpaqueLsaTlv
	SubTlvs() Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter
	setNil()
}

type Ospfv2OpaqueLsaTlvTypeEnum string

// Enum of Type on Ospfv2OpaqueLsaTlv
var Ospfv2OpaqueLsaTlvType = struct {
	TE_ROUTER_ADDRESS                Ospfv2OpaqueLsaTlvTypeEnum
	TE_LINK                          Ospfv2OpaqueLsaTlvTypeEnum
	TE_ROUTER_IPV6_ADDRESS           Ospfv2OpaqueLsaTlvTypeEnum
	TE_LINK_LOCAL                    Ospfv2OpaqueLsaTlvTypeEnum
	TE_NODE_ATTRIBUTE                Ospfv2OpaqueLsaTlvTypeEnum
	TE_OPTICAL_NODE_PROPERTY         Ospfv2OpaqueLsaTlvTypeEnum
	RI_INFORMATIONAL_CAPABILITIES    Ospfv2OpaqueLsaTlvTypeEnum
	RI_FUNCTIONAL_CAPABILITIES       Ospfv2OpaqueLsaTlvTypeEnum
	RI_TE_MESH_GROUP_IPV4            Ospfv2OpaqueLsaTlvTypeEnum
	RI_TE_MESH_GROUP_IPV6            Ospfv2OpaqueLsaTlvTypeEnum
	RI_TE_NODE_CAPABILITY_DESCRIPTOR Ospfv2OpaqueLsaTlvTypeEnum
	RI_PCED                          Ospfv2OpaqueLsaTlvTypeEnum
	RI_DYNAMIC_HOSTNAME              Ospfv2OpaqueLsaTlvTypeEnum
	RI_SR_ALGORITHM                  Ospfv2OpaqueLsaTlvTypeEnum
	RI_SID_LABEL_RANGE               Ospfv2OpaqueLsaTlvTypeEnum
	RI_NODE_ADMIN_TAG                Ospfv2OpaqueLsaTlvTypeEnum
	RI_SBFD_DISCRIMINATOR            Ospfv2OpaqueLsaTlvTypeEnum
	RI_NODE_MSD                      Ospfv2OpaqueLsaTlvTypeEnum
	RI_TUNNEL_ENCAPSULATIONS         Ospfv2OpaqueLsaTlvTypeEnum
	RI_SR_LOCAL_BLOCK                Ospfv2OpaqueLsaTlvTypeEnum
	RI_SRMS_PREFERENCE               Ospfv2OpaqueLsaTlvTypeEnum
	RI_FLEXIBLE_ALGORITHM_DEFINITION Ospfv2OpaqueLsaTlvTypeEnum
	RI_AREA_LEADER                   Ospfv2OpaqueLsaTlvTypeEnum
	RI_DYNAMIC_FLOODING              Ospfv2OpaqueLsaTlvTypeEnum
	RI_SRV6_CAPABILITIES             Ospfv2OpaqueLsaTlvTypeEnum
	RI_IP_ALGORITHM                  Ospfv2OpaqueLsaTlvTypeEnum
	EXTENDED_PREFIX_TLV              Ospfv2OpaqueLsaTlvTypeEnum
	EXTENDED_PREFIX_RANGE_TLV        Ospfv2OpaqueLsaTlvTypeEnum
	EXTENDED_LINK_TLV                Ospfv2OpaqueLsaTlvTypeEnum
	TTZ_ID                           Ospfv2OpaqueLsaTlvTypeEnum
	TTZ_ROUTER                       Ospfv2OpaqueLsaTlvTypeEnum
	TTZ_OPTIONS                      Ospfv2OpaqueLsaTlvTypeEnum
	DYNAMIC_FLOODING_AREA_ROUTER_IDS Ospfv2OpaqueLsaTlvTypeEnum
	DYNAMIC_FLOODING_PATH            Ospfv2OpaqueLsaTlvTypeEnum
	EXTENDED_INTER_AREA_ASBR_TLV     Ospfv2OpaqueLsaTlvTypeEnum
}{
	TE_ROUTER_ADDRESS:                Ospfv2OpaqueLsaTlvTypeEnum("te_router_address"),
	TE_LINK:                          Ospfv2OpaqueLsaTlvTypeEnum("te_link"),
	TE_ROUTER_IPV6_ADDRESS:           Ospfv2OpaqueLsaTlvTypeEnum("te_router_ipv6_address"),
	TE_LINK_LOCAL:                    Ospfv2OpaqueLsaTlvTypeEnum("te_link_local"),
	TE_NODE_ATTRIBUTE:                Ospfv2OpaqueLsaTlvTypeEnum("te_node_attribute"),
	TE_OPTICAL_NODE_PROPERTY:         Ospfv2OpaqueLsaTlvTypeEnum("te_optical_node_property"),
	RI_INFORMATIONAL_CAPABILITIES:    Ospfv2OpaqueLsaTlvTypeEnum("ri_informational_capabilities"),
	RI_FUNCTIONAL_CAPABILITIES:       Ospfv2OpaqueLsaTlvTypeEnum("ri_functional_capabilities"),
	RI_TE_MESH_GROUP_IPV4:            Ospfv2OpaqueLsaTlvTypeEnum("ri_te_mesh_group_ipv4"),
	RI_TE_MESH_GROUP_IPV6:            Ospfv2OpaqueLsaTlvTypeEnum("ri_te_mesh_group_ipv6"),
	RI_TE_NODE_CAPABILITY_DESCRIPTOR: Ospfv2OpaqueLsaTlvTypeEnum("ri_te_node_capability_descriptor"),
	RI_PCED:                          Ospfv2OpaqueLsaTlvTypeEnum("ri_pced"),
	RI_DYNAMIC_HOSTNAME:              Ospfv2OpaqueLsaTlvTypeEnum("ri_dynamic_hostname"),
	RI_SR_ALGORITHM:                  Ospfv2OpaqueLsaTlvTypeEnum("ri_sr_algorithm"),
	RI_SID_LABEL_RANGE:               Ospfv2OpaqueLsaTlvTypeEnum("ri_sid_label_range"),
	RI_NODE_ADMIN_TAG:                Ospfv2OpaqueLsaTlvTypeEnum("ri_node_admin_tag"),
	RI_SBFD_DISCRIMINATOR:            Ospfv2OpaqueLsaTlvTypeEnum("ri_sbfd_discriminator"),
	RI_NODE_MSD:                      Ospfv2OpaqueLsaTlvTypeEnum("ri_node_msd"),
	RI_TUNNEL_ENCAPSULATIONS:         Ospfv2OpaqueLsaTlvTypeEnum("ri_tunnel_encapsulations"),
	RI_SR_LOCAL_BLOCK:                Ospfv2OpaqueLsaTlvTypeEnum("ri_sr_local_block"),
	RI_SRMS_PREFERENCE:               Ospfv2OpaqueLsaTlvTypeEnum("ri_srms_preference"),
	RI_FLEXIBLE_ALGORITHM_DEFINITION: Ospfv2OpaqueLsaTlvTypeEnum("ri_flexible_algorithm_definition"),
	RI_AREA_LEADER:                   Ospfv2OpaqueLsaTlvTypeEnum("ri_area_leader"),
	RI_DYNAMIC_FLOODING:              Ospfv2OpaqueLsaTlvTypeEnum("ri_dynamic_flooding"),
	RI_SRV6_CAPABILITIES:             Ospfv2OpaqueLsaTlvTypeEnum("ri_srv6_capabilities"),
	RI_IP_ALGORITHM:                  Ospfv2OpaqueLsaTlvTypeEnum("ri_ip_algorithm"),
	EXTENDED_PREFIX_TLV:              Ospfv2OpaqueLsaTlvTypeEnum("extended_prefix_tlv"),
	EXTENDED_PREFIX_RANGE_TLV:        Ospfv2OpaqueLsaTlvTypeEnum("extended_prefix_range_tlv"),
	EXTENDED_LINK_TLV:                Ospfv2OpaqueLsaTlvTypeEnum("extended_link_tlv"),
	TTZ_ID:                           Ospfv2OpaqueLsaTlvTypeEnum("ttz_id"),
	TTZ_ROUTER:                       Ospfv2OpaqueLsaTlvTypeEnum("ttz_router"),
	TTZ_OPTIONS:                      Ospfv2OpaqueLsaTlvTypeEnum("ttz_options"),
	DYNAMIC_FLOODING_AREA_ROUTER_IDS: Ospfv2OpaqueLsaTlvTypeEnum("dynamic_flooding_area_router_ids"),
	DYNAMIC_FLOODING_PATH:            Ospfv2OpaqueLsaTlvTypeEnum("dynamic_flooding_path"),
	EXTENDED_INTER_AREA_ASBR_TLV:     Ospfv2OpaqueLsaTlvTypeEnum("extended_inter_area_asbr_tlv"),
}

func (obj *ospfv2OpaqueLsaTlv) Type() Ospfv2OpaqueLsaTlvTypeEnum {
	return Ospfv2OpaqueLsaTlvTypeEnum(obj.obj.Type.Enum().String())
}

// The TLV Type field. Its meaning is scoped by the parent LSA's opaque_type
// (IANA OSPFv2 TLV registries).
// Type returns a string
func (obj *ospfv2OpaqueLsaTlv) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *ospfv2OpaqueLsaTlv) SetType(value Ospfv2OpaqueLsaTlvTypeEnum) Ospfv2OpaqueLsaTlv {
	intValue, ok := otg.Ospfv2OpaqueLsaTlv_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2OpaqueLsaTlvTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2OpaqueLsaTlv_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// The TLV Length field, in octets, of the value field.
// Length returns a uint32
func (obj *ospfv2OpaqueLsaTlv) Length() uint32 {

	return *obj.obj.Length

}

// The TLV Length field, in octets, of the value field.
// Length returns a uint32
func (obj *ospfv2OpaqueLsaTlv) HasLength() bool {
	return obj.obj.Length != nil
}

// The TLV Length field, in octets, of the value field.
// SetLength sets the uint32 value in the Ospfv2OpaqueLsaTlv object
func (obj *ospfv2OpaqueLsaTlv) SetLength(value uint32) Ospfv2OpaqueLsaTlv {

	obj.obj.Length = &value
	return obj
}

// The TLV Value field, returned as a lowercase hexadecimal string.
// Value returns a string
func (obj *ospfv2OpaqueLsaTlv) Value() string {

	return *obj.obj.Value

}

// The TLV Value field, returned as a lowercase hexadecimal string.
// Value returns a string
func (obj *ospfv2OpaqueLsaTlv) HasValue() bool {
	return obj.obj.Value != nil
}

// The TLV Value field, returned as a lowercase hexadecimal string.
// SetValue sets the string value in the Ospfv2OpaqueLsaTlv object
func (obj *ospfv2OpaqueLsaTlv) SetValue(value string) Ospfv2OpaqueLsaTlv {

	obj.obj.Value = &value
	return obj
}

// The sub-TLVs nested within this TLV's value, in the generic type/length/value format (e.g. RFC 8665 Extended Prefix/Link Opaque LSA sub-TLVs).
// SubTlvs returns a []Ospfv2OpaqueLsaSubTlv
func (obj *ospfv2OpaqueLsaTlv) SubTlvs() Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter {
	if len(obj.obj.SubTlvs) == 0 {
		obj.obj.SubTlvs = []*otg.Ospfv2OpaqueLsaSubTlv{}
	}
	if obj.subTlvsHolder == nil {
		obj.subTlvsHolder = newOspfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter(&obj.obj.SubTlvs).setMsg(obj)
	}
	return obj.subTlvsHolder
}

type ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter struct {
	obj                        *ospfv2OpaqueLsaTlv
	ospfv2OpaqueLsaSubTlvSlice []Ospfv2OpaqueLsaSubTlv
	fieldPtr                   *[]*otg.Ospfv2OpaqueLsaSubTlv
}

func newOspfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter(ptr *[]*otg.Ospfv2OpaqueLsaSubTlv) Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter {
	return &ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter{fieldPtr: ptr}
}

type Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter interface {
	setMsg(*ospfv2OpaqueLsaTlv) Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter
	Items() []Ospfv2OpaqueLsaSubTlv
	Add() Ospfv2OpaqueLsaSubTlv
	Append(items ...Ospfv2OpaqueLsaSubTlv) Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter
	Set(index int, newObj Ospfv2OpaqueLsaSubTlv) Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter
	Clear() Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter
	clearHolderSlice() Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter
	appendHolderSlice(item Ospfv2OpaqueLsaSubTlv) Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter
}

func (obj *ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter) setMsg(msg *ospfv2OpaqueLsaTlv) Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2OpaqueLsaSubTlv{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter) Items() []Ospfv2OpaqueLsaSubTlv {
	return obj.ospfv2OpaqueLsaSubTlvSlice
}

func (obj *ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter) Add() Ospfv2OpaqueLsaSubTlv {
	newObj := &otg.Ospfv2OpaqueLsaSubTlv{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2OpaqueLsaSubTlv{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2OpaqueLsaSubTlvSlice = append(obj.ospfv2OpaqueLsaSubTlvSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter) Append(items ...Ospfv2OpaqueLsaSubTlv) Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2OpaqueLsaSubTlvSlice = append(obj.ospfv2OpaqueLsaSubTlvSlice, item)
	}
	return obj
}

func (obj *ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter) Set(index int, newObj Ospfv2OpaqueLsaSubTlv) Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2OpaqueLsaSubTlvSlice[index] = newObj
	return obj
}
func (obj *ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter) Clear() Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2OpaqueLsaSubTlv{}
		obj.ospfv2OpaqueLsaSubTlvSlice = []Ospfv2OpaqueLsaSubTlv{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter) clearHolderSlice() Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter {
	if len(obj.ospfv2OpaqueLsaSubTlvSlice) > 0 {
		obj.ospfv2OpaqueLsaSubTlvSlice = []Ospfv2OpaqueLsaSubTlv{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter) appendHolderSlice(item Ospfv2OpaqueLsaSubTlv) Ospfv2OpaqueLsaTlvOspfv2OpaqueLsaSubTlvIter {
	obj.ospfv2OpaqueLsaSubTlvSlice = append(obj.ospfv2OpaqueLsaSubTlvSlice, item)
	return obj
}

func (obj *ospfv2OpaqueLsaTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.SubTlvs) != 0 {

		if set_default {
			obj.SubTlvs().clearHolderSlice()
			for _, item := range obj.obj.SubTlvs {
				obj.SubTlvs().appendHolderSlice(&ospfv2OpaqueLsaSubTlv{obj: item})
			}
		}
		for _, item := range obj.SubTlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv2OpaqueLsaTlv) setDefault() {

}
