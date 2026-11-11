package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter *****
type patternFlowIcmpv6NeighborSolicitationTargetAddressCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	marshaller   marshalPatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	unMarshaller unMarshalPatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
}

func NewPatternFlowIcmpv6NeighborSolicitationTargetAddressCounter() PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter {
	obj := patternFlowIcmpv6NeighborSolicitationTargetAddressCounter{obj: &otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) msg() *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) setMsg(msg *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NeighborSolicitationTargetAddressCounter struct {
	obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter
}

type marshalPatternFlowIcmpv6NeighborSolicitationTargetAddressCounter interface {
	// ToProto marshals PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NeighborSolicitationTargetAddressCounter struct {
	obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter
}

type unMarshalPatternFlowIcmpv6NeighborSolicitationTargetAddressCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) (PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) Marshal() marshalPatternFlowIcmpv6NeighborSolicitationTargetAddressCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NeighborSolicitationTargetAddressCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationTargetAddressCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NeighborSolicitationTargetAddressCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) (PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) Clone() (PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NeighborSolicitationTargetAddressCounter()
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

// PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter is ipv6 counter pattern
type PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	// setMsg unmarshals PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter) PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	// validate validates PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	SetStart(value string) PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	SetStep(value string) PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	SetCount(value uint32) PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter object
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) SetStart(value string) PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter object
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) SetStep(value string) PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter object
func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) SetCount(value uint32) PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIcmpv6NeighborSolicitationTargetAddressCounter.Step"))
		}

	}

}

func (obj *patternFlowIcmpv6NeighborSolicitationTargetAddressCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("::1")
	}
	if obj.obj.Step == nil {
		obj.SetStep("::1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
