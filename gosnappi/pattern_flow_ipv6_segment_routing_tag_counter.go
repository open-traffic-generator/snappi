package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTagCounter *****
type patternFlowIpv6SegmentRoutingTagCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingTagCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingTagCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingTagCounter
}

func NewPatternFlowIpv6SegmentRoutingTagCounter() PatternFlowIpv6SegmentRoutingTagCounter {
	obj := patternFlowIpv6SegmentRoutingTagCounter{obj: &otg.PatternFlowIpv6SegmentRoutingTagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTagCounter) msg() *otg.PatternFlowIpv6SegmentRoutingTagCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTagCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTagCounter) PatternFlowIpv6SegmentRoutingTagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTagCounter struct {
	obj *patternFlowIpv6SegmentRoutingTagCounter
}

type marshalPatternFlowIpv6SegmentRoutingTagCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTagCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTagCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTagCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTagCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTagCounter struct {
	obj *patternFlowIpv6SegmentRoutingTagCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingTagCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTagCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTagCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTagCounter) (PatternFlowIpv6SegmentRoutingTagCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTagCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingTagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTagCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTagCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTagCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTagCounter) (PatternFlowIpv6SegmentRoutingTagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTagCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTagCounter) Clone() (PatternFlowIpv6SegmentRoutingTagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTagCounter()
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

// PatternFlowIpv6SegmentRoutingTagCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingTagCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTagCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingTagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTagCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTagCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingTagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTagCounter) PatternFlowIpv6SegmentRoutingTagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTagCounter
	// validate validates PatternFlowIpv6SegmentRoutingTagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingTagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTagCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingTagCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingTagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingTagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTagCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingTagCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingTagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingTagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTagCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingTagCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingTagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingTagCounter object
func (obj *patternFlowIpv6SegmentRoutingTagCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingTagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingTagCounter object
func (obj *patternFlowIpv6SegmentRoutingTagCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingTagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingTagCounter object
func (obj *patternFlowIpv6SegmentRoutingTagCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingTagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTagCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTagCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTagCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingTagCounter) setDefault() {
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
