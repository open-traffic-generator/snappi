package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6TrafficClassCounter *****
type patternFlowIpv6TrafficClassCounter struct {
	validation
	obj          *otg.PatternFlowIpv6TrafficClassCounter
	marshaller   marshalPatternFlowIpv6TrafficClassCounter
	unMarshaller unMarshalPatternFlowIpv6TrafficClassCounter
}

func NewPatternFlowIpv6TrafficClassCounter() PatternFlowIpv6TrafficClassCounter {
	obj := patternFlowIpv6TrafficClassCounter{obj: &otg.PatternFlowIpv6TrafficClassCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6TrafficClassCounter) msg() *otg.PatternFlowIpv6TrafficClassCounter {
	return obj.obj
}

func (obj *patternFlowIpv6TrafficClassCounter) setMsg(msg *otg.PatternFlowIpv6TrafficClassCounter) PatternFlowIpv6TrafficClassCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6TrafficClassCounter struct {
	obj *patternFlowIpv6TrafficClassCounter
}

type marshalPatternFlowIpv6TrafficClassCounter interface {
	// ToProto marshals PatternFlowIpv6TrafficClassCounter to protobuf object *otg.PatternFlowIpv6TrafficClassCounter
	ToProto() (*otg.PatternFlowIpv6TrafficClassCounter, error)
	// ToPbText marshals PatternFlowIpv6TrafficClassCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6TrafficClassCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6TrafficClassCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6TrafficClassCounter struct {
	obj *patternFlowIpv6TrafficClassCounter
}

type unMarshalPatternFlowIpv6TrafficClassCounter interface {
	// FromProto unmarshals PatternFlowIpv6TrafficClassCounter from protobuf object *otg.PatternFlowIpv6TrafficClassCounter
	FromProto(msg *otg.PatternFlowIpv6TrafficClassCounter) (PatternFlowIpv6TrafficClassCounter, error)
	// FromPbText unmarshals PatternFlowIpv6TrafficClassCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6TrafficClassCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6TrafficClassCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6TrafficClassCounter) Marshal() marshalPatternFlowIpv6TrafficClassCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6TrafficClassCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6TrafficClassCounter) Unmarshal() unMarshalPatternFlowIpv6TrafficClassCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6TrafficClassCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6TrafficClassCounter) ToProto() (*otg.PatternFlowIpv6TrafficClassCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6TrafficClassCounter) FromProto(msg *otg.PatternFlowIpv6TrafficClassCounter) (PatternFlowIpv6TrafficClassCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6TrafficClassCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6TrafficClassCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6TrafficClassCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6TrafficClassCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6TrafficClassCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6TrafficClassCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv6TrafficClassCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6TrafficClassCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6TrafficClassCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6TrafficClassCounter) Clone() (PatternFlowIpv6TrafficClassCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6TrafficClassCounter()
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

// PatternFlowIpv6TrafficClassCounter is integer counter pattern
type PatternFlowIpv6TrafficClassCounter interface {
	Validation
	// msg marshals PatternFlowIpv6TrafficClassCounter to protobuf object *otg.PatternFlowIpv6TrafficClassCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6TrafficClassCounter
	// setMsg unmarshals PatternFlowIpv6TrafficClassCounter from protobuf object *otg.PatternFlowIpv6TrafficClassCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6TrafficClassCounter) PatternFlowIpv6TrafficClassCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6TrafficClassCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6TrafficClassCounter
	// validate validates PatternFlowIpv6TrafficClassCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6TrafficClassCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv6TrafficClassCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv6TrafficClassCounter
	SetStart(value uint32) PatternFlowIpv6TrafficClassCounter
	// HasStart checks if Start has been set in PatternFlowIpv6TrafficClassCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv6TrafficClassCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv6TrafficClassCounter
	SetStep(value uint32) PatternFlowIpv6TrafficClassCounter
	// HasStep checks if Step has been set in PatternFlowIpv6TrafficClassCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv6TrafficClassCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv6TrafficClassCounter
	SetCount(value uint32) PatternFlowIpv6TrafficClassCounter
	// HasCount checks if Count has been set in PatternFlowIpv6TrafficClassCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6TrafficClassCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv6TrafficClassCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv6TrafficClassCounter object
func (obj *patternFlowIpv6TrafficClassCounter) SetStart(value uint32) PatternFlowIpv6TrafficClassCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6TrafficClassCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv6TrafficClassCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv6TrafficClassCounter object
func (obj *patternFlowIpv6TrafficClassCounter) SetStep(value uint32) PatternFlowIpv6TrafficClassCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6TrafficClassCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv6TrafficClassCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv6TrafficClassCounter object
func (obj *patternFlowIpv6TrafficClassCounter) SetCount(value uint32) PatternFlowIpv6TrafficClassCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv6TrafficClassCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6TrafficClassCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6TrafficClassCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6TrafficClassCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv6TrafficClassCounter) setDefault() {
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
