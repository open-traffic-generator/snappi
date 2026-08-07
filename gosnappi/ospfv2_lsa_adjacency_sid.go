package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaAdjacencySid *****
type ospfv2LsaAdjacencySid struct {
	validation
	obj          *otg.Ospfv2LsaAdjacencySid
	marshaller   marshalOspfv2LsaAdjacencySid
	unMarshaller unMarshalOspfv2LsaAdjacencySid
	flagsHolder  Ospfv2LsaAdjSidFlags
}

func NewOspfv2LsaAdjacencySid() Ospfv2LsaAdjacencySid {
	obj := ospfv2LsaAdjacencySid{obj: &otg.Ospfv2LsaAdjacencySid{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaAdjacencySid) msg() *otg.Ospfv2LsaAdjacencySid {
	return obj.obj
}

func (obj *ospfv2LsaAdjacencySid) setMsg(msg *otg.Ospfv2LsaAdjacencySid) Ospfv2LsaAdjacencySid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaAdjacencySid struct {
	obj *ospfv2LsaAdjacencySid
}

type marshalOspfv2LsaAdjacencySid interface {
	// ToProto marshals Ospfv2LsaAdjacencySid to protobuf object *otg.Ospfv2LsaAdjacencySid
	ToProto() (*otg.Ospfv2LsaAdjacencySid, error)
	// ToPbText marshals Ospfv2LsaAdjacencySid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaAdjacencySid to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaAdjacencySid to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaAdjacencySid struct {
	obj *ospfv2LsaAdjacencySid
}

type unMarshalOspfv2LsaAdjacencySid interface {
	// FromProto unmarshals Ospfv2LsaAdjacencySid from protobuf object *otg.Ospfv2LsaAdjacencySid
	FromProto(msg *otg.Ospfv2LsaAdjacencySid) (Ospfv2LsaAdjacencySid, error)
	// FromPbText unmarshals Ospfv2LsaAdjacencySid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaAdjacencySid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaAdjacencySid from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaAdjacencySid) Marshal() marshalOspfv2LsaAdjacencySid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaAdjacencySid{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaAdjacencySid) Unmarshal() unMarshalOspfv2LsaAdjacencySid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaAdjacencySid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaAdjacencySid) ToProto() (*otg.Ospfv2LsaAdjacencySid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaAdjacencySid) FromProto(msg *otg.Ospfv2LsaAdjacencySid) (Ospfv2LsaAdjacencySid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaAdjacencySid) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaAdjacencySid) FromPbText(value string) error {
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

func (m *marshalospfv2LsaAdjacencySid) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaAdjacencySid) FromYaml(value string) error {
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

func (m *marshalospfv2LsaAdjacencySid) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaAdjacencySid) FromJson(value string) error {
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

func (obj *ospfv2LsaAdjacencySid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaAdjacencySid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaAdjacencySid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaAdjacencySid) Clone() (Ospfv2LsaAdjacencySid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaAdjacencySid()
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

func (obj *ospfv2LsaAdjacencySid) setNil() {
	obj.flagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2LsaAdjacencySid is the learned OSPFv2 Adjacency-SID and its attributes, decoded from the Adj-SID / LAN Adj-SID
// sub-TLV of the Extended Link Opaque LSA (RFC 8665).
type Ospfv2LsaAdjacencySid interface {
	Validation
	// msg marshals Ospfv2LsaAdjacencySid to protobuf object *otg.Ospfv2LsaAdjacencySid
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaAdjacencySid
	// setMsg unmarshals Ospfv2LsaAdjacencySid from protobuf object *otg.Ospfv2LsaAdjacencySid
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaAdjacencySid) Ospfv2LsaAdjacencySid
	// provides marshal interface
	Marshal() marshalOspfv2LsaAdjacencySid
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaAdjacencySid
	// validate validates Ospfv2LsaAdjacencySid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaAdjacencySid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns Ospfv2LsaAdjacencySidTypeEnum, set in Ospfv2LsaAdjacencySid
	Type() Ospfv2LsaAdjacencySidTypeEnum
	// SetType assigns Ospfv2LsaAdjacencySidTypeEnum provided by user to Ospfv2LsaAdjacencySid
	SetType(value Ospfv2LsaAdjacencySidTypeEnum) Ospfv2LsaAdjacencySid
	// HasType checks if Type has been set in Ospfv2LsaAdjacencySid
	HasType() bool
	// Sids returns []uint32, set in Ospfv2LsaAdjacencySid.
	Sids() []uint32
	// SetSids assigns []uint32 provided by user to Ospfv2LsaAdjacencySid
	SetSids(value []uint32) Ospfv2LsaAdjacencySid
	// Flags returns Ospfv2LsaAdjSidFlags, set in Ospfv2LsaAdjacencySid.
	// Ospfv2LsaAdjSidFlags is one-octet flags of the OSPFv2 Adjacency-SID sub-TLV (RFC 8665).
	Flags() Ospfv2LsaAdjSidFlags
	// SetFlags assigns Ospfv2LsaAdjSidFlags provided by user to Ospfv2LsaAdjacencySid.
	// Ospfv2LsaAdjSidFlags is one-octet flags of the OSPFv2 Adjacency-SID sub-TLV (RFC 8665).
	SetFlags(value Ospfv2LsaAdjSidFlags) Ospfv2LsaAdjacencySid
	// HasFlags checks if Flags has been set in Ospfv2LsaAdjacencySid
	HasFlags() bool
	// Weight returns uint32, set in Ospfv2LsaAdjacencySid.
	Weight() uint32
	// SetWeight assigns uint32 provided by user to Ospfv2LsaAdjacencySid
	SetWeight(value uint32) Ospfv2LsaAdjacencySid
	// HasWeight checks if Weight has been set in Ospfv2LsaAdjacencySid
	HasWeight() bool
	setNil()
}

type Ospfv2LsaAdjacencySidTypeEnum string

// Enum of Type on Ospfv2LsaAdjacencySid
var Ospfv2LsaAdjacencySidType = struct {
	ADJ_SID     Ospfv2LsaAdjacencySidTypeEnum
	LAN_ADJ_SID Ospfv2LsaAdjacencySidTypeEnum
}{
	ADJ_SID:     Ospfv2LsaAdjacencySidTypeEnum("adj_sid"),
	LAN_ADJ_SID: Ospfv2LsaAdjacencySidTypeEnum("lan_adj_sid"),
}

func (obj *ospfv2LsaAdjacencySid) Type() Ospfv2LsaAdjacencySidTypeEnum {
	return Ospfv2LsaAdjacencySidTypeEnum(obj.obj.Type.Enum().String())
}

// Adjacency-SID type: Adjacency-SID (Extended Link sub-TLV Type 2) or LAN Adjacency-SID (Type 3).
// Type returns a string
func (obj *ospfv2LsaAdjacencySid) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *ospfv2LsaAdjacencySid) SetType(value Ospfv2LsaAdjacencySidTypeEnum) Ospfv2LsaAdjacencySid {
	intValue, ok := otg.Ospfv2LsaAdjacencySid_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2LsaAdjacencySidTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2LsaAdjacencySid_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// One or more SID/Label values or indices associated with the adjacency.
// Sids returns a []uint32
func (obj *ospfv2LsaAdjacencySid) Sids() []uint32 {
	if obj.obj.Sids == nil {
		obj.obj.Sids = make([]uint32, 0)
	}
	return obj.obj.Sids
}

// One or more SID/Label values or indices associated with the adjacency.
// SetSids sets the []uint32 value in the Ospfv2LsaAdjacencySid object
func (obj *ospfv2LsaAdjacencySid) SetSids(value []uint32) Ospfv2LsaAdjacencySid {

	if obj.obj.Sids == nil {
		obj.obj.Sids = make([]uint32, 0)
	}
	obj.obj.Sids = value

	return obj
}

// Flags associated with the Adjacency-SID.
// Flags returns a Ospfv2LsaAdjSidFlags
func (obj *ospfv2LsaAdjacencySid) Flags() Ospfv2LsaAdjSidFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewOspfv2LsaAdjSidFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &ospfv2LsaAdjSidFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// Flags associated with the Adjacency-SID.
// Flags returns a Ospfv2LsaAdjSidFlags
func (obj *ospfv2LsaAdjacencySid) HasFlags() bool {
	return obj.obj.Flags != nil
}

// Flags associated with the Adjacency-SID.
// SetFlags sets the Ospfv2LsaAdjSidFlags value in the Ospfv2LsaAdjacencySid object
func (obj *ospfv2LsaAdjacencySid) SetFlags(value Ospfv2LsaAdjSidFlags) Ospfv2LsaAdjacencySid {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// The weight of the Adjacency-SID for the purpose of load balancing.
// Weight returns a uint32
func (obj *ospfv2LsaAdjacencySid) Weight() uint32 {

	return *obj.obj.Weight

}

// The weight of the Adjacency-SID for the purpose of load balancing.
// Weight returns a uint32
func (obj *ospfv2LsaAdjacencySid) HasWeight() bool {
	return obj.obj.Weight != nil
}

// The weight of the Adjacency-SID for the purpose of load balancing.
// SetWeight sets the uint32 value in the Ospfv2LsaAdjacencySid object
func (obj *ospfv2LsaAdjacencySid) SetWeight(value uint32) Ospfv2LsaAdjacencySid {

	obj.obj.Weight = &value
	return obj
}

func (obj *ospfv2LsaAdjacencySid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

}

func (obj *ospfv2LsaAdjacencySid) setDefault() {

}
