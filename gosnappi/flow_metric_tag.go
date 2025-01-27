package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowMetricTag *****
type flowMetricTag struct {
	validation
	obj          *otg.FlowMetricTag
	marshaller   marshalFlowMetricTag
	unMarshaller unMarshalFlowMetricTag
	valueHolder  FlowMetricTagValue
}

func NewFlowMetricTag() FlowMetricTag {
	obj := flowMetricTag{obj: &otg.FlowMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *flowMetricTag) msg() *otg.FlowMetricTag {
	return obj.obj
}

func (obj *flowMetricTag) setMsg(msg *otg.FlowMetricTag) FlowMetricTag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowMetricTag struct {
	obj *flowMetricTag
}

type marshalFlowMetricTag interface {
	// ToProto marshals FlowMetricTag to protobuf object *otg.FlowMetricTag
	ToProto() (*otg.FlowMetricTag, error)
	// ToPbText marshals FlowMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowMetricTag struct {
	obj *flowMetricTag
}

type unMarshalFlowMetricTag interface {
	// FromProto unmarshals FlowMetricTag from protobuf object *otg.FlowMetricTag
	FromProto(msg *otg.FlowMetricTag) (FlowMetricTag, error)
	// FromPbText unmarshals FlowMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowMetricTag from JSON text
	FromJson(value string) error
}

func (obj *flowMetricTag) Marshal() marshalFlowMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowMetricTag) Unmarshal() unMarshalFlowMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowMetricTag) ToProto() (*otg.FlowMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowMetricTag) FromProto(msg *otg.FlowMetricTag) (FlowMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalflowMetricTag) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalflowMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalflowMetricTag) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalflowMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalflowMetricTag) ToJson() (string, error) {
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

func (m *unMarshalflowMetricTag) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *flowMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowMetricTag) Clone() (FlowMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowMetricTag()
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

func (obj *flowMetricTag) setNil() {
	obj.valueHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowMetricTag is description is TBD
type FlowMetricTag interface {
	Validation
	// msg marshals FlowMetricTag to protobuf object *otg.FlowMetricTag
	// and doesn't set defaults
	msg() *otg.FlowMetricTag
	// setMsg unmarshals FlowMetricTag from protobuf object *otg.FlowMetricTag
	// and doesn't set defaults
	setMsg(*otg.FlowMetricTag) FlowMetricTag
	// provides marshal interface
	Marshal() marshalFlowMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalFlowMetricTag
	// validate validates FlowMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in FlowMetricTag.
	Name() string
	// SetName assigns string provided by user to FlowMetricTag
	SetName(value string) FlowMetricTag
	// HasName checks if Name has been set in FlowMetricTag
	HasName() bool
	// Value returns FlowMetricTagValue, set in FlowMetricTag.
	// FlowMetricTagValue is a container for metric tag value
	Value() FlowMetricTagValue
	// SetValue assigns FlowMetricTagValue provided by user to FlowMetricTag.
	// FlowMetricTagValue is a container for metric tag value
	SetValue(value FlowMetricTagValue) FlowMetricTag
	// HasValue checks if Value has been set in FlowMetricTag
	HasValue() bool
	setNil()
}

// Name of packet field metric tag
// Name returns a string
func (obj *flowMetricTag) Name() string {

	return *obj.obj.Name

}

// Name of packet field metric tag
// Name returns a string
func (obj *flowMetricTag) HasName() bool {
	return obj.obj.Name != nil
}

// Name of packet field metric tag
// SetName sets the string value in the FlowMetricTag object
func (obj *flowMetricTag) SetName(value string) FlowMetricTag {

	obj.obj.Name = &value
	return obj
}

// description is TBD
// Value returns a FlowMetricTagValue
func (obj *flowMetricTag) Value() FlowMetricTagValue {
	if obj.obj.Value == nil {
		obj.obj.Value = NewFlowMetricTagValue().msg()
	}
	if obj.valueHolder == nil {
		obj.valueHolder = &flowMetricTagValue{obj: obj.obj.Value}
	}
	return obj.valueHolder
}

// description is TBD
// Value returns a FlowMetricTagValue
func (obj *flowMetricTag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the FlowMetricTagValue value in the FlowMetricTag object
func (obj *flowMetricTag) SetValue(value FlowMetricTagValue) FlowMetricTag {

	obj.valueHolder = nil
	obj.obj.Value = value.msg()

	return obj
}

func (obj *flowMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		obj.Value().validateObj(vObj, set_default)
	}

}

func (obj *flowMetricTag) setDefault() {

}
