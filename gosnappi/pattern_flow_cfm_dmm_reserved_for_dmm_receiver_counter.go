package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmmReservedForDmmReceiverCounter *****
type patternFlowCfmDmmReservedForDmmReceiverCounter struct {
	validation
	obj          *otg.PatternFlowCfmDmmReservedForDmmReceiverCounter
	marshaller   marshalPatternFlowCfmDmmReservedForDmmReceiverCounter
	unMarshaller unMarshalPatternFlowCfmDmmReservedForDmmReceiverCounter
}

func NewPatternFlowCfmDmmReservedForDmmReceiverCounter() PatternFlowCfmDmmReservedForDmmReceiverCounter {
	obj := patternFlowCfmDmmReservedForDmmReceiverCounter{obj: &otg.PatternFlowCfmDmmReservedForDmmReceiverCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) msg() *otg.PatternFlowCfmDmmReservedForDmmReceiverCounter {
	return obj.obj
}

func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) setMsg(msg *otg.PatternFlowCfmDmmReservedForDmmReceiverCounter) PatternFlowCfmDmmReservedForDmmReceiverCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmmReservedForDmmReceiverCounter struct {
	obj *patternFlowCfmDmmReservedForDmmReceiverCounter
}

type marshalPatternFlowCfmDmmReservedForDmmReceiverCounter interface {
	// ToProto marshals PatternFlowCfmDmmReservedForDmmReceiverCounter to protobuf object *otg.PatternFlowCfmDmmReservedForDmmReceiverCounter
	ToProto() (*otg.PatternFlowCfmDmmReservedForDmmReceiverCounter, error)
	// ToPbText marshals PatternFlowCfmDmmReservedForDmmReceiverCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmmReservedForDmmReceiverCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmmReservedForDmmReceiverCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmmReservedForDmmReceiverCounter struct {
	obj *patternFlowCfmDmmReservedForDmmReceiverCounter
}

type unMarshalPatternFlowCfmDmmReservedForDmmReceiverCounter interface {
	// FromProto unmarshals PatternFlowCfmDmmReservedForDmmReceiverCounter from protobuf object *otg.PatternFlowCfmDmmReservedForDmmReceiverCounter
	FromProto(msg *otg.PatternFlowCfmDmmReservedForDmmReceiverCounter) (PatternFlowCfmDmmReservedForDmmReceiverCounter, error)
	// FromPbText unmarshals PatternFlowCfmDmmReservedForDmmReceiverCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmmReservedForDmmReceiverCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmmReservedForDmmReceiverCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) Marshal() marshalPatternFlowCfmDmmReservedForDmmReceiverCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmmReservedForDmmReceiverCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) Unmarshal() unMarshalPatternFlowCfmDmmReservedForDmmReceiverCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmmReservedForDmmReceiverCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmmReservedForDmmReceiverCounter) ToProto() (*otg.PatternFlowCfmDmmReservedForDmmReceiverCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmmReservedForDmmReceiverCounter) FromProto(msg *otg.PatternFlowCfmDmmReservedForDmmReceiverCounter) (PatternFlowCfmDmmReservedForDmmReceiverCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmmReservedForDmmReceiverCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmmReceiverCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmmReservedForDmmReceiverCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmmReceiverCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmmReservedForDmmReceiverCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmmReceiverCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) Clone() (PatternFlowCfmDmmReservedForDmmReceiverCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmmReservedForDmmReceiverCounter()
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

// PatternFlowCfmDmmReservedForDmmReceiverCounter is integer counter pattern
type PatternFlowCfmDmmReservedForDmmReceiverCounter interface {
	Validation
	// msg marshals PatternFlowCfmDmmReservedForDmmReceiverCounter to protobuf object *otg.PatternFlowCfmDmmReservedForDmmReceiverCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmmReservedForDmmReceiverCounter
	// setMsg unmarshals PatternFlowCfmDmmReservedForDmmReceiverCounter from protobuf object *otg.PatternFlowCfmDmmReservedForDmmReceiverCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmmReservedForDmmReceiverCounter) PatternFlowCfmDmmReservedForDmmReceiverCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmmReservedForDmmReceiverCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmmReservedForDmmReceiverCounter
	// validate validates PatternFlowCfmDmmReservedForDmmReceiverCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmmReservedForDmmReceiverCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint64, set in PatternFlowCfmDmmReservedForDmmReceiverCounter.
	Start() uint64
	// SetStart assigns uint64 provided by user to PatternFlowCfmDmmReservedForDmmReceiverCounter
	SetStart(value uint64) PatternFlowCfmDmmReservedForDmmReceiverCounter
	// HasStart checks if Start has been set in PatternFlowCfmDmmReservedForDmmReceiverCounter
	HasStart() bool
	// Step returns uint64, set in PatternFlowCfmDmmReservedForDmmReceiverCounter.
	Step() uint64
	// SetStep assigns uint64 provided by user to PatternFlowCfmDmmReservedForDmmReceiverCounter
	SetStep(value uint64) PatternFlowCfmDmmReservedForDmmReceiverCounter
	// HasStep checks if Step has been set in PatternFlowCfmDmmReservedForDmmReceiverCounter
	HasStep() bool
	// Count returns uint64, set in PatternFlowCfmDmmReservedForDmmReceiverCounter.
	Count() uint64
	// SetCount assigns uint64 provided by user to PatternFlowCfmDmmReservedForDmmReceiverCounter
	SetCount(value uint64) PatternFlowCfmDmmReservedForDmmReceiverCounter
	// HasCount checks if Count has been set in PatternFlowCfmDmmReservedForDmmReceiverCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) Start() uint64 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint64 value in the PatternFlowCfmDmmReservedForDmmReceiverCounter object
func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) SetStart(value uint64) PatternFlowCfmDmmReservedForDmmReceiverCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) Step() uint64 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint64 value in the PatternFlowCfmDmmReservedForDmmReceiverCounter object
func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) SetStep(value uint64) PatternFlowCfmDmmReservedForDmmReceiverCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) Count() uint64 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint64 value in the PatternFlowCfmDmmReservedForDmmReceiverCounter object
func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) SetCount(value uint64) PatternFlowCfmDmmReservedForDmmReceiverCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowCfmDmmReservedForDmmReceiverCounter) setDefault() {
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
