package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetCapabilitiesResponse *****
type getCapabilitiesResponse struct {
	validation
	obj                        *otg.GetCapabilitiesResponse
	marshaller                 marshalGetCapabilitiesResponse
	unMarshaller               unMarshalGetCapabilitiesResponse
	capabilitiesResponseHolder CapabilitiesResponse
}

func NewGetCapabilitiesResponse() GetCapabilitiesResponse {
	obj := getCapabilitiesResponse{obj: &otg.GetCapabilitiesResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getCapabilitiesResponse) msg() *otg.GetCapabilitiesResponse {
	return obj.obj
}

func (obj *getCapabilitiesResponse) setMsg(msg *otg.GetCapabilitiesResponse) GetCapabilitiesResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetCapabilitiesResponse struct {
	obj *getCapabilitiesResponse
}

type marshalGetCapabilitiesResponse interface {
	// ToProto marshals GetCapabilitiesResponse to protobuf object *otg.GetCapabilitiesResponse
	ToProto() (*otg.GetCapabilitiesResponse, error)
	// ToPbText marshals GetCapabilitiesResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetCapabilitiesResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetCapabilitiesResponse to JSON text
	ToJson() (string, error)
}

type unMarshalgetCapabilitiesResponse struct {
	obj *getCapabilitiesResponse
}

type unMarshalGetCapabilitiesResponse interface {
	// FromProto unmarshals GetCapabilitiesResponse from protobuf object *otg.GetCapabilitiesResponse
	FromProto(msg *otg.GetCapabilitiesResponse) (GetCapabilitiesResponse, error)
	// FromPbText unmarshals GetCapabilitiesResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetCapabilitiesResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetCapabilitiesResponse from JSON text
	FromJson(value string) error
}

func (obj *getCapabilitiesResponse) Marshal() marshalGetCapabilitiesResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetCapabilitiesResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *getCapabilitiesResponse) Unmarshal() unMarshalGetCapabilitiesResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetCapabilitiesResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetCapabilitiesResponse) ToProto() (*otg.GetCapabilitiesResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetCapabilitiesResponse) FromProto(msg *otg.GetCapabilitiesResponse) (GetCapabilitiesResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetCapabilitiesResponse) ToPbText() (string, error) {
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

func (m *unMarshalgetCapabilitiesResponse) FromPbText(value string) error {
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

func (m *marshalgetCapabilitiesResponse) ToYaml() (string, error) {
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

func (m *unMarshalgetCapabilitiesResponse) FromYaml(value string) error {
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

func (m *marshalgetCapabilitiesResponse) ToJson() (string, error) {
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

func (m *unMarshalgetCapabilitiesResponse) FromJson(value string) error {
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

func (obj *getCapabilitiesResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getCapabilitiesResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getCapabilitiesResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getCapabilitiesResponse) Clone() (GetCapabilitiesResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetCapabilitiesResponse()
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

func (obj *getCapabilitiesResponse) setNil() {
	obj.capabilitiesResponseHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetCapabilitiesResponse is description is TBD
type GetCapabilitiesResponse interface {
	Validation
	// msg marshals GetCapabilitiesResponse to protobuf object *otg.GetCapabilitiesResponse
	// and doesn't set defaults
	msg() *otg.GetCapabilitiesResponse
	// setMsg unmarshals GetCapabilitiesResponse from protobuf object *otg.GetCapabilitiesResponse
	// and doesn't set defaults
	setMsg(*otg.GetCapabilitiesResponse) GetCapabilitiesResponse
	// provides marshal interface
	Marshal() marshalGetCapabilitiesResponse
	// provides unmarshal interface
	Unmarshal() unMarshalGetCapabilitiesResponse
	// validate validates GetCapabilitiesResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetCapabilitiesResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CapabilitiesResponse returns CapabilitiesResponse, set in GetCapabilitiesResponse.
	// CapabilitiesResponse is container for Capabiities Response from the traffic generator implementation.
	CapabilitiesResponse() CapabilitiesResponse
	// SetCapabilitiesResponse assigns CapabilitiesResponse provided by user to GetCapabilitiesResponse.
	// CapabilitiesResponse is container for Capabiities Response from the traffic generator implementation.
	SetCapabilitiesResponse(value CapabilitiesResponse) GetCapabilitiesResponse
	// HasCapabilitiesResponse checks if CapabilitiesResponse has been set in GetCapabilitiesResponse
	HasCapabilitiesResponse() bool
	setNil()
}

// description is TBD
// CapabilitiesResponse returns a CapabilitiesResponse
func (obj *getCapabilitiesResponse) CapabilitiesResponse() CapabilitiesResponse {
	if obj.obj.CapabilitiesResponse == nil {
		obj.obj.CapabilitiesResponse = NewCapabilitiesResponse().msg()
	}
	if obj.capabilitiesResponseHolder == nil {
		obj.capabilitiesResponseHolder = &capabilitiesResponse{obj: obj.obj.CapabilitiesResponse}
	}
	return obj.capabilitiesResponseHolder
}

// description is TBD
// CapabilitiesResponse returns a CapabilitiesResponse
func (obj *getCapabilitiesResponse) HasCapabilitiesResponse() bool {
	return obj.obj.CapabilitiesResponse != nil
}

// description is TBD
// SetCapabilitiesResponse sets the CapabilitiesResponse value in the GetCapabilitiesResponse object
func (obj *getCapabilitiesResponse) SetCapabilitiesResponse(value CapabilitiesResponse) GetCapabilitiesResponse {

	obj.capabilitiesResponseHolder = nil
	obj.obj.CapabilitiesResponse = value.msg()

	return obj
}

func (obj *getCapabilitiesResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CapabilitiesResponse != nil {

		obj.CapabilitiesResponse().validateObj(vObj, set_default)
	}

}

func (obj *getCapabilitiesResponse) setDefault() {

}
