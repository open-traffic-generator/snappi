package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetConfigResponse *****
type getConfigResponse struct {
	validation
	obj          *otg.GetConfigResponse
	marshaller   marshalGetConfigResponse
	unMarshaller unMarshalGetConfigResponse
	configHolder Config
}

func NewGetConfigResponse() GetConfigResponse {
	obj := getConfigResponse{obj: &otg.GetConfigResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getConfigResponse) msg() *otg.GetConfigResponse {
	return obj.obj
}

func (obj *getConfigResponse) setMsg(msg *otg.GetConfigResponse) GetConfigResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetConfigResponse struct {
	obj *getConfigResponse
}

type marshalGetConfigResponse interface {
	// ToProto marshals GetConfigResponse to protobuf object *otg.GetConfigResponse
	ToProto() (*otg.GetConfigResponse, error)
	// ToPbText marshals GetConfigResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetConfigResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetConfigResponse to JSON text
	ToJson() (string, error)
}

type unMarshalgetConfigResponse struct {
	obj *getConfigResponse
}

type unMarshalGetConfigResponse interface {
	// FromProto unmarshals GetConfigResponse from protobuf object *otg.GetConfigResponse
	FromProto(msg *otg.GetConfigResponse) (GetConfigResponse, error)
	// FromPbText unmarshals GetConfigResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetConfigResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetConfigResponse from JSON text
	FromJson(value string) error
}

func (obj *getConfigResponse) Marshal() marshalGetConfigResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetConfigResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *getConfigResponse) Unmarshal() unMarshalGetConfigResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetConfigResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetConfigResponse) ToProto() (*otg.GetConfigResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetConfigResponse) FromProto(msg *otg.GetConfigResponse) (GetConfigResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetConfigResponse) ToPbText() (string, error) {
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

func (m *unMarshalgetConfigResponse) FromPbText(value string) error {
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

func (m *marshalgetConfigResponse) ToYaml() (string, error) {
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

func (m *unMarshalgetConfigResponse) FromYaml(value string) error {
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

func (m *marshalgetConfigResponse) ToJson() (string, error) {
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

func (m *unMarshalgetConfigResponse) FromJson(value string) error {
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

func (obj *getConfigResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getConfigResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getConfigResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getConfigResponse) Clone() (GetConfigResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetConfigResponse()
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

func (obj *getConfigResponse) setNil() {
	obj.configHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetConfigResponse is description is TBD
type GetConfigResponse interface {
	Validation
	// msg marshals GetConfigResponse to protobuf object *otg.GetConfigResponse
	// and doesn't set defaults
	msg() *otg.GetConfigResponse
	// setMsg unmarshals GetConfigResponse from protobuf object *otg.GetConfigResponse
	// and doesn't set defaults
	setMsg(*otg.GetConfigResponse) GetConfigResponse
	// provides marshal interface
	Marshal() marshalGetConfigResponse
	// provides unmarshal interface
	Unmarshal() unMarshalGetConfigResponse
	// validate validates GetConfigResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetConfigResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Config returns Config, set in GetConfigResponse.
	// Config is a container for all models that are part of the configuration.
	Config() Config
	// SetConfig assigns Config provided by user to GetConfigResponse.
	// Config is a container for all models that are part of the configuration.
	SetConfig(value Config) GetConfigResponse
	// HasConfig checks if Config has been set in GetConfigResponse
	HasConfig() bool
	setNil()
}

// description is TBD
// Config returns a Config
func (obj *getConfigResponse) Config() Config {
	if obj.obj.Config == nil {
		obj.obj.Config = NewConfig().msg()
	}
	if obj.configHolder == nil {
		obj.configHolder = &config{obj: obj.obj.Config}
	}
	return obj.configHolder
}

// description is TBD
// Config returns a Config
func (obj *getConfigResponse) HasConfig() bool {
	return obj.obj.Config != nil
}

// description is TBD
// SetConfig sets the Config value in the GetConfigResponse object
func (obj *getConfigResponse) SetConfig(value Config) GetConfigResponse {

	obj.configHolder = nil
	obj.obj.Config = value.msg()

	return obj
}

func (obj *getConfigResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Config != nil {

		obj.Config().validateObj(vObj, set_default)
	}

}

func (obj *getConfigResponse) setDefault() {

}
