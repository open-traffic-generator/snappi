package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ControlActionResponse *****
type controlActionResponse struct {
	validation
	obj            *otg.ControlActionResponse
	marshaller     marshalControlActionResponse
	unMarshaller   unMarshalControlActionResponse
	responseHolder ActionResponse
}

func NewControlActionResponse() ControlActionResponse {
	obj := controlActionResponse{obj: &otg.ControlActionResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *controlActionResponse) msg() *otg.ControlActionResponse {
	return obj.obj
}

func (obj *controlActionResponse) setMsg(msg *otg.ControlActionResponse) ControlActionResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcontrolActionResponse struct {
	obj *controlActionResponse
}

type marshalControlActionResponse interface {
	// ToProto marshals ControlActionResponse to protobuf object *otg.ControlActionResponse
	ToProto() (*otg.ControlActionResponse, error)
	// ToPbText marshals ControlActionResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ControlActionResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals ControlActionResponse to JSON text
	ToJson() (string, error)
}

type unMarshalcontrolActionResponse struct {
	obj *controlActionResponse
}

type unMarshalControlActionResponse interface {
	// FromProto unmarshals ControlActionResponse from protobuf object *otg.ControlActionResponse
	FromProto(msg *otg.ControlActionResponse) (ControlActionResponse, error)
	// FromPbText unmarshals ControlActionResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ControlActionResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ControlActionResponse from JSON text
	FromJson(value string) error
}

func (obj *controlActionResponse) Marshal() marshalControlActionResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcontrolActionResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *controlActionResponse) Unmarshal() unMarshalControlActionResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcontrolActionResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcontrolActionResponse) ToProto() (*otg.ControlActionResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcontrolActionResponse) FromProto(msg *otg.ControlActionResponse) (ControlActionResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcontrolActionResponse) ToPbText() (string, error) {
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

func (m *unMarshalcontrolActionResponse) FromPbText(value string) error {
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

func (m *marshalcontrolActionResponse) ToYaml() (string, error) {
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

func (m *unMarshalcontrolActionResponse) FromYaml(value string) error {
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

func (m *marshalcontrolActionResponse) ToJson() (string, error) {
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

func (m *unMarshalcontrolActionResponse) FromJson(value string) error {
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

func (obj *controlActionResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *controlActionResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *controlActionResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *controlActionResponse) Clone() (ControlActionResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewControlActionResponse()
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

func (obj *controlActionResponse) setNil() {
	obj.responseHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ControlActionResponse is response for action triggered against configured resources along with warnings.
type ControlActionResponse interface {
	Validation
	// msg marshals ControlActionResponse to protobuf object *otg.ControlActionResponse
	// and doesn't set defaults
	msg() *otg.ControlActionResponse
	// setMsg unmarshals ControlActionResponse from protobuf object *otg.ControlActionResponse
	// and doesn't set defaults
	setMsg(*otg.ControlActionResponse) ControlActionResponse
	// provides marshal interface
	Marshal() marshalControlActionResponse
	// provides unmarshal interface
	Unmarshal() unMarshalControlActionResponse
	// validate validates ControlActionResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ControlActionResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Warnings returns []string, set in ControlActionResponse.
	Warnings() []string
	// SetWarnings assigns []string provided by user to ControlActionResponse
	SetWarnings(value []string) ControlActionResponse
	// Response returns ActionResponse, set in ControlActionResponse.
	// ActionResponse is response for action triggered against configured resources.
	Response() ActionResponse
	// SetResponse assigns ActionResponse provided by user to ControlActionResponse.
	// ActionResponse is response for action triggered against configured resources.
	SetResponse(value ActionResponse) ControlActionResponse
	// HasResponse checks if Response has been set in ControlActionResponse
	HasResponse() bool
	setNil()
}

// List of warnings generated while triggering specified action
// Warnings returns a []string
func (obj *controlActionResponse) Warnings() []string {
	if obj.obj.Warnings == nil {
		obj.obj.Warnings = make([]string, 0)
	}
	return obj.obj.Warnings
}

// List of warnings generated while triggering specified action
// SetWarnings sets the []string value in the ControlActionResponse object
func (obj *controlActionResponse) SetWarnings(value []string) ControlActionResponse {

	if obj.obj.Warnings == nil {
		obj.obj.Warnings = make([]string, 0)
	}
	obj.obj.Warnings = value

	return obj
}

// description is TBD
// Response returns a ActionResponse
func (obj *controlActionResponse) Response() ActionResponse {
	if obj.obj.Response == nil {
		obj.obj.Response = NewActionResponse().msg()
	}
	if obj.responseHolder == nil {
		obj.responseHolder = &actionResponse{obj: obj.obj.Response}
	}
	return obj.responseHolder
}

// description is TBD
// Response returns a ActionResponse
func (obj *controlActionResponse) HasResponse() bool {
	return obj.obj.Response != nil
}

// description is TBD
// SetResponse sets the ActionResponse value in the ControlActionResponse object
func (obj *controlActionResponse) SetResponse(value ActionResponse) ControlActionResponse {

	obj.responseHolder = nil
	obj.obj.Response = value.msg()

	return obj
}

func (obj *controlActionResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Response != nil {

		obj.Response().validateObj(vObj, set_default)
	}

}

func (obj *controlActionResponse) setDefault() {

}
