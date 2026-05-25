package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmCcmSequenceNumberCounter *****
type patternFlowCfmCcmSequenceNumberCounter struct {
	validation
	obj          *otg.PatternFlowCfmCcmSequenceNumberCounter
	marshaller   marshalPatternFlowCfmCcmSequenceNumberCounter
	unMarshaller unMarshalPatternFlowCfmCcmSequenceNumberCounter
}

func NewPatternFlowCfmCcmSequenceNumberCounter() PatternFlowCfmCcmSequenceNumberCounter {
	obj := patternFlowCfmCcmSequenceNumberCounter{obj: &otg.PatternFlowCfmCcmSequenceNumberCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmCcmSequenceNumberCounter) msg() *otg.PatternFlowCfmCcmSequenceNumberCounter {
	return obj.obj
}

func (obj *patternFlowCfmCcmSequenceNumberCounter) setMsg(msg *otg.PatternFlowCfmCcmSequenceNumberCounter) PatternFlowCfmCcmSequenceNumberCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmCcmSequenceNumberCounter struct {
	obj *patternFlowCfmCcmSequenceNumberCounter
}

type marshalPatternFlowCfmCcmSequenceNumberCounter interface {
	// ToProto marshals PatternFlowCfmCcmSequenceNumberCounter to protobuf object *otg.PatternFlowCfmCcmSequenceNumberCounter
	ToProto() (*otg.PatternFlowCfmCcmSequenceNumberCounter, error)
	// ToPbText marshals PatternFlowCfmCcmSequenceNumberCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmCcmSequenceNumberCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmCcmSequenceNumberCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmCcmSequenceNumberCounter struct {
	obj *patternFlowCfmCcmSequenceNumberCounter
}

type unMarshalPatternFlowCfmCcmSequenceNumberCounter interface {
	// FromProto unmarshals PatternFlowCfmCcmSequenceNumberCounter from protobuf object *otg.PatternFlowCfmCcmSequenceNumberCounter
	FromProto(msg *otg.PatternFlowCfmCcmSequenceNumberCounter) (PatternFlowCfmCcmSequenceNumberCounter, error)
	// FromPbText unmarshals PatternFlowCfmCcmSequenceNumberCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmCcmSequenceNumberCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmCcmSequenceNumberCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmCcmSequenceNumberCounter) Marshal() marshalPatternFlowCfmCcmSequenceNumberCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmCcmSequenceNumberCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmCcmSequenceNumberCounter) Unmarshal() unMarshalPatternFlowCfmCcmSequenceNumberCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmCcmSequenceNumberCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmCcmSequenceNumberCounter) ToProto() (*otg.PatternFlowCfmCcmSequenceNumberCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmCcmSequenceNumberCounter) FromProto(msg *otg.PatternFlowCfmCcmSequenceNumberCounter) (PatternFlowCfmCcmSequenceNumberCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmCcmSequenceNumberCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmSequenceNumberCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmCcmSequenceNumberCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmSequenceNumberCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmCcmSequenceNumberCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmSequenceNumberCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmCcmSequenceNumberCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmCcmSequenceNumberCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmCcmSequenceNumberCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmCcmSequenceNumberCounter) Clone() (PatternFlowCfmCcmSequenceNumberCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmCcmSequenceNumberCounter()
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

// PatternFlowCfmCcmSequenceNumberCounter is integer counter pattern
type PatternFlowCfmCcmSequenceNumberCounter interface {
	Validation
	// msg marshals PatternFlowCfmCcmSequenceNumberCounter to protobuf object *otg.PatternFlowCfmCcmSequenceNumberCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmCcmSequenceNumberCounter
	// setMsg unmarshals PatternFlowCfmCcmSequenceNumberCounter from protobuf object *otg.PatternFlowCfmCcmSequenceNumberCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmCcmSequenceNumberCounter) PatternFlowCfmCcmSequenceNumberCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmCcmSequenceNumberCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmCcmSequenceNumberCounter
	// validate validates PatternFlowCfmCcmSequenceNumberCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmCcmSequenceNumberCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowCfmCcmSequenceNumberCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowCfmCcmSequenceNumberCounter
	SetStart(value uint32) PatternFlowCfmCcmSequenceNumberCounter
	// HasStart checks if Start has been set in PatternFlowCfmCcmSequenceNumberCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowCfmCcmSequenceNumberCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowCfmCcmSequenceNumberCounter
	SetStep(value uint32) PatternFlowCfmCcmSequenceNumberCounter
	// HasStep checks if Step has been set in PatternFlowCfmCcmSequenceNumberCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowCfmCcmSequenceNumberCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowCfmCcmSequenceNumberCounter
	SetCount(value uint32) PatternFlowCfmCcmSequenceNumberCounter
	// HasCount checks if Count has been set in PatternFlowCfmCcmSequenceNumberCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmCcmSequenceNumberCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmCcmSequenceNumberCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowCfmCcmSequenceNumberCounter object
func (obj *patternFlowCfmCcmSequenceNumberCounter) SetStart(value uint32) PatternFlowCfmCcmSequenceNumberCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmCcmSequenceNumberCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmCcmSequenceNumberCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowCfmCcmSequenceNumberCounter object
func (obj *patternFlowCfmCcmSequenceNumberCounter) SetStep(value uint32) PatternFlowCfmCcmSequenceNumberCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmCcmSequenceNumberCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmCcmSequenceNumberCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowCfmCcmSequenceNumberCounter object
func (obj *patternFlowCfmCcmSequenceNumberCounter) SetCount(value uint32) PatternFlowCfmCcmSequenceNumberCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmCcmSequenceNumberCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowCfmCcmSequenceNumberCounter) setDefault() {
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
