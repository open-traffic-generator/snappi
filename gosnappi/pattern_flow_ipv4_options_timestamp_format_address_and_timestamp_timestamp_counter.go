package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter *****
type patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter struct {
	validation
	obj          *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	marshaller   marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	unMarshaller unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
}

func NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter {
	obj := patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter{obj: &otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
}

type marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
}

type unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter()
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

// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter is integer counter pattern
type PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	// validate validates PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	SetStart(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	// HasStart checks if Start has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	SetStep(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	// HasStep checks if Step has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	SetCount(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	// HasCount checks if Count has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) SetStart(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) SetStep(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) SetCount(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestampCounter) setDefault() {
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
