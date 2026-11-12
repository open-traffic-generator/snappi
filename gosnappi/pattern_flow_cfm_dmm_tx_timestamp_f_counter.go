package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmmTxTimestampFCounter *****
type patternFlowCfmDmmTxTimestampFCounter struct {
	validation
	obj          *otg.PatternFlowCfmDmmTxTimestampFCounter
	marshaller   marshalPatternFlowCfmDmmTxTimestampFCounter
	unMarshaller unMarshalPatternFlowCfmDmmTxTimestampFCounter
}

func NewPatternFlowCfmDmmTxTimestampFCounter() PatternFlowCfmDmmTxTimestampFCounter {
	obj := patternFlowCfmDmmTxTimestampFCounter{obj: &otg.PatternFlowCfmDmmTxTimestampFCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmmTxTimestampFCounter) msg() *otg.PatternFlowCfmDmmTxTimestampFCounter {
	return obj.obj
}

func (obj *patternFlowCfmDmmTxTimestampFCounter) setMsg(msg *otg.PatternFlowCfmDmmTxTimestampFCounter) PatternFlowCfmDmmTxTimestampFCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmmTxTimestampFCounter struct {
	obj *patternFlowCfmDmmTxTimestampFCounter
}

type marshalPatternFlowCfmDmmTxTimestampFCounter interface {
	// ToProto marshals PatternFlowCfmDmmTxTimestampFCounter to protobuf object *otg.PatternFlowCfmDmmTxTimestampFCounter
	ToProto() (*otg.PatternFlowCfmDmmTxTimestampFCounter, error)
	// ToPbText marshals PatternFlowCfmDmmTxTimestampFCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmmTxTimestampFCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmmTxTimestampFCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmmTxTimestampFCounter struct {
	obj *patternFlowCfmDmmTxTimestampFCounter
}

type unMarshalPatternFlowCfmDmmTxTimestampFCounter interface {
	// FromProto unmarshals PatternFlowCfmDmmTxTimestampFCounter from protobuf object *otg.PatternFlowCfmDmmTxTimestampFCounter
	FromProto(msg *otg.PatternFlowCfmDmmTxTimestampFCounter) (PatternFlowCfmDmmTxTimestampFCounter, error)
	// FromPbText unmarshals PatternFlowCfmDmmTxTimestampFCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmmTxTimestampFCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmmTxTimestampFCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmmTxTimestampFCounter) Marshal() marshalPatternFlowCfmDmmTxTimestampFCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmmTxTimestampFCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmmTxTimestampFCounter) Unmarshal() unMarshalPatternFlowCfmDmmTxTimestampFCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmmTxTimestampFCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmmTxTimestampFCounter) ToProto() (*otg.PatternFlowCfmDmmTxTimestampFCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmmTxTimestampFCounter) FromProto(msg *otg.PatternFlowCfmDmmTxTimestampFCounter) (PatternFlowCfmDmmTxTimestampFCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmmTxTimestampFCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmTxTimestampFCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmmTxTimestampFCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmTxTimestampFCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmmTxTimestampFCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmTxTimestampFCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmDmmTxTimestampFCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmTxTimestampFCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmTxTimestampFCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmmTxTimestampFCounter) Clone() (PatternFlowCfmDmmTxTimestampFCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmmTxTimestampFCounter()
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

// PatternFlowCfmDmmTxTimestampFCounter is integer counter pattern
type PatternFlowCfmDmmTxTimestampFCounter interface {
	Validation
	// msg marshals PatternFlowCfmDmmTxTimestampFCounter to protobuf object *otg.PatternFlowCfmDmmTxTimestampFCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmmTxTimestampFCounter
	// setMsg unmarshals PatternFlowCfmDmmTxTimestampFCounter from protobuf object *otg.PatternFlowCfmDmmTxTimestampFCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmmTxTimestampFCounter) PatternFlowCfmDmmTxTimestampFCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmmTxTimestampFCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmmTxTimestampFCounter
	// validate validates PatternFlowCfmDmmTxTimestampFCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmmTxTimestampFCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint64, set in PatternFlowCfmDmmTxTimestampFCounter.
	Start() uint64
	// SetStart assigns uint64 provided by user to PatternFlowCfmDmmTxTimestampFCounter
	SetStart(value uint64) PatternFlowCfmDmmTxTimestampFCounter
	// HasStart checks if Start has been set in PatternFlowCfmDmmTxTimestampFCounter
	HasStart() bool
	// Step returns uint64, set in PatternFlowCfmDmmTxTimestampFCounter.
	Step() uint64
	// SetStep assigns uint64 provided by user to PatternFlowCfmDmmTxTimestampFCounter
	SetStep(value uint64) PatternFlowCfmDmmTxTimestampFCounter
	// HasStep checks if Step has been set in PatternFlowCfmDmmTxTimestampFCounter
	HasStep() bool
	// Count returns uint64, set in PatternFlowCfmDmmTxTimestampFCounter.
	Count() uint64
	// SetCount assigns uint64 provided by user to PatternFlowCfmDmmTxTimestampFCounter
	SetCount(value uint64) PatternFlowCfmDmmTxTimestampFCounter
	// HasCount checks if Count has been set in PatternFlowCfmDmmTxTimestampFCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmmTxTimestampFCounter) Start() uint64 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmmTxTimestampFCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint64 value in the PatternFlowCfmDmmTxTimestampFCounter object
func (obj *patternFlowCfmDmmTxTimestampFCounter) SetStart(value uint64) PatternFlowCfmDmmTxTimestampFCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmmTxTimestampFCounter) Step() uint64 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmmTxTimestampFCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint64 value in the PatternFlowCfmDmmTxTimestampFCounter object
func (obj *patternFlowCfmDmmTxTimestampFCounter) SetStep(value uint64) PatternFlowCfmDmmTxTimestampFCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmmTxTimestampFCounter) Count() uint64 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmmTxTimestampFCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint64 value in the PatternFlowCfmDmmTxTimestampFCounter object
func (obj *patternFlowCfmDmmTxTimestampFCounter) SetCount(value uint64) PatternFlowCfmDmmTxTimestampFCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmDmmTxTimestampFCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowCfmDmmTxTimestampFCounter) setDefault() {
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
