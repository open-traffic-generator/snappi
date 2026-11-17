package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryGroupAddressCounter *****
type patternFlowIgmpv3MembershipQueryGroupAddressCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	marshaller   marshalPatternFlowIgmpv3MembershipQueryGroupAddressCounter
	unMarshaller unMarshalPatternFlowIgmpv3MembershipQueryGroupAddressCounter
}

func NewPatternFlowIgmpv3MembershipQueryGroupAddressCounter() PatternFlowIgmpv3MembershipQueryGroupAddressCounter {
	obj := patternFlowIgmpv3MembershipQueryGroupAddressCounter{obj: &otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) msg() *otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter) PatternFlowIgmpv3MembershipQueryGroupAddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryGroupAddressCounter struct {
	obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter
}

type marshalPatternFlowIgmpv3MembershipQueryGroupAddressCounter interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryGroupAddressCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryGroupAddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryGroupAddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryGroupAddressCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryGroupAddressCounter struct {
	obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter
}

type unMarshalPatternFlowIgmpv3MembershipQueryGroupAddressCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryGroupAddressCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter) (PatternFlowIgmpv3MembershipQueryGroupAddressCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryGroupAddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryGroupAddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryGroupAddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) Marshal() marshalPatternFlowIgmpv3MembershipQueryGroupAddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryGroupAddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryGroupAddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryGroupAddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryGroupAddressCounter) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryGroupAddressCounter) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter) (PatternFlowIgmpv3MembershipQueryGroupAddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryGroupAddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryGroupAddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryGroupAddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryGroupAddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryGroupAddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryGroupAddressCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) Clone() (PatternFlowIgmpv3MembershipQueryGroupAddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryGroupAddressCounter()
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

// PatternFlowIgmpv3MembershipQueryGroupAddressCounter is ipv4 counter pattern
type PatternFlowIgmpv3MembershipQueryGroupAddressCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryGroupAddressCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryGroupAddressCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryGroupAddressCounter) PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryGroupAddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryGroupAddressCounter
	// validate validates PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryGroupAddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIgmpv3MembershipQueryGroupAddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	SetStart(value string) PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIgmpv3MembershipQueryGroupAddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	SetStep(value string) PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3MembershipQueryGroupAddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	SetCount(value uint32) PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3MembershipQueryGroupAddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIgmpv3MembershipQueryGroupAddressCounter object
func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) SetStart(value string) PatternFlowIgmpv3MembershipQueryGroupAddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIgmpv3MembershipQueryGroupAddressCounter object
func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) SetStep(value string) PatternFlowIgmpv3MembershipQueryGroupAddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3MembershipQueryGroupAddressCounter object
func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) SetCount(value uint32) PatternFlowIgmpv3MembershipQueryGroupAddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv3MembershipQueryGroupAddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv3MembershipQueryGroupAddressCounter.Step"))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipQueryGroupAddressCounter) setDefault() {
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
