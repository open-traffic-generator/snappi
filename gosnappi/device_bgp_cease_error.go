package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBgpCeaseError *****
type deviceBgpCeaseError struct {
	validation
	obj          *otg.DeviceBgpCeaseError
	marshaller   marshalDeviceBgpCeaseError
	unMarshaller unMarshalDeviceBgpCeaseError
}

func NewDeviceBgpCeaseError() DeviceBgpCeaseError {
	obj := deviceBgpCeaseError{obj: &otg.DeviceBgpCeaseError{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBgpCeaseError) msg() *otg.DeviceBgpCeaseError {
	return obj.obj
}

func (obj *deviceBgpCeaseError) setMsg(msg *otg.DeviceBgpCeaseError) DeviceBgpCeaseError {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBgpCeaseError struct {
	obj *deviceBgpCeaseError
}

type marshalDeviceBgpCeaseError interface {
	// ToProto marshals DeviceBgpCeaseError to protobuf object *otg.DeviceBgpCeaseError
	ToProto() (*otg.DeviceBgpCeaseError, error)
	// ToPbText marshals DeviceBgpCeaseError to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBgpCeaseError to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBgpCeaseError to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals DeviceBgpCeaseError to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldeviceBgpCeaseError struct {
	obj *deviceBgpCeaseError
}

type unMarshalDeviceBgpCeaseError interface {
	// FromProto unmarshals DeviceBgpCeaseError from protobuf object *otg.DeviceBgpCeaseError
	FromProto(msg *otg.DeviceBgpCeaseError) (DeviceBgpCeaseError, error)
	// FromPbText unmarshals DeviceBgpCeaseError from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBgpCeaseError from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBgpCeaseError from JSON text
	FromJson(value string) error
}

func (obj *deviceBgpCeaseError) Marshal() marshalDeviceBgpCeaseError {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBgpCeaseError{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBgpCeaseError) Unmarshal() unMarshalDeviceBgpCeaseError {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBgpCeaseError{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBgpCeaseError) ToProto() (*otg.DeviceBgpCeaseError, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBgpCeaseError) FromProto(msg *otg.DeviceBgpCeaseError) (DeviceBgpCeaseError, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBgpCeaseError) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBgpCeaseError) FromPbText(value string) error {
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

func (m *marshaldeviceBgpCeaseError) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBgpCeaseError) FromYaml(value string) error {
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

func (m *marshaldeviceBgpCeaseError) ToJsonRaw() (string, error) {
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

func (m *marshaldeviceBgpCeaseError) ToJson() (string, error) {
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

func (m *unMarshaldeviceBgpCeaseError) FromJson(value string) error {
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

func (obj *deviceBgpCeaseError) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBgpCeaseError) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBgpCeaseError) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBgpCeaseError) Clone() (DeviceBgpCeaseError, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBgpCeaseError()
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

// DeviceBgpCeaseError is in the absence of any fatal errors, a BGP peer can close its BGP connection by sending the NOTIFICATION message with the  Error Code Cease. The 'hard_reset_code6_subcode9' subcode for Cease Notification can be used to signal a hard reset that will indicate that  Graceful Restart cannot be performed, even when Notification extensions to Graceful Restart procedure is supported.
type DeviceBgpCeaseError interface {
	Validation
	// msg marshals DeviceBgpCeaseError to protobuf object *otg.DeviceBgpCeaseError
	// and doesn't set defaults
	msg() *otg.DeviceBgpCeaseError
	// setMsg unmarshals DeviceBgpCeaseError from protobuf object *otg.DeviceBgpCeaseError
	// and doesn't set defaults
	setMsg(*otg.DeviceBgpCeaseError) DeviceBgpCeaseError
	// provides marshal interface
	Marshal() marshalDeviceBgpCeaseError
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBgpCeaseError
	// validate validates DeviceBgpCeaseError
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBgpCeaseError, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Subcode returns DeviceBgpCeaseErrorSubcodeEnum, set in DeviceBgpCeaseError
	Subcode() DeviceBgpCeaseErrorSubcodeEnum
	// SetSubcode assigns DeviceBgpCeaseErrorSubcodeEnum provided by user to DeviceBgpCeaseError
	SetSubcode(value DeviceBgpCeaseErrorSubcodeEnum) DeviceBgpCeaseError
	// HasSubcode checks if Subcode has been set in DeviceBgpCeaseError
	HasSubcode() bool
}

type DeviceBgpCeaseErrorSubcodeEnum string

// Enum of Subcode on DeviceBgpCeaseError
var DeviceBgpCeaseErrorSubcode = struct {
	MAX_NUMBER_PREFIX_REACHED_CODE6_SUBCODE1       DeviceBgpCeaseErrorSubcodeEnum
	ADMIN_SHUTDOWN_CODE6_SUBCODE2                  DeviceBgpCeaseErrorSubcodeEnum
	PEER_DELETED_CODE6_SUBCODE3                    DeviceBgpCeaseErrorSubcodeEnum
	ADMIN_RESET_CODE6_SUBCODE4                     DeviceBgpCeaseErrorSubcodeEnum
	CONNECTION_REJECT_CODE6_SUBCODE5               DeviceBgpCeaseErrorSubcodeEnum
	OTHER_CONFIG_CHANGES_CODE6_SUBCODE6            DeviceBgpCeaseErrorSubcodeEnum
	CONNECTION_COLLISION_RESOLUTION_CODE6_SUBCODE7 DeviceBgpCeaseErrorSubcodeEnum
	OUT_OF_RESOURCES_CODE6_SUBCODE8                DeviceBgpCeaseErrorSubcodeEnum
	BFD_SESSION_DOWN_CODE6_SUBCODE10               DeviceBgpCeaseErrorSubcodeEnum
	HARD_RESET_CODE6_SUBCODE9                      DeviceBgpCeaseErrorSubcodeEnum
}{
	MAX_NUMBER_PREFIX_REACHED_CODE6_SUBCODE1:       DeviceBgpCeaseErrorSubcodeEnum("max_number_prefix_reached_code6_subcode1"),
	ADMIN_SHUTDOWN_CODE6_SUBCODE2:                  DeviceBgpCeaseErrorSubcodeEnum("admin_shutdown_code6_subcode2"),
	PEER_DELETED_CODE6_SUBCODE3:                    DeviceBgpCeaseErrorSubcodeEnum("peer_deleted_code6_subcode3"),
	ADMIN_RESET_CODE6_SUBCODE4:                     DeviceBgpCeaseErrorSubcodeEnum("admin_reset_code6_subcode4"),
	CONNECTION_REJECT_CODE6_SUBCODE5:               DeviceBgpCeaseErrorSubcodeEnum("connection_reject_code6_subcode5"),
	OTHER_CONFIG_CHANGES_CODE6_SUBCODE6:            DeviceBgpCeaseErrorSubcodeEnum("other_config_changes_code6_subcode6"),
	CONNECTION_COLLISION_RESOLUTION_CODE6_SUBCODE7: DeviceBgpCeaseErrorSubcodeEnum("connection_collision_resolution_code6_subcode7"),
	OUT_OF_RESOURCES_CODE6_SUBCODE8:                DeviceBgpCeaseErrorSubcodeEnum("out_of_resources_code6_subcode8"),
	BFD_SESSION_DOWN_CODE6_SUBCODE10:               DeviceBgpCeaseErrorSubcodeEnum("bfd_session_down_code6_subcode10"),
	HARD_RESET_CODE6_SUBCODE9:                      DeviceBgpCeaseErrorSubcodeEnum("hard_reset_code6_subcode9"),
}

func (obj *deviceBgpCeaseError) Subcode() DeviceBgpCeaseErrorSubcodeEnum {
	return DeviceBgpCeaseErrorSubcodeEnum(obj.obj.Subcode.Enum().String())
}

// The Error Subcode to be sent to the peer in the Cease NOTIFICATION.
// Subcode returns a string
func (obj *deviceBgpCeaseError) HasSubcode() bool {
	return obj.obj.Subcode != nil
}

func (obj *deviceBgpCeaseError) SetSubcode(value DeviceBgpCeaseErrorSubcodeEnum) DeviceBgpCeaseError {
	intValue, ok := otg.DeviceBgpCeaseError_Subcode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceBgpCeaseErrorSubcodeEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceBgpCeaseError_Subcode_Enum(intValue)
	obj.obj.Subcode = &enumValue

	return obj
}

func (obj *deviceBgpCeaseError) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *deviceBgpCeaseError) setDefault() {
	if obj.obj.Subcode == nil {
		obj.SetSubcode(DeviceBgpCeaseErrorSubcode.ADMIN_SHUTDOWN_CODE6_SUBCODE2)

	}

}
