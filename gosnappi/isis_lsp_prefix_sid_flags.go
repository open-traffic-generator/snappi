package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspPrefixSidFlags *****
type isisLspPrefixSidFlags struct {
	validation
	obj          *otg.IsisLspPrefixSidFlags
	marshaller   marshalIsisLspPrefixSidFlags
	unMarshaller unMarshalIsisLspPrefixSidFlags
}

func NewIsisLspPrefixSidFlags() IsisLspPrefixSidFlags {
	obj := isisLspPrefixSidFlags{obj: &otg.IsisLspPrefixSidFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspPrefixSidFlags) msg() *otg.IsisLspPrefixSidFlags {
	return obj.obj
}

func (obj *isisLspPrefixSidFlags) setMsg(msg *otg.IsisLspPrefixSidFlags) IsisLspPrefixSidFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspPrefixSidFlags struct {
	obj *isisLspPrefixSidFlags
}

type marshalIsisLspPrefixSidFlags interface {
	// ToProto marshals IsisLspPrefixSidFlags to protobuf object *otg.IsisLspPrefixSidFlags
	ToProto() (*otg.IsisLspPrefixSidFlags, error)
	// ToPbText marshals IsisLspPrefixSidFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspPrefixSidFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspPrefixSidFlags to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspPrefixSidFlags struct {
	obj *isisLspPrefixSidFlags
}

type unMarshalIsisLspPrefixSidFlags interface {
	// FromProto unmarshals IsisLspPrefixSidFlags from protobuf object *otg.IsisLspPrefixSidFlags
	FromProto(msg *otg.IsisLspPrefixSidFlags) (IsisLspPrefixSidFlags, error)
	// FromPbText unmarshals IsisLspPrefixSidFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspPrefixSidFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspPrefixSidFlags from JSON text
	FromJson(value string) error
}

func (obj *isisLspPrefixSidFlags) Marshal() marshalIsisLspPrefixSidFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspPrefixSidFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspPrefixSidFlags) Unmarshal() unMarshalIsisLspPrefixSidFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspPrefixSidFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspPrefixSidFlags) ToProto() (*otg.IsisLspPrefixSidFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspPrefixSidFlags) FromProto(msg *otg.IsisLspPrefixSidFlags) (IsisLspPrefixSidFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspPrefixSidFlags) ToPbText() (string, error) {
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

func (m *unMarshalisisLspPrefixSidFlags) FromPbText(value string) error {
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

func (m *marshalisisLspPrefixSidFlags) ToYaml() (string, error) {
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

func (m *unMarshalisisLspPrefixSidFlags) FromYaml(value string) error {
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

func (m *marshalisisLspPrefixSidFlags) ToJson() (string, error) {
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

func (m *unMarshalisisLspPrefixSidFlags) FromJson(value string) error {
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

func (obj *isisLspPrefixSidFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspPrefixSidFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspPrefixSidFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspPrefixSidFlags) Clone() (IsisLspPrefixSidFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspPrefixSidFlags()
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

// IsisLspPrefixSidFlags is conatiner of 1-octet Flags associated with Prefix Segment-ID.
type IsisLspPrefixSidFlags interface {
	Validation
	// msg marshals IsisLspPrefixSidFlags to protobuf object *otg.IsisLspPrefixSidFlags
	// and doesn't set defaults
	msg() *otg.IsisLspPrefixSidFlags
	// setMsg unmarshals IsisLspPrefixSidFlags from protobuf object *otg.IsisLspPrefixSidFlags
	// and doesn't set defaults
	setMsg(*otg.IsisLspPrefixSidFlags) IsisLspPrefixSidFlags
	// provides marshal interface
	Marshal() marshalIsisLspPrefixSidFlags
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspPrefixSidFlags
	// validate validates IsisLspPrefixSidFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspPrefixSidFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RFlag returns bool, set in IsisLspPrefixSidFlags.
	RFlag() bool
	// SetRFlag assigns bool provided by user to IsisLspPrefixSidFlags
	SetRFlag(value bool) IsisLspPrefixSidFlags
	// HasRFlag checks if RFlag has been set in IsisLspPrefixSidFlags
	HasRFlag() bool
	// NFlag returns bool, set in IsisLspPrefixSidFlags.
	NFlag() bool
	// SetNFlag assigns bool provided by user to IsisLspPrefixSidFlags
	SetNFlag(value bool) IsisLspPrefixSidFlags
	// HasNFlag checks if NFlag has been set in IsisLspPrefixSidFlags
	HasNFlag() bool
	// PFlag returns bool, set in IsisLspPrefixSidFlags.
	PFlag() bool
	// SetPFlag assigns bool provided by user to IsisLspPrefixSidFlags
	SetPFlag(value bool) IsisLspPrefixSidFlags
	// HasPFlag checks if PFlag has been set in IsisLspPrefixSidFlags
	HasPFlag() bool
	// EFlag returns bool, set in IsisLspPrefixSidFlags.
	EFlag() bool
	// SetEFlag assigns bool provided by user to IsisLspPrefixSidFlags
	SetEFlag(value bool) IsisLspPrefixSidFlags
	// HasEFlag checks if EFlag has been set in IsisLspPrefixSidFlags
	HasEFlag() bool
	// VFlag returns bool, set in IsisLspPrefixSidFlags.
	VFlag() bool
	// SetVFlag assigns bool provided by user to IsisLspPrefixSidFlags
	SetVFlag(value bool) IsisLspPrefixSidFlags
	// HasVFlag checks if VFlag has been set in IsisLspPrefixSidFlags
	HasVFlag() bool
	// LFlag returns bool, set in IsisLspPrefixSidFlags.
	LFlag() bool
	// SetLFlag assigns bool provided by user to IsisLspPrefixSidFlags
	SetLFlag(value bool) IsisLspPrefixSidFlags
	// HasLFlag checks if LFlag has been set in IsisLspPrefixSidFlags
	HasLFlag() bool
}

// Readvertisment flag. When set, the prefix to which this Prefix-SID is attached, has been propagated by
// the router either from another level or from redistribution.
// RFlag returns a bool
func (obj *isisLspPrefixSidFlags) RFlag() bool {

	return *obj.obj.RFlag

}

// Readvertisment flag. When set, the prefix to which this Prefix-SID is attached, has been propagated by
// the router either from another level or from redistribution.
// RFlag returns a bool
func (obj *isisLspPrefixSidFlags) HasRFlag() bool {
	return obj.obj.RFlag != nil
}

// Readvertisment flag. When set, the prefix to which this Prefix-SID is attached, has been propagated by
// the router either from another level or from redistribution.
// SetRFlag sets the bool value in the IsisLspPrefixSidFlags object
func (obj *isisLspPrefixSidFlags) SetRFlag(value bool) IsisLspPrefixSidFlags {

	obj.obj.RFlag = &value
	return obj
}

// Node flag. When set, the Prefix-SID refers to the router identified by the prefix. Typically, the
// N-Flag is set on Prefix-SIDs attached to a router loopback address.
// NFlag returns a bool
func (obj *isisLspPrefixSidFlags) NFlag() bool {

	return *obj.obj.NFlag

}

// Node flag. When set, the Prefix-SID refers to the router identified by the prefix. Typically, the
// N-Flag is set on Prefix-SIDs attached to a router loopback address.
// NFlag returns a bool
func (obj *isisLspPrefixSidFlags) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// Node flag. When set, the Prefix-SID refers to the router identified by the prefix. Typically, the
// N-Flag is set on Prefix-SIDs attached to a router loopback address.
// SetNFlag sets the bool value in the IsisLspPrefixSidFlags object
func (obj *isisLspPrefixSidFlags) SetNFlag(value bool) IsisLspPrefixSidFlags {

	obj.obj.NFlag = &value
	return obj
}

// Penultimate-Hop-Popping flag. When set, then the penultimate hop MUST NOT pop the Prefix-SID before
// delivering the packet to the node that advertised the Prefix-SID.
// PFlag returns a bool
func (obj *isisLspPrefixSidFlags) PFlag() bool {

	return *obj.obj.PFlag

}

// Penultimate-Hop-Popping flag. When set, then the penultimate hop MUST NOT pop the Prefix-SID before
// delivering the packet to the node that advertised the Prefix-SID.
// PFlag returns a bool
func (obj *isisLspPrefixSidFlags) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// Penultimate-Hop-Popping flag. When set, then the penultimate hop MUST NOT pop the Prefix-SID before
// delivering the packet to the node that advertised the Prefix-SID.
// SetPFlag sets the bool value in the IsisLspPrefixSidFlags object
func (obj *isisLspPrefixSidFlags) SetPFlag(value bool) IsisLspPrefixSidFlags {

	obj.obj.PFlag = &value
	return obj
}

// Explicit-Null flag. When set, any upstream neighbor of the Prefix-SID originator MUST replace the
// Prefix-SID with a Prefix-SID having an Explicit-NULL value (0 for IPv4 and 2 for IPv6) before forwarding
// the packet.
// EFlag returns a bool
func (obj *isisLspPrefixSidFlags) EFlag() bool {

	return *obj.obj.EFlag

}

// Explicit-Null flag. When set, any upstream neighbor of the Prefix-SID originator MUST replace the
// Prefix-SID with a Prefix-SID having an Explicit-NULL value (0 for IPv4 and 2 for IPv6) before forwarding
// the packet.
// EFlag returns a bool
func (obj *isisLspPrefixSidFlags) HasEFlag() bool {
	return obj.obj.EFlag != nil
}

// Explicit-Null flag. When set, any upstream neighbor of the Prefix-SID originator MUST replace the
// Prefix-SID with a Prefix-SID having an Explicit-NULL value (0 for IPv4 and 2 for IPv6) before forwarding
// the packet.
// SetEFlag sets the bool value in the IsisLspPrefixSidFlags object
func (obj *isisLspPrefixSidFlags) SetEFlag(value bool) IsisLspPrefixSidFlags {

	obj.obj.EFlag = &value
	return obj
}

// Value flag. When set, the Prefix-SID carries avalue (instead of an index).
// VFlag returns a bool
func (obj *isisLspPrefixSidFlags) VFlag() bool {

	return *obj.obj.VFlag

}

// Value flag. When set, the Prefix-SID carries avalue (instead of an index).
// VFlag returns a bool
func (obj *isisLspPrefixSidFlags) HasVFlag() bool {
	return obj.obj.VFlag != nil
}

// Value flag. When set, the Prefix-SID carries avalue (instead of an index).
// SetVFlag sets the bool value in the IsisLspPrefixSidFlags object
func (obj *isisLspPrefixSidFlags) SetVFlag(value bool) IsisLspPrefixSidFlags {

	obj.obj.VFlag = &value
	return obj
}

// Local flag. When set, the value/index carried by the Prefix-SID has local significance.
// LFlag returns a bool
func (obj *isisLspPrefixSidFlags) LFlag() bool {

	return *obj.obj.LFlag

}

// Local flag. When set, the value/index carried by the Prefix-SID has local significance.
// LFlag returns a bool
func (obj *isisLspPrefixSidFlags) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// Local flag. When set, the value/index carried by the Prefix-SID has local significance.
// SetLFlag sets the bool value in the IsisLspPrefixSidFlags object
func (obj *isisLspPrefixSidFlags) SetLFlag(value bool) IsisLspPrefixSidFlags {

	obj.obj.LFlag = &value
	return obj
}

func (obj *isisLspPrefixSidFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspPrefixSidFlags) setDefault() {

}
