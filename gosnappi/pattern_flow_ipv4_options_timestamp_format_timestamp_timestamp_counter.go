package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter *****
type patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter struct {
	validation
	obj          *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	marshaller   marshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	unMarshaller unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
}

func NewPatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter() PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter {
	obj := patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter{obj: &otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) msg() *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter struct {
	obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
}

type marshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter struct {
	obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
}

type unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) (PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) Marshal() marshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) (PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) Clone() (PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter()
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

// PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter is integer counter pattern
type PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	// validate validates PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	SetStart(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	// HasStart checks if Start has been set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	SetStep(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	// HasStep checks if Step has been set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	SetCount(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	// HasCount checks if Count has been set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) SetStart(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) SetStep(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) SetCount(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) setDefault() {
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
