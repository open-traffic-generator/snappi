package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingTlvEgressNode *****
type flowIpv6SegmentRoutingTlvEgressNode struct {
	validation
	obj            *otg.FlowIpv6SegmentRoutingTlvEgressNode
	marshaller     marshalFlowIpv6SegmentRoutingTlvEgressNode
	unMarshaller   unMarshalFlowIpv6SegmentRoutingTlvEgressNode
	reservedHolder PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	flagsHolder    PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	valueHolder    PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
}

func NewFlowIpv6SegmentRoutingTlvEgressNode() FlowIpv6SegmentRoutingTlvEgressNode {
	obj := flowIpv6SegmentRoutingTlvEgressNode{obj: &otg.FlowIpv6SegmentRoutingTlvEgressNode{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingTlvEgressNode) msg() *otg.FlowIpv6SegmentRoutingTlvEgressNode {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingTlvEgressNode) setMsg(msg *otg.FlowIpv6SegmentRoutingTlvEgressNode) FlowIpv6SegmentRoutingTlvEgressNode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingTlvEgressNode struct {
	obj *flowIpv6SegmentRoutingTlvEgressNode
}

type marshalFlowIpv6SegmentRoutingTlvEgressNode interface {
	// ToProto marshals FlowIpv6SegmentRoutingTlvEgressNode to protobuf object *otg.FlowIpv6SegmentRoutingTlvEgressNode
	ToProto() (*otg.FlowIpv6SegmentRoutingTlvEgressNode, error)
	// ToPbText marshals FlowIpv6SegmentRoutingTlvEgressNode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingTlvEgressNode to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingTlvEgressNode to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingTlvEgressNode struct {
	obj *flowIpv6SegmentRoutingTlvEgressNode
}

type unMarshalFlowIpv6SegmentRoutingTlvEgressNode interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingTlvEgressNode from protobuf object *otg.FlowIpv6SegmentRoutingTlvEgressNode
	FromProto(msg *otg.FlowIpv6SegmentRoutingTlvEgressNode) (FlowIpv6SegmentRoutingTlvEgressNode, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingTlvEgressNode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingTlvEgressNode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingTlvEgressNode from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingTlvEgressNode) Marshal() marshalFlowIpv6SegmentRoutingTlvEgressNode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingTlvEgressNode{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingTlvEgressNode) Unmarshal() unMarshalFlowIpv6SegmentRoutingTlvEgressNode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingTlvEgressNode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingTlvEgressNode) ToProto() (*otg.FlowIpv6SegmentRoutingTlvEgressNode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingTlvEgressNode) FromProto(msg *otg.FlowIpv6SegmentRoutingTlvEgressNode) (FlowIpv6SegmentRoutingTlvEgressNode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingTlvEgressNode) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvEgressNode) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingTlvEgressNode) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvEgressNode) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingTlvEgressNode) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvEgressNode) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingTlvEgressNode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingTlvEgressNode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingTlvEgressNode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingTlvEgressNode) Clone() (FlowIpv6SegmentRoutingTlvEgressNode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingTlvEgressNode()
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

func (obj *flowIpv6SegmentRoutingTlvEgressNode) setNil() {
	obj.reservedHolder = nil
	obj.flagsHolder = nil
	obj.valueHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRoutingTlvEgressNode is a TLV that identifies the egress node of the SR policy.
type FlowIpv6SegmentRoutingTlvEgressNode interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingTlvEgressNode to protobuf object *otg.FlowIpv6SegmentRoutingTlvEgressNode
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingTlvEgressNode
	// setMsg unmarshals FlowIpv6SegmentRoutingTlvEgressNode from protobuf object *otg.FlowIpv6SegmentRoutingTlvEgressNode
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingTlvEgressNode) FlowIpv6SegmentRoutingTlvEgressNode
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingTlvEgressNode
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingTlvEgressNode
	// validate validates FlowIpv6SegmentRoutingTlvEgressNode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingTlvEgressNode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Reserved returns PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved, set in FlowIpv6SegmentRoutingTlvEgressNode.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved is reserved field. Must be 0.
	Reserved() PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
	// SetReserved assigns PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved provided by user to FlowIpv6SegmentRoutingTlvEgressNode.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved is reserved field. Must be 0.
	SetReserved(value PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) FlowIpv6SegmentRoutingTlvEgressNode
	// HasReserved checks if Reserved has been set in FlowIpv6SegmentRoutingTlvEgressNode
	HasReserved() bool
	// Flags returns PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags, set in FlowIpv6SegmentRoutingTlvEgressNode.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags is 8-bit flags field for this TLV.
	Flags() PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
	// SetFlags assigns PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags provided by user to FlowIpv6SegmentRoutingTlvEgressNode.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags is 8-bit flags field for this TLV.
	SetFlags(value PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) FlowIpv6SegmentRoutingTlvEgressNode
	// HasFlags checks if Flags has been set in FlowIpv6SegmentRoutingTlvEgressNode
	HasFlags() bool
	// Value returns PatternFlowIpv6SegmentRoutingTlvEgressNodeValue, set in FlowIpv6SegmentRoutingTlvEgressNode.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeValue is the 128-bit IPv6 address of the egress node.
	Value() PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
	// SetValue assigns PatternFlowIpv6SegmentRoutingTlvEgressNodeValue provided by user to FlowIpv6SegmentRoutingTlvEgressNode.
	// PatternFlowIpv6SegmentRoutingTlvEgressNodeValue is the 128-bit IPv6 address of the egress node.
	SetValue(value PatternFlowIpv6SegmentRoutingTlvEgressNodeValue) FlowIpv6SegmentRoutingTlvEgressNode
	// HasValue checks if Value has been set in FlowIpv6SegmentRoutingTlvEgressNode
	HasValue() bool
	setNil()
}

// description is TBD
// Reserved returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
func (obj *flowIpv6SegmentRoutingTlvEgressNode) Reserved() PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowIpv6SegmentRoutingTlvEgressNodeReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowIpv6SegmentRoutingTlvEgressNodeReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved
func (obj *flowIpv6SegmentRoutingTlvEgressNode) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved value in the FlowIpv6SegmentRoutingTlvEgressNode object
func (obj *flowIpv6SegmentRoutingTlvEgressNode) SetReserved(value PatternFlowIpv6SegmentRoutingTlvEgressNodeReserved) FlowIpv6SegmentRoutingTlvEgressNode {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// Flags returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
func (obj *flowIpv6SegmentRoutingTlvEgressNode) Flags() PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewPatternFlowIpv6SegmentRoutingTlvEgressNodeFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &patternFlowIpv6SegmentRoutingTlvEgressNodeFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// description is TBD
// Flags returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags
func (obj *flowIpv6SegmentRoutingTlvEgressNode) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags value in the FlowIpv6SegmentRoutingTlvEgressNode object
func (obj *flowIpv6SegmentRoutingTlvEgressNode) SetFlags(value PatternFlowIpv6SegmentRoutingTlvEgressNodeFlags) FlowIpv6SegmentRoutingTlvEgressNode {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// description is TBD
// Value returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
func (obj *flowIpv6SegmentRoutingTlvEgressNode) Value() PatternFlowIpv6SegmentRoutingTlvEgressNodeValue {
	if obj.obj.Value == nil {
		obj.obj.Value = NewPatternFlowIpv6SegmentRoutingTlvEgressNodeValue().msg()
	}
	if obj.valueHolder == nil {
		obj.valueHolder = &patternFlowIpv6SegmentRoutingTlvEgressNodeValue{obj: obj.obj.Value}
	}
	return obj.valueHolder
}

// description is TBD
// Value returns a PatternFlowIpv6SegmentRoutingTlvEgressNodeValue
func (obj *flowIpv6SegmentRoutingTlvEgressNode) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the PatternFlowIpv6SegmentRoutingTlvEgressNodeValue value in the FlowIpv6SegmentRoutingTlvEgressNode object
func (obj *flowIpv6SegmentRoutingTlvEgressNode) SetValue(value PatternFlowIpv6SegmentRoutingTlvEgressNodeValue) FlowIpv6SegmentRoutingTlvEgressNode {

	obj.valueHolder = nil
	obj.obj.Value = value.msg()

	return obj
}

func (obj *flowIpv6SegmentRoutingTlvEgressNode) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowIpv6SegmentRoutingTlvEgressNode) setDefault() {

}
