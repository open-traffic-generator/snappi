package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspSrgb *****
type isisLspSrgb struct {
	validation
	obj          *otg.IsisLspSrgb
	marshaller   marshalIsisLspSrgb
	unMarshaller unMarshalIsisLspSrgb
}

func NewIsisLspSrgb() IsisLspSrgb {
	obj := isisLspSrgb{obj: &otg.IsisLspSrgb{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspSrgb) msg() *otg.IsisLspSrgb {
	return obj.obj
}

func (obj *isisLspSrgb) setMsg(msg *otg.IsisLspSrgb) IsisLspSrgb {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspSrgb struct {
	obj *isisLspSrgb
}

type marshalIsisLspSrgb interface {
	// ToProto marshals IsisLspSrgb to protobuf object *otg.IsisLspSrgb
	ToProto() (*otg.IsisLspSrgb, error)
	// ToPbText marshals IsisLspSrgb to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspSrgb to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspSrgb to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspSrgb to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspSrgb struct {
	obj *isisLspSrgb
}

type unMarshalIsisLspSrgb interface {
	// FromProto unmarshals IsisLspSrgb from protobuf object *otg.IsisLspSrgb
	FromProto(msg *otg.IsisLspSrgb) (IsisLspSrgb, error)
	// FromPbText unmarshals IsisLspSrgb from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspSrgb from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspSrgb from JSON text
	FromJson(value string) error
}

func (obj *isisLspSrgb) Marshal() marshalIsisLspSrgb {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspSrgb{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspSrgb) Unmarshal() unMarshalIsisLspSrgb {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspSrgb{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspSrgb) ToProto() (*otg.IsisLspSrgb, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspSrgb) FromProto(msg *otg.IsisLspSrgb) (IsisLspSrgb, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspSrgb) ToPbText() (string, error) {
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

func (m *unMarshalisisLspSrgb) FromPbText(value string) error {
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

func (m *marshalisisLspSrgb) ToYaml() (string, error) {
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

func (m *unMarshalisisLspSrgb) FromYaml(value string) error {
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

func (m *marshalisisLspSrgb) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalisisLspSrgb) ToJson() (string, error) {
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

func (m *unMarshalisisLspSrgb) FromJson(value string) error {
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

func (obj *isisLspSrgb) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspSrgb) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspSrgb) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspSrgb) Clone() (IsisLspSrgb, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspSrgb()
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

// IsisLspSrgb is this contains the propeties of SRGB range. Reference: https://datatracker.ietf.org/doc/html/rfc8667#section-3.1-7.1.1
type IsisLspSrgb interface {
	Validation
	// msg marshals IsisLspSrgb to protobuf object *otg.IsisLspSrgb
	// and doesn't set defaults
	msg() *otg.IsisLspSrgb
	// setMsg unmarshals IsisLspSrgb from protobuf object *otg.IsisLspSrgb
	// and doesn't set defaults
	setMsg(*otg.IsisLspSrgb) IsisLspSrgb
	// provides marshal interface
	Marshal() marshalIsisLspSrgb
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspSrgb
	// validate validates IsisLspSrgb
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspSrgb, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StartingSid returns uint32, set in IsisLspSrgb.
	StartingSid() uint32
	// SetStartingSid assigns uint32 provided by user to IsisLspSrgb
	SetStartingSid(value uint32) IsisLspSrgb
	// HasStartingSid checks if StartingSid has been set in IsisLspSrgb
	HasStartingSid() bool
	// Range returns uint32, set in IsisLspSrgb.
	Range() uint32
	// SetRange assigns uint32 provided by user to IsisLspSrgb
	SetRange(value uint32) IsisLspSrgb
	// HasRange checks if Range has been set in IsisLspSrgb
	HasRange() bool
}

// The SID/Label sub-TLV contains the first value of the SRGB while the range contains the number of SRGB elements.
// StartingSid returns a uint32
func (obj *isisLspSrgb) StartingSid() uint32 {

	return *obj.obj.StartingSid

}

// The SID/Label sub-TLV contains the first value of the SRGB while the range contains the number of SRGB elements.
// StartingSid returns a uint32
func (obj *isisLspSrgb) HasStartingSid() bool {
	return obj.obj.StartingSid != nil
}

// The SID/Label sub-TLV contains the first value of the SRGB while the range contains the number of SRGB elements.
// SetStartingSid sets the uint32 value in the IsisLspSrgb object
func (obj *isisLspSrgb) SetStartingSid(value uint32) IsisLspSrgb {

	obj.obj.StartingSid = &value
	return obj
}

// This represents the number of SID in a SRGB range.
// Range returns a uint32
func (obj *isisLspSrgb) Range() uint32 {

	return *obj.obj.Range

}

// This represents the number of SID in a SRGB range.
// Range returns a uint32
func (obj *isisLspSrgb) HasRange() bool {
	return obj.obj.Range != nil
}

// This represents the number of SID in a SRGB range.
// SetRange sets the uint32 value in the IsisLspSrgb object
func (obj *isisLspSrgb) SetRange(value uint32) IsisLspSrgb {

	obj.obj.Range = &value
	return obj
}

func (obj *isisLspSrgb) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspSrgb) setDefault() {

}
