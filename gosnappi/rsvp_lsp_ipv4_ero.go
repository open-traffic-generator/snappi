package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpLspIpv4Ero *****
type rsvpLspIpv4Ero struct {
	validation
	obj          *otg.RsvpLspIpv4Ero
	marshaller   marshalRsvpLspIpv4Ero
	unMarshaller unMarshalRsvpLspIpv4Ero
}

func NewRsvpLspIpv4Ero() RsvpLspIpv4Ero {
	obj := rsvpLspIpv4Ero{obj: &otg.RsvpLspIpv4Ero{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpLspIpv4Ero) msg() *otg.RsvpLspIpv4Ero {
	return obj.obj
}

func (obj *rsvpLspIpv4Ero) setMsg(msg *otg.RsvpLspIpv4Ero) RsvpLspIpv4Ero {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpLspIpv4Ero struct {
	obj *rsvpLspIpv4Ero
}

type marshalRsvpLspIpv4Ero interface {
	// ToProto marshals RsvpLspIpv4Ero to protobuf object *otg.RsvpLspIpv4Ero
	ToProto() (*otg.RsvpLspIpv4Ero, error)
	// ToPbText marshals RsvpLspIpv4Ero to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpLspIpv4Ero to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpLspIpv4Ero to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals RsvpLspIpv4Ero to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrsvpLspIpv4Ero struct {
	obj *rsvpLspIpv4Ero
}

type unMarshalRsvpLspIpv4Ero interface {
	// FromProto unmarshals RsvpLspIpv4Ero from protobuf object *otg.RsvpLspIpv4Ero
	FromProto(msg *otg.RsvpLspIpv4Ero) (RsvpLspIpv4Ero, error)
	// FromPbText unmarshals RsvpLspIpv4Ero from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpLspIpv4Ero from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpLspIpv4Ero from JSON text
	FromJson(value string) error
}

func (obj *rsvpLspIpv4Ero) Marshal() marshalRsvpLspIpv4Ero {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpLspIpv4Ero{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpLspIpv4Ero) Unmarshal() unMarshalRsvpLspIpv4Ero {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpLspIpv4Ero{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpLspIpv4Ero) ToProto() (*otg.RsvpLspIpv4Ero, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpLspIpv4Ero) FromProto(msg *otg.RsvpLspIpv4Ero) (RsvpLspIpv4Ero, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpLspIpv4Ero) ToPbText() (string, error) {
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

func (m *unMarshalrsvpLspIpv4Ero) FromPbText(value string) error {
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

func (m *marshalrsvpLspIpv4Ero) ToYaml() (string, error) {
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

func (m *unMarshalrsvpLspIpv4Ero) FromYaml(value string) error {
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

func (m *marshalrsvpLspIpv4Ero) ToJsonRaw() (string, error) {
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

func (m *marshalrsvpLspIpv4Ero) ToJson() (string, error) {
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

func (m *unMarshalrsvpLspIpv4Ero) FromJson(value string) error {
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

func (obj *rsvpLspIpv4Ero) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpLspIpv4Ero) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpLspIpv4Ero) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpLspIpv4Ero) Clone() (RsvpLspIpv4Ero, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpLspIpv4Ero()
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

// RsvpLspIpv4Ero is this contains the list of sub-objects included in the Explicit Route Object(ERO) object send in the PATH message from the ingress. These sub-objects contain the intermediate hops to be traversed by the LSP while being forwarded  towards the egress endpoint.
type RsvpLspIpv4Ero interface {
	Validation
	// msg marshals RsvpLspIpv4Ero to protobuf object *otg.RsvpLspIpv4Ero
	// and doesn't set defaults
	msg() *otg.RsvpLspIpv4Ero
	// setMsg unmarshals RsvpLspIpv4Ero from protobuf object *otg.RsvpLspIpv4Ero
	// and doesn't set defaults
	setMsg(*otg.RsvpLspIpv4Ero) RsvpLspIpv4Ero
	// provides marshal interface
	Marshal() marshalRsvpLspIpv4Ero
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpLspIpv4Ero
	// validate validates RsvpLspIpv4Ero
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpLspIpv4Ero, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Prefix returns string, set in RsvpLspIpv4Ero.
	Prefix() string
	// SetPrefix assigns string provided by user to RsvpLspIpv4Ero
	SetPrefix(value string) RsvpLspIpv4Ero
	// HasPrefix checks if Prefix has been set in RsvpLspIpv4Ero
	HasPrefix() bool
	// Asn returns uint32, set in RsvpLspIpv4Ero.
	Asn() uint32
	// SetAsn assigns uint32 provided by user to RsvpLspIpv4Ero
	SetAsn(value uint32) RsvpLspIpv4Ero
	// HasAsn checks if Asn has been set in RsvpLspIpv4Ero
	HasAsn() bool
	// Type returns RsvpLspIpv4EroTypeEnum, set in RsvpLspIpv4Ero
	Type() RsvpLspIpv4EroTypeEnum
	// SetType assigns RsvpLspIpv4EroTypeEnum provided by user to RsvpLspIpv4Ero
	SetType(value RsvpLspIpv4EroTypeEnum) RsvpLspIpv4Ero
	// HasType checks if Type has been set in RsvpLspIpv4Ero
	HasType() bool
}

// The IPv4 prefix indicated by the ERO. Specified only when the ERO hop is an IPv4 prefix.
// Prefix returns a string
func (obj *rsvpLspIpv4Ero) Prefix() string {

	return *obj.obj.Prefix

}

// The IPv4 prefix indicated by the ERO. Specified only when the ERO hop is an IPv4 prefix.
// Prefix returns a string
func (obj *rsvpLspIpv4Ero) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The IPv4 prefix indicated by the ERO. Specified only when the ERO hop is an IPv4 prefix.
// SetPrefix sets the string value in the RsvpLspIpv4Ero object
func (obj *rsvpLspIpv4Ero) SetPrefix(value string) RsvpLspIpv4Ero {

	obj.obj.Prefix = &value
	return obj
}

// The autonomous system number indicated by the ERO. Specified only when the ERO hop is an  2 or 4-byte AS number.
// Asn returns a uint32
func (obj *rsvpLspIpv4Ero) Asn() uint32 {

	return *obj.obj.Asn

}

// The autonomous system number indicated by the ERO. Specified only when the ERO hop is an  2 or 4-byte AS number.
// Asn returns a uint32
func (obj *rsvpLspIpv4Ero) HasAsn() bool {
	return obj.obj.Asn != nil
}

// The autonomous system number indicated by the ERO. Specified only when the ERO hop is an  2 or 4-byte AS number.
// SetAsn sets the uint32 value in the RsvpLspIpv4Ero object
func (obj *rsvpLspIpv4Ero) SetAsn(value uint32) RsvpLspIpv4Ero {

	obj.obj.Asn = &value
	return obj
}

type RsvpLspIpv4EroTypeEnum string

// Enum of Type on RsvpLspIpv4Ero
var RsvpLspIpv4EroType = struct {
	IPV4                 RsvpLspIpv4EroTypeEnum
	IPV6                 RsvpLspIpv4EroTypeEnum
	ASN                  RsvpLspIpv4EroTypeEnum
	ASN4                 RsvpLspIpv4EroTypeEnum
	LABEL                RsvpLspIpv4EroTypeEnum
	UNNUMBERED_INTERFACE RsvpLspIpv4EroTypeEnum
}{
	IPV4:                 RsvpLspIpv4EroTypeEnum("ipv4"),
	IPV6:                 RsvpLspIpv4EroTypeEnum("ipv6"),
	ASN:                  RsvpLspIpv4EroTypeEnum("asn"),
	ASN4:                 RsvpLspIpv4EroTypeEnum("asn4"),
	LABEL:                RsvpLspIpv4EroTypeEnum("label"),
	UNNUMBERED_INTERFACE: RsvpLspIpv4EroTypeEnum("unnumbered_interface"),
}

func (obj *rsvpLspIpv4Ero) Type() RsvpLspIpv4EroTypeEnum {
	return RsvpLspIpv4EroTypeEnum(obj.obj.Type.Enum().String())
}

// The type indicated by the ERO.
// Type returns a string
func (obj *rsvpLspIpv4Ero) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *rsvpLspIpv4Ero) SetType(value RsvpLspIpv4EroTypeEnum) RsvpLspIpv4Ero {
	intValue, ok := otg.RsvpLspIpv4Ero_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RsvpLspIpv4EroTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.RsvpLspIpv4Ero_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

func (obj *rsvpLspIpv4Ero) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Prefix != nil {

		err := obj.validateIpv4(obj.Prefix())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RsvpLspIpv4Ero.Prefix"))
		}

	}

}

func (obj *rsvpLspIpv4Ero) setDefault() {

}
