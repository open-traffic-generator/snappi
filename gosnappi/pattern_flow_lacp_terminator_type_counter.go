package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpTerminatorTypeCounter *****
type patternFlowLacpTerminatorTypeCounter struct {
	validation
	obj          *otg.PatternFlowLacpTerminatorTypeCounter
	marshaller   marshalPatternFlowLacpTerminatorTypeCounter
	unMarshaller unMarshalPatternFlowLacpTerminatorTypeCounter
}

func NewPatternFlowLacpTerminatorTypeCounter() PatternFlowLacpTerminatorTypeCounter {
	obj := patternFlowLacpTerminatorTypeCounter{obj: &otg.PatternFlowLacpTerminatorTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpTerminatorTypeCounter) msg() *otg.PatternFlowLacpTerminatorTypeCounter {
	return obj.obj
}

func (obj *patternFlowLacpTerminatorTypeCounter) setMsg(msg *otg.PatternFlowLacpTerminatorTypeCounter) PatternFlowLacpTerminatorTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpTerminatorTypeCounter struct {
	obj *patternFlowLacpTerminatorTypeCounter
}

type marshalPatternFlowLacpTerminatorTypeCounter interface {
	// ToProto marshals PatternFlowLacpTerminatorTypeCounter to protobuf object *otg.PatternFlowLacpTerminatorTypeCounter
	ToProto() (*otg.PatternFlowLacpTerminatorTypeCounter, error)
	// ToPbText marshals PatternFlowLacpTerminatorTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpTerminatorTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpTerminatorTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpTerminatorTypeCounter struct {
	obj *patternFlowLacpTerminatorTypeCounter
}

type unMarshalPatternFlowLacpTerminatorTypeCounter interface {
	// FromProto unmarshals PatternFlowLacpTerminatorTypeCounter from protobuf object *otg.PatternFlowLacpTerminatorTypeCounter
	FromProto(msg *otg.PatternFlowLacpTerminatorTypeCounter) (PatternFlowLacpTerminatorTypeCounter, error)
	// FromPbText unmarshals PatternFlowLacpTerminatorTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpTerminatorTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpTerminatorTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpTerminatorTypeCounter) Marshal() marshalPatternFlowLacpTerminatorTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpTerminatorTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpTerminatorTypeCounter) Unmarshal() unMarshalPatternFlowLacpTerminatorTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpTerminatorTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpTerminatorTypeCounter) ToProto() (*otg.PatternFlowLacpTerminatorTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpTerminatorTypeCounter) FromProto(msg *otg.PatternFlowLacpTerminatorTypeCounter) (PatternFlowLacpTerminatorTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpTerminatorTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpTerminatorTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpTerminatorTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpTerminatorTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpTerminatorTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpTerminatorTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpTerminatorTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpTerminatorTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpTerminatorTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpTerminatorTypeCounter) Clone() (PatternFlowLacpTerminatorTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpTerminatorTypeCounter()
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

// PatternFlowLacpTerminatorTypeCounter is integer counter pattern
type PatternFlowLacpTerminatorTypeCounter interface {
	Validation
	// msg marshals PatternFlowLacpTerminatorTypeCounter to protobuf object *otg.PatternFlowLacpTerminatorTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpTerminatorTypeCounter
	// setMsg unmarshals PatternFlowLacpTerminatorTypeCounter from protobuf object *otg.PatternFlowLacpTerminatorTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpTerminatorTypeCounter) PatternFlowLacpTerminatorTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpTerminatorTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpTerminatorTypeCounter
	// validate validates PatternFlowLacpTerminatorTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpTerminatorTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpTerminatorTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpTerminatorTypeCounter
	SetStart(value uint32) PatternFlowLacpTerminatorTypeCounter
	// HasStart checks if Start has been set in PatternFlowLacpTerminatorTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpTerminatorTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpTerminatorTypeCounter
	SetStep(value uint32) PatternFlowLacpTerminatorTypeCounter
	// HasStep checks if Step has been set in PatternFlowLacpTerminatorTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpTerminatorTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpTerminatorTypeCounter
	SetCount(value uint32) PatternFlowLacpTerminatorTypeCounter
	// HasCount checks if Count has been set in PatternFlowLacpTerminatorTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpTerminatorTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpTerminatorTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpTerminatorTypeCounter object
func (obj *patternFlowLacpTerminatorTypeCounter) SetStart(value uint32) PatternFlowLacpTerminatorTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpTerminatorTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpTerminatorTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpTerminatorTypeCounter object
func (obj *patternFlowLacpTerminatorTypeCounter) SetStep(value uint32) PatternFlowLacpTerminatorTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpTerminatorTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpTerminatorTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpTerminatorTypeCounter object
func (obj *patternFlowLacpTerminatorTypeCounter) SetCount(value uint32) PatternFlowLacpTerminatorTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpTerminatorTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpTerminatorTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpTerminatorTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpTerminatorTypeCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpTerminatorTypeCounter) setDefault() {
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
