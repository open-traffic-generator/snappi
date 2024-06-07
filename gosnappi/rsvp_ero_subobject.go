package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpEroSubobject *****
type rsvpEroSubobject struct {
	validation
	obj          *otg.RsvpEroSubobject
	marshaller   marshalRsvpEroSubobject
	unMarshaller unMarshalRsvpEroSubobject
}

func NewRsvpEroSubobject() RsvpEroSubobject {
	obj := rsvpEroSubobject{obj: &otg.RsvpEroSubobject{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpEroSubobject) msg() *otg.RsvpEroSubobject {
	return obj.obj
}

func (obj *rsvpEroSubobject) setMsg(msg *otg.RsvpEroSubobject) RsvpEroSubobject {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpEroSubobject struct {
	obj *rsvpEroSubobject
}

type marshalRsvpEroSubobject interface {
	// ToProto marshals RsvpEroSubobject to protobuf object *otg.RsvpEroSubobject
	ToProto() (*otg.RsvpEroSubobject, error)
	// ToPbText marshals RsvpEroSubobject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpEroSubobject to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpEroSubobject to JSON text
	ToJson() (string, error)
}

type unMarshalrsvpEroSubobject struct {
	obj *rsvpEroSubobject
}

type unMarshalRsvpEroSubobject interface {
	// FromProto unmarshals RsvpEroSubobject from protobuf object *otg.RsvpEroSubobject
	FromProto(msg *otg.RsvpEroSubobject) (RsvpEroSubobject, error)
	// FromPbText unmarshals RsvpEroSubobject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpEroSubobject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpEroSubobject from JSON text
	FromJson(value string) error
}

func (obj *rsvpEroSubobject) Marshal() marshalRsvpEroSubobject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpEroSubobject{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpEroSubobject) Unmarshal() unMarshalRsvpEroSubobject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpEroSubobject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpEroSubobject) ToProto() (*otg.RsvpEroSubobject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpEroSubobject) FromProto(msg *otg.RsvpEroSubobject) (RsvpEroSubobject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpEroSubobject) ToPbText() (string, error) {
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

func (m *unMarshalrsvpEroSubobject) FromPbText(value string) error {
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

func (m *marshalrsvpEroSubobject) ToYaml() (string, error) {
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

func (m *unMarshalrsvpEroSubobject) FromYaml(value string) error {
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

func (m *marshalrsvpEroSubobject) ToJson() (string, error) {
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

func (m *unMarshalrsvpEroSubobject) FromJson(value string) error {
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

func (obj *rsvpEroSubobject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpEroSubobject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpEroSubobject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpEroSubobject) Clone() (RsvpEroSubobject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpEroSubobject()
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

// RsvpEroSubobject is configuration for the ERO sub-object.
type RsvpEroSubobject interface {
	Validation
	// msg marshals RsvpEroSubobject to protobuf object *otg.RsvpEroSubobject
	// and doesn't set defaults
	msg() *otg.RsvpEroSubobject
	// setMsg unmarshals RsvpEroSubobject from protobuf object *otg.RsvpEroSubobject
	// and doesn't set defaults
	setMsg(*otg.RsvpEroSubobject) RsvpEroSubobject
	// provides marshal interface
	Marshal() marshalRsvpEroSubobject
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpEroSubobject
	// validate validates RsvpEroSubobject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpEroSubobject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns RsvpEroSubobjectTypeEnum, set in RsvpEroSubobject
	Type() RsvpEroSubobjectTypeEnum
	// SetType assigns RsvpEroSubobjectTypeEnum provided by user to RsvpEroSubobject
	SetType(value RsvpEroSubobjectTypeEnum) RsvpEroSubobject
	// HasType checks if Type has been set in RsvpEroSubobject
	HasType() bool
	// Ipv4Address returns string, set in RsvpEroSubobject.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to RsvpEroSubobject
	SetIpv4Address(value string) RsvpEroSubobject
	// HasIpv4Address checks if Ipv4Address has been set in RsvpEroSubobject
	HasIpv4Address() bool
	// PrefixLength returns uint32, set in RsvpEroSubobject.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to RsvpEroSubobject
	SetPrefixLength(value uint32) RsvpEroSubobject
	// HasPrefixLength checks if PrefixLength has been set in RsvpEroSubobject
	HasPrefixLength() bool
	// AsNumber returns uint32, set in RsvpEroSubobject.
	AsNumber() uint32
	// SetAsNumber assigns uint32 provided by user to RsvpEroSubobject
	SetAsNumber(value uint32) RsvpEroSubobject
	// HasAsNumber checks if AsNumber has been set in RsvpEroSubobject
	HasAsNumber() bool
	// HopType returns RsvpEroSubobjectHopTypeEnum, set in RsvpEroSubobject
	HopType() RsvpEroSubobjectHopTypeEnum
	// SetHopType assigns RsvpEroSubobjectHopTypeEnum provided by user to RsvpEroSubobject
	SetHopType(value RsvpEroSubobjectHopTypeEnum) RsvpEroSubobject
	// HasHopType checks if HopType has been set in RsvpEroSubobject
	HasHopType() bool
}

type RsvpEroSubobjectTypeEnum string

// Enum of Type on RsvpEroSubobject
var RsvpEroSubobjectType = struct {
	IPV4      RsvpEroSubobjectTypeEnum
	AS_NUMBER RsvpEroSubobjectTypeEnum
}{
	IPV4:      RsvpEroSubobjectTypeEnum("ipv4"),
	AS_NUMBER: RsvpEroSubobjectTypeEnum("as_number"),
}

func (obj *rsvpEroSubobject) Type() RsvpEroSubobjectTypeEnum {
	return RsvpEroSubobjectTypeEnum(obj.obj.Type.Enum().String())
}

// The type of the ERO sub-object, one of IPv4 Address or AS Number.
// Type returns a string
func (obj *rsvpEroSubobject) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *rsvpEroSubobject) SetType(value RsvpEroSubobjectTypeEnum) RsvpEroSubobject {
	intValue, ok := otg.RsvpEroSubobject_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RsvpEroSubobjectTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.RsvpEroSubobject_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// IPv4 address that this LSP should traverse through. This field is applicable only if the value of 'type' is set to 'ipv4'.
// Ipv4Address returns a string
func (obj *rsvpEroSubobject) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// IPv4 address that this LSP should traverse through. This field is applicable only if the value of 'type' is set to 'ipv4'.
// Ipv4Address returns a string
func (obj *rsvpEroSubobject) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// IPv4 address that this LSP should traverse through. This field is applicable only if the value of 'type' is set to 'ipv4'.
// SetIpv4Address sets the string value in the RsvpEroSubobject object
func (obj *rsvpEroSubobject) SetIpv4Address(value string) RsvpEroSubobject {

	obj.obj.Ipv4Address = &value
	return obj
}

// Prefix length for the IPv4 address in the ERO sub-object. This field is applicable only if the value of 'type' is set to 'ipv4'.
// PrefixLength returns a uint32
func (obj *rsvpEroSubobject) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// Prefix length for the IPv4 address in the ERO sub-object. This field is applicable only if the value of 'type' is set to 'ipv4'.
// PrefixLength returns a uint32
func (obj *rsvpEroSubobject) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// Prefix length for the IPv4 address in the ERO sub-object. This field is applicable only if the value of 'type' is set to 'ipv4'.
// SetPrefixLength sets the uint32 value in the RsvpEroSubobject object
func (obj *rsvpEroSubobject) SetPrefixLength(value uint32) RsvpEroSubobject {

	obj.obj.PrefixLength = &value
	return obj
}

// Autonomous System number to be set in the ERO sub-object that this LSP should traverse through. This field is applicable only if the value of 'type' is set to 'as_number'. Note that as per RFC3209, 4-byte AS encoding is not supported.
// AsNumber returns a uint32
func (obj *rsvpEroSubobject) AsNumber() uint32 {

	return *obj.obj.AsNumber

}

// Autonomous System number to be set in the ERO sub-object that this LSP should traverse through. This field is applicable only if the value of 'type' is set to 'as_number'. Note that as per RFC3209, 4-byte AS encoding is not supported.
// AsNumber returns a uint32
func (obj *rsvpEroSubobject) HasAsNumber() bool {
	return obj.obj.AsNumber != nil
}

// Autonomous System number to be set in the ERO sub-object that this LSP should traverse through. This field is applicable only if the value of 'type' is set to 'as_number'. Note that as per RFC3209, 4-byte AS encoding is not supported.
// SetAsNumber sets the uint32 value in the RsvpEroSubobject object
func (obj *rsvpEroSubobject) SetAsNumber(value uint32) RsvpEroSubobject {

	obj.obj.AsNumber = &value
	return obj
}

type RsvpEroSubobjectHopTypeEnum string

// Enum of HopType on RsvpEroSubobject
var RsvpEroSubobjectHopType = struct {
	STRICT RsvpEroSubobjectHopTypeEnum
	LOOSE  RsvpEroSubobjectHopTypeEnum
}{
	STRICT: RsvpEroSubobjectHopTypeEnum("strict"),
	LOOSE:  RsvpEroSubobjectHopTypeEnum("loose"),
}

func (obj *rsvpEroSubobject) HopType() RsvpEroSubobjectHopTypeEnum {
	return RsvpEroSubobjectHopTypeEnum(obj.obj.HopType.Enum().String())
}

// The hop type of the ERO sub-object, one of Strict or Loose.
// HopType returns a string
func (obj *rsvpEroSubobject) HasHopType() bool {
	return obj.obj.HopType != nil
}

func (obj *rsvpEroSubobject) SetHopType(value RsvpEroSubobjectHopTypeEnum) RsvpEroSubobject {
	intValue, ok := otg.RsvpEroSubobject_HopType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RsvpEroSubobjectHopTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.RsvpEroSubobject_HopType_Enum(intValue)
	obj.obj.HopType = &enumValue

	return obj
}

func (obj *rsvpEroSubobject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4Address != nil {

		err := obj.validateIpv4(obj.Ipv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RsvpEroSubobject.Ipv4Address"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpEroSubobject.PrefixLength <= 32 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.AsNumber != nil {

		if *obj.obj.AsNumber > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpEroSubobject.AsNumber <= 65535 but Got %d", *obj.obj.AsNumber))
		}

	}

}

func (obj *rsvpEroSubobject) setDefault() {
	if obj.obj.Type == nil {
		obj.SetType(RsvpEroSubobjectType.IPV4)

	}
	if obj.obj.Ipv4Address == nil {
		obj.SetIpv4Address("0.0.0.0")
	}
	if obj.obj.PrefixLength == nil {
		obj.SetPrefixLength(32)
	}
	if obj.obj.AsNumber == nil {
		obj.SetAsNumber(0)
	}
	if obj.obj.HopType == nil {
		obj.SetHopType(RsvpEroSubobjectHopType.LOOSE)

	}

}
