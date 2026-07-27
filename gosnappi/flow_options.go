package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowOptions *****
type flowOptions struct {
	validation
	obj          *otg.FlowOptions
	marshaller   marshalFlowOptions
	unMarshaller unMarshalFlowOptions
}

func NewFlowOptions() FlowOptions {
	obj := flowOptions{obj: &otg.FlowOptions{}}
	obj.setDefault()
	return &obj
}

func (obj *flowOptions) msg() *otg.FlowOptions {
	return obj.obj
}

func (obj *flowOptions) setMsg(msg *otg.FlowOptions) FlowOptions {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowOptions struct {
	obj *flowOptions
}

type marshalFlowOptions interface {
	// ToProto marshals FlowOptions to protobuf object *otg.FlowOptions
	ToProto() (*otg.FlowOptions, error)
	// ToPbText marshals FlowOptions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowOptions to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowOptions to JSON text
	ToJson() (string, error)
}

type unMarshalflowOptions struct {
	obj *flowOptions
}

type unMarshalFlowOptions interface {
	// FromProto unmarshals FlowOptions from protobuf object *otg.FlowOptions
	FromProto(msg *otg.FlowOptions) (FlowOptions, error)
	// FromPbText unmarshals FlowOptions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowOptions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowOptions from JSON text
	FromJson(value string) error
}

func (obj *flowOptions) Marshal() marshalFlowOptions {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowOptions{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowOptions) Unmarshal() unMarshalFlowOptions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowOptions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowOptions) ToProto() (*otg.FlowOptions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowOptions) FromProto(msg *otg.FlowOptions) (FlowOptions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowOptions) ToPbText() (string, error) {
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

func (m *unMarshalflowOptions) FromPbText(value string) error {
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

func (m *marshalflowOptions) ToYaml() (string, error) {
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

func (m *unMarshalflowOptions) FromYaml(value string) error {
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

func (m *marshalflowOptions) ToJson() (string, error) {
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

func (m *unMarshalflowOptions) FromJson(value string) error {
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

func (obj *flowOptions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowOptions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowOptions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowOptions) Clone() (FlowOptions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowOptions()
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

// FlowOptions is common flow options that apply to all configured Flow objects.
type FlowOptions interface {
	Validation
	// msg marshals FlowOptions to protobuf object *otg.FlowOptions
	// and doesn't set defaults
	msg() *otg.FlowOptions
	// setMsg unmarshals FlowOptions from protobuf object *otg.FlowOptions
	// and doesn't set defaults
	setMsg(*otg.FlowOptions) FlowOptions
	// provides marshal interface
	Marshal() marshalFlowOptions
	// provides unmarshal interface
	Unmarshal() unMarshalFlowOptions
	// validate validates FlowOptions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowOptions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PacketLossDuration returns bool, set in FlowOptions.
	PacketLossDuration() bool
	// SetPacketLossDuration assigns bool provided by user to FlowOptions
	SetPacketLossDuration(value bool) FlowOptions
	// HasPacketLossDuration checks if PacketLossDuration has been set in FlowOptions
	HasPacketLossDuration() bool
}

// Estimated Convergence Time without received packets, calculated by lost frames at the expected Rx Rate. It applies to all flows and if set to true, it is reported in packet loss in flow statistics.
// PacketLossDuration returns a bool
func (obj *flowOptions) PacketLossDuration() bool {

	return *obj.obj.PacketLossDuration

}

// Estimated Convergence Time without received packets, calculated by lost frames at the expected Rx Rate. It applies to all flows and if set to true, it is reported in packet loss in flow statistics.
// PacketLossDuration returns a bool
func (obj *flowOptions) HasPacketLossDuration() bool {
	return obj.obj.PacketLossDuration != nil
}

// Estimated Convergence Time without received packets, calculated by lost frames at the expected Rx Rate. It applies to all flows and if set to true, it is reported in packet loss in flow statistics.
// SetPacketLossDuration sets the bool value in the FlowOptions object
func (obj *flowOptions) SetPacketLossDuration(value bool) FlowOptions {

	obj.obj.PacketLossDuration = &value
	return obj
}

func (obj *flowOptions) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowOptions) setDefault() {
	if obj.obj.PacketLossDuration == nil {
		obj.SetPacketLossDuration(false)
	}

}
