package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CPDUErrorIndexCounter *****
type patternFlowSnmpv2CPDUErrorIndexCounter struct {
	validation
	obj          *otg.PatternFlowSnmpv2CPDUErrorIndexCounter
	marshaller   marshalPatternFlowSnmpv2CPDUErrorIndexCounter
	unMarshaller unMarshalPatternFlowSnmpv2CPDUErrorIndexCounter
}

func NewPatternFlowSnmpv2CPDUErrorIndexCounter() PatternFlowSnmpv2CPDUErrorIndexCounter {
	obj := patternFlowSnmpv2CPDUErrorIndexCounter{obj: &otg.PatternFlowSnmpv2CPDUErrorIndexCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) msg() *otg.PatternFlowSnmpv2CPDUErrorIndexCounter {
	return obj.obj
}

func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) setMsg(msg *otg.PatternFlowSnmpv2CPDUErrorIndexCounter) PatternFlowSnmpv2CPDUErrorIndexCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CPDUErrorIndexCounter struct {
	obj *patternFlowSnmpv2CPDUErrorIndexCounter
}

type marshalPatternFlowSnmpv2CPDUErrorIndexCounter interface {
	// ToProto marshals PatternFlowSnmpv2CPDUErrorIndexCounter to protobuf object *otg.PatternFlowSnmpv2CPDUErrorIndexCounter
	ToProto() (*otg.PatternFlowSnmpv2CPDUErrorIndexCounter, error)
	// ToPbText marshals PatternFlowSnmpv2CPDUErrorIndexCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CPDUErrorIndexCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CPDUErrorIndexCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowSnmpv2CPDUErrorIndexCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowSnmpv2CPDUErrorIndexCounter struct {
	obj *patternFlowSnmpv2CPDUErrorIndexCounter
}

type unMarshalPatternFlowSnmpv2CPDUErrorIndexCounter interface {
	// FromProto unmarshals PatternFlowSnmpv2CPDUErrorIndexCounter from protobuf object *otg.PatternFlowSnmpv2CPDUErrorIndexCounter
	FromProto(msg *otg.PatternFlowSnmpv2CPDUErrorIndexCounter) (PatternFlowSnmpv2CPDUErrorIndexCounter, error)
	// FromPbText unmarshals PatternFlowSnmpv2CPDUErrorIndexCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CPDUErrorIndexCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CPDUErrorIndexCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) Marshal() marshalPatternFlowSnmpv2CPDUErrorIndexCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CPDUErrorIndexCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) Unmarshal() unMarshalPatternFlowSnmpv2CPDUErrorIndexCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CPDUErrorIndexCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CPDUErrorIndexCounter) ToProto() (*otg.PatternFlowSnmpv2CPDUErrorIndexCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CPDUErrorIndexCounter) FromProto(msg *otg.PatternFlowSnmpv2CPDUErrorIndexCounter) (PatternFlowSnmpv2CPDUErrorIndexCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CPDUErrorIndexCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CPDUErrorIndexCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CPDUErrorIndexCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CPDUErrorIndexCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CPDUErrorIndexCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowSnmpv2CPDUErrorIndexCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CPDUErrorIndexCounter) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) Clone() (PatternFlowSnmpv2CPDUErrorIndexCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CPDUErrorIndexCounter()
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

// PatternFlowSnmpv2CPDUErrorIndexCounter is integer counter pattern
type PatternFlowSnmpv2CPDUErrorIndexCounter interface {
	Validation
	// msg marshals PatternFlowSnmpv2CPDUErrorIndexCounter to protobuf object *otg.PatternFlowSnmpv2CPDUErrorIndexCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CPDUErrorIndexCounter
	// setMsg unmarshals PatternFlowSnmpv2CPDUErrorIndexCounter from protobuf object *otg.PatternFlowSnmpv2CPDUErrorIndexCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CPDUErrorIndexCounter) PatternFlowSnmpv2CPDUErrorIndexCounter
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CPDUErrorIndexCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CPDUErrorIndexCounter
	// validate validates PatternFlowSnmpv2CPDUErrorIndexCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CPDUErrorIndexCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowSnmpv2CPDUErrorIndexCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowSnmpv2CPDUErrorIndexCounter
	SetStart(value uint32) PatternFlowSnmpv2CPDUErrorIndexCounter
	// HasStart checks if Start has been set in PatternFlowSnmpv2CPDUErrorIndexCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowSnmpv2CPDUErrorIndexCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowSnmpv2CPDUErrorIndexCounter
	SetStep(value uint32) PatternFlowSnmpv2CPDUErrorIndexCounter
	// HasStep checks if Step has been set in PatternFlowSnmpv2CPDUErrorIndexCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowSnmpv2CPDUErrorIndexCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowSnmpv2CPDUErrorIndexCounter
	SetCount(value uint32) PatternFlowSnmpv2CPDUErrorIndexCounter
	// HasCount checks if Count has been set in PatternFlowSnmpv2CPDUErrorIndexCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowSnmpv2CPDUErrorIndexCounter object
func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) SetStart(value uint32) PatternFlowSnmpv2CPDUErrorIndexCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowSnmpv2CPDUErrorIndexCounter object
func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) SetStep(value uint32) PatternFlowSnmpv2CPDUErrorIndexCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowSnmpv2CPDUErrorIndexCounter object
func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) SetCount(value uint32) PatternFlowSnmpv2CPDUErrorIndexCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowSnmpv2CPDUErrorIndexCounter) setDefault() {
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
