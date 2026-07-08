package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmLbrTransactionIdCounter *****
type patternFlowCfmLbrTransactionIdCounter struct {
	validation
	obj          *otg.PatternFlowCfmLbrTransactionIdCounter
	marshaller   marshalPatternFlowCfmLbrTransactionIdCounter
	unMarshaller unMarshalPatternFlowCfmLbrTransactionIdCounter
}

func NewPatternFlowCfmLbrTransactionIdCounter() PatternFlowCfmLbrTransactionIdCounter {
	obj := patternFlowCfmLbrTransactionIdCounter{obj: &otg.PatternFlowCfmLbrTransactionIdCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmLbrTransactionIdCounter) msg() *otg.PatternFlowCfmLbrTransactionIdCounter {
	return obj.obj
}

func (obj *patternFlowCfmLbrTransactionIdCounter) setMsg(msg *otg.PatternFlowCfmLbrTransactionIdCounter) PatternFlowCfmLbrTransactionIdCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmLbrTransactionIdCounter struct {
	obj *patternFlowCfmLbrTransactionIdCounter
}

type marshalPatternFlowCfmLbrTransactionIdCounter interface {
	// ToProto marshals PatternFlowCfmLbrTransactionIdCounter to protobuf object *otg.PatternFlowCfmLbrTransactionIdCounter
	ToProto() (*otg.PatternFlowCfmLbrTransactionIdCounter, error)
	// ToPbText marshals PatternFlowCfmLbrTransactionIdCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmLbrTransactionIdCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmLbrTransactionIdCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmLbrTransactionIdCounter struct {
	obj *patternFlowCfmLbrTransactionIdCounter
}

type unMarshalPatternFlowCfmLbrTransactionIdCounter interface {
	// FromProto unmarshals PatternFlowCfmLbrTransactionIdCounter from protobuf object *otg.PatternFlowCfmLbrTransactionIdCounter
	FromProto(msg *otg.PatternFlowCfmLbrTransactionIdCounter) (PatternFlowCfmLbrTransactionIdCounter, error)
	// FromPbText unmarshals PatternFlowCfmLbrTransactionIdCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmLbrTransactionIdCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmLbrTransactionIdCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmLbrTransactionIdCounter) Marshal() marshalPatternFlowCfmLbrTransactionIdCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmLbrTransactionIdCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmLbrTransactionIdCounter) Unmarshal() unMarshalPatternFlowCfmLbrTransactionIdCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmLbrTransactionIdCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmLbrTransactionIdCounter) ToProto() (*otg.PatternFlowCfmLbrTransactionIdCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmLbrTransactionIdCounter) FromProto(msg *otg.PatternFlowCfmLbrTransactionIdCounter) (PatternFlowCfmLbrTransactionIdCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmLbrTransactionIdCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbrTransactionIdCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmLbrTransactionIdCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbrTransactionIdCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmLbrTransactionIdCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbrTransactionIdCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmLbrTransactionIdCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmLbrTransactionIdCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmLbrTransactionIdCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmLbrTransactionIdCounter) Clone() (PatternFlowCfmLbrTransactionIdCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmLbrTransactionIdCounter()
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

// PatternFlowCfmLbrTransactionIdCounter is integer counter pattern
type PatternFlowCfmLbrTransactionIdCounter interface {
	Validation
	// msg marshals PatternFlowCfmLbrTransactionIdCounter to protobuf object *otg.PatternFlowCfmLbrTransactionIdCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmLbrTransactionIdCounter
	// setMsg unmarshals PatternFlowCfmLbrTransactionIdCounter from protobuf object *otg.PatternFlowCfmLbrTransactionIdCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmLbrTransactionIdCounter) PatternFlowCfmLbrTransactionIdCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmLbrTransactionIdCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmLbrTransactionIdCounter
	// validate validates PatternFlowCfmLbrTransactionIdCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmLbrTransactionIdCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowCfmLbrTransactionIdCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowCfmLbrTransactionIdCounter
	SetStart(value uint32) PatternFlowCfmLbrTransactionIdCounter
	// HasStart checks if Start has been set in PatternFlowCfmLbrTransactionIdCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowCfmLbrTransactionIdCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowCfmLbrTransactionIdCounter
	SetStep(value uint32) PatternFlowCfmLbrTransactionIdCounter
	// HasStep checks if Step has been set in PatternFlowCfmLbrTransactionIdCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowCfmLbrTransactionIdCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowCfmLbrTransactionIdCounter
	SetCount(value uint32) PatternFlowCfmLbrTransactionIdCounter
	// HasCount checks if Count has been set in PatternFlowCfmLbrTransactionIdCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmLbrTransactionIdCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmLbrTransactionIdCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowCfmLbrTransactionIdCounter object
func (obj *patternFlowCfmLbrTransactionIdCounter) SetStart(value uint32) PatternFlowCfmLbrTransactionIdCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmLbrTransactionIdCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmLbrTransactionIdCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowCfmLbrTransactionIdCounter object
func (obj *patternFlowCfmLbrTransactionIdCounter) SetStep(value uint32) PatternFlowCfmLbrTransactionIdCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmLbrTransactionIdCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmLbrTransactionIdCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowCfmLbrTransactionIdCounter object
func (obj *patternFlowCfmLbrTransactionIdCounter) SetCount(value uint32) PatternFlowCfmLbrTransactionIdCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmLbrTransactionIdCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowCfmLbrTransactionIdCounter) setDefault() {
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
