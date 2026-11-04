package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingUsidLastEntryCounter *****
type patternFlowIpv6SegmentRoutingUsidLastEntryCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingUsidLastEntryCounter
}

func NewPatternFlowIpv6SegmentRoutingUsidLastEntryCounter() PatternFlowIpv6SegmentRoutingUsidLastEntryCounter {
	obj := patternFlowIpv6SegmentRoutingUsidLastEntryCounter{obj: &otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) msg() *otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter) PatternFlowIpv6SegmentRoutingUsidLastEntryCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingUsidLastEntryCounter struct {
	obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter
}

type marshalPatternFlowIpv6SegmentRoutingUsidLastEntryCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingUsidLastEntryCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingUsidLastEntryCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingUsidLastEntryCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingUsidLastEntryCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingUsidLastEntryCounter struct {
	obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingUsidLastEntryCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingUsidLastEntryCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter) (PatternFlowIpv6SegmentRoutingUsidLastEntryCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingUsidLastEntryCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingUsidLastEntryCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingUsidLastEntryCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingUsidLastEntryCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingUsidLastEntryCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidLastEntryCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingUsidLastEntryCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidLastEntryCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidLastEntryCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter) (PatternFlowIpv6SegmentRoutingUsidLastEntryCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidLastEntryCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidLastEntryCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidLastEntryCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidLastEntryCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidLastEntryCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidLastEntryCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) Clone() (PatternFlowIpv6SegmentRoutingUsidLastEntryCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingUsidLastEntryCounter()
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

// PatternFlowIpv6SegmentRoutingUsidLastEntryCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingUsidLastEntryCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingUsidLastEntryCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingUsidLastEntryCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingUsidLastEntryCounter) PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	// validate validates PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingUsidLastEntryCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingUsidLastEntryCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingUsidLastEntryCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingUsidLastEntryCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidLastEntryCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingUsidLastEntryCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidLastEntryCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingUsidLastEntryCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidLastEntryCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingUsidLastEntryCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidLastEntryCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidLastEntryCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidLastEntryCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntryCounter) setDefault() {
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
