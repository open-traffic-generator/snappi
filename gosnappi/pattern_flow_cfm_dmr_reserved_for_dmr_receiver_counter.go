package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmrReservedForDmrReceiverCounter *****
type patternFlowCfmDmrReservedForDmrReceiverCounter struct {
	validation
	obj          *otg.PatternFlowCfmDmrReservedForDmrReceiverCounter
	marshaller   marshalPatternFlowCfmDmrReservedForDmrReceiverCounter
	unMarshaller unMarshalPatternFlowCfmDmrReservedForDmrReceiverCounter
}

func NewPatternFlowCfmDmrReservedForDmrReceiverCounter() PatternFlowCfmDmrReservedForDmrReceiverCounter {
	obj := patternFlowCfmDmrReservedForDmrReceiverCounter{obj: &otg.PatternFlowCfmDmrReservedForDmrReceiverCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) msg() *otg.PatternFlowCfmDmrReservedForDmrReceiverCounter {
	return obj.obj
}

func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) setMsg(msg *otg.PatternFlowCfmDmrReservedForDmrReceiverCounter) PatternFlowCfmDmrReservedForDmrReceiverCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmrReservedForDmrReceiverCounter struct {
	obj *patternFlowCfmDmrReservedForDmrReceiverCounter
}

type marshalPatternFlowCfmDmrReservedForDmrReceiverCounter interface {
	// ToProto marshals PatternFlowCfmDmrReservedForDmrReceiverCounter to protobuf object *otg.PatternFlowCfmDmrReservedForDmrReceiverCounter
	ToProto() (*otg.PatternFlowCfmDmrReservedForDmrReceiverCounter, error)
	// ToPbText marshals PatternFlowCfmDmrReservedForDmrReceiverCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmrReservedForDmrReceiverCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmrReservedForDmrReceiverCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmrReservedForDmrReceiverCounter struct {
	obj *patternFlowCfmDmrReservedForDmrReceiverCounter
}

type unMarshalPatternFlowCfmDmrReservedForDmrReceiverCounter interface {
	// FromProto unmarshals PatternFlowCfmDmrReservedForDmrReceiverCounter from protobuf object *otg.PatternFlowCfmDmrReservedForDmrReceiverCounter
	FromProto(msg *otg.PatternFlowCfmDmrReservedForDmrReceiverCounter) (PatternFlowCfmDmrReservedForDmrReceiverCounter, error)
	// FromPbText unmarshals PatternFlowCfmDmrReservedForDmrReceiverCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmrReservedForDmrReceiverCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmrReservedForDmrReceiverCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) Marshal() marshalPatternFlowCfmDmrReservedForDmrReceiverCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmrReservedForDmrReceiverCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) Unmarshal() unMarshalPatternFlowCfmDmrReservedForDmrReceiverCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmrReservedForDmrReceiverCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmrReservedForDmrReceiverCounter) ToProto() (*otg.PatternFlowCfmDmrReservedForDmrReceiverCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmrReservedForDmrReceiverCounter) FromProto(msg *otg.PatternFlowCfmDmrReservedForDmrReceiverCounter) (PatternFlowCfmDmrReservedForDmrReceiverCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmrReservedForDmrReceiverCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrReservedForDmrReceiverCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmrReservedForDmrReceiverCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrReservedForDmrReceiverCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmrReservedForDmrReceiverCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrReservedForDmrReceiverCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) Clone() (PatternFlowCfmDmrReservedForDmrReceiverCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmrReservedForDmrReceiverCounter()
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

// PatternFlowCfmDmrReservedForDmrReceiverCounter is integer counter pattern
type PatternFlowCfmDmrReservedForDmrReceiverCounter interface {
	Validation
	// msg marshals PatternFlowCfmDmrReservedForDmrReceiverCounter to protobuf object *otg.PatternFlowCfmDmrReservedForDmrReceiverCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmrReservedForDmrReceiverCounter
	// setMsg unmarshals PatternFlowCfmDmrReservedForDmrReceiverCounter from protobuf object *otg.PatternFlowCfmDmrReservedForDmrReceiverCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmrReservedForDmrReceiverCounter) PatternFlowCfmDmrReservedForDmrReceiverCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmrReservedForDmrReceiverCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmrReservedForDmrReceiverCounter
	// validate validates PatternFlowCfmDmrReservedForDmrReceiverCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmrReservedForDmrReceiverCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint64, set in PatternFlowCfmDmrReservedForDmrReceiverCounter.
	Start() uint64
	// SetStart assigns uint64 provided by user to PatternFlowCfmDmrReservedForDmrReceiverCounter
	SetStart(value uint64) PatternFlowCfmDmrReservedForDmrReceiverCounter
	// HasStart checks if Start has been set in PatternFlowCfmDmrReservedForDmrReceiverCounter
	HasStart() bool
	// Step returns uint64, set in PatternFlowCfmDmrReservedForDmrReceiverCounter.
	Step() uint64
	// SetStep assigns uint64 provided by user to PatternFlowCfmDmrReservedForDmrReceiverCounter
	SetStep(value uint64) PatternFlowCfmDmrReservedForDmrReceiverCounter
	// HasStep checks if Step has been set in PatternFlowCfmDmrReservedForDmrReceiverCounter
	HasStep() bool
	// Count returns uint64, set in PatternFlowCfmDmrReservedForDmrReceiverCounter.
	Count() uint64
	// SetCount assigns uint64 provided by user to PatternFlowCfmDmrReservedForDmrReceiverCounter
	SetCount(value uint64) PatternFlowCfmDmrReservedForDmrReceiverCounter
	// HasCount checks if Count has been set in PatternFlowCfmDmrReservedForDmrReceiverCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) Start() uint64 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint64 value in the PatternFlowCfmDmrReservedForDmrReceiverCounter object
func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) SetStart(value uint64) PatternFlowCfmDmrReservedForDmrReceiverCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) Step() uint64 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint64 value in the PatternFlowCfmDmrReservedForDmrReceiverCounter object
func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) SetStep(value uint64) PatternFlowCfmDmrReservedForDmrReceiverCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) Count() uint64 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint64 value in the PatternFlowCfmDmrReservedForDmrReceiverCounter object
func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) SetCount(value uint64) PatternFlowCfmDmrReservedForDmrReceiverCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowCfmDmrReservedForDmrReceiverCounter) setDefault() {
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
