package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHEgressNodeTlvValueCounter *****
type patternFlowIpv6SRHEgressNodeTlvValueCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter
	marshaller   marshalPatternFlowIpv6SRHEgressNodeTlvValueCounter
	unMarshaller unMarshalPatternFlowIpv6SRHEgressNodeTlvValueCounter
}

func NewPatternFlowIpv6SRHEgressNodeTlvValueCounter() PatternFlowIpv6SRHEgressNodeTlvValueCounter {
	obj := patternFlowIpv6SRHEgressNodeTlvValueCounter{obj: &otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) msg() *otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) setMsg(msg *otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter) PatternFlowIpv6SRHEgressNodeTlvValueCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHEgressNodeTlvValueCounter struct {
	obj *patternFlowIpv6SRHEgressNodeTlvValueCounter
}

type marshalPatternFlowIpv6SRHEgressNodeTlvValueCounter interface {
	// ToProto marshals PatternFlowIpv6SRHEgressNodeTlvValueCounter to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter
	ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHEgressNodeTlvValueCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHEgressNodeTlvValueCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHEgressNodeTlvValueCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHEgressNodeTlvValueCounter struct {
	obj *patternFlowIpv6SRHEgressNodeTlvValueCounter
}

type unMarshalPatternFlowIpv6SRHEgressNodeTlvValueCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHEgressNodeTlvValueCounter from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter
	FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter) (PatternFlowIpv6SRHEgressNodeTlvValueCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHEgressNodeTlvValueCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHEgressNodeTlvValueCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHEgressNodeTlvValueCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvValueCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHEgressNodeTlvValueCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvValueCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHEgressNodeTlvValueCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvValueCounter) ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvValueCounter) FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter) (PatternFlowIpv6SRHEgressNodeTlvValueCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvValueCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvValueCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvValueCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvValueCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvValueCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvValueCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) Clone() (PatternFlowIpv6SRHEgressNodeTlvValueCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHEgressNodeTlvValueCounter()
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

// PatternFlowIpv6SRHEgressNodeTlvValueCounter is ipv6 counter pattern
type PatternFlowIpv6SRHEgressNodeTlvValueCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHEgressNodeTlvValueCounter to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter
	// setMsg unmarshals PatternFlowIpv6SRHEgressNodeTlvValueCounter from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHEgressNodeTlvValueCounter) PatternFlowIpv6SRHEgressNodeTlvValueCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvValueCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvValueCounter
	// validate validates PatternFlowIpv6SRHEgressNodeTlvValueCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHEgressNodeTlvValueCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv6SRHEgressNodeTlvValueCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv6SRHEgressNodeTlvValueCounter
	SetStart(value string) PatternFlowIpv6SRHEgressNodeTlvValueCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHEgressNodeTlvValueCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv6SRHEgressNodeTlvValueCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv6SRHEgressNodeTlvValueCounter
	SetStep(value string) PatternFlowIpv6SRHEgressNodeTlvValueCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHEgressNodeTlvValueCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvValueCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvValueCounter
	SetCount(value uint32) PatternFlowIpv6SRHEgressNodeTlvValueCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHEgressNodeTlvValueCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv6SRHEgressNodeTlvValueCounter object
func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) SetStart(value string) PatternFlowIpv6SRHEgressNodeTlvValueCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv6SRHEgressNodeTlvValueCounter object
func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) SetStep(value string) PatternFlowIpv6SRHEgressNodeTlvValueCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHEgressNodeTlvValueCounter object
func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) SetCount(value uint32) PatternFlowIpv6SRHEgressNodeTlvValueCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SRHEgressNodeTlvValueCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SRHEgressNodeTlvValueCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv6SRHEgressNodeTlvValueCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("::0")
	}
	if obj.obj.Step == nil {
		obj.SetStep("::1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
