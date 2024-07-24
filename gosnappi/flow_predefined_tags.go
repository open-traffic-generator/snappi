package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowPredefinedTags *****
type flowPredefinedTags struct {
	validation
	obj          *otg.FlowPredefinedTags
	marshaller   marshalFlowPredefinedTags
	unMarshaller unMarshalFlowPredefinedTags
}

func NewFlowPredefinedTags() FlowPredefinedTags {
	obj := flowPredefinedTags{obj: &otg.FlowPredefinedTags{}}
	obj.setDefault()
	return &obj
}

func (obj *flowPredefinedTags) msg() *otg.FlowPredefinedTags {
	return obj.obj
}

func (obj *flowPredefinedTags) setMsg(msg *otg.FlowPredefinedTags) FlowPredefinedTags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowPredefinedTags struct {
	obj *flowPredefinedTags
}

type marshalFlowPredefinedTags interface {
	// ToProto marshals FlowPredefinedTags to protobuf object *otg.FlowPredefinedTags
	ToProto() (*otg.FlowPredefinedTags, error)
	// ToPbText marshals FlowPredefinedTags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowPredefinedTags to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowPredefinedTags to JSON text
	ToJson() (string, error)
}

type unMarshalflowPredefinedTags struct {
	obj *flowPredefinedTags
}

type unMarshalFlowPredefinedTags interface {
	// FromProto unmarshals FlowPredefinedTags from protobuf object *otg.FlowPredefinedTags
	FromProto(msg *otg.FlowPredefinedTags) (FlowPredefinedTags, error)
	// FromPbText unmarshals FlowPredefinedTags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowPredefinedTags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowPredefinedTags from JSON text
	FromJson(value string) error
}

func (obj *flowPredefinedTags) Marshal() marshalFlowPredefinedTags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowPredefinedTags{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowPredefinedTags) Unmarshal() unMarshalFlowPredefinedTags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowPredefinedTags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowPredefinedTags) ToProto() (*otg.FlowPredefinedTags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowPredefinedTags) FromProto(msg *otg.FlowPredefinedTags) (FlowPredefinedTags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowPredefinedTags) ToPbText() (string, error) {
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

func (m *unMarshalflowPredefinedTags) FromPbText(value string) error {
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

func (m *marshalflowPredefinedTags) ToYaml() (string, error) {
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

func (m *unMarshalflowPredefinedTags) FromYaml(value string) error {
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

func (m *marshalflowPredefinedTags) ToJson() (string, error) {
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

func (m *unMarshalflowPredefinedTags) FromJson(value string) error {
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

func (obj *flowPredefinedTags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowPredefinedTags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowPredefinedTags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowPredefinedTags) Clone() (FlowPredefinedTags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowPredefinedTags()
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

// FlowPredefinedTags is list of predefined flow tracking options, outside packet fields, that can be enabled.
type FlowPredefinedTags interface {
	Validation
	// msg marshals FlowPredefinedTags to protobuf object *otg.FlowPredefinedTags
	// and doesn't set defaults
	msg() *otg.FlowPredefinedTags
	// setMsg unmarshals FlowPredefinedTags from protobuf object *otg.FlowPredefinedTags
	// and doesn't set defaults
	setMsg(*otg.FlowPredefinedTags) FlowPredefinedTags
	// provides marshal interface
	Marshal() marshalFlowPredefinedTags
	// provides unmarshal interface
	Unmarshal() unMarshalFlowPredefinedTags
	// validate validates FlowPredefinedTags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowPredefinedTags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RxName returns bool, set in FlowPredefinedTags.
	RxName() bool
	// SetRxName assigns bool provided by user to FlowPredefinedTags
	SetRxName(value bool) FlowPredefinedTags
	// HasRxName checks if RxName has been set in FlowPredefinedTags
	HasRxName() bool
}

// Enables Rx port or lag level disaggregation with predefined metrics tag name set as "rx_name".
// The Rx port / lag names can be found under tagged_metrics tag names in flow metrics response.
// RxName returns a bool
func (obj *flowPredefinedTags) RxName() bool {

	return *obj.obj.RxName

}

// Enables Rx port or lag level disaggregation with predefined metrics tag name set as "rx_name".
// The Rx port / lag names can be found under tagged_metrics tag names in flow metrics response.
// RxName returns a bool
func (obj *flowPredefinedTags) HasRxName() bool {
	return obj.obj.RxName != nil
}

// Enables Rx port or lag level disaggregation with predefined metrics tag name set as "rx_name".
// The Rx port / lag names can be found under tagged_metrics tag names in flow metrics response.
// SetRxName sets the bool value in the FlowPredefinedTags object
func (obj *flowPredefinedTags) SetRxName(value bool) FlowPredefinedTags {

	obj.obj.RxName = &value
	return obj
}

func (obj *flowPredefinedTags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowPredefinedTags) setDefault() {
	if obj.obj.RxName == nil {
		obj.SetRxName(false)
	}

}
