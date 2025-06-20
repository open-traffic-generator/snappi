package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRsvpReservedCounter *****
type patternFlowRsvpReservedCounter struct {
	validation
	obj          *otg.PatternFlowRsvpReservedCounter
	marshaller   marshalPatternFlowRsvpReservedCounter
	unMarshaller unMarshalPatternFlowRsvpReservedCounter
}

func NewPatternFlowRsvpReservedCounter() PatternFlowRsvpReservedCounter {
	obj := patternFlowRsvpReservedCounter{obj: &otg.PatternFlowRsvpReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRsvpReservedCounter) msg() *otg.PatternFlowRsvpReservedCounter {
	return obj.obj
}

func (obj *patternFlowRsvpReservedCounter) setMsg(msg *otg.PatternFlowRsvpReservedCounter) PatternFlowRsvpReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRsvpReservedCounter struct {
	obj *patternFlowRsvpReservedCounter
}

type marshalPatternFlowRsvpReservedCounter interface {
	// ToProto marshals PatternFlowRsvpReservedCounter to protobuf object *otg.PatternFlowRsvpReservedCounter
	ToProto() (*otg.PatternFlowRsvpReservedCounter, error)
	// ToPbText marshals PatternFlowRsvpReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRsvpReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRsvpReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRsvpReservedCounter struct {
	obj *patternFlowRsvpReservedCounter
}

type unMarshalPatternFlowRsvpReservedCounter interface {
	// FromProto unmarshals PatternFlowRsvpReservedCounter from protobuf object *otg.PatternFlowRsvpReservedCounter
	FromProto(msg *otg.PatternFlowRsvpReservedCounter) (PatternFlowRsvpReservedCounter, error)
	// FromPbText unmarshals PatternFlowRsvpReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRsvpReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRsvpReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRsvpReservedCounter) Marshal() marshalPatternFlowRsvpReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRsvpReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRsvpReservedCounter) Unmarshal() unMarshalPatternFlowRsvpReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRsvpReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRsvpReservedCounter) ToProto() (*otg.PatternFlowRsvpReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRsvpReservedCounter) FromProto(msg *otg.PatternFlowRsvpReservedCounter) (PatternFlowRsvpReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRsvpReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRsvpReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRsvpReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRsvpReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRsvpReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRsvpReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowRsvpReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRsvpReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRsvpReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRsvpReservedCounter) Clone() (PatternFlowRsvpReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRsvpReservedCounter()
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

// PatternFlowRsvpReservedCounter is integer counter pattern
type PatternFlowRsvpReservedCounter interface {
	Validation
	// msg marshals PatternFlowRsvpReservedCounter to protobuf object *otg.PatternFlowRsvpReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRsvpReservedCounter
	// setMsg unmarshals PatternFlowRsvpReservedCounter from protobuf object *otg.PatternFlowRsvpReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRsvpReservedCounter) PatternFlowRsvpReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRsvpReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRsvpReservedCounter
	// validate validates PatternFlowRsvpReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRsvpReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRsvpReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRsvpReservedCounter
	SetStart(value uint32) PatternFlowRsvpReservedCounter
	// HasStart checks if Start has been set in PatternFlowRsvpReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRsvpReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRsvpReservedCounter
	SetStep(value uint32) PatternFlowRsvpReservedCounter
	// HasStep checks if Step has been set in PatternFlowRsvpReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRsvpReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRsvpReservedCounter
	SetCount(value uint32) PatternFlowRsvpReservedCounter
	// HasCount checks if Count has been set in PatternFlowRsvpReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRsvpReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRsvpReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRsvpReservedCounter object
func (obj *patternFlowRsvpReservedCounter) SetStart(value uint32) PatternFlowRsvpReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRsvpReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRsvpReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRsvpReservedCounter object
func (obj *patternFlowRsvpReservedCounter) SetStep(value uint32) PatternFlowRsvpReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRsvpReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRsvpReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRsvpReservedCounter object
func (obj *patternFlowRsvpReservedCounter) SetCount(value uint32) PatternFlowRsvpReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRsvpReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRsvpReservedCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRsvpReservedCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRsvpReservedCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRsvpReservedCounter) setDefault() {
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
