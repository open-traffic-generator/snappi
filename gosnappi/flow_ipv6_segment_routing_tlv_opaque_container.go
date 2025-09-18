package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingTlvOpaqueContainer *****
type flowIpv6SegmentRoutingTlvOpaqueContainer struct {
	validation
	obj          *otg.FlowIpv6SegmentRoutingTlvOpaqueContainer
	marshaller   marshalFlowIpv6SegmentRoutingTlvOpaqueContainer
	unMarshaller unMarshalFlowIpv6SegmentRoutingTlvOpaqueContainer
}

func NewFlowIpv6SegmentRoutingTlvOpaqueContainer() FlowIpv6SegmentRoutingTlvOpaqueContainer {
	obj := flowIpv6SegmentRoutingTlvOpaqueContainer{obj: &otg.FlowIpv6SegmentRoutingTlvOpaqueContainer{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingTlvOpaqueContainer) msg() *otg.FlowIpv6SegmentRoutingTlvOpaqueContainer {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingTlvOpaqueContainer) setMsg(msg *otg.FlowIpv6SegmentRoutingTlvOpaqueContainer) FlowIpv6SegmentRoutingTlvOpaqueContainer {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingTlvOpaqueContainer struct {
	obj *flowIpv6SegmentRoutingTlvOpaqueContainer
}

type marshalFlowIpv6SegmentRoutingTlvOpaqueContainer interface {
	// ToProto marshals FlowIpv6SegmentRoutingTlvOpaqueContainer to protobuf object *otg.FlowIpv6SegmentRoutingTlvOpaqueContainer
	ToProto() (*otg.FlowIpv6SegmentRoutingTlvOpaqueContainer, error)
	// ToPbText marshals FlowIpv6SegmentRoutingTlvOpaqueContainer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingTlvOpaqueContainer to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingTlvOpaqueContainer to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingTlvOpaqueContainer struct {
	obj *flowIpv6SegmentRoutingTlvOpaqueContainer
}

type unMarshalFlowIpv6SegmentRoutingTlvOpaqueContainer interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingTlvOpaqueContainer from protobuf object *otg.FlowIpv6SegmentRoutingTlvOpaqueContainer
	FromProto(msg *otg.FlowIpv6SegmentRoutingTlvOpaqueContainer) (FlowIpv6SegmentRoutingTlvOpaqueContainer, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingTlvOpaqueContainer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingTlvOpaqueContainer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingTlvOpaqueContainer from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingTlvOpaqueContainer) Marshal() marshalFlowIpv6SegmentRoutingTlvOpaqueContainer {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingTlvOpaqueContainer{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingTlvOpaqueContainer) Unmarshal() unMarshalFlowIpv6SegmentRoutingTlvOpaqueContainer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingTlvOpaqueContainer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingTlvOpaqueContainer) ToProto() (*otg.FlowIpv6SegmentRoutingTlvOpaqueContainer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingTlvOpaqueContainer) FromProto(msg *otg.FlowIpv6SegmentRoutingTlvOpaqueContainer) (FlowIpv6SegmentRoutingTlvOpaqueContainer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingTlvOpaqueContainer) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvOpaqueContainer) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingTlvOpaqueContainer) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvOpaqueContainer) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingTlvOpaqueContainer) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvOpaqueContainer) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingTlvOpaqueContainer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingTlvOpaqueContainer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingTlvOpaqueContainer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingTlvOpaqueContainer) Clone() (FlowIpv6SegmentRoutingTlvOpaqueContainer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingTlvOpaqueContainer()
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

// FlowIpv6SegmentRoutingTlvOpaqueContainer is a TLV used to carry opaque data across the SR domain.
type FlowIpv6SegmentRoutingTlvOpaqueContainer interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingTlvOpaqueContainer to protobuf object *otg.FlowIpv6SegmentRoutingTlvOpaqueContainer
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingTlvOpaqueContainer
	// setMsg unmarshals FlowIpv6SegmentRoutingTlvOpaqueContainer from protobuf object *otg.FlowIpv6SegmentRoutingTlvOpaqueContainer
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingTlvOpaqueContainer) FlowIpv6SegmentRoutingTlvOpaqueContainer
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingTlvOpaqueContainer
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingTlvOpaqueContainer
	// validate validates FlowIpv6SegmentRoutingTlvOpaqueContainer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingTlvOpaqueContainer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns string, set in FlowIpv6SegmentRoutingTlvOpaqueContainer.
	Value() string
	// SetValue assigns string provided by user to FlowIpv6SegmentRoutingTlvOpaqueContainer
	SetValue(value string) FlowIpv6SegmentRoutingTlvOpaqueContainer
	// HasValue checks if Value has been set in FlowIpv6SegmentRoutingTlvOpaqueContainer
	HasValue() bool
}

// The variable-length opaque data carried by the TLV.
// Value returns a string
func (obj *flowIpv6SegmentRoutingTlvOpaqueContainer) Value() string {

	return *obj.obj.Value

}

// The variable-length opaque data carried by the TLV.
// Value returns a string
func (obj *flowIpv6SegmentRoutingTlvOpaqueContainer) HasValue() bool {
	return obj.obj.Value != nil
}

// The variable-length opaque data carried by the TLV.
// SetValue sets the string value in the FlowIpv6SegmentRoutingTlvOpaqueContainer object
func (obj *flowIpv6SegmentRoutingTlvOpaqueContainer) SetValue(value string) FlowIpv6SegmentRoutingTlvOpaqueContainer {

	obj.obj.Value = &value
	return obj
}

func (obj *flowIpv6SegmentRoutingTlvOpaqueContainer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateHex(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowIpv6SegmentRoutingTlvOpaqueContainer.Value"))
		}

	}

}

func (obj *flowIpv6SegmentRoutingTlvOpaqueContainer) setDefault() {
	if obj.obj.Value == nil {
		obj.SetValue("00")
	}

}
