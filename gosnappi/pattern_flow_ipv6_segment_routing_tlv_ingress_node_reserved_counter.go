package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter *****
type patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter() PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter {
	obj := patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter{obj: &otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) msg() *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
}

type marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) (PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) (PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) Clone() (PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter()
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

// PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	// validate validates PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeReservedCounter) setDefault() {
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
