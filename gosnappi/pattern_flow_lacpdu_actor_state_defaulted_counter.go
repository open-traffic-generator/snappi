package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateDefaultedCounter *****
type patternFlowLacpduActorStateDefaultedCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorStateDefaultedCounter
	marshaller   marshalPatternFlowLacpduActorStateDefaultedCounter
	unMarshaller unMarshalPatternFlowLacpduActorStateDefaultedCounter
}

func NewPatternFlowLacpduActorStateDefaultedCounter() PatternFlowLacpduActorStateDefaultedCounter {
	obj := patternFlowLacpduActorStateDefaultedCounter{obj: &otg.PatternFlowLacpduActorStateDefaultedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateDefaultedCounter) msg() *otg.PatternFlowLacpduActorStateDefaultedCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateDefaultedCounter) setMsg(msg *otg.PatternFlowLacpduActorStateDefaultedCounter) PatternFlowLacpduActorStateDefaultedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateDefaultedCounter struct {
	obj *patternFlowLacpduActorStateDefaultedCounter
}

type marshalPatternFlowLacpduActorStateDefaultedCounter interface {
	// ToProto marshals PatternFlowLacpduActorStateDefaultedCounter to protobuf object *otg.PatternFlowLacpduActorStateDefaultedCounter
	ToProto() (*otg.PatternFlowLacpduActorStateDefaultedCounter, error)
	// ToPbText marshals PatternFlowLacpduActorStateDefaultedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateDefaultedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateDefaultedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateDefaultedCounter struct {
	obj *patternFlowLacpduActorStateDefaultedCounter
}

type unMarshalPatternFlowLacpduActorStateDefaultedCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorStateDefaultedCounter from protobuf object *otg.PatternFlowLacpduActorStateDefaultedCounter
	FromProto(msg *otg.PatternFlowLacpduActorStateDefaultedCounter) (PatternFlowLacpduActorStateDefaultedCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateDefaultedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateDefaultedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateDefaultedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateDefaultedCounter) Marshal() marshalPatternFlowLacpduActorStateDefaultedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateDefaultedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateDefaultedCounter) Unmarshal() unMarshalPatternFlowLacpduActorStateDefaultedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateDefaultedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateDefaultedCounter) ToProto() (*otg.PatternFlowLacpduActorStateDefaultedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateDefaultedCounter) FromProto(msg *otg.PatternFlowLacpduActorStateDefaultedCounter) (PatternFlowLacpduActorStateDefaultedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateDefaultedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateDefaultedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateDefaultedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateDefaultedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateDefaultedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateDefaultedCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateDefaultedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateDefaultedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateDefaultedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateDefaultedCounter) Clone() (PatternFlowLacpduActorStateDefaultedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateDefaultedCounter()
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

// PatternFlowLacpduActorStateDefaultedCounter is integer counter pattern
type PatternFlowLacpduActorStateDefaultedCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateDefaultedCounter to protobuf object *otg.PatternFlowLacpduActorStateDefaultedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateDefaultedCounter
	// setMsg unmarshals PatternFlowLacpduActorStateDefaultedCounter from protobuf object *otg.PatternFlowLacpduActorStateDefaultedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateDefaultedCounter) PatternFlowLacpduActorStateDefaultedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateDefaultedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateDefaultedCounter
	// validate validates PatternFlowLacpduActorStateDefaultedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateDefaultedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorStateDefaultedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorStateDefaultedCounter
	SetStart(value uint32) PatternFlowLacpduActorStateDefaultedCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorStateDefaultedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorStateDefaultedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorStateDefaultedCounter
	SetStep(value uint32) PatternFlowLacpduActorStateDefaultedCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorStateDefaultedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorStateDefaultedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorStateDefaultedCounter
	SetCount(value uint32) PatternFlowLacpduActorStateDefaultedCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorStateDefaultedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateDefaultedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateDefaultedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorStateDefaultedCounter object
func (obj *patternFlowLacpduActorStateDefaultedCounter) SetStart(value uint32) PatternFlowLacpduActorStateDefaultedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateDefaultedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateDefaultedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorStateDefaultedCounter object
func (obj *patternFlowLacpduActorStateDefaultedCounter) SetStep(value uint32) PatternFlowLacpduActorStateDefaultedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateDefaultedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateDefaultedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorStateDefaultedCounter object
func (obj *patternFlowLacpduActorStateDefaultedCounter) SetCount(value uint32) PatternFlowLacpduActorStateDefaultedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorStateDefaultedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateDefaultedCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateDefaultedCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateDefaultedCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorStateDefaultedCounter) setDefault() {
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
