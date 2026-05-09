package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsHFlagCounter *****
type patternFlowIpv6SegmentRoutingFlagsHFlagCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingFlagsHFlagCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsHFlagCounter() PatternFlowIpv6SegmentRoutingFlagsHFlagCounter {
	obj := patternFlowIpv6SegmentRoutingFlagsHFlagCounter{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter) PatternFlowIpv6SegmentRoutingFlagsHFlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsHFlagCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter
}

type marshalPatternFlowIpv6SegmentRoutingFlagsHFlagCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsHFlagCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsHFlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsHFlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsHFlagCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsHFlagCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsHFlagCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsHFlagCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter) (PatternFlowIpv6SegmentRoutingFlagsHFlagCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsHFlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsHFlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsHFlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsHFlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsHFlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsHFlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsHFlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHFlagCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHFlagCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter) (PatternFlowIpv6SegmentRoutingFlagsHFlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHFlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHFlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHFlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHFlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHFlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHFlagCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) Clone() (PatternFlowIpv6SegmentRoutingFlagsHFlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsHFlagCounter()
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

// PatternFlowIpv6SegmentRoutingFlagsHFlagCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingFlagsHFlagCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsHFlagCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsHFlagCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsHFlagCounter) PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	// validate validates PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsHFlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsHFlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsHFlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsHFlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsHFlagCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsHFlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsHFlagCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsHFlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsHFlagCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsHFlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsHFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsHFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsHFlagCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlagCounter) setDefault() {
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
