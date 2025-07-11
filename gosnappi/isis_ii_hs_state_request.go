package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHsStateRequest *****
type isisIIHsStateRequest struct {
	validation
	obj          *otg.IsisIIHsStateRequest
	marshaller   marshalIsisIIHsStateRequest
	unMarshaller unMarshalIsisIIHsStateRequest
}

func NewIsisIIHsStateRequest() IsisIIHsStateRequest {
	obj := isisIIHsStateRequest{obj: &otg.IsisIIHsStateRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHsStateRequest) msg() *otg.IsisIIHsStateRequest {
	return obj.obj
}

func (obj *isisIIHsStateRequest) setMsg(msg *otg.IsisIIHsStateRequest) IsisIIHsStateRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHsStateRequest struct {
	obj *isisIIHsStateRequest
}

type marshalIsisIIHsStateRequest interface {
	// ToProto marshals IsisIIHsStateRequest to protobuf object *otg.IsisIIHsStateRequest
	ToProto() (*otg.IsisIIHsStateRequest, error)
	// ToPbText marshals IsisIIHsStateRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHsStateRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHsStateRequest to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHsStateRequest struct {
	obj *isisIIHsStateRequest
}

type unMarshalIsisIIHsStateRequest interface {
	// FromProto unmarshals IsisIIHsStateRequest from protobuf object *otg.IsisIIHsStateRequest
	FromProto(msg *otg.IsisIIHsStateRequest) (IsisIIHsStateRequest, error)
	// FromPbText unmarshals IsisIIHsStateRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHsStateRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHsStateRequest from JSON text
	FromJson(value string) error
}

func (obj *isisIIHsStateRequest) Marshal() marshalIsisIIHsStateRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHsStateRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHsStateRequest) Unmarshal() unMarshalIsisIIHsStateRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHsStateRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHsStateRequest) ToProto() (*otg.IsisIIHsStateRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHsStateRequest) FromProto(msg *otg.IsisIIHsStateRequest) (IsisIIHsStateRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHsStateRequest) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHsStateRequest) FromPbText(value string) error {
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

func (m *marshalisisIIHsStateRequest) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHsStateRequest) FromYaml(value string) error {
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

func (m *marshalisisIIHsStateRequest) ToJson() (string, error) {
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

func (m *unMarshalisisIIHsStateRequest) FromJson(value string) error {
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

func (obj *isisIIHsStateRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHsStateRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHsStateRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHsStateRequest) Clone() (IsisIIHsStateRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHsStateRequest()
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

// IsisIIHsStateRequest is the request to retrieve ISIS IIH information exchanged by the ISIS routers.
type IsisIIHsStateRequest interface {
	Validation
	// msg marshals IsisIIHsStateRequest to protobuf object *otg.IsisIIHsStateRequest
	// and doesn't set defaults
	msg() *otg.IsisIIHsStateRequest
	// setMsg unmarshals IsisIIHsStateRequest from protobuf object *otg.IsisIIHsStateRequest
	// and doesn't set defaults
	setMsg(*otg.IsisIIHsStateRequest) IsisIIHsStateRequest
	// provides marshal interface
	Marshal() marshalIsisIIHsStateRequest
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHsStateRequest
	// validate validates IsisIIHsStateRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHsStateRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IsisRouterNames returns []string, set in IsisIIHsStateRequest.
	IsisRouterNames() []string
	// SetIsisRouterNames assigns []string provided by user to IsisIIHsStateRequest
	SetIsisRouterNames(value []string) IsisIIHsStateRequest
}

// The names of ISIS routers for which learned information is requested. An empty list will return results of IIH States for all ISIS routers.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// IsisRouterNames returns a []string
func (obj *isisIIHsStateRequest) IsisRouterNames() []string {
	if obj.obj.IsisRouterNames == nil {
		obj.obj.IsisRouterNames = make([]string, 0)
	}
	return obj.obj.IsisRouterNames
}

// The names of ISIS routers for which learned information is requested. An empty list will return results of IIH States for all ISIS routers.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// SetIsisRouterNames sets the []string value in the IsisIIHsStateRequest object
func (obj *isisIIHsStateRequest) SetIsisRouterNames(value []string) IsisIIHsStateRequest {

	if obj.obj.IsisRouterNames == nil {
		obj.obj.IsisRouterNames = make([]string, 0)
	}
	obj.obj.IsisRouterNames = value

	return obj
}

func (obj *isisIIHsStateRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisIIHsStateRequest) setDefault() {

}
