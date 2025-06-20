package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4DscpPhbCounter *****
type patternFlowIpv4DscpPhbCounter struct {
	validation
	obj          *otg.PatternFlowIpv4DscpPhbCounter
	marshaller   marshalPatternFlowIpv4DscpPhbCounter
	unMarshaller unMarshalPatternFlowIpv4DscpPhbCounter
}

func NewPatternFlowIpv4DscpPhbCounter() PatternFlowIpv4DscpPhbCounter {
	obj := patternFlowIpv4DscpPhbCounter{obj: &otg.PatternFlowIpv4DscpPhbCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4DscpPhbCounter) msg() *otg.PatternFlowIpv4DscpPhbCounter {
	return obj.obj
}

func (obj *patternFlowIpv4DscpPhbCounter) setMsg(msg *otg.PatternFlowIpv4DscpPhbCounter) PatternFlowIpv4DscpPhbCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4DscpPhbCounter struct {
	obj *patternFlowIpv4DscpPhbCounter
}

type marshalPatternFlowIpv4DscpPhbCounter interface {
	// ToProto marshals PatternFlowIpv4DscpPhbCounter to protobuf object *otg.PatternFlowIpv4DscpPhbCounter
	ToProto() (*otg.PatternFlowIpv4DscpPhbCounter, error)
	// ToPbText marshals PatternFlowIpv4DscpPhbCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4DscpPhbCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4DscpPhbCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4DscpPhbCounter struct {
	obj *patternFlowIpv4DscpPhbCounter
}

type unMarshalPatternFlowIpv4DscpPhbCounter interface {
	// FromProto unmarshals PatternFlowIpv4DscpPhbCounter from protobuf object *otg.PatternFlowIpv4DscpPhbCounter
	FromProto(msg *otg.PatternFlowIpv4DscpPhbCounter) (PatternFlowIpv4DscpPhbCounter, error)
	// FromPbText unmarshals PatternFlowIpv4DscpPhbCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4DscpPhbCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4DscpPhbCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4DscpPhbCounter) Marshal() marshalPatternFlowIpv4DscpPhbCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4DscpPhbCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4DscpPhbCounter) Unmarshal() unMarshalPatternFlowIpv4DscpPhbCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4DscpPhbCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4DscpPhbCounter) ToProto() (*otg.PatternFlowIpv4DscpPhbCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4DscpPhbCounter) FromProto(msg *otg.PatternFlowIpv4DscpPhbCounter) (PatternFlowIpv4DscpPhbCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4DscpPhbCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpPhbCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4DscpPhbCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpPhbCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4DscpPhbCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpPhbCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4DscpPhbCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DscpPhbCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DscpPhbCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4DscpPhbCounter) Clone() (PatternFlowIpv4DscpPhbCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4DscpPhbCounter()
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

// PatternFlowIpv4DscpPhbCounter is integer counter pattern
type PatternFlowIpv4DscpPhbCounter interface {
	Validation
	// msg marshals PatternFlowIpv4DscpPhbCounter to protobuf object *otg.PatternFlowIpv4DscpPhbCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4DscpPhbCounter
	// setMsg unmarshals PatternFlowIpv4DscpPhbCounter from protobuf object *otg.PatternFlowIpv4DscpPhbCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4DscpPhbCounter) PatternFlowIpv4DscpPhbCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4DscpPhbCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4DscpPhbCounter
	// validate validates PatternFlowIpv4DscpPhbCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4DscpPhbCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4DscpPhbCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4DscpPhbCounter
	SetStart(value uint32) PatternFlowIpv4DscpPhbCounter
	// HasStart checks if Start has been set in PatternFlowIpv4DscpPhbCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4DscpPhbCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4DscpPhbCounter
	SetStep(value uint32) PatternFlowIpv4DscpPhbCounter
	// HasStep checks if Step has been set in PatternFlowIpv4DscpPhbCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4DscpPhbCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4DscpPhbCounter
	SetCount(value uint32) PatternFlowIpv4DscpPhbCounter
	// HasCount checks if Count has been set in PatternFlowIpv4DscpPhbCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4DscpPhbCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4DscpPhbCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4DscpPhbCounter object
func (obj *patternFlowIpv4DscpPhbCounter) SetStart(value uint32) PatternFlowIpv4DscpPhbCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4DscpPhbCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4DscpPhbCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4DscpPhbCounter object
func (obj *patternFlowIpv4DscpPhbCounter) SetStep(value uint32) PatternFlowIpv4DscpPhbCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4DscpPhbCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4DscpPhbCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4DscpPhbCounter object
func (obj *patternFlowIpv4DscpPhbCounter) SetCount(value uint32) PatternFlowIpv4DscpPhbCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4DscpPhbCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 63 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpPhbCounter.Start <= 63 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 63 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpPhbCounter.Step <= 63 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 63 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpPhbCounter.Count <= 63 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4DscpPhbCounter) setDefault() {
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
