package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorPortNumberCounter *****
type patternFlowLacpActorPortNumberCounter struct {
	validation
	obj          *otg.PatternFlowLacpActorPortNumberCounter
	marshaller   marshalPatternFlowLacpActorPortNumberCounter
	unMarshaller unMarshalPatternFlowLacpActorPortNumberCounter
}

func NewPatternFlowLacpActorPortNumberCounter() PatternFlowLacpActorPortNumberCounter {
	obj := patternFlowLacpActorPortNumberCounter{obj: &otg.PatternFlowLacpActorPortNumberCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorPortNumberCounter) msg() *otg.PatternFlowLacpActorPortNumberCounter {
	return obj.obj
}

func (obj *patternFlowLacpActorPortNumberCounter) setMsg(msg *otg.PatternFlowLacpActorPortNumberCounter) PatternFlowLacpActorPortNumberCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorPortNumberCounter struct {
	obj *patternFlowLacpActorPortNumberCounter
}

type marshalPatternFlowLacpActorPortNumberCounter interface {
	// ToProto marshals PatternFlowLacpActorPortNumberCounter to protobuf object *otg.PatternFlowLacpActorPortNumberCounter
	ToProto() (*otg.PatternFlowLacpActorPortNumberCounter, error)
	// ToPbText marshals PatternFlowLacpActorPortNumberCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorPortNumberCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorPortNumberCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorPortNumberCounter struct {
	obj *patternFlowLacpActorPortNumberCounter
}

type unMarshalPatternFlowLacpActorPortNumberCounter interface {
	// FromProto unmarshals PatternFlowLacpActorPortNumberCounter from protobuf object *otg.PatternFlowLacpActorPortNumberCounter
	FromProto(msg *otg.PatternFlowLacpActorPortNumberCounter) (PatternFlowLacpActorPortNumberCounter, error)
	// FromPbText unmarshals PatternFlowLacpActorPortNumberCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorPortNumberCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorPortNumberCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorPortNumberCounter) Marshal() marshalPatternFlowLacpActorPortNumberCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorPortNumberCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorPortNumberCounter) Unmarshal() unMarshalPatternFlowLacpActorPortNumberCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorPortNumberCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorPortNumberCounter) ToProto() (*otg.PatternFlowLacpActorPortNumberCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorPortNumberCounter) FromProto(msg *otg.PatternFlowLacpActorPortNumberCounter) (PatternFlowLacpActorPortNumberCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorPortNumberCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorPortNumberCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorPortNumberCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorPortNumberCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorPortNumberCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorPortNumberCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpActorPortNumberCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorPortNumberCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorPortNumberCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorPortNumberCounter) Clone() (PatternFlowLacpActorPortNumberCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorPortNumberCounter()
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

// PatternFlowLacpActorPortNumberCounter is integer counter pattern
type PatternFlowLacpActorPortNumberCounter interface {
	Validation
	// msg marshals PatternFlowLacpActorPortNumberCounter to protobuf object *otg.PatternFlowLacpActorPortNumberCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorPortNumberCounter
	// setMsg unmarshals PatternFlowLacpActorPortNumberCounter from protobuf object *otg.PatternFlowLacpActorPortNumberCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorPortNumberCounter) PatternFlowLacpActorPortNumberCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorPortNumberCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorPortNumberCounter
	// validate validates PatternFlowLacpActorPortNumberCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorPortNumberCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpActorPortNumberCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpActorPortNumberCounter
	SetStart(value uint32) PatternFlowLacpActorPortNumberCounter
	// HasStart checks if Start has been set in PatternFlowLacpActorPortNumberCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpActorPortNumberCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpActorPortNumberCounter
	SetStep(value uint32) PatternFlowLacpActorPortNumberCounter
	// HasStep checks if Step has been set in PatternFlowLacpActorPortNumberCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpActorPortNumberCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpActorPortNumberCounter
	SetCount(value uint32) PatternFlowLacpActorPortNumberCounter
	// HasCount checks if Count has been set in PatternFlowLacpActorPortNumberCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorPortNumberCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpActorPortNumberCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpActorPortNumberCounter object
func (obj *patternFlowLacpActorPortNumberCounter) SetStart(value uint32) PatternFlowLacpActorPortNumberCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorPortNumberCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpActorPortNumberCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpActorPortNumberCounter object
func (obj *patternFlowLacpActorPortNumberCounter) SetStep(value uint32) PatternFlowLacpActorPortNumberCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorPortNumberCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpActorPortNumberCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpActorPortNumberCounter object
func (obj *patternFlowLacpActorPortNumberCounter) SetCount(value uint32) PatternFlowLacpActorPortNumberCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpActorPortNumberCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorPortNumberCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorPortNumberCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorPortNumberCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpActorPortNumberCounter) setDefault() {
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
