package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1TeidCounter *****
type patternFlowGtpv1TeidCounter struct {
	validation
	obj          *otg.PatternFlowGtpv1TeidCounter
	marshaller   marshalPatternFlowGtpv1TeidCounter
	unMarshaller unMarshalPatternFlowGtpv1TeidCounter
}

func NewPatternFlowGtpv1TeidCounter() PatternFlowGtpv1TeidCounter {
	obj := patternFlowGtpv1TeidCounter{obj: &otg.PatternFlowGtpv1TeidCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1TeidCounter) msg() *otg.PatternFlowGtpv1TeidCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1TeidCounter) setMsg(msg *otg.PatternFlowGtpv1TeidCounter) PatternFlowGtpv1TeidCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1TeidCounter struct {
	obj *patternFlowGtpv1TeidCounter
}

type marshalPatternFlowGtpv1TeidCounter interface {
	// ToProto marshals PatternFlowGtpv1TeidCounter to protobuf object *otg.PatternFlowGtpv1TeidCounter
	ToProto() (*otg.PatternFlowGtpv1TeidCounter, error)
	// ToPbText marshals PatternFlowGtpv1TeidCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1TeidCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1TeidCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1TeidCounter struct {
	obj *patternFlowGtpv1TeidCounter
}

type unMarshalPatternFlowGtpv1TeidCounter interface {
	// FromProto unmarshals PatternFlowGtpv1TeidCounter from protobuf object *otg.PatternFlowGtpv1TeidCounter
	FromProto(msg *otg.PatternFlowGtpv1TeidCounter) (PatternFlowGtpv1TeidCounter, error)
	// FromPbText unmarshals PatternFlowGtpv1TeidCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1TeidCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1TeidCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1TeidCounter) Marshal() marshalPatternFlowGtpv1TeidCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1TeidCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1TeidCounter) Unmarshal() unMarshalPatternFlowGtpv1TeidCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1TeidCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1TeidCounter) ToProto() (*otg.PatternFlowGtpv1TeidCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1TeidCounter) FromProto(msg *otg.PatternFlowGtpv1TeidCounter) (PatternFlowGtpv1TeidCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1TeidCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1TeidCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1TeidCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1TeidCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1TeidCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1TeidCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv1TeidCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1TeidCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1TeidCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1TeidCounter) Clone() (PatternFlowGtpv1TeidCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1TeidCounter()
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

// PatternFlowGtpv1TeidCounter is integer counter pattern
type PatternFlowGtpv1TeidCounter interface {
	Validation
	// msg marshals PatternFlowGtpv1TeidCounter to protobuf object *otg.PatternFlowGtpv1TeidCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1TeidCounter
	// setMsg unmarshals PatternFlowGtpv1TeidCounter from protobuf object *otg.PatternFlowGtpv1TeidCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1TeidCounter) PatternFlowGtpv1TeidCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1TeidCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1TeidCounter
	// validate validates PatternFlowGtpv1TeidCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1TeidCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv1TeidCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv1TeidCounter
	SetStart(value uint32) PatternFlowGtpv1TeidCounter
	// HasStart checks if Start has been set in PatternFlowGtpv1TeidCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv1TeidCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv1TeidCounter
	SetStep(value uint32) PatternFlowGtpv1TeidCounter
	// HasStep checks if Step has been set in PatternFlowGtpv1TeidCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv1TeidCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv1TeidCounter
	SetCount(value uint32) PatternFlowGtpv1TeidCounter
	// HasCount checks if Count has been set in PatternFlowGtpv1TeidCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1TeidCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1TeidCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv1TeidCounter object
func (obj *patternFlowGtpv1TeidCounter) SetStart(value uint32) PatternFlowGtpv1TeidCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1TeidCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1TeidCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv1TeidCounter object
func (obj *patternFlowGtpv1TeidCounter) SetStep(value uint32) PatternFlowGtpv1TeidCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1TeidCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1TeidCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv1TeidCounter object
func (obj *patternFlowGtpv1TeidCounter) SetCount(value uint32) PatternFlowGtpv1TeidCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv1TeidCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowGtpv1TeidCounter) setDefault() {
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
