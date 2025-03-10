package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpSeqNumCounter *****
type patternFlowTcpSeqNumCounter struct {
	validation
	obj          *otg.PatternFlowTcpSeqNumCounter
	marshaller   marshalPatternFlowTcpSeqNumCounter
	unMarshaller unMarshalPatternFlowTcpSeqNumCounter
}

func NewPatternFlowTcpSeqNumCounter() PatternFlowTcpSeqNumCounter {
	obj := patternFlowTcpSeqNumCounter{obj: &otg.PatternFlowTcpSeqNumCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpSeqNumCounter) msg() *otg.PatternFlowTcpSeqNumCounter {
	return obj.obj
}

func (obj *patternFlowTcpSeqNumCounter) setMsg(msg *otg.PatternFlowTcpSeqNumCounter) PatternFlowTcpSeqNumCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpSeqNumCounter struct {
	obj *patternFlowTcpSeqNumCounter
}

type marshalPatternFlowTcpSeqNumCounter interface {
	// ToProto marshals PatternFlowTcpSeqNumCounter to protobuf object *otg.PatternFlowTcpSeqNumCounter
	ToProto() (*otg.PatternFlowTcpSeqNumCounter, error)
	// ToPbText marshals PatternFlowTcpSeqNumCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpSeqNumCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpSeqNumCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpSeqNumCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpSeqNumCounter struct {
	obj *patternFlowTcpSeqNumCounter
}

type unMarshalPatternFlowTcpSeqNumCounter interface {
	// FromProto unmarshals PatternFlowTcpSeqNumCounter from protobuf object *otg.PatternFlowTcpSeqNumCounter
	FromProto(msg *otg.PatternFlowTcpSeqNumCounter) (PatternFlowTcpSeqNumCounter, error)
	// FromPbText unmarshals PatternFlowTcpSeqNumCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpSeqNumCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpSeqNumCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpSeqNumCounter) Marshal() marshalPatternFlowTcpSeqNumCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpSeqNumCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpSeqNumCounter) Unmarshal() unMarshalPatternFlowTcpSeqNumCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpSeqNumCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpSeqNumCounter) ToProto() (*otg.PatternFlowTcpSeqNumCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpSeqNumCounter) FromProto(msg *otg.PatternFlowTcpSeqNumCounter) (PatternFlowTcpSeqNumCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpSeqNumCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpSeqNumCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpSeqNumCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpSeqNumCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpSeqNumCounter) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowTcpSeqNumCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpSeqNumCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpSeqNumCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpSeqNumCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpSeqNumCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpSeqNumCounter) Clone() (PatternFlowTcpSeqNumCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpSeqNumCounter()
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

// PatternFlowTcpSeqNumCounter is integer counter pattern
type PatternFlowTcpSeqNumCounter interface {
	Validation
	// msg marshals PatternFlowTcpSeqNumCounter to protobuf object *otg.PatternFlowTcpSeqNumCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpSeqNumCounter
	// setMsg unmarshals PatternFlowTcpSeqNumCounter from protobuf object *otg.PatternFlowTcpSeqNumCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpSeqNumCounter) PatternFlowTcpSeqNumCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpSeqNumCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpSeqNumCounter
	// validate validates PatternFlowTcpSeqNumCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpSeqNumCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpSeqNumCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpSeqNumCounter
	SetStart(value uint32) PatternFlowTcpSeqNumCounter
	// HasStart checks if Start has been set in PatternFlowTcpSeqNumCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpSeqNumCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpSeqNumCounter
	SetStep(value uint32) PatternFlowTcpSeqNumCounter
	// HasStep checks if Step has been set in PatternFlowTcpSeqNumCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpSeqNumCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpSeqNumCounter
	SetCount(value uint32) PatternFlowTcpSeqNumCounter
	// HasCount checks if Count has been set in PatternFlowTcpSeqNumCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpSeqNumCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpSeqNumCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpSeqNumCounter object
func (obj *patternFlowTcpSeqNumCounter) SetStart(value uint32) PatternFlowTcpSeqNumCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpSeqNumCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpSeqNumCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpSeqNumCounter object
func (obj *patternFlowTcpSeqNumCounter) SetStep(value uint32) PatternFlowTcpSeqNumCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpSeqNumCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpSeqNumCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpSeqNumCounter object
func (obj *patternFlowTcpSeqNumCounter) SetCount(value uint32) PatternFlowTcpSeqNumCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpSeqNumCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowTcpSeqNumCounter) setDefault() {
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
