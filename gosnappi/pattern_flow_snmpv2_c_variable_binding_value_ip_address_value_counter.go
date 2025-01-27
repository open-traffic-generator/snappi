package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter *****
type patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter struct {
	validation
	obj          *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	marshaller   marshalPatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	unMarshaller unMarshalPatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
}

func NewPatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter() PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter {
	obj := patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter{obj: &otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) msg() *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) setMsg(msg *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter struct {
	obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
}

type marshalPatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter interface {
	// ToProto marshals PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter, error)
	// ToPbText marshals PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter struct {
	obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
}

type unMarshalPatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter interface {
	// FromProto unmarshals PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) (PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) Marshal() marshalPatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) (PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) Clone() (PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter()
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

// PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter is ipv4 counter pattern
type PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	// setMsg unmarshals PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	// validate validates PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	SetStart(value string) PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	// HasStart checks if Start has been set in PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	HasStart() bool
	// Step returns string, set in PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	SetStep(value string) PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	// HasStep checks if Step has been set in PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	SetCount(value uint32) PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	// HasCount checks if Count has been set in PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) SetStart(value string) PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) SetStep(value string) PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter object
func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) SetCount(value uint32) PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowSnmpv2CVariableBindingValueIpAddressValueCounter.Step"))
		}

	}

}

func (obj *patternFlowSnmpv2CVariableBindingValueIpAddressValueCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("0.0.0.0")
	}
	if obj.obj.Step == nil {
		obj.SetStep("0.0.0.1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
