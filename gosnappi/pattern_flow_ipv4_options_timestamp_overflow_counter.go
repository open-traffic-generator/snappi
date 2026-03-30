package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampOverflowCounter *****
type patternFlowIpv4OptionsTimestampOverflowCounter struct {
	validation
	obj          *otg.PatternFlowIpv4OptionsTimestampOverflowCounter
	marshaller   marshalPatternFlowIpv4OptionsTimestampOverflowCounter
	unMarshaller unMarshalPatternFlowIpv4OptionsTimestampOverflowCounter
}

func NewPatternFlowIpv4OptionsTimestampOverflowCounter() PatternFlowIpv4OptionsTimestampOverflowCounter {
	obj := patternFlowIpv4OptionsTimestampOverflowCounter{obj: &otg.PatternFlowIpv4OptionsTimestampOverflowCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) msg() *otg.PatternFlowIpv4OptionsTimestampOverflowCounter {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampOverflowCounter) PatternFlowIpv4OptionsTimestampOverflowCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampOverflowCounter struct {
	obj *patternFlowIpv4OptionsTimestampOverflowCounter
}

type marshalPatternFlowIpv4OptionsTimestampOverflowCounter interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampOverflowCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampOverflowCounter
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampOverflowCounter, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampOverflowCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampOverflowCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampOverflowCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampOverflowCounter struct {
	obj *patternFlowIpv4OptionsTimestampOverflowCounter
}

type unMarshalPatternFlowIpv4OptionsTimestampOverflowCounter interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampOverflowCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampOverflowCounter
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampOverflowCounter) (PatternFlowIpv4OptionsTimestampOverflowCounter, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampOverflowCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampOverflowCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampOverflowCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) Marshal() marshalPatternFlowIpv4OptionsTimestampOverflowCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampOverflowCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampOverflowCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampOverflowCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampOverflowCounter) ToProto() (*otg.PatternFlowIpv4OptionsTimestampOverflowCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampOverflowCounter) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampOverflowCounter) (PatternFlowIpv4OptionsTimestampOverflowCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampOverflowCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampOverflowCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampOverflowCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampOverflowCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampOverflowCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampOverflowCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) Clone() (PatternFlowIpv4OptionsTimestampOverflowCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampOverflowCounter()
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

// PatternFlowIpv4OptionsTimestampOverflowCounter is integer counter pattern
type PatternFlowIpv4OptionsTimestampOverflowCounter interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampOverflowCounter to protobuf object *otg.PatternFlowIpv4OptionsTimestampOverflowCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampOverflowCounter
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampOverflowCounter from protobuf object *otg.PatternFlowIpv4OptionsTimestampOverflowCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampOverflowCounter) PatternFlowIpv4OptionsTimestampOverflowCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampOverflowCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampOverflowCounter
	// validate validates PatternFlowIpv4OptionsTimestampOverflowCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampOverflowCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4OptionsTimestampOverflowCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampOverflowCounter
	SetStart(value uint32) PatternFlowIpv4OptionsTimestampOverflowCounter
	// HasStart checks if Start has been set in PatternFlowIpv4OptionsTimestampOverflowCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4OptionsTimestampOverflowCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampOverflowCounter
	SetStep(value uint32) PatternFlowIpv4OptionsTimestampOverflowCounter
	// HasStep checks if Step has been set in PatternFlowIpv4OptionsTimestampOverflowCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4OptionsTimestampOverflowCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampOverflowCounter
	SetCount(value uint32) PatternFlowIpv4OptionsTimestampOverflowCounter
	// HasCount checks if Count has been set in PatternFlowIpv4OptionsTimestampOverflowCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4OptionsTimestampOverflowCounter object
func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) SetStart(value uint32) PatternFlowIpv4OptionsTimestampOverflowCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4OptionsTimestampOverflowCounter object
func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) SetStep(value uint32) PatternFlowIpv4OptionsTimestampOverflowCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4OptionsTimestampOverflowCounter object
func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) SetCount(value uint32) PatternFlowIpv4OptionsTimestampOverflowCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsTimestampOverflowCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsTimestampOverflowCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsTimestampOverflowCounter.Count <= 16 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4OptionsTimestampOverflowCounter) setDefault() {
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
