package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmrTxTimestampBCounter *****
type patternFlowCfmDmrTxTimestampBCounter struct {
	validation
	obj          *otg.PatternFlowCfmDmrTxTimestampBCounter
	marshaller   marshalPatternFlowCfmDmrTxTimestampBCounter
	unMarshaller unMarshalPatternFlowCfmDmrTxTimestampBCounter
}

func NewPatternFlowCfmDmrTxTimestampBCounter() PatternFlowCfmDmrTxTimestampBCounter {
	obj := patternFlowCfmDmrTxTimestampBCounter{obj: &otg.PatternFlowCfmDmrTxTimestampBCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmrTxTimestampBCounter) msg() *otg.PatternFlowCfmDmrTxTimestampBCounter {
	return obj.obj
}

func (obj *patternFlowCfmDmrTxTimestampBCounter) setMsg(msg *otg.PatternFlowCfmDmrTxTimestampBCounter) PatternFlowCfmDmrTxTimestampBCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmrTxTimestampBCounter struct {
	obj *patternFlowCfmDmrTxTimestampBCounter
}

type marshalPatternFlowCfmDmrTxTimestampBCounter interface {
	// ToProto marshals PatternFlowCfmDmrTxTimestampBCounter to protobuf object *otg.PatternFlowCfmDmrTxTimestampBCounter
	ToProto() (*otg.PatternFlowCfmDmrTxTimestampBCounter, error)
	// ToPbText marshals PatternFlowCfmDmrTxTimestampBCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmrTxTimestampBCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmrTxTimestampBCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmrTxTimestampBCounter struct {
	obj *patternFlowCfmDmrTxTimestampBCounter
}

type unMarshalPatternFlowCfmDmrTxTimestampBCounter interface {
	// FromProto unmarshals PatternFlowCfmDmrTxTimestampBCounter from protobuf object *otg.PatternFlowCfmDmrTxTimestampBCounter
	FromProto(msg *otg.PatternFlowCfmDmrTxTimestampBCounter) (PatternFlowCfmDmrTxTimestampBCounter, error)
	// FromPbText unmarshals PatternFlowCfmDmrTxTimestampBCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmrTxTimestampBCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmrTxTimestampBCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmrTxTimestampBCounter) Marshal() marshalPatternFlowCfmDmrTxTimestampBCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmrTxTimestampBCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmrTxTimestampBCounter) Unmarshal() unMarshalPatternFlowCfmDmrTxTimestampBCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmrTxTimestampBCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmrTxTimestampBCounter) ToProto() (*otg.PatternFlowCfmDmrTxTimestampBCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmrTxTimestampBCounter) FromProto(msg *otg.PatternFlowCfmDmrTxTimestampBCounter) (PatternFlowCfmDmrTxTimestampBCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmrTxTimestampBCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrTxTimestampBCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmrTxTimestampBCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrTxTimestampBCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmrTxTimestampBCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrTxTimestampBCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmDmrTxTimestampBCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrTxTimestampBCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrTxTimestampBCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmrTxTimestampBCounter) Clone() (PatternFlowCfmDmrTxTimestampBCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmrTxTimestampBCounter()
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

// PatternFlowCfmDmrTxTimestampBCounter is integer counter pattern
type PatternFlowCfmDmrTxTimestampBCounter interface {
	Validation
	// msg marshals PatternFlowCfmDmrTxTimestampBCounter to protobuf object *otg.PatternFlowCfmDmrTxTimestampBCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmrTxTimestampBCounter
	// setMsg unmarshals PatternFlowCfmDmrTxTimestampBCounter from protobuf object *otg.PatternFlowCfmDmrTxTimestampBCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmrTxTimestampBCounter) PatternFlowCfmDmrTxTimestampBCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmrTxTimestampBCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmrTxTimestampBCounter
	// validate validates PatternFlowCfmDmrTxTimestampBCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmrTxTimestampBCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint64, set in PatternFlowCfmDmrTxTimestampBCounter.
	Start() uint64
	// SetStart assigns uint64 provided by user to PatternFlowCfmDmrTxTimestampBCounter
	SetStart(value uint64) PatternFlowCfmDmrTxTimestampBCounter
	// HasStart checks if Start has been set in PatternFlowCfmDmrTxTimestampBCounter
	HasStart() bool
	// Step returns uint64, set in PatternFlowCfmDmrTxTimestampBCounter.
	Step() uint64
	// SetStep assigns uint64 provided by user to PatternFlowCfmDmrTxTimestampBCounter
	SetStep(value uint64) PatternFlowCfmDmrTxTimestampBCounter
	// HasStep checks if Step has been set in PatternFlowCfmDmrTxTimestampBCounter
	HasStep() bool
	// Count returns uint64, set in PatternFlowCfmDmrTxTimestampBCounter.
	Count() uint64
	// SetCount assigns uint64 provided by user to PatternFlowCfmDmrTxTimestampBCounter
	SetCount(value uint64) PatternFlowCfmDmrTxTimestampBCounter
	// HasCount checks if Count has been set in PatternFlowCfmDmrTxTimestampBCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmrTxTimestampBCounter) Start() uint64 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint64
func (obj *patternFlowCfmDmrTxTimestampBCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint64 value in the PatternFlowCfmDmrTxTimestampBCounter object
func (obj *patternFlowCfmDmrTxTimestampBCounter) SetStart(value uint64) PatternFlowCfmDmrTxTimestampBCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmrTxTimestampBCounter) Step() uint64 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint64
func (obj *patternFlowCfmDmrTxTimestampBCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint64 value in the PatternFlowCfmDmrTxTimestampBCounter object
func (obj *patternFlowCfmDmrTxTimestampBCounter) SetStep(value uint64) PatternFlowCfmDmrTxTimestampBCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmrTxTimestampBCounter) Count() uint64 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint64
func (obj *patternFlowCfmDmrTxTimestampBCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint64 value in the PatternFlowCfmDmrTxTimestampBCounter object
func (obj *patternFlowCfmDmrTxTimestampBCounter) SetCount(value uint64) PatternFlowCfmDmrTxTimestampBCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmDmrTxTimestampBCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowCfmDmrTxTimestampBCounter) setDefault() {
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
