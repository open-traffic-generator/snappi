package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter *****
type patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter struct {
	validation
	obj          *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	marshaller   marshalPatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	unMarshaller unMarshalPatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
}

func NewPatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter() PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter {
	obj := patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter{obj: &otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) msg() *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) setMsg(msg *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter struct {
	obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
}

type marshalPatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter interface {
	// ToProto marshals PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter, error)
	// ToPbText marshals PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter struct {
	obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
}

type unMarshalPatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter interface {
	// FromProto unmarshals PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) (PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) Marshal() marshalPatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) (PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) Clone() (PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter()
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

// PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter is integer counter pattern
type PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	// setMsg unmarshals PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	// validate validates PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	SetStart(value uint32) PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	// HasStart checks if Start has been set in PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	SetStep(value uint32) PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	// HasStep checks if Step has been set in PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	SetCount(value uint32) PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	// HasCount checks if Count has been set in PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) SetStart(value uint32) PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) SetStep(value uint32) PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) SetCount(value uint32) PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) setDefault() {
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
