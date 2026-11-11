package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter *****
type patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	marshaller   marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) (PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter()
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

// PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter is ipv6 counter pattern
type PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	SetStart(value string) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	SetStep(value string) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) SetStart(value string) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) SetStep(value string) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) SetCount(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter.Step"))
		}

	}

}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionPrefixCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("::")
	}
	if obj.obj.Step == nil {
		obj.SetStep("::1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
