package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter *****
type patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	marshaller   marshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter() PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter()
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

// PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter is integer counter pattern
type PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) setDefault() {
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
