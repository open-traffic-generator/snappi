package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsAlertCounter *****
type patternFlowIpv6SegmentRoutingFlagsAlertCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingFlagsAlertCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingFlagsAlertCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsAlertCounter() PatternFlowIpv6SegmentRoutingFlagsAlertCounter {
	obj := patternFlowIpv6SegmentRoutingFlagsAlertCounter{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter) PatternFlowIpv6SegmentRoutingFlagsAlertCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsAlertCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter
}

type marshalPatternFlowIpv6SegmentRoutingFlagsAlertCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsAlertCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsAlertCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsAlertCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsAlertCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsAlertCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsAlertCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsAlertCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter) (PatternFlowIpv6SegmentRoutingFlagsAlertCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsAlertCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsAlertCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsAlertCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsAlertCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsAlertCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsAlertCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsAlertCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsAlertCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsAlertCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter) (PatternFlowIpv6SegmentRoutingFlagsAlertCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsAlertCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsAlertCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsAlertCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsAlertCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsAlertCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsAlertCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) Clone() (PatternFlowIpv6SegmentRoutingFlagsAlertCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsAlertCounter()
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

// PatternFlowIpv6SegmentRoutingFlagsAlertCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingFlagsAlertCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsAlertCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsAlertCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsAlertCounter) PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsAlertCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsAlertCounter
	// validate validates PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsAlertCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsAlertCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsAlertCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsAlertCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsAlertCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsAlertCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsAlertCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsAlertCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsAlertCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsAlertCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsAlertCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsAlertCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsAlertCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlertCounter) setDefault() {
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
