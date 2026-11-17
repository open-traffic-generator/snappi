package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3SourceAddressCounter *****
type patternFlowIgmpv3SourceAddressCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3SourceAddressCounter
	marshaller   marshalPatternFlowIgmpv3SourceAddressCounter
	unMarshaller unMarshalPatternFlowIgmpv3SourceAddressCounter
}

func NewPatternFlowIgmpv3SourceAddressCounter() PatternFlowIgmpv3SourceAddressCounter {
	obj := patternFlowIgmpv3SourceAddressCounter{obj: &otg.PatternFlowIgmpv3SourceAddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3SourceAddressCounter) msg() *otg.PatternFlowIgmpv3SourceAddressCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3SourceAddressCounter) setMsg(msg *otg.PatternFlowIgmpv3SourceAddressCounter) PatternFlowIgmpv3SourceAddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3SourceAddressCounter struct {
	obj *patternFlowIgmpv3SourceAddressCounter
}

type marshalPatternFlowIgmpv3SourceAddressCounter interface {
	// ToProto marshals PatternFlowIgmpv3SourceAddressCounter to protobuf object *otg.PatternFlowIgmpv3SourceAddressCounter
	ToProto() (*otg.PatternFlowIgmpv3SourceAddressCounter, error)
	// ToPbText marshals PatternFlowIgmpv3SourceAddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3SourceAddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3SourceAddressCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3SourceAddressCounter struct {
	obj *patternFlowIgmpv3SourceAddressCounter
}

type unMarshalPatternFlowIgmpv3SourceAddressCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3SourceAddressCounter from protobuf object *otg.PatternFlowIgmpv3SourceAddressCounter
	FromProto(msg *otg.PatternFlowIgmpv3SourceAddressCounter) (PatternFlowIgmpv3SourceAddressCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3SourceAddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3SourceAddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3SourceAddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3SourceAddressCounter) Marshal() marshalPatternFlowIgmpv3SourceAddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3SourceAddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3SourceAddressCounter) Unmarshal() unMarshalPatternFlowIgmpv3SourceAddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3SourceAddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3SourceAddressCounter) ToProto() (*otg.PatternFlowIgmpv3SourceAddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3SourceAddressCounter) FromProto(msg *otg.PatternFlowIgmpv3SourceAddressCounter) (PatternFlowIgmpv3SourceAddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3SourceAddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3SourceAddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3SourceAddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3SourceAddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3SourceAddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3SourceAddressCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3SourceAddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3SourceAddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3SourceAddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3SourceAddressCounter) Clone() (PatternFlowIgmpv3SourceAddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3SourceAddressCounter()
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

// PatternFlowIgmpv3SourceAddressCounter is ipv4 counter pattern
type PatternFlowIgmpv3SourceAddressCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3SourceAddressCounter to protobuf object *otg.PatternFlowIgmpv3SourceAddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3SourceAddressCounter
	// setMsg unmarshals PatternFlowIgmpv3SourceAddressCounter from protobuf object *otg.PatternFlowIgmpv3SourceAddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3SourceAddressCounter) PatternFlowIgmpv3SourceAddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3SourceAddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3SourceAddressCounter
	// validate validates PatternFlowIgmpv3SourceAddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3SourceAddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIgmpv3SourceAddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIgmpv3SourceAddressCounter
	SetStart(value string) PatternFlowIgmpv3SourceAddressCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3SourceAddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIgmpv3SourceAddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIgmpv3SourceAddressCounter
	SetStep(value string) PatternFlowIgmpv3SourceAddressCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3SourceAddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3SourceAddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3SourceAddressCounter
	SetCount(value uint32) PatternFlowIgmpv3SourceAddressCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3SourceAddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIgmpv3SourceAddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIgmpv3SourceAddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIgmpv3SourceAddressCounter object
func (obj *patternFlowIgmpv3SourceAddressCounter) SetStart(value string) PatternFlowIgmpv3SourceAddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIgmpv3SourceAddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIgmpv3SourceAddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIgmpv3SourceAddressCounter object
func (obj *patternFlowIgmpv3SourceAddressCounter) SetStep(value string) PatternFlowIgmpv3SourceAddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3SourceAddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3SourceAddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3SourceAddressCounter object
func (obj *patternFlowIgmpv3SourceAddressCounter) SetCount(value uint32) PatternFlowIgmpv3SourceAddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3SourceAddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv3SourceAddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv3SourceAddressCounter.Step"))
		}

	}

}

func (obj *patternFlowIgmpv3SourceAddressCounter) setDefault() {
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
