package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter *****
type patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
}

func NewPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter {
	obj := patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter{obj: &otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
}

type marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) (PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) (PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter()
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

// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	// validate validates PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(32)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
