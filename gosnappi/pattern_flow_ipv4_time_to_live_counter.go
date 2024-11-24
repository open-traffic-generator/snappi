package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TimeToLiveCounter *****
type patternFlowIpv4TimeToLiveCounter struct {
	validation
	obj          *otg.PatternFlowIpv4TimeToLiveCounter
	marshaller   marshalPatternFlowIpv4TimeToLiveCounter
	unMarshaller unMarshalPatternFlowIpv4TimeToLiveCounter
}

func NewPatternFlowIpv4TimeToLiveCounter() PatternFlowIpv4TimeToLiveCounter {
	obj := patternFlowIpv4TimeToLiveCounter{obj: &otg.PatternFlowIpv4TimeToLiveCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TimeToLiveCounter) msg() *otg.PatternFlowIpv4TimeToLiveCounter {
	return obj.obj
}

func (obj *patternFlowIpv4TimeToLiveCounter) setMsg(msg *otg.PatternFlowIpv4TimeToLiveCounter) PatternFlowIpv4TimeToLiveCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TimeToLiveCounter struct {
	obj *patternFlowIpv4TimeToLiveCounter
}

type marshalPatternFlowIpv4TimeToLiveCounter interface {
	// ToProto marshals PatternFlowIpv4TimeToLiveCounter to protobuf object *otg.PatternFlowIpv4TimeToLiveCounter
	ToProto() (*otg.PatternFlowIpv4TimeToLiveCounter, error)
	// ToPbText marshals PatternFlowIpv4TimeToLiveCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TimeToLiveCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TimeToLiveCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TimeToLiveCounter struct {
	obj *patternFlowIpv4TimeToLiveCounter
}

type unMarshalPatternFlowIpv4TimeToLiveCounter interface {
	// FromProto unmarshals PatternFlowIpv4TimeToLiveCounter from protobuf object *otg.PatternFlowIpv4TimeToLiveCounter
	FromProto(msg *otg.PatternFlowIpv4TimeToLiveCounter) (PatternFlowIpv4TimeToLiveCounter, error)
	// FromPbText unmarshals PatternFlowIpv4TimeToLiveCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TimeToLiveCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TimeToLiveCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TimeToLiveCounter) Marshal() marshalPatternFlowIpv4TimeToLiveCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TimeToLiveCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TimeToLiveCounter) Unmarshal() unMarshalPatternFlowIpv4TimeToLiveCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TimeToLiveCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TimeToLiveCounter) ToProto() (*otg.PatternFlowIpv4TimeToLiveCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TimeToLiveCounter) FromProto(msg *otg.PatternFlowIpv4TimeToLiveCounter) (PatternFlowIpv4TimeToLiveCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TimeToLiveCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TimeToLiveCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TimeToLiveCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TimeToLiveCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TimeToLiveCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TimeToLiveCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4TimeToLiveCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TimeToLiveCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TimeToLiveCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TimeToLiveCounter) Clone() (PatternFlowIpv4TimeToLiveCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TimeToLiveCounter()
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

// PatternFlowIpv4TimeToLiveCounter is integer counter pattern
type PatternFlowIpv4TimeToLiveCounter interface {
	Validation
	// msg marshals PatternFlowIpv4TimeToLiveCounter to protobuf object *otg.PatternFlowIpv4TimeToLiveCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TimeToLiveCounter
	// setMsg unmarshals PatternFlowIpv4TimeToLiveCounter from protobuf object *otg.PatternFlowIpv4TimeToLiveCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TimeToLiveCounter) PatternFlowIpv4TimeToLiveCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TimeToLiveCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TimeToLiveCounter
	// validate validates PatternFlowIpv4TimeToLiveCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TimeToLiveCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4TimeToLiveCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4TimeToLiveCounter
	SetStart(value uint32) PatternFlowIpv4TimeToLiveCounter
	// HasStart checks if Start has been set in PatternFlowIpv4TimeToLiveCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4TimeToLiveCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4TimeToLiveCounter
	SetStep(value uint32) PatternFlowIpv4TimeToLiveCounter
	// HasStep checks if Step has been set in PatternFlowIpv4TimeToLiveCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4TimeToLiveCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4TimeToLiveCounter
	SetCount(value uint32) PatternFlowIpv4TimeToLiveCounter
	// HasCount checks if Count has been set in PatternFlowIpv4TimeToLiveCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TimeToLiveCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4TimeToLiveCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4TimeToLiveCounter object
func (obj *patternFlowIpv4TimeToLiveCounter) SetStart(value uint32) PatternFlowIpv4TimeToLiveCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TimeToLiveCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4TimeToLiveCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4TimeToLiveCounter object
func (obj *patternFlowIpv4TimeToLiveCounter) SetStep(value uint32) PatternFlowIpv4TimeToLiveCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TimeToLiveCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4TimeToLiveCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4TimeToLiveCounter object
func (obj *patternFlowIpv4TimeToLiveCounter) SetCount(value uint32) PatternFlowIpv4TimeToLiveCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4TimeToLiveCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TimeToLiveCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TimeToLiveCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TimeToLiveCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4TimeToLiveCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(64)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
