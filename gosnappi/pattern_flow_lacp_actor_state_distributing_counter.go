package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateDistributingCounter *****
type patternFlowLacpActorStateDistributingCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorStateDistributingCounter
	marshaller   marshalPatternFlowLacpActorStateDistributingCounter
	unMarshaller unMarshalPatternFlowLacpActorStateDistributingCounter
}

func NewPatternFlowLacpActorStateDistributingCounter() PatternFlowLacpActorStateDistributingCounter {
	obj := patternFlowLacpActorStateDistributingCounter{obj: &otg.PatternFlowLacpActorStateDistributingCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateDistributingCounter) msg() *otg.PatternFlowLacpActorStateDistributingCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorStateDistributingCounter) setMsg(msg *otg.PatternFlowLacpActorStateDistributingCounter) PatternFlowLacpActorStateDistributingCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateDistributingCounter struct {
	obj *patternFlowLacpActorStateDistributingCounter
}

type marshalPatternFlowLacpActorStateDistributingCounter interface {
	// ToProto marshals PatternFlowLacpActorStateDistributingCounter to protobuf object *otg.PatternFlowLacpActorStateDistributingCounter
	ToProto() (*otg.PatternFlowLacpActorStateDistributingCounter, error)
	// ToPbText marshals PatternFlowLacpActorStateDistributingCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateDistributingCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateDistributingCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateDistributingCounter struct {
	obj *patternFlowLacpActorStateDistributingCounter
}

type unMarshalPatternFlowLacpActorStateDistributingCounter interface {
	// FromProto unmarshals PatternFlowLacpActorStateDistributingCounter from protobuf object *otg.PatternFlowLacpActorStateDistributingCounter
	FromProto(msg *otg.PatternFlowLacpActorStateDistributingCounter) (PatternFlowLacpActorStateDistributingCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorStateDistributingCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateDistributingCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateDistributingCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateDistributingCounter) Marshal() marshalPatternFlowLacpActorStateDistributingCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateDistributingCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateDistributingCounter) Unmarshal() unMarshalPatternFlowLacpActorStateDistributingCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateDistributingCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateDistributingCounter) ToProto() (*otg.PatternFlowLacpActorStateDistributingCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateDistributingCounter) FromProto(msg *otg.PatternFlowLacpActorStateDistributingCounter) (PatternFlowLacpActorStateDistributingCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateDistributingCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateDistributingCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateDistributingCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateDistributingCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateDistributingCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateDistributingCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateDistributingCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateDistributingCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateDistributingCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateDistributingCounter) Clone() (PatternFlowLacpActorStateDistributingCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateDistributingCounter()
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

// PatternFlowLacpActorStateDistributingCounter is integer counter pattern
type PatternFlowLacpActorStateDistributingCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorStateDistributingCounter to protobuf object *otg.PatternFlowLacpActorStateDistributingCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateDistributingCounter
	// setMsg unmarshals PatternFlowLacpActorStateDistributingCounter from protobuf object *otg.PatternFlowLacpActorStateDistributingCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateDistributingCounter) PatternFlowLacpActorStateDistributingCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateDistributingCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateDistributingCounter
	// validate validates PatternFlowLacpActorStateDistributingCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateDistributingCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorStateDistributingCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorStateDistributingCounter
	SetStart(value uint32) PatternFlowLacpActorStateDistributingCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorStateDistributingCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorStateDistributingCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorStateDistributingCounter
	SetStep(value uint32) PatternFlowLacpActorStateDistributingCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorStateDistributingCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorStateDistributingCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorStateDistributingCounter
	SetCount(value uint32) PatternFlowLacpActorStateDistributingCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorStateDistributingCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateDistributingCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorStateDistributingCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorStateDistributingCounter object
func (obj *patternFlowLacpActorStateDistributingCounter) SetStart(value uint32) PatternFlowLacpActorStateDistributingCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateDistributingCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorStateDistributingCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorStateDistributingCounter object
func (obj *patternFlowLacpActorStateDistributingCounter) SetStep(value uint32) PatternFlowLacpActorStateDistributingCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateDistributingCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorStateDistributingCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorStateDistributingCounter object
func (obj *patternFlowLacpActorStateDistributingCounter) SetCount(value uint32) PatternFlowLacpActorStateDistributingCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorStateDistributingCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateDistributingCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateDistributingCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateDistributingCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorStateDistributingCounter) setDefault() {
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
