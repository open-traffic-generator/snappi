package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv2GroupAddressCounter *****
type patternFlowIgmpv2GroupAddressCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv2GroupAddressCounter
	marshaller   marshalPatternFlowIgmpv2GroupAddressCounter
	unMarshaller unMarshalPatternFlowIgmpv2GroupAddressCounter
}

func NewPatternFlowIgmpv2GroupAddressCounter() PatternFlowIgmpv2GroupAddressCounter {
	obj := patternFlowIgmpv2GroupAddressCounter{obj: &otg.PatternFlowIgmpv2GroupAddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv2GroupAddressCounter) msg() *otg.PatternFlowIgmpv2GroupAddressCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv2GroupAddressCounter) setMsg(msg *otg.PatternFlowIgmpv2GroupAddressCounter) PatternFlowIgmpv2GroupAddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv2GroupAddressCounter struct {
	obj *patternFlowIgmpv2GroupAddressCounter
}

type marshalPatternFlowIgmpv2GroupAddressCounter interface {
	// ToProto marshals PatternFlowIgmpv2GroupAddressCounter to protobuf object *otg.PatternFlowIgmpv2GroupAddressCounter
	ToProto() (*otg.PatternFlowIgmpv2GroupAddressCounter, error)
	// ToPbText marshals PatternFlowIgmpv2GroupAddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv2GroupAddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv2GroupAddressCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv2GroupAddressCounter struct {
	obj *patternFlowIgmpv2GroupAddressCounter
}

type unMarshalPatternFlowIgmpv2GroupAddressCounter interface {
	// FromProto unmarshals PatternFlowIgmpv2GroupAddressCounter from protobuf object *otg.PatternFlowIgmpv2GroupAddressCounter
	FromProto(msg *otg.PatternFlowIgmpv2GroupAddressCounter) (PatternFlowIgmpv2GroupAddressCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv2GroupAddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv2GroupAddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv2GroupAddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv2GroupAddressCounter) Marshal() marshalPatternFlowIgmpv2GroupAddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv2GroupAddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv2GroupAddressCounter) Unmarshal() unMarshalPatternFlowIgmpv2GroupAddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv2GroupAddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv2GroupAddressCounter) ToProto() (*otg.PatternFlowIgmpv2GroupAddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv2GroupAddressCounter) FromProto(msg *otg.PatternFlowIgmpv2GroupAddressCounter) (PatternFlowIgmpv2GroupAddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv2GroupAddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2GroupAddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv2GroupAddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2GroupAddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv2GroupAddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2GroupAddressCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv2GroupAddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2GroupAddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2GroupAddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv2GroupAddressCounter) Clone() (PatternFlowIgmpv2GroupAddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv2GroupAddressCounter()
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

// PatternFlowIgmpv2GroupAddressCounter is ipv4 counter pattern
type PatternFlowIgmpv2GroupAddressCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv2GroupAddressCounter to protobuf object *otg.PatternFlowIgmpv2GroupAddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv2GroupAddressCounter
	// setMsg unmarshals PatternFlowIgmpv2GroupAddressCounter from protobuf object *otg.PatternFlowIgmpv2GroupAddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv2GroupAddressCounter) PatternFlowIgmpv2GroupAddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv2GroupAddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv2GroupAddressCounter
	// validate validates PatternFlowIgmpv2GroupAddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv2GroupAddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIgmpv2GroupAddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIgmpv2GroupAddressCounter
	SetStart(value string) PatternFlowIgmpv2GroupAddressCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv2GroupAddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIgmpv2GroupAddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIgmpv2GroupAddressCounter
	SetStep(value string) PatternFlowIgmpv2GroupAddressCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv2GroupAddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv2GroupAddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv2GroupAddressCounter
	SetCount(value uint32) PatternFlowIgmpv2GroupAddressCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv2GroupAddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIgmpv2GroupAddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIgmpv2GroupAddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIgmpv2GroupAddressCounter object
func (obj *patternFlowIgmpv2GroupAddressCounter) SetStart(value string) PatternFlowIgmpv2GroupAddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIgmpv2GroupAddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIgmpv2GroupAddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIgmpv2GroupAddressCounter object
func (obj *patternFlowIgmpv2GroupAddressCounter) SetStep(value string) PatternFlowIgmpv2GroupAddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv2GroupAddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv2GroupAddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv2GroupAddressCounter object
func (obj *patternFlowIgmpv2GroupAddressCounter) SetCount(value uint32) PatternFlowIgmpv2GroupAddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv2GroupAddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv2GroupAddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv2GroupAddressCounter.Step"))
		}

	}

}

func (obj *patternFlowIgmpv2GroupAddressCounter) setDefault() {
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
