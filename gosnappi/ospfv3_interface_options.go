package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3InterfaceOptions *****
type ospfv3InterfaceOptions struct {
	validation
	obj          *otg.Ospfv3InterfaceOptions
	marshaller   marshalOspfv3InterfaceOptions
	unMarshaller unMarshalOspfv3InterfaceOptions
}

func NewOspfv3InterfaceOptions() Ospfv3InterfaceOptions {
	obj := ospfv3InterfaceOptions{obj: &otg.Ospfv3InterfaceOptions{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3InterfaceOptions) msg() *otg.Ospfv3InterfaceOptions {
	return obj.obj
}

func (obj *ospfv3InterfaceOptions) setMsg(msg *otg.Ospfv3InterfaceOptions) Ospfv3InterfaceOptions {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3InterfaceOptions struct {
	obj *ospfv3InterfaceOptions
}

type marshalOspfv3InterfaceOptions interface {
	// ToProto marshals Ospfv3InterfaceOptions to protobuf object *otg.Ospfv3InterfaceOptions
	ToProto() (*otg.Ospfv3InterfaceOptions, error)
	// ToPbText marshals Ospfv3InterfaceOptions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3InterfaceOptions to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3InterfaceOptions to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3InterfaceOptions struct {
	obj *ospfv3InterfaceOptions
}

type unMarshalOspfv3InterfaceOptions interface {
	// FromProto unmarshals Ospfv3InterfaceOptions from protobuf object *otg.Ospfv3InterfaceOptions
	FromProto(msg *otg.Ospfv3InterfaceOptions) (Ospfv3InterfaceOptions, error)
	// FromPbText unmarshals Ospfv3InterfaceOptions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3InterfaceOptions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3InterfaceOptions from JSON text
	FromJson(value string) error
}

func (obj *ospfv3InterfaceOptions) Marshal() marshalOspfv3InterfaceOptions {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3InterfaceOptions{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3InterfaceOptions) Unmarshal() unMarshalOspfv3InterfaceOptions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3InterfaceOptions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3InterfaceOptions) ToProto() (*otg.Ospfv3InterfaceOptions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3InterfaceOptions) FromProto(msg *otg.Ospfv3InterfaceOptions) (Ospfv3InterfaceOptions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3InterfaceOptions) ToPbText() (string, error) {
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

func (m *unMarshalospfv3InterfaceOptions) FromPbText(value string) error {
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

func (m *marshalospfv3InterfaceOptions) ToYaml() (string, error) {
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

func (m *unMarshalospfv3InterfaceOptions) FromYaml(value string) error {
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

func (m *marshalospfv3InterfaceOptions) ToJson() (string, error) {
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

func (m *unMarshalospfv3InterfaceOptions) FromJson(value string) error {
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

func (obj *ospfv3InterfaceOptions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceOptions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceOptions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3InterfaceOptions) Clone() (Ospfv3InterfaceOptions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3InterfaceOptions()
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

// Ospfv3InterfaceOptions is the Options field is present in OSPFv3 Hello packets, Database Description packets and all LSAs.
// The Options field enables OSPF routers to support (or not support) optional capabilities, and to
// communicate their capability level to other OSPF routers (https://datatracker.ietf.org/doc/html/rfc2740#appendix-A.2).
// When used in Hello packets, the Options field allows a router to reject a neighbor because of a capability mismatch.
type Ospfv3InterfaceOptions interface {
	Validation
	// msg marshals Ospfv3InterfaceOptions to protobuf object *otg.Ospfv3InterfaceOptions
	// and doesn't set defaults
	msg() *otg.Ospfv3InterfaceOptions
	// setMsg unmarshals Ospfv3InterfaceOptions from protobuf object *otg.Ospfv3InterfaceOptions
	// and doesn't set defaults
	setMsg(*otg.Ospfv3InterfaceOptions) Ospfv3InterfaceOptions
	// provides marshal interface
	Marshal() marshalOspfv3InterfaceOptions
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3InterfaceOptions
	// validate validates Ospfv3InterfaceOptions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3InterfaceOptions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DcBit returns bool, set in Ospfv3InterfaceOptions.
	DcBit() bool
	// SetDcBit assigns bool provided by user to Ospfv3InterfaceOptions
	SetDcBit(value bool) Ospfv3InterfaceOptions
	// HasDcBit checks if DcBit has been set in Ospfv3InterfaceOptions
	HasDcBit() bool
	// RBit returns bool, set in Ospfv3InterfaceOptions.
	RBit() bool
	// SetRBit assigns bool provided by user to Ospfv3InterfaceOptions
	SetRBit(value bool) Ospfv3InterfaceOptions
	// HasRBit checks if RBit has been set in Ospfv3InterfaceOptions
	HasRBit() bool
	// NBit returns bool, set in Ospfv3InterfaceOptions.
	NBit() bool
	// SetNBit assigns bool provided by user to Ospfv3InterfaceOptions
	SetNBit(value bool) Ospfv3InterfaceOptions
	// HasNBit checks if NBit has been set in Ospfv3InterfaceOptions
	HasNBit() bool
	// EBit returns bool, set in Ospfv3InterfaceOptions.
	EBit() bool
	// SetEBit assigns bool provided by user to Ospfv3InterfaceOptions
	SetEBit(value bool) Ospfv3InterfaceOptions
	// HasEBit checks if EBit has been set in Ospfv3InterfaceOptions
	HasEBit() bool
	// V6Bit returns bool, set in Ospfv3InterfaceOptions.
	V6Bit() bool
	// SetV6Bit assigns bool provided by user to Ospfv3InterfaceOptions
	SetV6Bit(value bool) Ospfv3InterfaceOptions
	// HasV6Bit checks if V6Bit has been set in Ospfv3InterfaceOptions
	HasV6Bit() bool
}

// Demand Circuit: This bit describes the router's handling of demand circuits.
// DcBit returns a bool
func (obj *ospfv3InterfaceOptions) DcBit() bool {

	return *obj.obj.DcBit

}

// Demand Circuit: This bit describes the router's handling of demand circuits.
// DcBit returns a bool
func (obj *ospfv3InterfaceOptions) HasDcBit() bool {
	return obj.obj.DcBit != nil
}

// Demand Circuit: This bit describes the router's handling of demand circuits.
// SetDcBit sets the bool value in the Ospfv3InterfaceOptions object
func (obj *ospfv3InterfaceOptions) SetDcBit(value bool) Ospfv3InterfaceOptions {

	obj.obj.DcBit = &value
	return obj
}

// Router: This bit indicates if the originator is an active router.
// RBit returns a bool
func (obj *ospfv3InterfaceOptions) RBit() bool {

	return *obj.obj.RBit

}

// Router: This bit indicates if the originator is an active router.
// RBit returns a bool
func (obj *ospfv3InterfaceOptions) HasRBit() bool {
	return obj.obj.RBit != nil
}

// Router: This bit indicates if the originator is an active router.
// SetRBit sets the bool value in the Ospfv3InterfaceOptions object
func (obj *ospfv3InterfaceOptions) SetRBit(value bool) Ospfv3InterfaceOptions {

	obj.obj.RBit = &value
	return obj
}

// NSSA Capability: This bit describes the handling of Type-7 LSAs, as specified in [Ref8], rfc2740.
// NBit returns a bool
func (obj *ospfv3InterfaceOptions) NBit() bool {

	return *obj.obj.NBit

}

// NSSA Capability: This bit describes the handling of Type-7 LSAs, as specified in [Ref8], rfc2740.
// NBit returns a bool
func (obj *ospfv3InterfaceOptions) HasNBit() bool {
	return obj.obj.NBit != nil
}

// NSSA Capability: This bit describes the handling of Type-7 LSAs, as specified in [Ref8], rfc2740.
// SetNBit sets the bool value in the Ospfv3InterfaceOptions object
func (obj *ospfv3InterfaceOptions) SetNBit(value bool) Ospfv3InterfaceOptions {

	obj.obj.NBit = &value
	return obj
}

// External Capability: This bit describes the router's willingness to receive and forward External-Attributes-LSAs,  as specified in [Ref1], rfc2740.
// EBit returns a bool
func (obj *ospfv3InterfaceOptions) EBit() bool {

	return *obj.obj.EBit

}

// External Capability: This bit describes the router's willingness to receive and forward External-Attributes-LSAs,  as specified in [Ref1], rfc2740.
// EBit returns a bool
func (obj *ospfv3InterfaceOptions) HasEBit() bool {
	return obj.obj.EBit != nil
}

// External Capability: This bit describes the router's willingness to receive and forward External-Attributes-LSAs,  as specified in [Ref1], rfc2740.
// SetEBit sets the bool value in the Ospfv3InterfaceOptions object
func (obj *ospfv3InterfaceOptions) SetEBit(value bool) Ospfv3InterfaceOptions {

	obj.obj.EBit = &value
	return obj
}

// V6: If set, the router/link should be included in IPv6 routing calculations.
// V6Bit returns a bool
func (obj *ospfv3InterfaceOptions) V6Bit() bool {

	return *obj.obj.V6Bit

}

// V6: If set, the router/link should be included in IPv6 routing calculations.
// V6Bit returns a bool
func (obj *ospfv3InterfaceOptions) HasV6Bit() bool {
	return obj.obj.V6Bit != nil
}

// V6: If set, the router/link should be included in IPv6 routing calculations.
// SetV6Bit sets the bool value in the Ospfv3InterfaceOptions object
func (obj *ospfv3InterfaceOptions) SetV6Bit(value bool) Ospfv3InterfaceOptions {

	obj.obj.V6Bit = &value
	return obj
}

func (obj *ospfv3InterfaceOptions) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv3InterfaceOptions) setDefault() {
	if obj.obj.DcBit == nil {
		obj.SetDcBit(false)
	}
	if obj.obj.RBit == nil {
		obj.SetRBit(true)
	}
	if obj.obj.NBit == nil {
		obj.SetNBit(false)
	}
	if obj.obj.EBit == nil {
		obj.SetEBit(true)
	}
	if obj.obj.V6Bit == nil {
		obj.SetV6Bit(true)
	}

}
