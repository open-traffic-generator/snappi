package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingSegmentSegmentCounter *****
type patternFlowIpv6SegmentRoutingSegmentSegmentCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingSegmentSegmentCounter
}

func NewPatternFlowIpv6SegmentRoutingSegmentSegmentCounter() PatternFlowIpv6SegmentRoutingSegmentSegmentCounter {
	obj := patternFlowIpv6SegmentRoutingSegmentSegmentCounter{obj: &otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) msg() *otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter) PatternFlowIpv6SegmentRoutingSegmentSegmentCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingSegmentSegmentCounter struct {
	obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter
}

type marshalPatternFlowIpv6SegmentRoutingSegmentSegmentCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingSegmentSegmentCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingSegmentSegmentCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingSegmentSegmentCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingSegmentSegmentCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingSegmentSegmentCounter struct {
	obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingSegmentSegmentCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingSegmentSegmentCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter) (PatternFlowIpv6SegmentRoutingSegmentSegmentCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingSegmentSegmentCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingSegmentSegmentCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingSegmentSegmentCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingSegmentSegmentCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingSegmentSegmentCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingSegmentSegmentCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingSegmentSegmentCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentSegmentCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentSegmentCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter) (PatternFlowIpv6SegmentRoutingSegmentSegmentCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentSegmentCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentSegmentCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentSegmentCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentSegmentCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentSegmentCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentSegmentCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) Clone() (PatternFlowIpv6SegmentRoutingSegmentSegmentCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingSegmentSegmentCounter()
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

// PatternFlowIpv6SegmentRoutingSegmentSegmentCounter is ipv6 counter pattern
type PatternFlowIpv6SegmentRoutingSegmentSegmentCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingSegmentSegmentCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingSegmentSegmentCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingSegmentSegmentCounter) PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	// validate validates PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingSegmentSegmentCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv6SegmentRoutingSegmentSegmentCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	SetStart(value string) PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv6SegmentRoutingSegmentSegmentCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	SetStep(value string) PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingSegmentSegmentCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingSegmentSegmentCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv6SegmentRoutingSegmentSegmentCounter object
func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) SetStart(value string) PatternFlowIpv6SegmentRoutingSegmentSegmentCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv6SegmentRoutingSegmentSegmentCounter object
func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) SetStep(value string) PatternFlowIpv6SegmentRoutingSegmentSegmentCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingSegmentSegmentCounter object
func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingSegmentSegmentCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingSegmentSegmentCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingSegmentSegmentCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingSegmentSegmentCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("::0")
	}
	if obj.obj.Step == nil {
		obj.SetStep("::1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
