package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2TeidFlagCounter *****
type patternFlowGtpv2TeidFlagCounter struct {
	validation
	obj          *otg.PatternFlowGtpv2TeidFlagCounter
	marshaller   marshalPatternFlowGtpv2TeidFlagCounter
	unMarshaller unMarshalPatternFlowGtpv2TeidFlagCounter
}

func NewPatternFlowGtpv2TeidFlagCounter() PatternFlowGtpv2TeidFlagCounter {
	obj := patternFlowGtpv2TeidFlagCounter{obj: &otg.PatternFlowGtpv2TeidFlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2TeidFlagCounter) msg() *otg.PatternFlowGtpv2TeidFlagCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2TeidFlagCounter) setMsg(msg *otg.PatternFlowGtpv2TeidFlagCounter) PatternFlowGtpv2TeidFlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2TeidFlagCounter struct {
	obj *patternFlowGtpv2TeidFlagCounter
}

type marshalPatternFlowGtpv2TeidFlagCounter interface {
	// ToProto marshals PatternFlowGtpv2TeidFlagCounter to protobuf object *otg.PatternFlowGtpv2TeidFlagCounter
	ToProto() (*otg.PatternFlowGtpv2TeidFlagCounter, error)
	// ToPbText marshals PatternFlowGtpv2TeidFlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2TeidFlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2TeidFlagCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2TeidFlagCounter struct {
	obj *patternFlowGtpv2TeidFlagCounter
}

type unMarshalPatternFlowGtpv2TeidFlagCounter interface {
	// FromProto unmarshals PatternFlowGtpv2TeidFlagCounter from protobuf object *otg.PatternFlowGtpv2TeidFlagCounter
	FromProto(msg *otg.PatternFlowGtpv2TeidFlagCounter) (PatternFlowGtpv2TeidFlagCounter, error)
	// FromPbText unmarshals PatternFlowGtpv2TeidFlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2TeidFlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2TeidFlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2TeidFlagCounter) Marshal() marshalPatternFlowGtpv2TeidFlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2TeidFlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2TeidFlagCounter) Unmarshal() unMarshalPatternFlowGtpv2TeidFlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2TeidFlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2TeidFlagCounter) ToProto() (*otg.PatternFlowGtpv2TeidFlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2TeidFlagCounter) FromProto(msg *otg.PatternFlowGtpv2TeidFlagCounter) (PatternFlowGtpv2TeidFlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2TeidFlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidFlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2TeidFlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidFlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2TeidFlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidFlagCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv2TeidFlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2TeidFlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2TeidFlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2TeidFlagCounter) Clone() (PatternFlowGtpv2TeidFlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2TeidFlagCounter()
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

// PatternFlowGtpv2TeidFlagCounter is integer counter pattern
type PatternFlowGtpv2TeidFlagCounter interface {
	Validation
	// msg marshals PatternFlowGtpv2TeidFlagCounter to protobuf object *otg.PatternFlowGtpv2TeidFlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2TeidFlagCounter
	// setMsg unmarshals PatternFlowGtpv2TeidFlagCounter from protobuf object *otg.PatternFlowGtpv2TeidFlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2TeidFlagCounter) PatternFlowGtpv2TeidFlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2TeidFlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2TeidFlagCounter
	// validate validates PatternFlowGtpv2TeidFlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2TeidFlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv2TeidFlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv2TeidFlagCounter
	SetStart(value uint32) PatternFlowGtpv2TeidFlagCounter
	// HasStart checks if Start has been set in PatternFlowGtpv2TeidFlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv2TeidFlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv2TeidFlagCounter
	SetStep(value uint32) PatternFlowGtpv2TeidFlagCounter
	// HasStep checks if Step has been set in PatternFlowGtpv2TeidFlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv2TeidFlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv2TeidFlagCounter
	SetCount(value uint32) PatternFlowGtpv2TeidFlagCounter
	// HasCount checks if Count has been set in PatternFlowGtpv2TeidFlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2TeidFlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2TeidFlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv2TeidFlagCounter object
func (obj *patternFlowGtpv2TeidFlagCounter) SetStart(value uint32) PatternFlowGtpv2TeidFlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2TeidFlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2TeidFlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv2TeidFlagCounter object
func (obj *patternFlowGtpv2TeidFlagCounter) SetStep(value uint32) PatternFlowGtpv2TeidFlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2TeidFlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2TeidFlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv2TeidFlagCounter object
func (obj *patternFlowGtpv2TeidFlagCounter) SetCount(value uint32) PatternFlowGtpv2TeidFlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv2TeidFlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2TeidFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2TeidFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2TeidFlagCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv2TeidFlagCounter) setDefault() {
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
