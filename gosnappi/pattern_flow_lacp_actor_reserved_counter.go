package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorReservedCounter *****
type patternFlowLacpActorReservedCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorReservedCounter
	marshaller   marshalPatternFlowLacpActorReservedCounter
	unMarshaller unMarshalPatternFlowLacpActorReservedCounter
}

func NewPatternFlowLacpActorReservedCounter() PatternFlowLacpActorReservedCounter {
	obj := patternFlowLacpActorReservedCounter{obj: &otg.PatternFlowLacpActorReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorReservedCounter) msg() *otg.PatternFlowLacpActorReservedCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorReservedCounter) setMsg(msg *otg.PatternFlowLacpActorReservedCounter) PatternFlowLacpActorReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorReservedCounter struct {
	obj *patternFlowLacpActorReservedCounter
}

type marshalPatternFlowLacpActorReservedCounter interface {
	// ToProto marshals PatternFlowLacpActorReservedCounter to protobuf object *otg.PatternFlowLacpActorReservedCounter
	ToProto() (*otg.PatternFlowLacpActorReservedCounter, error)
	// ToPbText marshals PatternFlowLacpActorReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorReservedCounter struct {
	obj *patternFlowLacpActorReservedCounter
}

type unMarshalPatternFlowLacpActorReservedCounter interface {
	// FromProto unmarshals PatternFlowLacpActorReservedCounter from protobuf object *otg.PatternFlowLacpActorReservedCounter
	FromProto(msg *otg.PatternFlowLacpActorReservedCounter) (PatternFlowLacpActorReservedCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorReservedCounter) Marshal() marshalPatternFlowLacpActorReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorReservedCounter) Unmarshal() unMarshalPatternFlowLacpActorReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorReservedCounter) ToProto() (*otg.PatternFlowLacpActorReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorReservedCounter) FromProto(msg *otg.PatternFlowLacpActorReservedCounter) (PatternFlowLacpActorReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorReservedCounter) Clone() (PatternFlowLacpActorReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorReservedCounter()
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

// PatternFlowLacpActorReservedCounter is integer counter pattern
type PatternFlowLacpActorReservedCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorReservedCounter to protobuf object *otg.PatternFlowLacpActorReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorReservedCounter
	// setMsg unmarshals PatternFlowLacpActorReservedCounter from protobuf object *otg.PatternFlowLacpActorReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorReservedCounter) PatternFlowLacpActorReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorReservedCounter
	// validate validates PatternFlowLacpActorReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorReservedCounter
	SetStart(value uint32) PatternFlowLacpActorReservedCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorReservedCounter
	SetStep(value uint32) PatternFlowLacpActorReservedCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorReservedCounter
	SetCount(value uint32) PatternFlowLacpActorReservedCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorReservedCounter object
func (obj *patternFlowLacpActorReservedCounter) SetStart(value uint32) PatternFlowLacpActorReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorReservedCounter object
func (obj *patternFlowLacpActorReservedCounter) SetStep(value uint32) PatternFlowLacpActorReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorReservedCounter object
func (obj *patternFlowLacpActorReservedCounter) SetCount(value uint32) PatternFlowLacpActorReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorReservedCounter.Start <= 16777215 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorReservedCounter.Step <= 16777215 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorReservedCounter.Count <= 16777215 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorReservedCounter) setDefault() {
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
