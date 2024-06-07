package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGreVersionCounter *****
type patternFlowGreVersionCounter struct {
	validation
	obj          *otg.PatternFlowGreVersionCounter
	marshaller   marshalPatternFlowGreVersionCounter
	unMarshaller unMarshalPatternFlowGreVersionCounter
}

func NewPatternFlowGreVersionCounter() PatternFlowGreVersionCounter {
	obj := patternFlowGreVersionCounter{obj: &otg.PatternFlowGreVersionCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreVersionCounter) msg() *otg.PatternFlowGreVersionCounter {
	return obj.obj
}

func (obj *patternFlowGreVersionCounter) setMsg(msg *otg.PatternFlowGreVersionCounter) PatternFlowGreVersionCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreVersionCounter struct {
	obj *patternFlowGreVersionCounter
}

type marshalPatternFlowGreVersionCounter interface {
	// ToProto marshals PatternFlowGreVersionCounter to protobuf object *otg.PatternFlowGreVersionCounter
	ToProto() (*otg.PatternFlowGreVersionCounter, error)
	// ToPbText marshals PatternFlowGreVersionCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreVersionCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreVersionCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGreVersionCounter struct {
	obj *patternFlowGreVersionCounter
}

type unMarshalPatternFlowGreVersionCounter interface {
	// FromProto unmarshals PatternFlowGreVersionCounter from protobuf object *otg.PatternFlowGreVersionCounter
	FromProto(msg *otg.PatternFlowGreVersionCounter) (PatternFlowGreVersionCounter, error)
	// FromPbText unmarshals PatternFlowGreVersionCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreVersionCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreVersionCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreVersionCounter) Marshal() marshalPatternFlowGreVersionCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreVersionCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreVersionCounter) Unmarshal() unMarshalPatternFlowGreVersionCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreVersionCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreVersionCounter) ToProto() (*otg.PatternFlowGreVersionCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreVersionCounter) FromProto(msg *otg.PatternFlowGreVersionCounter) (PatternFlowGreVersionCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreVersionCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreVersionCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreVersionCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreVersionCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreVersionCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreVersionCounter) FromJson(value string) error {
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

func (obj *patternFlowGreVersionCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreVersionCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreVersionCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreVersionCounter) Clone() (PatternFlowGreVersionCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreVersionCounter()
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

// PatternFlowGreVersionCounter is integer counter pattern
type PatternFlowGreVersionCounter interface {
	Validation
	// msg marshals PatternFlowGreVersionCounter to protobuf object *otg.PatternFlowGreVersionCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGreVersionCounter
	// setMsg unmarshals PatternFlowGreVersionCounter from protobuf object *otg.PatternFlowGreVersionCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreVersionCounter) PatternFlowGreVersionCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGreVersionCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreVersionCounter
	// validate validates PatternFlowGreVersionCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreVersionCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGreVersionCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGreVersionCounter
	SetStart(value uint32) PatternFlowGreVersionCounter
	// HasStart checks if Start has been set in PatternFlowGreVersionCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGreVersionCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGreVersionCounter
	SetStep(value uint32) PatternFlowGreVersionCounter
	// HasStep checks if Step has been set in PatternFlowGreVersionCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGreVersionCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGreVersionCounter
	SetCount(value uint32) PatternFlowGreVersionCounter
	// HasCount checks if Count has been set in PatternFlowGreVersionCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGreVersionCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGreVersionCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGreVersionCounter object
func (obj *patternFlowGreVersionCounter) SetStart(value uint32) PatternFlowGreVersionCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGreVersionCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGreVersionCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGreVersionCounter object
func (obj *patternFlowGreVersionCounter) SetStep(value uint32) PatternFlowGreVersionCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGreVersionCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGreVersionCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGreVersionCounter object
func (obj *patternFlowGreVersionCounter) SetCount(value uint32) PatternFlowGreVersionCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGreVersionCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreVersionCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreVersionCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreVersionCounter.Count <= 7 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGreVersionCounter) setDefault() {
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
