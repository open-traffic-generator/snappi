package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter *****
type patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter struct {
	validation
	obj          *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	marshaller   marshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	unMarshaller unMarshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
}

func NewPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter() PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter {
	obj := patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter{obj: &otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) msg() *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) setMsg(msg *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter struct {
	obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
}

type marshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter interface {
	// ToProto marshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter, error)
	// ToPbText marshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter struct {
	obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
}

type unMarshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter interface {
	// FromProto unmarshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) (PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) Marshal() marshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) (PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) Clone() (PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter()
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

// PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter is integer counter pattern
type PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	// setMsg unmarshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	// validate validates PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	SetStart(value uint32) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	// HasStart checks if Start has been set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	SetStep(value uint32) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	// HasStep checks if Step has been set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	SetCount(value uint32) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	// HasCount checks if Count has been set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) SetStart(value uint32) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) SetStep(value uint32) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) SetCount(value uint32) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) setDefault() {
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
