package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SetControlStateResponse *****
type setControlStateResponse struct {
	validation
	obj           *otg.SetControlStateResponse
	marshaller    marshalSetControlStateResponse
	unMarshaller  unMarshalSetControlStateResponse
	warningHolder Warning
}

func NewSetControlStateResponse() SetControlStateResponse {
	obj := setControlStateResponse{obj: &otg.SetControlStateResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *setControlStateResponse) msg() *otg.SetControlStateResponse {
	return obj.obj
}

func (obj *setControlStateResponse) setMsg(msg *otg.SetControlStateResponse) SetControlStateResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsetControlStateResponse struct {
	obj *setControlStateResponse
}

type marshalSetControlStateResponse interface {
	// ToProto marshals SetControlStateResponse to protobuf object *otg.SetControlStateResponse
	ToProto() (*otg.SetControlStateResponse, error)
	// ToPbText marshals SetControlStateResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SetControlStateResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals SetControlStateResponse to JSON text
	ToJson() (string, error)
}

type unMarshalsetControlStateResponse struct {
	obj *setControlStateResponse
}

type unMarshalSetControlStateResponse interface {
	// FromProto unmarshals SetControlStateResponse from protobuf object *otg.SetControlStateResponse
	FromProto(msg *otg.SetControlStateResponse) (SetControlStateResponse, error)
	// FromPbText unmarshals SetControlStateResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SetControlStateResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SetControlStateResponse from JSON text
	FromJson(value string) error
}

func (obj *setControlStateResponse) Marshal() marshalSetControlStateResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsetControlStateResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *setControlStateResponse) Unmarshal() unMarshalSetControlStateResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsetControlStateResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsetControlStateResponse) ToProto() (*otg.SetControlStateResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsetControlStateResponse) FromProto(msg *otg.SetControlStateResponse) (SetControlStateResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsetControlStateResponse) ToPbText() (string, error) {
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

func (m *unMarshalsetControlStateResponse) FromPbText(value string) error {
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

func (m *marshalsetControlStateResponse) ToYaml() (string, error) {
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

func (m *unMarshalsetControlStateResponse) FromYaml(value string) error {
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

func (m *marshalsetControlStateResponse) ToJson() (string, error) {
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

func (m *unMarshalsetControlStateResponse) FromJson(value string) error {
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

func (obj *setControlStateResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *setControlStateResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *setControlStateResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *setControlStateResponse) Clone() (SetControlStateResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSetControlStateResponse()
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

func (obj *setControlStateResponse) setNil() {
	obj.warningHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SetControlStateResponse is description is TBD
type SetControlStateResponse interface {
	Validation
	// msg marshals SetControlStateResponse to protobuf object *otg.SetControlStateResponse
	// and doesn't set defaults
	msg() *otg.SetControlStateResponse
	// setMsg unmarshals SetControlStateResponse from protobuf object *otg.SetControlStateResponse
	// and doesn't set defaults
	setMsg(*otg.SetControlStateResponse) SetControlStateResponse
	// provides marshal interface
	Marshal() marshalSetControlStateResponse
	// provides unmarshal interface
	Unmarshal() unMarshalSetControlStateResponse
	// validate validates SetControlStateResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SetControlStateResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Warning returns Warning, set in SetControlStateResponse.
	// Warning is a list of warnings that have occurred while executing the request.
	Warning() Warning
	// SetWarning assigns Warning provided by user to SetControlStateResponse.
	// Warning is a list of warnings that have occurred while executing the request.
	SetWarning(value Warning) SetControlStateResponse
	// HasWarning checks if Warning has been set in SetControlStateResponse
	HasWarning() bool
	setNil()
}

// description is TBD
// Warning returns a Warning
func (obj *setControlStateResponse) Warning() Warning {
	if obj.obj.Warning == nil {
		obj.obj.Warning = NewWarning().msg()
	}
	if obj.warningHolder == nil {
		obj.warningHolder = &warning{obj: obj.obj.Warning}
	}
	return obj.warningHolder
}

// description is TBD
// Warning returns a Warning
func (obj *setControlStateResponse) HasWarning() bool {
	return obj.obj.Warning != nil
}

// description is TBD
// SetWarning sets the Warning value in the SetControlStateResponse object
func (obj *setControlStateResponse) SetWarning(value Warning) SetControlStateResponse {

	obj.warningHolder = nil
	obj.obj.Warning = value.msg()

	return obj
}

func (obj *setControlStateResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Warning != nil {

		obj.Warning().validateObj(vObj, set_default)
	}

}

func (obj *setControlStateResponse) setDefault() {

}
