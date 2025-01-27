package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter *****
type patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	marshaller   marshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter() PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter {
	obj := patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter{obj: &otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) msg() *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
}

type marshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) (PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) (PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) Clone() (PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter()
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

// PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter is integer counter pattern
type PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	// validate validates PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) setDefault() {
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
