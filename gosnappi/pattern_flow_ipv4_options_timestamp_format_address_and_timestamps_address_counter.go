package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter *****
type patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter struct {
	validation
	obj          *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	marshaller   marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	unMarshaller unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
}

func NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter {
	obj := patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter{obj: &otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
}

type marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
}

type unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter()
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

// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter is ipv4 counter pattern
type PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	// validate validates PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	SetStart(value string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	// HasStart checks if Start has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	SetStep(value string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	// HasStep checks if Step has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	SetCount(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	// HasCount checks if Count has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) SetStart(value string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) SetStep(value string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) SetCount(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddressCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("0.0.0.0")
	}
	if obj.obj.Step == nil {
		obj.SetStep("0.0.0.1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
