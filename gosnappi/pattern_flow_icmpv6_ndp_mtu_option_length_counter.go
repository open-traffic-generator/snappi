package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpMtuOptionLengthCounter *****
type patternFlowIcmpv6NdpMtuOptionLengthCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter
	marshaller   marshalPatternFlowIcmpv6NdpMtuOptionLengthCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpMtuOptionLengthCounter
}

func NewPatternFlowIcmpv6NdpMtuOptionLengthCounter() PatternFlowIcmpv6NdpMtuOptionLengthCounter {
	obj := patternFlowIcmpv6NdpMtuOptionLengthCounter{obj: &otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) msg() *otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter) PatternFlowIcmpv6NdpMtuOptionLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpMtuOptionLengthCounter struct {
	obj *patternFlowIcmpv6NdpMtuOptionLengthCounter
}

type marshalPatternFlowIcmpv6NdpMtuOptionLengthCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpMtuOptionLengthCounter to protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpMtuOptionLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpMtuOptionLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpMtuOptionLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpMtuOptionLengthCounter struct {
	obj *patternFlowIcmpv6NdpMtuOptionLengthCounter
}

type unMarshalPatternFlowIcmpv6NdpMtuOptionLengthCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpMtuOptionLengthCounter from protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter) (PatternFlowIcmpv6NdpMtuOptionLengthCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpMtuOptionLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpMtuOptionLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpMtuOptionLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) Marshal() marshalPatternFlowIcmpv6NdpMtuOptionLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpMtuOptionLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpMtuOptionLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpMtuOptionLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpMtuOptionLengthCounter) ToProto() (*otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionLengthCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter) (PatternFlowIcmpv6NdpMtuOptionLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpMtuOptionLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpMtuOptionLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpMtuOptionLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) Clone() (PatternFlowIcmpv6NdpMtuOptionLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpMtuOptionLengthCounter()
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

// PatternFlowIcmpv6NdpMtuOptionLengthCounter is integer counter pattern
type PatternFlowIcmpv6NdpMtuOptionLengthCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpMtuOptionLengthCounter to protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpMtuOptionLengthCounter from protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpMtuOptionLengthCounter) PatternFlowIcmpv6NdpMtuOptionLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpMtuOptionLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpMtuOptionLengthCounter
	// validate validates PatternFlowIcmpv6NdpMtuOptionLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpMtuOptionLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpMtuOptionLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionLengthCounter
	SetStart(value uint32) PatternFlowIcmpv6NdpMtuOptionLengthCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpMtuOptionLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpMtuOptionLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionLengthCounter
	SetStep(value uint32) PatternFlowIcmpv6NdpMtuOptionLengthCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpMtuOptionLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpMtuOptionLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionLengthCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpMtuOptionLengthCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpMtuOptionLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpMtuOptionLengthCounter object
func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) SetStart(value uint32) PatternFlowIcmpv6NdpMtuOptionLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpMtuOptionLengthCounter object
func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) SetStep(value uint32) PatternFlowIcmpv6NdpMtuOptionLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpMtuOptionLengthCounter object
func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) SetCount(value uint32) PatternFlowIcmpv6NdpMtuOptionLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpMtuOptionLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpMtuOptionLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpMtuOptionLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6NdpMtuOptionLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
