package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpEcnNsCounter *****
type patternFlowTcpEcnNsCounter struct {
	validation
	obj          *otg.PatternFlowTcpEcnNsCounter
	marshaller   marshalPatternFlowTcpEcnNsCounter
	unMarshaller unMarshalPatternFlowTcpEcnNsCounter
}

func NewPatternFlowTcpEcnNsCounter() PatternFlowTcpEcnNsCounter {
	obj := patternFlowTcpEcnNsCounter{obj: &otg.PatternFlowTcpEcnNsCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpEcnNsCounter) msg() *otg.PatternFlowTcpEcnNsCounter {
	return obj.obj
}

func (obj *patternFlowTcpEcnNsCounter) setMsg(msg *otg.PatternFlowTcpEcnNsCounter) PatternFlowTcpEcnNsCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpEcnNsCounter struct {
	obj *patternFlowTcpEcnNsCounter
}

type marshalPatternFlowTcpEcnNsCounter interface {
	// ToProto marshals PatternFlowTcpEcnNsCounter to protobuf object *otg.PatternFlowTcpEcnNsCounter
	ToProto() (*otg.PatternFlowTcpEcnNsCounter, error)
	// ToPbText marshals PatternFlowTcpEcnNsCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpEcnNsCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpEcnNsCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpEcnNsCounter struct {
	obj *patternFlowTcpEcnNsCounter
}

type unMarshalPatternFlowTcpEcnNsCounter interface {
	// FromProto unmarshals PatternFlowTcpEcnNsCounter from protobuf object *otg.PatternFlowTcpEcnNsCounter
	FromProto(msg *otg.PatternFlowTcpEcnNsCounter) (PatternFlowTcpEcnNsCounter, error)
	// FromPbText unmarshals PatternFlowTcpEcnNsCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpEcnNsCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpEcnNsCounter from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpEcnNsCounter) Marshal() marshalPatternFlowTcpEcnNsCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpEcnNsCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpEcnNsCounter) Unmarshal() unMarshalPatternFlowTcpEcnNsCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpEcnNsCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpEcnNsCounter) ToProto() (*otg.PatternFlowTcpEcnNsCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpEcnNsCounter) FromProto(msg *otg.PatternFlowTcpEcnNsCounter) (PatternFlowTcpEcnNsCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpEcnNsCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnNsCounter) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpEcnNsCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnNsCounter) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpEcnNsCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnNsCounter) FromJson(value string) error {
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

func (obj *patternFlowTcpEcnNsCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnNsCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnNsCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpEcnNsCounter) Clone() (PatternFlowTcpEcnNsCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpEcnNsCounter()
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

// PatternFlowTcpEcnNsCounter is integer counter pattern
type PatternFlowTcpEcnNsCounter interface {
	Validation
	// msg marshals PatternFlowTcpEcnNsCounter to protobuf object *otg.PatternFlowTcpEcnNsCounter
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpEcnNsCounter
	// setMsg unmarshals PatternFlowTcpEcnNsCounter from protobuf object *otg.PatternFlowTcpEcnNsCounter
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpEcnNsCounter) PatternFlowTcpEcnNsCounter
	// provides marshal interface
	Marshal() marshalPatternFlowTcpEcnNsCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpEcnNsCounter
	// validate validates PatternFlowTcpEcnNsCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpEcnNsCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternFlowTcpEcnNsCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternFlowTcpEcnNsCounter
	SetStart(value uint32) PatternFlowTcpEcnNsCounter
	// HasStart checks if Start has been set in PatternFlowTcpEcnNsCounter
	HasStart() bool
	// Step returns uint32, set in PatternFlowTcpEcnNsCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternFlowTcpEcnNsCounter
	SetStep(value uint32) PatternFlowTcpEcnNsCounter
	// HasStep checks if Step has been set in PatternFlowTcpEcnNsCounter
	HasStep() bool
	// Count returns uint32, set in PatternFlowTcpEcnNsCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternFlowTcpEcnNsCounter
	SetCount(value uint32) PatternFlowTcpEcnNsCounter
	// HasCount checks if Count has been set in PatternFlowTcpEcnNsCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpEcnNsCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternFlowTcpEcnNsCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternFlowTcpEcnNsCounter object
func (obj *patternFlowTcpEcnNsCounter) SetStart(value uint32) PatternFlowTcpEcnNsCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpEcnNsCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternFlowTcpEcnNsCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternFlowTcpEcnNsCounter object
func (obj *patternFlowTcpEcnNsCounter) SetStep(value uint32) PatternFlowTcpEcnNsCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpEcnNsCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternFlowTcpEcnNsCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternFlowTcpEcnNsCounter object
func (obj *patternFlowTcpEcnNsCounter) SetCount(value uint32) PatternFlowTcpEcnNsCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternFlowTcpEcnNsCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnNsCounter.Start <= 1 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnNsCounter.Step <= 1 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnNsCounter.Count <= 2 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternFlowTcpEcnNsCounter) setDefault() {
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
