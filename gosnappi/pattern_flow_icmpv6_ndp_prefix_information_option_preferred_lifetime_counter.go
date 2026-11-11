package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter *****
type patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	marshaller   marshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter() PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter()
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

// PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter is integer counter pattern
type PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetimeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(604800)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
