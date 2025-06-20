package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpTargetProtocolAddrCounter *****
type patternFlowArpTargetProtocolAddrCounter struct {
	validation
	obj          *otg.PatternFlowArpTargetProtocolAddrCounter
	marshaller   marshalPatternFlowArpTargetProtocolAddrCounter
	unMarshaller unMarshalPatternFlowArpTargetProtocolAddrCounter
}

func NewPatternFlowArpTargetProtocolAddrCounter() PatternFlowArpTargetProtocolAddrCounter {
	obj := patternFlowArpTargetProtocolAddrCounter{obj: &otg.PatternFlowArpTargetProtocolAddrCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpTargetProtocolAddrCounter) msg() *otg.PatternFlowArpTargetProtocolAddrCounter {
	return obj.obj
}

func (obj *patternFlowArpTargetProtocolAddrCounter) setMsg(msg *otg.PatternFlowArpTargetProtocolAddrCounter) PatternFlowArpTargetProtocolAddrCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpTargetProtocolAddrCounter struct {
	obj *patternFlowArpTargetProtocolAddrCounter
}

type marshalPatternFlowArpTargetProtocolAddrCounter interface {
	// ToProto marshals PatternFlowArpTargetProtocolAddrCounter to protobuf object *otg.PatternFlowArpTargetProtocolAddrCounter
	ToProto() (*otg.PatternFlowArpTargetProtocolAddrCounter, error)
	// ToPbText marshals PatternFlowArpTargetProtocolAddrCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpTargetProtocolAddrCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpTargetProtocolAddrCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpTargetProtocolAddrCounter struct {
	obj *patternFlowArpTargetProtocolAddrCounter
}

type unMarshalPatternFlowArpTargetProtocolAddrCounter interface {
	// FromProto unmarshals PatternFlowArpTargetProtocolAddrCounter from protobuf object *otg.PatternFlowArpTargetProtocolAddrCounter
	FromProto(msg *otg.PatternFlowArpTargetProtocolAddrCounter) (PatternFlowArpTargetProtocolAddrCounter, error)
	// FromPbText unmarshals PatternFlowArpTargetProtocolAddrCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpTargetProtocolAddrCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpTargetProtocolAddrCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpTargetProtocolAddrCounter) Marshal() marshalPatternFlowArpTargetProtocolAddrCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpTargetProtocolAddrCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpTargetProtocolAddrCounter) Unmarshal() unMarshalPatternFlowArpTargetProtocolAddrCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpTargetProtocolAddrCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpTargetProtocolAddrCounter) ToProto() (*otg.PatternFlowArpTargetProtocolAddrCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpTargetProtocolAddrCounter) FromProto(msg *otg.PatternFlowArpTargetProtocolAddrCounter) (PatternFlowArpTargetProtocolAddrCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpTargetProtocolAddrCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetProtocolAddrCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpTargetProtocolAddrCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetProtocolAddrCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpTargetProtocolAddrCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetProtocolAddrCounter) FromJson(value string) error {
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

func (obj *patternFlowArpTargetProtocolAddrCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpTargetProtocolAddrCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpTargetProtocolAddrCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpTargetProtocolAddrCounter) Clone() (PatternFlowArpTargetProtocolAddrCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpTargetProtocolAddrCounter()
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

// PatternFlowArpTargetProtocolAddrCounter is ipv4 counter pattern
type PatternFlowArpTargetProtocolAddrCounter interface {
	Validation
	// msg marshals PatternFlowArpTargetProtocolAddrCounter to protobuf object *otg.PatternFlowArpTargetProtocolAddrCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowArpTargetProtocolAddrCounter
	// setMsg unmarshals PatternFlowArpTargetProtocolAddrCounter from protobuf object *otg.PatternFlowArpTargetProtocolAddrCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpTargetProtocolAddrCounter) PatternFlowArpTargetProtocolAddrCounter
	// provides marshal interface
	Marshal() marshalPatternFlowArpTargetProtocolAddrCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpTargetProtocolAddrCounter
	// validate validates PatternFlowArpTargetProtocolAddrCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpTargetProtocolAddrCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowArpTargetProtocolAddrCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowArpTargetProtocolAddrCounter
	SetStart(value string) PatternFlowArpTargetProtocolAddrCounter
	// HasStart checks if Start has been set in PatternFlowArpTargetProtocolAddrCounter
	HasStart() bool
	// Step returns string, set in PatternFlowArpTargetProtocolAddrCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowArpTargetProtocolAddrCounter
	SetStep(value string) PatternFlowArpTargetProtocolAddrCounter
	// HasStep checks if Step has been set in PatternFlowArpTargetProtocolAddrCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowArpTargetProtocolAddrCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowArpTargetProtocolAddrCounter
	SetCount(value uint32) PatternFlowArpTargetProtocolAddrCounter
	// HasCount checks if Count has been set in PatternFlowArpTargetProtocolAddrCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowArpTargetProtocolAddrCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowArpTargetProtocolAddrCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowArpTargetProtocolAddrCounter object
func (obj *patternFlowArpTargetProtocolAddrCounter) SetStart(value string) PatternFlowArpTargetProtocolAddrCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowArpTargetProtocolAddrCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowArpTargetProtocolAddrCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowArpTargetProtocolAddrCounter object
func (obj *patternFlowArpTargetProtocolAddrCounter) SetStep(value string) PatternFlowArpTargetProtocolAddrCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpTargetProtocolAddrCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpTargetProtocolAddrCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowArpTargetProtocolAddrCounter object
func (obj *patternFlowArpTargetProtocolAddrCounter) SetCount(value uint32) PatternFlowArpTargetProtocolAddrCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowArpTargetProtocolAddrCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetProtocolAddrCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetProtocolAddrCounter.Step"))
		}

	}

}

func (obj *patternFlowArpTargetProtocolAddrCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("0.0.0.0")
	}
	if obj.obj.Step == nil {
		obj.SetStep("0.0.0.1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
