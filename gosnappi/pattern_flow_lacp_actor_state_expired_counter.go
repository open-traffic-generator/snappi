package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateExpiredCounter *****
type patternFlowLacpActorStateExpiredCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorStateExpiredCounter
	marshaller   marshalPatternFlowLacpActorStateExpiredCounter
	unMarshaller unMarshalPatternFlowLacpActorStateExpiredCounter
}

func NewPatternFlowLacpActorStateExpiredCounter() PatternFlowLacpActorStateExpiredCounter {
	obj := patternFlowLacpActorStateExpiredCounter{obj: &otg.PatternFlowLacpActorStateExpiredCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateExpiredCounter) msg() *otg.PatternFlowLacpActorStateExpiredCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorStateExpiredCounter) setMsg(msg *otg.PatternFlowLacpActorStateExpiredCounter) PatternFlowLacpActorStateExpiredCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateExpiredCounter struct {
	obj *patternFlowLacpActorStateExpiredCounter
}

type marshalPatternFlowLacpActorStateExpiredCounter interface {
	// ToProto marshals PatternFlowLacpActorStateExpiredCounter to protobuf object *otg.PatternFlowLacpActorStateExpiredCounter
	ToProto() (*otg.PatternFlowLacpActorStateExpiredCounter, error)
	// ToPbText marshals PatternFlowLacpActorStateExpiredCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateExpiredCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateExpiredCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateExpiredCounter struct {
	obj *patternFlowLacpActorStateExpiredCounter
}

type unMarshalPatternFlowLacpActorStateExpiredCounter interface {
	// FromProto unmarshals PatternFlowLacpActorStateExpiredCounter from protobuf object *otg.PatternFlowLacpActorStateExpiredCounter
	FromProto(msg *otg.PatternFlowLacpActorStateExpiredCounter) (PatternFlowLacpActorStateExpiredCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorStateExpiredCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateExpiredCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateExpiredCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateExpiredCounter) Marshal() marshalPatternFlowLacpActorStateExpiredCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateExpiredCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateExpiredCounter) Unmarshal() unMarshalPatternFlowLacpActorStateExpiredCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateExpiredCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateExpiredCounter) ToProto() (*otg.PatternFlowLacpActorStateExpiredCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateExpiredCounter) FromProto(msg *otg.PatternFlowLacpActorStateExpiredCounter) (PatternFlowLacpActorStateExpiredCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateExpiredCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateExpiredCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateExpiredCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateExpiredCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateExpiredCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateExpiredCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateExpiredCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateExpiredCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateExpiredCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateExpiredCounter) Clone() (PatternFlowLacpActorStateExpiredCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateExpiredCounter()
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

// PatternFlowLacpActorStateExpiredCounter is integer counter pattern
type PatternFlowLacpActorStateExpiredCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorStateExpiredCounter to protobuf object *otg.PatternFlowLacpActorStateExpiredCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateExpiredCounter
	// setMsg unmarshals PatternFlowLacpActorStateExpiredCounter from protobuf object *otg.PatternFlowLacpActorStateExpiredCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateExpiredCounter) PatternFlowLacpActorStateExpiredCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateExpiredCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateExpiredCounter
	// validate validates PatternFlowLacpActorStateExpiredCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateExpiredCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorStateExpiredCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorStateExpiredCounter
	SetStart(value uint32) PatternFlowLacpActorStateExpiredCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorStateExpiredCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorStateExpiredCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorStateExpiredCounter
	SetStep(value uint32) PatternFlowLacpActorStateExpiredCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorStateExpiredCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorStateExpiredCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorStateExpiredCounter
	SetCount(value uint32) PatternFlowLacpActorStateExpiredCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorStateExpiredCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateExpiredCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateExpiredCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorStateExpiredCounter object
func (obj *patternFlowLacpActorStateExpiredCounter) SetStart(value uint32) PatternFlowLacpActorStateExpiredCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateExpiredCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateExpiredCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorStateExpiredCounter object
func (obj *patternFlowLacpActorStateExpiredCounter) SetStep(value uint32) PatternFlowLacpActorStateExpiredCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateExpiredCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateExpiredCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorStateExpiredCounter object
func (obj *patternFlowLacpActorStateExpiredCounter) SetCount(value uint32) PatternFlowLacpActorStateExpiredCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorStateExpiredCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateExpiredCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateExpiredCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateExpiredCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorStateExpiredCounter) setDefault() {
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
