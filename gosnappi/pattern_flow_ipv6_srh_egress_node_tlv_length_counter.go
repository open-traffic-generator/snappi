package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHEgressNodeTlvLengthCounter *****
type patternFlowIpv6SRHEgressNodeTlvLengthCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	marshaller   marshalPatternFlowIpv6SRHEgressNodeTlvLengthCounter
	unMarshaller unMarshalPatternFlowIpv6SRHEgressNodeTlvLengthCounter
}

func NewPatternFlowIpv6SRHEgressNodeTlvLengthCounter() PatternFlowIpv6SRHEgressNodeTlvLengthCounter {
	obj := patternFlowIpv6SRHEgressNodeTlvLengthCounter{obj: &otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) msg() *otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) setMsg(msg *otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter) PatternFlowIpv6SRHEgressNodeTlvLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHEgressNodeTlvLengthCounter struct {
	obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter
}

type marshalPatternFlowIpv6SRHEgressNodeTlvLengthCounter interface {
	// ToProto marshals PatternFlowIpv6SRHEgressNodeTlvLengthCounter to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHEgressNodeTlvLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHEgressNodeTlvLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHEgressNodeTlvLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHEgressNodeTlvLengthCounter struct {
	obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter
}

type unMarshalPatternFlowIpv6SRHEgressNodeTlvLengthCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHEgressNodeTlvLengthCounter from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter) (PatternFlowIpv6SRHEgressNodeTlvLengthCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHEgressNodeTlvLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHEgressNodeTlvLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHEgressNodeTlvLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHEgressNodeTlvLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHEgressNodeTlvLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvLengthCounter) ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvLengthCounter) FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter) (PatternFlowIpv6SRHEgressNodeTlvLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) Clone() (PatternFlowIpv6SRHEgressNodeTlvLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHEgressNodeTlvLengthCounter()
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

// PatternFlowIpv6SRHEgressNodeTlvLengthCounter is integer counter pattern
type PatternFlowIpv6SRHEgressNodeTlvLengthCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHEgressNodeTlvLengthCounter to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	// setMsg unmarshals PatternFlowIpv6SRHEgressNodeTlvLengthCounter from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHEgressNodeTlvLengthCounter) PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvLengthCounter
	// validate validates PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHEgressNodeTlvLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	SetStart(value uint32) PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	SetStep(value uint32) PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	SetCount(value uint32) PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHEgressNodeTlvLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHEgressNodeTlvLengthCounter object
func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) SetStart(value uint32) PatternFlowIpv6SRHEgressNodeTlvLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHEgressNodeTlvLengthCounter object
func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) SetStep(value uint32) PatternFlowIpv6SRHEgressNodeTlvLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHEgressNodeTlvLengthCounter object
func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) SetCount(value uint32) PatternFlowIpv6SRHEgressNodeTlvLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SRHEgressNodeTlvLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(18)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
