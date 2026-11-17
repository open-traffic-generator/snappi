package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter *****
type patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	marshaller   marshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter() PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter()
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

// PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter is integer counter pattern
type PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) setDefault() {
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
