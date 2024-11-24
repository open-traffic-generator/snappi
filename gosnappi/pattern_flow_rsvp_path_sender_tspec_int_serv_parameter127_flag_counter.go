package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter *****
type patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	marshaller   marshalPatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter() PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter {
	obj := patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter{obj: &otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) msg() *otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
}

type marshalPatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) (PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) (PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) Clone() (PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter()
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

// PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter is integer counter pattern
type PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	// validate validates PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) setDefault() {
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
