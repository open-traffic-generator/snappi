package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorSystemPriorityCounter *****
type patternFlowLacpActorSystemPriorityCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorSystemPriorityCounter
	marshaller   marshalPatternFlowLacpActorSystemPriorityCounter
	unMarshaller unMarshalPatternFlowLacpActorSystemPriorityCounter
}

func NewPatternFlowLacpActorSystemPriorityCounter() PatternFlowLacpActorSystemPriorityCounter {
	obj := patternFlowLacpActorSystemPriorityCounter{obj: &otg.PatternFlowLacpActorSystemPriorityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorSystemPriorityCounter) msg() *otg.PatternFlowLacpActorSystemPriorityCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorSystemPriorityCounter) setMsg(msg *otg.PatternFlowLacpActorSystemPriorityCounter) PatternFlowLacpActorSystemPriorityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorSystemPriorityCounter struct {
	obj *patternFlowLacpActorSystemPriorityCounter
}

type marshalPatternFlowLacpActorSystemPriorityCounter interface {
	// ToProto marshals PatternFlowLacpActorSystemPriorityCounter to protobuf object *otg.PatternFlowLacpActorSystemPriorityCounter
	ToProto() (*otg.PatternFlowLacpActorSystemPriorityCounter, error)
	// ToPbText marshals PatternFlowLacpActorSystemPriorityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorSystemPriorityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorSystemPriorityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorSystemPriorityCounter struct {
	obj *patternFlowLacpActorSystemPriorityCounter
}

type unMarshalPatternFlowLacpActorSystemPriorityCounter interface {
	// FromProto unmarshals PatternFlowLacpActorSystemPriorityCounter from protobuf object *otg.PatternFlowLacpActorSystemPriorityCounter
	FromProto(msg *otg.PatternFlowLacpActorSystemPriorityCounter) (PatternFlowLacpActorSystemPriorityCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorSystemPriorityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorSystemPriorityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorSystemPriorityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorSystemPriorityCounter) Marshal() marshalPatternFlowLacpActorSystemPriorityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorSystemPriorityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorSystemPriorityCounter) Unmarshal() unMarshalPatternFlowLacpActorSystemPriorityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorSystemPriorityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorSystemPriorityCounter) ToProto() (*otg.PatternFlowLacpActorSystemPriorityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorSystemPriorityCounter) FromProto(msg *otg.PatternFlowLacpActorSystemPriorityCounter) (PatternFlowLacpActorSystemPriorityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorSystemPriorityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorSystemPriorityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorSystemPriorityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorSystemPriorityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorSystemPriorityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorSystemPriorityCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorSystemPriorityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorSystemPriorityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorSystemPriorityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorSystemPriorityCounter) Clone() (PatternFlowLacpActorSystemPriorityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorSystemPriorityCounter()
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

// PatternFlowLacpActorSystemPriorityCounter is integer counter pattern
type PatternFlowLacpActorSystemPriorityCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorSystemPriorityCounter to protobuf object *otg.PatternFlowLacpActorSystemPriorityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorSystemPriorityCounter
	// setMsg unmarshals PatternFlowLacpActorSystemPriorityCounter from protobuf object *otg.PatternFlowLacpActorSystemPriorityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorSystemPriorityCounter) PatternFlowLacpActorSystemPriorityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorSystemPriorityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorSystemPriorityCounter
	// validate validates PatternFlowLacpActorSystemPriorityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorSystemPriorityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorSystemPriorityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorSystemPriorityCounter
	SetStart(value uint32) PatternFlowLacpActorSystemPriorityCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorSystemPriorityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorSystemPriorityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorSystemPriorityCounter
	SetStep(value uint32) PatternFlowLacpActorSystemPriorityCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorSystemPriorityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorSystemPriorityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorSystemPriorityCounter
	SetCount(value uint32) PatternFlowLacpActorSystemPriorityCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorSystemPriorityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorSystemPriorityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorSystemPriorityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorSystemPriorityCounter object
func (obj *patternFlowLacpActorSystemPriorityCounter) SetStart(value uint32) PatternFlowLacpActorSystemPriorityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorSystemPriorityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorSystemPriorityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorSystemPriorityCounter object
func (obj *patternFlowLacpActorSystemPriorityCounter) SetStep(value uint32) PatternFlowLacpActorSystemPriorityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorSystemPriorityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorSystemPriorityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorSystemPriorityCounter object
func (obj *patternFlowLacpActorSystemPriorityCounter) SetCount(value uint32) PatternFlowLacpActorSystemPriorityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorSystemPriorityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorSystemPriorityCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorSystemPriorityCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorSystemPriorityCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorSystemPriorityCounter) setDefault() {
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
