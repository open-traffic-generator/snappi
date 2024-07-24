package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesLocalPreference *****
type bgpAttributesLocalPreference struct {
	validation
	obj          *otg.BgpAttributesLocalPreference
	marshaller   marshalBgpAttributesLocalPreference
	unMarshaller unMarshalBgpAttributesLocalPreference
}

func NewBgpAttributesLocalPreference() BgpAttributesLocalPreference {
	obj := bgpAttributesLocalPreference{obj: &otg.BgpAttributesLocalPreference{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesLocalPreference) msg() *otg.BgpAttributesLocalPreference {
	return obj.obj
}

func (obj *bgpAttributesLocalPreference) setMsg(msg *otg.BgpAttributesLocalPreference) BgpAttributesLocalPreference {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesLocalPreference struct {
	obj *bgpAttributesLocalPreference
}

type marshalBgpAttributesLocalPreference interface {
	// ToProto marshals BgpAttributesLocalPreference to protobuf object *otg.BgpAttributesLocalPreference
	ToProto() (*otg.BgpAttributesLocalPreference, error)
	// ToPbText marshals BgpAttributesLocalPreference to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesLocalPreference to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesLocalPreference to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesLocalPreference struct {
	obj *bgpAttributesLocalPreference
}

type unMarshalBgpAttributesLocalPreference interface {
	// FromProto unmarshals BgpAttributesLocalPreference from protobuf object *otg.BgpAttributesLocalPreference
	FromProto(msg *otg.BgpAttributesLocalPreference) (BgpAttributesLocalPreference, error)
	// FromPbText unmarshals BgpAttributesLocalPreference from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesLocalPreference from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesLocalPreference from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesLocalPreference) Marshal() marshalBgpAttributesLocalPreference {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesLocalPreference{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesLocalPreference) Unmarshal() unMarshalBgpAttributesLocalPreference {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesLocalPreference{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesLocalPreference) ToProto() (*otg.BgpAttributesLocalPreference, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesLocalPreference) FromProto(msg *otg.BgpAttributesLocalPreference) (BgpAttributesLocalPreference, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesLocalPreference) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesLocalPreference) FromPbText(value string) error {
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

func (m *marshalbgpAttributesLocalPreference) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesLocalPreference) FromYaml(value string) error {
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

func (m *marshalbgpAttributesLocalPreference) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesLocalPreference) FromJson(value string) error {
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

func (obj *bgpAttributesLocalPreference) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesLocalPreference) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesLocalPreference) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesLocalPreference) Clone() (BgpAttributesLocalPreference, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesLocalPreference()
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

// BgpAttributesLocalPreference is optional LOCAL_PREFERENCE attribute sent to the peer to indicate the degree of preference
// for externally learned routes.This should be included only for internal peers.It is
// used for the selection of the path for the traffic leaving the AS.The route with the
// highest local preference value is preferred.
type BgpAttributesLocalPreference interface {
	Validation
	// msg marshals BgpAttributesLocalPreference to protobuf object *otg.BgpAttributesLocalPreference
	// and doesn't set defaults
	msg() *otg.BgpAttributesLocalPreference
	// setMsg unmarshals BgpAttributesLocalPreference from protobuf object *otg.BgpAttributesLocalPreference
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesLocalPreference) BgpAttributesLocalPreference
	// provides marshal interface
	Marshal() marshalBgpAttributesLocalPreference
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesLocalPreference
	// validate validates BgpAttributesLocalPreference
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesLocalPreference, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns uint32, set in BgpAttributesLocalPreference.
	Value() uint32
	// SetValue assigns uint32 provided by user to BgpAttributesLocalPreference
	SetValue(value uint32) BgpAttributesLocalPreference
	// HasValue checks if Value has been set in BgpAttributesLocalPreference
	HasValue() bool
}

// Value to be set in the LOCAL_PREFERENCE attribute The multi exit discriminator (MED) value used for route selection sent to the peer.
// Value returns a uint32
func (obj *bgpAttributesLocalPreference) Value() uint32 {

	return *obj.obj.Value

}

// Value to be set in the LOCAL_PREFERENCE attribute The multi exit discriminator (MED) value used for route selection sent to the peer.
// Value returns a uint32
func (obj *bgpAttributesLocalPreference) HasValue() bool {
	return obj.obj.Value != nil
}

// Value to be set in the LOCAL_PREFERENCE attribute The multi exit discriminator (MED) value used for route selection sent to the peer.
// SetValue sets the uint32 value in the BgpAttributesLocalPreference object
func (obj *bgpAttributesLocalPreference) SetValue(value uint32) BgpAttributesLocalPreference {

	obj.obj.Value = &value
	return obj
}

func (obj *bgpAttributesLocalPreference) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpAttributesLocalPreference) setDefault() {
	if obj.obj.Value == nil {
		obj.SetValue(100)
	}

}
