package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1SquenceNumberCounter *****
type patternFlowGtpv1SquenceNumberCounter struct {
	validation
	obj          *otg.PatternFlowGtpv1SquenceNumberCounter
	marshaller   marshalPatternFlowGtpv1SquenceNumberCounter
	unMarshaller unMarshalPatternFlowGtpv1SquenceNumberCounter
}

func NewPatternFlowGtpv1SquenceNumberCounter() PatternFlowGtpv1SquenceNumberCounter {
	obj := patternFlowGtpv1SquenceNumberCounter{obj: &otg.PatternFlowGtpv1SquenceNumberCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1SquenceNumberCounter) msg() *otg.PatternFlowGtpv1SquenceNumberCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1SquenceNumberCounter) setMsg(msg *otg.PatternFlowGtpv1SquenceNumberCounter) PatternFlowGtpv1SquenceNumberCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1SquenceNumberCounter struct {
	obj *patternFlowGtpv1SquenceNumberCounter
}

type marshalPatternFlowGtpv1SquenceNumberCounter interface {
	// ToProto marshals PatternFlowGtpv1SquenceNumberCounter to protobuf object *otg.PatternFlowGtpv1SquenceNumberCounter
	ToProto() (*otg.PatternFlowGtpv1SquenceNumberCounter, error)
	// ToPbText marshals PatternFlowGtpv1SquenceNumberCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1SquenceNumberCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1SquenceNumberCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1SquenceNumberCounter struct {
	obj *patternFlowGtpv1SquenceNumberCounter
}

type unMarshalPatternFlowGtpv1SquenceNumberCounter interface {
	// FromProto unmarshals PatternFlowGtpv1SquenceNumberCounter from protobuf object *otg.PatternFlowGtpv1SquenceNumberCounter
	FromProto(msg *otg.PatternFlowGtpv1SquenceNumberCounter) (PatternFlowGtpv1SquenceNumberCounter, error)
	// FromPbText unmarshals PatternFlowGtpv1SquenceNumberCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1SquenceNumberCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1SquenceNumberCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1SquenceNumberCounter) Marshal() marshalPatternFlowGtpv1SquenceNumberCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1SquenceNumberCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1SquenceNumberCounter) Unmarshal() unMarshalPatternFlowGtpv1SquenceNumberCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1SquenceNumberCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1SquenceNumberCounter) ToProto() (*otg.PatternFlowGtpv1SquenceNumberCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1SquenceNumberCounter) FromProto(msg *otg.PatternFlowGtpv1SquenceNumberCounter) (PatternFlowGtpv1SquenceNumberCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1SquenceNumberCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SquenceNumberCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1SquenceNumberCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SquenceNumberCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1SquenceNumberCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SquenceNumberCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv1SquenceNumberCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1SquenceNumberCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1SquenceNumberCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1SquenceNumberCounter) Clone() (PatternFlowGtpv1SquenceNumberCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1SquenceNumberCounter()
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

// PatternFlowGtpv1SquenceNumberCounter is integer counter pattern
type PatternFlowGtpv1SquenceNumberCounter interface {
	Validation
	// msg marshals PatternFlowGtpv1SquenceNumberCounter to protobuf object *otg.PatternFlowGtpv1SquenceNumberCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1SquenceNumberCounter
	// setMsg unmarshals PatternFlowGtpv1SquenceNumberCounter from protobuf object *otg.PatternFlowGtpv1SquenceNumberCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1SquenceNumberCounter) PatternFlowGtpv1SquenceNumberCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1SquenceNumberCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1SquenceNumberCounter
	// validate validates PatternFlowGtpv1SquenceNumberCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1SquenceNumberCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv1SquenceNumberCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv1SquenceNumberCounter
	SetStart(value uint32) PatternFlowGtpv1SquenceNumberCounter
	// HasStart checks if Start has been set in PatternFlowGtpv1SquenceNumberCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv1SquenceNumberCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv1SquenceNumberCounter
	SetStep(value uint32) PatternFlowGtpv1SquenceNumberCounter
	// HasStep checks if Step has been set in PatternFlowGtpv1SquenceNumberCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv1SquenceNumberCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv1SquenceNumberCounter
	SetCount(value uint32) PatternFlowGtpv1SquenceNumberCounter
	// HasCount checks if Count has been set in PatternFlowGtpv1SquenceNumberCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1SquenceNumberCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1SquenceNumberCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv1SquenceNumberCounter object
func (obj *patternFlowGtpv1SquenceNumberCounter) SetStart(value uint32) PatternFlowGtpv1SquenceNumberCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1SquenceNumberCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1SquenceNumberCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv1SquenceNumberCounter object
func (obj *patternFlowGtpv1SquenceNumberCounter) SetStep(value uint32) PatternFlowGtpv1SquenceNumberCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1SquenceNumberCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1SquenceNumberCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv1SquenceNumberCounter object
func (obj *patternFlowGtpv1SquenceNumberCounter) SetCount(value uint32) PatternFlowGtpv1SquenceNumberCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv1SquenceNumberCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SquenceNumberCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SquenceNumberCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65536 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SquenceNumberCounter.Count <= 65536 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv1SquenceNumberCounter) setDefault() {
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
