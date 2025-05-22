package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3LsasStateRequest *****
type ospfv3LsasStateRequest struct {
	validation
	obj          *otg.Ospfv3LsasStateRequest
	marshaller   marshalOspfv3LsasStateRequest
	unMarshaller unMarshalOspfv3LsasStateRequest
}

func NewOspfv3LsasStateRequest() Ospfv3LsasStateRequest {
	obj := ospfv3LsasStateRequest{obj: &otg.Ospfv3LsasStateRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3LsasStateRequest) msg() *otg.Ospfv3LsasStateRequest {
	return obj.obj
}

func (obj *ospfv3LsasStateRequest) setMsg(msg *otg.Ospfv3LsasStateRequest) Ospfv3LsasStateRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3LsasStateRequest struct {
	obj *ospfv3LsasStateRequest
}

type marshalOspfv3LsasStateRequest interface {
	// ToProto marshals Ospfv3LsasStateRequest to protobuf object *otg.Ospfv3LsasStateRequest
	ToProto() (*otg.Ospfv3LsasStateRequest, error)
	// ToPbText marshals Ospfv3LsasStateRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3LsasStateRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3LsasStateRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv3LsasStateRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv3LsasStateRequest struct {
	obj *ospfv3LsasStateRequest
}

type unMarshalOspfv3LsasStateRequest interface {
	// FromProto unmarshals Ospfv3LsasStateRequest from protobuf object *otg.Ospfv3LsasStateRequest
	FromProto(msg *otg.Ospfv3LsasStateRequest) (Ospfv3LsasStateRequest, error)
	// FromPbText unmarshals Ospfv3LsasStateRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3LsasStateRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3LsasStateRequest from JSON text
	FromJson(value string) error
}

func (obj *ospfv3LsasStateRequest) Marshal() marshalOspfv3LsasStateRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3LsasStateRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3LsasStateRequest) Unmarshal() unMarshalOspfv3LsasStateRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3LsasStateRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3LsasStateRequest) ToProto() (*otg.Ospfv3LsasStateRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3LsasStateRequest) FromProto(msg *otg.Ospfv3LsasStateRequest) (Ospfv3LsasStateRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3LsasStateRequest) ToPbText() (string, error) {
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

func (m *unMarshalospfv3LsasStateRequest) FromPbText(value string) error {
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

func (m *marshalospfv3LsasStateRequest) ToYaml() (string, error) {
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

func (m *unMarshalospfv3LsasStateRequest) FromYaml(value string) error {
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

func (m *marshalospfv3LsasStateRequest) ToJsonRaw() (string, error) {
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

func (m *marshalospfv3LsasStateRequest) ToJson() (string, error) {
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

func (m *unMarshalospfv3LsasStateRequest) FromJson(value string) error {
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

func (obj *ospfv3LsasStateRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3LsasStateRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3LsasStateRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3LsasStateRequest) Clone() (Ospfv3LsasStateRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3LsasStateRequest()
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

// Ospfv3LsasStateRequest is the request to retrieve OSPFv3 Link State Advertisements (LSA) information learned by the routers.
type Ospfv3LsasStateRequest interface {
	Validation
	// msg marshals Ospfv3LsasStateRequest to protobuf object *otg.Ospfv3LsasStateRequest
	// and doesn't set defaults
	msg() *otg.Ospfv3LsasStateRequest
	// setMsg unmarshals Ospfv3LsasStateRequest from protobuf object *otg.Ospfv3LsasStateRequest
	// and doesn't set defaults
	setMsg(*otg.Ospfv3LsasStateRequest) Ospfv3LsasStateRequest
	// provides marshal interface
	Marshal() marshalOspfv3LsasStateRequest
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3LsasStateRequest
	// validate validates Ospfv3LsasStateRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3LsasStateRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in Ospfv3LsasStateRequest.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to Ospfv3LsasStateRequest
	SetRouterNames(value []string) Ospfv3LsasStateRequest
}

// The names of OSPFv3 routers for which learned information is requested. An empty list will return results for all OSPFv3 routers.
//
// x-constraint:
// - /components/schemas/Ospfv3.RouterInstance/properties/name
//
// x-constraint:
// - /components/schemas/Ospfv3.RouterInstance/properties/name
//
// RouterNames returns a []string
func (obj *ospfv3LsasStateRequest) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of OSPFv3 routers for which learned information is requested. An empty list will return results for all OSPFv3 routers.
//
// x-constraint:
// - /components/schemas/Ospfv3.RouterInstance/properties/name
//
// x-constraint:
// - /components/schemas/Ospfv3.RouterInstance/properties/name
//
// SetRouterNames sets the []string value in the Ospfv3LsasStateRequest object
func (obj *ospfv3LsasStateRequest) SetRouterNames(value []string) Ospfv3LsasStateRequest {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

func (obj *ospfv3LsasStateRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv3LsasStateRequest) setDefault() {

}
