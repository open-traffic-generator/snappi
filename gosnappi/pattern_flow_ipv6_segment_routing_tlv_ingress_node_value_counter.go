package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter *****
type patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter() PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter {
	obj := patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter{obj: &otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) msg() *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
}

type marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter struct {
	obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) (PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) (PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) Clone() (PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter()
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

// PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter is ipv6 counter pattern
type PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	// validate validates PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	SetStart(value string) PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	SetStep(value string) PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) SetStart(value string) PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) SetStep(value string) PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter object
func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingTlvIngressNodeValueCounter) setDefault() {
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
