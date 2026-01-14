package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpEcnCwrCounter *****
type patternFlowTcpEcnCwrCounter struct {
	validation
	obj          *otg.PatternFlowTcpEcnCwrCounter
	marshaller   marshalPatternFlowTcpEcnCwrCounter
	unMarshaller unMarshalPatternFlowTcpEcnCwrCounter
}

func NewPatternFlowTcpEcnCwrCounter() PatternFlowTcpEcnCwrCounter {
	obj := patternFlowTcpEcnCwrCounter{obj: &otg.PatternFlowTcpEcnCwrCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpEcnCwrCounter) msg() *otg.PatternFlowTcpEcnCwrCounter {
	return obj.obj
}

func (obj *patternFlowTcpEcnCwrCounter) setMsg(msg *otg.PatternFlowTcpEcnCwrCounter) PatternFlowTcpEcnCwrCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpEcnCwrCounter struct {
	obj *patternFlowTcpEcnCwrCounter
}

type marshalPatternFlowTcpEcnCwrCounter interface {
	// ToProto marshals PatternFlowTcpEcnCwrCounter to protobuf object *otg.PatternFlowTcpEcnCwrCounter
	ToProto() (*otg.PatternFlowTcpEcnCwrCounter, error)
	// ToPbText marshals PatternFlowTcpEcnCwrCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpEcnCwrCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpEcnCwrCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpEcnCwrCounter struct {
	obj *patternFlowTcpEcnCwrCounter
}

type unMarshalPatternFlowTcpEcnCwrCounter interface {
	// FromProto unmarshals PatternFlowTcpEcnCwrCounter from protobuf object *otg.PatternFlowTcpEcnCwrCounter
	FromProto(msg *otg.PatternFlowTcpEcnCwrCounter) (PatternFlowTcpEcnCwrCounter, error)
	// FromPbText unmarshals PatternFlowTcpEcnCwrCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpEcnCwrCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpEcnCwrCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpEcnCwrCounter) Marshal() marshalPatternFlowTcpEcnCwrCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpEcnCwrCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpEcnCwrCounter) Unmarshal() unMarshalPatternFlowTcpEcnCwrCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpEcnCwrCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpEcnCwrCounter) ToProto() (*otg.PatternFlowTcpEcnCwrCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpEcnCwrCounter) FromProto(msg *otg.PatternFlowTcpEcnCwrCounter) (PatternFlowTcpEcnCwrCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpEcnCwrCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnCwrCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpEcnCwrCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnCwrCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpEcnCwrCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnCwrCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpEcnCwrCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnCwrCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnCwrCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpEcnCwrCounter) Clone() (PatternFlowTcpEcnCwrCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpEcnCwrCounter()
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

// PatternFlowTcpEcnCwrCounter is integer counter pattern
type PatternFlowTcpEcnCwrCounter interface {
	Validation
	// msg marshals PatternFlowTcpEcnCwrCounter to protobuf object *otg.PatternFlowTcpEcnCwrCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpEcnCwrCounter
	// setMsg unmarshals PatternFlowTcpEcnCwrCounter from protobuf object *otg.PatternFlowTcpEcnCwrCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpEcnCwrCounter) PatternFlowTcpEcnCwrCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpEcnCwrCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpEcnCwrCounter
	// validate validates PatternFlowTcpEcnCwrCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpEcnCwrCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpEcnCwrCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpEcnCwrCounter
	SetStart(value uint32) PatternFlowTcpEcnCwrCounter
	// HasStart checks if Start has been set in PatternFlowTcpEcnCwrCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpEcnCwrCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpEcnCwrCounter
	SetStep(value uint32) PatternFlowTcpEcnCwrCounter
	// HasStep checks if Step has been set in PatternFlowTcpEcnCwrCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpEcnCwrCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpEcnCwrCounter
	SetCount(value uint32) PatternFlowTcpEcnCwrCounter
	// HasCount checks if Count has been set in PatternFlowTcpEcnCwrCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpEcnCwrCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpEcnCwrCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpEcnCwrCounter object
func (obj *patternFlowTcpEcnCwrCounter) SetStart(value uint32) PatternFlowTcpEcnCwrCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpEcnCwrCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpEcnCwrCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpEcnCwrCounter object
func (obj *patternFlowTcpEcnCwrCounter) SetStep(value uint32) PatternFlowTcpEcnCwrCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpEcnCwrCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpEcnCwrCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpEcnCwrCounter object
func (obj *patternFlowTcpEcnCwrCounter) SetCount(value uint32) PatternFlowTcpEcnCwrCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpEcnCwrCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnCwrCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnCwrCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnCwrCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowTcpEcnCwrCounter) setDefault() {
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
