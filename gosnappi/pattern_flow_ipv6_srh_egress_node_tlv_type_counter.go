package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHEgressNodeTlvTypeCounter *****
type patternFlowIpv6SRHEgressNodeTlvTypeCounter struct {
	validation
	obj          *otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	marshaller   marshalPatternFlowIpv6SRHEgressNodeTlvTypeCounter
	unMarshaller unMarshalPatternFlowIpv6SRHEgressNodeTlvTypeCounter
}

func NewPatternFlowIpv6SRHEgressNodeTlvTypeCounter() PatternFlowIpv6SRHEgressNodeTlvTypeCounter {
	obj := patternFlowIpv6SRHEgressNodeTlvTypeCounter{obj: &otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) msg() *otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter {
	return obj.obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) setMsg(msg *otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter) PatternFlowIpv6SRHEgressNodeTlvTypeCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHEgressNodeTlvTypeCounter struct {
	obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter
}

type marshalPatternFlowIpv6SRHEgressNodeTlvTypeCounter interface {
	// ToProto marshals PatternFlowIpv6SRHEgressNodeTlvTypeCounter to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter, error)
	// ToPbText marshals PatternFlowIpv6SRHEgressNodeTlvTypeCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHEgressNodeTlvTypeCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHEgressNodeTlvTypeCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHEgressNodeTlvTypeCounter struct {
	obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter
}

type unMarshalPatternFlowIpv6SRHEgressNodeTlvTypeCounter interface {
	// FromProto unmarshals PatternFlowIpv6SRHEgressNodeTlvTypeCounter from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter) (PatternFlowIpv6SRHEgressNodeTlvTypeCounter, error)
	// FromPbText unmarshals PatternFlowIpv6SRHEgressNodeTlvTypeCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHEgressNodeTlvTypeCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHEgressNodeTlvTypeCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvTypeCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHEgressNodeTlvTypeCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvTypeCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHEgressNodeTlvTypeCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvTypeCounter) ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvTypeCounter) FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter) (PatternFlowIpv6SRHEgressNodeTlvTypeCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvTypeCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvTypeCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvTypeCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvTypeCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvTypeCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvTypeCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) Clone() (PatternFlowIpv6SRHEgressNodeTlvTypeCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHEgressNodeTlvTypeCounter()
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

// PatternFlowIpv6SRHEgressNodeTlvTypeCounter is integer counter pattern
type PatternFlowIpv6SRHEgressNodeTlvTypeCounter interface {
	Validation
	// msg marshals PatternFlowIpv6SRHEgressNodeTlvTypeCounter to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	// setMsg unmarshals PatternFlowIpv6SRHEgressNodeTlvTypeCounter from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHEgressNodeTlvTypeCounter) PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvTypeCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvTypeCounter
	// validate validates PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHEgressNodeTlvTypeCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvTypeCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	SetStart(value uint32) PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	// HasStart checks if Start has been set in PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvTypeCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	SetStep(value uint32) PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	// HasStep checks if Step has been set in PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvTypeCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	SetCount(value uint32) PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	// HasCount checks if Count has been set in PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6SRHEgressNodeTlvTypeCounter object
func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) SetStart(value uint32) PatternFlowIpv6SRHEgressNodeTlvTypeCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6SRHEgressNodeTlvTypeCounter object
func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) SetStep(value uint32) PatternFlowIpv6SRHEgressNodeTlvTypeCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6SRHEgressNodeTlvTypeCounter object
func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) SetCount(value uint32) PatternFlowIpv6SRHEgressNodeTlvTypeCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvTypeCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvTypeCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvTypeCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6SRHEgressNodeTlvTypeCounter) setDefault() {
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
