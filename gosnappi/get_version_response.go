package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetVersionResponse *****
type getVersionResponse struct {
	validation
	obj           *otg.GetVersionResponse
	marshaller    marshalGetVersionResponse
	unMarshaller  unMarshalGetVersionResponse
	versionHolder Version
}

func NewGetVersionResponse() GetVersionResponse {
	obj := getVersionResponse{obj: &otg.GetVersionResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getVersionResponse) msg() *otg.GetVersionResponse {
	return obj.obj
}

func (obj *getVersionResponse) setMsg(msg *otg.GetVersionResponse) GetVersionResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetVersionResponse struct {
	obj *getVersionResponse
}

type marshalGetVersionResponse interface {
	// ToProto marshals GetVersionResponse to protobuf object *otg.GetVersionResponse
	ToProto() (*otg.GetVersionResponse, error)
	// ToPbText marshals GetVersionResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetVersionResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetVersionResponse to JSON text
	ToJson() (string, error)
}

type unMarshalgetVersionResponse struct {
	obj *getVersionResponse
}

type unMarshalGetVersionResponse interface {
	// FromProto unmarshals GetVersionResponse from protobuf object *otg.GetVersionResponse
	FromProto(msg *otg.GetVersionResponse) (GetVersionResponse, error)
	// FromPbText unmarshals GetVersionResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetVersionResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetVersionResponse from JSON text
	FromJson(value string) error
}

func (obj *getVersionResponse) Marshal() marshalGetVersionResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetVersionResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *getVersionResponse) Unmarshal() unMarshalGetVersionResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetVersionResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetVersionResponse) ToProto() (*otg.GetVersionResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetVersionResponse) FromProto(msg *otg.GetVersionResponse) (GetVersionResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetVersionResponse) ToPbText() (string, error) {
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

func (m *unMarshalgetVersionResponse) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalgetVersionResponse) ToYaml() (string, error) {
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

func (m *unMarshalgetVersionResponse) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalgetVersionResponse) ToJson() (string, error) {
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

func (m *unMarshalgetVersionResponse) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *getVersionResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getVersionResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getVersionResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getVersionResponse) Clone() (GetVersionResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetVersionResponse()
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

func (obj *getVersionResponse) setNil() {
	obj.versionHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetVersionResponse is description is TBD
type GetVersionResponse interface {
	Validation
	// msg marshals GetVersionResponse to protobuf object *otg.GetVersionResponse
	// and doesn't set defaults
	msg() *otg.GetVersionResponse
	// setMsg unmarshals GetVersionResponse from protobuf object *otg.GetVersionResponse
	// and doesn't set defaults
	setMsg(*otg.GetVersionResponse) GetVersionResponse
	// provides marshal interface
	Marshal() marshalGetVersionResponse
	// provides unmarshal interface
	Unmarshal() unMarshalGetVersionResponse
	// validate validates GetVersionResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetVersionResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Version returns Version, set in GetVersionResponse.
	// Version is version details
	Version() Version
	// SetVersion assigns Version provided by user to GetVersionResponse.
	// Version is version details
	SetVersion(value Version) GetVersionResponse
	// HasVersion checks if Version has been set in GetVersionResponse
	HasVersion() bool
	setNil()
}

// description is TBD
// Version returns a Version
func (obj *getVersionResponse) Version() Version {
	if obj.obj.Version == nil {
		obj.obj.Version = NewVersion().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &version{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a Version
func (obj *getVersionResponse) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the Version value in the GetVersionResponse object
func (obj *getVersionResponse) SetVersion(value Version) GetVersionResponse {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

func (obj *getVersionResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(vObj, set_default)
	}

}

func (obj *getVersionResponse) setDefault() {

}
