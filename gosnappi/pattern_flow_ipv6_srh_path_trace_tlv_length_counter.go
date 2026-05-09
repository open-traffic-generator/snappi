package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPathTraceTlvLengthCounter *****
type patternFlowIpv6SRHPathTraceTlvLengthCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter
	marshaller   marshalPatternFlowIpv6SRHPathTraceTlvLengthCounter
	unMarshaller unMarshalPatternFlowIpv6SRHPathTraceTlvLengthCounter
}

func NewPatternFlowIpv6SRHPathTraceTlvLengthCounter() PatternFlowIpv6SRHPathTraceTlvLengthCounter {
	obj := patternFlowIpv6SRHPathTraceTlvLengthCounter{obj: &otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) msg() *otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) setMsg(msg *otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter) PatternFlowIpv6SRHPathTraceTlvLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPathTraceTlvLengthCounter struct {
	obj *patternFlowIpv6SRHPathTraceTlvLengthCounter
}

type marshalPatternFlowIpv6SRHPathTraceTlvLengthCounter interface {
	// ToProto marshals PatternFlowIpv6SRHPathTraceTlvLengthCounter to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter
	ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHPathTraceTlvLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPathTraceTlvLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPathTraceTlvLengthCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPathTraceTlvLengthCounter struct {
	obj *patternFlowIpv6SRHPathTraceTlvLengthCounter
}

type unMarshalPatternFlowIpv6SRHPathTraceTlvLengthCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHPathTraceTlvLengthCounter from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter
	FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter) (PatternFlowIpv6SRHPathTraceTlvLengthCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPathTraceTlvLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPathTraceTlvLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPathTraceTlvLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) Marshal() marshalPatternFlowIpv6SRHPathTraceTlvLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPathTraceTlvLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPathTraceTlvLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvLengthCounter) ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvLengthCounter) FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter) (PatternFlowIpv6SRHPathTraceTlvLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) Clone() (PatternFlowIpv6SRHPathTraceTlvLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPathTraceTlvLengthCounter()
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

// PatternFlowIpv6SRHPathTraceTlvLengthCounter is integer counter pattern
type PatternFlowIpv6SRHPathTraceTlvLengthCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPathTraceTlvLengthCounter to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter
	// setMsg unmarshals PatternFlowIpv6SRHPathTraceTlvLengthCounter from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPathTraceTlvLengthCounter) PatternFlowIpv6SRHPathTraceTlvLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPathTraceTlvLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvLengthCounter
	// validate validates PatternFlowIpv6SRHPathTraceTlvLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPathTraceTlvLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHPathTraceTlvLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvLengthCounter
	SetStart(value uint32) PatternFlowIpv6SRHPathTraceTlvLengthCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHPathTraceTlvLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHPathTraceTlvLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvLengthCounter
	SetStep(value uint32) PatternFlowIpv6SRHPathTraceTlvLengthCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHPathTraceTlvLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHPathTraceTlvLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvLengthCounter
	SetCount(value uint32) PatternFlowIpv6SRHPathTraceTlvLengthCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHPathTraceTlvLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvLengthCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) SetStart(value uint32) PatternFlowIpv6SRHPathTraceTlvLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvLengthCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) SetStep(value uint32) PatternFlowIpv6SRHPathTraceTlvLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvLengthCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) SetCount(value uint32) PatternFlowIpv6SRHPathTraceTlvLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvLengthCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvLengthCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvLengthCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SRHPathTraceTlvLengthCounter) setDefault() {
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
