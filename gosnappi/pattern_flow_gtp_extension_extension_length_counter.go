package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpExtensionExtensionLengthCounter *****
type patternFlowGtpExtensionExtensionLengthCounter struct {
	validation
	obj          *otg.PatternFlowGtpExtensionExtensionLengthCounter
	marshaller   marshalPatternFlowGtpExtensionExtensionLengthCounter
	unMarshaller unMarshalPatternFlowGtpExtensionExtensionLengthCounter
}

func NewPatternFlowGtpExtensionExtensionLengthCounter() PatternFlowGtpExtensionExtensionLengthCounter {
	obj := patternFlowGtpExtensionExtensionLengthCounter{obj: &otg.PatternFlowGtpExtensionExtensionLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) msg() *otg.PatternFlowGtpExtensionExtensionLengthCounter {
	return obj.obj
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) setMsg(msg *otg.PatternFlowGtpExtensionExtensionLengthCounter) PatternFlowGtpExtensionExtensionLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpExtensionExtensionLengthCounter struct {
	obj *patternFlowGtpExtensionExtensionLengthCounter
}

type marshalPatternFlowGtpExtensionExtensionLengthCounter interface {
	// ToProto marshals PatternFlowGtpExtensionExtensionLengthCounter to protobuf object *otg.PatternFlowGtpExtensionExtensionLengthCounter
	ToProto() (*otg.PatternFlowGtpExtensionExtensionLengthCounter, error)
	// ToPbText marshals PatternFlowGtpExtensionExtensionLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpExtensionExtensionLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpExtensionExtensionLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpExtensionExtensionLengthCounter struct {
	obj *patternFlowGtpExtensionExtensionLengthCounter
}

type unMarshalPatternFlowGtpExtensionExtensionLengthCounter interface {
	// FromProto unmarshals PatternFlowGtpExtensionExtensionLengthCounter from protobuf object *otg.PatternFlowGtpExtensionExtensionLengthCounter
	FromProto(msg *otg.PatternFlowGtpExtensionExtensionLengthCounter) (PatternFlowGtpExtensionExtensionLengthCounter, error)
	// FromPbText unmarshals PatternFlowGtpExtensionExtensionLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpExtensionExtensionLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpExtensionExtensionLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) Marshal() marshalPatternFlowGtpExtensionExtensionLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpExtensionExtensionLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) Unmarshal() unMarshalPatternFlowGtpExtensionExtensionLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpExtensionExtensionLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpExtensionExtensionLengthCounter) ToProto() (*otg.PatternFlowGtpExtensionExtensionLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpExtensionExtensionLengthCounter) FromProto(msg *otg.PatternFlowGtpExtensionExtensionLengthCounter) (PatternFlowGtpExtensionExtensionLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpExtensionExtensionLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionExtensionLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpExtensionExtensionLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionExtensionLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpExtensionExtensionLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionExtensionLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpExtensionExtensionLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) Clone() (PatternFlowGtpExtensionExtensionLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpExtensionExtensionLengthCounter()
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

// PatternFlowGtpExtensionExtensionLengthCounter is integer counter pattern
type PatternFlowGtpExtensionExtensionLengthCounter interface {
	Validation
	// msg marshals PatternFlowGtpExtensionExtensionLengthCounter to protobuf object *otg.PatternFlowGtpExtensionExtensionLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpExtensionExtensionLengthCounter
	// setMsg unmarshals PatternFlowGtpExtensionExtensionLengthCounter from protobuf object *otg.PatternFlowGtpExtensionExtensionLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpExtensionExtensionLengthCounter) PatternFlowGtpExtensionExtensionLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpExtensionExtensionLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpExtensionExtensionLengthCounter
	// validate validates PatternFlowGtpExtensionExtensionLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpExtensionExtensionLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpExtensionExtensionLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpExtensionExtensionLengthCounter
	SetStart(value uint32) PatternFlowGtpExtensionExtensionLengthCounter
	// HasStart checks if Start has been set in PatternFlowGtpExtensionExtensionLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpExtensionExtensionLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpExtensionExtensionLengthCounter
	SetStep(value uint32) PatternFlowGtpExtensionExtensionLengthCounter
	// HasStep checks if Step has been set in PatternFlowGtpExtensionExtensionLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpExtensionExtensionLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpExtensionExtensionLengthCounter
	SetCount(value uint32) PatternFlowGtpExtensionExtensionLengthCounter
	// HasCount checks if Count has been set in PatternFlowGtpExtensionExtensionLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpExtensionExtensionLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpExtensionExtensionLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpExtensionExtensionLengthCounter object
func (obj *patternFlowGtpExtensionExtensionLengthCounter) SetStart(value uint32) PatternFlowGtpExtensionExtensionLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpExtensionExtensionLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpExtensionExtensionLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpExtensionExtensionLengthCounter object
func (obj *patternFlowGtpExtensionExtensionLengthCounter) SetStep(value uint32) PatternFlowGtpExtensionExtensionLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpExtensionExtensionLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpExtensionExtensionLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpExtensionExtensionLengthCounter object
func (obj *patternFlowGtpExtensionExtensionLengthCounter) SetCount(value uint32) PatternFlowGtpExtensionExtensionLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionExtensionLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionExtensionLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionExtensionLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpExtensionExtensionLengthCounter) setDefault() {
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
