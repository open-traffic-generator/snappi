package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateTimeoutCounter *****
type patternFlowLacpduActorStateTimeoutCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorStateTimeoutCounter
	marshaller   marshalPatternFlowLacpduActorStateTimeoutCounter
	unMarshaller unMarshalPatternFlowLacpduActorStateTimeoutCounter
}

func NewPatternFlowLacpduActorStateTimeoutCounter() PatternFlowLacpduActorStateTimeoutCounter {
	obj := patternFlowLacpduActorStateTimeoutCounter{obj: &otg.PatternFlowLacpduActorStateTimeoutCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateTimeoutCounter) msg() *otg.PatternFlowLacpduActorStateTimeoutCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateTimeoutCounter) setMsg(msg *otg.PatternFlowLacpduActorStateTimeoutCounter) PatternFlowLacpduActorStateTimeoutCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateTimeoutCounter struct {
	obj *patternFlowLacpduActorStateTimeoutCounter
}

type marshalPatternFlowLacpduActorStateTimeoutCounter interface {
	// ToProto marshals PatternFlowLacpduActorStateTimeoutCounter to protobuf object *otg.PatternFlowLacpduActorStateTimeoutCounter
	ToProto() (*otg.PatternFlowLacpduActorStateTimeoutCounter, error)
	// ToPbText marshals PatternFlowLacpduActorStateTimeoutCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateTimeoutCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateTimeoutCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateTimeoutCounter struct {
	obj *patternFlowLacpduActorStateTimeoutCounter
}

type unMarshalPatternFlowLacpduActorStateTimeoutCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorStateTimeoutCounter from protobuf object *otg.PatternFlowLacpduActorStateTimeoutCounter
	FromProto(msg *otg.PatternFlowLacpduActorStateTimeoutCounter) (PatternFlowLacpduActorStateTimeoutCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateTimeoutCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateTimeoutCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateTimeoutCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateTimeoutCounter) Marshal() marshalPatternFlowLacpduActorStateTimeoutCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateTimeoutCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateTimeoutCounter) Unmarshal() unMarshalPatternFlowLacpduActorStateTimeoutCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateTimeoutCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateTimeoutCounter) ToProto() (*otg.PatternFlowLacpduActorStateTimeoutCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateTimeoutCounter) FromProto(msg *otg.PatternFlowLacpduActorStateTimeoutCounter) (PatternFlowLacpduActorStateTimeoutCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateTimeoutCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateTimeoutCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateTimeoutCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateTimeoutCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateTimeoutCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateTimeoutCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateTimeoutCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateTimeoutCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateTimeoutCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateTimeoutCounter) Clone() (PatternFlowLacpduActorStateTimeoutCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateTimeoutCounter()
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

// PatternFlowLacpduActorStateTimeoutCounter is integer counter pattern
type PatternFlowLacpduActorStateTimeoutCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateTimeoutCounter to protobuf object *otg.PatternFlowLacpduActorStateTimeoutCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateTimeoutCounter
	// setMsg unmarshals PatternFlowLacpduActorStateTimeoutCounter from protobuf object *otg.PatternFlowLacpduActorStateTimeoutCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateTimeoutCounter) PatternFlowLacpduActorStateTimeoutCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateTimeoutCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateTimeoutCounter
	// validate validates PatternFlowLacpduActorStateTimeoutCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateTimeoutCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorStateTimeoutCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorStateTimeoutCounter
	SetStart(value uint32) PatternFlowLacpduActorStateTimeoutCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorStateTimeoutCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorStateTimeoutCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorStateTimeoutCounter
	SetStep(value uint32) PatternFlowLacpduActorStateTimeoutCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorStateTimeoutCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorStateTimeoutCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorStateTimeoutCounter
	SetCount(value uint32) PatternFlowLacpduActorStateTimeoutCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorStateTimeoutCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateTimeoutCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateTimeoutCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorStateTimeoutCounter object
func (obj *patternFlowLacpduActorStateTimeoutCounter) SetStart(value uint32) PatternFlowLacpduActorStateTimeoutCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateTimeoutCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateTimeoutCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorStateTimeoutCounter object
func (obj *patternFlowLacpduActorStateTimeoutCounter) SetStep(value uint32) PatternFlowLacpduActorStateTimeoutCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateTimeoutCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateTimeoutCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorStateTimeoutCounter object
func (obj *patternFlowLacpduActorStateTimeoutCounter) SetCount(value uint32) PatternFlowLacpduActorStateTimeoutCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorStateTimeoutCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateTimeoutCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateTimeoutCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateTimeoutCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorStateTimeoutCounter) setDefault() {
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
