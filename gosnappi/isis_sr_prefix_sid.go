package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRPrefixSid *****
type isisSRPrefixSid struct {
	validation
	obj          *otg.IsisSRPrefixSid
	marshaller   marshalIsisSRPrefixSid
	unMarshaller unMarshalIsisSRPrefixSid
}

func NewIsisSRPrefixSid() IsisSRPrefixSid {
	obj := isisSRPrefixSid{obj: &otg.IsisSRPrefixSid{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRPrefixSid) msg() *otg.IsisSRPrefixSid {
	return obj.obj
}

func (obj *isisSRPrefixSid) setMsg(msg *otg.IsisSRPrefixSid) IsisSRPrefixSid {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRPrefixSid struct {
	obj *isisSRPrefixSid
}

type marshalIsisSRPrefixSid interface {
	// ToProto marshals IsisSRPrefixSid to protobuf object *otg.IsisSRPrefixSid
	ToProto() (*otg.IsisSRPrefixSid, error)
	// ToPbText marshals IsisSRPrefixSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRPrefixSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRPrefixSid to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRPrefixSid struct {
	obj *isisSRPrefixSid
}

type unMarshalIsisSRPrefixSid interface {
	// FromProto unmarshals IsisSRPrefixSid from protobuf object *otg.IsisSRPrefixSid
	FromProto(msg *otg.IsisSRPrefixSid) (IsisSRPrefixSid, error)
	// FromPbText unmarshals IsisSRPrefixSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRPrefixSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRPrefixSid from JSON text
	FromJson(value string) error
}

func (obj *isisSRPrefixSid) Marshal() marshalIsisSRPrefixSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRPrefixSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRPrefixSid) Unmarshal() unMarshalIsisSRPrefixSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRPrefixSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRPrefixSid) ToProto() (*otg.IsisSRPrefixSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRPrefixSid) FromProto(msg *otg.IsisSRPrefixSid) (IsisSRPrefixSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRPrefixSid) ToPbText() (string, error) {
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

func (m *unMarshalisisSRPrefixSid) FromPbText(value string) error {
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

func (m *marshalisisSRPrefixSid) ToYaml() (string, error) {
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

func (m *unMarshalisisSRPrefixSid) FromYaml(value string) error {
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

func (m *marshalisisSRPrefixSid) ToJson() (string, error) {
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

func (m *unMarshalisisSRPrefixSid) FromJson(value string) error {
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

func (obj *isisSRPrefixSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRPrefixSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRPrefixSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRPrefixSid) Clone() (IsisSRPrefixSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRPrefixSid()
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

// IsisSRPrefixSid is this contains the properties of IS-IS Prefix-SID and its attributes for  the extended Ipv4 and Ipv6 reachability. Reference: https://datatracker.ietf.org/doc/html/rfc8667#name-prefix-segment-identifier-p.
type IsisSRPrefixSid interface {
	Validation
	// msg marshals IsisSRPrefixSid to protobuf object *otg.IsisSRPrefixSid
	// and doesn't set defaults
	msg() *otg.IsisSRPrefixSid
	// setMsg unmarshals IsisSRPrefixSid from protobuf object *otg.IsisSRPrefixSid
	// and doesn't set defaults
	setMsg(*otg.IsisSRPrefixSid) IsisSRPrefixSid
	// provides marshal interface
	Marshal() marshalIsisSRPrefixSid
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRPrefixSid
	// validate validates IsisSRPrefixSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRPrefixSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns IsisSRPrefixSidChoiceEnum, set in IsisSRPrefixSid
	Choice() IsisSRPrefixSidChoiceEnum
	// setChoice assigns IsisSRPrefixSidChoiceEnum provided by user to IsisSRPrefixSid
	setChoice(value IsisSRPrefixSidChoiceEnum) IsisSRPrefixSid
	// HasChoice checks if Choice has been set in IsisSRPrefixSid
	HasChoice() bool
	// SidValue returns uint32, set in IsisSRPrefixSid.
	SidValue() uint32
	// SetSidValue assigns uint32 provided by user to IsisSRPrefixSid
	SetSidValue(value uint32) IsisSRPrefixSid
	// HasSidValue checks if SidValue has been set in IsisSRPrefixSid
	HasSidValue() bool
	// SidIndex returns uint32, set in IsisSRPrefixSid.
	SidIndex() uint32
	// SetSidIndex assigns uint32 provided by user to IsisSRPrefixSid
	SetSidIndex(value uint32) IsisSRPrefixSid
	// HasSidIndex checks if SidIndex has been set in IsisSRPrefixSid
	HasSidIndex() bool
	// RFlag returns bool, set in IsisSRPrefixSid.
	RFlag() bool
	// SetRFlag assigns bool provided by user to IsisSRPrefixSid
	SetRFlag(value bool) IsisSRPrefixSid
	// HasRFlag checks if RFlag has been set in IsisSRPrefixSid
	HasRFlag() bool
	// NFlag returns bool, set in IsisSRPrefixSid.
	NFlag() bool
	// SetNFlag assigns bool provided by user to IsisSRPrefixSid
	SetNFlag(value bool) IsisSRPrefixSid
	// HasNFlag checks if NFlag has been set in IsisSRPrefixSid
	HasNFlag() bool
	// PFlag returns bool, set in IsisSRPrefixSid.
	PFlag() bool
	// SetPFlag assigns bool provided by user to IsisSRPrefixSid
	SetPFlag(value bool) IsisSRPrefixSid
	// HasPFlag checks if PFlag has been set in IsisSRPrefixSid
	HasPFlag() bool
	// EFlag returns bool, set in IsisSRPrefixSid.
	EFlag() bool
	// SetEFlag assigns bool provided by user to IsisSRPrefixSid
	SetEFlag(value bool) IsisSRPrefixSid
	// HasEFlag checks if EFlag has been set in IsisSRPrefixSid
	HasEFlag() bool
	// LFlag returns bool, set in IsisSRPrefixSid.
	LFlag() bool
	// SetLFlag assigns bool provided by user to IsisSRPrefixSid
	SetLFlag(value bool) IsisSRPrefixSid
	// HasLFlag checks if LFlag has been set in IsisSRPrefixSid
	HasLFlag() bool
	// Algorithm returns uint32, set in IsisSRPrefixSid.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to IsisSRPrefixSid
	SetAlgorithm(value uint32) IsisSRPrefixSid
	// HasAlgorithm checks if Algorithm has been set in IsisSRPrefixSid
	HasAlgorithm() bool
}

type IsisSRPrefixSidChoiceEnum string

// Enum of Choice on IsisSRPrefixSid
var IsisSRPrefixSidChoice = struct {
	SID_VALUE IsisSRPrefixSidChoiceEnum
	SID_INDEX IsisSRPrefixSidChoiceEnum
}{
	SID_VALUE: IsisSRPrefixSidChoiceEnum("sid_value"),
	SID_INDEX: IsisSRPrefixSidChoiceEnum("sid_index"),
}

func (obj *isisSRPrefixSid) Choice() IsisSRPrefixSidChoiceEnum {
	return IsisSRPrefixSidChoiceEnum(obj.obj.Choice.Enum().String())
}

// Choice of whether Prefix-SID carries and absolute value or a relative index to the SRGB/SRLB Ranges.
// Please refer to device.isis.segment_routing.router_capability.sr_capability.srgb_ranges for the Segment Routing Global Block (SRGB) Descriptor and
// device.isis.segment_routing.router_capability.srlb_ranges for the SR Local Block (SRLB).
// - sid_value: Prefix-SID carries a value and then v_flag is set by the implementation.
// - sid_index: Prefix-SID carries an index and then v_flag is unset by the implementation.
// Choice returns a string
func (obj *isisSRPrefixSid) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *isisSRPrefixSid) setChoice(value IsisSRPrefixSidChoiceEnum) IsisSRPrefixSid {
	intValue, ok := otg.IsisSRPrefixSid_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisSRPrefixSidChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisSRPrefixSid_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.SidIndex = nil
	obj.obj.SidValue = nil

	if value == IsisSRPrefixSidChoice.SID_VALUE {
		defaultValue := uint32(16000)
		obj.obj.SidValue = &defaultValue
	}

	if value == IsisSRPrefixSidChoice.SID_INDEX {
		defaultValue := uint32(1)
		obj.obj.SidIndex = &defaultValue
	}

	return obj
}

// SID/Label as an absolute value that is associated with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SidValue returns a uint32
func (obj *isisSRPrefixSid) SidValue() uint32 {

	if obj.obj.SidValue == nil {
		obj.setChoice(IsisSRPrefixSidChoice.SID_VALUE)
	}

	return *obj.obj.SidValue

}

// SID/Label as an absolute value that is associated with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SidValue returns a uint32
func (obj *isisSRPrefixSid) HasSidValue() bool {
	return obj.obj.SidValue != nil
}

// SID/Label as an absolute value that is associated with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SetSidValue sets the uint32 value in the IsisSRPrefixSid object
func (obj *isisSRPrefixSid) SetSidValue(value uint32) IsisSRPrefixSid {
	obj.setChoice(IsisSRPrefixSidChoice.SID_VALUE)
	obj.obj.SidValue = &value
	return obj
}

// SID/Label Index that is associated with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SidIndex returns a uint32
func (obj *isisSRPrefixSid) SidIndex() uint32 {

	if obj.obj.SidIndex == nil {
		obj.setChoice(IsisSRPrefixSidChoice.SID_INDEX)
	}

	return *obj.obj.SidIndex

}

// SID/Label Index that is associated with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SidIndex returns a uint32
func (obj *isisSRPrefixSid) HasSidIndex() bool {
	return obj.obj.SidIndex != nil
}

// SID/Label Index that is associated with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SetSidIndex sets the uint32 value in the IsisSRPrefixSid object
func (obj *isisSRPrefixSid) SetSidIndex(value uint32) IsisSRPrefixSid {
	obj.setChoice(IsisSRPrefixSidChoice.SID_INDEX)
	obj.obj.SidIndex = &value
	return obj
}

// R-Flag: Re-advertisement Flag.
// If set, then the prefix to which this Prefix-SID is attached has been propagated by the router
// from either another level (i.e., from Level-1 to Level-2 or the opposite) or redistribution (e.g., from another protocol).
// RFlag returns a bool
func (obj *isisSRPrefixSid) RFlag() bool {

	return *obj.obj.RFlag

}

// R-Flag: Re-advertisement Flag.
// If set, then the prefix to which this Prefix-SID is attached has been propagated by the router
// from either another level (i.e., from Level-1 to Level-2 or the opposite) or redistribution (e.g., from another protocol).
// RFlag returns a bool
func (obj *isisSRPrefixSid) HasRFlag() bool {
	return obj.obj.RFlag != nil
}

// R-Flag: Re-advertisement Flag.
// If set, then the prefix to which this Prefix-SID is attached has been propagated by the router
// from either another level (i.e., from Level-1 to Level-2 or the opposite) or redistribution (e.g., from another protocol).
// SetRFlag sets the bool value in the IsisSRPrefixSid object
func (obj *isisSRPrefixSid) SetRFlag(value bool) IsisSRPrefixSid {

	obj.obj.RFlag = &value
	return obj
}

// N-Flag: Node-SID flag.
// If set, then the Prefix-SID refers to the router identified by the prefix.
// Typically, the N-Flag is set on Prefix-SIDs that are attached to a router loopback address.
// The N-Flag is set when the Prefix-SID is a Node-SID as described in [RFC8402].
// NFlag returns a bool
func (obj *isisSRPrefixSid) NFlag() bool {

	return *obj.obj.NFlag

}

// N-Flag: Node-SID flag.
// If set, then the Prefix-SID refers to the router identified by the prefix.
// Typically, the N-Flag is set on Prefix-SIDs that are attached to a router loopback address.
// The N-Flag is set when the Prefix-SID is a Node-SID as described in [RFC8402].
// NFlag returns a bool
func (obj *isisSRPrefixSid) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// N-Flag: Node-SID flag.
// If set, then the Prefix-SID refers to the router identified by the prefix.
// Typically, the N-Flag is set on Prefix-SIDs that are attached to a router loopback address.
// The N-Flag is set when the Prefix-SID is a Node-SID as described in [RFC8402].
// SetNFlag sets the bool value in the IsisSRPrefixSid object
func (obj *isisSRPrefixSid) SetNFlag(value bool) IsisSRPrefixSid {

	obj.obj.NFlag = &value
	return obj
}

// P-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the penultimate hop MUST NOT pop the Prefix-SID before delivering the packet to the node that advertised the Prefix-SID.
// PFlag returns a bool
func (obj *isisSRPrefixSid) PFlag() bool {

	return *obj.obj.PFlag

}

// P-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the penultimate hop MUST NOT pop the Prefix-SID before delivering the packet to the node that advertised the Prefix-SID.
// PFlag returns a bool
func (obj *isisSRPrefixSid) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// P-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the penultimate hop MUST NOT pop the Prefix-SID before delivering the packet to the node that advertised the Prefix-SID.
// SetPFlag sets the bool value in the IsisSRPrefixSid object
func (obj *isisSRPrefixSid) SetPFlag(value bool) IsisSRPrefixSid {

	obj.obj.PFlag = &value
	return obj
}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID with a Prefix-SID
// that has an Explicit NULL value (0 for IPv4 and 2 for IPv6) before forwarding the packet.
// EFlag returns a bool
func (obj *isisSRPrefixSid) EFlag() bool {

	return *obj.obj.EFlag

}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID with a Prefix-SID
// that has an Explicit NULL value (0 for IPv4 and 2 for IPv6) before forwarding the packet.
// EFlag returns a bool
func (obj *isisSRPrefixSid) HasEFlag() bool {
	return obj.obj.EFlag != nil
}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID with a Prefix-SID
// that has an Explicit NULL value (0 for IPv4 and 2 for IPv6) before forwarding the packet.
// SetEFlag sets the bool value in the IsisSRPrefixSid object
func (obj *isisSRPrefixSid) SetEFlag(value bool) IsisSRPrefixSid {

	obj.obj.EFlag = &value
	return obj
}

// The local flag. If set, then the value/index carried by
// the Prefix-SID has local significance and the Prefix SID is from Please refer to device.isis.segment_routing.router_capability.srlb_ranges.
// LFlag returns a bool
func (obj *isisSRPrefixSid) LFlag() bool {

	return *obj.obj.LFlag

}

// The local flag. If set, then the value/index carried by
// the Prefix-SID has local significance and the Prefix SID is from Please refer to device.isis.segment_routing.router_capability.srlb_ranges.
// LFlag returns a bool
func (obj *isisSRPrefixSid) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// The local flag. If set, then the value/index carried by
// the Prefix-SID has local significance and the Prefix SID is from Please refer to device.isis.segment_routing.router_capability.srlb_ranges.
// SetLFlag sets the bool value in the IsisSRPrefixSid object
func (obj *isisSRPrefixSid) SetLFlag(value bool) IsisSRPrefixSid {

	obj.obj.LFlag = &value
	return obj
}

// The Isis may use various algorithms when calculating reachability to other nodes or to prefixes attached to these nodes.
// Refernce: https://datatracker.ietf.org/doc/html/rfc8667#SRALGOSUBTLV
// Algorithm returns a uint32
func (obj *isisSRPrefixSid) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// The Isis may use various algorithms when calculating reachability to other nodes or to prefixes attached to these nodes.
// Refernce: https://datatracker.ietf.org/doc/html/rfc8667#SRALGOSUBTLV
// Algorithm returns a uint32
func (obj *isisSRPrefixSid) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// The Isis may use various algorithms when calculating reachability to other nodes or to prefixes attached to these nodes.
// Refernce: https://datatracker.ietf.org/doc/html/rfc8667#SRALGOSUBTLV
// SetAlgorithm sets the uint32 value in the IsisSRPrefixSid object
func (obj *isisSRPrefixSid) SetAlgorithm(value uint32) IsisSRPrefixSid {

	obj.obj.Algorithm = &value
	return obj
}

func (obj *isisSRPrefixSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SidValue != nil {

		if *obj.obj.SidValue < 1 || *obj.obj.SidValue > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= IsisSRPrefixSid.SidValue <= 1048575 but Got %d", *obj.obj.SidValue))
		}

	}

	if obj.obj.SidIndex != nil {

		if *obj.obj.SidIndex < 1 || *obj.obj.SidIndex > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= IsisSRPrefixSid.SidIndex <= 1048575 but Got %d", *obj.obj.SidIndex))
		}

	}

	if obj.obj.Algorithm != nil {

		if *obj.obj.Algorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRPrefixSid.Algorithm <= 255 but Got %d", *obj.obj.Algorithm))
		}

	}

}

func (obj *isisSRPrefixSid) setDefault() {
	var choices_set int = 0
	var choice IsisSRPrefixSidChoiceEnum

	if obj.obj.SidValue != nil {
		choices_set += 1
		choice = IsisSRPrefixSidChoice.SID_VALUE
	}

	if obj.obj.SidIndex != nil {
		choices_set += 1
		choice = IsisSRPrefixSidChoice.SID_INDEX
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(IsisSRPrefixSidChoice.SID_VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in IsisSRPrefixSid")
			}
		} else {
			intVal := otg.IsisSRPrefixSid_Choice_Enum_value[string(choice)]
			enumValue := otg.IsisSRPrefixSid_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.RFlag == nil {
		obj.SetRFlag(false)
	}
	if obj.obj.NFlag == nil {
		obj.SetNFlag(false)
	}
	if obj.obj.PFlag == nil {
		obj.SetPFlag(false)
	}
	if obj.obj.EFlag == nil {
		obj.SetEFlag(false)
	}
	if obj.obj.LFlag == nil {
		obj.SetLFlag(false)
	}
	if obj.obj.Algorithm == nil {
		obj.SetAlgorithm(0)
	}

}
