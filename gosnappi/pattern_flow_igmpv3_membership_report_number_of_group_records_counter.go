package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter *****
type patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	marshaller   marshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	unMarshaller unMarshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
}

func NewPatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter() PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter {
	obj := patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter{obj: &otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) msg() *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) setMsg(msg *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter struct {
	obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
}

type marshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter interface {
	// ToProto marshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter to protobuf object *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	ToProto() (*otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter struct {
	obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
}

type unMarshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter from protobuf object *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	FromProto(msg *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) (PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) Marshal() marshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) ToProto() (*otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) FromProto(msg *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) (PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) Clone() (PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter()
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

// PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter is integer counter pattern
type PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter to protobuf object *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	// setMsg unmarshals PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter from protobuf object *otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	// validate validates PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	SetStart(value uint32) PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	// HasStart checks if Start has been set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	SetStep(value uint32) PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	// HasStep checks if Step has been set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	SetCount(value uint32) PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	// HasCount checks if Count has been set in PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter object
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) SetStart(value uint32) PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter object
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) SetStep(value uint32) PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter object
func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) SetCount(value uint32) PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipReportNumberOfGroupRecordsCounter) setDefault() {
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
