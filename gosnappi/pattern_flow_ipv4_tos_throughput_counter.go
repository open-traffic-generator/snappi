package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosThroughputCounter *****
type patternFlowIpv4TosThroughputCounter struct {
	validation
	obj          *otg.PatternFlowIpv4TosThroughputCounter
	marshaller   marshalPatternFlowIpv4TosThroughputCounter
	unMarshaller unMarshalPatternFlowIpv4TosThroughputCounter
}

func NewPatternFlowIpv4TosThroughputCounter() PatternFlowIpv4TosThroughputCounter {
	obj := patternFlowIpv4TosThroughputCounter{obj: &otg.PatternFlowIpv4TosThroughputCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosThroughputCounter) msg() *otg.PatternFlowIpv4TosThroughputCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TosThroughputCounter) setMsg(msg *otg.PatternFlowIpv4TosThroughputCounter) PatternFlowIpv4TosThroughputCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosThroughputCounter struct {
	obj *patternFlowIpv4TosThroughputCounter
}

type marshalPatternFlowIpv4TosThroughputCounter interface {
	// ToProto marshals PatternFlowIpv4TosThroughputCounter to protobuf object *otg.PatternFlowIpv4TosThroughputCounter
	ToProto() (*otg.PatternFlowIpv4TosThroughputCounter, error)
	// ToPbText marshals PatternFlowIpv4TosThroughputCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosThroughputCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosThroughputCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosThroughputCounter struct {
	obj *patternFlowIpv4TosThroughputCounter
}

type unMarshalPatternFlowIpv4TosThroughputCounter interface {
	// FromProto unmarshals PatternFlowIpv4TosThroughputCounter from protobuf object *otg.PatternFlowIpv4TosThroughputCounter
	FromProto(msg *otg.PatternFlowIpv4TosThroughputCounter) (PatternFlowIpv4TosThroughputCounter, error)
	// FromPbText unmarshals PatternFlowIpv4TosThroughputCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosThroughputCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosThroughputCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosThroughputCounter) Marshal() marshalPatternFlowIpv4TosThroughputCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosThroughputCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosThroughputCounter) Unmarshal() unMarshalPatternFlowIpv4TosThroughputCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosThroughputCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosThroughputCounter) ToProto() (*otg.PatternFlowIpv4TosThroughputCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosThroughputCounter) FromProto(msg *otg.PatternFlowIpv4TosThroughputCounter) (PatternFlowIpv4TosThroughputCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosThroughputCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosThroughputCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosThroughputCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosThroughputCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosThroughputCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosThroughputCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosThroughputCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosThroughputCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosThroughputCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosThroughputCounter) Clone() (PatternFlowIpv4TosThroughputCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosThroughputCounter()
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

// PatternFlowIpv4TosThroughputCounter is integer counter pattern
type PatternFlowIpv4TosThroughputCounter interface {
	Validation
	// msg marshals PatternFlowIpv4TosThroughputCounter to protobuf object *otg.PatternFlowIpv4TosThroughputCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosThroughputCounter
	// setMsg unmarshals PatternFlowIpv4TosThroughputCounter from protobuf object *otg.PatternFlowIpv4TosThroughputCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosThroughputCounter) PatternFlowIpv4TosThroughputCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosThroughputCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosThroughputCounter
	// validate validates PatternFlowIpv4TosThroughputCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosThroughputCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4TosThroughputCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4TosThroughputCounter
	SetStart(value uint32) PatternFlowIpv4TosThroughputCounter
	// HasStart checks if Start has been set in PatternFlowIpv4TosThroughputCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4TosThroughputCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4TosThroughputCounter
	SetStep(value uint32) PatternFlowIpv4TosThroughputCounter
	// HasStep checks if Step has been set in PatternFlowIpv4TosThroughputCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4TosThroughputCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4TosThroughputCounter
	SetCount(value uint32) PatternFlowIpv4TosThroughputCounter
	// HasCount checks if Count has been set in PatternFlowIpv4TosThroughputCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TosThroughputCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TosThroughputCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4TosThroughputCounter object
func (obj *patternFlowIpv4TosThroughputCounter) SetStart(value uint32) PatternFlowIpv4TosThroughputCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TosThroughputCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TosThroughputCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4TosThroughputCounter object
func (obj *patternFlowIpv4TosThroughputCounter) SetStep(value uint32) PatternFlowIpv4TosThroughputCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TosThroughputCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TosThroughputCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4TosThroughputCounter object
func (obj *patternFlowIpv4TosThroughputCounter) SetCount(value uint32) PatternFlowIpv4TosThroughputCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4TosThroughputCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosThroughputCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosThroughputCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosThroughputCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4TosThroughputCounter) setDefault() {
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
