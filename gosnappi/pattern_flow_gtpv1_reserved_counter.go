package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1ReservedCounter *****
type patternFlowGtpv1ReservedCounter struct {
	validation
	obj          *otg.PatternFlowGtpv1ReservedCounter
	marshaller   marshalPatternFlowGtpv1ReservedCounter
	unMarshaller unMarshalPatternFlowGtpv1ReservedCounter
}

func NewPatternFlowGtpv1ReservedCounter() PatternFlowGtpv1ReservedCounter {
	obj := patternFlowGtpv1ReservedCounter{obj: &otg.PatternFlowGtpv1ReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1ReservedCounter) msg() *otg.PatternFlowGtpv1ReservedCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1ReservedCounter) setMsg(msg *otg.PatternFlowGtpv1ReservedCounter) PatternFlowGtpv1ReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1ReservedCounter struct {
	obj *patternFlowGtpv1ReservedCounter
}

type marshalPatternFlowGtpv1ReservedCounter interface {
	// ToProto marshals PatternFlowGtpv1ReservedCounter to protobuf object *otg.PatternFlowGtpv1ReservedCounter
	ToProto() (*otg.PatternFlowGtpv1ReservedCounter, error)
	// ToPbText marshals PatternFlowGtpv1ReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1ReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1ReservedCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1ReservedCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1ReservedCounter struct {
	obj *patternFlowGtpv1ReservedCounter
}

type unMarshalPatternFlowGtpv1ReservedCounter interface {
	// FromProto unmarshals PatternFlowGtpv1ReservedCounter from protobuf object *otg.PatternFlowGtpv1ReservedCounter
	FromProto(msg *otg.PatternFlowGtpv1ReservedCounter) (PatternFlowGtpv1ReservedCounter, error)
	// FromPbText unmarshals PatternFlowGtpv1ReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1ReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1ReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1ReservedCounter) Marshal() marshalPatternFlowGtpv1ReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1ReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1ReservedCounter) Unmarshal() unMarshalPatternFlowGtpv1ReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1ReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1ReservedCounter) ToProto() (*otg.PatternFlowGtpv1ReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1ReservedCounter) FromProto(msg *otg.PatternFlowGtpv1ReservedCounter) (PatternFlowGtpv1ReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1ReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1ReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1ReservedCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv1ReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv1ReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1ReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1ReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1ReservedCounter) Clone() (PatternFlowGtpv1ReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1ReservedCounter()
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

// PatternFlowGtpv1ReservedCounter is integer counter pattern
type PatternFlowGtpv1ReservedCounter interface {
	Validation
	// msg marshals PatternFlowGtpv1ReservedCounter to protobuf object *otg.PatternFlowGtpv1ReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1ReservedCounter
	// setMsg unmarshals PatternFlowGtpv1ReservedCounter from protobuf object *otg.PatternFlowGtpv1ReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1ReservedCounter) PatternFlowGtpv1ReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1ReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1ReservedCounter
	// validate validates PatternFlowGtpv1ReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1ReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv1ReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv1ReservedCounter
	SetStart(value uint32) PatternFlowGtpv1ReservedCounter
	// HasStart checks if Start has been set in PatternFlowGtpv1ReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv1ReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv1ReservedCounter
	SetStep(value uint32) PatternFlowGtpv1ReservedCounter
	// HasStep checks if Step has been set in PatternFlowGtpv1ReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv1ReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv1ReservedCounter
	SetCount(value uint32) PatternFlowGtpv1ReservedCounter
	// HasCount checks if Count has been set in PatternFlowGtpv1ReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1ReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1ReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv1ReservedCounter object
func (obj *patternFlowGtpv1ReservedCounter) SetStart(value uint32) PatternFlowGtpv1ReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1ReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1ReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv1ReservedCounter object
func (obj *patternFlowGtpv1ReservedCounter) SetStep(value uint32) PatternFlowGtpv1ReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1ReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1ReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv1ReservedCounter object
func (obj *patternFlowGtpv1ReservedCounter) SetCount(value uint32) PatternFlowGtpv1ReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv1ReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ReservedCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ReservedCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ReservedCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv1ReservedCounter) setDefault() {
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
