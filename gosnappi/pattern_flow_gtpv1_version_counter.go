package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1VersionCounter *****
type patternFlowGtpv1VersionCounter struct {
	validation
	obj          *otg.PatternFlowGtpv1VersionCounter
	marshaller   marshalPatternFlowGtpv1VersionCounter
	unMarshaller unMarshalPatternFlowGtpv1VersionCounter
}

func NewPatternFlowGtpv1VersionCounter() PatternFlowGtpv1VersionCounter {
	obj := patternFlowGtpv1VersionCounter{obj: &otg.PatternFlowGtpv1VersionCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1VersionCounter) msg() *otg.PatternFlowGtpv1VersionCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1VersionCounter) setMsg(msg *otg.PatternFlowGtpv1VersionCounter) PatternFlowGtpv1VersionCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1VersionCounter struct {
	obj *patternFlowGtpv1VersionCounter
}

type marshalPatternFlowGtpv1VersionCounter interface {
	// ToProto marshals PatternFlowGtpv1VersionCounter to protobuf object *otg.PatternFlowGtpv1VersionCounter
	ToProto() (*otg.PatternFlowGtpv1VersionCounter, error)
	// ToPbText marshals PatternFlowGtpv1VersionCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1VersionCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1VersionCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1VersionCounter struct {
	obj *patternFlowGtpv1VersionCounter
}

type unMarshalPatternFlowGtpv1VersionCounter interface {
	// FromProto unmarshals PatternFlowGtpv1VersionCounter from protobuf object *otg.PatternFlowGtpv1VersionCounter
	FromProto(msg *otg.PatternFlowGtpv1VersionCounter) (PatternFlowGtpv1VersionCounter, error)
	// FromPbText unmarshals PatternFlowGtpv1VersionCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1VersionCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1VersionCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1VersionCounter) Marshal() marshalPatternFlowGtpv1VersionCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1VersionCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1VersionCounter) Unmarshal() unMarshalPatternFlowGtpv1VersionCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1VersionCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1VersionCounter) ToProto() (*otg.PatternFlowGtpv1VersionCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1VersionCounter) FromProto(msg *otg.PatternFlowGtpv1VersionCounter) (PatternFlowGtpv1VersionCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1VersionCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1VersionCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1VersionCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1VersionCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1VersionCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1VersionCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv1VersionCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1VersionCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1VersionCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1VersionCounter) Clone() (PatternFlowGtpv1VersionCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1VersionCounter()
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

// PatternFlowGtpv1VersionCounter is integer counter pattern
type PatternFlowGtpv1VersionCounter interface {
	Validation
	// msg marshals PatternFlowGtpv1VersionCounter to protobuf object *otg.PatternFlowGtpv1VersionCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1VersionCounter
	// setMsg unmarshals PatternFlowGtpv1VersionCounter from protobuf object *otg.PatternFlowGtpv1VersionCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1VersionCounter) PatternFlowGtpv1VersionCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1VersionCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1VersionCounter
	// validate validates PatternFlowGtpv1VersionCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1VersionCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv1VersionCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv1VersionCounter
	SetStart(value uint32) PatternFlowGtpv1VersionCounter
	// HasStart checks if Start has been set in PatternFlowGtpv1VersionCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv1VersionCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv1VersionCounter
	SetStep(value uint32) PatternFlowGtpv1VersionCounter
	// HasStep checks if Step has been set in PatternFlowGtpv1VersionCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv1VersionCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv1VersionCounter
	SetCount(value uint32) PatternFlowGtpv1VersionCounter
	// HasCount checks if Count has been set in PatternFlowGtpv1VersionCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1VersionCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1VersionCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv1VersionCounter object
func (obj *patternFlowGtpv1VersionCounter) SetStart(value uint32) PatternFlowGtpv1VersionCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1VersionCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1VersionCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv1VersionCounter object
func (obj *patternFlowGtpv1VersionCounter) SetStep(value uint32) PatternFlowGtpv1VersionCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1VersionCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1VersionCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv1VersionCounter object
func (obj *patternFlowGtpv1VersionCounter) SetCount(value uint32) PatternFlowGtpv1VersionCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv1VersionCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1VersionCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1VersionCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1VersionCounter.Count <= 8 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv1VersionCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
