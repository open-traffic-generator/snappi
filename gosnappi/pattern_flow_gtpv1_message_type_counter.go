package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1MessageTypeCounter *****
type patternFlowGtpv1MessageTypeCounter struct {
	validation
	obj          *otg.PatternFlowGtpv1MessageTypeCounter
	marshaller   marshalPatternFlowGtpv1MessageTypeCounter
	unMarshaller unMarshalPatternFlowGtpv1MessageTypeCounter
}

func NewPatternFlowGtpv1MessageTypeCounter() PatternFlowGtpv1MessageTypeCounter {
	obj := patternFlowGtpv1MessageTypeCounter{obj: &otg.PatternFlowGtpv1MessageTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1MessageTypeCounter) msg() *otg.PatternFlowGtpv1MessageTypeCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1MessageTypeCounter) setMsg(msg *otg.PatternFlowGtpv1MessageTypeCounter) PatternFlowGtpv1MessageTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1MessageTypeCounter struct {
	obj *patternFlowGtpv1MessageTypeCounter
}

type marshalPatternFlowGtpv1MessageTypeCounter interface {
	// ToProto marshals PatternFlowGtpv1MessageTypeCounter to protobuf object *otg.PatternFlowGtpv1MessageTypeCounter
	ToProto() (*otg.PatternFlowGtpv1MessageTypeCounter, error)
	// ToPbText marshals PatternFlowGtpv1MessageTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1MessageTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1MessageTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1MessageTypeCounter struct {
	obj *patternFlowGtpv1MessageTypeCounter
}

type unMarshalPatternFlowGtpv1MessageTypeCounter interface {
	// FromProto unmarshals PatternFlowGtpv1MessageTypeCounter from protobuf object *otg.PatternFlowGtpv1MessageTypeCounter
	FromProto(msg *otg.PatternFlowGtpv1MessageTypeCounter) (PatternFlowGtpv1MessageTypeCounter, error)
	// FromPbText unmarshals PatternFlowGtpv1MessageTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1MessageTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1MessageTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1MessageTypeCounter) Marshal() marshalPatternFlowGtpv1MessageTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1MessageTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1MessageTypeCounter) Unmarshal() unMarshalPatternFlowGtpv1MessageTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1MessageTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1MessageTypeCounter) ToProto() (*otg.PatternFlowGtpv1MessageTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1MessageTypeCounter) FromProto(msg *otg.PatternFlowGtpv1MessageTypeCounter) (PatternFlowGtpv1MessageTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1MessageTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1MessageTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1MessageTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv1MessageTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1MessageTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1MessageTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1MessageTypeCounter) Clone() (PatternFlowGtpv1MessageTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1MessageTypeCounter()
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

// PatternFlowGtpv1MessageTypeCounter is integer counter pattern
type PatternFlowGtpv1MessageTypeCounter interface {
	Validation
	// msg marshals PatternFlowGtpv1MessageTypeCounter to protobuf object *otg.PatternFlowGtpv1MessageTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1MessageTypeCounter
	// setMsg unmarshals PatternFlowGtpv1MessageTypeCounter from protobuf object *otg.PatternFlowGtpv1MessageTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1MessageTypeCounter) PatternFlowGtpv1MessageTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1MessageTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1MessageTypeCounter
	// validate validates PatternFlowGtpv1MessageTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1MessageTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv1MessageTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv1MessageTypeCounter
	SetStart(value uint32) PatternFlowGtpv1MessageTypeCounter
	// HasStart checks if Start has been set in PatternFlowGtpv1MessageTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv1MessageTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv1MessageTypeCounter
	SetStep(value uint32) PatternFlowGtpv1MessageTypeCounter
	// HasStep checks if Step has been set in PatternFlowGtpv1MessageTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv1MessageTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv1MessageTypeCounter
	SetCount(value uint32) PatternFlowGtpv1MessageTypeCounter
	// HasCount checks if Count has been set in PatternFlowGtpv1MessageTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1MessageTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1MessageTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv1MessageTypeCounter object
func (obj *patternFlowGtpv1MessageTypeCounter) SetStart(value uint32) PatternFlowGtpv1MessageTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1MessageTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1MessageTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv1MessageTypeCounter object
func (obj *patternFlowGtpv1MessageTypeCounter) SetStep(value uint32) PatternFlowGtpv1MessageTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1MessageTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1MessageTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv1MessageTypeCounter object
func (obj *patternFlowGtpv1MessageTypeCounter) SetCount(value uint32) PatternFlowGtpv1MessageTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv1MessageTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv1MessageTypeCounter) setDefault() {
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
