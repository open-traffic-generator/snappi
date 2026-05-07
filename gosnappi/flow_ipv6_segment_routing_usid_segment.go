package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingUsidSegment *****
type flowIpv6SegmentRoutingUsidSegment struct {
	validation
	obj           *otg.FlowIpv6SegmentRoutingUsidSegment
	marshaller    marshalFlowIpv6SegmentRoutingUsidSegment
	unMarshaller  unMarshalFlowIpv6SegmentRoutingUsidSegment
	segmentHolder PatternFlowIpv6SegmentRoutingUsidSegmentSegment
}

func NewFlowIpv6SegmentRoutingUsidSegment() FlowIpv6SegmentRoutingUsidSegment {
	obj := flowIpv6SegmentRoutingUsidSegment{obj: &otg.FlowIpv6SegmentRoutingUsidSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingUsidSegment) msg() *otg.FlowIpv6SegmentRoutingUsidSegment {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingUsidSegment) setMsg(msg *otg.FlowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidSegment {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingUsidSegment struct {
	obj *flowIpv6SegmentRoutingUsidSegment
}

type marshalFlowIpv6SegmentRoutingUsidSegment interface {
	// ToProto marshals FlowIpv6SegmentRoutingUsidSegment to protobuf object *otg.FlowIpv6SegmentRoutingUsidSegment
	ToProto() (*otg.FlowIpv6SegmentRoutingUsidSegment, error)
	// ToPbText marshals FlowIpv6SegmentRoutingUsidSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingUsidSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingUsidSegment to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingUsidSegment struct {
	obj *flowIpv6SegmentRoutingUsidSegment
}

type unMarshalFlowIpv6SegmentRoutingUsidSegment interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingUsidSegment from protobuf object *otg.FlowIpv6SegmentRoutingUsidSegment
	FromProto(msg *otg.FlowIpv6SegmentRoutingUsidSegment) (FlowIpv6SegmentRoutingUsidSegment, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingUsidSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingUsidSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingUsidSegment from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingUsidSegment) Marshal() marshalFlowIpv6SegmentRoutingUsidSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingUsidSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingUsidSegment) Unmarshal() unMarshalFlowIpv6SegmentRoutingUsidSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingUsidSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingUsidSegment) ToProto() (*otg.FlowIpv6SegmentRoutingUsidSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingUsidSegment) FromProto(msg *otg.FlowIpv6SegmentRoutingUsidSegment) (FlowIpv6SegmentRoutingUsidSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingUsidSegment) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsidSegment) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingUsidSegment) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsidSegment) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingUsidSegment) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsidSegment) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingUsidSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUsidSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUsidSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingUsidSegment) Clone() (FlowIpv6SegmentRoutingUsidSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingUsidSegment()
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

func (obj *flowIpv6SegmentRoutingUsidSegment) setNil() {
	obj.segmentHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRoutingUsidSegment is one pre-computed uSID container entry in the SRH Segment List (RFC 9800 Section 4, RFC 8754 Section 2.1). The 128-bit value encodes the locator block and packed micro-SIDs and is supplied as a plain IPv6 address. No decomposition metadata is carried on the wire.
// Maximum micro-SIDs per container (RFC 9800 Section 4):
// floor((128 - locator_block_length) /
// (locator_node_length + function_length))
// Common formats:
// F3216 (lb=32, ln=16, fn=16): floor(96/32) = 3 uSIDs per container.
// F4816 (lb=48, ln=16, fn=16): floor(80/32) = 2 uSIDs per container.
// F3208 (lb=32, ln=8,  fn=8 ): floor(96/16) = 6 uSIDs per container.
//
// Example (F3216, lb=fc00::/32, node2-fn1 packed with node3-fn1):
// fc00:0:2:1:3:1::
type FlowIpv6SegmentRoutingUsidSegment interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingUsidSegment to protobuf object *otg.FlowIpv6SegmentRoutingUsidSegment
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingUsidSegment
	// setMsg unmarshals FlowIpv6SegmentRoutingUsidSegment from protobuf object *otg.FlowIpv6SegmentRoutingUsidSegment
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidSegment
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingUsidSegment
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingUsidSegment
	// validate validates FlowIpv6SegmentRoutingUsidSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingUsidSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Segment returns PatternFlowIpv6SegmentRoutingUsidSegmentSegment, set in FlowIpv6SegmentRoutingUsidSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentSegment is a pre-computed 128-bit uSID container value as an IPv6 address (RFC 9800 Section 4). Example: "fc00:0:2:1:3:1::".
	Segment() PatternFlowIpv6SegmentRoutingUsidSegmentSegment
	// SetSegment assigns PatternFlowIpv6SegmentRoutingUsidSegmentSegment provided by user to FlowIpv6SegmentRoutingUsidSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentSegment is a pre-computed 128-bit uSID container value as an IPv6 address (RFC 9800 Section 4). Example: "fc00:0:2:1:3:1::".
	SetSegment(value PatternFlowIpv6SegmentRoutingUsidSegmentSegment) FlowIpv6SegmentRoutingUsidSegment
	// HasSegment checks if Segment has been set in FlowIpv6SegmentRoutingUsidSegment
	HasSegment() bool
	setNil()
}

// description is TBD
// Segment returns a PatternFlowIpv6SegmentRoutingUsidSegmentSegment
func (obj *flowIpv6SegmentRoutingUsidSegment) Segment() PatternFlowIpv6SegmentRoutingUsidSegmentSegment {
	if obj.obj.Segment == nil {
		obj.obj.Segment = NewPatternFlowIpv6SegmentRoutingUsidSegmentSegment().msg()
	}
	if obj.segmentHolder == nil {
		obj.segmentHolder = &patternFlowIpv6SegmentRoutingUsidSegmentSegment{obj: obj.obj.Segment}
	}
	return obj.segmentHolder
}

// description is TBD
// Segment returns a PatternFlowIpv6SegmentRoutingUsidSegmentSegment
func (obj *flowIpv6SegmentRoutingUsidSegment) HasSegment() bool {
	return obj.obj.Segment != nil
}

// description is TBD
// SetSegment sets the PatternFlowIpv6SegmentRoutingUsidSegmentSegment value in the FlowIpv6SegmentRoutingUsidSegment object
func (obj *flowIpv6SegmentRoutingUsidSegment) SetSegment(value PatternFlowIpv6SegmentRoutingUsidSegmentSegment) FlowIpv6SegmentRoutingUsidSegment {

	obj.segmentHolder = nil
	obj.obj.Segment = value.msg()

	return obj
}

func (obj *flowIpv6SegmentRoutingUsidSegment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Segment != nil {

		obj.Segment().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6SegmentRoutingUsidSegment) setDefault() {

}
