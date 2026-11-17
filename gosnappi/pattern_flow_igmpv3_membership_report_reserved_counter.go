package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipReportReservedCounter *****
type patternFlowIgmpv3MembershipReportReservedCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipReportReservedCounter
	marshaller   marshalPatternFlowIgmpv3MembershipReportReservedCounter
	unMarshaller unMarshalPatternFlowIgmpv3MembershipReportReservedCounter
}

func NewPatternFlowIgmpv3MembershipReportReservedCounter() PatternFlowIgmpv3MembershipReportReservedCounter {
	obj := patternFlowIgmpv3MembershipReportReservedCounter{obj: &otg.PatternFlowIgmpv3MembershipReportReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipReportReservedCounter) msg() *otg.PatternFlowIgmpv3MembershipReportReservedCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipReportReservedCounter) setMsg(msg *otg.PatternFlowIgmpv3MembershipReportReservedCounter) PatternFlowIgmpv3MembershipReportReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipReportReservedCounter struct {
	obj *patternFlowIgmpv3MembershipReportReservedCounter
}

type marshalPatternFlowIgmpv3MembershipReportReservedCounter interface {
	// ToProto marshals PatternFlowIgmpv3MembershipReportReservedCounter to protobuf object *otg.PatternFlowIgmpv3MembershipReportReservedCounter
	ToProto() (*otg.PatternFlowIgmpv3MembershipReportReservedCounter, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipReportReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipReportReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipReportReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipReportReservedCounter struct {
	obj *patternFlowIgmpv3MembershipReportReservedCounter
}

type unMarshalPatternFlowIgmpv3MembershipReportReservedCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipReportReservedCounter from protobuf object *otg.PatternFlowIgmpv3MembershipReportReservedCounter
	FromProto(msg *otg.PatternFlowIgmpv3MembershipReportReservedCounter) (PatternFlowIgmpv3MembershipReportReservedCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipReportReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipReportReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipReportReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipReportReservedCounter) Marshal() marshalPatternFlowIgmpv3MembershipReportReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipReportReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipReportReservedCounter) Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipReportReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipReportReservedCounter) ToProto() (*otg.PatternFlowIgmpv3MembershipReportReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipReportReservedCounter) FromProto(msg *otg.PatternFlowIgmpv3MembershipReportReservedCounter) (PatternFlowIgmpv3MembershipReportReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipReportReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipReportReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipReportReservedCounter) Clone() (PatternFlowIgmpv3MembershipReportReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipReportReservedCounter()
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

// PatternFlowIgmpv3MembershipReportReservedCounter is integer counter pattern
type PatternFlowIgmpv3MembershipReportReservedCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipReportReservedCounter to protobuf object *otg.PatternFlowIgmpv3MembershipReportReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipReportReservedCounter
	// setMsg unmarshals PatternFlowIgmpv3MembershipReportReservedCounter from protobuf object *otg.PatternFlowIgmpv3MembershipReportReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipReportReservedCounter) PatternFlowIgmpv3MembershipReportReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipReportReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportReservedCounter
	// validate validates PatternFlowIgmpv3MembershipReportReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipReportReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3MembershipReportReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportReservedCounter
	SetStart(value uint32) PatternFlowIgmpv3MembershipReportReservedCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3MembershipReportReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3MembershipReportReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportReservedCounter
	SetStep(value uint32) PatternFlowIgmpv3MembershipReportReservedCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3MembershipReportReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3MembershipReportReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportReservedCounter
	SetCount(value uint32) PatternFlowIgmpv3MembershipReportReservedCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3MembershipReportReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipReportReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipReportReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3MembershipReportReservedCounter object
func (obj *patternFlowIgmpv3MembershipReportReservedCounter) SetStart(value uint32) PatternFlowIgmpv3MembershipReportReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipReportReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipReportReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3MembershipReportReservedCounter object
func (obj *patternFlowIgmpv3MembershipReportReservedCounter) SetStep(value uint32) PatternFlowIgmpv3MembershipReportReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipReportReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipReportReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3MembershipReportReservedCounter object
func (obj *patternFlowIgmpv3MembershipReportReservedCounter) SetCount(value uint32) PatternFlowIgmpv3MembershipReportReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipReportReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportReservedCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportReservedCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportReservedCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipReportReservedCounter) setDefault() {
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
