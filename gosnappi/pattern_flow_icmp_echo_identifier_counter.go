package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpEchoIdentifierCounter *****
type patternFlowIcmpEchoIdentifierCounter struct {
	validation
	obj          *otg.PatternFlowIcmpEchoIdentifierCounter
	marshaller   marshalPatternFlowIcmpEchoIdentifierCounter
	unMarshaller unMarshalPatternFlowIcmpEchoIdentifierCounter
}

func NewPatternFlowIcmpEchoIdentifierCounter() PatternFlowIcmpEchoIdentifierCounter {
	obj := patternFlowIcmpEchoIdentifierCounter{obj: &otg.PatternFlowIcmpEchoIdentifierCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpEchoIdentifierCounter) msg() *otg.PatternFlowIcmpEchoIdentifierCounter {
	return obj.obj
}

func (obj *patternFlowIcmpEchoIdentifierCounter) setMsg(msg *otg.PatternFlowIcmpEchoIdentifierCounter) PatternFlowIcmpEchoIdentifierCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpEchoIdentifierCounter struct {
	obj *patternFlowIcmpEchoIdentifierCounter
}

type marshalPatternFlowIcmpEchoIdentifierCounter interface {
	// ToProto marshals PatternFlowIcmpEchoIdentifierCounter to protobuf object *otg.PatternFlowIcmpEchoIdentifierCounter
	ToProto() (*otg.PatternFlowIcmpEchoIdentifierCounter, error)
	// ToPbText marshals PatternFlowIcmpEchoIdentifierCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpEchoIdentifierCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpEchoIdentifierCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpEchoIdentifierCounter struct {
	obj *patternFlowIcmpEchoIdentifierCounter
}

type unMarshalPatternFlowIcmpEchoIdentifierCounter interface {
	// FromProto unmarshals PatternFlowIcmpEchoIdentifierCounter from protobuf object *otg.PatternFlowIcmpEchoIdentifierCounter
	FromProto(msg *otg.PatternFlowIcmpEchoIdentifierCounter) (PatternFlowIcmpEchoIdentifierCounter, error)
	// FromPbText unmarshals PatternFlowIcmpEchoIdentifierCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpEchoIdentifierCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpEchoIdentifierCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpEchoIdentifierCounter) Marshal() marshalPatternFlowIcmpEchoIdentifierCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpEchoIdentifierCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpEchoIdentifierCounter) Unmarshal() unMarshalPatternFlowIcmpEchoIdentifierCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpEchoIdentifierCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpEchoIdentifierCounter) ToProto() (*otg.PatternFlowIcmpEchoIdentifierCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpEchoIdentifierCounter) FromProto(msg *otg.PatternFlowIcmpEchoIdentifierCounter) (PatternFlowIcmpEchoIdentifierCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpEchoIdentifierCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoIdentifierCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpEchoIdentifierCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoIdentifierCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpEchoIdentifierCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoIdentifierCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpEchoIdentifierCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoIdentifierCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoIdentifierCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpEchoIdentifierCounter) Clone() (PatternFlowIcmpEchoIdentifierCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpEchoIdentifierCounter()
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

// PatternFlowIcmpEchoIdentifierCounter is integer counter pattern
type PatternFlowIcmpEchoIdentifierCounter interface {
	Validation
	// msg marshals PatternFlowIcmpEchoIdentifierCounter to protobuf object *otg.PatternFlowIcmpEchoIdentifierCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpEchoIdentifierCounter
	// setMsg unmarshals PatternFlowIcmpEchoIdentifierCounter from protobuf object *otg.PatternFlowIcmpEchoIdentifierCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpEchoIdentifierCounter) PatternFlowIcmpEchoIdentifierCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpEchoIdentifierCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpEchoIdentifierCounter
	// validate validates PatternFlowIcmpEchoIdentifierCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpEchoIdentifierCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpEchoIdentifierCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpEchoIdentifierCounter
	SetStart(value uint32) PatternFlowIcmpEchoIdentifierCounter
	// HasStart checks if Start has been set in PatternFlowIcmpEchoIdentifierCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpEchoIdentifierCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpEchoIdentifierCounter
	SetStep(value uint32) PatternFlowIcmpEchoIdentifierCounter
	// HasStep checks if Step has been set in PatternFlowIcmpEchoIdentifierCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpEchoIdentifierCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpEchoIdentifierCounter
	SetCount(value uint32) PatternFlowIcmpEchoIdentifierCounter
	// HasCount checks if Count has been set in PatternFlowIcmpEchoIdentifierCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpEchoIdentifierCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpEchoIdentifierCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpEchoIdentifierCounter object
func (obj *patternFlowIcmpEchoIdentifierCounter) SetStart(value uint32) PatternFlowIcmpEchoIdentifierCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpEchoIdentifierCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpEchoIdentifierCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpEchoIdentifierCounter object
func (obj *patternFlowIcmpEchoIdentifierCounter) SetStep(value uint32) PatternFlowIcmpEchoIdentifierCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpEchoIdentifierCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpEchoIdentifierCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpEchoIdentifierCounter object
func (obj *patternFlowIcmpEchoIdentifierCounter) SetCount(value uint32) PatternFlowIcmpEchoIdentifierCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpEchoIdentifierCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoIdentifierCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoIdentifierCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoIdentifierCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpEchoIdentifierCounter) setDefault() {
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
