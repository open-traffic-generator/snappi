package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4IdentificationCounter *****
type patternFlowIpv4IdentificationCounter struct {
	validation
	obj          *otg.PatternFlowIpv4IdentificationCounter
	marshaller   marshalPatternFlowIpv4IdentificationCounter
	unMarshaller unMarshalPatternFlowIpv4IdentificationCounter
}

func NewPatternFlowIpv4IdentificationCounter() PatternFlowIpv4IdentificationCounter {
	obj := patternFlowIpv4IdentificationCounter{obj: &otg.PatternFlowIpv4IdentificationCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4IdentificationCounter) msg() *otg.PatternFlowIpv4IdentificationCounter {
	return obj.obj
}

func (obj *patternFlowIpv4IdentificationCounter) setMsg(msg *otg.PatternFlowIpv4IdentificationCounter) PatternFlowIpv4IdentificationCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4IdentificationCounter struct {
	obj *patternFlowIpv4IdentificationCounter
}

type marshalPatternFlowIpv4IdentificationCounter interface {
	// ToProto marshals PatternFlowIpv4IdentificationCounter to protobuf object *otg.PatternFlowIpv4IdentificationCounter
	ToProto() (*otg.PatternFlowIpv4IdentificationCounter, error)
	// ToPbText marshals PatternFlowIpv4IdentificationCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4IdentificationCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4IdentificationCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4IdentificationCounter struct {
	obj *patternFlowIpv4IdentificationCounter
}

type unMarshalPatternFlowIpv4IdentificationCounter interface {
	// FromProto unmarshals PatternFlowIpv4IdentificationCounter from protobuf object *otg.PatternFlowIpv4IdentificationCounter
	FromProto(msg *otg.PatternFlowIpv4IdentificationCounter) (PatternFlowIpv4IdentificationCounter, error)
	// FromPbText unmarshals PatternFlowIpv4IdentificationCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4IdentificationCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4IdentificationCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4IdentificationCounter) Marshal() marshalPatternFlowIpv4IdentificationCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4IdentificationCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4IdentificationCounter) Unmarshal() unMarshalPatternFlowIpv4IdentificationCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4IdentificationCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4IdentificationCounter) ToProto() (*otg.PatternFlowIpv4IdentificationCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4IdentificationCounter) FromProto(msg *otg.PatternFlowIpv4IdentificationCounter) (PatternFlowIpv4IdentificationCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4IdentificationCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4IdentificationCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4IdentificationCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4IdentificationCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4IdentificationCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4IdentificationCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4IdentificationCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4IdentificationCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4IdentificationCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4IdentificationCounter) Clone() (PatternFlowIpv4IdentificationCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4IdentificationCounter()
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

// PatternFlowIpv4IdentificationCounter is integer counter pattern
type PatternFlowIpv4IdentificationCounter interface {
	Validation
	// msg marshals PatternFlowIpv4IdentificationCounter to protobuf object *otg.PatternFlowIpv4IdentificationCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4IdentificationCounter
	// setMsg unmarshals PatternFlowIpv4IdentificationCounter from protobuf object *otg.PatternFlowIpv4IdentificationCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4IdentificationCounter) PatternFlowIpv4IdentificationCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4IdentificationCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4IdentificationCounter
	// validate validates PatternFlowIpv4IdentificationCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4IdentificationCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4IdentificationCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4IdentificationCounter
	SetStart(value uint32) PatternFlowIpv4IdentificationCounter
	// HasStart checks if Start has been set in PatternFlowIpv4IdentificationCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4IdentificationCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4IdentificationCounter
	SetStep(value uint32) PatternFlowIpv4IdentificationCounter
	// HasStep checks if Step has been set in PatternFlowIpv4IdentificationCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4IdentificationCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4IdentificationCounter
	SetCount(value uint32) PatternFlowIpv4IdentificationCounter
	// HasCount checks if Count has been set in PatternFlowIpv4IdentificationCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4IdentificationCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4IdentificationCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4IdentificationCounter object
func (obj *patternFlowIpv4IdentificationCounter) SetStart(value uint32) PatternFlowIpv4IdentificationCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4IdentificationCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4IdentificationCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4IdentificationCounter object
func (obj *patternFlowIpv4IdentificationCounter) SetStep(value uint32) PatternFlowIpv4IdentificationCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4IdentificationCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4IdentificationCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4IdentificationCounter object
func (obj *patternFlowIpv4IdentificationCounter) SetCount(value uint32) PatternFlowIpv4IdentificationCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4IdentificationCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4IdentificationCounter.Start <= 65535 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4IdentificationCounter.Step <= 65535 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4IdentificationCounter.Count <= 65535 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4IdentificationCounter) setDefault() {
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
