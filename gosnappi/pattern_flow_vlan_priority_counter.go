package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVlanPriorityCounter *****
type patternFlowVlanPriorityCounter struct {
	validation
	obj          *otg.PatternFlowVlanPriorityCounter
	marshaller   marshalPatternFlowVlanPriorityCounter
	unMarshaller unMarshalPatternFlowVlanPriorityCounter
}

func NewPatternFlowVlanPriorityCounter() PatternFlowVlanPriorityCounter {
	obj := patternFlowVlanPriorityCounter{obj: &otg.PatternFlowVlanPriorityCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVlanPriorityCounter) msg() *otg.PatternFlowVlanPriorityCounter {
	return obj.obj
}

func (obj *patternFlowVlanPriorityCounter) setMsg(msg *otg.PatternFlowVlanPriorityCounter) PatternFlowVlanPriorityCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVlanPriorityCounter struct {
	obj *patternFlowVlanPriorityCounter
}

type marshalPatternFlowVlanPriorityCounter interface {
	// ToProto marshals PatternFlowVlanPriorityCounter to protobuf object *otg.PatternFlowVlanPriorityCounter
	ToProto() (*otg.PatternFlowVlanPriorityCounter, error)
	// ToPbText marshals PatternFlowVlanPriorityCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVlanPriorityCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVlanPriorityCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVlanPriorityCounter struct {
	obj *patternFlowVlanPriorityCounter
}

type unMarshalPatternFlowVlanPriorityCounter interface {
	// FromProto unmarshals PatternFlowVlanPriorityCounter from protobuf object *otg.PatternFlowVlanPriorityCounter
	FromProto(msg *otg.PatternFlowVlanPriorityCounter) (PatternFlowVlanPriorityCounter, error)
	// FromPbText unmarshals PatternFlowVlanPriorityCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVlanPriorityCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVlanPriorityCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVlanPriorityCounter) Marshal() marshalPatternFlowVlanPriorityCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVlanPriorityCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVlanPriorityCounter) Unmarshal() unMarshalPatternFlowVlanPriorityCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVlanPriorityCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVlanPriorityCounter) ToProto() (*otg.PatternFlowVlanPriorityCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVlanPriorityCounter) FromProto(msg *otg.PatternFlowVlanPriorityCounter) (PatternFlowVlanPriorityCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVlanPriorityCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVlanPriorityCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowVlanPriorityCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVlanPriorityCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowVlanPriorityCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVlanPriorityCounter) FromJson(value string) error {
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

func (obj *patternFlowVlanPriorityCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVlanPriorityCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVlanPriorityCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVlanPriorityCounter) Clone() (PatternFlowVlanPriorityCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVlanPriorityCounter()
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

// PatternFlowVlanPriorityCounter is integer counter pattern
type PatternFlowVlanPriorityCounter interface {
	Validation
	// msg marshals PatternFlowVlanPriorityCounter to protobuf object *otg.PatternFlowVlanPriorityCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowVlanPriorityCounter
	// setMsg unmarshals PatternFlowVlanPriorityCounter from protobuf object *otg.PatternFlowVlanPriorityCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVlanPriorityCounter) PatternFlowVlanPriorityCounter
	// provides marshal interface
	Marshal() marshalPatternFlowVlanPriorityCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVlanPriorityCounter
	// validate validates PatternFlowVlanPriorityCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVlanPriorityCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowVlanPriorityCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowVlanPriorityCounter
	SetStart(value uint32) PatternFlowVlanPriorityCounter
	// HasStart checks if Start has been set in PatternFlowVlanPriorityCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowVlanPriorityCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowVlanPriorityCounter
	SetStep(value uint32) PatternFlowVlanPriorityCounter
	// HasStep checks if Step has been set in PatternFlowVlanPriorityCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowVlanPriorityCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowVlanPriorityCounter
	SetCount(value uint32) PatternFlowVlanPriorityCounter
	// HasCount checks if Count has been set in PatternFlowVlanPriorityCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVlanPriorityCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowVlanPriorityCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowVlanPriorityCounter object
func (obj *patternFlowVlanPriorityCounter) SetStart(value uint32) PatternFlowVlanPriorityCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVlanPriorityCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowVlanPriorityCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowVlanPriorityCounter object
func (obj *patternFlowVlanPriorityCounter) SetStep(value uint32) PatternFlowVlanPriorityCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVlanPriorityCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowVlanPriorityCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowVlanPriorityCounter object
func (obj *patternFlowVlanPriorityCounter) SetCount(value uint32) PatternFlowVlanPriorityCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowVlanPriorityCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanPriorityCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanPriorityCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanPriorityCounter.Count <= 8 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowVlanPriorityCounter) setDefault() {
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
