package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrtePreferenceSubTlv *****
type bgpSrtePreferenceSubTlv struct {
	validation
	obj          *otg.BgpSrtePreferenceSubTlv
	marshaller   marshalBgpSrtePreferenceSubTlv
	unMarshaller unMarshalBgpSrtePreferenceSubTlv
}

func NewBgpSrtePreferenceSubTlv() BgpSrtePreferenceSubTlv {
	obj := bgpSrtePreferenceSubTlv{obj: &otg.BgpSrtePreferenceSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrtePreferenceSubTlv) msg() *otg.BgpSrtePreferenceSubTlv {
	return obj.obj
}

func (obj *bgpSrtePreferenceSubTlv) setMsg(msg *otg.BgpSrtePreferenceSubTlv) BgpSrtePreferenceSubTlv {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrtePreferenceSubTlv struct {
	obj *bgpSrtePreferenceSubTlv
}

type marshalBgpSrtePreferenceSubTlv interface {
	// ToProto marshals BgpSrtePreferenceSubTlv to protobuf object *otg.BgpSrtePreferenceSubTlv
	ToProto() (*otg.BgpSrtePreferenceSubTlv, error)
	// ToPbText marshals BgpSrtePreferenceSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrtePreferenceSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrtePreferenceSubTlv to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpSrtePreferenceSubTlv to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpSrtePreferenceSubTlv struct {
	obj *bgpSrtePreferenceSubTlv
}

type unMarshalBgpSrtePreferenceSubTlv interface {
	// FromProto unmarshals BgpSrtePreferenceSubTlv from protobuf object *otg.BgpSrtePreferenceSubTlv
	FromProto(msg *otg.BgpSrtePreferenceSubTlv) (BgpSrtePreferenceSubTlv, error)
	// FromPbText unmarshals BgpSrtePreferenceSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrtePreferenceSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrtePreferenceSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrtePreferenceSubTlv) Marshal() marshalBgpSrtePreferenceSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrtePreferenceSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrtePreferenceSubTlv) Unmarshal() unMarshalBgpSrtePreferenceSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrtePreferenceSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrtePreferenceSubTlv) ToProto() (*otg.BgpSrtePreferenceSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrtePreferenceSubTlv) FromProto(msg *otg.BgpSrtePreferenceSubTlv) (BgpSrtePreferenceSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrtePreferenceSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrtePreferenceSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrtePreferenceSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrtePreferenceSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrtePreferenceSubTlv) ToJsonRaw() (string, error) {
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

func (m *marshalbgpSrtePreferenceSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrtePreferenceSubTlv) FromJson(value string) error {
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

func (obj *bgpSrtePreferenceSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrtePreferenceSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrtePreferenceSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrtePreferenceSubTlv) Clone() (BgpSrtePreferenceSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrtePreferenceSubTlv()
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

// BgpSrtePreferenceSubTlv is configuration for BGP preference sub TLV of the SR Policy candidate path.
type BgpSrtePreferenceSubTlv interface {
	Validation
	// msg marshals BgpSrtePreferenceSubTlv to protobuf object *otg.BgpSrtePreferenceSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrtePreferenceSubTlv
	// setMsg unmarshals BgpSrtePreferenceSubTlv from protobuf object *otg.BgpSrtePreferenceSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrtePreferenceSubTlv) BgpSrtePreferenceSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrtePreferenceSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrtePreferenceSubTlv
	// validate validates BgpSrtePreferenceSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrtePreferenceSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Preference returns uint32, set in BgpSrtePreferenceSubTlv.
	Preference() uint32
	// SetPreference assigns uint32 provided by user to BgpSrtePreferenceSubTlv
	SetPreference(value uint32) BgpSrtePreferenceSubTlv
	// HasPreference checks if Preference has been set in BgpSrtePreferenceSubTlv
	HasPreference() bool
}

// The preference value of the SR Policy candidate path.
// Preference returns a uint32
func (obj *bgpSrtePreferenceSubTlv) Preference() uint32 {

	return *obj.obj.Preference

}

// The preference value of the SR Policy candidate path.
// Preference returns a uint32
func (obj *bgpSrtePreferenceSubTlv) HasPreference() bool {
	return obj.obj.Preference != nil
}

// The preference value of the SR Policy candidate path.
// SetPreference sets the uint32 value in the BgpSrtePreferenceSubTlv object
func (obj *bgpSrtePreferenceSubTlv) SetPreference(value uint32) BgpSrtePreferenceSubTlv {

	obj.obj.Preference = &value
	return obj
}

func (obj *bgpSrtePreferenceSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpSrtePreferenceSubTlv) setDefault() {
	if obj.obj.Preference == nil {
		obj.SetPreference(0)
	}

}
