package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorPortPriorityCounter *****
type patternFlowLacpActorPortPriorityCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorPortPriorityCounter
	marshaller   marshalPatternFlowLacpActorPortPriorityCounter
	unMarshaller unMarshalPatternFlowLacpActorPortPriorityCounter
}

func NewPatternFlowLacpActorPortPriorityCounter() PatternFlowLacpActorPortPriorityCounter {
	obj := patternFlowLacpActorPortPriorityCounter{obj: &otg.PatternFlowLacpActorPortPriorityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorPortPriorityCounter) msg() *otg.PatternFlowLacpActorPortPriorityCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorPortPriorityCounter) setMsg(msg *otg.PatternFlowLacpActorPortPriorityCounter) PatternFlowLacpActorPortPriorityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorPortPriorityCounter struct {
	obj *patternFlowLacpActorPortPriorityCounter
}

type marshalPatternFlowLacpActorPortPriorityCounter interface {
	// ToProto marshals PatternFlowLacpActorPortPriorityCounter to protobuf object *otg.PatternFlowLacpActorPortPriorityCounter
	ToProto() (*otg.PatternFlowLacpActorPortPriorityCounter, error)
	// ToPbText marshals PatternFlowLacpActorPortPriorityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorPortPriorityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorPortPriorityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorPortPriorityCounter struct {
	obj *patternFlowLacpActorPortPriorityCounter
}

type unMarshalPatternFlowLacpActorPortPriorityCounter interface {
	// FromProto unmarshals PatternFlowLacpActorPortPriorityCounter from protobuf object *otg.PatternFlowLacpActorPortPriorityCounter
	FromProto(msg *otg.PatternFlowLacpActorPortPriorityCounter) (PatternFlowLacpActorPortPriorityCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorPortPriorityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorPortPriorityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorPortPriorityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorPortPriorityCounter) Marshal() marshalPatternFlowLacpActorPortPriorityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorPortPriorityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorPortPriorityCounter) Unmarshal() unMarshalPatternFlowLacpActorPortPriorityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorPortPriorityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorPortPriorityCounter) ToProto() (*otg.PatternFlowLacpActorPortPriorityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorPortPriorityCounter) FromProto(msg *otg.PatternFlowLacpActorPortPriorityCounter) (PatternFlowLacpActorPortPriorityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorPortPriorityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorPortPriorityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorPortPriorityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorPortPriorityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorPortPriorityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorPortPriorityCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorPortPriorityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorPortPriorityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorPortPriorityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorPortPriorityCounter) Clone() (PatternFlowLacpActorPortPriorityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorPortPriorityCounter()
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

// PatternFlowLacpActorPortPriorityCounter is integer counter pattern
type PatternFlowLacpActorPortPriorityCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorPortPriorityCounter to protobuf object *otg.PatternFlowLacpActorPortPriorityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorPortPriorityCounter
	// setMsg unmarshals PatternFlowLacpActorPortPriorityCounter from protobuf object *otg.PatternFlowLacpActorPortPriorityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorPortPriorityCounter) PatternFlowLacpActorPortPriorityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorPortPriorityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorPortPriorityCounter
	// validate validates PatternFlowLacpActorPortPriorityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorPortPriorityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorPortPriorityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorPortPriorityCounter
	SetStart(value uint32) PatternFlowLacpActorPortPriorityCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorPortPriorityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorPortPriorityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorPortPriorityCounter
	SetStep(value uint32) PatternFlowLacpActorPortPriorityCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorPortPriorityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorPortPriorityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorPortPriorityCounter
	SetCount(value uint32) PatternFlowLacpActorPortPriorityCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorPortPriorityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorPortPriorityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorPortPriorityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorPortPriorityCounter object
func (obj *patternFlowLacpActorPortPriorityCounter) SetStart(value uint32) PatternFlowLacpActorPortPriorityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorPortPriorityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorPortPriorityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorPortPriorityCounter object
func (obj *patternFlowLacpActorPortPriorityCounter) SetStep(value uint32) PatternFlowLacpActorPortPriorityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorPortPriorityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorPortPriorityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorPortPriorityCounter object
func (obj *patternFlowLacpActorPortPriorityCounter) SetCount(value uint32) PatternFlowLacpActorPortPriorityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorPortPriorityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorPortPriorityCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorPortPriorityCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorPortPriorityCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorPortPriorityCounter) setDefault() {
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
