package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorLengthCounter *****
type patternFlowLacpduActorLengthCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorLengthCounter
	marshaller   marshalPatternFlowLacpduActorLengthCounter
	unMarshaller unMarshalPatternFlowLacpduActorLengthCounter
}

func NewPatternFlowLacpduActorLengthCounter() PatternFlowLacpduActorLengthCounter {
	obj := patternFlowLacpduActorLengthCounter{obj: &otg.PatternFlowLacpduActorLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorLengthCounter) msg() *otg.PatternFlowLacpduActorLengthCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorLengthCounter) setMsg(msg *otg.PatternFlowLacpduActorLengthCounter) PatternFlowLacpduActorLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorLengthCounter struct {
	obj *patternFlowLacpduActorLengthCounter
}

type marshalPatternFlowLacpduActorLengthCounter interface {
	// ToProto marshals PatternFlowLacpduActorLengthCounter to protobuf object *otg.PatternFlowLacpduActorLengthCounter
	ToProto() (*otg.PatternFlowLacpduActorLengthCounter, error)
	// ToPbText marshals PatternFlowLacpduActorLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorLengthCounter struct {
	obj *patternFlowLacpduActorLengthCounter
}

type unMarshalPatternFlowLacpduActorLengthCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorLengthCounter from protobuf object *otg.PatternFlowLacpduActorLengthCounter
	FromProto(msg *otg.PatternFlowLacpduActorLengthCounter) (PatternFlowLacpduActorLengthCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorLengthCounter) Marshal() marshalPatternFlowLacpduActorLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorLengthCounter) Unmarshal() unMarshalPatternFlowLacpduActorLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorLengthCounter) ToProto() (*otg.PatternFlowLacpduActorLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorLengthCounter) FromProto(msg *otg.PatternFlowLacpduActorLengthCounter) (PatternFlowLacpduActorLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorLengthCounter) Clone() (PatternFlowLacpduActorLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorLengthCounter()
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

// PatternFlowLacpduActorLengthCounter is integer counter pattern
type PatternFlowLacpduActorLengthCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorLengthCounter to protobuf object *otg.PatternFlowLacpduActorLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorLengthCounter
	// setMsg unmarshals PatternFlowLacpduActorLengthCounter from protobuf object *otg.PatternFlowLacpduActorLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorLengthCounter) PatternFlowLacpduActorLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorLengthCounter
	// validate validates PatternFlowLacpduActorLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorLengthCounter
	SetStart(value uint32) PatternFlowLacpduActorLengthCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorLengthCounter
	SetStep(value uint32) PatternFlowLacpduActorLengthCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorLengthCounter
	SetCount(value uint32) PatternFlowLacpduActorLengthCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorLengthCounter object
func (obj *patternFlowLacpduActorLengthCounter) SetStart(value uint32) PatternFlowLacpduActorLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorLengthCounter object
func (obj *patternFlowLacpduActorLengthCounter) SetStep(value uint32) PatternFlowLacpduActorLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorLengthCounter object
func (obj *patternFlowLacpduActorLengthCounter) SetCount(value uint32) PatternFlowLacpduActorLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorLengthCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(20)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
