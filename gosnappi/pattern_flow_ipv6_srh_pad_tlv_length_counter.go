package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPadTlvLengthCounter *****
type patternFlowIpv6SRHPadTlvLengthCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHPadTlvLengthCounter
	marshaller   marshalPatternFlowIpv6SRHPadTlvLengthCounter
	unMarshaller unMarshalPatternFlowIpv6SRHPadTlvLengthCounter
}

func NewPatternFlowIpv6SRHPadTlvLengthCounter() PatternFlowIpv6SRHPadTlvLengthCounter {
	obj := patternFlowIpv6SRHPadTlvLengthCounter{obj: &otg.PatternFlowIpv6SRHPadTlvLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPadTlvLengthCounter) msg() *otg.PatternFlowIpv6SRHPadTlvLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPadTlvLengthCounter) setMsg(msg *otg.PatternFlowIpv6SRHPadTlvLengthCounter) PatternFlowIpv6SRHPadTlvLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPadTlvLengthCounter struct {
	obj *patternFlowIpv6SRHPadTlvLengthCounter
}

type marshalPatternFlowIpv6SRHPadTlvLengthCounter interface {
	// ToProto marshals PatternFlowIpv6SRHPadTlvLengthCounter to protobuf object *otg.PatternFlowIpv6SRHPadTlvLengthCounter
	ToProto() (*otg.PatternFlowIpv6SRHPadTlvLengthCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHPadTlvLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPadTlvLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPadTlvLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPadTlvLengthCounter struct {
	obj *patternFlowIpv6SRHPadTlvLengthCounter
}

type unMarshalPatternFlowIpv6SRHPadTlvLengthCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHPadTlvLengthCounter from protobuf object *otg.PatternFlowIpv6SRHPadTlvLengthCounter
	FromProto(msg *otg.PatternFlowIpv6SRHPadTlvLengthCounter) (PatternFlowIpv6SRHPadTlvLengthCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPadTlvLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPadTlvLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPadTlvLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPadTlvLengthCounter) Marshal() marshalPatternFlowIpv6SRHPadTlvLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPadTlvLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPadTlvLengthCounter) Unmarshal() unMarshalPatternFlowIpv6SRHPadTlvLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPadTlvLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPadTlvLengthCounter) ToProto() (*otg.PatternFlowIpv6SRHPadTlvLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPadTlvLengthCounter) FromProto(msg *otg.PatternFlowIpv6SRHPadTlvLengthCounter) (PatternFlowIpv6SRHPadTlvLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPadTlvLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPadTlvLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPadTlvLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPadTlvLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPadTlvLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPadTlvLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPadTlvLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPadTlvLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPadTlvLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPadTlvLengthCounter) Clone() (PatternFlowIpv6SRHPadTlvLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPadTlvLengthCounter()
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

// PatternFlowIpv6SRHPadTlvLengthCounter is integer counter pattern
type PatternFlowIpv6SRHPadTlvLengthCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPadTlvLengthCounter to protobuf object *otg.PatternFlowIpv6SRHPadTlvLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPadTlvLengthCounter
	// setMsg unmarshals PatternFlowIpv6SRHPadTlvLengthCounter from protobuf object *otg.PatternFlowIpv6SRHPadTlvLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPadTlvLengthCounter) PatternFlowIpv6SRHPadTlvLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPadTlvLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPadTlvLengthCounter
	// validate validates PatternFlowIpv6SRHPadTlvLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPadTlvLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHPadTlvLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHPadTlvLengthCounter
	SetStart(value uint32) PatternFlowIpv6SRHPadTlvLengthCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHPadTlvLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHPadTlvLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHPadTlvLengthCounter
	SetStep(value uint32) PatternFlowIpv6SRHPadTlvLengthCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHPadTlvLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHPadTlvLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHPadTlvLengthCounter
	SetCount(value uint32) PatternFlowIpv6SRHPadTlvLengthCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHPadTlvLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPadTlvLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPadTlvLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHPadTlvLengthCounter object
func (obj *patternFlowIpv6SRHPadTlvLengthCounter) SetStart(value uint32) PatternFlowIpv6SRHPadTlvLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPadTlvLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPadTlvLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHPadTlvLengthCounter object
func (obj *patternFlowIpv6SRHPadTlvLengthCounter) SetStep(value uint32) PatternFlowIpv6SRHPadTlvLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPadTlvLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPadTlvLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHPadTlvLengthCounter object
func (obj *patternFlowIpv6SRHPadTlvLengthCounter) SetCount(value uint32) PatternFlowIpv6SRHPadTlvLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHPadTlvLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPadTlvLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPadTlvLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPadTlvLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SRHPadTlvLengthCounter) setDefault() {
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
