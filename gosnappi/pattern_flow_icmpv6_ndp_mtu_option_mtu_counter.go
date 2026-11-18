package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpMtuOptionMtuCounter *****
type patternFlowIcmpv6NdpMtuOptionMtuCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter
	marshaller   marshalPatternFlowIcmpv6NdpMtuOptionMtuCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpMtuOptionMtuCounter
}

func NewPatternFlowIcmpv6NdpMtuOptionMtuCounter() PatternFlowIcmpv6NdpMtuOptionMtuCounter {
	obj := patternFlowIcmpv6NdpMtuOptionMtuCounter{obj: &otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) msg() *otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter) PatternFlowIcmpv6NdpMtuOptionMtuCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpMtuOptionMtuCounter struct {
	obj *patternFlowIcmpv6NdpMtuOptionMtuCounter
}

type marshalPatternFlowIcmpv6NdpMtuOptionMtuCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpMtuOptionMtuCounter to protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpMtuOptionMtuCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpMtuOptionMtuCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpMtuOptionMtuCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpMtuOptionMtuCounter struct {
	obj *patternFlowIcmpv6NdpMtuOptionMtuCounter
}

type unMarshalPatternFlowIcmpv6NdpMtuOptionMtuCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpMtuOptionMtuCounter from protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter) (PatternFlowIcmpv6NdpMtuOptionMtuCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpMtuOptionMtuCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpMtuOptionMtuCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpMtuOptionMtuCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) Marshal() marshalPatternFlowIcmpv6NdpMtuOptionMtuCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpMtuOptionMtuCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpMtuOptionMtuCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpMtuOptionMtuCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpMtuOptionMtuCounter) ToProto() (*otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionMtuCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter) (PatternFlowIcmpv6NdpMtuOptionMtuCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpMtuOptionMtuCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionMtuCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpMtuOptionMtuCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionMtuCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpMtuOptionMtuCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionMtuCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) Clone() (PatternFlowIcmpv6NdpMtuOptionMtuCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpMtuOptionMtuCounter()
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

// PatternFlowIcmpv6NdpMtuOptionMtuCounter is integer counter pattern
type PatternFlowIcmpv6NdpMtuOptionMtuCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpMtuOptionMtuCounter to protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpMtuOptionMtuCounter from protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpMtuOptionMtuCounter) PatternFlowIcmpv6NdpMtuOptionMtuCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpMtuOptionMtuCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpMtuOptionMtuCounter
	// validate validates PatternFlowIcmpv6NdpMtuOptionMtuCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpMtuOptionMtuCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpMtuOptionMtuCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionMtuCounter
	SetStart(value uint32) PatternFlowIcmpv6NdpMtuOptionMtuCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpMtuOptionMtuCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpMtuOptionMtuCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionMtuCounter
	SetStep(value uint32) PatternFlowIcmpv6NdpMtuOptionMtuCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpMtuOptionMtuCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpMtuOptionMtuCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionMtuCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpMtuOptionMtuCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpMtuOptionMtuCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpMtuOptionMtuCounter object
func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) SetStart(value uint32) PatternFlowIcmpv6NdpMtuOptionMtuCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpMtuOptionMtuCounter object
func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) SetStep(value uint32) PatternFlowIcmpv6NdpMtuOptionMtuCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpMtuOptionMtuCounter object
func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) SetCount(value uint32) PatternFlowIcmpv6NdpMtuOptionMtuCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowIcmpv6NdpMtuOptionMtuCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1500)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
