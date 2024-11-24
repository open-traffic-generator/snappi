package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6PayloadLengthCounter *****
type patternFlowIpv6PayloadLengthCounter struct {
	validation
	obj          *otg.PatternFlowIpv6PayloadLengthCounter
	marshaller   marshalPatternFlowIpv6PayloadLengthCounter
	unMarshaller unMarshalPatternFlowIpv6PayloadLengthCounter
}

func NewPatternFlowIpv6PayloadLengthCounter() PatternFlowIpv6PayloadLengthCounter {
	obj := patternFlowIpv6PayloadLengthCounter{obj: &otg.PatternFlowIpv6PayloadLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6PayloadLengthCounter) msg() *otg.PatternFlowIpv6PayloadLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv6PayloadLengthCounter) setMsg(msg *otg.PatternFlowIpv6PayloadLengthCounter) PatternFlowIpv6PayloadLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6PayloadLengthCounter struct {
	obj *patternFlowIpv6PayloadLengthCounter
}

type marshalPatternFlowIpv6PayloadLengthCounter interface {
	// ToProto marshals PatternFlowIpv6PayloadLengthCounter to protobuf object *otg.PatternFlowIpv6PayloadLengthCounter
	ToProto() (*otg.PatternFlowIpv6PayloadLengthCounter, error)
	// ToPbText marshals PatternFlowIpv6PayloadLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6PayloadLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6PayloadLengthCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv6PayloadLengthCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv6PayloadLengthCounter struct {
	obj *patternFlowIpv6PayloadLengthCounter
}

type unMarshalPatternFlowIpv6PayloadLengthCounter interface {
	// FromProto unmarshals PatternFlowIpv6PayloadLengthCounter from protobuf object *otg.PatternFlowIpv6PayloadLengthCounter
	FromProto(msg *otg.PatternFlowIpv6PayloadLengthCounter) (PatternFlowIpv6PayloadLengthCounter, error)
	// FromPbText unmarshals PatternFlowIpv6PayloadLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6PayloadLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6PayloadLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6PayloadLengthCounter) Marshal() marshalPatternFlowIpv6PayloadLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6PayloadLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6PayloadLengthCounter) Unmarshal() unMarshalPatternFlowIpv6PayloadLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6PayloadLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6PayloadLengthCounter) ToProto() (*otg.PatternFlowIpv6PayloadLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6PayloadLengthCounter) FromProto(msg *otg.PatternFlowIpv6PayloadLengthCounter) (PatternFlowIpv6PayloadLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6PayloadLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6PayloadLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6PayloadLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6PayloadLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6PayloadLengthCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv6PayloadLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6PayloadLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6PayloadLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6PayloadLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6PayloadLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6PayloadLengthCounter) Clone() (PatternFlowIpv6PayloadLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6PayloadLengthCounter()
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

// PatternFlowIpv6PayloadLengthCounter is integer counter pattern
type PatternFlowIpv6PayloadLengthCounter interface {
	Validation
	// msg marshals PatternFlowIpv6PayloadLengthCounter to protobuf object *otg.PatternFlowIpv6PayloadLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6PayloadLengthCounter
	// setMsg unmarshals PatternFlowIpv6PayloadLengthCounter from protobuf object *otg.PatternFlowIpv6PayloadLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6PayloadLengthCounter) PatternFlowIpv6PayloadLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6PayloadLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6PayloadLengthCounter
	// validate validates PatternFlowIpv6PayloadLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6PayloadLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6PayloadLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6PayloadLengthCounter
	SetStart(value uint32) PatternFlowIpv6PayloadLengthCounter
	// HasStart checks if Start has been set in PatternFlowIpv6PayloadLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6PayloadLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6PayloadLengthCounter
	SetStep(value uint32) PatternFlowIpv6PayloadLengthCounter
	// HasStep checks if Step has been set in PatternFlowIpv6PayloadLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6PayloadLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6PayloadLengthCounter
	SetCount(value uint32) PatternFlowIpv6PayloadLengthCounter
	// HasCount checks if Count has been set in PatternFlowIpv6PayloadLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6PayloadLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6PayloadLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6PayloadLengthCounter object
func (obj *patternFlowIpv6PayloadLengthCounter) SetStart(value uint32) PatternFlowIpv6PayloadLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6PayloadLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6PayloadLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6PayloadLengthCounter object
func (obj *patternFlowIpv6PayloadLengthCounter) SetStep(value uint32) PatternFlowIpv6PayloadLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6PayloadLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6PayloadLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6PayloadLengthCounter object
func (obj *patternFlowIpv6PayloadLengthCounter) SetCount(value uint32) PatternFlowIpv6PayloadLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6PayloadLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6PayloadLengthCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6PayloadLengthCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6PayloadLengthCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6PayloadLengthCounter) setDefault() {
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
