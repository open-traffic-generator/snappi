package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBgpMessageHeaderError *****
type deviceBgpMessageHeaderError struct {
	validation
	obj          *otg.DeviceBgpMessageHeaderError
	marshaller   marshalDeviceBgpMessageHeaderError
	unMarshaller unMarshalDeviceBgpMessageHeaderError
}

func NewDeviceBgpMessageHeaderError() DeviceBgpMessageHeaderError {
	obj := deviceBgpMessageHeaderError{obj: &otg.DeviceBgpMessageHeaderError{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBgpMessageHeaderError) msg() *otg.DeviceBgpMessageHeaderError {
	return obj.obj
}

func (obj *deviceBgpMessageHeaderError) setMsg(msg *otg.DeviceBgpMessageHeaderError) DeviceBgpMessageHeaderError {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBgpMessageHeaderError struct {
	obj *deviceBgpMessageHeaderError
}

type marshalDeviceBgpMessageHeaderError interface {
	// ToProto marshals DeviceBgpMessageHeaderError to protobuf object *otg.DeviceBgpMessageHeaderError
	ToProto() (*otg.DeviceBgpMessageHeaderError, error)
	// ToPbText marshals DeviceBgpMessageHeaderError to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBgpMessageHeaderError to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBgpMessageHeaderError to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBgpMessageHeaderError struct {
	obj *deviceBgpMessageHeaderError
}

type unMarshalDeviceBgpMessageHeaderError interface {
	// FromProto unmarshals DeviceBgpMessageHeaderError from protobuf object *otg.DeviceBgpMessageHeaderError
	FromProto(msg *otg.DeviceBgpMessageHeaderError) (DeviceBgpMessageHeaderError, error)
	// FromPbText unmarshals DeviceBgpMessageHeaderError from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBgpMessageHeaderError from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBgpMessageHeaderError from JSON text
	FromJson(value string) error
}

func (obj *deviceBgpMessageHeaderError) Marshal() marshalDeviceBgpMessageHeaderError {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBgpMessageHeaderError{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBgpMessageHeaderError) Unmarshal() unMarshalDeviceBgpMessageHeaderError {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBgpMessageHeaderError{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBgpMessageHeaderError) ToProto() (*otg.DeviceBgpMessageHeaderError, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBgpMessageHeaderError) FromProto(msg *otg.DeviceBgpMessageHeaderError) (DeviceBgpMessageHeaderError, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBgpMessageHeaderError) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBgpMessageHeaderError) FromPbText(value string) error {
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

func (m *marshaldeviceBgpMessageHeaderError) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBgpMessageHeaderError) FromYaml(value string) error {
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

func (m *marshaldeviceBgpMessageHeaderError) ToJson() (string, error) {
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

func (m *unMarshaldeviceBgpMessageHeaderError) FromJson(value string) error {
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

func (obj *deviceBgpMessageHeaderError) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBgpMessageHeaderError) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBgpMessageHeaderError) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBgpMessageHeaderError) Clone() (DeviceBgpMessageHeaderError, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBgpMessageHeaderError()
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

// DeviceBgpMessageHeaderError is all errors detected while processing the Message Header are indicated by sending the NOTIFICATION message  with the Error Code-Message Header Error. The Error Subcode elaborates on the specific nature of the error.
type DeviceBgpMessageHeaderError interface {
	Validation
	// msg marshals DeviceBgpMessageHeaderError to protobuf object *otg.DeviceBgpMessageHeaderError
	// and doesn't set defaults
	msg() *otg.DeviceBgpMessageHeaderError
	// setMsg unmarshals DeviceBgpMessageHeaderError from protobuf object *otg.DeviceBgpMessageHeaderError
	// and doesn't set defaults
	setMsg(*otg.DeviceBgpMessageHeaderError) DeviceBgpMessageHeaderError
	// provides marshal interface
	Marshal() marshalDeviceBgpMessageHeaderError
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBgpMessageHeaderError
	// validate validates DeviceBgpMessageHeaderError
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBgpMessageHeaderError, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Subcode returns DeviceBgpMessageHeaderErrorSubcodeEnum, set in DeviceBgpMessageHeaderError
	Subcode() DeviceBgpMessageHeaderErrorSubcodeEnum
	// SetSubcode assigns DeviceBgpMessageHeaderErrorSubcodeEnum provided by user to DeviceBgpMessageHeaderError
	SetSubcode(value DeviceBgpMessageHeaderErrorSubcodeEnum) DeviceBgpMessageHeaderError
	// HasSubcode checks if Subcode has been set in DeviceBgpMessageHeaderError
	HasSubcode() bool
}

type DeviceBgpMessageHeaderErrorSubcodeEnum string

// Enum of Subcode on DeviceBgpMessageHeaderError
var DeviceBgpMessageHeaderErrorSubcode = struct {
	CONNECTION_NOT_SYNCHRONIZED_CODE1_SUBCODE1 DeviceBgpMessageHeaderErrorSubcodeEnum
	BAD_MESSAGE_LENGTH_CODE1_SUBCODE2          DeviceBgpMessageHeaderErrorSubcodeEnum
	BAD_MESSAGE_TYPE_CODE1_SUBCODE3            DeviceBgpMessageHeaderErrorSubcodeEnum
}{
	CONNECTION_NOT_SYNCHRONIZED_CODE1_SUBCODE1: DeviceBgpMessageHeaderErrorSubcodeEnum("connection_not_synchronized_code1_subcode1"),
	BAD_MESSAGE_LENGTH_CODE1_SUBCODE2:          DeviceBgpMessageHeaderErrorSubcodeEnum("bad_message_length_code1_subcode2"),
	BAD_MESSAGE_TYPE_CODE1_SUBCODE3:            DeviceBgpMessageHeaderErrorSubcodeEnum("bad_message_type_code1_subcode3"),
}

func (obj *deviceBgpMessageHeaderError) Subcode() DeviceBgpMessageHeaderErrorSubcodeEnum {
	return DeviceBgpMessageHeaderErrorSubcodeEnum(obj.obj.Subcode.Enum().String())
}

// The Error Subcode indicates the specific type of error encountered during Message Header processing.
// Subcode returns a string
func (obj *deviceBgpMessageHeaderError) HasSubcode() bool {
	return obj.obj.Subcode != nil
}

func (obj *deviceBgpMessageHeaderError) SetSubcode(value DeviceBgpMessageHeaderErrorSubcodeEnum) DeviceBgpMessageHeaderError {
	intValue, ok := otg.DeviceBgpMessageHeaderError_Subcode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceBgpMessageHeaderErrorSubcodeEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceBgpMessageHeaderError_Subcode_Enum(intValue)
	obj.obj.Subcode = &enumValue

	return obj
}

func (obj *deviceBgpMessageHeaderError) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *deviceBgpMessageHeaderError) setDefault() {
	if obj.obj.Subcode == nil {
		obj.SetSubcode(DeviceBgpMessageHeaderErrorSubcode.CONNECTION_NOT_SYNCHRONIZED_CODE1_SUBCODE1)

	}

}
