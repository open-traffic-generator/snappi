package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2OpaqueLsaTlv *****
type ospfv2OpaqueLsaTlv struct {
	validation
	obj          *otg.Ospfv2OpaqueLsaTlv
	marshaller   marshalOspfv2OpaqueLsaTlv
	unMarshaller unMarshalOspfv2OpaqueLsaTlv
}

func NewOspfv2OpaqueLsaTlv() Ospfv2OpaqueLsaTlv {
	obj := ospfv2OpaqueLsaTlv{obj: &otg.Ospfv2OpaqueLsaTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2OpaqueLsaTlv) msg() *otg.Ospfv2OpaqueLsaTlv {
	return obj.obj
}

func (obj *ospfv2OpaqueLsaTlv) setMsg(msg *otg.Ospfv2OpaqueLsaTlv) Ospfv2OpaqueLsaTlv {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2OpaqueLsaTlv struct {
	obj *ospfv2OpaqueLsaTlv
}

type marshalOspfv2OpaqueLsaTlv interface {
	// ToProto marshals Ospfv2OpaqueLsaTlv to protobuf object *otg.Ospfv2OpaqueLsaTlv
	ToProto() (*otg.Ospfv2OpaqueLsaTlv, error)
	// ToPbText marshals Ospfv2OpaqueLsaTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2OpaqueLsaTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2OpaqueLsaTlv to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2OpaqueLsaTlv struct {
	obj *ospfv2OpaqueLsaTlv
}

type unMarshalOspfv2OpaqueLsaTlv interface {
	// FromProto unmarshals Ospfv2OpaqueLsaTlv from protobuf object *otg.Ospfv2OpaqueLsaTlv
	FromProto(msg *otg.Ospfv2OpaqueLsaTlv) (Ospfv2OpaqueLsaTlv, error)
	// FromPbText unmarshals Ospfv2OpaqueLsaTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2OpaqueLsaTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2OpaqueLsaTlv from JSON text
	FromJson(value string) error
}

func (obj *ospfv2OpaqueLsaTlv) Marshal() marshalOspfv2OpaqueLsaTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2OpaqueLsaTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2OpaqueLsaTlv) Unmarshal() unMarshalOspfv2OpaqueLsaTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2OpaqueLsaTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2OpaqueLsaTlv) ToProto() (*otg.Ospfv2OpaqueLsaTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2OpaqueLsaTlv) FromProto(msg *otg.Ospfv2OpaqueLsaTlv) (Ospfv2OpaqueLsaTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2OpaqueLsaTlv) ToPbText() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaTlv) FromPbText(value string) error {
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

func (m *marshalospfv2OpaqueLsaTlv) ToYaml() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaTlv) FromYaml(value string) error {
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

func (m *marshalospfv2OpaqueLsaTlv) ToJson() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaTlv) FromJson(value string) error {
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

func (obj *ospfv2OpaqueLsaTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2OpaqueLsaTlv) Clone() (Ospfv2OpaqueLsaTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2OpaqueLsaTlv()
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

// Ospfv2OpaqueLsaTlv is a top-level TLV carried in the body of an OSPFv2 Opaque LSA that is not decoded
// into a structured field elsewhere in the model, reported in the generic wire format
// every OSPFv2 Opaque LSA TLV shares: a 2-octet Type, a 2-octet Length and the Value
// octets (RFC 7770 Section 2.3).
// type is the numeric wire value, so a TLV type this model has never heard of is
// still fully representable and its content is never lost. The Value is reported as
// one opaque blob, including for a TLV whose Value is itself a sub-TLV sequence:
// nothing in the TLV header says the Value is a sub-TLV sequence, and a TLV reported
// here is by definition one this model does not decode.
type Ospfv2OpaqueLsaTlv interface {
	Validation
	// msg marshals Ospfv2OpaqueLsaTlv to protobuf object *otg.Ospfv2OpaqueLsaTlv
	// and doesn't set defaults
	msg() *otg.Ospfv2OpaqueLsaTlv
	// setMsg unmarshals Ospfv2OpaqueLsaTlv from protobuf object *otg.Ospfv2OpaqueLsaTlv
	// and doesn't set defaults
	setMsg(*otg.Ospfv2OpaqueLsaTlv) Ospfv2OpaqueLsaTlv
	// provides marshal interface
	Marshal() marshalOspfv2OpaqueLsaTlv
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2OpaqueLsaTlv
	// validate validates Ospfv2OpaqueLsaTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2OpaqueLsaTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns uint32, set in Ospfv2OpaqueLsaTlv.
	Type() uint32
	// SetType assigns uint32 provided by user to Ospfv2OpaqueLsaTlv
	SetType(value uint32) Ospfv2OpaqueLsaTlv
	// HasType checks if Type has been set in Ospfv2OpaqueLsaTlv
	HasType() bool
	// Length returns uint32, set in Ospfv2OpaqueLsaTlv.
	Length() uint32
	// SetLength assigns uint32 provided by user to Ospfv2OpaqueLsaTlv
	SetLength(value uint32) Ospfv2OpaqueLsaTlv
	// HasLength checks if Length has been set in Ospfv2OpaqueLsaTlv
	HasLength() bool
	// Value returns string, set in Ospfv2OpaqueLsaTlv.
	Value() string
	// SetValue assigns string provided by user to Ospfv2OpaqueLsaTlv
	SetValue(value string) Ospfv2OpaqueLsaTlv
	// HasValue checks if Value has been set in Ospfv2OpaqueLsaTlv
	HasValue() bool
}

// The TLV Type field, as the numeric value carried on the wire. Its meaning is
// scoped by the parent LSA's tlv_information.choice; the authoritative list of
// assigned TLV types is the IANA OSPFv2 parameters registry:
// https://www.iana.org/assignments/ospfv2-parameters/ospfv2-parameters.xhtml
// Type returns a uint32
func (obj *ospfv2OpaqueLsaTlv) Type() uint32 {

	return *obj.obj.Type

}

// The TLV Type field, as the numeric value carried on the wire. Its meaning is
// scoped by the parent LSA's tlv_information.choice; the authoritative list of
// assigned TLV types is the IANA OSPFv2 parameters registry:
// https://www.iana.org/assignments/ospfv2-parameters/ospfv2-parameters.xhtml
// Type returns a uint32
func (obj *ospfv2OpaqueLsaTlv) HasType() bool {
	return obj.obj.Type != nil
}

// The TLV Type field, as the numeric value carried on the wire. Its meaning is
// scoped by the parent LSA's tlv_information.choice; the authoritative list of
// assigned TLV types is the IANA OSPFv2 parameters registry:
// https://www.iana.org/assignments/ospfv2-parameters/ospfv2-parameters.xhtml
// SetType sets the uint32 value in the Ospfv2OpaqueLsaTlv object
func (obj *ospfv2OpaqueLsaTlv) SetType(value uint32) Ospfv2OpaqueLsaTlv {

	obj.obj.Type = &value
	return obj
}

// The TLV Length field, in octets, of the value field.
// Length returns a uint32
func (obj *ospfv2OpaqueLsaTlv) Length() uint32 {

	return *obj.obj.Length

}

// The TLV Length field, in octets, of the value field.
// Length returns a uint32
func (obj *ospfv2OpaqueLsaTlv) HasLength() bool {
	return obj.obj.Length != nil
}

// The TLV Length field, in octets, of the value field.
// SetLength sets the uint32 value in the Ospfv2OpaqueLsaTlv object
func (obj *ospfv2OpaqueLsaTlv) SetLength(value uint32) Ospfv2OpaqueLsaTlv {

	obj.obj.Length = &value
	return obj
}

// The raw byte contents of the TLV Value field, as hex characters. Two hex characters per octet, so the string is twice length characters long.
// Value returns a string
func (obj *ospfv2OpaqueLsaTlv) Value() string {

	return *obj.obj.Value

}

// The raw byte contents of the TLV Value field, as hex characters. Two hex characters per octet, so the string is twice length characters long.
// Value returns a string
func (obj *ospfv2OpaqueLsaTlv) HasValue() bool {
	return obj.obj.Value != nil
}

// The raw byte contents of the TLV Value field, as hex characters. Two hex characters per octet, so the string is twice length characters long.
// SetValue sets the string value in the Ospfv2OpaqueLsaTlv object
func (obj *ospfv2OpaqueLsaTlv) SetValue(value string) Ospfv2OpaqueLsaTlv {

	obj.obj.Value = &value
	return obj
}

func (obj *ospfv2OpaqueLsaTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		if *obj.obj.Type > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv2OpaqueLsaTlv.Type <= 65535 but Got %d", *obj.obj.Type))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv2OpaqueLsaTlv.Length <= 65535 but Got %d", *obj.obj.Length))
		}

	}

	if obj.obj.Value != nil {

		err := obj.validateHex(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2OpaqueLsaTlv.Value"))
		}

	}

}

func (obj *ospfv2OpaqueLsaTlv) setDefault() {

}
