package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SrcCounter *****
type patternFlowIpv6SrcCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SrcCounter
	marshaller   marshalPatternFlowIpv6SrcCounter
	unMarshaller unMarshalPatternFlowIpv6SrcCounter
}

func NewPatternFlowIpv6SrcCounter() PatternFlowIpv6SrcCounter {
	obj := patternFlowIpv6SrcCounter{obj: &otg.PatternFlowIpv6SrcCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SrcCounter) msg() *otg.PatternFlowIpv6SrcCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SrcCounter) setMsg(msg *otg.PatternFlowIpv6SrcCounter) PatternFlowIpv6SrcCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SrcCounter struct {
	obj *patternFlowIpv6SrcCounter
}

type marshalPatternFlowIpv6SrcCounter interface {
	// ToProto marshals PatternFlowIpv6SrcCounter to protobuf object *otg.PatternFlowIpv6SrcCounter
	ToProto() (*otg.PatternFlowIpv6SrcCounter, error)
	// ToPbText marshals PatternFlowIpv6SrcCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SrcCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SrcCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv6SrcCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv6SrcCounter struct {
	obj *patternFlowIpv6SrcCounter
}

type unMarshalPatternFlowIpv6SrcCounter interface {
	// FromProto unmarshals PatternFlowIpv6SrcCounter from protobuf object *otg.PatternFlowIpv6SrcCounter
	FromProto(msg *otg.PatternFlowIpv6SrcCounter) (PatternFlowIpv6SrcCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SrcCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SrcCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SrcCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SrcCounter) Marshal() marshalPatternFlowIpv6SrcCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SrcCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SrcCounter) Unmarshal() unMarshalPatternFlowIpv6SrcCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SrcCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SrcCounter) ToProto() (*otg.PatternFlowIpv6SrcCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SrcCounter) FromProto(msg *otg.PatternFlowIpv6SrcCounter) (PatternFlowIpv6SrcCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SrcCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SrcCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SrcCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SrcCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SrcCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv6SrcCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SrcCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SrcCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SrcCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SrcCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SrcCounter) Clone() (PatternFlowIpv6SrcCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SrcCounter()
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

// PatternFlowIpv6SrcCounter is ipv6 counter pattern
type PatternFlowIpv6SrcCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SrcCounter to protobuf object *otg.PatternFlowIpv6SrcCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SrcCounter
	// setMsg unmarshals PatternFlowIpv6SrcCounter from protobuf object *otg.PatternFlowIpv6SrcCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SrcCounter) PatternFlowIpv6SrcCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SrcCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SrcCounter
	// validate validates PatternFlowIpv6SrcCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SrcCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternFlowIpv6SrcCounter.
	Start() string
	// SetStart assigns string provided by user to PatternFlowIpv6SrcCounter
	SetStart(value string) PatternFlowIpv6SrcCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SrcCounter
	HasStart() bool
	// Step returns string, set in PatternFlowIpv6SrcCounter.
	Step() string
	// SetStep assigns string provided by user to PatternFlowIpv6SrcCounter
	SetStep(value string) PatternFlowIpv6SrcCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SrcCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SrcCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SrcCounter
	SetCount(value uint32) PatternFlowIpv6SrcCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SrcCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SrcCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternFlowIpv6SrcCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternFlowIpv6SrcCounter object
func (obj *patternFlowIpv6SrcCounter) SetStart(value string) PatternFlowIpv6SrcCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SrcCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternFlowIpv6SrcCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternFlowIpv6SrcCounter object
func (obj *patternFlowIpv6SrcCounter) SetStep(value string) PatternFlowIpv6SrcCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SrcCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SrcCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SrcCounter object
func (obj *patternFlowIpv6SrcCounter) SetCount(value uint32) PatternFlowIpv6SrcCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SrcCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SrcCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6SrcCounter.Step"))
		}

	}

}

func (obj *patternFlowIpv6SrcCounter) setDefault() {
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
