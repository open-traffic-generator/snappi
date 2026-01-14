package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2MessageTypeCounter *****
type patternFlowGtpv2MessageTypeCounter struct {
	validation
	obj          *otg.PatternFlowGtpv2MessageTypeCounter
	marshaller   marshalPatternFlowGtpv2MessageTypeCounter
	unMarshaller unMarshalPatternFlowGtpv2MessageTypeCounter
}

func NewPatternFlowGtpv2MessageTypeCounter() PatternFlowGtpv2MessageTypeCounter {
	obj := patternFlowGtpv2MessageTypeCounter{obj: &otg.PatternFlowGtpv2MessageTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2MessageTypeCounter) msg() *otg.PatternFlowGtpv2MessageTypeCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2MessageTypeCounter) setMsg(msg *otg.PatternFlowGtpv2MessageTypeCounter) PatternFlowGtpv2MessageTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2MessageTypeCounter struct {
	obj *patternFlowGtpv2MessageTypeCounter
}

type marshalPatternFlowGtpv2MessageTypeCounter interface {
	// ToProto marshals PatternFlowGtpv2MessageTypeCounter to protobuf object *otg.PatternFlowGtpv2MessageTypeCounter
	ToProto() (*otg.PatternFlowGtpv2MessageTypeCounter, error)
	// ToPbText marshals PatternFlowGtpv2MessageTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2MessageTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2MessageTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2MessageTypeCounter struct {
	obj *patternFlowGtpv2MessageTypeCounter
}

type unMarshalPatternFlowGtpv2MessageTypeCounter interface {
	// FromProto unmarshals PatternFlowGtpv2MessageTypeCounter from protobuf object *otg.PatternFlowGtpv2MessageTypeCounter
	FromProto(msg *otg.PatternFlowGtpv2MessageTypeCounter) (PatternFlowGtpv2MessageTypeCounter, error)
	// FromPbText unmarshals PatternFlowGtpv2MessageTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2MessageTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2MessageTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2MessageTypeCounter) Marshal() marshalPatternFlowGtpv2MessageTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2MessageTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2MessageTypeCounter) Unmarshal() unMarshalPatternFlowGtpv2MessageTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2MessageTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2MessageTypeCounter) ToProto() (*otg.PatternFlowGtpv2MessageTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2MessageTypeCounter) FromProto(msg *otg.PatternFlowGtpv2MessageTypeCounter) (PatternFlowGtpv2MessageTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2MessageTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2MessageTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2MessageTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv2MessageTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2MessageTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2MessageTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2MessageTypeCounter) Clone() (PatternFlowGtpv2MessageTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2MessageTypeCounter()
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

// PatternFlowGtpv2MessageTypeCounter is integer counter pattern
type PatternFlowGtpv2MessageTypeCounter interface {
	Validation
	// msg marshals PatternFlowGtpv2MessageTypeCounter to protobuf object *otg.PatternFlowGtpv2MessageTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2MessageTypeCounter
	// setMsg unmarshals PatternFlowGtpv2MessageTypeCounter from protobuf object *otg.PatternFlowGtpv2MessageTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2MessageTypeCounter) PatternFlowGtpv2MessageTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2MessageTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2MessageTypeCounter
	// validate validates PatternFlowGtpv2MessageTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2MessageTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv2MessageTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv2MessageTypeCounter
	SetStart(value uint32) PatternFlowGtpv2MessageTypeCounter
	// HasStart checks if Start has been set in PatternFlowGtpv2MessageTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv2MessageTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv2MessageTypeCounter
	SetStep(value uint32) PatternFlowGtpv2MessageTypeCounter
	// HasStep checks if Step has been set in PatternFlowGtpv2MessageTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv2MessageTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv2MessageTypeCounter
	SetCount(value uint32) PatternFlowGtpv2MessageTypeCounter
	// HasCount checks if Count has been set in PatternFlowGtpv2MessageTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2MessageTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2MessageTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv2MessageTypeCounter object
func (obj *patternFlowGtpv2MessageTypeCounter) SetStart(value uint32) PatternFlowGtpv2MessageTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2MessageTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2MessageTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv2MessageTypeCounter object
func (obj *patternFlowGtpv2MessageTypeCounter) SetStep(value uint32) PatternFlowGtpv2MessageTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2MessageTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2MessageTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv2MessageTypeCounter object
func (obj *patternFlowGtpv2MessageTypeCounter) SetCount(value uint32) PatternFlowGtpv2MessageTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv2MessageTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv2MessageTypeCounter) setDefault() {
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
