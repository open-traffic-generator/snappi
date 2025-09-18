package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsOamCounter *****
type patternFlowIpv6SegmentRoutingFlagsOamCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingFlagsOamCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingFlagsOamCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsOamCounter() PatternFlowIpv6SegmentRoutingFlagsOamCounter {
	obj := patternFlowIpv6SegmentRoutingFlagsOamCounter{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter) PatternFlowIpv6SegmentRoutingFlagsOamCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsOamCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsOamCounter
}

type marshalPatternFlowIpv6SegmentRoutingFlagsOamCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsOamCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsOamCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsOamCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsOamCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsOamCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsOamCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsOamCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsOamCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter) (PatternFlowIpv6SegmentRoutingFlagsOamCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsOamCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsOamCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsOamCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsOamCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsOamCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsOamCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsOamCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOamCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOamCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter) (PatternFlowIpv6SegmentRoutingFlagsOamCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOamCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOamCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOamCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOamCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOamCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOamCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) Clone() (PatternFlowIpv6SegmentRoutingFlagsOamCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsOamCounter()
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

// PatternFlowIpv6SegmentRoutingFlagsOamCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingFlagsOamCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsOamCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsOamCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsOamCounter) PatternFlowIpv6SegmentRoutingFlagsOamCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsOamCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsOamCounter
	// validate validates PatternFlowIpv6SegmentRoutingFlagsOamCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsOamCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsOamCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsOamCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsOamCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingFlagsOamCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsOamCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsOamCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsOamCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingFlagsOamCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsOamCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsOamCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsOamCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingFlagsOamCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsOamCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsOamCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsOamCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsOamCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsOamCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsOamCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsOamCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsOamCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsOamCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingFlagsOamCounter) setDefault() {
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
