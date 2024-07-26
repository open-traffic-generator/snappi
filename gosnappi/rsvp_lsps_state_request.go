package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpLspsStateRequest *****
type rsvpLspsStateRequest struct {
	validation
	obj          *otg.RsvpLspsStateRequest
	marshaller   marshalRsvpLspsStateRequest
	unMarshaller unMarshalRsvpLspsStateRequest
}

func NewRsvpLspsStateRequest() RsvpLspsStateRequest {
	obj := rsvpLspsStateRequest{obj: &otg.RsvpLspsStateRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpLspsStateRequest) msg() *otg.RsvpLspsStateRequest {
	return obj.obj
}

func (obj *rsvpLspsStateRequest) setMsg(msg *otg.RsvpLspsStateRequest) RsvpLspsStateRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpLspsStateRequest struct {
	obj *rsvpLspsStateRequest
}

type marshalRsvpLspsStateRequest interface {
	// ToProto marshals RsvpLspsStateRequest to protobuf object *otg.RsvpLspsStateRequest
	ToProto() (*otg.RsvpLspsStateRequest, error)
	// ToPbText marshals RsvpLspsStateRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpLspsStateRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpLspsStateRequest to JSON text
	ToJson() (string, error)
}

type unMarshalrsvpLspsStateRequest struct {
	obj *rsvpLspsStateRequest
}

type unMarshalRsvpLspsStateRequest interface {
	// FromProto unmarshals RsvpLspsStateRequest from protobuf object *otg.RsvpLspsStateRequest
	FromProto(msg *otg.RsvpLspsStateRequest) (RsvpLspsStateRequest, error)
	// FromPbText unmarshals RsvpLspsStateRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpLspsStateRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpLspsStateRequest from JSON text
	FromJson(value string) error
}

func (obj *rsvpLspsStateRequest) Marshal() marshalRsvpLspsStateRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpLspsStateRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpLspsStateRequest) Unmarshal() unMarshalRsvpLspsStateRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpLspsStateRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpLspsStateRequest) ToProto() (*otg.RsvpLspsStateRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpLspsStateRequest) FromProto(msg *otg.RsvpLspsStateRequest) (RsvpLspsStateRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpLspsStateRequest) ToPbText() (string, error) {
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

func (m *unMarshalrsvpLspsStateRequest) FromPbText(value string) error {
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

func (m *marshalrsvpLspsStateRequest) ToYaml() (string, error) {
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

func (m *unMarshalrsvpLspsStateRequest) FromYaml(value string) error {
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

func (m *marshalrsvpLspsStateRequest) ToJson() (string, error) {
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

func (m *unMarshalrsvpLspsStateRequest) FromJson(value string) error {
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

func (obj *rsvpLspsStateRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpLspsStateRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpLspsStateRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpLspsStateRequest) Clone() (RsvpLspsStateRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpLspsStateRequest()
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

// RsvpLspsStateRequest is the request to retrieve RSVP Label Switched Path (LSP) information learned by the router.
type RsvpLspsStateRequest interface {
	Validation
	// msg marshals RsvpLspsStateRequest to protobuf object *otg.RsvpLspsStateRequest
	// and doesn't set defaults
	msg() *otg.RsvpLspsStateRequest
	// setMsg unmarshals RsvpLspsStateRequest from protobuf object *otg.RsvpLspsStateRequest
	// and doesn't set defaults
	setMsg(*otg.RsvpLspsStateRequest) RsvpLspsStateRequest
	// provides marshal interface
	Marshal() marshalRsvpLspsStateRequest
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpLspsStateRequest
	// validate validates RsvpLspsStateRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpLspsStateRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RsvpRouterNames returns []string, set in RsvpLspsStateRequest.
	RsvpRouterNames() []string
	// SetRsvpRouterNames assigns []string provided by user to RsvpLspsStateRequest
	SetRsvpRouterNames(value []string) RsvpLspsStateRequest
}

// The names of RSVP-TE routers for which learned information is requested. An empty list will return results for all RSVP=TE routers.
//
// x-constraint:
// - /components/schemas/Device.Rsvp/properties/name
//
// x-constraint:
// - /components/schemas/Device.Rsvp/properties/name
//
// RsvpRouterNames returns a []string
func (obj *rsvpLspsStateRequest) RsvpRouterNames() []string {
	if obj.obj.RsvpRouterNames == nil {
		obj.obj.RsvpRouterNames = make([]string, 0)
	}
	return obj.obj.RsvpRouterNames
}

// The names of RSVP-TE routers for which learned information is requested. An empty list will return results for all RSVP=TE routers.
//
// x-constraint:
// - /components/schemas/Device.Rsvp/properties/name
//
// x-constraint:
// - /components/schemas/Device.Rsvp/properties/name
//
// SetRsvpRouterNames sets the []string value in the RsvpLspsStateRequest object
func (obj *rsvpLspsStateRequest) SetRsvpRouterNames(value []string) RsvpLspsStateRequest {

	if obj.obj.RsvpRouterNames == nil {
		obj.obj.RsvpRouterNames = make([]string, 0)
	}
	obj.obj.RsvpRouterNames = value

	return obj
}

func (obj *rsvpLspsStateRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *rsvpLspsStateRequest) setDefault() {

}
