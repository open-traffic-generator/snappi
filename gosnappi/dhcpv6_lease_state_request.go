package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6LeaseStateRequest *****
type dhcpv6LeaseStateRequest struct {
	validation
	obj          *otg.Dhcpv6LeaseStateRequest
	marshaller   marshalDhcpv6LeaseStateRequest
	unMarshaller unMarshalDhcpv6LeaseStateRequest
}

func NewDhcpv6LeaseStateRequest() Dhcpv6LeaseStateRequest {
	obj := dhcpv6LeaseStateRequest{obj: &otg.Dhcpv6LeaseStateRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6LeaseStateRequest) msg() *otg.Dhcpv6LeaseStateRequest {
	return obj.obj
}

func (obj *dhcpv6LeaseStateRequest) setMsg(msg *otg.Dhcpv6LeaseStateRequest) Dhcpv6LeaseStateRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6LeaseStateRequest struct {
	obj *dhcpv6LeaseStateRequest
}

type marshalDhcpv6LeaseStateRequest interface {
	// ToProto marshals Dhcpv6LeaseStateRequest to protobuf object *otg.Dhcpv6LeaseStateRequest
	ToProto() (*otg.Dhcpv6LeaseStateRequest, error)
	// ToPbText marshals Dhcpv6LeaseStateRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6LeaseStateRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6LeaseStateRequest to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6LeaseStateRequest struct {
	obj *dhcpv6LeaseStateRequest
}

type unMarshalDhcpv6LeaseStateRequest interface {
	// FromProto unmarshals Dhcpv6LeaseStateRequest from protobuf object *otg.Dhcpv6LeaseStateRequest
	FromProto(msg *otg.Dhcpv6LeaseStateRequest) (Dhcpv6LeaseStateRequest, error)
	// FromPbText unmarshals Dhcpv6LeaseStateRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6LeaseStateRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6LeaseStateRequest from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6LeaseStateRequest) Marshal() marshalDhcpv6LeaseStateRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6LeaseStateRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6LeaseStateRequest) Unmarshal() unMarshalDhcpv6LeaseStateRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6LeaseStateRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6LeaseStateRequest) ToProto() (*otg.Dhcpv6LeaseStateRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6LeaseStateRequest) FromProto(msg *otg.Dhcpv6LeaseStateRequest) (Dhcpv6LeaseStateRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6LeaseStateRequest) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6LeaseStateRequest) FromPbText(value string) error {
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

func (m *marshaldhcpv6LeaseStateRequest) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6LeaseStateRequest) FromYaml(value string) error {
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

func (m *marshaldhcpv6LeaseStateRequest) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6LeaseStateRequest) FromJson(value string) error {
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

func (obj *dhcpv6LeaseStateRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6LeaseStateRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6LeaseStateRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6LeaseStateRequest) Clone() (Dhcpv6LeaseStateRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6LeaseStateRequest()
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

// Dhcpv6LeaseStateRequest is the request to retrieve DHCP Server host allocated status.
type Dhcpv6LeaseStateRequest interface {
	Validation
	// msg marshals Dhcpv6LeaseStateRequest to protobuf object *otg.Dhcpv6LeaseStateRequest
	// and doesn't set defaults
	msg() *otg.Dhcpv6LeaseStateRequest
	// setMsg unmarshals Dhcpv6LeaseStateRequest from protobuf object *otg.Dhcpv6LeaseStateRequest
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6LeaseStateRequest) Dhcpv6LeaseStateRequest
	// provides marshal interface
	Marshal() marshalDhcpv6LeaseStateRequest
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6LeaseStateRequest
	// validate validates Dhcpv6LeaseStateRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6LeaseStateRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DhcpServerNames returns []string, set in Dhcpv6LeaseStateRequest.
	DhcpServerNames() []string
	// SetDhcpServerNames assigns []string provided by user to Dhcpv6LeaseStateRequest
	SetDhcpServerNames(value []string) Dhcpv6LeaseStateRequest
}

// The names of DHCPv6 server to return results for. An empty list will return results for all DHCPv6 servers.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6server/properties/name
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6server/properties/name
//
// DhcpServerNames returns a []string
func (obj *dhcpv6LeaseStateRequest) DhcpServerNames() []string {
	if obj.obj.DhcpServerNames == nil {
		obj.obj.DhcpServerNames = make([]string, 0)
	}
	return obj.obj.DhcpServerNames
}

// The names of DHCPv6 server to return results for. An empty list will return results for all DHCPv6 servers.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6server/properties/name
//
// x-constraint:
// - /components/schemas/Device.Dhcpv6server/properties/name
//
// SetDhcpServerNames sets the []string value in the Dhcpv6LeaseStateRequest object
func (obj *dhcpv6LeaseStateRequest) SetDhcpServerNames(value []string) Dhcpv6LeaseStateRequest {

	if obj.obj.DhcpServerNames == nil {
		obj.obj.DhcpServerNames = make([]string, 0)
	}
	obj.obj.DhcpServerNames = value

	return obj
}

func (obj *dhcpv6LeaseStateRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv6LeaseStateRequest) setDefault() {

}
