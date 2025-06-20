package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateConfigResponse *****
type updateConfigResponse struct {
	validation
	obj           *otg.UpdateConfigResponse
	marshaller    marshalUpdateConfigResponse
	unMarshaller  unMarshalUpdateConfigResponse
	warningHolder Warning
}

func NewUpdateConfigResponse() UpdateConfigResponse {
	obj := updateConfigResponse{obj: &otg.UpdateConfigResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *updateConfigResponse) msg() *otg.UpdateConfigResponse {
	return obj.obj
}

func (obj *updateConfigResponse) setMsg(msg *otg.UpdateConfigResponse) UpdateConfigResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateConfigResponse struct {
	obj *updateConfigResponse
}

type marshalUpdateConfigResponse interface {
	// ToProto marshals UpdateConfigResponse to protobuf object *otg.UpdateConfigResponse
	ToProto() (*otg.UpdateConfigResponse, error)
	// ToPbText marshals UpdateConfigResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateConfigResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateConfigResponse to JSON text
	ToJson() (string, error)
}

type unMarshalupdateConfigResponse struct {
	obj *updateConfigResponse
}

type unMarshalUpdateConfigResponse interface {
	// FromProto unmarshals UpdateConfigResponse from protobuf object *otg.UpdateConfigResponse
	FromProto(msg *otg.UpdateConfigResponse) (UpdateConfigResponse, error)
	// FromPbText unmarshals UpdateConfigResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateConfigResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateConfigResponse from JSON text
	FromJson(value string) error
}

func (obj *updateConfigResponse) Marshal() marshalUpdateConfigResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateConfigResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateConfigResponse) Unmarshal() unMarshalUpdateConfigResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateConfigResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateConfigResponse) ToProto() (*otg.UpdateConfigResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateConfigResponse) FromProto(msg *otg.UpdateConfigResponse) (UpdateConfigResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateConfigResponse) ToPbText() (string, error) {
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

func (m *unMarshalupdateConfigResponse) FromPbText(value string) error {
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

func (m *marshalupdateConfigResponse) ToYaml() (string, error) {
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

func (m *unMarshalupdateConfigResponse) FromYaml(value string) error {
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

func (m *marshalupdateConfigResponse) ToJson() (string, error) {
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

func (m *unMarshalupdateConfigResponse) FromJson(value string) error {
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

func (obj *updateConfigResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateConfigResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateConfigResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateConfigResponse) Clone() (UpdateConfigResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateConfigResponse()
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

func (obj *updateConfigResponse) setNil() {
	obj.warningHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateConfigResponse is description is TBD
type UpdateConfigResponse interface {
	Validation
	// msg marshals UpdateConfigResponse to protobuf object *otg.UpdateConfigResponse
	// and doesn't set defaults
	msg() *otg.UpdateConfigResponse
	// setMsg unmarshals UpdateConfigResponse from protobuf object *otg.UpdateConfigResponse
	// and doesn't set defaults
	setMsg(*otg.UpdateConfigResponse) UpdateConfigResponse
	// provides marshal interface
	Marshal() marshalUpdateConfigResponse
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateConfigResponse
	// validate validates UpdateConfigResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateConfigResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Warning returns Warning, set in UpdateConfigResponse.
	// Warning is a list of warnings that have occurred while executing the request.
	Warning() Warning
	// SetWarning assigns Warning provided by user to UpdateConfigResponse.
	// Warning is a list of warnings that have occurred while executing the request.
	SetWarning(value Warning) UpdateConfigResponse
	// HasWarning checks if Warning has been set in UpdateConfigResponse
	HasWarning() bool
	setNil()
}

// description is TBD
// Warning returns a Warning
func (obj *updateConfigResponse) Warning() Warning {
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
func (obj *updateConfigResponse) HasWarning() bool {
	return obj.obj.Warning != nil
}

// description is TBD
// SetWarning sets the Warning value in the UpdateConfigResponse object
func (obj *updateConfigResponse) SetWarning(value Warning) UpdateConfigResponse {

	obj.warningHolder = nil
	obj.obj.Warning = value.msg()

	return obj
}

func (obj *updateConfigResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Warning != nil {

		obj.Warning().validateObj(vObj, set_default)
	}

}

func (obj *updateConfigResponse) setDefault() {

}
