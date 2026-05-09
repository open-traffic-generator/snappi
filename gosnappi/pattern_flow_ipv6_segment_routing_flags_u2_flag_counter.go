package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter *****
type patternFlowIpv6SegmentRoutingFlagsU2FlagCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsU2FlagCounter() PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter {
	obj := patternFlowIpv6SegmentRoutingFlagsU2FlagCounter{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsU2FlagCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter
}

type marshalPatternFlowIpv6SegmentRoutingFlagsU2FlagCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsU2FlagCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsU2FlagCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) (PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsU2FlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsU2FlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsU2FlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsU2FlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) (PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) Clone() (PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsU2FlagCounter()
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

// PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter) PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	// validate validates PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsU2FlagCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingFlagsU2FlagCounter) setDefault() {
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
