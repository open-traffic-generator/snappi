package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingSegment *****
type flowIpv6SegmentRoutingSegment struct {
	validation
	obj                 *otg.FlowIpv6SegmentRoutingSegment
	marshaller          marshalFlowIpv6SegmentRoutingSegment
	unMarshaller        unMarshalFlowIpv6SegmentRoutingSegment
	segmentHolder       PatternFlowIpv6SegmentRoutingSegmentSegment
	usidContainerHolder FlowIpv6SegmentRoutingUSidContainer
}

func NewFlowIpv6SegmentRoutingSegment() FlowIpv6SegmentRoutingSegment {
	obj := flowIpv6SegmentRoutingSegment{obj: &otg.FlowIpv6SegmentRoutingSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingSegment) msg() *otg.FlowIpv6SegmentRoutingSegment {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingSegment) setMsg(msg *otg.FlowIpv6SegmentRoutingSegment) FlowIpv6SegmentRoutingSegment {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingSegment struct {
	obj *flowIpv6SegmentRoutingSegment
}

type marshalFlowIpv6SegmentRoutingSegment interface {
	// ToProto marshals FlowIpv6SegmentRoutingSegment to protobuf object *otg.FlowIpv6SegmentRoutingSegment
	ToProto() (*otg.FlowIpv6SegmentRoutingSegment, error)
	// ToPbText marshals FlowIpv6SegmentRoutingSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingSegment to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingSegment struct {
	obj *flowIpv6SegmentRoutingSegment
}

type unMarshalFlowIpv6SegmentRoutingSegment interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingSegment from protobuf object *otg.FlowIpv6SegmentRoutingSegment
	FromProto(msg *otg.FlowIpv6SegmentRoutingSegment) (FlowIpv6SegmentRoutingSegment, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingSegment from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingSegment) Marshal() marshalFlowIpv6SegmentRoutingSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingSegment) Unmarshal() unMarshalFlowIpv6SegmentRoutingSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingSegment) ToProto() (*otg.FlowIpv6SegmentRoutingSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingSegment) FromProto(msg *otg.FlowIpv6SegmentRoutingSegment) (FlowIpv6SegmentRoutingSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingSegment) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingSegment) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingSegment) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingSegment) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingSegment) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingSegment) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingSegment) Clone() (FlowIpv6SegmentRoutingSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingSegment()
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

func (obj *flowIpv6SegmentRoutingSegment) setNil() {
	obj.segmentHolder = nil
	obj.usidContainerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRoutingSegment is one entry in the SRH Segment List (RFC 8754 Section 2.1). Two encodings are supported: 'segment' - a single 128-bit SRv6 SID address (full-SID encoding, default). 'usid_container' - a packed 128-bit uSID container holding multiple micro-SIDs per RFC 9800. Set 'usid_container' to provide structured input; leave it unset and use 'segment' for standard full-SID entries.
type FlowIpv6SegmentRoutingSegment interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingSegment to protobuf object *otg.FlowIpv6SegmentRoutingSegment
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingSegment
	// setMsg unmarshals FlowIpv6SegmentRoutingSegment from protobuf object *otg.FlowIpv6SegmentRoutingSegment
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingSegment) FlowIpv6SegmentRoutingSegment
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingSegment
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingSegment
	// validate validates FlowIpv6SegmentRoutingSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Segment returns PatternFlowIpv6SegmentRoutingSegmentSegment, set in FlowIpv6SegmentRoutingSegment.
	// PatternFlowIpv6SegmentRoutingSegmentSegment is a 128-bit IPv6 SID address (RFC 8754 Section 2.1). Use for standard full-SID encoding where each segment carries a complete locator + function + argument value. Example: "2001:db8::1".
	Segment() PatternFlowIpv6SegmentRoutingSegmentSegment
	// SetSegment assigns PatternFlowIpv6SegmentRoutingSegmentSegment provided by user to FlowIpv6SegmentRoutingSegment.
	// PatternFlowIpv6SegmentRoutingSegmentSegment is a 128-bit IPv6 SID address (RFC 8754 Section 2.1). Use for standard full-SID encoding where each segment carries a complete locator + function + argument value. Example: "2001:db8::1".
	SetSegment(value PatternFlowIpv6SegmentRoutingSegmentSegment) FlowIpv6SegmentRoutingSegment
	// HasSegment checks if Segment has been set in FlowIpv6SegmentRoutingSegment
	HasSegment() bool
	// UsidContainer returns FlowIpv6SegmentRoutingUSidContainer, set in FlowIpv6SegmentRoutingSegment.
	// FlowIpv6SegmentRoutingUSidContainer is a uSID container segment (RFC 9800) - a 128-bit IPv6 address packed with multiple Micro-SIDs (uSIDs). The implementation assembles the 128-bit value by placing the locator block prefix in the high bits, then packing each uSID (locator_node + function) consecutively in units of (locator_node_length + function_length) bits, padding the remainder with zeros. Example using F3216 (lb=32, ln=16, fn=16, arg=0) and locator block fc00::/32:
	// usids = [{node: "0001", function: "0001"}, {node: "0002", function: "0001"}]
	// => container = fc00:0:1:1:2:1:: (2-hop path, nodes 1 and 2, function 1)
	// Reference: RFC 9800 Section 4.
	UsidContainer() FlowIpv6SegmentRoutingUSidContainer
	// SetUsidContainer assigns FlowIpv6SegmentRoutingUSidContainer provided by user to FlowIpv6SegmentRoutingSegment.
	// FlowIpv6SegmentRoutingUSidContainer is a uSID container segment (RFC 9800) - a 128-bit IPv6 address packed with multiple Micro-SIDs (uSIDs). The implementation assembles the 128-bit value by placing the locator block prefix in the high bits, then packing each uSID (locator_node + function) consecutively in units of (locator_node_length + function_length) bits, padding the remainder with zeros. Example using F3216 (lb=32, ln=16, fn=16, arg=0) and locator block fc00::/32:
	// usids = [{node: "0001", function: "0001"}, {node: "0002", function: "0001"}]
	// => container = fc00:0:1:1:2:1:: (2-hop path, nodes 1 and 2, function 1)
	// Reference: RFC 9800 Section 4.
	SetUsidContainer(value FlowIpv6SegmentRoutingUSidContainer) FlowIpv6SegmentRoutingSegment
	// HasUsidContainer checks if UsidContainer has been set in FlowIpv6SegmentRoutingSegment
	HasUsidContainer() bool
	setNil()
}

// description is TBD
// Segment returns a PatternFlowIpv6SegmentRoutingSegmentSegment
func (obj *flowIpv6SegmentRoutingSegment) Segment() PatternFlowIpv6SegmentRoutingSegmentSegment {
	if obj.obj.Segment == nil {
		obj.obj.Segment = NewPatternFlowIpv6SegmentRoutingSegmentSegment().msg()
	}
	if obj.segmentHolder == nil {
		obj.segmentHolder = &patternFlowIpv6SegmentRoutingSegmentSegment{obj: obj.obj.Segment}
	}
	return obj.segmentHolder
}

// description is TBD
// Segment returns a PatternFlowIpv6SegmentRoutingSegmentSegment
func (obj *flowIpv6SegmentRoutingSegment) HasSegment() bool {
	return obj.obj.Segment != nil
}

// description is TBD
// SetSegment sets the PatternFlowIpv6SegmentRoutingSegmentSegment value in the FlowIpv6SegmentRoutingSegment object
func (obj *flowIpv6SegmentRoutingSegment) SetSegment(value PatternFlowIpv6SegmentRoutingSegmentSegment) FlowIpv6SegmentRoutingSegment {

	obj.segmentHolder = nil
	obj.obj.Segment = value.msg()

	return obj
}

// Structured uSID container encoding (RFC 9800). When set, the implementation assembles the 128-bit container value from the constituent uSIDs and places it in the SRH segment list entry. Takes precedence over 'segment' if both are provided.
// UsidContainer returns a FlowIpv6SegmentRoutingUSidContainer
func (obj *flowIpv6SegmentRoutingSegment) UsidContainer() FlowIpv6SegmentRoutingUSidContainer {
	if obj.obj.UsidContainer == nil {
		obj.obj.UsidContainer = NewFlowIpv6SegmentRoutingUSidContainer().msg()
	}
	if obj.usidContainerHolder == nil {
		obj.usidContainerHolder = &flowIpv6SegmentRoutingUSidContainer{obj: obj.obj.UsidContainer}
	}
	return obj.usidContainerHolder
}

// Structured uSID container encoding (RFC 9800). When set, the implementation assembles the 128-bit container value from the constituent uSIDs and places it in the SRH segment list entry. Takes precedence over 'segment' if both are provided.
// UsidContainer returns a FlowIpv6SegmentRoutingUSidContainer
func (obj *flowIpv6SegmentRoutingSegment) HasUsidContainer() bool {
	return obj.obj.UsidContainer != nil
}

// Structured uSID container encoding (RFC 9800). When set, the implementation assembles the 128-bit container value from the constituent uSIDs and places it in the SRH segment list entry. Takes precedence over 'segment' if both are provided.
// SetUsidContainer sets the FlowIpv6SegmentRoutingUSidContainer value in the FlowIpv6SegmentRoutingSegment object
func (obj *flowIpv6SegmentRoutingSegment) SetUsidContainer(value FlowIpv6SegmentRoutingUSidContainer) FlowIpv6SegmentRoutingSegment {

	obj.usidContainerHolder = nil
	obj.obj.UsidContainer = value.msg()

	return obj
}

func (obj *flowIpv6SegmentRoutingSegment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Segment != nil {

		obj.Segment().validateObj(vObj, set_default)
	}

	if obj.obj.UsidContainer != nil {

		obj.UsidContainer().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6SegmentRoutingSegment) setDefault() {

}
