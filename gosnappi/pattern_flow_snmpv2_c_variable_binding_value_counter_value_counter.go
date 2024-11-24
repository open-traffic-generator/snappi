package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVariableBindingValueCounterValueCounter *****
type patternFlowSnmpv2CVariableBindingValueCounterValueCounter struct {
	validation
	obj          *otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	marshaller   marshalPatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	unMarshaller unMarshalPatternFlowSnmpv2CVariableBindingValueCounterValueCounter
}

func NewPatternFlowSnmpv2CVariableBindingValueCounterValueCounter() PatternFlowSnmpv2CVariableBindingValueCounterValueCounter {
	obj := patternFlowSnmpv2CVariableBindingValueCounterValueCounter{obj: &otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) msg() *otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) setMsg(msg *otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter) PatternFlowSnmpv2CVariableBindingValueCounterValueCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVariableBindingValueCounterValueCounter struct {
	obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter
}

type marshalPatternFlowSnmpv2CVariableBindingValueCounterValueCounter interface {
	// ToProto marshals PatternFlowSnmpv2CVariableBindingValueCounterValueCounter to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter, error)
	// ToPbText marshals PatternFlowSnmpv2CVariableBindingValueCounterValueCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVariableBindingValueCounterValueCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVariableBindingValueCounterValueCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowSnmpv2CVariableBindingValueCounterValueCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowSnmpv2CVariableBindingValueCounterValueCounter struct {
	obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter
}

type unMarshalPatternFlowSnmpv2CVariableBindingValueCounterValueCounter interface {
	// FromProto unmarshals PatternFlowSnmpv2CVariableBindingValueCounterValueCounter from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter) (PatternFlowSnmpv2CVariableBindingValueCounterValueCounter, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVariableBindingValueCounterValueCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVariableBindingValueCounterValueCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVariableBindingValueCounterValueCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) Marshal() marshalPatternFlowSnmpv2CVariableBindingValueCounterValueCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVariableBindingValueCounterValueCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueCounterValueCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVariableBindingValueCounterValueCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueCounterValueCounter) ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueCounterValueCounter) FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter) (PatternFlowSnmpv2CVariableBindingValueCounterValueCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueCounterValueCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueCounterValueCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueCounterValueCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueCounterValueCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueCounterValueCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueCounterValueCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueCounterValueCounter) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) Clone() (PatternFlowSnmpv2CVariableBindingValueCounterValueCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVariableBindingValueCounterValueCounter()
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

// PatternFlowSnmpv2CVariableBindingValueCounterValueCounter is integer counter pattern
type PatternFlowSnmpv2CVariableBindingValueCounterValueCounter interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVariableBindingValueCounterValueCounter to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	// setMsg unmarshals PatternFlowSnmpv2CVariableBindingValueCounterValueCounter from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVariableBindingValueCounterValueCounter) PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	// validate validates PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVariableBindingValueCounterValueCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowSnmpv2CVariableBindingValueCounterValueCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	SetStart(value uint32) PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	// HasStart checks if Start has been set in PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowSnmpv2CVariableBindingValueCounterValueCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	SetStep(value uint32) PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	// HasStep checks if Step has been set in PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowSnmpv2CVariableBindingValueCounterValueCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	SetCount(value uint32) PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	// HasCount checks if Count has been set in PatternFlowSnmpv2CVariableBindingValueCounterValueCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowSnmpv2CVariableBindingValueCounterValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) SetStart(value uint32) PatternFlowSnmpv2CVariableBindingValueCounterValueCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowSnmpv2CVariableBindingValueCounterValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) SetStep(value uint32) PatternFlowSnmpv2CVariableBindingValueCounterValueCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowSnmpv2CVariableBindingValueCounterValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) SetCount(value uint32) PatternFlowSnmpv2CVariableBindingValueCounterValueCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowSnmpv2CVariableBindingValueCounterValueCounter) setDefault() {
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
