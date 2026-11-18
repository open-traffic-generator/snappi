package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpMtuOptionReservedCounter *****
type patternFlowIcmpv6NdpMtuOptionReservedCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter
	marshaller   marshalPatternFlowIcmpv6NdpMtuOptionReservedCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpMtuOptionReservedCounter
}

func NewPatternFlowIcmpv6NdpMtuOptionReservedCounter() PatternFlowIcmpv6NdpMtuOptionReservedCounter {
	obj := patternFlowIcmpv6NdpMtuOptionReservedCounter{obj: &otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) msg() *otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter) PatternFlowIcmpv6NdpMtuOptionReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpMtuOptionReservedCounter struct {
	obj *patternFlowIcmpv6NdpMtuOptionReservedCounter
}

type marshalPatternFlowIcmpv6NdpMtuOptionReservedCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpMtuOptionReservedCounter to protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpMtuOptionReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpMtuOptionReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpMtuOptionReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpMtuOptionReservedCounter struct {
	obj *patternFlowIcmpv6NdpMtuOptionReservedCounter
}

type unMarshalPatternFlowIcmpv6NdpMtuOptionReservedCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpMtuOptionReservedCounter from protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter) (PatternFlowIcmpv6NdpMtuOptionReservedCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpMtuOptionReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpMtuOptionReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpMtuOptionReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) Marshal() marshalPatternFlowIcmpv6NdpMtuOptionReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpMtuOptionReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpMtuOptionReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpMtuOptionReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpMtuOptionReservedCounter) ToProto() (*otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionReservedCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter) (PatternFlowIcmpv6NdpMtuOptionReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpMtuOptionReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpMtuOptionReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpMtuOptionReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) Clone() (PatternFlowIcmpv6NdpMtuOptionReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpMtuOptionReservedCounter()
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

// PatternFlowIcmpv6NdpMtuOptionReservedCounter is integer counter pattern
type PatternFlowIcmpv6NdpMtuOptionReservedCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpMtuOptionReservedCounter to protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpMtuOptionReservedCounter from protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpMtuOptionReservedCounter) PatternFlowIcmpv6NdpMtuOptionReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpMtuOptionReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpMtuOptionReservedCounter
	// validate validates PatternFlowIcmpv6NdpMtuOptionReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpMtuOptionReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpMtuOptionReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionReservedCounter
	SetStart(value uint32) PatternFlowIcmpv6NdpMtuOptionReservedCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpMtuOptionReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpMtuOptionReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionReservedCounter
	SetStep(value uint32) PatternFlowIcmpv6NdpMtuOptionReservedCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpMtuOptionReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpMtuOptionReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionReservedCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpMtuOptionReservedCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpMtuOptionReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpMtuOptionReservedCounter object
func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) SetStart(value uint32) PatternFlowIcmpv6NdpMtuOptionReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpMtuOptionReservedCounter object
func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) SetStep(value uint32) PatternFlowIcmpv6NdpMtuOptionReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpMtuOptionReservedCounter object
func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) SetCount(value uint32) PatternFlowIcmpv6NdpMtuOptionReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpMtuOptionReservedCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpMtuOptionReservedCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpMtuOptionReservedCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6NdpMtuOptionReservedCounter) setDefault() {
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
