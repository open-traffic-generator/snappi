package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBgpFiniteStateMachineError *****
type deviceBgpFiniteStateMachineError struct {
	validation
	obj          *otg.DeviceBgpFiniteStateMachineError
	marshaller   marshalDeviceBgpFiniteStateMachineError
	unMarshaller unMarshalDeviceBgpFiniteStateMachineError
}

func NewDeviceBgpFiniteStateMachineError() DeviceBgpFiniteStateMachineError {
	obj := deviceBgpFiniteStateMachineError{obj: &otg.DeviceBgpFiniteStateMachineError{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBgpFiniteStateMachineError) msg() *otg.DeviceBgpFiniteStateMachineError {
	return obj.obj
}

func (obj *deviceBgpFiniteStateMachineError) setMsg(msg *otg.DeviceBgpFiniteStateMachineError) DeviceBgpFiniteStateMachineError {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBgpFiniteStateMachineError struct {
	obj *deviceBgpFiniteStateMachineError
}

type marshalDeviceBgpFiniteStateMachineError interface {
	// ToProto marshals DeviceBgpFiniteStateMachineError to protobuf object *otg.DeviceBgpFiniteStateMachineError
	ToProto() (*otg.DeviceBgpFiniteStateMachineError, error)
	// ToPbText marshals DeviceBgpFiniteStateMachineError to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBgpFiniteStateMachineError to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBgpFiniteStateMachineError to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals DeviceBgpFiniteStateMachineError to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldeviceBgpFiniteStateMachineError struct {
	obj *deviceBgpFiniteStateMachineError
}

type unMarshalDeviceBgpFiniteStateMachineError interface {
	// FromProto unmarshals DeviceBgpFiniteStateMachineError from protobuf object *otg.DeviceBgpFiniteStateMachineError
	FromProto(msg *otg.DeviceBgpFiniteStateMachineError) (DeviceBgpFiniteStateMachineError, error)
	// FromPbText unmarshals DeviceBgpFiniteStateMachineError from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBgpFiniteStateMachineError from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBgpFiniteStateMachineError from JSON text
	FromJson(value string) error
}

func (obj *deviceBgpFiniteStateMachineError) Marshal() marshalDeviceBgpFiniteStateMachineError {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBgpFiniteStateMachineError{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBgpFiniteStateMachineError) Unmarshal() unMarshalDeviceBgpFiniteStateMachineError {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBgpFiniteStateMachineError{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBgpFiniteStateMachineError) ToProto() (*otg.DeviceBgpFiniteStateMachineError, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBgpFiniteStateMachineError) FromProto(msg *otg.DeviceBgpFiniteStateMachineError) (DeviceBgpFiniteStateMachineError, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBgpFiniteStateMachineError) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBgpFiniteStateMachineError) FromPbText(value string) error {
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

func (m *marshaldeviceBgpFiniteStateMachineError) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBgpFiniteStateMachineError) FromYaml(value string) error {
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

func (m *marshaldeviceBgpFiniteStateMachineError) ToJsonRaw() (string, error) {
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

func (m *marshaldeviceBgpFiniteStateMachineError) ToJson() (string, error) {
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

func (m *unMarshaldeviceBgpFiniteStateMachineError) FromJson(value string) error {
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

func (obj *deviceBgpFiniteStateMachineError) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBgpFiniteStateMachineError) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBgpFiniteStateMachineError) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBgpFiniteStateMachineError) Clone() (DeviceBgpFiniteStateMachineError, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBgpFiniteStateMachineError()
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

// DeviceBgpFiniteStateMachineError is any error detected by the BGP Finite State Machine (e.g., receipt of an unexpected event) is indicated by  sending the NOTIFICATION message with the Error Code-Finite State Machine Error(Error Code 5). The Sub Code used is 0.  If a user wants to use non zero Sub Code then CustomError can be used.
type DeviceBgpFiniteStateMachineError interface {
	Validation
	// msg marshals DeviceBgpFiniteStateMachineError to protobuf object *otg.DeviceBgpFiniteStateMachineError
	// and doesn't set defaults
	msg() *otg.DeviceBgpFiniteStateMachineError
	// setMsg unmarshals DeviceBgpFiniteStateMachineError from protobuf object *otg.DeviceBgpFiniteStateMachineError
	// and doesn't set defaults
	setMsg(*otg.DeviceBgpFiniteStateMachineError) DeviceBgpFiniteStateMachineError
	// provides marshal interface
	Marshal() marshalDeviceBgpFiniteStateMachineError
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBgpFiniteStateMachineError
	// validate validates DeviceBgpFiniteStateMachineError
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBgpFiniteStateMachineError, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
}

func (obj *deviceBgpFiniteStateMachineError) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *deviceBgpFiniteStateMachineError) setDefault() {

}
