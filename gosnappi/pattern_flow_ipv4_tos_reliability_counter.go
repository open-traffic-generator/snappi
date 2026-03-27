package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosReliabilityCounter *****
type patternFlowIpv4TosReliabilityCounter struct {
	validation
	obj          *otg.PatternFlowIpv4TosReliabilityCounter
	marshaller   marshalPatternFlowIpv4TosReliabilityCounter
	unMarshaller unMarshalPatternFlowIpv4TosReliabilityCounter
}

func NewPatternFlowIpv4TosReliabilityCounter() PatternFlowIpv4TosReliabilityCounter {
	obj := patternFlowIpv4TosReliabilityCounter{obj: &otg.PatternFlowIpv4TosReliabilityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosReliabilityCounter) msg() *otg.PatternFlowIpv4TosReliabilityCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TosReliabilityCounter) setMsg(msg *otg.PatternFlowIpv4TosReliabilityCounter) PatternFlowIpv4TosReliabilityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosReliabilityCounter struct {
	obj *patternFlowIpv4TosReliabilityCounter
}

type marshalPatternFlowIpv4TosReliabilityCounter interface {
	// ToProto marshals PatternFlowIpv4TosReliabilityCounter to protobuf object *otg.PatternFlowIpv4TosReliabilityCounter
	ToProto() (*otg.PatternFlowIpv4TosReliabilityCounter, error)
	// ToPbText marshals PatternFlowIpv4TosReliabilityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosReliabilityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosReliabilityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosReliabilityCounter struct {
	obj *patternFlowIpv4TosReliabilityCounter
}

type unMarshalPatternFlowIpv4TosReliabilityCounter interface {
	// FromProto unmarshals PatternFlowIpv4TosReliabilityCounter from protobuf object *otg.PatternFlowIpv4TosReliabilityCounter
	FromProto(msg *otg.PatternFlowIpv4TosReliabilityCounter) (PatternFlowIpv4TosReliabilityCounter, error)
	// FromPbText unmarshals PatternFlowIpv4TosReliabilityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosReliabilityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosReliabilityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosReliabilityCounter) Marshal() marshalPatternFlowIpv4TosReliabilityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosReliabilityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosReliabilityCounter) Unmarshal() unMarshalPatternFlowIpv4TosReliabilityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosReliabilityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosReliabilityCounter) ToProto() (*otg.PatternFlowIpv4TosReliabilityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosReliabilityCounter) FromProto(msg *otg.PatternFlowIpv4TosReliabilityCounter) (PatternFlowIpv4TosReliabilityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosReliabilityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosReliabilityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosReliabilityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosReliabilityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosReliabilityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosReliabilityCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosReliabilityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosReliabilityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosReliabilityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosReliabilityCounter) Clone() (PatternFlowIpv4TosReliabilityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosReliabilityCounter()
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

// PatternFlowIpv4TosReliabilityCounter is integer counter pattern
type PatternFlowIpv4TosReliabilityCounter interface {
	Validation
	// msg marshals PatternFlowIpv4TosReliabilityCounter to protobuf object *otg.PatternFlowIpv4TosReliabilityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosReliabilityCounter
	// setMsg unmarshals PatternFlowIpv4TosReliabilityCounter from protobuf object *otg.PatternFlowIpv4TosReliabilityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosReliabilityCounter) PatternFlowIpv4TosReliabilityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosReliabilityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosReliabilityCounter
	// validate validates PatternFlowIpv4TosReliabilityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosReliabilityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4TosReliabilityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4TosReliabilityCounter
	SetStart(value uint32) PatternFlowIpv4TosReliabilityCounter
	// HasStart checks if Start has been set in PatternFlowIpv4TosReliabilityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4TosReliabilityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4TosReliabilityCounter
	SetStep(value uint32) PatternFlowIpv4TosReliabilityCounter
	// HasStep checks if Step has been set in PatternFlowIpv4TosReliabilityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4TosReliabilityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4TosReliabilityCounter
	SetCount(value uint32) PatternFlowIpv4TosReliabilityCounter
	// HasCount checks if Count has been set in PatternFlowIpv4TosReliabilityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TosReliabilityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TosReliabilityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4TosReliabilityCounter object
func (obj *patternFlowIpv4TosReliabilityCounter) SetStart(value uint32) PatternFlowIpv4TosReliabilityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TosReliabilityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TosReliabilityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4TosReliabilityCounter object
func (obj *patternFlowIpv4TosReliabilityCounter) SetStep(value uint32) PatternFlowIpv4TosReliabilityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TosReliabilityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TosReliabilityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4TosReliabilityCounter object
func (obj *patternFlowIpv4TosReliabilityCounter) SetCount(value uint32) PatternFlowIpv4TosReliabilityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4TosReliabilityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosReliabilityCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosReliabilityCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosReliabilityCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4TosReliabilityCounter) setDefault() {
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
