package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TotalLengthCounter *****
type patternFlowIpv4TotalLengthCounter struct {
	validation
	obj          *otg.PatternFlowIpv4TotalLengthCounter
	marshaller   marshalPatternFlowIpv4TotalLengthCounter
	unMarshaller unMarshalPatternFlowIpv4TotalLengthCounter
}

func NewPatternFlowIpv4TotalLengthCounter() PatternFlowIpv4TotalLengthCounter {
	obj := patternFlowIpv4TotalLengthCounter{obj: &otg.PatternFlowIpv4TotalLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TotalLengthCounter) msg() *otg.PatternFlowIpv4TotalLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TotalLengthCounter) setMsg(msg *otg.PatternFlowIpv4TotalLengthCounter) PatternFlowIpv4TotalLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TotalLengthCounter struct {
	obj *patternFlowIpv4TotalLengthCounter
}

type marshalPatternFlowIpv4TotalLengthCounter interface {
	// ToProto marshals PatternFlowIpv4TotalLengthCounter to protobuf object *otg.PatternFlowIpv4TotalLengthCounter
	ToProto() (*otg.PatternFlowIpv4TotalLengthCounter, error)
	// ToPbText marshals PatternFlowIpv4TotalLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TotalLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TotalLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TotalLengthCounter struct {
	obj *patternFlowIpv4TotalLengthCounter
}

type unMarshalPatternFlowIpv4TotalLengthCounter interface {
	// FromProto unmarshals PatternFlowIpv4TotalLengthCounter from protobuf object *otg.PatternFlowIpv4TotalLengthCounter
	FromProto(msg *otg.PatternFlowIpv4TotalLengthCounter) (PatternFlowIpv4TotalLengthCounter, error)
	// FromPbText unmarshals PatternFlowIpv4TotalLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TotalLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TotalLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TotalLengthCounter) Marshal() marshalPatternFlowIpv4TotalLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TotalLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TotalLengthCounter) Unmarshal() unMarshalPatternFlowIpv4TotalLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TotalLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TotalLengthCounter) ToProto() (*otg.PatternFlowIpv4TotalLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TotalLengthCounter) FromProto(msg *otg.PatternFlowIpv4TotalLengthCounter) (PatternFlowIpv4TotalLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TotalLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TotalLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TotalLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TotalLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TotalLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TotalLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4TotalLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TotalLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TotalLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TotalLengthCounter) Clone() (PatternFlowIpv4TotalLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TotalLengthCounter()
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

// PatternFlowIpv4TotalLengthCounter is integer counter pattern
type PatternFlowIpv4TotalLengthCounter interface {
	Validation
	// msg marshals PatternFlowIpv4TotalLengthCounter to protobuf object *otg.PatternFlowIpv4TotalLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TotalLengthCounter
	// setMsg unmarshals PatternFlowIpv4TotalLengthCounter from protobuf object *otg.PatternFlowIpv4TotalLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TotalLengthCounter) PatternFlowIpv4TotalLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TotalLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TotalLengthCounter
	// validate validates PatternFlowIpv4TotalLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TotalLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4TotalLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4TotalLengthCounter
	SetStart(value uint32) PatternFlowIpv4TotalLengthCounter
	// HasStart checks if Start has been set in PatternFlowIpv4TotalLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4TotalLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4TotalLengthCounter
	SetStep(value uint32) PatternFlowIpv4TotalLengthCounter
	// HasStep checks if Step has been set in PatternFlowIpv4TotalLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4TotalLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4TotalLengthCounter
	SetCount(value uint32) PatternFlowIpv4TotalLengthCounter
	// HasCount checks if Count has been set in PatternFlowIpv4TotalLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TotalLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TotalLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4TotalLengthCounter object
func (obj *patternFlowIpv4TotalLengthCounter) SetStart(value uint32) PatternFlowIpv4TotalLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TotalLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TotalLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4TotalLengthCounter object
func (obj *patternFlowIpv4TotalLengthCounter) SetStep(value uint32) PatternFlowIpv4TotalLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TotalLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TotalLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4TotalLengthCounter object
func (obj *patternFlowIpv4TotalLengthCounter) SetCount(value uint32) PatternFlowIpv4TotalLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4TotalLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TotalLengthCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TotalLengthCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TotalLengthCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4TotalLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(46)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
