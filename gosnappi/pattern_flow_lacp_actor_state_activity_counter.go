package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateActivityCounter *****
type patternFlowLacpActorStateActivityCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorStateActivityCounter
	marshaller   marshalPatternFlowLacpActorStateActivityCounter
	unMarshaller unMarshalPatternFlowLacpActorStateActivityCounter
}

func NewPatternFlowLacpActorStateActivityCounter() PatternFlowLacpActorStateActivityCounter {
	obj := patternFlowLacpActorStateActivityCounter{obj: &otg.PatternFlowLacpActorStateActivityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateActivityCounter) msg() *otg.PatternFlowLacpActorStateActivityCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorStateActivityCounter) setMsg(msg *otg.PatternFlowLacpActorStateActivityCounter) PatternFlowLacpActorStateActivityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateActivityCounter struct {
	obj *patternFlowLacpActorStateActivityCounter
}

type marshalPatternFlowLacpActorStateActivityCounter interface {
	// ToProto marshals PatternFlowLacpActorStateActivityCounter to protobuf object *otg.PatternFlowLacpActorStateActivityCounter
	ToProto() (*otg.PatternFlowLacpActorStateActivityCounter, error)
	// ToPbText marshals PatternFlowLacpActorStateActivityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateActivityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateActivityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateActivityCounter struct {
	obj *patternFlowLacpActorStateActivityCounter
}

type unMarshalPatternFlowLacpActorStateActivityCounter interface {
	// FromProto unmarshals PatternFlowLacpActorStateActivityCounter from protobuf object *otg.PatternFlowLacpActorStateActivityCounter
	FromProto(msg *otg.PatternFlowLacpActorStateActivityCounter) (PatternFlowLacpActorStateActivityCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorStateActivityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateActivityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateActivityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateActivityCounter) Marshal() marshalPatternFlowLacpActorStateActivityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateActivityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateActivityCounter) Unmarshal() unMarshalPatternFlowLacpActorStateActivityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateActivityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateActivityCounter) ToProto() (*otg.PatternFlowLacpActorStateActivityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateActivityCounter) FromProto(msg *otg.PatternFlowLacpActorStateActivityCounter) (PatternFlowLacpActorStateActivityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateActivityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateActivityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateActivityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateActivityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateActivityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateActivityCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateActivityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateActivityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateActivityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateActivityCounter) Clone() (PatternFlowLacpActorStateActivityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateActivityCounter()
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

// PatternFlowLacpActorStateActivityCounter is integer counter pattern
type PatternFlowLacpActorStateActivityCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorStateActivityCounter to protobuf object *otg.PatternFlowLacpActorStateActivityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateActivityCounter
	// setMsg unmarshals PatternFlowLacpActorStateActivityCounter from protobuf object *otg.PatternFlowLacpActorStateActivityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateActivityCounter) PatternFlowLacpActorStateActivityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateActivityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateActivityCounter
	// validate validates PatternFlowLacpActorStateActivityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateActivityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorStateActivityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorStateActivityCounter
	SetStart(value uint32) PatternFlowLacpActorStateActivityCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorStateActivityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorStateActivityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorStateActivityCounter
	SetStep(value uint32) PatternFlowLacpActorStateActivityCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorStateActivityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorStateActivityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorStateActivityCounter
	SetCount(value uint32) PatternFlowLacpActorStateActivityCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorStateActivityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateActivityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateActivityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorStateActivityCounter object
func (obj *patternFlowLacpActorStateActivityCounter) SetStart(value uint32) PatternFlowLacpActorStateActivityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateActivityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateActivityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorStateActivityCounter object
func (obj *patternFlowLacpActorStateActivityCounter) SetStep(value uint32) PatternFlowLacpActorStateActivityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateActivityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateActivityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorStateActivityCounter object
func (obj *patternFlowLacpActorStateActivityCounter) SetCount(value uint32) PatternFlowLacpActorStateActivityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorStateActivityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateActivityCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateActivityCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateActivityCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorStateActivityCounter) setDefault() {
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
