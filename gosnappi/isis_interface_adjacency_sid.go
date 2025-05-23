package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisInterfaceAdjacencySid *****
type isisInterfaceAdjacencySid struct {
	validation
	obj          *otg.IsisInterfaceAdjacencySid
	marshaller   marshalIsisInterfaceAdjacencySid
	unMarshaller unMarshalIsisInterfaceAdjacencySid
}

func NewIsisInterfaceAdjacencySid() IsisInterfaceAdjacencySid {
	obj := isisInterfaceAdjacencySid{obj: &otg.IsisInterfaceAdjacencySid{}}
	obj.setDefault()
	return &obj
}

func (obj *isisInterfaceAdjacencySid) msg() *otg.IsisInterfaceAdjacencySid {
	return obj.obj
}

func (obj *isisInterfaceAdjacencySid) setMsg(msg *otg.IsisInterfaceAdjacencySid) IsisInterfaceAdjacencySid {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisInterfaceAdjacencySid struct {
	obj *isisInterfaceAdjacencySid
}

type marshalIsisInterfaceAdjacencySid interface {
	// ToProto marshals IsisInterfaceAdjacencySid to protobuf object *otg.IsisInterfaceAdjacencySid
	ToProto() (*otg.IsisInterfaceAdjacencySid, error)
	// ToPbText marshals IsisInterfaceAdjacencySid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisInterfaceAdjacencySid to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisInterfaceAdjacencySid to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisInterfaceAdjacencySid to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisInterfaceAdjacencySid struct {
	obj *isisInterfaceAdjacencySid
}

type unMarshalIsisInterfaceAdjacencySid interface {
	// FromProto unmarshals IsisInterfaceAdjacencySid from protobuf object *otg.IsisInterfaceAdjacencySid
	FromProto(msg *otg.IsisInterfaceAdjacencySid) (IsisInterfaceAdjacencySid, error)
	// FromPbText unmarshals IsisInterfaceAdjacencySid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisInterfaceAdjacencySid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisInterfaceAdjacencySid from JSON text
	FromJson(value string) error
}

func (obj *isisInterfaceAdjacencySid) Marshal() marshalIsisInterfaceAdjacencySid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisInterfaceAdjacencySid{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisInterfaceAdjacencySid) Unmarshal() unMarshalIsisInterfaceAdjacencySid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisInterfaceAdjacencySid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisInterfaceAdjacencySid) ToProto() (*otg.IsisInterfaceAdjacencySid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisInterfaceAdjacencySid) FromProto(msg *otg.IsisInterfaceAdjacencySid) (IsisInterfaceAdjacencySid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisInterfaceAdjacencySid) ToPbText() (string, error) {
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

func (m *unMarshalisisInterfaceAdjacencySid) FromPbText(value string) error {
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

func (m *marshalisisInterfaceAdjacencySid) ToYaml() (string, error) {
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

func (m *unMarshalisisInterfaceAdjacencySid) FromYaml(value string) error {
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

func (m *marshalisisInterfaceAdjacencySid) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalisisInterfaceAdjacencySid) ToJson() (string, error) {
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

func (m *unMarshalisisInterfaceAdjacencySid) FromJson(value string) error {
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

func (obj *isisInterfaceAdjacencySid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisInterfaceAdjacencySid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisInterfaceAdjacencySid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisInterfaceAdjacencySid) Clone() (IsisInterfaceAdjacencySid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisInterfaceAdjacencySid()
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

// IsisInterfaceAdjacencySid is optional container for segment routing MPLS settings.
// If the container exists then the adjacency SID (segment identifier)
// sub TLV will be part of the packet.
// Reference: https://datatracker.ietf.org/doc/html/rfc8667#name-adjacency-segment-identifie.
type IsisInterfaceAdjacencySid interface {
	Validation
	// msg marshals IsisInterfaceAdjacencySid to protobuf object *otg.IsisInterfaceAdjacencySid
	// and doesn't set defaults
	msg() *otg.IsisInterfaceAdjacencySid
	// setMsg unmarshals IsisInterfaceAdjacencySid from protobuf object *otg.IsisInterfaceAdjacencySid
	// and doesn't set defaults
	setMsg(*otg.IsisInterfaceAdjacencySid) IsisInterfaceAdjacencySid
	// provides marshal interface
	Marshal() marshalIsisInterfaceAdjacencySid
	// provides unmarshal interface
	Unmarshal() unMarshalIsisInterfaceAdjacencySid
	// validate validates IsisInterfaceAdjacencySid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisInterfaceAdjacencySid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns IsisInterfaceAdjacencySidChoiceEnum, set in IsisInterfaceAdjacencySid
	Choice() IsisInterfaceAdjacencySidChoiceEnum
	// setChoice assigns IsisInterfaceAdjacencySidChoiceEnum provided by user to IsisInterfaceAdjacencySid
	setChoice(value IsisInterfaceAdjacencySidChoiceEnum) IsisInterfaceAdjacencySid
	// HasChoice checks if Choice has been set in IsisInterfaceAdjacencySid
	HasChoice() bool
	// SidValues returns []uint32, set in IsisInterfaceAdjacencySid.
	SidValues() []uint32
	// SetSidValues assigns []uint32 provided by user to IsisInterfaceAdjacencySid
	SetSidValues(value []uint32) IsisInterfaceAdjacencySid
	// SidIndices returns []uint32, set in IsisInterfaceAdjacencySid.
	SidIndices() []uint32
	// SetSidIndices assigns []uint32 provided by user to IsisInterfaceAdjacencySid
	SetSidIndices(value []uint32) IsisInterfaceAdjacencySid
	// FFlag returns bool, set in IsisInterfaceAdjacencySid.
	FFlag() bool
	// SetFFlag assigns bool provided by user to IsisInterfaceAdjacencySid
	SetFFlag(value bool) IsisInterfaceAdjacencySid
	// HasFFlag checks if FFlag has been set in IsisInterfaceAdjacencySid
	HasFFlag() bool
	// BFlag returns bool, set in IsisInterfaceAdjacencySid.
	BFlag() bool
	// SetBFlag assigns bool provided by user to IsisInterfaceAdjacencySid
	SetBFlag(value bool) IsisInterfaceAdjacencySid
	// HasBFlag checks if BFlag has been set in IsisInterfaceAdjacencySid
	HasBFlag() bool
	// LFlag returns bool, set in IsisInterfaceAdjacencySid.
	LFlag() bool
	// SetLFlag assigns bool provided by user to IsisInterfaceAdjacencySid
	SetLFlag(value bool) IsisInterfaceAdjacencySid
	// HasLFlag checks if LFlag has been set in IsisInterfaceAdjacencySid
	HasLFlag() bool
	// SFlag returns bool, set in IsisInterfaceAdjacencySid.
	SFlag() bool
	// SetSFlag assigns bool provided by user to IsisInterfaceAdjacencySid
	SetSFlag(value bool) IsisInterfaceAdjacencySid
	// HasSFlag checks if SFlag has been set in IsisInterfaceAdjacencySid
	HasSFlag() bool
	// PFlag returns bool, set in IsisInterfaceAdjacencySid.
	PFlag() bool
	// SetPFlag assigns bool provided by user to IsisInterfaceAdjacencySid
	SetPFlag(value bool) IsisInterfaceAdjacencySid
	// HasPFlag checks if PFlag has been set in IsisInterfaceAdjacencySid
	HasPFlag() bool
	// Weight returns uint32, set in IsisInterfaceAdjacencySid.
	Weight() uint32
	// SetWeight assigns uint32 provided by user to IsisInterfaceAdjacencySid
	SetWeight(value uint32) IsisInterfaceAdjacencySid
	// HasWeight checks if Weight has been set in IsisInterfaceAdjacencySid
	HasWeight() bool
}

type IsisInterfaceAdjacencySidChoiceEnum string

// Enum of Choice on IsisInterfaceAdjacencySid
var IsisInterfaceAdjacencySidChoice = struct {
	SID_VALUES  IsisInterfaceAdjacencySidChoiceEnum
	SID_INDICES IsisInterfaceAdjacencySidChoiceEnum
}{
	SID_VALUES:  IsisInterfaceAdjacencySidChoiceEnum("sid_values"),
	SID_INDICES: IsisInterfaceAdjacencySidChoiceEnum("sid_indices"),
}

func (obj *isisInterfaceAdjacencySid) Choice() IsisInterfaceAdjacencySidChoiceEnum {
	return IsisInterfaceAdjacencySidChoiceEnum(obj.obj.Choice.Enum().String())
}

// Choice of whether Adjacency-SID carries and absolute value or a relative index to the SRGB/SRLB Ranges.
// Please refer to device.isis.segment_routing.router_capability.sr_capability.srgb_ranges for the Segment Routing Global Block (SRGB) Descriptor and
// device.isis.segment_routing.router_capability.srlb_ranges for the SR Local Block (SRLB).
// A user needs to configure at least one entry of SID value or SID index.
// If no entry is configured, then an implementation may advertise appropriate default SID Value/Index based on the choice. e.g. the first value
// from the SRGB or SRLB range.
// - sid_values: Adjacency-SID carries one or more values and then v_flag is set by the implementation.
// - sid_indices: Adjacency-SID carries one or more indices and then v_flag is unset by the implementation.
// Choice returns a string
func (obj *isisInterfaceAdjacencySid) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *isisInterfaceAdjacencySid) setChoice(value IsisInterfaceAdjacencySidChoiceEnum) IsisInterfaceAdjacencySid {
	intValue, ok := otg.IsisInterfaceAdjacencySid_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisInterfaceAdjacencySidChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisInterfaceAdjacencySid_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.SidIndices = nil
	obj.obj.SidValues = nil
	return obj
}

// The corresponding Adjacency SID as one or more absolute Values for the link.
// SidValues returns a []uint32
func (obj *isisInterfaceAdjacencySid) SidValues() []uint32 {
	if obj.obj.SidValues == nil {

		obj.setChoice(IsisInterfaceAdjacencySidChoice.SID_VALUES)

	}
	return obj.obj.SidValues
}

// The corresponding Adjacency SID as one or more absolute Values for the link.
// SetSidValues sets the []uint32 value in the IsisInterfaceAdjacencySid object
func (obj *isisInterfaceAdjacencySid) SetSidValues(value []uint32) IsisInterfaceAdjacencySid {
	obj.setChoice(IsisInterfaceAdjacencySidChoice.SID_VALUES)
	if obj.obj.SidValues == nil {
		obj.obj.SidValues = make([]uint32, 0)
	}
	obj.obj.SidValues = value

	return obj
}

// One or more adjacency Indices are relative to ranges defined for SRGB or SRLB.
// SidIndices returns a []uint32
func (obj *isisInterfaceAdjacencySid) SidIndices() []uint32 {
	if obj.obj.SidIndices == nil {

		obj.setChoice(IsisInterfaceAdjacencySidChoice.SID_INDICES)

	}
	return obj.obj.SidIndices
}

// One or more adjacency Indices are relative to ranges defined for SRGB or SRLB.
// SetSidIndices sets the []uint32 value in the IsisInterfaceAdjacencySid object
func (obj *isisInterfaceAdjacencySid) SetSidIndices(value []uint32) IsisInterfaceAdjacencySid {
	obj.setChoice(IsisInterfaceAdjacencySidChoice.SID_INDICES)
	if obj.obj.SidIndices == nil {
		obj.obj.SidIndices = make([]uint32, 0)
	}
	obj.obj.SidIndices = value

	return obj
}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// FFlag returns a bool
func (obj *isisInterfaceAdjacencySid) FFlag() bool {

	return *obj.obj.FFlag

}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// FFlag returns a bool
func (obj *isisInterfaceAdjacencySid) HasFFlag() bool {
	return obj.obj.FFlag != nil
}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// SetFFlag sets the bool value in the IsisInterfaceAdjacencySid object
func (obj *isisInterfaceAdjacencySid) SetFFlag(value bool) IsisInterfaceAdjacencySid {

	obj.obj.FFlag = &value
	return obj
}

// The backup flag. If set, the Adj-SID is eligible for protection.
// BFlag returns a bool
func (obj *isisInterfaceAdjacencySid) BFlag() bool {

	return *obj.obj.BFlag

}

// The backup flag. If set, the Adj-SID is eligible for protection.
// BFlag returns a bool
func (obj *isisInterfaceAdjacencySid) HasBFlag() bool {
	return obj.obj.BFlag != nil
}

// The backup flag. If set, the Adj-SID is eligible for protection.
// SetBFlag sets the bool value in the IsisInterfaceAdjacencySid object
func (obj *isisInterfaceAdjacencySid) SetBFlag(value bool) IsisInterfaceAdjacencySid {

	obj.obj.BFlag = &value
	return obj
}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance. In this case, Adjacency_sid is from device.isis.segment_routing.router_capability.srlb_ranges.
// LFlag returns a bool
func (obj *isisInterfaceAdjacencySid) LFlag() bool {

	return *obj.obj.LFlag

}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance. In this case, Adjacency_sid is from device.isis.segment_routing.router_capability.srlb_ranges.
// LFlag returns a bool
func (obj *isisInterfaceAdjacencySid) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance. In this case, Adjacency_sid is from device.isis.segment_routing.router_capability.srlb_ranges.
// SetLFlag sets the bool value in the IsisInterfaceAdjacencySid object
func (obj *isisInterfaceAdjacencySid) SetLFlag(value bool) IsisInterfaceAdjacencySid {

	obj.obj.LFlag = &value
	return obj
}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SFlag returns a bool
func (obj *isisInterfaceAdjacencySid) SFlag() bool {

	return *obj.obj.SFlag

}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SFlag returns a bool
func (obj *isisInterfaceAdjacencySid) HasSFlag() bool {
	return obj.obj.SFlag != nil
}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SetSFlag sets the bool value in the IsisInterfaceAdjacencySid object
func (obj *isisInterfaceAdjacencySid) SetSFlag(value bool) IsisInterfaceAdjacencySid {

	obj.obj.SFlag = &value
	return obj
}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// PFlag returns a bool
func (obj *isisInterfaceAdjacencySid) PFlag() bool {

	return *obj.obj.PFlag

}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// PFlag returns a bool
func (obj *isisInterfaceAdjacencySid) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// SetPFlag sets the bool value in the IsisInterfaceAdjacencySid object
func (obj *isisInterfaceAdjacencySid) SetPFlag(value bool) IsisInterfaceAdjacencySid {

	obj.obj.PFlag = &value
	return obj
}

// The value represents the weight of the Adj-SID for the purpose of load balancing.
// Weight returns a uint32
func (obj *isisInterfaceAdjacencySid) Weight() uint32 {

	return *obj.obj.Weight

}

// The value represents the weight of the Adj-SID for the purpose of load balancing.
// Weight returns a uint32
func (obj *isisInterfaceAdjacencySid) HasWeight() bool {
	return obj.obj.Weight != nil
}

// The value represents the weight of the Adj-SID for the purpose of load balancing.
// SetWeight sets the uint32 value in the IsisInterfaceAdjacencySid object
func (obj *isisInterfaceAdjacencySid) SetWeight(value uint32) IsisInterfaceAdjacencySid {

	obj.obj.Weight = &value
	return obj
}

func (obj *isisInterfaceAdjacencySid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SidValues != nil {

		for _, item := range obj.obj.SidValues {
			if item < 16 || item > 1048575 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("16 <= IsisInterfaceAdjacencySid.SidValues <= 1048575 but Got %d", item))
			}

		}

	}

	if obj.obj.SidIndices != nil {

		for _, item := range obj.obj.SidIndices {
			if item > 4294967295 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("0 <= IsisInterfaceAdjacencySid.SidIndices <= 4294967295 but Got %d", item))
			}

		}

	}

	if obj.obj.Weight != nil {

		if *obj.obj.Weight > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisInterfaceAdjacencySid.Weight <= 255 but Got %d", *obj.obj.Weight))
		}

	}

}

func (obj *isisInterfaceAdjacencySid) setDefault() {
	var choices_set int = 0
	var choice IsisInterfaceAdjacencySidChoiceEnum

	if len(obj.obj.SidValues) > 0 {
		choices_set += 1
		choice = IsisInterfaceAdjacencySidChoice.SID_VALUES
	}

	if len(obj.obj.SidIndices) > 0 {
		choices_set += 1
		choice = IsisInterfaceAdjacencySidChoice.SID_INDICES
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(IsisInterfaceAdjacencySidChoice.SID_VALUES)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in IsisInterfaceAdjacencySid")
			}
		} else {
			intVal := otg.IsisInterfaceAdjacencySid_Choice_Enum_value[string(choice)]
			enumValue := otg.IsisInterfaceAdjacencySid_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.FFlag == nil {
		obj.SetFFlag(true)
	}
	if obj.obj.BFlag == nil {
		obj.SetBFlag(false)
	}
	if obj.obj.LFlag == nil {
		obj.SetLFlag(true)
	}
	if obj.obj.SFlag == nil {
		obj.SetSFlag(false)
	}
	if obj.obj.PFlag == nil {
		obj.SetPFlag(false)
	}
	if obj.obj.Weight == nil {
		obj.SetWeight(0)
	}

}
