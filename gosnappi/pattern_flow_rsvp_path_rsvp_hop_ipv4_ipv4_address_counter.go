package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter *****
type patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	marshaller   marshalPatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	unMarshaller unMarshalPatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
}

func NewPatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter() PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter {
	obj := patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter{obj: &otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) msg() *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) setMsg(msg *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter struct {
	obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
}

type marshalPatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter interface {
	// ToProto marshals PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter to protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	ToProto() (*otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter, error)
	// ToPbText marshals PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter struct {
	obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
}

type unMarshalPatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter from protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	FromProto(msg *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) (PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) Marshal() marshalPatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) Unmarshal() unMarshalPatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) ToProto() (*otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) FromProto(msg *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) (PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) Clone() (PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter()
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

// PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter is ipv4 counter pattern
type PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter to protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	// setMsg unmarshals PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter from protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	// validate validates PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	SetStart(value string) PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	SetStep(value string) PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	SetCount(value uint32) PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter object
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) SetStart(value string) PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter object
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) SetStep(value string) PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter object
func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) SetCount(value uint32) PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter.Step"))
		}

	}

}

func (obj *patternFlowRSVPPathRsvpHopIpv4Ipv4AddressCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("0.0.0.0")
	}
	if obj.obj.Step == nil {
		obj.SetStep("0.0.0.1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
