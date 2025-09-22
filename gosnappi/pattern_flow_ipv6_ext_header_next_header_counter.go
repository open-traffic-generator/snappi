package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6ExtHeaderNextHeaderCounter *****
type patternFlowIpv6ExtHeaderNextHeaderCounter struct {
	validation
	obj          *otg.PatternFlowIpv6ExtHeaderNextHeaderCounter
	marshaller   marshalPatternFlowIpv6ExtHeaderNextHeaderCounter
	unMarshaller unMarshalPatternFlowIpv6ExtHeaderNextHeaderCounter
}

func NewPatternFlowIpv6ExtHeaderNextHeaderCounter() PatternFlowIpv6ExtHeaderNextHeaderCounter {
	obj := patternFlowIpv6ExtHeaderNextHeaderCounter{obj: &otg.PatternFlowIpv6ExtHeaderNextHeaderCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) msg() *otg.PatternFlowIpv6ExtHeaderNextHeaderCounter {
	return obj.obj
}

func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) setMsg(msg *otg.PatternFlowIpv6ExtHeaderNextHeaderCounter) PatternFlowIpv6ExtHeaderNextHeaderCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6ExtHeaderNextHeaderCounter struct {
	obj *patternFlowIpv6ExtHeaderNextHeaderCounter
}

type marshalPatternFlowIpv6ExtHeaderNextHeaderCounter interface {
	// ToProto marshals PatternFlowIpv6ExtHeaderNextHeaderCounter to protobuf object *otg.PatternFlowIpv6ExtHeaderNextHeaderCounter
	ToProto() (*otg.PatternFlowIpv6ExtHeaderNextHeaderCounter, error)
	// ToPbText marshals PatternFlowIpv6ExtHeaderNextHeaderCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6ExtHeaderNextHeaderCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6ExtHeaderNextHeaderCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6ExtHeaderNextHeaderCounter struct {
	obj *patternFlowIpv6ExtHeaderNextHeaderCounter
}

type unMarshalPatternFlowIpv6ExtHeaderNextHeaderCounter interface {
	// FromProto unmarshals PatternFlowIpv6ExtHeaderNextHeaderCounter from protobuf object *otg.PatternFlowIpv6ExtHeaderNextHeaderCounter
	FromProto(msg *otg.PatternFlowIpv6ExtHeaderNextHeaderCounter) (PatternFlowIpv6ExtHeaderNextHeaderCounter, error)
	// FromPbText unmarshals PatternFlowIpv6ExtHeaderNextHeaderCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6ExtHeaderNextHeaderCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6ExtHeaderNextHeaderCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) Marshal() marshalPatternFlowIpv6ExtHeaderNextHeaderCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6ExtHeaderNextHeaderCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) Unmarshal() unMarshalPatternFlowIpv6ExtHeaderNextHeaderCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6ExtHeaderNextHeaderCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6ExtHeaderNextHeaderCounter) ToProto() (*otg.PatternFlowIpv6ExtHeaderNextHeaderCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6ExtHeaderNextHeaderCounter) FromProto(msg *otg.PatternFlowIpv6ExtHeaderNextHeaderCounter) (PatternFlowIpv6ExtHeaderNextHeaderCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6ExtHeaderNextHeaderCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6ExtHeaderNextHeaderCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6ExtHeaderNextHeaderCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6ExtHeaderNextHeaderCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6ExtHeaderNextHeaderCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6ExtHeaderNextHeaderCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) Clone() (PatternFlowIpv6ExtHeaderNextHeaderCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6ExtHeaderNextHeaderCounter()
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

// PatternFlowIpv6ExtHeaderNextHeaderCounter is integer counter pattern
type PatternFlowIpv6ExtHeaderNextHeaderCounter interface {
	Validation
	// msg marshals PatternFlowIpv6ExtHeaderNextHeaderCounter to protobuf object *otg.PatternFlowIpv6ExtHeaderNextHeaderCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6ExtHeaderNextHeaderCounter
	// setMsg unmarshals PatternFlowIpv6ExtHeaderNextHeaderCounter from protobuf object *otg.PatternFlowIpv6ExtHeaderNextHeaderCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6ExtHeaderNextHeaderCounter) PatternFlowIpv6ExtHeaderNextHeaderCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6ExtHeaderNextHeaderCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6ExtHeaderNextHeaderCounter
	// validate validates PatternFlowIpv6ExtHeaderNextHeaderCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6ExtHeaderNextHeaderCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6ExtHeaderNextHeaderCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6ExtHeaderNextHeaderCounter
	SetStart(value uint32) PatternFlowIpv6ExtHeaderNextHeaderCounter
	// HasStart checks if Start has been set in PatternFlowIpv6ExtHeaderNextHeaderCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6ExtHeaderNextHeaderCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6ExtHeaderNextHeaderCounter
	SetStep(value uint32) PatternFlowIpv6ExtHeaderNextHeaderCounter
	// HasStep checks if Step has been set in PatternFlowIpv6ExtHeaderNextHeaderCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6ExtHeaderNextHeaderCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6ExtHeaderNextHeaderCounter
	SetCount(value uint32) PatternFlowIpv6ExtHeaderNextHeaderCounter
	// HasCount checks if Count has been set in PatternFlowIpv6ExtHeaderNextHeaderCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6ExtHeaderNextHeaderCounter object
func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) SetStart(value uint32) PatternFlowIpv6ExtHeaderNextHeaderCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6ExtHeaderNextHeaderCounter object
func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) SetStep(value uint32) PatternFlowIpv6ExtHeaderNextHeaderCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6ExtHeaderNextHeaderCounter object
func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) SetCount(value uint32) PatternFlowIpv6ExtHeaderNextHeaderCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6ExtHeaderNextHeaderCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6ExtHeaderNextHeaderCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6ExtHeaderNextHeaderCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6ExtHeaderNextHeaderCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(59)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
