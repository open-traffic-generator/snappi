package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeleteConfigResponse *****
type deleteConfigResponse struct {
	validation
	obj           *otg.DeleteConfigResponse
	marshaller    marshalDeleteConfigResponse
	unMarshaller  unMarshalDeleteConfigResponse
	warningHolder Warning
}

func NewDeleteConfigResponse() DeleteConfigResponse {
	obj := deleteConfigResponse{obj: &otg.DeleteConfigResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *deleteConfigResponse) msg() *otg.DeleteConfigResponse {
	return obj.obj
}

func (obj *deleteConfigResponse) setMsg(msg *otg.DeleteConfigResponse) DeleteConfigResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeleteConfigResponse struct {
	obj *deleteConfigResponse
}

type marshalDeleteConfigResponse interface {
	// ToProto marshals DeleteConfigResponse to protobuf object *otg.DeleteConfigResponse
	ToProto() (*otg.DeleteConfigResponse, error)
	// ToPbText marshals DeleteConfigResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeleteConfigResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeleteConfigResponse to JSON text
	ToJson() (string, error)
}

type unMarshaldeleteConfigResponse struct {
	obj *deleteConfigResponse
}

type unMarshalDeleteConfigResponse interface {
	// FromProto unmarshals DeleteConfigResponse from protobuf object *otg.DeleteConfigResponse
	FromProto(msg *otg.DeleteConfigResponse) (DeleteConfigResponse, error)
	// FromPbText unmarshals DeleteConfigResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeleteConfigResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeleteConfigResponse from JSON text
	FromJson(value string) error
}

func (obj *deleteConfigResponse) Marshal() marshalDeleteConfigResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeleteConfigResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *deleteConfigResponse) Unmarshal() unMarshalDeleteConfigResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeleteConfigResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeleteConfigResponse) ToProto() (*otg.DeleteConfigResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeleteConfigResponse) FromProto(msg *otg.DeleteConfigResponse) (DeleteConfigResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeleteConfigResponse) ToPbText() (string, error) {
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

func (m *unMarshaldeleteConfigResponse) FromPbText(value string) error {
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

func (m *marshaldeleteConfigResponse) ToYaml() (string, error) {
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

func (m *unMarshaldeleteConfigResponse) FromYaml(value string) error {
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

func (m *marshaldeleteConfigResponse) ToJson() (string, error) {
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

func (m *unMarshaldeleteConfigResponse) FromJson(value string) error {
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

func (obj *deleteConfigResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deleteConfigResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deleteConfigResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deleteConfigResponse) Clone() (DeleteConfigResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeleteConfigResponse()
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

func (obj *deleteConfigResponse) setNil() {
	obj.warningHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeleteConfigResponse is description is TBD
type DeleteConfigResponse interface {
	Validation
	// msg marshals DeleteConfigResponse to protobuf object *otg.DeleteConfigResponse
	// and doesn't set defaults
	msg() *otg.DeleteConfigResponse
	// setMsg unmarshals DeleteConfigResponse from protobuf object *otg.DeleteConfigResponse
	// and doesn't set defaults
	setMsg(*otg.DeleteConfigResponse) DeleteConfigResponse
	// provides marshal interface
	Marshal() marshalDeleteConfigResponse
	// provides unmarshal interface
	Unmarshal() unMarshalDeleteConfigResponse
	// validate validates DeleteConfigResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeleteConfigResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Warning returns Warning, set in DeleteConfigResponse.
	// Warning is a list of warnings that have occurred while executing the request.
	Warning() Warning
	// SetWarning assigns Warning provided by user to DeleteConfigResponse.
	// Warning is a list of warnings that have occurred while executing the request.
	SetWarning(value Warning) DeleteConfigResponse
	// HasWarning checks if Warning has been set in DeleteConfigResponse
	HasWarning() bool
	setNil()
}

// description is TBD
// Warning returns a Warning
func (obj *deleteConfigResponse) Warning() Warning {
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
func (obj *deleteConfigResponse) HasWarning() bool {
	return obj.obj.Warning != nil
}

// description is TBD
// SetWarning sets the Warning value in the DeleteConfigResponse object
func (obj *deleteConfigResponse) SetWarning(value Warning) DeleteConfigResponse {

	obj.warningHolder = nil
	obj.obj.Warning = value.msg()

	return obj
}

func (obj *deleteConfigResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Warning != nil {

		obj.Warning().validateObj(vObj, set_default)
	}

}

func (obj *deleteConfigResponse) setDefault() {

}
