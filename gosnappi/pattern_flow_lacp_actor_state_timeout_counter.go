package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateTimeoutCounter *****
type patternFlowLacpActorStateTimeoutCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorStateTimeoutCounter
	marshaller   marshalPatternFlowLacpActorStateTimeoutCounter
	unMarshaller unMarshalPatternFlowLacpActorStateTimeoutCounter
}

func NewPatternFlowLacpActorStateTimeoutCounter() PatternFlowLacpActorStateTimeoutCounter {
	obj := patternFlowLacpActorStateTimeoutCounter{obj: &otg.PatternFlowLacpActorStateTimeoutCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateTimeoutCounter) msg() *otg.PatternFlowLacpActorStateTimeoutCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorStateTimeoutCounter) setMsg(msg *otg.PatternFlowLacpActorStateTimeoutCounter) PatternFlowLacpActorStateTimeoutCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateTimeoutCounter struct {
	obj *patternFlowLacpActorStateTimeoutCounter
}

type marshalPatternFlowLacpActorStateTimeoutCounter interface {
	// ToProto marshals PatternFlowLacpActorStateTimeoutCounter to protobuf object *otg.PatternFlowLacpActorStateTimeoutCounter
	ToProto() (*otg.PatternFlowLacpActorStateTimeoutCounter, error)
	// ToPbText marshals PatternFlowLacpActorStateTimeoutCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateTimeoutCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateTimeoutCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateTimeoutCounter struct {
	obj *patternFlowLacpActorStateTimeoutCounter
}

type unMarshalPatternFlowLacpActorStateTimeoutCounter interface {
	// FromProto unmarshals PatternFlowLacpActorStateTimeoutCounter from protobuf object *otg.PatternFlowLacpActorStateTimeoutCounter
	FromProto(msg *otg.PatternFlowLacpActorStateTimeoutCounter) (PatternFlowLacpActorStateTimeoutCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorStateTimeoutCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateTimeoutCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateTimeoutCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateTimeoutCounter) Marshal() marshalPatternFlowLacpActorStateTimeoutCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateTimeoutCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateTimeoutCounter) Unmarshal() unMarshalPatternFlowLacpActorStateTimeoutCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateTimeoutCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateTimeoutCounter) ToProto() (*otg.PatternFlowLacpActorStateTimeoutCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateTimeoutCounter) FromProto(msg *otg.PatternFlowLacpActorStateTimeoutCounter) (PatternFlowLacpActorStateTimeoutCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateTimeoutCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateTimeoutCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateTimeoutCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateTimeoutCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateTimeoutCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateTimeoutCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateTimeoutCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateTimeoutCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateTimeoutCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateTimeoutCounter) Clone() (PatternFlowLacpActorStateTimeoutCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateTimeoutCounter()
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

// PatternFlowLacpActorStateTimeoutCounter is integer counter pattern
type PatternFlowLacpActorStateTimeoutCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorStateTimeoutCounter to protobuf object *otg.PatternFlowLacpActorStateTimeoutCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateTimeoutCounter
	// setMsg unmarshals PatternFlowLacpActorStateTimeoutCounter from protobuf object *otg.PatternFlowLacpActorStateTimeoutCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateTimeoutCounter) PatternFlowLacpActorStateTimeoutCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateTimeoutCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateTimeoutCounter
	// validate validates PatternFlowLacpActorStateTimeoutCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateTimeoutCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorStateTimeoutCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorStateTimeoutCounter
	SetStart(value uint32) PatternFlowLacpActorStateTimeoutCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorStateTimeoutCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorStateTimeoutCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorStateTimeoutCounter
	SetStep(value uint32) PatternFlowLacpActorStateTimeoutCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorStateTimeoutCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorStateTimeoutCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorStateTimeoutCounter
	SetCount(value uint32) PatternFlowLacpActorStateTimeoutCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorStateTimeoutCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateTimeoutCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateTimeoutCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorStateTimeoutCounter object
func (obj *patternFlowLacpActorStateTimeoutCounter) SetStart(value uint32) PatternFlowLacpActorStateTimeoutCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateTimeoutCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateTimeoutCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorStateTimeoutCounter object
func (obj *patternFlowLacpActorStateTimeoutCounter) SetStep(value uint32) PatternFlowLacpActorStateTimeoutCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateTimeoutCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateTimeoutCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorStateTimeoutCounter object
func (obj *patternFlowLacpActorStateTimeoutCounter) SetCount(value uint32) PatternFlowLacpActorStateTimeoutCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorStateTimeoutCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateTimeoutCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateTimeoutCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateTimeoutCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorStateTimeoutCounter) setDefault() {
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
