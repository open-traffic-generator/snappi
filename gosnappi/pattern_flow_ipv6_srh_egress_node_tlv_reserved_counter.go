package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHEgressNodeTlvReservedCounter *****
type patternFlowIpv6SRHEgressNodeTlvReservedCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	marshaller   marshalPatternFlowIpv6SRHEgressNodeTlvReservedCounter
	unMarshaller unMarshalPatternFlowIpv6SRHEgressNodeTlvReservedCounter
}

func NewPatternFlowIpv6SRHEgressNodeTlvReservedCounter() PatternFlowIpv6SRHEgressNodeTlvReservedCounter {
	obj := patternFlowIpv6SRHEgressNodeTlvReservedCounter{obj: &otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) msg() *otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) setMsg(msg *otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter) PatternFlowIpv6SRHEgressNodeTlvReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHEgressNodeTlvReservedCounter struct {
	obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter
}

type marshalPatternFlowIpv6SRHEgressNodeTlvReservedCounter interface {
	// ToProto marshals PatternFlowIpv6SRHEgressNodeTlvReservedCounter to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHEgressNodeTlvReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHEgressNodeTlvReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHEgressNodeTlvReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHEgressNodeTlvReservedCounter struct {
	obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter
}

type unMarshalPatternFlowIpv6SRHEgressNodeTlvReservedCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHEgressNodeTlvReservedCounter from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter) (PatternFlowIpv6SRHEgressNodeTlvReservedCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHEgressNodeTlvReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHEgressNodeTlvReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHEgressNodeTlvReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHEgressNodeTlvReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHEgressNodeTlvReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvReservedCounter) ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvReservedCounter) FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter) (PatternFlowIpv6SRHEgressNodeTlvReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) Clone() (PatternFlowIpv6SRHEgressNodeTlvReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHEgressNodeTlvReservedCounter()
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

// PatternFlowIpv6SRHEgressNodeTlvReservedCounter is integer counter pattern
type PatternFlowIpv6SRHEgressNodeTlvReservedCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHEgressNodeTlvReservedCounter to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	// setMsg unmarshals PatternFlowIpv6SRHEgressNodeTlvReservedCounter from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHEgressNodeTlvReservedCounter) PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvReservedCounter
	// validate validates PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHEgressNodeTlvReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	SetStart(value uint32) PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	SetStep(value uint32) PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	SetCount(value uint32) PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHEgressNodeTlvReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHEgressNodeTlvReservedCounter object
func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) SetStart(value uint32) PatternFlowIpv6SRHEgressNodeTlvReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHEgressNodeTlvReservedCounter object
func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) SetStep(value uint32) PatternFlowIpv6SRHEgressNodeTlvReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHEgressNodeTlvReservedCounter object
func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) SetCount(value uint32) PatternFlowIpv6SRHEgressNodeTlvReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvReservedCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvReservedCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvReservedCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SRHEgressNodeTlvReservedCounter) setDefault() {
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
