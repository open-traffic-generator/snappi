package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQuerySFlagCounter *****
type patternFlowIgmpv3MembershipQuerySFlagCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipQuerySFlagCounter
	marshaller   marshalPatternFlowIgmpv3MembershipQuerySFlagCounter
	unMarshaller unMarshalPatternFlowIgmpv3MembershipQuerySFlagCounter
}

func NewPatternFlowIgmpv3MembershipQuerySFlagCounter() PatternFlowIgmpv3MembershipQuerySFlagCounter {
	obj := patternFlowIgmpv3MembershipQuerySFlagCounter{obj: &otg.PatternFlowIgmpv3MembershipQuerySFlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) msg() *otg.PatternFlowIgmpv3MembershipQuerySFlagCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) setMsg(msg *otg.PatternFlowIgmpv3MembershipQuerySFlagCounter) PatternFlowIgmpv3MembershipQuerySFlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQuerySFlagCounter struct {
	obj *patternFlowIgmpv3MembershipQuerySFlagCounter
}

type marshalPatternFlowIgmpv3MembershipQuerySFlagCounter interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQuerySFlagCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQuerySFlagCounter
	ToProto() (*otg.PatternFlowIgmpv3MembershipQuerySFlagCounter, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQuerySFlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQuerySFlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQuerySFlagCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQuerySFlagCounter struct {
	obj *patternFlowIgmpv3MembershipQuerySFlagCounter
}

type unMarshalPatternFlowIgmpv3MembershipQuerySFlagCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQuerySFlagCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQuerySFlagCounter
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQuerySFlagCounter) (PatternFlowIgmpv3MembershipQuerySFlagCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQuerySFlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQuerySFlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQuerySFlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) Marshal() marshalPatternFlowIgmpv3MembershipQuerySFlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQuerySFlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQuerySFlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQuerySFlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQuerySFlagCounter) ToProto() (*otg.PatternFlowIgmpv3MembershipQuerySFlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQuerySFlagCounter) FromProto(msg *otg.PatternFlowIgmpv3MembershipQuerySFlagCounter) (PatternFlowIgmpv3MembershipQuerySFlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQuerySFlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQuerySFlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQuerySFlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQuerySFlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQuerySFlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQuerySFlagCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) Clone() (PatternFlowIgmpv3MembershipQuerySFlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQuerySFlagCounter()
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

// PatternFlowIgmpv3MembershipQuerySFlagCounter is integer counter pattern
type PatternFlowIgmpv3MembershipQuerySFlagCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQuerySFlagCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQuerySFlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQuerySFlagCounter
	// setMsg unmarshals PatternFlowIgmpv3MembershipQuerySFlagCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQuerySFlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQuerySFlagCounter) PatternFlowIgmpv3MembershipQuerySFlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQuerySFlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQuerySFlagCounter
	// validate validates PatternFlowIgmpv3MembershipQuerySFlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQuerySFlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3MembershipQuerySFlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3MembershipQuerySFlagCounter
	SetStart(value uint32) PatternFlowIgmpv3MembershipQuerySFlagCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3MembershipQuerySFlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3MembershipQuerySFlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3MembershipQuerySFlagCounter
	SetStep(value uint32) PatternFlowIgmpv3MembershipQuerySFlagCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3MembershipQuerySFlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3MembershipQuerySFlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3MembershipQuerySFlagCounter
	SetCount(value uint32) PatternFlowIgmpv3MembershipQuerySFlagCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3MembershipQuerySFlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3MembershipQuerySFlagCounter object
func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) SetStart(value uint32) PatternFlowIgmpv3MembershipQuerySFlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3MembershipQuerySFlagCounter object
func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) SetStep(value uint32) PatternFlowIgmpv3MembershipQuerySFlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3MembershipQuerySFlagCounter object
func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) SetCount(value uint32) PatternFlowIgmpv3MembershipQuerySFlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQuerySFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQuerySFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQuerySFlagCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipQuerySFlagCounter) setDefault() {
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
