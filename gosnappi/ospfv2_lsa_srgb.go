package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaSrgb *****
type ospfv2LsaSrgb struct {
	validation
	obj          *otg.Ospfv2LsaSrgb
	marshaller   marshalOspfv2LsaSrgb
	unMarshaller unMarshalOspfv2LsaSrgb
}

func NewOspfv2LsaSrgb() Ospfv2LsaSrgb {
	obj := ospfv2LsaSrgb{obj: &otg.Ospfv2LsaSrgb{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaSrgb) msg() *otg.Ospfv2LsaSrgb {
	return obj.obj
}

func (obj *ospfv2LsaSrgb) setMsg(msg *otg.Ospfv2LsaSrgb) Ospfv2LsaSrgb {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaSrgb struct {
	obj *ospfv2LsaSrgb
}

type marshalOspfv2LsaSrgb interface {
	// ToProto marshals Ospfv2LsaSrgb to protobuf object *otg.Ospfv2LsaSrgb
	ToProto() (*otg.Ospfv2LsaSrgb, error)
	// ToPbText marshals Ospfv2LsaSrgb to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaSrgb to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaSrgb to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaSrgb struct {
	obj *ospfv2LsaSrgb
}

type unMarshalOspfv2LsaSrgb interface {
	// FromProto unmarshals Ospfv2LsaSrgb from protobuf object *otg.Ospfv2LsaSrgb
	FromProto(msg *otg.Ospfv2LsaSrgb) (Ospfv2LsaSrgb, error)
	// FromPbText unmarshals Ospfv2LsaSrgb from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaSrgb from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaSrgb from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaSrgb) Marshal() marshalOspfv2LsaSrgb {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaSrgb{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaSrgb) Unmarshal() unMarshalOspfv2LsaSrgb {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaSrgb{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaSrgb) ToProto() (*otg.Ospfv2LsaSrgb, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaSrgb) FromProto(msg *otg.Ospfv2LsaSrgb) (Ospfv2LsaSrgb, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaSrgb) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaSrgb) FromPbText(value string) error {
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

func (m *marshalospfv2LsaSrgb) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaSrgb) FromYaml(value string) error {
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

func (m *marshalospfv2LsaSrgb) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaSrgb) FromJson(value string) error {
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

func (obj *ospfv2LsaSrgb) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaSrgb) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaSrgb) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaSrgb) Clone() (Ospfv2LsaSrgb, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaSrgb()
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

// Ospfv2LsaSrgb is a learned Segment Routing Global Block (SRGB) range.
type Ospfv2LsaSrgb interface {
	Validation
	// msg marshals Ospfv2LsaSrgb to protobuf object *otg.Ospfv2LsaSrgb
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaSrgb
	// setMsg unmarshals Ospfv2LsaSrgb from protobuf object *otg.Ospfv2LsaSrgb
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaSrgb) Ospfv2LsaSrgb
	// provides marshal interface
	Marshal() marshalOspfv2LsaSrgb
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaSrgb
	// validate validates Ospfv2LsaSrgb
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaSrgb, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StartingSid returns uint32, set in Ospfv2LsaSrgb.
	StartingSid() uint32
	// SetStartingSid assigns uint32 provided by user to Ospfv2LsaSrgb
	SetStartingSid(value uint32) Ospfv2LsaSrgb
	// HasStartingSid checks if StartingSid has been set in Ospfv2LsaSrgb
	HasStartingSid() bool
	// Range returns uint32, set in Ospfv2LsaSrgb.
	Range() uint32
	// SetRange assigns uint32 provided by user to Ospfv2LsaSrgb
	SetRange(value uint32) Ospfv2LsaSrgb
	// HasRange checks if Range has been set in Ospfv2LsaSrgb
	HasRange() bool
}

// The first value (base SID/label) of the SRGB range.
// StartingSid returns a uint32
func (obj *ospfv2LsaSrgb) StartingSid() uint32 {

	return *obj.obj.StartingSid

}

// The first value (base SID/label) of the SRGB range.
// StartingSid returns a uint32
func (obj *ospfv2LsaSrgb) HasStartingSid() bool {
	return obj.obj.StartingSid != nil
}

// The first value (base SID/label) of the SRGB range.
// SetStartingSid sets the uint32 value in the Ospfv2LsaSrgb object
func (obj *ospfv2LsaSrgb) SetStartingSid(value uint32) Ospfv2LsaSrgb {

	obj.obj.StartingSid = &value
	return obj
}

// The number of SIDs in this SRGB range.
// Range returns a uint32
func (obj *ospfv2LsaSrgb) Range() uint32 {

	return *obj.obj.Range

}

// The number of SIDs in this SRGB range.
// Range returns a uint32
func (obj *ospfv2LsaSrgb) HasRange() bool {
	return obj.obj.Range != nil
}

// The number of SIDs in this SRGB range.
// SetRange sets the uint32 value in the Ospfv2LsaSrgb object
func (obj *ospfv2LsaSrgb) SetRange(value uint32) Ospfv2LsaSrgb {

	obj.obj.Range = &value
	return obj
}

func (obj *ospfv2LsaSrgb) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2LsaSrgb) setDefault() {

}
