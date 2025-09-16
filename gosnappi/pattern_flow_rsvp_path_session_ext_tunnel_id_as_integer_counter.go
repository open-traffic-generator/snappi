package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter *****
type patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	marshaller   marshalPatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	unMarshaller unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
}

func NewPatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter() PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter {
	obj := patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter{obj: &otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) msg() *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) setMsg(msg *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter struct {
	obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
}

type marshalPatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter interface {
	// ToProto marshals PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter to protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	ToProto() (*otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter struct {
	obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
}

type unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter from protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	FromProto(msg *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) (PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) Marshal() marshalPatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) Unmarshal() unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) ToProto() (*otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) FromProto(msg *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) (PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) Clone() (PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter()
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

// PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter is integer counter pattern
type PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter to protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	// setMsg unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter from protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	// validate validates PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	SetStart(value uint32) PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	SetStep(value uint32) PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	SetCount(value uint32) PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) SetStart(value uint32) PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) SetStep(value uint32) PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) SetCount(value uint32) PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) setDefault() {
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
