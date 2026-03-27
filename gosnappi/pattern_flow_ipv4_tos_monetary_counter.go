package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosMonetaryCounter *****
type patternFlowIpv4TosMonetaryCounter struct {
	validation
	obj          *otg.PatternFlowIpv4TosMonetaryCounter
	marshaller   marshalPatternFlowIpv4TosMonetaryCounter
	unMarshaller unMarshalPatternFlowIpv4TosMonetaryCounter
}

func NewPatternFlowIpv4TosMonetaryCounter() PatternFlowIpv4TosMonetaryCounter {
	obj := patternFlowIpv4TosMonetaryCounter{obj: &otg.PatternFlowIpv4TosMonetaryCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosMonetaryCounter) msg() *otg.PatternFlowIpv4TosMonetaryCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TosMonetaryCounter) setMsg(msg *otg.PatternFlowIpv4TosMonetaryCounter) PatternFlowIpv4TosMonetaryCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosMonetaryCounter struct {
	obj *patternFlowIpv4TosMonetaryCounter
}

type marshalPatternFlowIpv4TosMonetaryCounter interface {
	// ToProto marshals PatternFlowIpv4TosMonetaryCounter to protobuf object *otg.PatternFlowIpv4TosMonetaryCounter
	ToProto() (*otg.PatternFlowIpv4TosMonetaryCounter, error)
	// ToPbText marshals PatternFlowIpv4TosMonetaryCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosMonetaryCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosMonetaryCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosMonetaryCounter struct {
	obj *patternFlowIpv4TosMonetaryCounter
}

type unMarshalPatternFlowIpv4TosMonetaryCounter interface {
	// FromProto unmarshals PatternFlowIpv4TosMonetaryCounter from protobuf object *otg.PatternFlowIpv4TosMonetaryCounter
	FromProto(msg *otg.PatternFlowIpv4TosMonetaryCounter) (PatternFlowIpv4TosMonetaryCounter, error)
	// FromPbText unmarshals PatternFlowIpv4TosMonetaryCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosMonetaryCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosMonetaryCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosMonetaryCounter) Marshal() marshalPatternFlowIpv4TosMonetaryCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosMonetaryCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosMonetaryCounter) Unmarshal() unMarshalPatternFlowIpv4TosMonetaryCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosMonetaryCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosMonetaryCounter) ToProto() (*otg.PatternFlowIpv4TosMonetaryCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosMonetaryCounter) FromProto(msg *otg.PatternFlowIpv4TosMonetaryCounter) (PatternFlowIpv4TosMonetaryCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosMonetaryCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosMonetaryCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosMonetaryCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosMonetaryCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosMonetaryCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosMonetaryCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosMonetaryCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosMonetaryCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosMonetaryCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosMonetaryCounter) Clone() (PatternFlowIpv4TosMonetaryCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosMonetaryCounter()
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

// PatternFlowIpv4TosMonetaryCounter is integer counter pattern
type PatternFlowIpv4TosMonetaryCounter interface {
	Validation
	// msg marshals PatternFlowIpv4TosMonetaryCounter to protobuf object *otg.PatternFlowIpv4TosMonetaryCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosMonetaryCounter
	// setMsg unmarshals PatternFlowIpv4TosMonetaryCounter from protobuf object *otg.PatternFlowIpv4TosMonetaryCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosMonetaryCounter) PatternFlowIpv4TosMonetaryCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosMonetaryCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosMonetaryCounter
	// validate validates PatternFlowIpv4TosMonetaryCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosMonetaryCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4TosMonetaryCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4TosMonetaryCounter
	SetStart(value uint32) PatternFlowIpv4TosMonetaryCounter
	// HasStart checks if Start has been set in PatternFlowIpv4TosMonetaryCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4TosMonetaryCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4TosMonetaryCounter
	SetStep(value uint32) PatternFlowIpv4TosMonetaryCounter
	// HasStep checks if Step has been set in PatternFlowIpv4TosMonetaryCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4TosMonetaryCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4TosMonetaryCounter
	SetCount(value uint32) PatternFlowIpv4TosMonetaryCounter
	// HasCount checks if Count has been set in PatternFlowIpv4TosMonetaryCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TosMonetaryCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TosMonetaryCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4TosMonetaryCounter object
func (obj *patternFlowIpv4TosMonetaryCounter) SetStart(value uint32) PatternFlowIpv4TosMonetaryCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TosMonetaryCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TosMonetaryCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4TosMonetaryCounter object
func (obj *patternFlowIpv4TosMonetaryCounter) SetStep(value uint32) PatternFlowIpv4TosMonetaryCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TosMonetaryCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TosMonetaryCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4TosMonetaryCounter object
func (obj *patternFlowIpv4TosMonetaryCounter) SetCount(value uint32) PatternFlowIpv4TosMonetaryCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4TosMonetaryCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosMonetaryCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosMonetaryCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosMonetaryCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4TosMonetaryCounter) setDefault() {
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
