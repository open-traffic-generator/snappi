package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4HeaderLengthCounter *****
type patternFlowIpv4HeaderLengthCounter struct {
	validation
	obj          *otg.PatternFlowIpv4HeaderLengthCounter
	marshaller   marshalPatternFlowIpv4HeaderLengthCounter
	unMarshaller unMarshalPatternFlowIpv4HeaderLengthCounter
}

func NewPatternFlowIpv4HeaderLengthCounter() PatternFlowIpv4HeaderLengthCounter {
	obj := patternFlowIpv4HeaderLengthCounter{obj: &otg.PatternFlowIpv4HeaderLengthCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4HeaderLengthCounter) msg() *otg.PatternFlowIpv4HeaderLengthCounter {
	return obj.obj
}

func (obj *patternFlowIpv4HeaderLengthCounter) setMsg(msg *otg.PatternFlowIpv4HeaderLengthCounter) PatternFlowIpv4HeaderLengthCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4HeaderLengthCounter struct {
	obj *patternFlowIpv4HeaderLengthCounter
}

type marshalPatternFlowIpv4HeaderLengthCounter interface {
	// ToProto marshals PatternFlowIpv4HeaderLengthCounter to protobuf object *otg.PatternFlowIpv4HeaderLengthCounter
	ToProto() (*otg.PatternFlowIpv4HeaderLengthCounter, error)
	// ToPbText marshals PatternFlowIpv4HeaderLengthCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4HeaderLengthCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4HeaderLengthCounter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4HeaderLengthCounter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4HeaderLengthCounter struct {
	obj *patternFlowIpv4HeaderLengthCounter
}

type unMarshalPatternFlowIpv4HeaderLengthCounter interface {
	// FromProto unmarshals PatternFlowIpv4HeaderLengthCounter from protobuf object *otg.PatternFlowIpv4HeaderLengthCounter
	FromProto(msg *otg.PatternFlowIpv4HeaderLengthCounter) (PatternFlowIpv4HeaderLengthCounter, error)
	// FromPbText unmarshals PatternFlowIpv4HeaderLengthCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4HeaderLengthCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4HeaderLengthCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4HeaderLengthCounter) Marshal() marshalPatternFlowIpv4HeaderLengthCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4HeaderLengthCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4HeaderLengthCounter) Unmarshal() unMarshalPatternFlowIpv4HeaderLengthCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4HeaderLengthCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4HeaderLengthCounter) ToProto() (*otg.PatternFlowIpv4HeaderLengthCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4HeaderLengthCounter) FromProto(msg *otg.PatternFlowIpv4HeaderLengthCounter) (PatternFlowIpv4HeaderLengthCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4HeaderLengthCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4HeaderLengthCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4HeaderLengthCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4HeaderLengthCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4HeaderLengthCounter) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4HeaderLengthCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4HeaderLengthCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4HeaderLengthCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4HeaderLengthCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4HeaderLengthCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4HeaderLengthCounter) Clone() (PatternFlowIpv4HeaderLengthCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4HeaderLengthCounter()
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

// PatternFlowIpv4HeaderLengthCounter is integer counter pattern
type PatternFlowIpv4HeaderLengthCounter interface {
	Validation
	// msg marshals PatternFlowIpv4HeaderLengthCounter to protobuf object *otg.PatternFlowIpv4HeaderLengthCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4HeaderLengthCounter
	// setMsg unmarshals PatternFlowIpv4HeaderLengthCounter from protobuf object *otg.PatternFlowIpv4HeaderLengthCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4HeaderLengthCounter) PatternFlowIpv4HeaderLengthCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4HeaderLengthCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4HeaderLengthCounter
	// validate validates PatternFlowIpv4HeaderLengthCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4HeaderLengthCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4HeaderLengthCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4HeaderLengthCounter
	SetStart(value uint32) PatternFlowIpv4HeaderLengthCounter
	// HasStart checks if Start has been set in PatternFlowIpv4HeaderLengthCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4HeaderLengthCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4HeaderLengthCounter
	SetStep(value uint32) PatternFlowIpv4HeaderLengthCounter
	// HasStep checks if Step has been set in PatternFlowIpv4HeaderLengthCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4HeaderLengthCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4HeaderLengthCounter
	SetCount(value uint32) PatternFlowIpv4HeaderLengthCounter
	// HasCount checks if Count has been set in PatternFlowIpv4HeaderLengthCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4HeaderLengthCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4HeaderLengthCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4HeaderLengthCounter object
func (obj *patternFlowIpv4HeaderLengthCounter) SetStart(value uint32) PatternFlowIpv4HeaderLengthCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4HeaderLengthCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4HeaderLengthCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4HeaderLengthCounter object
func (obj *patternFlowIpv4HeaderLengthCounter) SetStep(value uint32) PatternFlowIpv4HeaderLengthCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4HeaderLengthCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4HeaderLengthCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4HeaderLengthCounter object
func (obj *patternFlowIpv4HeaderLengthCounter) SetCount(value uint32) PatternFlowIpv4HeaderLengthCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4HeaderLengthCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4HeaderLengthCounter.Start <= 15 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4HeaderLengthCounter.Step <= 15 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4HeaderLengthCounter.Count <= 15 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4HeaderLengthCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(5)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
