package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorLengthCounter *****
type patternFlowLacpActorLengthCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorLengthCounter
	marshaller   marshalPatternFlowLacpActorLengthCounter
	unMarshaller unMarshalPatternFlowLacpActorLengthCounter
}

func NewPatternFlowLacpActorLengthCounter() PatternFlowLacpActorLengthCounter {
	obj := patternFlowLacpActorLengthCounter{obj: &otg.PatternFlowLacpActorLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorLengthCounter) msg() *otg.PatternFlowLacpActorLengthCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorLengthCounter) setMsg(msg *otg.PatternFlowLacpActorLengthCounter) PatternFlowLacpActorLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorLengthCounter struct {
	obj *patternFlowLacpActorLengthCounter
}

type marshalPatternFlowLacpActorLengthCounter interface {
	// ToProto marshals PatternFlowLacpActorLengthCounter to protobuf object *otg.PatternFlowLacpActorLengthCounter
	ToProto() (*otg.PatternFlowLacpActorLengthCounter, error)
	// ToPbText marshals PatternFlowLacpActorLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorLengthCounter struct {
	obj *patternFlowLacpActorLengthCounter
}

type unMarshalPatternFlowLacpActorLengthCounter interface {
	// FromProto unmarshals PatternFlowLacpActorLengthCounter from protobuf object *otg.PatternFlowLacpActorLengthCounter
	FromProto(msg *otg.PatternFlowLacpActorLengthCounter) (PatternFlowLacpActorLengthCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorLengthCounter) Marshal() marshalPatternFlowLacpActorLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorLengthCounter) Unmarshal() unMarshalPatternFlowLacpActorLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorLengthCounter) ToProto() (*otg.PatternFlowLacpActorLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorLengthCounter) FromProto(msg *otg.PatternFlowLacpActorLengthCounter) (PatternFlowLacpActorLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorLengthCounter) Clone() (PatternFlowLacpActorLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorLengthCounter()
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

// PatternFlowLacpActorLengthCounter is integer counter pattern
type PatternFlowLacpActorLengthCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorLengthCounter to protobuf object *otg.PatternFlowLacpActorLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorLengthCounter
	// setMsg unmarshals PatternFlowLacpActorLengthCounter from protobuf object *otg.PatternFlowLacpActorLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorLengthCounter) PatternFlowLacpActorLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorLengthCounter
	// validate validates PatternFlowLacpActorLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorLengthCounter
	SetStart(value uint32) PatternFlowLacpActorLengthCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorLengthCounter
	SetStep(value uint32) PatternFlowLacpActorLengthCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorLengthCounter
	SetCount(value uint32) PatternFlowLacpActorLengthCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorLengthCounter object
func (obj *patternFlowLacpActorLengthCounter) SetStart(value uint32) PatternFlowLacpActorLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorLengthCounter object
func (obj *patternFlowLacpActorLengthCounter) SetStep(value uint32) PatternFlowLacpActorLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorLengthCounter object
func (obj *patternFlowLacpActorLengthCounter) SetCount(value uint32) PatternFlowLacpActorLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorLengthCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(20)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
