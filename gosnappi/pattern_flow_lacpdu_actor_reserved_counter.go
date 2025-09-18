package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorReservedCounter *****
type patternFlowLacpduActorReservedCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorReservedCounter
	marshaller   marshalPatternFlowLacpduActorReservedCounter
	unMarshaller unMarshalPatternFlowLacpduActorReservedCounter
}

func NewPatternFlowLacpduActorReservedCounter() PatternFlowLacpduActorReservedCounter {
	obj := patternFlowLacpduActorReservedCounter{obj: &otg.PatternFlowLacpduActorReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorReservedCounter) msg() *otg.PatternFlowLacpduActorReservedCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorReservedCounter) setMsg(msg *otg.PatternFlowLacpduActorReservedCounter) PatternFlowLacpduActorReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorReservedCounter struct {
	obj *patternFlowLacpduActorReservedCounter
}

type marshalPatternFlowLacpduActorReservedCounter interface {
	// ToProto marshals PatternFlowLacpduActorReservedCounter to protobuf object *otg.PatternFlowLacpduActorReservedCounter
	ToProto() (*otg.PatternFlowLacpduActorReservedCounter, error)
	// ToPbText marshals PatternFlowLacpduActorReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorReservedCounter struct {
	obj *patternFlowLacpduActorReservedCounter
}

type unMarshalPatternFlowLacpduActorReservedCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorReservedCounter from protobuf object *otg.PatternFlowLacpduActorReservedCounter
	FromProto(msg *otg.PatternFlowLacpduActorReservedCounter) (PatternFlowLacpduActorReservedCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorReservedCounter) Marshal() marshalPatternFlowLacpduActorReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorReservedCounter) Unmarshal() unMarshalPatternFlowLacpduActorReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorReservedCounter) ToProto() (*otg.PatternFlowLacpduActorReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorReservedCounter) FromProto(msg *otg.PatternFlowLacpduActorReservedCounter) (PatternFlowLacpduActorReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorReservedCounter) Clone() (PatternFlowLacpduActorReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorReservedCounter()
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

// PatternFlowLacpduActorReservedCounter is integer counter pattern
type PatternFlowLacpduActorReservedCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorReservedCounter to protobuf object *otg.PatternFlowLacpduActorReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorReservedCounter
	// setMsg unmarshals PatternFlowLacpduActorReservedCounter from protobuf object *otg.PatternFlowLacpduActorReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorReservedCounter) PatternFlowLacpduActorReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorReservedCounter
	// validate validates PatternFlowLacpduActorReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorReservedCounter
	SetStart(value uint32) PatternFlowLacpduActorReservedCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorReservedCounter
	SetStep(value uint32) PatternFlowLacpduActorReservedCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorReservedCounter
	SetCount(value uint32) PatternFlowLacpduActorReservedCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorReservedCounter object
func (obj *patternFlowLacpduActorReservedCounter) SetStart(value uint32) PatternFlowLacpduActorReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorReservedCounter object
func (obj *patternFlowLacpduActorReservedCounter) SetStep(value uint32) PatternFlowLacpduActorReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorReservedCounter object
func (obj *patternFlowLacpduActorReservedCounter) SetCount(value uint32) PatternFlowLacpduActorReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorReservedCounter.Start <= 16777215 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorReservedCounter.Step <= 16777215 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorReservedCounter.Count <= 16777215 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorReservedCounter) setDefault() {
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
