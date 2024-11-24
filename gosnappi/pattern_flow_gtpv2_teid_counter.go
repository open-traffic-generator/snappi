package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2TeidCounter *****
type patternFlowGtpv2TeidCounter struct {
	validation
	obj          *otg.PatternFlowGtpv2TeidCounter
	marshaller   marshalPatternFlowGtpv2TeidCounter
	unMarshaller unMarshalPatternFlowGtpv2TeidCounter
}

func NewPatternFlowGtpv2TeidCounter() PatternFlowGtpv2TeidCounter {
	obj := patternFlowGtpv2TeidCounter{obj: &otg.PatternFlowGtpv2TeidCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2TeidCounter) msg() *otg.PatternFlowGtpv2TeidCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2TeidCounter) setMsg(msg *otg.PatternFlowGtpv2TeidCounter) PatternFlowGtpv2TeidCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2TeidCounter struct {
	obj *patternFlowGtpv2TeidCounter
}

type marshalPatternFlowGtpv2TeidCounter interface {
	// ToProto marshals PatternFlowGtpv2TeidCounter to protobuf object *otg.PatternFlowGtpv2TeidCounter
	ToProto() (*otg.PatternFlowGtpv2TeidCounter, error)
	// ToPbText marshals PatternFlowGtpv2TeidCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2TeidCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2TeidCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv2TeidCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv2TeidCounter struct {
	obj *patternFlowGtpv2TeidCounter
}

type unMarshalPatternFlowGtpv2TeidCounter interface {
	// FromProto unmarshals PatternFlowGtpv2TeidCounter from protobuf object *otg.PatternFlowGtpv2TeidCounter
	FromProto(msg *otg.PatternFlowGtpv2TeidCounter) (PatternFlowGtpv2TeidCounter, error)
	// FromPbText unmarshals PatternFlowGtpv2TeidCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2TeidCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2TeidCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2TeidCounter) Marshal() marshalPatternFlowGtpv2TeidCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2TeidCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2TeidCounter) Unmarshal() unMarshalPatternFlowGtpv2TeidCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2TeidCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2TeidCounter) ToProto() (*otg.PatternFlowGtpv2TeidCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2TeidCounter) FromProto(msg *otg.PatternFlowGtpv2TeidCounter) (PatternFlowGtpv2TeidCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2TeidCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2TeidCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2TeidCounter) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowGtpv2TeidCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv2TeidCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2TeidCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2TeidCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2TeidCounter) Clone() (PatternFlowGtpv2TeidCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2TeidCounter()
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

// PatternFlowGtpv2TeidCounter is integer counter pattern
type PatternFlowGtpv2TeidCounter interface {
	Validation
	// msg marshals PatternFlowGtpv2TeidCounter to protobuf object *otg.PatternFlowGtpv2TeidCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2TeidCounter
	// setMsg unmarshals PatternFlowGtpv2TeidCounter from protobuf object *otg.PatternFlowGtpv2TeidCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2TeidCounter) PatternFlowGtpv2TeidCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2TeidCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2TeidCounter
	// validate validates PatternFlowGtpv2TeidCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2TeidCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv2TeidCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv2TeidCounter
	SetStart(value uint32) PatternFlowGtpv2TeidCounter
	// HasStart checks if Start has been set in PatternFlowGtpv2TeidCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv2TeidCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv2TeidCounter
	SetStep(value uint32) PatternFlowGtpv2TeidCounter
	// HasStep checks if Step has been set in PatternFlowGtpv2TeidCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv2TeidCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv2TeidCounter
	SetCount(value uint32) PatternFlowGtpv2TeidCounter
	// HasCount checks if Count has been set in PatternFlowGtpv2TeidCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2TeidCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2TeidCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv2TeidCounter object
func (obj *patternFlowGtpv2TeidCounter) SetStart(value uint32) PatternFlowGtpv2TeidCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2TeidCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2TeidCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv2TeidCounter object
func (obj *patternFlowGtpv2TeidCounter) SetStep(value uint32) PatternFlowGtpv2TeidCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2TeidCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2TeidCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv2TeidCounter object
func (obj *patternFlowGtpv2TeidCounter) SetCount(value uint32) PatternFlowGtpv2TeidCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv2TeidCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowGtpv2TeidCounter) setDefault() {
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
