package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmCfmShortMaNameLengthCounter *****
type patternFlowCfmCfmShortMaNameLengthCounter struct {
	validation
	obj          *otg.PatternFlowCfmCfmShortMaNameLengthCounter
	marshaller   marshalPatternFlowCfmCfmShortMaNameLengthCounter
	unMarshaller unMarshalPatternFlowCfmCfmShortMaNameLengthCounter
}

func NewPatternFlowCfmCfmShortMaNameLengthCounter() PatternFlowCfmCfmShortMaNameLengthCounter {
	obj := patternFlowCfmCfmShortMaNameLengthCounter{obj: &otg.PatternFlowCfmCfmShortMaNameLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmCfmShortMaNameLengthCounter) msg() *otg.PatternFlowCfmCfmShortMaNameLengthCounter {
	return obj.obj
}

func (obj *patternFlowCfmCfmShortMaNameLengthCounter) setMsg(msg *otg.PatternFlowCfmCfmShortMaNameLengthCounter) PatternFlowCfmCfmShortMaNameLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmCfmShortMaNameLengthCounter struct {
	obj *patternFlowCfmCfmShortMaNameLengthCounter
}

type marshalPatternFlowCfmCfmShortMaNameLengthCounter interface {
	// ToProto marshals PatternFlowCfmCfmShortMaNameLengthCounter to protobuf object *otg.PatternFlowCfmCfmShortMaNameLengthCounter
	ToProto() (*otg.PatternFlowCfmCfmShortMaNameLengthCounter, error)
	// ToPbText marshals PatternFlowCfmCfmShortMaNameLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmCfmShortMaNameLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmCfmShortMaNameLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmCfmShortMaNameLengthCounter struct {
	obj *patternFlowCfmCfmShortMaNameLengthCounter
}

type unMarshalPatternFlowCfmCfmShortMaNameLengthCounter interface {
	// FromProto unmarshals PatternFlowCfmCfmShortMaNameLengthCounter from protobuf object *otg.PatternFlowCfmCfmShortMaNameLengthCounter
	FromProto(msg *otg.PatternFlowCfmCfmShortMaNameLengthCounter) (PatternFlowCfmCfmShortMaNameLengthCounter, error)
	// FromPbText unmarshals PatternFlowCfmCfmShortMaNameLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmCfmShortMaNameLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmCfmShortMaNameLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmCfmShortMaNameLengthCounter) Marshal() marshalPatternFlowCfmCfmShortMaNameLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmCfmShortMaNameLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmCfmShortMaNameLengthCounter) Unmarshal() unMarshalPatternFlowCfmCfmShortMaNameLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmCfmShortMaNameLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmCfmShortMaNameLengthCounter) ToProto() (*otg.PatternFlowCfmCfmShortMaNameLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmCfmShortMaNameLengthCounter) FromProto(msg *otg.PatternFlowCfmCfmShortMaNameLengthCounter) (PatternFlowCfmCfmShortMaNameLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmCfmShortMaNameLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmCfmShortMaNameLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmCfmShortMaNameLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmCfmShortMaNameLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmCfmShortMaNameLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmCfmShortMaNameLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmCfmShortMaNameLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmCfmShortMaNameLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmCfmShortMaNameLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmCfmShortMaNameLengthCounter) Clone() (PatternFlowCfmCfmShortMaNameLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmCfmShortMaNameLengthCounter()
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

// PatternFlowCfmCfmShortMaNameLengthCounter is integer counter pattern
type PatternFlowCfmCfmShortMaNameLengthCounter interface {
	Validation
	// msg marshals PatternFlowCfmCfmShortMaNameLengthCounter to protobuf object *otg.PatternFlowCfmCfmShortMaNameLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmCfmShortMaNameLengthCounter
	// setMsg unmarshals PatternFlowCfmCfmShortMaNameLengthCounter from protobuf object *otg.PatternFlowCfmCfmShortMaNameLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmCfmShortMaNameLengthCounter) PatternFlowCfmCfmShortMaNameLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmCfmShortMaNameLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmCfmShortMaNameLengthCounter
	// validate validates PatternFlowCfmCfmShortMaNameLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmCfmShortMaNameLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowCfmCfmShortMaNameLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowCfmCfmShortMaNameLengthCounter
	SetStart(value uint32) PatternFlowCfmCfmShortMaNameLengthCounter
	// HasStart checks if Start has been set in PatternFlowCfmCfmShortMaNameLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowCfmCfmShortMaNameLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowCfmCfmShortMaNameLengthCounter
	SetStep(value uint32) PatternFlowCfmCfmShortMaNameLengthCounter
	// HasStep checks if Step has been set in PatternFlowCfmCfmShortMaNameLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowCfmCfmShortMaNameLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowCfmCfmShortMaNameLengthCounter
	SetCount(value uint32) PatternFlowCfmCfmShortMaNameLengthCounter
	// HasCount checks if Count has been set in PatternFlowCfmCfmShortMaNameLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmCfmShortMaNameLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmCfmShortMaNameLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowCfmCfmShortMaNameLengthCounter object
func (obj *patternFlowCfmCfmShortMaNameLengthCounter) SetStart(value uint32) PatternFlowCfmCfmShortMaNameLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmCfmShortMaNameLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmCfmShortMaNameLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowCfmCfmShortMaNameLengthCounter object
func (obj *patternFlowCfmCfmShortMaNameLengthCounter) SetStep(value uint32) PatternFlowCfmCfmShortMaNameLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmCfmShortMaNameLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmCfmShortMaNameLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowCfmCfmShortMaNameLengthCounter object
func (obj *patternFlowCfmCfmShortMaNameLengthCounter) SetCount(value uint32) PatternFlowCfmCfmShortMaNameLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmCfmShortMaNameLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmCfmShortMaNameLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmCfmShortMaNameLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmCfmShortMaNameLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowCfmCfmShortMaNameLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
