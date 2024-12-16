package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRSrlb *****
type isisSRSrlb struct {
	validation
	obj          *otg.IsisSRSrlb
	marshaller   marshalIsisSRSrlb
	unMarshaller unMarshalIsisSRSrlb
}

func NewIsisSRSrlb() IsisSRSrlb {
	obj := isisSRSrlb{obj: &otg.IsisSRSrlb{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRSrlb) msg() *otg.IsisSRSrlb {
	return obj.obj
}

func (obj *isisSRSrlb) setMsg(msg *otg.IsisSRSrlb) IsisSRSrlb {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRSrlb struct {
	obj *isisSRSrlb
}

type marshalIsisSRSrlb interface {
	// ToProto marshals IsisSRSrlb to protobuf object *otg.IsisSRSrlb
	ToProto() (*otg.IsisSRSrlb, error)
	// ToPbText marshals IsisSRSrlb to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRSrlb to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRSrlb to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRSrlb struct {
	obj *isisSRSrlb
}

type unMarshalIsisSRSrlb interface {
	// FromProto unmarshals IsisSRSrlb from protobuf object *otg.IsisSRSrlb
	FromProto(msg *otg.IsisSRSrlb) (IsisSRSrlb, error)
	// FromPbText unmarshals IsisSRSrlb from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRSrlb from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRSrlb from JSON text
	FromJson(value string) error
}

func (obj *isisSRSrlb) Marshal() marshalIsisSRSrlb {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRSrlb{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRSrlb) Unmarshal() unMarshalIsisSRSrlb {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRSrlb{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRSrlb) ToProto() (*otg.IsisSRSrlb, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRSrlb) FromProto(msg *otg.IsisSRSrlb) (IsisSRSrlb, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRSrlb) ToPbText() (string, error) {
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

func (m *unMarshalisisSRSrlb) FromPbText(value string) error {
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

func (m *marshalisisSRSrlb) ToYaml() (string, error) {
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

func (m *unMarshalisisSRSrlb) FromYaml(value string) error {
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

func (m *marshalisisSRSrlb) ToJson() (string, error) {
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

func (m *unMarshalisisSRSrlb) FromJson(value string) error {
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

func (obj *isisSRSrlb) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRSrlb) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRSrlb) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRSrlb) Clone() (IsisSRSrlb, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRSrlb()
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

// IsisSRSrlb is the SR Local Block (SRLB) sub-TLV contains the range of labels the node has reserved for Local SIDs.
// Local SIDs are used for Adjacency SIDs or locally significant Prefix SIDs and may also be allocated by components other than the IS-IS protocol.
// As an example, an application or a controller may instruct the router to allocate a specific Local SID. Therefore,
// in order for such applications or controllers to know what Local SIDs are available in the router,
// it is required that the router advertises its SRLB.
// Reference: https://datatracker.ietf.org/doc/html/rfc8667#name-sr-local-block-sub-tlv.
type IsisSRSrlb interface {
	Validation
	// msg marshals IsisSRSrlb to protobuf object *otg.IsisSRSrlb
	// and doesn't set defaults
	msg() *otg.IsisSRSrlb
	// setMsg unmarshals IsisSRSrlb from protobuf object *otg.IsisSRSrlb
	// and doesn't set defaults
	setMsg(*otg.IsisSRSrlb) IsisSRSrlb
	// provides marshal interface
	Marshal() marshalIsisSRSrlb
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRSrlb
	// validate validates IsisSRSrlb
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRSrlb, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StartingSid returns uint32, set in IsisSRSrlb.
	StartingSid() uint32
	// SetStartingSid assigns uint32 provided by user to IsisSRSrlb
	SetStartingSid(value uint32) IsisSRSrlb
	// HasStartingSid checks if StartingSid has been set in IsisSRSrlb
	HasStartingSid() bool
	// Range returns uint32, set in IsisSRSrlb.
	Range() uint32
	// SetRange assigns uint32 provided by user to IsisSRSrlb
	SetRange(value uint32) IsisSRSrlb
	// HasRange checks if Range has been set in IsisSRSrlb
	HasRange() bool
}

// The SID/Label sub-TLV contains the first value of the SRLB while the range contains the number of SRLB elements.
// StartingSid returns a uint32
func (obj *isisSRSrlb) StartingSid() uint32 {

	return *obj.obj.StartingSid

}

// The SID/Label sub-TLV contains the first value of the SRLB while the range contains the number of SRLB elements.
// StartingSid returns a uint32
func (obj *isisSRSrlb) HasStartingSid() bool {
	return obj.obj.StartingSid != nil
}

// The SID/Label sub-TLV contains the first value of the SRLB while the range contains the number of SRLB elements.
// SetStartingSid sets the uint32 value in the IsisSRSrlb object
func (obj *isisSRSrlb) SetStartingSid(value uint32) IsisSRSrlb {

	obj.obj.StartingSid = &value
	return obj
}

// This represents the number of SIDs in a SRGB range.
// Range returns a uint32
func (obj *isisSRSrlb) Range() uint32 {

	return *obj.obj.Range

}

// This represents the number of SIDs in a SRGB range.
// Range returns a uint32
func (obj *isisSRSrlb) HasRange() bool {
	return obj.obj.Range != nil
}

// This represents the number of SIDs in a SRGB range.
// SetRange sets the uint32 value in the IsisSRSrlb object
func (obj *isisSRSrlb) SetRange(value uint32) IsisSRSrlb {

	obj.obj.Range = &value
	return obj
}

func (obj *isisSRSrlb) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StartingSid != nil {

		if *obj.obj.StartingSid < 1 || *obj.obj.StartingSid > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= IsisSRSrlb.StartingSid <= 16777215 but Got %d", *obj.obj.StartingSid))
		}

	}

	if obj.obj.Range != nil {

		if *obj.obj.Range < 1 || *obj.obj.Range > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= IsisSRSrlb.Range <= 16777215 but Got %d", *obj.obj.Range))
		}

	}

}

func (obj *isisSRSrlb) setDefault() {
	if obj.obj.StartingSid == nil {
		obj.SetStartingSid(16000)
	}
	if obj.obj.Range == nil {
		obj.SetRange(8000)
	}

}
