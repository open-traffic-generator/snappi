package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpEchoSequenceNumberCounter *****
type patternFlowIcmpEchoSequenceNumberCounter struct {
	validation
	obj          *otg.PatternFlowIcmpEchoSequenceNumberCounter
	marshaller   marshalPatternFlowIcmpEchoSequenceNumberCounter
	unMarshaller unMarshalPatternFlowIcmpEchoSequenceNumberCounter
}

func NewPatternFlowIcmpEchoSequenceNumberCounter() PatternFlowIcmpEchoSequenceNumberCounter {
	obj := patternFlowIcmpEchoSequenceNumberCounter{obj: &otg.PatternFlowIcmpEchoSequenceNumberCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) msg() *otg.PatternFlowIcmpEchoSequenceNumberCounter {
	return obj.obj
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) setMsg(msg *otg.PatternFlowIcmpEchoSequenceNumberCounter) PatternFlowIcmpEchoSequenceNumberCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpEchoSequenceNumberCounter struct {
	obj *patternFlowIcmpEchoSequenceNumberCounter
}

type marshalPatternFlowIcmpEchoSequenceNumberCounter interface {
	// ToProto marshals PatternFlowIcmpEchoSequenceNumberCounter to protobuf object *otg.PatternFlowIcmpEchoSequenceNumberCounter
	ToProto() (*otg.PatternFlowIcmpEchoSequenceNumberCounter, error)
	// ToPbText marshals PatternFlowIcmpEchoSequenceNumberCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpEchoSequenceNumberCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpEchoSequenceNumberCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpEchoSequenceNumberCounter struct {
	obj *patternFlowIcmpEchoSequenceNumberCounter
}

type unMarshalPatternFlowIcmpEchoSequenceNumberCounter interface {
	// FromProto unmarshals PatternFlowIcmpEchoSequenceNumberCounter from protobuf object *otg.PatternFlowIcmpEchoSequenceNumberCounter
	FromProto(msg *otg.PatternFlowIcmpEchoSequenceNumberCounter) (PatternFlowIcmpEchoSequenceNumberCounter, error)
	// FromPbText unmarshals PatternFlowIcmpEchoSequenceNumberCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpEchoSequenceNumberCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpEchoSequenceNumberCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) Marshal() marshalPatternFlowIcmpEchoSequenceNumberCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpEchoSequenceNumberCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) Unmarshal() unMarshalPatternFlowIcmpEchoSequenceNumberCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpEchoSequenceNumberCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpEchoSequenceNumberCounter) ToProto() (*otg.PatternFlowIcmpEchoSequenceNumberCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpEchoSequenceNumberCounter) FromProto(msg *otg.PatternFlowIcmpEchoSequenceNumberCounter) (PatternFlowIcmpEchoSequenceNumberCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpEchoSequenceNumberCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoSequenceNumberCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpEchoSequenceNumberCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoSequenceNumberCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpEchoSequenceNumberCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoSequenceNumberCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpEchoSequenceNumberCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) Clone() (PatternFlowIcmpEchoSequenceNumberCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpEchoSequenceNumberCounter()
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

// PatternFlowIcmpEchoSequenceNumberCounter is integer counter pattern
type PatternFlowIcmpEchoSequenceNumberCounter interface {
	Validation
	// msg marshals PatternFlowIcmpEchoSequenceNumberCounter to protobuf object *otg.PatternFlowIcmpEchoSequenceNumberCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpEchoSequenceNumberCounter
	// setMsg unmarshals PatternFlowIcmpEchoSequenceNumberCounter from protobuf object *otg.PatternFlowIcmpEchoSequenceNumberCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpEchoSequenceNumberCounter) PatternFlowIcmpEchoSequenceNumberCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpEchoSequenceNumberCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpEchoSequenceNumberCounter
	// validate validates PatternFlowIcmpEchoSequenceNumberCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpEchoSequenceNumberCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpEchoSequenceNumberCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpEchoSequenceNumberCounter
	SetStart(value uint32) PatternFlowIcmpEchoSequenceNumberCounter
	// HasStart checks if Start has been set in PatternFlowIcmpEchoSequenceNumberCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpEchoSequenceNumberCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpEchoSequenceNumberCounter
	SetStep(value uint32) PatternFlowIcmpEchoSequenceNumberCounter
	// HasStep checks if Step has been set in PatternFlowIcmpEchoSequenceNumberCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpEchoSequenceNumberCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpEchoSequenceNumberCounter
	SetCount(value uint32) PatternFlowIcmpEchoSequenceNumberCounter
	// HasCount checks if Count has been set in PatternFlowIcmpEchoSequenceNumberCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpEchoSequenceNumberCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpEchoSequenceNumberCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpEchoSequenceNumberCounter object
func (obj *patternFlowIcmpEchoSequenceNumberCounter) SetStart(value uint32) PatternFlowIcmpEchoSequenceNumberCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpEchoSequenceNumberCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpEchoSequenceNumberCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpEchoSequenceNumberCounter object
func (obj *patternFlowIcmpEchoSequenceNumberCounter) SetStep(value uint32) PatternFlowIcmpEchoSequenceNumberCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpEchoSequenceNumberCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpEchoSequenceNumberCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpEchoSequenceNumberCounter object
func (obj *patternFlowIcmpEchoSequenceNumberCounter) SetCount(value uint32) PatternFlowIcmpEchoSequenceNumberCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoSequenceNumberCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoSequenceNumberCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoSequenceNumberCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpEchoSequenceNumberCounter) setDefault() {
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
