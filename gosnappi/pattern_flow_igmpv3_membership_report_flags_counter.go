package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipReportFlagsCounter *****
type patternFlowIgmpv3MembershipReportFlagsCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipReportFlagsCounter
	marshaller   marshalPatternFlowIgmpv3MembershipReportFlagsCounter
	unMarshaller unMarshalPatternFlowIgmpv3MembershipReportFlagsCounter
}

func NewPatternFlowIgmpv3MembershipReportFlagsCounter() PatternFlowIgmpv3MembershipReportFlagsCounter {
	obj := patternFlowIgmpv3MembershipReportFlagsCounter{obj: &otg.PatternFlowIgmpv3MembershipReportFlagsCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) msg() *otg.PatternFlowIgmpv3MembershipReportFlagsCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) setMsg(msg *otg.PatternFlowIgmpv3MembershipReportFlagsCounter) PatternFlowIgmpv3MembershipReportFlagsCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipReportFlagsCounter struct {
	obj *patternFlowIgmpv3MembershipReportFlagsCounter
}

type marshalPatternFlowIgmpv3MembershipReportFlagsCounter interface {
	// ToProto marshals PatternFlowIgmpv3MembershipReportFlagsCounter to protobuf object *otg.PatternFlowIgmpv3MembershipReportFlagsCounter
	ToProto() (*otg.PatternFlowIgmpv3MembershipReportFlagsCounter, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipReportFlagsCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipReportFlagsCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipReportFlagsCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipReportFlagsCounter struct {
	obj *patternFlowIgmpv3MembershipReportFlagsCounter
}

type unMarshalPatternFlowIgmpv3MembershipReportFlagsCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipReportFlagsCounter from protobuf object *otg.PatternFlowIgmpv3MembershipReportFlagsCounter
	FromProto(msg *otg.PatternFlowIgmpv3MembershipReportFlagsCounter) (PatternFlowIgmpv3MembershipReportFlagsCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipReportFlagsCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipReportFlagsCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipReportFlagsCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) Marshal() marshalPatternFlowIgmpv3MembershipReportFlagsCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipReportFlagsCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportFlagsCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipReportFlagsCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipReportFlagsCounter) ToProto() (*otg.PatternFlowIgmpv3MembershipReportFlagsCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipReportFlagsCounter) FromProto(msg *otg.PatternFlowIgmpv3MembershipReportFlagsCounter) (PatternFlowIgmpv3MembershipReportFlagsCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipReportFlagsCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportFlagsCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportFlagsCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportFlagsCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportFlagsCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportFlagsCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) Clone() (PatternFlowIgmpv3MembershipReportFlagsCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipReportFlagsCounter()
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

// PatternFlowIgmpv3MembershipReportFlagsCounter is integer counter pattern
type PatternFlowIgmpv3MembershipReportFlagsCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipReportFlagsCounter to protobuf object *otg.PatternFlowIgmpv3MembershipReportFlagsCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipReportFlagsCounter
	// setMsg unmarshals PatternFlowIgmpv3MembershipReportFlagsCounter from protobuf object *otg.PatternFlowIgmpv3MembershipReportFlagsCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipReportFlagsCounter) PatternFlowIgmpv3MembershipReportFlagsCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipReportFlagsCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportFlagsCounter
	// validate validates PatternFlowIgmpv3MembershipReportFlagsCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipReportFlagsCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3MembershipReportFlagsCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportFlagsCounter
	SetStart(value uint32) PatternFlowIgmpv3MembershipReportFlagsCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3MembershipReportFlagsCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3MembershipReportFlagsCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportFlagsCounter
	SetStep(value uint32) PatternFlowIgmpv3MembershipReportFlagsCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3MembershipReportFlagsCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3MembershipReportFlagsCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportFlagsCounter
	SetCount(value uint32) PatternFlowIgmpv3MembershipReportFlagsCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3MembershipReportFlagsCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3MembershipReportFlagsCounter object
func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) SetStart(value uint32) PatternFlowIgmpv3MembershipReportFlagsCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3MembershipReportFlagsCounter object
func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) SetStep(value uint32) PatternFlowIgmpv3MembershipReportFlagsCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3MembershipReportFlagsCounter object
func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) SetCount(value uint32) PatternFlowIgmpv3MembershipReportFlagsCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportFlagsCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportFlagsCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportFlagsCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipReportFlagsCounter) setDefault() {
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
