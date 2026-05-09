package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHIngressNodeTlvValueCounter *****
type patternFlowIpv6SRHIngressNodeTlvValueCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter
	marshaller   marshalPatternFlowIpv6SRHIngressNodeTlvValueCounter
	unMarshaller unMarshalPatternFlowIpv6SRHIngressNodeTlvValueCounter
}

func NewPatternFlowIpv6SRHIngressNodeTlvValueCounter() PatternFlowIpv6SRHIngressNodeTlvValueCounter {
	obj := patternFlowIpv6SRHIngressNodeTlvValueCounter{obj: &otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) msg() *otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) setMsg(msg *otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter) PatternFlowIpv6SRHIngressNodeTlvValueCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHIngressNodeTlvValueCounter struct {
	obj *patternFlowIpv6SRHIngressNodeTlvValueCounter
}

type marshalPatternFlowIpv6SRHIngressNodeTlvValueCounter interface {
	// ToProto marshals PatternFlowIpv6SRHIngressNodeTlvValueCounter to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter
	ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHIngressNodeTlvValueCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHIngressNodeTlvValueCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHIngressNodeTlvValueCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHIngressNodeTlvValueCounter struct {
	obj *patternFlowIpv6SRHIngressNodeTlvValueCounter
}

type unMarshalPatternFlowIpv6SRHIngressNodeTlvValueCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHIngressNodeTlvValueCounter from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter
	FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter) (PatternFlowIpv6SRHIngressNodeTlvValueCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHIngressNodeTlvValueCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHIngressNodeTlvValueCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHIngressNodeTlvValueCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvValueCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHIngressNodeTlvValueCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvValueCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHIngressNodeTlvValueCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvValueCounter) ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvValueCounter) FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter) (PatternFlowIpv6SRHIngressNodeTlvValueCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvValueCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvValueCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvValueCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvValueCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvValueCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvValueCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) Clone() (PatternFlowIpv6SRHIngressNodeTlvValueCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHIngressNodeTlvValueCounter()
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

// PatternFlowIpv6SRHIngressNodeTlvValueCounter is ipv6 counter pattern
type PatternFlowIpv6SRHIngressNodeTlvValueCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHIngressNodeTlvValueCounter to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter
	// setMsg unmarshals PatternFlowIpv6SRHIngressNodeTlvValueCounter from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHIngressNodeTlvValueCounter) PatternFlowIpv6SRHIngressNodeTlvValueCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvValueCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvValueCounter
	// validate validates PatternFlowIpv6SRHIngressNodeTlvValueCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHIngressNodeTlvValueCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv6SRHIngressNodeTlvValueCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv6SRHIngressNodeTlvValueCounter
	SetStart(value string) PatternFlowIpv6SRHIngressNodeTlvValueCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHIngressNodeTlvValueCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv6SRHIngressNodeTlvValueCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv6SRHIngressNodeTlvValueCounter
	SetStep(value string) PatternFlowIpv6SRHIngressNodeTlvValueCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHIngressNodeTlvValueCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvValueCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvValueCounter
	SetCount(value uint32) PatternFlowIpv6SRHIngressNodeTlvValueCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHIngressNodeTlvValueCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv6SRHIngressNodeTlvValueCounter object
func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) SetStart(value string) PatternFlowIpv6SRHIngressNodeTlvValueCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv6SRHIngressNodeTlvValueCounter object
func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) SetStep(value string) PatternFlowIpv6SRHIngressNodeTlvValueCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHIngressNodeTlvValueCounter object
func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) SetCount(value uint32) PatternFlowIpv6SRHIngressNodeTlvValueCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SRHIngressNodeTlvValueCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SRHIngressNodeTlvValueCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv6SRHIngressNodeTlvValueCounter) setDefault() {
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
