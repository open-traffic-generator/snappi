package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2V4RRExtdPrefixFlags *****
type ospfv2V4RRExtdPrefixFlags struct {
	validation
	obj          *otg.Ospfv2V4RRExtdPrefixFlags
	marshaller   marshalOspfv2V4RRExtdPrefixFlags
	unMarshaller unMarshalOspfv2V4RRExtdPrefixFlags
}

func NewOspfv2V4RRExtdPrefixFlags() Ospfv2V4RRExtdPrefixFlags {
	obj := ospfv2V4RRExtdPrefixFlags{obj: &otg.Ospfv2V4RRExtdPrefixFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2V4RRExtdPrefixFlags) msg() *otg.Ospfv2V4RRExtdPrefixFlags {
	return obj.obj
}

func (obj *ospfv2V4RRExtdPrefixFlags) setMsg(msg *otg.Ospfv2V4RRExtdPrefixFlags) Ospfv2V4RRExtdPrefixFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2V4RRExtdPrefixFlags struct {
	obj *ospfv2V4RRExtdPrefixFlags
}

type marshalOspfv2V4RRExtdPrefixFlags interface {
	// ToProto marshals Ospfv2V4RRExtdPrefixFlags to protobuf object *otg.Ospfv2V4RRExtdPrefixFlags
	ToProto() (*otg.Ospfv2V4RRExtdPrefixFlags, error)
	// ToPbText marshals Ospfv2V4RRExtdPrefixFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2V4RRExtdPrefixFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2V4RRExtdPrefixFlags to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2V4RRExtdPrefixFlags struct {
	obj *ospfv2V4RRExtdPrefixFlags
}

type unMarshalOspfv2V4RRExtdPrefixFlags interface {
	// FromProto unmarshals Ospfv2V4RRExtdPrefixFlags from protobuf object *otg.Ospfv2V4RRExtdPrefixFlags
	FromProto(msg *otg.Ospfv2V4RRExtdPrefixFlags) (Ospfv2V4RRExtdPrefixFlags, error)
	// FromPbText unmarshals Ospfv2V4RRExtdPrefixFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2V4RRExtdPrefixFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2V4RRExtdPrefixFlags from JSON text
	FromJson(value string) error
}

func (obj *ospfv2V4RRExtdPrefixFlags) Marshal() marshalOspfv2V4RRExtdPrefixFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2V4RRExtdPrefixFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2V4RRExtdPrefixFlags) Unmarshal() unMarshalOspfv2V4RRExtdPrefixFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2V4RRExtdPrefixFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2V4RRExtdPrefixFlags) ToProto() (*otg.Ospfv2V4RRExtdPrefixFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2V4RRExtdPrefixFlags) FromProto(msg *otg.Ospfv2V4RRExtdPrefixFlags) (Ospfv2V4RRExtdPrefixFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2V4RRExtdPrefixFlags) ToPbText() (string, error) {
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

func (m *unMarshalospfv2V4RRExtdPrefixFlags) FromPbText(value string) error {
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

func (m *marshalospfv2V4RRExtdPrefixFlags) ToYaml() (string, error) {
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

func (m *unMarshalospfv2V4RRExtdPrefixFlags) FromYaml(value string) error {
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

func (m *marshalospfv2V4RRExtdPrefixFlags) ToJson() (string, error) {
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

func (m *unMarshalospfv2V4RRExtdPrefixFlags) FromJson(value string) error {
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

func (obj *ospfv2V4RRExtdPrefixFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2V4RRExtdPrefixFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2V4RRExtdPrefixFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2V4RRExtdPrefixFlags) Clone() (Ospfv2V4RRExtdPrefixFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2V4RRExtdPrefixFlags()
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

// Ospfv2V4RRExtdPrefixFlags is one-octet field contains flags applicable to the prefix. https://datatracker.ietf.org/doc/html/rfc7684.
type Ospfv2V4RRExtdPrefixFlags interface {
	Validation
	// msg marshals Ospfv2V4RRExtdPrefixFlags to protobuf object *otg.Ospfv2V4RRExtdPrefixFlags
	// and doesn't set defaults
	msg() *otg.Ospfv2V4RRExtdPrefixFlags
	// setMsg unmarshals Ospfv2V4RRExtdPrefixFlags from protobuf object *otg.Ospfv2V4RRExtdPrefixFlags
	// and doesn't set defaults
	setMsg(*otg.Ospfv2V4RRExtdPrefixFlags) Ospfv2V4RRExtdPrefixFlags
	// provides marshal interface
	Marshal() marshalOspfv2V4RRExtdPrefixFlags
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2V4RRExtdPrefixFlags
	// validate validates Ospfv2V4RRExtdPrefixFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2V4RRExtdPrefixFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AFlag returns bool, set in Ospfv2V4RRExtdPrefixFlags.
	AFlag() bool
	// SetAFlag assigns bool provided by user to Ospfv2V4RRExtdPrefixFlags
	SetAFlag(value bool) Ospfv2V4RRExtdPrefixFlags
	// HasAFlag checks if AFlag has been set in Ospfv2V4RRExtdPrefixFlags
	HasAFlag() bool
	// NFlag returns bool, set in Ospfv2V4RRExtdPrefixFlags.
	NFlag() bool
	// SetNFlag assigns bool provided by user to Ospfv2V4RRExtdPrefixFlags
	SetNFlag(value bool) Ospfv2V4RRExtdPrefixFlags
	// HasNFlag checks if NFlag has been set in Ospfv2V4RRExtdPrefixFlags
	HasNFlag() bool
}

// 0x80 - (Attach Flag): An Area Border Router (ABR)
// generating an OSPFv2 Extended Prefix TLV for an inter-area
// prefix that is locally connected or attached in another
// connected area SHOULD set this flag.
// AFlag returns a bool
func (obj *ospfv2V4RRExtdPrefixFlags) AFlag() bool {

	return *obj.obj.AFlag

}

// 0x80 - (Attach Flag): An Area Border Router (ABR)
// generating an OSPFv2 Extended Prefix TLV for an inter-area
// prefix that is locally connected or attached in another
// connected area SHOULD set this flag.
// AFlag returns a bool
func (obj *ospfv2V4RRExtdPrefixFlags) HasAFlag() bool {
	return obj.obj.AFlag != nil
}

// 0x80 - (Attach Flag): An Area Border Router (ABR)
// generating an OSPFv2 Extended Prefix TLV for an inter-area
// prefix that is locally connected or attached in another
// connected area SHOULD set this flag.
// SetAFlag sets the bool value in the Ospfv2V4RRExtdPrefixFlags object
func (obj *ospfv2V4RRExtdPrefixFlags) SetAFlag(value bool) Ospfv2V4RRExtdPrefixFlags {

	obj.obj.AFlag = &value
	return obj
}

// N-Flag (Node Flag): Set when the prefix identifies the
// advertising router, i.e., the prefix is a host prefix
// advertising a globally reachable address typically associated
// with a loopback address.
// NFlag returns a bool
func (obj *ospfv2V4RRExtdPrefixFlags) NFlag() bool {

	return *obj.obj.NFlag

}

// N-Flag (Node Flag): Set when the prefix identifies the
// advertising router, i.e., the prefix is a host prefix
// advertising a globally reachable address typically associated
// with a loopback address.
// NFlag returns a bool
func (obj *ospfv2V4RRExtdPrefixFlags) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// N-Flag (Node Flag): Set when the prefix identifies the
// advertising router, i.e., the prefix is a host prefix
// advertising a globally reachable address typically associated
// with a loopback address.
// SetNFlag sets the bool value in the Ospfv2V4RRExtdPrefixFlags object
func (obj *ospfv2V4RRExtdPrefixFlags) SetNFlag(value bool) Ospfv2V4RRExtdPrefixFlags {

	obj.obj.NFlag = &value
	return obj
}

func (obj *ospfv2V4RRExtdPrefixFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2V4RRExtdPrefixFlags) setDefault() {
	if obj.obj.AFlag == nil {
		obj.SetAFlag(false)
	}
	if obj.obj.NFlag == nil {
		obj.SetNFlag(false)
	}

}
