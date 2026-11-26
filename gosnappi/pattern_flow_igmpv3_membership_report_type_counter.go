package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipReportTypeCounter *****
type patternFlowIgmpv3MembershipReportTypeCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipReportTypeCounter
	marshaller   marshalPatternFlowIgmpv3MembershipReportTypeCounter
	unMarshaller unMarshalPatternFlowIgmpv3MembershipReportTypeCounter
}

func NewPatternFlowIgmpv3MembershipReportTypeCounter() PatternFlowIgmpv3MembershipReportTypeCounter {
	obj := patternFlowIgmpv3MembershipReportTypeCounter{obj: &otg.PatternFlowIgmpv3MembershipReportTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipReportTypeCounter) msg() *otg.PatternFlowIgmpv3MembershipReportTypeCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipReportTypeCounter) setMsg(msg *otg.PatternFlowIgmpv3MembershipReportTypeCounter) PatternFlowIgmpv3MembershipReportTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipReportTypeCounter struct {
	obj *patternFlowIgmpv3MembershipReportTypeCounter
}

type marshalPatternFlowIgmpv3MembershipReportTypeCounter interface {
	// ToProto marshals PatternFlowIgmpv3MembershipReportTypeCounter to protobuf object *otg.PatternFlowIgmpv3MembershipReportTypeCounter
	ToProto() (*otg.PatternFlowIgmpv3MembershipReportTypeCounter, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipReportTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipReportTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipReportTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipReportTypeCounter struct {
	obj *patternFlowIgmpv3MembershipReportTypeCounter
}

type unMarshalPatternFlowIgmpv3MembershipReportTypeCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipReportTypeCounter from protobuf object *otg.PatternFlowIgmpv3MembershipReportTypeCounter
	FromProto(msg *otg.PatternFlowIgmpv3MembershipReportTypeCounter) (PatternFlowIgmpv3MembershipReportTypeCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipReportTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipReportTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipReportTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipReportTypeCounter) Marshal() marshalPatternFlowIgmpv3MembershipReportTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipReportTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipReportTypeCounter) Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipReportTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipReportTypeCounter) ToProto() (*otg.PatternFlowIgmpv3MembershipReportTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipReportTypeCounter) FromProto(msg *otg.PatternFlowIgmpv3MembershipReportTypeCounter) (PatternFlowIgmpv3MembershipReportTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipReportTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipReportTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipReportTypeCounter) Clone() (PatternFlowIgmpv3MembershipReportTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipReportTypeCounter()
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

// PatternFlowIgmpv3MembershipReportTypeCounter is integer counter pattern
type PatternFlowIgmpv3MembershipReportTypeCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipReportTypeCounter to protobuf object *otg.PatternFlowIgmpv3MembershipReportTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipReportTypeCounter
	// setMsg unmarshals PatternFlowIgmpv3MembershipReportTypeCounter from protobuf object *otg.PatternFlowIgmpv3MembershipReportTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipReportTypeCounter) PatternFlowIgmpv3MembershipReportTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipReportTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportTypeCounter
	// validate validates PatternFlowIgmpv3MembershipReportTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipReportTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3MembershipReportTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportTypeCounter
	SetStart(value uint32) PatternFlowIgmpv3MembershipReportTypeCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3MembershipReportTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3MembershipReportTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportTypeCounter
	SetStep(value uint32) PatternFlowIgmpv3MembershipReportTypeCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3MembershipReportTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3MembershipReportTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportTypeCounter
	SetCount(value uint32) PatternFlowIgmpv3MembershipReportTypeCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3MembershipReportTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipReportTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipReportTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3MembershipReportTypeCounter object
func (obj *patternFlowIgmpv3MembershipReportTypeCounter) SetStart(value uint32) PatternFlowIgmpv3MembershipReportTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipReportTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipReportTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3MembershipReportTypeCounter object
func (obj *patternFlowIgmpv3MembershipReportTypeCounter) SetStep(value uint32) PatternFlowIgmpv3MembershipReportTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipReportTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipReportTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3MembershipReportTypeCounter object
func (obj *patternFlowIgmpv3MembershipReportTypeCounter) SetCount(value uint32) PatternFlowIgmpv3MembershipReportTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipReportTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipReportTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(34)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
