package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSrPolicyPreference *****
type bgpAttributesSrPolicyPreference struct {
	validation
	obj          *otg.BgpAttributesSrPolicyPreference
	marshaller   marshalBgpAttributesSrPolicyPreference
	unMarshaller unMarshalBgpAttributesSrPolicyPreference
}

func NewBgpAttributesSrPolicyPreference() BgpAttributesSrPolicyPreference {
	obj := bgpAttributesSrPolicyPreference{obj: &otg.BgpAttributesSrPolicyPreference{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSrPolicyPreference) msg() *otg.BgpAttributesSrPolicyPreference {
	return obj.obj
}

func (obj *bgpAttributesSrPolicyPreference) setMsg(msg *otg.BgpAttributesSrPolicyPreference) BgpAttributesSrPolicyPreference {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSrPolicyPreference struct {
	obj *bgpAttributesSrPolicyPreference
}

type marshalBgpAttributesSrPolicyPreference interface {
	// ToProto marshals BgpAttributesSrPolicyPreference to protobuf object *otg.BgpAttributesSrPolicyPreference
	ToProto() (*otg.BgpAttributesSrPolicyPreference, error)
	// ToPbText marshals BgpAttributesSrPolicyPreference to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSrPolicyPreference to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSrPolicyPreference to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSrPolicyPreference to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSrPolicyPreference struct {
	obj *bgpAttributesSrPolicyPreference
}

type unMarshalBgpAttributesSrPolicyPreference interface {
	// FromProto unmarshals BgpAttributesSrPolicyPreference from protobuf object *otg.BgpAttributesSrPolicyPreference
	FromProto(msg *otg.BgpAttributesSrPolicyPreference) (BgpAttributesSrPolicyPreference, error)
	// FromPbText unmarshals BgpAttributesSrPolicyPreference from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSrPolicyPreference from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSrPolicyPreference from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSrPolicyPreference) Marshal() marshalBgpAttributesSrPolicyPreference {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSrPolicyPreference{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSrPolicyPreference) Unmarshal() unMarshalBgpAttributesSrPolicyPreference {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSrPolicyPreference{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSrPolicyPreference) ToProto() (*otg.BgpAttributesSrPolicyPreference, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSrPolicyPreference) FromProto(msg *otg.BgpAttributesSrPolicyPreference) (BgpAttributesSrPolicyPreference, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSrPolicyPreference) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyPreference) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSrPolicyPreference) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyPreference) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSrPolicyPreference) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSrPolicyPreference) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyPreference) FromJson(value string) error {
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

func (obj *bgpAttributesSrPolicyPreference) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSrPolicyPreference) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSrPolicyPreference) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSrPolicyPreference) Clone() (BgpAttributesSrPolicyPreference, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSrPolicyPreference()
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

// BgpAttributesSrPolicyPreference is optional Preference sub-tlv (Type 12) is used to select the best candidate path for an SR Policy.
// It is defined in Section 2.4.1 of draft-ietf-idr-sr-policy-safi-02 .

type BgpAttributesSrPolicyPreference interface {
	Validation
	// msg marshals BgpAttributesSrPolicyPreference to protobuf object *otg.BgpAttributesSrPolicyPreference
	// and doesn't set defaults
	msg() *otg.BgpAttributesSrPolicyPreference
	// setMsg unmarshals BgpAttributesSrPolicyPreference from protobuf object *otg.BgpAttributesSrPolicyPreference
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSrPolicyPreference) BgpAttributesSrPolicyPreference
	// provides marshal interface
	Marshal() marshalBgpAttributesSrPolicyPreference
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSrPolicyPreference
	// validate validates BgpAttributesSrPolicyPreference
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSrPolicyPreference, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns uint32, set in BgpAttributesSrPolicyPreference.
	Value() uint32
	// SetValue assigns uint32 provided by user to BgpAttributesSrPolicyPreference
	SetValue(value uint32) BgpAttributesSrPolicyPreference
	// HasValue checks if Value has been set in BgpAttributesSrPolicyPreference
	HasValue() bool
}

// Value to be carried in the Preference sub-tlv.
// Value returns a uint32
func (obj *bgpAttributesSrPolicyPreference) Value() uint32 {

	return *obj.obj.Value

}

// Value to be carried in the Preference sub-tlv.
// Value returns a uint32
func (obj *bgpAttributesSrPolicyPreference) HasValue() bool {
	return obj.obj.Value != nil
}

// Value to be carried in the Preference sub-tlv.
// SetValue sets the uint32 value in the BgpAttributesSrPolicyPreference object
func (obj *bgpAttributesSrPolicyPreference) SetValue(value uint32) BgpAttributesSrPolicyPreference {

	obj.obj.Value = &value
	return obj
}

func (obj *bgpAttributesSrPolicyPreference) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpAttributesSrPolicyPreference) setDefault() {
	if obj.obj.Value == nil {
		obj.SetValue(0)
	}

}
