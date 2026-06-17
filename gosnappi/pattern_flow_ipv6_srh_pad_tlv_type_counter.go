package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPadTlvTypeCounter *****
type patternFlowIpv6SRHPadTlvTypeCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHPadTlvTypeCounter
	marshaller   marshalPatternFlowIpv6SRHPadTlvTypeCounter
	unMarshaller unMarshalPatternFlowIpv6SRHPadTlvTypeCounter
}

func NewPatternFlowIpv6SRHPadTlvTypeCounter() PatternFlowIpv6SRHPadTlvTypeCounter {
	obj := patternFlowIpv6SRHPadTlvTypeCounter{obj: &otg.PatternFlowIpv6SRHPadTlvTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPadTlvTypeCounter) msg() *otg.PatternFlowIpv6SRHPadTlvTypeCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPadTlvTypeCounter) setMsg(msg *otg.PatternFlowIpv6SRHPadTlvTypeCounter) PatternFlowIpv6SRHPadTlvTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPadTlvTypeCounter struct {
	obj *patternFlowIpv6SRHPadTlvTypeCounter
}

type marshalPatternFlowIpv6SRHPadTlvTypeCounter interface {
	// ToProto marshals PatternFlowIpv6SRHPadTlvTypeCounter to protobuf object *otg.PatternFlowIpv6SRHPadTlvTypeCounter
	ToProto() (*otg.PatternFlowIpv6SRHPadTlvTypeCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHPadTlvTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPadTlvTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPadTlvTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPadTlvTypeCounter struct {
	obj *patternFlowIpv6SRHPadTlvTypeCounter
}

type unMarshalPatternFlowIpv6SRHPadTlvTypeCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHPadTlvTypeCounter from protobuf object *otg.PatternFlowIpv6SRHPadTlvTypeCounter
	FromProto(msg *otg.PatternFlowIpv6SRHPadTlvTypeCounter) (PatternFlowIpv6SRHPadTlvTypeCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPadTlvTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPadTlvTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPadTlvTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPadTlvTypeCounter) Marshal() marshalPatternFlowIpv6SRHPadTlvTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPadTlvTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPadTlvTypeCounter) Unmarshal() unMarshalPatternFlowIpv6SRHPadTlvTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPadTlvTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPadTlvTypeCounter) ToProto() (*otg.PatternFlowIpv6SRHPadTlvTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPadTlvTypeCounter) FromProto(msg *otg.PatternFlowIpv6SRHPadTlvTypeCounter) (PatternFlowIpv6SRHPadTlvTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPadTlvTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPadTlvTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPadTlvTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPadTlvTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPadTlvTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPadTlvTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPadTlvTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPadTlvTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPadTlvTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPadTlvTypeCounter) Clone() (PatternFlowIpv6SRHPadTlvTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPadTlvTypeCounter()
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

// PatternFlowIpv6SRHPadTlvTypeCounter is integer counter pattern
type PatternFlowIpv6SRHPadTlvTypeCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPadTlvTypeCounter to protobuf object *otg.PatternFlowIpv6SRHPadTlvTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPadTlvTypeCounter
	// setMsg unmarshals PatternFlowIpv6SRHPadTlvTypeCounter from protobuf object *otg.PatternFlowIpv6SRHPadTlvTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPadTlvTypeCounter) PatternFlowIpv6SRHPadTlvTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPadTlvTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPadTlvTypeCounter
	// validate validates PatternFlowIpv6SRHPadTlvTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPadTlvTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHPadTlvTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHPadTlvTypeCounter
	SetStart(value uint32) PatternFlowIpv6SRHPadTlvTypeCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHPadTlvTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHPadTlvTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHPadTlvTypeCounter
	SetStep(value uint32) PatternFlowIpv6SRHPadTlvTypeCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHPadTlvTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHPadTlvTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHPadTlvTypeCounter
	SetCount(value uint32) PatternFlowIpv6SRHPadTlvTypeCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHPadTlvTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPadTlvTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPadTlvTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHPadTlvTypeCounter object
func (obj *patternFlowIpv6SRHPadTlvTypeCounter) SetStart(value uint32) PatternFlowIpv6SRHPadTlvTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPadTlvTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPadTlvTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHPadTlvTypeCounter object
func (obj *patternFlowIpv6SRHPadTlvTypeCounter) SetStep(value uint32) PatternFlowIpv6SRHPadTlvTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPadTlvTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPadTlvTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHPadTlvTypeCounter object
func (obj *patternFlowIpv6SRHPadTlvTypeCounter) SetCount(value uint32) PatternFlowIpv6SRHPadTlvTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHPadTlvTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPadTlvTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPadTlvTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPadTlvTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SRHPadTlvTypeCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(4)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
