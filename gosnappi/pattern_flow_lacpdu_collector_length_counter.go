package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduCollectorLengthCounter *****
type patternFlowLacpduCollectorLengthCounter struct {
	validation
	obj          *otg.PatternFlowLacpduCollectorLengthCounter
	marshaller   marshalPatternFlowLacpduCollectorLengthCounter
	unMarshaller unMarshalPatternFlowLacpduCollectorLengthCounter
}

func NewPatternFlowLacpduCollectorLengthCounter() PatternFlowLacpduCollectorLengthCounter {
	obj := patternFlowLacpduCollectorLengthCounter{obj: &otg.PatternFlowLacpduCollectorLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduCollectorLengthCounter) msg() *otg.PatternFlowLacpduCollectorLengthCounter {
	return obj.obj
}

func (obj *patternFlowLacpduCollectorLengthCounter) setMsg(msg *otg.PatternFlowLacpduCollectorLengthCounter) PatternFlowLacpduCollectorLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduCollectorLengthCounter struct {
	obj *patternFlowLacpduCollectorLengthCounter
}

type marshalPatternFlowLacpduCollectorLengthCounter interface {
	// ToProto marshals PatternFlowLacpduCollectorLengthCounter to protobuf object *otg.PatternFlowLacpduCollectorLengthCounter
	ToProto() (*otg.PatternFlowLacpduCollectorLengthCounter, error)
	// ToPbText marshals PatternFlowLacpduCollectorLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduCollectorLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduCollectorLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduCollectorLengthCounter struct {
	obj *patternFlowLacpduCollectorLengthCounter
}

type unMarshalPatternFlowLacpduCollectorLengthCounter interface {
	// FromProto unmarshals PatternFlowLacpduCollectorLengthCounter from protobuf object *otg.PatternFlowLacpduCollectorLengthCounter
	FromProto(msg *otg.PatternFlowLacpduCollectorLengthCounter) (PatternFlowLacpduCollectorLengthCounter, error)
	// FromPbText unmarshals PatternFlowLacpduCollectorLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduCollectorLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduCollectorLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduCollectorLengthCounter) Marshal() marshalPatternFlowLacpduCollectorLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduCollectorLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduCollectorLengthCounter) Unmarshal() unMarshalPatternFlowLacpduCollectorLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduCollectorLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduCollectorLengthCounter) ToProto() (*otg.PatternFlowLacpduCollectorLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduCollectorLengthCounter) FromProto(msg *otg.PatternFlowLacpduCollectorLengthCounter) (PatternFlowLacpduCollectorLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduCollectorLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduCollectorLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduCollectorLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduCollectorLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduCollectorLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduCollectorLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduCollectorLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduCollectorLengthCounter) Clone() (PatternFlowLacpduCollectorLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduCollectorLengthCounter()
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

// PatternFlowLacpduCollectorLengthCounter is integer counter pattern
type PatternFlowLacpduCollectorLengthCounter interface {
	Validation
	// msg marshals PatternFlowLacpduCollectorLengthCounter to protobuf object *otg.PatternFlowLacpduCollectorLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduCollectorLengthCounter
	// setMsg unmarshals PatternFlowLacpduCollectorLengthCounter from protobuf object *otg.PatternFlowLacpduCollectorLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduCollectorLengthCounter) PatternFlowLacpduCollectorLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduCollectorLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduCollectorLengthCounter
	// validate validates PatternFlowLacpduCollectorLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduCollectorLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduCollectorLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduCollectorLengthCounter
	SetStart(value uint32) PatternFlowLacpduCollectorLengthCounter
	// HasStart checks if Start has been set in PatternFlowLacpduCollectorLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduCollectorLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduCollectorLengthCounter
	SetStep(value uint32) PatternFlowLacpduCollectorLengthCounter
	// HasStep checks if Step has been set in PatternFlowLacpduCollectorLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduCollectorLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduCollectorLengthCounter
	SetCount(value uint32) PatternFlowLacpduCollectorLengthCounter
	// HasCount checks if Count has been set in PatternFlowLacpduCollectorLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduCollectorLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduCollectorLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduCollectorLengthCounter object
func (obj *patternFlowLacpduCollectorLengthCounter) SetStart(value uint32) PatternFlowLacpduCollectorLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduCollectorLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduCollectorLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduCollectorLengthCounter object
func (obj *patternFlowLacpduCollectorLengthCounter) SetStep(value uint32) PatternFlowLacpduCollectorLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduCollectorLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduCollectorLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduCollectorLengthCounter object
func (obj *patternFlowLacpduCollectorLengthCounter) SetCount(value uint32) PatternFlowLacpduCollectorLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduCollectorLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduCollectorLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduCollectorLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduCollectorLengthCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduCollectorLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(16)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
