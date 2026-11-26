package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmmReservedForDmrCounter *****
type patternFlowCfmDmmReservedForDmrCounter struct {
	validation
	obj          *otg.PatternFlowCfmDmmReservedForDmrCounter
	marshaller   marshalPatternFlowCfmDmmReservedForDmrCounter
	unMarshaller unMarshalPatternFlowCfmDmmReservedForDmrCounter
}

func NewPatternFlowCfmDmmReservedForDmrCounter() PatternFlowCfmDmmReservedForDmrCounter {
	obj := patternFlowCfmDmmReservedForDmrCounter{obj: &otg.PatternFlowCfmDmmReservedForDmrCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmmReservedForDmrCounter) msg() *otg.PatternFlowCfmDmmReservedForDmrCounter {
	return obj.obj
}

func (obj *patternFlowCfmDmmReservedForDmrCounter) setMsg(msg *otg.PatternFlowCfmDmmReservedForDmrCounter) PatternFlowCfmDmmReservedForDmrCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmmReservedForDmrCounter struct {
	obj *patternFlowCfmDmmReservedForDmrCounter
}

type marshalPatternFlowCfmDmmReservedForDmrCounter interface {
	// ToProto marshals PatternFlowCfmDmmReservedForDmrCounter to protobuf object *otg.PatternFlowCfmDmmReservedForDmrCounter
	ToProto() (*otg.PatternFlowCfmDmmReservedForDmrCounter, error)
	// ToPbText marshals PatternFlowCfmDmmReservedForDmrCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmmReservedForDmrCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmmReservedForDmrCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmmReservedForDmrCounter struct {
	obj *patternFlowCfmDmmReservedForDmrCounter
}

type unMarshalPatternFlowCfmDmmReservedForDmrCounter interface {
	// FromProto unmarshals PatternFlowCfmDmmReservedForDmrCounter from protobuf object *otg.PatternFlowCfmDmmReservedForDmrCounter
	FromProto(msg *otg.PatternFlowCfmDmmReservedForDmrCounter) (PatternFlowCfmDmmReservedForDmrCounter, error)
	// FromPbText unmarshals PatternFlowCfmDmmReservedForDmrCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmmReservedForDmrCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmmReservedForDmrCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmmReservedForDmrCounter) Marshal() marshalPatternFlowCfmDmmReservedForDmrCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmmReservedForDmrCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmmReservedForDmrCounter) Unmarshal() unMarshalPatternFlowCfmDmmReservedForDmrCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmmReservedForDmrCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmmReservedForDmrCounter) ToProto() (*otg.PatternFlowCfmDmmReservedForDmrCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmmReservedForDmrCounter) FromProto(msg *otg.PatternFlowCfmDmmReservedForDmrCounter) (PatternFlowCfmDmmReservedForDmrCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmmReservedForDmrCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmrCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmmReservedForDmrCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmrCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmmReservedForDmrCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmrCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmDmmReservedForDmrCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmReservedForDmrCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmReservedForDmrCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmmReservedForDmrCounter) Clone() (PatternFlowCfmDmmReservedForDmrCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmmReservedForDmrCounter()
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

// PatternFlowCfmDmmReservedForDmrCounter is integer counter pattern
type PatternFlowCfmDmmReservedForDmrCounter interface {
	Validation
	// msg marshals PatternFlowCfmDmmReservedForDmrCounter to protobuf object *otg.PatternFlowCfmDmmReservedForDmrCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmmReservedForDmrCounter
	// setMsg unmarshals PatternFlowCfmDmmReservedForDmrCounter from protobuf object *otg.PatternFlowCfmDmmReservedForDmrCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmmReservedForDmrCounter) PatternFlowCfmDmmReservedForDmrCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmmReservedForDmrCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmmReservedForDmrCounter
	// validate validates PatternFlowCfmDmmReservedForDmrCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmmReservedForDmrCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint64, set in PatternFlowCfmDmmReservedForDmrCounter.
	Start() uint64
	// SetStart assigns uint64 provided by user to PatternFlowCfmDmmReservedForDmrCounter
	SetStart(value uint64) PatternFlowCfmDmmReservedForDmrCounter
	// HasStart checks if Start has been set in PatternFlowCfmDmmReservedForDmrCounter
	HasStart() bool
	// Step returns uint64, set in PatternFlowCfmDmmReservedForDmrCounter.
	Step() uint64
	// SetStep assigns uint64 provided by user to PatternFlowCfmDmmReservedForDmrCounter
	SetStep(value uint64) PatternFlowCfmDmmReservedForDmrCounter
	// HasStep checks if Step has been set in PatternFlowCfmDmmReservedForDmrCounter
	HasStep() bool
	// Count returns uint64, set in PatternFlowCfmDmmReservedForDmrCounter.
	Count() uint64
	// SetCount assigns uint64 provided by user to PatternFlowCfmDmmReservedForDmrCounter
	SetCount(value uint64) PatternFlowCfmDmmReservedForDmrCounter
	// HasCount checks if Count has been set in PatternFlowCfmDmmReservedForDmrCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrCounter) Start() uint64 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint64 value in the PatternFlowCfmDmmReservedForDmrCounter object
func (obj *patternFlowCfmDmmReservedForDmrCounter) SetStart(value uint64) PatternFlowCfmDmmReservedForDmrCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrCounter) Step() uint64 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint64 value in the PatternFlowCfmDmmReservedForDmrCounter object
func (obj *patternFlowCfmDmmReservedForDmrCounter) SetStep(value uint64) PatternFlowCfmDmmReservedForDmrCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrCounter) Count() uint64 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint64 value in the PatternFlowCfmDmmReservedForDmrCounter object
func (obj *patternFlowCfmDmmReservedForDmrCounter) SetCount(value uint64) PatternFlowCfmDmmReservedForDmrCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmDmmReservedForDmrCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowCfmDmmReservedForDmrCounter) setDefault() {
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
