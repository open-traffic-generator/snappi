package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRPrefixSID *****
type isisSRPrefixSID struct {
	validation
	obj          *otg.IsisSRPrefixSID
	marshaller   marshalIsisSRPrefixSID
	unMarshaller unMarshalIsisSRPrefixSID
}

func NewIsisSRPrefixSID() IsisSRPrefixSID {
	obj := isisSRPrefixSID{obj: &otg.IsisSRPrefixSID{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRPrefixSID) msg() *otg.IsisSRPrefixSID {
	return obj.obj
}

func (obj *isisSRPrefixSID) setMsg(msg *otg.IsisSRPrefixSID) IsisSRPrefixSID {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRPrefixSID struct {
	obj *isisSRPrefixSID
}

type marshalIsisSRPrefixSID interface {
	// ToProto marshals IsisSRPrefixSID to protobuf object *otg.IsisSRPrefixSID
	ToProto() (*otg.IsisSRPrefixSID, error)
	// ToPbText marshals IsisSRPrefixSID to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRPrefixSID to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRPrefixSID to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRPrefixSID struct {
	obj *isisSRPrefixSID
}

type unMarshalIsisSRPrefixSID interface {
	// FromProto unmarshals IsisSRPrefixSID from protobuf object *otg.IsisSRPrefixSID
	FromProto(msg *otg.IsisSRPrefixSID) (IsisSRPrefixSID, error)
	// FromPbText unmarshals IsisSRPrefixSID from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRPrefixSID from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRPrefixSID from JSON text
	FromJson(value string) error
}

func (obj *isisSRPrefixSID) Marshal() marshalIsisSRPrefixSID {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRPrefixSID{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRPrefixSID) Unmarshal() unMarshalIsisSRPrefixSID {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRPrefixSID{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRPrefixSID) ToProto() (*otg.IsisSRPrefixSID, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRPrefixSID) FromProto(msg *otg.IsisSRPrefixSID) (IsisSRPrefixSID, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRPrefixSID) ToPbText() (string, error) {
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

func (m *unMarshalisisSRPrefixSID) FromPbText(value string) error {
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

func (m *marshalisisSRPrefixSID) ToYaml() (string, error) {
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

func (m *unMarshalisisSRPrefixSID) FromYaml(value string) error {
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

func (m *marshalisisSRPrefixSID) ToJson() (string, error) {
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

func (m *unMarshalisisSRPrefixSID) FromJson(value string) error {
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

func (obj *isisSRPrefixSID) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRPrefixSID) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRPrefixSID) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRPrefixSID) Clone() (IsisSRPrefixSID, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRPrefixSID()
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

// IsisSRPrefixSID is this contains the properties of IS-IS Prefix-SID and its attributes for  the extended Ipv4 and Ipv6 reachability. Refernce: https://datatracker.ietf.org/doc/html/rfc8667#name-prefix-segment-identifier-p.
type IsisSRPrefixSID interface {
	Validation
	// msg marshals IsisSRPrefixSID to protobuf object *otg.IsisSRPrefixSID
	// and doesn't set defaults
	msg() *otg.IsisSRPrefixSID
	// setMsg unmarshals IsisSRPrefixSID from protobuf object *otg.IsisSRPrefixSID
	// and doesn't set defaults
	setMsg(*otg.IsisSRPrefixSID) IsisSRPrefixSID
	// provides marshal interface
	Marshal() marshalIsisSRPrefixSID
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRPrefixSID
	// validate validates IsisSRPrefixSID
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRPrefixSID, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns IsisSRPrefixSIDChoiceEnum, set in IsisSRPrefixSID
	Choice() IsisSRPrefixSIDChoiceEnum
	// setChoice assigns IsisSRPrefixSIDChoiceEnum provided by user to IsisSRPrefixSID
	setChoice(value IsisSRPrefixSIDChoiceEnum) IsisSRPrefixSID
	// HasChoice checks if Choice has been set in IsisSRPrefixSID
	HasChoice() bool
	// SidValue returns uint32, set in IsisSRPrefixSID.
	SidValue() uint32
	// SetSidValue assigns uint32 provided by user to IsisSRPrefixSID
	SetSidValue(value uint32) IsisSRPrefixSID
	// HasSidValue checks if SidValue has been set in IsisSRPrefixSID
	HasSidValue() bool
	// SidIndex returns uint32, set in IsisSRPrefixSID.
	SidIndex() uint32
	// SetSidIndex assigns uint32 provided by user to IsisSRPrefixSID
	SetSidIndex(value uint32) IsisSRPrefixSID
	// HasSidIndex checks if SidIndex has been set in IsisSRPrefixSID
	HasSidIndex() bool
	// RFlag returns bool, set in IsisSRPrefixSID.
	RFlag() bool
	// SetRFlag assigns bool provided by user to IsisSRPrefixSID
	SetRFlag(value bool) IsisSRPrefixSID
	// HasRFlag checks if RFlag has been set in IsisSRPrefixSID
	HasRFlag() bool
	// NFlag returns bool, set in IsisSRPrefixSID.
	NFlag() bool
	// SetNFlag assigns bool provided by user to IsisSRPrefixSID
	SetNFlag(value bool) IsisSRPrefixSID
	// HasNFlag checks if NFlag has been set in IsisSRPrefixSID
	HasNFlag() bool
	// PFlag returns bool, set in IsisSRPrefixSID.
	PFlag() bool
	// SetPFlag assigns bool provided by user to IsisSRPrefixSID
	SetPFlag(value bool) IsisSRPrefixSID
	// HasPFlag checks if PFlag has been set in IsisSRPrefixSID
	HasPFlag() bool
	// EFlag returns bool, set in IsisSRPrefixSID.
	EFlag() bool
	// SetEFlag assigns bool provided by user to IsisSRPrefixSID
	SetEFlag(value bool) IsisSRPrefixSID
	// HasEFlag checks if EFlag has been set in IsisSRPrefixSID
	HasEFlag() bool
	// LFlag returns bool, set in IsisSRPrefixSID.
	LFlag() bool
	// SetLFlag assigns bool provided by user to IsisSRPrefixSID
	SetLFlag(value bool) IsisSRPrefixSID
	// HasLFlag checks if LFlag has been set in IsisSRPrefixSID
	HasLFlag() bool
	// Algorithm returns uint32, set in IsisSRPrefixSID.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to IsisSRPrefixSID
	SetAlgorithm(value uint32) IsisSRPrefixSID
	// HasAlgorithm checks if Algorithm has been set in IsisSRPrefixSID
	HasAlgorithm() bool
}

type IsisSRPrefixSIDChoiceEnum string

// Enum of Choice on IsisSRPrefixSID
var IsisSRPrefixSIDChoice = struct {
	SID_VALUE IsisSRPrefixSIDChoiceEnum
	SID_INDEX IsisSRPrefixSIDChoiceEnum
}{
	SID_VALUE: IsisSRPrefixSIDChoiceEnum("sid_value"),
	SID_INDEX: IsisSRPrefixSIDChoiceEnum("sid_index"),
}

func (obj *isisSRPrefixSID) Choice() IsisSRPrefixSIDChoiceEnum {
	return IsisSRPrefixSIDChoiceEnum(obj.obj.Choice.Enum().String())
}

// The choice of Prefix-SID carries a value instead of an index.
// Refer to "srgb_ranges" under IsisSR.Srlb under Isis.RouterCapability  or the SR Local Block (SRLB)
// and IsisSR.Srgb under Isis.SRCapability for the Segment Routing Global Block (SRGB) Descriptor.
// - sid_value: Prefix-SID carries a value
// - sid_index: Prefix-SID carries an index.
// Choice returns a string
func (obj *isisSRPrefixSID) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *isisSRPrefixSID) setChoice(value IsisSRPrefixSIDChoiceEnum) IsisSRPrefixSID {
	intValue, ok := otg.IsisSRPrefixSID_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisSRPrefixSIDChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisSRPrefixSID_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.SidIndex = nil
	obj.obj.SidValue = nil

	if value == IsisSRPrefixSIDChoice.SID_VALUE {
		defaultValue := uint32(16000)
		obj.obj.SidValue = &defaultValue
	}

	if value == IsisSRPrefixSIDChoice.SID_INDEX {
		defaultValue := uint32(0)
		obj.obj.SidIndex = &defaultValue
	}

	return obj
}

// SID/Label value that is associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SidValue returns a uint32
func (obj *isisSRPrefixSID) SidValue() uint32 {

	if obj.obj.SidValue == nil {
		obj.setChoice(IsisSRPrefixSIDChoice.SID_VALUE)
	}

	return *obj.obj.SidValue

}

// SID/Label value that is associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SidValue returns a uint32
func (obj *isisSRPrefixSID) HasSidValue() bool {
	return obj.obj.SidValue != nil
}

// SID/Label value that is associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SetSidValue sets the uint32 value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetSidValue(value uint32) IsisSRPrefixSID {
	obj.setChoice(IsisSRPrefixSIDChoice.SID_VALUE)
	obj.obj.SidValue = &value
	return obj
}

// SID/Label Index that is associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SidIndex returns a uint32
func (obj *isisSRPrefixSID) SidIndex() uint32 {

	if obj.obj.SidIndex == nil {
		obj.setChoice(IsisSRPrefixSIDChoice.SID_INDEX)
	}

	return *obj.obj.SidIndex

}

// SID/Label Index that is associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SidIndex returns a uint32
func (obj *isisSRPrefixSID) HasSidIndex() bool {
	return obj.obj.SidIndex != nil
}

// SID/Label Index that is associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SetSidIndex sets the uint32 value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetSidIndex(value uint32) IsisSRPrefixSID {
	obj.setChoice(IsisSRPrefixSIDChoice.SID_INDEX)
	obj.obj.SidIndex = &value
	return obj
}

// R-Flag: Re-advertisement Flag.
// If set, then the prefix to which this Prefix-SID is attached has been propagated by the router
// from either another level (i.e., from Level-1 to Level-2 or the opposite) or redistribution (e.g., from another protocol).
// RFlag returns a bool
func (obj *isisSRPrefixSID) RFlag() bool {

	return *obj.obj.RFlag

}

// R-Flag: Re-advertisement Flag.
// If set, then the prefix to which this Prefix-SID is attached has been propagated by the router
// from either another level (i.e., from Level-1 to Level-2 or the opposite) or redistribution (e.g., from another protocol).
// RFlag returns a bool
func (obj *isisSRPrefixSID) HasRFlag() bool {
	return obj.obj.RFlag != nil
}

// R-Flag: Re-advertisement Flag.
// If set, then the prefix to which this Prefix-SID is attached has been propagated by the router
// from either another level (i.e., from Level-1 to Level-2 or the opposite) or redistribution (e.g., from another protocol).
// SetRFlag sets the bool value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetRFlag(value bool) IsisSRPrefixSID {

	obj.obj.RFlag = &value
	return obj
}

// N-Flag: Node-SID flag.
// If set, then the Prefix-SID refers to the router identified by the prefix.
// Typically, the N-Flag is set on Prefix-SIDs that are attached to a router loopback address.
// The N-Flag is set when the Prefix-SID is a Node-SID as described in [RFC8402].
// NFlag returns a bool
func (obj *isisSRPrefixSID) NFlag() bool {

	return *obj.obj.NFlag

}

// N-Flag: Node-SID flag.
// If set, then the Prefix-SID refers to the router identified by the prefix.
// Typically, the N-Flag is set on Prefix-SIDs that are attached to a router loopback address.
// The N-Flag is set when the Prefix-SID is a Node-SID as described in [RFC8402].
// NFlag returns a bool
func (obj *isisSRPrefixSID) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// N-Flag: Node-SID flag.
// If set, then the Prefix-SID refers to the router identified by the prefix.
// Typically, the N-Flag is set on Prefix-SIDs that are attached to a router loopback address.
// The N-Flag is set when the Prefix-SID is a Node-SID as described in [RFC8402].
// SetNFlag sets the bool value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetNFlag(value bool) IsisSRPrefixSID {

	obj.obj.NFlag = &value
	return obj
}

// P-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the Prefix-SID refers to the router identified by the prefix.
// Typically, the N-Flag is set on Prefix-SIDs that are attached to a router loopback address.
// The N-Flag is set when the Prefix-SID is a Node-SID as described in [RFC8402].
// PFlag returns a bool
func (obj *isisSRPrefixSID) PFlag() bool {

	return *obj.obj.PFlag

}

// P-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the Prefix-SID refers to the router identified by the prefix.
// Typically, the N-Flag is set on Prefix-SIDs that are attached to a router loopback address.
// The N-Flag is set when the Prefix-SID is a Node-SID as described in [RFC8402].
// PFlag returns a bool
func (obj *isisSRPrefixSID) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// P-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the Prefix-SID refers to the router identified by the prefix.
// Typically, the N-Flag is set on Prefix-SIDs that are attached to a router loopback address.
// The N-Flag is set when the Prefix-SID is a Node-SID as described in [RFC8402].
// SetPFlag sets the bool value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetPFlag(value bool) IsisSRPrefixSID {

	obj.obj.PFlag = &value
	return obj
}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID with a Prefix-SID
// that has an Explicit NULL value (0 for IPv4 and 2 for IPv6) before forwarding the packet.
// EFlag returns a bool
func (obj *isisSRPrefixSID) EFlag() bool {

	return *obj.obj.EFlag

}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID with a Prefix-SID
// that has an Explicit NULL value (0 for IPv4 and 2 for IPv6) before forwarding the packet.
// EFlag returns a bool
func (obj *isisSRPrefixSID) HasEFlag() bool {
	return obj.obj.EFlag != nil
}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID with a Prefix-SID
// that has an Explicit NULL value (0 for IPv4 and 2 for IPv6) before forwarding the packet.
// SetEFlag sets the bool value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetEFlag(value bool) IsisSRPrefixSID {

	obj.obj.EFlag = &value
	return obj
}

// The local flag.  If set, then the value/index carried by
// the Prefix-SID has local significance and the Prefix SID is from "srgb_ranges" under IsisSR.Srlb under Isis.RouterCapability.
// LFlag returns a bool
func (obj *isisSRPrefixSID) LFlag() bool {

	return *obj.obj.LFlag

}

// The local flag.  If set, then the value/index carried by
// the Prefix-SID has local significance and the Prefix SID is from "srgb_ranges" under IsisSR.Srlb under Isis.RouterCapability.
// LFlag returns a bool
func (obj *isisSRPrefixSID) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// The local flag.  If set, then the value/index carried by
// the Prefix-SID has local significance and the Prefix SID is from "srgb_ranges" under IsisSR.Srlb under Isis.RouterCapability.
// SetLFlag sets the bool value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetLFlag(value bool) IsisSRPrefixSID {

	obj.obj.LFlag = &value
	return obj
}

// The Isis may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these
// nodes.
// Algorithm returns a uint32
func (obj *isisSRPrefixSID) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// The Isis may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these
// nodes.
// Algorithm returns a uint32
func (obj *isisSRPrefixSID) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// The Isis may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these
// nodes.
// SetAlgorithm sets the uint32 value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetAlgorithm(value uint32) IsisSRPrefixSID {

	obj.obj.Algorithm = &value
	return obj
}

func (obj *isisSRPrefixSID) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Algorithm != nil {

		if *obj.obj.Algorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRPrefixSID.Algorithm <= 255 but Got %d", *obj.obj.Algorithm))
		}

	}

}

func (obj *isisSRPrefixSID) setDefault() {
	var choices_set int = 0
	var choice IsisSRPrefixSIDChoiceEnum

	if obj.obj.SidValue != nil {
		choices_set += 1
		choice = IsisSRPrefixSIDChoice.SID_VALUE
	}

	if obj.obj.SidIndex != nil {
		choices_set += 1
		choice = IsisSRPrefixSIDChoice.SID_INDEX
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(IsisSRPrefixSIDChoice.SID_VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in IsisSRPrefixSID")
			}
		} else {
			intVal := otg.IsisSRPrefixSID_Choice_Enum_value[string(choice)]
			enumValue := otg.IsisSRPrefixSID_Choice_Enum(intVal)
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
