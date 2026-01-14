package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingSegmentsLeftCounter *****
type patternFlowIpv6SegmentRoutingSegmentsLeftCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingSegmentsLeftCounter
}

func NewPatternFlowIpv6SegmentRoutingSegmentsLeftCounter() PatternFlowIpv6SegmentRoutingSegmentsLeftCounter {
	obj := patternFlowIpv6SegmentRoutingSegmentsLeftCounter{obj: &otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) msg() *otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter) PatternFlowIpv6SegmentRoutingSegmentsLeftCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingSegmentsLeftCounter struct {
	obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter
}

type marshalPatternFlowIpv6SegmentRoutingSegmentsLeftCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingSegmentsLeftCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingSegmentsLeftCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingSegmentsLeftCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingSegmentsLeftCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingSegmentsLeftCounter struct {
	obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingSegmentsLeftCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingSegmentsLeftCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter) (PatternFlowIpv6SegmentRoutingSegmentsLeftCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingSegmentsLeftCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingSegmentsLeftCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingSegmentsLeftCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingSegmentsLeftCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingSegmentsLeftCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingSegmentsLeftCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingSegmentsLeftCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentsLeftCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentsLeftCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter) (PatternFlowIpv6SegmentRoutingSegmentsLeftCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentsLeftCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentsLeftCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentsLeftCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentsLeftCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentsLeftCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentsLeftCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) Clone() (PatternFlowIpv6SegmentRoutingSegmentsLeftCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingSegmentsLeftCounter()
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

// PatternFlowIpv6SegmentRoutingSegmentsLeftCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingSegmentsLeftCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingSegmentsLeftCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingSegmentsLeftCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingSegmentsLeftCounter) PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	// validate validates PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingSegmentsLeftCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingSegmentsLeftCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingSegmentsLeftCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingSegmentsLeftCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingSegmentsLeftCounter object
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingSegmentsLeftCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingSegmentsLeftCounter object
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingSegmentsLeftCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingSegmentsLeftCounter object
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingSegmentsLeftCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingSegmentsLeftCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingSegmentsLeftCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingSegmentsLeftCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeftCounter) setDefault() {
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
