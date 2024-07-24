package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVlanIdCounter *****
type patternFlowVlanIdCounter struct {
	validation
	obj          *otg.PatternFlowVlanIdCounter
	marshaller   marshalPatternFlowVlanIdCounter
	unMarshaller unMarshalPatternFlowVlanIdCounter
}

func NewPatternFlowVlanIdCounter() PatternFlowVlanIdCounter {
	obj := patternFlowVlanIdCounter{obj: &otg.PatternFlowVlanIdCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVlanIdCounter) msg() *otg.PatternFlowVlanIdCounter {
	return obj.obj
}

func (obj *patternFlowVlanIdCounter) setMsg(msg *otg.PatternFlowVlanIdCounter) PatternFlowVlanIdCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVlanIdCounter struct {
	obj *patternFlowVlanIdCounter
}

type marshalPatternFlowVlanIdCounter interface {
	// ToProto marshals PatternFlowVlanIdCounter to protobuf object *otg.PatternFlowVlanIdCounter
	ToProto() (*otg.PatternFlowVlanIdCounter, error)
	// ToPbText marshals PatternFlowVlanIdCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVlanIdCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVlanIdCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVlanIdCounter struct {
	obj *patternFlowVlanIdCounter
}

type unMarshalPatternFlowVlanIdCounter interface {
	// FromProto unmarshals PatternFlowVlanIdCounter from protobuf object *otg.PatternFlowVlanIdCounter
	FromProto(msg *otg.PatternFlowVlanIdCounter) (PatternFlowVlanIdCounter, error)
	// FromPbText unmarshals PatternFlowVlanIdCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVlanIdCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVlanIdCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVlanIdCounter) Marshal() marshalPatternFlowVlanIdCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVlanIdCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVlanIdCounter) Unmarshal() unMarshalPatternFlowVlanIdCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVlanIdCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVlanIdCounter) ToProto() (*otg.PatternFlowVlanIdCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVlanIdCounter) FromProto(msg *otg.PatternFlowVlanIdCounter) (PatternFlowVlanIdCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVlanIdCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVlanIdCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowVlanIdCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVlanIdCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowVlanIdCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVlanIdCounter) FromJson(value string) error {
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

func (obj *patternFlowVlanIdCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVlanIdCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVlanIdCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVlanIdCounter) Clone() (PatternFlowVlanIdCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVlanIdCounter()
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

// PatternFlowVlanIdCounter is integer counter pattern
type PatternFlowVlanIdCounter interface {
	Validation
	// msg marshals PatternFlowVlanIdCounter to protobuf object *otg.PatternFlowVlanIdCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowVlanIdCounter
	// setMsg unmarshals PatternFlowVlanIdCounter from protobuf object *otg.PatternFlowVlanIdCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVlanIdCounter) PatternFlowVlanIdCounter
	// provides marshal interface
	Marshal() marshalPatternFlowVlanIdCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVlanIdCounter
	// validate validates PatternFlowVlanIdCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVlanIdCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowVlanIdCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowVlanIdCounter
	SetStart(value uint32) PatternFlowVlanIdCounter
	// HasStart checks if Start has been set in PatternFlowVlanIdCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowVlanIdCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowVlanIdCounter
	SetStep(value uint32) PatternFlowVlanIdCounter
	// HasStep checks if Step has been set in PatternFlowVlanIdCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowVlanIdCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowVlanIdCounter
	SetCount(value uint32) PatternFlowVlanIdCounter
	// HasCount checks if Count has been set in PatternFlowVlanIdCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVlanIdCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVlanIdCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowVlanIdCounter object
func (obj *patternFlowVlanIdCounter) SetStart(value uint32) PatternFlowVlanIdCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVlanIdCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVlanIdCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowVlanIdCounter object
func (obj *patternFlowVlanIdCounter) SetStep(value uint32) PatternFlowVlanIdCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVlanIdCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVlanIdCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowVlanIdCounter object
func (obj *patternFlowVlanIdCounter) SetCount(value uint32) PatternFlowVlanIdCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowVlanIdCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanIdCounter.Start <= 4095 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanIdCounter.Step <= 4095 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanIdCounter.Count <= 4095 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowVlanIdCounter) setDefault() {
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
