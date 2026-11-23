package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmrTxTimestampFCounter *****
type patternFlowCfmDmrTxTimestampFCounter struct {
	validation
	obj          *otg.PatternFlowCfmDmrTxTimestampFCounter
	marshaller   marshalPatternFlowCfmDmrTxTimestampFCounter
	unMarshaller unMarshalPatternFlowCfmDmrTxTimestampFCounter
}

func NewPatternFlowCfmDmrTxTimestampFCounter() PatternFlowCfmDmrTxTimestampFCounter {
	obj := patternFlowCfmDmrTxTimestampFCounter{obj: &otg.PatternFlowCfmDmrTxTimestampFCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmrTxTimestampFCounter) msg() *otg.PatternFlowCfmDmrTxTimestampFCounter {
	return obj.obj
}

func (obj *patternFlowCfmDmrTxTimestampFCounter) setMsg(msg *otg.PatternFlowCfmDmrTxTimestampFCounter) PatternFlowCfmDmrTxTimestampFCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmrTxTimestampFCounter struct {
	obj *patternFlowCfmDmrTxTimestampFCounter
}

type marshalPatternFlowCfmDmrTxTimestampFCounter interface {
	// ToProto marshals PatternFlowCfmDmrTxTimestampFCounter to protobuf object *otg.PatternFlowCfmDmrTxTimestampFCounter
	ToProto() (*otg.PatternFlowCfmDmrTxTimestampFCounter, error)
	// ToPbText marshals PatternFlowCfmDmrTxTimestampFCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmrTxTimestampFCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmrTxTimestampFCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmrTxTimestampFCounter struct {
	obj *patternFlowCfmDmrTxTimestampFCounter
}

type unMarshalPatternFlowCfmDmrTxTimestampFCounter interface {
	// FromProto unmarshals PatternFlowCfmDmrTxTimestampFCounter from protobuf object *otg.PatternFlowCfmDmrTxTimestampFCounter
	FromProto(msg *otg.PatternFlowCfmDmrTxTimestampFCounter) (PatternFlowCfmDmrTxTimestampFCounter, error)
	// FromPbText unmarshals PatternFlowCfmDmrTxTimestampFCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmrTxTimestampFCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmrTxTimestampFCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmrTxTimestampFCounter) Marshal() marshalPatternFlowCfmDmrTxTimestampFCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmrTxTimestampFCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmrTxTimestampFCounter) Unmarshal() unMarshalPatternFlowCfmDmrTxTimestampFCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmrTxTimestampFCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmrTxTimestampFCounter) ToProto() (*otg.PatternFlowCfmDmrTxTimestampFCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmrTxTimestampFCounter) FromProto(msg *otg.PatternFlowCfmDmrTxTimestampFCounter) (PatternFlowCfmDmrTxTimestampFCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmrTxTimestampFCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrTxTimestampFCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmrTxTimestampFCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrTxTimestampFCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmrTxTimestampFCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrTxTimestampFCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmDmrTxTimestampFCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrTxTimestampFCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrTxTimestampFCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmrTxTimestampFCounter) Clone() (PatternFlowCfmDmrTxTimestampFCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmrTxTimestampFCounter()
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

// PatternFlowCfmDmrTxTimestampFCounter is integer counter pattern
type PatternFlowCfmDmrTxTimestampFCounter interface {
	Validation
	// msg marshals PatternFlowCfmDmrTxTimestampFCounter to protobuf object *otg.PatternFlowCfmDmrTxTimestampFCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmrTxTimestampFCounter
	// setMsg unmarshals PatternFlowCfmDmrTxTimestampFCounter from protobuf object *otg.PatternFlowCfmDmrTxTimestampFCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmrTxTimestampFCounter) PatternFlowCfmDmrTxTimestampFCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmrTxTimestampFCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmrTxTimestampFCounter
	// validate validates PatternFlowCfmDmrTxTimestampFCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmrTxTimestampFCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint64, set in PatternFlowCfmDmrTxTimestampFCounter.
	Start() uint64
	// SetStart assigns uint64 provided by user to PatternFlowCfmDmrTxTimestampFCounter
	SetStart(value uint64) PatternFlowCfmDmrTxTimestampFCounter
	// HasStart checks if Start has been set in PatternFlowCfmDmrTxTimestampFCounter
	HasStart() bool
	// Step returns uint64, set in PatternFlowCfmDmrTxTimestampFCounter.
	Step() uint64
	// SetStep assigns uint64 provided by user to PatternFlowCfmDmrTxTimestampFCounter
	SetStep(value uint64) PatternFlowCfmDmrTxTimestampFCounter
	// HasStep checks if Step has been set in PatternFlowCfmDmrTxTimestampFCounter
	HasStep() bool
	// Count returns uint64, set in PatternFlowCfmDmrTxTimestampFCounter.
	Count() uint64
	// SetCount assigns uint64 provided by user to PatternFlowCfmDmrTxTimestampFCounter
	SetCount(value uint64) PatternFlowCfmDmrTxTimestampFCounter
	// HasCount checks if Count has been set in PatternFlowCfmDmrTxTimestampFCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmrTxTimestampFCounter) Start() uint64 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmrTxTimestampFCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint64 value in the PatternFlowCfmDmrTxTimestampFCounter object
func (obj *patternFlowCfmDmrTxTimestampFCounter) SetStart(value uint64) PatternFlowCfmDmrTxTimestampFCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmrTxTimestampFCounter) Step() uint64 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmrTxTimestampFCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint64 value in the PatternFlowCfmDmrTxTimestampFCounter object
func (obj *patternFlowCfmDmrTxTimestampFCounter) SetStep(value uint64) PatternFlowCfmDmrTxTimestampFCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmrTxTimestampFCounter) Count() uint64 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmrTxTimestampFCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint64 value in the PatternFlowCfmDmrTxTimestampFCounter object
func (obj *patternFlowCfmDmrTxTimestampFCounter) SetCount(value uint64) PatternFlowCfmDmrTxTimestampFCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmDmrTxTimestampFCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowCfmDmrTxTimestampFCounter) setDefault() {
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
