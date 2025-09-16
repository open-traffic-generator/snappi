package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6VersionCounter *****
type patternFlowIpv6VersionCounter struct {
	validation
	obj          *otg.PatternFlowIpv6VersionCounter
	marshaller   marshalPatternFlowIpv6VersionCounter
	unMarshaller unMarshalPatternFlowIpv6VersionCounter
}

func NewPatternFlowIpv6VersionCounter() PatternFlowIpv6VersionCounter {
	obj := patternFlowIpv6VersionCounter{obj: &otg.PatternFlowIpv6VersionCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6VersionCounter) msg() *otg.PatternFlowIpv6VersionCounter {
	return obj.obj
}

func (obj *patternFlowIpv6VersionCounter) setMsg(msg *otg.PatternFlowIpv6VersionCounter) PatternFlowIpv6VersionCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6VersionCounter struct {
	obj *patternFlowIpv6VersionCounter
}

type marshalPatternFlowIpv6VersionCounter interface {
	// ToProto marshals PatternFlowIpv6VersionCounter to protobuf object *otg.PatternFlowIpv6VersionCounter
	ToProto() (*otg.PatternFlowIpv6VersionCounter, error)
	// ToPbText marshals PatternFlowIpv6VersionCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6VersionCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6VersionCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6VersionCounter struct {
	obj *patternFlowIpv6VersionCounter
}

type unMarshalPatternFlowIpv6VersionCounter interface {
	// FromProto unmarshals PatternFlowIpv6VersionCounter from protobuf object *otg.PatternFlowIpv6VersionCounter
	FromProto(msg *otg.PatternFlowIpv6VersionCounter) (PatternFlowIpv6VersionCounter, error)
	// FromPbText unmarshals PatternFlowIpv6VersionCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6VersionCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6VersionCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6VersionCounter) Marshal() marshalPatternFlowIpv6VersionCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6VersionCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6VersionCounter) Unmarshal() unMarshalPatternFlowIpv6VersionCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6VersionCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6VersionCounter) ToProto() (*otg.PatternFlowIpv6VersionCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6VersionCounter) FromProto(msg *otg.PatternFlowIpv6VersionCounter) (PatternFlowIpv6VersionCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6VersionCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6VersionCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6VersionCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6VersionCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6VersionCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6VersionCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6VersionCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6VersionCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6VersionCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6VersionCounter) Clone() (PatternFlowIpv6VersionCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6VersionCounter()
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

// PatternFlowIpv6VersionCounter is integer counter pattern
type PatternFlowIpv6VersionCounter interface {
	Validation
	// msg marshals PatternFlowIpv6VersionCounter to protobuf object *otg.PatternFlowIpv6VersionCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6VersionCounter
	// setMsg unmarshals PatternFlowIpv6VersionCounter from protobuf object *otg.PatternFlowIpv6VersionCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6VersionCounter) PatternFlowIpv6VersionCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6VersionCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6VersionCounter
	// validate validates PatternFlowIpv6VersionCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6VersionCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6VersionCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6VersionCounter
	SetStart(value uint32) PatternFlowIpv6VersionCounter
	// HasStart checks if Start has been set in PatternFlowIpv6VersionCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6VersionCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6VersionCounter
	SetStep(value uint32) PatternFlowIpv6VersionCounter
	// HasStep checks if Step has been set in PatternFlowIpv6VersionCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6VersionCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6VersionCounter
	SetCount(value uint32) PatternFlowIpv6VersionCounter
	// HasCount checks if Count has been set in PatternFlowIpv6VersionCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6VersionCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6VersionCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6VersionCounter object
func (obj *patternFlowIpv6VersionCounter) SetStart(value uint32) PatternFlowIpv6VersionCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6VersionCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6VersionCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6VersionCounter object
func (obj *patternFlowIpv6VersionCounter) SetStep(value uint32) PatternFlowIpv6VersionCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6VersionCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6VersionCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6VersionCounter object
func (obj *patternFlowIpv6VersionCounter) SetCount(value uint32) PatternFlowIpv6VersionCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6VersionCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6VersionCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6VersionCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6VersionCounter.Count <= 15 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6VersionCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(6)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
