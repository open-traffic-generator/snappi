package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateCollectingCounter *****
type patternFlowLacpduActorStateCollectingCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorStateCollectingCounter
	marshaller   marshalPatternFlowLacpduActorStateCollectingCounter
	unMarshaller unMarshalPatternFlowLacpduActorStateCollectingCounter
}

func NewPatternFlowLacpduActorStateCollectingCounter() PatternFlowLacpduActorStateCollectingCounter {
	obj := patternFlowLacpduActorStateCollectingCounter{obj: &otg.PatternFlowLacpduActorStateCollectingCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateCollectingCounter) msg() *otg.PatternFlowLacpduActorStateCollectingCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateCollectingCounter) setMsg(msg *otg.PatternFlowLacpduActorStateCollectingCounter) PatternFlowLacpduActorStateCollectingCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateCollectingCounter struct {
	obj *patternFlowLacpduActorStateCollectingCounter
}

type marshalPatternFlowLacpduActorStateCollectingCounter interface {
	// ToProto marshals PatternFlowLacpduActorStateCollectingCounter to protobuf object *otg.PatternFlowLacpduActorStateCollectingCounter
	ToProto() (*otg.PatternFlowLacpduActorStateCollectingCounter, error)
	// ToPbText marshals PatternFlowLacpduActorStateCollectingCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateCollectingCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateCollectingCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateCollectingCounter struct {
	obj *patternFlowLacpduActorStateCollectingCounter
}

type unMarshalPatternFlowLacpduActorStateCollectingCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorStateCollectingCounter from protobuf object *otg.PatternFlowLacpduActorStateCollectingCounter
	FromProto(msg *otg.PatternFlowLacpduActorStateCollectingCounter) (PatternFlowLacpduActorStateCollectingCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateCollectingCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateCollectingCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateCollectingCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateCollectingCounter) Marshal() marshalPatternFlowLacpduActorStateCollectingCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateCollectingCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateCollectingCounter) Unmarshal() unMarshalPatternFlowLacpduActorStateCollectingCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateCollectingCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateCollectingCounter) ToProto() (*otg.PatternFlowLacpduActorStateCollectingCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateCollectingCounter) FromProto(msg *otg.PatternFlowLacpduActorStateCollectingCounter) (PatternFlowLacpduActorStateCollectingCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateCollectingCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateCollectingCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateCollectingCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateCollectingCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateCollectingCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateCollectingCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateCollectingCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateCollectingCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateCollectingCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateCollectingCounter) Clone() (PatternFlowLacpduActorStateCollectingCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateCollectingCounter()
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

// PatternFlowLacpduActorStateCollectingCounter is integer counter pattern
type PatternFlowLacpduActorStateCollectingCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateCollectingCounter to protobuf object *otg.PatternFlowLacpduActorStateCollectingCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateCollectingCounter
	// setMsg unmarshals PatternFlowLacpduActorStateCollectingCounter from protobuf object *otg.PatternFlowLacpduActorStateCollectingCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateCollectingCounter) PatternFlowLacpduActorStateCollectingCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateCollectingCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateCollectingCounter
	// validate validates PatternFlowLacpduActorStateCollectingCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateCollectingCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorStateCollectingCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorStateCollectingCounter
	SetStart(value uint32) PatternFlowLacpduActorStateCollectingCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorStateCollectingCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorStateCollectingCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorStateCollectingCounter
	SetStep(value uint32) PatternFlowLacpduActorStateCollectingCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorStateCollectingCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorStateCollectingCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorStateCollectingCounter
	SetCount(value uint32) PatternFlowLacpduActorStateCollectingCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorStateCollectingCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateCollectingCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateCollectingCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorStateCollectingCounter object
func (obj *patternFlowLacpduActorStateCollectingCounter) SetStart(value uint32) PatternFlowLacpduActorStateCollectingCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateCollectingCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateCollectingCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorStateCollectingCounter object
func (obj *patternFlowLacpduActorStateCollectingCounter) SetStep(value uint32) PatternFlowLacpduActorStateCollectingCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateCollectingCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateCollectingCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorStateCollectingCounter object
func (obj *patternFlowLacpduActorStateCollectingCounter) SetCount(value uint32) PatternFlowLacpduActorStateCollectingCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorStateCollectingCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateCollectingCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateCollectingCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateCollectingCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorStateCollectingCounter) setDefault() {
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
