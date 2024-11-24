package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpAckNumCounter *****
type patternFlowTcpAckNumCounter struct {
	validation
	obj          *otg.PatternFlowTcpAckNumCounter
	marshaller   marshalPatternFlowTcpAckNumCounter
	unMarshaller unMarshalPatternFlowTcpAckNumCounter
}

func NewPatternFlowTcpAckNumCounter() PatternFlowTcpAckNumCounter {
	obj := patternFlowTcpAckNumCounter{obj: &otg.PatternFlowTcpAckNumCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpAckNumCounter) msg() *otg.PatternFlowTcpAckNumCounter {
	return obj.obj
}

func (obj *patternFlowTcpAckNumCounter) setMsg(msg *otg.PatternFlowTcpAckNumCounter) PatternFlowTcpAckNumCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpAckNumCounter struct {
	obj *patternFlowTcpAckNumCounter
}

type marshalPatternFlowTcpAckNumCounter interface {
	// ToProto marshals PatternFlowTcpAckNumCounter to protobuf object *otg.PatternFlowTcpAckNumCounter
	ToProto() (*otg.PatternFlowTcpAckNumCounter, error)
	// ToPbText marshals PatternFlowTcpAckNumCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpAckNumCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpAckNumCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpAckNumCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpAckNumCounter struct {
	obj *patternFlowTcpAckNumCounter
}

type unMarshalPatternFlowTcpAckNumCounter interface {
	// FromProto unmarshals PatternFlowTcpAckNumCounter from protobuf object *otg.PatternFlowTcpAckNumCounter
	FromProto(msg *otg.PatternFlowTcpAckNumCounter) (PatternFlowTcpAckNumCounter, error)
	// FromPbText unmarshals PatternFlowTcpAckNumCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpAckNumCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpAckNumCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpAckNumCounter) Marshal() marshalPatternFlowTcpAckNumCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpAckNumCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpAckNumCounter) Unmarshal() unMarshalPatternFlowTcpAckNumCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpAckNumCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpAckNumCounter) ToProto() (*otg.PatternFlowTcpAckNumCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpAckNumCounter) FromProto(msg *otg.PatternFlowTcpAckNumCounter) (PatternFlowTcpAckNumCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpAckNumCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpAckNumCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpAckNumCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpAckNumCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpAckNumCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowTcpAckNumCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpAckNumCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpAckNumCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpAckNumCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpAckNumCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpAckNumCounter) Clone() (PatternFlowTcpAckNumCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpAckNumCounter()
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

// PatternFlowTcpAckNumCounter is integer counter pattern
type PatternFlowTcpAckNumCounter interface {
	Validation
	// msg marshals PatternFlowTcpAckNumCounter to protobuf object *otg.PatternFlowTcpAckNumCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpAckNumCounter
	// setMsg unmarshals PatternFlowTcpAckNumCounter from protobuf object *otg.PatternFlowTcpAckNumCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpAckNumCounter) PatternFlowTcpAckNumCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpAckNumCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpAckNumCounter
	// validate validates PatternFlowTcpAckNumCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpAckNumCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpAckNumCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpAckNumCounter
	SetStart(value uint32) PatternFlowTcpAckNumCounter
	// HasStart checks if Start has been set in PatternFlowTcpAckNumCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpAckNumCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpAckNumCounter
	SetStep(value uint32) PatternFlowTcpAckNumCounter
	// HasStep checks if Step has been set in PatternFlowTcpAckNumCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpAckNumCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpAckNumCounter
	SetCount(value uint32) PatternFlowTcpAckNumCounter
	// HasCount checks if Count has been set in PatternFlowTcpAckNumCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpAckNumCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpAckNumCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpAckNumCounter object
func (obj *patternFlowTcpAckNumCounter) SetStart(value uint32) PatternFlowTcpAckNumCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpAckNumCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpAckNumCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpAckNumCounter object
func (obj *patternFlowTcpAckNumCounter) SetStep(value uint32) PatternFlowTcpAckNumCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpAckNumCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpAckNumCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpAckNumCounter object
func (obj *patternFlowTcpAckNumCounter) SetCount(value uint32) PatternFlowTcpAckNumCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpAckNumCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowTcpAckNumCounter) setDefault() {
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
