package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetStatesResponse *****
type getStatesResponse struct {
	validation
	obj                  *otg.GetStatesResponse
	marshaller           marshalGetStatesResponse
	unMarshaller         unMarshalGetStatesResponse
	statesResponseHolder StatesResponse
}

func NewGetStatesResponse() GetStatesResponse {
	obj := getStatesResponse{obj: &otg.GetStatesResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getStatesResponse) msg() *otg.GetStatesResponse {
	return obj.obj
}

func (obj *getStatesResponse) setMsg(msg *otg.GetStatesResponse) GetStatesResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetStatesResponse struct {
	obj *getStatesResponse
}

type marshalGetStatesResponse interface {
	// ToProto marshals GetStatesResponse to protobuf object *otg.GetStatesResponse
	ToProto() (*otg.GetStatesResponse, error)
	// ToPbText marshals GetStatesResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetStatesResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetStatesResponse to JSON text
	ToJson() (string, error)
}

type unMarshalgetStatesResponse struct {
	obj *getStatesResponse
}

type unMarshalGetStatesResponse interface {
	// FromProto unmarshals GetStatesResponse from protobuf object *otg.GetStatesResponse
	FromProto(msg *otg.GetStatesResponse) (GetStatesResponse, error)
	// FromPbText unmarshals GetStatesResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetStatesResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetStatesResponse from JSON text
	FromJson(value string) error
}

func (obj *getStatesResponse) Marshal() marshalGetStatesResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetStatesResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *getStatesResponse) Unmarshal() unMarshalGetStatesResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetStatesResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetStatesResponse) ToProto() (*otg.GetStatesResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetStatesResponse) FromProto(msg *otg.GetStatesResponse) (GetStatesResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetStatesResponse) ToPbText() (string, error) {
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

func (m *unMarshalgetStatesResponse) FromPbText(value string) error {
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

func (m *marshalgetStatesResponse) ToYaml() (string, error) {
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

func (m *unMarshalgetStatesResponse) FromYaml(value string) error {
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

func (m *marshalgetStatesResponse) ToJson() (string, error) {
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

func (m *unMarshalgetStatesResponse) FromJson(value string) error {
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

func (obj *getStatesResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getStatesResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getStatesResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getStatesResponse) Clone() (GetStatesResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetStatesResponse()
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

func (obj *getStatesResponse) setNil() {
	obj.statesResponseHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetStatesResponse is description is TBD
type GetStatesResponse interface {
	Validation
	// msg marshals GetStatesResponse to protobuf object *otg.GetStatesResponse
	// and doesn't set defaults
	msg() *otg.GetStatesResponse
	// setMsg unmarshals GetStatesResponse from protobuf object *otg.GetStatesResponse
	// and doesn't set defaults
	setMsg(*otg.GetStatesResponse) GetStatesResponse
	// provides marshal interface
	Marshal() marshalGetStatesResponse
	// provides unmarshal interface
	Unmarshal() unMarshalGetStatesResponse
	// validate validates GetStatesResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetStatesResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StatesResponse returns StatesResponse, set in GetStatesResponse.
	// StatesResponse is response containing chosen traffic generator states
	StatesResponse() StatesResponse
	// SetStatesResponse assigns StatesResponse provided by user to GetStatesResponse.
	// StatesResponse is response containing chosen traffic generator states
	SetStatesResponse(value StatesResponse) GetStatesResponse
	// HasStatesResponse checks if StatesResponse has been set in GetStatesResponse
	HasStatesResponse() bool
	setNil()
}

// description is TBD
// StatesResponse returns a StatesResponse
func (obj *getStatesResponse) StatesResponse() StatesResponse {
	if obj.obj.StatesResponse == nil {
		obj.obj.StatesResponse = NewStatesResponse().msg()
	}
	if obj.statesResponseHolder == nil {
		obj.statesResponseHolder = &statesResponse{obj: obj.obj.StatesResponse}
	}
	return obj.statesResponseHolder
}

// description is TBD
// StatesResponse returns a StatesResponse
func (obj *getStatesResponse) HasStatesResponse() bool {
	return obj.obj.StatesResponse != nil
}

// description is TBD
// SetStatesResponse sets the StatesResponse value in the GetStatesResponse object
func (obj *getStatesResponse) SetStatesResponse(value StatesResponse) GetStatesResponse {

	obj.statesResponseHolder = nil
	obj.obj.StatesResponse = value.msg()

	return obj
}

func (obj *getStatesResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StatesResponse != nil {

		obj.StatesResponse().validateObj(vObj, set_default)
	}

}

func (obj *getStatesResponse) setDefault() {

}
