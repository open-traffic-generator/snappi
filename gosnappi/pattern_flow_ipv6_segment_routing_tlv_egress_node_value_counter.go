package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter *****
type patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter() PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter {
	obj := patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter{obj: &otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) msg() *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
}

type marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) (PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) (PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) Clone() (PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter()
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

// PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter is ipv6 counter pattern
type PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	// validate validates PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	SetStart(value string) PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	SetStep(value string) PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) SetStart(value string) PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) SetStep(value string) PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvEgressNodeValueCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("::")
	}
	if obj.obj.Step == nil {
		obj.SetStep("::1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
