package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CapabilitiesPath *****
type capabilitiesPath struct {
	validation
	obj          *otg.CapabilitiesPath
	marshaller   marshalCapabilitiesPath
	unMarshaller unMarshalCapabilitiesPath
}

func NewCapabilitiesPath() CapabilitiesPath {
	obj := capabilitiesPath{obj: &otg.CapabilitiesPath{}}
	obj.setDefault()
	return &obj
}

func (obj *capabilitiesPath) msg() *otg.CapabilitiesPath {
	return obj.obj
}

func (obj *capabilitiesPath) setMsg(msg *otg.CapabilitiesPath) CapabilitiesPath {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcapabilitiesPath struct {
	obj *capabilitiesPath
}

type marshalCapabilitiesPath interface {
	// ToProto marshals CapabilitiesPath to protobuf object *otg.CapabilitiesPath
	ToProto() (*otg.CapabilitiesPath, error)
	// ToPbText marshals CapabilitiesPath to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CapabilitiesPath to YAML text
	ToYaml() (string, error)
	// ToJson marshals CapabilitiesPath to JSON text
	ToJson() (string, error)
}

type unMarshalcapabilitiesPath struct {
	obj *capabilitiesPath
}

type unMarshalCapabilitiesPath interface {
	// FromProto unmarshals CapabilitiesPath from protobuf object *otg.CapabilitiesPath
	FromProto(msg *otg.CapabilitiesPath) (CapabilitiesPath, error)
	// FromPbText unmarshals CapabilitiesPath from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CapabilitiesPath from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CapabilitiesPath from JSON text
	FromJson(value string) error
}

func (obj *capabilitiesPath) Marshal() marshalCapabilitiesPath {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcapabilitiesPath{obj: obj}
	}
	return obj.marshaller
}

func (obj *capabilitiesPath) Unmarshal() unMarshalCapabilitiesPath {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcapabilitiesPath{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcapabilitiesPath) ToProto() (*otg.CapabilitiesPath, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcapabilitiesPath) FromProto(msg *otg.CapabilitiesPath) (CapabilitiesPath, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcapabilitiesPath) ToPbText() (string, error) {
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

func (m *unMarshalcapabilitiesPath) FromPbText(value string) error {
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

func (m *marshalcapabilitiesPath) ToYaml() (string, error) {
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

func (m *unMarshalcapabilitiesPath) FromYaml(value string) error {
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

func (m *marshalcapabilitiesPath) ToJson() (string, error) {
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

func (m *unMarshalcapabilitiesPath) FromJson(value string) error {
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

func (obj *capabilitiesPath) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *capabilitiesPath) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *capabilitiesPath) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *capabilitiesPath) Clone() (CapabilitiesPath, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCapabilitiesPath()
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

// CapabilitiesPath is a path indicating one specific object and all its children OR leaf attribute OR  a specific value of leaf 'choice' field of a leaf object of the model
// hierarchy that is supported or not supported by a specific port_type , depending on whether the request_type is included_paths or excluded_path.
type CapabilitiesPath interface {
	Validation
	// msg marshals CapabilitiesPath to protobuf object *otg.CapabilitiesPath
	// and doesn't set defaults
	msg() *otg.CapabilitiesPath
	// setMsg unmarshals CapabilitiesPath from protobuf object *otg.CapabilitiesPath
	// and doesn't set defaults
	setMsg(*otg.CapabilitiesPath) CapabilitiesPath
	// provides marshal interface
	Marshal() marshalCapabilitiesPath
	// provides unmarshal interface
	Unmarshal() unMarshalCapabilitiesPath
	// validate validates CapabilitiesPath
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CapabilitiesPath, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Path returns string, set in CapabilitiesPath.
	Path() string
	// SetPath assigns string provided by user to CapabilitiesPath
	SetPath(value string) CapabilitiesPath
	// HasPath checks if Path has been set in CapabilitiesPath
	HasPath() bool
}

// A path indicating one specific object and all its children, one specific attribute at a leaf of the model hierachy or the specific value of 'choice' value at the leaf of the model
// hierarchy that is supported or not supported by a specific port_type , depending on whether the request_type is included_paths or excluded_path.
// The path string should always be of the generic format `/<node1>/<node2>/..../<node-n>[/<attribute[/choice/<enum value of choice attribute >]]` where node-x is a valid child object of node-(x-1) and
// optional `attribute` is a child attribute of node-n .
//
// Path returns a string
func (obj *capabilitiesPath) Path() string {

	return *obj.obj.Path

}

// A path indicating one specific object and all its children, one specific attribute at a leaf of the model hierachy or the specific value of 'choice' value at the leaf of the model
// hierarchy that is supported or not supported by a specific port_type , depending on whether the request_type is included_paths or excluded_path.
// The path string should always be of the generic format `/<node1>/<node2>/..../<node-n>[/<attribute[/choice/<enum value of choice attribute >]]` where node-x is a valid child object of node-(x-1) and
// optional `attribute` is a child attribute of node-n .
//
// Path returns a string
func (obj *capabilitiesPath) HasPath() bool {
	return obj.obj.Path != nil
}

// A path indicating one specific object and all its children, one specific attribute at a leaf of the model hierachy or the specific value of 'choice' value at the leaf of the model
// hierarchy that is supported or not supported by a specific port_type , depending on whether the request_type is included_paths or excluded_path.
// The path string should always be of the generic format `/<node1>/<node2>/..../<node-n>[/<attribute[/choice/<enum value of choice attribute >]]` where node-x is a valid child object of node-(x-1) and
// optional `attribute` is a child attribute of node-n .
//
// SetPath sets the string value in the CapabilitiesPath object
func (obj *capabilitiesPath) SetPath(value string) CapabilitiesPath {

	obj.obj.Path = &value
	return obj
}

func (obj *capabilitiesPath) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *capabilitiesPath) setDefault() {

}
