package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter *****
type patternFlowSnmpv2CVariableBindingValueIntegerValueCounter struct {
	validation
	obj          *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	marshaller   marshalPatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	unMarshaller unMarshalPatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
}

func NewPatternFlowSnmpv2CVariableBindingValueIntegerValueCounter() PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter {
	obj := patternFlowSnmpv2CVariableBindingValueIntegerValueCounter{obj: &otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) msg() *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) setMsg(msg *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVariableBindingValueIntegerValueCounter struct {
	obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter
}

type marshalPatternFlowSnmpv2CVariableBindingValueIntegerValueCounter interface {
	// ToProto marshals PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter, error)
	// ToPbText marshals PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowSnmpv2CVariableBindingValueIntegerValueCounter struct {
	obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter
}

type unMarshalPatternFlowSnmpv2CVariableBindingValueIntegerValueCounter interface {
	// FromProto unmarshals PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) (PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) Marshal() marshalPatternFlowSnmpv2CVariableBindingValueIntegerValueCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVariableBindingValueIntegerValueCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueIntegerValueCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVariableBindingValueIntegerValueCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) (PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) Clone() (PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVariableBindingValueIntegerValueCounter()
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

// PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter is integer counter pattern
type PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	// setMsg unmarshals PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter) PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	// validate validates PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns int32, set in PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter.
	Start() int32
	// SetStart assigns int32 provided by user to PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	SetStart(value int32) PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	// HasStart checks if Start has been set in PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	HasStart() bool
	// Step returns int32, set in PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter.
	Step() int32
	// SetStep assigns int32 provided by user to PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	SetStep(value int32) PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	// HasStep checks if Step has been set in PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	HasStep() bool
	// Count returns int32, set in PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter.
	Count() int32
	// SetCount assigns int32 provided by user to PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	SetCount(value int32) PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	// HasCount checks if Count has been set in PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter
	HasCount() bool
}

// description is TBD
// Start returns a int32
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) Start() int32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a int32
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the int32 value in the PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) SetStart(value int32) PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a int32
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) Step() int32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a int32
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the int32 value in the PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) SetStep(value int32) PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a int32
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) Count() int32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a int32
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the int32 value in the PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) SetCount(value int32) PatternFlowSnmpv2CVariableBindingValueIntegerValueCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowSnmpv2CVariableBindingValueIntegerValueCounter) setDefault() {
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
