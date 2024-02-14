package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter *****
type patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	marshaller   marshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter() PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter {
	obj := patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter{obj: &otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) msg() *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
}

type marshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) (PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) (PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) Clone() (PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter()
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

// PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter is integer counter pattern
type PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	// validate validates PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) setDefault() {
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
