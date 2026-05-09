package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPathTraceTlvValueCounter *****
type patternFlowIpv6SRHPathTraceTlvValueCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHPathTraceTlvValueCounter
	marshaller   marshalPatternFlowIpv6SRHPathTraceTlvValueCounter
	unMarshaller unMarshalPatternFlowIpv6SRHPathTraceTlvValueCounter
}

func NewPatternFlowIpv6SRHPathTraceTlvValueCounter() PatternFlowIpv6SRHPathTraceTlvValueCounter {
	obj := patternFlowIpv6SRHPathTraceTlvValueCounter{obj: &otg.PatternFlowIpv6SRHPathTraceTlvValueCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) msg() *otg.PatternFlowIpv6SRHPathTraceTlvValueCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) setMsg(msg *otg.PatternFlowIpv6SRHPathTraceTlvValueCounter) PatternFlowIpv6SRHPathTraceTlvValueCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPathTraceTlvValueCounter struct {
	obj *patternFlowIpv6SRHPathTraceTlvValueCounter
}

type marshalPatternFlowIpv6SRHPathTraceTlvValueCounter interface {
	// ToProto marshals PatternFlowIpv6SRHPathTraceTlvValueCounter to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvValueCounter
	ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvValueCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHPathTraceTlvValueCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPathTraceTlvValueCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPathTraceTlvValueCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPathTraceTlvValueCounter struct {
	obj *patternFlowIpv6SRHPathTraceTlvValueCounter
}

type unMarshalPatternFlowIpv6SRHPathTraceTlvValueCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHPathTraceTlvValueCounter from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvValueCounter
	FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvValueCounter) (PatternFlowIpv6SRHPathTraceTlvValueCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPathTraceTlvValueCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPathTraceTlvValueCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPathTraceTlvValueCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) Marshal() marshalPatternFlowIpv6SRHPathTraceTlvValueCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPathTraceTlvValueCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvValueCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPathTraceTlvValueCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvValueCounter) ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvValueCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvValueCounter) FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvValueCounter) (PatternFlowIpv6SRHPathTraceTlvValueCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvValueCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvValueCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvValueCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvValueCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvValueCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvValueCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) Clone() (PatternFlowIpv6SRHPathTraceTlvValueCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPathTraceTlvValueCounter()
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

// PatternFlowIpv6SRHPathTraceTlvValueCounter is integer counter pattern
type PatternFlowIpv6SRHPathTraceTlvValueCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPathTraceTlvValueCounter to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvValueCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPathTraceTlvValueCounter
	// setMsg unmarshals PatternFlowIpv6SRHPathTraceTlvValueCounter from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvValueCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPathTraceTlvValueCounter) PatternFlowIpv6SRHPathTraceTlvValueCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPathTraceTlvValueCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvValueCounter
	// validate validates PatternFlowIpv6SRHPathTraceTlvValueCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPathTraceTlvValueCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHPathTraceTlvValueCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvValueCounter
	SetStart(value uint32) PatternFlowIpv6SRHPathTraceTlvValueCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHPathTraceTlvValueCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHPathTraceTlvValueCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvValueCounter
	SetStep(value uint32) PatternFlowIpv6SRHPathTraceTlvValueCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHPathTraceTlvValueCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHPathTraceTlvValueCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvValueCounter
	SetCount(value uint32) PatternFlowIpv6SRHPathTraceTlvValueCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHPathTraceTlvValueCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvValueCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) SetStart(value uint32) PatternFlowIpv6SRHPathTraceTlvValueCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvValueCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) SetStep(value uint32) PatternFlowIpv6SRHPathTraceTlvValueCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvValueCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) SetCount(value uint32) PatternFlowIpv6SRHPathTraceTlvValueCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowIpv6SRHPathTraceTlvValueCounter) setDefault() {
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
