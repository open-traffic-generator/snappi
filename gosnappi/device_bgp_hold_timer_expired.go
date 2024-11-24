package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBgpHoldTimerExpired *****
type deviceBgpHoldTimerExpired struct {
	validation
	obj          *otg.DeviceBgpHoldTimerExpired
	marshaller   marshalDeviceBgpHoldTimerExpired
	unMarshaller unMarshalDeviceBgpHoldTimerExpired
}

func NewDeviceBgpHoldTimerExpired() DeviceBgpHoldTimerExpired {
	obj := deviceBgpHoldTimerExpired{obj: &otg.DeviceBgpHoldTimerExpired{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBgpHoldTimerExpired) msg() *otg.DeviceBgpHoldTimerExpired {
	return obj.obj
}

func (obj *deviceBgpHoldTimerExpired) setMsg(msg *otg.DeviceBgpHoldTimerExpired) DeviceBgpHoldTimerExpired {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBgpHoldTimerExpired struct {
	obj *deviceBgpHoldTimerExpired
}

type marshalDeviceBgpHoldTimerExpired interface {
	// ToProto marshals DeviceBgpHoldTimerExpired to protobuf object *otg.DeviceBgpHoldTimerExpired
	ToProto() (*otg.DeviceBgpHoldTimerExpired, error)
	// ToPbText marshals DeviceBgpHoldTimerExpired to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBgpHoldTimerExpired to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBgpHoldTimerExpired to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals DeviceBgpHoldTimerExpired to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldeviceBgpHoldTimerExpired struct {
	obj *deviceBgpHoldTimerExpired
}

type unMarshalDeviceBgpHoldTimerExpired interface {
	// FromProto unmarshals DeviceBgpHoldTimerExpired from protobuf object *otg.DeviceBgpHoldTimerExpired
	FromProto(msg *otg.DeviceBgpHoldTimerExpired) (DeviceBgpHoldTimerExpired, error)
	// FromPbText unmarshals DeviceBgpHoldTimerExpired from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBgpHoldTimerExpired from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBgpHoldTimerExpired from JSON text
	FromJson(value string) error
}

func (obj *deviceBgpHoldTimerExpired) Marshal() marshalDeviceBgpHoldTimerExpired {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBgpHoldTimerExpired{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBgpHoldTimerExpired) Unmarshal() unMarshalDeviceBgpHoldTimerExpired {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBgpHoldTimerExpired{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBgpHoldTimerExpired) ToProto() (*otg.DeviceBgpHoldTimerExpired, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBgpHoldTimerExpired) FromProto(msg *otg.DeviceBgpHoldTimerExpired) (DeviceBgpHoldTimerExpired, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBgpHoldTimerExpired) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBgpHoldTimerExpired) FromPbText(value string) error {
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

func (m *marshaldeviceBgpHoldTimerExpired) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBgpHoldTimerExpired) FromYaml(value string) error {
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

func (m *marshaldeviceBgpHoldTimerExpired) ToJsonRaw() (string, error) {
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

func (m *marshaldeviceBgpHoldTimerExpired) ToJson() (string, error) {
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

func (m *unMarshaldeviceBgpHoldTimerExpired) FromJson(value string) error {
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

func (obj *deviceBgpHoldTimerExpired) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBgpHoldTimerExpired) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBgpHoldTimerExpired) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBgpHoldTimerExpired) Clone() (DeviceBgpHoldTimerExpired, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBgpHoldTimerExpired()
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

// DeviceBgpHoldTimerExpired is if a system does not receive successive KEEPALIVE, UPDATE, and/or NOTIFICATION messages within the period specified  in the Hold Time field of the OPEN message, then the NOTIFICATION message with the Hold Timer Expired Error Code(Error Code 4) is  sent and the BGP connection is closed. The Sub Code used is 0. If a user wants to use non zero Sub Code then CustomError can be used.
type DeviceBgpHoldTimerExpired interface {
	Validation
	// msg marshals DeviceBgpHoldTimerExpired to protobuf object *otg.DeviceBgpHoldTimerExpired
	// and doesn't set defaults
	msg() *otg.DeviceBgpHoldTimerExpired
	// setMsg unmarshals DeviceBgpHoldTimerExpired from protobuf object *otg.DeviceBgpHoldTimerExpired
	// and doesn't set defaults
	setMsg(*otg.DeviceBgpHoldTimerExpired) DeviceBgpHoldTimerExpired
	// provides marshal interface
	Marshal() marshalDeviceBgpHoldTimerExpired
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBgpHoldTimerExpired
	// validate validates DeviceBgpHoldTimerExpired
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBgpHoldTimerExpired, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
}

func (obj *deviceBgpHoldTimerExpired) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *deviceBgpHoldTimerExpired) setDefault() {

}
