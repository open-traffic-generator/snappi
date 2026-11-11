package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter *****
type patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter struct {
	validation
	obj          *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	marshaller   marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	unMarshaller unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
}

func NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter() PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter {
	obj := patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter{obj: &otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) msg() *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) setMsg(msg *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter struct {
	obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
}

type marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter interface {
	// ToProto marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter to protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	ToProto() (*otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter, error)
	// ToPbText marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter struct {
	obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
}

type unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter from protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	FromProto(msg *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) (PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) Marshal() marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) Unmarshal() unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) ToProto() (*otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) FromProto(msg *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) (PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) Clone() (PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter()
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

// PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter is integer counter pattern
type PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter to protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	// setMsg unmarshals PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter from protobuf object *otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	// validate validates PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	SetStart(value uint32) PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	// HasStart checks if Start has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	SetStep(value uint32) PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	// HasStep checks if Step has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	SetCount(value uint32) PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	// HasCount checks if Count has been set in PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) SetStart(value uint32) PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) SetStep(value uint32) PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter object
func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) SetCount(value uint32) PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIcmpv6NdpLinkLayerAddressOptionLengthCounter) setDefault() {
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
