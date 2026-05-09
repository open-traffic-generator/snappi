package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter *****
type patternFlowIpv6SRHOpaqueContainerTlvLengthCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	marshaller   marshalPatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	unMarshaller unMarshalPatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
}

func NewPatternFlowIpv6SRHOpaqueContainerTlvLengthCounter() PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter {
	obj := patternFlowIpv6SRHOpaqueContainerTlvLengthCounter{obj: &otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) msg() *otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) setMsg(msg *otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHOpaqueContainerTlvLengthCounter struct {
	obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter
}

type marshalPatternFlowIpv6SRHOpaqueContainerTlvLengthCounter interface {
	// ToProto marshals PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter to protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	ToProto() (*otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHOpaqueContainerTlvLengthCounter struct {
	obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter
}

type unMarshalPatternFlowIpv6SRHOpaqueContainerTlvLengthCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter from protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	FromProto(msg *otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) (PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) Marshal() marshalPatternFlowIpv6SRHOpaqueContainerTlvLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHOpaqueContainerTlvLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) Unmarshal() unMarshalPatternFlowIpv6SRHOpaqueContainerTlvLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHOpaqueContainerTlvLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) ToProto() (*otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) FromProto(msg *otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) (PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) Clone() (PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHOpaqueContainerTlvLengthCounter()
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

// PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter is integer counter pattern
type PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter to protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	// setMsg unmarshals PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter from protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter) PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	// validate validates PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	SetStart(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	SetStep(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	SetCount(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) SetStart(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) SetStep(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) SetCount(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHOpaqueContainerTlvLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvLengthCounter) setDefault() {
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
