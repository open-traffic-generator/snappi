package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorPortNumberCounter *****
type patternFlowLacpduActorPortNumberCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorPortNumberCounter
	marshaller   marshalPatternFlowLacpduActorPortNumberCounter
	unMarshaller unMarshalPatternFlowLacpduActorPortNumberCounter
}

func NewPatternFlowLacpduActorPortNumberCounter() PatternFlowLacpduActorPortNumberCounter {
	obj := patternFlowLacpduActorPortNumberCounter{obj: &otg.PatternFlowLacpduActorPortNumberCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorPortNumberCounter) msg() *otg.PatternFlowLacpduActorPortNumberCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorPortNumberCounter) setMsg(msg *otg.PatternFlowLacpduActorPortNumberCounter) PatternFlowLacpduActorPortNumberCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorPortNumberCounter struct {
	obj *patternFlowLacpduActorPortNumberCounter
}

type marshalPatternFlowLacpduActorPortNumberCounter interface {
	// ToProto marshals PatternFlowLacpduActorPortNumberCounter to protobuf object *otg.PatternFlowLacpduActorPortNumberCounter
	ToProto() (*otg.PatternFlowLacpduActorPortNumberCounter, error)
	// ToPbText marshals PatternFlowLacpduActorPortNumberCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorPortNumberCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorPortNumberCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorPortNumberCounter struct {
	obj *patternFlowLacpduActorPortNumberCounter
}

type unMarshalPatternFlowLacpduActorPortNumberCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorPortNumberCounter from protobuf object *otg.PatternFlowLacpduActorPortNumberCounter
	FromProto(msg *otg.PatternFlowLacpduActorPortNumberCounter) (PatternFlowLacpduActorPortNumberCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorPortNumberCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorPortNumberCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorPortNumberCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorPortNumberCounter) Marshal() marshalPatternFlowLacpduActorPortNumberCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorPortNumberCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorPortNumberCounter) Unmarshal() unMarshalPatternFlowLacpduActorPortNumberCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorPortNumberCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorPortNumberCounter) ToProto() (*otg.PatternFlowLacpduActorPortNumberCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorPortNumberCounter) FromProto(msg *otg.PatternFlowLacpduActorPortNumberCounter) (PatternFlowLacpduActorPortNumberCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorPortNumberCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorPortNumberCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorPortNumberCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorPortNumberCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorPortNumberCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorPortNumberCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorPortNumberCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorPortNumberCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorPortNumberCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorPortNumberCounter) Clone() (PatternFlowLacpduActorPortNumberCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorPortNumberCounter()
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

// PatternFlowLacpduActorPortNumberCounter is integer counter pattern
type PatternFlowLacpduActorPortNumberCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorPortNumberCounter to protobuf object *otg.PatternFlowLacpduActorPortNumberCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorPortNumberCounter
	// setMsg unmarshals PatternFlowLacpduActorPortNumberCounter from protobuf object *otg.PatternFlowLacpduActorPortNumberCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorPortNumberCounter) PatternFlowLacpduActorPortNumberCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorPortNumberCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorPortNumberCounter
	// validate validates PatternFlowLacpduActorPortNumberCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorPortNumberCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorPortNumberCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorPortNumberCounter
	SetStart(value uint32) PatternFlowLacpduActorPortNumberCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorPortNumberCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorPortNumberCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorPortNumberCounter
	SetStep(value uint32) PatternFlowLacpduActorPortNumberCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorPortNumberCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorPortNumberCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorPortNumberCounter
	SetCount(value uint32) PatternFlowLacpduActorPortNumberCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorPortNumberCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorPortNumberCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorPortNumberCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorPortNumberCounter object
func (obj *patternFlowLacpduActorPortNumberCounter) SetStart(value uint32) PatternFlowLacpduActorPortNumberCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorPortNumberCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorPortNumberCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorPortNumberCounter object
func (obj *patternFlowLacpduActorPortNumberCounter) SetStep(value uint32) PatternFlowLacpduActorPortNumberCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorPortNumberCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorPortNumberCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorPortNumberCounter object
func (obj *patternFlowLacpduActorPortNumberCounter) SetCount(value uint32) PatternFlowLacpduActorPortNumberCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorPortNumberCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorPortNumberCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorPortNumberCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorPortNumberCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorPortNumberCounter) setDefault() {
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
