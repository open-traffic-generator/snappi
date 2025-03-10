package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter *****
type patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	marshaller   marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	unMarshaller unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
}

func NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter() PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter {
	obj := patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter{obj: &otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) msg() *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) setMsg(msg *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter struct {
	obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
}

type marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter interface {
	// ToProto marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter to protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	ToProto() (*otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter, error)
	// ToPbText marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter struct {
	obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
}

type unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter from protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	FromProto(msg *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) (PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) Marshal() marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) Unmarshal() unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) ToProto() (*otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) FromProto(msg *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) (PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) Clone() (PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter()
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

// PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter is integer counter pattern
type PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter to protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	// setMsg unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter from protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	// validate validates PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	SetStart(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	SetStep(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	SetCount(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) SetStart(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) SetStep(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) SetCount(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) setDefault() {
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
