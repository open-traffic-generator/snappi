package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorSystemIdCounter *****
type patternFlowLacpActorSystemIdCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorSystemIdCounter
	marshaller   marshalPatternFlowLacpActorSystemIdCounter
	unMarshaller unMarshalPatternFlowLacpActorSystemIdCounter
}

func NewPatternFlowLacpActorSystemIdCounter() PatternFlowLacpActorSystemIdCounter {
	obj := patternFlowLacpActorSystemIdCounter{obj: &otg.PatternFlowLacpActorSystemIdCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorSystemIdCounter) msg() *otg.PatternFlowLacpActorSystemIdCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorSystemIdCounter) setMsg(msg *otg.PatternFlowLacpActorSystemIdCounter) PatternFlowLacpActorSystemIdCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorSystemIdCounter struct {
	obj *patternFlowLacpActorSystemIdCounter
}

type marshalPatternFlowLacpActorSystemIdCounter interface {
	// ToProto marshals PatternFlowLacpActorSystemIdCounter to protobuf object *otg.PatternFlowLacpActorSystemIdCounter
	ToProto() (*otg.PatternFlowLacpActorSystemIdCounter, error)
	// ToPbText marshals PatternFlowLacpActorSystemIdCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorSystemIdCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorSystemIdCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorSystemIdCounter struct {
	obj *patternFlowLacpActorSystemIdCounter
}

type unMarshalPatternFlowLacpActorSystemIdCounter interface {
	// FromProto unmarshals PatternFlowLacpActorSystemIdCounter from protobuf object *otg.PatternFlowLacpActorSystemIdCounter
	FromProto(msg *otg.PatternFlowLacpActorSystemIdCounter) (PatternFlowLacpActorSystemIdCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorSystemIdCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorSystemIdCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorSystemIdCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorSystemIdCounter) Marshal() marshalPatternFlowLacpActorSystemIdCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorSystemIdCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorSystemIdCounter) Unmarshal() unMarshalPatternFlowLacpActorSystemIdCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorSystemIdCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorSystemIdCounter) ToProto() (*otg.PatternFlowLacpActorSystemIdCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorSystemIdCounter) FromProto(msg *otg.PatternFlowLacpActorSystemIdCounter) (PatternFlowLacpActorSystemIdCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorSystemIdCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorSystemIdCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorSystemIdCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorSystemIdCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorSystemIdCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorSystemIdCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorSystemIdCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorSystemIdCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorSystemIdCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorSystemIdCounter) Clone() (PatternFlowLacpActorSystemIdCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorSystemIdCounter()
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

// PatternFlowLacpActorSystemIdCounter is mac counter pattern
type PatternFlowLacpActorSystemIdCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorSystemIdCounter to protobuf object *otg.PatternFlowLacpActorSystemIdCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorSystemIdCounter
	// setMsg unmarshals PatternFlowLacpActorSystemIdCounter from protobuf object *otg.PatternFlowLacpActorSystemIdCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorSystemIdCounter) PatternFlowLacpActorSystemIdCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorSystemIdCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorSystemIdCounter
	// validate validates PatternFlowLacpActorSystemIdCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorSystemIdCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowLacpActorSystemIdCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowLacpActorSystemIdCounter
	SetStart(value string) PatternFlowLacpActorSystemIdCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorSystemIdCounter
	HasStart() bool
	// Step returns string, set in PatternFlowLacpActorSystemIdCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowLacpActorSystemIdCounter
	SetStep(value string) PatternFlowLacpActorSystemIdCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorSystemIdCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorSystemIdCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorSystemIdCounter
	SetCount(value uint32) PatternFlowLacpActorSystemIdCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorSystemIdCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowLacpActorSystemIdCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowLacpActorSystemIdCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowLacpActorSystemIdCounter object
func (obj *patternFlowLacpActorSystemIdCounter) SetStart(value string) PatternFlowLacpActorSystemIdCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowLacpActorSystemIdCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowLacpActorSystemIdCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowLacpActorSystemIdCounter object
func (obj *patternFlowLacpActorSystemIdCounter) SetStep(value string) PatternFlowLacpActorSystemIdCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorSystemIdCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorSystemIdCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorSystemIdCounter object
func (obj *patternFlowLacpActorSystemIdCounter) SetCount(value uint32) PatternFlowLacpActorSystemIdCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorSystemIdCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpActorSystemIdCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpActorSystemIdCounter.Step"))
		}

	}

}

func (obj *patternFlowLacpActorSystemIdCounter) setDefault() {
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
