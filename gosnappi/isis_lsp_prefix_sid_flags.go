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
	// Readvertisment returns bool, set in IsisLspPrefixSidFlags.
	Readvertisment() bool
	// SetReadvertisment assigns bool provided by user to IsisLspPrefixSidFlags
	SetReadvertisment(value bool) IsisLspPrefixSidFlags
	// HasReadvertisment checks if Readvertisment has been set in IsisLspPrefixSidFlags
	HasReadvertisment() bool
	// Node returns bool, set in IsisLspPrefixSidFlags.
	Node() bool
	// SetNode assigns bool provided by user to IsisLspPrefixSidFlags
	SetNode(value bool) IsisLspPrefixSidFlags
	// HasNode checks if Node has been set in IsisLspPrefixSidFlags
	HasNode() bool
	// NoPhp returns bool, set in IsisLspPrefixSidFlags.
	NoPhp() bool
	// SetNoPhp assigns bool provided by user to IsisLspPrefixSidFlags
	SetNoPhp(value bool) IsisLspPrefixSidFlags
	// HasNoPhp checks if NoPhp has been set in IsisLspPrefixSidFlags
	HasNoPhp() bool
	// ExplicityNull returns bool, set in IsisLspPrefixSidFlags.
	ExplicityNull() bool
	// SetExplicityNull assigns bool provided by user to IsisLspPrefixSidFlags
	SetExplicityNull(value bool) IsisLspPrefixSidFlags
	// HasExplicityNull checks if ExplicityNull has been set in IsisLspPrefixSidFlags
	HasExplicityNull() bool
	// Value returns bool, set in IsisLspPrefixSidFlags.
	Value() bool
	// SetValue assigns bool provided by user to IsisLspPrefixSidFlags
	SetValue(value bool) IsisLspPrefixSidFlags
	// HasValue checks if Value has been set in IsisLspPrefixSidFlags
	HasValue() bool
	// Local returns bool, set in IsisLspPrefixSidFlags.
	Local() bool
	// SetLocal assigns bool provided by user to IsisLspPrefixSidFlags
	SetLocal(value bool) IsisLspPrefixSidFlags
	// HasLocal checks if Local has been set in IsisLspPrefixSidFlags
	HasLocal() bool
}

// Readvertisment flag. When set, the prefix to which this Prefix-SID is attached, has been propagated by
// the router either from another level or from redistribution.
// Readvertisment returns a bool
func (obj *isisLspPrefixSidFlags) Readvertisment() bool {

	return *obj.obj.Readvertisment

}

// Readvertisment flag. When set, the prefix to which this Prefix-SID is attached, has been propagated by
// the router either from another level or from redistribution.
// Readvertisment returns a bool
func (obj *isisLspPrefixSidFlags) HasReadvertisment() bool {
	return obj.obj.Readvertisment != nil
}

// Readvertisment flag. When set, the prefix to which this Prefix-SID is attached, has been propagated by
// the router either from another level or from redistribution.
// SetReadvertisment sets the bool value in the IsisLspPrefixSidFlags object
func (obj *isisLspPrefixSidFlags) SetReadvertisment(value bool) IsisLspPrefixSidFlags {

	obj.obj.Readvertisment = &value
	return obj
}

// Node flag. When set, the Prefix-SID refers to the router identified by the prefix. Typically, the
// N-Flag is set on Prefix-SIDs attached to a router loopback address.
// Node returns a bool
func (obj *isisLspPrefixSidFlags) Node() bool {

	return *obj.obj.Node

}

// Node flag. When set, the Prefix-SID refers to the router identified by the prefix. Typically, the
// N-Flag is set on Prefix-SIDs attached to a router loopback address.
// Node returns a bool
func (obj *isisLspPrefixSidFlags) HasNode() bool {
	return obj.obj.Node != nil
}

// Node flag. When set, the Prefix-SID refers to the router identified by the prefix. Typically, the
// N-Flag is set on Prefix-SIDs attached to a router loopback address.
// SetNode sets the bool value in the IsisLspPrefixSidFlags object
func (obj *isisLspPrefixSidFlags) SetNode(value bool) IsisLspPrefixSidFlags {

	obj.obj.Node = &value
	return obj
}

// Penultimate-Hop-Popping flag. When set, then the penultimate hop MUST NOT pop the Prefix-SID before
// delivering the packet to the node that advertised the Prefix-SID.
// NoPhp returns a bool
func (obj *isisLspPrefixSidFlags) NoPhp() bool {

	return *obj.obj.NoPhp

}

// Penultimate-Hop-Popping flag. When set, then the penultimate hop MUST NOT pop the Prefix-SID before
// delivering the packet to the node that advertised the Prefix-SID.
// NoPhp returns a bool
func (obj *isisLspPrefixSidFlags) HasNoPhp() bool {
	return obj.obj.NoPhp != nil
}

// Penultimate-Hop-Popping flag. When set, then the penultimate hop MUST NOT pop the Prefix-SID before
// delivering the packet to the node that advertised the Prefix-SID.
// SetNoPhp sets the bool value in the IsisLspPrefixSidFlags object
func (obj *isisLspPrefixSidFlags) SetNoPhp(value bool) IsisLspPrefixSidFlags {

	obj.obj.NoPhp = &value
	return obj
}

// Explicit-Null flag. When set, any upstream neighbor of the Prefix-SID originator MUST replace the
// Prefix-SID with a Prefix-SID having an Explicit-NULL value (0 for IPv4 and 2 for IPv6) before forwarding
// the packet.
// ExplicityNull returns a bool
func (obj *isisLspPrefixSidFlags) ExplicityNull() bool {

	return *obj.obj.ExplicityNull

}

// Explicit-Null flag. When set, any upstream neighbor of the Prefix-SID originator MUST replace the
// Prefix-SID with a Prefix-SID having an Explicit-NULL value (0 for IPv4 and 2 for IPv6) before forwarding
// the packet.
// ExplicityNull returns a bool
func (obj *isisLspPrefixSidFlags) HasExplicityNull() bool {
	return obj.obj.ExplicityNull != nil
}

// Explicit-Null flag. When set, any upstream neighbor of the Prefix-SID originator MUST replace the
// Prefix-SID with a Prefix-SID having an Explicit-NULL value (0 for IPv4 and 2 for IPv6) before forwarding
// the packet.
// SetExplicityNull sets the bool value in the IsisLspPrefixSidFlags object
func (obj *isisLspPrefixSidFlags) SetExplicityNull(value bool) IsisLspPrefixSidFlags {

	obj.obj.ExplicityNull = &value
	return obj
}

// Value flag. When set, the Prefix-SID carries avalue (instead of an index). By default the flag is UNSET.
// Value returns a bool
func (obj *isisLspPrefixSidFlags) Value() bool {

	return *obj.obj.Value

}

// Value flag. When set, the Prefix-SID carries avalue (instead of an index). By default the flag is UNSET.
// Value returns a bool
func (obj *isisLspPrefixSidFlags) HasValue() bool {
	return obj.obj.Value != nil
}

// Value flag. When set, the Prefix-SID carries avalue (instead of an index). By default the flag is UNSET.
// SetValue sets the bool value in the IsisLspPrefixSidFlags object
func (obj *isisLspPrefixSidFlags) SetValue(value bool) IsisLspPrefixSidFlags {

	obj.obj.Value = &value
	return obj
}

// Local flag. When set, the value/index carried by the Prefix-SID has local significance.
// By default the flag is UNSET.
// Local returns a bool
func (obj *isisLspPrefixSidFlags) Local() bool {

	return *obj.obj.Local

}

// Local flag. When set, the value/index carried by the Prefix-SID has local significance.
// By default the flag is UNSET.
// Local returns a bool
func (obj *isisLspPrefixSidFlags) HasLocal() bool {
	return obj.obj.Local != nil
}

// Local flag. When set, the value/index carried by the Prefix-SID has local significance.
// By default the flag is UNSET.
// SetLocal sets the bool value in the IsisLspPrefixSidFlags object
func (obj *isisLspPrefixSidFlags) SetLocal(value bool) IsisLspPrefixSidFlags {

	obj.obj.Local = &value
	return obj
}

func (obj *isisLspPrefixSidFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspPrefixSidFlags) setDefault() {

}
