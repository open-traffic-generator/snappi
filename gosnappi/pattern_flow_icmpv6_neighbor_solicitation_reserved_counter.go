package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NeighborSolicitationReservedCounter *****
type patternFlowIcmpv6NeighborSolicitationReservedCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter
	marshaller   marshalPatternFlowIcmpv6NeighborSolicitationReservedCounter
	unMarshaller unMarshalPatternFlowIcmpv6NeighborSolicitationReservedCounter
}

func NewPatternFlowIcmpv6NeighborSolicitationReservedCounter() PatternFlowIcmpv6NeighborSolicitationReservedCounter {
	obj := patternFlowIcmpv6NeighborSolicitationReservedCounter{obj: &otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) msg() *otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) setMsg(msg *otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter) PatternFlowIcmpv6NeighborSolicitationReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NeighborSolicitationReservedCounter struct {
	obj *patternFlowIcmpv6NeighborSolicitationReservedCounter
}

type marshalPatternFlowIcmpv6NeighborSolicitationReservedCounter interface {
	// ToProto marshals PatternFlowIcmpv6NeighborSolicitationReservedCounter to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter
	ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NeighborSolicitationReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NeighborSolicitationReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NeighborSolicitationReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NeighborSolicitationReservedCounter struct {
	obj *patternFlowIcmpv6NeighborSolicitationReservedCounter
}

type unMarshalPatternFlowIcmpv6NeighborSolicitationReservedCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NeighborSolicitationReservedCounter from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter
	FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter) (PatternFlowIcmpv6NeighborSolicitationReservedCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NeighborSolicitationReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NeighborSolicitationReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NeighborSolicitationReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) Marshal() marshalPatternFlowIcmpv6NeighborSolicitationReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NeighborSolicitationReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NeighborSolicitationReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationReservedCounter) ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationReservedCounter) FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter) (PatternFlowIcmpv6NeighborSolicitationReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) Clone() (PatternFlowIcmpv6NeighborSolicitationReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NeighborSolicitationReservedCounter()
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

// PatternFlowIcmpv6NeighborSolicitationReservedCounter is integer counter pattern
type PatternFlowIcmpv6NeighborSolicitationReservedCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NeighborSolicitationReservedCounter to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter
	// setMsg unmarshals PatternFlowIcmpv6NeighborSolicitationReservedCounter from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NeighborSolicitationReservedCounter) PatternFlowIcmpv6NeighborSolicitationReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NeighborSolicitationReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationReservedCounter
	// validate validates PatternFlowIcmpv6NeighborSolicitationReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NeighborSolicitationReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NeighborSolicitationReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationReservedCounter
	SetStart(value uint32) PatternFlowIcmpv6NeighborSolicitationReservedCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NeighborSolicitationReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NeighborSolicitationReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationReservedCounter
	SetStep(value uint32) PatternFlowIcmpv6NeighborSolicitationReservedCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NeighborSolicitationReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NeighborSolicitationReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationReservedCounter
	SetCount(value uint32) PatternFlowIcmpv6NeighborSolicitationReservedCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NeighborSolicitationReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationReservedCounter object
func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) SetStart(value uint32) PatternFlowIcmpv6NeighborSolicitationReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationReservedCounter object
func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) SetStep(value uint32) PatternFlowIcmpv6NeighborSolicitationReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationReservedCounter object
func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) SetCount(value uint32) PatternFlowIcmpv6NeighborSolicitationReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowIcmpv6NeighborSolicitationReservedCounter) setDefault() {
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
