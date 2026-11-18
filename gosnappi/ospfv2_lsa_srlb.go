package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaSrlb *****
type ospfv2LsaSrlb struct {
	validation
	obj          *otg.Ospfv2LsaSrlb
	marshaller   marshalOspfv2LsaSrlb
	unMarshaller unMarshalOspfv2LsaSrlb
}

func NewOspfv2LsaSrlb() Ospfv2LsaSrlb {
	obj := ospfv2LsaSrlb{obj: &otg.Ospfv2LsaSrlb{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaSrlb) msg() *otg.Ospfv2LsaSrlb {
	return obj.obj
}

func (obj *ospfv2LsaSrlb) setMsg(msg *otg.Ospfv2LsaSrlb) Ospfv2LsaSrlb {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaSrlb struct {
	obj *ospfv2LsaSrlb
}

type marshalOspfv2LsaSrlb interface {
	// ToProto marshals Ospfv2LsaSrlb to protobuf object *otg.Ospfv2LsaSrlb
	ToProto() (*otg.Ospfv2LsaSrlb, error)
	// ToPbText marshals Ospfv2LsaSrlb to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaSrlb to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaSrlb to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaSrlb struct {
	obj *ospfv2LsaSrlb
}

type unMarshalOspfv2LsaSrlb interface {
	// FromProto unmarshals Ospfv2LsaSrlb from protobuf object *otg.Ospfv2LsaSrlb
	FromProto(msg *otg.Ospfv2LsaSrlb) (Ospfv2LsaSrlb, error)
	// FromPbText unmarshals Ospfv2LsaSrlb from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaSrlb from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaSrlb from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaSrlb) Marshal() marshalOspfv2LsaSrlb {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaSrlb{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaSrlb) Unmarshal() unMarshalOspfv2LsaSrlb {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaSrlb{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaSrlb) ToProto() (*otg.Ospfv2LsaSrlb, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaSrlb) FromProto(msg *otg.Ospfv2LsaSrlb) (Ospfv2LsaSrlb, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaSrlb) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaSrlb) FromPbText(value string) error {
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

func (m *marshalospfv2LsaSrlb) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaSrlb) FromYaml(value string) error {
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

func (m *marshalospfv2LsaSrlb) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaSrlb) FromJson(value string) error {
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

func (obj *ospfv2LsaSrlb) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaSrlb) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaSrlb) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaSrlb) Clone() (Ospfv2LsaSrlb, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaSrlb()
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

// Ospfv2LsaSrlb is a learned SR Local Block (SRLB) range.
type Ospfv2LsaSrlb interface {
	Validation
	// msg marshals Ospfv2LsaSrlb to protobuf object *otg.Ospfv2LsaSrlb
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaSrlb
	// setMsg unmarshals Ospfv2LsaSrlb from protobuf object *otg.Ospfv2LsaSrlb
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaSrlb) Ospfv2LsaSrlb
	// provides marshal interface
	Marshal() marshalOspfv2LsaSrlb
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaSrlb
	// validate validates Ospfv2LsaSrlb
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaSrlb, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StartingSid returns uint32, set in Ospfv2LsaSrlb.
	StartingSid() uint32
	// SetStartingSid assigns uint32 provided by user to Ospfv2LsaSrlb
	SetStartingSid(value uint32) Ospfv2LsaSrlb
	// HasStartingSid checks if StartingSid has been set in Ospfv2LsaSrlb
	HasStartingSid() bool
	// Range returns uint32, set in Ospfv2LsaSrlb.
	Range() uint32
	// SetRange assigns uint32 provided by user to Ospfv2LsaSrlb
	SetRange(value uint32) Ospfv2LsaSrlb
	// HasRange checks if Range has been set in Ospfv2LsaSrlb
	HasRange() bool
}

// The first value (base SID/label) of the SRLB range.
// StartingSid returns a uint32
func (obj *ospfv2LsaSrlb) StartingSid() uint32 {

	return *obj.obj.StartingSid

}

// The first value (base SID/label) of the SRLB range.
// StartingSid returns a uint32
func (obj *ospfv2LsaSrlb) HasStartingSid() bool {
	return obj.obj.StartingSid != nil
}

// The first value (base SID/label) of the SRLB range.
// SetStartingSid sets the uint32 value in the Ospfv2LsaSrlb object
func (obj *ospfv2LsaSrlb) SetStartingSid(value uint32) Ospfv2LsaSrlb {

	obj.obj.StartingSid = &value
	return obj
}

// The number of SIDs in this SRLB range.
// Range returns a uint32
func (obj *ospfv2LsaSrlb) Range() uint32 {

	return *obj.obj.Range

}

// The number of SIDs in this SRLB range.
// Range returns a uint32
func (obj *ospfv2LsaSrlb) HasRange() bool {
	return obj.obj.Range != nil
}

// The number of SIDs in this SRLB range.
// SetRange sets the uint32 value in the Ospfv2LsaSrlb object
func (obj *ospfv2LsaSrlb) SetRange(value uint32) Ospfv2LsaSrlb {

	obj.obj.Range = &value
	return obj
}

func (obj *ospfv2LsaSrlb) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2LsaSrlb) setDefault() {

}
