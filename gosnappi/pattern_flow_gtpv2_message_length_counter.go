package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2MessageLengthCounter *****
type patternFlowGtpv2MessageLengthCounter struct {
	validation
	obj          *otg.PatternFlowGtpv2MessageLengthCounter
	marshaller   marshalPatternFlowGtpv2MessageLengthCounter
	unMarshaller unMarshalPatternFlowGtpv2MessageLengthCounter
}

func NewPatternFlowGtpv2MessageLengthCounter() PatternFlowGtpv2MessageLengthCounter {
	obj := patternFlowGtpv2MessageLengthCounter{obj: &otg.PatternFlowGtpv2MessageLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2MessageLengthCounter) msg() *otg.PatternFlowGtpv2MessageLengthCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2MessageLengthCounter) setMsg(msg *otg.PatternFlowGtpv2MessageLengthCounter) PatternFlowGtpv2MessageLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2MessageLengthCounter struct {
	obj *patternFlowGtpv2MessageLengthCounter
}

type marshalPatternFlowGtpv2MessageLengthCounter interface {
	// ToProto marshals PatternFlowGtpv2MessageLengthCounter to protobuf object *otg.PatternFlowGtpv2MessageLengthCounter
	ToProto() (*otg.PatternFlowGtpv2MessageLengthCounter, error)
	// ToPbText marshals PatternFlowGtpv2MessageLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2MessageLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2MessageLengthCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv2MessageLengthCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv2MessageLengthCounter struct {
	obj *patternFlowGtpv2MessageLengthCounter
}

type unMarshalPatternFlowGtpv2MessageLengthCounter interface {
	// FromProto unmarshals PatternFlowGtpv2MessageLengthCounter from protobuf object *otg.PatternFlowGtpv2MessageLengthCounter
	FromProto(msg *otg.PatternFlowGtpv2MessageLengthCounter) (PatternFlowGtpv2MessageLengthCounter, error)
	// FromPbText unmarshals PatternFlowGtpv2MessageLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2MessageLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2MessageLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2MessageLengthCounter) Marshal() marshalPatternFlowGtpv2MessageLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2MessageLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2MessageLengthCounter) Unmarshal() unMarshalPatternFlowGtpv2MessageLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2MessageLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2MessageLengthCounter) ToProto() (*otg.PatternFlowGtpv2MessageLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2MessageLengthCounter) FromProto(msg *otg.PatternFlowGtpv2MessageLengthCounter) (PatternFlowGtpv2MessageLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2MessageLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2MessageLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2MessageLengthCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv2MessageLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv2MessageLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2MessageLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2MessageLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2MessageLengthCounter) Clone() (PatternFlowGtpv2MessageLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2MessageLengthCounter()
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

// PatternFlowGtpv2MessageLengthCounter is integer counter pattern
type PatternFlowGtpv2MessageLengthCounter interface {
	Validation
	// msg marshals PatternFlowGtpv2MessageLengthCounter to protobuf object *otg.PatternFlowGtpv2MessageLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2MessageLengthCounter
	// setMsg unmarshals PatternFlowGtpv2MessageLengthCounter from protobuf object *otg.PatternFlowGtpv2MessageLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2MessageLengthCounter) PatternFlowGtpv2MessageLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2MessageLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2MessageLengthCounter
	// validate validates PatternFlowGtpv2MessageLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2MessageLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv2MessageLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv2MessageLengthCounter
	SetStart(value uint32) PatternFlowGtpv2MessageLengthCounter
	// HasStart checks if Start has been set in PatternFlowGtpv2MessageLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv2MessageLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv2MessageLengthCounter
	SetStep(value uint32) PatternFlowGtpv2MessageLengthCounter
	// HasStep checks if Step has been set in PatternFlowGtpv2MessageLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv2MessageLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv2MessageLengthCounter
	SetCount(value uint32) PatternFlowGtpv2MessageLengthCounter
	// HasCount checks if Count has been set in PatternFlowGtpv2MessageLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2MessageLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2MessageLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv2MessageLengthCounter object
func (obj *patternFlowGtpv2MessageLengthCounter) SetStart(value uint32) PatternFlowGtpv2MessageLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2MessageLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2MessageLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv2MessageLengthCounter object
func (obj *patternFlowGtpv2MessageLengthCounter) SetStep(value uint32) PatternFlowGtpv2MessageLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2MessageLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2MessageLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv2MessageLengthCounter object
func (obj *patternFlowGtpv2MessageLengthCounter) SetCount(value uint32) PatternFlowGtpv2MessageLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv2MessageLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageLengthCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageLengthCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageLengthCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv2MessageLengthCounter) setDefault() {
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
