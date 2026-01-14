package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpOperationCounter *****
type patternFlowArpOperationCounter struct {
	validation
	obj          *otg.PatternFlowArpOperationCounter
	marshaller   marshalPatternFlowArpOperationCounter
	unMarshaller unMarshalPatternFlowArpOperationCounter
}

func NewPatternFlowArpOperationCounter() PatternFlowArpOperationCounter {
	obj := patternFlowArpOperationCounter{obj: &otg.PatternFlowArpOperationCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpOperationCounter) msg() *otg.PatternFlowArpOperationCounter {
	return obj.obj
}

func (obj *patternFlowArpOperationCounter) setMsg(msg *otg.PatternFlowArpOperationCounter) PatternFlowArpOperationCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpOperationCounter struct {
	obj *patternFlowArpOperationCounter
}

type marshalPatternFlowArpOperationCounter interface {
	// ToProto marshals PatternFlowArpOperationCounter to protobuf object *otg.PatternFlowArpOperationCounter
	ToProto() (*otg.PatternFlowArpOperationCounter, error)
	// ToPbText marshals PatternFlowArpOperationCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpOperationCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpOperationCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpOperationCounter struct {
	obj *patternFlowArpOperationCounter
}

type unMarshalPatternFlowArpOperationCounter interface {
	// FromProto unmarshals PatternFlowArpOperationCounter from protobuf object *otg.PatternFlowArpOperationCounter
	FromProto(msg *otg.PatternFlowArpOperationCounter) (PatternFlowArpOperationCounter, error)
	// FromPbText unmarshals PatternFlowArpOperationCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpOperationCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpOperationCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpOperationCounter) Marshal() marshalPatternFlowArpOperationCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpOperationCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpOperationCounter) Unmarshal() unMarshalPatternFlowArpOperationCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpOperationCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpOperationCounter) ToProto() (*otg.PatternFlowArpOperationCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpOperationCounter) FromProto(msg *otg.PatternFlowArpOperationCounter) (PatternFlowArpOperationCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpOperationCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpOperationCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpOperationCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpOperationCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpOperationCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpOperationCounter) FromJson(value string) error {
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

func (obj *patternFlowArpOperationCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpOperationCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpOperationCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpOperationCounter) Clone() (PatternFlowArpOperationCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpOperationCounter()
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

// PatternFlowArpOperationCounter is integer counter pattern
type PatternFlowArpOperationCounter interface {
	Validation
	// msg marshals PatternFlowArpOperationCounter to protobuf object *otg.PatternFlowArpOperationCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowArpOperationCounter
	// setMsg unmarshals PatternFlowArpOperationCounter from protobuf object *otg.PatternFlowArpOperationCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpOperationCounter) PatternFlowArpOperationCounter
	// provides marshal interface
	Marshal() marshalPatternFlowArpOperationCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpOperationCounter
	// validate validates PatternFlowArpOperationCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpOperationCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowArpOperationCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowArpOperationCounter
	SetStart(value uint32) PatternFlowArpOperationCounter
	// HasStart checks if Start has been set in PatternFlowArpOperationCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowArpOperationCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowArpOperationCounter
	SetStep(value uint32) PatternFlowArpOperationCounter
	// HasStep checks if Step has been set in PatternFlowArpOperationCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowArpOperationCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowArpOperationCounter
	SetCount(value uint32) PatternFlowArpOperationCounter
	// HasCount checks if Count has been set in PatternFlowArpOperationCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowArpOperationCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowArpOperationCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowArpOperationCounter object
func (obj *patternFlowArpOperationCounter) SetStart(value uint32) PatternFlowArpOperationCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowArpOperationCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowArpOperationCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowArpOperationCounter object
func (obj *patternFlowArpOperationCounter) SetStep(value uint32) PatternFlowArpOperationCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpOperationCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowArpOperationCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowArpOperationCounter object
func (obj *patternFlowArpOperationCounter) SetCount(value uint32) PatternFlowArpOperationCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowArpOperationCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpOperationCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpOperationCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpOperationCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowArpOperationCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
