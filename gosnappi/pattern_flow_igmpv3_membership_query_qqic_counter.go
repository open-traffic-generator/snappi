package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryQqicCounter *****
type patternFlowIgmpv3MembershipQueryQqicCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipQueryQqicCounter
	marshaller   marshalPatternFlowIgmpv3MembershipQueryQqicCounter
	unMarshaller unMarshalPatternFlowIgmpv3MembershipQueryQqicCounter
}

func NewPatternFlowIgmpv3MembershipQueryQqicCounter() PatternFlowIgmpv3MembershipQueryQqicCounter {
	obj := patternFlowIgmpv3MembershipQueryQqicCounter{obj: &otg.PatternFlowIgmpv3MembershipQueryQqicCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) msg() *otg.PatternFlowIgmpv3MembershipQueryQqicCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryQqicCounter) PatternFlowIgmpv3MembershipQueryQqicCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryQqicCounter struct {
	obj *patternFlowIgmpv3MembershipQueryQqicCounter
}

type marshalPatternFlowIgmpv3MembershipQueryQqicCounter interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryQqicCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryQqicCounter
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryQqicCounter, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryQqicCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryQqicCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryQqicCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryQqicCounter struct {
	obj *patternFlowIgmpv3MembershipQueryQqicCounter
}

type unMarshalPatternFlowIgmpv3MembershipQueryQqicCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryQqicCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryQqicCounter
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryQqicCounter) (PatternFlowIgmpv3MembershipQueryQqicCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryQqicCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryQqicCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryQqicCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) Marshal() marshalPatternFlowIgmpv3MembershipQueryQqicCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryQqicCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryQqicCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryQqicCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryQqicCounter) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryQqicCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQqicCounter) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryQqicCounter) (PatternFlowIgmpv3MembershipQueryQqicCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryQqicCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQqicCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryQqicCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQqicCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryQqicCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQqicCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) Clone() (PatternFlowIgmpv3MembershipQueryQqicCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryQqicCounter()
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

// PatternFlowIgmpv3MembershipQueryQqicCounter is integer counter pattern
type PatternFlowIgmpv3MembershipQueryQqicCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryQqicCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryQqicCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryQqicCounter
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryQqicCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryQqicCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryQqicCounter) PatternFlowIgmpv3MembershipQueryQqicCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryQqicCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryQqicCounter
	// validate validates PatternFlowIgmpv3MembershipQueryQqicCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryQqicCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3MembershipQueryQqicCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryQqicCounter
	SetStart(value uint32) PatternFlowIgmpv3MembershipQueryQqicCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3MembershipQueryQqicCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3MembershipQueryQqicCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryQqicCounter
	SetStep(value uint32) PatternFlowIgmpv3MembershipQueryQqicCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3MembershipQueryQqicCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3MembershipQueryQqicCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryQqicCounter
	SetCount(value uint32) PatternFlowIgmpv3MembershipQueryQqicCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3MembershipQueryQqicCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3MembershipQueryQqicCounter object
func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) SetStart(value uint32) PatternFlowIgmpv3MembershipQueryQqicCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3MembershipQueryQqicCounter object
func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) SetStep(value uint32) PatternFlowIgmpv3MembershipQueryQqicCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3MembershipQueryQqicCounter object
func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) SetCount(value uint32) PatternFlowIgmpv3MembershipQueryQqicCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryQqicCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryQqicCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryQqicCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipQueryQqicCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(125)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
