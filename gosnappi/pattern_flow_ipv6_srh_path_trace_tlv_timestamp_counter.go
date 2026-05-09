package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPathTraceTlvTimestampCounter *****
type patternFlowIpv6SRHPathTraceTlvTimestampCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	marshaller   marshalPatternFlowIpv6SRHPathTraceTlvTimestampCounter
	unMarshaller unMarshalPatternFlowIpv6SRHPathTraceTlvTimestampCounter
}

func NewPatternFlowIpv6SRHPathTraceTlvTimestampCounter() PatternFlowIpv6SRHPathTraceTlvTimestampCounter {
	obj := patternFlowIpv6SRHPathTraceTlvTimestampCounter{obj: &otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) msg() *otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) setMsg(msg *otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter) PatternFlowIpv6SRHPathTraceTlvTimestampCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPathTraceTlvTimestampCounter struct {
	obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter
}

type marshalPatternFlowIpv6SRHPathTraceTlvTimestampCounter interface {
	// ToProto marshals PatternFlowIpv6SRHPathTraceTlvTimestampCounter to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHPathTraceTlvTimestampCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPathTraceTlvTimestampCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPathTraceTlvTimestampCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPathTraceTlvTimestampCounter struct {
	obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter
}

type unMarshalPatternFlowIpv6SRHPathTraceTlvTimestampCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHPathTraceTlvTimestampCounter from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter) (PatternFlowIpv6SRHPathTraceTlvTimestampCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPathTraceTlvTimestampCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPathTraceTlvTimestampCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPathTraceTlvTimestampCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) Marshal() marshalPatternFlowIpv6SRHPathTraceTlvTimestampCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPathTraceTlvTimestampCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvTimestampCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPathTraceTlvTimestampCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvTimestampCounter) ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvTimestampCounter) FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter) (PatternFlowIpv6SRHPathTraceTlvTimestampCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvTimestampCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvTimestampCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvTimestampCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvTimestampCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvTimestampCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvTimestampCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) Clone() (PatternFlowIpv6SRHPathTraceTlvTimestampCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPathTraceTlvTimestampCounter()
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

// PatternFlowIpv6SRHPathTraceTlvTimestampCounter is integer counter pattern
type PatternFlowIpv6SRHPathTraceTlvTimestampCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPathTraceTlvTimestampCounter to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	// setMsg unmarshals PatternFlowIpv6SRHPathTraceTlvTimestampCounter from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPathTraceTlvTimestampCounter) PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPathTraceTlvTimestampCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvTimestampCounter
	// validate validates PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPathTraceTlvTimestampCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHPathTraceTlvTimestampCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	SetStart(value uint32) PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHPathTraceTlvTimestampCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	SetStep(value uint32) PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHPathTraceTlvTimestampCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	SetCount(value uint32) PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvTimestampCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) SetStart(value uint32) PatternFlowIpv6SRHPathTraceTlvTimestampCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvTimestampCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) SetStep(value uint32) PatternFlowIpv6SRHPathTraceTlvTimestampCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvTimestampCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) SetCount(value uint32) PatternFlowIpv6SRHPathTraceTlvTimestampCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestampCounter) setDefault() {
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
