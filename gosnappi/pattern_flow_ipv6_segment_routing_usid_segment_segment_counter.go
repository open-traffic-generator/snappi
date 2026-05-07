package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter *****
type patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
}

func NewPatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter() PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter {
	obj := patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter{obj: &otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
}

type marshalPatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) (PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) (PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter()
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

// PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter is ipv6 counter pattern
type PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	// validate validates PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	SetStart(value string) PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	SetStep(value string) PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) SetStart(value string) PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) SetStep(value string) PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentSegmentCounter) setDefault() {
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
