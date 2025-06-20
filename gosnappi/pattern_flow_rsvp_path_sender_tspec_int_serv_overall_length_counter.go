package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter *****
type patternFlowRSVPPathSenderTspecIntServOverallLengthCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	marshaller   marshalPatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServOverallLengthCounter() PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter {
	obj := patternFlowRSVPPathSenderTspecIntServOverallLengthCounter{obj: &otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) msg() *otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServOverallLengthCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter
}

type marshalPatternFlowRSVPPathSenderTspecIntServOverallLengthCounter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServOverallLengthCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServOverallLengthCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) (PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServOverallLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServOverallLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServOverallLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServOverallLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) (PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) Clone() (PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServOverallLengthCounter()
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

// PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter is integer counter pattern
type PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter) PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	// validate validates PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServOverallLengthCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServOverallLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(7)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
