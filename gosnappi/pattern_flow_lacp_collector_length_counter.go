package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpCollectorLengthCounter *****
type patternFlowLacpCollectorLengthCounter struct {
	validation
	obj          *otg.PatternFlowLacpCollectorLengthCounter
	marshaller   marshalPatternFlowLacpCollectorLengthCounter
	unMarshaller unMarshalPatternFlowLacpCollectorLengthCounter
}

func NewPatternFlowLacpCollectorLengthCounter() PatternFlowLacpCollectorLengthCounter {
	obj := patternFlowLacpCollectorLengthCounter{obj: &otg.PatternFlowLacpCollectorLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpCollectorLengthCounter) msg() *otg.PatternFlowLacpCollectorLengthCounter {
	return obj.obj
}

func (obj *patternFlowLacpCollectorLengthCounter) setMsg(msg *otg.PatternFlowLacpCollectorLengthCounter) PatternFlowLacpCollectorLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpCollectorLengthCounter struct {
	obj *patternFlowLacpCollectorLengthCounter
}

type marshalPatternFlowLacpCollectorLengthCounter interface {
	// ToProto marshals PatternFlowLacpCollectorLengthCounter to protobuf object *otg.PatternFlowLacpCollectorLengthCounter
	ToProto() (*otg.PatternFlowLacpCollectorLengthCounter, error)
	// ToPbText marshals PatternFlowLacpCollectorLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpCollectorLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpCollectorLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpCollectorLengthCounter struct {
	obj *patternFlowLacpCollectorLengthCounter
}

type unMarshalPatternFlowLacpCollectorLengthCounter interface {
	// FromProto unmarshals PatternFlowLacpCollectorLengthCounter from protobuf object *otg.PatternFlowLacpCollectorLengthCounter
	FromProto(msg *otg.PatternFlowLacpCollectorLengthCounter) (PatternFlowLacpCollectorLengthCounter, error)
	// FromPbText unmarshals PatternFlowLacpCollectorLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpCollectorLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpCollectorLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpCollectorLengthCounter) Marshal() marshalPatternFlowLacpCollectorLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpCollectorLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpCollectorLengthCounter) Unmarshal() unMarshalPatternFlowLacpCollectorLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpCollectorLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpCollectorLengthCounter) ToProto() (*otg.PatternFlowLacpCollectorLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpCollectorLengthCounter) FromProto(msg *otg.PatternFlowLacpCollectorLengthCounter) (PatternFlowLacpCollectorLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpCollectorLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpCollectorLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpCollectorLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpCollectorLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpCollectorLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpCollectorLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpCollectorLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpCollectorLengthCounter) Clone() (PatternFlowLacpCollectorLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpCollectorLengthCounter()
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

// PatternFlowLacpCollectorLengthCounter is integer counter pattern
type PatternFlowLacpCollectorLengthCounter interface {
	Validation
	// msg marshals PatternFlowLacpCollectorLengthCounter to protobuf object *otg.PatternFlowLacpCollectorLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpCollectorLengthCounter
	// setMsg unmarshals PatternFlowLacpCollectorLengthCounter from protobuf object *otg.PatternFlowLacpCollectorLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpCollectorLengthCounter) PatternFlowLacpCollectorLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpCollectorLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpCollectorLengthCounter
	// validate validates PatternFlowLacpCollectorLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpCollectorLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpCollectorLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpCollectorLengthCounter
	SetStart(value uint32) PatternFlowLacpCollectorLengthCounter
	// HasStart checks if Start has been set in PatternFlowLacpCollectorLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpCollectorLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpCollectorLengthCounter
	SetStep(value uint32) PatternFlowLacpCollectorLengthCounter
	// HasStep checks if Step has been set in PatternFlowLacpCollectorLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpCollectorLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpCollectorLengthCounter
	SetCount(value uint32) PatternFlowLacpCollectorLengthCounter
	// HasCount checks if Count has been set in PatternFlowLacpCollectorLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpCollectorLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpCollectorLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpCollectorLengthCounter object
func (obj *patternFlowLacpCollectorLengthCounter) SetStart(value uint32) PatternFlowLacpCollectorLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpCollectorLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpCollectorLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpCollectorLengthCounter object
func (obj *patternFlowLacpCollectorLengthCounter) SetStep(value uint32) PatternFlowLacpCollectorLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpCollectorLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpCollectorLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpCollectorLengthCounter object
func (obj *patternFlowLacpCollectorLengthCounter) SetCount(value uint32) PatternFlowLacpCollectorLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpCollectorLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpCollectorLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpCollectorLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpCollectorLengthCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpCollectorLengthCounter) setDefault() {
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
