package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateActivityCounter *****
type patternFlowLacpduActorStateActivityCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorStateActivityCounter
	marshaller   marshalPatternFlowLacpduActorStateActivityCounter
	unMarshaller unMarshalPatternFlowLacpduActorStateActivityCounter
}

func NewPatternFlowLacpduActorStateActivityCounter() PatternFlowLacpduActorStateActivityCounter {
	obj := patternFlowLacpduActorStateActivityCounter{obj: &otg.PatternFlowLacpduActorStateActivityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateActivityCounter) msg() *otg.PatternFlowLacpduActorStateActivityCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateActivityCounter) setMsg(msg *otg.PatternFlowLacpduActorStateActivityCounter) PatternFlowLacpduActorStateActivityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateActivityCounter struct {
	obj *patternFlowLacpduActorStateActivityCounter
}

type marshalPatternFlowLacpduActorStateActivityCounter interface {
	// ToProto marshals PatternFlowLacpduActorStateActivityCounter to protobuf object *otg.PatternFlowLacpduActorStateActivityCounter
	ToProto() (*otg.PatternFlowLacpduActorStateActivityCounter, error)
	// ToPbText marshals PatternFlowLacpduActorStateActivityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateActivityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateActivityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateActivityCounter struct {
	obj *patternFlowLacpduActorStateActivityCounter
}

type unMarshalPatternFlowLacpduActorStateActivityCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorStateActivityCounter from protobuf object *otg.PatternFlowLacpduActorStateActivityCounter
	FromProto(msg *otg.PatternFlowLacpduActorStateActivityCounter) (PatternFlowLacpduActorStateActivityCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateActivityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateActivityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateActivityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateActivityCounter) Marshal() marshalPatternFlowLacpduActorStateActivityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateActivityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateActivityCounter) Unmarshal() unMarshalPatternFlowLacpduActorStateActivityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateActivityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateActivityCounter) ToProto() (*otg.PatternFlowLacpduActorStateActivityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateActivityCounter) FromProto(msg *otg.PatternFlowLacpduActorStateActivityCounter) (PatternFlowLacpduActorStateActivityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateActivityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateActivityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateActivityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateActivityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateActivityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateActivityCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateActivityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateActivityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateActivityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateActivityCounter) Clone() (PatternFlowLacpduActorStateActivityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateActivityCounter()
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

// PatternFlowLacpduActorStateActivityCounter is integer counter pattern
type PatternFlowLacpduActorStateActivityCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateActivityCounter to protobuf object *otg.PatternFlowLacpduActorStateActivityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateActivityCounter
	// setMsg unmarshals PatternFlowLacpduActorStateActivityCounter from protobuf object *otg.PatternFlowLacpduActorStateActivityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateActivityCounter) PatternFlowLacpduActorStateActivityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateActivityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateActivityCounter
	// validate validates PatternFlowLacpduActorStateActivityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateActivityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorStateActivityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorStateActivityCounter
	SetStart(value uint32) PatternFlowLacpduActorStateActivityCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorStateActivityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorStateActivityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorStateActivityCounter
	SetStep(value uint32) PatternFlowLacpduActorStateActivityCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorStateActivityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorStateActivityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorStateActivityCounter
	SetCount(value uint32) PatternFlowLacpduActorStateActivityCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorStateActivityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateActivityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateActivityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorStateActivityCounter object
func (obj *patternFlowLacpduActorStateActivityCounter) SetStart(value uint32) PatternFlowLacpduActorStateActivityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateActivityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateActivityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorStateActivityCounter object
func (obj *patternFlowLacpduActorStateActivityCounter) SetStep(value uint32) PatternFlowLacpduActorStateActivityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateActivityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateActivityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorStateActivityCounter object
func (obj *patternFlowLacpduActorStateActivityCounter) SetCount(value uint32) PatternFlowLacpduActorStateActivityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorStateActivityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateActivityCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateActivityCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateActivityCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorStateActivityCounter) setDefault() {
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
