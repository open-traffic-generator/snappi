package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisInterfaceAdjacencySID *****
type isisInterfaceAdjacencySID struct {
	validation
	obj          *otg.IsisInterfaceAdjacencySID
	marshaller   marshalIsisInterfaceAdjacencySID
	unMarshaller unMarshalIsisInterfaceAdjacencySID
}

func NewIsisInterfaceAdjacencySID() IsisInterfaceAdjacencySID {
	obj := isisInterfaceAdjacencySID{obj: &otg.IsisInterfaceAdjacencySID{}}
	obj.setDefault()
	return &obj
}

func (obj *isisInterfaceAdjacencySID) msg() *otg.IsisInterfaceAdjacencySID {
	return obj.obj
}

func (obj *isisInterfaceAdjacencySID) setMsg(msg *otg.IsisInterfaceAdjacencySID) IsisInterfaceAdjacencySID {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisInterfaceAdjacencySID struct {
	obj *isisInterfaceAdjacencySID
}

type marshalIsisInterfaceAdjacencySID interface {
	// ToProto marshals IsisInterfaceAdjacencySID to protobuf object *otg.IsisInterfaceAdjacencySID
	ToProto() (*otg.IsisInterfaceAdjacencySID, error)
	// ToPbText marshals IsisInterfaceAdjacencySID to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisInterfaceAdjacencySID to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisInterfaceAdjacencySID to JSON text
	ToJson() (string, error)
}

type unMarshalisisInterfaceAdjacencySID struct {
	obj *isisInterfaceAdjacencySID
}

type unMarshalIsisInterfaceAdjacencySID interface {
	// FromProto unmarshals IsisInterfaceAdjacencySID from protobuf object *otg.IsisInterfaceAdjacencySID
	FromProto(msg *otg.IsisInterfaceAdjacencySID) (IsisInterfaceAdjacencySID, error)
	// FromPbText unmarshals IsisInterfaceAdjacencySID from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisInterfaceAdjacencySID from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisInterfaceAdjacencySID from JSON text
	FromJson(value string) error
}

func (obj *isisInterfaceAdjacencySID) Marshal() marshalIsisInterfaceAdjacencySID {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisInterfaceAdjacencySID{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisInterfaceAdjacencySID) Unmarshal() unMarshalIsisInterfaceAdjacencySID {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisInterfaceAdjacencySID{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisInterfaceAdjacencySID) ToProto() (*otg.IsisInterfaceAdjacencySID, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisInterfaceAdjacencySID) FromProto(msg *otg.IsisInterfaceAdjacencySID) (IsisInterfaceAdjacencySID, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisInterfaceAdjacencySID) ToPbText() (string, error) {
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

func (m *unMarshalisisInterfaceAdjacencySID) FromPbText(value string) error {
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

func (m *marshalisisInterfaceAdjacencySID) ToYaml() (string, error) {
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

func (m *unMarshalisisInterfaceAdjacencySID) FromYaml(value string) error {
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

func (m *marshalisisInterfaceAdjacencySID) ToJson() (string, error) {
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

func (m *unMarshalisisInterfaceAdjacencySID) FromJson(value string) error {
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

func (obj *isisInterfaceAdjacencySID) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisInterfaceAdjacencySID) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisInterfaceAdjacencySID) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisInterfaceAdjacencySID) Clone() (IsisInterfaceAdjacencySID, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisInterfaceAdjacencySID()
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

// IsisInterfaceAdjacencySID is optional container for segment routing MPLS settings.
// If the container exists then the adjacency SID (segment identifier)
// sub TLV will be part of the packet.
// Reference: https://datatracker.ietf.org/doc/html/rfc8667#name-adjacency-segment-identifie.
type IsisInterfaceAdjacencySID interface {
	Validation
	// msg marshals IsisInterfaceAdjacencySID to protobuf object *otg.IsisInterfaceAdjacencySID
	// and doesn't set defaults
	msg() *otg.IsisInterfaceAdjacencySID
	// setMsg unmarshals IsisInterfaceAdjacencySID from protobuf object *otg.IsisInterfaceAdjacencySID
	// and doesn't set defaults
	setMsg(*otg.IsisInterfaceAdjacencySID) IsisInterfaceAdjacencySID
	// provides marshal interface
	Marshal() marshalIsisInterfaceAdjacencySID
	// provides unmarshal interface
	Unmarshal() unMarshalIsisInterfaceAdjacencySID
	// validate validates IsisInterfaceAdjacencySID
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisInterfaceAdjacencySID, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns IsisInterfaceAdjacencySIDChoiceEnum, set in IsisInterfaceAdjacencySID
	Choice() IsisInterfaceAdjacencySIDChoiceEnum
	// setChoice assigns IsisInterfaceAdjacencySIDChoiceEnum provided by user to IsisInterfaceAdjacencySID
	setChoice(value IsisInterfaceAdjacencySIDChoiceEnum) IsisInterfaceAdjacencySID
	// HasChoice checks if Choice has been set in IsisInterfaceAdjacencySID
	HasChoice() bool
	// SidValue returns uint32, set in IsisInterfaceAdjacencySID.
	SidValue() uint32
	// SetSidValue assigns uint32 provided by user to IsisInterfaceAdjacencySID
	SetSidValue(value uint32) IsisInterfaceAdjacencySID
	// HasSidValue checks if SidValue has been set in IsisInterfaceAdjacencySID
	HasSidValue() bool
	// SidIndex returns uint32, set in IsisInterfaceAdjacencySID.
	SidIndex() uint32
	// SetSidIndex assigns uint32 provided by user to IsisInterfaceAdjacencySID
	SetSidIndex(value uint32) IsisInterfaceAdjacencySID
	// HasSidIndex checks if SidIndex has been set in IsisInterfaceAdjacencySID
	HasSidIndex() bool
	// FFlag returns bool, set in IsisInterfaceAdjacencySID.
	FFlag() bool
	// SetFFlag assigns bool provided by user to IsisInterfaceAdjacencySID
	SetFFlag(value bool) IsisInterfaceAdjacencySID
	// HasFFlag checks if FFlag has been set in IsisInterfaceAdjacencySID
	HasFFlag() bool
	// BFlag returns bool, set in IsisInterfaceAdjacencySID.
	BFlag() bool
	// SetBFlag assigns bool provided by user to IsisInterfaceAdjacencySID
	SetBFlag(value bool) IsisInterfaceAdjacencySID
	// HasBFlag checks if BFlag has been set in IsisInterfaceAdjacencySID
	HasBFlag() bool
	// LFlag returns bool, set in IsisInterfaceAdjacencySID.
	LFlag() bool
	// SetLFlag assigns bool provided by user to IsisInterfaceAdjacencySID
	SetLFlag(value bool) IsisInterfaceAdjacencySID
	// HasLFlag checks if LFlag has been set in IsisInterfaceAdjacencySID
	HasLFlag() bool
	// SFlag returns bool, set in IsisInterfaceAdjacencySID.
	SFlag() bool
	// SetSFlag assigns bool provided by user to IsisInterfaceAdjacencySID
	SetSFlag(value bool) IsisInterfaceAdjacencySID
	// HasSFlag checks if SFlag has been set in IsisInterfaceAdjacencySID
	HasSFlag() bool
	// PFlag returns bool, set in IsisInterfaceAdjacencySID.
	PFlag() bool
	// SetPFlag assigns bool provided by user to IsisInterfaceAdjacencySID
	SetPFlag(value bool) IsisInterfaceAdjacencySID
	// HasPFlag checks if PFlag has been set in IsisInterfaceAdjacencySID
	HasPFlag() bool
	// Weight returns uint32, set in IsisInterfaceAdjacencySID.
	Weight() uint32
	// SetWeight assigns uint32 provided by user to IsisInterfaceAdjacencySID
	SetWeight(value uint32) IsisInterfaceAdjacencySID
	// HasWeight checks if Weight has been set in IsisInterfaceAdjacencySID
	HasWeight() bool
}

type IsisInterfaceAdjacencySIDChoiceEnum string

// Enum of Choice on IsisInterfaceAdjacencySID
var IsisInterfaceAdjacencySIDChoice = struct {
	SID_VALUE IsisInterfaceAdjacencySIDChoiceEnum
	SID_INDEX IsisInterfaceAdjacencySIDChoiceEnum
}{
	SID_VALUE: IsisInterfaceAdjacencySIDChoiceEnum("sid_value"),
	SID_INDEX: IsisInterfaceAdjacencySIDChoiceEnum("sid_index"),
}

func (obj *isisInterfaceAdjacencySID) Choice() IsisInterfaceAdjacencySIDChoiceEnum {
	return IsisInterfaceAdjacencySIDChoiceEnum(obj.obj.Choice.Enum().String())
}

// Choice of whether Adjacency-SID carries and absolute value or a relative index to the SRGB/SRLB Ranges.
// Please refer to device.isis.segment_routing.router_capability.sr_capability.srgb_ranges for the Segment Routing Global Block (SRGB) Descriptor and
// device.isis.segment_routing.router_capability.srlb_ranges for the SR Local Block (SRLB).
// - sid_value: Adjacency-SID carries a value and then v_flag is set by the implementation.
// - sid_index: Adjacency-SID carries an index and then v_flag is unset by the implementation.
// Choice returns a string
func (obj *isisInterfaceAdjacencySID) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *isisInterfaceAdjacencySID) setChoice(value IsisInterfaceAdjacencySIDChoiceEnum) IsisInterfaceAdjacencySID {
	intValue, ok := otg.IsisInterfaceAdjacencySID_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisInterfaceAdjacencySIDChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisInterfaceAdjacencySID_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.SidIndex = nil
	obj.obj.SidValue = nil

	if value == IsisInterfaceAdjacencySIDChoice.SID_VALUE {
		defaultValue := uint32(1)
		obj.obj.SidValue = &defaultValue
	}

	if value == IsisInterfaceAdjacencySIDChoice.SID_INDEX {
		defaultValue := uint32(1)
		obj.obj.SidIndex = &defaultValue
	}

	return obj
}

// The corresponding Adjacency SID as an absolute Value for the link that.
// SidValue returns a uint32
func (obj *isisInterfaceAdjacencySID) SidValue() uint32 {

	if obj.obj.SidValue == nil {
		obj.setChoice(IsisInterfaceAdjacencySIDChoice.SID_VALUE)
	}

	return *obj.obj.SidValue

}

// The corresponding Adjacency SID as an absolute Value for the link that.
// SidValue returns a uint32
func (obj *isisInterfaceAdjacencySID) HasSidValue() bool {
	return obj.obj.SidValue != nil
}

// The corresponding Adjacency SID as an absolute Value for the link that.
// SetSidValue sets the uint32 value in the IsisInterfaceAdjacencySID object
func (obj *isisInterfaceAdjacencySID) SetSidValue(value uint32) IsisInterfaceAdjacencySID {
	obj.setChoice(IsisInterfaceAdjacencySIDChoice.SID_VALUE)
	obj.obj.SidValue = &value
	return obj
}

// Adjacency Index is relative to ranges defined for SRGB or SRLB.
// SidIndex returns a uint32
func (obj *isisInterfaceAdjacencySID) SidIndex() uint32 {

	if obj.obj.SidIndex == nil {
		obj.setChoice(IsisInterfaceAdjacencySIDChoice.SID_INDEX)
	}

	return *obj.obj.SidIndex

}

// Adjacency Index is relative to ranges defined for SRGB or SRLB.
// SidIndex returns a uint32
func (obj *isisInterfaceAdjacencySID) HasSidIndex() bool {
	return obj.obj.SidIndex != nil
}

// Adjacency Index is relative to ranges defined for SRGB or SRLB.
// SetSidIndex sets the uint32 value in the IsisInterfaceAdjacencySID object
func (obj *isisInterfaceAdjacencySID) SetSidIndex(value uint32) IsisInterfaceAdjacencySID {
	obj.setChoice(IsisInterfaceAdjacencySIDChoice.SID_INDEX)
	obj.obj.SidIndex = &value
	return obj
}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// FFlag returns a bool
func (obj *isisInterfaceAdjacencySID) FFlag() bool {

	return *obj.obj.FFlag

}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// FFlag returns a bool
func (obj *isisInterfaceAdjacencySID) HasFFlag() bool {
	return obj.obj.FFlag != nil
}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// SetFFlag sets the bool value in the IsisInterfaceAdjacencySID object
func (obj *isisInterfaceAdjacencySID) SetFFlag(value bool) IsisInterfaceAdjacencySID {

	obj.obj.FFlag = &value
	return obj
}

// The backup flag. If set, the Adj-SID is eligible for protection.
// BFlag returns a bool
func (obj *isisInterfaceAdjacencySID) BFlag() bool {

	return *obj.obj.BFlag

}

// The backup flag. If set, the Adj-SID is eligible for protection.
// BFlag returns a bool
func (obj *isisInterfaceAdjacencySID) HasBFlag() bool {
	return obj.obj.BFlag != nil
}

// The backup flag. If set, the Adj-SID is eligible for protection.
// SetBFlag sets the bool value in the IsisInterfaceAdjacencySID object
func (obj *isisInterfaceAdjacencySID) SetBFlag(value bool) IsisInterfaceAdjacencySID {

	obj.obj.BFlag = &value
	return obj
}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance. In this case, Adjacency_sid is from device.isis.segment_routing.router_capability.srlb_ranges.
// LFlag returns a bool
func (obj *isisInterfaceAdjacencySID) LFlag() bool {

	return *obj.obj.LFlag

}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance. In this case, Adjacency_sid is from device.isis.segment_routing.router_capability.srlb_ranges.
// LFlag returns a bool
func (obj *isisInterfaceAdjacencySID) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance. In this case, Adjacency_sid is from device.isis.segment_routing.router_capability.srlb_ranges.
// SetLFlag sets the bool value in the IsisInterfaceAdjacencySID object
func (obj *isisInterfaceAdjacencySID) SetLFlag(value bool) IsisInterfaceAdjacencySID {

	obj.obj.LFlag = &value
	return obj
}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SFlag returns a bool
func (obj *isisInterfaceAdjacencySID) SFlag() bool {

	return *obj.obj.SFlag

}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SFlag returns a bool
func (obj *isisInterfaceAdjacencySID) HasSFlag() bool {
	return obj.obj.SFlag != nil
}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SetSFlag sets the bool value in the IsisInterfaceAdjacencySID object
func (obj *isisInterfaceAdjacencySID) SetSFlag(value bool) IsisInterfaceAdjacencySID {

	obj.obj.SFlag = &value
	return obj
}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// PFlag returns a bool
func (obj *isisInterfaceAdjacencySID) PFlag() bool {

	return *obj.obj.PFlag

}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// PFlag returns a bool
func (obj *isisInterfaceAdjacencySID) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// SetPFlag sets the bool value in the IsisInterfaceAdjacencySID object
func (obj *isisInterfaceAdjacencySID) SetPFlag(value bool) IsisInterfaceAdjacencySID {

	obj.obj.PFlag = &value
	return obj
}

// The value represents the weight of the Adj-SID for the purpose of load balancing.
// Weight returns a uint32
func (obj *isisInterfaceAdjacencySID) Weight() uint32 {

	return *obj.obj.Weight

}

// The value represents the weight of the Adj-SID for the purpose of load balancing.
// Weight returns a uint32
func (obj *isisInterfaceAdjacencySID) HasWeight() bool {
	return obj.obj.Weight != nil
}

// The value represents the weight of the Adj-SID for the purpose of load balancing.
// SetWeight sets the uint32 value in the IsisInterfaceAdjacencySID object
func (obj *isisInterfaceAdjacencySID) SetWeight(value uint32) IsisInterfaceAdjacencySID {

	obj.obj.Weight = &value
	return obj
}

func (obj *isisInterfaceAdjacencySID) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SidValue != nil {

		if *obj.obj.SidValue < 1 || *obj.obj.SidValue > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= IsisInterfaceAdjacencySID.SidValue <= 1048575 but Got %d", *obj.obj.SidValue))
		}

	}

	if obj.obj.SidIndex != nil {

		if *obj.obj.SidIndex > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisInterfaceAdjacencySID.SidIndex <= 1048575 but Got %d", *obj.obj.SidIndex))
		}

	}

	if obj.obj.Weight != nil {

		if *obj.obj.Weight > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisInterfaceAdjacencySID.Weight <= 255 but Got %d", *obj.obj.Weight))
		}

	}

}

func (obj *isisInterfaceAdjacencySID) setDefault() {
	var choices_set int = 0
	var choice IsisInterfaceAdjacencySIDChoiceEnum

	if obj.obj.SidValue != nil {
		choices_set += 1
		choice = IsisInterfaceAdjacencySIDChoice.SID_VALUE
	}

	if obj.obj.SidIndex != nil {
		choices_set += 1
		choice = IsisInterfaceAdjacencySIDChoice.SID_INDEX
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(IsisInterfaceAdjacencySIDChoice.SID_VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in IsisInterfaceAdjacencySID")
			}
		} else {
			intVal := otg.IsisInterfaceAdjacencySID_Choice_Enum_value[string(choice)]
			enumValue := otg.IsisInterfaceAdjacencySID_Choice_Enum(intVal)
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
