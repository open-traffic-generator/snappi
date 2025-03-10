package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServZeroBitCounter *****
type patternFlowRSVPPathSenderTspecIntServZeroBitCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	marshaller   marshalPatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTspecIntServZeroBitCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServZeroBitCounter() PatternFlowRSVPPathSenderTspecIntServZeroBitCounter {
	obj := patternFlowRSVPPathSenderTspecIntServZeroBitCounter{obj: &otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) msg() *otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter) PatternFlowRSVPPathSenderTspecIntServZeroBitCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServZeroBitCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter
}

type marshalPatternFlowRSVPPathSenderTspecIntServZeroBitCounter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServZeroBitCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServZeroBitCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServZeroBitCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServZeroBitCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSenderTspecIntServZeroBitCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServZeroBitCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServZeroBitCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServZeroBitCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter) (PatternFlowRSVPPathSenderTspecIntServZeroBitCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServZeroBitCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServZeroBitCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServZeroBitCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServZeroBitCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServZeroBitCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServZeroBitCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServZeroBitCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServZeroBitCounter) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServZeroBitCounter) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter) (PatternFlowRSVPPathSenderTspecIntServZeroBitCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServZeroBitCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServZeroBitCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServZeroBitCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServZeroBitCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServZeroBitCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServZeroBitCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServZeroBitCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) Clone() (PatternFlowRSVPPathSenderTspecIntServZeroBitCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServZeroBitCounter()
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

// PatternFlowRSVPPathSenderTspecIntServZeroBitCounter is integer counter pattern
type PatternFlowRSVPPathSenderTspecIntServZeroBitCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServZeroBitCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServZeroBitCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServZeroBitCounter) PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	// validate validates PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServZeroBitCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTspecIntServZeroBitCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTspecIntServZeroBitCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTspecIntServZeroBitCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServZeroBitCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServZeroBitCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServZeroBitCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServZeroBitCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServZeroBitCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServZeroBitCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServZeroBitCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServZeroBitCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServZeroBitCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBitCounter) setDefault() {
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
