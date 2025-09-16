package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpNLRIPrefixPathId *****
type bgpNLRIPrefixPathId struct {
	validation
	obj          *otg.BgpNLRIPrefixPathId
	marshaller   marshalBgpNLRIPrefixPathId
	unMarshaller unMarshalBgpNLRIPrefixPathId
}

func NewBgpNLRIPrefixPathId() BgpNLRIPrefixPathId {
	obj := bgpNLRIPrefixPathId{obj: &otg.BgpNLRIPrefixPathId{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpNLRIPrefixPathId) msg() *otg.BgpNLRIPrefixPathId {
	return obj.obj
}

func (obj *bgpNLRIPrefixPathId) setMsg(msg *otg.BgpNLRIPrefixPathId) BgpNLRIPrefixPathId {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpNLRIPrefixPathId struct {
	obj *bgpNLRIPrefixPathId
}

type marshalBgpNLRIPrefixPathId interface {
	// ToProto marshals BgpNLRIPrefixPathId to protobuf object *otg.BgpNLRIPrefixPathId
	ToProto() (*otg.BgpNLRIPrefixPathId, error)
	// ToPbText marshals BgpNLRIPrefixPathId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpNLRIPrefixPathId to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpNLRIPrefixPathId to JSON text
	ToJson() (string, error)
}

type unMarshalbgpNLRIPrefixPathId struct {
	obj *bgpNLRIPrefixPathId
}

type unMarshalBgpNLRIPrefixPathId interface {
	// FromProto unmarshals BgpNLRIPrefixPathId from protobuf object *otg.BgpNLRIPrefixPathId
	FromProto(msg *otg.BgpNLRIPrefixPathId) (BgpNLRIPrefixPathId, error)
	// FromPbText unmarshals BgpNLRIPrefixPathId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpNLRIPrefixPathId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpNLRIPrefixPathId from JSON text
	FromJson(value string) error
}

func (obj *bgpNLRIPrefixPathId) Marshal() marshalBgpNLRIPrefixPathId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpNLRIPrefixPathId{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpNLRIPrefixPathId) Unmarshal() unMarshalBgpNLRIPrefixPathId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpNLRIPrefixPathId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpNLRIPrefixPathId) ToProto() (*otg.BgpNLRIPrefixPathId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpNLRIPrefixPathId) FromProto(msg *otg.BgpNLRIPrefixPathId) (BgpNLRIPrefixPathId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpNLRIPrefixPathId) ToPbText() (string, error) {
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

func (m *unMarshalbgpNLRIPrefixPathId) FromPbText(value string) error {
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

func (m *marshalbgpNLRIPrefixPathId) ToYaml() (string, error) {
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

func (m *unMarshalbgpNLRIPrefixPathId) FromYaml(value string) error {
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

func (m *marshalbgpNLRIPrefixPathId) ToJson() (string, error) {
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

func (m *unMarshalbgpNLRIPrefixPathId) FromJson(value string) error {
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

func (obj *bgpNLRIPrefixPathId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpNLRIPrefixPathId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpNLRIPrefixPathId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpNLRIPrefixPathId) Clone() (BgpNLRIPrefixPathId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpNLRIPrefixPathId()
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

// BgpNLRIPrefixPathId is optional field in the NLRI carrying Path Id of the prefix.
type BgpNLRIPrefixPathId interface {
	Validation
	// msg marshals BgpNLRIPrefixPathId to protobuf object *otg.BgpNLRIPrefixPathId
	// and doesn't set defaults
	msg() *otg.BgpNLRIPrefixPathId
	// setMsg unmarshals BgpNLRIPrefixPathId from protobuf object *otg.BgpNLRIPrefixPathId
	// and doesn't set defaults
	setMsg(*otg.BgpNLRIPrefixPathId) BgpNLRIPrefixPathId
	// provides marshal interface
	Marshal() marshalBgpNLRIPrefixPathId
	// provides unmarshal interface
	Unmarshal() unMarshalBgpNLRIPrefixPathId
	// validate validates BgpNLRIPrefixPathId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpNLRIPrefixPathId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns uint32, set in BgpNLRIPrefixPathId.
	Value() uint32
	// SetValue assigns uint32 provided by user to BgpNLRIPrefixPathId
	SetValue(value uint32) BgpNLRIPrefixPathId
	// HasValue checks if Value has been set in BgpNLRIPrefixPathId
	HasValue() bool
}

// The value of the optional Path ID of the prefix.
// Value returns a uint32
func (obj *bgpNLRIPrefixPathId) Value() uint32 {

	return *obj.obj.Value

}

// The value of the optional Path ID of the prefix.
// Value returns a uint32
func (obj *bgpNLRIPrefixPathId) HasValue() bool {
	return obj.obj.Value != nil
}

// The value of the optional Path ID of the prefix.
// SetValue sets the uint32 value in the BgpNLRIPrefixPathId object
func (obj *bgpNLRIPrefixPathId) SetValue(value uint32) BgpNLRIPrefixPathId {

	obj.obj.Value = &value
	return obj
}

func (obj *bgpNLRIPrefixPathId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpNLRIPrefixPathId) setDefault() {
	if obj.obj.Value == nil {
		obj.SetValue(1)
	}

}
