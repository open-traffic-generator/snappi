package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingLastEntryCounter *****
type patternFlowIpv6SegmentRoutingLastEntryCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingLastEntryCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingLastEntryCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingLastEntryCounter
}

func NewPatternFlowIpv6SegmentRoutingLastEntryCounter() PatternFlowIpv6SegmentRoutingLastEntryCounter {
	obj := patternFlowIpv6SegmentRoutingLastEntryCounter{obj: &otg.PatternFlowIpv6SegmentRoutingLastEntryCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) msg() *otg.PatternFlowIpv6SegmentRoutingLastEntryCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingLastEntryCounter) PatternFlowIpv6SegmentRoutingLastEntryCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingLastEntryCounter struct {
	obj *patternFlowIpv6SegmentRoutingLastEntryCounter
}

type marshalPatternFlowIpv6SegmentRoutingLastEntryCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingLastEntryCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingLastEntryCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingLastEntryCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingLastEntryCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingLastEntryCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingLastEntryCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingLastEntryCounter struct {
	obj *patternFlowIpv6SegmentRoutingLastEntryCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingLastEntryCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingLastEntryCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingLastEntryCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingLastEntryCounter) (PatternFlowIpv6SegmentRoutingLastEntryCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingLastEntryCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingLastEntryCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingLastEntryCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingLastEntryCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingLastEntryCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingLastEntryCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingLastEntryCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingLastEntryCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingLastEntryCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingLastEntryCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingLastEntryCounter) (PatternFlowIpv6SegmentRoutingLastEntryCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingLastEntryCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingLastEntryCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingLastEntryCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingLastEntryCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingLastEntryCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingLastEntryCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) Clone() (PatternFlowIpv6SegmentRoutingLastEntryCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingLastEntryCounter()
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

// PatternFlowIpv6SegmentRoutingLastEntryCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingLastEntryCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingLastEntryCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingLastEntryCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingLastEntryCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingLastEntryCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingLastEntryCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingLastEntryCounter) PatternFlowIpv6SegmentRoutingLastEntryCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingLastEntryCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingLastEntryCounter
	// validate validates PatternFlowIpv6SegmentRoutingLastEntryCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingLastEntryCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingLastEntryCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingLastEntryCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingLastEntryCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingLastEntryCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingLastEntryCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingLastEntryCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingLastEntryCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingLastEntryCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingLastEntryCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingLastEntryCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingLastEntryCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingLastEntryCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingLastEntryCounter object
func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingLastEntryCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingLastEntryCounter object
func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingLastEntryCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingLastEntryCounter object
func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingLastEntryCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingLastEntryCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingLastEntryCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingLastEntryCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingLastEntryCounter) setDefault() {
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
