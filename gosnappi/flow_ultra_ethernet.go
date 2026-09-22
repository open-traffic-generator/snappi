package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowUltraEthernet *****
type flowUltraEthernet struct {
	validation
	obj          *otg.FlowUltraEthernet
	marshaller   marshalFlowUltraEthernet
	unMarshaller unMarshalFlowUltraEthernet
}

func NewFlowUltraEthernet() FlowUltraEthernet {
	obj := flowUltraEthernet{obj: &otg.FlowUltraEthernet{}}
	obj.setDefault()
	return &obj
}

func (obj *flowUltraEthernet) msg() *otg.FlowUltraEthernet {
	return obj.obj
}

func (obj *flowUltraEthernet) setMsg(msg *otg.FlowUltraEthernet) FlowUltraEthernet {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowUltraEthernet struct {
	obj *flowUltraEthernet
}

type marshalFlowUltraEthernet interface {
	// ToProto marshals FlowUltraEthernet to protobuf object *otg.FlowUltraEthernet
	ToProto() (*otg.FlowUltraEthernet, error)
	// ToPbText marshals FlowUltraEthernet to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowUltraEthernet to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowUltraEthernet to JSON text
	ToJson() (string, error)
}

type unMarshalflowUltraEthernet struct {
	obj *flowUltraEthernet
}

type unMarshalFlowUltraEthernet interface {
	// FromProto unmarshals FlowUltraEthernet from protobuf object *otg.FlowUltraEthernet
	FromProto(msg *otg.FlowUltraEthernet) (FlowUltraEthernet, error)
	// FromPbText unmarshals FlowUltraEthernet from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowUltraEthernet from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowUltraEthernet from JSON text
	FromJson(value string) error
}

func (obj *flowUltraEthernet) Marshal() marshalFlowUltraEthernet {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowUltraEthernet{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowUltraEthernet) Unmarshal() unMarshalFlowUltraEthernet {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowUltraEthernet{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowUltraEthernet) ToProto() (*otg.FlowUltraEthernet, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowUltraEthernet) FromProto(msg *otg.FlowUltraEthernet) (FlowUltraEthernet, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowUltraEthernet) ToPbText() (string, error) {
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

func (m *unMarshalflowUltraEthernet) FromPbText(value string) error {
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

func (m *marshalflowUltraEthernet) ToYaml() (string, error) {
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

func (m *unMarshalflowUltraEthernet) FromYaml(value string) error {
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

func (m *marshalflowUltraEthernet) ToJson() (string, error) {
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

func (m *unMarshalflowUltraEthernet) FromJson(value string) error {
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

func (obj *flowUltraEthernet) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowUltraEthernet) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowUltraEthernet) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowUltraEthernet) Clone() (FlowUltraEthernet, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowUltraEthernet()
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

// FlowUltraEthernet is per flow Ultra Ethernet (UEC) settings.
//
// Reference: UE-Specification-1.0.3 Section 5.1.
type FlowUltraEthernet interface {
	Validation
	// msg marshals FlowUltraEthernet to protobuf object *otg.FlowUltraEthernet
	// and doesn't set defaults
	msg() *otg.FlowUltraEthernet
	// setMsg unmarshals FlowUltraEthernet from protobuf object *otg.FlowUltraEthernet
	// and doesn't set defaults
	setMsg(*otg.FlowUltraEthernet) FlowUltraEthernet
	// provides marshal interface
	Marshal() marshalFlowUltraEthernet
	// provides unmarshal interface
	Unmarshal() unMarshalFlowUltraEthernet
	// validate validates FlowUltraEthernet
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowUltraEthernet, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LlrEligible returns bool, set in FlowUltraEthernet.
	LlrEligible() bool
	// SetLlrEligible assigns bool provided by user to FlowUltraEthernet
	SetLlrEligible(value bool) FlowUltraEthernet
	// HasLlrEligible checks if LlrEligible has been set in FlowUltraEthernet
	HasLlrEligible() bool
}

// Marks this flow's frames as Link Layer Retry (LLR) eligible, independent of
// the port level LLR enable. When true, the frames of this flow participate in
// LLR (sequence numbering and replay on loss); when false, they are transmitted
// as LLR-ineligible. Applies only on ports where LLR is enabled.
// LlrEligible returns a bool
func (obj *flowUltraEthernet) LlrEligible() bool {

	return *obj.obj.LlrEligible

}

// Marks this flow's frames as Link Layer Retry (LLR) eligible, independent of
// the port level LLR enable. When true, the frames of this flow participate in
// LLR (sequence numbering and replay on loss); when false, they are transmitted
// as LLR-ineligible. Applies only on ports where LLR is enabled.
// LlrEligible returns a bool
func (obj *flowUltraEthernet) HasLlrEligible() bool {
	return obj.obj.LlrEligible != nil
}

// Marks this flow's frames as Link Layer Retry (LLR) eligible, independent of
// the port level LLR enable. When true, the frames of this flow participate in
// LLR (sequence numbering and replay on loss); when false, they are transmitted
// as LLR-ineligible. Applies only on ports where LLR is enabled.
// SetLlrEligible sets the bool value in the FlowUltraEthernet object
func (obj *flowUltraEthernet) SetLlrEligible(value bool) FlowUltraEthernet {

	obj.obj.LlrEligible = &value
	return obj
}

func (obj *flowUltraEthernet) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowUltraEthernet) setDefault() {
	if obj.obj.LlrEligible == nil {
		obj.SetLlrEligible(false)
	}

}
