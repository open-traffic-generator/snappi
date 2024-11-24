package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIpv6PingRequest *****
type actionProtocolIpv6PingRequest struct {
	validation
	obj          *otg.ActionProtocolIpv6PingRequest
	marshaller   marshalActionProtocolIpv6PingRequest
	unMarshaller unMarshalActionProtocolIpv6PingRequest
}

func NewActionProtocolIpv6PingRequest() ActionProtocolIpv6PingRequest {
	obj := actionProtocolIpv6PingRequest{obj: &otg.ActionProtocolIpv6PingRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIpv6PingRequest) msg() *otg.ActionProtocolIpv6PingRequest {
	return obj.obj
}

func (obj *actionProtocolIpv6PingRequest) setMsg(msg *otg.ActionProtocolIpv6PingRequest) ActionProtocolIpv6PingRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIpv6PingRequest struct {
	obj *actionProtocolIpv6PingRequest
}

type marshalActionProtocolIpv6PingRequest interface {
	// ToProto marshals ActionProtocolIpv6PingRequest to protobuf object *otg.ActionProtocolIpv6PingRequest
	ToProto() (*otg.ActionProtocolIpv6PingRequest, error)
	// ToPbText marshals ActionProtocolIpv6PingRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIpv6PingRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIpv6PingRequest to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIpv6PingRequest struct {
	obj *actionProtocolIpv6PingRequest
}

type unMarshalActionProtocolIpv6PingRequest interface {
	// FromProto unmarshals ActionProtocolIpv6PingRequest from protobuf object *otg.ActionProtocolIpv6PingRequest
	FromProto(msg *otg.ActionProtocolIpv6PingRequest) (ActionProtocolIpv6PingRequest, error)
	// FromPbText unmarshals ActionProtocolIpv6PingRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIpv6PingRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIpv6PingRequest from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIpv6PingRequest) Marshal() marshalActionProtocolIpv6PingRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIpv6PingRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIpv6PingRequest) Unmarshal() unMarshalActionProtocolIpv6PingRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIpv6PingRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIpv6PingRequest) ToProto() (*otg.ActionProtocolIpv6PingRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIpv6PingRequest) FromProto(msg *otg.ActionProtocolIpv6PingRequest) (ActionProtocolIpv6PingRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIpv6PingRequest) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIpv6PingRequest) FromPbText(value string) error {
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

func (m *marshalactionProtocolIpv6PingRequest) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIpv6PingRequest) FromYaml(value string) error {
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

func (m *marshalactionProtocolIpv6PingRequest) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIpv6PingRequest) FromJson(value string) error {
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

func (obj *actionProtocolIpv6PingRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIpv6PingRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIpv6PingRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIpv6PingRequest) Clone() (ActionProtocolIpv6PingRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIpv6PingRequest()
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

// ActionProtocolIpv6PingRequest is under Review: Most ping request parameters are still TBD.
//
// Under Review: Most ping request parameters are still TBD.
//
// Request for initiating ping between a single source and destination pair.
// For ping request, 1 IPv6 ICMP Echo Request shall be sent and wait for ping response to either succeed or time out. The API wait timeout for each request shall be 300ms.
type ActionProtocolIpv6PingRequest interface {
	Validation
	// msg marshals ActionProtocolIpv6PingRequest to protobuf object *otg.ActionProtocolIpv6PingRequest
	// and doesn't set defaults
	msg() *otg.ActionProtocolIpv6PingRequest
	// setMsg unmarshals ActionProtocolIpv6PingRequest from protobuf object *otg.ActionProtocolIpv6PingRequest
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIpv6PingRequest) ActionProtocolIpv6PingRequest
	// provides marshal interface
	Marshal() marshalActionProtocolIpv6PingRequest
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIpv6PingRequest
	// validate validates ActionProtocolIpv6PingRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIpv6PingRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SrcName returns string, set in ActionProtocolIpv6PingRequest.
	SrcName() string
	// SetSrcName assigns string provided by user to ActionProtocolIpv6PingRequest
	SetSrcName(value string) ActionProtocolIpv6PingRequest
	// HasSrcName checks if SrcName has been set in ActionProtocolIpv6PingRequest
	HasSrcName() bool
	// DstIp returns string, set in ActionProtocolIpv6PingRequest.
	DstIp() string
	// SetDstIp assigns string provided by user to ActionProtocolIpv6PingRequest
	SetDstIp(value string) ActionProtocolIpv6PingRequest
	// HasDstIp checks if DstIp has been set in ActionProtocolIpv6PingRequest
	HasDstIp() bool
}

// Name of source IPv6 interface to be used.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// SrcName returns a string
func (obj *actionProtocolIpv6PingRequest) SrcName() string {

	return *obj.obj.SrcName

}

// Name of source IPv6 interface to be used.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// SrcName returns a string
func (obj *actionProtocolIpv6PingRequest) HasSrcName() bool {
	return obj.obj.SrcName != nil
}

// Name of source IPv6 interface to be used.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// SetSrcName sets the string value in the ActionProtocolIpv6PingRequest object
func (obj *actionProtocolIpv6PingRequest) SetSrcName(value string) ActionProtocolIpv6PingRequest {

	obj.obj.SrcName = &value
	return obj
}

// Destination IPv6 address to ping.
// DstIp returns a string
func (obj *actionProtocolIpv6PingRequest) DstIp() string {

	return *obj.obj.DstIp

}

// Destination IPv6 address to ping.
// DstIp returns a string
func (obj *actionProtocolIpv6PingRequest) HasDstIp() bool {
	return obj.obj.DstIp != nil
}

// Destination IPv6 address to ping.
// SetDstIp sets the string value in the ActionProtocolIpv6PingRequest object
func (obj *actionProtocolIpv6PingRequest) SetDstIp(value string) ActionProtocolIpv6PingRequest {

	obj.obj.DstIp = &value
	return obj
}

func (obj *actionProtocolIpv6PingRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	obj.addWarnings("ActionProtocolIpv6PingRequest is under review, Most ping request parameters are still TBD.")

	if obj.obj.DstIp != nil {

		err := obj.validateIpv6(obj.DstIp())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on ActionProtocolIpv6PingRequest.DstIp"))
		}

	}

}

func (obj *actionProtocolIpv6PingRequest) setDefault() {

}
