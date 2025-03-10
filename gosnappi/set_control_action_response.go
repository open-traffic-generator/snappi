package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SetControlActionResponse *****
type setControlActionResponse struct {
	validation
	obj                         *otg.SetControlActionResponse
	marshaller                  marshalSetControlActionResponse
	unMarshaller                unMarshalSetControlActionResponse
	controlActionResponseHolder ControlActionResponse
}

func NewSetControlActionResponse() SetControlActionResponse {
	obj := setControlActionResponse{obj: &otg.SetControlActionResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *setControlActionResponse) msg() *otg.SetControlActionResponse {
	return obj.obj
}

func (obj *setControlActionResponse) setMsg(msg *otg.SetControlActionResponse) SetControlActionResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsetControlActionResponse struct {
	obj *setControlActionResponse
}

type marshalSetControlActionResponse interface {
	// ToProto marshals SetControlActionResponse to protobuf object *otg.SetControlActionResponse
	ToProto() (*otg.SetControlActionResponse, error)
	// ToPbText marshals SetControlActionResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SetControlActionResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals SetControlActionResponse to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals SetControlActionResponse to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalsetControlActionResponse struct {
	obj *setControlActionResponse
}

type unMarshalSetControlActionResponse interface {
	// FromProto unmarshals SetControlActionResponse from protobuf object *otg.SetControlActionResponse
	FromProto(msg *otg.SetControlActionResponse) (SetControlActionResponse, error)
	// FromPbText unmarshals SetControlActionResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SetControlActionResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SetControlActionResponse from JSON text
	FromJson(value string) error
}

func (obj *setControlActionResponse) Marshal() marshalSetControlActionResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsetControlActionResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *setControlActionResponse) Unmarshal() unMarshalSetControlActionResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsetControlActionResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsetControlActionResponse) ToProto() (*otg.SetControlActionResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsetControlActionResponse) FromProto(msg *otg.SetControlActionResponse) (SetControlActionResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsetControlActionResponse) ToPbText() (string, error) {
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

func (m *unMarshalsetControlActionResponse) FromPbText(value string) error {
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

func (m *marshalsetControlActionResponse) ToYaml() (string, error) {
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

func (m *unMarshalsetControlActionResponse) FromYaml(value string) error {
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

func (m *marshalsetControlActionResponse) ToJsonRaw() (string, error) {
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

func (m *marshalsetControlActionResponse) ToJson() (string, error) {
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

func (m *unMarshalsetControlActionResponse) FromJson(value string) error {
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

func (obj *setControlActionResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *setControlActionResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *setControlActionResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *setControlActionResponse) Clone() (SetControlActionResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSetControlActionResponse()
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

func (obj *setControlActionResponse) setNil() {
	obj.controlActionResponseHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SetControlActionResponse is description is TBD
type SetControlActionResponse interface {
	Validation
	// msg marshals SetControlActionResponse to protobuf object *otg.SetControlActionResponse
	// and doesn't set defaults
	msg() *otg.SetControlActionResponse
	// setMsg unmarshals SetControlActionResponse from protobuf object *otg.SetControlActionResponse
	// and doesn't set defaults
	setMsg(*otg.SetControlActionResponse) SetControlActionResponse
	// provides marshal interface
	Marshal() marshalSetControlActionResponse
	// provides unmarshal interface
	Unmarshal() unMarshalSetControlActionResponse
	// validate validates SetControlActionResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SetControlActionResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ControlActionResponse returns ControlActionResponse, set in SetControlActionResponse.
	// ControlActionResponse is response for action triggered against configured resources along with warnings.
	ControlActionResponse() ControlActionResponse
	// SetControlActionResponse assigns ControlActionResponse provided by user to SetControlActionResponse.
	// ControlActionResponse is response for action triggered against configured resources along with warnings.
	SetControlActionResponse(value ControlActionResponse) SetControlActionResponse
	// HasControlActionResponse checks if ControlActionResponse has been set in SetControlActionResponse
	HasControlActionResponse() bool
	setNil()
}

// description is TBD
// ControlActionResponse returns a ControlActionResponse
func (obj *setControlActionResponse) ControlActionResponse() ControlActionResponse {
	if obj.obj.ControlActionResponse == nil {
		obj.obj.ControlActionResponse = NewControlActionResponse().msg()
	}
	if obj.controlActionResponseHolder == nil {
		obj.controlActionResponseHolder = &controlActionResponse{obj: obj.obj.ControlActionResponse}
	}
	return obj.controlActionResponseHolder
}

// description is TBD
// ControlActionResponse returns a ControlActionResponse
func (obj *setControlActionResponse) HasControlActionResponse() bool {
	return obj.obj.ControlActionResponse != nil
}

// description is TBD
// SetControlActionResponse sets the ControlActionResponse value in the SetControlActionResponse object
func (obj *setControlActionResponse) SetControlActionResponse(value ControlActionResponse) SetControlActionResponse {

	obj.controlActionResponseHolder = nil
	obj.obj.ControlActionResponse = value.msg()

	return obj
}

func (obj *setControlActionResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ControlActionResponse != nil {

		obj.ControlActionResponse().validateObj(vObj, set_default)
	}

}

func (obj *setControlActionResponse) setDefault() {

}
