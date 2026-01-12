package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4DontFragmentCounter *****
type patternFlowIpv4DontFragmentCounter struct {
	validation
	obj          *otg.PatternFlowIpv4DontFragmentCounter
	marshaller   marshalPatternFlowIpv4DontFragmentCounter
	unMarshaller unMarshalPatternFlowIpv4DontFragmentCounter
}

func NewPatternFlowIpv4DontFragmentCounter() PatternFlowIpv4DontFragmentCounter {
	obj := patternFlowIpv4DontFragmentCounter{obj: &otg.PatternFlowIpv4DontFragmentCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4DontFragmentCounter) msg() *otg.PatternFlowIpv4DontFragmentCounter {
	return obj.obj
}

func (obj *patternFlowIpv4DontFragmentCounter) setMsg(msg *otg.PatternFlowIpv4DontFragmentCounter) PatternFlowIpv4DontFragmentCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4DontFragmentCounter struct {
	obj *patternFlowIpv4DontFragmentCounter
}

type marshalPatternFlowIpv4DontFragmentCounter interface {
	// ToProto marshals PatternFlowIpv4DontFragmentCounter to protobuf object *otg.PatternFlowIpv4DontFragmentCounter
	ToProto() (*otg.PatternFlowIpv4DontFragmentCounter, error)
	// ToPbText marshals PatternFlowIpv4DontFragmentCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4DontFragmentCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4DontFragmentCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4DontFragmentCounter struct {
	obj *patternFlowIpv4DontFragmentCounter
}

type unMarshalPatternFlowIpv4DontFragmentCounter interface {
	// FromProto unmarshals PatternFlowIpv4DontFragmentCounter from protobuf object *otg.PatternFlowIpv4DontFragmentCounter
	FromProto(msg *otg.PatternFlowIpv4DontFragmentCounter) (PatternFlowIpv4DontFragmentCounter, error)
	// FromPbText unmarshals PatternFlowIpv4DontFragmentCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4DontFragmentCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4DontFragmentCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4DontFragmentCounter) Marshal() marshalPatternFlowIpv4DontFragmentCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4DontFragmentCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4DontFragmentCounter) Unmarshal() unMarshalPatternFlowIpv4DontFragmentCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4DontFragmentCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4DontFragmentCounter) ToProto() (*otg.PatternFlowIpv4DontFragmentCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4DontFragmentCounter) FromProto(msg *otg.PatternFlowIpv4DontFragmentCounter) (PatternFlowIpv4DontFragmentCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4DontFragmentCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DontFragmentCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4DontFragmentCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DontFragmentCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4DontFragmentCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DontFragmentCounter) FromJson(value string) error {
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

func (obj *patternFlowIpv4DontFragmentCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DontFragmentCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DontFragmentCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4DontFragmentCounter) Clone() (PatternFlowIpv4DontFragmentCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4DontFragmentCounter()
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

// PatternFlowIpv4DontFragmentCounter is integer counter pattern
type PatternFlowIpv4DontFragmentCounter interface {
	Validation
	// msg marshals PatternFlowIpv4DontFragmentCounter to protobuf object *otg.PatternFlowIpv4DontFragmentCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4DontFragmentCounter
	// setMsg unmarshals PatternFlowIpv4DontFragmentCounter from protobuf object *otg.PatternFlowIpv4DontFragmentCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4DontFragmentCounter) PatternFlowIpv4DontFragmentCounter
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4DontFragmentCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4DontFragmentCounter
	// validate validates PatternFlowIpv4DontFragmentCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4DontFragmentCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowIpv4DontFragmentCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowIpv4DontFragmentCounter
	SetStart(value uint32) PatternFlowIpv4DontFragmentCounter
	// HasStart checks if Start has been set in PatternFlowIpv4DontFragmentCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowIpv4DontFragmentCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowIpv4DontFragmentCounter
	SetStep(value uint32) PatternFlowIpv4DontFragmentCounter
	// HasStep checks if Step has been set in PatternFlowIpv4DontFragmentCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowIpv4DontFragmentCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowIpv4DontFragmentCounter
	SetCount(value uint32) PatternFlowIpv4DontFragmentCounter
	// HasCount checks if Count has been set in PatternFlowIpv4DontFragmentCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4DontFragmentCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowIpv4DontFragmentCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowIpv4DontFragmentCounter object
func (obj *patternFlowIpv4DontFragmentCounter) SetStart(value uint32) PatternFlowIpv4DontFragmentCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4DontFragmentCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowIpv4DontFragmentCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowIpv4DontFragmentCounter object
func (obj *patternFlowIpv4DontFragmentCounter) SetStep(value uint32) PatternFlowIpv4DontFragmentCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4DontFragmentCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowIpv4DontFragmentCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowIpv4DontFragmentCounter object
func (obj *patternFlowIpv4DontFragmentCounter) SetCount(value uint32) PatternFlowIpv4DontFragmentCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowIpv4DontFragmentCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DontFragmentCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DontFragmentCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DontFragmentCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowIpv4DontFragmentCounter) setDefault() {
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
