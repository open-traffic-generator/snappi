package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6EchoSequenceNumberCounter *****
type patternFlowIcmpv6EchoSequenceNumberCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6EchoSequenceNumberCounter
	marshaller   marshalPatternFlowIcmpv6EchoSequenceNumberCounter
	unMarshaller unMarshalPatternFlowIcmpv6EchoSequenceNumberCounter
}

func NewPatternFlowIcmpv6EchoSequenceNumberCounter() PatternFlowIcmpv6EchoSequenceNumberCounter {
	obj := patternFlowIcmpv6EchoSequenceNumberCounter{obj: &otg.PatternFlowIcmpv6EchoSequenceNumberCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) msg() *otg.PatternFlowIcmpv6EchoSequenceNumberCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) setMsg(msg *otg.PatternFlowIcmpv6EchoSequenceNumberCounter) PatternFlowIcmpv6EchoSequenceNumberCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6EchoSequenceNumberCounter struct {
	obj *patternFlowIcmpv6EchoSequenceNumberCounter
}

type marshalPatternFlowIcmpv6EchoSequenceNumberCounter interface {
	// ToProto marshals PatternFlowIcmpv6EchoSequenceNumberCounter to protobuf object *otg.PatternFlowIcmpv6EchoSequenceNumberCounter
	ToProto() (*otg.PatternFlowIcmpv6EchoSequenceNumberCounter, error)
	// ToPbText marshals PatternFlowIcmpv6EchoSequenceNumberCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6EchoSequenceNumberCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6EchoSequenceNumberCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6EchoSequenceNumberCounter struct {
	obj *patternFlowIcmpv6EchoSequenceNumberCounter
}

type unMarshalPatternFlowIcmpv6EchoSequenceNumberCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6EchoSequenceNumberCounter from protobuf object *otg.PatternFlowIcmpv6EchoSequenceNumberCounter
	FromProto(msg *otg.PatternFlowIcmpv6EchoSequenceNumberCounter) (PatternFlowIcmpv6EchoSequenceNumberCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6EchoSequenceNumberCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6EchoSequenceNumberCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6EchoSequenceNumberCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) Marshal() marshalPatternFlowIcmpv6EchoSequenceNumberCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6EchoSequenceNumberCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) Unmarshal() unMarshalPatternFlowIcmpv6EchoSequenceNumberCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6EchoSequenceNumberCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6EchoSequenceNumberCounter) ToProto() (*otg.PatternFlowIcmpv6EchoSequenceNumberCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6EchoSequenceNumberCounter) FromProto(msg *otg.PatternFlowIcmpv6EchoSequenceNumberCounter) (PatternFlowIcmpv6EchoSequenceNumberCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6EchoSequenceNumberCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoSequenceNumberCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoSequenceNumberCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoSequenceNumberCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoSequenceNumberCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoSequenceNumberCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) Clone() (PatternFlowIcmpv6EchoSequenceNumberCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6EchoSequenceNumberCounter()
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

// PatternFlowIcmpv6EchoSequenceNumberCounter is integer counter pattern
type PatternFlowIcmpv6EchoSequenceNumberCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6EchoSequenceNumberCounter to protobuf object *otg.PatternFlowIcmpv6EchoSequenceNumberCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6EchoSequenceNumberCounter
	// setMsg unmarshals PatternFlowIcmpv6EchoSequenceNumberCounter from protobuf object *otg.PatternFlowIcmpv6EchoSequenceNumberCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6EchoSequenceNumberCounter) PatternFlowIcmpv6EchoSequenceNumberCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6EchoSequenceNumberCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6EchoSequenceNumberCounter
	// validate validates PatternFlowIcmpv6EchoSequenceNumberCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6EchoSequenceNumberCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6EchoSequenceNumberCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6EchoSequenceNumberCounter
	SetStart(value uint32) PatternFlowIcmpv6EchoSequenceNumberCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6EchoSequenceNumberCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6EchoSequenceNumberCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6EchoSequenceNumberCounter
	SetStep(value uint32) PatternFlowIcmpv6EchoSequenceNumberCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6EchoSequenceNumberCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6EchoSequenceNumberCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6EchoSequenceNumberCounter
	SetCount(value uint32) PatternFlowIcmpv6EchoSequenceNumberCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6EchoSequenceNumberCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6EchoSequenceNumberCounter object
func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) SetStart(value uint32) PatternFlowIcmpv6EchoSequenceNumberCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6EchoSequenceNumberCounter object
func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) SetStep(value uint32) PatternFlowIcmpv6EchoSequenceNumberCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6EchoSequenceNumberCounter object
func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) SetCount(value uint32) PatternFlowIcmpv6EchoSequenceNumberCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoSequenceNumberCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoSequenceNumberCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoSequenceNumberCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6EchoSequenceNumberCounter) setDefault() {
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
