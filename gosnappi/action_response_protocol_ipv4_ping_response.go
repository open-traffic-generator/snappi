package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionResponseProtocolIpv4PingResponse *****
type actionResponseProtocolIpv4PingResponse struct {
	validation
	obj          *otg.ActionResponseProtocolIpv4PingResponse
	marshaller   marshalActionResponseProtocolIpv4PingResponse
	unMarshaller unMarshalActionResponseProtocolIpv4PingResponse
}

func NewActionResponseProtocolIpv4PingResponse() ActionResponseProtocolIpv4PingResponse {
	obj := actionResponseProtocolIpv4PingResponse{obj: &otg.ActionResponseProtocolIpv4PingResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *actionResponseProtocolIpv4PingResponse) msg() *otg.ActionResponseProtocolIpv4PingResponse {
	return obj.obj
}

func (obj *actionResponseProtocolIpv4PingResponse) setMsg(msg *otg.ActionResponseProtocolIpv4PingResponse) ActionResponseProtocolIpv4PingResponse {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionResponseProtocolIpv4PingResponse struct {
	obj *actionResponseProtocolIpv4PingResponse
}

type marshalActionResponseProtocolIpv4PingResponse interface {
	// ToProto marshals ActionResponseProtocolIpv4PingResponse to protobuf object *otg.ActionResponseProtocolIpv4PingResponse
	ToProto() (*otg.ActionResponseProtocolIpv4PingResponse, error)
	// ToPbText marshals ActionResponseProtocolIpv4PingResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionResponseProtocolIpv4PingResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionResponseProtocolIpv4PingResponse to JSON text
	ToJson() (string, error)
}

type unMarshalactionResponseProtocolIpv4PingResponse struct {
	obj *actionResponseProtocolIpv4PingResponse
}

type unMarshalActionResponseProtocolIpv4PingResponse interface {
	// FromProto unmarshals ActionResponseProtocolIpv4PingResponse from protobuf object *otg.ActionResponseProtocolIpv4PingResponse
	FromProto(msg *otg.ActionResponseProtocolIpv4PingResponse) (ActionResponseProtocolIpv4PingResponse, error)
	// FromPbText unmarshals ActionResponseProtocolIpv4PingResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionResponseProtocolIpv4PingResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionResponseProtocolIpv4PingResponse from JSON text
	FromJson(value string) error
}

func (obj *actionResponseProtocolIpv4PingResponse) Marshal() marshalActionResponseProtocolIpv4PingResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionResponseProtocolIpv4PingResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionResponseProtocolIpv4PingResponse) Unmarshal() unMarshalActionResponseProtocolIpv4PingResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionResponseProtocolIpv4PingResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionResponseProtocolIpv4PingResponse) ToProto() (*otg.ActionResponseProtocolIpv4PingResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionResponseProtocolIpv4PingResponse) FromProto(msg *otg.ActionResponseProtocolIpv4PingResponse) (ActionResponseProtocolIpv4PingResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionResponseProtocolIpv4PingResponse) ToPbText() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv4PingResponse) FromPbText(value string) error {
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

func (m *marshalactionResponseProtocolIpv4PingResponse) ToYaml() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv4PingResponse) FromYaml(value string) error {
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

func (m *marshalactionResponseProtocolIpv4PingResponse) ToJson() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv4PingResponse) FromJson(value string) error {
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

func (obj *actionResponseProtocolIpv4PingResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionResponseProtocolIpv4PingResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionResponseProtocolIpv4PingResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionResponseProtocolIpv4PingResponse) Clone() (ActionResponseProtocolIpv4PingResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionResponseProtocolIpv4PingResponse()
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

// ActionResponseProtocolIpv4PingResponse is response for ping initiated between a single source and destination pair.
type ActionResponseProtocolIpv4PingResponse interface {
	Validation
	// msg marshals ActionResponseProtocolIpv4PingResponse to protobuf object *otg.ActionResponseProtocolIpv4PingResponse
	// and doesn't set defaults
	msg() *otg.ActionResponseProtocolIpv4PingResponse
	// setMsg unmarshals ActionResponseProtocolIpv4PingResponse from protobuf object *otg.ActionResponseProtocolIpv4PingResponse
	// and doesn't set defaults
	setMsg(*otg.ActionResponseProtocolIpv4PingResponse) ActionResponseProtocolIpv4PingResponse
	// provides marshal interface
	Marshal() marshalActionResponseProtocolIpv4PingResponse
	// provides unmarshal interface
	Unmarshal() unMarshalActionResponseProtocolIpv4PingResponse
	// validate validates ActionResponseProtocolIpv4PingResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionResponseProtocolIpv4PingResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SrcName returns string, set in ActionResponseProtocolIpv4PingResponse.
	SrcName() string
	// SetSrcName assigns string provided by user to ActionResponseProtocolIpv4PingResponse
	SetSrcName(value string) ActionResponseProtocolIpv4PingResponse
	// DstIp returns string, set in ActionResponseProtocolIpv4PingResponse.
	DstIp() string
	// SetDstIp assigns string provided by user to ActionResponseProtocolIpv4PingResponse
	SetDstIp(value string) ActionResponseProtocolIpv4PingResponse
	// Result returns ActionResponseProtocolIpv4PingResponseResultEnum, set in ActionResponseProtocolIpv4PingResponse
	Result() ActionResponseProtocolIpv4PingResponseResultEnum
	// SetResult assigns ActionResponseProtocolIpv4PingResponseResultEnum provided by user to ActionResponseProtocolIpv4PingResponse
	SetResult(value ActionResponseProtocolIpv4PingResponseResultEnum) ActionResponseProtocolIpv4PingResponse
}

// Name of source IPv4 interface used for ping.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// SrcName returns a string
func (obj *actionResponseProtocolIpv4PingResponse) SrcName() string {

	return *obj.obj.SrcName

}

// Name of source IPv4 interface used for ping.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// SetSrcName sets the string value in the ActionResponseProtocolIpv4PingResponse object
func (obj *actionResponseProtocolIpv4PingResponse) SetSrcName(value string) ActionResponseProtocolIpv4PingResponse {

	obj.obj.SrcName = &value
	return obj
}

// Destination IPv4 address used for ping.
// DstIp returns a string
func (obj *actionResponseProtocolIpv4PingResponse) DstIp() string {

	return *obj.obj.DstIp

}

// Destination IPv4 address used for ping.
// SetDstIp sets the string value in the ActionResponseProtocolIpv4PingResponse object
func (obj *actionResponseProtocolIpv4PingResponse) SetDstIp(value string) ActionResponseProtocolIpv4PingResponse {

	obj.obj.DstIp = &value
	return obj
}

type ActionResponseProtocolIpv4PingResponseResultEnum string

// Enum of Result on ActionResponseProtocolIpv4PingResponse
var ActionResponseProtocolIpv4PingResponseResult = struct {
	SUCCEEDED ActionResponseProtocolIpv4PingResponseResultEnum
	FAILED    ActionResponseProtocolIpv4PingResponseResultEnum
}{
	SUCCEEDED: ActionResponseProtocolIpv4PingResponseResultEnum("succeeded"),
	FAILED:    ActionResponseProtocolIpv4PingResponseResultEnum("failed"),
}

func (obj *actionResponseProtocolIpv4PingResponse) Result() ActionResponseProtocolIpv4PingResponseResultEnum {
	return ActionResponseProtocolIpv4PingResponseResultEnum(obj.obj.Result.Enum().String())
}

func (obj *actionResponseProtocolIpv4PingResponse) SetResult(value ActionResponseProtocolIpv4PingResponseResultEnum) ActionResponseProtocolIpv4PingResponse {
	intValue, ok := otg.ActionResponseProtocolIpv4PingResponse_Result_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionResponseProtocolIpv4PingResponseResultEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionResponseProtocolIpv4PingResponse_Result_Enum(intValue)
	obj.obj.Result = &enumValue

	return obj
}

func (obj *actionResponseProtocolIpv4PingResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// SrcName is required
	if obj.obj.SrcName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "SrcName is required field on interface ActionResponseProtocolIpv4PingResponse")
	}

	// DstIp is required
	if obj.obj.DstIp == nil {
		vObj.validationErrors = append(vObj.validationErrors, "DstIp is required field on interface ActionResponseProtocolIpv4PingResponse")
	}
	if obj.obj.DstIp != nil {

		err := obj.validateIpv4(obj.DstIp())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on ActionResponseProtocolIpv4PingResponse.DstIp"))
		}

	}

	// Result is required
	if obj.obj.Result == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Result is required field on interface ActionResponseProtocolIpv4PingResponse")
	}
}

func (obj *actionResponseProtocolIpv4PingResponse) setDefault() {

}
