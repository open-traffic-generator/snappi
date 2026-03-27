package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorSystemIdCounter *****
type patternFlowLacpduActorSystemIdCounter struct {
	validation
	obj          *otg.PatternFlowLacpduActorSystemIdCounter
	marshaller   marshalPatternFlowLacpduActorSystemIdCounter
	unMarshaller unMarshalPatternFlowLacpduActorSystemIdCounter
}

func NewPatternFlowLacpduActorSystemIdCounter() PatternFlowLacpduActorSystemIdCounter {
	obj := patternFlowLacpduActorSystemIdCounter{obj: &otg.PatternFlowLacpduActorSystemIdCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorSystemIdCounter) msg() *otg.PatternFlowLacpduActorSystemIdCounter {
	return obj.obj
}

func (obj *patternFlowLacpduActorSystemIdCounter) setMsg(msg *otg.PatternFlowLacpduActorSystemIdCounter) PatternFlowLacpduActorSystemIdCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorSystemIdCounter struct {
	obj *patternFlowLacpduActorSystemIdCounter
}

type marshalPatternFlowLacpduActorSystemIdCounter interface {
	// ToProto marshals PatternFlowLacpduActorSystemIdCounter to protobuf object *otg.PatternFlowLacpduActorSystemIdCounter
	ToProto() (*otg.PatternFlowLacpduActorSystemIdCounter, error)
	// ToPbText marshals PatternFlowLacpduActorSystemIdCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorSystemIdCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorSystemIdCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorSystemIdCounter struct {
	obj *patternFlowLacpduActorSystemIdCounter
}

type unMarshalPatternFlowLacpduActorSystemIdCounter interface {
	// FromProto unmarshals PatternFlowLacpduActorSystemIdCounter from protobuf object *otg.PatternFlowLacpduActorSystemIdCounter
	FromProto(msg *otg.PatternFlowLacpduActorSystemIdCounter) (PatternFlowLacpduActorSystemIdCounter, error)
	// FromPbText unmarshals PatternFlowLacpduActorSystemIdCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorSystemIdCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorSystemIdCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorSystemIdCounter) Marshal() marshalPatternFlowLacpduActorSystemIdCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorSystemIdCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorSystemIdCounter) Unmarshal() unMarshalPatternFlowLacpduActorSystemIdCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorSystemIdCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorSystemIdCounter) ToProto() (*otg.PatternFlowLacpduActorSystemIdCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorSystemIdCounter) FromProto(msg *otg.PatternFlowLacpduActorSystemIdCounter) (PatternFlowLacpduActorSystemIdCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorSystemIdCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorSystemIdCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorSystemIdCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorSystemIdCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorSystemIdCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorSystemIdCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorSystemIdCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorSystemIdCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorSystemIdCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorSystemIdCounter) Clone() (PatternFlowLacpduActorSystemIdCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorSystemIdCounter()
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

// PatternFlowLacpduActorSystemIdCounter is mac counter pattern
type PatternFlowLacpduActorSystemIdCounter interface {
	Validation
	// msg marshals PatternFlowLacpduActorSystemIdCounter to protobuf object *otg.PatternFlowLacpduActorSystemIdCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorSystemIdCounter
	// setMsg unmarshals PatternFlowLacpduActorSystemIdCounter from protobuf object *otg.PatternFlowLacpduActorSystemIdCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorSystemIdCounter) PatternFlowLacpduActorSystemIdCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorSystemIdCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorSystemIdCounter
	// validate validates PatternFlowLacpduActorSystemIdCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorSystemIdCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowLacpduActorSystemIdCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowLacpduActorSystemIdCounter
	SetStart(value string) PatternFlowLacpduActorSystemIdCounter
	// HasStart checks if Start has been set in PatternFlowLacpduActorSystemIdCounter
	HasStart() bool
	// Step returns string, set in PatternFlowLacpduActorSystemIdCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowLacpduActorSystemIdCounter
	SetStep(value string) PatternFlowLacpduActorSystemIdCounter
	// HasStep checks if Step has been set in PatternFlowLacpduActorSystemIdCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduActorSystemIdCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduActorSystemIdCounter
	SetCount(value uint32) PatternFlowLacpduActorSystemIdCounter
	// HasCount checks if Count has been set in PatternFlowLacpduActorSystemIdCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowLacpduActorSystemIdCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowLacpduActorSystemIdCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowLacpduActorSystemIdCounter object
func (obj *patternFlowLacpduActorSystemIdCounter) SetStart(value string) PatternFlowLacpduActorSystemIdCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowLacpduActorSystemIdCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowLacpduActorSystemIdCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowLacpduActorSystemIdCounter object
func (obj *patternFlowLacpduActorSystemIdCounter) SetStep(value string) PatternFlowLacpduActorSystemIdCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorSystemIdCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduActorSystemIdCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduActorSystemIdCounter object
func (obj *patternFlowLacpduActorSystemIdCounter) SetCount(value uint32) PatternFlowLacpduActorSystemIdCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduActorSystemIdCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpduActorSystemIdCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpduActorSystemIdCounter.Step"))
		}

	}

}

func (obj *patternFlowLacpduActorSystemIdCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("00:00:00:00:00:00")
	}
	if obj.obj.Step == nil {
		obj.SetStep("00:00:00:00:00:01")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
