package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowPacketLossDuration *****
type flowPacketLossDuration struct {
	validation
	obj          *otg.FlowPacketLossDuration
	marshaller   marshalFlowPacketLossDuration
	unMarshaller unMarshalFlowPacketLossDuration
}

func NewFlowPacketLossDuration() FlowPacketLossDuration {
	obj := flowPacketLossDuration{obj: &otg.FlowPacketLossDuration{}}
	obj.setDefault()
	return &obj
}

func (obj *flowPacketLossDuration) msg() *otg.FlowPacketLossDuration {
	return obj.obj
}

func (obj *flowPacketLossDuration) setMsg(msg *otg.FlowPacketLossDuration) FlowPacketLossDuration {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowPacketLossDuration struct {
	obj *flowPacketLossDuration
}

type marshalFlowPacketLossDuration interface {
	// ToProto marshals FlowPacketLossDuration to protobuf object *otg.FlowPacketLossDuration
	ToProto() (*otg.FlowPacketLossDuration, error)
	// ToPbText marshals FlowPacketLossDuration to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowPacketLossDuration to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowPacketLossDuration to JSON text
	ToJson() (string, error)
}

type unMarshalflowPacketLossDuration struct {
	obj *flowPacketLossDuration
}

type unMarshalFlowPacketLossDuration interface {
	// FromProto unmarshals FlowPacketLossDuration from protobuf object *otg.FlowPacketLossDuration
	FromProto(msg *otg.FlowPacketLossDuration) (FlowPacketLossDuration, error)
	// FromPbText unmarshals FlowPacketLossDuration from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowPacketLossDuration from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowPacketLossDuration from JSON text
	FromJson(value string) error
}

func (obj *flowPacketLossDuration) Marshal() marshalFlowPacketLossDuration {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowPacketLossDuration{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowPacketLossDuration) Unmarshal() unMarshalFlowPacketLossDuration {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowPacketLossDuration{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowPacketLossDuration) ToProto() (*otg.FlowPacketLossDuration, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowPacketLossDuration) FromProto(msg *otg.FlowPacketLossDuration) (FlowPacketLossDuration, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowPacketLossDuration) ToPbText() (string, error) {
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

func (m *unMarshalflowPacketLossDuration) FromPbText(value string) error {
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

func (m *marshalflowPacketLossDuration) ToYaml() (string, error) {
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

func (m *unMarshalflowPacketLossDuration) FromYaml(value string) error {
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

func (m *marshalflowPacketLossDuration) ToJson() (string, error) {
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

func (m *unMarshalflowPacketLossDuration) FromJson(value string) error {
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

func (obj *flowPacketLossDuration) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowPacketLossDuration) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowPacketLossDuration) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowPacketLossDuration) Clone() (FlowPacketLossDuration, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowPacketLossDuration()
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

// FlowPacketLossDuration is the container for packet loss duration metrics. The container will be empty if
// options.flow_options.packet_loss_duration has not been enabled during set_config.
type FlowPacketLossDuration interface {
	Validation
	// msg marshals FlowPacketLossDuration to protobuf object *otg.FlowPacketLossDuration
	// and doesn't set defaults
	msg() *otg.FlowPacketLossDuration
	// setMsg unmarshals FlowPacketLossDuration from protobuf object *otg.FlowPacketLossDuration
	// and doesn't set defaults
	setMsg(*otg.FlowPacketLossDuration) FlowPacketLossDuration
	// provides marshal interface
	Marshal() marshalFlowPacketLossDuration
	// provides unmarshal interface
	Unmarshal() unMarshalFlowPacketLossDuration
	// validate validates FlowPacketLossDuration
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowPacketLossDuration, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns float32, set in FlowPacketLossDuration.
	Value() float32
	// SetValue assigns float32 provided by user to FlowPacketLossDuration
	SetValue(value float32) FlowPacketLossDuration
	// HasValue checks if Value has been set in FlowPacketLossDuration
	HasValue() bool
}

// Estimated packet loss duration in milliseconds, calculated based on the number
// of lost packets and Rx Rate.
// Value returns a float32
func (obj *flowPacketLossDuration) Value() float32 {

	return *obj.obj.Value

}

// Estimated packet loss duration in milliseconds, calculated based on the number
// of lost packets and Rx Rate.
// Value returns a float32
func (obj *flowPacketLossDuration) HasValue() bool {
	return obj.obj.Value != nil
}

// Estimated packet loss duration in milliseconds, calculated based on the number
// of lost packets and Rx Rate.
// SetValue sets the float32 value in the FlowPacketLossDuration object
func (obj *flowPacketLossDuration) SetValue(value float32) FlowPacketLossDuration {

	obj.obj.Value = &value
	return obj
}

func (obj *flowPacketLossDuration) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowPacketLossDuration) setDefault() {

}
