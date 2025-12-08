package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConfigTagValuePair *****
type configTagValuePair struct {
	validation
	obj          *otg.ConfigTagValuePair
	marshaller   marshalConfigTagValuePair
	unMarshaller unMarshalConfigTagValuePair
}

func NewConfigTagValuePair() ConfigTagValuePair {
	obj := configTagValuePair{obj: &otg.ConfigTagValuePair{}}
	obj.setDefault()
	return &obj
}

func (obj *configTagValuePair) msg() *otg.ConfigTagValuePair {
	return obj.obj
}

func (obj *configTagValuePair) setMsg(msg *otg.ConfigTagValuePair) ConfigTagValuePair {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconfigTagValuePair struct {
	obj *configTagValuePair
}

type marshalConfigTagValuePair interface {
	// ToProto marshals ConfigTagValuePair to protobuf object *otg.ConfigTagValuePair
	ToProto() (*otg.ConfigTagValuePair, error)
	// ToPbText marshals ConfigTagValuePair to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConfigTagValuePair to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConfigTagValuePair to JSON text
	ToJson() (string, error)
}

type unMarshalconfigTagValuePair struct {
	obj *configTagValuePair
}

type unMarshalConfigTagValuePair interface {
	// FromProto unmarshals ConfigTagValuePair from protobuf object *otg.ConfigTagValuePair
	FromProto(msg *otg.ConfigTagValuePair) (ConfigTagValuePair, error)
	// FromPbText unmarshals ConfigTagValuePair from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConfigTagValuePair from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConfigTagValuePair from JSON text
	FromJson(value string) error
}

func (obj *configTagValuePair) Marshal() marshalConfigTagValuePair {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconfigTagValuePair{obj: obj}
	}
	return obj.marshaller
}

func (obj *configTagValuePair) Unmarshal() unMarshalConfigTagValuePair {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconfigTagValuePair{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconfigTagValuePair) ToProto() (*otg.ConfigTagValuePair, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconfigTagValuePair) FromProto(msg *otg.ConfigTagValuePair) (ConfigTagValuePair, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconfigTagValuePair) ToPbText() (string, error) {
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

func (m *unMarshalconfigTagValuePair) FromPbText(value string) error {
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

func (m *marshalconfigTagValuePair) ToYaml() (string, error) {
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

func (m *unMarshalconfigTagValuePair) FromYaml(value string) error {
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

func (m *marshalconfigTagValuePair) ToJson() (string, error) {
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

func (m *unMarshalconfigTagValuePair) FromJson(value string) error {
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

func (obj *configTagValuePair) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *configTagValuePair) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *configTagValuePair) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *configTagValuePair) Clone() (ConfigTagValuePair, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConfigTagValuePair()
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

// ConfigTagValuePair is one tag value pair being provided in the metadata as part of one set_config .
type ConfigTagValuePair interface {
	Validation
	// msg marshals ConfigTagValuePair to protobuf object *otg.ConfigTagValuePair
	// and doesn't set defaults
	msg() *otg.ConfigTagValuePair
	// setMsg unmarshals ConfigTagValuePair from protobuf object *otg.ConfigTagValuePair
	// and doesn't set defaults
	setMsg(*otg.ConfigTagValuePair) ConfigTagValuePair
	// provides marshal interface
	Marshal() marshalConfigTagValuePair
	// provides unmarshal interface
	Unmarshal() unMarshalConfigTagValuePair
	// validate validates ConfigTagValuePair
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConfigTagValuePair, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Tag returns string, set in ConfigTagValuePair.
	Tag() string
	// SetTag assigns string provided by user to ConfigTagValuePair
	SetTag(value string) ConfigTagValuePair
	// HasTag checks if Tag has been set in ConfigTagValuePair
	HasTag() bool
	// Value returns string, set in ConfigTagValuePair.
	Value() string
	// SetValue assigns string provided by user to ConfigTagValuePair
	SetValue(value string) ConfigTagValuePair
	// HasValue checks if Value has been set in ConfigTagValuePair
	HasValue() bool
}

// A generic tag field which can be used to distinguish multiple information being supplied via the metadata.
// e.g. "sub-test" or "timestamp"
//
// Tag returns a string
func (obj *configTagValuePair) Tag() string {

	return *obj.obj.Tag

}

// A generic tag field which can be used to distinguish multiple information being supplied via the metadata.
// e.g. "sub-test" or "timestamp"
//
// Tag returns a string
func (obj *configTagValuePair) HasTag() bool {
	return obj.obj.Tag != nil
}

// A generic tag field which can be used to distinguish multiple information being supplied via the metadata.
// e.g. "sub-test" or "timestamp"
//
// SetTag sets the string value in the ConfigTagValuePair object
func (obj *configTagValuePair) SetTag(value string) ConfigTagValuePair {

	obj.obj.Tag = &value
	return obj
}

// The specific value associated with the tag provided.
// e.g. current sub-test name or timestamp when the set_config was done from the test.
//
// Value returns a string
func (obj *configTagValuePair) Value() string {

	return *obj.obj.Value

}

// The specific value associated with the tag provided.
// e.g. current sub-test name or timestamp when the set_config was done from the test.
//
// Value returns a string
func (obj *configTagValuePair) HasValue() bool {
	return obj.obj.Value != nil
}

// The specific value associated with the tag provided.
// e.g. current sub-test name or timestamp when the set_config was done from the test.
//
// SetValue sets the string value in the ConfigTagValuePair object
func (obj *configTagValuePair) SetValue(value string) ConfigTagValuePair {

	obj.obj.Value = &value
	return obj
}

func (obj *configTagValuePair) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *configTagValuePair) setDefault() {

}
