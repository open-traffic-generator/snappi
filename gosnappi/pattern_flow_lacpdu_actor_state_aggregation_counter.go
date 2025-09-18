package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateAggregationCounter *****
type patternFlowLacpduActorStateAggregationCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorStateAggregationCounter
	marshaller   marshalPatternFlowLacpduActorStateAggregationCounter
	unMarshaller unMarshalPatternFlowLacpduActorStateAggregationCounter
}

func NewPatternFlowLacpduActorStateAggregationCounter() PatternFlowLacpduActorStateAggregationCounter {
	obj := patternFlowLacpduActorStateAggregationCounter{obj: &otg.PatternFlowLacpduActorStateAggregationCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateAggregationCounter) msg() *otg.PatternFlowLacpduActorStateAggregationCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateAggregationCounter) setMsg(msg *otg.PatternFlowLacpduActorStateAggregationCounter) PatternFlowLacpduActorStateAggregationCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateAggregationCounter struct {
	obj *patternFlowLacpduActorStateAggregationCounter
}

type marshalPatternFlowLacpduActorStateAggregationCounter interface {
	// ToProto marshals PatternFlowLacpduActorStateAggregationCounter to protobuf object *otg.PatternFlowLacpduActorStateAggregationCounter
	ToProto() (*otg.PatternFlowLacpduActorStateAggregationCounter, error)
	// ToPbText marshals PatternFlowLacpduActorStateAggregationCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateAggregationCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateAggregationCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateAggregationCounter struct {
	obj *patternFlowLacpduActorStateAggregationCounter
}

type unMarshalPatternFlowLacpduActorStateAggregationCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorStateAggregationCounter from protobuf object *otg.PatternFlowLacpduActorStateAggregationCounter
	FromProto(msg *otg.PatternFlowLacpduActorStateAggregationCounter) (PatternFlowLacpduActorStateAggregationCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateAggregationCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateAggregationCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateAggregationCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateAggregationCounter) Marshal() marshalPatternFlowLacpduActorStateAggregationCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateAggregationCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateAggregationCounter) Unmarshal() unMarshalPatternFlowLacpduActorStateAggregationCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateAggregationCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateAggregationCounter) ToProto() (*otg.PatternFlowLacpduActorStateAggregationCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateAggregationCounter) FromProto(msg *otg.PatternFlowLacpduActorStateAggregationCounter) (PatternFlowLacpduActorStateAggregationCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateAggregationCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateAggregationCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateAggregationCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateAggregationCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateAggregationCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateAggregationCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateAggregationCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateAggregationCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateAggregationCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateAggregationCounter) Clone() (PatternFlowLacpduActorStateAggregationCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateAggregationCounter()
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

// PatternFlowLacpduActorStateAggregationCounter is integer counter pattern
type PatternFlowLacpduActorStateAggregationCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateAggregationCounter to protobuf object *otg.PatternFlowLacpduActorStateAggregationCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateAggregationCounter
	// setMsg unmarshals PatternFlowLacpduActorStateAggregationCounter from protobuf object *otg.PatternFlowLacpduActorStateAggregationCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateAggregationCounter) PatternFlowLacpduActorStateAggregationCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateAggregationCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateAggregationCounter
	// validate validates PatternFlowLacpduActorStateAggregationCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateAggregationCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorStateAggregationCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorStateAggregationCounter
	SetStart(value uint32) PatternFlowLacpduActorStateAggregationCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorStateAggregationCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorStateAggregationCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorStateAggregationCounter
	SetStep(value uint32) PatternFlowLacpduActorStateAggregationCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorStateAggregationCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorStateAggregationCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorStateAggregationCounter
	SetCount(value uint32) PatternFlowLacpduActorStateAggregationCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorStateAggregationCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateAggregationCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateAggregationCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorStateAggregationCounter object
func (obj *patternFlowLacpduActorStateAggregationCounter) SetStart(value uint32) PatternFlowLacpduActorStateAggregationCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateAggregationCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateAggregationCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorStateAggregationCounter object
func (obj *patternFlowLacpduActorStateAggregationCounter) SetStep(value uint32) PatternFlowLacpduActorStateAggregationCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateAggregationCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateAggregationCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorStateAggregationCounter object
func (obj *patternFlowLacpduActorStateAggregationCounter) SetCount(value uint32) PatternFlowLacpduActorStateAggregationCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorStateAggregationCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateAggregationCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateAggregationCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateAggregationCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorStateAggregationCounter) setDefault() {
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
