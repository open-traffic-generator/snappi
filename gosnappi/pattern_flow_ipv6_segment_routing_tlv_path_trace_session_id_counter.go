package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter *****
type patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter() PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter {
	obj := patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter{obj: &otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
}

type marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) (PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) (PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter()
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

// PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	// validate validates PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(0)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
