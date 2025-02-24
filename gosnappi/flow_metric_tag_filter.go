package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowMetricTagFilter *****
type flowMetricTagFilter struct {
	validation
	obj          *otg.FlowMetricTagFilter
	marshaller   marshalFlowMetricTagFilter
	unMarshaller unMarshalFlowMetricTagFilter
}

func NewFlowMetricTagFilter() FlowMetricTagFilter {
	obj := flowMetricTagFilter{obj: &otg.FlowMetricTagFilter{}}
	obj.setDefault()
	return &obj
}

func (obj *flowMetricTagFilter) msg() *otg.FlowMetricTagFilter {
	return obj.obj
}

func (obj *flowMetricTagFilter) setMsg(msg *otg.FlowMetricTagFilter) FlowMetricTagFilter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowMetricTagFilter struct {
	obj *flowMetricTagFilter
}

type marshalFlowMetricTagFilter interface {
	// ToProto marshals FlowMetricTagFilter to protobuf object *otg.FlowMetricTagFilter
	ToProto() (*otg.FlowMetricTagFilter, error)
	// ToPbText marshals FlowMetricTagFilter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowMetricTagFilter to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowMetricTagFilter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowMetricTagFilter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowMetricTagFilter struct {
	obj *flowMetricTagFilter
}

type unMarshalFlowMetricTagFilter interface {
	// FromProto unmarshals FlowMetricTagFilter from protobuf object *otg.FlowMetricTagFilter
	FromProto(msg *otg.FlowMetricTagFilter) (FlowMetricTagFilter, error)
	// FromPbText unmarshals FlowMetricTagFilter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowMetricTagFilter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowMetricTagFilter from JSON text
	FromJson(value string) error
}

func (obj *flowMetricTagFilter) Marshal() marshalFlowMetricTagFilter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowMetricTagFilter{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowMetricTagFilter) Unmarshal() unMarshalFlowMetricTagFilter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowMetricTagFilter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowMetricTagFilter) ToProto() (*otg.FlowMetricTagFilter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowMetricTagFilter) FromProto(msg *otg.FlowMetricTagFilter) (FlowMetricTagFilter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowMetricTagFilter) ToPbText() (string, error) {
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

func (m *unMarshalflowMetricTagFilter) FromPbText(value string) error {
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

func (m *marshalflowMetricTagFilter) ToYaml() (string, error) {
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

func (m *unMarshalflowMetricTagFilter) FromYaml(value string) error {
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

func (m *marshalflowMetricTagFilter) ToJsonRaw() (string, error) {
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

func (m *marshalflowMetricTagFilter) ToJson() (string, error) {
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

func (m *unMarshalflowMetricTagFilter) FromJson(value string) error {
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

func (obj *flowMetricTagFilter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowMetricTagFilter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowMetricTagFilter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowMetricTagFilter) Clone() (FlowMetricTagFilter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowMetricTagFilter()
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

// FlowMetricTagFilter is a container for filtering ingress and/or egress metric tags.
// The Tx stats may not be applicable in both the request and response filter.
type FlowMetricTagFilter interface {
	Validation
	// msg marshals FlowMetricTagFilter to protobuf object *otg.FlowMetricTagFilter
	// and doesn't set defaults
	msg() *otg.FlowMetricTagFilter
	// setMsg unmarshals FlowMetricTagFilter from protobuf object *otg.FlowMetricTagFilter
	// and doesn't set defaults
	setMsg(*otg.FlowMetricTagFilter) FlowMetricTagFilter
	// provides marshal interface
	Marshal() marshalFlowMetricTagFilter
	// provides unmarshal interface
	Unmarshal() unMarshalFlowMetricTagFilter
	// validate validates FlowMetricTagFilter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowMetricTagFilter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in FlowMetricTagFilter.
	Name() string
	// SetName assigns string provided by user to FlowMetricTagFilter
	SetName(value string) FlowMetricTagFilter
	// HasName checks if Name has been set in FlowMetricTagFilter
	HasName() bool
	// Values returns []string, set in FlowMetricTagFilter.
	Values() []string
	// SetValues assigns []string provided by user to FlowMetricTagFilter
	SetValues(value []string) FlowMetricTagFilter
}

// A metric tag name that MUST exist in a flow packet or
// flow egress_packet configuration
// Name returns a string
func (obj *flowMetricTagFilter) Name() string {

	return *obj.obj.Name

}

// A metric tag name that MUST exist in a flow packet or
// flow egress_packet configuration
// Name returns a string
func (obj *flowMetricTagFilter) HasName() bool {
	return obj.obj.Name != nil
}

// A metric tag name that MUST exist in a flow packet or
// flow egress_packet configuration
// SetName sets the string value in the FlowMetricTagFilter object
func (obj *flowMetricTagFilter) SetName(value string) FlowMetricTagFilter {

	obj.obj.Name = &value
	return obj
}

// A list of filters that can be applied to the metric tag name.
// By default all values will be included in the flow metric results.
// Values returns a []string
func (obj *flowMetricTagFilter) Values() []string {
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	return obj.obj.Values
}

// A list of filters that can be applied to the metric tag name.
// By default all values will be included in the flow metric results.
// SetValues sets the []string value in the FlowMetricTagFilter object
func (obj *flowMetricTagFilter) SetValues(value []string) FlowMetricTagFilter {

	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

func (obj *flowMetricTagFilter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowMetricTagFilter) setDefault() {

}
