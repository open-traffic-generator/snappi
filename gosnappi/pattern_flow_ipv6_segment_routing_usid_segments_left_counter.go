package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter *****
type patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
}

func NewPatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter() PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter {
	obj := patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter{obj: &otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
}

type marshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) (PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) (PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter()
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

// PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	// validate validates PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) setDefault() {
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
