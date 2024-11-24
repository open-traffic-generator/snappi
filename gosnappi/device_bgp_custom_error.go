package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBgpCustomError *****
type deviceBgpCustomError struct {
	validation
	obj          *otg.DeviceBgpCustomError
	marshaller   marshalDeviceBgpCustomError
	unMarshaller unMarshalDeviceBgpCustomError
}

func NewDeviceBgpCustomError() DeviceBgpCustomError {
	obj := deviceBgpCustomError{obj: &otg.DeviceBgpCustomError{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBgpCustomError) msg() *otg.DeviceBgpCustomError {
	return obj.obj
}

func (obj *deviceBgpCustomError) setMsg(msg *otg.DeviceBgpCustomError) DeviceBgpCustomError {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBgpCustomError struct {
	obj *deviceBgpCustomError
}

type marshalDeviceBgpCustomError interface {
	// ToProto marshals DeviceBgpCustomError to protobuf object *otg.DeviceBgpCustomError
	ToProto() (*otg.DeviceBgpCustomError, error)
	// ToPbText marshals DeviceBgpCustomError to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBgpCustomError to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBgpCustomError to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals DeviceBgpCustomError to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldeviceBgpCustomError struct {
	obj *deviceBgpCustomError
}

type unMarshalDeviceBgpCustomError interface {
	// FromProto unmarshals DeviceBgpCustomError from protobuf object *otg.DeviceBgpCustomError
	FromProto(msg *otg.DeviceBgpCustomError) (DeviceBgpCustomError, error)
	// FromPbText unmarshals DeviceBgpCustomError from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBgpCustomError from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBgpCustomError from JSON text
	FromJson(value string) error
}

func (obj *deviceBgpCustomError) Marshal() marshalDeviceBgpCustomError {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBgpCustomError{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBgpCustomError) Unmarshal() unMarshalDeviceBgpCustomError {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBgpCustomError{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBgpCustomError) ToProto() (*otg.DeviceBgpCustomError, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBgpCustomError) FromProto(msg *otg.DeviceBgpCustomError) (DeviceBgpCustomError, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBgpCustomError) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBgpCustomError) FromPbText(value string) error {
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

func (m *marshaldeviceBgpCustomError) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBgpCustomError) FromYaml(value string) error {
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

func (m *marshaldeviceBgpCustomError) ToJsonRaw() (string, error) {
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

func (m *marshaldeviceBgpCustomError) ToJson() (string, error) {
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

func (m *unMarshaldeviceBgpCustomError) FromJson(value string) error {
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

func (obj *deviceBgpCustomError) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBgpCustomError) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBgpCustomError) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBgpCustomError) Clone() (DeviceBgpCustomError, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBgpCustomError()
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

// DeviceBgpCustomError is a BGP peer can send NOTIFICATION message with user defined Error Code and Error Subcode.
type DeviceBgpCustomError interface {
	Validation
	// msg marshals DeviceBgpCustomError to protobuf object *otg.DeviceBgpCustomError
	// and doesn't set defaults
	msg() *otg.DeviceBgpCustomError
	// setMsg unmarshals DeviceBgpCustomError from protobuf object *otg.DeviceBgpCustomError
	// and doesn't set defaults
	setMsg(*otg.DeviceBgpCustomError) DeviceBgpCustomError
	// provides marshal interface
	Marshal() marshalDeviceBgpCustomError
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBgpCustomError
	// validate validates DeviceBgpCustomError
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBgpCustomError, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Code returns uint32, set in DeviceBgpCustomError.
	Code() uint32
	// SetCode assigns uint32 provided by user to DeviceBgpCustomError
	SetCode(value uint32) DeviceBgpCustomError
	// HasCode checks if Code has been set in DeviceBgpCustomError
	HasCode() bool
	// Subcode returns uint32, set in DeviceBgpCustomError.
	Subcode() uint32
	// SetSubcode assigns uint32 provided by user to DeviceBgpCustomError
	SetSubcode(value uint32) DeviceBgpCustomError
	// HasSubcode checks if Subcode has been set in DeviceBgpCustomError
	HasSubcode() bool
}

// The Error code to be sent in the NOTIFICATION message to peer.
// Code returns a uint32
func (obj *deviceBgpCustomError) Code() uint32 {

	return *obj.obj.Code

}

// The Error code to be sent in the NOTIFICATION message to peer.
// Code returns a uint32
func (obj *deviceBgpCustomError) HasCode() bool {
	return obj.obj.Code != nil
}

// The Error code to be sent in the NOTIFICATION message to peer.
// SetCode sets the uint32 value in the DeviceBgpCustomError object
func (obj *deviceBgpCustomError) SetCode(value uint32) DeviceBgpCustomError {

	obj.obj.Code = &value
	return obj
}

// The Error Subcode to be sent in the NOTIFICATION message to peer.
// Subcode returns a uint32
func (obj *deviceBgpCustomError) Subcode() uint32 {

	return *obj.obj.Subcode

}

// The Error Subcode to be sent in the NOTIFICATION message to peer.
// Subcode returns a uint32
func (obj *deviceBgpCustomError) HasSubcode() bool {
	return obj.obj.Subcode != nil
}

// The Error Subcode to be sent in the NOTIFICATION message to peer.
// SetSubcode sets the uint32 value in the DeviceBgpCustomError object
func (obj *deviceBgpCustomError) SetSubcode(value uint32) DeviceBgpCustomError {

	obj.obj.Subcode = &value
	return obj
}

func (obj *deviceBgpCustomError) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *deviceBgpCustomError) setDefault() {

}
