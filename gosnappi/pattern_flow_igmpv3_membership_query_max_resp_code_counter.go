package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter *****
type patternFlowIgmpv3MembershipQueryMaxRespCodeCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	marshaller   marshalPatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	unMarshaller unMarshalPatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
}

func NewPatternFlowIgmpv3MembershipQueryMaxRespCodeCounter() PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter {
	obj := patternFlowIgmpv3MembershipQueryMaxRespCodeCounter{obj: &otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) msg() *otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryMaxRespCodeCounter struct {
	obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter
}

type marshalPatternFlowIgmpv3MembershipQueryMaxRespCodeCounter interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryMaxRespCodeCounter struct {
	obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter
}

type unMarshalPatternFlowIgmpv3MembershipQueryMaxRespCodeCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) (PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) Marshal() marshalPatternFlowIgmpv3MembershipQueryMaxRespCodeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryMaxRespCodeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryMaxRespCodeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryMaxRespCodeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) (PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) Clone() (PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryMaxRespCodeCounter()
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

// PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter is integer counter pattern
type PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	// validate validates PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	SetStart(value uint32) PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	SetStep(value uint32) PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	SetCount(value uint32) PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter object
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) SetStart(value uint32) PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter object
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) SetStep(value uint32) PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter object
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) SetCount(value uint32) PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCodeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(100)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
