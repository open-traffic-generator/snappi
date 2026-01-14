package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPppAddressCounter *****
type patternFlowPppAddressCounter struct {
	validation
	obj          *otg.PatternFlowPppAddressCounter
	marshaller   marshalPatternFlowPppAddressCounter
	unMarshaller unMarshalPatternFlowPppAddressCounter
}

func NewPatternFlowPppAddressCounter() PatternFlowPppAddressCounter {
	obj := patternFlowPppAddressCounter{obj: &otg.PatternFlowPppAddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPppAddressCounter) msg() *otg.PatternFlowPppAddressCounter {
	return obj.obj
}

func (obj *patternFlowPppAddressCounter) setMsg(msg *otg.PatternFlowPppAddressCounter) PatternFlowPppAddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPppAddressCounter struct {
	obj *patternFlowPppAddressCounter
}

type marshalPatternFlowPppAddressCounter interface {
	// ToProto marshals PatternFlowPppAddressCounter to protobuf object *otg.PatternFlowPppAddressCounter
	ToProto() (*otg.PatternFlowPppAddressCounter, error)
	// ToPbText marshals PatternFlowPppAddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPppAddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPppAddressCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPppAddressCounter struct {
	obj *patternFlowPppAddressCounter
}

type unMarshalPatternFlowPppAddressCounter interface {
	// FromProto unmarshals PatternFlowPppAddressCounter from protobuf object *otg.PatternFlowPppAddressCounter
	FromProto(msg *otg.PatternFlowPppAddressCounter) (PatternFlowPppAddressCounter, error)
	// FromPbText unmarshals PatternFlowPppAddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPppAddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPppAddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPppAddressCounter) Marshal() marshalPatternFlowPppAddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPppAddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPppAddressCounter) Unmarshal() unMarshalPatternFlowPppAddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPppAddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPppAddressCounter) ToProto() (*otg.PatternFlowPppAddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPppAddressCounter) FromProto(msg *otg.PatternFlowPppAddressCounter) (PatternFlowPppAddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPppAddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPppAddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowPppAddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPppAddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowPppAddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPppAddressCounter) FromJson(value string) error {
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

func (obj *patternFlowPppAddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPppAddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPppAddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPppAddressCounter) Clone() (PatternFlowPppAddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPppAddressCounter()
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

// PatternFlowPppAddressCounter is integer counter pattern
type PatternFlowPppAddressCounter interface {
	Validation
	// msg marshals PatternFlowPppAddressCounter to protobuf object *otg.PatternFlowPppAddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowPppAddressCounter
	// setMsg unmarshals PatternFlowPppAddressCounter from protobuf object *otg.PatternFlowPppAddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPppAddressCounter) PatternFlowPppAddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowPppAddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPppAddressCounter
	// validate validates PatternFlowPppAddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPppAddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowPppAddressCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowPppAddressCounter
	SetStart(value uint32) PatternFlowPppAddressCounter
	// HasStart checks if Start has been set in PatternFlowPppAddressCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowPppAddressCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowPppAddressCounter
	SetStep(value uint32) PatternFlowPppAddressCounter
	// HasStep checks if Step has been set in PatternFlowPppAddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowPppAddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowPppAddressCounter
	SetCount(value uint32) PatternFlowPppAddressCounter
	// HasCount checks if Count has been set in PatternFlowPppAddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPppAddressCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowPppAddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowPppAddressCounter object
func (obj *patternFlowPppAddressCounter) SetStart(value uint32) PatternFlowPppAddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPppAddressCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowPppAddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowPppAddressCounter object
func (obj *patternFlowPppAddressCounter) SetStep(value uint32) PatternFlowPppAddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPppAddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowPppAddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowPppAddressCounter object
func (obj *patternFlowPppAddressCounter) SetCount(value uint32) PatternFlowPppAddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowPppAddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppAddressCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppAddressCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppAddressCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowPppAddressCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(255)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
