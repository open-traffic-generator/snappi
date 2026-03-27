package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPppControlCounter *****
type patternFlowPppControlCounter struct {
	validation
	obj          *otg.PatternFlowPppControlCounter
	marshaller   marshalPatternFlowPppControlCounter
	unMarshaller unMarshalPatternFlowPppControlCounter
}

func NewPatternFlowPppControlCounter() PatternFlowPppControlCounter {
	obj := patternFlowPppControlCounter{obj: &otg.PatternFlowPppControlCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPppControlCounter) msg() *otg.PatternFlowPppControlCounter {
	return obj.obj
}

func (obj *patternFlowPppControlCounter) setMsg(msg *otg.PatternFlowPppControlCounter) PatternFlowPppControlCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPppControlCounter struct {
	obj *patternFlowPppControlCounter
}

type marshalPatternFlowPppControlCounter interface {
	// ToProto marshals PatternFlowPppControlCounter to protobuf object *otg.PatternFlowPppControlCounter
	ToProto() (*otg.PatternFlowPppControlCounter, error)
	// ToPbText marshals PatternFlowPppControlCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPppControlCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPppControlCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPppControlCounter struct {
	obj *patternFlowPppControlCounter
}

type unMarshalPatternFlowPppControlCounter interface {
	// FromProto unmarshals PatternFlowPppControlCounter from protobuf object *otg.PatternFlowPppControlCounter
	FromProto(msg *otg.PatternFlowPppControlCounter) (PatternFlowPppControlCounter, error)
	// FromPbText unmarshals PatternFlowPppControlCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPppControlCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPppControlCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPppControlCounter) Marshal() marshalPatternFlowPppControlCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPppControlCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPppControlCounter) Unmarshal() unMarshalPatternFlowPppControlCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPppControlCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPppControlCounter) ToProto() (*otg.PatternFlowPppControlCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPppControlCounter) FromProto(msg *otg.PatternFlowPppControlCounter) (PatternFlowPppControlCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPppControlCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPppControlCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPppControlCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPppControlCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPppControlCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPppControlCounter) FromJson(value string) error {
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

func (obj *patternFlowPppControlCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPppControlCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPppControlCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPppControlCounter) Clone() (PatternFlowPppControlCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPppControlCounter()
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

// PatternFlowPppControlCounter is integer counter pattern
type PatternFlowPppControlCounter interface {
	Validation
	// msg marshals PatternFlowPppControlCounter to protobuf object *otg.PatternFlowPppControlCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowPppControlCounter
	// setMsg unmarshals PatternFlowPppControlCounter from protobuf object *otg.PatternFlowPppControlCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPppControlCounter) PatternFlowPppControlCounter
	// provides marshal interface
	Marshal() marshalPatternFlowPppControlCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPppControlCounter
	// validate validates PatternFlowPppControlCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPppControlCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPppControlCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPppControlCounter
	SetStart(value uint32) PatternFlowPppControlCounter
	// HasStart checks if Start has been set in PatternFlowPppControlCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPppControlCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPppControlCounter
	SetStep(value uint32) PatternFlowPppControlCounter
	// HasStep checks if Step has been set in PatternFlowPppControlCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPppControlCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPppControlCounter
	SetCount(value uint32) PatternFlowPppControlCounter
	// HasCount checks if Count has been set in PatternFlowPppControlCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPppControlCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPppControlCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPppControlCounter object
func (obj *patternFlowPppControlCounter) SetStart(value uint32) PatternFlowPppControlCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPppControlCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPppControlCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPppControlCounter object
func (obj *patternFlowPppControlCounter) SetStep(value uint32) PatternFlowPppControlCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPppControlCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPppControlCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPppControlCounter object
func (obj *patternFlowPppControlCounter) SetCount(value uint32) PatternFlowPppControlCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPppControlCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppControlCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppControlCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppControlCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPppControlCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(3)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
