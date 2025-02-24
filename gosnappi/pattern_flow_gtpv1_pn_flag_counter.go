package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1PnFlagCounter *****
type patternFlowGtpv1PnFlagCounter struct {
	validation
	obj          *otg.PatternFlowGtpv1PnFlagCounter
	marshaller   marshalPatternFlowGtpv1PnFlagCounter
	unMarshaller unMarshalPatternFlowGtpv1PnFlagCounter
}

func NewPatternFlowGtpv1PnFlagCounter() PatternFlowGtpv1PnFlagCounter {
	obj := patternFlowGtpv1PnFlagCounter{obj: &otg.PatternFlowGtpv1PnFlagCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1PnFlagCounter) msg() *otg.PatternFlowGtpv1PnFlagCounter {
	return obj.obj
}

func (obj *patternFlowGtpv1PnFlagCounter) setMsg(msg *otg.PatternFlowGtpv1PnFlagCounter) PatternFlowGtpv1PnFlagCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1PnFlagCounter struct {
	obj *patternFlowGtpv1PnFlagCounter
}

type marshalPatternFlowGtpv1PnFlagCounter interface {
	// ToProto marshals PatternFlowGtpv1PnFlagCounter to protobuf object *otg.PatternFlowGtpv1PnFlagCounter
	ToProto() (*otg.PatternFlowGtpv1PnFlagCounter, error)
	// ToPbText marshals PatternFlowGtpv1PnFlagCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1PnFlagCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1PnFlagCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1PnFlagCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1PnFlagCounter struct {
	obj *patternFlowGtpv1PnFlagCounter
}

type unMarshalPatternFlowGtpv1PnFlagCounter interface {
	// FromProto unmarshals PatternFlowGtpv1PnFlagCounter from protobuf object *otg.PatternFlowGtpv1PnFlagCounter
	FromProto(msg *otg.PatternFlowGtpv1PnFlagCounter) (PatternFlowGtpv1PnFlagCounter, error)
	// FromPbText unmarshals PatternFlowGtpv1PnFlagCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1PnFlagCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1PnFlagCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1PnFlagCounter) Marshal() marshalPatternFlowGtpv1PnFlagCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1PnFlagCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1PnFlagCounter) Unmarshal() unMarshalPatternFlowGtpv1PnFlagCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1PnFlagCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1PnFlagCounter) ToProto() (*otg.PatternFlowGtpv1PnFlagCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1PnFlagCounter) FromProto(msg *otg.PatternFlowGtpv1PnFlagCounter) (PatternFlowGtpv1PnFlagCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1PnFlagCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1PnFlagCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1PnFlagCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1PnFlagCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1PnFlagCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv1PnFlagCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1PnFlagCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv1PnFlagCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1PnFlagCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1PnFlagCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1PnFlagCounter) Clone() (PatternFlowGtpv1PnFlagCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1PnFlagCounter()
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

// PatternFlowGtpv1PnFlagCounter is integer counter pattern
type PatternFlowGtpv1PnFlagCounter interface {
	Validation
	// msg marshals PatternFlowGtpv1PnFlagCounter to protobuf object *otg.PatternFlowGtpv1PnFlagCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1PnFlagCounter
	// setMsg unmarshals PatternFlowGtpv1PnFlagCounter from protobuf object *otg.PatternFlowGtpv1PnFlagCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1PnFlagCounter) PatternFlowGtpv1PnFlagCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1PnFlagCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1PnFlagCounter
	// validate validates PatternFlowGtpv1PnFlagCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1PnFlagCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv1PnFlagCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv1PnFlagCounter
	SetStart(value uint32) PatternFlowGtpv1PnFlagCounter
	// HasStart checks if Start has been set in PatternFlowGtpv1PnFlagCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv1PnFlagCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv1PnFlagCounter
	SetStep(value uint32) PatternFlowGtpv1PnFlagCounter
	// HasStep checks if Step has been set in PatternFlowGtpv1PnFlagCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv1PnFlagCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv1PnFlagCounter
	SetCount(value uint32) PatternFlowGtpv1PnFlagCounter
	// HasCount checks if Count has been set in PatternFlowGtpv1PnFlagCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1PnFlagCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv1PnFlagCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv1PnFlagCounter object
func (obj *patternFlowGtpv1PnFlagCounter) SetStart(value uint32) PatternFlowGtpv1PnFlagCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1PnFlagCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv1PnFlagCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv1PnFlagCounter object
func (obj *patternFlowGtpv1PnFlagCounter) SetStep(value uint32) PatternFlowGtpv1PnFlagCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1PnFlagCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv1PnFlagCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv1PnFlagCounter object
func (obj *patternFlowGtpv1PnFlagCounter) SetCount(value uint32) PatternFlowGtpv1PnFlagCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv1PnFlagCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1PnFlagCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1PnFlagCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1PnFlagCounter.Count <= 1 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv1PnFlagCounter) setDefault() {
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
