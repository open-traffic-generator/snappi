package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter *****
type patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter struct {
	validation
	obj          *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	marshaller   marshalPatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	unMarshaller unMarshalPatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
}

func NewPatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter() PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter {
	obj := patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter{obj: &otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) msg() *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter {
	return obj.obj
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) setMsg(msg *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter struct {
	obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
}

type marshalPatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter interface {
	// ToProto marshals PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter to protobuf object *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	ToProto() (*otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter, error)
	// ToPbText marshals PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter struct {
	obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
}

type unMarshalPatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter interface {
	// FromProto unmarshals PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter from protobuf object *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	FromProto(msg *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) (PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter, error)
	// FromPbText unmarshals PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) Marshal() marshalPatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) Unmarshal() unMarshalPatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) ToProto() (*otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) FromProto(msg *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) (PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) Clone() (PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter()
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

// PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter is integer counter pattern
type PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter interface {
	Validation
	// msg marshals PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter to protobuf object *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	// setMsg unmarshals PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter from protobuf object *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	// validate validates PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	SetStart(value uint32) PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	// HasStart checks if Start has been set in PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	SetStep(value uint32) PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	// HasStep checks if Step has been set in PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	SetCount(value uint32) PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	// HasCount checks if Count has been set in PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter object
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) SetStart(value uint32) PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter object
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) SetStep(value uint32) PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter object
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) SetCount(value uint32) PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) setDefault() {
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
