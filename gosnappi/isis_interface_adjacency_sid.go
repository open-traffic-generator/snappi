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
	flagsHolder  IsisInterfaceAdjSidFlags
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
	obj.setNil()
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
	m.obj.setNil()
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
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
	m.obj.setNil()
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

func (obj *isisInterfaceAdjacencySid) setNil() {
	obj.flagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
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
	// SidValue returns uint32, set in IsisInterfaceAdjacencySid.
	SidValue() uint32
	// SetSidValue assigns uint32 provided by user to IsisInterfaceAdjacencySid
	SetSidValue(value uint32) IsisInterfaceAdjacencySid
	// HasSidValue checks if SidValue has been set in IsisInterfaceAdjacencySid
	HasSidValue() bool
	// SidIndex returns uint32, set in IsisInterfaceAdjacencySid.
	SidIndex() uint32
	// SetSidIndex assigns uint32 provided by user to IsisInterfaceAdjacencySid
	SetSidIndex(value uint32) IsisInterfaceAdjacencySid
	// HasSidIndex checks if SidIndex has been set in IsisInterfaceAdjacencySid
	HasSidIndex() bool
	// Flags returns IsisInterfaceAdjSidFlags, set in IsisInterfaceAdjacencySid.
	// IsisInterfaceAdjSidFlags is conatiner of 1-octet Flags associated with Adjacency Segment-ID.
	Flags() IsisInterfaceAdjSidFlags
	// SetFlags assigns IsisInterfaceAdjSidFlags provided by user to IsisInterfaceAdjacencySid.
	// IsisInterfaceAdjSidFlags is conatiner of 1-octet Flags associated with Adjacency Segment-ID.
	SetFlags(value IsisInterfaceAdjSidFlags) IsisInterfaceAdjacencySid
	// HasFlags checks if Flags has been set in IsisInterfaceAdjacencySid
	HasFlags() bool
	// Weight returns uint32, set in IsisInterfaceAdjacencySid.
	Weight() uint32
	// SetWeight assigns uint32 provided by user to IsisInterfaceAdjacencySid
	SetWeight(value uint32) IsisInterfaceAdjacencySid
	// HasWeight checks if Weight has been set in IsisInterfaceAdjacencySid
	HasWeight() bool
	setNil()
}

type IsisInterfaceAdjacencySidChoiceEnum string

// Enum of Choice on IsisInterfaceAdjacencySid
var IsisInterfaceAdjacencySidChoice = struct {
	SID_VALUE IsisInterfaceAdjacencySidChoiceEnum
	SID_INDEX IsisInterfaceAdjacencySidChoiceEnum
}{
	SID_VALUE: IsisInterfaceAdjacencySidChoiceEnum("sid_value"),
	SID_INDEX: IsisInterfaceAdjacencySidChoiceEnum("sid_index"),
}

func (obj *isisInterfaceAdjacencySid) Choice() IsisInterfaceAdjacencySidChoiceEnum {
	return IsisInterfaceAdjacencySidChoiceEnum(obj.obj.Choice.Enum().String())
}

// Choice of whether Adjacency-SID carries and absolute value or a relative index to the SRGB/SRLB Ranges.
// Please refer to device.isis.segment_routing.router_capability.sr_capability.srgb_ranges for the Segment Routing Global Block (SRGB) Descriptor and
// device.isis.segment_routing.router_capability.srlb_ranges for the SR Local Block (SRLB).
// - sid_value: Adjacency-SID carries a value and then v_flag is set by the implementation.
// - sid_index: Adjacency-SID carries an index and then v_flag is unset by the implementation.
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
	obj.obj.SidIndex = nil
	obj.obj.SidValue = nil

	if value == IsisInterfaceAdjacencySidChoice.SID_VALUE {
		defaultValue := uint32(1)
		obj.obj.SidValue = &defaultValue
	}

	if value == IsisInterfaceAdjacencySidChoice.SID_INDEX {
		defaultValue := uint32(1)
		obj.obj.SidIndex = &defaultValue
	}

	return obj
}

// The corresponding Adjacency SID as an absolute Value for the link that.
// SidValue returns a uint32
func (obj *isisInterfaceAdjacencySid) SidValue() uint32 {

	if obj.obj.SidValue == nil {
		obj.setChoice(IsisInterfaceAdjacencySidChoice.SID_VALUE)
	}

	return *obj.obj.SidValue

}

// The corresponding Adjacency SID as an absolute Value for the link that.
// SidValue returns a uint32
func (obj *isisInterfaceAdjacencySid) HasSidValue() bool {
	return obj.obj.SidValue != nil
}

// The corresponding Adjacency SID as an absolute Value for the link that.
// SetSidValue sets the uint32 value in the IsisInterfaceAdjacencySid object
func (obj *isisInterfaceAdjacencySid) SetSidValue(value uint32) IsisInterfaceAdjacencySid {
	obj.setChoice(IsisInterfaceAdjacencySidChoice.SID_VALUE)
	obj.obj.SidValue = &value
	return obj
}

// Adjacency Index is relative to ranges defined for SRGB or SRLB.
// SidIndex returns a uint32
func (obj *isisInterfaceAdjacencySid) SidIndex() uint32 {

	if obj.obj.SidIndex == nil {
		obj.setChoice(IsisInterfaceAdjacencySidChoice.SID_INDEX)
	}

	return *obj.obj.SidIndex

}

// Adjacency Index is relative to ranges defined for SRGB or SRLB.
// SidIndex returns a uint32
func (obj *isisInterfaceAdjacencySid) HasSidIndex() bool {
	return obj.obj.SidIndex != nil
}

// Adjacency Index is relative to ranges defined for SRGB or SRLB.
// SetSidIndex sets the uint32 value in the IsisInterfaceAdjacencySid object
func (obj *isisInterfaceAdjacencySid) SetSidIndex(value uint32) IsisInterfaceAdjacencySid {
	obj.setChoice(IsisInterfaceAdjacencySidChoice.SID_INDEX)
	obj.obj.SidIndex = &value
	return obj
}

// Flags associated with Adjacency Segment-ID.
// Flags returns a IsisInterfaceAdjSidFlags
func (obj *isisInterfaceAdjacencySid) Flags() IsisInterfaceAdjSidFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewIsisInterfaceAdjSidFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &isisInterfaceAdjSidFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// Flags associated with Adjacency Segment-ID.
// Flags returns a IsisInterfaceAdjSidFlags
func (obj *isisInterfaceAdjacencySid) HasFlags() bool {
	return obj.obj.Flags != nil
}

// Flags associated with Adjacency Segment-ID.
// SetFlags sets the IsisInterfaceAdjSidFlags value in the IsisInterfaceAdjacencySid object
func (obj *isisInterfaceAdjacencySid) SetFlags(value IsisInterfaceAdjSidFlags) IsisInterfaceAdjacencySid {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

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

	if obj.obj.SidValue != nil {

		if *obj.obj.SidValue < 1 || *obj.obj.SidValue > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= IsisInterfaceAdjacencySid.SidValue <= 1048575 but Got %d", *obj.obj.SidValue))
		}

	}

	if obj.obj.SidIndex != nil {

		if *obj.obj.SidIndex > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisInterfaceAdjacencySid.SidIndex <= 1048575 but Got %d", *obj.obj.SidIndex))
		}

	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
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

	if obj.obj.SidValue != nil {
		choices_set += 1
		choice = IsisInterfaceAdjacencySidChoice.SID_VALUE
	}

	if obj.obj.SidIndex != nil {
		choices_set += 1
		choice = IsisInterfaceAdjacencySidChoice.SID_INDEX
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(IsisInterfaceAdjacencySidChoice.SID_VALUE)

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

	if obj.obj.Weight == nil {
		obj.SetWeight(0)
	}

}
