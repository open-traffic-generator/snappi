package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4Dscp *****
type flowIpv4Dscp struct {
	validation
	obj          *otg.FlowIpv4Dscp
	marshaller   marshalFlowIpv4Dscp
	unMarshaller unMarshalFlowIpv4Dscp
	phbHolder    PatternFlowIpv4DscpPhb
	ecnHolder    PatternFlowIpv4DscpEcn
}

func NewFlowIpv4Dscp() FlowIpv4Dscp {
	obj := flowIpv4Dscp{obj: &otg.FlowIpv4Dscp{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4Dscp) msg() *otg.FlowIpv4Dscp {
	return obj.obj
}

func (obj *flowIpv4Dscp) setMsg(msg *otg.FlowIpv4Dscp) FlowIpv4Dscp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4Dscp struct {
	obj *flowIpv4Dscp
}

type marshalFlowIpv4Dscp interface {
	// ToProto marshals FlowIpv4Dscp to protobuf object *otg.FlowIpv4Dscp
	ToProto() (*otg.FlowIpv4Dscp, error)
	// ToPbText marshals FlowIpv4Dscp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4Dscp to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4Dscp to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4Dscp struct {
	obj *flowIpv4Dscp
}

type unMarshalFlowIpv4Dscp interface {
	// FromProto unmarshals FlowIpv4Dscp from protobuf object *otg.FlowIpv4Dscp
	FromProto(msg *otg.FlowIpv4Dscp) (FlowIpv4Dscp, error)
	// FromPbText unmarshals FlowIpv4Dscp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4Dscp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4Dscp from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4Dscp) Marshal() marshalFlowIpv4Dscp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4Dscp{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4Dscp) Unmarshal() unMarshalFlowIpv4Dscp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4Dscp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4Dscp) ToProto() (*otg.FlowIpv4Dscp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4Dscp) FromProto(msg *otg.FlowIpv4Dscp) (FlowIpv4Dscp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4Dscp) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4Dscp) FromPbText(value string) error {
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

func (m *marshalflowIpv4Dscp) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4Dscp) FromYaml(value string) error {
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

func (m *marshalflowIpv4Dscp) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4Dscp) FromJson(value string) error {
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

func (obj *flowIpv4Dscp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4Dscp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4Dscp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4Dscp) Clone() (FlowIpv4Dscp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4Dscp()
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

func (obj *flowIpv4Dscp) setNil() {
	obj.phbHolder = nil
	obj.ecnHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv4Dscp is differentiated services code point (DSCP) packet field.
type FlowIpv4Dscp interface {
	Validation
	// msg marshals FlowIpv4Dscp to protobuf object *otg.FlowIpv4Dscp
	// and doesn't set defaults
	msg() *otg.FlowIpv4Dscp
	// setMsg unmarshals FlowIpv4Dscp from protobuf object *otg.FlowIpv4Dscp
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4Dscp) FlowIpv4Dscp
	// provides marshal interface
	Marshal() marshalFlowIpv4Dscp
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4Dscp
	// validate validates FlowIpv4Dscp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4Dscp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Phb returns PatternFlowIpv4DscpPhb, set in FlowIpv4Dscp.
	// PatternFlowIpv4DscpPhb is per hop behavior
	Phb() PatternFlowIpv4DscpPhb
	// SetPhb assigns PatternFlowIpv4DscpPhb provided by user to FlowIpv4Dscp.
	// PatternFlowIpv4DscpPhb is per hop behavior
	SetPhb(value PatternFlowIpv4DscpPhb) FlowIpv4Dscp
	// HasPhb checks if Phb has been set in FlowIpv4Dscp
	HasPhb() bool
	// Ecn returns PatternFlowIpv4DscpEcn, set in FlowIpv4Dscp.
	// PatternFlowIpv4DscpEcn is explicit congestion notification
	Ecn() PatternFlowIpv4DscpEcn
	// SetEcn assigns PatternFlowIpv4DscpEcn provided by user to FlowIpv4Dscp.
	// PatternFlowIpv4DscpEcn is explicit congestion notification
	SetEcn(value PatternFlowIpv4DscpEcn) FlowIpv4Dscp
	// HasEcn checks if Ecn has been set in FlowIpv4Dscp
	HasEcn() bool
	setNil()
}

// description is TBD
// Phb returns a PatternFlowIpv4DscpPhb
func (obj *flowIpv4Dscp) Phb() PatternFlowIpv4DscpPhb {
	if obj.obj.Phb == nil {
		obj.obj.Phb = NewPatternFlowIpv4DscpPhb().msg()
	}
	if obj.phbHolder == nil {
		obj.phbHolder = &patternFlowIpv4DscpPhb{obj: obj.obj.Phb}
	}
	return obj.phbHolder
}

// description is TBD
// Phb returns a PatternFlowIpv4DscpPhb
func (obj *flowIpv4Dscp) HasPhb() bool {
	return obj.obj.Phb != nil
}

// description is TBD
// SetPhb sets the PatternFlowIpv4DscpPhb value in the FlowIpv4Dscp object
func (obj *flowIpv4Dscp) SetPhb(value PatternFlowIpv4DscpPhb) FlowIpv4Dscp {

	obj.phbHolder = nil
	obj.obj.Phb = value.msg()

	return obj
}

// description is TBD
// Ecn returns a PatternFlowIpv4DscpEcn
func (obj *flowIpv4Dscp) Ecn() PatternFlowIpv4DscpEcn {
	if obj.obj.Ecn == nil {
		obj.obj.Ecn = NewPatternFlowIpv4DscpEcn().msg()
	}
	if obj.ecnHolder == nil {
		obj.ecnHolder = &patternFlowIpv4DscpEcn{obj: obj.obj.Ecn}
	}
	return obj.ecnHolder
}

// description is TBD
// Ecn returns a PatternFlowIpv4DscpEcn
func (obj *flowIpv4Dscp) HasEcn() bool {
	return obj.obj.Ecn != nil
}

// description is TBD
// SetEcn sets the PatternFlowIpv4DscpEcn value in the FlowIpv4Dscp object
func (obj *flowIpv4Dscp) SetEcn(value PatternFlowIpv4DscpEcn) FlowIpv4Dscp {

	obj.ecnHolder = nil
	obj.obj.Ecn = value.msg()

	return obj
}

func (obj *flowIpv4Dscp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Phb != nil {

		obj.Phb().validateObj(vObj, set_default)
	}

	if obj.obj.Ecn != nil {

		obj.Ecn().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv4Dscp) setDefault() {

}
