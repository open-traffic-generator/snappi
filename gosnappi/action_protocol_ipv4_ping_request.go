package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIpv4PingRequest *****
type actionProtocolIpv4PingRequest struct {
	validation
	obj          *otg.ActionProtocolIpv4PingRequest
	marshaller   marshalActionProtocolIpv4PingRequest
	unMarshaller unMarshalActionProtocolIpv4PingRequest
}

func NewActionProtocolIpv4PingRequest() ActionProtocolIpv4PingRequest {
	obj := actionProtocolIpv4PingRequest{obj: &otg.ActionProtocolIpv4PingRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIpv4PingRequest) msg() *otg.ActionProtocolIpv4PingRequest {
	return obj.obj
}

func (obj *actionProtocolIpv4PingRequest) setMsg(msg *otg.ActionProtocolIpv4PingRequest) ActionProtocolIpv4PingRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIpv4PingRequest struct {
	obj *actionProtocolIpv4PingRequest
}

type marshalActionProtocolIpv4PingRequest interface {
	// ToProto marshals ActionProtocolIpv4PingRequest to protobuf object *otg.ActionProtocolIpv4PingRequest
	ToProto() (*otg.ActionProtocolIpv4PingRequest, error)
	// ToPbText marshals ActionProtocolIpv4PingRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIpv4PingRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIpv4PingRequest to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIpv4PingRequest struct {
	obj *actionProtocolIpv4PingRequest
}

type unMarshalActionProtocolIpv4PingRequest interface {
	// FromProto unmarshals ActionProtocolIpv4PingRequest from protobuf object *otg.ActionProtocolIpv4PingRequest
	FromProto(msg *otg.ActionProtocolIpv4PingRequest) (ActionProtocolIpv4PingRequest, error)
	// FromPbText unmarshals ActionProtocolIpv4PingRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIpv4PingRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIpv4PingRequest from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIpv4PingRequest) Marshal() marshalActionProtocolIpv4PingRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIpv4PingRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIpv4PingRequest) Unmarshal() unMarshalActionProtocolIpv4PingRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIpv4PingRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIpv4PingRequest) ToProto() (*otg.ActionProtocolIpv4PingRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIpv4PingRequest) FromProto(msg *otg.ActionProtocolIpv4PingRequest) (ActionProtocolIpv4PingRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIpv4PingRequest) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIpv4PingRequest) FromPbText(value string) error {
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

func (m *marshalactionProtocolIpv4PingRequest) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIpv4PingRequest) FromYaml(value string) error {
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

func (m *marshalactionProtocolIpv4PingRequest) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIpv4PingRequest) FromJson(value string) error {
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

func (obj *actionProtocolIpv4PingRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIpv4PingRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIpv4PingRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIpv4PingRequest) Clone() (ActionProtocolIpv4PingRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIpv4PingRequest()
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

// ActionProtocolIpv4PingRequest is request for initiating ping between a single source and destination pair.
// For ping request, 1 IPv4 ICMP Echo Request shall be sent and wait for ping response to either succeed or time out. The API wait timeout for each request shall be 300ms.
type ActionProtocolIpv4PingRequest interface {
	Validation
	// msg marshals ActionProtocolIpv4PingRequest to protobuf object *otg.ActionProtocolIpv4PingRequest
	// and doesn't set defaults
	msg() *otg.ActionProtocolIpv4PingRequest
	// setMsg unmarshals ActionProtocolIpv4PingRequest from protobuf object *otg.ActionProtocolIpv4PingRequest
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIpv4PingRequest) ActionProtocolIpv4PingRequest
	// provides marshal interface
	Marshal() marshalActionProtocolIpv4PingRequest
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIpv4PingRequest
	// validate validates ActionProtocolIpv4PingRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIpv4PingRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SrcName returns string, set in ActionProtocolIpv4PingRequest.
	SrcName() string
	// SetSrcName assigns string provided by user to ActionProtocolIpv4PingRequest
	SetSrcName(value string) ActionProtocolIpv4PingRequest
	// HasSrcName checks if SrcName has been set in ActionProtocolIpv4PingRequest
	HasSrcName() bool
	// DstIp returns string, set in ActionProtocolIpv4PingRequest.
	DstIp() string
	// SetDstIp assigns string provided by user to ActionProtocolIpv4PingRequest
	SetDstIp(value string) ActionProtocolIpv4PingRequest
	// HasDstIp checks if DstIp has been set in ActionProtocolIpv4PingRequest
	HasDstIp() bool
}

// Name of source IPv4 interface to be used.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// SrcName returns a string
func (obj *actionProtocolIpv4PingRequest) SrcName() string {

	return *obj.obj.SrcName

}

// Name of source IPv4 interface to be used.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// SrcName returns a string
func (obj *actionProtocolIpv4PingRequest) HasSrcName() bool {
	return obj.obj.SrcName != nil
}

// Name of source IPv4 interface to be used.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// SetSrcName sets the string value in the ActionProtocolIpv4PingRequest object
func (obj *actionProtocolIpv4PingRequest) SetSrcName(value string) ActionProtocolIpv4PingRequest {

	obj.obj.SrcName = &value
	return obj
}

// Destination IPv4 address to ping.
// DstIp returns a string
func (obj *actionProtocolIpv4PingRequest) DstIp() string {

	return *obj.obj.DstIp

}

// Destination IPv4 address to ping.
// DstIp returns a string
func (obj *actionProtocolIpv4PingRequest) HasDstIp() bool {
	return obj.obj.DstIp != nil
}

// Destination IPv4 address to ping.
// SetDstIp sets the string value in the ActionProtocolIpv4PingRequest object
func (obj *actionProtocolIpv4PingRequest) SetDstIp(value string) ActionProtocolIpv4PingRequest {

	obj.obj.DstIp = &value
	return obj
}

func (obj *actionProtocolIpv4PingRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.DstIp != nil {

		err := obj.validateIpv4(obj.DstIp())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on ActionProtocolIpv4PingRequest.DstIp"))
		}

	}

}

func (obj *actionProtocolIpv4PingRequest) setDefault() {

}
