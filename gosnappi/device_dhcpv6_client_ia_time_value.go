package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceDhcpv6ClientIaTimeValue *****
type deviceDhcpv6ClientIaTimeValue struct {
	validation
	obj          *otg.DeviceDhcpv6ClientIaTimeValue
	marshaller   marshalDeviceDhcpv6ClientIaTimeValue
	unMarshaller unMarshalDeviceDhcpv6ClientIaTimeValue
}

func NewDeviceDhcpv6ClientIaTimeValue() DeviceDhcpv6ClientIaTimeValue {
	obj := deviceDhcpv6ClientIaTimeValue{obj: &otg.DeviceDhcpv6ClientIaTimeValue{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceDhcpv6ClientIaTimeValue) msg() *otg.DeviceDhcpv6ClientIaTimeValue {
	return obj.obj
}

func (obj *deviceDhcpv6ClientIaTimeValue) setMsg(msg *otg.DeviceDhcpv6ClientIaTimeValue) DeviceDhcpv6ClientIaTimeValue {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceDhcpv6ClientIaTimeValue struct {
	obj *deviceDhcpv6ClientIaTimeValue
}

type marshalDeviceDhcpv6ClientIaTimeValue interface {
	// ToProto marshals DeviceDhcpv6ClientIaTimeValue to protobuf object *otg.DeviceDhcpv6ClientIaTimeValue
	ToProto() (*otg.DeviceDhcpv6ClientIaTimeValue, error)
	// ToPbText marshals DeviceDhcpv6ClientIaTimeValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceDhcpv6ClientIaTimeValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceDhcpv6ClientIaTimeValue to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceDhcpv6ClientIaTimeValue struct {
	obj *deviceDhcpv6ClientIaTimeValue
}

type unMarshalDeviceDhcpv6ClientIaTimeValue interface {
	// FromProto unmarshals DeviceDhcpv6ClientIaTimeValue from protobuf object *otg.DeviceDhcpv6ClientIaTimeValue
	FromProto(msg *otg.DeviceDhcpv6ClientIaTimeValue) (DeviceDhcpv6ClientIaTimeValue, error)
	// FromPbText unmarshals DeviceDhcpv6ClientIaTimeValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceDhcpv6ClientIaTimeValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceDhcpv6ClientIaTimeValue from JSON text
	FromJson(value string) error
}

func (obj *deviceDhcpv6ClientIaTimeValue) Marshal() marshalDeviceDhcpv6ClientIaTimeValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceDhcpv6ClientIaTimeValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceDhcpv6ClientIaTimeValue) Unmarshal() unMarshalDeviceDhcpv6ClientIaTimeValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceDhcpv6ClientIaTimeValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceDhcpv6ClientIaTimeValue) ToProto() (*otg.DeviceDhcpv6ClientIaTimeValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceDhcpv6ClientIaTimeValue) FromProto(msg *otg.DeviceDhcpv6ClientIaTimeValue) (DeviceDhcpv6ClientIaTimeValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceDhcpv6ClientIaTimeValue) ToPbText() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientIaTimeValue) FromPbText(value string) error {
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

func (m *marshaldeviceDhcpv6ClientIaTimeValue) ToYaml() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientIaTimeValue) FromYaml(value string) error {
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

func (m *marshaldeviceDhcpv6ClientIaTimeValue) ToJson() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientIaTimeValue) FromJson(value string) error {
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

func (obj *deviceDhcpv6ClientIaTimeValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientIaTimeValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientIaTimeValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceDhcpv6ClientIaTimeValue) Clone() (DeviceDhcpv6ClientIaTimeValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceDhcpv6ClientIaTimeValue()
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

// DeviceDhcpv6ClientIaTimeValue is the container for the suggested times at which the client contacts the server or any available server.
type DeviceDhcpv6ClientIaTimeValue interface {
	Validation
	// msg marshals DeviceDhcpv6ClientIaTimeValue to protobuf object *otg.DeviceDhcpv6ClientIaTimeValue
	// and doesn't set defaults
	msg() *otg.DeviceDhcpv6ClientIaTimeValue
	// setMsg unmarshals DeviceDhcpv6ClientIaTimeValue from protobuf object *otg.DeviceDhcpv6ClientIaTimeValue
	// and doesn't set defaults
	setMsg(*otg.DeviceDhcpv6ClientIaTimeValue) DeviceDhcpv6ClientIaTimeValue
	// provides marshal interface
	Marshal() marshalDeviceDhcpv6ClientIaTimeValue
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceDhcpv6ClientIaTimeValue
	// validate validates DeviceDhcpv6ClientIaTimeValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceDhcpv6ClientIaTimeValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// T1 returns uint32, set in DeviceDhcpv6ClientIaTimeValue.
	T1() uint32
	// SetT1 assigns uint32 provided by user to DeviceDhcpv6ClientIaTimeValue
	SetT1(value uint32) DeviceDhcpv6ClientIaTimeValue
	// HasT1 checks if T1 has been set in DeviceDhcpv6ClientIaTimeValue
	HasT1() bool
	// T2 returns uint32, set in DeviceDhcpv6ClientIaTimeValue.
	T2() uint32
	// SetT2 assigns uint32 provided by user to DeviceDhcpv6ClientIaTimeValue
	SetT2(value uint32) DeviceDhcpv6ClientIaTimeValue
	// HasT2 checks if T2 has been set in DeviceDhcpv6ClientIaTimeValue
	HasT2() bool
}

// The suggested time at which the client contacts the server from which the addresses were obtained to  extend the lifetimes of the addresses assigned. T1 is a time duration relative to the current time expressed  in units of seconds. If set to 0 server will ignore it. If the maximum value is specified it means infinite time.
// T1 returns a uint32
func (obj *deviceDhcpv6ClientIaTimeValue) T1() uint32 {

	return *obj.obj.T1

}

// The suggested time at which the client contacts the server from which the addresses were obtained to  extend the lifetimes of the addresses assigned. T1 is a time duration relative to the current time expressed  in units of seconds. If set to 0 server will ignore it. If the maximum value is specified it means infinite time.
// T1 returns a uint32
func (obj *deviceDhcpv6ClientIaTimeValue) HasT1() bool {
	return obj.obj.T1 != nil
}

// The suggested time at which the client contacts the server from which the addresses were obtained to  extend the lifetimes of the addresses assigned. T1 is a time duration relative to the current time expressed  in units of seconds. If set to 0 server will ignore it. If the maximum value is specified it means infinite time.
// SetT1 sets the uint32 value in the DeviceDhcpv6ClientIaTimeValue object
func (obj *deviceDhcpv6ClientIaTimeValue) SetT1(value uint32) DeviceDhcpv6ClientIaTimeValue {

	obj.obj.T1 = &value
	return obj
}

// The suggested time at which the client contacts any available server to extend the lifetimes of the addresses assigned.  T2 is a time duration relative to the current time expressed in units of seconds. If set to 0 server will ignore it. If  the maximum value is specified it means infinite time
// T2 returns a uint32
func (obj *deviceDhcpv6ClientIaTimeValue) T2() uint32 {

	return *obj.obj.T2

}

// The suggested time at which the client contacts any available server to extend the lifetimes of the addresses assigned.  T2 is a time duration relative to the current time expressed in units of seconds. If set to 0 server will ignore it. If  the maximum value is specified it means infinite time
// T2 returns a uint32
func (obj *deviceDhcpv6ClientIaTimeValue) HasT2() bool {
	return obj.obj.T2 != nil
}

// The suggested time at which the client contacts any available server to extend the lifetimes of the addresses assigned.  T2 is a time duration relative to the current time expressed in units of seconds. If set to 0 server will ignore it. If  the maximum value is specified it means infinite time
// SetT2 sets the uint32 value in the DeviceDhcpv6ClientIaTimeValue object
func (obj *deviceDhcpv6ClientIaTimeValue) SetT2(value uint32) DeviceDhcpv6ClientIaTimeValue {

	obj.obj.T2 = &value
	return obj
}

func (obj *deviceDhcpv6ClientIaTimeValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.T1 != nil {

		if *obj.obj.T1 > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= DeviceDhcpv6ClientIaTimeValue.T1 <= 4294967295 but Got %d", *obj.obj.T1))
		}

	}

	if obj.obj.T2 != nil {

		if *obj.obj.T2 > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= DeviceDhcpv6ClientIaTimeValue.T2 <= 4294967295 but Got %d", *obj.obj.T2))
		}

	}

}

func (obj *deviceDhcpv6ClientIaTimeValue) setDefault() {
	if obj.obj.T1 == nil {
		obj.SetT1(302400)
	}
	if obj.obj.T2 == nil {
		obj.SetT2(483840)
	}

}
