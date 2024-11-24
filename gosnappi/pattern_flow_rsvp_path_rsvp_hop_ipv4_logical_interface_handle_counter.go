package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter *****
type patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	marshaller   marshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	unMarshaller unMarshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
}

func NewPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter() PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter {
	obj := patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter{obj: &otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) msg() *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) setMsg(msg *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter struct {
	obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
}

type marshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter interface {
	// ToProto marshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter to protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	ToProto() (*otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter, error)
	// ToPbText marshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter struct {
	obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
}

type unMarshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter from protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	FromProto(msg *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) (PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) Marshal() marshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) Unmarshal() unMarshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) ToProto() (*otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) FromProto(msg *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) (PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) Clone() (PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter()
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

// PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter is integer counter pattern
type PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter to protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	// setMsg unmarshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter from protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	// validate validates PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	SetStart(value uint32) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	SetStep(value uint32) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	SetCount(value uint32) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter object
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) SetStart(value uint32) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter object
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) SetStep(value uint32) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter object
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) SetCount(value uint32) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) setDefault() {
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
