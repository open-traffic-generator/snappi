package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv4LeaseStateRequest *****
type dhcpv4LeaseStateRequest struct {
	validation
	obj          *otg.Dhcpv4LeaseStateRequest
	marshaller   marshalDhcpv4LeaseStateRequest
	unMarshaller unMarshalDhcpv4LeaseStateRequest
}

func NewDhcpv4LeaseStateRequest() Dhcpv4LeaseStateRequest {
	obj := dhcpv4LeaseStateRequest{obj: &otg.Dhcpv4LeaseStateRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv4LeaseStateRequest) msg() *otg.Dhcpv4LeaseStateRequest {
	return obj.obj
}

func (obj *dhcpv4LeaseStateRequest) setMsg(msg *otg.Dhcpv4LeaseStateRequest) Dhcpv4LeaseStateRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv4LeaseStateRequest struct {
	obj *dhcpv4LeaseStateRequest
}

type marshalDhcpv4LeaseStateRequest interface {
	// ToProto marshals Dhcpv4LeaseStateRequest to protobuf object *otg.Dhcpv4LeaseStateRequest
	ToProto() (*otg.Dhcpv4LeaseStateRequest, error)
	// ToPbText marshals Dhcpv4LeaseStateRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv4LeaseStateRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv4LeaseStateRequest to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv4LeaseStateRequest struct {
	obj *dhcpv4LeaseStateRequest
}

type unMarshalDhcpv4LeaseStateRequest interface {
	// FromProto unmarshals Dhcpv4LeaseStateRequest from protobuf object *otg.Dhcpv4LeaseStateRequest
	FromProto(msg *otg.Dhcpv4LeaseStateRequest) (Dhcpv4LeaseStateRequest, error)
	// FromPbText unmarshals Dhcpv4LeaseStateRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv4LeaseStateRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv4LeaseStateRequest from JSON text
	FromJson(value string) error
}

func (obj *dhcpv4LeaseStateRequest) Marshal() marshalDhcpv4LeaseStateRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv4LeaseStateRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv4LeaseStateRequest) Unmarshal() unMarshalDhcpv4LeaseStateRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv4LeaseStateRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv4LeaseStateRequest) ToProto() (*otg.Dhcpv4LeaseStateRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv4LeaseStateRequest) FromProto(msg *otg.Dhcpv4LeaseStateRequest) (Dhcpv4LeaseStateRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv4LeaseStateRequest) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv4LeaseStateRequest) FromPbText(value string) error {
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

func (m *marshaldhcpv4LeaseStateRequest) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv4LeaseStateRequest) FromYaml(value string) error {
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

func (m *marshaldhcpv4LeaseStateRequest) ToJson() (string, error) {
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

func (m *unMarshaldhcpv4LeaseStateRequest) FromJson(value string) error {
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

func (obj *dhcpv4LeaseStateRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv4LeaseStateRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv4LeaseStateRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv4LeaseStateRequest) Clone() (Dhcpv4LeaseStateRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv4LeaseStateRequest()
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

// Dhcpv4LeaseStateRequest is the request to retrieve DHCP Server host allocated status.
type Dhcpv4LeaseStateRequest interface {
	Validation
	// msg marshals Dhcpv4LeaseStateRequest to protobuf object *otg.Dhcpv4LeaseStateRequest
	// and doesn't set defaults
	msg() *otg.Dhcpv4LeaseStateRequest
	// setMsg unmarshals Dhcpv4LeaseStateRequest from protobuf object *otg.Dhcpv4LeaseStateRequest
	// and doesn't set defaults
	setMsg(*otg.Dhcpv4LeaseStateRequest) Dhcpv4LeaseStateRequest
	// provides marshal interface
	Marshal() marshalDhcpv4LeaseStateRequest
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv4LeaseStateRequest
	// validate validates Dhcpv4LeaseStateRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv4LeaseStateRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DhcpServerNames returns []string, set in Dhcpv4LeaseStateRequest.
	DhcpServerNames() []string
	// SetDhcpServerNames assigns []string provided by user to Dhcpv4LeaseStateRequest
	SetDhcpServerNames(value []string) Dhcpv4LeaseStateRequest
}

// The names of DHCPv4 server to return results for. An empty list will return results for all DHCPv4 servers.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4server/properties/name
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4server/properties/name
//
// DhcpServerNames returns a []string
func (obj *dhcpv4LeaseStateRequest) DhcpServerNames() []string {
	if obj.obj.DhcpServerNames == nil {
		obj.obj.DhcpServerNames = make([]string, 0)
	}
	return obj.obj.DhcpServerNames
}

// The names of DHCPv4 server to return results for. An empty list will return results for all DHCPv4 servers.
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4server/properties/name
//
// x-constraint:
// - /components/schemas/Device.Dhcpv4server/properties/name
//
// SetDhcpServerNames sets the []string value in the Dhcpv4LeaseStateRequest object
func (obj *dhcpv4LeaseStateRequest) SetDhcpServerNames(value []string) Dhcpv4LeaseStateRequest {

	if obj.obj.DhcpServerNames == nil {
		obj.obj.DhcpServerNames = make([]string, 0)
	}
	obj.obj.DhcpServerNames = value

	return obj
}

func (obj *dhcpv4LeaseStateRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv4LeaseStateRequest) setDefault() {

}
