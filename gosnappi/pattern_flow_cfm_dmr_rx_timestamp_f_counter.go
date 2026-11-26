package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmrRxTimestampFCounter *****
type patternFlowCfmDmrRxTimestampFCounter struct {
	validation
	obj          *otg.PatternFlowCfmDmrRxTimestampFCounter
	marshaller   marshalPatternFlowCfmDmrRxTimestampFCounter
	unMarshaller unMarshalPatternFlowCfmDmrRxTimestampFCounter
}

func NewPatternFlowCfmDmrRxTimestampFCounter() PatternFlowCfmDmrRxTimestampFCounter {
	obj := patternFlowCfmDmrRxTimestampFCounter{obj: &otg.PatternFlowCfmDmrRxTimestampFCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmrRxTimestampFCounter) msg() *otg.PatternFlowCfmDmrRxTimestampFCounter {
	return obj.obj
}

func (obj *patternFlowCfmDmrRxTimestampFCounter) setMsg(msg *otg.PatternFlowCfmDmrRxTimestampFCounter) PatternFlowCfmDmrRxTimestampFCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmrRxTimestampFCounter struct {
	obj *patternFlowCfmDmrRxTimestampFCounter
}

type marshalPatternFlowCfmDmrRxTimestampFCounter interface {
	// ToProto marshals PatternFlowCfmDmrRxTimestampFCounter to protobuf object *otg.PatternFlowCfmDmrRxTimestampFCounter
	ToProto() (*otg.PatternFlowCfmDmrRxTimestampFCounter, error)
	// ToPbText marshals PatternFlowCfmDmrRxTimestampFCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmrRxTimestampFCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmrRxTimestampFCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmrRxTimestampFCounter struct {
	obj *patternFlowCfmDmrRxTimestampFCounter
}

type unMarshalPatternFlowCfmDmrRxTimestampFCounter interface {
	// FromProto unmarshals PatternFlowCfmDmrRxTimestampFCounter from protobuf object *otg.PatternFlowCfmDmrRxTimestampFCounter
	FromProto(msg *otg.PatternFlowCfmDmrRxTimestampFCounter) (PatternFlowCfmDmrRxTimestampFCounter, error)
	// FromPbText unmarshals PatternFlowCfmDmrRxTimestampFCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmrRxTimestampFCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmrRxTimestampFCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmrRxTimestampFCounter) Marshal() marshalPatternFlowCfmDmrRxTimestampFCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmrRxTimestampFCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmrRxTimestampFCounter) Unmarshal() unMarshalPatternFlowCfmDmrRxTimestampFCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmrRxTimestampFCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmrRxTimestampFCounter) ToProto() (*otg.PatternFlowCfmDmrRxTimestampFCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmrRxTimestampFCounter) FromProto(msg *otg.PatternFlowCfmDmrRxTimestampFCounter) (PatternFlowCfmDmrRxTimestampFCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmrRxTimestampFCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrRxTimestampFCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmrRxTimestampFCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrRxTimestampFCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmrRxTimestampFCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrRxTimestampFCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmDmrRxTimestampFCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrRxTimestampFCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrRxTimestampFCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmrRxTimestampFCounter) Clone() (PatternFlowCfmDmrRxTimestampFCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmrRxTimestampFCounter()
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

// PatternFlowCfmDmrRxTimestampFCounter is integer counter pattern
type PatternFlowCfmDmrRxTimestampFCounter interface {
	Validation
	// msg marshals PatternFlowCfmDmrRxTimestampFCounter to protobuf object *otg.PatternFlowCfmDmrRxTimestampFCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmrRxTimestampFCounter
	// setMsg unmarshals PatternFlowCfmDmrRxTimestampFCounter from protobuf object *otg.PatternFlowCfmDmrRxTimestampFCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmrRxTimestampFCounter) PatternFlowCfmDmrRxTimestampFCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmrRxTimestampFCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmrRxTimestampFCounter
	// validate validates PatternFlowCfmDmrRxTimestampFCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmrRxTimestampFCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint64, set in PatternFlowCfmDmrRxTimestampFCounter.
	Start() uint64
	// SetStart assigns uint64 provided by user to PatternFlowCfmDmrRxTimestampFCounter
	SetStart(value uint64) PatternFlowCfmDmrRxTimestampFCounter
	// HasStart checks if Start has been set in PatternFlowCfmDmrRxTimestampFCounter
	HasStart() bool
	// Step returns uint64, set in PatternFlowCfmDmrRxTimestampFCounter.
	Step() uint64
	// SetStep assigns uint64 provided by user to PatternFlowCfmDmrRxTimestampFCounter
	SetStep(value uint64) PatternFlowCfmDmrRxTimestampFCounter
	// HasStep checks if Step has been set in PatternFlowCfmDmrRxTimestampFCounter
	HasStep() bool
	// Count returns uint64, set in PatternFlowCfmDmrRxTimestampFCounter.
	Count() uint64
	// SetCount assigns uint64 provided by user to PatternFlowCfmDmrRxTimestampFCounter
	SetCount(value uint64) PatternFlowCfmDmrRxTimestampFCounter
	// HasCount checks if Count has been set in PatternFlowCfmDmrRxTimestampFCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmrRxTimestampFCounter) Start() uint64 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmrRxTimestampFCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint64 value in the PatternFlowCfmDmrRxTimestampFCounter object
func (obj *patternFlowCfmDmrRxTimestampFCounter) SetStart(value uint64) PatternFlowCfmDmrRxTimestampFCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmrRxTimestampFCounter) Step() uint64 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmrRxTimestampFCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint64 value in the PatternFlowCfmDmrRxTimestampFCounter object
func (obj *patternFlowCfmDmrRxTimestampFCounter) SetStep(value uint64) PatternFlowCfmDmrRxTimestampFCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmrRxTimestampFCounter) Count() uint64 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmrRxTimestampFCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint64 value in the PatternFlowCfmDmrRxTimestampFCounter object
func (obj *patternFlowCfmDmrRxTimestampFCounter) SetCount(value uint64) PatternFlowCfmDmrRxTimestampFCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmDmrRxTimestampFCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowCfmDmrRxTimestampFCounter) setDefault() {
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
