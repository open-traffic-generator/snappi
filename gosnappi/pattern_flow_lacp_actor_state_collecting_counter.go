package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateCollectingCounter *****
type patternFlowLacpActorStateCollectingCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorStateCollectingCounter
	marshaller   marshalPatternFlowLacpActorStateCollectingCounter
	unMarshaller unMarshalPatternFlowLacpActorStateCollectingCounter
}

func NewPatternFlowLacpActorStateCollectingCounter() PatternFlowLacpActorStateCollectingCounter {
	obj := patternFlowLacpActorStateCollectingCounter{obj: &otg.PatternFlowLacpActorStateCollectingCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateCollectingCounter) msg() *otg.PatternFlowLacpActorStateCollectingCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorStateCollectingCounter) setMsg(msg *otg.PatternFlowLacpActorStateCollectingCounter) PatternFlowLacpActorStateCollectingCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateCollectingCounter struct {
	obj *patternFlowLacpActorStateCollectingCounter
}

type marshalPatternFlowLacpActorStateCollectingCounter interface {
	// ToProto marshals PatternFlowLacpActorStateCollectingCounter to protobuf object *otg.PatternFlowLacpActorStateCollectingCounter
	ToProto() (*otg.PatternFlowLacpActorStateCollectingCounter, error)
	// ToPbText marshals PatternFlowLacpActorStateCollectingCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateCollectingCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateCollectingCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateCollectingCounter struct {
	obj *patternFlowLacpActorStateCollectingCounter
}

type unMarshalPatternFlowLacpActorStateCollectingCounter interface {
	// FromProto unmarshals PatternFlowLacpActorStateCollectingCounter from protobuf object *otg.PatternFlowLacpActorStateCollectingCounter
	FromProto(msg *otg.PatternFlowLacpActorStateCollectingCounter) (PatternFlowLacpActorStateCollectingCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorStateCollectingCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateCollectingCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateCollectingCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateCollectingCounter) Marshal() marshalPatternFlowLacpActorStateCollectingCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateCollectingCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateCollectingCounter) Unmarshal() unMarshalPatternFlowLacpActorStateCollectingCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateCollectingCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateCollectingCounter) ToProto() (*otg.PatternFlowLacpActorStateCollectingCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateCollectingCounter) FromProto(msg *otg.PatternFlowLacpActorStateCollectingCounter) (PatternFlowLacpActorStateCollectingCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateCollectingCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateCollectingCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateCollectingCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateCollectingCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateCollectingCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateCollectingCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateCollectingCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateCollectingCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateCollectingCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateCollectingCounter) Clone() (PatternFlowLacpActorStateCollectingCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateCollectingCounter()
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

// PatternFlowLacpActorStateCollectingCounter is integer counter pattern
type PatternFlowLacpActorStateCollectingCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorStateCollectingCounter to protobuf object *otg.PatternFlowLacpActorStateCollectingCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateCollectingCounter
	// setMsg unmarshals PatternFlowLacpActorStateCollectingCounter from protobuf object *otg.PatternFlowLacpActorStateCollectingCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateCollectingCounter) PatternFlowLacpActorStateCollectingCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateCollectingCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateCollectingCounter
	// validate validates PatternFlowLacpActorStateCollectingCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateCollectingCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorStateCollectingCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorStateCollectingCounter
	SetStart(value uint32) PatternFlowLacpActorStateCollectingCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorStateCollectingCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorStateCollectingCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorStateCollectingCounter
	SetStep(value uint32) PatternFlowLacpActorStateCollectingCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorStateCollectingCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorStateCollectingCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorStateCollectingCounter
	SetCount(value uint32) PatternFlowLacpActorStateCollectingCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorStateCollectingCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateCollectingCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateCollectingCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorStateCollectingCounter object
func (obj *patternFlowLacpActorStateCollectingCounter) SetStart(value uint32) PatternFlowLacpActorStateCollectingCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateCollectingCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateCollectingCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorStateCollectingCounter object
func (obj *patternFlowLacpActorStateCollectingCounter) SetStep(value uint32) PatternFlowLacpActorStateCollectingCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateCollectingCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateCollectingCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorStateCollectingCounter object
func (obj *patternFlowLacpActorStateCollectingCounter) SetCount(value uint32) PatternFlowLacpActorStateCollectingCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorStateCollectingCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateCollectingCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateCollectingCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateCollectingCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorStateCollectingCounter) setDefault() {
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
