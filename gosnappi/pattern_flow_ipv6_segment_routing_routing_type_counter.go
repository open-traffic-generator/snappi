package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingRoutingTypeCounter *****
type patternFlowIpv6SegmentRoutingRoutingTypeCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingRoutingTypeCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingRoutingTypeCounter
}

func NewPatternFlowIpv6SegmentRoutingRoutingTypeCounter() PatternFlowIpv6SegmentRoutingRoutingTypeCounter {
	obj := patternFlowIpv6SegmentRoutingRoutingTypeCounter{obj: &otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) msg() *otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter) PatternFlowIpv6SegmentRoutingRoutingTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingRoutingTypeCounter struct {
	obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter
}

type marshalPatternFlowIpv6SegmentRoutingRoutingTypeCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingRoutingTypeCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingRoutingTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingRoutingTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingRoutingTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingRoutingTypeCounter struct {
	obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingRoutingTypeCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingRoutingTypeCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter) (PatternFlowIpv6SegmentRoutingRoutingTypeCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingRoutingTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingRoutingTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingRoutingTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingRoutingTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingRoutingTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingRoutingTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingRoutingTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingRoutingTypeCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingRoutingTypeCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter) (PatternFlowIpv6SegmentRoutingRoutingTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingRoutingTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingRoutingTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingRoutingTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingRoutingTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingRoutingTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingRoutingTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) Clone() (PatternFlowIpv6SegmentRoutingRoutingTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingRoutingTypeCounter()
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

// PatternFlowIpv6SegmentRoutingRoutingTypeCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingRoutingTypeCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingRoutingTypeCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingRoutingTypeCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingRoutingTypeCounter) PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingRoutingTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingRoutingTypeCounter
	// validate validates PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingRoutingTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingRoutingTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingRoutingTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingRoutingTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingRoutingTypeCounter object
func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingRoutingTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingRoutingTypeCounter object
func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingRoutingTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingRoutingTypeCounter object
func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingRoutingTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingRoutingTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingRoutingTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingRoutingTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingRoutingTypeCounter) setDefault() {
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
