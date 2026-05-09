package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPathTraceTlvReservedCounter *****
type patternFlowIpv6SRHPathTraceTlvReservedCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter
	marshaller   marshalPatternFlowIpv6SRHPathTraceTlvReservedCounter
	unMarshaller unMarshalPatternFlowIpv6SRHPathTraceTlvReservedCounter
}

func NewPatternFlowIpv6SRHPathTraceTlvReservedCounter() PatternFlowIpv6SRHPathTraceTlvReservedCounter {
	obj := patternFlowIpv6SRHPathTraceTlvReservedCounter{obj: &otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) msg() *otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) setMsg(msg *otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter) PatternFlowIpv6SRHPathTraceTlvReservedCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPathTraceTlvReservedCounter struct {
	obj *patternFlowIpv6SRHPathTraceTlvReservedCounter
}

type marshalPatternFlowIpv6SRHPathTraceTlvReservedCounter interface {
	// ToProto marshals PatternFlowIpv6SRHPathTraceTlvReservedCounter to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter
	ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHPathTraceTlvReservedCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPathTraceTlvReservedCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPathTraceTlvReservedCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPathTraceTlvReservedCounter struct {
	obj *patternFlowIpv6SRHPathTraceTlvReservedCounter
}

type unMarshalPatternFlowIpv6SRHPathTraceTlvReservedCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHPathTraceTlvReservedCounter from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter
	FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter) (PatternFlowIpv6SRHPathTraceTlvReservedCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPathTraceTlvReservedCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPathTraceTlvReservedCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPathTraceTlvReservedCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) Marshal() marshalPatternFlowIpv6SRHPathTraceTlvReservedCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPathTraceTlvReservedCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvReservedCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPathTraceTlvReservedCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvReservedCounter) ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvReservedCounter) FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter) (PatternFlowIpv6SRHPathTraceTlvReservedCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvReservedCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvReservedCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvReservedCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvReservedCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvReservedCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvReservedCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) Clone() (PatternFlowIpv6SRHPathTraceTlvReservedCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPathTraceTlvReservedCounter()
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

// PatternFlowIpv6SRHPathTraceTlvReservedCounter is integer counter pattern
type PatternFlowIpv6SRHPathTraceTlvReservedCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPathTraceTlvReservedCounter to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter
	// setMsg unmarshals PatternFlowIpv6SRHPathTraceTlvReservedCounter from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPathTraceTlvReservedCounter) PatternFlowIpv6SRHPathTraceTlvReservedCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPathTraceTlvReservedCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvReservedCounter
	// validate validates PatternFlowIpv6SRHPathTraceTlvReservedCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPathTraceTlvReservedCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHPathTraceTlvReservedCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvReservedCounter
	SetStart(value uint32) PatternFlowIpv6SRHPathTraceTlvReservedCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHPathTraceTlvReservedCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHPathTraceTlvReservedCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvReservedCounter
	SetStep(value uint32) PatternFlowIpv6SRHPathTraceTlvReservedCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHPathTraceTlvReservedCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHPathTraceTlvReservedCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvReservedCounter
	SetCount(value uint32) PatternFlowIpv6SRHPathTraceTlvReservedCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHPathTraceTlvReservedCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvReservedCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) SetStart(value uint32) PatternFlowIpv6SRHPathTraceTlvReservedCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvReservedCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) SetStep(value uint32) PatternFlowIpv6SRHPathTraceTlvReservedCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvReservedCounter object
func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) SetCount(value uint32) PatternFlowIpv6SRHPathTraceTlvReservedCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvReservedCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvReservedCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvReservedCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SRHPathTraceTlvReservedCounter) setDefault() {
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
