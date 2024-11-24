package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Version *****
type version struct {
	validation
	obj          *otg.Version
	marshaller   marshalVersion
	unMarshaller unMarshalVersion
}

func NewVersion() Version {
	obj := version{obj: &otg.Version{}}
	obj.setDefault()
	return &obj
}

func (obj *version) msg() *otg.Version {
	return obj.obj
}

func (obj *version) setMsg(msg *otg.Version) Version {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalversion struct {
	obj *version
}

type marshalVersion interface {
	// ToProto marshals Version to protobuf object *otg.Version
	ToProto() (*otg.Version, error)
	// ToPbText marshals Version to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Version to YAML text
	ToYaml() (string, error)
	// ToJson marshals Version to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Version to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalversion struct {
	obj *version
}

type unMarshalVersion interface {
	// FromProto unmarshals Version from protobuf object *otg.Version
	FromProto(msg *otg.Version) (Version, error)
	// FromPbText unmarshals Version from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Version from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Version from JSON text
	FromJson(value string) error
}

func (obj *version) Marshal() marshalVersion {
	if obj.marshaller == nil {
		obj.marshaller = &marshalversion{obj: obj}
	}
	return obj.marshaller
}

func (obj *version) Unmarshal() unMarshalVersion {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalversion{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalversion) ToProto() (*otg.Version, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalversion) FromProto(msg *otg.Version) (Version, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalversion) ToPbText() (string, error) {
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

func (m *unMarshalversion) FromPbText(value string) error {
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

func (m *marshalversion) ToYaml() (string, error) {
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

func (m *unMarshalversion) FromYaml(value string) error {
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

func (m *marshalversion) ToJsonRaw() (string, error) {
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

func (m *marshalversion) ToJson() (string, error) {
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

func (m *unMarshalversion) FromJson(value string) error {
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

func (obj *version) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *version) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *version) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *version) Clone() (Version, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVersion()
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

// Version is version details
type Version interface {
	Validation
	// msg marshals Version to protobuf object *otg.Version
	// and doesn't set defaults
	msg() *otg.Version
	// setMsg unmarshals Version from protobuf object *otg.Version
	// and doesn't set defaults
	setMsg(*otg.Version) Version
	// provides marshal interface
	Marshal() marshalVersion
	// provides unmarshal interface
	Unmarshal() unMarshalVersion
	// validate validates Version
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Version, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ApiSpecVersion returns string, set in Version.
	ApiSpecVersion() string
	// SetApiSpecVersion assigns string provided by user to Version
	SetApiSpecVersion(value string) Version
	// HasApiSpecVersion checks if ApiSpecVersion has been set in Version
	HasApiSpecVersion() bool
	// SdkVersion returns string, set in Version.
	SdkVersion() string
	// SetSdkVersion assigns string provided by user to Version
	SetSdkVersion(value string) Version
	// HasSdkVersion checks if SdkVersion has been set in Version
	HasSdkVersion() bool
	// AppVersion returns string, set in Version.
	AppVersion() string
	// SetAppVersion assigns string provided by user to Version
	SetAppVersion(value string) Version
	// HasAppVersion checks if AppVersion has been set in Version
	HasAppVersion() bool
}

// Version of API specification
// ApiSpecVersion returns a string
func (obj *version) ApiSpecVersion() string {

	return *obj.obj.ApiSpecVersion

}

// Version of API specification
// ApiSpecVersion returns a string
func (obj *version) HasApiSpecVersion() bool {
	return obj.obj.ApiSpecVersion != nil
}

// Version of API specification
// SetApiSpecVersion sets the string value in the Version object
func (obj *version) SetApiSpecVersion(value string) Version {

	obj.obj.ApiSpecVersion = &value
	return obj
}

// Version of SDK generated from API specification
// SdkVersion returns a string
func (obj *version) SdkVersion() string {

	return *obj.obj.SdkVersion

}

// Version of SDK generated from API specification
// SdkVersion returns a string
func (obj *version) HasSdkVersion() bool {
	return obj.obj.SdkVersion != nil
}

// Version of SDK generated from API specification
// SetSdkVersion sets the string value in the Version object
func (obj *version) SetSdkVersion(value string) Version {

	obj.obj.SdkVersion = &value
	return obj
}

// Version of application consuming or serving the API
// AppVersion returns a string
func (obj *version) AppVersion() string {

	return *obj.obj.AppVersion

}

// Version of application consuming or serving the API
// AppVersion returns a string
func (obj *version) HasAppVersion() bool {
	return obj.obj.AppVersion != nil
}

// Version of application consuming or serving the API
// SetAppVersion sets the string value in the Version object
func (obj *version) SetAppVersion(value string) Version {

	obj.obj.AppVersion = &value
	return obj
}

func (obj *version) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *version) setDefault() {
	if obj.obj.ApiSpecVersion == nil {
		obj.SetApiSpecVersion("")
	}
	if obj.obj.SdkVersion == nil {
		obj.SetSdkVersion("")
	}
	if obj.obj.AppVersion == nil {
		obj.SetAppVersion("")
	}

}
