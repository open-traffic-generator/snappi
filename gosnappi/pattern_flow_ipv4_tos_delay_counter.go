package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosDelayCounter *****
type patternFlowIpv4TosDelayCounter struct {
	validation
	obj          *otg.PatternFlowIpv4TosDelayCounter
	marshaller   marshalPatternFlowIpv4TosDelayCounter
	unMarshaller unMarshalPatternFlowIpv4TosDelayCounter
}

func NewPatternFlowIpv4TosDelayCounter() PatternFlowIpv4TosDelayCounter {
	obj := patternFlowIpv4TosDelayCounter{obj: &otg.PatternFlowIpv4TosDelayCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosDelayCounter) msg() *otg.PatternFlowIpv4TosDelayCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TosDelayCounter) setMsg(msg *otg.PatternFlowIpv4TosDelayCounter) PatternFlowIpv4TosDelayCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosDelayCounter struct {
	obj *patternFlowIpv4TosDelayCounter
}

type marshalPatternFlowIpv4TosDelayCounter interface {
	// ToProto marshals PatternFlowIpv4TosDelayCounter to protobuf object *otg.PatternFlowIpv4TosDelayCounter
	ToProto() (*otg.PatternFlowIpv4TosDelayCounter, error)
	// ToPbText marshals PatternFlowIpv4TosDelayCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosDelayCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosDelayCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosDelayCounter struct {
	obj *patternFlowIpv4TosDelayCounter
}

type unMarshalPatternFlowIpv4TosDelayCounter interface {
	// FromProto unmarshals PatternFlowIpv4TosDelayCounter from protobuf object *otg.PatternFlowIpv4TosDelayCounter
	FromProto(msg *otg.PatternFlowIpv4TosDelayCounter) (PatternFlowIpv4TosDelayCounter, error)
	// FromPbText unmarshals PatternFlowIpv4TosDelayCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosDelayCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosDelayCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosDelayCounter) Marshal() marshalPatternFlowIpv4TosDelayCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosDelayCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosDelayCounter) Unmarshal() unMarshalPatternFlowIpv4TosDelayCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosDelayCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosDelayCounter) ToProto() (*otg.PatternFlowIpv4TosDelayCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosDelayCounter) FromProto(msg *otg.PatternFlowIpv4TosDelayCounter) (PatternFlowIpv4TosDelayCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosDelayCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosDelayCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosDelayCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosDelayCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosDelayCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosDelayCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosDelayCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosDelayCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosDelayCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosDelayCounter) Clone() (PatternFlowIpv4TosDelayCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosDelayCounter()
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

// PatternFlowIpv4TosDelayCounter is integer counter pattern
type PatternFlowIpv4TosDelayCounter interface {
	Validation
	// msg marshals PatternFlowIpv4TosDelayCounter to protobuf object *otg.PatternFlowIpv4TosDelayCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosDelayCounter
	// setMsg unmarshals PatternFlowIpv4TosDelayCounter from protobuf object *otg.PatternFlowIpv4TosDelayCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosDelayCounter) PatternFlowIpv4TosDelayCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosDelayCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosDelayCounter
	// validate validates PatternFlowIpv4TosDelayCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosDelayCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4TosDelayCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4TosDelayCounter
	SetStart(value uint32) PatternFlowIpv4TosDelayCounter
	// HasStart checks if Start has been set in PatternFlowIpv4TosDelayCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4TosDelayCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4TosDelayCounter
	SetStep(value uint32) PatternFlowIpv4TosDelayCounter
	// HasStep checks if Step has been set in PatternFlowIpv4TosDelayCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4TosDelayCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4TosDelayCounter
	SetCount(value uint32) PatternFlowIpv4TosDelayCounter
	// HasCount checks if Count has been set in PatternFlowIpv4TosDelayCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TosDelayCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TosDelayCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4TosDelayCounter object
func (obj *patternFlowIpv4TosDelayCounter) SetStart(value uint32) PatternFlowIpv4TosDelayCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TosDelayCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TosDelayCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4TosDelayCounter object
func (obj *patternFlowIpv4TosDelayCounter) SetStep(value uint32) PatternFlowIpv4TosDelayCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TosDelayCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TosDelayCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4TosDelayCounter object
func (obj *patternFlowIpv4TosDelayCounter) SetCount(value uint32) PatternFlowIpv4TosDelayCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4TosDelayCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosDelayCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosDelayCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosDelayCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4TosDelayCounter) setDefault() {
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
