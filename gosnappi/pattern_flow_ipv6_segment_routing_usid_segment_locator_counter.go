package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter *****
type patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
}

func NewPatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter {
	obj := patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter{obj: &otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
}

type marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) (PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) (PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter()
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

// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter is ipv6 counter pattern
type PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	// validate validates PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	SetStart(value string) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	SetStep(value string) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) SetStart(value string) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) SetStep(value string) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentLocatorCounter) setDefault() {
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
