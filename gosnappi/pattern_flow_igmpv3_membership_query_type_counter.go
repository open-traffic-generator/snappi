package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryTypeCounter *****
type patternFlowIgmpv3MembershipQueryTypeCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipQueryTypeCounter
	marshaller   marshalPatternFlowIgmpv3MembershipQueryTypeCounter
	unMarshaller unMarshalPatternFlowIgmpv3MembershipQueryTypeCounter
}

func NewPatternFlowIgmpv3MembershipQueryTypeCounter() PatternFlowIgmpv3MembershipQueryTypeCounter {
	obj := patternFlowIgmpv3MembershipQueryTypeCounter{obj: &otg.PatternFlowIgmpv3MembershipQueryTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) msg() *otg.PatternFlowIgmpv3MembershipQueryTypeCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryTypeCounter) PatternFlowIgmpv3MembershipQueryTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryTypeCounter struct {
	obj *patternFlowIgmpv3MembershipQueryTypeCounter
}

type marshalPatternFlowIgmpv3MembershipQueryTypeCounter interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryTypeCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryTypeCounter
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryTypeCounter, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryTypeCounter struct {
	obj *patternFlowIgmpv3MembershipQueryTypeCounter
}

type unMarshalPatternFlowIgmpv3MembershipQueryTypeCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryTypeCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryTypeCounter
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryTypeCounter) (PatternFlowIgmpv3MembershipQueryTypeCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) Marshal() marshalPatternFlowIgmpv3MembershipQueryTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryTypeCounter) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryTypeCounter) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryTypeCounter) (PatternFlowIgmpv3MembershipQueryTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) Clone() (PatternFlowIgmpv3MembershipQueryTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryTypeCounter()
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

// PatternFlowIgmpv3MembershipQueryTypeCounter is integer counter pattern
type PatternFlowIgmpv3MembershipQueryTypeCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryTypeCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryTypeCounter
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryTypeCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryTypeCounter) PatternFlowIgmpv3MembershipQueryTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryTypeCounter
	// validate validates PatternFlowIgmpv3MembershipQueryTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3MembershipQueryTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryTypeCounter
	SetStart(value uint32) PatternFlowIgmpv3MembershipQueryTypeCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3MembershipQueryTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3MembershipQueryTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryTypeCounter
	SetStep(value uint32) PatternFlowIgmpv3MembershipQueryTypeCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3MembershipQueryTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3MembershipQueryTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryTypeCounter
	SetCount(value uint32) PatternFlowIgmpv3MembershipQueryTypeCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3MembershipQueryTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3MembershipQueryTypeCounter object
func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) SetStart(value uint32) PatternFlowIgmpv3MembershipQueryTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3MembershipQueryTypeCounter object
func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) SetStep(value uint32) PatternFlowIgmpv3MembershipQueryTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3MembershipQueryTypeCounter object
func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) SetCount(value uint32) PatternFlowIgmpv3MembershipQueryTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipQueryTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(17)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
