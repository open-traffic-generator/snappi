package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmmReservedForDmrReceiverCounter *****
type patternFlowCfmDmmReservedForDmrReceiverCounter struct {
	validation
	obj          *otg.PatternFlowCfmDmmReservedForDmrReceiverCounter
	marshaller   marshalPatternFlowCfmDmmReservedForDmrReceiverCounter
	unMarshaller unMarshalPatternFlowCfmDmmReservedForDmrReceiverCounter
}

func NewPatternFlowCfmDmmReservedForDmrReceiverCounter() PatternFlowCfmDmmReservedForDmrReceiverCounter {
	obj := patternFlowCfmDmmReservedForDmrReceiverCounter{obj: &otg.PatternFlowCfmDmmReservedForDmrReceiverCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) msg() *otg.PatternFlowCfmDmmReservedForDmrReceiverCounter {
	return obj.obj
}

func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) setMsg(msg *otg.PatternFlowCfmDmmReservedForDmrReceiverCounter) PatternFlowCfmDmmReservedForDmrReceiverCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmmReservedForDmrReceiverCounter struct {
	obj *patternFlowCfmDmmReservedForDmrReceiverCounter
}

type marshalPatternFlowCfmDmmReservedForDmrReceiverCounter interface {
	// ToProto marshals PatternFlowCfmDmmReservedForDmrReceiverCounter to protobuf object *otg.PatternFlowCfmDmmReservedForDmrReceiverCounter
	ToProto() (*otg.PatternFlowCfmDmmReservedForDmrReceiverCounter, error)
	// ToPbText marshals PatternFlowCfmDmmReservedForDmrReceiverCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmmReservedForDmrReceiverCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmmReservedForDmrReceiverCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmmReservedForDmrReceiverCounter struct {
	obj *patternFlowCfmDmmReservedForDmrReceiverCounter
}

type unMarshalPatternFlowCfmDmmReservedForDmrReceiverCounter interface {
	// FromProto unmarshals PatternFlowCfmDmmReservedForDmrReceiverCounter from protobuf object *otg.PatternFlowCfmDmmReservedForDmrReceiverCounter
	FromProto(msg *otg.PatternFlowCfmDmmReservedForDmrReceiverCounter) (PatternFlowCfmDmmReservedForDmrReceiverCounter, error)
	// FromPbText unmarshals PatternFlowCfmDmmReservedForDmrReceiverCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmmReservedForDmrReceiverCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmmReservedForDmrReceiverCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) Marshal() marshalPatternFlowCfmDmmReservedForDmrReceiverCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmmReservedForDmrReceiverCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) Unmarshal() unMarshalPatternFlowCfmDmmReservedForDmrReceiverCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmmReservedForDmrReceiverCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmmReservedForDmrReceiverCounter) ToProto() (*otg.PatternFlowCfmDmmReservedForDmrReceiverCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmmReservedForDmrReceiverCounter) FromProto(msg *otg.PatternFlowCfmDmmReservedForDmrReceiverCounter) (PatternFlowCfmDmmReservedForDmrReceiverCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmmReservedForDmrReceiverCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmrReceiverCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmmReservedForDmrReceiverCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmrReceiverCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmmReservedForDmrReceiverCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmrReceiverCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) Clone() (PatternFlowCfmDmmReservedForDmrReceiverCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmmReservedForDmrReceiverCounter()
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

// PatternFlowCfmDmmReservedForDmrReceiverCounter is integer counter pattern
type PatternFlowCfmDmmReservedForDmrReceiverCounter interface {
	Validation
	// msg marshals PatternFlowCfmDmmReservedForDmrReceiverCounter to protobuf object *otg.PatternFlowCfmDmmReservedForDmrReceiverCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmmReservedForDmrReceiverCounter
	// setMsg unmarshals PatternFlowCfmDmmReservedForDmrReceiverCounter from protobuf object *otg.PatternFlowCfmDmmReservedForDmrReceiverCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmmReservedForDmrReceiverCounter) PatternFlowCfmDmmReservedForDmrReceiverCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmmReservedForDmrReceiverCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmmReservedForDmrReceiverCounter
	// validate validates PatternFlowCfmDmmReservedForDmrReceiverCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmmReservedForDmrReceiverCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint64, set in PatternFlowCfmDmmReservedForDmrReceiverCounter.
	Start() uint64
	// SetStart assigns uint64 provided by user to PatternFlowCfmDmmReservedForDmrReceiverCounter
	SetStart(value uint64) PatternFlowCfmDmmReservedForDmrReceiverCounter
	// HasStart checks if Start has been set in PatternFlowCfmDmmReservedForDmrReceiverCounter
	HasStart() bool
	// Step returns uint64, set in PatternFlowCfmDmmReservedForDmrReceiverCounter.
	Step() uint64
	// SetStep assigns uint64 provided by user to PatternFlowCfmDmmReservedForDmrReceiverCounter
	SetStep(value uint64) PatternFlowCfmDmmReservedForDmrReceiverCounter
	// HasStep checks if Step has been set in PatternFlowCfmDmmReservedForDmrReceiverCounter
	HasStep() bool
	// Count returns uint64, set in PatternFlowCfmDmmReservedForDmrReceiverCounter.
	Count() uint64
	// SetCount assigns uint64 provided by user to PatternFlowCfmDmmReservedForDmrReceiverCounter
	SetCount(value uint64) PatternFlowCfmDmmReservedForDmrReceiverCounter
	// HasCount checks if Count has been set in PatternFlowCfmDmmReservedForDmrReceiverCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) Start() uint64 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint64 value in the PatternFlowCfmDmmReservedForDmrReceiverCounter object
func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) SetStart(value uint64) PatternFlowCfmDmmReservedForDmrReceiverCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) Step() uint64 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint64 value in the PatternFlowCfmDmmReservedForDmrReceiverCounter object
func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) SetStep(value uint64) PatternFlowCfmDmmReservedForDmrReceiverCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) Count() uint64 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint64 value in the PatternFlowCfmDmmReservedForDmrReceiverCounter object
func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) SetCount(value uint64) PatternFlowCfmDmmReservedForDmrReceiverCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowCfmDmmReservedForDmrReceiverCounter) setDefault() {
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
