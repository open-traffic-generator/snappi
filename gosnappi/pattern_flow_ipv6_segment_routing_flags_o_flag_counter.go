package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsOFlagCounter *****
type patternFlowIpv6SegmentRoutingFlagsOFlagCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingFlagsOFlagCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsOFlagCounter() PatternFlowIpv6SegmentRoutingFlagsOFlagCounter {
	obj := patternFlowIpv6SegmentRoutingFlagsOFlagCounter{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter) PatternFlowIpv6SegmentRoutingFlagsOFlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsOFlagCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter
}

type marshalPatternFlowIpv6SegmentRoutingFlagsOFlagCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsOFlagCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsOFlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsOFlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsOFlagCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsOFlagCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsOFlagCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsOFlagCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter) (PatternFlowIpv6SegmentRoutingFlagsOFlagCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsOFlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsOFlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsOFlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsOFlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsOFlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsOFlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsOFlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOFlagCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOFlagCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter) (PatternFlowIpv6SegmentRoutingFlagsOFlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOFlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOFlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOFlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOFlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOFlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOFlagCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) Clone() (PatternFlowIpv6SegmentRoutingFlagsOFlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsOFlagCounter()
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

// PatternFlowIpv6SegmentRoutingFlagsOFlagCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingFlagsOFlagCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsOFlagCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsOFlagCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsOFlagCounter) PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	// validate validates PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsOFlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsOFlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsOFlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsOFlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsOFlagCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsOFlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsOFlagCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsOFlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsOFlagCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsOFlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsOFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsOFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsOFlagCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlagCounter) setDefault() {
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
