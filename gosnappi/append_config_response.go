package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** AppendConfigResponse *****
type appendConfigResponse struct {
	validation
	obj           *otg.AppendConfigResponse
	marshaller    marshalAppendConfigResponse
	unMarshaller  unMarshalAppendConfigResponse
	warningHolder Warning
}

func NewAppendConfigResponse() AppendConfigResponse {
	obj := appendConfigResponse{obj: &otg.AppendConfigResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *appendConfigResponse) msg() *otg.AppendConfigResponse {
	return obj.obj
}

func (obj *appendConfigResponse) setMsg(msg *otg.AppendConfigResponse) AppendConfigResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalappendConfigResponse struct {
	obj *appendConfigResponse
}

type marshalAppendConfigResponse interface {
	// ToProto marshals AppendConfigResponse to protobuf object *otg.AppendConfigResponse
	ToProto() (*otg.AppendConfigResponse, error)
	// ToPbText marshals AppendConfigResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals AppendConfigResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals AppendConfigResponse to JSON text
	ToJson() (string, error)
}

type unMarshalappendConfigResponse struct {
	obj *appendConfigResponse
}

type unMarshalAppendConfigResponse interface {
	// FromProto unmarshals AppendConfigResponse from protobuf object *otg.AppendConfigResponse
	FromProto(msg *otg.AppendConfigResponse) (AppendConfigResponse, error)
	// FromPbText unmarshals AppendConfigResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals AppendConfigResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals AppendConfigResponse from JSON text
	FromJson(value string) error
}

func (obj *appendConfigResponse) Marshal() marshalAppendConfigResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalappendConfigResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *appendConfigResponse) Unmarshal() unMarshalAppendConfigResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalappendConfigResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalappendConfigResponse) ToProto() (*otg.AppendConfigResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalappendConfigResponse) FromProto(msg *otg.AppendConfigResponse) (AppendConfigResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalappendConfigResponse) ToPbText() (string, error) {
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

func (m *unMarshalappendConfigResponse) FromPbText(value string) error {
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

func (m *marshalappendConfigResponse) ToYaml() (string, error) {
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

func (m *unMarshalappendConfigResponse) FromYaml(value string) error {
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

func (m *marshalappendConfigResponse) ToJson() (string, error) {
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

func (m *unMarshalappendConfigResponse) FromJson(value string) error {
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

func (obj *appendConfigResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *appendConfigResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *appendConfigResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *appendConfigResponse) Clone() (AppendConfigResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewAppendConfigResponse()
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

func (obj *appendConfigResponse) setNil() {
	obj.warningHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// AppendConfigResponse is description is TBD
type AppendConfigResponse interface {
	Validation
	// msg marshals AppendConfigResponse to protobuf object *otg.AppendConfigResponse
	// and doesn't set defaults
	msg() *otg.AppendConfigResponse
	// setMsg unmarshals AppendConfigResponse from protobuf object *otg.AppendConfigResponse
	// and doesn't set defaults
	setMsg(*otg.AppendConfigResponse) AppendConfigResponse
	// provides marshal interface
	Marshal() marshalAppendConfigResponse
	// provides unmarshal interface
	Unmarshal() unMarshalAppendConfigResponse
	// validate validates AppendConfigResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (AppendConfigResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Warning returns Warning, set in AppendConfigResponse.
	// Warning is a list of warnings that have occurred while executing the request.
	Warning() Warning
	// SetWarning assigns Warning provided by user to AppendConfigResponse.
	// Warning is a list of warnings that have occurred while executing the request.
	SetWarning(value Warning) AppendConfigResponse
	// HasWarning checks if Warning has been set in AppendConfigResponse
	HasWarning() bool
	setNil()
}

// description is TBD
// Warning returns a Warning
func (obj *appendConfigResponse) Warning() Warning {
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
func (obj *appendConfigResponse) HasWarning() bool {
	return obj.obj.Warning != nil
}

// description is TBD
// SetWarning sets the Warning value in the AppendConfigResponse object
func (obj *appendConfigResponse) SetWarning(value Warning) AppendConfigResponse {

	obj.warningHolder = nil
	obj.obj.Warning = value.msg()

	return obj
}

func (obj *appendConfigResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Warning != nil {

		obj.Warning().validateObj(vObj, set_default)
	}

}

func (obj *appendConfigResponse) setDefault() {

}
