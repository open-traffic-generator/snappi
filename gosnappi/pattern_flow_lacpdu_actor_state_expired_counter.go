package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateExpiredCounter *****
type patternFlowLacpduActorStateExpiredCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorStateExpiredCounter
	marshaller   marshalPatternFlowLacpduActorStateExpiredCounter
	unMarshaller unMarshalPatternFlowLacpduActorStateExpiredCounter
}

func NewPatternFlowLacpduActorStateExpiredCounter() PatternFlowLacpduActorStateExpiredCounter {
	obj := patternFlowLacpduActorStateExpiredCounter{obj: &otg.PatternFlowLacpduActorStateExpiredCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateExpiredCounter) msg() *otg.PatternFlowLacpduActorStateExpiredCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateExpiredCounter) setMsg(msg *otg.PatternFlowLacpduActorStateExpiredCounter) PatternFlowLacpduActorStateExpiredCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateExpiredCounter struct {
	obj *patternFlowLacpduActorStateExpiredCounter
}

type marshalPatternFlowLacpduActorStateExpiredCounter interface {
	// ToProto marshals PatternFlowLacpduActorStateExpiredCounter to protobuf object *otg.PatternFlowLacpduActorStateExpiredCounter
	ToProto() (*otg.PatternFlowLacpduActorStateExpiredCounter, error)
	// ToPbText marshals PatternFlowLacpduActorStateExpiredCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateExpiredCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateExpiredCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateExpiredCounter struct {
	obj *patternFlowLacpduActorStateExpiredCounter
}

type unMarshalPatternFlowLacpduActorStateExpiredCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorStateExpiredCounter from protobuf object *otg.PatternFlowLacpduActorStateExpiredCounter
	FromProto(msg *otg.PatternFlowLacpduActorStateExpiredCounter) (PatternFlowLacpduActorStateExpiredCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateExpiredCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateExpiredCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateExpiredCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateExpiredCounter) Marshal() marshalPatternFlowLacpduActorStateExpiredCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateExpiredCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateExpiredCounter) Unmarshal() unMarshalPatternFlowLacpduActorStateExpiredCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateExpiredCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateExpiredCounter) ToProto() (*otg.PatternFlowLacpduActorStateExpiredCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateExpiredCounter) FromProto(msg *otg.PatternFlowLacpduActorStateExpiredCounter) (PatternFlowLacpduActorStateExpiredCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateExpiredCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateExpiredCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateExpiredCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateExpiredCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateExpiredCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateExpiredCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateExpiredCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateExpiredCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateExpiredCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateExpiredCounter) Clone() (PatternFlowLacpduActorStateExpiredCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateExpiredCounter()
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

// PatternFlowLacpduActorStateExpiredCounter is integer counter pattern
type PatternFlowLacpduActorStateExpiredCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateExpiredCounter to protobuf object *otg.PatternFlowLacpduActorStateExpiredCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateExpiredCounter
	// setMsg unmarshals PatternFlowLacpduActorStateExpiredCounter from protobuf object *otg.PatternFlowLacpduActorStateExpiredCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateExpiredCounter) PatternFlowLacpduActorStateExpiredCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateExpiredCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateExpiredCounter
	// validate validates PatternFlowLacpduActorStateExpiredCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateExpiredCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorStateExpiredCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorStateExpiredCounter
	SetStart(value uint32) PatternFlowLacpduActorStateExpiredCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorStateExpiredCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorStateExpiredCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorStateExpiredCounter
	SetStep(value uint32) PatternFlowLacpduActorStateExpiredCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorStateExpiredCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorStateExpiredCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorStateExpiredCounter
	SetCount(value uint32) PatternFlowLacpduActorStateExpiredCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorStateExpiredCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateExpiredCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateExpiredCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorStateExpiredCounter object
func (obj *patternFlowLacpduActorStateExpiredCounter) SetStart(value uint32) PatternFlowLacpduActorStateExpiredCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateExpiredCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateExpiredCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorStateExpiredCounter object
func (obj *patternFlowLacpduActorStateExpiredCounter) SetStep(value uint32) PatternFlowLacpduActorStateExpiredCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateExpiredCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateExpiredCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorStateExpiredCounter object
func (obj *patternFlowLacpduActorStateExpiredCounter) SetCount(value uint32) PatternFlowLacpduActorStateExpiredCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorStateExpiredCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateExpiredCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateExpiredCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateExpiredCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorStateExpiredCounter) setDefault() {
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
