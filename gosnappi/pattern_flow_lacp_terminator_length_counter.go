package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpTerminatorLengthCounter *****
type patternFlowLacpTerminatorLengthCounter struct {
	validation
	obj          *otg.PatternFlowLacpTerminatorLengthCounter
	marshaller   marshalPatternFlowLacpTerminatorLengthCounter
	unMarshaller unMarshalPatternFlowLacpTerminatorLengthCounter
}

func NewPatternFlowLacpTerminatorLengthCounter() PatternFlowLacpTerminatorLengthCounter {
	obj := patternFlowLacpTerminatorLengthCounter{obj: &otg.PatternFlowLacpTerminatorLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpTerminatorLengthCounter) msg() *otg.PatternFlowLacpTerminatorLengthCounter {
	return obj.obj
}

func (obj *patternFlowLacpTerminatorLengthCounter) setMsg(msg *otg.PatternFlowLacpTerminatorLengthCounter) PatternFlowLacpTerminatorLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpTerminatorLengthCounter struct {
	obj *patternFlowLacpTerminatorLengthCounter
}

type marshalPatternFlowLacpTerminatorLengthCounter interface {
	// ToProto marshals PatternFlowLacpTerminatorLengthCounter to protobuf object *otg.PatternFlowLacpTerminatorLengthCounter
	ToProto() (*otg.PatternFlowLacpTerminatorLengthCounter, error)
	// ToPbText marshals PatternFlowLacpTerminatorLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpTerminatorLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpTerminatorLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpTerminatorLengthCounter struct {
	obj *patternFlowLacpTerminatorLengthCounter
}

type unMarshalPatternFlowLacpTerminatorLengthCounter interface {
	// FromProto unmarshals PatternFlowLacpTerminatorLengthCounter from protobuf object *otg.PatternFlowLacpTerminatorLengthCounter
	FromProto(msg *otg.PatternFlowLacpTerminatorLengthCounter) (PatternFlowLacpTerminatorLengthCounter, error)
	// FromPbText unmarshals PatternFlowLacpTerminatorLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpTerminatorLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpTerminatorLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpTerminatorLengthCounter) Marshal() marshalPatternFlowLacpTerminatorLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpTerminatorLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpTerminatorLengthCounter) Unmarshal() unMarshalPatternFlowLacpTerminatorLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpTerminatorLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpTerminatorLengthCounter) ToProto() (*otg.PatternFlowLacpTerminatorLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpTerminatorLengthCounter) FromProto(msg *otg.PatternFlowLacpTerminatorLengthCounter) (PatternFlowLacpTerminatorLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpTerminatorLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpTerminatorLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpTerminatorLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpTerminatorLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpTerminatorLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpTerminatorLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpTerminatorLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpTerminatorLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpTerminatorLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpTerminatorLengthCounter) Clone() (PatternFlowLacpTerminatorLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpTerminatorLengthCounter()
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

// PatternFlowLacpTerminatorLengthCounter is integer counter pattern
type PatternFlowLacpTerminatorLengthCounter interface {
	Validation
	// msg marshals PatternFlowLacpTerminatorLengthCounter to protobuf object *otg.PatternFlowLacpTerminatorLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpTerminatorLengthCounter
	// setMsg unmarshals PatternFlowLacpTerminatorLengthCounter from protobuf object *otg.PatternFlowLacpTerminatorLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpTerminatorLengthCounter) PatternFlowLacpTerminatorLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpTerminatorLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpTerminatorLengthCounter
	// validate validates PatternFlowLacpTerminatorLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpTerminatorLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpTerminatorLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpTerminatorLengthCounter
	SetStart(value uint32) PatternFlowLacpTerminatorLengthCounter
	// HasStart checks if Start has been set in PatternFlowLacpTerminatorLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpTerminatorLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpTerminatorLengthCounter
	SetStep(value uint32) PatternFlowLacpTerminatorLengthCounter
	// HasStep checks if Step has been set in PatternFlowLacpTerminatorLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpTerminatorLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpTerminatorLengthCounter
	SetCount(value uint32) PatternFlowLacpTerminatorLengthCounter
	// HasCount checks if Count has been set in PatternFlowLacpTerminatorLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpTerminatorLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpTerminatorLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpTerminatorLengthCounter object
func (obj *patternFlowLacpTerminatorLengthCounter) SetStart(value uint32) PatternFlowLacpTerminatorLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpTerminatorLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpTerminatorLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpTerminatorLengthCounter object
func (obj *patternFlowLacpTerminatorLengthCounter) SetStep(value uint32) PatternFlowLacpTerminatorLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpTerminatorLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpTerminatorLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpTerminatorLengthCounter object
func (obj *patternFlowLacpTerminatorLengthCounter) SetCount(value uint32) PatternFlowLacpTerminatorLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpTerminatorLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpTerminatorLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpTerminatorLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpTerminatorLengthCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpTerminatorLengthCounter) setDefault() {
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
