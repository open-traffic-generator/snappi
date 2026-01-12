package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpEcnEchoCounter *****
type patternFlowTcpEcnEchoCounter struct {
	validation
	obj          *otg.PatternFlowTcpEcnEchoCounter
	marshaller   marshalPatternFlowTcpEcnEchoCounter
	unMarshaller unMarshalPatternFlowTcpEcnEchoCounter
}

func NewPatternFlowTcpEcnEchoCounter() PatternFlowTcpEcnEchoCounter {
	obj := patternFlowTcpEcnEchoCounter{obj: &otg.PatternFlowTcpEcnEchoCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpEcnEchoCounter) msg() *otg.PatternFlowTcpEcnEchoCounter {
	return obj.obj
}

func (obj *patternFlowTcpEcnEchoCounter) setMsg(msg *otg.PatternFlowTcpEcnEchoCounter) PatternFlowTcpEcnEchoCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpEcnEchoCounter struct {
	obj *patternFlowTcpEcnEchoCounter
}

type marshalPatternFlowTcpEcnEchoCounter interface {
	// ToProto marshals PatternFlowTcpEcnEchoCounter to protobuf object *otg.PatternFlowTcpEcnEchoCounter
	ToProto() (*otg.PatternFlowTcpEcnEchoCounter, error)
	// ToPbText marshals PatternFlowTcpEcnEchoCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpEcnEchoCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpEcnEchoCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpEcnEchoCounter struct {
	obj *patternFlowTcpEcnEchoCounter
}

type unMarshalPatternFlowTcpEcnEchoCounter interface {
	// FromProto unmarshals PatternFlowTcpEcnEchoCounter from protobuf object *otg.PatternFlowTcpEcnEchoCounter
	FromProto(msg *otg.PatternFlowTcpEcnEchoCounter) (PatternFlowTcpEcnEchoCounter, error)
	// FromPbText unmarshals PatternFlowTcpEcnEchoCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpEcnEchoCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpEcnEchoCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpEcnEchoCounter) Marshal() marshalPatternFlowTcpEcnEchoCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpEcnEchoCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpEcnEchoCounter) Unmarshal() unMarshalPatternFlowTcpEcnEchoCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpEcnEchoCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpEcnEchoCounter) ToProto() (*otg.PatternFlowTcpEcnEchoCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpEcnEchoCounter) FromProto(msg *otg.PatternFlowTcpEcnEchoCounter) (PatternFlowTcpEcnEchoCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpEcnEchoCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnEchoCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpEcnEchoCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnEchoCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpEcnEchoCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnEchoCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpEcnEchoCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnEchoCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnEchoCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpEcnEchoCounter) Clone() (PatternFlowTcpEcnEchoCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpEcnEchoCounter()
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

// PatternFlowTcpEcnEchoCounter is integer counter pattern
type PatternFlowTcpEcnEchoCounter interface {
	Validation
	// msg marshals PatternFlowTcpEcnEchoCounter to protobuf object *otg.PatternFlowTcpEcnEchoCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpEcnEchoCounter
	// setMsg unmarshals PatternFlowTcpEcnEchoCounter from protobuf object *otg.PatternFlowTcpEcnEchoCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpEcnEchoCounter) PatternFlowTcpEcnEchoCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpEcnEchoCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpEcnEchoCounter
	// validate validates PatternFlowTcpEcnEchoCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpEcnEchoCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpEcnEchoCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpEcnEchoCounter
	SetStart(value uint32) PatternFlowTcpEcnEchoCounter
	// HasStart checks if Start has been set in PatternFlowTcpEcnEchoCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpEcnEchoCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpEcnEchoCounter
	SetStep(value uint32) PatternFlowTcpEcnEchoCounter
	// HasStep checks if Step has been set in PatternFlowTcpEcnEchoCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpEcnEchoCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpEcnEchoCounter
	SetCount(value uint32) PatternFlowTcpEcnEchoCounter
	// HasCount checks if Count has been set in PatternFlowTcpEcnEchoCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpEcnEchoCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpEcnEchoCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpEcnEchoCounter object
func (obj *patternFlowTcpEcnEchoCounter) SetStart(value uint32) PatternFlowTcpEcnEchoCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpEcnEchoCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpEcnEchoCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpEcnEchoCounter object
func (obj *patternFlowTcpEcnEchoCounter) SetStep(value uint32) PatternFlowTcpEcnEchoCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpEcnEchoCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpEcnEchoCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpEcnEchoCounter object
func (obj *patternFlowTcpEcnEchoCounter) SetCount(value uint32) PatternFlowTcpEcnEchoCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpEcnEchoCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnEchoCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnEchoCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnEchoCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowTcpEcnEchoCounter) setDefault() {
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
