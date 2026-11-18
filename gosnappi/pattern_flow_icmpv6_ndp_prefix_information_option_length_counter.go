package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter *****
type patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	marshaller   marshalPatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter() PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter()
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

// PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter is integer counter pattern
type PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(4)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
