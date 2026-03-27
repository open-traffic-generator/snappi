package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter *****
type patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	marshaller   marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	unMarshaller unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
}

func NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter {
	obj := patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter{obj: &otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) msg() *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) setMsg(msg *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter struct {
	obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
}

type marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter interface {
	// ToProto marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter to protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	ToProto() (*otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter, error)
	// ToPbText marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter struct {
	obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
}

type unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter from protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	FromProto(msg *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) Marshal() marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) Unmarshal() unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) ToProto() (*otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) FromProto(msg *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) Clone() (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter()
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

// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter is integer counter pattern
type PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter to protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	// setMsg unmarshals PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter from protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	// validate validates PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	SetStart(value uint32) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	SetStep(value uint32) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	SetCount(value uint32) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) SetStart(value uint32) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) SetStep(value uint32) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter object
func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) SetCount(value uint32) PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBitCounter) setDefault() {
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
