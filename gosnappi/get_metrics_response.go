package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetMetricsResponse *****
type getMetricsResponse struct {
	validation
	obj                   *otg.GetMetricsResponse
	marshaller            marshalGetMetricsResponse
	unMarshaller          unMarshalGetMetricsResponse
	metricsResponseHolder MetricsResponse
}

func NewGetMetricsResponse() GetMetricsResponse {
	obj := getMetricsResponse{obj: &otg.GetMetricsResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getMetricsResponse) msg() *otg.GetMetricsResponse {
	return obj.obj
}

func (obj *getMetricsResponse) setMsg(msg *otg.GetMetricsResponse) GetMetricsResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetMetricsResponse struct {
	obj *getMetricsResponse
}

type marshalGetMetricsResponse interface {
	// ToProto marshals GetMetricsResponse to protobuf object *otg.GetMetricsResponse
	ToProto() (*otg.GetMetricsResponse, error)
	// ToPbText marshals GetMetricsResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetMetricsResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetMetricsResponse to JSON text
	ToJson() (string, error)
}

type unMarshalgetMetricsResponse struct {
	obj *getMetricsResponse
}

type unMarshalGetMetricsResponse interface {
	// FromProto unmarshals GetMetricsResponse from protobuf object *otg.GetMetricsResponse
	FromProto(msg *otg.GetMetricsResponse) (GetMetricsResponse, error)
	// FromPbText unmarshals GetMetricsResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetMetricsResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetMetricsResponse from JSON text
	FromJson(value string) error
}

func (obj *getMetricsResponse) Marshal() marshalGetMetricsResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetMetricsResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *getMetricsResponse) Unmarshal() unMarshalGetMetricsResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetMetricsResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetMetricsResponse) ToProto() (*otg.GetMetricsResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetMetricsResponse) FromProto(msg *otg.GetMetricsResponse) (GetMetricsResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetMetricsResponse) ToPbText() (string, error) {
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

func (m *unMarshalgetMetricsResponse) FromPbText(value string) error {
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

func (m *marshalgetMetricsResponse) ToYaml() (string, error) {
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

func (m *unMarshalgetMetricsResponse) FromYaml(value string) error {
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

func (m *marshalgetMetricsResponse) ToJson() (string, error) {
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

func (m *unMarshalgetMetricsResponse) FromJson(value string) error {
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

func (obj *getMetricsResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getMetricsResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getMetricsResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getMetricsResponse) Clone() (GetMetricsResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetMetricsResponse()
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

func (obj *getMetricsResponse) setNil() {
	obj.metricsResponseHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetMetricsResponse is description is TBD
type GetMetricsResponse interface {
	Validation
	// msg marshals GetMetricsResponse to protobuf object *otg.GetMetricsResponse
	// and doesn't set defaults
	msg() *otg.GetMetricsResponse
	// setMsg unmarshals GetMetricsResponse from protobuf object *otg.GetMetricsResponse
	// and doesn't set defaults
	setMsg(*otg.GetMetricsResponse) GetMetricsResponse
	// provides marshal interface
	Marshal() marshalGetMetricsResponse
	// provides unmarshal interface
	Unmarshal() unMarshalGetMetricsResponse
	// validate validates GetMetricsResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetMetricsResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MetricsResponse returns MetricsResponse, set in GetMetricsResponse.
	// MetricsResponse is response containing chosen traffic generator metrics.
	MetricsResponse() MetricsResponse
	// SetMetricsResponse assigns MetricsResponse provided by user to GetMetricsResponse.
	// MetricsResponse is response containing chosen traffic generator metrics.
	SetMetricsResponse(value MetricsResponse) GetMetricsResponse
	// HasMetricsResponse checks if MetricsResponse has been set in GetMetricsResponse
	HasMetricsResponse() bool
	setNil()
}

// description is TBD
// MetricsResponse returns a MetricsResponse
func (obj *getMetricsResponse) MetricsResponse() MetricsResponse {
	if obj.obj.MetricsResponse == nil {
		obj.obj.MetricsResponse = NewMetricsResponse().msg()
	}
	if obj.metricsResponseHolder == nil {
		obj.metricsResponseHolder = &metricsResponse{obj: obj.obj.MetricsResponse}
	}
	return obj.metricsResponseHolder
}

// description is TBD
// MetricsResponse returns a MetricsResponse
func (obj *getMetricsResponse) HasMetricsResponse() bool {
	return obj.obj.MetricsResponse != nil
}

// description is TBD
// SetMetricsResponse sets the MetricsResponse value in the GetMetricsResponse object
func (obj *getMetricsResponse) SetMetricsResponse(value MetricsResponse) GetMetricsResponse {

	obj.metricsResponseHolder = nil
	obj.obj.MetricsResponse = value.msg()

	return obj
}

func (obj *getMetricsResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MetricsResponse != nil {

		obj.MetricsResponse().validateObj(vObj, set_default)
	}

}

func (obj *getMetricsResponse) setDefault() {

}
