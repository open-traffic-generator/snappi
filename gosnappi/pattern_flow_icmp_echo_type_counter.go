package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpEchoTypeCounter *****
type patternFlowIcmpEchoTypeCounter struct {
	validation
	obj          *otg.PatternFlowIcmpEchoTypeCounter
	marshaller   marshalPatternFlowIcmpEchoTypeCounter
	unMarshaller unMarshalPatternFlowIcmpEchoTypeCounter
}

func NewPatternFlowIcmpEchoTypeCounter() PatternFlowIcmpEchoTypeCounter {
	obj := patternFlowIcmpEchoTypeCounter{obj: &otg.PatternFlowIcmpEchoTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpEchoTypeCounter) msg() *otg.PatternFlowIcmpEchoTypeCounter {
	return obj.obj
}

func (obj *patternFlowIcmpEchoTypeCounter) setMsg(msg *otg.PatternFlowIcmpEchoTypeCounter) PatternFlowIcmpEchoTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpEchoTypeCounter struct {
	obj *patternFlowIcmpEchoTypeCounter
}

type marshalPatternFlowIcmpEchoTypeCounter interface {
	// ToProto marshals PatternFlowIcmpEchoTypeCounter to protobuf object *otg.PatternFlowIcmpEchoTypeCounter
	ToProto() (*otg.PatternFlowIcmpEchoTypeCounter, error)
	// ToPbText marshals PatternFlowIcmpEchoTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpEchoTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpEchoTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpEchoTypeCounter struct {
	obj *patternFlowIcmpEchoTypeCounter
}

type unMarshalPatternFlowIcmpEchoTypeCounter interface {
	// FromProto unmarshals PatternFlowIcmpEchoTypeCounter from protobuf object *otg.PatternFlowIcmpEchoTypeCounter
	FromProto(msg *otg.PatternFlowIcmpEchoTypeCounter) (PatternFlowIcmpEchoTypeCounter, error)
	// FromPbText unmarshals PatternFlowIcmpEchoTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpEchoTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpEchoTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpEchoTypeCounter) Marshal() marshalPatternFlowIcmpEchoTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpEchoTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpEchoTypeCounter) Unmarshal() unMarshalPatternFlowIcmpEchoTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpEchoTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpEchoTypeCounter) ToProto() (*otg.PatternFlowIcmpEchoTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpEchoTypeCounter) FromProto(msg *otg.PatternFlowIcmpEchoTypeCounter) (PatternFlowIcmpEchoTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpEchoTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpEchoTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpEchoTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpEchoTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpEchoTypeCounter) Clone() (PatternFlowIcmpEchoTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpEchoTypeCounter()
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

// PatternFlowIcmpEchoTypeCounter is integer counter pattern
type PatternFlowIcmpEchoTypeCounter interface {
	Validation
	// msg marshals PatternFlowIcmpEchoTypeCounter to protobuf object *otg.PatternFlowIcmpEchoTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpEchoTypeCounter
	// setMsg unmarshals PatternFlowIcmpEchoTypeCounter from protobuf object *otg.PatternFlowIcmpEchoTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpEchoTypeCounter) PatternFlowIcmpEchoTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpEchoTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpEchoTypeCounter
	// validate validates PatternFlowIcmpEchoTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpEchoTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpEchoTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpEchoTypeCounter
	SetStart(value uint32) PatternFlowIcmpEchoTypeCounter
	// HasStart checks if Start has been set in PatternFlowIcmpEchoTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpEchoTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpEchoTypeCounter
	SetStep(value uint32) PatternFlowIcmpEchoTypeCounter
	// HasStep checks if Step has been set in PatternFlowIcmpEchoTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpEchoTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpEchoTypeCounter
	SetCount(value uint32) PatternFlowIcmpEchoTypeCounter
	// HasCount checks if Count has been set in PatternFlowIcmpEchoTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpEchoTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpEchoTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpEchoTypeCounter object
func (obj *patternFlowIcmpEchoTypeCounter) SetStart(value uint32) PatternFlowIcmpEchoTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpEchoTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpEchoTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpEchoTypeCounter object
func (obj *patternFlowIcmpEchoTypeCounter) SetStep(value uint32) PatternFlowIcmpEchoTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpEchoTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpEchoTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpEchoTypeCounter object
func (obj *patternFlowIcmpEchoTypeCounter) SetCount(value uint32) PatternFlowIcmpEchoTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpEchoTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpEchoTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(8)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
