package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter *****
type patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	marshaller   marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	unMarshaller unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
}

func NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter() PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter {
	obj := patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter{obj: &otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) msg() *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) setMsg(msg *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter struct {
	obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
}

type marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter interface {
	// ToProto marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter to protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	ToProto() (*otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter, error)
	// ToPbText marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter struct {
	obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
}

type unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter from protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	FromProto(msg *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) (PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) Marshal() marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) Unmarshal() unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) ToProto() (*otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) FromProto(msg *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) (PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) Clone() (PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter()
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

// PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter is ipv4 counter pattern
type PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter to protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	// setMsg unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter from protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	// validate validates PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	SetStart(value string) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	SetStep(value string) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	SetCount(value uint32) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) SetStart(value string) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) SetStep(value string) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) SetCount(value uint32) PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter.Step"))
		}

	}

}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4AddressCounter) setDefault() {
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
