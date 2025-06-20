package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4SrcCounter *****
type patternFlowIpv4SrcCounter struct {
	validation
	obj          *otg.PatternFlowIpv4SrcCounter
	marshaller   marshalPatternFlowIpv4SrcCounter
	unMarshaller unMarshalPatternFlowIpv4SrcCounter
}

func NewPatternFlowIpv4SrcCounter() PatternFlowIpv4SrcCounter {
	obj := patternFlowIpv4SrcCounter{obj: &otg.PatternFlowIpv4SrcCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4SrcCounter) msg() *otg.PatternFlowIpv4SrcCounter {
	return obj.obj
}

func (obj *patternFlowIpv4SrcCounter) setMsg(msg *otg.PatternFlowIpv4SrcCounter) PatternFlowIpv4SrcCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4SrcCounter struct {
	obj *patternFlowIpv4SrcCounter
}

type marshalPatternFlowIpv4SrcCounter interface {
	// ToProto marshals PatternFlowIpv4SrcCounter to protobuf object *otg.PatternFlowIpv4SrcCounter
	ToProto() (*otg.PatternFlowIpv4SrcCounter, error)
	// ToPbText marshals PatternFlowIpv4SrcCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4SrcCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4SrcCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4SrcCounter struct {
	obj *patternFlowIpv4SrcCounter
}

type unMarshalPatternFlowIpv4SrcCounter interface {
	// FromProto unmarshals PatternFlowIpv4SrcCounter from protobuf object *otg.PatternFlowIpv4SrcCounter
	FromProto(msg *otg.PatternFlowIpv4SrcCounter) (PatternFlowIpv4SrcCounter, error)
	// FromPbText unmarshals PatternFlowIpv4SrcCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4SrcCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4SrcCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4SrcCounter) Marshal() marshalPatternFlowIpv4SrcCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4SrcCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4SrcCounter) Unmarshal() unMarshalPatternFlowIpv4SrcCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4SrcCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4SrcCounter) ToProto() (*otg.PatternFlowIpv4SrcCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4SrcCounter) FromProto(msg *otg.PatternFlowIpv4SrcCounter) (PatternFlowIpv4SrcCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4SrcCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4SrcCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4SrcCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4SrcCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4SrcCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4SrcCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4SrcCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4SrcCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4SrcCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4SrcCounter) Clone() (PatternFlowIpv4SrcCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4SrcCounter()
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

// PatternFlowIpv4SrcCounter is ipv4 counter pattern
type PatternFlowIpv4SrcCounter interface {
	Validation
	// msg marshals PatternFlowIpv4SrcCounter to protobuf object *otg.PatternFlowIpv4SrcCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4SrcCounter
	// setMsg unmarshals PatternFlowIpv4SrcCounter from protobuf object *otg.PatternFlowIpv4SrcCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4SrcCounter) PatternFlowIpv4SrcCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4SrcCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4SrcCounter
	// validate validates PatternFlowIpv4SrcCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4SrcCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv4SrcCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv4SrcCounter
	SetStart(value string) PatternFlowIpv4SrcCounter
	// HasStart checks if Start has been set in PatternFlowIpv4SrcCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv4SrcCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv4SrcCounter
	SetStep(value string) PatternFlowIpv4SrcCounter
	// HasStep checks if Step has been set in PatternFlowIpv4SrcCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4SrcCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4SrcCounter
	SetCount(value uint32) PatternFlowIpv4SrcCounter
	// HasCount checks if Count has been set in PatternFlowIpv4SrcCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv4SrcCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv4SrcCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv4SrcCounter object
func (obj *patternFlowIpv4SrcCounter) SetStart(value string) PatternFlowIpv4SrcCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv4SrcCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv4SrcCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv4SrcCounter object
func (obj *patternFlowIpv4SrcCounter) SetStep(value string) PatternFlowIpv4SrcCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4SrcCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4SrcCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4SrcCounter object
func (obj *patternFlowIpv4SrcCounter) SetCount(value uint32) PatternFlowIpv4SrcCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4SrcCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4SrcCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4SrcCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv4SrcCounter) setDefault() {
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
