package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter *****
type patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	marshaller   marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
}

func NewPatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter() PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter {
	obj := patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter{obj: &otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) msg() *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter struct {
	obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
}

type marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter to protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter struct {
	obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
}

type unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter from protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) (PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) Marshal() marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) ToProto() (*otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) (PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) Clone() (PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter()
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

// PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter is integer counter pattern
type PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter to protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter from protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	// validate validates PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	SetStart(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	SetStep(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) SetStart(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) SetStep(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) SetCount(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionLengthCounter) setDefault() {
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
