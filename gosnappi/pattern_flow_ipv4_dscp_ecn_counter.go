package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4DscpEcnCounter *****
type patternFlowIpv4DscpEcnCounter struct {
	validation
	obj          *otg.PatternFlowIpv4DscpEcnCounter
	marshaller   marshalPatternFlowIpv4DscpEcnCounter
	unMarshaller unMarshalPatternFlowIpv4DscpEcnCounter
}

func NewPatternFlowIpv4DscpEcnCounter() PatternFlowIpv4DscpEcnCounter {
	obj := patternFlowIpv4DscpEcnCounter{obj: &otg.PatternFlowIpv4DscpEcnCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4DscpEcnCounter) msg() *otg.PatternFlowIpv4DscpEcnCounter {
	return obj.obj
}

func (obj *patternFlowIpv4DscpEcnCounter) setMsg(msg *otg.PatternFlowIpv4DscpEcnCounter) PatternFlowIpv4DscpEcnCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4DscpEcnCounter struct {
	obj *patternFlowIpv4DscpEcnCounter
}

type marshalPatternFlowIpv4DscpEcnCounter interface {
	// ToProto marshals PatternFlowIpv4DscpEcnCounter to protobuf object *otg.PatternFlowIpv4DscpEcnCounter
	ToProto() (*otg.PatternFlowIpv4DscpEcnCounter, error)
	// ToPbText marshals PatternFlowIpv4DscpEcnCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4DscpEcnCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4DscpEcnCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4DscpEcnCounter struct {
	obj *patternFlowIpv4DscpEcnCounter
}

type unMarshalPatternFlowIpv4DscpEcnCounter interface {
	// FromProto unmarshals PatternFlowIpv4DscpEcnCounter from protobuf object *otg.PatternFlowIpv4DscpEcnCounter
	FromProto(msg *otg.PatternFlowIpv4DscpEcnCounter) (PatternFlowIpv4DscpEcnCounter, error)
	// FromPbText unmarshals PatternFlowIpv4DscpEcnCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4DscpEcnCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4DscpEcnCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4DscpEcnCounter) Marshal() marshalPatternFlowIpv4DscpEcnCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4DscpEcnCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4DscpEcnCounter) Unmarshal() unMarshalPatternFlowIpv4DscpEcnCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4DscpEcnCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4DscpEcnCounter) ToProto() (*otg.PatternFlowIpv4DscpEcnCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4DscpEcnCounter) FromProto(msg *otg.PatternFlowIpv4DscpEcnCounter) (PatternFlowIpv4DscpEcnCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4DscpEcnCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpEcnCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4DscpEcnCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpEcnCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4DscpEcnCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpEcnCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4DscpEcnCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DscpEcnCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DscpEcnCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4DscpEcnCounter) Clone() (PatternFlowIpv4DscpEcnCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4DscpEcnCounter()
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

// PatternFlowIpv4DscpEcnCounter is integer counter pattern
type PatternFlowIpv4DscpEcnCounter interface {
	Validation
	// msg marshals PatternFlowIpv4DscpEcnCounter to protobuf object *otg.PatternFlowIpv4DscpEcnCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4DscpEcnCounter
	// setMsg unmarshals PatternFlowIpv4DscpEcnCounter from protobuf object *otg.PatternFlowIpv4DscpEcnCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4DscpEcnCounter) PatternFlowIpv4DscpEcnCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4DscpEcnCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4DscpEcnCounter
	// validate validates PatternFlowIpv4DscpEcnCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4DscpEcnCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4DscpEcnCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4DscpEcnCounter
	SetStart(value uint32) PatternFlowIpv4DscpEcnCounter
	// HasStart checks if Start has been set in PatternFlowIpv4DscpEcnCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4DscpEcnCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4DscpEcnCounter
	SetStep(value uint32) PatternFlowIpv4DscpEcnCounter
	// HasStep checks if Step has been set in PatternFlowIpv4DscpEcnCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4DscpEcnCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4DscpEcnCounter
	SetCount(value uint32) PatternFlowIpv4DscpEcnCounter
	// HasCount checks if Count has been set in PatternFlowIpv4DscpEcnCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4DscpEcnCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4DscpEcnCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4DscpEcnCounter object
func (obj *patternFlowIpv4DscpEcnCounter) SetStart(value uint32) PatternFlowIpv4DscpEcnCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4DscpEcnCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4DscpEcnCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4DscpEcnCounter object
func (obj *patternFlowIpv4DscpEcnCounter) SetStep(value uint32) PatternFlowIpv4DscpEcnCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4DscpEcnCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4DscpEcnCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4DscpEcnCounter object
func (obj *patternFlowIpv4DscpEcnCounter) SetCount(value uint32) PatternFlowIpv4DscpEcnCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4DscpEcnCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpEcnCounter.Start <= 3 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpEcnCounter.Step <= 3 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpEcnCounter.Count <= 3 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4DscpEcnCounter) setDefault() {
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
