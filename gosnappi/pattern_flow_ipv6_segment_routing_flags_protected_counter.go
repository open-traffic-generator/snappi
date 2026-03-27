package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsProtectedCounter *****
type patternFlowIpv6SegmentRoutingFlagsProtectedCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingFlagsProtectedCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsProtectedCounter() PatternFlowIpv6SegmentRoutingFlagsProtectedCounter {
	obj := patternFlowIpv6SegmentRoutingFlagsProtectedCounter{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter) PatternFlowIpv6SegmentRoutingFlagsProtectedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsProtectedCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter
}

type marshalPatternFlowIpv6SegmentRoutingFlagsProtectedCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsProtectedCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsProtectedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsProtectedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsProtectedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsProtectedCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsProtectedCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsProtectedCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter) (PatternFlowIpv6SegmentRoutingFlagsProtectedCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsProtectedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsProtectedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsProtectedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsProtectedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsProtectedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsProtectedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsProtectedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsProtectedCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsProtectedCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter) (PatternFlowIpv6SegmentRoutingFlagsProtectedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsProtectedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsProtectedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsProtectedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsProtectedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsProtectedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsProtectedCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) Clone() (PatternFlowIpv6SegmentRoutingFlagsProtectedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsProtectedCounter()
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

// PatternFlowIpv6SegmentRoutingFlagsProtectedCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingFlagsProtectedCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsProtectedCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsProtectedCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsProtectedCounter) PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	// validate validates PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsProtectedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsProtectedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsProtectedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsProtectedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsProtectedCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsProtectedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsProtectedCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsProtectedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsProtectedCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsProtectedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsProtectedCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsProtectedCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsProtectedCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtectedCounter) setDefault() {
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
