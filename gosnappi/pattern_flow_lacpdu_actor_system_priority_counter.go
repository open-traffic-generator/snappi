package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorSystemPriorityCounter *****
type patternFlowLacpduActorSystemPriorityCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorSystemPriorityCounter
	marshaller   marshalPatternFlowLacpduActorSystemPriorityCounter
	unMarshaller unMarshalPatternFlowLacpduActorSystemPriorityCounter
}

func NewPatternFlowLacpduActorSystemPriorityCounter() PatternFlowLacpduActorSystemPriorityCounter {
	obj := patternFlowLacpduActorSystemPriorityCounter{obj: &otg.PatternFlowLacpduActorSystemPriorityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorSystemPriorityCounter) msg() *otg.PatternFlowLacpduActorSystemPriorityCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorSystemPriorityCounter) setMsg(msg *otg.PatternFlowLacpduActorSystemPriorityCounter) PatternFlowLacpduActorSystemPriorityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorSystemPriorityCounter struct {
	obj *patternFlowLacpduActorSystemPriorityCounter
}

type marshalPatternFlowLacpduActorSystemPriorityCounter interface {
	// ToProto marshals PatternFlowLacpduActorSystemPriorityCounter to protobuf object *otg.PatternFlowLacpduActorSystemPriorityCounter
	ToProto() (*otg.PatternFlowLacpduActorSystemPriorityCounter, error)
	// ToPbText marshals PatternFlowLacpduActorSystemPriorityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorSystemPriorityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorSystemPriorityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorSystemPriorityCounter struct {
	obj *patternFlowLacpduActorSystemPriorityCounter
}

type unMarshalPatternFlowLacpduActorSystemPriorityCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorSystemPriorityCounter from protobuf object *otg.PatternFlowLacpduActorSystemPriorityCounter
	FromProto(msg *otg.PatternFlowLacpduActorSystemPriorityCounter) (PatternFlowLacpduActorSystemPriorityCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorSystemPriorityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorSystemPriorityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorSystemPriorityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorSystemPriorityCounter) Marshal() marshalPatternFlowLacpduActorSystemPriorityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorSystemPriorityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorSystemPriorityCounter) Unmarshal() unMarshalPatternFlowLacpduActorSystemPriorityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorSystemPriorityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorSystemPriorityCounter) ToProto() (*otg.PatternFlowLacpduActorSystemPriorityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorSystemPriorityCounter) FromProto(msg *otg.PatternFlowLacpduActorSystemPriorityCounter) (PatternFlowLacpduActorSystemPriorityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorSystemPriorityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorSystemPriorityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorSystemPriorityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorSystemPriorityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorSystemPriorityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorSystemPriorityCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorSystemPriorityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorSystemPriorityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorSystemPriorityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorSystemPriorityCounter) Clone() (PatternFlowLacpduActorSystemPriorityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorSystemPriorityCounter()
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

// PatternFlowLacpduActorSystemPriorityCounter is integer counter pattern
type PatternFlowLacpduActorSystemPriorityCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorSystemPriorityCounter to protobuf object *otg.PatternFlowLacpduActorSystemPriorityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorSystemPriorityCounter
	// setMsg unmarshals PatternFlowLacpduActorSystemPriorityCounter from protobuf object *otg.PatternFlowLacpduActorSystemPriorityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorSystemPriorityCounter) PatternFlowLacpduActorSystemPriorityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorSystemPriorityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorSystemPriorityCounter
	// validate validates PatternFlowLacpduActorSystemPriorityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorSystemPriorityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorSystemPriorityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorSystemPriorityCounter
	SetStart(value uint32) PatternFlowLacpduActorSystemPriorityCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorSystemPriorityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorSystemPriorityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorSystemPriorityCounter
	SetStep(value uint32) PatternFlowLacpduActorSystemPriorityCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorSystemPriorityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorSystemPriorityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorSystemPriorityCounter
	SetCount(value uint32) PatternFlowLacpduActorSystemPriorityCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorSystemPriorityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorSystemPriorityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorSystemPriorityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorSystemPriorityCounter object
func (obj *patternFlowLacpduActorSystemPriorityCounter) SetStart(value uint32) PatternFlowLacpduActorSystemPriorityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorSystemPriorityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorSystemPriorityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorSystemPriorityCounter object
func (obj *patternFlowLacpduActorSystemPriorityCounter) SetStep(value uint32) PatternFlowLacpduActorSystemPriorityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorSystemPriorityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorSystemPriorityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorSystemPriorityCounter object
func (obj *patternFlowLacpduActorSystemPriorityCounter) SetCount(value uint32) PatternFlowLacpduActorSystemPriorityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorSystemPriorityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorSystemPriorityCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorSystemPriorityCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorSystemPriorityCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorSystemPriorityCounter) setDefault() {
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
