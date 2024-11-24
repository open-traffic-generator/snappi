package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspsStateRequest *****
type isisLspsStateRequest struct {
	validation
	obj          *otg.IsisLspsStateRequest
	marshaller   marshalIsisLspsStateRequest
	unMarshaller unMarshalIsisLspsStateRequest
}

func NewIsisLspsStateRequest() IsisLspsStateRequest {
	obj := isisLspsStateRequest{obj: &otg.IsisLspsStateRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspsStateRequest) msg() *otg.IsisLspsStateRequest {
	return obj.obj
}

func (obj *isisLspsStateRequest) setMsg(msg *otg.IsisLspsStateRequest) IsisLspsStateRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspsStateRequest struct {
	obj *isisLspsStateRequest
}

type marshalIsisLspsStateRequest interface {
	// ToProto marshals IsisLspsStateRequest to protobuf object *otg.IsisLspsStateRequest
	ToProto() (*otg.IsisLspsStateRequest, error)
	// ToPbText marshals IsisLspsStateRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspsStateRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspsStateRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspsStateRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspsStateRequest struct {
	obj *isisLspsStateRequest
}

type unMarshalIsisLspsStateRequest interface {
	// FromProto unmarshals IsisLspsStateRequest from protobuf object *otg.IsisLspsStateRequest
	FromProto(msg *otg.IsisLspsStateRequest) (IsisLspsStateRequest, error)
	// FromPbText unmarshals IsisLspsStateRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspsStateRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspsStateRequest from JSON text
	FromJson(value string) error
}

func (obj *isisLspsStateRequest) Marshal() marshalIsisLspsStateRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspsStateRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspsStateRequest) Unmarshal() unMarshalIsisLspsStateRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspsStateRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspsStateRequest) ToProto() (*otg.IsisLspsStateRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspsStateRequest) FromProto(msg *otg.IsisLspsStateRequest) (IsisLspsStateRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspsStateRequest) ToPbText() (string, error) {
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

func (m *unMarshalisisLspsStateRequest) FromPbText(value string) error {
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

func (m *marshalisisLspsStateRequest) ToYaml() (string, error) {
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

func (m *unMarshalisisLspsStateRequest) FromYaml(value string) error {
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

func (m *marshalisisLspsStateRequest) ToJsonRaw() (string, error) {
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

func (m *marshalisisLspsStateRequest) ToJson() (string, error) {
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

func (m *unMarshalisisLspsStateRequest) FromJson(value string) error {
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

func (obj *isisLspsStateRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspsStateRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspsStateRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspsStateRequest) Clone() (IsisLspsStateRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspsStateRequest()
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

// IsisLspsStateRequest is the request to retrieve ISIS Link State PDU (LSP) information learned by the router.
type IsisLspsStateRequest interface {
	Validation
	// msg marshals IsisLspsStateRequest to protobuf object *otg.IsisLspsStateRequest
	// and doesn't set defaults
	msg() *otg.IsisLspsStateRequest
	// setMsg unmarshals IsisLspsStateRequest from protobuf object *otg.IsisLspsStateRequest
	// and doesn't set defaults
	setMsg(*otg.IsisLspsStateRequest) IsisLspsStateRequest
	// provides marshal interface
	Marshal() marshalIsisLspsStateRequest
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspsStateRequest
	// validate validates IsisLspsStateRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspsStateRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IsisRouterNames returns []string, set in IsisLspsStateRequest.
	IsisRouterNames() []string
	// SetIsisRouterNames assigns []string provided by user to IsisLspsStateRequest
	SetIsisRouterNames(value []string) IsisLspsStateRequest
}

// The names of ISIS routers for which learned information is requested. An empty list will return results for all ISIS routers.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// IsisRouterNames returns a []string
func (obj *isisLspsStateRequest) IsisRouterNames() []string {
	if obj.obj.IsisRouterNames == nil {
		obj.obj.IsisRouterNames = make([]string, 0)
	}
	return obj.obj.IsisRouterNames
}

// The names of ISIS routers for which learned information is requested. An empty list will return results for all ISIS routers.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// SetIsisRouterNames sets the []string value in the IsisLspsStateRequest object
func (obj *isisLspsStateRequest) SetIsisRouterNames(value []string) IsisLspsStateRequest {

	if obj.obj.IsisRouterNames == nil {
		obj.obj.IsisRouterNames = make([]string, 0)
	}
	obj.obj.IsisRouterNames = value

	return obj
}

func (obj *isisLspsStateRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspsStateRequest) setDefault() {

}
