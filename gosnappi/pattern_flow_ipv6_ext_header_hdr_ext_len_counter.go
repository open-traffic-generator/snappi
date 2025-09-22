package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6ExtHeaderHdrExtLenCounter *****
type patternFlowIpv6ExtHeaderHdrExtLenCounter struct {
	validation
	obj          *otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter
	marshaller   marshalPatternFlowIpv6ExtHeaderHdrExtLenCounter
	unMarshaller unMarshalPatternFlowIpv6ExtHeaderHdrExtLenCounter
}

func NewPatternFlowIpv6ExtHeaderHdrExtLenCounter() PatternFlowIpv6ExtHeaderHdrExtLenCounter {
	obj := patternFlowIpv6ExtHeaderHdrExtLenCounter{obj: &otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) msg() *otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter {
	return obj.obj
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) setMsg(msg *otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter) PatternFlowIpv6ExtHeaderHdrExtLenCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6ExtHeaderHdrExtLenCounter struct {
	obj *patternFlowIpv6ExtHeaderHdrExtLenCounter
}

type marshalPatternFlowIpv6ExtHeaderHdrExtLenCounter interface {
	// ToProto marshals PatternFlowIpv6ExtHeaderHdrExtLenCounter to protobuf object *otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter
	ToProto() (*otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter, error)
	// ToPbText marshals PatternFlowIpv6ExtHeaderHdrExtLenCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6ExtHeaderHdrExtLenCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6ExtHeaderHdrExtLenCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6ExtHeaderHdrExtLenCounter struct {
	obj *patternFlowIpv6ExtHeaderHdrExtLenCounter
}

type unMarshalPatternFlowIpv6ExtHeaderHdrExtLenCounter interface {
	// FromProto unmarshals PatternFlowIpv6ExtHeaderHdrExtLenCounter from protobuf object *otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter
	FromProto(msg *otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter) (PatternFlowIpv6ExtHeaderHdrExtLenCounter, error)
	// FromPbText unmarshals PatternFlowIpv6ExtHeaderHdrExtLenCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6ExtHeaderHdrExtLenCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6ExtHeaderHdrExtLenCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) Marshal() marshalPatternFlowIpv6ExtHeaderHdrExtLenCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6ExtHeaderHdrExtLenCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) Unmarshal() unMarshalPatternFlowIpv6ExtHeaderHdrExtLenCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6ExtHeaderHdrExtLenCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6ExtHeaderHdrExtLenCounter) ToProto() (*otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6ExtHeaderHdrExtLenCounter) FromProto(msg *otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter) (PatternFlowIpv6ExtHeaderHdrExtLenCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6ExtHeaderHdrExtLenCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6ExtHeaderHdrExtLenCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6ExtHeaderHdrExtLenCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6ExtHeaderHdrExtLenCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6ExtHeaderHdrExtLenCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6ExtHeaderHdrExtLenCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) Clone() (PatternFlowIpv6ExtHeaderHdrExtLenCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6ExtHeaderHdrExtLenCounter()
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

// PatternFlowIpv6ExtHeaderHdrExtLenCounter is integer counter pattern
type PatternFlowIpv6ExtHeaderHdrExtLenCounter interface {
	Validation
	// msg marshals PatternFlowIpv6ExtHeaderHdrExtLenCounter to protobuf object *otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter
	// setMsg unmarshals PatternFlowIpv6ExtHeaderHdrExtLenCounter from protobuf object *otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6ExtHeaderHdrExtLenCounter) PatternFlowIpv6ExtHeaderHdrExtLenCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6ExtHeaderHdrExtLenCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6ExtHeaderHdrExtLenCounter
	// validate validates PatternFlowIpv6ExtHeaderHdrExtLenCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6ExtHeaderHdrExtLenCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6ExtHeaderHdrExtLenCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6ExtHeaderHdrExtLenCounter
	SetStart(value uint32) PatternFlowIpv6ExtHeaderHdrExtLenCounter
	// HasStart checks if Start has been set in PatternFlowIpv6ExtHeaderHdrExtLenCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6ExtHeaderHdrExtLenCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6ExtHeaderHdrExtLenCounter
	SetStep(value uint32) PatternFlowIpv6ExtHeaderHdrExtLenCounter
	// HasStep checks if Step has been set in PatternFlowIpv6ExtHeaderHdrExtLenCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6ExtHeaderHdrExtLenCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6ExtHeaderHdrExtLenCounter
	SetCount(value uint32) PatternFlowIpv6ExtHeaderHdrExtLenCounter
	// HasCount checks if Count has been set in PatternFlowIpv6ExtHeaderHdrExtLenCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6ExtHeaderHdrExtLenCounter object
func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) SetStart(value uint32) PatternFlowIpv6ExtHeaderHdrExtLenCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6ExtHeaderHdrExtLenCounter object
func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) SetStep(value uint32) PatternFlowIpv6ExtHeaderHdrExtLenCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6ExtHeaderHdrExtLenCounter object
func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) SetCount(value uint32) PatternFlowIpv6ExtHeaderHdrExtLenCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6ExtHeaderHdrExtLenCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6ExtHeaderHdrExtLenCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6ExtHeaderHdrExtLenCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6ExtHeaderHdrExtLenCounter) setDefault() {
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
