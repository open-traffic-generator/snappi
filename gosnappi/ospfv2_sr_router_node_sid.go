package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2SRRouterNodeSid *****
type ospfv2SRRouterNodeSid struct {
	validation
	obj                        *otg.Ospfv2SRRouterNodeSid
	marshaller                 marshalOspfv2SRRouterNodeSid
	unMarshaller               unMarshalOspfv2SRRouterNodeSid
	additionalPrefixSidsHolder Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter
}

func NewOspfv2SRRouterNodeSid() Ospfv2SRRouterNodeSid {
	obj := ospfv2SRRouterNodeSid{obj: &otg.Ospfv2SRRouterNodeSid{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2SRRouterNodeSid) msg() *otg.Ospfv2SRRouterNodeSid {
	return obj.obj
}

func (obj *ospfv2SRRouterNodeSid) setMsg(msg *otg.Ospfv2SRRouterNodeSid) Ospfv2SRRouterNodeSid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2SRRouterNodeSid struct {
	obj *ospfv2SRRouterNodeSid
}

type marshalOspfv2SRRouterNodeSid interface {
	// ToProto marshals Ospfv2SRRouterNodeSid to protobuf object *otg.Ospfv2SRRouterNodeSid
	ToProto() (*otg.Ospfv2SRRouterNodeSid, error)
	// ToPbText marshals Ospfv2SRRouterNodeSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2SRRouterNodeSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2SRRouterNodeSid to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2SRRouterNodeSid struct {
	obj *ospfv2SRRouterNodeSid
}

type unMarshalOspfv2SRRouterNodeSid interface {
	// FromProto unmarshals Ospfv2SRRouterNodeSid from protobuf object *otg.Ospfv2SRRouterNodeSid
	FromProto(msg *otg.Ospfv2SRRouterNodeSid) (Ospfv2SRRouterNodeSid, error)
	// FromPbText unmarshals Ospfv2SRRouterNodeSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2SRRouterNodeSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2SRRouterNodeSid from JSON text
	FromJson(value string) error
}

func (obj *ospfv2SRRouterNodeSid) Marshal() marshalOspfv2SRRouterNodeSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2SRRouterNodeSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2SRRouterNodeSid) Unmarshal() unMarshalOspfv2SRRouterNodeSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2SRRouterNodeSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2SRRouterNodeSid) ToProto() (*otg.Ospfv2SRRouterNodeSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2SRRouterNodeSid) FromProto(msg *otg.Ospfv2SRRouterNodeSid) (Ospfv2SRRouterNodeSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2SRRouterNodeSid) ToPbText() (string, error) {
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

func (m *unMarshalospfv2SRRouterNodeSid) FromPbText(value string) error {
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

func (m *marshalospfv2SRRouterNodeSid) ToYaml() (string, error) {
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

func (m *unMarshalospfv2SRRouterNodeSid) FromYaml(value string) error {
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

func (m *marshalospfv2SRRouterNodeSid) ToJson() (string, error) {
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

func (m *unMarshalospfv2SRRouterNodeSid) FromJson(value string) error {
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

func (obj *ospfv2SRRouterNodeSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2SRRouterNodeSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2SRRouterNodeSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2SRRouterNodeSid) Clone() (Ospfv2SRRouterNodeSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2SRRouterNodeSid()
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

func (obj *ospfv2SRRouterNodeSid) setNil() {
	obj.additionalPrefixSidsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2SRRouterNodeSid is the Node (loopback) Prefix-SID advertised by this router for its own loopback address.
// It is advertised as a Prefix-SID sub-TLV inside the Extended Prefix TLV of the Extended
// Prefix Opaque LSA, together with the one-octet Extended Prefix flags.
// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-prefix-sid-sub-tlv.
type Ospfv2SRRouterNodeSid interface {
	Validation
	// msg marshals Ospfv2SRRouterNodeSid to protobuf object *otg.Ospfv2SRRouterNodeSid
	// and doesn't set defaults
	msg() *otg.Ospfv2SRRouterNodeSid
	// setMsg unmarshals Ospfv2SRRouterNodeSid from protobuf object *otg.Ospfv2SRRouterNodeSid
	// and doesn't set defaults
	setMsg(*otg.Ospfv2SRRouterNodeSid) Ospfv2SRRouterNodeSid
	// provides marshal interface
	Marshal() marshalOspfv2SRRouterNodeSid
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2SRRouterNodeSid
	// validate validates Ospfv2SRRouterNodeSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2SRRouterNodeSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv2SRRouterNodeSidChoiceEnum, set in Ospfv2SRRouterNodeSid
	Choice() Ospfv2SRRouterNodeSidChoiceEnum
	// setChoice assigns Ospfv2SRRouterNodeSidChoiceEnum provided by user to Ospfv2SRRouterNodeSid
	setChoice(value Ospfv2SRRouterNodeSidChoiceEnum) Ospfv2SRRouterNodeSid
	// HasChoice checks if Choice has been set in Ospfv2SRRouterNodeSid
	HasChoice() bool
	// SidValue returns uint32, set in Ospfv2SRRouterNodeSid.
	SidValue() uint32
	// SetSidValue assigns uint32 provided by user to Ospfv2SRRouterNodeSid
	SetSidValue(value uint32) Ospfv2SRRouterNodeSid
	// HasSidValue checks if SidValue has been set in Ospfv2SRRouterNodeSid
	HasSidValue() bool
	// SidIndex returns uint32, set in Ospfv2SRRouterNodeSid.
	SidIndex() uint32
	// SetSidIndex assigns uint32 provided by user to Ospfv2SRRouterNodeSid
	SetSidIndex(value uint32) Ospfv2SRRouterNodeSid
	// HasSidIndex checks if SidIndex has been set in Ospfv2SRRouterNodeSid
	HasSidIndex() bool
	// NpFlag returns bool, set in Ospfv2SRRouterNodeSid.
	NpFlag() bool
	// SetNpFlag assigns bool provided by user to Ospfv2SRRouterNodeSid
	SetNpFlag(value bool) Ospfv2SRRouterNodeSid
	// HasNpFlag checks if NpFlag has been set in Ospfv2SRRouterNodeSid
	HasNpFlag() bool
	// MFlag returns bool, set in Ospfv2SRRouterNodeSid.
	MFlag() bool
	// SetMFlag assigns bool provided by user to Ospfv2SRRouterNodeSid
	SetMFlag(value bool) Ospfv2SRRouterNodeSid
	// HasMFlag checks if MFlag has been set in Ospfv2SRRouterNodeSid
	HasMFlag() bool
	// EFlag returns bool, set in Ospfv2SRRouterNodeSid.
	EFlag() bool
	// SetEFlag assigns bool provided by user to Ospfv2SRRouterNodeSid
	SetEFlag(value bool) Ospfv2SRRouterNodeSid
	// HasEFlag checks if EFlag has been set in Ospfv2SRRouterNodeSid
	HasEFlag() bool
	// Algorithm returns uint32, set in Ospfv2SRRouterNodeSid.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to Ospfv2SRRouterNodeSid
	SetAlgorithm(value uint32) Ospfv2SRRouterNodeSid
	// HasAlgorithm checks if Algorithm has been set in Ospfv2SRRouterNodeSid
	HasAlgorithm() bool
	// NFlag returns bool, set in Ospfv2SRRouterNodeSid.
	NFlag() bool
	// SetNFlag assigns bool provided by user to Ospfv2SRRouterNodeSid
	SetNFlag(value bool) Ospfv2SRRouterNodeSid
	// HasNFlag checks if NFlag has been set in Ospfv2SRRouterNodeSid
	HasNFlag() bool
	// AFlag returns bool, set in Ospfv2SRRouterNodeSid.
	AFlag() bool
	// SetAFlag assigns bool provided by user to Ospfv2SRRouterNodeSid
	SetAFlag(value bool) Ospfv2SRRouterNodeSid
	// HasAFlag checks if AFlag has been set in Ospfv2SRRouterNodeSid
	HasAFlag() bool
	// AdditionalPrefixSids returns Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIterIter, set in Ospfv2SRRouterNodeSid
	AdditionalPrefixSids() Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter
	setNil()
}

type Ospfv2SRRouterNodeSidChoiceEnum string

// Enum of Choice on Ospfv2SRRouterNodeSid
var Ospfv2SRRouterNodeSidChoice = struct {
	SID_VALUE Ospfv2SRRouterNodeSidChoiceEnum
	SID_INDEX Ospfv2SRRouterNodeSidChoiceEnum
}{
	SID_VALUE: Ospfv2SRRouterNodeSidChoiceEnum("sid_value"),
	SID_INDEX: Ospfv2SRRouterNodeSidChoiceEnum("sid_index"),
}

func (obj *ospfv2SRRouterNodeSid) Choice() Ospfv2SRRouterNodeSidChoiceEnum {
	return Ospfv2SRRouterNodeSidChoiceEnum(obj.obj.Choice.Enum().String())
}

// Choice of whether the Node Prefix-SID carries an absolute value (local label) or a
// relative index into the SRGB. This choice sets the V-Flag (Value/Index) and the L-Flag
// (Local/Global) as follows:
// - sid_index: V-Flag and L-Flag are unset (both 0). The SID carries a 4-octet index that
// is an offset into the SRGB. Please refer to device.ospfv2.segment_routing.srgb_ranges.
// - sid_value: V-Flag and L-Flag are set (both 1). The SID carries a 3-octet local label
// value with local significance.
// Choice returns a string
func (obj *ospfv2SRRouterNodeSid) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv2SRRouterNodeSid) setChoice(value Ospfv2SRRouterNodeSidChoiceEnum) Ospfv2SRRouterNodeSid {
	intValue, ok := otg.Ospfv2SRRouterNodeSid_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2SRRouterNodeSidChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2SRRouterNodeSid_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.SidIndex = nil
	obj.obj.SidValue = nil

	if value == Ospfv2SRRouterNodeSidChoice.SID_VALUE {
		defaultValue := uint32(16)
		obj.obj.SidValue = &defaultValue
	}

	if value == Ospfv2SRRouterNodeSidChoice.SID_INDEX {
		defaultValue := uint32(0)
		obj.obj.SidIndex = &defaultValue
	}

	return obj
}

// The Node SID/Label as an absolute local label value. Used when the choice is sid_value.
// SidValue returns a uint32
func (obj *ospfv2SRRouterNodeSid) SidValue() uint32 {

	if obj.obj.SidValue == nil {
		obj.setChoice(Ospfv2SRRouterNodeSidChoice.SID_VALUE)
	}

	return *obj.obj.SidValue

}

// The Node SID/Label as an absolute local label value. Used when the choice is sid_value.
// SidValue returns a uint32
func (obj *ospfv2SRRouterNodeSid) HasSidValue() bool {
	return obj.obj.SidValue != nil
}

// The Node SID/Label as an absolute local label value. Used when the choice is sid_value.
// SetSidValue sets the uint32 value in the Ospfv2SRRouterNodeSid object
func (obj *ospfv2SRRouterNodeSid) SetSidValue(value uint32) Ospfv2SRRouterNodeSid {
	obj.setChoice(Ospfv2SRRouterNodeSidChoice.SID_VALUE)
	obj.obj.SidValue = &value
	return obj
}

// The Node SID/Label index, an offset into the SRGB. Used when the choice is sid_index.
// SidIndex returns a uint32
func (obj *ospfv2SRRouterNodeSid) SidIndex() uint32 {

	if obj.obj.SidIndex == nil {
		obj.setChoice(Ospfv2SRRouterNodeSidChoice.SID_INDEX)
	}

	return *obj.obj.SidIndex

}

// The Node SID/Label index, an offset into the SRGB. Used when the choice is sid_index.
// SidIndex returns a uint32
func (obj *ospfv2SRRouterNodeSid) HasSidIndex() bool {
	return obj.obj.SidIndex != nil
}

// The Node SID/Label index, an offset into the SRGB. Used when the choice is sid_index.
// SetSidIndex sets the uint32 value in the Ospfv2SRRouterNodeSid object
func (obj *ospfv2SRRouterNodeSid) SetSidIndex(value uint32) Ospfv2SRRouterNodeSid {
	obj.setChoice(Ospfv2SRRouterNodeSidChoice.SID_INDEX)
	obj.obj.SidIndex = &value
	return obj
}

// NP-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the penultimate hop MUST NOT pop the Prefix-SID before delivering the
// packet to this node.
// NpFlag returns a bool
func (obj *ospfv2SRRouterNodeSid) NpFlag() bool {

	return *obj.obj.NpFlag

}

// NP-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the penultimate hop MUST NOT pop the Prefix-SID before delivering the
// packet to this node.
// NpFlag returns a bool
func (obj *ospfv2SRRouterNodeSid) HasNpFlag() bool {
	return obj.obj.NpFlag != nil
}

// NP-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the penultimate hop MUST NOT pop the Prefix-SID before delivering the
// packet to this node.
// SetNpFlag sets the bool value in the Ospfv2SRRouterNodeSid object
func (obj *ospfv2SRRouterNodeSid) SetNpFlag(value bool) Ospfv2SRRouterNodeSid {

	obj.obj.NpFlag = &value
	return obj
}

// M-Flag: Mapping Server Flag.
// If set, then the SID was advertised by an SR Mapping Server. When set, the NP-Flag and
// the E-Flag MUST be ignored on reception.
// MFlag returns a bool
func (obj *ospfv2SRRouterNodeSid) MFlag() bool {

	return *obj.obj.MFlag

}

// M-Flag: Mapping Server Flag.
// If set, then the SID was advertised by an SR Mapping Server. When set, the NP-Flag and
// the E-Flag MUST be ignored on reception.
// MFlag returns a bool
func (obj *ospfv2SRRouterNodeSid) HasMFlag() bool {
	return obj.obj.MFlag != nil
}

// M-Flag: Mapping Server Flag.
// If set, then the SID was advertised by an SR Mapping Server. When set, the NP-Flag and
// the E-Flag MUST be ignored on reception.
// SetMFlag sets the bool value in the Ospfv2SRRouterNodeSid object
func (obj *ospfv2SRRouterNodeSid) SetMFlag(value bool) Ospfv2SRRouterNodeSid {

	obj.obj.MFlag = &value
	return obj
}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID
// with the Explicit-NULL label before forwarding the packet.
// EFlag returns a bool
func (obj *ospfv2SRRouterNodeSid) EFlag() bool {

	return *obj.obj.EFlag

}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID
// with the Explicit-NULL label before forwarding the packet.
// EFlag returns a bool
func (obj *ospfv2SRRouterNodeSid) HasEFlag() bool {
	return obj.obj.EFlag != nil
}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID
// with the Explicit-NULL label before forwarding the packet.
// SetEFlag sets the bool value in the Ospfv2SRRouterNodeSid object
func (obj *ospfv2SRRouterNodeSid) SetEFlag(value bool) Ospfv2SRRouterNodeSid {

	obj.obj.EFlag = &value
	return obj
}

// The Segment Routing Algorithm the Node Prefix-SID is associated with, e.g. 0 for SPF
// or 1 for Strict SPF.
// Algorithm returns a uint32
func (obj *ospfv2SRRouterNodeSid) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// The Segment Routing Algorithm the Node Prefix-SID is associated with, e.g. 0 for SPF
// or 1 for Strict SPF.
// Algorithm returns a uint32
func (obj *ospfv2SRRouterNodeSid) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// The Segment Routing Algorithm the Node Prefix-SID is associated with, e.g. 0 for SPF
// or 1 for Strict SPF.
// SetAlgorithm sets the uint32 value in the Ospfv2SRRouterNodeSid object
func (obj *ospfv2SRRouterNodeSid) SetAlgorithm(value uint32) Ospfv2SRRouterNodeSid {

	obj.obj.Algorithm = &value
	return obj
}

// N-Flag (Node Flag) of the Extended Prefix TLV flags. Set when the prefix identifies the
// advertising router, i.e. it is a host prefix advertising a globally reachable address
// typically associated with a loopback address. This is normally set for a Node SID.
// NFlag returns a bool
func (obj *ospfv2SRRouterNodeSid) NFlag() bool {

	return *obj.obj.NFlag

}

// N-Flag (Node Flag) of the Extended Prefix TLV flags. Set when the prefix identifies the
// advertising router, i.e. it is a host prefix advertising a globally reachable address
// typically associated with a loopback address. This is normally set for a Node SID.
// NFlag returns a bool
func (obj *ospfv2SRRouterNodeSid) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// N-Flag (Node Flag) of the Extended Prefix TLV flags. Set when the prefix identifies the
// advertising router, i.e. it is a host prefix advertising a globally reachable address
// typically associated with a loopback address. This is normally set for a Node SID.
// SetNFlag sets the bool value in the Ospfv2SRRouterNodeSid object
func (obj *ospfv2SRRouterNodeSid) SetNFlag(value bool) Ospfv2SRRouterNodeSid {

	obj.obj.NFlag = &value
	return obj
}

// A-Flag (Attach Flag) of the Extended Prefix TLV flags. An Area Border Router (ABR)
// generating an OSPFv2 Extended Prefix TLV for an inter-area prefix that is locally
// connected or attached in another connected area SHOULD set this flag.
// AFlag returns a bool
func (obj *ospfv2SRRouterNodeSid) AFlag() bool {

	return *obj.obj.AFlag

}

// A-Flag (Attach Flag) of the Extended Prefix TLV flags. An Area Border Router (ABR)
// generating an OSPFv2 Extended Prefix TLV for an inter-area prefix that is locally
// connected or attached in another connected area SHOULD set this flag.
// AFlag returns a bool
func (obj *ospfv2SRRouterNodeSid) HasAFlag() bool {
	return obj.obj.AFlag != nil
}

// A-Flag (Attach Flag) of the Extended Prefix TLV flags. An Area Border Router (ABR)
// generating an OSPFv2 Extended Prefix TLV for an inter-area prefix that is locally
// connected or attached in another connected area SHOULD set this flag.
// SetAFlag sets the bool value in the Ospfv2SRRouterNodeSid object
func (obj *ospfv2SRRouterNodeSid) SetAFlag(value bool) Ospfv2SRRouterNodeSid {

	obj.obj.AFlag = &value
	return obj
}

// An optional list of additional Node Prefix-SIDs advertised for the same loopback
// prefix but with different Segment Routing algorithms (one Prefix-SID sub-TLV per
// algorithm).
// AdditionalPrefixSids returns a []Ospfv2SRPrefixSid
func (obj *ospfv2SRRouterNodeSid) AdditionalPrefixSids() Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter {
	if len(obj.obj.AdditionalPrefixSids) == 0 {
		obj.obj.AdditionalPrefixSids = []*otg.Ospfv2SRPrefixSid{}
	}
	if obj.additionalPrefixSidsHolder == nil {
		obj.additionalPrefixSidsHolder = newOspfv2SRRouterNodeSidOspfv2SRPrefixSidIter(&obj.obj.AdditionalPrefixSids).setMsg(obj)
	}
	return obj.additionalPrefixSidsHolder
}

type ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter struct {
	obj                    *ospfv2SRRouterNodeSid
	ospfv2SRPrefixSidSlice []Ospfv2SRPrefixSid
	fieldPtr               *[]*otg.Ospfv2SRPrefixSid
}

func newOspfv2SRRouterNodeSidOspfv2SRPrefixSidIter(ptr *[]*otg.Ospfv2SRPrefixSid) Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter {
	return &ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter{fieldPtr: ptr}
}

type Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter interface {
	setMsg(*ospfv2SRRouterNodeSid) Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter
	Items() []Ospfv2SRPrefixSid
	Add() Ospfv2SRPrefixSid
	Append(items ...Ospfv2SRPrefixSid) Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter
	Set(index int, newObj Ospfv2SRPrefixSid) Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter
	Clear() Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter
	clearHolderSlice() Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter
	appendHolderSlice(item Ospfv2SRPrefixSid) Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter
}

func (obj *ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter) setMsg(msg *ospfv2SRRouterNodeSid) Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2SRPrefixSid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter) Items() []Ospfv2SRPrefixSid {
	return obj.ospfv2SRPrefixSidSlice
}

func (obj *ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter) Add() Ospfv2SRPrefixSid {
	newObj := &otg.Ospfv2SRPrefixSid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2SRPrefixSid{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2SRPrefixSidSlice = append(obj.ospfv2SRPrefixSidSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter) Append(items ...Ospfv2SRPrefixSid) Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2SRPrefixSidSlice = append(obj.ospfv2SRPrefixSidSlice, item)
	}
	return obj
}

func (obj *ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter) Set(index int, newObj Ospfv2SRPrefixSid) Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2SRPrefixSidSlice[index] = newObj
	return obj
}
func (obj *ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter) Clear() Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2SRPrefixSid{}
		obj.ospfv2SRPrefixSidSlice = []Ospfv2SRPrefixSid{}
	}
	return obj
}
func (obj *ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter) clearHolderSlice() Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter {
	if len(obj.ospfv2SRPrefixSidSlice) > 0 {
		obj.ospfv2SRPrefixSidSlice = []Ospfv2SRPrefixSid{}
	}
	return obj
}
func (obj *ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter) appendHolderSlice(item Ospfv2SRPrefixSid) Ospfv2SRRouterNodeSidOspfv2SRPrefixSidIter {
	obj.ospfv2SRPrefixSidSlice = append(obj.ospfv2SRPrefixSidSlice, item)
	return obj
}

func (obj *ospfv2SRRouterNodeSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SidValue != nil {

		if *obj.obj.SidValue < 16 || *obj.obj.SidValue > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("16 <= Ospfv2SRRouterNodeSid.SidValue <= 1048575 but Got %d", *obj.obj.SidValue))
		}

	}

	if obj.obj.SidIndex != nil {

		if *obj.obj.SidIndex > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv2SRRouterNodeSid.SidIndex <= 4294967295 but Got %d", *obj.obj.SidIndex))
		}

	}

	if obj.obj.Algorithm != nil {

		if *obj.obj.Algorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv2SRRouterNodeSid.Algorithm <= 255 but Got %d", *obj.obj.Algorithm))
		}

	}

	if len(obj.obj.AdditionalPrefixSids) != 0 {

		if set_default {
			obj.AdditionalPrefixSids().clearHolderSlice()
			for _, item := range obj.obj.AdditionalPrefixSids {
				obj.AdditionalPrefixSids().appendHolderSlice(&ospfv2SRPrefixSid{obj: item})
			}
		}
		for _, item := range obj.AdditionalPrefixSids().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv2SRRouterNodeSid) setDefault() {
	var choices_set int = 0
	var choice Ospfv2SRRouterNodeSidChoiceEnum

	if obj.obj.SidValue != nil {
		choices_set += 1
		choice = Ospfv2SRRouterNodeSidChoice.SID_VALUE
	}

	if obj.obj.SidIndex != nil {
		choices_set += 1
		choice = Ospfv2SRRouterNodeSidChoice.SID_INDEX
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv2SRRouterNodeSidChoice.SID_INDEX)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv2SRRouterNodeSid")
			}
		} else {
			intVal := otg.Ospfv2SRRouterNodeSid_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv2SRRouterNodeSid_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.NpFlag == nil {
		obj.SetNpFlag(true)
	}
	if obj.obj.MFlag == nil {
		obj.SetMFlag(false)
	}
	if obj.obj.EFlag == nil {
		obj.SetEFlag(false)
	}
	if obj.obj.Algorithm == nil {
		obj.SetAlgorithm(0)
	}
	if obj.obj.NFlag == nil {
		obj.SetNFlag(true)
	}
	if obj.obj.AFlag == nil {
		obj.SetAFlag(false)
	}

}
