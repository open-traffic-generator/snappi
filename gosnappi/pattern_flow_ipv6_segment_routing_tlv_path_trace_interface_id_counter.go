package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter *****
type patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter {
	obj := patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter{obj: &otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
}

type marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter()
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

// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	// validate validates PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter.Start <= 4095 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter.Step <= 4095 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter.Count <= 4095 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) setDefault() {
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
