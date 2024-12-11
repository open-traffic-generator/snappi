package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRSrgb *****
type isisSRSrgb struct {
	validation
	obj          *otg.IsisSRSrgb
	marshaller   marshalIsisSRSrgb
	unMarshaller unMarshalIsisSRSrgb
}

func NewIsisSRSrgb() IsisSRSrgb {
	obj := isisSRSrgb{obj: &otg.IsisSRSrgb{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRSrgb) msg() *otg.IsisSRSrgb {
	return obj.obj
}

func (obj *isisSRSrgb) setMsg(msg *otg.IsisSRSrgb) IsisSRSrgb {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRSrgb struct {
	obj *isisSRSrgb
}

type marshalIsisSRSrgb interface {
	// ToProto marshals IsisSRSrgb to protobuf object *otg.IsisSRSrgb
	ToProto() (*otg.IsisSRSrgb, error)
	// ToPbText marshals IsisSRSrgb to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRSrgb to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRSrgb to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRSrgb struct {
	obj *isisSRSrgb
}

type unMarshalIsisSRSrgb interface {
	// FromProto unmarshals IsisSRSrgb from protobuf object *otg.IsisSRSrgb
	FromProto(msg *otg.IsisSRSrgb) (IsisSRSrgb, error)
	// FromPbText unmarshals IsisSRSrgb from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRSrgb from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRSrgb from JSON text
	FromJson(value string) error
}

func (obj *isisSRSrgb) Marshal() marshalIsisSRSrgb {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRSrgb{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRSrgb) Unmarshal() unMarshalIsisSRSrgb {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRSrgb{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRSrgb) ToProto() (*otg.IsisSRSrgb, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRSrgb) FromProto(msg *otg.IsisSRSrgb) (IsisSRSrgb, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRSrgb) ToPbText() (string, error) {
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

func (m *unMarshalisisSRSrgb) FromPbText(value string) error {
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

func (m *marshalisisSRSrgb) ToYaml() (string, error) {
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

func (m *unMarshalisisSRSrgb) FromYaml(value string) error {
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

func (m *marshalisisSRSrgb) ToJson() (string, error) {
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

func (m *unMarshalisisSRSrgb) FromJson(value string) error {
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

func (obj *isisSRSrgb) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRSrgb) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRSrgb) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRSrgb) Clone() (IsisSRSrgb, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRSrgb()
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

// IsisSRSrgb is this contains the propeties of Segment Routing Global Block (SRGB) Descriptor. Reference: https://datatracker.ietf.org/doc/html/rfc8667#section-3.1-7.1.1
type IsisSRSrgb interface {
	Validation
	// msg marshals IsisSRSrgb to protobuf object *otg.IsisSRSrgb
	// and doesn't set defaults
	msg() *otg.IsisSRSrgb
	// setMsg unmarshals IsisSRSrgb from protobuf object *otg.IsisSRSrgb
	// and doesn't set defaults
	setMsg(*otg.IsisSRSrgb) IsisSRSrgb
	// provides marshal interface
	Marshal() marshalIsisSRSrgb
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRSrgb
	// validate validates IsisSRSrgb
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRSrgb, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Range returns uint32, set in IsisSRSrgb.
	Range() uint32
	// SetRange assigns uint32 provided by user to IsisSRSrgb
	SetRange(value uint32) IsisSRSrgb
	// HasRange checks if Range has been set in IsisSRSrgb
	HasRange() bool
	// StartingSid returns uint32, set in IsisSRSrgb.
	StartingSid() uint32
	// SetStartingSid assigns uint32 provided by user to IsisSRSrgb
	SetStartingSid(value uint32) IsisSRSrgb
	// HasStartingSid checks if StartingSid has been set in IsisSRSrgb
	HasStartingSid() bool
}

// This represents the number of SID in a SRGB range.
// Range returns a uint32
func (obj *isisSRSrgb) Range() uint32 {

	return *obj.obj.Range

}

// This represents the number of SID in a SRGB range.
// Range returns a uint32
func (obj *isisSRSrgb) HasRange() bool {
	return obj.obj.Range != nil
}

// This represents the number of SID in a SRGB range.
// SetRange sets the uint32 value in the IsisSRSrgb object
func (obj *isisSRSrgb) SetRange(value uint32) IsisSRSrgb {

	obj.obj.Range = &value
	return obj
}

// The SID/Label sub-TLV contains the first value of the SRGB while the range contains the number of SRGB elements.
// StartingSid returns a uint32
func (obj *isisSRSrgb) StartingSid() uint32 {

	return *obj.obj.StartingSid

}

// The SID/Label sub-TLV contains the first value of the SRGB while the range contains the number of SRGB elements.
// StartingSid returns a uint32
func (obj *isisSRSrgb) HasStartingSid() bool {
	return obj.obj.StartingSid != nil
}

// The SID/Label sub-TLV contains the first value of the SRGB while the range contains the number of SRGB elements.
// SetStartingSid sets the uint32 value in the IsisSRSrgb object
func (obj *isisSRSrgb) SetStartingSid(value uint32) IsisSRSrgb {

	obj.obj.StartingSid = &value
	return obj
}

func (obj *isisSRSrgb) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Range != nil {

		if *obj.obj.Range < 1 || *obj.obj.Range > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= IsisSRSrgb.Range <= 16777215 but Got %d", *obj.obj.Range))
		}

	}

	if obj.obj.StartingSid != nil {

		if *obj.obj.StartingSid < 1 || *obj.obj.StartingSid > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= IsisSRSrgb.StartingSid <= 16777215 but Got %d", *obj.obj.StartingSid))
		}

	}

}

func (obj *isisSRSrgb) setDefault() {
	if obj.obj.Range == nil {
		obj.SetRange(8000)
	}
	if obj.obj.StartingSid == nil {
		obj.SetStartingSid(16000)
	}

}
