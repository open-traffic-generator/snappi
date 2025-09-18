package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter *****
type patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter() PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter {
	obj := patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter{obj: &otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) msg() *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
}

type marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) (PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) (PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) Clone() (PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter()
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

// PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	// validate validates PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeReservedCounter) setDefault() {
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
