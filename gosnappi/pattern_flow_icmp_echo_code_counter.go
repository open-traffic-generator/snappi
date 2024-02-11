package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpEchoCodeCounter *****
type patternFlowIcmpEchoCodeCounter struct {
	validation
	obj          *otg.PatternFlowIcmpEchoCodeCounter
	marshaller   marshalPatternFlowIcmpEchoCodeCounter
	unMarshaller unMarshalPatternFlowIcmpEchoCodeCounter
}

func NewPatternFlowIcmpEchoCodeCounter() PatternFlowIcmpEchoCodeCounter {
	obj := patternFlowIcmpEchoCodeCounter{obj: &otg.PatternFlowIcmpEchoCodeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpEchoCodeCounter) msg() *otg.PatternFlowIcmpEchoCodeCounter {
	return obj.obj
}

func (obj *patternFlowIcmpEchoCodeCounter) setMsg(msg *otg.PatternFlowIcmpEchoCodeCounter) PatternFlowIcmpEchoCodeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpEchoCodeCounter struct {
	obj *patternFlowIcmpEchoCodeCounter
}

type marshalPatternFlowIcmpEchoCodeCounter interface {
	// ToProto marshals PatternFlowIcmpEchoCodeCounter to protobuf object *otg.PatternFlowIcmpEchoCodeCounter
	ToProto() (*otg.PatternFlowIcmpEchoCodeCounter, error)
	// ToPbText marshals PatternFlowIcmpEchoCodeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpEchoCodeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpEchoCodeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpEchoCodeCounter struct {
	obj *patternFlowIcmpEchoCodeCounter
}

type unMarshalPatternFlowIcmpEchoCodeCounter interface {
	// FromProto unmarshals PatternFlowIcmpEchoCodeCounter from protobuf object *otg.PatternFlowIcmpEchoCodeCounter
	FromProto(msg *otg.PatternFlowIcmpEchoCodeCounter) (PatternFlowIcmpEchoCodeCounter, error)
	// FromPbText unmarshals PatternFlowIcmpEchoCodeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpEchoCodeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpEchoCodeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpEchoCodeCounter) Marshal() marshalPatternFlowIcmpEchoCodeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpEchoCodeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpEchoCodeCounter) Unmarshal() unMarshalPatternFlowIcmpEchoCodeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpEchoCodeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpEchoCodeCounter) ToProto() (*otg.PatternFlowIcmpEchoCodeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpEchoCodeCounter) FromProto(msg *otg.PatternFlowIcmpEchoCodeCounter) (PatternFlowIcmpEchoCodeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpEchoCodeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoCodeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpEchoCodeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoCodeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpEchoCodeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoCodeCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpEchoCodeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoCodeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoCodeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpEchoCodeCounter) Clone() (PatternFlowIcmpEchoCodeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpEchoCodeCounter()
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

// PatternFlowIcmpEchoCodeCounter is integer counter pattern
type PatternFlowIcmpEchoCodeCounter interface {
	Validation
	// msg marshals PatternFlowIcmpEchoCodeCounter to protobuf object *otg.PatternFlowIcmpEchoCodeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpEchoCodeCounter
	// setMsg unmarshals PatternFlowIcmpEchoCodeCounter from protobuf object *otg.PatternFlowIcmpEchoCodeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpEchoCodeCounter) PatternFlowIcmpEchoCodeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpEchoCodeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpEchoCodeCounter
	// validate validates PatternFlowIcmpEchoCodeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpEchoCodeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpEchoCodeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpEchoCodeCounter
	SetStart(value uint32) PatternFlowIcmpEchoCodeCounter
	// HasStart checks if Start has been set in PatternFlowIcmpEchoCodeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpEchoCodeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpEchoCodeCounter
	SetStep(value uint32) PatternFlowIcmpEchoCodeCounter
	// HasStep checks if Step has been set in PatternFlowIcmpEchoCodeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpEchoCodeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpEchoCodeCounter
	SetCount(value uint32) PatternFlowIcmpEchoCodeCounter
	// HasCount checks if Count has been set in PatternFlowIcmpEchoCodeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpEchoCodeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpEchoCodeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpEchoCodeCounter object
func (obj *patternFlowIcmpEchoCodeCounter) SetStart(value uint32) PatternFlowIcmpEchoCodeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpEchoCodeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpEchoCodeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpEchoCodeCounter object
func (obj *patternFlowIcmpEchoCodeCounter) SetStep(value uint32) PatternFlowIcmpEchoCodeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpEchoCodeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpEchoCodeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpEchoCodeCounter object
func (obj *patternFlowIcmpEchoCodeCounter) SetCount(value uint32) PatternFlowIcmpEchoCodeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpEchoCodeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoCodeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoCodeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoCodeCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpEchoCodeCounter) setDefault() {
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
