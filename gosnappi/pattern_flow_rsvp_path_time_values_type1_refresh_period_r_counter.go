package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter *****
type patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	marshaller   marshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	unMarshaller unMarshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
}

func NewPatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter() PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter {
	obj := patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter{obj: &otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) msg() *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) setMsg(msg *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter struct {
	obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
}

type marshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter interface {
	// ToProto marshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter to protobuf object *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	ToProto() (*otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter, error)
	// ToPbText marshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter struct {
	obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
}

type unMarshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter from protobuf object *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	FromProto(msg *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) (PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) Marshal() marshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) Unmarshal() unMarshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) ToProto() (*otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) FromProto(msg *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) (PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) Clone() (PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter()
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

// PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter is integer counter pattern
type PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter to protobuf object *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	// setMsg unmarshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter from protobuf object *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	// validate validates PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	SetStart(value uint32) PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	SetStep(value uint32) PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	SetCount(value uint32) PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter object
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) SetStart(value uint32) PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter object
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) SetStep(value uint32) PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter object
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) SetCount(value uint32) PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(30000)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
