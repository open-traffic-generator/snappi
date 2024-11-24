package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter *****
type patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	marshaller   marshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter() PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter {
	obj := patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter{obj: &otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) msg() *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
}

type marshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) (PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) (PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) Clone() (PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter()
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

// PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter is integer counter pattern
type PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	// validate validates PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(6)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
