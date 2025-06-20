package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetCaptureResponse *****
type getCaptureResponse struct {
	validation
	obj          *otg.GetCaptureResponse
	marshaller   marshalGetCaptureResponse
	unMarshaller unMarshalGetCaptureResponse
}

func NewGetCaptureResponse() GetCaptureResponse {
	obj := getCaptureResponse{obj: &otg.GetCaptureResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getCaptureResponse) msg() *otg.GetCaptureResponse {
	return obj.obj
}

func (obj *getCaptureResponse) setMsg(msg *otg.GetCaptureResponse) GetCaptureResponse {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetCaptureResponse struct {
	obj *getCaptureResponse
}

type marshalGetCaptureResponse interface {
	// ToProto marshals GetCaptureResponse to protobuf object *otg.GetCaptureResponse
	ToProto() (*otg.GetCaptureResponse, error)
	// ToPbText marshals GetCaptureResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetCaptureResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetCaptureResponse to JSON text
	ToJson() (string, error)
}

type unMarshalgetCaptureResponse struct {
	obj *getCaptureResponse
}

type unMarshalGetCaptureResponse interface {
	// FromProto unmarshals GetCaptureResponse from protobuf object *otg.GetCaptureResponse
	FromProto(msg *otg.GetCaptureResponse) (GetCaptureResponse, error)
	// FromPbText unmarshals GetCaptureResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetCaptureResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetCaptureResponse from JSON text
	FromJson(value string) error
}

func (obj *getCaptureResponse) Marshal() marshalGetCaptureResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetCaptureResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *getCaptureResponse) Unmarshal() unMarshalGetCaptureResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetCaptureResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetCaptureResponse) ToProto() (*otg.GetCaptureResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetCaptureResponse) FromProto(msg *otg.GetCaptureResponse) (GetCaptureResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetCaptureResponse) ToPbText() (string, error) {
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

func (m *unMarshalgetCaptureResponse) FromPbText(value string) error {
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

func (m *marshalgetCaptureResponse) ToYaml() (string, error) {
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

func (m *unMarshalgetCaptureResponse) FromYaml(value string) error {
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

func (m *marshalgetCaptureResponse) ToJson() (string, error) {
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

func (m *unMarshalgetCaptureResponse) FromJson(value string) error {
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

func (obj *getCaptureResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getCaptureResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getCaptureResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getCaptureResponse) Clone() (GetCaptureResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetCaptureResponse()
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

// GetCaptureResponse is description is TBD
type GetCaptureResponse interface {
	Validation
	// msg marshals GetCaptureResponse to protobuf object *otg.GetCaptureResponse
	// and doesn't set defaults
	msg() *otg.GetCaptureResponse
	// setMsg unmarshals GetCaptureResponse from protobuf object *otg.GetCaptureResponse
	// and doesn't set defaults
	setMsg(*otg.GetCaptureResponse) GetCaptureResponse
	// provides marshal interface
	Marshal() marshalGetCaptureResponse
	// provides unmarshal interface
	Unmarshal() unMarshalGetCaptureResponse
	// validate validates GetCaptureResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetCaptureResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ResponseBytes returns []byte, set in GetCaptureResponse.
	ResponseBytes() []byte
	// SetResponseBytes assigns []byte provided by user to GetCaptureResponse
	SetResponseBytes(value []byte) GetCaptureResponse
	// HasResponseBytes checks if ResponseBytes has been set in GetCaptureResponse
	HasResponseBytes() bool
}

// description is TBD
// ResponseBytes returns a []byte
func (obj *getCaptureResponse) ResponseBytes() []byte {

	return obj.obj.ResponseBytes
}

// description is TBD
// ResponseBytes returns a []byte
func (obj *getCaptureResponse) HasResponseBytes() bool {
	return obj.obj.ResponseBytes != nil
}

// description is TBD
// SetResponseBytes sets the []byte value in the GetCaptureResponse object
func (obj *getCaptureResponse) SetResponseBytes(value []byte) GetCaptureResponse {

	obj.obj.ResponseBytes = value
	return obj
}

func (obj *getCaptureResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *getCaptureResponse) setDefault() {

}
