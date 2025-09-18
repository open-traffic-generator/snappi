package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingTlvIngressNode *****
type flowIpv6SegmentRoutingTlvIngressNode struct {
	validation
	obj            *otg.FlowIpv6SegmentRoutingTlvIngressNode
	marshaller     marshalFlowIpv6SegmentRoutingTlvIngressNode
	unMarshaller   unMarshalFlowIpv6SegmentRoutingTlvIngressNode
	reservedHolder PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	flagsHolder    PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	valueHolder    PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
}

func NewFlowIpv6SegmentRoutingTlvIngressNode() FlowIpv6SegmentRoutingTlvIngressNode {
	obj := flowIpv6SegmentRoutingTlvIngressNode{obj: &otg.FlowIpv6SegmentRoutingTlvIngressNode{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingTlvIngressNode) msg() *otg.FlowIpv6SegmentRoutingTlvIngressNode {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingTlvIngressNode) setMsg(msg *otg.FlowIpv6SegmentRoutingTlvIngressNode) FlowIpv6SegmentRoutingTlvIngressNode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingTlvIngressNode struct {
	obj *flowIpv6SegmentRoutingTlvIngressNode
}

type marshalFlowIpv6SegmentRoutingTlvIngressNode interface {
	// ToProto marshals FlowIpv6SegmentRoutingTlvIngressNode to protobuf object *otg.FlowIpv6SegmentRoutingTlvIngressNode
	ToProto() (*otg.FlowIpv6SegmentRoutingTlvIngressNode, error)
	// ToPbText marshals FlowIpv6SegmentRoutingTlvIngressNode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingTlvIngressNode to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingTlvIngressNode to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingTlvIngressNode struct {
	obj *flowIpv6SegmentRoutingTlvIngressNode
}

type unMarshalFlowIpv6SegmentRoutingTlvIngressNode interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingTlvIngressNode from protobuf object *otg.FlowIpv6SegmentRoutingTlvIngressNode
	FromProto(msg *otg.FlowIpv6SegmentRoutingTlvIngressNode) (FlowIpv6SegmentRoutingTlvIngressNode, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingTlvIngressNode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingTlvIngressNode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingTlvIngressNode from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingTlvIngressNode) Marshal() marshalFlowIpv6SegmentRoutingTlvIngressNode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingTlvIngressNode{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingTlvIngressNode) Unmarshal() unMarshalFlowIpv6SegmentRoutingTlvIngressNode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingTlvIngressNode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingTlvIngressNode) ToProto() (*otg.FlowIpv6SegmentRoutingTlvIngressNode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingTlvIngressNode) FromProto(msg *otg.FlowIpv6SegmentRoutingTlvIngressNode) (FlowIpv6SegmentRoutingTlvIngressNode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingTlvIngressNode) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvIngressNode) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingTlvIngressNode) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvIngressNode) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingTlvIngressNode) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvIngressNode) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingTlvIngressNode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingTlvIngressNode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingTlvIngressNode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingTlvIngressNode) Clone() (FlowIpv6SegmentRoutingTlvIngressNode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingTlvIngressNode()
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

func (obj *flowIpv6SegmentRoutingTlvIngressNode) setNil() {
	obj.reservedHolder = nil
	obj.flagsHolder = nil
	obj.valueHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRoutingTlvIngressNode is a TLV that may be used to carry the IPv6 address of the ingress node of the SR domain.
type FlowIpv6SegmentRoutingTlvIngressNode interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingTlvIngressNode to protobuf object *otg.FlowIpv6SegmentRoutingTlvIngressNode
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingTlvIngressNode
	// setMsg unmarshals FlowIpv6SegmentRoutingTlvIngressNode from protobuf object *otg.FlowIpv6SegmentRoutingTlvIngressNode
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingTlvIngressNode) FlowIpv6SegmentRoutingTlvIngressNode
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingTlvIngressNode
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingTlvIngressNode
	// validate validates FlowIpv6SegmentRoutingTlvIngressNode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingTlvIngressNode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Reserved returns PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved, set in FlowIpv6SegmentRoutingTlvIngressNode.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved is reserved field. Must be 0.
	Reserved() PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
	// SetReserved assigns PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved provided by user to FlowIpv6SegmentRoutingTlvIngressNode.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved is reserved field. Must be 0.
	SetReserved(value PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) FlowIpv6SegmentRoutingTlvIngressNode
	// HasReserved checks if Reserved has been set in FlowIpv6SegmentRoutingTlvIngressNode
	HasReserved() bool
	// Flags returns PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags, set in FlowIpv6SegmentRoutingTlvIngressNode.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags is 8-bit flags field for this TLV.
	Flags() PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
	// SetFlags assigns PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags provided by user to FlowIpv6SegmentRoutingTlvIngressNode.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags is 8-bit flags field for this TLV.
	SetFlags(value PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) FlowIpv6SegmentRoutingTlvIngressNode
	// HasFlags checks if Flags has been set in FlowIpv6SegmentRoutingTlvIngressNode
	HasFlags() bool
	// Value returns PatternFlowIpv6SegmentRoutingTlvIngressNodeValue, set in FlowIpv6SegmentRoutingTlvIngressNode.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeValue is the 128-bit IPv6 address of the ingress node.
	Value() PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
	// SetValue assigns PatternFlowIpv6SegmentRoutingTlvIngressNodeValue provided by user to FlowIpv6SegmentRoutingTlvIngressNode.
	// PatternFlowIpv6SegmentRoutingTlvIngressNodeValue is the 128-bit IPv6 address of the ingress node.
	SetValue(value PatternFlowIpv6SegmentRoutingTlvIngressNodeValue) FlowIpv6SegmentRoutingTlvIngressNode
	// HasValue checks if Value has been set in FlowIpv6SegmentRoutingTlvIngressNode
	HasValue() bool
	setNil()
}

// description is TBD
// Reserved returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
func (obj *flowIpv6SegmentRoutingTlvIngressNode) Reserved() PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowIpv6SegmentRoutingTlvIngressNodeReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowIpv6SegmentRoutingTlvIngressNodeReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved
func (obj *flowIpv6SegmentRoutingTlvIngressNode) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved value in the FlowIpv6SegmentRoutingTlvIngressNode object
func (obj *flowIpv6SegmentRoutingTlvIngressNode) SetReserved(value PatternFlowIpv6SegmentRoutingTlvIngressNodeReserved) FlowIpv6SegmentRoutingTlvIngressNode {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// Flags returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
func (obj *flowIpv6SegmentRoutingTlvIngressNode) Flags() PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewPatternFlowIpv6SegmentRoutingTlvIngressNodeFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &patternFlowIpv6SegmentRoutingTlvIngressNodeFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// description is TBD
// Flags returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags
func (obj *flowIpv6SegmentRoutingTlvIngressNode) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags value in the FlowIpv6SegmentRoutingTlvIngressNode object
func (obj *flowIpv6SegmentRoutingTlvIngressNode) SetFlags(value PatternFlowIpv6SegmentRoutingTlvIngressNodeFlags) FlowIpv6SegmentRoutingTlvIngressNode {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// description is TBD
// Value returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
func (obj *flowIpv6SegmentRoutingTlvIngressNode) Value() PatternFlowIpv6SegmentRoutingTlvIngressNodeValue {
	if obj.obj.Value == nil {
		obj.obj.Value = NewPatternFlowIpv6SegmentRoutingTlvIngressNodeValue().msg()
	}
	if obj.valueHolder == nil {
		obj.valueHolder = &patternFlowIpv6SegmentRoutingTlvIngressNodeValue{obj: obj.obj.Value}
	}
	return obj.valueHolder
}

// description is TBD
// Value returns a PatternFlowIpv6SegmentRoutingTlvIngressNodeValue
func (obj *flowIpv6SegmentRoutingTlvIngressNode) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the PatternFlowIpv6SegmentRoutingTlvIngressNodeValue value in the FlowIpv6SegmentRoutingTlvIngressNode object
func (obj *flowIpv6SegmentRoutingTlvIngressNode) SetValue(value PatternFlowIpv6SegmentRoutingTlvIngressNodeValue) FlowIpv6SegmentRoutingTlvIngressNode {

	obj.valueHolder = nil
	obj.obj.Value = value.msg()

	return obj
}

func (obj *flowIpv6SegmentRoutingTlvIngressNode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(vObj, set_default)
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.Value != nil {

		obj.Value().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6SegmentRoutingTlvIngressNode) setDefault() {

}
