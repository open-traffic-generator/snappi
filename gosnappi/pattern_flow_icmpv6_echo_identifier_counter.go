package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6EchoIdentifierCounter *****
type patternFlowIcmpv6EchoIdentifierCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6EchoIdentifierCounter
	marshaller   marshalPatternFlowIcmpv6EchoIdentifierCounter
	unMarshaller unMarshalPatternFlowIcmpv6EchoIdentifierCounter
}

func NewPatternFlowIcmpv6EchoIdentifierCounter() PatternFlowIcmpv6EchoIdentifierCounter {
	obj := patternFlowIcmpv6EchoIdentifierCounter{obj: &otg.PatternFlowIcmpv6EchoIdentifierCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) msg() *otg.PatternFlowIcmpv6EchoIdentifierCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) setMsg(msg *otg.PatternFlowIcmpv6EchoIdentifierCounter) PatternFlowIcmpv6EchoIdentifierCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6EchoIdentifierCounter struct {
	obj *patternFlowIcmpv6EchoIdentifierCounter
}

type marshalPatternFlowIcmpv6EchoIdentifierCounter interface {
	// ToProto marshals PatternFlowIcmpv6EchoIdentifierCounter to protobuf object *otg.PatternFlowIcmpv6EchoIdentifierCounter
	ToProto() (*otg.PatternFlowIcmpv6EchoIdentifierCounter, error)
	// ToPbText marshals PatternFlowIcmpv6EchoIdentifierCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6EchoIdentifierCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6EchoIdentifierCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6EchoIdentifierCounter struct {
	obj *patternFlowIcmpv6EchoIdentifierCounter
}

type unMarshalPatternFlowIcmpv6EchoIdentifierCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6EchoIdentifierCounter from protobuf object *otg.PatternFlowIcmpv6EchoIdentifierCounter
	FromProto(msg *otg.PatternFlowIcmpv6EchoIdentifierCounter) (PatternFlowIcmpv6EchoIdentifierCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6EchoIdentifierCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6EchoIdentifierCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6EchoIdentifierCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) Marshal() marshalPatternFlowIcmpv6EchoIdentifierCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6EchoIdentifierCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) Unmarshal() unMarshalPatternFlowIcmpv6EchoIdentifierCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6EchoIdentifierCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6EchoIdentifierCounter) ToProto() (*otg.PatternFlowIcmpv6EchoIdentifierCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6EchoIdentifierCounter) FromProto(msg *otg.PatternFlowIcmpv6EchoIdentifierCounter) (PatternFlowIcmpv6EchoIdentifierCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6EchoIdentifierCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoIdentifierCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoIdentifierCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoIdentifierCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoIdentifierCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoIdentifierCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6EchoIdentifierCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) Clone() (PatternFlowIcmpv6EchoIdentifierCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6EchoIdentifierCounter()
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

// PatternFlowIcmpv6EchoIdentifierCounter is integer counter pattern
type PatternFlowIcmpv6EchoIdentifierCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6EchoIdentifierCounter to protobuf object *otg.PatternFlowIcmpv6EchoIdentifierCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6EchoIdentifierCounter
	// setMsg unmarshals PatternFlowIcmpv6EchoIdentifierCounter from protobuf object *otg.PatternFlowIcmpv6EchoIdentifierCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6EchoIdentifierCounter) PatternFlowIcmpv6EchoIdentifierCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6EchoIdentifierCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6EchoIdentifierCounter
	// validate validates PatternFlowIcmpv6EchoIdentifierCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6EchoIdentifierCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6EchoIdentifierCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6EchoIdentifierCounter
	SetStart(value uint32) PatternFlowIcmpv6EchoIdentifierCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6EchoIdentifierCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6EchoIdentifierCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6EchoIdentifierCounter
	SetStep(value uint32) PatternFlowIcmpv6EchoIdentifierCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6EchoIdentifierCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6EchoIdentifierCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6EchoIdentifierCounter
	SetCount(value uint32) PatternFlowIcmpv6EchoIdentifierCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6EchoIdentifierCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6EchoIdentifierCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6EchoIdentifierCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6EchoIdentifierCounter object
func (obj *patternFlowIcmpv6EchoIdentifierCounter) SetStart(value uint32) PatternFlowIcmpv6EchoIdentifierCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6EchoIdentifierCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6EchoIdentifierCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6EchoIdentifierCounter object
func (obj *patternFlowIcmpv6EchoIdentifierCounter) SetStep(value uint32) PatternFlowIcmpv6EchoIdentifierCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6EchoIdentifierCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6EchoIdentifierCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6EchoIdentifierCounter object
func (obj *patternFlowIcmpv6EchoIdentifierCounter) SetCount(value uint32) PatternFlowIcmpv6EchoIdentifierCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoIdentifierCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoIdentifierCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoIdentifierCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6EchoIdentifierCounter) setDefault() {
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
