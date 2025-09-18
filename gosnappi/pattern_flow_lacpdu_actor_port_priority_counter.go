package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorPortPriorityCounter *****
type patternFlowLacpduActorPortPriorityCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorPortPriorityCounter
	marshaller   marshalPatternFlowLacpduActorPortPriorityCounter
	unMarshaller unMarshalPatternFlowLacpduActorPortPriorityCounter
}

func NewPatternFlowLacpduActorPortPriorityCounter() PatternFlowLacpduActorPortPriorityCounter {
	obj := patternFlowLacpduActorPortPriorityCounter{obj: &otg.PatternFlowLacpduActorPortPriorityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorPortPriorityCounter) msg() *otg.PatternFlowLacpduActorPortPriorityCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorPortPriorityCounter) setMsg(msg *otg.PatternFlowLacpduActorPortPriorityCounter) PatternFlowLacpduActorPortPriorityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorPortPriorityCounter struct {
	obj *patternFlowLacpduActorPortPriorityCounter
}

type marshalPatternFlowLacpduActorPortPriorityCounter interface {
	// ToProto marshals PatternFlowLacpduActorPortPriorityCounter to protobuf object *otg.PatternFlowLacpduActorPortPriorityCounter
	ToProto() (*otg.PatternFlowLacpduActorPortPriorityCounter, error)
	// ToPbText marshals PatternFlowLacpduActorPortPriorityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorPortPriorityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorPortPriorityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorPortPriorityCounter struct {
	obj *patternFlowLacpduActorPortPriorityCounter
}

type unMarshalPatternFlowLacpduActorPortPriorityCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorPortPriorityCounter from protobuf object *otg.PatternFlowLacpduActorPortPriorityCounter
	FromProto(msg *otg.PatternFlowLacpduActorPortPriorityCounter) (PatternFlowLacpduActorPortPriorityCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorPortPriorityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorPortPriorityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorPortPriorityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorPortPriorityCounter) Marshal() marshalPatternFlowLacpduActorPortPriorityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorPortPriorityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorPortPriorityCounter) Unmarshal() unMarshalPatternFlowLacpduActorPortPriorityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorPortPriorityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorPortPriorityCounter) ToProto() (*otg.PatternFlowLacpduActorPortPriorityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorPortPriorityCounter) FromProto(msg *otg.PatternFlowLacpduActorPortPriorityCounter) (PatternFlowLacpduActorPortPriorityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorPortPriorityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorPortPriorityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorPortPriorityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorPortPriorityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorPortPriorityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorPortPriorityCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorPortPriorityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorPortPriorityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorPortPriorityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorPortPriorityCounter) Clone() (PatternFlowLacpduActorPortPriorityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorPortPriorityCounter()
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

// PatternFlowLacpduActorPortPriorityCounter is integer counter pattern
type PatternFlowLacpduActorPortPriorityCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorPortPriorityCounter to protobuf object *otg.PatternFlowLacpduActorPortPriorityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorPortPriorityCounter
	// setMsg unmarshals PatternFlowLacpduActorPortPriorityCounter from protobuf object *otg.PatternFlowLacpduActorPortPriorityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorPortPriorityCounter) PatternFlowLacpduActorPortPriorityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorPortPriorityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorPortPriorityCounter
	// validate validates PatternFlowLacpduActorPortPriorityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorPortPriorityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorPortPriorityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorPortPriorityCounter
	SetStart(value uint32) PatternFlowLacpduActorPortPriorityCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorPortPriorityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorPortPriorityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorPortPriorityCounter
	SetStep(value uint32) PatternFlowLacpduActorPortPriorityCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorPortPriorityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorPortPriorityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorPortPriorityCounter
	SetCount(value uint32) PatternFlowLacpduActorPortPriorityCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorPortPriorityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorPortPriorityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorPortPriorityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorPortPriorityCounter object
func (obj *patternFlowLacpduActorPortPriorityCounter) SetStart(value uint32) PatternFlowLacpduActorPortPriorityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorPortPriorityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorPortPriorityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorPortPriorityCounter object
func (obj *patternFlowLacpduActorPortPriorityCounter) SetStep(value uint32) PatternFlowLacpduActorPortPriorityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorPortPriorityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorPortPriorityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorPortPriorityCounter object
func (obj *patternFlowLacpduActorPortPriorityCounter) SetCount(value uint32) PatternFlowLacpduActorPortPriorityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorPortPriorityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorPortPriorityCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorPortPriorityCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorPortPriorityCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorPortPriorityCounter) setDefault() {
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
