package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1EFlagCounter *****
type patternFlowGtpv1EFlagCounter struct {
	validation
	obj          *otg.PatternFlowGtpv1EFlagCounter
	marshaller   marshalPatternFlowGtpv1EFlagCounter
	unMarshaller unMarshalPatternFlowGtpv1EFlagCounter
}

func NewPatternFlowGtpv1EFlagCounter() PatternFlowGtpv1EFlagCounter {
	obj := patternFlowGtpv1EFlagCounter{obj: &otg.PatternFlowGtpv1EFlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1EFlagCounter) msg() *otg.PatternFlowGtpv1EFlagCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1EFlagCounter) setMsg(msg *otg.PatternFlowGtpv1EFlagCounter) PatternFlowGtpv1EFlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1EFlagCounter struct {
	obj *patternFlowGtpv1EFlagCounter
}

type marshalPatternFlowGtpv1EFlagCounter interface {
	// ToProto marshals PatternFlowGtpv1EFlagCounter to protobuf object *otg.PatternFlowGtpv1EFlagCounter
	ToProto() (*otg.PatternFlowGtpv1EFlagCounter, error)
	// ToPbText marshals PatternFlowGtpv1EFlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1EFlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1EFlagCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1EFlagCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1EFlagCounter struct {
	obj *patternFlowGtpv1EFlagCounter
}

type unMarshalPatternFlowGtpv1EFlagCounter interface {
	// FromProto unmarshals PatternFlowGtpv1EFlagCounter from protobuf object *otg.PatternFlowGtpv1EFlagCounter
	FromProto(msg *otg.PatternFlowGtpv1EFlagCounter) (PatternFlowGtpv1EFlagCounter, error)
	// FromPbText unmarshals PatternFlowGtpv1EFlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1EFlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1EFlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1EFlagCounter) Marshal() marshalPatternFlowGtpv1EFlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1EFlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1EFlagCounter) Unmarshal() unMarshalPatternFlowGtpv1EFlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1EFlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1EFlagCounter) ToProto() (*otg.PatternFlowGtpv1EFlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1EFlagCounter) FromProto(msg *otg.PatternFlowGtpv1EFlagCounter) (PatternFlowGtpv1EFlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1EFlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1EFlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1EFlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1EFlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1EFlagCounter) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowGtpv1EFlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1EFlagCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv1EFlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1EFlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1EFlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1EFlagCounter) Clone() (PatternFlowGtpv1EFlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1EFlagCounter()
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

// PatternFlowGtpv1EFlagCounter is integer counter pattern
type PatternFlowGtpv1EFlagCounter interface {
	Validation
	// msg marshals PatternFlowGtpv1EFlagCounter to protobuf object *otg.PatternFlowGtpv1EFlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1EFlagCounter
	// setMsg unmarshals PatternFlowGtpv1EFlagCounter from protobuf object *otg.PatternFlowGtpv1EFlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1EFlagCounter) PatternFlowGtpv1EFlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1EFlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1EFlagCounter
	// validate validates PatternFlowGtpv1EFlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1EFlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv1EFlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv1EFlagCounter
	SetStart(value uint32) PatternFlowGtpv1EFlagCounter
	// HasStart checks if Start has been set in PatternFlowGtpv1EFlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv1EFlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv1EFlagCounter
	SetStep(value uint32) PatternFlowGtpv1EFlagCounter
	// HasStep checks if Step has been set in PatternFlowGtpv1EFlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv1EFlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv1EFlagCounter
	SetCount(value uint32) PatternFlowGtpv1EFlagCounter
	// HasCount checks if Count has been set in PatternFlowGtpv1EFlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1EFlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1EFlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv1EFlagCounter object
func (obj *patternFlowGtpv1EFlagCounter) SetStart(value uint32) PatternFlowGtpv1EFlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1EFlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1EFlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv1EFlagCounter object
func (obj *patternFlowGtpv1EFlagCounter) SetStep(value uint32) PatternFlowGtpv1EFlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1EFlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1EFlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv1EFlagCounter object
func (obj *patternFlowGtpv1EFlagCounter) SetCount(value uint32) PatternFlowGtpv1EFlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv1EFlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1EFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1EFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1EFlagCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv1EFlagCounter) setDefault() {
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
