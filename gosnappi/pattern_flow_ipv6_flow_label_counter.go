package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6FlowLabelCounter *****
type patternFlowIpv6FlowLabelCounter struct {
	validation
	obj          *otg.PatternFlowIpv6FlowLabelCounter
	marshaller   marshalPatternFlowIpv6FlowLabelCounter
	unMarshaller unMarshalPatternFlowIpv6FlowLabelCounter
}

func NewPatternFlowIpv6FlowLabelCounter() PatternFlowIpv6FlowLabelCounter {
	obj := patternFlowIpv6FlowLabelCounter{obj: &otg.PatternFlowIpv6FlowLabelCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6FlowLabelCounter) msg() *otg.PatternFlowIpv6FlowLabelCounter {
	return obj.obj
}

func (obj *patternFlowIpv6FlowLabelCounter) setMsg(msg *otg.PatternFlowIpv6FlowLabelCounter) PatternFlowIpv6FlowLabelCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6FlowLabelCounter struct {
	obj *patternFlowIpv6FlowLabelCounter
}

type marshalPatternFlowIpv6FlowLabelCounter interface {
	// ToProto marshals PatternFlowIpv6FlowLabelCounter to protobuf object *otg.PatternFlowIpv6FlowLabelCounter
	ToProto() (*otg.PatternFlowIpv6FlowLabelCounter, error)
	// ToPbText marshals PatternFlowIpv6FlowLabelCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6FlowLabelCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6FlowLabelCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv6FlowLabelCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv6FlowLabelCounter struct {
	obj *patternFlowIpv6FlowLabelCounter
}

type unMarshalPatternFlowIpv6FlowLabelCounter interface {
	// FromProto unmarshals PatternFlowIpv6FlowLabelCounter from protobuf object *otg.PatternFlowIpv6FlowLabelCounter
	FromProto(msg *otg.PatternFlowIpv6FlowLabelCounter) (PatternFlowIpv6FlowLabelCounter, error)
	// FromPbText unmarshals PatternFlowIpv6FlowLabelCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6FlowLabelCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6FlowLabelCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6FlowLabelCounter) Marshal() marshalPatternFlowIpv6FlowLabelCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6FlowLabelCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6FlowLabelCounter) Unmarshal() unMarshalPatternFlowIpv6FlowLabelCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6FlowLabelCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6FlowLabelCounter) ToProto() (*otg.PatternFlowIpv6FlowLabelCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6FlowLabelCounter) FromProto(msg *otg.PatternFlowIpv6FlowLabelCounter) (PatternFlowIpv6FlowLabelCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6FlowLabelCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6FlowLabelCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6FlowLabelCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6FlowLabelCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6FlowLabelCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv6FlowLabelCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6FlowLabelCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6FlowLabelCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6FlowLabelCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6FlowLabelCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6FlowLabelCounter) Clone() (PatternFlowIpv6FlowLabelCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6FlowLabelCounter()
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

// PatternFlowIpv6FlowLabelCounter is integer counter pattern
type PatternFlowIpv6FlowLabelCounter interface {
	Validation
	// msg marshals PatternFlowIpv6FlowLabelCounter to protobuf object *otg.PatternFlowIpv6FlowLabelCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6FlowLabelCounter
	// setMsg unmarshals PatternFlowIpv6FlowLabelCounter from protobuf object *otg.PatternFlowIpv6FlowLabelCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6FlowLabelCounter) PatternFlowIpv6FlowLabelCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6FlowLabelCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6FlowLabelCounter
	// validate validates PatternFlowIpv6FlowLabelCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6FlowLabelCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6FlowLabelCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6FlowLabelCounter
	SetStart(value uint32) PatternFlowIpv6FlowLabelCounter
	// HasStart checks if Start has been set in PatternFlowIpv6FlowLabelCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6FlowLabelCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6FlowLabelCounter
	SetStep(value uint32) PatternFlowIpv6FlowLabelCounter
	// HasStep checks if Step has been set in PatternFlowIpv6FlowLabelCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6FlowLabelCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6FlowLabelCounter
	SetCount(value uint32) PatternFlowIpv6FlowLabelCounter
	// HasCount checks if Count has been set in PatternFlowIpv6FlowLabelCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6FlowLabelCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6FlowLabelCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6FlowLabelCounter object
func (obj *patternFlowIpv6FlowLabelCounter) SetStart(value uint32) PatternFlowIpv6FlowLabelCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6FlowLabelCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6FlowLabelCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6FlowLabelCounter object
func (obj *patternFlowIpv6FlowLabelCounter) SetStep(value uint32) PatternFlowIpv6FlowLabelCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6FlowLabelCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6FlowLabelCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6FlowLabelCounter object
func (obj *patternFlowIpv6FlowLabelCounter) SetCount(value uint32) PatternFlowIpv6FlowLabelCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6FlowLabelCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6FlowLabelCounter.Start <= 1048575 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6FlowLabelCounter.Step <= 1048575 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6FlowLabelCounter.Count <= 1048575 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6FlowLabelCounter) setDefault() {
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
