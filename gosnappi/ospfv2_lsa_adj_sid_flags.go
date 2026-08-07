package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaAdjSidFlags *****
type ospfv2LsaAdjSidFlags struct {
	validation
	obj          *otg.Ospfv2LsaAdjSidFlags
	marshaller   marshalOspfv2LsaAdjSidFlags
	unMarshaller unMarshalOspfv2LsaAdjSidFlags
}

func NewOspfv2LsaAdjSidFlags() Ospfv2LsaAdjSidFlags {
	obj := ospfv2LsaAdjSidFlags{obj: &otg.Ospfv2LsaAdjSidFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaAdjSidFlags) msg() *otg.Ospfv2LsaAdjSidFlags {
	return obj.obj
}

func (obj *ospfv2LsaAdjSidFlags) setMsg(msg *otg.Ospfv2LsaAdjSidFlags) Ospfv2LsaAdjSidFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaAdjSidFlags struct {
	obj *ospfv2LsaAdjSidFlags
}

type marshalOspfv2LsaAdjSidFlags interface {
	// ToProto marshals Ospfv2LsaAdjSidFlags to protobuf object *otg.Ospfv2LsaAdjSidFlags
	ToProto() (*otg.Ospfv2LsaAdjSidFlags, error)
	// ToPbText marshals Ospfv2LsaAdjSidFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaAdjSidFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaAdjSidFlags to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaAdjSidFlags struct {
	obj *ospfv2LsaAdjSidFlags
}

type unMarshalOspfv2LsaAdjSidFlags interface {
	// FromProto unmarshals Ospfv2LsaAdjSidFlags from protobuf object *otg.Ospfv2LsaAdjSidFlags
	FromProto(msg *otg.Ospfv2LsaAdjSidFlags) (Ospfv2LsaAdjSidFlags, error)
	// FromPbText unmarshals Ospfv2LsaAdjSidFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaAdjSidFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaAdjSidFlags from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaAdjSidFlags) Marshal() marshalOspfv2LsaAdjSidFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaAdjSidFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaAdjSidFlags) Unmarshal() unMarshalOspfv2LsaAdjSidFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaAdjSidFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaAdjSidFlags) ToProto() (*otg.Ospfv2LsaAdjSidFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaAdjSidFlags) FromProto(msg *otg.Ospfv2LsaAdjSidFlags) (Ospfv2LsaAdjSidFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaAdjSidFlags) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaAdjSidFlags) FromPbText(value string) error {
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

func (m *marshalospfv2LsaAdjSidFlags) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaAdjSidFlags) FromYaml(value string) error {
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

func (m *marshalospfv2LsaAdjSidFlags) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaAdjSidFlags) FromJson(value string) error {
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

func (obj *ospfv2LsaAdjSidFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaAdjSidFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaAdjSidFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaAdjSidFlags) Clone() (Ospfv2LsaAdjSidFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaAdjSidFlags()
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

// Ospfv2LsaAdjSidFlags is one-octet flags of the OSPFv2 Adjacency-SID sub-TLV (RFC 8665).
type Ospfv2LsaAdjSidFlags interface {
	Validation
	// msg marshals Ospfv2LsaAdjSidFlags to protobuf object *otg.Ospfv2LsaAdjSidFlags
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaAdjSidFlags
	// setMsg unmarshals Ospfv2LsaAdjSidFlags from protobuf object *otg.Ospfv2LsaAdjSidFlags
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaAdjSidFlags) Ospfv2LsaAdjSidFlags
	// provides marshal interface
	Marshal() marshalOspfv2LsaAdjSidFlags
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaAdjSidFlags
	// validate validates Ospfv2LsaAdjSidFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaAdjSidFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// BFlag returns bool, set in Ospfv2LsaAdjSidFlags.
	BFlag() bool
	// SetBFlag assigns bool provided by user to Ospfv2LsaAdjSidFlags
	SetBFlag(value bool) Ospfv2LsaAdjSidFlags
	// HasBFlag checks if BFlag has been set in Ospfv2LsaAdjSidFlags
	HasBFlag() bool
	// GFlag returns bool, set in Ospfv2LsaAdjSidFlags.
	GFlag() bool
	// SetGFlag assigns bool provided by user to Ospfv2LsaAdjSidFlags
	SetGFlag(value bool) Ospfv2LsaAdjSidFlags
	// HasGFlag checks if GFlag has been set in Ospfv2LsaAdjSidFlags
	HasGFlag() bool
	// PFlag returns bool, set in Ospfv2LsaAdjSidFlags.
	PFlag() bool
	// SetPFlag assigns bool provided by user to Ospfv2LsaAdjSidFlags
	SetPFlag(value bool) Ospfv2LsaAdjSidFlags
	// HasPFlag checks if PFlag has been set in Ospfv2LsaAdjSidFlags
	HasPFlag() bool
	// VFlag returns bool, set in Ospfv2LsaAdjSidFlags.
	VFlag() bool
	// SetVFlag assigns bool provided by user to Ospfv2LsaAdjSidFlags
	SetVFlag(value bool) Ospfv2LsaAdjSidFlags
	// HasVFlag checks if VFlag has been set in Ospfv2LsaAdjSidFlags
	HasVFlag() bool
	// LFlag returns bool, set in Ospfv2LsaAdjSidFlags.
	LFlag() bool
	// SetLFlag assigns bool provided by user to Ospfv2LsaAdjSidFlags
	SetLFlag(value bool) Ospfv2LsaAdjSidFlags
	// HasLFlag checks if LFlag has been set in Ospfv2LsaAdjSidFlags
	HasLFlag() bool
}

// B-Flag (Backup): the Adjacency-SID is eligible for protection.
// BFlag returns a bool
func (obj *ospfv2LsaAdjSidFlags) BFlag() bool {

	return *obj.obj.BFlag

}

// B-Flag (Backup): the Adjacency-SID is eligible for protection.
// BFlag returns a bool
func (obj *ospfv2LsaAdjSidFlags) HasBFlag() bool {
	return obj.obj.BFlag != nil
}

// B-Flag (Backup): the Adjacency-SID is eligible for protection.
// SetBFlag sets the bool value in the Ospfv2LsaAdjSidFlags object
func (obj *ospfv2LsaAdjSidFlags) SetBFlag(value bool) Ospfv2LsaAdjSidFlags {

	obj.obj.BFlag = &value
	return obj
}

// G-Flag (Group): the Adjacency-SID refers to a group of adjacencies.
// GFlag returns a bool
func (obj *ospfv2LsaAdjSidFlags) GFlag() bool {

	return *obj.obj.GFlag

}

// G-Flag (Group): the Adjacency-SID refers to a group of adjacencies.
// GFlag returns a bool
func (obj *ospfv2LsaAdjSidFlags) HasGFlag() bool {
	return obj.obj.GFlag != nil
}

// G-Flag (Group): the Adjacency-SID refers to a group of adjacencies.
// SetGFlag sets the bool value in the Ospfv2LsaAdjSidFlags object
func (obj *ospfv2LsaAdjSidFlags) SetGFlag(value bool) Ospfv2LsaAdjSidFlags {

	obj.obj.GFlag = &value
	return obj
}

// P-Flag (Persistent): the Adjacency-SID is persistently allocated.
// PFlag returns a bool
func (obj *ospfv2LsaAdjSidFlags) PFlag() bool {

	return *obj.obj.PFlag

}

// P-Flag (Persistent): the Adjacency-SID is persistently allocated.
// PFlag returns a bool
func (obj *ospfv2LsaAdjSidFlags) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// P-Flag (Persistent): the Adjacency-SID is persistently allocated.
// SetPFlag sets the bool value in the Ospfv2LsaAdjSidFlags object
func (obj *ospfv2LsaAdjSidFlags) SetPFlag(value bool) Ospfv2LsaAdjSidFlags {

	obj.obj.PFlag = &value
	return obj
}

// V-Flag (Value): if set, the Adj-SID carries an absolute value (label); if clear, an index.
// VFlag returns a bool
func (obj *ospfv2LsaAdjSidFlags) VFlag() bool {

	return *obj.obj.VFlag

}

// V-Flag (Value): if set, the Adj-SID carries an absolute value (label); if clear, an index.
// VFlag returns a bool
func (obj *ospfv2LsaAdjSidFlags) HasVFlag() bool {
	return obj.obj.VFlag != nil
}

// V-Flag (Value): if set, the Adj-SID carries an absolute value (label); if clear, an index.
// SetVFlag sets the bool value in the Ospfv2LsaAdjSidFlags object
func (obj *ospfv2LsaAdjSidFlags) SetVFlag(value bool) Ospfv2LsaAdjSidFlags {

	obj.obj.VFlag = &value
	return obj
}

// L-Flag (Local): if set, the value/index carried by the Adj-SID has local significance.
// LFlag returns a bool
func (obj *ospfv2LsaAdjSidFlags) LFlag() bool {

	return *obj.obj.LFlag

}

// L-Flag (Local): if set, the value/index carried by the Adj-SID has local significance.
// LFlag returns a bool
func (obj *ospfv2LsaAdjSidFlags) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// L-Flag (Local): if set, the value/index carried by the Adj-SID has local significance.
// SetLFlag sets the bool value in the Ospfv2LsaAdjSidFlags object
func (obj *ospfv2LsaAdjSidFlags) SetLFlag(value bool) Ospfv2LsaAdjSidFlags {

	obj.obj.LFlag = &value
	return obj
}

func (obj *ospfv2LsaAdjSidFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2LsaAdjSidFlags) setDefault() {

}
