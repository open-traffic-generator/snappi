package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowUdpLengthCounter *****
type patternFlowUdpLengthCounter struct {
	validation
	obj          *otg.PatternFlowUdpLengthCounter
	marshaller   marshalPatternFlowUdpLengthCounter
	unMarshaller unMarshalPatternFlowUdpLengthCounter
}

func NewPatternFlowUdpLengthCounter() PatternFlowUdpLengthCounter {
	obj := patternFlowUdpLengthCounter{obj: &otg.PatternFlowUdpLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowUdpLengthCounter) msg() *otg.PatternFlowUdpLengthCounter {
	return obj.obj
}

func (obj *patternFlowUdpLengthCounter) setMsg(msg *otg.PatternFlowUdpLengthCounter) PatternFlowUdpLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowUdpLengthCounter struct {
	obj *patternFlowUdpLengthCounter
}

type marshalPatternFlowUdpLengthCounter interface {
	// ToProto marshals PatternFlowUdpLengthCounter to protobuf object *otg.PatternFlowUdpLengthCounter
	ToProto() (*otg.PatternFlowUdpLengthCounter, error)
	// ToPbText marshals PatternFlowUdpLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowUdpLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowUdpLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowUdpLengthCounter struct {
	obj *patternFlowUdpLengthCounter
}

type unMarshalPatternFlowUdpLengthCounter interface {
	// FromProto unmarshals PatternFlowUdpLengthCounter from protobuf object *otg.PatternFlowUdpLengthCounter
	FromProto(msg *otg.PatternFlowUdpLengthCounter) (PatternFlowUdpLengthCounter, error)
	// FromPbText unmarshals PatternFlowUdpLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowUdpLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowUdpLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowUdpLengthCounter) Marshal() marshalPatternFlowUdpLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowUdpLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowUdpLengthCounter) Unmarshal() unMarshalPatternFlowUdpLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowUdpLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowUdpLengthCounter) ToProto() (*otg.PatternFlowUdpLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowUdpLengthCounter) FromProto(msg *otg.PatternFlowUdpLengthCounter) (PatternFlowUdpLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowUdpLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowUdpLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowUdpLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowUdpLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowUdpLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowUdpLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowUdpLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowUdpLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowUdpLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowUdpLengthCounter) Clone() (PatternFlowUdpLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowUdpLengthCounter()
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

// PatternFlowUdpLengthCounter is integer counter pattern
type PatternFlowUdpLengthCounter interface {
	Validation
	// msg marshals PatternFlowUdpLengthCounter to protobuf object *otg.PatternFlowUdpLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowUdpLengthCounter
	// setMsg unmarshals PatternFlowUdpLengthCounter from protobuf object *otg.PatternFlowUdpLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowUdpLengthCounter) PatternFlowUdpLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowUdpLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowUdpLengthCounter
	// validate validates PatternFlowUdpLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowUdpLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowUdpLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowUdpLengthCounter
	SetStart(value uint32) PatternFlowUdpLengthCounter
	// HasStart checks if Start has been set in PatternFlowUdpLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowUdpLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowUdpLengthCounter
	SetStep(value uint32) PatternFlowUdpLengthCounter
	// HasStep checks if Step has been set in PatternFlowUdpLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowUdpLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowUdpLengthCounter
	SetCount(value uint32) PatternFlowUdpLengthCounter
	// HasCount checks if Count has been set in PatternFlowUdpLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowUdpLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowUdpLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowUdpLengthCounter object
func (obj *patternFlowUdpLengthCounter) SetStart(value uint32) PatternFlowUdpLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowUdpLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowUdpLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowUdpLengthCounter object
func (obj *patternFlowUdpLengthCounter) SetStep(value uint32) PatternFlowUdpLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowUdpLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowUdpLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowUdpLengthCounter object
func (obj *patternFlowUdpLengthCounter) SetCount(value uint32) PatternFlowUdpLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowUdpLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpLengthCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpLengthCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpLengthCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowUdpLengthCounter) setDefault() {
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
