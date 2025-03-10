package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesCustomCommunity *****
type bgpAttributesCustomCommunity struct {
	validation
	obj          *otg.BgpAttributesCustomCommunity
	marshaller   marshalBgpAttributesCustomCommunity
	unMarshaller unMarshalBgpAttributesCustomCommunity
}

func NewBgpAttributesCustomCommunity() BgpAttributesCustomCommunity {
	obj := bgpAttributesCustomCommunity{obj: &otg.BgpAttributesCustomCommunity{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesCustomCommunity) msg() *otg.BgpAttributesCustomCommunity {
	return obj.obj
}

func (obj *bgpAttributesCustomCommunity) setMsg(msg *otg.BgpAttributesCustomCommunity) BgpAttributesCustomCommunity {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesCustomCommunity struct {
	obj *bgpAttributesCustomCommunity
}

type marshalBgpAttributesCustomCommunity interface {
	// ToProto marshals BgpAttributesCustomCommunity to protobuf object *otg.BgpAttributesCustomCommunity
	ToProto() (*otg.BgpAttributesCustomCommunity, error)
	// ToPbText marshals BgpAttributesCustomCommunity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesCustomCommunity to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesCustomCommunity to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesCustomCommunity to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesCustomCommunity struct {
	obj *bgpAttributesCustomCommunity
}

type unMarshalBgpAttributesCustomCommunity interface {
	// FromProto unmarshals BgpAttributesCustomCommunity from protobuf object *otg.BgpAttributesCustomCommunity
	FromProto(msg *otg.BgpAttributesCustomCommunity) (BgpAttributesCustomCommunity, error)
	// FromPbText unmarshals BgpAttributesCustomCommunity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesCustomCommunity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesCustomCommunity from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesCustomCommunity) Marshal() marshalBgpAttributesCustomCommunity {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesCustomCommunity{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesCustomCommunity) Unmarshal() unMarshalBgpAttributesCustomCommunity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesCustomCommunity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesCustomCommunity) ToProto() (*otg.BgpAttributesCustomCommunity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesCustomCommunity) FromProto(msg *otg.BgpAttributesCustomCommunity) (BgpAttributesCustomCommunity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesCustomCommunity) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesCustomCommunity) FromPbText(value string) error {
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

func (m *marshalbgpAttributesCustomCommunity) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesCustomCommunity) FromYaml(value string) error {
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

func (m *marshalbgpAttributesCustomCommunity) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesCustomCommunity) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesCustomCommunity) FromJson(value string) error {
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

func (obj *bgpAttributesCustomCommunity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesCustomCommunity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesCustomCommunity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesCustomCommunity) Clone() (BgpAttributesCustomCommunity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesCustomCommunity()
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

// BgpAttributesCustomCommunity is user defined COMMUNITY attribute containing 2 byte AS and custom 2 byte value defined by the administrator of the domain.
type BgpAttributesCustomCommunity interface {
	Validation
	// msg marshals BgpAttributesCustomCommunity to protobuf object *otg.BgpAttributesCustomCommunity
	// and doesn't set defaults
	msg() *otg.BgpAttributesCustomCommunity
	// setMsg unmarshals BgpAttributesCustomCommunity from protobuf object *otg.BgpAttributesCustomCommunity
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesCustomCommunity) BgpAttributesCustomCommunity
	// provides marshal interface
	Marshal() marshalBgpAttributesCustomCommunity
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesCustomCommunity
	// validate validates BgpAttributesCustomCommunity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesCustomCommunity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AsNumber returns uint32, set in BgpAttributesCustomCommunity.
	AsNumber() uint32
	// SetAsNumber assigns uint32 provided by user to BgpAttributesCustomCommunity
	SetAsNumber(value uint32) BgpAttributesCustomCommunity
	// HasAsNumber checks if AsNumber has been set in BgpAttributesCustomCommunity
	HasAsNumber() bool
	// Custom returns string, set in BgpAttributesCustomCommunity.
	Custom() string
	// SetCustom assigns string provided by user to BgpAttributesCustomCommunity
	SetCustom(value string) BgpAttributesCustomCommunity
	// HasCustom checks if Custom has been set in BgpAttributesCustomCommunity
	HasCustom() bool
}

// First two octets of the community value containing a 2 byte AS number.
// AsNumber returns a uint32
func (obj *bgpAttributesCustomCommunity) AsNumber() uint32 {

	return *obj.obj.AsNumber

}

// First two octets of the community value containing a 2 byte AS number.
// AsNumber returns a uint32
func (obj *bgpAttributesCustomCommunity) HasAsNumber() bool {
	return obj.obj.AsNumber != nil
}

// First two octets of the community value containing a 2 byte AS number.
// SetAsNumber sets the uint32 value in the BgpAttributesCustomCommunity object
func (obj *bgpAttributesCustomCommunity) SetAsNumber(value uint32) BgpAttributesCustomCommunity {

	obj.obj.AsNumber = &value
	return obj
}

// Last two octets of the community value in hex.  If user provides less than 4 hex bytes, it should be left-padded with 0s.
// Custom returns a string
func (obj *bgpAttributesCustomCommunity) Custom() string {

	return *obj.obj.Custom

}

// Last two octets of the community value in hex.  If user provides less than 4 hex bytes, it should be left-padded with 0s.
// Custom returns a string
func (obj *bgpAttributesCustomCommunity) HasCustom() bool {
	return obj.obj.Custom != nil
}

// Last two octets of the community value in hex.  If user provides less than 4 hex bytes, it should be left-padded with 0s.
// SetCustom sets the string value in the BgpAttributesCustomCommunity object
func (obj *bgpAttributesCustomCommunity) SetCustom(value string) BgpAttributesCustomCommunity {

	obj.obj.Custom = &value
	return obj
}

func (obj *bgpAttributesCustomCommunity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.AsNumber != nil {

		if *obj.obj.AsNumber > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAttributesCustomCommunity.AsNumber <= 65535 but Got %d", *obj.obj.AsNumber))
		}

	}

	if obj.obj.Custom != nil {

		if len(*obj.obj.Custom) > 4 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of BgpAttributesCustomCommunity.Custom <= 4 but Got %d",
					len(*obj.obj.Custom)))
		}

	}

}

func (obj *bgpAttributesCustomCommunity) setDefault() {
	if obj.obj.AsNumber == nil {
		obj.SetAsNumber(0)
	}
	if obj.obj.Custom == nil {
		obj.SetCustom("0000")
	}

}
