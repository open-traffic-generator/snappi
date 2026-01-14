package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVlanCfiCounter *****
type patternFlowVlanCfiCounter struct {
	validation
	obj          *otg.PatternFlowVlanCfiCounter
	marshaller   marshalPatternFlowVlanCfiCounter
	unMarshaller unMarshalPatternFlowVlanCfiCounter
}

func NewPatternFlowVlanCfiCounter() PatternFlowVlanCfiCounter {
	obj := patternFlowVlanCfiCounter{obj: &otg.PatternFlowVlanCfiCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVlanCfiCounter) msg() *otg.PatternFlowVlanCfiCounter {
	return obj.obj
}

func (obj *patternFlowVlanCfiCounter) setMsg(msg *otg.PatternFlowVlanCfiCounter) PatternFlowVlanCfiCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVlanCfiCounter struct {
	obj *patternFlowVlanCfiCounter
}

type marshalPatternFlowVlanCfiCounter interface {
	// ToProto marshals PatternFlowVlanCfiCounter to protobuf object *otg.PatternFlowVlanCfiCounter
	ToProto() (*otg.PatternFlowVlanCfiCounter, error)
	// ToPbText marshals PatternFlowVlanCfiCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVlanCfiCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVlanCfiCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVlanCfiCounter struct {
	obj *patternFlowVlanCfiCounter
}

type unMarshalPatternFlowVlanCfiCounter interface {
	// FromProto unmarshals PatternFlowVlanCfiCounter from protobuf object *otg.PatternFlowVlanCfiCounter
	FromProto(msg *otg.PatternFlowVlanCfiCounter) (PatternFlowVlanCfiCounter, error)
	// FromPbText unmarshals PatternFlowVlanCfiCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVlanCfiCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVlanCfiCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVlanCfiCounter) Marshal() marshalPatternFlowVlanCfiCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVlanCfiCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVlanCfiCounter) Unmarshal() unMarshalPatternFlowVlanCfiCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVlanCfiCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVlanCfiCounter) ToProto() (*otg.PatternFlowVlanCfiCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVlanCfiCounter) FromProto(msg *otg.PatternFlowVlanCfiCounter) (PatternFlowVlanCfiCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVlanCfiCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVlanCfiCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowVlanCfiCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVlanCfiCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowVlanCfiCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVlanCfiCounter) FromJson(value string) error {
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

func (obj *patternFlowVlanCfiCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVlanCfiCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVlanCfiCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVlanCfiCounter) Clone() (PatternFlowVlanCfiCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVlanCfiCounter()
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

// PatternFlowVlanCfiCounter is integer counter pattern
type PatternFlowVlanCfiCounter interface {
	Validation
	// msg marshals PatternFlowVlanCfiCounter to protobuf object *otg.PatternFlowVlanCfiCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowVlanCfiCounter
	// setMsg unmarshals PatternFlowVlanCfiCounter from protobuf object *otg.PatternFlowVlanCfiCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVlanCfiCounter) PatternFlowVlanCfiCounter
	// provides marshal interface
	Marshal() marshalPatternFlowVlanCfiCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVlanCfiCounter
	// validate validates PatternFlowVlanCfiCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVlanCfiCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowVlanCfiCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowVlanCfiCounter
	SetStart(value uint32) PatternFlowVlanCfiCounter
	// HasStart checks if Start has been set in PatternFlowVlanCfiCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowVlanCfiCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowVlanCfiCounter
	SetStep(value uint32) PatternFlowVlanCfiCounter
	// HasStep checks if Step has been set in PatternFlowVlanCfiCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowVlanCfiCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowVlanCfiCounter
	SetCount(value uint32) PatternFlowVlanCfiCounter
	// HasCount checks if Count has been set in PatternFlowVlanCfiCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVlanCfiCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVlanCfiCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowVlanCfiCounter object
func (obj *patternFlowVlanCfiCounter) SetStart(value uint32) PatternFlowVlanCfiCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVlanCfiCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVlanCfiCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowVlanCfiCounter object
func (obj *patternFlowVlanCfiCounter) SetStep(value uint32) PatternFlowVlanCfiCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVlanCfiCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVlanCfiCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowVlanCfiCounter object
func (obj *patternFlowVlanCfiCounter) SetCount(value uint32) PatternFlowVlanCfiCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowVlanCfiCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanCfiCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanCfiCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanCfiCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowVlanCfiCounter) setDefault() {
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
