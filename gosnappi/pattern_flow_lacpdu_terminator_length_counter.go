package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduTerminatorLengthCounter *****
type patternFlowLacpduTerminatorLengthCounter struct {
	validation
	obj          *otg.PatternFlowLacpduTerminatorLengthCounter
	marshaller   marshalPatternFlowLacpduTerminatorLengthCounter
	unMarshaller unMarshalPatternFlowLacpduTerminatorLengthCounter
}

func NewPatternFlowLacpduTerminatorLengthCounter() PatternFlowLacpduTerminatorLengthCounter {
	obj := patternFlowLacpduTerminatorLengthCounter{obj: &otg.PatternFlowLacpduTerminatorLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduTerminatorLengthCounter) msg() *otg.PatternFlowLacpduTerminatorLengthCounter {
	return obj.obj
}

func (obj *patternFlowLacpduTerminatorLengthCounter) setMsg(msg *otg.PatternFlowLacpduTerminatorLengthCounter) PatternFlowLacpduTerminatorLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduTerminatorLengthCounter struct {
	obj *patternFlowLacpduTerminatorLengthCounter
}

type marshalPatternFlowLacpduTerminatorLengthCounter interface {
	// ToProto marshals PatternFlowLacpduTerminatorLengthCounter to protobuf object *otg.PatternFlowLacpduTerminatorLengthCounter
	ToProto() (*otg.PatternFlowLacpduTerminatorLengthCounter, error)
	// ToPbText marshals PatternFlowLacpduTerminatorLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduTerminatorLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduTerminatorLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduTerminatorLengthCounter struct {
	obj *patternFlowLacpduTerminatorLengthCounter
}

type unMarshalPatternFlowLacpduTerminatorLengthCounter interface {
	// FromProto unmarshals PatternFlowLacpduTerminatorLengthCounter from protobuf object *otg.PatternFlowLacpduTerminatorLengthCounter
	FromProto(msg *otg.PatternFlowLacpduTerminatorLengthCounter) (PatternFlowLacpduTerminatorLengthCounter, error)
	// FromPbText unmarshals PatternFlowLacpduTerminatorLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduTerminatorLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduTerminatorLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduTerminatorLengthCounter) Marshal() marshalPatternFlowLacpduTerminatorLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduTerminatorLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduTerminatorLengthCounter) Unmarshal() unMarshalPatternFlowLacpduTerminatorLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduTerminatorLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduTerminatorLengthCounter) ToProto() (*otg.PatternFlowLacpduTerminatorLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduTerminatorLengthCounter) FromProto(msg *otg.PatternFlowLacpduTerminatorLengthCounter) (PatternFlowLacpduTerminatorLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduTerminatorLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduTerminatorLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduTerminatorLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduTerminatorLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduTerminatorLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduTerminatorLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduTerminatorLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduTerminatorLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduTerminatorLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduTerminatorLengthCounter) Clone() (PatternFlowLacpduTerminatorLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduTerminatorLengthCounter()
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

// PatternFlowLacpduTerminatorLengthCounter is integer counter pattern
type PatternFlowLacpduTerminatorLengthCounter interface {
	Validation
	// msg marshals PatternFlowLacpduTerminatorLengthCounter to protobuf object *otg.PatternFlowLacpduTerminatorLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduTerminatorLengthCounter
	// setMsg unmarshals PatternFlowLacpduTerminatorLengthCounter from protobuf object *otg.PatternFlowLacpduTerminatorLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduTerminatorLengthCounter) PatternFlowLacpduTerminatorLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduTerminatorLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduTerminatorLengthCounter
	// validate validates PatternFlowLacpduTerminatorLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduTerminatorLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduTerminatorLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduTerminatorLengthCounter
	SetStart(value uint32) PatternFlowLacpduTerminatorLengthCounter
	// HasStart checks if Start has been set in PatternFlowLacpduTerminatorLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduTerminatorLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduTerminatorLengthCounter
	SetStep(value uint32) PatternFlowLacpduTerminatorLengthCounter
	// HasStep checks if Step has been set in PatternFlowLacpduTerminatorLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduTerminatorLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduTerminatorLengthCounter
	SetCount(value uint32) PatternFlowLacpduTerminatorLengthCounter
	// HasCount checks if Count has been set in PatternFlowLacpduTerminatorLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduTerminatorLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduTerminatorLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduTerminatorLengthCounter object
func (obj *patternFlowLacpduTerminatorLengthCounter) SetStart(value uint32) PatternFlowLacpduTerminatorLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduTerminatorLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduTerminatorLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduTerminatorLengthCounter object
func (obj *patternFlowLacpduTerminatorLengthCounter) SetStep(value uint32) PatternFlowLacpduTerminatorLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduTerminatorLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduTerminatorLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduTerminatorLengthCounter object
func (obj *patternFlowLacpduTerminatorLengthCounter) SetCount(value uint32) PatternFlowLacpduTerminatorLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduTerminatorLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduTerminatorLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduTerminatorLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduTerminatorLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduTerminatorLengthCounter) setDefault() {
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
