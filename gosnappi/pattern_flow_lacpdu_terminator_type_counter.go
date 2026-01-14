package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduTerminatorTypeCounter *****
type patternFlowLacpduTerminatorTypeCounter struct {
	validation
	obj          *otg.PatternFlowLacpduTerminatorTypeCounter
	marshaller   marshalPatternFlowLacpduTerminatorTypeCounter
	unMarshaller unMarshalPatternFlowLacpduTerminatorTypeCounter
}

func NewPatternFlowLacpduTerminatorTypeCounter() PatternFlowLacpduTerminatorTypeCounter {
	obj := patternFlowLacpduTerminatorTypeCounter{obj: &otg.PatternFlowLacpduTerminatorTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduTerminatorTypeCounter) msg() *otg.PatternFlowLacpduTerminatorTypeCounter {
	return obj.obj
}

func (obj *patternFlowLacpduTerminatorTypeCounter) setMsg(msg *otg.PatternFlowLacpduTerminatorTypeCounter) PatternFlowLacpduTerminatorTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduTerminatorTypeCounter struct {
	obj *patternFlowLacpduTerminatorTypeCounter
}

type marshalPatternFlowLacpduTerminatorTypeCounter interface {
	// ToProto marshals PatternFlowLacpduTerminatorTypeCounter to protobuf object *otg.PatternFlowLacpduTerminatorTypeCounter
	ToProto() (*otg.PatternFlowLacpduTerminatorTypeCounter, error)
	// ToPbText marshals PatternFlowLacpduTerminatorTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduTerminatorTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduTerminatorTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduTerminatorTypeCounter struct {
	obj *patternFlowLacpduTerminatorTypeCounter
}

type unMarshalPatternFlowLacpduTerminatorTypeCounter interface {
	// FromProto unmarshals PatternFlowLacpduTerminatorTypeCounter from protobuf object *otg.PatternFlowLacpduTerminatorTypeCounter
	FromProto(msg *otg.PatternFlowLacpduTerminatorTypeCounter) (PatternFlowLacpduTerminatorTypeCounter, error)
	// FromPbText unmarshals PatternFlowLacpduTerminatorTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduTerminatorTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduTerminatorTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduTerminatorTypeCounter) Marshal() marshalPatternFlowLacpduTerminatorTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduTerminatorTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduTerminatorTypeCounter) Unmarshal() unMarshalPatternFlowLacpduTerminatorTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduTerminatorTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduTerminatorTypeCounter) ToProto() (*otg.PatternFlowLacpduTerminatorTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduTerminatorTypeCounter) FromProto(msg *otg.PatternFlowLacpduTerminatorTypeCounter) (PatternFlowLacpduTerminatorTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduTerminatorTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduTerminatorTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduTerminatorTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduTerminatorTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduTerminatorTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduTerminatorTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduTerminatorTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduTerminatorTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduTerminatorTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduTerminatorTypeCounter) Clone() (PatternFlowLacpduTerminatorTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduTerminatorTypeCounter()
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

// PatternFlowLacpduTerminatorTypeCounter is integer counter pattern
type PatternFlowLacpduTerminatorTypeCounter interface {
	Validation
	// msg marshals PatternFlowLacpduTerminatorTypeCounter to protobuf object *otg.PatternFlowLacpduTerminatorTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduTerminatorTypeCounter
	// setMsg unmarshals PatternFlowLacpduTerminatorTypeCounter from protobuf object *otg.PatternFlowLacpduTerminatorTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduTerminatorTypeCounter) PatternFlowLacpduTerminatorTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduTerminatorTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduTerminatorTypeCounter
	// validate validates PatternFlowLacpduTerminatorTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduTerminatorTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduTerminatorTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduTerminatorTypeCounter
	SetStart(value uint32) PatternFlowLacpduTerminatorTypeCounter
	// HasStart checks if Start has been set in PatternFlowLacpduTerminatorTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduTerminatorTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduTerminatorTypeCounter
	SetStep(value uint32) PatternFlowLacpduTerminatorTypeCounter
	// HasStep checks if Step has been set in PatternFlowLacpduTerminatorTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduTerminatorTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduTerminatorTypeCounter
	SetCount(value uint32) PatternFlowLacpduTerminatorTypeCounter
	// HasCount checks if Count has been set in PatternFlowLacpduTerminatorTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduTerminatorTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduTerminatorTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduTerminatorTypeCounter object
func (obj *patternFlowLacpduTerminatorTypeCounter) SetStart(value uint32) PatternFlowLacpduTerminatorTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduTerminatorTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduTerminatorTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduTerminatorTypeCounter object
func (obj *patternFlowLacpduTerminatorTypeCounter) SetStep(value uint32) PatternFlowLacpduTerminatorTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduTerminatorTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduTerminatorTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduTerminatorTypeCounter object
func (obj *patternFlowLacpduTerminatorTypeCounter) SetCount(value uint32) PatternFlowLacpduTerminatorTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduTerminatorTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduTerminatorTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduTerminatorTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduTerminatorTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduTerminatorTypeCounter) setDefault() {
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
