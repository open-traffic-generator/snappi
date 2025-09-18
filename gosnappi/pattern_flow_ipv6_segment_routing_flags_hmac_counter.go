package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsHmacCounter *****
type patternFlowIpv6SegmentRoutingFlagsHmacCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingFlagsHmacCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingFlagsHmacCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsHmacCounter() PatternFlowIpv6SegmentRoutingFlagsHmacCounter {
	obj := patternFlowIpv6SegmentRoutingFlagsHmacCounter{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter) PatternFlowIpv6SegmentRoutingFlagsHmacCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsHmacCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter
}

type marshalPatternFlowIpv6SegmentRoutingFlagsHmacCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsHmacCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsHmacCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsHmacCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsHmacCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsHmacCounter struct {
	obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsHmacCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsHmacCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter) (PatternFlowIpv6SegmentRoutingFlagsHmacCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsHmacCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsHmacCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsHmacCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsHmacCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsHmacCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsHmacCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsHmacCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHmacCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHmacCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter) (PatternFlowIpv6SegmentRoutingFlagsHmacCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHmacCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHmacCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHmacCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHmacCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHmacCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHmacCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) Clone() (PatternFlowIpv6SegmentRoutingFlagsHmacCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsHmacCounter()
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

// PatternFlowIpv6SegmentRoutingFlagsHmacCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingFlagsHmacCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsHmacCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsHmacCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsHmacCounter) PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsHmacCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsHmacCounter
	// validate validates PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsHmacCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsHmacCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsHmacCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsHmacCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsHmacCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingFlagsHmacCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsHmacCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingFlagsHmacCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsHmacCounter object
func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingFlagsHmacCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsHmacCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsHmacCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsHmacCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmacCounter) setDefault() {
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
