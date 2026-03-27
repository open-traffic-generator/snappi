package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4VersionCounter *****
type patternFlowIpv4VersionCounter struct {
	validation
	obj          *otg.PatternFlowIpv4VersionCounter
	marshaller   marshalPatternFlowIpv4VersionCounter
	unMarshaller unMarshalPatternFlowIpv4VersionCounter
}

func NewPatternFlowIpv4VersionCounter() PatternFlowIpv4VersionCounter {
	obj := patternFlowIpv4VersionCounter{obj: &otg.PatternFlowIpv4VersionCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4VersionCounter) msg() *otg.PatternFlowIpv4VersionCounter {
	return obj.obj
}

func (obj *patternFlowIpv4VersionCounter) setMsg(msg *otg.PatternFlowIpv4VersionCounter) PatternFlowIpv4VersionCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4VersionCounter struct {
	obj *patternFlowIpv4VersionCounter
}

type marshalPatternFlowIpv4VersionCounter interface {
	// ToProto marshals PatternFlowIpv4VersionCounter to protobuf object *otg.PatternFlowIpv4VersionCounter
	ToProto() (*otg.PatternFlowIpv4VersionCounter, error)
	// ToPbText marshals PatternFlowIpv4VersionCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4VersionCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4VersionCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4VersionCounter struct {
	obj *patternFlowIpv4VersionCounter
}

type unMarshalPatternFlowIpv4VersionCounter interface {
	// FromProto unmarshals PatternFlowIpv4VersionCounter from protobuf object *otg.PatternFlowIpv4VersionCounter
	FromProto(msg *otg.PatternFlowIpv4VersionCounter) (PatternFlowIpv4VersionCounter, error)
	// FromPbText unmarshals PatternFlowIpv4VersionCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4VersionCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4VersionCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4VersionCounter) Marshal() marshalPatternFlowIpv4VersionCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4VersionCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4VersionCounter) Unmarshal() unMarshalPatternFlowIpv4VersionCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4VersionCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4VersionCounter) ToProto() (*otg.PatternFlowIpv4VersionCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4VersionCounter) FromProto(msg *otg.PatternFlowIpv4VersionCounter) (PatternFlowIpv4VersionCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4VersionCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4VersionCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4VersionCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4VersionCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4VersionCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4VersionCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4VersionCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4VersionCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4VersionCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4VersionCounter) Clone() (PatternFlowIpv4VersionCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4VersionCounter()
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

// PatternFlowIpv4VersionCounter is integer counter pattern
type PatternFlowIpv4VersionCounter interface {
	Validation
	// msg marshals PatternFlowIpv4VersionCounter to protobuf object *otg.PatternFlowIpv4VersionCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4VersionCounter
	// setMsg unmarshals PatternFlowIpv4VersionCounter from protobuf object *otg.PatternFlowIpv4VersionCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4VersionCounter) PatternFlowIpv4VersionCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4VersionCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4VersionCounter
	// validate validates PatternFlowIpv4VersionCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4VersionCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4VersionCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4VersionCounter
	SetStart(value uint32) PatternFlowIpv4VersionCounter
	// HasStart checks if Start has been set in PatternFlowIpv4VersionCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4VersionCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4VersionCounter
	SetStep(value uint32) PatternFlowIpv4VersionCounter
	// HasStep checks if Step has been set in PatternFlowIpv4VersionCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4VersionCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4VersionCounter
	SetCount(value uint32) PatternFlowIpv4VersionCounter
	// HasCount checks if Count has been set in PatternFlowIpv4VersionCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4VersionCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4VersionCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4VersionCounter object
func (obj *patternFlowIpv4VersionCounter) SetStart(value uint32) PatternFlowIpv4VersionCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4VersionCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4VersionCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4VersionCounter object
func (obj *patternFlowIpv4VersionCounter) SetStep(value uint32) PatternFlowIpv4VersionCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4VersionCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4VersionCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4VersionCounter object
func (obj *patternFlowIpv4VersionCounter) SetCount(value uint32) PatternFlowIpv4VersionCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4VersionCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4VersionCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4VersionCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4VersionCounter.Count <= 16 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4VersionCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(4)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
