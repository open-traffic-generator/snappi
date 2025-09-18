package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter *****
type patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter {
	obj := patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter{obj: &otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
}

type marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter()
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

// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	// validate validates PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter.Count <= 15 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) setDefault() {
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
