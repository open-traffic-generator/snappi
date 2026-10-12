package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2SRSrgb *****
type ospfv2SRSrgb struct {
	validation
	obj          *otg.Ospfv2SRSrgb
	marshaller   marshalOspfv2SRSrgb
	unMarshaller unMarshalOspfv2SRSrgb
}

func NewOspfv2SRSrgb() Ospfv2SRSrgb {
	obj := ospfv2SRSrgb{obj: &otg.Ospfv2SRSrgb{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2SRSrgb) msg() *otg.Ospfv2SRSrgb {
	return obj.obj
}

func (obj *ospfv2SRSrgb) setMsg(msg *otg.Ospfv2SRSrgb) Ospfv2SRSrgb {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2SRSrgb struct {
	obj *ospfv2SRSrgb
}

type marshalOspfv2SRSrgb interface {
	// ToProto marshals Ospfv2SRSrgb to protobuf object *otg.Ospfv2SRSrgb
	ToProto() (*otg.Ospfv2SRSrgb, error)
	// ToPbText marshals Ospfv2SRSrgb to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2SRSrgb to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2SRSrgb to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2SRSrgb struct {
	obj *ospfv2SRSrgb
}

type unMarshalOspfv2SRSrgb interface {
	// FromProto unmarshals Ospfv2SRSrgb from protobuf object *otg.Ospfv2SRSrgb
	FromProto(msg *otg.Ospfv2SRSrgb) (Ospfv2SRSrgb, error)
	// FromPbText unmarshals Ospfv2SRSrgb from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2SRSrgb from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2SRSrgb from JSON text
	FromJson(value string) error
}

func (obj *ospfv2SRSrgb) Marshal() marshalOspfv2SRSrgb {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2SRSrgb{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2SRSrgb) Unmarshal() unMarshalOspfv2SRSrgb {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2SRSrgb{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2SRSrgb) ToProto() (*otg.Ospfv2SRSrgb, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2SRSrgb) FromProto(msg *otg.Ospfv2SRSrgb) (Ospfv2SRSrgb, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2SRSrgb) ToPbText() (string, error) {
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

func (m *unMarshalospfv2SRSrgb) FromPbText(value string) error {
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

func (m *marshalospfv2SRSrgb) ToYaml() (string, error) {
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

func (m *unMarshalospfv2SRSrgb) FromYaml(value string) error {
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

func (m *marshalospfv2SRSrgb) ToJson() (string, error) {
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

func (m *unMarshalospfv2SRSrgb) FromJson(value string) error {
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

func (obj *ospfv2SRSrgb) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2SRSrgb) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2SRSrgb) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2SRSrgb) Clone() (Ospfv2SRSrgb, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2SRSrgb()
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

// Ospfv2SRSrgb is this contains the properties of a single Segment Routing Global Block (SRGB) range.
// The SID/Label Range TLV carries the size of the range and a mandatory SID/Label sub-TLV
// that provides the first (base) SID/Label value of the range.
// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-sidlabel-range-tlv.
type Ospfv2SRSrgb interface {
	Validation
	// msg marshals Ospfv2SRSrgb to protobuf object *otg.Ospfv2SRSrgb
	// and doesn't set defaults
	msg() *otg.Ospfv2SRSrgb
	// setMsg unmarshals Ospfv2SRSrgb from protobuf object *otg.Ospfv2SRSrgb
	// and doesn't set defaults
	setMsg(*otg.Ospfv2SRSrgb) Ospfv2SRSrgb
	// provides marshal interface
	Marshal() marshalOspfv2SRSrgb
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2SRSrgb
	// validate validates Ospfv2SRSrgb
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2SRSrgb, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StartingSid returns uint32, set in Ospfv2SRSrgb.
	StartingSid() uint32
	// SetStartingSid assigns uint32 provided by user to Ospfv2SRSrgb
	SetStartingSid(value uint32) Ospfv2SRSrgb
	// HasStartingSid checks if StartingSid has been set in Ospfv2SRSrgb
	HasStartingSid() bool
	// Range returns uint32, set in Ospfv2SRSrgb.
	Range() uint32
	// SetRange assigns uint32 provided by user to Ospfv2SRSrgb
	SetRange(value uint32) Ospfv2SRSrgb
	// HasRange checks if Range has been set in Ospfv2SRSrgb
	HasRange() bool
}

// The first value of the SRGB range provided by the SID/Label sub-TLV.
// StartingSid returns a uint32
func (obj *ospfv2SRSrgb) StartingSid() uint32 {

	return *obj.obj.StartingSid

}

// The first value of the SRGB range provided by the SID/Label sub-TLV.
// StartingSid returns a uint32
func (obj *ospfv2SRSrgb) HasStartingSid() bool {
	return obj.obj.StartingSid != nil
}

// The first value of the SRGB range provided by the SID/Label sub-TLV.
// SetStartingSid sets the uint32 value in the Ospfv2SRSrgb object
func (obj *ospfv2SRSrgb) SetStartingSid(value uint32) Ospfv2SRSrgb {

	obj.obj.StartingSid = &value
	return obj
}

// The Range Size, i.e. the number of SIDs in this SRGB range. MUST be greater than zero.
// Range returns a uint32
func (obj *ospfv2SRSrgb) Range() uint32 {

	return *obj.obj.Range

}

// The Range Size, i.e. the number of SIDs in this SRGB range. MUST be greater than zero.
// Range returns a uint32
func (obj *ospfv2SRSrgb) HasRange() bool {
	return obj.obj.Range != nil
}

// The Range Size, i.e. the number of SIDs in this SRGB range. MUST be greater than zero.
// SetRange sets the uint32 value in the Ospfv2SRSrgb object
func (obj *ospfv2SRSrgb) SetRange(value uint32) Ospfv2SRSrgb {

	obj.obj.Range = &value
	return obj
}

func (obj *ospfv2SRSrgb) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StartingSid != nil {

		if *obj.obj.StartingSid < 1 || *obj.obj.StartingSid > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Ospfv2SRSrgb.StartingSid <= 4294967295 but Got %d", *obj.obj.StartingSid))
		}

	}

	if obj.obj.Range != nil {

		if *obj.obj.Range < 1 || *obj.obj.Range > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Ospfv2SRSrgb.Range <= 16777215 but Got %d", *obj.obj.Range))
		}

	}

}

func (obj *ospfv2SRSrgb) setDefault() {
	if obj.obj.StartingSid == nil {
		obj.SetStartingSid(16000)
	}
	if obj.obj.Range == nil {
		obj.SetRange(8000)
	}

}
