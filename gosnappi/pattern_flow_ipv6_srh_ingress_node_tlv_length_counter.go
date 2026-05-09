package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHIngressNodeTlvLengthCounter *****
type patternFlowIpv6SRHIngressNodeTlvLengthCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	marshaller   marshalPatternFlowIpv6SRHIngressNodeTlvLengthCounter
	unMarshaller unMarshalPatternFlowIpv6SRHIngressNodeTlvLengthCounter
}

func NewPatternFlowIpv6SRHIngressNodeTlvLengthCounter() PatternFlowIpv6SRHIngressNodeTlvLengthCounter {
	obj := patternFlowIpv6SRHIngressNodeTlvLengthCounter{obj: &otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) msg() *otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) setMsg(msg *otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter) PatternFlowIpv6SRHIngressNodeTlvLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHIngressNodeTlvLengthCounter struct {
	obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter
}

type marshalPatternFlowIpv6SRHIngressNodeTlvLengthCounter interface {
	// ToProto marshals PatternFlowIpv6SRHIngressNodeTlvLengthCounter to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHIngressNodeTlvLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHIngressNodeTlvLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHIngressNodeTlvLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHIngressNodeTlvLengthCounter struct {
	obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter
}

type unMarshalPatternFlowIpv6SRHIngressNodeTlvLengthCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHIngressNodeTlvLengthCounter from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter) (PatternFlowIpv6SRHIngressNodeTlvLengthCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHIngressNodeTlvLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHIngressNodeTlvLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHIngressNodeTlvLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHIngressNodeTlvLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHIngressNodeTlvLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvLengthCounter) ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvLengthCounter) FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter) (PatternFlowIpv6SRHIngressNodeTlvLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) Clone() (PatternFlowIpv6SRHIngressNodeTlvLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHIngressNodeTlvLengthCounter()
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

// PatternFlowIpv6SRHIngressNodeTlvLengthCounter is integer counter pattern
type PatternFlowIpv6SRHIngressNodeTlvLengthCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHIngressNodeTlvLengthCounter to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	// setMsg unmarshals PatternFlowIpv6SRHIngressNodeTlvLengthCounter from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHIngressNodeTlvLengthCounter) PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvLengthCounter
	// validate validates PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHIngressNodeTlvLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	SetStart(value uint32) PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	SetStep(value uint32) PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	SetCount(value uint32) PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHIngressNodeTlvLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHIngressNodeTlvLengthCounter object
func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) SetStart(value uint32) PatternFlowIpv6SRHIngressNodeTlvLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHIngressNodeTlvLengthCounter object
func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) SetStep(value uint32) PatternFlowIpv6SRHIngressNodeTlvLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHIngressNodeTlvLengthCounter object
func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) SetCount(value uint32) PatternFlowIpv6SRHIngressNodeTlvLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SRHIngressNodeTlvLengthCounter) setDefault() {
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
