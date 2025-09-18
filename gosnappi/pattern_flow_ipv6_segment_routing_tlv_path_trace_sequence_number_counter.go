package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter *****
type patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter() PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter {
	obj := patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter{obj: &otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
}

type marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) (PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) (PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter()
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

// PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	// validate validates PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) setDefault() {
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
