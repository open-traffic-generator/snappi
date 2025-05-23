package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspSrlb *****
type isisLspSrlb struct {
	validation
	obj          *otg.IsisLspSrlb
	marshaller   marshalIsisLspSrlb
	unMarshaller unMarshalIsisLspSrlb
}

func NewIsisLspSrlb() IsisLspSrlb {
	obj := isisLspSrlb{obj: &otg.IsisLspSrlb{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspSrlb) msg() *otg.IsisLspSrlb {
	return obj.obj
}

func (obj *isisLspSrlb) setMsg(msg *otg.IsisLspSrlb) IsisLspSrlb {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspSrlb struct {
	obj *isisLspSrlb
}

type marshalIsisLspSrlb interface {
	// ToProto marshals IsisLspSrlb to protobuf object *otg.IsisLspSrlb
	ToProto() (*otg.IsisLspSrlb, error)
	// ToPbText marshals IsisLspSrlb to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspSrlb to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspSrlb to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspSrlb to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspSrlb struct {
	obj *isisLspSrlb
}

type unMarshalIsisLspSrlb interface {
	// FromProto unmarshals IsisLspSrlb from protobuf object *otg.IsisLspSrlb
	FromProto(msg *otg.IsisLspSrlb) (IsisLspSrlb, error)
	// FromPbText unmarshals IsisLspSrlb from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspSrlb from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspSrlb from JSON text
	FromJson(value string) error
}

func (obj *isisLspSrlb) Marshal() marshalIsisLspSrlb {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspSrlb{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspSrlb) Unmarshal() unMarshalIsisLspSrlb {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspSrlb{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspSrlb) ToProto() (*otg.IsisLspSrlb, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspSrlb) FromProto(msg *otg.IsisLspSrlb) (IsisLspSrlb, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspSrlb) ToPbText() (string, error) {
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

func (m *unMarshalisisLspSrlb) FromPbText(value string) error {
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

func (m *marshalisisLspSrlb) ToYaml() (string, error) {
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

func (m *unMarshalisisLspSrlb) FromYaml(value string) error {
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

func (m *marshalisisLspSrlb) ToJsonRaw() (string, error) {
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

func (m *marshalisisLspSrlb) ToJson() (string, error) {
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

func (m *unMarshalisisLspSrlb) FromJson(value string) error {
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

func (obj *isisLspSrlb) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspSrlb) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspSrlb) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspSrlb) Clone() (IsisLspSrlb, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspSrlb()
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

// IsisLspSrlb is this contains the propeties of SRLB. The SR Local Block (SRLB) sub-TLV contains the range of labels the node has reserved  for Local SIDs. Local SIDs are used, e.g., for Adj-SIDs, and may also be allocated by components other than the IS-IS protocol Reference: https://datatracker.ietf.org/doc/html/rfc8667#name-sr-local-block-sub-tlv.
type IsisLspSrlb interface {
	Validation
	// msg marshals IsisLspSrlb to protobuf object *otg.IsisLspSrlb
	// and doesn't set defaults
	msg() *otg.IsisLspSrlb
	// setMsg unmarshals IsisLspSrlb from protobuf object *otg.IsisLspSrlb
	// and doesn't set defaults
	setMsg(*otg.IsisLspSrlb) IsisLspSrlb
	// provides marshal interface
	Marshal() marshalIsisLspSrlb
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspSrlb
	// validate validates IsisLspSrlb
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspSrlb, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StartingSid returns uint32, set in IsisLspSrlb.
	StartingSid() uint32
	// SetStartingSid assigns uint32 provided by user to IsisLspSrlb
	SetStartingSid(value uint32) IsisLspSrlb
	// HasStartingSid checks if StartingSid has been set in IsisLspSrlb
	HasStartingSid() bool
	// Range returns uint32, set in IsisLspSrlb.
	Range() uint32
	// SetRange assigns uint32 provided by user to IsisLspSrlb
	SetRange(value uint32) IsisLspSrlb
	// HasRange checks if Range has been set in IsisLspSrlb
	HasRange() bool
}

// The SID/Label sub-TLV contains the first value of the SRGB while the range contains the number of SRGB elements.
// StartingSid returns a uint32
func (obj *isisLspSrlb) StartingSid() uint32 {

	return *obj.obj.StartingSid

}

// The SID/Label sub-TLV contains the first value of the SRGB while the range contains the number of SRGB elements.
// StartingSid returns a uint32
func (obj *isisLspSrlb) HasStartingSid() bool {
	return obj.obj.StartingSid != nil
}

// The SID/Label sub-TLV contains the first value of the SRGB while the range contains the number of SRGB elements.
// SetStartingSid sets the uint32 value in the IsisLspSrlb object
func (obj *isisLspSrlb) SetStartingSid(value uint32) IsisLspSrlb {

	obj.obj.StartingSid = &value
	return obj
}

// This represents the number of SID in a SRGB range.
// Range returns a uint32
func (obj *isisLspSrlb) Range() uint32 {

	return *obj.obj.Range

}

// This represents the number of SID in a SRGB range.
// Range returns a uint32
func (obj *isisLspSrlb) HasRange() bool {
	return obj.obj.Range != nil
}

// This represents the number of SID in a SRGB range.
// SetRange sets the uint32 value in the IsisLspSrlb object
func (obj *isisLspSrlb) SetRange(value uint32) IsisLspSrlb {

	obj.obj.Range = &value
	return obj
}

func (obj *isisLspSrlb) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspSrlb) setDefault() {

}
