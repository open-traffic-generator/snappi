package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorKeyCounter *****
type patternFlowLacpduActorKeyCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorKeyCounter
	marshaller   marshalPatternFlowLacpduActorKeyCounter
	unMarshaller unMarshalPatternFlowLacpduActorKeyCounter
}

func NewPatternFlowLacpduActorKeyCounter() PatternFlowLacpduActorKeyCounter {
	obj := patternFlowLacpduActorKeyCounter{obj: &otg.PatternFlowLacpduActorKeyCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorKeyCounter) msg() *otg.PatternFlowLacpduActorKeyCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorKeyCounter) setMsg(msg *otg.PatternFlowLacpduActorKeyCounter) PatternFlowLacpduActorKeyCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorKeyCounter struct {
	obj *patternFlowLacpduActorKeyCounter
}

type marshalPatternFlowLacpduActorKeyCounter interface {
	// ToProto marshals PatternFlowLacpduActorKeyCounter to protobuf object *otg.PatternFlowLacpduActorKeyCounter
	ToProto() (*otg.PatternFlowLacpduActorKeyCounter, error)
	// ToPbText marshals PatternFlowLacpduActorKeyCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorKeyCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorKeyCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorKeyCounter struct {
	obj *patternFlowLacpduActorKeyCounter
}

type unMarshalPatternFlowLacpduActorKeyCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorKeyCounter from protobuf object *otg.PatternFlowLacpduActorKeyCounter
	FromProto(msg *otg.PatternFlowLacpduActorKeyCounter) (PatternFlowLacpduActorKeyCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorKeyCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorKeyCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorKeyCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorKeyCounter) Marshal() marshalPatternFlowLacpduActorKeyCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorKeyCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorKeyCounter) Unmarshal() unMarshalPatternFlowLacpduActorKeyCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorKeyCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorKeyCounter) ToProto() (*otg.PatternFlowLacpduActorKeyCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorKeyCounter) FromProto(msg *otg.PatternFlowLacpduActorKeyCounter) (PatternFlowLacpduActorKeyCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorKeyCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorKeyCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorKeyCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorKeyCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorKeyCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorKeyCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorKeyCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorKeyCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorKeyCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorKeyCounter) Clone() (PatternFlowLacpduActorKeyCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorKeyCounter()
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

// PatternFlowLacpduActorKeyCounter is integer counter pattern
type PatternFlowLacpduActorKeyCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorKeyCounter to protobuf object *otg.PatternFlowLacpduActorKeyCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorKeyCounter
	// setMsg unmarshals PatternFlowLacpduActorKeyCounter from protobuf object *otg.PatternFlowLacpduActorKeyCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorKeyCounter) PatternFlowLacpduActorKeyCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorKeyCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorKeyCounter
	// validate validates PatternFlowLacpduActorKeyCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorKeyCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduActorKeyCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduActorKeyCounter
	SetStart(value uint32) PatternFlowLacpduActorKeyCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorKeyCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduActorKeyCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduActorKeyCounter
	SetStep(value uint32) PatternFlowLacpduActorKeyCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorKeyCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorKeyCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorKeyCounter
	SetCount(value uint32) PatternFlowLacpduActorKeyCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorKeyCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorKeyCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduActorKeyCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduActorKeyCounter object
func (obj *patternFlowLacpduActorKeyCounter) SetStart(value uint32) PatternFlowLacpduActorKeyCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorKeyCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduActorKeyCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduActorKeyCounter object
func (obj *patternFlowLacpduActorKeyCounter) SetStep(value uint32) PatternFlowLacpduActorKeyCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorKeyCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorKeyCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorKeyCounter object
func (obj *patternFlowLacpduActorKeyCounter) SetCount(value uint32) PatternFlowLacpduActorKeyCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorKeyCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorKeyCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorKeyCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorKeyCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduActorKeyCounter) setDefault() {
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
