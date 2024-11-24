package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1MessageLengthCounter *****
type patternFlowGtpv1MessageLengthCounter struct {
	validation
	obj          *otg.PatternFlowGtpv1MessageLengthCounter
	marshaller   marshalPatternFlowGtpv1MessageLengthCounter
	unMarshaller unMarshalPatternFlowGtpv1MessageLengthCounter
}

func NewPatternFlowGtpv1MessageLengthCounter() PatternFlowGtpv1MessageLengthCounter {
	obj := patternFlowGtpv1MessageLengthCounter{obj: &otg.PatternFlowGtpv1MessageLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1MessageLengthCounter) msg() *otg.PatternFlowGtpv1MessageLengthCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1MessageLengthCounter) setMsg(msg *otg.PatternFlowGtpv1MessageLengthCounter) PatternFlowGtpv1MessageLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1MessageLengthCounter struct {
	obj *patternFlowGtpv1MessageLengthCounter
}

type marshalPatternFlowGtpv1MessageLengthCounter interface {
	// ToProto marshals PatternFlowGtpv1MessageLengthCounter to protobuf object *otg.PatternFlowGtpv1MessageLengthCounter
	ToProto() (*otg.PatternFlowGtpv1MessageLengthCounter, error)
	// ToPbText marshals PatternFlowGtpv1MessageLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1MessageLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1MessageLengthCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1MessageLengthCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1MessageLengthCounter struct {
	obj *patternFlowGtpv1MessageLengthCounter
}

type unMarshalPatternFlowGtpv1MessageLengthCounter interface {
	// FromProto unmarshals PatternFlowGtpv1MessageLengthCounter from protobuf object *otg.PatternFlowGtpv1MessageLengthCounter
	FromProto(msg *otg.PatternFlowGtpv1MessageLengthCounter) (PatternFlowGtpv1MessageLengthCounter, error)
	// FromPbText unmarshals PatternFlowGtpv1MessageLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1MessageLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1MessageLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1MessageLengthCounter) Marshal() marshalPatternFlowGtpv1MessageLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1MessageLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1MessageLengthCounter) Unmarshal() unMarshalPatternFlowGtpv1MessageLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1MessageLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1MessageLengthCounter) ToProto() (*otg.PatternFlowGtpv1MessageLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1MessageLengthCounter) FromProto(msg *otg.PatternFlowGtpv1MessageLengthCounter) (PatternFlowGtpv1MessageLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1MessageLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1MessageLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1MessageLengthCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv1MessageLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv1MessageLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1MessageLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1MessageLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1MessageLengthCounter) Clone() (PatternFlowGtpv1MessageLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1MessageLengthCounter()
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

// PatternFlowGtpv1MessageLengthCounter is integer counter pattern
type PatternFlowGtpv1MessageLengthCounter interface {
	Validation
	// msg marshals PatternFlowGtpv1MessageLengthCounter to protobuf object *otg.PatternFlowGtpv1MessageLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1MessageLengthCounter
	// setMsg unmarshals PatternFlowGtpv1MessageLengthCounter from protobuf object *otg.PatternFlowGtpv1MessageLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1MessageLengthCounter) PatternFlowGtpv1MessageLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1MessageLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1MessageLengthCounter
	// validate validates PatternFlowGtpv1MessageLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1MessageLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv1MessageLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv1MessageLengthCounter
	SetStart(value uint32) PatternFlowGtpv1MessageLengthCounter
	// HasStart checks if Start has been set in PatternFlowGtpv1MessageLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv1MessageLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv1MessageLengthCounter
	SetStep(value uint32) PatternFlowGtpv1MessageLengthCounter
	// HasStep checks if Step has been set in PatternFlowGtpv1MessageLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv1MessageLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv1MessageLengthCounter
	SetCount(value uint32) PatternFlowGtpv1MessageLengthCounter
	// HasCount checks if Count has been set in PatternFlowGtpv1MessageLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1MessageLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1MessageLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv1MessageLengthCounter object
func (obj *patternFlowGtpv1MessageLengthCounter) SetStart(value uint32) PatternFlowGtpv1MessageLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1MessageLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1MessageLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv1MessageLengthCounter object
func (obj *patternFlowGtpv1MessageLengthCounter) SetStep(value uint32) PatternFlowGtpv1MessageLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1MessageLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1MessageLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv1MessageLengthCounter object
func (obj *patternFlowGtpv1MessageLengthCounter) SetCount(value uint32) PatternFlowGtpv1MessageLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv1MessageLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageLengthCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageLengthCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageLengthCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv1MessageLengthCounter) setDefault() {
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
