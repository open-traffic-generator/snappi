package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtCommunity *****
type bgpExtCommunity struct {
	validation
	obj          *otg.BgpExtCommunity
	marshaller   marshalBgpExtCommunity
	unMarshaller unMarshalBgpExtCommunity
}

func NewBgpExtCommunity() BgpExtCommunity {
	obj := bgpExtCommunity{obj: &otg.BgpExtCommunity{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtCommunity) msg() *otg.BgpExtCommunity {
	return obj.obj
}

func (obj *bgpExtCommunity) setMsg(msg *otg.BgpExtCommunity) BgpExtCommunity {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtCommunity struct {
	obj *bgpExtCommunity
}

type marshalBgpExtCommunity interface {
	// ToProto marshals BgpExtCommunity to protobuf object *otg.BgpExtCommunity
	ToProto() (*otg.BgpExtCommunity, error)
	// ToPbText marshals BgpExtCommunity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtCommunity to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtCommunity to JSON text
	ToJson() (string, error)
}

type unMarshalbgpExtCommunity struct {
	obj *bgpExtCommunity
}

type unMarshalBgpExtCommunity interface {
	// FromProto unmarshals BgpExtCommunity from protobuf object *otg.BgpExtCommunity
	FromProto(msg *otg.BgpExtCommunity) (BgpExtCommunity, error)
	// FromPbText unmarshals BgpExtCommunity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtCommunity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtCommunity from JSON text
	FromJson(value string) error
}

func (obj *bgpExtCommunity) Marshal() marshalBgpExtCommunity {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtCommunity{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtCommunity) Unmarshal() unMarshalBgpExtCommunity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtCommunity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtCommunity) ToProto() (*otg.BgpExtCommunity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtCommunity) FromProto(msg *otg.BgpExtCommunity) (BgpExtCommunity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtCommunity) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtCommunity) FromPbText(value string) error {
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

func (m *marshalbgpExtCommunity) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtCommunity) FromYaml(value string) error {
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

func (m *marshalbgpExtCommunity) ToJson() (string, error) {
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

func (m *unMarshalbgpExtCommunity) FromJson(value string) error {
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

func (obj *bgpExtCommunity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtCommunity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtCommunity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtCommunity) Clone() (BgpExtCommunity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtCommunity()
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

// BgpExtCommunity is the Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an 8-Bytes value. It is divided into two main parts. The first 2 Bytes of the community encode a type and sub-type fields and the last 6 Bytes carry a unique set of data in a  format defined by the type and sub-type field. Extended communities provide a larger  range for grouping or categorizing communities.
type BgpExtCommunity interface {
	Validation
	// msg marshals BgpExtCommunity to protobuf object *otg.BgpExtCommunity
	// and doesn't set defaults
	msg() *otg.BgpExtCommunity
	// setMsg unmarshals BgpExtCommunity from protobuf object *otg.BgpExtCommunity
	// and doesn't set defaults
	setMsg(*otg.BgpExtCommunity) BgpExtCommunity
	// provides marshal interface
	Marshal() marshalBgpExtCommunity
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtCommunity
	// validate validates BgpExtCommunity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtCommunity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns BgpExtCommunityTypeEnum, set in BgpExtCommunity
	Type() BgpExtCommunityTypeEnum
	// SetType assigns BgpExtCommunityTypeEnum provided by user to BgpExtCommunity
	SetType(value BgpExtCommunityTypeEnum) BgpExtCommunity
	// HasType checks if Type has been set in BgpExtCommunity
	HasType() bool
	// Subtype returns BgpExtCommunitySubtypeEnum, set in BgpExtCommunity
	Subtype() BgpExtCommunitySubtypeEnum
	// SetSubtype assigns BgpExtCommunitySubtypeEnum provided by user to BgpExtCommunity
	SetSubtype(value BgpExtCommunitySubtypeEnum) BgpExtCommunity
	// HasSubtype checks if Subtype has been set in BgpExtCommunity
	HasSubtype() bool
	// Value returns string, set in BgpExtCommunity.
	Value() string
	// SetValue assigns string provided by user to BgpExtCommunity
	SetValue(value string) BgpExtCommunity
	// HasValue checks if Value has been set in BgpExtCommunity
	HasValue() bool
}

type BgpExtCommunityTypeEnum string

// Enum of Type on BgpExtCommunity
var BgpExtCommunityType = struct {
	ADMINISTRATOR_AS_2OCTET                BgpExtCommunityTypeEnum
	ADMINISTRATOR_IPV4_ADDRESS             BgpExtCommunityTypeEnum
	ADMINISTRATOR_AS_4OCTET                BgpExtCommunityTypeEnum
	OPAQUE                                 BgpExtCommunityTypeEnum
	EVPN                                   BgpExtCommunityTypeEnum
	ADMINISTRATOR_AS_2OCTET_LINK_BANDWIDTH BgpExtCommunityTypeEnum
}{
	ADMINISTRATOR_AS_2OCTET:                BgpExtCommunityTypeEnum("administrator_as_2octet"),
	ADMINISTRATOR_IPV4_ADDRESS:             BgpExtCommunityTypeEnum("administrator_ipv4_address"),
	ADMINISTRATOR_AS_4OCTET:                BgpExtCommunityTypeEnum("administrator_as_4octet"),
	OPAQUE:                                 BgpExtCommunityTypeEnum("opaque"),
	EVPN:                                   BgpExtCommunityTypeEnum("evpn"),
	ADMINISTRATOR_AS_2OCTET_LINK_BANDWIDTH: BgpExtCommunityTypeEnum("administrator_as_2octet_link_bandwidth"),
}

func (obj *bgpExtCommunity) Type() BgpExtCommunityTypeEnum {
	return BgpExtCommunityTypeEnum(obj.obj.Type.Enum().String())
}

// Extended Community Type field of 1 Byte.
// - administrator_as_2octet: Two-Octet AS Specific Extended Community (RFC 4360).
// - administrator_ipv4_address: IPv4 Address Specific Extended Community (RFC 4360).
// - administrator_as_4octet:  4-Octet AS Specific Extended Community (RFC 5668).
// - opaque: Opaque Extended Community (RFC 7432).
// - evpn: EVPN Extended Community (RFC 7153).
// - administrator_as_2octet_link_bandwidth : Link Bandwidth Extended Community (RFC 7153).
// Type returns a string
func (obj *bgpExtCommunity) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *bgpExtCommunity) SetType(value BgpExtCommunityTypeEnum) BgpExtCommunity {
	intValue, ok := otg.BgpExtCommunity_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpExtCommunityTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpExtCommunity_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

type BgpExtCommunitySubtypeEnum string

// Enum of Subtype on BgpExtCommunity
var BgpExtCommunitySubtype = struct {
	ROUTE_TARGET       BgpExtCommunitySubtypeEnum
	ORIGIN             BgpExtCommunitySubtypeEnum
	EXTENDED_BANDWIDTH BgpExtCommunitySubtypeEnum
	COLOR              BgpExtCommunitySubtypeEnum
	ENCAPSULATION      BgpExtCommunitySubtypeEnum
	MAC_ADDRESS        BgpExtCommunitySubtypeEnum
}{
	ROUTE_TARGET:       BgpExtCommunitySubtypeEnum("route_target"),
	ORIGIN:             BgpExtCommunitySubtypeEnum("origin"),
	EXTENDED_BANDWIDTH: BgpExtCommunitySubtypeEnum("extended_bandwidth"),
	COLOR:              BgpExtCommunitySubtypeEnum("color"),
	ENCAPSULATION:      BgpExtCommunitySubtypeEnum("encapsulation"),
	MAC_ADDRESS:        BgpExtCommunitySubtypeEnum("mac_address"),
}

func (obj *bgpExtCommunity) Subtype() BgpExtCommunitySubtypeEnum {
	return BgpExtCommunitySubtypeEnum(obj.obj.Subtype.Enum().String())
}

// Extended Community Sub Type field of 1 Byte.
// - route_target: Route Target.
// - origin: Origin.
// - extended_bandwidth: Specifies the link bandwidth.
// - color: Specifies the color value.
// - encapsulation: Specifies the Encapsulation Extended Community.
// - mac_address: Specifies the Extended community MAC address.
// Subtype returns a string
func (obj *bgpExtCommunity) HasSubtype() bool {
	return obj.obj.Subtype != nil
}

func (obj *bgpExtCommunity) SetSubtype(value BgpExtCommunitySubtypeEnum) BgpExtCommunity {
	intValue, ok := otg.BgpExtCommunity_Subtype_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpExtCommunitySubtypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpExtCommunity_Subtype_Enum(intValue)
	obj.obj.Subtype = &enumValue

	return obj
}

// Extended Community value of 6 Bytes. Example - for the Opaque type and Color subtype value can be '0000000000c8'  for the color value 200.
// Value returns a string
func (obj *bgpExtCommunity) Value() string {

	return *obj.obj.Value

}

// Extended Community value of 6 Bytes. Example - for the Opaque type and Color subtype value can be '0000000000c8'  for the color value 200.
// Value returns a string
func (obj *bgpExtCommunity) HasValue() bool {
	return obj.obj.Value != nil
}

// Extended Community value of 6 Bytes. Example - for the Opaque type and Color subtype value can be '0000000000c8'  for the color value 200.
// SetValue sets the string value in the BgpExtCommunity object
func (obj *bgpExtCommunity) SetValue(value string) BgpExtCommunity {

	obj.obj.Value = &value
	return obj
}

func (obj *bgpExtCommunity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateHex(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpExtCommunity.Value"))
		}

	}

}

func (obj *bgpExtCommunity) setDefault() {

}
