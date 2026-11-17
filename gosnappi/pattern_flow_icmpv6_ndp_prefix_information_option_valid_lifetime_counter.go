package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter *****
type patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	marshaller   marshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter() PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter()
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

// PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter is integer counter pattern
type PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(4294967295)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
