package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2Options *****
type ospfv2Options struct {
	validation
	obj          *otg.Ospfv2Options
	marshaller   marshalOspfv2Options
	unMarshaller unMarshalOspfv2Options
}

func NewOspfv2Options() Ospfv2Options {
	obj := ospfv2Options{obj: &otg.Ospfv2Options{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2Options) msg() *otg.Ospfv2Options {
	return obj.obj
}

func (obj *ospfv2Options) setMsg(msg *otg.Ospfv2Options) Ospfv2Options {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2Options struct {
	obj *ospfv2Options
}

type marshalOspfv2Options interface {
	// ToProto marshals Ospfv2Options to protobuf object *otg.Ospfv2Options
	ToProto() (*otg.Ospfv2Options, error)
	// ToPbText marshals Ospfv2Options to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2Options to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2Options to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2Options struct {
	obj *ospfv2Options
}

type unMarshalOspfv2Options interface {
	// FromProto unmarshals Ospfv2Options from protobuf object *otg.Ospfv2Options
	FromProto(msg *otg.Ospfv2Options) (Ospfv2Options, error)
	// FromPbText unmarshals Ospfv2Options from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2Options from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2Options from JSON text
	FromJson(value string) error
}

func (obj *ospfv2Options) Marshal() marshalOspfv2Options {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2Options{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2Options) Unmarshal() unMarshalOspfv2Options {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2Options{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2Options) ToProto() (*otg.Ospfv2Options, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2Options) FromProto(msg *otg.Ospfv2Options) (Ospfv2Options, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2Options) ToPbText() (string, error) {
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

func (m *unMarshalospfv2Options) FromPbText(value string) error {
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

func (m *marshalospfv2Options) ToYaml() (string, error) {
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

func (m *unMarshalospfv2Options) FromYaml(value string) error {
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

func (m *marshalospfv2Options) ToJson() (string, error) {
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

func (m *unMarshalospfv2Options) FromJson(value string) error {
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

func (obj *ospfv2Options) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2Options) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2Options) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2Options) Clone() (Ospfv2Options, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2Options()
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

// Ospfv2Options is the OSPFv2 Options field is present Database Description packets and all LSAs.
// This enables OSPF routers to support (or not support) optional capabilities,
// and to communicate their capability level to other OSPF routers.
// When capabilities are exchanged in Database Description packets a
// router can choose not to forward certain LSAs to a neighbor because
// of its reduced functionality.
// Reference: A.2 The Options field: https://www.rfc-editor.org/rfc/rfc2328#page-46.
type Ospfv2Options interface {
	Validation
	// msg marshals Ospfv2Options to protobuf object *otg.Ospfv2Options
	// and doesn't set defaults
	msg() *otg.Ospfv2Options
	// setMsg unmarshals Ospfv2Options from protobuf object *otg.Ospfv2Options
	// and doesn't set defaults
	setMsg(*otg.Ospfv2Options) Ospfv2Options
	// provides marshal interface
	Marshal() marshalOspfv2Options
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2Options
	// validate validates Ospfv2Options
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2Options, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TBit returns bool, set in Ospfv2Options.
	TBit() bool
	// SetTBit assigns bool provided by user to Ospfv2Options
	SetTBit(value bool) Ospfv2Options
	// HasTBit checks if TBit has been set in Ospfv2Options
	HasTBit() bool
	// EBit returns bool, set in Ospfv2Options.
	EBit() bool
	// SetEBit assigns bool provided by user to Ospfv2Options
	SetEBit(value bool) Ospfv2Options
	// HasEBit checks if EBit has been set in Ospfv2Options
	HasEBit() bool
	// McBit returns bool, set in Ospfv2Options.
	McBit() bool
	// SetMcBit assigns bool provided by user to Ospfv2Options
	SetMcBit(value bool) Ospfv2Options
	// HasMcBit checks if McBit has been set in Ospfv2Options
	HasMcBit() bool
	// NpBit returns bool, set in Ospfv2Options.
	NpBit() bool
	// SetNpBit assigns bool provided by user to Ospfv2Options
	SetNpBit(value bool) Ospfv2Options
	// HasNpBit checks if NpBit has been set in Ospfv2Options
	HasNpBit() bool
	// EaBit returns bool, set in Ospfv2Options.
	EaBit() bool
	// SetEaBit assigns bool provided by user to Ospfv2Options
	SetEaBit(value bool) Ospfv2Options
	// HasEaBit checks if EaBit has been set in Ospfv2Options
	HasEaBit() bool
	// DcBit returns bool, set in Ospfv2Options.
	DcBit() bool
	// SetDcBit assigns bool provided by user to Ospfv2Options
	SetDcBit(value bool) Ospfv2Options
	// HasDcBit checks if DcBit has been set in Ospfv2Options
	HasDcBit() bool
	// OBit returns bool, set in Ospfv2Options.
	OBit() bool
	// SetOBit assigns bool provided by user to Ospfv2Options
	SetOBit(value bool) Ospfv2Options
	// HasOBit checks if OBit has been set in Ospfv2Options
	HasOBit() bool
	// UnusedBit returns bool, set in Ospfv2Options.
	UnusedBit() bool
	// SetUnusedBit assigns bool provided by user to Ospfv2Options
	SetUnusedBit(value bool) Ospfv2Options
	// HasUnusedBit checks if UnusedBit has been set in Ospfv2Options
	HasUnusedBit() bool
	// EnableRouterLsaBBit returns bool, set in Ospfv2Options.
	EnableRouterLsaBBit() bool
	// SetEnableRouterLsaBBit assigns bool provided by user to Ospfv2Options
	SetEnableRouterLsaBBit(value bool) Ospfv2Options
	// HasEnableRouterLsaBBit checks if EnableRouterLsaBBit has been set in Ospfv2Options
	HasEnableRouterLsaBBit() bool
	// EnableRouterLsaEBit returns bool, set in Ospfv2Options.
	EnableRouterLsaEBit() bool
	// SetEnableRouterLsaEBit assigns bool provided by user to Ospfv2Options
	SetEnableRouterLsaEBit(value bool) Ospfv2Options
	// HasEnableRouterLsaEBit checks if EnableRouterLsaEBit has been set in Ospfv2Options
	HasEnableRouterLsaEBit() bool
}

// Type of Service: 0th-bit: describes OSPFv2's TOS capability.
// TBit returns a bool
func (obj *ospfv2Options) TBit() bool {

	return *obj.obj.TBit

}

// Type of Service: 0th-bit: describes OSPFv2's TOS capability.
// TBit returns a bool
func (obj *ospfv2Options) HasTBit() bool {
	return obj.obj.TBit != nil
}

// Type of Service: 0th-bit: describes OSPFv2's TOS capability.
// SetTBit sets the bool value in the Ospfv2Options object
func (obj *ospfv2Options) SetTBit(value bool) Ospfv2Options {

	obj.obj.TBit = &value
	return obj
}

// External Capability: 1st-bit: describes the way AS-external-LSAs are flooded.
// EBit returns a bool
func (obj *ospfv2Options) EBit() bool {

	return *obj.obj.EBit

}

// External Capability: 1st-bit: describes the way AS-external-LSAs are flooded.
// EBit returns a bool
func (obj *ospfv2Options) HasEBit() bool {
	return obj.obj.EBit != nil
}

// External Capability: 1st-bit: describes the way AS-external-LSAs are flooded.
// SetEBit sets the bool value in the Ospfv2Options object
func (obj *ospfv2Options) SetEBit(value bool) Ospfv2Options {

	obj.obj.EBit = &value
	return obj
}

// Multicast Capability: 2nd-bit: describes whether IP multicast datagrams are forwarded according to  the specifications in [Ref18], rfc2328.
// McBit returns a bool
func (obj *ospfv2Options) McBit() bool {

	return *obj.obj.McBit

}

// Multicast Capability: 2nd-bit: describes whether IP multicast datagrams are forwarded according to  the specifications in [Ref18], rfc2328.
// McBit returns a bool
func (obj *ospfv2Options) HasMcBit() bool {
	return obj.obj.McBit != nil
}

// Multicast Capability: 2nd-bit: describes whether IP multicast datagrams are forwarded according to  the specifications in [Ref18], rfc2328.
// SetMcBit sets the bool value in the Ospfv2Options object
func (obj *ospfv2Options) SetMcBit(value bool) Ospfv2Options {

	obj.obj.McBit = &value
	return obj
}

// NSSA Capability: 3rd-bit: describes the handling of Type-7 LSAs, as specified in [Ref19], rfc2328.
// NpBit returns a bool
func (obj *ospfv2Options) NpBit() bool {

	return *obj.obj.NpBit

}

// NSSA Capability: 3rd-bit: describes the handling of Type-7 LSAs, as specified in [Ref19], rfc2328.
// NpBit returns a bool
func (obj *ospfv2Options) HasNpBit() bool {
	return obj.obj.NpBit != nil
}

// NSSA Capability: 3rd-bit: describes the handling of Type-7 LSAs, as specified in [Ref19], rfc2328.
// SetNpBit sets the bool value in the Ospfv2Options object
func (obj *ospfv2Options) SetNpBit(value bool) Ospfv2Options {

	obj.obj.NpBit = &value
	return obj
}

// External Attribute: 4th-bit: describes the router's willingness to receive and forward External-Attributes-LSAs,  as specified in [Ref20], rfc2328.
// EaBit returns a bool
func (obj *ospfv2Options) EaBit() bool {

	return *obj.obj.EaBit

}

// External Attribute: 4th-bit: describes the router's willingness to receive and forward External-Attributes-LSAs,  as specified in [Ref20], rfc2328.
// EaBit returns a bool
func (obj *ospfv2Options) HasEaBit() bool {
	return obj.obj.EaBit != nil
}

// External Attribute: 4th-bit: describes the router's willingness to receive and forward External-Attributes-LSAs,  as specified in [Ref20], rfc2328.
// SetEaBit sets the bool value in the Ospfv2Options object
func (obj *ospfv2Options) SetEaBit(value bool) Ospfv2Options {

	obj.obj.EaBit = &value
	return obj
}

// Demand Circuit: 4th-bit: describes the router's handling of demand circuits, as specified in [Ref21], rfc2328.
// DcBit returns a bool
func (obj *ospfv2Options) DcBit() bool {

	return *obj.obj.DcBit

}

// Demand Circuit: 4th-bit: describes the router's handling of demand circuits, as specified in [Ref21], rfc2328.
// DcBit returns a bool
func (obj *ospfv2Options) HasDcBit() bool {
	return obj.obj.DcBit != nil
}

// Demand Circuit: 4th-bit: describes the router's handling of demand circuits, as specified in [Ref21], rfc2328.
// SetDcBit sets the bool value in the Ospfv2Options object
func (obj *ospfv2Options) SetDcBit(value bool) Ospfv2Options {

	obj.obj.DcBit = &value
	return obj
}

// Opaque LSA's Forwarded: 6th-bit: describes the router's willingness to receive and forward Opaque-LSAs, rfc2370.
// OBit returns a bool
func (obj *ospfv2Options) OBit() bool {

	return *obj.obj.OBit

}

// Opaque LSA's Forwarded: 6th-bit: describes the router's willingness to receive and forward Opaque-LSAs, rfc2370.
// OBit returns a bool
func (obj *ospfv2Options) HasOBit() bool {
	return obj.obj.OBit != nil
}

// Opaque LSA's Forwarded: 6th-bit: describes the router's willingness to receive and forward Opaque-LSAs, rfc2370.
// SetOBit sets the bool value in the Ospfv2Options object
func (obj *ospfv2Options) SetOBit(value bool) Ospfv2Options {

	obj.obj.OBit = &value
	return obj
}

// Opaque LSA's Forwarded: 7th-bit: unused bit.
// UnusedBit returns a bool
func (obj *ospfv2Options) UnusedBit() bool {

	return *obj.obj.UnusedBit

}

// Opaque LSA's Forwarded: 7th-bit: unused bit.
// UnusedBit returns a bool
func (obj *ospfv2Options) HasUnusedBit() bool {
	return obj.obj.UnusedBit != nil
}

// Opaque LSA's Forwarded: 7th-bit: unused bit.
// SetUnusedBit sets the bool value in the Ospfv2Options object
func (obj *ospfv2Options) SetUnusedBit(value bool) Ospfv2Options {

	obj.obj.UnusedBit = &value
	return obj
}

// Enable to indicate that the router acts as an Area Border Router.
// EnableRouterLsaBBit returns a bool
func (obj *ospfv2Options) EnableRouterLsaBBit() bool {

	return *obj.obj.EnableRouterLsaBBit

}

// Enable to indicate that the router acts as an Area Border Router.
// EnableRouterLsaBBit returns a bool
func (obj *ospfv2Options) HasEnableRouterLsaBBit() bool {
	return obj.obj.EnableRouterLsaBBit != nil
}

// Enable to indicate that the router acts as an Area Border Router.
// SetEnableRouterLsaBBit sets the bool value in the Ospfv2Options object
func (obj *ospfv2Options) SetEnableRouterLsaBBit(value bool) Ospfv2Options {

	obj.obj.EnableRouterLsaBBit = &value
	return obj
}

// Enable to indicate that the router acts as an AS Boundary Router.
// EnableRouterLsaEBit returns a bool
func (obj *ospfv2Options) EnableRouterLsaEBit() bool {

	return *obj.obj.EnableRouterLsaEBit

}

// Enable to indicate that the router acts as an AS Boundary Router.
// EnableRouterLsaEBit returns a bool
func (obj *ospfv2Options) HasEnableRouterLsaEBit() bool {
	return obj.obj.EnableRouterLsaEBit != nil
}

// Enable to indicate that the router acts as an AS Boundary Router.
// SetEnableRouterLsaEBit sets the bool value in the Ospfv2Options object
func (obj *ospfv2Options) SetEnableRouterLsaEBit(value bool) Ospfv2Options {

	obj.obj.EnableRouterLsaEBit = &value
	return obj
}

func (obj *ospfv2Options) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2Options) setDefault() {
	if obj.obj.TBit == nil {
		obj.SetTBit(false)
	}
	if obj.obj.EBit == nil {
		obj.SetEBit(false)
	}
	if obj.obj.McBit == nil {
		obj.SetMcBit(false)
	}
	if obj.obj.NpBit == nil {
		obj.SetNpBit(false)
	}
	if obj.obj.EaBit == nil {
		obj.SetEaBit(false)
	}
	if obj.obj.DcBit == nil {
		obj.SetDcBit(false)
	}
	if obj.obj.OBit == nil {
		obj.SetOBit(false)
	}
	if obj.obj.UnusedBit == nil {
		obj.SetUnusedBit(false)
	}
	if obj.obj.EnableRouterLsaBBit == nil {
		obj.SetEnableRouterLsaBBit(false)
	}
	if obj.obj.EnableRouterLsaEBit == nil {
		obj.SetEnableRouterLsaEBit(false)
	}

}
