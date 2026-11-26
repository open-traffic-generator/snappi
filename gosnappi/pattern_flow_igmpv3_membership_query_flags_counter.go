package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryFlagsCounter *****
type patternFlowIgmpv3MembershipQueryFlagsCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipQueryFlagsCounter
	marshaller   marshalPatternFlowIgmpv3MembershipQueryFlagsCounter
	unMarshaller unMarshalPatternFlowIgmpv3MembershipQueryFlagsCounter
}

func NewPatternFlowIgmpv3MembershipQueryFlagsCounter() PatternFlowIgmpv3MembershipQueryFlagsCounter {
	obj := patternFlowIgmpv3MembershipQueryFlagsCounter{obj: &otg.PatternFlowIgmpv3MembershipQueryFlagsCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) msg() *otg.PatternFlowIgmpv3MembershipQueryFlagsCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryFlagsCounter) PatternFlowIgmpv3MembershipQueryFlagsCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryFlagsCounter struct {
	obj *patternFlowIgmpv3MembershipQueryFlagsCounter
}

type marshalPatternFlowIgmpv3MembershipQueryFlagsCounter interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryFlagsCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryFlagsCounter
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryFlagsCounter, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryFlagsCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryFlagsCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryFlagsCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryFlagsCounter struct {
	obj *patternFlowIgmpv3MembershipQueryFlagsCounter
}

type unMarshalPatternFlowIgmpv3MembershipQueryFlagsCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryFlagsCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryFlagsCounter
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryFlagsCounter) (PatternFlowIgmpv3MembershipQueryFlagsCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryFlagsCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryFlagsCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryFlagsCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) Marshal() marshalPatternFlowIgmpv3MembershipQueryFlagsCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryFlagsCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryFlagsCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryFlagsCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryFlagsCounter) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryFlagsCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryFlagsCounter) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryFlagsCounter) (PatternFlowIgmpv3MembershipQueryFlagsCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryFlagsCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryFlagsCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryFlagsCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryFlagsCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryFlagsCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryFlagsCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) Clone() (PatternFlowIgmpv3MembershipQueryFlagsCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryFlagsCounter()
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

// PatternFlowIgmpv3MembershipQueryFlagsCounter is integer counter pattern
type PatternFlowIgmpv3MembershipQueryFlagsCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryFlagsCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryFlagsCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryFlagsCounter
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryFlagsCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryFlagsCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryFlagsCounter) PatternFlowIgmpv3MembershipQueryFlagsCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryFlagsCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryFlagsCounter
	// validate validates PatternFlowIgmpv3MembershipQueryFlagsCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryFlagsCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3MembershipQueryFlagsCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryFlagsCounter
	SetStart(value uint32) PatternFlowIgmpv3MembershipQueryFlagsCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3MembershipQueryFlagsCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3MembershipQueryFlagsCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryFlagsCounter
	SetStep(value uint32) PatternFlowIgmpv3MembershipQueryFlagsCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3MembershipQueryFlagsCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3MembershipQueryFlagsCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryFlagsCounter
	SetCount(value uint32) PatternFlowIgmpv3MembershipQueryFlagsCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3MembershipQueryFlagsCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3MembershipQueryFlagsCounter object
func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) SetStart(value uint32) PatternFlowIgmpv3MembershipQueryFlagsCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3MembershipQueryFlagsCounter object
func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) SetStep(value uint32) PatternFlowIgmpv3MembershipQueryFlagsCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3MembershipQueryFlagsCounter object
func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) SetCount(value uint32) PatternFlowIgmpv3MembershipQueryFlagsCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryFlagsCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryFlagsCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryFlagsCounter.Count <= 16 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipQueryFlagsCounter) setDefault() {
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
