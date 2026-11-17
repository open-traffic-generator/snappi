package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter *****
type patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	marshaller   marshalPatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	unMarshaller unMarshalPatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
}

func NewPatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter() PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter {
	obj := patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter{obj: &otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) msg() *otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter struct {
	obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
}

type marshalPatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter struct {
	obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
}

type unMarshalPatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) (PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) Marshal() marshalPatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) (PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) Clone() (PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter()
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

// PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter is integer counter pattern
type PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter to protobuf object *otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter from protobuf object *otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	// validate validates PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	SetStart(value uint32) PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	SetStep(value uint32) PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	SetCount(value uint32) PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter object
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) SetStart(value uint32) PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter object
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) SetStep(value uint32) PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter object
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) SetCount(value uint32) PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) setDefault() {
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
