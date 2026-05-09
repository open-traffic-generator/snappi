package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPathTraceTlvPmDataCounter *****
type patternFlowIpv6SRHPathTraceTlvPmDataCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	marshaller   marshalPatternFlowIpv6SRHPathTraceTlvPmDataCounter
	unMarshaller unMarshalPatternFlowIpv6SRHPathTraceTlvPmDataCounter
}

func NewPatternFlowIpv6SRHPathTraceTlvPmDataCounter() PatternFlowIpv6SRHPathTraceTlvPmDataCounter {
	obj := patternFlowIpv6SRHPathTraceTlvPmDataCounter{obj: &otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) msg() *otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) setMsg(msg *otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter) PatternFlowIpv6SRHPathTraceTlvPmDataCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPathTraceTlvPmDataCounter struct {
	obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter
}

type marshalPatternFlowIpv6SRHPathTraceTlvPmDataCounter interface {
	// ToProto marshals PatternFlowIpv6SRHPathTraceTlvPmDataCounter to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHPathTraceTlvPmDataCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPathTraceTlvPmDataCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPathTraceTlvPmDataCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPathTraceTlvPmDataCounter struct {
	obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter
}

type unMarshalPatternFlowIpv6SRHPathTraceTlvPmDataCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHPathTraceTlvPmDataCounter from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter) (PatternFlowIpv6SRHPathTraceTlvPmDataCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPathTraceTlvPmDataCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPathTraceTlvPmDataCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPathTraceTlvPmDataCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) Marshal() marshalPatternFlowIpv6SRHPathTraceTlvPmDataCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPathTraceTlvPmDataCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvPmDataCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPathTraceTlvPmDataCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvPmDataCounter) ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvPmDataCounter) FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter) (PatternFlowIpv6SRHPathTraceTlvPmDataCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvPmDataCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvPmDataCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvPmDataCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvPmDataCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvPmDataCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvPmDataCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) Clone() (PatternFlowIpv6SRHPathTraceTlvPmDataCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPathTraceTlvPmDataCounter()
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

// PatternFlowIpv6SRHPathTraceTlvPmDataCounter is integer counter pattern
type PatternFlowIpv6SRHPathTraceTlvPmDataCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPathTraceTlvPmDataCounter to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	// setMsg unmarshals PatternFlowIpv6SRHPathTraceTlvPmDataCounter from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPathTraceTlvPmDataCounter) PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPathTraceTlvPmDataCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvPmDataCounter
	// validate validates PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPathTraceTlvPmDataCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHPathTraceTlvPmDataCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	SetStart(value uint32) PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHPathTraceTlvPmDataCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	SetStep(value uint32) PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHPathTraceTlvPmDataCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	SetCount(value uint32) PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvPmDataCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) SetStart(value uint32) PatternFlowIpv6SRHPathTraceTlvPmDataCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvPmDataCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) SetStep(value uint32) PatternFlowIpv6SRHPathTraceTlvPmDataCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvPmDataCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) SetCount(value uint32) PatternFlowIpv6SRHPathTraceTlvPmDataCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowIpv6SRHPathTraceTlvPmDataCounter) setDefault() {
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
