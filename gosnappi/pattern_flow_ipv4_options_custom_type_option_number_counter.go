package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsCustomTypeOptionNumberCounter *****
type patternFlowIpv4OptionsCustomTypeOptionNumberCounter struct {
	validation
	obj          *otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	marshaller   marshalPatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	unMarshaller unMarshalPatternFlowIpv4OptionsCustomTypeOptionNumberCounter
}

func NewPatternFlowIpv4OptionsCustomTypeOptionNumberCounter() PatternFlowIpv4OptionsCustomTypeOptionNumberCounter {
	obj := patternFlowIpv4OptionsCustomTypeOptionNumberCounter{obj: &otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) msg() *otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) setMsg(msg *otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter) PatternFlowIpv4OptionsCustomTypeOptionNumberCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsCustomTypeOptionNumberCounter struct {
	obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter
}

type marshalPatternFlowIpv4OptionsCustomTypeOptionNumberCounter interface {
	// ToProto marshals PatternFlowIpv4OptionsCustomTypeOptionNumberCounter to protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	ToProto() (*otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter, error)
	// ToPbText marshals PatternFlowIpv4OptionsCustomTypeOptionNumberCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsCustomTypeOptionNumberCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsCustomTypeOptionNumberCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsCustomTypeOptionNumberCounter struct {
	obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter
}

type unMarshalPatternFlowIpv4OptionsCustomTypeOptionNumberCounter interface {
	// FromProto unmarshals PatternFlowIpv4OptionsCustomTypeOptionNumberCounter from protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	FromProto(msg *otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter) (PatternFlowIpv4OptionsCustomTypeOptionNumberCounter, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsCustomTypeOptionNumberCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsCustomTypeOptionNumberCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsCustomTypeOptionNumberCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) Marshal() marshalPatternFlowIpv4OptionsCustomTypeOptionNumberCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsCustomTypeOptionNumberCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) Unmarshal() unMarshalPatternFlowIpv4OptionsCustomTypeOptionNumberCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsCustomTypeOptionNumberCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionNumberCounter) ToProto() (*otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionNumberCounter) FromProto(msg *otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter) (PatternFlowIpv4OptionsCustomTypeOptionNumberCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionNumberCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionNumberCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionNumberCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionNumberCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionNumberCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionNumberCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) Clone() (PatternFlowIpv4OptionsCustomTypeOptionNumberCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsCustomTypeOptionNumberCounter()
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

// PatternFlowIpv4OptionsCustomTypeOptionNumberCounter is integer counter pattern
type PatternFlowIpv4OptionsCustomTypeOptionNumberCounter interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsCustomTypeOptionNumberCounter to protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	// setMsg unmarshals PatternFlowIpv4OptionsCustomTypeOptionNumberCounter from protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsCustomTypeOptionNumberCounter) PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	// validate validates PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsCustomTypeOptionNumberCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4OptionsCustomTypeOptionNumberCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	SetStart(value uint32) PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	// HasStart checks if Start has been set in PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4OptionsCustomTypeOptionNumberCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	SetStep(value uint32) PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	// HasStep checks if Step has been set in PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4OptionsCustomTypeOptionNumberCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	SetCount(value uint32) PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	// HasCount checks if Count has been set in PatternFlowIpv4OptionsCustomTypeOptionNumberCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4OptionsCustomTypeOptionNumberCounter object
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) SetStart(value uint32) PatternFlowIpv4OptionsCustomTypeOptionNumberCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4OptionsCustomTypeOptionNumberCounter object
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) SetStep(value uint32) PatternFlowIpv4OptionsCustomTypeOptionNumberCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4OptionsCustomTypeOptionNumberCounter object
func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) SetCount(value uint32) PatternFlowIpv4OptionsCustomTypeOptionNumberCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsCustomTypeOptionNumberCounter.Start <= 31 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsCustomTypeOptionNumberCounter.Step <= 31 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsCustomTypeOptionNumberCounter.Count <= 32 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4OptionsCustomTypeOptionNumberCounter) setDefault() {
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
