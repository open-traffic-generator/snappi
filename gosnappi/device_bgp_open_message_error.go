package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBgpOpenMessageError *****
type deviceBgpOpenMessageError struct {
	validation
	obj          *otg.DeviceBgpOpenMessageError
	marshaller   marshalDeviceBgpOpenMessageError
	unMarshaller unMarshalDeviceBgpOpenMessageError
}

func NewDeviceBgpOpenMessageError() DeviceBgpOpenMessageError {
	obj := deviceBgpOpenMessageError{obj: &otg.DeviceBgpOpenMessageError{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBgpOpenMessageError) msg() *otg.DeviceBgpOpenMessageError {
	return obj.obj
}

func (obj *deviceBgpOpenMessageError) setMsg(msg *otg.DeviceBgpOpenMessageError) DeviceBgpOpenMessageError {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBgpOpenMessageError struct {
	obj *deviceBgpOpenMessageError
}

type marshalDeviceBgpOpenMessageError interface {
	// ToProto marshals DeviceBgpOpenMessageError to protobuf object *otg.DeviceBgpOpenMessageError
	ToProto() (*otg.DeviceBgpOpenMessageError, error)
	// ToPbText marshals DeviceBgpOpenMessageError to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBgpOpenMessageError to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBgpOpenMessageError to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals DeviceBgpOpenMessageError to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldeviceBgpOpenMessageError struct {
	obj *deviceBgpOpenMessageError
}

type unMarshalDeviceBgpOpenMessageError interface {
	// FromProto unmarshals DeviceBgpOpenMessageError from protobuf object *otg.DeviceBgpOpenMessageError
	FromProto(msg *otg.DeviceBgpOpenMessageError) (DeviceBgpOpenMessageError, error)
	// FromPbText unmarshals DeviceBgpOpenMessageError from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBgpOpenMessageError from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBgpOpenMessageError from JSON text
	FromJson(value string) error
}

func (obj *deviceBgpOpenMessageError) Marshal() marshalDeviceBgpOpenMessageError {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBgpOpenMessageError{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBgpOpenMessageError) Unmarshal() unMarshalDeviceBgpOpenMessageError {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBgpOpenMessageError{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBgpOpenMessageError) ToProto() (*otg.DeviceBgpOpenMessageError, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBgpOpenMessageError) FromProto(msg *otg.DeviceBgpOpenMessageError) (DeviceBgpOpenMessageError, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBgpOpenMessageError) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBgpOpenMessageError) FromPbText(value string) error {
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

func (m *marshaldeviceBgpOpenMessageError) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBgpOpenMessageError) FromYaml(value string) error {
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

func (m *marshaldeviceBgpOpenMessageError) ToJsonRaw() (string, error) {
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

func (m *marshaldeviceBgpOpenMessageError) ToJson() (string, error) {
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

func (m *unMarshaldeviceBgpOpenMessageError) FromJson(value string) error {
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

func (obj *deviceBgpOpenMessageError) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBgpOpenMessageError) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBgpOpenMessageError) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBgpOpenMessageError) Clone() (DeviceBgpOpenMessageError, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBgpOpenMessageError()
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

// DeviceBgpOpenMessageError is all errors detected while processing the OPEN message are indicated by sending the NOTIFICATION message  with the Error Code-Open Message Error. The Error Subcode elaborates on the specific nature of the error.
type DeviceBgpOpenMessageError interface {
	Validation
	// msg marshals DeviceBgpOpenMessageError to protobuf object *otg.DeviceBgpOpenMessageError
	// and doesn't set defaults
	msg() *otg.DeviceBgpOpenMessageError
	// setMsg unmarshals DeviceBgpOpenMessageError from protobuf object *otg.DeviceBgpOpenMessageError
	// and doesn't set defaults
	setMsg(*otg.DeviceBgpOpenMessageError) DeviceBgpOpenMessageError
	// provides marshal interface
	Marshal() marshalDeviceBgpOpenMessageError
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBgpOpenMessageError
	// validate validates DeviceBgpOpenMessageError
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBgpOpenMessageError, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Subcode returns DeviceBgpOpenMessageErrorSubcodeEnum, set in DeviceBgpOpenMessageError
	Subcode() DeviceBgpOpenMessageErrorSubcodeEnum
	// SetSubcode assigns DeviceBgpOpenMessageErrorSubcodeEnum provided by user to DeviceBgpOpenMessageError
	SetSubcode(value DeviceBgpOpenMessageErrorSubcodeEnum) DeviceBgpOpenMessageError
	// HasSubcode checks if Subcode has been set in DeviceBgpOpenMessageError
	HasSubcode() bool
}

type DeviceBgpOpenMessageErrorSubcodeEnum string

// Enum of Subcode on DeviceBgpOpenMessageError
var DeviceBgpOpenMessageErrorSubcode = struct {
	UNSUPPORTED_VERSION_NUMBER_CODE2_SUBCODE1     DeviceBgpOpenMessageErrorSubcodeEnum
	ERROR_PEER_AS_CODE2_SUBCODE2                  DeviceBgpOpenMessageErrorSubcodeEnum
	ERROR_BGP_ID_CODE2_SUBCODE3                   DeviceBgpOpenMessageErrorSubcodeEnum
	UNSUPPORTED_OPTIONAL_PARAMETER_CODE2_SUBCODE4 DeviceBgpOpenMessageErrorSubcodeEnum
	AUTH_FAILED_CODE2_SUBCODE5                    DeviceBgpOpenMessageErrorSubcodeEnum
	UNSUPPORTED_HOLD_TIME_CODE2_SUBCODE6          DeviceBgpOpenMessageErrorSubcodeEnum
	UNSUPPORTED_CAPABILITY_CODE2_SUBCODE7         DeviceBgpOpenMessageErrorSubcodeEnum
}{
	UNSUPPORTED_VERSION_NUMBER_CODE2_SUBCODE1:     DeviceBgpOpenMessageErrorSubcodeEnum("unsupported_version_number_code2_subcode1"),
	ERROR_PEER_AS_CODE2_SUBCODE2:                  DeviceBgpOpenMessageErrorSubcodeEnum("error_peer_as_code2_subcode2"),
	ERROR_BGP_ID_CODE2_SUBCODE3:                   DeviceBgpOpenMessageErrorSubcodeEnum("error_bgp_id_code2_subcode3"),
	UNSUPPORTED_OPTIONAL_PARAMETER_CODE2_SUBCODE4: DeviceBgpOpenMessageErrorSubcodeEnum("unsupported_optional_parameter_code2_subcode4"),
	AUTH_FAILED_CODE2_SUBCODE5:                    DeviceBgpOpenMessageErrorSubcodeEnum("auth_failed_code2_subcode5"),
	UNSUPPORTED_HOLD_TIME_CODE2_SUBCODE6:          DeviceBgpOpenMessageErrorSubcodeEnum("unsupported_hold_time_code2_subcode6"),
	UNSUPPORTED_CAPABILITY_CODE2_SUBCODE7:         DeviceBgpOpenMessageErrorSubcodeEnum("unsupported_capability_code2_subcode7"),
}

func (obj *deviceBgpOpenMessageError) Subcode() DeviceBgpOpenMessageErrorSubcodeEnum {
	return DeviceBgpOpenMessageErrorSubcodeEnum(obj.obj.Subcode.Enum().String())
}

// The Error Subcode indicates the specific type of error encountered during OPEN message processing.
// Subcode returns a string
func (obj *deviceBgpOpenMessageError) HasSubcode() bool {
	return obj.obj.Subcode != nil
}

func (obj *deviceBgpOpenMessageError) SetSubcode(value DeviceBgpOpenMessageErrorSubcodeEnum) DeviceBgpOpenMessageError {
	intValue, ok := otg.DeviceBgpOpenMessageError_Subcode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceBgpOpenMessageErrorSubcodeEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceBgpOpenMessageError_Subcode_Enum(intValue)
	obj.obj.Subcode = &enumValue

	return obj
}

func (obj *deviceBgpOpenMessageError) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *deviceBgpOpenMessageError) setDefault() {
	if obj.obj.Subcode == nil {
		obj.SetSubcode(DeviceBgpOpenMessageErrorSubcode.UNSUPPORTED_VERSION_NUMBER_CODE2_SUBCODE1)

	}

}
