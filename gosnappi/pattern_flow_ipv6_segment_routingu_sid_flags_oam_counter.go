package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter *****
type patternFlowIpv6SegmentRoutinguSidFlagsOamCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
}

func NewPatternFlowIpv6SegmentRoutinguSidFlagsOamCounter() PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter {
	obj := patternFlowIpv6SegmentRoutinguSidFlagsOamCounter{obj: &otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) msg() *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutinguSidFlagsOamCounter struct {
	obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter
}

type marshalPatternFlowIpv6SegmentRoutinguSidFlagsOamCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutinguSidFlagsOamCounter struct {
	obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter
}

type unMarshalPatternFlowIpv6SegmentRoutinguSidFlagsOamCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) (PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) Marshal() marshalPatternFlowIpv6SegmentRoutinguSidFlagsOamCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutinguSidFlagsOamCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutinguSidFlagsOamCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutinguSidFlagsOamCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) (PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) Clone() (PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutinguSidFlagsOamCounter()
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

// PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	// validate validates PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter object
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter object
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter object
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOamCounter) setDefault() {
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
