package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAddPath *****
type bgpAddPath struct {
	validation
	obj          *otg.BgpAddPath
	marshaller   marshalBgpAddPath
	unMarshaller unMarshalBgpAddPath
}

func NewBgpAddPath() BgpAddPath {
	obj := bgpAddPath{obj: &otg.BgpAddPath{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAddPath) msg() *otg.BgpAddPath {
	return obj.obj
}

func (obj *bgpAddPath) setMsg(msg *otg.BgpAddPath) BgpAddPath {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAddPath struct {
	obj *bgpAddPath
}

type marshalBgpAddPath interface {
	// ToProto marshals BgpAddPath to protobuf object *otg.BgpAddPath
	ToProto() (*otg.BgpAddPath, error)
	// ToPbText marshals BgpAddPath to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAddPath to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAddPath to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAddPath to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAddPath struct {
	obj *bgpAddPath
}

type unMarshalBgpAddPath interface {
	// FromProto unmarshals BgpAddPath from protobuf object *otg.BgpAddPath
	FromProto(msg *otg.BgpAddPath) (BgpAddPath, error)
	// FromPbText unmarshals BgpAddPath from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAddPath from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAddPath from JSON text
	FromJson(value string) error
}

func (obj *bgpAddPath) Marshal() marshalBgpAddPath {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAddPath{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAddPath) Unmarshal() unMarshalBgpAddPath {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAddPath{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAddPath) ToProto() (*otg.BgpAddPath, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAddPath) FromProto(msg *otg.BgpAddPath) (BgpAddPath, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAddPath) ToPbText() (string, error) {
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

func (m *unMarshalbgpAddPath) FromPbText(value string) error {
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

func (m *marshalbgpAddPath) ToYaml() (string, error) {
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

func (m *unMarshalbgpAddPath) FromYaml(value string) error {
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

func (m *marshalbgpAddPath) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAddPath) ToJson() (string, error) {
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

func (m *unMarshalbgpAddPath) FromJson(value string) error {
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

func (obj *bgpAddPath) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAddPath) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAddPath) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAddPath) Clone() (BgpAddPath, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAddPath()
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

// BgpAddPath is the BGP Additional Paths feature is a BGP extension that allows the  advertisement of multiple paths for the same prefix without the new  paths implicitly replacing any previous paths.
type BgpAddPath interface {
	Validation
	// msg marshals BgpAddPath to protobuf object *otg.BgpAddPath
	// and doesn't set defaults
	msg() *otg.BgpAddPath
	// setMsg unmarshals BgpAddPath from protobuf object *otg.BgpAddPath
	// and doesn't set defaults
	setMsg(*otg.BgpAddPath) BgpAddPath
	// provides marshal interface
	Marshal() marshalBgpAddPath
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAddPath
	// validate validates BgpAddPath
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAddPath, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PathId returns uint32, set in BgpAddPath.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpAddPath
	SetPathId(value uint32) BgpAddPath
	// HasPathId checks if PathId has been set in BgpAddPath
	HasPathId() bool
}

// The id of the additional path.
// PathId returns a uint32
func (obj *bgpAddPath) PathId() uint32 {

	return *obj.obj.PathId

}

// The id of the additional path.
// PathId returns a uint32
func (obj *bgpAddPath) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The id of the additional path.
// SetPathId sets the uint32 value in the BgpAddPath object
func (obj *bgpAddPath) SetPathId(value uint32) BgpAddPath {

	obj.obj.PathId = &value
	return obj
}

func (obj *bgpAddPath) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpAddPath) setDefault() {
	if obj.obj.PathId == nil {
		obj.SetPathId(1)
	}

}
