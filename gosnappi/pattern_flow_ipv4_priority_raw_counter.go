package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4PriorityRawCounter *****
type patternFlowIpv4PriorityRawCounter struct {
	validation
	obj          *otg.PatternFlowIpv4PriorityRawCounter
	marshaller   marshalPatternFlowIpv4PriorityRawCounter
	unMarshaller unMarshalPatternFlowIpv4PriorityRawCounter
}

func NewPatternFlowIpv4PriorityRawCounter() PatternFlowIpv4PriorityRawCounter {
	obj := patternFlowIpv4PriorityRawCounter{obj: &otg.PatternFlowIpv4PriorityRawCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4PriorityRawCounter) msg() *otg.PatternFlowIpv4PriorityRawCounter {
	return obj.obj
}

func (obj *patternFlowIpv4PriorityRawCounter) setMsg(msg *otg.PatternFlowIpv4PriorityRawCounter) PatternFlowIpv4PriorityRawCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4PriorityRawCounter struct {
	obj *patternFlowIpv4PriorityRawCounter
}

type marshalPatternFlowIpv4PriorityRawCounter interface {
	// ToProto marshals PatternFlowIpv4PriorityRawCounter to protobuf object *otg.PatternFlowIpv4PriorityRawCounter
	ToProto() (*otg.PatternFlowIpv4PriorityRawCounter, error)
	// ToPbText marshals PatternFlowIpv4PriorityRawCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4PriorityRawCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4PriorityRawCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4PriorityRawCounter struct {
	obj *patternFlowIpv4PriorityRawCounter
}

type unMarshalPatternFlowIpv4PriorityRawCounter interface {
	// FromProto unmarshals PatternFlowIpv4PriorityRawCounter from protobuf object *otg.PatternFlowIpv4PriorityRawCounter
	FromProto(msg *otg.PatternFlowIpv4PriorityRawCounter) (PatternFlowIpv4PriorityRawCounter, error)
	// FromPbText unmarshals PatternFlowIpv4PriorityRawCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4PriorityRawCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4PriorityRawCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4PriorityRawCounter) Marshal() marshalPatternFlowIpv4PriorityRawCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4PriorityRawCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4PriorityRawCounter) Unmarshal() unMarshalPatternFlowIpv4PriorityRawCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4PriorityRawCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4PriorityRawCounter) ToProto() (*otg.PatternFlowIpv4PriorityRawCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4PriorityRawCounter) FromProto(msg *otg.PatternFlowIpv4PriorityRawCounter) (PatternFlowIpv4PriorityRawCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4PriorityRawCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4PriorityRawCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4PriorityRawCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4PriorityRawCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4PriorityRawCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4PriorityRawCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4PriorityRawCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4PriorityRawCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4PriorityRawCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4PriorityRawCounter) Clone() (PatternFlowIpv4PriorityRawCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4PriorityRawCounter()
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

// PatternFlowIpv4PriorityRawCounter is integer counter pattern
type PatternFlowIpv4PriorityRawCounter interface {
	Validation
	// msg marshals PatternFlowIpv4PriorityRawCounter to protobuf object *otg.PatternFlowIpv4PriorityRawCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4PriorityRawCounter
	// setMsg unmarshals PatternFlowIpv4PriorityRawCounter from protobuf object *otg.PatternFlowIpv4PriorityRawCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4PriorityRawCounter) PatternFlowIpv4PriorityRawCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4PriorityRawCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4PriorityRawCounter
	// validate validates PatternFlowIpv4PriorityRawCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4PriorityRawCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4PriorityRawCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4PriorityRawCounter
	SetStart(value uint32) PatternFlowIpv4PriorityRawCounter
	// HasStart checks if Start has been set in PatternFlowIpv4PriorityRawCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4PriorityRawCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4PriorityRawCounter
	SetStep(value uint32) PatternFlowIpv4PriorityRawCounter
	// HasStep checks if Step has been set in PatternFlowIpv4PriorityRawCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4PriorityRawCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4PriorityRawCounter
	SetCount(value uint32) PatternFlowIpv4PriorityRawCounter
	// HasCount checks if Count has been set in PatternFlowIpv4PriorityRawCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4PriorityRawCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4PriorityRawCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4PriorityRawCounter object
func (obj *patternFlowIpv4PriorityRawCounter) SetStart(value uint32) PatternFlowIpv4PriorityRawCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4PriorityRawCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4PriorityRawCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4PriorityRawCounter object
func (obj *patternFlowIpv4PriorityRawCounter) SetStep(value uint32) PatternFlowIpv4PriorityRawCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4PriorityRawCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4PriorityRawCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4PriorityRawCounter object
func (obj *patternFlowIpv4PriorityRawCounter) SetCount(value uint32) PatternFlowIpv4PriorityRawCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4PriorityRawCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4PriorityRawCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4PriorityRawCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4PriorityRawCounter.Count <= 256 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4PriorityRawCounter) setDefault() {
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
