package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHIngressNodeTlvReservedCounter *****
type patternFlowIpv6SRHIngressNodeTlvReservedCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	marshaller   marshalPatternFlowIpv6SRHIngressNodeTlvReservedCounter
	unMarshaller unMarshalPatternFlowIpv6SRHIngressNodeTlvReservedCounter
}

func NewPatternFlowIpv6SRHIngressNodeTlvReservedCounter() PatternFlowIpv6SRHIngressNodeTlvReservedCounter {
	obj := patternFlowIpv6SRHIngressNodeTlvReservedCounter{obj: &otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) msg() *otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) setMsg(msg *otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter) PatternFlowIpv6SRHIngressNodeTlvReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHIngressNodeTlvReservedCounter struct {
	obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter
}

type marshalPatternFlowIpv6SRHIngressNodeTlvReservedCounter interface {
	// ToProto marshals PatternFlowIpv6SRHIngressNodeTlvReservedCounter to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHIngressNodeTlvReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHIngressNodeTlvReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHIngressNodeTlvReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHIngressNodeTlvReservedCounter struct {
	obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter
}

type unMarshalPatternFlowIpv6SRHIngressNodeTlvReservedCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHIngressNodeTlvReservedCounter from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter) (PatternFlowIpv6SRHIngressNodeTlvReservedCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHIngressNodeTlvReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHIngressNodeTlvReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHIngressNodeTlvReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHIngressNodeTlvReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHIngressNodeTlvReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvReservedCounter) ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvReservedCounter) FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter) (PatternFlowIpv6SRHIngressNodeTlvReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) Clone() (PatternFlowIpv6SRHIngressNodeTlvReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHIngressNodeTlvReservedCounter()
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

// PatternFlowIpv6SRHIngressNodeTlvReservedCounter is integer counter pattern
type PatternFlowIpv6SRHIngressNodeTlvReservedCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHIngressNodeTlvReservedCounter to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	// setMsg unmarshals PatternFlowIpv6SRHIngressNodeTlvReservedCounter from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHIngressNodeTlvReservedCounter) PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvReservedCounter
	// validate validates PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHIngressNodeTlvReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	SetStart(value uint32) PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	SetStep(value uint32) PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	SetCount(value uint32) PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHIngressNodeTlvReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHIngressNodeTlvReservedCounter object
func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) SetStart(value uint32) PatternFlowIpv6SRHIngressNodeTlvReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHIngressNodeTlvReservedCounter object
func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) SetStep(value uint32) PatternFlowIpv6SRHIngressNodeTlvReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHIngressNodeTlvReservedCounter object
func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) SetCount(value uint32) PatternFlowIpv6SRHIngressNodeTlvReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvReservedCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvReservedCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvReservedCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SRHIngressNodeTlvReservedCounter) setDefault() {
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
