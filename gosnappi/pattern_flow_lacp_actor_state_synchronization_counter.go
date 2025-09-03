package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateSynchronizationCounter *****
type patternFlowLacpActorStateSynchronizationCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorStateSynchronizationCounter
	marshaller   marshalPatternFlowLacpActorStateSynchronizationCounter
	unMarshaller unMarshalPatternFlowLacpActorStateSynchronizationCounter
}

func NewPatternFlowLacpActorStateSynchronizationCounter() PatternFlowLacpActorStateSynchronizationCounter {
	obj := patternFlowLacpActorStateSynchronizationCounter{obj: &otg.PatternFlowLacpActorStateSynchronizationCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateSynchronizationCounter) msg() *otg.PatternFlowLacpActorStateSynchronizationCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorStateSynchronizationCounter) setMsg(msg *otg.PatternFlowLacpActorStateSynchronizationCounter) PatternFlowLacpActorStateSynchronizationCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateSynchronizationCounter struct {
	obj *patternFlowLacpActorStateSynchronizationCounter
}

type marshalPatternFlowLacpActorStateSynchronizationCounter interface {
	// ToProto marshals PatternFlowLacpActorStateSynchronizationCounter to protobuf object *otg.PatternFlowLacpActorStateSynchronizationCounter
	ToProto() (*otg.PatternFlowLacpActorStateSynchronizationCounter, error)
	// ToPbText marshals PatternFlowLacpActorStateSynchronizationCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateSynchronizationCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateSynchronizationCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateSynchronizationCounter struct {
	obj *patternFlowLacpActorStateSynchronizationCounter
}

type unMarshalPatternFlowLacpActorStateSynchronizationCounter interface {
	// FromProto unmarshals PatternFlowLacpActorStateSynchronizationCounter from protobuf object *otg.PatternFlowLacpActorStateSynchronizationCounter
	FromProto(msg *otg.PatternFlowLacpActorStateSynchronizationCounter) (PatternFlowLacpActorStateSynchronizationCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorStateSynchronizationCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateSynchronizationCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateSynchronizationCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateSynchronizationCounter) Marshal() marshalPatternFlowLacpActorStateSynchronizationCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateSynchronizationCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateSynchronizationCounter) Unmarshal() unMarshalPatternFlowLacpActorStateSynchronizationCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateSynchronizationCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateSynchronizationCounter) ToProto() (*otg.PatternFlowLacpActorStateSynchronizationCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateSynchronizationCounter) FromProto(msg *otg.PatternFlowLacpActorStateSynchronizationCounter) (PatternFlowLacpActorStateSynchronizationCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateSynchronizationCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateSynchronizationCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateSynchronizationCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateSynchronizationCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateSynchronizationCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateSynchronizationCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateSynchronizationCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateSynchronizationCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateSynchronizationCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateSynchronizationCounter) Clone() (PatternFlowLacpActorStateSynchronizationCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateSynchronizationCounter()
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

// PatternFlowLacpActorStateSynchronizationCounter is integer counter pattern
type PatternFlowLacpActorStateSynchronizationCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorStateSynchronizationCounter to protobuf object *otg.PatternFlowLacpActorStateSynchronizationCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateSynchronizationCounter
	// setMsg unmarshals PatternFlowLacpActorStateSynchronizationCounter from protobuf object *otg.PatternFlowLacpActorStateSynchronizationCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateSynchronizationCounter) PatternFlowLacpActorStateSynchronizationCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateSynchronizationCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateSynchronizationCounter
	// validate validates PatternFlowLacpActorStateSynchronizationCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateSynchronizationCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorStateSynchronizationCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorStateSynchronizationCounter
	SetStart(value uint32) PatternFlowLacpActorStateSynchronizationCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorStateSynchronizationCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorStateSynchronizationCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorStateSynchronizationCounter
	SetStep(value uint32) PatternFlowLacpActorStateSynchronizationCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorStateSynchronizationCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorStateSynchronizationCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorStateSynchronizationCounter
	SetCount(value uint32) PatternFlowLacpActorStateSynchronizationCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorStateSynchronizationCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateSynchronizationCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateSynchronizationCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorStateSynchronizationCounter object
func (obj *patternFlowLacpActorStateSynchronizationCounter) SetStart(value uint32) PatternFlowLacpActorStateSynchronizationCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateSynchronizationCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateSynchronizationCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorStateSynchronizationCounter object
func (obj *patternFlowLacpActorStateSynchronizationCounter) SetStep(value uint32) PatternFlowLacpActorStateSynchronizationCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateSynchronizationCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateSynchronizationCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorStateSynchronizationCounter object
func (obj *patternFlowLacpActorStateSynchronizationCounter) SetCount(value uint32) PatternFlowLacpActorStateSynchronizationCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorStateSynchronizationCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateSynchronizationCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateSynchronizationCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateSynchronizationCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorStateSynchronizationCounter) setDefault() {
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
