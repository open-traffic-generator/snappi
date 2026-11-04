package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6UsidDstLocatorLengthCounter *****
type patternFlowIpv6UsidDstLocatorLengthCounter struct {
	validation
	obj          *otg.PatternFlowIpv6UsidDstLocatorLengthCounter
	marshaller   marshalPatternFlowIpv6UsidDstLocatorLengthCounter
	unMarshaller unMarshalPatternFlowIpv6UsidDstLocatorLengthCounter
}

func NewPatternFlowIpv6UsidDstLocatorLengthCounter() PatternFlowIpv6UsidDstLocatorLengthCounter {
	obj := patternFlowIpv6UsidDstLocatorLengthCounter{obj: &otg.PatternFlowIpv6UsidDstLocatorLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) msg() *otg.PatternFlowIpv6UsidDstLocatorLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) setMsg(msg *otg.PatternFlowIpv6UsidDstLocatorLengthCounter) PatternFlowIpv6UsidDstLocatorLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6UsidDstLocatorLengthCounter struct {
	obj *patternFlowIpv6UsidDstLocatorLengthCounter
}

type marshalPatternFlowIpv6UsidDstLocatorLengthCounter interface {
	// ToProto marshals PatternFlowIpv6UsidDstLocatorLengthCounter to protobuf object *otg.PatternFlowIpv6UsidDstLocatorLengthCounter
	ToProto() (*otg.PatternFlowIpv6UsidDstLocatorLengthCounter, error)
	// ToPbText marshals PatternFlowIpv6UsidDstLocatorLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6UsidDstLocatorLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6UsidDstLocatorLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6UsidDstLocatorLengthCounter struct {
	obj *patternFlowIpv6UsidDstLocatorLengthCounter
}

type unMarshalPatternFlowIpv6UsidDstLocatorLengthCounter interface {
	// FromProto unmarshals PatternFlowIpv6UsidDstLocatorLengthCounter from protobuf object *otg.PatternFlowIpv6UsidDstLocatorLengthCounter
	FromProto(msg *otg.PatternFlowIpv6UsidDstLocatorLengthCounter) (PatternFlowIpv6UsidDstLocatorLengthCounter, error)
	// FromPbText unmarshals PatternFlowIpv6UsidDstLocatorLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6UsidDstLocatorLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6UsidDstLocatorLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) Marshal() marshalPatternFlowIpv6UsidDstLocatorLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6UsidDstLocatorLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) Unmarshal() unMarshalPatternFlowIpv6UsidDstLocatorLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6UsidDstLocatorLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6UsidDstLocatorLengthCounter) ToProto() (*otg.PatternFlowIpv6UsidDstLocatorLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6UsidDstLocatorLengthCounter) FromProto(msg *otg.PatternFlowIpv6UsidDstLocatorLengthCounter) (PatternFlowIpv6UsidDstLocatorLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6UsidDstLocatorLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6UsidDstLocatorLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6UsidDstLocatorLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6UsidDstLocatorLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6UsidDstLocatorLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6UsidDstLocatorLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) Clone() (PatternFlowIpv6UsidDstLocatorLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6UsidDstLocatorLengthCounter()
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

// PatternFlowIpv6UsidDstLocatorLengthCounter is integer counter pattern
type PatternFlowIpv6UsidDstLocatorLengthCounter interface {
	Validation
	// msg marshals PatternFlowIpv6UsidDstLocatorLengthCounter to protobuf object *otg.PatternFlowIpv6UsidDstLocatorLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6UsidDstLocatorLengthCounter
	// setMsg unmarshals PatternFlowIpv6UsidDstLocatorLengthCounter from protobuf object *otg.PatternFlowIpv6UsidDstLocatorLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6UsidDstLocatorLengthCounter) PatternFlowIpv6UsidDstLocatorLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6UsidDstLocatorLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6UsidDstLocatorLengthCounter
	// validate validates PatternFlowIpv6UsidDstLocatorLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6UsidDstLocatorLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6UsidDstLocatorLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6UsidDstLocatorLengthCounter
	SetStart(value uint32) PatternFlowIpv6UsidDstLocatorLengthCounter
	// HasStart checks if Start has been set in PatternFlowIpv6UsidDstLocatorLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6UsidDstLocatorLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6UsidDstLocatorLengthCounter
	SetStep(value uint32) PatternFlowIpv6UsidDstLocatorLengthCounter
	// HasStep checks if Step has been set in PatternFlowIpv6UsidDstLocatorLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6UsidDstLocatorLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6UsidDstLocatorLengthCounter
	SetCount(value uint32) PatternFlowIpv6UsidDstLocatorLengthCounter
	// HasCount checks if Count has been set in PatternFlowIpv6UsidDstLocatorLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6UsidDstLocatorLengthCounter object
func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) SetStart(value uint32) PatternFlowIpv6UsidDstLocatorLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6UsidDstLocatorLengthCounter object
func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) SetStep(value uint32) PatternFlowIpv6UsidDstLocatorLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6UsidDstLocatorLengthCounter object
func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) SetCount(value uint32) PatternFlowIpv6UsidDstLocatorLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6UsidDstLocatorLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6UsidDstLocatorLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6UsidDstLocatorLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6UsidDstLocatorLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(32)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
