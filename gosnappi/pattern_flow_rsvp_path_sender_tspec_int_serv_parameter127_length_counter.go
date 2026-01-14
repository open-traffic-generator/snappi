package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter *****
type patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	marshaller   marshalPatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter() PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter {
	obj := patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter{obj: &otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) msg() *otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
}

type marshalPatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) (PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) (PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) Clone() (PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter()
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

// PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter is integer counter pattern
type PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	// validate validates PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(5)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
