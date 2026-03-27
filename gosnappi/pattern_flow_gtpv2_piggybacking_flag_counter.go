package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2PiggybackingFlagCounter *****
type patternFlowGtpv2PiggybackingFlagCounter struct {
	validation
	obj          *otg.PatternFlowGtpv2PiggybackingFlagCounter
	marshaller   marshalPatternFlowGtpv2PiggybackingFlagCounter
	unMarshaller unMarshalPatternFlowGtpv2PiggybackingFlagCounter
}

func NewPatternFlowGtpv2PiggybackingFlagCounter() PatternFlowGtpv2PiggybackingFlagCounter {
	obj := patternFlowGtpv2PiggybackingFlagCounter{obj: &otg.PatternFlowGtpv2PiggybackingFlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) msg() *otg.PatternFlowGtpv2PiggybackingFlagCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) setMsg(msg *otg.PatternFlowGtpv2PiggybackingFlagCounter) PatternFlowGtpv2PiggybackingFlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2PiggybackingFlagCounter struct {
	obj *patternFlowGtpv2PiggybackingFlagCounter
}

type marshalPatternFlowGtpv2PiggybackingFlagCounter interface {
	// ToProto marshals PatternFlowGtpv2PiggybackingFlagCounter to protobuf object *otg.PatternFlowGtpv2PiggybackingFlagCounter
	ToProto() (*otg.PatternFlowGtpv2PiggybackingFlagCounter, error)
	// ToPbText marshals PatternFlowGtpv2PiggybackingFlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2PiggybackingFlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2PiggybackingFlagCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2PiggybackingFlagCounter struct {
	obj *patternFlowGtpv2PiggybackingFlagCounter
}

type unMarshalPatternFlowGtpv2PiggybackingFlagCounter interface {
	// FromProto unmarshals PatternFlowGtpv2PiggybackingFlagCounter from protobuf object *otg.PatternFlowGtpv2PiggybackingFlagCounter
	FromProto(msg *otg.PatternFlowGtpv2PiggybackingFlagCounter) (PatternFlowGtpv2PiggybackingFlagCounter, error)
	// FromPbText unmarshals PatternFlowGtpv2PiggybackingFlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2PiggybackingFlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2PiggybackingFlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) Marshal() marshalPatternFlowGtpv2PiggybackingFlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2PiggybackingFlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) Unmarshal() unMarshalPatternFlowGtpv2PiggybackingFlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2PiggybackingFlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2PiggybackingFlagCounter) ToProto() (*otg.PatternFlowGtpv2PiggybackingFlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2PiggybackingFlagCounter) FromProto(msg *otg.PatternFlowGtpv2PiggybackingFlagCounter) (PatternFlowGtpv2PiggybackingFlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2PiggybackingFlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2PiggybackingFlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2PiggybackingFlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2PiggybackingFlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2PiggybackingFlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2PiggybackingFlagCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv2PiggybackingFlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) Clone() (PatternFlowGtpv2PiggybackingFlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2PiggybackingFlagCounter()
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

// PatternFlowGtpv2PiggybackingFlagCounter is integer counter pattern
type PatternFlowGtpv2PiggybackingFlagCounter interface {
	Validation
	// msg marshals PatternFlowGtpv2PiggybackingFlagCounter to protobuf object *otg.PatternFlowGtpv2PiggybackingFlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2PiggybackingFlagCounter
	// setMsg unmarshals PatternFlowGtpv2PiggybackingFlagCounter from protobuf object *otg.PatternFlowGtpv2PiggybackingFlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2PiggybackingFlagCounter) PatternFlowGtpv2PiggybackingFlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2PiggybackingFlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2PiggybackingFlagCounter
	// validate validates PatternFlowGtpv2PiggybackingFlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2PiggybackingFlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv2PiggybackingFlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv2PiggybackingFlagCounter
	SetStart(value uint32) PatternFlowGtpv2PiggybackingFlagCounter
	// HasStart checks if Start has been set in PatternFlowGtpv2PiggybackingFlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv2PiggybackingFlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv2PiggybackingFlagCounter
	SetStep(value uint32) PatternFlowGtpv2PiggybackingFlagCounter
	// HasStep checks if Step has been set in PatternFlowGtpv2PiggybackingFlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv2PiggybackingFlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv2PiggybackingFlagCounter
	SetCount(value uint32) PatternFlowGtpv2PiggybackingFlagCounter
	// HasCount checks if Count has been set in PatternFlowGtpv2PiggybackingFlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2PiggybackingFlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2PiggybackingFlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv2PiggybackingFlagCounter object
func (obj *patternFlowGtpv2PiggybackingFlagCounter) SetStart(value uint32) PatternFlowGtpv2PiggybackingFlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2PiggybackingFlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2PiggybackingFlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv2PiggybackingFlagCounter object
func (obj *patternFlowGtpv2PiggybackingFlagCounter) SetStep(value uint32) PatternFlowGtpv2PiggybackingFlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2PiggybackingFlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2PiggybackingFlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv2PiggybackingFlagCounter object
func (obj *patternFlowGtpv2PiggybackingFlagCounter) SetCount(value uint32) PatternFlowGtpv2PiggybackingFlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2PiggybackingFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2PiggybackingFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2PiggybackingFlagCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv2PiggybackingFlagCounter) setDefault() {
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
