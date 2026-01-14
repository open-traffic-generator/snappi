package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1SFlagCounter *****
type patternFlowGtpv1SFlagCounter struct {
	validation
	obj          *otg.PatternFlowGtpv1SFlagCounter
	marshaller   marshalPatternFlowGtpv1SFlagCounter
	unMarshaller unMarshalPatternFlowGtpv1SFlagCounter
}

func NewPatternFlowGtpv1SFlagCounter() PatternFlowGtpv1SFlagCounter {
	obj := patternFlowGtpv1SFlagCounter{obj: &otg.PatternFlowGtpv1SFlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1SFlagCounter) msg() *otg.PatternFlowGtpv1SFlagCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1SFlagCounter) setMsg(msg *otg.PatternFlowGtpv1SFlagCounter) PatternFlowGtpv1SFlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1SFlagCounter struct {
	obj *patternFlowGtpv1SFlagCounter
}

type marshalPatternFlowGtpv1SFlagCounter interface {
	// ToProto marshals PatternFlowGtpv1SFlagCounter to protobuf object *otg.PatternFlowGtpv1SFlagCounter
	ToProto() (*otg.PatternFlowGtpv1SFlagCounter, error)
	// ToPbText marshals PatternFlowGtpv1SFlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1SFlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1SFlagCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1SFlagCounter struct {
	obj *patternFlowGtpv1SFlagCounter
}

type unMarshalPatternFlowGtpv1SFlagCounter interface {
	// FromProto unmarshals PatternFlowGtpv1SFlagCounter from protobuf object *otg.PatternFlowGtpv1SFlagCounter
	FromProto(msg *otg.PatternFlowGtpv1SFlagCounter) (PatternFlowGtpv1SFlagCounter, error)
	// FromPbText unmarshals PatternFlowGtpv1SFlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1SFlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1SFlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1SFlagCounter) Marshal() marshalPatternFlowGtpv1SFlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1SFlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1SFlagCounter) Unmarshal() unMarshalPatternFlowGtpv1SFlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1SFlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1SFlagCounter) ToProto() (*otg.PatternFlowGtpv1SFlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1SFlagCounter) FromProto(msg *otg.PatternFlowGtpv1SFlagCounter) (PatternFlowGtpv1SFlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1SFlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SFlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1SFlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SFlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1SFlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SFlagCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv1SFlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1SFlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1SFlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1SFlagCounter) Clone() (PatternFlowGtpv1SFlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1SFlagCounter()
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

// PatternFlowGtpv1SFlagCounter is integer counter pattern
type PatternFlowGtpv1SFlagCounter interface {
	Validation
	// msg marshals PatternFlowGtpv1SFlagCounter to protobuf object *otg.PatternFlowGtpv1SFlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1SFlagCounter
	// setMsg unmarshals PatternFlowGtpv1SFlagCounter from protobuf object *otg.PatternFlowGtpv1SFlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1SFlagCounter) PatternFlowGtpv1SFlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1SFlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1SFlagCounter
	// validate validates PatternFlowGtpv1SFlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1SFlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv1SFlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv1SFlagCounter
	SetStart(value uint32) PatternFlowGtpv1SFlagCounter
	// HasStart checks if Start has been set in PatternFlowGtpv1SFlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv1SFlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv1SFlagCounter
	SetStep(value uint32) PatternFlowGtpv1SFlagCounter
	// HasStep checks if Step has been set in PatternFlowGtpv1SFlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv1SFlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv1SFlagCounter
	SetCount(value uint32) PatternFlowGtpv1SFlagCounter
	// HasCount checks if Count has been set in PatternFlowGtpv1SFlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1SFlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1SFlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv1SFlagCounter object
func (obj *patternFlowGtpv1SFlagCounter) SetStart(value uint32) PatternFlowGtpv1SFlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1SFlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1SFlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv1SFlagCounter object
func (obj *patternFlowGtpv1SFlagCounter) SetStep(value uint32) PatternFlowGtpv1SFlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1SFlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1SFlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv1SFlagCounter object
func (obj *patternFlowGtpv1SFlagCounter) SetCount(value uint32) PatternFlowGtpv1SFlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv1SFlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SFlagCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv1SFlagCounter) setDefault() {
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
