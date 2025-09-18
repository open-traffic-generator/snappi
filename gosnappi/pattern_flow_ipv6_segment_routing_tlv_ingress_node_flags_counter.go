package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter *****
type patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter() PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter {
	obj := patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter{obj: &otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) msg() *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
}

type marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) (PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) (PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) Clone() (PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter()
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

// PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	// validate validates PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeFlagsCounter) setDefault() {
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
