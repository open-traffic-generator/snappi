package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter *****
type patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter struct {
	validation
	obj          *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	marshaller   marshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	unMarshaller unMarshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter() PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter {
	obj := patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter{obj: &otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) msg() *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
}

type marshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter struct {
	obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) (PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) (PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) Clone() (PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter()
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

// PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter is integer counter pattern
type PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	// validate validates PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	// HasStart checks if Start has been set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	// HasStep checks if Step has been set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	// HasCount checks if Count has been set in PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) SetStart(value uint32) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) SetStep(value uint32) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter object
func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) SetCount(value uint32) PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServParameterIdTokenBucketTspecCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(127)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
