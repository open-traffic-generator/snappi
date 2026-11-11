package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter *****
type patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	marshaller   marshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter() PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter()
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

// PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter is integer counter pattern
type PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) SetStart(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) SetStep(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) setDefault() {
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
