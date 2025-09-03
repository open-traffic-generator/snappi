package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateAggregationCounter *****
type patternFlowLacpActorStateAggregationCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorStateAggregationCounter
	marshaller   marshalPatternFlowLacpActorStateAggregationCounter
	unMarshaller unMarshalPatternFlowLacpActorStateAggregationCounter
}

func NewPatternFlowLacpActorStateAggregationCounter() PatternFlowLacpActorStateAggregationCounter {
	obj := patternFlowLacpActorStateAggregationCounter{obj: &otg.PatternFlowLacpActorStateAggregationCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateAggregationCounter) msg() *otg.PatternFlowLacpActorStateAggregationCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorStateAggregationCounter) setMsg(msg *otg.PatternFlowLacpActorStateAggregationCounter) PatternFlowLacpActorStateAggregationCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateAggregationCounter struct {
	obj *patternFlowLacpActorStateAggregationCounter
}

type marshalPatternFlowLacpActorStateAggregationCounter interface {
	// ToProto marshals PatternFlowLacpActorStateAggregationCounter to protobuf object *otg.PatternFlowLacpActorStateAggregationCounter
	ToProto() (*otg.PatternFlowLacpActorStateAggregationCounter, error)
	// ToPbText marshals PatternFlowLacpActorStateAggregationCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateAggregationCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateAggregationCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateAggregationCounter struct {
	obj *patternFlowLacpActorStateAggregationCounter
}

type unMarshalPatternFlowLacpActorStateAggregationCounter interface {
	// FromProto unmarshals PatternFlowLacpActorStateAggregationCounter from protobuf object *otg.PatternFlowLacpActorStateAggregationCounter
	FromProto(msg *otg.PatternFlowLacpActorStateAggregationCounter) (PatternFlowLacpActorStateAggregationCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorStateAggregationCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateAggregationCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateAggregationCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateAggregationCounter) Marshal() marshalPatternFlowLacpActorStateAggregationCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateAggregationCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateAggregationCounter) Unmarshal() unMarshalPatternFlowLacpActorStateAggregationCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateAggregationCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateAggregationCounter) ToProto() (*otg.PatternFlowLacpActorStateAggregationCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateAggregationCounter) FromProto(msg *otg.PatternFlowLacpActorStateAggregationCounter) (PatternFlowLacpActorStateAggregationCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateAggregationCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateAggregationCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateAggregationCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateAggregationCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateAggregationCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateAggregationCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateAggregationCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateAggregationCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateAggregationCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateAggregationCounter) Clone() (PatternFlowLacpActorStateAggregationCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateAggregationCounter()
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

// PatternFlowLacpActorStateAggregationCounter is integer counter pattern
type PatternFlowLacpActorStateAggregationCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorStateAggregationCounter to protobuf object *otg.PatternFlowLacpActorStateAggregationCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateAggregationCounter
	// setMsg unmarshals PatternFlowLacpActorStateAggregationCounter from protobuf object *otg.PatternFlowLacpActorStateAggregationCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateAggregationCounter) PatternFlowLacpActorStateAggregationCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateAggregationCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateAggregationCounter
	// validate validates PatternFlowLacpActorStateAggregationCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateAggregationCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorStateAggregationCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorStateAggregationCounter
	SetStart(value uint32) PatternFlowLacpActorStateAggregationCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorStateAggregationCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorStateAggregationCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorStateAggregationCounter
	SetStep(value uint32) PatternFlowLacpActorStateAggregationCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorStateAggregationCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorStateAggregationCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorStateAggregationCounter
	SetCount(value uint32) PatternFlowLacpActorStateAggregationCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorStateAggregationCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateAggregationCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateAggregationCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorStateAggregationCounter object
func (obj *patternFlowLacpActorStateAggregationCounter) SetStart(value uint32) PatternFlowLacpActorStateAggregationCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateAggregationCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateAggregationCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorStateAggregationCounter object
func (obj *patternFlowLacpActorStateAggregationCounter) SetStep(value uint32) PatternFlowLacpActorStateAggregationCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateAggregationCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateAggregationCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorStateAggregationCounter object
func (obj *patternFlowLacpActorStateAggregationCounter) SetCount(value uint32) PatternFlowLacpActorStateAggregationCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorStateAggregationCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateAggregationCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateAggregationCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateAggregationCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorStateAggregationCounter) setDefault() {
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
