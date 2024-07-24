package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6EchoTypeCounter *****
type patternFlowIcmpv6EchoTypeCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6EchoTypeCounter
	marshaller   marshalPatternFlowIcmpv6EchoTypeCounter
	unMarshaller unMarshalPatternFlowIcmpv6EchoTypeCounter
}

func NewPatternFlowIcmpv6EchoTypeCounter() PatternFlowIcmpv6EchoTypeCounter {
	obj := patternFlowIcmpv6EchoTypeCounter{obj: &otg.PatternFlowIcmpv6EchoTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6EchoTypeCounter) msg() *otg.PatternFlowIcmpv6EchoTypeCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoTypeCounter) setMsg(msg *otg.PatternFlowIcmpv6EchoTypeCounter) PatternFlowIcmpv6EchoTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6EchoTypeCounter struct {
	obj *patternFlowIcmpv6EchoTypeCounter
}

type marshalPatternFlowIcmpv6EchoTypeCounter interface {
	// ToProto marshals PatternFlowIcmpv6EchoTypeCounter to protobuf object *otg.PatternFlowIcmpv6EchoTypeCounter
	ToProto() (*otg.PatternFlowIcmpv6EchoTypeCounter, error)
	// ToPbText marshals PatternFlowIcmpv6EchoTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6EchoTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6EchoTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6EchoTypeCounter struct {
	obj *patternFlowIcmpv6EchoTypeCounter
}

type unMarshalPatternFlowIcmpv6EchoTypeCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6EchoTypeCounter from protobuf object *otg.PatternFlowIcmpv6EchoTypeCounter
	FromProto(msg *otg.PatternFlowIcmpv6EchoTypeCounter) (PatternFlowIcmpv6EchoTypeCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6EchoTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6EchoTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6EchoTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6EchoTypeCounter) Marshal() marshalPatternFlowIcmpv6EchoTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6EchoTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6EchoTypeCounter) Unmarshal() unMarshalPatternFlowIcmpv6EchoTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6EchoTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6EchoTypeCounter) ToProto() (*otg.PatternFlowIcmpv6EchoTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6EchoTypeCounter) FromProto(msg *otg.PatternFlowIcmpv6EchoTypeCounter) (PatternFlowIcmpv6EchoTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6EchoTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6EchoTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6EchoTypeCounter) Clone() (PatternFlowIcmpv6EchoTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6EchoTypeCounter()
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

// PatternFlowIcmpv6EchoTypeCounter is integer counter pattern
type PatternFlowIcmpv6EchoTypeCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6EchoTypeCounter to protobuf object *otg.PatternFlowIcmpv6EchoTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6EchoTypeCounter
	// setMsg unmarshals PatternFlowIcmpv6EchoTypeCounter from protobuf object *otg.PatternFlowIcmpv6EchoTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6EchoTypeCounter) PatternFlowIcmpv6EchoTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6EchoTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6EchoTypeCounter
	// validate validates PatternFlowIcmpv6EchoTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6EchoTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6EchoTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6EchoTypeCounter
	SetStart(value uint32) PatternFlowIcmpv6EchoTypeCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6EchoTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6EchoTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6EchoTypeCounter
	SetStep(value uint32) PatternFlowIcmpv6EchoTypeCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6EchoTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6EchoTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6EchoTypeCounter
	SetCount(value uint32) PatternFlowIcmpv6EchoTypeCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6EchoTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6EchoTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6EchoTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6EchoTypeCounter object
func (obj *patternFlowIcmpv6EchoTypeCounter) SetStart(value uint32) PatternFlowIcmpv6EchoTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6EchoTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6EchoTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6EchoTypeCounter object
func (obj *patternFlowIcmpv6EchoTypeCounter) SetStep(value uint32) PatternFlowIcmpv6EchoTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6EchoTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6EchoTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6EchoTypeCounter object
func (obj *patternFlowIcmpv6EchoTypeCounter) SetCount(value uint32) PatternFlowIcmpv6EchoTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoTypeCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6EchoTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(128)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
