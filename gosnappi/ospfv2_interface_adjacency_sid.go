package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2InterfaceAdjacencySid *****
type ospfv2InterfaceAdjacencySid struct {
	validation
	obj          *otg.Ospfv2InterfaceAdjacencySid
	marshaller   marshalOspfv2InterfaceAdjacencySid
	unMarshaller unMarshalOspfv2InterfaceAdjacencySid
}

func NewOspfv2InterfaceAdjacencySid() Ospfv2InterfaceAdjacencySid {
	obj := ospfv2InterfaceAdjacencySid{obj: &otg.Ospfv2InterfaceAdjacencySid{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2InterfaceAdjacencySid) msg() *otg.Ospfv2InterfaceAdjacencySid {
	return obj.obj
}

func (obj *ospfv2InterfaceAdjacencySid) setMsg(msg *otg.Ospfv2InterfaceAdjacencySid) Ospfv2InterfaceAdjacencySid {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2InterfaceAdjacencySid struct {
	obj *ospfv2InterfaceAdjacencySid
}

type marshalOspfv2InterfaceAdjacencySid interface {
	// ToProto marshals Ospfv2InterfaceAdjacencySid to protobuf object *otg.Ospfv2InterfaceAdjacencySid
	ToProto() (*otg.Ospfv2InterfaceAdjacencySid, error)
	// ToPbText marshals Ospfv2InterfaceAdjacencySid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2InterfaceAdjacencySid to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2InterfaceAdjacencySid to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2InterfaceAdjacencySid struct {
	obj *ospfv2InterfaceAdjacencySid
}

type unMarshalOspfv2InterfaceAdjacencySid interface {
	// FromProto unmarshals Ospfv2InterfaceAdjacencySid from protobuf object *otg.Ospfv2InterfaceAdjacencySid
	FromProto(msg *otg.Ospfv2InterfaceAdjacencySid) (Ospfv2InterfaceAdjacencySid, error)
	// FromPbText unmarshals Ospfv2InterfaceAdjacencySid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2InterfaceAdjacencySid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2InterfaceAdjacencySid from JSON text
	FromJson(value string) error
}

func (obj *ospfv2InterfaceAdjacencySid) Marshal() marshalOspfv2InterfaceAdjacencySid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2InterfaceAdjacencySid{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2InterfaceAdjacencySid) Unmarshal() unMarshalOspfv2InterfaceAdjacencySid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2InterfaceAdjacencySid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2InterfaceAdjacencySid) ToProto() (*otg.Ospfv2InterfaceAdjacencySid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2InterfaceAdjacencySid) FromProto(msg *otg.Ospfv2InterfaceAdjacencySid) (Ospfv2InterfaceAdjacencySid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2InterfaceAdjacencySid) ToPbText() (string, error) {
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

func (m *unMarshalospfv2InterfaceAdjacencySid) FromPbText(value string) error {
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

func (m *marshalospfv2InterfaceAdjacencySid) ToYaml() (string, error) {
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

func (m *unMarshalospfv2InterfaceAdjacencySid) FromYaml(value string) error {
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

func (m *marshalospfv2InterfaceAdjacencySid) ToJson() (string, error) {
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

func (m *unMarshalospfv2InterfaceAdjacencySid) FromJson(value string) error {
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

func (obj *ospfv2InterfaceAdjacencySid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceAdjacencySid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2InterfaceAdjacencySid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2InterfaceAdjacencySid) Clone() (Ospfv2InterfaceAdjacencySid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2InterfaceAdjacencySid()
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

// Ospfv2InterfaceAdjacencySid is optional container for the OSPFv2 Adjacency SID sub-TLV.
// The Adjacency SID sub-TLV is carried inside the OSPFv2 Extended Link TLV of the
// Extended Link Opaque LSA and describes a segment associated with a specific adjacency
// (typically a point-to-point link).
// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-adjacency-sid-sub-tlv.
type Ospfv2InterfaceAdjacencySid interface {
	Validation
	// msg marshals Ospfv2InterfaceAdjacencySid to protobuf object *otg.Ospfv2InterfaceAdjacencySid
	// and doesn't set defaults
	msg() *otg.Ospfv2InterfaceAdjacencySid
	// setMsg unmarshals Ospfv2InterfaceAdjacencySid from protobuf object *otg.Ospfv2InterfaceAdjacencySid
	// and doesn't set defaults
	setMsg(*otg.Ospfv2InterfaceAdjacencySid) Ospfv2InterfaceAdjacencySid
	// provides marshal interface
	Marshal() marshalOspfv2InterfaceAdjacencySid
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2InterfaceAdjacencySid
	// validate validates Ospfv2InterfaceAdjacencySid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2InterfaceAdjacencySid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv2InterfaceAdjacencySidChoiceEnum, set in Ospfv2InterfaceAdjacencySid
	Choice() Ospfv2InterfaceAdjacencySidChoiceEnum
	// setChoice assigns Ospfv2InterfaceAdjacencySidChoiceEnum provided by user to Ospfv2InterfaceAdjacencySid
	setChoice(value Ospfv2InterfaceAdjacencySidChoiceEnum) Ospfv2InterfaceAdjacencySid
	// HasChoice checks if Choice has been set in Ospfv2InterfaceAdjacencySid
	HasChoice() bool
	// SidValues returns []uint32, set in Ospfv2InterfaceAdjacencySid.
	SidValues() []uint32
	// SetSidValues assigns []uint32 provided by user to Ospfv2InterfaceAdjacencySid
	SetSidValues(value []uint32) Ospfv2InterfaceAdjacencySid
	// SidIndices returns []uint32, set in Ospfv2InterfaceAdjacencySid.
	SidIndices() []uint32
	// SetSidIndices assigns []uint32 provided by user to Ospfv2InterfaceAdjacencySid
	SetSidIndices(value []uint32) Ospfv2InterfaceAdjacencySid
	// BFlag returns bool, set in Ospfv2InterfaceAdjacencySid.
	BFlag() bool
	// SetBFlag assigns bool provided by user to Ospfv2InterfaceAdjacencySid
	SetBFlag(value bool) Ospfv2InterfaceAdjacencySid
	// HasBFlag checks if BFlag has been set in Ospfv2InterfaceAdjacencySid
	HasBFlag() bool
	// GFlag returns bool, set in Ospfv2InterfaceAdjacencySid.
	GFlag() bool
	// SetGFlag assigns bool provided by user to Ospfv2InterfaceAdjacencySid
	SetGFlag(value bool) Ospfv2InterfaceAdjacencySid
	// HasGFlag checks if GFlag has been set in Ospfv2InterfaceAdjacencySid
	HasGFlag() bool
	// PFlag returns bool, set in Ospfv2InterfaceAdjacencySid.
	PFlag() bool
	// SetPFlag assigns bool provided by user to Ospfv2InterfaceAdjacencySid
	SetPFlag(value bool) Ospfv2InterfaceAdjacencySid
	// HasPFlag checks if PFlag has been set in Ospfv2InterfaceAdjacencySid
	HasPFlag() bool
	// Weight returns uint32, set in Ospfv2InterfaceAdjacencySid.
	Weight() uint32
	// SetWeight assigns uint32 provided by user to Ospfv2InterfaceAdjacencySid
	SetWeight(value uint32) Ospfv2InterfaceAdjacencySid
	// HasWeight checks if Weight has been set in Ospfv2InterfaceAdjacencySid
	HasWeight() bool
}

type Ospfv2InterfaceAdjacencySidChoiceEnum string

// Enum of Choice on Ospfv2InterfaceAdjacencySid
var Ospfv2InterfaceAdjacencySidChoice = struct {
	SID_VALUES  Ospfv2InterfaceAdjacencySidChoiceEnum
	SID_INDICES Ospfv2InterfaceAdjacencySidChoiceEnum
}{
	SID_VALUES:  Ospfv2InterfaceAdjacencySidChoiceEnum("sid_values"),
	SID_INDICES: Ospfv2InterfaceAdjacencySidChoiceEnum("sid_indices"),
}

func (obj *ospfv2InterfaceAdjacencySid) Choice() Ospfv2InterfaceAdjacencySidChoiceEnum {
	return Ospfv2InterfaceAdjacencySidChoiceEnum(obj.obj.Choice.Enum().String())
}

// Choice of whether the Adjacency SID carries absolute values (local labels) or relative
// indices. This choice sets the V-Flag (Value/Index) and the L-Flag (Local/Global) of the
// Adjacency SID sub-TLV as follows:
// - sid_values: V-Flag and L-Flag are set (both 1). Each Adj-SID carries a 3-octet local
// label value with local significance, typically from the SRLB. Please refer to
// device.ospfv2.segment_routing.srlb_ranges.
// - sid_indices: V-Flag and L-Flag are unset (both 0). Each Adj-SID carries a 4-octet
// index that is an offset into the SRGB.
// A user needs to configure at least one entry of SID value or SID index.
// Choice returns a string
func (obj *ospfv2InterfaceAdjacencySid) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv2InterfaceAdjacencySid) setChoice(value Ospfv2InterfaceAdjacencySidChoiceEnum) Ospfv2InterfaceAdjacencySid {
	intValue, ok := otg.Ospfv2InterfaceAdjacencySid_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2InterfaceAdjacencySidChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2InterfaceAdjacencySid_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.SidIndices = nil
	obj.obj.SidValues = nil
	return obj
}

// The corresponding Adjacency SID as one or more absolute local label values for the link. Used when the choice is sid_values.
// SidValues returns a []uint32
func (obj *ospfv2InterfaceAdjacencySid) SidValues() []uint32 {
	if obj.obj.SidValues == nil {

		obj.setChoice(Ospfv2InterfaceAdjacencySidChoice.SID_VALUES)

	}
	return obj.obj.SidValues
}

// The corresponding Adjacency SID as one or more absolute local label values for the link. Used when the choice is sid_values.
// SetSidValues sets the []uint32 value in the Ospfv2InterfaceAdjacencySid object
func (obj *ospfv2InterfaceAdjacencySid) SetSidValues(value []uint32) Ospfv2InterfaceAdjacencySid {
	obj.setChoice(Ospfv2InterfaceAdjacencySidChoice.SID_VALUES)
	if obj.obj.SidValues == nil {
		obj.obj.SidValues = make([]uint32, 0)
	}
	obj.obj.SidValues = value

	return obj
}

// One or more Adjacency SID indices, relative to the SRGB. Used when the choice is sid_indices.
// SidIndices returns a []uint32
func (obj *ospfv2InterfaceAdjacencySid) SidIndices() []uint32 {
	if obj.obj.SidIndices == nil {

		obj.setChoice(Ospfv2InterfaceAdjacencySidChoice.SID_INDICES)

	}
	return obj.obj.SidIndices
}

// One or more Adjacency SID indices, relative to the SRGB. Used when the choice is sid_indices.
// SetSidIndices sets the []uint32 value in the Ospfv2InterfaceAdjacencySid object
func (obj *ospfv2InterfaceAdjacencySid) SetSidIndices(value []uint32) Ospfv2InterfaceAdjacencySid {
	obj.setChoice(Ospfv2InterfaceAdjacencySidChoice.SID_INDICES)
	if obj.obj.SidIndices == nil {
		obj.obj.SidIndices = make([]uint32, 0)
	}
	obj.obj.SidIndices = value

	return obj
}

// B-Flag: Backup Flag.
// If set, the Adjacency SID is eligible for protection, for example using
// Fast Reroute (FRR) / Loop-Free Alternate (LFA).
// BFlag returns a bool
func (obj *ospfv2InterfaceAdjacencySid) BFlag() bool {

	return *obj.obj.BFlag

}

// B-Flag: Backup Flag.
// If set, the Adjacency SID is eligible for protection, for example using
// Fast Reroute (FRR) / Loop-Free Alternate (LFA).
// BFlag returns a bool
func (obj *ospfv2InterfaceAdjacencySid) HasBFlag() bool {
	return obj.obj.BFlag != nil
}

// B-Flag: Backup Flag.
// If set, the Adjacency SID is eligible for protection, for example using
// Fast Reroute (FRR) / Loop-Free Alternate (LFA).
// SetBFlag sets the bool value in the Ospfv2InterfaceAdjacencySid object
func (obj *ospfv2InterfaceAdjacencySid) SetBFlag(value bool) Ospfv2InterfaceAdjacencySid {

	obj.obj.BFlag = &value
	return obj
}

// G-Flag: Group Flag.
// When set, the G-Flag indicates that the Adj-SID refers to a group of adjacencies and
// therefore MAY be assigned to other adjacencies as well.
// GFlag returns a bool
func (obj *ospfv2InterfaceAdjacencySid) GFlag() bool {

	return *obj.obj.GFlag

}

// G-Flag: Group Flag.
// When set, the G-Flag indicates that the Adj-SID refers to a group of adjacencies and
// therefore MAY be assigned to other adjacencies as well.
// GFlag returns a bool
func (obj *ospfv2InterfaceAdjacencySid) HasGFlag() bool {
	return obj.obj.GFlag != nil
}

// G-Flag: Group Flag.
// When set, the G-Flag indicates that the Adj-SID refers to a group of adjacencies and
// therefore MAY be assigned to other adjacencies as well.
// SetGFlag sets the bool value in the Ospfv2InterfaceAdjacencySid object
func (obj *ospfv2InterfaceAdjacencySid) SetGFlag(value bool) Ospfv2InterfaceAdjacencySid {

	obj.obj.GFlag = &value
	return obj
}

// P-Flag: Persistent Flag.
// When set, the P-Flag indicates that the Adj-SID is persistently allocated, i.e. the
// Adj-SID value remains consistent across router restart and/or interface flap.
// PFlag returns a bool
func (obj *ospfv2InterfaceAdjacencySid) PFlag() bool {

	return *obj.obj.PFlag

}

// P-Flag: Persistent Flag.
// When set, the P-Flag indicates that the Adj-SID is persistently allocated, i.e. the
// Adj-SID value remains consistent across router restart and/or interface flap.
// PFlag returns a bool
func (obj *ospfv2InterfaceAdjacencySid) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// P-Flag: Persistent Flag.
// When set, the P-Flag indicates that the Adj-SID is persistently allocated, i.e. the
// Adj-SID value remains consistent across router restart and/or interface flap.
// SetPFlag sets the bool value in the Ospfv2InterfaceAdjacencySid object
func (obj *ospfv2InterfaceAdjacencySid) SetPFlag(value bool) Ospfv2InterfaceAdjacencySid {

	obj.obj.PFlag = &value
	return obj
}

// The value represents the weight of the Adj-SID for the purpose of load balancing.
// Weight returns a uint32
func (obj *ospfv2InterfaceAdjacencySid) Weight() uint32 {

	return *obj.obj.Weight

}

// The value represents the weight of the Adj-SID for the purpose of load balancing.
// Weight returns a uint32
func (obj *ospfv2InterfaceAdjacencySid) HasWeight() bool {
	return obj.obj.Weight != nil
}

// The value represents the weight of the Adj-SID for the purpose of load balancing.
// SetWeight sets the uint32 value in the Ospfv2InterfaceAdjacencySid object
func (obj *ospfv2InterfaceAdjacencySid) SetWeight(value uint32) Ospfv2InterfaceAdjacencySid {

	obj.obj.Weight = &value
	return obj
}

func (obj *ospfv2InterfaceAdjacencySid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SidValues != nil {

		for _, item := range obj.obj.SidValues {
			if item < 16 || item > 1048575 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("16 <= Ospfv2InterfaceAdjacencySid.SidValues <= 1048575 but Got %d", item))
			}

		}

	}

	if obj.obj.SidIndices != nil {

		for _, item := range obj.obj.SidIndices {
			if item > 4294967295 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("0 <= Ospfv2InterfaceAdjacencySid.SidIndices <= 4294967295 but Got %d", item))
			}

		}

	}

	if obj.obj.Weight != nil {

		if *obj.obj.Weight > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv2InterfaceAdjacencySid.Weight <= 255 but Got %d", *obj.obj.Weight))
		}

	}

}

func (obj *ospfv2InterfaceAdjacencySid) setDefault() {
	var choices_set int = 0
	var choice Ospfv2InterfaceAdjacencySidChoiceEnum

	if len(obj.obj.SidValues) > 0 {
		choices_set += 1
		choice = Ospfv2InterfaceAdjacencySidChoice.SID_VALUES
	}

	if len(obj.obj.SidIndices) > 0 {
		choices_set += 1
		choice = Ospfv2InterfaceAdjacencySidChoice.SID_INDICES
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv2InterfaceAdjacencySidChoice.SID_VALUES)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv2InterfaceAdjacencySid")
			}
		} else {
			intVal := otg.Ospfv2InterfaceAdjacencySid_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv2InterfaceAdjacencySid_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.BFlag == nil {
		obj.SetBFlag(false)
	}
	if obj.obj.GFlag == nil {
		obj.SetGFlag(false)
	}
	if obj.obj.PFlag == nil {
		obj.SetPFlag(false)
	}
	if obj.obj.Weight == nil {
		obj.SetWeight(0)
	}

}
