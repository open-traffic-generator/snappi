package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRPrefixSidFlags *****
type isisSRPrefixSidFlags struct {
	validation
	obj          *otg.IsisSRPrefixSidFlags
	marshaller   marshalIsisSRPrefixSidFlags
	unMarshaller unMarshalIsisSRPrefixSidFlags
}

func NewIsisSRPrefixSidFlags() IsisSRPrefixSidFlags {
	obj := isisSRPrefixSidFlags{obj: &otg.IsisSRPrefixSidFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRPrefixSidFlags) msg() *otg.IsisSRPrefixSidFlags {
	return obj.obj
}

func (obj *isisSRPrefixSidFlags) setMsg(msg *otg.IsisSRPrefixSidFlags) IsisSRPrefixSidFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRPrefixSidFlags struct {
	obj *isisSRPrefixSidFlags
}

type marshalIsisSRPrefixSidFlags interface {
	// ToProto marshals IsisSRPrefixSidFlags to protobuf object *otg.IsisSRPrefixSidFlags
	ToProto() (*otg.IsisSRPrefixSidFlags, error)
	// ToPbText marshals IsisSRPrefixSidFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRPrefixSidFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRPrefixSidFlags to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRPrefixSidFlags struct {
	obj *isisSRPrefixSidFlags
}

type unMarshalIsisSRPrefixSidFlags interface {
	// FromProto unmarshals IsisSRPrefixSidFlags from protobuf object *otg.IsisSRPrefixSidFlags
	FromProto(msg *otg.IsisSRPrefixSidFlags) (IsisSRPrefixSidFlags, error)
	// FromPbText unmarshals IsisSRPrefixSidFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRPrefixSidFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRPrefixSidFlags from JSON text
	FromJson(value string) error
}

func (obj *isisSRPrefixSidFlags) Marshal() marshalIsisSRPrefixSidFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRPrefixSidFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRPrefixSidFlags) Unmarshal() unMarshalIsisSRPrefixSidFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRPrefixSidFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRPrefixSidFlags) ToProto() (*otg.IsisSRPrefixSidFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRPrefixSidFlags) FromProto(msg *otg.IsisSRPrefixSidFlags) (IsisSRPrefixSidFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRPrefixSidFlags) ToPbText() (string, error) {
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

func (m *unMarshalisisSRPrefixSidFlags) FromPbText(value string) error {
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

func (m *marshalisisSRPrefixSidFlags) ToYaml() (string, error) {
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

func (m *unMarshalisisSRPrefixSidFlags) FromYaml(value string) error {
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

func (m *marshalisisSRPrefixSidFlags) ToJson() (string, error) {
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

func (m *unMarshalisisSRPrefixSidFlags) FromJson(value string) error {
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

func (obj *isisSRPrefixSidFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRPrefixSidFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRPrefixSidFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRPrefixSidFlags) Clone() (IsisSRPrefixSidFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRPrefixSidFlags()
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

// IsisSRPrefixSidFlags is conatiner of 1-octet Flags associated with Prefix Segment-ID.
type IsisSRPrefixSidFlags interface {
	Validation
	// msg marshals IsisSRPrefixSidFlags to protobuf object *otg.IsisSRPrefixSidFlags
	// and doesn't set defaults
	msg() *otg.IsisSRPrefixSidFlags
	// setMsg unmarshals IsisSRPrefixSidFlags from protobuf object *otg.IsisSRPrefixSidFlags
	// and doesn't set defaults
	setMsg(*otg.IsisSRPrefixSidFlags) IsisSRPrefixSidFlags
	// provides marshal interface
	Marshal() marshalIsisSRPrefixSidFlags
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRPrefixSidFlags
	// validate validates IsisSRPrefixSidFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRPrefixSidFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RFlag returns bool, set in IsisSRPrefixSidFlags.
	RFlag() bool
	// SetRFlag assigns bool provided by user to IsisSRPrefixSidFlags
	SetRFlag(value bool) IsisSRPrefixSidFlags
	// HasRFlag checks if RFlag has been set in IsisSRPrefixSidFlags
	HasRFlag() bool
	// NFlag returns bool, set in IsisSRPrefixSidFlags.
	NFlag() bool
	// SetNFlag assigns bool provided by user to IsisSRPrefixSidFlags
	SetNFlag(value bool) IsisSRPrefixSidFlags
	// HasNFlag checks if NFlag has been set in IsisSRPrefixSidFlags
	HasNFlag() bool
	// PFlag returns bool, set in IsisSRPrefixSidFlags.
	PFlag() bool
	// SetPFlag assigns bool provided by user to IsisSRPrefixSidFlags
	SetPFlag(value bool) IsisSRPrefixSidFlags
	// HasPFlag checks if PFlag has been set in IsisSRPrefixSidFlags
	HasPFlag() bool
	// EFlag returns bool, set in IsisSRPrefixSidFlags.
	EFlag() bool
	// SetEFlag assigns bool provided by user to IsisSRPrefixSidFlags
	SetEFlag(value bool) IsisSRPrefixSidFlags
	// HasEFlag checks if EFlag has been set in IsisSRPrefixSidFlags
	HasEFlag() bool
	// LFlag returns bool, set in IsisSRPrefixSidFlags.
	LFlag() bool
	// SetLFlag assigns bool provided by user to IsisSRPrefixSidFlags
	SetLFlag(value bool) IsisSRPrefixSidFlags
	// HasLFlag checks if LFlag has been set in IsisSRPrefixSidFlags
	HasLFlag() bool
}

// R-Flag: Re-advertisement Flag.
// If set, then the prefix to which this Prefix-SID is attached has been propagated by the router
// from either another level (i.e., from Level-1 to Level-2 or the opposite) or redistribution (e.g., from another protocol).
// RFlag returns a bool
func (obj *isisSRPrefixSidFlags) RFlag() bool {

	return *obj.obj.RFlag

}

// R-Flag: Re-advertisement Flag.
// If set, then the prefix to which this Prefix-SID is attached has been propagated by the router
// from either another level (i.e., from Level-1 to Level-2 or the opposite) or redistribution (e.g., from another protocol).
// RFlag returns a bool
func (obj *isisSRPrefixSidFlags) HasRFlag() bool {
	return obj.obj.RFlag != nil
}

// R-Flag: Re-advertisement Flag.
// If set, then the prefix to which this Prefix-SID is attached has been propagated by the router
// from either another level (i.e., from Level-1 to Level-2 or the opposite) or redistribution (e.g., from another protocol).
// SetRFlag sets the bool value in the IsisSRPrefixSidFlags object
func (obj *isisSRPrefixSidFlags) SetRFlag(value bool) IsisSRPrefixSidFlags {

	obj.obj.RFlag = &value
	return obj
}

// N-Flag: Node-SID flag.
// If set, then the Prefix-SID refers to the router identified by the prefix.
// Typically, the N-Flag is set on Prefix-SIDs that are attached to a router loopback address.
// The N-Flag is set when the Prefix-SID is a Node-SID as described in [RFC8402].
// NFlag returns a bool
func (obj *isisSRPrefixSidFlags) NFlag() bool {

	return *obj.obj.NFlag

}

// N-Flag: Node-SID flag.
// If set, then the Prefix-SID refers to the router identified by the prefix.
// Typically, the N-Flag is set on Prefix-SIDs that are attached to a router loopback address.
// The N-Flag is set when the Prefix-SID is a Node-SID as described in [RFC8402].
// NFlag returns a bool
func (obj *isisSRPrefixSidFlags) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// N-Flag: Node-SID flag.
// If set, then the Prefix-SID refers to the router identified by the prefix.
// Typically, the N-Flag is set on Prefix-SIDs that are attached to a router loopback address.
// The N-Flag is set when the Prefix-SID is a Node-SID as described in [RFC8402].
// SetNFlag sets the bool value in the IsisSRPrefixSidFlags object
func (obj *isisSRPrefixSidFlags) SetNFlag(value bool) IsisSRPrefixSidFlags {

	obj.obj.NFlag = &value
	return obj
}

// P-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the penultimate hop MUST NOT pop the Prefix-SID before delivering the packet to the node that advertised the Prefix-SID.
// PFlag returns a bool
func (obj *isisSRPrefixSidFlags) PFlag() bool {

	return *obj.obj.PFlag

}

// P-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the penultimate hop MUST NOT pop the Prefix-SID before delivering the packet to the node that advertised the Prefix-SID.
// PFlag returns a bool
func (obj *isisSRPrefixSidFlags) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// P-Flag: No-PHP (No Penultimate Hop-Popping) Flag.
// If set, then the penultimate hop MUST NOT pop the Prefix-SID before delivering the packet to the node that advertised the Prefix-SID.
// SetPFlag sets the bool value in the IsisSRPrefixSidFlags object
func (obj *isisSRPrefixSidFlags) SetPFlag(value bool) IsisSRPrefixSidFlags {

	obj.obj.PFlag = &value
	return obj
}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID with a Prefix-SID
// that has an Explicit NULL value (0 for IPv4 and 2 for IPv6) before forwarding the packet.
// EFlag returns a bool
func (obj *isisSRPrefixSidFlags) EFlag() bool {

	return *obj.obj.EFlag

}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID with a Prefix-SID
// that has an Explicit NULL value (0 for IPv4 and 2 for IPv6) before forwarding the packet.
// EFlag returns a bool
func (obj *isisSRPrefixSidFlags) HasEFlag() bool {
	return obj.obj.EFlag != nil
}

// E-Flag: Explicit-Null Flag.
// If set, any upstream neighbor of the Prefix-SID originator MUST replace the Prefix-SID with a Prefix-SID
// that has an Explicit NULL value (0 for IPv4 and 2 for IPv6) before forwarding the packet.
// SetEFlag sets the bool value in the IsisSRPrefixSidFlags object
func (obj *isisSRPrefixSidFlags) SetEFlag(value bool) IsisSRPrefixSidFlags {

	obj.obj.EFlag = &value
	return obj
}

// The local flag. If set, then the value/index carried by
// the Prefix-SID has local significance and the Prefix SID is from Please refer to device.isis.segment_routing.router_capability.srlb_ranges.
// LFlag returns a bool
func (obj *isisSRPrefixSidFlags) LFlag() bool {

	return *obj.obj.LFlag

}

// The local flag. If set, then the value/index carried by
// the Prefix-SID has local significance and the Prefix SID is from Please refer to device.isis.segment_routing.router_capability.srlb_ranges.
// LFlag returns a bool
func (obj *isisSRPrefixSidFlags) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// The local flag. If set, then the value/index carried by
// the Prefix-SID has local significance and the Prefix SID is from Please refer to device.isis.segment_routing.router_capability.srlb_ranges.
// SetLFlag sets the bool value in the IsisSRPrefixSidFlags object
func (obj *isisSRPrefixSidFlags) SetLFlag(value bool) IsisSRPrefixSidFlags {

	obj.obj.LFlag = &value
	return obj
}

func (obj *isisSRPrefixSidFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisSRPrefixSidFlags) setDefault() {
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

}
