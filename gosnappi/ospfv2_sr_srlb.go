package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2SRSrlb *****
type ospfv2SRSrlb struct {
	validation
	obj          *otg.Ospfv2SRSrlb
	marshaller   marshalOspfv2SRSrlb
	unMarshaller unMarshalOspfv2SRSrlb
}

func NewOspfv2SRSrlb() Ospfv2SRSrlb {
	obj := ospfv2SRSrlb{obj: &otg.Ospfv2SRSrlb{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2SRSrlb) msg() *otg.Ospfv2SRSrlb {
	return obj.obj
}

func (obj *ospfv2SRSrlb) setMsg(msg *otg.Ospfv2SRSrlb) Ospfv2SRSrlb {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2SRSrlb struct {
	obj *ospfv2SRSrlb
}

type marshalOspfv2SRSrlb interface {
	// ToProto marshals Ospfv2SRSrlb to protobuf object *otg.Ospfv2SRSrlb
	ToProto() (*otg.Ospfv2SRSrlb, error)
	// ToPbText marshals Ospfv2SRSrlb to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2SRSrlb to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2SRSrlb to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2SRSrlb struct {
	obj *ospfv2SRSrlb
}

type unMarshalOspfv2SRSrlb interface {
	// FromProto unmarshals Ospfv2SRSrlb from protobuf object *otg.Ospfv2SRSrlb
	FromProto(msg *otg.Ospfv2SRSrlb) (Ospfv2SRSrlb, error)
	// FromPbText unmarshals Ospfv2SRSrlb from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2SRSrlb from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2SRSrlb from JSON text
	FromJson(value string) error
}

func (obj *ospfv2SRSrlb) Marshal() marshalOspfv2SRSrlb {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2SRSrlb{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2SRSrlb) Unmarshal() unMarshalOspfv2SRSrlb {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2SRSrlb{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2SRSrlb) ToProto() (*otg.Ospfv2SRSrlb, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2SRSrlb) FromProto(msg *otg.Ospfv2SRSrlb) (Ospfv2SRSrlb, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2SRSrlb) ToPbText() (string, error) {
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

func (m *unMarshalospfv2SRSrlb) FromPbText(value string) error {
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

func (m *marshalospfv2SRSrlb) ToYaml() (string, error) {
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

func (m *unMarshalospfv2SRSrlb) FromYaml(value string) error {
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

func (m *marshalospfv2SRSrlb) ToJson() (string, error) {
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

func (m *unMarshalospfv2SRSrlb) FromJson(value string) error {
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

func (obj *ospfv2SRSrlb) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2SRSrlb) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2SRSrlb) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2SRSrlb) Clone() (Ospfv2SRSrlb, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2SRSrlb()
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

// Ospfv2SRSrlb is this contains the properties of a single SR Local Block (SRLB) range.
// The SR Local Block TLV carries the size of the range and a mandatory SID/Label sub-TLV
// that provides the first (base) SID/Label value of the local block.
// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-sr-local-block-tlv.
type Ospfv2SRSrlb interface {
	Validation
	// msg marshals Ospfv2SRSrlb to protobuf object *otg.Ospfv2SRSrlb
	// and doesn't set defaults
	msg() *otg.Ospfv2SRSrlb
	// setMsg unmarshals Ospfv2SRSrlb from protobuf object *otg.Ospfv2SRSrlb
	// and doesn't set defaults
	setMsg(*otg.Ospfv2SRSrlb) Ospfv2SRSrlb
	// provides marshal interface
	Marshal() marshalOspfv2SRSrlb
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2SRSrlb
	// validate validates Ospfv2SRSrlb
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2SRSrlb, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StartingSid returns uint32, set in Ospfv2SRSrlb.
	StartingSid() uint32
	// SetStartingSid assigns uint32 provided by user to Ospfv2SRSrlb
	SetStartingSid(value uint32) Ospfv2SRSrlb
	// HasStartingSid checks if StartingSid has been set in Ospfv2SRSrlb
	HasStartingSid() bool
	// Range returns uint32, set in Ospfv2SRSrlb.
	Range() uint32
	// SetRange assigns uint32 provided by user to Ospfv2SRSrlb
	SetRange(value uint32) Ospfv2SRSrlb
	// HasRange checks if Range has been set in Ospfv2SRSrlb
	HasRange() bool
}

// The first value of the SRLB range provided by the SID/Label sub-TLV.
// StartingSid returns a uint32
func (obj *ospfv2SRSrlb) StartingSid() uint32 {

	return *obj.obj.StartingSid

}

// The first value of the SRLB range provided by the SID/Label sub-TLV.
// StartingSid returns a uint32
func (obj *ospfv2SRSrlb) HasStartingSid() bool {
	return obj.obj.StartingSid != nil
}

// The first value of the SRLB range provided by the SID/Label sub-TLV.
// SetStartingSid sets the uint32 value in the Ospfv2SRSrlb object
func (obj *ospfv2SRSrlb) SetStartingSid(value uint32) Ospfv2SRSrlb {

	obj.obj.StartingSid = &value
	return obj
}

// The Range Size, i.e. the number of SIDs in this SRLB range. MUST be greater than zero.
// Range returns a uint32
func (obj *ospfv2SRSrlb) Range() uint32 {

	return *obj.obj.Range

}

// The Range Size, i.e. the number of SIDs in this SRLB range. MUST be greater than zero.
// Range returns a uint32
func (obj *ospfv2SRSrlb) HasRange() bool {
	return obj.obj.Range != nil
}

// The Range Size, i.e. the number of SIDs in this SRLB range. MUST be greater than zero.
// SetRange sets the uint32 value in the Ospfv2SRSrlb object
func (obj *ospfv2SRSrlb) SetRange(value uint32) Ospfv2SRSrlb {

	obj.obj.Range = &value
	return obj
}

func (obj *ospfv2SRSrlb) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StartingSid != nil {

		if *obj.obj.StartingSid < 1 || *obj.obj.StartingSid > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Ospfv2SRSrlb.StartingSid <= 4294967295 but Got %d", *obj.obj.StartingSid))
		}

	}

	if obj.obj.Range != nil {

		if *obj.obj.Range < 1 || *obj.obj.Range > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Ospfv2SRSrlb.Range <= 16777215 but Got %d", *obj.obj.Range))
		}

	}

}

func (obj *ospfv2SRSrlb) setDefault() {
	if obj.obj.StartingSid == nil {
		obj.SetStartingSid(16000)
	}
	if obj.obj.Range == nil {
		obj.SetRange(8000)
	}

}
