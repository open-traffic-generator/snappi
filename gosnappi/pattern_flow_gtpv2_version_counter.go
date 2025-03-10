package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2VersionCounter *****
type patternFlowGtpv2VersionCounter struct {
	validation
	obj          *otg.PatternFlowGtpv2VersionCounter
	marshaller   marshalPatternFlowGtpv2VersionCounter
	unMarshaller unMarshalPatternFlowGtpv2VersionCounter
}

func NewPatternFlowGtpv2VersionCounter() PatternFlowGtpv2VersionCounter {
	obj := patternFlowGtpv2VersionCounter{obj: &otg.PatternFlowGtpv2VersionCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2VersionCounter) msg() *otg.PatternFlowGtpv2VersionCounter {
	return obj.obj
}

func (obj *patternFlowGtpv2VersionCounter) setMsg(msg *otg.PatternFlowGtpv2VersionCounter) PatternFlowGtpv2VersionCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2VersionCounter struct {
	obj *patternFlowGtpv2VersionCounter
}

type marshalPatternFlowGtpv2VersionCounter interface {
	// ToProto marshals PatternFlowGtpv2VersionCounter to protobuf object *otg.PatternFlowGtpv2VersionCounter
	ToProto() (*otg.PatternFlowGtpv2VersionCounter, error)
	// ToPbText marshals PatternFlowGtpv2VersionCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2VersionCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2VersionCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv2VersionCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv2VersionCounter struct {
	obj *patternFlowGtpv2VersionCounter
}

type unMarshalPatternFlowGtpv2VersionCounter interface {
	// FromProto unmarshals PatternFlowGtpv2VersionCounter from protobuf object *otg.PatternFlowGtpv2VersionCounter
	FromProto(msg *otg.PatternFlowGtpv2VersionCounter) (PatternFlowGtpv2VersionCounter, error)
	// FromPbText unmarshals PatternFlowGtpv2VersionCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2VersionCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2VersionCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2VersionCounter) Marshal() marshalPatternFlowGtpv2VersionCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2VersionCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2VersionCounter) Unmarshal() unMarshalPatternFlowGtpv2VersionCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2VersionCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2VersionCounter) ToProto() (*otg.PatternFlowGtpv2VersionCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2VersionCounter) FromProto(msg *otg.PatternFlowGtpv2VersionCounter) (PatternFlowGtpv2VersionCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2VersionCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2VersionCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2VersionCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2VersionCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2VersionCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv2VersionCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2VersionCounter) FromJson(value string) error {
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

func (obj *patternFlowGtpv2VersionCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2VersionCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2VersionCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2VersionCounter) Clone() (PatternFlowGtpv2VersionCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2VersionCounter()
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

// PatternFlowGtpv2VersionCounter is integer counter pattern
type PatternFlowGtpv2VersionCounter interface {
	Validation
	// msg marshals PatternFlowGtpv2VersionCounter to protobuf object *otg.PatternFlowGtpv2VersionCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2VersionCounter
	// setMsg unmarshals PatternFlowGtpv2VersionCounter from protobuf object *otg.PatternFlowGtpv2VersionCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2VersionCounter) PatternFlowGtpv2VersionCounter
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2VersionCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2VersionCounter
	// validate validates PatternFlowGtpv2VersionCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2VersionCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowGtpv2VersionCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowGtpv2VersionCounter
	SetStart(value uint32) PatternFlowGtpv2VersionCounter
	// HasStart checks if Start has been set in PatternFlowGtpv2VersionCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowGtpv2VersionCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowGtpv2VersionCounter
	SetStep(value uint32) PatternFlowGtpv2VersionCounter
	// HasStep checks if Step has been set in PatternFlowGtpv2VersionCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowGtpv2VersionCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowGtpv2VersionCounter
	SetCount(value uint32) PatternFlowGtpv2VersionCounter
	// HasCount checks if Count has been set in PatternFlowGtpv2VersionCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2VersionCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowGtpv2VersionCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowGtpv2VersionCounter object
func (obj *patternFlowGtpv2VersionCounter) SetStart(value uint32) PatternFlowGtpv2VersionCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2VersionCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowGtpv2VersionCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowGtpv2VersionCounter object
func (obj *patternFlowGtpv2VersionCounter) SetStep(value uint32) PatternFlowGtpv2VersionCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2VersionCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowGtpv2VersionCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowGtpv2VersionCounter object
func (obj *patternFlowGtpv2VersionCounter) SetCount(value uint32) PatternFlowGtpv2VersionCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowGtpv2VersionCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2VersionCounter.Start <= 7 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2VersionCounter.Step <= 7 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2VersionCounter.Count <= 7 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowGtpv2VersionCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(2)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
