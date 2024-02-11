package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6DstCounter *****
type patternFlowIpv6DstCounter struct {
	validation
	obj          *otg.PatternFlowIpv6DstCounter
	marshaller   marshalPatternFlowIpv6DstCounter
	unMarshaller unMarshalPatternFlowIpv6DstCounter
}

func NewPatternFlowIpv6DstCounter() PatternFlowIpv6DstCounter {
	obj := patternFlowIpv6DstCounter{obj: &otg.PatternFlowIpv6DstCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6DstCounter) msg() *otg.PatternFlowIpv6DstCounter {
	return obj.obj
}

func (obj *patternFlowIpv6DstCounter) setMsg(msg *otg.PatternFlowIpv6DstCounter) PatternFlowIpv6DstCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6DstCounter struct {
	obj *patternFlowIpv6DstCounter
}

type marshalPatternFlowIpv6DstCounter interface {
	// ToProto marshals PatternFlowIpv6DstCounter to protobuf object *otg.PatternFlowIpv6DstCounter
	ToProto() (*otg.PatternFlowIpv6DstCounter, error)
	// ToPbText marshals PatternFlowIpv6DstCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6DstCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6DstCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6DstCounter struct {
	obj *patternFlowIpv6DstCounter
}

type unMarshalPatternFlowIpv6DstCounter interface {
	// FromProto unmarshals PatternFlowIpv6DstCounter from protobuf object *otg.PatternFlowIpv6DstCounter
	FromProto(msg *otg.PatternFlowIpv6DstCounter) (PatternFlowIpv6DstCounter, error)
	// FromPbText unmarshals PatternFlowIpv6DstCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6DstCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6DstCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6DstCounter) Marshal() marshalPatternFlowIpv6DstCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6DstCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6DstCounter) Unmarshal() unMarshalPatternFlowIpv6DstCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6DstCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6DstCounter) ToProto() (*otg.PatternFlowIpv6DstCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6DstCounter) FromProto(msg *otg.PatternFlowIpv6DstCounter) (PatternFlowIpv6DstCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6DstCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6DstCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6DstCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6DstCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6DstCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6DstCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6DstCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6DstCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6DstCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6DstCounter) Clone() (PatternFlowIpv6DstCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6DstCounter()
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

// PatternFlowIpv6DstCounter is ipv6 counter pattern
type PatternFlowIpv6DstCounter interface {
	Validation
	// msg marshals PatternFlowIpv6DstCounter to protobuf object *otg.PatternFlowIpv6DstCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6DstCounter
	// setMsg unmarshals PatternFlowIpv6DstCounter from protobuf object *otg.PatternFlowIpv6DstCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6DstCounter) PatternFlowIpv6DstCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6DstCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6DstCounter
	// validate validates PatternFlowIpv6DstCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6DstCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv6DstCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv6DstCounter
	SetStart(value string) PatternFlowIpv6DstCounter
	// HasStart checks if Start has been set in PatternFlowIpv6DstCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv6DstCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv6DstCounter
	SetStep(value string) PatternFlowIpv6DstCounter
	// HasStep checks if Step has been set in PatternFlowIpv6DstCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6DstCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6DstCounter
	SetCount(value uint32) PatternFlowIpv6DstCounter
	// HasCount checks if Count has been set in PatternFlowIpv6DstCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6DstCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6DstCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv6DstCounter object
func (obj *patternFlowIpv6DstCounter) SetStart(value string) PatternFlowIpv6DstCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6DstCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6DstCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv6DstCounter object
func (obj *patternFlowIpv6DstCounter) SetStep(value string) PatternFlowIpv6DstCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6DstCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6DstCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6DstCounter object
func (obj *patternFlowIpv6DstCounter) SetCount(value uint32) PatternFlowIpv6DstCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6DstCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6DstCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6DstCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv6DstCounter) setDefault() {
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
