package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv1VersionCounter *****
type patternFlowIgmpv1VersionCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv1VersionCounter
	marshaller   marshalPatternFlowIgmpv1VersionCounter
	unMarshaller unMarshalPatternFlowIgmpv1VersionCounter
}

func NewPatternFlowIgmpv1VersionCounter() PatternFlowIgmpv1VersionCounter {
	obj := patternFlowIgmpv1VersionCounter{obj: &otg.PatternFlowIgmpv1VersionCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv1VersionCounter) msg() *otg.PatternFlowIgmpv1VersionCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv1VersionCounter) setMsg(msg *otg.PatternFlowIgmpv1VersionCounter) PatternFlowIgmpv1VersionCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv1VersionCounter struct {
	obj *patternFlowIgmpv1VersionCounter
}

type marshalPatternFlowIgmpv1VersionCounter interface {
	// ToProto marshals PatternFlowIgmpv1VersionCounter to protobuf object *otg.PatternFlowIgmpv1VersionCounter
	ToProto() (*otg.PatternFlowIgmpv1VersionCounter, error)
	// ToPbText marshals PatternFlowIgmpv1VersionCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv1VersionCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv1VersionCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv1VersionCounter struct {
	obj *patternFlowIgmpv1VersionCounter
}

type unMarshalPatternFlowIgmpv1VersionCounter interface {
	// FromProto unmarshals PatternFlowIgmpv1VersionCounter from protobuf object *otg.PatternFlowIgmpv1VersionCounter
	FromProto(msg *otg.PatternFlowIgmpv1VersionCounter) (PatternFlowIgmpv1VersionCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv1VersionCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv1VersionCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv1VersionCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv1VersionCounter) Marshal() marshalPatternFlowIgmpv1VersionCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv1VersionCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv1VersionCounter) Unmarshal() unMarshalPatternFlowIgmpv1VersionCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv1VersionCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv1VersionCounter) ToProto() (*otg.PatternFlowIgmpv1VersionCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv1VersionCounter) FromProto(msg *otg.PatternFlowIgmpv1VersionCounter) (PatternFlowIgmpv1VersionCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv1VersionCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1VersionCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv1VersionCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1VersionCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv1VersionCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1VersionCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv1VersionCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1VersionCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1VersionCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv1VersionCounter) Clone() (PatternFlowIgmpv1VersionCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv1VersionCounter()
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

// PatternFlowIgmpv1VersionCounter is integer counter pattern
type PatternFlowIgmpv1VersionCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv1VersionCounter to protobuf object *otg.PatternFlowIgmpv1VersionCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv1VersionCounter
	// setMsg unmarshals PatternFlowIgmpv1VersionCounter from protobuf object *otg.PatternFlowIgmpv1VersionCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv1VersionCounter) PatternFlowIgmpv1VersionCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv1VersionCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv1VersionCounter
	// validate validates PatternFlowIgmpv1VersionCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv1VersionCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv1VersionCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv1VersionCounter
	SetStart(value uint32) PatternFlowIgmpv1VersionCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv1VersionCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv1VersionCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv1VersionCounter
	SetStep(value uint32) PatternFlowIgmpv1VersionCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv1VersionCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv1VersionCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv1VersionCounter
	SetCount(value uint32) PatternFlowIgmpv1VersionCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv1VersionCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv1VersionCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv1VersionCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv1VersionCounter object
func (obj *patternFlowIgmpv1VersionCounter) SetStart(value uint32) PatternFlowIgmpv1VersionCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv1VersionCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv1VersionCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv1VersionCounter object
func (obj *patternFlowIgmpv1VersionCounter) SetStep(value uint32) PatternFlowIgmpv1VersionCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv1VersionCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv1VersionCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv1VersionCounter object
func (obj *patternFlowIgmpv1VersionCounter) SetCount(value uint32) PatternFlowIgmpv1VersionCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv1VersionCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1VersionCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1VersionCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1VersionCounter.Count <= 16 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv1VersionCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
