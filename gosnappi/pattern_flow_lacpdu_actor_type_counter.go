package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorTypeCounter *****
type patternFlowLacpduActorTypeCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorTypeCounter
	marshaller   marshalPatternFlowLacpduActorTypeCounter
	unMarshaller unMarshalPatternFlowLacpduActorTypeCounter
}

func NewPatternFlowLacpduActorTypeCounter() PatternFlowLacpduActorTypeCounter {
	obj := patternFlowLacpduActorTypeCounter{obj: &otg.PatternFlowLacpduActorTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorTypeCounter) msg() *otg.PatternFlowLacpduActorTypeCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorTypeCounter) setMsg(msg *otg.PatternFlowLacpduActorTypeCounter) PatternFlowLacpduActorTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorTypeCounter struct {
	obj *patternFlowLacpduActorTypeCounter
}

type marshalPatternFlowLacpduActorTypeCounter interface {
	// ToProto marshals PatternFlowLacpduActorTypeCounter to protobuf object *otg.PatternFlowLacpduActorTypeCounter
	ToProto() (*otg.PatternFlowLacpduActorTypeCounter, error)
	// ToPbText marshals PatternFlowLacpduActorTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorTypeCounter struct {
	obj *patternFlowLacpduActorTypeCounter
}

type unMarshalPatternFlowLacpduActorTypeCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorTypeCounter from protobuf object *otg.PatternFlowLacpduActorTypeCounter
	FromProto(msg *otg.PatternFlowLacpduActorTypeCounter) (PatternFlowLacpduActorTypeCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorTypeCounter) Marshal() marshalPatternFlowLacpduActorTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorTypeCounter) Unmarshal() unMarshalPatternFlowLacpduActorTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorTypeCounter) ToProto() (*otg.PatternFlowLacpduActorTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorTypeCounter) FromProto(msg *otg.PatternFlowLacpduActorTypeCounter) (PatternFlowLacpduActorTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorTypeCounter) Clone() (PatternFlowLacpduActorTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorTypeCounter()
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

// PatternFlowLacpduActorTypeCounter is integer counter pattern
type PatternFlowLacpduActorTypeCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorTypeCounter to protobuf object *otg.PatternFlowLacpduActorTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorTypeCounter
	// setMsg unmarshals PatternFlowLacpduActorTypeCounter from protobuf object *otg.PatternFlowLacpduActorTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorTypeCounter) PatternFlowLacpduActorTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorTypeCounter
	// validate validates PatternFlowLacpduActorTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorTypeCounter
	SetStart(value uint32) PatternFlowLacpduActorTypeCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorTypeCounter
	SetStep(value uint32) PatternFlowLacpduActorTypeCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorTypeCounter
	SetCount(value uint32) PatternFlowLacpduActorTypeCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorTypeCounter object
func (obj *patternFlowLacpduActorTypeCounter) SetStart(value uint32) PatternFlowLacpduActorTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorTypeCounter object
func (obj *patternFlowLacpduActorTypeCounter) SetStep(value uint32) PatternFlowLacpduActorTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorTypeCounter object
func (obj *patternFlowLacpduActorTypeCounter) SetCount(value uint32) PatternFlowLacpduActorTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorTypeCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
