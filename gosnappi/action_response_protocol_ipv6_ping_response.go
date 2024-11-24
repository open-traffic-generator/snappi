package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionResponseProtocolIpv6PingResponse *****
type actionResponseProtocolIpv6PingResponse struct {
	validation
	obj          *otg.ActionResponseProtocolIpv6PingResponse
	marshaller   marshalActionResponseProtocolIpv6PingResponse
	unMarshaller unMarshalActionResponseProtocolIpv6PingResponse
}

func NewActionResponseProtocolIpv6PingResponse() ActionResponseProtocolIpv6PingResponse {
	obj := actionResponseProtocolIpv6PingResponse{obj: &otg.ActionResponseProtocolIpv6PingResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *actionResponseProtocolIpv6PingResponse) msg() *otg.ActionResponseProtocolIpv6PingResponse {
	return obj.obj
}

func (obj *actionResponseProtocolIpv6PingResponse) setMsg(msg *otg.ActionResponseProtocolIpv6PingResponse) ActionResponseProtocolIpv6PingResponse {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionResponseProtocolIpv6PingResponse struct {
	obj *actionResponseProtocolIpv6PingResponse
}

type marshalActionResponseProtocolIpv6PingResponse interface {
	// ToProto marshals ActionResponseProtocolIpv6PingResponse to protobuf object *otg.ActionResponseProtocolIpv6PingResponse
	ToProto() (*otg.ActionResponseProtocolIpv6PingResponse, error)
	// ToPbText marshals ActionResponseProtocolIpv6PingResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionResponseProtocolIpv6PingResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionResponseProtocolIpv6PingResponse to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ActionResponseProtocolIpv6PingResponse to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalactionResponseProtocolIpv6PingResponse struct {
	obj *actionResponseProtocolIpv6PingResponse
}

type unMarshalActionResponseProtocolIpv6PingResponse interface {
	// FromProto unmarshals ActionResponseProtocolIpv6PingResponse from protobuf object *otg.ActionResponseProtocolIpv6PingResponse
	FromProto(msg *otg.ActionResponseProtocolIpv6PingResponse) (ActionResponseProtocolIpv6PingResponse, error)
	// FromPbText unmarshals ActionResponseProtocolIpv6PingResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionResponseProtocolIpv6PingResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionResponseProtocolIpv6PingResponse from JSON text
	FromJson(value string) error
}

func (obj *actionResponseProtocolIpv6PingResponse) Marshal() marshalActionResponseProtocolIpv6PingResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionResponseProtocolIpv6PingResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionResponseProtocolIpv6PingResponse) Unmarshal() unMarshalActionResponseProtocolIpv6PingResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionResponseProtocolIpv6PingResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionResponseProtocolIpv6PingResponse) ToProto() (*otg.ActionResponseProtocolIpv6PingResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionResponseProtocolIpv6PingResponse) FromProto(msg *otg.ActionResponseProtocolIpv6PingResponse) (ActionResponseProtocolIpv6PingResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionResponseProtocolIpv6PingResponse) ToPbText() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv6PingResponse) FromPbText(value string) error {
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

func (m *marshalactionResponseProtocolIpv6PingResponse) ToYaml() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv6PingResponse) FromYaml(value string) error {
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

func (m *marshalactionResponseProtocolIpv6PingResponse) ToJsonRaw() (string, error) {
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

func (m *marshalactionResponseProtocolIpv6PingResponse) ToJson() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv6PingResponse) FromJson(value string) error {
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

func (obj *actionResponseProtocolIpv6PingResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionResponseProtocolIpv6PingResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionResponseProtocolIpv6PingResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionResponseProtocolIpv6PingResponse) Clone() (ActionResponseProtocolIpv6PingResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionResponseProtocolIpv6PingResponse()
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

// ActionResponseProtocolIpv6PingResponse is response for ping initiated between a single source and destination pair.
type ActionResponseProtocolIpv6PingResponse interface {
	Validation
	// msg marshals ActionResponseProtocolIpv6PingResponse to protobuf object *otg.ActionResponseProtocolIpv6PingResponse
	// and doesn't set defaults
	msg() *otg.ActionResponseProtocolIpv6PingResponse
	// setMsg unmarshals ActionResponseProtocolIpv6PingResponse from protobuf object *otg.ActionResponseProtocolIpv6PingResponse
	// and doesn't set defaults
	setMsg(*otg.ActionResponseProtocolIpv6PingResponse) ActionResponseProtocolIpv6PingResponse
	// provides marshal interface
	Marshal() marshalActionResponseProtocolIpv6PingResponse
	// provides unmarshal interface
	Unmarshal() unMarshalActionResponseProtocolIpv6PingResponse
	// validate validates ActionResponseProtocolIpv6PingResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionResponseProtocolIpv6PingResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SrcName returns string, set in ActionResponseProtocolIpv6PingResponse.
	SrcName() string
	// SetSrcName assigns string provided by user to ActionResponseProtocolIpv6PingResponse
	SetSrcName(value string) ActionResponseProtocolIpv6PingResponse
	// DstIp returns string, set in ActionResponseProtocolIpv6PingResponse.
	DstIp() string
	// SetDstIp assigns string provided by user to ActionResponseProtocolIpv6PingResponse
	SetDstIp(value string) ActionResponseProtocolIpv6PingResponse
	// Result returns ActionResponseProtocolIpv6PingResponseResultEnum, set in ActionResponseProtocolIpv6PingResponse
	Result() ActionResponseProtocolIpv6PingResponseResultEnum
	// SetResult assigns ActionResponseProtocolIpv6PingResponseResultEnum provided by user to ActionResponseProtocolIpv6PingResponse
	SetResult(value ActionResponseProtocolIpv6PingResponseResultEnum) ActionResponseProtocolIpv6PingResponse
}

// Name of source IPv6 interface used for ping.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// SrcName returns a string
func (obj *actionResponseProtocolIpv6PingResponse) SrcName() string {

	return *obj.obj.SrcName

}

// Name of source IPv6 interface used for ping.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// SetSrcName sets the string value in the ActionResponseProtocolIpv6PingResponse object
func (obj *actionResponseProtocolIpv6PingResponse) SetSrcName(value string) ActionResponseProtocolIpv6PingResponse {

	obj.obj.SrcName = &value
	return obj
}

// Destination IPv6 address used for ping.
// DstIp returns a string
func (obj *actionResponseProtocolIpv6PingResponse) DstIp() string {

	return *obj.obj.DstIp

}

// Destination IPv6 address used for ping.
// SetDstIp sets the string value in the ActionResponseProtocolIpv6PingResponse object
func (obj *actionResponseProtocolIpv6PingResponse) SetDstIp(value string) ActionResponseProtocolIpv6PingResponse {

	obj.obj.DstIp = &value
	return obj
}

type ActionResponseProtocolIpv6PingResponseResultEnum string

// Enum of Result on ActionResponseProtocolIpv6PingResponse
var ActionResponseProtocolIpv6PingResponseResult = struct {
	SUCCEEDED ActionResponseProtocolIpv6PingResponseResultEnum
	FAILED    ActionResponseProtocolIpv6PingResponseResultEnum
}{
	SUCCEEDED: ActionResponseProtocolIpv6PingResponseResultEnum("succeeded"),
	FAILED:    ActionResponseProtocolIpv6PingResponseResultEnum("failed"),
}

func (obj *actionResponseProtocolIpv6PingResponse) Result() ActionResponseProtocolIpv6PingResponseResultEnum {
	return ActionResponseProtocolIpv6PingResponseResultEnum(obj.obj.Result.Enum().String())
}

func (obj *actionResponseProtocolIpv6PingResponse) SetResult(value ActionResponseProtocolIpv6PingResponseResultEnum) ActionResponseProtocolIpv6PingResponse {
	intValue, ok := otg.ActionResponseProtocolIpv6PingResponse_Result_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionResponseProtocolIpv6PingResponseResultEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionResponseProtocolIpv6PingResponse_Result_Enum(intValue)
	obj.obj.Result = &enumValue

	return obj
}

func (obj *actionResponseProtocolIpv6PingResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// SrcName is required
	if obj.obj.SrcName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "SrcName is required field on interface ActionResponseProtocolIpv6PingResponse")
	}

	// DstIp is required
	if obj.obj.DstIp == nil {
		vObj.validationErrors = append(vObj.validationErrors, "DstIp is required field on interface ActionResponseProtocolIpv6PingResponse")
	}
	if obj.obj.DstIp != nil {

		err := obj.validateIpv6(obj.DstIp())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on ActionResponseProtocolIpv6PingResponse.DstIp"))
		}

	}

	// Result is required
	if obj.obj.Result == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Result is required field on interface ActionResponseProtocolIpv6PingResponse")
	}
}

func (obj *actionResponseProtocolIpv6PingResponse) setDefault() {

}
