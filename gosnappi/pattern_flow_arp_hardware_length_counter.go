package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpHardwareLengthCounter *****
type patternFlowArpHardwareLengthCounter struct {
	validation
	obj          *otg.PatternFlowArpHardwareLengthCounter
	marshaller   marshalPatternFlowArpHardwareLengthCounter
	unMarshaller unMarshalPatternFlowArpHardwareLengthCounter
}

func NewPatternFlowArpHardwareLengthCounter() PatternFlowArpHardwareLengthCounter {
	obj := patternFlowArpHardwareLengthCounter{obj: &otg.PatternFlowArpHardwareLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpHardwareLengthCounter) msg() *otg.PatternFlowArpHardwareLengthCounter {
	return obj.obj
}

func (obj *patternFlowArpHardwareLengthCounter) setMsg(msg *otg.PatternFlowArpHardwareLengthCounter) PatternFlowArpHardwareLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpHardwareLengthCounter struct {
	obj *patternFlowArpHardwareLengthCounter
}

type marshalPatternFlowArpHardwareLengthCounter interface {
	// ToProto marshals PatternFlowArpHardwareLengthCounter to protobuf object *otg.PatternFlowArpHardwareLengthCounter
	ToProto() (*otg.PatternFlowArpHardwareLengthCounter, error)
	// ToPbText marshals PatternFlowArpHardwareLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpHardwareLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpHardwareLengthCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowArpHardwareLengthCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowArpHardwareLengthCounter struct {
	obj *patternFlowArpHardwareLengthCounter
}

type unMarshalPatternFlowArpHardwareLengthCounter interface {
	// FromProto unmarshals PatternFlowArpHardwareLengthCounter from protobuf object *otg.PatternFlowArpHardwareLengthCounter
	FromProto(msg *otg.PatternFlowArpHardwareLengthCounter) (PatternFlowArpHardwareLengthCounter, error)
	// FromPbText unmarshals PatternFlowArpHardwareLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpHardwareLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpHardwareLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpHardwareLengthCounter) Marshal() marshalPatternFlowArpHardwareLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpHardwareLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpHardwareLengthCounter) Unmarshal() unMarshalPatternFlowArpHardwareLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpHardwareLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpHardwareLengthCounter) ToProto() (*otg.PatternFlowArpHardwareLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpHardwareLengthCounter) FromProto(msg *otg.PatternFlowArpHardwareLengthCounter) (PatternFlowArpHardwareLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpHardwareLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpHardwareLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpHardwareLengthCounter) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowArpHardwareLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowArpHardwareLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpHardwareLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpHardwareLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpHardwareLengthCounter) Clone() (PatternFlowArpHardwareLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpHardwareLengthCounter()
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

// PatternFlowArpHardwareLengthCounter is integer counter pattern
type PatternFlowArpHardwareLengthCounter interface {
	Validation
	// msg marshals PatternFlowArpHardwareLengthCounter to protobuf object *otg.PatternFlowArpHardwareLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowArpHardwareLengthCounter
	// setMsg unmarshals PatternFlowArpHardwareLengthCounter from protobuf object *otg.PatternFlowArpHardwareLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpHardwareLengthCounter) PatternFlowArpHardwareLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowArpHardwareLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpHardwareLengthCounter
	// validate validates PatternFlowArpHardwareLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpHardwareLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowArpHardwareLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowArpHardwareLengthCounter
	SetStart(value uint32) PatternFlowArpHardwareLengthCounter
	// HasStart checks if Start has been set in PatternFlowArpHardwareLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowArpHardwareLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowArpHardwareLengthCounter
	SetStep(value uint32) PatternFlowArpHardwareLengthCounter
	// HasStep checks if Step has been set in PatternFlowArpHardwareLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowArpHardwareLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowArpHardwareLengthCounter
	SetCount(value uint32) PatternFlowArpHardwareLengthCounter
	// HasCount checks if Count has been set in PatternFlowArpHardwareLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowArpHardwareLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowArpHardwareLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowArpHardwareLengthCounter object
func (obj *patternFlowArpHardwareLengthCounter) SetStart(value uint32) PatternFlowArpHardwareLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowArpHardwareLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowArpHardwareLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowArpHardwareLengthCounter object
func (obj *patternFlowArpHardwareLengthCounter) SetStep(value uint32) PatternFlowArpHardwareLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpHardwareLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpHardwareLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowArpHardwareLengthCounter object
func (obj *patternFlowArpHardwareLengthCounter) SetCount(value uint32) PatternFlowArpHardwareLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowArpHardwareLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareLengthCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowArpHardwareLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(6)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
