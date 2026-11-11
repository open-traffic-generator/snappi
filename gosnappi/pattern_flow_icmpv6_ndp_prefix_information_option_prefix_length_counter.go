package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter *****
type patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	marshaller   marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter()
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

// PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter is integer counter pattern
type PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(64)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
