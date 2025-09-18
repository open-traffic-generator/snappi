package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter *****
type patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter() PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter {
	obj := patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter{obj: &otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) msg() *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
}

type marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) (PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) (PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) Clone() (PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter()
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

// PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	// validate validates PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeFlagsCounter) setDefault() {
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
