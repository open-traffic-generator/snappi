package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter *****
type patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter struct {
	validation
	obj          *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	marshaller   marshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	unMarshaller unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
}

func NewPatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter() PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter {
	obj := patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter{obj: &otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) msg() *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter struct {
	obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
}

type marshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter struct {
	obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
}

type unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) (PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) Marshal() marshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) (PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) Clone() (PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter()
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

// PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter is integer counter pattern
type PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	// validate validates PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	SetStart(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	// HasStart checks if Start has been set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	SetStep(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	// HasStep checks if Step has been set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	SetCount(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	// HasCount checks if Count has been set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) SetStart(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) SetStep(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) SetCount(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) setDefault() {
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
