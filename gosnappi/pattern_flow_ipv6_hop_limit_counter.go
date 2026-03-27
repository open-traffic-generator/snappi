package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6HopLimitCounter *****
type patternFlowIpv6HopLimitCounter struct {
	validation
	obj          *otg.PatternFlowIpv6HopLimitCounter
	marshaller   marshalPatternFlowIpv6HopLimitCounter
	unMarshaller unMarshalPatternFlowIpv6HopLimitCounter
}

func NewPatternFlowIpv6HopLimitCounter() PatternFlowIpv6HopLimitCounter {
	obj := patternFlowIpv6HopLimitCounter{obj: &otg.PatternFlowIpv6HopLimitCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6HopLimitCounter) msg() *otg.PatternFlowIpv6HopLimitCounter {
	return obj.obj
}

func (obj *patternFlowIpv6HopLimitCounter) setMsg(msg *otg.PatternFlowIpv6HopLimitCounter) PatternFlowIpv6HopLimitCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6HopLimitCounter struct {
	obj *patternFlowIpv6HopLimitCounter
}

type marshalPatternFlowIpv6HopLimitCounter interface {
	// ToProto marshals PatternFlowIpv6HopLimitCounter to protobuf object *otg.PatternFlowIpv6HopLimitCounter
	ToProto() (*otg.PatternFlowIpv6HopLimitCounter, error)
	// ToPbText marshals PatternFlowIpv6HopLimitCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6HopLimitCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6HopLimitCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6HopLimitCounter struct {
	obj *patternFlowIpv6HopLimitCounter
}

type unMarshalPatternFlowIpv6HopLimitCounter interface {
	// FromProto unmarshals PatternFlowIpv6HopLimitCounter from protobuf object *otg.PatternFlowIpv6HopLimitCounter
	FromProto(msg *otg.PatternFlowIpv6HopLimitCounter) (PatternFlowIpv6HopLimitCounter, error)
	// FromPbText unmarshals PatternFlowIpv6HopLimitCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6HopLimitCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6HopLimitCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6HopLimitCounter) Marshal() marshalPatternFlowIpv6HopLimitCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6HopLimitCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6HopLimitCounter) Unmarshal() unMarshalPatternFlowIpv6HopLimitCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6HopLimitCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6HopLimitCounter) ToProto() (*otg.PatternFlowIpv6HopLimitCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6HopLimitCounter) FromProto(msg *otg.PatternFlowIpv6HopLimitCounter) (PatternFlowIpv6HopLimitCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6HopLimitCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6HopLimitCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6HopLimitCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6HopLimitCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6HopLimitCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6HopLimitCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6HopLimitCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6HopLimitCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6HopLimitCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6HopLimitCounter) Clone() (PatternFlowIpv6HopLimitCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6HopLimitCounter()
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

// PatternFlowIpv6HopLimitCounter is integer counter pattern
type PatternFlowIpv6HopLimitCounter interface {
	Validation
	// msg marshals PatternFlowIpv6HopLimitCounter to protobuf object *otg.PatternFlowIpv6HopLimitCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6HopLimitCounter
	// setMsg unmarshals PatternFlowIpv6HopLimitCounter from protobuf object *otg.PatternFlowIpv6HopLimitCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6HopLimitCounter) PatternFlowIpv6HopLimitCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6HopLimitCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6HopLimitCounter
	// validate validates PatternFlowIpv6HopLimitCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6HopLimitCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6HopLimitCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6HopLimitCounter
	SetStart(value uint32) PatternFlowIpv6HopLimitCounter
	// HasStart checks if Start has been set in PatternFlowIpv6HopLimitCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6HopLimitCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6HopLimitCounter
	SetStep(value uint32) PatternFlowIpv6HopLimitCounter
	// HasStep checks if Step has been set in PatternFlowIpv6HopLimitCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6HopLimitCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6HopLimitCounter
	SetCount(value uint32) PatternFlowIpv6HopLimitCounter
	// HasCount checks if Count has been set in PatternFlowIpv6HopLimitCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6HopLimitCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6HopLimitCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6HopLimitCounter object
func (obj *patternFlowIpv6HopLimitCounter) SetStart(value uint32) PatternFlowIpv6HopLimitCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6HopLimitCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6HopLimitCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6HopLimitCounter object
func (obj *patternFlowIpv6HopLimitCounter) SetStep(value uint32) PatternFlowIpv6HopLimitCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6HopLimitCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6HopLimitCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6HopLimitCounter object
func (obj *patternFlowIpv6HopLimitCounter) SetCount(value uint32) PatternFlowIpv6HopLimitCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6HopLimitCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6HopLimitCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6HopLimitCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6HopLimitCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6HopLimitCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(64)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
