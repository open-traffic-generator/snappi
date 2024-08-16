package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6InterfaceStateRequest *****
type dhcpv6InterfaceStateRequest struct {
	validation
	obj          *otg.Dhcpv6InterfaceStateRequest
	marshaller   marshalDhcpv6InterfaceStateRequest
	unMarshaller unMarshalDhcpv6InterfaceStateRequest
}

func NewDhcpv6InterfaceStateRequest() Dhcpv6InterfaceStateRequest {
	obj := dhcpv6InterfaceStateRequest{obj: &otg.Dhcpv6InterfaceStateRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6InterfaceStateRequest) msg() *otg.Dhcpv6InterfaceStateRequest {
	return obj.obj
}

func (obj *dhcpv6InterfaceStateRequest) setMsg(msg *otg.Dhcpv6InterfaceStateRequest) Dhcpv6InterfaceStateRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6InterfaceStateRequest struct {
	obj *dhcpv6InterfaceStateRequest
}

type marshalDhcpv6InterfaceStateRequest interface {
	// ToProto marshals Dhcpv6InterfaceStateRequest to protobuf object *otg.Dhcpv6InterfaceStateRequest
	ToProto() (*otg.Dhcpv6InterfaceStateRequest, error)
	// ToPbText marshals Dhcpv6InterfaceStateRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6InterfaceStateRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6InterfaceStateRequest to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6InterfaceStateRequest struct {
	obj *dhcpv6InterfaceStateRequest
}

type unMarshalDhcpv6InterfaceStateRequest interface {
	// FromProto unmarshals Dhcpv6InterfaceStateRequest from protobuf object *otg.Dhcpv6InterfaceStateRequest
	FromProto(msg *otg.Dhcpv6InterfaceStateRequest) (Dhcpv6InterfaceStateRequest, error)
	// FromPbText unmarshals Dhcpv6InterfaceStateRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6InterfaceStateRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6InterfaceStateRequest from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6InterfaceStateRequest) Marshal() marshalDhcpv6InterfaceStateRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6InterfaceStateRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6InterfaceStateRequest) Unmarshal() unMarshalDhcpv6InterfaceStateRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6InterfaceStateRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6InterfaceStateRequest) ToProto() (*otg.Dhcpv6InterfaceStateRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6InterfaceStateRequest) FromProto(msg *otg.Dhcpv6InterfaceStateRequest) (Dhcpv6InterfaceStateRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6InterfaceStateRequest) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceStateRequest) FromPbText(value string) error {
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

func (m *marshaldhcpv6InterfaceStateRequest) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceStateRequest) FromYaml(value string) error {
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

func (m *marshaldhcpv6InterfaceStateRequest) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceStateRequest) FromJson(value string) error {
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

func (obj *dhcpv6InterfaceStateRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6InterfaceStateRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6InterfaceStateRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6InterfaceStateRequest) Clone() (Dhcpv6InterfaceStateRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6InterfaceStateRequest()
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

// Dhcpv6InterfaceStateRequest is the request for assigned IPv6 address information associated with DHCP Client sessions.
type Dhcpv6InterfaceStateRequest interface {
	Validation
	// msg marshals Dhcpv6InterfaceStateRequest to protobuf object *otg.Dhcpv6InterfaceStateRequest
	// and doesn't set defaults
	msg() *otg.Dhcpv6InterfaceStateRequest
	// setMsg unmarshals Dhcpv6InterfaceStateRequest from protobuf object *otg.Dhcpv6InterfaceStateRequest
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6InterfaceStateRequest) Dhcpv6InterfaceStateRequest
	// provides marshal interface
	Marshal() marshalDhcpv6InterfaceStateRequest
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6InterfaceStateRequest
	// validate validates Dhcpv6InterfaceStateRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6InterfaceStateRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DhcpClientNames returns []string, set in Dhcpv6InterfaceStateRequest.
	DhcpClientNames() []string
	// SetDhcpClientNames assigns []string provided by user to Dhcpv6InterfaceStateRequest
	SetDhcpClientNames(value []string) Dhcpv6InterfaceStateRequest
}

// The names of DHCPv6 client to return results for. An empty list will return results for all DHCPv6 Client address information.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6client/properties/name
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6client/properties/name
//
// DhcpClientNames returns a []string
func (obj *dhcpv6InterfaceStateRequest) DhcpClientNames() []string {
	if obj.obj.DhcpClientNames == nil {
		obj.obj.DhcpClientNames = make([]string, 0)
	}
	return obj.obj.DhcpClientNames
}

// The names of DHCPv6 client to return results for. An empty list will return results for all DHCPv6 Client address information.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6client/properties/name
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6client/properties/name
//
// SetDhcpClientNames sets the []string value in the Dhcpv6InterfaceStateRequest object
func (obj *dhcpv6InterfaceStateRequest) SetDhcpClientNames(value []string) Dhcpv6InterfaceStateRequest {

	if obj.obj.DhcpClientNames == nil {
		obj.obj.DhcpClientNames = make([]string, 0)
	}
	obj.obj.DhcpClientNames = value

	return obj
}

func (obj *dhcpv6InterfaceStateRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv6InterfaceStateRequest) setDefault() {

}
