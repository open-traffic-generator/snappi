package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryQrvCounter *****
type patternFlowIgmpv3MembershipQueryQrvCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipQueryQrvCounter
	marshaller   marshalPatternFlowIgmpv3MembershipQueryQrvCounter
	unMarshaller unMarshalPatternFlowIgmpv3MembershipQueryQrvCounter
}

func NewPatternFlowIgmpv3MembershipQueryQrvCounter() PatternFlowIgmpv3MembershipQueryQrvCounter {
	obj := patternFlowIgmpv3MembershipQueryQrvCounter{obj: &otg.PatternFlowIgmpv3MembershipQueryQrvCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) msg() *otg.PatternFlowIgmpv3MembershipQueryQrvCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryQrvCounter) PatternFlowIgmpv3MembershipQueryQrvCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryQrvCounter struct {
	obj *patternFlowIgmpv3MembershipQueryQrvCounter
}

type marshalPatternFlowIgmpv3MembershipQueryQrvCounter interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryQrvCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryQrvCounter
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryQrvCounter, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryQrvCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryQrvCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryQrvCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryQrvCounter struct {
	obj *patternFlowIgmpv3MembershipQueryQrvCounter
}

type unMarshalPatternFlowIgmpv3MembershipQueryQrvCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryQrvCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryQrvCounter
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryQrvCounter) (PatternFlowIgmpv3MembershipQueryQrvCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryQrvCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryQrvCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryQrvCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) Marshal() marshalPatternFlowIgmpv3MembershipQueryQrvCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryQrvCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryQrvCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryQrvCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryQrvCounter) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryQrvCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQrvCounter) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryQrvCounter) (PatternFlowIgmpv3MembershipQueryQrvCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryQrvCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQrvCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryQrvCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQrvCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryQrvCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQrvCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) Clone() (PatternFlowIgmpv3MembershipQueryQrvCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryQrvCounter()
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

// PatternFlowIgmpv3MembershipQueryQrvCounter is integer counter pattern
type PatternFlowIgmpv3MembershipQueryQrvCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryQrvCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryQrvCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryQrvCounter
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryQrvCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryQrvCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryQrvCounter) PatternFlowIgmpv3MembershipQueryQrvCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryQrvCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryQrvCounter
	// validate validates PatternFlowIgmpv3MembershipQueryQrvCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryQrvCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3MembershipQueryQrvCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryQrvCounter
	SetStart(value uint32) PatternFlowIgmpv3MembershipQueryQrvCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3MembershipQueryQrvCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3MembershipQueryQrvCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryQrvCounter
	SetStep(value uint32) PatternFlowIgmpv3MembershipQueryQrvCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3MembershipQueryQrvCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3MembershipQueryQrvCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryQrvCounter
	SetCount(value uint32) PatternFlowIgmpv3MembershipQueryQrvCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3MembershipQueryQrvCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3MembershipQueryQrvCounter object
func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) SetStart(value uint32) PatternFlowIgmpv3MembershipQueryQrvCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3MembershipQueryQrvCounter object
func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) SetStep(value uint32) PatternFlowIgmpv3MembershipQueryQrvCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3MembershipQueryQrvCounter object
func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) SetCount(value uint32) PatternFlowIgmpv3MembershipQueryQrvCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryQrvCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryQrvCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryQrvCounter.Count <= 8 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipQueryQrvCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(2)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
