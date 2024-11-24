package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv1GroupAddressCounter *****
type patternFlowIgmpv1GroupAddressCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv1GroupAddressCounter
	marshaller   marshalPatternFlowIgmpv1GroupAddressCounter
	unMarshaller unMarshalPatternFlowIgmpv1GroupAddressCounter
}

func NewPatternFlowIgmpv1GroupAddressCounter() PatternFlowIgmpv1GroupAddressCounter {
	obj := patternFlowIgmpv1GroupAddressCounter{obj: &otg.PatternFlowIgmpv1GroupAddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv1GroupAddressCounter) msg() *otg.PatternFlowIgmpv1GroupAddressCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv1GroupAddressCounter) setMsg(msg *otg.PatternFlowIgmpv1GroupAddressCounter) PatternFlowIgmpv1GroupAddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv1GroupAddressCounter struct {
	obj *patternFlowIgmpv1GroupAddressCounter
}

type marshalPatternFlowIgmpv1GroupAddressCounter interface {
	// ToProto marshals PatternFlowIgmpv1GroupAddressCounter to protobuf object *otg.PatternFlowIgmpv1GroupAddressCounter
	ToProto() (*otg.PatternFlowIgmpv1GroupAddressCounter, error)
	// ToPbText marshals PatternFlowIgmpv1GroupAddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv1GroupAddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv1GroupAddressCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIgmpv1GroupAddressCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIgmpv1GroupAddressCounter struct {
	obj *patternFlowIgmpv1GroupAddressCounter
}

type unMarshalPatternFlowIgmpv1GroupAddressCounter interface {
	// FromProto unmarshals PatternFlowIgmpv1GroupAddressCounter from protobuf object *otg.PatternFlowIgmpv1GroupAddressCounter
	FromProto(msg *otg.PatternFlowIgmpv1GroupAddressCounter) (PatternFlowIgmpv1GroupAddressCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv1GroupAddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv1GroupAddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv1GroupAddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv1GroupAddressCounter) Marshal() marshalPatternFlowIgmpv1GroupAddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv1GroupAddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv1GroupAddressCounter) Unmarshal() unMarshalPatternFlowIgmpv1GroupAddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv1GroupAddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv1GroupAddressCounter) ToProto() (*otg.PatternFlowIgmpv1GroupAddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv1GroupAddressCounter) FromProto(msg *otg.PatternFlowIgmpv1GroupAddressCounter) (PatternFlowIgmpv1GroupAddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv1GroupAddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1GroupAddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv1GroupAddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1GroupAddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv1GroupAddressCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIgmpv1GroupAddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1GroupAddressCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv1GroupAddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1GroupAddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1GroupAddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv1GroupAddressCounter) Clone() (PatternFlowIgmpv1GroupAddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv1GroupAddressCounter()
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

// PatternFlowIgmpv1GroupAddressCounter is ipv4 counter pattern
type PatternFlowIgmpv1GroupAddressCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv1GroupAddressCounter to protobuf object *otg.PatternFlowIgmpv1GroupAddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv1GroupAddressCounter
	// setMsg unmarshals PatternFlowIgmpv1GroupAddressCounter from protobuf object *otg.PatternFlowIgmpv1GroupAddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv1GroupAddressCounter) PatternFlowIgmpv1GroupAddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv1GroupAddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv1GroupAddressCounter
	// validate validates PatternFlowIgmpv1GroupAddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv1GroupAddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIgmpv1GroupAddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIgmpv1GroupAddressCounter
	SetStart(value string) PatternFlowIgmpv1GroupAddressCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv1GroupAddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIgmpv1GroupAddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIgmpv1GroupAddressCounter
	SetStep(value string) PatternFlowIgmpv1GroupAddressCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv1GroupAddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv1GroupAddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv1GroupAddressCounter
	SetCount(value uint32) PatternFlowIgmpv1GroupAddressCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv1GroupAddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIgmpv1GroupAddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIgmpv1GroupAddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIgmpv1GroupAddressCounter object
func (obj *patternFlowIgmpv1GroupAddressCounter) SetStart(value string) PatternFlowIgmpv1GroupAddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIgmpv1GroupAddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIgmpv1GroupAddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIgmpv1GroupAddressCounter object
func (obj *patternFlowIgmpv1GroupAddressCounter) SetStep(value string) PatternFlowIgmpv1GroupAddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv1GroupAddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv1GroupAddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv1GroupAddressCounter object
func (obj *patternFlowIgmpv1GroupAddressCounter) SetCount(value uint32) PatternFlowIgmpv1GroupAddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv1GroupAddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv1GroupAddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv1GroupAddressCounter.Step"))
		}

	}

}

func (obj *patternFlowIgmpv1GroupAddressCounter) setDefault() {
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
