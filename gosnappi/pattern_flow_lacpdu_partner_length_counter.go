package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerLengthCounter *****
type patternFlowLacpduPartnerLengthCounter struct {
	validation
	obj          *otg.PatternFlowLacpduPartnerLengthCounter
	marshaller   marshalPatternFlowLacpduPartnerLengthCounter
	unMarshaller unMarshalPatternFlowLacpduPartnerLengthCounter
}

func NewPatternFlowLacpduPartnerLengthCounter() PatternFlowLacpduPartnerLengthCounter {
	obj := patternFlowLacpduPartnerLengthCounter{obj: &otg.PatternFlowLacpduPartnerLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerLengthCounter) msg() *otg.PatternFlowLacpduPartnerLengthCounter {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerLengthCounter) setMsg(msg *otg.PatternFlowLacpduPartnerLengthCounter) PatternFlowLacpduPartnerLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerLengthCounter struct {
	obj *patternFlowLacpduPartnerLengthCounter
}

type marshalPatternFlowLacpduPartnerLengthCounter interface {
	// ToProto marshals PatternFlowLacpduPartnerLengthCounter to protobuf object *otg.PatternFlowLacpduPartnerLengthCounter
	ToProto() (*otg.PatternFlowLacpduPartnerLengthCounter, error)
	// ToPbText marshals PatternFlowLacpduPartnerLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerLengthCounter struct {
	obj *patternFlowLacpduPartnerLengthCounter
}

type unMarshalPatternFlowLacpduPartnerLengthCounter interface {
	// FromProto unmarshals PatternFlowLacpduPartnerLengthCounter from protobuf object *otg.PatternFlowLacpduPartnerLengthCounter
	FromProto(msg *otg.PatternFlowLacpduPartnerLengthCounter) (PatternFlowLacpduPartnerLengthCounter, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerLengthCounter) Marshal() marshalPatternFlowLacpduPartnerLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerLengthCounter) Unmarshal() unMarshalPatternFlowLacpduPartnerLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerLengthCounter) ToProto() (*otg.PatternFlowLacpduPartnerLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerLengthCounter) FromProto(msg *otg.PatternFlowLacpduPartnerLengthCounter) (PatternFlowLacpduPartnerLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerLengthCounter) Clone() (PatternFlowLacpduPartnerLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerLengthCounter()
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

// PatternFlowLacpduPartnerLengthCounter is integer counter pattern
type PatternFlowLacpduPartnerLengthCounter interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerLengthCounter to protobuf object *otg.PatternFlowLacpduPartnerLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerLengthCounter
	// setMsg unmarshals PatternFlowLacpduPartnerLengthCounter from protobuf object *otg.PatternFlowLacpduPartnerLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerLengthCounter) PatternFlowLacpduPartnerLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerLengthCounter
	// validate validates PatternFlowLacpduPartnerLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowLacpduPartnerLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowLacpduPartnerLengthCounter
	SetStart(value uint32) PatternFlowLacpduPartnerLengthCounter
	// HasStart checks if Start has been set in PatternFlowLacpduPartnerLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowLacpduPartnerLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowLacpduPartnerLengthCounter
	SetStep(value uint32) PatternFlowLacpduPartnerLengthCounter
	// HasStep checks if Step has been set in PatternFlowLacpduPartnerLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowLacpduPartnerLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowLacpduPartnerLengthCounter
	SetCount(value uint32) PatternFlowLacpduPartnerLengthCounter
	// HasCount checks if Count has been set in PatternFlowLacpduPartnerLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowLacpduPartnerLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowLacpduPartnerLengthCounter object
func (obj *patternFlowLacpduPartnerLengthCounter) SetStart(value uint32) PatternFlowLacpduPartnerLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowLacpduPartnerLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowLacpduPartnerLengthCounter object
func (obj *patternFlowLacpduPartnerLengthCounter) SetStep(value uint32) PatternFlowLacpduPartnerLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowLacpduPartnerLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowLacpduPartnerLengthCounter object
func (obj *patternFlowLacpduPartnerLengthCounter) SetCount(value uint32) PatternFlowLacpduPartnerLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowLacpduPartnerLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowLacpduPartnerLengthCounter) setDefault() {
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
