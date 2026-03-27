package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateDistributingCounter *****
type patternFlowLacpduActorStateDistributingCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorStateDistributingCounter
	marshaller   marshalPatternFlowLacpduActorStateDistributingCounter
	unMarshaller unMarshalPatternFlowLacpduActorStateDistributingCounter
}

func NewPatternFlowLacpduActorStateDistributingCounter() PatternFlowLacpduActorStateDistributingCounter {
	obj := patternFlowLacpduActorStateDistributingCounter{obj: &otg.PatternFlowLacpduActorStateDistributingCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateDistributingCounter) msg() *otg.PatternFlowLacpduActorStateDistributingCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateDistributingCounter) setMsg(msg *otg.PatternFlowLacpduActorStateDistributingCounter) PatternFlowLacpduActorStateDistributingCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateDistributingCounter struct {
	obj *patternFlowLacpduActorStateDistributingCounter
}

type marshalPatternFlowLacpduActorStateDistributingCounter interface {
	// ToProto marshals PatternFlowLacpduActorStateDistributingCounter to protobuf object *otg.PatternFlowLacpduActorStateDistributingCounter
	ToProto() (*otg.PatternFlowLacpduActorStateDistributingCounter, error)
	// ToPbText marshals PatternFlowLacpduActorStateDistributingCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateDistributingCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateDistributingCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateDistributingCounter struct {
	obj *patternFlowLacpduActorStateDistributingCounter
}

type unMarshalPatternFlowLacpduActorStateDistributingCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorStateDistributingCounter from protobuf object *otg.PatternFlowLacpduActorStateDistributingCounter
	FromProto(msg *otg.PatternFlowLacpduActorStateDistributingCounter) (PatternFlowLacpduActorStateDistributingCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateDistributingCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateDistributingCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateDistributingCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateDistributingCounter) Marshal() marshalPatternFlowLacpduActorStateDistributingCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateDistributingCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateDistributingCounter) Unmarshal() unMarshalPatternFlowLacpduActorStateDistributingCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateDistributingCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateDistributingCounter) ToProto() (*otg.PatternFlowLacpduActorStateDistributingCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateDistributingCounter) FromProto(msg *otg.PatternFlowLacpduActorStateDistributingCounter) (PatternFlowLacpduActorStateDistributingCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateDistributingCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateDistributingCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateDistributingCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateDistributingCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateDistributingCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateDistributingCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateDistributingCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateDistributingCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateDistributingCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateDistributingCounter) Clone() (PatternFlowLacpduActorStateDistributingCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateDistributingCounter()
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

// PatternFlowLacpduActorStateDistributingCounter is integer counter pattern
type PatternFlowLacpduActorStateDistributingCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateDistributingCounter to protobuf object *otg.PatternFlowLacpduActorStateDistributingCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateDistributingCounter
	// setMsg unmarshals PatternFlowLacpduActorStateDistributingCounter from protobuf object *otg.PatternFlowLacpduActorStateDistributingCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateDistributingCounter) PatternFlowLacpduActorStateDistributingCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateDistributingCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateDistributingCounter
	// validate validates PatternFlowLacpduActorStateDistributingCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateDistributingCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorStateDistributingCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorStateDistributingCounter
	SetStart(value uint32) PatternFlowLacpduActorStateDistributingCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorStateDistributingCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorStateDistributingCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorStateDistributingCounter
	SetStep(value uint32) PatternFlowLacpduActorStateDistributingCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorStateDistributingCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorStateDistributingCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorStateDistributingCounter
	SetCount(value uint32) PatternFlowLacpduActorStateDistributingCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorStateDistributingCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateDistributingCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorStateDistributingCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorStateDistributingCounter object
func (obj *patternFlowLacpduActorStateDistributingCounter) SetStart(value uint32) PatternFlowLacpduActorStateDistributingCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateDistributingCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorStateDistributingCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorStateDistributingCounter object
func (obj *patternFlowLacpduActorStateDistributingCounter) SetStep(value uint32) PatternFlowLacpduActorStateDistributingCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateDistributingCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorStateDistributingCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorStateDistributingCounter object
func (obj *patternFlowLacpduActorStateDistributingCounter) SetCount(value uint32) PatternFlowLacpduActorStateDistributingCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorStateDistributingCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateDistributingCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateDistributingCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateDistributingCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorStateDistributingCounter) setDefault() {
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
