package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpCommunity *****
type bgpCommunity struct {
	validation
	obj          *otg.BgpCommunity
	marshaller   marshalBgpCommunity
	unMarshaller unMarshalBgpCommunity
}

func NewBgpCommunity() BgpCommunity {
	obj := bgpCommunity{obj: &otg.BgpCommunity{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpCommunity) msg() *otg.BgpCommunity {
	return obj.obj
}

func (obj *bgpCommunity) setMsg(msg *otg.BgpCommunity) BgpCommunity {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpCommunity struct {
	obj *bgpCommunity
}

type marshalBgpCommunity interface {
	// ToProto marshals BgpCommunity to protobuf object *otg.BgpCommunity
	ToProto() (*otg.BgpCommunity, error)
	// ToPbText marshals BgpCommunity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpCommunity to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpCommunity to JSON text
	ToJson() (string, error)
}

type unMarshalbgpCommunity struct {
	obj *bgpCommunity
}

type unMarshalBgpCommunity interface {
	// FromProto unmarshals BgpCommunity from protobuf object *otg.BgpCommunity
	FromProto(msg *otg.BgpCommunity) (BgpCommunity, error)
	// FromPbText unmarshals BgpCommunity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpCommunity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpCommunity from JSON text
	FromJson(value string) error
}

func (obj *bgpCommunity) Marshal() marshalBgpCommunity {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpCommunity{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpCommunity) Unmarshal() unMarshalBgpCommunity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpCommunity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpCommunity) ToProto() (*otg.BgpCommunity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpCommunity) FromProto(msg *otg.BgpCommunity) (BgpCommunity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpCommunity) ToPbText() (string, error) {
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

func (m *unMarshalbgpCommunity) FromPbText(value string) error {
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

func (m *marshalbgpCommunity) ToYaml() (string, error) {
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

func (m *unMarshalbgpCommunity) FromYaml(value string) error {
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

func (m *marshalbgpCommunity) ToJson() (string, error) {
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

func (m *unMarshalbgpCommunity) FromJson(value string) error {
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

func (obj *bgpCommunity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpCommunity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpCommunity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpCommunity) Clone() (BgpCommunity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpCommunity()
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

// BgpCommunity is bGP communities provide additional capability for tagging routes and  for modifying BGP routing policy on upstream and downstream routers. BGP community is a 32-bit number which is broken into 16-bit AS number and  a 16-bit custom value.
type BgpCommunity interface {
	Validation
	// msg marshals BgpCommunity to protobuf object *otg.BgpCommunity
	// and doesn't set defaults
	msg() *otg.BgpCommunity
	// setMsg unmarshals BgpCommunity from protobuf object *otg.BgpCommunity
	// and doesn't set defaults
	setMsg(*otg.BgpCommunity) BgpCommunity
	// provides marshal interface
	Marshal() marshalBgpCommunity
	// provides unmarshal interface
	Unmarshal() unMarshalBgpCommunity
	// validate validates BgpCommunity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpCommunity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns BgpCommunityTypeEnum, set in BgpCommunity
	Type() BgpCommunityTypeEnum
	// SetType assigns BgpCommunityTypeEnum provided by user to BgpCommunity
	SetType(value BgpCommunityTypeEnum) BgpCommunity
	// HasType checks if Type has been set in BgpCommunity
	HasType() bool
	// AsNumber returns uint32, set in BgpCommunity.
	AsNumber() uint32
	// SetAsNumber assigns uint32 provided by user to BgpCommunity
	SetAsNumber(value uint32) BgpCommunity
	// HasAsNumber checks if AsNumber has been set in BgpCommunity
	HasAsNumber() bool
	// AsCustom returns uint32, set in BgpCommunity.
	AsCustom() uint32
	// SetAsCustom assigns uint32 provided by user to BgpCommunity
	SetAsCustom(value uint32) BgpCommunity
	// HasAsCustom checks if AsCustom has been set in BgpCommunity
	HasAsCustom() bool
}

type BgpCommunityTypeEnum string

// Enum of Type on BgpCommunity
var BgpCommunityType = struct {
	MANUAL_AS_NUMBER    BgpCommunityTypeEnum
	NO_EXPORT           BgpCommunityTypeEnum
	NO_ADVERTISED       BgpCommunityTypeEnum
	NO_EXPORT_SUBCONFED BgpCommunityTypeEnum
	LLGR_STALE          BgpCommunityTypeEnum
	NO_LLGR             BgpCommunityTypeEnum
}{
	MANUAL_AS_NUMBER:    BgpCommunityTypeEnum("manual_as_number"),
	NO_EXPORT:           BgpCommunityTypeEnum("no_export"),
	NO_ADVERTISED:       BgpCommunityTypeEnum("no_advertised"),
	NO_EXPORT_SUBCONFED: BgpCommunityTypeEnum("no_export_subconfed"),
	LLGR_STALE:          BgpCommunityTypeEnum("llgr_stale"),
	NO_LLGR:             BgpCommunityTypeEnum("no_llgr"),
}

func (obj *bgpCommunity) Type() BgpCommunityTypeEnum {
	return BgpCommunityTypeEnum(obj.obj.Type.Enum().String())
}

// The type of community AS number.
// Type returns a string
func (obj *bgpCommunity) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *bgpCommunity) SetType(value BgpCommunityTypeEnum) BgpCommunity {
	intValue, ok := otg.BgpCommunity_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpCommunityTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpCommunity_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// First two octets of 32 bit community AS number.
// AsNumber returns a uint32
func (obj *bgpCommunity) AsNumber() uint32 {

	return *obj.obj.AsNumber

}

// First two octets of 32 bit community AS number.
// AsNumber returns a uint32
func (obj *bgpCommunity) HasAsNumber() bool {
	return obj.obj.AsNumber != nil
}

// First two octets of 32 bit community AS number.
// SetAsNumber sets the uint32 value in the BgpCommunity object
func (obj *bgpCommunity) SetAsNumber(value uint32) BgpCommunity {

	obj.obj.AsNumber = &value
	return obj
}

// Last two octets of the community value.
// AsCustom returns a uint32
func (obj *bgpCommunity) AsCustom() uint32 {

	return *obj.obj.AsCustom

}

// Last two octets of the community value.
// AsCustom returns a uint32
func (obj *bgpCommunity) HasAsCustom() bool {
	return obj.obj.AsCustom != nil
}

// Last two octets of the community value.
// SetAsCustom sets the uint32 value in the BgpCommunity object
func (obj *bgpCommunity) SetAsCustom(value uint32) BgpCommunity {

	obj.obj.AsCustom = &value
	return obj
}

func (obj *bgpCommunity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.AsNumber != nil {

		if *obj.obj.AsNumber > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpCommunity.AsNumber <= 65535 but Got %d", *obj.obj.AsNumber))
		}

	}

	if obj.obj.AsCustom != nil {

		if *obj.obj.AsCustom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpCommunity.AsCustom <= 65535 but Got %d", *obj.obj.AsCustom))
		}

	}

}

func (obj *bgpCommunity) setDefault() {
	if obj.obj.AsNumber == nil {
		obj.SetAsNumber(0)
	}
	if obj.obj.AsCustom == nil {
		obj.SetAsCustom(0)
	}

}
