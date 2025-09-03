package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerLengthCounter *****
type patternFlowLacpPartnerLengthCounter struct {
	validation
	obj          *otg.PatternFlowLacpPartnerLengthCounter
	marshaller   marshalPatternFlowLacpPartnerLengthCounter
	unMarshaller unMarshalPatternFlowLacpPartnerLengthCounter
}

func NewPatternFlowLacpPartnerLengthCounter() PatternFlowLacpPartnerLengthCounter {
	obj := patternFlowLacpPartnerLengthCounter{obj: &otg.PatternFlowLacpPartnerLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerLengthCounter) msg() *otg.PatternFlowLacpPartnerLengthCounter {
	return obj.obj
}

func (obj *patternFlowLacpPartnerLengthCounter) setMsg(msg *otg.PatternFlowLacpPartnerLengthCounter) PatternFlowLacpPartnerLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerLengthCounter struct {
	obj *patternFlowLacpPartnerLengthCounter
}

type marshalPatternFlowLacpPartnerLengthCounter interface {
	// ToProto marshals PatternFlowLacpPartnerLengthCounter to protobuf object *otg.PatternFlowLacpPartnerLengthCounter
	ToProto() (*otg.PatternFlowLacpPartnerLengthCounter, error)
	// ToPbText marshals PatternFlowLacpPartnerLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerLengthCounter struct {
	obj *patternFlowLacpPartnerLengthCounter
}

type unMarshalPatternFlowLacpPartnerLengthCounter interface {
	// FromProto unmarshals PatternFlowLacpPartnerLengthCounter from protobuf object *otg.PatternFlowLacpPartnerLengthCounter
	FromProto(msg *otg.PatternFlowLacpPartnerLengthCounter) (PatternFlowLacpPartnerLengthCounter, error)
	// FromPbText unmarshals PatternFlowLacpPartnerLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerLengthCounter) Marshal() marshalPatternFlowLacpPartnerLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerLengthCounter) Unmarshal() unMarshalPatternFlowLacpPartnerLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerLengthCounter) ToProto() (*otg.PatternFlowLacpPartnerLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerLengthCounter) FromProto(msg *otg.PatternFlowLacpPartnerLengthCounter) (PatternFlowLacpPartnerLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerLengthCounter) Clone() (PatternFlowLacpPartnerLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerLengthCounter()
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

// PatternFlowLacpPartnerLengthCounter is integer counter pattern
type PatternFlowLacpPartnerLengthCounter interface {
	Validation
	// msg marshals PatternFlowLacpPartnerLengthCounter to protobuf object *otg.PatternFlowLacpPartnerLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerLengthCounter
	// setMsg unmarshals PatternFlowLacpPartnerLengthCounter from protobuf object *otg.PatternFlowLacpPartnerLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerLengthCounter) PatternFlowLacpPartnerLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerLengthCounter
	// validate validates PatternFlowLacpPartnerLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpPartnerLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpPartnerLengthCounter
	SetStart(value uint32) PatternFlowLacpPartnerLengthCounter
	// HasStart checks if Start has been set in PatternFlowLacpPartnerLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpPartnerLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpPartnerLengthCounter
	SetStep(value uint32) PatternFlowLacpPartnerLengthCounter
	// HasStep checks if Step has been set in PatternFlowLacpPartnerLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpPartnerLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpPartnerLengthCounter
	SetCount(value uint32) PatternFlowLacpPartnerLengthCounter
	// HasCount checks if Count has been set in PatternFlowLacpPartnerLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpPartnerLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpPartnerLengthCounter object
func (obj *patternFlowLacpPartnerLengthCounter) SetStart(value uint32) PatternFlowLacpPartnerLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpPartnerLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpPartnerLengthCounter object
func (obj *patternFlowLacpPartnerLengthCounter) SetStep(value uint32) PatternFlowLacpPartnerLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpPartnerLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpPartnerLengthCounter object
func (obj *patternFlowLacpPartnerLengthCounter) SetCount(value uint32) PatternFlowLacpPartnerLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpPartnerLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerLengthCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpPartnerLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(20)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
