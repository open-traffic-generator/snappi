package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6UsidDstLocatorCounter *****
type patternFlowIpv6UsidDstLocatorCounter struct {
	validation
	obj          *otg.PatternFlowIpv6UsidDstLocatorCounter
	marshaller   marshalPatternFlowIpv6UsidDstLocatorCounter
	unMarshaller unMarshalPatternFlowIpv6UsidDstLocatorCounter
}

func NewPatternFlowIpv6UsidDstLocatorCounter() PatternFlowIpv6UsidDstLocatorCounter {
	obj := patternFlowIpv6UsidDstLocatorCounter{obj: &otg.PatternFlowIpv6UsidDstLocatorCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6UsidDstLocatorCounter) msg() *otg.PatternFlowIpv6UsidDstLocatorCounter {
	return obj.obj
}

func (obj *patternFlowIpv6UsidDstLocatorCounter) setMsg(msg *otg.PatternFlowIpv6UsidDstLocatorCounter) PatternFlowIpv6UsidDstLocatorCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6UsidDstLocatorCounter struct {
	obj *patternFlowIpv6UsidDstLocatorCounter
}

type marshalPatternFlowIpv6UsidDstLocatorCounter interface {
	// ToProto marshals PatternFlowIpv6UsidDstLocatorCounter to protobuf object *otg.PatternFlowIpv6UsidDstLocatorCounter
	ToProto() (*otg.PatternFlowIpv6UsidDstLocatorCounter, error)
	// ToPbText marshals PatternFlowIpv6UsidDstLocatorCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6UsidDstLocatorCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6UsidDstLocatorCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6UsidDstLocatorCounter struct {
	obj *patternFlowIpv6UsidDstLocatorCounter
}

type unMarshalPatternFlowIpv6UsidDstLocatorCounter interface {
	// FromProto unmarshals PatternFlowIpv6UsidDstLocatorCounter from protobuf object *otg.PatternFlowIpv6UsidDstLocatorCounter
	FromProto(msg *otg.PatternFlowIpv6UsidDstLocatorCounter) (PatternFlowIpv6UsidDstLocatorCounter, error)
	// FromPbText unmarshals PatternFlowIpv6UsidDstLocatorCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6UsidDstLocatorCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6UsidDstLocatorCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6UsidDstLocatorCounter) Marshal() marshalPatternFlowIpv6UsidDstLocatorCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6UsidDstLocatorCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6UsidDstLocatorCounter) Unmarshal() unMarshalPatternFlowIpv6UsidDstLocatorCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6UsidDstLocatorCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6UsidDstLocatorCounter) ToProto() (*otg.PatternFlowIpv6UsidDstLocatorCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6UsidDstLocatorCounter) FromProto(msg *otg.PatternFlowIpv6UsidDstLocatorCounter) (PatternFlowIpv6UsidDstLocatorCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6UsidDstLocatorCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6UsidDstLocatorCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6UsidDstLocatorCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6UsidDstLocatorCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6UsidDstLocatorCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6UsidDstLocatorCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6UsidDstLocatorCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6UsidDstLocatorCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6UsidDstLocatorCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6UsidDstLocatorCounter) Clone() (PatternFlowIpv6UsidDstLocatorCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6UsidDstLocatorCounter()
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

// PatternFlowIpv6UsidDstLocatorCounter is ipv6 counter pattern
type PatternFlowIpv6UsidDstLocatorCounter interface {
	Validation
	// msg marshals PatternFlowIpv6UsidDstLocatorCounter to protobuf object *otg.PatternFlowIpv6UsidDstLocatorCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6UsidDstLocatorCounter
	// setMsg unmarshals PatternFlowIpv6UsidDstLocatorCounter from protobuf object *otg.PatternFlowIpv6UsidDstLocatorCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6UsidDstLocatorCounter) PatternFlowIpv6UsidDstLocatorCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6UsidDstLocatorCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6UsidDstLocatorCounter
	// validate validates PatternFlowIpv6UsidDstLocatorCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6UsidDstLocatorCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv6UsidDstLocatorCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv6UsidDstLocatorCounter
	SetStart(value string) PatternFlowIpv6UsidDstLocatorCounter
	// HasStart checks if Start has been set in PatternFlowIpv6UsidDstLocatorCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv6UsidDstLocatorCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv6UsidDstLocatorCounter
	SetStep(value string) PatternFlowIpv6UsidDstLocatorCounter
	// HasStep checks if Step has been set in PatternFlowIpv6UsidDstLocatorCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6UsidDstLocatorCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6UsidDstLocatorCounter
	SetCount(value uint32) PatternFlowIpv6UsidDstLocatorCounter
	// HasCount checks if Count has been set in PatternFlowIpv6UsidDstLocatorCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6UsidDstLocatorCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6UsidDstLocatorCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv6UsidDstLocatorCounter object
func (obj *patternFlowIpv6UsidDstLocatorCounter) SetStart(value string) PatternFlowIpv6UsidDstLocatorCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6UsidDstLocatorCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6UsidDstLocatorCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv6UsidDstLocatorCounter object
func (obj *patternFlowIpv6UsidDstLocatorCounter) SetStep(value string) PatternFlowIpv6UsidDstLocatorCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6UsidDstLocatorCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6UsidDstLocatorCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6UsidDstLocatorCounter object
func (obj *patternFlowIpv6UsidDstLocatorCounter) SetCount(value uint32) PatternFlowIpv6UsidDstLocatorCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6UsidDstLocatorCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6UsidDstLocatorCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6UsidDstLocatorCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv6UsidDstLocatorCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("::0")
	}
	if obj.obj.Step == nil {
		obj.SetStep("::1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
