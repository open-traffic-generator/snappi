package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowPacketLoss *****
type flowPacketLoss struct {
	validation
	obj          *otg.FlowPacketLoss
	marshaller   marshalFlowPacketLoss
	unMarshaller unMarshalFlowPacketLoss
}

func NewFlowPacketLoss() FlowPacketLoss {
	obj := flowPacketLoss{obj: &otg.FlowPacketLoss{}}
	obj.setDefault()
	return &obj
}

func (obj *flowPacketLoss) msg() *otg.FlowPacketLoss {
	return obj.obj
}

func (obj *flowPacketLoss) setMsg(msg *otg.FlowPacketLoss) FlowPacketLoss {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowPacketLoss struct {
	obj *flowPacketLoss
}

type marshalFlowPacketLoss interface {
	// ToProto marshals FlowPacketLoss to protobuf object *otg.FlowPacketLoss
	ToProto() (*otg.FlowPacketLoss, error)
	// ToPbText marshals FlowPacketLoss to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowPacketLoss to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowPacketLoss to JSON text
	ToJson() (string, error)
}

type unMarshalflowPacketLoss struct {
	obj *flowPacketLoss
}

type unMarshalFlowPacketLoss interface {
	// FromProto unmarshals FlowPacketLoss from protobuf object *otg.FlowPacketLoss
	FromProto(msg *otg.FlowPacketLoss) (FlowPacketLoss, error)
	// FromPbText unmarshals FlowPacketLoss from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowPacketLoss from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowPacketLoss from JSON text
	FromJson(value string) error
}

func (obj *flowPacketLoss) Marshal() marshalFlowPacketLoss {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowPacketLoss{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowPacketLoss) Unmarshal() unMarshalFlowPacketLoss {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowPacketLoss{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowPacketLoss) ToProto() (*otg.FlowPacketLoss, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowPacketLoss) FromProto(msg *otg.FlowPacketLoss) (FlowPacketLoss, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowPacketLoss) ToPbText() (string, error) {
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

func (m *unMarshalflowPacketLoss) FromPbText(value string) error {
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

func (m *marshalflowPacketLoss) ToYaml() (string, error) {
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

func (m *unMarshalflowPacketLoss) FromYaml(value string) error {
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

func (m *marshalflowPacketLoss) ToJson() (string, error) {
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

func (m *unMarshalflowPacketLoss) FromJson(value string) error {
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

func (obj *flowPacketLoss) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowPacketLoss) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowPacketLoss) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowPacketLoss) Clone() (FlowPacketLoss, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowPacketLoss()
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

// FlowPacketLoss is the container for packet loss duration metrics. The container will be empty if
// options.flow_options.packet_loss_duration has not been enabled during set_config.
type FlowPacketLoss interface {
	Validation
	// msg marshals FlowPacketLoss to protobuf object *otg.FlowPacketLoss
	// and doesn't set defaults
	msg() *otg.FlowPacketLoss
	// setMsg unmarshals FlowPacketLoss from protobuf object *otg.FlowPacketLoss
	// and doesn't set defaults
	setMsg(*otg.FlowPacketLoss) FlowPacketLoss
	// provides marshal interface
	Marshal() marshalFlowPacketLoss
	// provides unmarshal interface
	Unmarshal() unMarshalFlowPacketLoss
	// validate validates FlowPacketLoss
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowPacketLoss, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Duration returns float32, set in FlowPacketLoss.
	Duration() float32
	// SetDuration assigns float32 provided by user to FlowPacketLoss
	SetDuration(value float32) FlowPacketLoss
	// HasDuration checks if Duration has been set in FlowPacketLoss
	HasDuration() bool
}

// Estimated time in milli second for the lost packets at Rx Rate
// Duration returns a float32
func (obj *flowPacketLoss) Duration() float32 {

	return *obj.obj.Duration

}

// Estimated time in milli second for the lost packets at Rx Rate
// Duration returns a float32
func (obj *flowPacketLoss) HasDuration() bool {
	return obj.obj.Duration != nil
}

// Estimated time in milli second for the lost packets at Rx Rate
// SetDuration sets the float32 value in the FlowPacketLoss object
func (obj *flowPacketLoss) SetDuration(value float32) FlowPacketLoss {

	obj.obj.Duration = &value
	return obj
}

func (obj *flowPacketLoss) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowPacketLoss) setDefault() {

}
