package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateDefaultedCounter *****
type patternFlowLacpActorStateDefaultedCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorStateDefaultedCounter
	marshaller   marshalPatternFlowLacpActorStateDefaultedCounter
	unMarshaller unMarshalPatternFlowLacpActorStateDefaultedCounter
}

func NewPatternFlowLacpActorStateDefaultedCounter() PatternFlowLacpActorStateDefaultedCounter {
	obj := patternFlowLacpActorStateDefaultedCounter{obj: &otg.PatternFlowLacpActorStateDefaultedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateDefaultedCounter) msg() *otg.PatternFlowLacpActorStateDefaultedCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorStateDefaultedCounter) setMsg(msg *otg.PatternFlowLacpActorStateDefaultedCounter) PatternFlowLacpActorStateDefaultedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateDefaultedCounter struct {
	obj *patternFlowLacpActorStateDefaultedCounter
}

type marshalPatternFlowLacpActorStateDefaultedCounter interface {
	// ToProto marshals PatternFlowLacpActorStateDefaultedCounter to protobuf object *otg.PatternFlowLacpActorStateDefaultedCounter
	ToProto() (*otg.PatternFlowLacpActorStateDefaultedCounter, error)
	// ToPbText marshals PatternFlowLacpActorStateDefaultedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateDefaultedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateDefaultedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateDefaultedCounter struct {
	obj *patternFlowLacpActorStateDefaultedCounter
}

type unMarshalPatternFlowLacpActorStateDefaultedCounter interface {
	// FromProto unmarshals PatternFlowLacpActorStateDefaultedCounter from protobuf object *otg.PatternFlowLacpActorStateDefaultedCounter
	FromProto(msg *otg.PatternFlowLacpActorStateDefaultedCounter) (PatternFlowLacpActorStateDefaultedCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorStateDefaultedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateDefaultedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateDefaultedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateDefaultedCounter) Marshal() marshalPatternFlowLacpActorStateDefaultedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateDefaultedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateDefaultedCounter) Unmarshal() unMarshalPatternFlowLacpActorStateDefaultedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateDefaultedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateDefaultedCounter) ToProto() (*otg.PatternFlowLacpActorStateDefaultedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateDefaultedCounter) FromProto(msg *otg.PatternFlowLacpActorStateDefaultedCounter) (PatternFlowLacpActorStateDefaultedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateDefaultedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateDefaultedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateDefaultedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateDefaultedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateDefaultedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateDefaultedCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateDefaultedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateDefaultedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateDefaultedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateDefaultedCounter) Clone() (PatternFlowLacpActorStateDefaultedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateDefaultedCounter()
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

// PatternFlowLacpActorStateDefaultedCounter is integer counter pattern
type PatternFlowLacpActorStateDefaultedCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorStateDefaultedCounter to protobuf object *otg.PatternFlowLacpActorStateDefaultedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateDefaultedCounter
	// setMsg unmarshals PatternFlowLacpActorStateDefaultedCounter from protobuf object *otg.PatternFlowLacpActorStateDefaultedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateDefaultedCounter) PatternFlowLacpActorStateDefaultedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateDefaultedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateDefaultedCounter
	// validate validates PatternFlowLacpActorStateDefaultedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateDefaultedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorStateDefaultedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorStateDefaultedCounter
	SetStart(value uint32) PatternFlowLacpActorStateDefaultedCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorStateDefaultedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorStateDefaultedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorStateDefaultedCounter
	SetStep(value uint32) PatternFlowLacpActorStateDefaultedCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorStateDefaultedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorStateDefaultedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorStateDefaultedCounter
	SetCount(value uint32) PatternFlowLacpActorStateDefaultedCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorStateDefaultedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateDefaultedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateDefaultedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorStateDefaultedCounter object
func (obj *patternFlowLacpActorStateDefaultedCounter) SetStart(value uint32) PatternFlowLacpActorStateDefaultedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateDefaultedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateDefaultedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorStateDefaultedCounter object
func (obj *patternFlowLacpActorStateDefaultedCounter) SetStep(value uint32) PatternFlowLacpActorStateDefaultedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateDefaultedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateDefaultedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorStateDefaultedCounter object
func (obj *patternFlowLacpActorStateDefaultedCounter) SetCount(value uint32) PatternFlowLacpActorStateDefaultedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorStateDefaultedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateDefaultedCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateDefaultedCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateDefaultedCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorStateDefaultedCounter) setDefault() {
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
