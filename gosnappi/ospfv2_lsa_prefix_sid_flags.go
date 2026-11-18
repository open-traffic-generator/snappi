package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaPrefixSidFlags *****
type ospfv2LsaPrefixSidFlags struct {
	validation
	obj          *otg.Ospfv2LsaPrefixSidFlags
	marshaller   marshalOspfv2LsaPrefixSidFlags
	unMarshaller unMarshalOspfv2LsaPrefixSidFlags
}

func NewOspfv2LsaPrefixSidFlags() Ospfv2LsaPrefixSidFlags {
	obj := ospfv2LsaPrefixSidFlags{obj: &otg.Ospfv2LsaPrefixSidFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaPrefixSidFlags) msg() *otg.Ospfv2LsaPrefixSidFlags {
	return obj.obj
}

func (obj *ospfv2LsaPrefixSidFlags) setMsg(msg *otg.Ospfv2LsaPrefixSidFlags) Ospfv2LsaPrefixSidFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaPrefixSidFlags struct {
	obj *ospfv2LsaPrefixSidFlags
}

type marshalOspfv2LsaPrefixSidFlags interface {
	// ToProto marshals Ospfv2LsaPrefixSidFlags to protobuf object *otg.Ospfv2LsaPrefixSidFlags
	ToProto() (*otg.Ospfv2LsaPrefixSidFlags, error)
	// ToPbText marshals Ospfv2LsaPrefixSidFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaPrefixSidFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaPrefixSidFlags to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaPrefixSidFlags struct {
	obj *ospfv2LsaPrefixSidFlags
}

type unMarshalOspfv2LsaPrefixSidFlags interface {
	// FromProto unmarshals Ospfv2LsaPrefixSidFlags from protobuf object *otg.Ospfv2LsaPrefixSidFlags
	FromProto(msg *otg.Ospfv2LsaPrefixSidFlags) (Ospfv2LsaPrefixSidFlags, error)
	// FromPbText unmarshals Ospfv2LsaPrefixSidFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaPrefixSidFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaPrefixSidFlags from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaPrefixSidFlags) Marshal() marshalOspfv2LsaPrefixSidFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaPrefixSidFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaPrefixSidFlags) Unmarshal() unMarshalOspfv2LsaPrefixSidFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaPrefixSidFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaPrefixSidFlags) ToProto() (*otg.Ospfv2LsaPrefixSidFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaPrefixSidFlags) FromProto(msg *otg.Ospfv2LsaPrefixSidFlags) (Ospfv2LsaPrefixSidFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaPrefixSidFlags) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaPrefixSidFlags) FromPbText(value string) error {
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

func (m *marshalospfv2LsaPrefixSidFlags) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaPrefixSidFlags) FromYaml(value string) error {
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

func (m *marshalospfv2LsaPrefixSidFlags) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaPrefixSidFlags) FromJson(value string) error {
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

func (obj *ospfv2LsaPrefixSidFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaPrefixSidFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaPrefixSidFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaPrefixSidFlags) Clone() (Ospfv2LsaPrefixSidFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaPrefixSidFlags()
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

// Ospfv2LsaPrefixSidFlags is one-octet flags of the OSPFv2 Prefix-SID sub-TLV (RFC 8665).
type Ospfv2LsaPrefixSidFlags interface {
	Validation
	// msg marshals Ospfv2LsaPrefixSidFlags to protobuf object *otg.Ospfv2LsaPrefixSidFlags
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaPrefixSidFlags
	// setMsg unmarshals Ospfv2LsaPrefixSidFlags from protobuf object *otg.Ospfv2LsaPrefixSidFlags
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaPrefixSidFlags) Ospfv2LsaPrefixSidFlags
	// provides marshal interface
	Marshal() marshalOspfv2LsaPrefixSidFlags
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaPrefixSidFlags
	// validate validates Ospfv2LsaPrefixSidFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaPrefixSidFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// NpFlag returns bool, set in Ospfv2LsaPrefixSidFlags.
	NpFlag() bool
	// SetNpFlag assigns bool provided by user to Ospfv2LsaPrefixSidFlags
	SetNpFlag(value bool) Ospfv2LsaPrefixSidFlags
	// HasNpFlag checks if NpFlag has been set in Ospfv2LsaPrefixSidFlags
	HasNpFlag() bool
	// MFlag returns bool, set in Ospfv2LsaPrefixSidFlags.
	MFlag() bool
	// SetMFlag assigns bool provided by user to Ospfv2LsaPrefixSidFlags
	SetMFlag(value bool) Ospfv2LsaPrefixSidFlags
	// HasMFlag checks if MFlag has been set in Ospfv2LsaPrefixSidFlags
	HasMFlag() bool
	// EFlag returns bool, set in Ospfv2LsaPrefixSidFlags.
	EFlag() bool
	// SetEFlag assigns bool provided by user to Ospfv2LsaPrefixSidFlags
	SetEFlag(value bool) Ospfv2LsaPrefixSidFlags
	// HasEFlag checks if EFlag has been set in Ospfv2LsaPrefixSidFlags
	HasEFlag() bool
	// VFlag returns bool, set in Ospfv2LsaPrefixSidFlags.
	VFlag() bool
	// SetVFlag assigns bool provided by user to Ospfv2LsaPrefixSidFlags
	SetVFlag(value bool) Ospfv2LsaPrefixSidFlags
	// HasVFlag checks if VFlag has been set in Ospfv2LsaPrefixSidFlags
	HasVFlag() bool
	// LFlag returns bool, set in Ospfv2LsaPrefixSidFlags.
	LFlag() bool
	// SetLFlag assigns bool provided by user to Ospfv2LsaPrefixSidFlags
	SetLFlag(value bool) Ospfv2LsaPrefixSidFlags
	// HasLFlag checks if LFlag has been set in Ospfv2LsaPrefixSidFlags
	HasLFlag() bool
}

// NP-Flag (No-PHP): if set, the penultimate hop MUST NOT pop the Prefix-SID before delivering the packet to the advertising node.
// NpFlag returns a bool
func (obj *ospfv2LsaPrefixSidFlags) NpFlag() bool {

	return *obj.obj.NpFlag

}

// NP-Flag (No-PHP): if set, the penultimate hop MUST NOT pop the Prefix-SID before delivering the packet to the advertising node.
// NpFlag returns a bool
func (obj *ospfv2LsaPrefixSidFlags) HasNpFlag() bool {
	return obj.obj.NpFlag != nil
}

// NP-Flag (No-PHP): if set, the penultimate hop MUST NOT pop the Prefix-SID before delivering the packet to the advertising node.
// SetNpFlag sets the bool value in the Ospfv2LsaPrefixSidFlags object
func (obj *ospfv2LsaPrefixSidFlags) SetNpFlag(value bool) Ospfv2LsaPrefixSidFlags {

	obj.obj.NpFlag = &value
	return obj
}

// M-Flag (Mapping Server): if set, the SID was advertised by an SR Mapping Server.
// MFlag returns a bool
func (obj *ospfv2LsaPrefixSidFlags) MFlag() bool {

	return *obj.obj.MFlag

}

// M-Flag (Mapping Server): if set, the SID was advertised by an SR Mapping Server.
// MFlag returns a bool
func (obj *ospfv2LsaPrefixSidFlags) HasMFlag() bool {
	return obj.obj.MFlag != nil
}

// M-Flag (Mapping Server): if set, the SID was advertised by an SR Mapping Server.
// SetMFlag sets the bool value in the Ospfv2LsaPrefixSidFlags object
func (obj *ospfv2LsaPrefixSidFlags) SetMFlag(value bool) Ospfv2LsaPrefixSidFlags {

	obj.obj.MFlag = &value
	return obj
}

// E-Flag (Explicit-Null): if set, the upstream neighbor MUST replace the Prefix-SID with the Explicit-NULL label before forwarding.
// EFlag returns a bool
func (obj *ospfv2LsaPrefixSidFlags) EFlag() bool {

	return *obj.obj.EFlag

}

// E-Flag (Explicit-Null): if set, the upstream neighbor MUST replace the Prefix-SID with the Explicit-NULL label before forwarding.
// EFlag returns a bool
func (obj *ospfv2LsaPrefixSidFlags) HasEFlag() bool {
	return obj.obj.EFlag != nil
}

// E-Flag (Explicit-Null): if set, the upstream neighbor MUST replace the Prefix-SID with the Explicit-NULL label before forwarding.
// SetEFlag sets the bool value in the Ospfv2LsaPrefixSidFlags object
func (obj *ospfv2LsaPrefixSidFlags) SetEFlag(value bool) Ospfv2LsaPrefixSidFlags {

	obj.obj.EFlag = &value
	return obj
}

// V-Flag (Value): if set, the Prefix-SID carries an absolute value (label); if clear, an index.
// VFlag returns a bool
func (obj *ospfv2LsaPrefixSidFlags) VFlag() bool {

	return *obj.obj.VFlag

}

// V-Flag (Value): if set, the Prefix-SID carries an absolute value (label); if clear, an index.
// VFlag returns a bool
func (obj *ospfv2LsaPrefixSidFlags) HasVFlag() bool {
	return obj.obj.VFlag != nil
}

// V-Flag (Value): if set, the Prefix-SID carries an absolute value (label); if clear, an index.
// SetVFlag sets the bool value in the Ospfv2LsaPrefixSidFlags object
func (obj *ospfv2LsaPrefixSidFlags) SetVFlag(value bool) Ospfv2LsaPrefixSidFlags {

	obj.obj.VFlag = &value
	return obj
}

// L-Flag (Local): if set, the value/index carried by the Prefix-SID has local significance.
// LFlag returns a bool
func (obj *ospfv2LsaPrefixSidFlags) LFlag() bool {

	return *obj.obj.LFlag

}

// L-Flag (Local): if set, the value/index carried by the Prefix-SID has local significance.
// LFlag returns a bool
func (obj *ospfv2LsaPrefixSidFlags) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// L-Flag (Local): if set, the value/index carried by the Prefix-SID has local significance.
// SetLFlag sets the bool value in the Ospfv2LsaPrefixSidFlags object
func (obj *ospfv2LsaPrefixSidFlags) SetLFlag(value bool) Ospfv2LsaPrefixSidFlags {

	obj.obj.LFlag = &value
	return obj
}

func (obj *ospfv2LsaPrefixSidFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2LsaPrefixSidFlags) setDefault() {

}
