package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServVersionCounter *****
type patternFlowRSVPPathSenderTspecIntServVersionCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter
	marshaller   marshalPatternFlowRSVPPathSenderTspecIntServVersionCounter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTspecIntServVersionCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServVersionCounter() PatternFlowRSVPPathSenderTspecIntServVersionCounter {
	obj := patternFlowRSVPPathSenderTspecIntServVersionCounter{obj: &otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) msg() *otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter) PatternFlowRSVPPathSenderTspecIntServVersionCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServVersionCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServVersionCounter
}

type marshalPatternFlowRSVPPathSenderTspecIntServVersionCounter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServVersionCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServVersionCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServVersionCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServVersionCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServVersionCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServVersionCounter
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServVersionCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServVersionCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter) (PatternFlowRSVPPathSenderTspecIntServVersionCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServVersionCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServVersionCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServVersionCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServVersionCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServVersionCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServVersionCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServVersionCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServVersionCounter) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServVersionCounter) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter) (PatternFlowRSVPPathSenderTspecIntServVersionCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServVersionCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServVersionCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServVersionCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServVersionCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServVersionCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServVersionCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) Clone() (PatternFlowRSVPPathSenderTspecIntServVersionCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServVersionCounter()
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

// PatternFlowRSVPPathSenderTspecIntServVersionCounter is integer counter pattern
type PatternFlowRSVPPathSenderTspecIntServVersionCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServVersionCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServVersionCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServVersionCounter) PatternFlowRSVPPathSenderTspecIntServVersionCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServVersionCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServVersionCounter
	// validate validates PatternFlowRSVPPathSenderTspecIntServVersionCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServVersionCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTspecIntServVersionCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServVersionCounter
	SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServVersionCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTspecIntServVersionCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTspecIntServVersionCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServVersionCounter
	SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServVersionCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTspecIntServVersionCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTspecIntServVersionCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServVersionCounter
	SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServVersionCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTspecIntServVersionCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServVersionCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServVersionCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServVersionCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServVersionCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServVersionCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServVersionCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServVersionCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServVersionCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServVersionCounter.Count <= 16 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServVersionCounter) setDefault() {
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
