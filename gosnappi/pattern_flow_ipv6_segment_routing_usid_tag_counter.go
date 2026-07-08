package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingUsidTagCounter *****
type patternFlowIpv6SegmentRoutingUsidTagCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SegmentRoutingUsidTagCounter
	marshaller   marshalPatternFlowIpv6SegmentRoutingUsidTagCounter
	unMarshaller unMarshalPatternFlowIpv6SegmentRoutingUsidTagCounter
}

func NewPatternFlowIpv6SegmentRoutingUsidTagCounter() PatternFlowIpv6SegmentRoutingUsidTagCounter {
	obj := patternFlowIpv6SegmentRoutingUsidTagCounter{obj: &otg.PatternFlowIpv6SegmentRoutingUsidTagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) msg() *otg.PatternFlowIpv6SegmentRoutingUsidTagCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingUsidTagCounter) PatternFlowIpv6SegmentRoutingUsidTagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingUsidTagCounter struct {
	obj *patternFlowIpv6SegmentRoutingUsidTagCounter
}

type marshalPatternFlowIpv6SegmentRoutingUsidTagCounter interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingUsidTagCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidTagCounter
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidTagCounter, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingUsidTagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingUsidTagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingUsidTagCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingUsidTagCounter struct {
	obj *patternFlowIpv6SegmentRoutingUsidTagCounter
}

type unMarshalPatternFlowIpv6SegmentRoutingUsidTagCounter interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingUsidTagCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidTagCounter
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidTagCounter) (PatternFlowIpv6SegmentRoutingUsidTagCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingUsidTagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingUsidTagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingUsidTagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) Marshal() marshalPatternFlowIpv6SegmentRoutingUsidTagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingUsidTagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidTagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingUsidTagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidTagCounter) ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidTagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidTagCounter) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidTagCounter) (PatternFlowIpv6SegmentRoutingUsidTagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidTagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidTagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidTagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidTagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidTagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidTagCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) Clone() (PatternFlowIpv6SegmentRoutingUsidTagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingUsidTagCounter()
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

// PatternFlowIpv6SegmentRoutingUsidTagCounter is integer counter pattern
type PatternFlowIpv6SegmentRoutingUsidTagCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingUsidTagCounter to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidTagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingUsidTagCounter
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingUsidTagCounter from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidTagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingUsidTagCounter) PatternFlowIpv6SegmentRoutingUsidTagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingUsidTagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidTagCounter
	// validate validates PatternFlowIpv6SegmentRoutingUsidTagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingUsidTagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SegmentRoutingUsidTagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidTagCounter
	SetStart(value uint32) PatternFlowIpv6SegmentRoutingUsidTagCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SegmentRoutingUsidTagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SegmentRoutingUsidTagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidTagCounter
	SetStep(value uint32) PatternFlowIpv6SegmentRoutingUsidTagCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SegmentRoutingUsidTagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SegmentRoutingUsidTagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidTagCounter
	SetCount(value uint32) PatternFlowIpv6SegmentRoutingUsidTagCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SegmentRoutingUsidTagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidTagCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) SetStart(value uint32) PatternFlowIpv6SegmentRoutingUsidTagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidTagCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) SetStep(value uint32) PatternFlowIpv6SegmentRoutingUsidTagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidTagCounter object
func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) SetCount(value uint32) PatternFlowIpv6SegmentRoutingUsidTagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidTagCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidTagCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidTagCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SegmentRoutingUsidTagCounter) setDefault() {
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
