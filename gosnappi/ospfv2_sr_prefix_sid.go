package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2SRPrefixSid *****
type ospfv2SRPrefixSid struct {
	validation
	obj          *otg.Ospfv2SRPrefixSid
	marshaller   marshalOspfv2SRPrefixSid
	unMarshaller unMarshalOspfv2SRPrefixSid
}

func NewOspfv2SRPrefixSid() Ospfv2SRPrefixSid {
	obj := ospfv2SRPrefixSid{obj: &otg.Ospfv2SRPrefixSid{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2SRPrefixSid) msg() *otg.Ospfv2SRPrefixSid {
	return obj.obj
}

func (obj *ospfv2SRPrefixSid) setMsg(msg *otg.Ospfv2SRPrefixSid) Ospfv2SRPrefixSid {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2SRPrefixSid struct {
	obj *ospfv2SRPrefixSid
}

type marshalOspfv2SRPrefixSid interface {
	// ToProto marshals Ospfv2SRPrefixSid to protobuf object *otg.Ospfv2SRPrefixSid
	ToProto() (*otg.Ospfv2SRPrefixSid, error)
	// ToPbText marshals Ospfv2SRPrefixSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2SRPrefixSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2SRPrefixSid to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2SRPrefixSid struct {
	obj *ospfv2SRPrefixSid
}

type unMarshalOspfv2SRPrefixSid interface {
	// FromProto unmarshals Ospfv2SRPrefixSid from protobuf object *otg.Ospfv2SRPrefixSid
	FromProto(msg *otg.Ospfv2SRPrefixSid) (Ospfv2SRPrefixSid, error)
	// FromPbText unmarshals Ospfv2SRPrefixSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2SRPrefixSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2SRPrefixSid from JSON text
	FromJson(value string) error
}

func (obj *ospfv2SRPrefixSid) Marshal() marshalOspfv2SRPrefixSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2SRPrefixSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2SRPrefixSid) Unmarshal() unMarshalOspfv2SRPrefixSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2SRPrefixSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2SRPrefixSid) ToProto() (*otg.Ospfv2SRPrefixSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2SRPrefixSid) FromProto(msg *otg.Ospfv2SRPrefixSid) (Ospfv2SRPrefixSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2SRPrefixSid) ToPbText() (string, error) {
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

func (m *unMarshalospfv2SRPrefixSid) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalospfv2SRPrefixSid) ToYaml() (string, error) {
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

func (m *unMarshalospfv2SRPrefixSid) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalospfv2SRPrefixSid) ToJson() (string, error) {
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

func (m *unMarshalospfv2SRPrefixSid) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *ospfv2SRPrefixSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2SRPrefixSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2SRPrefixSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2SRPrefixSid) Clone() (Ospfv2SRPrefixSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2SRPrefixSid()
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

// Ospfv2SRPrefixSid is this contains the properties of an OSPFv2 Prefix-SID sub-TLV and its attributes.
// The Prefix-SID sub-TLV is carried inside the OSPFv2 Extended Prefix TLV of the
// Extended Prefix Opaque LSA and is associated with a specific IPv4 prefix.
// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-prefix-sid-sub-tlv.
type Ospfv2SRPrefixSid interface {
	Validation
	// msg marshals Ospfv2SRPrefixSid to protobuf object *otg.Ospfv2SRPrefixSid
	// and doesn't set defaults
	msg() *otg.Ospfv2SRPrefixSid
	// setMsg unmarshals Ospfv2SRPrefixSid from protobuf object *otg.Ospfv2SRPrefixSid
	// and doesn't set defaults
	setMsg(*otg.Ospfv2SRPrefixSid) Ospfv2SRPrefixSid
	// provides marshal interface
	Marshal() marshalOspfv2SRPrefixSid
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2SRPrefixSid
	// validate validates Ospfv2SRPrefixSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2SRPrefixSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv2SRPrefixSidChoiceEnum, set in Ospfv2SRPrefixSid
	Choice() Ospfv2SRPrefixSidChoiceEnum
	// setChoice assigns Ospfv2SRPrefixSidChoiceEnum provided by user to Ospfv2SRPrefixSid
	setChoice(value Ospfv2SRPrefixSidChoiceEnum) Ospfv2SRPrefixSid
	// HasChoice checks if Choice has been set in Ospfv2SRPrefixSid
	HasChoice() bool
	// SidValues returns []uint32, set in Ospfv2SRPrefixSid.
	SidValues() []uint32
	// SetSidValues assigns []uint32 provided by user to Ospfv2SRPrefixSid
	SetSidValues(value []uint32) Ospfv2SRPrefixSid
	// SidIndices returns []uint32, set in Ospfv2SRPrefixSid.
	SidIndices() []uint32
	// SetSidIndices assigns []uint32 provided by user to Ospfv2SRPrefixSid
	SetSidIndices(value []uint32) Ospfv2SRPrefixSid
	// NpFlag returns bool, set in Ospfv2SRPrefixSid.
	NpFlag() bool
	// SetNpFlag assigns bool provided by user to Ospfv2SRPrefixSid
	SetNpFlag(value bool) Ospfv2SRPrefixSid
	// HasNpFlag checks if NpFlag has been set in Ospfv2SRPrefixSid
	HasNpFlag() bool
	// MFlag returns bool, set in Ospfv2SRPrefixSid.
	MFlag() bool
	// SetMFlag assigns bool provided by user to Ospfv2SRPrefixSid
	SetMFlag(value bool) Ospfv2SRPrefixSid
	// HasMFlag checks if MFlag has been set in Ospfv2SRPrefixSid
	HasMFlag() bool
	// EFlag returns bool, set in Ospfv2SRPrefixSid.
	EFlag() bool
	// SetEFlag assigns bool provided by user to Ospfv2SRPrefixSid
	SetEFlag(value bool) Ospfv2SRPrefixSid
	// HasEFlag checks if EFlag has been set in Ospfv2SRPrefixSid
	HasEFlag() bool
	// Algorithm returns uint32, set in Ospfv2SRPrefixSid.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to Ospfv2SRPrefixSid
	SetAlgorithm(value uint32) Ospfv2SRPrefixSid
	// HasAlgorithm checks if Algorithm has been set in Ospfv2SRPrefixSid
	HasAlgorithm() bool
}

type Ospfv2SRPrefixSidChoiceEnum string

// Enum of Choice on Ospfv2SRPrefixSid
var Ospfv2SRPrefixSidChoice = struct {
	SID_VALUES  Ospfv2SRPrefixSidChoiceEnum
	SID_INDICES Ospfv2SRPrefixSidChoiceEnum
}{
	SID_VALUES:  Ospfv2SRPrefixSidChoiceEnum("sid_values"),
	SID_INDICES: Ospfv2SRPrefixSidChoiceEnum("sid_indices"),
}

func (obj *ospfv2SRPrefixSid) Choice() Ospfv2SRPrefixSidChoiceEnum {
	return Ospfv2SRPrefixSidChoiceEnum(obj.obj.Choice.Enum().String())
}

// Choice of whether the Prefix-SID carries absolute values (local labels) or relative
// indices into the SRGB. This choice sets the V-Flag (Value/Index) and the L-Flag
// (Local/Global) of the Prefix-SID sub-TLV as follows:
// - sid_indices: V-Flag and L-Flag are unset (both 0). Each Prefix-SID carries a 4-octet
// index that is an offset into the SRGB advertised by the router. Please refer to
// device.ospfv2.segment_routing.srgb_ranges.
// - sid_values: V-Flag and L-Flag are set (both 1). Each Prefix-SID carries a 3-octet
// local label value with local significance.
// A user needs to configure at least one entry of SID value or SID index. If no entry is
// configured, an implementation may advertise an appropriate default SID value/index
// based on the choice, e.g. the first value from the SRGB range.
// Choice returns a string
func (obj *ospfv2SRPrefixSid) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv2SRPrefixSid) setChoice(value Ospfv2SRPrefixSidChoiceEnum) Ospfv2SRPrefixSid {
	intValue, ok := otg.Ospfv2SRPrefixSid_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2SRPrefixSidChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2SRPrefixSid_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.SidIndices = nil
	obj.obj.SidValues = nil
	return obj
}

// SID/Label as one or more absolute local label values associated with the IGP Prefix segment attached to the specific IPv4 prefix. Used when the choice is sid_values.
// SidValues returns a []uint32
func (obj *ospfv2SRPrefixSid) SidValues() []uint32 {
	if obj.obj.SidValues == nil {

		obj.setChoice(Ospfv2SRPrefixSidChoice.SID_VALUES)

	}
	return obj.obj.SidValues
}

// SID/Label as one or more absolute local label values associated with the IGP Prefix segment attached to the specific IPv4 prefix. Used when the choice is sid_values.
// SetSidValues sets the []uint32 value in the Ospfv2SRPrefixSid object
func (obj *ospfv2SRPrefixSid) SetSidValues(value []uint32) Ospfv2SRPrefixSid {
	obj.setChoice(Ospfv2SRPrefixSidChoice.SID_VALUES)
	if obj.obj.SidValues == nil {
		obj.obj.SidValues = make([]uint32, 0)
	}
	obj.obj.SidValues = value

	return obj
}

// One or more SID/Label indices associated with the IGP Prefix segment attached to the specific IPv4 prefix. Each index is an offset into the SRGB. Used when the choice is sid_indices.
// SidIndices returns a []uint32
func (obj *ospfv2SRPrefixSid) SidIndices() []uint32 {
	if obj.obj.SidIndices == nil {

		obj.setChoice(Ospfv2SRPrefixSidChoice.SID_INDICES)

	}
	return obj.obj.SidIndices
}

// One or more SID/Label indices associated with the IGP Prefix segment attached to the specific IPv4 prefix. Each index is an offset into the SRGB. Used when the choice is sid_indices.
// SetSidIndices sets the []uint32 value in the Ospfv2SRPrefixSid object
func (obj *ospfv2SRPrefixSid) SetSidIndices(value []uint32) Ospfv2SRPrefixSid {
	obj.setChoice(Ospfv2SRPrefixSidChoice.SID_INDICES)
	if obj.obj.SidIndices == nil {
		obj.obj.SidIndices = make([]uint32, 0)
	}
	obj.obj.SidIndices = value

	return obj
}

// NP-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the penultimate hop MUST NOT pop the Prefix-SID before delivering the
// packet to the node that advertised the Prefix-SID.
// NpFlag returns a bool
func (obj *ospfv2SRPrefixSid) NpFlag() bool {

	return *obj.obj.NpFlag

}

// NP-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the penultimate hop MUST NOT pop the Prefix-SID before delivering the
// packet to the node that advertised the Prefix-SID.
// NpFlag returns a bool
func (obj *ospfv2SRPrefixSid) HasNpFlag() bool {
	return obj.obj.NpFlag != nil
}

// NP-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the penultimate hop MUST NOT pop the Prefix-SID before delivering the
// packet to the node that advertised the Prefix-SID.
// SetNpFlag sets the bool value in the Ospfv2SRPrefixSid object
func (obj *ospfv2SRPrefixSid) SetNpFlag(value bool) Ospfv2SRPrefixSid {

	obj.obj.NpFlag = &value
	return obj
}

// M-Flag: Mapping Server Flag.
// If set, then the SID was advertised by an SR Mapping Server. When set, the NP-Flag and
// the E-Flag MUST be ignored on reception.
// MFlag returns a bool
func (obj *ospfv2SRPrefixSid) MFlag() bool {

	return *obj.obj.MFlag

}

// M-Flag: Mapping Server Flag.
// If set, then the SID was advertised by an SR Mapping Server. When set, the NP-Flag and
// the E-Flag MUST be ignored on reception.
// MFlag returns a bool
func (obj *ospfv2SRPrefixSid) HasMFlag() bool {
	return obj.obj.MFlag != nil
}

// M-Flag: Mapping Server Flag.
// If set, then the SID was advertised by an SR Mapping Server. When set, the NP-Flag and
// the E-Flag MUST be ignored on reception.
// SetMFlag sets the bool value in the Ospfv2SRPrefixSid object
func (obj *ospfv2SRPrefixSid) SetMFlag(value bool) Ospfv2SRPrefixSid {

	obj.obj.MFlag = &value
	return obj
}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID
// with the Explicit-NULL label before forwarding the packet.
// EFlag returns a bool
func (obj *ospfv2SRPrefixSid) EFlag() bool {

	return *obj.obj.EFlag

}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID
// with the Explicit-NULL label before forwarding the packet.
// EFlag returns a bool
func (obj *ospfv2SRPrefixSid) HasEFlag() bool {
	return obj.obj.EFlag != nil
}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID
// with the Explicit-NULL label before forwarding the packet.
// SetEFlag sets the bool value in the Ospfv2SRPrefixSid object
func (obj *ospfv2SRPrefixSid) SetEFlag(value bool) Ospfv2SRPrefixSid {

	obj.obj.EFlag = &value
	return obj
}

// The Segment Routing Algorithm the Prefix-SID is associated with. The value matches an
// algorithm advertised in the SR-Algorithm TLV, e.g. 0 for SPF or 1 for Strict SPF.
// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-sr-algorithm-tlv.
// Algorithm returns a uint32
func (obj *ospfv2SRPrefixSid) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// The Segment Routing Algorithm the Prefix-SID is associated with. The value matches an
// algorithm advertised in the SR-Algorithm TLV, e.g. 0 for SPF or 1 for Strict SPF.
// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-sr-algorithm-tlv.
// Algorithm returns a uint32
func (obj *ospfv2SRPrefixSid) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// The Segment Routing Algorithm the Prefix-SID is associated with. The value matches an
// algorithm advertised in the SR-Algorithm TLV, e.g. 0 for SPF or 1 for Strict SPF.
// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-sr-algorithm-tlv.
// SetAlgorithm sets the uint32 value in the Ospfv2SRPrefixSid object
func (obj *ospfv2SRPrefixSid) SetAlgorithm(value uint32) Ospfv2SRPrefixSid {

	obj.obj.Algorithm = &value
	return obj
}

func (obj *ospfv2SRPrefixSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SidValues != nil {

		for _, item := range obj.obj.SidValues {
			if item < 16 || item > 1048575 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("16 <= Ospfv2SRPrefixSid.SidValues <= 1048575 but Got %d", item))
			}

		}

	}

	if obj.obj.SidIndices != nil {

		for _, item := range obj.obj.SidIndices {
			if item > 4294967295 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("0 <= Ospfv2SRPrefixSid.SidIndices <= 4294967295 but Got %d", item))
			}

		}

	}

	if obj.obj.Algorithm != nil {

		if *obj.obj.Algorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv2SRPrefixSid.Algorithm <= 255 but Got %d", *obj.obj.Algorithm))
		}

	}

}

func (obj *ospfv2SRPrefixSid) setDefault() {
	var choices_set int = 0
	var choice Ospfv2SRPrefixSidChoiceEnum

	if len(obj.obj.SidValues) > 0 {
		choices_set += 1
		choice = Ospfv2SRPrefixSidChoice.SID_VALUES
	}

	if len(obj.obj.SidIndices) > 0 {
		choices_set += 1
		choice = Ospfv2SRPrefixSidChoice.SID_INDICES
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv2SRPrefixSidChoice.SID_INDICES)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv2SRPrefixSid")
			}
		} else {
			intVal := otg.Ospfv2SRPrefixSid_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv2SRPrefixSid_Choice_Enum(intVal)
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

}
