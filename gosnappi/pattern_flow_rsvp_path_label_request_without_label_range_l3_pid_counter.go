package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter *****
type patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	marshaller   marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	unMarshaller unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
}

func NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter() PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter {
	obj := patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter{obj: &otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) msg() *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) setMsg(msg *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter struct {
	obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
}

type marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter interface {
	// ToProto marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter to protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	ToProto() (*otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter, error)
	// ToPbText marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter struct {
	obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
}

type unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter from protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	FromProto(msg *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) (PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) Marshal() marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) Unmarshal() unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) ToProto() (*otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) FromProto(msg *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) (PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) Clone() (PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter()
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

// PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter is integer counter pattern
type PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter to protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	// setMsg unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter from protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	// validate validates PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	SetStart(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	SetStep(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	SetCount(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) SetStart(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) SetStep(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) SetCount(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeL3PidCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(2048)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
