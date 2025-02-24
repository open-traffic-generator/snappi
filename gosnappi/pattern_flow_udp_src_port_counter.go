package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowUdpSrcPortCounter *****
type patternFlowUdpSrcPortCounter struct {
	validation
	obj          *otg.PatternFlowUdpSrcPortCounter
	marshaller   marshalPatternFlowUdpSrcPortCounter
	unMarshaller unMarshalPatternFlowUdpSrcPortCounter
}

func NewPatternFlowUdpSrcPortCounter() PatternFlowUdpSrcPortCounter {
	obj := patternFlowUdpSrcPortCounter{obj: &otg.PatternFlowUdpSrcPortCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowUdpSrcPortCounter) msg() *otg.PatternFlowUdpSrcPortCounter {
	return obj.obj
}

func (obj *patternFlowUdpSrcPortCounter) setMsg(msg *otg.PatternFlowUdpSrcPortCounter) PatternFlowUdpSrcPortCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowUdpSrcPortCounter struct {
	obj *patternFlowUdpSrcPortCounter
}

type marshalPatternFlowUdpSrcPortCounter interface {
	// ToProto marshals PatternFlowUdpSrcPortCounter to protobuf object *otg.PatternFlowUdpSrcPortCounter
	ToProto() (*otg.PatternFlowUdpSrcPortCounter, error)
	// ToPbText marshals PatternFlowUdpSrcPortCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowUdpSrcPortCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowUdpSrcPortCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowUdpSrcPortCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowUdpSrcPortCounter struct {
	obj *patternFlowUdpSrcPortCounter
}

type unMarshalPatternFlowUdpSrcPortCounter interface {
	// FromProto unmarshals PatternFlowUdpSrcPortCounter from protobuf object *otg.PatternFlowUdpSrcPortCounter
	FromProto(msg *otg.PatternFlowUdpSrcPortCounter) (PatternFlowUdpSrcPortCounter, error)
	// FromPbText unmarshals PatternFlowUdpSrcPortCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowUdpSrcPortCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowUdpSrcPortCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowUdpSrcPortCounter) Marshal() marshalPatternFlowUdpSrcPortCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowUdpSrcPortCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowUdpSrcPortCounter) Unmarshal() unMarshalPatternFlowUdpSrcPortCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowUdpSrcPortCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowUdpSrcPortCounter) ToProto() (*otg.PatternFlowUdpSrcPortCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowUdpSrcPortCounter) FromProto(msg *otg.PatternFlowUdpSrcPortCounter) (PatternFlowUdpSrcPortCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowUdpSrcPortCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowUdpSrcPortCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowUdpSrcPortCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowUdpSrcPortCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowUdpSrcPortCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowUdpSrcPortCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowUdpSrcPortCounter) FromJson(value string) error {
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

func (obj *patternFlowUdpSrcPortCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowUdpSrcPortCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowUdpSrcPortCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowUdpSrcPortCounter) Clone() (PatternFlowUdpSrcPortCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowUdpSrcPortCounter()
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

// PatternFlowUdpSrcPortCounter is integer counter pattern
type PatternFlowUdpSrcPortCounter interface {
	Validation
	// msg marshals PatternFlowUdpSrcPortCounter to protobuf object *otg.PatternFlowUdpSrcPortCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowUdpSrcPortCounter
	// setMsg unmarshals PatternFlowUdpSrcPortCounter from protobuf object *otg.PatternFlowUdpSrcPortCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowUdpSrcPortCounter) PatternFlowUdpSrcPortCounter
	// provides marshal interface
	Marshal() marshalPatternFlowUdpSrcPortCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowUdpSrcPortCounter
	// validate validates PatternFlowUdpSrcPortCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowUdpSrcPortCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowUdpSrcPortCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowUdpSrcPortCounter
	SetStart(value uint32) PatternFlowUdpSrcPortCounter
	// HasStart checks if Start has been set in PatternFlowUdpSrcPortCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowUdpSrcPortCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowUdpSrcPortCounter
	SetStep(value uint32) PatternFlowUdpSrcPortCounter
	// HasStep checks if Step has been set in PatternFlowUdpSrcPortCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowUdpSrcPortCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowUdpSrcPortCounter
	SetCount(value uint32) PatternFlowUdpSrcPortCounter
	// HasCount checks if Count has been set in PatternFlowUdpSrcPortCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowUdpSrcPortCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowUdpSrcPortCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowUdpSrcPortCounter object
func (obj *patternFlowUdpSrcPortCounter) SetStart(value uint32) PatternFlowUdpSrcPortCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowUdpSrcPortCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowUdpSrcPortCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowUdpSrcPortCounter object
func (obj *patternFlowUdpSrcPortCounter) SetStep(value uint32) PatternFlowUdpSrcPortCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowUdpSrcPortCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowUdpSrcPortCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowUdpSrcPortCounter object
func (obj *patternFlowUdpSrcPortCounter) SetCount(value uint32) PatternFlowUdpSrcPortCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowUdpSrcPortCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpSrcPortCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpSrcPortCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpSrcPortCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowUdpSrcPortCounter) setDefault() {
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
