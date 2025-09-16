package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6EchoCodeCounter *****
type patternFlowIcmpv6EchoCodeCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6EchoCodeCounter
	marshaller   marshalPatternFlowIcmpv6EchoCodeCounter
	unMarshaller unMarshalPatternFlowIcmpv6EchoCodeCounter
}

func NewPatternFlowIcmpv6EchoCodeCounter() PatternFlowIcmpv6EchoCodeCounter {
	obj := patternFlowIcmpv6EchoCodeCounter{obj: &otg.PatternFlowIcmpv6EchoCodeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6EchoCodeCounter) msg() *otg.PatternFlowIcmpv6EchoCodeCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoCodeCounter) setMsg(msg *otg.PatternFlowIcmpv6EchoCodeCounter) PatternFlowIcmpv6EchoCodeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6EchoCodeCounter struct {
	obj *patternFlowIcmpv6EchoCodeCounter
}

type marshalPatternFlowIcmpv6EchoCodeCounter interface {
	// ToProto marshals PatternFlowIcmpv6EchoCodeCounter to protobuf object *otg.PatternFlowIcmpv6EchoCodeCounter
	ToProto() (*otg.PatternFlowIcmpv6EchoCodeCounter, error)
	// ToPbText marshals PatternFlowIcmpv6EchoCodeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6EchoCodeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6EchoCodeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6EchoCodeCounter struct {
	obj *patternFlowIcmpv6EchoCodeCounter
}

type unMarshalPatternFlowIcmpv6EchoCodeCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6EchoCodeCounter from protobuf object *otg.PatternFlowIcmpv6EchoCodeCounter
	FromProto(msg *otg.PatternFlowIcmpv6EchoCodeCounter) (PatternFlowIcmpv6EchoCodeCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6EchoCodeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6EchoCodeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6EchoCodeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6EchoCodeCounter) Marshal() marshalPatternFlowIcmpv6EchoCodeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6EchoCodeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6EchoCodeCounter) Unmarshal() unMarshalPatternFlowIcmpv6EchoCodeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6EchoCodeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6EchoCodeCounter) ToProto() (*otg.PatternFlowIcmpv6EchoCodeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6EchoCodeCounter) FromProto(msg *otg.PatternFlowIcmpv6EchoCodeCounter) (PatternFlowIcmpv6EchoCodeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6EchoCodeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoCodeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoCodeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoCodeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoCodeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoCodeCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6EchoCodeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoCodeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoCodeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6EchoCodeCounter) Clone() (PatternFlowIcmpv6EchoCodeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6EchoCodeCounter()
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

// PatternFlowIcmpv6EchoCodeCounter is integer counter pattern
type PatternFlowIcmpv6EchoCodeCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6EchoCodeCounter to protobuf object *otg.PatternFlowIcmpv6EchoCodeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6EchoCodeCounter
	// setMsg unmarshals PatternFlowIcmpv6EchoCodeCounter from protobuf object *otg.PatternFlowIcmpv6EchoCodeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6EchoCodeCounter) PatternFlowIcmpv6EchoCodeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6EchoCodeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6EchoCodeCounter
	// validate validates PatternFlowIcmpv6EchoCodeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6EchoCodeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6EchoCodeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6EchoCodeCounter
	SetStart(value uint32) PatternFlowIcmpv6EchoCodeCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6EchoCodeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6EchoCodeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6EchoCodeCounter
	SetStep(value uint32) PatternFlowIcmpv6EchoCodeCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6EchoCodeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6EchoCodeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6EchoCodeCounter
	SetCount(value uint32) PatternFlowIcmpv6EchoCodeCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6EchoCodeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6EchoCodeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6EchoCodeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6EchoCodeCounter object
func (obj *patternFlowIcmpv6EchoCodeCounter) SetStart(value uint32) PatternFlowIcmpv6EchoCodeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6EchoCodeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6EchoCodeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6EchoCodeCounter object
func (obj *patternFlowIcmpv6EchoCodeCounter) SetStep(value uint32) PatternFlowIcmpv6EchoCodeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6EchoCodeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6EchoCodeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6EchoCodeCounter object
func (obj *patternFlowIcmpv6EchoCodeCounter) SetCount(value uint32) PatternFlowIcmpv6EchoCodeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoCodeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoCodeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoCodeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoCodeCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6EchoCodeCounter) setDefault() {
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
