package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LdpBindingsStateRequest *****
type ldpBindingsStateRequest struct {
	validation
	obj          *otg.LdpBindingsStateRequest
	marshaller   marshalLdpBindingsStateRequest
	unMarshaller unMarshalLdpBindingsStateRequest
}

func NewLdpBindingsStateRequest() LdpBindingsStateRequest {
	obj := ldpBindingsStateRequest{obj: &otg.LdpBindingsStateRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *ldpBindingsStateRequest) msg() *otg.LdpBindingsStateRequest {
	return obj.obj
}

func (obj *ldpBindingsStateRequest) setMsg(msg *otg.LdpBindingsStateRequest) LdpBindingsStateRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalldpBindingsStateRequest struct {
	obj *ldpBindingsStateRequest
}

type marshalLdpBindingsStateRequest interface {
	// ToProto marshals LdpBindingsStateRequest to protobuf object *otg.LdpBindingsStateRequest
	ToProto() (*otg.LdpBindingsStateRequest, error)
	// ToPbText marshals LdpBindingsStateRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LdpBindingsStateRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals LdpBindingsStateRequest to JSON text
	ToJson() (string, error)
}

type unMarshalldpBindingsStateRequest struct {
	obj *ldpBindingsStateRequest
}

type unMarshalLdpBindingsStateRequest interface {
	// FromProto unmarshals LdpBindingsStateRequest from protobuf object *otg.LdpBindingsStateRequest
	FromProto(msg *otg.LdpBindingsStateRequest) (LdpBindingsStateRequest, error)
	// FromPbText unmarshals LdpBindingsStateRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LdpBindingsStateRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LdpBindingsStateRequest from JSON text
	FromJson(value string) error
}

func (obj *ldpBindingsStateRequest) Marshal() marshalLdpBindingsStateRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalldpBindingsStateRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *ldpBindingsStateRequest) Unmarshal() unMarshalLdpBindingsStateRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalldpBindingsStateRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalldpBindingsStateRequest) ToProto() (*otg.LdpBindingsStateRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalldpBindingsStateRequest) FromProto(msg *otg.LdpBindingsStateRequest) (LdpBindingsStateRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalldpBindingsStateRequest) ToPbText() (string, error) {
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

func (m *unMarshalldpBindingsStateRequest) FromPbText(value string) error {
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

func (m *marshalldpBindingsStateRequest) ToYaml() (string, error) {
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

func (m *unMarshalldpBindingsStateRequest) FromYaml(value string) error {
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

func (m *marshalldpBindingsStateRequest) ToJson() (string, error) {
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

func (m *unMarshalldpBindingsStateRequest) FromJson(value string) error {
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

func (obj *ldpBindingsStateRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ldpBindingsStateRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ldpBindingsStateRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ldpBindingsStateRequest) Clone() (LdpBindingsStateRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLdpBindingsStateRequest()
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

// LdpBindingsStateRequest is the request to retrieve the FEC-label bindings received by LDP routers.
type LdpBindingsStateRequest interface {
	Validation
	// msg marshals LdpBindingsStateRequest to protobuf object *otg.LdpBindingsStateRequest
	// and doesn't set defaults
	msg() *otg.LdpBindingsStateRequest
	// setMsg unmarshals LdpBindingsStateRequest from protobuf object *otg.LdpBindingsStateRequest
	// and doesn't set defaults
	setMsg(*otg.LdpBindingsStateRequest) LdpBindingsStateRequest
	// provides marshal interface
	Marshal() marshalLdpBindingsStateRequest
	// provides unmarshal interface
	Unmarshal() unMarshalLdpBindingsStateRequest
	// validate validates LdpBindingsStateRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LdpBindingsStateRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LdpRouterNames returns []string, set in LdpBindingsStateRequest.
	LdpRouterNames() []string
	// SetLdpRouterNames assigns []string provided by user to LdpBindingsStateRequest
	SetLdpRouterNames(value []string) LdpBindingsStateRequest
}

// The names of LDP routers for which received label bindings are requested. An empty list will return results for all LDP routers.
//
// x-constraint:
// - /components/schemas/Device.LdpRouter/properties/name
//
// LdpRouterNames returns a []string
func (obj *ldpBindingsStateRequest) LdpRouterNames() []string {
	if obj.obj.LdpRouterNames == nil {
		obj.obj.LdpRouterNames = make([]string, 0)
	}
	return obj.obj.LdpRouterNames
}

// The names of LDP routers for which received label bindings are requested. An empty list will return results for all LDP routers.
//
// x-constraint:
// - /components/schemas/Device.LdpRouter/properties/name
//
// SetLdpRouterNames sets the []string value in the LdpBindingsStateRequest object
func (obj *ldpBindingsStateRequest) SetLdpRouterNames(value []string) LdpBindingsStateRequest {

	if obj.obj.LdpRouterNames == nil {
		obj.obj.LdpRouterNames = make([]string, 0)
	}
	obj.obj.LdpRouterNames = value

	return obj
}

func (obj *ldpBindingsStateRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ldpBindingsStateRequest) setDefault() {

}
