package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsasStateRequest *****
type ospfv2LsasStateRequest struct {
	validation
	obj          *otg.Ospfv2LsasStateRequest
	marshaller   marshalOspfv2LsasStateRequest
	unMarshaller unMarshalOspfv2LsasStateRequest
}

func NewOspfv2LsasStateRequest() Ospfv2LsasStateRequest {
	obj := ospfv2LsasStateRequest{obj: &otg.Ospfv2LsasStateRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsasStateRequest) msg() *otg.Ospfv2LsasStateRequest {
	return obj.obj
}

func (obj *ospfv2LsasStateRequest) setMsg(msg *otg.Ospfv2LsasStateRequest) Ospfv2LsasStateRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsasStateRequest struct {
	obj *ospfv2LsasStateRequest
}

type marshalOspfv2LsasStateRequest interface {
	// ToProto marshals Ospfv2LsasStateRequest to protobuf object *otg.Ospfv2LsasStateRequest
	ToProto() (*otg.Ospfv2LsasStateRequest, error)
	// ToPbText marshals Ospfv2LsasStateRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsasStateRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsasStateRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv2LsasStateRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv2LsasStateRequest struct {
	obj *ospfv2LsasStateRequest
}

type unMarshalOspfv2LsasStateRequest interface {
	// FromProto unmarshals Ospfv2LsasStateRequest from protobuf object *otg.Ospfv2LsasStateRequest
	FromProto(msg *otg.Ospfv2LsasStateRequest) (Ospfv2LsasStateRequest, error)
	// FromPbText unmarshals Ospfv2LsasStateRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsasStateRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsasStateRequest from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsasStateRequest) Marshal() marshalOspfv2LsasStateRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsasStateRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsasStateRequest) Unmarshal() unMarshalOspfv2LsasStateRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsasStateRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsasStateRequest) ToProto() (*otg.Ospfv2LsasStateRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsasStateRequest) FromProto(msg *otg.Ospfv2LsasStateRequest) (Ospfv2LsasStateRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsasStateRequest) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsasStateRequest) FromPbText(value string) error {
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

func (m *marshalospfv2LsasStateRequest) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsasStateRequest) FromYaml(value string) error {
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

func (m *marshalospfv2LsasStateRequest) ToJsonRaw() (string, error) {
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

func (m *marshalospfv2LsasStateRequest) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsasStateRequest) FromJson(value string) error {
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

func (obj *ospfv2LsasStateRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsasStateRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsasStateRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsasStateRequest) Clone() (Ospfv2LsasStateRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsasStateRequest()
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

// Ospfv2LsasStateRequest is the request to retrieve OSPFv2 Link State Advertisements (LSA) information learned by the routers.
type Ospfv2LsasStateRequest interface {
	Validation
	// msg marshals Ospfv2LsasStateRequest to protobuf object *otg.Ospfv2LsasStateRequest
	// and doesn't set defaults
	msg() *otg.Ospfv2LsasStateRequest
	// setMsg unmarshals Ospfv2LsasStateRequest from protobuf object *otg.Ospfv2LsasStateRequest
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsasStateRequest) Ospfv2LsasStateRequest
	// provides marshal interface
	Marshal() marshalOspfv2LsasStateRequest
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsasStateRequest
	// validate validates Ospfv2LsasStateRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsasStateRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in Ospfv2LsasStateRequest.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to Ospfv2LsasStateRequest
	SetRouterNames(value []string) Ospfv2LsasStateRequest
}

// The names of OSPFv2 routers for which learned information is requested. An empty list will return results for all OSPFv2 routers.
//
// x-constraint:
// - /components/schemas/Device.Ospfv2Router/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ospfv2Router/properties/name
//
// RouterNames returns a []string
func (obj *ospfv2LsasStateRequest) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of OSPFv2 routers for which learned information is requested. An empty list will return results for all OSPFv2 routers.
//
// x-constraint:
// - /components/schemas/Device.Ospfv2Router/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ospfv2Router/properties/name
//
// SetRouterNames sets the []string value in the Ospfv2LsasStateRequest object
func (obj *ospfv2LsasStateRequest) SetRouterNames(value []string) Ospfv2LsasStateRequest {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

func (obj *ospfv2LsasStateRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2LsasStateRequest) setDefault() {

}
