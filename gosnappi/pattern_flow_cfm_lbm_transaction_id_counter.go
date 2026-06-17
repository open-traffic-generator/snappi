package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmLbmTransactionIdCounter *****
type patternFlowCfmLbmTransactionIdCounter struct {
	validation
	obj          *otg.PatternFlowCfmLbmTransactionIdCounter
	marshaller   marshalPatternFlowCfmLbmTransactionIdCounter
	unMarshaller unMarshalPatternFlowCfmLbmTransactionIdCounter
}

func NewPatternFlowCfmLbmTransactionIdCounter() PatternFlowCfmLbmTransactionIdCounter {
	obj := patternFlowCfmLbmTransactionIdCounter{obj: &otg.PatternFlowCfmLbmTransactionIdCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmLbmTransactionIdCounter) msg() *otg.PatternFlowCfmLbmTransactionIdCounter {
	return obj.obj
}

func (obj *patternFlowCfmLbmTransactionIdCounter) setMsg(msg *otg.PatternFlowCfmLbmTransactionIdCounter) PatternFlowCfmLbmTransactionIdCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmLbmTransactionIdCounter struct {
	obj *patternFlowCfmLbmTransactionIdCounter
}

type marshalPatternFlowCfmLbmTransactionIdCounter interface {
	// ToProto marshals PatternFlowCfmLbmTransactionIdCounter to protobuf object *otg.PatternFlowCfmLbmTransactionIdCounter
	ToProto() (*otg.PatternFlowCfmLbmTransactionIdCounter, error)
	// ToPbText marshals PatternFlowCfmLbmTransactionIdCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmLbmTransactionIdCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmLbmTransactionIdCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmLbmTransactionIdCounter struct {
	obj *patternFlowCfmLbmTransactionIdCounter
}

type unMarshalPatternFlowCfmLbmTransactionIdCounter interface {
	// FromProto unmarshals PatternFlowCfmLbmTransactionIdCounter from protobuf object *otg.PatternFlowCfmLbmTransactionIdCounter
	FromProto(msg *otg.PatternFlowCfmLbmTransactionIdCounter) (PatternFlowCfmLbmTransactionIdCounter, error)
	// FromPbText unmarshals PatternFlowCfmLbmTransactionIdCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmLbmTransactionIdCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmLbmTransactionIdCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmLbmTransactionIdCounter) Marshal() marshalPatternFlowCfmLbmTransactionIdCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmLbmTransactionIdCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmLbmTransactionIdCounter) Unmarshal() unMarshalPatternFlowCfmLbmTransactionIdCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmLbmTransactionIdCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmLbmTransactionIdCounter) ToProto() (*otg.PatternFlowCfmLbmTransactionIdCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmLbmTransactionIdCounter) FromProto(msg *otg.PatternFlowCfmLbmTransactionIdCounter) (PatternFlowCfmLbmTransactionIdCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmLbmTransactionIdCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbmTransactionIdCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmLbmTransactionIdCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbmTransactionIdCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmLbmTransactionIdCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbmTransactionIdCounter) FromJson(value string) error {
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

func (obj *patternFlowCfmLbmTransactionIdCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmLbmTransactionIdCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmLbmTransactionIdCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmLbmTransactionIdCounter) Clone() (PatternFlowCfmLbmTransactionIdCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmLbmTransactionIdCounter()
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

// PatternFlowCfmLbmTransactionIdCounter is integer counter pattern
type PatternFlowCfmLbmTransactionIdCounter interface {
	Validation
	// msg marshals PatternFlowCfmLbmTransactionIdCounter to protobuf object *otg.PatternFlowCfmLbmTransactionIdCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmLbmTransactionIdCounter
	// setMsg unmarshals PatternFlowCfmLbmTransactionIdCounter from protobuf object *otg.PatternFlowCfmLbmTransactionIdCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmLbmTransactionIdCounter) PatternFlowCfmLbmTransactionIdCounter
	// provides marshal interface
	Marshal() marshalPatternFlowCfmLbmTransactionIdCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmLbmTransactionIdCounter
	// validate validates PatternFlowCfmLbmTransactionIdCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmLbmTransactionIdCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowCfmLbmTransactionIdCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowCfmLbmTransactionIdCounter
	SetStart(value uint32) PatternFlowCfmLbmTransactionIdCounter
	// HasStart checks if Start has been set in PatternFlowCfmLbmTransactionIdCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowCfmLbmTransactionIdCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowCfmLbmTransactionIdCounter
	SetStep(value uint32) PatternFlowCfmLbmTransactionIdCounter
	// HasStep checks if Step has been set in PatternFlowCfmLbmTransactionIdCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowCfmLbmTransactionIdCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowCfmLbmTransactionIdCounter
	SetCount(value uint32) PatternFlowCfmLbmTransactionIdCounter
	// HasCount checks if Count has been set in PatternFlowCfmLbmTransactionIdCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmLbmTransactionIdCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowCfmLbmTransactionIdCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowCfmLbmTransactionIdCounter object
func (obj *patternFlowCfmLbmTransactionIdCounter) SetStart(value uint32) PatternFlowCfmLbmTransactionIdCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmLbmTransactionIdCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowCfmLbmTransactionIdCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowCfmLbmTransactionIdCounter object
func (obj *patternFlowCfmLbmTransactionIdCounter) SetStep(value uint32) PatternFlowCfmLbmTransactionIdCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmLbmTransactionIdCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowCfmLbmTransactionIdCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowCfmLbmTransactionIdCounter object
func (obj *patternFlowCfmLbmTransactionIdCounter) SetCount(value uint32) PatternFlowCfmLbmTransactionIdCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowCfmLbmTransactionIdCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowCfmLbmTransactionIdCounter) setDefault() {
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
