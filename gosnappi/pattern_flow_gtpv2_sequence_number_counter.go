package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2SequenceNumberCounter *****
type patternFlowGtpv2SequenceNumberCounter struct {
	validation
	obj          *otg.PatternFlowGtpv2SequenceNumberCounter
	marshaller   marshalPatternFlowGtpv2SequenceNumberCounter
	unMarshaller unMarshalPatternFlowGtpv2SequenceNumberCounter
}

func NewPatternFlowGtpv2SequenceNumberCounter() PatternFlowGtpv2SequenceNumberCounter {
	obj := patternFlowGtpv2SequenceNumberCounter{obj: &otg.PatternFlowGtpv2SequenceNumberCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2SequenceNumberCounter) msg() *otg.PatternFlowGtpv2SequenceNumberCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2SequenceNumberCounter) setMsg(msg *otg.PatternFlowGtpv2SequenceNumberCounter) PatternFlowGtpv2SequenceNumberCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2SequenceNumberCounter struct {
	obj *patternFlowGtpv2SequenceNumberCounter
}

type marshalPatternFlowGtpv2SequenceNumberCounter interface {
	// ToProto marshals PatternFlowGtpv2SequenceNumberCounter to protobuf object *otg.PatternFlowGtpv2SequenceNumberCounter
	ToProto() (*otg.PatternFlowGtpv2SequenceNumberCounter, error)
	// ToPbText marshals PatternFlowGtpv2SequenceNumberCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2SequenceNumberCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2SequenceNumberCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv2SequenceNumberCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv2SequenceNumberCounter struct {
	obj *patternFlowGtpv2SequenceNumberCounter
}

type unMarshalPatternFlowGtpv2SequenceNumberCounter interface {
	// FromProto unmarshals PatternFlowGtpv2SequenceNumberCounter from protobuf object *otg.PatternFlowGtpv2SequenceNumberCounter
	FromProto(msg *otg.PatternFlowGtpv2SequenceNumberCounter) (PatternFlowGtpv2SequenceNumberCounter, error)
	// FromPbText unmarshals PatternFlowGtpv2SequenceNumberCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2SequenceNumberCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2SequenceNumberCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2SequenceNumberCounter) Marshal() marshalPatternFlowGtpv2SequenceNumberCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2SequenceNumberCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2SequenceNumberCounter) Unmarshal() unMarshalPatternFlowGtpv2SequenceNumberCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2SequenceNumberCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2SequenceNumberCounter) ToProto() (*otg.PatternFlowGtpv2SequenceNumberCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2SequenceNumberCounter) FromProto(msg *otg.PatternFlowGtpv2SequenceNumberCounter) (PatternFlowGtpv2SequenceNumberCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2SequenceNumberCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2SequenceNumberCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2SequenceNumberCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2SequenceNumberCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2SequenceNumberCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv2SequenceNumberCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2SequenceNumberCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv2SequenceNumberCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2SequenceNumberCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2SequenceNumberCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2SequenceNumberCounter) Clone() (PatternFlowGtpv2SequenceNumberCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2SequenceNumberCounter()
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

// PatternFlowGtpv2SequenceNumberCounter is integer counter pattern
type PatternFlowGtpv2SequenceNumberCounter interface {
	Validation
	// msg marshals PatternFlowGtpv2SequenceNumberCounter to protobuf object *otg.PatternFlowGtpv2SequenceNumberCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2SequenceNumberCounter
	// setMsg unmarshals PatternFlowGtpv2SequenceNumberCounter from protobuf object *otg.PatternFlowGtpv2SequenceNumberCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2SequenceNumberCounter) PatternFlowGtpv2SequenceNumberCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2SequenceNumberCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2SequenceNumberCounter
	// validate validates PatternFlowGtpv2SequenceNumberCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2SequenceNumberCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv2SequenceNumberCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv2SequenceNumberCounter
	SetStart(value uint32) PatternFlowGtpv2SequenceNumberCounter
	// HasStart checks if Start has been set in PatternFlowGtpv2SequenceNumberCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv2SequenceNumberCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv2SequenceNumberCounter
	SetStep(value uint32) PatternFlowGtpv2SequenceNumberCounter
	// HasStep checks if Step has been set in PatternFlowGtpv2SequenceNumberCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv2SequenceNumberCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv2SequenceNumberCounter
	SetCount(value uint32) PatternFlowGtpv2SequenceNumberCounter
	// HasCount checks if Count has been set in PatternFlowGtpv2SequenceNumberCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2SequenceNumberCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2SequenceNumberCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv2SequenceNumberCounter object
func (obj *patternFlowGtpv2SequenceNumberCounter) SetStart(value uint32) PatternFlowGtpv2SequenceNumberCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2SequenceNumberCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2SequenceNumberCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv2SequenceNumberCounter object
func (obj *patternFlowGtpv2SequenceNumberCounter) SetStep(value uint32) PatternFlowGtpv2SequenceNumberCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2SequenceNumberCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2SequenceNumberCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv2SequenceNumberCounter object
func (obj *patternFlowGtpv2SequenceNumberCounter) SetCount(value uint32) PatternFlowGtpv2SequenceNumberCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv2SequenceNumberCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2SequenceNumberCounter.Start <= 16777215 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2SequenceNumberCounter.Step <= 16777215 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2SequenceNumberCounter.Count <= 16777215 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv2SequenceNumberCounter) setDefault() {
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
