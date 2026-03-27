package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosUnusedCounter *****
type patternFlowIpv4TosUnusedCounter struct {
	validation
	obj          *otg.PatternFlowIpv4TosUnusedCounter
	marshaller   marshalPatternFlowIpv4TosUnusedCounter
	unMarshaller unMarshalPatternFlowIpv4TosUnusedCounter
}

func NewPatternFlowIpv4TosUnusedCounter() PatternFlowIpv4TosUnusedCounter {
	obj := patternFlowIpv4TosUnusedCounter{obj: &otg.PatternFlowIpv4TosUnusedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosUnusedCounter) msg() *otg.PatternFlowIpv4TosUnusedCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TosUnusedCounter) setMsg(msg *otg.PatternFlowIpv4TosUnusedCounter) PatternFlowIpv4TosUnusedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosUnusedCounter struct {
	obj *patternFlowIpv4TosUnusedCounter
}

type marshalPatternFlowIpv4TosUnusedCounter interface {
	// ToProto marshals PatternFlowIpv4TosUnusedCounter to protobuf object *otg.PatternFlowIpv4TosUnusedCounter
	ToProto() (*otg.PatternFlowIpv4TosUnusedCounter, error)
	// ToPbText marshals PatternFlowIpv4TosUnusedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosUnusedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosUnusedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosUnusedCounter struct {
	obj *patternFlowIpv4TosUnusedCounter
}

type unMarshalPatternFlowIpv4TosUnusedCounter interface {
	// FromProto unmarshals PatternFlowIpv4TosUnusedCounter from protobuf object *otg.PatternFlowIpv4TosUnusedCounter
	FromProto(msg *otg.PatternFlowIpv4TosUnusedCounter) (PatternFlowIpv4TosUnusedCounter, error)
	// FromPbText unmarshals PatternFlowIpv4TosUnusedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosUnusedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosUnusedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosUnusedCounter) Marshal() marshalPatternFlowIpv4TosUnusedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosUnusedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosUnusedCounter) Unmarshal() unMarshalPatternFlowIpv4TosUnusedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosUnusedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosUnusedCounter) ToProto() (*otg.PatternFlowIpv4TosUnusedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosUnusedCounter) FromProto(msg *otg.PatternFlowIpv4TosUnusedCounter) (PatternFlowIpv4TosUnusedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosUnusedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosUnusedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosUnusedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosUnusedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosUnusedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosUnusedCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosUnusedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosUnusedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosUnusedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosUnusedCounter) Clone() (PatternFlowIpv4TosUnusedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosUnusedCounter()
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

// PatternFlowIpv4TosUnusedCounter is integer counter pattern
type PatternFlowIpv4TosUnusedCounter interface {
	Validation
	// msg marshals PatternFlowIpv4TosUnusedCounter to protobuf object *otg.PatternFlowIpv4TosUnusedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosUnusedCounter
	// setMsg unmarshals PatternFlowIpv4TosUnusedCounter from protobuf object *otg.PatternFlowIpv4TosUnusedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosUnusedCounter) PatternFlowIpv4TosUnusedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosUnusedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosUnusedCounter
	// validate validates PatternFlowIpv4TosUnusedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosUnusedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4TosUnusedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4TosUnusedCounter
	SetStart(value uint32) PatternFlowIpv4TosUnusedCounter
	// HasStart checks if Start has been set in PatternFlowIpv4TosUnusedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4TosUnusedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4TosUnusedCounter
	SetStep(value uint32) PatternFlowIpv4TosUnusedCounter
	// HasStep checks if Step has been set in PatternFlowIpv4TosUnusedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4TosUnusedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4TosUnusedCounter
	SetCount(value uint32) PatternFlowIpv4TosUnusedCounter
	// HasCount checks if Count has been set in PatternFlowIpv4TosUnusedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TosUnusedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TosUnusedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4TosUnusedCounter object
func (obj *patternFlowIpv4TosUnusedCounter) SetStart(value uint32) PatternFlowIpv4TosUnusedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TosUnusedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TosUnusedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4TosUnusedCounter object
func (obj *patternFlowIpv4TosUnusedCounter) SetStep(value uint32) PatternFlowIpv4TosUnusedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TosUnusedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TosUnusedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4TosUnusedCounter object
func (obj *patternFlowIpv4TosUnusedCounter) SetCount(value uint32) PatternFlowIpv4TosUnusedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4TosUnusedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosUnusedCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosUnusedCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosUnusedCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4TosUnusedCounter) setDefault() {
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
