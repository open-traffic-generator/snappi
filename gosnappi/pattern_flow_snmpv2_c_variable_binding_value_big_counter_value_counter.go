package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter *****
type patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter struct {
	validation
	obj          *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	marshaller   marshalPatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	unMarshaller unMarshalPatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
}

func NewPatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter() PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter {
	obj := patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter{obj: &otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) msg() *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) setMsg(msg *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter struct {
	obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
}

type marshalPatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter interface {
	// ToProto marshals PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter, error)
	// ToPbText marshals PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter struct {
	obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
}

type unMarshalPatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter interface {
	// FromProto unmarshals PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) (PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) Marshal() marshalPatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) (PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) Clone() (PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter()
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

// PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter is integer counter pattern
type PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	// setMsg unmarshals PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	// validate validates PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint64, set in PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter.
	Start() uint64
	// SetStart assigns uint64 provided by user to PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	SetStart(value uint64) PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	// HasStart checks if Start has been set in PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	HasStart() bool
	// Step returns uint64, set in PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter.
	Step() uint64
	// SetStep assigns uint64 provided by user to PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	SetStep(value uint64) PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	// HasStep checks if Step has been set in PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	HasStep() bool
	// Count returns uint64, set in PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter.
	Count() uint64
	// SetCount assigns uint64 provided by user to PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	SetCount(value uint64) PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	// HasCount checks if Count has been set in PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint64
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) Start() uint64 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint64
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint64 value in the PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) SetStart(value uint64) PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint64
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) Step() uint64 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint64
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint64 value in the PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) SetStep(value uint64) PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint64
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) Count() uint64 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint64
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint64 value in the PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) SetCount(value uint64) PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) setDefault() {
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
