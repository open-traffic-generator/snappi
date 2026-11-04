package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter *****
type patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter struct {
	validation
	obj          *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	marshaller   marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	unMarshaller unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
}

func NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter {
	obj := patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter{obj: &otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
}

type marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
}

type unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter()
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

// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter is integer counter pattern
type PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	// validate validates PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	SetStart(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	// HasStart checks if Start has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	SetStep(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	// HasStep checks if Step has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	SetCount(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	// HasCount checks if Count has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) SetStart(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) SetStep(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) SetCount(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestampCounter) setDefault() {
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
