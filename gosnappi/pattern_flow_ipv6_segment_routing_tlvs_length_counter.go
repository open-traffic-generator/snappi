package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvsLengthCounter *****
type patternFlowIpv6SegmentRoutingTlvsLengthCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingTlvsLengthCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingTlvsLengthCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvsLengthCounter() PatternFlowIpv6SegmentRoutingTlvsLengthCounter {
	obj := patternFlowIpv6SegmentRoutingTlvsLengthCounter{obj: &otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) msg() *otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter) PatternFlowIpv6SegmentRoutingTlvsLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvsLengthCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter
}

type marshalPatternFlowIpv6SegmentRoutingTlvsLengthCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvsLengthCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvsLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvsLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvsLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvsLengthCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvsLengthCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvsLengthCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter) (PatternFlowIpv6SegmentRoutingTlvsLengthCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvsLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvsLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvsLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvsLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvsLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvsLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvsLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvsLengthCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvsLengthCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter) (PatternFlowIpv6SegmentRoutingTlvsLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvsLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvsLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvsLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvsLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvsLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvsLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) Clone() (PatternFlowIpv6SegmentRoutingTlvsLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvsLengthCounter()
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

// PatternFlowIpv6SegmentRoutingTlvsLengthCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingTlvsLengthCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvsLengthCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvsLengthCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvsLengthCounter) PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvsLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvsLengthCounter
	// validate validates PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvsLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingTlvsLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingTlvsLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingTlvsLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingTlvsLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvsLengthCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvsLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvsLengthCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvsLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvsLengthCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvsLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvsLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvsLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvsLengthCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvsLengthCounter) setDefault() {
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
