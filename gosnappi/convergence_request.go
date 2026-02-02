package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConvergenceRequest *****
type convergenceRequest struct {
	validation
	obj          *otg.ConvergenceRequest
	marshaller   marshalConvergenceRequest
	unMarshaller unMarshalConvergenceRequest
}

func NewConvergenceRequest() ConvergenceRequest {
	obj := convergenceRequest{obj: &otg.ConvergenceRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *convergenceRequest) msg() *otg.ConvergenceRequest {
	return obj.obj
}

func (obj *convergenceRequest) setMsg(msg *otg.ConvergenceRequest) ConvergenceRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconvergenceRequest struct {
	obj *convergenceRequest
}

type marshalConvergenceRequest interface {
	// ToProto marshals ConvergenceRequest to protobuf object *otg.ConvergenceRequest
	ToProto() (*otg.ConvergenceRequest, error)
	// ToPbText marshals ConvergenceRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConvergenceRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConvergenceRequest to JSON text
	ToJson() (string, error)
}

type unMarshalconvergenceRequest struct {
	obj *convergenceRequest
}

type unMarshalConvergenceRequest interface {
	// FromProto unmarshals ConvergenceRequest from protobuf object *otg.ConvergenceRequest
	FromProto(msg *otg.ConvergenceRequest) (ConvergenceRequest, error)
	// FromPbText unmarshals ConvergenceRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConvergenceRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConvergenceRequest from JSON text
	FromJson(value string) error
}

func (obj *convergenceRequest) Marshal() marshalConvergenceRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconvergenceRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *convergenceRequest) Unmarshal() unMarshalConvergenceRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconvergenceRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconvergenceRequest) ToProto() (*otg.ConvergenceRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconvergenceRequest) FromProto(msg *otg.ConvergenceRequest) (ConvergenceRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconvergenceRequest) ToPbText() (string, error) {
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

func (m *unMarshalconvergenceRequest) FromPbText(value string) error {
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

func (m *marshalconvergenceRequest) ToYaml() (string, error) {
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

func (m *unMarshalconvergenceRequest) FromYaml(value string) error {
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

func (m *marshalconvergenceRequest) ToJson() (string, error) {
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

func (m *unMarshalconvergenceRequest) FromJson(value string) error {
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

func (obj *convergenceRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *convergenceRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *convergenceRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *convergenceRequest) Clone() (ConvergenceRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConvergenceRequest()
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

// ConvergenceRequest is container for requesting control-plane and data-plane convergence time metrics for flows.
type ConvergenceRequest interface {
	Validation
	// msg marshals ConvergenceRequest to protobuf object *otg.ConvergenceRequest
	// and doesn't set defaults
	msg() *otg.ConvergenceRequest
	// setMsg unmarshals ConvergenceRequest from protobuf object *otg.ConvergenceRequest
	// and doesn't set defaults
	setMsg(*otg.ConvergenceRequest) ConvergenceRequest
	// provides marshal interface
	Marshal() marshalConvergenceRequest
	// provides unmarshal interface
	Unmarshal() unMarshalConvergenceRequest
	// validate validates ConvergenceRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConvergenceRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlowNames returns []string, set in ConvergenceRequest.
	FlowNames() []string
	// SetFlowNames assigns []string provided by user to ConvergenceRequest
	SetFlowNames(value []string) ConvergenceRequest
}

// Convergence metrics will be retrieved for these flow names.
// If no flow names are specified then convergence metrics for all flows will be returned.
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// FlowNames returns a []string
func (obj *convergenceRequest) FlowNames() []string {
	if obj.obj.FlowNames == nil {
		obj.obj.FlowNames = make([]string, 0)
	}
	return obj.obj.FlowNames
}

// Convergence metrics will be retrieved for these flow names.
// If no flow names are specified then convergence metrics for all flows will be returned.
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// SetFlowNames sets the []string value in the ConvergenceRequest object
func (obj *convergenceRequest) SetFlowNames(value []string) ConvergenceRequest {

	if obj.obj.FlowNames == nil {
		obj.obj.FlowNames = make([]string, 0)
	}
	obj.obj.FlowNames = value

	return obj
}

func (obj *convergenceRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *convergenceRequest) setDefault() {

}
