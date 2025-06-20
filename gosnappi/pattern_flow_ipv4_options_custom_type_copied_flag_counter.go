package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter *****
type patternFlowIpv4OptionsCustomTypeCopiedFlagCounter struct {
	validation
	obj          *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	marshaller   marshalPatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	unMarshaller unMarshalPatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
}

func NewPatternFlowIpv4OptionsCustomTypeCopiedFlagCounter() PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter {
	obj := patternFlowIpv4OptionsCustomTypeCopiedFlagCounter{obj: &otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) msg() *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) setMsg(msg *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsCustomTypeCopiedFlagCounter struct {
	obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter
}

type marshalPatternFlowIpv4OptionsCustomTypeCopiedFlagCounter interface {
	// ToProto marshals PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter to protobuf object *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	ToProto() (*otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter, error)
	// ToPbText marshals PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsCustomTypeCopiedFlagCounter struct {
	obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter
}

type unMarshalPatternFlowIpv4OptionsCustomTypeCopiedFlagCounter interface {
	// FromProto unmarshals PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter from protobuf object *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	FromProto(msg *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) (PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) Marshal() marshalPatternFlowIpv4OptionsCustomTypeCopiedFlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsCustomTypeCopiedFlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) Unmarshal() unMarshalPatternFlowIpv4OptionsCustomTypeCopiedFlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsCustomTypeCopiedFlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) ToProto() (*otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) FromProto(msg *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) (PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) Clone() (PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsCustomTypeCopiedFlagCounter()
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

// PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter is integer counter pattern
type PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter to protobuf object *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	// setMsg unmarshals PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter from protobuf object *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	// validate validates PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	SetStart(value uint32) PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	// HasStart checks if Start has been set in PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	SetStep(value uint32) PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	// HasStep checks if Step has been set in PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	SetCount(value uint32) PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	// HasCount checks if Count has been set in PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter object
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) SetStart(value uint32) PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter object
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) SetStep(value uint32) PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter object
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) SetCount(value uint32) PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlagCounter) setDefault() {
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
