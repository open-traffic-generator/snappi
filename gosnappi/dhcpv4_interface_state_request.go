package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv4InterfaceStateRequest *****
type dhcpv4InterfaceStateRequest struct {
	validation
	obj          *otg.Dhcpv4InterfaceStateRequest
	marshaller   marshalDhcpv4InterfaceStateRequest
	unMarshaller unMarshalDhcpv4InterfaceStateRequest
}

func NewDhcpv4InterfaceStateRequest() Dhcpv4InterfaceStateRequest {
	obj := dhcpv4InterfaceStateRequest{obj: &otg.Dhcpv4InterfaceStateRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv4InterfaceStateRequest) msg() *otg.Dhcpv4InterfaceStateRequest {
	return obj.obj
}

func (obj *dhcpv4InterfaceStateRequest) setMsg(msg *otg.Dhcpv4InterfaceStateRequest) Dhcpv4InterfaceStateRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv4InterfaceStateRequest struct {
	obj *dhcpv4InterfaceStateRequest
}

type marshalDhcpv4InterfaceStateRequest interface {
	// ToProto marshals Dhcpv4InterfaceStateRequest to protobuf object *otg.Dhcpv4InterfaceStateRequest
	ToProto() (*otg.Dhcpv4InterfaceStateRequest, error)
	// ToPbText marshals Dhcpv4InterfaceStateRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv4InterfaceStateRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv4InterfaceStateRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv4InterfaceStateRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv4InterfaceStateRequest struct {
	obj *dhcpv4InterfaceStateRequest
}

type unMarshalDhcpv4InterfaceStateRequest interface {
	// FromProto unmarshals Dhcpv4InterfaceStateRequest from protobuf object *otg.Dhcpv4InterfaceStateRequest
	FromProto(msg *otg.Dhcpv4InterfaceStateRequest) (Dhcpv4InterfaceStateRequest, error)
	// FromPbText unmarshals Dhcpv4InterfaceStateRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv4InterfaceStateRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv4InterfaceStateRequest from JSON text
	FromJson(value string) error
}

func (obj *dhcpv4InterfaceStateRequest) Marshal() marshalDhcpv4InterfaceStateRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv4InterfaceStateRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv4InterfaceStateRequest) Unmarshal() unMarshalDhcpv4InterfaceStateRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv4InterfaceStateRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv4InterfaceStateRequest) ToProto() (*otg.Dhcpv4InterfaceStateRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv4InterfaceStateRequest) FromProto(msg *otg.Dhcpv4InterfaceStateRequest) (Dhcpv4InterfaceStateRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv4InterfaceStateRequest) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv4InterfaceStateRequest) FromPbText(value string) error {
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

func (m *marshaldhcpv4InterfaceStateRequest) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv4InterfaceStateRequest) FromYaml(value string) error {
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

func (m *marshaldhcpv4InterfaceStateRequest) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv4InterfaceStateRequest) ToJson() (string, error) {
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

func (m *unMarshaldhcpv4InterfaceStateRequest) FromJson(value string) error {
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

func (obj *dhcpv4InterfaceStateRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv4InterfaceStateRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv4InterfaceStateRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv4InterfaceStateRequest) Clone() (Dhcpv4InterfaceStateRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv4InterfaceStateRequest()
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

// Dhcpv4InterfaceStateRequest is the request for assigned IPv4 address information associated with DHCP Client sessions.
type Dhcpv4InterfaceStateRequest interface {
	Validation
	// msg marshals Dhcpv4InterfaceStateRequest to protobuf object *otg.Dhcpv4InterfaceStateRequest
	// and doesn't set defaults
	msg() *otg.Dhcpv4InterfaceStateRequest
	// setMsg unmarshals Dhcpv4InterfaceStateRequest from protobuf object *otg.Dhcpv4InterfaceStateRequest
	// and doesn't set defaults
	setMsg(*otg.Dhcpv4InterfaceStateRequest) Dhcpv4InterfaceStateRequest
	// provides marshal interface
	Marshal() marshalDhcpv4InterfaceStateRequest
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv4InterfaceStateRequest
	// validate validates Dhcpv4InterfaceStateRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv4InterfaceStateRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DhcpClientNames returns []string, set in Dhcpv4InterfaceStateRequest.
	DhcpClientNames() []string
	// SetDhcpClientNames assigns []string provided by user to Dhcpv4InterfaceStateRequest
	SetDhcpClientNames(value []string) Dhcpv4InterfaceStateRequest
}

// The names of DHCPv4 client to return results for. An empty list will return results for all DHCPv4 Client address information.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// DhcpClientNames returns a []string
func (obj *dhcpv4InterfaceStateRequest) DhcpClientNames() []string {
	if obj.obj.DhcpClientNames == nil {
		obj.obj.DhcpClientNames = make([]string, 0)
	}
	return obj.obj.DhcpClientNames
}

// The names of DHCPv4 client to return results for. An empty list will return results for all DHCPv4 Client address information.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4client/properties/name
//
// SetDhcpClientNames sets the []string value in the Dhcpv4InterfaceStateRequest object
func (obj *dhcpv4InterfaceStateRequest) SetDhcpClientNames(value []string) Dhcpv4InterfaceStateRequest {

	if obj.obj.DhcpClientNames == nil {
		obj.obj.DhcpClientNames = make([]string, 0)
	}
	obj.obj.DhcpClientNames = value

	return obj
}

func (obj *dhcpv4InterfaceStateRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv4InterfaceStateRequest) setDefault() {

}
