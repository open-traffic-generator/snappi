package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter *****
type patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter struct {
	validation
	obj          *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	marshaller   marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	unMarshaller unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
}

func NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter {
	obj := patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter{obj: &otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
}

type marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter struct {
	obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
}

type unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter()
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

// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter is ipv4 counter pattern
type PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	// validate validates PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	SetStart(value string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	// HasStart checks if Start has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	SetStep(value string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	// HasStep checks if Step has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	SetCount(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	// HasCount checks if Count has been set in PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) SetStart(value string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) SetStep(value string) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter object
func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) SetCount(value uint32) PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddressCounter) setDefault() {
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
