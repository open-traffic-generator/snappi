package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBgpUpdateMessageError *****
type deviceBgpUpdateMessageError struct {
	validation
	obj          *otg.DeviceBgpUpdateMessageError
	marshaller   marshalDeviceBgpUpdateMessageError
	unMarshaller unMarshalDeviceBgpUpdateMessageError
}

func NewDeviceBgpUpdateMessageError() DeviceBgpUpdateMessageError {
	obj := deviceBgpUpdateMessageError{obj: &otg.DeviceBgpUpdateMessageError{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBgpUpdateMessageError) msg() *otg.DeviceBgpUpdateMessageError {
	return obj.obj
}

func (obj *deviceBgpUpdateMessageError) setMsg(msg *otg.DeviceBgpUpdateMessageError) DeviceBgpUpdateMessageError {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBgpUpdateMessageError struct {
	obj *deviceBgpUpdateMessageError
}

type marshalDeviceBgpUpdateMessageError interface {
	// ToProto marshals DeviceBgpUpdateMessageError to protobuf object *otg.DeviceBgpUpdateMessageError
	ToProto() (*otg.DeviceBgpUpdateMessageError, error)
	// ToPbText marshals DeviceBgpUpdateMessageError to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBgpUpdateMessageError to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBgpUpdateMessageError to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBgpUpdateMessageError struct {
	obj *deviceBgpUpdateMessageError
}

type unMarshalDeviceBgpUpdateMessageError interface {
	// FromProto unmarshals DeviceBgpUpdateMessageError from protobuf object *otg.DeviceBgpUpdateMessageError
	FromProto(msg *otg.DeviceBgpUpdateMessageError) (DeviceBgpUpdateMessageError, error)
	// FromPbText unmarshals DeviceBgpUpdateMessageError from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBgpUpdateMessageError from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBgpUpdateMessageError from JSON text
	FromJson(value string) error
}

func (obj *deviceBgpUpdateMessageError) Marshal() marshalDeviceBgpUpdateMessageError {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBgpUpdateMessageError{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBgpUpdateMessageError) Unmarshal() unMarshalDeviceBgpUpdateMessageError {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBgpUpdateMessageError{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBgpUpdateMessageError) ToProto() (*otg.DeviceBgpUpdateMessageError, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBgpUpdateMessageError) FromProto(msg *otg.DeviceBgpUpdateMessageError) (DeviceBgpUpdateMessageError, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBgpUpdateMessageError) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBgpUpdateMessageError) FromPbText(value string) error {
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

func (m *marshaldeviceBgpUpdateMessageError) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBgpUpdateMessageError) FromYaml(value string) error {
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

func (m *marshaldeviceBgpUpdateMessageError) ToJson() (string, error) {
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

func (m *unMarshaldeviceBgpUpdateMessageError) FromJson(value string) error {
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

func (obj *deviceBgpUpdateMessageError) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBgpUpdateMessageError) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBgpUpdateMessageError) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBgpUpdateMessageError) Clone() (DeviceBgpUpdateMessageError, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBgpUpdateMessageError()
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

// DeviceBgpUpdateMessageError is all errors detected while processing the UPDATE message are indicated by sending the NOTIFICATION message  with the Error Code-Update Message Error. The Error Subcode elaborates on the specific nature of the error.
type DeviceBgpUpdateMessageError interface {
	Validation
	// msg marshals DeviceBgpUpdateMessageError to protobuf object *otg.DeviceBgpUpdateMessageError
	// and doesn't set defaults
	msg() *otg.DeviceBgpUpdateMessageError
	// setMsg unmarshals DeviceBgpUpdateMessageError from protobuf object *otg.DeviceBgpUpdateMessageError
	// and doesn't set defaults
	setMsg(*otg.DeviceBgpUpdateMessageError) DeviceBgpUpdateMessageError
	// provides marshal interface
	Marshal() marshalDeviceBgpUpdateMessageError
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBgpUpdateMessageError
	// validate validates DeviceBgpUpdateMessageError
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBgpUpdateMessageError, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Subcode returns DeviceBgpUpdateMessageErrorSubcodeEnum, set in DeviceBgpUpdateMessageError
	Subcode() DeviceBgpUpdateMessageErrorSubcodeEnum
	// SetSubcode assigns DeviceBgpUpdateMessageErrorSubcodeEnum provided by user to DeviceBgpUpdateMessageError
	SetSubcode(value DeviceBgpUpdateMessageErrorSubcodeEnum) DeviceBgpUpdateMessageError
	// HasSubcode checks if Subcode has been set in DeviceBgpUpdateMessageError
	HasSubcode() bool
}

type DeviceBgpUpdateMessageErrorSubcodeEnum string

// Enum of Subcode on DeviceBgpUpdateMessageError
var DeviceBgpUpdateMessageErrorSubcode = struct {
	MALFORMED_ATTRIB_LIST_CODE3_SUBCODE1         DeviceBgpUpdateMessageErrorSubcodeEnum
	UNRECOGNIZED_WELLKNOWN_ATTRIB_CODE3_SUBCODE2 DeviceBgpUpdateMessageErrorSubcodeEnum
	WELLKNOWN_ATTRIB_MISSING_CODE3_SUBCODE3      DeviceBgpUpdateMessageErrorSubcodeEnum
	ATTRIB_FLAGS_ERROR_CODE3_SUBCODE4            DeviceBgpUpdateMessageErrorSubcodeEnum
	ATTRIB_LENGTH_ERROR_CODE3_SUBCODE5           DeviceBgpUpdateMessageErrorSubcodeEnum
	INVALID_ORIGIN_ATTRIB_CODE3_SUBCODE6         DeviceBgpUpdateMessageErrorSubcodeEnum
	AS_ROUTING_LOOP_CODE3_SUBCODE7               DeviceBgpUpdateMessageErrorSubcodeEnum
	INVALID_NHOP_ATTRIB_CODE3_SUBCODE8           DeviceBgpUpdateMessageErrorSubcodeEnum
	ERROR_OPTIONAL_ATTRIB_CODE3_SUBCODE9         DeviceBgpUpdateMessageErrorSubcodeEnum
	INVALID_NETWORK_FIELD_CODE3_SUBCODE10        DeviceBgpUpdateMessageErrorSubcodeEnum
	ABNORMAL_ASPATH_CODE3_SUBCODE11              DeviceBgpUpdateMessageErrorSubcodeEnum
}{
	MALFORMED_ATTRIB_LIST_CODE3_SUBCODE1:         DeviceBgpUpdateMessageErrorSubcodeEnum("malformed_attrib_list_code3_subcode1"),
	UNRECOGNIZED_WELLKNOWN_ATTRIB_CODE3_SUBCODE2: DeviceBgpUpdateMessageErrorSubcodeEnum("unrecognized_wellknown_attrib_code3_subcode2"),
	WELLKNOWN_ATTRIB_MISSING_CODE3_SUBCODE3:      DeviceBgpUpdateMessageErrorSubcodeEnum("wellknown_attrib_missing_code3_subcode3"),
	ATTRIB_FLAGS_ERROR_CODE3_SUBCODE4:            DeviceBgpUpdateMessageErrorSubcodeEnum("attrib_flags_error_code3_subcode4"),
	ATTRIB_LENGTH_ERROR_CODE3_SUBCODE5:           DeviceBgpUpdateMessageErrorSubcodeEnum("attrib_length_error_code3_subcode5"),
	INVALID_ORIGIN_ATTRIB_CODE3_SUBCODE6:         DeviceBgpUpdateMessageErrorSubcodeEnum("invalid_origin_attrib_code3_subcode6"),
	AS_ROUTING_LOOP_CODE3_SUBCODE7:               DeviceBgpUpdateMessageErrorSubcodeEnum("as_routing_loop_code3_subcode7"),
	INVALID_NHOP_ATTRIB_CODE3_SUBCODE8:           DeviceBgpUpdateMessageErrorSubcodeEnum("invalid_nhop_attrib_code3_subcode8"),
	ERROR_OPTIONAL_ATTRIB_CODE3_SUBCODE9:         DeviceBgpUpdateMessageErrorSubcodeEnum("error_optional_attrib_code3_subcode9"),
	INVALID_NETWORK_FIELD_CODE3_SUBCODE10:        DeviceBgpUpdateMessageErrorSubcodeEnum("invalid_network_field_code3_subcode10"),
	ABNORMAL_ASPATH_CODE3_SUBCODE11:              DeviceBgpUpdateMessageErrorSubcodeEnum("abnormal_aspath_code3_subcode11"),
}

func (obj *deviceBgpUpdateMessageError) Subcode() DeviceBgpUpdateMessageErrorSubcodeEnum {
	return DeviceBgpUpdateMessageErrorSubcodeEnum(obj.obj.Subcode.Enum().String())
}

// The Error Subcode, the specific type of error encountered during UPDATE processing.
// Subcode returns a string
func (obj *deviceBgpUpdateMessageError) HasSubcode() bool {
	return obj.obj.Subcode != nil
}

func (obj *deviceBgpUpdateMessageError) SetSubcode(value DeviceBgpUpdateMessageErrorSubcodeEnum) DeviceBgpUpdateMessageError {
	intValue, ok := otg.DeviceBgpUpdateMessageError_Subcode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceBgpUpdateMessageErrorSubcodeEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceBgpUpdateMessageError_Subcode_Enum(intValue)
	obj.obj.Subcode = &enumValue

	return obj
}

func (obj *deviceBgpUpdateMessageError) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *deviceBgpUpdateMessageError) setDefault() {
	if obj.obj.Subcode == nil {
		obj.SetSubcode(DeviceBgpUpdateMessageErrorSubcode.MALFORMED_ATTRIB_LIST_CODE3_SUBCODE1)

	}

}
