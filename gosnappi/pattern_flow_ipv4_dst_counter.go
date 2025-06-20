package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4DstCounter *****
type patternFlowIpv4DstCounter struct {
	validation
	obj          *otg.PatternFlowIpv4DstCounter
	marshaller   marshalPatternFlowIpv4DstCounter
	unMarshaller unMarshalPatternFlowIpv4DstCounter
}

func NewPatternFlowIpv4DstCounter() PatternFlowIpv4DstCounter {
	obj := patternFlowIpv4DstCounter{obj: &otg.PatternFlowIpv4DstCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4DstCounter) msg() *otg.PatternFlowIpv4DstCounter {
	return obj.obj
}

func (obj *patternFlowIpv4DstCounter) setMsg(msg *otg.PatternFlowIpv4DstCounter) PatternFlowIpv4DstCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4DstCounter struct {
	obj *patternFlowIpv4DstCounter
}

type marshalPatternFlowIpv4DstCounter interface {
	// ToProto marshals PatternFlowIpv4DstCounter to protobuf object *otg.PatternFlowIpv4DstCounter
	ToProto() (*otg.PatternFlowIpv4DstCounter, error)
	// ToPbText marshals PatternFlowIpv4DstCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4DstCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4DstCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4DstCounter struct {
	obj *patternFlowIpv4DstCounter
}

type unMarshalPatternFlowIpv4DstCounter interface {
	// FromProto unmarshals PatternFlowIpv4DstCounter from protobuf object *otg.PatternFlowIpv4DstCounter
	FromProto(msg *otg.PatternFlowIpv4DstCounter) (PatternFlowIpv4DstCounter, error)
	// FromPbText unmarshals PatternFlowIpv4DstCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4DstCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4DstCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4DstCounter) Marshal() marshalPatternFlowIpv4DstCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4DstCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4DstCounter) Unmarshal() unMarshalPatternFlowIpv4DstCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4DstCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4DstCounter) ToProto() (*otg.PatternFlowIpv4DstCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4DstCounter) FromProto(msg *otg.PatternFlowIpv4DstCounter) (PatternFlowIpv4DstCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4DstCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DstCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4DstCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DstCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4DstCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DstCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4DstCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DstCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DstCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4DstCounter) Clone() (PatternFlowIpv4DstCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4DstCounter()
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

// PatternFlowIpv4DstCounter is ipv4 counter pattern
type PatternFlowIpv4DstCounter interface {
	Validation
	// msg marshals PatternFlowIpv4DstCounter to protobuf object *otg.PatternFlowIpv4DstCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4DstCounter
	// setMsg unmarshals PatternFlowIpv4DstCounter from protobuf object *otg.PatternFlowIpv4DstCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4DstCounter) PatternFlowIpv4DstCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4DstCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4DstCounter
	// validate validates PatternFlowIpv4DstCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4DstCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv4DstCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv4DstCounter
	SetStart(value string) PatternFlowIpv4DstCounter
	// HasStart checks if Start has been set in PatternFlowIpv4DstCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv4DstCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv4DstCounter
	SetStep(value string) PatternFlowIpv4DstCounter
	// HasStep checks if Step has been set in PatternFlowIpv4DstCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4DstCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4DstCounter
	SetCount(value uint32) PatternFlowIpv4DstCounter
	// HasCount checks if Count has been set in PatternFlowIpv4DstCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv4DstCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv4DstCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv4DstCounter object
func (obj *patternFlowIpv4DstCounter) SetStart(value string) PatternFlowIpv4DstCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv4DstCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv4DstCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv4DstCounter object
func (obj *patternFlowIpv4DstCounter) SetStep(value string) PatternFlowIpv4DstCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4DstCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4DstCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4DstCounter object
func (obj *patternFlowIpv4DstCounter) SetCount(value uint32) PatternFlowIpv4DstCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4DstCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4DstCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4DstCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv4DstCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("0.0.0.0")
	}
	if obj.obj.Step == nil {
		obj.SetStep("0.0.0.1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
