package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter *****
type patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	marshaller   marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	unMarshaller unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
}

func NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter() PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter {
	obj := patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter{obj: &otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) msg() *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) setMsg(msg *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter struct {
	obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
}

type marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter interface {
	// ToProto marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter to protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	ToProto() (*otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter, error)
	// ToPbText marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter struct {
	obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
}

type unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter from protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	FromProto(msg *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) (PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) Marshal() marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) Unmarshal() unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) ToProto() (*otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) FromProto(msg *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) (PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) Clone() (PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter()
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

// PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter is integer counter pattern
type PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter to protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	// setMsg unmarshals PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter from protobuf object *otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	// validate validates PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	SetStart(value uint32) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	SetStep(value uint32) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	SetCount(value uint32) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) SetStart(value uint32) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) SetStep(value uint32) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter object
func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) SetCount(value uint32) PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(32)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
